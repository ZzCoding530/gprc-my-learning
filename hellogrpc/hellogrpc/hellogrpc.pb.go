// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.5.0
// source: hellogrpc/hellogrpc.proto

package hellogrpc

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

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputContent string  `protobuf:"bytes,1,opt,name=inputContent,proto3" json:"inputContent,omitempty"`
	Author       *Author `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellogrpc_hellogrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_hellogrpc_hellogrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_hellogrpc_hellogrpc_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetInputContent() string {
	if x != nil {
		return x.InputContent
	}
	return ""
}

func (x *Input) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Gender string `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellogrpc_hellogrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_hellogrpc_hellogrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_hellogrpc_hellogrpc_proto_rawDescGZIP(), []int{1}
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type OutPut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutputContent string `protobuf:"bytes,1,opt,name=outputContent,proto3" json:"outputContent,omitempty"`
	Call          string `protobuf:"bytes,2,opt,name=call,proto3" json:"call,omitempty"`
}

func (x *OutPut) Reset() {
	*x = OutPut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellogrpc_hellogrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutPut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutPut) ProtoMessage() {}

func (x *OutPut) ProtoReflect() protoreflect.Message {
	mi := &file_hellogrpc_hellogrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutPut.ProtoReflect.Descriptor instead.
func (*OutPut) Descriptor() ([]byte, []int) {
	return file_hellogrpc_hellogrpc_proto_rawDescGZIP(), []int{2}
}

func (x *OutPut) GetOutputContent() string {
	if x != nil {
		return x.OutputContent
	}
	return ""
}

func (x *OutPut) GetCall() string {
	if x != nil {
		return x.Call
	}
	return ""
}

var File_hellogrpc_hellogrpc_proto protoreflect.FileDescriptor

var file_hellogrpc_hellogrpc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x56, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x34,
	0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x50, 0x75, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x32, 0x85, 0x01, 0x0a, 0x09, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x10, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x75,
	0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x47, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x79, 0x12, 0x10, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x75, 0x74, 0x22, 0x00,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a,
	0x7a, 0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x35, 0x33, 0x30, 0x2f, 0x67, 0x70, 0x72, 0x63, 0x2d,
	0x6d, 0x79, 0x2d, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hellogrpc_hellogrpc_proto_rawDescOnce sync.Once
	file_hellogrpc_hellogrpc_proto_rawDescData = file_hellogrpc_hellogrpc_proto_rawDesc
)

func file_hellogrpc_hellogrpc_proto_rawDescGZIP() []byte {
	file_hellogrpc_hellogrpc_proto_rawDescOnce.Do(func() {
		file_hellogrpc_hellogrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_hellogrpc_hellogrpc_proto_rawDescData)
	})
	return file_hellogrpc_hellogrpc_proto_rawDescData
}

var file_hellogrpc_hellogrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hellogrpc_hellogrpc_proto_goTypes = []interface{}{
	(*Input)(nil),  // 0: hellogrpc.Input
	(*Author)(nil), // 1: hellogrpc.Author
	(*OutPut)(nil), // 2: hellogrpc.OutPut
}
var file_hellogrpc_hellogrpc_proto_depIdxs = []int32{
	1, // 0: hellogrpc.Input.author:type_name -> hellogrpc.Author
	0, // 1: hellogrpc.Translate.TranslateToEnglish:input_type -> hellogrpc.Input
	0, // 2: hellogrpc.Translate.TranslateToGermany:input_type -> hellogrpc.Input
	2, // 3: hellogrpc.Translate.TranslateToEnglish:output_type -> hellogrpc.OutPut
	2, // 4: hellogrpc.Translate.TranslateToGermany:output_type -> hellogrpc.OutPut
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hellogrpc_hellogrpc_proto_init() }
func file_hellogrpc_hellogrpc_proto_init() {
	if File_hellogrpc_hellogrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hellogrpc_hellogrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_hellogrpc_hellogrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_hellogrpc_hellogrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutPut); i {
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
			RawDescriptor: file_hellogrpc_hellogrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hellogrpc_hellogrpc_proto_goTypes,
		DependencyIndexes: file_hellogrpc_hellogrpc_proto_depIdxs,
		MessageInfos:      file_hellogrpc_hellogrpc_proto_msgTypes,
	}.Build()
	File_hellogrpc_hellogrpc_proto = out.File
	file_hellogrpc_hellogrpc_proto_rawDesc = nil
	file_hellogrpc_hellogrpc_proto_goTypes = nil
	file_hellogrpc_hellogrpc_proto_depIdxs = nil
}
