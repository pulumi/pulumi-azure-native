


package fluidrelay

type ProvisioningState string

const (
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
	ProvisioningStateFailed    = ProvisioningState("Failed")
	ProvisioningStateCanceled  = ProvisioningState("Canceled")
)

func init() {
}
