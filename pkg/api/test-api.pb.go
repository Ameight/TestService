// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: test-api.proto

package api

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

type GetInfoResponseTest int32

const (
	GetInfoResponse_int64 GetInfoResponseTest = 0
	GetInfoResponse_int32 GetInfoResponseTest = 1
	GetInfoResponse_int16 GetInfoResponseTest = 2
)

// Enum value maps for GetInfoResponseTest.
var (
	GetInfoResponseTest_name = map[int32]string{
		0: "int64",
		1: "int32",
		2: "int16",
	}
	GetInfoResponseTest_value = map[string]int32{
		"int64": 0,
		"int32": 1,
		"int16": 2,
	}
)

func (x GetInfoResponseTest) Enum() *GetInfoResponseTest {
	p := new(GetInfoResponseTest)
	*p = x
	return p
}

func (x GetInfoResponseTest) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetInfoResponseTest) Descriptor() protoreflect.EnumDescriptor {
	return file_test_api_proto_enumTypes[0].Descriptor()
}

func (GetInfoResponseTest) Type() protoreflect.EnumType {
	return &file_test_api_proto_enumTypes[0]
}

func (x GetInfoResponseTest) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetInfoResponseTest.Descriptor instead.
func (GetInfoResponseTest) EnumDescriptor() ([]byte, []int) {
	return file_test_api_proto_rawDescGZIP(), []int{3, 0}
}

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str  string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_test_api_proto_rawDescGZIP(), []int{0}
}

func (x *TestRequest) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *TestRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_test_api_proto_rawDescGZIP(), []int{1}
}

func (x *TestResponse) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_test_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetInfoRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number1 int64               `protobuf:"varint,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2 GetInfoResponseTest `protobuf:"varint,2,opt,name=number2,proto3,enum=api.GetInfoResponseTest" json:"number2,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_test_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetInfoResponse) GetNumber1() int64 {
	if x != nil {
		return x.Number1
	}
	return 0
}

func (x *GetInfoResponse) GetNumber2() GetInfoResponseTest {
	if x != nil {
		return x.Number2
	}
	return GetInfoResponse_int64
}

var File_test_api_proto protoreflect.FileDescriptor

var file_test_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x33, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0c, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x22, 0x28, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x89, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x31, 0x12, 0x33, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x22, 0x27, 0x0a, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x12, 0x09, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x31, 0x36,
	0x10, 0x02, 0x32, 0x71, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x74, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_api_proto_rawDescOnce sync.Once
	file_test_api_proto_rawDescData = file_test_api_proto_rawDesc
)

func file_test_api_proto_rawDescGZIP() []byte {
	file_test_api_proto_rawDescOnce.Do(func() {
		file_test_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_api_proto_rawDescData)
	})
	return file_test_api_proto_rawDescData
}

var file_test_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_test_api_proto_goTypes = []interface{}{
	(GetInfoResponseTest)(0), // 0: api.GetInfoResponse.test
	(*TestRequest)(nil),      // 1: api.TestRequest
	(*TestResponse)(nil),     // 2: api.TestResponse
	(*GetInfoRequest)(nil),   // 3: api.GetInfoRequest
	(*GetInfoResponse)(nil),  // 4: api.GetInfoResponse
}
var file_test_api_proto_depIdxs = []int32{
	0, // 0: api.GetInfoResponse.number2:type_name -> api.GetInfoResponse.test
	1, // 1: api.Test.testMethod:input_type -> api.TestRequest
	3, // 2: api.Test.getInfo:input_type -> api.GetInfoRequest
	2, // 3: api.Test.testMethod:output_type -> api.TestResponse
	4, // 4: api.Test.getInfo:output_type -> api.GetInfoResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_test_api_proto_init() }
func file_test_api_proto_init() {
	if File_test_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequest); i {
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
		file_test_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResponse); i {
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
		file_test_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoRequest); i {
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
		file_test_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponse); i {
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
			RawDescriptor: file_test_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_api_proto_goTypes,
		DependencyIndexes: file_test_api_proto_depIdxs,
		EnumInfos:         file_test_api_proto_enumTypes,
		MessageInfos:      file_test_api_proto_msgTypes,
	}.Build()
	File_test_api_proto = out.File
	file_test_api_proto_rawDesc = nil
	file_test_api_proto_goTypes = nil
	file_test_api_proto_depIdxs = nil
}
