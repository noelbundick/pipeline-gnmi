// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_mgmt_frr_states_detail.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_frrs_frr

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

// RSVP states for one FRR LSP
type RsvpMgmtFrrStatesDetail_KEYS struct {
	DestinationAddress   string   `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	DestinationPort      uint32   `protobuf:"varint,2,opt,name=destination_port,json=destinationPort" json:"destination_port,omitempty"`
	Protocol             uint32   `protobuf:"varint,3,opt,name=protocol" json:"protocol,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,4,opt,name=extended_tunnel_id,json=extendedTunnelId" json:"extended_tunnel_id,omitempty"`
	SessionType          string   `protobuf:"bytes,5,opt,name=session_type,json=sessionType" json:"session_type,omitempty"`
	P2MpId               uint32   `protobuf:"varint,6,opt,name=p2_mp_id,json=p2MpId" json:"p2_mp_id,omitempty"`
	SourceAddress        string   `protobuf:"bytes,7,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	SourcePort           uint32   `protobuf:"varint,8,opt,name=source_port,json=sourcePort" json:"source_port,omitempty"`
	SubGroupOrigin       string   `protobuf:"bytes,9,opt,name=sub_group_origin,json=subGroupOrigin" json:"sub_group_origin,omitempty"`
	SubGroupId           uint32   `protobuf:"varint,10,opt,name=sub_group_id,json=subGroupId" json:"sub_group_id,omitempty"`
	VrfName              string   `protobuf:"bytes,11,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) Reset()         { *m = RsvpMgmtFrrStatesDetail_KEYS{} }
func (m *RsvpMgmtFrrStatesDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtFrrStatesDetail_KEYS) ProtoMessage()    {}
func (*RsvpMgmtFrrStatesDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_frr_states_detail_8f4900b5d76805cc, []int{0}
}
func (m *RsvpMgmtFrrStatesDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtFrrStatesDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtFrrStatesDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS.Merge(dst, src)
}
func (m *RsvpMgmtFrrStatesDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS.Size(m)
}
func (m *RsvpMgmtFrrStatesDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetDestinationPort() uint32 {
	if m != nil {
		return m.DestinationPort
	}
	return 0
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetSessionType() string {
	if m != nil {
		return m.SessionType
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetSourcePort() uint32 {
	if m != nil {
		return m.SourcePort
	}
	return 0
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetSubGroupOrigin() string {
	if m != nil {
		return m.SubGroupOrigin
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type RsvpMgmtFrrStatesDetail struct {
	// RSVP Session Information
	Session *RsvpMgmtSessionInfo `protobuf:"bytes,50,opt,name=session" json:"session,omitempty"`
	// RSVP S2L Sub-LSP information
	S2LSubLsp *RsvpMgmtS2LSubLspIpv4 `protobuf:"bytes,51,opt,name=s2_l_sub_lsp,json=s2LSubLsp" json:"s2_l_sub_lsp,omitempty"`
	// RSVP FRR Path States
	PathStatus string `protobuf:"bytes,52,opt,name=path_status,json=pathStatus" json:"path_status,omitempty"`
	// RSVP FRR Reservation States
	ReservationStatus    string   `protobuf:"bytes,53,opt,name=reservation_status,json=reservationStatus" json:"reservation_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtFrrStatesDetail) Reset()         { *m = RsvpMgmtFrrStatesDetail{} }
func (m *RsvpMgmtFrrStatesDetail) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtFrrStatesDetail) ProtoMessage()    {}
func (*RsvpMgmtFrrStatesDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_frr_states_detail_8f4900b5d76805cc, []int{1}
}
func (m *RsvpMgmtFrrStatesDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail.Unmarshal(m, b)
}
func (m *RsvpMgmtFrrStatesDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtFrrStatesDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtFrrStatesDetail.Merge(dst, src)
}
func (m *RsvpMgmtFrrStatesDetail) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail.Size(m)
}
func (m *RsvpMgmtFrrStatesDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtFrrStatesDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtFrrStatesDetail proto.InternalMessageInfo

func (m *RsvpMgmtFrrStatesDetail) GetSession() *RsvpMgmtSessionInfo {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *RsvpMgmtFrrStatesDetail) GetS2LSubLsp() *RsvpMgmtS2LSubLspIpv4 {
	if m != nil {
		return m.S2LSubLsp
	}
	return nil
}

func (m *RsvpMgmtFrrStatesDetail) GetPathStatus() string {
	if m != nil {
		return m.PathStatus
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail) GetReservationStatus() string {
	if m != nil {
		return m.ReservationStatus
	}
	return ""
}

// RSVP S2L Sub-LSP
type RsvpMgmtS2LSubLspIpv4 struct {
	// S2L Sub-LSP Destination Address
	S2LDestinationAddress string   `protobuf:"bytes,1,opt,name=s2_l_destination_address,json=s2LDestinationAddress" json:"s2_l_destination_address,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *RsvpMgmtS2LSubLspIpv4) Reset()         { *m = RsvpMgmtS2LSubLspIpv4{} }
func (m *RsvpMgmtS2LSubLspIpv4) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtS2LSubLspIpv4) ProtoMessage()    {}
func (*RsvpMgmtS2LSubLspIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_frr_states_detail_8f4900b5d76805cc, []int{2}
}
func (m *RsvpMgmtS2LSubLspIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtS2LSubLspIpv4.Unmarshal(m, b)
}
func (m *RsvpMgmtS2LSubLspIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtS2LSubLspIpv4.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtS2LSubLspIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtS2LSubLspIpv4.Merge(dst, src)
}
func (m *RsvpMgmtS2LSubLspIpv4) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtS2LSubLspIpv4.Size(m)
}
func (m *RsvpMgmtS2LSubLspIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtS2LSubLspIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtS2LSubLspIpv4 proto.InternalMessageInfo

func (m *RsvpMgmtS2LSubLspIpv4) GetS2LDestinationAddress() string {
	if m != nil {
		return m.S2LDestinationAddress
	}
	return ""
}

// RSVP UDP IPv4 Session
type RsvpMgmtSessionUdpIpv4 struct {
	// Destination address
	DestinationAddress string `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// Protocol type (originally defined in RFC 790, further values in subsequent RFCs)
	Protocol uint32 `protobuf:"varint,2,opt,name=protocol" json:"protocol,omitempty"`
	// The Session Destination Port
	DestinationPort      uint32   `protobuf:"varint,3,opt,name=destination_port,json=destinationPort" json:"destination_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtSessionUdpIpv4) Reset()         { *m = RsvpMgmtSessionUdpIpv4{} }
func (m *RsvpMgmtSessionUdpIpv4) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtSessionUdpIpv4) ProtoMessage()    {}
func (*RsvpMgmtSessionUdpIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_frr_states_detail_8f4900b5d76805cc, []int{3}
}
func (m *RsvpMgmtSessionUdpIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtSessionUdpIpv4.Unmarshal(m, b)
}
func (m *RsvpMgmtSessionUdpIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtSessionUdpIpv4.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtSessionUdpIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtSessionUdpIpv4.Merge(dst, src)
}
func (m *RsvpMgmtSessionUdpIpv4) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtSessionUdpIpv4.Size(m)
}
func (m *RsvpMgmtSessionUdpIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtSessionUdpIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtSessionUdpIpv4 proto.InternalMessageInfo

func (m *RsvpMgmtSessionUdpIpv4) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RsvpMgmtSessionUdpIpv4) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *RsvpMgmtSessionUdpIpv4) GetDestinationPort() uint32 {
	if m != nil {
		return m.DestinationPort
	}
	return 0
}

// RSVP LSP-Tunnel-IPv4 Session
type RsvpMgmtSessionLspTunnelIpv4 struct {
	// Destination address
	DestinationAddress string `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// The Session Tunnel ID
	TunnelId uint32 `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	// The Session Extended Tunnel ID
	ExtendedTunnelId     string   `protobuf:"bytes,3,opt,name=extended_tunnel_id,json=extendedTunnelId" json:"extended_tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtSessionLspTunnelIpv4) Reset()         { *m = RsvpMgmtSessionLspTunnelIpv4{} }
func (m *RsvpMgmtSessionLspTunnelIpv4) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtSessionLspTunnelIpv4) ProtoMessage()    {}
func (*RsvpMgmtSessionLspTunnelIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_frr_states_detail_8f4900b5d76805cc, []int{4}
}
func (m *RsvpMgmtSessionLspTunnelIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4.Unmarshal(m, b)
}
func (m *RsvpMgmtSessionLspTunnelIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtSessionLspTunnelIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4.Merge(dst, src)
}
func (m *RsvpMgmtSessionLspTunnelIpv4) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4.Size(m)
}
func (m *RsvpMgmtSessionLspTunnelIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4 proto.InternalMessageInfo

func (m *RsvpMgmtSessionLspTunnelIpv4) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RsvpMgmtSessionLspTunnelIpv4) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *RsvpMgmtSessionLspTunnelIpv4) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

