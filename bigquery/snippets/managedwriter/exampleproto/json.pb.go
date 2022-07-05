// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The BigQuery Storage API expects protocol buffer data to be encoded in the
// proto2 wire format. This allows it to disambiguate missing optional fields
// from default values without the need for wrapper types.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: json.proto

package exampleproto

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

// Define a message type representing the rows in your table.
type JsonData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The following types map directly between protocol buffers and their
	// corresponding BigQuery data types.
	BoolCol    *bool    `protobuf:"varint,1,opt,name=bool_col,json=boolCol" json:"bool_col,omitempty"`
	Float64Col *float64 `protobuf:"fixed64,3,opt,name=float64_col,json=float64Col" json:"float64_col,omitempty"`
	Int64Col   *int64   `protobuf:"varint,4,opt,name=int64_col,json=int64Col" json:"int64_col,omitempty"`
	StringCol  *string  `protobuf:"bytes,5,opt,name=string_col,json=stringCol" json:"string_col,omitempty"`
}

func (x *JsonData) Reset() {
	*x = JsonData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_json_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonData) ProtoMessage() {}

func (x *JsonData) ProtoReflect() protoreflect.Message {
	mi := &file_json_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonData.ProtoReflect.Descriptor instead.
func (*JsonData) Descriptor() ([]byte, []int) {
	return file_json_proto_rawDescGZIP(), []int{0}
}

func (x *JsonData) GetBoolCol() bool {
	if x != nil && x.BoolCol != nil {
		return *x.BoolCol
	}
	return false
}

func (x *JsonData) GetFloat64Col() float64 {
	if x != nil && x.Float64Col != nil {
		return *x.Float64Col
	}
	return 0
}

func (x *JsonData) GetInt64Col() int64 {
	if x != nil && x.Int64Col != nil {
		return *x.Int64Col
	}
	return 0
}

func (x *JsonData) GetStringCol() string {
	if x != nil && x.StringCol != nil {
		return *x.StringCol
	}
	return ""
}

var File_json_proto protoreflect.FileDescriptor

var file_json_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x08, 0x4a,
	0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6c, 0x5f,
	0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x43,
	0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x5f, 0x63, 0x6f,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34,
	0x43, 0x6f, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x63, 0x6f, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x43, 0x6f, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x42,
	0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x64, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70,
	0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_json_proto_rawDescOnce sync.Once
	file_json_proto_rawDescData = file_json_proto_rawDesc
)

func file_json_proto_rawDescGZIP() []byte {
	file_json_proto_rawDescOnce.Do(func() {
		file_json_proto_rawDescData = protoimpl.X.CompressGZIP(file_json_proto_rawDescData)
	})
	return file_json_proto_rawDescData
}

var file_json_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_json_proto_goTypes = []interface{}{
	(*JsonData)(nil), // 0: exampleproto.JsonData
}
var file_json_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_json_proto_init() }
func file_json_proto_init() {
	if File_json_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_json_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonData); i {
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
			RawDescriptor: file_json_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_json_proto_goTypes,
		DependencyIndexes: file_json_proto_depIdxs,
		MessageInfos:      file_json_proto_msgTypes,
	}.Build()
	File_json_proto = out.File
	file_json_proto_rawDesc = nil
	file_json_proto_goTypes = nil
	file_json_proto_depIdxs = nil
}
