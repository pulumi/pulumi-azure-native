


package v20220301preview

type AccountConfiguration string

const (
	AccountConfigurationFree     = AccountConfiguration("Free")
	AccountConfigurationCapacity = AccountConfiguration("Capacity")
)

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned,UserAssigned")
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
