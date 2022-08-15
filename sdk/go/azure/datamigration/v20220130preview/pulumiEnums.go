


package v20220130preview

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

type MySqlTargetPlatformType string

const (
	MySqlTargetPlatformTypeSqlServer       = MySqlTargetPlatformType("SqlServer")
	MySqlTargetPlatformTypeAzureDbForMySQL = MySqlTargetPlatformType("AzureDbForMySQL")
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
	ServerLevelPermissionsGroupMigrationFromSqlServerToAzureVM     = ServerLevelPermissionsGroup("MigrationFromSqlServerToAzureVM")
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

type TaskType string

const (
	TaskType_Connect_MongoDb                                        = TaskType("Connect.MongoDb")
	TaskType_ConnectToSource_SqlServer                              = TaskType("ConnectToSource.SqlServer")
	TaskType_ConnectToSource_SqlServer_Sync                         = TaskType("ConnectToSource.SqlServer.Sync")
	TaskType_ConnectToSource_PostgreSql_Sync                        = TaskType("ConnectToSource.PostgreSql.Sync")
	TaskType_ConnectToSource_MySql                                  = TaskType("ConnectToSource.MySql")
	TaskType_ConnectToSource_Oracle_Sync                            = TaskType("ConnectToSource.Oracle.Sync")
	TaskType_ConnectToTarget_SqlDb                                  = TaskType("ConnectToTarget.SqlDb")
	TaskType_ConnectToTarget_SqlDb_Sync                             = TaskType("ConnectToTarget.SqlDb.Sync")
	TaskType_ConnectToTarget_AzureDbForPostgreSql_Sync              = TaskType("ConnectToTarget.AzureDbForPostgreSql.Sync")
	TaskType_ConnectToTarget_Oracle_AzureDbForPostgreSql_Sync       = TaskType("ConnectToTarget.Oracle.AzureDbForPostgreSql.Sync")
	TaskType_ConnectToTarget_AzureSqlDbMI                           = TaskType("ConnectToTarget.AzureSqlDbMI")
	TaskType_ConnectToTarget_AzureSqlDbMI_Sync_LRS                  = TaskType("ConnectToTarget.AzureSqlDbMI.Sync.LRS")
	TaskType_ConnectToTarget_AzureDbForMySql                        = TaskType("ConnectToTarget.AzureDbForMySql")
	TaskType_GetUserTables_Sql                                      = TaskType("GetUserTables.Sql")
	TaskType_GetUserTables_AzureSqlDb_Sync                          = TaskType("GetUserTables.AzureSqlDb.Sync")
	TaskTypeGetUserTablesOracle                                     = TaskType("GetUserTablesOracle")
	TaskTypeGetUserTablesPostgreSql                                 = TaskType("GetUserTablesPostgreSql")
	TaskTypeGetUserTablesMySql                                      = TaskType("GetUserTablesMySql")
	TaskType_Migrate_MongoDb                                        = TaskType("Migrate.MongoDb")
	TaskType_Migrate_SqlServer_AzureSqlDbMI                         = TaskType("Migrate.SqlServer.AzureSqlDbMI")
	TaskType_Migrate_SqlServer_AzureSqlDbMI_Sync_LRS                = TaskType("Migrate.SqlServer.AzureSqlDbMI.Sync.LRS")
	TaskType_Migrate_SqlServer_SqlDb                                = TaskType("Migrate.SqlServer.SqlDb")
	TaskType_Migrate_SqlServer_AzureSqlDb_Sync                      = TaskType("Migrate.SqlServer.AzureSqlDb.Sync")
	TaskType_Migrate_MySql_AzureDbForMySql_Sync                     = TaskType("Migrate.MySql.AzureDbForMySql.Sync")
	TaskType_Migrate_MySql_AzureDbForMySql                          = TaskType("Migrate.MySql.AzureDbForMySql")
	TaskType_Migrate_PostgreSql_AzureDbForPostgreSql_SyncV2         = TaskType("Migrate.PostgreSql.AzureDbForPostgreSql.SyncV2")
	TaskType_Migrate_Oracle_AzureDbForPostgreSql_Sync               = TaskType("Migrate.Oracle.AzureDbForPostgreSql.Sync")
	TaskType_ValidateMigrationInput_SqlServer_SqlDb_Sync            = TaskType("ValidateMigrationInput.SqlServer.SqlDb.Sync")
	TaskType_ValidateMigrationInput_SqlServer_AzureSqlDbMI          = TaskType("ValidateMigrationInput.SqlServer.AzureSqlDbMI")
	TaskType_ValidateMigrationInput_SqlServer_AzureSqlDbMI_Sync_LRS = TaskType("ValidateMigrationInput.SqlServer.AzureSqlDbMI.Sync.LRS")
	TaskType_Validate_MongoDb                                       = TaskType("Validate.MongoDb")
	TaskType_Validate_Oracle_AzureDbPostgreSql_Sync                 = TaskType("Validate.Oracle.AzureDbPostgreSql.Sync")
	TaskType_GetTDECertificates_Sql                                 = TaskType("GetTDECertificates.Sql")
	TaskType_Migrate_Ssis                                           = TaskType("Migrate.Ssis")
	TaskType_Service_Check_OCI                                      = TaskType("Service.Check.OCI")
	TaskType_Service_Upload_OCI                                     = TaskType("Service.Upload.OCI")
	TaskType_Service_Install_OCI                                    = TaskType("Service.Install.OCI")
	TaskTypeMigrateSchemaSqlServerSqlDb                             = TaskType("MigrateSchemaSqlServerSqlDb")
)

func init() {
}
