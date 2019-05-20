// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pce_stats_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_pce_stdby_statistics

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

// PCE Stats Information
type PceStatsBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceStatsBag_KEYS) Reset()         { *m = PceStatsBag_KEYS{} }
func (m *PceStatsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PceStatsBag_KEYS) ProtoMessage()    {}
func (*PceStatsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_pce_stats_bag_cdfbce16c8d23f0c, []int{0}
}
func (m *PceStatsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceStatsBag_KEYS.Unmarshal(m, b)
}
func (m *PceStatsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceStatsBag_KEYS.Marshal(b, m, deterministic)
}
func (dst *PceStatsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceStatsBag_KEYS.Merge(dst, src)
}
func (m *PceStatsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PceStatsBag_KEYS.Size(m)
}
func (m *PceStatsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PceStatsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PceStatsBag_KEYS proto.InternalMessageInfo

type PceStatsBag struct {
	// Neighbor Count
	Neighbors uint32 `protobuf:"varint,50,opt,name=neighbors" json:"neighbors,omitempty"`
	// Tunnel Count
	Tunnels uint32 `protobuf:"varint,51,opt,name=tunnels" json:"tunnels,omitempty"`
	// Total number of PCEReq
	PceReqTotal uint32 `protobuf:"varint,52,opt,name=pce_req_total,json=pceReqTotal" json:"pce_req_total,omitempty"`
	// Number of pending req
	PceReqPending uint32 `protobuf:"varint,53,opt,name=pce_req_pending,json=pceReqPending" json:"pce_req_pending,omitempty"`
	// Number of timedout req
	PceReqTimedOut uint32 `protobuf:"varint,54,opt,name=pce_req_timed_out,json=pceReqTimedOut" json:"pce_req_timed_out,omitempty"`
	// Maximum number of TCP file descriptors used
	MaxFileDescriptors uint32 `protobuf:"varint,55,opt,name=max_file_descriptors,json=maxFileDescriptors" json:"max_file_descriptors,omitempty"`
	// Number of TCP file descriptors currently in use
	CurrentFileDescriptors uint32 `protobuf:"varint,56,opt,name=current_file_descriptors,json=currentFileDescriptors" json:"current_file_descriptors,omitempty"`
	// Max Input Queue Depth
	MaximumInQueueDepth []uint32 `protobuf:"varint,57,rep,packed,name=maximum_in_queue_depth,json=maximumInQueueDepth" json:"maximum_in_queue_depth,omitempty"`
	// Average Input Queue Depth
	AverageInQueueDepth []uint32 `protobuf:"varint,58,rep,packed,name=average_in_queue_depth,json=averageInQueueDepth" json:"average_in_queue_depth,omitempty"`
	// Current PCE queue length
	CurrentPceqLength uint32 `protobuf:"varint,59,opt,name=current_pceq_length,json=currentPceqLength" json:"current_pceq_length,omitempty"`
	// Current TE queue length
	CurrentTeqLength uint32 `protobuf:"varint,60,opt,name=current_teq_length,json=currentTeqLength" json:"current_teq_length,omitempty"`
	// Max PCE queue length
	MaxPceqLength uint32 `protobuf:"varint,61,opt,name=max_pceq_length,json=maxPceqLength" json:"max_pceq_length,omitempty"`
	// Max TE queue length
	MaxTeqLength uint32 `protobuf:"varint,62,opt,name=max_teq_length,json=maxTeqLength" json:"max_teq_length,omitempty"`
	// IGP Statistics
	IgpStatistics        *PceIgpStatsBag `protobuf:"bytes,63,opt,name=igp_statistics,json=igpStatistics" json:"igp_statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PceStatsBag) Reset()         { *m = PceStatsBag{} }
func (m *PceStatsBag) String() string { return proto.CompactTextString(m) }
func (*PceStatsBag) ProtoMessage()    {}
func (*PceStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_pce_stats_bag_cdfbce16c8d23f0c, []int{1}
}
func (m *PceStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceStatsBag.Unmarshal(m, b)
}
func (m *PceStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceStatsBag.Marshal(b, m, deterministic)
}
func (dst *PceStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceStatsBag.Merge(dst, src)
}
func (m *PceStatsBag) XXX_Size() int {
	return xxx_messageInfo_PceStatsBag.Size(m)
}
func (m *PceStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceStatsBag proto.InternalMessageInfo

func (m *PceStatsBag) GetNeighbors() uint32 {
	if m != nil {
		return m.Neighbors
	}
	return 0
}

func (m *PceStatsBag) GetTunnels() uint32 {
	if m != nil {
		return m.Tunnels
	}
	return 0
}

func (m *PceStatsBag) GetPceReqTotal() uint32 {
	if m != nil {
		return m.PceReqTotal
	}
	return 0
}

func (m *PceStatsBag) GetPceReqPending() uint32 {
	if m != nil {
		return m.PceReqPending
	}
	return 0
}

func (m *PceStatsBag) GetPceReqTimedOut() uint32 {
	if m != nil {
		return m.PceReqTimedOut
	}
	return 0
}

func (m *PceStatsBag) GetMaxFileDescriptors() uint32 {
	if m != nil {
		return m.MaxFileDescriptors
	}
	return 0
}

func (m *PceStatsBag) GetCurrentFileDescriptors() uint32 {
	if m != nil {
		return m.CurrentFileDescriptors
	}
	return 0
}

func (m *PceStatsBag) GetMaximumInQueueDepth() []uint32 {
	if m != nil {
		return m.MaximumInQueueDepth
	}
	return nil
}

func (m *PceStatsBag) GetAverageInQueueDepth() []uint32 {
	if m != nil {
		return m.AverageInQueueDepth
	}
	return nil
}

func (m *PceStatsBag) GetCurrentPceqLength() uint32 {
	if m != nil {
		return m.CurrentPceqLength
	}
	return 0
}

func (m *PceStatsBag) GetCurrentTeqLength() uint32 {
	if m != nil {
		return m.CurrentTeqLength
	}
	return 0
}

func (m *PceStatsBag) GetMaxPceqLength() uint32 {
	if m != nil {
		return m.MaxPceqLength
	}
	return 0
}

func (m *PceStatsBag) GetMaxTeqLength() uint32 {
	if m != nil {
		return m.MaxTeqLength
	}
	return 0
}

func (m *PceStatsBag) GetIgpStatistics() *PceIgpStatsBag {
	if m != nil {
		return m.IgpStatistics
	}
	return nil
}

// PCE IGP Stats
type PceIgpStatsBag struct {
	// ABR Lookup Min
	AbrLookupMin uint64 `protobuf:"varint,1,opt,name=abr_lookup_min,json=abrLookupMin" json:"abr_lookup_min,omitempty"`
	// ABR Lookup Max
	AbrLookupMax uint64 `protobuf:"varint,2,opt,name=abr_lookup_max,json=abrLookupMax" json:"abr_lookup_max,omitempty"`
	// ABR Lookup Avg
	AbrLookupAvg uint64 `protobuf:"varint,3,opt,name=abr_lookup_avg,json=abrLookupAvg" json:"abr_lookup_avg,omitempty"`
	// ABR Lookup Timeout
	AbrLookupTimeout uint64 `protobuf:"varint,4,opt,name=abr_lookup_timeout,json=abrLookupTimeout" json:"abr_lookup_timeout,omitempty"`
	// ABR Lookup Complete
	AbrLookupComplete    uint64   `protobuf:"varint,5,opt,name=abr_lookup_complete,json=abrLookupComplete" json:"abr_lookup_complete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceIgpStatsBag) Reset()         { *m = PceIgpStatsBag{} }
func (m *PceIgpStatsBag) String() string { return proto.CompactTextString(m) }
func (*PceIgpStatsBag) ProtoMessage()    {}
func (*PceIgpStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_pce_stats_bag_cdfbce16c8d23f0c, []int{2}
}
func (m *PceIgpStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceIgpStatsBag.Unmarshal(m, b)
}
func (m *PceIgpStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceIgpStatsBag.Marshal(b, m, deterministic)
}
func (dst *PceIgpStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceIgpStatsBag.Merge(dst, src)
}
func (m *PceIgpStatsBag) XXX_Size() int {
	return xxx_messageInfo_PceIgpStatsBag.Size(m)
}
func (m *PceIgpStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceIgpStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceIgpStatsBag proto.InternalMessageInfo

func (m *PceIgpStatsBag) GetAbrLookupMin() uint64 {
	if m != nil {
		return m.AbrLookupMin
	}
	return 0
}

func (m *PceIgpStatsBag) GetAbrLookupMax() uint64 {
	if m != nil {
		return m.AbrLookupMax
	}
	return 0
}

func (m *PceIgpStatsBag) GetAbrLookupAvg() uint64 {
	if m != nil {
		return m.AbrLookupAvg
	}
	return 0
}

func (m *PceIgpStatsBag) GetAbrLookupTimeout() uint64 {
	if m != nil {
		return m.AbrLookupTimeout
	}
	return 0
}

func (m *PceIgpStatsBag) GetAbrLookupComplete() uint64 {
	if m != nil {
		return m.AbrLookupComplete
	}
	return 0
}

func init() {
	proto.RegisterType((*PceStatsBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce_stdby.statistics.pce_stats_bag_KEYS")
	proto.RegisterType((*PceStatsBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce_stdby.statistics.pce_stats_bag")
	proto.RegisterType((*PceIgpStatsBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce_stdby.statistics.pce_igp_stats_bag")
}

func init() { proto.RegisterFile("pce_stats_bag.proto", fileDescriptor_pce_stats_bag_cdfbce16c8d23f0c) }

var fileDescriptor_pce_stats_bag_cdfbce16c8d23f0c = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdd, 0x6e, 0xd3, 0x30,
	0x1c, 0xc5, 0x55, 0xb6, 0x81, 0xf0, 0xd6, 0x6e, 0x75, 0xa7, 0xc9, 0x17, 0x5c, 0x54, 0x15, 0x9a,
	0x8a, 0x84, 0x22, 0xb4, 0xf2, 0x31, 0xbe, 0x85, 0x18, 0x93, 0x10, 0x43, 0x8c, 0xae, 0x37, 0x5c,
	0x59, 0x4e, 0xfa, 0x27, 0xb5, 0x88, 0x3f, 0x6a, 0x3b, 0x55, 0x78, 0x2a, 0x1e, 0x8a, 0x17, 0x41,
	0x76, 0xd2, 0x26, 0x0d, 0x77, 0xbb, 0xcc, 0x39, 0xbf, 0x73, 0xf4, 0x57, 0x72, 0x82, 0x06, 0x3a,
	0x01, 0x6a, 0x1d, 0x73, 0x96, 0xc6, 0x2c, 0x8d, 0xb4, 0x51, 0x4e, 0xe1, 0x49, 0xc2, 0x6d, 0xa2,
	0x28, 0x57, 0x96, 0x16, 0x86, 0x0a, 0x9d, 0x59, 0xea, 0x80, 0x2a, 0x0d, 0x26, 0x0a, 0x0f, 0x65,
	0x66, 0x1e, 0xff, 0x8e, 0x7c, 0x92, 0x5b, 0xc7, 0x13, 0x3b, 0x3a, 0x46, 0x78, 0xab, 0x8b, 0x7e,
	0xf9, 0xf4, 0xe3, 0x66, 0xf4, 0x67, 0x0f, 0x75, 0xb7, 0x64, 0xfc, 0x00, 0xdd, 0x97, 0xc0, 0xd3,
	0x45, 0xac, 0x8c, 0x25, 0x67, 0xc3, 0xce, 0xb8, 0x3b, 0xad, 0x05, 0x4c, 0xd0, 0x3d, 0x97, 0x4b,
	0x09, 0x99, 0x25, 0x93, 0xe0, 0xad, 0x1f, 0xf1, 0xa8, 0x2c, 0x32, 0xb0, 0xa4, 0x4e, 0x39, 0x96,
	0x91, 0xa7, 0xc1, 0xdf, 0xd7, 0x09, 0x4c, 0x61, 0x39, 0xf3, 0x12, 0x3e, 0x45, 0x87, 0x6b, 0x46,
	0x83, 0x9c, 0x73, 0x99, 0x92, 0x67, 0x81, 0xea, 0x96, 0xd4, 0x75, 0x29, 0xe2, 0x47, 0xa8, 0xbf,
	0xe9, 0xe2, 0x02, 0xe6, 0x54, 0xe5, 0x8e, 0x3c, 0x0f, 0x64, 0xaf, 0xea, 0xf3, 0xf2, 0xb7, 0xdc,
	0xe1, 0x27, 0xe8, 0x58, 0xb0, 0x82, 0xfe, 0xe4, 0x19, 0xd0, 0x39, 0xd8, 0xc4, 0x70, 0xed, 0xfc,
	0xe5, 0x2f, 0x02, 0x8d, 0x05, 0x2b, 0x2e, 0x79, 0x06, 0x17, 0xb5, 0x83, 0xcf, 0x11, 0x49, 0x72,
	0x63, 0x40, 0xba, 0xff, 0x53, 0xe7, 0x21, 0x75, 0x52, 0xf9, 0xed, 0xe4, 0x04, 0x9d, 0x08, 0x56,
	0x70, 0x91, 0x0b, 0xca, 0x25, 0x5d, 0xe6, 0x90, 0xfb, 0xb4, 0x76, 0x0b, 0xf2, 0x72, 0xb8, 0x33,
	0xee, 0x4e, 0x07, 0x95, 0xfb, 0x59, 0x7e, 0xf7, 0xde, 0x85, 0xb7, 0x7c, 0x88, 0xad, 0xc0, 0xb0,
	0x14, 0xda, 0xa1, 0x57, 0x65, 0xa8, 0x72, 0xb7, 0x42, 0x11, 0x1a, 0xac, 0x6f, 0xd4, 0x09, 0x2c,
	0x69, 0x06, 0x32, 0x75, 0x0b, 0xf2, 0x3a, 0x9c, 0xd7, 0xaf, 0xac, 0xeb, 0x04, 0x96, 0x57, 0xc1,
	0xc0, 0x8f, 0x11, 0x5e, 0xf3, 0xae, 0xc6, 0xdf, 0x04, 0xfc, 0xa8, 0x72, 0x66, 0x1b, 0xfa, 0x14,
	0x1d, 0xfa, 0x77, 0xd6, 0x6c, 0x7e, 0x5b, 0x7e, 0x06, 0xc1, 0x8a, 0x46, 0xeb, 0x43, 0xd4, 0xf3,
	0x5c, 0xa3, 0xf1, 0x5d, 0xc0, 0x0e, 0x04, 0x2b, 0xea, 0x36, 0x81, 0x7a, 0x3c, 0xd5, 0xb4, 0x9e,
	0x1a, 0x79, 0x3f, 0xec, 0x8c, 0xf7, 0xcf, 0x2e, 0xa3, 0x5b, 0xcc, 0x34, 0xf2, 0xe2, 0xba, 0x2e,
	0x0c, 0x72, 0xda, 0xe5, 0xa9, 0xbe, 0xa9, 0x77, 0xfc, 0xb7, 0x53, 0x8e, 0x63, 0x0b, 0xf2, 0xa7,
	0xb2, 0xd8, 0xd0, 0x4c, 0xa9, 0x5f, 0xb9, 0xa6, 0x82, 0x4b, 0xd2, 0x19, 0x76, 0xc6, 0xbb, 0xd3,
	0x03, 0x16, 0x9b, 0xab, 0x20, 0x7e, 0xe5, 0xb2, 0x4d, 0xb1, 0x82, 0xdc, 0x69, 0x53, 0xac, 0x68,
	0x51, 0x6c, 0x95, 0x92, 0x9d, 0x16, 0xf5, 0x61, 0x95, 0xfa, 0x57, 0xde, 0xa0, 0xfc, 0x4c, 0xfd,
	0x48, 0x77, 0x03, 0x79, 0xb4, 0x21, 0x67, 0xa5, 0xee, 0x3f, 0x68, 0x83, 0x4e, 0x94, 0xd0, 0x19,
	0x38, 0x20, 0x7b, 0x01, 0xef, 0x6f, 0xf0, 0x8f, 0x95, 0x11, 0xdf, 0x0d, 0x7f, 0xfa, 0xe4, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0xa3, 0xd8, 0x7d, 0x00, 0x04, 0x00, 0x00,
}