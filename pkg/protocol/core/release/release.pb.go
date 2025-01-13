// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: release.proto

package pbrelease

import (
	base "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/base"
	strategy "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/strategy"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Release source resource reference: pkg/dal/table/release.go
type Release struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *ReleaseSpec          `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status     *ReleaseStatus        `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Attachment *ReleaseAttachment    `protobuf:"bytes,4,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.CreatedRevision `protobuf:"bytes,5,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	mi := &file_release_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_release_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_release_proto_rawDescGZIP(), []int{0}
}

func (x *Release) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Release) GetSpec() *ReleaseSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Release) GetStatus() *ReleaseStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Release) GetAttachment() *ReleaseAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *Release) GetRevision() *base.CreatedRevision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// ReleaseSpec source resource reference: pkg/dal/table/release.go
type ReleaseSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Memo       string `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	Deprecated bool   `protobuf:"varint,3,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	PublishNum uint32 `protobuf:"varint,4,opt,name=publish_num,json=publishNum,proto3" json:"publish_num,omitempty"`
}

func (x *ReleaseSpec) Reset() {
	*x = ReleaseSpec{}
	mi := &file_release_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReleaseSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseSpec) ProtoMessage() {}

func (x *ReleaseSpec) ProtoReflect() protoreflect.Message {
	mi := &file_release_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseSpec.ProtoReflect.Descriptor instead.
func (*ReleaseSpec) Descriptor() ([]byte, []int) {
	return file_release_proto_rawDescGZIP(), []int{1}
}

func (x *ReleaseSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReleaseSpec) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *ReleaseSpec) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

func (x *ReleaseSpec) GetPublishNum() uint32 {
	if x != nil {
		return x.PublishNum
	}
	return 0
}

// ReleaseStatus status that not in db
type ReleaseStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublishStatus  string                         `protobuf:"bytes,1,opt,name=publish_status,json=publishStatus,proto3" json:"publish_status,omitempty"`
	ReleasedGroups []*ReleaseStatus_ReleasedGroup `protobuf:"bytes,2,rep,name=released_groups,json=releasedGroups,proto3" json:"released_groups,omitempty"`
	// 是否全量发布过,或者发布过默认分组
	FullyReleased  bool   `protobuf:"varint,3,opt,name=fully_released,json=fullyReleased,proto3" json:"fully_released,omitempty"`
	StrategyStatus string `protobuf:"bytes,4,opt,name=strategy_status,json=strategyStatus,proto3" json:"strategy_status,omitempty"`
}

func (x *ReleaseStatus) Reset() {
	*x = ReleaseStatus{}
	mi := &file_release_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReleaseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseStatus) ProtoMessage() {}

func (x *ReleaseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_release_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseStatus.ProtoReflect.Descriptor instead.
func (*ReleaseStatus) Descriptor() ([]byte, []int) {
	return file_release_proto_rawDescGZIP(), []int{2}
}

func (x *ReleaseStatus) GetPublishStatus() string {
	if x != nil {
		return x.PublishStatus
	}
	return ""
}

func (x *ReleaseStatus) GetReleasedGroups() []*ReleaseStatus_ReleasedGroup {
	if x != nil {
		return x.ReleasedGroups
	}
	return nil
}

func (x *ReleaseStatus) GetFullyReleased() bool {
	if x != nil {
		return x.FullyReleased
	}
	return false
}

func (x *ReleaseStatus) GetStrategyStatus() string {
	if x != nil {
		return x.StrategyStatus
	}
	return ""
}

// ReleaseAttachment source resource reference: pkg/dal/table/release.go
type ReleaseAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	AppId uint32 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *ReleaseAttachment) Reset() {
	*x = ReleaseAttachment{}
	mi := &file_release_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReleaseAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseAttachment) ProtoMessage() {}

