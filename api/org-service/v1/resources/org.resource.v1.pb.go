// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.6
// source: api/org-service/v1/resources/org.resource.v1.proto

package resourcev1

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

type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_api_org_service_v1_resources_org_resource_v1_proto_rawDescGZIP(), []int{0}
}

func (x *PingReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason   string            `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message  string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data     *PingRespData     `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PingResp) Reset() {
	*x = PingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResp) ProtoMessage() {}

func (x *PingResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResp.ProtoReflect.Descriptor instead.
func (*PingResp) Descriptor() ([]byte, []int) {
	return file_api_org_service_v1_resources_org_resource_v1_proto_rawDescGZIP(), []int{1}
}

func (x *PingResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PingResp) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *PingResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PingResp) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PingResp) GetData() *PingRespData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PingRespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingRespData) Reset() {
	*x = PingRespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRespData) ProtoMessage() {}

func (x *PingRespData) ProtoReflect() protoreflect.Message {
	mi := &file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRespData.ProtoReflect.Descriptor instead.
func (*PingRespData) Descriptor() ([]byte, []int) {
	return file_api_org_service_v1_resources_org_resource_v1_proto_rawDescGZIP(), []int{2}
}

func (x *PingRespData) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Org ENGINE InnoDB CHARSET utf8mb4 COMMENT '组织'
type Org struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// created_time 创建时间
	CreatedTime string `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// updated_time 最后修改时间
	UpdatedTime string `protobuf:"bytes,3,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// deleted_time 删除时间
	DeletedTime uint64 `protobuf:"varint,4,opt,name=deleted_time,json=deletedTime,proto3" json:"deleted_time,omitempty"`
	// org_id uuid
	OrgId uint64 `protobuf:"varint,5,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// org_name 组织名称
	OrgName string `protobuf:"bytes,6,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	// org_avatar 组织头像
	OrgAvatar string `protobuf:"bytes,7,opt,name=org_avatar,json=orgAvatar,proto3" json:"org_avatar,omitempty"`
	// org_contact_name 组织联系名称
	OrgContactName string `protobuf:"bytes,8,opt,name=org_contact_name,json=orgContactName,proto3" json:"org_contact_name,omitempty"`
	// org_contact_phone 组织联系手机
	OrgContactPhone string `protobuf:"bytes,9,opt,name=org_contact_phone,json=orgContactPhone,proto3" json:"org_contact_phone,omitempty"`
	// org_contact_email 组织联系邮箱
	OrgContactEmail string `protobuf:"bytes,10,opt,name=org_contact_email,json=orgContactEmail,proto3" json:"org_contact_email,omitempty"`
	// org_type 组织状态；1：个人版，2：标准组织，3：。。。
	OrgType uint32 `protobuf:"varint,11,opt,name=org_type,json=orgType,proto3" json:"org_type,omitempty"`
	// org_status 状态；0：INIT，1：ENABLE，2：DISABLE，3：DELETED
	OrgStatus uint32 `protobuf:"varint,12,opt,name=org_status,json=orgStatus,proto3" json:"org_status,omitempty"`
	// org_industry_id 组织行业
	OrgIndustryId uint64 `protobuf:"varint,13,opt,name=org_industry_id,json=orgIndustryId,proto3" json:"org_industry_id,omitempty"`
	// org_scale_id 组织规模
	OrgScaleId uint64 `protobuf:"varint,14,opt,name=org_scale_id,json=orgScaleId,proto3" json:"org_scale_id,omitempty"`
	// org_address 组织地址
	OrgAddress string `protobuf:"bytes,15,opt,name=org_address,json=orgAddress,proto3" json:"org_address,omitempty"`
	// org_zip_code 组织地址邮编
	OrgZipCode string `protobuf:"bytes,16,opt,name=org_zip_code,json=orgZipCode,proto3" json:"org_zip_code,omitempty"`
	// org_creator_id 组织创建者
	OrgCreatorId uint64 `protobuf:"varint,17,opt,name=org_creator_id,json=orgCreatorId,proto3" json:"org_creator_id,omitempty"`
}

func (x *Org) Reset() {
	*x = Org{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Org) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Org) ProtoMessage() {}

func (x *Org) ProtoReflect() protoreflect.Message {
	mi := &file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Org.ProtoReflect.Descriptor instead.
func (*Org) Descriptor() ([]byte, []int) {
	return file_api_org_service_v1_resources_org_resource_v1_proto_rawDescGZIP(), []int{3}
}

func (x *Org) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Org) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *Org) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *Org) GetDeletedTime() uint64 {
	if x != nil {
		return x.DeletedTime
	}
	return 0
}

func (x *Org) GetOrgId() uint64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *Org) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *Org) GetOrgAvatar() string {
	if x != nil {
		return x.OrgAvatar
	}
	return ""
}

func (x *Org) GetOrgContactName() string {
	if x != nil {
		return x.OrgContactName
	}
	return ""
}

