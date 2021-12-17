


package v20181001

type FeatureFlags string

const (
	FeatureFlagsServiceMode            = FeatureFlags("ServiceMode")
	FeatureFlagsEnableConnectivityLogs = FeatureFlags("EnableConnectivityLogs")
)

type SignalRSkuTier string

const (
	SignalRSkuTierFree     = SignalRSkuTier("Free")
	SignalRSkuTierBasic    = SignalRSkuTier("Basic")
	SignalRSkuTierStandard = SignalRSkuTier("Standard")
	SignalRSkuTierPremium  = SignalRSkuTier("Premium")
)

func init() {
}