func (x *ReleaseAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_release_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseAttachment.ProtoReflect.Descriptor instead.
func (*ReleaseAttachment) Descriptor() ([]byte, []int) {
	return file_release_proto_rawDescGZIP(), []int{3}
}

func (x *ReleaseAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *ReleaseAttachment) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

// PublishRecord list publish relate field
type PublishRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublishTime       string          `protobuf:"bytes,1,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	Name              string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Scope             *strategy.Scope `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
	Creator           string          `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	FullyReleased     bool            `protobuf:"varint,5,opt,name=fully_released,json=fullyReleased,proto3" json:"fully_released,omitempty"`
	UpdatedAt         string          `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	FinalApprovalTime string          `protobuf:"bytes,7,opt,name=final_approval_time,json=finalApprovalTime,proto3" json:"final_approval_time,omitempty"`
}

func (x *PublishRecord) Reset() {
	*x = PublishRecord{}
	mi := &file_release_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRecord) ProtoMessage() {}

func (x *PublishRecord) ProtoReflect() protoreflect.Message {
	mi := &file_release_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRecord.ProtoReflect.Descriptor instead.
func (*PublishRecord) Descriptor() ([]byte, []int) {
	return file_release_proto_rawDescGZIP(), []int{4}
}

func (x *PublishRecord) GetPublishTime() string {
	if x != nil {
		return x.PublishTime
	}
	return ""
}

func (x *PublishRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PublishRecord) GetScope() *strategy.Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *PublishRecord) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *PublishRecord) GetFullyReleased() bool {
	if x != nil {
		return x.FullyReleased
	}
	return false
}

func (x *PublishRecord) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PublishRecord) GetFinalApprovalTime() string {
	if x != nil {
		return x.FinalApprovalTime
	}
	return ""
}

type ReleaseStatus_ReleasedGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mode        string           `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
	OldSelector *structpb.Struct `protobuf:"bytes,4,opt,name=old_selector,json=oldSelector,proto3" json:"old_selector,omitempty"`
	NewSelector *structpb.Struct `protobuf:"bytes,5,opt,name=new_selector,json=newSelector,proto3" json:"new_selector,omitempty"`
	Uid         string           `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	Edited      bool             `protobuf:"varint,7,opt,name=edited,proto3" json:"edited,omitempty"`
}

func (x *ReleaseStatus_ReleasedGroup) Reset() {
	*x = ReleaseStatus_ReleasedGroup{}
	mi := &file_release_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReleaseStatus_ReleasedGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseStatus_ReleasedGroup) ProtoMessage() {}

func (x *ReleaseStatus_ReleasedGroup) ProtoReflect() protoreflect.Message {
	mi := &file_release_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseStatus_ReleasedGroup.ProtoReflect.Descriptor instead.
func (*ReleaseStatus_ReleasedGroup) Descriptor() ([]byte, []int) {
	return file_release_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ReleaseStatus_ReleasedGroup) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReleaseStatus_ReleasedGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReleaseStatus_ReleasedGroup) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *ReleaseStatus_ReleasedGroup) GetOldSelector() *structpb.Struct {
	if x != nil {
		return x.OldSelector
	}
	return nil
}

func (x *ReleaseStatus_ReleasedGroup) GetNewSelector() *structpb.Struct {
	if x != nil {
		return x.NewSelector
	}
	return nil
}

func (x *ReleaseStatus_ReleasedGroup) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ReleaseStatus_ReleasedGroup) GetEdited() bool {
	if x != nil {
		return x.Edited
	}
	return false
}

var File_release_proto protoreflect.FileDescriptor

