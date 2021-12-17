


package v20210210

type AgentAutoUpdateStatus string

const (
	AgentAutoUpdateStatusDisabled = AgentAutoUpdateStatus("Disabled")
	AgentAutoUpdateStatusEnabled  = AgentAutoUpdateStatus("Enabled")
)

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

type DiskAccountType string

const (
	DiskAccountType_Standard_LRS    = DiskAccountType("Standard_LRS")
	DiskAccountType_Premium_LRS     = DiskAccountType("Premium_LRS")
	DiskAccountType_StandardSSD_LRS = DiskAccountType("StandardSSD_LRS")
)

type FailoverDeploymentModel string

const (
	FailoverDeploymentModelNotApplicable   = FailoverDeploymentModel("NotApplicable")
	FailoverDeploymentModelClassic         = FailoverDeploymentModel("Classic")
	FailoverDeploymentModelResourceManager = FailoverDeploymentModel("ResourceManager")
)

type HealthStatus string

const (
	HealthStatusPassed          = HealthStatus("Passed")
	HealthStatusActionRequired  = HealthStatus("ActionRequired")
	HealthStatusActionSuggested = HealthStatus("ActionSuggested")
	HealthStatusInvalid         = HealthStatus("Invalid")
)

type InfrastructureEncryptionState string

const (
	InfrastructureEncryptionStateEnabled  = InfrastructureEncryptionState("Enabled")
	InfrastructureEncryptionStateDisabled = InfrastructureEncryptionState("Disabled")
)

type LastBackupStatus string

const (
	LastBackupStatusInvalid   = LastBackupStatus("Invalid")
	LastBackupStatusHealthy   = LastBackupStatus("Healthy")
	LastBackupStatusUnhealthy = LastBackupStatus("Unhealthy")
	LastBackupStatusIRPending = LastBackupStatus("IRPending")
)

type LicenseType string

const (
	LicenseTypeNotSpecified  = LicenseType("NotSpecified")
	LicenseTypeNoLicenseType = LicenseType("NoLicenseType")
	LicenseTypeWindowsServer = LicenseType("WindowsServer")
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

type PossibleOperationsDirections string

const (
	PossibleOperationsDirectionsPrimaryToRecovery = PossibleOperationsDirections("PrimaryToRecovery")
	PossibleOperationsDirectionsRecoveryToPrimary = PossibleOperationsDirections("RecoveryToPrimary")
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

type RecoveryPlanGroupType string

const (
	RecoveryPlanGroupTypeShutdown = RecoveryPlanGroupType("Shutdown")
	RecoveryPlanGroupTypeBoot     = RecoveryPlanGroupType("Boot")
	RecoveryPlanGroupTypeFailover = RecoveryPlanGroupType("Failover")
)

type ReplicationProtectedItemOperation string

const (
	ReplicationProtectedItemOperationReverseReplicate    = ReplicationProtectedItemOperation("ReverseReplicate")
	ReplicationProtectedItemOperationCommit              = ReplicationProtectedItemOperation("Commit")
	ReplicationProtectedItemOperationPlannedFailover     = ReplicationProtectedItemOperation("PlannedFailover")
	ReplicationProtectedItemOperationUnplannedFailover   = ReplicationProtectedItemOperation("UnplannedFailover")
	ReplicationProtectedItemOperationDisableProtection   = ReplicationProtectedItemOperation("DisableProtection")
	ReplicationProtectedItemOperationTestFailover        = ReplicationProtectedItemOperation("TestFailover")
	ReplicationProtectedItemOperationTestFailoverCleanup = ReplicationProtectedItemOperation("TestFailoverCleanup")
	ReplicationProtectedItemOperationFailback            = ReplicationProtectedItemOperation("Failback")
	ReplicationProtectedItemOperationFinalizeFailback    = ReplicationProtectedItemOperation("FinalizeFailback")
	ReplicationProtectedItemOperationCancelFailover      = ReplicationProtectedItemOperation("CancelFailover")
	ReplicationProtectedItemOperationChangePit           = ReplicationProtectedItemOperation("ChangePit")
	ReplicationProtectedItemOperationRepairReplication   = ReplicationProtectedItemOperation("RepairReplication")
	ReplicationProtectedItemOperationSwitchProtection    = ReplicationProtectedItemOperation("SwitchProtection")
	ReplicationProtectedItemOperationCompleteMigration   = ReplicationProtectedItemOperation("CompleteMigration")
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

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
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

type SetMultiVmSyncStatus string

const (
	SetMultiVmSyncStatusEnable  = SetMultiVmSyncStatus("Enable")
	SetMultiVmSyncStatusDisable = SetMultiVmSyncStatus("Disable")
)

type SkuName string

const (
	SkuNameStandard = SkuName("Standard")
	SkuNameRS0      = SkuName("RS0")
)

type SqlServerLicenseType string

const (
	SqlServerLicenseTypeNotSpecified  = SqlServerLicenseType("NotSpecified")
	SqlServerLicenseTypeNoLicenseType = SqlServerLicenseType("NoLicenseType")
	SqlServerLicenseTypePAYG          = SqlServerLicenseType("PAYG")
	SqlServerLicenseTypeAHUB          = SqlServerLicenseType("AHUB")
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
