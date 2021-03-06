// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: config/monitoring.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	metricstores "github.com/spinnaker/kleat/api/client/metricstores"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configurable metric store types.
type MetricStoreType int32

const (
	MetricStoreType_unspecified MetricStoreType = 0
	MetricStoreType_datadog     MetricStoreType = 1
	MetricStoreType_newrelic    MetricStoreType = 2
	MetricStoreType_prometheus  MetricStoreType = 3
	MetricStoreType_stackdriver MetricStoreType = 4
)

// Enum value maps for MetricStoreType.
var (
	MetricStoreType_name = map[int32]string{
		0: "unspecified",
		1: "datadog",
		2: "newrelic",
		3: "prometheus",
		4: "stackdriver",
	}
	MetricStoreType_value = map[string]int32{
		"unspecified": 0,
		"datadog":     1,
		"newrelic":    2,
		"prometheus":  3,
		"stackdriver": 4,
	}
)

func (x MetricStoreType) Enum() *MetricStoreType {
	p := new(MetricStoreType)
	*p = x
	return p
}

func (x MetricStoreType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricStoreType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_monitoring_proto_enumTypes[0].Descriptor()
}

func (MetricStoreType) Type() protoreflect.EnumType {
	return &file_config_monitoring_proto_enumTypes[0]
}

func (x MetricStoreType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricStoreType.Descriptor instead.
func (MetricStoreType) EnumDescriptor() ([]byte, []int) {
	return file_config_monitoring_proto_rawDescGZIP(), []int{0}
}

// Configuration for the spinnaker-monitoring microservice.
// The monitoring protos use snake_case for backwards compatibility with
// Halyard-generated hal configs and the spinnaker-monitoring microservice.
// All new protos should be added using camelCase for consistency with the
// rest of the hal config.
type Monitoring struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for the Datadog metric store.
	Datadog *metricstores.Datadog `protobuf:"bytes,1,opt,name=datadog,proto3" json:"datadog,omitempty"`
	// Configuration for the New Relic metric store.
	Newrelic *metricstores.Newrelic `protobuf:"bytes,2,opt,name=newrelic,proto3" json:"newrelic,omitempty"`
	// Configuration for the Prometheus metric store.
	Prometheus *metricstores.Prometheus `protobuf:"bytes,3,opt,name=prometheus,proto3" json:"prometheus,omitempty"`
	// Configuration for the Stackdriver metric store.
	Stackdriver *metricstores.Stackdriver `protobuf:"bytes,4,opt,name=stackdriver,proto3" json:"stackdriver,omitempty"`
	// Configuration for monitoring period and enabled metric stores.
	Monitor *Monitoring_Monitor `protobuf:"bytes,5,opt,name=monitor,proto3" json:"monitor,omitempty"`
}

func (x *Monitoring) Reset() {
	*x = Monitoring{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_monitoring_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monitoring) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monitoring) ProtoMessage() {}

func (x *Monitoring) ProtoReflect() protoreflect.Message {
	mi := &file_config_monitoring_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monitoring.ProtoReflect.Descriptor instead.
func (*Monitoring) Descriptor() ([]byte, []int) {
	return file_config_monitoring_proto_rawDescGZIP(), []int{0}
}

func (x *Monitoring) GetDatadog() *metricstores.Datadog {
	if x != nil {
		return x.Datadog
	}
	return nil
}

func (x *Monitoring) GetNewrelic() *metricstores.Newrelic {
	if x != nil {
		return x.Newrelic
	}
	return nil
}

func (x *Monitoring) GetPrometheus() *metricstores.Prometheus {
	if x != nil {
		return x.Prometheus
	}
	return nil
}

func (x *Monitoring) GetStackdriver() *metricstores.Stackdriver {
	if x != nil {
		return x.Stackdriver
	}
	return nil
}

func (x *Monitoring) GetMonitor() *Monitoring_Monitor {
	if x != nil {
		return x.Monitor
	}
	return nil
}

// Configuration for monitoring period and enabled metric stores.
type Monitoring_Monitor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Polling period for the monitoring daemon (seconds). Defaults to 30.
	Period int32 `protobuf:"varint,1,opt,name=period,proto3" json:"period,omitempty"`
	// List of enabled metric stores.
	MetricStore []MetricStoreType `protobuf:"varint,2,rep,packed,name=metricStore,json=metric_store,proto3,enum=proto.config.MetricStoreType" json:"metricStore,omitempty"`
}

