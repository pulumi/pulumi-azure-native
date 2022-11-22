


package v20221001

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

type DaysOfWeek string

const (
	DaysOfWeekMonday    = DaysOfWeek("Monday")
	DaysOfWeekTuesday   = DaysOfWeek("Tuesday")
	DaysOfWeekWednesday = DaysOfWeek("Wednesday")
	DaysOfWeekThursday  = DaysOfWeek("Thursday")
	DaysOfWeekFriday    = DaysOfWeek("Friday")
	DaysOfWeekSaturday  = DaysOfWeek("Saturday")
	DaysOfWeekSunday    = DaysOfWeek("Sunday")
)

type ExportType string

const (
	ExportTypeUsage         = ExportType("Usage")
	ExportTypeActualCost    = ExportType("ActualCost")
	ExportTypeAmortizedCost = ExportType("AmortizedCost")
)

type FileFormat string

const (
	FileFormatCsv = FileFormat("Csv")
)

type FormatType string

const (
	FormatTypeCsv = FormatType("Csv")
)

type FunctionType string

const (
	FunctionTypeSum = FunctionType("Sum")
)

type GranularityType string

const (
	GranularityTypeDaily = GranularityType("Daily")
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

type RecurrenceType string

const (
	RecurrenceTypeDaily    = RecurrenceType("Daily")
	RecurrenceTypeWeekly   = RecurrenceType("Weekly")
	RecurrenceTypeMonthly  = RecurrenceType("Monthly")
	RecurrenceTypeAnnually = RecurrenceType("Annually")
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

type ScheduleFrequency string

const (
	// Cost analysis data will be emailed every day.
	ScheduleFrequencyDaily = ScheduleFrequency("Daily")
	// Cost analysis data will be emailed every week.
	ScheduleFrequencyWeekly = ScheduleFrequency("Weekly")
	// Cost analysis data will be emailed every month.
	ScheduleFrequencyMonthly = ScheduleFrequency("Monthly")
)

type ScheduledActionKind string

const (
	// Cost analysis data will be emailed.
	ScheduledActionKindEmail = ScheduledActionKind("Email")
	// Cost anomaly information will be emailed. Available only on subscription scope at daily frequency. If no anomaly is detected on the resource, an email won't be sent.
	ScheduledActionKindInsightAlert = ScheduledActionKind("InsightAlert")
)

type ScheduledActionStatus string

const (
	// Scheduled action is saved but will not be run.
	ScheduledActionStatusDisabled = ScheduledActionStatus("Disabled")
	// Scheduled action is saved and will be run.
	ScheduledActionStatusEnabled = ScheduledActionStatus("Enabled")
	// Scheduled action is expired.
	ScheduledActionStatusExpired = ScheduledActionStatus("Expired")
)

type StatusType string

const (
	StatusTypeActive   = StatusType("Active")
	StatusTypeInactive = StatusType("Inactive")
)

type TimeframeType string

const (
	TimeframeTypeMonthToDate         = TimeframeType("MonthToDate")
	TimeframeTypeBillingMonthToDate  = TimeframeType("BillingMonthToDate")
	TimeframeTypeTheLastMonth        = TimeframeType("TheLastMonth")
	TimeframeTypeTheLastBillingMonth = TimeframeType("TheLastBillingMonth")
	TimeframeTypeWeekToDate          = TimeframeType("WeekToDate")
	TimeframeTypeCustom              = TimeframeType("Custom")
)

type WeeksOfMonth string

const (
	WeeksOfMonthFirst  = WeeksOfMonth("First")
	WeeksOfMonthSecond = WeeksOfMonth("Second")
	WeeksOfMonthThird  = WeeksOfMonth("Third")
	WeeksOfMonthFourth = WeeksOfMonth("Fourth")
	WeeksOfMonthLast   = WeeksOfMonth("Last")
)

func init() {
}
