// proto/sharing/sharing.proto

// Версия ProtoBuf

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: sharing/sharing.proto

// Текущий пакет - указывает пространство имен для сервиса и сообщений. Помогает избегать конфликтов имен.

package sharingv1

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

// Объект, который отправляется при вызове RPC-метода (ручки) Register.
type PullerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"` // IP of the agent.
}

func (x *PullerRequest) Reset() {
	*x = PullerRequest{}
	mi := &file_sharing_sharing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullerRequest) ProtoMessage() {}

func (x *PullerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sharing_sharing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullerRequest.ProtoReflect.Descriptor instead.
func (*PullerRequest) Descriptor() ([]byte, []int) {
	return file_sharing_sharing_proto_rawDescGZIP(), []int{0}
}

func (x *PullerRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

// Объект, который метод (ручка) вернёт.
type PullerResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImagePath string `protobuf:"bytes,1,opt,name=imagePath,proto3" json:"imagePath,omitempty"` // User ID of the registered user.
	Param     string `protobuf:"bytes,2,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *PullerResponce) Reset() {
	*x = PullerResponce{}
	mi := &file_sharing_sharing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullerResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullerResponce) ProtoMessage() {}

func (x *PullerResponce) ProtoReflect() protoreflect.Message {
	mi := &file_sharing_sharing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullerResponce.ProtoReflect.Descriptor instead.
func (*PullerResponce) Descriptor() ([]byte, []int) {
	return file_sharing_sharing_proto_rawDescGZIP(), []int{1}
}

func (x *PullerResponce) GetImagePath() string {
	if x != nil {
		return x.ImagePath
	}
	return ""
}

func (x *PullerResponce) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

// То же самое для метода Login()
type ResulterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`               // Email of the user to login.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`         // Password of the user to login.
	AppId    int32  `protobuf:"varint,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"` // ID of the app to login to.
}

func (x *ResulterRequest) Reset() {
	*x = ResulterRequest{}
	mi := &file_sharing_sharing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResulterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResulterRequest) ProtoMessage() {}

func (x *ResulterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sharing_sharing_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResulterRequest.ProtoReflect.Descriptor instead.
func (*ResulterRequest) Descriptor() ([]byte, []int) {
	return file_sharing_sharing_proto_rawDescGZIP(), []int{2}
}

func (x *ResulterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResulterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ResulterRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type ResulterResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // Auth token of the logged in user.
}

func (x *ResulterResponce) Reset() {
	*x = ResulterResponce{}
	mi := &file_sharing_sharing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResulterResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResulterResponce) ProtoMessage() {}

func (x *ResulterResponce) ProtoReflect() protoreflect.Message {
	mi := &file_sharing_sharing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResulterResponce.ProtoReflect.Descriptor instead.
func (*ResulterResponce) Descriptor() ([]byte, []int) {
	return file_sharing_sharing_proto_rawDescGZIP(), []int{3}
}

func (x *ResulterResponce) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_sharing_sharing_proto protoreflect.FileDescriptor

var file_sharing_sharing_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67,
	0x22, 0x1f, 0x0a, 0x0d, 0x50, 0x75, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x22, 0x44, 0x0a, 0x0e, 0x50, 0x75, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x5a, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x85, 0x01,
	0x0a, 0x07, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x06, 0x50, 0x75, 0x6c,
	0x6c, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x75,
	0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x18, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x61, 0x72, 0x69, 0x61, 0x65, 0x76,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x61, 0x72,
	0x69, 0x6e, 0x67, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sharing_sharing_proto_rawDescOnce sync.Once
	file_sharing_sharing_proto_rawDescData = file_sharing_sharing_proto_rawDesc
)

func file_sharing_sharing_proto_rawDescGZIP() []byte {
	file_sharing_sharing_proto_rawDescOnce.Do(func() {
		file_sharing_sharing_proto_rawDescData = protoimpl.X.CompressGZIP(file_sharing_sharing_proto_rawDescData)
	})
	return file_sharing_sharing_proto_rawDescData
}

var file_sharing_sharing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sharing_sharing_proto_goTypes = []any{
	(*PullerRequest)(nil),    // 0: sharing.PullerRequest
	(*PullerResponce)(nil),   // 1: sharing.PullerResponce
	(*ResulterRequest)(nil),  // 2: sharing.ResulterRequest
	(*ResulterResponce)(nil), // 3: sharing.ResulterResponce
}
var file_sharing_sharing_proto_depIdxs = []int32{
	0, // 0: sharing.Sharing.Puller:input_type -> sharing.PullerRequest
	2, // 1: sharing.Sharing.Resulter:input_type -> sharing.ResulterRequest
	1, // 2: sharing.Sharing.Puller:output_type -> sharing.PullerResponce
	3, // 3: sharing.Sharing.Resulter:output_type -> sharing.ResulterResponce
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sharing_sharing_proto_init() }
func file_sharing_sharing_proto_init() {
	if File_sharing_sharing_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sharing_sharing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sharing_sharing_proto_goTypes,
		DependencyIndexes: file_sharing_sharing_proto_depIdxs,
		MessageInfos:      file_sharing_sharing_proto_msgTypes,
	}.Build()
	File_sharing_sharing_proto = out.File
	file_sharing_sharing_proto_rawDesc = nil
	file_sharing_sharing_proto_goTypes = nil
	file_sharing_sharing_proto_depIdxs = nil
}