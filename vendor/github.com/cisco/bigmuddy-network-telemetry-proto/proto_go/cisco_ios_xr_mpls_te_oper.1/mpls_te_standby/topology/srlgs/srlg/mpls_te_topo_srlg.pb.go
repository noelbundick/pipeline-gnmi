// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_topo_srlg.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_topology_srlgs_srlg

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

// A SRLG entry in the MPLS TE topology
type MplsTeTopoSrlg_KEYS struct {
	SrlgNumber           uint32   `protobuf:"varint,1,opt,name=srlg_number,json=srlgNumber" json:"srlg_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTopoSrlg_KEYS) Reset()         { *m = MplsTeTopoSrlg_KEYS{} }
func (m *MplsTeTopoSrlg_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeTopoSrlg_KEYS) ProtoMessage()    {}
func (*MplsTeTopoSrlg_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_topo_srlg_8f6a700e427582ab, []int{0}
}
func (m *MplsTeTopoSrlg_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTopoSrlg_KEYS.Unmarshal(m, b)
}
func (m *MplsTeTopoSrlg_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTopoSrlg_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTeTopoSrlg_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTopoSrlg_KEYS.Merge(dst, src)
}
func (m *MplsTeTopoSrlg_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeTopoSrlg_KEYS.Size(m)
}
func (m *MplsTeTopoSrlg_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTopoSrlg_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTopoSrlg_KEYS proto.InternalMessageInfo

func (m *MplsTeTopoSrlg_KEYS) GetSrlgNumber() uint32 {
	if m != nil {
		return m.SrlgNumber
	}
	return 0
}

type MplsTeTopoSrlg struct {
	// SRLG name
	SrlgName string `protobuf:"bytes,50,opt,name=srlg_name,json=srlgName" json:"srlg_name,omitempty"`
	// The admin weight that is added to the link if the SRLG is shared with the protected link
	AdminWeight uint32 `protobuf:"varint,51,opt,name=admin_weight,json=adminWeight" json:"admin_weight,omitempty"`
	// Topology areas in this SRLG
	SrlgAreas            []*MplsTeTopoSrlgArea `protobuf:"bytes,52,rep,name=srlg_areas,json=srlgAreas" json:"srlg_areas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MplsTeTopoSrlg) Reset()         { *m = MplsTeTopoSrlg{} }
func (m *MplsTeTopoSrlg) String() string { return proto.CompactTextString(m) }
func (*MplsTeTopoSrlg) ProtoMessage()    {}
func (*MplsTeTopoSrlg) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_topo_srlg_8f6a700e427582ab, []int{1}
}
func (m *MplsTeTopoSrlg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTopoSrlg.Unmarshal(m, b)
}
func (m *MplsTeTopoSrlg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTopoSrlg.Marshal(b, m, deterministic)
}
func (dst *MplsTeTopoSrlg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTopoSrlg.Merge(dst, src)
}
func (m *MplsTeTopoSrlg) XXX_Size() int {
	return xxx_messageInfo_MplsTeTopoSrlg.Size(m)
}
func (m *MplsTeTopoSrlg) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTopoSrlg.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTopoSrlg proto.InternalMessageInfo

func (m *MplsTeTopoSrlg) GetSrlgName() string {
	if m != nil {
		return m.SrlgName
	}
	return ""
}

func (m *MplsTeTopoSrlg) GetAdminWeight() uint32 {
	if m != nil {
		return m.AdminWeight
	}
	return 0
}

func (m *MplsTeTopoSrlg) GetSrlgAreas() []*MplsTeTopoSrlgArea {
	if m != nil {
		return m.SrlgAreas
	}
	return nil
}

