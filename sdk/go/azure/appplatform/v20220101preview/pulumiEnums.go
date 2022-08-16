


package v20220101preview

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

type StorageType string

const (
	StorageTypeStorageAccount = StorageType("StorageAccount")
)

type Type string

const (
	TypeAzureFileVolume = Type("AzureFileVolume")
)

func init() {
}
