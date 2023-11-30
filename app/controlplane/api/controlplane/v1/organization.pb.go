//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: controlplane/v1/organization.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type OrganizationServiceListMembershipsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrganizationServiceListMembershipsRequest) Reset() {
	*x = OrganizationServiceListMembershipsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_organization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationServiceListMembershipsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationServiceListMembershipsRequest) ProtoMessage() {}

func (x *OrganizationServiceListMembershipsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_organization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationServiceListMembershipsRequest.ProtoReflect.Descriptor instead.
func (*OrganizationServiceListMembershipsRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_organization_proto_rawDescGZIP(), []int{0}
}

type OrganizationServiceListMembershipsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*OrgMembershipItem `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *OrganizationServiceListMembershipsResponse) Reset() {
	*x = OrganizationServiceListMembershipsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_organization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationServiceListMembershipsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationServiceListMembershipsResponse) ProtoMessage() {}

func (x *OrganizationServiceListMembershipsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_organization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationServiceListMembershipsResponse.ProtoReflect.Descriptor instead.
func (*OrganizationServiceListMembershipsResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_organization_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationServiceListMembershipsResponse) GetResult() []*OrgMembershipItem {
	if x != nil {
		return x.Result
	}
	return nil
}

type OrganizationServiceUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// "optional" allow us to detect if the value is explicitly set
	// and not just the default balue
	Name *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *OrganizationServiceUpdateRequest) Reset() {
	*x = OrganizationServiceUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_organization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationServiceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationServiceUpdateRequest) ProtoMessage() {}

func (x *OrganizationServiceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_organization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationServiceUpdateRequest.ProtoReflect.Descriptor instead.
func (*OrganizationServiceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_organization_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationServiceUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrganizationServiceUpdateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type OrganizationServiceUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *OrgItem `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OrganizationServiceUpdateResponse) Reset() {
	*x = OrganizationServiceUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_organization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationServiceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationServiceUpdateResponse) ProtoMessage() {}

func (x *OrganizationServiceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_organization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationServiceUpdateResponse.ProtoReflect.Descriptor instead.
func (*OrganizationServiceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_organization_proto_rawDescGZIP(), []int{3}
}

func (x *OrganizationServiceUpdateResponse) GetResult() *OrgItem {
	if x != nil {
		return x.Result
	}
	return nil
}

type SetCurrentMembershipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MembershipId string `protobuf:"bytes,1,opt,name=membership_id,json=membershipId,proto3" json:"membership_id,omitempty"`
}

func (x *SetCurrentMembershipRequest) Reset() {
	*x = SetCurrentMembershipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_organization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCurrentMembershipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCurrentMembershipRequest) ProtoMessage() {}

func (x *SetCurrentMembershipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_organization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCurrentMembershipRequest.ProtoReflect.Descriptor instead.
func (*SetCurrentMembershipRequest) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_organization_proto_rawDescGZIP(), []int{4}
}

func (x *SetCurrentMembershipRequest) GetMembershipId() string {
	if x != nil {
		return x.MembershipId
	}
	return ""
}

type SetCurrentMembershipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *OrgMembershipItem `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SetCurrentMembershipResponse) Reset() {
	*x = SetCurrentMembershipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controlplane_v1_organization_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCurrentMembershipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCurrentMembershipResponse) ProtoMessage() {}

func (x *SetCurrentMembershipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controlplane_v1_organization_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCurrentMembershipResponse.ProtoReflect.Descriptor instead.
func (*SetCurrentMembershipResponse) Descriptor() ([]byte, []int) {
	return file_controlplane_v1_organization_proto_rawDescGZIP(), []int{5}
}

func (x *SetCurrentMembershipResponse) GetResult() *OrgMembershipItem {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_controlplane_v1_organization_proto protoreflect.FileDescriptor

var file_controlplane_v1_organization_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x29, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x68, 0x0a, 0x2a, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x5e,
	0x0a, 0x20, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x55,
	0x0a, 0x21, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x4c, 0x0a, 0x1b, 0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x1c, 0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0x88, 0x03, 0x0a, 0x13, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x3a, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x31,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x2c, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x6f, 0x6f, 0x70,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controlplane_v1_organization_proto_rawDescOnce sync.Once
	file_controlplane_v1_organization_proto_rawDescData = file_controlplane_v1_organization_proto_rawDesc
)

func file_controlplane_v1_organization_proto_rawDescGZIP() []byte {
	file_controlplane_v1_organization_proto_rawDescOnce.Do(func() {
		file_controlplane_v1_organization_proto_rawDescData = protoimpl.X.CompressGZIP(file_controlplane_v1_organization_proto_rawDescData)
	})
	return file_controlplane_v1_organization_proto_rawDescData
}

var file_controlplane_v1_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_controlplane_v1_organization_proto_goTypes = []interface{}{
	(*OrganizationServiceListMembershipsRequest)(nil),  // 0: controlplane.v1.OrganizationServiceListMembershipsRequest
	(*OrganizationServiceListMembershipsResponse)(nil), // 1: controlplane.v1.OrganizationServiceListMembershipsResponse
	(*OrganizationServiceUpdateRequest)(nil),           // 2: controlplane.v1.OrganizationServiceUpdateRequest
	(*OrganizationServiceUpdateResponse)(nil),          // 3: controlplane.v1.OrganizationServiceUpdateResponse
	(*SetCurrentMembershipRequest)(nil),                // 4: controlplane.v1.SetCurrentMembershipRequest
	(*SetCurrentMembershipResponse)(nil),               // 5: controlplane.v1.SetCurrentMembershipResponse
	(*OrgMembershipItem)(nil),                          // 6: controlplane.v1.OrgMembershipItem
	(*OrgItem)(nil),                                    // 7: controlplane.v1.OrgItem
}
var file_controlplane_v1_organization_proto_depIdxs = []int32{
	6, // 0: controlplane.v1.OrganizationServiceListMembershipsResponse.result:type_name -> controlplane.v1.OrgMembershipItem
	7, // 1: controlplane.v1.OrganizationServiceUpdateResponse.result:type_name -> controlplane.v1.OrgItem
	6, // 2: controlplane.v1.SetCurrentMembershipResponse.result:type_name -> controlplane.v1.OrgMembershipItem
	0, // 3: controlplane.v1.OrganizationService.ListMemberships:input_type -> controlplane.v1.OrganizationServiceListMembershipsRequest
	2, // 4: controlplane.v1.OrganizationService.Update:input_type -> controlplane.v1.OrganizationServiceUpdateRequest
	4, // 5: controlplane.v1.OrganizationService.SetCurrentMembership:input_type -> controlplane.v1.SetCurrentMembershipRequest
	1, // 6: controlplane.v1.OrganizationService.ListMemberships:output_type -> controlplane.v1.OrganizationServiceListMembershipsResponse
	3, // 7: controlplane.v1.OrganizationService.Update:output_type -> controlplane.v1.OrganizationServiceUpdateResponse
	5, // 8: controlplane.v1.OrganizationService.SetCurrentMembership:output_type -> controlplane.v1.SetCurrentMembershipResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_controlplane_v1_organization_proto_init() }
func file_controlplane_v1_organization_proto_init() {
	if File_controlplane_v1_organization_proto != nil {
		return
	}
	file_controlplane_v1_response_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controlplane_v1_organization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationServiceListMembershipsRequest); i {
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
		file_controlplane_v1_organization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationServiceListMembershipsResponse); i {
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
		file_controlplane_v1_organization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationServiceUpdateRequest); i {
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
		file_controlplane_v1_organization_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationServiceUpdateResponse); i {
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
		file_controlplane_v1_organization_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCurrentMembershipRequest); i {
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
		file_controlplane_v1_organization_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCurrentMembershipResponse); i {
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
	file_controlplane_v1_organization_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controlplane_v1_organization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controlplane_v1_organization_proto_goTypes,
		DependencyIndexes: file_controlplane_v1_organization_proto_depIdxs,
		MessageInfos:      file_controlplane_v1_organization_proto_msgTypes,
	}.Build()
	File_controlplane_v1_organization_proto = out.File
	file_controlplane_v1_organization_proto_rawDesc = nil
	file_controlplane_v1_organization_proto_goTypes = nil
	file_controlplane_v1_organization_proto_depIdxs = nil
}
