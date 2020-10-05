// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Shopify/sarama (interfaces: ClusterAdmin)

// Package mocks is a generated GoMock package.
package mocks

import (
	sarama "github.com/Shopify/sarama"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClusterAdmin is a mock of ClusterAdmin interface
type MockClusterAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockClusterAdminMockRecorder
}

// MockClusterAdminMockRecorder is the mock recorder for MockClusterAdmin
type MockClusterAdminMockRecorder struct {
	mock *MockClusterAdmin
}

// NewMockClusterAdmin creates a new mock instance
func NewMockClusterAdmin(ctrl *gomock.Controller) *MockClusterAdmin {
	mock := &MockClusterAdmin{ctrl: ctrl}
	mock.recorder = &MockClusterAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterAdmin) EXPECT() *MockClusterAdminMockRecorder {
	return m.recorder
}

// AlterConfig mocks base method
func (m *MockClusterAdmin) AlterConfig(arg0 sarama.ConfigResourceType, arg1 string, arg2 map[string]*string, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlterConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AlterConfig indicates an expected call of AlterConfig
func (mr *MockClusterAdminMockRecorder) AlterConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlterConfig", reflect.TypeOf((*MockClusterAdmin)(nil).AlterConfig), arg0, arg1, arg2, arg3)
}

// AlterPartitionReassignments mocks base method
func (m *MockClusterAdmin) AlterPartitionReassignments(arg0 string, arg1 [][]int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlterPartitionReassignments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AlterPartitionReassignments indicates an expected call of AlterPartitionReassignments
func (mr *MockClusterAdminMockRecorder) AlterPartitionReassignments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlterPartitionReassignments", reflect.TypeOf((*MockClusterAdmin)(nil).AlterPartitionReassignments), arg0, arg1)
}

// Close mocks base method
func (m *MockClusterAdmin) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockClusterAdminMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClusterAdmin)(nil).Close))
}

// CreateACL mocks base method
func (m *MockClusterAdmin) CreateACL(arg0 sarama.Resource, arg1 sarama.Acl) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateACL", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateACL indicates an expected call of CreateACL
func (mr *MockClusterAdminMockRecorder) CreateACL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateACL", reflect.TypeOf((*MockClusterAdmin)(nil).CreateACL), arg0, arg1)
}

