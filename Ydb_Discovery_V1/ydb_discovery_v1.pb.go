// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ydb_discovery_v1.proto

package Ydb_Discovery_V1

import (
	Ydb_Discovery "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Discovery"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ydb_discovery_v1_proto protoreflect.FileDescriptor

var file_ydb_discovery_v1_proto_rawDesc = []byte{
	0x0a, 0x16, 0x79, 0x64, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x56, 0x31, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb5, 0x01, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x57, 0x68, 0x6f, 0x41, 0x6d,
	0x49, 0x12, 0x1c, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x51,
	0x0a, 0x15, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x59, 0x64, 0x62, 0x5f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_discovery_v1_proto_goTypes = []interface{}{
	(*Ydb_Discovery.ListEndpointsRequest)(nil),  // 0: Ydb.Discovery.ListEndpointsRequest
	(*Ydb_Discovery.WhoAmIRequest)(nil),         // 1: Ydb.Discovery.WhoAmIRequest
	(*Ydb_Discovery.ListEndpointsResponse)(nil), // 2: Ydb.Discovery.ListEndpointsResponse
	(*Ydb_Discovery.WhoAmIResponse)(nil),        // 3: Ydb.Discovery.WhoAmIResponse
}
var file_ydb_discovery_v1_proto_depIdxs = []int32{
	0, // 0: Ydb.Discovery.V1.DiscoveryService.ListEndpoints:input_type -> Ydb.Discovery.ListEndpointsRequest
	1, // 1: Ydb.Discovery.V1.DiscoveryService.WhoAmI:input_type -> Ydb.Discovery.WhoAmIRequest
	2, // 2: Ydb.Discovery.V1.DiscoveryService.ListEndpoints:output_type -> Ydb.Discovery.ListEndpointsResponse
	3, // 3: Ydb.Discovery.V1.DiscoveryService.WhoAmI:output_type -> Ydb.Discovery.WhoAmIResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_discovery_v1_proto_init() }
func file_ydb_discovery_v1_proto_init() {
	if File_ydb_discovery_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_discovery_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_discovery_v1_proto_goTypes,
		DependencyIndexes: file_ydb_discovery_v1_proto_depIdxs,
	}.Build()
	File_ydb_discovery_v1_proto = out.File
	file_ydb_discovery_v1_proto_rawDesc = nil
	file_ydb_discovery_v1_proto_goTypes = nil
	file_ydb_discovery_v1_proto_depIdxs = nil
}
