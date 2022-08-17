


package policyinsights

type ComplianceState string

const (
	// The resource is in compliance with the policy.
	ComplianceStateCompliant = ComplianceState("Compliant")
	// The resource is not in compliance with the policy.
	ComplianceStateNonCompliant = ComplianceState("NonCompliant")
	// The compliance state of the resource is not known.
	ComplianceStateUnknown = ComplianceState("Unknown")
)

type ResourceDiscoveryMode string

const (
	// Remediate resources that are already known to be non-compliant.
	ResourceDiscoveryModeExistingNonCompliant = ResourceDiscoveryMode("ExistingNonCompliant")
	// Re-evaluate the compliance state of resources and then remediate the resources found to be non-compliant.
	ResourceDiscoveryModeReEvaluateCompliance = ResourceDiscoveryMode("ReEvaluateCompliance")
)

func init() {
}
