


package v20220901preview

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
	case "azure-native:devcenter/v20220901preview:AttachedNetworkByDevCenter":
		r = &AttachedNetworkByDevCenter{}
	case "azure-native:devcenter/v20220901preview:Catalog":
		r = &Catalog{}
	case "azure-native:devcenter/v20220901preview:DevBoxDefinition":
		r = &DevBoxDefinition{}
	case "azure-native:devcenter/v20220901preview:DevCenter":
		r = &DevCenter{}
	case "azure-native:devcenter/v20220901preview:EnvironmentType":
		r = &EnvironmentType{}
	case "azure-native:devcenter/v20220901preview:Gallery":
		r = &Gallery{}
	case "azure-native:devcenter/v20220901preview:NetworkConnection":
		r = &NetworkConnection{}
	case "azure-native:devcenter/v20220901preview:Pool":
		r = &Pool{}
	case "azure-native:devcenter/v20220901preview:Project":
		r = &Project{}
	case "azure-native:devcenter/v20220901preview:ProjectEnvironmentType":
		r = &ProjectEnvironmentType{}
	case "azure-native:devcenter/v20220901preview:Schedule":
		r = &Schedule{}
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
		"devcenter/v20220901preview",
		&module{version},
	)
}
