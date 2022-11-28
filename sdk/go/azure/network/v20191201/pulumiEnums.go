


package v20191201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

type ApplicationGatewayCustomErrorStatusCode string

const (
	ApplicationGatewayCustomErrorStatusCodeHttpStatus403 = ApplicationGatewayCustomErrorStatusCode("HttpStatus403")
	ApplicationGatewayCustomErrorStatusCodeHttpStatus502 = ApplicationGatewayCustomErrorStatusCode("HttpStatus502")
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
	AzureFirewallApplicationRuleProtocolTypeMssql = AzureFirewallApplicationRuleProtocolType("Mssql")
)

type AzureFirewallNatRCActionType string

const (
	AzureFirewallNatRCActionTypeSnat = AzureFirewallNatRCActionType("Snat")
	AzureFirewallNatRCActionTypeDnat = AzureFirewallNatRCActionType("Dnat")
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

type AzureFirewallSkuName string

const (
	AzureFirewallSkuName_AZFW_VNet = AzureFirewallSkuName("AZFW_VNet")
	AzureFirewallSkuName_AZFW_Hub  = AzureFirewallSkuName("AZFW_Hub")
)

type AzureFirewallSkuTier string

const (
	AzureFirewallSkuTierStandard = AzureFirewallSkuTier("Standard")
)

type AzureFirewallThreatIntelMode string

const (
	AzureFirewallThreatIntelModeAlert = AzureFirewallThreatIntelMode("Alert")
	AzureFirewallThreatIntelModeDeny  = AzureFirewallThreatIntelMode("Deny")
	AzureFirewallThreatIntelModeOff   = AzureFirewallThreatIntelMode("Off")
)

type ConnectionMonitorEndpointFilterItemType string

const (
	ConnectionMonitorEndpointFilterItemTypeAgentAddress = ConnectionMonitorEndpointFilterItemType("AgentAddress")
)

type ConnectionMonitorEndpointFilterType string

const (
	ConnectionMonitorEndpointFilterTypeInclude = ConnectionMonitorEndpointFilterType("Include")
)

type ConnectionMonitorTestConfigurationProtocol string

const (
	ConnectionMonitorTestConfigurationProtocolTcp  = ConnectionMonitorTestConfigurationProtocol("Tcp")
	ConnectionMonitorTestConfigurationProtocolHttp = ConnectionMonitorTestConfigurationProtocol("Http")
	ConnectionMonitorTestConfigurationProtocolIcmp = ConnectionMonitorTestConfigurationProtocol("Icmp")
)

type DdosCustomPolicyProtocol string

const (
	DdosCustomPolicyProtocolTcp = DdosCustomPolicyProtocol("Tcp")
	DdosCustomPolicyProtocolUdp = DdosCustomPolicyProtocol("Udp")
	DdosCustomPolicyProtocolSyn = DdosCustomPolicyProtocol("Syn")
)

type DdosCustomPolicyTriggerSensitivityOverride string

const (
	DdosCustomPolicyTriggerSensitivityOverrideRelaxed = DdosCustomPolicyTriggerSensitivityOverride("Relaxed")
	DdosCustomPolicyTriggerSensitivityOverrideLow     = DdosCustomPolicyTriggerSensitivityOverride("Low")
	DdosCustomPolicyTriggerSensitivityOverrideDefault = DdosCustomPolicyTriggerSensitivityOverride("Default")
	DdosCustomPolicyTriggerSensitivityOverrideHigh    = DdosCustomPolicyTriggerSensitivityOverride("High")
)

type DdosSettingsProtectionCoverage string

const (
	DdosSettingsProtectionCoverageBasic    = DdosSettingsProtectionCoverage("Basic")
	DdosSettingsProtectionCoverageStandard = DdosSettingsProtectionCoverage("Standard")
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
	ExpressRouteCircuitSkuTierBasic    = ExpressRouteCircuitSkuTier("Basic")
	ExpressRouteCircuitSkuTierLocal    = ExpressRouteCircuitSkuTier("Local")
)

type ExpressRouteLinkAdminState string

const (
	ExpressRouteLinkAdminStateEnabled  = ExpressRouteLinkAdminState("Enabled")
	ExpressRouteLinkAdminStateDisabled = ExpressRouteLinkAdminState("Disabled")
)

type ExpressRouteLinkMacSecCipher string

const (
	ExpressRouteLinkMacSecCipher_Gcm_Aes_128 = ExpressRouteLinkMacSecCipher("gcm-aes-128")
	ExpressRouteLinkMacSecCipher_Gcm_Aes_256 = ExpressRouteLinkMacSecCipher("gcm-aes-256")
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

type ExpressRoutePortsEncapsulation string

const (
	ExpressRoutePortsEncapsulationDot1Q = ExpressRoutePortsEncapsulation("Dot1Q")
	ExpressRoutePortsEncapsulationQinQ  = ExpressRoutePortsEncapsulation("QinQ")
)

type FirewallPolicyFilterRuleActionType string

const (
	FirewallPolicyFilterRuleActionTypeAllow = FirewallPolicyFilterRuleActionType("Allow")
	FirewallPolicyFilterRuleActionTypeDeny  = FirewallPolicyFilterRuleActionType("Deny")
)

type FirewallPolicyNatRuleActionType string

const (
	FirewallPolicyNatRuleActionTypeDNAT = FirewallPolicyNatRuleActionType("DNAT")
)

type FirewallPolicyRuleConditionApplicationProtocolType string

const (
	FirewallPolicyRuleConditionApplicationProtocolTypeHttp  = FirewallPolicyRuleConditionApplicationProtocolType("Http")
	FirewallPolicyRuleConditionApplicationProtocolTypeHttps = FirewallPolicyRuleConditionApplicationProtocolType("Https")
)

type FirewallPolicyRuleConditionNetworkProtocol string

const (
	FirewallPolicyRuleConditionNetworkProtocolTCP  = FirewallPolicyRuleConditionNetworkProtocol("TCP")
	FirewallPolicyRuleConditionNetworkProtocolUDP  = FirewallPolicyRuleConditionNetworkProtocol("UDP")
	FirewallPolicyRuleConditionNetworkProtocolAny  = FirewallPolicyRuleConditionNetworkProtocol("Any")
	FirewallPolicyRuleConditionNetworkProtocolICMP = FirewallPolicyRuleConditionNetworkProtocol("ICMP")
)

type FirewallPolicyRuleConditionType string

const (
	FirewallPolicyRuleConditionTypeApplicationRuleCondition = FirewallPolicyRuleConditionType("ApplicationRuleCondition")
	FirewallPolicyRuleConditionTypeNetworkRuleCondition     = FirewallPolicyRuleConditionType("NetworkRuleCondition")
	FirewallPolicyRuleConditionTypeNatRuleCondition         = FirewallPolicyRuleConditionType("NatRuleCondition")
)

type FirewallPolicyRuleType string

const (
	FirewallPolicyRuleTypeFirewallPolicyNatRule    = FirewallPolicyRuleType("FirewallPolicyNatRule")
	FirewallPolicyRuleTypeFirewallPolicyFilterRule = FirewallPolicyRuleType("FirewallPolicyFilterRule")
)

type FlowLogFormatType string

const (
	FlowLogFormatTypeJSON = FlowLogFormatType("JSON")
)

type HTTPConfigurationMethod string

const (
	HTTPConfigurationMethodGet  = HTTPConfigurationMethod("Get")
	HTTPConfigurationMethodPost = HTTPConfigurationMethod("Post")
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

type LoadBalancerOutboundRuleProtocol string

const (
	LoadBalancerOutboundRuleProtocolTcp = LoadBalancerOutboundRuleProtocol("Tcp")
	LoadBalancerOutboundRuleProtocolUdp = LoadBalancerOutboundRuleProtocol("Udp")
	LoadBalancerOutboundRuleProtocolAll = LoadBalancerOutboundRuleProtocol("All")
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

type ManagedRuleEnabledState string

const (
	ManagedRuleEnabledStateDisabled = ManagedRuleEnabledState("Disabled")
)

type NatGatewaySkuName string

const (
	NatGatewaySkuNameStandard = NatGatewaySkuName("Standard")
)

type OutputType string

const (
	OutputTypeWorkspace = OutputType("Workspace")
)

type OwaspCrsExclusionEntryMatchVariable string

const (
	OwaspCrsExclusionEntryMatchVariableRequestHeaderNames = OwaspCrsExclusionEntryMatchVariable("RequestHeaderNames")
	OwaspCrsExclusionEntryMatchVariableRequestCookieNames = OwaspCrsExclusionEntryMatchVariable("RequestCookieNames")
	OwaspCrsExclusionEntryMatchVariableRequestArgNames    = OwaspCrsExclusionEntryMatchVariable("RequestArgNames")
)

type OwaspCrsExclusionEntrySelectorMatchOperator string

const (
	OwaspCrsExclusionEntrySelectorMatchOperatorEquals     = OwaspCrsExclusionEntrySelectorMatchOperator("Equals")
	OwaspCrsExclusionEntrySelectorMatchOperatorContains   = OwaspCrsExclusionEntrySelectorMatchOperator("Contains")
	OwaspCrsExclusionEntrySelectorMatchOperatorStartsWith = OwaspCrsExclusionEntrySelectorMatchOperator("StartsWith")
	OwaspCrsExclusionEntrySelectorMatchOperatorEndsWith   = OwaspCrsExclusionEntrySelectorMatchOperator("EndsWith")
	OwaspCrsExclusionEntrySelectorMatchOperatorEqualsAny  = OwaspCrsExclusionEntrySelectorMatchOperator("EqualsAny")
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

type PreferredIPVersion string

const (
	PreferredIPVersionIPv4 = PreferredIPVersion("IPv4")
	PreferredIPVersionIPv6 = PreferredIPVersion("IPv6")
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

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

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
	SecurityRuleProtocolIcmp     = SecurityRuleProtocol("Icmp")
	SecurityRuleProtocolEsp      = SecurityRuleProtocol("Esp")
	SecurityRuleProtocolAsterisk = SecurityRuleProtocol("*")
	SecurityRuleProtocolAh       = SecurityRuleProtocol("Ah")
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

type VirtualNetworkGatewayConnectionProtocol string

const (
	VirtualNetworkGatewayConnectionProtocolIKEv2 = VirtualNetworkGatewayConnectionProtocol("IKEv2")
	VirtualNetworkGatewayConnectionProtocolIKEv1 = VirtualNetworkGatewayConnectionProtocol("IKEv1")
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
	VirtualNetworkGatewaySkuNameVpnGw4           = VirtualNetworkGatewaySkuName("VpnGw4")
	VirtualNetworkGatewaySkuNameVpnGw5           = VirtualNetworkGatewaySkuName("VpnGw5")
	VirtualNetworkGatewaySkuNameVpnGw1AZ         = VirtualNetworkGatewaySkuName("VpnGw1AZ")
	VirtualNetworkGatewaySkuNameVpnGw2AZ         = VirtualNetworkGatewaySkuName("VpnGw2AZ")
	VirtualNetworkGatewaySkuNameVpnGw3AZ         = VirtualNetworkGatewaySkuName("VpnGw3AZ")
	VirtualNetworkGatewaySkuNameVpnGw4AZ         = VirtualNetworkGatewaySkuName("VpnGw4AZ")
	VirtualNetworkGatewaySkuNameVpnGw5AZ         = VirtualNetworkGatewaySkuName("VpnGw5AZ")
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
	VirtualNetworkGatewaySkuTierVpnGw4           = VirtualNetworkGatewaySkuTier("VpnGw4")
	VirtualNetworkGatewaySkuTierVpnGw5           = VirtualNetworkGatewaySkuTier("VpnGw5")
	VirtualNetworkGatewaySkuTierVpnGw1AZ         = VirtualNetworkGatewaySkuTier("VpnGw1AZ")
	VirtualNetworkGatewaySkuTierVpnGw2AZ         = VirtualNetworkGatewaySkuTier("VpnGw2AZ")
	VirtualNetworkGatewaySkuTierVpnGw3AZ         = VirtualNetworkGatewaySkuTier("VpnGw3AZ")
	VirtualNetworkGatewaySkuTierVpnGw4AZ         = VirtualNetworkGatewaySkuTier("VpnGw4AZ")
	VirtualNetworkGatewaySkuTierVpnGw5AZ         = VirtualNetworkGatewaySkuTier("VpnGw5AZ")
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

type VpnAuthenticationType string

const (
	VpnAuthenticationTypeCertificate = VpnAuthenticationType("Certificate")
	VpnAuthenticationTypeRadius      = VpnAuthenticationType("Radius")
	VpnAuthenticationTypeAAD         = VpnAuthenticationType("AAD")
)

type VpnClientProtocol string

const (
	VpnClientProtocolIkeV2   = VpnClientProtocol("IkeV2")
	VpnClientProtocolSSTP    = VpnClientProtocol("SSTP")
	VpnClientProtocolOpenVPN = VpnClientProtocol("OpenVPN")
)

type VpnGatewayGeneration string

const (
	VpnGatewayGenerationNone        = VpnGatewayGeneration("None")
	VpnGatewayGenerationGeneration1 = VpnGatewayGeneration("Generation1")
	VpnGatewayGenerationGeneration2 = VpnGatewayGeneration("Generation2")
)

type VpnGatewayTunnelingProtocol string

const (
	VpnGatewayTunnelingProtocolIkeV2   = VpnGatewayTunnelingProtocol("IkeV2")
	VpnGatewayTunnelingProtocolOpenVPN = VpnGatewayTunnelingProtocol("OpenVPN")
)

type VpnType string

const (
	VpnTypePolicyBased = VpnType("PolicyBased")
	VpnTypeRouteBased  = VpnType("RouteBased")
)

type WebApplicationFirewallAction string

const (
	WebApplicationFirewallActionAllow = WebApplicationFirewallAction("Allow")
	WebApplicationFirewallActionBlock = WebApplicationFirewallAction("Block")
	WebApplicationFirewallActionLog   = WebApplicationFirewallAction("Log")
)

type WebApplicationFirewallEnabledState string

const (
	WebApplicationFirewallEnabledStateDisabled = WebApplicationFirewallEnabledState("Disabled")
	WebApplicationFirewallEnabledStateEnabled  = WebApplicationFirewallEnabledState("Enabled")
)

type WebApplicationFirewallMatchVariable string

const (
	WebApplicationFirewallMatchVariableRemoteAddr     = WebApplicationFirewallMatchVariable("RemoteAddr")
	WebApplicationFirewallMatchVariableRequestMethod  = WebApplicationFirewallMatchVariable("RequestMethod")
	WebApplicationFirewallMatchVariableQueryString    = WebApplicationFirewallMatchVariable("QueryString")
	WebApplicationFirewallMatchVariablePostArgs       = WebApplicationFirewallMatchVariable("PostArgs")
	WebApplicationFirewallMatchVariableRequestUri     = WebApplicationFirewallMatchVariable("RequestUri")
	WebApplicationFirewallMatchVariableRequestHeaders = WebApplicationFirewallMatchVariable("RequestHeaders")
	WebApplicationFirewallMatchVariableRequestBody    = WebApplicationFirewallMatchVariable("RequestBody")
	WebApplicationFirewallMatchVariableRequestCookies = WebApplicationFirewallMatchVariable("RequestCookies")
)

type WebApplicationFirewallMode string

const (
	WebApplicationFirewallModePrevention = WebApplicationFirewallMode("Prevention")
	WebApplicationFirewallModeDetection  = WebApplicationFirewallMode("Detection")
)

type WebApplicationFirewallOperator string

const (
	WebApplicationFirewallOperatorIPMatch            = WebApplicationFirewallOperator("IPMatch")
	WebApplicationFirewallOperatorEqual              = WebApplicationFirewallOperator("Equal")
	WebApplicationFirewallOperatorContains           = WebApplicationFirewallOperator("Contains")
	WebApplicationFirewallOperatorLessThan           = WebApplicationFirewallOperator("LessThan")
	WebApplicationFirewallOperatorGreaterThan        = WebApplicationFirewallOperator("GreaterThan")
	WebApplicationFirewallOperatorLessThanOrEqual    = WebApplicationFirewallOperator("LessThanOrEqual")
	WebApplicationFirewallOperatorGreaterThanOrEqual = WebApplicationFirewallOperator("GreaterThanOrEqual")
	WebApplicationFirewallOperatorBeginsWith         = WebApplicationFirewallOperator("BeginsWith")
	WebApplicationFirewallOperatorEndsWith           = WebApplicationFirewallOperator("EndsWith")
	WebApplicationFirewallOperatorRegex              = WebApplicationFirewallOperator("Regex")
	WebApplicationFirewallOperatorGeoMatch           = WebApplicationFirewallOperator("GeoMatch")
)

type WebApplicationFirewallRuleType string

const (
	WebApplicationFirewallRuleTypeMatchRule = WebApplicationFirewallRuleType("MatchRule")
	WebApplicationFirewallRuleTypeInvalid   = WebApplicationFirewallRuleType("Invalid")
)

type WebApplicationFirewallTransform string

const (
	WebApplicationFirewallTransformLowercase        = WebApplicationFirewallTransform("Lowercase")
	WebApplicationFirewallTransformTrim             = WebApplicationFirewallTransform("Trim")
	WebApplicationFirewallTransformUrlDecode        = WebApplicationFirewallTransform("UrlDecode")
	WebApplicationFirewallTransformUrlEncode        = WebApplicationFirewallTransform("UrlEncode")
	WebApplicationFirewallTransformRemoveNulls      = WebApplicationFirewallTransform("RemoveNulls")
	WebApplicationFirewallTransformHtmlEntityDecode = WebApplicationFirewallTransform("HtmlEntityDecode")
)

func init() {
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
