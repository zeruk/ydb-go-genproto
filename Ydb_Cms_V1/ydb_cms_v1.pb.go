// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ydb_cms_v1.proto

package Ydb_Cms_V1

import (
	Ydb_Cms "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Cms"
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

var File_ydb_cms_v1_proto protoreflect.FileDescriptor

var file_ydb_cms_v1_proto_rawDesc = []byte{
	0x0a, 0x10, 0x79, 0x64, 0x62, 0x5f, 0x63, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e, 0x56, 0x31, 0x1a, 0x14,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x63, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9c, 0x04, 0x0a, 0x0a, 0x43, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x1d, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e, 0x41, 0x6c,
	0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e, 0x41, 0x6c, 0x74,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x27, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x43, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x43, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x4b, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x59, 0x64, 0x62, 0x5f, 0x43, 0x6d, 0x73, 0x5f, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_cms_v1_proto_goTypes = []interface{}{
	(*Ydb_Cms.CreateDatabaseRequest)(nil),           // 0: Ydb.Cms.CreateDatabaseRequest
	(*Ydb_Cms.GetDatabaseStatusRequest)(nil),        // 1: Ydb.Cms.GetDatabaseStatusRequest
	(*Ydb_Cms.AlterDatabaseRequest)(nil),            // 2: Ydb.Cms.AlterDatabaseRequest
	(*Ydb_Cms.ListDatabasesRequest)(nil),            // 3: Ydb.Cms.ListDatabasesRequest
	(*Ydb_Cms.RemoveDatabaseRequest)(nil),           // 4: Ydb.Cms.RemoveDatabaseRequest
	(*Ydb_Cms.DescribeDatabaseOptionsRequest)(nil),  // 5: Ydb.Cms.DescribeDatabaseOptionsRequest
	(*Ydb_Cms.CreateDatabaseResponse)(nil),          // 6: Ydb.Cms.CreateDatabaseResponse
	(*Ydb_Cms.GetDatabaseStatusResponse)(nil),       // 7: Ydb.Cms.GetDatabaseStatusResponse
	(*Ydb_Cms.AlterDatabaseResponse)(nil),           // 8: Ydb.Cms.AlterDatabaseResponse
	(*Ydb_Cms.ListDatabasesResponse)(nil),           // 9: Ydb.Cms.ListDatabasesResponse
	(*Ydb_Cms.RemoveDatabaseResponse)(nil),          // 10: Ydb.Cms.RemoveDatabaseResponse
	(*Ydb_Cms.DescribeDatabaseOptionsResponse)(nil), // 11: Ydb.Cms.DescribeDatabaseOptionsResponse
}
var file_ydb_cms_v1_proto_depIdxs = []int32{
	0,  // 0: Ydb.Cms.V1.CmsService.CreateDatabase:input_type -> Ydb.Cms.CreateDatabaseRequest
	1,  // 1: Ydb.Cms.V1.CmsService.GetDatabaseStatus:input_type -> Ydb.Cms.GetDatabaseStatusRequest
	2,  // 2: Ydb.Cms.V1.CmsService.AlterDatabase:input_type -> Ydb.Cms.AlterDatabaseRequest
	3,  // 3: Ydb.Cms.V1.CmsService.ListDatabases:input_type -> Ydb.Cms.ListDatabasesRequest
	4,  // 4: Ydb.Cms.V1.CmsService.RemoveDatabase:input_type -> Ydb.Cms.RemoveDatabaseRequest
	5,  // 5: Ydb.Cms.V1.CmsService.DescribeDatabaseOptions:input_type -> Ydb.Cms.DescribeDatabaseOptionsRequest
	6,  // 6: Ydb.Cms.V1.CmsService.CreateDatabase:output_type -> Ydb.Cms.CreateDatabaseResponse
	7,  // 7: Ydb.Cms.V1.CmsService.GetDatabaseStatus:output_type -> Ydb.Cms.GetDatabaseStatusResponse
	8,  // 8: Ydb.Cms.V1.CmsService.AlterDatabase:output_type -> Ydb.Cms.AlterDatabaseResponse
	9,  // 9: Ydb.Cms.V1.CmsService.ListDatabases:output_type -> Ydb.Cms.ListDatabasesResponse
	10, // 10: Ydb.Cms.V1.CmsService.RemoveDatabase:output_type -> Ydb.Cms.RemoveDatabaseResponse
	11, // 11: Ydb.Cms.V1.CmsService.DescribeDatabaseOptions:output_type -> Ydb.Cms.DescribeDatabaseOptionsResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_cms_v1_proto_init() }
func file_ydb_cms_v1_proto_init() {
	if File_ydb_cms_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_cms_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_cms_v1_proto_goTypes,
		DependencyIndexes: file_ydb_cms_v1_proto_depIdxs,
	}.Build()
	File_ydb_cms_v1_proto = out.File
	file_ydb_cms_v1_proto_rawDesc = nil
	file_ydb_cms_v1_proto_goTypes = nil
	file_ydb_cms_v1_proto_depIdxs = nil
}
