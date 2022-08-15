


package v20220101preview

type CoreNetworkType string

const (
	CoreNetworkType_5GC = CoreNetworkType("5GC")
	CoreNetworkTypeEPC  = CoreNetworkType("EPC")
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
	NaptEnabledEnabled  = NaptEnabled("Enabled")
	NaptEnabledDisabled = NaptEnabled("Disabled")
)

type PduSessionType string

const (
	PduSessionTypeIPv4 = PduSessionType("IPv4")
	PduSessionTypeIPv6 = PduSessionType("IPv6")
)

type PreemptionCapability string

const (
	PreemptionCapabilityNotPreempt = PreemptionCapability("NotPreempt")
	PreemptionCapabilityMayPreempt = PreemptionCapability("MayPreempt")
)

type PreemptionVulnerability string

const (
	PreemptionVulnerabilityNotPreemptable = PreemptionVulnerability("NotPreemptable")
	PreemptionVulnerabilityPreemptable    = PreemptionVulnerability("Preemptable")
)

type SdfDirection string

const (
	SdfDirectionUplink        = SdfDirection("Uplink")
	SdfDirectionDownlink      = SdfDirection("Downlink")
	SdfDirectionBidirectional = SdfDirection("Bidirectional")
)

type TrafficControlPermission string

const (
	TrafficControlPermissionEnabled = TrafficControlPermission("Enabled")
	TrafficControlPermissionBlocked = TrafficControlPermission("Blocked")
)

func init() {
}
