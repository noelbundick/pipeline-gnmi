// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_path_protection.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_tunnels_tunnel_path_protections_tunnel_path_protection

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

// MPLS TE Path Protected Information
type MplsTePathProtection_KEYS struct {
	CType                string   `protobuf:"bytes,1,opt,name=c_type,json=cType" json:"c_type,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	TunnelName           string   `protobuf:"bytes,3,opt,name=tunnel_name,json=tunnelName" json:"tunnel_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTePathProtection_KEYS) Reset()         { *m = MplsTePathProtection_KEYS{} }
func (m *MplsTePathProtection_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTePathProtection_KEYS) ProtoMessage()    {}
func (*MplsTePathProtection_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_path_protection_a2caa39ccbfa33ab, []int{0}
}
func (m *MplsTePathProtection_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTePathProtection_KEYS.Unmarshal(m, b)
}
func (m *MplsTePathProtection_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTePathProtection_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTePathProtection_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTePathProtection_KEYS.Merge(dst, src)
}
func (m *MplsTePathProtection_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTePathProtection_KEYS.Size(m)
}
func (m *MplsTePathProtection_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTePathProtection_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTePathProtection_KEYS proto.InternalMessageInfo

func (m *MplsTePathProtection_KEYS) GetCType() string {
	if m != nil {
		return m.CType
	}
	return ""
}

func (m *MplsTePathProtection_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *MplsTePathProtection_KEYS) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

type MplsTePathProtection struct {
	// Tunnel UP
	IsTunnelUp bool `protobuf:"varint,50,opt,name=is_tunnel_up,json=isTunnelUp" json:"is_tunnel_up,omitempty"`
	// Source
	SourceAddress string `protobuf:"bytes,51,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	// Destination
	DestinationAddress string `protobuf:"bytes,52,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// Extended Tunnel ID
	ExtendedTunnelId string `protobuf:"bytes,53,opt,name=extended_tunnel_id,json=extendedTunnelId" json:"extended_tunnel_id,omitempty"`
	// Path Protect Info
	PathProtection *MplsTePpInfo `protobuf:"bytes,54,opt,name=path_protection,json=pathProtection" json:"path_protection,omitempty"`
	//  Current LSP info
	CurrentLsp *MplsTePathProtectionLsp `protobuf:"bytes,55,opt,name=current_lsp,json=currentLsp" json:"current_lsp,omitempty"`
	//  Standby LSP info
	StandbyLsp           *MplsTePathProtectionLsp `protobuf:"bytes,56,opt,name=standby_lsp,json=standbyLsp" json:"standby_lsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MplsTePathProtection) Reset()         { *m = MplsTePathProtection{} }
func (m *MplsTePathProtection) String() string { return proto.CompactTextString(m) }
func (*MplsTePathProtection) ProtoMessage()    {}
func (*MplsTePathProtection) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_path_protection_a2caa39ccbfa33ab, []int{1}
}
func (m *MplsTePathProtection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTePathProtection.Unmarshal(m, b)
}
func (m *MplsTePathProtection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTePathProtection.Marshal(b, m, deterministic)
}
func (dst *MplsTePathProtection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTePathProtection.Merge(dst, src)
}
func (m *MplsTePathProtection) XXX_Size() int {
	return xxx_messageInfo_MplsTePathProtection.Size(m)
}
func (m *MplsTePathProtection) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTePathProtection.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTePathProtection proto.InternalMessageInfo

func (m *MplsTePathProtection) GetIsTunnelUp() bool {
	if m != nil {
		return m.IsTunnelUp
	}
	return false
}

func (m *MplsTePathProtection) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsTePathProtection) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTePathProtection) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *MplsTePathProtection) GetPathProtection() *MplsTePpInfo {
	if m != nil {
		return m.PathProtection
	}
	return nil
}

func (m *MplsTePathProtection) GetCurrentLsp() *MplsTePathProtectionLsp {
	if m != nil {
		return m.CurrentLsp
	}
	return nil
}

func (m *MplsTePathProtection) GetStandbyLsp() *MplsTePathProtectionLsp {
	if m != nil {
		return m.StandbyLsp
	}
	return nil
}

// RSVP ERO IPV4 subobject
type RsvpMgmtEroIpv4Subobj struct {
	// ERO Entry Is Strict
	IsStrictRoute bool `protobuf:"varint,1,opt,name=is_strict_route,json=isStrictRoute" json:"is_strict_route,omitempty"`
	// The ERO IPV4 Address
	EroAddress string `protobuf:"bytes,2,opt,name=ero_address,json=eroAddress" json:"ero_address,omitempty"`
	// ERO Prefix Length
	PrefixLength         uint32   `protobuf:"varint,3,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtEroIpv4Subobj) Reset()         { *m = RsvpMgmtEroIpv4Subobj{} }
func (m *RsvpMgmtEroIpv4Subobj) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtEroIpv4Subobj) ProtoMessage()    {}
func (*RsvpMgmtEroIpv4Subobj) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_path_protection_a2caa39ccbfa33ab, []int{2}
}
func (m *RsvpMgmtEroIpv4Subobj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtEroIpv4Subobj.Unmarshal(m, b)
}
func (m *RsvpMgmtEroIpv4Subobj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtEroIpv4Subobj.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtEroIpv4Subobj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtEroIpv4Subobj.Merge(dst, src)
}
func (m *RsvpMgmtEroIpv4Subobj) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtEroIpv4Subobj.Size(m)
}
func (m *RsvpMgmtEroIpv4Subobj) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtEroIpv4Subobj.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtEroIpv4Subobj proto.InternalMessageInfo

