


package v20181015

type AddRemove string

const (
	// Indicates that a user is adding a favorite lab
	AddRemoveAdd = AddRemove("Add")
	// Indicates that a user is removing a favorite lab
	AddRemoveRemove = AddRemove("Remove")
)

type ConfigurationState string

const (
	// User either hasn't started configuring their template
	// or they haven't started the configuration process.
	ConfigurationStateNotApplicable = ConfigurationState("NotApplicable")
	// User is finished modifying the template.
	ConfigurationStateCompleted = ConfigurationState("Completed")
)

type LabUserAccessMode string

const (
	// Only users registered with the lab can access VMs.
	LabUserAccessModeRestricted = LabUserAccessMode("Restricted")
	// Any user can register with the lab and access its VMs.
	LabUserAccessModeOpen = LabUserAccessMode("Open")
)

type ManagedLabVmSize string

const (
	// The base VM size
	ManagedLabVmSizeBasic = ManagedLabVmSize("Basic")
	// The standard or default VM size
	ManagedLabVmSizeStandard = ManagedLabVmSize("Standard")
	// The most performant VM size
	ManagedLabVmSizePerformance = ManagedLabVmSize("Performance")
)

func init() {
}
