


package v20210610preview

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
	case "azure-native:hybridcompute/v20210610preview:Machine":
		r = &Machine{}
	case "azure-native:hybridcompute/v20210610preview:MachineExtension":
		r = &MachineExtension{}
	case "azure-native:hybridcompute/v20210610preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:hybridcompute/v20210610preview:PrivateLinkScope":
		r = &PrivateLinkScope{}
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
		"hybridcompute/v20210610preview",
		&module{version},
	)
}
