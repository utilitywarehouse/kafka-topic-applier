package apply

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/ghodss/yaml"
	"google.golang.org/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/pb/kta"
)

type Applier struct {
	svc kta.TopicBotClient
}

func NewApplier(svc kta.TopicBotClient) *Applier {
	return &Applier{svc: svc}
}

func (l *Applier) Apply(filePath string) error {
	ctx := context.Background()
	existingTopics, err := l.svc.List(ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	topicExists := make(map[string]bool)

	for _, t := range existingTopics.Topics {
		topicExists[t.Name] = true
	}

	tt, err := getTopicsFromFile(filePath)
	if err != nil {
		return err
	}

	for _, t := range tt.Topics {
		switch {
		case topicExists[t.Name] && !t.ShouldBeRemoved:
			continue
		case !topicExists[t.Name] && t.ShouldBeRemoved:
			continue
		case t.ShouldBeRemoved:
			log.Printf("removing topic %s\n", t.Name)
			_, err := l.svc.Delete(ctx, t)
			if err != nil {
				log.Print(err)
			}
		default:
			log.Printf("creating topic %s\n", t.Name)
			_, err := l.svc.Create(ctx, t)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getTopicsFromFile(filePath string) (*kta.Topics, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "could not read file")
	}

	tt := kta.Topics{}
	if err := yaml.Unmarshal(b, &tt); err != nil {
		return nil, errors.Wrap(err, "could not de-serialise topics")
	}

	return &tt, nil
}
