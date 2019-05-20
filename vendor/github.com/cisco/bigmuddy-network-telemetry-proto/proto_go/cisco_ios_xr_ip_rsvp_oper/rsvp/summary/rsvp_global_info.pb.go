// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_global_info.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_summary

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

// RSVP global information
type RsvpGlobalInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpGlobalInfo_KEYS) Reset()         { *m = RsvpGlobalInfo_KEYS{} }
func (m *RsvpGlobalInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpGlobalInfo_KEYS) ProtoMessage()    {}
func (*RsvpGlobalInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_global_info_d30ce4ecb4c384b9, []int{0}
}
func (m *RsvpGlobalInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpGlobalInfo_KEYS.Unmarshal(m, b)
}
func (m *RsvpGlobalInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpGlobalInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpGlobalInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpGlobalInfo_KEYS.Merge(dst, src)
}
func (m *RsvpGlobalInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpGlobalInfo_KEYS.Size(m)
}
func (m *RsvpGlobalInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpGlobalInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpGlobalInfo_KEYS proto.InternalMessageInfo

type RsvpGlobalInfo struct {
	// Issu status
	IssuStatus *RsvpNsrIssuStatusInfo `protobuf:"bytes,50,opt,name=issu_status,json=issuStatus" json:"issu_status,omitempty"`
	// NSR status
	NsrStatus *RsvpNsrIssuStatusInfo `protobuf:"bytes,51,opt,name=nsr_status,json=nsrStatus" json:"nsr_status,omitempty"`
	// Total interfaces
	Interfaces uint32 `protobuf:"varint,52,opt,name=interfaces" json:"interfaces,omitempty"`
	// Total LSPs
	LsPs uint32 `protobuf:"varint,53,opt,name=ls_ps,json=lsPs" json:"ls_ps,omitempty"`
	// All database counters
	DatabaseCounters     *RsvpMgmtAllDbCounters `protobuf:"bytes,54,opt,name=database_counters,json=databaseCounters" json:"database_counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RsvpGlobalInfo) Reset()         { *m = RsvpGlobalInfo{} }
func (m *RsvpGlobalInfo) String() string { return proto.CompactTextString(m) }
func (*RsvpGlobalInfo) ProtoMessage()    {}
func (*RsvpGlobalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_global_info_d30ce4ecb4c384b9, []int{1}
}
func (m *RsvpGlobalInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpGlobalInfo.Unmarshal(m, b)
}
func (m *RsvpGlobalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpGlobalInfo.Marshal(b, m, deterministic)
}
func (dst *RsvpGlobalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpGlobalInfo.Merge(dst, src)
}
func (m *RsvpGlobalInfo) XXX_Size() int {
	return xxx_messageInfo_RsvpGlobalInfo.Size(m)
}
func (m *RsvpGlobalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpGlobalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpGlobalInfo proto.InternalMessageInfo

func (m *RsvpGlobalInfo) GetIssuStatus() *RsvpNsrIssuStatusInfo {
	if m != nil {
		return m.IssuStatus
	}
	return nil
}

func (m *RsvpGlobalInfo) GetNsrStatus() *RsvpNsrIssuStatusInfo {
	if m != nil {
		return m.NsrStatus
	}
	return nil
}

func (m *RsvpGlobalInfo) GetInterfaces() uint32 {
	if m != nil {
		return m.Interfaces
	}
	return 0
}

func (m *RsvpGlobalInfo) GetLsPs() uint32 {
	if m != nil {
		return m.LsPs
	}
	return 0
}

func (m *RsvpGlobalInfo) GetDatabaseCounters() *RsvpMgmtAllDbCounters {
	if m != nil {
		return m.DatabaseCounters
	}
	return nil
}

// All database counters
type RsvpMgmtAllDbCounters struct {
	// Number of sessions
	Sessions uint32 `protobuf:"varint,1,opt,name=sessions" json:"sessions,omitempty"`
	// Number of locally created and incoming path states
	IncomingPaths uint32 `protobuf:"varint,2,opt,name=incoming_paths,json=incomingPaths" json:"incoming_paths,omitempty"`
	// Number of outgoing path states
	OutgoingPaths uint32 `protobuf:"varint,3,opt,name=outgoing_paths,json=outgoingPaths" json:"outgoing_paths,omitempty"`
	// Number of locally created and incoming reservation states
	IncomingReservations uint32 `protobuf:"varint,4,opt,name=incoming_reservations,json=incomingReservations" json:"incoming_reservations,omitempty"`
	// Number of outgoing reservation states
	OutgoingReservations uint32 `protobuf:"varint,5,opt,name=outgoing_reservations,json=outgoingReservations" json:"outgoing_reservations,omitempty"`
	// Number of Interfaces
	Interfaces           uint32   `protobuf:"varint,6,opt,name=interfaces" json:"interfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtAllDbCounters) Reset()         { *m = RsvpMgmtAllDbCounters{} }
func (m *RsvpMgmtAllDbCounters) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtAllDbCounters) ProtoMessage()    {}
func (*RsvpMgmtAllDbCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_global_info_d30ce4ecb4c384b9, []int{2}
}
func (m *RsvpMgmtAllDbCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtAllDbCounters.Unmarshal(m, b)
}
func (m *RsvpMgmtAllDbCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtAllDbCounters.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtAllDbCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtAllDbCounters.Merge(dst, src)
}
func (m *RsvpMgmtAllDbCounters) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtAllDbCounters.Size(m)
}
func (m *RsvpMgmtAllDbCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtAllDbCounters.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtAllDbCounters proto.InternalMessageInfo

func (m *RsvpMgmtAllDbCounters) GetSessions() uint32 {
	if m != nil {
		return m.Sessions
	}
	return 0
}

func (m *RsvpMgmtAllDbCounters) GetIncomingPaths() uint32 {
	if m != nil {
		return m.IncomingPaths
	}
	return 0
}

func (m *RsvpMgmtAllDbCounters) GetOutgoingPaths() uint32 {
	if m != nil {
		return m.OutgoingPaths
	}
	return 0
}

func (m *RsvpMgmtAllDbCounters) GetIncomingReservations() uint32 {
	if m != nil {
		return m.IncomingReservations
	}
	return 0
}

func (m *RsvpMgmtAllDbCounters) GetOutgoingReservations() uint32 {
	if m != nil {
		return m.OutgoingReservations
	}
	return 0
}

func (m *RsvpMgmtAllDbCounters) GetInterfaces() uint32 {
	if m != nil {
		return m.Interfaces
	}
	return 0
}

// IDT status information
type RsvpIdtStatus struct {
	// Sync status
	SyncStatus string `protobuf:"bytes,1,opt,name=sync_status,json=syncStatus" json:"sync_status,omitempty"`
	// Not ready reason
	NotReadyReason string `protobuf:"bytes,2,opt,name=not_ready_reason,json=notReadyReason" json:"not_ready_reason,omitempty"`
	// IDT start timestamp in seconds
	IdtStartTime uint32 `protobuf:"varint,3,opt,name=idt_start_time,json=idtStartTime" json:"idt_start_time,omitempty"`
	// IDT end timestamp in seconds
	IdtEndTime uint32 `protobuf:"varint,4,opt,name=idt_end_time,json=idtEndTime" json:"idt_end_time,omitempty"`
	// Declare ready timestamp in seconds
	DeclareTime uint32 `protobuf:"varint,5,opt,name=declare_time,json=declareTime" json:"declare_time,omitempty"`
	// Withdraw ready timestamp in seconds
	WithdrawTime         uint32   `protobuf:"varint,6,opt,name=withdraw_time,json=withdrawTime" json:"withdraw_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpIdtStatus) Reset()         { *m = RsvpIdtStatus{} }
func (m *RsvpIdtStatus) String() string { return proto.CompactTextString(m) }
func (*RsvpIdtStatus) ProtoMessage()    {}
func (*RsvpIdtStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_global_info_d30ce4ecb4c384b9, []int{3}
}
func (m *RsvpIdtStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpIdtStatus.Unmarshal(m, b)
}
func (m *RsvpIdtStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpIdtStatus.Marshal(b, m, deterministic)
}
func (dst *RsvpIdtStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpIdtStatus.Merge(dst, src)
}
func (m *RsvpIdtStatus) XXX_Size() int {
	return xxx_messageInfo_RsvpIdtStatus.Size(m)
}
func (m *RsvpIdtStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpIdtStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpIdtStatus proto.InternalMessageInfo

func (m *RsvpIdtStatus) GetSyncStatus() string {
	if m != nil {
		return m.SyncStatus
	}
	return ""
}

func (m *RsvpIdtStatus) GetNotReadyReason() string {
	if m != nil {
		return m.NotReadyReason
	}
	return ""
}

func (m *RsvpIdtStatus) GetIdtStartTime() uint32 {
	if m != nil {
		return m.IdtStartTime
	}
	return 0
}

func (m *RsvpIdtStatus) GetIdtEndTime() uint32 {
	if m != nil {
		return m.IdtEndTime
	}
	return 0
}

func (m *RsvpIdtStatus) GetDeclareTime() uint32 {
	if m != nil {
		return m.DeclareTime
	}
	return 0
}

func (m *RsvpIdtStatus) GetWithdrawTime() uint32 {
	if m != nil {
		return m.WithdrawTime
	}
	return 0
}

// NSR/ISSU status information
type RsvpNsrIssuStatusInfo struct {
	// Process role
	Role string `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	// IDT status
	IdtStatus *RsvpIdtStatus `protobuf:"bytes,2,opt,name=idt_status,json=idtStatus" json:"idt_status,omitempty"`
	// Previous IDT status
	PreviousIdtStatus    *RsvpIdtStatus `protobuf:"bytes,3,opt,name=previous_idt_status,json=previousIdtStatus" json:"previous_idt_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RsvpNsrIssuStatusInfo) Reset()         { *m = RsvpNsrIssuStatusInfo{} }
func (m *RsvpNsrIssuStatusInfo) String() string { return proto.CompactTextString(m) }
func (*RsvpNsrIssuStatusInfo) ProtoMessage()    {}
func (*RsvpNsrIssuStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_global_info_d30ce4ecb4c384b9, []int{4}
}
func (m *RsvpNsrIssuStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpNsrIssuStatusInfo.Unmarshal(m, b)
}
func (m *RsvpNsrIssuStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpNsrIssuStatusInfo.Marshal(b, m, deterministic)
}
func (dst *RsvpNsrIssuStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpNsrIssuStatusInfo.Merge(dst, src)
}
func (m *RsvpNsrIssuStatusInfo) XXX_Size() int {
	return xxx_messageInfo_RsvpNsrIssuStatusInfo.Size(m)
}
func (m *RsvpNsrIssuStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpNsrIssuStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpNsrIssuStatusInfo proto.InternalMessageInfo

func (m *RsvpNsrIssuStatusInfo) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *RsvpNsrIssuStatusInfo) GetIdtStatus() *RsvpIdtStatus {
	if m != nil {
		return m.IdtStatus
	}
	return nil
}

func (m *RsvpNsrIssuStatusInfo) GetPreviousIdtStatus() *RsvpIdtStatus {
	if m != nil {
		return m.PreviousIdtStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*RsvpGlobalInfo_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.summary.rsvp_global_info_KEYS")
	proto.RegisterType((*RsvpGlobalInfo)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.summary.rsvp_global_info")
	proto.RegisterType((*RsvpMgmtAllDbCounters)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.summary.rsvp_mgmt_all_db_counters")
	proto.RegisterType((*RsvpIdtStatus)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.summary.rsvp_idt_status")
	proto.RegisterType((*RsvpNsrIssuStatusInfo)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.summary.rsvp_nsr_issu_status_info")
}

func init() {
	proto.RegisterFile("rsvp_global_info.proto", fileDescriptor_rsvp_global_info_d30ce4ecb4c384b9)
}

var fileDescriptor_rsvp_global_info_d30ce4ecb4c384b9 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4b, 0x8f, 0xd3, 0x3c,
	0x14, 0x55, 0x3b, 0x9d, 0xd1, 0xd7, 0xdb, 0xc7, 0xd7, 0xf1, 0x30, 0x10, 0x58, 0x40, 0x29, 0x0f,
	0x75, 0xd5, 0x45, 0xcb, 0x63, 0x8d, 0xd0, 0x2c, 0x10, 0x9b, 0x51, 0x8a, 0x90, 0x58, 0x19, 0x27,
	0xf6, 0x64, 0x2c, 0x25, 0x76, 0xe4, 0xeb, 0x74, 0xe8, 0xaf, 0xe0, 0x8f, 0xb2, 0xe0, 0x0f, 0x20,
	0x21, 0x3b, 0x4e, 0x5f, 0x02, 0x89, 0xd7, 0x26, 0xb2, 0xcf, 0x3d, 0xf7, 0x1c, 0xdf, 0x87, 0x02,
	0xb7, 0x0d, 0xae, 0x4a, 0x9a, 0xe5, 0x3a, 0x61, 0x39, 0x95, 0xea, 0x4a, 0xcf, 0x4a, 0xa3, 0xad,
	0x26, 0x4f, 0x53, 0x89, 0xa9, 0xa6, 0x52, 0x23, 0xfd, 0x64, 0xa8, 0x2c, 0xa9, 0xe7, 0xe9, 0x52,
	0x98, 0x99, 0x3b, 0xcd, 0xb0, 0x2a, 0x0a, 0x66, 0xd6, 0x93, 0x3b, 0x70, 0x7e, 0xa8, 0x40, 0xdf,
	0x5e, 0x7c, 0x58, 0x4e, 0xbe, 0xb5, 0x61, 0x74, 0x18, 0x21, 0x09, 0xf4, 0x24, 0x62, 0x45, 0xd1,
	0x32, 0x5b, 0x61, 0x34, 0x1f, 0xb7, 0xa6, 0xbd, 0xf9, 0xab, 0xd9, 0xaf, 0x79, 0xf9, 0x0b, 0x55,
	0x68, 0xe8, 0x8e, 0x86, 0xd7, 0x8d, 0xc1, 0x21, 0x4b, 0x0f, 0x90, 0x8f, 0x00, 0x8e, 0x13, 0x2c,
	0x16, 0xff, 0xca, 0xa2, 0xab, 0xd0, 0x04, 0x87, 0xfb, 0x00, 0x52, 0x59, 0x61, 0xae, 0x58, 0x2a,
	0x30, 0x7a, 0x36, 0x6e, 0x4d, 0x07, 0xf1, 0x0e, 0x42, 0xce, 0xe0, 0x38, 0x47, 0x5a, 0x62, 0xf4,
	0xdc, 0x87, 0x3a, 0x39, 0x5e, 0x22, 0x51, 0x70, 0xca, 0x99, 0x65, 0x09, 0x43, 0x41, 0x53, 0x5d,
	0x39, 0x36, 0x46, 0x2f, 0xfe, 0xe0, 0x75, 0x45, 0x56, 0x58, 0xca, 0xf2, 0x9c, 0xf2, 0x64, 0x23,
	0x14, 0x8f, 0x1a, 0xed, 0xd7, 0x01, 0x99, 0x7c, 0x6e, 0xc3, 0xdd, 0x9f, 0xf2, 0xc9, 0x3d, 0xf8,
	0x0f, 0x05, 0xa2, 0xd4, 0x0a, 0xa3, 0x96, 0x7f, 0xe5, 0xe6, 0x4e, 0x9e, 0xc0, 0x50, 0xaa, 0x54,
	0x17, 0x52, 0x65, 0xb4, 0x64, 0xf6, 0x1a, 0xa3, 0xb6, 0x67, 0x0c, 0x1a, 0xf4, 0xd2, 0x81, 0x8e,
	0xa6, 0x2b, 0x9b, 0xe9, 0x2d, 0xed, 0xa8, 0xa6, 0x35, 0x68, 0x4d, 0x5b, 0xc0, 0xf9, 0x46, 0xcd,
	0x08, 0x14, 0x66, 0xc5, 0xac, 0xb7, 0xed, 0x78, 0xf6, 0xad, 0x26, 0x18, 0xef, 0xc4, 0x5c, 0xd2,
	0x46, 0x7b, 0x2f, 0xe9, 0xb8, 0x4e, 0x6a, 0x82, 0x7b, 0x49, 0xfb, 0x63, 0x39, 0x39, 0x1c, 0xcb,
	0xe4, 0x4b, 0x0b, 0xfe, 0xf7, 0x1d, 0x91, 0xdc, 0x86, 0xd1, 0x92, 0x07, 0xd0, 0xc3, 0xb5, 0x4a,
	0x9b, 0x6d, 0x71, 0xad, 0xe8, 0xc6, 0xe0, 0xa0, 0x30, 0xeb, 0x29, 0x8c, 0x94, 0xb6, 0xd4, 0x08,
	0xc6, 0xd7, 0xee, 0x8b, 0x5a, 0xf9, 0x76, 0x74, 0xe3, 0xa1, 0xd2, 0x36, 0x76, 0x70, 0xec, 0x51,
	0xf2, 0x18, 0x86, 0x41, 0xd8, 0x58, 0x6a, 0x65, 0x21, 0x42, 0x3f, 0xfa, 0x92, 0xdb, 0xa5, 0x03,
	0xdf, 0xc9, 0x42, 0x90, 0x31, 0xb8, 0x3b, 0x15, 0x8a, 0xd7, 0x9c, 0x4e, 0x78, 0x26, 0xb7, 0x17,
	0x8a, 0x7b, 0xc6, 0x43, 0xe8, 0x73, 0x91, 0xe6, 0xcc, 0x88, 0x9a, 0x51, 0x97, 0xdc, 0x0b, 0x98,
	0xa7, 0x3c, 0x82, 0xc1, 0x8d, 0xb4, 0xd7, 0xdc, 0xb0, 0x9b, 0x9a, 0x53, 0x17, 0xdb, 0x6f, 0x40,
	0x47, 0x9a, 0x7c, 0x6d, 0x85, 0x05, 0xf8, 0xd1, 0x3a, 0x13, 0x02, 0x1d, 0xa3, 0x73, 0x11, 0x2a,
	0xf6, 0x67, 0xf2, 0x1e, 0x60, 0xdb, 0x1a, 0x5f, 0x65, 0x6f, 0xfe, 0xf2, 0xb7, 0x76, 0x73, 0x9b,
	0x1e, 0x77, 0xeb, 0xb2, 0x5d, 0x0f, 0x33, 0x38, 0x2b, 0x8d, 0x58, 0x49, 0xed, 0xcc, 0xb7, 0x06,
	0x47, 0x7f, 0x67, 0x70, 0xda, 0x68, 0xbe, 0x69, 0x8c, 0x92, 0x13, 0xff, 0xef, 0x5a, 0x7c, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0xc9, 0x2e, 0xd9, 0x6f, 0xd5, 0x04, 0x00, 0x00,
}