// Code generated by protoc-gen-go. DO NOT EDIT.
// source: te_stats_bag_lsp_t.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_signalling_counters_tails_counters_tails_counter

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

// MPLS TE LSP Statistics
type TeStatsBagLspT_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,2,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	TunnelId             uint32   `protobuf:"varint,3,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,4,opt,name=extended_tunnel_id,json=extendedTunnelId" json:"extended_tunnel_id,omitempty"`
	LspId                uint32   `protobuf:"varint,5,opt,name=lsp_id,json=lspId" json:"lsp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsBagLspT_KEYS) Reset()         { *m = TeStatsBagLspT_KEYS{} }
func (m *TeStatsBagLspT_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagLspT_KEYS) ProtoMessage()    {}
func (*TeStatsBagLspT_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_bag_lsp_t_a15aa9202c376e6d, []int{0}
}
func (m *TeStatsBagLspT_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsBagLspT_KEYS.Unmarshal(m, b)
}
func (m *TeStatsBagLspT_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsBagLspT_KEYS.Marshal(b, m, deterministic)
}
func (dst *TeStatsBagLspT_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsBagLspT_KEYS.Merge(dst, src)
}
func (m *TeStatsBagLspT_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeStatsBagLspT_KEYS.Size(m)
}
func (m *TeStatsBagLspT_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsBagLspT_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsBagLspT_KEYS proto.InternalMessageInfo

func (m *TeStatsBagLspT_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *TeStatsBagLspT_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *TeStatsBagLspT_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *TeStatsBagLspT_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *TeStatsBagLspT_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

type TeStatsBagLspT struct {
	// Tunnel Name
	TunnelName string `protobuf:"bytes,50,opt,name=tunnel_name,json=tunnelName" json:"tunnel_name,omitempty"`
	// LSP statistics
	Statistics *TeStatsSigT `protobuf:"bytes,51,opt,name=statistics" json:"statistics,omitempty"`
	// List of S2L Statistics
	S2LStatistics        []*TeStatsBagS2LT `protobuf:"bytes,52,rep,name=s2_l_statistics,json=s2LStatistics" json:"s2_l_statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TeStatsBagLspT) Reset()         { *m = TeStatsBagLspT{} }
func (m *TeStatsBagLspT) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagLspT) ProtoMessage()    {}
func (*TeStatsBagLspT) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_bag_lsp_t_a15aa9202c376e6d, []int{1}
}
func (m *TeStatsBagLspT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsBagLspT.Unmarshal(m, b)
}
func (m *TeStatsBagLspT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsBagLspT.Marshal(b, m, deterministic)
}
func (dst *TeStatsBagLspT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsBagLspT.Merge(dst, src)
}
func (m *TeStatsBagLspT) XXX_Size() int {
	return xxx_messageInfo_TeStatsBagLspT.Size(m)
}
func (m *TeStatsBagLspT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsBagLspT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsBagLspT proto.InternalMessageInfo

func (m *TeStatsBagLspT) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *TeStatsBagLspT) GetStatistics() *TeStatsSigT {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *TeStatsBagLspT) GetS2LStatistics() []*TeStatsBagS2LT {
	if m != nil {
		return m.S2LStatistics
	}
	return nil
}

