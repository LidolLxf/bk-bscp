// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: released_ci.proto

package pbrci

import (
	base "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/base"
	commit "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/commit"
	config_item "github.com/TencentBlueKing/bk-bscp/pkg/protocol/core/config-item"
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

// ReleasedConfigItem source resource reference: pkg/dal/table/release_ci.go
type ReleasedConfigItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32                            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ReleaseId    uint32                            `protobuf:"varint,2,opt,name=release_id,json=releaseId,proto3" json:"release_id,omitempty"`
	CommitId     uint32                            `protobuf:"varint,3,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	CommitSpec   *commit.ReleasedCommitSpec        `protobuf:"bytes,4,opt,name=commit_spec,json=commitSpec,proto3" json:"commit_spec,omitempty"`
	ConfigItemId uint32                            `protobuf:"varint,5,opt,name=config_item_id,json=configItemId,proto3" json:"config_item_id,omitempty"`
	Spec         *config_item.ConfigItemSpec       `protobuf:"bytes,6,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment   *config_item.ConfigItemAttachment `protobuf:"bytes,7,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision     *base.Revision                    `protobuf:"bytes,8,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *ReleasedConfigItem) Reset() {
	*x = ReleasedConfigItem{}
	mi := &file_released_ci_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReleasedConfigItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleasedConfigItem) ProtoMessage() {}

func (x *ReleasedConfigItem) ProtoReflect() protoreflect.Message {
	mi := &file_released_ci_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleasedConfigItem.ProtoReflect.Descriptor instead.
func (*ReleasedConfigItem) Descriptor() ([]byte, []int) {
	return file_released_ci_proto_rawDescGZIP(), []int{0}
}

func (x *ReleasedConfigItem) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReleasedConfigItem) GetReleaseId() uint32 {
	if x != nil {
		return x.ReleaseId
	}
	return 0
}

func (x *ReleasedConfigItem) GetCommitId() uint32 {
	if x != nil {
		return x.CommitId
	}
	return 0
}

func (x *ReleasedConfigItem) GetCommitSpec() *commit.ReleasedCommitSpec {
	if x != nil {
		return x.CommitSpec
	}
	return nil
}

func (x *ReleasedConfigItem) GetConfigItemId() uint32 {
	if x != nil {
		return x.ConfigItemId
	}
	return 0
}

func (x *ReleasedConfigItem) GetSpec() *config_item.ConfigItemSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *ReleasedConfigItem) GetAttachment() *config_item.ConfigItemAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *ReleasedConfigItem) GetRevision() *base.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

var File_released_ci_proto protoreflect.FileDescriptor

var file_released_ci_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x63, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x62, 0x72, 0x63, 0x69, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2d, 0x69,
	0x74, 0x65, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x03, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x32, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x22, 0x92, 0x41, 0x1f, 0x32, 0x1d, 0xe6,
	0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0xe7, 0x89, 0x88, 0xe6, 0x9c, 0xac, 0xe6, 0x96, 0x87, 0xe4, 0xbb,
	0xb6, 0xe9, 0x85, 0x8d, 0xe7, 0xbd, 0xae, 0xe9, 0xa1, 0xb9, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x32, 0x0a, 0x0a, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x13, 0x92, 0x41, 0x10, 0x32, 0x0e, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a,
	0xa1, 0xe7, 0x89, 0x88, 0xe6, 0x9c, 0xac, 0x49, 0x44, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x22, 0x92, 0x41, 0x1f, 0x32, 0x1d, 0xe6, 0x96,
	0x87, 0xe4, 0xbb, 0xb6, 0xe9, 0x85, 0x8d, 0xe7, 0xbd, 0xae, 0xe9, 0xa1, 0xb9, 0xe7, 0x89, 0x88,
	0xe6, 0x9c, 0xac, 0xe8, 0xae, 0xb0, 0xe5, 0xbd, 0x95, 0x49, 0x44, 0x52, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x3c, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x16, 0x92, 0x41,
	0x13, 0x32, 0x11, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe9, 0x85, 0x8d, 0xe7, 0xbd, 0xae, 0xe9,
	0xa1, 0xb9, 0x49, 0x44, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x62, 0x63, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74,
	0x65, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3a, 0x0a, 0x0a,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x62, 0x63, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74,
	0x65, 0x6d, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x75, 0x65,
	0x4b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6b, 0x2d, 0x62, 0x73, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x2d, 0x63, 0x69, 0x3b, 0x70, 0x62, 0x72, 0x63, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_released_ci_proto_rawDescOnce sync.Once
	file_released_ci_proto_rawDescData = file_released_ci_proto_rawDesc
)

func file_released_ci_proto_rawDescGZIP() []byte {
	file_released_ci_proto_rawDescOnce.Do(func() {
		file_released_ci_proto_rawDescData = protoimpl.X.CompressGZIP(file_released_ci_proto_rawDescData)
	})
	return file_released_ci_proto_rawDescData
}

var file_released_ci_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_released_ci_proto_goTypes = []any{
	(*ReleasedConfigItem)(nil),               // 0: pbrci.ReleasedConfigItem
	(*commit.ReleasedCommitSpec)(nil),        // 1: pbcommit.ReleasedCommitSpec
	(*config_item.ConfigItemSpec)(nil),       // 2: pbci.ConfigItemSpec
	(*config_item.ConfigItemAttachment)(nil), // 3: pbci.ConfigItemAttachment
	(*base.Revision)(nil),                    // 4: pbbase.Revision
}
var file_released_ci_proto_depIdxs = []int32{
	1, // 0: pbrci.ReleasedConfigItem.commit_spec:type_name -> pbcommit.ReleasedCommitSpec
	2, // 1: pbrci.ReleasedConfigItem.spec:type_name -> pbci.ConfigItemSpec
	3, // 2: pbrci.ReleasedConfigItem.attachment:type_name -> pbci.ConfigItemAttachment
	4, // 3: pbrci.ReleasedConfigItem.revision:type_name -> pbbase.Revision
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_released_ci_proto_init() }
func file_released_ci_proto_init() {
	if File_released_ci_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_released_ci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_released_ci_proto_goTypes,
		DependencyIndexes: file_released_ci_proto_depIdxs,
		MessageInfos:      file_released_ci_proto_msgTypes,
	}.Build()
	File_released_ci_proto = out.File
	file_released_ci_proto_rawDesc = nil
	file_released_ci_proto_goTypes = nil
	file_released_ci_proto_depIdxs = nil
}
