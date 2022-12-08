


package v20201201

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
	case "azure-native:synapse/v20201201:BigDataPool":
		r = &BigDataPool{}
	case "azure-native:synapse/v20201201:IntegrationRuntime":
		r = &IntegrationRuntime{}
	case "azure-native:synapse/v20201201:IpFirewallRule":
		r = &IpFirewallRule{}
	case "azure-native:synapse/v20201201:Key":
		r = &Key{}
	case "azure-native:synapse/v20201201:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:synapse/v20201201:PrivateLinkHub":
		r = &PrivateLinkHub{}
	case "azure-native:synapse/v20201201:SqlPool":
		r = &SqlPool{}
	case "azure-native:synapse/v20201201:SqlPoolSensitivityLabel":
		r = &SqlPoolSensitivityLabel{}
	case "azure-native:synapse/v20201201:SqlPoolTransparentDataEncryption":
		r = &SqlPoolTransparentDataEncryption{}
	case "azure-native:synapse/v20201201:SqlPoolVulnerabilityAssessment":
		r = &SqlPoolVulnerabilityAssessment{}
	case "azure-native:synapse/v20201201:SqlPoolVulnerabilityAssessmentRuleBaseline":
		r = &SqlPoolVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:synapse/v20201201:SqlPoolWorkloadClassifier":
		r = &SqlPoolWorkloadClassifier{}
	case "azure-native:synapse/v20201201:SqlPoolWorkloadGroup":
		r = &SqlPoolWorkloadGroup{}
	case "azure-native:synapse/v20201201:Workspace":
		r = &Workspace{}
	case "azure-native:synapse/v20201201:WorkspaceAadAdmin":
		r = &WorkspaceAadAdmin{}
	case "azure-native:synapse/v20201201:WorkspaceManagedSqlServerVulnerabilityAssessment":
		r = &WorkspaceManagedSqlServerVulnerabilityAssessment{}
	case "azure-native:synapse/v20201201:WorkspaceSqlAadAdmin":
		r = &WorkspaceSqlAadAdmin{}
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
		"synapse/v20201201",
		&module{version},
	)
}
