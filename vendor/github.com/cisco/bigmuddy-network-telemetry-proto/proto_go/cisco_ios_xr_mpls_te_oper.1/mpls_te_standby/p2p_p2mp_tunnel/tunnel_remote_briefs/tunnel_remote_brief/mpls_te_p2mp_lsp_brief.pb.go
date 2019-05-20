// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_p2mp_lsp_brief.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_p2p_p2mp_tunnel_tunnel_remote_briefs_tunnel_remote_brief

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

// A P2MP LSP
type MplsTeP2MpLspBrief_KEYS struct {
	LspId                uint32   `protobuf:"varint,1,opt,name=lsp_id,json=lspId" json:"lsp_id,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,3,opt,name=extended_tunnel_id,json=extendedTunnelId" json:"extended_tunnel_id,omitempty"`
	SourceAddress        string   `protobuf:"bytes,4,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	CType                string   `protobuf:"bytes,5,opt,name=c_type,json=cType" json:"c_type,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,6,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	P2MpId               uint32   `protobuf:"varint,7,opt,name=p2_mp_id,json=p2MpId" json:"p2_mp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeP2MpLspBrief_KEYS) Reset()         { *m = MplsTeP2MpLspBrief_KEYS{} }
func (m *MplsTeP2MpLspBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeP2MpLspBrief_KEYS) ProtoMessage()    {}
func (*MplsTeP2MpLspBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_p2mp_lsp_brief_b5d9d2c46b014819, []int{0}
}
func (m *MplsTeP2MpLspBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeP2MpLspBrief_KEYS.Unmarshal(m, b)
}
func (m *MplsTeP2MpLspBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeP2MpLspBrief_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTeP2MpLspBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeP2MpLspBrief_KEYS.Merge(dst, src)
}
func (m *MplsTeP2MpLspBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeP2MpLspBrief_KEYS.Size(m)
}
func (m *MplsTeP2MpLspBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeP2MpLspBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeP2MpLspBrief_KEYS proto.InternalMessageInfo

func (m *MplsTeP2MpLspBrief_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *MplsTeP2MpLspBrief_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *MplsTeP2MpLspBrief_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *MplsTeP2MpLspBrief_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsTeP2MpLspBrief_KEYS) GetCType() string {
	if m != nil {
		return m.CType
	}
	return ""
}

func (m *MplsTeP2MpLspBrief_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTeP2MpLspBrief_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

type MplsTeP2MpLspBrief struct {
	// Signaled Name
	SignaledName string `protobuf:"bytes,50,opt,name=signaled_name,json=signaledName" json:"signaled_name,omitempty"`
	// FEC for the LSP
	LspFec *TeLspFecT `protobuf:"bytes,51,opt,name=lsp_fec,json=lspFec" json:"lsp_fec,omitempty"`
	// array of S2L structures
	S2LList []*MplsTeP2MpS2LBrief `protobuf:"bytes,52,rep,name=s2_l_list,json=s2LList" json:"s2_l_list,omitempty"`
	// Time in seconds since the LSP was up
	Uptime               uint32   `protobuf:"varint,53,opt,name=uptime" json:"uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeP2MpLspBrief) Reset()         { *m = MplsTeP2MpLspBrief{} }
func (m *MplsTeP2MpLspBrief) String() string { return proto.CompactTextString(m) }
func (*MplsTeP2MpLspBrief) ProtoMessage()    {}
func (*MplsTeP2MpLspBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_p2mp_lsp_brief_b5d9d2c46b014819, []int{1}
}
func (m *MplsTeP2MpLspBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeP2MpLspBrief.Unmarshal(m, b)
}
func (m *MplsTeP2MpLspBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeP2MpLspBrief.Marshal(b, m, deterministic)
}
func (dst *MplsTeP2MpLspBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeP2MpLspBrief.Merge(dst, src)
}
func (m *MplsTeP2MpLspBrief) XXX_Size() int {
	return xxx_messageInfo_MplsTeP2MpLspBrief.Size(m)
}
func (m *MplsTeP2MpLspBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeP2MpLspBrief.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeP2MpLspBrief proto.InternalMessageInfo

func (m *MplsTeP2MpLspBrief) GetSignaledName() string {
	if m != nil {
		return m.SignaledName
	}
	return ""
}

func (m *MplsTeP2MpLspBrief) GetLspFec() *TeLspFecT {
	if m != nil {
		return m.LspFec
	}
	return nil
}

func (m *MplsTeP2MpLspBrief) GetS2LList() []*MplsTeP2MpS2LBrief {
	if m != nil {
		return m.S2LList
	}
	return nil
}

func (m *MplsTeP2MpLspBrief) GetUptime() uint32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

// A sub-LSP FEC
type TeS2LFecT struct {
	// sub-LSP subgroup ID
	S2LFecSubgroupId uint32 `protobuf:"varint,1,opt,name=s2_l_fec_subgroup_id,json=s2LFecSubgroupId" json:"s2_l_fec_subgroup_id,omitempty"`
	// LSP ID
	S2LFecLspId uint32 `protobuf:"varint,2,opt,name=s2_l_fec_lsp_id,json=s2LFecLspId" json:"s2_l_fec_lsp_id,omitempty"`
	// Tunnel ID
	S2LFecTunnelId uint32 `protobuf:"varint,3,opt,name=s2_l_fec_tunnel_id,json=s2LFecTunnelId" json:"s2_l_fec_tunnel_id,omitempty"`
	// Extended tunnel ID
	S2LFecExtendedTunnelId string `protobuf:"bytes,4,opt,name=s2_l_fec_extended_tunnel_id,json=s2LFecExtendedTunnelId" json:"s2_l_fec_extended_tunnel_id,omitempty"`
	// LSP source address
	S2LFecSource string `protobuf:"bytes,5,opt,name=s2_l_fec_source,json=s2LFecSource" json:"s2_l_fec_source,omitempty"`
	// sub-LSP destination address
	S2LFecDest string `protobuf:"bytes,6,opt,name=s2_l_fec_dest,json=s2LFecDest" json:"s2_l_fec_dest,omitempty"`
	// P2MP ID
	S2LFecP2MpId uint32 `protobuf:"varint,7,opt,name=s2_l_fec_p2_mp_id,json=s2LFecP2MpId" json:"s2_l_fec_p2_mp_id,omitempty"`
	// Subgroup Originator
	S2LFecSubgroupOriginator string `protobuf:"bytes,8,opt,name=s2_l_fec_subgroup_originator,json=s2LFecSubgroupOriginator" json:"s2_l_fec_subgroup_originator,omitempty"`
	// Session identifier (ctype)
	S2LFecCType string `protobuf:"bytes,9,opt,name=s2_l_fec_c_type,json=s2LFecCType" json:"s2_l_fec_c_type,omitempty"`
	// VRF; currently only for GMPLS tunnels
	S2LFecVrf            string   `protobuf:"bytes,10,opt,name=s2_l_fec_vrf,json=s2LFecVrf" json:"s2_l_fec_vrf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeS2LFecT) Reset()         { *m = TeS2LFecT{} }
func (m *TeS2LFecT) String() string { return proto.CompactTextString(m) }
func (*TeS2LFecT) ProtoMessage()    {}
func (*TeS2LFecT) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_p2mp_lsp_brief_b5d9d2c46b014819, []int{2}
}
func (m *TeS2LFecT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeS2LFecT.Unmarshal(m, b)
}
func (m *TeS2LFecT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeS2LFecT.Marshal(b, m, deterministic)
}
func (dst *TeS2LFecT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeS2LFecT.Merge(dst, src)
}
func (m *TeS2LFecT) XXX_Size() int {
	return xxx_messageInfo_TeS2LFecT.Size(m)
}
func (m *TeS2LFecT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeS2LFecT.DiscardUnknown(m)
}

var xxx_messageInfo_TeS2LFecT proto.InternalMessageInfo

func (m *TeS2LFecT) GetS2LFecSubgroupId() uint32 {
	if m != nil {
		return m.S2LFecSubgroupId
	}
	return 0
}

func (m *TeS2LFecT) GetS2LFecLspId() uint32 {
	if m != nil {
		return m.S2LFecLspId
	}
	return 0
}

func (m *TeS2LFecT) GetS2LFecTunnelId() uint32 {
	if m != nil {
		return m.S2LFecTunnelId
	}
	return 0
}

func (m *TeS2LFecT) GetS2LFecExtendedTunnelId() string {
	if m != nil {
		return m.S2LFecExtendedTunnelId
	}
	return ""
}

func (m *TeS2LFecT) GetS2LFecSource() string {
	if m != nil {
		return m.S2LFecSource
	}
	return ""
}

func (m *TeS2LFecT) GetS2LFecDest() string {
	if m != nil {
		return m.S2LFecDest
	}
	return ""
}

func (m *TeS2LFecT) GetS2LFecP2MpId() uint32 {
	if m != nil {
		return m.S2LFecP2MpId
	}
	return 0
}

func (m *TeS2LFecT) GetS2LFecSubgroupOriginator() string {
	if m != nil {
		return m.S2LFecSubgroupOriginator
	}
	return ""
}

func (m *TeS2LFecT) GetS2LFecCType() string {
	if m != nil {
		return m.S2LFecCType
	}
	return ""
}

func (m *TeS2LFecT) GetS2LFecVrf() string {
	if m != nil {
		return m.S2LFecVrf
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
	return fileDescriptor_mpls_te_p2mp_lsp_brief_b5d9d2c46b014819, []int{3}
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
	return fileDescriptor_mpls_te_p2mp_lsp_brief_b5d9d2c46b014819, []int{4}
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

// A P2MP S2L
type MplsTeP2MpS2LBrief struct {
	// FEC for the S2L
	S2LFec *TeS2LFecT `protobuf:"bytes,1,opt,name=s2_l_fec,json=s2LFec" json:"s2_l_fec,omitempty"`
	// Time in seconds since the S2L was up
	Uptime uint32 `protobuf:"varint,2,opt,name=uptime" json:"uptime,omitempty"`
	// Egress Interface
	EgressInterface string `protobuf:"bytes,3,opt,name=egress_interface,json=egressInterface" json:"egress_interface,omitempty"`
	// Ingress Interface
	IngressInterface string `protobuf:"bytes,4,opt,name=ingress_interface,json=ingressInterface" json:"ingress_interface,omitempty"`
	// Role of S2L
	Role string `protobuf:"bytes,5,opt,name=role" json:"role,omitempty"`
	// Oper state of S2L
	OperState            bool     `protobuf:"varint,6,opt,name=oper_state,json=operState" json:"oper_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeP2MpS2LBrief) Reset()         { *m = MplsTeP2MpS2LBrief{} }
func (m *MplsTeP2MpS2LBrief) String() string { return proto.CompactTextString(m) }
func (*MplsTeP2MpS2LBrief) ProtoMessage()    {}
func (*MplsTeP2MpS2LBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_p2mp_lsp_brief_b5d9d2c46b014819, []int{5}
}
func (m *MplsTeP2MpS2LBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeP2MpS2LBrief.Unmarshal(m, b)
}
func (m *MplsTeP2MpS2LBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeP2MpS2LBrief.Marshal(b, m, deterministic)
}
func (dst *MplsTeP2MpS2LBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeP2MpS2LBrief.Merge(dst, src)
}
func (m *MplsTeP2MpS2LBrief) XXX_Size() int {
	return xxx_messageInfo_MplsTeP2MpS2LBrief.Size(m)
}
func (m *MplsTeP2MpS2LBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeP2MpS2LBrief.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeP2MpS2LBrief proto.InternalMessageInfo

func (m *MplsTeP2MpS2LBrief) GetS2LFec() *TeS2LFecT {
	if m != nil {
		return m.S2LFec
	}
	return nil
}

func (m *MplsTeP2MpS2LBrief) GetUptime() uint32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *MplsTeP2MpS2LBrief) GetEgressInterface() string {
	if m != nil {
		return m.EgressInterface
	}
	return ""
}

func (m *MplsTeP2MpS2LBrief) GetIngressInterface() string {
	if m != nil {
		return m.IngressInterface
	}
	return ""
}

func (m *MplsTeP2MpS2LBrief) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *MplsTeP2MpS2LBrief) GetOperState() bool {
	if m != nil {
		return m.OperState
	}
	return false
}

func init() {
	proto.RegisterType((*MplsTeP2MpLspBrief_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2p_p2mp_tunnel.tunnel_remote_briefs.tunnel_remote_brief.mpls_te_p2mp_lsp_brief_KEYS")
	proto.RegisterType((*MplsTeP2MpLspBrief)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2p_p2mp_tunnel.tunnel_remote_briefs.tunnel_remote_brief.mpls_te_p2mp_lsp_brief")
	proto.RegisterType((*TeS2LFecT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2p_p2mp_tunnel.tunnel_remote_briefs.tunnel_remote_brief.te_s2l_fec_t")
	proto.RegisterType((*TeLspFecCtypeDataT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2p_p2mp_tunnel.tunnel_remote_briefs.tunnel_remote_brief.te_lsp_fec_ctype_data_t")
	proto.RegisterType((*TeLspFecT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2p_p2mp_tunnel.tunnel_remote_briefs.tunnel_remote_brief.te_lsp_fec_t")
	proto.RegisterType((*MplsTeP2MpS2LBrief)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2p_p2mp_tunnel.tunnel_remote_briefs.tunnel_remote_brief.mpls_te_p2mp_s2l_brief")
}

func init() {
	proto.RegisterFile("mpls_te_p2mp_lsp_brief.proto", fileDescriptor_mpls_te_p2mp_lsp_brief_b5d9d2c46b014819)
}

var fileDescriptor_mpls_te_p2mp_lsp_brief_b5d9d2c46b014819 = []byte{
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xd1, 0x6e, 0xd3, 0x48,
	0x14, 0x86, 0xe5, 0xb4, 0x4d, 0xe2, 0x93, 0xa4, 0x4d, 0xa7, 0xd9, 0xd4, 0x52, 0xbb, 0xda, 0x6c,
	0x76, 0x57, 0x9b, 0x5d, 0x50, 0x2a, 0xb9, 0xc0, 0x0d, 0x12, 0x12, 0xa2, 0xad, 0x14, 0x11, 0xa0,
	0x72, 0x2b, 0x24, 0xae, 0x46, 0x8e, 0x3d, 0x8e, 0x2c, 0x39, 0xf6, 0xc8, 0x33, 0x41, 0x2d, 0xbc,
	0x00, 0xf0, 0x00, 0x5c, 0xf3, 0x02, 0x3c, 0x07, 0xbc, 0x0c, 0x2f, 0xc0, 0x15, 0x9a, 0x39, 0xb6,
	0xe3, 0xa4, 0xe1, 0x92, 0x5e, 0x25, 0x3e, 0xe7, 0x9f, 0xe3, 0x33, 0xff, 0x7c, 0x73, 0x12, 0x38,
	0x9c, 0xf1, 0x48, 0x50, 0xc9, 0x28, 0xb7, 0x67, 0x9c, 0x46, 0x82, 0xd3, 0x49, 0x1a, 0xb2, 0x60,
	0xc8, 0xd3, 0x44, 0x26, 0x64, 0xe2, 0x85, 0xc2, 0x4b, 0x68, 0x98, 0x08, 0x7a, 0x95, 0xd2, 0x5c,
	0x9a, 0x70, 0x96, 0x0e, 0xf3, 0x07, 0x21, 0xdd, 0xd8, 0x9f, 0x5c, 0x0f, 0xb9, 0xcd, 0xb1, 0x86,
	0x9c, 0xc7, 0x31, 0x8b, 0x86, 0xf8, 0x41, 0x53, 0x36, 0x4b, 0x24, 0xc3, 0xa2, 0x62, 0x5d, 0xb0,
	0xff, 0xae, 0x02, 0x07, 0xeb, 0x9b, 0xa0, 0x4f, 0x4f, 0x5f, 0x5d, 0x90, 0xdf, 0xa0, 0xaa, 0x22,
	0xa1, 0x6f, 0x19, 0x3d, 0x63, 0xd0, 0x72, 0xb6, 0x22, 0xc1, 0x47, 0x3e, 0x39, 0x00, 0x33, 0xab,
	0x16, 0xfa, 0x56, 0x45, 0x67, 0xea, 0x18, 0x18, 0xf9, 0xe4, 0x2e, 0x10, 0x76, 0x25, 0x59, 0xec,
	0x33, 0x9f, 0x2e, 0x54, 0x1b, 0x3d, 0x63, 0x60, 0x3a, 0xed, 0x3c, 0x73, 0x99, 0xab, 0xff, 0x81,
	0x6d, 0x91, 0xcc, 0x53, 0x8f, 0x51, 0xd7, 0xf7, 0x53, 0x26, 0x84, 0xb5, 0xa9, 0x95, 0x2d, 0x8c,
	0x3e, 0xc6, 0xa0, 0x6a, 0xc4, 0xa3, 0xf2, 0x9a, 0x33, 0x6b, 0x4b, 0xa7, 0xb7, 0xbc, 0xcb, 0x6b,
	0xce, 0xc8, 0x11, 0xec, 0xf9, 0x4c, 0xc8, 0x30, 0x76, 0x65, 0x98, 0xc4, 0x45, 0x89, 0xaa, 0xd6,
	0x90, 0x52, 0x2a, 0xaf, 0x63, 0x41, 0x9d, 0xdb, 0x74, 0xa6, 0xb7, 0x54, 0xd3, 0x8d, 0x57, 0xb9,
	0xfd, 0x8c, 0x8f, 0xfc, 0xfe, 0xf7, 0x0a, 0x74, 0xd7, 0x5b, 0x41, 0xfe, 0x82, 0x96, 0x08, 0xa7,
	0xb1, 0x1b, 0x31, 0x9f, 0xc6, 0xee, 0x8c, 0x59, 0xb6, 0xae, 0xdf, 0xcc, 0x83, 0xcf, 0xdd, 0x19,
	0x23, 0xef, 0x0d, 0xa8, 0xa9, 0x25, 0x01, 0xf3, 0xac, 0xe3, 0x9e, 0x31, 0x68, 0xd8, 0x7c, 0xf8,
	0xeb, 0x4f, 0x70, 0x28, 0x19, 0xcd, 0xde, 0x4a, 0xa5, 0xa3, 0x0e, 0xeb, 0x8c, 0x79, 0xe4, 0xa3,
	0x01, 0xa6, 0xb0, 0x69, 0x44, 0xa3, 0x50, 0x48, 0xeb, 0x5e, 0x6f, 0x63, 0xd0, 0xb0, 0xdf, 0xdc,
	0x46, 0x37, 0x4b, 0x06, 0x0a, 0x3b, 0xc2, 0xb0, 0x53, 0x13, 0xf6, 0x78, 0x1c, 0x0a, 0x49, 0xba,
	0x50, 0x9d, 0x73, 0x19, 0xce, 0x98, 0x75, 0x1f, 0xcd, 0xc7, 0xa7, 0xfe, 0x97, 0x0d, 0x68, 0xaa,
	0xf7, 0xdb, 0x11, 0xee, 0x84, 0x0c, 0xa1, 0xa3, 0x37, 0xa0, 0x9e, 0xc4, 0x7c, 0x32, 0x4d, 0x93,
	0x79, 0x09, 0xc3, 0xb6, 0xb0, 0xc7, 0x67, 0xcc, 0xbb, 0xc8, 0x12, 0x23, 0x9f, 0xfc, 0x0d, 0x3b,
	0x85, 0x3e, 0x23, 0x16, 0xb9, 0x6c, 0xa0, 0x74, 0xac, 0xb9, 0xfd, 0x1f, 0x48, 0xa1, 0x5a, 0x46,
	0xb3, 0xe5, 0x6c, 0xa3, 0xb0, 0x00, 0xf3, 0x21, 0x1c, 0x14, 0xda, 0x35, 0x3c, 0x23, 0xa5, 0x5d,
	0x5c, 0x74, 0x7a, 0x93, 0xea, 0x45, 0x3b, 0x08, 0x72, 0xc6, 0x6d, 0x33, 0xeb, 0x5c, 0xc7, 0xc8,
	0x9f, 0xd0, 0x2a, 0x64, 0x0a, 0xd6, 0x0c, 0x5c, 0x40, 0xd1, 0x09, 0x13, 0x92, 0xfc, 0x0b, 0xbb,
	0x85, 0x64, 0x85, 0xdc, 0xac, 0xd6, 0xb9, 0xe6, 0x97, 0x3c, 0x82, 0xc3, 0x9b, 0x8e, 0x25, 0x69,
	0x38, 0x55, 0x37, 0x20, 0x49, 0xad, 0xba, 0x2e, 0x6d, 0x2d, 0x3b, 0xf7, 0xa2, 0xc8, 0x2f, 0x39,
	0x98, 0x5d, 0x35, 0x53, 0x2f, 0xc9, 0x1c, 0x7c, 0xa2, 0x2f, 0xdc, 0x1f, 0xd0, 0x2c, 0x54, 0xaf,
	0xd3, 0xc0, 0x02, 0x2d, 0x31, 0x51, 0xf2, 0x32, 0x0d, 0xfa, 0x9f, 0x0c, 0xd8, 0x2f, 0x31, 0xe9,
	0xa9, 0x42, 0xd4, 0x77, 0xa5, 0x4b, 0x25, 0x39, 0x04, 0x28, 0x55, 0x37, 0xf4, 0xd2, 0x7a, 0x90,
	0x97, 0x3e, 0x82, 0x0e, 0xb7, 0x29, 0xde, 0xbb, 0xd2, 0xcd, 0xd5, 0xe7, 0x68, 0x3a, 0xbb, 0xdc,
	0x3e, 0x1f, 0x0b, 0x7e, 0xb2, 0x48, 0x90, 0x07, 0x60, 0xe5, 0xc6, 0xe5, 0x03, 0xa0, 0x70, 0x08,
	0xcf, 0xb4, 0x13, 0xa0, 0x8b, 0x59, 0x1a, 0x9d, 0xea, 0x7f, 0xab, 0x68, 0xd8, 0x8a, 0x6b, 0x93,
	0xf7, 0xb5, 0x34, 0xe9, 0x54, 0x5f, 0x08, 0x4d, 0x1f, 0x5a, 0xcb, 0xbc, 0x64, 0x60, 0x05, 0x25,
	0x58, 0x8e, 0xa1, 0xfb, 0x13, 0x4e, 0x70, 0xee, 0xed, 0x05, 0x6b, 0x20, 0xf9, 0x1d, 0x5f, 0x9b,
	0xf1, 0x81, 0x40, 0x99, 0x41, 0x01, 0xc7, 0x67, 0x03, 0x3a, 0xab, 0xfb, 0x0b, 0xe3, 0x20, 0xd1,
	0x24, 0x35, 0xec, 0xb7, 0xb7, 0x3c, 0x5d, 0xca, 0x27, 0xe9, 0x90, 0x65, 0x63, 0x47, 0x71, 0x90,
	0x90, 0x7d, 0xa8, 0xe5, 0x54, 0x20, 0xc6, 0xd5, 0x00, 0x91, 0xf8, 0xba, 0x3a, 0x59, 0x8b, 0xc1,
	0x40, 0x3e, 0x18, 0x50, 0xcf, 0x79, 0xd2, 0xc6, 0xdf, 0xde, 0xd4, 0x2c, 0x66, 0x8d, 0x53, 0x45,
	0x7a, 0x4b, 0xc3, 0xa9, 0x52, 0x1e, 0x4e, 0xe4, 0x3f, 0x68, 0xb3, 0xa9, 0xfa, 0xf5, 0xa0, 0x61,
	0x2c, 0x59, 0x1a, 0xb8, 0x1e, 0xcb, 0x8e, 0x75, 0x07, 0xe3, 0xa3, 0x3c, 0x4c, 0xee, 0xc0, 0x6e,
	0x18, 0xaf, 0x6a, 0xf1, 0x64, 0xdb, 0x59, 0x62, 0x21, 0x26, 0xb0, 0x99, 0x26, 0x51, 0x3e, 0x19,
	0xf4, 0x77, 0xc5, 0x84, 0xda, 0xa9, 0xda, 0xa1, 0x64, 0xda, 0xc7, 0xba, 0x63, 0xaa, 0xc8, 0x85,
	0x0a, 0x4c, 0xaa, 0xfa, 0xaf, 0xc1, 0xf1, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x32, 0xbb,
	0xa8, 0x3a, 0x08, 0x00, 0x00,
}
