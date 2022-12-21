


package v20221101preview

type ApmType string

const (
	ApmTypeApplicationInsights = ApmType("ApplicationInsights")
	ApmTypeAppDynamics         = ApmType("AppDynamics")
	ApmTypeDynatrace           = ApmType("Dynatrace")
	ApmTypeNewRelic            = ApmType("NewRelic")
	ApmTypeElasticAPM          = ApmType("ElasticAPM")
)

type BackendProtocol string

const (
	BackendProtocolGRPC    = BackendProtocol("GRPC")
	BackendProtocolDefault = BackendProtocol("Default")
)

type BindingType string

const (
	BindingTypeApplicationInsights = BindingType("ApplicationInsights")
	BindingTypeApacheSkyWalking    = BindingType("ApacheSkyWalking")
	BindingTypeAppDynamics         = BindingType("AppDynamics")
	BindingTypeDynatrace           = BindingType("Dynatrace")
	BindingTypeNewRelic            = BindingType("NewRelic")
	BindingTypeElasticAPM          = BindingType("ElasticAPM")
	BindingTypeCACertificates      = BindingType("CACertificates")
)

type DevToolPortalFeatureState string

const (
	// Enable the plugin in Dev Tool Portal.
	DevToolPortalFeatureStateEnabled = DevToolPortalFeatureState("Enabled")
	// Disable the plugin in Dev Tool Portal.
	DevToolPortalFeatureStateDisabled = DevToolPortalFeatureState("Disabled")
)

type GatewayRouteConfigProtocol string

const (
	GatewayRouteConfigProtocolHTTP  = GatewayRouteConfigProtocol("HTTP")
	GatewayRouteConfigProtocolHTTPS = GatewayRouteConfigProtocol("HTTPS")
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

type SessionAffinity string

const (
	SessionAffinityCookie = SessionAffinity("Cookie")
	SessionAffinityNone   = SessionAffinity("None")
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
