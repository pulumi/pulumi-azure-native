


package v20150601preview

type Protocol string

const (
	ProtocolTCP = Protocol("TCP")
	ProtocolUDP = Protocol("UDP")
	ProtocolAll = Protocol("*")
)

type Status string

const (
	StatusRevoked   = Status("Revoked")
	StatusInitiated = Status("Initiated")
)

type StatusReason string

const (
	StatusReasonExpired               = StatusReason("Expired")
	StatusReasonUserRequested         = StatusReason("UserRequested")
	StatusReasonNewerRequestInitiated = StatusReason("NewerRequestInitiated")
)

func init() {
}
