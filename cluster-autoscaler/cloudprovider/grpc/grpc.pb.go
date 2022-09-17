/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package grpccloudprovider

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type InstanceErrorClass int32

const (
	InstanceErrorClass_ERROR_UNDEFINED        InstanceErrorClass = 0
	InstanceErrorClass_ERROR_OUT_OF_RESOURCES InstanceErrorClass = 1
	InstanceErrorClass_ERROR_OTHER            InstanceErrorClass = 99
)

// Enum value maps for InstanceErrorClass.
var (
	InstanceErrorClass_name = map[int32]string{
		0:  "ERROR_UNDEFINED",
		1:  "ERROR_OUT_OF_RESOURCES",
		99: "ERROR_OTHER",
	}
	InstanceErrorClass_value = map[string]int32{
		"ERROR_UNDEFINED":        0,
		"ERROR_OUT_OF_RESOURCES": 1,
		"ERROR_OTHER":            99,
	}
)

func (x InstanceErrorClass) Enum() *InstanceErrorClass {
	p := new(InstanceErrorClass)
	*p = x
	return p
}

func (x InstanceErrorClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstanceErrorClass) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudprovider_grpc_grpc_proto_enumTypes[0].Descriptor()
}

func (InstanceErrorClass) Type() protoreflect.EnumType {
	return &file_cloudprovider_grpc_grpc_proto_enumTypes[0]
}

func (x InstanceErrorClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstanceErrorClass.Descriptor instead.
func (InstanceErrorClass) EnumDescriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{0}
}

type InstanceState int32

const (
	InstanceState_STATE_UNDEFINED     InstanceState = 0
	InstanceState_STATE_RUNNING       InstanceState = 1
	InstanceState_STATE_BEING_CREATED InstanceState = 2
	InstanceState_STATE_BEING_DELETED InstanceState = 3
)

// Enum value maps for InstanceState.
var (
	InstanceState_name = map[int32]string{
		0: "STATE_UNDEFINED",
		1: "STATE_RUNNING",
		2: "STATE_BEING_CREATED",
		3: "STATE_BEING_DELETED",
	}
	InstanceState_value = map[string]int32{
		"STATE_UNDEFINED":     0,
		"STATE_RUNNING":       1,
		"STATE_BEING_CREATED": 2,
		"STATE_BEING_DELETED": 3,
	}
)

func (x InstanceState) Enum() *InstanceState {
	p := new(InstanceState)
	*p = x
	return p
}

func (x InstanceState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstanceState) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudprovider_grpc_grpc_proto_enumTypes[1].Descriptor()
}

func (InstanceState) Type() protoreflect.EnumType {
	return &file_cloudprovider_grpc_grpc_proto_enumTypes[1]
}

func (x InstanceState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstanceState.Descriptor instead.
func (InstanceState) EnumDescriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{1}
}

type CloudProviderServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderID string `protobuf:"bytes,1,opt,name=providerID,proto3" json:"providerID,omitempty"`
}

func (x *CloudProviderServiceRequest) Reset() {
	*x = CloudProviderServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudProviderServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudProviderServiceRequest) ProtoMessage() {}

func (x *CloudProviderServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudProviderServiceRequest.ProtoReflect.Descriptor instead.
func (*CloudProviderServiceRequest) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *CloudProviderServiceRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

type KubeAdmConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KubeAdmAddress        string   `protobuf:"bytes,1,opt,name=kubeAdmAddress,proto3" json:"kubeAdmAddress,omitempty"`
	KubeAdmToken          string   `protobuf:"bytes,2,opt,name=kubeAdmToken,proto3" json:"kubeAdmToken,omitempty"`
	KubeAdmCACert         string   `protobuf:"bytes,3,opt,name=kubeAdmCACert,proto3" json:"kubeAdmCACert,omitempty"`
	KubeAdmExtraArguments []string `protobuf:"bytes,4,rep,name=kubeAdmExtraArguments,proto3" json:"kubeAdmExtraArguments,omitempty"`
}

func (x *KubeAdmConfig) Reset() {
	*x = KubeAdmConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeAdmConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeAdmConfig) ProtoMessage() {}

func (x *KubeAdmConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeAdmConfig.ProtoReflect.Descriptor instead.
func (*KubeAdmConfig) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *KubeAdmConfig) GetKubeAdmAddress() string {
	if x != nil {
		return x.KubeAdmAddress
	}
	return ""
}

func (x *KubeAdmConfig) GetKubeAdmToken() string {
	if x != nil {
		return x.KubeAdmToken
	}
	return ""
}

func (x *KubeAdmConfig) GetKubeAdmCACert() string {
	if x != nil {
		return x.KubeAdmCACert
	}
	return ""
}

func (x *KubeAdmConfig) GetKubeAdmExtraArguments() []string {
	if x != nil {
		return x.KubeAdmExtraArguments
	}
	return nil
}

type NodeGroupDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeGroupID         string            `protobuf:"bytes,1,opt,name=nodeGroupID,proto3" json:"nodeGroupID,omitempty"`
	MinSize             int32             `protobuf:"varint,2,opt,name=minSize,proto3" json:"minSize,omitempty"`
	MaxSize             int32             `protobuf:"varint,3,opt,name=maxSize,proto3" json:"maxSize,omitempty"`
	Provisionned        bool              `protobuf:"varint,4,opt,name=provisionned,proto3" json:"provisionned,omitempty"`
	IncludeExistingNode bool              `protobuf:"varint,5,opt,name=includeExistingNode,proto3" json:"includeExistingNode,omitempty"`
	Labels              map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NodeGroupDef) Reset() {
	*x = NodeGroupDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeGroupDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGroupDef) ProtoMessage() {}

func (x *NodeGroupDef) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGroupDef.ProtoReflect.Descriptor instead.
func (*NodeGroupDef) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *NodeGroupDef) GetNodeGroupID() string {
	if x != nil {
		return x.NodeGroupID
	}
	return ""
}

func (x *NodeGroupDef) GetMinSize() int32 {
	if x != nil {
		return x.MinSize
	}
	return 0
}

func (x *NodeGroupDef) GetMaxSize() int32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *NodeGroupDef) GetProvisionned() bool {
	if x != nil {
		return x.Provisionned
	}
	return false
}

func (x *NodeGroupDef) GetIncludeExistingNode() bool {
	if x != nil {
		return x.IncludeExistingNode
	}
	return false
}

func (x *NodeGroupDef) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderID           string           `protobuf:"bytes,1,opt,name=providerID,proto3" json:"providerID,omitempty"`
	ResourceLimiter      *ResourceLimiter `protobuf:"bytes,2,opt,name=resourceLimiter,proto3" json:"resourceLimiter,omitempty"`
	KubeAdmConfiguration *KubeAdmConfig   `protobuf:"bytes,3,opt,name=kubeAdmConfiguration,proto3" json:"kubeAdmConfiguration,omitempty"`
	Nodes                []*NodeGroupDef  `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
	AutoProvisionned     bool             `protobuf:"varint,5,opt,name=autoProvisionned,proto3" json:"autoProvisionned,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *ConnectRequest) GetResourceLimiter() *ResourceLimiter {
	if x != nil {
		return x.ResourceLimiter
	}
	return nil
}

func (x *ConnectRequest) GetKubeAdmConfiguration() *KubeAdmConfig {
	if x != nil {
		return x.KubeAdmConfiguration
	}
	return nil
}

func (x *ConnectRequest) GetNodes() []*NodeGroupDef {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *ConnectRequest) GetAutoProvisionned() bool {
	if x != nil {
		return x.AutoProvisionned
	}
	return false
}

type ConnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*ConnectReply_Error
	//	*ConnectReply_Connected
	Response isConnectReply_Response `protobuf_oneof:"response"`
}

func (x *ConnectReply) Reset() {
	*x = ConnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectReply) ProtoMessage() {}

func (x *ConnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectReply.ProtoReflect.Descriptor instead.
func (*ConnectReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{4}
}

func (m *ConnectReply) GetResponse() isConnectReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *ConnectReply) GetError() *Error {
	if x, ok := x.GetResponse().(*ConnectReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *ConnectReply) GetConnected() bool {
	if x, ok := x.GetResponse().(*ConnectReply_Connected); ok {
		return x.Connected
	}
	return false
}

type isConnectReply_Response interface {
	isConnectReply_Response()
}

type ConnectReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type ConnectReply_Connected struct {
	Connected bool `protobuf:"varint,2,opt,name=connected,proto3,oneof"`
}

func (*ConnectReply_Error) isConnectReply_Response() {}

func (*ConnectReply_Connected) isConnectReply_Response() {}

type NameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NameReply) Reset() {
	*x = NameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameReply) ProtoMessage() {}

