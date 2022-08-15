


package v20190901preview

type DirectPeeringType string

const (
	DirectPeeringTypeEdge     = DirectPeeringType("Edge")
	DirectPeeringTypeTransit  = DirectPeeringType("Transit")
	DirectPeeringTypeCdn      = DirectPeeringType("Cdn")
	DirectPeeringTypeInternal = DirectPeeringType("Internal")
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

type Name string

const (
	Name_Basic_Exchange_Free      = Name("Basic_Exchange_Free")
	Name_Basic_Direct_Free        = Name("Basic_Direct_Free")
	Name_Premium_Exchange_Metered = Name("Premium_Exchange_Metered")
	Name_Premium_Direct_Free      = Name("Premium_Direct_Free")
	Name_Premium_Direct_Metered   = Name("Premium_Direct_Metered")
	Name_Premium_Direct_Unlimited = Name("Premium_Direct_Unlimited")
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
