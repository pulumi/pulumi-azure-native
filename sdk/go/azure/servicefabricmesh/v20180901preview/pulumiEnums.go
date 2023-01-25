


package v20180901preview

type ApplicationScopedVolumeKind string

const (
	// Provides Service Fabric High Availability Volume Disk
	ApplicationScopedVolumeKindServiceFabricVolumeDisk = ApplicationScopedVolumeKind("ServiceFabricVolumeDisk")
)

type AutoScalingMechanismKind string

const (
	// Indicates that scaling should be performed by adding or removing replicas.
	AutoScalingMechanismKindAddRemoveReplica = AutoScalingMechanismKind("AddRemoveReplica")
)

type AutoScalingMetricKind string

const (
	// Indicates that the metric is one of resources, like cpu or memory.
	AutoScalingMetricKindResource = AutoScalingMetricKind("Resource")
)

type AutoScalingResourceMetricName string

const (
	// Indicates that the resource is CPU cores.
	AutoScalingResourceMetricNameCpu = AutoScalingResourceMetricName("cpu")
	// Indicates that the resource is memory in GB.
	AutoScalingResourceMetricNameMemoryInGB = AutoScalingResourceMetricName("memoryInGB")
)

type AutoScalingTriggerKind string

const (
	// Indicates that scaling should be performed based on average load of all replicas in the service.
	AutoScalingTriggerKindAverageLoad = AutoScalingTriggerKind("AverageLoad")
)

type DiagnosticsSinkKind string

const (
	// Indicates an invalid sink kind. All Service Fabric enumerations have the invalid type.
	DiagnosticsSinkKindInvalid = DiagnosticsSinkKind("Invalid")
	// Diagnostics settings for Geneva.
	DiagnosticsSinkKindAzureInternalMonitoringPipeline = DiagnosticsSinkKind("AzureInternalMonitoringPipeline")
)

type HeaderMatchType string

const (
	HeaderMatchTypeExact = HeaderMatchType("exact")
)

type NetworkKind string

const (
	// Indicates a container network local to a single Service Fabric cluster. The value is 1.
	NetworkKindLocal = NetworkKind("Local")
)

type OperatingSystemType string

const (
	// The required operating system is Linux.
	OperatingSystemTypeLinux = OperatingSystemType("Linux")
	// The required operating system is Windows.
	OperatingSystemTypeWindows = OperatingSystemType("Windows")
)

type PathMatchType string

const (
	PathMatchTypePrefix = PathMatchType("prefix")
)

type SecretKind string

const (
	// A simple secret resource whose plaintext value is provided by the user.
	SecretKindInlinedValue = SecretKind("inlinedValue")
)

type SizeTypes string

const (
	SizeTypesSmall  = SizeTypes("Small")
	SizeTypesMedium = SizeTypes("Medium")
	SizeTypesLarge  = SizeTypes("Large")
)

type VolumeProvider string

const (
	// Provides volumes that are backed by Azure Files.
	VolumeProviderSFAzureFile = VolumeProvider("SFAzureFile")
)

func init() {
}
