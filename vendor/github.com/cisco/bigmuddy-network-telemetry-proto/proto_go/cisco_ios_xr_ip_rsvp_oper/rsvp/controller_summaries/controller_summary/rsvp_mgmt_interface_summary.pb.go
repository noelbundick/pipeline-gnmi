// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_mgmt_interface_summary.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_controller_summaries_controller_summary

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Summary Form of RSVP interface Info
type RsvpMgmtInterfaceSummary_KEYS struct {
	ControllerName       string   `protobuf:"bytes,1,opt,name=controller_name,json=controllerName" json:"controller_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtInterfaceSummary_KEYS) Reset()         { *m = RsvpMgmtInterfaceSummary_KEYS{} }
func (m *RsvpMgmtInterfaceSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtInterfaceSummary_KEYS) ProtoMessage()    {}
func (*RsvpMgmtInterfaceSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_interface_summary_6e585e968840893d, []int{0}
}
func (m *RsvpMgmtInterfaceSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtInterfaceSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtInterfaceSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS.Merge(dst, src)
}
func (m *RsvpMgmtInterfaceSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS.Size(m)
}
func (m *RsvpMgmtInterfaceSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtInterfaceSummary_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtInterfaceSummary_KEYS) GetControllerName() string {
	if m != nil {
		return m.ControllerName
	}
	return ""
}

type RsvpMgmtInterfaceSummary struct {
	// Interface Name
	InterfaceName string `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Interface bandwidth information
	BandwidthInformation *RsvpMgmtDsteModeInterfaceBw `protobuf:"bytes,51,opt,name=bandwidth_information,json=bandwidthInformation" json:"bandwidth_information,omitempty"`
	// Number of locally created and incoming path states
	PathsIn uint32 `protobuf:"varint,52,opt,name=paths_in,json=pathsIn" json:"paths_in,omitempty"`
	// Number of outgoing path states
	PathsOut uint32 `protobuf:"varint,53,opt,name=paths_out,json=pathsOut" json:"paths_out,omitempty"`
	// Number of locally created and incoming reservation states
	ReservationsIn uint32 `protobuf:"varint,54,opt,name=reservations_in,json=reservationsIn" json:"reservations_in,omitempty"`
	// Number of outgoing reservation states
	ReservationsOut      uint32   `protobuf:"varint,55,opt,name=reservations_out,json=reservationsOut" json:"reservations_out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtInterfaceSummary) Reset()         { *m = RsvpMgmtInterfaceSummary{} }
func (m *RsvpMgmtInterfaceSummary) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtInterfaceSummary) ProtoMessage()    {}
func (*RsvpMgmtInterfaceSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_interface_summary_6e585e968840893d, []int{1}
}
func (m *RsvpMgmtInterfaceSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary.Unmarshal(m, b)
}
func (m *RsvpMgmtInterfaceSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtInterfaceSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtInterfaceSummary.Merge(dst, src)
}
func (m *RsvpMgmtInterfaceSummary) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtInterfaceSummary.Size(m)
}
func (m *RsvpMgmtInterfaceSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtInterfaceSummary.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtInterfaceSummary proto.InternalMessageInfo

func (m *RsvpMgmtInterfaceSummary) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *RsvpMgmtInterfaceSummary) GetBandwidthInformation() *RsvpMgmtDsteModeInterfaceBw {
	if m != nil {
		return m.BandwidthInformation
	}
	return nil
}

func (m *RsvpMgmtInterfaceSummary) GetPathsIn() uint32 {
	if m != nil {
		return m.PathsIn
	}
	return 0
}

func (m *RsvpMgmtInterfaceSummary) GetPathsOut() uint32 {
	if m != nil {
		return m.PathsOut
	}
	return 0
}

func (m *RsvpMgmtInterfaceSummary) GetReservationsIn() uint32 {
	if m != nil {
		return m.ReservationsIn
	}
	return 0
}

func (m *RsvpMgmtInterfaceSummary) GetReservationsOut() uint32 {
	if m != nil {
		return m.ReservationsOut
	}
	return 0
}

// Prestandard DSTE interface bandwidth information
type RsvpMgmtInterfaceBwPrestdDste struct {
	// Bandwidth (bits per second) now allocated
	AllocatedBitRate uint64 `protobuf:"varint,1,opt,name=allocated_bit_rate,json=allocatedBitRate" json:"allocated_bit_rate,omitempty"`
	// Max bandwidth (bits per second) allowed per flow
	MaxFlowBandwidth uint64 `protobuf:"varint,2,opt,name=max_flow_bandwidth,json=maxFlowBandwidth" json:"max_flow_bandwidth,omitempty"`
	// Max bandwidth (bits per second) allowed
	MaxBandwidth uint64 `protobuf:"varint,3,opt,name=max_bandwidth,json=maxBandwidth" json:"max_bandwidth,omitempty"`
	// Max bandwidth (bits per second) allowed in subpool
	MaxSubpoolBandwidth uint64 `protobuf:"varint,4,opt,name=max_subpool_bandwidth,json=maxSubpoolBandwidth" json:"max_subpool_bandwidth,omitempty"`
	// True if the Max B/W is an absolute value and false if its a percentage
	IsMaxBandwidthAbsolute bool `protobuf:"varint,5,opt,name=is_max_bandwidth_absolute,json=isMaxBandwidthAbsolute" json:"is_max_bandwidth_absolute,omitempty"`
	// True if the Max sub-pool B/W is an absolute value and false if its a percentage
	IsMaxSubpoolBandwidthAbsolute bool     `protobuf:"varint,6,opt,name=is_max_subpool_bandwidth_absolute,json=isMaxSubpoolBandwidthAbsolute" json:"is_max_subpool_bandwidth_absolute,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *RsvpMgmtInterfaceBwPrestdDste) Reset()         { *m = RsvpMgmtInterfaceBwPrestdDste{} }
func (m *RsvpMgmtInterfaceBwPrestdDste) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtInterfaceBwPrestdDste) ProtoMessage()    {}
func (*RsvpMgmtInterfaceBwPrestdDste) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_interface_summary_6e585e968840893d, []int{2}
}
func (m *RsvpMgmtInterfaceBwPrestdDste) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste.Unmarshal(m, b)
}
func (m *RsvpMgmtInterfaceBwPrestdDste) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtInterfaceBwPrestdDste) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste.Merge(dst, src)
}
func (m *RsvpMgmtInterfaceBwPrestdDste) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste.Size(m)
}
func (m *RsvpMgmtInterfaceBwPrestdDste) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtInterfaceBwPrestdDste proto.InternalMessageInfo

