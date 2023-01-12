


package v20220802preview

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
	case "azure-native:containerservice/v20220802preview:AgentPool":
		r = &AgentPool{}
	case "azure-native:containerservice/v20220802preview:MaintenanceConfiguration":
		r = &MaintenanceConfiguration{}
	case "azure-native:containerservice/v20220802preview:ManagedCluster":
		r = &ManagedCluster{}
	case "azure-native:containerservice/v20220802preview:ManagedClusterSnapshot":
		r = &ManagedClusterSnapshot{}
	case "azure-native:containerservice/v20220802preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:containerservice/v20220802preview:Snapshot":
		r = &Snapshot{}
	case "azure-native:containerservice/v20220802preview:TrustedAccessRoleBinding":
		r = &TrustedAccessRoleBinding{}
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
		"containerservice/v20220802preview",
		&module{version},
	)
}
