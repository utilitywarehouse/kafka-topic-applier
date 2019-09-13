package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/ptypes/empty"
	cli "github.com/jawher/mow.cli"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/pb/kta"
	"google.golang.org/grpc"
)

const (
	appDescription = "App that produces printable invoice models"
	appName        = "client"
)

func main() {
	app := cli.App(appName, appDescription)

	grpcPort := app.Int(cli.IntOpt{
		Name:   "grpc-port",
		Desc:   "The port to listen on for API GRPC connections",
		Value:  8090,
		EnvVar: "GRPC_PORT",
	})
	ctx := context.Background()

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", *grpcPort), grpc.WithInsecure())
	if err != nil {
		log.Fatal("conn err", err)
	}

	grpcClient := kta.NewTopicBotClient(conn)
	defer conn.Close()

	app.Command("list", "list topics", func(cmd *cli.Cmd) {
		useYAML := cmd.Bool(cli.BoolOpt{
			Name:  "yaml",
			Desc:  "Show YAML results",
			Value: false,
		})

		cmd.Action = func() {
			resp, err := grpcClient.List(ctx, &empty.Empty{})
			if err != nil {
				log.Fatal("failed to run command", err)
			}

			if *useYAML {
				y, err := yaml.Marshal(resp)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(y))
			} else {
				fmt.Println(resp)
			}
		}
	})

	app.Command("delete", "delete topic", func(cmd *cli.Cmd) {
		name := cmd.String(cli.StringOpt{
			Name: "topic",
			Desc: "Topic Name",
		})

		cmd.Action = func() {
			resp, err := grpcClient.Delete(ctx, &kta.Topic{Name: *name})
			if err != nil {
				log.Fatal("failed to run command", err)
			}
			fmt.Println(resp)
		}
	})

	app.Command("create", "create topic", func(cmd *cli.Cmd) {
		name := cmd.String(cli.StringOpt{
			Name: "topic",
			Desc: "Topic Name",
		})
		partitions := cmd.Int(cli.IntOpt{
			Name:   "partitions",
			Desc:   "Number of partitions to create the topic in",
			EnvVar: "PARTITIONS",
			Value:  15,
		})
		replicationFactor := cmd.Int(cli.IntOpt{
			Name:   "replication-factor",
			Desc:   "How many times to replicate the topics",
			EnvVar: "REPLICATION_FACTOR",
			Value:  1, // must not be greater than available brokers, otherwise will fail
		})
		topicCompression := cmd.String(cli.StringOpt{
			Name:   "topic-compression-type",
			Desc:   "Compression type for topic",
			EnvVar: "TOPIC_COMPRESSION",
		})
		maxMessageBytes := cmd.String(cli.StringOpt{
			Name:   "topic-max-message-bytes",
			Desc:   "The largest record batch size allowed on topic",
			EnvVar: "TOPIC_MAX_MESSAGE_BYTES",
		})
		maxRetentionBytes := cmd.String(cli.StringOpt{
			Name:   "topic-max-retention-bytes",
			Desc:   "The highest retention size allowed on topic",
			EnvVar: "TOPIC_MAX_RETENTION_BYTES",
		})
		maxRetentionMS := cmd.String(cli.StringOpt{
			Name:   "topic-max-retention-ms",
			Desc:   "The highest retention time (ms) allowed on topic",
			EnvVar: "TOPIC_MAX_RETENTION_MS",
		})
		cleanupPolicyType := cmd.String(cli.StringOpt{
			Name:   "topic-cleanup-policy-type",
			Desc:   "Defines how the logs are stored [compact, delete]",
			EnvVar: "TOPIC_CLEANUP_POLICY_TYPE",
		})

		cmd.Action = func() {
			resp, err := grpcClient.Create(ctx, &kta.Topic{
				Name:              *name,
				Partitions:        int32(*partitions),
				ReplicationFactor: int32(*replicationFactor),
				TopicCompression:  *topicCompression,
				MaxMessageBytes:   *maxMessageBytes,
				MaxRetentionBytes: *maxRetentionBytes,
				MaxRetentionTime:  *maxRetentionMS,
				CleanupPolicy:     *cleanupPolicyType,
			})
			if err != nil {
				log.Fatal("failed to run command", err)
			}
			fmt.Println(resp)
		}
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal("failed to run app", err)
	}

}
