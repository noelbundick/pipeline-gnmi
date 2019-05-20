// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pce_neighbor_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_pce_peers_peer

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

// PCE Neighbor Information
type PceNeighborBag_KEYS struct {
	PeerAddress          string   `protobuf:"bytes,1,opt,name=peer_address,json=peerAddress" json:"peer_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceNeighborBag_KEYS) Reset()         { *m = PceNeighborBag_KEYS{} }
func (m *PceNeighborBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PceNeighborBag_KEYS) ProtoMessage()    {}
func (*PceNeighborBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_pce_neighbor_bag_e0cc86011fb806be, []int{0}
}
func (m *PceNeighborBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceNeighborBag_KEYS.Unmarshal(m, b)
}
func (m *PceNeighborBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceNeighborBag_KEYS.Marshal(b, m, deterministic)
}
func (dst *PceNeighborBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceNeighborBag_KEYS.Merge(dst, src)
}
func (m *PceNeighborBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PceNeighborBag_KEYS.Size(m)
}
func (m *PceNeighborBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PceNeighborBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PceNeighborBag_KEYS proto.InternalMessageInfo

func (m *PceNeighborBag_KEYS) GetPeerAddress() string {
	if m != nil {
		return m.PeerAddress
	}
	return ""
}

type PceNeighborBag struct {
	// Error (for display only)
	Error string `protobuf:"bytes,50,opt,name=error" json:"error,omitempty"`
	// PCE State
	PceState string `protobuf:"bytes,51,opt,name=pce_state,json=pceState" json:"pce_state,omitempty"`
	// PCE Precedence
	Precedence uint32 `protobuf:"varint,52,opt,name=precedence" json:"precedence,omitempty"`
	// Stateful
	Stateful bool `protobuf:"varint,53,opt,name=stateful" json:"stateful,omitempty"`
	// Update capability
	CapabilityUpdate bool `protobuf:"varint,54,opt,name=capability_update,json=capabilityUpdate" json:"capability_update,omitempty"`
	// Instantiation capability
	CapabilityInstantiate bool `protobuf:"varint,55,opt,name=capability_instantiate,json=capabilityInstantiate" json:"capability_instantiate,omitempty"`
	// Segment Routing capability
	CapabilitySegmentRouting bool `protobuf:"varint,56,opt,name=capability_segment_routing,json=capabilitySegmentRouting" json:"capability_segment_routing,omitempty"`
	// Triggered Synchronization capability
	TriggeredSyncCapability bool `protobuf:"varint,57,opt,name=triggered_sync_capability,json=triggeredSyncCapability" json:"triggered_sync_capability,omitempty"`
	// DB version capability
	CapabilityDbVersion bool `protobuf:"varint,58,opt,name=capability_db_version,json=capabilityDbVersion" json:"capability_db_version,omitempty"`
	// Delta Synchronization capability
	DeltaSyncCapability bool `protobuf:"varint,59,opt,name=delta_sync_capability,json=deltaSyncCapability" json:"delta_sync_capability,omitempty"`
	// PCEP Up Time
	PcepUpTime uint32 `protobuf:"varint,60,opt,name=pcep_up_time,json=pcepUpTime" json:"pcep_up_time,omitempty"`
	// Keepalive count
	Keepalives uint32 `protobuf:"varint,61,opt,name=keepalives" json:"keepalives,omitempty"`
	// Candidate
	Candidate bool `protobuf:"varint,62,opt,name=candidate" json:"candidate,omitempty"`
	// Statically Configured
	StaticallyConfigured bool `protobuf:"varint,63,opt,name=statically_configured,json=staticallyConfigured" json:"statically_configured,omitempty"`
	// MD5 Authentication Enabled
	Md5Enabled bool `protobuf:"varint,64,opt,name=md5_enabled,json=md5Enabled" json:"md5_enabled,omitempty"`
	// Keychain based Authentication Enabled
	KeychainEnabled bool `protobuf:"varint,65,opt,name=keychain_enabled,json=keychainEnabled" json:"keychain_enabled,omitempty"`
	// IGPs
	IgPs []*PcePeerIgpBag `protobuf:"bytes,66,rep,name=ig_ps,json=igPs" json:"ig_ps,omitempty"`
	// Negotiated KA
	NegotiatedKeepalive uint32 `protobuf:"varint,67,opt,name=negotiated_keepalive,json=negotiatedKeepalive" json:"negotiated_keepalive,omitempty"`
	// Negotiated DT
	NegotatedDeadTime uint32 `protobuf:"varint,68,opt,name=negotated_dead_time,json=negotatedDeadTime" json:"negotated_dead_time,omitempty"`
	// PCEReq Rx
	PceReqRx uint32 `protobuf:"varint,69,opt,name=pce_req_rx,json=pceReqRx" json:"pce_req_rx,omitempty"`
	// PCEReq Tx
	PceReqTx uint32 `protobuf:"varint,70,opt,name=pce_req_tx,json=pceReqTx" json:"pce_req_tx,omitempty"`
	// PCERep Rx
	PceRepRx uint32 `protobuf:"varint,71,opt,name=pce_rep_rx,json=pceRepRx" json:"pce_rep_rx,omitempty"`
	// PCERep Tx
	PceRepTx uint32 `protobuf:"varint,72,opt,name=pce_rep_tx,json=pceRepTx" json:"pce_rep_tx,omitempty"`
	// PCEErr Rx
	PceErrRx uint32 `protobuf:"varint,73,opt,name=pce_err_rx,json=pceErrRx" json:"pce_err_rx,omitempty"`
	// PCEErr Tx
	PceErrTx uint32 `protobuf:"varint,74,opt,name=pce_err_tx,json=pceErrTx" json:"pce_err_tx,omitempty"`
	// PCEOpen Tx
	PceOpenTx uint32 `protobuf:"varint,75,opt,name=pce_open_tx,json=pceOpenTx" json:"pce_open_tx,omitempty"`
	// PCEOpen Rx
	PceOpenRx uint32 `protobuf:"varint,76,opt,name=pce_open_rx,json=pceOpenRx" json:"pce_open_rx,omitempty"`
	// PCERpt Rx
	PceRptRx uint32 `protobuf:"varint,77,opt,name=pce_rpt_rx,json=pceRptRx" json:"pce_rpt_rx,omitempty"`
	// PCERpt Tx
	PceRptTx uint32 `protobuf:"varint,78,opt,name=pce_rpt_tx,json=pceRptTx" json:"pce_rpt_tx,omitempty"`
	// PCEUpd Rx
	PceUpdRx uint32 `protobuf:"varint,79,opt,name=pce_upd_rx,json=pceUpdRx" json:"pce_upd_rx,omitempty"`
	// PCEUpd Tx
	PceUpdTx uint32 `protobuf:"varint,80,opt,name=pce_upd_tx,json=pceUpdTx" json:"pce_upd_tx,omitempty"`
	// PCEInit Rx
	PceInitRx uint32 `protobuf:"varint,81,opt,name=pce_init_rx,json=pceInitRx" json:"pce_init_rx,omitempty"`
	// PCEInit_Tx
	PceInitTx uint32 `protobuf:"varint,82,opt,name=pce_init_tx,json=pceInitTx" json:"pce_init_tx,omitempty"`
	// PCE Keepalive Tx
	PceKeepaliveTx uint64 `protobuf:"varint,83,opt,name=pce_keepalive_tx,json=pceKeepaliveTx" json:"pce_keepalive_tx,omitempty"`
	// PCE Keepalive Rx
	PceKeepaliveRx uint64 `protobuf:"varint,84,opt,name=pce_keepalive_rx,json=pceKeepaliveRx" json:"pce_keepalive_rx,omitempty"`
	// Mininum reply time from peer in ms
	ReqReplyMin uint64 `protobuf:"varint,85,opt,name=req_reply_min,json=reqReplyMin" json:"req_reply_min,omitempty"`
	// Maximum reply time from peer in ms
	ReqReplyMax uint64 `protobuf:"varint,86,opt,name=req_reply_max,json=reqReplyMax" json:"req_reply_max,omitempty"`
	// Average reply time from peer in ms
	ReqReplyAvg uint64 `protobuf:"varint,87,opt,name=req_reply_avg,json=reqReplyAvg" json:"req_reply_avg,omitempty"`
	// Request timeout count
	RequestTimeouts uint64 `protobuf:"varint,88,opt,name=request_timeouts,json=requestTimeouts" json:"request_timeouts,omitempty"`
	// Last sent PCEErr
	LastTxPceErr *PceerrBag `protobuf:"bytes,89,opt,name=last_tx_pce_err,json=lastTxPceErr" json:"last_tx_pce_err,omitempty"`
	// Last received PCEErr
	LastRxPceErr *PceerrBag `protobuf:"bytes,90,opt,name=last_rx_pce_err,json=lastRxPceErr" json:"last_rx_pce_err,omitempty"`
	// Local PCEP session ID
	LocalSid uint32 `protobuf:"varint,91,opt,name=local_sid,json=localSid" json:"local_sid,omitempty"`
	// Remote PCEP session ID
	RemoteSid uint32 `protobuf:"varint,92,opt,name=remote_sid,json=remoteSid" json:"remote_sid,omitempty"`
	// Minimum keepalive interval for the peer
	MinKeepaliveInterval uint32 `protobuf:"varint,93,opt,name=min_keepalive_interval,json=minKeepaliveInterval" json:"min_keepalive_interval,omitempty"`
	// Maximum dead interval for the peer
	MaxDeadInterval      uint32   `protobuf:"varint,94,opt,name=max_dead_interval,json=maxDeadInterval" json:"max_dead_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceNeighborBag) Reset()         { *m = PceNeighborBag{} }
func (m *PceNeighborBag) String() string { return proto.CompactTextString(m) }
func (*PceNeighborBag) ProtoMessage()    {}
func (*PceNeighborBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_pce_neighbor_bag_e0cc86011fb806be, []int{1}
}
func (m *PceNeighborBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceNeighborBag.Unmarshal(m, b)
}
func (m *PceNeighborBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceNeighborBag.Marshal(b, m, deterministic)
}
func (dst *PceNeighborBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceNeighborBag.Merge(dst, src)
}
func (m *PceNeighborBag) XXX_Size() int {
	return xxx_messageInfo_PceNeighborBag.Size(m)
}
func (m *PceNeighborBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceNeighborBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceNeighborBag proto.InternalMessageInfo

func (m *PceNeighborBag) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *PceNeighborBag) GetPceState() string {
	if m != nil {
		return m.PceState
	}
	return ""
}

func (m *PceNeighborBag) GetPrecedence() uint32 {
	if m != nil {
		return m.Precedence
	}
	return 0
}

func (m *PceNeighborBag) GetStateful() bool {
	if m != nil {
		return m.Stateful
	}
	return false
}

func (m *PceNeighborBag) GetCapabilityUpdate() bool {
	if m != nil {
		return m.CapabilityUpdate
	}
	return false
}

func (m *PceNeighborBag) GetCapabilityInstantiate() bool {
	if m != nil {
		return m.CapabilityInstantiate
	}
	return false
}

func (m *PceNeighborBag) GetCapabilitySegmentRouting() bool {
	if m != nil {
		return m.CapabilitySegmentRouting
	}
	return false
}

func (m *PceNeighborBag) GetTriggeredSyncCapability() bool {
	if m != nil {
		return m.TriggeredSyncCapability
	}
	return false
}

func (m *PceNeighborBag) GetCapabilityDbVersion() bool {
	if m != nil {
		return m.CapabilityDbVersion
	}
	return false
}

func (m *PceNeighborBag) GetDeltaSyncCapability() bool {
	if m != nil {
		return m.DeltaSyncCapability
	}
	return false
}

func (m *PceNeighborBag) GetPcepUpTime() uint32 {
	if m != nil {
		return m.PcepUpTime
	}
	return 0
}

func (m *PceNeighborBag) GetKeepalives() uint32 {
	if m != nil {
		return m.Keepalives
	}
	return 0
}

func (m *PceNeighborBag) GetCandidate() bool {
	if m != nil {
		return m.Candidate
	}
	return false
}

func (m *PceNeighborBag) GetStaticallyConfigured() bool {
	if m != nil {
		return m.StaticallyConfigured
	}
	return false
}

func (m *PceNeighborBag) GetMd5Enabled() bool {
	if m != nil {
		return m.Md5Enabled
	}
	return false
}

func (m *PceNeighborBag) GetKeychainEnabled() bool {
	if m != nil {
		return m.KeychainEnabled
	}
	return false
}

func (m *PceNeighborBag) GetIgPs() []*PcePeerIgpBag {
	if m != nil {
		return m.IgPs
	}
	return nil
}

func (m *PceNeighborBag) GetNegotiatedKeepalive() uint32 {
	if m != nil {
		return m.NegotiatedKeepalive
	}
	return 0
}

func (m *PceNeighborBag) GetNegotatedDeadTime() uint32 {
	if m != nil {
		return m.NegotatedDeadTime
	}
	return 0
}

func (m *PceNeighborBag) GetPceReqRx() uint32 {
	if m != nil {
		return m.PceReqRx
	}
	return 0
}

func (m *PceNeighborBag) GetPceReqTx() uint32 {
	if m != nil {
		return m.PceReqTx
	}
	return 0
}

func (m *PceNeighborBag) GetPceRepRx() uint32 {
	if m != nil {
		return m.PceRepRx
	}
	return 0
}

func (m *PceNeighborBag) GetPceRepTx() uint32 {
	if m != nil {
		return m.PceRepTx
	}
	return 0
}

func (m *PceNeighborBag) GetPceErrRx() uint32 {
	if m != nil {
		return m.PceErrRx
	}
	return 0
}

func (m *PceNeighborBag) GetPceErrTx() uint32 {
	if m != nil {
		return m.PceErrTx
	}
	return 0
}

func (m *PceNeighborBag) GetPceOpenTx() uint32 {
	if m != nil {
		return m.PceOpenTx
	}
	return 0
}

func (m *PceNeighborBag) GetPceOpenRx() uint32 {
	if m != nil {
		return m.PceOpenRx
	}
	return 0
}

func (m *PceNeighborBag) GetPceRptRx() uint32 {
	if m != nil {
		return m.PceRptRx
	}
	return 0
}

func (m *PceNeighborBag) GetPceRptTx() uint32 {
	if m != nil {
		return m.PceRptTx
	}
	return 0
}

func (m *PceNeighborBag) GetPceUpdRx() uint32 {
	if m != nil {
		return m.PceUpdRx
	}
	return 0
}

func (m *PceNeighborBag) GetPceUpdTx() uint32 {
	if m != nil {
		return m.PceUpdTx
	}
	return 0
}

func (m *PceNeighborBag) GetPceInitRx() uint32 {
	if m != nil {
		return m.PceInitRx
	}
	return 0
}

func (m *PceNeighborBag) GetPceInitTx() uint32 {
	if m != nil {
		return m.PceInitTx
	}
	return 0
}

func (m *PceNeighborBag) GetPceKeepaliveTx() uint64 {
	if m != nil {
		return m.PceKeepaliveTx
	}
	return 0
}

func (m *PceNeighborBag) GetPceKeepaliveRx() uint64 {
	if m != nil {
		return m.PceKeepaliveRx
	}
	return 0
}

func (m *PceNeighborBag) GetReqReplyMin() uint64 {
	if m != nil {
		return m.ReqReplyMin
	}
	return 0
}

func (m *PceNeighborBag) GetReqReplyMax() uint64 {
	if m != nil {
		return m.ReqReplyMax
	}
	return 0
}

func (m *PceNeighborBag) GetReqReplyAvg() uint64 {
	if m != nil {
		return m.ReqReplyAvg
	}
	return 0
}

func (m *PceNeighborBag) GetRequestTimeouts() uint64 {
	if m != nil {
		return m.RequestTimeouts
	}
	return 0
}

func (m *PceNeighborBag) GetLastTxPceErr() *PceerrBag {
	if m != nil {
		return m.LastTxPceErr
	}
	return nil
}

func (m *PceNeighborBag) GetLastRxPceErr() *PceerrBag {
	if m != nil {
		return m.LastRxPceErr
	}
	return nil
}

func (m *PceNeighborBag) GetLocalSid() uint32 {
	if m != nil {
		return m.LocalSid
	}
	return 0
}

func (m *PceNeighborBag) GetRemoteSid() uint32 {
	if m != nil {
		return m.RemoteSid
	}
	return 0
}

func (m *PceNeighborBag) GetMinKeepaliveInterval() uint32 {
	if m != nil {
		return m.MinKeepaliveInterval
	}
	return 0
}

func (m *PceNeighborBag) GetMaxDeadInterval() uint32 {
	if m != nil {
		return m.MaxDeadInterval
	}
	return 0
}

// PCE Error information
type PceerrBag struct {
	// PCE Error Type
	PceErrType uint32 `protobuf:"varint,1,opt,name=pce_err_type,json=pceErrType" json:"pce_err_type,omitempty"`
	// PCE Error Value
	PceErrValue uint32 `protobuf:"varint,2,opt,name=pce_err_value,json=pceErrValue" json:"pce_err_value,omitempty"`
	// Has Open object
	HasOpenObject bool `protobuf:"varint,3,opt,name=has_open_object,json=hasOpenObject" json:"has_open_object,omitempty"`
	// Open Version
	OpenVersion uint32 `protobuf:"varint,4,opt,name=open_version,json=openVersion" json:"open_version,omitempty"`
	// Open Keepalive (seconds)
	OpenKeepalive uint32 `protobuf:"varint,5,opt,name=open_keepalive,json=openKeepalive" json:"open_keepalive,omitempty"`
	// Open Deadtime (seconds)
	OpenDeadTime uint32 `protobuf:"varint,6,opt,name=open_dead_time,json=openDeadTime" json:"open_dead_time,omitempty"`
	// Has RP object
	HasRpObject bool `protobuf:"varint,7,opt,name=has_rp_object,json=hasRpObject" json:"has_rp_object,omitempty"`
	// RP request ID
	RpRequestId          uint32   `protobuf:"varint,8,opt,name=rp_request_id,json=rpRequestId" json:"rp_request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceerrBag) Reset()         { *m = PceerrBag{} }
func (m *PceerrBag) String() string { return proto.CompactTextString(m) }
func (*PceerrBag) ProtoMessage()    {}
func (*PceerrBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_pce_neighbor_bag_e0cc86011fb806be, []int{2}
}
func (m *PceerrBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceerrBag.Unmarshal(m, b)
}
func (m *PceerrBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceerrBag.Marshal(b, m, deterministic)
}
func (dst *PceerrBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceerrBag.Merge(dst, src)
}
func (m *PceerrBag) XXX_Size() int {
	return xxx_messageInfo_PceerrBag.Size(m)
}
func (m *PceerrBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceerrBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceerrBag proto.InternalMessageInfo

func (m *PceerrBag) GetPceErrType() uint32 {
	if m != nil {
		return m.PceErrType
	}
	return 0
}

func (m *PceerrBag) GetPceErrValue() uint32 {
	if m != nil {
		return m.PceErrValue
	}
	return 0
}

func (m *PceerrBag) GetHasOpenObject() bool {
	if m != nil {
		return m.HasOpenObject
	}
	return false
}

func (m *PceerrBag) GetOpenVersion() uint32 {
	if m != nil {
		return m.OpenVersion
	}
	return 0
}

func (m *PceerrBag) GetOpenKeepalive() uint32 {
	if m != nil {
		return m.OpenKeepalive
	}
	return 0
}

func (m *PceerrBag) GetOpenDeadTime() uint32 {
	if m != nil {
		return m.OpenDeadTime
	}
	return 0
}

func (m *PceerrBag) GetHasRpObject() bool {
	if m != nil {
		return m.HasRpObject
	}
	return false
}

func (m *PceerrBag) GetRpRequestId() uint32 {
	if m != nil {
		return m.RpRequestId
	}
	return 0
}

// PCE IGP information
type PcePeerIgpBag struct {
	// Type
	IgpType string `protobuf:"bytes,1,opt,name=igp_type,json=igpType" json:"igp_type,omitempty"`
	// Instance ID
	IgpInstanceId        string   `protobuf:"bytes,2,opt,name=igp_instance_id,json=igpInstanceId" json:"igp_instance_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcePeerIgpBag) Reset()         { *m = PcePeerIgpBag{} }
func (m *PcePeerIgpBag) String() string { return proto.CompactTextString(m) }
func (*PcePeerIgpBag) ProtoMessage()    {}
func (*PcePeerIgpBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_pce_neighbor_bag_e0cc86011fb806be, []int{3}
}
func (m *PcePeerIgpBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePeerIgpBag.Unmarshal(m, b)
}
func (m *PcePeerIgpBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePeerIgpBag.Marshal(b, m, deterministic)
}
func (dst *PcePeerIgpBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePeerIgpBag.Merge(dst, src)
}
func (m *PcePeerIgpBag) XXX_Size() int {
	return xxx_messageInfo_PcePeerIgpBag.Size(m)
}
func (m *PcePeerIgpBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePeerIgpBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcePeerIgpBag proto.InternalMessageInfo

func (m *PcePeerIgpBag) GetIgpType() string {
	if m != nil {
		return m.IgpType
	}
	return ""
}

func (m *PcePeerIgpBag) GetIgpInstanceId() string {
	if m != nil {
		return m.IgpInstanceId
	}
	return ""
}

func init() {
	proto.RegisterType((*PceNeighborBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce.peers.peer.pce_neighbor_bag_KEYS")
	proto.RegisterType((*PceNeighborBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce.peers.peer.pce_neighbor_bag")
	proto.RegisterType((*PceerrBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce.peers.peer.pceerr_bag")
	proto.RegisterType((*PcePeerIgpBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce.peers.peer.pce_peer_igp_bag")
}

func init() {
	proto.RegisterFile("pce_neighbor_bag.proto", fileDescriptor_pce_neighbor_bag_e0cc86011fb806be)
}

var fileDescriptor_pce_neighbor_bag_e0cc86011fb806be = []byte{
	// 1071 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xfb, 0x72, 0x1b, 0x35,
	0x14, 0xc6, 0xc7, 0xbd, 0xc6, 0x72, 0x53, 0xb7, 0x6a, 0x5a, 0xd4, 0x0b, 0xc5, 0x78, 0xa0, 0x63,
	0x60, 0xf0, 0x0c, 0x69, 0x03, 0x34, 0x14, 0x4a, 0x48, 0x02, 0x98, 0x50, 0x12, 0xd6, 0x4e, 0xa0,
	0xdc, 0x84, 0xbc, 0x3a, 0xdd, 0x88, 0xee, 0x45, 0xd1, 0xca, 0x9e, 0xf5, 0x6b, 0xf1, 0x38, 0x3c,
	0x0d, 0xa3, 0x23, 0x7b, 0x77, 0xed, 0xf2, 0x4f, 0x67, 0xf8, 0x27, 0x93, 0xfd, 0x7e, 0xdf, 0xb7,
	0x3a, 0x23, 0x9d, 0xa3, 0x35, 0xb9, 0xa5, 0x43, 0xe0, 0x29, 0xa8, 0xe8, 0x74, 0x9c, 0x19, 0x3e,
	0x16, 0x51, 0x5f, 0x9b, 0xcc, 0x66, 0xf4, 0xc3, 0x50, 0xe5, 0x61, 0xc6, 0x55, 0x96, 0xf3, 0xc2,
	0xf0, 0x44, 0xc7, 0x39, 0xb7, 0xc0, 0x33, 0x0d, 0xa6, 0x8f, 0x0f, 0x3a, 0x84, 0xbe, 0x06, 0x30,
	0x39, 0xfe, 0xed, 0x6e, 0x93, 0x9b, 0xab, 0x2f, 0xe2, 0x07, 0xfb, 0xcf, 0x87, 0xf4, 0x6d, 0x72,
	0xc5, 0x19, 0xb8, 0x90, 0xd2, 0x40, 0x9e, 0xb3, 0x46, 0xa7, 0xd1, 0x6b, 0x06, 0x2d, 0xa7, 0xed,
	0x78, 0xa9, 0xfb, 0xcf, 0x55, 0x72, 0x6d, 0x35, 0x4c, 0x37, 0xc8, 0x45, 0x30, 0x26, 0x33, 0x6c,
	0x13, 0x03, 0xfe, 0x81, 0xde, 0x25, 0x4d, 0xe7, 0xcc, 0xad, 0xb0, 0xc0, 0x1e, 0x22, 0x59, 0xd3,
	0x21, 0x0c, 0xdd, 0x33, 0xbd, 0x4f, 0x88, 0x36, 0x10, 0x82, 0x84, 0x34, 0x04, 0xf6, 0xa8, 0xd3,
	0xe8, 0xad, 0x07, 0x35, 0x85, 0xde, 0x21, 0x6b, 0x18, 0x7c, 0x31, 0x89, 0xd9, 0x56, 0xa7, 0xd1,
	0x5b, 0x0b, 0xca, 0x67, 0xfa, 0x01, 0xb9, 0x1e, 0x0a, 0x2d, 0xc6, 0x2a, 0x56, 0x76, 0xc6, 0x27,
	0x5a, 0xba, 0x05, 0x3e, 0x46, 0xd3, 0xb5, 0x0a, 0x1c, 0xa3, 0x4e, 0xb7, 0xc8, 0xad, 0x9a, 0x59,
	0xa5, 0xb9, 0x15, 0xa9, 0x55, 0x2e, 0xf1, 0x09, 0x26, 0x6e, 0x56, 0x74, 0x50, 0x41, 0xfa, 0x84,
	0xdc, 0xa9, 0xc5, 0x72, 0x88, 0x12, 0x48, 0x2d, 0x37, 0xd9, 0xc4, 0xaa, 0x34, 0x62, 0x9f, 0x62,
	0x94, 0x55, 0x8e, 0xa1, 0x37, 0x04, 0x9e, 0xd3, 0x6d, 0x72, 0xdb, 0x1a, 0x15, 0x45, 0x60, 0x40,
	0xf2, 0x7c, 0x96, 0x86, 0xbc, 0xb2, 0xb2, 0xc7, 0x18, 0x7e, 0xa3, 0x34, 0x0c, 0x67, 0x69, 0xb8,
	0x5b, 0x62, 0xba, 0x49, 0x6a, 0x25, 0x71, 0x39, 0xe6, 0x53, 0x30, 0xb9, 0xca, 0x52, 0xb6, 0x8d,
	0xb9, 0x1b, 0x15, 0xdc, 0x1b, 0x9f, 0x78, 0xe4, 0x32, 0x12, 0x62, 0x2b, 0x5e, 0x59, 0xeb, 0x33,
	0x9f, 0x41, 0xb8, 0xb2, 0x4e, 0x87, 0x5c, 0xd1, 0x21, 0x68, 0x3e, 0xd1, 0xdc, 0xaa, 0x04, 0xd8,
	0x93, 0xf9, 0x19, 0x84, 0xa0, 0x8f, 0xf5, 0x48, 0x25, 0x78, 0x46, 0x2f, 0x01, 0xb4, 0x88, 0xd5,
	0x14, 0x72, 0xf6, 0xb9, 0xe7, 0x95, 0x42, 0xef, 0x91, 0x66, 0x28, 0x52, 0xa9, 0x70, 0xff, 0xbf,
	0xc0, 0x95, 0x2a, 0x81, 0x3e, 0x24, 0x37, 0xdd, 0x89, 0xa9, 0x50, 0xc4, 0xf1, 0x8c, 0x87, 0x59,
	0xfa, 0x42, 0x45, 0x13, 0x03, 0x92, 0x3d, 0x45, 0xe7, 0x46, 0x05, 0x77, 0x4b, 0x46, 0xdf, 0x22,
	0xad, 0x44, 0x6e, 0x71, 0x48, 0xc5, 0x38, 0x06, 0xc9, 0xbe, 0x44, 0x2b, 0x49, 0xe4, 0xd6, 0xbe,
	0x57, 0xe8, 0x7b, 0xe4, 0xda, 0x4b, 0x98, 0x85, 0xa7, 0x42, 0xa5, 0xa5, 0x6b, 0x07, 0x5d, 0xed,
	0x85, 0xbe, 0xb0, 0x8e, 0xc8, 0x45, 0x15, 0x71, 0x9d, 0xb3, 0xaf, 0x3a, 0xe7, 0x7b, 0xad, 0xcd,
	0xa7, 0xfd, 0xd7, 0x9a, 0x92, 0xbe, 0xeb, 0x5d, 0x9c, 0x06, 0x15, 0x69, 0xd7, 0xe5, 0xc1, 0x05,
	0x15, 0x1d, 0xe5, 0xf4, 0x23, 0xb2, 0x91, 0x42, 0x94, 0x61, 0x97, 0x48, 0x5e, 0xee, 0x06, 0xdb,
	0xc5, 0xed, 0xb9, 0x51, 0xb1, 0x83, 0x05, 0xa2, 0x7d, 0xe2, 0x65, 0x4c, 0x48, 0x10, 0xd2, 0x6f,
	0xf8, 0x1e, 0x26, 0xae, 0x97, 0x68, 0x0f, 0x84, 0xc4, 0x7d, 0xbf, 0x47, 0xdc, 0x29, 0x70, 0x03,
	0x67, 0xdc, 0x14, 0x6c, 0x1f, 0x6d, 0x6e, 0x72, 0x02, 0x38, 0x0b, 0x8a, 0x3a, 0xb5, 0x05, 0xfb,
	0xba, 0x4e, 0x47, 0x35, 0xaa, 0x5d, 0xf6, 0x9b, 0x1a, 0xd5, 0xc1, 0x12, 0xb5, 0x05, 0xfb, 0xb6,
	0x4e, 0xab, 0x2c, 0x18, 0xe3, 0xb2, 0x83, 0x92, 0xee, 0x1b, 0x13, 0x2c, 0x51, 0x5b, 0xb0, 0xef,
	0xea, 0x74, 0x54, 0xd0, 0xfb, 0xa4, 0xe5, 0x68, 0xa6, 0x21, 0x75, 0xf8, 0x00, 0xb1, 0x9b, 0xff,
	0x43, 0x0d, 0xe9, 0x0a, 0x37, 0x05, 0xfb, 0x7e, 0x89, 0xd7, 0x2a, 0xd3, 0xd6, 0xe1, 0x67, 0x55,
	0x65, 0xda, 0x2e, 0x53, 0x5b, 0xb0, 0x1f, 0xea, 0xb4, 0xaa, 0x7b, 0xa2, 0xa5, 0xcb, 0x1e, 0x96,
	0xf4, 0x58, 0xcb, 0x60, 0x89, 0xda, 0x82, 0x1d, 0xd5, 0x69, 0x55, 0x97, 0x4a, 0x15, 0x2e, 0xfc,
	0x63, 0x59, 0xd7, 0x20, 0x55, 0x6e, 0xe5, 0x3a, 0xb7, 0x05, 0x0b, 0x96, 0xf8, 0xa8, 0xa0, 0x3d,
	0x7f, 0x1d, 0x96, 0x7d, 0xe0, 0x4c, 0xc3, 0x4e, 0xa3, 0x77, 0x21, 0xb8, 0xaa, 0x43, 0x28, 0x7b,
	0xe0, 0xbf, 0x9c, 0xa6, 0x60, 0xa3, 0x57, 0x9d, 0x41, 0x41, 0xbb, 0x64, 0x1d, 0xcf, 0x1e, 0x74,
	0x3c, 0xe3, 0x89, 0x4a, 0xd9, 0x31, 0xda, 0x5a, 0x06, 0xce, 0x02, 0xa7, 0x3d, 0x53, 0xe9, 0x8a,
	0x47, 0x14, 0xec, 0x64, 0xc5, 0x23, 0x56, 0xde, 0x23, 0xa6, 0x11, 0xfb, 0x69, 0xd9, 0xb3, 0x33,
	0x8d, 0xdc, 0x3c, 0x19, 0x38, 0x9b, 0x40, 0x6e, 0xb1, 0x29, 0xb3, 0x89, 0xcd, 0xd9, 0xcf, 0x68,
	0x6b, 0xcf, 0xf5, 0xd1, 0x5c, 0xa6, 0x7f, 0x92, 0x76, 0x2c, 0x9c, 0xaf, 0xe0, 0xf3, 0x46, 0x60,
	0xcf, 0x3b, 0x8d, 0x5e, 0x6b, 0xf3, 0xf1, 0xeb, 0x4f, 0x96, 0xeb, 0x22, 0x37, 0x53, 0x57, 0xdc,
	0x1b, 0x47, 0xc5, 0x11, 0xb6, 0x51, 0xb9, 0x82, 0xa9, 0x56, 0xf8, 0xe5, 0x7f, 0x59, 0x21, 0x58,
	0xac, 0x70, 0x97, 0x34, 0xe3, 0x2c, 0x14, 0x31, 0xcf, 0x95, 0x64, 0xbf, 0xfa, 0x5e, 0x40, 0x61,
	0xa8, 0x24, 0x7d, 0x93, 0x10, 0x03, 0x49, 0x66, 0x01, 0xe9, 0x6f, 0xfe, 0xa8, 0xbd, 0xe2, 0xf0,
	0x23, 0x72, 0x2b, 0x51, 0x69, 0xed, 0x00, 0x55, 0x6a, 0xc1, 0x4c, 0x45, 0xcc, 0x7e, 0x47, 0xeb,
	0x46, 0xa2, 0xd2, 0xf2, 0x18, 0x07, 0x73, 0x46, 0xdf, 0x27, 0xd7, 0x13, 0x51, 0xf8, 0xb1, 0x2f,
	0x03, 0x7f, 0x60, 0xa0, 0x9d, 0x88, 0xc2, 0x0d, 0xfd, 0xc2, 0xdb, 0xfd, 0xfb, 0x1c, 0xf6, 0xea,
	0xbc, 0xf4, 0xf9, 0x0d, 0xed, 0x27, 0x6e, 0xa6, 0x01, 0x3f, 0xc7, 0xfe, 0x86, 0x76, 0x33, 0x37,
	0xd3, 0xe0, 0x4e, 0x78, 0xe1, 0x98, 0x8a, 0x78, 0x02, 0xec, 0x1c, 0x5a, 0x5a, 0xde, 0x72, 0xe2,
	0x24, 0xfa, 0x80, 0xb4, 0x4f, 0x45, 0xee, 0x27, 0x2f, 0x1b, 0xff, 0x05, 0xa1, 0x65, 0xe7, 0xf1,
	0xc2, 0x5c, 0x3f, 0x15, 0xb9, 0x9b, 0xbe, 0x43, 0x14, 0xdd, 0xc7, 0x1f, 0x3d, 0x8b, 0xcf, 0xcd,
	0x05, 0xff, 0x2a, 0xa7, 0x2d, 0x3e, 0x33, 0xef, 0x92, 0xab, 0x68, 0xa9, 0x6e, 0xbd, 0x8b, 0x68,
	0x5a, 0x77, 0x6a, 0x75, 0xdf, 0xbd, 0x33, 0xb7, 0x55, 0x57, 0xdd, 0x25, 0xb4, 0xe1, 0xfb, 0xcb,
	0x5b, 0xae, 0x4b, 0x5c, 0x01, 0xdc, 0xe8, 0x45, 0x55, 0x97, 0xb1, 0xaa, 0xd6, 0xa9, 0xc8, 0x03,
	0x3d, 0xaf, 0xc9, 0x75, 0xb0, 0xe6, 0x8b, 0x06, 0x55, 0x92, 0xad, 0xf9, 0xa2, 0x8c, 0x0e, 0xbc,
	0x36, 0x90, 0xdd, 0x63, 0x3f, 0x57, 0xf5, 0xab, 0x9a, 0xde, 0x26, 0x6b, 0xee, 0xdf, 0x72, 0xd7,
	0x9a, 0xc1, 0x65, 0x15, 0x69, 0xdc, 0xb2, 0x07, 0xa4, 0xed, 0x90, 0xff, 0x21, 0xe0, 0x26, 0x5b,
	0xe2, 0xa6, 0x35, 0x83, 0x75, 0x15, 0xe9, 0xc1, 0x5c, 0x1d, 0xc8, 0xf1, 0x25, 0xfc, 0x69, 0xf5,
	0xf0, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0xc0, 0x7f, 0x76, 0x74, 0x09, 0x00, 0x00,
}