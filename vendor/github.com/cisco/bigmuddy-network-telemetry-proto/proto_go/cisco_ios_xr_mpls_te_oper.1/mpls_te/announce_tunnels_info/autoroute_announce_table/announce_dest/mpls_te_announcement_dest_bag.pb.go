// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_announcement_dest_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_announce_tunnels_info_autoroute_announce_table_announce_dest

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

// The summary info of announcements for a destination and an IGP area
type MplsTeAnnouncementDestBag_KEYS struct {
	DestinationAddress   string   `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	Protocol             string   `protobuf:"bytes,2,opt,name=protocol" json:"protocol,omitempty"`
	Area                 uint32   `protobuf:"varint,3,opt,name=area" json:"area,omitempty"`
	IgpId                string   `protobuf:"bytes,4,opt,name=igp_id,json=igpId" json:"igp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeAnnouncementDestBag_KEYS) Reset()         { *m = MplsTeAnnouncementDestBag_KEYS{} }
func (m *MplsTeAnnouncementDestBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeAnnouncementDestBag_KEYS) ProtoMessage()    {}
func (*MplsTeAnnouncementDestBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_announcement_dest_bag_be31fd0d431ba514, []int{0}
}
func (m *MplsTeAnnouncementDestBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAnnouncementDestBag_KEYS.Unmarshal(m, b)
}
func (m *MplsTeAnnouncementDestBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAnnouncementDestBag_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTeAnnouncementDestBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAnnouncementDestBag_KEYS.Merge(dst, src)
}
func (m *MplsTeAnnouncementDestBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeAnnouncementDestBag_KEYS.Size(m)
}
func (m *MplsTeAnnouncementDestBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAnnouncementDestBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAnnouncementDestBag_KEYS proto.InternalMessageInfo

func (m *MplsTeAnnouncementDestBag_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTeAnnouncementDestBag_KEYS) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *MplsTeAnnouncementDestBag_KEYS) GetArea() uint32 {
	if m != nil {
		return m.Area
	}
	return 0
}

func (m *MplsTeAnnouncementDestBag_KEYS) GetIgpId() string {
	if m != nil {
		return m.IgpId
	}
	return ""
}

type MplsTeAnnouncementDestBag struct {
	// IGP type
	IgpType string `protobuf:"bytes,50,opt,name=igp_type,json=igpType" json:"igp_type,omitempty"`
	// IGP Instance name
	IgpInstance string `protobuf:"bytes,51,opt,name=igp_instance,json=igpInstance" json:"igp_instance,omitempty"`
	// IGP Area ID
	IgpArea uint32 `protobuf:"varint,52,opt,name=igp_area,json=igpArea" json:"igp_area,omitempty"`
	// Destination
	Destination string `protobuf:"bytes,53,opt,name=destination" json:"destination,omitempty"`
	// The number of announced tunnels
	AnnouncedTunnelsCount uint32 `protobuf:"varint,54,opt,name=announced_tunnels_count,json=announcedTunnelsCount" json:"announced_tunnels_count,omitempty"`
	// The number of tunnels that are announced as shortcuts
	AutoroutedTunnelsCount uint32 `protobuf:"varint,55,opt,name=autorouted_tunnels_count,json=autoroutedTunnelsCount" json:"autorouted_tunnels_count,omitempty"`
	// The number of tunnels that are announced as forwarding adjacencies
	ForwardingAdjacencyTunnelsCount uint32 `protobuf:"varint,56,opt,name=forwarding_adjacency_tunnels_count,json=forwardingAdjacencyTunnelsCount" json:"forwarding_adjacency_tunnels_count,omitempty"`
	// The format for the area: IPv4 address or a positive integer
	AreaFormat string `protobuf:"bytes,57,opt,name=area_format,json=areaFormat" json:"area_format,omitempty"`
	// Announced tunnels for this destination
	TotalTunnels         []*MplsTeAutorouteBag `protobuf:"bytes,58,rep,name=total_tunnels,json=totalTunnels" json:"total_tunnels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MplsTeAnnouncementDestBag) Reset()         { *m = MplsTeAnnouncementDestBag{} }
func (m *MplsTeAnnouncementDestBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeAnnouncementDestBag) ProtoMessage()    {}
func (*MplsTeAnnouncementDestBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_announcement_dest_bag_be31fd0d431ba514, []int{1}
}
func (m *MplsTeAnnouncementDestBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAnnouncementDestBag.Unmarshal(m, b)
}
func (m *MplsTeAnnouncementDestBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAnnouncementDestBag.Marshal(b, m, deterministic)
}
func (dst *MplsTeAnnouncementDestBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAnnouncementDestBag.Merge(dst, src)
}
func (m *MplsTeAnnouncementDestBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeAnnouncementDestBag.Size(m)
}
func (m *MplsTeAnnouncementDestBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAnnouncementDestBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAnnouncementDestBag proto.InternalMessageInfo

func (m *MplsTeAnnouncementDestBag) GetIgpType() string {
	if m != nil {
		return m.IgpType
	}
	return ""
}

func (m *MplsTeAnnouncementDestBag) GetIgpInstance() string {
	if m != nil {
		return m.IgpInstance
	}
	return ""
}

func (m *MplsTeAnnouncementDestBag) GetIgpArea() uint32 {
	if m != nil {
		return m.IgpArea
	}
	return 0
}

func (m *MplsTeAnnouncementDestBag) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *MplsTeAnnouncementDestBag) GetAnnouncedTunnelsCount() uint32 {
	if m != nil {
		return m.AnnouncedTunnelsCount
	}
	return 0
}

func (m *MplsTeAnnouncementDestBag) GetAutoroutedTunnelsCount() uint32 {
	if m != nil {
		return m.AutoroutedTunnelsCount
	}
	return 0
}

func (m *MplsTeAnnouncementDestBag) GetForwardingAdjacencyTunnelsCount() uint32 {
	if m != nil {
		return m.ForwardingAdjacencyTunnelsCount
	}
	return 0
}

func (m *MplsTeAnnouncementDestBag) GetAreaFormat() string {
	if m != nil {
		return m.AreaFormat
	}
	return ""
}

func (m *MplsTeAnnouncementDestBag) GetTotalTunnels() []*MplsTeAutorouteBag {
	if m != nil {
		return m.TotalTunnels
	}
	return nil
}

// IGP area information including AFI
type MplsTeAreaAfiInfo struct {
	// The IGP Instance and Area ID
	IgpAreaId string `protobuf:"bytes,1,opt,name=igp_area_id,json=igpAreaId" json:"igp_area_id,omitempty"`
	// AFI
	Afi                  string   `protobuf:"bytes,2,opt,name=afi" json:"afi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeAreaAfiInfo) Reset()         { *m = MplsTeAreaAfiInfo{} }
func (m *MplsTeAreaAfiInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeAreaAfiInfo) ProtoMessage()    {}
func (*MplsTeAreaAfiInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_announcement_dest_bag_be31fd0d431ba514, []int{2}
}
func (m *MplsTeAreaAfiInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAreaAfiInfo.Unmarshal(m, b)
}
func (m *MplsTeAreaAfiInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAreaAfiInfo.Marshal(b, m, deterministic)
}
func (dst *MplsTeAreaAfiInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAreaAfiInfo.Merge(dst, src)
}
func (m *MplsTeAreaAfiInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeAreaAfiInfo.Size(m)
}
func (m *MplsTeAreaAfiInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAreaAfiInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAreaAfiInfo proto.InternalMessageInfo

func (m *MplsTeAreaAfiInfo) GetIgpAreaId() string {
	if m != nil {
		return m.IgpAreaId
	}
	return ""
}

func (m *MplsTeAreaAfiInfo) GetAfi() string {
	if m != nil {
		return m.Afi
	}
	return ""
}

// MPLS TE Autoroute Information
type MplsTeAutorouteBag struct {
	// The tunnel name
	TunnelName string `protobuf:"bytes,1,opt,name=tunnel_name,json=tunnelName" json:"tunnel_name,omitempty"`
	// The tunnel signaled-name
	TunnelSigName string `protobuf:"bytes,2,opt,name=tunnel_sig_name,json=tunnelSigName" json:"tunnel_sig_name,omitempty"`
	// Tunnel loadshare
	TunnelLoadshare uint32 `protobuf:"varint,3,opt,name=tunnel_loadshare,json=tunnelLoadshare" json:"tunnel_loadshare,omitempty"`
	// Announcement type
	AnnounceType string `protobuf:"bytes,4,opt,name=announce_type,json=announceType" json:"announce_type,omitempty"`
	// Autoroute mode
	Mode string `protobuf:"bytes,5,opt,name=mode" json:"mode,omitempty"`
	// IGP metric
	IgpMetric int32 `protobuf:"zigzag32,6,opt,name=igp_metric,json=igpMetric" json:"igp_metric,omitempty"`
	// Hold time in seconds
	HoldTime uint32 `protobuf:"varint,7,opt,name=hold_time,json=holdTime" json:"hold_time,omitempty"`
	// IGP Area Format
	AreaFormat string `protobuf:"bytes,8,opt,name=area_format,json=areaFormat" json:"area_format,omitempty"`
	// True if this is auto-tunnel mesh AA
	IsAutoMeshAa bool `protobuf:"varint,9,opt,name=is_auto_mesh_aa,json=isAutoMeshAa" json:"is_auto_mesh_aa,omitempty"`
	// True if this is a segment routing tunnel
	IsSr bool `protobuf:"varint,10,opt,name=is_sr,json=isSr" json:"is_sr,omitempty"`
	// True if this is a strict SID SR tunnel
	IsSrStrict bool `protobuf:"varint,11,opt,name=is_sr_strict,json=isSrStrict" json:"is_sr_strict,omitempty"`
	// MeshGroup ID for AA
	MeshGroupIdaa uint32 `protobuf:"varint,12,opt,name=mesh_group_idaa,json=meshGroupIdaa" json:"mesh_group_idaa,omitempty"`
	// List of IGPs to which it is announced
	IgPs                 []*MplsTeAreaAfiInfo `protobuf:"bytes,13,rep,name=ig_ps,json=igPs" json:"ig_ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MplsTeAutorouteBag) Reset()         { *m = MplsTeAutorouteBag{} }
func (m *MplsTeAutorouteBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeAutorouteBag) ProtoMessage()    {}
func (*MplsTeAutorouteBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_announcement_dest_bag_be31fd0d431ba514, []int{3}
}
func (m *MplsTeAutorouteBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAutorouteBag.Unmarshal(m, b)
}
func (m *MplsTeAutorouteBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAutorouteBag.Marshal(b, m, deterministic)
}
func (dst *MplsTeAutorouteBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAutorouteBag.Merge(dst, src)
}
func (m *MplsTeAutorouteBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeAutorouteBag.Size(m)
}
func (m *MplsTeAutorouteBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAutorouteBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAutorouteBag proto.InternalMessageInfo

func (m *MplsTeAutorouteBag) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *MplsTeAutorouteBag) GetTunnelSigName() string {
	if m != nil {
		return m.TunnelSigName
	}
	return ""
}

func (m *MplsTeAutorouteBag) GetTunnelLoadshare() uint32 {
	if m != nil {
		return m.TunnelLoadshare
	}
	return 0
}

func (m *MplsTeAutorouteBag) GetAnnounceType() string {
	if m != nil {
		return m.AnnounceType
	}
	return ""
}

func (m *MplsTeAutorouteBag) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *MplsTeAutorouteBag) GetIgpMetric() int32 {
	if m != nil {
		return m.IgpMetric
	}
	return 0
}

func (m *MplsTeAutorouteBag) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *MplsTeAutorouteBag) GetAreaFormat() string {
	if m != nil {
		return m.AreaFormat
	}
	return ""
}

func (m *MplsTeAutorouteBag) GetIsAutoMeshAa() bool {
	if m != nil {
		return m.IsAutoMeshAa
	}
	return false
}

func (m *MplsTeAutorouteBag) GetIsSr() bool {
	if m != nil {
		return m.IsSr
	}
	return false
}

func (m *MplsTeAutorouteBag) GetIsSrStrict() bool {
	if m != nil {
		return m.IsSrStrict
	}
	return false
}

func (m *MplsTeAutorouteBag) GetMeshGroupIdaa() uint32 {
	if m != nil {
		return m.MeshGroupIdaa
	}
	return 0
}

func (m *MplsTeAutorouteBag) GetIgPs() []*MplsTeAreaAfiInfo {
	if m != nil {
		return m.IgPs
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsTeAnnouncementDestBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.announce_tunnels_info.autoroute_announce_table.announce_dest.mpls_te_announcement_dest_bag_KEYS")
	proto.RegisterType((*MplsTeAnnouncementDestBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.announce_tunnels_info.autoroute_announce_table.announce_dest.mpls_te_announcement_dest_bag")
	proto.RegisterType((*MplsTeAreaAfiInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.announce_tunnels_info.autoroute_announce_table.announce_dest.mpls_te_area_afi_info")
	proto.RegisterType((*MplsTeAutorouteBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.announce_tunnels_info.autoroute_announce_table.announce_dest.mpls_te_autoroute_bag")
}

func init() {
	proto.RegisterFile("mpls_te_announcement_dest_bag.proto", fileDescriptor_mpls_te_announcement_dest_bag_be31fd0d431ba514)
}

var fileDescriptor_mpls_te_announcement_dest_bag_be31fd0d431ba514 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0x95, 0x6f, 0x93, 0x36, 0x99, 0x24, 0x6a, 0xef, 0x56, 0xbd, 0x77, 0xef, 0x45, 0xa5, 0x21,
	0x15, 0x28, 0xbc, 0x04, 0xa9, 0x85, 0x52, 0x78, 0x8b, 0x10, 0xa0, 0xa8, 0x14, 0xa1, 0xa4, 0x2f,
	0xbc, 0xb0, 0x9a, 0xd8, 0x1b, 0x67, 0x51, 0xbc, 0x6b, 0x79, 0xd7, 0x82, 0xfe, 0x02, 0x3f, 0x80,
	0x84, 0xf8, 0x0c, 0x3e, 0x10, 0xed, 0xd8, 0x4e, 0xdd, 0x22, 0xf5, 0x0d, 0xde, 0x36, 0x67, 0xce,
	0x39, 0x3b, 0x33, 0x7b, 0x62, 0x38, 0x4c, 0xd2, 0x95, 0x15, 0x4e, 0x0a, 0xd4, 0xda, 0xe4, 0x3a,
	0x94, 0x89, 0xd4, 0x4e, 0x44, 0xd2, 0x3a, 0x31, 0xc7, 0x78, 0x94, 0x66, 0xc6, 0x19, 0xf6, 0x21,
	0x54, 0x36, 0x34, 0x42, 0x19, 0x2b, 0x3e, 0x67, 0xa2, 0x52, 0x98, 0x54, 0x66, 0xa3, 0xf2, 0xc7,
	0xa8, 0x92, 0x0b, 0x97, 0x6b, 0x2d, 0x57, 0x56, 0x28, 0xbd, 0x30, 0x23, 0xcc, 0x9d, 0xc9, 0x4c,
	0x5e, 0xb3, 0x17, 0x0e, 0xe7, 0xab, 0x1a, 0xdd, 0xdf, 0x34, 0xf8, 0x1e, 0xc0, 0xe0, 0xd6, 0x3e,
	0xc4, 0xd9, 0xcb, 0xf7, 0x33, 0xf6, 0x08, 0x76, 0x3d, 0xa0, 0x34, 0x3a, 0x65, 0xb4, 0xc0, 0x28,
	0xca, 0xa4, 0xb5, 0x3c, 0xe8, 0x07, 0xc3, 0xf6, 0x94, 0xd5, 0x4a, 0xe3, 0xa2, 0xc2, 0xfe, 0x87,
	0x16, 0x0d, 0x10, 0x9a, 0x15, 0xff, 0x8b, 0x58, 0xeb, 0xdf, 0x8c, 0x41, 0x03, 0x33, 0x89, 0x7c,
	0xa3, 0x1f, 0x0c, 0x7b, 0x53, 0x3a, 0xb3, 0x3d, 0xd8, 0x54, 0x71, 0x2a, 0x54, 0xc4, 0x1b, 0xc4,
	0x6e, 0xaa, 0x38, 0x9d, 0x44, 0x83, 0xaf, 0x0d, 0xd8, 0xbf, 0xb5, 0x3d, 0xf6, 0x1f, 0xb4, 0xbc,
	0xd0, 0x5d, 0xa6, 0x92, 0x1f, 0x91, 0x74, 0x4b, 0xc5, 0xe9, 0xc5, 0x65, 0x2a, 0xd9, 0x3d, 0xe8,
	0x92, 0xa7, 0xb6, 0x0e, 0x75, 0x28, 0xf9, 0x31, 0x95, 0x3b, 0xde, 0xb9, 0x84, 0x2a, 0x35, 0xb5,
	0xf3, 0x98, 0xda, 0xf1, 0xea, 0xb1, 0xef, 0xa8, 0x0f, 0x9d, 0xda, 0x5c, 0xfc, 0x49, 0x21, 0xae,
	0x41, 0xec, 0x04, 0xfe, 0xad, 0x7a, 0x8a, 0xd6, 0xcb, 0x0f, 0x4d, 0xae, 0x1d, 0x3f, 0x21, 0xaf,
	0xbd, 0x75, 0xf9, 0xa2, 0xa8, 0xbe, 0xf0, 0x45, 0x76, 0x0a, 0x7c, 0xfd, 0x3a, 0x37, 0x85, 0x4f,
	0x49, 0xf8, 0xcf, 0x55, 0xfd, 0x9a, 0xf2, 0x0c, 0x06, 0x0b, 0x93, 0x7d, 0xc2, 0x2c, 0x52, 0x3a,
	0x16, 0x18, 0x7d, 0xc4, 0x50, 0xea, 0xf0, 0xf2, 0x86, 0xc7, 0x29, 0x79, 0x1c, 0x5c, 0x31, 0xc7,
	0x15, 0xf1, 0x9a, 0xd9, 0x01, 0x74, 0xfc, 0xdc, 0x62, 0x61, 0xb2, 0x04, 0x1d, 0x7f, 0x46, 0x03,
	0x82, 0x87, 0x5e, 0x11, 0xc2, 0xbe, 0x05, 0xd0, 0x73, 0xc6, 0xe1, 0xaa, 0xf2, 0xe7, 0xcf, 0xfb,
	0x1b, 0xc3, 0xce, 0x51, 0x3e, 0xfa, 0xbd, 0xa1, 0x1c, 0xad, 0x5f, 0x7c, 0x4d, 0x9f, 0x63, 0x3c,
	0xed, 0x52, 0x2f, 0xe5, 0x08, 0x83, 0x09, 0xec, 0xad, 0x69, 0x7e, 0x0a, 0x5c, 0x28, 0xba, 0x83,
	0xdd, 0x85, 0x4e, 0xf5, 0xa4, 0x3e, 0x4e, 0x45, 0x44, 0xdb, 0xe5, 0xab, 0x4e, 0x22, 0xb6, 0x03,
	0x1b, 0xb8, 0x50, 0x65, 0x28, 0xfd, 0x71, 0xf0, 0xa3, 0x51, 0xf3, 0xaa, 0x5f, 0xe9, 0x57, 0x54,
	0xf4, 0x2f, 0x34, 0x26, 0xb2, 0xf4, 0x82, 0x02, 0x7a, 0x8b, 0x89, 0x64, 0x0f, 0x60, 0xbb, 0x24,
	0x58, 0x15, 0x17, 0xa4, 0xc2, 0xb8, 0x57, 0xc0, 0x33, 0x15, 0x13, 0xef, 0x21, 0xec, 0x94, 0xbc,
	0x95, 0xc1, 0xc8, 0x2e, 0x31, 0x93, 0x65, 0xfc, 0x4b, 0xfd, 0x9b, 0x0a, 0x66, 0x87, 0xd0, 0xbb,
	0x5a, 0x8e, 0x4f, 0x75, 0xf1, 0x87, 0xe8, 0x56, 0x20, 0x45, 0x9b, 0x41, 0x23, 0x31, 0x91, 0xe4,
	0x4d, 0xaa, 0xd1, 0x99, 0xed, 0x03, 0xf8, 0xc1, 0x13, 0xe9, 0x32, 0x15, 0xf2, 0xcd, 0x7e, 0x30,
	0xfc, 0x9b, 0xe6, 0x3e, 0x27, 0x80, 0xdd, 0x81, 0xf6, 0xd2, 0xac, 0x22, 0xe1, 0x54, 0x22, 0xf9,
	0x16, 0xdd, 0xdd, 0xf2, 0xc0, 0x85, 0x4a, 0xe4, 0xcd, 0x2c, 0xb4, 0x7e, 0xc9, 0xc2, 0x7d, 0xd8,
	0x56, 0x96, 0xb6, 0x23, 0x12, 0x69, 0x97, 0x02, 0x91, 0xb7, 0xfb, 0xc1, 0xb0, 0x35, 0xed, 0x2a,
	0x3b, 0xce, 0x9d, 0x39, 0x97, 0x76, 0x39, 0x46, 0xb6, 0x0b, 0x4d, 0x65, 0x85, 0xcd, 0x38, 0x50,
	0xb1, 0xa1, 0xec, 0x2c, 0x63, 0x7d, 0xe8, 0x12, 0x28, 0xac, 0x6f, 0xc4, 0xf1, 0x0e, 0xd5, 0xc0,
	0xd7, 0x66, 0x84, 0xf8, 0x35, 0x92, 0x6b, 0x9c, 0x99, 0xdc, 0x7f, 0x04, 0x10, 0x79, 0x97, 0x3a,
	0xec, 0x79, 0xf8, 0xb5, 0x47, 0x27, 0x11, 0x22, 0xfb, 0x12, 0x40, 0x53, 0xc5, 0x22, 0xb5, 0xbc,
	0xf7, 0x87, 0x93, 0x58, 0x8f, 0xd8, 0xb4, 0xa1, 0xe2, 0x77, 0x76, 0xbe, 0x49, 0x1f, 0xb4, 0xe3,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xce, 0xca, 0x07, 0xc8, 0x05, 0x00, 0x00,
}
