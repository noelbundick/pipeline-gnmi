// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_autotun_mesh_config_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_auto_tunnel_mesh_config

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

// Auto-tunnel mesh feature configuration information
type MplsTeAutotunMeshConfigInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeAutotunMeshConfigInfo_KEYS) Reset()         { *m = MplsTeAutotunMeshConfigInfo_KEYS{} }
func (m *MplsTeAutotunMeshConfigInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeAutotunMeshConfigInfo_KEYS) ProtoMessage()    {}
func (*MplsTeAutotunMeshConfigInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_autotun_mesh_config_info_c806e8688cd30edd, []int{0}
}
func (m *MplsTeAutotunMeshConfigInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAutotunMeshConfigInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTeAutotunMeshConfigInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAutotunMeshConfigInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTeAutotunMeshConfigInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAutotunMeshConfigInfo_KEYS.Merge(dst, src)
}
func (m *MplsTeAutotunMeshConfigInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeAutotunMeshConfigInfo_KEYS.Size(m)
}
func (m *MplsTeAutotunMeshConfigInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAutotunMeshConfigInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAutotunMeshConfigInfo_KEYS proto.InternalMessageInfo

type MplsTeAutotunMeshConfigInfo struct {
	// Indicate if auto-tunnel mesh feature is configured
	IsConfigured bool `protobuf:"varint,50,opt,name=is_configured,json=isConfigured" json:"is_configured,omitempty"`
	// Configured value of unused removal timer in seconds
	UnusedRemovalTimeoutConfigured uint32 `protobuf:"varint,51,opt,name=unused_removal_timeout_configured,json=unusedRemovalTimeoutConfigured" json:"unused_removal_timeout_configured,omitempty"`
	// Lower bound of configured tunnel ID range
	MinTunnelId uint32 `protobuf:"varint,52,opt,name=min_tunnel_id,json=minTunnelId" json:"min_tunnel_id,omitempty"`
	// Upper bound of configured tunnel ID range
	MaxTunnelId          uint32   `protobuf:"varint,53,opt,name=max_tunnel_id,json=maxTunnelId" json:"max_tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeAutotunMeshConfigInfo) Reset()         { *m = MplsTeAutotunMeshConfigInfo{} }
func (m *MplsTeAutotunMeshConfigInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeAutotunMeshConfigInfo) ProtoMessage()    {}
func (*MplsTeAutotunMeshConfigInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_autotun_mesh_config_info_c806e8688cd30edd, []int{1}
}
func (m *MplsTeAutotunMeshConfigInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAutotunMeshConfigInfo.Unmarshal(m, b)
}
func (m *MplsTeAutotunMeshConfigInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAutotunMeshConfigInfo.Marshal(b, m, deterministic)
}
func (dst *MplsTeAutotunMeshConfigInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAutotunMeshConfigInfo.Merge(dst, src)
}
func (m *MplsTeAutotunMeshConfigInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeAutotunMeshConfigInfo.Size(m)
}
func (m *MplsTeAutotunMeshConfigInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAutotunMeshConfigInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAutotunMeshConfigInfo proto.InternalMessageInfo

func (m *MplsTeAutotunMeshConfigInfo) GetIsConfigured() bool {
	if m != nil {
		return m.IsConfigured
	}
	return false
}

func (m *MplsTeAutotunMeshConfigInfo) GetUnusedRemovalTimeoutConfigured() uint32 {
	if m != nil {
		return m.UnusedRemovalTimeoutConfigured
	}
	return 0
}

func (m *MplsTeAutotunMeshConfigInfo) GetMinTunnelId() uint32 {
	if m != nil {
		return m.MinTunnelId
	}
	return 0
}

func (m *MplsTeAutotunMeshConfigInfo) GetMaxTunnelId() uint32 {
	if m != nil {
		return m.MaxTunnelId
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeAutotunMeshConfigInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.auto_tunnel.mesh.config.mpls_te_autotun_mesh_config_info_KEYS")
	proto.RegisterType((*MplsTeAutotunMeshConfigInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.auto_tunnel.mesh.config.mpls_te_autotun_mesh_config_info")
}

func init() {
	proto.RegisterFile("mpls_te_autotun_mesh_config_info.proto", fileDescriptor_mpls_te_autotun_mesh_config_info_c806e8688cd30edd)
}

var fileDescriptor_mpls_te_autotun_mesh_config_info_c806e8688cd30edd = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0x05, 0x70, 0x65, 0x41, 0xc8, 0x90, 0x25, 0x53, 0x27, 0x14, 0x82, 0x80, 0x4e, 0x1e, 0x28,
	0x0c, 0xcc, 0x88, 0xa1, 0x62, 0x0b, 0x5d, 0x98, 0x4e, 0x26, 0xb9, 0xc2, 0x49, 0xf1, 0x5d, 0xe4,
	0x3f, 0x28, 0x9f, 0x95, 0x4f, 0x83, 0x1a, 0xbb, 0x28, 0x5b, 0x47, 0x3f, 0xfd, 0xde, 0xb3, 0x6c,
	0x75, 0x67, 0xc7, 0xc1, 0x43, 0x40, 0x30, 0x31, 0x48, 0x88, 0x0c, 0x16, 0xfd, 0x37, 0x74, 0xc2,
	0x7b, 0xfa, 0x02, 0xe2, 0xbd, 0xe8, 0xd1, 0x49, 0x90, 0xea, 0xb9, 0x23, 0xdf, 0x09, 0x90, 0x78,
	0x98, 0x1c, 0x1c, 0x4b, 0x32, 0xa2, 0xd3, 0xf9, 0xa0, 0x0f, 0x0b, 0x10, 0x22, 0x33, 0x0e, 0xfa,
	0xb0, 0xa2, 0xd3, 0x4a, 0x73, 0xaf, 0x6e, 0x4f, 0x5d, 0x02, 0x6f, 0xaf, 0x1f, 0xef, 0xcd, 0x6f,
	0xa1, 0xea, 0x53, 0xb2, 0xba, 0x51, 0x25, 0xf9, 0x9c, 0x44, 0x87, 0xfd, 0xea, 0xa1, 0x2e, 0xd6,
	0xe7, 0xed, 0x25, 0xf9, 0x97, 0xff, 0xac, 0xda, 0xaa, 0xeb, 0xc8, 0xd1, 0x63, 0x0f, 0x0e, 0xad,
	0xfc, 0x98, 0x01, 0x02, 0x59, 0x94, 0x18, 0x96, 0xc5, 0x4d, 0x5d, 0xac, 0xcb, 0xf6, 0x2a, 0xc1,
	0x36, 0xb9, 0x5d, 0x62, 0x8b, 0xa9, 0x46, 0x95, 0x96, 0x38, 0xbf, 0x0b, 0xa8, 0x5f, 0x3d, 0xce,
	0xb5, 0x0b, 0x4b, 0xbc, 0x9b, 0xb3, 0x6d, 0x32, 0x66, 0x5a, 0x98, 0xa7, 0x6c, 0xcc, 0x74, 0x34,
	0x9f, 0x67, 0xf3, 0x3f, 0x6e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x76, 0xdd, 0x9b, 0xcc, 0x71,
	0x01, 0x00, 0x00,
}
