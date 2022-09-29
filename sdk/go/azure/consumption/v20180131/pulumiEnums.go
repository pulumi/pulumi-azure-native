


package v20180131

type CategoryType string

const (
	CategoryTypeCost  = CategoryType("Cost")
	CategoryTypeUsage = CategoryType("Usage")
)

type OperatorType string

const (
	OperatorTypeEqualTo              = OperatorType("EqualTo")
	OperatorTypeGreaterThan          = OperatorType("GreaterThan")
	OperatorTypeGreaterThanOrEqualTo = OperatorType("GreaterThanOrEqualTo")
)

type TimeGrainType string

const (
	TimeGrainTypeMonthly   = TimeGrainType("Monthly")
	TimeGrainTypeQuarterly = TimeGrainType("Quarterly")
	TimeGrainTypeAnnually  = TimeGrainType("Annually")
)

func init() {
}
