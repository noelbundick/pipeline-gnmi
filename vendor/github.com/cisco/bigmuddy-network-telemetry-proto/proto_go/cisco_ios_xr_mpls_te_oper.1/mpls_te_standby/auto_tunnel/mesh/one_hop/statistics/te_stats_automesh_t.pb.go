// Code generated by protoc-gen-go. DO NOT EDIT.
// source: te_stats_automesh_t.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_auto_tunnel_mesh_one_hop_statistics

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

// Counters for TE Automesh
type TeStatsAutomeshT_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsAutomeshT_KEYS) Reset()         { *m = TeStatsAutomeshT_KEYS{} }
func (m *TeStatsAutomeshT_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeStatsAutomeshT_KEYS) ProtoMessage()    {}
func (*TeStatsAutomeshT_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_automesh_t_83c0b64a9ff38c2c, []int{0}
}
func (m *TeStatsAutomeshT_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsAutomeshT_KEYS.Unmarshal(m, b)
}
func (m *TeStatsAutomeshT_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsAutomeshT_KEYS.Marshal(b, m, deterministic)
}
func (dst *TeStatsAutomeshT_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsAutomeshT_KEYS.Merge(dst, src)
}
func (m *TeStatsAutomeshT_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeStatsAutomeshT_KEYS.Size(m)
}
func (m *TeStatsAutomeshT_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsAutomeshT_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsAutomeshT_KEYS proto.InternalMessageInfo

type TeStatsAutomeshT struct {
	// Number of connected automesh tunnels
	Created uint32 `protobuf:"varint,50,opt,name=created" json:"created,omitempty"`
	// Number of connected automesh tunnels
	Connected uint32 `protobuf:"varint,51,opt,name=connected" json:"connected,omitempty"`
	// Number of automesh tunnels removed while unused
	RemovedUnused uint32 `protobuf:"varint,52,opt,name=removed_unused,json=removedUnused" json:"removed_unused,omitempty"`
	// Number of automesh tunnels removed while up and in use
	RemovedInUse uint32 `protobuf:"varint,53,opt,name=removed_in_use,json=removedInUse" json:"removed_in_use,omitempty"`
	// Number of automesh tunnel attempts rejected because the total number exceeds the range
	RemovedRangeExceeded uint32 `protobuf:"varint,54,opt,name=removed_range_exceeded,json=removedRangeExceeded" json:"removed_range_exceeded,omitempty"`
	// Time at which these were last cleared in seconds since (in seconds since 1st Jan 1970 00:00:00)'
	LastClearedTime      uint32   `protobuf:"varint,55,opt,name=last_cleared_time,json=lastClearedTime" json:"last_cleared_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsAutomeshT) Reset()         { *m = TeStatsAutomeshT{} }
func (m *TeStatsAutomeshT) String() string { return proto.CompactTextString(m) }
func (*TeStatsAutomeshT) ProtoMessage()    {}
func (*TeStatsAutomeshT) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_automesh_t_83c0b64a9ff38c2c, []int{1}
}
func (m *TeStatsAutomeshT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsAutomeshT.Unmarshal(m, b)
}
func (m *TeStatsAutomeshT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsAutomeshT.Marshal(b, m, deterministic)
}
func (dst *TeStatsAutomeshT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsAutomeshT.Merge(dst, src)
}
func (m *TeStatsAutomeshT) XXX_Size() int {
	return xxx_messageInfo_TeStatsAutomeshT.Size(m)
}
func (m *TeStatsAutomeshT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsAutomeshT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsAutomeshT proto.InternalMessageInfo

func (m *TeStatsAutomeshT) GetCreated() uint32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *TeStatsAutomeshT) GetConnected() uint32 {
	if m != nil {
		return m.Connected
	}
	return 0
}

func (m *TeStatsAutomeshT) GetRemovedUnused() uint32 {
	if m != nil {
		return m.RemovedUnused
	}
	return 0
}

func (m *TeStatsAutomeshT) GetRemovedInUse() uint32 {
	if m != nil {
		return m.RemovedInUse
	}
	return 0
}

func (m *TeStatsAutomeshT) GetRemovedRangeExceeded() uint32 {
	if m != nil {
		return m.RemovedRangeExceeded
	}
	return 0
}

func (m *TeStatsAutomeshT) GetLastClearedTime() uint32 {
	if m != nil {
		return m.LastClearedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*TeStatsAutomeshT_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.auto_tunnel.mesh.one_hop.statistics.te_stats_automesh_t_KEYS")
	proto.RegisterType((*TeStatsAutomeshT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.auto_tunnel.mesh.one_hop.statistics.te_stats_automesh_t")
}

func init() {
	proto.RegisterFile("te_stats_automesh_t.proto", fileDescriptor_te_stats_automesh_t_83c0b64a9ff38c2c)
}

var fileDescriptor_te_stats_automesh_t_83c0b64a9ff38c2c = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x40, 0xe9, 0xe5, 0xfb, 0x30, 0x58, 0xc5, 0x28, 0x12, 0xc5, 0x83, 0x14, 0x05, 0xf1, 0x90,
	0x83, 0xad, 0xfa, 0x03, 0xa4, 0x07, 0x11, 0x2f, 0xd5, 0x1e, 0x3c, 0x0d, 0x69, 0x32, 0xd8, 0xc0,
	0x6e, 0xb2, 0x64, 0x26, 0x52, 0xff, 0xbb, 0x07, 0x49, 0xdc, 0xa5, 0x97, 0x1e, 0xe7, 0xbd, 0x37,
	0x13, 0x88, 0x38, 0x63, 0x04, 0x62, 0xc3, 0x04, 0x26, 0x73, 0x6c, 0x91, 0xd6, 0xc0, 0xba, 0x4b,
	0x91, 0xa3, 0x7c, 0xb5, 0x9e, 0x6c, 0x04, 0x1f, 0x09, 0x36, 0x09, 0xda, 0xae, 0x21, 0x60, 0x84,
	0xd8, 0x61, 0xd2, 0xc3, 0x40, 0x6c, 0x82, 0x5b, 0x7d, 0xeb, 0xb2, 0x0b, 0x9c, 0x43, 0xc0, 0x46,
	0x97, 0x1b, 0x3a, 0x06, 0x84, 0x75, 0xec, 0x74, 0x39, 0xed, 0x89, 0xbd, 0xa5, 0xc9, 0xb9, 0x50,
	0x3b, 0xde, 0x82, 0x97, 0xf9, 0xc7, 0xdb, 0xe4, 0x67, 0x24, 0x8e, 0x77, 0x48, 0xa9, 0xc4, 0x7f,
	0x9b, 0xd0, 0x30, 0x3a, 0x75, 0x77, 0x39, 0xba, 0x19, 0x2f, 0x86, 0x51, 0x5e, 0x88, 0x3d, 0x1b,
	0x43, 0x40, 0x5b, 0xdc, 0xb4, 0xba, 0x2d, 0x90, 0xd7, 0xe2, 0x20, 0x61, 0x1b, 0xbf, 0xd0, 0x41,
	0x0e, 0x99, 0xd0, 0xa9, 0x59, 0x4d, 0xc6, 0x3d, 0x5d, 0x56, 0x28, 0xaf, 0xb6, 0x99, 0x0f, 0x90,
	0x09, 0xd5, 0x7d, 0xcd, 0xf6, 0x7b, 0xfa, 0x1c, 0x96, 0x84, 0x72, 0x26, 0x4e, 0x87, 0x2a, 0x99,
	0xf0, 0x89, 0x80, 0x1b, 0x8b, 0xe8, 0xd0, 0xa9, 0x87, 0x5a, 0x9f, 0xf4, 0x76, 0x51, 0xe4, 0xbc,
	0x77, 0xf2, 0x56, 0x1c, 0x35, 0x86, 0x18, 0x6c, 0x83, 0x26, 0xa1, 0x03, 0xf6, 0x2d, 0xaa, 0xc7,
	0xba, 0x70, 0x58, 0xc4, 0xd3, 0x1f, 0x7f, 0xf7, 0x2d, 0xae, 0xfe, 0xd5, 0x0f, 0x9f, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x48, 0x6b, 0x1d, 0xdf, 0x8d, 0x01, 0x00, 0x00,
}