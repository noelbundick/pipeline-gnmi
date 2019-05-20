// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_bfd_counters.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_bfd_counters

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

// TE BFD Counters
type MplsTeBfdCounters_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeBfdCounters_KEYS) Reset()         { *m = MplsTeBfdCounters_KEYS{} }
func (m *MplsTeBfdCounters_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeBfdCounters_KEYS) ProtoMessage()    {}
func (*MplsTeBfdCounters_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_bfd_counters_07111d881e5cf620, []int{0}
}
func (m *MplsTeBfdCounters_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBfdCounters_KEYS.Unmarshal(m, b)
}
func (m *MplsTeBfdCounters_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBfdCounters_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTeBfdCounters_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBfdCounters_KEYS.Merge(dst, src)
}
func (m *MplsTeBfdCounters_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeBfdCounters_KEYS.Size(m)
}
func (m *MplsTeBfdCounters_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBfdCounters_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBfdCounters_KEYS proto.InternalMessageInfo

type MplsTeBfdCounters struct {
	// The timestamp when these stats are cleared relative to Jan 1, 1970
	LastClearedTimestamp uint32 `protobuf:"varint,50,opt,name=last_cleared_timestamp,json=lastClearedTimestamp" json:"last_cleared_timestamp,omitempty"`
	// BFD over Head-end LSPs Cumulative Counters
	BfdOverLspHeadCounters *MplsTeBfdLspCounters `protobuf:"bytes,51,opt,name=bfd_over_lsp_head_counters,json=bfdOverLspHeadCounters" json:"bfd_over_lsp_head_counters,omitempty"`
	// SBFD over Head-end LSPs Cumulative Counters
	SbfdOverLspHeadCounters *MplsTeBfdLspCounters `protobuf:"bytes,52,opt,name=sbfd_over_lsp_head_counters,json=sbfdOverLspHeadCounters" json:"sbfd_over_lsp_head_counters,omitempty"`
	// BFD over Tail-end LSPs Cumulative Counters
	BfdOverLspTailCounters *MplsTeBfdSessionCounters `protobuf:"bytes,53,opt,name=bfd_over_lsp_tail_counters,json=bfdOverLspTailCounters" json:"bfd_over_lsp_tail_counters,omitempty"`
	// BFD over Links Cumulative Counters
	BfDoLmCounters       *MplsTeBfdSessionCounters `protobuf:"bytes,54,opt,name=bf_do_lm_counters,json=bfDoLmCounters" json:"bf_do_lm_counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MplsTeBfdCounters) Reset()         { *m = MplsTeBfdCounters{} }
func (m *MplsTeBfdCounters) String() string { return proto.CompactTextString(m) }
func (*MplsTeBfdCounters) ProtoMessage()    {}
func (*MplsTeBfdCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_bfd_counters_07111d881e5cf620, []int{1}
}
func (m *MplsTeBfdCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBfdCounters.Unmarshal(m, b)
}
func (m *MplsTeBfdCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBfdCounters.Marshal(b, m, deterministic)
}
func (dst *MplsTeBfdCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBfdCounters.Merge(dst, src)
}
func (m *MplsTeBfdCounters) XXX_Size() int {
	return xxx_messageInfo_MplsTeBfdCounters.Size(m)
}
func (m *MplsTeBfdCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBfdCounters.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBfdCounters proto.InternalMessageInfo

func (m *MplsTeBfdCounters) GetLastClearedTimestamp() uint32 {
	if m != nil {
		return m.LastClearedTimestamp
	}
	return 0
}

func (m *MplsTeBfdCounters) GetBfdOverLspHeadCounters() *MplsTeBfdLspCounters {
	if m != nil {
		return m.BfdOverLspHeadCounters
	}
	return nil
}

func (m *MplsTeBfdCounters) GetSbfdOverLspHeadCounters() *MplsTeBfdLspCounters {
	if m != nil {
		return m.SbfdOverLspHeadCounters
	}
	return nil
}

func (m *MplsTeBfdCounters) GetBfdOverLspTailCounters() *MplsTeBfdSessionCounters {
	if m != nil {
		return m.BfdOverLspTailCounters
	}
	return nil
}

func (m *MplsTeBfdCounters) GetBfDoLmCounters() *MplsTeBfdSessionCounters {
	if m != nil {
		return m.BfDoLmCounters
	}
	return nil
}

// TE BFDOverLSP Counters
type MplsTeBfdLspCounters struct {
	// The number of BFDOverLSP session create events
	SessionCreateEvents uint32 `protobuf:"varint,1,opt,name=session_create_events,json=sessionCreateEvents" json:"session_create_events,omitempty"`
	// The number of BFDOverLSP session up events
	SessionUpEvents uint32 `protobuf:"varint,2,opt,name=session_up_events,json=sessionUpEvents" json:"session_up_events,omitempty"`
	// The number of BFDOverLSP session creation failed events
	SessionCreationFailedEvents uint32 `protobuf:"varint,3,opt,name=session_creation_failed_events,json=sessionCreationFailedEvents" json:"session_creation_failed_events,omitempty"`
	// The number of BFDOverLSP session down events
	SessionDownEvents uint32 `protobuf:"varint,4,opt,name=session_down_events,json=sessionDownEvents" json:"session_down_events,omitempty"`
	// The number of BFDOverLSP session admin down events
	SessionAdminDownEvents uint32 `protobuf:"varint,5,opt,name=session_admin_down_events,json=sessionAdminDownEvents" json:"session_admin_down_events,omitempty"`
	// The number of BFDOverLSP session gracefully delete events
	SessionGracefullyDeleteEvents uint32 `protobuf:"varint,6,opt,name=session_gracefully_delete_events,json=sessionGracefullyDeleteEvents" json:"session_gracefully_delete_events,omitempty"`
	// The number of BFDOverLSP session non gracefully delete events
	SessionNonGracefullyDeleteEvents uint32 `protobuf:"varint,7,opt,name=session_non_gracefully_delete_events,json=sessionNonGracefullyDeleteEvents" json:"session_non_gracefully_delete_events,omitempty"`
	// The number of BFDOverLSP session create timeout events
	SessionCreateTimeoutEvents uint32 `protobuf:"varint,8,opt,name=session_create_timeout_events,json=sessionCreateTimeoutEvents" json:"session_create_timeout_events,omitempty"`
	// The number of BFDOverLSP session replay events
	SessionReplayEvents  uint32   `protobuf:"varint,9,opt,name=session_replay_events,json=sessionReplayEvents" json:"session_replay_events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeBfdLspCounters) Reset()         { *m = MplsTeBfdLspCounters{} }
func (m *MplsTeBfdLspCounters) String() string { return proto.CompactTextString(m) }
func (*MplsTeBfdLspCounters) ProtoMessage()    {}
func (*MplsTeBfdLspCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_bfd_counters_07111d881e5cf620, []int{2}
}
func (m *MplsTeBfdLspCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBfdLspCounters.Unmarshal(m, b)
}
func (m *MplsTeBfdLspCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBfdLspCounters.Marshal(b, m, deterministic)
}
func (dst *MplsTeBfdLspCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBfdLspCounters.Merge(dst, src)
}
func (m *MplsTeBfdLspCounters) XXX_Size() int {
	return xxx_messageInfo_MplsTeBfdLspCounters.Size(m)
}
func (m *MplsTeBfdLspCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBfdLspCounters.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBfdLspCounters proto.InternalMessageInfo

func (m *MplsTeBfdLspCounters) GetSessionCreateEvents() uint32 {
	if m != nil {
		return m.SessionCreateEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionUpEvents() uint32 {
	if m != nil {
		return m.SessionUpEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionCreationFailedEvents() uint32 {
	if m != nil {
		return m.SessionCreationFailedEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionDownEvents() uint32 {
	if m != nil {
		return m.SessionDownEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionAdminDownEvents() uint32 {
	if m != nil {
		return m.SessionAdminDownEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionGracefullyDeleteEvents() uint32 {
	if m != nil {
		return m.SessionGracefullyDeleteEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionNonGracefullyDeleteEvents() uint32 {
	if m != nil {
		return m.SessionNonGracefullyDeleteEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionCreateTimeoutEvents() uint32 {
	if m != nil {
		return m.SessionCreateTimeoutEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionReplayEvents() uint32 {
	if m != nil {
		return m.SessionReplayEvents
	}
	return 0
}

// TE BFDoLM counters
type MplsTeBfdSessionCounters struct {
	// The number of BFDOverLM session create events
	SessionCreateEvents uint32 `protobuf:"varint,1,opt,name=session_create_events,json=sessionCreateEvents" json:"session_create_events,omitempty"`
	// The number of BFDOverLM session up events
	SessionUpEvents uint32 `protobuf:"varint,2,opt,name=session_up_events,json=sessionUpEvents" json:"session_up_events,omitempty"`
	// The number of BFDOverLM session creation failed events
	SessionCreationFailedEvents uint32 `protobuf:"varint,3,opt,name=session_creation_failed_events,json=sessionCreationFailedEvents" json:"session_creation_failed_events,omitempty"`
	// The number of BFDOverLM session down events
	SessionDownEvents uint32 `protobuf:"varint,4,opt,name=session_down_events,json=sessionDownEvents" json:"session_down_events,omitempty"`
	// The number of BFDOverLM session admin down events
	SessionAdminDownEvents uint32 `protobuf:"varint,5,opt,name=session_admin_down_events,json=sessionAdminDownEvents" json:"session_admin_down_events,omitempty"`
	// The number of BFDOverLM session gracefully delete events
	SessionGracefullyDeleteEvents uint32 `protobuf:"varint,6,opt,name=session_gracefully_delete_events,json=sessionGracefullyDeleteEvents" json:"session_gracefully_delete_events,omitempty"`
	// The number of BFDOverLM session non gracefully delete events
	SessionNonGracefullyDeleteEvents uint32 `protobuf:"varint,7,opt,name=session_non_gracefully_delete_events,json=sessionNonGracefullyDeleteEvents" json:"session_non_gracefully_delete_events,omitempty"`
	// The number of BFDOverLSP session replay events
	SessionReplayEvents  uint32   `protobuf:"varint,8,opt,name=session_replay_events,json=sessionReplayEvents" json:"session_replay_events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeBfdSessionCounters) Reset()         { *m = MplsTeBfdSessionCounters{} }
func (m *MplsTeBfdSessionCounters) String() string { return proto.CompactTextString(m) }
func (*MplsTeBfdSessionCounters) ProtoMessage()    {}
func (*MplsTeBfdSessionCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_bfd_counters_07111d881e5cf620, []int{3}
}
func (m *MplsTeBfdSessionCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBfdSessionCounters.Unmarshal(m, b)
}
func (m *MplsTeBfdSessionCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBfdSessionCounters.Marshal(b, m, deterministic)
}
func (dst *MplsTeBfdSessionCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBfdSessionCounters.Merge(dst, src)
}
func (m *MplsTeBfdSessionCounters) XXX_Size() int {
	return xxx_messageInfo_MplsTeBfdSessionCounters.Size(m)
}
func (m *MplsTeBfdSessionCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBfdSessionCounters.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBfdSessionCounters proto.InternalMessageInfo

func (m *MplsTeBfdSessionCounters) GetSessionCreateEvents() uint32 {
	if m != nil {
		return m.SessionCreateEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionUpEvents() uint32 {
	if m != nil {
		return m.SessionUpEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionCreationFailedEvents() uint32 {
	if m != nil {
		return m.SessionCreationFailedEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionDownEvents() uint32 {
	if m != nil {
		return m.SessionDownEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionAdminDownEvents() uint32 {
	if m != nil {
		return m.SessionAdminDownEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionGracefullyDeleteEvents() uint32 {
	if m != nil {
		return m.SessionGracefullyDeleteEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionNonGracefullyDeleteEvents() uint32 {
	if m != nil {
		return m.SessionNonGracefullyDeleteEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionReplayEvents() uint32 {
	if m != nil {
		return m.SessionReplayEvents
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeBfdCounters_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bfd.counters.mpls_te_bfd_counters_KEYS")
	proto.RegisterType((*MplsTeBfdCounters)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bfd.counters.mpls_te_bfd_counters")
	proto.RegisterType((*MplsTeBfdLspCounters)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bfd.counters.mpls_te_bfd_lsp_counters")
	proto.RegisterType((*MplsTeBfdSessionCounters)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bfd.counters.mpls_te_bfd_session_counters")
}

func init() {
	proto.RegisterFile("mpls_te_bfd_counters.proto", fileDescriptor_mpls_te_bfd_counters_07111d881e5cf620)
}

var fileDescriptor_mpls_te_bfd_counters_07111d881e5cf620 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x95, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0x59, 0x1b, 0x6b, 0x1d, 0x51, 0xe9, 0x5a, 0x63, 0x9a, 0x58, 0x09, 0xc1, 0x43, 0xf1,
	0xb0, 0x87, 0xb4, 0x0a, 0x1e, 0x4b, 0x52, 0x5b, 0x30, 0x54, 0x88, 0xf1, 0xe0, 0x69, 0x98, 0xdd,
	0x79, 0x57, 0x17, 0x66, 0x77, 0x86, 0x99, 0x49, 0x62, 0xef, 0x22, 0x7e, 0x26, 0x3f, 0x8d, 0xdf,
	0xc3, 0x8b, 0xec, 0x64, 0xde, 0xfc, 0x23, 0x1b, 0x10, 0xf5, 0x20, 0x78, 0x4b, 0x78, 0x9e, 0xe7,
	0x37, 0xcf, 0xce, 0xfb, 0xb2, 0x4b, 0x9a, 0xb9, 0x12, 0x86, 0x5a, 0xa0, 0x71, 0xca, 0x69, 0x22,
	0xc7, 0x85, 0x05, 0x6d, 0x22, 0xa5, 0xa5, 0x95, 0x61, 0x94, 0x64, 0x26, 0x91, 0x34, 0x93, 0x86,
	0x7e, 0xd2, 0x14, 0x8d, 0x52, 0x81, 0x8e, 0xfc, 0x9f, 0x28, 0x4e, 0x79, 0x84, 0xa9, 0x4e, 0x8b,
	0x1c, 0x6e, 0xa2, 0xd1, 0xd7, 0xe7, 0xef, 0xdf, 0x76, 0xbe, 0xd7, 0xc8, 0xc1, 0x26, 0x35, 0x3c,
	0x25, 0x75, 0xc1, 0x8c, 0xa5, 0x89, 0x00, 0xa6, 0x81, 0x53, 0x9b, 0xe5, 0x60, 0x2c, 0xcb, 0x55,
	0xa3, 0xdb, 0x0e, 0x8e, 0xef, 0x0e, 0x0f, 0x4a, 0xb5, 0x37, 0x13, 0x47, 0xa8, 0x85, 0x9f, 0x03,
	0xd2, 0x2c, 0x31, 0x72, 0x02, 0x9a, 0x0a, 0xa3, 0xe8, 0x47, 0x60, 0x0b, 0x68, 0xe3, 0xa4, 0x1d,
	0x1c, 0xdf, 0xe9, 0x5e, 0xfe, 0xe2, 0x13, 0x44, 0xcb, 0x05, 0x4b, 0x28, 0x0a, 0xc3, 0x7a, 0x9c,
	0xf2, 0x37, 0x13, 0xd0, 0x03, 0xa3, 0x2e, 0x81, 0xf1, 0x1e, 0x96, 0xff, 0x12, 0x90, 0x96, 0xd9,
	0xd2, 0xe3, 0xf4, 0x0f, 0xf7, 0x78, 0x64, 0x2a, 0x8a, 0x7c, 0x5d, 0xbf, 0x0f, 0xcb, 0x32, 0xb1,
	0xe8, 0xf1, 0xdc, 0xf5, 0x18, 0xfc, 0x4e, 0x0f, 0x03, 0xc6, 0x64, 0xb2, 0xd8, 0x78, 0x27, 0x23,
	0x96, 0x89, 0x79, 0x95, 0x29, 0xd9, 0x8f, 0x53, 0xca, 0x25, 0x15, 0xf9, 0xa2, 0xc0, 0x8b, 0xbf,
	0x50, 0xe0, 0x5e, 0x9c, 0xf6, 0xe5, 0x20, 0xc7, 0x83, 0x3b, 0xdf, 0x6a, 0xa4, 0x51, 0x75, 0x73,
	0x61, 0x97, 0x3c, 0x9c, 0x03, 0x34, 0x30, 0x0b, 0x14, 0x26, 0x50, 0x58, 0xd3, 0x08, 0xdc, 0x96,
	0x3d, 0xf0, 0x62, 0xcf, 0x69, 0xe7, 0x4e, 0x0a, 0x9f, 0x91, 0x7d, 0xcc, 0x8c, 0x15, 0xfa, 0x6f,
	0x38, 0xff, 0x7d, 0x2f, 0xbc, 0x53, 0xde, 0xdb, 0x23, 0x4f, 0x56, 0xf8, 0xe5, 0x8f, 0x94, 0x65,
	0x02, 0x38, 0x06, 0x77, 0x5c, 0xb0, 0xb5, 0x7c, 0x50, 0x26, 0x8b, 0x57, 0xce, 0xe3, 0x21, 0x11,
	0xc1, 0x1e, 0x94, 0xcb, 0x69, 0x81, 0xc9, 0x9a, 0x4b, 0x62, 0x97, 0xbe, 0x9c, 0x16, 0xde, 0xff,
	0x92, 0x1c, 0xa2, 0x9f, 0xf1, 0x3c, 0x5b, 0x4d, 0xdd, 0x74, 0xa9, 0xba, 0x37, 0x9c, 0x95, 0xfa,
	0x52, 0xf4, 0x82, 0xb4, 0x31, 0xfa, 0x41, 0xb3, 0x04, 0xd2, 0xb1, 0x10, 0xd7, 0x94, 0x83, 0x80,
	0xc5, 0xd5, 0xec, 0x3a, 0xc2, 0x91, 0xf7, 0x5d, 0xcc, 0x6d, 0x7d, 0xe7, 0xf2, 0xa0, 0x2b, 0xf2,
	0x14, 0x41, 0xc5, 0x36, 0xd8, 0x2d, 0x07, 0xc3, 0x43, 0xaf, 0x2a, 0x79, 0x67, 0xe4, 0x68, 0x6d,
	0x50, 0xe5, 0x1b, 0x41, 0x8e, 0x2d, 0x82, 0xf6, 0x1c, 0xa8, 0xb9, 0x32, 0xb0, 0xd1, 0xcc, 0xe2,
	0x11, 0x4b, 0xb3, 0xd6, 0xa0, 0x04, 0xbb, 0xc6, 0xe8, 0xed, 0x95, 0x59, 0x0f, 0x9d, 0x36, 0xcb,
	0x74, 0x7e, 0xec, 0x90, 0xc7, 0xdb, 0xb6, 0xed, 0xff, 0x02, 0xfd, 0x2b, 0x0b, 0x54, 0x39, 0xfd,
	0xbd, 0xca, 0xe9, 0xc7, 0xbb, 0xee, 0x8b, 0x77, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x9d,
	0x81, 0x8a, 0x0f, 0x07, 0x00, 0x00,
}
