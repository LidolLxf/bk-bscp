// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: config_item.proto

package pbci

import (
	base "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/base"
	commit "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/commit"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// ConfigItem source resource reference: pkg/dal/table/config_item.go
type ConfigItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ConfigItemId uint32                `protobuf:"varint,2,opt,name=config_item_id,json=configItemId,proto3" json:"config_item_id,omitempty"`
	FileState    string                `protobuf:"bytes,3,opt,name=file_state,json=fileState,proto3" json:"file_state,omitempty"`
	Spec         *ConfigItemSpec       `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	CommitSpec   *commit.CommitSpec    `protobuf:"bytes,5,opt,name=commit_spec,json=commitSpec,proto3" json:"commit_spec,omitempty"`
	Attachment   *ConfigItemAttachment `protobuf:"bytes,6,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision     *base.Revision        `protobuf:"bytes,7,opt,name=revision,proto3" json:"revision,omitempty"`
	IsConflict   bool                  `protobuf:"varint,8,opt,name=is_conflict,json=isConflict,proto3" json:"is_conflict,omitempty"`
}

func (x *ConfigItem) Reset() {
	*x = ConfigItem{}
	mi := &file_config_item_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigItem) ProtoMessage() {}

func (x *ConfigItem) ProtoReflect() protoreflect.Message {
	mi := &file_config_item_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigItem.ProtoReflect.Descriptor instead.
func (*ConfigItem) Descriptor() ([]byte, []int) {
	return file_config_item_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigItem) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConfigItem) GetConfigItemId() uint32 {
	if x != nil {
		return x.ConfigItemId
	}
	return 0
}

func (x *ConfigItem) GetFileState() string {
	if x != nil {
		return x.FileState
	}
	return ""
}

func (x *ConfigItem) GetSpec() *ConfigItemSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *ConfigItem) GetCommitSpec() *commit.CommitSpec {
	if x != nil {
		return x.CommitSpec
	}
	return nil
}

func (x *ConfigItem) GetAttachment() *ConfigItemAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *ConfigItem) GetRevision() *base.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

func (x *ConfigItem) GetIsConflict() bool {
	if x != nil {
		return x.IsConflict
	}
	return false
}

// ConfigItemSpec source resource reference: pkg/dal/table/config_item.go
type ConfigItemSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path       string          `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	FileType   string          `protobuf:"bytes,3,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	FileMode   string          `protobuf:"bytes,4,opt,name=file_mode,json=fileMode,proto3" json:"file_mode,omitempty"`
	Memo       string          `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	Permission *FilePermission `protobuf:"bytes,6,opt,name=permission,proto3" json:"permission,omitempty"`
	Charset    string          `protobuf:"bytes,7,opt,name=charset,proto3" json:"charset,omitempty"`
}

func (x *ConfigItemSpec) Reset() {
	*x = ConfigItemSpec{}
	mi := &file_config_item_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigItemSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigItemSpec) ProtoMessage() {}

func (x *ConfigItemSpec) ProtoReflect() protoreflect.Message {
	mi := &file_config_item_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigItemSpec.ProtoReflect.Descriptor instead.
func (*ConfigItemSpec) Descriptor() ([]byte, []int) {
	return file_config_item_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigItemSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConfigItemSpec) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ConfigItemSpec) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *ConfigItemSpec) GetFileMode() string {
	if x != nil {
		return x.FileMode
	}
	return ""
}

func (x *ConfigItemSpec) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *ConfigItemSpec) GetPermission() *FilePermission {
	if x != nil {
		return x.Permission
	}
	return nil
}

func (x *ConfigItemSpec) GetCharset() string {
	if x != nil {
		return x.Charset
	}
	return ""
}

// ConfigItemAttachment source resource reference: pkg/dal/table/config_item.go
type ConfigItemAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	AppId uint32 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *ConfigItemAttachment) Reset() {
	*x = ConfigItemAttachment{}
	mi := &file_config_item_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigItemAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigItemAttachment) ProtoMessage() {}

func (x *ConfigItemAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_config_item_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigItemAttachment.ProtoReflect.Descriptor instead.
func (*ConfigItemAttachment) Descriptor() ([]byte, []int) {
	return file_config_item_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigItemAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *ConfigItemAttachment) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

// FilePermission source resource reference: pkg/dal/table/config_item.go
type FilePermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	UserGroup string `protobuf:"bytes,2,opt,name=user_group,json=userGroup,proto3" json:"user_group,omitempty"`
	Privilege string `protobuf:"bytes,3,opt,name=privilege,proto3" json:"privilege,omitempty"`
}

func (x *FilePermission) Reset() {
	*x = FilePermission{}
	mi := &file_config_item_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilePermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilePermission) ProtoMessage() {}

func (x *FilePermission) ProtoReflect() protoreflect.Message {
	mi := &file_config_item_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilePermission.ProtoReflect.Descriptor instead.
func (*FilePermission) Descriptor() ([]byte, []int) {
	return file_config_item_proto_rawDescGZIP(), []int{3}
}

func (x *FilePermission) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *FilePermission) GetUserGroup() string {
	if x != nil {
		return x.UserGroup
	}
	return ""
}

func (x *FilePermission) GetPrivilege() string {
	if x != nil {
		return x.Privilege
	}
	return ""
}

// ListConfigItemCounts source resource reference: pkg/dal/table/config_item.go
type ListConfigItemCounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId    uint32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Count    uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	UpdateAt string `protobuf:"bytes,3,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (x *ListConfigItemCounts) Reset() {
	*x = ListConfigItemCounts{}
	mi := &file_config_item_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListConfigItemCounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConfigItemCounts) ProtoMessage() {}

