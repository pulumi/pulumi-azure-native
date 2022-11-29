


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AwsEnvironmentData struct {
	EnvironmentType    string      `pulumi:"environmentType"`
	OrganizationalData interface{} `pulumi:"organizationalData"`
}

type AwsEnvironmentDataResponse struct {
	EnvironmentType    string      `pulumi:"environmentType"`
	OrganizationalData interface{} `pulumi:"organizationalData"`
}

type AwsOrganizationalDataMaster struct {
	ExcludedAccountIds         []string `pulumi:"excludedAccountIds"`
	OrganizationMembershipType string   `pulumi:"organizationMembershipType"`
	StacksetName               *string  `pulumi:"stacksetName"`
}

type AwsOrganizationalDataMasterResponse struct {
	ExcludedAccountIds         []string `pulumi:"excludedAccountIds"`
	OrganizationMembershipType string   `pulumi:"organizationMembershipType"`
	StacksetName               *string  `pulumi:"stacksetName"`
}

type AwsOrganizationalDataMember struct {
	OrganizationMembershipType string  `pulumi:"organizationMembershipType"`
	ParentHierarchyId          *string `pulumi:"parentHierarchyId"`
}

type AwsOrganizationalDataMemberResponse struct {
	OrganizationMembershipType string  `pulumi:"organizationMembershipType"`
	ParentHierarchyId          *string `pulumi:"parentHierarchyId"`
}

type AzureDevOpsScopeEnvironmentData struct {
	EnvironmentType string `pulumi:"environmentType"`
}

type AzureDevOpsScopeEnvironmentDataResponse struct {
	EnvironmentType string `pulumi:"environmentType"`
}

type CspmMonitorAwsOffering struct {
	NativeCloudConnection *CspmMonitorAwsOfferingNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                       `pulumi:"offeringType"`
}

type CspmMonitorAwsOfferingNativeCloudConnection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type CspmMonitorAwsOfferingResponse struct {
	Description           string                                               `pulumi:"description"`
	NativeCloudConnection *CspmMonitorAwsOfferingResponseNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                               `pulumi:"offeringType"`
}

type CspmMonitorAwsOfferingResponseNativeCloudConnection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type CspmMonitorAzureDevOpsOffering struct {
	OfferingType string `pulumi:"offeringType"`
}

type CspmMonitorAzureDevOpsOfferingResponse struct {
	Description  string `pulumi:"description"`
	OfferingType string `pulumi:"offeringType"`
}

type CspmMonitorGcpOffering struct {
	NativeCloudConnection *CspmMonitorGcpOfferingNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                       `pulumi:"offeringType"`
}

type CspmMonitorGcpOfferingNativeCloudConnection struct {
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string `pulumi:"workloadIdentityProviderId"`
}

type CspmMonitorGcpOfferingResponse struct {
	Description           string                                               `pulumi:"description"`
	NativeCloudConnection *CspmMonitorGcpOfferingResponseNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                               `pulumi:"offeringType"`
}

type CspmMonitorGcpOfferingResponseNativeCloudConnection struct {
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string `pulumi:"workloadIdentityProviderId"`
}

type CspmMonitorGithubOffering struct {
	OfferingType string `pulumi:"offeringType"`
}

type CspmMonitorGithubOfferingResponse struct {
	Description  string `pulumi:"description"`
	OfferingType string `pulumi:"offeringType"`
}

type DefenderCspmAwsOffering struct {
	OfferingType string                             `pulumi:"offeringType"`
	VmScanners   *DefenderCspmAwsOfferingVmScanners `pulumi:"vmScanners"`
}

type DefenderCspmAwsOfferingConfiguration struct {
	CloudRoleArn  *string           `pulumi:"cloudRoleArn"`
	ExclusionTags map[string]string `pulumi:"exclusionTags"`
	ScanningMode  *string           `pulumi:"scanningMode"`
}

type DefenderCspmAwsOfferingResponse struct {
	Description  string                                     `pulumi:"description"`
	OfferingType string                                     `pulumi:"offeringType"`
	VmScanners   *DefenderCspmAwsOfferingResponseVmScanners `pulumi:"vmScanners"`
}

