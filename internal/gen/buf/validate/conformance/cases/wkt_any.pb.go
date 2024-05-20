// Copyright 2023-2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/wkt_any.proto

package cases

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type AnyNone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *anypb.Any `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *AnyNone) Reset() {
	*x = AnyNone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyNone) ProtoMessage() {}

func (x *AnyNone) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyNone.ProtoReflect.Descriptor instead.
func (*AnyNone) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_any_proto_rawDescGZIP(), []int{0}
}

func (x *AnyNone) GetVal() *anypb.Any {
	if x != nil {
		return x.Val
	}
	return nil
}

type AnyRequired struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *anypb.Any `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *AnyRequired) Reset() {
	*x = AnyRequired{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyRequired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyRequired) ProtoMessage() {}

func (x *AnyRequired) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyRequired.ProtoReflect.Descriptor instead.
func (*AnyRequired) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_any_proto_rawDescGZIP(), []int{1}
}

func (x *AnyRequired) GetVal() *anypb.Any {
	if x != nil {
		return x.Val
	}
	return nil
}

type AnyIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *anypb.Any `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *AnyIn) Reset() {
	*x = AnyIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyIn) ProtoMessage() {}

func (x *AnyIn) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyIn.ProtoReflect.Descriptor instead.
func (*AnyIn) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_any_proto_rawDescGZIP(), []int{2}
}

func (x *AnyIn) GetVal() *anypb.Any {
	if x != nil {
		return x.Val
	}
	return nil
}

type AnyNotIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *anypb.Any `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *AnyNotIn) Reset() {
	*x = AnyNotIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyNotIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyNotIn) ProtoMessage() {}

func (x *AnyNotIn) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyNotIn.ProtoReflect.Descriptor instead.
func (*AnyNotIn) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_wkt_any_proto_rawDescGZIP(), []int{3}
}

func (x *AnyNotIn) GetVal() *anypb.Any {
	if x != nil {
		return x.Val
	}
	return nil
}

var File_buf_validate_conformance_cases_wkt_any_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_wkt_any_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x77, 0x6b, 0x74, 0x5f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x07, 0x41, 0x6e, 0x79, 0x4e, 0x6f, 0x6e,
	0x65, 0x12, 0x26, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x3d, 0x0a, 0x0b, 0x41, 0x6e, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x65, 0x0a, 0x05, 0x41, 0x6e, 0x79, 0x49,
	0x6e, 0x12, 0x5c, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x42, 0x34, 0xba, 0x48, 0x31, 0xa2, 0x01, 0x2e, 0x12, 0x2c, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22,
	0x69, 0x0a, 0x08, 0x41, 0x6e, 0x79, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x12, 0x5d, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x35,
	0xba, 0x48, 0x32, 0xa2, 0x01, 0x2f, 0x1a, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x42, 0x9f, 0x02, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x42, 0x0b, 0x57, 0x6b, 0x74, 0x41, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73,
	0x65, 0x73, 0xa2, 0x02, 0x04, 0x42, 0x56, 0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x73, 0xca, 0x02, 0x1e, 0x42, 0x75, 0x66,
	0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0xe2, 0x02, 0x2a, 0x42, 0x75,
	0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x42, 0x75, 0x66, 0x3a, 0x3a,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_validate_conformance_cases_wkt_any_proto_rawDescOnce sync.Once
	file_buf_validate_conformance_cases_wkt_any_proto_rawDescData = file_buf_validate_conformance_cases_wkt_any_proto_rawDesc
)

func file_buf_validate_conformance_cases_wkt_any_proto_rawDescGZIP() []byte {
	file_buf_validate_conformance_cases_wkt_any_proto_rawDescOnce.Do(func() {
		file_buf_validate_conformance_cases_wkt_any_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_conformance_cases_wkt_any_proto_rawDescData)
	})
	return file_buf_validate_conformance_cases_wkt_any_proto_rawDescData
}

var file_buf_validate_conformance_cases_wkt_any_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buf_validate_conformance_cases_wkt_any_proto_goTypes = []interface{}{
	(*AnyNone)(nil),     // 0: buf.validate.conformance.cases.AnyNone
	(*AnyRequired)(nil), // 1: buf.validate.conformance.cases.AnyRequired
	(*AnyIn)(nil),       // 2: buf.validate.conformance.cases.AnyIn
	(*AnyNotIn)(nil),    // 3: buf.validate.conformance.cases.AnyNotIn
	(*anypb.Any)(nil),   // 4: google.protobuf.Any
}
var file_buf_validate_conformance_cases_wkt_any_proto_depIdxs = []int32{
	4, // 0: buf.validate.conformance.cases.AnyNone.val:type_name -> google.protobuf.Any
	4, // 1: buf.validate.conformance.cases.AnyRequired.val:type_name -> google.protobuf.Any
	4, // 2: buf.validate.conformance.cases.AnyIn.val:type_name -> google.protobuf.Any
	4, // 3: buf.validate.conformance.cases.AnyNotIn.val:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_wkt_any_proto_init() }
func file_buf_validate_conformance_cases_wkt_any_proto_init() {
	if File_buf_validate_conformance_cases_wkt_any_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyNone); i {
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
		file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyRequired); i {
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
		file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyIn); i {
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
		file_buf_validate_conformance_cases_wkt_any_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyNotIn); i {
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
			RawDescriptor: file_buf_validate_conformance_cases_wkt_any_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_wkt_any_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_wkt_any_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_wkt_any_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_wkt_any_proto = out.File
	file_buf_validate_conformance_cases_wkt_any_proto_rawDesc = nil
	file_buf_validate_conformance_cases_wkt_any_proto_goTypes = nil
	file_buf_validate_conformance_cases_wkt_any_proto_depIdxs = nil
}
