// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_mgmt_count_message.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_standby_counters_message_summary

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

// Counters for all messages
type RsvpMgmtCountMessage_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtCountMessage_KEYS) Reset()         { *m = RsvpMgmtCountMessage_KEYS{} }
func (m *RsvpMgmtCountMessage_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtCountMessage_KEYS) ProtoMessage()    {}
func (*RsvpMgmtCountMessage_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_count_message_39c13105804d766d, []int{0}
}
func (m *RsvpMgmtCountMessage_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtCountMessage_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtCountMessage_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtCountMessage_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtCountMessage_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtCountMessage_KEYS.Merge(dst, src)
}
func (m *RsvpMgmtCountMessage_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtCountMessage_KEYS.Size(m)
}
func (m *RsvpMgmtCountMessage_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtCountMessage_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtCountMessage_KEYS proto.InternalMessageInfo

type RsvpMgmtCountMessage struct {
	// Count of messages received
	ReceivedMessages *RsvpMgmtCounts `protobuf:"bytes,50,opt,name=received_messages,json=receivedMessages" json:"received_messages,omitempty"`
	// Count of messages transmitted
	TransmittedMessages *RsvpMgmtCounts `protobuf:"bytes,51,opt,name=transmitted_messages,json=transmittedMessages" json:"transmitted_messages,omitempty"`
	// Count of Bundle messages received
	BundleReceivedMessages *RsvpMgmtCounts `protobuf:"bytes,52,opt,name=bundle_received_messages,json=bundleReceivedMessages" json:"bundle_received_messages,omitempty"`
	// Count of Bundle messages transmitted
	BundleTransmittedMessages *RsvpMgmtCounts `protobuf:"bytes,53,opt,name=bundle_transmitted_messages,json=bundleTransmittedMessages" json:"bundle_transmitted_messages,omitempty"`
	// Count of messages retransmitted
	RetransmittedMessages uint32 `protobuf:"varint,54,opt,name=retransmitted_messages,json=retransmittedMessages" json:"retransmitted_messages,omitempty"`
	// Count of Out of Order messages
	OutOfOrderMessages uint32 `protobuf:"varint,55,opt,name=out_of_order_messages,json=outOfOrderMessages" json:"out_of_order_messages,omitempty"`
	// Count of Rate Limited messages
	RateLimitedMessages  uint32   `protobuf:"varint,56,opt,name=rate_limited_messages,json=rateLimitedMessages" json:"rate_limited_messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtCountMessage) Reset()         { *m = RsvpMgmtCountMessage{} }
func (m *RsvpMgmtCountMessage) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtCountMessage) ProtoMessage()    {}
func (*RsvpMgmtCountMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_count_message_39c13105804d766d, []int{1}
}
func (m *RsvpMgmtCountMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtCountMessage.Unmarshal(m, b)
}
func (m *RsvpMgmtCountMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtCountMessage.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtCountMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtCountMessage.Merge(dst, src)
}
func (m *RsvpMgmtCountMessage) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtCountMessage.Size(m)
}
func (m *RsvpMgmtCountMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtCountMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtCountMessage proto.InternalMessageInfo

func (m *RsvpMgmtCountMessage) GetReceivedMessages() *RsvpMgmtCounts {
	if m != nil {
		return m.ReceivedMessages
	}
	return nil
}

func (m *RsvpMgmtCountMessage) GetTransmittedMessages() *RsvpMgmtCounts {
	if m != nil {
		return m.TransmittedMessages
	}
	return nil
}

func (m *RsvpMgmtCountMessage) GetBundleReceivedMessages() *RsvpMgmtCounts {
	if m != nil {
		return m.BundleReceivedMessages
	}
	return nil
}

func (m *RsvpMgmtCountMessage) GetBundleTransmittedMessages() *RsvpMgmtCounts {
	if m != nil {
		return m.BundleTransmittedMessages
	}
	return nil
}

func (m *RsvpMgmtCountMessage) GetRetransmittedMessages() uint32 {
	if m != nil {
		return m.RetransmittedMessages
	}
	return 0
}

func (m *RsvpMgmtCountMessage) GetOutOfOrderMessages() uint32 {
	if m != nil {
		return m.OutOfOrderMessages
	}
	return 0
}

func (m *RsvpMgmtCountMessage) GetRateLimitedMessages() uint32 {
	if m != nil {
		return m.RateLimitedMessages
	}
	return 0
}

// Counters for all known RSVP message types
type RsvpMgmtCounts struct {
	// Count of Path messages
	Path uint32 `protobuf:"varint,1,opt,name=path" json:"path,omitempty"`
	// Count of Reservation messages
	Reservation uint32 `protobuf:"varint,2,opt,name=reservation" json:"reservation,omitempty"`
	// Count of PathError messages
	PathError uint32 `protobuf:"varint,3,opt,name=path_error,json=pathError" json:"path_error,omitempty"`
	// Count of ReservationError messages
	ReservationError uint32 `protobuf:"varint,4,opt,name=reservation_error,json=reservationError" json:"reservation_error,omitempty"`
	// Count of PathTear messages
	PathTear uint32 `protobuf:"varint,5,opt,name=path_tear,json=pathTear" json:"path_tear,omitempty"`
	// Count of ReservationTear messages
	ReservationTear uint32 `protobuf:"varint,6,opt,name=reservation_tear,json=reservationTear" json:"reservation_tear,omitempty"`
	// Count of ReservationConfirm messages
	ReservationConfirm uint32 `protobuf:"varint,7,opt,name=reservation_confirm,json=reservationConfirm" json:"reservation_confirm,omitempty"`
	// Count of Bundle messages
	Bundle uint32 `protobuf:"varint,8,opt,name=bundle" json:"bundle,omitempty"`
	// Count of ACK messages
	Ack uint32 `protobuf:"varint,9,opt,name=ack" json:"ack,omitempty"`
	// Count of Srefresh messages
	SRefresh uint32 `protobuf:"varint,10,opt,name=s_refresh,json=sRefresh" json:"s_refresh,omitempty"`
	// Count of Hello messages
	Hello uint32 `protobuf:"varint,11,opt,name=hello" json:"hello,omitempty"`
	// Count of Integrity Challenge messages
	Challenge uint32 `protobuf:"varint,12,opt,name=challenge" json:"challenge,omitempty"`
	// Count of Integrity Response messages
	Response             uint32   `protobuf:"varint,13,opt,name=response" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtCounts) Reset()         { *m = RsvpMgmtCounts{} }
func (m *RsvpMgmtCounts) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtCounts) ProtoMessage()    {}
func (*RsvpMgmtCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_count_message_39c13105804d766d, []int{2}
}
func (m *RsvpMgmtCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtCounts.Unmarshal(m, b)
}
func (m *RsvpMgmtCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtCounts.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtCounts.Merge(dst, src)
}
func (m *RsvpMgmtCounts) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtCounts.Size(m)
}
func (m *RsvpMgmtCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtCounts.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtCounts proto.InternalMessageInfo

func (m *RsvpMgmtCounts) GetPath() uint32 {
	if m != nil {
		return m.Path
	}
	return 0
}

func (m *RsvpMgmtCounts) GetReservation() uint32 {
	if m != nil {
		return m.Reservation
	}
	return 0
}

func (m *RsvpMgmtCounts) GetPathError() uint32 {
	if m != nil {
		return m.PathError
	}
	return 0
}

func (m *RsvpMgmtCounts) GetReservationError() uint32 {
	if m != nil {
		return m.ReservationError
	}
	return 0
}

func (m *RsvpMgmtCounts) GetPathTear() uint32 {
	if m != nil {
		return m.PathTear
	}
	return 0
}

func (m *RsvpMgmtCounts) GetReservationTear() uint32 {
	if m != nil {
		return m.ReservationTear
	}
	return 0
}

func (m *RsvpMgmtCounts) GetReservationConfirm() uint32 {
	if m != nil {
		return m.ReservationConfirm
	}
	return 0
}

func (m *RsvpMgmtCounts) GetBundle() uint32 {
	if m != nil {
		return m.Bundle
	}
	return 0
}

func (m *RsvpMgmtCounts) GetAck() uint32 {
	if m != nil {
		return m.Ack
	}
	return 0
}

func (m *RsvpMgmtCounts) GetSRefresh() uint32 {
	if m != nil {
		return m.SRefresh
	}
	return 0
}

func (m *RsvpMgmtCounts) GetHello() uint32 {
	if m != nil {
		return m.Hello
	}
	return 0
}

func (m *RsvpMgmtCounts) GetChallenge() uint32 {
	if m != nil {
		return m.Challenge
	}
	return 0
}

func (m *RsvpMgmtCounts) GetResponse() uint32 {
	if m != nil {
		return m.Response
	}
	return 0
}

func init() {
	proto.RegisterType((*RsvpMgmtCountMessage_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.counters.message_summary.rsvp_mgmt_count_message_KEYS")
	proto.RegisterType((*RsvpMgmtCountMessage)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.counters.message_summary.rsvp_mgmt_count_message")
	proto.RegisterType((*RsvpMgmtCounts)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.counters.message_summary.rsvp_mgmt_counts")
}

func init() {
	proto.RegisterFile("rsvp_mgmt_count_message.proto", fileDescriptor_rsvp_mgmt_count_message_39c13105804d766d)
}

var fileDescriptor_rsvp_mgmt_count_message_39c13105804d766d = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4f, 0x6f, 0x12, 0x41,
	0x14, 0xcf, 0x4a, 0x41, 0x78, 0xd8, 0x88, 0x43, 0xc1, 0xd1, 0xb6, 0x86, 0x70, 0xaa, 0x31, 0x59,
	0x23, 0xb5, 0xea, 0xcd, 0x83, 0xe9, 0x49, 0x4d, 0x23, 0xf6, 0xe2, 0x69, 0x32, 0x2c, 0x0f, 0xd8,
	0xb8, 0xbb, 0xb3, 0x79, 0x33, 0x10, 0x7b, 0xf1, 0xe4, 0xcd, 0x93, 0xdf, 0xc1, 0x0f, 0xe6, 0x47,
	0x31, 0x33, 0xb3, 0xac, 0x5b, 0xa4, 0xa7, 0x86, 0xdb, 0x7b, 0xbf, 0x3f, 0xfc, 0x7e, 0xbc, 0xcd,
	0x2e, 0x1c, 0x93, 0x5e, 0xe5, 0x22, 0x9d, 0xa7, 0x46, 0x44, 0x6a, 0x99, 0x19, 0x91, 0xa2, 0xd6,
	0x72, 0x8e, 0x61, 0x4e, 0xca, 0x28, 0xf6, 0x36, 0x8a, 0x75, 0xa4, 0x44, 0xac, 0xb4, 0xf8, 0x46,
	0x22, 0xce, 0x85, 0x93, 0xab, 0x1c, 0x29, 0x74, 0x93, 0x36, 0x32, 0x9b, 0x4e, 0xae, 0x42, 0xe7,
	0x45, 0xd2, 0x61, 0x61, 0x17, 0x7a, 0x99, 0xa6, 0x92, 0xae, 0x86, 0x4f, 0xe0, 0xe8, 0x86, 0x04,
	0xf1, 0xfe, 0xfc, 0xcb, 0xe7, 0xe1, 0x9f, 0x3a, 0x3c, 0xbc, 0x41, 0xc0, 0xbe, 0xc3, 0x03, 0xc2,
	0x08, 0xe3, 0x15, 0x4e, 0xd7, 0x98, 0xe6, 0xa3, 0x41, 0x70, 0xd2, 0x1e, 0x7d, 0x0a, 0x6f, 0x59,
	0x2c, 0xdc, 0x08, 0xd5, 0xe3, 0xce, 0x3a, 0xeb, 0x63, 0x11, 0xc5, 0x7e, 0x04, 0x70, 0x60, 0x48,
	0x66, 0x3a, 0x8d, 0x8d, 0xa9, 0x76, 0x38, 0xdd, 0x55, 0x87, 0x6e, 0x25, 0xae, 0xac, 0xf1, 0x33,
	0x00, 0x3e, 0x59, 0x66, 0xd3, 0x04, 0xc5, 0xff, 0xe7, 0x78, 0xb9, 0xab, 0x2a, 0x7d, 0x1f, 0x39,
	0xde, 0x3c, 0xca, 0xaf, 0x00, 0x0e, 0x8b, 0x36, 0x5b, 0x6f, 0x73, 0xb6, 0xab, 0x42, 0x8f, 0x7c,
	0xea, 0xe5, 0x96, 0x0b, 0x9d, 0x41, 0x9f, 0x70, 0x6b, 0x9b, 0x57, 0x83, 0xe0, 0x64, 0x7f, 0xdc,
	0xbb, 0xc6, 0x96, 0xb6, 0x17, 0xd0, 0x53, 0x4b, 0x23, 0xd4, 0x4c, 0x28, 0x9a, 0x22, 0xfd, 0x73,
	0xbd, 0x76, 0x2e, 0xa6, 0x96, 0xe6, 0x62, 0x76, 0x61, 0xa9, 0xd2, 0x32, 0x82, 0x1e, 0x49, 0x83,
	0x22, 0x89, 0xd3, 0xf8, 0x5a, 0xd0, 0x1b, 0x67, 0xe9, 0x5a, 0xf2, 0x83, 0xe7, 0xd6, 0x9e, 0xe1,
	0xef, 0x1a, 0x74, 0x36, 0xff, 0x0d, 0x63, 0xb0, 0x97, 0x4b, 0xb3, 0xe0, 0x81, 0xf3, 0xb9, 0x99,
	0x0d, 0xa0, 0x4d, 0xa8, 0x91, 0x56, 0xd2, 0xc4, 0x2a, 0xe3, 0x77, 0x1c, 0x55, 0x85, 0xd8, 0x31,
	0x80, 0x55, 0x0a, 0x24, 0x52, 0xc4, 0x6b, 0x4e, 0xd0, 0xb2, 0xc8, 0xb9, 0x05, 0xd8, 0x33, 0xfb,
	0xc2, 0x94, 0xea, 0x42, 0xb5, 0xe7, 0x54, 0x9d, 0x0a, 0xe1, 0xc5, 0x87, 0xe0, 0x9c, 0xc2, 0xa0,
	0x24, 0x5e, 0x77, 0xa2, 0xa6, 0x05, 0x2e, 0x51, 0x12, 0x7b, 0x0a, 0x55, 0x83, 0xd7, 0x34, 0x9c,
	0xe6, 0x7e, 0x05, 0x77, 0xd2, 0xe7, 0xd0, 0xad, 0x4a, 0x23, 0x95, 0xcd, 0x62, 0x4a, 0xf9, 0x5d,
	0x7f, 0xc3, 0x0a, 0xf5, 0xce, 0x33, 0xac, 0x0f, 0x0d, 0xff, 0x28, 0x79, 0xd3, 0x69, 0x8a, 0x8d,
	0x75, 0xa0, 0x26, 0xa3, 0xaf, 0xbc, 0xe5, 0x40, 0x3b, 0xda, 0x8a, 0x5a, 0x10, 0xce, 0x08, 0xf5,
	0x82, 0x83, 0xaf, 0xa8, 0xc7, 0x7e, 0x67, 0x07, 0x50, 0x5f, 0x60, 0x92, 0x28, 0xde, 0x76, 0x84,
	0x5f, 0xd8, 0x11, 0xb4, 0xa2, 0x85, 0x4c, 0x12, 0xcc, 0xe6, 0xc8, 0xef, 0xf9, 0x03, 0x95, 0x00,
	0x7b, 0x0c, 0x4d, 0x42, 0x9d, 0xab, 0x4c, 0x23, 0xdf, 0xf7, 0xbf, 0xb7, 0xde, 0x27, 0x0d, 0xf7,
	0xc5, 0x3b, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x7f, 0x2c, 0x98, 0x12, 0x05, 0x00, 0x00,
}