


package v20210201preview

type ConditionOperator string

const (
	ConditionOperatorEquals             = ConditionOperator("Equals")
	ConditionOperatorGreaterThan        = ConditionOperator("GreaterThan")
	ConditionOperatorGreaterThanOrEqual = ConditionOperator("GreaterThanOrEqual")
	ConditionOperatorLessThan           = ConditionOperator("LessThan")
	ConditionOperatorLessThanOrEqual    = ConditionOperator("LessThanOrEqual")
)

type DimensionOperator string

const (
	DimensionOperatorInclude = DimensionOperator("Include")
	DimensionOperatorExclude = DimensionOperator("Exclude")
)

type Kind string

const (
	KindLogAlert    = Kind("LogAlert")
	KindLogToMetric = Kind("LogToMetric")
)

type TimeAggregation string

const (
	TimeAggregationCount   = TimeAggregation("Count")
	TimeAggregationAverage = TimeAggregation("Average")
	TimeAggregationMinimum = TimeAggregation("Minimum")
	TimeAggregationMaximum = TimeAggregation("Maximum")
	TimeAggregationTotal   = TimeAggregation("Total")
)

func init() {
}
