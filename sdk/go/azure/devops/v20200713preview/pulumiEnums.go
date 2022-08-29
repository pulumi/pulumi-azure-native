


package v20200713preview

type AuthorizationType string

const (
	AuthorizationTypePersonalAccessToken = AuthorizationType("personalAccessToken")
)

type CodeRepositoryType string

const (
	CodeRepositoryTypeGitHub  = CodeRepositoryType("gitHub")
	CodeRepositoryTypeVstsGit = CodeRepositoryType("vstsGit")
)

type PipelineTypeEnum string

const (
	PipelineTypeEnumGithubWorkflow = PipelineTypeEnum("githubWorkflow")
	PipelineTypeEnumAzurePipeline  = PipelineTypeEnum("azurePipeline")
)

func init() {
}
