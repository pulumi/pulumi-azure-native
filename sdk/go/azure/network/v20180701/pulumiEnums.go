


package v20180701

type Access string

const (
	AccessAllow = Access("Allow")
	AccessDeny  = Access("Deny")
)

type ApplicationGatewayCookieBasedAffinity string

const (
	ApplicationGatewayCookieBasedAffinityEnabled  = ApplicationGatewayCookieBasedAffinity("Enabled")
	ApplicationGatewayCookieBasedAffinityDisabled = ApplicationGatewayCookieBasedAffinity("Disabled")
)

type ApplicationGatewayFirewallMode string

const (
	ApplicationGatewayFirewallModeDetection  = ApplicationGatewayFirewallMode("Detection")
	ApplicationGatewayFirewallModePrevention = ApplicationGatewayFirewallMode("Prevention")
)

type ApplicationGatewayProtocol string

const (
	ApplicationGatewayProtocolHttp  = ApplicationGatewayProtocol("Http")
	ApplicationGatewayProtocolHttps = ApplicationGatewayProtocol("Https")
)

type ApplicationGatewayRedirectType string

const (
	ApplicationGatewayRedirectTypePermanent = ApplicationGatewayRedirectType("Permanent")
	ApplicationGatewayRedirectTypeFound     = ApplicationGatewayRedirectType("Found")
	ApplicationGatewayRedirectTypeSeeOther  = ApplicationGatewayRedirectType("SeeOther")
	ApplicationGatewayRedirectTypeTemporary = ApplicationGatewayRedirectType("Temporary")
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
	ApplicationGatewaySkuName_WAF_Medium      = ApplicationGatewaySkuName("WAF_Medium")
	ApplicationGatewaySkuName_WAF_Large       = ApplicationGatewaySkuName("WAF_Large")
	ApplicationGatewaySkuName_Standard_v2     = ApplicationGatewaySkuName("Standard_v2")
	ApplicationGatewaySkuName_WAF_v2          = ApplicationGatewaySkuName("WAF_v2")
)

type ApplicationGatewaySslCipherSuite string

