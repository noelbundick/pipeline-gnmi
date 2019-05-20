// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rsvp_mgmt_authentication_compact.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_authentication_briefs_authentication_brief

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

// Authentication Compact Information
type RsvpMgmtAuthenticationCompact_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,2,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	ModeId               string   `protobuf:"bytes,3,opt,name=mode_id,json=modeId" json:"mode_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtAuthenticationCompact_KEYS) Reset()         { *m = RsvpMgmtAuthenticationCompact_KEYS{} }
func (m *RsvpMgmtAuthenticationCompact_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtAuthenticationCompact_KEYS) ProtoMessage()    {}
func (*RsvpMgmtAuthenticationCompact_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_authentication_compact_6e5ab9549bded42b, []int{0}
}
func (m *RsvpMgmtAuthenticationCompact_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtAuthenticationCompact_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtAuthenticationCompact_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtAuthenticationCompact_KEYS.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtAuthenticationCompact_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtAuthenticationCompact_KEYS.Merge(dst, src)
}
func (m *RsvpMgmtAuthenticationCompact_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtAuthenticationCompact_KEYS.Size(m)
}
func (m *RsvpMgmtAuthenticationCompact_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtAuthenticationCompact_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtAuthenticationCompact_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtAuthenticationCompact_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *RsvpMgmtAuthenticationCompact_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RsvpMgmtAuthenticationCompact_KEYS) GetModeId() string {
	if m != nil {
		return m.ModeId
	}
	return ""
}

func (m *RsvpMgmtAuthenticationCompact_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type RsvpMgmtAuthenticationCompact struct {
	// Source address
	SourceAddress string `protobuf:"bytes,50,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	// Destination address
	DestinationAddress string `protobuf:"bytes,51,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// Neighbor address
	NeighborAddress string `protobuf:"bytes,52,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	// Direction
	Direction string `protobuf:"bytes,53,opt,name=direction" json:"direction,omitempty"`
	// Key type
	KeyType string `protobuf:"bytes,54,opt,name=key_type,json=keyType" json:"key_type,omitempty"`
	// Key source
	KeySource string `protobuf:"bytes,55,opt,name=key_source,json=keySource" json:"key_source,omitempty"`
	// Key ID
	KeyId uint64 `protobuf:"varint,56,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	// Key validity
	KeyIdValid           uint32   `protobuf:"varint,57,opt,name=key_id_valid,json=keyIdValid" json:"key_id_valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtAuthenticationCompact) Reset()         { *m = RsvpMgmtAuthenticationCompact{} }
