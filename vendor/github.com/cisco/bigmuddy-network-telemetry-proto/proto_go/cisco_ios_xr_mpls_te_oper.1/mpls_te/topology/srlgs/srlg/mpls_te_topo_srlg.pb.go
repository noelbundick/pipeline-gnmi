// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_topo_srlg.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_topology_srlgs_srlg

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
	return fileDescriptor_mpls_te_topo_srlg_000388a1b6826c4f, []int{0}
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
	return fileDescriptor_mpls_te_topo_srlg_000388a1b6826c4f, []int{1}
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
	return fileDescriptor_mpls_te_topo_srlg_000388a1b6826c4f, []int{2}
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
	return fileDescriptor_mpls_te_topo_srlg_000388a1b6826c4f, []int{3}
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
	return fileDescriptor_mpls_te_topo_srlg_000388a1b6826c4f, []int{4}
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
	return fileDescriptor_mpls_te_topo_srlg_000388a1b6826c4f, []int{5}
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
	return fileDescriptor_mpls_te_topo_srlg_000388a1b6826c4f, []int{6}
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
	proto.RegisterType((*MplsTeTopoSrlg_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.topology.srlgs.srlg.mpls_te_topo_srlg_KEYS")
	proto.RegisterType((*MplsTeTopoSrlg)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.topology.srlgs.srlg.mpls_te_topo_srlg")
	proto.RegisterType((*TeAddrTypeIpv4Unnum)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.topology.srlgs.srlg.te_addr_type_ipv4_unnum")
	proto.RegisterType((*Addr)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.topology.srlgs.srlg.addr")
	proto.RegisterType((*TeAddrT_)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.topology.srlgs.srlg.te_addr_t_")
	proto.RegisterType((*MplsTeTopoSrlgLink)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.topology.srlgs.srlg.mpls_te_topo_srlg_link")
	proto.RegisterType((*MplsTeTopoSrlgArea)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.topology.srlgs.srlg.mpls_te_topo_srlg_area")
}

func init() {
	proto.RegisterFile("mpls_te_topo_srlg.proto", fileDescriptor_mpls_te_topo_srlg_000388a1b6826c4f)
}

var fileDescriptor_mpls_te_topo_srlg_000388a1b6826c4f = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x3d, 0x6f, 0x13, 0x41,
	0x10, 0xd5, 0x25, 0x26, 0xb1, 0xe7, 0x62, 0x22, 0x56, 0x80, 0x0f, 0x51, 0x60, 0xae, 0x00, 0x57,
	0x57, 0x38, 0xa1, 0x40, 0x54, 0x2e, 0x00, 0x59, 0x40, 0x8a, 0x23, 0x08, 0x51, 0x8d, 0xce, 0xbe,
	0xc9, 0xb1, 0xc2, 0xf7, 0xa1, 0xdd, 0x35, 0x24, 0x7f, 0x00, 0x7e, 0x17, 0x0d, 0x3f, 0x80, 0x5f,
	0x84, 0x66, 0xee, 0x23, 0x96, 0x1c, 0x1a, 0x43, 0x63, 0xef, 0x3e, 0xcf, 0x7b, 0x33, 0x6f, 0xf4,
	0xd6, 0x30, 0xca, 0xab, 0x95, 0x45, 0x47, 0xe8, 0xca, 0xaa, 0x44, 0x6b, 0x56, 0x59, 0x54, 0x99,
	0xd2, 0x95, 0xea, 0xd9, 0x52, 0xdb, 0x65, 0x89, 0xba, 0xb4, 0x78, 0x69, 0xb0, 0xad, 0x2a, 0x2b,
	0x32, 0x51, 0x73, 0x89, 0x98, 0xb2, 0x2a, 0xb3, 0xab, 0x88, 0x69, 0x56, 0x3e, 0xc3, 0xe7, 0x70,
	0x7f, 0x4b, 0x11, 0xdf, 0xbc, 0xfc, 0xf4, 0x5e, 0x3d, 0x02, 0x5f, 0x2e, 0xc5, 0x3a, 0x5f, 0x90,
	0x09, 0xbc, 0xb1, 0x37, 0x19, 0xc6, 0xc0, 0xd0, 0x99, 0x20, 0xe1, 0x4f, 0x0f, 0xee, 0x6c, 0x71,
	0xd5, 0x43, 0x18, 0xd4, 0xb4, 0x24, 0xa7, 0x60, 0x3a, 0xf6, 0x26, 0x83, 0xb8, 0x2f, 0xa4, 0x24,
	0x27, 0xf5, 0x18, 0x8e, 0x92, 0x34, 0xd7, 0x05, 0x7e, 0x23, 0x9d, 0x7d, 0x76, 0xc1, 0x89, 0x88,
	0xfa, 0x82, 0x7d, 0x14, 0x48, 0xad, 0x40, 0x7a, 0x60, 0x62, 0x28, 0xb1, 0xc1, 0xe9, 0x78, 0x7f,
	0xe2, 0x4f, 0xdf, 0x45, 0x3b, 0x99, 0x8b, 0xb6, 0x9d, 0xb1, 0x6a, 0x2c, 0x03, 0xce, 0x58, 0x3f,
	0x44, 0x18, 0x39, 0xc2, 0x24, 0x4d, 0x0d, 0xba, 0xab, 0x8a, 0x50, 0x57, 0x5f, 0x4f, 0x71, 0x5d,
	0x14, 0xeb, 0x9c, 0x8d, 0x98, 0x72, 0xed, 0xc8, 0xa0, 0x4e, 0xc5, 0xfd, 0x20, 0xee, 0xd7, 0xc0,
	0x3c, 0x55, 0x4f, 0xe1, 0x58, 0x17, 0x8e, 0xcc, 0x45, 0xb2, 0x24, 0xd4, 0x45, 0x4a, 0x97, 0xc1,
	0x9e, 0x78, 0xb9, 0xdd, 0xc1, 0x73, 0x46, 0xc3, 0x5f, 0x1e, 0xf4, 0x58, 0x5e, 0x29, 0xe8, 0x71,
	0x87, 0x46, 0x49, 0xce, 0xbc, 0x0e, 0x69, 0xc8, 0x05, 0x64, 0xad, 0x48, 0x0c, 0x62, 0x9f, 0xb1,
	0x59, 0x0d, 0xa9, 0xef, 0x1e, 0x8c, 0xae, 0x87, 0x5a, 0x90, 0xa1, 0xb4, 0x2b, 0xdf, 0x1f, 0x7b,
	0x13, 0x7f, 0x7a, 0xb6, 0xe3, 0x72, 0xfe, 0xe2, 0x3b, 0xbe, 0xc7, 0xe7, 0x0f, 0x5d, 0xb7, 0x66,
	0x90, 0x70, 0x01, 0xd0, 0x31, 0x50, 0x9d, 0xc3, 0x61, 0x73, 0x13, 0x43, 0xfe, 0xf4, 0xc5, 0x8e,
	0x53, 0xb0, 0x44, 0x7c, 0xe0, 0x88, 0xdb, 0x84, 0xbf, 0xbd, 0x9b, 0xd2, 0xb8, 0xd2, 0xc5, 0x17,
	0x5e, 0x15, 0x7f, 0x77, 0xde, 0xeb, 0x35, 0xfa, 0x8c, 0xb5, 0xab, 0xb2, 0x70, 0x77, 0xb3, 0x04,
	0x33, 0x2a, 0xc8, 0xe8, 0xa5, 0x6c, 0xd5, 0x9f, 0xce, 0xfe, 0x75, 0x4d, 0x18, 0xab, 0x8d, 0x6e,
	0xaf, 0x6b, 0x71, 0x35, 0x86, 0x23, 0x47, 0x78, 0x1d, 0x94, 0x7d, 0x99, 0x0b, 0x1c, 0xc5, 0x4d,
	0x54, 0xc2, 0x1f, 0x7b, 0x37, 0x99, 0xe2, 0x20, 0xaa, 0x07, 0xd0, 0xd7, 0x59, 0x85, 0x1b, 0xb9,
	0x38, 0xd4, 0x59, 0x75, 0xde, 0x46, 0x23, 0xab, 0x50, 0x17, 0xd6, 0x25, 0xc5, 0x92, 0xba, 0x68,
	0x64, 0xd5, 0xbc, 0x81, 0x5a, 0x36, 0x2b, 0x49, 0xdb, 0xa1, 0xb0, 0x39, 0xd7, 0xea, 0x09, 0x1c,
	0xb7, 0x3f, 0xe1, 0x45, 0x69, 0xf2, 0xc4, 0x05, 0x3d, 0x11, 0x18, 0x36, 0x15, 0xaf, 0x04, 0xec,
	0x1e, 0x1b, 0x1b, 0xb3, 0xc1, 0xad, 0xff, 0xfc, 0xd8, 0x58, 0xb5, 0x7e, 0x6c, 0x6f, 0x59, 0x7f,
	0x71, 0x20, 0xff, 0x54, 0x27, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x47, 0x7f, 0xf6, 0x9c, 0xc4,
	0x04, 0x00, 0x00,
}
