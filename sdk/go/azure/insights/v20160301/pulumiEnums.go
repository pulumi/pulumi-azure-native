


package v20160301

type ConditionOperator string

const (
	ConditionOperatorGreaterThan        = ConditionOperator("GreaterThan")
	ConditionOperatorGreaterThanOrEqual = ConditionOperator("GreaterThanOrEqual")
	ConditionOperatorLessThan           = ConditionOperator("LessThan")
	ConditionOperatorLessThanOrEqual    = ConditionOperator("LessThanOrEqual")
)

type TimeAggregationOperator string

const (
	TimeAggregationOperatorAverage = TimeAggregationOperator("Average")
	TimeAggregationOperatorMinimum = TimeAggregationOperator("Minimum")
	TimeAggregationOperatorMaximum = TimeAggregationOperator("Maximum")
	TimeAggregationOperatorTotal   = TimeAggregationOperator("Total")
	TimeAggregationOperatorLast    = TimeAggregationOperator("Last")
)

func init() {
}
