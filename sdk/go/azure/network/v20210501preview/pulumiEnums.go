


package v20210501preview

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

type NetworkIntentPolicyBasedService string

const (
	NetworkIntentPolicyBasedServiceNone = NetworkIntentPolicyBasedService("None")
	NetworkIntentPolicyBasedServiceAll  = NetworkIntentPolicyBasedService("All")
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

type SecurityType string

const (
	SecurityTypeAdminPolicy = SecurityType("AdminPolicy")
	SecurityTypeUserPolicy  = SecurityType("UserPolicy")
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