func (x *NameReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameReply.ProtoReflect.Descriptor instead.
func (*NameReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *NameReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NodeGroupsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeGroups []*NodeGroup `protobuf:"bytes,1,rep,name=nodeGroups,proto3" json:"nodeGroups,omitempty"`
}

func (x *NodeGroupsReply) Reset() {
	*x = NodeGroupsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeGroupsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGroupsReply) ProtoMessage() {}

func (x *NodeGroupsReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGroupsReply.ProtoReflect.Descriptor instead.
func (*NodeGroupsReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *NodeGroupsReply) GetNodeGroups() []*NodeGroup {
	if x != nil {
		return x.NodeGroups
	}
	return nil
}

type NodeGroupForNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderID string `protobuf:"bytes,1,opt,name=providerID,proto3" json:"providerID,omitempty"`
	Node       string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *NodeGroupForNodeRequest) Reset() {
	*x = NodeGroupForNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeGroupForNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGroupForNodeRequest) ProtoMessage() {}

func (x *NodeGroupForNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGroupForNodeRequest.ProtoReflect.Descriptor instead.
func (*NodeGroupForNodeRequest) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *NodeGroupForNodeRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *NodeGroupForNodeRequest) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

type NodeGroupes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*NodeGroup `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *NodeGroupes) Reset() {
	*x = NodeGroupes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeGroupes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGroupes) ProtoMessage() {}

func (x *NodeGroupes) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGroupes.ProtoReflect.Descriptor instead.
func (*NodeGroupes) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *NodeGroupes) GetItems() []*NodeGroup {
	if x != nil {
		return x.Items
	}
	return nil
}

type NodeGroupForNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*NodeGroupForNodeReply_Error
	//	*NodeGroupForNodeReply_NodeGroup
	Response isNodeGroupForNodeReply_Response `protobuf_oneof:"response"`
}

func (x *NodeGroupForNodeReply) Reset() {
	*x = NodeGroupForNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeGroupForNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGroupForNodeReply) ProtoMessage() {}

func (x *NodeGroupForNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGroupForNodeReply.ProtoReflect.Descriptor instead.
func (*NodeGroupForNodeReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{9}
}

func (m *NodeGroupForNodeReply) GetResponse() isNodeGroupForNodeReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *NodeGroupForNodeReply) GetError() *Error {
	if x, ok := x.GetResponse().(*NodeGroupForNodeReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *NodeGroupForNodeReply) GetNodeGroup() *NodeGroup {
	if x, ok := x.GetResponse().(*NodeGroupForNodeReply_NodeGroup); ok {
		return x.NodeGroup
	}
	return nil
}

type isNodeGroupForNodeReply_Response interface {
	isNodeGroupForNodeReply_Response()
}

type NodeGroupForNodeReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type NodeGroupForNodeReply_NodeGroup struct {
	NodeGroup *NodeGroup `protobuf:"bytes,2,opt,name=nodeGroup,proto3,oneof"`
}

func (*NodeGroupForNodeReply_Error) isNodeGroupForNodeReply_Response() {}

func (*NodeGroupForNodeReply_NodeGroup) isNodeGroupForNodeReply_Response() {}

type PricingModelReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*PricingModelReply_Error
	//	*PricingModelReply_PriceModel
	Response isPricingModelReply_Response `protobuf_oneof:"response"`
}

func (x *PricingModelReply) Reset() {
	*x = PricingModelReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricingModelReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricingModelReply) ProtoMessage() {}

func (x *PricingModelReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricingModelReply.ProtoReflect.Descriptor instead.
func (*PricingModelReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{10}
}

func (m *PricingModelReply) GetResponse() isPricingModelReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *PricingModelReply) GetError() *Error {
	if x, ok := x.GetResponse().(*PricingModelReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *PricingModelReply) GetPriceModel() *PricingModel {
	if x, ok := x.GetResponse().(*PricingModelReply_PriceModel); ok {
		return x.PriceModel
	}
	return nil
}

type isPricingModelReply_Response interface {
	isPricingModelReply_Response()
}

type PricingModelReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type PricingModelReply_PriceModel struct {
	PriceModel *PricingModel `protobuf:"bytes,2,opt,name=priceModel,proto3,oneof"`
}

func (*PricingModelReply_Error) isPricingModelReply_Response() {}

func (*PricingModelReply_PriceModel) isPricingModelReply_Response() {}

type AvailableMachineTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineType []string `protobuf:"bytes,1,rep,name=machineType,proto3" json:"machineType,omitempty"`
}

func (x *AvailableMachineTypes) Reset() {
	*x = AvailableMachineTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableMachineTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableMachineTypes) ProtoMessage() {}

func (x *AvailableMachineTypes) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableMachineTypes.ProtoReflect.Descriptor instead.
func (*AvailableMachineTypes) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{11}
}

func (x *AvailableMachineTypes) GetMachineType() []string {
	if x != nil {
		return x.MachineType
	}
	return nil
}

type AvailableMachineTypesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*AvailableMachineTypesReply_Error
	//	*AvailableMachineTypesReply_AvailableMachineTypes
	Response isAvailableMachineTypesReply_Response `protobuf_oneof:"response"`
}

func (x *AvailableMachineTypesReply) Reset() {
	*x = AvailableMachineTypesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableMachineTypesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableMachineTypesReply) ProtoMessage() {}

func (x *AvailableMachineTypesReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableMachineTypesReply.ProtoReflect.Descriptor instead.
func (*AvailableMachineTypesReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{12}
}

func (m *AvailableMachineTypesReply) GetResponse() isAvailableMachineTypesReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *AvailableMachineTypesReply) GetError() *Error {
	if x, ok := x.GetResponse().(*AvailableMachineTypesReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *AvailableMachineTypesReply) GetAvailableMachineTypes() *AvailableMachineTypes {
	if x, ok := x.GetResponse().(*AvailableMachineTypesReply_AvailableMachineTypes); ok {
		return x.AvailableMachineTypes
	}
	return nil
}

type isAvailableMachineTypesReply_Response interface {
	isAvailableMachineTypesReply_Response()
}

type AvailableMachineTypesReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type AvailableMachineTypesReply_AvailableMachineTypes struct {
	AvailableMachineTypes *AvailableMachineTypes `protobuf:"bytes,2,opt,name=availableMachineTypes,proto3,oneof"`
}

func (*AvailableMachineTypesReply_Error) isAvailableMachineTypesReply_Response() {}

func (*AvailableMachineTypesReply_AvailableMachineTypes) isAvailableMachineTypesReply_Response() {}

type NewNodeGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderID     string            `protobuf:"bytes,1,opt,name=providerID,proto3" json:"providerID,omitempty"`
	MachineType    string            `protobuf:"bytes,2,opt,name=machineType,proto3" json:"machineType,omitempty"`
	NodeGroupID    string            `protobuf:"bytes,3,opt,name=nodeGroupID,proto3" json:"nodeGroupID,omitempty"`
	MinNodeSize    int32             `protobuf:"varint,4,opt,name=minNodeSize,proto3" json:"minNodeSize,omitempty"`
	MaxNodeSize    int32             `protobuf:"varint,5,opt,name=maxNodeSize,proto3" json:"maxNodeSize,omitempty"`
	Labels         map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SystemLabels   map[string]string `protobuf:"bytes,7,rep,name=systemLabels,proto3" json:"systemLabels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Taints         []*v1.Taint       `protobuf:"bytes,8,rep,name=taints,proto3" json:"taints,omitempty"`
	ExtraResources map[string]string `protobuf:"bytes,9,rep,name=extraResources,proto3" json:"extraResources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NewNodeGroupRequest) Reset() {
	*x = NewNodeGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewNodeGroupRequest) ProtoMessage() {}

func (x *NewNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*NewNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{13}
}

func (x *NewNodeGroupRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *NewNodeGroupRequest) GetMachineType() string {
	if x != nil {
		return x.MachineType
	}
	return ""
}

func (x *NewNodeGroupRequest) GetNodeGroupID() string {
	if x != nil {
		return x.NodeGroupID
	}
	return ""
}

func (x *NewNodeGroupRequest) GetMinNodeSize() int32 {
	if x != nil {
		return x.MinNodeSize
	}
	return 0
}

func (x *NewNodeGroupRequest) GetMaxNodeSize() int32 {
	if x != nil {
		return x.MaxNodeSize
	}
	return 0
}

func (x *NewNodeGroupRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *NewNodeGroupRequest) GetSystemLabels() map[string]string {
	if x != nil {
		return x.SystemLabels
	}
	return nil
}

func (x *NewNodeGroupRequest) GetTaints() []*v1.Taint {
	if x != nil {
		return x.Taints
	}
	return nil
}

func (x *NewNodeGroupRequest) GetExtraResources() map[string]string {
	if x != nil {
		return x.ExtraResources
	}
	return nil
}

type NewNodeGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*NewNodeGroupReply_Error
	//	*NewNodeGroupReply_NodeGroup
	Response isNewNodeGroupReply_Response `protobuf_oneof:"response"`
}

func (x *NewNodeGroupReply) Reset() {
	*x = NewNodeGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewNodeGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewNodeGroupReply) ProtoMessage() {}

func (x *NewNodeGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewNodeGroupReply.ProtoReflect.Descriptor instead.
func (*NewNodeGroupReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{14}
}

func (m *NewNodeGroupReply) GetResponse() isNewNodeGroupReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *NewNodeGroupReply) GetError() *Error {
	if x, ok := x.GetResponse().(*NewNodeGroupReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *NewNodeGroupReply) GetNodeGroup() *NodeGroup {
	if x, ok := x.GetResponse().(*NewNodeGroupReply_NodeGroup); ok {
		return x.NodeGroup
	}
	return nil
}

type isNewNodeGroupReply_Response interface {
	isNewNodeGroupReply_Response()
}

type NewNodeGroupReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type NewNodeGroupReply_NodeGroup struct {
	NodeGroup *NodeGroup `protobuf:"bytes,2,opt,name=nodeGroup,proto3,oneof"`
}

func (*NewNodeGroupReply_Error) isNewNodeGroupReply_Response() {}

func (*NewNodeGroupReply_NodeGroup) isNewNodeGroupReply_Response() {}

type ResourceLimiterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*ResourceLimiterReply_Error
	//	*ResourceLimiterReply_ResourceLimiter
	Response isResourceLimiterReply_Response `protobuf_oneof:"response"`
}

func (x *ResourceLimiterReply) Reset() {
	*x = ResourceLimiterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceLimiterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceLimiterReply) ProtoMessage() {}

func (x *ResourceLimiterReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceLimiterReply.ProtoReflect.Descriptor instead.
func (*ResourceLimiterReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{15}
}

func (m *ResourceLimiterReply) GetResponse() isResourceLimiterReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *ResourceLimiterReply) GetError() *Error {
	if x, ok := x.GetResponse().(*ResourceLimiterReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *ResourceLimiterReply) GetResourceLimiter() *ResourceLimiter {
	if x, ok := x.GetResponse().(*ResourceLimiterReply_ResourceLimiter); ok {
		return x.ResourceLimiter
	}
	return nil
}

type isResourceLimiterReply_Response interface {
	isResourceLimiterReply_Response()
}

type ResourceLimiterReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type ResourceLimiterReply_ResourceLimiter struct {
	ResourceLimiter *ResourceLimiter `protobuf:"bytes,2,opt,name=resourceLimiter,proto3,oneof"`
}

func (*ResourceLimiterReply_Error) isResourceLimiterReply_Response() {}

func (*ResourceLimiterReply_ResourceLimiter) isResourceLimiterReply_Response() {}

type GPULabelReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*GPULabelReply_Error
	//	*GPULabelReply_Gpulabel
	Response isGPULabelReply_Response `protobuf_oneof:"response"`
}

func (x *GPULabelReply) Reset() {
	*x = GPULabelReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GPULabelReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPULabelReply) ProtoMessage() {}

func (x *GPULabelReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPULabelReply.ProtoReflect.Descriptor instead.
func (*GPULabelReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{16}
}

func (m *GPULabelReply) GetResponse() isGPULabelReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *GPULabelReply) GetError() *Error {
	if x, ok := x.GetResponse().(*GPULabelReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *GPULabelReply) GetGpulabel() string {
	if x, ok := x.GetResponse().(*GPULabelReply_Gpulabel); ok {
		return x.Gpulabel
	}
	return ""
}

type isGPULabelReply_Response interface {
	isGPULabelReply_Response()
}

type GPULabelReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type GPULabelReply_Gpulabel struct {
	Gpulabel string `protobuf:"bytes,2,opt,name=gpulabel,proto3,oneof"`
}

func (*GPULabelReply_Error) isGPULabelReply_Response() {}

func (*GPULabelReply_Gpulabel) isGPULabelReply_Response() {}

type GetAvailableGPUTypesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailableGpuTypes map[string]string `protobuf:"bytes,2,rep,name=availableGpuTypes,proto3" json:"availableGpuTypes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetAvailableGPUTypesReply) Reset() {
	*x = GetAvailableGPUTypesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailableGPUTypesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableGPUTypesReply) ProtoMessage() {}

func (x *GetAvailableGPUTypesReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableGPUTypesReply.ProtoReflect.Descriptor instead.
func (*GetAvailableGPUTypesReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{17}
}

func (x *GetAvailableGPUTypesReply) GetAvailableGpuTypes() map[string]string {
	if x != nil {
		return x.AvailableGpuTypes
	}
	return nil
}

type CleanupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CleanupReply) Reset() {
	*x = CleanupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CleanupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CleanupReply) ProtoMessage() {}

