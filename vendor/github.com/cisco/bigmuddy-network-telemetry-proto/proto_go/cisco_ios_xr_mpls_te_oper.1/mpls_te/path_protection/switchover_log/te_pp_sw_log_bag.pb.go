// Code generated by protoc-gen-go. DO NOT EDIT.
// source: te_pp_sw_log_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_path_protection_switchover_log

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

// The full log of the path protection switchover events
type TePpSwLogBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TePpSwLogBag_KEYS) Reset()         { *m = TePpSwLogBag_KEYS{} }
func (m *TePpSwLogBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*TePpSwLogBag_KEYS) ProtoMessage()    {}
func (*TePpSwLogBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_pp_sw_log_bag_4518f01dc0f8d3d1, []int{0}
}
func (m *TePpSwLogBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TePpSwLogBag_KEYS.Unmarshal(m, b)
}
func (m *TePpSwLogBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TePpSwLogBag_KEYS.Marshal(b, m, deterministic)
}
func (dst *TePpSwLogBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TePpSwLogBag_KEYS.Merge(dst, src)
}
func (m *TePpSwLogBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_TePpSwLogBag_KEYS.Size(m)
}
func (m *TePpSwLogBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TePpSwLogBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TePpSwLogBag_KEYS proto.InternalMessageInfo

type TePpSwLogBag struct {
	// Total number of path protection switchover events. This could be more than the entries in the array
	PathProtectionSwitchovers uint32 `protobuf:"varint,50,opt,name=path_protection_switchovers,json=pathProtectionSwitchovers" json:"path_protection_switchovers,omitempty"`
	// The maximum delay for a switchover in milliseconds
	MaximumSwitchoverMillisec uint32 `protobuf:"varint,51,opt,name=maximum_switchover_millisec,json=maximumSwitchoverMillisec" json:"maximum_switchover_millisec,omitempty"`
	// The average delay for a switchover in milliseconds
	AverageSwitchoverMillisec uint32 `protobuf:"varint,52,opt,name=average_switchover_millisec,json=averageSwitchoverMillisec" json:"average_switchover_millisec,omitempty"`
	// The array of path protection switchover entries
	PathProtectionSwitchoverEntries []*TePpSwLogEntryBag `protobuf:"bytes,53,rep,name=path_protection_switchover_entries,json=pathProtectionSwitchoverEntries" json:"path_protection_switchover_entries,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}             `json:"-"`
	XXX_unrecognized                []byte               `json:"-"`
	XXX_sizecache                   int32                `json:"-"`
}

func (m *TePpSwLogBag) Reset()         { *m = TePpSwLogBag{} }
func (m *TePpSwLogBag) String() string { return proto.CompactTextString(m) }
func (*TePpSwLogBag) ProtoMessage()    {}
func (*TePpSwLogBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_pp_sw_log_bag_4518f01dc0f8d3d1, []int{1}
}
func (m *TePpSwLogBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TePpSwLogBag.Unmarshal(m, b)
}
func (m *TePpSwLogBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TePpSwLogBag.Marshal(b, m, deterministic)
}
func (dst *TePpSwLogBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TePpSwLogBag.Merge(dst, src)
}
func (m *TePpSwLogBag) XXX_Size() int {
	return xxx_messageInfo_TePpSwLogBag.Size(m)
}
func (m *TePpSwLogBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TePpSwLogBag.DiscardUnknown(m)
}

var xxx_messageInfo_TePpSwLogBag proto.InternalMessageInfo

func (m *TePpSwLogBag) GetPathProtectionSwitchovers() uint32 {
	if m != nil {
		return m.PathProtectionSwitchovers
	}
	return 0
}

func (m *TePpSwLogBag) GetMaximumSwitchoverMillisec() uint32 {
	if m != nil {
		return m.MaximumSwitchoverMillisec
	}
	return 0
}

func (m *TePpSwLogBag) GetAverageSwitchoverMillisec() uint32 {
	if m != nil {
		return m.AverageSwitchoverMillisec
	}
	return 0
}

func (m *TePpSwLogBag) GetPathProtectionSwitchoverEntries() []*TePpSwLogEntryBag {
	if m != nil {
		return m.PathProtectionSwitchoverEntries
	}
	return nil
}

// HASI ID Event
type HasiIdEvent struct {
	// LocalID
	LocalId uint64 `protobuf:"varint,1,opt,name=local_id,json=localId" json:"local_id,omitempty"`
	// RemoteID
	RemoreId uint64 `protobuf:"varint,2,opt,name=remore_id,json=remoreId" json:"remore_id,omitempty"`
	// Time
	Time uint32 `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	// Count
	Count                uint64   `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasiIdEvent) Reset()         { *m = HasiIdEvent{} }
func (m *HasiIdEvent) String() string { return proto.CompactTextString(m) }
func (*HasiIdEvent) ProtoMessage()    {}
func (*HasiIdEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_pp_sw_log_bag_4518f01dc0f8d3d1, []int{2}
}
func (m *HasiIdEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasiIdEvent.Unmarshal(m, b)
}
func (m *HasiIdEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasiIdEvent.Marshal(b, m, deterministic)
}
func (dst *HasiIdEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasiIdEvent.Merge(dst, src)
}
func (m *HasiIdEvent) XXX_Size() int {
	return xxx_messageInfo_HasiIdEvent.Size(m)
}
func (m *HasiIdEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_HasiIdEvent.DiscardUnknown(m)
}

var xxx_messageInfo_HasiIdEvent proto.InternalMessageInfo

func (m *HasiIdEvent) GetLocalId() uint64 {
	if m != nil {
		return m.LocalId
	}
	return 0
}

func (m *HasiIdEvent) GetRemoreId() uint64 {
	if m != nil {
		return m.RemoreId
	}
	return 0
}

func (m *HasiIdEvent) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *HasiIdEvent) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// HASI Slave Object ID events
type HasiSObjIdEvents struct {
	// Create
	Create *HasiIdEvent `protobuf:"bytes,1,opt,name=create" json:"create,omitempty"`
	// Updates
	Updates *HasiIdEvent `protobuf:"bytes,2,opt,name=updates" json:"updates,omitempty"`
	// ApplicationSyncs
	ApplicationSyncs *HasiIdEvent `protobuf:"bytes,3,opt,name=application_syncs,json=applicationSyncs" json:"application_syncs,omitempty"`
	// ApplicationNaks
	ApplicationNaks *HasiIdEvent `protobuf:"bytes,4,opt,name=application_naks,json=applicationNaks" json:"application_naks,omitempty"`
	// SummaryReplaySyncs
	SummaryReplaySyncs *HasiIdEvent `protobuf:"bytes,5,opt,name=summary_replay_syncs,json=summaryReplaySyncs" json:"summary_replay_syncs,omitempty"`
	// SummaryReplayNaks
	SummaryReplayNaks *HasiIdEvent `protobuf:"bytes,6,opt,name=summary_replay_naks,json=summaryReplayNaks" json:"summary_replay_naks,omitempty"`
	// CacheNakRetries
	CacheNakRetries      *HasiIdEvent `protobuf:"bytes,7,opt,name=cache_nak_retries,json=cacheNakRetries" json:"cache_nak_retries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HasiSObjIdEvents) Reset()         { *m = HasiSObjIdEvents{} }
func (m *HasiSObjIdEvents) String() string { return proto.CompactTextString(m) }
func (*HasiSObjIdEvents) ProtoMessage()    {}
func (*HasiSObjIdEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_pp_sw_log_bag_4518f01dc0f8d3d1, []int{3}
}
func (m *HasiSObjIdEvents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasiSObjIdEvents.Unmarshal(m, b)
}
func (m *HasiSObjIdEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasiSObjIdEvents.Marshal(b, m, deterministic)
}
func (dst *HasiSObjIdEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasiSObjIdEvents.Merge(dst, src)
}
func (m *HasiSObjIdEvents) XXX_Size() int {
	return xxx_messageInfo_HasiSObjIdEvents.Size(m)
}
func (m *HasiSObjIdEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_HasiSObjIdEvents.DiscardUnknown(m)
}

var xxx_messageInfo_HasiSObjIdEvents proto.InternalMessageInfo

func (m *HasiSObjIdEvents) GetCreate() *HasiIdEvent {
	if m != nil {
		return m.Create
	}
	return nil
}

func (m *HasiSObjIdEvents) GetUpdates() *HasiIdEvent {
	if m != nil {
		return m.Updates
	}
	return nil
}

func (m *HasiSObjIdEvents) GetApplicationSyncs() *HasiIdEvent {
	if m != nil {
		return m.ApplicationSyncs
	}
	return nil
}

func (m *HasiSObjIdEvents) GetApplicationNaks() *HasiIdEvent {
	if m != nil {
		return m.ApplicationNaks
	}
	return nil
}

func (m *HasiSObjIdEvents) GetSummaryReplaySyncs() *HasiIdEvent {
	if m != nil {
		return m.SummaryReplaySyncs
	}
	return nil
}

func (m *HasiSObjIdEvents) GetSummaryReplayNaks() *HasiIdEvent {
	if m != nil {
		return m.SummaryReplayNaks
	}
	return nil
}

func (m *HasiSObjIdEvents) GetCacheNakRetries() *HasiIdEvent {
	if m != nil {
		return m.CacheNakRetries
	}
	return nil
}

// HASI Slave Object Context
type HasiSObjectCtx struct {
	// ObjectSyncID
	ObjectSyncId uint64 `protobuf:"varint,1,opt,name=object_sync_id,json=objectSyncId" json:"object_sync_id,omitempty"`
	// ObjectType
	ObjectType uint32 `protobuf:"varint,2,opt,name=object_type,json=objectType" json:"object_type,omitempty"`
	// ObjectIdEvents
	ObjectIdEvents       *HasiSObjIdEvents `protobuf:"bytes,3,opt,name=object_id_events,json=objectIdEvents" json:"object_id_events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HasiSObjectCtx) Reset()         { *m = HasiSObjectCtx{} }
func (m *HasiSObjectCtx) String() string { return proto.CompactTextString(m) }
func (*HasiSObjectCtx) ProtoMessage()    {}
func (*HasiSObjectCtx) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_pp_sw_log_bag_4518f01dc0f8d3d1, []int{4}
}
func (m *HasiSObjectCtx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasiSObjectCtx.Unmarshal(m, b)
}
func (m *HasiSObjectCtx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasiSObjectCtx.Marshal(b, m, deterministic)
}
func (dst *HasiSObjectCtx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasiSObjectCtx.Merge(dst, src)
}
func (m *HasiSObjectCtx) XXX_Size() int {
	return xxx_messageInfo_HasiSObjectCtx.Size(m)
}
func (m *HasiSObjectCtx) XXX_DiscardUnknown() {
	xxx_messageInfo_HasiSObjectCtx.DiscardUnknown(m)
}

var xxx_messageInfo_HasiSObjectCtx proto.InternalMessageInfo

func (m *HasiSObjectCtx) GetObjectSyncId() uint64 {
	if m != nil {
		return m.ObjectSyncId
	}
	return 0
}

func (m *HasiSObjectCtx) GetObjectType() uint32 {
	if m != nil {
		return m.ObjectType
	}
	return 0
}

func (m *HasiSObjectCtx) GetObjectIdEvents() *HasiSObjIdEvents {
	if m != nil {
		return m.ObjectIdEvents
	}
	return nil
}

// HASI Master Slave ID Events
type HasiMIdEvents struct {
	// ApplicationSyncs
	ApplicationSyncs *HasiIdEvent `protobuf:"bytes,1,opt,name=application_syncs,json=applicationSyncs" json:"application_syncs,omitempty"`
	// ImplicitDeletes
	ImplicitDeletes *HasiIdEvent `protobuf:"bytes,2,opt,name=implicit_deletes,json=implicitDeletes" json:"implicit_deletes,omitempty"`
	// ApplicationNaks
	ApplicationNaks *HasiIdEvent `protobuf:"bytes,3,opt,name=application_naks,json=applicationNaks" json:"application_naks,omitempty"`
	// SummaryReplaySyncs
	SummaryReplaySyncs *HasiIdEvent `protobuf:"bytes,4,opt,name=summary_replay_syncs,json=summaryReplaySyncs" json:"summary_replay_syncs,omitempty"`
	// SummaryReplayNaks
	SummaryReplayNaks    *HasiIdEvent `protobuf:"bytes,5,opt,name=summary_replay_naks,json=summaryReplayNaks" json:"summary_replay_naks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HasiMIdEvents) Reset()         { *m = HasiMIdEvents{} }
func (m *HasiMIdEvents) String() string { return proto.CompactTextString(m) }
func (*HasiMIdEvents) ProtoMessage()    {}
func (*HasiMIdEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_pp_sw_log_bag_4518f01dc0f8d3d1, []int{5}
}
func (m *HasiMIdEvents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasiMIdEvents.Unmarshal(m, b)
}
func (m *HasiMIdEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasiMIdEvents.Marshal(b, m, deterministic)
}
func (dst *HasiMIdEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasiMIdEvents.Merge(dst, src)
}
func (m *HasiMIdEvents) XXX_Size() int {
	return xxx_messageInfo_HasiMIdEvents.Size(m)
}
func (m *HasiMIdEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_HasiMIdEvents.DiscardUnknown(m)
}

var xxx_messageInfo_HasiMIdEvents proto.InternalMessageInfo

func (m *HasiMIdEvents) GetApplicationSyncs() *HasiIdEvent {
	if m != nil {
		return m.ApplicationSyncs
	}
	return nil
}

func (m *HasiMIdEvents) GetImplicitDeletes() *HasiIdEvent {
	if m != nil {
		return m.ImplicitDeletes
	}
	return nil
}

func (m *HasiMIdEvents) GetApplicationNaks() *HasiIdEvent {
	if m != nil {
		return m.ApplicationNaks
	}
	return nil
}

func (m *HasiMIdEvents) GetSummaryReplaySyncs() *HasiIdEvent {
	if m != nil {
		return m.SummaryReplaySyncs
	}
	return nil
}

func (m *HasiMIdEvents) GetSummaryReplayNaks() *HasiIdEvent {
	if m != nil {
		return m.SummaryReplayNaks
	}
	return nil
}

// HASI Master Object Client Info
type HasiMObjectSlaveInfo struct {
	// SlaveId
	SlaveId uint32 `protobuf:"varint,1,opt,name=slave_id,json=slaveId" json:"slave_id,omitempty"`
	// IdEvents
	IdEvents             *HasiMIdEvents `protobuf:"bytes,2,opt,name=id_events,json=idEvents" json:"id_events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HasiMObjectSlaveInfo) Reset()         { *m = HasiMObjectSlaveInfo{} }
func (m *HasiMObjectSlaveInfo) String() string { return proto.CompactTextString(m) }
func (*HasiMObjectSlaveInfo) ProtoMessage()    {}
func (*HasiMObjectSlaveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_pp_sw_log_bag_4518f01dc0f8d3d1, []int{6}
}
func (m *HasiMObjectSlaveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasiMObjectSlaveInfo.Unmarshal(m, b)
}
func (m *HasiMObjectSlaveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasiMObjectSlaveInfo.Marshal(b, m, deterministic)
}
func (dst *HasiMObjectSlaveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasiMObjectSlaveInfo.Merge(dst, src)
}
func (m *HasiMObjectSlaveInfo) XXX_Size() int {
	return xxx_messageInfo_HasiMObjectSlaveInfo.Size(m)
}
func (m *HasiMObjectSlaveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HasiMObjectSlaveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HasiMObjectSlaveInfo proto.InternalMessageInfo

func (m *HasiMObjectSlaveInfo) GetSlaveId() uint32 {
	if m != nil {
		return m.SlaveId
	}
	return 0
}

func (m *HasiMObjectSlaveInfo) GetIdEvents() *HasiMIdEvents {
	if m != nil {
		return m.IdEvents
	}
	return nil
}

// HASI Master Object Context
type HasiMObjectCtx struct {
	// ObjectSyncID
	ObjectSyncId uint64 `protobuf:"varint,1,opt,name=object_sync_id,json=objectSyncId" json:"object_sync_id,omitempty"`
	// ObjectType
	ObjectType uint32 `protobuf:"varint,2,opt,name=object_type,json=objectType" json:"object_type,omitempty"`
	// Queue
	Queue uint32 `protobuf:"varint,3,opt,name=queue" json:"queue,omitempty"`
	// CreateID
	CreateId uint64 `protobuf:"varint,4,opt,name=create_id,json=createId" json:"create_id,omitempty"`
	// CreateTime
	CreateTime uint32 `protobuf:"varint,5,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// UpdateTime
	UpdateTime uint32 `protobuf:"varint,6,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	// DeleteTime
	DeleteTime uint32 `protobuf:"varint,7,opt,name=delete_time,json=deleteTime" json:"delete_time,omitempty"`
	// DeleteContextLength
	DeleteCtxLenght uint32 `protobuf:"varint,8,opt,name=delete_ctx_lenght,json=deleteCtxLenght" json:"delete_ctx_lenght,omitempty"`
	// SlaveInformation
	SlaveInformation []*HasiMObjectSlaveInfo `protobuf:"bytes,9,rep,name=slave_information,json=slaveInformation" json:"slave_information,omitempty"`
	// HasHistory
	HasHistory bool `protobuf:"varint,10,opt,name=has_history,json=hasHistory" json:"has_history,omitempty"`
	// PreActiveHistory
	PreActiveHistory     *HasiSObjIdEvents `protobuf:"bytes,11,opt,name=pre_active_history,json=preActiveHistory" json:"pre_active_history,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HasiMObjectCtx) Reset()         { *m = HasiMObjectCtx{} }
func (m *HasiMObjectCtx) String() string { return proto.CompactTextString(m) }
func (*HasiMObjectCtx) ProtoMessage()    {}
func (*HasiMObjectCtx) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_pp_sw_log_bag_4518f01dc0f8d3d1, []int{7}
}
func (m *HasiMObjectCtx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasiMObjectCtx.Unmarshal(m, b)
}
func (m *HasiMObjectCtx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasiMObjectCtx.Marshal(b, m, deterministic)
}
func (dst *HasiMObjectCtx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasiMObjectCtx.Merge(dst, src)
}
func (m *HasiMObjectCtx) XXX_Size() int {
	return xxx_messageInfo_HasiMObjectCtx.Size(m)
}
func (m *HasiMObjectCtx) XXX_DiscardUnknown() {
	xxx_messageInfo_HasiMObjectCtx.DiscardUnknown(m)
}

var xxx_messageInfo_HasiMObjectCtx proto.InternalMessageInfo

func (m *HasiMObjectCtx) GetObjectSyncId() uint64 {
	if m != nil {
		return m.ObjectSyncId
	}
	return 0
}

func (m *HasiMObjectCtx) GetObjectType() uint32 {
	if m != nil {
		return m.ObjectType
	}
	return 0
}

func (m *HasiMObjectCtx) GetQueue() uint32 {
	if m != nil {
		return m.Queue
	}
	return 0
}

func (m *HasiMObjectCtx) GetCreateId() uint64 {
	if m != nil {
		return m.CreateId
	}
	return 0
}

func (m *HasiMObjectCtx) GetCreateTime() uint32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *HasiMObjectCtx) GetUpdateTime() uint32 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *HasiMObjectCtx) GetDeleteTime() uint32 {
	if m != nil {
		return m.DeleteTime
	}
	return 0
}

func (m *HasiMObjectCtx) GetDeleteCtxLenght() uint32 {
	if m != nil {
		return m.DeleteCtxLenght
	}
	return 0
}

func (m *HasiMObjectCtx) GetSlaveInformation() []*HasiMObjectSlaveInfo {
	if m != nil {
		return m.SlaveInformation
	}
	return nil
}

func (m *HasiMObjectCtx) GetHasHistory() bool {
	if m != nil {
		return m.HasHistory
	}
	return false
}

func (m *HasiMObjectCtx) GetPreActiveHistory() *HasiSObjIdEvents {
	if m != nil {
		return m.PreActiveHistory
	}
	return nil
}

// Path protection log entry
type TePpSwLogEntryBag struct {
	// The index number of the path protection switch over event
	PathProtectionSwitchoverEventIndex uint32 `protobuf:"varint,1,opt,name=path_protection_switchover_event_index,json=pathProtectionSwitchoverEventIndex" json:"path_protection_switchover_event_index,omitempty"`
	// The ID of the tunnel that experienced switchover
	PathProtectionTunnelId uint32 `protobuf:"varint,2,opt,name=path_protection_tunnel_id,json=pathProtectionTunnelId" json:"path_protection_tunnel_id,omitempty"`
	// The LSP ID from which the traffic was switched over
	FromLspId uint32 `protobuf:"varint,3,opt,name=from_lsp_id,json=fromLspId" json:"from_lsp_id,omitempty"`
	// The LSP ID to which the traffic was switched over
	ToLspId uint32 `protobuf:"varint,4,opt,name=to_lsp_id,json=toLspId" json:"to_lsp_id,omitempty"`
	// The date when the error that caused the switchover was detected. This date is the number of seconds since Jan 1st 1970
	DateOfErrorDetection uint32 `protobuf:"varint,5,opt,name=date_of_error_detection,json=dateOfErrorDetection" json:"date_of_error_detection,omitempty"`
	// The milliseconds offset of the date when the error that caused  the switchover was detected.
	DateOfErrorDetectionMillisec uint32 `protobuf:"varint,6,opt,name=date_of_error_detection_millisec,json=dateOfErrorDetectionMillisec" json:"date_of_error_detection_millisec,omitempty"`
	// The time in milliseconds between the detection of the error and switching the traffic
	SwitchoverDurationMillisec uint32 `protobuf:"varint,7,opt,name=switchover_duration_millisec,json=switchoverDurationMillisec" json:"switchover_duration_millisec,omitempty"`
	// The reason that caused the path protection switchover
	PathProtectionSwitchoverReason string   `protobuf:"bytes,8,opt,name=path_protection_switchover_reason,json=pathProtectionSwitchoverReason" json:"path_protection_switchover_reason,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *TePpSwLogEntryBag) Reset()         { *m = TePpSwLogEntryBag{} }
func (m *TePpSwLogEntryBag) String() string { return proto.CompactTextString(m) }
func (*TePpSwLogEntryBag) ProtoMessage()    {}
func (*TePpSwLogEntryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_te_pp_sw_log_bag_4518f01dc0f8d3d1, []int{8}
}
func (m *TePpSwLogEntryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TePpSwLogEntryBag.Unmarshal(m, b)
}
func (m *TePpSwLogEntryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TePpSwLogEntryBag.Marshal(b, m, deterministic)
}
func (dst *TePpSwLogEntryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TePpSwLogEntryBag.Merge(dst, src)
}
func (m *TePpSwLogEntryBag) XXX_Size() int {
	return xxx_messageInfo_TePpSwLogEntryBag.Size(m)
}
func (m *TePpSwLogEntryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TePpSwLogEntryBag.DiscardUnknown(m)
}

var xxx_messageInfo_TePpSwLogEntryBag proto.InternalMessageInfo

func (m *TePpSwLogEntryBag) GetPathProtectionSwitchoverEventIndex() uint32 {
	if m != nil {
		return m.PathProtectionSwitchoverEventIndex
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetPathProtectionTunnelId() uint32 {
	if m != nil {
		return m.PathProtectionTunnelId
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetFromLspId() uint32 {
	if m != nil {
		return m.FromLspId
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetToLspId() uint32 {
	if m != nil {
		return m.ToLspId
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetDateOfErrorDetection() uint32 {
	if m != nil {
		return m.DateOfErrorDetection
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetDateOfErrorDetectionMillisec() uint32 {
	if m != nil {
		return m.DateOfErrorDetectionMillisec
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetSwitchoverDurationMillisec() uint32 {
	if m != nil {
		return m.SwitchoverDurationMillisec
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetPathProtectionSwitchoverReason() string {
	if m != nil {
		return m.PathProtectionSwitchoverReason
	}
	return ""
}

func init() {
	proto.RegisterType((*TePpSwLogBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.path_protection.switchover_log.te_pp_sw_log_bag_KEYS")
	proto.RegisterType((*TePpSwLogBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.path_protection.switchover_log.te_pp_sw_log_bag")
	proto.RegisterType((*HasiIdEvent)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.path_protection.switchover_log.hasi_id_event")
	proto.RegisterType((*HasiSObjIdEvents)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.path_protection.switchover_log.hasi_s_obj_id_events")
	proto.RegisterType((*HasiSObjectCtx)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.path_protection.switchover_log.hasi_s_object_ctx")
	proto.RegisterType((*HasiMIdEvents)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.path_protection.switchover_log.hasi_m_id_events")
	proto.RegisterType((*HasiMObjectSlaveInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.path_protection.switchover_log.hasi_m_object_slave_info")
	proto.RegisterType((*HasiMObjectCtx)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.path_protection.switchover_log.hasi_m_object_ctx")
	proto.RegisterType((*TePpSwLogEntryBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.path_protection.switchover_log.te_pp_sw_log_entry_bag")
}

func init() {
	proto.RegisterFile("te_pp_sw_log_bag.proto", fileDescriptor_te_pp_sw_log_bag_4518f01dc0f8d3d1)
}

var fileDescriptor_te_pp_sw_log_bag_4518f01dc0f8d3d1 = []byte{
	// 987 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0xd6, 0x64, 0xfd, 0x5b, 0xc6, 0xac, 0xdd, 0x98, 0x64, 0x36, 0x89, 0xb2, 0x66, 0x84, 0xd0,
	0x8a, 0x83, 0x0f, 0x1b, 0x72, 0xe0, 0x82, 0x82, 0xd8, 0x45, 0x58, 0x84, 0x04, 0xcd, 0xae, 0x10,
	0x70, 0x69, 0xf5, 0xce, 0xf4, 0xda, 0xcd, 0xce, 0x4c, 0x4f, 0xba, 0xdb, 0x8e, 0x0d, 0x48, 0x80,
	0x84, 0xc4, 0x89, 0xc7, 0x80, 0xc7, 0xe1, 0x0d, 0xb8, 0x72, 0xe5, 0xce, 0x09, 0xf5, 0x8f, 0x3d,
	0xb6, 0x63, 0xe7, 0x92, 0x9d, 0xe4, 0xe6, 0xae, 0xfa, 0xaa, 0xbe, 0xee, 0xea, 0xfa, 0xda, 0x35,
	0x70, 0x53, 0x51, 0x9c, 0xe7, 0x58, 0x3e, 0xc3, 0x09, 0x1f, 0xe1, 0x0b, 0x32, 0x1a, 0xe4, 0x82,
	0x2b, 0x8e, 0x1e, 0x46, 0x4c, 0x46, 0x1c, 0x33, 0x2e, 0xf1, 0x4c, 0xe0, 0x34, 0x4f, 0x24, 0x56,
	0x14, 0xf3, 0x9c, 0x8a, 0x81, 0x5b, 0x0c, 0x72, 0xa2, 0xc6, 0x58, 0xa3, 0x69, 0xa4, 0x18, 0xcf,
	0x06, 0xf2, 0x19, 0x53, 0xd1, 0x98, 0x4f, 0xa9, 0xd0, 0xb9, 0x82, 0x5b, 0xf0, 0xf6, 0x66, 0x6e,
	0xfc, 0xf9, 0xe9, 0x37, 0x67, 0xc1, 0x7f, 0x37, 0xa0, 0xb3, 0xe9, 0x41, 0x1f, 0xc1, 0x9d, 0x8d,
	0x7c, 0xb8, 0xc8, 0x27, 0xfd, 0xe3, 0xbe, 0x77, 0xd4, 0x0e, 0x0f, 0x34, 0xe4, 0xcb, 0x25, 0xe2,
	0xac, 0x00, 0xe8, 0xf8, 0x94, 0xcc, 0x58, 0x3a, 0x49, 0x57, 0xe2, 0x70, 0xca, 0x92, 0x84, 0x49,
	0x1a, 0xf9, 0xf7, 0x6d, 0xbc, 0x83, 0x14, 0x81, 0x5f, 0x38, 0x80, 0x8e, 0x27, 0x53, 0x2a, 0xc8,
	0x88, 0x6e, 0x8d, 0xff, 0xc0, 0xc6, 0x3b, 0xc8, 0x96, 0xf8, 0x3f, 0x3c, 0x08, 0x76, 0x1f, 0x00,
	0xd3, 0x4c, 0x09, 0x46, 0xa5, 0xff, 0xa0, 0xbf, 0x77, 0xd4, 0x3a, 0xfe, 0x7a, 0xf0, 0xb2, 0xd5,
	0x1d, 0xac, 0x15, 0x50, 0x67, 0x9f, 0xeb, 0x32, 0x86, 0x87, 0xbb, 0x2a, 0x74, 0x6a, 0x37, 0x10,
	0x3c, 0x85, 0xf6, 0x98, 0x48, 0x86, 0x59, 0x8c, 0xe9, 0x94, 0x66, 0x0a, 0x1d, 0x40, 0x23, 0xe1,
	0x11, 0x49, 0x30, 0x8b, 0x7d, 0xaf, 0xef, 0x1d, 0x55, 0xc2, 0xba, 0x59, 0x0f, 0x63, 0x74, 0x07,
	0x9a, 0x82, 0xa6, 0x5c, 0x50, 0xed, 0xbb, 0x61, 0x7c, 0x0d, 0x6b, 0x18, 0xc6, 0x08, 0x41, 0x45,
	0xb1, 0x94, 0xfa, 0x7b, 0xa6, 0x32, 0xe6, 0x37, 0xea, 0x41, 0x35, 0xe2, 0x93, 0x4c, 0xf9, 0x15,
	0x03, 0xb6, 0x8b, 0xe0, 0x9f, 0x1a, 0xf4, 0x0c, 0xa7, 0xc4, 0xfc, 0xe2, 0xbb, 0x25, 0xb3, 0x44,
	0x23, 0xa8, 0x45, 0x82, 0x12, 0x45, 0x0d, 0x71, 0xeb, 0xf8, 0xc9, 0xcb, 0x97, 0x65, 0xed, 0x6c,
	0xa1, 0x4b, 0x8f, 0x18, 0xd4, 0x27, 0x79, 0x4c, 0x14, 0x95, 0xe6, 0x18, 0x25, 0x30, 0x2d, 0xf2,
	0xa3, 0x1f, 0xa1, 0x4b, 0xf2, 0x3c, 0x61, 0x11, 0xb1, 0x2d, 0x30, 0xcf, 0x22, 0x69, 0x6a, 0x54,
	0x02, 0x69, 0x67, 0x85, 0xe9, 0x4c, 0x13, 0xa1, 0xef, 0x61, 0xd5, 0x86, 0x33, 0x72, 0x25, 0xcd,
	0x5d, 0x94, 0x40, 0xbe, 0xbf, 0x42, 0xf4, 0x98, 0x5c, 0x49, 0xf4, 0x8b, 0x07, 0x3d, 0x39, 0x49,
	0x53, 0x22, 0xe6, 0x58, 0xd0, 0x3c, 0x21, 0x73, 0x77, 0xfa, 0x6a, 0x39, 0x1b, 0x40, 0x8e, 0x2c,
	0x34, 0x5c, 0xf6, 0xfc, 0x3f, 0xc1, 0x5b, 0x1b, 0x5b, 0x30, 0x25, 0xa8, 0x95, 0xb3, 0x83, 0xee,
	0xda, 0x0e, 0x4c, 0x11, 0x7e, 0x80, 0x6e, 0x44, 0xa2, 0x31, 0xd5, 0xbc, 0x58, 0x50, 0x2b, 0xfa,
	0x7a, 0x49, 0x37, 0x60, 0x98, 0x1e, 0x93, 0xab, 0xd0, 0xf2, 0x04, 0x7f, 0x7b, 0xd0, 0x2d, 0x84,
	0x46, 0x23, 0x85, 0x23, 0x35, 0x43, 0xef, 0xc2, 0x9b, 0x6e, 0xa5, 0xaf, 0xa3, 0x90, 0xf9, 0x1b,
	0xd6, 0xaa, 0x0b, 0x37, 0x8c, 0xd1, 0x21, 0xb4, 0x1c, 0x4a, 0xcd, 0x73, 0x6a, 0x64, 0xd2, 0x0e,
	0xc1, 0x9a, 0xce, 0xe7, 0x39, 0x45, 0x3f, 0x7b, 0xd0, 0x71, 0x88, 0xa5, 0x82, 0x5d, 0x63, 0x7f,
	0x75, 0x4d, 0x27, 0xdb, 0x78, 0x1f, 0x42, 0xb7, 0xef, 0x61, 0x7c, 0x6a, 0xd6, 0xc1, 0xef, 0x55,
	0xe8, 0x18, 0x60, 0xba, 0xf2, 0x88, 0x6c, 0x15, 0x9c, 0xf7, 0x0a, 0x05, 0xc7, 0x52, 0x6d, 0x63,
	0x0a, 0xc7, 0x34, 0xa1, 0x25, 0x3e, 0x31, 0xfb, 0x0b, 0xa2, 0x13, 0xcb, 0xb3, 0x55, 0xec, 0x7b,
	0xaf, 0x5b, 0xec, 0x95, 0xd7, 0x2e, 0xf6, 0xea, 0xab, 0x12, 0x7b, 0xf0, 0xa7, 0x07, 0xbe, 0xeb,
	0xc7, 0x85, 0xc2, 0x12, 0x32, 0xa5, 0x98, 0x65, 0x97, 0x5c, 0xff, 0xaf, 0xba, 0x95, 0x15, 0x5c,
	0x3b, 0xac, 0x9b, 0xf5, 0x30, 0x46, 0x1c, 0x9a, 0x85, 0x84, 0x6c, 0xb7, 0x84, 0xd7, 0xb4, 0xdd,
	0x15, 0x65, 0x84, 0x0d, 0xb6, 0x10, 0xce, 0x5f, 0x15, 0xf7, 0x30, 0xa4, 0x25, 0x3c, 0x0c, 0x3d,
	0xa8, 0x3e, 0x9d, 0xd0, 0xc9, 0x62, 0x12, 0xb0, 0x0b, 0x3d, 0x3b, 0xd8, 0x3f, 0x5f, 0x9d, 0xd7,
	0x8e, 0x03, 0x0d, 0x6b, 0xb0, 0x39, 0x9d, 0xd3, 0x8c, 0x10, 0x55, 0x9b, 0xd3, 0x9a, 0xce, 0xf5,
	0x20, 0x71, 0x08, 0x2d, 0xfb, 0x87, 0x6a, 0x01, 0x35, 0x0b, 0xb0, 0xa6, 0x05, 0xc0, 0xca, 0xcd,
	0x02, 0xea, 0x16, 0x60, 0x4d, 0x06, 0xf0, 0x3e, 0x74, 0x1d, 0x20, 0x52, 0x33, 0x9c, 0xd0, 0x6c,
	0x34, 0x56, 0x7e, 0xc3, 0xc0, 0xf6, 0xad, 0xe3, 0x13, 0x35, 0x7b, 0x64, 0xcc, 0xe8, 0x37, 0x0f,
	0xba, 0xc5, 0xcd, 0x89, 0xd4, 0xb4, 0xb9, 0xdf, 0x34, 0xa3, 0xda, 0xb7, 0xd7, 0x76, 0x31, 0xcf,
	0xb5, 0x48, 0xd8, 0xb1, 0x0d, 0x51, 0x70, 0xea, 0x63, 0x8d, 0x89, 0xc4, 0x63, 0x26, 0x15, 0x17,
	0x73, 0x1f, 0xfa, 0xde, 0x51, 0x23, 0x84, 0x31, 0x91, 0x9f, 0x59, 0x0b, 0xfa, 0xd5, 0x03, 0x94,
	0x0b, 0x8a, 0x49, 0xa4, 0xd8, 0x94, 0x2e, 0x81, 0xad, 0x52, 0xdf, 0xe1, 0x4e, 0x2e, 0xe8, 0xc7,
	0x86, 0xd0, 0x6d, 0x23, 0xf8, 0x77, 0x6f, 0xe3, 0xc3, 0x61, 0x39, 0x81, 0xa2, 0x10, 0xde, 0x7b,
	0xd1, 0x1c, 0xac, 0x33, 0x62, 0x96, 0xc5, 0x74, 0xe6, 0x54, 0x11, 0xec, 0x9c, 0x58, 0x35, 0x74,
	0xa8, 0x91, 0xe8, 0x43, 0x38, 0xd8, 0xcc, 0xa9, 0x26, 0x59, 0x46, 0x93, 0xc5, 0x60, 0xda, 0x0e,
	0x6f, 0xae, 0xa7, 0x39, 0x37, 0xee, 0x61, 0x8c, 0xee, 0x41, 0xeb, 0x52, 0xf0, 0x14, 0x27, 0x32,
	0xd7, 0x60, 0xdb, 0xa3, 0x4d, 0x6d, 0x7a, 0x24, 0xf3, 0x61, 0x8c, 0x6e, 0x43, 0x53, 0xf1, 0x85,
	0xb7, 0x62, 0x75, 0xaa, 0xb8, 0xf5, 0x3d, 0x80, 0x5b, 0xa6, 0x07, 0xf9, 0x25, 0xa6, 0x42, 0x70,
	0x81, 0x63, 0xea, 0xd2, 0xbb, 0x96, 0xed, 0x69, 0xf7, 0x93, 0xcb, 0x53, 0xed, 0x3c, 0x59, 0xf8,
	0xd0, 0xa7, 0xd0, 0xdf, 0x11, 0x56, 0x7c, 0x4f, 0xd8, 0x8e, 0xbe, 0xbb, 0x2d, 0x7e, 0xf9, 0x49,
	0xf1, 0x10, 0xee, 0xae, 0x54, 0x2e, 0x9e, 0x08, 0xb2, 0x9e, 0xc3, 0x36, 0xfd, 0xed, 0x02, 0x73,
	0xe2, 0x20, 0xcb, 0x0c, 0x43, 0x78, 0xe7, 0x05, 0x77, 0x21, 0x28, 0x91, 0x3c, 0x33, 0xa2, 0x68,
	0x86, 0xf7, 0x76, 0x5d, 0x43, 0x68, 0x50, 0x17, 0x35, 0xf3, 0x59, 0x78, 0xff, 0xff, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x22, 0xe3, 0xc3, 0x0d, 0x30, 0x0e, 0x00, 0x00,
}
