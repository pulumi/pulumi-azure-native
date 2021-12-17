


package v20201201preview

type ExportType string

const (
	ExportTypeUsage         = ExportType("Usage")
	ExportTypeActualCost    = ExportType("ActualCost")
	ExportTypeAmortizedCost = ExportType("AmortizedCost")
)

type FormatType string

const (
	FormatTypeCsv = FormatType("Csv")
)

type GranularityType string

const (
	GranularityTypeDaily = GranularityType("Daily")
)

type RecurrenceType string

const (
	RecurrenceTypeDaily    = RecurrenceType("Daily")
	RecurrenceTypeWeekly   = RecurrenceType("Weekly")
	RecurrenceTypeMonthly  = RecurrenceType("Monthly")
	RecurrenceTypeAnnually = RecurrenceType("Annually")
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

func init() {
}
