// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_tail_bfd_lsp_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_bfd_tail_infos_tail_info

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

// TE Tail End BFDOverLSP Info
type MplsTeTailBfdLspInfo_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	LspId                uint32   `protobuf:"varint,3,opt,name=lsp_id,json=lspId" json:"lsp_id,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,5,opt,name=extended_tunnel_id,json=extendedTunnelId" json:"extended_tunnel_id,omitempty"`
	CType                string   `protobuf:"bytes,6,opt,name=c_type,json=cType" json:"c_type,omitempty"`
	P2MpId               uint32   `protobuf:"varint,7,opt,name=p2_mp_id,json=p2MpId" json:"p2_mp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTailBfdLspInfo_KEYS) Reset()         { *m = MplsTeTailBfdLspInfo_KEYS{} }
func (m *MplsTeTailBfdLspInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeTailBfdLspInfo_KEYS) ProtoMessage()    {}
func (*MplsTeTailBfdLspInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_tail_bfd_lsp_info_25fd4978b5763f4b, []int{0}
}
func (m *MplsTeTailBfdLspInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTeTailBfdLspInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTeTailBfdLspInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS.Merge(dst, src)
}
func (m *MplsTeTailBfdLspInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS.Size(m)
}
func (m *MplsTeTailBfdLspInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS proto.InternalMessageInfo

func (m *MplsTeTailBfdLspInfo_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetCType() string {
	if m != nil {
		return m.CType
	}
	return ""
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

type MplsTeTailBfdLspInfo struct {
	// Signaled Name
	SignaledName string `protobuf:"bytes,50,opt,name=signaled_name,json=signaledName" json:"signaled_name,omitempty"`
	// FEC for the LSP
	LspFec *TeLspFecT `protobuf:"bytes,51,opt,name=lsp_fec,json=lspFec" json:"lsp_fec,omitempty"`
	// BFD Session State
	BfdSessionState      string   `protobuf:"bytes,52,opt,name=bfd_session_state,json=bfdSessionState" json:"bfd_session_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTailBfdLspInfo) Reset()         { *m = MplsTeTailBfdLspInfo{} }
func (m *MplsTeTailBfdLspInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeTailBfdLspInfo) ProtoMessage()    {}
func (*MplsTeTailBfdLspInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_tail_bfd_lsp_info_25fd4978b5763f4b, []int{1}
}
func (m *MplsTeTailBfdLspInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTailBfdLspInfo.Unmarshal(m, b)
}
func (m *MplsTeTailBfdLspInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTailBfdLspInfo.Marshal(b, m, deterministic)
}
func (dst *MplsTeTailBfdLspInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTailBfdLspInfo.Merge(dst, src)
}
func (m *MplsTeTailBfdLspInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeTailBfdLspInfo.Size(m)
}
func (m *MplsTeTailBfdLspInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTailBfdLspInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTailBfdLspInfo proto.InternalMessageInfo

func (m *MplsTeTailBfdLspInfo) GetSignaledName() string {
	if m != nil {
		return m.SignaledName
	}
	return ""
}

func (m *MplsTeTailBfdLspInfo) GetLspFec() *TeLspFecT {
	if m != nil {
		return m.LspFec
	}
	return nil
}

func (m *MplsTeTailBfdLspInfo) GetBfdSessionState() string {
	if m != nil {
		return m.BfdSessionState
	}
	return ""
}

// C-type-specific LSP FEC data
type TeLspFecCtypeDataT struct {
	FecCType string `protobuf:"bytes,1,opt,name=fec_c_type,json=fecCType" json:"fec_c_type,omitempty"`
	// P2P LSP destination
	P2PLspDestination string `protobuf:"bytes,2,opt,name=p2_p_lsp_destination,json=p2PLspDestination" json:"p2_p_lsp_destination,omitempty"`
	// P2MP ID
	FecDestinationP2MpId uint32   `protobuf:"varint,3,opt,name=fec_destination_p2_mp_id,json=fecDestinationP2MpId" json:"fec_destination_p2_mp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeLspFecCtypeDataT) Reset()         { *m = TeLspFecCtypeDataT{} }
func (m *TeLspFecCtypeDataT) String() string { return proto.CompactTextString(m) }
func (*TeLspFecCtypeDataT) ProtoMessage()    {}
func (*TeLspFecCtypeDataT) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_tail_bfd_lsp_info_25fd4978b5763f4b, []int{2}
}
func (m *TeLspFecCtypeDataT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeLspFecCtypeDataT.Unmarshal(m, b)
}
func (m *TeLspFecCtypeDataT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeLspFecCtypeDataT.Marshal(b, m, deterministic)
}
func (dst *TeLspFecCtypeDataT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeLspFecCtypeDataT.Merge(dst, src)
}
func (m *TeLspFecCtypeDataT) XXX_Size() int {
	return xxx_messageInfo_TeLspFecCtypeDataT.Size(m)
}
func (m *TeLspFecCtypeDataT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeLspFecCtypeDataT.DiscardUnknown(m)
}

var xxx_messageInfo_TeLspFecCtypeDataT proto.InternalMessageInfo

func (m *TeLspFecCtypeDataT) GetFecCType() string {
	if m != nil {
		return m.FecCType
	}
	return ""
}

func (m *TeLspFecCtypeDataT) GetP2PLspDestination() string {
	if m != nil {
		return m.P2PLspDestination
	}
	return ""
}

func (m *TeLspFecCtypeDataT) GetFecDestinationP2MpId() uint32 {
	if m != nil {
		return m.FecDestinationP2MpId
	}
	return 0
}

// A LSP FEC
type TeLspFecT struct {
	// LSP ID
	FecLspId uint32 `protobuf:"varint,1,opt,name=fec_lsp_id,json=fecLspId" json:"fec_lsp_id,omitempty"`
	// Tunnel ID
	FecTunnelId uint32 `protobuf:"varint,2,opt,name=fec_tunnel_id,json=fecTunnelId" json:"fec_tunnel_id,omitempty"`
	// Extended tunnel ID
	FecExtendedTunnelId string `protobuf:"bytes,3,opt,name=fec_extended_tunnel_id,json=fecExtendedTunnelId" json:"fec_extended_tunnel_id,omitempty"`
	// Tunnel source address
	FecSource string `protobuf:"bytes,4,opt,name=fec_source,json=fecSource" json:"fec_source,omitempty"`
	// Destination or P2MP ID
	FecDestinationInfo *TeLspFecCtypeDataT `protobuf:"bytes,5,opt,name=fec_destination_info,json=fecDestinationInfo" json:"fec_destination_info,omitempty"`
	// VRF; currently only for GMPLS tunnels
	FecVrf               string   `protobuf:"bytes,6,opt,name=fec_vrf,json=fecVrf" json:"fec_vrf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeLspFecT) Reset()         { *m = TeLspFecT{} }
func (m *TeLspFecT) String() string { return proto.CompactTextString(m) }
func (*TeLspFecT) ProtoMessage()    {}
func (*TeLspFecT) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_tail_bfd_lsp_info_25fd4978b5763f4b, []int{3}
}
func (m *TeLspFecT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeLspFecT.Unmarshal(m, b)
}
func (m *TeLspFecT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeLspFecT.Marshal(b, m, deterministic)
}
func (dst *TeLspFecT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeLspFecT.Merge(dst, src)
}
func (m *TeLspFecT) XXX_Size() int {
	return xxx_messageInfo_TeLspFecT.Size(m)
}
func (m *TeLspFecT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeLspFecT.DiscardUnknown(m)
}

var xxx_messageInfo_TeLspFecT proto.InternalMessageInfo

func (m *TeLspFecT) GetFecLspId() uint32 {
	if m != nil {
		return m.FecLspId
	}
	return 0
}

func (m *TeLspFecT) GetFecTunnelId() uint32 {
	if m != nil {
		return m.FecTunnelId
	}
	return 0
}

func (m *TeLspFecT) GetFecExtendedTunnelId() string {
	if m != nil {
		return m.FecExtendedTunnelId
	}
	return ""
}

func (m *TeLspFecT) GetFecSource() string {
	if m != nil {
		return m.FecSource
	}
	return ""
}

func (m *TeLspFecT) GetFecDestinationInfo() *TeLspFecCtypeDataT {
	if m != nil {
		return m.FecDestinationInfo
	}
	return nil
}

func (m *TeLspFecT) GetFecVrf() string {
	if m != nil {
		return m.FecVrf
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsTeTailBfdLspInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.tail_infos.tail_info.mpls_te_tail_bfd_lsp_info_KEYS")
	proto.RegisterType((*MplsTeTailBfdLspInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.tail_infos.tail_info.mpls_te_tail_bfd_lsp_info")
	proto.RegisterType((*TeLspFecCtypeDataT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.tail_infos.tail_info.te_lsp_fec_ctype_data_t")
	proto.RegisterType((*TeLspFecT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.tail_infos.tail_info.te_lsp_fec_t")
}

func init() {
	proto.RegisterFile("mpls_te_tail_bfd_lsp_info.proto", fileDescriptor_mpls_te_tail_bfd_lsp_info_25fd4978b5763f4b)
}

var fileDescriptor_mpls_te_tail_bfd_lsp_info_25fd4978b5763f4b = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x8a, 0x13, 0x4d,
	0x10, 0x65, 0xb2, 0x5f, 0xfe, 0x2a, 0x9b, 0x4f, 0xb7, 0x37, 0xba, 0x23, 0xfe, 0x85, 0x88, 0x10,
	0x44, 0x66, 0x61, 0x22, 0xde, 0xfb, 0xb3, 0x42, 0x70, 0x95, 0x90, 0x2c, 0x82, 0x78, 0xd1, 0x4c,
	0xba, 0xab, 0x65, 0x60, 0x32, 0xd3, 0x4c, 0xf7, 0xca, 0xe6, 0x19, 0xc4, 0x77, 0xf0, 0x85, 0xbc,
	0xf6, 0x75, 0xa4, 0xab, 0x33, 0xc9, 0x18, 0xdc, 0x2b, 0xbd, 0x4b, 0x4e, 0x55, 0x9d, 0xaa, 0x3e,
	0xe7, 0x30, 0xf0, 0x70, 0xa5, 0x33, 0xc3, 0x2d, 0x72, 0x9b, 0xa4, 0x19, 0x5f, 0x2a, 0xc9, 0x33,
	0xa3, 0x79, 0x9a, 0xab, 0x22, 0xd2, 0x65, 0x61, 0x0b, 0xf6, 0x52, 0xa4, 0x46, 0x14, 0x3c, 0x2d,
	0x0c, 0xbf, 0x2a, 0x79, 0xd5, 0x5d, 0x68, 0x2c, 0xa3, 0xea, 0x8f, 0xb1, 0x49, 0x2e, 0x97, 0xeb,
	0x68, 0xa9, 0x64, 0x44, 0x34, 0x6e, 0xdc, 0xec, 0x7e, 0x8e, 0xbe, 0x36, 0xe0, 0xc1, 0xb5, 0x7b,
	0xf8, 0xdb, 0xb3, 0x8f, 0x0b, 0xf6, 0x18, 0xfe, 0x37, 0xc5, 0x65, 0x29, 0x90, 0x27, 0x52, 0x96,
	0x68, 0x4c, 0x18, 0x0c, 0x83, 0x71, 0x77, 0xde, 0xf7, 0xe8, 0x0b, 0x0f, 0xb2, 0xbb, 0xd0, 0xb5,
	0x97, 0x79, 0x8e, 0x19, 0x4f, 0x65, 0xd8, 0x18, 0x06, 0xe3, 0xfe, 0xbc, 0xe3, 0x81, 0xa9, 0x64,
	0xb7, 0xa0, 0x45, 0xa4, 0x32, 0x3c, 0xa0, 0x4a, 0x33, 0x33, 0x7a, 0x2a, 0xd9, 0x29, 0x1c, 0x4b,
	0x34, 0x36, 0xcd, 0x13, 0x9b, 0x16, 0xf9, 0x96, 0xff, 0x3f, 0xe2, 0x67, 0xb5, 0x52, 0xb5, 0xe4,
	0x29, 0x30, 0xbc, 0xb2, 0x98, 0x4b, 0x94, 0x7c, 0xb7, 0xad, 0x49, 0xfd, 0x37, 0xab, 0xca, 0x45,
	0x6d, 0xab, 0xe0, 0x76, 0xad, 0x31, 0x6c, 0x51, 0x47, 0x53, 0x5c, 0xac, 0x35, 0xb2, 0x10, 0x3a,
	0x3a, 0xe6, 0x2b, 0x3a, 0xa7, 0x4d, 0xe7, 0xb4, 0x74, 0xfc, 0x4e, 0x4f, 0xe5, 0xe8, 0x67, 0x00,
	0x77, 0xae, 0x55, 0x83, 0x3d, 0x82, 0xbe, 0x49, 0x3f, 0xe7, 0x49, 0x86, 0x92, 0xe7, 0xc9, 0x0a,
	0xc3, 0x98, 0x58, 0x0f, 0x2b, 0xf0, 0x7d, 0xb2, 0x42, 0x96, 0x42, 0xdb, 0x0d, 0x28, 0x14, 0xe1,
	0x64, 0x18, 0x8c, 0x7b, 0xf1, 0x2c, 0xfa, 0x7b, 0x9b, 0x22, 0x8b, 0x7c, 0xc3, 0xca, 0xed, 0xdc,
	0x49, 0xf9, 0x06, 0x05, 0x7b, 0x02, 0x47, 0xee, 0x3e, 0x83, 0xc6, 0x38, 0xf5, 0x8c, 0x4d, 0x2c,
	0x86, 0xcf, 0xe8, 0xa6, 0x1b, 0x4b, 0x25, 0x17, 0x1e, 0x5f, 0x38, 0x78, 0xf4, 0x3d, 0x80, 0x93,
	0x1a, 0x89, 0x70, 0xaa, 0x70, 0x99, 0xd8, 0x84, 0x5b, 0x76, 0x0f, 0x80, 0x30, 0x2f, 0x95, 0x37,
	0xb7, 0xa3, 0x50, 0xbc, 0x22, 0xb5, 0x4e, 0x61, 0xa0, 0x63, 0xae, 0x69, 0xb4, 0xe6, 0x08, 0x59,
	0xdc, 0x9d, 0x1f, 0xe9, 0x78, 0x76, 0x6e, 0xf4, 0xeb, 0x5d, 0x81, 0x3d, 0x87, 0xd0, 0xd1, 0xd5,
	0x8d, 0xdd, 0xca, 0xed, 0xdd, 0x1f, 0x28, 0x14, 0xb5, 0x89, 0x99, 0x17, 0xff, 0x47, 0x03, 0x0e,
	0xeb, 0xef, 0xac, 0xee, 0xda, 0x04, 0x27, 0xf0, 0x91, 0x52, 0x28, 0xce, 0x29, 0x3b, 0x23, 0xe8,
	0x53, 0xdb, 0x5e, 0xe6, 0x7a, 0x0a, 0xc5, 0x36, 0x00, 0x13, 0xb8, 0xed, 0x7a, 0xfe, 0x10, 0x99,
	0x03, 0xba, 0xfe, 0x58, 0xa1, 0x38, 0xdb, 0x4f, 0xcd, 0x7d, 0xbf, 0xd6, 0xa7, 0x7b, 0x93, 0xc5,
	0xae, 0x42, 0xb1, 0x20, 0x80, 0x7d, 0x0b, 0x60, 0xb0, 0xff, 0x3e, 0xe7, 0x11, 0xa5, 0xb0, 0x17,
	0x7f, 0xfa, 0xc7, 0x76, 0xd7, 0x9d, 0x9a, 0xb3, 0xdf, 0x85, 0x9b, 0xba, 0x54, 0x9e, 0x40, 0xdb,
	0xf5, 0x7d, 0x29, 0xd5, 0x26, 0xe5, 0x2d, 0x85, 0xe2, 0x43, 0xa9, 0x96, 0x2d, 0xfa, 0x4a, 0x4c,
	0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xff, 0x28, 0xd9, 0x38, 0x48, 0x04, 0x00, 0x00,
}
