// Copyright 2023 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: google/spanner/admin/instance/v1/common.proto

package instancepb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Indicates the expected fulfillment period of an operation.
type FulfillmentPeriod int32

const (
	// Not specified.
	FulfillmentPeriod_FULFILLMENT_PERIOD_UNSPECIFIED FulfillmentPeriod = 0
	// Normal fulfillment period. The operation is expected to complete within
	// minutes.
	FulfillmentPeriod_FULFILLMENT_PERIOD_NORMAL FulfillmentPeriod = 1
	// Extended fulfillment period. It can take up to an hour for the operation
	// to complete.
	FulfillmentPeriod_FULFILLMENT_PERIOD_EXTENDED FulfillmentPeriod = 2
)

// Enum value maps for FulfillmentPeriod.
var (
	FulfillmentPeriod_name = map[int32]string{
		0: "FULFILLMENT_PERIOD_UNSPECIFIED",
		1: "FULFILLMENT_PERIOD_NORMAL",
		2: "FULFILLMENT_PERIOD_EXTENDED",
	}
	FulfillmentPeriod_value = map[string]int32{
		"FULFILLMENT_PERIOD_UNSPECIFIED": 0,
		"FULFILLMENT_PERIOD_NORMAL":      1,
		"FULFILLMENT_PERIOD_EXTENDED":    2,
	}
)

func (x FulfillmentPeriod) Enum() *FulfillmentPeriod {
	p := new(FulfillmentPeriod)
	*p = x
	return p
}

func (x FulfillmentPeriod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FulfillmentPeriod) Descriptor() protoreflect.EnumDescriptor {
	return file_google_spanner_admin_instance_v1_common_proto_enumTypes[0].Descriptor()
}

func (FulfillmentPeriod) Type() protoreflect.EnumType {
	return &file_google_spanner_admin_instance_v1_common_proto_enumTypes[0]
}

func (x FulfillmentPeriod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FulfillmentPeriod.Descriptor instead.
func (FulfillmentPeriod) EnumDescriptor() ([]byte, []int) {
	return file_google_spanner_admin_instance_v1_common_proto_rawDescGZIP(), []int{0}
}

// Encapsulates progress related information for a Cloud Spanner long
// running instance operations.
type OperationProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Percent completion of the operation.
	// Values are between 0 and 100 inclusive.
	ProgressPercent int32 `protobuf:"varint,1,opt,name=progress_percent,json=progressPercent,proto3" json:"progress_percent,omitempty"`
	// Time the request was received.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// If set, the time at which this operation failed or was completed
	// successfully.
	EndTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *OperationProgress) Reset() {
	*x = OperationProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_admin_instance_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationProgress) ProtoMessage() {}

func (x *OperationProgress) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_admin_instance_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationProgress.ProtoReflect.Descriptor instead.
func (*OperationProgress) Descriptor() ([]byte, []int) {
	return file_google_spanner_admin_instance_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *OperationProgress) GetProgressPercent() int32 {
	if x != nil {
		return x.ProgressPercent
	}
	return 0
}

func (x *OperationProgress) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *OperationProgress) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_google_spanner_admin_instance_v1_common_proto protoreflect.FileDescriptor

var file_google_spanner_admin_instance_v1_common_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x77, 0x0a, 0x11, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x22, 0x0a, 0x1e, 0x46, 0x55,
	0x4c, 0x46, 0x49, 0x4c, 0x4c, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d,
	0x0a, 0x19, 0x46, 0x55, 0x4c, 0x46, 0x49, 0x4c, 0x4c, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x45,
	0x52, 0x49, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x1f, 0x0a,
	0x1b, 0x46, 0x55, 0x4c, 0x46, 0x49, 0x4c, 0x4c, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x45, 0x52,
	0x49, 0x4f, 0x44, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x02, 0x42, 0xfd,
	0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x70, 0x62, 0x3b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x70, 0x62, 0xaa, 0x02,
	0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x70,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x2b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x3a, 0x3a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_spanner_admin_instance_v1_common_proto_rawDescOnce sync.Once
	file_google_spanner_admin_instance_v1_common_proto_rawDescData = file_google_spanner_admin_instance_v1_common_proto_rawDesc
)

func file_google_spanner_admin_instance_v1_common_proto_rawDescGZIP() []byte {
	file_google_spanner_admin_instance_v1_common_proto_rawDescOnce.Do(func() {
		file_google_spanner_admin_instance_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_spanner_admin_instance_v1_common_proto_rawDescData)
	})
	return file_google_spanner_admin_instance_v1_common_proto_rawDescData
}

var file_google_spanner_admin_instance_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_spanner_admin_instance_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_spanner_admin_instance_v1_common_proto_goTypes = []interface{}{
	(FulfillmentPeriod)(0),      // 0: google.spanner.admin.instance.v1.FulfillmentPeriod
	(*OperationProgress)(nil),   // 1: google.spanner.admin.instance.v1.OperationProgress
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_google_spanner_admin_instance_v1_common_proto_depIdxs = []int32{
	2, // 0: google.spanner.admin.instance.v1.OperationProgress.start_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.spanner.admin.instance.v1.OperationProgress.end_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_spanner_admin_instance_v1_common_proto_init() }
func file_google_spanner_admin_instance_v1_common_proto_init() {
	if File_google_spanner_admin_instance_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_spanner_admin_instance_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationProgress); i {
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
			RawDescriptor: file_google_spanner_admin_instance_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_spanner_admin_instance_v1_common_proto_goTypes,
		DependencyIndexes: file_google_spanner_admin_instance_v1_common_proto_depIdxs,
		EnumInfos:         file_google_spanner_admin_instance_v1_common_proto_enumTypes,
		MessageInfos:      file_google_spanner_admin_instance_v1_common_proto_msgTypes,
	}.Build()
	File_google_spanner_admin_instance_v1_common_proto = out.File
	file_google_spanner_admin_instance_v1_common_proto_rawDesc = nil
	file_google_spanner_admin_instance_v1_common_proto_goTypes = nil
	file_google_spanner_admin_instance_v1_common_proto_depIdxs = nil
}
