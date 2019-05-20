// Code generated by protoc-gen-go. DO NOT EDIT.
// source: te_sh_igp.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_igp_areas_igp_area

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

// An IGP instance
type TeShIgp_KEYS struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	ProcessTag           string   `protobuf:"bytes,2,opt,name=process_tag,json=processTag" json:"process_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeShIgp_KEYS) Reset()         { *m = TeShIgp_KEYS{} }
func (m *TeShIgp_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeShIgp_KEYS) ProtoMessage()    {}
func (*TeShIgp_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_sh_igp_d10e4fc10764a87c, []int{0}
}
func (m *TeShIgp_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeShIgp_KEYS.Unmarshal(m, b)
}
func (m *TeShIgp_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeShIgp_KEYS.Marshal(b, m, deterministic)
}
func (dst *TeShIgp_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeShIgp_KEYS.Merge(dst, src)
}
func (m *TeShIgp_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeShIgp_KEYS.Size(m)
}
func (m *TeShIgp_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeShIgp_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeShIgp_KEYS proto.InternalMessageInfo

func (m *TeShIgp_KEYS) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *TeShIgp_KEYS) GetProcessTag() string {
	if m != nil {
		return m.ProcessTag
	}
	return ""
}

type TeShIgp struct {
	// IGP type
	IgpType string `protobuf:"bytes,50,opt,name=igp_type,json=igpType" json:"igp_type,omitempty"`
	//  The IGP instance name
	InstanceName string `protobuf:"bytes,51,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	//  The IGP instance's system/router-id; interpret appropriately depending on IGP type
	IgpSystemId string `protobuf:"bytes,52,opt,name=igp_system_id,json=igpSystemId" json:"igp_system_id,omitempty"`
	// Configured TE router-id
	ConfiguredTeRouterId string `protobuf:"bytes,53,opt,name=configured_te_router_id,json=configuredTeRouterId" json:"configured_te_router_id,omitempty"`
	// Global router-id
	GlobalRouterId string `protobuf:"bytes,54,opt,name=global_router_id,json=globalRouterId" json:"global_router_id,omitempty"`
	// Secondary router-ids
	SecondaryRouterId [][]byte `protobuf:"bytes,55,rep,name=secondary_router_id,json=secondaryRouterId,proto3" json:"secondary_router_id,omitempty"`
	// Global router-id for GMPLS (optical TE)
	GloballRouterIdOptical string `protobuf:"bytes,56,opt,name=globall_router_id_optical,json=globallRouterIdOptical" json:"globall_router_id_optical,omitempty"`
	// TE router-id in use
	InUseTeRouterId string `protobuf:"bytes,57,opt,name=in_use_te_router_id,json=inUseTeRouterId" json:"in_use_te_router_id,omitempty"`
	// Flag to indicate whether the IGP connection is open or not
	IsConnectionUp bool `protobuf:"varint,58,opt,name=is_connection_up,json=isConnectionUp" json:"is_connection_up,omitempty"`
	// Number of times IGP connection has gone up
	ConnectionUpCount uint32 `protobuf:"varint,59,opt,name=connection_up_count,json=connectionUpCount" json:"connection_up_count,omitempty"`
	// Number of times IGP connection has gone down
	ConnectionDownCount uint32 `protobuf:"varint,60,opt,name=connection_down_count,json=connectionDownCount" json:"connection_down_count,omitempty"`
	// IGP area information
	AreaList             []*TeShIgpArea `protobuf:"bytes,61,rep,name=area_list,json=areaList" json:"area_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TeShIgp) Reset()         { *m = TeShIgp{} }
func (m *TeShIgp) String() string { return proto.CompactTextString(m) }
func (*TeShIgp) ProtoMessage()    {}
func (*TeShIgp) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_sh_igp_d10e4fc10764a87c, []int{1}
}
func (m *TeShIgp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeShIgp.Unmarshal(m, b)
}
func (m *TeShIgp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeShIgp.Marshal(b, m, deterministic)
}
func (dst *TeShIgp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeShIgp.Merge(dst, src)
}
func (m *TeShIgp) XXX_Size() int {
	return xxx_messageInfo_TeShIgp.Size(m)
}
func (m *TeShIgp) XXX_DiscardUnknown() {
	xxx_messageInfo_TeShIgp.DiscardUnknown(m)
}

var xxx_messageInfo_TeShIgp proto.InternalMessageInfo

func (m *TeShIgp) GetIgpType() string {
	if m != nil {
		return m.IgpType
	}
	return ""
}

func (m *TeShIgp) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *TeShIgp) GetIgpSystemId() string {
	if m != nil {
		return m.IgpSystemId
	}
	return ""
}

func (m *TeShIgp) GetConfiguredTeRouterId() string {
	if m != nil {
		return m.ConfiguredTeRouterId
	}
	return ""
}

func (m *TeShIgp) GetGlobalRouterId() string {
	if m != nil {
		return m.GlobalRouterId
	}
	return ""
}

func (m *TeShIgp) GetSecondaryRouterId() [][]byte {
	if m != nil {
		return m.SecondaryRouterId
	}
	return nil
}

func (m *TeShIgp) GetGloballRouterIdOptical() string {
	if m != nil {
		return m.GloballRouterIdOptical
	}
	return ""
}

func (m *TeShIgp) GetInUseTeRouterId() string {
	if m != nil {
		return m.InUseTeRouterId
	}
	return ""
}

func (m *TeShIgp) GetIsConnectionUp() bool {
	if m != nil {
		return m.IsConnectionUp
	}
	return false
}

func (m *TeShIgp) GetConnectionUpCount() uint32 {
	if m != nil {
		return m.ConnectionUpCount
	}
	return 0
}

func (m *TeShIgp) GetConnectionDownCount() uint32 {
	if m != nil {
		return m.ConnectionDownCount
	}
	return 0
}

func (m *TeShIgp) GetAreaList() []*TeShIgpArea {
	if m != nil {
		return m.AreaList
	}
	return nil
}

// Statistics for an IGP-area
type TeShIgpAreaStats struct {
	// Number of adjacency request messages sent
	AreaAdjacencyRequestMessages uint32 `protobuf:"varint,1,opt,name=area_adjacency_request_messages,json=areaAdjacencyRequestMessages" json:"area_adjacency_request_messages,omitempty"`
	// Number of adjacency announcement messages received
	AreaAdjacencyAnnounceMessages uint32 `protobuf:"varint,2,opt,name=area_adjacency_announce_messages,json=areaAdjacencyAnnounceMessages" json:"area_adjacency_announce_messages,omitempty"`
	// Number of local LSA floods sent
	AreaFloodMessages uint32 `protobuf:"varint,3,opt,name=area_flood_messages,json=areaFloodMessages" json:"area_flood_messages,omitempty"`
	// Number of LSA announcement messages received
	AreaLsaAnnounceMessages uint32 `protobuf:"varint,4,opt,name=area_lsa_announce_messages,json=areaLsaAnnounceMessages" json:"area_lsa_announce_messages,omitempty"`
	// Number of LSA fragment announcement messages received
	AreaLsaFragmentAnnounceMessages uint32 `protobuf:"varint,5,opt,name=area_lsa_fragment_announce_messages,json=areaLsaFragmentAnnounceMessages" json:"area_lsa_fragment_announce_messages,omitempty"`
	// Number of LSA delete messages received
	AreaLsaDeleteMessages uint32 `protobuf:"varint,6,opt,name=area_lsa_delete_messages,json=areaLsaDeleteMessages" json:"area_lsa_delete_messages,omitempty"`
	// Number of LSA fragment delete messages received
	AreaLsaFragmentDeleteMessages uint32 `protobuf:"varint,7,opt,name=area_lsa_fragment_delete_messages,json=areaLsaFragmentDeleteMessages" json:"area_lsa_fragment_delete_messages,omitempty"`
	// Number of tunnel announcement messages sent
	AreaTunnelAnnounceMessages uint32   `protobuf:"varint,8,opt,name=area_tunnel_announce_messages,json=areaTunnelAnnounceMessages" json:"area_tunnel_announce_messages,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *TeShIgpAreaStats) Reset()         { *m = TeShIgpAreaStats{} }
