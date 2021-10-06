


package v20210701

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
	case "azure-native:compute/v20210701:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:compute/v20210701:CapacityReservation":
		r = &CapacityReservation{}
	case "azure-native:compute/v20210701:CapacityReservationGroup":
		r = &CapacityReservationGroup{}
	case "azure-native:compute/v20210701:DedicatedHost":
		r = &DedicatedHost{}
	case "azure-native:compute/v20210701:DedicatedHostGroup":
		r = &DedicatedHostGroup{}
	case "azure-native:compute/v20210701:Gallery":
		r = &Gallery{}
	case "azure-native:compute/v20210701:GalleryApplication":
		r = &GalleryApplication{}
	case "azure-native:compute/v20210701:GalleryApplicationVersion":
		r = &GalleryApplicationVersion{}
	case "azure-native:compute/v20210701:GalleryImage":
		r = &GalleryImage{}
	case "azure-native:compute/v20210701:GalleryImageVersion":
		r = &GalleryImageVersion{}
	case "azure-native:compute/v20210701:Image":
		r = &Image{}
	case "azure-native:compute/v20210701:ProximityPlacementGroup":
		r = &ProximityPlacementGroup{}
	case "azure-native:compute/v20210701:RestorePoint":
		r = &RestorePoint{}
	case "azure-native:compute/v20210701:RestorePointCollection":
		r = &RestorePointCollection{}
	case "azure-native:compute/v20210701:SshPublicKey":
		r = &SshPublicKey{}
	case "azure-native:compute/v20210701:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:compute/v20210701:VirtualMachineExtension":
		r = &VirtualMachineExtension{}
	case "azure-native:compute/v20210701:VirtualMachineRunCommandByVirtualMachine":
		r = &VirtualMachineRunCommandByVirtualMachine{}
	case "azure-native:compute/v20210701:VirtualMachineScaleSet":
		r = &VirtualMachineScaleSet{}
	case "azure-native:compute/v20210701:VirtualMachineScaleSetExtension":
		r = &VirtualMachineScaleSetExtension{}
	case "azure-native:compute/v20210701:VirtualMachineScaleSetVM":
		r = &VirtualMachineScaleSetVM{}
	case "azure-native:compute/v20210701:VirtualMachineScaleSetVMExtension":
		r = &VirtualMachineScaleSetVMExtension{}
	case "azure-native:compute/v20210701:VirtualMachineScaleSetVMRunCommand":
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
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"compute/v20210701",
		&module{version},
	)
}