func (m *RsvpMgmtEroIpv4Subobj) GetIsStrictRoute() bool {
	if m != nil {
		return m.IsStrictRoute
	}
	return false
}

func (m *RsvpMgmtEroIpv4Subobj) GetEroAddress() string {
	if m != nil {
		return m.EroAddress
	}
	return ""
}

func (m *RsvpMgmtEroIpv4Subobj) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

// RSVP ERO UNNUM subobject
type RsvpMgmtEroUnnumSubobj struct {
	// ERO Entry Is Strict
	IsStrictRoute bool `protobuf:"varint,1,opt,name=is_strict_route,json=isStrictRoute" json:"is_strict_route,omitempty"`
	// The Interface ID in ERO
	EroInterfaceId uint32 `protobuf:"varint,2,opt,name=ero_interface_id,json=eroInterfaceId" json:"ero_interface_id,omitempty"`
	// The Router ID in ERO
	EroRouterId string `protobuf:"bytes,3,opt,name=ero_router_id,json=eroRouterId" json:"ero_router_id,omitempty"`
	// Status of ERO
	Status               string   `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtEroUnnumSubobj) Reset()         { *m = RsvpMgmtEroUnnumSubobj{} }
func (m *RsvpMgmtEroUnnumSubobj) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtEroUnnumSubobj) ProtoMessage()    {}
func (*RsvpMgmtEroUnnumSubobj) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_path_protection_a2caa39ccbfa33ab, []int{3}
}
func (m *RsvpMgmtEroUnnumSubobj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtEroUnnumSubobj.Unmarshal(m, b)
}
func (m *RsvpMgmtEroUnnumSubobj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtEroUnnumSubobj.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtEroUnnumSubobj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtEroUnnumSubobj.Merge(dst, src)
}
func (m *RsvpMgmtEroUnnumSubobj) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtEroUnnumSubobj.Size(m)
}
func (m *RsvpMgmtEroUnnumSubobj) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtEroUnnumSubobj.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtEroUnnumSubobj proto.InternalMessageInfo

func (m *RsvpMgmtEroUnnumSubobj) GetIsStrictRoute() bool {
	if m != nil {
		return m.IsStrictRoute
	}
	return false
}

func (m *RsvpMgmtEroUnnumSubobj) GetEroInterfaceId() uint32 {
	if m != nil {
		return m.EroInterfaceId
	}
	return 0
}

func (m *RsvpMgmtEroUnnumSubobj) GetEroRouterId() string {
	if m != nil {
		return m.EroRouterId
	}
	return ""
}

func (m *RsvpMgmtEroUnnumSubobj) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

// Union of the different RSVP ERO types
type RsvpMgmtEroSubobj struct {
	EroType string `protobuf:"bytes,1,opt,name=ero_type,json=eroType" json:"ero_type,omitempty"`
	// IPV4 ERO Sub Object
	Ipv4EroSubObject *RsvpMgmtEroIpv4Subobj `protobuf:"bytes,2,opt,name=ipv4_ero_sub_object,json=ipv4EroSubObject" json:"ipv4_ero_sub_object,omitempty"`
	// Unnumbered ERO Sub Object
	UnnumberedEroSubObject *RsvpMgmtEroUnnumSubobj `protobuf:"bytes,3,opt,name=unnumbered_ero_sub_object,json=unnumberedEroSubObject" json:"unnumbered_ero_sub_object,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *RsvpMgmtEroSubobj) Reset()         { *m = RsvpMgmtEroSubobj{} }
func (m *RsvpMgmtEroSubobj) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtEroSubobj) ProtoMessage()    {}
func (*RsvpMgmtEroSubobj) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_path_protection_a2caa39ccbfa33ab, []int{4}
}
func (m *RsvpMgmtEroSubobj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtEroSubobj.Unmarshal(m, b)
}
func (m *RsvpMgmtEroSubobj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtEroSubobj.Marshal(b, m, deterministic)
}
func (dst *RsvpMgmtEroSubobj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtEroSubobj.Merge(dst, src)
}
func (m *RsvpMgmtEroSubobj) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtEroSubobj.Size(m)
}
func (m *RsvpMgmtEroSubobj) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtEroSubobj.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtEroSubobj proto.InternalMessageInfo

func (m *RsvpMgmtEroSubobj) GetEroType() string {
	if m != nil {
		return m.EroType
	}
	return ""
}

func (m *RsvpMgmtEroSubobj) GetIpv4EroSubObject() *RsvpMgmtEroIpv4Subobj {
	if m != nil {
		return m.Ipv4EroSubObject
	}
	return nil
}

