


package v20160601

type DayOfWeek string

const (
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
)

type MonthOfYear string

const (
	MonthOfYearInvalid   = MonthOfYear("Invalid")
	MonthOfYearJanuary   = MonthOfYear("January")
	MonthOfYearFebruary  = MonthOfYear("February")
	MonthOfYearMarch     = MonthOfYear("March")
	MonthOfYearApril     = MonthOfYear("April")
	MonthOfYearMay       = MonthOfYear("May")
	MonthOfYearJune      = MonthOfYear("June")
	MonthOfYearJuly      = MonthOfYear("July")
	MonthOfYearAugust    = MonthOfYear("August")
	MonthOfYearSeptember = MonthOfYear("September")
	MonthOfYearOctober   = MonthOfYear("October")
	MonthOfYearNovember  = MonthOfYear("November")
	MonthOfYearDecember  = MonthOfYear("December")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
)

type RetentionDurationType string

const (
	RetentionDurationTypeInvalid = RetentionDurationType("Invalid")
	RetentionDurationTypeDays    = RetentionDurationType("Days")
	RetentionDurationTypeWeeks   = RetentionDurationType("Weeks")
	RetentionDurationTypeMonths  = RetentionDurationType("Months")
	RetentionDurationTypeYears   = RetentionDurationType("Years")
)

type RetentionScheduleFormat string

const (
	RetentionScheduleFormatInvalid = RetentionScheduleFormat("Invalid")
	RetentionScheduleFormatDaily   = RetentionScheduleFormat("Daily")
	RetentionScheduleFormatWeekly  = RetentionScheduleFormat("Weekly")
)

type ScheduleRunType string

const (
	ScheduleRunTypeInvalid = ScheduleRunType("Invalid")
	ScheduleRunTypeDaily   = ScheduleRunType("Daily")
	ScheduleRunTypeWeekly  = ScheduleRunType("Weekly")
)

type SkuName string

const (
	SkuNameStandard = SkuName("Standard")
	SkuNameRS0      = SkuName("RS0")
)

type WeekOfMonth string

const (
	WeekOfMonthFirst  = WeekOfMonth("First")
	WeekOfMonthSecond = WeekOfMonth("Second")
	WeekOfMonthThird  = WeekOfMonth("Third")
	WeekOfMonthFourth = WeekOfMonth("Fourth")
	WeekOfMonthLast   = WeekOfMonth("Last")
)

func init() {
}
