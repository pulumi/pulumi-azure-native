


package v20211201preview

type AbsoluteMarker string

const (
	AbsoluteMarkerAllBackup    = AbsoluteMarker("AllBackup")
	AbsoluteMarkerFirstOfDay   = AbsoluteMarker("FirstOfDay")
	AbsoluteMarkerFirstOfMonth = AbsoluteMarker("FirstOfMonth")
	AbsoluteMarkerFirstOfWeek  = AbsoluteMarker("FirstOfWeek")
	AbsoluteMarkerFirstOfYear  = AbsoluteMarker("FirstOfYear")
)

type DataStoreTypes string

const (
	DataStoreTypesOperationalStore = DataStoreTypes("OperationalStore")
	DataStoreTypesVaultStore       = DataStoreTypes("VaultStore")
	DataStoreTypesArchiveStore     = DataStoreTypes("ArchiveStore")
)

type DayOfWeek string

const (
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
)

type Month string

const (
	MonthApril     = Month("April")
	MonthAugust    = Month("August")
	MonthDecember  = Month("December")
	MonthFebruary  = Month("February")
	MonthJanuary   = Month("January")
	MonthJuly      = Month("July")
	MonthJune      = Month("June")
	MonthMarch     = Month("March")
	MonthMay       = Month("May")
	MonthNovember  = Month("November")
	MonthOctober   = Month("October")
	MonthSeptember = Month("September")
)

type SecretStoreType string

const (
	SecretStoreTypeInvalid       = SecretStoreType("Invalid")
	SecretStoreTypeAzureKeyVault = SecretStoreType("AzureKeyVault")
)

type StorageSettingStoreTypes string

const (
	StorageSettingStoreTypesArchiveStore  = StorageSettingStoreTypes("ArchiveStore")
	StorageSettingStoreTypesSnapshotStore = StorageSettingStoreTypes("SnapshotStore")
	StorageSettingStoreTypesVaultStore    = StorageSettingStoreTypes("VaultStore")
)

type StorageSettingTypes string

const (
	StorageSettingTypesGeoRedundant     = StorageSettingTypes("GeoRedundant")
	StorageSettingTypesLocallyRedundant = StorageSettingTypes("LocallyRedundant")
)

type WeekNumber string

const (
	WeekNumberFirst  = WeekNumber("First")
	WeekNumberFourth = WeekNumber("Fourth")
	WeekNumberLast   = WeekNumber("Last")
	WeekNumberSecond = WeekNumber("Second")
	WeekNumberThird  = WeekNumber("Third")
)

func init() {
}
