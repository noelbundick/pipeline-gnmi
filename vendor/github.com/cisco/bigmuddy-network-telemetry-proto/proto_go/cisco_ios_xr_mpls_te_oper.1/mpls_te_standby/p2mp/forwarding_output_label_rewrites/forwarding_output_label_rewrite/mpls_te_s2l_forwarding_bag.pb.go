// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_s2l_forwarding_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_p2mp_forwarding_output_label_rewrites_forwarding_output_label_rewrite

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

// Output label forwarding details for S2Ls
type MplsTeS2LForwardingBag_KEYS struct {
	TunnelId             uint32   `protobuf:"varint,1,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	LspId                uint32   `protobuf:"varint,2,opt,name=lsp_id,json=lspId" json:"lsp_id,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,3,opt,name=extended_tunnel_id,json=extendedTunnelId" json:"extended_tunnel_id,omitempty"`
	SourceAddress        string   `protobuf:"bytes,4,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	P2MpId               uint32   `protobuf:"varint,5,opt,name=p2_mp_id,json=p2MpId" json:"p2_mp_id,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,6,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	NextHopAddress       string   `protobuf:"bytes,7,opt,name=next_hop_address,json=nextHopAddress" json:"next_hop_address,omitempty"`
	PreviousHopAddress   string   `protobuf:"bytes,8,opt,name=previous_hop_address,json=previousHopAddress" json:"previous_hop_address,omitempty"`
	CType                string   `protobuf:"bytes,9,opt,name=c_type,json=cType" json:"c_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeS2LForwardingBag_KEYS) Reset()         { *m = MplsTeS2LForwardingBag_KEYS{} }
func (m *MplsTeS2LForwardingBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeS2LForwardingBag_KEYS) ProtoMessage()    {}
func (*MplsTeS2LForwardingBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37, []int{0}
}
func (m *MplsTeS2LForwardingBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeS2LForwardingBag_KEYS.Unmarshal(m, b)
}
func (m *MplsTeS2LForwardingBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeS2LForwardingBag_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTeS2LForwardingBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeS2LForwardingBag_KEYS.Merge(dst, src)
}
func (m *MplsTeS2LForwardingBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeS2LForwardingBag_KEYS.Size(m)
}
func (m *MplsTeS2LForwardingBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeS2LForwardingBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeS2LForwardingBag_KEYS proto.InternalMessageInfo

func (m *MplsTeS2LForwardingBag_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *MplsTeS2LForwardingBag_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *MplsTeS2LForwardingBag_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *MplsTeS2LForwardingBag_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsTeS2LForwardingBag_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *MplsTeS2LForwardingBag_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTeS2LForwardingBag_KEYS) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *MplsTeS2LForwardingBag_KEYS) GetPreviousHopAddress() string {
	if m != nil {
		return m.PreviousHopAddress
	}
	return ""
}

func (m *MplsTeS2LForwardingBag_KEYS) GetCType() string {
	if m != nil {
		return m.CType
	}
	return ""
}

type MplsTeS2LForwardingBag struct {
	// Output rewrite shared between S2Ls
	S2LOutputRewrite *MplsTeS2LOutputRwBag `protobuf:"bytes,50,opt,name=s2_l_output_rewrite,json=s2LOutputRewrite" json:"s2_l_output_rewrite,omitempty"`
	// Input interface of the S2Ls
	OriginalInputInterface string `protobuf:"bytes,51,opt,name=original_input_interface,json=originalInputInterface" json:"original_input_interface,omitempty"`
	// The output interface of the S2Ls
	OutputInterfaceName string `protobuf:"bytes,52,opt,name=output_interface_name,json=outputInterfaceName" json:"output_interface_name,omitempty"`
	// Backup tunnel name
	BackupTunnelName string `protobuf:"bytes,53,opt,name=backup_tunnel_name,json=backupTunnelName" json:"backup_tunnel_name,omitempty"`
	// Unique identifiers for the S2L sharing the rewrite
	S2L []*MplsTeS2LForwardingS2LIdBag `protobuf:"bytes,54,rep,name=s2_l,json=s2L" json:"s2_l,omitempty"`
	// TRUE if the s2l path is Segment-Routing
	IsSegmentRouting bool `protobuf:"varint,55,opt,name=is_segment_routing,json=isSegmentRouting" json:"is_segment_routing,omitempty"`
	// Segment-Routing Paths
	S2LsrPaths           []*MplsTeSrS2LPathBag `protobuf:"bytes,56,rep,name=s2_lsr_paths,json=s2LsrPaths" json:"s2_lsr_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MplsTeS2LForwardingBag) Reset()         { *m = MplsTeS2LForwardingBag{} }
func (m *MplsTeS2LForwardingBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeS2LForwardingBag) ProtoMessage()    {}
func (*MplsTeS2LForwardingBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37, []int{1}
}
func (m *MplsTeS2LForwardingBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeS2LForwardingBag.Unmarshal(m, b)
}
func (m *MplsTeS2LForwardingBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeS2LForwardingBag.Marshal(b, m, deterministic)
}
func (dst *MplsTeS2LForwardingBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeS2LForwardingBag.Merge(dst, src)
}
func (m *MplsTeS2LForwardingBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeS2LForwardingBag.Size(m)
}
func (m *MplsTeS2LForwardingBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeS2LForwardingBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeS2LForwardingBag proto.InternalMessageInfo

func (m *MplsTeS2LForwardingBag) GetS2LOutputRewrite() *MplsTeS2LOutputRwBag {
	if m != nil {
		return m.S2LOutputRewrite
	}
	return nil
}

func (m *MplsTeS2LForwardingBag) GetOriginalInputInterface() string {
	if m != nil {
		return m.OriginalInputInterface
	}
	return ""
}

func (m *MplsTeS2LForwardingBag) GetOutputInterfaceName() string {
	if m != nil {
		return m.OutputInterfaceName
	}
	return ""
}

func (m *MplsTeS2LForwardingBag) GetBackupTunnelName() string {
	if m != nil {
		return m.BackupTunnelName
	}
	return ""
}

func (m *MplsTeS2LForwardingBag) GetS2L() []*MplsTeS2LForwardingS2LIdBag {
	if m != nil {
		return m.S2L
	}
	return nil
}

func (m *MplsTeS2LForwardingBag) GetIsSegmentRouting() bool {
	if m != nil {
		return m.IsSegmentRouting
	}
	return false
}

func (m *MplsTeS2LForwardingBag) GetS2LsrPaths() []*MplsTeSrS2LPathBag {
	if m != nil {
		return m.S2LsrPaths
	}
	return nil
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
	return fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37, []int{2}
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
	return fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37, []int{3}
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

// LMRIB S2L Output label rewrite FEC subfamily
type TeS2LOutputRwFecSubfamilyBag struct {
	// LSP FEC
	LspFec *TeLspFecT `protobuf:"bytes,1,opt,name=lsp_fec,json=lspFec" json:"lsp_fec,omitempty"`
	// Next hop address
	NextHopAddress string `protobuf:"bytes,2,opt,name=next_hop_address,json=nextHopAddress" json:"next_hop_address,omitempty"`
	// Previous hop address
	PreviousHopAddress   string   `protobuf:"bytes,3,opt,name=previous_hop_address,json=previousHopAddress" json:"previous_hop_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeS2LOutputRwFecSubfamilyBag) Reset()         { *m = TeS2LOutputRwFecSubfamilyBag{} }
func (m *TeS2LOutputRwFecSubfamilyBag) String() string { return proto.CompactTextString(m) }
func (*TeS2LOutputRwFecSubfamilyBag) ProtoMessage()    {}
func (*TeS2LOutputRwFecSubfamilyBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37, []int{4}
}
func (m *TeS2LOutputRwFecSubfamilyBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeS2LOutputRwFecSubfamilyBag.Unmarshal(m, b)
}
func (m *TeS2LOutputRwFecSubfamilyBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeS2LOutputRwFecSubfamilyBag.Marshal(b, m, deterministic)
}
func (dst *TeS2LOutputRwFecSubfamilyBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeS2LOutputRwFecSubfamilyBag.Merge(dst, src)
}
func (m *TeS2LOutputRwFecSubfamilyBag) XXX_Size() int {
	return xxx_messageInfo_TeS2LOutputRwFecSubfamilyBag.Size(m)
}
func (m *TeS2LOutputRwFecSubfamilyBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TeS2LOutputRwFecSubfamilyBag.DiscardUnknown(m)
}

var xxx_messageInfo_TeS2LOutputRwFecSubfamilyBag proto.InternalMessageInfo

func (m *TeS2LOutputRwFecSubfamilyBag) GetLspFec() *TeLspFecT {
	if m != nil {
		return m.LspFec
	}
	return nil
}

func (m *TeS2LOutputRwFecSubfamilyBag) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *TeS2LOutputRwFecSubfamilyBag) GetPreviousHopAddress() string {
	if m != nil {
		return m.PreviousHopAddress
	}
	return ""
}

// Output rewrite tunnel information
type TeS2LOutputRwOutputInfoBag struct {
	// Physical interface associated with the entry
	PhysicaInterfaceName string `protobuf:"bytes,1,opt,name=physica_interface_name,json=physicaInterfaceName" json:"physica_interface_name,omitempty"`
	// Interface handle of the associated tunnel
	TunnelInterfaceName string `protobuf:"bytes,2,opt,name=tunnel_interface_name,json=tunnelInterfaceName" json:"tunnel_interface_name,omitempty"`
	// Interface name of the parent intf
	ParentInterfaceName string `protobuf:"bytes,3,opt,name=parent_interface_name,json=parentInterfaceName" json:"parent_interface_name,omitempty"`
	// Next hop address associated with the entry
	NextHopAddress string `protobuf:"bytes,4,opt,name=next_hop_address,json=nextHopAddress" json:"next_hop_address,omitempty"`
	// Output label associated with the entry
	OutLabel uint32 `protobuf:"varint,5,opt,name=out_label,json=outLabel" json:"out_label,omitempty"`
	// Segment-Routing labels stack
	SrLabelStack         []uint32 `protobuf:"varint,6,rep,packed,name=sr_label_stack,json=srLabelStack" json:"sr_label_stack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeS2LOutputRwOutputInfoBag) Reset()         { *m = TeS2LOutputRwOutputInfoBag{} }
func (m *TeS2LOutputRwOutputInfoBag) String() string { return proto.CompactTextString(m) }
func (*TeS2LOutputRwOutputInfoBag) ProtoMessage()    {}
func (*TeS2LOutputRwOutputInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37, []int{5}
}
func (m *TeS2LOutputRwOutputInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeS2LOutputRwOutputInfoBag.Unmarshal(m, b)
}
func (m *TeS2LOutputRwOutputInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeS2LOutputRwOutputInfoBag.Marshal(b, m, deterministic)
}
func (dst *TeS2LOutputRwOutputInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeS2LOutputRwOutputInfoBag.Merge(dst, src)
}
func (m *TeS2LOutputRwOutputInfoBag) XXX_Size() int {
	return xxx_messageInfo_TeS2LOutputRwOutputInfoBag.Size(m)
}
func (m *TeS2LOutputRwOutputInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TeS2LOutputRwOutputInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_TeS2LOutputRwOutputInfoBag proto.InternalMessageInfo

func (m *TeS2LOutputRwOutputInfoBag) GetPhysicaInterfaceName() string {
	if m != nil {
		return m.PhysicaInterfaceName
	}
	return ""
}

func (m *TeS2LOutputRwOutputInfoBag) GetTunnelInterfaceName() string {
	if m != nil {
		return m.TunnelInterfaceName
	}
	return ""
}

func (m *TeS2LOutputRwOutputInfoBag) GetParentInterfaceName() string {
	if m != nil {
		return m.ParentInterfaceName
	}
	return ""
}

func (m *TeS2LOutputRwOutputInfoBag) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *TeS2LOutputRwOutputInfoBag) GetOutLabel() uint32 {
	if m != nil {
		return m.OutLabel
	}
	return 0
}

func (m *TeS2LOutputRwOutputInfoBag) GetSrLabelStack() []uint32 {
	if m != nil {
		return m.SrLabelStack
	}
	return nil
}

// An entry of output label rewrite data
type TeS2LOutputRwFieldsBag struct {
	// Time stamp of the entry (secs since 1/1/70)
	Timestamp uint32 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	// Role of the associated S2L
	OutRewriteRole string `protobuf:"bytes,2,opt,name=out_rewrite_role,json=outRewriteRole" json:"out_rewrite_role,omitempty"`
	// Local label of the rewrite
	Label uint32 `protobuf:"varint,3,opt,name=label" json:"label,omitempty"`
	// Properties of the S2L rewrite
	PrimaryS2L *TeS2LOutputRwOutputInfoBag `protobuf:"bytes,4,opt,name=primary_s2_l,json=primaryS2L" json:"primary_s2_l,omitempty"`
	// Properties of the S2L's backup tunnel rewrite
	BackupTunnelRewrite *TeS2LOutputRwOutputInfoBag `protobuf:"bytes,5,opt,name=backup_tunnel_rewrite,json=backupTunnelRewrite" json:"backup_tunnel_rewrite,omitempty"`
	// Backup is active
	BackupActive bool `protobuf:"varint,6,opt,name=backup_active,json=backupActive" json:"backup_active,omitempty"`
	// Source of the S2L
	S2LSource string `protobuf:"bytes,7,opt,name=s2_l_source,json=s2LSource" json:"s2_l_source,omitempty"`
	// Imposition of explicit NULL
	ExplicitNull string `protobuf:"bytes,8,opt,name=explicit_null,json=explicitNull" json:"explicit_null,omitempty"`
	// Protocol transported in the S2L
	ProtocolTransported  uint32   `protobuf:"varint,9,opt,name=protocol_transported,json=protocolTransported" json:"protocol_transported,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeS2LOutputRwFieldsBag) Reset()         { *m = TeS2LOutputRwFieldsBag{} }
func (m *TeS2LOutputRwFieldsBag) String() string { return proto.CompactTextString(m) }
func (*TeS2LOutputRwFieldsBag) ProtoMessage()    {}
func (*TeS2LOutputRwFieldsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37, []int{6}
}
func (m *TeS2LOutputRwFieldsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeS2LOutputRwFieldsBag.Unmarshal(m, b)
}
func (m *TeS2LOutputRwFieldsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeS2LOutputRwFieldsBag.Marshal(b, m, deterministic)
}
func (dst *TeS2LOutputRwFieldsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeS2LOutputRwFieldsBag.Merge(dst, src)
}
func (m *TeS2LOutputRwFieldsBag) XXX_Size() int {
	return xxx_messageInfo_TeS2LOutputRwFieldsBag.Size(m)
}
func (m *TeS2LOutputRwFieldsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TeS2LOutputRwFieldsBag.DiscardUnknown(m)
}

var xxx_messageInfo_TeS2LOutputRwFieldsBag proto.InternalMessageInfo

func (m *TeS2LOutputRwFieldsBag) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TeS2LOutputRwFieldsBag) GetOutRewriteRole() string {
	if m != nil {
		return m.OutRewriteRole
	}
	return ""
}

func (m *TeS2LOutputRwFieldsBag) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *TeS2LOutputRwFieldsBag) GetPrimaryS2L() *TeS2LOutputRwOutputInfoBag {
	if m != nil {
		return m.PrimaryS2L
	}
	return nil
}

func (m *TeS2LOutputRwFieldsBag) GetBackupTunnelRewrite() *TeS2LOutputRwOutputInfoBag {
	if m != nil {
		return m.BackupTunnelRewrite
	}
	return nil
}

func (m *TeS2LOutputRwFieldsBag) GetBackupActive() bool {
	if m != nil {
		return m.BackupActive
	}
	return false
}

func (m *TeS2LOutputRwFieldsBag) GetS2LSource() string {
	if m != nil {
		return m.S2LSource
	}
	return ""
}

func (m *TeS2LOutputRwFieldsBag) GetExplicitNull() string {
	if m != nil {
		return m.ExplicitNull
	}
	return ""
}

func (m *TeS2LOutputRwFieldsBag) GetProtocolTransported() uint32 {
	if m != nil {
		return m.ProtocolTransported
	}
	return 0
}

// S2L output rewrite
type MplsTeS2LOutputRwBag struct {
	// Subfamily identifiers
	Subfamily *TeS2LOutputRwFecSubfamilyBag `protobuf:"bytes,1,opt,name=subfamily" json:"subfamily,omitempty"`
	// Successful rewrite details
	SuccessfulRewrite *TeS2LOutputRwFieldsBag `protobuf:"bytes,2,opt,name=successful_rewrite,json=successfulRewrite" json:"successful_rewrite,omitempty"`
	// Failed rewrite details
	FailedRewrite *TeS2LOutputRwFieldsBag `protobuf:"bytes,3,opt,name=failed_rewrite,json=failedRewrite" json:"failed_rewrite,omitempty"`
	// Pending rewrite details
	PendingRewrite       *TeS2LOutputRwFieldsBag `protobuf:"bytes,4,opt,name=pending_rewrite,json=pendingRewrite" json:"pending_rewrite,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MplsTeS2LOutputRwBag) Reset()         { *m = MplsTeS2LOutputRwBag{} }
func (m *MplsTeS2LOutputRwBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeS2LOutputRwBag) ProtoMessage()    {}
func (*MplsTeS2LOutputRwBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37, []int{7}
}
func (m *MplsTeS2LOutputRwBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeS2LOutputRwBag.Unmarshal(m, b)
}
func (m *MplsTeS2LOutputRwBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeS2LOutputRwBag.Marshal(b, m, deterministic)
}
func (dst *MplsTeS2LOutputRwBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeS2LOutputRwBag.Merge(dst, src)
}
func (m *MplsTeS2LOutputRwBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeS2LOutputRwBag.Size(m)
}
func (m *MplsTeS2LOutputRwBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeS2LOutputRwBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeS2LOutputRwBag proto.InternalMessageInfo

func (m *MplsTeS2LOutputRwBag) GetSubfamily() *TeS2LOutputRwFecSubfamilyBag {
	if m != nil {
		return m.Subfamily
	}
	return nil
}

func (m *MplsTeS2LOutputRwBag) GetSuccessfulRewrite() *TeS2LOutputRwFieldsBag {
	if m != nil {
		return m.SuccessfulRewrite
	}
	return nil
}

func (m *MplsTeS2LOutputRwBag) GetFailedRewrite() *TeS2LOutputRwFieldsBag {
	if m != nil {
		return m.FailedRewrite
	}
	return nil
}

func (m *MplsTeS2LOutputRwBag) GetPendingRewrite() *TeS2LOutputRwFieldsBag {
	if m != nil {
		return m.PendingRewrite
	}
	return nil
}

// S2L Forwarding identifiers
type MplsTeS2LForwardingS2LIdBag struct {
	// Destination
	DestinationAddress string `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// S2L Sub-Group ID
	SubGroupId uint32 `protobuf:"varint,2,opt,name=sub_group_id,json=subGroupId" json:"sub_group_id,omitempty"`
	// S2L Sub-Group Originator ID
	SubGroupOriginalId   string   `protobuf:"bytes,3,opt,name=sub_group_original_id,json=subGroupOriginalId" json:"sub_group_original_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeS2LForwardingS2LIdBag) Reset()         { *m = MplsTeS2LForwardingS2LIdBag{} }
func (m *MplsTeS2LForwardingS2LIdBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeS2LForwardingS2LIdBag) ProtoMessage()    {}
func (*MplsTeS2LForwardingS2LIdBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37, []int{8}
}
func (m *MplsTeS2LForwardingS2LIdBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeS2LForwardingS2LIdBag.Unmarshal(m, b)
}
func (m *MplsTeS2LForwardingS2LIdBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeS2LForwardingS2LIdBag.Marshal(b, m, deterministic)
}
func (dst *MplsTeS2LForwardingS2LIdBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeS2LForwardingS2LIdBag.Merge(dst, src)
}
func (m *MplsTeS2LForwardingS2LIdBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeS2LForwardingS2LIdBag.Size(m)
}
func (m *MplsTeS2LForwardingS2LIdBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeS2LForwardingS2LIdBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeS2LForwardingS2LIdBag proto.InternalMessageInfo

func (m *MplsTeS2LForwardingS2LIdBag) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTeS2LForwardingS2LIdBag) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

func (m *MplsTeS2LForwardingS2LIdBag) GetSubGroupOriginalId() string {
	if m != nil {
		return m.SubGroupOriginalId
	}
	return ""
}

// Segment-Routing S2L path
type MplsTeSrS2LPathBag struct {
	// True if path is primary
	IsPrimary bool `protobuf:"varint,1,opt,name=is_primary,json=isPrimary" json:"is_primary,omitempty"`
	// True if path is backup
	IsBackup bool `protobuf:"varint,2,opt,name=is_backup,json=isBackup" json:"is_backup,omitempty"`
	// Outgoing Interface Name
	OutgoingInterface string `protobuf:"bytes,3,opt,name=outgoing_interface,json=outgoingInterface" json:"outgoing_interface,omitempty"`
	// Path Identifier
	PathId uint32 `protobuf:"varint,4,opt,name=path_id,json=pathId" json:"path_id,omitempty"`
	// Backup Path Identifier
	BackupPathId uint32 `protobuf:"varint,5,opt,name=backup_path_id,json=backupPathId" json:"backup_path_id,omitempty"`
	// Outgoing Labels Stack
	OutgoingLabelsStack []uint32 `protobuf:"varint,6,rep,packed,name=outgoing_labels_stack,json=outgoingLabelsStack" json:"outgoing_labels_stack,omitempty"`
	// Next hop
	NextHop              string   `protobuf:"bytes,7,opt,name=next_hop,json=nextHop" json:"next_hop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeSrS2LPathBag) Reset()         { *m = MplsTeSrS2LPathBag{} }
func (m *MplsTeSrS2LPathBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeSrS2LPathBag) ProtoMessage()    {}
func (*MplsTeSrS2LPathBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37, []int{9}
}
func (m *MplsTeSrS2LPathBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeSrS2LPathBag.Unmarshal(m, b)
}
func (m *MplsTeSrS2LPathBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeSrS2LPathBag.Marshal(b, m, deterministic)
}
func (dst *MplsTeSrS2LPathBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeSrS2LPathBag.Merge(dst, src)
}
func (m *MplsTeSrS2LPathBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeSrS2LPathBag.Size(m)
}
func (m *MplsTeSrS2LPathBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeSrS2LPathBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeSrS2LPathBag proto.InternalMessageInfo

func (m *MplsTeSrS2LPathBag) GetIsPrimary() bool {
	if m != nil {
		return m.IsPrimary
	}
	return false
}

func (m *MplsTeSrS2LPathBag) GetIsBackup() bool {
	if m != nil {
		return m.IsBackup
	}
	return false
}

func (m *MplsTeSrS2LPathBag) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *MplsTeSrS2LPathBag) GetPathId() uint32 {
	if m != nil {
		return m.PathId
	}
	return 0
}

func (m *MplsTeSrS2LPathBag) GetBackupPathId() uint32 {
	if m != nil {
		return m.BackupPathId
	}
	return 0
}

func (m *MplsTeSrS2LPathBag) GetOutgoingLabelsStack() []uint32 {
	if m != nil {
		return m.OutgoingLabelsStack
	}
	return nil
}

func (m *MplsTeSrS2LPathBag) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsTeS2LForwardingBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2mp.forwarding_output_label_rewrites.forwarding_output_label_rewrite.mpls_te_s2l_forwarding_bag_KEYS")
	proto.RegisterType((*MplsTeS2LForwardingBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2mp.forwarding_output_label_rewrites.forwarding_output_label_rewrite.mpls_te_s2l_forwarding_bag")
	proto.RegisterType((*TeLspFecCtypeDataT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2mp.forwarding_output_label_rewrites.forwarding_output_label_rewrite.te_lsp_fec_ctype_data_t")
	proto.RegisterType((*TeLspFecT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2mp.forwarding_output_label_rewrites.forwarding_output_label_rewrite.te_lsp_fec_t")
	proto.RegisterType((*TeS2LOutputRwFecSubfamilyBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2mp.forwarding_output_label_rewrites.forwarding_output_label_rewrite.te_s2l_output_rw_fec_subfamily_bag")
	proto.RegisterType((*TeS2LOutputRwOutputInfoBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2mp.forwarding_output_label_rewrites.forwarding_output_label_rewrite.te_s2l_output_rw_output_info_bag")
	proto.RegisterType((*TeS2LOutputRwFieldsBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2mp.forwarding_output_label_rewrites.forwarding_output_label_rewrite.te_s2l_output_rw_fields_bag")
	proto.RegisterType((*MplsTeS2LOutputRwBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2mp.forwarding_output_label_rewrites.forwarding_output_label_rewrite.mpls_te_s2l_output_rw_bag")
	proto.RegisterType((*MplsTeS2LForwardingS2LIdBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2mp.forwarding_output_label_rewrites.forwarding_output_label_rewrite.mpls_te_s2l_forwarding_s2l_id_bag")
	proto.RegisterType((*MplsTeSrS2LPathBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.p2mp.forwarding_output_label_rewrites.forwarding_output_label_rewrite.mpls_te_sr_s2l_path_bag")
}

func init() {
	proto.RegisterFile("mpls_te_s2l_forwarding_bag.proto", fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37)
}

var fileDescriptor_mpls_te_s2l_forwarding_bag_95c2ffceaa55fe37 = []byte{
	// 1316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xcf, 0x6f, 0x1c, 0xc5,
	0x12, 0xd6, 0xec, 0x7a, 0xed, 0xdd, 0xb2, 0xd7, 0xcf, 0xe9, 0xb5, 0xe3, 0xcd, 0x4b, 0xde, 0x7b,
	0xfb, 0x36, 0x20, 0xed, 0x01, 0x1c, 0x32, 0x09, 0x21, 0xd7, 0x00, 0x01, 0x56, 0x2c, 0x89, 0x35,
	0x8e, 0x90, 0x38, 0xb5, 0x66, 0x67, 0x7a, 0xd6, 0xad, 0xcc, 0x4e, 0xb7, 0xba, 0x7b, 0x12, 0xfb,
	0xc0, 0x1d, 0x21, 0x2e, 0x04, 0x09, 0xc1, 0x8d, 0x13, 0x12, 0x42, 0x5c, 0x50, 0xc4, 0x01, 0x38,
	0xf3, 0xb7, 0xf0, 0x37, 0xc0, 0x01, 0xd4, 0xbf, 0x66, 0x77, 0xed, 0x2c, 0x11, 0x17, 0xec, 0xe3,
	0x54, 0x7d, 0xd5, 0xf3, 0x75, 0x75, 0xf5, 0x57, 0xd5, 0xd0, 0x9b, 0xf2, 0x5c, 0x62, 0x45, 0xb0,
	0x0c, 0x73, 0x9c, 0x31, 0xf1, 0x38, 0x16, 0x29, 0x2d, 0x26, 0x78, 0x1c, 0x4f, 0xf6, 0xb8, 0x60,
	0x8a, 0x21, 0x96, 0x50, 0x99, 0x30, 0x4c, 0x99, 0xc4, 0x47, 0x02, 0x7b, 0x38, 0xe3, 0x44, 0xec,
	0x55, 0xb1, 0x2a, 0x2e, 0xd2, 0xf1, 0xf1, 0x1e, 0x0f, 0xa7, 0x7c, 0x6f, 0x6e, 0x11, 0x56, 0x2a,
	0x5e, 0x2a, 0x9c, 0xc7, 0x63, 0x92, 0x63, 0x41, 0x1e, 0x0b, 0xaa, 0x88, 0x7c, 0x1e, 0xa0, 0xff,
	0x6b, 0x0d, 0xfe, 0xb7, 0x9c, 0x15, 0x7e, 0xf7, 0xee, 0x07, 0x07, 0xe8, 0x32, 0xb4, 0x54, 0x59,
	0x14, 0x24, 0xc7, 0x34, 0xed, 0x06, 0xbd, 0x60, 0xd0, 0x8e, 0x9a, 0xd6, 0x30, 0x4c, 0xd1, 0x0e,
	0xac, 0xe6, 0x92, 0x6b, 0x4f, 0xcd, 0x78, 0x1a, 0xb9, 0xe4, 0xc3, 0x14, 0xbd, 0x04, 0x88, 0x1c,
	0x29, 0x52, 0xa4, 0x24, 0xc5, 0xb3, 0xe0, 0x7a, 0x2f, 0x18, 0xb4, 0xa2, 0x2d, 0xef, 0x79, 0xe0,
	0x17, 0x79, 0x11, 0x36, 0x25, 0x2b, 0x45, 0x42, 0x70, 0x9c, 0xa6, 0x82, 0x48, 0xd9, 0x5d, 0x31,
	0xc8, 0xb6, 0xb5, 0xde, 0xb1, 0x46, 0xd4, 0x85, 0x26, 0x0f, 0xf1, 0xd4, 0xfc, 0xad, 0x61, 0xfe,
	0xb6, 0xca, 0xc3, 0xf7, 0xf4, 0xef, 0xae, 0x41, 0x27, 0x25, 0x52, 0xd1, 0x22, 0x56, 0x94, 0x15,
	0xd5, 0x2a, 0xab, 0x66, 0x15, 0x34, 0xe7, 0xf2, 0x4b, 0x0d, 0x60, 0xab, 0x20, 0x47, 0x0a, 0x1f,
	0x32, 0x5e, 0xa1, 0xd7, 0x0c, 0x7a, 0x53, 0xdb, 0xdf, 0x61, 0xdc, 0x23, 0x5f, 0x81, 0x6d, 0x2e,
	0xc8, 0x23, 0xca, 0x4a, 0xb9, 0x80, 0x6e, 0xda, 0xb5, 0xbd, 0x6f, 0x2e, 0x62, 0x07, 0x56, 0x13,
	0xac, 0x8e, 0x39, 0xe9, 0xb6, 0x0c, 0xa6, 0x91, 0x3c, 0x38, 0xe6, 0xa4, 0xff, 0x5b, 0x03, 0xfe,
	0xbd, 0x3c, 0xd5, 0xe8, 0xc7, 0x00, 0x3a, 0x32, 0xc4, 0xb9, 0x3f, 0x27, 0x77, 0x42, 0xdd, 0xb0,
	0x17, 0x0c, 0xd6, 0xc3, 0x8f, 0x83, 0xbd, 0x7f, 0xb8, 0x34, 0xf6, 0xe6, 0xb9, 0x7a, 0x4e, 0x8f,
	0x35, 0xd5, 0x68, 0x4b, 0x86, 0xa3, 0xfb, 0xc6, 0x12, 0x59, 0x2c, 0xba, 0x0d, 0x5d, 0x26, 0xe8,
	0x84, 0x16, 0x71, 0x8e, 0x69, 0xa1, 0xa1, 0xb4, 0x50, 0x44, 0x64, 0x71, 0x42, 0xba, 0x37, 0x4c,
	0x12, 0x2e, 0x7a, 0xff, 0x50, 0xbb, 0x87, 0xde, 0x8b, 0x42, 0xd8, 0x71, 0x8b, 0x57, 0x11, 0xb8,
	0x88, 0xa7, 0xa4, 0x7b, 0xd3, 0x84, 0x75, 0xac, 0xb3, 0xc2, 0xdf, 0x8b, 0xa7, 0x44, 0x17, 0xd7,
	0x38, 0x4e, 0x1e, 0x96, 0xdc, 0x97, 0x96, 0x09, 0x78, 0xd5, 0x16, 0x97, 0xf5, 0xd8, 0xd2, 0x32,
	0xe8, 0xaf, 0x03, 0x58, 0xd1, 0x89, 0xed, 0xde, 0xea, 0xd5, 0x07, 0xeb, 0xe1, 0x93, 0xb3, 0xcd,
	0xe4, 0x1c, 0x56, 0x7f, 0xd2, 0xd4, 0x64, 0xb4, 0x2e, 0xc3, 0x91, 0xde, 0x16, 0x95, 0x58, 0x92,
	0xc9, 0x94, 0x14, 0x0a, 0x0b, 0x56, 0x2a, 0x5a, 0x4c, 0xba, 0xaf, 0xf5, 0x82, 0x41, 0x33, 0xda,
	0xa2, 0xf2, 0xc0, 0x3a, 0x22, 0x6b, 0x47, 0xdf, 0x06, 0xb0, 0xa1, 0xb7, 0x25, 0x05, 0xe6, 0xb1,
	0x3a, 0x94, 0xdd, 0xdb, 0x66, 0x7b, 0x1f, 0x9d, 0xe1, 0xf6, 0x84, 0xd9, 0x92, 0x66, 0x63, 0x36,
	0x05, 0x32, 0x1c, 0x49, 0xb1, 0xaf, 0xc9, 0xf5, 0xbf, 0x0a, 0x60, 0x57, 0x11, 0xac, 0xa5, 0x22,
	0x23, 0x09, 0x4e, 0xf4, 0xf5, 0xc0, 0x69, 0xac, 0x62, 0xac, 0xd0, 0x15, 0x00, 0x63, 0xb3, 0x77,
	0x26, 0x30, 0xc7, 0xd8, 0xcc, 0x48, 0xf2, 0x86, 0xbe, 0x36, 0xe8, 0x1a, 0x6c, 0xf3, 0x10, 0x73,
	0x13, 0x3a, 0x77, 0x91, 0x8d, 0xdc, 0xb4, 0xa2, 0x0b, 0x3c, 0xdc, 0x1f, 0x49, 0xfe, 0xe6, 0xcc,
	0x81, 0x6e, 0x41, 0x57, 0x2f, 0x37, 0xaf, 0x07, 0x95, 0x6a, 0xd4, 0x8d, 0x6a, 0x6c, 0x67, 0x24,
	0x99, 0x8b, 0xd8, 0x37, 0x1a, 0xd2, 0xff, 0xbd, 0x06, 0x1b, 0x73, 0x14, 0x2b, 0x5e, 0x4e, 0xde,
	0x9c, 0xf0, 0x65, 0x24, 0x19, 0x19, 0x85, 0xeb, 0x43, 0xdb, 0xc0, 0x2a, 0x71, 0xb3, 0xfa, 0xb7,
	0x9e, 0x91, 0xa4, 0xd2, 0xb5, 0x1b, 0x70, 0x51, 0x63, 0x96, 0x2a, 0x61, 0x27, 0x23, 0xc9, 0xdd,
	0x93, 0x62, 0xf8, 0x1f, 0xfb, 0x5b, 0x2b, 0x7d, 0x4e, 0x08, 0x5b, 0x19, 0x49, 0x0e, 0x8c, 0x01,
	0xfd, 0x14, 0xc0, 0xf6, 0xc9, 0xfd, 0xd1, 0x22, 0x63, 0x46, 0x11, 0xcf, 0xe4, 0xfc, 0x97, 0x9c,
	0x6b, 0x84, 0x16, 0xd3, 0x3c, 0x2c, 0x32, 0x86, 0x76, 0x61, 0x4d, 0xe3, 0x1e, 0x89, 0xcc, 0x89,
	0xf3, 0x6a, 0x46, 0x92, 0xf7, 0x45, 0xd6, 0xff, 0xb2, 0x06, 0xfd, 0x53, 0x62, 0x63, 0xf2, 0x50,
	0x8e, 0xb3, 0x78, 0x4a, 0xf3, 0x63, 0xa3, 0x92, 0x9f, 0x07, 0xb0, 0xe6, 0x7e, 0x66, 0x4e, 0x64,
	0x3d, 0xfc, 0xf0, 0x2c, 0xf7, 0xab, 0x22, 0xdd, 0xfd, 0xde, 0x22, 0xc9, 0x33, 0x1b, 0x4a, 0xed,
	0x6f, 0x35, 0x94, 0xfa, 0xb2, 0x86, 0xd2, 0xff, 0xae, 0x06, 0xbd, 0x53, 0xb9, 0xa9, 0x54, 0x33,
	0x63, 0x26, 0x33, 0x37, 0xe1, 0x22, 0x3f, 0x3c, 0x96, 0x34, 0x89, 0x4f, 0x2a, 0xa9, 0xbd, 0x51,
	0xdb, 0xce, 0xbb, 0x28, 0xa5, 0x21, 0xec, 0xf8, 0xa2, 0x5c, 0x0c, 0xb2, 0xdc, 0x3b, 0xae, 0xcf,
	0x9f, 0x8c, 0xe1, 0xb1, 0xd0, 0x1a, 0x75, 0x22, 0xc6, 0x15, 0xb5, 0x75, 0x2e, 0xc6, 0x3c, 0x2b,
	0x3d, 0x2b, 0xcf, 0x4c, 0xcf, 0x65, 0x68, 0x31, 0x9f, 0x72, 0xd7, 0xe5, 0x9b, 0xac, 0x54, 0x23,
	0xfd, 0x8d, 0x5e, 0x80, 0x4d, 0x29, 0xdc, 0x71, 0x48, 0x15, 0x27, 0x0f, 0xbb, 0xab, 0xbd, 0xfa,
	0xa0, 0x1d, 0x6d, 0x48, 0x61, 0x00, 0x07, 0xda, 0xd6, 0xff, 0xa2, 0x01, 0x97, 0x4f, 0xd7, 0x12,
	0x25, 0x79, 0x2a, 0x4d, 0xaa, 0xae, 0x40, 0x4b, 0xd1, 0x29, 0x91, 0x2a, 0x9e, 0x72, 0x77, 0xaf,
	0x67, 0x06, 0x4d, 0x95, 0xcd, 0xfa, 0x2f, 0x16, 0x2c, 0xf7, 0xd9, 0xd8, 0x64, 0x55, 0xc7, 0x8b,
	0x58, 0x4e, 0xd0, 0x36, 0x34, 0x2c, 0xcd, 0xba, 0x1b, 0x7d, 0x0c, 0xc7, 0xa7, 0x01, 0x6c, 0x70,
	0x41, 0xa7, 0xb1, 0x38, 0xc6, 0xa6, 0xef, 0xac, 0x98, 0x3a, 0xfd, 0xf4, 0x4c, 0x2e, 0xe6, 0x5f,
	0xd6, 0x4c, 0x04, 0x8e, 0xe7, 0x41, 0x38, 0x42, 0xbf, 0x04, 0xb0, 0xb3, 0xd8, 0x55, 0xfd, 0x04,
	0xd2, 0x38, 0xb7, 0xfc, 0x3b, 0xf3, 0xcd, 0xde, 0xcf, 0x22, 0x57, 0xa1, 0xed, 0xf6, 0x11, 0x27,
	0x8a, 0x3e, 0x22, 0x46, 0x68, 0x9a, 0xd1, 0x86, 0x35, 0xde, 0x31, 0x36, 0xf4, 0x5f, 0x58, 0x37,
	0xc3, 0x96, 0x53, 0x59, 0x3b, 0xfa, 0xb5, 0x64, 0x38, 0x72, 0x2a, 0x7b, 0x15, 0xda, 0xe4, 0x88,
	0xe7, 0x34, 0xa1, 0x0a, 0x17, 0x65, 0x9e, 0xbb, 0x71, 0x6f, 0xc3, 0x1b, 0xef, 0x95, 0x79, 0x8e,
	0xae, 0xeb, 0x9b, 0xcc, 0x14, 0x4b, 0x58, 0x8e, 0x95, 0x88, 0x0b, 0xc9, 0x99, 0x50, 0x24, 0x35,
	0x63, 0x5f, 0x3b, 0xea, 0x78, 0xdf, 0x83, 0x99, 0xab, 0xff, 0x47, 0x03, 0x2e, 0x2d, 0x1d, 0xac,
	0xd0, 0xf7, 0x01, 0xb4, 0x2a, 0xbd, 0x73, 0xfa, 0xf6, 0xd9, 0x39, 0xc8, 0xfb, 0x29, 0x1d, 0x8e,
	0x66, 0x34, 0xd1, 0xcf, 0x01, 0x20, 0x59, 0x26, 0x09, 0x91, 0x32, 0x2b, 0x67, 0x55, 0x53, 0x33,
	0xec, 0x3f, 0x39, 0x0f, 0xec, 0xab, 0x9b, 0x1f, 0x5d, 0x98, 0x11, 0xf5, 0xe5, 0xf2, 0x34, 0x80,
	0xcd, 0x2c, 0xa6, 0x39, 0x49, 0x2b, 0xea, 0xf5, 0xf3, 0x48, 0xbd, 0x6d, 0x49, 0x7a, 0xda, 0x3f,
	0x04, 0xf0, 0x2f, 0x4e, 0x0a, 0xb3, 0x96, 0xe7, 0xbd, 0x72, 0x1e, 0x79, 0x6f, 0x3a, 0x96, 0x8e,
	0x78, 0xff, 0x9b, 0x00, 0xfe, 0xff, 0xdc, 0x81, 0x78, 0xd9, 0x83, 0x2e, 0x58, 0xfa, 0xa0, 0xeb,
	0xc1, 0x86, 0x2c, 0xc7, 0x78, 0x22, 0x58, 0x39, 0xf7, 0x1a, 0x05, 0x59, 0x8e, 0xdf, 0xd6, 0xa6,
	0x61, 0x8a, 0xae, 0xc3, 0xce, 0x0c, 0x31, 0x7b, 0xad, 0xf8, 0x59, 0x0c, 0x79, 0xe8, 0x7d, 0xff,
	0x50, 0x49, 0xfb, 0x4f, 0x6a, 0xb0, 0xbb, 0x64, 0xba, 0xd5, 0x63, 0x1a, 0x95, 0xd8, 0x09, 0xa8,
	0x21, 0xd6, 0x8c, 0x5a, 0x54, 0xee, 0x5b, 0x83, 0x6e, 0x63, 0x54, 0x27, 0x40, 0x6b, 0x8e, 0x21,
	0xd3, 0x8c, 0x9a, 0x54, 0xbe, 0x6e, 0xbe, 0xd1, 0xcb, 0x80, 0x58, 0xa9, 0x26, 0x4c, 0x6f, 0x7a,
	0xf6, 0x50, 0xb2, 0x3c, 0x2e, 0x78, 0xcf, 0xec, 0x8d, 0xb4, 0x0b, 0x6b, 0xe6, 0xb7, 0x34, 0x35,
	0x47, 0xac, 0x9f, 0xbd, 0xb1, 0x3a, 0x1c, 0xa6, 0xba, 0x1d, 0x3a, 0xa9, 0xf3, 0x7e, 0xdb, 0x30,
	0x9d, 0xd6, 0xed, 0x5b, 0x94, 0x7d, 0x62, 0xd9, 0xbf, 0x99, 0x73, 0x93, 0x0b, 0xbd, 0xb3, 0xe3,
	0x9d, 0xa6, 0x83, 0x4a, 0xd3, 0x42, 0xd1, 0x25, 0x68, 0xfa, 0x7e, 0xed, 0xc4, 0x71, 0xcd, 0xf5,
	0xe9, 0xf1, 0xaa, 0xd1, 0xb5, 0x1b, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xda, 0xae, 0x30, 0x63,
	0xce, 0x10, 0x00, 0x00,
}