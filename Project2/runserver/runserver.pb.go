//runserver.proto
//Only user created file in this directory
//Do not modify .pb.go files

//Need to make changes here if:
// Token definition changes
// More rpc functions are created

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

	ID            string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	NAME          string            `protobuf:"bytes,2,opt,name=NAME,proto3" json:"NAME,omitempty"`
	LOW           uint64            `protobuf:"varint,3,opt,name=LOW,proto3" json:"LOW,omitempty"`
	MID           uint64            `protobuf:"varint,4,opt,name=MID,proto3" json:"MID,omitempty"`
	HIGH          uint64            `protobuf:"varint,5,opt,name=HIGH,proto3" json:"HIGH,omitempty"`
	PARTIAL_VALUE uint64            `protobuf:"varint,6,opt,name=PARTIAL_VALUE,json=PARTIALVALUE,proto3" json:"PARTIAL_VALUE,omitempty"`
	FINAL_VALUE   uint64            `protobuf:"varint,7,opt,name=FINAL_VALUE,json=FINALVALUE,proto3" json:"FINAL_VALUE,omitempty"`
	WRITER        map[string]string `protobuf:"bytes,8,rep,name=WRITER,proto3" json:"WRITER,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	READER        map[string]string `protobuf:"bytes,9,rep,name=READER,proto3" json:"READER,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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

func (x *Token) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Token) GetNAME() string {
	if x != nil {
		return x.NAME
	}
	return ""
}

func (x *Token) GetLOW() uint64 {
	if x != nil {
		return x.LOW
	}
	return 0
}

func (x *Token) GetMID() uint64 {
	if x != nil {
		return x.MID
	}
	return 0
}

func (x *Token) GetHIGH() uint64 {
	if x != nil {
		return x.HIGH
	}
	return 0
}

func (x *Token) GetPARTIAL_VALUE() uint64 {
	if x != nil {
		return x.PARTIAL_VALUE
	}
	return 0
}

func (x *Token) GetFINAL_VALUE() uint64 {
	if x != nil {
		return x.FINAL_VALUE
	}
	return 0
}

func (x *Token) GetWRITER() map[string]string {
	if x != nil {
		return x.WRITER
	}
	return nil
}

func (x *Token) GetREADER() map[string]string {
	if x != nil {
		return x.READER
	}
	return nil
}

var File_runserver_runserver_proto protoreflect.FileDescriptor

var file_runserver_runserver_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x75, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x75, 0x6e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x8b, 0x03, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x41, 0x4d, 0x45, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x4c, 0x4f, 0x57, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x4d, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x48, 0x49, 0x47, 0x48, 0x12, 0x23, 0x0a, 0x0d,
	0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0c, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x12, 0x1f, 0x0a, 0x0b, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x12, 0x34, 0x0a, 0x06, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52, 0x12, 0x34, 0x0a, 0x06, 0x52, 0x45, 0x41, 0x44,
	0x45, 0x52, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x52, 0x45, 0x41, 0x44, 0x45,
	0x52, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x52, 0x45, 0x41, 0x44, 0x45, 0x52, 0x1a, 0x39,
	0x0a, 0x0b, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x52, 0x45, 0x41,
	0x44, 0x45, 0x52, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0xeb, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e,
	0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a,
	0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x44, 0x72, 0x6f, 0x70, 0x12, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x72, 0x75,
	0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2b, 0x0a,
	0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x1a, 0x10, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x15, 0x5a, 0x13, 0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x72, 0x75, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_runserver_runserver_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_runserver_runserver_proto_goTypes = []interface{}{
	(*Token)(nil), // 0: runserver.Token
	nil,           // 1: runserver.Token.WRITEREntry
	nil,           // 2: runserver.Token.READEREntry
}
var file_runserver_runserver_proto_depIdxs = []int32{
	1, // 0: runserver.Token.WRITER:type_name -> runserver.Token.WRITEREntry
	2, // 1: runserver.Token.READER:type_name -> runserver.Token.READEREntry
	0, // 2: runserver.RunService.Create:input_type -> runserver.Token
	0, // 3: runserver.RunService.Drop:input_type -> runserver.Token
	0, // 4: runserver.RunService.Write:input_type -> runserver.Token
	0, // 5: runserver.RunService.Read:input_type -> runserver.Token
	0, // 6: runserver.RunService.Test:input_type -> runserver.Token
	0, // 7: runserver.RunService.Create:output_type -> runserver.Token
	0, // 8: runserver.RunService.Drop:output_type -> runserver.Token
	0, // 9: runserver.RunService.Write:output_type -> runserver.Token
	0, // 10: runserver.RunService.Read:output_type -> runserver.Token
	0, // 11: runserver.RunService.Test:output_type -> runserver.Token
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
			NumMessages:   3,
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
