


package v20200315preview

type DiskPoolTier string

const (
	DiskPoolTierBasic    = DiskPoolTier("Basic")
	DiskPoolTierStandard = DiskPoolTier("Standard")
	DiskPoolTierPremium  = DiskPoolTier("Premium")
)

func init() {
}
