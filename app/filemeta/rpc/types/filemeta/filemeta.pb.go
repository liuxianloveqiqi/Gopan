// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: filemeta.proto

package filemeta

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

// 文件元信息
type FileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FileSha1   string `protobuf:"bytes,2,opt,name=FileSha1,proto3" json:"FileSha1,omitempty"`
	FileSize   int64  `protobuf:"varint,3,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	FileName   string `protobuf:"bytes,4,opt,name=FileName,proto3" json:"FileName,omitempty"`
	FileAddr   string `protobuf:"bytes,6,opt,name=FileAddr,proto3" json:"FileAddr,omitempty"`
	Status     int64  `protobuf:"varint,7,opt,name=Status,proto3" json:"Status,omitempty"`
	CreateTime string `protobuf:"bytes,8,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime string `protobuf:"bytes,9,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *FileMeta) Reset() {
	*x = FileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filemeta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMeta) ProtoMessage() {}

func (x *FileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_filemeta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMeta.ProtoReflect.Descriptor instead.
func (*FileMeta) Descriptor() ([]byte, []int) {
	return file_filemeta_proto_rawDescGZIP(), []int{0}
}

func (x *FileMeta) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FileMeta) GetFileSha1() string {
	if x != nil {
		return x.FileSha1
	}
	return ""
}

func (x *FileMeta) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FileMeta) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileMeta) GetFileAddr() string {
	if x != nil {
		return x.FileAddr
	}
	return ""
}

func (x *FileMeta) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FileMeta) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *FileMeta) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type UserFileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FileSha1   string `protobuf:"bytes,2,opt,name=FileSha1,proto3" json:"FileSha1,omitempty"`
	FileSize   int64  `protobuf:"varint,3,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	FileName   string `protobuf:"bytes,4,opt,name=FileName,proto3" json:"FileName,omitempty"`
	Status     int64  `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	CreateTime string `protobuf:"bytes,6,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime string `protobuf:"bytes,7,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	UserId     int64  `protobuf:"varint,8,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *UserFileMeta) Reset() {
	*x = UserFileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filemeta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFileMeta) ProtoMessage() {}

func (x *UserFileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_filemeta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFileMeta.ProtoReflect.Descriptor instead.
func (*UserFileMeta) Descriptor() ([]byte, []int) {
	return file_filemeta_proto_rawDescGZIP(), []int{1}
}

func (x *UserFileMeta) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserFileMeta) GetFileSha1() string {
	if x != nil {
		return x.FileSha1
	}
	return ""
}

func (x *UserFileMeta) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *UserFileMeta) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UserFileMeta) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserFileMeta) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *UserFileMeta) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *UserFileMeta) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetFileMetaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileSha1 string `protobuf:"bytes,1,opt,name=FileSha1,proto3" json:"FileSha1,omitempty"`
}

func (x *GetFileMetaReq) Reset() {
	*x = GetFileMetaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filemeta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileMetaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileMetaReq) ProtoMessage() {}

func (x *GetFileMetaReq) ProtoReflect() protoreflect.Message {
	mi := &file_filemeta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileMetaReq.ProtoReflect.Descriptor instead.
func (*GetFileMetaReq) Descriptor() ([]byte, []int) {
	return file_filemeta_proto_rawDescGZIP(), []int{2}
}

func (x *GetFileMetaReq) GetFileSha1() string {
	if x != nil {
		return x.FileSha1
	}
	return ""
}

type GetUserFileMetaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetUserFileMetaReq) Reset() {
	*x = GetUserFileMetaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filemeta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserFileMetaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserFileMetaReq) ProtoMessage() {}

