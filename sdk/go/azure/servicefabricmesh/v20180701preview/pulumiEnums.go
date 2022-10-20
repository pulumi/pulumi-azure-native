


package v20180701preview

type DiagnosticsSinkKind string

const (
	// Indicates an invalid sink kind. All Service Fabric enumerations have the invalid type.
	DiagnosticsSinkKindInvalid = DiagnosticsSinkKind("Invalid")
	// Diagnostics settings for Geneva.
	DiagnosticsSinkKindAzureInternalMonitoringPipeline = DiagnosticsSinkKind("AzureInternalMonitoringPipeline")
)

type HealthState string

const (
	// Indicates an invalid health state. All Service Fabric enumerations have the invalid type. The value is zero.
	HealthStateInvalid = HealthState("Invalid")
	// Indicates the health state is okay. The value is 1.
	HealthStateOk = HealthState("Ok")
	// Indicates the health state is at a warning level. The value is 2.
	HealthStateWarning = HealthState("Warning")
	// Indicates the health state is at an error level. Error health state should be investigated, as they can impact the correct functionality of the cluster. The value is 3.
	HealthStateError = HealthState("Error")
	// Indicates an unknown health status. The value is 65535.
	HealthStateUnknown = HealthState("Unknown")
)

type IngressQoSLevel string

const (
	IngressQoSLevelBronze = IngressQoSLevel("Bronze")
)

type OperatingSystemTypes string

const (
	OperatingSystemTypesLinux   = OperatingSystemTypes("Linux")
	OperatingSystemTypesWindows = OperatingSystemTypes("Windows")
)

type VolumeProvider string

const (
	VolumeProviderSFAzureFile = VolumeProvider("SFAzureFile")
)

func init() {
}
