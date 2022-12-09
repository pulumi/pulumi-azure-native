


package v20220501preview

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
	case "azure-native:sql/v20220501preview:BackupShortTermRetentionPolicy":
		r = &BackupShortTermRetentionPolicy{}
	case "azure-native:sql/v20220501preview:DataMaskingPolicy":
		r = &DataMaskingPolicy{}
	case "azure-native:sql/v20220501preview:Database":
		r = &Database{}
	case "azure-native:sql/v20220501preview:DatabaseAdvisor":
		r = &DatabaseAdvisor{}
	case "azure-native:sql/v20220501preview:DatabaseBlobAuditingPolicy":
		r = &DatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20220501preview:DatabaseSecurityAlertPolicy":
		r = &DatabaseSecurityAlertPolicy{}
	case "azure-native:sql/v20220501preview:DatabaseSqlVulnerabilityAssessmentRuleBaseline":
		r = &DatabaseSqlVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220501preview:DatabaseVulnerabilityAssessment":
		r = &DatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20220501preview:DatabaseVulnerabilityAssessmentRuleBaseline":
		r = &DatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220501preview:DistributedAvailabilityGroup":
		r = &DistributedAvailabilityGroup{}
	case "azure-native:sql/v20220501preview:ElasticPool":
		r = &ElasticPool{}
	case "azure-native:sql/v20220501preview:EncryptionProtector":
		r = &EncryptionProtector{}
	case "azure-native:sql/v20220501preview:ExtendedDatabaseBlobAuditingPolicy":
		r = &ExtendedDatabaseBlobAuditingPolicy{}
	case "azure-native:sql/v20220501preview:ExtendedServerBlobAuditingPolicy":
		r = &ExtendedServerBlobAuditingPolicy{}
	case "azure-native:sql/v20220501preview:FailoverGroup":
		r = &FailoverGroup{}
	case "azure-native:sql/v20220501preview:FirewallRule":
		r = &FirewallRule{}
	case "azure-native:sql/v20220501preview:GeoBackupPolicy":
		r = &GeoBackupPolicy{}
	case "azure-native:sql/v20220501preview:IPv6FirewallRule":
		r = &IPv6FirewallRule{}
	case "azure-native:sql/v20220501preview:InstanceFailoverGroup":
		r = &InstanceFailoverGroup{}
	case "azure-native:sql/v20220501preview:InstancePool":
		r = &InstancePool{}
	case "azure-native:sql/v20220501preview:Job":
		r = &Job{}
	case "azure-native:sql/v20220501preview:JobAgent":
		r = &JobAgent{}
	case "azure-native:sql/v20220501preview:JobCredential":
		r = &JobCredential{}
	case "azure-native:sql/v20220501preview:JobStep":
		r = &JobStep{}
	case "azure-native:sql/v20220501preview:JobTargetGroup":
		r = &JobTargetGroup{}
	case "azure-native:sql/v20220501preview:LongTermRetentionPolicy":
		r = &LongTermRetentionPolicy{}
	case "azure-native:sql/v20220501preview:ManagedDatabase":
		r = &ManagedDatabase{}
	case "azure-native:sql/v20220501preview:ManagedDatabaseSensitivityLabel":
		r = &ManagedDatabaseSensitivityLabel{}
	case "azure-native:sql/v20220501preview:ManagedDatabaseVulnerabilityAssessment":
		r = &ManagedDatabaseVulnerabilityAssessment{}
	case "azure-native:sql/v20220501preview:ManagedDatabaseVulnerabilityAssessmentRuleBaseline":
		r = &ManagedDatabaseVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220501preview:ManagedInstance":
		r = &ManagedInstance{}
	case "azure-native:sql/v20220501preview:ManagedInstanceAdministrator":
		r = &ManagedInstanceAdministrator{}
	case "azure-native:sql/v20220501preview:ManagedInstanceAzureADOnlyAuthentication":
		r = &ManagedInstanceAzureADOnlyAuthentication{}
	case "azure-native:sql/v20220501preview:ManagedInstanceKey":
		r = &ManagedInstanceKey{}
	case "azure-native:sql/v20220501preview:ManagedInstanceLongTermRetentionPolicy":
		r = &ManagedInstanceLongTermRetentionPolicy{}
	case "azure-native:sql/v20220501preview:ManagedInstancePrivateEndpointConnection":
		r = &ManagedInstancePrivateEndpointConnection{}
	case "azure-native:sql/v20220501preview:ManagedInstanceVulnerabilityAssessment":
		r = &ManagedInstanceVulnerabilityAssessment{}
	case "azure-native:sql/v20220501preview:ManagedServerDnsAlias":
		r = &ManagedServerDnsAlias{}
	case "azure-native:sql/v20220501preview:OutboundFirewallRule":
		r = &OutboundFirewallRule{}
	case "azure-native:sql/v20220501preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:sql/v20220501preview:SensitivityLabel":
		r = &SensitivityLabel{}
	case "azure-native:sql/v20220501preview:Server":
		r = &Server{}
	case "azure-native:sql/v20220501preview:ServerAdvisor":
		r = &ServerAdvisor{}
	case "azure-native:sql/v20220501preview:ServerAzureADAdministrator":
		r = &ServerAzureADAdministrator{}
	case "azure-native:sql/v20220501preview:ServerAzureADOnlyAuthentication":
		r = &ServerAzureADOnlyAuthentication{}
	case "azure-native:sql/v20220501preview:ServerBlobAuditingPolicy":
		r = &ServerBlobAuditingPolicy{}
	case "azure-native:sql/v20220501preview:ServerDnsAlias":
		r = &ServerDnsAlias{}
	case "azure-native:sql/v20220501preview:ServerKey":
		r = &ServerKey{}
	case "azure-native:sql/v20220501preview:ServerSecurityAlertPolicy":
		r = &ServerSecurityAlertPolicy{}
	case "azure-native:sql/v20220501preview:ServerTrustCertificate":
		r = &ServerTrustCertificate{}
	case "azure-native:sql/v20220501preview:ServerTrustGroup":
		r = &ServerTrustGroup{}
	case "azure-native:sql/v20220501preview:ServerVulnerabilityAssessment":
		r = &ServerVulnerabilityAssessment{}
	case "azure-native:sql/v20220501preview:SqlVulnerabilityAssessmentRuleBaseline":
		r = &SqlVulnerabilityAssessmentRuleBaseline{}
	case "azure-native:sql/v20220501preview:SqlVulnerabilityAssessmentsSetting":
		r = &SqlVulnerabilityAssessmentsSetting{}
	case "azure-native:sql/v20220501preview:SyncAgent":
		r = &SyncAgent{}
	case "azure-native:sql/v20220501preview:SyncGroup":
		r = &SyncGroup{}
	case "azure-native:sql/v20220501preview:SyncMember":
		r = &SyncMember{}
	case "azure-native:sql/v20220501preview:TransparentDataEncryption":
		r = &TransparentDataEncryption{}
	case "azure-native:sql/v20220501preview:VirtualNetworkRule":
		r = &VirtualNetworkRule{}
	case "azure-native:sql/v20220501preview:WorkloadClassifier":
		r = &WorkloadClassifier{}
	case "azure-native:sql/v20220501preview:WorkloadGroup":
		r = &WorkloadGroup{}
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
		"sql/v20220501preview",
		&module{version},
	)
}