func (x *CleanupReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CleanupReply.ProtoReflect.Descriptor instead.
func (*CleanupReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{18}
}

func (x *CleanupReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type RefreshReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RefreshReply) Reset() {
	*x = RefreshReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshReply) ProtoMessage() {}

func (x *RefreshReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshReply.ProtoReflect.Descriptor instead.
func (*RefreshReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{19}
}

func (x *RefreshReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type BelongsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderID  string `protobuf:"bytes,1,opt,name=providerID,proto3" json:"providerID,omitempty"`
	NodeGroupID string `protobuf:"bytes,2,opt,name=nodeGroupID,proto3" json:"nodeGroupID,omitempty"`
	Node        string `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *BelongsRequest) Reset() {
	*x = BelongsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BelongsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BelongsRequest) ProtoMessage() {}

func (x *BelongsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BelongsRequest.ProtoReflect.Descriptor instead.
func (*BelongsRequest) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{20}
}

func (x *BelongsRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *BelongsRequest) GetNodeGroupID() string {
	if x != nil {
		return x.NodeGroupID
	}
	return ""
}

func (x *BelongsRequest) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

type BelongsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*BelongsReply_Error
	//	*BelongsReply_Belongs
	Response isBelongsReply_Response `protobuf_oneof:"response"`
}

func (x *BelongsReply) Reset() {
	*x = BelongsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BelongsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BelongsReply) ProtoMessage() {}

func (x *BelongsReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BelongsReply.ProtoReflect.Descriptor instead.
func (*BelongsReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{21}
}

func (m *BelongsReply) GetResponse() isBelongsReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *BelongsReply) GetError() *Error {
	if x, ok := x.GetResponse().(*BelongsReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *BelongsReply) GetBelongs() bool {
	if x, ok := x.GetResponse().(*BelongsReply_Belongs); ok {
		return x.Belongs
	}
	return false
}

type isBelongsReply_Response interface {
	isBelongsReply_Response()
}

type BelongsReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type BelongsReply_Belongs struct {
	Belongs bool `protobuf:"varint,2,opt,name=belongs,proto3,oneof"`
}

func (*BelongsReply_Error) isBelongsReply_Response() {}

func (*BelongsReply_Belongs) isBelongsReply_Response() {}

type NodeGroupServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderID  string `protobuf:"bytes,1,opt,name=providerID,proto3" json:"providerID,omitempty"`
	NodeGroupID string `protobuf:"bytes,2,opt,name=nodeGroupID,proto3" json:"nodeGroupID,omitempty"`
}

func (x *NodeGroupServiceRequest) Reset() {
	*x = NodeGroupServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeGroupServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGroupServiceRequest) ProtoMessage() {}

func (x *NodeGroupServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGroupServiceRequest.ProtoReflect.Descriptor instead.
func (*NodeGroupServiceRequest) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{22}
}

func (x *NodeGroupServiceRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *NodeGroupServiceRequest) GetNodeGroupID() string {
	if x != nil {
		return x.NodeGroupID
	}
	return ""
}

type MaxSizeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxSize int32 `protobuf:"varint,1,opt,name=maxSize,proto3" json:"maxSize,omitempty"`
}

func (x *MaxSizeReply) Reset() {
	*x = MaxSizeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxSizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxSizeReply) ProtoMessage() {}

func (x *MaxSizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxSizeReply.ProtoReflect.Descriptor instead.
func (*MaxSizeReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{23}
}

func (x *MaxSizeReply) GetMaxSize() int32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

type MinSizeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinSize int32 `protobuf:"varint,1,opt,name=minSize,proto3" json:"minSize,omitempty"`
}

func (x *MinSizeReply) Reset() {
	*x = MinSizeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinSizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinSizeReply) ProtoMessage() {}

func (x *MinSizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinSizeReply.ProtoReflect.Descriptor instead.
func (*MinSizeReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{24}
}

func (x *MinSizeReply) GetMinSize() int32 {
	if x != nil {
		return x.MinSize
	}
	return 0
}

type TargetSizeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*TargetSizeReply_Error
	//	*TargetSizeReply_TargetSize
	Response isTargetSizeReply_Response `protobuf_oneof:"response"`
}

func (x *TargetSizeReply) Reset() {
	*x = TargetSizeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetSizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetSizeReply) ProtoMessage() {}

func (x *TargetSizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetSizeReply.ProtoReflect.Descriptor instead.
func (*TargetSizeReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{25}
}

func (m *TargetSizeReply) GetResponse() isTargetSizeReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *TargetSizeReply) GetError() *Error {
	if x, ok := x.GetResponse().(*TargetSizeReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *TargetSizeReply) GetTargetSize() int32 {
	if x, ok := x.GetResponse().(*TargetSizeReply_TargetSize); ok {
		return x.TargetSize
	}
	return 0
}

type isTargetSizeReply_Response interface {
	isTargetSizeReply_Response()
}

type TargetSizeReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type TargetSizeReply_TargetSize struct {
	TargetSize int32 `protobuf:"varint,2,opt,name=targetSize,proto3,oneof"`
}

func (*TargetSizeReply_Error) isTargetSizeReply_Response() {}

func (*TargetSizeReply_TargetSize) isTargetSizeReply_Response() {}

type IncreaseSizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderID  string `protobuf:"bytes,1,opt,name=providerID,proto3" json:"providerID,omitempty"`
	NodeGroupID string `protobuf:"bytes,2,opt,name=nodeGroupID,proto3" json:"nodeGroupID,omitempty"`
	Delta       int32  `protobuf:"varint,3,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *IncreaseSizeRequest) Reset() {
	*x = IncreaseSizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[26]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseSizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseSizeRequest) ProtoMessage() {}

func (x *IncreaseSizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[26]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseSizeRequest.ProtoReflect.Descriptor instead.
func (*IncreaseSizeRequest) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{26}
}

func (x *IncreaseSizeRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *IncreaseSizeRequest) GetNodeGroupID() string {
	if x != nil {
		return x.NodeGroupID
	}
	return ""
}

func (x *IncreaseSizeRequest) GetDelta() int32 {
	if x != nil {
		return x.Delta
	}
	return 0
}

type IncreaseSizeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *IncreaseSizeReply) Reset() {
	*x = IncreaseSizeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[27]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseSizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseSizeReply) ProtoMessage() {}

func (x *IncreaseSizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[27]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseSizeReply.ProtoReflect.Descriptor instead.
func (*IncreaseSizeReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{27}
}

func (x *IncreaseSizeReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type DeleteNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderID  string   `protobuf:"bytes,1,opt,name=providerID,proto3" json:"providerID,omitempty"`
	NodeGroupID string   `protobuf:"bytes,2,opt,name=nodeGroupID,proto3" json:"nodeGroupID,omitempty"`
	Node        []string `protobuf:"bytes,3,rep,name=node,proto3" json:"node,omitempty"`
}

func (x *DeleteNodesRequest) Reset() {
	*x = DeleteNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[28]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodesRequest) ProtoMessage() {}

func (x *DeleteNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[28]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodesRequest.ProtoReflect.Descriptor instead.
func (*DeleteNodesRequest) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{28}
}

func (x *DeleteNodesRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *DeleteNodesRequest) GetNodeGroupID() string {
	if x != nil {
		return x.NodeGroupID
	}
	return ""
}

func (x *DeleteNodesRequest) GetNode() []string {
	if x != nil {
		return x.Node
	}
	return nil
}

type DeleteNodesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteNodesReply) Reset() {
	*x = DeleteNodesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[29]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodesReply) ProtoMessage() {}

func (x *DeleteNodesReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[29]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodesReply.ProtoReflect.Descriptor instead.
func (*DeleteNodesReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{29}
}

func (x *DeleteNodesReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type DecreaseTargetSizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderID  string `protobuf:"bytes,1,opt,name=providerID,proto3" json:"providerID,omitempty"`
	NodeGroupID string `protobuf:"bytes,2,opt,name=nodeGroupID,proto3" json:"nodeGroupID,omitempty"`
	Delta       int32  `protobuf:"varint,3,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *DecreaseTargetSizeRequest) Reset() {
	*x = DecreaseTargetSizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[30]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecreaseTargetSizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseTargetSizeRequest) ProtoMessage() {}

func (x *DecreaseTargetSizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[30]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseTargetSizeRequest.ProtoReflect.Descriptor instead.
func (*DecreaseTargetSizeRequest) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{30}
}

func (x *DecreaseTargetSizeRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *DecreaseTargetSizeRequest) GetNodeGroupID() string {
	if x != nil {
		return x.NodeGroupID
	}
	return ""
}

func (x *DecreaseTargetSizeRequest) GetDelta() int32 {
	if x != nil {
		return x.Delta
	}
	return 0
}

type DecreaseTargetSizeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DecreaseTargetSizeReply) Reset() {
	*x = DecreaseTargetSizeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[31]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecreaseTargetSizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseTargetSizeReply) ProtoMessage() {}

func (x *DecreaseTargetSizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[31]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseTargetSizeReply.ProtoReflect.Descriptor instead.
func (*DecreaseTargetSizeReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{31}
}

func (x *DecreaseTargetSizeReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type IdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *IdReply) Reset() {
	*x = IdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[32]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReply) ProtoMessage() {}

func (x *IdReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[32]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReply.ProtoReflect.Descriptor instead.
func (*IdReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{32}
}

func (x *IdReply) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type DebugReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DebugReply) Reset() {
	*x = DebugReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[33]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugReply) ProtoMessage() {}

func (x *DebugReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[33]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugReply.ProtoReflect.Descriptor instead.
func (*DebugReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{33}
}

func (x *DebugReply) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type Instances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Instance `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Instances) Reset() {
	*x = Instances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[34]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instances) ProtoMessage() {}

func (x *Instances) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[34]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instances.ProtoReflect.Descriptor instead.
func (*Instances) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{34}
}

func (x *Instances) GetItems() []*Instance {
	if x != nil {
		return x.Items
	}
	return nil
}

type NodesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*NodesReply_Error
	//	*NodesReply_Instances
	Response isNodesReply_Response `protobuf_oneof:"response"`
}

func (x *NodesReply) Reset() {
	*x = NodesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[35]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodesReply) ProtoMessage() {}

func (x *NodesReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[35]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodesReply.ProtoReflect.Descriptor instead.
func (*NodesReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{35}
}

func (m *NodesReply) GetResponse() isNodesReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *NodesReply) GetError() *Error {
	if x, ok := x.GetResponse().(*NodesReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *NodesReply) GetInstances() *Instances {
	if x, ok := x.GetResponse().(*NodesReply_Instances); ok {
		return x.Instances
	}
	return nil
}

type isNodesReply_Response interface {
	isNodesReply_Response()
}

type NodesReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type NodesReply_Instances struct {
	Instances *Instances `protobuf:"bytes,2,opt,name=instances,proto3,oneof"`
}

func (*NodesReply_Error) isNodesReply_Response() {}

func (*NodesReply_Instances) isNodesReply_Response() {}

type ImageStateSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size     int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	NumNodes int32 `protobuf:"varint,2,opt,name=numNodes,proto3" json:"numNodes,omitempty"`
}

func (x *ImageStateSummary) Reset() {
	*x = ImageStateSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[36]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageStateSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageStateSummary) ProtoMessage() {}

func (x *ImageStateSummary) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[36]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageStateSummary.ProtoReflect.Descriptor instead.
func (*ImageStateSummary) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{36}
}

func (x *ImageStateSummary) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ImageStateSummary) GetNumNodes() int32 {
	if x != nil {
		return x.NumNodes
	}
	return 0
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MilliCPU         int64 `protobuf:"varint,1,opt,name=milliCPU,proto3" json:"milliCPU,omitempty"`
	Memory           int64 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	EphemeralStorage int64 `protobuf:"varint,3,opt,name=ephemeralStorage,proto3" json:"ephemeralStorage,omitempty"`
	// We store allowedPodNumber (which is Node.Status.Allocatable.Pods().Value())
	// explicitly as int, to avoid conversions and improve performance.
	AllowedPodNumber int32 `protobuf:"varint,4,opt,name=allowedPodNumber,proto3" json:"allowedPodNumber,omitempty"`
	// ScalarResources
	ScalarResources map[string]int64 `protobuf:"bytes,5,rep,name=scalarResources,proto3" json:"scalarResources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[37]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[37]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{37}
}

func (x *Resource) GetMilliCPU() int64 {
	if x != nil {
		return x.MilliCPU
	}
	return 0
}

func (x *Resource) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Resource) GetEphemeralStorage() int64 {
	if x != nil {
		return x.EphemeralStorage
	}
	return 0
}

func (x *Resource) GetAllowedPodNumber() int32 {
	if x != nil {
		return x.AllowedPodNumber
	}
	return 0
}

func (x *Resource) GetScalarResources() map[string]int64 {
	if x != nil {
		return x.ScalarResources
	}
	return nil
}

type TransientSchedulerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllocatableVolumesCount int32 `protobuf:"varint,1,opt,name=allocatableVolumesCount,proto3" json:"allocatableVolumesCount,omitempty"`
	RequestedVolumes        int32 `protobuf:"varint,2,opt,name=requestedVolumes,proto3" json:"requestedVolumes,omitempty"`
}

func (x *TransientSchedulerInfo) Reset() {
	*x = TransientSchedulerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[38]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransientSchedulerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransientSchedulerInfo) ProtoMessage() {}

func (x *TransientSchedulerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[38]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransientSchedulerInfo.ProtoReflect.Descriptor instead.
func (*TransientSchedulerInfo) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{38}
}

func (x *TransientSchedulerInfo) GetAllocatableVolumesCount() int32 {
	if x != nil {
		return x.AllocatableVolumesCount
	}
	return 0
}

func (x *TransientSchedulerInfo) GetRequestedVolumes() int32 {
	if x != nil {
		return x.RequestedVolumes
	}
	return 0
}

type ProtocolPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ProtocolPort) Reset() {
	*x = ProtocolPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[39]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolPort) ProtoMessage() {}

func (x *ProtocolPort) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[39]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolPort.ProtoReflect.Descriptor instead.
func (*ProtocolPort) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{39}
}

func (x *ProtocolPort) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ProtocolPort) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type HostPortInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostPortInfo map[string]*ProtocolPort `protobuf:"bytes,1,rep,name=hostPortInfo,proto3" json:"hostPortInfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HostPortInfo) Reset() {
	*x = HostPortInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[40]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostPortInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostPortInfo) ProtoMessage() {}

func (x *HostPortInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[40]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostPortInfo.ProtoReflect.Descriptor instead.
func (*HostPortInfo) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{40}
}