type DefenderCspmAwsOfferingResponseConfiguration struct {
	CloudRoleArn  *string           `pulumi:"cloudRoleArn"`
	ExclusionTags map[string]string `pulumi:"exclusionTags"`
	ScanningMode  *string           `pulumi:"scanningMode"`
}

type DefenderCspmAwsOfferingResponseVmScanners struct {
	Configuration *DefenderCspmAwsOfferingResponseConfiguration `pulumi:"configuration"`
	Enabled       *bool                                         `pulumi:"enabled"`
}

type DefenderCspmAwsOfferingVmScanners struct {
	Configuration *DefenderCspmAwsOfferingConfiguration `pulumi:"configuration"`
	Enabled       *bool                                 `pulumi:"enabled"`
}

type DefenderCspmGcpOffering struct {
	OfferingType string `pulumi:"offeringType"`
}

type DefenderCspmGcpOfferingResponse struct {
	Description  string `pulumi:"description"`
	OfferingType string `pulumi:"offeringType"`
}

type DefenderFoDatabasesAwsOffering struct {
	ArcAutoProvisioning *DefenderFoDatabasesAwsOfferingArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	OfferingType        string                                             `pulumi:"offeringType"`
	Rds                 *DefenderFoDatabasesAwsOfferingRds                 `pulumi:"rds"`
}

type DefenderFoDatabasesAwsOfferingArcAutoProvisioning struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
	Enabled      *bool   `pulumi:"enabled"`
}

type DefenderFoDatabasesAwsOfferingRds struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
	Enabled      *bool   `pulumi:"enabled"`
}

type DefenderFoDatabasesAwsOfferingResponse struct {
	ArcAutoProvisioning *DefenderFoDatabasesAwsOfferingResponseArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	Description         string                                                     `pulumi:"description"`
	OfferingType        string                                                     `pulumi:"offeringType"`
	Rds                 *DefenderFoDatabasesAwsOfferingResponseRds                 `pulumi:"rds"`
}

type DefenderFoDatabasesAwsOfferingResponseArcAutoProvisioning struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
	Enabled      *bool   `pulumi:"enabled"`
}

type DefenderFoDatabasesAwsOfferingResponseRds struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
	Enabled      *bool   `pulumi:"enabled"`
}

type DefenderForContainersAwsOffering struct {
	AutoProvisioning                       *bool                                                                 `pulumi:"autoProvisioning"`
	CloudWatchToKinesis                    *DefenderForContainersAwsOfferingCloudWatchToKinesis                  `pulumi:"cloudWatchToKinesis"`
	ContainerVulnerabilityAssessment       *DefenderForContainersAwsOfferingContainerVulnerabilityAssessment     `pulumi:"containerVulnerabilityAssessment"`
	ContainerVulnerabilityAssessmentTask   *DefenderForContainersAwsOfferingContainerVulnerabilityAssessmentTask `pulumi:"containerVulnerabilityAssessmentTask"`
	EnableContainerVulnerabilityAssessment *bool                                                                 `pulumi:"enableContainerVulnerabilityAssessment"`
	KinesisToS3                            *DefenderForContainersAwsOfferingKinesisToS3                          `pulumi:"kinesisToS3"`
	KubeAuditRetentionTime                 *float64                                                              `pulumi:"kubeAuditRetentionTime"`
	KubernetesScubaReader                  *DefenderForContainersAwsOfferingKubernetesScubaReader                `pulumi:"kubernetesScubaReader"`
	KubernetesService                      *DefenderForContainersAwsOfferingKubernetesService                    `pulumi:"kubernetesService"`
	OfferingType                           string                                                                `pulumi:"offeringType"`
	ScubaExternalId                        *string                                                               `pulumi:"scubaExternalId"`
}

