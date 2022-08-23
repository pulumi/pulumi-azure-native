


package v20180201preview

type BuildTaskStatus string

const (
	BuildTaskStatusDisabled = BuildTaskStatus("Disabled")
	BuildTaskStatusEnabled  = BuildTaskStatus("Enabled")
)

type OsType string

const (
	OsTypeWindows = OsType("Windows")
	OsTypeLinux   = OsType("Linux")
)

type SourceControlType string

const (
	SourceControlTypeGithub                  = SourceControlType("Github")
	SourceControlTypeVisualStudioTeamService = SourceControlType("VisualStudioTeamService")
)

type TokenType string

const (
	TokenTypePAT   = TokenType("PAT")
	TokenTypeOAuth = TokenType("OAuth")
)

func init() {
}
