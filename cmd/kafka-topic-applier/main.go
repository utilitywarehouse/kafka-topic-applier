package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/apply"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/pb/kta"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"

	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
	"github.com/utilitywarehouse/go-operational/op"
)

const (
	appName = "kafka-topic-applier"
	appDesc = "Maintain Kafka Topics"
)

var gitHash string // populated at compile time

func main() {
	app := cli.App(appName, appDesc)

	opsPort := app.Int(cli.IntOpt{
		Name:   "ops-port",
		Desc:   "The HTTP ops port",
		EnvVar: "OPS_PORT",
		Value:  8081,
	})
	grpcPort := app.Int(cli.IntOpt{
		Name:   "grpc-port",
		Desc:   "The port to listen on for API GRPC connections",
		Value:  8090,
		EnvVar: "GRPC_PORT",
	})
	logLevel := app.String(cli.StringOpt{
		Name:   "log-level",
		Desc:   "log level [debug|info|warn|error]",
		EnvVar: "LOG_LEVEL",
		Value:  "debug",
	})
	logFormat := app.String(cli.StringOpt{
		Name:   "log-format",
		Desc:   "Log format, if set to text will use text as logging format, otherwise will use json",
		EnvVar: "LOG_FORMAT",
		Value:  "json",
	})
	team := app.String(cli.StringOpt{
		Name:   "team",
		Desc:   "Who's running this",
		EnvVar: "TEAM",
		Value:  "billing",
	})
	kafkaBrokers := app.String(cli.StringOpt{
		Name:   "kafka-brokers",
		Desc:   "CSV delimited list of brokers",
		EnvVar: "KAFKA_BROKERS",
		Value:  "localhost:9092",
	})
	kafkaVersion := app.String(cli.StringOpt{
		Name:   "kafka-version",
		Desc:   "Version Kafka is running on",
		EnvVar: "KAFKA_VERSION",
		Value:  "2.0.1",
	})
	topicsFile := app.String(cli.StringOpt{
		Name:   "topics-file",
		Desc:   "yaml source of topics",
		EnvVar: "TOPICS_FILE",
		Value:  "/tmp/topics.yaml",
	})
	kafkaTimeout := app.Int(cli.IntOpt{
		Name:   "kafka-timeout",
		Desc:   "How long to wait for kafka to timeout in seconds",
		Value:  3,
		EnvVar: "KAFKA_TIMEOUT",
	})

	app.Action = func() {
		configureLogger(*logLevel, *logFormat)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		config := sarama.NewConfig()
		config.Version = getKafkaVersion(*kafkaVersion)
		config.Admin.Timeout = time.Duration(1000000000 * *kafkaTimeout)

		ca, err := sarama.NewClusterAdmin(strings.Split(*kafkaBrokers, ","), config)
		if err != nil {
			log.Fatal("failed to create cluster admin: ", err)
		}
		defer ca.Close()

		svc := service.New(ca)

		grpcLogger := configureGrpcLogger()

		grpcServer := initialiseGRPCServer(svc, grpcLogger)
		go startGRPCServer(grpcServer, grpcPort)
		defer grpcServer.GracefulStop()

		opsServer := initialiseOpsServer(opsPort, team)
		go startOpsServer(opsServer)
		defer opsServer.Shutdown(context.Background())

		localSvc := fmt.Sprintf("localhost:%d", *grpcPort)
		svcGrpcClientConn := initialiseGRPCClientConnection(ctx, &localSvc)
		defer svcGrpcClientConn.Close()

		svcCli := kta.NewTopicBotClient(svcGrpcClientConn)

		a := apply.NewApplier(svcCli)
		a.Apply(*topicsFile)

		waitForShutdown()
	}

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Panic("unable to run app")
	}
}

func initialiseOpsServer(opsPort *int, team *string) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", *opsPort),
		Handler: newOpHandler(team),
	}
}

func startOpsServer(opsServer *http.Server) {
	opsErr := opsServer.ListenAndServe()
	switch opsErr {
	case http.ErrServerClosed:
		log.WithError(opsErr).Warn("ops server shutdown")
	default:
		log.WithError(opsErr).Panic("unable to start ops http server")
	}
}

func initialiseGRPCServer(svc kta.TopicBotServer, logger *log.Logger) *grpc.Server {
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(log.NewEntry(logger)),
		)),
	)
	kta.RegisterTopicBotServer(grpcServer, svc)

	healthpb.RegisterHealthServer(grpcServer, health.NewServer())
	return grpcServer
}

func startGRPCServer(grpcServer *grpc.Server, grpcPort *int) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.WithFields(log.Fields{"port": *grpcPort}).
			WithError(err).
			Panic("failed to listen on GRPC port")
	}
	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.WithError(err).Panic("failed to serve GRPC connections")
		}
	}()
}

func newOpHandler(team *string) http.Handler {
	return op.NewHandler(op.
		NewStatus(appName, appDesc).
		AddOwner(*team, fmt.Sprintf("#%s", *team)).
		SetRevision(gitHash).
		ReadyAlways())
}

func configureLogger(level, format string) {
	l, err := log.ParseLevel(level)
	if err != nil {
		log.WithFields(log.Fields{"log_level": level}).
			WithError(err).
			Panic("invalid log level")
	}
	log.SetLevel(l)

	format = strings.ToLower(format)
	if format != "text" && format != "json" {
		log.Panicf("invalid log format: %s", format)
	}
	if format == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func configureGrpcLogger() *log.Logger {
	l := log.New()
	l.SetLevel(log.WarnLevel)
	grpc_logrus.ReplaceGrpcLogger(log.NewEntry(l))
	return l
}

func waitForShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Warn("shutting down")
}

func initialiseGRPCClientConnection(ctx context.Context, grpcClientAddress *string) *grpc.ClientConn {
	grpcClientConn, err := grpc.DialContext(ctx,
		*grpcClientAddress,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			[]grpc_retry.CallOption{
				grpc_retry.WithBackoff(grpc_retry.BackoffLinearWithJitter(100*time.Millisecond, 0.1)),
				grpc_retry.WithMax(3),
				grpc_retry.WithCodes(codes.Unknown, codes.DeadlineExceeded, codes.Internal, codes.Unavailable),
			}...,
		)),
	)
	if err != nil {
		log.WithFields(log.Fields{"grpc_client_address": *grpcClientAddress}).
			WithError(err).
			Panic("grpc client connection failed")
	}
	return grpcClientConn
}

func getKafkaVersion(v string) sarama.KafkaVersion {
	version, err := sarama.ParseKafkaVersion(v)
	if err != nil {
		log.WithError(err).Panic("error parsing version")
	}

	return version
}
