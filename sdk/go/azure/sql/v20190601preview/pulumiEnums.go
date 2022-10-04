


package v20190601preview

type AdministratorType string

const (
	AdministratorTypeActiveDirectory = AdministratorType("ActiveDirectory")
)

type CatalogCollationType string

const (
	CatalogCollationType_DATABASE_DEFAULT             = CatalogCollationType("DATABASE_DEFAULT")
	CatalogCollationType_SQL_Latin1_General_CP1_CI_AS = CatalogCollationType("SQL_Latin1_General_CP1_CI_AS")
)

type CreateMode string

const (
	CreateModeDefault                        = CreateMode("Default")
	CreateModeCopy                           = CreateMode("Copy")
	CreateModeSecondary                      = CreateMode("Secondary")
	CreateModePointInTimeRestore             = CreateMode("PointInTimeRestore")
	CreateModeRestore                        = CreateMode("Restore")
	CreateModeRecovery                       = CreateMode("Recovery")
	CreateModeRestoreExternalBackup          = CreateMode("RestoreExternalBackup")
	CreateModeRestoreExternalBackupSecondary = CreateMode("RestoreExternalBackupSecondary")
	CreateModeRestoreLongTermRetentionBackup = CreateMode("RestoreLongTermRetentionBackup")
	CreateModeOnlineSecondary                = CreateMode("OnlineSecondary")
)

type DatabaseLicenseType string

const (
	DatabaseLicenseTypeLicenseIncluded = DatabaseLicenseType("LicenseIncluded")
	DatabaseLicenseTypeBasePrice       = DatabaseLicenseType("BasePrice")
)

type DatabaseReadScale string

const (
	DatabaseReadScaleEnabled  = DatabaseReadScale("Enabled")
	DatabaseReadScaleDisabled = DatabaseReadScale("Disabled")
)

type IdentityType string

const (
	IdentityTypeNone                         = IdentityType("None")
	IdentityTypeSystemAssigned               = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned                 = IdentityType("UserAssigned")
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned,UserAssigned")
)

type ManagedDatabaseCreateMode string

const (
	ManagedDatabaseCreateModeDefault                        = ManagedDatabaseCreateMode("Default")
	ManagedDatabaseCreateModeRestoreExternalBackup          = ManagedDatabaseCreateMode("RestoreExternalBackup")
	ManagedDatabaseCreateModePointInTimeRestore             = ManagedDatabaseCreateMode("PointInTimeRestore")
	ManagedDatabaseCreateModeRecovery                       = ManagedDatabaseCreateMode("Recovery")
	ManagedDatabaseCreateModeRestoreLongTermRetentionBackup = ManagedDatabaseCreateMode("RestoreLongTermRetentionBackup")
)

type SampleName string

const (
	SampleNameAdventureWorksLT       = SampleName("AdventureWorksLT")
	SampleNameWideWorldImportersStd  = SampleName("WideWorldImportersStd")
	SampleNameWideWorldImportersFull = SampleName("WideWorldImportersFull")
)

type ServerPublicNetworkAccess string

const (
	ServerPublicNetworkAccessEnabled  = ServerPublicNetworkAccess("Enabled")
	ServerPublicNetworkAccessDisabled = ServerPublicNetworkAccess("Disabled")
)

type StorageAccountType string

const (
	StorageAccountTypeGRS = StorageAccountType("GRS")
	StorageAccountTypeLRS = StorageAccountType("LRS")
	StorageAccountTypeZRS = StorageAccountType("ZRS")
)

type SyncConflictResolutionPolicy string

const (
	SyncConflictResolutionPolicyHubWin    = SyncConflictResolutionPolicy("HubWin")
	SyncConflictResolutionPolicyMemberWin = SyncConflictResolutionPolicy("MemberWin")
)

type SyncDirection string

const (
	SyncDirectionBidirectional     = SyncDirection("Bidirectional")
	SyncDirectionOneWayMemberToHub = SyncDirection("OneWayMemberToHub")
	SyncDirectionOneWayHubToMember = SyncDirection("OneWayHubToMember")
)

type SyncMemberDbType string

const (
	SyncMemberDbTypeAzureSqlDatabase  = SyncMemberDbType("AzureSqlDatabase")
	SyncMemberDbTypeSqlServerDatabase = SyncMemberDbType("SqlServerDatabase")
)

func init() {
}