type DefenderForContainersAwsOfferingCloudWatchToKinesis struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingContainerVulnerabilityAssessment struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingContainerVulnerabilityAssessmentTask struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingKinesisToS3 struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingKubernetesScubaReader struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingKubernetesService struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponse struct {
	AutoProvisioning                       *bool                                                                         `pulumi:"autoProvisioning"`
	CloudWatchToKinesis                    *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis                  `pulumi:"cloudWatchToKinesis"`
	ContainerVulnerabilityAssessment       *DefenderForContainersAwsOfferingResponseContainerVulnerabilityAssessment     `pulumi:"containerVulnerabilityAssessment"`
	ContainerVulnerabilityAssessmentTask   *DefenderForContainersAwsOfferingResponseContainerVulnerabilityAssessmentTask `pulumi:"containerVulnerabilityAssessmentTask"`
	Description                            string                                                                        `pulumi:"description"`
	EnableContainerVulnerabilityAssessment *bool                                                                         `pulumi:"enableContainerVulnerabilityAssessment"`
	KinesisToS3                            *DefenderForContainersAwsOfferingResponseKinesisToS3                          `pulumi:"kinesisToS3"`
	KubeAuditRetentionTime                 *float64                                                                      `pulumi:"kubeAuditRetentionTime"`
	KubernetesScubaReader                  *DefenderForContainersAwsOfferingResponseKubernetesScubaReader                `pulumi:"kubernetesScubaReader"`
	KubernetesService                      *DefenderForContainersAwsOfferingResponseKubernetesService                    `pulumi:"kubernetesService"`
	OfferingType                           string                                                                        `pulumi:"offeringType"`
	ScubaExternalId                        *string                                                                       `pulumi:"scubaExternalId"`
}

type DefenderForContainersAwsOfferingResponseCloudWatchToKinesis struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponseContainerVulnerabilityAssessment struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponseContainerVulnerabilityAssessmentTask struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponseKinesisToS3 struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponseKubernetesScubaReader struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponseKubernetesService struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersGcpOffering struct {
	AuditLogsAutoProvisioningFlag     *bool                                                              `pulumi:"auditLogsAutoProvisioningFlag"`
	DataPipelineNativeCloudConnection *DefenderForContainersGcpOfferingDataPipelineNativeCloudConnection `pulumi:"dataPipelineNativeCloudConnection"`
	DefenderAgentAutoProvisioningFlag *bool                                                              `pulumi:"defenderAgentAutoProvisioningFlag"`
	NativeCloudConnection             *DefenderForContainersGcpOfferingNativeCloudConnection             `pulumi:"nativeCloudConnection"`
	OfferingType                      string                                                             `pulumi:"offeringType"`
	PolicyAgentAutoProvisioningFlag   *bool                                                              `pulumi:"policyAgentAutoProvisioningFlag"`
}

type DefenderForContainersGcpOfferingDataPipelineNativeCloudConnection struct {
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string `pulumi:"workloadIdentityProviderId"`
}

type DefenderForContainersGcpOfferingNativeCloudConnection struct {
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string `pulumi:"workloadIdentityProviderId"`
}

type DefenderForContainersGcpOfferingResponse struct {
	AuditLogsAutoProvisioningFlag     *bool                                                                      `pulumi:"auditLogsAutoProvisioningFlag"`
	DataPipelineNativeCloudConnection *DefenderForContainersGcpOfferingResponseDataPipelineNativeCloudConnection `pulumi:"dataPipelineNativeCloudConnection"`
	DefenderAgentAutoProvisioningFlag *bool                                                                      `pulumi:"defenderAgentAutoProvisioningFlag"`
	Description                       string                                                                     `pulumi:"description"`
	NativeCloudConnection             *DefenderForContainersGcpOfferingResponseNativeCloudConnection             `pulumi:"nativeCloudConnection"`
	OfferingType                      string                                                                     `pulumi:"offeringType"`
	PolicyAgentAutoProvisioningFlag   *bool                                                                      `pulumi:"policyAgentAutoProvisioningFlag"`
}

type DefenderForContainersGcpOfferingResponseDataPipelineNativeCloudConnection struct {
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string `pulumi:"workloadIdentityProviderId"`
}

type DefenderForContainersGcpOfferingResponseNativeCloudConnection struct {
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string `pulumi:"workloadIdentityProviderId"`
}

type DefenderForDatabasesGcpOffering struct {
	ArcAutoProvisioning                     *DefenderForDatabasesGcpOfferingArcAutoProvisioning                     `pulumi:"arcAutoProvisioning"`
	DefenderForDatabasesArcAutoProvisioning *DefenderForDatabasesGcpOfferingDefenderForDatabasesArcAutoProvisioning `pulumi:"defenderForDatabasesArcAutoProvisioning"`
	OfferingType                            string                                                                  `pulumi:"offeringType"`
}