func (m *RsvpMgmtEroSubobj) GetUnnumberedEroSubObject() *RsvpMgmtEroUnnumSubobj {
	if m != nil {
		return m.UnnumberedEroSubObject
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
	return fileDescriptor_mpls_te_path_protection_a2caa39ccbfa33ab, []int{5}
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

// MPLS TE Path Protected Switchover Information
type MplsTePpInfo struct {
	// The date when the switchover was completed. This date is the number of seconds since Jan 1st 1970
	TimeOfSwitchoverSec uint32 `protobuf:"varint,1,opt,name=time_of_switchover_sec,json=timeOfSwitchoverSec" json:"time_of_switchover_sec,omitempty"`
	// Number of times switchover occurred
	SwitchoverTotal uint32 `protobuf:"varint,2,opt,name=switchover_total,json=switchoverTotal" json:"switchover_total,omitempty"`
	// Total times, Standby LSP ready for use
	SwitchoverReady uint32 `protobuf:"varint,3,opt,name=switchover_ready,json=switchoverReady" json:"switchover_ready,omitempty"`
	// Number of times the standby LSP was reoptimized
	StandbyReoptimizedNumber uint32 `protobuf:"varint,4,opt,name=standby_reoptimized_number,json=standbyReoptimizedNumber" json:"standby_reoptimized_number,omitempty"`
	// Reason for last switchover
	SwitchoverReason uint32 `protobuf:"varint,5,opt,name=switchover_reason,json=switchoverReason" json:"switchover_reason,omitempty"`
	// Standby path diversity type
	DiversityType string `protobuf:"bytes,6,opt,name=diversity_type,json=diversityType" json:"diversity_type,omitempty"`
	// Is Path Protection Configured
	IsPathProtectConfigured bool `protobuf:"varint,7,opt,name=is_path_protect_configured,json=isPathProtectConfigured" json:"is_path_protect_configured,omitempty"`
	// The path option ID (level/index) of the configured explicit protecting path
	PathProtectionProtectedById uint32 `protobuf:"varint,8,opt,name=path_protection_protected_by_id,json=pathProtectionProtectedById" json:"path_protection_protected_by_id,omitempty"`
	// Flag to indicate whether tunnel has a path-option which is valid for path-protection
	ValidPathProtectionPathOptionExists bool `protobuf:"varint,9,opt,name=valid_path_protection_path_option_exists,json=validPathProtectionPathOptionExists" json:"valid_path_protection_path_option_exists,omitempty"`
	// Is Switchover Underway
	IsPathProtectSwitchOverUnderway bool `protobuf:"varint,10,opt,name=is_path_protect_switch_over_underway,json=isPathProtectSwitchOverUnderway" json:"is_path_protect_switch_over_underway,omitempty"`
	// Path Protected Switchover Information
	Switchover *TePpSwLogEntryBag `protobuf:"bytes,11,opt,name=switchover" json:"switchover,omitempty"`
	// Remaining time until path protection reoptimization (seconds)
	ReoptimizationTimeRemaining uint32   `protobuf:"varint,12,opt,name=reoptimization_time_remaining,json=reoptimizationTimeRemaining" json:"reoptimization_time_remaining,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *MplsTePpInfo) Reset()         { *m = MplsTePpInfo{} }
func (m *MplsTePpInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTePpInfo) ProtoMessage()    {}
func (*MplsTePpInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_path_protection_a2caa39ccbfa33ab, []int{6}
}
func (m *MplsTePpInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTePpInfo.Unmarshal(m, b)
}
func (m *MplsTePpInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTePpInfo.Marshal(b, m, deterministic)
}
func (dst *MplsTePpInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTePpInfo.Merge(dst, src)
}
func (m *MplsTePpInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTePpInfo.Size(m)
}
func (m *MplsTePpInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTePpInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTePpInfo proto.InternalMessageInfo

func (m *MplsTePpInfo) GetTimeOfSwitchoverSec() uint32 {
	if m != nil {
		return m.TimeOfSwitchoverSec
	}
	return 0
}

func (m *MplsTePpInfo) GetSwitchoverTotal() uint32 {
	if m != nil {
		return m.SwitchoverTotal
	}
	return 0
}

func (m *MplsTePpInfo) GetSwitchoverReady() uint32 {
	if m != nil {
		return m.SwitchoverReady
	}
	return 0
}

func (m *MplsTePpInfo) GetStandbyReoptimizedNumber() uint32 {
	if m != nil {
		return m.StandbyReoptimizedNumber
	}
	return 0
}

func (m *MplsTePpInfo) GetSwitchoverReason() uint32 {
	if m != nil {
		return m.SwitchoverReason
	}
	return 0
}

func (m *MplsTePpInfo) GetDiversityType() string {
	if m != nil {
		return m.DiversityType
	}
	return ""
}

func (m *MplsTePpInfo) GetIsPathProtectConfigured() bool {
	if m != nil {
		return m.IsPathProtectConfigured
	}
	return false
}

func (m *MplsTePpInfo) GetPathProtectionProtectedById() uint32 {
	if m != nil {
		return m.PathProtectionProtectedById
	}
	return 0
}

func (m *MplsTePpInfo) GetValidPathProtectionPathOptionExists() bool {
	if m != nil {
		return m.ValidPathProtectionPathOptionExists
	}
	return false
}

func (m *MplsTePpInfo) GetIsPathProtectSwitchOverUnderway() bool {
	if m != nil {
		return m.IsPathProtectSwitchOverUnderway
	}
	return false
}

func (m *MplsTePpInfo) GetSwitchover() *TePpSwLogEntryBag {
	if m != nil {
		return m.Switchover
	}
	return nil
}

func (m *MplsTePpInfo) GetReoptimizationTimeRemaining() uint32 {
	if m != nil {
		return m.ReoptimizationTimeRemaining
	}
	return 0
}

// Segment-routing outgoing fowarding info
type MplsTeS2LSrOutgoingFwdInfo struct {
	// Output interface of LSP
	LspOutputInterface string `protobuf:"bytes,1,opt,name=lsp_output_interface,json=lspOutputInterface" json:"lsp_output_interface,omitempty"`
	// Output label of the LSP
	LspOutputLabel       uint32   `protobuf:"varint,2,opt,name=lsp_output_label,json=lspOutputLabel" json:"lsp_output_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeS2LSrOutgoingFwdInfo) Reset()         { *m = MplsTeS2LSrOutgoingFwdInfo{} }
func (m *MplsTeS2LSrOutgoingFwdInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeS2LSrOutgoingFwdInfo) ProtoMessage()    {}
func (*MplsTeS2LSrOutgoingFwdInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_path_protection_a2caa39ccbfa33ab, []int{7}
}
func (m *MplsTeS2LSrOutgoingFwdInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo.Unmarshal(m, b)
}
func (m *MplsTeS2LSrOutgoingFwdInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo.Marshal(b, m, deterministic)
}
func (dst *MplsTeS2LSrOutgoingFwdInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo.Merge(dst, src)
}
func (m *MplsTeS2LSrOutgoingFwdInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo.Size(m)
}
func (m *MplsTeS2LSrOutgoingFwdInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo proto.InternalMessageInfo

func (m *MplsTeS2LSrOutgoingFwdInfo) GetLspOutputInterface() string {
	if m != nil {
		return m.LspOutputInterface
	}
	return ""
}

func (m *MplsTeS2LSrOutgoingFwdInfo) GetLspOutputLabel() uint32 {
	if m != nil {
		return m.LspOutputLabel
	}
	return 0
}

// MPLS TE Path Protected LSP Information
type MplsTePathProtectionLsp struct {
	// LSP Uptime
	LspUptime uint32 `protobuf:"varint,1,opt,name=lsp_uptime,json=lspUptime" json:"lsp_uptime,omitempty"`
	// LSP ID
	PathProtectionLspId uint32 `protobuf:"varint,2,opt,name=path_protection_lsp_id,json=pathProtectionLspId" json:"path_protection_lsp_id,omitempty"`
	// Local label
	LspLocalLabel uint32 `protobuf:"varint,3,opt,name=lsp_local_label,json=lspLocalLabel" json:"lsp_local_label,omitempty"`
	// Segment-routing iutgoing info of LSP
	SrlspOutgoingInfo []*MplsTeS2LSrOutgoingFwdInfo `protobuf:"bytes,4,rep,name=srlsp_outgoing_info,json=srlspOutgoingInfo" json:"srlsp_outgoing_info,omitempty"`
	// Output interface of LSP
	LspOutputInterface string `protobuf:"bytes,5,opt,name=lsp_output_interface,json=lspOutputInterface" json:"lsp_output_interface,omitempty"`
	// Output label of the LSP
	LspOutputLabel uint32 `protobuf:"varint,6,opt,name=lsp_output_label,json=lspOutputLabel" json:"lsp_output_label,omitempty"`
	// Path used by LSP
	LspHop []*RsvpMgmtEroSubobj `protobuf:"bytes,7,rep,name=lsp_hop,json=lspHop" json:"lsp_hop,omitempty"`
	// LSP State
	LspState             string   `protobuf:"bytes,8,opt,name=lsp_state,json=lspState" json:"lsp_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTePathProtectionLsp) Reset()         { *m = MplsTePathProtectionLsp{} }
func (m *MplsTePathProtectionLsp) String() string { return proto.CompactTextString(m) }
func (*MplsTePathProtectionLsp) ProtoMessage()    {}
func (*MplsTePathProtectionLsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_path_protection_a2caa39ccbfa33ab, []int{8}
}
func (m *MplsTePathProtectionLsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTePathProtectionLsp.Unmarshal(m, b)
}
func (m *MplsTePathProtectionLsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTePathProtectionLsp.Marshal(b, m, deterministic)
}
func (dst *MplsTePathProtectionLsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTePathProtectionLsp.Merge(dst, src)
}
func (m *MplsTePathProtectionLsp) XXX_Size() int {
	return xxx_messageInfo_MplsTePathProtectionLsp.Size(m)
}
func (m *MplsTePathProtectionLsp) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTePathProtectionLsp.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTePathProtectionLsp proto.InternalMessageInfo

func (m *MplsTePathProtectionLsp) GetLspUptime() uint32 {
	if m != nil {
		return m.LspUptime
	}
	return 0
}

func (m *MplsTePathProtectionLsp) GetPathProtectionLspId() uint32 {
	if m != nil {
		return m.PathProtectionLspId
	}
	return 0
}

func (m *MplsTePathProtectionLsp) GetLspLocalLabel() uint32 {
	if m != nil {
		return m.LspLocalLabel
	}
	return 0
}

func (m *MplsTePathProtectionLsp) GetSrlspOutgoingInfo() []*MplsTeS2LSrOutgoingFwdInfo {
	if m != nil {
		return m.SrlspOutgoingInfo
	}
	return nil
}

func (m *MplsTePathProtectionLsp) GetLspOutputInterface() string {
	if m != nil {
		return m.LspOutputInterface
	}
	return ""
}

func (m *MplsTePathProtectionLsp) GetLspOutputLabel() uint32 {
	if m != nil {
		return m.LspOutputLabel
	}
	return 0
}

func (m *MplsTePathProtectionLsp) GetLspHop() []*RsvpMgmtEroSubobj {
	if m != nil {
		return m.LspHop
	}
	return nil
}

func (m *MplsTePathProtectionLsp) GetLspState() string {
	if m != nil {
		return m.LspState
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsTePathProtection_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.mpls_te_path_protection_KEYS")
	proto.RegisterType((*MplsTePathProtection)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.mpls_te_path_protection")
	proto.RegisterType((*RsvpMgmtEroIpv4Subobj)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.rsvp_mgmt_ero_ipv4_subobj")
	proto.RegisterType((*RsvpMgmtEroUnnumSubobj)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.rsvp_mgmt_ero_unnum_subobj")
	proto.RegisterType((*RsvpMgmtEroSubobj)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.rsvp_mgmt_ero_subobj")
	proto.RegisterType((*TePpSwLogEntryBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.te_pp_sw_log_entry_bag")
	proto.RegisterType((*MplsTePpInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.mpls_te_pp_info")
	proto.RegisterType((*MplsTeS2LSrOutgoingFwdInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.mpls_te_s2l_sr_outgoing_fwd_info")
	proto.RegisterType((*MplsTePathProtectionLsp)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.mpls_te_path_protection_lsp")
}

func init() {
	proto.RegisterFile("mpls_te_path_protection.proto", fileDescriptor_mpls_te_path_protection_a2caa39ccbfa33ab)
}

var fileDescriptor_mpls_te_path_protection_a2caa39ccbfa33ab = []byte{
	// 1261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0x36, 0x89, 0xe3, 0x3c, 0xd7, 0x49, 0x3a, 0x29, 0xe9, 0x26, 0xfd, 0x67, 0x5c, 0xa8,
	0x8c, 0x40, 0x01, 0xa5, 0x2d, 0x7f, 0x04, 0x07, 0x28, 0x0d, 0xc2, 0x22, 0xad, 0xa3, 0x4d, 0x22,
	0xc1, 0x69, 0xb4, 0xde, 0x1d, 0x3b, 0x53, 0xad, 0x77, 0x46, 0x33, 0xb3, 0x4e, 0xdc, 0x03, 0x57,
	0xca, 0x0d, 0x71, 0xe0, 0x03, 0x70, 0xec, 0x01, 0x3e, 0x07, 0x9f, 0x86, 0xaf, 0x80, 0xe6, 0xed,
	0xac, 0xff, 0x6c, 0xe3, 0x4a, 0x3d, 0x24, 0x9c, 0x92, 0x79, 0xef, 0x37, 0x6f, 0x7f, 0xef, 0xff,
	0x18, 0x6e, 0x0f, 0x64, 0xa2, 0xa9, 0x61, 0x54, 0x86, 0xe6, 0x84, 0x4a, 0x25, 0x0c, 0x8b, 0x0c,
	0x17, 0xe9, 0x8e, 0xfd, 0x57, 0x90, 0x1f, 0x23, 0xae, 0x23, 0x41, 0xb9, 0xd0, 0xf4, 0x4c, 0xd1,
	0x02, 0x2b, 0x24, 0x53, 0x3b, 0xee, 0xb0, 0x63, 0xb2, 0x34, 0x65, 0x89, 0x76, 0x7f, 0xcb, 0x76,
	0xe6, 0xc9, 0x9b, 0x1a, 0x6e, 0xcd, 0xf9, 0x34, 0xfd, 0x61, 0xef, 0xa7, 0x43, 0xf2, 0x0e, 0x54,
	0x22, 0x6a, 0x46, 0x92, 0xf9, 0x5e, 0xc3, 0x6b, 0xad, 0x04, 0x4b, 0xd1, 0xd1, 0x48, 0x32, 0x72,
	0x13, 0x56, 0x9c, 0x41, 0x1e, 0xfb, 0x57, 0x1a, 0x5e, 0xab, 0x1e, 0x54, 0x73, 0x41, 0x3b, 0x26,
	0x77, 0xa1, 0xe6, 0x94, 0x69, 0x38, 0x60, 0xfe, 0x02, 0x5e, 0x84, 0x5c, 0xf4, 0x2c, 0x1c, 0xb0,
	0xe6, 0xcb, 0x25, 0xb8, 0x31, 0xe7, 0xab, 0xa4, 0x01, 0x57, 0xb9, 0xa6, 0xee, 0x7e, 0x26, 0xfd,
	0xdd, 0x86, 0xd7, 0xaa, 0x06, 0xc0, 0xf5, 0x11, 0x8a, 0x8e, 0x25, 0x79, 0x1f, 0x56, 0xb5, 0xc8,
	0x54, 0xc4, 0x68, 0x18, 0xc7, 0x8a, 0x69, 0xed, 0x3f, 0xc0, 0x2f, 0xd4, 0x73, 0xe9, 0x37, 0xb9,
	0x90, 0x7c, 0x0c, 0x1b, 0x31, 0xd3, 0x86, 0xa7, 0x21, 0x7a, 0x53, 0x60, 0x1f, 0x22, 0x96, 0x4c,
	0xa9, 0x8a, 0x0b, 0x1f, 0x01, 0x61, 0x67, 0x86, 0xa5, 0x31, 0x8b, 0xe9, 0xc4, 0xb9, 0x47, 0x88,
	0x5f, 0x2f, 0x34, 0x47, 0x85, 0x93, 0xbf, 0x7b, 0xb0, 0x56, 0xe2, 0xee, 0x7f, 0xda, 0xf0, 0x5a,
	0xb5, 0x5d, 0xbe, 0x73, 0x51, 0xd9, 0xda, 0x19, 0x07, 0x4d, 0x52, 0x9e, 0xf6, 0x44, 0xb0, 0x6a,
	0x01, 0x07, 0x93, 0xe0, 0xfd, 0xe1, 0x41, 0x2d, 0xca, 0x94, 0x62, 0xa9, 0xa1, 0x89, 0x96, 0xfe,
	0x67, 0x48, 0x28, 0xbb, 0x04, 0x42, 0xa5, 0xda, 0x49, 0xb4, 0x0c, 0xc0, 0x31, 0xd9, 0xd7, 0x12,
	0x89, 0x69, 0x13, 0xa6, 0x71, 0x77, 0x84, 0xc4, 0x3e, 0xff, 0x5f, 0x89, 0x39, 0x26, 0xfb, 0x5a,
	0x36, 0x7f, 0xf5, 0x60, 0x4b, 0xe9, 0xa1, 0xa4, 0x83, 0xfe, 0xc0, 0x50, 0xa6, 0x04, 0xe5, 0x72,
	0xf8, 0x90, 0xea, 0xac, 0x2b, 0xba, 0xcf, 0xc9, 0x7d, 0x58, 0xe3, 0x9a, 0x6a, 0xa3, 0x78, 0x64,
	0xa8, 0x12, 0x99, 0xc9, 0xdb, 0xa0, 0x1a, 0xd4, 0xb9, 0x3e, 0x44, 0x69, 0x60, 0x85, 0xb6, 0xe2,
	0xed, 0xd5, 0xa2, 0xc6, 0xae, 0xe4, 0x15, 0xcf, 0x94, 0x28, 0x6a, 0xeb, 0x1e, 0xd4, 0xa5, 0x62,
	0x3d, 0x7e, 0x46, 0x13, 0x96, 0xf6, 0xcd, 0x09, 0x36, 0x45, 0x3d, 0xb8, 0x9a, 0x0b, 0xf7, 0x51,
	0xd6, 0x7c, 0xe5, 0xc1, 0xf6, 0x2c, 0x97, 0x2c, 0x4d, 0xb3, 0xc1, 0xdb, 0x92, 0x69, 0xc1, 0x3a,
	0xfa, 0x91, 0x1a, 0xa6, 0x7a, 0x61, 0xc4, 0x26, 0x2d, 0xba, 0xca, 0x94, 0x68, 0x17, 0xe2, 0x76,
	0x4c, 0x9a, 0x50, 0xb7, 0x48, 0xb4, 0xa5, 0x2c, 0x2c, 0x6f, 0x55, 0xeb, 0x0b, 0x9a, 0x52, 0xed,
	0x98, 0x6c, 0x42, 0x45, 0x9b, 0xd0, 0x64, 0xda, 0x5f, 0x44, 0xa5, 0x3b, 0x35, 0x5f, 0x2e, 0xc0,
	0xf5, 0x59, 0xb2, 0x8e, 0xe6, 0x16, 0x54, 0xed, 0x69, 0x6a, 0x66, 0x2c, 0x33, 0x25, 0x70, 0x6a,
	0xfc, 0xe9, 0xc1, 0x06, 0x86, 0xd7, 0xc1, 0xa9, 0xe8, 0x3e, 0x67, 0x91, 0x41, 0x76, 0xb5, 0x5d,
	0x7d, 0x71, 0xd5, 0x30, 0x37, 0xc3, 0xc1, 0xba, 0x3d, 0xec, 0x29, 0x71, 0x98, 0x75, 0x3b, 0x48,
	0x86, 0xfc, 0xed, 0xc1, 0x16, 0xc6, 0xbd, 0xcb, 0x14, 0x8b, 0xcb, 0x54, 0x17, 0x90, 0xaa, 0xb9,
	0x2c, 0xaa, 0xd3, 0x05, 0x10, 0x6c, 0x4e, 0x68, 0x4d, 0x33, 0x6e, 0xfe, 0xbb, 0x00, 0x9b, 0xf9,
	0x50, 0xd0, 0xa7, 0x34, 0x11, 0x7d, 0xca, 0x52, 0xa3, 0x46, 0xb4, 0x1b, 0xf6, 0x49, 0x00, 0xf7,
	0xcb, 0x1d, 0xa0, 0x4f, 0xb9, 0x89, 0x4e, 0xc4, 0x90, 0x29, 0xca, 0x86, 0x76, 0x50, 0xf0, 0x34,
	0x66, 0x67, 0x98, 0xaa, 0x7a, 0xd0, 0x9c, 0x1d, 0x28, 0x87, 0x63, 0xec, 0x9e, 0x85, 0xb6, 0x2d,
	0x92, 0x7c, 0x01, 0x5b, 0x65, 0x9b, 0xe5, 0x5d, 0xb0, 0x39, 0x6b, 0x66, 0x3c, 0x34, 0xef, 0x40,
	0xad, 0xa7, 0xc4, 0xc0, 0x76, 0x61, 0x51, 0x6e, 0xf5, 0x60, 0xc5, 0x8a, 0xf6, 0xb5, 0x6c, 0xc7,
	0x64, 0x1b, 0x56, 0x8c, 0x28, 0xb4, 0x8b, 0xa8, 0x5d, 0x36, 0x22, 0xd7, 0x3d, 0x82, 0x1b, 0x71,
	0x68, 0x83, 0xdc, 0xa3, 0x4c, 0x29, 0xa1, 0x68, 0xcc, 0x8a, 0xb9, 0xbb, 0x84, 0xc8, 0xeb, 0x56,
	0xdd, 0xe9, 0xed, 0x59, 0xe5, 0x93, 0x42, 0x47, 0xbe, 0x83, 0xc6, 0x9c, 0x6b, 0x74, 0xc0, 0x93,
	0x84, 0x6b, 0x16, 0xf9, 0x15, 0xbc, 0x7f, 0xeb, 0xbc, 0xfb, 0x4f, 0x1d, 0x86, 0x7c, 0x0d, 0xb7,
	0xa6, 0x22, 0x17, 0x67, 0x2a, 0x9c, 0xb5, 0xb1, 0x8c, 0x36, 0xb6, 0x27, 0x98, 0x27, 0x0e, 0x32,
	0xb6, 0xd0, 0x86, 0x77, 0xdf, 0x90, 0x0b, 0xc5, 0x42, 0x2d, 0x52, 0xbf, 0x8a, 0x1d, 0x73, 0x67,
	0x5e, 0x1a, 0x02, 0x44, 0x35, 0xff, 0xaa, 0xc0, 0x5a, 0x69, 0x17, 0x90, 0x07, 0xb0, 0x69, 0xf8,
	0x00, 0x1d, 0x9d, 0x32, 0x6b, 0xa9, 0xe5, 0xa9, 0xdd, 0xb0, 0xda, 0x4e, 0x6f, 0x62, 0xeb, 0x90,
	0x45, 0xe4, 0x03, 0x58, 0x9f, 0x02, 0x1b, 0x61, 0xc2, 0xc4, 0xa5, 0x70, 0x6d, 0x22, 0x3f, 0xb2,
	0xe2, 0x12, 0x54, 0xb1, 0x30, 0x1e, 0xb9, 0x04, 0x4e, 0x41, 0x03, 0x2b, 0x26, 0x5f, 0xc1, 0x76,
	0x31, 0xec, 0x15, 0x13, 0xd2, 0xf0, 0x01, 0x7f, 0xc1, 0x62, 0x9a, 0x17, 0xaf, 0xcb, 0xab, 0xef,
	0x10, 0xc1, 0x04, 0xf0, 0x0c, 0xf5, 0xe4, 0x43, 0xb8, 0xf6, 0x7a, 0x5c, 0xf2, 0x14, 0xaf, 0xeb,
	0x52, 0x24, 0xec, 0x63, 0x20, 0xe6, 0x43, 0xa6, 0x34, 0x37, 0xa3, 0x7c, 0xe6, 0x54, 0xf2, 0xc7,
	0xc0, 0x58, 0x8a, 0x93, 0xe7, 0x4b, 0xd8, 0xe6, 0x7a, 0xa6, 0xd7, 0x68, 0x24, 0xd2, 0x1e, 0xef,
	0x67, 0x8a, 0xc5, 0x98, 0xbb, 0x6a, 0x70, 0x83, 0xeb, 0x83, 0x49, 0xd8, 0xbf, 0x1d, 0xab, 0xc9,
	0x13, 0xb8, 0x5b, 0x4e, 0x9c, 0xfb, 0x97, 0xc5, 0xb4, 0x3b, 0xb2, 0xb5, 0x5a, 0x45, 0x7a, 0x37,
	0x67, 0xd3, 0x76, 0x50, 0x80, 0x1e, 0x8f, 0xda, 0x31, 0x39, 0x86, 0xd6, 0x30, 0x4c, 0x78, 0xfc,
	0xda, 0x4a, 0xc2, 0xb3, 0x0d, 0x82, 0x48, 0x29, 0x3b, 0xe3, 0xda, 0x68, 0x7f, 0x05, 0x09, 0xdd,
	0x43, 0xfc, 0xc1, 0xac, 0xcd, 0xd0, 0x9c, 0x74, 0x10, 0xbb, 0x87, 0x50, 0xf2, 0x14, 0xde, 0x2b,
	0x7b, 0x96, 0x07, 0x89, 0x62, 0xf8, 0xb2, 0x34, 0x66, 0xea, 0x34, 0x1c, 0xf9, 0x80, 0x26, 0xef,
	0xce, 0xf8, 0x98, 0xd7, 0x42, 0x67, 0xc8, 0xd4, 0xb1, 0x83, 0x91, 0xdf, 0x3c, 0x80, 0x49, 0x90,
	0xfd, 0x1a, 0x8e, 0x3b, 0x79, 0x71, 0xe3, 0xee, 0xfc, 0xb9, 0x15, 0x4c, 0x71, 0x20, 0x8f, 0xe1,
	0xf6, 0xb8, 0x8a, 0xf2, 0xa6, 0xc3, 0x3a, 0x57, 0x6c, 0x10, 0xf2, 0x94, 0xa7, 0x7d, 0xff, 0x6a,
	0x1e, 0xfc, 0x59, 0xd0, 0x11, 0x1f, 0xb0, 0xa0, 0x80, 0x34, 0x7f, 0x86, 0x46, 0xc1, 0x5a, 0xef,
	0x26, 0x54, 0x2b, 0x2a, 0x32, 0xd3, 0x17, 0x3c, 0xed, 0xd3, 0xde, 0x69, 0x9c, 0x37, 0xd0, 0x27,
	0x70, 0xdd, 0x4e, 0x1e, 0x91, 0x19, 0x99, 0x99, 0xc9, 0xfa, 0x74, 0x4b, 0x8c, 0x24, 0x5a, 0x76,
	0x50, 0x35, 0xde, 0xa0, 0x76, 0xd3, 0x4e, 0xdd, 0x48, 0xc2, 0x2e, 0x2b, 0xba, 0x67, 0x75, 0x8c,
	0xde, 0xb7, 0xd2, 0xe6, 0x3f, 0x8b, 0x70, 0xf3, 0x0d, 0x4f, 0x12, 0x72, 0x1b, 0xc0, 0x5a, 0xca,
	0xac, 0x03, 0xcc, 0x35, 0xec, 0x4a, 0xa2, 0xe5, 0x31, 0x0a, 0x6c, 0x6f, 0x9f, 0x73, 0x6b, 0x32,
	0x6f, 0x37, 0x66, 0x0b, 0x2f, 0x1f, 0x98, 0xf7, 0x61, 0xcd, 0x82, 0x12, 0x11, 0x85, 0x89, 0x23,
	0x97, 0xf7, 0x6b, 0x3d, 0xd1, 0x72, 0xdf, 0x4a, 0x91, 0x1b, 0x79, 0xe5, 0xc1, 0x86, 0x56, 0xce,
	0x91, 0x3c, 0x26, 0x36, 0x1e, 0xfe, 0x62, 0x63, 0xa1, 0x55, 0xdb, 0x7d, 0x71, 0xf1, 0x6f, 0xb4,
	0x79, 0x19, 0x09, 0xae, 0x21, 0xad, 0x8e, 0x93, 0xb7, 0xdf, 0x94, 0xa4, 0xa5, 0xb7, 0x4a, 0x52,
	0xe5, 0xbc, 0x24, 0x91, 0x5f, 0x3c, 0x58, 0xb6, 0xd0, 0x13, 0x21, 0xfd, 0x65, 0x74, 0x3e, 0xbd,
	0xac, 0x3d, 0xef, 0x36, 0x7c, 0x25, 0xd1, 0xf2, 0x7b, 0x21, 0xed, 0xcf, 0x2b, 0x4b, 0xc4, 0x3e,
	0xb5, 0x98, 0x5b, 0x09, 0xd5, 0x44, 0xcb, 0x43, 0x7b, 0xee, 0x56, 0xf0, 0x37, 0xe1, 0x83, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x23, 0x0d, 0xcd, 0x28, 0x34, 0x0e, 0x00, 0x00,
}