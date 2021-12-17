


package v20210201preview

type BackupItemType string

const (
	BackupItemTypeInvalid           = BackupItemType("Invalid")
	BackupItemTypeVM                = BackupItemType("VM")
	BackupItemTypeFileFolder        = BackupItemType("FileFolder")
	BackupItemTypeAzureSqlDb        = BackupItemType("AzureSqlDb")
	BackupItemTypeSQLDB             = BackupItemType("SQLDB")
	BackupItemTypeExchange          = BackupItemType("Exchange")
	BackupItemTypeSharepoint        = BackupItemType("Sharepoint")
	BackupItemTypeVMwareVM          = BackupItemType("VMwareVM")
	BackupItemTypeSystemState       = BackupItemType("SystemState")
	BackupItemTypeClient            = BackupItemType("Client")
	BackupItemTypeGenericDataSource = BackupItemType("GenericDataSource")
	BackupItemTypeSQLDataBase       = BackupItemType("SQLDataBase")
	BackupItemTypeAzureFileShare    = BackupItemType("AzureFileShare")
	BackupItemTypeSAPHanaDatabase   = BackupItemType("SAPHanaDatabase")
	BackupItemTypeSAPAseDatabase    = BackupItemType("SAPAseDatabase")
)

type BackupManagementType string

const (
	BackupManagementTypeInvalid           = BackupManagementType("Invalid")
	BackupManagementTypeAzureIaasVM       = BackupManagementType("AzureIaasVM")
	BackupManagementTypeMAB               = BackupManagementType("MAB")
	BackupManagementTypeDPM               = BackupManagementType("DPM")
	BackupManagementTypeAzureBackupServer = BackupManagementType("AzureBackupServer")
	BackupManagementTypeAzureSql          = BackupManagementType("AzureSql")
	BackupManagementTypeAzureStorage      = BackupManagementType("AzureStorage")
	BackupManagementTypeAzureWorkload     = BackupManagementType("AzureWorkload")
	BackupManagementTypeDefaultBackup     = BackupManagementType("DefaultBackup")
)

type ContainerType string

const (
	ContainerTypeInvalid                    = ContainerType("Invalid")
	ContainerTypeUnknown                    = ContainerType("Unknown")
	ContainerTypeIaasVMContainer            = ContainerType("IaasVMContainer")
	ContainerTypeIaasVMServiceContainer     = ContainerType("IaasVMServiceContainer")
	ContainerTypeDPMContainer               = ContainerType("DPMContainer")
	ContainerTypeAzureBackupServerContainer = ContainerType("AzureBackupServerContainer")
	ContainerTypeMABContainer               = ContainerType("MABContainer")
	ContainerTypeCluster                    = ContainerType("Cluster")
	ContainerTypeAzureSqlContainer          = ContainerType("AzureSqlContainer")
	ContainerTypeWindows                    = ContainerType("Windows")
	ContainerTypeVCenter                    = ContainerType("VCenter")
	ContainerTypeVMAppContainer             = ContainerType("VMAppContainer")
	ContainerTypeSQLAGWorkLoadContainer     = ContainerType("SQLAGWorkLoadContainer")
	ContainerTypeStorageContainer           = ContainerType("StorageContainer")
	ContainerTypeGenericContainer           = ContainerType("GenericContainer")
)

type CreateMode string

const (
	CreateModeInvalid = CreateMode("Invalid")
	CreateModeDefault = CreateMode("Default")
	CreateModeRecover = CreateMode("Recover")
)

type DataSourceType string

