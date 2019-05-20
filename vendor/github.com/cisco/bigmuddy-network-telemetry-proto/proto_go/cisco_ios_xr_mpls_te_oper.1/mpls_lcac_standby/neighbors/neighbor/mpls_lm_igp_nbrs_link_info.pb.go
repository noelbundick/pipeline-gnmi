// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_lm_igp_nbrs_link_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_lcac_standby_neighbors_neighbor

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

// Neighbor information based on the link
type MplsLmIgpNbrsLinkInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmIgpNbrsLinkInfo_KEYS) Reset()         { *m = MplsLmIgpNbrsLinkInfo_KEYS{} }
func (m *MplsLmIgpNbrsLinkInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLmIgpNbrsLinkInfo_KEYS) ProtoMessage()    {}
func (*MplsLmIgpNbrsLinkInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_lm_igp_nbrs_link_info_1ba3d008fc654925, []int{0}
}
func (m *MplsLmIgpNbrsLinkInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmIgpNbrsLinkInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsLmIgpNbrsLinkInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmIgpNbrsLinkInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsLmIgpNbrsLinkInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmIgpNbrsLinkInfo_KEYS.Merge(dst, src)
}
func (m *MplsLmIgpNbrsLinkInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLmIgpNbrsLinkInfo_KEYS.Size(m)
}
func (m *MplsLmIgpNbrsLinkInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmIgpNbrsLinkInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmIgpNbrsLinkInfo_KEYS proto.InternalMessageInfo

func (m *MplsLmIgpNbrsLinkInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type MplsLmIgpNbrsLinkInfo struct {
	// Neighbors of the specified link id
	Neighbors            []*MplsLmIgpNbrInfo `protobuf:"bytes,50,rep,name=neighbors" json:"neighbors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MplsLmIgpNbrsLinkInfo) Reset()         { *m = MplsLmIgpNbrsLinkInfo{} }
func (m *MplsLmIgpNbrsLinkInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmIgpNbrsLinkInfo) ProtoMessage()    {}
func (*MplsLmIgpNbrsLinkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_lm_igp_nbrs_link_info_1ba3d008fc654925, []int{1}
}
func (m *MplsLmIgpNbrsLinkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmIgpNbrsLinkInfo.Unmarshal(m, b)
}
func (m *MplsLmIgpNbrsLinkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmIgpNbrsLinkInfo.Marshal(b, m, deterministic)
}
func (dst *MplsLmIgpNbrsLinkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmIgpNbrsLinkInfo.Merge(dst, src)
}
func (m *MplsLmIgpNbrsLinkInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmIgpNbrsLinkInfo.Size(m)
}
func (m *MplsLmIgpNbrsLinkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmIgpNbrsLinkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmIgpNbrsLinkInfo proto.InternalMessageInfo

func (m *MplsLmIgpNbrsLinkInfo) GetNeighbors() []*MplsLmIgpNbrInfo {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

// Neighbor information
type MplsLmIgpNbrInfo struct {
	// The interface on which this neighbor is discovered
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// The neighbor's id
	NeighborId string `protobuf:"bytes,2,opt,name=neighbor_id,json=neighborId" json:"neighbor_id,omitempty"`
	// The area id
	AreaId string `protobuf:"bytes,3,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	// The neighbor's IP address
	NeighborAddress      string   `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmIgpNbrInfo) Reset()         { *m = MplsLmIgpNbrInfo{} }
func (m *MplsLmIgpNbrInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmIgpNbrInfo) ProtoMessage()    {}
func (*MplsLmIgpNbrInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_lm_igp_nbrs_link_info_1ba3d008fc654925, []int{2}
}
func (m *MplsLmIgpNbrInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmIgpNbrInfo.Unmarshal(m, b)
}
func (m *MplsLmIgpNbrInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmIgpNbrInfo.Marshal(b, m, deterministic)
}
func (dst *MplsLmIgpNbrInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmIgpNbrInfo.Merge(dst, src)
}
func (m *MplsLmIgpNbrInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmIgpNbrInfo.Size(m)
}
func (m *MplsLmIgpNbrInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmIgpNbrInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmIgpNbrInfo proto.InternalMessageInfo

func (m *MplsLmIgpNbrInfo) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *MplsLmIgpNbrInfo) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *MplsLmIgpNbrInfo) GetAreaId() string {
	if m != nil {
		return m.AreaId
	}
	return ""
}

func (m *MplsLmIgpNbrInfo) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsLmIgpNbrsLinkInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.neighbors.neighbor.mpls_lm_igp_nbrs_link_info_KEYS")
	proto.RegisterType((*MplsLmIgpNbrsLinkInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.neighbors.neighbor.mpls_lm_igp_nbrs_link_info")
	proto.RegisterType((*MplsLmIgpNbrInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.neighbors.neighbor.mpls_lm_igp_nbr_info")
}

func init() {
	proto.RegisterFile("mpls_lm_igp_nbrs_link_info.proto", fileDescriptor_mpls_lm_igp_nbrs_link_info_1ba3d008fc654925)
}

var fileDescriptor_mpls_lm_igp_nbrs_link_info_1ba3d008fc654925 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0x59, 0x2b, 0x95, 0x4e, 0xf1, 0x0f, 0x41, 0x70, 0xf1, 0xd2, 0x65, 0x41, 0xa8, 0x97,
	0x1c, 0xea, 0x5d, 0xf0, 0x20, 0x58, 0x04, 0x0f, 0xd5, 0x8b, 0xa7, 0x21, 0x9b, 0x4c, 0x6b, 0x70,
	0x37, 0x59, 0x26, 0x7b, 0xd0, 0xb7, 0xf0, 0x19, 0x7c, 0x52, 0x69, 0xca, 0x46, 0x10, 0x17, 0x04,
	0x6f, 0xc9, 0x6f, 0xbe, 0x99, 0x2f, 0x61, 0xa0, 0x68, 0xda, 0x3a, 0x60, 0xdd, 0xa0, 0xdd, 0xb4,
	0xe8, 0x2a, 0x0e, 0x58, 0x5b, 0xf7, 0x8a, 0xd6, 0xad, 0xbd, 0x6c, 0xd9, 0x77, 0x5e, 0x5c, 0x6b,
	0x1b, 0xb4, 0x47, 0xeb, 0x03, 0xbe, 0x31, 0x46, 0xbc, 0x23, 0xf4, 0x2d, 0xb1, 0xdc, 0xf5, 0x6a,
	0xa5, 0x31, 0x74, 0xca, 0x99, 0xea, 0x5d, 0x3a, 0xb2, 0x9b, 0x97, 0xca, 0x73, 0x48, 0xa7, 0xf2,
	0x0e, 0x66, 0xc3, 0x0e, 0xbc, 0xbf, 0x7d, 0x7e, 0x14, 0x17, 0x70, 0x64, 0x5d, 0x47, 0xbc, 0x56,
	0x9a, 0xd0, 0xa9, 0x86, 0xf2, 0xac, 0xc8, 0xe6, 0x93, 0xd5, 0x61, 0x4a, 0x1f, 0x54, 0x43, 0xe5,
	0x47, 0x06, 0xe7, 0xc3, 0xa3, 0x04, 0xc3, 0x24, 0xe9, 0xf3, 0x45, 0x31, 0x9a, 0x4f, 0x17, 0x4f,
	0xf2, 0x7f, 0x8f, 0x97, 0x3f, 0x74, 0x51, 0xb4, 0xfa, 0xd6, 0x94, 0x9f, 0x19, 0x9c, 0xfe, 0xc6,
	0xfc, 0xf1, 0x4b, 0x62, 0x06, 0xd3, 0x7e, 0x18, 0x5a, 0x93, 0xef, 0x45, 0x06, 0xfa, 0x68, 0x69,
	0xc4, 0x19, 0x1c, 0x28, 0x26, 0xb5, 0x2d, 0x8e, 0x62, 0x71, 0xbc, 0xbd, 0x2e, 0x8d, 0xb8, 0x84,
	0x93, 0xd4, 0xa9, 0x8c, 0x61, 0x0a, 0x21, 0xdf, 0x8f, 0xc4, 0x71, 0x9f, 0xdf, 0xec, 0xe2, 0x6a,
	0x1c, 0x17, 0x79, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x74, 0x4d, 0x63, 0xec, 0x01, 0x00,
	0x00,
}