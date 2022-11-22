


package datamigration

type AuthenticationType string

const (
	AuthenticationTypeNone                      = AuthenticationType("None")
	AuthenticationTypeWindowsAuthentication     = AuthenticationType("WindowsAuthentication")
	AuthenticationTypeSqlAuthentication         = AuthenticationType("SqlAuthentication")
	AuthenticationTypeActiveDirectoryIntegrated = AuthenticationType("ActiveDirectoryIntegrated")
	AuthenticationTypeActiveDirectoryPassword   = AuthenticationType("ActiveDirectoryPassword")
)

type BackupMode string

const (
	BackupModeCreateBackup   = BackupMode("CreateBackup")
	BackupModeExistingBackup = BackupMode("ExistingBackup")
)

type ProjectSourcePlatform string

const (
	ProjectSourcePlatformSQL     = ProjectSourcePlatform("SQL")
	ProjectSourcePlatformUnknown = ProjectSourcePlatform("Unknown")
)

type ProjectTargetPlatform string

const (
	ProjectTargetPlatformSQLDB   = ProjectTargetPlatform("SQLDB")
	ProjectTargetPlatformUnknown = ProjectTargetPlatform("Unknown")
)

type ResourceType string

const (
	ResourceTypeSqlMi = ResourceType("SqlMi")
	ResourceTypeSqlVm = ResourceType("SqlVm")
	ResourceTypeSqlDb = ResourceType("SqlDb")
)

type ServerLevelPermissionsGroup string

const (
	ServerLevelPermissionsGroupDefault                             = ServerLevelPermissionsGroup("Default")
	ServerLevelPermissionsGroupMigrationFromSqlServerToAzureDB     = ServerLevelPermissionsGroup("MigrationFromSqlServerToAzureDB")
	ServerLevelPermissionsGroupMigrationFromSqlServerToAzureMI     = ServerLevelPermissionsGroup("MigrationFromSqlServerToAzureMI")
	ServerLevelPermissionsGroupMigrationFromMySQLToAzureDBForMySQL = ServerLevelPermissionsGroup("MigrationFromMySQLToAzureDBForMySQL")
)

type SqlSourcePlatform string

const (
	SqlSourcePlatformSqlOnPrem = SqlSourcePlatform("SqlOnPrem")
)

func init() {
}
