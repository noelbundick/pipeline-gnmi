// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_tp_detail_mid_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_tp_tp_midpoints_tp_detail_midpoints_tp_detail_midpoint

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

// Transport profile mid-point LSP detail information
type MplsTpDetailMidInfo_KEYS struct {
	TpMidpointName       string   `protobuf:"bytes,1,opt,name=tp_midpoint_name,json=tpMidpointName" json:"tp_midpoint_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTpDetailMidInfo_KEYS) Reset()         { *m = MplsTpDetailMidInfo_KEYS{} }
func (m *MplsTpDetailMidInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTpDetailMidInfo_KEYS) ProtoMessage()    {}
func (*MplsTpDetailMidInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_tp_detail_mid_info_22e0bb4fdf384efe, []int{0}
}
func (m *MplsTpDetailMidInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTpDetailMidInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTpDetailMidInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTpDetailMidInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTpDetailMidInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTpDetailMidInfo_KEYS.Merge(dst, src)
}
func (m *MplsTpDetailMidInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTpDetailMidInfo_KEYS.Size(m)
}
func (m *MplsTpDetailMidInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTpDetailMidInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTpDetailMidInfo_KEYS proto.InternalMessageInfo

func (m *MplsTpDetailMidInfo_KEYS) GetTpMidpointName() string {
	if m != nil {
		return m.TpMidpointName
	}
	return ""
}

type MplsTpDetailMidInfo struct {
	// Brief LSP information
	BriefLspInformation *MplsTpMidInfo `protobuf:"bytes,50,opt,name=brief_lsp_information,json=briefLspInformation" json:"brief_lsp_information,omitempty"`
	// Forward LSP information
	ForwardLsp *MplsTpMidLspInfo `protobuf:"bytes,51,opt,name=forward_lsp,json=forwardLsp" json:"forward_lsp,omitempty"`
	// Reverse LSP information
	ReverseLsp *MplsTpMidLspInfo `protobuf:"bytes,52,opt,name=reverse_lsp,json=reverseLsp" json:"reverse_lsp,omitempty"`
	// Forward LSP reserved bandwidth in Kbps
	ForwardLspReservedBandwidth uint32 `protobuf:"varint,53,opt,name=forward_lsp_reserved_bandwidth,json=forwardLspReservedBandwidth" json:"forward_lsp_reserved_bandwidth,omitempty"`
	// Reverse LSP reserved bandwidth in Kbps
	ReverseLspReservedBandwidth uint32   `protobuf:"varint,54,opt,name=reverse_lsp_reserved_bandwidth,json=reverseLspReservedBandwidth" json:"reverse_lsp_reserved_bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *MplsTpDetailMidInfo) Reset()         { *m = MplsTpDetailMidInfo{} }
func (m *MplsTpDetailMidInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTpDetailMidInfo) ProtoMessage()    {}
func (*MplsTpDetailMidInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_tp_detail_mid_info_22e0bb4fdf384efe, []int{1}
}
func (m *MplsTpDetailMidInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTpDetailMidInfo.Unmarshal(m, b)
}
func (m *MplsTpDetailMidInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTpDetailMidInfo.Marshal(b, m, deterministic)
}
func (dst *MplsTpDetailMidInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTpDetailMidInfo.Merge(dst, src)
}
func (m *MplsTpDetailMidInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTpDetailMidInfo.Size(m)
}
func (m *MplsTpDetailMidInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTpDetailMidInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTpDetailMidInfo proto.InternalMessageInfo

func (m *MplsTpDetailMidInfo) GetBriefLspInformation() *MplsTpMidInfo {
	if m != nil {
		return m.BriefLspInformation
	}
	return nil
}

func (m *MplsTpDetailMidInfo) GetForwardLsp() *MplsTpMidLspInfo {
	if m != nil {
		return m.ForwardLsp
	}
	return nil
}

func (m *MplsTpDetailMidInfo) GetReverseLsp() *MplsTpMidLspInfo {
	if m != nil {
		return m.ReverseLsp
	}
	return nil
}

func (m *MplsTpDetailMidInfo) GetForwardLspReservedBandwidth() uint32 {
	if m != nil {
		return m.ForwardLspReservedBandwidth
	}
	return 0
}

func (m *MplsTpDetailMidInfo) GetReverseLspReservedBandwidth() uint32 {
	if m != nil {
		return m.ReverseLspReservedBandwidth
	}
	return 0
}

// Transport profile mid-point LSP detailed information
type MplsTpMidLspInfo struct {
	// Outgoing label
	OutLabel uint32 `protobuf:"varint,1,opt,name=out_label,json=outLabel" json:"out_label,omitempty"`
	// Outgoing TP link ID
	OutLink uint32 `protobuf:"varint,2,opt,name=out_link,json=outLink" json:"out_link,omitempty"`
	// OAM Refresh time in seconds
	OamRefreshInterval uint32 `protobuf:"varint,3,opt,name=oam_refresh_interval,json=oamRefreshInterval" json:"oam_refresh_interval,omitempty"`
	// Outgoing interface
	OutgoingInterface string `protobuf:"bytes,4,opt,name=outgoing_interface,json=outgoingInterface" json:"outgoing_interface,omitempty"`
	// Outgoing next-hop in IPv4 address format
	NextHopAddress string `protobuf:"bytes,5,opt,name=next_hop_address,json=nextHopAddress" json:"next_hop_address,omitempty"`
	// Incoming label
	InLabel uint32 `protobuf:"varint,6,opt,name=in_label,json=inLabel" json:"in_label,omitempty"`
	// Configured bandwidth in Kbps
	Bandwidth            uint32   `protobuf:"varint,7,opt,name=bandwidth" json:"bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTpMidLspInfo) Reset()         { *m = MplsTpMidLspInfo{} }
func (m *MplsTpMidLspInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTpMidLspInfo) ProtoMessage()    {}
func (*MplsTpMidLspInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_tp_detail_mid_info_22e0bb4fdf384efe, []int{2}
}
func (m *MplsTpMidLspInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTpMidLspInfo.Unmarshal(m, b)
}
func (m *MplsTpMidLspInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTpMidLspInfo.Marshal(b, m, deterministic)
}
func (dst *MplsTpMidLspInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTpMidLspInfo.Merge(dst, src)
}
func (m *MplsTpMidLspInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTpMidLspInfo.Size(m)
}
func (m *MplsTpMidLspInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTpMidLspInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTpMidLspInfo proto.InternalMessageInfo

func (m *MplsTpMidLspInfo) GetOutLabel() uint32 {
	if m != nil {
		return m.OutLabel
	}
	return 0
}

func (m *MplsTpMidLspInfo) GetOutLink() uint32 {
	if m != nil {
		return m.OutLink
	}
	return 0
}

func (m *MplsTpMidLspInfo) GetOamRefreshInterval() uint32 {
	if m != nil {
		return m.OamRefreshInterval
	}
	return 0
}

func (m *MplsTpMidLspInfo) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *MplsTpMidLspInfo) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *MplsTpMidLspInfo) GetInLabel() uint32 {
	if m != nil {
		return m.InLabel
	}
	return 0
}

func (m *MplsTpMidLspInfo) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

// Transport profile mid-point LSP information
type MplsTpMidInfo struct {
	// Mid Point Name
	MidpointName string `protobuf:"bytes,1,opt,name=midpoint_name,json=midpointName" json:"midpoint_name,omitempty"`
	// Tunnel Name
	TunnelName string `protobuf:"bytes,2,opt,name=tunnel_name,json=tunnelName" json:"tunnel_name,omitempty"`
	// Source Node ID
	SourceNodeId string `protobuf:"bytes,3,opt,name=source_node_id,json=sourceNodeId" json:"source_node_id,omitempty"`
	// Source Global ID
	SourceGlobalId uint32 `protobuf:"varint,4,opt,name=source_global_id,json=sourceGlobalId" json:"source_global_id,omitempty"`
	// Source Tunnel ID
	SourceTunnelId uint32 `protobuf:"varint,5,opt,name=source_tunnel_id,json=sourceTunnelId" json:"source_tunnel_id,omitempty"`
	// LSP ID
	LspId uint32 `protobuf:"varint,6,opt,name=lsp_id,json=lspId" json:"lsp_id,omitempty"`
	// Destination Node ID
	DestinationNodeId string `protobuf:"bytes,7,opt,name=destination_node_id,json=destinationNodeId" json:"destination_node_id,omitempty"`
	// Destination Global ID
	DestinationGlobalId uint32 `protobuf:"varint,8,opt,name=destination_global_id,json=destinationGlobalId" json:"destination_global_id,omitempty"`
	// Destination Tunnel ID
	DestinationTunnelId uint32 `protobuf:"varint,9,opt,name=destination_tunnel_id,json=destinationTunnelId" json:"destination_tunnel_id,omitempty"`
	// Forward LSP State
	ForwardLspState string `protobuf:"bytes,10,opt,name=forward_lsp_state,json=forwardLspState" json:"forward_lsp_state,omitempty"`
	// Reverse LSP State
	ReverseLspState      string   `protobuf:"bytes,11,opt,name=reverse_lsp_state,json=reverseLspState" json:"reverse_lsp_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTpMidInfo) Reset()         { *m = MplsTpMidInfo{} }
func (m *MplsTpMidInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTpMidInfo) ProtoMessage()    {}
func (*MplsTpMidInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_tp_detail_mid_info_22e0bb4fdf384efe, []int{3}
}
func (m *MplsTpMidInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTpMidInfo.Unmarshal(m, b)
}
func (m *MplsTpMidInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTpMidInfo.Marshal(b, m, deterministic)
}
func (dst *MplsTpMidInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTpMidInfo.Merge(dst, src)
}
func (m *MplsTpMidInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTpMidInfo.Size(m)
}
func (m *MplsTpMidInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTpMidInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTpMidInfo proto.InternalMessageInfo

func (m *MplsTpMidInfo) GetMidpointName() string {
	if m != nil {
		return m.MidpointName
	}
	return ""
}

func (m *MplsTpMidInfo) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *MplsTpMidInfo) GetSourceNodeId() string {
	if m != nil {
		return m.SourceNodeId
	}
	return ""
}

func (m *MplsTpMidInfo) GetSourceGlobalId() uint32 {
	if m != nil {
		return m.SourceGlobalId
	}
	return 0
}

func (m *MplsTpMidInfo) GetSourceTunnelId() uint32 {
	if m != nil {
		return m.SourceTunnelId
	}
	return 0
}

func (m *MplsTpMidInfo) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *MplsTpMidInfo) GetDestinationNodeId() string {
	if m != nil {
		return m.DestinationNodeId
	}
	return ""
}

func (m *MplsTpMidInfo) GetDestinationGlobalId() uint32 {
	if m != nil {
		return m.DestinationGlobalId
	}
	return 0
}

func (m *MplsTpMidInfo) GetDestinationTunnelId() uint32 {
	if m != nil {
		return m.DestinationTunnelId
	}
	return 0
}

func (m *MplsTpMidInfo) GetForwardLspState() string {
	if m != nil {
		return m.ForwardLspState
	}
	return ""
}

func (m *MplsTpMidInfo) GetReverseLspState() string {
	if m != nil {
		return m.ReverseLspState
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsTpDetailMidInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_midpoints.tp_detail_midpoints.tp_detail_midpoint.mpls_tp_detail_mid_info_KEYS")
	proto.RegisterType((*MplsTpDetailMidInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_midpoints.tp_detail_midpoints.tp_detail_midpoint.mpls_tp_detail_mid_info")
	proto.RegisterType((*MplsTpMidLspInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_midpoints.tp_detail_midpoints.tp_detail_midpoint.mpls_tp_mid_lsp_info")
	proto.RegisterType((*MplsTpMidInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_midpoints.tp_detail_midpoints.tp_detail_midpoint.mpls_tp_mid_info")
}

func init() {
	proto.RegisterFile("mpls_tp_detail_mid_info.proto", fileDescriptor_mpls_tp_detail_mid_info_22e0bb4fdf384efe)
}

var fileDescriptor_mpls_tp_detail_mid_info_22e0bb4fdf384efe = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x4d, 0x6f, 0x13, 0x3d,
	0x10, 0xc7, 0x95, 0xbe, 0xc7, 0x69, 0xfb, 0xb4, 0x6e, 0xab, 0x67, 0x51, 0x0b, 0x54, 0x81, 0x43,
	0x84, 0x44, 0x84, 0x5a, 0xe0, 0x0e, 0x08, 0xd1, 0x15, 0xa5, 0x87, 0x2d, 0x1c, 0x38, 0x59, 0x4e,
	0x3c, 0x49, 0xad, 0xee, 0xda, 0x96, 0xed, 0xb4, 0xfd, 0x10, 0x9c, 0xb8, 0xf2, 0xc9, 0x38, 0xf1,
	0x55, 0x90, 0xc7, 0x4e, 0x76, 0xfb, 0x76, 0x04, 0x8e, 0xf9, 0xcf, 0x7f, 0x66, 0x7e, 0x33, 0xf1,
	0x2c, 0x79, 0x58, 0x99, 0xd2, 0x31, 0x6f, 0x98, 0x00, 0xcf, 0x65, 0xc9, 0x2a, 0x29, 0x98, 0x54,
	0x23, 0xdd, 0x37, 0x56, 0x7b, 0x4d, 0xbf, 0x0c, 0xa5, 0x1b, 0x6a, 0x26, 0xb5, 0x63, 0x57, 0x96,
	0x45, 0x2f, 0x30, 0x6d, 0xc0, 0xf6, 0x53, 0x62, 0xdf, 0x9b, 0x90, 0x64, 0xb4, 0x54, 0xde, 0xf5,
	0xaf, 0x15, 0xba, 0x57, 0xeb, 0x1e, 0x91, 0xbd, 0x7b, 0xfa, 0xb2, 0x8f, 0xef, 0xbf, 0x9e, 0xd2,
	0x1e, 0xd9, 0x68, 0x94, 0x65, 0x8a, 0x57, 0x90, 0xb5, 0xf6, 0x5b, 0xbd, 0x76, 0xb1, 0xee, 0xcd,
	0xa7, 0x24, 0x9f, 0xf0, 0x0a, 0xba, 0xbf, 0x16, 0xc8, 0xff, 0xf7, 0x94, 0xa2, 0x3f, 0x5a, 0x64,
	0x67, 0x60, 0x25, 0x8c, 0x58, 0xe9, 0x0c, 0x4a, 0xb6, 0xe2, 0x5e, 0x6a, 0x95, 0x1d, 0xec, 0xb7,
	0x7a, 0x9d, 0x83, 0x71, 0xff, 0x8f, 0x4c, 0x37, 0xcd, 0x9d, 0x81, 0x14, 0x5b, 0x48, 0x71, 0xec,
	0x4c, 0x5e, 0x33, 0xd0, 0x6f, 0x2d, 0xd2, 0x19, 0x69, 0x7b, 0xc9, 0xad, 0x08, 0x7c, 0xd9, 0x21,
	0x32, 0x9d, 0xff, 0x05, 0xa6, 0xe9, 0x36, 0x0a, 0x92, 0xfa, 0x1f, 0x3b, 0x83, 0x38, 0x16, 0x2e,
	0xc0, 0x3a, 0x40, 0x9c, 0x97, 0xff, 0x00, 0x27, 0xf5, 0x0f, 0x38, 0xef, 0xc8, 0xa3, 0xc6, 0x72,
	0x98, 0x05, 0x07, 0xf6, 0x02, 0x04, 0x1b, 0x70, 0x25, 0x2e, 0xa5, 0xf0, 0x67, 0xd9, 0xab, 0xfd,
	0x56, 0x6f, 0xad, 0xd8, 0xad, 0x47, 0x28, 0x92, 0xe7, 0xed, 0xd4, 0x12, 0x8a, 0x34, 0x46, 0xba,
	0xab, 0xc8, 0xeb, 0x58, 0xa4, 0x6e, 0x7c, 0xab, 0x48, 0xf7, 0xfb, 0x1c, 0xd9, 0xbe, 0x0b, 0x97,
	0xee, 0x92, 0xb6, 0x9e, 0x78, 0x56, 0xf2, 0x01, 0x94, 0xf8, 0x3a, 0xd7, 0x8a, 0x15, 0x3d, 0xf1,
	0xc7, 0xe1, 0x37, 0x7d, 0x40, 0x56, 0x30, 0x28, 0xd5, 0x79, 0x36, 0x87, 0xb1, 0xe5, 0x10, 0x93,
	0xea, 0x9c, 0xbe, 0x20, 0xdb, 0x9a, 0x57, 0xcc, 0xc2, 0xc8, 0x82, 0x3b, 0x63, 0x52, 0x79, 0xb0,
	0x17, 0xbc, 0xcc, 0xe6, 0xd1, 0x46, 0x35, 0xaf, 0x8a, 0x18, 0xca, 0x53, 0x84, 0x3e, 0x27, 0x54,
	0x4f, 0xfc, 0x58, 0x4b, 0x35, 0x8e, 0xf6, 0x11, 0x1f, 0x42, 0xb6, 0x80, 0x07, 0xb1, 0x39, 0x8d,
	0xe4, 0xd3, 0x40, 0xb8, 0x1e, 0x05, 0x57, 0x9e, 0x9d, 0x69, 0xc3, 0xb8, 0x10, 0x16, 0x9c, 0xcb,
	0x16, 0xe3, 0xf5, 0x04, 0xfd, 0x48, 0x9b, 0x37, 0x51, 0x0d, 0x94, 0x52, 0xa5, 0x09, 0x96, 0x22,
	0xa5, 0x54, 0x71, 0x80, 0x3d, 0xd2, 0xae, 0xd7, 0xb4, 0x8c, 0xb1, 0x5a, 0xe8, 0xfe, 0x9c, 0x27,
	0x1b, 0x37, 0x9f, 0x39, 0x7d, 0x42, 0xd6, 0xee, 0x3a, 0xd9, 0xd5, 0xaa, 0x71, 0xb0, 0xf4, 0x31,
	0xe9, 0xf8, 0x89, 0x52, 0x50, 0x46, 0xcb, 0x1c, 0x5a, 0x48, 0x94, 0xd0, 0xf0, 0x94, 0xac, 0x3b,
	0x3d, 0xb1, 0x43, 0x60, 0x4a, 0x0b, 0x60, 0x52, 0xe0, 0x62, 0xda, 0xc5, 0x6a, 0x54, 0x4f, 0xb4,
	0x80, 0x5c, 0x84, 0x19, 0x93, 0x6b, 0x5c, 0xea, 0x01, 0x2f, 0x83, 0x6f, 0x01, 0x29, 0x53, 0xf6,
	0x07, 0x94, 0xaf, 0x39, 0x53, 0x5f, 0x29, 0x70, 0x1b, 0x33, 0xe7, 0x67, 0x94, 0x73, 0x41, 0x77,
	0xc8, 0x12, 0xfe, 0xb9, 0x22, 0xed, 0x62, 0xb1, 0x74, 0x26, 0x17, 0xb4, 0x4f, 0xb6, 0x04, 0x38,
	0x2f, 0x15, 0xde, 0xed, 0x8c, 0x6a, 0x39, 0xae, 0xbf, 0x11, 0x4a, 0x68, 0x07, 0x64, 0xa7, 0xe9,
	0xaf, 0xf9, 0x56, 0xb0, 0x6a, 0xb3, 0xd8, 0x0c, 0xf2, 0x46, 0x4e, 0x4d, 0xda, 0xbe, 0x95, 0x33,
	0xc3, 0x7d, 0x46, 0x36, 0x9b, 0x27, 0xe2, 0x3c, 0xf7, 0x90, 0x11, 0xa4, 0xfa, 0xaf, 0xbe, 0x8a,
	0xd3, 0x20, 0x07, 0x6f, 0xf3, 0x12, 0xa2, 0xb7, 0x13, 0xbd, 0xf5, 0xe3, 0x47, 0xef, 0x60, 0x09,
	0x3f, 0xfd, 0x87, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xad, 0x80, 0x51, 0x1b, 0x06, 0x00,
	0x00,
}