// TE IPv4 unnumbered address type
type TeAddrTypeIpv4Unnum struct {
	// IPv4 router ID
	RouterId string `protobuf:"bytes,1,opt,name=router_id,json=routerId" json:"router_id,omitempty"`
	// Interface index
	InterfaceIndex       uint32   `protobuf:"varint,2,opt,name=interface_index,json=interfaceIndex" json:"interface_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeAddrTypeIpv4Unnum) Reset()         { *m = TeAddrTypeIpv4Unnum{} }
func (m *TeAddrTypeIpv4Unnum) String() string { return proto.CompactTextString(m) }
func (*TeAddrTypeIpv4Unnum) ProtoMessage()    {}
func (*TeAddrTypeIpv4Unnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_topo_srlg_8f6a700e427582ab, []int{2}
}
func (m *TeAddrTypeIpv4Unnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeAddrTypeIpv4Unnum.Unmarshal(m, b)
}
func (m *TeAddrTypeIpv4Unnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeAddrTypeIpv4Unnum.Marshal(b, m, deterministic)
}
func (dst *TeAddrTypeIpv4Unnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeAddrTypeIpv4Unnum.Merge(dst, src)
}
func (m *TeAddrTypeIpv4Unnum) XXX_Size() int {
	return xxx_messageInfo_TeAddrTypeIpv4Unnum.Size(m)
}
func (m *TeAddrTypeIpv4Unnum) XXX_DiscardUnknown() {
	xxx_messageInfo_TeAddrTypeIpv4Unnum.DiscardUnknown(m)
}

var xxx_messageInfo_TeAddrTypeIpv4Unnum proto.InternalMessageInfo

func (m *TeAddrTypeIpv4Unnum) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *TeAddrTypeIpv4Unnum) GetInterfaceIndex() uint32 {
	if m != nil {
		return m.InterfaceIndex
	}
	return 0
}

type Addr struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// IPv4 address
	Ipv4Address string `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address" json:"ipv4_address,omitempty"`
	// IPv4 unnumbered address
	Ipv4UnnumberedAddress *TeAddrTypeIpv4Unnum `protobuf:"bytes,3,opt,name=ipv4_unnumbered_address,json=ipv4UnnumberedAddress" json:"ipv4_unnumbered_address,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}             `json:"-"`
	XXX_unrecognized      []byte               `json:"-"`
	XXX_sizecache         int32                `json:"-"`
}

func (m *Addr) Reset()         { *m = Addr{} }
func (m *Addr) String() string { return proto.CompactTextString(m) }
func (*Addr) ProtoMessage()    {}
func (*Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_topo_srlg_8f6a700e427582ab, []int{3}
}
func (m *Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addr.Unmarshal(m, b)
}
func (m *Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addr.Marshal(b, m, deterministic)
}
func (dst *Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addr.Merge(dst, src)
}
func (m *Addr) XXX_Size() int {
	return xxx_messageInfo_Addr.Size(m)
}
func (m *Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_Addr.DiscardUnknown(m)
}

var xxx_messageInfo_Addr proto.InternalMessageInfo

func (m *Addr) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Addr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *Addr) GetIpv4UnnumberedAddress() *TeAddrTypeIpv4Unnum {
	if m != nil {
		return m.Ipv4UnnumberedAddress
	}
	return nil
}

type TeAddrT_ struct {
	// TE Address
	TeAddr               *Addr    `protobuf:"bytes,1,opt,name=te_addr,json=teAddr" json:"te_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeAddrT_) Reset()         { *m = TeAddrT_{} }
func (m *TeAddrT_) String() string { return proto.CompactTextString(m) }
func (*TeAddrT_) ProtoMessage()    {}
func (*TeAddrT_) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_topo_srlg_8f6a700e427582ab, []int{4}
}
func (m *TeAddrT_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeAddrT_.Unmarshal(m, b)
}
func (m *TeAddrT_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeAddrT_.Marshal(b, m, deterministic)
}
func (dst *TeAddrT_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeAddrT_.Merge(dst, src)
}
func (m *TeAddrT_) XXX_Size() int {
	return xxx_messageInfo_TeAddrT_.Size(m)
}
func (m *TeAddrT_) XXX_DiscardUnknown() {
	xxx_messageInfo_TeAddrT_.DiscardUnknown(m)
}

var xxx_messageInfo_TeAddrT_ proto.InternalMessageInfo

func (m *TeAddrT_) GetTeAddr() *Addr {
	if m != nil {
		return m.TeAddr
	}
	return nil
}

