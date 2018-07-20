// Code generated by protoc-gen-go.
// source: ipv4_rib_edm_route.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_routes_route is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_route.proto

It has these top-level messages:
	Ipv4RibEdmRoute_KEYS
	Ipv4RibEdmRoute
	Ipv4RibEdmPath
	Ipv4RibEdmPathItem
	RibEdmLocalLabel
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_routes_route

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

// Information of a rib route head and rib proto route
type Ipv4RibEdmRoute_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	Address        string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	PrefixLength   uint32 `protobuf:"varint,6,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	NextHopAddress string `protobuf:"bytes,7,opt,name=next_hop_address,json=nextHopAddress" json:"next_hop_address,omitempty"`
	InterfaceName  string `protobuf:"bytes,8,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *Ipv4RibEdmRoute_KEYS) Reset()                    { *m = Ipv4RibEdmRoute_KEYS{} }
func (m *Ipv4RibEdmRoute_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute_KEYS) ProtoMessage()               {}
func (*Ipv4RibEdmRoute_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv4RibEdmRoute_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv4RibEdmRoute_KEYS) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv4RibEdmRoute struct {
	// Route prefix
	Prefix string `protobuf:"bytes,50,opt,name=prefix" json:"prefix,omitempty"`
	// Length of prefix
	PrefixLength uint32 `protobuf:"varint,51,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	// Route version
	RouteVersion uint32 `protobuf:"varint,52,opt,name=route_version,json=routeVersion" json:"route_version,omitempty"`
	// Protocol advertising the route
	ProtocolId uint32 `protobuf:"varint,53,opt,name=protocol_id,json=protocolId" json:"protocol_id,omitempty"`
	//  Name of Protocol
	ProtocolName string `protobuf:"bytes,54,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Instance name
	Instance string `protobuf:"bytes,55,opt,name=instance" json:"instance,omitempty"`
	// Client adding the route to RIB
	ClientId uint32 `protobuf:"varint,56,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Route type
	RouteType uint32 `protobuf:"varint,57,opt,name=route_type,json=routeType" json:"route_type,omitempty"`
	// Route priority
	Priority uint32 `protobuf:"varint,58,opt,name=priority" json:"priority,omitempty"`
	// SVD Type of route
	SvdType uint32 `protobuf:"varint,59,opt,name=svd_type,json=svdType" json:"svd_type,omitempty"`
	// Route flags
	Flags uint32 `protobuf:"varint,60,opt,name=flags" json:"flags,omitempty"`
	// Extended Route flags
	ExtendedFlags uint64 `protobuf:"varint,61,opt,name=extended_flags,json=extendedFlags" json:"extended_flags,omitempty"`
	// Opaque proto specific info
	Tag uint32 `protobuf:"varint,62,opt,name=tag" json:"tag,omitempty"`
	// Distance of the route
	Distance uint32 `protobuf:"varint,63,opt,name=distance" json:"distance,omitempty"`
	// Diversion distance of the route
	DiversionDistance uint32 `protobuf:"varint,64,opt,name=diversion_distance,json=diversionDistance" json:"diversion_distance,omitempty"`
	// Route metric
	Metric uint32 `protobuf:"varint,65,opt,name=metric" json:"metric,omitempty"`
	// Number of paths
	PathsCount uint32 `protobuf:"varint,66,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// BGP Attribute ID
	AttributeIdentity uint32 `protobuf:"varint,67,opt,name=attribute_identity,json=attributeIdentity" json:"attribute_identity,omitempty"`
	// BGP Traffic Index
	TrafficIndex uint32 `protobuf:"varint,68,opt,name=traffic_index,json=trafficIndex" json:"traffic_index,omitempty"`
	// Route ip precedence
	RoutePrecedence uint32 `protobuf:"varint,69,opt,name=route_precedence,json=routePrecedence" json:"route_precedence,omitempty"`
	// Route qos group
	QosGroup uint32 `protobuf:"varint,70,opt,name=qos_group,json=qosGroup" json:"qos_group,omitempty"`
	// Flow tag
	FlowTag uint32 `protobuf:"varint,71,opt,name=flow_tag,json=flowTag" json:"flow_tag,omitempty"`
	// Forward Class
	FwdClass uint32 `protobuf:"varint,72,opt,name=fwd_class,json=fwdClass" json:"fwd_class,omitempty"`
	// Number of pic paths in this route
	PicCount uint32 `protobuf:"varint,73,opt,name=pic_count,json=picCount" json:"pic_count,omitempty"`
	// Is the route active or backup
	Active bool `protobuf:"varint,74,opt,name=active" json:"active,omitempty"`
	// Route has a diversion path
	Diversion bool `protobuf:"varint,75,opt,name=diversion" json:"diversion,omitempty"`
	// Diversion route protocol name
	DiversionProtoName string `protobuf:"bytes,76,opt,name=diversion_proto_name,json=diversionProtoName" json:"diversion_proto_name,omitempty"`
	// Age of route (seconds)
	RouteAge uint32 `protobuf:"varint,77,opt,name=route_age,json=routeAge" json:"route_age,omitempty"`
	// Local label of the route
	RouteLabel uint32 `protobuf:"varint,78,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Route Version
	Version uint32 `protobuf:"varint,79,opt,name=version" json:"version,omitempty"`
	// Table Version
	TblVersion uint64 `protobuf:"varint,80,opt,name=tbl_version,json=tblVersion" json:"tbl_version,omitempty"`
	// Route modification time(nanoseconds)
	RouteModifyTime uint64 `protobuf:"varint,81,opt,name=route_modify_time,json=routeModifyTime" json:"route_modify_time,omitempty"`
	// Path(s) of the route
	RoutePath *Ipv4RibEdmPath `protobuf:"bytes,82,opt,name=route_path,json=routePath" json:"route_path,omitempty"`
}

func (m *Ipv4RibEdmRoute) Reset()                    { *m = Ipv4RibEdmRoute{} }
func (m *Ipv4RibEdmRoute) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute) ProtoMessage()               {}
func (*Ipv4RibEdmRoute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4RibEdmRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteVersion() uint32 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetSvdType() uint32 {
	if m != nil {
		return m.SvdType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetExtendedFlags() uint64 {
	if m != nil {
		return m.ExtendedFlags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDiversionDistance() uint32 {
	if m != nil {
		return m.DiversionDistance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetAttributeIdentity() uint32 {
	if m != nil {
		return m.AttributeIdentity
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTrafficIndex() uint32 {
	if m != nil {
		return m.TrafficIndex
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePrecedence() uint32 {
	if m != nil {
		return m.RoutePrecedence
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlowTag() uint32 {
	if m != nil {
		return m.FlowTag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFwdClass() uint32 {
	if m != nil {
		return m.FwdClass
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPicCount() uint32 {
	if m != nil {
		return m.PicCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversion() bool {
	if m != nil {
		return m.Diversion
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversionProtoName() string {
	if m != nil {
		return m.DiversionProtoName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetRouteAge() uint32 {
	if m != nil {
		return m.RouteAge
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTblVersion() uint64 {
	if m != nil {
		return m.TblVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteModifyTime() uint64 {
	if m != nil {
		return m.RouteModifyTime
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePath() *Ipv4RibEdmPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

// Information of a rib path
type Ipv4RibEdmPath struct {
	// Next path
	Ipv4RibEdmPath []*Ipv4RibEdmPathItem `protobuf:"bytes,1,rep,name=ipv4_rib_edm_path,json=ipv4RibEdmPath" json:"ipv4_rib_edm_path,omitempty"`
}

func (m *Ipv4RibEdmPath) Reset()                    { *m = Ipv4RibEdmPath{} }
func (m *Ipv4RibEdmPath) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPath) ProtoMessage()               {}
func (*Ipv4RibEdmPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv4RibEdmPath) GetIpv4RibEdmPath() []*Ipv4RibEdmPathItem {
	if m != nil {
		return m.Ipv4RibEdmPath
	}
	return nil
}

type Ipv4RibEdmPathItem struct {
	// Nexthop
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// Infosource
	InformationSource string `protobuf:"bytes,2,opt,name=information_source,json=informationSource" json:"information_source,omitempty"`
	// V6 nexthop
	V6Nexthop string `protobuf:"bytes,3,opt,name=v6_nexthop,json=v6Nexthop" json:"v6_nexthop,omitempty"`
	// Interface Name
	InterfaceName string `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Metrics
	Metric uint32 `protobuf:"varint,5,opt,name=metric" json:"metric,omitempty"`
	// Load Metrics
	LoadMetric uint32 `protobuf:"varint,6,opt,name=load_metric,json=loadMetric" json:"load_metric,omitempty"`
	// Flags extended to 64 bits
	Flags64 uint64 `protobuf:"varint,7,opt,name=flags64" json:"flags64,omitempty"`
	// Flags
	Flags uint32 `protobuf:"varint,8,opt,name=flags" json:"flags,omitempty"`
	// Private Flags
	PrivateFlags uint32 `protobuf:"varint,9,opt,name=private_flags,json=privateFlags" json:"private_flags,omitempty"`
	// Looping path
	Looped bool `protobuf:"varint,10,opt,name=looped" json:"looped,omitempty"`
	// Nexthop tableid
	NextHopTableId uint32 `protobuf:"varint,11,opt,name=next_hop_table_id,json=nextHopTableId" json:"next_hop_table_id,omitempty"`
	// VRF Name of the nh table
	NextHopVrfName string `protobuf:"bytes,12,opt,name=next_hop_vrf_name,json=nextHopVrfName" json:"next_hop_vrf_name,omitempty"`
	// NH table name
	NextHopTableName string `protobuf:"bytes,13,opt,name=next_hop_table_name,json=nextHopTableName" json:"next_hop_table_name,omitempty"`
	// NH afi
	NextHopAfi uint32 `protobuf:"varint,14,opt,name=next_hop_afi,json=nextHopAfi" json:"next_hop_afi,omitempty"`
	// NH safi
	NextHopSafi uint32 `protobuf:"varint,15,opt,name=next_hop_safi,json=nextHopSafi" json:"next_hop_safi,omitempty"`
	// Label associated with this path
	RouteLabel uint32 `protobuf:"varint,16,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Tunnel ID associated with this path
	TunnelId uint32 `protobuf:"varint,17,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	// Path id of this path
	Pathid uint32 `protobuf:"varint,18,opt,name=pathid" json:"pathid,omitempty"`
	// Path id of this path's backup
	BackupPathid uint32 `protobuf:"varint,19,opt,name=backup_pathid,json=backupPathid" json:"backup_pathid,omitempty"`
	// Refcnt of backup
	RefCntOfBackup uint32 `protobuf:"varint,20,opt,name=ref_cnt_of_backup,json=refCntOfBackup" json:"ref_cnt_of_backup,omitempty"`
	// Number of extended communities
	NumberOfExtendedCommunities uint32 `protobuf:"varint,21,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities" json:"number_of_extended_communities,omitempty"`
	// MVPN attribute present
	MvpnPresent bool `protobuf:"varint,22,opt,name=mvpn_present,json=mvpnPresent" json:"mvpn_present,omitempty"`
	// Path RT present
	PathrtPresent           bool `protobuf:"varint,23,opt,name=pathrt_present,json=pathrtPresent" json:"pathrt_present,omitempty"`
	VrfimportrtPresent      bool `protobuf:"varint,24,opt,name=vrfimportrt_present,json=vrfimportrtPresent" json:"vrfimportrt_present,omitempty"`
	SourceasrtPresent       bool `protobuf:"varint,25,opt,name=sourceasrt_present,json=sourceasrtPresent" json:"sourceasrt_present,omitempty"`
	SourcerdPresent         bool `protobuf:"varint,26,opt,name=sourcerd_present,json=sourcerdPresent" json:"sourcerd_present,omitempty"`
	SegmentedNexthopPresent bool `protobuf:"varint,27,opt,name=segmented_nexthop_present,json=segmentedNexthopPresent" json:"segmented_nexthop_present,omitempty"`
	// NHID associated with this path
	NextHopId uint32 `protobuf:"varint,28,opt,name=next_hop_id,json=nextHopId" json:"next_hop_id,omitempty"`
	// NHID references
	NextHopIdRefcount uint32 `protobuf:"varint,29,opt,name=next_hop_id_refcount,json=nextHopIdRefcount" json:"next_hop_id_refcount,omitempty"`
	// OSPF area associated with the path
	OspfAreaId string `protobuf:"bytes,30,opt,name=ospf_area_id,json=ospfAreaId" json:"ospf_area_id,omitempty"`
	// Remote backup node address
	RemoteBackupAddr [][]byte `protobuf:"bytes,31,rep,name=remote_backup_addr,json=remoteBackupAddr,proto3" json:"remote_backup_addr,omitempty"`
	// Path has a label stack
	HasLabelstk bool `protobuf:"varint,32,opt,name=has_labelstk,json=hasLabelstk" json:"has_labelstk,omitempty"`
	// Number of labels in stack
	NumLabels uint32 `protobuf:"varint,33,opt,name=num_labels,json=numLabels" json:"num_labels,omitempty"`
	// Labels for this path
	Labelstk []uint32 `protobuf:"varint,34,rep,packed,name=labelstk" json:"labelstk,omitempty"`
	// binding Label for this path
	BindingLabel uint32 `protobuf:"varint,35,opt,name=binding_label,json=bindingLabel" json:"binding_label,omitempty"`
	// Fib nhid encap id
	NhidFeid uint64 `protobuf:"varint,36,opt,name=nhid_feid,json=nhidFeid" json:"nhid_feid,omitempty"`
	// Fib mpls encap id
	MplsFeid uint64 `protobuf:"varint,37,opt,name=mpls_feid,json=mplsFeid" json:"mpls_feid,omitempty"`
}

