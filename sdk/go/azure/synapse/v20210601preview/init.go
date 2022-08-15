


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
	case "azure-native:synapse/v20210601preview:BigDataPool":
		r = &BigDataPool{}
	case "azure-native:synapse/v20210601preview:EventGridDataConnection":
		r = &EventGridDataConnection{}
	case "azure-native:synapse/v20210601preview:EventHubDataConnection":
		r = &EventHubDataConnection{}
	case "azure-native:synapse/v20210601preview:IntegrationRuntime":
		r = &IntegrationRuntime{}
	case "azure-native:synapse/v20210601preview:IotHubDataConnection":
		r = &IotHubDataConnection{}
	case "azure-native:synapse/v20210601preview:IpFirewallRule":
		r = &IpFirewallRule{}
	case "azure-native:synapse/v20210601preview:Key":
		r = &Key{}
	case "azure-native:synapse/v20210601preview:KustoPool":
		r = &KustoPool{}
	case "azure-native:synapse/v20210601preview:KustoPoolAttachedDatabaseConfiguration":
		r = &KustoPoolAttachedDatabaseConfiguration{}
	case "azure-native:synapse/v20210601preview:KustoPoolDataConnection":
		r = &KustoPoolDataConnection{}
	case "azure-native:synapse/v20210601preview:KustoPoolDatabase":
		r = &KustoPoolDatabase{}
	case "azure-native:synapse/v20210601preview:KustoPoolDatabasePrincipalAssignment":
		r = &KustoPoolDatabasePrincipalAssignment{}
	case "azure-native:synapse/v20210601preview:KustoPoolPrincipalAssignment":
		r = &KustoPoolPrincipalAssignment{}
	case "azure-native:synapse/v20210601preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:synapse/v20210601preview:PrivateLinkHub":
		r = &PrivateLinkHub{}
	case "azure-native:synapse/v20210601preview:ReadOnlyFollowingDatabase":
		r = &ReadOnlyFollowingDatabase{}
	case "azure-native:synapse/v20210601preview:ReadWriteDatabase":
		r = &ReadWriteDatabase{}
	case "azure-native:synapse/v20210601preview:SqlPool":
		r = &SqlPool{}
	case "azure-native:synapse/v20210601preview:SqlPoolSensitivityLabel":
		r = &SqlPoolSensitivityLabel{}
	case "azure-native:synapse/v20210601preview:SqlPoolTransparentDataEncryption":
		r = &SqlPoolTransparentDataEncryption{}
	case "azure-native:synapse/v20210601preview:SqlPoolVulnerabilityAssessment":
		r = &SqlPoolVulnerabilityAssessment{}
	case "azure-native:synapse/v20210601preview:SqlPoolVulnerabilityAssessmentRuleBaseline":
		r = &SqlPoolVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:synapse/v20210601preview:SqlPoolWorkloadClassifier":
		r = &SqlPoolWorkloadClassifier{}
	case "azure-native:synapse/v20210601preview:SqlPoolWorkloadGroup":
		r = &SqlPoolWorkloadGroup{}
	case "azure-native:synapse/v20210601preview:Workspace":
		r = &Workspace{}
	case "azure-native:synapse/v20210601preview:WorkspaceAadAdmin":
		r = &WorkspaceAadAdmin{}
	case "azure-native:synapse/v20210601preview:WorkspaceManagedSqlServerVulnerabilityAssessment":
		r = &WorkspaceManagedSqlServerVulnerabilityAssessment{}
	case "azure-native:synapse/v20210601preview:WorkspaceSqlAadAdmin":
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
		"synapse/v20210601preview",
		&module{version},
	)
}
