


package v20180901preview

type ProvisioningState string

const (
	ProvisioningStateAccepted  = ProvisioningState("Accepted")
	ProvisioningStateCreating  = ProvisioningState("Creating")
	ProvisioningStateDeleting  = ProvisioningState("Deleting")
	ProvisioningStateFailed    = ProvisioningState("Failed")
	ProvisioningStateMoving    = ProvisioningState("Moving")
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
)

func init() {
}
