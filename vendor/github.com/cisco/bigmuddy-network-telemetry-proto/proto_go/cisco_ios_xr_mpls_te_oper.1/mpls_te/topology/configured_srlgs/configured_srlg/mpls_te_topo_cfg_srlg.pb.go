// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_topo_cfg_srlg.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_topology_configured_srlgs_configured_srlg

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

// A configured SRLG entry
type MplsTeTopoCfgSrlg_KEYS struct {
	SrlgNumber           uint32   `protobuf:"varint,1,opt,name=srlg_number,json=srlgNumber" json:"srlg_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTopoCfgSrlg_KEYS) Reset()         { *m = MplsTeTopoCfgSrlg_KEYS{} }
func (m *MplsTeTopoCfgSrlg_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeTopoCfgSrlg_KEYS) ProtoMessage()    {}
func (*MplsTeTopoCfgSrlg_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_topo_cfg_srlg_43dc56dd2c3cf2cd, []int{0}
}
func (m *MplsTeTopoCfgSrlg_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTopoCfgSrlg_KEYS.Unmarshal(m, b)
}
func (m *MplsTeTopoCfgSrlg_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTopoCfgSrlg_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTeTopoCfgSrlg_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTopoCfgSrlg_KEYS.Merge(dst, src)
}
func (m *MplsTeTopoCfgSrlg_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeTopoCfgSrlg_KEYS.Size(m)
}
func (m *MplsTeTopoCfgSrlg_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTopoCfgSrlg_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTopoCfgSrlg_KEYS proto.InternalMessageInfo

func (m *MplsTeTopoCfgSrlg_KEYS) GetSrlgNumber() uint32 {
	if m != nil {
		return m.SrlgNumber
	}
	return 0
}

type MplsTeTopoCfgSrlg struct {
	// SRLG number
	SrlgNumber uint32 `protobuf:"varint,50,opt,name=srlg_number,json=srlgNumber" json:"srlg_number,omitempty"`
	// SRLG name
	SrlgName string `protobuf:"bytes,51,opt,name=srlg_name,json=srlgName" json:"srlg_name,omitempty"`
	// The admin weight that is added to the link if the SRLG is shared with the protected link
	AdminWeight uint32 `protobuf:"varint,52,opt,name=admin_weight,json=adminWeight" json:"admin_weight,omitempty"`
	// Set to TRUE if the admin weight is explicitely configured
	IsAdminWeightConfigured bool `protobuf:"varint,53,opt,name=is_admin_weight_configured,json=isAdminWeightConfigured" json:"is_admin_weight_configured,omitempty"`
	// Link associated with the SRLG
	SrlgLink             []*MplsTeTopoCfgSrlgLink `protobuf:"bytes,54,rep,name=srlg_link,json=srlgLink" json:"srlg_link,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MplsTeTopoCfgSrlg) Reset()         { *m = MplsTeTopoCfgSrlg{} }
func (m *MplsTeTopoCfgSrlg) String() string { return proto.CompactTextString(m) }
func (*MplsTeTopoCfgSrlg) ProtoMessage()    {}
func (*MplsTeTopoCfgSrlg) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_topo_cfg_srlg_43dc56dd2c3cf2cd, []int{1}
}
func (m *MplsTeTopoCfgSrlg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTopoCfgSrlg.Unmarshal(m, b)
}
func (m *MplsTeTopoCfgSrlg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTopoCfgSrlg.Marshal(b, m, deterministic)
}
func (dst *MplsTeTopoCfgSrlg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTopoCfgSrlg.Merge(dst, src)
}
func (m *MplsTeTopoCfgSrlg) XXX_Size() int {
	return xxx_messageInfo_MplsTeTopoCfgSrlg.Size(m)
}
func (m *MplsTeTopoCfgSrlg) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTopoCfgSrlg.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTopoCfgSrlg proto.InternalMessageInfo

func (m *MplsTeTopoCfgSrlg) GetSrlgNumber() uint32 {
	if m != nil {
		return m.SrlgNumber
	}
	return 0
}

func (m *MplsTeTopoCfgSrlg) GetSrlgName() string {
	if m != nil {
		return m.SrlgName
	}
	return ""
}

func (m *MplsTeTopoCfgSrlg) GetAdminWeight() uint32 {
	if m != nil {
		return m.AdminWeight
	}
	return 0
}

func (m *MplsTeTopoCfgSrlg) GetIsAdminWeightConfigured() bool {
	if m != nil {
		return m.IsAdminWeightConfigured
	}
	return false
}

func (m *MplsTeTopoCfgSrlg) GetSrlgLink() []*MplsTeTopoCfgSrlgLink {
	if m != nil {
		return m.SrlgLink
	}
	return nil
}

