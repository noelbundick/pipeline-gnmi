// Code generated by protoc-gen-go. DO NOT EDIT.
// source: te_nsr_issu_status_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_issu_detail

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

// NSR/ISSU sync status information
type TeNsrIssuStatusInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeNsrIssuStatusInfo_KEYS) Reset()         { *m = TeNsrIssuStatusInfo_KEYS{} }
func (m *TeNsrIssuStatusInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeNsrIssuStatusInfo_KEYS) ProtoMessage()    {}
func (*TeNsrIssuStatusInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_nsr_issu_status_info_da8dbefc83966d3d, []int{0}
}
func (m *TeNsrIssuStatusInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeNsrIssuStatusInfo_KEYS.Unmarshal(m, b)
}
func (m *TeNsrIssuStatusInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeNsrIssuStatusInfo_KEYS.Marshal(b, m, deterministic)
}
func (dst *TeNsrIssuStatusInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeNsrIssuStatusInfo_KEYS.Merge(dst, src)
}
func (m *TeNsrIssuStatusInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeNsrIssuStatusInfo_KEYS.Size(m)
}
func (m *TeNsrIssuStatusInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeNsrIssuStatusInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeNsrIssuStatusInfo_KEYS proto.InternalMessageInfo

type TeNsrIssuStatusInfo struct {
	// Process role
	Role string `protobuf:"bytes,50,opt,name=role" json:"role,omitempty"`
	// Sync information for the NSR and ISSU
	SyncStatusInformation *TeSyncStatusInfo `protobuf:"bytes,51,opt,name=sync_status_information,json=syncStatusInformation" json:"sync_status_information,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *TeNsrIssuStatusInfo) Reset()         { *m = TeNsrIssuStatusInfo{} }
func (m *TeNsrIssuStatusInfo) String() string { return proto.CompactTextString(m) }
func (*TeNsrIssuStatusInfo) ProtoMessage()    {}
func (*TeNsrIssuStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_nsr_issu_status_info_da8dbefc83966d3d, []int{1}
}
func (m *TeNsrIssuStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeNsrIssuStatusInfo.Unmarshal(m, b)
}
func (m *TeNsrIssuStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeNsrIssuStatusInfo.Marshal(b, m, deterministic)
}
func (dst *TeNsrIssuStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeNsrIssuStatusInfo.Merge(dst, src)
}
func (m *TeNsrIssuStatusInfo) XXX_Size() int {
	return xxx_messageInfo_TeNsrIssuStatusInfo.Size(m)
}
func (m *TeNsrIssuStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeNsrIssuStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeNsrIssuStatusInfo proto.InternalMessageInfo

func (m *TeNsrIssuStatusInfo) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *TeNsrIssuStatusInfo) GetSyncStatusInformation() *TeSyncStatusInfo {
	if m != nil {
		return m.SyncStatusInformation
	}
	return nil
}

// IDT status information
type TeIdtStatus struct {
	// Ready status
	IsReadyStatus bool `protobuf:"varint,1,opt,name=is_ready_status,json=isReadyStatus" json:"is_ready_status,omitempty"`
	// Not ready reason
	Reason string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	// IDT start timestampe in seconds
	IdtStartTime uint32 `protobuf:"varint,3,opt,name=idt_start_time,json=idtStartTime" json:"idt_start_time,omitempty"`
	// IDT end timestampe in seconds
	IdtEndTime uint32 `protobuf:"varint,4,opt,name=idt_end_time,json=idtEndTime" json:"idt_end_time,omitempty"`
	// Declare ready timestampe in seconds
	DeclareTime uint32 `protobuf:"varint,5,opt,name=declare_time,json=declareTime" json:"declare_time,omitempty"`
	// Withdraw ready timestampe in seconds
	WithdrawTime         uint32   `protobuf:"varint,6,opt,name=withdraw_time,json=withdrawTime" json:"withdraw_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeIdtStatus) Reset()         { *m = TeIdtStatus{} }
func (m *TeIdtStatus) String() string { return proto.CompactTextString(m) }
func (*TeIdtStatus) ProtoMessage()    {}
func (*TeIdtStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_nsr_issu_status_info_da8dbefc83966d3d, []int{2}
}
func (m *TeIdtStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeIdtStatus.Unmarshal(m, b)
}
func (m *TeIdtStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeIdtStatus.Marshal(b, m, deterministic)
}
func (dst *TeIdtStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeIdtStatus.Merge(dst, src)
}
func (m *TeIdtStatus) XXX_Size() int {
	return xxx_messageInfo_TeIdtStatus.Size(m)
}
func (m *TeIdtStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TeIdtStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TeIdtStatus proto.InternalMessageInfo

func (m *TeIdtStatus) GetIsReadyStatus() bool {
	if m != nil {
		return m.IsReadyStatus
	}
	return false
}

func (m *TeIdtStatus) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *TeIdtStatus) GetIdtStartTime() uint32 {
	if m != nil {
		return m.IdtStartTime
	}
	return 0
}

func (m *TeIdtStatus) GetIdtEndTime() uint32 {
	if m != nil {
		return m.IdtEndTime
	}
	return 0
}

func (m *TeIdtStatus) GetDeclareTime() uint32 {
	if m != nil {
		return m.DeclareTime
	}
	return 0
}

func (m *TeIdtStatus) GetWithdrawTime() uint32 {
	if m != nil {
		return m.WithdrawTime
	}
	return 0
}

// Sync pending VIF information
type TeVifPendingInfo struct {
	// Pending reason
	PendingReason string `protobuf:"bytes,1,opt,name=pending_reason,json=pendingReason" json:"pending_reason,omitempty"`
	// Tunnel name
	TunnelName           string   `protobuf:"bytes,2,opt,name=tunnel_name,json=tunnelName" json:"tunnel_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeVifPendingInfo) Reset()         { *m = TeVifPendingInfo{} }
func (m *TeVifPendingInfo) String() string { return proto.CompactTextString(m) }
func (*TeVifPendingInfo) ProtoMessage()    {}
func (*TeVifPendingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_nsr_issu_status_info_da8dbefc83966d3d, []int{3}
}
func (m *TeVifPendingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeVifPendingInfo.Unmarshal(m, b)
}
func (m *TeVifPendingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeVifPendingInfo.Marshal(b, m, deterministic)
}
func (dst *TeVifPendingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeVifPendingInfo.Merge(dst, src)
}
func (m *TeVifPendingInfo) XXX_Size() int {
	return xxx_messageInfo_TeVifPendingInfo.Size(m)
}
func (m *TeVifPendingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeVifPendingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeVifPendingInfo proto.InternalMessageInfo

func (m *TeVifPendingInfo) GetPendingReason() string {
	if m != nil {
		return m.PendingReason
	}
	return ""
}

func (m *TeVifPendingInfo) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

// Sync pending S2L information
type TeS2LPendingInfo struct {
	// Pending reason
	PendingReason string `protobuf:"bytes,1,opt,name=pending_reason,json=pendingReason" json:"pending_reason,omitempty"`
	// Signaled name
	SignaledName string `protobuf:"bytes,2,opt,name=signaled_name,json=signaledName" json:"signaled_name,omitempty"`
	// S2L role
	S2LRole              string   `protobuf:"bytes,3,opt,name=s2_l_role,json=s2LRole" json:"s2_l_role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeS2LPendingInfo) Reset()         { *m = TeS2LPendingInfo{} }
func (m *TeS2LPendingInfo) String() string { return proto.CompactTextString(m) }
func (*TeS2LPendingInfo) ProtoMessage()    {}
func (*TeS2LPendingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_nsr_issu_status_info_da8dbefc83966d3d, []int{4}
}
func (m *TeS2LPendingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeS2LPendingInfo.Unmarshal(m, b)
}
func (m *TeS2LPendingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeS2LPendingInfo.Marshal(b, m, deterministic)
}
func (dst *TeS2LPendingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeS2LPendingInfo.Merge(dst, src)
}
func (m *TeS2LPendingInfo) XXX_Size() int {
	return xxx_messageInfo_TeS2LPendingInfo.Size(m)
}
func (m *TeS2LPendingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeS2LPendingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeS2LPendingInfo proto.InternalMessageInfo

func (m *TeS2LPendingInfo) GetPendingReason() string {
	if m != nil {
		return m.PendingReason
	}
	return ""
}

func (m *TeS2LPendingInfo) GetSignaledName() string {
	if m != nil {
		return m.SignaledName
	}
	return ""
}

func (m *TeS2LPendingInfo) GetS2LRole() string {
	if m != nil {
		return m.S2LRole
	}
	return ""
}

// Sync Status information
type TeSyncIdtInfo struct {
	// Current IDT information
	CurrentIdtInfo *TeIdtStatus `protobuf:"bytes,1,opt,name=current_idt_info,json=currentIdtInfo" json:"current_idt_info,omitempty"`
	// Previous IDT information
	PreviousIdtStatus    *TeIdtStatus `protobuf:"bytes,2,opt,name=previous_idt_status,json=previousIdtStatus" json:"previous_idt_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TeSyncIdtInfo) Reset()         { *m = TeSyncIdtInfo{} }
func (m *TeSyncIdtInfo) String() string { return proto.CompactTextString(m) }
func (*TeSyncIdtInfo) ProtoMessage()    {}
func (*TeSyncIdtInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_nsr_issu_status_info_da8dbefc83966d3d, []int{5}
}
func (m *TeSyncIdtInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSyncIdtInfo.Unmarshal(m, b)
}
func (m *TeSyncIdtInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSyncIdtInfo.Marshal(b, m, deterministic)
}
func (dst *TeSyncIdtInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSyncIdtInfo.Merge(dst, src)
}
func (m *TeSyncIdtInfo) XXX_Size() int {
	return xxx_messageInfo_TeSyncIdtInfo.Size(m)
}
func (m *TeSyncIdtInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSyncIdtInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeSyncIdtInfo proto.InternalMessageInfo

func (m *TeSyncIdtInfo) GetCurrentIdtInfo() *TeIdtStatus {
	if m != nil {
		return m.CurrentIdtInfo
	}
	return nil
}

func (m *TeSyncIdtInfo) GetPreviousIdtStatus() *TeIdtStatus {
	if m != nil {
		return m.PreviousIdtStatus
	}
	return nil
}

// Master Sync status information
type TeSyncStatusMasterInfo struct {
	// Sync status IDT information
	Idt                  *TeSyncIdtInfo `protobuf:"bytes,1,opt,name=idt" json:"idt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TeSyncStatusMasterInfo) Reset()         { *m = TeSyncStatusMasterInfo{} }
func (m *TeSyncStatusMasterInfo) String() string { return proto.CompactTextString(m) }
func (*TeSyncStatusMasterInfo) ProtoMessage()    {}
func (*TeSyncStatusMasterInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_nsr_issu_status_info_da8dbefc83966d3d, []int{6}
}
func (m *TeSyncStatusMasterInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSyncStatusMasterInfo.Unmarshal(m, b)
}
func (m *TeSyncStatusMasterInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSyncStatusMasterInfo.Marshal(b, m, deterministic)
}
func (dst *TeSyncStatusMasterInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSyncStatusMasterInfo.Merge(dst, src)
}
func (m *TeSyncStatusMasterInfo) XXX_Size() int {
	return xxx_messageInfo_TeSyncStatusMasterInfo.Size(m)
}
func (m *TeSyncStatusMasterInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSyncStatusMasterInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeSyncStatusMasterInfo proto.InternalMessageInfo

func (m *TeSyncStatusMasterInfo) GetIdt() *TeSyncIdtInfo {
	if m != nil {
		return m.Idt
	}
	return nil
}

// Sync status slave Information
type TeSyncStatusSlaveInfo struct {
	// Sync status IDT information
	Idt *TeSyncIdtInfo `protobuf:"bytes,1,opt,name=idt" json:"idt,omitempty"`
	// Pending tunnels details
	VifPending []*TeVifPendingInfo `protobuf:"bytes,2,rep,name=vif_pending,json=vifPending" json:"vif_pending,omitempty"`
	// Pending sub-LSPs details
	S2LPending []*TeS2LPendingInfo `protobuf:"bytes,3,rep,name=s2_l_pending,json=s2LPending" json:"s2_l_pending,omitempty"`
	// Tunnels in sync
	InsyncTunnels uint32 `protobuf:"varint,4,opt,name=insync_tunnels,json=insyncTunnels" json:"insync_tunnels,omitempty"`
	// Sub-LSPs in sync
	InsyncSubLsPs uint32 `protobuf:"varint,5,opt,name=insync_sub_ls_ps,json=insyncSubLsPs" json:"insync_sub_ls_ps,omitempty"`
	// Tunnels in sync pending
	PendingTunnels uint32 `protobuf:"varint,6,opt,name=pending_tunnels,json=pendingTunnels" json:"pending_tunnels,omitempty"`
	// Sub-LSPs in sync pending
	PendingSubLsPs       uint32   `protobuf:"varint,7,opt,name=pending_sub_ls_ps,json=pendingSubLsPs" json:"pending_sub_ls_ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeSyncStatusSlaveInfo) Reset()         { *m = TeSyncStatusSlaveInfo{} }
func (m *TeSyncStatusSlaveInfo) String() string { return proto.CompactTextString(m) }
func (*TeSyncStatusSlaveInfo) ProtoMessage()    {}
func (*TeSyncStatusSlaveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_nsr_issu_status_info_da8dbefc83966d3d, []int{7}
}
func (m *TeSyncStatusSlaveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSyncStatusSlaveInfo.Unmarshal(m, b)
}
func (m *TeSyncStatusSlaveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSyncStatusSlaveInfo.Marshal(b, m, deterministic)
}
func (dst *TeSyncStatusSlaveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSyncStatusSlaveInfo.Merge(dst, src)
}
func (m *TeSyncStatusSlaveInfo) XXX_Size() int {
	return xxx_messageInfo_TeSyncStatusSlaveInfo.Size(m)
}
func (m *TeSyncStatusSlaveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSyncStatusSlaveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeSyncStatusSlaveInfo proto.InternalMessageInfo

func (m *TeSyncStatusSlaveInfo) GetIdt() *TeSyncIdtInfo {
	if m != nil {
		return m.Idt
	}
	return nil
}

func (m *TeSyncStatusSlaveInfo) GetVifPending() []*TeVifPendingInfo {
	if m != nil {
		return m.VifPending
	}
	return nil
}

func (m *TeSyncStatusSlaveInfo) GetS2LPending() []*TeS2LPendingInfo {
	if m != nil {
		return m.S2LPending
	}
	return nil
}

func (m *TeSyncStatusSlaveInfo) GetInsyncTunnels() uint32 {
	if m != nil {
		return m.InsyncTunnels
	}
	return 0
}

func (m *TeSyncStatusSlaveInfo) GetInsyncSubLsPs() uint32 {
	if m != nil {
		return m.InsyncSubLsPs
	}
	return 0
}

func (m *TeSyncStatusSlaveInfo) GetPendingTunnels() uint32 {
	if m != nil {
		return m.PendingTunnels
	}
	return 0
}

func (m *TeSyncStatusSlaveInfo) GetPendingSubLsPs() uint32 {
	if m != nil {
		return m.PendingSubLsPs
	}
	return 0
}

// Sync information for the NSR and ISSU based on master or slave role
type TeSyncStatusInfo struct {
	SyncShowType string `protobuf:"bytes,1,opt,name=sync_show_type,json=syncShowType" json:"sync_show_type,omitempty"`
	// Slave sync information
	SlaveSyncInformation *TeSyncStatusSlaveInfo `protobuf:"bytes,2,opt,name=slave_sync_information,json=slaveSyncInformation" json:"slave_sync_information,omitempty"`
	// Master sync information
	MasterSyncInformation *TeSyncStatusMasterInfo `protobuf:"bytes,3,opt,name=master_sync_information,json=masterSyncInformation" json:"master_sync_information,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                `json:"-"`
	XXX_unrecognized      []byte                  `json:"-"`
	XXX_sizecache         int32                   `json:"-"`
}

func (m *TeSyncStatusInfo) Reset()         { *m = TeSyncStatusInfo{} }
func (m *TeSyncStatusInfo) String() string { return proto.CompactTextString(m) }
func (*TeSyncStatusInfo) ProtoMessage()    {}
func (*TeSyncStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_nsr_issu_status_info_da8dbefc83966d3d, []int{8}
}
func (m *TeSyncStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSyncStatusInfo.Unmarshal(m, b)
}
func (m *TeSyncStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSyncStatusInfo.Marshal(b, m, deterministic)
}
func (dst *TeSyncStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSyncStatusInfo.Merge(dst, src)
}
func (m *TeSyncStatusInfo) XXX_Size() int {
	return xxx_messageInfo_TeSyncStatusInfo.Size(m)
}
func (m *TeSyncStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSyncStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeSyncStatusInfo proto.InternalMessageInfo

func (m *TeSyncStatusInfo) GetSyncShowType() string {
	if m != nil {
		return m.SyncShowType
	}
	return ""
}

func (m *TeSyncStatusInfo) GetSlaveSyncInformation() *TeSyncStatusSlaveInfo {
	if m != nil {
		return m.SlaveSyncInformation
	}
	return nil
}

func (m *TeSyncStatusInfo) GetMasterSyncInformation() *TeSyncStatusMasterInfo {
	if m != nil {
		return m.MasterSyncInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*TeNsrIssuStatusInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_nsr_issu_status_info_KEYS")
	proto.RegisterType((*TeNsrIssuStatusInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_nsr_issu_status_info")
	proto.RegisterType((*TeIdtStatus)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_idt_status")
	proto.RegisterType((*TeVifPendingInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_vif_pending_info")
	proto.RegisterType((*TeS2LPendingInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_s2l_pending_info")
	proto.RegisterType((*TeSyncIdtInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_sync_idt_info")
	proto.RegisterType((*TeSyncStatusMasterInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_sync_status_master_info")
	proto.RegisterType((*TeSyncStatusSlaveInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_sync_status_slave_info")
	proto.RegisterType((*TeSyncStatusInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_sync_status_info")
}

func init() {
	proto.RegisterFile("te_nsr_issu_status_info.proto", fileDescriptor_te_nsr_issu_status_info_da8dbefc83966d3d)
}

var fileDescriptor_te_nsr_issu_status_info_da8dbefc83966d3d = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x51, 0x6f, 0xd3, 0x3a,
	0x18, 0x55, 0x96, 0xdd, 0xed, 0xee, 0x6b, 0xd3, 0x6d, 0xde, 0xdd, 0xd6, 0x3b, 0xdd, 0x0b, 0x25,
	0x63, 0x50, 0x1e, 0xe8, 0x43, 0xf6, 0x8a, 0x84, 0x84, 0x34, 0x89, 0x8a, 0x09, 0x8d, 0x74, 0x2f,
	0x3c, 0x20, 0x2b, 0x8b, 0xdd, 0xcd, 0x52, 0x62, 0x47, 0xb6, 0xd3, 0x52, 0x24, 0x90, 0xf8, 0x39,
	0xfc, 0x0a, 0x7e, 0x09, 0xe2, 0x95, 0x9f, 0x81, 0x62, 0x3b, 0x6b, 0x09, 0xda, 0xc3, 0x3a, 0xf1,
	0x16, 0x1f, 0x9d, 0x9e, 0xf3, 0x7d, 0xc9, 0x39, 0x2e, 0xfc, 0xaf, 0x29, 0xe6, 0x4a, 0x62, 0xa6,
	0x54, 0x89, 0x95, 0x4e, 0x74, 0xa9, 0x30, 0xe3, 0x63, 0x31, 0x28, 0xa4, 0xd0, 0x02, 0x3d, 0x4d,
	0x99, 0x4a, 0x05, 0x66, 0x42, 0xe1, 0xf7, 0x12, 0xe7, 0x45, 0xa6, 0xb0, 0xa6, 0x58, 0x14, 0x54,
	0x0e, 0xdc, 0x61, 0x50, 0xfd, 0x72, 0x40, 0xa8, 0x4e, 0x58, 0x16, 0xde, 0x83, 0xff, 0x6e, 0xd0,
	0xc3, 0xaf, 0x4e, 0xde, 0x8e, 0xc2, 0x2f, 0x1e, 0xec, 0xdf, 0x40, 0x40, 0x08, 0x56, 0xa5, 0xc8,
	0x68, 0x37, 0xea, 0x79, 0xfd, 0x8d, 0xd8, 0x3c, 0xa3, 0x0f, 0xb0, 0xaf, 0x66, 0x3c, 0x5d, 0xe4,
	0xc9, 0x3c, 0xd1, 0x4c, 0xf0, 0xee, 0x71, 0xcf, 0xeb, 0xb7, 0xa2, 0x17, 0x83, 0x5b, 0x0d, 0x38,
	0xd0, 0x14, 0x37, 0x05, 0xe3, 0xdd, 0x0a, 0x19, 0x19, 0x60, 0x38, 0x37, 0x08, 0xbf, 0x7b, 0x10,
	0x68, 0x8a, 0x19, 0xd1, 0x8e, 0x8d, 0x1e, 0xc1, 0x26, 0x53, 0x58, 0xd2, 0x84, 0xcc, 0x1c, 0xd4,
	0xf5, 0x7a, 0x5e, 0xff, 0xef, 0x38, 0x60, 0x2a, 0xae, 0x50, 0x2b, 0x82, 0xf6, 0x60, 0x4d, 0xd2,
	0x44, 0x09, 0xde, 0x5d, 0x31, 0xbb, 0xb8, 0x13, 0x7a, 0x08, 0x1d, 0xa7, 0x26, 0x35, 0xd6, 0x2c,
	0xa7, 0x5d, 0xbf, 0xe7, 0xf5, 0x83, 0xb8, 0xcd, 0x88, 0x1e, 0x55, 0xe0, 0x39, 0xcb, 0x29, 0xea,
	0x41, 0x75, 0xc6, 0x94, 0x13, 0xcb, 0x59, 0x35, 0x1c, 0x60, 0x44, 0x9f, 0x70, 0x62, 0x18, 0x0f,
	0xa0, 0x4d, 0x68, 0x9a, 0x25, 0x92, 0x5a, 0xc6, 0x5f, 0x86, 0xd1, 0x72, 0x98, 0xa1, 0x1c, 0x42,
	0x30, 0x65, 0xfa, 0x8a, 0xc8, 0x64, 0x6a, 0x39, 0x6b, 0xd6, 0xa9, 0x06, 0x2b, 0x52, 0xf8, 0x0e,
	0x76, 0x34, 0xc5, 0x13, 0x36, 0xc6, 0x05, 0xe5, 0x84, 0xf1, 0x4b, 0xfb, 0x21, 0x8e, 0xa0, 0x53,
	0x9f, 0xdd, 0x1a, 0x9e, 0x59, 0x23, 0x70, 0x68, 0x6c, 0xb7, 0xb9, 0x0f, 0x2d, 0x5d, 0x72, 0x4e,
	0x33, 0xcc, 0x93, 0x9c, 0xba, 0x55, 0xc1, 0x42, 0xaf, 0x93, 0x9c, 0x86, 0x1f, 0x8d, 0xbc, 0x8a,
	0xb2, 0xa5, 0xe4, 0x0f, 0x21, 0x50, 0xec, 0x92, 0x27, 0x19, 0x25, 0x8b, 0x06, 0xed, 0x1a, 0xac,
	0x2c, 0xd0, 0x01, 0x6c, 0xa8, 0x08, 0x67, 0xd8, 0x04, 0xc7, 0x37, 0x84, 0x75, 0x15, 0x9d, 0xc6,
	0x22, 0xa3, 0xe1, 0x0f, 0x0f, 0xb6, 0xea, 0xcf, 0x5d, 0xbd, 0x50, 0x63, 0x3e, 0x86, 0xad, 0xb4,
	0x94, 0x92, 0x72, 0x7d, 0x8d, 0x19, 0xfb, 0x56, 0xf4, 0xec, 0xf6, 0x49, 0x9a, 0x47, 0x23, 0xee,
	0x38, 0xd5, 0x21, 0xd1, 0x55, 0x86, 0x50, 0x06, 0x3b, 0x85, 0xa4, 0x13, 0x26, 0xaa, 0x90, 0x5d,
	0xd3, 0xcc, 0x0e, 0x77, 0xb5, 0xda, 0xae, 0x85, 0x87, 0x26, 0x35, 0xba, 0x54, 0xa1, 0x80, 0x83,
	0x46, 0xb0, 0xf3, 0x44, 0x69, 0x2a, 0xed, 0xce, 0x6f, 0xc0, 0x67, 0x44, 0xbb, 0x35, 0x9f, 0x2f,
	0x59, 0x98, 0xfa, 0x6d, 0xc5, 0x95, 0x56, 0xf8, 0xcd, 0x87, 0x7f, 0x1b, 0x8e, 0x2a, 0x4b, 0x26,
	0xf4, 0x4f, 0x19, 0xa2, 0x14, 0x5a, 0x0b, 0x39, 0xed, 0xae, 0xf4, 0xfc, 0xe5, 0xca, 0xdf, 0x0c,
	0x7b, 0x0c, 0x13, 0x36, 0x3e, 0xb3, 0x00, 0x22, 0xd0, 0x36, 0x69, 0xaa, 0x5d, 0xfc, 0x65, 0x5d,
	0x9a, 0x99, 0x8f, 0x41, 0x45, 0xa7, 0xb5, 0xcb, 0x11, 0x74, 0x18, 0x37, 0x2b, 0xda, 0xae, 0x28,
	0xd7, 0xf0, 0xc0, 0xa2, 0xe7, 0x16, 0x44, 0x8f, 0x61, 0xcb, 0xd1, 0x54, 0x79, 0x81, 0x33, 0x85,
	0x0b, 0xe5, 0x8a, 0xee, 0x88, 0xa3, 0xf2, 0xe2, 0x54, 0x9d, 0x55, 0xc4, 0xcd, 0xda, 0xab, 0x16,
	0xb4, 0x65, 0xaf, 0x6b, 0x56, 0x2b, 0x3e, 0x81, 0xed, 0x9a, 0x38, 0x97, 0x5c, 0xff, 0x85, 0xea,
	0x34, 0xc3, 0xaf, 0x2b, 0xb6, 0xbb, 0x8d, 0xab, 0xb2, 0xba, 0xc1, 0x2c, 0x76, 0x25, 0xa6, 0x58,
	0xcf, 0x0a, 0xea, 0xba, 0xdb, 0x36, 0x03, 0x5d, 0x89, 0xe9, 0xf9, 0xac, 0xa0, 0xe8, 0x13, 0xec,
	0xd9, 0x34, 0xd8, 0x0f, 0xb9, 0x70, 0x69, 0xdb, 0xfc, 0xbf, 0xbc, 0xdb, 0xa5, 0x3d, 0x4f, 0x5a,
	0xfc, 0x8f, 0x79, 0x1e, 0xcd, 0x78, 0xba, 0x70, 0x73, 0xa3, 0xcf, 0x1e, 0xec, 0xbb, 0x02, 0xfc,
	0x36, 0x81, 0x6f, 0x26, 0x18, 0xde, 0x6d, 0x82, 0x85, 0x76, 0xc5, 0xbb, 0xf6, 0xd0, 0x98, 0xe1,
	0x62, 0xcd, 0xfc, 0x7f, 0x1e, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x18, 0x1b, 0xc3, 0x08, 0x60,
	0x07, 0x00, 0x00,
}
