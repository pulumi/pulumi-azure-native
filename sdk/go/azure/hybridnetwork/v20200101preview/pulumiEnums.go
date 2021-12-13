


package v20200101preview

type DeviceType string

const (
	DeviceTypeUnknown        = DeviceType("Unknown")
	DeviceTypeAzureStackEdge = DeviceType("AzureStackEdge")
)

type DiskCreateOptionTypes string

const (
	DiskCreateOptionTypesUnknown = DiskCreateOptionTypes("Unknown")
	DiskCreateOptionTypesEmpty   = DiskCreateOptionTypes("Empty")
)

type IPAllocationMethod string

const (
	IPAllocationMethodUnknown = IPAllocationMethod("Unknown")
	IPAllocationMethodStatic  = IPAllocationMethod("Static")
	IPAllocationMethodDynamic = IPAllocationMethod("Dynamic")
)

type IPVersion string

const (
	IPVersionUnknown = IPVersion("Unknown")
	IPVersionIPv4    = IPVersion("IPv4")
)

type NetworkFunctionRoleConfigurationType string

const (
	NetworkFunctionRoleConfigurationTypeUnknown        = NetworkFunctionRoleConfigurationType("Unknown")
	NetworkFunctionRoleConfigurationTypeVirtualMachine = NetworkFunctionRoleConfigurationType("VirtualMachine")
)

type OperatingSystemTypes string

const (
	OperatingSystemTypesUnknown = OperatingSystemTypes("Unknown")
	OperatingSystemTypesWindows = OperatingSystemTypes("Windows")
	OperatingSystemTypesLinux   = OperatingSystemTypes("Linux")
)

type SkuDeploymentMode string

const (
	SkuDeploymentModeUnknown         = SkuDeploymentMode("Unknown")
	SkuDeploymentModeAzure           = SkuDeploymentMode("Azure")
	SkuDeploymentModePrivateEdgeZone = SkuDeploymentMode("PrivateEdgeZone")
)

type SkuType string

const (
	SkuTypeUnknown           = SkuType("Unknown")
	SkuTypeEvolvedPacketCore = SkuType("EvolvedPacketCore")
	SkuTypeSDWAN             = SkuType("SDWAN")
	SkuTypeFirewall          = SkuType("Firewall")
)

type VMSwitchType string

const (
	VMSwitchTypeUnknown    = VMSwitchType("Unknown")
	VMSwitchTypeManagement = VMSwitchType("Management")
	VMSwitchTypeWan        = VMSwitchType("Wan")
	VMSwitchTypeLan        = VMSwitchType("Lan")
)

type VirtualMachineSizeTypes string

const (
	VirtualMachineSizeTypesUnknown           = VirtualMachineSizeTypes("Unknown")
	VirtualMachineSizeTypes_Standard_D1_v2   = VirtualMachineSizeTypes("Standard_D1_v2")
	VirtualMachineSizeTypes_Standard_D2_v2   = VirtualMachineSizeTypes("Standard_D2_v2")
	VirtualMachineSizeTypes_Standard_D3_v2   = VirtualMachineSizeTypes("Standard_D3_v2")
	VirtualMachineSizeTypes_Standard_D4_v2   = VirtualMachineSizeTypes("Standard_D4_v2")
	VirtualMachineSizeTypes_Standard_D5_v2   = VirtualMachineSizeTypes("Standard_D5_v2")
	VirtualMachineSizeTypes_Standard_D11_v2  = VirtualMachineSizeTypes("Standard_D11_v2")
	VirtualMachineSizeTypes_Standard_D12_v2  = VirtualMachineSizeTypes("Standard_D12_v2")
	VirtualMachineSizeTypes_Standard_D13_v2  = VirtualMachineSizeTypes("Standard_D13_v2")
	VirtualMachineSizeTypes_Standard_DS1_v2  = VirtualMachineSizeTypes("Standard_DS1_v2")
	VirtualMachineSizeTypes_Standard_DS2_v2  = VirtualMachineSizeTypes("Standard_DS2_v2")
	VirtualMachineSizeTypes_Standard_DS3_v2  = VirtualMachineSizeTypes("Standard_DS3_v2")
	VirtualMachineSizeTypes_Standard_DS4_v2  = VirtualMachineSizeTypes("Standard_DS4_v2")
	VirtualMachineSizeTypes_Standard_DS5_v2  = VirtualMachineSizeTypes("Standard_DS5_v2")
	VirtualMachineSizeTypes_Standard_DS11_v2 = VirtualMachineSizeTypes("Standard_DS11_v2")
	VirtualMachineSizeTypes_Standard_DS12_v2 = VirtualMachineSizeTypes("Standard_DS12_v2")
	VirtualMachineSizeTypes_Standard_DS13_v2 = VirtualMachineSizeTypes("Standard_DS13_v2")
	VirtualMachineSizeTypes_Standard_F1      = VirtualMachineSizeTypes("Standard_F1")
	VirtualMachineSizeTypes_Standard_F2      = VirtualMachineSizeTypes("Standard_F2")
	VirtualMachineSizeTypes_Standard_F4      = VirtualMachineSizeTypes("Standard_F4")
	VirtualMachineSizeTypes_Standard_F8      = VirtualMachineSizeTypes("Standard_F8")
	VirtualMachineSizeTypes_Standard_F16     = VirtualMachineSizeTypes("Standard_F16")
	VirtualMachineSizeTypes_Standard_F1s     = VirtualMachineSizeTypes("Standard_F1s")
	VirtualMachineSizeTypes_Standard_F2s     = VirtualMachineSizeTypes("Standard_F2s")
	VirtualMachineSizeTypes_Standard_F4s     = VirtualMachineSizeTypes("Standard_F4s")
	VirtualMachineSizeTypes_Standard_F8s     = VirtualMachineSizeTypes("Standard_F8s")
	VirtualMachineSizeTypes_Standard_F16s    = VirtualMachineSizeTypes("Standard_F16s")
)

func init() {
}