type DefenderForDatabasesGcpOfferingArcAutoProvisioning struct {
	Enabled *bool `pulumi:"enabled"`
}

type DefenderForDatabasesGcpOfferingDefenderForDatabasesArcAutoProvisioning struct {
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string `pulumi:"workloadIdentityProviderId"`
}

type DefenderForDatabasesGcpOfferingResponse struct {
	ArcAutoProvisioning                     *DefenderForDatabasesGcpOfferingResponseArcAutoProvisioning                     `pulumi:"arcAutoProvisioning"`
	DefenderForDatabasesArcAutoProvisioning *DefenderForDatabasesGcpOfferingResponseDefenderForDatabasesArcAutoProvisioning `pulumi:"defenderForDatabasesArcAutoProvisioning"`
	Description                             string                                                                          `pulumi:"description"`
	OfferingType                            string                                                                          `pulumi:"offeringType"`
}

type DefenderForDatabasesGcpOfferingResponseArcAutoProvisioning struct {
	Enabled *bool `pulumi:"enabled"`
}

type DefenderForDatabasesGcpOfferingResponseDefenderForDatabasesArcAutoProvisioning struct {
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string `pulumi:"workloadIdentityProviderId"`
}

type DefenderForDevOpsAzureDevOpsOffering struct {
	OfferingType string `pulumi:"offeringType"`
}

type DefenderForDevOpsAzureDevOpsOfferingResponse struct {
	Description  string `pulumi:"description"`
	OfferingType string `pulumi:"offeringType"`
}

type DefenderForDevOpsGithubOffering struct {
	OfferingType string `pulumi:"offeringType"`
}

type DefenderForDevOpsGithubOfferingResponse struct {
	Description  string `pulumi:"description"`
	OfferingType string `pulumi:"offeringType"`
}

type DefenderForServersAwsOffering struct {
	ArcAutoProvisioning *DefenderForServersAwsOfferingArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersAwsOfferingDefenderForServers  `pulumi:"defenderForServers"`
	MdeAutoProvisioning *DefenderForServersAwsOfferingMdeAutoProvisioning `pulumi:"mdeAutoProvisioning"`
	OfferingType        string                                            `pulumi:"offeringType"`
	SubPlan             *DefenderForServersAwsOfferingSubPlan             `pulumi:"subPlan"`
	VaAutoProvisioning  *DefenderForServersAwsOfferingVaAutoProvisioning  `pulumi:"vaAutoProvisioning"`
	VmScanners          *DefenderForServersAwsOfferingVmScanners          `pulumi:"vmScanners"`
}

type DefenderForServersAwsOfferingArcAutoProvisioning struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
	Enabled      *bool   `pulumi:"enabled"`
}

type DefenderForServersAwsOfferingConfiguration struct {
	Type *string `pulumi:"type"`
}

type DefenderForServersAwsOfferingConfigurationConfiguration struct {
	CloudRoleArn  *string           `pulumi:"cloudRoleArn"`
	ExclusionTags map[string]string `pulumi:"exclusionTags"`
	ScanningMode  *string           `pulumi:"scanningMode"`
}

type DefenderForServersAwsOfferingDefenderForServers struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForServersAwsOfferingMdeAutoProvisioning struct {
	Configuration interface{} `pulumi:"configuration"`
	Enabled       *bool       `pulumi:"enabled"`
}

type DefenderForServersAwsOfferingResponse struct {
	ArcAutoProvisioning *DefenderForServersAwsOfferingResponseArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersAwsOfferingResponseDefenderForServers  `pulumi:"defenderForServers"`
	Description         string                                                    `pulumi:"description"`
	MdeAutoProvisioning *DefenderForServersAwsOfferingResponseMdeAutoProvisioning `pulumi:"mdeAutoProvisioning"`
	OfferingType        string                                                    `pulumi:"offeringType"`
	SubPlan             *DefenderForServersAwsOfferingResponseSubPlan             `pulumi:"subPlan"`
	VaAutoProvisioning  *DefenderForServersAwsOfferingResponseVaAutoProvisioning  `pulumi:"vaAutoProvisioning"`
	VmScanners          *DefenderForServersAwsOfferingResponseVmScanners          `pulumi:"vmScanners"`
}

