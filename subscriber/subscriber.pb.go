// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: subscriber/subscriber.proto

package subscriber

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_subscriber_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_subscriber_proto_msgTypes[0]
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
	return file_subscriber_subscriber_proto_rawDescGZIP(), []int{0}
}

func (x *Topic) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type REvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic   string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Message []byte `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *REvent) Reset() {
	*x = REvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriber_subscriber_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *REvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*REvent) ProtoMessage() {}

func (x *REvent) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_subscriber_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use REvent.ProtoReflect.Descriptor instead.
func (*REvent) Descriptor() ([]byte, []int) {
	return file_subscriber_subscriber_proto_rawDescGZIP(), []int{1}
}

func (x *REvent) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *REvent) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_subscriber_subscriber_proto protoreflect.FileDescriptor

var file_subscriber_subscriber_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a,
	0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x38, 0x0a, 0x06,
	0x52, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x2e, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x06, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x1a, 0x07, 0x2e, 0x52, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x72, 0x65, 0x65, 0x72, 0x61, 0x6d, 0x37, 0x37, 0x2f, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscriber_subscriber_proto_rawDescOnce sync.Once
	file_subscriber_subscriber_proto_rawDescData = file_subscriber_subscriber_proto_rawDesc
)

func file_subscriber_subscriber_proto_rawDescGZIP() []byte {
	file_subscriber_subscriber_proto_rawDescOnce.Do(func() {
		file_subscriber_subscriber_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscriber_subscriber_proto_rawDescData)
	})
	return file_subscriber_subscriber_proto_rawDescData
}

var file_subscriber_subscriber_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_subscriber_subscriber_proto_goTypes = []interface{}{
	(*Topic)(nil),  // 0: Topic
	(*REvent)(nil), // 1: REvent
}
var file_subscriber_subscriber_proto_depIdxs = []int32{
	0, // 0: Subscriber.Subscribe:input_type -> Topic
	1, // 1: Subscriber.Subscribe:output_type -> REvent
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_subscriber_subscriber_proto_init() }
func file_subscriber_subscriber_proto_init() {
	if File_subscriber_subscriber_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subscriber_subscriber_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_subscriber_subscriber_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*REvent); i {
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
			RawDescriptor: file_subscriber_subscriber_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subscriber_subscriber_proto_goTypes,
		DependencyIndexes: file_subscriber_subscriber_proto_depIdxs,
		MessageInfos:      file_subscriber_subscriber_proto_msgTypes,
	}.Build()
	File_subscriber_subscriber_proto = out.File
	file_subscriber_subscriber_proto_rawDesc = nil
	file_subscriber_subscriber_proto_goTypes = nil
	file_subscriber_subscriber_proto_depIdxs = nil
}
