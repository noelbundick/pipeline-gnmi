// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_mgmt_nbr_detail.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_interface_neighbor_details_interface_neighbor_detail

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

// Detail Info for RSVP Neighbor
type RsvpMgmtNbrDetail_KEYS struct {
	NeighborAddress      string   `protobuf:"bytes,1,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtNbrDetail_KEYS) Reset()         { *m = RsvpMgmtNbrDetail_KEYS{} }
func (m *RsvpMgmtNbrDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtNbrDetail_KEYS) ProtoMessage()    {}
func (*RsvpMgmtNbrDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_nbr_detail_98b57ca0f2656cc1, []int{0}
}
func (m *RsvpMgmtNbrDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtNbrDetail_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtNbrDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtNbrDetail_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtNbrDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtNbrDetail_KEYS.Merge(dst, src)
}
func (m *RsvpMgmtNbrDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtNbrDetail_KEYS.Size(m)
}
func (m *RsvpMgmtNbrDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtNbrDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtNbrDetail_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtNbrDetail_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type RsvpMgmtNbrDetail struct {
	// Neighbor node address
	NodeAddress string `protobuf:"bytes,50,opt,name=node_address,json=nodeAddress" json:"node_address,omitempty"`
	// Detail list of I/F Neighbors
	InterfaceNeighborListDetail []*RsvpMgmtIfNbrDetail `protobuf:"bytes,51,rep,name=interface_neighbor_list_detail,json=interfaceNeighborListDetail" json:"interface_neighbor_list_detail,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}               `json:"-"`
	XXX_unrecognized            []byte                 `json:"-"`
	XXX_sizecache               int32                  `json:"-"`
}