// RSVP UNI IPv4 Session
type RsvpMgmtSessionUniIpv4 struct {
	// Destination address
	DestinationAddress string `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// The Session Tunnel ID
	TunnelId uint32 `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	// The Session Extended Address
	ExtendedAddress      string   `protobuf:"bytes,3,opt,name=extended_address,json=extendedAddress" json:"extended_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtSessionUniIpv4) Reset()         { *m = RsvpMgmtSessionUniIpv4{} }
func (m *RsvpMgmtSessionUniIpv4) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtSessionUniIpv4) ProtoMessage()    {}
func (*RsvpMgmtSessionUniIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_frr_states_detail_8f4900b5d76805cc, []int{5}
}
func (m *RsvpMgmtSessionUniIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtSessionUniIpv4.Unmarshal(m, b)
}
func (m *RsvpMgmtSessionUniIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtSessionUniIpv4.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtSessionUniIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtSessionUniIpv4.Merge(dst, src)
}
func (m *RsvpMgmtSessionUniIpv4) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtSessionUniIpv4.Size(m)
}
func (m *RsvpMgmtSessionUniIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtSessionUniIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtSessionUniIpv4 proto.InternalMessageInfo

func (m *RsvpMgmtSessionUniIpv4) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RsvpMgmtSessionUniIpv4) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *RsvpMgmtSessionUniIpv4) GetExtendedAddress() string {
	if m != nil {
		return m.ExtendedAddress
	}
	return ""
}

// RSVP P2MP-LSP-Tunnel-IPv4 Session
type RsvpMgmtSessionP2MpLspTunnelIpv4 struct {
	// The Point to Multipoint ID
	P2MpId uint32 `protobuf:"varint,1,opt,name=p2_mp_id,json=p2MpId" json:"p2_mp_id,omitempty"`
	// The Session Tunnel ID
	TunnelId uint32 `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	// The Session Extended Tunnel ID
	ExtendedTunnelId     string   `protobuf:"bytes,3,opt,name=extended_tunnel_id,json=extendedTunnelId" json:"extended_tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) Reset()         { *m = RsvpMgmtSessionP2MpLspTunnelIpv4{} }
func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtSessionP2MpLspTunnelIpv4) ProtoMessage()    {}
func (*RsvpMgmtSessionP2MpLspTunnelIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_frr_states_detail_8f4900b5d76805cc, []int{6}
}
func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4.Unmarshal(m, b)
}
func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtSessionP2MpLspTunnelIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4.Merge(dst, src)
}
func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4.Size(m)
}
func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4 proto.InternalMessageInfo

func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

// Union of the different RSVP session types
type RsvpSessionUnion struct {
	SessionType string `protobuf:"bytes,1,opt,name=session_type,json=sessionType" json:"session_type,omitempty"`
	// UDP IPv4 session
	Ipv4 *RsvpMgmtSessionUdpIpv4 `protobuf:"bytes,2,opt,name=ipv4" json:"ipv4,omitempty"`
	// IPv4 LSP session
	Ipv4LspSession *RsvpMgmtSessionLspTunnelIpv4 `protobuf:"bytes,3,opt,name=ipv4_lsp_session,json=ipv4LspSession" json:"ipv4_lsp_session,omitempty"`
	// IPv4 UNI session
	Ipv4UniSession *RsvpMgmtSessionUniIpv4 `protobuf:"bytes,4,opt,name=ipv4_uni_session,json=ipv4UniSession" json:"ipv4_uni_session,omitempty"`
	// IPv4 P2MP LSP session
	Ipv4P2MpLspSession   *RsvpMgmtSessionP2MpLspTunnelIpv4 `protobuf:"bytes,5,opt,name=ipv4_p2_mp_lsp_session,json=ipv4P2MpLspSession" json:"ipv4_p2_mp_lsp_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *RsvpSessionUnion) Reset()         { *m = RsvpSessionUnion{} }
func (m *RsvpSessionUnion) String() string { return proto.CompactTextString(m) }
func (*RsvpSessionUnion) ProtoMessage()    {}
func (*RsvpSessionUnion) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_frr_states_detail_8f4900b5d76805cc, []int{7}
}
func (m *RsvpSessionUnion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpSessionUnion.Unmarshal(m, b)
}
func (m *RsvpSessionUnion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpSessionUnion.Marshal(b, m, deterministic)
}
func (dst *RsvpSessionUnion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpSessionUnion.Merge(dst, src)
}
func (m *RsvpSessionUnion) XXX_Size() int {
	return xxx_messageInfo_RsvpSessionUnion.Size(m)
}
func (m *RsvpSessionUnion) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpSessionUnion.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpSessionUnion proto.InternalMessageInfo

