// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: protos/addressbook.proto

package clientGRPCJsonXml

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Format  *string `protobuf:"bytes,1,req,name=format" json:"format,omitempty"`
	Payload *string `protobuf:"bytes,2,req,name=payload" json:"payload,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_addressbook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_protos_addressbook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_protos_addressbook_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetFormat() string {
	if x != nil && x.Format != nil {
		return *x.Format
	}
	return ""
}

func (x *Message) GetPayload() string {
	if x != nil && x.Payload != nil {
		return *x.Payload
	}
	return ""
}

var File_protos_addressbook_proto protoreflect.FileDescriptor

var file_protos_addressbook_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x3b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x32, 0x9e, 0x02, 0x0a, 0x0a, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x2e,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0a, 0x45,
	0x64, 0x69, 0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x14, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x47, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x14, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01,
}

var (
	file_protos_addressbook_proto_rawDescOnce sync.Once
	file_protos_addressbook_proto_rawDescData = file_protos_addressbook_proto_rawDesc
)

func file_protos_addressbook_proto_rawDescGZIP() []byte {
	file_protos_addressbook_proto_rawDescOnce.Do(func() {
		file_protos_addressbook_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_addressbook_proto_rawDescData)
	})
	return file_protos_addressbook_proto_rawDescData
}

var file_protos_addressbook_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_addressbook_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: addressbook.Message
}
var file_protos_addressbook_proto_depIdxs = []int32{
	0, // 0: addressbook.RPCService.GetPersonByPhoneNumber:input_type -> addressbook.Message
	0, // 1: addressbook.RPCService.EditPeople:input_type -> addressbook.Message
	0, // 2: addressbook.RPCService.ListPeopleByPhoneType:input_type -> addressbook.Message
	0, // 3: addressbook.RPCService.GetPeopleById:input_type -> addressbook.Message
	0, // 4: addressbook.RPCService.GetPersonByPhoneNumber:output_type -> addressbook.Message
	0, // 5: addressbook.RPCService.EditPeople:output_type -> addressbook.Message
	0, // 6: addressbook.RPCService.ListPeopleByPhoneType:output_type -> addressbook.Message
	0, // 7: addressbook.RPCService.GetPeopleById:output_type -> addressbook.Message
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_addressbook_proto_init() }
func file_protos_addressbook_proto_init() {
	if File_protos_addressbook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_addressbook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_protos_addressbook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_addressbook_proto_goTypes,
		DependencyIndexes: file_protos_addressbook_proto_depIdxs,
		MessageInfos:      file_protos_addressbook_proto_msgTypes,
	}.Build()
	File_protos_addressbook_proto = out.File
	file_protos_addressbook_proto_rawDesc = nil
	file_protos_addressbook_proto_goTypes = nil
	file_protos_addressbook_proto_depIdxs = nil
}