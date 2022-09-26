


package v20210601

type ActionType string

const (
	ActionTypeAllow    = ActionType("Allow")
	ActionTypeBlock    = ActionType("Block")
	ActionTypeLog      = ActionType("Log")
	ActionTypeRedirect = ActionType("Redirect")
)

type BackendEnabledState string

const (
	BackendEnabledStateEnabled  = BackendEnabledState("Enabled")
	BackendEnabledStateDisabled = BackendEnabledState("Disabled")
)

type CustomRuleEnabledState string

const (
	CustomRuleEnabledStateDisabled = CustomRuleEnabledState("Disabled")
	CustomRuleEnabledStateEnabled  = CustomRuleEnabledState("Enabled")
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

type FrontDoorMatchVariable string

const (
	FrontDoorMatchVariableRemoteAddr    = FrontDoorMatchVariable("RemoteAddr")
	FrontDoorMatchVariableRequestMethod = FrontDoorMatchVariable("RequestMethod")
	FrontDoorMatchVariableQueryString   = FrontDoorMatchVariable("QueryString")
	FrontDoorMatchVariablePostArgs      = FrontDoorMatchVariable("PostArgs")
	FrontDoorMatchVariableRequestUri    = FrontDoorMatchVariable("RequestUri")
	FrontDoorMatchVariableRequestHeader = FrontDoorMatchVariable("RequestHeader")
	FrontDoorMatchVariableRequestBody   = FrontDoorMatchVariable("RequestBody")
	FrontDoorMatchVariableCookies       = FrontDoorMatchVariable("Cookies")
	FrontDoorMatchVariableSocketAddr    = FrontDoorMatchVariable("SocketAddr")
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

type ManagedRuleEnabledState string

const (
	ManagedRuleEnabledStateDisabled = ManagedRuleEnabledState("Disabled")
	ManagedRuleEnabledStateEnabled  = ManagedRuleEnabledState("Enabled")
)

type ManagedRuleExclusionMatchVariable string

const (
	ManagedRuleExclusionMatchVariableRequestHeaderNames      = ManagedRuleExclusionMatchVariable("RequestHeaderNames")
	ManagedRuleExclusionMatchVariableRequestCookieNames      = ManagedRuleExclusionMatchVariable("RequestCookieNames")
	ManagedRuleExclusionMatchVariableQueryStringArgNames     = ManagedRuleExclusionMatchVariable("QueryStringArgNames")
	ManagedRuleExclusionMatchVariableRequestBodyPostArgNames = ManagedRuleExclusionMatchVariable("RequestBodyPostArgNames")
	ManagedRuleExclusionMatchVariableRequestBodyJsonArgNames = ManagedRuleExclusionMatchVariable("RequestBodyJsonArgNames")
)

type ManagedRuleExclusionSelectorMatchOperator string

const (
	ManagedRuleExclusionSelectorMatchOperatorEquals     = ManagedRuleExclusionSelectorMatchOperator("Equals")
	ManagedRuleExclusionSelectorMatchOperatorContains   = ManagedRuleExclusionSelectorMatchOperator("Contains")
	ManagedRuleExclusionSelectorMatchOperatorStartsWith = ManagedRuleExclusionSelectorMatchOperator("StartsWith")
	ManagedRuleExclusionSelectorMatchOperatorEndsWith   = ManagedRuleExclusionSelectorMatchOperator("EndsWith")
	ManagedRuleExclusionSelectorMatchOperatorEqualsAny  = ManagedRuleExclusionSelectorMatchOperator("EqualsAny")
)

type ManagedRuleSetActionType string

const (
	ManagedRuleSetActionTypeBlock    = ManagedRuleSetActionType("Block")
	ManagedRuleSetActionTypeLog      = ManagedRuleSetActionType("Log")
	ManagedRuleSetActionTypeRedirect = ManagedRuleSetActionType("Redirect")
)

type MatchProcessingBehavior string

const (
	MatchProcessingBehaviorContinue = MatchProcessingBehavior("Continue")
	MatchProcessingBehaviorStop     = MatchProcessingBehavior("Stop")
)

type Operator string

const (
	OperatorAny                = Operator("Any")
	OperatorIPMatch            = Operator("IPMatch")
	OperatorGeoMatch           = Operator("GeoMatch")
	OperatorEqual              = Operator("Equal")
	OperatorContains           = Operator("Contains")
	OperatorLessThan           = Operator("LessThan")
	OperatorGreaterThan        = Operator("GreaterThan")
	OperatorLessThanOrEqual    = Operator("LessThanOrEqual")
	OperatorGreaterThanOrEqual = Operator("GreaterThanOrEqual")
	OperatorBeginsWith         = Operator("BeginsWith")
	OperatorEndsWith           = Operator("EndsWith")
	OperatorRegEx              = Operator("RegEx")
)

type PolicyEnabledState string

const (
	PolicyEnabledStateDisabled = PolicyEnabledState("Disabled")
	PolicyEnabledStateEnabled  = PolicyEnabledState("Enabled")
)

type PolicyMode string

const (
	PolicyModePrevention = PolicyMode("Prevention")
	PolicyModeDetection  = PolicyMode("Detection")
)

type PolicyRequestBodyCheck string

const (
	PolicyRequestBodyCheckDisabled = PolicyRequestBodyCheck("Disabled")
	PolicyRequestBodyCheckEnabled  = PolicyRequestBodyCheck("Enabled")
)

type RoutingRuleEnabledState string

const (
	RoutingRuleEnabledStateEnabled  = RoutingRuleEnabledState("Enabled")
	RoutingRuleEnabledStateDisabled = RoutingRuleEnabledState("Disabled")
)

type RuleType string

const (
	RuleTypeMatchRule     = RuleType("MatchRule")
	RuleTypeRateLimitRule = RuleType("RateLimitRule")
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

type SkuName string

const (
	SkuName_Classic_AzureFrontDoor  = SkuName("Classic_AzureFrontDoor")
	SkuName_Standard_AzureFrontDoor = SkuName("Standard_AzureFrontDoor")
	SkuName_Premium_AzureFrontDoor  = SkuName("Premium_AzureFrontDoor")
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

type TransformType string

const (
	TransformTypeLowercase   = TransformType("Lowercase")
	TransformTypeUppercase   = TransformType("Uppercase")
	TransformTypeTrim        = TransformType("Trim")
	TransformTypeUrlDecode   = TransformType("UrlDecode")
	TransformTypeUrlEncode   = TransformType("UrlEncode")
	TransformTypeRemoveNulls = TransformType("RemoveNulls")
)

func init() {
}
