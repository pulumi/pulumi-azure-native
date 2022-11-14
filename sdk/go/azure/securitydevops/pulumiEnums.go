


package securitydevops

type AutoDiscovery string

const (
	AutoDiscoveryDisabled = AutoDiscovery("Disabled")
	AutoDiscoveryEnabled  = AutoDiscovery("Enabled")
)

type ProvisioningState string

const (
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
	ProvisioningStateFailed    = ProvisioningState("Failed")
	ProvisioningStateCanceled  = ProvisioningState("Canceled")
)

func init() {
}
