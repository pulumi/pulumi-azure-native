


package v20210501

type MaintenanceScope string

const (
	// This maintenance scope controls installation of azure platform updates i.e. services on physical nodes hosting customer VMs.
	MaintenanceScopeHost = MaintenanceScope("Host")
	// This maintenance scope controls os image installation on VM/VMSS
	MaintenanceScopeOSImage = MaintenanceScope("OSImage")
	// This maintenance scope controls extension installation on VM/VMSS
	MaintenanceScopeExtension = MaintenanceScope("Extension")
	// This maintenance scope controls installation of windows and linux packages on VM/VMSS
	MaintenanceScopeInGuestPatch = MaintenanceScope("InGuestPatch")
	// This maintenance scope controls installation of SQL server platform updates.
	MaintenanceScopeSQLDB = MaintenanceScope("SQLDB")
	// This maintenance scope controls installation of SQL managed instance platform update.
	MaintenanceScopeSQLManagedInstance = MaintenanceScope("SQLManagedInstance")
)

type Visibility string

const (
	// Only visible to users with permissions.
	VisibilityCustom = Visibility("Custom")
	// Visible to all users.
	VisibilityPublic = Visibility("Public")
)

func init() {
}
