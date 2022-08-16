


package v20200601

type PersistedConnectionStatus string

const (
	PersistedConnectionStatusPending      = PersistedConnectionStatus("Pending")
	PersistedConnectionStatusApproved     = PersistedConnectionStatus("Approved")
	PersistedConnectionStatusRejected     = PersistedConnectionStatus("Rejected")
	PersistedConnectionStatusDisconnected = PersistedConnectionStatus("Disconnected")
)

type ResourceProvisioningState string

const (
	ResourceProvisioningStateCreating  = ResourceProvisioningState("Creating")
	ResourceProvisioningStateUpdating  = ResourceProvisioningState("Updating")
	ResourceProvisioningStateDeleting  = ResourceProvisioningState("Deleting")
	ResourceProvisioningStateSucceeded = ResourceProvisioningState("Succeeded")
	ResourceProvisioningStateCanceled  = ResourceProvisioningState("Canceled")
	ResourceProvisioningStateFailed    = ResourceProvisioningState("Failed")
)

func init() {
}
