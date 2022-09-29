


package v20210601

type DaysOfWeek string

const (
	DaysOfWeekMonday    = DaysOfWeek("Monday")
	DaysOfWeekTuesday   = DaysOfWeek("Tuesday")
	DaysOfWeekWednesday = DaysOfWeek("Wednesday")
	DaysOfWeekThursday  = DaysOfWeek("Thursday")
	DaysOfWeekFriday    = DaysOfWeek("Friday")
	DaysOfWeekSaturday  = DaysOfWeek("Saturday")
	DaysOfWeekSunday    = DaysOfWeek("Sunday")
)

type DirectoryType string

const (
	DirectoryTypeActiveDirectory = DirectoryType("ActiveDirectory")
)

type JsonWebKeyEncryptionAlgorithm string

const (
	JsonWebKeyEncryptionAlgorithm_RSA_OAEP     = JsonWebKeyEncryptionAlgorithm("RSA-OAEP")
	JsonWebKeyEncryptionAlgorithm_RSA_OAEP_256 = JsonWebKeyEncryptionAlgorithm("RSA-OAEP-256")
	JsonWebKeyEncryptionAlgorithm_RSA1_5       = JsonWebKeyEncryptionAlgorithm("RSA1_5")
)

type OSType string

const (
	OSTypeWindows = OSType("Windows")
	OSTypeLinux   = OSType("Linux")
)

type PrivateIPAllocationMethod string

const (
	PrivateIPAllocationMethodDynamic = PrivateIPAllocationMethod("dynamic")
	PrivateIPAllocationMethodStatic  = PrivateIPAllocationMethod("static")
)

type PrivateLink string

const (
	PrivateLinkDisabled = PrivateLink("Disabled")
	PrivateLinkEnabled  = PrivateLink("Enabled")
)

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved = PrivateLinkServiceConnectionStatus("Approved")
	PrivateLinkServiceConnectionStatusRejected = PrivateLinkServiceConnectionStatus("Rejected")
	PrivateLinkServiceConnectionStatusPending  = PrivateLinkServiceConnectionStatus("Pending")
	PrivateLinkServiceConnectionStatusRemoved  = PrivateLinkServiceConnectionStatus("Removed")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

type ResourceProviderConnection string

const (
	ResourceProviderConnectionInbound  = ResourceProviderConnection("Inbound")
	ResourceProviderConnectionOutbound = ResourceProviderConnection("Outbound")
)

type Tier string

const (
	TierStandard = Tier("Standard")
	TierPremium  = Tier("Premium")
)

func init() {
}