func (m *Ipv4RibEdmPathItem) Reset()                    { *m = Ipv4RibEdmPathItem{} }
func (m *Ipv4RibEdmPathItem) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPathItem) ProtoMessage()               {}
func (*Ipv4RibEdmPathItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ipv4RibEdmPathItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInformationSource() string {
	if m != nil {
		return m.InformationSource
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetV6Nexthop() string {
	if m != nil {
		return m.V6Nexthop
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags64() uint64 {
	if m != nil {
		return m.Flags64
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPrivateFlags() uint32 {
	if m != nil {
		return m.PrivateFlags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLooped() bool {
	if m != nil {
		return m.Looped
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopVrfName() string {
	if m != nil {
		return m.NextHopVrfName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableName() string {
	if m != nil {
		return m.NextHopTableName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopAfi() uint32 {
	if m != nil {
		return m.NextHopAfi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopSafi() uint32 {
	if m != nil {
		return m.NextHopSafi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPathid() uint32 {
	if m != nil {
		return m.Pathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetBackupPathid() uint32 {
	if m != nil {
		return m.BackupPathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRefCntOfBackup() uint32 {
	if m != nil {
		return m.RefCntOfBackup
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMvpnPresent() bool {
	if m != nil {
		return m.MvpnPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetPathrtPresent() bool {
	if m != nil {
		return m.PathrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetVrfimportrtPresent() bool {
	if m != nil {
		return m.VrfimportrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourceasrtPresent() bool {
	if m != nil {
		return m.SourceasrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourcerdPresent() bool {
	if m != nil {
		return m.SourcerdPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSegmentedNexthopPresent() bool {
	if m != nil {
		return m.SegmentedNexthopPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopId() uint32 {
	if m != nil {
		return m.NextHopId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopIdRefcount() uint32 {
	if m != nil {
		return m.NextHopIdRefcount
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetOspfAreaId() string {
	if m != nil {
		return m.OspfAreaId
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetRemoteBackupAddr() [][]byte {
	if m != nil {
		return m.RemoteBackupAddr
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetHasLabelstk() bool {
	if m != nil {
		return m.HasLabelstk
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNumLabels() uint32 {
	if m != nil {
		return m.NumLabels
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLabelstk() []uint32 {
	if m != nil {
		return m.Labelstk
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetBindingLabel() uint32 {
	if m != nil {
		return m.BindingLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNhidFeid() uint64 {
	if m != nil {
		return m.NhidFeid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMplsFeid() uint64 {
	if m != nil {
		return m.MplsFeid
	}
	return 0
}

// Information of local label for route head
type RibEdmLocalLabel struct {
	// Protocol Name
	ProtocolName string `protobuf:"bytes,1,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Client ID
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Stale
	Stale uint32 `protobuf:"varint,3,opt,name=stale" json:"stale,omitempty"`
	// Mirrored
	Mirrored uint32 `protobuf:"varint,4,opt,name=mirrored" json:"mirrored,omitempty"`
	// Merge disable
	MergeDisable uint32 `protobuf:"varint,5,opt,name=merge_disable,json=mergeDisable" json:"merge_disable,omitempty"`
	// Redist only
	RedistOnly uint32 `protobuf:"varint,6,opt,name=redist_only,json=redistOnly" json:"redist_only,omitempty"`
	// Local label
	Label uint32 `protobuf:"varint,7,opt,name=label" json:"label,omitempty"`
	// Distance
	Distance uint32 `protobuf:"varint,8,opt,name=distance" json:"distance,omitempty"`
}

func (m *RibEdmLocalLabel) Reset()                    { *m = RibEdmLocalLabel{} }
func (m *RibEdmLocalLabel) String() string            { return proto.CompactTextString(m) }
func (*RibEdmLocalLabel) ProtoMessage()               {}
func (*RibEdmLocalLabel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RibEdmLocalLabel) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *RibEdmLocalLabel) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *RibEdmLocalLabel) GetStale() uint32 {
	if m != nil {
		return m.Stale
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMirrored() uint32 {
	if m != nil {
		return m.Mirrored
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMergeDisable() uint32 {
	if m != nil {
		return m.MergeDisable
	}
	return 0
}

func (m *RibEdmLocalLabel) GetRedistOnly() uint32 {
	if m != nil {
		return m.RedistOnly
	}
	return 0
}

func (m *RibEdmLocalLabel) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *RibEdmLocalLabel) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv4_rib_edm_route")
	proto.RegisterType((*Ipv4RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv4_rib_edm_path")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv4_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.rib_edm_local_label")
}

func init() { proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x5b, 0x6e, 0x5b, 0x37,
	0x10, 0x85, 0xfc, 0x94, 0x68, 0xcb, 0xb1, 0xae, 0x5d, 0x9b, 0x8e, 0xf3, 0x50, 0x9c, 0x06, 0x90,
	0x8b, 0x5a, 0x29, 0x92, 0xd4, 0x6d, 0xd3, 0xa7, 0xe3, 0x38, 0x89, 0x1a, 0x27, 0x76, 0x95, 0x20,
	0x40, 0xbf, 0x08, 0xea, 0x92, 0x57, 0x22, 0x72, 0x5f, 0x21, 0x29, 0xc5, 0xfe, 0x29, 0xd0, 0x25,
	0xf4, 0xa7, 0x0b, 0xe8, 0x0e, 0xba, 0x85, 0xae, 0xa0, 0x9f, 0x5d, 0x47, 0x57, 0x50, 0x70, 0x86,
	0xf7, 0x4a, 0x7e, 0xf4, 0xbb, 0xfd, 0x89, 0x32, 0xe7, 0x1c, 0xf2, 0x0e, 0x67, 0x86, 0x33, 0x34,
	0xa1, 0x2a, 0x1f, 0x3d, 0x60, 0x5a, 0xf5, 0x98, 0x14, 0x09, 0xd3, 0xd9, 0xd0, 0xca, 0x76, 0xae,
	0x33, 0x9b, 0x05, 0x3f, 0x85, 0xca, 0x84, 0x19, 0x53, 0x99, 0x61, 0x27, 0x9a, 0xa9, 0x1c, 0x44,
	0xa0, 0xce, 0x72, 0xa9, 0xdb, 0xce, 0x32, 0x56, 0xf4, 0x4e, 0xdb, 0x23, 0x1d, 0x19, 0xf7, 0x4f,
	0x9b, 0x47, 0xa6, 0xcd, 0xa3, 0xb6, 0x71, 0xbf, 0x86, 0x47, 0x6d, 0xbf, 0x06, 0x36, 0x65, 0x96,
	0xf7, 0x62, 0xc9, 0x52, 0x9e, 0x48, 0xf3, 0x6f, 0x44, 0x1b, 0x00, 0x83, 0x3f, 0x5b, 0xbf, 0x4d,
	0x91, 0xf5, 0x8b, 0xce, 0xb1, 0xe7, 0x07, 0x3f, 0xbe, 0x0a, 0x36, 0x48, 0x75, 0xa4, 0x23, 0x58,
	0x44, 0x2b, 0xcd, 0x4a, 0xab, 0xd6, 0x9d, 0x1f, 0xe9, 0xe8, 0x25, 0x4f, 0x64, 0xb0, 0x4e, 0xe6,
	0xb9, 0x67, 0xa6, 0x80, 0x99, 0xe3, 0x48, 0x6c, 0x90, 0xaa, 0x29, 0x98, 0x69, 0x5c, 0x63, 0x3c,
	0xd5, 0x22, 0xcb, 0xe7, 0x7d, 0xa1, 0x33, 0x20, 0x59, 0x02, 0xfc, 0xb5, 0x83, 0x41, 0x49, 0xc9,
	0x3c, 0x17, 0x42, 0x4b, 0x63, 0xe8, 0x2c, 0xee, 0xe1, 0xcd, 0xe0, 0x36, 0xa9, 0xe7, 0x5a, 0x46,
	0xea, 0x84, 0xc5, 0x32, 0xed, 0xdb, 0x01, 0x9d, 0x6b, 0x56, 0x5a, 0xf5, 0xee, 0x22, 0x82, 0x87,
	0x80, 0xb9, 0x0f, 0xa5, 0xf2, 0xc4, 0xb2, 0x41, 0x96, 0xb3, 0x62, 0x9f, 0x79, 0xfc, 0x90, 0xc3,
	0x9f, 0x65, 0xf9, 0x9e, 0xdf, 0xee, 0x0e, 0x59, 0x52, 0xa9, 0x95, 0x3a, 0xe2, 0xa1, 0x77, 0xa8,
	0x0a, 0xba, 0x7a, 0x89, 0x3a, 0x7f, 0xb6, 0xfe, 0xac, 0x91, 0xe0, 0x62, 0x90, 0x82, 0x35, 0x32,
	0x87, 0xdf, 0xa5, 0xf7, 0x30, 0x06, 0x68, 0x5d, 0x74, 0xf2, 0xfe, 0x25, 0x4e, 0xde, 0x26, 0x75,
	0x8c, 0xc6, 0x48, 0x6a, 0xa3, 0xb2, 0x94, 0x3e, 0x40, 0x11, 0x80, 0x6f, 0x10, 0x0b, 0x6e, 0x92,
	0x05, 0x28, 0x93, 0x30, 0x8b, 0x99, 0x12, 0xf4, 0x53, 0x90, 0x90, 0x02, 0xea, 0x08, 0xfc, 0x94,
	0x17, 0x80, 0xff, 0xbb, 0xe0, 0xc9, 0x62, 0x01, 0x42, 0x38, 0xaf, 0x92, 0xaa, 0x4a, 0x8d, 0xe5,
	0x69, 0x28, 0xe9, 0x67, 0xc0, 0x97, 0x76, 0xb0, 0x49, 0x6a, 0x61, 0xac, 0x64, 0x6a, 0xdd, 0xfe,
	0x9f, 0xc3, 0xfe, 0x55, 0x04, 0x3a, 0x22, 0xb8, 0x4e, 0x88, 0xcf, 0xd8, 0x69, 0x2e, 0xe9, 0x17,
	0xc0, 0xd6, 0x30, 0x57, 0xa7, 0x39, 0xec, 0x9b, 0x6b, 0x95, 0x69, 0x65, 0x4f, 0xe9, 0x43, 0x5c,
	0x5a, 0xd8, 0x50, 0x07, 0x23, 0x81, 0x0b, 0xbf, 0x04, 0x6e, 0xde, 0x8c, 0x04, 0x2c, 0x5b, 0x25,
	0xb3, 0x51, 0xcc, 0xfb, 0x86, 0x7e, 0x05, 0x38, 0x1a, 0x2e, 0x15, 0xf2, 0xc4, 0xca, 0x54, 0x48,
	0xc1, 0x90, 0xfe, 0xba, 0x59, 0x69, 0xcd, 0x74, 0xeb, 0x05, 0xfa, 0x04, 0x64, 0xcb, 0x64, 0xda,
	0xf2, 0x3e, 0xfd, 0x06, 0x96, 0xba, 0xff, 0x3a, 0x2f, 0x84, 0xf2, 0xa7, 0xfb, 0x16, 0xbd, 0x28,
	0xec, 0x60, 0x87, 0x04, 0x42, 0xf9, 0x00, 0xb3, 0x52, 0xf5, 0x1d, 0xa8, 0x1a, 0x25, 0xf3, 0xb8,
	0x90, 0xaf, 0x91, 0xb9, 0x44, 0x5a, 0xad, 0x42, 0xba, 0x07, 0x12, 0x6f, 0x41, 0x1a, 0xb8, 0x1d,
	0x18, 0x16, 0x66, 0xc3, 0xd4, 0xd2, 0x47, 0x3e, 0x0d, 0x0e, 0xda, 0x77, 0x88, 0xfb, 0x0e, 0xb7,
	0x56, 0xab, 0x9e, 0x0b, 0x96, 0x12, 0x32, 0xb5, 0x2e, 0x26, 0xfb, 0xf8, 0x9d, 0x92, 0xe9, 0x78,
	0xc2, 0x65, 0xcd, 0x6a, 0x1e, 0x45, 0x2a, 0x64, 0x2a, 0x15, 0xf2, 0x84, 0x3e, 0xc6, 0xdc, 0x7b,
	0xb0, 0xe3, 0xb0, 0x60, 0xbb, 0xb8, 0x2e, 0xb9, 0x96, 0xa1, 0x14, 0xd2, 0x79, 0x7e, 0x00, 0xba,
	0x2b, 0x80, 0x1f, 0x97, 0xb0, 0x4b, 0xe2, 0xbb, 0xcc, 0xb0, 0xbe, 0xce, 0x86, 0x39, 0x7d, 0x82,
	0x31, 0x78, 0x97, 0x99, 0xa7, 0xce, 0x76, 0x99, 0x88, 0xe2, 0xec, 0x3d, 0x73, 0x61, 0x7b, 0x8a,
	0x99, 0x70, 0xf6, 0x6b, 0xde, 0x77, 0xeb, 0xa2, 0xf7, 0x82, 0x85, 0x31, 0x37, 0x86, 0x3e, 0xc3,
	0x75, 0xd1, 0x7b, 0xb1, 0xef, 0x6c, 0x47, 0xe6, 0x2a, 0xf4, 0x47, 0xee, 0xf8, 0xf4, 0xaa, 0x10,
	0x0f, 0xbc, 0x46, 0xe6, 0x78, 0x68, 0xd5, 0x48, 0xd2, 0xef, 0x9b, 0x95, 0x56, 0xb5, 0xeb, 0xad,
	0xe0, 0x1a, 0xa9, 0x95, 0x61, 0xa5, 0xcf, 0x81, 0x1a, 0x03, 0xc1, 0x27, 0x64, 0x75, 0x9c, 0x0e,
	0x28, 0x51, 0x2c, 0xda, 0x43, 0x28, 0xca, 0x71, 0xaa, 0x8e, 0x1d, 0x05, 0xa5, 0xbb, 0x49, 0xb0,
	0xde, 0x18, 0xef, 0x4b, 0xfa, 0x02, 0x9d, 0x00, 0x60, 0xaf, 0x2f, 0x5d, 0x5a, 0x90, 0x8c, 0x79,
	0x4f, 0xc6, 0xf4, 0x25, 0xa6, 0x05, 0xa0, 0x43, 0x87, 0xb8, 0x3e, 0x52, 0xf8, 0x72, 0x84, 0x27,
	0x1f, 0x8d, 0x2f, 0x96, 0xed, 0xc5, 0xe5, 0xdd, 0x3b, 0x86, 0x52, 0x23, 0xb6, 0x17, 0x17, 0x37,
	0xef, 0x23, 0xd2, 0xc0, 0xbd, 0x93, 0x4c, 0xa8, 0xe8, 0x94, 0x59, 0x95, 0x48, 0xfa, 0x03, 0xc8,
	0x30, 0xfc, 0x2f, 0x00, 0x7f, 0xad, 0x12, 0x19, 0xfc, 0x5e, 0x29, 0xee, 0x89, 0x2b, 0x09, 0xda,
	0x6d, 0x56, 0x5a, 0x0b, 0xf7, 0x7e, 0xa9, 0xb4, 0xff, 0xdb, 0xd6, 0xde, 0x3e, 0xd3, 0xb1, 0x9c,
	0x67, 0xfe, 0xee, 0x1e, 0x73, 0x3b, 0xd8, 0xfa, 0xab, 0x42, 0x1a, 0x17, 0x04, 0xc1, 0x1f, 0x97,
	0xa1, 0xb4, 0xd2, 0x9c, 0x6e, 0x2d, 0xdc, 0xfb, 0xf5, 0xff, 0x77, 0x20, 0xa6, 0xac, 0x4c, 0xba,
	0x4b, 0x0e, 0xef, 0xaa, 0xde, 0x81, 0x48, 0xe0, 0x68, 0x7f, 0x13, 0xb2, 0x76, 0xb9, 0x74, 0x72,
	0xb0, 0x54, 0xce, 0x0e, 0x96, 0x1d, 0x12, 0xa8, 0x34, 0xca, 0x74, 0xc2, 0xad, 0x2b, 0x4e, 0x93,
	0x0d, 0x75, 0x58, 0xcc, 0xb6, 0xc6, 0x04, 0xf3, 0x0a, 0x08, 0xd7, 0x19, 0x47, 0xbb, 0xcc, 0x4d,
	0x93, 0x41, 0x96, 0xfb, 0x41, 0x57, 0x1b, 0xed, 0xbe, 0x44, 0xe0, 0x92, 0xb9, 0x32, 0x73, 0xc9,
	0x5c, 0x99, 0xe8, 0x37, 0xb3, 0xe7, 0xfb, 0x4d, 0x9c, 0x71, 0xc1, 0x3c, 0x89, 0x33, 0x8e, 0x38,
	0xe8, 0x05, 0x0a, 0x28, 0x99, 0x87, 0x1e, 0xb9, 0xfb, 0x00, 0x06, 0xdb, 0x4c, 0xb7, 0x30, 0xc7,
	0xcd, 0xb5, 0x3a, 0xd9, 0x5c, 0x61, 0x4c, 0xa8, 0x11, 0xb7, 0xd2, 0xf7, 0xd6, 0x5a, 0x31, 0x91,
	0x00, 0xc4, 0xd6, 0xba, 0x46, 0xe6, 0xe2, 0x2c, 0xcb, 0xa5, 0xa0, 0x04, 0xef, 0x34, 0x5a, 0xc1,
	0x36, 0x69, 0x94, 0xe3, 0x14, 0x53, 0xa3, 0x04, 0x5d, 0x80, 0x0d, 0x8a, 0x79, 0x0a, 0xa3, 0xbb,
	0x73, 0x56, 0x5a, 0x3e, 0x1d, 0x16, 0xcf, 0x8c, 0xde, 0x37, 0xfe, 0x05, 0xb1, 0x43, 0x56, 0xce,
	0xed, 0x0a, 0xe2, 0x3a, 0x88, 0x97, 0x27, 0xf7, 0x05, 0x79, 0x93, 0x2c, 0x8e, 0x67, 0x7a, 0xa4,
	0xe8, 0x12, 0xc6, 0xa4, 0x98, 0xe7, 0x91, 0x0a, 0xb6, 0x48, 0xbd, 0x54, 0x18, 0x27, 0xb9, 0x02,
	0x92, 0x05, 0x2f, 0x79, 0xc5, 0x23, 0x75, 0xbe, 0x63, 0x2c, 0x5f, 0xe8, 0x18, 0x9b, 0xa4, 0x66,
	0x87, 0x69, 0x2a, 0x61, 0xdc, 0x36, 0xb0, 0xdf, 0x20, 0xd0, 0x11, 0x30, 0xef, 0xb9, 0x1d, 0x28,
	0x41, 0x03, 0x4c, 0x17, 0x5a, 0x2e, 0xba, 0x3d, 0x1e, 0xbe, 0x1d, 0xe6, 0xcc, 0xd3, 0x2b, 0x18,
	0x5d, 0x04, 0x8f, 0x51, 0xb4, 0x4d, 0x1a, 0x5a, 0x46, 0x2c, 0x4c, 0x2d, 0xcb, 0x22, 0x86, 0x14,
	0x5d, 0xc5, 0x28, 0x6a, 0x19, 0xed, 0xa7, 0xf6, 0x28, 0x7a, 0x04, 0x68, 0xb0, 0x4f, 0x6e, 0xa4,
	0xc3, 0xa4, 0x27, 0xb5, 0x53, 0x96, 0x43, 0x31, 0xcc, 0x92, 0x64, 0x98, 0x2a, 0xab, 0xa4, 0xa1,
	0x1f, 0xc0, 0xba, 0x4d, 0x54, 0x1d, 0x45, 0x07, 0x5e, 0xb3, 0x3f, 0x96, 0x04, 0xb7, 0xc8, 0x62,
	0x32, 0xca, 0x5d, 0x9b, 0x95, 0x46, 0xa6, 0x96, 0xae, 0x41, 0x4e, 0x17, 0x1c, 0x76, 0x8c, 0x90,
	0xab, 0x52, 0xe7, 0xb0, 0xb6, 0xa5, 0x68, 0x1d, 0x44, 0x75, 0x44, 0x0b, 0xd9, 0x5d, 0xb2, 0x32,
	0xd2, 0x91, 0x4a, 0xf2, 0x4c, 0xdb, 0x09, 0x2d, 0x05, 0x6d, 0x30, 0x41, 0x15, 0x0b, 0x76, 0x48,
	0x80, 0xf7, 0x87, 0x9b, 0x09, 0xfd, 0x06, 0xe8, 0x1b, 0x63, 0xa6, 0x90, 0x6f, 0x93, 0x65, 0x04,
	0xb5, 0x28, 0xc5, 0x57, 0x41, 0x7c, 0xa5, 0xc0, 0x0b, 0xe9, 0x43, 0xb2, 0x61, 0x64, 0x3f, 0x91,
	0xa9, 0x95, 0xa2, 0xb8, 0x7d, 0xe5, 0x9a, 0x4d, 0x58, 0xb3, 0x5e, 0x0a, 0xfc, 0x65, 0x2c, 0xd6,
	0xde, 0x20, 0x0b, 0x65, 0x7d, 0x28, 0x41, 0xaf, 0xe1, 0x6b, 0xc6, 0x57, 0x47, 0x47, 0x04, 0x77,
	0xc9, 0xea, 0x04, 0xcf, 0xb4, 0x8c, 0x70, 0xf4, 0x5d, 0xc7, 0x29, 0x5e, 0x0a, 0xbb, 0x9e, 0x70,
	0x25, 0x99, 0x99, 0x3c, 0x62, 0x5c, 0x4b, 0xee, 0x76, 0xbc, 0x01, 0xa5, 0x4b, 0x1c, 0xb6, 0xa7,
	0x25, 0xef, 0x88, 0xe0, 0x63, 0x12, 0x68, 0x99, 0x64, 0x56, 0xfa, 0x7c, 0xc3, 0x6b, 0x94, 0xde,
	0x6c, 0x4e, 0xb7, 0x16, 0xbb, 0xcb, 0xc8, 0x60, 0xca, 0xdd, 0x7b, 0xd4, 0x65, 0x6c, 0xc0, 0x0d,
	0x96, 0xa6, 0xb1, 0x6f, 0x69, 0x13, 0x33, 0x36, 0xe0, 0xe6, 0xd0, 0x43, 0xae, 0xed, 0xa4, 0xc3,
	0xc4, 0x4b, 0xe8, 0x2d, 0x7f, 0x84, 0x61, 0x82, 0x02, 0xf7, 0x14, 0x2a, 0x57, 0x6f, 0x35, 0xa7,
	0x5d, 0xf1, 0x16, 0x36, 0x14, 0xa9, 0x4a, 0x85, 0x4a, 0xfb, 0xbe, 0xf8, 0x6f, 0xfb, 0x22, 0x45,
	0xb0, 0x2c, 0xff, 0x74, 0xa0, 0x04, 0x8b, 0xa4, 0x12, 0xf4, 0x43, 0xe8, 0x2c, 0x55, 0x07, 0x3c,
	0x91, 0x4a, 0x38, 0x32, 0xc9, 0x63, 0x83, 0xe4, 0x1d, 0x24, 0x1d, 0xe0, 0xc8, 0xad, 0x9f, 0xa7,
	0xc8, 0x4a, 0xd1, 0x6f, 0xe3, 0x2c, 0xe4, 0x31, 0x7e, 0xe5, 0xe2, 0x03, 0xb5, 0x72, 0xc9, 0x03,
	0xf5, 0xcc, 0x23, 0x74, 0xea, 0xdc, 0x23, 0x74, 0x95, 0xcc, 0x1a, 0xcb, 0x63, 0xfc, 0x73, 0xa2,
	0xde, 0x45, 0xc3, 0x1d, 0x35, 0x51, 0x5a, 0x67, 0x5a, 0x0a, 0xe8, 0xad, 0xf5, 0x6e, 0x69, 0xbb,
	0x6f, 0x26, 0x52, 0xf7, 0xa5, 0x7b, 0xf1, 0xb9, 0x06, 0xe2, 0xbb, 0xeb, 0x22, 0x80, 0x8f, 0x11,
	0x83, 0x56, 0x20, 0xdd, 0x93, 0x90, 0x65, 0x69, 0x7c, 0x5a, 0xf4, 0x58, 0x84, 0x8e, 0xd2, 0xf8,
	0xd4, 0x7d, 0x17, 0x03, 0x35, 0x8f, 0xdf, 0xc5, 0xf3, 0x4c, 0xbe, 0x36, 0xab, 0x67, 0x5f, 0x9b,
	0xbd, 0x39, 0x38, 0xd4, 0xfd, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x52, 0x26, 0xc6, 0xaf, 0xee,
	0x0d, 0x00, 0x00,
}