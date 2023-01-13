


package v20221101

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
	case "azure-native:mobilenetwork/v20221101:AttachedDataNetwork":
		r = &AttachedDataNetwork{}
	case "azure-native:mobilenetwork/v20221101:DataNetwork":
		r = &DataNetwork{}
	case "azure-native:mobilenetwork/v20221101:MobileNetwork":
		r = &MobileNetwork{}
	case "azure-native:mobilenetwork/v20221101:PacketCoreControlPlane":
		r = &PacketCoreControlPlane{}
	case "azure-native:mobilenetwork/v20221101:PacketCoreDataPlane":
		r = &PacketCoreDataPlane{}
	case "azure-native:mobilenetwork/v20221101:Service":
		r = &Service{}
	case "azure-native:mobilenetwork/v20221101:Sim":
		r = &Sim{}
	case "azure-native:mobilenetwork/v20221101:SimGroup":
		r = &SimGroup{}
	case "azure-native:mobilenetwork/v20221101:SimPolicy":
		r = &SimPolicy{}
	case "azure-native:mobilenetwork/v20221101:Site":
		r = &Site{}
	case "azure-native:mobilenetwork/v20221101:Slice":
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
		"mobilenetwork/v20221101",
		&module{version},
	)
}
