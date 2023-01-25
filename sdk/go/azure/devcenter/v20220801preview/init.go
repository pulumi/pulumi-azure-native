


package v20220801preview

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
	case "azure-native:devcenter/v20220801preview:AttachedNetworkByDevCenter":
		r = &AttachedNetworkByDevCenter{}
	case "azure-native:devcenter/v20220801preview:Catalog":
		r = &Catalog{}
	case "azure-native:devcenter/v20220801preview:DevBoxDefinition":
		r = &DevBoxDefinition{}
	case "azure-native:devcenter/v20220801preview:DevCenter":
		r = &DevCenter{}
	case "azure-native:devcenter/v20220801preview:EnvironmentType":
		r = &EnvironmentType{}
	case "azure-native:devcenter/v20220801preview:Gallery":
		r = &Gallery{}
	case "azure-native:devcenter/v20220801preview:NetworkConnection":
		r = &NetworkConnection{}
	case "azure-native:devcenter/v20220801preview:Pool":
		r = &Pool{}
	case "azure-native:devcenter/v20220801preview:Project":
		r = &Project{}
	case "azure-native:devcenter/v20220801preview:ProjectEnvironmentType":
		r = &ProjectEnvironmentType{}
	case "azure-native:devcenter/v20220801preview:Schedule":
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
		"devcenter/v20220801preview",
		&module{version},
	)
}
