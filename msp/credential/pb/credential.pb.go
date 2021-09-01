// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: credential.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	pb "github.com/erda-project/erda-proto-go/core/services/authentication/credentials/accesskey/pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAccessKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectType pb.SubjectTypeEnum_SubjectType `protobuf:"varint,1,opt,name=subjectType,proto3,enum=erda.core.services.authentication.credentials.accesskey.SubjectTypeEnum_SubjectType" json:"subjectType,omitempty"`
	Subject     string                         `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Description string                         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateAccessKeyRequest) Reset() {
	*x = CreateAccessKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccessKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccessKeyRequest) ProtoMessage() {}

func (x *CreateAccessKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccessKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateAccessKeyRequest) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccessKeyRequest) GetSubjectType() pb.SubjectTypeEnum_SubjectType {
	if x != nil {
		return x.SubjectType
	}
	return pb.SubjectTypeEnum_NOT_SPECIFIED
}

func (x *CreateAccessKeyRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *CreateAccessKeyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateAccessKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *pb.AccessKeysItem `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateAccessKeyResponse) Reset() {
	*x = CreateAccessKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccessKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccessKeyResponse) ProtoMessage() {}

func (x *CreateAccessKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccessKeyResponse.ProtoReflect.Descriptor instead.
func (*CreateAccessKeyResponse) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccessKeyResponse) GetData() *pb.AccessKeysItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteAccessKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAccessKeyRequest) Reset() {
	*x = DeleteAccessKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccessKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccessKeyRequest) ProtoMessage() {}

func (x *DeleteAccessKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccessKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccessKeyRequest) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteAccessKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAccessKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAccessKeyResponse) Reset() {
	*x = DeleteAccessKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccessKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccessKeyResponse) ProtoMessage() {}

func (x *DeleteAccessKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccessKeyResponse.ProtoReflect.Descriptor instead.
func (*DeleteAccessKeyResponse) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{3}
}

type GetAccessKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAccessKeyRequest) Reset() {
	*x = GetAccessKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessKeyRequest) ProtoMessage() {}

func (x *GetAccessKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessKeyRequest.ProtoReflect.Descriptor instead.
func (*GetAccessKeyRequest) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{4}
}

func (x *GetAccessKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAccessKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *pb.AccessKeysItem `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAccessKeyResponse) Reset() {
	*x = GetAccessKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessKeyResponse) ProtoMessage() {}

func (x *GetAccessKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessKeyResponse.ProtoReflect.Descriptor instead.
func (*GetAccessKeyResponse) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{5}
}

func (x *GetAccessKeyResponse) GetData() *pb.AccessKeysItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type DownloadAccessKeyFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DownloadAccessKeyFileRequest) Reset() {
	*x = DownloadAccessKeyFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadAccessKeyFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadAccessKeyFileRequest) ProtoMessage() {}

func (x *DownloadAccessKeyFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadAccessKeyFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadAccessKeyFileRequest) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{6}
}

func (x *DownloadAccessKeyFileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DownloadAccessKeyFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DownloadAccessKeyFileResponse) Reset() {
	*x = DownloadAccessKeyFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadAccessKeyFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadAccessKeyFileResponse) ProtoMessage() {}

func (x *DownloadAccessKeyFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadAccessKeyFileResponse.ProtoReflect.Descriptor instead.
func (*DownloadAccessKeyFileResponse) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{7}
}

type QueryAccessKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      pb.StatusEnum_Status           `protobuf:"varint,1,opt,name=status,proto3,enum=erda.core.services.authentication.credentials.accesskey.StatusEnum_Status" json:"status,omitempty"`
	SubjectType pb.SubjectTypeEnum_SubjectType `protobuf:"varint,2,opt,name=subjectType,proto3,enum=erda.core.services.authentication.credentials.accesskey.SubjectTypeEnum_SubjectType" json:"subjectType,omitempty"`
	Subject     string                         `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	AccessKey   string                         `protobuf:"bytes,4,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	PageNo      int64                          `protobuf:"varint,5,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize    int64                          `protobuf:"varint,6,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *QueryAccessKeysRequest) Reset() {
	*x = QueryAccessKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAccessKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAccessKeysRequest) ProtoMessage() {}

func (x *QueryAccessKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAccessKeysRequest.ProtoReflect.Descriptor instead.
func (*QueryAccessKeysRequest) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{8}
}

