


package devtestlab

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

type EnvironmentPermission string

const (
	EnvironmentPermissionReader      = EnvironmentPermission("Reader")
	EnvironmentPermissionContributor = EnvironmentPermission("Contributor")
)

type HostCachingOptions string

const (
	HostCachingOptionsNone      = HostCachingOptions("None")
	HostCachingOptionsReadOnly  = HostCachingOptions("ReadOnly")
	HostCachingOptionsReadWrite = HostCachingOptions("ReadWrite")
)

type LinuxOsState string

const (
	LinuxOsStateNonDeprovisioned     = LinuxOsState("NonDeprovisioned")
	LinuxOsStateDeprovisionRequested = LinuxOsState("DeprovisionRequested")
	LinuxOsStateDeprovisionApplied   = LinuxOsState("DeprovisionApplied")
)

type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                         = ManagedIdentityType("None")
	ManagedIdentityTypeSystemAssigned               = ManagedIdentityType("SystemAssigned")
	ManagedIdentityTypeUserAssigned                 = ManagedIdentityType("UserAssigned")
	ManagedIdentityType_SystemAssigned_UserAssigned = ManagedIdentityType("SystemAssigned,UserAssigned")
)

type NotificationChannelEventType string

const (
	NotificationChannelEventTypeAutoShutdown = NotificationChannelEventType("AutoShutdown")
	NotificationChannelEventTypeCost         = NotificationChannelEventType("Cost")
)

type PolicyEvaluatorType string

const (
	PolicyEvaluatorTypeAllowedValuesPolicy = PolicyEvaluatorType("AllowedValuesPolicy")
	PolicyEvaluatorTypeMaxValuePolicy      = PolicyEvaluatorType("MaxValuePolicy")
)

type PolicyFactName string

const (
	PolicyFactNameUserOwnedLabVmCount         = PolicyFactName("UserOwnedLabVmCount")
	PolicyFactNameUserOwnedLabPremiumVmCount  = PolicyFactName("UserOwnedLabPremiumVmCount")
	PolicyFactNameLabVmCount                  = PolicyFactName("LabVmCount")
	PolicyFactNameLabPremiumVmCount           = PolicyFactName("LabPremiumVmCount")
	PolicyFactNameLabVmSize                   = PolicyFactName("LabVmSize")
	PolicyFactNameGalleryImage                = PolicyFactName("GalleryImage")
	PolicyFactNameUserOwnedLabVmCountInSubnet = PolicyFactName("UserOwnedLabVmCountInSubnet")
	PolicyFactNameLabTargetCost               = PolicyFactName("LabTargetCost")
	PolicyFactNameEnvironmentTemplate         = PolicyFactName("EnvironmentTemplate")
	PolicyFactNameScheduleEditPermission      = PolicyFactName("ScheduleEditPermission")
)

type PolicyStatus string

const (
	PolicyStatusEnabled  = PolicyStatus("Enabled")
	PolicyStatusDisabled = PolicyStatus("Disabled")
)

type PremiumDataDisk string

const (
	PremiumDataDiskDisabled = PremiumDataDisk("Disabled")
	PremiumDataDiskEnabled  = PremiumDataDisk("Enabled")
)

type SourceControlType string

const (
	SourceControlTypeVsoGit         = SourceControlType("VsoGit")
	SourceControlTypeGitHub         = SourceControlType("GitHub")
	SourceControlTypeStorageAccount = SourceControlType("StorageAccount")
)

type StorageType string

const (
	StorageTypeStandard    = StorageType("Standard")
	StorageTypePremium     = StorageType("Premium")
	StorageTypeStandardSSD = StorageType("StandardSSD")
)

type TransportProtocol string

const (
	TransportProtocolTcp = TransportProtocol("Tcp")
	TransportProtocolUdp = TransportProtocol("Udp")
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
