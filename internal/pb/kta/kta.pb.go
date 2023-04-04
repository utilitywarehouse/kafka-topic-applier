// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/kta.proto

package kta

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Topics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []*Topic `protobuf:"bytes,3,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *Topics) Reset() {
	*x = Topics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topics) ProtoMessage() {}

func (x *Topics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topics.ProtoReflect.Descriptor instead.
func (*Topics) Descriptor() ([]byte, []int) {
	return file_proto_kta_proto_rawDescGZIP(), []int{0}
}

func (x *Topics) GetTopics() []*Topic {
	if x != nil {
		return x.Topics
	}
	return nil
}

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Partitions        int32  `protobuf:"varint,2,opt,name=partitions,proto3" json:"partitions,omitempty"`
	ReplicationFactor int32  `protobuf:"varint,3,opt,name=replication_factor,json=replicationFactor,proto3" json:"replication_factor,omitempty"`
	TopicCompression  string `protobuf:"bytes,4,opt,name=topic_compression,json=topicCompression,proto3" json:"topic_compression,omitempty"`
	MaxMessageBytes   string `protobuf:"bytes,5,opt,name=max_message_bytes,json=maxMessageBytes,proto3" json:"max_message_bytes,omitempty"`
	MaxRetentionBytes string `protobuf:"bytes,6,opt,name=max_retention_bytes,json=maxRetentionBytes,proto3" json:"max_retention_bytes,omitempty"`
	MaxRetentionTime  string `protobuf:"bytes,7,opt,name=max_retention_time,json=maxRetentionTime,proto3" json:"max_retention_time,omitempty"`
	CleanupPolicy     string `protobuf:"bytes,8,opt,name=cleanup_policy,json=cleanupPolicy,proto3" json:"cleanup_policy,omitempty"`
	ShouldBeRemoved   bool   `protobuf:"varint,9,opt,name=should_be_removed,json=shouldBeRemoved,proto3" json:"should_be_removed,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_proto_kta_proto_rawDescGZIP(), []int{1}
}

func (x *Topic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Topic) GetPartitions() int32 {
	if x != nil {
		return x.Partitions
	}
	return 0
}

func (x *Topic) GetReplicationFactor() int32 {
	if x != nil {
		return x.ReplicationFactor
	}
	return 0
}

func (x *Topic) GetTopicCompression() string {
	if x != nil {
		return x.TopicCompression
	}
	return ""
}

func (x *Topic) GetMaxMessageBytes() string {
	if x != nil {
		return x.MaxMessageBytes
	}
	return ""
}

func (x *Topic) GetMaxRetentionBytes() string {
	if x != nil {
		return x.MaxRetentionBytes
	}
	return ""
}

func (x *Topic) GetMaxRetentionTime() string {
	if x != nil {
		return x.MaxRetentionTime
	}
	return ""
}

func (x *Topic) GetCleanupPolicy() string {
	if x != nil {
		return x.CleanupPolicy
	}
	return ""
}

func (x *Topic) GetShouldBeRemoved() bool {
	if x != nil {
		return x.ShouldBeRemoved
	}
	return false
}

var File_proto_kta_proto protoreflect.FileDescriptor

var file_proto_kta_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x6b, 0x74, 0x61, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x06, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x22, 0x0a,
	0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x6b, 0x74, 0x61, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x22, 0xf4, 0x02, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2d, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2b,
	0x0a, 0x11, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x6d,
	0x61, 0x78, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x72,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x72,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2a, 0x0a, 0x11,
	0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x42,
	0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x32, 0x99, 0x01, 0x0a, 0x08, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x42, 0x6f, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x0a, 0x2e, 0x6b, 0x74, 0x61, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x6b, 0x74, 0x61, 0x2e, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0a,
	0x2e, 0x6b, 0x74, 0x61, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x6b, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_kta_proto_rawDescOnce sync.Once
	file_proto_kta_proto_rawDescData = file_proto_kta_proto_rawDesc
)

func file_proto_kta_proto_rawDescGZIP() []byte {
	file_proto_kta_proto_rawDescOnce.Do(func() {
		file_proto_kta_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_kta_proto_rawDescData)
	})
	return file_proto_kta_proto_rawDescData
}

var file_proto_kta_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_kta_proto_goTypes = []interface{}{
	(*Topics)(nil),        // 0: kta.Topics
	(*Topic)(nil),         // 1: kta.Topic
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_proto_kta_proto_depIdxs = []int32{
	1, // 0: kta.Topics.topics:type_name -> kta.Topic
	1, // 1: kta.TopicBot.Create:input_type -> kta.Topic
	2, // 2: kta.TopicBot.List:input_type -> google.protobuf.Empty
	1, // 3: kta.TopicBot.Delete:input_type -> kta.Topic
	2, // 4: kta.TopicBot.Create:output_type -> google.protobuf.Empty
	0, // 5: kta.TopicBot.List:output_type -> kta.Topics
	2, // 6: kta.TopicBot.Delete:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_kta_proto_init() }
func file_proto_kta_proto_init() {
	if File_proto_kta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_kta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_kta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_kta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kta_proto_goTypes,
		DependencyIndexes: file_proto_kta_proto_depIdxs,
		MessageInfos:      file_proto_kta_proto_msgTypes,
	}.Build()
	File_proto_kta_proto = out.File
	file_proto_kta_proto_rawDesc = nil
	file_proto_kta_proto_goTypes = nil
	file_proto_kta_proto_depIdxs = nil
}
