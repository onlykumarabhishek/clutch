// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: chaos/experimentation/v1/list_view_item.proto

package experimentationv1

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

// ListViewItem an abstraction for a list item.
type ListViewItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of a list item.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The properties map.
	Properties *PropertiesMap `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *ListViewItem) Reset() {
	*x = ListViewItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_experimentation_v1_list_view_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListViewItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListViewItem) ProtoMessage() {}

func (x *ListViewItem) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_experimentation_v1_list_view_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListViewItem.ProtoReflect.Descriptor instead.
func (*ListViewItem) Descriptor() ([]byte, []int) {
	return file_chaos_experimentation_v1_list_view_item_proto_rawDescGZIP(), []int{0}
}

func (x *ListViewItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListViewItem) GetProperties() *PropertiesMap {
	if x != nil {
		return x.Properties
	}
	return nil
}

var File_chaos_experimentation_v1_list_view_item_proto protoreflect.FileDescriptor

var file_chaos_experimentation_v1_list_view_item_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x29, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x4d, 0x61, 0x70, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x42, 0x4f, 0x5a, 0x4d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63,
	0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chaos_experimentation_v1_list_view_item_proto_rawDescOnce sync.Once
	file_chaos_experimentation_v1_list_view_item_proto_rawDescData = file_chaos_experimentation_v1_list_view_item_proto_rawDesc
)

func file_chaos_experimentation_v1_list_view_item_proto_rawDescGZIP() []byte {
	file_chaos_experimentation_v1_list_view_item_proto_rawDescOnce.Do(func() {
		file_chaos_experimentation_v1_list_view_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaos_experimentation_v1_list_view_item_proto_rawDescData)
	})
	return file_chaos_experimentation_v1_list_view_item_proto_rawDescData
}

var file_chaos_experimentation_v1_list_view_item_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chaos_experimentation_v1_list_view_item_proto_goTypes = []interface{}{
	(*ListViewItem)(nil),  // 0: clutch.chaos.experimentation.v1.ListViewItem
	(*PropertiesMap)(nil), // 1: clutch.chaos.experimentation.v1.PropertiesMap
}
var file_chaos_experimentation_v1_list_view_item_proto_depIdxs = []int32{
	1, // 0: clutch.chaos.experimentation.v1.ListViewItem.properties:type_name -> clutch.chaos.experimentation.v1.PropertiesMap
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chaos_experimentation_v1_list_view_item_proto_init() }
func file_chaos_experimentation_v1_list_view_item_proto_init() {
	if File_chaos_experimentation_v1_list_view_item_proto != nil {
		return
	}
	file_chaos_experimentation_v1_properties_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chaos_experimentation_v1_list_view_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListViewItem); i {
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
			RawDescriptor: file_chaos_experimentation_v1_list_view_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chaos_experimentation_v1_list_view_item_proto_goTypes,
		DependencyIndexes: file_chaos_experimentation_v1_list_view_item_proto_depIdxs,
		MessageInfos:      file_chaos_experimentation_v1_list_view_item_proto_msgTypes,
	}.Build()
	File_chaos_experimentation_v1_list_view_item_proto = out.File
	file_chaos_experimentation_v1_list_view_item_proto_rawDesc = nil
	file_chaos_experimentation_v1_list_view_item_proto_goTypes = nil
	file_chaos_experimentation_v1_list_view_item_proto_depIdxs = nil
}
