


package v20220501

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
	case "azure-native:avs/v20220501:Addon":
		r = &Addon{}
	case "azure-native:avs/v20220501:Authorization":
		r = &Authorization{}
	case "azure-native:avs/v20220501:CloudLink":
		r = &CloudLink{}
	case "azure-native:avs/v20220501:Cluster":
		r = &Cluster{}
	case "azure-native:avs/v20220501:Datastore":
		r = &Datastore{}
	case "azure-native:avs/v20220501:GlobalReachConnection":
		r = &GlobalReachConnection{}
	case "azure-native:avs/v20220501:HcxEnterpriseSite":
		r = &HcxEnterpriseSite{}
	case "azure-native:avs/v20220501:PlacementPolicy":
		r = &PlacementPolicy{}
	case "azure-native:avs/v20220501:PrivateCloud":
		r = &PrivateCloud{}
	case "azure-native:avs/v20220501:ScriptExecution":
		r = &ScriptExecution{}
	case "azure-native:avs/v20220501:WorkloadNetworkDhcp":
		r = &WorkloadNetworkDhcp{}
	case "azure-native:avs/v20220501:WorkloadNetworkDnsService":
		r = &WorkloadNetworkDnsService{}
	case "azure-native:avs/v20220501:WorkloadNetworkDnsZone":
		r = &WorkloadNetworkDnsZone{}
	case "azure-native:avs/v20220501:WorkloadNetworkPortMirroring":
		r = &WorkloadNetworkPortMirroring{}
	case "azure-native:avs/v20220501:WorkloadNetworkPublicIP":
		r = &WorkloadNetworkPublicIP{}
	case "azure-native:avs/v20220501:WorkloadNetworkSegment":
		r = &WorkloadNetworkSegment{}
	case "azure-native:avs/v20220501:WorkloadNetworkVMGroup":
		r = &WorkloadNetworkVMGroup{}
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
		"avs/v20220501",
		&module{version},
	)
}
