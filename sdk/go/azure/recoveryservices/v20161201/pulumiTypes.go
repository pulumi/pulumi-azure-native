


package v20161201

type AzureBackupServerContainer struct {
	BackupManagementType *string                   `pulumi:"backupManagementType"`
	CanReRegister        *bool                     `pulumi:"canReRegister"`
	ContainerId          *string                   `pulumi:"containerId"`
	ContainerType        *string                   `pulumi:"containerType"`
	DpmAgentVersion      *string                   `pulumi:"dpmAgentVersion"`
	DpmServers           []string                  `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                   `pulumi:"friendlyName"`
	HealthStatus         *string                   `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                  `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                   `pulumi:"protectionStatus"`
	RegistrationStatus   *string                   `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                     `pulumi:"upgradeAvailable"`
}

type AzureBackupServerContainerResponse struct {
	BackupManagementType *string                           `pulumi:"backupManagementType"`
	CanReRegister        *bool                             `pulumi:"canReRegister"`
	ContainerId          *string                           `pulumi:"containerId"`
	ContainerType        *string                           `pulumi:"containerType"`
	DpmAgentVersion      *string                           `pulumi:"dpmAgentVersion"`
	DpmServers           []string                          `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                           `pulumi:"friendlyName"`
	HealthStatus         *string                           `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                          `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                           `pulumi:"protectionStatus"`
	RegistrationStatus   *string                           `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                             `pulumi:"upgradeAvailable"`
}

type AzureIaaSClassicComputeVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         *string `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type AzureIaaSClassicComputeVMContainerResponse struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         *string `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type AzureIaaSComputeVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         *string `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type AzureIaaSComputeVMContainerResponse struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         *string `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type AzureSQLAGWorkloadContainerProtectionContainer struct {
	BackupManagementType *string                             `pulumi:"backupManagementType"`
	ContainerType        *string                             `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                             `pulumi:"friendlyName"`
	HealthStatus         *string                             `pulumi:"healthStatus"`
	LastUpdatedTime      *string                             `pulumi:"lastUpdatedTime"`
	OperationType        *string                             `pulumi:"operationType"`
	RegistrationStatus   *string                             `pulumi:"registrationStatus"`
	SourceResourceId     *string                             `pulumi:"sourceResourceId"`
	WorkloadType         *string                             `pulumi:"workloadType"`
}

type AzureSQLAGWorkloadContainerProtectionContainerResponse struct {
	BackupManagementType *string                                     `pulumi:"backupManagementType"`
	ContainerType        *string                                     `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                                     `pulumi:"friendlyName"`
	HealthStatus         *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime      *string                                     `pulumi:"lastUpdatedTime"`
	OperationType        *string                                     `pulumi:"operationType"`
	RegistrationStatus   *string                                     `pulumi:"registrationStatus"`
	SourceResourceId     *string                                     `pulumi:"sourceResourceId"`
	WorkloadType         *string                                     `pulumi:"workloadType"`
}

type AzureSqlContainer struct {
	BackupManagementType *string `pulumi:"backupManagementType"`
	ContainerType        *string `pulumi:"containerType"`
	FriendlyName         *string `pulumi:"friendlyName"`
	HealthStatus         *string `pulumi:"healthStatus"`
	RegistrationStatus   *string `pulumi:"registrationStatus"`
}

type AzureSqlContainerResponse struct {
	BackupManagementType *string `pulumi:"backupManagementType"`
	ContainerType        *string `pulumi:"containerType"`
	FriendlyName         *string `pulumi:"friendlyName"`
	HealthStatus         *string `pulumi:"healthStatus"`
	RegistrationStatus   *string `pulumi:"registrationStatus"`
}

type AzureStorageContainer struct {
	BackupManagementType  *string  `pulumi:"backupManagementType"`
	ContainerType         *string  `pulumi:"containerType"`
	FriendlyName          *string  `pulumi:"friendlyName"`
	HealthStatus          *string  `pulumi:"healthStatus"`
	ProtectedItemCount    *float64 `pulumi:"protectedItemCount"`
	RegistrationStatus    *string  `pulumi:"registrationStatus"`
	ResourceGroup         *string  `pulumi:"resourceGroup"`
	SourceResourceId      *string  `pulumi:"sourceResourceId"`
	StorageAccountVersion *string  `pulumi:"storageAccountVersion"`
}