func (x *HostPortInfo) GetHostPortInfo() map[string]*ProtocolPort {
	if x != nil {
		return x.HostPortInfo
	}
	return nil
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node              string                        `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Pods              []string                      `protobuf:"bytes,2,rep,name=pods,proto3" json:"pods,omitempty"`
	UsedPorts         *HostPortInfo                 `protobuf:"bytes,3,opt,name=usedPorts,proto3" json:"usedPorts,omitempty"`
	RequestedResource *Resource                     `protobuf:"bytes,4,opt,name=requestedResource,proto3" json:"requestedResource,omitempty"`
	NonzeroRequest    *Resource                     `protobuf:"bytes,5,opt,name=nonzeroRequest,proto3" json:"nonzeroRequest,omitempty"`
	ImageStates       map[string]*ImageStateSummary `protobuf:"bytes,6,rep,name=imageStates,proto3" json:"imageStates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TransientInfo     *TransientSchedulerInfo       `protobuf:"bytes,7,opt,name=TransientInfo,proto3" json:"TransientInfo,omitempty"`
	Generation        int64                         `protobuf:"varint,8,opt,name=generation,proto3" json:"generation,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[41]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[41]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{41}
}

func (x *NodeInfo) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *NodeInfo) GetPods() []string {
	if x != nil {
		return x.Pods
	}
	return nil
}

func (x *NodeInfo) GetUsedPorts() *HostPortInfo {
	if x != nil {
		return x.UsedPorts
	}
	return nil
}

func (x *NodeInfo) GetRequestedResource() *Resource {
	if x != nil {
		return x.RequestedResource
	}
	return nil
}

func (x *NodeInfo) GetNonzeroRequest() *Resource {
	if x != nil {
		return x.NonzeroRequest
	}
	return nil
}

func (x *NodeInfo) GetImageStates() map[string]*ImageStateSummary {
	if x != nil {
		return x.ImageStates
	}
	return nil
}

func (x *NodeInfo) GetTransientInfo() *TransientSchedulerInfo {
	if x != nil {
		return x.TransientInfo
	}
	return nil
}

func (x *NodeInfo) GetGeneration() int64 {
	if x != nil {
		return x.Generation
	}
	return 0
}

type TemplateNodeInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*TemplateNodeInfoReply_Error
	//	*TemplateNodeInfoReply_NodeInfo
	Response isTemplateNodeInfoReply_Response `protobuf_oneof:"response"`
}

func (x *TemplateNodeInfoReply) Reset() {
	*x = TemplateNodeInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[42]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateNodeInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateNodeInfoReply) ProtoMessage() {}

func (x *TemplateNodeInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[42]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateNodeInfoReply.ProtoReflect.Descriptor instead.
func (*TemplateNodeInfoReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{42}
}

func (m *TemplateNodeInfoReply) GetResponse() isTemplateNodeInfoReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *TemplateNodeInfoReply) GetError() *Error {
	if x, ok := x.GetResponse().(*TemplateNodeInfoReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *TemplateNodeInfoReply) GetNodeInfo() *NodeInfo {
	if x, ok := x.GetResponse().(*TemplateNodeInfoReply_NodeInfo); ok {
		return x.NodeInfo
	}
	return nil
}

type isTemplateNodeInfoReply_Response interface {
	isTemplateNodeInfoReply_Response()
}

type TemplateNodeInfoReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type TemplateNodeInfoReply_NodeInfo struct {
	NodeInfo *NodeInfo `protobuf:"bytes,2,opt,name=nodeInfo,proto3,oneof"`
}

func (*TemplateNodeInfoReply_Error) isTemplateNodeInfoReply_Response() {}

func (*TemplateNodeInfoReply_NodeInfo) isTemplateNodeInfoReply_Response() {}

type ExistReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *ExistReply) Reset() {
	*x = ExistReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[43]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistReply) ProtoMessage() {}

func (x *ExistReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[43]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistReply.ProtoReflect.Descriptor instead.
func (*ExistReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{43}
}

func (x *ExistReply) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type CreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*CreateReply_Error
	//	*CreateReply_NodeGroup
	Response isCreateReply_Response `protobuf_oneof:"response"`
}

func (x *CreateReply) Reset() {
	*x = CreateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[44]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReply) ProtoMessage() {}

func (x *CreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[44]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReply.ProtoReflect.Descriptor instead.
func (*CreateReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{44}
}

func (m *CreateReply) GetResponse() isCreateReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *CreateReply) GetError() *Error {
	if x, ok := x.GetResponse().(*CreateReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *CreateReply) GetNodeGroup() *NodeGroup {
	if x, ok := x.GetResponse().(*CreateReply_NodeGroup); ok {
		return x.NodeGroup
	}
	return nil
}

type isCreateReply_Response interface {
	isCreateReply_Response()
}

type CreateReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type CreateReply_NodeGroup struct {
	NodeGroup *NodeGroup `protobuf:"bytes,2,opt,name=nodeGroup,proto3,oneof"`
}

func (*CreateReply_Error) isCreateReply_Response() {}

func (*CreateReply_NodeGroup) isCreateReply_Response() {}

type DeleteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteReply) Reset() {
	*x = DeleteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[45]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReply) ProtoMessage() {}

func (x *DeleteReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[45]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReply.ProtoReflect.Descriptor instead.
func (*DeleteReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{45}
}

func (x *DeleteReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type AutoprovisionedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Autoprovisioned bool `protobuf:"varint,1,opt,name=autoprovisioned,proto3" json:"autoprovisioned,omitempty"`
}

func (x *AutoprovisionedReply) Reset() {
	*x = AutoprovisionedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[46]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoprovisionedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoprovisionedReply) ProtoMessage() {}

func (x *AutoprovisionedReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[46]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoprovisionedReply.ProtoReflect.Descriptor instead.
func (*AutoprovisionedReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{46}
}

func (x *AutoprovisionedReply) GetAutoprovisioned() bool {
	if x != nil {
		return x.Autoprovisioned
	}
	return false
}

type NodePriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderID string `protobuf:"bytes,1,opt,name=providerID,proto3" json:"providerID,omitempty"`
	Node       string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	StartTime  int64  `protobuf:"varint,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime    int64  `protobuf:"varint,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *NodePriceRequest) Reset() {
	*x = NodePriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[47]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodePriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodePriceRequest) ProtoMessage() {}

func (x *NodePriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[47]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodePriceRequest.ProtoReflect.Descriptor instead.
func (*NodePriceRequest) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{47}
}

func (x *NodePriceRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *NodePriceRequest) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *NodePriceRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *NodePriceRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type NodePriceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*NodePriceReply_Error
	//	*NodePriceReply_Price
	Response isNodePriceReply_Response `protobuf_oneof:"response"`
}

func (x *NodePriceReply) Reset() {
	*x = NodePriceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[48]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodePriceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodePriceReply) ProtoMessage() {}

func (x *NodePriceReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[48]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodePriceReply.ProtoReflect.Descriptor instead.
func (*NodePriceReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{48}
}

func (m *NodePriceReply) GetResponse() isNodePriceReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *NodePriceReply) GetError() *Error {
	if x, ok := x.GetResponse().(*NodePriceReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *NodePriceReply) GetPrice() float64 {
	if x, ok := x.GetResponse().(*NodePriceReply_Price); ok {
		return x.Price
	}
	return 0
}

type isNodePriceReply_Response interface {
	isNodePriceReply_Response()
}

type NodePriceReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type NodePriceReply_Price struct {
	Price float64 `protobuf:"fixed64,2,opt,name=price,proto3,oneof"`
}

func (*NodePriceReply_Error) isNodePriceReply_Response() {}

func (*NodePriceReply_Price) isNodePriceReply_Response() {}

type PodPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderID string `protobuf:"bytes,1,opt,name=providerID,proto3" json:"providerID,omitempty"`
	Pod        string `protobuf:"bytes,2,opt,name=pod,proto3" json:"pod,omitempty"`
	StartTime  int64  `protobuf:"varint,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime    int64  `protobuf:"varint,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *PodPriceRequest) Reset() {
	*x = PodPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[49]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodPriceRequest) ProtoMessage() {}

func (x *PodPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[49]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodPriceRequest.ProtoReflect.Descriptor instead.
func (*PodPriceRequest) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{49}
}

func (x *PodPriceRequest) GetProviderID() string {
	if x != nil {
		return x.ProviderID
	}
	return ""
}

func (x *PodPriceRequest) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

func (x *PodPriceRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *PodPriceRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type PodPriceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*PodPriceReply_Error
	//	*PodPriceReply_Price
	Response isPodPriceReply_Response `protobuf_oneof:"response"`
}

func (x *PodPriceReply) Reset() {
	*x = PodPriceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[50]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodPriceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodPriceReply) ProtoMessage() {}

func (x *PodPriceReply) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[50]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodPriceReply.ProtoReflect.Descriptor instead.
func (*PodPriceReply) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{50}
}

func (m *PodPriceReply) GetResponse() isPodPriceReply_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *PodPriceReply) GetError() *Error {
	if x, ok := x.GetResponse().(*PodPriceReply_Error); ok {
		return x.Error
	}
	return nil
}

func (x *PodPriceReply) GetPrice() float64 {
	if x, ok := x.GetResponse().(*PodPriceReply_Price); ok {
		return x.Price
	}
	return 0
}

type isPodPriceReply_Response interface {
	isPodPriceReply_Response()
}

type PodPriceReply_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type PodPriceReply_Price struct {
	Price float64 `protobuf:"fixed64,2,opt,name=price,proto3,oneof"`
}

func (*PodPriceReply_Error) isPodPriceReply_Response() {}

func (*PodPriceReply_Price) isPodPriceReply_Response() {}

/////////////////////////////////////////////////////////////////////////////////////////////////////////
// General messages types
/////////////////////////////////////////////////////////////////////////////////////////////////////////
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[51]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[51]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{51}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type NodeGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NodeGroup) Reset() {
	*x = NodeGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[52]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGroup) ProtoMessage() {}

func (x *NodeGroup) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[52]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGroup.ProtoReflect.Descriptor instead.
func (*NodeGroup) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{52}
}

func (x *NodeGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PricingModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PricingModel) Reset() {
	*x = PricingModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[53]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricingModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricingModel) ProtoMessage() {}

func (x *PricingModel) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[53]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricingModel.ProtoReflect.Descriptor instead.
func (*PricingModel) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{53}
}

func (x *PricingModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResourceLimiter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinLimits map[string]int64 `protobuf:"bytes,1,rep,name=minLimits,proto3" json:"minLimits,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MaxLimits map[string]int64 `protobuf:"bytes,2,rep,name=maxLimits,proto3" json:"maxLimits,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ResourceLimiter) Reset() {
	*x = ResourceLimiter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[54]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceLimiter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceLimiter) ProtoMessage() {}

func (x *ResourceLimiter) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[54]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceLimiter.ProtoReflect.Descriptor instead.
func (*ResourceLimiter) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{54}
}

func (x *ResourceLimiter) GetMinLimits() map[string]int64 {
	if x != nil {
		return x.MinLimits
	}
	return nil
}

func (x *ResourceLimiter) GetMaxLimits() map[string]int64 {
	if x != nil {
		return x.MaxLimits
	}
	return nil
}

type InstanceErrorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorClass InstanceErrorClass `protobuf:"varint,1,opt,name=errorClass,proto3,enum=grpccloudprovider.InstanceErrorClass" json:"errorClass,omitempty"`
	// ErrorCode is cloud-provider specific error code for error condition
	ErrorCode string `protobuf:"bytes,2,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	// ErrorMessage is human readable description of error condition
	ErrorMessage string `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *InstanceErrorInfo) Reset() {
	*x = InstanceErrorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[55]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceErrorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceErrorInfo) ProtoMessage() {}

func (x *InstanceErrorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[55]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceErrorInfo.ProtoReflect.Descriptor instead.
func (*InstanceErrorInfo) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{55}
}

func (x *InstanceErrorInfo) GetErrorClass() InstanceErrorClass {
	if x != nil {
		return x.ErrorClass
	}
	return InstanceErrorClass_ERROR_UNDEFINED
}

func (x *InstanceErrorInfo) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *InstanceErrorInfo) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type InstanceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State     InstanceState      `protobuf:"varint,1,opt,name=state,proto3,enum=grpccloudprovider.InstanceState" json:"state,omitempty"`
	ErrorInfo *InstanceErrorInfo `protobuf:"bytes,2,opt,name=errorInfo,proto3" json:"errorInfo,omitempty"`
}

func (x *InstanceStatus) Reset() {
	*x = InstanceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[56]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceStatus) ProtoMessage() {}

func (x *InstanceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[56]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceStatus.ProtoReflect.Descriptor instead.
func (*InstanceStatus) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{56}
}

func (x *InstanceStatus) GetState() InstanceState {
	if x != nil {
		return x.State
	}
	return InstanceState_STATE_UNDEFINED
}

func (x *InstanceStatus) GetErrorInfo() *InstanceErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status *InstanceStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[57]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_grpc_grpc_proto_msgTypes[57]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_cloudprovider_grpc_grpc_proto_rawDescGZIP(), []int{57}
}

func (x *Instance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Instance) GetStatus() *InstanceStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_cloudprovider_grpc_grpc_proto protoreflect.FileDescriptor

