


package purview

type PublicNetworkAccess string

const (
	PublicNetworkAccessNotSpecified = PublicNetworkAccess("NotSpecified")
	PublicNetworkAccessEnabled      = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled     = PublicNetworkAccess("Disabled")
)

type Status string

const (
	StatusUnknown      = Status("Unknown")
	StatusPending      = Status("Pending")
	StatusApproved     = Status("Approved")
	StatusRejected     = Status("Rejected")
	StatusDisconnected = Status("Disconnected")
)

type Type string

const (
	TypeSystemAssigned = Type("SystemAssigned")
)

func init() {
}
