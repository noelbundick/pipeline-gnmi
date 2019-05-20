// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_mgmt_authentication_detail.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_standby_authentication_details_authentication_detail

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

// Authentication Detail Information
type RsvpMgmtAuthenticationDetail_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,2,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	ModeId               string   `protobuf:"bytes,3,opt,name=mode_id,json=modeId" json:"mode_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtAuthenticationDetail_KEYS) Reset()         { *m = RsvpMgmtAuthenticationDetail_KEYS{} }
func (m *RsvpMgmtAuthenticationDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtAuthenticationDetail_KEYS) ProtoMessage()    {}
func (*RsvpMgmtAuthenticationDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_authentication_detail_d8b948686a9b1bbf, []int{0}
}
func (m *RsvpMgmtAuthenticationDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtAuthenticationDetail_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtAuthenticationDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtAuthenticationDetail_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtAuthenticationDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtAuthenticationDetail_KEYS.Merge(dst, src)
}
func (m *RsvpMgmtAuthenticationDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtAuthenticationDetail_KEYS.Size(m)
}
func (m *RsvpMgmtAuthenticationDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtAuthenticationDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtAuthenticationDetail_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtAuthenticationDetail_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *RsvpMgmtAuthenticationDetail_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RsvpMgmtAuthenticationDetail_KEYS) GetModeId() string {
	if m != nil {
		return m.ModeId
	}
	return ""
}

func (m *RsvpMgmtAuthenticationDetail_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type RsvpMgmtAuthenticationDetail struct {
	// Basic authentication data
	AuthCompact *RsvpMgmtAuthenticationCompact `protobuf:"bytes,50,opt,name=auth_compact,json=authCompact" json:"auth_compact,omitempty"`
	// Key status
	KeyStatus uint32 `protobuf:"varint,51,opt,name=key_status,json=keyStatus" json:"key_status,omitempty"`
	// Direction
	KeyDigestInfo uint32 `protobuf:"varint,52,opt,name=key_digest_info,json=keyDigestInfo" json:"key_digest_info,omitempty"`
	// Lifetime (seconds)
	Lifetime uint32 `protobuf:"varint,53,opt,name=lifetime" json:"lifetime,omitempty"`
	// Remaining lifetime (seconds)
	LifetimeLeft uint32 `protobuf:"varint,54,opt,name=lifetime_left,json=lifetimeLeft" json:"lifetime_left,omitempty"`
	// Challenge status
	ChallengeStatus      string                 `protobuf:"bytes,55,opt,name=challenge_status,json=challengeStatus" json:"challenge_status,omitempty"`
	DirectionInfo        *RsvpMgmtAuthDirection `protobuf:"bytes,56,opt,name=direction_info,json=directionInfo" json:"direction_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RsvpMgmtAuthenticationDetail) Reset()         { *m = RsvpMgmtAuthenticationDetail{} }
func (m *RsvpMgmtAuthenticationDetail) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtAuthenticationDetail) ProtoMessage()    {}
func (*RsvpMgmtAuthenticationDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_authentication_detail_d8b948686a9b1bbf, []int{1}
}
func (m *RsvpMgmtAuthenticationDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtAuthenticationDetail.Unmarshal(m, b)
}
func (m *RsvpMgmtAuthenticationDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtAuthenticationDetail.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtAuthenticationDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtAuthenticationDetail.Merge(dst, src)
}
func (m *RsvpMgmtAuthenticationDetail) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtAuthenticationDetail.Size(m)
}
func (m *RsvpMgmtAuthenticationDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtAuthenticationDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtAuthenticationDetail proto.InternalMessageInfo

func (m *RsvpMgmtAuthenticationDetail) GetAuthCompact() *RsvpMgmtAuthenticationCompact {
	if m != nil {
		return m.AuthCompact
	}
	return nil
}

func (m *RsvpMgmtAuthenticationDetail) GetKeyStatus() uint32 {
	if m != nil {
		return m.KeyStatus
	}
	return 0
}

func (m *RsvpMgmtAuthenticationDetail) GetKeyDigestInfo() uint32 {
	if m != nil {
		return m.KeyDigestInfo
	}
	return 0
}

func (m *RsvpMgmtAuthenticationDetail) GetLifetime() uint32 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *RsvpMgmtAuthenticationDetail) GetLifetimeLeft() uint32 {
	if m != nil {
		return m.LifetimeLeft
	}
	return 0
}

func (m *RsvpMgmtAuthenticationDetail) GetChallengeStatus() string {
	if m != nil {
		return m.ChallengeStatus
	}
	return ""
}

func (m *RsvpMgmtAuthenticationDetail) GetDirectionInfo() *RsvpMgmtAuthDirection {
	if m != nil {
		return m.DirectionInfo
	}
	return nil
}

