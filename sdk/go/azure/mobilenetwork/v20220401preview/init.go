


package v20220401preview

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
	case "azure-native:mobilenetwork/v20220401preview:AttachedDataNetwork":
		r = &AttachedDataNetwork{}
	case "azure-native:mobilenetwork/v20220401preview:DataNetwork":
		r = &DataNetwork{}
	case "azure-native:mobilenetwork/v20220401preview:MobileNetwork":
		r = &MobileNetwork{}
	case "azure-native:mobilenetwork/v20220401preview:PacketCoreControlPlane":
		r = &PacketCoreControlPlane{}
	case "azure-native:mobilenetwork/v20220401preview:PacketCoreDataPlane":
		r = &PacketCoreDataPlane{}
	case "azure-native:mobilenetwork/v20220401preview:Service":
		r = &Service{}
	case "azure-native:mobilenetwork/v20220401preview:Sim":
		r = &Sim{}
	case "azure-native:mobilenetwork/v20220401preview:SimGroup":
		r = &SimGroup{}
	case "azure-native:mobilenetwork/v20220401preview:SimPolicy":
		r = &SimPolicy{}
	case "azure-native:mobilenetwork/v20220401preview:Site":
		r = &Site{}
	case "azure-native:mobilenetwork/v20220401preview:Slice":
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
		"mobilenetwork/v20220401preview",
		&module{version},
	)
}
