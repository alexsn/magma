// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/mconfig/mconfigs.proto

package mconfig // import "magma/orc8r/cloud/go/protos/mconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protos "magma/orc8r/cloud/go/protos"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ------------------------------------------------------------------------------
// Control Proxy configs
// ------------------------------------------------------------------------------
type ControlProxy struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ControlProxy) Reset()         { *m = ControlProxy{} }
func (m *ControlProxy) String() string { return proto.CompactTextString(m) }
func (*ControlProxy) ProtoMessage()    {}
func (*ControlProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_5bc13885cab21ab6, []int{0}
}
func (m *ControlProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlProxy.Unmarshal(m, b)
}
func (m *ControlProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlProxy.Marshal(b, m, deterministic)
}
func (dst *ControlProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlProxy.Merge(dst, src)
}
func (m *ControlProxy) XXX_Size() int {
	return xxx_messageInfo_ControlProxy.Size(m)
}
func (m *ControlProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ControlProxy proto.InternalMessageInfo

func (m *ControlProxy) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

// ------------------------------------------------------------------------------
// DnsD configs
// ------------------------------------------------------------------------------
type DnsD struct {
	LogLevel             protos.LogLevel                 `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	EnableCaching        bool                            `protobuf:"varint,2,opt,name=enable_caching,json=enableCaching,proto3" json:"enable_caching,omitempty"`
	LocalTTL             int32                           `protobuf:"varint,3,opt,name=localTTL,proto3" json:"localTTL,omitempty"`
	Records              []*NetworkDNSConfigRecordsItems `protobuf:"bytes,4,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DnsD) Reset()         { *m = DnsD{} }
func (m *DnsD) String() string { return proto.CompactTextString(m) }
func (*DnsD) ProtoMessage()    {}
func (*DnsD) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_5bc13885cab21ab6, []int{1}
}
func (m *DnsD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsD.Unmarshal(m, b)
}
func (m *DnsD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsD.Marshal(b, m, deterministic)
}
func (dst *DnsD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsD.Merge(dst, src)
}
func (m *DnsD) XXX_Size() int {
	return xxx_messageInfo_DnsD.Size(m)
}
func (m *DnsD) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsD.DiscardUnknown(m)
}

var xxx_messageInfo_DnsD proto.InternalMessageInfo

func (m *DnsD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *DnsD) GetEnableCaching() bool {
	if m != nil {
		return m.EnableCaching
	}
	return false
}

func (m *DnsD) GetLocalTTL() int32 {
	if m != nil {
		return m.LocalTTL
	}
	return 0
}

func (m *DnsD) GetRecords() []*NetworkDNSConfigRecordsItems {
	if m != nil {
		return m.Records
	}
	return nil
}

