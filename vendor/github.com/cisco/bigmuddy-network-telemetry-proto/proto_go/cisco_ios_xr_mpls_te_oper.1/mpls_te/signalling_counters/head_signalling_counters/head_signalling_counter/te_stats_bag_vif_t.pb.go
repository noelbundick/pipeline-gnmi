// Code generated by protoc-gen-go. DO NOT EDIT.
// source: te_stats_bag_vif_t.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_signalling_counters_head_signalling_counters_head_signalling_counter

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

// MPLS TE VIF Statistics
type TeStatsBagVifT_KEYS struct {
	CType                string   `protobuf:"bytes,1,opt,name=c_type,json=cType" json:"c_type,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsBagVifT_KEYS) Reset()         { *m = TeStatsBagVifT_KEYS{} }
func (m *TeStatsBagVifT_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagVifT_KEYS) ProtoMessage()    {}
func (*TeStatsBagVifT_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_bag_vif_t_68277d045fe6c9aa, []int{0}
}
func (m *TeStatsBagVifT_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsBagVifT_KEYS.Unmarshal(m, b)
}
func (m *TeStatsBagVifT_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsBagVifT_KEYS.Marshal(b, m, deterministic)
}
func (dst *TeStatsBagVifT_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsBagVifT_KEYS.Merge(dst, src)
}
func (m *TeStatsBagVifT_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeStatsBagVifT_KEYS.Size(m)
}
func (m *TeStatsBagVifT_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsBagVifT_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsBagVifT_KEYS proto.InternalMessageInfo

func (m *TeStatsBagVifT_KEYS) GetCType() string {
	if m != nil {
		return m.CType
	}
	return ""
}

func (m *TeStatsBagVifT_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

type TeStatsBagVifT struct {
	// Tunnel Name
	TunnelName string `protobuf:"bytes,50,opt,name=tunnel_name,json=tunnelName" json:"tunnel_name,omitempty"`
	// Tunnel Signalled-Name
	TunnelSigName string `protobuf:"bytes,51,opt,name=tunnel_sig_name,json=tunnelSigName" json:"tunnel_sig_name,omitempty"`
	// LSP ID
	LspId uint32 `protobuf:"varint,52,opt,name=lsp_id,json=lspId" json:"lsp_id,omitempty"`
	// VIF stats
	Statistics *TeStatsSigT `protobuf:"bytes,53,opt,name=statistics" json:"statistics,omitempty"`
	// List of Destination Stats
	DestinationStatistics []*TeStatsBagDestT `protobuf:"bytes,54,rep,name=destination_statistics,json=destinationStatistics" json:"destination_statistics,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}           `json:"-"`
	XXX_unrecognized      []byte             `json:"-"`
	XXX_sizecache         int32              `json:"-"`
}

func (m *TeStatsBagVifT) Reset()         { *m = TeStatsBagVifT{} }
func (m *TeStatsBagVifT) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagVifT) ProtoMessage()    {}
func (*TeStatsBagVifT) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_bag_vif_t_68277d045fe6c9aa, []int{1}
}
func (m *TeStatsBagVifT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsBagVifT.Unmarshal(m, b)
}
func (m *TeStatsBagVifT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsBagVifT.Marshal(b, m, deterministic)
}
func (dst *TeStatsBagVifT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsBagVifT.Merge(dst, src)
}
func (m *TeStatsBagVifT) XXX_Size() int {
	return xxx_messageInfo_TeStatsBagVifT.Size(m)
}
func (m *TeStatsBagVifT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsBagVifT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsBagVifT proto.InternalMessageInfo

func (m *TeStatsBagVifT) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *TeStatsBagVifT) GetTunnelSigName() string {
	if m != nil {
		return m.TunnelSigName
	}
	return ""
}

