


package v20220701

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
	case "azure-native:network/v20220701:DnsForwardingRuleset":
		r = &DnsForwardingRuleset{}
	case "azure-native:network/v20220701:DnsResolver":
		r = &DnsResolver{}
	case "azure-native:network/v20220701:ForwardingRule":
		r = &ForwardingRule{}
	case "azure-native:network/v20220701:InboundEndpoint":
		r = &InboundEndpoint{}
	case "azure-native:network/v20220701:OutboundEndpoint":
		r = &OutboundEndpoint{}
	case "azure-native:network/v20220701:VirtualNetworkLink":
		r = &VirtualNetworkLink{}
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
		"network/v20220701",
		&module{version},
	)
}
