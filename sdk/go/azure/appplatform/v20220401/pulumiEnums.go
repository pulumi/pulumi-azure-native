


package v20220401

type BindingType string

const (
	BindingTypeApplicationInsights = BindingType("ApplicationInsights")
	BindingTypeApacheSkyWalking    = BindingType("ApacheSkyWalking")
	BindingTypeAppDynamics         = BindingType("AppDynamics")
	BindingTypeDynatrace           = BindingType("Dynatrace")
	BindingTypeNewRelic            = BindingType("NewRelic")
	BindingTypeElasticAPM          = BindingType("ElasticAPM")
)

type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                         = ManagedIdentityType("None")
	ManagedIdentityTypeSystemAssigned               = ManagedIdentityType("SystemAssigned")
	ManagedIdentityTypeUserAssigned                 = ManagedIdentityType("UserAssigned")
	ManagedIdentityType_SystemAssigned_UserAssigned = ManagedIdentityType("SystemAssigned,UserAssigned")
)

func init() {
}
