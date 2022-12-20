


package migrate

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
	case "azure-native:migrate:Assessment":
		r = &Assessment{}
	case "azure-native:migrate:Group":
		r = &Group{}
	case "azure-native:migrate:HyperVCollector":
		r = &HyperVCollector{}
	case "azure-native:migrate:ImportCollector":
		r = &ImportCollector{}
	case "azure-native:migrate:MigrateProject":
		r = &MigrateProject{}
	case "azure-native:migrate:MoveCollection":
		r = &MoveCollection{}
	case "azure-native:migrate:MoveResource":
		r = &MoveResource{}
	case "azure-native:migrate:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:migrate:Project":
		r = &Project{}
	case "azure-native:migrate:ServerCollector":
		r = &ServerCollector{}
	case "azure-native:migrate:Solution":
		r = &Solution{}
	case "azure-native:migrate:VMwareCollector":
		r = &VMwareCollector{}
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
		"migrate",
		&module{version},
	)
}