var file_cloudprovider_grpc_grpc_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x1a, 0x22, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x1b, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0xb7, 0x01, 0x0a, 0x0d, 0x4b, 0x75, 0x62, 0x65, 0x41, 0x64,
	0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x41,
	0x64, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6b, 0x75, 0x62, 0x65, 0x41, 0x64, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x6b, 0x75, 0x62, 0x65, 0x41, 0x64, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x75, 0x62, 0x65, 0x41, 0x64, 0x6d, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x6b, 0x75, 0x62, 0x65, 0x41, 0x64, 0x6d, 0x43, 0x41,
	0x43, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x75, 0x62, 0x65,
	0x41, 0x64, 0x6d, 0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x6b, 0x75, 0x62,
	0x65, 0x41, 0x64, 0x6d, 0x45, 0x78, 0x74, 0x72, 0x61, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x6b, 0x75, 0x62, 0x65, 0x41, 0x64,
	0x6d, 0x45, 0x78, 0x74, 0x72, 0x61, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0xba, 0x02, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x66,
	0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x6e, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x43, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x66, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb7, 0x02, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x4c, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x52, 0x0f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x54, 0x0a,
	0x14, 0x6b, 0x75, 0x62, 0x65, 0x41, 0x64, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x41, 0x64, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x6b,
	0x75, 0x62, 0x65, 0x41, 0x64, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x44, 0x65, 0x66, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x75,
	0x74, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x6e, 0x65, 0x64, 0x22, 0x6c, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48,
	0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x4d, 0x0a, 0x17, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x46, 0x6f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x41, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x15, 0x4e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x6f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x94,
	0x01, 0x0a, 0x11, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x0a, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x15, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x22, 0xbc, 0x01, 0x0a, 0x1a, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x60, 0x0a, 0x15, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48, 0x00, 0x52, 0x15, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xbd, 0x05, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x69, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x61, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x4a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x5c, 0x0a, 0x0c, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x74, 0x61, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x38, 0x73, 0x2e,
	0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x62, 0x0a, 0x0e,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x77, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x13,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x8f, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xa4, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4e, 0x0a, 0x0f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x0a, 0x0d, 0x47, 0x50, 0x55, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x08, 0x67,
	0x70, 0x75, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x67, 0x70, 0x75, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x50, 0x55, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x71, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x47, 0x70, 0x75, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x47,
	0x50, 0x55, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x70, 0x75, 0x54, 0x79, 0x70, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x70,
	0x75, 0x54, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x44, 0x0a, 0x16, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x47, 0x70, 0x75, 0x54, 0x79, 0x70, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x0c,
	0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3e, 0x0a, 0x0c,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x66, 0x0a, 0x0e,
	0x42, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20,
	0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x6f, 0x64, 0x65, 0x22, 0x68, 0x0a, 0x0c, 0x42, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x07, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x62, 0x65, 0x6c, 0x6f, 0x6e,
	0x67, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b,
	0x0a, 0x17, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x22, 0x28, 0x0a, 0x0c, 0x4d,
	0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61,
	0x78, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x4d, 0x69, 0x6e, 0x53, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x71, 0x0a, 0x0f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x6d, 0x0a, 0x13, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x22, 0x43, 0x0a, 0x11, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x6a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x22, 0x42, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x73, 0x0a, 0x19, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61,
	0x73, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x22, 0x49, 0x0a, 0x17, 0x44,
	0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x25, 0x0a, 0x07, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x0a,
	0x0a, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48,
	0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x43, 0x0a, 0x11, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x75, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e,
	0x75, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xb6, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x43, 0x50, 0x55,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x43, 0x50, 0x55,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x70, 0x68, 0x65,
	0x6d, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x50,
	0x6f, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x50, 0x6f, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x5a, 0x0a, 0x0f, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x73, 0x63, 0x61,
	0x6c, 0x61, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x42, 0x0a, 0x14,
	0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x7e, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x17, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73,
	0x22, 0x3e, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0xc7, 0x01, 0x0a, 0x0c, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x55, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x6f, 0x73, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x60, 0x0a, 0x11, 0x48, 0x6f, 0x73, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa8, 0x04, 0x0a, 0x08, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x64, 0x73, 0x12,
	0x3d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x75, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x49,
	0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x6e, 0x6f, 0x6e,
	0x7a, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0e,
	0x6e, 0x6f, 0x6e, 0x7a, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e,
	0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x4f,
	0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x64, 0x0a, 0x10, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x90, 0x01, 0x0a, 0x15, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x39, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x0a, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x89,
	0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x3c, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x0a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x40, 0x0a, 0x14, 0x41, 0x75, 0x74,
	0x6f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x22, 0x7e, 0x0a, 0x10, 0x4e,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x0e, 0x4e,
	0x6f, 0x64, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x16, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x7b, 0x0a, 0x0f, 0x50, 0x6f, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x65, 0x0a, 0x0d, 0x50, 0x6f, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x00, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x1b, 0x0a, 0x09,
	0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0c, 0x50, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xaf, 0x02, 0x0a, 0x0f, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x4f, 0x0a,
	0x09, 0x6d, 0x69, 0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x69, 0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x4f,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x78, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x1a,
	0x3c, 0x0a, 0x0e, 0x4d, 0x69, 0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a,
	0x0e, 0x4d, 0x61, 0x78, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9c, 0x01, 0x0a, 0x11,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x45, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x0a, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x55, 0x0a, 0x08, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2a, 0x56, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x52, 0x45, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x53, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x63, 0x2a, 0x69, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x45, 0x49, 0x4e, 0x47,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x42, 0x45, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x03, 0x32, 0xd4, 0x09, 0x0a, 0x14, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x12, 0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x10, 0x4e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x6f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x6f, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x6f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x12, 0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x7b, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x26, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x77, 0x4e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x4e, 0x65, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x08, 0x47, 0x50, 0x55, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x50, 0x55, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x50, 0x55, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x2e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x47,
	0x50, 0x55, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5c,
	0x0a, 0x07, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x12, 0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6c,
	0x65, 0x61, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x07,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x2e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0xfe, 0x0a, 0x0a, 0x10, 0x4e,
	0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x58, 0x0a, 0x07, 0x4d, 0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x78, 0x53, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x07, 0x4d, 0x69, 0x6e,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x69, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0c, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x26, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x12, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x70, 0x0a, 0x12, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65,
	0x61, 0x73, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73,
	0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x4e, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x54, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x2a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x6a,
	0x0a, 0x10, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x05, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x68, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x65, 0x64, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x07, 0x42, 0x65,
	0x6c, 0x6f, 0x6e, 0x67, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x65, 0x6c, 0x6f, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x65, 0x6c,
	0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0xc0, 0x01, 0x0a, 0x13,
	0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x08, 0x50, 0x6f,
	0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x64, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x50,
	0x6f, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x43,
	0x0a, 0x2c, 0x69, 0x6f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x42, 0x11,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x47, 0x72, 0x70,
	0x63, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudprovider_grpc_grpc_proto_rawDescOnce sync.Once
	file_cloudprovider_grpc_grpc_proto_rawDescData = file_cloudprovider_grpc_grpc_proto_rawDesc
)

func file_cloudprovider_grpc_grpc_proto_rawDescGZIP() []byte {
	file_cloudprovider_grpc_grpc_proto_rawDescOnce.Do(func() {
		file_cloudprovider_grpc_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudprovider_grpc_grpc_proto_rawDescData)
	})
	return file_cloudprovider_grpc_grpc_proto_rawDescData
}

