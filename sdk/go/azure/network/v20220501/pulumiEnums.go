


package v20220501

type ActionType string

const (
	ActionTypeAllow    = ActionType("Allow")
	ActionTypeBlock    = ActionType("Block")
	ActionTypeLog      = ActionType("Log")
	ActionTypeRedirect = ActionType("Redirect")
)

type CustomRuleEnabledState string

const (
	CustomRuleEnabledStateDisabled = CustomRuleEnabledState("Disabled")
	CustomRuleEnabledStateEnabled  = CustomRuleEnabledState("Enabled")
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

type RuleType string

const (
	RuleTypeMatchRule     = RuleType("MatchRule")
	RuleTypeRateLimitRule = RuleType("RateLimitRule")
)

type SkuName string

const (
	SkuName_Classic_AzureFrontDoor  = SkuName("Classic_AzureFrontDoor")
	SkuName_Standard_AzureFrontDoor = SkuName("Standard_AzureFrontDoor")
	SkuName_Premium_AzureFrontDoor  = SkuName("Premium_AzureFrontDoor")
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
