


package v20220401preview

type AddressPrefixType string

const (
	AddressPrefixTypeIPPrefix   = AddressPrefixType("IPPrefix")
	AddressPrefixTypeServiceTag = AddressPrefixType("ServiceTag")
)

type AdminRuleKind string

const (
	AdminRuleKindCustom  = AdminRuleKind("Custom")
	AdminRuleKindDefault = AdminRuleKind("Default")
)

type AllowedEndpointRecordType string

const (
	AllowedEndpointRecordTypeDomainName  = AllowedEndpointRecordType("DomainName")
	AllowedEndpointRecordTypeIPv4Address = AllowedEndpointRecordType("IPv4Address")
	AllowedEndpointRecordTypeIPv6Address = AllowedEndpointRecordType("IPv6Address")
	AllowedEndpointRecordTypeAny         = AllowedEndpointRecordType("Any")
)

type AlwaysServe string

const (
	AlwaysServeEnabled  = AlwaysServe("Enabled")
	AlwaysServeDisabled = AlwaysServe("Disabled")
)

type ConfigurationType string

const (
	ConfigurationTypeSecurityAdmin = ConfigurationType("SecurityAdmin")
	ConfigurationTypeSecurityUser  = ConfigurationType("SecurityUser")
	ConfigurationTypeConnectivity  = ConfigurationType("Connectivity")
)

type ConnectivityTopology string

const (
	ConnectivityTopologyHubAndSpoke = ConnectivityTopology("HubAndSpoke")
	ConnectivityTopologyMesh        = ConnectivityTopology("Mesh")
)

type DeleteExistingNSGs string

const (
	DeleteExistingNSGsFalse = DeleteExistingNSGs("False")
	DeleteExistingNSGsTrue  = DeleteExistingNSGs("True")
)

type DeleteExistingPeering string

const (
	DeleteExistingPeeringFalse = DeleteExistingPeering("False")
	DeleteExistingPeeringTrue  = DeleteExistingPeering("True")
)

type EndpointMonitorStatus string

const (
	EndpointMonitorStatusCheckingEndpoint = EndpointMonitorStatus("CheckingEndpoint")
	EndpointMonitorStatusOnline           = EndpointMonitorStatus("Online")
	EndpointMonitorStatusDegraded         = EndpointMonitorStatus("Degraded")
	EndpointMonitorStatusDisabled         = EndpointMonitorStatus("Disabled")
	EndpointMonitorStatusInactive         = EndpointMonitorStatus("Inactive")
	EndpointMonitorStatusStopped          = EndpointMonitorStatus("Stopped")
)

type EndpointStatus string

const (
	EndpointStatusEnabled  = EndpointStatus("Enabled")
	EndpointStatusDisabled = EndpointStatus("Disabled")
)

type GroupConnectivity string

const (
	GroupConnectivityNone              = GroupConnectivity("None")
	GroupConnectivityDirectlyConnected = GroupConnectivity("DirectlyConnected")
)

type IsGlobal string

const (
	IsGlobalFalse = IsGlobal("False")
	IsGlobalTrue  = IsGlobal("True")
)

type MonitorProtocol string

const (
	MonitorProtocolHTTP  = MonitorProtocol("HTTP")
	MonitorProtocolHTTPS = MonitorProtocol("HTTPS")
	MonitorProtocolTCP   = MonitorProtocol("TCP")
)

type NetworkIntentPolicyBasedService string

const (
	NetworkIntentPolicyBasedServiceNone = NetworkIntentPolicyBasedService("None")
	NetworkIntentPolicyBasedServiceAll  = NetworkIntentPolicyBasedService("All")
)

type ProfileMonitorStatus string

const (
	ProfileMonitorStatusCheckingEndpoints = ProfileMonitorStatus("CheckingEndpoints")
	ProfileMonitorStatusOnline            = ProfileMonitorStatus("Online")
	ProfileMonitorStatusDegraded          = ProfileMonitorStatus("Degraded")
	ProfileMonitorStatusDisabled          = ProfileMonitorStatus("Disabled")
	ProfileMonitorStatusInactive          = ProfileMonitorStatus("Inactive")
)

type ProfileStatus string

const (
	ProfileStatusEnabled  = ProfileStatus("Enabled")
	ProfileStatusDisabled = ProfileStatus("Disabled")
)

type SecurityConfigurationRuleAccess string

const (
	SecurityConfigurationRuleAccessAllow       = SecurityConfigurationRuleAccess("Allow")
	SecurityConfigurationRuleAccessDeny        = SecurityConfigurationRuleAccess("Deny")
	SecurityConfigurationRuleAccessAlwaysAllow = SecurityConfigurationRuleAccess("AlwaysAllow")
)

type SecurityConfigurationRuleDirection string

const (
	SecurityConfigurationRuleDirectionInbound  = SecurityConfigurationRuleDirection("Inbound")
	SecurityConfigurationRuleDirectionOutbound = SecurityConfigurationRuleDirection("Outbound")
)

type SecurityConfigurationRuleProtocol string

const (
	SecurityConfigurationRuleProtocolTcp  = SecurityConfigurationRuleProtocol("Tcp")
	SecurityConfigurationRuleProtocolUdp  = SecurityConfigurationRuleProtocol("Udp")
	SecurityConfigurationRuleProtocolIcmp = SecurityConfigurationRuleProtocol("Icmp")
	SecurityConfigurationRuleProtocolEsp  = SecurityConfigurationRuleProtocol("Esp")
	SecurityConfigurationRuleProtocolAny  = SecurityConfigurationRuleProtocol("Any")
	SecurityConfigurationRuleProtocolAh   = SecurityConfigurationRuleProtocol("Ah")
)

type TrafficRoutingMethod string

const (
	TrafficRoutingMethodPerformance = TrafficRoutingMethod("Performance")
	TrafficRoutingMethodPriority    = TrafficRoutingMethod("Priority")
	TrafficRoutingMethodWeighted    = TrafficRoutingMethod("Weighted")
	TrafficRoutingMethodGeographic  = TrafficRoutingMethod("Geographic")
	TrafficRoutingMethodMultiValue  = TrafficRoutingMethod("MultiValue")
	TrafficRoutingMethodSubnet      = TrafficRoutingMethod("Subnet")
)

type TrafficViewEnrollmentStatus string

const (
	TrafficViewEnrollmentStatusEnabled  = TrafficViewEnrollmentStatus("Enabled")
	TrafficViewEnrollmentStatusDisabled = TrafficViewEnrollmentStatus("Disabled")
)

type UseHubGateway string

const (
	UseHubGatewayFalse = UseHubGateway("False")
	UseHubGatewayTrue  = UseHubGateway("True")
)

type UserRuleKind string

const (
	UserRuleKindCustom  = UserRuleKind("Custom")
	UserRuleKindDefault = UserRuleKind("Default")
)

func init() {
}
