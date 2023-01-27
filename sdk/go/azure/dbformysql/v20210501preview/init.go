


package v20210501preview

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
	case "azure-native:dbformysql/v20210501preview:Configuration":
		r = &Configuration{}
	case "azure-native:dbformysql/v20210501preview:Database":
		r = &Database{}
	case "azure-native:dbformysql/v20210501preview:FirewallRule":
		r = &FirewallRule{}
	case "azure-native:dbformysql/v20210501preview:Server":
		r = &Server{}
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
		"dbformysql/v20210501preview",
		&module{version},
	)
}