func (x *Org) GetOrgContactPhone() string {
	if x != nil {
		return x.OrgContactPhone
	}
	return ""
}

func (x *Org) GetOrgContactEmail() string {
	if x != nil {
		return x.OrgContactEmail
	}
	return ""
}

func (x *Org) GetOrgType() uint32 {
	if x != nil {
		return x.OrgType
	}
	return 0
}

func (x *Org) GetOrgStatus() uint32 {
	if x != nil {
		return x.OrgStatus
	}
	return 0
}

func (x *Org) GetOrgIndustryId() uint64 {
	if x != nil {
		return x.OrgIndustryId
	}
	return 0
}

func (x *Org) GetOrgScaleId() uint64 {
	if x != nil {
		return x.OrgScaleId
	}
	return 0
}

func (x *Org) GetOrgAddress() string {
	if x != nil {
		return x.OrgAddress
	}
	return ""
}

func (x *Org) GetOrgZipCode() string {
	if x != nil {
		return x.OrgZipCode
	}
	return ""
}

func (x *Org) GetOrgCreatorId() uint64 {
	if x != nil {
		return x.OrgCreatorId
	}
	return 0
}

var File_api_org_service_v1_resources_org_resource_v1_proto protoreflect.FileDescriptor

var file_api_org_service_v1_resources_org_resource_v1_proto_rawDesc = []byte{
	0x0a, 0x32, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6f,
	0x72, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x22, 0x23, 0x0a,
	0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x95, 0x02, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76,
	0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0xbe, 0x04, 0x0a, 0x03, 0x4f, 0x72, 0x67, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x67, 0x5f, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x67,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x72, 0x67, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6f, 0x72, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x72, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x72, 0x67,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x6f, 0x72, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x72, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6f, 0x72, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6f, 0x72, 0x67,
	0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x72,
	0x67, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x6f, 0x72, 0x67, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x72, 0x67, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x72, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a,
	0x0c, 0x6f, 0x72, 0x67, 0x5f, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x67, 0x5a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x24, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x42, 0x88, 0x01, 0x0a, 0x17, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76,
	0x31, 0x42, 0x14, 0x53, 0x61, 0x61, 0x73, 0x41, 0x70, 0x69, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x73,
	0x61, 0x61, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x67,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_org_service_v1_resources_org_resource_v1_proto_rawDescOnce sync.Once
	file_api_org_service_v1_resources_org_resource_v1_proto_rawDescData = file_api_org_service_v1_resources_org_resource_v1_proto_rawDesc
)

func file_api_org_service_v1_resources_org_resource_v1_proto_rawDescGZIP() []byte {
	file_api_org_service_v1_resources_org_resource_v1_proto_rawDescOnce.Do(func() {
		file_api_org_service_v1_resources_org_resource_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_org_service_v1_resources_org_resource_v1_proto_rawDescData)
	})
	return file_api_org_service_v1_resources_org_resource_v1_proto_rawDescData
}

var file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_org_service_v1_resources_org_resource_v1_proto_goTypes = []interface{}{
	(*PingReq)(nil),      // 0: saas.api.org.resourcev1.PingReq
	(*PingResp)(nil),     // 1: saas.api.org.resourcev1.PingResp
	(*PingRespData)(nil), // 2: saas.api.org.resourcev1.PingRespData
	(*Org)(nil),          // 3: saas.api.org.resourcev1.Org
	nil,                  // 4: saas.api.org.resourcev1.PingResp.MetadataEntry
}
var file_api_org_service_v1_resources_org_resource_v1_proto_depIdxs = []int32{
	4, // 0: saas.api.org.resourcev1.PingResp.metadata:type_name -> saas.api.org.resourcev1.PingResp.MetadataEntry
	2, // 1: saas.api.org.resourcev1.PingResp.data:type_name -> saas.api.org.resourcev1.PingRespData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_org_service_v1_resources_org_resource_v1_proto_init() }
func file_api_org_service_v1_resources_org_resource_v1_proto_init() {
	if File_api_org_service_v1_resources_org_resource_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
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
		file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResp); i {
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
		file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRespData); i {
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
		file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Org); i {
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
			RawDescriptor: file_api_org_service_v1_resources_org_resource_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_org_service_v1_resources_org_resource_v1_proto_goTypes,
		DependencyIndexes: file_api_org_service_v1_resources_org_resource_v1_proto_depIdxs,
		MessageInfos:      file_api_org_service_v1_resources_org_resource_v1_proto_msgTypes,
	}.Build()
	File_api_org_service_v1_resources_org_resource_v1_proto = out.File
	file_api_org_service_v1_resources_org_resource_v1_proto_rawDesc = nil
	file_api_org_service_v1_resources_org_resource_v1_proto_goTypes = nil
	file_api_org_service_v1_resources_org_resource_v1_proto_depIdxs = nil
}
