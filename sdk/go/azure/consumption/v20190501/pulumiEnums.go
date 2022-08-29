


package v20190501

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

type ThresholdType string

const (
	ThresholdTypeActual = ThresholdType("Actual")
)

type TimeGrainType string

const (
	TimeGrainTypeMonthly        = TimeGrainType("Monthly")
	TimeGrainTypeQuarterly      = TimeGrainType("Quarterly")
	TimeGrainTypeAnnually       = TimeGrainType("Annually")
	TimeGrainTypeBillingMonth   = TimeGrainType("BillingMonth")
	TimeGrainTypeBillingQuarter = TimeGrainType("BillingQuarter")
	TimeGrainTypeBillingAnnual  = TimeGrainType("BillingAnnual")
)

func init() {
}
