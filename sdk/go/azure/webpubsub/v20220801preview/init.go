


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
	case "azure-native:webpubsub/v20220801preview:WebPubSub":
		r = &WebPubSub{}
	case "azure-native:webpubsub/v20220801preview:WebPubSubCustomCertificate":
		r = &WebPubSubCustomCertificate{}
	case "azure-native:webpubsub/v20220801preview:WebPubSubCustomDomain":
		r = &WebPubSubCustomDomain{}
	case "azure-native:webpubsub/v20220801preview:WebPubSubHub":
		r = &WebPubSubHub{}
	case "azure-native:webpubsub/v20220801preview:WebPubSubPrivateEndpointConnection":
		r = &WebPubSubPrivateEndpointConnection{}
	case "azure-native:webpubsub/v20220801preview:WebPubSubSharedPrivateLinkResource":
		r = &WebPubSubSharedPrivateLinkResource{}
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
		"webpubsub/v20220801preview",
		&module{version},
	)
}
