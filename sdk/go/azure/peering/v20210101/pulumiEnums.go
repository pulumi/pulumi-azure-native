


package v20210101

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

type Family string

const (
	FamilyDirect   = Family("Direct")
	FamilyExchange = Family("Exchange")
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

type Size string

const (
	SizeFree      = Size("Free")
	SizeMetered   = Size("Metered")
	SizeUnlimited = Size("Unlimited")
)

type Tier string

const (
	TierBasic   = Tier("Basic")
	TierPremium = Tier("Premium")
)

type ValidationState string

const (
	ValidationStateNone     = ValidationState("None")
	ValidationStatePending  = ValidationState("Pending")
	ValidationStateApproved = ValidationState("Approved")
	ValidationStateFailed   = ValidationState("Failed")
)

func init() {
}
