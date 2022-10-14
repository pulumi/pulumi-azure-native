


package v20171001preview

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

type ElasticPoolLicenseType string

const (
	ElasticPoolLicenseTypeLicenseIncluded = ElasticPoolLicenseType("LicenseIncluded")
	ElasticPoolLicenseTypeBasePrice       = ElasticPoolLicenseType("BasePrice")
)

type ReadOnlyEndpointFailoverPolicy string

const (
	ReadOnlyEndpointFailoverPolicyDisabled = ReadOnlyEndpointFailoverPolicy("Disabled")
	ReadOnlyEndpointFailoverPolicyEnabled  = ReadOnlyEndpointFailoverPolicy("Enabled")
)

type ReadWriteEndpointFailoverPolicy string

const (
	ReadWriteEndpointFailoverPolicyManual    = ReadWriteEndpointFailoverPolicy("Manual")
	ReadWriteEndpointFailoverPolicyAutomatic = ReadWriteEndpointFailoverPolicy("Automatic")
)

type SampleName string

const (
	SampleNameAdventureWorksLT       = SampleName("AdventureWorksLT")
	SampleNameWideWorldImportersStd  = SampleName("WideWorldImportersStd")
	SampleNameWideWorldImportersFull = SampleName("WideWorldImportersFull")
)

type ServerKeyType string

const (
	ServerKeyTypeServiceManaged = ServerKeyType("ServiceManaged")
	ServerKeyTypeAzureKeyVault  = ServerKeyType("AzureKeyVault")
)

func init() {
}
