// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: resolver/v1/resolver_api.proto

package resolverv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AutocompleteResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The primary identifer of the resource
	// This value is used when rendering the autocomplete results on the frontend
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The label can provide more detailed information about the result
	// This is dispalyed inline with the auto complete result
	// eg: what region an ec2 instance resides in
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *AutocompleteResult) Reset() {
	*x = AutocompleteResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_resolver_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutocompleteResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutocompleteResult) ProtoMessage() {}

func (x *AutocompleteResult) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_resolver_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutocompleteResult.ProtoReflect.Descriptor instead.
func (*AutocompleteResult) Descriptor() ([]byte, []int) {
	return file_resolver_v1_resolver_api_proto_rawDescGZIP(), []int{0}
}

func (x *AutocompleteResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AutocompleteResult) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type AutocompleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type URL of the desired result.
	Want   string `protobuf:"bytes,1,opt,name=want,proto3" json:"want,omitempty"`
	Search string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
	Limit  uint64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *AutocompleteRequest) Reset() {
	*x = AutocompleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_resolver_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutocompleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutocompleteRequest) ProtoMessage() {}

func (x *AutocompleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_resolver_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutocompleteRequest.ProtoReflect.Descriptor instead.
func (*AutocompleteRequest) Descriptor() ([]byte, []int) {
	return file_resolver_v1_resolver_api_proto_rawDescGZIP(), []int{1}
}

func (x *AutocompleteRequest) GetWant() string {
	if x != nil {
		return x.Want
	}
	return ""
}

func (x *AutocompleteRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *AutocompleteRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type AutocompleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*AutocompleteResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *AutocompleteResponse) Reset() {
	*x = AutocompleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_resolver_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutocompleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutocompleteResponse) ProtoMessage() {}

func (x *AutocompleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_resolver_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutocompleteResponse.ProtoReflect.Descriptor instead.
func (*AutocompleteResponse) Descriptor() ([]byte, []int) {
	return file_resolver_v1_resolver_api_proto_rawDescGZIP(), []int{2}
}

func (x *AutocompleteResponse) GetResults() []*AutocompleteResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type ResolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type URL of the desired result.
	Want string `protobuf:"bytes,1,opt,name=want,proto3" json:"want,omitempty"`
	// Filled in object schemas.
	Have *anypb.Any `protobuf:"bytes,2,opt,name=have,proto3" json:"have,omitempty"`
	// The maximum number of results to return.
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ResolveRequest) Reset() {
	*x = ResolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_resolver_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveRequest) ProtoMessage() {}

func (x *ResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_resolver_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveRequest.ProtoReflect.Descriptor instead.
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return file_resolver_v1_resolver_api_proto_rawDescGZIP(), []int{3}
}

func (x *ResolveRequest) GetWant() string {
	if x != nil {
		return x.Want
	}
	return ""
}

func (x *ResolveRequest) GetHave() *anypb.Any {
	if x != nil {
		return x.Have
	}
	return nil
}

func (x *ResolveRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ResolveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results         []*anypb.Any     `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	PartialFailures []*status.Status `protobuf:"bytes,2,rep,name=partial_failures,json=partialFailures,proto3" json:"partial_failures,omitempty"`
}

func (x *ResolveResponse) Reset() {
	*x = ResolveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_resolver_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveResponse) ProtoMessage() {}

func (x *ResolveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_resolver_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveResponse.ProtoReflect.Descriptor instead.
func (*ResolveResponse) Descriptor() ([]byte, []int) {
	return file_resolver_v1_resolver_api_proto_rawDescGZIP(), []int{4}
}

func (x *ResolveResponse) GetResults() []*anypb.Any {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ResolveResponse) GetPartialFailures() []*status.Status {
	if x != nil {
		return x.PartialFailures
	}
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type URL of the desired result.
	Want string `protobuf:"bytes,1,opt,name=want,proto3" json:"want,omitempty"`
	// Free-form text query.
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// The maximum number of results to return.
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_resolver_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_resolver_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_resolver_v1_resolver_api_proto_rawDescGZIP(), []int{5}
}

func (x *SearchRequest) GetWant() string {
	if x != nil {
		return x.Want
	}
	return ""
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results         []*anypb.Any     `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	PartialFailures []*status.Status `protobuf:"bytes,2,rep,name=partial_failures,json=partialFailures,proto3" json:"partial_failures,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_resolver_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_resolver_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_resolver_v1_resolver_api_proto_rawDescGZIP(), []int{6}
}

func (x *SearchResponse) GetResults() []*anypb.Any {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SearchResponse) GetPartialFailures() []*status.Status {
	if x != nil {
		return x.PartialFailures
	}
	return nil
}

type GetObjectSchemasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
}

func (x *GetObjectSchemasRequest) Reset() {
	*x = GetObjectSchemasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_resolver_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectSchemasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectSchemasRequest) ProtoMessage() {}

func (x *GetObjectSchemasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_resolver_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectSchemasRequest.ProtoReflect.Descriptor instead.
func (*GetObjectSchemasRequest) Descriptor() ([]byte, []int) {
	return file_resolver_v1_resolver_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetObjectSchemasRequest) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

type GetObjectSchemasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeUrl string    `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Schemas []*Schema `protobuf:"bytes,2,rep,name=schemas,proto3" json:"schemas,omitempty"`
}

func (x *GetObjectSchemasResponse) Reset() {
	*x = GetObjectSchemasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_resolver_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectSchemasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectSchemasResponse) ProtoMessage() {}

func (x *GetObjectSchemasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_resolver_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectSchemasResponse.ProtoReflect.Descriptor instead.
func (*GetObjectSchemasResponse) Descriptor() ([]byte, []int) {
	return file_resolver_v1_resolver_api_proto_rawDescGZIP(), []int{8}
}

func (x *GetObjectSchemasResponse) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *GetObjectSchemasResponse) GetSchemas() []*Schema {
	if x != nil {
		return x.Schemas
	}
	return nil
}

var File_resolver_v1_resolver_api_proto protoreflect.FileDescriptor

var file_resolver_v1_resolver_api_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22,
	0x69, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x77, 0x61, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x77,
	0x61, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x58, 0x0a, 0x14, 0x41, 0x75,
	0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x22, 0x77, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x77, 0x61, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x77,
	0x61, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x68, 0x61, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x04, 0x68, 0x61, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x80, 0x01,
	0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x12, 0x3d, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73,
	0x22, 0x61, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x04, 0x77, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x77, 0x61, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x7f, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65,
	0x55, 0x72, 0x6c, 0x22, 0x6b, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6c,
	0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x32, 0xaf, 0x04, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x41, 0x50, 0x49,
	0x12, 0x9d, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x2b, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x3a, 0x01, 0x2a, 0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x02,
	0x12, 0x75, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x21, 0x2e, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x01,
	0x2a, 0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x02, 0x12, 0x79, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x12, 0x22, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0xaa, 0xe1, 0x1c, 0x02,
	0x08, 0x02, 0x12, 0x8d, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63,
	0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xaa, 0xe1, 0x1c, 0x02,
	0x08, 0x02, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resolver_v1_resolver_api_proto_rawDescOnce sync.Once
	file_resolver_v1_resolver_api_proto_rawDescData = file_resolver_v1_resolver_api_proto_rawDesc
)

