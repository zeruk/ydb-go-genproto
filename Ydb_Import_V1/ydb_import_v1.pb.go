// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ydb_import_v1.proto

package Ydb_Import_V1

import (
	Ydb_Import "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Import"
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

var File_ydb_import_v1_proto protoreflect.FileDescriptor

var file_ydb_import_v1_proto_rawDesc = []byte{
	0x0a, 0x13, 0x79, 0x64, 0x62, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x59, 0x64, 0x62, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x56, 0x31, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62,
	0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xaf, 0x01,
	0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x51, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x33, 0x12,
	0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x33, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x33, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1d, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x4c, 0x0a, 0x13, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x2e, 0x76, 0x31, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x59, 0x64, 0x62, 0x5f, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_import_v1_proto_goTypes = []interface{}{
	(*Ydb_Import.ImportFromS3Request)(nil),  // 0: Ydb.Import.ImportFromS3Request
	(*Ydb_Import.ImportDataRequest)(nil),    // 1: Ydb.Import.ImportDataRequest
	(*Ydb_Import.ImportFromS3Response)(nil), // 2: Ydb.Import.ImportFromS3Response
	(*Ydb_Import.ImportDataResponse)(nil),   // 3: Ydb.Import.ImportDataResponse
}
var file_ydb_import_v1_proto_depIdxs = []int32{
	0, // 0: Ydb.Import.V1.ImportService.ImportFromS3:input_type -> Ydb.Import.ImportFromS3Request
	1, // 1: Ydb.Import.V1.ImportService.ImportData:input_type -> Ydb.Import.ImportDataRequest
	2, // 2: Ydb.Import.V1.ImportService.ImportFromS3:output_type -> Ydb.Import.ImportFromS3Response
	3, // 3: Ydb.Import.V1.ImportService.ImportData:output_type -> Ydb.Import.ImportDataResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_import_v1_proto_init() }
func file_ydb_import_v1_proto_init() {
	if File_ydb_import_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_import_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_import_v1_proto_goTypes,
		DependencyIndexes: file_ydb_import_v1_proto_depIdxs,
	}.Build()
	File_ydb_import_v1_proto = out.File
	file_ydb_import_v1_proto_rawDesc = nil
	file_ydb_import_v1_proto_goTypes = nil
	file_ydb_import_v1_proto_depIdxs = nil
}
