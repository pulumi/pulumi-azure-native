


package v20210808preview

type ActionType string

const (
	ActionTypeAddActionGroups       = ActionType("AddActionGroups")
	ActionTypeRemoveAllActionGroups = ActionType("RemoveAllActionGroups")
)

type DaysOfWeek string

const (
	DaysOfWeekSunday    = DaysOfWeek("Sunday")
	DaysOfWeekMonday    = DaysOfWeek("Monday")
	DaysOfWeekTuesday   = DaysOfWeek("Tuesday")
	DaysOfWeekWednesday = DaysOfWeek("Wednesday")
	DaysOfWeekThursday  = DaysOfWeek("Thursday")
	DaysOfWeekFriday    = DaysOfWeek("Friday")
	DaysOfWeekSaturday  = DaysOfWeek("Saturday")
)

type Field string

const (
	FieldSeverity            = Field("Severity")
	FieldMonitorService      = Field("MonitorService")
	FieldMonitorCondition    = Field("MonitorCondition")
	FieldSignalType          = Field("SignalType")
	FieldTargetResourceType  = Field("TargetResourceType")
	FieldTargetResource      = Field("TargetResource")
	FieldTargetResourceGroup = Field("TargetResourceGroup")
	FieldAlertRuleId         = Field("AlertRuleId")
	FieldAlertRuleName       = Field("AlertRuleName")
	FieldDescription         = Field("Description")
	FieldAlertContext        = Field("AlertContext")
)

type Operator string

const (
	OperatorEquals         = Operator("Equals")
	OperatorNotEquals      = Operator("NotEquals")
	OperatorContains       = Operator("Contains")
	OperatorDoesNotContain = Operator("DoesNotContain")
)

type RecurrenceType string

const (
	RecurrenceTypeDaily   = RecurrenceType("Daily")
	RecurrenceTypeWeekly  = RecurrenceType("Weekly")
	RecurrenceTypeMonthly = RecurrenceType("Monthly")
)

func init() {
}
