


package v20200501

type Status string

const (
	StatusApproved     = Status("Approved")
	StatusPending      = Status("Pending")
	StatusRejected     = Status("Rejected")
	StatusDisconnected = Status("Disconnected")
)

func init() {
}
