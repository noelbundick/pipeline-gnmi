// Code generated by protoc-gen-go. DO NOT EDIT.
// source: te_stats_bag_lsp_t.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_signalling_counters_remote_signalling_counters_remote_signalling_counter

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
	CType                string   `protobuf:"bytes,1,opt,name=c_type,json=cType" json:"c_type,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,3,opt,name=extended_tunnel_id,json=extendedTunnelId" json:"extended_tunnel_id,omitempty"`
	P2MpId               uint32   `protobuf:"varint,4,opt,name=p2_mp_id,json=p2MpId" json:"p2_mp_id,omitempty"`
	LspId                uint32   `protobuf:"varint,5,opt,name=lsp_id,json=lspId" json:"lsp_id,omitempty"`
	SourceAddress        string   `protobuf:"bytes,6,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,7,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	SubGroupOriginator   string   `protobuf:"bytes,8,opt,name=sub_group_originator,json=subGroupOriginator" json:"sub_group_originator,omitempty"`
	SubGroupId           uint32   `protobuf:"varint,9,opt,name=sub_group_id,json=subGroupId" json:"sub_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsBagLspT_KEYS) Reset()         { *m = TeStatsBagLspT_KEYS{} }
func (m *TeStatsBagLspT_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagLspT_KEYS) ProtoMessage()    {}
func (*TeStatsBagLspT_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_bag_lsp_t_a6690e53ca105512, []int{0}
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

func (m *TeStatsBagLspT_KEYS) GetCType() string {
	if m != nil {
		return m.CType
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

func (m *TeStatsBagLspT_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *TeStatsBagLspT_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

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

func (m *TeStatsBagLspT_KEYS) GetSubGroupOriginator() string {
	if m != nil {
		return m.SubGroupOriginator
	}
	return ""
}

func (m *TeStatsBagLspT_KEYS) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
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
	return fileDescriptor_te_stats_bag_lsp_t_a6690e53ca105512, []int{1}
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
	return fileDescriptor_te_stats_bag_lsp_t_a6690e53ca105512, []int{2}
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
	return fileDescriptor_te_stats_bag_lsp_t_a6690e53ca105512, []int{3}
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
	proto.RegisterType((*TeStatsBagLspT_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.remote_signalling_counters.remote_signalling_counter.te_stats_bag_lsp_t_KEYS")
	proto.RegisterType((*TeStatsBagLspT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.remote_signalling_counters.remote_signalling_counter.te_stats_bag_lsp_t")
	proto.RegisterType((*TeStatsSigT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.remote_signalling_counters.remote_signalling_counter.te_stats_sig_t")
	proto.RegisterType((*TeStatsBagS2LT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.remote_signalling_counters.remote_signalling_counter.te_stats_bag_s2l_t")
}

func init() {
	proto.RegisterFile("te_stats_bag_lsp_t.proto", fileDescriptor_te_stats_bag_lsp_t_a6690e53ca105512)
}

var fileDescriptor_te_stats_bag_lsp_t_a6690e53ca105512 = []byte{
	// 783 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0xe5, 0xa4, 0x4d, 0x93, 0x49, 0x93, 0xa6, 0x93, 0xfe, 0x58, 0xb0, 0x20, 0x44, 0x42,
	0xca, 0x02, 0x52, 0xe4, 0xb2, 0x62, 0x57, 0x4a, 0x85, 0x2a, 0xfe, 0xd3, 0xb2, 0x60, 0x35, 0x9a,
	0xd8, 0xa3, 0xd4, 0xaa, 0xe3, 0x31, 0x77, 0xc6, 0xc1, 0xd9, 0xb2, 0x60, 0xc7, 0x53, 0x20, 0xb1,
	0xe5, 0x11, 0x78, 0x0d, 0x1e, 0x07, 0xcd, 0x4c, 0x9c, 0xda, 0xb1, 0x23, 0xc1, 0x06, 0xd8, 0xb5,
	0xf7, 0xdc, 0xcf, 0x73, 0xcf, 0x1c, 0xd7, 0xb7, 0xc8, 0x96, 0x8c, 0x08, 0x49, 0xa5, 0x20, 0x63,
	0x3a, 0x21, 0x81, 0x88, 0x88, 0x1c, 0x46, 0xc0, 0x25, 0xc7, 0xe0, 0xfa, 0xc2, 0xe5, 0xc4, 0xe7,
	0x82, 0x24, 0x40, 0xa6, 0x51, 0x20, 0x88, 0x64, 0x84, 0x47, 0x0c, 0x86, 0xe9, 0x2f, 0x42, 0xd2,
	0xd0, 0x1b, 0xcf, 0x87, 0xc2, 0x9f, 0x84, 0x34, 0x08, 0xfc, 0x70, 0x42, 0x5c, 0x1e, 0x87, 0x92,
	0x81, 0x18, 0x02, 0x9b, 0x72, 0xd5, 0xf2, 0x27, 0x52, 0xff, 0x67, 0x05, 0x1d, 0x16, 0x07, 0x22,
	0xcf, 0xcf, 0xde, 0x5f, 0xe0, 0x7d, 0x54, 0x73, 0x89, 0x9c, 0x47, 0xcc, 0xb6, 0x7a, 0xd6, 0xa0,
	0x31, 0xda, 0x74, 0x2f, 0xe7, 0x11, 0xc3, 0xb7, 0x51, 0x43, 0xc6, 0x61, 0xc8, 0x02, 0xe2, 0x7b,
	0x76, 0xa5, 0x67, 0x0d, 0x5a, 0xa3, 0xba, 0x29, 0x9c, 0x7b, 0xf8, 0x3e, 0xc2, 0x2c, 0x91, 0x2c,
	0xf4, 0x98, 0x47, 0x6e, 0xba, 0xaa, 0x9a, 0xef, 0xa4, 0xca, 0x65, 0xda, 0x6d, 0xa3, 0x7a, 0xe4,
	0x90, 0x69, 0xa4, 0x7a, 0x36, 0xf4, 0x93, 0x6a, 0x91, 0xf3, 0x32, 0x3a, 0xf7, 0xd4, 0xd9, 0x6a,
	0x12, 0xdf, 0xb3, 0x37, 0x75, 0x7d, 0x33, 0x10, 0xaa, 0x7c, 0x0f, 0xb5, 0x05, 0x8f, 0xc1, 0x65,
	0x84, 0x7a, 0x1e, 0x30, 0x21, 0xec, 0x9a, 0x7e, 0x74, 0xcb, 0x54, 0x4f, 0x4c, 0x11, 0x1f, 0xa1,
	0xae, 0xc7, 0x84, 0xf4, 0x43, 0x2a, 0x7d, 0x1e, 0x2e, 0x7b, 0xb7, 0x74, 0x2f, 0xce, 0x48, 0x29,
	0xf0, 0x10, 0xed, 0x89, 0x78, 0x4c, 0x26, 0xc0, 0xe3, 0x88, 0x70, 0xf0, 0x27, 0x4a, 0xe7, 0x60,
	0xd7, 0x0d, 0x21, 0xe2, 0xf1, 0x33, 0x25, 0xbd, 0x5e, 0x2a, 0xb8, 0x87, 0xb6, 0x6f, 0x08, 0xdf,
	0xb3, 0x1b, 0x7a, 0x4c, 0x94, 0x76, 0x9e, 0x7b, 0xfd, 0x2f, 0x55, 0x84, 0x8b, 0x57, 0x8b, 0xef,
	0xa0, 0xe6, 0xe2, 0x62, 0x42, 0x3a, 0x65, 0xb6, 0xa3, 0x4f, 0x40, 0xa6, 0xf4, 0x8a, 0x4e, 0x19,
	0xfe, 0x6a, 0x21, 0xa4, 0x20, 0x5f, 0x48, 0xdf, 0x15, 0xf6, 0x71, 0xcf, 0x1a, 0x34, 0x9d, 0x4f,
	0xd6, 0xf0, 0xef, 0xbf, 0x1d, 0xc3, 0xe5, 0xf8, 0xc2, 0x9f, 0x10, 0x39, 0xca, 0x8c, 0x85, 0xbf,
	0x5b, 0x68, 0x47, 0x38, 0x24, 0x20, 0x99, 0x51, 0x1f, 0xf5, 0xaa, 0x83, 0xa6, 0xf3, 0xf9, 0xdf,
	0x8e, 0xaa, 0x6e, 0x5a, 0x38, 0x01, 0x91, 0xa3, 0x96, 0x70, 0x5e, 0x5c, 0x2c, 0xa7, 0xeb, 0x7f,
	0x6b, 0xa0, 0x76, 0xde, 0x10, 0x1e, 0xa0, 0x8e, 0x4c, 0x08, 0x9b, 0xb1, 0x50, 0x92, 0x38, 0xbc,
	0x0e, 0xf9, 0xc7, 0x50, 0xbf, 0xeb, 0xad, 0x51, 0x5b, 0x26, 0x67, 0xaa, 0xfc, 0xce, 0x54, 0xf1,
	0x11, 0xda, 0x93, 0x09, 0x89, 0xa8, 0xbc, 0x22, 0x2e, 0x30, 0x2a, 0x99, 0xa1, 0x16, 0xef, 0xff,
	0xae, 0x4c, 0xde, 0x50, 0x79, 0x75, 0xaa, 0x15, 0xcd, 0xe5, 0x80, 0x2b, 0x1a, 0x4e, 0x52, 0xa0,
	0x9a, 0x03, 0xb4, 0x52, 0x00, 0x3c, 0x16, 0xb0, 0xe5, 0x09, 0x1b, 0x59, 0xe0, 0xa9, 0x56, 0x0c,
	0xf0, 0x00, 0x75, 0x53, 0x80, 0x01, 0x70, 0x58, 0xf4, 0x9b, 0xbf, 0x97, 0x8e, 0xe9, 0x3f, 0x53,
	0x42, 0xf6, 0xf9, 0xc0, 0xc4, 0x2c, 0xef, 0xa0, 0x96, 0x3e, 0x7f, 0xc4, 0xc4, 0xac, 0xe8, 0xc0,
	0x00, 0x59, 0x07, 0x5b, 0x39, 0xa0, 0xe0, 0x40, 0x03, 0x39, 0x07, 0xf5, 0x2c, 0x50, 0x74, 0xa0,
	0x81, 0xac, 0x83, 0x46, 0xea, 0x40, 0xf5, 0x67, 0x1c, 0x3c, 0x46, 0xb7, 0x52, 0xc3, 0xc0, 0xd8,
	0x8c, 0x06, 0xe4, 0x43, 0xcc, 0x60, 0xbe, 0xa0, 0x90, 0xa6, 0x0e, 0x8c, 0xef, 0x91, 0xd6, 0xdf,
	0x2a, 0xd9, 0xb0, 0x03, 0xd4, 0x81, 0xd5, 0xa4, 0x9b, 0x26, 0x69, 0x28, 0x24, 0x0d, 0x65, 0x49,
	0x6f, 0x1b, 0x17, 0x50, 0x96, 0x34, 0x94, 0x25, 0xdd, 0xca, 0x01, 0xf9, 0x7b, 0x82, 0xb2, 0xa4,
	0xdb, 0x59, 0x60, 0xe5, 0x9e, 0xa0, 0x24, 0xe9, 0x1d, 0x73, 0x4f, 0x50, 0x92, 0x34, 0x94, 0x25,
	0xdd, 0x49, 0x9f, 0x5f, 0x92, 0x34, 0x94, 0x25, 0xbd, 0x9b, 0x03, 0x0a, 0x0e, 0x8a, 0x49, 0xe3,
	0x2c, 0x50, 0x74, 0x50, 0x48, 0xba, 0x9b, 0x3a, 0x28, 0x26, 0x0d, 0xeb, 0x93, 0xde, 0x33, 0x49,
	0x43, 0x79, 0xd2, 0xc7, 0xe8, 0x40, 0x26, 0x64, 0x4c, 0xdd, 0xeb, 0x38, 0x22, 0x54, 0xa8, 0x4f,
	0xc5, 0x82, 0xdb, 0xd7, 0x5c, 0x57, 0x26, 0x4f, 0xb4, 0x78, 0xa2, 0xb5, 0xec, 0x81, 0x2b, 0x10,
	0xa4, 0x63, 0x1e, 0xa4, 0x07, 0xe6, 0x40, 0x58, 0x0c, 0xab, 0x56, 0x9e, 0xfa, 0x41, 0x10, 0xc9,
	0x25, 0x0d, 0xcc, 0x17, 0xc9, 0x3e, 0x34, 0xd6, 0x8c, 0x72, 0xa9, 0x84, 0x53, 0x55, 0xc7, 0x77,
	0xd1, 0xf6, 0xa2, 0xdb, 0xf4, 0xd9, 0xba, 0xaf, 0x69, 0x6a, 0xba, 0xa5, 0xff, 0xa3, 0xb2, 0xb2,
	0x38, 0xf4, 0xe7, 0x6c, 0xed, 0x8e, 0xb2, 0x7e, 0x7b, 0x47, 0x55, 0x56, 0x77, 0xd4, 0xba, 0x45,
	0x59, 0x5d, 0xbb, 0x28, 0x57, 0x96, 0xd3, 0xc6, 0x7f, 0xb9, 0x9c, 0xc6, 0x35, 0xfd, 0x0f, 0xd5,
	0xf1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x27, 0xa1, 0x5f, 0x6c, 0x09, 0x00, 0x00,
}
