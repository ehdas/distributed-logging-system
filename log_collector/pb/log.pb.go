// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.3
// source: proto/log.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ServiceName   string                 `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Level         string                 `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp     int64                  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	mi := &file_proto_log_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_proto_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *LogRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *LogRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type LogResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogResponse) Reset() {
	*x = LogResponse{}
	mi := &file_proto_log_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponse) ProtoMessage() {}

func (x *LogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponse.ProtoReflect.Descriptor instead.
func (*LogResponse) Descriptor() ([]byte, []int) {
	return file_proto_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_log_proto protoreflect.FileDescriptor

const file_proto_log_proto_rawDesc = "" +
	"\n" +
	"\x0fproto/log.proto\x12\x03log\"}\n" +
	"\n" +
	"LogRequest\x12!\n" +
	"\fservice_name\x18\x01 \x01(\tR\vserviceName\x12\x14\n" +
	"\x05level\x18\x02 \x01(\tR\x05level\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\x12\x1c\n" +
	"\ttimestamp\x18\x04 \x01(\x03R\ttimestamp\"%\n" +
	"\vLogResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2:\n" +
	"\n" +
	"LogService\x12,\n" +
	"\aSendLog\x12\x0f.log.LogRequest\x1a\x10.log.LogResponseB\x05Z\x03/pbb\x06proto3"

var (
	file_proto_log_proto_rawDescOnce sync.Once
	file_proto_log_proto_rawDescData []byte
)

func file_proto_log_proto_rawDescGZIP() []byte {
	file_proto_log_proto_rawDescOnce.Do(func() {
		file_proto_log_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_log_proto_rawDesc), len(file_proto_log_proto_rawDesc)))
	})
	return file_proto_log_proto_rawDescData
}

var file_proto_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_log_proto_goTypes = []any{
	(*LogRequest)(nil),  // 0: log.LogRequest
	(*LogResponse)(nil), // 1: log.LogResponse
}
var file_proto_log_proto_depIdxs = []int32{
	0, // 0: log.LogService.SendLog:input_type -> log.LogRequest
	1, // 1: log.LogService.SendLog:output_type -> log.LogResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_log_proto_init() }
func file_proto_log_proto_init() {
	if File_proto_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_log_proto_rawDesc), len(file_proto_log_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_log_proto_goTypes,
		DependencyIndexes: file_proto_log_proto_depIdxs,
		MessageInfos:      file_proto_log_proto_msgTypes,
	}.Build()
	File_proto_log_proto = out.File
	file_proto_log_proto_goTypes = nil
	file_proto_log_proto_depIdxs = nil
}
