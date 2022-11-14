


package v20180715preview

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

type MongoDbReplication string

const (
	MongoDbReplicationDisabled   = MongoDbReplication("Disabled")
	MongoDbReplicationOneTime    = MongoDbReplication("OneTime")
	MongoDbReplicationContinuous = MongoDbReplication("Continuous")
)

type MongoDbShardKeyOrder string

const (
	MongoDbShardKeyOrderForward = MongoDbShardKeyOrder("Forward")
	MongoDbShardKeyOrderReverse = MongoDbShardKeyOrder("Reverse")
	MongoDbShardKeyOrderHashed  = MongoDbShardKeyOrder("Hashed")
)

type ProjectSourcePlatform string

const (
	ProjectSourcePlatformSQL        = ProjectSourcePlatform("SQL")
	ProjectSourcePlatformMySQL      = ProjectSourcePlatform("MySQL")
	ProjectSourcePlatformPostgreSql = ProjectSourcePlatform("PostgreSql")
	ProjectSourcePlatformMongoDb    = ProjectSourcePlatform("MongoDb")
	ProjectSourcePlatformUnknown    = ProjectSourcePlatform("Unknown")
)

type ProjectTargetPlatform string

const (
	ProjectTargetPlatformSQLDB                = ProjectTargetPlatform("SQLDB")
	ProjectTargetPlatformSQLMI                = ProjectTargetPlatform("SQLMI")
	ProjectTargetPlatformAzureDbForMySql      = ProjectTargetPlatform("AzureDbForMySql")
	ProjectTargetPlatformAzureDbForPostgreSql = ProjectTargetPlatform("AzureDbForPostgreSql")
	ProjectTargetPlatformMongoDb              = ProjectTargetPlatform("MongoDb")
	ProjectTargetPlatformUnknown              = ProjectTargetPlatform("Unknown")
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

type SsisMigrationOverwriteOption string

const (
	SsisMigrationOverwriteOptionIgnore    = SsisMigrationOverwriteOption("Ignore")
	SsisMigrationOverwriteOptionOverwrite = SsisMigrationOverwriteOption("Overwrite")
)

type SsisStoreType string

const (
	SsisStoreTypeSsisCatalog = SsisStoreType("SsisCatalog")
)

func init() {
}
