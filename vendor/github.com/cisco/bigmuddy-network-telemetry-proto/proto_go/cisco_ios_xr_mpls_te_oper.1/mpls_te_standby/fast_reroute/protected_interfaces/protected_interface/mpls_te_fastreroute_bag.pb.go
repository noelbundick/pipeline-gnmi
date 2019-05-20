// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_fastreroute_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_fast_reroute_protected_interfaces_protected_interface

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

// MPLS TE Fast Reroute Information
type MplsTeFastrerouteBag_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	BackupTunnelName     string   `protobuf:"bytes,2,opt,name=backup_tunnel_name,json=backupTunnelName" json:"backup_tunnel_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeFastrerouteBag_KEYS) Reset()         { *m = MplsTeFastrerouteBag_KEYS{} }
func (m *MplsTeFastrerouteBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeFastrerouteBag_KEYS) ProtoMessage()    {}
func (*MplsTeFastrerouteBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_fastreroute_bag_d62ab19bd2550326, []int{0}
}
func (m *MplsTeFastrerouteBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeFastrerouteBag_KEYS.Unmarshal(m, b)
}
func (m *MplsTeFastrerouteBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeFastrerouteBag_KEYS.Marshal(b, m, deterministic)
}
func (dst *MplsTeFastrerouteBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeFastrerouteBag_KEYS.Merge(dst, src)
}
func (m *MplsTeFastrerouteBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeFastrerouteBag_KEYS.Size(m)
}
func (m *MplsTeFastrerouteBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeFastrerouteBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeFastrerouteBag_KEYS proto.InternalMessageInfo

func (m *MplsTeFastrerouteBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *MplsTeFastrerouteBag_KEYS) GetBackupTunnelName() string {
	if m != nil {
		return m.BackupTunnelName
	}
	return ""
}

type MplsTeFastrerouteBag struct {
	// Backup tunnel ID
	BackupTunnelId uint32 `protobuf:"varint,50,opt,name=backup_tunnel_id,json=backupTunnelId" json:"backup_tunnel_id,omitempty"`
	// Backup tunnel name
	BackupTunnelName string `protobuf:"bytes,51,opt,name=backup_tunnel_name,json=backupTunnelName" json:"backup_tunnel_name,omitempty"`
	// Backup state
	BackupStatus string `protobuf:"bytes,52,opt,name=backup_status,json=backupStatus" json:"backup_status,omitempty"`
	// Backup type
	BackupType string `protobuf:"bytes,53,opt,name=backup_type,json=backupType" json:"backup_type,omitempty"`
	// Backup usage
	BackupUsage string `protobuf:"bytes,54,opt,name=backup_usage,json=backupUsage" json:"backup_usage,omitempty"`
	// Protected interface Autobackup config
	ProtInterfaceAutobackupConfig string `protobuf:"bytes,55,opt,name=prot_interface_autobackup_config,json=protInterfaceAutobackupConfig" json:"prot_interface_autobackup_config,omitempty"`
	// Protected interface SRLG config
	ProtInterfaceSrlgConfig string `protobuf:"bytes,56,opt,name=prot_interface_srlg_config,json=protInterfaceSrlgConfig" json:"prot_interface_srlg_config,omitempty"`
	// Attribute Set Name
	TunnelAttributeSetName string `protobuf:"bytes,57,opt,name=tunnel_attribute_set_name,json=tunnelAttributeSetName" json:"tunnel_attribute_set_name,omitempty"`
	// TRUE if the auto-backup has an attribute set defined
	HasAttributeSet bool `protobuf:"varint,58,opt,name=has_attribute_set,json=hasAttributeSet" json:"has_attribute_set,omitempty"`
	// Flag to indicate the existence of attribute set found in database
	IsAttributeSetInDb bool `protobuf:"varint,59,opt,name=is_attribute_set_in_db,json=isAttributeSetInDb" json:"is_attribute_set_in_db,omitempty"`
	// Indicates if the recreate timer is running
	RecreateTimerIsRunning bool `protobuf:"varint,60,opt,name=recreate_timer_is_running,json=recreateTimerIsRunning" json:"recreate_timer_is_running,omitempty"`
	// Time Remaining in Recreate Timer (seconds)
	RecreateRemainingTime uint32 `protobuf:"varint,61,opt,name=recreate_remaining_time,json=recreateRemainingTime" json:"recreate_remaining_time,omitempty"`
	// Backup's source
	SourceAddress string `protobuf:"bytes,62,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	// Backup's destination
	DestinationAddress string `protobuf:"bytes,63,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// Backup's general status
	GeneralStatus string `protobuf:"bytes,64,opt,name=general_status,json=generalStatus" json:"general_status,omitempty"`
	// Backup's connection's status
	ConnectionStatus string `protobuf:"bytes,65,opt,name=connection_status,json=connectionStatus" json:"connection_status,omitempty"`
	// The output intf of the tunnel
	OutputInterfaceName string `protobuf:"bytes,66,opt,name=output_interface_name,json=outputInterfaceName" json:"output_interface_name,omitempty"`
	// Bandwidth pool type
	BandwidthPoolType string `protobuf:"bytes,67,opt,name=bandwidth_pool_type,json=bandwidthPoolType" json:"bandwidth_pool_type,omitempty"`
	// Bandwidth limit type
	BandwidthLimitType string `protobuf:"bytes,68,opt,name=bandwidth_limit_type,json=bandwidthLimitType" json:"bandwidth_limit_type,omitempty"`
	// Bandwidth (kbps)
	Bandwidth uint32 `protobuf:"varint,69,opt,name=bandwidth" json:"bandwidth,omitempty"`
	// Tunnel instance
	TunnelInstance uint32 `protobuf:"varint,70,opt,name=tunnel_instance,json=tunnelInstance" json:"tunnel_instance,omitempty"`
	// Bandwidth currently in use (kbps)
	InUseBandwidth uint32 `protobuf:"varint,71,opt,name=in_use_bandwidth,json=inUseBandwidth" json:"in_use_bandwidth,omitempty"`
	// Bandwidth soft preempted and rerouted over the backup(kbps)
	SoftPreemptedInUseBandwidth uint32 `protobuf:"varint,72,opt,name=soft_preempted_in_use_bandwidth,json=softPreemptedInUseBandwidth" json:"soft_preempted_in_use_bandwidth,omitempty"`
	// Number of LSPs
	LsPs uint32 `protobuf:"varint,73,opt,name=ls_ps,json=lsPs" json:"ls_ps,omitempty"`
	// Number of S2L Families
	S2LFamilies uint32 `protobuf:"varint,74,opt,name=s2_l_families,json=s2LFamilies" json:"s2_l_families,omitempty"`
	// Number of P2MP S2Ls
	S2Ls uint32 `protobuf:"varint,75,opt,name=s2_ls,json=s2Ls" json:"s2_ls,omitempty"`
	// Number of LSPs in FRR active state
	FrrActiveLsPs uint32 `protobuf:"varint,76,opt,name=frr_active_ls_ps,json=frrActiveLsPs" json:"frr_active_ls_ps,omitempty"`
	// Number of soft preempted LSPs routed over backup
	FrrActiveSoftPreemptedLsPs uint32   `protobuf:"varint,77,opt,name=frr_active_soft_preempted_ls_ps,json=frrActiveSoftPreemptedLsPs" json:"frr_active_soft_preempted_ls_ps,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *MplsTeFastrerouteBag) Reset()         { *m = MplsTeFastrerouteBag{} }