// RSVP Authentication Send Type Counters
type RsvpMgmtAuthSendTypeCounters struct {
	// Messages sent
	AuthenticationSent uint32 `protobuf:"varint,1,opt,name=authentication_sent,json=authenticationSent" json:"authentication_sent,omitempty"`
	// Failures
	AuthenticationFailures uint32 `protobuf:"varint,2,opt,name=authentication_failures,json=authenticationFailures" json:"authentication_failures,omitempty"`
	// Challenges received
	AuthenticationSendChallengesReceived uint32 `protobuf:"varint,3,opt,name=authentication_send_challenges_received,json=authenticationSendChallengesReceived" json:"authentication_send_challenges_received,omitempty"`
	// Challenge responses sent
	AuthenticationChallengeResponsesSent uint32   `protobuf:"varint,4,opt,name=authentication_challenge_responses_sent,json=authenticationChallengeResponsesSent" json:"authentication_challenge_responses_sent,omitempty"`
	XXX_NoUnkeyedLiteral                 struct{} `json:"-"`
	XXX_unrecognized                     []byte   `json:"-"`
	XXX_sizecache                        int32    `json:"-"`
}

func (m *RsvpMgmtAuthSendTypeCounters) Reset()         { *m = RsvpMgmtAuthSendTypeCounters{} }
func (m *RsvpMgmtAuthSendTypeCounters) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtAuthSendTypeCounters) ProtoMessage()    {}
func (*RsvpMgmtAuthSendTypeCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_authentication_detail_d8b948686a9b1bbf, []int{2}
}
func (m *RsvpMgmtAuthSendTypeCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtAuthSendTypeCounters.Unmarshal(m, b)
}
func (m *RsvpMgmtAuthSendTypeCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtAuthSendTypeCounters.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtAuthSendTypeCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtAuthSendTypeCounters.Merge(dst, src)
}
func (m *RsvpMgmtAuthSendTypeCounters) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtAuthSendTypeCounters.Size(m)
}
func (m *RsvpMgmtAuthSendTypeCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtAuthSendTypeCounters.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtAuthSendTypeCounters proto.InternalMessageInfo

func (m *RsvpMgmtAuthSendTypeCounters) GetAuthenticationSent() uint32 {
	if m != nil {
		return m.AuthenticationSent
	}
	return 0
}

func (m *RsvpMgmtAuthSendTypeCounters) GetAuthenticationFailures() uint32 {
	if m != nil {
		return m.AuthenticationFailures
	}
	return 0
}

func (m *RsvpMgmtAuthSendTypeCounters) GetAuthenticationSendChallengesReceived() uint32 {
	if m != nil {
		return m.AuthenticationSendChallengesReceived
	}
	return 0
}

func (m *RsvpMgmtAuthSendTypeCounters) GetAuthenticationChallengeResponsesSent() uint32 {
	if m != nil {
		return m.AuthenticationChallengeResponsesSent
	}
	return 0
}

// RSVP Authentication Receive Type Counters
type RsvpMgmtAuthRecvTypeCounters struct {
	// Valid messages
	AuthenticationReceivedValidMessages uint32 `protobuf:"varint,1,opt,name=authentication_received_valid_messages,json=authenticationReceivedValidMessages" json:"authentication_received_valid_messages,omitempty"`
	// Challenges sent
	AuthenticationReceivedChallengesSent uint32 `protobuf:"varint,2,opt,name=authentication_received_challenges_sent,json=authenticationReceivedChallengesSent" json:"authentication_received_challenges_sent,omitempty"`
	// Challenge responses received
	AuthenticationReceivedChallengeResponse uint32 `protobuf:"varint,3,opt,name=authentication_received_challenge_response,json=authenticationReceivedChallengeResponse" json:"authentication_received_challenge_response,omitempty"`
	// Challenges resent
	AuthenticationReceivedChallengesResent uint32 `protobuf:"varint,4,opt,name=authentication_received_challenges_resent,json=authenticationReceivedChallengesResent" json:"authentication_received_challenges_resent,omitempty"`
	// Challenge timeouts
	AuthenticationReceivedChallengeTimeouts uint32 `protobuf:"varint,5,opt,name=authentication_received_challenge_timeouts,json=authenticationReceivedChallengeTimeouts" json:"authentication_received_challenge_timeouts,omitempty"`
	// Authentication received during challenge
	AuthenticationReceivedDuringChallenge uint32 `protobuf:"varint,6,opt,name=authentication_received_during_challenge,json=authenticationReceivedDuringChallenge" json:"authentication_received_during_challenge,omitempty"`
	// Authentication received incomplete
	AuthenticationReceivedIncomplete uint32 `protobuf:"varint,7,opt,name=authentication_received_incomplete,json=authenticationReceivedIncomplete" json:"authentication_received_incomplete,omitempty"`
	// Authentication received with no integrity
	AuthenticationReceivedNoIntegrity uint32 `protobuf:"varint,8,opt,name=authentication_received_no_integrity,json=authenticationReceivedNoIntegrity" json:"authentication_received_no_integrity,omitempty"`
	// Authentication received with bad digest
	AuthenticationReceivedBadDigest uint32 `protobuf:"varint,9,opt,name=authentication_received_bad_digest,json=authenticationReceivedBadDigest" json:"authentication_received_bad_digest,omitempty"`
	// Authentication received with wrong digest type
	AuthenticationReceivedWrongDigestType uint32 `protobuf:"varint,10,opt,name=authentication_received_wrong_digest_type,json=authenticationReceivedWrongDigestType" json:"authentication_received_wrong_digest_type,omitempty"`
	// Authentication received with duplicate sequence number
	AuthenticationReceivedSequenceNumberDuplicate uint32 `protobuf:"varint,11,opt,name=authentication_received_sequence_number_duplicate,json=authenticationReceivedSequenceNumberDuplicate" json:"authentication_received_sequence_number_duplicate,omitempty"`
	// Authentication received with sequence number out of range
	AuthenticationReceivedSequenceNumberOutofRange uint32 `protobuf:"varint,12,opt,name=authentication_received_sequence_number_outof_range,json=authenticationReceivedSequenceNumberOutofRange" json:"authentication_received_sequence_number_outof_range,omitempty"`
	// Incorect challenge responses received
	AuthenticationReceivedWrongChallengesResponse uint32 `protobuf:"varint,13,opt,name=authentication_received_wrong_challenges_response,json=authenticationReceivedWrongChallengesResponse" json:"authentication_received_wrong_challenges_response,omitempty"`
	// Duplicate challenge responses received
	AuthenticationReceivedChallengesResponseDuplicate uint32 `protobuf:"varint,14,opt,name=authentication_received_challenges_response_duplicate,json=authenticationReceivedChallengesResponseDuplicate" json:"authentication_received_challenges_response_duplicate,omitempty"`
	// Challenge responses received late
	AuthenticationReceivedResponseLate uint32 `protobuf:"varint,15,opt,name=authentication_received_response_late,json=authenticationReceivedResponseLate" json:"authentication_received_response_late,omitempty"`
	// Authentication received with bad mesage format
	AuthenticationReceivedBadMessageFormat uint32   `protobuf:"varint,16,opt,name=authentication_received_bad_message_format,json=authenticationReceivedBadMessageFormat" json:"authentication_received_bad_message_format,omitempty"`
	XXX_NoUnkeyedLiteral                   struct{} `json:"-"`
	XXX_unrecognized                       []byte   `json:"-"`
	XXX_sizecache                          int32    `json:"-"`
}

func (m *RsvpMgmtAuthRecvTypeCounters) Reset()         { *m = RsvpMgmtAuthRecvTypeCounters{} }
func (m *RsvpMgmtAuthRecvTypeCounters) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtAuthRecvTypeCounters) ProtoMessage()    {}
func (*RsvpMgmtAuthRecvTypeCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_authentication_detail_d8b948686a9b1bbf, []int{3}
}
func (m *RsvpMgmtAuthRecvTypeCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtAuthRecvTypeCounters.Unmarshal(m, b)
}
func (m *RsvpMgmtAuthRecvTypeCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtAuthRecvTypeCounters.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtAuthRecvTypeCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtAuthRecvTypeCounters.Merge(dst, src)
}
func (m *RsvpMgmtAuthRecvTypeCounters) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtAuthRecvTypeCounters.Size(m)
}
func (m *RsvpMgmtAuthRecvTypeCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtAuthRecvTypeCounters.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtAuthRecvTypeCounters proto.InternalMessageInfo

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedValidMessages() uint32 {
	if m != nil {
		return m.AuthenticationReceivedValidMessages
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedChallengesSent() uint32 {
	if m != nil {
		return m.AuthenticationReceivedChallengesSent
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedChallengeResponse() uint32 {
	if m != nil {
		return m.AuthenticationReceivedChallengeResponse
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedChallengesResent() uint32 {
	if m != nil {
		return m.AuthenticationReceivedChallengesResent
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedChallengeTimeouts() uint32 {
	if m != nil {
		return m.AuthenticationReceivedChallengeTimeouts
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedDuringChallenge() uint32 {
	if m != nil {
		return m.AuthenticationReceivedDuringChallenge
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedIncomplete() uint32 {
	if m != nil {
		return m.AuthenticationReceivedIncomplete
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedNoIntegrity() uint32 {
	if m != nil {
		return m.AuthenticationReceivedNoIntegrity
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedBadDigest() uint32 {
	if m != nil {
		return m.AuthenticationReceivedBadDigest
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedWrongDigestType() uint32 {
	if m != nil {
		return m.AuthenticationReceivedWrongDigestType
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedSequenceNumberDuplicate() uint32 {
	if m != nil {
		return m.AuthenticationReceivedSequenceNumberDuplicate
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedSequenceNumberOutofRange() uint32 {
	if m != nil {
		return m.AuthenticationReceivedSequenceNumberOutofRange
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedWrongChallengesResponse() uint32 {
	if m != nil {
		return m.AuthenticationReceivedWrongChallengesResponse
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedChallengesResponseDuplicate() uint32 {
	if m != nil {
		return m.AuthenticationReceivedChallengesResponseDuplicate
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedResponseLate() uint32 {
	if m != nil {
		return m.AuthenticationReceivedResponseLate
	}
	return 0
}

func (m *RsvpMgmtAuthRecvTypeCounters) GetAuthenticationReceivedBadMessageFormat() uint32 {
	if m != nil {
		return m.AuthenticationReceivedBadMessageFormat
	}
	return 0
}

// RSVP Authentication Send Type Information
type RsvpMgmtAuthDirectionSend struct {
	// Sequence number
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	// Counters
	Counters             *RsvpMgmtAuthSendTypeCounters `protobuf:"bytes,2,opt,name=counters" json:"counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *RsvpMgmtAuthDirectionSend) Reset()         { *m = RsvpMgmtAuthDirectionSend{} }
func (m *RsvpMgmtAuthDirectionSend) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtAuthDirectionSend) ProtoMessage()    {}
func (*RsvpMgmtAuthDirectionSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_authentication_detail_d8b948686a9b1bbf, []int{4}
}
func (m *RsvpMgmtAuthDirectionSend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtAuthDirectionSend.Unmarshal(m, b)
}
func (m *RsvpMgmtAuthDirectionSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtAuthDirectionSend.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtAuthDirectionSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtAuthDirectionSend.Merge(dst, src)
}
func (m *RsvpMgmtAuthDirectionSend) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtAuthDirectionSend.Size(m)
}
func (m *RsvpMgmtAuthDirectionSend) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtAuthDirectionSend.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtAuthDirectionSend proto.InternalMessageInfo

func (m *RsvpMgmtAuthDirectionSend) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RsvpMgmtAuthDirectionSend) GetCounters() *RsvpMgmtAuthSendTypeCounters {
	if m != nil {
		return m.Counters
	}
	return nil
}

// RSVP Authentication Recieve Type Information
type RsvpMgmtAuthDirectionRecv struct {
	// Sequence number
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	// Sequence window size
	SequenceWindowSize uint32 `protobuf:"varint,2,opt,name=sequence_window_size,json=sequenceWindowSize" json:"sequence_window_size,omitempty"`
	// Sequence window count
	SequenceWindowCount uint32 `protobuf:"varint,3,opt,name=sequence_window_count,json=sequenceWindowCount" json:"sequence_window_count,omitempty"`
	// Sequence window
	SequenceWindow []uint64 `protobuf:"varint,4,rep,packed,name=sequence_window,json=sequenceWindow" json:"sequence_window,omitempty"`
	// Counters
	Counters             *RsvpMgmtAuthRecvTypeCounters `protobuf:"bytes,5,opt,name=counters" json:"counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *RsvpMgmtAuthDirectionRecv) Reset()         { *m = RsvpMgmtAuthDirectionRecv{} }
func (m *RsvpMgmtAuthDirectionRecv) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtAuthDirectionRecv) ProtoMessage()    {}
func (*RsvpMgmtAuthDirectionRecv) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_authentication_detail_d8b948686a9b1bbf, []int{5}
}
func (m *RsvpMgmtAuthDirectionRecv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtAuthDirectionRecv.Unmarshal(m, b)
}
func (m *RsvpMgmtAuthDirectionRecv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtAuthDirectionRecv.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtAuthDirectionRecv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtAuthDirectionRecv.Merge(dst, src)
}
func (m *RsvpMgmtAuthDirectionRecv) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtAuthDirectionRecv.Size(m)
}
func (m *RsvpMgmtAuthDirectionRecv) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtAuthDirectionRecv.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtAuthDirectionRecv proto.InternalMessageInfo

func (m *RsvpMgmtAuthDirectionRecv) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RsvpMgmtAuthDirectionRecv) GetSequenceWindowSize() uint32 {
	if m != nil {
		return m.SequenceWindowSize
	}
	return 0
}

func (m *RsvpMgmtAuthDirectionRecv) GetSequenceWindowCount() uint32 {
	if m != nil {
		return m.SequenceWindowCount
	}
	return 0
}

func (m *RsvpMgmtAuthDirectionRecv) GetSequenceWindow() []uint64 {
	if m != nil {
		return m.SequenceWindow
	}
	return nil
}

func (m *RsvpMgmtAuthDirectionRecv) GetCounters() *RsvpMgmtAuthRecvTypeCounters {
	if m != nil {
		return m.Counters
	}
	return nil
}

// Union of the different RSVP Authentication Direction types
type RsvpMgmtAuthDirection struct {
	AuthDirection        string                     `protobuf:"bytes,1,opt,name=auth_direction,json=authDirection" json:"auth_direction,omitempty"`
	SendInfo             *RsvpMgmtAuthDirectionSend `protobuf:"bytes,2,opt,name=send_info,json=sendInfo" json:"send_info,omitempty"`
	ReceiveInfo          *RsvpMgmtAuthDirectionRecv `protobuf:"bytes,3,opt,name=receive_info,json=receiveInfo" json:"receive_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RsvpMgmtAuthDirection) Reset()         { *m = RsvpMgmtAuthDirection{} }
func (m *RsvpMgmtAuthDirection) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtAuthDirection) ProtoMessage()    {}
func (*RsvpMgmtAuthDirection) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_authentication_detail_d8b948686a9b1bbf, []int{6}
}
func (m *RsvpMgmtAuthDirection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtAuthDirection.Unmarshal(m, b)
}
func (m *RsvpMgmtAuthDirection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtAuthDirection.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtAuthDirection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtAuthDirection.Merge(dst, src)
}
func (m *RsvpMgmtAuthDirection) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtAuthDirection.Size(m)
}
func (m *RsvpMgmtAuthDirection) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtAuthDirection.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtAuthDirection proto.InternalMessageInfo

func (m *RsvpMgmtAuthDirection) GetAuthDirection() string {
	if m != nil {
		return m.AuthDirection
	}
	return ""
}

func (m *RsvpMgmtAuthDirection) GetSendInfo() *RsvpMgmtAuthDirectionSend {
	if m != nil {
		return m.SendInfo
	}
	return nil
}

func (m *RsvpMgmtAuthDirection) GetReceiveInfo() *RsvpMgmtAuthDirectionRecv {
	if m != nil {
		return m.ReceiveInfo
	}
	return nil
}

// Authentication Compact Information
type RsvpMgmtAuthenticationCompact struct {
	// Source address
	SourceAddress string `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	// Destination address
	DestinationAddress string `protobuf:"bytes,2,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// Neighbor address
	NeighborAddress string `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	// Direction
	Direction string `protobuf:"bytes,4,opt,name=direction" json:"direction,omitempty"`
	// Key type
	KeyType string `protobuf:"bytes,5,opt,name=key_type,json=keyType" json:"key_type,omitempty"`
	// Key source
	KeySource string `protobuf:"bytes,6,opt,name=key_source,json=keySource" json:"key_source,omitempty"`
	// Key ID
	KeyId uint64 `protobuf:"varint,7,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	// Key validity
	KeyIdValid           uint32   `protobuf:"varint,8,opt,name=key_id_valid,json=keyIdValid" json:"key_id_valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtAuthenticationCompact) Reset()         { *m = RsvpMgmtAuthenticationCompact{} }
func (m *RsvpMgmtAuthenticationCompact) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtAuthenticationCompact) ProtoMessage()    {}
func (*RsvpMgmtAuthenticationCompact) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_authentication_detail_d8b948686a9b1bbf, []int{7}
}
func (m *RsvpMgmtAuthenticationCompact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtAuthenticationCompact.Unmarshal(m, b)
}
func (m *RsvpMgmtAuthenticationCompact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtAuthenticationCompact.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtAuthenticationCompact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtAuthenticationCompact.Merge(dst, src)
}
func (m *RsvpMgmtAuthenticationCompact) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtAuthenticationCompact.Size(m)
}
func (m *RsvpMgmtAuthenticationCompact) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtAuthenticationCompact.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtAuthenticationCompact proto.InternalMessageInfo

func (m *RsvpMgmtAuthenticationCompact) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *RsvpMgmtAuthenticationCompact) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RsvpMgmtAuthenticationCompact) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *RsvpMgmtAuthenticationCompact) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *RsvpMgmtAuthenticationCompact) GetKeyType() string {
	if m != nil {
		return m.KeyType
	}
	return ""
}

func (m *RsvpMgmtAuthenticationCompact) GetKeySource() string {
	if m != nil {
		return m.KeySource
	}
	return ""
}

func (m *RsvpMgmtAuthenticationCompact) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *RsvpMgmtAuthenticationCompact) GetKeyIdValid() uint32 {
	if m != nil {
		return m.KeyIdValid
	}
	return 0
}

func init() {
	proto.RegisterType((*RsvpMgmtAuthenticationDetail_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.authentication_details.authentication_detail.rsvp_mgmt_authentication_detail_KEYS")
	proto.RegisterType((*RsvpMgmtAuthenticationDetail)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.authentication_details.authentication_detail.rsvp_mgmt_authentication_detail")
	proto.RegisterType((*RsvpMgmtAuthSendTypeCounters)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.authentication_details.authentication_detail.rsvp_mgmt_auth_send_type_counters")
	proto.RegisterType((*RsvpMgmtAuthRecvTypeCounters)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.authentication_details.authentication_detail.rsvp_mgmt_auth_recv_type_counters")
	proto.RegisterType((*RsvpMgmtAuthDirectionSend)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.authentication_details.authentication_detail.rsvp_mgmt_auth_direction_send")
	proto.RegisterType((*RsvpMgmtAuthDirectionRecv)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.authentication_details.authentication_detail.rsvp_mgmt_auth_direction_recv")
	proto.RegisterType((*RsvpMgmtAuthDirection)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.authentication_details.authentication_detail.rsvp_mgmt_auth_direction")
	proto.RegisterType((*RsvpMgmtAuthenticationCompact)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.authentication_details.authentication_detail.rsvp_mgmt_authentication_compact")
}

func init() {
	proto.RegisterFile("rsvp_mgmt_authentication_detail.proto", fileDescriptor_rsvp_mgmt_authentication_detail_d8b948686a9b1bbf)
}

var fileDescriptor_rsvp_mgmt_authentication_detail_d8b948686a9b1bbf = []byte{
	// 1149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x96, 0x5d, 0xe7, 0xc7, 0xe3, 0x38, 0x89, 0xa6, 0x94, 0x2e, 0x88, 0xaa, 0xae, 0x9b, 0x5f,
	0x24, 0x0c, 0x4d, 0x28, 0xe5, 0x96, 0x26, 0x54, 0x8a, 0x1a, 0x52, 0xb1, 0x2e, 0x84, 0xc2, 0xc5,
	0x68, 0xb3, 0x73, 0xec, 0x8c, 0xec, 0xdd, 0x31, 0x3b, 0xb3, 0x09, 0xee, 0x03, 0x70, 0x8b, 0x84,
	0x40, 0xe2, 0x49, 0xb8, 0x41, 0xe2, 0x19, 0x78, 0x03, 0xae, 0xb8, 0xe2, 0x25, 0xd0, 0x9c, 0xdd,
	0x59, 0xc7, 0xce, 0xae, 0xed, 0x9b, 0x86, 0xbb, 0xdd, 0x33, 0xdf, 0xf9, 0xce, 0x39, 0xdf, 0xfc,
	0x9c, 0x19, 0xb2, 0x19, 0xa9, 0x8b, 0x01, 0x0b, 0xba, 0x81, 0x66, 0x5e, 0xac, 0xcf, 0x21, 0xd4,
	0xc2, 0xf7, 0xb4, 0x90, 0x21, 0xe3, 0xa0, 0x3d, 0xd1, 0x6f, 0x0d, 0x22, 0xa9, 0x25, 0x6d, 0xfb,
	0x42, 0xf9, 0x92, 0x09, 0xa9, 0xd8, 0x0f, 0x11, 0x13, 0x03, 0x86, 0x6e, 0x72, 0x00, 0x51, 0x0b,
	0xbf, 0x94, 0xf6, 0x42, 0x7e, 0x36, 0x6c, 0xe5, 0x72, 0xa8, 0x7c, 0x73, 0xf3, 0xcf, 0x12, 0xd9,
	0x98, 0x11, 0x9e, 0x3d, 0xff, 0xfc, 0x55, 0x9b, 0x6e, 0x92, 0x55, 0x25, 0xe3, 0xc8, 0x07, 0xe6,
	0x71, 0x1e, 0x81, 0x52, 0x4e, 0xa9, 0x51, 0xda, 0xa9, 0xba, 0xf5, 0xc4, 0xfa, 0x59, 0x62, 0xa4,
	0x1f, 0x92, 0xdb, 0x1c, 0x94, 0x16, 0x61, 0xc2, 0x60, 0xb1, 0x65, 0xc4, 0xd2, 0x2b, 0x43, 0xd6,
	0xe1, 0x2e, 0x59, 0x0a, 0x24, 0x07, 0x26, 0xb8, 0x73, 0x0b, 0x41, 0x8b, 0xe6, 0xf7, 0x88, 0x9b,
	0x80, 0x22, 0xd4, 0x10, 0x75, 0x3c, 0x1f, 0x58, 0xe8, 0x05, 0xe0, 0x54, 0x92, 0x80, 0x99, 0xf5,
	0xc4, 0x0b, 0xa0, 0xf9, 0x63, 0x85, 0xdc, 0x9f, 0x51, 0x00, 0xfd, 0xad, 0x44, 0x56, 0xcc, 0x08,
	0xf3, 0x65, 0x30, 0xf0, 0x7c, 0xed, 0xec, 0x35, 0x4a, 0x3b, 0xb5, 0xbd, 0xb8, 0xf5, 0x06, 0x14,
	0x6d, 0x15, 0x26, 0x93, 0x06, 0x77, 0x6b, 0xc6, 0x7e, 0x90, 0xfc, 0xd0, 0x7b, 0x84, 0xf4, 0x60,
	0x68, 0x22, 0xe9, 0x58, 0x39, 0xfb, 0x8d, 0xd2, 0x4e, 0xdd, 0xad, 0xf6, 0x60, 0xd8, 0x46, 0x03,
	0xdd, 0x22, 0x6b, 0x66, 0x98, 0x8b, 0x2e, 0x28, 0xcd, 0x44, 0xd8, 0x91, 0xce, 0xc7, 0x88, 0xa9,
	0xf7, 0x60, 0x78, 0x88, 0xd6, 0xa3, 0xb0, 0x23, 0xe9, 0xbb, 0x64, 0xb9, 0x2f, 0x3a, 0xa0, 0x45,
	0x00, 0xce, 0x63, 0x04, 0x64, 0xff, 0xf4, 0x21, 0xa9, 0xdb, 0x6f, 0xd6, 0x87, 0x8e, 0x76, 0x3e,
	0x41, 0xc0, 0x8a, 0x35, 0x1e, 0x43, 0x47, 0xd3, 0x5d, 0xb2, 0xee, 0x9f, 0x7b, 0xfd, 0x3e, 0x84,
	0x5d, 0xb0, 0xd9, 0x3c, 0x41, 0xbd, 0xd7, 0x32, 0x7b, 0x9a, 0xd3, 0x2f, 0x25, 0xb2, 0xca, 0x45,
	0x04, 0x3e, 0x56, 0x85, 0x39, 0x7d, 0x8a, 0x7a, 0x06, 0x37, 0xa0, 0x27, 0xcb, 0x22, 0xbb, 0xf5,
	0xec, 0xd3, 0x48, 0xd0, 0xfc, 0xa3, 0x4c, 0x1e, 0x4c, 0x60, 0x15, 0x84, 0x9c, 0xe9, 0xe1, 0x00,
	0x98, 0x2f, 0x63, 0xb3, 0x6a, 0x70, 0x7d, 0x4e, 0x84, 0x51, 0x10, 0x6a, 0x5c, 0xcb, 0x75, 0x97,
	0x8e, 0x0f, 0xb5, 0x21, 0xd4, 0xf4, 0x09, 0xb9, 0x3b, 0xe1, 0xd0, 0xf1, 0x44, 0x3f, 0x8e, 0x20,
	0x59, 0xd4, 0x75, 0xf7, 0xed, 0xf1, 0xe1, 0x67, 0xe9, 0x28, 0xfd, 0x8a, 0x6c, 0x5f, 0x8f, 0xc4,
	0x59, 0xa6, 0xa6, 0x62, 0x11, 0xf8, 0x20, 0x2e, 0x20, 0x59, 0xf8, 0x75, 0x77, 0xe3, 0x5a, 0x74,
	0x7e, 0x90, 0x81, 0xdd, 0x14, 0x9b, 0x43, 0x3b, 0x9a, 0xb7, 0x08, 0xd4, 0x40, 0x86, 0x0a, 0x54,
	0x52, 0x54, 0x25, 0x8f, 0x36, 0xa3, 0x74, 0x2d, 0xd8, 0x94, 0xd9, 0xfc, 0xa7, 0x76, 0x4d, 0xbd,
	0x08, 0xfc, 0x8b, 0x09, 0xf5, 0xda, 0x64, 0x6b, 0x22, 0xb8, 0xad, 0x81, 0x5d, 0x78, 0x7d, 0xc1,
	0x59, 0x00, 0x4a, 0x79, 0x5d, 0x50, 0xa9, 0xa0, 0x0f, 0xc7, 0xd1, 0xb6, 0x88, 0xaf, 0x0d, 0xf6,
	0x8b, 0x14, 0x9a, 0x53, 0x51, 0x46, 0x7a, 0x45, 0x2c, 0xac, 0xa8, 0x9c, 0x57, 0x91, 0x65, 0x1d,
	0x89, 0x85, 0x13, 0xf7, 0x1d, 0x79, 0x7f, 0x26, 0x6d, 0xa6, 0x58, 0x3a, 0x05, 0xdb, 0x33, 0x98,
	0xad, 0x66, 0xf4, 0x15, 0xd9, 0x9d, 0x23, 0xe7, 0x08, 0xae, 0xcc, 0xc3, 0xd6, 0xac, 0xac, 0x5d,
	0x44, 0xcf, 0x97, 0xb7, 0xd9, 0xb7, 0x32, 0xd6, 0xca, 0x59, 0x98, 0x2b, 0xef, 0x97, 0x29, 0x9c,
	0x9e, 0x92, 0x9d, 0x22, 0x72, 0x1e, 0x47, 0x22, 0xec, 0x8e, 0x62, 0x38, 0x8b, 0x48, 0xbd, 0x99,
	0x4f, 0x7d, 0x88, 0xe8, 0x2c, 0x00, 0x3d, 0x26, 0xcd, 0x22, 0x62, 0x11, 0x9a, 0x93, 0xaf, 0x0f,
	0x1a, 0x9c, 0x25, 0xa4, 0x6c, 0xe4, 0x53, 0x1e, 0x65, 0x38, 0xfa, 0x82, 0x6c, 0x14, 0xb1, 0x85,
	0x92, 0x99, 0x06, 0xd0, 0x8d, 0x84, 0x1e, 0x3a, 0xcb, 0xc8, 0xf7, 0x20, 0x9f, 0xef, 0x44, 0x1e,
	0x59, 0x20, 0x7d, 0x5e, 0x9c, 0xde, 0x99, 0xc7, 0xd3, 0xf3, 0xd5, 0xa9, 0x22, 0xdd, 0xfd, 0x7c,
	0xba, 0xa7, 0x1e, 0x4f, 0x0e, 0x5c, 0xfa, 0x4d, 0xf1, 0xe4, 0x5f, 0x46, 0x32, 0xec, 0xda, 0xe3,
	0xda, 0xec, 0x1d, 0x87, 0x4c, 0x53, 0xf1, 0xd4, 0xc0, 0x13, 0xd6, 0x97, 0xc3, 0x01, 0xd0, 0x73,
	0xf2, 0xa8, 0x88, 0x59, 0xc1, 0xf7, 0x31, 0x84, 0xa6, 0x15, 0xc6, 0xc1, 0x19, 0x44, 0x8c, 0xc7,
	0x83, 0xbe, 0x01, 0x81, 0x53, 0xc3, 0x08, 0x1f, 0xe4, 0x47, 0x68, 0xa7, 0x6e, 0x27, 0xe8, 0x75,
	0x68, 0x9d, 0x68, 0x8f, 0xec, 0xcf, 0x1b, 0x49, 0xc6, 0x5a, 0x76, 0x58, 0xe4, 0x99, 0x35, 0xb1,
	0x82, 0xb1, 0x5a, 0xf3, 0xc4, 0x7a, 0x61, 0xdc, 0x5c, 0xe3, 0x35, 0xad, 0xac, 0x44, 0xb0, 0xf1,
	0x3d, 0x93, 0xec, 0xc8, 0xfa, 0xb4, 0xb2, 0x50, 0xb8, 0xb1, 0xad, 0x93, 0xec, 0xcb, 0x01, 0x79,
	0x3c, 0xdf, 0xbe, 0x44, 0xf8, 0x15, 0x11, 0x57, 0x31, 0xda, 0xa3, 0x39, 0xf6, 0x28, 0x7a, 0x8e,
	0x84, 0xfc, 0x92, 0x6c, 0x16, 0x45, 0xcc, 0xc2, 0xf4, 0x4d, 0x84, 0x35, 0x8c, 0xd0, 0xcc, 0x8f,
	0x60, 0x79, 0x8f, 0x0d, 0xe5, 0xb7, 0xc5, 0x27, 0x80, 0x59, 0xac, 0xe9, 0x19, 0xcb, 0x3a, 0x32,
	0x0a, 0x3c, 0xed, 0xac, 0x4f, 0x3b, 0x5d, 0x9e, 0x7a, 0xf6, 0x9c, 0x7d, 0x86, 0xe8, 0xe6, 0x5f,
	0x25, 0x72, 0xaf, 0xa8, 0xa3, 0x62, 0x83, 0x32, 0x57, 0x09, 0xbb, 0x02, 0xf0, 0x14, 0xaf, 0xb8,
	0xd9, 0x3f, 0xfd, 0xb9, 0x44, 0x96, 0x6d, 0x33, 0xc0, 0xc3, 0xb8, 0xb6, 0x77, 0x71, 0x13, 0x4d,
	0xff, 0x7a, 0x23, 0x77, 0xb3, 0x3c, 0x9a, 0x7f, 0x97, 0xa7, 0x94, 0x64, 0x9a, 0xd8, 0xd4, 0x92,
	0x3e, 0x22, 0x6f, 0x65, 0x0b, 0xfe, 0x52, 0x84, 0x5c, 0x5e, 0x32, 0x25, 0x5e, 0x43, 0xda, 0x6a,
	0xa8, 0x1d, 0x3b, 0xc5, 0xa1, 0xb6, 0x78, 0x0d, 0x74, 0x8f, 0xdc, 0x99, 0xf4, 0xc0, 0x5c, 0xd2,
	0x1e, 0x72, 0x7b, 0xdc, 0xe5, 0xc0, 0x0c, 0xd1, 0x6d, 0xb2, 0x36, 0xe1, 0xe3, 0x54, 0x1a, 0xb7,
	0x76, 0x2a, 0xee, 0xea, 0x38, 0x7a, 0x5c, 0xe1, 0x85, 0x9b, 0x53, 0xf8, 0x7a, 0xb3, 0xbf, 0xa2,
	0xf0, 0xbf, 0x65, 0xe2, 0x14, 0x29, 0x6c, 0xee, 0xe9, 0xe3, 0x16, 0xfb, 0x30, 0x30, 0xd6, 0xc3,
	0x0c, 0xf6, 0x53, 0x89, 0x54, 0x71, 0x1a, 0xf1, 0xc2, 0x98, 0xac, 0x9d, 0xe8, 0x46, 0x2f, 0x8c,
	0xb8, 0x8a, 0xcc, 0xcc, 0x87, 0x1c, 0xef, 0xcc, 0xbf, 0x96, 0xc8, 0x4a, 0xba, 0xb1, 0x92, 0xa4,
	0x6e, 0xfd, 0x1f, 0x49, 0x19, 0xe1, 0xdd, 0x5a, 0x9a, 0x07, 0x5e, 0x64, 0x7f, 0x2f, 0x93, 0xc6,
	0xac, 0x47, 0xc4, 0x1b, 0x7b, 0x8e, 0xed, 0x92, 0xf5, 0x10, 0x44, 0xf7, 0xfc, 0x4c, 0x46, 0x19,
	0x3a, 0x79, 0x97, 0xad, 0x59, 0xbb, 0x85, 0xbe, 0x47, 0xaa, 0xa3, 0x39, 0x4f, 0xde, 0x66, 0x23,
	0x03, 0x7d, 0x87, 0x2c, 0x9b, 0x97, 0x0b, 0xf6, 0xc0, 0x05, 0x1c, 0x5c, 0xea, 0xc1, 0x10, 0xbb,
	0x9c, 0x7d, 0xf3, 0x60, 0xa6, 0x78, 0xcd, 0xa8, 0x26, 0x6f, 0x1e, 0x34, 0xd0, 0x3b, 0x64, 0xd1,
	0x0c, 0x0b, 0x8e, 0xd7, 0x85, 0x8a, 0xbb, 0xd0, 0x83, 0xe1, 0x11, 0xa7, 0x0d, 0xb2, 0x92, 0x98,
	0x93, 0xab, 0x66, 0xda, 0xfb, 0x09, 0x0e, 0xe2, 0x85, 0xf2, 0x6c, 0x11, 0xdf, 0xc9, 0xfb, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xf7, 0x05, 0x93, 0x50, 0x0f, 0x00, 0x00,
}