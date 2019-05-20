// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_tp_link_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_tp_tp_links_tp_links_tp_link

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

// Transport profile link information
type MplsTeTpLinkInfo_KEYS struct {
	TpLinkId             uint32   `protobuf:"varint,1,opt,name=tp_link_id,json=tpLinkId" json:"tp_link_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTpLinkInfo_KEYS) Reset()         { *m = MplsTeTpLinkInfo_KEYS{} }
func (m *MplsTeTpLinkInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeTpLinkInfo_KEYS) ProtoMessage()    {}
func (*MplsTeTpLinkInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_tp_link_info_a0f8c4ebf321d349, []int{0}
}
func (m *MplsTeTpLinkInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTpLinkInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTeTpLinkInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTpLinkInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTeTpLinkInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTpLinkInfo_KEYS.Merge(dst, src)
}
func (m *MplsTeTpLinkInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeTpLinkInfo_KEYS.Size(m)
}
func (m *MplsTeTpLinkInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTpLinkInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTpLinkInfo_KEYS proto.InternalMessageInfo

func (m *MplsTeTpLinkInfo_KEYS) GetTpLinkId() uint32 {
	if m != nil {
		return m.TpLinkId
	}
	return 0
}

type MplsTeTpLinkInfo struct {
	// Transport profile link identifier
	LinkId uint32 `protobuf:"varint,50,opt,name=link_id,json=linkId" json:"link_id,omitempty"`
	// Transport profile link interface name
	Interface string `protobuf:"bytes,51,opt,name=interface" json:"interface,omitempty"`
	// Transport profile next-hop in IPv4 address format
	NextHopAddress string `protobuf:"bytes,52,opt,name=next_hop_address,json=nextHopAddress" json:"next_hop_address,omitempty"`
	// Link state
	LinkState string `protobuf:"bytes,53,opt,name=link_state,json=linkState" json:"link_state,omitempty"`
	// Available bandwidth in Kbps
	AvailableBandwidth   uint64   `protobuf:"varint,54,opt,name=available_bandwidth,json=availableBandwidth" json:"available_bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTpLinkInfo) Reset()         { *m = MplsTeTpLinkInfo{} }
func (m *MplsTeTpLinkInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeTpLinkInfo) ProtoMessage()    {}
func (*MplsTeTpLinkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_tp_link_info_a0f8c4ebf321d349, []int{1}
}
func (m *MplsTeTpLinkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTpLinkInfo.Unmarshal(m, b)
}
func (m *MplsTeTpLinkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTpLinkInfo.Marshal(b, m, deterministic)
}
func (dst *MplsTeTpLinkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTpLinkInfo.Merge(dst, src)
}
func (m *MplsTeTpLinkInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeTpLinkInfo.Size(m)
}
func (m *MplsTeTpLinkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTpLinkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTpLinkInfo proto.InternalMessageInfo

func (m *MplsTeTpLinkInfo) GetLinkId() uint32 {
	if m != nil {
		return m.LinkId
	}
	return 0
}

func (m *MplsTeTpLinkInfo) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *MplsTeTpLinkInfo) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *MplsTeTpLinkInfo) GetLinkState() string {
	if m != nil {
		return m.LinkState
	}
	return ""
}

func (m *MplsTeTpLinkInfo) GetAvailableBandwidth() uint64 {
	if m != nil {
		return m.AvailableBandwidth
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeTpLinkInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_links.tp_links.tp_link.mpls_te_tp_link_info_KEYS")
	proto.RegisterType((*MplsTeTpLinkInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_links.tp_links.tp_link.mpls_te_tp_link_info")
}

func init() {
	proto.RegisterFile("mpls_te_tp_link_info.proto", fileDescriptor_mpls_te_tp_link_info_a0f8c4ebf321d349)
}

var fileDescriptor_mpls_te_tp_link_info_a0f8c4ebf321d349 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0x45, 0x09, 0x48, 0xb5, 0x0f, 0x14, 0x89, 0x82, 0x51, 0x2a, 0x0c, 0x5d, 0xcd, 0x6a, 0x04,
	0xab, 0x82, 0xb8, 0x52, 0x10, 0x14, 0x5d, 0x4d, 0x57, 0xae, 0x1e, 0x99, 0x49, 0x4a, 0x43, 0x63,
	0x12, 0x92, 0x87, 0xf6, 0x0f, 0xfd, 0x2d, 0x99, 0x8c, 0x53, 0x17, 0xed, 0x2e, 0x39, 0xf7, 0x9e,
	0x1b, 0x08, 0x5c, 0x7c, 0x06, 0x9b, 0x90, 0x34, 0x52, 0x40, 0x6b, 0xdc, 0x0a, 0x8d, 0x5b, 0xf8,
	0x2a, 0x44, 0x4f, 0x9e, 0x3f, 0xb4, 0x26, 0xb5, 0x1e, 0x8d, 0x4f, 0xb8, 0x8e, 0x38, 0x14, 0x7d,
	0xd0, 0xb1, 0xea, 0x2f, 0xa1, 0xfa, 0xb3, 0xd2, 0xd6, 0x61, 0x7a, 0x0f, 0xe7, 0xbb, 0xa6, 0xf1,
	0xed, 0xf9, 0x63, 0xce, 0x27, 0x00, 0x1b, 0xa8, 0x04, 0x2b, 0x58, 0x79, 0x58, 0x1f, 0x50, 0x78,
	0x37, 0x6e, 0xf5, 0xaa, 0xa6, 0x3f, 0x0c, 0x4e, 0x77, 0xb9, 0xfc, 0x0c, 0xf6, 0x07, 0xe7, 0x3a,
	0x3b, 0x23, 0x9b, 0x0d, 0x3e, 0x81, 0xb1, 0x71, 0xa4, 0xe3, 0x42, 0xb6, 0x5a, 0xcc, 0x0a, 0x56,
	0x8e, 0xeb, 0x7f, 0xc0, 0x4b, 0x38, 0x76, 0x7a, 0x4d, 0xb8, 0xf4, 0x01, 0xa5, 0x52, 0x51, 0xa7,
	0x24, 0x6e, 0x72, 0xe9, 0xa8, 0xe3, 0x2f, 0x3e, 0x3c, 0xf6, 0x94, 0x5f, 0x02, 0xe4, 0x07, 0x12,
	0x49, 0xd2, 0xe2, 0xb6, 0x1f, 0xea, 0xc8, 0xbc, 0x03, 0xfc, 0x0a, 0x4e, 0xe4, 0x97, 0x34, 0x56,
	0x36, 0x56, 0x63, 0x23, 0x9d, 0xfa, 0x36, 0x8a, 0x96, 0xe2, 0xae, 0x60, 0xe5, 0x5e, 0xcd, 0x37,
	0xd1, 0xd3, 0x90, 0x34, 0xa3, 0xfc, 0x91, 0xb3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x32,
	0x14, 0xe1, 0x66, 0x01, 0x00, 0x00,
}