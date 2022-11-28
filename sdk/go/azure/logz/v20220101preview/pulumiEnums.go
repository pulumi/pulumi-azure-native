


package v20220101preview

type ManagedIdentityTypes string

const (
	ManagedIdentityTypesSystemAssigned = ManagedIdentityTypes("SystemAssigned")
	ManagedIdentityTypesUserAssigned   = ManagedIdentityTypes("UserAssigned")
)

type MarketplaceSubscriptionStatus string

const (
	MarketplaceSubscriptionStatusActive    = MarketplaceSubscriptionStatus("Active")
	MarketplaceSubscriptionStatusSuspended = MarketplaceSubscriptionStatus("Suspended")
)

type MonitoringStatus string

const (
	MonitoringStatusEnabled  = MonitoringStatus("Enabled")
	MonitoringStatusDisabled = MonitoringStatus("Disabled")
)

type TagAction string

const (
	TagActionInclude = TagAction("Include")
	TagActionExclude = TagAction("Exclude")
)

func init() {
}
