package apply

import (
	"context"
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/pb/kta"
	"github.com/utilitywarehouse/uwgolib/log"
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
	existingTopicMap := make(map[string]bool)

	for _, t := range existingTopics.Topics {
		existingTopicMap[t.Name] = true
	}

	tt, err := getTopicsFromFile(filePath)
	if err != nil {
		return err
	}

	for _, t := range tt.Topics {
		switch {
		case existingTopicMap[t.Name] && !t.ShouldBeRemoved:
			continue
		case t.ShouldBeRemoved:
			log.Infof("removing topic %s", t.Name)
			_, err := l.svc.Delete(ctx, t)
			if err != nil {
				return err
			}
		default:
			log.Infof("creating topic %s", t.Name)
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
