


package v20170801preview

type ContainerGroupIpAddressType string

const (
	ContainerGroupIpAddressTypePublic = ContainerGroupIpAddressType("Public")
)

type ContainerGroupNetworkProtocol string

const (
	ContainerGroupNetworkProtocolTCP = ContainerGroupNetworkProtocol("TCP")
	ContainerGroupNetworkProtocolUDP = ContainerGroupNetworkProtocol("UDP")
)

type ContainerRestartPolicy string

const (
	ContainerRestartPolicyAlways = ContainerRestartPolicy("always")
)

type OperatingSystemTypes string

const (
	OperatingSystemTypesWindows = OperatingSystemTypes("Windows")
	OperatingSystemTypesLinux   = OperatingSystemTypes("Linux")
)

func init() {
}
