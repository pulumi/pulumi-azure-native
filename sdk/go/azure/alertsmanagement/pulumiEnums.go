


package alertsmanagement

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

type AlertRuleState string

const (
	AlertRuleStateEnabled  = AlertRuleState("Enabled")
	AlertRuleStateDisabled = AlertRuleState("Disabled")
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

type Severity string

const (
	SeveritySev0 = Severity("Sev0")
	SeveritySev1 = Severity("Sev1")
	SeveritySev2 = Severity("Sev2")
	SeveritySev3 = Severity("Sev3")
	SeveritySev4 = Severity("Sev4")
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
