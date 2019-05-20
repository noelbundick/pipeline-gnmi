// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_mgmt_graceful_restart.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_graceful_restart

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

// RSVP Graceful Restart Info
type RsvpMgmtGracefulRestart_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtGracefulRestart_KEYS) Reset()         { *m = RsvpMgmtGracefulRestart_KEYS{} }
func (m *RsvpMgmtGracefulRestart_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtGracefulRestart_KEYS) ProtoMessage()    {}
func (*RsvpMgmtGracefulRestart_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_graceful_restart_a4d6e6c9ed9a3cc9, []int{0}
}
func (m *RsvpMgmtGracefulRestart_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtGracefulRestart_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtGracefulRestart_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtGracefulRestart_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtGracefulRestart_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtGracefulRestart_KEYS.Merge(dst, src)
}
func (m *RsvpMgmtGracefulRestart_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtGracefulRestart_KEYS.Size(m)
}
func (m *RsvpMgmtGracefulRestart_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtGracefulRestart_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtGracefulRestart_KEYS proto.InternalMessageInfo

type RsvpMgmtGracefulRestart struct {
	// Whether GR is enabled
	IsGrEnabled bool `protobuf:"varint,50,opt,name=is_gr_enabled,json=isGrEnabled" json:"is_gr_enabled,omitempty"`
	// Global neighbor count
	GlobalNeighbors uint32 `protobuf:"varint,51,opt,name=global_neighbors,json=globalNeighbors" json:"global_neighbors,omitempty"`
	// Local node address
	LocalNodeAddress []*RsvpMgmtLocalNodeIdIpv4 `protobuf:"bytes,52,rep,name=local_node_address,json=localNodeAddress" json:"local_node_address,omitempty"`
	// Restart time (milliseconds)
	RestartTime uint32 `protobuf:"varint,53,opt,name=restart_time,json=restartTime" json:"restart_time,omitempty"`
	// Recovery time (milliseconds)
	RecoveryTime uint32 `protobuf:"varint,54,opt,name=recovery_time,json=recoveryTime" json:"recovery_time,omitempty"`
	// Whether recovery timer is running
	IsRecoveryTimerRunning bool `protobuf:"varint,55,opt,name=is_recovery_timer_running,json=isRecoveryTimerRunning" json:"is_recovery_timer_running,omitempty"`
	// How much recovery timer remains
	RecoveryTimeLeft *RsvpMgmtTimespec `protobuf:"bytes,56,opt,name=recovery_time_left,json=recoveryTimeLeft" json:"recovery_time_left,omitempty"`
	// Time at which recovery timer will expire
	RecoveryTimerExpTime *RsvpMgmtTimespec `protobuf:"bytes,57,opt,name=recovery_timer_exp_time,json=recoveryTimerExpTime" json:"recovery_timer_exp_time,omitempty"`
	// Interval at which hello messages are sent
	HelloInterval uint32 `protobuf:"varint,58,opt,name=hello_interval,json=helloInterval" json:"hello_interval,omitempty"`
	// Max number of hellos missed before hellos declared down
	MissedHellos uint32 `protobuf:"varint,59,opt,name=missed_hellos,json=missedHellos" json:"missed_hellos,omitempty"`
	// Total number of pending states
	PendingStates        uint32   `protobuf:"varint,60,opt,name=pending_states,json=pendingStates" json:"pending_states,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtGracefulRestart) Reset()         { *m = RsvpMgmtGracefulRestart{} }
func (m *RsvpMgmtGracefulRestart) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtGracefulRestart) ProtoMessage()    {}
func (*RsvpMgmtGracefulRestart) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_graceful_restart_a4d6e6c9ed9a3cc9, []int{1}
}
func (m *RsvpMgmtGracefulRestart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtGracefulRestart.Unmarshal(m, b)
}
func (m *RsvpMgmtGracefulRestart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtGracefulRestart.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtGracefulRestart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtGracefulRestart.Merge(dst, src)
}
func (m *RsvpMgmtGracefulRestart) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtGracefulRestart.Size(m)
}
func (m *RsvpMgmtGracefulRestart) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtGracefulRestart.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtGracefulRestart proto.InternalMessageInfo

func (m *RsvpMgmtGracefulRestart) GetIsGrEnabled() bool {
	if m != nil {
		return m.IsGrEnabled
	}
	return false
}

func (m *RsvpMgmtGracefulRestart) GetGlobalNeighbors() uint32 {
	if m != nil {
		return m.GlobalNeighbors
	}
	return 0
}

func (m *RsvpMgmtGracefulRestart) GetLocalNodeAddress() []*RsvpMgmtLocalNodeIdIpv4 {
	if m != nil {
		return m.LocalNodeAddress
	}
	return nil
}

func (m *RsvpMgmtGracefulRestart) GetRestartTime() uint32 {
	if m != nil {
		return m.RestartTime
	}
	return 0
}

func (m *RsvpMgmtGracefulRestart) GetRecoveryTime() uint32 {
	if m != nil {
		return m.RecoveryTime
	}
	return 0
}

func (m *RsvpMgmtGracefulRestart) GetIsRecoveryTimerRunning() bool {
	if m != nil {
		return m.IsRecoveryTimerRunning
	}
	return false
}

func (m *RsvpMgmtGracefulRestart) GetRecoveryTimeLeft() *RsvpMgmtTimespec {
	if m != nil {
		return m.RecoveryTimeLeft
	}
	return nil
}

func (m *RsvpMgmtGracefulRestart) GetRecoveryTimerExpTime() *RsvpMgmtTimespec {
	if m != nil {
		return m.RecoveryTimerExpTime
	}
	return nil
}

func (m *RsvpMgmtGracefulRestart) GetHelloInterval() uint32 {
	if m != nil {
		return m.HelloInterval
	}
	return 0
}

func (m *RsvpMgmtGracefulRestart) GetMissedHellos() uint32 {
	if m != nil {
		return m.MissedHellos
	}
	return 0
}

func (m *RsvpMgmtGracefulRestart) GetPendingStates() uint32 {
	if m != nil {
		return m.PendingStates
	}
	return 0
}

// RSVP Time Spec
type RsvpMgmtTimespec struct {
	// Time Value in Seconds
	Seconds int32 `protobuf:"zigzag32,1,opt,name=seconds" json:"seconds,omitempty"`
	// Time Value in Nano-seconds
	Nanoseconds          int32    `protobuf:"zigzag32,2,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtTimespec) Reset()         { *m = RsvpMgmtTimespec{} }
func (m *RsvpMgmtTimespec) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtTimespec) ProtoMessage()    {}
func (*RsvpMgmtTimespec) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_graceful_restart_a4d6e6c9ed9a3cc9, []int{2}
}
func (m *RsvpMgmtTimespec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtTimespec.Unmarshal(m, b)
}
func (m *RsvpMgmtTimespec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtTimespec.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtTimespec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtTimespec.Merge(dst, src)
}
func (m *RsvpMgmtTimespec) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtTimespec.Size(m)
}
func (m *RsvpMgmtTimespec) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtTimespec.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtTimespec proto.InternalMessageInfo

