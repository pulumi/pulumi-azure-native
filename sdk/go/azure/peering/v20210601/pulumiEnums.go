


package v20210601

type DirectPeeringType string

const (
	DirectPeeringTypeEdge     = DirectPeeringType("Edge")
	DirectPeeringTypeTransit  = DirectPeeringType("Transit")
	DirectPeeringTypeCdn      = DirectPeeringType("Cdn")
	DirectPeeringTypeInternal = DirectPeeringType("Internal")
	DirectPeeringTypeIx       = DirectPeeringType("Ix")
	DirectPeeringTypeIxRs     = DirectPeeringType("IxRs")
	DirectPeeringTypeVoice    = DirectPeeringType("Voice")
)

type Kind string

const (
	KindDirect   = Kind("Direct")
	KindExchange = Kind("Exchange")
)

type Role string

const (
	RoleNoc        = Role("Noc")
	RolePolicy     = Role("Policy")
	RoleTechnical  = Role("Technical")
	RoleService    = Role("Service")
	RoleEscalation = Role("Escalation")
	RoleOther      = Role("Other")
)

type SessionAddressProvider string

const (
	SessionAddressProviderMicrosoft = SessionAddressProvider("Microsoft")
	SessionAddressProviderPeer      = SessionAddressProvider("Peer")
)

func init() {
}
