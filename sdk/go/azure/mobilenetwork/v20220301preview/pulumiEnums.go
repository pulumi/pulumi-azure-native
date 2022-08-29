


package v20220301preview

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