const (
	DataSourceTypeInvalid           = DataSourceType("Invalid")
	DataSourceTypeVM                = DataSourceType("VM")
	DataSourceTypeFileFolder        = DataSourceType("FileFolder")
	DataSourceTypeAzureSqlDb        = DataSourceType("AzureSqlDb")
	DataSourceTypeSQLDB             = DataSourceType("SQLDB")
	DataSourceTypeExchange          = DataSourceType("Exchange")
	DataSourceTypeSharepoint        = DataSourceType("Sharepoint")
	DataSourceTypeVMwareVM          = DataSourceType("VMwareVM")
	DataSourceTypeSystemState       = DataSourceType("SystemState")
	DataSourceTypeClient            = DataSourceType("Client")
	DataSourceTypeGenericDataSource = DataSourceType("GenericDataSource")
	DataSourceTypeSQLDataBase       = DataSourceType("SQLDataBase")
	DataSourceTypeAzureFileShare    = DataSourceType("AzureFileShare")
	DataSourceTypeSAPHanaDatabase   = DataSourceType("SAPHanaDatabase")
	DataSourceTypeSAPAseDatabase    = DataSourceType("SAPAseDatabase")
)

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

type HealthStatus string

const (
	HealthStatusPassed          = HealthStatus("Passed")
	HealthStatusActionRequired  = HealthStatus("ActionRequired")
	HealthStatusActionSuggested = HealthStatus("ActionSuggested")
	HealthStatusInvalid         = HealthStatus("Invalid")
)

type LastBackupStatus string

