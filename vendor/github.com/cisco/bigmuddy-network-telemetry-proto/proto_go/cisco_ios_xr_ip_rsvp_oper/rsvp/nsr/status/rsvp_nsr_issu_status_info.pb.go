// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_nsr_issu_status_info.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_nsr_status

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

// NSR/ISSU status information
type RsvpNsrIssuStatusInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpNsrIssuStatusInfo_KEYS) Reset()         { *m = RsvpNsrIssuStatusInfo_KEYS{} }
func (m *RsvpNsrIssuStatusInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpNsrIssuStatusInfo_KEYS) ProtoMessage()    {}
func (*RsvpNsrIssuStatusInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_nsr_issu_status_info_baa3def3c18385da, []int{0}
}
func (m *RsvpNsrIssuStatusInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpNsrIssuStatusInfo_KEYS.Unmarshal(m, b)
}
func (m *RsvpNsrIssuStatusInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpNsrIssuStatusInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpNsrIssuStatusInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpNsrIssuStatusInfo_KEYS.Merge(dst, src)
}
func (m *RsvpNsrIssuStatusInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpNsrIssuStatusInfo_KEYS.Size(m)
}
func (m *RsvpNsrIssuStatusInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpNsrIssuStatusInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpNsrIssuStatusInfo_KEYS proto.InternalMessageInfo

type RsvpNsrIssuStatusInfo struct {
	// Process role
	Role string `protobuf:"bytes,50,opt,name=role" json:"role,omitempty"`
	// IDT status
	IdtStatus *RsvpIdtStatus `protobuf:"bytes,51,opt,name=idt_status,json=idtStatus" json:"idt_status,omitempty"`
	// Previous IDT status
	PreviousIdtStatus    *RsvpIdtStatus `protobuf:"bytes,52,opt,name=previous_idt_status,json=previousIdtStatus" json:"previous_idt_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RsvpNsrIssuStatusInfo) Reset()         { *m = RsvpNsrIssuStatusInfo{} }
func (m *RsvpNsrIssuStatusInfo) String() string { return proto.CompactTextString(m) }
func (*RsvpNsrIssuStatusInfo) ProtoMessage()    {}
func (*RsvpNsrIssuStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_nsr_issu_status_info_baa3def3c18385da, []int{1}
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
	return fileDescriptor_rsvp_nsr_issu_status_info_baa3def3c18385da, []int{2}
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

func init() {
	proto.RegisterType((*RsvpNsrIssuStatusInfo_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.nsr.status.rsvp_nsr_issu_status_info_KEYS")
	proto.RegisterType((*RsvpNsrIssuStatusInfo)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.nsr.status.rsvp_nsr_issu_status_info")
	proto.RegisterType((*RsvpIdtStatus)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.nsr.status.rsvp_idt_status")
}

func init() {
	proto.RegisterFile("rsvp_nsr_issu_status_info.proto", fileDescriptor_rsvp_nsr_issu_status_info_baa3def3c18385da)
}

var fileDescriptor_rsvp_nsr_issu_status_info_baa3def3c18385da = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0x95, 0xef, 0xb7, 0x54, 0xea, 0xf5, 0x07, 0x60, 0x96, 0xb0, 0xd0, 0x50, 0x18, 0xc2,
	0x92, 0xa1, 0x65, 0x62, 0xef, 0x80, 0xd8, 0x52, 0x96, 0x4e, 0xa7, 0x10, 0x1b, 0x61, 0xd4, 0xda,
	0xd1, 0xd9, 0x6d, 0xe9, 0xff, 0xcc, 0xcc, 0x8c, 0x7c, 0x49, 0x54, 0x09, 0xa9, 0x12, 0x12, 0x4b,
	0x74, 0xf7, 0xf4, 0xb9, 0xf7, 0xee, 0x22, 0xc3, 0x98, 0xdc, 0xb6, 0x42, 0xe3, 0x08, 0xb5, 0x73,
	0x1b, 0x74, 0xbe, 0xf0, 0x1b, 0x87, 0xda, 0xbc, 0xda, 0xac, 0x22, 0xeb, 0xad, 0xb8, 0x2b, 0xb5,
	0x2b, 0x2d, 0x6a, 0xeb, 0xf0, 0x83, 0x50, 0x57, 0xc8, 0x03, 0xb6, 0x52, 0x94, 0x85, 0x2a, 0x33,
	0x8e, 0xb2, 0x7a, 0x6a, 0x92, 0xc0, 0xd5, 0x51, 0x37, 0x7c, 0x9a, 0x2f, 0x17, 0x93, 0xaf, 0x08,
	0x2e, 0x8f, 0x22, 0x42, 0x40, 0x87, 0xec, 0x4a, 0xc5, 0xd3, 0x24, 0x4a, 0x7b, 0x39, 0xd7, 0x62,
	0x09, 0xa0, 0xa5, 0x6f, 0xb0, 0x78, 0x96, 0x44, 0x69, 0x7f, 0xfa, 0x90, 0xfd, 0x7a, 0x27, 0xee,
	0xf1, 0xe0, 0x90, 0xf7, 0xb4, 0xf4, 0x0b, 0x2e, 0xc5, 0x3b, 0x5c, 0x54, 0xa4, 0xb6, 0xda, 0x86,
	0xfc, 0x43, 0xc6, 0xfd, 0x9f, 0x33, 0xce, 0x5b, 0xdb, 0xc7, 0x36, 0x6b, 0xf2, 0x19, 0xc1, 0xe9,
	0x0f, 0x4c, 0x8c, 0xa1, 0xef, 0xf6, 0xa6, 0x6c, 0x73, 0x23, 0xbe, 0x1a, 0x82, 0xd4, 0x2c, 0x98,
	0xc2, 0x99, 0xb1, 0x1e, 0x49, 0x15, 0x72, 0x1f, 0xbe, 0xce, 0x9a, 0xf8, 0x1f, 0x53, 0x23, 0x63,
	0x7d, 0x1e, 0xe4, 0x9c, 0x55, 0x71, 0x0b, 0xa3, 0xc6, 0x98, 0x3c, 0x7a, 0xbd, 0x56, 0xf1, 0xff,
	0x24, 0x4a, 0x87, 0xf9, 0xa0, 0xbe, 0x96, 0xfc, 0xb3, 0x5e, 0x2b, 0x91, 0x40, 0xe8, 0x51, 0x19,
	0x59, 0x33, 0x1d, 0x66, 0xc2, 0xff, 0x9d, 0x1b, 0xc9, 0xc4, 0x35, 0x0c, 0xa4, 0x2a, 0x57, 0x05,
	0xa9, 0x9a, 0x38, 0x61, 0xa2, 0xdf, 0x68, 0x8c, 0xdc, 0xc0, 0x70, 0xa7, 0xfd, 0x9b, 0xa4, 0x62,
	0x57, 0x33, 0xdd, 0x3a, 0xa9, 0x15, 0x03, 0xf4, 0xd2, 0xe5, 0xb7, 0x33, 0xfb, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x87, 0x41, 0xd3, 0xfc, 0x5e, 0x02, 0x00, 0x00,
}