func file_resolver_v1_resolver_api_proto_rawDescGZIP() []byte {
	file_resolver_v1_resolver_api_proto_rawDescOnce.Do(func() {
		file_resolver_v1_resolver_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_resolver_v1_resolver_api_proto_rawDescData)
	})
	return file_resolver_v1_resolver_api_proto_rawDescData
}

var file_resolver_v1_resolver_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_resolver_v1_resolver_api_proto_goTypes = []interface{}{
	(*AutocompleteResult)(nil),       // 0: clutch.resolver.v1.AutocompleteResult
	(*AutocompleteRequest)(nil),      // 1: clutch.resolver.v1.AutocompleteRequest
	(*AutocompleteResponse)(nil),     // 2: clutch.resolver.v1.AutocompleteResponse
	(*ResolveRequest)(nil),           // 3: clutch.resolver.v1.ResolveRequest
	(*ResolveResponse)(nil),          // 4: clutch.resolver.v1.ResolveResponse
	(*SearchRequest)(nil),            // 5: clutch.resolver.v1.SearchRequest
	(*SearchResponse)(nil),           // 6: clutch.resolver.v1.SearchResponse
	(*GetObjectSchemasRequest)(nil),  // 7: clutch.resolver.v1.GetObjectSchemasRequest
	(*GetObjectSchemasResponse)(nil), // 8: clutch.resolver.v1.GetObjectSchemasResponse
	(*anypb.Any)(nil),                // 9: google.protobuf.Any
	(*status.Status)(nil),            // 10: google.rpc.Status
	(*Schema)(nil),                   // 11: clutch.resolver.v1.Schema
}
var file_resolver_v1_resolver_api_proto_depIdxs = []int32{
	0,  // 0: clutch.resolver.v1.AutocompleteResponse.results:type_name -> clutch.resolver.v1.AutocompleteResult
	9,  // 1: clutch.resolver.v1.ResolveRequest.have:type_name -> google.protobuf.Any
	9,  // 2: clutch.resolver.v1.ResolveResponse.results:type_name -> google.protobuf.Any
	10, // 3: clutch.resolver.v1.ResolveResponse.partial_failures:type_name -> google.rpc.Status
	9,  // 4: clutch.resolver.v1.SearchResponse.results:type_name -> google.protobuf.Any
	10, // 5: clutch.resolver.v1.SearchResponse.partial_failures:type_name -> google.rpc.Status
	11, // 6: clutch.resolver.v1.GetObjectSchemasResponse.schemas:type_name -> clutch.resolver.v1.Schema
	7,  // 7: clutch.resolver.v1.ResolverAPI.GetObjectSchemas:input_type -> clutch.resolver.v1.GetObjectSchemasRequest
	5,  // 8: clutch.resolver.v1.ResolverAPI.Search:input_type -> clutch.resolver.v1.SearchRequest
	3,  // 9: clutch.resolver.v1.ResolverAPI.Resolve:input_type -> clutch.resolver.v1.ResolveRequest
	1,  // 10: clutch.resolver.v1.ResolverAPI.Autocomplete:input_type -> clutch.resolver.v1.AutocompleteRequest
	8,  // 11: clutch.resolver.v1.ResolverAPI.GetObjectSchemas:output_type -> clutch.resolver.v1.GetObjectSchemasResponse
	6,  // 12: clutch.resolver.v1.ResolverAPI.Search:output_type -> clutch.resolver.v1.SearchResponse
	4,  // 13: clutch.resolver.v1.ResolverAPI.Resolve:output_type -> clutch.resolver.v1.ResolveResponse
	2,  // 14: clutch.resolver.v1.ResolverAPI.Autocomplete:output_type -> clutch.resolver.v1.AutocompleteResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_resolver_v1_resolver_api_proto_init() }
func file_resolver_v1_resolver_api_proto_init() {
	if File_resolver_v1_resolver_api_proto != nil {
		return
	}
	file_resolver_v1_schema_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_resolver_v1_resolver_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutocompleteResult); i {
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
		file_resolver_v1_resolver_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutocompleteRequest); i {
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
		file_resolver_v1_resolver_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutocompleteResponse); i {
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
		file_resolver_v1_resolver_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveRequest); i {
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
		file_resolver_v1_resolver_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveResponse); i {
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
		file_resolver_v1_resolver_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_resolver_v1_resolver_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_resolver_v1_resolver_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectSchemasRequest); i {
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
		file_resolver_v1_resolver_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectSchemasResponse); i {
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
			RawDescriptor: file_resolver_v1_resolver_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resolver_v1_resolver_api_proto_goTypes,
		DependencyIndexes: file_resolver_v1_resolver_api_proto_depIdxs,
		MessageInfos:      file_resolver_v1_resolver_api_proto_msgTypes,
	}.Build()
	File_resolver_v1_resolver_api_proto = out.File
	file_resolver_v1_resolver_api_proto_rawDesc = nil
	file_resolver_v1_resolver_api_proto_goTypes = nil
	file_resolver_v1_resolver_api_proto_depIdxs = nil
}
