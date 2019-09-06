// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kta.proto

package kta

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Topics struct {
	Topics []*Topic `protobuf:"bytes,3,rep,name=topics,proto3" json:"topics,omitempty" yaml:"topics"`
}

func (m *Topics) Reset()      { *m = Topics{} }
func (*Topics) ProtoMessage() {}
func (*Topics) Descriptor() ([]byte, []int) {
	return fileDescriptor_d23c639e4ed5e2ef, []int{0}
}
func (m *Topics) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Topics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Topics.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Topics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topics.Merge(m, src)
}
func (m *Topics) XXX_Size() int {
	return m.Size()
}
func (m *Topics) XXX_DiscardUnknown() {
	xxx_messageInfo_Topics.DiscardUnknown(m)
}

var xxx_messageInfo_Topics proto.InternalMessageInfo

func (m *Topics) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

type Topic struct {
	Name              string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	Partitions        int32  `protobuf:"varint,2,opt,name=partitions,proto3" json:"partitions,omitempty" yaml:"partitions"`
	ReplicationFactor int32  `protobuf:"varint,3,opt,name=replication_factor,json=replicationFactor,proto3" json:"replication_factor,omitempty" yaml:"replication_factor"`
	TopicCompression  string `protobuf:"bytes,4,opt,name=topic_compression,json=topicCompression,proto3" json:"topic_compression,omitempty" yaml:"topic_compression"`
	MaxMessageBytes   string `protobuf:"bytes,5,opt,name=max_message_bytes,json=maxMessageBytes,proto3" json:"max_message_bytes,omitempty" yaml:"max_message_bytes"`
	MaxRetentionBytes string `protobuf:"bytes,6,opt,name=max_retention_bytes,json=maxRetentionBytes,proto3" json:"max_retention_bytes,omitempty" yaml:"max_retention_bytes"`
	MaxRetentionTime  string `protobuf:"bytes,7,opt,name=max_retention_time,json=maxRetentionTime,proto3" json:"max_retention_time,omitempty" yaml:"max_retention_time"`
	CleanupPolicy     string `protobuf:"bytes,8,opt,name=cleanup_policy,json=cleanupPolicy,proto3" json:"cleanup_policy,omitempty" yaml:"cleanup_policy"`
	ShouldBeRemoved   bool   `protobuf:"varint,9,opt,name=should_be_removed,json=shouldBeRemoved,proto3" json:"should_be_removed,omitempty" yaml:"should_be_removed"`
}

func (m *Topic) Reset()      { *m = Topic{} }
func (*Topic) ProtoMessage() {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_d23c639e4ed5e2ef, []int{1}
}
func (m *Topic) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return m.Size()
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Topic) GetPartitions() int32 {
	if m != nil {
		return m.Partitions
	}
	return 0
}

func (m *Topic) GetReplicationFactor() int32 {
	if m != nil {
		return m.ReplicationFactor
	}
	return 0
}

func (m *Topic) GetTopicCompression() string {
	if m != nil {
		return m.TopicCompression
	}
	return ""
}

func (m *Topic) GetMaxMessageBytes() string {
	if m != nil {
		return m.MaxMessageBytes
	}
	return ""
}

func (m *Topic) GetMaxRetentionBytes() string {
	if m != nil {
		return m.MaxRetentionBytes
	}
	return ""
}

func (m *Topic) GetMaxRetentionTime() string {
	if m != nil {
		return m.MaxRetentionTime
	}
	return ""
}

func (m *Topic) GetCleanupPolicy() string {
	if m != nil {
		return m.CleanupPolicy
	}
	return ""
}

func (m *Topic) GetShouldBeRemoved() bool {
	if m != nil {
		return m.ShouldBeRemoved
	}
	return false
}

func init() {
	proto.RegisterType((*Topics)(nil), "kta.Topics")
	proto.RegisterType((*Topic)(nil), "kta.Topic")
}

func init() { proto.RegisterFile("kta.proto", fileDescriptor_d23c639e4ed5e2ef) }