func (x *GetUserFileMetaReq) ProtoReflect() protoreflect.Message {
	mi := &file_filemeta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserFileMetaReq.ProtoReflect.Descriptor instead.
func (*GetUserFileMetaReq) Descriptor() ([]byte, []int) {
	return file_filemeta_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserFileMetaReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserFileMetaResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserFileMetaList []*UserFileMeta `protobuf:"bytes,1,rep,name=UserFileMetaList,proto3" json:"UserFileMetaList,omitempty"`
}

func (x *GetUserFileMetaResp) Reset() {
	*x = GetUserFileMetaResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filemeta_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserFileMetaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserFileMetaResp) ProtoMessage() {}

func (x *GetUserFileMetaResp) ProtoReflect() protoreflect.Message {
	mi := &file_filemeta_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserFileMetaResp.ProtoReflect.Descriptor instead.
func (*GetUserFileMetaResp) Descriptor() ([]byte, []int) {
	return file_filemeta_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserFileMetaResp) GetUserFileMetaList() []*UserFileMeta {
	if x != nil {
		return x.UserFileMetaList
	}
	return nil
}

type UpdataUserFileMetaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=FileName,proto3" json:"FileName,omitempty"`
	FileSha1 string `protobuf:"bytes,3,opt,name=FileSha1,proto3" json:"FileSha1,omitempty"`
}

func (x *UpdataUserFileMetaReq) Reset() {
	*x = UpdataUserFileMetaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filemeta_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdataUserFileMetaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdataUserFileMetaReq) ProtoMessage() {}

func (x *UpdataUserFileMetaReq) ProtoReflect() protoreflect.Message {
	mi := &file_filemeta_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdataUserFileMetaReq.ProtoReflect.Descriptor instead.
func (*UpdataUserFileMetaReq) Descriptor() ([]byte, []int) {
	return file_filemeta_proto_rawDescGZIP(), []int{5}
}

func (x *UpdataUserFileMetaReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdataUserFileMetaReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UpdataUserFileMetaReq) GetFileSha1() string {
	if x != nil {
		return x.FileSha1
	}
	return ""
}

type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filemeta_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_filemeta_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_filemeta_proto_rawDescGZIP(), []int{6}
}

func (x *CommonResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_filemeta_proto protoreflect.FileDescriptor

var file_filemeta_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xe2, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61,
	0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61,
	0x31, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xe2, 0x01,
	0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x31,
	0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x57,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x10, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x67, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x61, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x31,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x31,
	0x22, 0x3a, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd8, 0x01, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x1a, 0x10, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x4a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x1a, 0x1b, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x47,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x61, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x61, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filemeta_proto_rawDescOnce sync.Once
	file_filemeta_proto_rawDescData = file_filemeta_proto_rawDesc
)

func file_filemeta_proto_rawDescGZIP() []byte {
	file_filemeta_proto_rawDescOnce.Do(func() {
		file_filemeta_proto_rawDescData = protoimpl.X.CompressGZIP(file_filemeta_proto_rawDescData)
	})
	return file_filemeta_proto_rawDescData
}

var file_filemeta_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_filemeta_proto_goTypes = []interface{}{
	(*FileMeta)(nil),              // 0: upload.FileMeta
	(*UserFileMeta)(nil),          // 1: upload.UserFileMeta
	(*GetFileMetaReq)(nil),        // 2: upload.GetFileMetaReq
	(*GetUserFileMetaReq)(nil),    // 3: upload.GetUserFileMetaReq
	(*GetUserFileMetaResp)(nil),   // 4: upload.GetUserFileMetaResp
	(*UpdataUserFileMetaReq)(nil), // 5: upload.UpdataUserFileMetaReq
	(*CommonResp)(nil),            // 6: upload.CommonResp
}
var file_filemeta_proto_depIdxs = []int32{
	1, // 0: upload.GetUserFileMetaResp.UserFileMetaList:type_name -> upload.UserFileMeta
	2, // 1: upload.filemeta.GetFileMeta:input_type -> upload.GetFileMetaReq
	3, // 2: upload.filemeta.GetUserFileMeta:input_type -> upload.GetUserFileMetaReq
	5, // 3: upload.filemeta.UpdataUserFileMeta:input_type -> upload.UpdataUserFileMetaReq
	0, // 4: upload.filemeta.GetFileMeta:output_type -> upload.FileMeta
	4, // 5: upload.filemeta.GetUserFileMeta:output_type -> upload.GetUserFileMetaResp
	6, // 6: upload.filemeta.UpdataUserFileMeta:output_type -> upload.CommonResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_filemeta_proto_init() }
func file_filemeta_proto_init() {
	if File_filemeta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filemeta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMeta); i {
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
		file_filemeta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFileMeta); i {
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
		file_filemeta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileMetaReq); i {
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
		file_filemeta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserFileMetaReq); i {
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
		file_filemeta_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserFileMetaResp); i {
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
		file_filemeta_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdataUserFileMetaReq); i {
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
		file_filemeta_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
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
			RawDescriptor: file_filemeta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filemeta_proto_goTypes,
		DependencyIndexes: file_filemeta_proto_depIdxs,
		MessageInfos:      file_filemeta_proto_msgTypes,
	}.Build()
	File_filemeta_proto = out.File
	file_filemeta_proto_rawDesc = nil
	file_filemeta_proto_goTypes = nil
	file_filemeta_proto_depIdxs = nil
}
