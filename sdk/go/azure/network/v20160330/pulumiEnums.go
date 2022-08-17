


package v20160330

type ApplicationGatewayCookieBasedAffinity string

const (
	ApplicationGatewayCookieBasedAffinityEnabled  = ApplicationGatewayCookieBasedAffinity("Enabled")
	ApplicationGatewayCookieBasedAffinityDisabled = ApplicationGatewayCookieBasedAffinity("Disabled")
)

type ApplicationGatewayProtocol string

const (
	ApplicationGatewayProtocolHttp  = ApplicationGatewayProtocol("Http")
	ApplicationGatewayProtocolHttps = ApplicationGatewayProtocol("Https")
)

type ApplicationGatewayRequestRoutingRuleType string

const (
	ApplicationGatewayRequestRoutingRuleTypeBasic            = ApplicationGatewayRequestRoutingRuleType("Basic")
	ApplicationGatewayRequestRoutingRuleTypePathBasedRouting = ApplicationGatewayRequestRoutingRuleType("PathBasedRouting")
)

type ApplicationGatewaySkuName string

const (
	ApplicationGatewaySkuName_Standard_Small  = ApplicationGatewaySkuName("Standard_Small")
	ApplicationGatewaySkuName_Standard_Medium = ApplicationGatewaySkuName("Standard_Medium")
	ApplicationGatewaySkuName_Standard_Large  = ApplicationGatewaySkuName("Standard_Large")
)

type ApplicationGatewayTier string

const (
	ApplicationGatewayTierStandard = ApplicationGatewayTier("Standard")
)

type AuthorizationUseStatus string

const (
	AuthorizationUseStatusAvailable = AuthorizationUseStatus("Available")
	AuthorizationUseStatusInUse     = AuthorizationUseStatus("InUse")
)

type ExpressRouteCircuitPeeringAdvertisedPublicPrefixState string

const (
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateNotConfigured    = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("NotConfigured")
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateConfiguring      = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("Configuring")
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateConfigured       = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("Configured")
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateValidationNeeded = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("ValidationNeeded")
)

type ExpressRouteCircuitPeeringStateEnum string

const (
	ExpressRouteCircuitPeeringStateEnumDisabled = ExpressRouteCircuitPeeringStateEnum("Disabled")
	ExpressRouteCircuitPeeringStateEnumEnabled  = ExpressRouteCircuitPeeringStateEnum("Enabled")
)

type ExpressRouteCircuitPeeringTypeEnum string

const (
	ExpressRouteCircuitPeeringTypeEnumAzurePublicPeering  = ExpressRouteCircuitPeeringTypeEnum("AzurePublicPeering")
	ExpressRouteCircuitPeeringTypeEnumAzurePrivatePeering = ExpressRouteCircuitPeeringTypeEnum("AzurePrivatePeering")
	ExpressRouteCircuitPeeringTypeEnumMicrosoftPeering    = ExpressRouteCircuitPeeringTypeEnum("MicrosoftPeering")
)

type ExpressRouteCircuitSkuFamily string

const (
	ExpressRouteCircuitSkuFamilyUnlimitedData = ExpressRouteCircuitSkuFamily("UnlimitedData")
	ExpressRouteCircuitSkuFamilyMeteredData   = ExpressRouteCircuitSkuFamily("MeteredData")
)

type ExpressRouteCircuitSkuTier string

const (
	ExpressRouteCircuitSkuTierStandard = ExpressRouteCircuitSkuTier("Standard")
	ExpressRouteCircuitSkuTierPremium  = ExpressRouteCircuitSkuTier("Premium")
)

type IPAllocationMethod string

const (
	IPAllocationMethodStatic  = IPAllocationMethod("Static")
	IPAllocationMethodDynamic = IPAllocationMethod("Dynamic")
)

type IPVersion string

const (
	IPVersionIPv4 = IPVersion("IPv4")
	IPVersionIPv6 = IPVersion("IPv6")
)

type LoadDistribution string

const (
	LoadDistributionDefault          = LoadDistribution("Default")
	LoadDistributionSourceIP         = LoadDistribution("SourceIP")
	LoadDistributionSourceIPProtocol = LoadDistribution("SourceIPProtocol")
)

type ProbeProtocol string

const (
	ProbeProtocolHttp = ProbeProtocol("Http")
	ProbeProtocolTcp  = ProbeProtocol("Tcp")
)