const (
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384   = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256   = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA      = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA      = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_DHE_RSA_WITH_AES_256_GCM_SHA384     = ApplicationGatewaySslCipherSuite("TLS_DHE_RSA_WITH_AES_256_GCM_SHA384")
	ApplicationGatewaySslCipherSuite_TLS_DHE_RSA_WITH_AES_128_GCM_SHA256     = ApplicationGatewaySslCipherSuite("TLS_DHE_RSA_WITH_AES_128_GCM_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_DHE_RSA_WITH_AES_256_CBC_SHA        = ApplicationGatewaySslCipherSuite("TLS_DHE_RSA_WITH_AES_256_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_DHE_RSA_WITH_AES_128_CBC_SHA        = ApplicationGatewaySslCipherSuite("TLS_DHE_RSA_WITH_AES_128_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_256_GCM_SHA384         = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_256_GCM_SHA384")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_128_GCM_SHA256         = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_128_GCM_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_256_CBC_SHA256         = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_256_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_128_CBC_SHA256         = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_128_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_256_CBC_SHA            = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_256_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_128_CBC_SHA            = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_128_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384 = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256 = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA    = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA    = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_DHE_DSS_WITH_AES_256_CBC_SHA256     = ApplicationGatewaySslCipherSuite("TLS_DHE_DSS_WITH_AES_256_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_DHE_DSS_WITH_AES_128_CBC_SHA256     = ApplicationGatewaySslCipherSuite("TLS_DHE_DSS_WITH_AES_128_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_DHE_DSS_WITH_AES_256_CBC_SHA        = ApplicationGatewaySslCipherSuite("TLS_DHE_DSS_WITH_AES_256_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_DHE_DSS_WITH_AES_128_CBC_SHA        = ApplicationGatewaySslCipherSuite("TLS_DHE_DSS_WITH_AES_128_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_3DES_EDE_CBC_SHA           = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_3DES_EDE_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA       = ApplicationGatewaySslCipherSuite("TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256   = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384   = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384")
)

type ApplicationGatewaySslPolicyName string

const (
	ApplicationGatewaySslPolicyNameAppGwSslPolicy20150501  = ApplicationGatewaySslPolicyName("AppGwSslPolicy20150501")
	ApplicationGatewaySslPolicyNameAppGwSslPolicy20170401  = ApplicationGatewaySslPolicyName("AppGwSslPolicy20170401")
	ApplicationGatewaySslPolicyNameAppGwSslPolicy20170401S = ApplicationGatewaySslPolicyName("AppGwSslPolicy20170401S")
)

type ApplicationGatewaySslPolicyType string

const (
	ApplicationGatewaySslPolicyTypePredefined = ApplicationGatewaySslPolicyType("Predefined")
	ApplicationGatewaySslPolicyTypeCustom     = ApplicationGatewaySslPolicyType("Custom")
)

type ApplicationGatewaySslProtocol string

const (
	ApplicationGatewaySslProtocol_TLSv1_0 = ApplicationGatewaySslProtocol("TLSv1_0")
	ApplicationGatewaySslProtocol_TLSv1_1 = ApplicationGatewaySslProtocol("TLSv1_1")
	ApplicationGatewaySslProtocol_TLSv1_2 = ApplicationGatewaySslProtocol("TLSv1_2")
)

type ApplicationGatewayTier string

const (
	ApplicationGatewayTierStandard     = ApplicationGatewayTier("Standard")
	ApplicationGatewayTierWAF          = ApplicationGatewayTier("WAF")
	ApplicationGatewayTier_Standard_v2 = ApplicationGatewayTier("Standard_v2")
	ApplicationGatewayTier_WAF_v2      = ApplicationGatewayTier("WAF_v2")
)

type AuthorizationUseStatus string

const (
	AuthorizationUseStatusAvailable = AuthorizationUseStatus("Available")
	AuthorizationUseStatusInUse     = AuthorizationUseStatus("InUse")
)

type AzureFirewallApplicationRuleProtocolType string

const (
	AzureFirewallApplicationRuleProtocolTypeHttp  = AzureFirewallApplicationRuleProtocolType("Http")
	AzureFirewallApplicationRuleProtocolTypeHttps = AzureFirewallApplicationRuleProtocolType("Https")
)

type AzureFirewallNetworkRuleProtocol string

const (
	AzureFirewallNetworkRuleProtocolTCP  = AzureFirewallNetworkRuleProtocol("TCP")
	AzureFirewallNetworkRuleProtocolUDP  = AzureFirewallNetworkRuleProtocol("UDP")
	AzureFirewallNetworkRuleProtocolAny  = AzureFirewallNetworkRuleProtocol("Any")
	AzureFirewallNetworkRuleProtocolICMP = AzureFirewallNetworkRuleProtocol("ICMP")
)

type AzureFirewallRCActionType string

const (
	AzureFirewallRCActionTypeAllow = AzureFirewallRCActionType("Allow")
	AzureFirewallRCActionTypeDeny  = AzureFirewallRCActionType("Deny")
)

type DhGroup string

const (
	DhGroupNone        = DhGroup("None")
	DhGroupDHGroup1    = DhGroup("DHGroup1")
	DhGroupDHGroup2    = DhGroup("DHGroup2")
	DhGroupDHGroup14   = DhGroup("DHGroup14")
	DhGroupDHGroup2048 = DhGroup("DHGroup2048")
	DhGroupECP256      = DhGroup("ECP256")
	DhGroupECP384      = DhGroup("ECP384")
	DhGroupDHGroup24   = DhGroup("DHGroup24")
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

type ExpressRoutePeeringState string

const (
	ExpressRoutePeeringStateDisabled = ExpressRoutePeeringState("Disabled")
	ExpressRoutePeeringStateEnabled  = ExpressRoutePeeringState("Enabled")
)

type ExpressRoutePeeringType string

const (
	ExpressRoutePeeringTypeAzurePublicPeering  = ExpressRoutePeeringType("AzurePublicPeering")
	ExpressRoutePeeringTypeAzurePrivatePeering = ExpressRoutePeeringType("AzurePrivatePeering")
	ExpressRoutePeeringTypeMicrosoftPeering    = ExpressRoutePeeringType("MicrosoftPeering")
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

type IkeEncryption string

const (
	IkeEncryptionDES       = IkeEncryption("DES")
	IkeEncryptionDES3      = IkeEncryption("DES3")
	IkeEncryptionAES128    = IkeEncryption("AES128")
	IkeEncryptionAES192    = IkeEncryption("AES192")
	IkeEncryptionAES256    = IkeEncryption("AES256")
	IkeEncryptionGCMAES256 = IkeEncryption("GCMAES256")
	IkeEncryptionGCMAES128 = IkeEncryption("GCMAES128")
)

type IkeIntegrity string

const (
	IkeIntegrityMD5       = IkeIntegrity("MD5")
	IkeIntegritySHA1      = IkeIntegrity("SHA1")
	IkeIntegritySHA256    = IkeIntegrity("SHA256")
	IkeIntegritySHA384    = IkeIntegrity("SHA384")
	IkeIntegrityGCMAES256 = IkeIntegrity("GCMAES256")
	IkeIntegrityGCMAES128 = IkeIntegrity("GCMAES128")
)

type IpsecEncryption string

const (
	IpsecEncryptionNone      = IpsecEncryption("None")
	IpsecEncryptionDES       = IpsecEncryption("DES")
	IpsecEncryptionDES3      = IpsecEncryption("DES3")
	IpsecEncryptionAES128    = IpsecEncryption("AES128")
	IpsecEncryptionAES192    = IpsecEncryption("AES192")
	IpsecEncryptionAES256    = IpsecEncryption("AES256")
	IpsecEncryptionGCMAES128 = IpsecEncryption("GCMAES128")
	IpsecEncryptionGCMAES192 = IpsecEncryption("GCMAES192")
	IpsecEncryptionGCMAES256 = IpsecEncryption("GCMAES256")
)

type IpsecIntegrity string

const (
	IpsecIntegrityMD5       = IpsecIntegrity("MD5")
	IpsecIntegritySHA1      = IpsecIntegrity("SHA1")
	IpsecIntegritySHA256    = IpsecIntegrity("SHA256")
	IpsecIntegrityGCMAES128 = IpsecIntegrity("GCMAES128")
	IpsecIntegrityGCMAES192 = IpsecIntegrity("GCMAES192")
	IpsecIntegrityGCMAES256 = IpsecIntegrity("GCMAES256")
)

type LoadBalancerSkuName string

const (
	LoadBalancerSkuNameBasic    = LoadBalancerSkuName("Basic")
	LoadBalancerSkuNameStandard = LoadBalancerSkuName("Standard")
)

type LoadDistribution string

const (
	LoadDistributionDefault          = LoadDistribution("Default")
	LoadDistributionSourceIP         = LoadDistribution("SourceIP")
	LoadDistributionSourceIPProtocol = LoadDistribution("SourceIPProtocol")
)

type PcProtocol string

const (
	PcProtocolTCP = PcProtocol("TCP")
	PcProtocolUDP = PcProtocol("UDP")
	PcProtocolAny = PcProtocol("Any")
)

type PfsGroup string

const (
	PfsGroupNone    = PfsGroup("None")
	PfsGroupPFS1    = PfsGroup("PFS1")
	PfsGroupPFS2    = PfsGroup("PFS2")
	PfsGroupPFS2048 = PfsGroup("PFS2048")
	PfsGroupECP256  = PfsGroup("ECP256")
	PfsGroupECP384  = PfsGroup("ECP384")
	PfsGroupPFS24   = PfsGroup("PFS24")
	PfsGroupPFS14   = PfsGroup("PFS14")
	PfsGroupPFSMM   = PfsGroup("PFSMM")
)

type ProbeProtocol string

const (
	ProbeProtocolHttp  = ProbeProtocol("Http")
	ProbeProtocolTcp   = ProbeProtocol("Tcp")
	ProbeProtocolHttps = ProbeProtocol("Https")
)

type PublicIPAddressSkuName string

const (
	PublicIPAddressSkuNameBasic    = PublicIPAddressSkuName("Basic")
	PublicIPAddressSkuNameStandard = PublicIPAddressSkuName("Standard")
)

type PublicIPPrefixSkuName string

const (
	PublicIPPrefixSkuNameStandard = PublicIPPrefixSkuName("Standard")
)

type RouteFilterRuleTypeEnum string

const (
	RouteFilterRuleTypeEnumCommunity = RouteFilterRuleTypeEnum("Community")
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
	TransportProtocolAll = TransportProtocol("All")
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
	VirtualNetworkGatewaySkuNameBasic            = VirtualNetworkGatewaySkuName("Basic")
	VirtualNetworkGatewaySkuNameHighPerformance  = VirtualNetworkGatewaySkuName("HighPerformance")
	VirtualNetworkGatewaySkuNameStandard         = VirtualNetworkGatewaySkuName("Standard")
	VirtualNetworkGatewaySkuNameUltraPerformance = VirtualNetworkGatewaySkuName("UltraPerformance")
	VirtualNetworkGatewaySkuNameVpnGw1           = VirtualNetworkGatewaySkuName("VpnGw1")
	VirtualNetworkGatewaySkuNameVpnGw2           = VirtualNetworkGatewaySkuName("VpnGw2")
	VirtualNetworkGatewaySkuNameVpnGw3           = VirtualNetworkGatewaySkuName("VpnGw3")
	VirtualNetworkGatewaySkuNameVpnGw1AZ         = VirtualNetworkGatewaySkuName("VpnGw1AZ")
	VirtualNetworkGatewaySkuNameVpnGw2AZ         = VirtualNetworkGatewaySkuName("VpnGw2AZ")
	VirtualNetworkGatewaySkuNameVpnGw3AZ         = VirtualNetworkGatewaySkuName("VpnGw3AZ")
	VirtualNetworkGatewaySkuNameErGw1AZ          = VirtualNetworkGatewaySkuName("ErGw1AZ")
	VirtualNetworkGatewaySkuNameErGw2AZ          = VirtualNetworkGatewaySkuName("ErGw2AZ")
	VirtualNetworkGatewaySkuNameErGw3AZ          = VirtualNetworkGatewaySkuName("ErGw3AZ")
)

type VirtualNetworkGatewaySkuTier string

const (
	VirtualNetworkGatewaySkuTierBasic            = VirtualNetworkGatewaySkuTier("Basic")
	VirtualNetworkGatewaySkuTierHighPerformance  = VirtualNetworkGatewaySkuTier("HighPerformance")
	VirtualNetworkGatewaySkuTierStandard         = VirtualNetworkGatewaySkuTier("Standard")
	VirtualNetworkGatewaySkuTierUltraPerformance = VirtualNetworkGatewaySkuTier("UltraPerformance")
	VirtualNetworkGatewaySkuTierVpnGw1           = VirtualNetworkGatewaySkuTier("VpnGw1")
	VirtualNetworkGatewaySkuTierVpnGw2           = VirtualNetworkGatewaySkuTier("VpnGw2")
	VirtualNetworkGatewaySkuTierVpnGw3           = VirtualNetworkGatewaySkuTier("VpnGw3")
	VirtualNetworkGatewaySkuTierVpnGw1AZ         = VirtualNetworkGatewaySkuTier("VpnGw1AZ")
	VirtualNetworkGatewaySkuTierVpnGw2AZ         = VirtualNetworkGatewaySkuTier("VpnGw2AZ")
	VirtualNetworkGatewaySkuTierVpnGw3AZ         = VirtualNetworkGatewaySkuTier("VpnGw3AZ")
	VirtualNetworkGatewaySkuTierErGw1AZ          = VirtualNetworkGatewaySkuTier("ErGw1AZ")
	VirtualNetworkGatewaySkuTierErGw2AZ          = VirtualNetworkGatewaySkuTier("ErGw2AZ")
	VirtualNetworkGatewaySkuTierErGw3AZ          = VirtualNetworkGatewaySkuTier("ErGw3AZ")
)

type VirtualNetworkGatewayTypeEnum string

const (
	VirtualNetworkGatewayTypeEnumVpn          = VirtualNetworkGatewayTypeEnum("Vpn")
	VirtualNetworkGatewayTypeEnumExpressRoute = VirtualNetworkGatewayTypeEnum("ExpressRoute")
)

type VirtualNetworkPeeringStateEnum string

const (
	VirtualNetworkPeeringStateEnumInitiated    = VirtualNetworkPeeringStateEnum("Initiated")
	VirtualNetworkPeeringStateEnumConnected    = VirtualNetworkPeeringStateEnum("Connected")
	VirtualNetworkPeeringStateEnumDisconnected = VirtualNetworkPeeringStateEnum("Disconnected")
)

type VpnClientProtocol string

const (
	VpnClientProtocolIkeV2   = VpnClientProtocol("IkeV2")
	VpnClientProtocolSSTP    = VpnClientProtocol("SSTP")
	VpnClientProtocolOpenVPN = VpnClientProtocol("OpenVPN")
)

type VpnType string

const (
	VpnTypePolicyBased = VpnType("PolicyBased")
	VpnTypeRouteBased  = VpnType("RouteBased")
)

func init() {
}
