


package v20211201

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
	case "azure-native:avs/v20211201:Addon":
		r = &Addon{}
	case "azure-native:avs/v20211201:Authorization":
		r = &Authorization{}
	case "azure-native:avs/v20211201:CloudLink":
		r = &CloudLink{}
	case "azure-native:avs/v20211201:Cluster":
		r = &Cluster{}
	case "azure-native:avs/v20211201:Datastore":
		r = &Datastore{}
	case "azure-native:avs/v20211201:GlobalReachConnection":
		r = &GlobalReachConnection{}
	case "azure-native:avs/v20211201:HcxEnterpriseSite":
		r = &HcxEnterpriseSite{}
	case "azure-native:avs/v20211201:PlacementPolicy":
		r = &PlacementPolicy{}
	case "azure-native:avs/v20211201:PrivateCloud":
		r = &PrivateCloud{}
	case "azure-native:avs/v20211201:ScriptExecution":
		r = &ScriptExecution{}
	case "azure-native:avs/v20211201:WorkloadNetworkDhcp":
		r = &WorkloadNetworkDhcp{}
	case "azure-native:avs/v20211201:WorkloadNetworkDnsService":
		r = &WorkloadNetworkDnsService{}
	case "azure-native:avs/v20211201:WorkloadNetworkDnsZone":
		r = &WorkloadNetworkDnsZone{}
	case "azure-native:avs/v20211201:WorkloadNetworkPortMirroring":
		r = &WorkloadNetworkPortMirroring{}
	case "azure-native:avs/v20211201:WorkloadNetworkPublicIP":
		r = &WorkloadNetworkPublicIP{}
	case "azure-native:avs/v20211201:WorkloadNetworkSegment":
		r = &WorkloadNetworkSegment{}
	case "azure-native:avs/v20211201:WorkloadNetworkVMGroup":
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
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"avs/v20211201",
		&module{version},
	)
}
