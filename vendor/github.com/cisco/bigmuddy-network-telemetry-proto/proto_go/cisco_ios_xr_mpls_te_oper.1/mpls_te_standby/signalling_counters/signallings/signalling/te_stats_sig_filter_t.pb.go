// Code generated by protoc-gen-go. DO NOT EDIT.
// source: te_stats_sig_filter_t.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_signalling_counters_signallings_signalling

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

// TE Signaling Filter Data
type TeStatsSigFilterT_KEYS struct {
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

func (m *TeStatsSigFilterT_KEYS) Reset()         { *m = TeStatsSigFilterT_KEYS{} }
func (m *TeStatsSigFilterT_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigFilterT_KEYS) ProtoMessage()    {}
func (*TeStatsSigFilterT_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_sig_filter_t_6ab65823fc8d819e, []int{0}
}
func (m *TeStatsSigFilterT_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsSigFilterT_KEYS.Unmarshal(m, b)
}
func (m *TeStatsSigFilterT_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsSigFilterT_KEYS.Marshal(b, m, deterministic)
}
func (dst *TeStatsSigFilterT_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsSigFilterT_KEYS.Merge(dst, src)
}
func (m *TeStatsSigFilterT_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeStatsSigFilterT_KEYS.Size(m)
}
func (m *TeStatsSigFilterT_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsSigFilterT_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsSigFilterT_KEYS proto.InternalMessageInfo

func (m *TeStatsSigFilterT_KEYS) GetCType() string {
	if m != nil {
		return m.CType
	}
	return ""
}

func (m *TeStatsSigFilterT_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *TeStatsSigFilterT_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *TeStatsSigFilterT_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *TeStatsSigFilterT_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *TeStatsSigFilterT_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *TeStatsSigFilterT_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *TeStatsSigFilterT_KEYS) GetSubGroupOriginator() string {
	if m != nil {
		return m.SubGroupOriginator
	}
	return ""
}

func (m *TeStatsSigFilterT_KEYS) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

type TeStatsSigFilterT struct {
	// TE Signaling Filter data
	TeSignallingFilterData *TeStatsSigFilterDataU `protobuf:"bytes,50,opt,name=te_signalling_filter_data,json=teSignallingFilterData" json:"te_signalling_filter_data,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}               `json:"-"`
	XXX_unrecognized       []byte                 `json:"-"`
	XXX_sizecache          int32                  `json:"-"`
}

func (m *TeStatsSigFilterT) Reset()         { *m = TeStatsSigFilterT{} }
func (m *TeStatsSigFilterT) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigFilterT) ProtoMessage()    {}
func (*TeStatsSigFilterT) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_sig_filter_t_6ab65823fc8d819e, []int{1}
}
func (m *TeStatsSigFilterT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsSigFilterT.Unmarshal(m, b)
}
func (m *TeStatsSigFilterT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsSigFilterT.Marshal(b, m, deterministic)
}
func (dst *TeStatsSigFilterT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsSigFilterT.Merge(dst, src)
}
func (m *TeStatsSigFilterT) XXX_Size() int {
	return xxx_messageInfo_TeStatsSigFilterT.Size(m)
}
func (m *TeStatsSigFilterT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsSigFilterT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsSigFilterT proto.InternalMessageInfo

func (m *TeStatsSigFilterT) GetTeSignallingFilterData() *TeStatsSigFilterDataU {
	if m != nil {
		return m.TeSignallingFilterData
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
	return fileDescriptor_te_stats_sig_filter_t_6ab65823fc8d819e, []int{2}
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
	return fileDescriptor_te_stats_sig_filter_t_6ab65823fc8d819e, []int{3}
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

// MPLS TE LSP Statistics
type TeStatsBagLspT struct {
	// Tunnel Name
	TunnelName string `protobuf:"bytes,1,opt,name=tunnel_name,json=tunnelName" json:"tunnel_name,omitempty"`
	// LSP statistics
	Statistics *TeStatsSigT `protobuf:"bytes,2,opt,name=statistics" json:"statistics,omitempty"`
	// List of S2L Statistics
	S2LStatistics        []*TeStatsBagS2LT `protobuf:"bytes,3,rep,name=s2_l_statistics,json=s2LStatistics" json:"s2_l_statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TeStatsBagLspT) Reset()         { *m = TeStatsBagLspT{} }
func (m *TeStatsBagLspT) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagLspT) ProtoMessage()    {}
func (*TeStatsBagLspT) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_sig_filter_t_6ab65823fc8d819e, []int{4}
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
	return fileDescriptor_te_stats_sig_filter_t_6ab65823fc8d819e, []int{5}
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

// MPLS TE VIF Statistics
type TeStatsBagVifT struct {
	// Tunnel Name
	TunnelName string `protobuf:"bytes,1,opt,name=tunnel_name,json=tunnelName" json:"tunnel_name,omitempty"`
	// Tunnel Signalled-Name
	TunnelSigName string `protobuf:"bytes,2,opt,name=tunnel_sig_name,json=tunnelSigName" json:"tunnel_sig_name,omitempty"`
	// LSP ID
	LspId uint32 `protobuf:"varint,3,opt,name=lsp_id,json=lspId" json:"lsp_id,omitempty"`
	// VIF stats
	Statistics *TeStatsSigT `protobuf:"bytes,4,opt,name=statistics" json:"statistics,omitempty"`
	// List of Destination Stats
	DestinationStatistics []*TeStatsBagDestT `protobuf:"bytes,5,rep,name=destination_statistics,json=destinationStatistics" json:"destination_statistics,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}           `json:"-"`
	XXX_unrecognized      []byte             `json:"-"`
	XXX_sizecache         int32              `json:"-"`
}

func (m *TeStatsBagVifT) Reset()         { *m = TeStatsBagVifT{} }
func (m *TeStatsBagVifT) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagVifT) ProtoMessage()    {}
func (*TeStatsBagVifT) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_sig_filter_t_6ab65823fc8d819e, []int{6}
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

// Union of TE Signalling Filter data
type TeStatsSigFilterDataU struct {
	StatisticsFilter string `protobuf:"bytes,1,opt,name=statistics_filter,json=statisticsFilter" json:"statistics_filter,omitempty"`
	// VIF data
	TeSignallingFilterVif *TeStatsBagVifT `protobuf:"bytes,2,opt,name=te_signalling_filter_vif,json=teSignallingFilterVif" json:"te_signalling_filter_vif,omitempty"`
	// LSP data
	TeSignallingFilterLsp *TeStatsBagLspT `protobuf:"bytes,3,opt,name=te_signalling_filter_lsp,json=teSignallingFilterLsp" json:"te_signalling_filter_lsp,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}        `json:"-"`
	XXX_unrecognized      []byte          `json:"-"`
	XXX_sizecache         int32           `json:"-"`
}

func (m *TeStatsSigFilterDataU) Reset()         { *m = TeStatsSigFilterDataU{} }
func (m *TeStatsSigFilterDataU) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigFilterDataU) ProtoMessage()    {}
func (*TeStatsSigFilterDataU) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_sig_filter_t_6ab65823fc8d819e, []int{7}
}
func (m *TeStatsSigFilterDataU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsSigFilterDataU.Unmarshal(m, b)
}
func (m *TeStatsSigFilterDataU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsSigFilterDataU.Marshal(b, m, deterministic)
}
func (dst *TeStatsSigFilterDataU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsSigFilterDataU.Merge(dst, src)
}
func (m *TeStatsSigFilterDataU) XXX_Size() int {
	return xxx_messageInfo_TeStatsSigFilterDataU.Size(m)
}
func (m *TeStatsSigFilterDataU) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsSigFilterDataU.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsSigFilterDataU proto.InternalMessageInfo