func (m *RsvpMgmtAuthenticationCompact) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtAuthenticationCompact) ProtoMessage()    {}
func (*RsvpMgmtAuthenticationCompact) Descriptor() ([]byte, []int) {
	return fileDescriptor_rsvp_mgmt_authentication_compact_6e5ab9549bded42b, []int{1}
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
	proto.RegisterType((*RsvpMgmtAuthenticationCompact_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.authentication_briefs.authentication_brief.rsvp_mgmt_authentication_compact_KEYS")
	proto.RegisterType((*RsvpMgmtAuthenticationCompact)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.authentication_briefs.authentication_brief.rsvp_mgmt_authentication_compact")
}

func init() {
	proto.RegisterFile("rsvp_mgmt_authentication_compact.proto", fileDescriptor_rsvp_mgmt_authentication_compact_6e5ab9549bded42b)
}

var fileDescriptor_rsvp_mgmt_authentication_compact_6e5ab9549bded42b = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4d, 0x4b, 0x2b, 0x31,
	0x14, 0x86, 0x99, 0xde, 0xde, 0xf6, 0xf6, 0x70, 0xab, 0x12, 0x11, 0x23, 0x28, 0x0c, 0x85, 0x4a,
	0xdd, 0x54, 0xb0, 0x7e, 0x2e, 0x5d, 0xb8, 0x28, 0x82, 0x8b, 0x56, 0x04, 0x57, 0x21, 0x4d, 0x4e,
	0xdb, 0x30, 0xce, 0x64, 0x48, 0xd2, 0xe2, 0xfc, 0x31, 0xb7, 0xfe, 0x35, 0x99, 0x84, 0xa9, 0x22,
	0x95, 0xe2, 0x6e, 0xe6, 0x79, 0x9f, 0x37, 0xe7, 0x04, 0x02, 0xc7, 0xc6, 0x2e, 0x73, 0x96, 0xce,
	0x52, 0xc7, 0xf8, 0xc2, 0xcd, 0x31, 0x73, 0x4a, 0x70, 0xa7, 0x74, 0xc6, 0x84, 0x4e, 0x73, 0x2e,
	0x5c, 0x3f, 0x37, 0xda, 0x69, 0x32, 0x14, 0xca, 0x0a, 0xcd, 0x94, 0xb6, 0xec, 0xd5, 0x30, 0x95,
	0x33, 0xdf, 0xd3, 0x39, 0x9a, 0x7e, 0xf9, 0xd5, 0xff, 0x56, 0x9e, 0x18, 0x85, 0x53, 0xbb, 0x96,
	0x76, 0xde, 0x23, 0xe8, 0x6e, 0x9a, 0xca, 0xee, 0xef, 0x9e, 0xc7, 0xa4, 0x0b, 0x5b, 0x56, 0x2f,
	0x8c, 0x40, 0xc6, 0xa5, 0x34, 0x68, 0x2d, 0x8d, 0xe2, 0xa8, 0xd7, 0x1a, 0xb5, 0x03, 0xbd, 0x0d,
	0x90, 0x9c, 0xc2, 0xae, 0x44, 0xeb, 0x54, 0x16, 0x8e, 0xa8, 0xdc, 0x9a, 0x77, 0xc9, 0x97, 0xa8,
	0x2a, 0xec, 0x43, 0x33, 0xd5, 0x12, 0x99, 0x92, 0xf4, 0x8f, 0x97, 0x1a, 0xe5, 0xef, 0x50, 0x96,
	0x03, 0x55, 0xe6, 0xd0, 0x4c, 0xb9, 0x40, 0x96, 0xf1, 0x14, 0x69, 0x3d, 0x0c, 0x5c, 0xd1, 0x07,
	0x9e, 0x62, 0xe7, 0xad, 0x06, 0xf1, 0xa6, 0x1b, 0xac, 0x59, 0xfe, 0xec, 0x17, 0xcb, 0x0f, 0x7e,
	0x5c, 0xfe, 0x04, 0x76, 0x32, 0x54, 0xb3, 0xf9, 0x44, 0x9b, 0x95, 0x7d, 0xee, 0xed, 0xed, 0x8a,
	0x57, 0xea, 0x21, 0xb4, 0xa4, 0x32, 0x28, 0xca, 0x3a, 0xbd, 0xf0, 0xce, 0x27, 0x20, 0x07, 0xf0,
	0x2f, 0xc1, 0x82, 0xb9, 0x22, 0x47, 0x7a, 0xe9, 0xc3, 0x66, 0x82, 0xc5, 0x63, 0x91, 0x23, 0x39,
	0x02, 0x28, 0xa3, 0xb0, 0x29, 0xbd, 0x0a, 0xcd, 0x04, 0x8b, 0xb1, 0x07, 0x64, 0x0f, 0x1a, 0x65,
	0xac, 0x24, 0xbd, 0x8e, 0xa3, 0x5e, 0x7d, 0xf4, 0x37, 0xc1, 0x62, 0x28, 0x49, 0x0c, 0xff, 0x03,
	0x66, 0x4b, 0xfe, 0xa2, 0x24, 0xbd, 0x89, 0xa3, 0x5e, 0x7b, 0x04, 0x3e, 0x7c, 0x2a, 0xc9, 0xa4,
	0xe1, 0x1f, 0xd3, 0xe0, 0x23, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x6e, 0x3a, 0xf3, 0x76, 0x02, 0x00,
	0x00,
}