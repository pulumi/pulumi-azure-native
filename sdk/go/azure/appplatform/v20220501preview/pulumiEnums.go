


package v20220501preview

type BindingType string

const (
	BindingTypeApplicationInsights = BindingType("ApplicationInsights")
	BindingTypeApacheSkyWalking    = BindingType("ApacheSkyWalking")
	BindingTypeAppDynamics         = BindingType("AppDynamics")
	BindingTypeDynatrace           = BindingType("Dynatrace")
	BindingTypeNewRelic            = BindingType("NewRelic")
	BindingTypeElasticAPM          = BindingType("ElasticAPM")
)

type HTTPSchemeType string

const (
	HTTPSchemeTypeHTTP  = HTTPSchemeType("HTTP")
	HTTPSchemeTypeHTTPS = HTTPSchemeType("HTTPS")
)

type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                         = ManagedIdentityType("None")
	ManagedIdentityTypeSystemAssigned               = ManagedIdentityType("SystemAssigned")
	ManagedIdentityTypeUserAssigned                 = ManagedIdentityType("UserAssigned")
	ManagedIdentityType_SystemAssigned_UserAssigned = ManagedIdentityType("SystemAssigned,UserAssigned")
)

type ProbeActionType string

const (
	ProbeActionTypeHTTPGetAction   = ProbeActionType("HTTPGetAction")
	ProbeActionTypeTCPSocketAction = ProbeActionType("TCPSocketAction")
	ProbeActionTypeExecAction      = ProbeActionType("ExecAction")
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
