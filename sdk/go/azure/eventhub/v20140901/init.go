


package v20140901

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
	case "azure-native:eventhub/v20140901:ConsumerGroup":
		r = &ConsumerGroup{}
	case "azure-native:eventhub/v20140901:EventHub":
		r = &EventHub{}
	case "azure-native:eventhub/v20140901:EventHubAuthorizationRule":
		r = &EventHubAuthorizationRule{}
	case "azure-native:eventhub/v20140901:Namespace":
		r = &Namespace{}
	case "azure-native:eventhub/v20140901:NamespaceAuthorizationRule":
		r = &NamespaceAuthorizationRule{}
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
		"eventhub/v20140901",
		&module{version},
	)
}
