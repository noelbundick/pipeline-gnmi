// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_sync_count_info.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_counters_nsr

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

// RSVP sync sent and received counters information
type RsvpSyncCountInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpSyncCountInfo_KEYS) Reset()         { *m = RsvpSyncCountInfo_KEYS{} }
func (m *RsvpSyncCountInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpSyncCountInfo_KEYS) ProtoMessage()    {}
func (*RsvpSyncCountInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_sync_count_info_3f5f88f30750e455, []int{0}
}
func (m *RsvpSyncCountInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpSyncCountInfo_KEYS.Unmarshal(m, b)
}
func (m *RsvpSyncCountInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpSyncCountInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpSyncCountInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpSyncCountInfo_KEYS.Merge(dst, src)
}
func (m *RsvpSyncCountInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpSyncCountInfo_KEYS.Size(m)
}
func (m *RsvpSyncCountInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpSyncCountInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpSyncCountInfo_KEYS proto.InternalMessageInfo

type RsvpSyncCountInfo struct {
	// The timestamp, in seconds, when these statistics are cleared since 00:00:00 UTC, January 1, 1970
	LastClearedTimestamp uint32 `protobuf:"varint,50,opt,name=last_cleared_timestamp,json=lastClearedTimestamp" json:"last_cleared_timestamp,omitempty"`
	// Process role
	RsvpProcessRole string `protobuf:"bytes,51,opt,name=rsvp_process_role,json=rsvpProcessRole" json:"rsvp_process_role,omitempty"`
	// Last IDT number of states
	LastIdtStates uint32 `protobuf:"varint,52,opt,name=last_idt_states,json=lastIdtStates" json:"last_idt_states,omitempty"`
	// Total number of states
	TotalStates uint32 `protobuf:"varint,53,opt,name=total_states,json=totalStates" json:"total_states,omitempty"`
	// Total number of deletions
	TotalDeletions uint32 `protobuf:"varint,54,opt,name=total_deletions,json=totalDeletions" json:"total_deletions,omitempty"`
	// Total number of NACKs
	TotalNacks uint64 `protobuf:"varint,55,opt,name=total_nacks,json=totalNacks" json:"total_nacks,omitempty"`
	// Total number of IDTs
	TotalIdTs            uint32   `protobuf:"varint,56,opt,name=total_id_ts,json=totalIdTs" json:"total_id_ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpSyncCountInfo) Reset()         { *m = RsvpSyncCountInfo{} }
func (m *RsvpSyncCountInfo) String() string { return proto.CompactTextString(m) }
func (*RsvpSyncCountInfo) ProtoMessage()    {}
func (*RsvpSyncCountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_sync_count_info_3f5f88f30750e455, []int{1}
}
func (m *RsvpSyncCountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpSyncCountInfo.Unmarshal(m, b)
}
func (m *RsvpSyncCountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpSyncCountInfo.Marshal(b, m, deterministic)
}
func (dst *RsvpSyncCountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpSyncCountInfo.Merge(dst, src)
}
func (m *RsvpSyncCountInfo) XXX_Size() int {
	return xxx_messageInfo_RsvpSyncCountInfo.Size(m)
}
func (m *RsvpSyncCountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpSyncCountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpSyncCountInfo proto.InternalMessageInfo

func (m *RsvpSyncCountInfo) GetLastClearedTimestamp() uint32 {
	if m != nil {
		return m.LastClearedTimestamp
	}
	return 0
}

func (m *RsvpSyncCountInfo) GetRsvpProcessRole() string {
	if m != nil {
		return m.RsvpProcessRole
	}
	return ""
}

func (m *RsvpSyncCountInfo) GetLastIdtStates() uint32 {
	if m != nil {
		return m.LastIdtStates
	}
	return 0
}

func (m *RsvpSyncCountInfo) GetTotalStates() uint32 {
	if m != nil {
		return m.TotalStates
	}
	return 0
}

func (m *RsvpSyncCountInfo) GetTotalDeletions() uint32 {
	if m != nil {
		return m.TotalDeletions
	}
	return 0
}

func (m *RsvpSyncCountInfo) GetTotalNacks() uint64 {
	if m != nil {
		return m.TotalNacks
	}
	return 0
}

func (m *RsvpSyncCountInfo) GetTotalIdTs() uint32 {
	if m != nil {
		return m.TotalIdTs
	}
	return 0
}

func init() {
	proto.RegisterType((*RsvpSyncCountInfo_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.counters.nsr.rsvp_sync_count_info_KEYS")
	proto.RegisterType((*RsvpSyncCountInfo)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.counters.nsr.rsvp_sync_count_info")
}

func init() {
	proto.RegisterFile("rsvp_sync_count_info.proto", fileDescriptor_rsvp_sync_count_info_3f5f88f30750e455)
}

var fileDescriptor_rsvp_sync_count_info_3f5f88f30750e455 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0x47, 0xa9, 0x88, 0xd0, 0xd5, 0x5a, 0x5c, 0x8a, 0x44, 0x05, 0xad, 0x3d, 0x68, 0x50, 0xc8,
	0xc1, 0xd6, 0x3f, 0x77, 0xf5, 0x50, 0x04, 0x91, 0xb4, 0x17, 0x4f, 0x43, 0xdc, 0xac, 0xb0, 0xb8,
	0xdd, 0x59, 0x76, 0x46, 0xd1, 0x0f, 0xe5, 0x77, 0x94, 0x4c, 0xcc, 0xad, 0xb7, 0xf0, 0xde, 0x9b,
	0x5f, 0x0e, 0xab, 0x0e, 0x13, 0x7d, 0x45, 0xa0, 0x9f, 0x60, 0xc0, 0xe0, 0x67, 0x60, 0x70, 0xe1,
	0x1d, 0x8b, 0x98, 0x90, 0x51, 0x5f, 0x1a, 0x47, 0x06, 0xc1, 0x21, 0xc1, 0x77, 0x02, 0x17, 0x41,
	0x5a, 0x8c, 0x36, 0x15, 0xcd, 0x57, 0x21, 0x07, 0x36, 0x51, 0x11, 0x28, 0x4d, 0x8e, 0xd4, 0xc1,
	0xba, 0x29, 0x78, 0x7a, 0x7c, 0x5d, 0x4c, 0x7e, 0x37, 0xd4, 0x68, 0x9d, 0xd5, 0x33, 0xb5, 0xef,
	0x2b, 0x62, 0x30, 0xde, 0x56, 0xc9, 0xd6, 0xc0, 0x6e, 0x65, 0x89, 0xab, 0x55, 0xcc, 0xae, 0xc6,
	0xbd, 0x7c, 0x50, 0x8e, 0x1a, 0x7b, 0xdf, 0xca, 0x65, 0xe7, 0xf4, 0x85, 0xda, 0x93, 0xb5, 0x98,
	0xd0, 0x58, 0x22, 0x48, 0xe8, 0x6d, 0x36, 0x1d, 0xf7, 0xf2, 0x7e, 0x39, 0x6c, 0xc4, 0x4b, 0xcb,
	0x4b, 0xf4, 0x56, 0x9f, 0xa9, 0xa1, 0xfc, 0xc1, 0xd5, 0x0c, 0xc4, 0x15, 0x5b, 0xca, 0x66, 0x32,
	0x3d, 0x68, 0xf0, 0xbc, 0xe6, 0x85, 0x40, 0x7d, 0xaa, 0x76, 0x18, 0xb9, 0xf2, 0x5d, 0x74, 0x2d,
	0xd1, 0xb6, 0xb0, 0xff, 0xe4, 0x5c, 0x0d, 0xdb, 0xa4, 0xb6, 0xde, 0xb2, 0xc3, 0x40, 0xd9, 0x8d,
	0x54, 0xbb, 0x82, 0x1f, 0x3a, 0xaa, 0x4f, 0x54, 0x7b, 0x07, 0xa1, 0x32, 0x1f, 0x94, 0xdd, 0x8e,
	0x7b, 0xf9, 0x66, 0xa9, 0x04, 0x3d, 0x37, 0x44, 0x1f, 0x77, 0x81, 0xab, 0x81, 0x29, 0xbb, 0x93,
	0x95, 0xbe, 0xa0, 0x79, 0xbd, 0xa4, 0xb7, 0x2d, 0x79, 0x80, 0xe9, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x37, 0x37, 0xfc, 0x8e, 0x9e, 0x01, 0x00, 0x00,
}