type DefenderForServersAwsOfferingResponseArcAutoProvisioning struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
	Enabled      *bool   `pulumi:"enabled"`
}

type DefenderForServersAwsOfferingResponseConfiguration struct {
	Type *string `pulumi:"type"`
}

type DefenderForServersAwsOfferingResponseConfigurationConfiguration struct {
	CloudRoleArn  *string           `pulumi:"cloudRoleArn"`
	ExclusionTags map[string]string `pulumi:"exclusionTags"`
	ScanningMode  *string           `pulumi:"scanningMode"`
}

type DefenderForServersAwsOfferingResponseDefenderForServers struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForServersAwsOfferingResponseMdeAutoProvisioning struct {
	Configuration interface{} `pulumi:"configuration"`
	Enabled       *bool       `pulumi:"enabled"`
}

type DefenderForServersAwsOfferingResponseSubPlan struct {
	Type *string `pulumi:"type"`
}

type DefenderForServersAwsOfferingResponseVaAutoProvisioning struct {
	Configuration *DefenderForServersAwsOfferingResponseConfiguration `pulumi:"configuration"`
	Enabled       *bool                                               `pulumi:"enabled"`
}

type DefenderForServersAwsOfferingResponseVmScanners struct {
	Configuration *DefenderForServersAwsOfferingResponseConfigurationConfiguration `pulumi:"configuration"`
	Enabled       *bool                                                            `pulumi:"enabled"`
}

type DefenderForServersAwsOfferingSubPlan struct {
	Type *string `pulumi:"type"`
}

type DefenderForServersAwsOfferingVaAutoProvisioning struct {
	Configuration *DefenderForServersAwsOfferingConfiguration `pulumi:"configuration"`
	Enabled       *bool                                       `pulumi:"enabled"`
}

type DefenderForServersAwsOfferingVmScanners struct {
	Configuration *DefenderForServersAwsOfferingConfigurationConfiguration `pulumi:"configuration"`
	Enabled       *bool                                                    `pulumi:"enabled"`
}

type DefenderForServersGcpOffering struct {
	ArcAutoProvisioning *DefenderForServersGcpOfferingArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersGcpOfferingDefenderForServers  `pulumi:"defenderForServers"`
	MdeAutoProvisioning *DefenderForServersGcpOfferingMdeAutoProvisioning `pulumi:"mdeAutoProvisioning"`
	OfferingType        string                                            `pulumi:"offeringType"`
	SubPlan             *DefenderForServersGcpOfferingSubPlan             `pulumi:"subPlan"`
	VaAutoProvisioning  *DefenderForServersGcpOfferingVaAutoProvisioning  `pulumi:"vaAutoProvisioning"`
}

type DefenderForServersGcpOfferingArcAutoProvisioning struct {
	Enabled *bool `pulumi:"enabled"`
}

type DefenderForServersGcpOfferingConfiguration struct {
	Type *string `pulumi:"type"`
}

type DefenderForServersGcpOfferingDefenderForServers struct {
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string `pulumi:"workloadIdentityProviderId"`
}

type DefenderForServersGcpOfferingMdeAutoProvisioning struct {
	Configuration interface{} `pulumi:"configuration"`
	Enabled       *bool       `pulumi:"enabled"`
}

type DefenderForServersGcpOfferingResponse struct {
	ArcAutoProvisioning *DefenderForServersGcpOfferingResponseArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersGcpOfferingResponseDefenderForServers  `pulumi:"defenderForServers"`
	Description         string                                                    `pulumi:"description"`
	MdeAutoProvisioning *DefenderForServersGcpOfferingResponseMdeAutoProvisioning `pulumi:"mdeAutoProvisioning"`
	OfferingType        string                                                    `pulumi:"offeringType"`
	SubPlan             *DefenderForServersGcpOfferingResponseSubPlan             `pulumi:"subPlan"`
	VaAutoProvisioning  *DefenderForServersGcpOfferingResponseVaAutoProvisioning  `pulumi:"vaAutoProvisioning"`
}

type DefenderForServersGcpOfferingResponseArcAutoProvisioning struct {
	Enabled *bool `pulumi:"enabled"`
}

type DefenderForServersGcpOfferingResponseConfiguration struct {
	Type *string `pulumi:"type"`
}

