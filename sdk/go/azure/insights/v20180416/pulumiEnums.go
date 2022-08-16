


package v20180416

type AlertSeverity string

const (
	AlertSeverityZero  = AlertSeverity("0")
	AlertSeverityOne   = AlertSeverity("1")
	AlertSeverityTwo   = AlertSeverity("2")
	AlertSeverityThree = AlertSeverity("3")
	AlertSeverityFour  = AlertSeverity("4")
)

type ConditionalOperator string

const (
	ConditionalOperatorGreaterThanOrEqual = ConditionalOperator("GreaterThanOrEqual")
	ConditionalOperatorLessThanOrEqual    = ConditionalOperator("LessThanOrEqual")
	ConditionalOperatorGreaterThan        = ConditionalOperator("GreaterThan")
	ConditionalOperatorLessThan           = ConditionalOperator("LessThan")
	ConditionalOperatorEqual              = ConditionalOperator("Equal")
)

type Enabled string

const (
	EnabledTrue  = Enabled("true")
	EnabledFalse = Enabled("false")
)

type MetricTriggerType string

const (
	MetricTriggerTypeConsecutive = MetricTriggerType("Consecutive")
	MetricTriggerTypeTotal       = MetricTriggerType("Total")
)

type Operator string

const (
	OperatorInclude = Operator("Include")
)

type QueryType string

const (
	QueryTypeResultCount = QueryType("ResultCount")
)

func init() {
}