func (m *RsvpMgmtTimespec) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *RsvpMgmtTimespec) GetNanoseconds() int32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

// Local Node-id
type RsvpMgmtLocalNodeIdIpv4 struct {
	// Local node address
	LocalNodeIpAddress string `protobuf:"bytes,1,opt,name=local_node_ip_address,json=localNodeIpAddress" json:"local_node_ip_address,omitempty"`
	// GR local node-id app type
	ApplicationType      string   `protobuf:"bytes,2,opt,name=application_type,json=applicationType" json:"application_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtLocalNodeIdIpv4) Reset()         { *m = RsvpMgmtLocalNodeIdIpv4{} }
func (m *RsvpMgmtLocalNodeIdIpv4) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtLocalNodeIdIpv4) ProtoMessage()    {}
func (*RsvpMgmtLocalNodeIdIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_graceful_restart_a4d6e6c9ed9a3cc9, []int{3}
}
func (m *RsvpMgmtLocalNodeIdIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtLocalNodeIdIpv4.Unmarshal(m, b)
}
func (m *RsvpMgmtLocalNodeIdIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtLocalNodeIdIpv4.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtLocalNodeIdIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtLocalNodeIdIpv4.Merge(dst, src)
}
func (m *RsvpMgmtLocalNodeIdIpv4) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtLocalNodeIdIpv4.Size(m)
}
func (m *RsvpMgmtLocalNodeIdIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtLocalNodeIdIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtLocalNodeIdIpv4 proto.InternalMessageInfo

func (m *RsvpMgmtLocalNodeIdIpv4) GetLocalNodeIpAddress() string {
	if m != nil {
		return m.LocalNodeIpAddress
	}
	return ""
}

func (m *RsvpMgmtLocalNodeIdIpv4) GetApplicationType() string {
	if m != nil {
		return m.ApplicationType
	}
	return ""
}

func init() {
	proto.RegisterType((*RsvpMgmtGracefulRestart_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.graceful_restart.rsvp_mgmt_graceful_restart_KEYS")
	proto.RegisterType((*RsvpMgmtGracefulRestart)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.graceful_restart.rsvp_mgmt_graceful_restart")
	proto.RegisterType((*RsvpMgmtTimespec)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.graceful_restart.rsvp_mgmt_timespec")
	proto.RegisterType((*RsvpMgmtLocalNodeIdIpv4)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.graceful_restart.rsvp_mgmt_local_node_id_ipv4")
}

func init() {
	proto.RegisterFile("rsvp_mgmt_graceful_restart.proto", fileDescriptor_rsvp_mgmt_graceful_restart_a4d6e6c9ed9a3cc9)
}

var fileDescriptor_rsvp_mgmt_graceful_restart_a4d6e6c9ed9a3cc9 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x5b, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x59, 0xc5, 0x4b, 0x27, 0x5d, 0xbb, 0x0e, 0x5e, 0x46, 0x11, 0x4c, 0x23, 0x42, 0xfa,
	0x12, 0x71, 0x5b, 0x2f, 0x55, 0x5f, 0x44, 0x16, 0x2d, 0x6a, 0x91, 0xb4, 0x2f, 0x3e, 0x1d, 0x66,
	0x93, 0xb3, 0xe9, 0xc0, 0x64, 0x66, 0x9c, 0x99, 0x2e, 0xbb, 0xea, 0x57, 0xf4, 0x3b, 0x49, 0x26,
	0x9b, 0x9a, 0x56, 0x2a, 0x88, 0xbe, 0x25, 0xbf, 0xf3, 0x3f, 0xb7, 0x3f, 0x67, 0x48, 0x6c, 0xdd,
	0xdc, 0x40, 0x5d, 0xd5, 0x1e, 0x2a, 0xcb, 0x0b, 0x9c, 0x1d, 0x4b, 0xb0, 0xe8, 0x3c, 0xb7, 0x3e,
	0x33, 0x56, 0x7b, 0x4d, 0x1f, 0x15, 0xc2, 0x15, 0x1a, 0x84, 0x76, 0xb0, 0xb0, 0x20, 0x0c, 0x84,
	0x0c, 0x6d, 0xd0, 0x66, 0xcd, 0x57, 0x76, 0x36, 0x2d, 0xd9, 0x24, 0xf7, 0xcf, 0x2f, 0x0a, 0xef,
	0x27, 0x9f, 0x0f, 0x92, 0x1f, 0x97, 0xc8, 0xdd, 0xf3, 0x35, 0x34, 0x21, 0x43, 0xe1, 0xa0, 0xb2,
	0x80, 0x8a, 0x4f, 0x25, 0x96, 0x6c, 0x1c, 0x0f, 0xd2, 0xab, 0x79, 0x24, 0xdc, 0x5b, 0x3b, 0x69,
	0x11, 0xdd, 0x22, 0xa3, 0x4a, 0xea, 0x29, 0x97, 0xa0, 0x50, 0x54, 0x47, 0x53, 0x6d, 0x1d, 0xdb,
	0x8e, 0x07, 0xe9, 0x30, 0xdf, 0x68, 0xf9, 0x7e, 0x87, 0xe9, 0x37, 0x42, 0xa5, 0x2e, 0x1a, 0xa5,
	0x2e, 0x11, 0x78, 0x59, 0x5a, 0x74, 0x8e, 0xed, 0xc4, 0x17, 0xd3, 0x68, 0xfc, 0x31, 0xfb, 0xcb,
	0xf5, 0xb2, 0x5f, 0x73, 0xf7, 0x8a, 0x8a, 0x12, 0x84, 0x99, 0xef, 0xe4, 0xa3, 0xc0, 0xf6, 0x75,
	0x89, 0xaf, 0xdb, 0x36, 0x74, 0x93, 0xac, 0x77, 0xab, 0x7b, 0x51, 0x23, 0x7b, 0x12, 0x66, 0x8c,
	0x56, 0xec, 0x50, 0xd4, 0x48, 0x1f, 0x90, 0xa1, 0xc5, 0x42, 0xcf, 0xd1, 0x2e, 0x5b, 0xcd, 0xd3,
	0xa0, 0x59, 0xef, 0x60, 0x10, 0xed, 0x92, 0x3b, 0xc2, 0xc1, 0x29, 0x9d, 0x05, 0x7b, 0xac, 0x94,
	0x50, 0x15, 0x7b, 0x16, 0xfc, 0xb9, 0x25, 0x5c, 0xde, 0x4b, 0xb1, 0x79, 0x1b, 0xa5, 0x5f, 0x08,
	0x3d, 0x95, 0x07, 0x12, 0x67, 0x9e, 0x3d, 0x8f, 0x07, 0x69, 0x34, 0x7e, 0xf3, 0x0f, 0xfb, 0x37,
	0xb5, 0x9c, 0xc1, 0x22, 0x1f, 0xf5, 0x27, 0xfd, 0x80, 0x33, 0x4f, 0xbf, 0x92, 0xdb, 0x67, 0x46,
	0xc5, 0x85, 0x69, 0x97, 0xdb, 0xfd, 0x7f, 0x7d, 0x6f, 0xf4, 0xfb, 0xda, 0xc9, 0xc2, 0x04, 0xa7,
	0x1e, 0x92, 0x6b, 0x47, 0x28, 0xa5, 0x06, 0xa1, 0x3c, 0xda, 0x39, 0x97, 0xec, 0x45, 0xf0, 0x73,
	0x18, 0xe8, 0xde, 0x0a, 0x36, 0xae, 0xd7, 0xc2, 0x39, 0x2c, 0x21, 0x70, 0xc7, 0x5e, 0xb6, 0xae,
	0xb7, 0xf0, 0x5d, 0x60, 0x4d, 0x2d, 0x83, 0xaa, 0x14, 0xaa, 0x02, 0xe7, 0xb9, 0x47, 0xc7, 0x5e,
	0xb5, 0xb5, 0x56, 0xf4, 0x20, 0xc0, 0xe4, 0x13, 0xa1, 0xbf, 0x8f, 0x47, 0x19, 0xb9, 0xe2, 0xb0,
	0xd0, 0xaa, 0x74, 0x6c, 0x10, 0x0f, 0xd2, 0xeb, 0x79, 0xf7, 0x4b, 0x63, 0x12, 0x29, 0xae, 0x74,
	0x17, 0xbd, 0x10, 0xa2, 0x7d, 0x94, 0x7c, 0x27, 0xf7, 0xfe, 0x74, 0x68, 0xf4, 0x31, 0xb9, 0xd9,
	0xa7, 0xe6, 0xe4, 0xac, 0x9b, 0x4e, 0x6b, 0x39, 0x3d, 0xb9, 0xc3, 0x3d, 0xd3, 0x5d, 0xe2, 0x16,
	0x19, 0x71, 0x63, 0xa4, 0x28, 0xb8, 0x17, 0x5a, 0x81, 0x5f, 0x1a, 0x0c, 0x9d, 0xd7, 0xf2, 0x8d,
	0x1e, 0x3f, 0x5c, 0x1a, 0x9c, 0x5e, 0x0e, 0x4f, 0x7f, 0xfb, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc1, 0xd1, 0x7d, 0xe3, 0x1e, 0x04, 0x00, 0x00,
}
