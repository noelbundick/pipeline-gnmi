// Code generated by protoc-gen-go. DO NOT EDIT.
// source: te_stats_sig_t.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_signalling_counters_signalling_summary

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

// Send-Recv count for TE Signaling
type TeStatsSigT_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsSigT_KEYS) Reset()         { *m = TeStatsSigT_KEYS{} }
func (m *TeStatsSigT_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigT_KEYS) ProtoMessage()    {}
func (*TeStatsSigT_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_sig_t_d3d47aac80ce22f1, []int{0}
}
func (m *TeStatsSigT_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsSigT_KEYS.Unmarshal(m, b)
}
func (m *TeStatsSigT_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsSigT_KEYS.Marshal(b, m, deterministic)
}
func (dst *TeStatsSigT_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsSigT_KEYS.Merge(dst, src)
}
func (m *TeStatsSigT_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeStatsSigT_KEYS.Size(m)
}
func (m *TeStatsSigT_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsSigT_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsSigT_KEYS proto.InternalMessageInfo

type TeStatsSigT struct {
	// Unknown TX events
	TxEventUnknown uint32 `protobuf:"varint,50,opt,name=tx_event_unknown,json=txEventUnknown" json:"tx_event_unknown,omitempty"`
	// TX Path Create event
	TxPathCreateEvent uint32 `protobuf:"varint,51,opt,name=tx_path_create_event,json=txPathCreateEvent" json:"tx_path_create_event,omitempty"`
	// TX Path Change event
	TxPathChangeEvent uint32 `protobuf:"varint,52,opt,name=tx_path_change_event,json=txPathChangeEvent" json:"tx_path_change_event,omitempty"`
	// TX Path Delete event
	TxPathDeleteEvent uint32 `protobuf:"varint,53,opt,name=tx_path_delete_event,json=txPathDeleteEvent" json:"tx_path_delete_event,omitempty"`
	// TX Path Error event
	TxPathErrorEvent uint32 `protobuf:"varint,54,opt,name=tx_path_error_event,json=txPathErrorEvent" json:"tx_path_error_event,omitempty"`
	// TX Resv Create event
	TxResvCreateEvent uint32 `protobuf:"varint,55,opt,name=tx_resv_create_event,json=txResvCreateEvent" json:"tx_resv_create_event,omitempty"`
	// TX Resv Change event
	TxResvChangeEvent uint32 `protobuf:"varint,56,opt,name=tx_resv_change_event,json=txResvChangeEvent" json:"tx_resv_change_event,omitempty"`
	// TX Resv Delete event
	TxResvDeleteEvent uint32 `protobuf:"varint,57,opt,name=tx_resv_delete_event,json=txResvDeleteEvent" json:"tx_resv_delete_event,omitempty"`
	// TX Resv Error event
	TxResvErrorEvent uint32 `protobuf:"varint,58,opt,name=tx_resv_error_event,json=txResvErrorEvent" json:"tx_resv_error_event,omitempty"`
	// TX Path Reeval Query event
	TxPathReevalQueryEvent uint32 `protobuf:"varint,59,opt,name=tx_path_reeval_query_event,json=txPathReevalQueryEvent" json:"tx_path_reeval_query_event,omitempty"`
	// RX Unknown events
	RxEventUnknown uint32 `protobuf:"varint,60,opt,name=rx_event_unknown,json=rxEventUnknown" json:"rx_event_unknown,omitempty"`
	// RX Path Create event
	RxPathCreateEvent uint32 `protobuf:"varint,61,opt,name=rx_path_create_event,json=rxPathCreateEvent" json:"rx_path_create_event,omitempty"`
	// RX Path Change event
	RxPathChangeEvent uint32 `protobuf:"varint,62,opt,name=rx_path_change_event,json=rxPathChangeEvent" json:"rx_path_change_event,omitempty"`
	// RX Path Delete event
	RxPathDeleteEvent uint32 `protobuf:"varint,63,opt,name=rx_path_delete_event,json=rxPathDeleteEvent" json:"rx_path_delete_event,omitempty"`
	// RX Path Error event
	RxPathErrorEvent uint32 `protobuf:"varint,64,opt,name=rx_path_error_event,json=rxPathErrorEvent" json:"rx_path_error_event,omitempty"`
	// RX Resv Create event
	RxResvCreateEvent uint32 `protobuf:"varint,65,opt,name=rx_resv_create_event,json=rxResvCreateEvent" json:"rx_resv_create_event,omitempty"`
	// RX Resv Change event
	RxResvChangeEvent uint32 `protobuf:"varint,66,opt,name=rx_resv_change_event,json=rxResvChangeEvent" json:"rx_resv_change_event,omitempty"`
	// RX Resv Delete event
	RxResvDeleteEvent uint32 `protobuf:"varint,67,opt,name=rx_resv_delete_event,json=rxResvDeleteEvent" json:"rx_resv_delete_event,omitempty"`
	// RX Resv Error event
	RxResvErrorEvent uint32 `protobuf:"varint,68,opt,name=rx_resv_error_event,json=rxResvErrorEvent" json:"rx_resv_error_event,omitempty"`
	// RX Path Reeval Query event
	RxPathReevalQueryEvent uint32 `protobuf:"varint,69,opt,name=rx_path_reeval_query_event,json=rxPathReevalQueryEvent" json:"rx_path_reeval_query_event,omitempty"`
	// Backup Assign event
	TxBackupAssignEvent uint32 `protobuf:"varint,70,opt,name=tx_backup_assign_event,json=txBackupAssignEvent" json:"tx_backup_assign_event,omitempty"`
	// Error on Backup Assign event
	RxBackupAssignErrEvent uint32 `protobuf:"varint,71,opt,name=rx_backup_assign_err_event,json=rxBackupAssignErrEvent" json:"rx_backup_assign_err_event,omitempty"`
	// Total TE Signalling event count
	EventsTotalCount uint32 `protobuf:"varint,72,opt,name=events_total_count,json=eventsTotalCount" json:"events_total_count,omitempty"`
	// TE Signaling event count
	EventsCount          uint32   `protobuf:"varint,73,opt,name=events_count,json=eventsCount" json:"events_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsSigT) Reset()         { *m = TeStatsSigT{} }
func (m *TeStatsSigT) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigT) ProtoMessage()    {}
func (*TeStatsSigT) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_stats_sig_t_d3d47aac80ce22f1, []int{1}
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

func init() {
	proto.RegisterType((*TeStatsSigT_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signalling_summary.te_stats_sig_t_KEYS")
	proto.RegisterType((*TeStatsSigT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signalling_summary.te_stats_sig_t")
}

func init() {
	proto.RegisterFile("te_stats_sig_t.proto", fileDescriptor_te_stats_sig_t_d3d47aac80ce22f1)
}

var fileDescriptor_te_stats_sig_t_d3d47aac80ce22f1 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x4b, 0x8f, 0xd3, 0x30,
	0x14, 0x85, 0xc5, 0x06, 0x09, 0x03, 0xa3, 0x92, 0x0e, 0xa3, 0x8a, 0x15, 0xcc, 0xaa, 0x0b, 0x28,
	0x12, 0xe5, 0x39, 0x3c, 0x67, 0x3a, 0xe5, 0x21, 0x36, 0x43, 0x81, 0x05, 0xab, 0x2b, 0x37, 0x63,
	0xb5, 0x51, 0x53, 0x3b, 0x5c, 0xdf, 0x84, 0xf4, 0x0f, 0xf1, 0x3b, 0x51, 0xec, 0x18, 0xd9, 0xb5,
	0x67, 0x99, 0x7b, 0xce, 0xe7, 0x9b, 0xe3, 0x23, 0x99, 0x1d, 0x92, 0x00, 0x4d, 0x9c, 0x34, 0xe8,
	0x62, 0x05, 0x34, 0xa9, 0x50, 0x91, 0xca, 0x2e, 0xf2, 0x42, 0xe7, 0x0a, 0x0a, 0xa5, 0xa1, 0x45,
	0xd8, 0x56, 0xa5, 0x06, 0x12, 0xa0, 0x2a, 0x81, 0x13, 0xf7, 0xa1, 0x89, 0xcb, 0xcb, 0xe5, 0x6e,
	0xa2, 0x8b, 0x95, 0xe4, 0x65, 0x59, 0xc8, 0x15, 0xe4, 0xaa, 0x96, 0x24, 0x50, 0xfb, 0x33, 0x5d,
	0x6f, 0xb7, 0x1c, 0x77, 0xc7, 0x77, 0xd9, 0x30, 0xdc, 0x04, 0x5f, 0xe7, 0xbf, 0xbe, 0x1f, 0xff,
	0xbd, 0xc1, 0x0e, 0xc2, 0x79, 0x36, 0x66, 0x03, 0x6a, 0x41, 0x34, 0x42, 0x12, 0xd4, 0x72, 0x23,
	0xd5, 0x1f, 0x39, 0x7a, 0x72, 0xff, 0xda, 0xf8, 0xf6, 0xe2, 0x80, 0xda, 0x79, 0x37, 0xfe, 0x69,
	0xa7, 0xd9, 0x63, 0x76, 0x48, 0x2d, 0x54, 0x9c, 0xd6, 0x90, 0xa3, 0xe0, 0x24, 0x2c, 0x35, 0x9a,
	0x1a, 0xf7, 0x1d, 0x6a, 0x2f, 0x38, 0xad, 0x67, 0x46, 0x31, 0x5c, 0x00, 0xac, 0xb9, 0x5c, 0x39,
	0xe0, 0x69, 0x00, 0x18, 0x25, 0x02, 0x2e, 0x45, 0x29, 0xfe, 0x6f, 0x78, 0xe6, 0x03, 0xe7, 0x46,
	0xb1, 0xc0, 0x23, 0x36, 0x74, 0x80, 0x40, 0x54, 0xd8, 0xfb, 0x9f, 0x1b, 0xff, 0xc0, 0xfa, 0xe7,
	0x9d, 0xe0, 0x9f, 0x8f, 0x42, 0x37, 0x61, 0x82, 0x17, 0xee, 0xfc, 0x85, 0xd0, 0x4d, 0x9c, 0xc0,
	0x02, 0x7e, 0x82, 0x97, 0x01, 0x10, 0x25, 0x30, 0x40, 0x90, 0xe0, 0x95, 0x0f, 0xc4, 0x09, 0x0c,
	0xe0, 0x27, 0x38, 0x71, 0x09, 0x3a, 0xbf, 0x97, 0xe0, 0x84, 0xdd, 0x73, 0x81, 0x51, 0x88, 0x86,
	0x97, 0xf0, 0xbb, 0x16, 0xb8, 0xeb, 0xa9, 0xd7, 0x86, 0x3a, 0xb2, 0xb9, 0x17, 0x46, 0xff, 0xd6,
	0xc9, 0x96, 0x1d, 0xb3, 0x01, 0xee, 0x37, 0xfd, 0xc6, 0x36, 0x8d, 0x51, 0xd3, 0x98, 0x6a, 0xfa,
	0xad, 0x4d, 0x81, 0xa9, 0xa6, 0x31, 0xd5, 0xf4, 0xbb, 0x00, 0x08, 0xef, 0x09, 0x53, 0x4d, 0xbf,
	0xf7, 0x81, 0xbd, 0x7b, 0xc2, 0x44, 0xd3, 0x1f, 0xec, 0x3d, 0x61, 0xa2, 0x69, 0x4c, 0x35, 0x7d,
	0xea, 0xce, 0x4f, 0x34, 0x8d, 0xa9, 0xa6, 0xcf, 0x02, 0x20, 0x4a, 0x10, 0x37, 0x3d, 0xf3, 0x81,
	0x38, 0x41, 0xd4, 0xf4, 0xb9, 0x4b, 0x10, 0x37, 0x8d, 0x57, 0x37, 0x3d, 0xb7, 0x4d, 0x63, 0xba,
	0xe9, 0x29, 0x3b, 0xa2, 0x16, 0x96, 0x3c, 0xdf, 0xd4, 0x15, 0x70, 0xdd, 0xbd, 0x0f, 0x3d, 0xf7,
	0xd1, 0x70, 0x43, 0x6a, 0xcf, 0x8c, 0x78, 0x6a, 0x34, 0x7f, 0xe1, 0x1e, 0x84, 0xee, 0x37, 0x3f,
	0xb9, 0x85, 0x01, 0x88, 0xfd, 0xcf, 0x3e, 0x64, 0x99, 0xb1, 0x69, 0x20, 0x45, 0xbc, 0xb4, 0x4f,
	0xd3, 0xe8, 0xb3, 0x8d, 0x66, 0x95, 0x1f, 0x9d, 0x30, 0xeb, 0xe6, 0xd9, 0x03, 0x76, 0xab, 0x77,
	0x5b, 0xdf, 0x17, 0xe3, 0xbb, 0x69, 0x67, 0xc6, 0xb2, 0xbc, 0x6e, 0x1e, 0xc6, 0xe9, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd3, 0xdb, 0x4e, 0x61, 0x30, 0x05, 0x00, 0x00,
}