func (m *RsvpMgmtNbrDetail) Reset()         { *m = RsvpMgmtNbrDetail{} }
func (m *RsvpMgmtNbrDetail) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtNbrDetail) ProtoMessage()    {}
func (*RsvpMgmtNbrDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_nbr_detail_98b57ca0f2656cc1, []int{1}
}
func (m *RsvpMgmtNbrDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtNbrDetail.Unmarshal(m, b)
}
func (m *RsvpMgmtNbrDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtNbrDetail.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtNbrDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtNbrDetail.Merge(dst, src)
}
func (m *RsvpMgmtNbrDetail) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtNbrDetail.Size(m)
}
func (m *RsvpMgmtNbrDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtNbrDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtNbrDetail proto.InternalMessageInfo

func (m *RsvpMgmtNbrDetail) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

func (m *RsvpMgmtNbrDetail) GetInterfaceNeighborListDetail() []*RsvpMgmtIfNbrDetail {
	if m != nil {
		return m.InterfaceNeighborListDetail
	}
	return nil
}

// Detail Info for RSVP Interface Neighbor
type RsvpMgmtIfNbrDetail struct {
	// Interface Neighbor address
	InterfaceNeighborAddress string `protobuf:"bytes,1,opt,name=interface_neighbor_address,json=interfaceNeighborAddress" json:"interface_neighbor_address,omitempty"`
	// Neighbor's Interface handle
	NeighborInterfaceName string `protobuf:"bytes,2,opt,name=neighbor_interface_name,json=neighborInterfaceName" json:"neighbor_interface_name,omitempty"`
	// Is Neighbor's RR enable
	IsRrEnabled bool `protobuf:"varint,3,opt,name=is_rr_enabled,json=isRrEnabled" json:"is_rr_enabled,omitempty"`
	// Neighbor's epoch value
	NeighborEpoch uint32 `protobuf:"varint,4,opt,name=neighbor_epoch,json=neighborEpoch" json:"neighbor_epoch,omitempty"`
	// Number of out of order msgs
	OutOfOrderMessages uint32 `protobuf:"varint,5,opt,name=out_of_order_messages,json=outOfOrderMessages" json:"out_of_order_messages,omitempty"`
	// Number of retransmitted msgs
	RetransmittedMessages uint32   `protobuf:"varint,6,opt,name=retransmitted_messages,json=retransmittedMessages" json:"retransmitted_messages,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *RsvpMgmtIfNbrDetail) Reset()         { *m = RsvpMgmtIfNbrDetail{} }
func (m *RsvpMgmtIfNbrDetail) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtIfNbrDetail) ProtoMessage()    {}
func (*RsvpMgmtIfNbrDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_nbr_detail_98b57ca0f2656cc1, []int{2}
}
func (m *RsvpMgmtIfNbrDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtIfNbrDetail.Unmarshal(m, b)
}
func (m *RsvpMgmtIfNbrDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtIfNbrDetail.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtIfNbrDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtIfNbrDetail.Merge(dst, src)
}
func (m *RsvpMgmtIfNbrDetail) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtIfNbrDetail.Size(m)
}
func (m *RsvpMgmtIfNbrDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtIfNbrDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtIfNbrDetail proto.InternalMessageInfo

func (m *RsvpMgmtIfNbrDetail) GetInterfaceNeighborAddress() string {
	if m != nil {
		return m.InterfaceNeighborAddress
	}
	return ""
}

func (m *RsvpMgmtIfNbrDetail) GetNeighborInterfaceName() string {
	if m != nil {
		return m.NeighborInterfaceName
	}
	return ""
}

func (m *RsvpMgmtIfNbrDetail) GetIsRrEnabled() bool {
	if m != nil {
		return m.IsRrEnabled
	}
	return false
}

func (m *RsvpMgmtIfNbrDetail) GetNeighborEpoch() uint32 {
	if m != nil {
		return m.NeighborEpoch
	}
	return 0
}

func (m *RsvpMgmtIfNbrDetail) GetOutOfOrderMessages() uint32 {
	if m != nil {
		return m.OutOfOrderMessages
	}
	return 0
}

func (m *RsvpMgmtIfNbrDetail) GetRetransmittedMessages() uint32 {
	if m != nil {
		return m.RetransmittedMessages
	}
	return 0
}

func init() {
	proto.RegisterType((*RsvpMgmtNbrDetail_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.interface_neighbor_details.interface_neighbor_detail.rsvp_mgmt_nbr_detail_KEYS")
	proto.RegisterType((*RsvpMgmtNbrDetail)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.interface_neighbor_details.interface_neighbor_detail.rsvp_mgmt_nbr_detail")
	proto.RegisterType((*RsvpMgmtIfNbrDetail)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.interface_neighbor_details.interface_neighbor_detail.rsvp_mgmt_if_nbr_detail")
}

func init() {
	proto.RegisterFile("rsvp_mgmt_nbr_detail.proto", fileDescriptor_rsvp_mgmt_nbr_detail_98b57ca0f2656cc1)
}

var fileDescriptor_rsvp_mgmt_nbr_detail_98b57ca0f2656cc1 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0xfb, 0x7d, 0x45, 0xa7, 0x56, 0x65, 0xb0, 0x36, 0x56, 0x90, 0x58, 0x10, 0xe2,
	0x26, 0x60, 0x8b, 0xae, 0xdc, 0x08, 0x56, 0x10, 0xff, 0x14, 0xd2, 0x95, 0xab, 0x61, 0x92, 0xdc,
	0xb4, 0x03, 0x49, 0x26, 0xcc, 0x9d, 0x8a, 0xaf, 0xe5, 0x3b, 0xf8, 0x3a, 0xbe, 0x83, 0x24, 0x4d,
	0xd2, 0x16, 0xd3, 0xa5, 0xbb, 0xe1, 0x9c, 0xfb, 0x3b, 0xf7, 0x72, 0x18, 0xd2, 0x57, 0xf8, 0x9e,
	0xb2, 0x78, 0x16, 0x6b, 0x96, 0x78, 0x8a, 0x05, 0xa0, 0xb9, 0x88, 0x9c, 0x54, 0x49, 0x2d, 0xe9,
	0xd4, 0x17, 0xe8, 0x4b, 0x26, 0x24, 0xb2, 0x0f, 0xc5, 0x44, 0xca, 0xf2, 0x59, 0x99, 0x82, 0x72,
	0xb2, 0x97, 0x23, 0x12, 0x0d, 0x2a, 0xe4, 0x3e, 0xb0, 0x04, 0xc4, 0x6c, 0xee, 0xc9, 0x92, 0xc7,
	0xed, 0xd6, 0xe0, 0x81, 0x9c, 0xd4, 0xad, 0x64, 0x4f, 0xe3, 0xb7, 0x29, 0xbd, 0x24, 0x87, 0xd5,
	0x3c, 0x0f, 0x02, 0x05, 0x88, 0xa6, 0x61, 0x19, 0xf6, 0xae, 0x7b, 0x50, 0xea, 0x77, 0x4b, 0x79,
	0xf0, 0x6d, 0x90, 0xa3, 0xba, 0x20, 0x7a, 0x4e, 0xf6, 0x12, 0x19, 0x40, 0xc5, 0x0f, 0x73, 0xbe,
	0x9d, 0x69, 0x05, 0x4b, 0x3f, 0x0d, 0x72, 0x56, 0x73, 0x61, 0x24, 0x50, 0x17, 0x29, 0xe6, 0xc8,
	0x6a, 0xda, 0xed, 0x61, 0xe4, 0xfc, 0x41, 0x05, 0xce, 0xea, 0x6c, 0x11, 0xae, 0x5d, 0xee, 0x9e,
	0x56, 0xc8, 0x6b, 0x41, 0x3c, 0x0b, 0xd4, 0xf7, 0xcb, 0xde, 0xbe, 0x1a, 0xa4, 0xb7, 0x05, 0xa4,
	0xb7, 0xa4, 0x5f, 0xb3, 0x6d, 0xb3, 0x40, 0xf3, 0x57, 0x78, 0xd9, 0xc6, 0x0d, 0xe9, 0x55, 0xcc,
	0x5a, 0x0c, 0x8f, 0xc1, 0x6c, 0xe4, 0x68, 0xb7, 0xb4, 0x1f, 0xab, 0x08, 0x1e, 0x03, 0x1d, 0x90,
	0x8e, 0x40, 0xa6, 0x14, 0x83, 0x84, 0x7b, 0x11, 0x04, 0x66, 0xd3, 0x32, 0xec, 0x1d, 0xb7, 0x2d,
	0xd0, 0x55, 0xe3, 0xa5, 0x44, 0x2f, 0xc8, 0x7e, 0x95, 0x0d, 0xa9, 0xf4, 0xe7, 0xe6, 0x3f, 0xcb,
	0xb0, 0x3b, 0x6e, 0xa7, 0x54, 0xc7, 0x99, 0x48, 0xaf, 0x48, 0x57, 0x2e, 0x34, 0x93, 0x21, 0x93,
	0x2a, 0x00, 0xc5, 0x62, 0x40, 0xe4, 0x33, 0x40, 0xf3, 0x7f, 0x3e, 0x4d, 0xe5, 0x42, 0x4f, 0xc2,
	0x49, 0x66, 0xbd, 0x14, 0x0e, 0xbd, 0x26, 0xc7, 0x0a, 0xb4, 0xe2, 0x09, 0xc6, 0x42, 0x6b, 0x08,
	0x56, 0x4c, 0x2b, 0x67, 0xba, 0x1b, 0x6e, 0x89, 0x79, 0xad, 0xfc, 0x6b, 0x8f, 0x7e, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xc3, 0x6c, 0xeb, 0x6d, 0xf8, 0x02, 0x00, 0x00,
}
