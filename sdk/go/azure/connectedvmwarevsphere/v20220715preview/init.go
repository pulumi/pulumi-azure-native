


package v20220715preview

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
	case "azure-native:connectedvmwarevsphere/v20220715preview:Cluster":
		r = &Cluster{}
	case "azure-native:connectedvmwarevsphere/v20220715preview:Datastore":
		r = &Datastore{}
	case "azure-native:connectedvmwarevsphere/v20220715preview:GuestAgent":
		r = &GuestAgent{}
	case "azure-native:connectedvmwarevsphere/v20220715preview:Host":
		r = &Host{}
	case "azure-native:connectedvmwarevsphere/v20220715preview:HybridIdentityMetadatum":
		r = &HybridIdentityMetadatum{}
	case "azure-native:connectedvmwarevsphere/v20220715preview:InventoryItem":
		r = &InventoryItem{}
	case "azure-native:connectedvmwarevsphere/v20220715preview:MachineExtension":
		r = &MachineExtension{}
	case "azure-native:connectedvmwarevsphere/v20220715preview:ResourcePool":
		r = &ResourcePool{}
	case "azure-native:connectedvmwarevsphere/v20220715preview:VCenter":
		r = &VCenter{}
	case "azure-native:connectedvmwarevsphere/v20220715preview:VirtualMachine":
		r = &VirtualMachine{}
	case "azure-native:connectedvmwarevsphere/v20220715preview:VirtualMachineTemplate":
		r = &VirtualMachineTemplate{}
	case "azure-native:connectedvmwarevsphere/v20220715preview:VirtualNetwork":
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
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"connectedvmwarevsphere/v20220715preview",
		&module{version},
	)
}
