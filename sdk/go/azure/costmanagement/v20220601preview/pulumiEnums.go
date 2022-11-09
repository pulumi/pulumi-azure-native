


package v20220601preview

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

type FileFormat string

const (
	FileFormatCsv = FileFormat("Csv")
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
	// Scheduled action is saved but will not be executed.
	ScheduledActionStatusDisabled = ScheduledActionStatus("Disabled")
	// Scheduled action is saved and will be executed.
	ScheduledActionStatusEnabled = ScheduledActionStatus("Enabled")
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
