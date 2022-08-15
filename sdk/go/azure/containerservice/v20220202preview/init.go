


package v20220202preview

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
	case "azure-native:containerservice/v20220202preview:AgentPool":
		r = &AgentPool{}
	case "azure-native:containerservice/v20220202preview:MaintenanceConfiguration":
		r = &MaintenanceConfiguration{}
	case "azure-native:containerservice/v20220202preview:ManagedCluster":
		r = &ManagedCluster{}
	case "azure-native:containerservice/v20220202preview:ManagedClusterSnapshot":
		r = &ManagedClusterSnapshot{}
	case "azure-native:containerservice/v20220202preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:containerservice/v20220202preview:Snapshot":
		r = &Snapshot{}
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
		"containerservice/v20220202preview",
		&module{version},
	)
}
