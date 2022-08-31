


package v20180901

type ContainerGroupIpAddressType string

const (
	ContainerGroupIpAddressTypePublic  = ContainerGroupIpAddressType("Public")
	ContainerGroupIpAddressTypePrivate = ContainerGroupIpAddressType("Private")
)

type ContainerGroupNetworkProtocol string

const (
	ContainerGroupNetworkProtocolTCP = ContainerGroupNetworkProtocol("TCP")
	ContainerGroupNetworkProtocolUDP = ContainerGroupNetworkProtocol("UDP")
)

type ContainerGroupRestartPolicy string

const (
	ContainerGroupRestartPolicyAlways    = ContainerGroupRestartPolicy("Always")
	ContainerGroupRestartPolicyOnFailure = ContainerGroupRestartPolicy("OnFailure")
	ContainerGroupRestartPolicyNever     = ContainerGroupRestartPolicy("Never")
)

type ContainerNetworkProtocol string

const (
	ContainerNetworkProtocolTCP = ContainerNetworkProtocol("TCP")
	ContainerNetworkProtocolUDP = ContainerNetworkProtocol("UDP")
)

type LogAnalyticsLogType string

const (
	LogAnalyticsLogTypeContainerInsights     = LogAnalyticsLogType("ContainerInsights")
	LogAnalyticsLogTypeContainerInstanceLogs = LogAnalyticsLogType("ContainerInstanceLogs")
)

type OperatingSystemTypes string

const (
	OperatingSystemTypesWindows = OperatingSystemTypes("Windows")
	OperatingSystemTypesLinux   = OperatingSystemTypes("Linux")
)

func init() {
}
