


package maintenance

type MaintenanceScope string

const (
	MaintenanceScopeAll        = MaintenanceScope("All")
	MaintenanceScopeHost       = MaintenanceScope("Host")
	MaintenanceScopeResource   = MaintenanceScope("Resource")
	MaintenanceScopeInResource = MaintenanceScope("InResource")
)

func init() {
}
