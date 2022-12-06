


package v20200101

type BackendEnabledState string

const (
	BackendEnabledStateEnabled  = BackendEnabledState("Enabled")
	BackendEnabledStateDisabled = BackendEnabledState("Disabled")
)

type DynamicCompressionEnabled string

const (
	DynamicCompressionEnabledEnabled  = DynamicCompressionEnabled("Enabled")
	DynamicCompressionEnabledDisabled = DynamicCompressionEnabled("Disabled")
)

type EnforceCertificateNameCheckEnabledState string

const (
	EnforceCertificateNameCheckEnabledStateEnabled  = EnforceCertificateNameCheckEnabledState("Enabled")
	EnforceCertificateNameCheckEnabledStateDisabled = EnforceCertificateNameCheckEnabledState("Disabled")
)

type FrontDoorEnabledState string

const (
	FrontDoorEnabledStateEnabled  = FrontDoorEnabledState("Enabled")
	FrontDoorEnabledStateDisabled = FrontDoorEnabledState("Disabled")
)

type FrontDoorForwardingProtocol string

const (
	FrontDoorForwardingProtocolHttpOnly     = FrontDoorForwardingProtocol("HttpOnly")
	FrontDoorForwardingProtocolHttpsOnly    = FrontDoorForwardingProtocol("HttpsOnly")
	FrontDoorForwardingProtocolMatchRequest = FrontDoorForwardingProtocol("MatchRequest")
)

type FrontDoorHealthProbeMethod string

const (
	FrontDoorHealthProbeMethodGET  = FrontDoorHealthProbeMethod("GET")
	FrontDoorHealthProbeMethodHEAD = FrontDoorHealthProbeMethod("HEAD")
)

type FrontDoorProtocol string

const (
	FrontDoorProtocolHttp  = FrontDoorProtocol("Http")
	FrontDoorProtocolHttps = FrontDoorProtocol("Https")
)

type FrontDoorQuery string

const (
	FrontDoorQueryStripNone      = FrontDoorQuery("StripNone")
	FrontDoorQueryStripAll       = FrontDoorQuery("StripAll")
	FrontDoorQueryStripOnly      = FrontDoorQuery("StripOnly")
	FrontDoorQueryStripAllExcept = FrontDoorQuery("StripAllExcept")
)

type FrontDoorRedirectProtocol string

const (
	FrontDoorRedirectProtocolHttpOnly     = FrontDoorRedirectProtocol("HttpOnly")
	FrontDoorRedirectProtocolHttpsOnly    = FrontDoorRedirectProtocol("HttpsOnly")
	FrontDoorRedirectProtocolMatchRequest = FrontDoorRedirectProtocol("MatchRequest")
)

type FrontDoorRedirectType string

const (
	FrontDoorRedirectTypeMoved             = FrontDoorRedirectType("Moved")
	FrontDoorRedirectTypeFound             = FrontDoorRedirectType("Found")
	FrontDoorRedirectTypeTemporaryRedirect = FrontDoorRedirectType("TemporaryRedirect")
	FrontDoorRedirectTypePermanentRedirect = FrontDoorRedirectType("PermanentRedirect")
)

type HeaderActionType string

const (
	HeaderActionTypeAppend    = HeaderActionType("Append")
	HeaderActionTypeDelete    = HeaderActionType("Delete")
	HeaderActionTypeOverwrite = HeaderActionType("Overwrite")
)

type HealthProbeEnabled string

const (
	HealthProbeEnabledEnabled  = HealthProbeEnabled("Enabled")
	HealthProbeEnabledDisabled = HealthProbeEnabled("Disabled")
)

type MatchProcessingBehavior string

const (
	MatchProcessingBehaviorContinue = MatchProcessingBehavior("Continue")
	MatchProcessingBehaviorStop     = MatchProcessingBehavior("Stop")
)

type RoutingRuleEnabledState string

const (
	RoutingRuleEnabledStateEnabled  = RoutingRuleEnabledState("Enabled")
	RoutingRuleEnabledStateDisabled = RoutingRuleEnabledState("Disabled")
)

type RulesEngineMatchVariable string

const (
	RulesEngineMatchVariableIsMobile                 = RulesEngineMatchVariable("IsMobile")
	RulesEngineMatchVariableRemoteAddr               = RulesEngineMatchVariable("RemoteAddr")
	RulesEngineMatchVariableRequestMethod            = RulesEngineMatchVariable("RequestMethod")
	RulesEngineMatchVariableQueryString              = RulesEngineMatchVariable("QueryString")
	RulesEngineMatchVariablePostArgs                 = RulesEngineMatchVariable("PostArgs")
	RulesEngineMatchVariableRequestUri               = RulesEngineMatchVariable("RequestUri")
	RulesEngineMatchVariableRequestPath              = RulesEngineMatchVariable("RequestPath")
	RulesEngineMatchVariableRequestFilename          = RulesEngineMatchVariable("RequestFilename")
	RulesEngineMatchVariableRequestFilenameExtension = RulesEngineMatchVariable("RequestFilenameExtension")
	RulesEngineMatchVariableRequestHeader            = RulesEngineMatchVariable("RequestHeader")
	RulesEngineMatchVariableRequestBody              = RulesEngineMatchVariable("RequestBody")
	RulesEngineMatchVariableRequestScheme            = RulesEngineMatchVariable("RequestScheme")
)

type RulesEngineOperator string

const (
	RulesEngineOperatorAny                = RulesEngineOperator("Any")
	RulesEngineOperatorIPMatch            = RulesEngineOperator("IPMatch")
	RulesEngineOperatorGeoMatch           = RulesEngineOperator("GeoMatch")
	RulesEngineOperatorEqual              = RulesEngineOperator("Equal")
	RulesEngineOperatorContains           = RulesEngineOperator("Contains")
	RulesEngineOperatorLessThan           = RulesEngineOperator("LessThan")
	RulesEngineOperatorGreaterThan        = RulesEngineOperator("GreaterThan")
	RulesEngineOperatorLessThanOrEqual    = RulesEngineOperator("LessThanOrEqual")
	RulesEngineOperatorGreaterThanOrEqual = RulesEngineOperator("GreaterThanOrEqual")
	RulesEngineOperatorBeginsWith         = RulesEngineOperator("BeginsWith")
	RulesEngineOperatorEndsWith           = RulesEngineOperator("EndsWith")
)

type SessionAffinityEnabledState string

const (
	SessionAffinityEnabledStateEnabled  = SessionAffinityEnabledState("Enabled")
	SessionAffinityEnabledStateDisabled = SessionAffinityEnabledState("Disabled")
)

type Transform string

const (
	TransformLowercase   = Transform("Lowercase")
	TransformUppercase   = Transform("Uppercase")
	TransformTrim        = Transform("Trim")
	TransformUrlDecode   = Transform("UrlDecode")
	TransformUrlEncode   = Transform("UrlEncode")
	TransformRemoveNulls = Transform("RemoveNulls")
)

func init() {
}