func (m *RsvpMgmtInterfaceBwPrestdDste) GetAllocatedBitRate() uint64 {
	if m != nil {
		return m.AllocatedBitRate
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwPrestdDste) GetMaxFlowBandwidth() uint64 {
	if m != nil {
		return m.MaxFlowBandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwPrestdDste) GetMaxBandwidth() uint64 {
	if m != nil {
		return m.MaxBandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwPrestdDste) GetMaxSubpoolBandwidth() uint64 {
	if m != nil {
		return m.MaxSubpoolBandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwPrestdDste) GetIsMaxBandwidthAbsolute() bool {
	if m != nil {
		return m.IsMaxBandwidthAbsolute
	}
	return false
}

func (m *RsvpMgmtInterfaceBwPrestdDste) GetIsMaxSubpoolBandwidthAbsolute() bool {
	if m != nil {
		return m.IsMaxSubpoolBandwidthAbsolute
	}
	return false
}

// RSVP interface bandwidth info with standard based DSTE enabled
type RsvpMgmtInterfaceBwStdDste struct {
	// Bandwidth (bits per second) now allocated
	AllocatedBitRate uint64 `protobuf:"varint,1,opt,name=allocated_bit_rate,json=allocatedBitRate" json:"allocated_bit_rate,omitempty"`
	// Max bandwidth (bits per second) allowed per flow
	MaxFlowBandwidth uint64 `protobuf:"varint,2,opt,name=max_flow_bandwidth,json=maxFlowBandwidth" json:"max_flow_bandwidth,omitempty"`
	// Max bandwidth (bits per second) allowed
	MaxBandwidth uint64 `protobuf:"varint,3,opt,name=max_bandwidth,json=maxBandwidth" json:"max_bandwidth,omitempty"`
	// Max bandwidth (bits per second) allowed in BC0 pool
	MaxPool0Bandwidth uint64 `protobuf:"varint,4,opt,name=max_pool0_bandwidth,json=maxPool0Bandwidth" json:"max_pool0_bandwidth,omitempty"`
	// Max bandwidth (bits per second) allowed in BC1 pool
	MaxPool1Bandwidth uint64 `protobuf:"varint,5,opt,name=max_pool1_bandwidth,json=maxPool1Bandwidth" json:"max_pool1_bandwidth,omitempty"`
	// True if the Max B/W is an absolute value and false if its a percentage
	IsMaxBandwidthAbsolute bool `protobuf:"varint,6,opt,name=is_max_bandwidth_absolute,json=isMaxBandwidthAbsolute" json:"is_max_bandwidth_absolute,omitempty"`
	// True if the Max BC0 B/W is an absolute value and false if its a percentage
	IsMaxBc0BandwidthAbsolute bool `protobuf:"varint,7,opt,name=is_max_bc0_bandwidth_absolute,json=isMaxBc0BandwidthAbsolute" json:"is_max_bc0_bandwidth_absolute,omitempty"`
	// True if the Max BC1 sub-pool B/W is an absolute value and false if its a percentage
	IsMaxBc1BandwidthAbsolute bool     `protobuf:"varint,8,opt,name=is_max_bc1_bandwidth_absolute,json=isMaxBc1BandwidthAbsolute" json:"is_max_bc1_bandwidth_absolute,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *RsvpMgmtInterfaceBwStdDste) Reset()         { *m = RsvpMgmtInterfaceBwStdDste{} }
func (m *RsvpMgmtInterfaceBwStdDste) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtInterfaceBwStdDste) ProtoMessage()    {}
func (*RsvpMgmtInterfaceBwStdDste) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_interface_summary_6e585e968840893d, []int{3}
}
func (m *RsvpMgmtInterfaceBwStdDste) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtInterfaceBwStdDste.Unmarshal(m, b)
}
func (m *RsvpMgmtInterfaceBwStdDste) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtInterfaceBwStdDste.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtInterfaceBwStdDste) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtInterfaceBwStdDste.Merge(dst, src)
}
func (m *RsvpMgmtInterfaceBwStdDste) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtInterfaceBwStdDste.Size(m)
}
func (m *RsvpMgmtInterfaceBwStdDste) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtInterfaceBwStdDste.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtInterfaceBwStdDste proto.InternalMessageInfo

func (m *RsvpMgmtInterfaceBwStdDste) GetAllocatedBitRate() uint64 {
	if m != nil {
		return m.AllocatedBitRate
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwStdDste) GetMaxFlowBandwidth() uint64 {
	if m != nil {
		return m.MaxFlowBandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwStdDste) GetMaxBandwidth() uint64 {
	if m != nil {
		return m.MaxBandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwStdDste) GetMaxPool0Bandwidth() uint64 {
	if m != nil {
		return m.MaxPool0Bandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwStdDste) GetMaxPool1Bandwidth() uint64 {
	if m != nil {
		return m.MaxPool1Bandwidth
	}
	return 0
}

func (m *RsvpMgmtInterfaceBwStdDste) GetIsMaxBandwidthAbsolute() bool {
	if m != nil {
		return m.IsMaxBandwidthAbsolute
	}
	return false
}

func (m *RsvpMgmtInterfaceBwStdDste) GetIsMaxBc0BandwidthAbsolute() bool {
	if m != nil {
		return m.IsMaxBc0BandwidthAbsolute
	}
	return false
}

func (m *RsvpMgmtInterfaceBwStdDste) GetIsMaxBc1BandwidthAbsolute() bool {
	if m != nil {
		return m.IsMaxBc1BandwidthAbsolute
	}
	return false
}

// Union of the different RSVP interface bandwidth types
type RsvpMgmtDsteModeInterfaceBw struct {
	DsteMode string `protobuf:"bytes,1,opt,name=dste_mode,json=dsteMode" json:"dste_mode,omitempty"`
	// Prestandard DSTE interface information
	PreStandardDsteInterface *RsvpMgmtInterfaceBwPrestdDste `protobuf:"bytes,2,opt,name=pre_standard_dste_interface,json=preStandardDsteInterface" json:"pre_standard_dste_interface,omitempty"`
	// Standard DSTE interface information
	StandardDsteInterface *RsvpMgmtInterfaceBwStdDste `protobuf:"bytes,3,opt,name=standard_dste_interface,json=standardDsteInterface" json:"standard_dste_interface,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                    `json:"-"`
	XXX_unrecognized      []byte                      `json:"-"`
	XXX_sizecache         int32                       `json:"-"`
}

func (m *RsvpMgmtDsteModeInterfaceBw) Reset()         { *m = RsvpMgmtDsteModeInterfaceBw{} }
func (m *RsvpMgmtDsteModeInterfaceBw) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtDsteModeInterfaceBw) ProtoMessage()    {}
func (*RsvpMgmtDsteModeInterfaceBw) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_interface_summary_6e585e968840893d, []int{4}
}
func (m *RsvpMgmtDsteModeInterfaceBw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw.Unmarshal(m, b)
}
func (m *RsvpMgmtDsteModeInterfaceBw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtDsteModeInterfaceBw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw.Merge(dst, src)
}
func (m *RsvpMgmtDsteModeInterfaceBw) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw.Size(m)
}
func (m *RsvpMgmtDsteModeInterfaceBw) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtDsteModeInterfaceBw proto.InternalMessageInfo

func (m *RsvpMgmtDsteModeInterfaceBw) GetDsteMode() string {
	if m != nil {
		return m.DsteMode
	}
	return ""
}

func (m *RsvpMgmtDsteModeInterfaceBw) GetPreStandardDsteInterface() *RsvpMgmtInterfaceBwPrestdDste {
	if m != nil {
		return m.PreStandardDsteInterface
	}
	return nil
}

func (m *RsvpMgmtDsteModeInterfaceBw) GetStandardDsteInterface() *RsvpMgmtInterfaceBwStdDste {
	if m != nil {
		return m.StandardDsteInterface
	}
	return nil
}

func init() {
	proto.RegisterType((*RsvpMgmtInterfaceSummary_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.controller_summaries.controller_summary.rsvp_mgmt_interface_summary_KEYS")
	proto.RegisterType((*RsvpMgmtInterfaceSummary)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.controller_summaries.controller_summary.rsvp_mgmt_interface_summary")
	proto.RegisterType((*RsvpMgmtInterfaceBwPrestdDste)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.controller_summaries.controller_summary.rsvp_mgmt_interface_bw_prestd_dste")
	proto.RegisterType((*RsvpMgmtInterfaceBwStdDste)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.controller_summaries.controller_summary.rsvp_mgmt_interface_bw_std_dste")
	proto.RegisterType((*RsvpMgmtDsteModeInterfaceBw)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.controller_summaries.controller_summary.rsvp_mgmt_dste_mode_interface_bw")
}

func init() {
	proto.RegisterFile("rsvp_mgmt_interface_summary.proto", fileDescriptor_rsvp_mgmt_interface_summary_6e585e968840893d)
}

var fileDescriptor_rsvp_mgmt_interface_summary_6e585e968840893d = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xdf, 0x6e, 0xd3, 0x3c,
	0x18, 0xc6, 0x95, 0x75, 0xeb, 0x3a, 0xef, 0x6b, 0xbf, 0x61, 0x28, 0xa4, 0xaa, 0x26, 0xba, 0x22,
	0x44, 0x91, 0x50, 0xd4, 0x76, 0xfc, 0x11, 0x67, 0x30, 0xc1, 0x44, 0x35, 0x8d, 0xa1, 0xf4, 0x88,
	0x23, 0xcb, 0x49, 0xdc, 0xd5, 0x28, 0x8e, 0x23, 0xdb, 0x59, 0xb3, 0x8b, 0xe0, 0x12, 0x10, 0xc7,
	0x5c, 0x06, 0x37, 0xc2, 0x35, 0x70, 0x09, 0x28, 0x4e, 0x9b, 0xa4, 0x34, 0xad, 0x38, 0x18, 0x12,
	0x67, 0xdb, 0xfb, 0xfe, 0x9e, 0xc7, 0x4f, 0xde, 0x57, 0x76, 0xc1, 0x91, 0x90, 0x57, 0x21, 0x62,
	0x97, 0x4c, 0x21, 0x1a, 0x28, 0x22, 0x26, 0xd8, 0x25, 0x48, 0x46, 0x8c, 0x61, 0x71, 0x6d, 0x85,
	0x82, 0x2b, 0x0e, 0x4f, 0x5d, 0x2a, 0x5d, 0x8e, 0x28, 0x97, 0x28, 0x16, 0x88, 0x86, 0x48, 0x4b,
	0x78, 0x48, 0x84, 0x95, 0xfc, 0x65, 0xb9, 0x3c, 0x50, 0x82, 0xfb, 0x3e, 0x11, 0x73, 0x21, 0x25,
	0x72, 0xb5, 0x78, 0xdd, 0x3d, 0x03, 0x9d, 0x0d, 0x87, 0xa1, 0xb3, 0xb7, 0x1f, 0xc7, 0xf0, 0x11,
	0xf8, 0xbf, 0xa0, 0x0c, 0x30, 0x23, 0xa6, 0xd1, 0x31, 0x7a, 0x7b, 0x76, 0x23, 0x2f, 0xbf, 0xc7,
	0x8c, 0x74, 0x7f, 0x6e, 0x81, 0xf6, 0x06, 0x37, 0xf8, 0x10, 0x34, 0xf2, 0xa2, 0xf6, 0x19, 0x6a,
	0x9f, 0x7a, 0x56, 0x4d, 0x6c, 0xe0, 0x17, 0x03, 0x34, 0x1d, 0x1c, 0x78, 0x33, 0xea, 0xa9, 0x29,
	0xa2, 0xc1, 0x84, 0x0b, 0x86, 0x15, 0xe5, 0x81, 0x79, 0xdc, 0x31, 0x7a, 0xfb, 0xc3, 0xa9, 0x75,
	0x33, 0x1f, 0x6f, 0xe5, 0x59, 0x3d, 0xa9, 0x08, 0x62, 0xdc, 0x23, 0x85, 0xd4, 0xce, 0xcc, 0xbe,
	0x93, 0xc5, 0x18, 0xe5, 0x29, 0x60, 0x0b, 0xd4, 0x42, 0xac, 0xa6, 0x12, 0xd1, 0xc0, 0x7c, 0xda,
	0x31, 0x7a, 0x75, 0x7b, 0x57, 0xff, 0x3f, 0x0a, 0x60, 0x1b, 0xec, 0xa5, 0x2d, 0x1e, 0x29, 0xf3,
	0x99, 0xee, 0xa5, 0xec, 0x45, 0xa4, 0x92, 0x39, 0x0a, 0x22, 0x89, 0xb8, 0xd2, 0x36, 0x5a, 0xfe,
	0x5c, 0x23, 0x8d, 0x62, 0x79, 0x14, 0xc0, 0xc7, 0xe0, 0x60, 0x09, 0x4c, 0xcc, 0x5e, 0x68, 0x72,
	0xc9, 0xe0, 0x22, 0x52, 0xdd, 0x1f, 0x5b, 0xa0, 0x5b, 0x36, 0x72, 0x67, 0x86, 0x42, 0x41, 0xa4,
	0xf2, 0xf4, 0xa7, 0xc1, 0x27, 0x00, 0x62, 0xdf, 0xe7, 0x2e, 0x56, 0xc4, 0x43, 0x0e, 0x55, 0x48,
	0x60, 0x95, 0x6e, 0x71, 0xdb, 0x3e, 0xc8, 0x3a, 0x27, 0x54, 0xd9, 0x38, 0xa5, 0x19, 0x8e, 0xd1,
	0xc4, 0xe7, 0x33, 0x94, 0x4d, 0xc0, 0xdc, 0x4a, 0x69, 0x86, 0xe3, 0x53, 0x9f, 0xcf, 0x4e, 0x16,
	0x75, 0xf8, 0x00, 0xd4, 0x13, 0x3a, 0x07, 0x2b, 0x1a, 0xfc, 0x8f, 0xe1, 0x38, 0x87, 0x86, 0xa0,
	0x99, 0x40, 0x32, 0x72, 0x42, 0xce, 0xfd, 0x02, 0xbc, 0xad, 0xe1, 0xdb, 0x0c, 0xc7, 0xe3, 0xb4,
	0x97, 0x6b, 0x5e, 0x82, 0x16, 0x95, 0x68, 0xc9, 0x1b, 0x61, 0x47, 0x72, 0x3f, 0x52, 0xc4, 0xdc,
	0xe9, 0x18, 0xbd, 0x9a, 0x7d, 0x97, 0xca, 0xf3, 0xc2, 0x31, 0xaf, 0xe7, 0x5d, 0xf8, 0x0e, 0x1c,
	0xcd, 0xa5, 0x2b, 0x27, 0xe6, 0x16, 0x55, 0x6d, 0x71, 0xa8, 0x2d, 0x7e, 0x3f, 0x7c, 0xe1, 0xd4,
	0xfd, 0x5e, 0x01, 0xf7, 0xd7, 0x0c, 0xf8, 0xdf, 0x99, 0xae, 0x05, 0x92, 0x01, 0xa2, 0xe4, 0x0b,
	0xfa, 0x2b, 0xb3, 0xbd, 0xc5, 0x70, 0xfc, 0x21, 0xe9, 0x94, 0xf2, 0x83, 0x02, 0xbf, 0xb3, 0xc4,
	0x0f, 0xfe, 0x70, 0x13, 0xd5, 0x8d, 0x9b, 0x78, 0x05, 0x0e, 0x17, 0x52, 0xb7, 0x5f, 0x26, 0xdf,
	0xd5, 0xf2, 0x56, 0x2a, 0x77, 0xfb, 0x9b, 0x1d, 0x06, 0x65, 0x0e, 0xb5, 0x25, 0x87, 0xc1, 0xea,
	0x0e, 0x3f, 0x57, 0x8a, 0xaf, 0x5c, 0xf9, 0x5d, 0x4f, 0xae, 0x6e, 0xd6, 0x99, 0xbf, 0x6f, 0xb5,
	0xa4, 0x70, 0xce, 0x3d, 0x02, 0xbf, 0x19, 0xa0, 0x1d, 0x0a, 0x82, 0xa4, 0xc2, 0x81, 0x87, 0x45,
	0xba, 0xf7, 0x5c, 0xaf, 0xb7, 0xb7, 0x3f, 0xfc, 0x74, 0xf3, 0x0f, 0xd3, 0xba, 0x1b, 0x6d, 0x9b,
	0xa1, 0x20, 0xe3, 0x79, 0x9a, 0x37, 0x52, 0x91, 0xd1, 0x02, 0x84, 0x5f, 0x0d, 0x70, 0x6f, 0x5d,
	0xce, 0x8a, 0xce, 0x79, 0xf9, 0x97, 0x73, 0x66, 0x21, 0x9b, 0xb2, 0x2c, 0xa1, 0x53, 0xd5, 0xbf,
	0x61, 0xc7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x13, 0x44, 0xcc, 0xbe, 0xe8, 0x06, 0x00, 0x00,
}