func (x *QueryAccessKeysRequest) GetStatus() pb.StatusEnum_Status {
	if x != nil {
		return x.Status
	}
	return pb.StatusEnum_NOT_SPECIFIED
}

func (x *QueryAccessKeysRequest) GetSubjectType() pb.SubjectTypeEnum_SubjectType {
	if x != nil {
		return x.SubjectType
	}
	return pb.SubjectTypeEnum_NOT_SPECIFIED
}

func (x *QueryAccessKeysRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *QueryAccessKeysRequest) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *QueryAccessKeysRequest) GetPageNo() int64 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *QueryAccessKeysRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type QueryAccessKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *QueryAccessKeysData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *QueryAccessKeysResponse) Reset() {
	*x = QueryAccessKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAccessKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAccessKeysResponse) ProtoMessage() {}

func (x *QueryAccessKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAccessKeysResponse.ProtoReflect.Descriptor instead.
func (*QueryAccessKeysResponse) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{9}
}

func (x *QueryAccessKeysResponse) GetData() *QueryAccessKeysData {
	if x != nil {
		return x.Data
	}
	return nil
}

type QueryAccessKeysData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*pb.AccessKeysItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64                `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *QueryAccessKeysData) Reset() {
	*x = QueryAccessKeysData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credential_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAccessKeysData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAccessKeysData) ProtoMessage() {}

func (x *QueryAccessKeysData) ProtoReflect() protoreflect.Message {
	mi := &file_credential_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAccessKeysData.ProtoReflect.Descriptor instead.
func (*QueryAccessKeysData) Descriptor() ([]byte, []int) {
	return file_credential_proto_rawDescGZIP(), []int{10}
}

func (x *QueryAccessKeysData) GetList() []*pb.AccessKeysItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *QueryAccessKeysData) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_credential_proto protoreflect.FileDescriptor

var file_credential_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x42, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x76, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x54, 0x2e, 0x65,
	0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x19, 0x0a,
	0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x73, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x2e, 0x0a, 0x1c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe0, 0x02, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x62, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x4a, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x76, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x54, 0x2e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x6b, 0x65, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x57, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x88, 0x01, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x5b, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65,
	0x79, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xae, 0x06, 0x0a,
	0x10, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x95, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x2b, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x73, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2d, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x9a, 0x01, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x2b, 0x2e,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x72, 0x64,
	0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26,
	0x2a, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2d, 0x6b, 0x65, 0x79,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x91, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2d, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xb0, 0x01, 0x0a, 0x15, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x31, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2a, 0x22, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2d,
	0x6b, 0x65, 0x79, 0x73, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x9d, 0x01,
	0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x73, 0x12, 0x2b, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x29, 0x22, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2d, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x39, 0x5a,
	0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61,
	0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_credential_proto_rawDescOnce sync.Once
	file_credential_proto_rawDescData = file_credential_proto_rawDesc
)

func file_credential_proto_rawDescGZIP() []byte {
	file_credential_proto_rawDescOnce.Do(func() {
		file_credential_proto_rawDescData = protoimpl.X.CompressGZIP(file_credential_proto_rawDescData)
	})
	return file_credential_proto_rawDescData
}

var file_credential_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_credential_proto_goTypes = []interface{}{
	(*CreateAccessKeyRequest)(nil),        // 0: erda.msp.credential.CreateAccessKeyRequest
	(*CreateAccessKeyResponse)(nil),       // 1: erda.msp.credential.CreateAccessKeyResponse
	(*DeleteAccessKeyRequest)(nil),        // 2: erda.msp.credential.DeleteAccessKeyRequest
	(*DeleteAccessKeyResponse)(nil),       // 3: erda.msp.credential.DeleteAccessKeyResponse
	(*GetAccessKeyRequest)(nil),           // 4: erda.msp.credential.GetAccessKeyRequest
	(*GetAccessKeyResponse)(nil),          // 5: erda.msp.credential.GetAccessKeyResponse
	(*DownloadAccessKeyFileRequest)(nil),  // 6: erda.msp.credential.DownloadAccessKeyFileRequest
	(*DownloadAccessKeyFileResponse)(nil), // 7: erda.msp.credential.DownloadAccessKeyFileResponse
	(*QueryAccessKeysRequest)(nil),        // 8: erda.msp.credential.QueryAccessKeysRequest
	(*QueryAccessKeysResponse)(nil),       // 9: erda.msp.credential.QueryAccessKeysResponse
	(*QueryAccessKeysData)(nil),           // 10: erda.msp.credential.QueryAccessKeysData
	(pb.SubjectTypeEnum_SubjectType)(0),   // 11: erda.core.services.authentication.credentials.accesskey.SubjectTypeEnum.SubjectType
	(*pb.AccessKeysItem)(nil),             // 12: erda.core.services.authentication.credentials.accesskey.AccessKeysItem
	(pb.StatusEnum_Status)(0),             // 13: erda.core.services.authentication.credentials.accesskey.StatusEnum.Status
}
var file_credential_proto_depIdxs = []int32{
	11, // 0: erda.msp.credential.CreateAccessKeyRequest.subjectType:type_name -> erda.core.services.authentication.credentials.accesskey.SubjectTypeEnum.SubjectType
	12, // 1: erda.msp.credential.CreateAccessKeyResponse.data:type_name -> erda.core.services.authentication.credentials.accesskey.AccessKeysItem
	12, // 2: erda.msp.credential.GetAccessKeyResponse.data:type_name -> erda.core.services.authentication.credentials.accesskey.AccessKeysItem
	13, // 3: erda.msp.credential.QueryAccessKeysRequest.status:type_name -> erda.core.services.authentication.credentials.accesskey.StatusEnum.Status
	11, // 4: erda.msp.credential.QueryAccessKeysRequest.subjectType:type_name -> erda.core.services.authentication.credentials.accesskey.SubjectTypeEnum.SubjectType
	10, // 5: erda.msp.credential.QueryAccessKeysResponse.data:type_name -> erda.msp.credential.QueryAccessKeysData
	12, // 6: erda.msp.credential.QueryAccessKeysData.list:type_name -> erda.core.services.authentication.credentials.accesskey.AccessKeysItem
	0,  // 7: erda.msp.credential.AccessKeyService.CreateAccessKey:input_type -> erda.msp.credential.CreateAccessKeyRequest
	2,  // 8: erda.msp.credential.AccessKeyService.DeleteAccessKey:input_type -> erda.msp.credential.DeleteAccessKeyRequest
	4,  // 9: erda.msp.credential.AccessKeyService.GetAccessKey:input_type -> erda.msp.credential.GetAccessKeyRequest
	6,  // 10: erda.msp.credential.AccessKeyService.DownloadAccessKeyFile:input_type -> erda.msp.credential.DownloadAccessKeyFileRequest
	8,  // 11: erda.msp.credential.AccessKeyService.QueryAccessKeys:input_type -> erda.msp.credential.QueryAccessKeysRequest
	1,  // 12: erda.msp.credential.AccessKeyService.CreateAccessKey:output_type -> erda.msp.credential.CreateAccessKeyResponse
	3,  // 13: erda.msp.credential.AccessKeyService.DeleteAccessKey:output_type -> erda.msp.credential.DeleteAccessKeyResponse
	5,  // 14: erda.msp.credential.AccessKeyService.GetAccessKey:output_type -> erda.msp.credential.GetAccessKeyResponse
	7,  // 15: erda.msp.credential.AccessKeyService.DownloadAccessKeyFile:output_type -> erda.msp.credential.DownloadAccessKeyFileResponse
	9,  // 16: erda.msp.credential.AccessKeyService.QueryAccessKeys:output_type -> erda.msp.credential.QueryAccessKeysResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_credential_proto_init() }
func file_credential_proto_init() {
	if File_credential_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_credential_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccessKeyRequest); i {
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
		file_credential_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccessKeyResponse); i {
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
		file_credential_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccessKeyRequest); i {
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
		file_credential_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccessKeyResponse); i {
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
		file_credential_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessKeyRequest); i {
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
		file_credential_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessKeyResponse); i {
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
		file_credential_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadAccessKeyFileRequest); i {
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
		file_credential_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadAccessKeyFileResponse); i {
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
		file_credential_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAccessKeysRequest); i {
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
		file_credential_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAccessKeysResponse); i {
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
		file_credential_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAccessKeysData); i {
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
			RawDescriptor: file_credential_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_credential_proto_goTypes,
		DependencyIndexes: file_credential_proto_depIdxs,
		MessageInfos:      file_credential_proto_msgTypes,
	}.Build()
	File_credential_proto = out.File
	file_credential_proto_rawDesc = nil
	file_credential_proto_goTypes = nil
	file_credential_proto_depIdxs = nil
}
