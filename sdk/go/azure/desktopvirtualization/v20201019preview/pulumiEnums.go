


package v20201019preview

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
	HostPoolTypePooled   = HostPoolType("Pooled")
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

type RegistrationTokenOperation string

const (
	RegistrationTokenOperationDelete = RegistrationTokenOperation("Delete")
	RegistrationTokenOperationNone   = RegistrationTokenOperation("None")
	RegistrationTokenOperationUpdate = RegistrationTokenOperation("Update")
)

type RemoteApplicationType string

const (
	RemoteApplicationTypeInBuilt         = RemoteApplicationType("InBuilt")
	RemoteApplicationTypeMsixApplication = RemoteApplicationType("MsixApplication")
)

type SSOSecretType string

const (
	SSOSecretTypeSharedKey             = SSOSecretType("SharedKey")
	SSOSecretTypeCertificate           = SSOSecretType("Certificate")
	SSOSecretTypeSharedKeyInKeyVault   = SSOSecretType("SharedKeyInKeyVault")
	SSOSecretTypeCertificateInKeyVault = SSOSecretType("CertificateInKeyVault")
)

func init() {
}