type NetworkDNSConfigRecordsItems struct {
	ARecord              []string `protobuf:"bytes,1,rep,name=a_record,json=aRecord,proto3" json:"a_record,omitempty"`
	AaaaRecord           []string `protobuf:"bytes,2,rep,name=aaaa_record,json=aaaaRecord,proto3" json:"aaaa_record,omitempty"`
	CnameRecord          []string `protobuf:"bytes,3,rep,name=cname_record,json=cnameRecord,proto3" json:"cname_record,omitempty"`
	Domain               string   `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkDNSConfigRecordsItems) Reset()         { *m = NetworkDNSConfigRecordsItems{} }
func (m *NetworkDNSConfigRecordsItems) String() string { return proto.CompactTextString(m) }
func (*NetworkDNSConfigRecordsItems) ProtoMessage()    {}
func (*NetworkDNSConfigRecordsItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_5bc13885cab21ab6, []int{2}
}
func (m *NetworkDNSConfigRecordsItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkDNSConfigRecordsItems.Unmarshal(m, b)
}
func (m *NetworkDNSConfigRecordsItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkDNSConfigRecordsItems.Marshal(b, m, deterministic)
}
func (dst *NetworkDNSConfigRecordsItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkDNSConfigRecordsItems.Merge(dst, src)
}
func (m *NetworkDNSConfigRecordsItems) XXX_Size() int {
	return xxx_messageInfo_NetworkDNSConfigRecordsItems.Size(m)
}
func (m *NetworkDNSConfigRecordsItems) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkDNSConfigRecordsItems.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkDNSConfigRecordsItems proto.InternalMessageInfo

func (m *NetworkDNSConfigRecordsItems) GetARecord() []string {
	if m != nil {
		return m.ARecord
	}
	return nil
}

func (m *NetworkDNSConfigRecordsItems) GetAaaaRecord() []string {
	if m != nil {
		return m.AaaaRecord
	}
	return nil
}

func (m *NetworkDNSConfigRecordsItems) GetCnameRecord() []string {
	if m != nil {
		return m.CnameRecord
	}
	return nil
}

func (m *NetworkDNSConfigRecordsItems) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type ImageSpec struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Order                int64    `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageSpec) Reset()         { *m = ImageSpec{} }
func (m *ImageSpec) String() string { return proto.CompactTextString(m) }
func (*ImageSpec) ProtoMessage()    {}
func (*ImageSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_5bc13885cab21ab6, []int{3}
}
func (m *ImageSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageSpec.Unmarshal(m, b)
}
func (m *ImageSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageSpec.Marshal(b, m, deterministic)
}
func (dst *ImageSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageSpec.Merge(dst, src)
}
func (m *ImageSpec) XXX_Size() int {
	return xxx_messageInfo_ImageSpec.Size(m)
}
func (m *ImageSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ImageSpec proto.InternalMessageInfo

func (m *ImageSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageSpec) GetOrder() int64 {
	if m != nil {
		return m.Order
	}
	return 0
}

type MagmaD struct {
	LogLevel protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	// Interval for the gateways to send checkin rpc calls to the cloud.
	CheckinInterval int32 `protobuf:"varint,2,opt,name=checkin_interval,json=checkinInterval,proto3" json:"checkin_interval,omitempty"`
	// Checkin rpc timeout
	CheckinTimeout int32 `protobuf:"varint,3,opt,name=checkin_timeout,json=checkinTimeout,proto3" json:"checkin_timeout,omitempty"`
	// Enables autoupgrading of the magma package
	AutoupgradeEnabled bool `protobuf:"varint,4,opt,name=autoupgrade_enabled,json=autoupgradeEnabled,proto3" json:"autoupgrade_enabled,omitempty"`
	// Interval to poll for package upgrades
	AutoupgradePollInterval int32 `protobuf:"varint,5,opt,name=autoupgrade_poll_interval,json=autoupgradePollInterval,proto3" json:"autoupgrade_poll_interval,omitempty"`
	// The magma package version the gateway should upgrade to
	PackageVersion string `protobuf:"bytes,6,opt,name=package_version,json=packageVersion,proto3" json:"package_version,omitempty"`
	// List of upgrade images
	Images []*ImageSpec `protobuf:"bytes,7,rep,name=images,proto3" json:"images,omitempty"`
	// For streamer, should be left unused by gateway
	TierId       string          `protobuf:"bytes,8,opt,name=tier_id,json=tierId,proto3" json:"tier_id,omitempty"`
	FeatureFlags map[string]bool `protobuf:"bytes,9,rep,name=feature_flags,json=featureFlags,proto3" json:"feature_flags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// List of dynamic_services
	DynamicServices      []string `protobuf:"bytes,10,rep,name=dynamic_services,json=dynamicServices,proto3" json:"dynamic_services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MagmaD) Reset()         { *m = MagmaD{} }
func (m *MagmaD) String() string { return proto.CompactTextString(m) }
func (*MagmaD) ProtoMessage()    {}
func (*MagmaD) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_5bc13885cab21ab6, []int{4}
}
func (m *MagmaD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MagmaD.Unmarshal(m, b)
}
func (m *MagmaD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MagmaD.Marshal(b, m, deterministic)
}
func (dst *MagmaD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MagmaD.Merge(dst, src)
}
func (m *MagmaD) XXX_Size() int {
	return xxx_messageInfo_MagmaD.Size(m)
}
func (m *MagmaD) XXX_DiscardUnknown() {
	xxx_messageInfo_MagmaD.DiscardUnknown(m)
}

var xxx_messageInfo_MagmaD proto.InternalMessageInfo

func (m *MagmaD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *MagmaD) GetCheckinInterval() int32 {
	if m != nil {
		return m.CheckinInterval
	}
	return 0
}

func (m *MagmaD) GetCheckinTimeout() int32 {
	if m != nil {
		return m.CheckinTimeout
	}
	return 0
}

func (m *MagmaD) GetAutoupgradeEnabled() bool {
	if m != nil {
		return m.AutoupgradeEnabled
	}
	return false
}

func (m *MagmaD) GetAutoupgradePollInterval() int32 {
	if m != nil {
		return m.AutoupgradePollInterval
	}
	return 0
}

func (m *MagmaD) GetPackageVersion() string {
	if m != nil {
		return m.PackageVersion
	}
	return ""
}

func (m *MagmaD) GetImages() []*ImageSpec {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *MagmaD) GetTierId() string {
	if m != nil {
		return m.TierId
	}
	return ""
}

func (m *MagmaD) GetFeatureFlags() map[string]bool {
	if m != nil {
		return m.FeatureFlags
	}
	return nil
}

func (m *MagmaD) GetDynamicServices() []string {
	if m != nil {
		return m.DynamicServices
	}
	return nil
}

type DirectoryD struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DirectoryD) Reset()         { *m = DirectoryD{} }
func (m *DirectoryD) String() string { return proto.CompactTextString(m) }
func (*DirectoryD) ProtoMessage()    {}
func (*DirectoryD) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_5bc13885cab21ab6, []int{5}
}
func (m *DirectoryD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectoryD.Unmarshal(m, b)
}
func (m *DirectoryD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectoryD.Marshal(b, m, deterministic)
}
func (dst *DirectoryD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectoryD.Merge(dst, src)
}
func (m *DirectoryD) XXX_Size() int {
	return xxx_messageInfo_DirectoryD.Size(m)
}
func (m *DirectoryD) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectoryD.DiscardUnknown(m)
}

