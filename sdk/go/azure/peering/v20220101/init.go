


package v20220101

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
	case "azure-native:peering/v20220101:ConnectionMonitorTest":
		r = &ConnectionMonitorTest{}
	case "azure-native:peering/v20220101:PeerAsn":
		r = &PeerAsn{}
	case "azure-native:peering/v20220101:Peering":
		r = &Peering{}
	case "azure-native:peering/v20220101:PeeringService":
		r = &PeeringService{}
	case "azure-native:peering/v20220101:Prefix":
		r = &Prefix{}
	case "azure-native:peering/v20220101:RegisteredAsn":
		r = &RegisteredAsn{}
	case "azure-native:peering/v20220101:RegisteredPrefix":
		r = &RegisteredPrefix{}
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
		"peering/v20220101",
		&module{version},
	)
}