func (m *RsvpSessionUnion) GetSessionType() string {
	if m != nil {
		return m.SessionType
	}
	return ""
}

func (m *RsvpSessionUnion) GetIpv4() *RsvpMgmtSessionUdpIpv4 {
	if m != nil {
		return m.Ipv4
	}
	return nil
}

func (m *RsvpSessionUnion) GetIpv4LspSession() *RsvpMgmtSessionLspTunnelIpv4 {
	if m != nil {
		return m.Ipv4LspSession
	}
	return nil
}

func (m *RsvpSessionUnion) GetIpv4UniSession() *RsvpMgmtSessionUniIpv4 {
	if m != nil {
		return m.Ipv4UniSession
	}
	return nil
}

func (m *RsvpSessionUnion) GetIpv4P2MpLspSession() *RsvpMgmtSessionP2MpLspTunnelIpv4 {
	if m != nil {
		return m.Ipv4P2MpLspSession
	}
	return nil
}

// RSVP Session Info
type RsvpMgmtSessionInfo struct {
	// RSVP Session
	RsvpSession          *RsvpSessionUnion `protobuf:"bytes,1,opt,name=rsvp_session,json=rsvpSession" json:"rsvp_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RsvpMgmtSessionInfo) Reset()         { *m = RsvpMgmtSessionInfo{} }
func (m *RsvpMgmtSessionInfo) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtSessionInfo) ProtoMessage()    {}
func (*RsvpMgmtSessionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_frr_states_detail_8f4900b5d76805cc, []int{8}
}
func (m *RsvpMgmtSessionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtSessionInfo.Unmarshal(m, b)
}
func (m *RsvpMgmtSessionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtSessionInfo.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtSessionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtSessionInfo.Merge(dst, src)
}
func (m *RsvpMgmtSessionInfo) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtSessionInfo.Size(m)
}
func (m *RsvpMgmtSessionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtSessionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtSessionInfo proto.InternalMessageInfo

func (m *RsvpMgmtSessionInfo) GetRsvpSession() *RsvpSessionUnion {
	if m != nil {
		return m.RsvpSession
	}
	return nil
}

func init() {
	proto.RegisterType((*RsvpMgmtFrrStatesDetail_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frrs.frr.rsvp_mgmt_frr_states_detail_KEYS")
	proto.RegisterType((*RsvpMgmtFrrStatesDetail)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frrs.frr.rsvp_mgmt_frr_states_detail")
	proto.RegisterType((*RsvpMgmtS2LSubLspIpv4)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frrs.frr.rsvp_mgmt_s2l_sub_lsp_ipv4")
	proto.RegisterType((*RsvpMgmtSessionUdpIpv4)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frrs.frr.rsvp_mgmt_session_udp_ipv4")
	proto.RegisterType((*RsvpMgmtSessionLspTunnelIpv4)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frrs.frr.rsvp_mgmt_session_lsp_tunnel_ipv4")
	proto.RegisterType((*RsvpMgmtSessionUniIpv4)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frrs.frr.rsvp_mgmt_session_uni_ipv4")
	proto.RegisterType((*RsvpMgmtSessionP2MpLspTunnelIpv4)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frrs.frr.rsvp_mgmt_session_p2mp_lsp_tunnel_ipv4")
	proto.RegisterType((*RsvpSessionUnion)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frrs.frr.rsvp_session_union")
	proto.RegisterType((*RsvpMgmtSessionInfo)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frrs.frr.rsvp_mgmt_session_info")
}

func init() {
	proto.RegisterFile("rsvp_mgmt_frr_states_detail.proto", fileDescriptor_rsvp_mgmt_frr_states_detail_8f4900b5d76805cc)
}

var fileDescriptor_rsvp_mgmt_frr_states_detail_8f4900b5d76805cc = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4f, 0x6f, 0xd3, 0x4a,
	0x10, 0x97, 0x9b, 0xfe, 0x49, 0x26, 0x69, 0x9b, 0xb7, 0x4f, 0xaf, 0xf2, 0x6b, 0x0f, 0xa4, 0x91,
	0x80, 0x54, 0x82, 0x20, 0xb9, 0x45, 0x1c, 0x38, 0x20, 0x04, 0x08, 0x15, 0x52, 0x5a, 0x39, 0xad,
	0x50, 0x2f, 0xac, 0x92, 0x78, 0x53, 0x56, 0x8a, 0xbd, 0xab, 0xdd, 0x75, 0xd4, 0x9e, 0x39, 0x73,
	0x83, 0x33, 0xdf, 0x80, 0x13, 0xdf, 0x8e, 0x0b, 0xda, 0xb1, 0x9d, 0xb8, 0x8d, 0x5b, 0xd1, 0x16,
	0x2e, 0x6e, 0x33, 0x3b, 0xfb, 0xfb, 0x33, 0x3b, 0x33, 0xb0, 0xa9, 0xf4, 0x58, 0xd2, 0xf0, 0x24,
	0x34, 0x74, 0xa8, 0x14, 0xd5, 0xa6, 0x67, 0x98, 0xa6, 0x01, 0x33, 0x3d, 0x3e, 0x6a, 0x4b, 0x25,
	0x8c, 0x20, 0xf7, 0x07, 0x5c, 0x0f, 0x04, 0xe5, 0x42, 0xd3, 0x53, 0x45, 0xb9, 0xa4, 0x78, 0x45,
	0x48, 0xa6, 0xda, 0xf6, 0xbf, 0xf6, 0x50, 0x29, 0x6d, 0x3f, 0xcd, 0x1f, 0x25, 0x68, 0x5c, 0x01,
	0x47, 0xdf, 0xbe, 0x3a, 0xee, 0x92, 0x47, 0xf0, 0x6f, 0xc0, 0xb4, 0xe1, 0x51, 0xcf, 0x70, 0x11,
	0xd1, 0x5e, 0x10, 0x28, 0xa6, 0xb5, 0xeb, 0x34, 0x9c, 0x56, 0xc5, 0x27, 0xb9, 0xa3, 0xe7, 0xc9,
	0x09, 0xd9, 0x82, 0x7a, 0xfe, 0x82, 0x14, 0xca, 0xb8, 0x73, 0x0d, 0xa7, 0xb5, 0xec, 0xaf, 0xe6,
	0xe2, 0x07, 0x42, 0x19, 0xb2, 0x0e, 0x65, 0x94, 0x3c, 0x10, 0x23, 0xb7, 0x84, 0x29, 0x93, 0xdf,
	0xe4, 0x01, 0x10, 0x76, 0x6a, 0x58, 0x14, 0xb0, 0x80, 0x9a, 0x38, 0x8a, 0xd8, 0x88, 0xf2, 0xc0,
	0x9d, 0x47, 0xda, 0x7a, 0x76, 0x72, 0x88, 0x07, 0xbb, 0x01, 0xd9, 0x84, 0x9a, 0x66, 0x5a, 0x5b,
	0x42, 0x73, 0x26, 0x99, 0xbb, 0x80, 0x79, 0xd5, 0x34, 0x76, 0x78, 0x26, 0x19, 0x71, 0xa1, 0x2c,
	0x3d, 0x1a, 0x4a, 0x0b, 0xb3, 0x88, 0x64, 0x8b, 0xd2, 0xdb, 0x93, 0xbb, 0x01, 0xb9, 0x0b, 0x2b,
	0x5a, 0xc4, 0x6a, 0xc0, 0x26, 0xee, 0x96, 0xf0, 0xfa, 0x72, 0x12, 0xcd, 0x8c, 0xdd, 0x81, 0x6a,
	0x9a, 0x86, 0x9e, 0xca, 0x88, 0x01, 0x49, 0x08, 0xed, 0xb4, 0xa0, 0xae, 0xe3, 0x3e, 0x3d, 0x51,
	0x22, 0x96, 0x54, 0x28, 0x7e, 0xc2, 0x23, 0xb7, 0x82, 0x48, 0x2b, 0x3a, 0xee, 0xbf, 0xb6, 0xe1,
	0x7d, 0x8c, 0x92, 0x06, 0xd4, 0xa6, 0x99, 0x3c, 0x70, 0x21, 0xc5, 0x4a, 0xb3, 0x76, 0x03, 0xf2,
	0x3f, 0x94, 0xc7, 0x6a, 0x48, 0xa3, 0x5e, 0xc8, 0xdc, 0x2a, 0x62, 0x2c, 0x8d, 0xd5, 0xf0, 0x5d,
	0x2f, 0x64, 0xcd, 0xef, 0x73, 0xb0, 0x71, 0xc5, 0xb3, 0x91, 0x63, 0x58, 0x4a, 0x7d, 0xbb, 0x5e,
	0xc3, 0x69, 0x55, 0xbd, 0x67, 0xed, 0xdf, 0xec, 0x88, 0xf6, 0x14, 0x36, 0xab, 0x26, 0x8f, 0x86,
	0xc2, 0xcf, 0xf0, 0x48, 0x00, 0x35, 0xed, 0xd1, 0x11, 0xb5, 0xe2, 0x47, 0x5a, 0xba, 0xdb, 0x88,
	0xff, 0xe2, 0x26, 0xf8, 0xde, 0x04, 0x85, 0x72, 0x39, 0xde, 0xf1, 0x2b, 0xda, 0xeb, 0x74, 0xe3,
	0x7e, 0x47, 0x4b, 0x5b, 0x68, 0xd9, 0x33, 0x1f, 0xd1, 0x56, 0xac, 0xdd, 0x1d, 0xb4, 0x0f, 0x36,
	0xd4, 0xc5, 0x08, 0x79, 0x08, 0x44, 0x31, 0xcd, 0xd4, 0x38, 0x69, 0xb1, 0x34, 0xef, 0x31, 0xe6,
	0xfd, 0x93, 0x3b, 0x49, 0xd2, 0x9b, 0x47, 0xb0, 0x7e, 0x39, 0x31, 0x79, 0x02, 0x2e, 0x7a, 0xba,
	0xbc, 0xcb, 0xff, 0xd3, 0x5e, 0xe7, 0xe5, 0x4c, 0xa3, 0x37, 0xbf, 0x38, 0xe7, 0x70, 0xd3, 0x82,
	0xc5, 0x41, 0x8a, 0x7b, 0xed, 0xc1, 0xc9, 0x4f, 0xc3, 0xdc, 0x85, 0x69, 0x28, 0x1a, 0xaa, 0x52,
	0xe1, 0x50, 0x35, 0xbf, 0x39, 0xf9, 0x25, 0x91, 0xc9, 0xb2, 0x76, 0xb3, 0x31, 0xba, 0x91, 0xba,
	0x0d, 0xa8, 0x4c, 0xc7, 0x30, 0x95, 0x67, 0xb2, 0xf1, 0x2b, 0x1e, 0xd6, 0x52, 0xf1, 0xb0, 0x36,
	0xbf, 0x16, 0x17, 0x2e, 0xe2, 0x7f, 0x43, 0xda, 0x16, 0x4c, 0x04, 0x4c, 0xa0, 0x12, 0x61, 0xab,
	0x59, 0x3c, 0x7b, 0xd0, 0xcf, 0x0e, 0xdc, 0x9b, 0xd5, 0x25, 0xbd, 0x50, 0xce, 0x94, 0x2f, 0xbf,
	0x4c, 0x9c, 0x73, 0xcb, 0xe4, 0x0f, 0xd6, 0xe9, 0x67, 0x09, 0x08, 0xea, 0xc9, 0x95, 0x48, 0x44,
	0x33, 0xbb, 0xce, 0x99, 0xdd, 0x75, 0xef, 0x61, 0xde, 0xca, 0x44, 0xfe, 0x1b, 0xce, 0xe7, 0x85,
	0x76, 0xf6, 0x11, 0x90, 0x18, 0xa8, 0xdb, 0xbf, 0x58, 0x8f, 0x6c, 0xc9, 0x94, 0x90, 0xe4, 0xcd,
	0x2d, 0x48, 0x2e, 0x54, 0xd7, 0x5f, 0xb1, 0xdf, 0x8e, 0x96, 0xdd, 0x74, 0xed, 0x84, 0x29, 0xab,
	0x6d, 0x91, 0x8c, 0x75, 0xfe, 0xf6, 0xd6, 0xd2, 0x86, 0x4b, 0xe8, 0x8e, 0x22, 0x9e, 0xd1, 0x7d,
	0x72, 0x60, 0x0d, 0xf9, 0x92, 0x27, 0xce, 0x7b, 0x5d, 0x40, 0xd6, 0xfd, 0x5b, 0xb0, 0x16, 0xb5,
	0x93, 0x4f, 0xec, 0xf7, 0xc0, 0xdb, 0x93, 0x53, 0xd3, 0xcd, 0x53, 0x58, 0x2b, 0x5e, 0xc7, 0xe4,
	0x03, 0xd4, 0xf2, 0x6d, 0x81, 0x0d, 0x50, 0xf5, 0x9e, 0x5e, 0x4f, 0xd4, 0xb9, 0x9e, 0xf2, 0xab,
	0x36, 0x96, 0x32, 0xf7, 0x17, 0x71, 0xed, 0x6c, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x79,
	0xe6, 0xc6, 0x6c, 0x08, 0x00, 0x00,
}
