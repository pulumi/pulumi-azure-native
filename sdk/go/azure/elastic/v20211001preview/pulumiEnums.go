


package v20211001preview

type ManagedIdentityTypes string

const (
	ManagedIdentityTypesSystemAssigned = ManagedIdentityTypes("SystemAssigned")
)

type MonitoringStatus string

const (
	MonitoringStatusEnabled  = MonitoringStatus("Enabled")
	MonitoringStatusDisabled = MonitoringStatus("Disabled")
)

type ProvisioningState string

const (
	ProvisioningStateAccepted     = ProvisioningState("Accepted")
	ProvisioningStateCreating     = ProvisioningState("Creating")
	ProvisioningStateUpdating     = ProvisioningState("Updating")
	ProvisioningStateDeleting     = ProvisioningState("Deleting")
	ProvisioningStateSucceeded    = ProvisioningState("Succeeded")
	ProvisioningStateFailed       = ProvisioningState("Failed")
	ProvisioningStateCanceled     = ProvisioningState("Canceled")
	ProvisioningStateDeleted      = ProvisioningState("Deleted")
	ProvisioningStateNotSpecified = ProvisioningState("NotSpecified")
)

type TagAction string

const (
	TagActionInclude = TagAction("Include")
	TagActionExclude = TagAction("Exclude")
)

func init() {
}