func (m *TeStatsSigFilterDataU) GetStatisticsFilter() string {
	if m != nil {
		return m.StatisticsFilter
	}
	return ""
}

func (m *TeStatsSigFilterDataU) GetTeSignallingFilterVif() *TeStatsBagVifT {
	if m != nil {
		return m.TeSignallingFilterVif
	}
	return nil
}

func (m *TeStatsSigFilterDataU) GetTeSignallingFilterLsp() *TeStatsBagLspT {
	if m != nil {
		return m.TeSignallingFilterLsp
	}
	return nil
}

func init() {
	proto.RegisterType((*TeStatsSigFilterT_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_sig_filter_t_KEYS")
	proto.RegisterType((*TeStatsSigFilterT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_sig_filter_t")
	proto.RegisterType((*TeStatsSigT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_sig_t")
	proto.RegisterType((*TeStatsBagS2LT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_bag_s2l_t")
	proto.RegisterType((*TeStatsBagLspT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_bag_lsp_t")
	proto.RegisterType((*TeStatsBagDestT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_bag_dest_t")
	proto.RegisterType((*TeStatsBagVifT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_bag_vif_t")
	proto.RegisterType((*TeStatsSigFilterDataU)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_sig_filter_data_u")
}

func init() {
	proto.RegisterFile("te_stats_sig_filter_t.proto", fileDescriptor_te_stats_sig_filter_t_6ab65823fc8d819e)
}

var fileDescriptor_te_stats_sig_filter_t_6ab65823fc8d819e = []byte{
	// 973 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x4f, 0x73, 0xe3, 0x34,
	0x14, 0x1f, 0x27, 0xdb, 0x6c, 0xa3, 0x34, 0x6d, 0xaa, 0x34, 0xc1, 0xec, 0x1e, 0x08, 0x99, 0x81,
	0xc9, 0x0c, 0x90, 0x65, 0xbc, 0x37, 0x6e, 0xcb, 0x6e, 0x61, 0x3a, 0x2c, 0xff, 0xdc, 0xc2, 0x0c,
	0x27, 0x8d, 0x62, 0xab, 0xa9, 0x66, 0x5d, 0xdb, 0x48, 0xcf, 0xc1, 0xbd, 0xf3, 0x01, 0xe0, 0x0b,
	0x30, 0xcb, 0x85, 0x0b, 0x9f, 0x83, 0x33, 0x1f, 0x80, 0x0b, 0x1f, 0x85, 0x91, 0x64, 0x27, 0x72,
	0xec, 0xcc, 0xf4, 0xd2, 0xc2, 0x81, 0x5b, 0xf2, 0x7e, 0xbf, 0x9f, 0xf4, 0x9e, 0x7e, 0x2f, 0x4f,
	0x0a, 0x7a, 0x0c, 0x8c, 0x48, 0xa0, 0x20, 0x89, 0xe4, 0x4b, 0x72, 0xc9, 0x23, 0x60, 0x82, 0xc0,
	0x3c, 0x15, 0x09, 0x24, 0xf8, 0x22, 0xe0, 0x32, 0x48, 0x08, 0x4f, 0x24, 0xc9, 0x05, 0xb9, 0x4e,
	0x23, 0x49, 0x80, 0x91, 0x24, 0x65, 0x62, 0x5e, 0x7e, 0x91, 0x40, 0xe3, 0x70, 0x71, 0x33, 0x97,
	0x7c, 0x19, 0xd3, 0x28, 0xe2, 0xf1, 0x92, 0x04, 0x49, 0x16, 0x03, 0x13, 0xd2, 0x8a, 0xd9, 0x9f,
	0xa7, 0x7f, 0xb5, 0xd0, 0xa3, 0xc6, 0x5d, 0xc9, 0x67, 0xa7, 0xdf, 0x9d, 0xe3, 0x11, 0xea, 0x04,
	0x04, 0x6e, 0x52, 0xe6, 0x3a, 0x13, 0x67, 0xd6, 0xf5, 0xf7, 0x82, 0x8b, 0x9b, 0x94, 0xe1, 0xc7,
	0xa8, 0x0b, 0x59, 0x1c, 0xb3, 0x88, 0xf0, 0xd0, 0x6d, 0x4d, 0x9c, 0x59, 0xdf, 0xdf, 0x37, 0x81,
	0xb3, 0x10, 0xbf, 0x8f, 0x30, 0xcb, 0x81, 0xc5, 0x21, 0x0b, 0xc9, 0x86, 0xd5, 0xd6, 0xfa, 0x41,
	0x89, 0x5c, 0x94, 0x6c, 0x17, 0xed, 0xa7, 0x1e, 0xb9, 0x4e, 0x15, 0xe7, 0x81, 0x5e, 0xa9, 0x93,
	0x7a, 0x9f, 0xa7, 0x67, 0xa1, 0xda, 0x3b, 0x92, 0x3a, 0xbe, 0xa7, 0xe3, 0x7b, 0x91, 0x54, 0xe1,
	0x77, 0xd0, 0xa1, 0x4c, 0x32, 0x11, 0x30, 0x42, 0xc3, 0x50, 0x30, 0x29, 0xdd, 0x8e, 0x5e, 0xba,
	0x6f, 0xa2, 0xcf, 0x4c, 0x10, 0x3f, 0x41, 0xc3, 0x90, 0x49, 0xe0, 0x31, 0x05, 0x9e, 0xc4, 0x6b,
	0xee, 0x43, 0xcd, 0xc5, 0x16, 0x54, 0x0a, 0x3e, 0x44, 0x27, 0x32, 0x5b, 0x90, 0xa5, 0x48, 0xb2,
	0x94, 0x24, 0x82, 0x2f, 0x15, 0x9e, 0x08, 0x77, 0xdf, 0x28, 0x64, 0xb6, 0xf8, 0x54, 0x41, 0x5f,
	0xae, 0x11, 0x3c, 0x41, 0x07, 0x1b, 0x05, 0x0f, 0xdd, 0xae, 0x4e, 0x13, 0x95, 0xcc, 0xb3, 0x70,
	0xfa, 0x87, 0x83, 0x46, 0x8d, 0xa7, 0x8b, 0x7f, 0x77, 0xd0, 0x9b, 0x0a, 0xd9, 0x38, 0x55, 0x40,
	0x21, 0x05, 0xea, 0x7a, 0x13, 0x67, 0xd6, 0xf3, 0xd2, 0xf9, 0x5d, 0x58, 0x3e, 0x6f, 0x4a, 0x48,
	0xed, 0x4a, 0x32, 0x7f, 0x0c, 0xec, 0x7c, 0x4d, 0xfc, 0x44, 0x43, 0x2f, 0x28, 0xd0, 0xe9, 0x6f,
	0x5d, 0x74, 0x58, 0x91, 0x01, 0x9e, 0xa1, 0x01, 0xe4, 0x84, 0xad, 0x58, 0x0c, 0x24, 0x8b, 0x5f,
	0xc5, 0xc9, 0x0f, 0xb1, 0xee, 0x91, 0xbe, 0x7f, 0x08, 0xf9, 0xa9, 0x0a, 0x7f, 0x63, 0xa2, 0xf8,
	0x09, 0x3a, 0x81, 0x9c, 0xa4, 0x14, 0xae, 0x48, 0x20, 0x18, 0x05, 0x66, 0x54, 0x45, 0xdf, 0x1c,
	0x43, 0xfe, 0x15, 0x85, 0xab, 0xe7, 0x1a, 0xd1, 0xba, 0x8a, 0xe0, 0x8a, 0xc6, 0xcb, 0x52, 0xd0,
	0xae, 0x08, 0x34, 0x52, 0x13, 0x84, 0x2c, 0x62, 0xeb, 0x1d, 0x1e, 0xd8, 0x82, 0x17, 0x1a, 0x31,
	0x82, 0x0f, 0xd0, 0xb0, 0x14, 0x30, 0x21, 0x12, 0x51, 0xf0, 0x4d, 0x9f, 0x0d, 0x0c, 0xff, 0x54,
	0x01, 0xf6, 0xfa, 0x82, 0xc9, 0x55, 0xb5, 0x82, 0x4e, 0xb9, 0xbe, 0xcf, 0xe4, 0xaa, 0x5e, 0x81,
	0x11, 0xd8, 0x15, 0x3c, 0xac, 0x08, 0x6a, 0x15, 0x68, 0x41, 0xa5, 0x82, 0x7d, 0x5b, 0x50, 0xaf,
	0x40, 0x0b, 0xec, 0x0a, 0xba, 0x65, 0x05, 0x8a, 0x6f, 0x55, 0xf0, 0x11, 0x7a, 0x54, 0x16, 0x2c,
	0x18, 0x5b, 0xd1, 0x88, 0x7c, 0x9f, 0x31, 0x71, 0x53, 0xa8, 0x90, 0x56, 0x8d, 0x4d, 0xdd, 0xbe,
	0xc6, 0xbf, 0x56, 0xb0, 0xd1, 0xce, 0xd0, 0x40, 0x6c, 0x3b, 0xdd, 0x33, 0x4e, 0x8b, 0x9a, 0xd3,
	0xa2, 0xc9, 0xe9, 0x03, 0x53, 0x85, 0x68, 0x72, 0x5a, 0x34, 0x39, 0xdd, 0xaf, 0x08, 0xaa, 0xe7,
	0x24, 0x9a, 0x9c, 0x3e, 0xb4, 0x05, 0x5b, 0xe7, 0x24, 0x1a, 0x9c, 0x3e, 0x32, 0xe7, 0x24, 0x1a,
	0x9c, 0x16, 0x4d, 0x4e, 0x0f, 0xca, 0xf5, 0x1b, 0x9c, 0x16, 0x4d, 0x4e, 0x1f, 0x57, 0x04, 0xb5,
	0x0a, 0xea, 0x4e, 0x63, 0x5b, 0x50, 0xaf, 0xa0, 0xe6, 0xf4, 0xb0, 0xac, 0xa0, 0xee, 0xb4, 0xd8,
	0xed, 0xf4, 0x89, 0x71, 0x5a, 0x34, 0x3b, 0xfd, 0x14, 0x8d, 0x21, 0x27, 0x0b, 0x1a, 0xbc, 0xca,
	0x52, 0x42, 0xa5, 0x9a, 0x19, 0x85, 0x6e, 0xa4, 0x75, 0x43, 0xc8, 0x3f, 0xd6, 0xe0, 0x33, 0x8d,
	0xd9, 0x1b, 0x6e, 0x89, 0x44, 0x99, 0xe6, 0xb8, 0xdc, 0xb0, 0x22, 0x14, 0x45, 0xb2, 0xea, 0xaa,
	0x50, 0x1f, 0x24, 0x81, 0x04, 0x68, 0x64, 0x46, 0x97, 0xfb, 0x86, 0x29, 0xcd, 0x20, 0x17, 0x0a,
	0x78, 0xae, 0xe2, 0xf8, 0x6d, 0x74, 0x50, 0xb0, 0x0d, 0xcf, 0xd5, 0xbc, 0x9e, 0x89, 0x69, 0xca,
	0xf4, 0x75, 0x0b, 0xe1, 0xf5, 0xa0, 0x5a, 0xd0, 0x25, 0x91, 0x5e, 0x44, 0x60, 0xe7, 0x6c, 0x77,
	0x6e, 0x3d, 0xdb, 0x5b, 0xdb, 0xb3, 0x7d, 0xd7, 0x05, 0xd3, 0xde, 0x79, 0xc1, 0xfc, 0xe8, 0x20,
	0xa4, 0x12, 0xe3, 0x12, 0x78, 0x20, 0xf5, 0x70, 0xea, 0x79, 0xe1, 0x3d, 0xcc, 0x78, 0xf0, 0xad,
	0x7d, 0xa7, 0x7f, 0x6e, 0x1f, 0x91, 0xba, 0x64, 0x01, 0xbf, 0x85, 0x7a, 0xc5, 0x65, 0x1d, 0xd3,
	0xeb, 0xf2, 0xba, 0x47, 0x26, 0xf4, 0x05, 0xbd, 0x66, 0xdb, 0xe9, 0xb7, 0xfe, 0x9d, 0xf4, 0xf1,
	0x4f, 0x0e, 0x3a, 0x92, 0x1e, 0x89, 0x88, 0x95, 0x4b, 0x7b, 0xd2, 0x9e, 0xf5, 0xbc, 0xab, 0x3b,
	0xce, 0x65, 0xdd, 0x4e, 0x7e, 0x5f, 0x7a, 0x2f, 0xcf, 0x37, 0x27, 0xfa, 0x77, 0x0b, 0x0d, 0x2b,
	0x2c, 0x65, 0x3e, 0x81, 0x5d, 0x1d, 0xe2, 0xdc, 0xb6, 0x43, 0xfe, 0x3f, 0xe2, 0xf5, 0x11, 0xff,
	0xd2, 0xde, 0x6a, 0xda, 0x15, 0xbf, 0xbc, 0x4d, 0xd3, 0xbe, 0x8b, 0x8e, 0x0a, 0x82, 0x2a, 0x53,
	0x93, 0x5a, 0xe6, 0xb5, 0x68, 0xc2, 0xe7, 0x7c, 0xa9, 0x79, 0x9b, 0xb7, 0x66, 0xdb, 0x7e, 0x6b,
	0xfe, 0x37, 0x7e, 0xb2, 0xf8, 0xb5, 0x83, 0xc6, 0x76, 0x27, 0x59, 0x29, 0xed, 0x69, 0x5f, 0xf8,
	0x3d, 0xf8, 0x62, 0x9a, 0xda, 0x1f, 0x59, 0x89, 0x58, 0x06, 0xfd, 0xdc, 0x6e, 0xfe, 0x1f, 0x61,
	0x1e, 0x96, 0xf8, 0x3d, 0x74, 0xbc, 0x49, 0xba, 0xc0, 0x0a, 0xbb, 0x06, 0x1b, 0xc0, 0xbc, 0x38,
	0xf1, 0xaf, 0x0e, 0x72, 0x1b, 0xdf, 0xc6, 0x2b, 0x7e, 0x59, 0xfc, 0x28, 0xee, 0xa3, 0x11, 0x75,
	0x8b, 0xf9, 0xa3, 0xfa, 0x93, 0xf8, 0x5b, 0x7e, 0xb9, 0x3b, 0xc7, 0x48, 0xa6, 0xba, 0x87, 0xee,
	0x27, 0x47, 0x3d, 0xbb, 0x9b, 0x72, 0x7c, 0x29, 0xd3, 0x45, 0x47, 0xff, 0x71, 0x7c, 0xfa, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x60, 0x04, 0x1e, 0x57, 0x0e, 0x00, 0x00,
}
