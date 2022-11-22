


package v20220930preview

type AlertsState string

const (
	AlertsStateEnabled  = AlertsState("Enabled")
	AlertsStateDisabled = AlertsState("Disabled")
)

type ImmutabilityState string

const (
	ImmutabilityStateDisabled = ImmutabilityState("Disabled")
	ImmutabilityStateUnlocked = ImmutabilityState("Unlocked")
	ImmutabilityStateLocked   = ImmutabilityState("Locked")
)

type InfrastructureEncryptionState string

const (
	InfrastructureEncryptionStateEnabled  = InfrastructureEncryptionState("Enabled")
	InfrastructureEncryptionStateDisabled = InfrastructureEncryptionState("Disabled")
)

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
)

type SkuName string

const (
	SkuNameStandard = SkuName("Standard")
	SkuNameRS0      = SkuName("RS0")
)

func init() {
}
