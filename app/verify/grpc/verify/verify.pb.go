// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: verify.proto

package verify

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListInformRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerifyInfo string `protobuf:"bytes,1,opt,name=verifyInfo,proto3" json:"verifyInfo,omitempty"`
	PageIndex  int64  `protobuf:"varint,2,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize   int64  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	VerifyType string `protobuf:"bytes,4,opt,name=verifyType,proto3" json:"verifyType,omitempty"`
	SocialName string `protobuf:"bytes,5,opt,name=socialName,proto3" json:"socialName,omitempty"`
	StartTime  string `protobuf:"bytes,6,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime    string `protobuf:"bytes,7,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *ListInformRequest) Reset() {
	*x = ListInformRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInformRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInformRequest) ProtoMessage() {}

func (x *ListInformRequest) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInformRequest.ProtoReflect.Descriptor instead.
func (*ListInformRequest) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{0}
}

func (x *ListInformRequest) GetVerifyInfo() string {
	if x != nil {
		return x.VerifyInfo
	}
	return ""
}

func (x *ListInformRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *ListInformRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListInformRequest) GetVerifyType() string {
	if x != nil {
		return x.VerifyType
	}
	return ""
}

func (x *ListInformRequest) GetSocialName() string {
	if x != nil {
		return x.SocialName
	}
	return ""
}

func (x *ListInformRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ListInformRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type ListInform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerifyInfo string `protobuf:"bytes,1,opt,name=verifyInfo,proto3" json:"verifyInfo,omitempty"`
	VerifyType string `protobuf:"bytes,2,opt,name=verifyType,proto3" json:"verifyType,omitempty"`
	SocialName string `protobuf:"bytes,3,opt,name=socialName,proto3" json:"socialName,omitempty"`
	CreateTime string `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *ListInform) Reset() {
	*x = ListInform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInform) ProtoMessage() {}

func (x *ListInform) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInform.ProtoReflect.Descriptor instead.
func (*ListInform) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{1}
}

func (x *ListInform) GetVerifyInfo() string {
	if x != nil {
		return x.VerifyInfo
	}
	return ""
}

func (x *ListInform) GetVerifyType() string {
	if x != nil {
		return x.VerifyType
	}
	return ""
}

func (x *ListInform) GetSocialName() string {
	if x != nil {
		return x.SocialName
	}
	return ""
}

func (x *ListInform) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

type ListInformResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*ListInform `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Count int32         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListInformResponse) Reset() {
	*x = ListInformResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInformResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInformResponse) ProtoMessage() {}

func (x *ListInformResponse) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInformResponse.ProtoReflect.Descriptor instead.
func (*ListInformResponse) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{2}
}

func (x *ListInformResponse) GetList() []*ListInform {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListInformResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_verify_proto protoreflect.FileDescriptor

var file_verify_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x22, 0xe5, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x8c,
	0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1e, 0x0a,
	0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a,
	0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x52, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0x4d, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x43, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x19, 0x2e, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_verify_proto_rawDescOnce sync.Once
	file_verify_proto_rawDescData = file_verify_proto_rawDesc
)

func file_verify_proto_rawDescGZIP() []byte {
	file_verify_proto_rawDescOnce.Do(func() {
		file_verify_proto_rawDescData = protoimpl.X.CompressGZIP(file_verify_proto_rawDescData)
	})
	return file_verify_proto_rawDescData
}

var file_verify_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_verify_proto_goTypes = []interface{}{
	(*ListInformRequest)(nil),  // 0: verify.ListInformRequest
	(*ListInform)(nil),         // 1: verify.ListInform
	(*ListInformResponse)(nil), // 2: verify.ListInformResponse
}
var file_verify_proto_depIdxs = []int32{
	1, // 0: verify.ListInformResponse.list:type_name -> verify.ListInform
	0, // 1: verify.Verify.ListInform:input_type -> verify.ListInformRequest
	2, // 2: verify.Verify.ListInform:output_type -> verify.ListInformResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_verify_proto_init() }
func file_verify_proto_init() {
	if File_verify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_verify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInformRequest); i {
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
		file_verify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInform); i {
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
		file_verify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInformResponse); i {
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
			RawDescriptor: file_verify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_verify_proto_goTypes,
		DependencyIndexes: file_verify_proto_depIdxs,
		MessageInfos:      file_verify_proto_msgTypes,
	}.Build()
	File_verify_proto = out.File
	file_verify_proto_rawDesc = nil
	file_verify_proto_goTypes = nil
	file_verify_proto_depIdxs = nil
}