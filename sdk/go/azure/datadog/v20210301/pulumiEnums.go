


package v20210301

type ManagedIdentityTypes string

const (
	ManagedIdentityTypesSystemAssigned = ManagedIdentityTypes("SystemAssigned")
	ManagedIdentityTypesUserAssigned   = ManagedIdentityTypes("UserAssigned")
)

type MonitoringStatus string

const (
	MonitoringStatusEnabled  = MonitoringStatus("Enabled")
	MonitoringStatusDisabled = MonitoringStatus("Disabled")
)

func init() {
}
