


package v20190123preview

type ApplicationGroupType string

const (
	ApplicationGroupTypeRemoteApp = ApplicationGroupType("RemoteApp")
	ApplicationGroupTypeDesktop   = ApplicationGroupType("Desktop")
)

type CommandLineSetting string

const (
	CommandLineSettingDoNotAllow = CommandLineSetting("DoNotAllow")
	CommandLineSettingAllow      = CommandLineSetting("Allow")
	CommandLineSettingRequire    = CommandLineSetting("Require")
)

type HostPoolType string

const (
	HostPoolTypePersonal = HostPoolType("Personal")
	HostPoolTypeShared   = HostPoolType("Shared")
)

type LoadBalancerType string

const (
	LoadBalancerTypeBreadthFirst = LoadBalancerType("BreadthFirst")
	LoadBalancerTypeDepthFirst   = LoadBalancerType("DepthFirst")
	LoadBalancerTypePersistent   = LoadBalancerType("Persistent")
)

type PersonalDesktopAssignmentType string

const (
	PersonalDesktopAssignmentTypeAutomatic = PersonalDesktopAssignmentType("Automatic")
	PersonalDesktopAssignmentTypeDirect    = PersonalDesktopAssignmentType("Direct")
)

type PreferredAppGroupType string

const (
	PreferredAppGroupTypeNone             = PreferredAppGroupType("None")
	PreferredAppGroupTypeDesktop          = PreferredAppGroupType("Desktop")
	PreferredAppGroupTypeRailApplications = PreferredAppGroupType("RailApplications")
)

func init() {
}
