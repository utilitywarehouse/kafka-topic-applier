package service

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/pb/kta"
)

const (
	compressionType = "compression.type"
	maxMessageBytes = "max.message.bytes"
	retentionBytes  = "retention.bytes"
	retentionMS     = "retention.ms"
	cleanupPolicy   = "cleanup.policy"
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
	err := s.clusterAdmin.CreateTopic(req.Name, mapTopicToTopicDetail(req), false)
	if err != nil {
		return nil, fmt.Errorf("failed to create topic %s, err: %v", req.Name, err)
	}
	return &empty.Empty{}, nil
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
	err := s.clusterAdmin.DeleteTopic(req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to delete topic %s, err: %v", req.Name, err)
	}
	return &empty.Empty{}, nil
}

func mapTopicDetailToTopic(n string, td sarama.TopicDetail) *kta.Topic {
	return &kta.Topic{
		Name:              n,
		Partitions:        td.NumPartitions,
		ReplicationFactor: int32(td.ReplicationFactor),
		TopicCompression:  fixNullString(td.ConfigEntries[compressionType]),
		MaxMessageBytes:   fixNullString(td.ConfigEntries[maxMessageBytes]),
		MaxRetentionBytes: fixNullString(td.ConfigEntries[retentionBytes]),
		MaxRetentionTime:  fixNullString(td.ConfigEntries[retentionMS]),
		CleanupPolicy:     fixNullString(td.ConfigEntries[cleanupPolicy]),
	}
}

func mapTopicToTopicDetail(t *kta.Topic) *sarama.TopicDetail {
	topicConfig := map[string]*string{}
	if t.TopicCompression != "" {
		topicConfig[compressionType] = &t.TopicCompression
	}
	if t.MaxMessageBytes != "" {
		topicConfig[maxMessageBytes] = &t.MaxMessageBytes
	}
	if t.MaxRetentionBytes != "" {
		topicConfig[retentionBytes] = &t.MaxRetentionBytes
	}
	if t.MaxRetentionTime != "" {
		topicConfig[retentionMS] = &t.MaxRetentionTime
	}
	if t.CleanupPolicy != "" {
		topicConfig[cleanupPolicy] = &t.CleanupPolicy
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
