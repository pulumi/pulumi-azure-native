


package v20211001

type ResourceDiscoveryMode string

const (
	// Remediate resources that are already known to be non-compliant.
	ResourceDiscoveryModeExistingNonCompliant = ResourceDiscoveryMode("ExistingNonCompliant")
	// Re-evaluate the compliance state of resources and then remediate the resources found to be non-compliant.
	ResourceDiscoveryModeReEvaluateCompliance = ResourceDiscoveryMode("ReEvaluateCompliance")
)

func init() {
}
