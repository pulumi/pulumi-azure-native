


package v20190505preview

type ActionRuleStatus string

const (
	ActionRuleStatusEnabled  = ActionRuleStatus("Enabled")
	ActionRuleStatusDisabled = ActionRuleStatus("Disabled")
)

type ActionRuleType string

const (
	ActionRuleTypeSuppression = ActionRuleType("Suppression")
	ActionRuleTypeActionGroup = ActionRuleType("ActionGroup")
	ActionRuleTypeDiagnostics = ActionRuleType("Diagnostics")
)

type Operator string

const (
	OperatorEquals         = Operator("Equals")
	OperatorNotEquals      = Operator("NotEquals")
	OperatorContains       = Operator("Contains")
	OperatorDoesNotContain = Operator("DoesNotContain")
)

type ScopeType string

const (
	ScopeTypeResourceGroup = ScopeType("ResourceGroup")
	ScopeTypeResource      = ScopeType("Resource")
	ScopeTypeSubscription  = ScopeType("Subscription")
)

type SuppressionType string

const (
	SuppressionTypeAlways  = SuppressionType("Always")
	SuppressionTypeOnce    = SuppressionType("Once")
	SuppressionTypeDaily   = SuppressionType("Daily")
	SuppressionTypeWeekly  = SuppressionType("Weekly")
	SuppressionTypeMonthly = SuppressionType("Monthly")
)

func init() {
}
