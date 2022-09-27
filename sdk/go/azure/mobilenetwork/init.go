


package mobilenetwork

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
	case "azure-native:mobilenetwork:AttachedDataNetwork":
		r = &AttachedDataNetwork{}
	case "azure-native:mobilenetwork:DataNetwork":
		r = &DataNetwork{}
	case "azure-native:mobilenetwork:MobileNetwork":
		r = &MobileNetwork{}
	case "azure-native:mobilenetwork:PacketCoreControlPlane":
		r = &PacketCoreControlPlane{}
	case "azure-native:mobilenetwork:PacketCoreDataPlane":
		r = &PacketCoreDataPlane{}
	case "azure-native:mobilenetwork:Service":
		r = &Service{}
	case "azure-native:mobilenetwork:Sim":
		r = &Sim{}
	case "azure-native:mobilenetwork:SimGroup":
		r = &SimGroup{}
	case "azure-native:mobilenetwork:SimPolicy":
		r = &SimPolicy{}
	case "azure-native:mobilenetwork:Site":
		r = &Site{}
	case "azure-native:mobilenetwork:Slice":
		r = &Slice{}
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
		"mobilenetwork",
		&module{version},
	)
}
