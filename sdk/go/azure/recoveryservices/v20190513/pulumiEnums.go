


package v20190513

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

func init() {
}
