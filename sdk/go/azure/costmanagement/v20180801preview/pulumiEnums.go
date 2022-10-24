


package v20180801preview

type ConnectorStatus string

const (
	ConnectorStatusActive    = ConnectorStatus("active")
	ConnectorStatusError     = ConnectorStatus("error")
	ConnectorStatusSuspended = ConnectorStatus("suspended")
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

type ReportType string

const (
	ReportTypeUsage = ReportType("Usage")
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

func init() {
}
