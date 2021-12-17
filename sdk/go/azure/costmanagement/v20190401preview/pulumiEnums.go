


package v20190401preview

type AccumulatedType string

const (
	AccumulatedTypeTrue  = AccumulatedType("true")
	AccumulatedTypeFalse = AccumulatedType("false")
)

type CategoryType string

const (
	CategoryTypeCost  = CategoryType("Cost")
	CategoryTypeUsage = CategoryType("Usage")
)

type ChartType string

const (
	ChartTypeArea          = ChartType("Area")
	ChartTypeLine          = ChartType("Line")
	ChartTypeStackedColumn = ChartType("StackedColumn")
	ChartTypeGroupedColumn = ChartType("GroupedColumn")
	ChartTypeTable         = ChartType("Table")
)

type FunctionType string

const (
	FunctionTypeSum = FunctionType("Sum")
)

type GranularityType string

const (
	GranularityTypeDaily   = GranularityType("Daily")
	GranularityTypeMonthly = GranularityType("Monthly")
)

type KpiTypeType string

const (
	KpiTypeTypeForecast = KpiTypeType("Forecast")
	KpiTypeTypeBudget   = KpiTypeType("Budget")
)

type MetricType string

const (
	MetricTypeActualCost    = MetricType("ActualCost")
	MetricTypeAmortizedCost = MetricType("AmortizedCost")
	MetricTypeAHUB          = MetricType("AHUB")
)

type NotificationOperatorType string

const (
	NotificationOperatorTypeEqualTo              = NotificationOperatorType("EqualTo")
	NotificationOperatorTypeGreaterThan          = NotificationOperatorType("GreaterThan")
	NotificationOperatorTypeGreaterThanOrEqualTo = NotificationOperatorType("GreaterThanOrEqualTo")
)

type OperatorType string

const (
	OperatorTypeIn       = OperatorType("In")
	OperatorTypeContains = OperatorType("Contains")
)

type PivotTypeType string

const (
	PivotTypeTypeDimension = PivotTypeType("Dimension")
	PivotTypeTypeTagKey    = PivotTypeType("TagKey")
)

type ReportConfigColumnType string

const (
	ReportConfigColumnTypeTag       = ReportConfigColumnType("Tag")
	ReportConfigColumnTypeDimension = ReportConfigColumnType("Dimension")
)

type ReportType string

const (
	ReportTypeUsage = ReportType("Usage")
)

type TimeGrainType string

const (
	TimeGrainTypeMonthly   = TimeGrainType("Monthly")
	TimeGrainTypeQuarterly = TimeGrainType("Quarterly")
	TimeGrainTypeAnnually  = TimeGrainType("Annually")
)

type TimeframeType string

const (
	TimeframeTypeWeekToDate  = TimeframeType("WeekToDate")
	TimeframeTypeMonthToDate = TimeframeType("MonthToDate")
	TimeframeTypeYearToDate  = TimeframeType("YearToDate")
	TimeframeTypeCustom      = TimeframeType("Custom")
)

func init() {
}
