


package v20210312preview

type ProvisioningState string

const (
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
	ProvisioningStateFailed    = ProvisioningState("Failed")
	ProvisioningStateCanceled  = ProvisioningState("Canceled")
)

func init() {
}
