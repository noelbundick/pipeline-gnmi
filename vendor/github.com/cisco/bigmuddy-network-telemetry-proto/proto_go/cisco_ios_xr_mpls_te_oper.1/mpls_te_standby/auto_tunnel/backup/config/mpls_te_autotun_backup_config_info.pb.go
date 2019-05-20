// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_autotun_backup_config_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_auto_tunnel_backup_config

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
	return fileDescriptor_mpls_te_autotun_backup_config_info_89f5043223ca2636, []int{0}
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
	return fileDescriptor_mpls_te_autotun_backup_config_info_89f5043223ca2636, []int{1}
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
	proto.RegisterType((*MplsTeAutotunBackupConfigInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.auto_tunnel.backup.config.mpls_te_autotun_backup_config_info_KEYS")
	proto.RegisterType((*MplsTeAutotunBackupConfigInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.auto_tunnel.backup.config.mpls_te_autotun_backup_config_info")
}

func init() {
	proto.RegisterFile("mpls_te_autotun_backup_config_info.proto", fileDescriptor_mpls_te_autotun_backup_config_info_89f5043223ca2636)
}

var fileDescriptor_mpls_te_autotun_backup_config_info_89f5043223ca2636 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbf, 0x4e, 0x84, 0x40,
	0x10, 0xc6, 0x73, 0x16, 0xc6, 0xac, 0xa2, 0x09, 0x15, 0x95, 0x39, 0xb1, 0x38, 0x6c, 0x28, 0x3c,
	0xf5, 0x05, 0x88, 0xc5, 0xc5, 0x0e, 0xaf, 0xb1, 0x9a, 0x2c, 0xcb, 0x62, 0x26, 0xc2, 0x0c, 0xd9,
	0x3f, 0x06, 0x1f, 0xc0, 0xf7, 0xbe, 0x1c, 0x0b, 0x17, 0xba, 0x2b, 0xe7, 0x97, 0xdf, 0xf7, 0x7d,
	0xc9, 0x88, 0xac, 0xeb, 0x5b, 0x0b, 0x4e, 0x83, 0xf4, 0x8e, 0x9d, 0x27, 0xa8, 0xa4, 0xfa, 0xf1,
	0x3d, 0x28, 0xa6, 0x06, 0xbf, 0x01, 0xa9, 0xe1, 0xbc, 0x37, 0xec, 0x38, 0x2e, 0x14, 0x5a, 0xc5,
	0x80, 0x6c, 0x61, 0x30, 0x30, 0xc7, 0xb8, 0xd7, 0x26, 0x9f, 0x0f, 0xeb, 0x24, 0xd5, 0xd5, 0x5f,
	0x7e, 0xec, 0x02, 0xe7, 0x89, 0x74, 0x9b, 0x87, 0xbe, 0x3c, 0xf4, 0xa5, 0x4f, 0x62, 0x73, 0x7e,
	0x10, 0x3e, 0xde, 0xbf, 0x3e, 0xd3, 0xff, 0x0b, 0x91, 0x9e, 0x77, 0xe3, 0x47, 0x11, 0xa1, 0x9d,
	0x88, 0x37, 0xba, 0x4e, 0x9e, 0xd7, 0xab, 0xec, 0xaa, 0xbc, 0x41, 0x5b, 0x9c, 0x58, 0xbc, 0x11,
	0x77, 0x48, 0x4e, 0x9b, 0x46, 0x2a, 0x0d, 0x8a, 0x3d, 0xb9, 0x64, 0xbb, 0x5e, 0x65, 0x51, 0x79,
	0x7b, 0xc2, 0xc5, 0x91, 0xc6, 0x3b, 0xf1, 0xe0, 0xc9, 0x5b, 0x5d, 0x83, 0xd1, 0x1d, 0xff, 0xca,
	0x16, 0x1c, 0x76, 0x9a, 0xbd, 0x5b, 0x2e, 0xbc, 0x8c, 0xd1, 0xfb, 0x20, 0x96, 0xc1, 0xdb, 0x07,
	0x6d, 0xb1, 0x99, 0x8a, 0xa8, 0x43, 0x9a, 0xde, 0x00, 0x58, 0x27, 0xaf, 0x63, 0xec, 0xba, 0x43,
	0xda, 0x8f, 0x6c, 0x17, 0x1c, 0x39, 0x2c, 0x9c, 0xb7, 0xc9, 0x91, 0xc3, 0xec, 0x54, 0x97, 0xe3,
	0xfb, 0xb7, 0x87, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0xc6, 0x93, 0x15, 0xaa, 0x01, 0x00, 0x00,
}