var fileDescriptor_d23c639e4ed5e2ef = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0xd3, 0x3c,
	0x18, 0xc7, 0xe3, 0xb7, 0x5d, 0xdf, 0xd6, 0xd5, 0xe8, 0x6a, 0x34, 0x94, 0x15, 0x70, 0x2a, 0x73,
	0xe9, 0x65, 0x99, 0x34, 0xd8, 0x85, 0x0b, 0x28, 0x03, 0x04, 0x62, 0x20, 0x64, 0xed, 0x1e, 0xb9,
	0x99, 0x97, 0x45, 0x8b, 0xeb, 0x28, 0x71, 0xd1, 0x7a, 0xe3, 0x23, 0x70, 0xe5, 0x1b, 0xf0, 0x51,
	0x38, 0xf6, 0xb8, 0x53, 0x44, 0xd3, 0x0b, 0x37, 0xa6, 0x7c, 0x02, 0x14, 0x7b, 0x53, 0xb3, 0x55,
	0x48, 0xdc, 0xf2, 0xfc, 0x9f, 0xdf, 0xf3, 0x8b, 0x6d, 0x3d, 0xb0, 0x73, 0xae, 0x98, 0x9b, 0xa4,
	0x52, 0x49, 0xd4, 0x38, 0x57, 0x6c, 0xb0, 0x1b, 0x46, 0xea, 0x6c, 0x3a, 0x76, 0x03, 0x29, 0xf6,
	0x42, 0x19, 0xca, 0x3d, 0xdd, 0x1b, 0x4f, 0x4f, 0x75, 0xa5, 0x0b, 0xfd, 0x65, 0x66, 0x06, 0x0f,
	0x43, 0x29, 0xc3, 0x98, 0xaf, 0x28, 0x2e, 0x12, 0x35, 0x33, 0x4d, 0xf2, 0x02, 0xb6, 0x8e, 0x65,
	0x12, 0x05, 0x19, 0x3a, 0x80, 0x2d, 0xa5, 0xbf, 0xec, 0xc6, 0xb0, 0x31, 0xea, 0xee, 0x43, 0xb7,
	0xfa, 0xad, 0x6e, 0x7a, 0xfd, 0x32, 0x77, 0x36, 0x67, 0x4c, 0xc4, 0xcf, 0x89, 0x61, 0x08, 0xbd,
	0x86, 0xc9, 0xef, 0x26, 0xdc, 0xd0, 0x10, 0x7a, 0x02, 0x9b, 0x13, 0x26, 0xb8, 0x0d, 0x86, 0x60,
	0xd4, 0xf1, 0x7a, 0x65, 0xee, 0x74, 0xcd, 0x48, 0x95, 0x12, 0xaa, 0x9b, 0xe8, 0x00, 0xc2, 0x84,
	0xa5, 0x2a, 0x52, 0x91, 0x9c, 0x64, 0xf6, 0x7f, 0x43, 0x30, 0xda, 0xf0, 0xb6, 0xcb, 0xdc, 0xe9,
	0x1b, 0x74, 0xd5, 0x23, 0xb4, 0x06, 0xa2, 0x23, 0x88, 0x52, 0x9e, 0xc4, 0x51, 0xc0, 0xaa, 0xda,
	0x3f, 0x65, 0x81, 0x92, 0xa9, 0xdd, 0xd0, 0xe3, 0x8f, 0xcb, 0xdc, 0xd9, 0x31, 0xe3, 0xeb, 0x0c,
	0xa1, 0xfd, 0x5a, 0xf8, 0x46, 0x67, 0xe8, 0x1d, 0xec, 0xeb, 0xd3, 0xfb, 0x81, 0x14, 0x49, 0xca,
	0xb3, 0x2c, 0x92, 0x13, 0xbb, 0xa9, 0x8f, 0xfd, 0xa8, 0xcc, 0x1d, 0xbb, 0x76, 0xd3, 0x3a, 0x42,
	0xe8, 0x96, 0xce, 0x0e, 0x57, 0x11, 0x7a, 0x0b, 0xfb, 0x82, 0x5d, 0xf8, 0x82, 0x67, 0x19, 0x0b,
	0xb9, 0x3f, 0x9e, 0x29, 0x9e, 0xd9, 0x1b, 0x77, 0x55, 0x6b, 0x08, 0xa1, 0x3d, 0xc1, 0x2e, 0x3e,
	0x98, 0xc8, 0xab, 0x12, 0xf4, 0x11, 0xde, 0xaf, 0xb0, 0x94, 0x2b, 0x3e, 0xd1, 0x17, 0x30, 0xae,
	0x96, 0x76, 0xe1, 0x32, 0x77, 0x06, 0x2b, 0xd7, 0x1d, 0x88, 0xd0, 0xea, 0x10, 0xf4, 0x26, 0x34,
	0xbe, 0xf7, 0x10, 0xdd, 0x46, 0x55, 0x24, 0xb8, 0xfd, 0xbf, 0xd6, 0xd5, 0x9e, 0x6c, 0x9d, 0x21,
	0x74, 0xab, 0x6e, 0x3b, 0x8e, 0x04, 0x47, 0x2f, 0xe1, 0xbd, 0x20, 0xe6, 0x6c, 0x32, 0x4d, 0xfc,
	0x44, 0xc6, 0x51, 0x30, 0xb3, 0xdb, 0x5a, 0xb4, 0x53, 0xe6, 0xce, 0xb6, 0x11, 0xdd, 0xee, 0x13,
	0xba, 0x79, 0x1d, 0x7c, 0xd2, 0x75, 0xf5, 0x50, 0xd9, 0x99, 0x9c, 0xc6, 0x27, 0xfe, 0x98, 0xfb,
	0x29, 0x17, 0xf2, 0x33, 0x3f, 0xb1, 0x3b, 0x43, 0x30, 0x6a, 0xd7, 0x1f, 0x6a, 0x0d, 0x21, 0xb4,
	0x67, 0x32, 0x8f, 0x53, 0x93, 0xec, 0x7f, 0x03, 0xb0, 0x6d, 0xd6, 0x52, 0x2a, 0xe4, 0xc2, 0xd6,
	0x61, 0xca, 0x99, 0xe2, 0xa8, 0xb6, 0xaf, 0x83, 0x07, 0xae, 0xd9, 0x79, 0xf7, 0x66, 0xe7, 0xdd,
	0xd7, 0xd5, 0xce, 0x13, 0x0b, 0xed, 0xc2, 0xe6, 0x51, 0x94, 0x29, 0xf4, 0x17, 0x62, 0xd0, 0x5d,
	0x59, 0x32, 0x62, 0x55, 0xfa, 0x57, 0x3c, 0xe6, 0xff, 0xaa, 0xf7, 0x9e, 0xcd, 0x17, 0xd8, 0xba,
	0x5c, 0x60, 0xeb, 0x6a, 0x81, 0xc1, 0x97, 0x02, 0x83, 0xef, 0x05, 0x06, 0x3f, 0x0a, 0x0c, 0xe6,
	0x05, 0x06, 0x3f, 0x0b, 0x0c, 0x7e, 0x15, 0xd8, 0xba, 0x2a, 0x30, 0xf8, 0xba, 0xc4, 0xd6, 0x7c,
	0x89, 0xad, 0xcb, 0x25, 0xb6, 0xc6, 0x2d, 0xed, 0x79, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0xc6, 0x6e, 0x90, 0xe9, 0x03, 0x00, 0x00,
}