// A link associated with configured SRLG
type MplsTeTopoCfgSrlgLink struct {
	// Local address
	LocalAddress string `protobuf:"bytes,1,opt,name=local_address,json=localAddress" json:"local_address,omitempty"`
	// Remote address
	RemoteAddress        string   `protobuf:"bytes,2,opt,name=remote_address,json=remoteAddress" json:"remote_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTopoCfgSrlgLink) Reset()         { *m = MplsTeTopoCfgSrlgLink{} }
func (m *MplsTeTopoCfgSrlgLink) String() string { return proto.CompactTextString(m) }
func (*MplsTeTopoCfgSrlgLink) ProtoMessage()    {}
func (*MplsTeTopoCfgSrlgLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_topo_cfg_srlg_43dc56dd2c3cf2cd, []int{2}
}
func (m *MplsTeTopoCfgSrlgLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTopoCfgSrlgLink.Unmarshal(m, b)
}
func (m *MplsTeTopoCfgSrlgLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTopoCfgSrlgLink.Marshal(b, m, deterministic)
}
func (dst *MplsTeTopoCfgSrlgLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTopoCfgSrlgLink.Merge(dst, src)
}
func (m *MplsTeTopoCfgSrlgLink) XXX_Size() int {
	return xxx_messageInfo_MplsTeTopoCfgSrlgLink.Size(m)
}
func (m *MplsTeTopoCfgSrlgLink) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTopoCfgSrlgLink.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTopoCfgSrlgLink proto.InternalMessageInfo

func (m *MplsTeTopoCfgSrlgLink) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *MplsTeTopoCfgSrlgLink) GetRemoteAddress() string {
	if m != nil {
		return m.RemoteAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsTeTopoCfgSrlg_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.topology.configured_srlgs.configured_srlg.mpls_te_topo_cfg_srlg_KEYS")
	proto.RegisterType((*MplsTeTopoCfgSrlg)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.topology.configured_srlgs.configured_srlg.mpls_te_topo_cfg_srlg")
	proto.RegisterType((*MplsTeTopoCfgSrlgLink)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.topology.configured_srlgs.configured_srlg.mpls_te_topo_cfg_srlg_link")
}

func init() {
	proto.RegisterFile("mpls_te_topo_cfg_srlg.proto", fileDescriptor_mpls_te_topo_cfg_srlg_43dc56dd2c3cf2cd)
}

var fileDescriptor_mpls_te_topo_cfg_srlg_43dc56dd2c3cf2cd = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x49, 0x05, 0x69, 0x2e, 0x8d, 0xc3, 0x81, 0x18, 0xda, 0xc1, 0x18, 0x11, 0x32, 0xdd,
	0xd0, 0xaa, 0x8b, 0x38, 0x14, 0x71, 0xaa, 0x38, 0xc4, 0x41, 0x9c, 0x1e, 0x69, 0x72, 0xbd, 0x1e,
	0xbd, 0xcb, 0x85, 0xbb, 0x14, 0x75, 0xf7, 0x63, 0xf9, 0xe1, 0x24, 0x97, 0xd8, 0x4a, 0x88, 0x5b,
	0xc7, 0xfb, 0xbd, 0xdf, 0x3f, 0x8f, 0xf7, 0x27, 0x68, 0x22, 0x4b, 0x61, 0xa0, 0xa2, 0x50, 0xa9,
	0x52, 0x41, 0xb6, 0x62, 0x60, 0xb4, 0x60, 0xa4, 0xd4, 0xaa, 0x52, 0x78, 0x91, 0x71, 0x93, 0x29,
	0xe0, 0xca, 0xc0, 0x87, 0x86, 0x5f, 0x53, 0x95, 0x54, 0x93, 0xf6, 0x41, 0xea, 0x98, 0x50, 0xec,
	0x93, 0x64, 0xaa, 0x58, 0x71, 0xb6, 0xd5, 0x34, 0xb7, 0x5f, 0x30, 0x5d, 0x10, 0xdd, 0xa3, 0x71,
	0xef, 0x2e, 0x58, 0x3c, 0xbe, 0xbd, 0xe0, 0x73, 0xe4, 0xd9, 0x47, 0xb1, 0x95, 0x4b, 0xaa, 0x03,
	0x27, 0x74, 0x62, 0x3f, 0x41, 0x35, 0x7a, 0xb6, 0x24, 0xfa, 0x1e, 0xa0, 0xd3, 0xde, 0x7c, 0x37,
	0x3a, 0xed, 0x46, 0xf1, 0x04, 0xb9, 0x8d, 0x90, 0x4a, 0x1a, 0xcc, 0x42, 0x27, 0x76, 0x93, 0xa1,
	0x1d, 0xa7, 0x92, 0xe2, 0x0b, 0x34, 0x4a, 0x73, 0xc9, 0x0b, 0x78, 0xa7, 0x9c, 0xad, 0xab, 0xe0,
	0xda, 0xc6, 0x3d, 0xcb, 0x5e, 0x2d, 0xc2, 0x77, 0x68, 0xcc, 0x0d, 0xfc, 0xb5, 0x60, 0x7f, 0x5c,
	0x70, 0x13, 0x3a, 0xf1, 0x30, 0x39, 0xe3, 0x66, 0xbe, 0x8f, 0x3c, 0xec, 0xc6, 0xf8, 0xcb, 0x69,
	0xb7, 0x0b, 0x5e, 0x6c, 0x82, 0xdb, 0xf0, 0x28, 0xf6, 0xa6, 0x8c, 0x1c, 0xb0, 0x58, 0xd2, 0xdf,
	0x6a, 0xbd, 0xae, 0x39, 0xf3, 0x89, 0x17, 0x9b, 0x68, 0xfd, 0x5f, 0xfb, 0xb5, 0x87, 0x2f, 0x91,
	0x2f, 0x54, 0x96, 0x0a, 0x48, 0xf3, 0x5c, 0x53, 0x63, 0x6c, 0xff, 0x6e, 0x32, 0xb2, 0x70, 0xde,
	0x30, 0x7c, 0x85, 0x4e, 0x34, 0x95, 0xaa, 0xa2, 0x3b, 0x6b, 0x60, 0x2d, 0xbf, 0xa1, 0xad, 0xb6,
	0x3c, 0xb6, 0xff, 0xce, 0xec, 0x27, 0x00, 0x00, 0xff, 0xff, 0x40, 0xc0, 0xe3, 0x6f, 0x5a, 0x02,
	0x00, 0x00,
}
