


package v20220101preview

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
	case "azure-native:servicebus/v20220101preview:DisasterRecoveryConfig":
		r = &DisasterRecoveryConfig{}
	case "azure-native:servicebus/v20220101preview:MigrationConfig":
		r = &MigrationConfig{}
	case "azure-native:servicebus/v20220101preview:Namespace":
		r = &Namespace{}
	case "azure-native:servicebus/v20220101preview:NamespaceAuthorizationRule":
		r = &NamespaceAuthorizationRule{}
	case "azure-native:servicebus/v20220101preview:NamespaceNetworkRuleSet":
		r = &NamespaceNetworkRuleSet{}
	case "azure-native:servicebus/v20220101preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:servicebus/v20220101preview:Queue":
		r = &Queue{}
	case "azure-native:servicebus/v20220101preview:QueueAuthorizationRule":
		r = &QueueAuthorizationRule{}
	case "azure-native:servicebus/v20220101preview:Rule":
		r = &Rule{}
	case "azure-native:servicebus/v20220101preview:Subscription":
		r = &Subscription{}
	case "azure-native:servicebus/v20220101preview:Topic":
		r = &Topic{}
	case "azure-native:servicebus/v20220101preview:TopicAuthorizationRule":
		r = &TopicAuthorizationRule{}
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
		"servicebus/v20220101preview",
		&module{version},
	)
}