func (x *Monitoring_Monitor) Reset() {
	*x = Monitoring_Monitor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_monitoring_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monitoring_Monitor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monitoring_Monitor) ProtoMessage() {}

func (x *Monitoring_Monitor) ProtoReflect() protoreflect.Message {
	mi := &file_config_monitoring_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monitoring_Monitor.ProtoReflect.Descriptor instead.
func (*Monitoring_Monitor) Descriptor() ([]byte, []int) {
	return file_config_monitoring_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Monitoring_Monitor) GetPeriod() int32 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *Monitoring_Monitor) GetMetricStore() []MetricStoreType {
	if x != nil {
		return x.MetricStore
	}
	return nil
}

var File_config_monitoring_proto protoreflect.FileDescriptor

var file_config_monitoring_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x73, 0x2f, 0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa1, 0x03, 0x0a, 0x0a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x35,
	0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x52, 0x07, 0x64, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x12, 0x38, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x77,
	0x72, 0x65, 0x6c, 0x69, 0x63, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x12,
	0x3e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x65, 0x75, 0x73, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x12,
	0x41, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x3a, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x1a, 0x63,
	0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x40, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2a, 0x5e, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x64,
	0x6f, 0x67, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69, 0x63,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x10, 0x04, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_monitoring_proto_rawDescOnce sync.Once
	file_config_monitoring_proto_rawDescData = file_config_monitoring_proto_rawDesc
)

func file_config_monitoring_proto_rawDescGZIP() []byte {
	file_config_monitoring_proto_rawDescOnce.Do(func() {
		file_config_monitoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_monitoring_proto_rawDescData)
	})
	return file_config_monitoring_proto_rawDescData
}

var file_config_monitoring_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_monitoring_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_monitoring_proto_goTypes = []interface{}{
	(MetricStoreType)(0),             // 0: proto.config.MetricStoreType
	(*Monitoring)(nil),               // 1: proto.config.Monitoring
	(*Monitoring_Monitor)(nil),       // 2: proto.config.Monitoring.Monitor
	(*metricstores.Datadog)(nil),     // 3: proto.metricstores.Datadog
	(*metricstores.Newrelic)(nil),    // 4: proto.metricstores.Newrelic
	(*metricstores.Prometheus)(nil),  // 5: proto.metricstores.Prometheus
	(*metricstores.Stackdriver)(nil), // 6: proto.metricstores.Stackdriver
}
var file_config_monitoring_proto_depIdxs = []int32{
	3, // 0: proto.config.Monitoring.datadog:type_name -> proto.metricstores.Datadog
	4, // 1: proto.config.Monitoring.newrelic:type_name -> proto.metricstores.Newrelic
	5, // 2: proto.config.Monitoring.prometheus:type_name -> proto.metricstores.Prometheus
	6, // 3: proto.config.Monitoring.stackdriver:type_name -> proto.metricstores.Stackdriver
	2, // 4: proto.config.Monitoring.monitor:type_name -> proto.config.Monitoring.Monitor
	0, // 5: proto.config.Monitoring.Monitor.metricStore:type_name -> proto.config.MetricStoreType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_config_monitoring_proto_init() }
func file_config_monitoring_proto_init() {
	if File_config_monitoring_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_monitoring_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monitoring); i {
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
		file_config_monitoring_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monitoring_Monitor); i {
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
			RawDescriptor: file_config_monitoring_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_monitoring_proto_goTypes,
		DependencyIndexes: file_config_monitoring_proto_depIdxs,
		EnumInfos:         file_config_monitoring_proto_enumTypes,
		MessageInfos:      file_config_monitoring_proto_msgTypes,
	}.Build()
	File_config_monitoring_proto = out.File
	file_config_monitoring_proto_rawDesc = nil
	file_config_monitoring_proto_goTypes = nil
	file_config_monitoring_proto_depIdxs = nil
}
