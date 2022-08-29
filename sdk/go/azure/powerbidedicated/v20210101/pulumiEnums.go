


package v20210101

type CapacitySkuTier string

const (
	CapacitySkuTier_PBIE_Azure     = CapacitySkuTier("PBIE_Azure")
	CapacitySkuTierPremium         = CapacitySkuTier("Premium")
	CapacitySkuTierAutoPremiumHost = CapacitySkuTier("AutoPremiumHost")
)

type IdentityType string

const (
	IdentityTypeUser            = IdentityType("User")
	IdentityTypeApplication     = IdentityType("Application")
	IdentityTypeManagedIdentity = IdentityType("ManagedIdentity")
	IdentityTypeKey             = IdentityType("Key")
)

type Mode string

const (
	ModeGen1 = Mode("Gen1")
	ModeGen2 = Mode("Gen2")
)

type VCoreSkuTier string

const (
	VCoreSkuTierAutoScale = VCoreSkuTier("AutoScale")
)

func init() {
}