// CreatePartitions mocks base method
func (m *MockClusterAdmin) CreatePartitions(arg0 string, arg1 int32, arg2 [][]int32, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePartitions", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePartitions indicates an expected call of CreatePartitions
func (mr *MockClusterAdminMockRecorder) CreatePartitions(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePartitions", reflect.TypeOf((*MockClusterAdmin)(nil).CreatePartitions), arg0, arg1, arg2, arg3)
}

// CreateTopic mocks base method
func (m *MockClusterAdmin) CreateTopic(arg0 string, arg1 *sarama.TopicDetail, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopic", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTopic indicates an expected call of CreateTopic
func (mr *MockClusterAdminMockRecorder) CreateTopic(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopic", reflect.TypeOf((*MockClusterAdmin)(nil).CreateTopic), arg0, arg1, arg2)
}

// DeleteACL mocks base method
func (m *MockClusterAdmin) DeleteACL(arg0 sarama.AclFilter, arg1 bool) ([]sarama.MatchingAcl, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteACL", arg0, arg1)
	ret0, _ := ret[0].([]sarama.MatchingAcl)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteACL indicates an expected call of DeleteACL
func (mr *MockClusterAdminMockRecorder) DeleteACL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteACL", reflect.TypeOf((*MockClusterAdmin)(nil).DeleteACL), arg0, arg1)
}

// DeleteConsumerGroup mocks base method
func (m *MockClusterAdmin) DeleteConsumerGroup(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConsumerGroup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConsumerGroup indicates an expected call of DeleteConsumerGroup
func (mr *MockClusterAdminMockRecorder) DeleteConsumerGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConsumerGroup", reflect.TypeOf((*MockClusterAdmin)(nil).DeleteConsumerGroup), arg0)
}

// DeleteRecords mocks base method
func (m *MockClusterAdmin) DeleteRecords(arg0 string, arg1 map[int32]int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecords", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecords indicates an expected call of DeleteRecords
func (mr *MockClusterAdminMockRecorder) DeleteRecords(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecords", reflect.TypeOf((*MockClusterAdmin)(nil).DeleteRecords), arg0, arg1)
}

// DeleteTopic mocks base method
func (m *MockClusterAdmin) DeleteTopic(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTopic", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTopic indicates an expected call of DeleteTopic
func (mr *MockClusterAdminMockRecorder) DeleteTopic(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTopic", reflect.TypeOf((*MockClusterAdmin)(nil).DeleteTopic), arg0)
}

// DescribeCluster mocks base method
func (m *MockClusterAdmin) DescribeCluster() ([]*sarama.Broker, int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCluster")
	ret0, _ := ret[0].([]*sarama.Broker)
	ret1, _ := ret[1].(int32)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DescribeCluster indicates an expected call of DescribeCluster
func (mr *MockClusterAdminMockRecorder) DescribeCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCluster", reflect.TypeOf((*MockClusterAdmin)(nil).DescribeCluster))
}

// DescribeConfig mocks base method
func (m *MockClusterAdmin) DescribeConfig(arg0 sarama.ConfigResource) ([]sarama.ConfigEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeConfig", arg0)
	ret0, _ := ret[0].([]sarama.ConfigEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConfig indicates an expected call of DescribeConfig
func (mr *MockClusterAdminMockRecorder) DescribeConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConfig", reflect.TypeOf((*MockClusterAdmin)(nil).DescribeConfig), arg0)
}

// DescribeConsumerGroups mocks base method
func (m *MockClusterAdmin) DescribeConsumerGroups(arg0 []string) ([]*sarama.GroupDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeConsumerGroups", arg0)
	ret0, _ := ret[0].([]*sarama.GroupDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConsumerGroups indicates an expected call of DescribeConsumerGroups
func (mr *MockClusterAdminMockRecorder) DescribeConsumerGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConsumerGroups", reflect.TypeOf((*MockClusterAdmin)(nil).DescribeConsumerGroups), arg0)
}

// DescribeLogDirs mocks base method
func (m *MockClusterAdmin) DescribeLogDirs(arg0 []int32) (map[int32][]sarama.DescribeLogDirsResponseDirMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLogDirs", arg0)
	ret0, _ := ret[0].(map[int32][]sarama.DescribeLogDirsResponseDirMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLogDirs indicates an expected call of DescribeLogDirs
func (mr *MockClusterAdminMockRecorder) DescribeLogDirs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLogDirs", reflect.TypeOf((*MockClusterAdmin)(nil).DescribeLogDirs), arg0)
}

// DescribeTopics mocks base method
func (m *MockClusterAdmin) DescribeTopics(arg0 []string) ([]*sarama.TopicMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTopics", arg0)
	ret0, _ := ret[0].([]*sarama.TopicMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTopics indicates an expected call of DescribeTopics
func (mr *MockClusterAdminMockRecorder) DescribeTopics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTopics", reflect.TypeOf((*MockClusterAdmin)(nil).DescribeTopics), arg0)
}

// ListAcls mocks base method
func (m *MockClusterAdmin) ListAcls(arg0 sarama.AclFilter) ([]sarama.ResourceAcls, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAcls", arg0)
	ret0, _ := ret[0].([]sarama.ResourceAcls)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAcls indicates an expected call of ListAcls
func (mr *MockClusterAdminMockRecorder) ListAcls(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAcls", reflect.TypeOf((*MockClusterAdmin)(nil).ListAcls), arg0)
}

// ListConsumerGroupOffsets mocks base method
func (m *MockClusterAdmin) ListConsumerGroupOffsets(arg0 string, arg1 map[string][]int32) (*sarama.OffsetFetchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConsumerGroupOffsets", arg0, arg1)
	ret0, _ := ret[0].(*sarama.OffsetFetchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConsumerGroupOffsets indicates an expected call of ListConsumerGroupOffsets
func (mr *MockClusterAdminMockRecorder) ListConsumerGroupOffsets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConsumerGroupOffsets", reflect.TypeOf((*MockClusterAdmin)(nil).ListConsumerGroupOffsets), arg0, arg1)
}

// ListConsumerGroups mocks base method
func (m *MockClusterAdmin) ListConsumerGroups() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConsumerGroups")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConsumerGroups indicates an expected call of ListConsumerGroups
func (mr *MockClusterAdminMockRecorder) ListConsumerGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConsumerGroups", reflect.TypeOf((*MockClusterAdmin)(nil).ListConsumerGroups))
}

// ListPartitionReassignments mocks base method
func (m *MockClusterAdmin) ListPartitionReassignments(arg0 string, arg1 []int32) (map[string]map[int32]*sarama.PartitionReplicaReassignmentsStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPartitionReassignments", arg0, arg1)
	ret0, _ := ret[0].(map[string]map[int32]*sarama.PartitionReplicaReassignmentsStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPartitionReassignments indicates an expected call of ListPartitionReassignments
func (mr *MockClusterAdminMockRecorder) ListPartitionReassignments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPartitionReassignments", reflect.TypeOf((*MockClusterAdmin)(nil).ListPartitionReassignments), arg0, arg1)
}

// ListTopics mocks base method
func (m *MockClusterAdmin) ListTopics() (map[string]sarama.TopicDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTopics")
	ret0, _ := ret[0].(map[string]sarama.TopicDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTopics indicates an expected call of ListTopics
func (mr *MockClusterAdminMockRecorder) ListTopics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTopics", reflect.TypeOf((*MockClusterAdmin)(nil).ListTopics))
}
