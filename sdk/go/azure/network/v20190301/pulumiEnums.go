


package v20190301

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

type RuleType string

const (
	RuleTypeMatchRule     = RuleType("MatchRule")
	RuleTypeRateLimitRule = RuleType("RateLimitRule")
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
