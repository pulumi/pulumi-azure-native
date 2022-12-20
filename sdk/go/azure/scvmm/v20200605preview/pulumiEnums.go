


package v20200605preview

type AllocationMethod string

const (
	AllocationMethodDynamic = AllocationMethod("Dynamic")
	AllocationMethodStatic  = AllocationMethod("Static")
)

type CreateDiffDisk string

const (
	CreateDiffDiskFalse = CreateDiffDisk("false")
	CreateDiffDiskTrue  = CreateDiffDisk("true")
)

type DynamicMemoryEnabled string

const (
	DynamicMemoryEnabledFalse = DynamicMemoryEnabled("false")
	DynamicMemoryEnabledTrue  = DynamicMemoryEnabled("true")
)

type InventoryType string

const (
	InventoryTypeCloud                  = InventoryType("Cloud")
	InventoryTypeVirtualNetwork         = InventoryType("VirtualNetwork")
	InventoryTypeVirtualMachineTemplate = InventoryType("VirtualMachineTemplate")
	InventoryTypeVirtualMachine         = InventoryType("VirtualMachine")
)

type LimitCpuForMigration string

const (
	LimitCpuForMigrationFalse = LimitCpuForMigration("false")
	LimitCpuForMigrationTrue  = LimitCpuForMigration("true")
)

func init() {
}