func (x *ListConfigItemCounts) ProtoReflect() protoreflect.Message {
	mi := &file_config_item_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConfigItemCounts.ProtoReflect.Descriptor instead.
func (*ListConfigItemCounts) Descriptor() ([]byte, []int) {
	return file_config_item_proto_rawDescGZIP(), []int{4}
}

func (x *ListConfigItemCounts) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *ListConfigItemCounts) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListConfigItemCounts) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

var File_config_item_proto protoreflect.FileDescriptor

var file_config_item_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x63, 0x69, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xee, 0x03, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x26, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x16,
	0x92, 0x41, 0x13, 0x32, 0x11, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe9, 0x85, 0x8d, 0xe7, 0xbd,
	0xae, 0xe9, 0xa1, 0xb9, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x16, 0x92, 0x41, 0x13, 0x32, 0x11, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe9,
	0x85, 0x8d, 0xe7, 0xbd, 0xae, 0xe9, 0xa1, 0xb9, 0x49, 0x44, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x5e, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3f, 0x92, 0x41,
	0x3c, 0x32, 0x3a, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe9, 0x85, 0x8d, 0xe7, 0xbd, 0xae, 0xe9,
	0xa1, 0xb9, 0xe7, 0x8a, 0xb6, 0xe6, 0x80, 0x81, 0xef, 0xbc, 0x9a, 0x28, 0x41, 0x44, 0x44, 0xe3,
	0x80, 0x81, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0xe3, 0x80, 0x81, 0x52, 0x45, 0x56, 0x49, 0x53,
	0x45, 0xe3, 0x80, 0x81, 0x55, 0x4e, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x29, 0x52, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x63, 0x69, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0a, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3a, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x70, 0x62, 0x63, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69,
	0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x42, 0x2e, 0x92, 0x41, 0x2b, 0x32, 0x29, 0xe6,
	0x98, 0xaf, 0xe5, 0x90, 0xa6, 0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0xe5, 0x86, 0xb2, 0xe7, 0xaa,
	0x81, 0xef, 0xbc, 0x9a, 0xe6, 0x98, 0xaf, 0x3d, 0x74, 0x72, 0x75, 0x65, 0xef, 0xbc, 0x8c, 0xe5,
	0x90, 0xa6, 0x3d, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x6c, 0x69, 0x63, 0x74, 0x22, 0xea, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49,
	0x74, 0x65, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41, 0x0b, 0x32, 0x09, 0xe6, 0x96, 0x87, 0xe4,
	0xbb, 0xb6, 0xe5, 0x90, 0x8d, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x32, 0x0c,
	0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe8, 0xb7, 0xaf, 0xe5, 0xbe, 0x84, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x60, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x43, 0x92, 0x41, 0x40, 0x32, 0x3e, 0xe9, 0x85, 0x8d, 0xe7,
	0xbd, 0xae, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0xa0, 0xbc, 0xe5, 0xbc, 0x8f, 0xef, 0xbc,
	0x9a, 0xe6, 0x96, 0x87, 0xe6, 0x9c, 0xac, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0x3d, 0x66, 0x69,
	0x6c, 0x65, 0x2c, 0x20, 0xe4, 0xba, 0x8c, 0xe8, 0xbf, 0x9b, 0xe5, 0x88, 0xb6, 0xe6, 0x96, 0x87,
	0xe4, 0xbb, 0xb6, 0x3d, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x92, 0x41, 0x14, 0x32, 0x0c, 0xe6, 0x96,
	0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0xa8, 0xa1, 0xe5, 0xbc, 0x8f, 0x3a, 0x04, 0x75, 0x6e, 0x69, 0x78,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x32, 0x0c, 0xe6,
	0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0x8f, 0x8f, 0xe8, 0xbf, 0xb0, 0x52, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x12, 0x34, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x63, 0x69, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x72, 0x73,
	0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65,
	0x74, 0x22, 0x62, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x62, 0x69, 0x7a,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0d, 0x92, 0x41, 0x0a, 0x32, 0x08,
	0xe4, 0xb8, 0x9a, 0xe5, 0x8a, 0xa1, 0x49, 0x44, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x0d, 0x92, 0x41, 0x0a, 0x32, 0x08, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x49, 0x44, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41, 0x0b, 0x32, 0x09, 0xe7, 0x94, 0xa8,
	0xe6, 0x88, 0xb7, 0xe5, 0x90, 0x8d, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x11, 0x92, 0x41, 0x0e, 0x32, 0x0c, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0xbb, 0x84,
	0xe5, 0x90, 0x8d, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x34,
	0x0a, 0x09, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x16, 0x92, 0x41, 0x13, 0x32, 0x0c, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe6, 0x9d,
	0x83, 0xe9, 0x99, 0x90, 0x4a, 0x03, 0x36, 0x34, 0x34, 0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x67, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0d, 0x92,
	0x41, 0x0a, 0x32, 0x08, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x49, 0x44, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x0b, 0x92, 0x41, 0x08, 0x32, 0x06, 0xe6, 0x80, 0xbb, 0xe6, 0x95, 0xb0, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x32, 0x0c,
	0xe6, 0x9b, 0xb4, 0xe6, 0x96, 0xb0, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4, 0x52, 0x08, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x75, 0x65,
	0x4b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6b, 0x2d, 0x62, 0x73, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2d, 0x69, 0x74, 0x65, 0x6d, 0x3b, 0x70, 0x62, 0x63, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_item_proto_rawDescOnce sync.Once
	file_config_item_proto_rawDescData = file_config_item_proto_rawDesc
)

