//runserver.proto
//Only user created file in this directory
//Do not modify .pb.go files

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: runserver/runserver.proto

package runserver

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

// Message definition
type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Low          uint64 `protobuf:"varint,3,opt,name=low,proto3" json:"low,omitempty"`
	Mid          uint64 `protobuf:"varint,4,opt,name=mid,proto3" json:"mid,omitempty"`
	High         uint64 `protobuf:"varint,5,opt,name=high,proto3" json:"high,omitempty"`
	PartialValue uint64 `protobuf:"varint,6,opt,name=partial_value,json=partialValue,proto3" json:"partial_value,omitempty"`
	FinalValue   uint64 `protobuf:"varint,7,opt,name=final_value,json=finalValue,proto3" json:"final_value,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runserver_runserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_runserver_runserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_runserver_runserver_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Token) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Token) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Token) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *Token) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Token) GetPartialValue() uint64 {
	if x != nil {
		return x.PartialValue
	}
	return 0
}

func (x *Token) GetFinalValue() uint64 {
	if x != nil {
		return x.FinalValue
	}
	return 0
}

var File_runserver_runserver_proto protoreflect.FileDescriptor

var file_runserver_runserver_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x75, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x75, 0x6e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xa9, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x32, 0xbf, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x72, 0x75,
	0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10, 0x2e,
	0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x2a, 0x0a, 0x04, 0x64, 0x72, 0x6f, 0x70, 0x12, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2b, 0x0a, 0x05, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64,
	0x12, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x15, 0x5a, 0x13, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_runserver_runserver_proto_rawDescOnce sync.Once
	file_runserver_runserver_proto_rawDescData = file_runserver_runserver_proto_rawDesc
)

func file_runserver_runserver_proto_rawDescGZIP() []byte {
	file_runserver_runserver_proto_rawDescOnce.Do(func() {
		file_runserver_runserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_runserver_runserver_proto_rawDescData)
	})
	return file_runserver_runserver_proto_rawDescData
}

var file_runserver_runserver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_runserver_runserver_proto_goTypes = []interface{}{
	(*Token)(nil), // 0: runserver.Token
}
var file_runserver_runserver_proto_depIdxs = []int32{
	0, // 0: runserver.RunService.create:input_type -> runserver.Token
	0, // 1: runserver.RunService.drop:input_type -> runserver.Token
	0, // 2: runserver.RunService.write:input_type -> runserver.Token
	0, // 3: runserver.RunService.read:input_type -> runserver.Token
	0, // 4: runserver.RunService.create:output_type -> runserver.Token
	0, // 5: runserver.RunService.drop:output_type -> runserver.Token
	0, // 6: runserver.RunService.write:output_type -> runserver.Token
	0, // 7: runserver.RunService.read:output_type -> runserver.Token
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_runserver_runserver_proto_init() }
func file_runserver_runserver_proto_init() {
	if File_runserver_runserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runserver_runserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_runserver_runserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runserver_runserver_proto_goTypes,
		DependencyIndexes: file_runserver_runserver_proto_depIdxs,
		MessageInfos:      file_runserver_runserver_proto_msgTypes,
	}.Build()
	File_runserver_runserver_proto = out.File
	file_runserver_runserver_proto_rawDesc = nil
	file_runserver_runserver_proto_goTypes = nil
	file_runserver_runserver_proto_depIdxs = nil
}