// SRLG link address data
type MplsTeTopoSrlgLink struct {
	// Link address
	LinkAddress string `protobuf:"bytes,1,opt,name=link_address,json=linkAddress" json:"link_address,omitempty"`
	// Link address
	LinkAddressGeneric *TeAddrT_ `protobuf:"bytes,2,opt,name=link_address_generic,json=linkAddressGeneric" json:"link_address_generic,omitempty"`
	// Link TE router-id
	TeRouterId           string   `protobuf:"bytes,3,opt,name=te_router_id,json=teRouterId" json:"te_router_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTopoSrlgLink) Reset()         { *m = MplsTeTopoSrlgLink{} }
func (m *MplsTeTopoSrlgLink) String() string { return proto.CompactTextString(m) }
func (*MplsTeTopoSrlgLink) ProtoMessage()    {}
func (*MplsTeTopoSrlgLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_topo_srlg_8f6a700e427582ab, []int{5}
}
func (m *MplsTeTopoSrlgLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTopoSrlgLink.Unmarshal(m, b)
}
func (m *MplsTeTopoSrlgLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTopoSrlgLink.Marshal(b, m, deterministic)
}
func (dst *MplsTeTopoSrlgLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTopoSrlgLink.Merge(dst, src)
}
func (m *MplsTeTopoSrlgLink) XXX_Size() int {
	return xxx_messageInfo_MplsTeTopoSrlgLink.Size(m)
}
func (m *MplsTeTopoSrlgLink) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTopoSrlgLink.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTopoSrlgLink proto.InternalMessageInfo

func (m *MplsTeTopoSrlgLink) GetLinkAddress() string {
	if m != nil {
		return m.LinkAddress
	}
	return ""
}

func (m *MplsTeTopoSrlgLink) GetLinkAddressGeneric() *TeAddrT_ {
	if m != nil {
		return m.LinkAddressGeneric
	}
	return nil
}

func (m *MplsTeTopoSrlgLink) GetTeRouterId() string {
	if m != nil {
		return m.TeRouterId
	}
	return ""
}

// A MPLS TE topology link in a SRLG
type MplsTeTopoSrlgArea struct {
	// IGP type
	IgpType string `protobuf:"bytes,1,opt,name=igp_type,json=igpType" json:"igp_type,omitempty"`
	// IGP Instance name
	IgpInstance string `protobuf:"bytes,2,opt,name=igp_instance,json=igpInstance" json:"igp_instance,omitempty"`
	// IGP Area ID
	IgpArea uint32 `protobuf:"varint,3,opt,name=igp_area,json=igpArea" json:"igp_area,omitempty"`
	// IGP-area format
	IgpAreaFormat string `protobuf:"bytes,4,opt,name=igp_area_format,json=igpAreaFormat" json:"igp_area_format,omitempty"`
	// Links in this SRLG/area
	SrlgLinks            []*MplsTeTopoSrlgLink `protobuf:"bytes,5,rep,name=srlg_links,json=srlgLinks" json:"srlg_links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MplsTeTopoSrlgArea) Reset()         { *m = MplsTeTopoSrlgArea{} }
func (m *MplsTeTopoSrlgArea) String() string { return proto.CompactTextString(m) }
func (*MplsTeTopoSrlgArea) ProtoMessage()    {}
func (*MplsTeTopoSrlgArea) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_topo_srlg_8f6a700e427582ab, []int{6}
}
func (m *MplsTeTopoSrlgArea) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTopoSrlgArea.Unmarshal(m, b)
}
func (m *MplsTeTopoSrlgArea) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTopoSrlgArea.Marshal(b, m, deterministic)
}
func (dst *MplsTeTopoSrlgArea) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTopoSrlgArea.Merge(dst, src)
}
func (m *MplsTeTopoSrlgArea) XXX_Size() int {
	return xxx_messageInfo_MplsTeTopoSrlgArea.Size(m)
}
func (m *MplsTeTopoSrlgArea) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTopoSrlgArea.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTopoSrlgArea proto.InternalMessageInfo

func (m *MplsTeTopoSrlgArea) GetIgpType() string {
	if m != nil {
		return m.IgpType
	}
	return ""
}

func (m *MplsTeTopoSrlgArea) GetIgpInstance() string {
	if m != nil {
		return m.IgpInstance
	}
	return ""
}

func (m *MplsTeTopoSrlgArea) GetIgpArea() uint32 {
	if m != nil {
		return m.IgpArea
	}
	return 0
}