func file_config_item_proto_rawDescGZIP() []byte {
	file_config_item_proto_rawDescOnce.Do(func() {
		file_config_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_item_proto_rawDescData)
	})
	return file_config_item_proto_rawDescData
}

var file_config_item_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_config_item_proto_goTypes = []any{
	(*ConfigItem)(nil),           // 0: pbci.ConfigItem
	(*ConfigItemSpec)(nil),       // 1: pbci.ConfigItemSpec
	(*ConfigItemAttachment)(nil), // 2: pbci.ConfigItemAttachment
	(*FilePermission)(nil),       // 3: pbci.FilePermission
	(*ListConfigItemCounts)(nil), // 4: pbci.ListConfigItemCounts
	(*commit.CommitSpec)(nil),    // 5: pbcommit.CommitSpec
	(*base.Revision)(nil),        // 6: pbbase.Revision
}
var file_config_item_proto_depIdxs = []int32{
	1, // 0: pbci.ConfigItem.spec:type_name -> pbci.ConfigItemSpec
	5, // 1: pbci.ConfigItem.commit_spec:type_name -> pbcommit.CommitSpec
	2, // 2: pbci.ConfigItem.attachment:type_name -> pbci.ConfigItemAttachment
	6, // 3: pbci.ConfigItem.revision:type_name -> pbbase.Revision
	3, // 4: pbci.ConfigItemSpec.permission:type_name -> pbci.FilePermission
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_config_item_proto_init() }
func file_config_item_proto_init() {
	if File_config_item_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_item_proto_goTypes,
		DependencyIndexes: file_config_item_proto_depIdxs,
		MessageInfos:      file_config_item_proto_msgTypes,
	}.Build()
	File_config_item_proto = out.File
	file_config_item_proto_rawDesc = nil
	file_config_item_proto_goTypes = nil
	file_config_item_proto_depIdxs = nil
}