type AzureStorageContainerResponse struct {
	BackupManagementType  *string  `pulumi:"backupManagementType"`
	ContainerType         *string  `pulumi:"containerType"`
	FriendlyName          *string  `pulumi:"friendlyName"`
	HealthStatus          *string  `pulumi:"healthStatus"`
	ProtectedItemCount    *float64 `pulumi:"protectedItemCount"`
	RegistrationStatus    *string  `pulumi:"registrationStatus"`
	ResourceGroup         *string  `pulumi:"resourceGroup"`
	SourceResourceId      *string  `pulumi:"sourceResourceId"`
	StorageAccountVersion *string  `pulumi:"storageAccountVersion"`
}

type AzureVMAppContainerProtectionContainer struct {
	BackupManagementType *string                             `pulumi:"backupManagementType"`
	ContainerType        *string                             `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                             `pulumi:"friendlyName"`
	HealthStatus         *string                             `pulumi:"healthStatus"`
	LastUpdatedTime      *string                             `pulumi:"lastUpdatedTime"`
	OperationType        *string                             `pulumi:"operationType"`
	RegistrationStatus   *string                             `pulumi:"registrationStatus"`
	SourceResourceId     *string                             `pulumi:"sourceResourceId"`
	WorkloadType         *string                             `pulumi:"workloadType"`
}

type AzureVMAppContainerProtectionContainerResponse struct {
	BackupManagementType *string                                     `pulumi:"backupManagementType"`
	ContainerType        *string                                     `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                                     `pulumi:"friendlyName"`
	HealthStatus         *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime      *string                                     `pulumi:"lastUpdatedTime"`
	OperationType        *string                                     `pulumi:"operationType"`
	RegistrationStatus   *string                                     `pulumi:"registrationStatus"`
	SourceResourceId     *string                                     `pulumi:"sourceResourceId"`
	WorkloadType         *string                                     `pulumi:"workloadType"`
}

type AzureWorkloadContainer struct {
	BackupManagementType *string                             `pulumi:"backupManagementType"`
	ContainerType        *string                             `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                             `pulumi:"friendlyName"`
	HealthStatus         *string                             `pulumi:"healthStatus"`
	LastUpdatedTime      *string                             `pulumi:"lastUpdatedTime"`
	OperationType        *string                             `pulumi:"operationType"`
	RegistrationStatus   *string                             `pulumi:"registrationStatus"`
	SourceResourceId     *string                             `pulumi:"sourceResourceId"`
	WorkloadType         *string                             `pulumi:"workloadType"`
}

type AzureWorkloadContainerExtendedInfo struct {
	HostServerName *string                `pulumi:"hostServerName"`
	InquiryInfo    *InquiryInfo           `pulumi:"inquiryInfo"`
	NodesList      []DistributedNodesInfo `pulumi:"nodesList"`
}

type AzureWorkloadContainerExtendedInfoResponse struct {
	HostServerName *string                        `pulumi:"hostServerName"`
	InquiryInfo    *InquiryInfoResponse           `pulumi:"inquiryInfo"`
	NodesList      []DistributedNodesInfoResponse `pulumi:"nodesList"`
}

type AzureWorkloadContainerResponse struct {
	BackupManagementType *string                                     `pulumi:"backupManagementType"`
	ContainerType        *string                                     `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                                     `pulumi:"friendlyName"`
	HealthStatus         *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime      *string                                     `pulumi:"lastUpdatedTime"`
	OperationType        *string                                     `pulumi:"operationType"`
	RegistrationStatus   *string                                     `pulumi:"registrationStatus"`
	SourceResourceId     *string                                     `pulumi:"sourceResourceId"`
	WorkloadType         *string                                     `pulumi:"workloadType"`
}

type ContainerIdentityInfo struct {
	AadTenantId              *string `pulumi:"aadTenantId"`
	Audience                 *string `pulumi:"audience"`
	ServicePrincipalClientId *string `pulumi:"servicePrincipalClientId"`
	UniqueName               *string `pulumi:"uniqueName"`
}

