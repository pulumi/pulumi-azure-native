


package v20200401preview

type ChangeDetailsMode string

const (
	ChangeDetailsModeNone    = ChangeDetailsMode("None")
	ChangeDetailsModeInclude = ChangeDetailsMode("Include")
	ChangeDetailsModeExclude = ChangeDetailsMode("Exclude")
)

type ManagedIdentityTypes string

const (
	ManagedIdentityTypesNone           = ManagedIdentityTypes("None")
	ManagedIdentityTypesSystemAssigned = ManagedIdentityTypes("SystemAssigned")
)

type NotificationsState string

const (
	NotificationsStateNone     = NotificationsState("None")
	NotificationsStateEnabled  = NotificationsState("Enabled")
	NotificationsStateDisabled = NotificationsState("Disabled")
)

func init() {
}