func (m *MplsTeTopoSrlgArea) GetIgpAreaFormat() string {
	if m != nil {
		return m.IgpAreaFormat
	}
	return ""
}

func (m *MplsTeTopoSrlgArea) GetSrlgLinks() []*MplsTeTopoSrlgLink {
	if m != nil {
		return m.SrlgLinks
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsTeTopoSrlg_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.topology.srlgs.srlg.mpls_te_topo_srlg_KEYS")
	proto.RegisterType((*MplsTeTopoSrlg)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.topology.srlgs.srlg.mpls_te_topo_srlg")
	proto.RegisterType((*TeAddrTypeIpv4Unnum)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.topology.srlgs.srlg.te_addr_type_ipv4_unnum")
	proto.RegisterType((*Addr)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.topology.srlgs.srlg.addr")
	proto.RegisterType((*TeAddrT_)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.topology.srlgs.srlg.te_addr_t_")
	proto.RegisterType((*MplsTeTopoSrlgLink)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.topology.srlgs.srlg.mpls_te_topo_srlg_link")
	proto.RegisterType((*MplsTeTopoSrlgArea)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.topology.srlgs.srlg.mpls_te_topo_srlg_area")
}

func init() {
	proto.RegisterFile("mpls_te_topo_srlg.proto", fileDescriptor_mpls_te_topo_srlg_8f6a700e427582ab)
}

var fileDescriptor_mpls_te_topo_srlg_8f6a700e427582ab = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0xd0, 0x26, 0xe3, 0x86, 0x8a, 0x15, 0x10, 0x23, 0x0e, 0x04, 0x1f, 0x20, 0x27,
	0x1f, 0xd2, 0x5e, 0x38, 0x70, 0xa8, 0x10, 0xa0, 0x08, 0xc4, 0xc1, 0x50, 0x10, 0x12, 0xd2, 0xc8,
	0xb1, 0xa7, 0x66, 0x21, 0xfe, 0xd0, 0x7a, 0x03, 0x8d, 0xf8, 0x0f, 0xf0, 0xbb, 0xb8, 0xf0, 0x03,
	0xf8, 0x35, 0x68, 0xc6, 0x1f, 0x8d, 0x94, 0x72, 0x0a, 0x5c, 0x92, 0xdd, 0x97, 0x99, 0xf7, 0x66,
	0x9e, 0xde, 0x06, 0xc6, 0x59, 0xb9, 0xac, 0xd0, 0x12, 0xda, 0xa2, 0x2c, 0xb0, 0x32, 0xcb, 0x34,
	0x28, 0x4d, 0x61, 0x0b, 0xf5, 0x38, 0xd6, 0x55, 0x5c, 0xa0, 0x2e, 0x2a, 0xbc, 0x30, 0xd8, 0x56,
	0x15, 0x25, 0x99, 0xa0, 0xbd, 0x54, 0x36, 0xca, 0x93, 0xc5, 0x3a, 0xe0, 0xd6, 0x65, 0x91, 0xae,
	0x03, 0x6e, 0xaf, 0xe4, 0xd3, 0x7f, 0x04, 0xb7, 0xb7, 0x98, 0xf1, 0xc5, 0xd3, 0xf7, 0xaf, 0xd5,
	0x3d, 0x70, 0xe5, 0x92, 0xaf, 0xb2, 0x05, 0x19, 0xcf, 0x99, 0x38, 0xd3, 0x51, 0x08, 0x0c, 0xbd,
	0x12, 0xc4, 0xff, 0xe9, 0xc0, 0x8d, 0xad, 0x5e, 0x75, 0x17, 0x86, 0x75, 0x5b, 0x94, 0x91, 0x37,
	0x9b, 0x38, 0xd3, 0x61, 0x38, 0x90, 0xa6, 0x28, 0x23, 0x75, 0x1f, 0x0e, 0xa3, 0x24, 0xd3, 0x39,
	0x7e, 0x25, 0x9d, 0x7e, 0xb4, 0xde, 0xb1, 0x90, 0xba, 0x82, 0xbd, 0x13, 0x48, 0x59, 0x10, 0x0d,
	0x8c, 0x0c, 0x45, 0x95, 0x77, 0x32, 0xe9, 0x4d, 0xdd, 0xd9, 0x59, 0xb0, 0xd3, 0x92, 0xc1, 0xf6,
	0x86, 0xcc, 0x1e, 0xca, 0xa0, 0xa7, 0xac, 0xe3, 0x23, 0x8c, 0x2d, 0x61, 0x94, 0x24, 0x06, 0xed,
	0xba, 0x24, 0xd4, 0xe5, 0x97, 0x13, 0x5c, 0xe5, 0xf9, 0x2a, 0xe3, 0x85, 0x4c, 0xb1, 0xb2, 0x64,
	0x50, 0x27, 0xe2, 0xc2, 0x30, 0x1c, 0xd4, 0xc0, 0x3c, 0x51, 0x0f, 0xe1, 0x48, 0xe7, 0x96, 0xcc,
	0x79, 0x14, 0x13, 0xea, 0x3c, 0xa1, 0x0b, 0x6f, 0x4f, 0x76, 0xba, 0xde, 0xc1, 0x73, 0x46, 0xfd,
	0x5f, 0x0e, 0xf4, 0x99, 0x5e, 0x29, 0xe8, 0xb3, 0x42, 0xc3, 0x24, 0x67, 0xb6, 0x45, 0x04, 0xb9,
	0x80, 0xaa, 0x4a, 0x28, 0x86, 0xa1, 0xcb, 0xd8, 0x69, 0x0d, 0xa9, 0xef, 0x0e, 0x8c, 0x2f, 0x87,
	0x5a, 0x90, 0xa1, 0xa4, 0x2b, 0xef, 0x4d, 0x9c, 0xa9, 0x3b, 0x7b, 0xbb, 0xa3, 0x49, 0x7f, 0xd9,
	0x3f, 0xbc, 0xc5, 0xe7, 0xb3, 0x4e, 0xb5, 0x19, 0xc8, 0xff, 0x04, 0xd0, 0x75, 0xa0, 0xfa, 0x00,
	0x07, 0xcd, 0x4d, 0x16, 0x73, 0x67, 0x4f, 0x76, 0x9c, 0x86, 0xa9, 0xc2, 0x7d, 0x4b, 0x2c, 0xe7,
	0xff, 0x76, 0xae, 0x4a, 0xe9, 0x52, 0xe7, 0x9f, 0xd9, 0x3a, 0xfe, 0xee, 0xbc, 0xa8, 0x6d, 0x75,
	0x19, 0x6b, 0xad, 0xfb, 0x06, 0x37, 0x37, 0x4b, 0x30, 0xa5, 0x9c, 0x8c, 0x8e, 0xc5, 0x65, 0x77,
	0x36, 0xff, 0x57, 0xb6, 0x61, 0xa8, 0x36, 0x54, 0x9f, 0xd7, 0x22, 0x6a, 0x02, 0x87, 0x96, 0xf0,
	0x32, 0x40, 0x3d, 0x99, 0x0f, 0x2c, 0x85, 0x4d, 0x84, 0xfc, 0x1f, 0x7b, 0x57, 0x2d, 0xc7, 0x01,
	0x55, 0x77, 0x60, 0xa0, 0xd3, 0x12, 0x37, 0xf2, 0x72, 0xa0, 0xd3, 0xf2, 0x4d, 0x1b, 0x99, 0xb4,
	0x44, 0x9d, 0xf3, 0x6c, 0x31, 0x75, 0x91, 0x49, 0xcb, 0x79, 0x03, 0xb5, 0xdd, 0xcc, 0x24, 0xb2,
	0x23, 0xe9, 0xe6, 0xbc, 0xab, 0x07, 0x70, 0xd4, 0xfe, 0x84, 0xe7, 0x85, 0xc9, 0x22, 0xeb, 0xf5,
	0x85, 0x60, 0xd4, 0x54, 0x3c, 0x13, 0xb0, 0x7b, 0x8c, 0xbc, 0x58, 0xe5, 0x5d, 0xfb, 0x4f, 0x8f,
	0x91, 0xd9, 0xeb, 0xc7, 0xf8, 0x92, 0x75, 0x16, 0xfb, 0xf2, 0xcf, 0x76, 0xfc, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x0c, 0x9f, 0x98, 0x8c, 0xf4, 0x04, 0x00, 0x00,
}
