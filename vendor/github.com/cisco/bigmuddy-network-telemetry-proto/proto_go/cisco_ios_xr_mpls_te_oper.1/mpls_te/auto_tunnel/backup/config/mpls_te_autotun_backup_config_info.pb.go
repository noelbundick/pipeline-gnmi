// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_autotun_backup_config_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_auto_tunnel_backup_config

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

// Auto-tunnel backup feature configuration information
type MplsTeAutotunBackupConfigInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeAutotunBackupConfigInfo_KEYS) Reset()         { *m = MplsTeAutotunBackupConfigInfo_KEYS{} }
func (m *MplsTeAutotunBackupConfigInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeAutotunBackupConfigInfo_KEYS) ProtoMessage()    {}
func (*MplsTeAutotunBackupConfigInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_autotun_backup_config_info_4c03f5045ecb1659, []int{0}
}
func (m *MplsTeAutotunBackupConfigInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAutotunBackupConfigInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTeAutotunBackupConfigInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAutotunBackupConfigInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTeAutotunBackupConfigInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAutotunBackupConfigInfo_KEYS.Merge(dst, src)
}
func (m *MplsTeAutotunBackupConfigInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeAutotunBackupConfigInfo_KEYS.Size(m)
}
func (m *MplsTeAutotunBackupConfigInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAutotunBackupConfigInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAutotunBackupConfigInfo_KEYS proto.InternalMessageInfo

type MplsTeAutotunBackupConfigInfo struct {
	// Indicate if auto-tunnel bacukp feature is configured
	IsConfigured bool `protobuf:"varint,50,opt,name=is_configured,json=isConfigured" json:"is_configured,omitempty"`
	// Number of TE link interfaces with auto-tunnel backup configured
	InterfaceCount uint32 `protobuf:"varint,51,opt,name=interface_count,json=interfaceCount" json:"interface_count,omitempty"`
	// Configured value of unused removal timer in seconds
	UnusedRemovalTimeoutConfigured uint32 `protobuf:"varint,52,opt,name=unused_removal_timeout_configured,json=unusedRemovalTimeoutConfigured" json:"unused_removal_timeout_configured,omitempty"`
	// Lower bound of configured tunnel ID range
	MinTunnelId uint32 `protobuf:"varint,53,opt,name=min_tunnel_id,json=minTunnelId" json:"min_tunnel_id,omitempty"`
	// Upper bound of configured tunnel ID range
	MaxTunnelId          uint32   `protobuf:"varint,54,opt,name=max_tunnel_id,json=maxTunnelId" json:"max_tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeAutotunBackupConfigInfo) Reset()         { *m = MplsTeAutotunBackupConfigInfo{} }
func (m *MplsTeAutotunBackupConfigInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeAutotunBackupConfigInfo) ProtoMessage()    {}
func (*MplsTeAutotunBackupConfigInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_autotun_backup_config_info_4c03f5045ecb1659, []int{1}
}
func (m *MplsTeAutotunBackupConfigInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAutotunBackupConfigInfo.Unmarshal(m, b)
}
func (m *MplsTeAutotunBackupConfigInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAutotunBackupConfigInfo.Marshal(b, m, deterministic)
}
func (dst *MplsTeAutotunBackupConfigInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAutotunBackupConfigInfo.Merge(dst, src)
}
func (m *MplsTeAutotunBackupConfigInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeAutotunBackupConfigInfo.Size(m)
}
func (m *MplsTeAutotunBackupConfigInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAutotunBackupConfigInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAutotunBackupConfigInfo proto.InternalMessageInfo

func (m *MplsTeAutotunBackupConfigInfo) GetIsConfigured() bool {
	if m != nil {
		return m.IsConfigured
	}
	return false
}

func (m *MplsTeAutotunBackupConfigInfo) GetInterfaceCount() uint32 {
	if m != nil {
		return m.InterfaceCount
	}
	return 0
}

func (m *MplsTeAutotunBackupConfigInfo) GetUnusedRemovalTimeoutConfigured() uint32 {
	if m != nil {
		return m.UnusedRemovalTimeoutConfigured
	}
	return 0
}

func (m *MplsTeAutotunBackupConfigInfo) GetMinTunnelId() uint32 {
	if m != nil {
		return m.MinTunnelId
	}
	return 0
}

func (m *MplsTeAutotunBackupConfigInfo) GetMaxTunnelId() uint32 {
	if m != nil {
		return m.MaxTunnelId
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeAutotunBackupConfigInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.auto_tunnel.backup.config.mpls_te_autotun_backup_config_info_KEYS")
	proto.RegisterType((*MplsTeAutotunBackupConfigInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.auto_tunnel.backup.config.mpls_te_autotun_backup_config_info")
}

func init() {
	proto.RegisterFile("mpls_te_autotun_backup_config_info.proto", fileDescriptor_mpls_te_autotun_backup_config_info_4c03f5045ecb1659)
}

var fileDescriptor_mpls_te_autotun_backup_config_info_4c03f5045ecb1659 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbd, 0x4e, 0xeb, 0x40,
	0x10, 0x85, 0x95, 0x5b, 0x5c, 0xa1, 0x05, 0x83, 0xe4, 0xca, 0x15, 0x0a, 0xa6, 0x88, 0x69, 0xb6,
	0x20, 0x40, 0x43, 0x19, 0x51, 0x44, 0x74, 0x26, 0x0d, 0xd5, 0x68, 0xb3, 0x5e, 0xa3, 0x11, 0xde,
	0x19, 0x6b, 0x7f, 0x90, 0x5f, 0x80, 0xf7, 0x46, 0xf1, 0x3a, 0x91, 0xbb, 0x94, 0xf3, 0xe9, 0x3b,
	0xe7, 0x48, 0x23, 0x2a, 0xdb, 0x77, 0x1e, 0x82, 0x01, 0x15, 0x03, 0x87, 0x48, 0xb0, 0x57, 0xfa,
	0x3b, 0xf6, 0xa0, 0x99, 0x5a, 0xfc, 0x02, 0xa4, 0x96, 0x65, 0xef, 0x38, 0x70, 0xfe, 0xaa, 0xd1,
	0x6b, 0x06, 0x64, 0x0f, 0x83, 0x83, 0x63, 0x8c, 0x7b, 0xe3, 0xe4, 0x74, 0xc8, 0x43, 0x07, 0x84,
	0x48, 0x64, 0x3a, 0x99, 0x7a, 0x64, 0xea, 0x29, 0x1f, 0xc4, 0xea, 0xfc, 0x10, 0xbc, 0xbf, 0x7d,
	0x7e, 0x94, 0xbf, 0xff, 0x44, 0x79, 0xde, 0xcd, 0xef, 0x45, 0x86, 0x7e, 0x22, 0xd1, 0x99, 0xa6,
	0x78, 0x5c, 0x2e, 0xaa, 0x8b, 0xfa, 0x0a, 0xfd, 0xe6, 0xc4, 0xf2, 0x95, 0xb8, 0x41, 0x0a, 0xc6,
	0xb5, 0x4a, 0x1b, 0xd0, 0x1c, 0x29, 0x14, 0xeb, 0xe5, 0xa2, 0xca, 0xea, 0xeb, 0x13, 0xde, 0x1c,
	0x68, 0xbe, 0x15, 0x77, 0x91, 0xa2, 0x37, 0x0d, 0x38, 0x63, 0xf9, 0x47, 0x75, 0x10, 0xd0, 0x1a,
	0x8e, 0x61, 0xbe, 0xf0, 0x34, 0x46, 0x6f, 0x93, 0x58, 0x27, 0x6f, 0x97, 0xb4, 0xd9, 0x66, 0x29,
	0x32, 0x8b, 0x34, 0xbd, 0x01, 0xb0, 0x29, 0x9e, 0xc7, 0xd8, 0xa5, 0x45, 0xda, 0x8d, 0x6c, 0x9b,
	0x1c, 0x35, 0xcc, 0x9c, 0x97, 0xc9, 0x51, 0xc3, 0xd1, 0xd9, 0xff, 0x1f, 0xdf, 0xbe, 0xfe, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x02, 0xf7, 0x78, 0xd1, 0xa2, 0x01, 0x00, 0x00,
}
