


package v20220901

type ComplianceState string

const (
	// The resource is in compliance with the policy.
	ComplianceStateCompliant = ComplianceState("Compliant")
	// The resource is not in compliance with the policy.
	ComplianceStateNonCompliant = ComplianceState("NonCompliant")
	// The compliance state of the resource is not known.
	ComplianceStateUnknown = ComplianceState("Unknown")
)

func init() {
}
