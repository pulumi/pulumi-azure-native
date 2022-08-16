


package v20220201

type AccountConfiguration string

const (
	AccountConfigurationFree     = AccountConfiguration("Free")
	AccountConfigurationCapacity = AccountConfiguration("Capacity")
)

type ModelingFeatures string

const (
	ModelingFeaturesBasic    = ModelingFeatures("Basic")
	ModelingFeaturesStandard = ModelingFeatures("Standard")
	ModelingFeaturesPremium  = ModelingFeatures("Premium")
)

type ModelingFrequency string

const (
	ModelingFrequencyLow    = ModelingFrequency("Low")
	ModelingFrequencyMedium = ModelingFrequency("Medium")
	ModelingFrequencyHigh   = ModelingFrequency("High")
)

type ModelingSize string

const (
	ModelingSizeSmall  = ModelingSize("Small")
	ModelingSizeMedium = ModelingSize("Medium")
	ModelingSizeLarge  = ModelingSize("Large")
)

type PrincipalType string

const (
	PrincipalTypeApplication = PrincipalType("Application")
	PrincipalTypeUser        = PrincipalType("User")
)

func init() {
}