var file_cloudprovider_grpc_grpc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cloudprovider_grpc_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 68)
var file_cloudprovider_grpc_grpc_proto_goTypes = []interface{}{
	(InstanceErrorClass)(0),             // 0: grpccloudprovider.InstanceErrorClass
	(InstanceState)(0),                  // 1: grpccloudprovider.InstanceState
	(*CloudProviderServiceRequest)(nil), // 2: grpccloudprovider.CloudProviderServiceRequest
	(*KubeAdmConfig)(nil),               // 3: grpccloudprovider.KubeAdmConfig
	(*NodeGroupDef)(nil),                // 4: grpccloudprovider.NodeGroupDef
	(*ConnectRequest)(nil),              // 5: grpccloudprovider.ConnectRequest
	(*ConnectReply)(nil),                // 6: grpccloudprovider.ConnectReply
	(*NameReply)(nil),                   // 7: grpccloudprovider.NameReply
	(*NodeGroupsReply)(nil),             // 8: grpccloudprovider.NodeGroupsReply
	(*NodeGroupForNodeRequest)(nil),     // 9: grpccloudprovider.NodeGroupForNodeRequest
	(*NodeGroupes)(nil),                 // 10: grpccloudprovider.NodeGroupes
	(*NodeGroupForNodeReply)(nil),       // 11: grpccloudprovider.NodeGroupForNodeReply
	(*PricingModelReply)(nil),           // 12: grpccloudprovider.PricingModelReply
	(*AvailableMachineTypes)(nil),       // 13: grpccloudprovider.AvailableMachineTypes
	(*AvailableMachineTypesReply)(nil),  // 14: grpccloudprovider.AvailableMachineTypesReply
	(*NewNodeGroupRequest)(nil),         // 15: grpccloudprovider.NewNodeGroupRequest
	(*NewNodeGroupReply)(nil),           // 16: grpccloudprovider.NewNodeGroupReply
	(*ResourceLimiterReply)(nil),        // 17: grpccloudprovider.ResourceLimiterReply
	(*GPULabelReply)(nil),               // 18: grpccloudprovider.GPULabelReply
	(*GetAvailableGPUTypesReply)(nil),   // 19: grpccloudprovider.GetAvailableGPUTypesReply
	(*CleanupReply)(nil),                // 20: grpccloudprovider.CleanupReply
	(*RefreshReply)(nil),                // 21: grpccloudprovider.RefreshReply
	(*BelongsRequest)(nil),              // 22: grpccloudprovider.BelongsRequest
	(*BelongsReply)(nil),                // 23: grpccloudprovider.BelongsReply
	(*NodeGroupServiceRequest)(nil),     // 24: grpccloudprovider.NodeGroupServiceRequest
	(*MaxSizeReply)(nil),                // 25: grpccloudprovider.MaxSizeReply
	(*MinSizeReply)(nil),                // 26: grpccloudprovider.MinSizeReply
	(*TargetSizeReply)(nil),             // 27: grpccloudprovider.TargetSizeReply
	(*IncreaseSizeRequest)(nil),         // 28: grpccloudprovider.IncreaseSizeRequest
	(*IncreaseSizeReply)(nil),           // 29: grpccloudprovider.IncreaseSizeReply
	(*DeleteNodesRequest)(nil),          // 30: grpccloudprovider.DeleteNodesRequest
	(*DeleteNodesReply)(nil),            // 31: grpccloudprovider.DeleteNodesReply
	(*DecreaseTargetSizeRequest)(nil),   // 32: grpccloudprovider.DecreaseTargetSizeRequest
	(*DecreaseTargetSizeReply)(nil),     // 33: grpccloudprovider.DecreaseTargetSizeReply
	(*IdReply)(nil),                     // 34: grpccloudprovider.IdReply
	(*DebugReply)(nil),                  // 35: grpccloudprovider.DebugReply
	(*Instances)(nil),                   // 36: grpccloudprovider.Instances
	(*NodesReply)(nil),                  // 37: grpccloudprovider.NodesReply
	(*ImageStateSummary)(nil),           // 38: grpccloudprovider.ImageStateSummary
	(*Resource)(nil),                    // 39: grpccloudprovider.Resource
	(*TransientSchedulerInfo)(nil),      // 40: grpccloudprovider.TransientSchedulerInfo
	(*ProtocolPort)(nil),                // 41: grpccloudprovider.ProtocolPort
	(*HostPortInfo)(nil),                // 42: grpccloudprovider.HostPortInfo
	(*NodeInfo)(nil),                    // 43: grpccloudprovider.NodeInfo
	(*TemplateNodeInfoReply)(nil),       // 44: grpccloudprovider.TemplateNodeInfoReply
	(*ExistReply)(nil),                  // 45: grpccloudprovider.ExistReply
	(*CreateReply)(nil),                 // 46: grpccloudprovider.CreateReply
	(*DeleteReply)(nil),                 // 47: grpccloudprovider.DeleteReply
	(*AutoprovisionedReply)(nil),        // 48: grpccloudprovider.AutoprovisionedReply
	(*NodePriceRequest)(nil),            // 49: grpccloudprovider.NodePriceRequest
	(*NodePriceReply)(nil),              // 50: grpccloudprovider.NodePriceReply
	(*PodPriceRequest)(nil),             // 51: grpccloudprovider.PodPriceRequest
	(*PodPriceReply)(nil),               // 52: grpccloudprovider.PodPriceReply
	(*Error)(nil),                       // 53: grpccloudprovider.Error
	(*NodeGroup)(nil),                   // 54: grpccloudprovider.NodeGroup
	(*PricingModel)(nil),                // 55: grpccloudprovider.PricingModel
	(*ResourceLimiter)(nil),             // 56: grpccloudprovider.ResourceLimiter
	(*InstanceErrorInfo)(nil),           // 57: grpccloudprovider.InstanceErrorInfo
	(*InstanceStatus)(nil),              // 58: grpccloudprovider.InstanceStatus
	(*Instance)(nil),                    // 59: grpccloudprovider.Instance
	nil,                                 // 60: grpccloudprovider.NodeGroupDef.LabelsEntry
	nil,                                 // 61: grpccloudprovider.NewNodeGroupRequest.LabelsEntry
	nil,                                 // 62: grpccloudprovider.NewNodeGroupRequest.SystemLabelsEntry
	nil,                                 // 63: grpccloudprovider.NewNodeGroupRequest.ExtraResourcesEntry
	nil,                                 // 64: grpccloudprovider.GetAvailableGPUTypesReply.AvailableGpuTypesEntry
	nil,                                 // 65: grpccloudprovider.Resource.ScalarResourcesEntry
	nil,                                 // 66: grpccloudprovider.HostPortInfo.HostPortInfoEntry
	nil,                                 // 67: grpccloudprovider.NodeInfo.ImageStatesEntry
	nil,                                 // 68: grpccloudprovider.ResourceLimiter.MinLimitsEntry
	nil,                                 // 69: grpccloudprovider.ResourceLimiter.MaxLimitsEntry
	(*v1.Taint)(nil),                    // 70: k8s.io.api.core.v1.Taint
}
var file_cloudprovider_grpc_grpc_proto_depIdxs = []int32{
	60, // 0: grpccloudprovider.NodeGroupDef.labels:type_name -> grpccloudprovider.NodeGroupDef.LabelsEntry
	56, // 1: grpccloudprovider.ConnectRequest.resourceLimiter:type_name -> grpccloudprovider.ResourceLimiter
	3,  // 2: grpccloudprovider.ConnectRequest.kubeAdmConfiguration:type_name -> grpccloudprovider.KubeAdmConfig
	4,  // 3: grpccloudprovider.ConnectRequest.nodes:type_name -> grpccloudprovider.NodeGroupDef
	53, // 4: grpccloudprovider.ConnectReply.error:type_name -> grpccloudprovider.Error
	54, // 5: grpccloudprovider.NodeGroupsReply.nodeGroups:type_name -> grpccloudprovider.NodeGroup
	54, // 6: grpccloudprovider.NodeGroupes.items:type_name -> grpccloudprovider.NodeGroup
	53, // 7: grpccloudprovider.NodeGroupForNodeReply.error:type_name -> grpccloudprovider.Error
	54, // 8: grpccloudprovider.NodeGroupForNodeReply.nodeGroup:type_name -> grpccloudprovider.NodeGroup
	53, // 9: grpccloudprovider.PricingModelReply.error:type_name -> grpccloudprovider.Error
	55, // 10: grpccloudprovider.PricingModelReply.priceModel:type_name -> grpccloudprovider.PricingModel
	53, // 11: grpccloudprovider.AvailableMachineTypesReply.error:type_name -> grpccloudprovider.Error
	13, // 12: grpccloudprovider.AvailableMachineTypesReply.availableMachineTypes:type_name -> grpccloudprovider.AvailableMachineTypes
	61, // 13: grpccloudprovider.NewNodeGroupRequest.labels:type_name -> grpccloudprovider.NewNodeGroupRequest.LabelsEntry
	62, // 14: grpccloudprovider.NewNodeGroupRequest.systemLabels:type_name -> grpccloudprovider.NewNodeGroupRequest.SystemLabelsEntry
	70, // 15: grpccloudprovider.NewNodeGroupRequest.taints:type_name -> k8s.io.api.core.v1.Taint
	63, // 16: grpccloudprovider.NewNodeGroupRequest.extraResources:type_name -> grpccloudprovider.NewNodeGroupRequest.ExtraResourcesEntry
	53, // 17: grpccloudprovider.NewNodeGroupReply.error:type_name -> grpccloudprovider.Error
	54, // 18: grpccloudprovider.NewNodeGroupReply.nodeGroup:type_name -> grpccloudprovider.NodeGroup
	53, // 19: grpccloudprovider.ResourceLimiterReply.error:type_name -> grpccloudprovider.Error
	56, // 20: grpccloudprovider.ResourceLimiterReply.resourceLimiter:type_name -> grpccloudprovider.ResourceLimiter
	53, // 21: grpccloudprovider.GPULabelReply.error:type_name -> grpccloudprovider.Error
	64, // 22: grpccloudprovider.GetAvailableGPUTypesReply.availableGpuTypes:type_name -> grpccloudprovider.GetAvailableGPUTypesReply.AvailableGpuTypesEntry
	53, // 23: grpccloudprovider.CleanupReply.error:type_name -> grpccloudprovider.Error
	53, // 24: grpccloudprovider.RefreshReply.error:type_name -> grpccloudprovider.Error
	53, // 25: grpccloudprovider.BelongsReply.error:type_name -> grpccloudprovider.Error
	53, // 26: grpccloudprovider.TargetSizeReply.error:type_name -> grpccloudprovider.Error
	53, // 27: grpccloudprovider.IncreaseSizeReply.error:type_name -> grpccloudprovider.Error
	53, // 28: grpccloudprovider.DeleteNodesReply.error:type_name -> grpccloudprovider.Error
	53, // 29: grpccloudprovider.DecreaseTargetSizeReply.error:type_name -> grpccloudprovider.Error
	59, // 30: grpccloudprovider.Instances.items:type_name -> grpccloudprovider.Instance
	53, // 31: grpccloudprovider.NodesReply.error:type_name -> grpccloudprovider.Error
	36, // 32: grpccloudprovider.NodesReply.instances:type_name -> grpccloudprovider.Instances
	65, // 33: grpccloudprovider.Resource.scalarResources:type_name -> grpccloudprovider.Resource.ScalarResourcesEntry
	66, // 34: grpccloudprovider.HostPortInfo.hostPortInfo:type_name -> grpccloudprovider.HostPortInfo.HostPortInfoEntry
	42, // 35: grpccloudprovider.NodeInfo.usedPorts:type_name -> grpccloudprovider.HostPortInfo
	39, // 36: grpccloudprovider.NodeInfo.requestedResource:type_name -> grpccloudprovider.Resource
	39, // 37: grpccloudprovider.NodeInfo.nonzeroRequest:type_name -> grpccloudprovider.Resource
	67, // 38: grpccloudprovider.NodeInfo.imageStates:type_name -> grpccloudprovider.NodeInfo.ImageStatesEntry
	40, // 39: grpccloudprovider.NodeInfo.TransientInfo:type_name -> grpccloudprovider.TransientSchedulerInfo
	53, // 40: grpccloudprovider.TemplateNodeInfoReply.error:type_name -> grpccloudprovider.Error
	43, // 41: grpccloudprovider.TemplateNodeInfoReply.nodeInfo:type_name -> grpccloudprovider.NodeInfo
	53, // 42: grpccloudprovider.CreateReply.error:type_name -> grpccloudprovider.Error
	54, // 43: grpccloudprovider.CreateReply.nodeGroup:type_name -> grpccloudprovider.NodeGroup
	53, // 44: grpccloudprovider.DeleteReply.error:type_name -> grpccloudprovider.Error
	53, // 45: grpccloudprovider.NodePriceReply.error:type_name -> grpccloudprovider.Error
	53, // 46: grpccloudprovider.PodPriceReply.error:type_name -> grpccloudprovider.Error
	68, // 47: grpccloudprovider.ResourceLimiter.minLimits:type_name -> grpccloudprovider.ResourceLimiter.MinLimitsEntry
	69, // 48: grpccloudprovider.ResourceLimiter.maxLimits:type_name -> grpccloudprovider.ResourceLimiter.MaxLimitsEntry
	0,  // 49: grpccloudprovider.InstanceErrorInfo.errorClass:type_name -> grpccloudprovider.InstanceErrorClass
	1,  // 50: grpccloudprovider.InstanceStatus.state:type_name -> grpccloudprovider.InstanceState
	57, // 51: grpccloudprovider.InstanceStatus.errorInfo:type_name -> grpccloudprovider.InstanceErrorInfo
	58, // 52: grpccloudprovider.Instance.status:type_name -> grpccloudprovider.InstanceStatus
	41, // 53: grpccloudprovider.HostPortInfo.HostPortInfoEntry.value:type_name -> grpccloudprovider.ProtocolPort
	38, // 54: grpccloudprovider.NodeInfo.ImageStatesEntry.value:type_name -> grpccloudprovider.ImageStateSummary
	5,  // 55: grpccloudprovider.CloudProviderService.Connect:input_type -> grpccloudprovider.ConnectRequest
	2,  // 56: grpccloudprovider.CloudProviderService.Name:input_type -> grpccloudprovider.CloudProviderServiceRequest
	2,  // 57: grpccloudprovider.CloudProviderService.NodeGroups:input_type -> grpccloudprovider.CloudProviderServiceRequest
	9,  // 58: grpccloudprovider.CloudProviderService.NodeGroupForNode:input_type -> grpccloudprovider.NodeGroupForNodeRequest
	2,  // 59: grpccloudprovider.CloudProviderService.Pricing:input_type -> grpccloudprovider.CloudProviderServiceRequest
	2,  // 60: grpccloudprovider.CloudProviderService.GetAvailableMachineTypes:input_type -> grpccloudprovider.CloudProviderServiceRequest
	15, // 61: grpccloudprovider.CloudProviderService.NewNodeGroup:input_type -> grpccloudprovider.NewNodeGroupRequest
	2,  // 62: grpccloudprovider.CloudProviderService.GetResourceLimiter:input_type -> grpccloudprovider.CloudProviderServiceRequest
	2,  // 63: grpccloudprovider.CloudProviderService.GPULabel:input_type -> grpccloudprovider.CloudProviderServiceRequest
	2,  // 64: grpccloudprovider.CloudProviderService.GetAvailableGPUTypes:input_type -> grpccloudprovider.CloudProviderServiceRequest
	2,  // 65: grpccloudprovider.CloudProviderService.Cleanup:input_type -> grpccloudprovider.CloudProviderServiceRequest
	2,  // 66: grpccloudprovider.CloudProviderService.Refresh:input_type -> grpccloudprovider.CloudProviderServiceRequest
	24, // 67: grpccloudprovider.NodeGroupService.MaxSize:input_type -> grpccloudprovider.NodeGroupServiceRequest
	24, // 68: grpccloudprovider.NodeGroupService.MinSize:input_type -> grpccloudprovider.NodeGroupServiceRequest
	24, // 69: grpccloudprovider.NodeGroupService.TargetSize:input_type -> grpccloudprovider.NodeGroupServiceRequest
	28, // 70: grpccloudprovider.NodeGroupService.IncreaseSize:input_type -> grpccloudprovider.IncreaseSizeRequest
	30, // 71: grpccloudprovider.NodeGroupService.DeleteNodes:input_type -> grpccloudprovider.DeleteNodesRequest
	32, // 72: grpccloudprovider.NodeGroupService.DecreaseTargetSize:input_type -> grpccloudprovider.DecreaseTargetSizeRequest
	24, // 73: grpccloudprovider.NodeGroupService.Id:input_type -> grpccloudprovider.NodeGroupServiceRequest
	24, // 74: grpccloudprovider.NodeGroupService.Debug:input_type -> grpccloudprovider.NodeGroupServiceRequest
	24, // 75: grpccloudprovider.NodeGroupService.Nodes:input_type -> grpccloudprovider.NodeGroupServiceRequest
	24, // 76: grpccloudprovider.NodeGroupService.TemplateNodeInfo:input_type -> grpccloudprovider.NodeGroupServiceRequest
	24, // 77: grpccloudprovider.NodeGroupService.Exist:input_type -> grpccloudprovider.NodeGroupServiceRequest
	24, // 78: grpccloudprovider.NodeGroupService.Create:input_type -> grpccloudprovider.NodeGroupServiceRequest
	24, // 79: grpccloudprovider.NodeGroupService.Delete:input_type -> grpccloudprovider.NodeGroupServiceRequest
	24, // 80: grpccloudprovider.NodeGroupService.Autoprovisioned:input_type -> grpccloudprovider.NodeGroupServiceRequest
	22, // 81: grpccloudprovider.NodeGroupService.Belongs:input_type -> grpccloudprovider.BelongsRequest
	49, // 82: grpccloudprovider.PricingModelService.NodePrice:input_type -> grpccloudprovider.NodePriceRequest
	51, // 83: grpccloudprovider.PricingModelService.PodPrice:input_type -> grpccloudprovider.PodPriceRequest
	6,  // 84: grpccloudprovider.CloudProviderService.Connect:output_type -> grpccloudprovider.ConnectReply
	7,  // 85: grpccloudprovider.CloudProviderService.Name:output_type -> grpccloudprovider.NameReply
	8,  // 86: grpccloudprovider.CloudProviderService.NodeGroups:output_type -> grpccloudprovider.NodeGroupsReply
	11, // 87: grpccloudprovider.CloudProviderService.NodeGroupForNode:output_type -> grpccloudprovider.NodeGroupForNodeReply
	12, // 88: grpccloudprovider.CloudProviderService.Pricing:output_type -> grpccloudprovider.PricingModelReply
	14, // 89: grpccloudprovider.CloudProviderService.GetAvailableMachineTypes:output_type -> grpccloudprovider.AvailableMachineTypesReply
	16, // 90: grpccloudprovider.CloudProviderService.NewNodeGroup:output_type -> grpccloudprovider.NewNodeGroupReply
	17, // 91: grpccloudprovider.CloudProviderService.GetResourceLimiter:output_type -> grpccloudprovider.ResourceLimiterReply
	18, // 92: grpccloudprovider.CloudProviderService.GPULabel:output_type -> grpccloudprovider.GPULabelReply
	19, // 93: grpccloudprovider.CloudProviderService.GetAvailableGPUTypes:output_type -> grpccloudprovider.GetAvailableGPUTypesReply
	20, // 94: grpccloudprovider.CloudProviderService.Cleanup:output_type -> grpccloudprovider.CleanupReply
	21, // 95: grpccloudprovider.CloudProviderService.Refresh:output_type -> grpccloudprovider.RefreshReply
	25, // 96: grpccloudprovider.NodeGroupService.MaxSize:output_type -> grpccloudprovider.MaxSizeReply
	26, // 97: grpccloudprovider.NodeGroupService.MinSize:output_type -> grpccloudprovider.MinSizeReply
	27, // 98: grpccloudprovider.NodeGroupService.TargetSize:output_type -> grpccloudprovider.TargetSizeReply
	29, // 99: grpccloudprovider.NodeGroupService.IncreaseSize:output_type -> grpccloudprovider.IncreaseSizeReply
	31, // 100: grpccloudprovider.NodeGroupService.DeleteNodes:output_type -> grpccloudprovider.DeleteNodesReply
	33, // 101: grpccloudprovider.NodeGroupService.DecreaseTargetSize:output_type -> grpccloudprovider.DecreaseTargetSizeReply
	34, // 102: grpccloudprovider.NodeGroupService.Id:output_type -> grpccloudprovider.IdReply
	35, // 103: grpccloudprovider.NodeGroupService.Debug:output_type -> grpccloudprovider.DebugReply
	37, // 104: grpccloudprovider.NodeGroupService.Nodes:output_type -> grpccloudprovider.NodesReply
	44, // 105: grpccloudprovider.NodeGroupService.TemplateNodeInfo:output_type -> grpccloudprovider.TemplateNodeInfoReply
	45, // 106: grpccloudprovider.NodeGroupService.Exist:output_type -> grpccloudprovider.ExistReply
	46, // 107: grpccloudprovider.NodeGroupService.Create:output_type -> grpccloudprovider.CreateReply
	47, // 108: grpccloudprovider.NodeGroupService.Delete:output_type -> grpccloudprovider.DeleteReply
	48, // 109: grpccloudprovider.NodeGroupService.Autoprovisioned:output_type -> grpccloudprovider.AutoprovisionedReply
	23, // 110: grpccloudprovider.NodeGroupService.Belongs:output_type -> grpccloudprovider.BelongsReply
	50, // 111: grpccloudprovider.PricingModelService.NodePrice:output_type -> grpccloudprovider.NodePriceReply
	52, // 112: grpccloudprovider.PricingModelService.PodPrice:output_type -> grpccloudprovider.PodPriceReply
	84, // [84:113] is the sub-list for method output_type
	55, // [55:84] is the sub-list for method input_type
	55, // [55:55] is the sub-list for extension type_name
	55, // [55:55] is the sub-list for extension extendee
	0,  // [0:55] is the sub-list for field type_name
}

