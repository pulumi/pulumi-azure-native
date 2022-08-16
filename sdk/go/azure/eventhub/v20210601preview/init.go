


package v20210601preview

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
	case "azure-native:eventhub/v20210601preview:Cluster":
		r = &Cluster{}
	case "azure-native:eventhub/v20210601preview:ConsumerGroup":
		r = &ConsumerGroup{}
	case "azure-native:eventhub/v20210601preview:DisasterRecoveryConfig":
		r = &DisasterRecoveryConfig{}
	case "azure-native:eventhub/v20210601preview:EventHub":
		r = &EventHub{}
	case "azure-native:eventhub/v20210601preview:EventHubAuthorizationRule":
		r = &EventHubAuthorizationRule{}
	case "azure-native:eventhub/v20210601preview:Namespace":
		r = &Namespace{}
	case "azure-native:eventhub/v20210601preview:NamespaceAuthorizationRule":
		r = &NamespaceAuthorizationRule{}
	case "azure-native:eventhub/v20210601preview:NamespaceNetworkRuleSet":
		r = &NamespaceNetworkRuleSet{}
	case "azure-native:eventhub/v20210601preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
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
		"eventhub/v20210601preview",
		&module{version},
	)
}
