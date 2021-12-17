


package v20190701preview

type AuthorizationType string

const (
	AuthorizationTypePersonalAccessToken = AuthorizationType("personalAccessToken")
)

type CodeRepositoryType string

const (
	CodeRepositoryTypeGitHub  = CodeRepositoryType("gitHub")
	CodeRepositoryTypeVstsGit = CodeRepositoryType("vstsGit")
)

func init() {
}
