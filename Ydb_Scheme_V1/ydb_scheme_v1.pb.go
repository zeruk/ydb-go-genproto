// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ydb_scheme_v1.proto

package Ydb_Scheme_V1

import (
	Ydb_Scheme "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Scheme"
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

var File_ydb_scheme_v1_proto protoreflect.FileDescriptor

var file_ydb_scheme_v1_proto_rawDesc = []byte{
	0x0a, 0x13, 0x79, 0x64, 0x62, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x2e, 0x56, 0x31, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcc, 0x03,
	0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x54, 0x0a, 0x0d, 0x4d, 0x61, 0x6b, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x4d, 0x61,
	0x6b, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e,
	0x4d, 0x61, 0x6b, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x54, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4b, 0x0a, 0x12,
	0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x59, 0x64, 0x62, 0x5f,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_ydb_scheme_v1_proto_goTypes = []interface{}{
	(*Ydb_Scheme.MakeDirectoryRequest)(nil),      // 0: Ydb.Scheme.MakeDirectoryRequest
	(*Ydb_Scheme.RemoveDirectoryRequest)(nil),    // 1: Ydb.Scheme.RemoveDirectoryRequest
	(*Ydb_Scheme.ListDirectoryRequest)(nil),      // 2: Ydb.Scheme.ListDirectoryRequest
	(*Ydb_Scheme.DescribePathRequest)(nil),       // 3: Ydb.Scheme.DescribePathRequest
	(*Ydb_Scheme.ModifyPermissionsRequest)(nil),  // 4: Ydb.Scheme.ModifyPermissionsRequest
	(*Ydb_Scheme.MakeDirectoryResponse)(nil),     // 5: Ydb.Scheme.MakeDirectoryResponse
	(*Ydb_Scheme.RemoveDirectoryResponse)(nil),   // 6: Ydb.Scheme.RemoveDirectoryResponse
	(*Ydb_Scheme.ListDirectoryResponse)(nil),     // 7: Ydb.Scheme.ListDirectoryResponse
	(*Ydb_Scheme.DescribePathResponse)(nil),      // 8: Ydb.Scheme.DescribePathResponse
	(*Ydb_Scheme.ModifyPermissionsResponse)(nil), // 9: Ydb.Scheme.ModifyPermissionsResponse
}
var file_ydb_scheme_v1_proto_depIdxs = []int32{
	0, // 0: Ydb.Scheme.V1.SchemeService.MakeDirectory:input_type -> Ydb.Scheme.MakeDirectoryRequest
	1, // 1: Ydb.Scheme.V1.SchemeService.RemoveDirectory:input_type -> Ydb.Scheme.RemoveDirectoryRequest
	2, // 2: Ydb.Scheme.V1.SchemeService.ListDirectory:input_type -> Ydb.Scheme.ListDirectoryRequest
	3, // 3: Ydb.Scheme.V1.SchemeService.DescribePath:input_type -> Ydb.Scheme.DescribePathRequest
	4, // 4: Ydb.Scheme.V1.SchemeService.ModifyPermissions:input_type -> Ydb.Scheme.ModifyPermissionsRequest
	5, // 5: Ydb.Scheme.V1.SchemeService.MakeDirectory:output_type -> Ydb.Scheme.MakeDirectoryResponse
	6, // 6: Ydb.Scheme.V1.SchemeService.RemoveDirectory:output_type -> Ydb.Scheme.RemoveDirectoryResponse
	7, // 7: Ydb.Scheme.V1.SchemeService.ListDirectory:output_type -> Ydb.Scheme.ListDirectoryResponse
	8, // 8: Ydb.Scheme.V1.SchemeService.DescribePath:output_type -> Ydb.Scheme.DescribePathResponse
	9, // 9: Ydb.Scheme.V1.SchemeService.ModifyPermissions:output_type -> Ydb.Scheme.ModifyPermissionsResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_scheme_v1_proto_init() }
func file_ydb_scheme_v1_proto_init() {
	if File_ydb_scheme_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_scheme_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_scheme_v1_proto_goTypes,
		DependencyIndexes: file_ydb_scheme_v1_proto_depIdxs,
	}.Build()
	File_ydb_scheme_v1_proto = out.File
	file_ydb_scheme_v1_proto_rawDesc = nil
	file_ydb_scheme_v1_proto_goTypes = nil
	file_ydb_scheme_v1_proto_depIdxs = nil
}
