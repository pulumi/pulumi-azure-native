


package v20220801

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:compute/v20220801:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:compute/v20220801:CapacityReservation":
		r = &CapacityReservation{}
	case "azure-native:compute/v20220801:CapacityReservationGroup":
		r = &CapacityReservationGroup{}
	case "azure-native:compute/v20220801:DedicatedHost":
		r = &DedicatedHost{}
	case "azure-native:compute/v20220801:DedicatedHostGroup":
		r = &DedicatedHostGroup{}
	case "azure-native:compute/v20220801:Image":
		r = &Image{}
	case "azure-native:compute/v20220801:ProximityPlacementGroup":
		r = &ProximityPlacementGroup{}
	case "azure-native:compute/v20220801:RestorePoint":
		r = &RestorePoint{}
	case "azure-native:compute/v20220801:RestorePointCollection":
		r = &RestorePointCollection{}
	case "azure-native:compute/v20220801:SshPublicKey":
		r = &SshPublicKey{}
	case "azure-native:compute/v20220801:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:compute/v20220801:VirtualMachineExtension":
		r = &VirtualMachineExtension{}
	case "azure-native:compute/v20220801:VirtualMachineRunCommandByVirtualMachine":
		r = &VirtualMachineRunCommandByVirtualMachine{}
	case "azure-native:compute/v20220801:VirtualMachineScaleSet":
		r = &VirtualMachineScaleSet{}
	case "azure-native:compute/v20220801:VirtualMachineScaleSetExtension":
		r = &VirtualMachineScaleSetExtension{}
	case "azure-native:compute/v20220801:VirtualMachineScaleSetVM":
		r = &VirtualMachineScaleSetVM{}
	case "azure-native:compute/v20220801:VirtualMachineScaleSetVMExtension":
		r = &VirtualMachineScaleSetVMExtension{}
	case "azure-native:compute/v20220801:VirtualMachineScaleSetVMRunCommand":
		r = &VirtualMachineScaleSetVMRunCommand{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"compute/v20220801",
		&module{version},
	)
}
