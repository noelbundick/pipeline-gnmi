// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_lm_summary_common_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_lcac_link_summary

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

// Summary information for all the areas
type MplsLmSummaryCommonInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmSummaryCommonInfo_KEYS) Reset()         { *m = MplsLmSummaryCommonInfo_KEYS{} }
func (m *MplsLmSummaryCommonInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLmSummaryCommonInfo_KEYS) ProtoMessage()    {}
func (*MplsLmSummaryCommonInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_lm_summary_common_info_a897a8843af9f76b, []int{0}
}
func (m *MplsLmSummaryCommonInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsLmSummaryCommonInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsLmSummaryCommonInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS.Merge(dst, src)
}
func (m *MplsLmSummaryCommonInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS.Size(m)
}
func (m *MplsLmSummaryCommonInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS proto.InternalMessageInfo

type MplsLmSummaryCommonInfo struct {
	// TRUE if Role is Standby, Active otherwise
	IsRoleStandby bool `protobuf:"varint,50,opt,name=is_role_standby,json=isRoleStandby" json:"is_role_standby,omitempty"`
	// Total number of links
	Links uint32 `protobuf:"varint,51,opt,name=links" json:"links,omitempty"`
	// Maximum number of links supported
	MaximumLinks uint32 `protobuf:"varint,52,opt,name=maximum_links,json=maximumLinks" json:"maximum_links,omitempty"`
	// TRUE if flooding is enabled
	IsFloodingEnabled bool `protobuf:"varint,53,opt,name=is_flooding_enabled,json=isFloodingEnabled" json:"is_flooding_enabled,omitempty"`
	// Summary info for the areas
	AreasSummary         []*MplsLmSummaryAreaInfo `protobuf:"bytes,54,rep,name=areas_summary,json=areasSummary" json:"areas_summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MplsLmSummaryCommonInfo) Reset()         { *m = MplsLmSummaryCommonInfo{} }
func (m *MplsLmSummaryCommonInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmSummaryCommonInfo) ProtoMessage()    {}
func (*MplsLmSummaryCommonInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_lm_summary_common_info_a897a8843af9f76b, []int{1}
}
func (m *MplsLmSummaryCommonInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmSummaryCommonInfo.Unmarshal(m, b)
}
func (m *MplsLmSummaryCommonInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmSummaryCommonInfo.Marshal(b, m, deterministic)
}
func (dst *MplsLmSummaryCommonInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmSummaryCommonInfo.Merge(dst, src)
}
func (m *MplsLmSummaryCommonInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmSummaryCommonInfo.Size(m)
}
func (m *MplsLmSummaryCommonInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmSummaryCommonInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmSummaryCommonInfo proto.InternalMessageInfo

func (m *MplsLmSummaryCommonInfo) GetIsRoleStandby() bool {
	if m != nil {
		return m.IsRoleStandby
	}
	return false
}

func (m *MplsLmSummaryCommonInfo) GetLinks() uint32 {
	if m != nil {
		return m.Links
	}
	return 0
}

func (m *MplsLmSummaryCommonInfo) GetMaximumLinks() uint32 {
	if m != nil {
		return m.MaximumLinks
	}
	return 0
}

func (m *MplsLmSummaryCommonInfo) GetIsFloodingEnabled() bool {
	if m != nil {
		return m.IsFloodingEnabled
	}
	return false
}

func (m *MplsLmSummaryCommonInfo) GetAreasSummary() []*MplsLmSummaryAreaInfo {
	if m != nil {
		return m.AreasSummary
	}
	return nil
}

// Summary information of an area
type MplsLmSummaryAreaInfo struct {
	// Area id
	AreaId string `protobuf:"bytes,1,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	// Protocol running over the area
	Protocol string `protobuf:"bytes,2,opt,name=protocol" json:"protocol,omitempty"`
	// TRUE if flooding occurred in this area
	IsFlooded bool `protobuf:"varint,3,opt,name=is_flooded,json=isFlooded" json:"is_flooded,omitempty"`
	// TRUE if periodic flooding is on
	IsPeriodicFloodingOn bool `protobuf:"varint,4,opt,name=is_periodic_flooding_on,json=isPeriodicFloodingOn" json:"is_periodic_flooding_on,omitempty"`
	// Flooding period in seconds
	PeriodicFloodingInterval uint32 `protobuf:"varint,5,opt,name=periodic_flooding_interval,json=periodicFloodingInterval" json:"periodic_flooding_interval,omitempty"`
	// Number of flooded links
	LinksFlooded uint32 `protobuf:"varint,6,opt,name=links_flooded,json=linksFlooded" json:"links_flooded,omitempty"`
	// IGP id of local node
	SystemId string `protobuf:"bytes,7,opt,name=system_id,json=systemId" json:"system_id,omitempty"`
	// Local router id
	LocalNodeRouterId string `protobuf:"bytes,8,opt,name=local_node_router_id,json=localNodeRouterId" json:"local_node_router_id,omitempty"`
	// Number of IGP neighbors
	IgpNeighbors         uint32   `protobuf:"varint,9,opt,name=igp_neighbors,json=igpNeighbors" json:"igp_neighbors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmSummaryAreaInfo) Reset()         { *m = MplsLmSummaryAreaInfo{} }
func (m *MplsLmSummaryAreaInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmSummaryAreaInfo) ProtoMessage()    {}
func (*MplsLmSummaryAreaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_lm_summary_common_info_a897a8843af9f76b, []int{2}
}
func (m *MplsLmSummaryAreaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmSummaryAreaInfo.Unmarshal(m, b)
}
func (m *MplsLmSummaryAreaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmSummaryAreaInfo.Marshal(b, m, deterministic)
}
func (dst *MplsLmSummaryAreaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmSummaryAreaInfo.Merge(dst, src)
}
func (m *MplsLmSummaryAreaInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmSummaryAreaInfo.Size(m)
}
func (m *MplsLmSummaryAreaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmSummaryAreaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmSummaryAreaInfo proto.InternalMessageInfo

func (m *MplsLmSummaryAreaInfo) GetAreaId() string {
	if m != nil {
		return m.AreaId
	}
	return ""
}

func (m *MplsLmSummaryAreaInfo) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *MplsLmSummaryAreaInfo) GetIsFlooded() bool {
	if m != nil {
		return m.IsFlooded
	}
	return false
}

func (m *MplsLmSummaryAreaInfo) GetIsPeriodicFloodingOn() bool {
	if m != nil {
		return m.IsPeriodicFloodingOn
	}
	return false
}

func (m *MplsLmSummaryAreaInfo) GetPeriodicFloodingInterval() uint32 {
	if m != nil {
		return m.PeriodicFloodingInterval
	}
	return 0
}

func (m *MplsLmSummaryAreaInfo) GetLinksFlooded() uint32 {
	if m != nil {
		return m.LinksFlooded
	}
	return 0
}

func (m *MplsLmSummaryAreaInfo) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *MplsLmSummaryAreaInfo) GetLocalNodeRouterId() string {
	if m != nil {
		return m.LocalNodeRouterId
	}
	return ""
}

func (m *MplsLmSummaryAreaInfo) GetIgpNeighbors() uint32 {
	if m != nil {
		return m.IgpNeighbors
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsLmSummaryCommonInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.link_summary.mpls_lm_summary_common_info_KEYS")
	proto.RegisterType((*MplsLmSummaryCommonInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.link_summary.mpls_lm_summary_common_info")
	proto.RegisterType((*MplsLmSummaryAreaInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.link_summary.mpls_lm_summary_area_info")
}

func init() {
	proto.RegisterFile("mpls_lm_summary_common_info.proto", fileDescriptor_mpls_lm_summary_common_info_a897a8843af9f76b)
}

var fileDescriptor_mpls_lm_summary_common_info_a897a8843af9f76b = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0x26, 0xa9, 0x4d, 0x93, 0xb5, 0x41, 0xba, 0x06, 0xba, 0xb6, 0x08, 0x67, 0x04, 0xc9, 0xd3,
	0x29, 0xad, 0xf5, 0xc9, 0xd7, 0x0a, 0xa1, 0x52, 0xe5, 0xf2, 0xe4, 0xd3, 0x70, 0xb9, 0xdd, 0xc6,
	0xc1, 0xdd, 0x9d, 0x63, 0xf7, 0x22, 0xcd, 0xdf, 0xf1, 0x77, 0xfa, 0x20, 0x37, 0x7b, 0x57, 0xa1,
	0x62, 0xa1, 0x8f, 0x33, 0xdf, 0x37, 0xf3, 0xcd, 0xf7, 0x8d, 0x78, 0xe5, 0x6a, 0x1b, 0xc1, 0x3a,
	0x88, 0x5b, 0xe7, 0xca, 0xb0, 0x83, 0x8a, 0x9c, 0x23, 0x0f, 0xe8, 0x6f, 0x28, 0xaf, 0x03, 0x35,
	0x24, 0xdf, 0x55, 0x18, 0x2b, 0x02, 0xa4, 0x08, 0xb7, 0x01, 0x98, 0xdf, 0x18, 0xa0, 0xda, 0x84,
	0x3c, 0x0d, 0x57, 0x65, 0x95, 0x5b, 0xf4, 0x3f, 0xfa, 0x1d, 0xf3, 0xb9, 0xc8, 0x1e, 0x58, 0x0b,
	0x57, 0x97, 0xdf, 0x56, 0xf3, 0x5f, 0x43, 0x71, 0xfa, 0x00, 0x49, 0xbe, 0x11, 0xcf, 0x30, 0x42,
	0x20, 0x6b, 0x20, 0x36, 0xa5, 0xd7, 0xeb, 0x9d, 0x3a, 0xcb, 0x06, 0x8b, 0x71, 0x31, 0xc5, 0x58,
	0x90, 0x35, 0xab, 0xd4, 0x94, 0x33, 0xb1, 0xdf, 0x6a, 0x47, 0x75, 0x9e, 0x0d, 0x16, 0xd3, 0x22,
	0x15, 0xf2, 0xb5, 0x98, 0xba, 0xf2, 0x16, 0xdd, 0xd6, 0x41, 0x42, 0xdf, 0x33, 0x7a, 0xd8, 0x35,
	0x3f, 0x33, 0x29, 0x17, 0xcf, 0x31, 0xc2, 0x8d, 0x25, 0xd2, 0xe8, 0x37, 0x60, 0x7c, 0xb9, 0xb6,
	0x46, 0xab, 0x0b, 0x96, 0x39, 0xc2, 0xf8, 0xa9, 0x43, 0x2e, 0x13, 0x20, 0x6b, 0x31, 0x2d, 0x83,
	0x29, 0x63, 0x7f, 0xaf, 0xfa, 0x90, 0xed, 0x2d, 0x9e, 0x9e, 0x5d, 0xe5, 0x8f, 0x0d, 0x28, 0xbf,
	0x6f, 0xbc, 0x5d, 0xcb, 0xb6, 0x8b, 0x43, 0x56, 0x58, 0x75, 0x41, 0xfe, 0x1e, 0x8a, 0x17, 0xff,
	0xe5, 0xca, 0x63, 0x71, 0x90, 0x0a, 0xad, 0x06, 0xd9, 0x60, 0x31, 0x29, 0x46, 0x6d, 0xb9, 0xd4,
	0xf2, 0x44, 0x8c, 0xf9, 0x75, 0x15, 0x59, 0x35, 0x64, 0xe4, 0xae, 0x96, 0x2f, 0x85, 0xe8, 0x4d,
	0x1b, 0xad, 0xf6, 0xd8, 0xeb, 0xa4, 0xf3, 0x6a, 0xb4, 0xbc, 0x10, 0xc7, 0x18, 0xa1, 0x36, 0x01,
	0x49, 0x63, 0xf5, 0x37, 0x1c, 0xf2, 0xea, 0x09, 0x73, 0x67, 0x18, 0xbf, 0x76, 0x68, 0x9f, 0xcf,
	0x17, 0x2f, 0x3f, 0x8a, 0x93, 0x7f, 0x67, 0xd0, 0x37, 0x26, 0xfc, 0x2c, 0xad, 0xda, 0xe7, 0xf0,
	0x55, 0x7d, 0x6f, 0x6e, 0xd9, 0xe1, 0xed, 0xb7, 0xf8, 0x4b, 0x77, 0x67, 0x8d, 0xd2, 0xb7, 0xb8,
	0xd9, 0x5f, 0x76, 0x2a, 0x26, 0x71, 0x17, 0x1b, 0xe3, 0x5a, 0xbf, 0x07, 0xc9, 0x55, 0x6a, 0x2c,
	0xb5, 0x7c, 0x2b, 0x66, 0x96, 0xaa, 0xd2, 0x82, 0x27, 0x6d, 0x20, 0xd0, 0xb6, 0x31, 0xa1, 0xe5,
	0x8d, 0x99, 0x77, 0xc4, 0xd8, 0x35, 0x69, 0x53, 0x30, 0xb2, 0xd4, 0xad, 0x24, 0x6e, 0x6a, 0xf0,
	0x06, 0x37, 0xdf, 0xd7, 0x14, 0xa2, 0x9a, 0x24, 0x49, 0xdc, 0xd4, 0xd7, 0x7d, 0x6f, 0x3d, 0xe2,
	0xd4, 0xce, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x5e, 0xa7, 0x0e, 0x25, 0x03, 0x00, 0x00,
}