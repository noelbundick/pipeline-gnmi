// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_mgmt_glbl_nbr_detail.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_global_neighbor_details_global_neighbor_detail

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

// Detail Info for RSVP Global Neighbor
type RsvpMgmtGlblNbrDetail_KEYS struct {
	NeighborAddress      string   `protobuf:"bytes,1,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtGlblNbrDetail_KEYS) Reset()         { *m = RsvpMgmtGlblNbrDetail_KEYS{} }
func (m *RsvpMgmtGlblNbrDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtGlblNbrDetail_KEYS) ProtoMessage()    {}
func (*RsvpMgmtGlblNbrDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_glbl_nbr_detail_e58546f2004883fb, []int{0}
}
func (m *RsvpMgmtGlblNbrDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtGlblNbrDetail_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtGlblNbrDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtGlblNbrDetail_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtGlblNbrDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtGlblNbrDetail_KEYS.Merge(dst, src)
}
func (m *RsvpMgmtGlblNbrDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtGlblNbrDetail_KEYS.Size(m)
}
func (m *RsvpMgmtGlblNbrDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtGlblNbrDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtGlblNbrDetail_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtGlblNbrDetail_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type RsvpMgmtGlblNbrDetail struct {
	// Is GR enabled
	IsGrEnabled bool `protobuf:"varint,50,opt,name=is_gr_enabled,json=isGrEnabled" json:"is_gr_enabled,omitempty"`
	// Neighbor node address
	NodeAddress string `protobuf:"bytes,51,opt,name=node_address,json=nodeAddress" json:"node_address,omitempty"`
	// Current restart state
	RestartState string `protobuf:"bytes,52,opt,name=restart_state,json=restartState" json:"restart_state,omitempty"`
	// Global Neighbor Flags
	GlobalNeighborFlags *RsvpMgmtGlblNbrFlags `protobuf:"bytes,53,opt,name=global_neighbor_flags,json=globalNeighborFlags" json:"global_neighbor_flags,omitempty"`
	// Local node address
	LocalNodeAddress [][]byte `protobuf:"bytes,54,rep,name=local_node_address,json=localNodeAddress,proto3" json:"local_node_address,omitempty"`
	// Interface Neighbor List
	InterfaceNeighborList [][]byte `protobuf:"bytes,55,rep,name=interface_neighbor_list,json=interfaceNeighborList,proto3" json:"interface_neighbor_list,omitempty"`
	// GR Restart Time (milliseconds)
	RestartTime uint32 `protobuf:"varint,56,opt,name=restart_time,json=restartTime" json:"restart_time,omitempty"`
	// Is GR restart timer running
	IsRestartTimerRunning bool `protobuf:"varint,57,opt,name=is_restart_timer_running,json=isRestartTimerRunning" json:"is_restart_timer_running,omitempty"`
	// How much restart time remains
	RestartTimeLeft *RsvpMgmtTimespec `protobuf:"bytes,58,opt,name=restart_time_left,json=restartTimeLeft" json:"restart_time_left,omitempty"`
	// When will restart timer expire
	RestartTimerExpiryTime *RsvpMgmtTimespec `protobuf:"bytes,59,opt,name=restart_timer_expiry_time,json=restartTimerExpiryTime" json:"restart_timer_expiry_time,omitempty"`
	// GR Recovery Time (milliseconds)
	RecoveryTime uint32 `protobuf:"varint,60,opt,name=recovery_time,json=recoveryTime" json:"recovery_time,omitempty"`
	// Is RSVP recovery timer running
	IsRecoveryTimerRunning bool `protobuf:"varint,61,opt,name=is_recovery_timer_running,json=isRecoveryTimerRunning" json:"is_recovery_timer_running,omitempty"`
	// How much recovery timer remains
	RecoveryTimeLeft *RsvpMgmtTimespec `protobuf:"bytes,62,opt,name=recovery_time_left,json=recoveryTimeLeft" json:"recovery_time_left,omitempty"`
	// Time at which recovery timer will expire
	RecoveryTimerExpTime *RsvpMgmtTimespec `protobuf:"bytes,63,opt,name=recovery_timer_exp_time,json=recoveryTimerExpTime" json:"recovery_timer_exp_time,omitempty"`
	// Hello Interval (milliseconds)
	HelloInterval uint32 `protobuf:"varint,64,opt,name=hello_interval,json=helloInterval" json:"hello_interval,omitempty"`
	// Hello missed count
	MissedHellos uint32 `protobuf:"varint,65,opt,name=missed_hellos,json=missedHellos" json:"missed_hellos,omitempty"`
	// Neighbor's hello state
	NeighborHelloState []string `protobuf:"bytes,66,rep,name=neighbor_hello_state,json=neighborHelloState" json:"neighbor_hello_state,omitempty"`
	// Hello up time
	UpTime []*RsvpMgmtTimespec `protobuf:"bytes,67,rep,name=up_time,json=upTime" json:"up_time,omitempty"`
	// Time when communication was lost
	LostCommunicationTime []*RsvpMgmtTimespec `protobuf:"bytes,68,rep,name=lost_communication_time,json=lostCommunicationTime" json:"lost_communication_time,omitempty"`
	// Reason why communication was lost
	LostCommunicationReason []string `protobuf:"bytes,69,rep,name=lost_communication_reason,json=lostCommunicationReason" json:"lost_communication_reason,omitempty"`
	// Total number of times communication got lost
	LostCommunicationTotal []uint32 `protobuf:"varint,70,rep,packed,name=lost_communication_total,json=lostCommunicationTotal" json:"lost_communication_total,omitempty"`
	// Number of pending states for this neighbor
	PendingStates        uint32   `protobuf:"varint,71,opt,name=pending_states,json=pendingStates" json:"pending_states,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtGlblNbrDetail) Reset()         { *m = RsvpMgmtGlblNbrDetail{} }
func (m *RsvpMgmtGlblNbrDetail) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtGlblNbrDetail) ProtoMessage()    {}
func (*RsvpMgmtGlblNbrDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_glbl_nbr_detail_e58546f2004883fb, []int{1}
}
func (m *RsvpMgmtGlblNbrDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtGlblNbrDetail.Unmarshal(m, b)
}
func (m *RsvpMgmtGlblNbrDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtGlblNbrDetail.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtGlblNbrDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtGlblNbrDetail.Merge(dst, src)
}
func (m *RsvpMgmtGlblNbrDetail) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtGlblNbrDetail.Size(m)
}
func (m *RsvpMgmtGlblNbrDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtGlblNbrDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtGlblNbrDetail proto.InternalMessageInfo

func (m *RsvpMgmtGlblNbrDetail) GetIsGrEnabled() bool {
	if m != nil {
		return m.IsGrEnabled
	}
	return false
}

func (m *RsvpMgmtGlblNbrDetail) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

func (m *RsvpMgmtGlblNbrDetail) GetRestartState() string {
	if m != nil {
		return m.RestartState
	}
	return ""
}

func (m *RsvpMgmtGlblNbrDetail) GetGlobalNeighborFlags() *RsvpMgmtGlblNbrFlags {
	if m != nil {
		return m.GlobalNeighborFlags
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetLocalNodeAddress() [][]byte {
	if m != nil {
		return m.LocalNodeAddress
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetInterfaceNeighborList() [][]byte {
	if m != nil {
		return m.InterfaceNeighborList
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetRestartTime() uint32 {
	if m != nil {
		return m.RestartTime
	}
	return 0
}

func (m *RsvpMgmtGlblNbrDetail) GetIsRestartTimerRunning() bool {
	if m != nil {
		return m.IsRestartTimerRunning
	}
	return false
}

func (m *RsvpMgmtGlblNbrDetail) GetRestartTimeLeft() *RsvpMgmtTimespec {
	if m != nil {
		return m.RestartTimeLeft
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetRestartTimerExpiryTime() *RsvpMgmtTimespec {
	if m != nil {
		return m.RestartTimerExpiryTime
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetRecoveryTime() uint32 {
	if m != nil {
		return m.RecoveryTime
	}
	return 0
}

func (m *RsvpMgmtGlblNbrDetail) GetIsRecoveryTimerRunning() bool {
	if m != nil {
		return m.IsRecoveryTimerRunning
	}
	return false
}

func (m *RsvpMgmtGlblNbrDetail) GetRecoveryTimeLeft() *RsvpMgmtTimespec {
	if m != nil {
		return m.RecoveryTimeLeft
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetRecoveryTimerExpTime() *RsvpMgmtTimespec {
	if m != nil {
		return m.RecoveryTimerExpTime
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetHelloInterval() uint32 {
	if m != nil {
		return m.HelloInterval
	}
	return 0
}

func (m *RsvpMgmtGlblNbrDetail) GetMissedHellos() uint32 {
	if m != nil {
		return m.MissedHellos
	}
	return 0
}

func (m *RsvpMgmtGlblNbrDetail) GetNeighborHelloState() []string {
	if m != nil {
		return m.NeighborHelloState
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetUpTime() []*RsvpMgmtTimespec {
	if m != nil {
		return m.UpTime
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetLostCommunicationTime() []*RsvpMgmtTimespec {
	if m != nil {
		return m.LostCommunicationTime
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetLostCommunicationReason() []string {
	if m != nil {
		return m.LostCommunicationReason
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetLostCommunicationTotal() []uint32 {
	if m != nil {
		return m.LostCommunicationTotal
	}
	return nil
}

func (m *RsvpMgmtGlblNbrDetail) GetPendingStates() uint32 {
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
	return fileDescriptor_rsvp_mgmt_glbl_nbr_detail_e58546f2004883fb, []int{2}
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

// Global Neighbor Flags
type RsvpMgmtGlblNbrFlags struct {
	// OUNI Application
	IsApplicationOuni bool `protobuf:"varint,1,opt,name=is_application_ouni,json=isApplicationOuni" json:"is_application_ouni,omitempty"`
	// MPLS Application
	IsApplicationMpls    bool     `protobuf:"varint,2,opt,name=is_application_mpls,json=isApplicationMpls" json:"is_application_mpls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtGlblNbrFlags) Reset()         { *m = RsvpMgmtGlblNbrFlags{} }
func (m *RsvpMgmtGlblNbrFlags) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtGlblNbrFlags) ProtoMessage()    {}
func (*RsvpMgmtGlblNbrFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_glbl_nbr_detail_e58546f2004883fb, []int{3}
}
func (m *RsvpMgmtGlblNbrFlags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtGlblNbrFlags.Unmarshal(m, b)
}
func (m *RsvpMgmtGlblNbrFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtGlblNbrFlags.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtGlblNbrFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtGlblNbrFlags.Merge(dst, src)
}
func (m *RsvpMgmtGlblNbrFlags) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtGlblNbrFlags.Size(m)
}
func (m *RsvpMgmtGlblNbrFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtGlblNbrFlags.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtGlblNbrFlags proto.InternalMessageInfo

func (m *RsvpMgmtGlblNbrFlags) GetIsApplicationOuni() bool {
	if m != nil {
		return m.IsApplicationOuni
	}
	return false
}

func (m *RsvpMgmtGlblNbrFlags) GetIsApplicationMpls() bool {
	if m != nil {
		return m.IsApplicationMpls
	}
	return false
}

func init() {
	proto.RegisterType((*RsvpMgmtGlblNbrDetail_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.global_neighbor_details.global_neighbor_detail.rsvp_mgmt_glbl_nbr_detail_KEYS")
	proto.RegisterType((*RsvpMgmtGlblNbrDetail)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.global_neighbor_details.global_neighbor_detail.rsvp_mgmt_glbl_nbr_detail")
	proto.RegisterType((*RsvpMgmtTimespec)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.global_neighbor_details.global_neighbor_detail.rsvp_mgmt_timespec")
	proto.RegisterType((*RsvpMgmtGlblNbrFlags)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.global_neighbor_details.global_neighbor_detail.rsvp_mgmt_glbl_nbr_flags")
}

func init() {
	proto.RegisterFile("rsvp_mgmt_glbl_nbr_detail.proto", fileDescriptor_rsvp_mgmt_glbl_nbr_detail_e58546f2004883fb)
}

var fileDescriptor_rsvp_mgmt_glbl_nbr_detail_e58546f2004883fb = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0x85, 0x66, 0x20, 0x1f, 0xb4, 0xbd, 0x24, 0x4c, 0x1c, 0xd3, 0x2f, 0x9b, 0xe6, 0x61, 0x80,
	0x07, 0x0c, 0xc6, 0x90, 0x6c, 0xf9, 0xda, 0xfa, 0x91, 0xa6, 0x4e, 0x5a, 0xe4, 0xa3, 0x85, 0xd2,
	0x97, 0x3e, 0x11, 0xb2, 0x44, 0x2b, 0x04, 0x28, 0x52, 0x20, 0xe9, 0x20, 0xcd, 0x7f, 0x28, 0xfa,
	0xda, 0xa2, 0xfd, 0x5d, 0xfd, 0x3d, 0x85, 0x2e, 0x2d, 0x5b, 0x49, 0xec, 0xb7, 0xd6, 0x6f, 0xf6,
	0xb9, 0xf7, 0xf2, 0x9c, 0x73, 0x0f, 0x2d, 0x19, 0xfd, 0xaa, 0xcd, 0x75, 0x46, 0xd3, 0x24, 0xb5,
	0x34, 0x11, 0x7d, 0x41, 0x65, 0x5f, 0xd3, 0x98, 0xd9, 0x90, 0x8b, 0x6e, 0xa6, 0x95, 0x55, 0xf8,
	0x3c, 0xe2, 0x26, 0x52, 0x94, 0x2b, 0x43, 0x6f, 0x34, 0xe5, 0x19, 0x85, 0x01, 0x95, 0x31, 0xdd,
	0xcd, 0x3f, 0x75, 0x13, 0xa1, 0xfa, 0xa1, 0xa0, 0x92, 0xf1, 0xe4, 0xaa, 0xaf, 0x8a, 0x61, 0x33,
	0x03, 0x6f, 0x9f, 0xa2, 0x5f, 0x66, 0x32, 0xd2, 0xd3, 0xde, 0xdb, 0x4b, 0xfc, 0x27, 0x5a, 0x1d,
	0x0f, 0x85, 0x71, 0xac, 0x99, 0x31, 0xc4, 0xf3, 0xbd, 0xce, 0x72, 0xb0, 0x52, 0xe0, 0x87, 0x0e,
	0x6e, 0x7f, 0xad, 0xa3, 0xd6, 0xcc, 0xd3, 0x70, 0x1b, 0xd5, 0xb9, 0xa1, 0x89, 0xa6, 0x4c, 0x86,
	0x7d, 0xc1, 0x62, 0xb2, 0xe5, 0x7b, 0x9d, 0xa5, 0xa0, 0xca, 0xcd, 0x89, 0xee, 0x39, 0x08, 0xff,
	0x86, 0x6a, 0x52, 0xc5, 0x6c, 0x4c, 0xb4, 0x0d, 0x44, 0xd5, 0x1c, 0x1b, 0x91, 0xe0, 0xdf, 0x51,
	0x5d, 0x33, 0x63, 0x43, 0x6d, 0xa9, 0xb1, 0xa1, 0x65, 0xe4, 0x1f, 0xe8, 0xa9, 0x8d, 0xc0, 0xcb,
	0x1c, 0xc3, 0x9f, 0x3d, 0xd4, 0xb8, 0xef, 0x78, 0x20, 0xc2, 0xc4, 0x90, 0x7f, 0x7d, 0xaf, 0x53,
	0xdd, 0x4a, 0xba, 0xdf, 0x75, 0x8d, 0xdd, 0x29, 0xae, 0x81, 0x2e, 0x58, 0x77, 0x03, 0x17, 0xa3,
	0xfe, 0xe3, 0x1c, 0xc4, 0x7f, 0x21, 0x2c, 0x54, 0x94, 0x1f, 0x53, 0xf6, 0xba, 0xe3, 0x57, 0x3a,
	0xb5, 0x60, 0x15, 0x2a, 0x17, 0x25, 0xc3, 0x3b, 0xa8, 0xc9, 0xa5, 0x65, 0x7a, 0x10, 0x46, 0x6c,
	0x42, 0x2c, 0xb8, 0xb1, 0x64, 0x17, 0x46, 0x1a, 0xe3, 0x72, 0x41, 0x73, 0xc6, 0x8d, 0xcd, 0x77,
	0x59, 0x2c, 0xca, 0xf2, 0x94, 0x91, 0x3d, 0xdf, 0xeb, 0xd4, 0x83, 0xea, 0x08, 0x7b, 0xc3, 0x53,
	0x86, 0x77, 0x11, 0xe1, 0x86, 0x96, 0xbb, 0x34, 0xd5, 0x43, 0x29, 0xb9, 0x4c, 0xc8, 0x3e, 0xa4,
	0xd3, 0xe0, 0x26, 0x98, 0x0c, 0xe8, 0xc0, 0x15, 0xf1, 0x7b, 0x0f, 0xad, 0x95, 0xc7, 0xa8, 0x60,
	0x03, 0x4b, 0x0e, 0x60, 0xb7, 0xe1, 0x0f, 0xdb, 0x6d, 0xce, 0x64, 0x32, 0x16, 0x05, 0x2b, 0x25,
	0x13, 0x67, 0x6c, 0x60, 0xf1, 0x17, 0x0f, 0xb5, 0xee, 0xda, 0x60, 0x37, 0x19, 0xd7, 0xef, 0x9c,
	0xf3, 0xff, 0xe6, 0xa5, 0x6b, 0xb3, 0xa4, 0x4b, 0xf7, 0x40, 0x01, 0xec, 0x19, 0xee, 0x6c, 0xa4,
	0xae, 0x59, 0xa1, 0xe8, 0x7f, 0xc8, 0xa2, 0x56, 0x80, 0xd0, 0xb4, 0x8f, 0x5a, 0x10, 0x46, 0xa9,
	0x6f, 0x92, 0xc6, 0x23, 0x48, 0x63, 0x33, 0x4f, 0x63, 0x32, 0x32, 0x8e, 0xe3, 0x83, 0x87, 0xf0,
	0x9d, 0x41, 0x97, 0xc7, 0xe3, 0x79, 0xf9, 0x5e, 0x2d, 0x1b, 0x81, 0x40, 0x3e, 0x7a, 0xa8, 0x79,
	0xcf, 0x0a, 0xbb, 0xc9, 0x9c, 0xf9, 0x27, 0xf3, 0x92, 0xb5, 0x51, 0x96, 0x95, 0xe7, 0x01, 0x7b,
	0xfe, 0x03, 0xfd, 0x7c, 0xc5, 0x84, 0x50, 0x14, 0x7e, 0x36, 0xd7, 0xa1, 0x20, 0x4f, 0x21, 0x8d,
	0x3a, 0xa0, 0x2f, 0x47, 0x60, 0x9e, 0x59, 0xca, 0x8d, 0x61, 0x31, 0x05, 0xdc, 0x90, 0x43, 0x97,
	0x99, 0x03, 0x5f, 0x00, 0x86, 0xff, 0x46, 0x1b, 0x63, 0x1d, 0xee, 0x50, 0xf7, 0x4c, 0x7a, 0xe6,
	0x57, 0x3a, 0xcb, 0x01, 0x2e, 0x6a, 0xd0, 0xed, 0x9e, 0x4c, 0xb7, 0x68, 0x71, 0x38, 0xda, 0xc3,
	0x91, 0x5f, 0x99, 0xcf, 0x1e, 0x16, 0x86, 0xce, 0xf9, 0x27, 0x0f, 0x35, 0x85, 0x32, 0x96, 0x46,
	0x2a, 0x4d, 0x87, 0x92, 0x47, 0xa1, 0xe5, 0x4a, 0x3a, 0x31, 0xcf, 0xe7, 0x25, 0xa6, 0x91, 0x2b,
	0x38, 0x2a, 0x0b, 0x00, 0x6d, 0x07, 0xa8, 0x35, 0x45, 0x9a, 0x66, 0xa1, 0x51, 0x92, 0xf4, 0x60,
	0x9d, 0xcd, 0x07, 0x93, 0x01, 0x94, 0xf1, 0x1e, 0x22, 0xd3, 0x6c, 0x29, 0x1b, 0x0a, 0x72, 0xec,
	0x57, 0x3a, 0xf5, 0x60, 0xf3, 0x21, 0x69, 0x5e, 0xcd, 0xef, 0x42, 0xc6, 0x64, 0xcc, 0x65, 0xe2,
	0x82, 0x33, 0xe4, 0xc4, 0xdd, 0x85, 0x11, 0x0a, 0x99, 0x99, 0xf6, 0x6b, 0x84, 0x1f, 0x3a, 0xc1,
	0x04, 0x2d, 0x1a, 0x16, 0x29, 0x19, 0xbb, 0x17, 0xe2, 0x5a, 0x50, 0x7c, 0xc5, 0x3e, 0xaa, 0xca,
	0x50, 0xaa, 0xa2, 0xfa, 0x13, 0x54, 0xcb, 0x50, 0xfb, 0x16, 0x91, 0x59, 0xef, 0x0c, 0xdc, 0x45,
	0xeb, 0xdc, 0xd0, 0x30, 0xcb, 0x44, 0x61, 0x45, 0x0d, 0x25, 0x07, 0x8e, 0xa5, 0x60, 0x8d, 0x9b,
	0xc3, 0x49, 0xe5, 0xd5, 0x50, 0xf2, 0x29, 0xfd, 0x69, 0x26, 0x1c, 0xeb, 0xfd, 0xfe, 0xf3, 0x4c,
	0x98, 0xfe, 0x02, 0xfc, 0x93, 0xd8, 0xfe, 0x16, 0x00, 0x00, 0xff, 0xff, 0x89, 0x4a, 0x10, 0x0e,
	0x6c, 0x08, 0x00, 0x00,
}
