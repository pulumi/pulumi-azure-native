


package v20200701preview

type MaintenanceScope string

const (
	MaintenanceScopeAll                = MaintenanceScope("All")
	MaintenanceScopeHost               = MaintenanceScope("Host")
	MaintenanceScopeResource           = MaintenanceScope("Resource")
	MaintenanceScopeInResource         = MaintenanceScope("InResource")
	MaintenanceScopeOSImage            = MaintenanceScope("OSImage")
	MaintenanceScopeExtension          = MaintenanceScope("Extension")
	MaintenanceScopeInGuestPatch       = MaintenanceScope("InGuestPatch")
	MaintenanceScopeSQLDB              = MaintenanceScope("SQLDB")
	MaintenanceScopeSQLManagedInstance = MaintenanceScope("SQLManagedInstance")
)

type Visibility string

const (
	VisibilityCustom = Visibility("Custom")
	VisibilityPublic = Visibility("Public")
)

func init() {
}
