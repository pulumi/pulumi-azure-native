


package v20220715preview

type DiskMode string

const (
	DiskModePersistent                 = DiskMode("persistent")
	DiskMode_Independent_persistent    = DiskMode("independent_persistent")
	DiskMode_Independent_nonpersistent = DiskMode("independent_nonpersistent")
)

type DiskType string

const (
	DiskTypeFlat        = DiskType("flat")
	DiskTypePmem        = DiskType("pmem")
	DiskTypeRawphysical = DiskType("rawphysical")
	DiskTypeRawvirtual  = DiskType("rawvirtual")
	DiskTypeSparse      = DiskType("sparse")
	DiskTypeSesparse    = DiskType("sesparse")
	DiskTypeUnknown     = DiskType("unknown")
)

type FirmwareType string

const (
	FirmwareTypeBios = FirmwareType("bios")
	FirmwareTypeEfi  = FirmwareType("efi")
)

type IPAddressAllocationMethod string

const (
	IPAddressAllocationMethodUnset     = IPAddressAllocationMethod("unset")
	IPAddressAllocationMethodDynamic   = IPAddressAllocationMethod("dynamic")
	IPAddressAllocationMethodStatic    = IPAddressAllocationMethod("static")
	IPAddressAllocationMethodLinklayer = IPAddressAllocationMethod("linklayer")
	IPAddressAllocationMethodRandom    = IPAddressAllocationMethod("random")
	IPAddressAllocationMethodOther     = IPAddressAllocationMethod("other")
)

type IdentityType string

const (
	IdentityTypeNone           = IdentityType("None")
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
)

type InventoryType string

const (
	InventoryTypeResourcePool           = InventoryType("ResourcePool")
	InventoryTypeVirtualMachine         = InventoryType("VirtualMachine")
	InventoryTypeVirtualMachineTemplate = InventoryType("VirtualMachineTemplate")
	InventoryTypeVirtualNetwork         = InventoryType("VirtualNetwork")
	InventoryTypeCluster                = InventoryType("Cluster")
	InventoryTypeDatastore              = InventoryType("Datastore")
	InventoryTypeHost                   = InventoryType("Host")
)

type NICType string

const (
	NICTypeVmxnet3 = NICType("vmxnet3")
	NICTypeVmxnet2 = NICType("vmxnet2")
	NICTypeVmxnet  = NICType("vmxnet")
	NICTypeE1000   = NICType("e1000")
	NICTypeE1000e  = NICType("e1000e")
	NICTypePcnet32 = NICType("pcnet32")
)

type OsType string

const (
	OsTypeWindows = OsType("Windows")
	OsTypeLinux   = OsType("Linux")
	OsTypeOther   = OsType("Other")
)

type PowerOnBootOption string

const (
	PowerOnBootOptionEnabled  = PowerOnBootOption("enabled")
	PowerOnBootOptionDisabled = PowerOnBootOption("disabled")
)

type ProvisioningAction string

const (
	ProvisioningActionInstall   = ProvisioningAction("install")
	ProvisioningActionUninstall = ProvisioningAction("uninstall")
	ProvisioningActionRepair    = ProvisioningAction("repair")
)

func init() {
}