func (m *MplsTeFastrerouteBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeFastrerouteBag) ProtoMessage()    {}
func (*MplsTeFastrerouteBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_fastreroute_bag_d62ab19bd2550326, []int{1}
}
func (m *MplsTeFastrerouteBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeFastrerouteBag.Unmarshal(m, b)
}
func (m *MplsTeFastrerouteBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeFastrerouteBag.Marshal(b, m, deterministic)
}
func (dst *MplsTeFastrerouteBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeFastrerouteBag.Merge(dst, src)
}
func (m *MplsTeFastrerouteBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeFastrerouteBag.Size(m)
}
func (m *MplsTeFastrerouteBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeFastrerouteBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeFastrerouteBag proto.InternalMessageInfo

func (m *MplsTeFastrerouteBag) GetBackupTunnelId() uint32 {
	if m != nil {
		return m.BackupTunnelId
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetBackupTunnelName() string {
	if m != nil {
		return m.BackupTunnelName
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBackupStatus() string {
	if m != nil {
		return m.BackupStatus
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBackupType() string {
	if m != nil {
		return m.BackupType
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBackupUsage() string {
	if m != nil {
		return m.BackupUsage
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetProtInterfaceAutobackupConfig() string {
	if m != nil {
		return m.ProtInterfaceAutobackupConfig
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetProtInterfaceSrlgConfig() string {
	if m != nil {
		return m.ProtInterfaceSrlgConfig
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetTunnelAttributeSetName() string {
	if m != nil {
		return m.TunnelAttributeSetName
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetHasAttributeSet() bool {
	if m != nil {
		return m.HasAttributeSet
	}
	return false
}

func (m *MplsTeFastrerouteBag) GetIsAttributeSetInDb() bool {
	if m != nil {
		return m.IsAttributeSetInDb
	}
	return false
}

func (m *MplsTeFastrerouteBag) GetRecreateTimerIsRunning() bool {
	if m != nil {
		return m.RecreateTimerIsRunning
	}
	return false
}

func (m *MplsTeFastrerouteBag) GetRecreateRemainingTime() uint32 {
	if m != nil {
		return m.RecreateRemainingTime
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetGeneralStatus() string {
	if m != nil {
		return m.GeneralStatus
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetConnectionStatus() string {
	if m != nil {
		return m.ConnectionStatus
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetOutputInterfaceName() string {
	if m != nil {
		return m.OutputInterfaceName
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBandwidthPoolType() string {
	if m != nil {
		return m.BandwidthPoolType
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBandwidthLimitType() string {
	if m != nil {
		return m.BandwidthLimitType
	}
	return ""
}

func (m *MplsTeFastrerouteBag) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetTunnelInstance() uint32 {
	if m != nil {
		return m.TunnelInstance
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetInUseBandwidth() uint32 {
	if m != nil {
		return m.InUseBandwidth
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetSoftPreemptedInUseBandwidth() uint32 {
	if m != nil {
		return m.SoftPreemptedInUseBandwidth
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetLsPs() uint32 {
	if m != nil {
		return m.LsPs
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetS2LFamilies() uint32 {
	if m != nil {
		return m.S2LFamilies
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetS2Ls() uint32 {
	if m != nil {
		return m.S2Ls
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetFrrActiveLsPs() uint32 {
	if m != nil {
		return m.FrrActiveLsPs
	}
	return 0
}

func (m *MplsTeFastrerouteBag) GetFrrActiveSoftPreemptedLsPs() uint32 {
	if m != nil {
		return m.FrrActiveSoftPreemptedLsPs
	}
	return 0
}

// Statistics - count and timestamp on an item
type MplsTeStatItemT struct {
	// Number of occurence
	Count uint32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	// Timestamp, when occured last
	LastTimeItOccured    uint32   `protobuf:"varint,2,opt,name=last_time_it_occured,json=lastTimeItOccured" json:"last_time_it_occured,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeStatItemT) Reset()         { *m = MplsTeStatItemT{} }
func (m *MplsTeStatItemT) String() string { return proto.CompactTextString(m) }
func (*MplsTeStatItemT) ProtoMessage()    {}
func (*MplsTeStatItemT) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_fastreroute_bag_d62ab19bd2550326, []int{2}
}
func (m *MplsTeStatItemT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeStatItemT.Unmarshal(m, b)
}
func (m *MplsTeStatItemT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeStatItemT.Marshal(b, m, deterministic)
}
func (dst *MplsTeStatItemT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeStatItemT.Merge(dst, src)
}
func (m *MplsTeStatItemT) XXX_Size() int {
	return xxx_messageInfo_MplsTeStatItemT.Size(m)
}
func (m *MplsTeStatItemT) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeStatItemT.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeStatItemT proto.InternalMessageInfo

func (m *MplsTeStatItemT) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MplsTeStatItemT) GetLastTimeItOccured() uint32 {
	if m != nil {
		return m.LastTimeItOccured
	}
	return 0
}

// Statistics of all types of item operations and timestamps
type MplsTeStatsItemT struct {
	// Statistics on all types of operation on an item
	ItemStatistics       []*MplsTeStatItemT `protobuf:"bytes,1,rep,name=item_statistics,json=itemStatistics" json:"item_statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MplsTeStatsItemT) Reset()         { *m = MplsTeStatsItemT{} }
func (m *MplsTeStatsItemT) String() string { return proto.CompactTextString(m) }
func (*MplsTeStatsItemT) ProtoMessage()    {}
func (*MplsTeStatsItemT) Descriptor() ([]byte, []int) {
	return fileDescriptor_mpls_te_fastreroute_bag_d62ab19bd2550326, []int{3}
}
func (m *MplsTeStatsItemT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeStatsItemT.Unmarshal(m, b)
}
func (m *MplsTeStatsItemT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeStatsItemT.Marshal(b, m, deterministic)
}
func (dst *MplsTeStatsItemT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeStatsItemT.Merge(dst, src)
}
func (m *MplsTeStatsItemT) XXX_Size() int {
	return xxx_messageInfo_MplsTeStatsItemT.Size(m)
}
func (m *MplsTeStatsItemT) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeStatsItemT.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeStatsItemT proto.InternalMessageInfo

func (m *MplsTeStatsItemT) GetItemStatistics() []*MplsTeStatItemT {
	if m != nil {
		return m.ItemStatistics
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsTeFastrerouteBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fast_reroute.protected_interfaces.protected_interface.mpls_te_fastreroute_bag_KEYS")
	proto.RegisterType((*MplsTeFastrerouteBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fast_reroute.protected_interfaces.protected_interface.mpls_te_fastreroute_bag")
	proto.RegisterType((*MplsTeStatItemT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fast_reroute.protected_interfaces.protected_interface.mpls_te_stat_item_t")
	proto.RegisterType((*MplsTeStatsItemT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.fast_reroute.protected_interfaces.protected_interface.mpls_te_stats_item_t")
}

func init() {
	proto.RegisterFile("mpls_te_fastreroute_bag.proto", fileDescriptor_mpls_te_fastreroute_bag_d62ab19bd2550326)
}

var fileDescriptor_mpls_te_fastreroute_bag_d62ab19bd2550326 = []byte{
	// 840 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0x5f, 0x6f, 0x1b, 0x45,
	0x14, 0xc5, 0x65, 0xa0, 0x88, 0xde, 0xe0, 0xfc, 0x19, 0xa7, 0xc9, 0x50, 0x5a, 0xc5, 0x18, 0x55,
	0xb5, 0x00, 0xb9, 0xc8, 0x85, 0x42, 0x29, 0xff, 0xdc, 0xa4, 0x2d, 0x4b, 0x0d, 0x44, 0x76, 0xfb,
	0x80, 0x84, 0x34, 0x1a, 0xef, 0x8e, 0x9d, 0x11, 0xbb, 0x33, 0xab, 0xb9, 0xb3, 0x40, 0x3e, 0x0c,
	0x4f, 0x7c, 0x07, 0x3e, 0x1f, 0x9a, 0x3b, 0xbb, 0x1b, 0x3b, 0x4a, 0x1e, 0x79, 0x4a, 0x7c, 0xce,
	0xef, 0x9c, 0xdd, 0x9d, 0xbd, 0x33, 0x0b, 0x77, 0x8b, 0x32, 0x47, 0xe1, 0x95, 0x58, 0x4a, 0xf4,
	0x4e, 0x39, 0x5b, 0x79, 0x25, 0x16, 0x72, 0x35, 0x2a, 0x9d, 0xf5, 0x96, 0x89, 0x54, 0x63, 0x6a,
	0x85, 0xb6, 0x28, 0xfe, 0x72, 0xa2, 0x61, 0x6d, 0xa9, 0xdc, 0xa8, 0xf9, 0x81, 0x5e, 0x9a, 0x6c,
	0x71, 0x3e, 0x0a, 0x05, 0xa2, 0x6e, 0xa0, 0xb4, 0x4a, 0xbd, 0xca, 0x84, 0x36, 0x5e, 0xb9, 0xa5,
	0x4c, 0x15, 0x5e, 0x25, 0x0e, 0x10, 0xee, 0x5c, 0x73, 0x07, 0xe2, 0xe5, 0xb3, 0x5f, 0xe7, 0xec,
	0x1e, 0x6c, 0xb7, 0xb0, 0x30, 0xb2, 0x50, 0xbc, 0xd3, 0xef, 0x0c, 0x6f, 0xce, 0xba, 0xad, 0xfa,
	0xb3, 0x2c, 0x14, 0xfb, 0x04, 0xd8, 0x42, 0xa6, 0xbf, 0x57, 0xa5, 0xf0, 0x95, 0x31, 0x2a, 0x8f,
	0xe8, 0x1b, 0x84, 0xee, 0x46, 0xe7, 0x15, 0x19, 0x81, 0x1e, 0xfc, 0x03, 0x70, 0x78, 0xcd, 0x55,
	0xd9, 0x10, 0x76, 0x37, 0x9b, 0x74, 0xc6, 0xc7, 0xfd, 0xce, 0xb0, 0x3b, 0xdb, 0x5e, 0xef, 0x49,
	0xb2, 0x6b, 0xae, 0xf9, 0xf0, 0xea, 0x6b, 0xb2, 0x0f, 0xa1, 0x5b, 0xd3, 0xe8, 0xa5, 0xaf, 0x90,
	0x7f, 0x46, 0xe0, 0xbb, 0x51, 0x9c, 0x93, 0xc6, 0x8e, 0x60, 0xab, 0xa9, 0x3c, 0x2f, 0x15, 0xff,
	0x9c, 0x10, 0xa8, 0xbb, 0xce, 0x4b, 0xc5, 0x3e, 0x80, 0x3a, 0x20, 0x2a, 0x94, 0x2b, 0xc5, 0x1f,
	0x11, 0x51, 0x87, 0x5e, 0x07, 0x89, 0xbd, 0x80, 0x7e, 0x58, 0xe8, 0x8b, 0x35, 0x16, 0xb2, 0xf2,
	0xb6, 0x4e, 0xa5, 0xd6, 0x2c, 0xf5, 0x8a, 0x7f, 0x41, 0xb1, 0xbb, 0x81, 0x4b, 0x1a, 0x6c, 0xd2,
	0x52, 0xc7, 0x04, 0xb1, 0x27, 0x70, 0xfb, 0x52, 0x11, 0xba, 0x7c, 0xd5, 0x54, 0x7c, 0x49, 0x15,
	0x87, 0x1b, 0x15, 0x73, 0x97, 0xaf, 0xea, 0xf0, 0x63, 0x78, 0xaf, 0x5e, 0x15, 0xe9, 0xbd, 0xd3,
	0x8b, 0xb0, 0xbc, 0xa8, 0x7c, 0x5c, 0xa3, 0xc7, 0x94, 0x3d, 0x88, 0xc0, 0xa4, 0xf1, 0xe7, 0xca,
	0xd3, 0x4a, 0x7d, 0x04, 0x7b, 0x67, 0x12, 0x37, 0x73, 0xfc, 0xab, 0x7e, 0x67, 0xf8, 0xce, 0x6c,
	0xe7, 0x4c, 0xe2, 0x3a, 0xcf, 0xc6, 0x70, 0xa0, 0x2f, 0xa1, 0x42, 0x1b, 0x91, 0x2d, 0xf8, 0x13,
	0x0a, 0x30, 0xbd, 0xc1, 0x27, 0xe6, 0x64, 0x11, 0x6e, 0xcd, 0xa9, 0xd4, 0x29, 0xe9, 0x95, 0xf0,
	0xba, 0x50, 0x4e, 0x68, 0x14, 0xae, 0x32, 0x46, 0x9b, 0x15, 0xff, 0x9a, 0x62, 0x07, 0x0d, 0xf0,
	0x2a, 0xf8, 0x09, 0xce, 0xa2, 0xcb, 0x1e, 0xc1, 0x61, 0x1b, 0x75, 0xaa, 0x90, 0x3a, 0xa8, 0x54,
	0xc2, 0xbf, 0xa1, 0x19, 0xb9, 0xd5, 0xd8, 0xb3, 0xc6, 0x0d, 0x0d, 0x61, 0x8a, 0xd1, 0x56, 0x2e,
	0xbc, 0x8b, 0x2c, 0x73, 0x0a, 0x91, 0x7f, 0x1b, 0xa7, 0x38, 0xaa, 0x93, 0x28, 0xb2, 0x07, 0xd0,
	0xcb, 0x14, 0x7a, 0x6d, 0xa4, 0xd7, 0xd6, 0xb4, 0xec, 0x77, 0xc4, 0xb2, 0x35, 0xab, 0x09, 0xdc,
	0x83, 0xed, 0x95, 0x32, 0xca, 0xc9, 0xbc, 0x99, 0xaa, 0xef, 0x63, 0x6f, 0xad, 0xd6, 0x63, 0xf5,
	0x31, 0xec, 0xa5, 0xd6, 0x18, 0x95, 0x52, 0x6d, 0x4d, 0x4e, 0xe2, 0xa0, 0x5e, 0x18, 0x35, 0x3c,
	0x86, 0x5b, 0xb6, 0xf2, 0x65, 0xb5, 0xfe, 0xe2, 0xe9, 0xad, 0x3d, 0xa5, 0x40, 0x2f, 0x9a, 0xc9,
	0xc6, 0xf6, 0x1b, 0x41, 0x6f, 0x21, 0x4d, 0xf6, 0xa7, 0xce, 0xfc, 0x99, 0x28, 0xad, 0xcd, 0xe3,
	0xfc, 0x1e, 0x53, 0x62, 0xaf, 0xb5, 0x4e, 0xad, 0xcd, 0x69, 0x8c, 0x3f, 0x85, 0xfd, 0x0b, 0x3e,
	0xd7, 0x85, 0xf6, 0x31, 0x70, 0x12, 0x9f, 0xb4, 0xf5, 0xa6, 0xc1, 0xa2, 0xc4, 0x1d, 0xb8, 0xd9,
	0xaa, 0xfc, 0x19, 0xad, 0xf5, 0x85, 0xc0, 0xee, 0xc3, 0x4e, 0xb3, 0x5b, 0x4d, 0x38, 0x8f, 0x52,
	0xc5, 0x9f, 0xc7, 0x3d, 0x1b, 0xe5, 0xa4, 0x56, 0xc3, 0xee, 0xd6, 0x46, 0x54, 0x18, 0xf6, 0x7a,
	0xd3, 0xf6, 0x22, 0x92, 0xda, 0xbc, 0x46, 0xf5, 0xb4, 0xad, 0x3c, 0x81, 0x23, 0xb4, 0x4b, 0x2f,
	0x4a, 0xa7, 0x54, 0x51, 0xc6, 0x43, 0xeb, 0x52, 0xf0, 0x07, 0x0a, 0xbe, 0x1f, 0xb0, 0xd3, 0x86,
	0x4a, 0x36, 0x5b, 0x7a, 0x70, 0x23, 0x47, 0x51, 0x22, 0x4f, 0x88, 0x7d, 0x2b, 0xc7, 0x53, 0x64,
	0x03, 0xe8, 0xe2, 0x58, 0xe4, 0x62, 0x29, 0x0b, 0x9d, 0x6b, 0x85, 0xfc, 0x47, 0x32, 0xb7, 0x70,
	0x3c, 0x7d, 0x5e, 0x4b, 0x21, 0x18, 0x18, 0xe4, 0x2f, 0x63, 0x10, 0xc7, 0x53, 0x64, 0xf7, 0x61,
	0x77, 0xe9, 0x9c, 0x90, 0xa9, 0xd7, 0x7f, 0x28, 0x11, 0x8b, 0xa7, 0xe4, 0x77, 0x97, 0xce, 0x4d,
	0x48, 0x9e, 0x86, 0x2b, 0x1c, 0xc3, 0xd1, 0x1a, 0x78, 0xe9, 0x39, 0x62, 0xee, 0x27, 0xca, 0xdd,
	0x6e, 0x73, 0xf3, 0xf5, 0xa7, 0x08, 0x25, 0x83, 0xdf, 0xa0, 0xb7, 0x76, 0xc6, 0x7b, 0xa1, 0xbd,
	0x2a, 0x84, 0x67, 0xfb, 0x70, 0x23, 0xb5, 0x95, 0xf1, 0x74, 0x10, 0x77, 0x67, 0xf1, 0x07, 0x7b,
	0x00, 0xfb, 0x79, 0xf8, 0x00, 0x84, 0xbd, 0x20, 0xb4, 0x17, 0x36, 0x4d, 0x2b, 0xa7, 0x32, 0x3a,
	0x82, 0xbb, 0xb3, 0xbd, 0xe0, 0x85, 0x9d, 0x90, 0xf8, 0x5f, 0xa2, 0x31, 0xf8, 0xb7, 0x03, 0xfb,
	0xeb, 0xf5, 0xd8, 0xf4, 0xff, 0xdd, 0x81, 0x1d, 0xfa, 0x37, 0xa8, 0x1a, 0xbd, 0x4e, 0x91, 0x77,
	0xfa, 0x6f, 0x0e, 0xb7, 0xc6, 0x7e, 0xf4, 0x3f, 0x7f, 0x8d, 0x46, 0x57, 0x3c, 0xef, 0x6c, 0x3b,
	0xfc, 0x9d, 0xb7, 0xf7, 0xb2, 0x78, 0x9b, 0xbe, 0x8c, 0x0f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0xb7, 0xbb, 0xa0, 0x3a, 0x07, 0x00, 0x00,
}
