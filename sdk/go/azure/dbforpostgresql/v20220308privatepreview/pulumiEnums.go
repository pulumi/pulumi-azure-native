


package v20220308privatepreview

type ServerVersion string

const (
	ServerVersion_14 = ServerVersion("14")
	ServerVersion_13 = ServerVersion("13")
	ServerVersion_12 = ServerVersion("12")
	ServerVersion_11 = ServerVersion("11")
)

type SkuTier string

const (
	SkuTierBurstable       = SkuTier("Burstable")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
)

func init() {
}
