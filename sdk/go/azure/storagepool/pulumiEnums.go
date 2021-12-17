


package storagepool

type DiskPoolTier string

const (
	DiskPoolTierBasic    = DiskPoolTier("Basic")
	DiskPoolTierStandard = DiskPoolTier("Standard")
	DiskPoolTierPremium  = DiskPoolTier("Premium")
)

func init() {
}
