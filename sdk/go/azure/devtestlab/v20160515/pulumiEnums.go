


package v20160515

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

type LinuxOsState string

const (
	LinuxOsStateNonDeprovisioned     = LinuxOsState("NonDeprovisioned")
	LinuxOsStateDeprovisionRequested = LinuxOsState("DeprovisionRequested")
	LinuxOsStateDeprovisionApplied   = LinuxOsState("DeprovisionApplied")
)

type NotificationChannelEventType string

const (
	NotificationChannelEventTypeAutoShutdown = NotificationChannelEventType("AutoShutdown")
	NotificationChannelEventTypeCost         = NotificationChannelEventType("Cost")
)

type NotificationStatus string

const (
	NotificationStatusDisabled = NotificationStatus("Disabled")
	NotificationStatusEnabled  = NotificationStatus("Enabled")
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
	SourceControlTypeVsoGit = SourceControlType("VsoGit")
	SourceControlTypeGitHub = SourceControlType("GitHub")
)

type StorageType string

const (
	StorageTypeStandard = StorageType("Standard")
	StorageTypePremium  = StorageType("Premium")
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

type VirtualMachineCreationSource string

const (
	VirtualMachineCreationSourceFromCustomImage  = VirtualMachineCreationSource("FromCustomImage")
	VirtualMachineCreationSourceFromGalleryImage = VirtualMachineCreationSource("FromGalleryImage")
)

type WindowsOsState string

const (
	WindowsOsStateNonSysprepped    = WindowsOsState("NonSysprepped")
	WindowsOsStateSysprepRequested = WindowsOsState("SysprepRequested")
	WindowsOsStateSysprepApplied   = WindowsOsState("SysprepApplied")
)

func init() {
}
