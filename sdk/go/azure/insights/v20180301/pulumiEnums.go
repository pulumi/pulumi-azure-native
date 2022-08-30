


package v20180301

type AggregationTypeEnum string

const (
	AggregationTypeEnumAverage = AggregationTypeEnum("Average")
	AggregationTypeEnumCount   = AggregationTypeEnum("Count")
	AggregationTypeEnumMinimum = AggregationTypeEnum("Minimum")
	AggregationTypeEnumMaximum = AggregationTypeEnum("Maximum")
	AggregationTypeEnumTotal   = AggregationTypeEnum("Total")
)

type CriterionType string

const (
	CriterionTypeStaticThresholdCriterion  = CriterionType("StaticThresholdCriterion")
	CriterionTypeDynamicThresholdCriterion = CriterionType("DynamicThresholdCriterion")
)

type DynamicThresholdOperator string

const (
	DynamicThresholdOperatorGreaterThan       = DynamicThresholdOperator("GreaterThan")
	DynamicThresholdOperatorLessThan          = DynamicThresholdOperator("LessThan")
	DynamicThresholdOperatorGreaterOrLessThan = DynamicThresholdOperator("GreaterOrLessThan")
)

type DynamicThresholdSensitivity string

const (
	DynamicThresholdSensitivityLow    = DynamicThresholdSensitivity("Low")
	DynamicThresholdSensitivityMedium = DynamicThresholdSensitivity("Medium")
	DynamicThresholdSensitivityHigh   = DynamicThresholdSensitivity("High")
)

type Odatatype string

const (
	Odatatype_Microsoft_Azure_Monitor_SingleResourceMultipleMetricCriteria   = Odatatype("Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria")
	Odatatype_Microsoft_Azure_Monitor_MultipleResourceMultipleMetricCriteria = Odatatype("Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria")
	Odatatype_Microsoft_Azure_Monitor_WebtestLocationAvailabilityCriteria    = Odatatype("Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria")
)

type Operator string

const (
	OperatorEquals             = Operator("Equals")
	OperatorGreaterThan        = Operator("GreaterThan")
	OperatorGreaterThanOrEqual = Operator("GreaterThanOrEqual")
	OperatorLessThan           = Operator("LessThan")
	OperatorLessThanOrEqual    = Operator("LessThanOrEqual")
)

func init() {
}