func (m *TeShIgpAreaStats) String() string { return proto.CompactTextString(m) }
func (*TeShIgpAreaStats) ProtoMessage()    {}
func (*TeShIgpAreaStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_sh_igp_d10e4fc10764a87c, []int{2}
}
func (m *TeShIgpAreaStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeShIgpAreaStats.Unmarshal(m, b)
}
func (m *TeShIgpAreaStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeShIgpAreaStats.Marshal(b, m, deterministic)
}
func (dst *TeShIgpAreaStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeShIgpAreaStats.Merge(dst, src)
}
func (m *TeShIgpAreaStats) XXX_Size() int {
	return xxx_messageInfo_TeShIgpAreaStats.Size(m)
}
func (m *TeShIgpAreaStats) XXX_DiscardUnknown() {
	xxx_messageInfo_TeShIgpAreaStats.DiscardUnknown(m)
}

var xxx_messageInfo_TeShIgpAreaStats proto.InternalMessageInfo

func (m *TeShIgpAreaStats) GetAreaAdjacencyRequestMessages() uint32 {
	if m != nil {
		return m.AreaAdjacencyRequestMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaAdjacencyAnnounceMessages() uint32 {
	if m != nil {
		return m.AreaAdjacencyAnnounceMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaFloodMessages() uint32 {
	if m != nil {
		return m.AreaFloodMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaLsaAnnounceMessages() uint32 {
	if m != nil {
		return m.AreaLsaAnnounceMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaLsaFragmentAnnounceMessages() uint32 {
	if m != nil {
		return m.AreaLsaFragmentAnnounceMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaLsaDeleteMessages() uint32 {
	if m != nil {
		return m.AreaLsaDeleteMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaLsaFragmentDeleteMessages() uint32 {
	if m != nil {
		return m.AreaLsaFragmentDeleteMessages
	}
	return 0
}

func (m *TeShIgpAreaStats) GetAreaTunnelAnnounceMessages() uint32 {
	if m != nil {
		return m.AreaTunnelAnnounceMessages
	}
	return 0
}

// Data relating to an active IGP area
type TeShIgpAreaActive struct {
	// Number of interfaces running over this area
	InterfacesCount uint32 `protobuf:"varint,1,opt,name=interfaces_count,json=interfacesCount" json:"interfaces_count,omitempty"`
	// Flag to indicate IDT for link adjacencies was received
	LinkIdtReceived bool `protobuf:"varint,2,opt,name=link_idt_received,json=linkIdtReceived" json:"link_idt_received,omitempty"`
	// Flag to indicate IDT topology was received
	TopologyIdtReceived bool `protobuf:"varint,3,opt,name=topology_idt_received,json=topologyIdtReceived" json:"topology_idt_received,omitempty"`
	// Flag to indicate if the area is SR strict
	SrStrict bool `protobuf:"varint,4,opt,name=sr_strict,json=srStrict" json:"sr_strict,omitempty"`
	// Number of p2p tunnel heads whose path was calculated over this area
	P2PHeadsCount uint32 `protobuf:"varint,5,opt,name=p2_p_heads_count,json=p2PHeadsCount" json:"p2_p_heads_count,omitempty"`
	// Number of p2p tunnel heads which have been autoroute-announced into this area
	P2PAutorouteAnnouncedCount uint32 `protobuf:"varint,6,opt,name=p2_p_autoroute_announced_count,json=p2PAutorouteAnnouncedCount" json:"p2_p_autoroute_announced_count,omitempty"`
	// Number of p2p tunnel heads which have been announced as forwarding adjacencies in this area
	P2PForwardingAdjacencyCount uint32 `protobuf:"varint,7,opt,name=p2_p_forwarding_adjacency_count,json=p2PForwardingAdjacencyCount" json:"p2_p_forwarding_adjacency_count,omitempty"`
	// Number of P2MP destinations whose path was calculated over this area
	P2MpDestinationCount uint32 `protobuf:"varint,8,opt,name=p2_mp_destination_count,json=p2MpDestinationCount" json:"p2_mp_destination_count,omitempty"`
	// Number of tunnels with a loose hop ERO expanded in this area
	TunnelLooseHops uint32 `protobuf:"varint,9,opt,name=tunnel_loose_hops,json=tunnelLooseHops" json:"tunnel_loose_hops,omitempty"`
	// Area communication statistics
	AreaStatistics       *TeShIgpAreaStats `protobuf:"bytes,10,opt,name=area_statistics,json=areaStatistics" json:"area_statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TeShIgpAreaActive) Reset()         { *m = TeShIgpAreaActive{} }
func (m *TeShIgpAreaActive) String() string { return proto.CompactTextString(m) }
func (*TeShIgpAreaActive) ProtoMessage()    {}
func (*TeShIgpAreaActive) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_sh_igp_d10e4fc10764a87c, []int{3}
}
func (m *TeShIgpAreaActive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeShIgpAreaActive.Unmarshal(m, b)
}
func (m *TeShIgpAreaActive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeShIgpAreaActive.Marshal(b, m, deterministic)
}
func (dst *TeShIgpAreaActive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeShIgpAreaActive.Merge(dst, src)
}
func (m *TeShIgpAreaActive) XXX_Size() int {
	return xxx_messageInfo_TeShIgpAreaActive.Size(m)
}
func (m *TeShIgpAreaActive) XXX_DiscardUnknown() {
	xxx_messageInfo_TeShIgpAreaActive.DiscardUnknown(m)
}

var xxx_messageInfo_TeShIgpAreaActive proto.InternalMessageInfo

func (m *TeShIgpAreaActive) GetInterfacesCount() uint32 {
	if m != nil {
		return m.InterfacesCount
	}
	return 0
}

func (m *TeShIgpAreaActive) GetLinkIdtReceived() bool {
	if m != nil {
		return m.LinkIdtReceived
	}
	return false
}

func (m *TeShIgpAreaActive) GetTopologyIdtReceived() bool {
	if m != nil {
		return m.TopologyIdtReceived
	}
	return false
}

func (m *TeShIgpAreaActive) GetSrStrict() bool {
	if m != nil {
		return m.SrStrict
	}
	return false
}

func (m *TeShIgpAreaActive) GetP2PHeadsCount() uint32 {
	if m != nil {
		return m.P2PHeadsCount
	}
	return 0
}

func (m *TeShIgpAreaActive) GetP2PAutorouteAnnouncedCount() uint32 {
	if m != nil {
		return m.P2PAutorouteAnnouncedCount
	}
	return 0
}

func (m *TeShIgpAreaActive) GetP2PForwardingAdjacencyCount() uint32 {
	if m != nil {
		return m.P2PForwardingAdjacencyCount
	}
	return 0
}

func (m *TeShIgpAreaActive) GetP2MpDestinationCount() uint32 {
	if m != nil {
		return m.P2MpDestinationCount
	}
	return 0
}

func (m *TeShIgpAreaActive) GetTunnelLooseHops() uint32 {
	if m != nil {
		return m.TunnelLooseHops
	}
	return 0
}

func (m *TeShIgpAreaActive) GetAreaStatistics() *TeShIgpAreaStats {
	if m != nil {
		return m.AreaStatistics
	}
	return nil
}

// An IGP area
type TeShIgpArea struct {
	// Internal area index
	AreaIndex uint32 `protobuf:"varint,1,opt,name=area_index,json=areaIndex" json:"area_index,omitempty"`
	// Area number
	AreaNumber uint32 `protobuf:"varint,2,opt,name=area_number,json=areaNumber" json:"area_number,omitempty"`
	// IGP Area Format
	AreaFormat string `protobuf:"bytes,3,opt,name=area_format,json=areaFormat" json:"area_format,omitempty"`
	// Indicates whether or not the area is correctly configured under the IGP submode to run TE
	IsConfigReady bool `protobuf:"varint,4,opt,name=is_config_ready,json=isConfigReady" json:"is_config_ready,omitempty"`
	// Data which applies only to an active area
	ActiveData           *TeShIgpAreaActive `protobuf:"bytes,5,opt,name=active_data,json=activeData" json:"active_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TeShIgpArea) Reset()         { *m = TeShIgpArea{} }
func (m *TeShIgpArea) String() string { return proto.CompactTextString(m) }
func (*TeShIgpArea) ProtoMessage()    {}
func (*TeShIgpArea) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_sh_igp_d10e4fc10764a87c, []int{4}
}
func (m *TeShIgpArea) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeShIgpArea.Unmarshal(m, b)
}
func (m *TeShIgpArea) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeShIgpArea.Marshal(b, m, deterministic)
}
func (dst *TeShIgpArea) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeShIgpArea.Merge(dst, src)
}
func (m *TeShIgpArea) XXX_Size() int {
	return xxx_messageInfo_TeShIgpArea.Size(m)
}
func (m *TeShIgpArea) XXX_DiscardUnknown() {
	xxx_messageInfo_TeShIgpArea.DiscardUnknown(m)
}

var xxx_messageInfo_TeShIgpArea proto.InternalMessageInfo

func (m *TeShIgpArea) GetAreaIndex() uint32 {
	if m != nil {
		return m.AreaIndex
	}
	return 0
}

func (m *TeShIgpArea) GetAreaNumber() uint32 {
	if m != nil {
		return m.AreaNumber
	}
	return 0
}

func (m *TeShIgpArea) GetAreaFormat() string {
	if m != nil {
		return m.AreaFormat
	}
	return ""
}

func (m *TeShIgpArea) GetIsConfigReady() bool {
	if m != nil {
		return m.IsConfigReady
	}
	return false
}

func (m *TeShIgpArea) GetActiveData() *TeShIgpAreaActive {
	if m != nil {
		return m.ActiveData
	}
	return nil
}

func init() {
	proto.RegisterType((*TeShIgp_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.igp_areas.igp_area.te_sh_igp_KEYS")
	proto.RegisterType((*TeShIgp)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.igp_areas.igp_area.te_sh_igp")
	proto.RegisterType((*TeShIgpAreaStats)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.igp_areas.igp_area.te_sh_igp_area_stats")
	proto.RegisterType((*TeShIgpAreaActive)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.igp_areas.igp_area.te_sh_igp_area_active")
	proto.RegisterType((*TeShIgpArea)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.igp_areas.igp_area.te_sh_igp_area")
}

func init() { proto.RegisterFile("te_sh_igp.proto", fileDescriptor_te_sh_igp_d10e4fc10764a87c) }

var fileDescriptor_te_sh_igp_d10e4fc10764a87c = []byte{
	// 965 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x52, 0x1b, 0xb7,
	0x17, 0x1f, 0x3e, 0xfe, 0x89, 0x7d, 0x1c, 0x63, 0x58, 0xe0, 0xcf, 0x86, 0x34, 0xc5, 0x75, 0x66,
	0x5a, 0xb7, 0xd3, 0xf1, 0xc5, 0x26, 0x34, 0x4d, 0xd3, 0x5e, 0xd0, 0x38, 0x14, 0x1a, 0x48, 0x99,
	0x85, 0x5c, 0xf4, 0x4a, 0x23, 0x76, 0xe5, 0x45, 0xed, 0x5a, 0x52, 0x25, 0x39, 0xc4, 0x4f, 0xd0,
	0x47, 0xe9, 0xbb, 0xf4, 0x11, 0xfa, 0x34, 0x1d, 0x1d, 0xad, 0xe5, 0x35, 0x70, 0xd5, 0xc9, 0x0d,
	0x58, 0xe7, 0xf7, 0x21, 0xe9, 0x9c, 0xa3, 0x63, 0x43, 0xc7, 0x32, 0x62, 0xae, 0x08, 0x2f, 0xd4,
	0x40, 0x69, 0x69, 0x65, 0xf4, 0x2c, 0xe3, 0x26, 0x93, 0x84, 0x4b, 0x43, 0x3e, 0x68, 0x32, 0x56,
	0xa5, 0x21, 0x96, 0x11, 0xa9, 0x98, 0x1e, 0x54, 0x8b, 0x01, 0x2f, 0x14, 0xa1, 0x9a, 0x51, 0x13,
	0x3e, 0xf5, 0x4e, 0x61, 0x2d, 0x18, 0x91, 0x37, 0xaf, 0x7f, 0x3d, 0x8f, 0x76, 0xa1, 0x81, 0x86,
	0x99, 0x2c, 0xe3, 0xa5, 0xee, 0x52, 0xbf, 0x99, 0x86, 0x75, 0xb4, 0x07, 0x2d, 0xa5, 0x65, 0xc6,
	0x8c, 0x21, 0x96, 0x16, 0xf1, 0x32, 0xc2, 0x50, 0x85, 0x2e, 0x68, 0xd1, 0xfb, 0x67, 0x15, 0x9a,
	0xc1, 0x2f, 0x7a, 0x08, 0x0d, 0x67, 0x6b, 0xa7, 0x8a, 0xc5, 0x09, 0x72, 0xef, 0xf3, 0x42, 0x5d,
	0x4c, 0x15, 0x8b, 0x9e, 0x40, 0x9b, 0x0b, 0x63, 0xa9, 0xc8, 0x18, 0x11, 0x74, 0xcc, 0xe2, 0xa7,
	0x88, 0x3f, 0x98, 0x05, 0xdf, 0xd2, 0x31, 0x8b, 0x7a, 0xd0, 0x76, 0x7a, 0x33, 0x35, 0x96, 0x8d,
	0x09, 0xcf, 0xe3, 0x67, 0x48, 0x6a, 0xf1, 0x42, 0x9d, 0x63, 0xec, 0x38, 0x8f, 0xf6, 0x61, 0x27,
	0x93, 0x62, 0xc4, 0x8b, 0x89, 0x66, 0xb9, 0xbb, 0xb1, 0x96, 0x13, 0xcb, 0xb4, 0x63, 0xef, 0x23,
	0x7b, 0x6b, 0x0e, 0x5f, 0xb0, 0x14, 0xc1, 0xe3, 0x3c, 0xea, 0xc3, 0x7a, 0x51, 0xca, 0x4b, 0x5a,
	0xd6, 0xf8, 0xdf, 0x20, 0x7f, 0xcd, 0xc7, 0x03, 0x73, 0x00, 0x9b, 0x86, 0x65, 0x52, 0xe4, 0x54,
	0x4f, 0x6b, 0xe4, 0xe7, 0xdd, 0x95, 0xfe, 0x83, 0x74, 0x23, 0x40, 0x81, 0xff, 0x02, 0x1e, 0x7a,
	0x87, 0x9a, 0x35, 0x91, 0xca, 0xf2, 0x8c, 0x96, 0xf1, 0xb7, 0xb8, 0xc5, 0xff, 0x2b, 0xc2, 0x4c,
	0xf3, 0x8b, 0x47, 0xa3, 0xaf, 0x61, 0x93, 0x0b, 0x32, 0x31, 0x6c, 0xf1, 0x1e, 0x2f, 0x50, 0xd4,
	0xe1, 0xe2, 0x9d, 0x61, 0x8b, 0x57, 0xe0, 0x86, 0x64, 0x52, 0x08, 0x96, 0x59, 0x2e, 0x05, 0x99,
	0xa8, 0xf8, 0xbb, 0xee, 0x52, 0xbf, 0x91, 0xae, 0x71, 0xf3, 0x2a, 0x84, 0xdf, 0x29, 0x77, 0x85,
	0x05, 0x1a, 0xc9, 0xe4, 0x44, 0xd8, 0xf8, 0x65, 0x77, 0xa9, 0xdf, 0x4e, 0x37, 0xb2, 0x1a, 0xf5,
	0x95, 0x03, 0xa2, 0x04, 0xb6, 0x6b, 0xfc, 0x5c, 0x5e, 0x8b, 0x4a, 0xf1, 0x3d, 0x2a, 0x6a, 0x66,
	0x43, 0x79, 0x2d, 0xbc, 0x86, 0x42, 0xd3, 0x35, 0x14, 0x29, 0xb9, 0xb1, 0xf1, 0x0f, 0xdd, 0x95,
	0x7e, 0x2b, 0x19, 0x0e, 0xfe, 0x4b, 0x4b, 0x0e, 0xe6, 0xfd, 0xe8, 0x96, 0x69, 0xc3, 0xfd, 0x3d,
	0xe1, 0xc6, 0xf6, 0xfe, 0x5a, 0x85, 0xad, 0x45, 0x90, 0x18, 0x4b, 0xad, 0x89, 0x5e, 0xc3, 0x1e,
	0xae, 0x68, 0xfe, 0x1b, 0xcd, 0x98, 0xc8, 0xa6, 0x44, 0xb3, 0x3f, 0x26, 0xcc, 0x58, 0x32, 0x66,
	0xc6, 0xd0, 0x82, 0x19, 0xec, 0xe4, 0x76, 0xfa, 0x89, 0xa3, 0x1d, 0xcc, 0x58, 0xa9, 0x27, 0x9d,
	0x56, 0x9c, 0xe8, 0x27, 0xe8, 0xde, 0xb0, 0xa1, 0x42, 0xc8, 0x89, 0x6b, 0xd1, 0xe0, 0xb3, 0x8c,
	0x3e, 0x8f, 0x17, 0x7c, 0x0e, 0x2a, 0x56, 0x30, 0x1a, 0xc0, 0x26, 0x1a, 0x8d, 0x4a, 0x29, 0xf3,
	0xb9, 0x76, 0xc5, 0xe7, 0xdb, 0x41, 0x87, 0x0e, 0x09, 0xfc, 0x97, 0xb0, 0xeb, 0x73, 0x67, 0xe8,
	0x1d, 0x5b, 0xae, 0xa2, 0x6c, 0x07, 0xd3, 0x60, 0xe8, 0xad, 0xcd, 0x4e, 0xe0, 0x49, 0x10, 0x8f,
	0x34, 0x2d, 0xc6, 0x4c, 0xd8, 0x3b, 0x5c, 0xfe, 0x87, 0x2e, 0x7b, 0x95, 0xcb, 0x61, 0x45, 0xbc,
	0xe5, 0xf6, 0x1c, 0xe2, 0xe0, 0x96, 0xb3, 0x92, 0xd9, 0x9a, 0xc5, 0x3d, 0xb4, 0xd8, 0xae, 0x2c,
	0x86, 0x88, 0x06, 0xe1, 0x11, 0x7c, 0x76, 0xfb, 0x18, 0x37, 0x1d, 0xee, 0xcf, 0xb3, 0x57, 0x3b,
	0xc4, 0x0d, 0xa7, 0x03, 0x40, 0x02, 0xb1, 0x13, 0x21, 0x58, 0x79, 0xc7, 0x55, 0x1a, 0xe8, 0x82,
	0x29, 0xbb, 0x40, 0xce, 0xcd, 0x5b, 0xf4, 0xfe, 0x5e, 0x85, 0xed, 0x1b, 0x9d, 0x42, 0x33, 0xcb,
	0xdf, 0xb3, 0xe8, 0x4b, 0x58, 0xe7, 0xc2, 0x32, 0x3d, 0xa2, 0x19, 0x33, 0x55, 0x57, 0xfb, 0xde,
	0xe8, 0xcc, 0xe3, 0xbe, 0xa3, 0xbf, 0x82, 0x8d, 0x92, 0x8b, 0xdf, 0x09, 0xcf, 0x2d, 0xd1, 0x2c,
	0x63, 0xfc, 0x3d, 0xcb, 0xb1, 0xfe, 0x8d, 0xb4, 0xe3, 0x80, 0xe3, 0xdc, 0xa6, 0x55, 0xd8, 0xbd,
	0x18, 0x2b, 0x95, 0x2c, 0x65, 0x31, 0x5d, 0xe4, 0xaf, 0x20, 0x7f, 0x73, 0x06, 0xd6, 0x35, 0x8f,
	0xa0, 0x69, 0x34, 0x31, 0x56, 0xf3, 0xcc, 0x62, 0x91, 0x1b, 0x69, 0xc3, 0xe8, 0x73, 0x5c, 0x47,
	0x5f, 0xc0, 0xba, 0x4a, 0x88, 0x22, 0x57, 0x8c, 0xe6, 0xb3, 0x73, 0xfa, 0x12, 0xb6, 0x55, 0x72,
	0x76, 0xe4, 0xa2, 0xfe, 0x94, 0x3f, 0xc2, 0xa7, 0x48, 0xa4, 0x13, 0x2b, 0x71, 0x64, 0x84, 0x84,
	0xe5, 0x95, 0xcc, 0x97, 0x6d, 0x57, 0x25, 0x67, 0x07, 0x33, 0xce, 0x2c, 0x61, 0xb9, 0xf7, 0x18,
	0xc2, 0x1e, 0x7a, 0x8c, 0xa4, 0xbe, 0xa6, 0x3a, 0xe7, 0xa2, 0xa8, 0xbd, 0x01, 0x6f, 0xe2, 0x2b,
	0xf7, 0x48, 0x25, 0x67, 0x87, 0x81, 0x14, 0x1e, 0x80, 0x77, 0xd9, 0x87, 0x1d, 0x95, 0x90, 0xb1,
	0x22, 0x39, 0x33, 0x96, 0x0b, 0x8a, 0xc3, 0xc3, 0xab, 0x7d, 0xc5, 0xb6, 0x54, 0x72, 0xaa, 0x86,
	0x73, 0x30, 0xa4, 0xb9, 0xaa, 0x74, 0x29, 0xa5, 0x61, 0xe4, 0x4a, 0x2a, 0x13, 0x37, 0x7d, 0x49,
	0x3c, 0x70, 0xe2, 0xe2, 0x47, 0x52, 0x99, 0xc8, 0x40, 0x27, 0x3c, 0x7b, 0x6e, 0x2c, 0xcf, 0x4c,
	0x0c, 0xdd, 0xa5, 0x7e, 0x2b, 0xf9, 0xf9, 0x63, 0x8c, 0x1a, 0x3f, 0x4d, 0xd2, 0x35, 0xf7, 0xf9,
	0x3c, 0xec, 0xd0, 0xfb, 0x73, 0xb9, 0xfe, 0x1d, 0xe9, 0xc0, 0xe8, 0x31, 0x00, 0x0a, 0xb8, 0xc8,
	0xd9, 0x87, 0xaa, 0x7f, 0x70, 0xfc, 0x1d, 0xbb, 0x80, 0xfb, 0x9a, 0x44, 0x58, 0x4c, 0xc6, 0x97,
	0x4c, 0x57, 0x33, 0x03, 0x15, 0x6f, 0x31, 0x12, 0x08, 0x23, 0xa9, 0xc7, 0xd4, 0x62, 0x93, 0x34,
	0x3d, 0xe1, 0x10, 0x23, 0xd1, 0xe7, 0xd0, 0xf1, 0xb3, 0x7d, 0xc4, 0x0b, 0xa2, 0x19, 0xcd, 0xa7,
	0x55, 0x87, 0xb4, 0x71, 0xb4, 0x8f, 0x78, 0x91, 0xba, 0x60, 0x54, 0x42, 0xcb, 0x37, 0x36, 0xc9,
	0xa9, 0xa5, 0xd8, 0x21, 0xad, 0xe4, 0xcd, 0x47, 0x49, 0x86, 0xf7, 0x4d, 0xc1, 0xff, 0x1f, 0x52,
	0x4b, 0x2f, 0xef, 0xe1, 0x0f, 0x81, 0xa7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x59, 0x71,
	0xa0, 0x7c, 0x08, 0x00, 0x00,
}