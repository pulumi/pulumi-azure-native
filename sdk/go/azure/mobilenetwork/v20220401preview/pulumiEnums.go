


package v20220401preview

type BillingSku string

const (
	// Evaluation package plan
	BillingSkuEvaluationPackage = BillingSku("EvaluationPackage")
	// Flagship starter package plan
	BillingSkuFlagshipStarterPackage = BillingSku("FlagshipStarterPackage")
	// Edge site 2Gbps plan
	BillingSkuEdgeSite2GBPS = BillingSku("EdgeSite2GBPS")
	// Edge site 3Gbps plan
	BillingSkuEdgeSite3GBPS = BillingSku("EdgeSite3GBPS")
	// Edge site 4Gbps plan
	BillingSkuEdgeSite4GBPS = BillingSku("EdgeSite4GBPS")
	// Medium package plan
	BillingSkuMediumPackage = BillingSku("MediumPackage")
	// Large package plan
	BillingSkuLargePackage = BillingSku("LargePackage")
)

type CoreNetworkType string

const (
	// 5G core
	CoreNetworkType_5GC = CoreNetworkType("5GC")
	// EPC / 4G core
	CoreNetworkTypeEPC = CoreNetworkType("EPC")
)

type CreatedByType string

const (
	CreatedByTypeUser            = CreatedByType("User")
	CreatedByTypeApplication     = CreatedByType("Application")
	CreatedByTypeManagedIdentity = CreatedByType("ManagedIdentity")
	CreatedByTypeKey             = CreatedByType("Key")
)

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned,UserAssigned")
)

type NaptEnabled string

const (
	// NAPT is enabled
	NaptEnabledEnabled = NaptEnabled("Enabled")
	// NAPT is disabled
	NaptEnabledDisabled = NaptEnabled("Disabled")
)

type PduSessionType string

const (
	PduSessionTypeIPv4 = PduSessionType("IPv4")
	PduSessionTypeIPv6 = PduSessionType("IPv6")
)

type PlatformType string

const (
	// If this option is chosen, you must set one of "azureStackEdgeDevice", "connectedCluster" or "customLocation". If multiple are set then "customLocation" will take precedence over "connectedCluster" which takes precedence over "azureStackEdgeDevice".
	PlatformType_AKS_HCI = PlatformType("AKS-HCI")
	// If this option is chosen, you must set one of "connectedCluster" or "customLocation". If multiple are set then "customLocation" will take precedence over "connectedCluster".
	PlatformTypeBaseVM = PlatformType("BaseVM")
)

type PreemptionCapability string

const (
	// Cannot preempt
	PreemptionCapabilityNotPreempt = PreemptionCapability("NotPreempt")
	// May preempt
	PreemptionCapabilityMayPreempt = PreemptionCapability("MayPreempt")
)

type PreemptionVulnerability string

const (
	// Cannot be preempted
	PreemptionVulnerabilityNotPreemptable = PreemptionVulnerability("NotPreemptable")
	// May be preempted
	PreemptionVulnerabilityPreemptable = PreemptionVulnerability("Preemptable")
)

type SdfDirection string

const (
	// Traffic flowing from the UE to the data network.
	SdfDirectionUplink = SdfDirection("Uplink")
	// Traffic flowing from the data network to the UE.
	SdfDirectionDownlink = SdfDirection("Downlink")
	// Traffic flowing both to and from the UE.
	SdfDirectionBidirectional = SdfDirection("Bidirectional")
)

type TrafficControlPermission string

const (
	// Traffic matching this rule is allowed to flow.
	TrafficControlPermissionEnabled = TrafficControlPermission("Enabled")
	// Traffic matching this rule is not allowed to flow.
	TrafficControlPermissionBlocked = TrafficControlPermission("Blocked")
)

func init() {
}
