


package v20221101preview

type AbsoluteMarker string

const (
	AbsoluteMarkerAllBackup    = AbsoluteMarker("AllBackup")
	AbsoluteMarkerFirstOfDay   = AbsoluteMarker("FirstOfDay")
	AbsoluteMarkerFirstOfMonth = AbsoluteMarker("FirstOfMonth")
	AbsoluteMarkerFirstOfWeek  = AbsoluteMarker("FirstOfWeek")
	AbsoluteMarkerFirstOfYear  = AbsoluteMarker("FirstOfYear")
)

type AlertsState string

const (
	AlertsStateEnabled  = AlertsState("Enabled")
	AlertsStateDisabled = AlertsState("Disabled")
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

type ImmutabilityState string

const (
	ImmutabilityStateDisabled = ImmutabilityState("Disabled")
	ImmutabilityStateUnlocked = ImmutabilityState("Unlocked")
	ImmutabilityStateLocked   = ImmutabilityState("Locked")
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

type SoftDeleteState string

const (
	// Soft Delete is turned off for the BackupVault
	SoftDeleteStateOff = SoftDeleteState("Off")
	// Soft Delete is enabled for the BackupVault but can be turned off
	SoftDeleteStateOn = SoftDeleteState("On")
	// Soft Delete is permanently enabled for the BackupVault and the setting cannot be changed
	SoftDeleteStateAlwaysOn = SoftDeleteState("AlwaysOn")
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
	StorageSettingTypesZoneRedundant    = StorageSettingTypes("ZoneRedundant")
)

type ValidationType string

const (
	ValidationTypeShallowValidation = ValidationType("ShallowValidation")
	ValidationTypeDeepValidation    = ValidationType("DeepValidation")
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
