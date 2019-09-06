package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/mocks"
	"github.com/utilitywarehouse/kafka-topic-applier/internal/service"
)

type testSuite struct {
	ctx  context.Context
	svc  *service.Service
	mock *mockz
}

type mockz struct {
	cli *mocks.MockClusterAdmin
}

func testSetup(t *testing.T) (testSuite, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	ca := mocks.NewMockClusterAdmin(ctrl)

	svc := service.New(ca)

	return testSuite{
		ctx,
		svc,
		&mockz{},
	}, ctrl
}

func TestSuccessfullyCreate(t *testing.T) {
	_, ctrl := testSetup(t)
	defer ctrl.Finish()
}
