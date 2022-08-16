


package v20201019preview

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
	case "azure-native:desktopvirtualization/v20201019preview:Application":
		r = &Application{}
	case "azure-native:desktopvirtualization/v20201019preview:ApplicationGroup":
		r = &ApplicationGroup{}
	case "azure-native:desktopvirtualization/v20201019preview:HostPool":
		r = &HostPool{}
	case "azure-native:desktopvirtualization/v20201019preview:MSIXPackage":
		r = &MSIXPackage{}
	case "azure-native:desktopvirtualization/v20201019preview:Workspace":
		r = &Workspace{}
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
		"desktopvirtualization/v20201019preview",
		&module{version},
	)
}