const (
	LastBackupStatusInvalid   = LastBackupStatus("Invalid")
	LastBackupStatusHealthy   = LastBackupStatus("Healthy")
	LastBackupStatusUnhealthy = LastBackupStatus("Unhealthy")
	LastBackupStatusIRPending = LastBackupStatus("IRPending")
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

type OperationType string

const (
	OperationTypeInvalid    = OperationType("Invalid")
	OperationTypeRegister   = OperationType("Register")
	OperationTypeReregister = OperationType("Reregister")
)

type PolicyType string

const (
	PolicyTypeInvalid      = PolicyType("Invalid")
	PolicyTypeFull         = PolicyType("Full")
	PolicyTypeDifferential = PolicyType("Differential")
	PolicyTypeLog          = PolicyType("Log")
	PolicyTypeCopyOnlyFull = PolicyType("CopyOnlyFull")
	PolicyTypeIncremental  = PolicyType("Incremental")
)

type PrivateEndpointConnectionStatus string

const (
	PrivateEndpointConnectionStatusPending      = PrivateEndpointConnectionStatus("Pending")
	PrivateEndpointConnectionStatusApproved     = PrivateEndpointConnectionStatus("Approved")
	PrivateEndpointConnectionStatusRejected     = PrivateEndpointConnectionStatus("Rejected")
	PrivateEndpointConnectionStatusDisconnected = PrivateEndpointConnectionStatus("Disconnected")
)

type ProtectedItemHealthStatus string

const (
	ProtectedItemHealthStatusInvalid      = ProtectedItemHealthStatus("Invalid")
	ProtectedItemHealthStatusHealthy      = ProtectedItemHealthStatus("Healthy")
	ProtectedItemHealthStatusUnhealthy    = ProtectedItemHealthStatus("Unhealthy")
	ProtectedItemHealthStatusNotReachable = ProtectedItemHealthStatus("NotReachable")
	ProtectedItemHealthStatusIRPending    = ProtectedItemHealthStatus("IRPending")
)

type ProtectedItemStateEnum string

const (
	ProtectedItemStateEnumInvalid           = ProtectedItemStateEnum("Invalid")
	ProtectedItemStateEnumIRPending         = ProtectedItemStateEnum("IRPending")
	ProtectedItemStateEnumProtected         = ProtectedItemStateEnum("Protected")
	ProtectedItemStateEnumProtectionError   = ProtectedItemStateEnum("ProtectionError")
	ProtectedItemStateEnumProtectionStopped = ProtectedItemStateEnum("ProtectionStopped")
	ProtectedItemStateEnumProtectionPaused  = ProtectedItemStateEnum("ProtectionPaused")
)

type ProtectionState string

const (
	ProtectionStateInvalid           = ProtectionState("Invalid")
	ProtectionStateIRPending         = ProtectionState("IRPending")
	ProtectionStateProtected         = ProtectionState("Protected")
	ProtectionStateProtectionError   = ProtectionState("ProtectionError")
	ProtectionStateProtectionStopped = ProtectionState("ProtectionStopped")
	ProtectionStateProtectionPaused  = ProtectionState("ProtectionPaused")
)

type ProtectionStatus string

const (
	ProtectionStatusInvalid          = ProtectionStatus("Invalid")
	ProtectionStatusNotProtected     = ProtectionStatus("NotProtected")
	ProtectionStatusProtecting       = ProtectionStatus("Protecting")
	ProtectionStatusProtected        = ProtectionStatus("Protected")
	ProtectionStatusProtectionFailed = ProtectionStatus("ProtectionFailed")
)

type ProvisioningState string

const (
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
	ProvisioningStateDeleting  = ProvisioningState("Deleting")
	ProvisioningStateFailed    = ProvisioningState("Failed")
	ProvisioningStatePending   = ProvisioningState("Pending")
)

type ResourceHealthStatus string

const (
	ResourceHealthStatusHealthy             = ResourceHealthStatus("Healthy")
	ResourceHealthStatusTransientDegraded   = ResourceHealthStatus("TransientDegraded")
	ResourceHealthStatusPersistentDegraded  = ResourceHealthStatus("PersistentDegraded")
	ResourceHealthStatusTransientUnhealthy  = ResourceHealthStatus("TransientUnhealthy")
	ResourceHealthStatusPersistentUnhealthy = ResourceHealthStatus("PersistentUnhealthy")
	ResourceHealthStatusInvalid             = ResourceHealthStatus("Invalid")
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

type WeekOfMonth string

const (
	WeekOfMonthFirst   = WeekOfMonth("First")
	WeekOfMonthSecond  = WeekOfMonth("Second")
	WeekOfMonthThird   = WeekOfMonth("Third")
	WeekOfMonthFourth  = WeekOfMonth("Fourth")
	WeekOfMonthLast    = WeekOfMonth("Last")
	WeekOfMonthInvalid = WeekOfMonth("Invalid")
)

type WorkloadItemType string

const (
	WorkloadItemTypeInvalid         = WorkloadItemType("Invalid")
	WorkloadItemTypeSQLInstance     = WorkloadItemType("SQLInstance")
	WorkloadItemTypeSQLDataBase     = WorkloadItemType("SQLDataBase")
	WorkloadItemTypeSAPHanaSystem   = WorkloadItemType("SAPHanaSystem")
	WorkloadItemTypeSAPHanaDatabase = WorkloadItemType("SAPHanaDatabase")
	WorkloadItemTypeSAPAseSystem    = WorkloadItemType("SAPAseSystem")
	WorkloadItemTypeSAPAseDatabase  = WorkloadItemType("SAPAseDatabase")
)

type WorkloadType string

const (
	WorkloadTypeInvalid           = WorkloadType("Invalid")
	WorkloadTypeVM                = WorkloadType("VM")
	WorkloadTypeFileFolder        = WorkloadType("FileFolder")
	WorkloadTypeAzureSqlDb        = WorkloadType("AzureSqlDb")
	WorkloadTypeSQLDB             = WorkloadType("SQLDB")
	WorkloadTypeExchange          = WorkloadType("Exchange")
	WorkloadTypeSharepoint        = WorkloadType("Sharepoint")
	WorkloadTypeVMwareVM          = WorkloadType("VMwareVM")
	WorkloadTypeSystemState       = WorkloadType("SystemState")
	WorkloadTypeClient            = WorkloadType("Client")
	WorkloadTypeGenericDataSource = WorkloadType("GenericDataSource")
	WorkloadTypeSQLDataBase       = WorkloadType("SQLDataBase")
	WorkloadTypeAzureFileShare    = WorkloadType("AzureFileShare")
	WorkloadTypeSAPHanaDatabase   = WorkloadType("SAPHanaDatabase")
	WorkloadTypeSAPAseDatabase    = WorkloadType("SAPAseDatabase")
)

func init() {
}