var xxx_messageInfo_DirectoryD proto.InternalMessageInfo

func (m *DirectoryD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

// ------------------------------------------------------------------------------
// MetricsD configs
// ------------------------------------------------------------------------------
type MetricsD struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MetricsD) Reset()         { *m = MetricsD{} }
func (m *MetricsD) String() string { return proto.CompactTextString(m) }
func (*MetricsD) ProtoMessage()    {}
func (*MetricsD) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_5bc13885cab21ab6, []int{6}
}
func (m *MetricsD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsD.Unmarshal(m, b)
}
func (m *MetricsD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsD.Marshal(b, m, deterministic)
}
func (dst *MetricsD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsD.Merge(dst, src)
}
func (m *MetricsD) XXX_Size() int {
	return xxx_messageInfo_MetricsD.Size(m)
}
func (m *MetricsD) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsD.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsD proto.InternalMessageInfo

func (m *MetricsD) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

// ------------------------------------------------------------------------------
// State configs
// ------------------------------------------------------------------------------
type State struct {
	LogLevel             protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_5bc13885cab21ab6, []int{7}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func init() {
	proto.RegisterType((*ControlProxy)(nil), "magma.mconfig.ControlProxy")
	proto.RegisterType((*DnsD)(nil), "magma.mconfig.DnsD")
	proto.RegisterType((*NetworkDNSConfigRecordsItems)(nil), "magma.mconfig.NetworkDNSConfigRecordsItems")
	proto.RegisterType((*ImageSpec)(nil), "magma.mconfig.ImageSpec")
	proto.RegisterType((*MagmaD)(nil), "magma.mconfig.MagmaD")
	proto.RegisterMapType((map[string]bool)(nil), "magma.mconfig.MagmaD.FeatureFlagsEntry")
	proto.RegisterType((*DirectoryD)(nil), "magma.mconfig.DirectoryD")
	proto.RegisterType((*MetricsD)(nil), "magma.mconfig.MetricsD")
	proto.RegisterType((*State)(nil), "magma.mconfig.State")
}

func init() {
	proto.RegisterFile("orc8r/protos/mconfig/mconfigs.proto", fileDescriptor_mconfigs_5bc13885cab21ab6)
}

var fileDescriptor_mconfigs_5bc13885cab21ab6 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x6d, 0x4b, 0x1b, 0x41,
	0x10, 0xc7, 0x39, 0xf3, 0x3c, 0x6a, 0xb4, 0xdb, 0x07, 0x4f, 0x29, 0x34, 0x3d, 0x11, 0x53, 0x0a,
	0x49, 0xb1, 0x14, 0xc4, 0x42, 0x5b, 0x34, 0x0a, 0x81, 0x28, 0x72, 0x91, 0xbe, 0xe8, 0x9b, 0x63,
	0xdd, 0x9b, 0x9c, 0x4b, 0xf6, 0x6e, 0xc3, 0xde, 0x26, 0x6d, 0x3e, 0x48, 0x3f, 0x45, 0xbf, 0x46,
	0x3f, 0x58, 0xb9, 0xdd, 0x8d, 0x8d, 0x16, 0xfa, 0x22, 0xaf, 0xb2, 0xf3, 0x9b, 0xff, 0x4c, 0x76,
	0x67, 0xe6, 0x06, 0xf6, 0xa5, 0x62, 0xc7, 0xaa, 0x3b, 0x51, 0x52, 0xcb, 0xbc, 0x9b, 0x32, 0x99,
	0x8d, 0x78, 0xb2, 0xf8, 0xcd, 0x3b, 0x86, 0x93, 0xcd, 0x94, 0x26, 0x29, 0xed, 0x38, 0xba, 0xb7,
	0xfb, 0x20, 0x86, 0xc9, 0x34, 0x95, 0x99, 0x55, 0x06, 0xa7, 0xb0, 0x71, 0x26, 0x33, 0xad, 0xa4,
	0xb8, 0x56, 0xf2, 0xc7, 0x9c, 0x1c, 0x41, 0x43, 0xc8, 0x24, 0x12, 0x38, 0x43, 0xe1, 0x7b, 0x2d,
	0xaf, 0xdd, 0x3c, 0x7a, 0xde, 0xb1, 0xd9, 0x4c, 0x92, 0xce, 0x40, 0x26, 0x83, 0xc2, 0x19, 0xd6,
	0x85, 0x3b, 0x05, 0xbf, 0x3d, 0x28, 0xf7, 0xb2, 0xbc, 0xb7, 0x4a, 0x30, 0x39, 0x80, 0x26, 0x66,
	0xf4, 0x56, 0x60, 0xc4, 0x28, 0xbb, 0xe3, 0x59, 0xe2, 0xaf, 0xb5, 0xbc, 0x76, 0x3d, 0xdc, 0xb4,
	0xf4, 0xcc, 0x42, 0xb2, 0x07, 0x75, 0x21, 0x19, 0x15, 0x37, 0x37, 0x03, 0xbf, 0xd4, 0xf2, 0xda,
	0x95, 0xf0, 0xde, 0x26, 0xe7, 0x50, 0x53, 0xc8, 0xa4, 0x8a, 0x73, 0xbf, 0xdc, 0x2a, 0xb5, 0xd7,
	0x8f, 0xde, 0x76, 0x1e, 0xbc, 0xbf, 0x73, 0x85, 0xfa, 0xbb, 0x54, 0xe3, 0xde, 0xd5, 0xf0, 0xcc,
	0x80, 0xd0, 0xaa, 0xfb, 0x1a, 0xd3, 0x3c, 0x5c, 0xc4, 0x06, 0x3f, 0x3d, 0x78, 0xf9, 0x3f, 0x25,
	0xd9, 0x85, 0x3a, 0x8d, 0xac, 0xda, 0xf7, 0x5a, 0xa5, 0x76, 0x23, 0xac, 0x51, 0x2b, 0x20, 0xaf,
	0x60, 0x9d, 0x52, 0x7a, 0xef, 0x5d, 0x33, 0x5e, 0x28, 0x90, 0x13, 0xbc, 0x86, 0x0d, 0x96, 0xd1,
	0x14, 0x17, 0x8a, 0x92, 0x51, 0xac, 0x1b, 0xe6, 0x24, 0x2f, 0xa0, 0x1a, 0xcb, 0x94, 0xf2, 0xcc,
	0x2f, 0xb7, 0xbc, 0x76, 0x23, 0x74, 0x56, 0xf0, 0x01, 0x1a, 0xfd, 0x94, 0x26, 0x38, 0x9c, 0x20,
	0x23, 0x04, 0xca, 0x45, 0x88, 0xa9, 0x6e, 0x23, 0x34, 0x67, 0xf2, 0x0c, 0x2a, 0x52, 0xc5, 0xa8,
	0x4c, 0xe5, 0x4a, 0xa1, 0x35, 0x82, 0x5f, 0x65, 0xa8, 0x5e, 0x16, 0x65, 0x58, 0xad, 0x2f, 0x6f,
	0x60, 0x9b, 0xdd, 0x21, 0x1b, 0xf3, 0x2c, 0xe2, 0x99, 0x46, 0x35, 0xa3, 0xc2, 0xe4, 0xaf, 0x84,
	0x5b, 0x8e, 0xf7, 0x1d, 0x26, 0x87, 0xb0, 0x40, 0x91, 0xe6, 0x29, 0xca, 0xa9, 0x76, 0x2d, 0x6a,
	0x3a, 0x7c, 0x63, 0x29, 0xe9, 0xc2, 0x53, 0x3a, 0xd5, 0x72, 0x3a, 0x49, 0x14, 0x8d, 0x31, 0xb2,
	0x1d, 0x8e, 0xcd, 0x73, 0xeb, 0x21, 0x59, 0x72, 0x9d, 0x5b, 0x0f, 0x39, 0x81, 0xdd, 0xe5, 0x80,
	0x89, 0x14, 0xe2, 0xef, 0x6d, 0x2a, 0xe6, 0x3f, 0x76, 0x96, 0x04, 0xd7, 0x52, 0x88, 0xe5, 0x5b,
	0x4d, 0x28, 0x1b, 0xd3, 0x04, 0xa3, 0x19, 0xaa, 0x9c, 0xcb, 0xcc, 0xaf, 0x9a, 0xa2, 0x35, 0x1d,
	0xfe, 0x6a, 0x29, 0x79, 0x07, 0x55, 0x5e, 0xd4, 0x37, 0xf7, 0x6b, 0x66, 0x7a, 0xfc, 0x47, 0xd3,
	0x73, 0x5f, 0xfc, 0xd0, 0xe9, 0xc8, 0x0e, 0xd4, 0x34, 0x47, 0x15, 0xf1, 0xd8, 0xaf, 0xdb, 0x56,
	0x15, 0x66, 0x3f, 0x26, 0x03, 0xd8, 0x1c, 0x21, 0xd5, 0x53, 0x85, 0xd1, 0x48, 0xd0, 0x24, 0xf7,
	0x1b, 0x26, 0xe3, 0xe1, 0xa3, 0x8c, 0xb6, 0x2d, 0x9d, 0x0b, 0x2b, 0xbd, 0x28, 0x94, 0xe7, 0x99,
	0x56, 0xf3, 0x70, 0x63, 0xb4, 0x84, 0x8a, 0x16, 0xc4, 0xf3, 0x8c, 0xa6, 0x9c, 0x45, 0x39, 0xaa,
	0x19, 0x67, 0x98, 0xfb, 0x60, 0xe6, 0x66, 0xcb, 0xf1, 0xa1, 0xc3, 0x7b, 0x9f, 0xe1, 0xc9, 0x3f,
	0xd9, 0xc8, 0x36, 0x94, 0xc6, 0x38, 0x77, 0xa3, 0x52, 0x1c, 0x8b, 0x49, 0x99, 0x51, 0x31, 0x45,
	0xf7, 0x8d, 0x59, 0xe3, 0x64, 0xed, 0xd8, 0x0b, 0xbe, 0x00, 0xf4, 0xb8, 0x42, 0xa6, 0xa5, 0x9a,
	0xaf, 0x34, 0x30, 0xc1, 0x27, 0xa8, 0x5f, 0xa2, 0x56, 0x9c, 0xad, 0xb6, 0x08, 0x82, 0x8f, 0x50,
	0x19, 0x6a, 0xaa, 0x71, 0x95, 0xe0, 0xd3, 0x83, 0x6f, 0xfb, 0x46, 0xd1, 0xb5, 0x9b, 0x8e, 0x09,
	0x39, 0x8d, 0xbb, 0x89, 0x7c, 0xb4, 0x26, 0x6f, 0xab, 0xc6, 0x7e, 0xff, 0x27, 0x00, 0x00, 0xff,
	0xff, 0xc1, 0x57, 0x22, 0xa5, 0x45, 0x05, 0x00, 0x00,
}
