


package v20200201preview

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
