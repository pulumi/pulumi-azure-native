


package costmanagement

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

type ConnectorBillingModel string

const (
	ConnectorBillingModelTrial       = ConnectorBillingModel("trial")
	ConnectorBillingModelAutoUpgrade = ConnectorBillingModel("autoUpgrade")
	ConnectorBillingModelPremium     = ConnectorBillingModel("premium")
	ConnectorBillingModelExpired     = ConnectorBillingModel("expired")
)

type CostAllocationPolicyType string

const (
	CostAllocationPolicyTypeFixedProportion = CostAllocationPolicyType("FixedProportion")
)

type CostAllocationResourceType string

const (
	// Indicates an Azure dimension such as a subscription id or resource group name is being used for allocation.
	CostAllocationResourceTypeDimension = CostAllocationResourceType("Dimension")
	// Allocates cost based on Azure Tag key value pairs.
	CostAllocationResourceTypeTag = CostAllocationResourceType("Tag")
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
	FunctionTypeAvg = FunctionType("Avg")
	FunctionTypeMax = FunctionType("Max")
	FunctionTypeMin = FunctionType("Min")
	FunctionTypeSum = FunctionType("Sum")
)

type GranularityType string

const (
	GranularityTypeDaily  = GranularityType("Daily")
	GranularityTypeHourly = GranularityType("Hourly")
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

type ReportColumnType string

const (
	ReportColumnTypeTag       = ReportColumnType("Tag")
	ReportColumnTypeDimension = ReportColumnType("Dimension")
)

type ReportConfigColumnType string

const (
	ReportConfigColumnTypeTag       = ReportConfigColumnType("Tag")
	ReportConfigColumnTypeDimension = ReportConfigColumnType("Dimension")
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

type RuleStatus string

const (
	// Rule is saved but not used to allocate costs.
	RuleStatusNotActive = RuleStatus("NotActive")
	// Rule is saved and impacting cost allocation.
	RuleStatusActive = RuleStatus("Active")
	// Rule is saved and cost allocation is being updated. Readonly value that cannot be submitted in a put request.
	RuleStatusProcessing = RuleStatus("Processing")
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
)

type ScheduledActionStatus string

const (
	// Scheduled action is saved but will not be executed.
	ScheduledActionStatusDisabled = ScheduledActionStatus("Disabled")
	// Scheduled action is saved and will be executed.
	ScheduledActionStatusEnabled = ScheduledActionStatus("Enabled")
)

type StatusType string

const (
	StatusTypeActive   = StatusType("Active")
	StatusTypeInactive = StatusType("Inactive")
)

type TimeframeType string

const (
	TimeframeTypeWeekToDate  = TimeframeType("WeekToDate")
	TimeframeTypeMonthToDate = TimeframeType("MonthToDate")
	TimeframeTypeCustom      = TimeframeType("Custom")
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
