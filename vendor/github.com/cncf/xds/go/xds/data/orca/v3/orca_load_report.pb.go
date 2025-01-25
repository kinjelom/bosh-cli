// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.1
// source: xds/data/orca/v3/orca_load_report.proto

package v3

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type OrcaLoadReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuUtilization float64 `protobuf:"fixed64,1,opt,name=cpu_utilization,json=cpuUtilization,proto3" json:"cpu_utilization,omitempty"`
	MemUtilization float64 `protobuf:"fixed64,2,opt,name=mem_utilization,json=memUtilization,proto3" json:"mem_utilization,omitempty"`
	// Deprecated: Marked as deprecated in xds/data/orca/v3/orca_load_report.proto.
	Rps                    uint64             `protobuf:"varint,3,opt,name=rps,proto3" json:"rps,omitempty"`
	RequestCost            map[string]float64 `protobuf:"bytes,4,rep,name=request_cost,json=requestCost,proto3" json:"request_cost,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Utilization            map[string]float64 `protobuf:"bytes,5,rep,name=utilization,proto3" json:"utilization,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	RpsFractional          float64            `protobuf:"fixed64,6,opt,name=rps_fractional,json=rpsFractional,proto3" json:"rps_fractional,omitempty"`
	Eps                    float64            `protobuf:"fixed64,7,opt,name=eps,proto3" json:"eps,omitempty"`
	NamedMetrics           map[string]float64 `protobuf:"bytes,8,rep,name=named_metrics,json=namedMetrics,proto3" json:"named_metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	ApplicationUtilization float64            `protobuf:"fixed64,9,opt,name=application_utilization,json=applicationUtilization,proto3" json:"application_utilization,omitempty"`
}

func (x *OrcaLoadReport) Reset() {
	*x = OrcaLoadReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_data_orca_v3_orca_load_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrcaLoadReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrcaLoadReport) ProtoMessage() {}

func (x *OrcaLoadReport) ProtoReflect() protoreflect.Message {
	mi := &file_xds_data_orca_v3_orca_load_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrcaLoadReport.ProtoReflect.Descriptor instead.
func (*OrcaLoadReport) Descriptor() ([]byte, []int) {
	return file_xds_data_orca_v3_orca_load_report_proto_rawDescGZIP(), []int{0}
}

func (x *OrcaLoadReport) GetCpuUtilization() float64 {
	if x != nil {
		return x.CpuUtilization
	}
	return 0
}

func (x *OrcaLoadReport) GetMemUtilization() float64 {
	if x != nil {
		return x.MemUtilization
	}
	return 0
}

// Deprecated: Marked as deprecated in xds/data/orca/v3/orca_load_report.proto.
func (x *OrcaLoadReport) GetRps() uint64 {
	if x != nil {
		return x.Rps
	}
	return 0
}

func (x *OrcaLoadReport) GetRequestCost() map[string]float64 {
	if x != nil {
		return x.RequestCost
	}
	return nil
}

func (x *OrcaLoadReport) GetUtilization() map[string]float64 {
	if x != nil {
		return x.Utilization
	}
	return nil
}

func (x *OrcaLoadReport) GetRpsFractional() float64 {
	if x != nil {
		return x.RpsFractional
	}
	return 0
}

func (x *OrcaLoadReport) GetEps() float64 {
	if x != nil {
		return x.Eps
	}
	return 0
}

func (x *OrcaLoadReport) GetNamedMetrics() map[string]float64 {
	if x != nil {
		return x.NamedMetrics
	}
	return nil
}

func (x *OrcaLoadReport) GetApplicationUtilization() float64 {
	if x != nil {
		return x.ApplicationUtilization
	}
	return 0
}

var File_xds_data_orca_v3_orca_load_report_proto protoreflect.FileDescriptor

var file_xds_data_orca_v3_orca_load_report_proto_rawDesc = []byte{
	0x0a, 0x27, 0x78, 0x64, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6f, 0x72, 0x63, 0x61, 0x2f,
	0x76, 0x33, 0x2f, 0x6f, 0x72, 0x63, 0x61, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x78, 0x64, 0x73, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x76, 0x33, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x06, 0x0a, 0x0e, 0x4f, 0x72, 0x63, 0x61, 0x4c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x37, 0x0a, 0x0f, 0x63, 0x70, 0x75, 0x5f, 0x75,
	0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x52, 0x0e, 0x63, 0x70, 0x75, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x40, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x42, 0x17, 0xfa, 0x42, 0x14, 0x12, 0x12,
	0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x03, 0x72, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x03, 0x72, 0x70, 0x73, 0x12, 0x54, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x78, 0x64, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x76,
	0x33, 0x2e, 0x4f, 0x72, 0x63, 0x61, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x71,
	0x0a, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f,
	0x72, 0x63, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x72, 0x63, 0x61, 0x4c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x1c, 0xfa, 0x42, 0x19, 0x9a, 0x01, 0x16, 0x2a, 0x14,
	0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f, 0x29, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x52, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x35, 0x0a, 0x0e, 0x72, 0x70, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09,
	0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x0d, 0x72, 0x70, 0x73, 0x46, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x70, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x29, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x03, 0x65, 0x70, 0x73, 0x12, 0x57, 0x0a, 0x0d, 0x6e, 0x61,
	0x6d, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x72, 0x63,
	0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x72, 0x63, 0x61, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x12, 0x47, 0x0a, 0x17, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x01, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x29, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x52, 0x16, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3e, 0x0a, 0x10,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10,
	0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11,
	0x4e, 0x61, 0x6d, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x5d, 0x0a,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x78, 0x64, 0x73, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x76, 0x33, 0x42, 0x13, 0x4f, 0x72,
	0x63, 0x61, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x78, 0x64, 0x73, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x6f, 0x72, 0x63, 0x61, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xds_data_orca_v3_orca_load_report_proto_rawDescOnce sync.Once
	file_xds_data_orca_v3_orca_load_report_proto_rawDescData = file_xds_data_orca_v3_orca_load_report_proto_rawDesc
)

func file_xds_data_orca_v3_orca_load_report_proto_rawDescGZIP() []byte {
	file_xds_data_orca_v3_orca_load_report_proto_rawDescOnce.Do(func() {
		file_xds_data_orca_v3_orca_load_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_data_orca_v3_orca_load_report_proto_rawDescData)
	})
	return file_xds_data_orca_v3_orca_load_report_proto_rawDescData
}

var file_xds_data_orca_v3_orca_load_report_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_xds_data_orca_v3_orca_load_report_proto_goTypes = []interface{}{
	(*OrcaLoadReport)(nil), // 0: xds.data.orca.v3.OrcaLoadReport
	nil,                    // 1: xds.data.orca.v3.OrcaLoadReport.RequestCostEntry
	nil,                    // 2: xds.data.orca.v3.OrcaLoadReport.UtilizationEntry
	nil,                    // 3: xds.data.orca.v3.OrcaLoadReport.NamedMetricsEntry
}
var file_xds_data_orca_v3_orca_load_report_proto_depIdxs = []int32{
	1, // 0: xds.data.orca.v3.OrcaLoadReport.request_cost:type_name -> xds.data.orca.v3.OrcaLoadReport.RequestCostEntry
	2, // 1: xds.data.orca.v3.OrcaLoadReport.utilization:type_name -> xds.data.orca.v3.OrcaLoadReport.UtilizationEntry
	3, // 2: xds.data.orca.v3.OrcaLoadReport.named_metrics:type_name -> xds.data.orca.v3.OrcaLoadReport.NamedMetricsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_xds_data_orca_v3_orca_load_report_proto_init() }
func file_xds_data_orca_v3_orca_load_report_proto_init() {
	if File_xds_data_orca_v3_orca_load_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xds_data_orca_v3_orca_load_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrcaLoadReport); i {
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
			RawDescriptor: file_xds_data_orca_v3_orca_load_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xds_data_orca_v3_orca_load_report_proto_goTypes,
		DependencyIndexes: file_xds_data_orca_v3_orca_load_report_proto_depIdxs,
		MessageInfos:      file_xds_data_orca_v3_orca_load_report_proto_msgTypes,
	}.Build()
	File_xds_data_orca_v3_orca_load_report_proto = out.File
	file_xds_data_orca_v3_orca_load_report_proto_rawDesc = nil
	file_xds_data_orca_v3_orca_load_report_proto_goTypes = nil
	file_xds_data_orca_v3_orca_load_report_proto_depIdxs = nil
}