// Send-Recv count for TE Signaling
type TeStatsSigT struct {
	// Unknown TX events
	TxEventUnknown uint32 `protobuf:"varint,1,opt,name=tx_event_unknown,json=txEventUnknown" json:"tx_event_unknown,omitempty"`
	// TX Path Create event
	TxPathCreateEvent uint32 `protobuf:"varint,2,opt,name=tx_path_create_event,json=txPathCreateEvent" json:"tx_path_create_event,omitempty"`
	// TX Path Change event
	TxPathChangeEvent uint32 `protobuf:"varint,3,opt,name=tx_path_change_event,json=txPathChangeEvent" json:"tx_path_change_event,omitempty"`
	// TX Path Delete event
	TxPathDeleteEvent uint32 `protobuf:"varint,4,opt,name=tx_path_delete_event,json=txPathDeleteEvent" json:"tx_path_delete_event,omitempty"`
	// TX Path Error event
	TxPathErrorEvent uint32 `protobuf:"varint,5,opt,name=tx_path_error_event,json=txPathErrorEvent" json:"tx_path_error_event,omitempty"`
	// TX Resv Create event
	TxResvCreateEvent uint32 `protobuf:"varint,6,opt,name=tx_resv_create_event,json=txResvCreateEvent" json:"tx_resv_create_event,omitempty"`
	// TX Resv Change event
	TxResvChangeEvent uint32 `protobuf:"varint,7,opt,name=tx_resv_change_event,json=txResvChangeEvent" json:"tx_resv_change_event,omitempty"`
	// TX Resv Delete event
	TxResvDeleteEvent uint32 `protobuf:"varint,8,opt,name=tx_resv_delete_event,json=txResvDeleteEvent" json:"tx_resv_delete_event,omitempty"`
	// TX Resv Error event
	TxResvErrorEvent uint32 `protobuf:"varint,9,opt,name=tx_resv_error_event,json=txResvErrorEvent" json:"tx_resv_error_event,omitempty"`
	// TX Path Reeval Query event
	TxPathReevalQueryEvent uint32 `protobuf:"varint,10,opt,name=tx_path_reeval_query_event,json=txPathReevalQueryEvent" json:"tx_path_reeval_query_event,omitempty"`
	// RX Unknown events
	RxEventUnknown uint32 `protobuf:"varint,11,opt,name=rx_event_unknown,json=rxEventUnknown" json:"rx_event_unknown,omitempty"`
	// RX Path Create event
	RxPathCreateEvent uint32 `protobuf:"varint,12,opt,name=rx_path_create_event,json=rxPathCreateEvent" json:"rx_path_create_event,omitempty"`
	// RX Path Change event
	RxPathChangeEvent uint32 `protobuf:"varint,13,opt,name=rx_path_change_event,json=rxPathChangeEvent" json:"rx_path_change_event,omitempty"`
	// RX Path Delete event
	RxPathDeleteEvent uint32 `protobuf:"varint,14,opt,name=rx_path_delete_event,json=rxPathDeleteEvent" json:"rx_path_delete_event,omitempty"`
	// RX Path Error event
	RxPathErrorEvent uint32 `protobuf:"varint,15,opt,name=rx_path_error_event,json=rxPathErrorEvent" json:"rx_path_error_event,omitempty"`
	// RX Resv Create event
	RxResvCreateEvent uint32 `protobuf:"varint,16,opt,name=rx_resv_create_event,json=rxResvCreateEvent" json:"rx_resv_create_event,omitempty"`
	// RX Resv Change event
	RxResvChangeEvent uint32 `protobuf:"varint,17,opt,name=rx_resv_change_event,json=rxResvChangeEvent" json:"rx_resv_change_event,omitempty"`
	// RX Resv Delete event
	RxResvDeleteEvent uint32 `protobuf:"varint,18,opt,name=rx_resv_delete_event,json=rxResvDeleteEvent" json:"rx_resv_delete_event,omitempty"`
	// RX Resv Error event
	RxResvErrorEvent uint32 `protobuf:"varint,19,opt,name=rx_resv_error_event,json=rxResvErrorEvent" json:"rx_resv_error_event,omitempty"`
	// RX Path Reeval Query event
	RxPathReevalQueryEvent uint32 `protobuf:"varint,20,opt,name=rx_path_reeval_query_event,json=rxPathReevalQueryEvent" json:"rx_path_reeval_query_event,omitempty"`
	// Backup Assign event
	TxBackupAssignEvent uint32 `protobuf:"varint,21,opt,name=tx_backup_assign_event,json=txBackupAssignEvent" json:"tx_backup_assign_event,omitempty"`
	// Error on Backup Assign event
	RxBackupAssignErrEvent uint32 `protobuf:"varint,22,opt,name=rx_backup_assign_err_event,json=rxBackupAssignErrEvent" json:"rx_backup_assign_err_event,omitempty"`
	// Total TE Signalling event count
	EventsTotalCount uint32 `protobuf:"varint,23,opt,name=events_total_count,json=eventsTotalCount" json:"events_total_count,omitempty"`
	// TE Signaling event count
	EventsCount          uint32   `protobuf:"varint,24,opt,name=events_count,json=eventsCount" json:"events_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsSigT) Reset()         { *m = TeStatsSigT{} }
func (m *TeStatsSigT) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigT) ProtoMessage()    {}
func (*TeStatsSigT) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_bag_lsp_t_a15aa9202c376e6d, []int{2}
}
func (m *TeStatsSigT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsSigT.Unmarshal(m, b)
}
func (m *TeStatsSigT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsSigT.Marshal(b, m, deterministic)
}
func (dst *TeStatsSigT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsSigT.Merge(dst, src)
}
func (m *TeStatsSigT) XXX_Size() int {
	return xxx_messageInfo_TeStatsSigT.Size(m)
}
func (m *TeStatsSigT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsSigT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsSigT proto.InternalMessageInfo

func (m *TeStatsSigT) GetTxEventUnknown() uint32 {
	if m != nil {
		return m.TxEventUnknown
	}
	return 0
}

func (m *TeStatsSigT) GetTxPathCreateEvent() uint32 {
	if m != nil {
		return m.TxPathCreateEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxPathChangeEvent() uint32 {
	if m != nil {
		return m.TxPathChangeEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxPathDeleteEvent() uint32 {
	if m != nil {
		return m.TxPathDeleteEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxPathErrorEvent() uint32 {
	if m != nil {
		return m.TxPathErrorEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxResvCreateEvent() uint32 {
	if m != nil {
		return m.TxResvCreateEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxResvChangeEvent() uint32 {
	if m != nil {
		return m.TxResvChangeEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxResvDeleteEvent() uint32 {
	if m != nil {
		return m.TxResvDeleteEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxResvErrorEvent() uint32 {
	if m != nil {
		return m.TxResvErrorEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxPathReevalQueryEvent() uint32 {
	if m != nil {
		return m.TxPathReevalQueryEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxEventUnknown() uint32 {
	if m != nil {
		return m.RxEventUnknown
	}
	return 0
}

func (m *TeStatsSigT) GetRxPathCreateEvent() uint32 {
	if m != nil {
		return m.RxPathCreateEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxPathChangeEvent() uint32 {
	if m != nil {
		return m.RxPathChangeEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxPathDeleteEvent() uint32 {
	if m != nil {
		return m.RxPathDeleteEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxPathErrorEvent() uint32 {
	if m != nil {
		return m.RxPathErrorEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxResvCreateEvent() uint32 {
	if m != nil {
		return m.RxResvCreateEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxResvChangeEvent() uint32 {
	if m != nil {
		return m.RxResvChangeEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxResvDeleteEvent() uint32 {
	if m != nil {
		return m.RxResvDeleteEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxResvErrorEvent() uint32 {
	if m != nil {
		return m.RxResvErrorEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxPathReevalQueryEvent() uint32 {
	if m != nil {
		return m.RxPathReevalQueryEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxBackupAssignEvent() uint32 {
	if m != nil {
		return m.TxBackupAssignEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxBackupAssignErrEvent() uint32 {
	if m != nil {
		return m.RxBackupAssignErrEvent
	}
	return 0
}

func (m *TeStatsSigT) GetEventsTotalCount() uint32 {
	if m != nil {
		return m.EventsTotalCount
	}
	return 0
}

func (m *TeStatsSigT) GetEventsCount() uint32 {
	if m != nil {
		return m.EventsCount
	}
	return 0
}

// MPLS TE S2L Statistics
type TeStatsBagS2LT struct {
	// Subgroup Originator
	SubGroupOriginator string `protobuf:"bytes,1,opt,name=sub_group_originator,json=subGroupOriginator" json:"sub_group_originator,omitempty"`
	// subgroup ID
	SubGroupId uint32 `protobuf:"varint,2,opt,name=sub_group_id,json=subGroupId" json:"sub_group_id,omitempty"`
	// Destination address
	DestinationAddress string `protobuf:"bytes,3,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// S2L stats
	Statistics           *TeStatsSigT `protobuf:"bytes,4,opt,name=statistics" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TeStatsBagS2LT) Reset()         { *m = TeStatsBagS2LT{} }
func (m *TeStatsBagS2LT) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagS2LT) ProtoMessage()    {}
func (*TeStatsBagS2LT) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_bag_lsp_t_a15aa9202c376e6d, []int{3}
}
func (m *TeStatsBagS2LT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsBagS2LT.Unmarshal(m, b)
}
func (m *TeStatsBagS2LT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsBagS2LT.Marshal(b, m, deterministic)
}
func (dst *TeStatsBagS2LT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsBagS2LT.Merge(dst, src)
}
func (m *TeStatsBagS2LT) XXX_Size() int {
	return xxx_messageInfo_TeStatsBagS2LT.Size(m)
}
func (m *TeStatsBagS2LT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsBagS2LT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsBagS2LT proto.InternalMessageInfo

func (m *TeStatsBagS2LT) GetSubGroupOriginator() string {
	if m != nil {
		return m.SubGroupOriginator
	}
	return ""
}

func (m *TeStatsBagS2LT) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

func (m *TeStatsBagS2LT) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *TeStatsBagS2LT) GetStatistics() *TeStatsSigT {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func init() {
	proto.RegisterType((*TeStatsBagLspT_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.tails_counters.tails_counter.te_stats_bag_lsp_t_KEYS")
	proto.RegisterType((*TeStatsBagLspT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.tails_counters.tails_counter.te_stats_bag_lsp_t")
	proto.RegisterType((*TeStatsSigT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.tails_counters.tails_counter.te_stats_sig_t")
	proto.RegisterType((*TeStatsBagS2LT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.tails_counters.tails_counter.te_stats_bag_s2l_t")
}

func init() {
	proto.RegisterFile("te_stats_bag_lsp_t.proto", fileDescriptor_te_stats_bag_lsp_t_a15aa9202c376e6d)
}

var fileDescriptor_te_stats_bag_lsp_t_a15aa9202c376e6d = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x96, 0xcf, 0x72, 0xd3, 0x3a,
	0x14, 0xc6, 0x27, 0x4d, 0xdb, 0xdb, 0x28, 0x4d, 0x9a, 0x2a, 0xfd, 0xe3, 0xb9, 0x77, 0x71, 0x43,
	0x66, 0x98, 0xc9, 0x02, 0x52, 0x26, 0x65, 0xc5, 0xae, 0x94, 0x0e, 0xd3, 0x81, 0xe1, 0x8f, 0x5b,
	0x16, 0xb0, 0xd1, 0x28, 0xb1, 0xc6, 0x35, 0x75, 0x65, 0x73, 0x74, 0x1c, 0xdc, 0x57, 0xe0, 0x01,
	0x78, 0x84, 0xbe, 0x10, 0x4b, 0x5e, 0x86, 0xb1, 0x64, 0xa5, 0x76, 0xec, 0x2c, 0x81, 0x5d, 0x73,
	0xbe, 0xf3, 0x93, 0xce, 0xa7, 0xcf, 0x23, 0x95, 0x38, 0x28, 0x98, 0x42, 0x8e, 0x8a, 0x4d, 0xb9,
	0xcf, 0x42, 0x15, 0x33, 0x1c, 0xc7, 0x10, 0x61, 0x44, 0x3f, 0xcd, 0x02, 0x35, 0x8b, 0x58, 0x10,
	0x29, 0x96, 0x02, 0xbb, 0x89, 0x43, 0xc5, 0x50, 0xb0, 0x28, 0x16, 0x30, 0xb6, 0x3f, 0x14, 0x72,
	0xe9, 0x4d, 0x6f, 0xc7, 0x2a, 0xf0, 0x25, 0x0f, 0xc3, 0x40, 0xfa, 0x6c, 0x16, 0x25, 0x12, 0x05,
	0xa8, 0x31, 0xf2, 0x20, 0x54, 0x2b, 0x7e, 0x0e, 0x7f, 0x34, 0xc8, 0x61, 0x75, 0x63, 0xf6, 0xea,
	0xec, 0xe3, 0x05, 0x7d, 0x48, 0xba, 0x2a, 0x4a, 0x60, 0x26, 0x18, 0xf7, 0x3c, 0x10, 0x4a, 0x39,
	0x8d, 0x41, 0x63, 0xd4, 0x72, 0x3b, 0xa6, 0x7a, 0x62, 0x8a, 0xf4, 0x88, 0xf4, 0x3d, 0xa1, 0x30,
	0x90, 0x1c, 0x83, 0x48, 0x2e, 0x7a, 0xd7, 0x74, 0x2f, 0x2d, 0x48, 0x16, 0xf8, 0x8f, 0xb4, 0x30,
	0x91, 0x52, 0x84, 0x2c, 0xf0, 0x9c, 0xe6, 0xa0, 0x31, 0xea, 0xb8, 0x5b, 0xa6, 0x70, 0xee, 0xd1,
	0x47, 0x84, 0x8a, 0x14, 0x85, 0xf4, 0x84, 0xc7, 0xee, 0xbb, 0xd6, 0xf5, 0x62, 0x3d, 0xab, 0x5c,
	0xda, 0xee, 0x7d, 0xb2, 0x99, 0x0d, 0x1c, 0x78, 0xce, 0x86, 0x5e, 0x67, 0x23, 0x54, 0xf1, 0xb9,
	0x37, 0xfc, 0xb9, 0x46, 0x68, 0xd5, 0x15, 0xfd, 0x9f, 0xb4, 0xf3, 0x25, 0x25, 0xbf, 0x11, 0xce,
	0x44, 0x2f, 0x4a, 0x4c, 0xe9, 0x0d, 0xbf, 0x11, 0xf4, 0x5b, 0x83, 0x90, 0x0c, 0x0a, 0x14, 0x06,
	0x33, 0xe5, 0x1c, 0x0f, 0x1a, 0xa3, 0xf6, 0xe4, 0xf3, 0xf8, 0xf7, 0x9d, 0xff, 0x78, 0x31, 0xa5,
	0x0a, 0x7c, 0x86, 0x6e, 0x61, 0x77, 0xfa, 0xbd, 0x41, 0x76, 0xd4, 0x84, 0x85, 0xac, 0x30, 0xd1,
	0xd3, 0x41, 0x73, 0xd4, 0x9e, 0xc8, 0x3f, 0x32, 0x51, 0x76, 0x6e, 0x6a, 0x12, 0x32, 0x74, 0x3b,
	0x6a, 0xf2, 0xfa, 0x62, 0x31, 0xc4, 0xf0, 0xae, 0x45, 0xba, 0xe5, 0xb9, 0xe9, 0x88, 0xf4, 0x30,
	0x65, 0x62, 0x2e, 0x24, 0xb2, 0x44, 0x5e, 0xcb, 0xe8, 0xab, 0xd4, 0x1f, 0x4b, 0xc7, 0xed, 0x62,
	0x7a, 0x96, 0x95, 0x3f, 0x98, 0x2a, 0x3d, 0x22, 0x7b, 0x98, 0xb2, 0x98, 0xe3, 0x15, 0x9b, 0x81,
	0xe0, 0x28, 0x0c, 0xa5, 0x3f, 0x97, 0x8e, 0xbb, 0x8b, 0xe9, 0x3b, 0x8e, 0x57, 0xa7, 0x5a, 0xd1,
	0x5c, 0x09, 0xb8, 0xe2, 0xd2, 0xb7, 0x40, 0xb3, 0x04, 0x68, 0xa5, 0x02, 0x78, 0x22, 0x14, 0x8b,
	0x1d, 0xd6, 0x8b, 0xc0, 0x0b, 0xad, 0x18, 0xe0, 0x31, 0xe9, 0x5b, 0x40, 0x00, 0x44, 0x90, 0xf7,
	0x9b, 0x2f, 0xaa, 0x67, 0xfa, 0xcf, 0x32, 0xa1, 0xb8, 0x3e, 0x08, 0x35, 0x2f, 0x3b, 0xd8, 0xb4,
	0xeb, 0xbb, 0x42, 0xcd, 0xab, 0x0e, 0x0c, 0x50, 0x74, 0xf0, 0x4f, 0x09, 0xa8, 0x38, 0xd0, 0x40,
	0xc9, 0xc1, 0x56, 0x11, 0xa8, 0x3a, 0xd0, 0x40, 0xd1, 0x41, 0xcb, 0x3a, 0xc8, 0xfa, 0x0b, 0x0e,
	0x9e, 0x91, 0x7f, 0xad, 0x61, 0x10, 0x62, 0xce, 0x43, 0xf6, 0x25, 0x11, 0x70, 0x9b, 0x53, 0x44,
	0x53, 0x07, 0xc6, 0xb7, 0xab, 0xf5, 0xf7, 0x99, 0x6c, 0xd8, 0x11, 0xe9, 0xc1, 0x72, 0xd2, 0x6d,
	0x93, 0x34, 0x54, 0x92, 0x86, 0xba, 0xa4, 0xb7, 0x8d, 0x0b, 0xa8, 0x4b, 0x1a, 0xea, 0x92, 0xee,
	0x94, 0x80, 0xf2, 0x39, 0x41, 0x5d, 0xd2, 0xdd, 0x22, 0xb0, 0x74, 0x4e, 0x50, 0x93, 0xf4, 0x8e,
	0x39, 0x27, 0xa8, 0x49, 0x1a, 0xea, 0x92, 0xee, 0xd9, 0xf5, 0x6b, 0x92, 0x86, 0xba, 0xa4, 0x77,
	0x4b, 0x40, 0xc5, 0x41, 0x35, 0x69, 0x5a, 0x04, 0xaa, 0x0e, 0x2a, 0x49, 0xf7, 0xad, 0x83, 0x6a,
	0xd2, 0xb0, 0x3a, 0xe9, 0x3d, 0x93, 0x34, 0xd4, 0x27, 0x7d, 0x4c, 0x0e, 0x30, 0x65, 0x53, 0x3e,
	0xbb, 0x4e, 0x62, 0xc6, 0x55, 0x76, 0x91, 0xe4, 0xdc, 0xbe, 0xe6, 0xfa, 0x98, 0x3e, 0xd7, 0xe2,
	0x89, 0xd6, 0x8a, 0x1b, 0x2e, 0x41, 0x60, 0xc7, 0x3c, 0xb0, 0x1b, 0x96, 0x40, 0xc8, 0x87, 0xcd,
	0xae, 0xfe, 0xec, 0x0f, 0xc5, 0x30, 0x42, 0x1e, 0x9a, 0x1b, 0xc9, 0x39, 0x34, 0xd6, 0x8c, 0x72,
	0x99, 0x09, 0xa7, 0x59, 0x9d, 0x3e, 0x20, 0xdb, 0x79, 0xb7, 0xe9, 0x73, 0x74, 0x5f, 0xdb, 0xd4,
	0x74, 0xcb, 0xf0, 0x6e, 0xf9, 0x19, 0xd0, 0xd7, 0x19, 0x7d, 0x42, 0xf6, 0x54, 0x32, 0x65, 0x3e,
	0x44, 0x49, 0xcc, 0x22, 0x08, 0xfc, 0xec, 0x7d, 0x8a, 0x20, 0x7f, 0xdd, 0xa8, 0x4a, 0xa6, 0x2f,
	0x33, 0xe9, 0xed, 0x42, 0xa1, 0x03, 0xb2, 0x7d, 0x4f, 0x04, 0x5e, 0x7e, 0x59, 0x11, 0xdb, 0x79,
	0xee, 0xad, 0x7a, 0x04, 0x9b, 0x2b, 0x1f, 0xc1, 0xa5, 0xa7, 0x66, 0xfd, 0x6f, 0x3e, 0x35, 0xd3,
	0x4d, 0xfd, 0x8f, 0xc6, 0xf1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x20, 0x77, 0x4e, 0x84,
	0x08, 0x00, 0x00,
}
