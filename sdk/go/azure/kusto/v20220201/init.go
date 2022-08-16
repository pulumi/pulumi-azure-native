


package v20220201

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
	case "azure-native:kusto/v20220201:AttachedDatabaseConfiguration":
		r = &AttachedDatabaseConfiguration{}
	case "azure-native:kusto/v20220201:Cluster":
		r = &Cluster{}
	case "azure-native:kusto/v20220201:ClusterPrincipalAssignment":
		r = &ClusterPrincipalAssignment{}
	case "azure-native:kusto/v20220201:DataConnection":
		r = &DataConnection{}
	case "azure-native:kusto/v20220201:Database":
		r = &Database{}
	case "azure-native:kusto/v20220201:DatabasePrincipalAssignment":
		r = &DatabasePrincipalAssignment{}
	case "azure-native:kusto/v20220201:EventGridDataConnection":
		r = &EventGridDataConnection{}
	case "azure-native:kusto/v20220201:EventHubDataConnection":
		r = &EventHubDataConnection{}
	case "azure-native:kusto/v20220201:IotHubDataConnection":
		r = &IotHubDataConnection{}
	case "azure-native:kusto/v20220201:ManagedPrivateEndpoint":
		r = &ManagedPrivateEndpoint{}
	case "azure-native:kusto/v20220201:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:kusto/v20220201:ReadOnlyFollowingDatabase":
		r = &ReadOnlyFollowingDatabase{}
	case "azure-native:kusto/v20220201:ReadWriteDatabase":
		r = &ReadWriteDatabase{}
	case "azure-native:kusto/v20220201:Script":
		r = &Script{}
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
		"kusto/v20220201",
		&module{version},
	)
}
