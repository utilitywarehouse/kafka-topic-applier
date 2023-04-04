package service_test

import (
	"context"
	"log"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/ptypes/empty"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/mocks"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/pb/kta"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/service"
)

type testSuite struct {
	ctx  context.Context
	svc  *service.Service
	mock *mockz
}

type mockz struct {
	ca *mocks.MockClusterAdmin
}

func testSetup(t *testing.T) (testSuite, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	ca := mocks.NewMockClusterAdmin(ctrl)

	svc := service.New(ca)

	return testSuite{
		ctx,
		svc,
		&mockz{ca},
	}, ctrl
}

func TestSuccessfullyCreate(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ts.mock.ca.EXPECT().CreateTopic("name", gomock.Any(), false).Return(nil).Times(1)

	_, err := ts.svc.Create(ts.ctx, &kta.Topic{Name: "name"})
	if err != nil {
		log.Fatal(err)
	}

}

func TestSuccessfullyList(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	tt := make(map[string]sarama.TopicDetail)
	ts.mock.ca.EXPECT().ListTopics().Return(tt, nil).Times(1)

	_, err := ts.svc.List(ts.ctx, &empty.Empty{})
	if err != nil {
		log.Fatal(err)
	}
}

func TestSuccessfullyDelete(t *testing.T) {
	ts, ctrl := testSetup(t)
	defer ctrl.Finish()

	ts.mock.ca.EXPECT().DeleteTopic("name").Return(nil).Times(1)

	_, err := ts.svc.Delete(ts.ctx, &kta.Topic{Name: "name"})
	if err != nil {
		log.Fatal(err)
	}
}
