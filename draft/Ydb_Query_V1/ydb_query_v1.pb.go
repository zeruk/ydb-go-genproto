// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: draft/ydb_query_v1.proto

package Ydb_Query_V1

import (
	Ydb_Query "github.com/ydb-platform/ydb-go-genproto/draft/protos/Ydb_Query"
	Ydb_Operations "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Operations"
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

var File_draft_ydb_query_v1_proto protoreflect.FileDescriptor

var file_draft_ydb_query_v1_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x59, 0x64, 0x62, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x56, 0x31, 0x1a, 0x1c, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79,
	0x64, 0x62, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x98, 0x08, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x50,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x42, 0x65, 0x67,
	0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x42, 0x65,
	0x67, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x13, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0c,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1e, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72,
	0x74, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x12, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x61, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x12, 0x1c, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x12, 0x1d, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x1e, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x55, 0x0a,
	0x17, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x59, 0x64, 0x62, 0x5f, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_draft_ydb_query_v1_proto_goTypes = []interface{}{
	(*Ydb_Query.CreateSessionRequest)(nil),        // 0: Ydb.Query.CreateSessionRequest
	(*Ydb_Query.DeleteSessionRequest)(nil),        // 1: Ydb.Query.DeleteSessionRequest
	(*Ydb_Query.PingSessionRequest)(nil),          // 2: Ydb.Query.PingSessionRequest
	(*Ydb_Query.BeginTransactionRequest)(nil),     // 3: Ydb.Query.BeginTransactionRequest
	(*Ydb_Query.CommitTransactionRequest)(nil),    // 4: Ydb.Query.CommitTransactionRequest
	(*Ydb_Query.RollbackTransactionRequest)(nil),  // 5: Ydb.Query.RollbackTransactionRequest
	(*Ydb_Query.ExecuteQueryRequest)(nil),         // 6: Ydb.Query.ExecuteQueryRequest
	(*Ydb_Query.ExecuteScriptRequest)(nil),        // 7: Ydb.Query.ExecuteScriptRequest
	(*Ydb_Query.FetchScriptResultsRequest)(nil),   // 8: Ydb.Query.FetchScriptResultsRequest
	(*Ydb_Query.SaveScriptRequest)(nil),           // 9: Ydb.Query.SaveScriptRequest
	(*Ydb_Query.ListScriptsRequest)(nil),          // 10: Ydb.Query.ListScriptsRequest
	(*Ydb_Query.DeleteScriptRequest)(nil),         // 11: Ydb.Query.DeleteScriptRequest
	(*Ydb_Query.CreateSessionResponse)(nil),       // 12: Ydb.Query.CreateSessionResponse
	(*Ydb_Query.DeleteSessionResponse)(nil),       // 13: Ydb.Query.DeleteSessionResponse
	(*Ydb_Query.PingSessionResponse)(nil),         // 14: Ydb.Query.PingSessionResponse
	(*Ydb_Query.BeginTransactionResponse)(nil),    // 15: Ydb.Query.BeginTransactionResponse
	(*Ydb_Query.CommitTransactionResponse)(nil),   // 16: Ydb.Query.CommitTransactionResponse
	(*Ydb_Query.RollbackTransactionResponse)(nil), // 17: Ydb.Query.RollbackTransactionResponse
	(*Ydb_Query.ExecuteQueryResponsePart)(nil),    // 18: Ydb.Query.ExecuteQueryResponsePart
	(*Ydb_Operations.Operation)(nil),              // 19: Ydb.Operations.Operation
	(*Ydb_Query.FetchScriptResultsResponse)(nil),  // 20: Ydb.Query.FetchScriptResultsResponse
	(*Ydb_Query.SaveScriptResponse)(nil),          // 21: Ydb.Query.SaveScriptResponse
	(*Ydb_Query.ListScriptsResponse)(nil),         // 22: Ydb.Query.ListScriptsResponse
	(*Ydb_Query.DeleteScriptResponse)(nil),        // 23: Ydb.Query.DeleteScriptResponse
}
var file_draft_ydb_query_v1_proto_depIdxs = []int32{
	0,  // 0: Ydb.Query.V1.QueryService.CreateSession:input_type -> Ydb.Query.CreateSessionRequest
	1,  // 1: Ydb.Query.V1.QueryService.DeleteSession:input_type -> Ydb.Query.DeleteSessionRequest
	2,  // 2: Ydb.Query.V1.QueryService.PingSession:input_type -> Ydb.Query.PingSessionRequest
	3,  // 3: Ydb.Query.V1.QueryService.BeginTransaction:input_type -> Ydb.Query.BeginTransactionRequest
	4,  // 4: Ydb.Query.V1.QueryService.CommitTransaction:input_type -> Ydb.Query.CommitTransactionRequest
	5,  // 5: Ydb.Query.V1.QueryService.RollbackTransaction:input_type -> Ydb.Query.RollbackTransactionRequest
	6,  // 6: Ydb.Query.V1.QueryService.ExecuteQuery:input_type -> Ydb.Query.ExecuteQueryRequest
	7,  // 7: Ydb.Query.V1.QueryService.ExecuteScript:input_type -> Ydb.Query.ExecuteScriptRequest
	8,  // 8: Ydb.Query.V1.QueryService.FetchScriptResults:input_type -> Ydb.Query.FetchScriptResultsRequest
	9,  // 9: Ydb.Query.V1.QueryService.SaveScript:input_type -> Ydb.Query.SaveScriptRequest
	10, // 10: Ydb.Query.V1.QueryService.ListScripts:input_type -> Ydb.Query.ListScriptsRequest
	11, // 11: Ydb.Query.V1.QueryService.DeleteScript:input_type -> Ydb.Query.DeleteScriptRequest
	12, // 12: Ydb.Query.V1.QueryService.CreateSession:output_type -> Ydb.Query.CreateSessionResponse
	13, // 13: Ydb.Query.V1.QueryService.DeleteSession:output_type -> Ydb.Query.DeleteSessionResponse
	14, // 14: Ydb.Query.V1.QueryService.PingSession:output_type -> Ydb.Query.PingSessionResponse
	15, // 15: Ydb.Query.V1.QueryService.BeginTransaction:output_type -> Ydb.Query.BeginTransactionResponse
	16, // 16: Ydb.Query.V1.QueryService.CommitTransaction:output_type -> Ydb.Query.CommitTransactionResponse
	17, // 17: Ydb.Query.V1.QueryService.RollbackTransaction:output_type -> Ydb.Query.RollbackTransactionResponse
	18, // 18: Ydb.Query.V1.QueryService.ExecuteQuery:output_type -> Ydb.Query.ExecuteQueryResponsePart
	19, // 19: Ydb.Query.V1.QueryService.ExecuteScript:output_type -> Ydb.Operations.Operation
	20, // 20: Ydb.Query.V1.QueryService.FetchScriptResults:output_type -> Ydb.Query.FetchScriptResultsResponse
	21, // 21: Ydb.Query.V1.QueryService.SaveScript:output_type -> Ydb.Query.SaveScriptResponse
	22, // 22: Ydb.Query.V1.QueryService.ListScripts:output_type -> Ydb.Query.ListScriptsResponse
	23, // 23: Ydb.Query.V1.QueryService.DeleteScript:output_type -> Ydb.Query.DeleteScriptResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_draft_ydb_query_v1_proto_init() }
func file_draft_ydb_query_v1_proto_init() {
	if File_draft_ydb_query_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_draft_ydb_query_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_draft_ydb_query_v1_proto_goTypes,
		DependencyIndexes: file_draft_ydb_query_v1_proto_depIdxs,
	}.Build()
	File_draft_ydb_query_v1_proto = out.File
	file_draft_ydb_query_v1_proto_rawDesc = nil
	file_draft_ydb_query_v1_proto_goTypes = nil
	file_draft_ydb_query_v1_proto_depIdxs = nil
}