func init() { file_cloudprovider_grpc_grpc_proto_init() }
func file_cloudprovider_grpc_grpc_proto_init() {
	if File_cloudprovider_grpc_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudprovider_grpc_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudProviderServiceRequest); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeAdmConfig); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeGroupDef); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeGroupsReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeGroupForNodeRequest); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeGroupes); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeGroupForNodeReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricingModelReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableMachineTypes); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableMachineTypesReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewNodeGroupRequest); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewNodeGroupReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceLimiterReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GPULabelReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvailableGPUTypesReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CleanupReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BelongsRequest); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BelongsReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeGroupServiceRequest); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxSizeReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinSizeReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetSizeReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[26].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseSizeRequest); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[27].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseSizeReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[28].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodesRequest); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[29].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodesReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[30].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecreaseTargetSizeRequest); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[31].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecreaseTargetSizeReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[32].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[33].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[34].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instances); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[35].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodesReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[36].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageStateSummary); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[37].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[38].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransientSchedulerInfo); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[39].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolPort); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[40].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostPortInfo); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[41].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[42].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateNodeInfoReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[43].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[44].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[45].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[46].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoprovisionedReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[47].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodePriceRequest); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[48].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodePriceReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[49].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodPriceRequest); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[50].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodPriceReply); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[51].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[52].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeGroup); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[53].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricingModel); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[54].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceLimiter); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[55].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceErrorInfo); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[56].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceStatus); i {
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
		file_cloudprovider_grpc_grpc_proto_msgTypes[57].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instance); i {
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
	file_cloudprovider_grpc_grpc_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ConnectReply_Error)(nil),
		(*ConnectReply_Connected)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*NodeGroupForNodeReply_Error)(nil),
		(*NodeGroupForNodeReply_NodeGroup)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[10].OneofWrappers = []interface{}{
		(*PricingModelReply_Error)(nil),
		(*PricingModelReply_PriceModel)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[12].OneofWrappers = []interface{}{
		(*AvailableMachineTypesReply_Error)(nil),
		(*AvailableMachineTypesReply_AvailableMachineTypes)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[14].OneofWrappers = []interface{}{
		(*NewNodeGroupReply_Error)(nil),
		(*NewNodeGroupReply_NodeGroup)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[15].OneofWrappers = []interface{}{
		(*ResourceLimiterReply_Error)(nil),
		(*ResourceLimiterReply_ResourceLimiter)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[16].OneofWrappers = []interface{}{
		(*GPULabelReply_Error)(nil),
		(*GPULabelReply_Gpulabel)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[21].OneofWrappers = []interface{}{
		(*BelongsReply_Error)(nil),
		(*BelongsReply_Belongs)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[25].OneofWrappers = []interface{}{
		(*TargetSizeReply_Error)(nil),
		(*TargetSizeReply_TargetSize)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[35].OneofWrappers = []interface{}{
		(*NodesReply_Error)(nil),
		(*NodesReply_Instances)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[42].OneofWrappers = []interface{}{
		(*TemplateNodeInfoReply_Error)(nil),
		(*TemplateNodeInfoReply_NodeInfo)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[44].OneofWrappers = []interface{}{
		(*CreateReply_Error)(nil),
		(*CreateReply_NodeGroup)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[48].OneofWrappers = []interface{}{
		(*NodePriceReply_Error)(nil),
		(*NodePriceReply_Price)(nil),
	}
	file_cloudprovider_grpc_grpc_proto_msgTypes[50].OneofWrappers = []interface{}{
		(*PodPriceReply_Error)(nil),
		(*PodPriceReply_Price)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudprovider_grpc_grpc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   68,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_cloudprovider_grpc_grpc_proto_goTypes,
		DependencyIndexes: file_cloudprovider_grpc_grpc_proto_depIdxs,
		EnumInfos:         file_cloudprovider_grpc_grpc_proto_enumTypes,
		MessageInfos:      file_cloudprovider_grpc_grpc_proto_msgTypes,
	}.Build()
	File_cloudprovider_grpc_grpc_proto = out.File
	file_cloudprovider_grpc_grpc_proto_rawDesc = nil
	file_cloudprovider_grpc_grpc_proto_goTypes = nil
	file_cloudprovider_grpc_grpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CloudProviderServiceClient is the client API for CloudProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CloudProviderServiceClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error)
	Name(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*NameReply, error)
	NodeGroups(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*NodeGroupsReply, error)
	NodeGroupForNode(ctx context.Context, in *NodeGroupForNodeRequest, opts ...grpc.CallOption) (*NodeGroupForNodeReply, error)
	Pricing(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*PricingModelReply, error)
	GetAvailableMachineTypes(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*AvailableMachineTypesReply, error)
	NewNodeGroup(ctx context.Context, in *NewNodeGroupRequest, opts ...grpc.CallOption) (*NewNodeGroupReply, error)
	GetResourceLimiter(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*ResourceLimiterReply, error)
	GPULabel(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*GPULabelReply, error)
	GetAvailableGPUTypes(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*GetAvailableGPUTypesReply, error)
	Cleanup(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*CleanupReply, error)
	Refresh(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*RefreshReply, error)
}

type cloudProviderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudProviderServiceClient(cc grpc.ClientConnInterface) CloudProviderServiceClient {
	return &cloudProviderServiceClient{cc}
}

func (c *cloudProviderServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error) {
	out := new(ConnectReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) Name(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*NameReply, error) {
	out := new(NameReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) NodeGroups(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*NodeGroupsReply, error) {
	out := new(NodeGroupsReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/NodeGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) NodeGroupForNode(ctx context.Context, in *NodeGroupForNodeRequest, opts ...grpc.CallOption) (*NodeGroupForNodeReply, error) {
	out := new(NodeGroupForNodeReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/NodeGroupForNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) Pricing(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*PricingModelReply, error) {
	out := new(PricingModelReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/Pricing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) GetAvailableMachineTypes(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*AvailableMachineTypesReply, error) {
	out := new(AvailableMachineTypesReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/GetAvailableMachineTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) NewNodeGroup(ctx context.Context, in *NewNodeGroupRequest, opts ...grpc.CallOption) (*NewNodeGroupReply, error) {
	out := new(NewNodeGroupReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/NewNodeGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) GetResourceLimiter(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*ResourceLimiterReply, error) {
	out := new(ResourceLimiterReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/GetResourceLimiter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) GPULabel(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*GPULabelReply, error) {
	out := new(GPULabelReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/GPULabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) GetAvailableGPUTypes(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*GetAvailableGPUTypesReply, error) {
	out := new(GetAvailableGPUTypesReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/GetAvailableGPUTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) Cleanup(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*CleanupReply, error) {
	out := new(CleanupReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/Cleanup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) Refresh(ctx context.Context, in *CloudProviderServiceRequest, opts ...grpc.CallOption) (*RefreshReply, error) {
	out := new(RefreshReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.CloudProviderService/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudProviderServiceServer is the server API for CloudProviderService service.
type CloudProviderServiceServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectReply, error)
	Name(context.Context, *CloudProviderServiceRequest) (*NameReply, error)
	NodeGroups(context.Context, *CloudProviderServiceRequest) (*NodeGroupsReply, error)
	NodeGroupForNode(context.Context, *NodeGroupForNodeRequest) (*NodeGroupForNodeReply, error)
	Pricing(context.Context, *CloudProviderServiceRequest) (*PricingModelReply, error)
	GetAvailableMachineTypes(context.Context, *CloudProviderServiceRequest) (*AvailableMachineTypesReply, error)
	NewNodeGroup(context.Context, *NewNodeGroupRequest) (*NewNodeGroupReply, error)
	GetResourceLimiter(context.Context, *CloudProviderServiceRequest) (*ResourceLimiterReply, error)
	GPULabel(context.Context, *CloudProviderServiceRequest) (*GPULabelReply, error)
	GetAvailableGPUTypes(context.Context, *CloudProviderServiceRequest) (*GetAvailableGPUTypesReply, error)
	Cleanup(context.Context, *CloudProviderServiceRequest) (*CleanupReply, error)
	Refresh(context.Context, *CloudProviderServiceRequest) (*RefreshReply, error)
}

// UnimplementedCloudProviderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCloudProviderServiceServer struct {
}

func (*UnimplementedCloudProviderServiceServer) Connect(context.Context, *ConnectRequest) (*ConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedCloudProviderServiceServer) Name(context.Context, *CloudProviderServiceRequest) (*NameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (*UnimplementedCloudProviderServiceServer) NodeGroups(context.Context, *CloudProviderServiceRequest) (*NodeGroupsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroups not implemented")
}
func (*UnimplementedCloudProviderServiceServer) NodeGroupForNode(context.Context, *NodeGroupForNodeRequest) (*NodeGroupForNodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGroupForNode not implemented")
}
func (*UnimplementedCloudProviderServiceServer) Pricing(context.Context, *CloudProviderServiceRequest) (*PricingModelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pricing not implemented")
}
func (*UnimplementedCloudProviderServiceServer) GetAvailableMachineTypes(context.Context, *CloudProviderServiceRequest) (*AvailableMachineTypesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableMachineTypes not implemented")
}
func (*UnimplementedCloudProviderServiceServer) NewNodeGroup(context.Context, *NewNodeGroupRequest) (*NewNodeGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewNodeGroup not implemented")
}
func (*UnimplementedCloudProviderServiceServer) GetResourceLimiter(context.Context, *CloudProviderServiceRequest) (*ResourceLimiterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceLimiter not implemented")
}
func (*UnimplementedCloudProviderServiceServer) GPULabel(context.Context, *CloudProviderServiceRequest) (*GPULabelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GPULabel not implemented")
}
func (*UnimplementedCloudProviderServiceServer) GetAvailableGPUTypes(context.Context, *CloudProviderServiceRequest) (*GetAvailableGPUTypesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableGPUTypes not implemented")
}
func (*UnimplementedCloudProviderServiceServer) Cleanup(context.Context, *CloudProviderServiceRequest) (*CleanupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cleanup not implemented")
}
func (*UnimplementedCloudProviderServiceServer) Refresh(context.Context, *CloudProviderServiceRequest) (*RefreshReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

func RegisterCloudProviderServiceServer(s *grpc.Server, srv CloudProviderServiceServer) {
	s.RegisterService(&_CloudProviderService_serviceDesc, srv)
}

func _CloudProviderService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudProviderServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).Name(ctx, req.(*CloudProviderServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_NodeGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudProviderServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).NodeGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/NodeGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).NodeGroups(ctx, req.(*CloudProviderServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_NodeGroupForNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupForNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).NodeGroupForNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/NodeGroupForNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).NodeGroupForNode(ctx, req.(*NodeGroupForNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_Pricing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudProviderServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).Pricing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/Pricing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).Pricing(ctx, req.(*CloudProviderServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_GetAvailableMachineTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudProviderServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).GetAvailableMachineTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/GetAvailableMachineTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).GetAvailableMachineTypes(ctx, req.(*CloudProviderServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_NewNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).NewNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/NewNodeGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).NewNodeGroup(ctx, req.(*NewNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_GetResourceLimiter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudProviderServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).GetResourceLimiter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/GetResourceLimiter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).GetResourceLimiter(ctx, req.(*CloudProviderServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_GPULabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudProviderServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).GPULabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/GPULabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).GPULabel(ctx, req.(*CloudProviderServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_GetAvailableGPUTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudProviderServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).GetAvailableGPUTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/GetAvailableGPUTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).GetAvailableGPUTypes(ctx, req.(*CloudProviderServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_Cleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudProviderServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).Cleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/Cleanup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).Cleanup(ctx, req.(*CloudProviderServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudProviderServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.CloudProviderService/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).Refresh(ctx, req.(*CloudProviderServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudProviderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpccloudprovider.CloudProviderService",
	HandlerType: (*CloudProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _CloudProviderService_Connect_Handler,
		},
		{
			MethodName: "Name",
			Handler:    _CloudProviderService_Name_Handler,
		},
		{
			MethodName: "NodeGroups",
			Handler:    _CloudProviderService_NodeGroups_Handler,
		},
		{
			MethodName: "NodeGroupForNode",
			Handler:    _CloudProviderService_NodeGroupForNode_Handler,
		},
		{
			MethodName: "Pricing",
			Handler:    _CloudProviderService_Pricing_Handler,
		},
		{
			MethodName: "GetAvailableMachineTypes",
			Handler:    _CloudProviderService_GetAvailableMachineTypes_Handler,
		},
		{
			MethodName: "NewNodeGroup",
			Handler:    _CloudProviderService_NewNodeGroup_Handler,
		},
		{
			MethodName: "GetResourceLimiter",
			Handler:    _CloudProviderService_GetResourceLimiter_Handler,
		},
		{
			MethodName: "GPULabel",
			Handler:    _CloudProviderService_GPULabel_Handler,
		},
		{
			MethodName: "GetAvailableGPUTypes",
			Handler:    _CloudProviderService_GetAvailableGPUTypes_Handler,
		},
		{
			MethodName: "Cleanup",
			Handler:    _CloudProviderService_Cleanup_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _CloudProviderService_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudprovider/grpc/grpc.proto",
}

// NodeGroupServiceClient is the client API for NodeGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeGroupServiceClient interface {
	MaxSize(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*MaxSizeReply, error)
	MinSize(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*MinSizeReply, error)
	TargetSize(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*TargetSizeReply, error)
	IncreaseSize(ctx context.Context, in *IncreaseSizeRequest, opts ...grpc.CallOption) (*IncreaseSizeReply, error)
	DeleteNodes(ctx context.Context, in *DeleteNodesRequest, opts ...grpc.CallOption) (*DeleteNodesReply, error)
	DecreaseTargetSize(ctx context.Context, in *DecreaseTargetSizeRequest, opts ...grpc.CallOption) (*DecreaseTargetSizeReply, error)
	Id(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*IdReply, error)
	Debug(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*DebugReply, error)
	Nodes(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*NodesReply, error)
	TemplateNodeInfo(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*TemplateNodeInfoReply, error)
	Exist(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*ExistReply, error)
	Create(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*CreateReply, error)
	Delete(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*DeleteReply, error)
	Autoprovisioned(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*AutoprovisionedReply, error)
	Belongs(ctx context.Context, in *BelongsRequest, opts ...grpc.CallOption) (*BelongsReply, error)
}

type nodeGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeGroupServiceClient(cc grpc.ClientConnInterface) NodeGroupServiceClient {
	return &nodeGroupServiceClient{cc}
}

func (c *nodeGroupServiceClient) MaxSize(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*MaxSizeReply, error) {
	out := new(MaxSizeReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/MaxSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) MinSize(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*MinSizeReply, error) {
	out := new(MinSizeReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/MinSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) TargetSize(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*TargetSizeReply, error) {
	out := new(TargetSizeReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/TargetSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) IncreaseSize(ctx context.Context, in *IncreaseSizeRequest, opts ...grpc.CallOption) (*IncreaseSizeReply, error) {
	out := new(IncreaseSizeReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/IncreaseSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) DeleteNodes(ctx context.Context, in *DeleteNodesRequest, opts ...grpc.CallOption) (*DeleteNodesReply, error) {
	out := new(DeleteNodesReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/DeleteNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) DecreaseTargetSize(ctx context.Context, in *DecreaseTargetSizeRequest, opts ...grpc.CallOption) (*DecreaseTargetSizeReply, error) {
	out := new(DecreaseTargetSizeReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/DecreaseTargetSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) Id(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*IdReply, error) {
	out := new(IdReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/Id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) Debug(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*DebugReply, error) {
	out := new(DebugReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/Debug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) Nodes(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*NodesReply, error) {
	out := new(NodesReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/Nodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) TemplateNodeInfo(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*TemplateNodeInfoReply, error) {
	out := new(TemplateNodeInfoReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/TemplateNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) Exist(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*ExistReply, error) {
	out := new(ExistReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/Exist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) Create(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) Delete(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*DeleteReply, error) {
	out := new(DeleteReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) Autoprovisioned(ctx context.Context, in *NodeGroupServiceRequest, opts ...grpc.CallOption) (*AutoprovisionedReply, error) {
	out := new(AutoprovisionedReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/Autoprovisioned", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) Belongs(ctx context.Context, in *BelongsRequest, opts ...grpc.CallOption) (*BelongsReply, error) {
	out := new(BelongsReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.NodeGroupService/Belongs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeGroupServiceServer is the server API for NodeGroupService service.
type NodeGroupServiceServer interface {
	MaxSize(context.Context, *NodeGroupServiceRequest) (*MaxSizeReply, error)
	MinSize(context.Context, *NodeGroupServiceRequest) (*MinSizeReply, error)
	TargetSize(context.Context, *NodeGroupServiceRequest) (*TargetSizeReply, error)
	IncreaseSize(context.Context, *IncreaseSizeRequest) (*IncreaseSizeReply, error)
	DeleteNodes(context.Context, *DeleteNodesRequest) (*DeleteNodesReply, error)
	DecreaseTargetSize(context.Context, *DecreaseTargetSizeRequest) (*DecreaseTargetSizeReply, error)
	Id(context.Context, *NodeGroupServiceRequest) (*IdReply, error)
	Debug(context.Context, *NodeGroupServiceRequest) (*DebugReply, error)
	Nodes(context.Context, *NodeGroupServiceRequest) (*NodesReply, error)
	TemplateNodeInfo(context.Context, *NodeGroupServiceRequest) (*TemplateNodeInfoReply, error)
	Exist(context.Context, *NodeGroupServiceRequest) (*ExistReply, error)
	Create(context.Context, *NodeGroupServiceRequest) (*CreateReply, error)
	Delete(context.Context, *NodeGroupServiceRequest) (*DeleteReply, error)
	Autoprovisioned(context.Context, *NodeGroupServiceRequest) (*AutoprovisionedReply, error)
	Belongs(context.Context, *BelongsRequest) (*BelongsReply, error)
}

// UnimplementedNodeGroupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeGroupServiceServer struct {
}

func (*UnimplementedNodeGroupServiceServer) MaxSize(context.Context, *NodeGroupServiceRequest) (*MaxSizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaxSize not implemented")
}
func (*UnimplementedNodeGroupServiceServer) MinSize(context.Context, *NodeGroupServiceRequest) (*MinSizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MinSize not implemented")
}
func (*UnimplementedNodeGroupServiceServer) TargetSize(context.Context, *NodeGroupServiceRequest) (*TargetSizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TargetSize not implemented")
}
func (*UnimplementedNodeGroupServiceServer) IncreaseSize(context.Context, *IncreaseSizeRequest) (*IncreaseSizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncreaseSize not implemented")
}
func (*UnimplementedNodeGroupServiceServer) DeleteNodes(context.Context, *DeleteNodesRequest) (*DeleteNodesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodes not implemented")
}
func (*UnimplementedNodeGroupServiceServer) DecreaseTargetSize(context.Context, *DecreaseTargetSizeRequest) (*DecreaseTargetSizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecreaseTargetSize not implemented")
}
func (*UnimplementedNodeGroupServiceServer) Id(context.Context, *NodeGroupServiceRequest) (*IdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Id not implemented")
}
func (*UnimplementedNodeGroupServiceServer) Debug(context.Context, *NodeGroupServiceRequest) (*DebugReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Debug not implemented")
}
func (*UnimplementedNodeGroupServiceServer) Nodes(context.Context, *NodeGroupServiceRequest) (*NodesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nodes not implemented")
}
func (*UnimplementedNodeGroupServiceServer) TemplateNodeInfo(context.Context, *NodeGroupServiceRequest) (*TemplateNodeInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TemplateNodeInfo not implemented")
}
func (*UnimplementedNodeGroupServiceServer) Exist(context.Context, *NodeGroupServiceRequest) (*ExistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exist not implemented")
}
func (*UnimplementedNodeGroupServiceServer) Create(context.Context, *NodeGroupServiceRequest) (*CreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedNodeGroupServiceServer) Delete(context.Context, *NodeGroupServiceRequest) (*DeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedNodeGroupServiceServer) Autoprovisioned(context.Context, *NodeGroupServiceRequest) (*AutoprovisionedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Autoprovisioned not implemented")
}
func (*UnimplementedNodeGroupServiceServer) Belongs(context.Context, *BelongsRequest) (*BelongsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Belongs not implemented")
}

func RegisterNodeGroupServiceServer(s *grpc.Server, srv NodeGroupServiceServer) {
	s.RegisterService(&_NodeGroupService_serviceDesc, srv)
}

func _NodeGroupService_MaxSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).MaxSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/MaxSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).MaxSize(ctx, req.(*NodeGroupServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_MinSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).MinSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/MinSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).MinSize(ctx, req.(*NodeGroupServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_TargetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).TargetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/TargetSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).TargetSize(ctx, req.(*NodeGroupServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_IncreaseSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncreaseSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).IncreaseSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/IncreaseSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).IncreaseSize(ctx, req.(*IncreaseSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_DeleteNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).DeleteNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/DeleteNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).DeleteNodes(ctx, req.(*DeleteNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_DecreaseTargetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecreaseTargetSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).DecreaseTargetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/DecreaseTargetSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).DecreaseTargetSize(ctx, req.(*DecreaseTargetSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_Id_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).Id(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/Id",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).Id(ctx, req.(*NodeGroupServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_Debug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).Debug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/Debug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).Debug(ctx, req.(*NodeGroupServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_Nodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).Nodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/Nodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).Nodes(ctx, req.(*NodeGroupServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_TemplateNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).TemplateNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/TemplateNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).TemplateNodeInfo(ctx, req.(*NodeGroupServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_Exist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).Exist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/Exist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).Exist(ctx, req.(*NodeGroupServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).Create(ctx, req.(*NodeGroupServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).Delete(ctx, req.(*NodeGroupServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_Autoprovisioned_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGroupServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).Autoprovisioned(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/Autoprovisioned",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).Autoprovisioned(ctx, req.(*NodeGroupServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_Belongs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BelongsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).Belongs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.NodeGroupService/Belongs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).Belongs(ctx, req.(*BelongsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeGroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpccloudprovider.NodeGroupService",
	HandlerType: (*NodeGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MaxSize",
			Handler:    _NodeGroupService_MaxSize_Handler,
		},
		{
			MethodName: "MinSize",
			Handler:    _NodeGroupService_MinSize_Handler,
		},
		{
			MethodName: "TargetSize",
			Handler:    _NodeGroupService_TargetSize_Handler,
		},
		{
			MethodName: "IncreaseSize",
			Handler:    _NodeGroupService_IncreaseSize_Handler,
		},
		{
			MethodName: "DeleteNodes",
			Handler:    _NodeGroupService_DeleteNodes_Handler,
		},
		{
			MethodName: "DecreaseTargetSize",
			Handler:    _NodeGroupService_DecreaseTargetSize_Handler,
		},
		{
			MethodName: "Id",
			Handler:    _NodeGroupService_Id_Handler,
		},
		{
			MethodName: "Debug",
			Handler:    _NodeGroupService_Debug_Handler,
		},
		{
			MethodName: "Nodes",
			Handler:    _NodeGroupService_Nodes_Handler,
		},
		{
			MethodName: "TemplateNodeInfo",
			Handler:    _NodeGroupService_TemplateNodeInfo_Handler,
		},
		{
			MethodName: "Exist",
			Handler:    _NodeGroupService_Exist_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _NodeGroupService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NodeGroupService_Delete_Handler,
		},
		{
			MethodName: "Autoprovisioned",
			Handler:    _NodeGroupService_Autoprovisioned_Handler,
		},
		{
			MethodName: "Belongs",
			Handler:    _NodeGroupService_Belongs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudprovider/grpc/grpc.proto",
}

// PricingModelServiceClient is the client API for PricingModelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PricingModelServiceClient interface {
	NodePrice(ctx context.Context, in *NodePriceRequest, opts ...grpc.CallOption) (*NodePriceReply, error)
	PodPrice(ctx context.Context, in *PodPriceRequest, opts ...grpc.CallOption) (*PodPriceReply, error)
}

type pricingModelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPricingModelServiceClient(cc grpc.ClientConnInterface) PricingModelServiceClient {
	return &pricingModelServiceClient{cc}
}

func (c *pricingModelServiceClient) NodePrice(ctx context.Context, in *NodePriceRequest, opts ...grpc.CallOption) (*NodePriceReply, error) {
	out := new(NodePriceReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.PricingModelService/NodePrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricingModelServiceClient) PodPrice(ctx context.Context, in *PodPriceRequest, opts ...grpc.CallOption) (*PodPriceReply, error) {
	out := new(PodPriceReply)
	err := c.cc.Invoke(ctx, "/grpccloudprovider.PricingModelService/PodPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PricingModelServiceServer is the server API for PricingModelService service.
type PricingModelServiceServer interface {
	NodePrice(context.Context, *NodePriceRequest) (*NodePriceReply, error)
	PodPrice(context.Context, *PodPriceRequest) (*PodPriceReply, error)
}

// UnimplementedPricingModelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPricingModelServiceServer struct {
}

func (*UnimplementedPricingModelServiceServer) NodePrice(context.Context, *NodePriceRequest) (*NodePriceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodePrice not implemented")
}
func (*UnimplementedPricingModelServiceServer) PodPrice(context.Context, *PodPriceRequest) (*PodPriceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PodPrice not implemented")
}

func RegisterPricingModelServiceServer(s *grpc.Server, srv PricingModelServiceServer) {
	s.RegisterService(&_PricingModelService_serviceDesc, srv)
}

func _PricingModelService_NodePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodePriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingModelServiceServer).NodePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.PricingModelService/NodePrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingModelServiceServer).NodePrice(ctx, req.(*NodePriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricingModelService_PodPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingModelServiceServer).PodPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccloudprovider.PricingModelService/PodPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingModelServiceServer).PodPrice(ctx, req.(*PodPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PricingModelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpccloudprovider.PricingModelService",
	HandlerType: (*PricingModelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NodePrice",
			Handler:    _PricingModelService_NodePrice_Handler,
		},
		{
			MethodName: "PodPrice",
			Handler:    _PricingModelService_PodPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudprovider/grpc/grpc.proto",
}
