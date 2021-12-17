


package v20180901

type Architecture string

const (
	ArchitectureAmd64 = Architecture("amd64")
	ArchitectureX86   = Architecture("x86")
	ArchitectureArm   = Architecture("arm")
)

type BaseImageTriggerType string

const (
	BaseImageTriggerTypeAll     = BaseImageTriggerType("All")
	BaseImageTriggerTypeRuntime = BaseImageTriggerType("Runtime")
)

type OS string

const (
	OSWindows = OS("Windows")
	OSLinux   = OS("Linux")
)

type SecretObjectType string

const (
	SecretObjectTypeOpaque = SecretObjectType("Opaque")
)

type SourceControlType string

const (
	SourceControlTypeGithub                  = SourceControlType("Github")
	SourceControlTypeVisualStudioTeamService = SourceControlType("VisualStudioTeamService")
)

type SourceRegistryLoginMode string

const (
	SourceRegistryLoginModeNone    = SourceRegistryLoginMode("None")
	SourceRegistryLoginModeDefault = SourceRegistryLoginMode("Default")
)

type SourceTriggerEvent string

const (
	SourceTriggerEventCommit      = SourceTriggerEvent("commit")
	SourceTriggerEventPullrequest = SourceTriggerEvent("pullrequest")
)

type TaskStatus string

const (
	TaskStatusDisabled = TaskStatus("Disabled")
	TaskStatusEnabled  = TaskStatus("Enabled")
)

type TokenType string

const (
	TokenTypePAT   = TokenType("PAT")
	TokenTypeOAuth = TokenType("OAuth")
)

type TriggerStatus string

const (
	TriggerStatusDisabled = TriggerStatus("Disabled")
	TriggerStatusEnabled  = TriggerStatus("Enabled")
)

type Variant string

const (
	VariantV6 = Variant("v6")
	VariantV7 = Variant("v7")
	VariantV8 = Variant("v8")
)

func init() {
}