type ContainerIdentityInfoResponse struct {
	AadTenantId              *string `pulumi:"aadTenantId"`
	Audience                 *string `pulumi:"audience"`
	ServicePrincipalClientId *string `pulumi:"servicePrincipalClientId"`
	UniqueName               *string `pulumi:"uniqueName"`
}

type DPMContainerExtendedInfo struct {
	LastRefreshedAt *string `pulumi:"lastRefreshedAt"`
}

type DPMContainerExtendedInfoResponse struct {
	LastRefreshedAt *string `pulumi:"lastRefreshedAt"`
}

type DistributedNodesInfo struct {
	NodeName *string `pulumi:"nodeName"`
	Status   *string `pulumi:"status"`
}

type DistributedNodesInfoResponse struct {
	ErrorDetail *ErrorDetailResponse `pulumi:"errorDetail"`
	NodeName    *string              `pulumi:"nodeName"`
	Status      *string              `pulumi:"status"`
}

type DpmContainer struct {
	BackupManagementType *string                   `pulumi:"backupManagementType"`
	CanReRegister        *bool                     `pulumi:"canReRegister"`
	ContainerId          *string                   `pulumi:"containerId"`
	ContainerType        *string                   `pulumi:"containerType"`
	DpmAgentVersion      *string                   `pulumi:"dpmAgentVersion"`
	DpmServers           []string                  `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                   `pulumi:"friendlyName"`
	HealthStatus         *string                   `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                  `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                   `pulumi:"protectionStatus"`
	RegistrationStatus   *string                   `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                     `pulumi:"upgradeAvailable"`
}

type DpmContainerResponse struct {
	BackupManagementType *string                           `pulumi:"backupManagementType"`
	CanReRegister        *bool                             `pulumi:"canReRegister"`
	ContainerId          *string                           `pulumi:"containerId"`
	ContainerType        *string                           `pulumi:"containerType"`
	DpmAgentVersion      *string                           `pulumi:"dpmAgentVersion"`
	DpmServers           []string                          `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                           `pulumi:"friendlyName"`
	HealthStatus         *string                           `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                          `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                           `pulumi:"protectionStatus"`
	RegistrationStatus   *string                           `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                             `pulumi:"upgradeAvailable"`
}

type ErrorDetailResponse struct {
	Code            string   `pulumi:"code"`
	Message         string   `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
}

type GenericContainer struct {
	BackupManagementType *string                       `pulumi:"backupManagementType"`
	ContainerType        *string                       `pulumi:"containerType"`
	ExtendedInformation  *GenericContainerExtendedInfo `pulumi:"extendedInformation"`
	FabricName           *string                       `pulumi:"fabricName"`
	FriendlyName         *string                       `pulumi:"friendlyName"`
	HealthStatus         *string                       `pulumi:"healthStatus"`
	RegistrationStatus   *string                       `pulumi:"registrationStatus"`
}

type GenericContainerExtendedInfo struct {
	ContainerIdentityInfo *ContainerIdentityInfo `pulumi:"containerIdentityInfo"`
	RawCertData           *string                `pulumi:"rawCertData"`
	ServiceEndpoints      map[string]string      `pulumi:"serviceEndpoints"`
}

type GenericContainerExtendedInfoResponse struct {
	ContainerIdentityInfo *ContainerIdentityInfoResponse `pulumi:"containerIdentityInfo"`
	RawCertData           *string                        `pulumi:"rawCertData"`
	ServiceEndpoints      map[string]string              `pulumi:"serviceEndpoints"`
}

type GenericContainerResponse struct {
	BackupManagementType *string                               `pulumi:"backupManagementType"`
	ContainerType        *string                               `pulumi:"containerType"`
	ExtendedInformation  *GenericContainerExtendedInfoResponse `pulumi:"extendedInformation"`
	FabricName           *string                               `pulumi:"fabricName"`
	FriendlyName         *string                               `pulumi:"friendlyName"`
	HealthStatus         *string                               `pulumi:"healthStatus"`
	RegistrationStatus   *string                               `pulumi:"registrationStatus"`
}

type IaaSVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         *string `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type IaaSVMContainerResponse struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         *string `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type InquiryInfo struct {
	InquiryDetails []WorkloadInquiryDetails `pulumi:"inquiryDetails"`
	Status         *string                  `pulumi:"status"`
}

type InquiryInfoResponse struct {
	ErrorDetail    *ErrorDetailResponse             `pulumi:"errorDetail"`
	InquiryDetails []WorkloadInquiryDetailsResponse `pulumi:"inquiryDetails"`
	Status         *string                          `pulumi:"status"`
}

type InquiryValidation struct {
	Status *string `pulumi:"status"`
}

type InquiryValidationResponse struct {
	AdditionalDetail string               `pulumi:"additionalDetail"`
	ErrorDetail      *ErrorDetailResponse `pulumi:"errorDetail"`
	Status           *string              `pulumi:"status"`
}

type MABContainerHealthDetails struct {
	Code            *int     `pulumi:"code"`
	Message         *string  `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
	Title           *string  `pulumi:"title"`
}