func (m *TeStatsBagVifT) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *TeStatsBagVifT) GetStatistics() *TeStatsSigT {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *TeStatsBagVifT) GetDestinationStatistics() []*TeStatsBagDestT {
	if m != nil {
		return m.DestinationStatistics
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
	return fileDescriptor_te_stats_bag_vif_t_68277d045fe6c9aa, []int{2}
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
	return fileDescriptor_te_stats_bag_vif_t_68277d045fe6c9aa, []int{3}
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

// MPLS TE Destination Statistics
type TeStatsBagDestT struct {
	// Destination address
	DestinationAddress string `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// Destination stats
	Statistics *TeStatsSigT `protobuf:"bytes,2,opt,name=statistics" json:"statistics,omitempty"`
	// List of S2L Stats
	S2LStatistics        []*TeStatsBagS2LT `protobuf:"bytes,3,rep,name=s2_l_statistics,json=s2LStatistics" json:"s2_l_statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TeStatsBagDestT) Reset()         { *m = TeStatsBagDestT{} }
func (m *TeStatsBagDestT) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagDestT) ProtoMessage()    {}
func (*TeStatsBagDestT) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_bag_vif_t_68277d045fe6c9aa, []int{4}
}
func (m *TeStatsBagDestT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsBagDestT.Unmarshal(m, b)
}
func (m *TeStatsBagDestT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsBagDestT.Marshal(b, m, deterministic)
}
func (dst *TeStatsBagDestT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsBagDestT.Merge(dst, src)
}
func (m *TeStatsBagDestT) XXX_Size() int {
	return xxx_messageInfo_TeStatsBagDestT.Size(m)
}
func (m *TeStatsBagDestT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsBagDestT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsBagDestT proto.InternalMessageInfo

func (m *TeStatsBagDestT) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *TeStatsBagDestT) GetStatistics() *TeStatsSigT {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *TeStatsBagDestT) GetS2LStatistics() []*TeStatsBagS2LT {
	if m != nil {
		return m.S2LStatistics
	}
	return nil
}

func init() {
	proto.RegisterType((*TeStatsBagVifT_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.signalling_counters.head_signalling_counters.head_signalling_counter.te_stats_bag_vif_t_KEYS")
	proto.RegisterType((*TeStatsBagVifT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.signalling_counters.head_signalling_counters.head_signalling_counter.te_stats_bag_vif_t")
	proto.RegisterType((*TeStatsSigT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.signalling_counters.head_signalling_counters.head_signalling_counter.te_stats_sig_t")
	proto.RegisterType((*TeStatsBagS2LT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.signalling_counters.head_signalling_counters.head_signalling_counter.te_stats_bag_s2l_t")
	proto.RegisterType((*TeStatsBagDestT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.signalling_counters.head_signalling_counters.head_signalling_counter.te_stats_bag_dest_t")
}

func init() {
	proto.RegisterFile("te_stats_bag_vif_t.proto", fileDescriptor_te_stats_bag_vif_t_68277d045fe6c9aa)
}

var fileDescriptor_te_stats_bag_vif_t_68277d045fe6c9aa = []byte{
	// 767 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0x4f, 0x72, 0xd3, 0x30,
	0x14, 0xc6, 0xc7, 0x4d, 0x1b, 0x1a, 0xa5, 0x69, 0x53, 0xa5, 0x49, 0x33, 0xb0, 0x20, 0x64, 0xc1,
	0x64, 0x01, 0x29, 0x93, 0x02, 0x0b, 0x76, 0xa5, 0x74, 0x98, 0x0e, 0xff, 0xdd, 0xb2, 0x60, 0xa5,
	0x51, 0x6c, 0xd5, 0xf1, 0xd4, 0xb5, 0xcd, 0x93, 0x1c, 0x52, 0xd8, 0x72, 0x06, 0x0e, 0xc0, 0x82,
	0x03, 0xf4, 0x08, 0x9c, 0x8c, 0x91, 0x64, 0xa7, 0x72, 0xec, 0xce, 0xc0, 0xa6, 0x74, 0x19, 0x7d,
	0xef, 0x27, 0xbd, 0x4f, 0x9f, 0x47, 0x2f, 0xa8, 0x2b, 0x18, 0xe1, 0x82, 0x0a, 0x4e, 0xc6, 0xd4,
	0x23, 0x53, 0xff, 0x84, 0x88, 0x61, 0x0c, 0x91, 0x88, 0xf0, 0x89, 0xe3, 0x73, 0x27, 0x22, 0x7e,
	0xc4, 0xc9, 0x0c, 0xc8, 0x59, 0x1c, 0x70, 0x22, 0x18, 0x89, 0x62, 0x06, 0xc3, 0xf4, 0xc7, 0x90,
	0xfb, 0x5e, 0x48, 0x83, 0xc0, 0x0f, 0x3d, 0xe2, 0x44, 0x49, 0x28, 0x18, 0xf0, 0xe1, 0x84, 0x51,
	0x97, 0xfc, 0x83, 0xd0, 0x7f, 0x83, 0xb6, 0x8b, 0x3d, 0x90, 0x57, 0x07, 0x9f, 0x8e, 0x70, 0x1b,
	0x55, 0x1d, 0x22, 0xce, 0x63, 0xd6, 0xb5, 0x7a, 0xd6, 0xa0, 0x66, 0xaf, 0x38, 0xc7, 0xe7, 0x31,
	0xc3, 0x77, 0x50, 0x4d, 0x24, 0x61, 0xc8, 0x02, 0xe2, 0xbb, 0xdd, 0xa5, 0x9e, 0x35, 0x68, 0xd8,
	0xab, 0x7a, 0xe1, 0xd0, 0xed, 0xff, 0xae, 0x20, 0x5c, 0xdc, 0x0f, 0xdf, 0x45, 0xf5, 0x94, 0x09,
	0xe9, 0x19, 0xeb, 0x8e, 0xd4, 0x7e, 0x48, 0x2f, 0xbd, 0xa5, 0x67, 0x0c, 0xdf, 0x47, 0x1b, 0x69,
	0x01, 0xf7, 0x3d, 0x5d, 0xb4, 0xab, 0x8a, 0x1a, 0x7a, 0xf9, 0xc8, 0xf7, 0x54, 0x5d, 0x1b, 0x55,
	0x03, 0x1e, 0xcb, 0x93, 0x1f, 0xab, 0x93, 0x57, 0x02, 0x1e, 0x1f, 0xba, 0xf8, 0x87, 0x85, 0x90,
	0x3c, 0xd3, 0xe7, 0xc2, 0x77, 0x78, 0xf7, 0x49, 0xcf, 0x1a, 0xd4, 0x47, 0xd3, 0xe1, 0xf5, 0xdc,
	0xe1, 0x70, 0x6e, 0x58, 0xf6, 0x2e, 0x6c, 0xa3, 0x13, 0x7c, 0x61, 0xa1, 0x8e, 0xcb, 0xb8, 0xf0,
	0x43, 0x2a, 0xfc, 0x28, 0x24, 0x46, 0x93, 0x4f, 0x7b, 0x95, 0x41, 0x7d, 0xf4, 0xed, 0xda, 0x9b,
	0x94, 0xa9, 0xc8, 0x96, 0x88, 0xb0, 0xdb, 0x46, 0x6b, 0x47, 0xf3, 0xce, 0xfa, 0xbf, 0x6a, 0x68,
	0x3d, 0xef, 0x09, 0x0f, 0x50, 0x53, 0xcc, 0x08, 0x9b, 0xb2, 0x50, 0x90, 0x24, 0x3c, 0x0d, 0xa3,
	0x2f, 0xa1, 0xfa, 0x2a, 0x1a, 0xf6, 0xba, 0x98, 0x1d, 0xc8, 0xe5, 0x8f, 0x7a, 0x15, 0xef, 0xa0,
	0x2d, 0x31, 0x23, 0x31, 0x15, 0x13, 0xe2, 0x00, 0xa3, 0x82, 0x69, 0x2a, 0xfd, 0x52, 0x36, 0xc5,
	0xec, 0x3d, 0x15, 0x93, 0x7d, 0xa5, 0x28, 0x2e, 0x07, 0x4c, 0x68, 0xe8, 0x65, 0x40, 0x25, 0x07,
	0x28, 0xa5, 0x00, 0xb8, 0x2c, 0x60, 0xf3, 0x13, 0x96, 0x4d, 0xe0, 0x85, 0x52, 0x34, 0xf0, 0x10,
	0xb5, 0x32, 0x80, 0x01, 0x44, 0x90, 0xd6, 0xaf, 0xa8, 0xfa, 0xa6, 0xae, 0x3f, 0x90, 0x82, 0xb9,
	0x3f, 0x30, 0x3e, 0xcd, 0x3b, 0xa8, 0x66, 0xfb, 0xdb, 0x8c, 0x4f, 0x8b, 0x0e, 0x34, 0x60, 0x3a,
	0xb8, 0x95, 0x03, 0x0a, 0x0e, 0x14, 0x90, 0x73, 0xb0, 0x6a, 0x02, 0x45, 0x07, 0x0a, 0x30, 0x1d,
	0xd4, 0x32, 0x07, 0xb2, 0xde, 0x70, 0xf0, 0x0c, 0xdd, 0xce, 0x0c, 0x03, 0x63, 0x53, 0x1a, 0x90,
	0xcf, 0x09, 0x83, 0xf3, 0x94, 0x42, 0x8a, 0xea, 0x68, 0xdf, 0xb6, 0xd2, 0x3f, 0x48, 0x59, 0xb3,
	0x03, 0xd4, 0x84, 0xc5, 0xa4, 0xeb, 0x3a, 0x69, 0x28, 0x24, 0x0d, 0x65, 0x49, 0xaf, 0x69, 0x17,
	0x50, 0x96, 0x34, 0x94, 0x25, 0xdd, 0xc8, 0x01, 0xf9, 0x7b, 0x82, 0xb2, 0xa4, 0xd7, 0x4d, 0x60,
	0xe1, 0x9e, 0xa0, 0x24, 0xe9, 0x0d, 0x7d, 0x4f, 0x50, 0x92, 0x34, 0x94, 0x25, 0xdd, 0xcc, 0xf6,
	0x2f, 0x49, 0x1a, 0xca, 0x92, 0xde, 0xcc, 0x01, 0x05, 0x07, 0xc5, 0xa4, 0xb1, 0x09, 0x14, 0x1d,
	0x14, 0x92, 0x6e, 0x65, 0x0e, 0x8a, 0x49, 0xc3, 0xd5, 0x49, 0x6f, 0xe9, 0xa4, 0xa1, 0x3c, 0xe9,
	0x5d, 0xd4, 0x11, 0x33, 0x32, 0xa6, 0xce, 0x69, 0x12, 0x13, 0xca, 0xe5, 0x9b, 0x91, 0x72, 0x6d,
	0xc5, 0xb5, 0xc4, 0xec, 0xb9, 0x12, 0xf7, 0x94, 0x66, 0x1e, 0xb8, 0x00, 0x41, 0xd6, 0x66, 0x27,
	0x3b, 0x30, 0x07, 0x42, 0xda, 0xec, 0x03, 0x84, 0x55, 0x19, 0x27, 0x22, 0x12, 0x34, 0xd0, 0x4f,
	0x53, 0x77, 0x5b, 0x5b, 0xd3, 0xca, 0xb1, 0x14, 0xf6, 0xe5, 0x3a, 0xbe, 0x87, 0xd6, 0xd2, 0x6a,
	0x5d, 0xd7, 0x55, 0x75, 0x75, 0xbd, 0xa6, 0x4a, 0xfa, 0x17, 0x4b, 0x0b, 0xd3, 0x86, 0x8f, 0x02,
	0x22, 0xf0, 0x23, 0xb4, 0xc5, 0x93, 0x31, 0xf1, 0x20, 0x4a, 0x62, 0x12, 0x81, 0xef, 0xc9, 0x27,
	0x2e, 0x82, 0x74, 0x8c, 0x61, 0x9e, 0x8c, 0x5f, 0x4a, 0xe9, 0xdd, 0x5c, 0xc1, 0x3d, 0xb4, 0x76,
	0x49, 0xcc, 0xc7, 0x1a, 0xca, 0x2a, 0x0f, 0x5d, 0xbc, 0x83, 0x5a, 0xe6, 0x3b, 0x4e, 0x5d, 0x17,
	0x18, 0xe7, 0xea, 0x91, 0xaa, 0xd9, 0xd8, 0x90, 0xf6, 0xb4, 0xb2, 0x38, 0x92, 0x96, 0x6f, 0xca,
	0x48, 0xea, 0x7f, 0xaf, 0xa0, 0x56, 0xc9, 0x30, 0xb8, 0xca, 0xa1, 0xf5, 0xb7, 0x0e, 0x97, 0x6e,
	0xcc, 0xd0, 0xfd, 0x69, 0xa1, 0x0d, 0x3e, 0x22, 0x81, 0x39, 0x6d, 0x2b, 0x6a, 0xda, 0x7e, 0xfd,
	0x2f, 0xd3, 0x56, 0x7d, 0x95, 0x76, 0x83, 0x8f, 0x5e, 0x5f, 0x0e, 0xd9, 0x71, 0x55, 0xfd, 0xcf,
	0xdb, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x79, 0xba, 0x2b, 0xf4, 0x03, 0x0a, 0x00, 0x00,
}