var file_release_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x70, 0x62, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x13,
	0x92, 0x41, 0x10, 0x32, 0x0e, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0xe7, 0x89, 0x88, 0xe6, 0x9c,
	0xac, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xd9, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x32, 0x0c, 0xe7, 0x89, 0x88,
	0xe6, 0x9c, 0xac, 0xe5, 0x90, 0x8d, 0xe7, 0xa7, 0xb0, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x25, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x92,
	0x41, 0x0e, 0x32, 0x0c, 0xe7, 0x89, 0x88, 0xe6, 0x9c, 0xac, 0xe6, 0x8f, 0x8f, 0xe8, 0xbf, 0xb0,
	0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x48, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x28, 0x92, 0x41, 0x25, 0x32,
	0x23, 0xe6, 0x98, 0xaf, 0xe5, 0x90, 0xa6, 0xe5, 0xbc, 0x83, 0xe7, 0x94, 0xa8, 0xef, 0xbc, 0x9a,
	0xe6, 0x98, 0xaf, 0x3d, 0x74, 0x72, 0x75, 0x65, 0xef, 0xbc, 0x8c, 0xe5, 0x90, 0xa6, 0x3d, 0x66,
	0x61, 0x6c, 0x73, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x32, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x32, 0x0c, 0xe5, 0x8f, 0x91, 0xe5,
	0xb8, 0x83, 0xe6, 0x95, 0xb0, 0xe9, 0x87, 0x8f, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x4e, 0x75, 0x6d, 0x22, 0x87, 0x06, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x79, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x52,
	0x92, 0x41, 0x4f, 0x32, 0x4d, 0xe7, 0x89, 0x88, 0xe6, 0x9c, 0xac, 0xe7, 0x8a, 0xb6, 0xe6, 0x80,
	0x81, 0xef, 0xbc, 0x88, 0xe6, 0x9e, 0x9a, 0xe4, 0xb8, 0xbe, 0xe5, 0x80, 0xbc, 0xef, 0xbc, 0x9a,
	0x6e, 0x6f, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0xe3, 0x80, 0x81, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0xe3,
	0x80, 0x81, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0xef,
	0xbc, 0x89, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x4f, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x62, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x5f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x66, 0x75, 0x6c, 0x6c,
	0x79, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0xd9, 0x03, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x23, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x13, 0x92, 0x41, 0x10, 0x32, 0x0e, 0xe7, 0x89, 0x88, 0xe6, 0x9c, 0xac, 0xe5, 0x88, 0x86,
	0xe7, 0xbb, 0x84, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0x92, 0x41, 0x17, 0x32, 0x15, 0xe6, 0x9c,
	0x8d, 0xe5, 0x8a, 0xa1, 0xe7, 0x89, 0x88, 0xe6, 0x9c, 0xac, 0xe5, 0x88, 0x86, 0xe7, 0xbb, 0x84,
	0xe5, 0x90, 0x8d, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3b, 0x92, 0x41, 0x38, 0x32, 0x36, 0xe5, 0x88,
	0x86, 0xe7, 0xbb, 0x84, 0xe7, 0xb1, 0xbb, 0xe5, 0x9e, 0x8b, 0xef, 0xbc, 0x88, 0xe6, 0x9e, 0x9a,
	0xe4, 0xb8, 0xbe, 0xe5, 0x80, 0xbc, 0xef, 0xbc, 0x9a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0xe3,
	0x80, 0x81, 0x64, 0x65, 0x62, 0x75, 0x67, 0xe3, 0x80, 0x81, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0xef, 0xbc, 0x89, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x5c, 0x0a, 0x0c, 0x6f, 0x6c,
	0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x20, 0x92, 0x41, 0x1d, 0x32, 0x1b,
	0xe5, 0xbd, 0x93, 0xe5, 0x89, 0x8d, 0xe7, 0xba, 0xbf, 0xe4, 0xb8, 0x8a, 0xe5, 0x88, 0x86, 0xe7,
	0xbb, 0x84, 0xe9, 0x80, 0x89, 0xe6, 0x8b, 0xa9, 0xe5, 0x99, 0xa8, 0x52, 0x0b, 0x6f, 0x6c, 0x64,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x5c, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x20, 0x92, 0x41, 0x1d, 0x32, 0x1b, 0xe4, 0xbf,
	0xae, 0xe6, 0x94, 0xb9, 0xe5, 0x90, 0x8e, 0xe7, 0x9a, 0x84, 0xe5, 0x88, 0x86, 0xe7, 0xbb, 0x84,
	0xe9, 0x80, 0x89, 0xe6, 0x8b, 0xa9, 0xe5, 0x99, 0xa8, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1d, 0x92, 0x41, 0x1a, 0x32, 0x18, 0xe8, 0xb0, 0x83, 0xe8, 0xaf, 0x95,
	0xe5, 0x88, 0x86, 0xe7, 0xbb, 0x84, 0xe7, 0x9a, 0x84, 0xe5, 0xae, 0x9e, 0xe4, 0xbe, 0x8b, 0x55,
	0x49, 0x44, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x65, 0x64, 0x69, 0x74, 0x65,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x42, 0x1d, 0x92, 0x41, 0x1a, 0x32, 0x18, 0xe6, 0x98,
	0xaf, 0xe5, 0x90, 0xa6, 0xe4, 0xb8, 0xba, 0xe5, 0xb7, 0xb2, 0xe7, 0xbc, 0x96, 0xe8, 0xbe, 0x91,
	0xe7, 0x8a, 0xb6, 0xe6, 0x80, 0x81, 0x52, 0x06, 0x65, 0x64, 0x69, 0x74, 0x65, 0x64, 0x22, 0x5f,
	0x0a, 0x11, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x0d, 0x92, 0x41, 0x0a, 0x32, 0x08, 0xe4, 0xb8, 0x9a, 0xe5, 0x8a, 0xa1,
	0x49, 0x44, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0d, 0x92, 0x41, 0x0a, 0x32, 0x08,
	0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x49, 0x44, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22,
	0xff, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x66,
	0x75, 0x6c, 0x6c, 0x79, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x75, 0x65, 0x4b, 0x69, 0x6e, 0x67, 0x2f,
	0x62, 0x6b, 0x2d, 0x62, 0x73, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x3b, 0x70, 0x62, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_release_proto_rawDescOnce sync.Once
	file_release_proto_rawDescData = file_release_proto_rawDesc
)

