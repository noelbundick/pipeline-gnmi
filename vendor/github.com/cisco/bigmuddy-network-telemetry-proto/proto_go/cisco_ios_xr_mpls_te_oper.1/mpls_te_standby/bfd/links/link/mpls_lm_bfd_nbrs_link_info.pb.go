// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_lm_bfd_nbrs_link_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_bfd_links_link

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

// BFD Neighbor information based on the link
type MplsLmBfdNbrsLinkInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmBfdNbrsLinkInfo_KEYS) Reset()         { *m = MplsLmBfdNbrsLinkInfo_KEYS{} }
func (m *MplsLmBfdNbrsLinkInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLmBfdNbrsLinkInfo_KEYS) ProtoMessage()    {}
func (*MplsLmBfdNbrsLinkInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_lm_bfd_nbrs_link_info_2ce6ef9822fd4219, []int{0}
}
func (m *MplsLmBfdNbrsLinkInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsLmBfdNbrsLinkInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsLmBfdNbrsLinkInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS.Merge(dst, src)
}
func (m *MplsLmBfdNbrsLinkInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS.Size(m)
}
func (m *MplsLmBfdNbrsLinkInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS proto.InternalMessageInfo

func (m *MplsLmBfdNbrsLinkInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type MplsLmBfdNbrsLinkInfo struct {
	// Neighbors of the specified link id
	Neighbors            []*MplsLmBfdNbrInfo `protobuf:"bytes,50,rep,name=neighbors" json:"neighbors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MplsLmBfdNbrsLinkInfo) Reset()         { *m = MplsLmBfdNbrsLinkInfo{} }
func (m *MplsLmBfdNbrsLinkInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmBfdNbrsLinkInfo) ProtoMessage()    {}
func (*MplsLmBfdNbrsLinkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_lm_bfd_nbrs_link_info_2ce6ef9822fd4219, []int{1}
}
func (m *MplsLmBfdNbrsLinkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo.Unmarshal(m, b)
}
func (m *MplsLmBfdNbrsLinkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo.Marshal(b, m, deterministic)
}
func (dst *MplsLmBfdNbrsLinkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmBfdNbrsLinkInfo.Merge(dst, src)
}
func (m *MplsLmBfdNbrsLinkInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo.Size(m)
}
func (m *MplsLmBfdNbrsLinkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmBfdNbrsLinkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmBfdNbrsLinkInfo proto.InternalMessageInfo

func (m *MplsLmBfdNbrsLinkInfo) GetNeighbors() []*MplsLmBfdNbrInfo {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

// BFD Neighbor information
type MplsLmBfdNbrInfo struct {
	// The neighbor's IP address
	NeighborAddress string `protobuf:"bytes,1,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	// TRUE if the BFD session is Up on this link
	IsBfdUp              bool     `protobuf:"varint,2,opt,name=is_bfd_up,json=isBfdUp" json:"is_bfd_up,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmBfdNbrInfo) Reset()         { *m = MplsLmBfdNbrInfo{} }
func (m *MplsLmBfdNbrInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmBfdNbrInfo) ProtoMessage()    {}
func (*MplsLmBfdNbrInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_lm_bfd_nbrs_link_info_2ce6ef9822fd4219, []int{2}
}
func (m *MplsLmBfdNbrInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmBfdNbrInfo.Unmarshal(m, b)
}
func (m *MplsLmBfdNbrInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmBfdNbrInfo.Marshal(b, m, deterministic)
}
func (dst *MplsLmBfdNbrInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmBfdNbrInfo.Merge(dst, src)
}
func (m *MplsLmBfdNbrInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmBfdNbrInfo.Size(m)
}
func (m *MplsLmBfdNbrInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmBfdNbrInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmBfdNbrInfo proto.InternalMessageInfo

func (m *MplsLmBfdNbrInfo) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *MplsLmBfdNbrInfo) GetIsBfdUp() bool {
	if m != nil {
		return m.IsBfdUp
	}
	return false
}

func init() {
	proto.RegisterType((*MplsLmBfdNbrsLinkInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.links.link.mpls_lm_bfd_nbrs_link_info_KEYS")
	proto.RegisterType((*MplsLmBfdNbrsLinkInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.links.link.mpls_lm_bfd_nbrs_link_info")
	proto.RegisterType((*MplsLmBfdNbrInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.links.link.mpls_lm_bfd_nbr_info")
}

func init() {
	proto.RegisterFile("mpls_lm_bfd_nbrs_link_info.proto", fileDescriptor_mpls_lm_bfd_nbrs_link_info_2ce6ef9822fd4219)
}

var fileDescriptor_mpls_lm_bfd_nbrs_link_info_2ce6ef9822fd4219 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x3f, 0x4b, 0xc5, 0x30,
	0x14, 0xc5, 0xa9, 0x82, 0xda, 0x88, 0x7f, 0x08, 0x0e, 0xe5, 0x2d, 0x96, 0x82, 0x50, 0x97, 0x0c,
	0xcf, 0xc5, 0x55, 0x41, 0x10, 0x84, 0x37, 0x54, 0x1c, 0x1c, 0xe4, 0x92, 0x34, 0x89, 0x5e, 0x6c,
	0x93, 0x90, 0x1b, 0x41, 0xbf, 0x82, 0x9f, 0x5a, 0x8c, 0x56, 0x41, 0xec, 0xf2, 0x96, 0x40, 0x0e,
	0xe7, 0x9c, 0xdf, 0xbd, 0x97, 0xd5, 0x63, 0x18, 0x08, 0x86, 0x11, 0x94, 0xd5, 0xe0, 0x54, 0x24,
	0x18, 0xd0, 0x3d, 0x03, 0x3a, 0xeb, 0x45, 0x88, 0x3e, 0x79, 0x7e, 0xde, 0x23, 0xf5, 0x1e, 0xd0,
	0x13, 0xbc, 0x46, 0xc8, 0xf6, 0x64, 0xc0, 0x07, 0x13, 0xc5, 0xf4, 0xa1, 0x24, 0x9d, 0x56, 0x6f,
	0x42, 0x59, 0x2d, 0x3e, 0xe3, 0x94, 0xdf, 0xe6, 0x9a, 0x1d, 0xcf, 0xb7, 0xc3, 0xcd, 0xd5, 0xfd,
	0x2d, 0x3f, 0x61, 0xfb, 0xe8, 0x92, 0x89, 0x56, 0xf6, 0x06, 0x9c, 0x1c, 0x4d, 0x55, 0xd4, 0x45,
	0x5b, 0x76, 0x7b, 0x3f, 0xea, 0x4a, 0x8e, 0xa6, 0x79, 0x2f, 0xd8, 0x62, 0xbe, 0x8a, 0x0f, 0xac,
	0x74, 0x06, 0x1f, 0x9f, 0x94, 0x8f, 0x54, 0x2d, 0xeb, 0xcd, 0x76, 0x77, 0xb9, 0x12, 0xeb, 0x8e,
	0x2d, 0xfe, 0x80, 0x32, 0xa2, 0xfb, 0x05, 0x34, 0x0f, 0xec, 0xe8, 0x3f, 0x0b, 0x3f, 0x65, 0x87,
	0x93, 0x09, 0xa4, 0xd6, 0xd1, 0x10, 0x7d, 0x6f, 0x73, 0x30, 0xe9, 0x17, 0x5f, 0x32, 0x5f, 0xb0,
	0x12, 0x29, 0xa7, 0x5f, 0x42, 0xb5, 0x51, 0x17, 0xed, 0x4e, 0xb7, 0x8d, 0x74, 0x69, 0xf5, 0x5d,
	0x50, 0x5b, 0xf9, 0xec, 0x67, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x9d, 0x09, 0x76, 0x9a,
	0x01, 0x00, 0x00,
}
