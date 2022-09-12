


package v20150521preview

type CustomImageOsType string

const (
	CustomImageOsTypeWindows = CustomImageOsType("Windows")
	CustomImageOsTypeLinux   = CustomImageOsType("Linux")
	CustomImageOsTypeNone    = CustomImageOsType("None")
)

type EnableStatus string

const (
	EnableStatusEnabled  = EnableStatus("Enabled")
	EnableStatusDisabled = EnableStatus("Disabled")
)

type LabStorageType string

const (
	LabStorageTypeStandard = LabStorageType("Standard")
	LabStorageTypePremium  = LabStorageType("Premium")
)

type LinuxOsState string

const (
	LinuxOsStateNonDeprovisioned     = LinuxOsState("NonDeprovisioned")
	LinuxOsStateDeprovisionRequested = LinuxOsState("DeprovisionRequested")
	LinuxOsStateDeprovisionApplied   = LinuxOsState("DeprovisionApplied")
)

type PolicyEvaluatorType string

const (
	PolicyEvaluatorTypeAllowedValuesPolicy = PolicyEvaluatorType("AllowedValuesPolicy")
	PolicyEvaluatorTypeMaxValuePolicy      = PolicyEvaluatorType("MaxValuePolicy")
)

type PolicyFactName string

const (
	PolicyFactNameUserOwnedLabVmCount         = PolicyFactName("UserOwnedLabVmCount")
	PolicyFactNameLabVmCount                  = PolicyFactName("LabVmCount")
	PolicyFactNameLabVmSize                   = PolicyFactName("LabVmSize")
	PolicyFactNameGalleryImage                = PolicyFactName("GalleryImage")
	PolicyFactNameUserOwnedLabVmCountInSubnet = PolicyFactName("UserOwnedLabVmCountInSubnet")
)

type PolicyStatus string

const (
	PolicyStatusEnabled  = PolicyStatus("Enabled")
	PolicyStatusDisabled = PolicyStatus("Disabled")
)

type SourceControlType string

const (
	SourceControlTypeVsoGit = SourceControlType("VsoGit")
	SourceControlTypeGitHub = SourceControlType("GitHub")
)

type TaskType string

const (
	TaskTypeLabVmsShutdownTask = TaskType("LabVmsShutdownTask")
	TaskTypeLabVmsStartupTask  = TaskType("LabVmsStartupTask")
	TaskTypeLabBillingTask     = TaskType("LabBillingTask")
)

type UsagePermissionType string

const (
	UsagePermissionTypeDefault = UsagePermissionType("Default")
	UsagePermissionTypeDeny    = UsagePermissionType("Deny")
	UsagePermissionTypeAllow   = UsagePermissionType("Allow")
)

type WindowsOsState string

const (
	WindowsOsStateNonSysprepped    = WindowsOsState("NonSysprepped")
	WindowsOsStateSysprepRequested = WindowsOsState("SysprepRequested")
	WindowsOsStateSysprepApplied   = WindowsOsState("SysprepApplied")
)

func init() {
}
