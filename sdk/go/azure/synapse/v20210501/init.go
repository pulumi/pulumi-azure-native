


package v20210501

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
	case "azure-native:synapse/v20210501:BigDataPool":
		r = &BigDataPool{}
	case "azure-native:synapse/v20210501:IntegrationRuntime":
		r = &IntegrationRuntime{}
	case "azure-native:synapse/v20210501:IpFirewallRule":
		r = &IpFirewallRule{}
	case "azure-native:synapse/v20210501:Key":
		r = &Key{}
	case "azure-native:synapse/v20210501:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:synapse/v20210501:PrivateLinkHub":
		r = &PrivateLinkHub{}
	case "azure-native:synapse/v20210501:SqlPool":
		r = &SqlPool{}
	case "azure-native:synapse/v20210501:SqlPoolSensitivityLabel":
		r = &SqlPoolSensitivityLabel{}
	case "azure-native:synapse/v20210501:SqlPoolTransparentDataEncryption":
		r = &SqlPoolTransparentDataEncryption{}
	case "azure-native:synapse/v20210501:SqlPoolVulnerabilityAssessment":
		r = &SqlPoolVulnerabilityAssessment{}
	case "azure-native:synapse/v20210501:SqlPoolVulnerabilityAssessmentRuleBaseline":
		r = &SqlPoolVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:synapse/v20210501:SqlPoolWorkloadClassifier":
		r = &SqlPoolWorkloadClassifier{}
	case "azure-native:synapse/v20210501:SqlPoolWorkloadGroup":
		r = &SqlPoolWorkloadGroup{}
	case "azure-native:synapse/v20210501:Workspace":
		r = &Workspace{}
	case "azure-native:synapse/v20210501:WorkspaceAadAdmin":
		r = &WorkspaceAadAdmin{}
	case "azure-native:synapse/v20210501:WorkspaceManagedSqlServerVulnerabilityAssessment":
		r = &WorkspaceManagedSqlServerVulnerabilityAssessment{}
	case "azure-native:synapse/v20210501:WorkspaceSqlAadAdmin":
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
		"synapse/v20210501",
		&module{version},
	)
}
