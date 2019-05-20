// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_mgmt_count_prefix_interface_info.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_counters_prefix_filtering_interfaces_interfaces_interface

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

// Prefix filtering interface counters
type RsvpMgmtCountPrefixInterfaceInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtCountPrefixInterfaceInfo_KEYS) Reset()         { *m = RsvpMgmtCountPrefixInterfaceInfo_KEYS{} }
func (m *RsvpMgmtCountPrefixInterfaceInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtCountPrefixInterfaceInfo_KEYS) ProtoMessage()    {}
func (*RsvpMgmtCountPrefixInterfaceInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_count_prefix_interface_info_929979ccd364172a, []int{0}
}
func (m *RsvpMgmtCountPrefixInterfaceInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtCountPrefixInterfaceInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtCountPrefixInterfaceInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo_KEYS.Merge(dst, src)
}
func (m *RsvpMgmtCountPrefixInterfaceInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo_KEYS.Size(m)
}
func (m *RsvpMgmtCountPrefixInterfaceInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtCountPrefixInterfaceInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type RsvpMgmtCountPrefixInterfaceInfo struct {
	// Count of messages which got forwarded
	Forwarded *RsvpMgmtCountPrefixMsg `protobuf:"bytes,50,opt,name=forwarded" json:"forwarded,omitempty"`
	// Count of locally destined messages
	LocallyDestined *RsvpMgmtCountPrefixMsg `protobuf:"bytes,51,opt,name=locally_destined,json=locallyDestined" json:"locally_destined,omitempty"`
	// Count of messages dropped
	Dropped *RsvpMgmtCountPrefixMsg `protobuf:"bytes,52,opt,name=dropped" json:"dropped,omitempty"`
	// Count of messages which got dropped due to default ACL action
	DefaultActionDropped *RsvpMgmtCountPrefixMsg `protobuf:"bytes,53,opt,name=default_action_dropped,json=defaultActionDropped" json:"default_action_dropped,omitempty"`
	// Count of messages which were processed due to default ACL action
	DefaultActionProcessed *RsvpMgmtCountPrefixMsg `protobuf:"bytes,54,opt,name=default_action_processed,json=defaultActionProcessed" json:"default_action_processed,omitempty"`
	// Count of total messages
	Total                *RsvpMgmtCountPrefixMsg `protobuf:"bytes,55,opt,name=total" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RsvpMgmtCountPrefixInterfaceInfo) Reset()         { *m = RsvpMgmtCountPrefixInterfaceInfo{} }
func (m *RsvpMgmtCountPrefixInterfaceInfo) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtCountPrefixInterfaceInfo) ProtoMessage()    {}
func (*RsvpMgmtCountPrefixInterfaceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_count_prefix_interface_info_929979ccd364172a, []int{1}
}
func (m *RsvpMgmtCountPrefixInterfaceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo.Unmarshal(m, b)
}
func (m *RsvpMgmtCountPrefixInterfaceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtCountPrefixInterfaceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo.Merge(dst, src)
}
func (m *RsvpMgmtCountPrefixInterfaceInfo) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo.Size(m)
}
func (m *RsvpMgmtCountPrefixInterfaceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtCountPrefixInterfaceInfo proto.InternalMessageInfo

func (m *RsvpMgmtCountPrefixInterfaceInfo) GetForwarded() *RsvpMgmtCountPrefixMsg {
	if m != nil {
		return m.Forwarded
	}
	return nil
}

func (m *RsvpMgmtCountPrefixInterfaceInfo) GetLocallyDestined() *RsvpMgmtCountPrefixMsg {
	if m != nil {
		return m.LocallyDestined
	}
	return nil
}

func (m *RsvpMgmtCountPrefixInterfaceInfo) GetDropped() *RsvpMgmtCountPrefixMsg {
	if m != nil {
		return m.Dropped
	}
	return nil
}

func (m *RsvpMgmtCountPrefixInterfaceInfo) GetDefaultActionDropped() *RsvpMgmtCountPrefixMsg {
	if m != nil {
		return m.DefaultActionDropped
	}
	return nil
}

func (m *RsvpMgmtCountPrefixInterfaceInfo) GetDefaultActionProcessed() *RsvpMgmtCountPrefixMsg {
	if m != nil {
		return m.DefaultActionProcessed
	}
	return nil
}

func (m *RsvpMgmtCountPrefixInterfaceInfo) GetTotal() *RsvpMgmtCountPrefixMsg {
	if m != nil {
		return m.Total
	}
	return nil
}

// Message types for prefix filtering
type RsvpMgmtCountPrefixMsg struct {
	// Count of Path messages
	Path uint32 `protobuf:"varint,1,opt,name=path" json:"path,omitempty"`
	// Count of PathTear messages
	PathTear uint32 `protobuf:"varint,2,opt,name=path_tear,json=pathTear" json:"path_tear,omitempty"`
	// Count of ReservationConfirm messages
	ReservationConfirm uint32 `protobuf:"varint,3,opt,name=reservation_confirm,json=reservationConfirm" json:"reservation_confirm,omitempty"`
	// Total count of messages
	Total                uint32   `protobuf:"varint,4,opt,name=total" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtCountPrefixMsg) Reset()         { *m = RsvpMgmtCountPrefixMsg{} }
func (m *RsvpMgmtCountPrefixMsg) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtCountPrefixMsg) ProtoMessage()    {}
func (*RsvpMgmtCountPrefixMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_count_prefix_interface_info_929979ccd364172a, []int{2}
}
func (m *RsvpMgmtCountPrefixMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtCountPrefixMsg.Unmarshal(m, b)
}
func (m *RsvpMgmtCountPrefixMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtCountPrefixMsg.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtCountPrefixMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtCountPrefixMsg.Merge(dst, src)
}
func (m *RsvpMgmtCountPrefixMsg) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtCountPrefixMsg.Size(m)
}
func (m *RsvpMgmtCountPrefixMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtCountPrefixMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtCountPrefixMsg proto.InternalMessageInfo

func (m *RsvpMgmtCountPrefixMsg) GetPath() uint32 {
	if m != nil {
		return m.Path
	}
	return 0
}

func (m *RsvpMgmtCountPrefixMsg) GetPathTear() uint32 {
	if m != nil {
		return m.PathTear
	}
	return 0
}

func (m *RsvpMgmtCountPrefixMsg) GetReservationConfirm() uint32 {
	if m != nil {
		return m.ReservationConfirm
	}
	return 0
}

func (m *RsvpMgmtCountPrefixMsg) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*RsvpMgmtCountPrefixInterfaceInfo_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.counters.prefix_filtering.interfaces.interfaces.interface.rsvp_mgmt_count_prefix_interface_info_KEYS")
	proto.RegisterType((*RsvpMgmtCountPrefixInterfaceInfo)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.counters.prefix_filtering.interfaces.interfaces.interface.rsvp_mgmt_count_prefix_interface_info")
	proto.RegisterType((*RsvpMgmtCountPrefixMsg)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.counters.prefix_filtering.interfaces.interfaces.interface.rsvp_mgmt_count_prefix_msg")
}

func init() {
	proto.RegisterFile("rsvp_mgmt_count_prefix_interface_info.proto", fileDescriptor_rsvp_mgmt_count_prefix_interface_info_929979ccd364172a)
}

var fileDescriptor_rsvp_mgmt_count_prefix_interface_info_929979ccd364172a = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0xd4, 0xcd, 0xaa, 0x13, 0x31,
	0x14, 0x07, 0x70, 0xa2, 0xd5, 0xda, 0x48, 0x55, 0x62, 0x29, 0x41, 0x37, 0xa5, 0x50, 0x28, 0x0a,
	0x23, 0xb4, 0x7e, 0xac, 0xc5, 0xba, 0x12, 0x44, 0x5a, 0x17, 0xba, 0x0a, 0x31, 0x39, 0x53, 0x03,
	0x33, 0x49, 0x38, 0x49, 0x6b, 0x7d, 0x05, 0x17, 0x2e, 0xc4, 0xbd, 0x7b, 0x17, 0x3e, 0x90, 0x4f,
	0x73, 0x69, 0x3a, 0xd3, 0xde, 0x7b, 0xb9, 0x85, 0x6e, 0x2e, 0xb3, 0x9a, 0x33, 0xe7, 0x9c, 0xf9,
	0xf3, 0x9b, 0x19, 0x08, 0x7d, 0x8a, 0x61, 0xed, 0x45, 0xb9, 0x2c, 0xa3, 0x50, 0x6e, 0x65, 0xa3,
	0xf0, 0x08, 0xb9, 0xd9, 0x08, 0x63, 0x23, 0x60, 0x2e, 0x15, 0x08, 0x63, 0x73, 0x97, 0x79, 0x74,
	0xd1, 0xb1, 0x4f, 0xca, 0x04, 0xe5, 0x84, 0x71, 0x41, 0x6c, 0x50, 0x18, 0x2f, 0xd2, 0xc3, 0xce,
	0x03, 0x66, 0xdb, 0x2a, 0x4b, 0x09, 0x80, 0x21, 0xab, 0x42, 0x72, 0x53, 0x44, 0x40, 0x63, 0x97,
	0xd9, 0x3e, 0x2e, 0x5c, 0x59, 0x0e, 0x17, 0xf4, 0xc9, 0x49, 0x10, 0xf1, 0xee, 0xed, 0xe7, 0x05,
	0x1b, 0xd1, 0x7b, 0x87, 0xb6, 0x95, 0x25, 0x70, 0x32, 0x20, 0xe3, 0xce, 0xbc, 0xbb, 0xef, 0xbe,
	0x97, 0x25, 0x0c, 0xff, 0xb7, 0xe9, 0xe8, 0xa4, 0x54, 0xf6, 0x8b, 0xd0, 0x4e, 0xee, 0xf0, 0x9b,
	0x44, 0x0d, 0x9a, 0x4f, 0x06, 0x64, 0x7c, 0x77, 0x12, 0xb3, 0xeb, 0x7a, 0xdb, 0xec, 0x08, 0xaa,
	0x0c, 0xcb, 0xf9, 0x81, 0xc1, 0xfe, 0x10, 0xfa, 0xa0, 0x70, 0x4a, 0x16, 0xc5, 0x77, 0xa1, 0x21,
	0x44, 0x63, 0x41, 0xf3, 0x69, 0x83, 0xb6, 0xfb, 0x95, 0x66, 0x56, 0x61, 0xd8, 0x4f, 0x42, 0xdb,
	0x1a, 0x9d, 0xf7, 0xa0, 0xf9, 0xf3, 0x06, 0x61, 0x35, 0x82, 0xfd, 0x25, 0xb4, 0xaf, 0x21, 0x97,
	0xab, 0x22, 0x0a, 0xa9, 0xa2, 0x71, 0x56, 0xd4, 0xbe, 0x17, 0x0d, 0xfa, 0x7a, 0x95, 0xe9, 0x75,
	0x22, 0xcd, 0x2a, 0xec, 0x3f, 0x42, 0xf9, 0x25, 0xac, 0x47, 0xa7, 0x20, 0x04, 0xd0, 0xfc, 0x65,
	0x83, 0xdc, 0xfe, 0x05, 0xee, 0x87, 0xda, 0xc4, 0x7e, 0x10, 0x7a, 0x2b, 0xba, 0x28, 0x0b, 0xfe,
	0xaa, 0x41, 0xdd, 0x8e, 0x30, 0xfc, 0x4d, 0xe8, 0xa3, 0xe3, 0x5b, 0x8c, 0xd1, 0x96, 0x97, 0xf1,
	0x6b, 0x3a, 0x18, 0xba, 0xf3, 0x54, 0xb3, 0xc7, 0xb4, 0xb3, 0xbd, 0x8a, 0x08, 0x12, 0xf9, 0x8d,
	0x34, 0xb8, 0xb3, 0x6d, 0x7c, 0x04, 0x89, 0xec, 0x19, 0x7d, 0x88, 0x10, 0x00, 0xd7, 0x32, 0xfd,
	0x09, 0xe5, 0x6c, 0x6e, 0xb0, 0xe4, 0x37, 0xd3, 0x1a, 0x3b, 0x37, 0x7a, 0xb3, 0x9b, 0xb0, 0x5e,
	0xfd, 0x31, 0x5a, 0x69, 0x65, 0x77, 0xf3, 0xe5, 0x76, 0x3a, 0x29, 0xa7, 0x67, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xb6, 0x49, 0x9c, 0x5c, 0x58, 0x05, 0x00, 0x00,
}