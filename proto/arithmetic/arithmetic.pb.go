// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: arithmetic/arithmetic.proto

package arithmetic

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message containing the two integers
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arithmetic_arithmetic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_arithmetic_arithmetic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_arithmetic_arithmetic_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Request) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

// The response message containing the result
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arithmetic_arithmetic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_arithmetic_arithmetic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_arithmetic_arithmetic_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_arithmetic_arithmetic_proto protoreflect.FileDescriptor

var file_arithmetic_arithmetic_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2f, 0x61, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x61,
	0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x62, 0x22, 0x22,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0xdd, 0x02, 0x0a, 0x0a, 0x41, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69,
	0x63, 0x12, 0x4f, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e, 0x61, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2f, 0x61,
	0x64, 0x64, 0x12, 0x52, 0x0a, 0x06, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x12, 0x13, 0x2e, 0x61,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a,
	0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74,
	0x69, 0x63, 0x2f, 0x64, 0x69, 0x76, 0x12, 0x54, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x79, 0x12, 0x13, 0x2e, 0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x65, 0x74, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2f, 0x6d, 0x75, 0x6c, 0x12, 0x54, 0x0a, 0x08,
	0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x2f, 0x73,
	0x75, 0x62, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x61, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74,
	0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_arithmetic_arithmetic_proto_rawDescOnce sync.Once
	file_arithmetic_arithmetic_proto_rawDescData = file_arithmetic_arithmetic_proto_rawDesc
)

func file_arithmetic_arithmetic_proto_rawDescGZIP() []byte {
	file_arithmetic_arithmetic_proto_rawDescOnce.Do(func() {
		file_arithmetic_arithmetic_proto_rawDescData = protoimpl.X.CompressGZIP(file_arithmetic_arithmetic_proto_rawDescData)
	})
	return file_arithmetic_arithmetic_proto_rawDescData
}

var file_arithmetic_arithmetic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_arithmetic_arithmetic_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: arithmetic.Request
	(*Response)(nil), // 1: arithmetic.Response
}
var file_arithmetic_arithmetic_proto_depIdxs = []int32{
	0, // 0: arithmetic.Arithmetic.Add:input_type -> arithmetic.Request
	0, // 1: arithmetic.Arithmetic.Divide:input_type -> arithmetic.Request
	0, // 2: arithmetic.Arithmetic.Multiply:input_type -> arithmetic.Request
	0, // 3: arithmetic.Arithmetic.Subtract:input_type -> arithmetic.Request
	1, // 4: arithmetic.Arithmetic.Add:output_type -> arithmetic.Response
	1, // 5: arithmetic.Arithmetic.Divide:output_type -> arithmetic.Response
	1, // 6: arithmetic.Arithmetic.Multiply:output_type -> arithmetic.Response
	1, // 7: arithmetic.Arithmetic.Subtract:output_type -> arithmetic.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_arithmetic_arithmetic_proto_init() }
func file_arithmetic_arithmetic_proto_init() {
	if File_arithmetic_arithmetic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_arithmetic_arithmetic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_arithmetic_arithmetic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_arithmetic_arithmetic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_arithmetic_arithmetic_proto_goTypes,
		DependencyIndexes: file_arithmetic_arithmetic_proto_depIdxs,
		MessageInfos:      file_arithmetic_arithmetic_proto_msgTypes,
	}.Build()
	File_arithmetic_arithmetic_proto = out.File
	file_arithmetic_arithmetic_proto_rawDesc = nil
	file_arithmetic_arithmetic_proto_goTypes = nil
	file_arithmetic_arithmetic_proto_depIdxs = nil
}