type RouteNextHopType string

const (
	RouteNextHopTypeVirtualNetworkGateway = RouteNextHopType("VirtualNetworkGateway")
	RouteNextHopTypeVnetLocal             = RouteNextHopType("VnetLocal")
	RouteNextHopTypeInternet              = RouteNextHopType("Internet")
	RouteNextHopTypeVirtualAppliance      = RouteNextHopType("VirtualAppliance")
	RouteNextHopTypeNone                  = RouteNextHopType("None")
)

type SecurityRuleAccess string

const (
	SecurityRuleAccessAllow = SecurityRuleAccess("Allow")
	SecurityRuleAccessDeny  = SecurityRuleAccess("Deny")
)

type SecurityRuleDirection string

const (
	SecurityRuleDirectionInbound  = SecurityRuleDirection("Inbound")
	SecurityRuleDirectionOutbound = SecurityRuleDirection("Outbound")
)

type SecurityRuleProtocol string

const (
	SecurityRuleProtocolTcp      = SecurityRuleProtocol("Tcp")
	SecurityRuleProtocolUdp      = SecurityRuleProtocol("Udp")
	SecurityRuleProtocolAsterisk = SecurityRuleProtocol("*")
)

type ServiceProviderProvisioningState string

const (
	ServiceProviderProvisioningStateNotProvisioned = ServiceProviderProvisioningState("NotProvisioned")
	ServiceProviderProvisioningStateProvisioning   = ServiceProviderProvisioningState("Provisioning")
	ServiceProviderProvisioningStateProvisioned    = ServiceProviderProvisioningState("Provisioned")
	ServiceProviderProvisioningStateDeprovisioning = ServiceProviderProvisioningState("Deprovisioning")
)

type TransportProtocol string

const (
	TransportProtocolUdp = TransportProtocol("Udp")
	TransportProtocolTcp = TransportProtocol("Tcp")
)

type VirtualNetworkGatewayConnectionStatus string

const (
	VirtualNetworkGatewayConnectionStatusUnknown      = VirtualNetworkGatewayConnectionStatus("Unknown")
	VirtualNetworkGatewayConnectionStatusConnecting   = VirtualNetworkGatewayConnectionStatus("Connecting")
	VirtualNetworkGatewayConnectionStatusConnected    = VirtualNetworkGatewayConnectionStatus("Connected")
	VirtualNetworkGatewayConnectionStatusNotConnected = VirtualNetworkGatewayConnectionStatus("NotConnected")
)

type VirtualNetworkGatewayConnectionType string

const (
	VirtualNetworkGatewayConnectionTypeIPsec        = VirtualNetworkGatewayConnectionType("IPsec")
	VirtualNetworkGatewayConnectionTypeVnet2Vnet    = VirtualNetworkGatewayConnectionType("Vnet2Vnet")
	VirtualNetworkGatewayConnectionTypeExpressRoute = VirtualNetworkGatewayConnectionType("ExpressRoute")
	VirtualNetworkGatewayConnectionTypeVPNClient    = VirtualNetworkGatewayConnectionType("VPNClient")
)

type VirtualNetworkGatewaySkuName string

const (
	VirtualNetworkGatewaySkuNameBasic           = VirtualNetworkGatewaySkuName("Basic")
	VirtualNetworkGatewaySkuNameHighPerformance = VirtualNetworkGatewaySkuName("HighPerformance")
	VirtualNetworkGatewaySkuNameStandard        = VirtualNetworkGatewaySkuName("Standard")
)

type VirtualNetworkGatewaySkuTier string

const (
	VirtualNetworkGatewaySkuTierBasic           = VirtualNetworkGatewaySkuTier("Basic")
	VirtualNetworkGatewaySkuTierHighPerformance = VirtualNetworkGatewaySkuTier("HighPerformance")
	VirtualNetworkGatewaySkuTierStandard        = VirtualNetworkGatewaySkuTier("Standard")
)

type VirtualNetworkGatewayTypeEnum string

const (
	VirtualNetworkGatewayTypeEnumVpn          = VirtualNetworkGatewayTypeEnum("Vpn")
	VirtualNetworkGatewayTypeEnumExpressRoute = VirtualNetworkGatewayTypeEnum("ExpressRoute")
)

type VpnType string

const (
	VpnTypePolicyBased = VpnType("PolicyBased")
	VpnTypeRouteBased  = VpnType("RouteBased")
)

func init() {
}
