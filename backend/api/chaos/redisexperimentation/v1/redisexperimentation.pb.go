// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: chaos/redisexperimentation/v1/redisexperimentation.proto

package redisexperimentationv1

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

// The configuration of a Redis fault.
type FaultConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The targeting of the fault describing what redis requests are being considered for faults.
	FaultTargeting *FaultTargeting `protobuf:"bytes,1,opt,name=fault_targeting,json=faultTargeting,proto3" json:"fault_targeting,omitempty"`
	// Types that are assignable to Fault:
	//	*FaultConfig_ErrorFault
	//	*FaultConfig_LatencyFault
	Fault isFaultConfig_Fault `protobuf_oneof:"fault"`
}

func (x *FaultConfig) Reset() {
	*x = FaultConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultConfig) ProtoMessage() {}

func (x *FaultConfig) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultConfig.ProtoReflect.Descriptor instead.
func (*FaultConfig) Descriptor() ([]byte, []int) {
	return file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescGZIP(), []int{0}
}

func (x *FaultConfig) GetFaultTargeting() *FaultTargeting {
	if x != nil {
		return x.FaultTargeting
	}
	return nil
}

func (m *FaultConfig) GetFault() isFaultConfig_Fault {
	if m != nil {
		return m.Fault
	}
	return nil
}

func (x *FaultConfig) GetErrorFault() *ErrorFault {
	if x, ok := x.GetFault().(*FaultConfig_ErrorFault); ok {
		return x.ErrorFault
	}
	return nil
}

func (x *FaultConfig) GetLatencyFault() *LatencyFault {
	if x, ok := x.GetFault().(*FaultConfig_LatencyFault); ok {
		return x.LatencyFault
	}
	return nil
}

type isFaultConfig_Fault interface {
	isFaultConfig_Fault()
}

type FaultConfig_ErrorFault struct {
	// The error fault responds to redis commands with an error.
	ErrorFault *ErrorFault `protobuf:"bytes,2,opt,name=error_fault,json=errorFault,proto3,oneof"`
}

type FaultConfig_LatencyFault struct {
	// The latency fault delays redis commands.
	LatencyFault *LatencyFault `protobuf:"bytes,3,opt,name=latency_fault,json=latencyFault,proto3,oneof"`
}

func (*FaultConfig_ErrorFault) isFaultConfig_Fault() {}

func (*FaultConfig_LatencyFault) isFaultConfig_Fault() {}

// The definition of a redis error fault.
type ErrorFault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The percentage of requests the fault should be applied to.
	Percentage *FaultPercentage `protobuf:"bytes,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *ErrorFault) Reset() {
	*x = ErrorFault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorFault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorFault) ProtoMessage() {}

func (x *ErrorFault) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorFault.ProtoReflect.Descriptor instead.
func (*ErrorFault) Descriptor() ([]byte, []int) {
	return file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorFault) GetPercentage() *FaultPercentage {
	if x != nil {
		return x.Percentage
	}
	return nil
}

// The definition of a latency fault.
type LatencyFault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The percentage of requests the fault should be applied to.
	Percentage *FaultPercentage `protobuf:"bytes,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *LatencyFault) Reset() {
	*x = LatencyFault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatencyFault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatencyFault) ProtoMessage() {}

func (x *LatencyFault) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatencyFault.ProtoReflect.Descriptor instead.
func (*LatencyFault) Descriptor() ([]byte, []int) {
	return file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescGZIP(), []int{2}
}

func (x *LatencyFault) GetPercentage() *FaultPercentage {
	if x != nil {
		return x.Percentage
	}
	return nil
}

// Enforce faults on upstream redis cluster.
type FaultTargeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The upstream redis cluster.
	UpstreamCluster *SingleCluster `protobuf:"bytes,1,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	// A single downstream cluster sending requests to redis upstream.
	DownstreamCluster *SingleCluster `protobuf:"bytes,2,opt,name=downstream_cluster,json=downstreamCluster,proto3" json:"downstream_cluster,omitempty"`
}

func (x *FaultTargeting) Reset() {
	*x = FaultTargeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultTargeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultTargeting) ProtoMessage() {}

func (x *FaultTargeting) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultTargeting.ProtoReflect.Descriptor instead.
func (*FaultTargeting) Descriptor() ([]byte, []int) {
	return file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescGZIP(), []int{3}
}

func (x *FaultTargeting) GetUpstreamCluster() *SingleCluster {
	if x != nil {
		return x.UpstreamCluster
	}
	return nil
}

func (x *FaultTargeting) GetDownstreamCluster() *SingleCluster {
	if x != nil {
		return x.DownstreamCluster
	}
	return nil
}

// A single cluster that is partaking in the fault injection.
type SingleCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of a cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SingleCluster) Reset() {
	*x = SingleCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleCluster) ProtoMessage() {}

func (x *SingleCluster) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleCluster.ProtoReflect.Descriptor instead.
func (*SingleCluster) Descriptor() ([]byte, []int) {
	return file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescGZIP(), []int{4}
}

func (x *SingleCluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The fault percentage controlling what percentage of requests considered for a fault injection
// should have the fault applied.
type FaultPercentage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The numerator of a percentage with a fixed denumerator equal to 100
	// (i.e. percentage equal to 50 results in 50/100 = 50%)
	Percentage uint32 `protobuf:"varint,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (x *FaultPercentage) Reset() {
	*x = FaultPercentage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaultPercentage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaultPercentage) ProtoMessage() {}

