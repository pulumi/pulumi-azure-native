


package v20151031

type ContentSourceType string

const (
	ContentSourceTypeEmbeddedContent = ContentSourceType("embeddedContent")
	ContentSourceTypeUri             = ContentSourceType("uri")
)

type RunbookTypeEnum string

const (
	RunbookTypeEnumScript                  = RunbookTypeEnum("Script")
	RunbookTypeEnumGraph                   = RunbookTypeEnum("Graph")
	RunbookTypeEnumPowerShellWorkflow      = RunbookTypeEnum("PowerShellWorkflow")
	RunbookTypeEnumPowerShell              = RunbookTypeEnum("PowerShell")
	RunbookTypeEnumGraphPowerShellWorkflow = RunbookTypeEnum("GraphPowerShellWorkflow")
	RunbookTypeEnumGraphPowerShell         = RunbookTypeEnum("GraphPowerShell")
)

type ScheduleDay string

const (
	ScheduleDayMonday    = ScheduleDay("Monday")
	ScheduleDayTuesday   = ScheduleDay("Tuesday")
	ScheduleDayWednesday = ScheduleDay("Wednesday")
	ScheduleDayThursday  = ScheduleDay("Thursday")
	ScheduleDayFriday    = ScheduleDay("Friday")
	ScheduleDaySaturday  = ScheduleDay("Saturday")
	ScheduleDaySunday    = ScheduleDay("Sunday")
)

type ScheduleFrequency string

const (
	ScheduleFrequencyOneTime = ScheduleFrequency("OneTime")
	ScheduleFrequencyDay     = ScheduleFrequency("Day")
	ScheduleFrequencyHour    = ScheduleFrequency("Hour")
	ScheduleFrequencyWeek    = ScheduleFrequency("Week")
	ScheduleFrequencyMonth   = ScheduleFrequency("Month")
	// The minimum allowed interval for Minute schedules is 15 minutes.
	ScheduleFrequencyMinute = ScheduleFrequency("Minute")
)

type SkuNameEnum string

const (
	SkuNameEnumFree  = SkuNameEnum("Free")
	SkuNameEnumBasic = SkuNameEnum("Basic")
)

func init() {
}
