


package v20221101

type AuthenticationType string

const (
	// Use AAD SSO to authenticate the user (this requires internet access).
	AuthenticationTypeAAD = AuthenticationType("AAD")
	// Use locally stored passwords to authenticate the user.
	AuthenticationTypePassword = AuthenticationType("Password")
)

type BillingSku string

const (
	// 100 Mbps, 20 active SIMs plan, 2 RANs
	BillingSkuG0 = BillingSku("G0")
	// 1 Gbps, 100 active SIMs plan, 5 RANs
	BillingSkuG1 = BillingSku("G1")
	// 2 Gbps, 200 active SIMs plan, 10 RANs
	BillingSkuG2 = BillingSku("G2")
	// 3 Gbps, 300 active SIMs plan
	BillingSkuG3 = BillingSku("G3")
	// 4 Gbps, 400 active SIMs plan
	BillingSkuG4 = BillingSku("G4")
	// 5 Gbps, 500 active SIMs plan
	BillingSkuG5 = BillingSku("G5")
	// 10 Gbps, 1000 active SIMs plan
	BillingSkuG10 = BillingSku("G10")
)

type CoreNetworkType string

const (
	// 5G core
	CoreNetworkType_5GC = CoreNetworkType("5GC")
	// EPC / 4G core
	CoreNetworkTypeEPC = CoreNetworkType("EPC")
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
	// If this option is chosen, you must set one of "azureStackEdgeDevice", "connectedCluster" or "customLocation". If multiple are set, they must be consistent with each other.
	PlatformType_AKS_HCI = PlatformType("AKS-HCI")
	// If this option is chosen, you must set one of "azureStackHciCluster", "connectedCluster" or "customLocation". If multiple are set, they must be consistent with each other.
	PlatformType_3P_AZURE_STACK_HCI = PlatformType("3P-AZURE-STACK-HCI")
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
