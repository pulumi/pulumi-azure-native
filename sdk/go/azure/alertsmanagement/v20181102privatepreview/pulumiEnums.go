


package v20181102privatepreview

type ActionRuleStatus string

const (
	ActionRuleStatusEnabled  = ActionRuleStatus("enabled")
	ActionRuleStatusDisabled = ActionRuleStatus("disabled")
)

type ScopeType string

const (
	ScopeTypeResourceGroup = ScopeType("ResourceGroup")
	ScopeTypeResource      = ScopeType("Resource")
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
