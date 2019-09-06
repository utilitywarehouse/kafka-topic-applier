package service

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/pb/kta"
)

// NewService creates a new service
//go:generate mockgen -package=mocks -destination=../mocks/clusterAdmin.go github.com/Shopify/sarama ClusterAdmin
func New(clusterAdmin sarama.ClusterAdmin) *Service {
	return &Service{
		clusterAdmin: clusterAdmin,
	}
}

// Service implements the TopicBot service
type Service struct {
	clusterAdmin sarama.ClusterAdmin
}

// Create adds a new topic to Kafka if it doesn't already exist
func (s *Service) Create(ctx context.Context, req *kta.Topic) (*empty.Empty, error) {
	return &empty.Empty{}, s.clusterAdmin.CreateTopic(req.Name, mapTopicToTopicDetail(req), false)
}

// List return existing Topics
func (s *Service) List(ctx context.Context, _ *empty.Empty) (*kta.Topics, error) {
	var topics kta.Topics
	tt, err := s.clusterAdmin.ListTopics()
	if err != nil {
		return nil, err
	}

	for i, t := range tt {
		topics.Topics = append(topics.Topics, mapTopicDetailToTopic(i, t))
	}

	return &topics, nil
}

// Delete removes a topic from Kafka
func (s *Service) Delete(ctx context.Context, req *kta.Topic) (*empty.Empty, error) {
	return &empty.Empty{}, s.clusterAdmin.DeleteTopic(req.Name)
}

func mapTopicDetailToTopic(n string, td sarama.TopicDetail) *kta.Topic {
	return &kta.Topic{
		Name:              n,
		Partitions:        td.NumPartitions,
		ReplicationFactor: int32(td.ReplicationFactor),
		TopicCompression:  fixNullString(td.ConfigEntries["compresion.type"]),
		MaxMessageBytes:   fixNullString(td.ConfigEntries["max.message.bytes"]),
		MaxRetentionBytes: fixNullString(td.ConfigEntries["retention.bytes"]),
		MaxRetentionTime:  fixNullString(td.ConfigEntries["retention.ms"]),
		CleanupPolicy:     fixNullString(td.ConfigEntries["cleanup.policy"]),
	}
}

func mapTopicToTopicDetail(t *kta.Topic) *sarama.TopicDetail {
	topicConfig := map[string]*string{}
	if t.TopicCompression != "" {
		topicConfig["compression.type"] = &t.TopicCompression
	}
	if t.MaxMessageBytes != "" {
		topicConfig["max.message.bytes"] = &t.MaxMessageBytes
	}
	if t.MaxRetentionBytes != "" {
		topicConfig["retention.bytes"] = &t.MaxRetentionBytes
	}
	if t.MaxRetentionTime != "" {
		topicConfig["retention.ms"] = &t.MaxRetentionTime
	}
	if t.CleanupPolicy != "" {
		topicConfig["cleanup.policy"] = &t.CleanupPolicy
	}

	return &sarama.TopicDetail{
		NumPartitions:     t.Partitions,
		ReplicationFactor: int16(t.ReplicationFactor),
		ConfigEntries:     topicConfig,
	}
}

func fixNullString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
