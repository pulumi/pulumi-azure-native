


package v20180901preview

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
	case "azure-native:servicefabricmesh/v20180901preview:Application":
		r = &Application{}
	case "azure-native:servicefabricmesh/v20180901preview:Gateway":
		r = &Gateway{}
	case "azure-native:servicefabricmesh/v20180901preview:Network":
		r = &Network{}
	case "azure-native:servicefabricmesh/v20180901preview:Secret":
		r = &Secret{}
	case "azure-native:servicefabricmesh/v20180901preview:SecretValue":
		r = &SecretValue{}
	case "azure-native:servicefabricmesh/v20180901preview:Volume":
		r = &Volume{}
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
		"servicefabricmesh/v20180901preview",
		&module{version},
	)
}