type MABContainerHealthDetailsResponse struct {
	Code            *int     `pulumi:"code"`
	Message         *string  `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
	Title           *string  `pulumi:"title"`
}

type MabContainer struct {
	AgentVersion              *string                     `pulumi:"agentVersion"`
	BackupManagementType      *string                     `pulumi:"backupManagementType"`
	CanReRegister             *bool                       `pulumi:"canReRegister"`
	ContainerHealthState      *string                     `pulumi:"containerHealthState"`
	ContainerId               *float64                    `pulumi:"containerId"`
	ContainerType             *string                     `pulumi:"containerType"`
	ExtendedInfo              *MabContainerExtendedInfo   `pulumi:"extendedInfo"`
	FriendlyName              *string                     `pulumi:"friendlyName"`
	HealthStatus              *string                     `pulumi:"healthStatus"`
	MabContainerHealthDetails []MABContainerHealthDetails `pulumi:"mabContainerHealthDetails"`
	ProtectedItemCount        *float64                    `pulumi:"protectedItemCount"`
	RegistrationStatus        *string                     `pulumi:"registrationStatus"`
}

type MabContainerExtendedInfo struct {
	BackupItemType   *string  `pulumi:"backupItemType"`
	BackupItems      []string `pulumi:"backupItems"`
	LastBackupStatus *string  `pulumi:"lastBackupStatus"`
	LastRefreshedAt  *string  `pulumi:"lastRefreshedAt"`
	PolicyName       *string  `pulumi:"policyName"`
}

type MabContainerExtendedInfoResponse struct {
	BackupItemType   *string  `pulumi:"backupItemType"`
	BackupItems      []string `pulumi:"backupItems"`
	LastBackupStatus *string  `pulumi:"lastBackupStatus"`
	LastRefreshedAt  *string  `pulumi:"lastRefreshedAt"`
	PolicyName       *string  `pulumi:"policyName"`
}

type MabContainerResponse struct {
	AgentVersion              *string                             `pulumi:"agentVersion"`
	BackupManagementType      *string                             `pulumi:"backupManagementType"`
	CanReRegister             *bool                               `pulumi:"canReRegister"`
	ContainerHealthState      *string                             `pulumi:"containerHealthState"`
	ContainerId               *float64                            `pulumi:"containerId"`
	ContainerType             *string                             `pulumi:"containerType"`
	ExtendedInfo              *MabContainerExtendedInfoResponse   `pulumi:"extendedInfo"`
	FriendlyName              *string                             `pulumi:"friendlyName"`
	HealthStatus              *string                             `pulumi:"healthStatus"`
	MabContainerHealthDetails []MABContainerHealthDetailsResponse `pulumi:"mabContainerHealthDetails"`
	ProtectedItemCount        *float64                            `pulumi:"protectedItemCount"`
	RegistrationStatus        *string                             `pulumi:"registrationStatus"`
}

type WorkloadInquiryDetails struct {
	InquiryValidation *InquiryValidation `pulumi:"inquiryValidation"`
	ItemCount         *float64           `pulumi:"itemCount"`
	Type              *string            `pulumi:"type"`
}

type WorkloadInquiryDetailsResponse struct {
	InquiryValidation *InquiryValidationResponse `pulumi:"inquiryValidation"`
	ItemCount         *float64                   `pulumi:"itemCount"`
	Type              *string                    `pulumi:"type"`
}

func init() {
}