type DefenderForServersGcpOfferingResponseDefenderForServers struct {
	ServiceAccountEmailAddress *string `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string `pulumi:"workloadIdentityProviderId"`
}

type DefenderForServersGcpOfferingResponseMdeAutoProvisioning struct {
	Configuration interface{} `pulumi:"configuration"`
	Enabled       *bool       `pulumi:"enabled"`
}

type DefenderForServersGcpOfferingResponseSubPlan struct {
	Type *string `pulumi:"type"`
}

type DefenderForServersGcpOfferingResponseVaAutoProvisioning struct {
	Configuration *DefenderForServersGcpOfferingResponseConfiguration `pulumi:"configuration"`
	Enabled       *bool                                               `pulumi:"enabled"`
}

type DefenderForServersGcpOfferingSubPlan struct {
	Type *string `pulumi:"type"`
}

type DefenderForServersGcpOfferingVaAutoProvisioning struct {
	Configuration *DefenderForServersGcpOfferingConfiguration `pulumi:"configuration"`
	Enabled       *bool                                       `pulumi:"enabled"`
}

type GcpOrganizationalDataMember struct {
	ManagementProjectNumber    *string `pulumi:"managementProjectNumber"`
	OrganizationMembershipType string  `pulumi:"organizationMembershipType"`
	ParentHierarchyId          *string `pulumi:"parentHierarchyId"`
}

type GcpOrganizationalDataMemberResponse struct {
	ManagementProjectNumber    *string `pulumi:"managementProjectNumber"`
	OrganizationMembershipType string  `pulumi:"organizationMembershipType"`
	ParentHierarchyId          *string `pulumi:"parentHierarchyId"`
}

type GcpOrganizationalDataOrganization struct {
	ExcludedProjectNumbers     []string `pulumi:"excludedProjectNumbers"`
	OrganizationMembershipType string   `pulumi:"organizationMembershipType"`
	ServiceAccountEmailAddress *string  `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string  `pulumi:"workloadIdentityProviderId"`
}

type GcpOrganizationalDataOrganizationResponse struct {
	ExcludedProjectNumbers     []string `pulumi:"excludedProjectNumbers"`
	OrganizationMembershipType string   `pulumi:"organizationMembershipType"`
	ServiceAccountEmailAddress *string  `pulumi:"serviceAccountEmailAddress"`
	WorkloadIdentityProviderId *string  `pulumi:"workloadIdentityProviderId"`
}

type GcpProjectDetails struct {
	ProjectId     *string `pulumi:"projectId"`
	ProjectNumber *string `pulumi:"projectNumber"`
}

type GcpProjectDetailsResponse struct {
	ProjectId              *string `pulumi:"projectId"`
	ProjectNumber          *string `pulumi:"projectNumber"`
	WorkloadIdentityPoolId string  `pulumi:"workloadIdentityPoolId"`
}

type GcpProjectEnvironmentData struct {
	EnvironmentType    string             `pulumi:"environmentType"`
	OrganizationalData interface{}        `pulumi:"organizationalData"`
	ProjectDetails     *GcpProjectDetails `pulumi:"projectDetails"`
}

type GcpProjectEnvironmentDataResponse struct {
	EnvironmentType    string                     `pulumi:"environmentType"`
	OrganizationalData interface{}                `pulumi:"organizationalData"`
	ProjectDetails     *GcpProjectDetailsResponse `pulumi:"projectDetails"`
}

type GithubScopeEnvironmentData struct {
	EnvironmentType string `pulumi:"environmentType"`
}

type GithubScopeEnvironmentDataResponse struct {
	EnvironmentType string `pulumi:"environmentType"`
}

type InformationProtectionAwsOffering struct {
	InformationProtection *InformationProtectionAwsOfferingInformationProtection `pulumi:"informationProtection"`
	OfferingType          string                                                 `pulumi:"offeringType"`
}

type InformationProtectionAwsOfferingInformationProtection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type InformationProtectionAwsOfferingResponse struct {
	Description           string                                                         `pulumi:"description"`
	InformationProtection *InformationProtectionAwsOfferingResponseInformationProtection `pulumi:"informationProtection"`
	OfferingType          string                                                         `pulumi:"offeringType"`
}

type InformationProtectionAwsOfferingResponseInformationProtection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
