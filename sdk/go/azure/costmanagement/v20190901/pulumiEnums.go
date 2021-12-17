


package v20190901

type ExportType string

const (
	ExportTypeUsage = ExportType("Usage")
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
	GranularityTypeDaily  = GranularityType("Daily")
	GranularityTypeHourly = GranularityType("Hourly")
)

type OperatorType string

const (
	OperatorTypeIn = OperatorType("In")
)

type QueryColumnType string

const (
	QueryColumnTypeTag       = QueryColumnType("Tag")
	QueryColumnTypeDimension = QueryColumnType("Dimension")
)

type RecurrenceType string

const (
	RecurrenceTypeDaily    = RecurrenceType("Daily")
	RecurrenceTypeWeekly   = RecurrenceType("Weekly")
	RecurrenceTypeMonthly  = RecurrenceType("Monthly")
	RecurrenceTypeAnnually = RecurrenceType("Annually")
)

type SortDirection string

const (
	SortDirectionAscending  = SortDirection("Ascending")
	SortDirectionDescending = SortDirection("Descending")
)

type StatusType string

const (
	StatusTypeActive   = StatusType("Active")
	StatusTypeInactive = StatusType("Inactive")
)

type TimeframeType string

const (
	TimeframeTypeWeekToDate   = TimeframeType("WeekToDate")
	TimeframeTypeMonthToDate  = TimeframeType("MonthToDate")
	TimeframeTypeYearToDate   = TimeframeType("YearToDate")
	TimeframeTypeTheLastWeek  = TimeframeType("TheLastWeek")
	TimeframeTypeTheLastMonth = TimeframeType("TheLastMonth")
	TimeframeTypeTheLastYear  = TimeframeType("TheLastYear")
	TimeframeTypeCustom       = TimeframeType("Custom")
)

func init() {
}