func (x *FaultPercentage) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaultPercentage.ProtoReflect.Descriptor instead.
func (*FaultPercentage) Descriptor() ([]byte, []int) {
	return file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescGZIP(), []int{5}
}

func (x *FaultPercentage) GetPercentage() uint32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

var File_chaos_redisexperimentation_v1_redisexperimentation_proto protoreflect.FileDescriptor

var file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x0b, 0x46, 0x61,
	0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x67, 0x0a, 0x0f, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x53, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68,
	0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x59, 0x0a, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x5f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x64, 0x69, 0x73, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x42, 0x0c, 0x0a, 0x05, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x22, 0x6d, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x5f,
	0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x22,
	0x6f, 0x0a, 0x0c, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x12,
	0x5f, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61,
	0x6f, 0x73, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x22, 0xd4, 0x01, 0x0a, 0x0e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x5e, 0x0a, 0x10, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x0f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x62, 0x0a, 0x12, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x11, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x0d, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x0f, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x2a, 0x04, 0x18, 0x64, 0x20, 0x00, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x64, 0x69, 0x73, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescOnce sync.Once
	file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescData = file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDesc
)

func file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescGZIP() []byte {
	file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescOnce.Do(func() {
		file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescData)
	})
	return file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDescData
}

var file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chaos_redisexperimentation_v1_redisexperimentation_proto_goTypes = []interface{}{
	(*FaultConfig)(nil),     // 0: clutch.chaos.redisexperimentation.v1.FaultConfig
	(*ErrorFault)(nil),      // 1: clutch.chaos.redisexperimentation.v1.ErrorFault
	(*LatencyFault)(nil),    // 2: clutch.chaos.redisexperimentation.v1.LatencyFault
	(*FaultTargeting)(nil),  // 3: clutch.chaos.redisexperimentation.v1.FaultTargeting
	(*SingleCluster)(nil),   // 4: clutch.chaos.redisexperimentation.v1.SingleCluster
	(*FaultPercentage)(nil), // 5: clutch.chaos.redisexperimentation.v1.FaultPercentage
}
var file_chaos_redisexperimentation_v1_redisexperimentation_proto_depIdxs = []int32{
	3, // 0: clutch.chaos.redisexperimentation.v1.FaultConfig.fault_targeting:type_name -> clutch.chaos.redisexperimentation.v1.FaultTargeting
	1, // 1: clutch.chaos.redisexperimentation.v1.FaultConfig.error_fault:type_name -> clutch.chaos.redisexperimentation.v1.ErrorFault
	2, // 2: clutch.chaos.redisexperimentation.v1.FaultConfig.latency_fault:type_name -> clutch.chaos.redisexperimentation.v1.LatencyFault
	5, // 3: clutch.chaos.redisexperimentation.v1.ErrorFault.percentage:type_name -> clutch.chaos.redisexperimentation.v1.FaultPercentage
	5, // 4: clutch.chaos.redisexperimentation.v1.LatencyFault.percentage:type_name -> clutch.chaos.redisexperimentation.v1.FaultPercentage
	4, // 5: clutch.chaos.redisexperimentation.v1.FaultTargeting.upstream_cluster:type_name -> clutch.chaos.redisexperimentation.v1.SingleCluster
	4, // 6: clutch.chaos.redisexperimentation.v1.FaultTargeting.downstream_cluster:type_name -> clutch.chaos.redisexperimentation.v1.SingleCluster
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_chaos_redisexperimentation_v1_redisexperimentation_proto_init() }
func file_chaos_redisexperimentation_v1_redisexperimentation_proto_init() {
	if File_chaos_redisexperimentation_v1_redisexperimentation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultConfig); i {
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
		file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorFault); i {
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
		file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatencyFault); i {
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
		file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultTargeting); i {
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
		file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleCluster); i {
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
		file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaultPercentage); i {
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
	file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FaultConfig_ErrorFault)(nil),
		(*FaultConfig_LatencyFault)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chaos_redisexperimentation_v1_redisexperimentation_proto_goTypes,
		DependencyIndexes: file_chaos_redisexperimentation_v1_redisexperimentation_proto_depIdxs,
		MessageInfos:      file_chaos_redisexperimentation_v1_redisexperimentation_proto_msgTypes,
	}.Build()
	File_chaos_redisexperimentation_v1_redisexperimentation_proto = out.File
	file_chaos_redisexperimentation_v1_redisexperimentation_proto_rawDesc = nil
	file_chaos_redisexperimentation_v1_redisexperimentation_proto_goTypes = nil
	file_chaos_redisexperimentation_v1_redisexperimentation_proto_depIdxs = nil
}
