


package v20180601preview

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
	case "azure-native:dbformariadb/v20180601preview:Configuration":
		r = &Configuration{}
	case "azure-native:dbformariadb/v20180601preview:Database":
		r = &Database{}
	case "azure-native:dbformariadb/v20180601preview:FirewallRule":
		r = &FirewallRule{}
	case "azure-native:dbformariadb/v20180601preview:Server":
		r = &Server{}
	case "azure-native:dbformariadb/v20180601preview:VirtualNetworkRule":
		r = &VirtualNetworkRule{}
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
		"dbformariadb/v20180601preview",
		&module{version},
	)
}
