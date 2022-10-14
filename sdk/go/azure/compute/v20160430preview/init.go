


package v20160430preview

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
	case "azure-native:compute/v20160430preview:AvailabilitySet":
		r = &AvailabilitySet{}
	case "azure-native:compute/v20160430preview:Disk":
		r = &Disk{}
	case "azure-native:compute/v20160430preview:Image":
		r = &Image{}
	case "azure-native:compute/v20160430preview:Snapshot":
		r = &Snapshot{}
	case "azure-native:compute/v20160430preview:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:compute/v20160430preview:VirtualMachineExtension":
		r = &VirtualMachineExtension{}
	case "azure-native:compute/v20160430preview:VirtualMachineScaleSet":
		r = &VirtualMachineScaleSet{}
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
		"compute/v20160430preview",
		&module{version},
	)
}
