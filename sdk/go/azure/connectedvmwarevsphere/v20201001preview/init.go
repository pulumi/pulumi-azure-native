


package v20201001preview

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
	case "azure-native:connectedvmwarevsphere/v20201001preview:Cluster":
		r = &Cluster{}
	case "azure-native:connectedvmwarevsphere/v20201001preview:Datastore":
		r = &Datastore{}
	case "azure-native:connectedvmwarevsphere/v20201001preview:GuestAgent":
		r = &GuestAgent{}
	case "azure-native:connectedvmwarevsphere/v20201001preview:Host":
		r = &Host{}
	case "azure-native:connectedvmwarevsphere/v20201001preview:HybridIdentityMetadatum":
		r = &HybridIdentityMetadatum{}
	case "azure-native:connectedvmwarevsphere/v20201001preview:InventoryItem":
		r = &InventoryItem{}
	case "azure-native:connectedvmwarevsphere/v20201001preview:MachineExtension":
		r = &MachineExtension{}
	case "azure-native:connectedvmwarevsphere/v20201001preview:ResourcePool":
		r = &ResourcePool{}
	case "azure-native:connectedvmwarevsphere/v20201001preview:VCenter":
		r = &VCenter{}
	case "azure-native:connectedvmwarevsphere/v20201001preview:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:connectedvmwarevsphere/v20201001preview:VirtualMachineTemplate":
		r = &VirtualMachineTemplate{}
	case "azure-native:connectedvmwarevsphere/v20201001preview:VirtualNetwork":
		r = &VirtualNetwork{}
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
		"connectedvmwarevsphere/v20201001preview",
		&module{version},
	)
}