func (this *Topics) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Topics)
	if !ok {
		that2, ok := that.(Topics)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Topics) != len(that1.Topics) {
		return false
	}
	for i := range this.Topics {
		if !this.Topics[i].Equal(that1.Topics[i]) {
			return false
		}
	}
	return true
}
func (this *Topic) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Topic)
	if !ok {
		that2, ok := that.(Topic)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Partitions != that1.Partitions {
		return false
	}
	if this.ReplicationFactor != that1.ReplicationFactor {
		return false
	}
	if this.TopicCompression != that1.TopicCompression {
		return false
	}
	if this.MaxMessageBytes != that1.MaxMessageBytes {
		return false
	}
	if this.MaxRetentionBytes != that1.MaxRetentionBytes {
		return false
	}
	if this.MaxRetentionTime != that1.MaxRetentionTime {
		return false
	}
	if this.CleanupPolicy != that1.CleanupPolicy {
		return false
	}
	if this.ShouldBeRemoved != that1.ShouldBeRemoved {
		return false
	}
	return true
}
func (this *Topics) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&kta.Topics{")
	if this.Topics != nil {
		s = append(s, "Topics: "+fmt.Sprintf("%#v", this.Topics)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Topic) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 13)
	s = append(s, "&kta.Topic{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Partitions: "+fmt.Sprintf("%#v", this.Partitions)+",\n")
	s = append(s, "ReplicationFactor: "+fmt.Sprintf("%#v", this.ReplicationFactor)+",\n")
	s = append(s, "TopicCompression: "+fmt.Sprintf("%#v", this.TopicCompression)+",\n")
	s = append(s, "MaxMessageBytes: "+fmt.Sprintf("%#v", this.MaxMessageBytes)+",\n")
	s = append(s, "MaxRetentionBytes: "+fmt.Sprintf("%#v", this.MaxRetentionBytes)+",\n")
	s = append(s, "MaxRetentionTime: "+fmt.Sprintf("%#v", this.MaxRetentionTime)+",\n")
	s = append(s, "CleanupPolicy: "+fmt.Sprintf("%#v", this.CleanupPolicy)+",\n")
	s = append(s, "ShouldBeRemoved: "+fmt.Sprintf("%#v", this.ShouldBeRemoved)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringKta(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopicBotClient is the client API for TopicBot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicBotClient interface {
	Create(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Topics, error)
	Delete(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*empty.Empty, error)
}

type topicBotClient struct {
	cc *grpc.ClientConn
}

func NewTopicBotClient(cc *grpc.ClientConn) TopicBotClient {
	return &topicBotClient{cc}
}

func (c *topicBotClient) Create(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kta.TopicBot/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicBotClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Topics, error) {
	out := new(Topics)
	err := c.cc.Invoke(ctx, "/kta.TopicBot/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicBotClient) Delete(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kta.TopicBot/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicBotServer is the server API for TopicBot service.
type TopicBotServer interface {
	Create(context.Context, *Topic) (*empty.Empty, error)
	List(context.Context, *empty.Empty) (*Topics, error)
	Delete(context.Context, *Topic) (*empty.Empty, error)
}

// UnimplementedTopicBotServer can be embedded to have forward compatible implementations.
type UnimplementedTopicBotServer struct {
}

func (*UnimplementedTopicBotServer) Create(ctx context.Context, req *Topic) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTopicBotServer) List(ctx context.Context, req *empty.Empty) (*Topics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedTopicBotServer) Delete(ctx context.Context, req *Topic) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterTopicBotServer(s *grpc.Server, srv TopicBotServer) {
	s.RegisterService(&_TopicBot_serviceDesc, srv)
}

func _TopicBot_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicBotServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kta.TopicBot/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicBotServer).Create(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicBot_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicBotServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kta.TopicBot/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicBotServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicBot_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicBotServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kta.TopicBot/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicBotServer).Delete(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicBot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kta.TopicBot",
	HandlerType: (*TopicBotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TopicBot_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TopicBot_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TopicBot_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kta.proto",
}

func (m *Topics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Topics) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Topics) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Topics) > 0 {
		for iNdEx := len(m.Topics) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Topics[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintKta(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	return len(dAtA) - i, nil
}

func (m *Topic) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Topic) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Topic) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ShouldBeRemoved {
		i--
		if m.ShouldBeRemoved {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.CleanupPolicy) > 0 {
		i -= len(m.CleanupPolicy)
		copy(dAtA[i:], m.CleanupPolicy)
		i = encodeVarintKta(dAtA, i, uint64(len(m.CleanupPolicy)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.MaxRetentionTime) > 0 {
		i -= len(m.MaxRetentionTime)
		copy(dAtA[i:], m.MaxRetentionTime)
		i = encodeVarintKta(dAtA, i, uint64(len(m.MaxRetentionTime)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.MaxRetentionBytes) > 0 {
		i -= len(m.MaxRetentionBytes)
		copy(dAtA[i:], m.MaxRetentionBytes)
		i = encodeVarintKta(dAtA, i, uint64(len(m.MaxRetentionBytes)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MaxMessageBytes) > 0 {
		i -= len(m.MaxMessageBytes)
		copy(dAtA[i:], m.MaxMessageBytes)
		i = encodeVarintKta(dAtA, i, uint64(len(m.MaxMessageBytes)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TopicCompression) > 0 {
		i -= len(m.TopicCompression)
		copy(dAtA[i:], m.TopicCompression)
		i = encodeVarintKta(dAtA, i, uint64(len(m.TopicCompression)))
		i--
		dAtA[i] = 0x22
	}
	if m.ReplicationFactor != 0 {
		i = encodeVarintKta(dAtA, i, uint64(m.ReplicationFactor))
		i--
		dAtA[i] = 0x18
	}
	if m.Partitions != 0 {
		i = encodeVarintKta(dAtA, i, uint64(m.Partitions))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintKta(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKta(dAtA []byte, offset int, v uint64) int {
	offset -= sovKta(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Topics) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Topics) > 0 {
		for _, e := range m.Topics {
			l = e.Size()
			n += 1 + l + sovKta(uint64(l))
		}
	}
	return n
}

func (m *Topic) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovKta(uint64(l))
	}
	if m.Partitions != 0 {
		n += 1 + sovKta(uint64(m.Partitions))
	}
	if m.ReplicationFactor != 0 {
		n += 1 + sovKta(uint64(m.ReplicationFactor))
	}
	l = len(m.TopicCompression)
	if l > 0 {
		n += 1 + l + sovKta(uint64(l))
	}
	l = len(m.MaxMessageBytes)
	if l > 0 {
		n += 1 + l + sovKta(uint64(l))
	}
	l = len(m.MaxRetentionBytes)
	if l > 0 {
		n += 1 + l + sovKta(uint64(l))
	}
	l = len(m.MaxRetentionTime)
	if l > 0 {
		n += 1 + l + sovKta(uint64(l))
	}
	l = len(m.CleanupPolicy)
	if l > 0 {
		n += 1 + l + sovKta(uint64(l))
	}
	if m.ShouldBeRemoved {
		n += 2
	}
	return n
}

func sovKta(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKta(x uint64) (n int) {
	return sovKta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Topics) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForTopics := "[]*Topic{"
	for _, f := range this.Topics {
		repeatedStringForTopics += strings.Replace(f.String(), "Topic", "Topic", 1) + ","
	}
	repeatedStringForTopics += "}"
	s := strings.Join([]string{`&Topics{`,
		`Topics:` + repeatedStringForTopics + `,`,
		`}`,
	}, "")
	return s
}
func (this *Topic) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Topic{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Partitions:` + fmt.Sprintf("%v", this.Partitions) + `,`,
		`ReplicationFactor:` + fmt.Sprintf("%v", this.ReplicationFactor) + `,`,
		`TopicCompression:` + fmt.Sprintf("%v", this.TopicCompression) + `,`,
		`MaxMessageBytes:` + fmt.Sprintf("%v", this.MaxMessageBytes) + `,`,
		`MaxRetentionBytes:` + fmt.Sprintf("%v", this.MaxRetentionBytes) + `,`,
		`MaxRetentionTime:` + fmt.Sprintf("%v", this.MaxRetentionTime) + `,`,
		`CleanupPolicy:` + fmt.Sprintf("%v", this.CleanupPolicy) + `,`,
		`ShouldBeRemoved:` + fmt.Sprintf("%v", this.ShouldBeRemoved) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringKta(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Topics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKta
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Topics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Topics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthKta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topics = append(m.Topics, &Topic{})
			if err := m.Topics[len(m.Topics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKta
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Topic) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKta
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Topic: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Topic: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Partitions", wireType)
			}
			m.Partitions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Partitions |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicationFactor", wireType)
			}
			m.ReplicationFactor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplicationFactor |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopicCompression", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TopicCompression = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMessageBytes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxMessageBytes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRetentionBytes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxRetentionBytes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRetentionTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxRetentionTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CleanupPolicy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CleanupPolicy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShouldBeRemoved", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ShouldBeRemoved = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipKta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKta
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipKta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKta
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowKta
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowKta
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthKta
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthKta
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKta
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipKta(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthKta
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthKta = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKta   = fmt.Errorf("proto: integer overflow")
)