func file_release_proto_rawDescGZIP() []byte {
	file_release_proto_rawDescOnce.Do(func() {
		file_release_proto_rawDescData = protoimpl.X.CompressGZIP(file_release_proto_rawDescData)
	})
	return file_release_proto_rawDescData
}

var file_release_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_release_proto_goTypes = []any{
	(*Release)(nil),                     // 0: pbrelease.Release
	(*ReleaseSpec)(nil),                 // 1: pbrelease.ReleaseSpec
	(*ReleaseStatus)(nil),               // 2: pbrelease.ReleaseStatus
	(*ReleaseAttachment)(nil),           // 3: pbrelease.ReleaseAttachment
	(*PublishRecord)(nil),               // 4: pbrelease.PublishRecord
	(*ReleaseStatus_ReleasedGroup)(nil), // 5: pbrelease.ReleaseStatus.ReleasedGroup
	(*base.CreatedRevision)(nil),        // 6: pbbase.CreatedRevision
	(*strategy.Scope)(nil),              // 7: pbstrategy.Scope
	(*structpb.Struct)(nil),             // 8: google.protobuf.Struct
}
var file_release_proto_depIdxs = []int32{
	1, // 0: pbrelease.Release.spec:type_name -> pbrelease.ReleaseSpec
	2, // 1: pbrelease.Release.status:type_name -> pbrelease.ReleaseStatus
	3, // 2: pbrelease.Release.attachment:type_name -> pbrelease.ReleaseAttachment
	6, // 3: pbrelease.Release.revision:type_name -> pbbase.CreatedRevision
	5, // 4: pbrelease.ReleaseStatus.released_groups:type_name -> pbrelease.ReleaseStatus.ReleasedGroup
	7, // 5: pbrelease.PublishRecord.scope:type_name -> pbstrategy.Scope
	8, // 6: pbrelease.ReleaseStatus.ReleasedGroup.old_selector:type_name -> google.protobuf.Struct
	8, // 7: pbrelease.ReleaseStatus.ReleasedGroup.new_selector:type_name -> google.protobuf.Struct
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_release_proto_init() }
func file_release_proto_init() {
	if File_release_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_release_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_release_proto_goTypes,
		DependencyIndexes: file_release_proto_depIdxs,
		MessageInfos:      file_release_proto_msgTypes,
	}.Build()
	File_release_proto = out.File
	file_release_proto_rawDesc = nil
	file_release_proto_goTypes = nil
	file_release_proto_depIdxs = nil
}
