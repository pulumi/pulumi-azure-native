


package v20180601

type ContainerGroupIpAddressType string

const (
	ContainerGroupIpAddressTypePublic = ContainerGroupIpAddressType("Public")
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

type OperatingSystemTypes string

const (
	OperatingSystemTypesWindows = OperatingSystemTypes("Windows")
	OperatingSystemTypesLinux   = OperatingSystemTypes("Linux")
)

func init() {
}
