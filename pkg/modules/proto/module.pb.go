// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: module.proto

package proto

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

// GeneratorRequest represents a request to generate something based on the project details
type GeneratorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Project represents the project name
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Stack represents the stack name
	Stack string `protobuf:"bytes,2,opt,name=stack,proto3" json:"stack,omitempty"`
	// App represents the application name, which is typically the same as the namespace of Kubernetes resources
	App string `protobuf:"bytes,3,opt,name=app,proto3" json:"app,omitempty"`
	// Workload represents the v1.Workload defined in the AppConfiguration
	Workload []byte `protobuf:"bytes,4,opt,name=workload,proto3" json:"workload,omitempty"`
	// DevModuleConfig is the developer's inputs of this module
	DevModuleConfig []byte `protobuf:"bytes,5,opt,name=dev_module_config,json=devModuleConfig,proto3" json:"dev_module_config,omitempty"`
	// PlatformModuleConfig is the platform engineer's inputs of this module
	PlatformModuleConfig []byte `protobuf:"bytes,6,opt,name=platform_module_config,json=platformModuleConfig,proto3" json:"platform_module_config,omitempty"`
	// RuntimeConfig is the runtime configurations defined in the workspace
	RuntimeConfig []byte `protobuf:"bytes,7,opt,name=runtime_config,json=runtimeConfig,proto3" json:"runtime_config,omitempty"`
}

func (x *GeneratorRequest) Reset() {
	*x = GeneratorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratorRequest) ProtoMessage() {}

func (x *GeneratorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratorRequest.ProtoReflect.Descriptor instead.
func (*GeneratorRequest) Descriptor() ([]byte, []int) {
	return file_module_proto_rawDescGZIP(), []int{0}
}

func (x *GeneratorRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GeneratorRequest) GetStack() string {
	if x != nil {
		return x.Stack
	}
	return ""
}

func (x *GeneratorRequest) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *GeneratorRequest) GetWorkload() []byte {
	if x != nil {
		return x.Workload
	}
	return nil
}

func (x *GeneratorRequest) GetDevModuleConfig() []byte {
	if x != nil {
		return x.DevModuleConfig
	}
	return nil
}

func (x *GeneratorRequest) GetPlatformModuleConfig() []byte {
	if x != nil {
		return x.PlatformModuleConfig
	}
	return nil
}

func (x *GeneratorRequest) GetRuntimeConfig() []byte {
	if x != nil {
		return x.RuntimeConfig
	}
	return nil
}

// GeneratorResponse represents the generate result of the generator.
type GeneratorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resources is a v1.Resource array, which represents the generated resources by this module.
	Resources [][]byte `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	Patchers  [][]byte `protobuf:"bytes,2,rep,name=patchers,proto3" json:"patchers,omitempty"`
}

func (x *GeneratorResponse) Reset() {
	*x = GeneratorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_module_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratorResponse) ProtoMessage() {}

func (x *GeneratorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_module_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratorResponse.ProtoReflect.Descriptor instead.
func (*GeneratorResponse) Descriptor() ([]byte, []int) {
	return file_module_proto_rawDescGZIP(), []int{1}
}

func (x *GeneratorResponse) GetResources() [][]byte {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *GeneratorResponse) GetPatchers() [][]byte {
	if x != nil {
		return x.Patchers
	}
	return nil
}

var File_module_proto protoreflect.FileDescriptor

var file_module_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9,
	0x01, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x64, 0x65,
	0x76, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a,
	0x16, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x4d, 0x0a, 0x11, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x08, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x32, 0x3b, 0x0a, 0x06, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x11, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_module_proto_rawDescOnce sync.Once
	file_module_proto_rawDescData = file_module_proto_rawDesc
)

func file_module_proto_rawDescGZIP() []byte {
	file_module_proto_rawDescOnce.Do(func() {
		file_module_proto_rawDescData = protoimpl.X.CompressGZIP(file_module_proto_rawDescData)
	})
	return file_module_proto_rawDescData
}

var file_module_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_module_proto_goTypes = []interface{}{
	(*GeneratorRequest)(nil),  // 0: GeneratorRequest
	(*GeneratorResponse)(nil), // 1: GeneratorResponse
}
var file_module_proto_depIdxs = []int32{
	0, // 0: Module.Generate:input_type -> GeneratorRequest
	1, // 1: Module.Generate:output_type -> GeneratorResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_module_proto_init() }
func file_module_proto_init() {
	if File_module_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_module_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratorRequest); i {
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
		file_module_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratorResponse); i {
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
			RawDescriptor: file_module_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_module_proto_goTypes,
		DependencyIndexes: file_module_proto_depIdxs,
		MessageInfos:      file_module_proto_msgTypes,
	}.Build()
	File_module_proto = out.File
	file_module_proto_rawDesc = nil
	file_module_proto_goTypes = nil
	file_module_proto_depIdxs = nil
}
