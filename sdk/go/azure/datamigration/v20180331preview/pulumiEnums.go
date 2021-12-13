


package v20180331preview

type AuthenticationType string

const (
	AuthenticationTypeNone                      = AuthenticationType("None")
	AuthenticationTypeWindowsAuthentication     = AuthenticationType("WindowsAuthentication")
	AuthenticationTypeSqlAuthentication         = AuthenticationType("SqlAuthentication")
	AuthenticationTypeActiveDirectoryIntegrated = AuthenticationType("ActiveDirectoryIntegrated")
	AuthenticationTypeActiveDirectoryPassword   = AuthenticationType("ActiveDirectoryPassword")
)

type ProjectSourcePlatform string

const (
	ProjectSourcePlatformSQL     = ProjectSourcePlatform("SQL")
	ProjectSourcePlatformUnknown = ProjectSourcePlatform("Unknown")
)

type ProjectTargetPlatform string

const (
	ProjectTargetPlatformSQLDB   = ProjectTargetPlatform("SQLDB")
	ProjectTargetPlatformSQLMI   = ProjectTargetPlatform("SQLMI")
	ProjectTargetPlatformUnknown = ProjectTargetPlatform("Unknown")
)

type ServerLevelPermissionsGroup string

const (
	ServerLevelPermissionsGroupDefault                         = ServerLevelPermissionsGroup("Default")
	ServerLevelPermissionsGroupMigrationFromSqlServerToAzureDB = ServerLevelPermissionsGroup("MigrationFromSqlServerToAzureDB")
)

func init() {
}
