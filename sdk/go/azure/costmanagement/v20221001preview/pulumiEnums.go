


package v20221001preview

type AccumulatedType string

const (
	AccumulatedTypeTrue  = AccumulatedType("true")
	AccumulatedTypeFalse = AccumulatedType("false")
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

type ReportConfigSortingType string

const (
	ReportConfigSortingTypeAscending  = ReportConfigSortingType("Ascending")
	ReportConfigSortingTypeDescending = ReportConfigSortingType("Descending")
)

type ReportGranularityType string

const (
	ReportGranularityTypeDaily   = ReportGranularityType("Daily")
	ReportGranularityTypeMonthly = ReportGranularityType("Monthly")
)

type ReportTimeframeType string

const (
	ReportTimeframeTypeWeekToDate  = ReportTimeframeType("WeekToDate")
	ReportTimeframeTypeMonthToDate = ReportTimeframeType("MonthToDate")
	ReportTimeframeTypeYearToDate  = ReportTimeframeType("YearToDate")
	ReportTimeframeTypeCustom      = ReportTimeframeType("Custom")
)

type ReportType string

const (
	ReportTypeUsage = ReportType("Usage")
)

type SettingsKind string

const (
	SettingsKindTaginheritance = SettingsKind("taginheritance")
)

func init() {
}
