


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AKS struct {
	ComputeLocation  *string        `pulumi:"computeLocation"`
	ComputeType      string         `pulumi:"computeType"`
	Description      *string        `pulumi:"description"`
	DisableLocalAuth *bool          `pulumi:"disableLocalAuth"`
	Properties       *AKSProperties `pulumi:"properties"`
	ResourceId       *string        `pulumi:"resourceId"`
}


func (val *AKS) Defaults() *AKS {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type AKSProperties struct {
	AgentCount                 *int                        `pulumi:"agentCount"`
	AgentVmSize                *string                     `pulumi:"agentVmSize"`
	AksNetworkingConfiguration *AksNetworkingConfiguration `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                *string                     `pulumi:"clusterFqdn"`
	ClusterPurpose             *string                     `pulumi:"clusterPurpose"`
	LoadBalancerSubnet         *string                     `pulumi:"loadBalancerSubnet"`
	LoadBalancerType           *string                     `pulumi:"loadBalancerType"`
	SslConfiguration           *SslConfiguration           `pulumi:"sslConfiguration"`
}


func (val *AKSProperties) Defaults() *AKSProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClusterPurpose) {
		clusterPurpose_ := "FastProd"
		tmp.ClusterPurpose = &clusterPurpose_
	}
	if isZero(tmp.LoadBalancerType) {
		loadBalancerType_ := "PublicIp"
		tmp.LoadBalancerType = &loadBalancerType_
	}
	return &tmp
}

type AKSResponse struct {
	ComputeLocation    *string                 `pulumi:"computeLocation"`
	ComputeType        string                  `pulumi:"computeType"`
	CreatedOn          string                  `pulumi:"createdOn"`
	Description        *string                 `pulumi:"description"`
	DisableLocalAuth   *bool                   `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                    `pulumi:"isAttachedCompute"`
	ModifiedOn         string                  `pulumi:"modifiedOn"`
	Properties         *AKSResponseProperties  `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                  `pulumi:"provisioningState"`
	ResourceId         *string                 `pulumi:"resourceId"`
}


func (val *AKSResponse) Defaults() *AKSResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type AKSResponseProperties struct {
	AgentCount                 *int                                `pulumi:"agentCount"`
	AgentVmSize                *string                             `pulumi:"agentVmSize"`
	AksNetworkingConfiguration *AksNetworkingConfigurationResponse `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                *string                             `pulumi:"clusterFqdn"`
	ClusterPurpose             *string                             `pulumi:"clusterPurpose"`
	LoadBalancerSubnet         *string                             `pulumi:"loadBalancerSubnet"`
	LoadBalancerType           *string                             `pulumi:"loadBalancerType"`
	SslConfiguration           *SslConfigurationResponse           `pulumi:"sslConfiguration"`
	SystemServices             []SystemServiceResponse             `pulumi:"systemServices"`
}


func (val *AKSResponseProperties) Defaults() *AKSResponseProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClusterPurpose) {
		clusterPurpose_ := "FastProd"
		tmp.ClusterPurpose = &clusterPurpose_
	}
	if isZero(tmp.LoadBalancerType) {
		loadBalancerType_ := "PublicIp"
		tmp.LoadBalancerType = &loadBalancerType_
	}
	return &tmp
}

type AccountKeyDatastoreCredentials struct {
	CredentialsType string                      `pulumi:"credentialsType"`
	Secrets         *AccountKeyDatastoreSecrets `pulumi:"secrets"`
}

type AccountKeyDatastoreCredentialsResponse struct {
	CredentialsType string `pulumi:"credentialsType"`
}

type AccountKeyDatastoreSecrets struct {
	Key         *string `pulumi:"key"`
	SecretsType string  `pulumi:"secretsType"`
}

type AksNetworkingConfiguration struct {
	DnsServiceIP     *string `pulumi:"dnsServiceIP"`
	DockerBridgeCidr *string `pulumi:"dockerBridgeCidr"`
	ServiceCidr      *string `pulumi:"serviceCidr"`
	SubnetId         *string `pulumi:"subnetId"`
}

type AksNetworkingConfigurationResponse struct {
	DnsServiceIP     *string `pulumi:"dnsServiceIP"`
	DockerBridgeCidr *string `pulumi:"dockerBridgeCidr"`
	ServiceCidr      *string `pulumi:"serviceCidr"`
	SubnetId         *string `pulumi:"subnetId"`
}

type AmlCompute struct {
	ComputeLocation  *string               `pulumi:"computeLocation"`
	ComputeType      string                `pulumi:"computeType"`
	Description      *string               `pulumi:"description"`
	DisableLocalAuth *bool                 `pulumi:"disableLocalAuth"`
	Properties       *AmlComputeProperties `pulumi:"properties"`
	ResourceId       *string               `pulumi:"resourceId"`
}


func (val *AmlCompute) Defaults() *AmlCompute {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type AmlComputeNodeInformationResponse struct {
	NodeId           string  `pulumi:"nodeId"`
	NodeState        string  `pulumi:"nodeState"`
	Port             float64 `pulumi:"port"`
	PrivateIpAddress string  `pulumi:"privateIpAddress"`
	PublicIpAddress  string  `pulumi:"publicIpAddress"`
	RunId            string  `pulumi:"runId"`
}

type AmlComputeProperties struct {
	EnableNodePublicIp          *bool                   `pulumi:"enableNodePublicIp"`
	IsolatedNetwork             *bool                   `pulumi:"isolatedNetwork"`
	OsType                      *string                 `pulumi:"osType"`
	RemoteLoginPortPublicAccess *string                 `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings               *ScaleSettings          `pulumi:"scaleSettings"`
	Subnet                      *ResourceId             `pulumi:"subnet"`
	UserAccountCredentials      *UserAccountCredentials `pulumi:"userAccountCredentials"`
	VirtualMachineImage         *VirtualMachineImage    `pulumi:"virtualMachineImage"`
	VmPriority                  *string                 `pulumi:"vmPriority"`
	VmSize                      *string                 `pulumi:"vmSize"`
}


func (val *AmlComputeProperties) Defaults() *AmlComputeProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableNodePublicIp) {
		enableNodePublicIp_ := true
		tmp.EnableNodePublicIp = &enableNodePublicIp_
	}
	if isZero(tmp.OsType) {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	if isZero(tmp.RemoteLoginPortPublicAccess) {
		remoteLoginPortPublicAccess_ := "NotSpecified"
		tmp.RemoteLoginPortPublicAccess = &remoteLoginPortPublicAccess_
	}
	tmp.ScaleSettings = tmp.ScaleSettings.Defaults()

	return &tmp
}

type AmlComputeResponse struct {
	ComputeLocation    *string                       `pulumi:"computeLocation"`
	ComputeType        string                        `pulumi:"computeType"`
	CreatedOn          string                        `pulumi:"createdOn"`
	Description        *string                       `pulumi:"description"`
	DisableLocalAuth   *bool                         `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                          `pulumi:"isAttachedCompute"`
	ModifiedOn         string                        `pulumi:"modifiedOn"`
	Properties         *AmlComputeResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse       `pulumi:"provisioningErrors"`
	ProvisioningState  string                        `pulumi:"provisioningState"`
	ResourceId         *string                       `pulumi:"resourceId"`
}


func (val *AmlComputeResponse) Defaults() *AmlComputeResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type AmlComputeResponseProperties struct {
	AllocationState               string                          `pulumi:"allocationState"`
	AllocationStateTransitionTime string                          `pulumi:"allocationStateTransitionTime"`
	CurrentNodeCount              int                             `pulumi:"currentNodeCount"`
	EnableNodePublicIp            *bool                           `pulumi:"enableNodePublicIp"`
	Errors                        []ErrorResponseResponse         `pulumi:"errors"`
	IsolatedNetwork               *bool                           `pulumi:"isolatedNetwork"`
	NodeStateCounts               NodeStateCountsResponse         `pulumi:"nodeStateCounts"`
	OsType                        *string                         `pulumi:"osType"`
	RemoteLoginPortPublicAccess   *string                         `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings                 *ScaleSettingsResponse          `pulumi:"scaleSettings"`
	Subnet                        *ResourceIdResponse             `pulumi:"subnet"`
	TargetNodeCount               int                             `pulumi:"targetNodeCount"`
	UserAccountCredentials        *UserAccountCredentialsResponse `pulumi:"userAccountCredentials"`
	VirtualMachineImage           *VirtualMachineImageResponse    `pulumi:"virtualMachineImage"`
	VmPriority                    *string                         `pulumi:"vmPriority"`
	VmSize                        *string                         `pulumi:"vmSize"`
}


func (val *AmlComputeResponseProperties) Defaults() *AmlComputeResponseProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableNodePublicIp) {
		enableNodePublicIp_ := true
		tmp.EnableNodePublicIp = &enableNodePublicIp_
	}
	if isZero(tmp.OsType) {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	if isZero(tmp.RemoteLoginPortPublicAccess) {
		remoteLoginPortPublicAccess_ := "NotSpecified"
		tmp.RemoteLoginPortPublicAccess = &remoteLoginPortPublicAccess_
	}
	tmp.ScaleSettings = tmp.ScaleSettings.Defaults()

	return &tmp
}

type AmlToken struct {
	IdentityType string `pulumi:"identityType"`
}

type AmlTokenResponse struct {
	IdentityType string `pulumi:"identityType"`
}

type AssignedUser struct {
	ObjectId string `pulumi:"objectId"`
	TenantId string `pulumi:"tenantId"`
}

type AssignedUserResponse struct {
	ObjectId string `pulumi:"objectId"`
	TenantId string `pulumi:"tenantId"`
}

type AutoPauseProperties struct {
	DelayInMinutes *int  `pulumi:"delayInMinutes"`
	Enabled        *bool `pulumi:"enabled"`
}

type AutoPausePropertiesResponse struct {
	DelayInMinutes *int  `pulumi:"delayInMinutes"`
	Enabled        *bool `pulumi:"enabled"`
}

type AutoScaleProperties struct {
	Enabled      *bool `pulumi:"enabled"`
	MaxNodeCount *int  `pulumi:"maxNodeCount"`
	MinNodeCount *int  `pulumi:"minNodeCount"`
}

type AutoScalePropertiesResponse struct {
	Enabled      *bool `pulumi:"enabled"`
	MaxNodeCount *int  `pulumi:"maxNodeCount"`
	MinNodeCount *int  `pulumi:"minNodeCount"`
}

type AutoScaleSettings struct {
	MaxInstances                *int    `pulumi:"maxInstances"`
	MinInstances                *int    `pulumi:"minInstances"`
	PollingInterval             *string `pulumi:"pollingInterval"`
	ScaleType                   string  `pulumi:"scaleType"`
	TargetUtilizationPercentage *int    `pulumi:"targetUtilizationPercentage"`
}

type AutoScaleSettingsResponse struct {
	MaxInstances                *int    `pulumi:"maxInstances"`
	MinInstances                *int    `pulumi:"minInstances"`
	PollingInterval             *string `pulumi:"pollingInterval"`
	ScaleType                   string  `pulumi:"scaleType"`
	TargetUtilizationPercentage *int    `pulumi:"targetUtilizationPercentage"`
}

type AzureBlobContents struct {
	AccountName   string      `pulumi:"accountName"`
	ContainerName string      `pulumi:"containerName"`
	ContentsType  string      `pulumi:"contentsType"`
	Credentials   interface{} `pulumi:"credentials"`
	Endpoint      string      `pulumi:"endpoint"`
	Protocol      string      `pulumi:"protocol"`
}

type AzureBlobContentsResponse struct {
	AccountName   string      `pulumi:"accountName"`
	ContainerName string      `pulumi:"containerName"`
	ContentsType  string      `pulumi:"contentsType"`
	Credentials   interface{} `pulumi:"credentials"`
	Endpoint      string      `pulumi:"endpoint"`
	Protocol      string      `pulumi:"protocol"`
}

type AzureDataLakeGen1Contents struct {
	ContentsType string      `pulumi:"contentsType"`
	Credentials  interface{} `pulumi:"credentials"`
	StoreName    string      `pulumi:"storeName"`
}

type AzureDataLakeGen1ContentsResponse struct {
	ContentsType string      `pulumi:"contentsType"`
	Credentials  interface{} `pulumi:"credentials"`
	StoreName    string      `pulumi:"storeName"`
}

type AzureDataLakeGen2Contents struct {
	AccountName   string      `pulumi:"accountName"`
	ContainerName string      `pulumi:"containerName"`
	ContentsType  string      `pulumi:"contentsType"`
	Credentials   interface{} `pulumi:"credentials"`
	Endpoint      string      `pulumi:"endpoint"`
	Protocol      string      `pulumi:"protocol"`
}

type AzureDataLakeGen2ContentsResponse struct {
	AccountName   string      `pulumi:"accountName"`
	ContainerName string      `pulumi:"containerName"`
	ContentsType  string      `pulumi:"contentsType"`
	Credentials   interface{} `pulumi:"credentials"`
	Endpoint      string      `pulumi:"endpoint"`
	Protocol      string      `pulumi:"protocol"`
}

type AzureFileContents struct {
	AccountName   string      `pulumi:"accountName"`
	ContainerName string      `pulumi:"containerName"`
	ContentsType  string      `pulumi:"contentsType"`
	Credentials   interface{} `pulumi:"credentials"`
	Endpoint      string      `pulumi:"endpoint"`
	Protocol      string      `pulumi:"protocol"`
}

type AzureFileContentsResponse struct {
	AccountName   string      `pulumi:"accountName"`
	ContainerName string      `pulumi:"containerName"`
	ContentsType  string      `pulumi:"contentsType"`
	Credentials   interface{} `pulumi:"credentials"`
	Endpoint      string      `pulumi:"endpoint"`
	Protocol      string      `pulumi:"protocol"`
}

type AzurePostgreSqlContents struct {
	ContentsType string      `pulumi:"contentsType"`
	Credentials  interface{} `pulumi:"credentials"`
	DatabaseName string      `pulumi:"databaseName"`
	EnableSSL    *bool       `pulumi:"enableSSL"`
	Endpoint     string      `pulumi:"endpoint"`
	PortNumber   int         `pulumi:"portNumber"`
	ServerName   string      `pulumi:"serverName"`
}

type AzurePostgreSqlContentsResponse struct {
	ContentsType string      `pulumi:"contentsType"`
	Credentials  interface{} `pulumi:"credentials"`
	DatabaseName string      `pulumi:"databaseName"`
	EnableSSL    *bool       `pulumi:"enableSSL"`
	Endpoint     string      `pulumi:"endpoint"`
	PortNumber   int         `pulumi:"portNumber"`
	ServerName   string      `pulumi:"serverName"`
}

type AzureSqlDatabaseContents struct {
	ContentsType string      `pulumi:"contentsType"`
	Credentials  interface{} `pulumi:"credentials"`
	DatabaseName string      `pulumi:"databaseName"`
	Endpoint     string      `pulumi:"endpoint"`
	PortNumber   int         `pulumi:"portNumber"`
	ServerName   string      `pulumi:"serverName"`
}

type AzureSqlDatabaseContentsResponse struct {
	ContentsType string      `pulumi:"contentsType"`
	Credentials  interface{} `pulumi:"credentials"`
	DatabaseName string      `pulumi:"databaseName"`
	Endpoint     string      `pulumi:"endpoint"`
	PortNumber   int         `pulumi:"portNumber"`
	ServerName   string      `pulumi:"serverName"`
}

type BanditPolicy struct {
	DelayEvaluation    *int     `pulumi:"delayEvaluation"`
	EvaluationInterval *int     `pulumi:"evaluationInterval"`
	PolicyType         string   `pulumi:"policyType"`
	SlackAmount        *float64 `pulumi:"slackAmount"`
	SlackFactor        *float64 `pulumi:"slackFactor"`
}

type BanditPolicyResponse struct {
	DelayEvaluation    *int     `pulumi:"delayEvaluation"`
	EvaluationInterval *int     `pulumi:"evaluationInterval"`
	PolicyType         string   `pulumi:"policyType"`
	SlackAmount        *float64 `pulumi:"slackAmount"`
	SlackFactor        *float64 `pulumi:"slackFactor"`
}

type BatchDeploymentType struct {
	CodeConfiguration    *CodeConfiguration        `pulumi:"codeConfiguration"`
	Compute              *ComputeConfiguration     `pulumi:"compute"`
	Description          *string                   `pulumi:"description"`
	EnvironmentId        *string                   `pulumi:"environmentId"`
	EnvironmentVariables map[string]string         `pulumi:"environmentVariables"`
	ErrorThreshold       *int                      `pulumi:"errorThreshold"`
	LoggingLevel         *string                   `pulumi:"loggingLevel"`
	MiniBatchSize        *float64                  `pulumi:"miniBatchSize"`
	Model                interface{}               `pulumi:"model"`
	OutputConfiguration  *BatchOutputConfiguration `pulumi:"outputConfiguration"`
	PartitionKeys        []string                  `pulumi:"partitionKeys"`
	Properties           map[string]string         `pulumi:"properties"`
	RetrySettings        *BatchRetrySettings       `pulumi:"retrySettings"`
}





type BatchDeploymentTypeInput interface {
	pulumi.Input

	ToBatchDeploymentTypeOutput() BatchDeploymentTypeOutput
	ToBatchDeploymentTypeOutputWithContext(context.Context) BatchDeploymentTypeOutput
}

type BatchDeploymentTypeArgs struct {
	CodeConfiguration    CodeConfigurationPtrInput        `pulumi:"codeConfiguration"`
	Compute              ComputeConfigurationPtrInput     `pulumi:"compute"`
	Description          pulumi.StringPtrInput            `pulumi:"description"`
	EnvironmentId        pulumi.StringPtrInput            `pulumi:"environmentId"`
	EnvironmentVariables pulumi.StringMapInput            `pulumi:"environmentVariables"`
	ErrorThreshold       pulumi.IntPtrInput               `pulumi:"errorThreshold"`
	LoggingLevel         pulumi.StringPtrInput            `pulumi:"loggingLevel"`
	MiniBatchSize        pulumi.Float64PtrInput           `pulumi:"miniBatchSize"`
	Model                pulumi.Input                     `pulumi:"model"`
	OutputConfiguration  BatchOutputConfigurationPtrInput `pulumi:"outputConfiguration"`
	PartitionKeys        pulumi.StringArrayInput          `pulumi:"partitionKeys"`
	Properties           pulumi.StringMapInput            `pulumi:"properties"`
	RetrySettings        BatchRetrySettingsPtrInput       `pulumi:"retrySettings"`
}

func (BatchDeploymentTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchDeploymentType)(nil)).Elem()
}

func (i BatchDeploymentTypeArgs) ToBatchDeploymentTypeOutput() BatchDeploymentTypeOutput {
	return i.ToBatchDeploymentTypeOutputWithContext(context.Background())
}

func (i BatchDeploymentTypeArgs) ToBatchDeploymentTypeOutputWithContext(ctx context.Context) BatchDeploymentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchDeploymentTypeOutput)
}

type BatchDeploymentTypeOutput struct{ *pulumi.OutputState }

func (BatchDeploymentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchDeploymentType)(nil)).Elem()
}

func (o BatchDeploymentTypeOutput) ToBatchDeploymentTypeOutput() BatchDeploymentTypeOutput {
	return o
}

func (o BatchDeploymentTypeOutput) ToBatchDeploymentTypeOutputWithContext(ctx context.Context) BatchDeploymentTypeOutput {
	return o
}

func (o BatchDeploymentTypeOutput) CodeConfiguration() CodeConfigurationPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *CodeConfiguration { return v.CodeConfiguration }).(CodeConfigurationPtrOutput)
}

func (o BatchDeploymentTypeOutput) Compute() ComputeConfigurationPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *ComputeConfiguration { return v.Compute }).(ComputeConfigurationPtrOutput)
}

func (o BatchDeploymentTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BatchDeploymentTypeOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o BatchDeploymentTypeOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v BatchDeploymentType) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o BatchDeploymentTypeOutput) ErrorThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *int { return v.ErrorThreshold }).(pulumi.IntPtrOutput)
}

func (o BatchDeploymentTypeOutput) LoggingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *string { return v.LoggingLevel }).(pulumi.StringPtrOutput)
}

func (o BatchDeploymentTypeOutput) MiniBatchSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *float64 { return v.MiniBatchSize }).(pulumi.Float64PtrOutput)
}

func (o BatchDeploymentTypeOutput) Model() pulumi.AnyOutput {
	return o.ApplyT(func(v BatchDeploymentType) interface{} { return v.Model }).(pulumi.AnyOutput)
}

func (o BatchDeploymentTypeOutput) OutputConfiguration() BatchOutputConfigurationPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *BatchOutputConfiguration { return v.OutputConfiguration }).(BatchOutputConfigurationPtrOutput)
}

func (o BatchDeploymentTypeOutput) PartitionKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BatchDeploymentType) []string { return v.PartitionKeys }).(pulumi.StringArrayOutput)
}

func (o BatchDeploymentTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v BatchDeploymentType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o BatchDeploymentTypeOutput) RetrySettings() BatchRetrySettingsPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *BatchRetrySettings { return v.RetrySettings }).(BatchRetrySettingsPtrOutput)
}

type BatchDeploymentResponse struct {
	CodeConfiguration    *CodeConfigurationResponse        `pulumi:"codeConfiguration"`
	Compute              *ComputeConfigurationResponse     `pulumi:"compute"`
	Description          *string                           `pulumi:"description"`
	EnvironmentId        *string                           `pulumi:"environmentId"`
	EnvironmentVariables map[string]string                 `pulumi:"environmentVariables"`
	ErrorThreshold       *int                              `pulumi:"errorThreshold"`
	LoggingLevel         *string                           `pulumi:"loggingLevel"`
	MiniBatchSize        *float64                          `pulumi:"miniBatchSize"`
	Model                interface{}                       `pulumi:"model"`
	OutputConfiguration  *BatchOutputConfigurationResponse `pulumi:"outputConfiguration"`
	PartitionKeys        []string                          `pulumi:"partitionKeys"`
	Properties           map[string]string                 `pulumi:"properties"`
	RetrySettings        *BatchRetrySettingsResponse       `pulumi:"retrySettings"`
}

type BatchDeploymentResponseOutput struct{ *pulumi.OutputState }

func (BatchDeploymentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchDeploymentResponse)(nil)).Elem()
}

func (o BatchDeploymentResponseOutput) ToBatchDeploymentResponseOutput() BatchDeploymentResponseOutput {
	return o
}

func (o BatchDeploymentResponseOutput) ToBatchDeploymentResponseOutputWithContext(ctx context.Context) BatchDeploymentResponseOutput {
	return o
}

func (o BatchDeploymentResponseOutput) CodeConfiguration() CodeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *CodeConfigurationResponse { return v.CodeConfiguration }).(CodeConfigurationResponsePtrOutput)
}

func (o BatchDeploymentResponseOutput) Compute() ComputeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *ComputeConfigurationResponse { return v.Compute }).(ComputeConfigurationResponsePtrOutput)
}

func (o BatchDeploymentResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BatchDeploymentResponseOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o BatchDeploymentResponseOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o BatchDeploymentResponseOutput) ErrorThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *int { return v.ErrorThreshold }).(pulumi.IntPtrOutput)
}

func (o BatchDeploymentResponseOutput) LoggingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *string { return v.LoggingLevel }).(pulumi.StringPtrOutput)
}

func (o BatchDeploymentResponseOutput) MiniBatchSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *float64 { return v.MiniBatchSize }).(pulumi.Float64PtrOutput)
}

func (o BatchDeploymentResponseOutput) Model() pulumi.AnyOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) interface{} { return v.Model }).(pulumi.AnyOutput)
}

func (o BatchDeploymentResponseOutput) OutputConfiguration() BatchOutputConfigurationResponsePtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *BatchOutputConfigurationResponse { return v.OutputConfiguration }).(BatchOutputConfigurationResponsePtrOutput)
}

func (o BatchDeploymentResponseOutput) PartitionKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) []string { return v.PartitionKeys }).(pulumi.StringArrayOutput)
}

func (o BatchDeploymentResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o BatchDeploymentResponseOutput) RetrySettings() BatchRetrySettingsResponsePtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *BatchRetrySettingsResponse { return v.RetrySettings }).(BatchRetrySettingsResponsePtrOutput)
}

type BatchEndpointType struct {
	AuthMode    *string           `pulumi:"authMode"`
	Description *string           `pulumi:"description"`
	Keys        *EndpointAuthKeys `pulumi:"keys"`
	Properties  map[string]string `pulumi:"properties"`
	Traffic     map[string]int    `pulumi:"traffic"`
}





type BatchEndpointTypeInput interface {
	pulumi.Input

	ToBatchEndpointTypeOutput() BatchEndpointTypeOutput
	ToBatchEndpointTypeOutputWithContext(context.Context) BatchEndpointTypeOutput
}

type BatchEndpointTypeArgs struct {
	AuthMode    pulumi.StringPtrInput    `pulumi:"authMode"`
	Description pulumi.StringPtrInput    `pulumi:"description"`
	Keys        EndpointAuthKeysPtrInput `pulumi:"keys"`
	Properties  pulumi.StringMapInput    `pulumi:"properties"`
	Traffic     pulumi.IntMapInput       `pulumi:"traffic"`
}

func (BatchEndpointTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchEndpointType)(nil)).Elem()
}

func (i BatchEndpointTypeArgs) ToBatchEndpointTypeOutput() BatchEndpointTypeOutput {
	return i.ToBatchEndpointTypeOutputWithContext(context.Background())
}

func (i BatchEndpointTypeArgs) ToBatchEndpointTypeOutputWithContext(ctx context.Context) BatchEndpointTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchEndpointTypeOutput)
}

type BatchEndpointTypeOutput struct{ *pulumi.OutputState }

func (BatchEndpointTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchEndpointType)(nil)).Elem()
}

func (o BatchEndpointTypeOutput) ToBatchEndpointTypeOutput() BatchEndpointTypeOutput {
	return o
}

func (o BatchEndpointTypeOutput) ToBatchEndpointTypeOutputWithContext(ctx context.Context) BatchEndpointTypeOutput {
	return o
}

func (o BatchEndpointTypeOutput) AuthMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchEndpointType) *string { return v.AuthMode }).(pulumi.StringPtrOutput)
}

func (o BatchEndpointTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchEndpointType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BatchEndpointTypeOutput) Keys() EndpointAuthKeysPtrOutput {
	return o.ApplyT(func(v BatchEndpointType) *EndpointAuthKeys { return v.Keys }).(EndpointAuthKeysPtrOutput)
}

func (o BatchEndpointTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v BatchEndpointType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o BatchEndpointTypeOutput) Traffic() pulumi.IntMapOutput {
	return o.ApplyT(func(v BatchEndpointType) map[string]int { return v.Traffic }).(pulumi.IntMapOutput)
}

type BatchEndpointResponse struct {
	AuthMode    *string           `pulumi:"authMode"`
	Description *string           `pulumi:"description"`
	Properties  map[string]string `pulumi:"properties"`
	ScoringUri  string            `pulumi:"scoringUri"`
	SwaggerUri  string            `pulumi:"swaggerUri"`
	Traffic     map[string]int    `pulumi:"traffic"`
}

type BatchEndpointResponseOutput struct{ *pulumi.OutputState }

func (BatchEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchEndpointResponse)(nil)).Elem()
}

func (o BatchEndpointResponseOutput) ToBatchEndpointResponseOutput() BatchEndpointResponseOutput {
	return o
}

func (o BatchEndpointResponseOutput) ToBatchEndpointResponseOutputWithContext(ctx context.Context) BatchEndpointResponseOutput {
	return o
}

func (o BatchEndpointResponseOutput) AuthMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchEndpointResponse) *string { return v.AuthMode }).(pulumi.StringPtrOutput)
}

func (o BatchEndpointResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchEndpointResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BatchEndpointResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v BatchEndpointResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o BatchEndpointResponseOutput) ScoringUri() pulumi.StringOutput {
	return o.ApplyT(func(v BatchEndpointResponse) string { return v.ScoringUri }).(pulumi.StringOutput)
}

func (o BatchEndpointResponseOutput) SwaggerUri() pulumi.StringOutput {
	return o.ApplyT(func(v BatchEndpointResponse) string { return v.SwaggerUri }).(pulumi.StringOutput)
}

func (o BatchEndpointResponseOutput) Traffic() pulumi.IntMapOutput {
	return o.ApplyT(func(v BatchEndpointResponse) map[string]int { return v.Traffic }).(pulumi.IntMapOutput)
}

type BatchOutputConfiguration struct {
	AppendRowFileName *string `pulumi:"appendRowFileName"`
	OutputAction      *string `pulumi:"outputAction"`
}





type BatchOutputConfigurationInput interface {
	pulumi.Input

	ToBatchOutputConfigurationOutput() BatchOutputConfigurationOutput
	ToBatchOutputConfigurationOutputWithContext(context.Context) BatchOutputConfigurationOutput
}

type BatchOutputConfigurationArgs struct {
	AppendRowFileName pulumi.StringPtrInput `pulumi:"appendRowFileName"`
	OutputAction      pulumi.StringPtrInput `pulumi:"outputAction"`
}

func (BatchOutputConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchOutputConfiguration)(nil)).Elem()
}

func (i BatchOutputConfigurationArgs) ToBatchOutputConfigurationOutput() BatchOutputConfigurationOutput {
	return i.ToBatchOutputConfigurationOutputWithContext(context.Background())
}

func (i BatchOutputConfigurationArgs) ToBatchOutputConfigurationOutputWithContext(ctx context.Context) BatchOutputConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchOutputConfigurationOutput)
}

func (i BatchOutputConfigurationArgs) ToBatchOutputConfigurationPtrOutput() BatchOutputConfigurationPtrOutput {
	return i.ToBatchOutputConfigurationPtrOutputWithContext(context.Background())
}

func (i BatchOutputConfigurationArgs) ToBatchOutputConfigurationPtrOutputWithContext(ctx context.Context) BatchOutputConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchOutputConfigurationOutput).ToBatchOutputConfigurationPtrOutputWithContext(ctx)
}









type BatchOutputConfigurationPtrInput interface {
	pulumi.Input

	ToBatchOutputConfigurationPtrOutput() BatchOutputConfigurationPtrOutput
	ToBatchOutputConfigurationPtrOutputWithContext(context.Context) BatchOutputConfigurationPtrOutput
}

type batchOutputConfigurationPtrType BatchOutputConfigurationArgs

func BatchOutputConfigurationPtr(v *BatchOutputConfigurationArgs) BatchOutputConfigurationPtrInput {
	return (*batchOutputConfigurationPtrType)(v)
}

func (*batchOutputConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchOutputConfiguration)(nil)).Elem()
}

func (i *batchOutputConfigurationPtrType) ToBatchOutputConfigurationPtrOutput() BatchOutputConfigurationPtrOutput {
	return i.ToBatchOutputConfigurationPtrOutputWithContext(context.Background())
}

func (i *batchOutputConfigurationPtrType) ToBatchOutputConfigurationPtrOutputWithContext(ctx context.Context) BatchOutputConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchOutputConfigurationPtrOutput)
}

type BatchOutputConfigurationOutput struct{ *pulumi.OutputState }

func (BatchOutputConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchOutputConfiguration)(nil)).Elem()
}

func (o BatchOutputConfigurationOutput) ToBatchOutputConfigurationOutput() BatchOutputConfigurationOutput {
	return o
}

func (o BatchOutputConfigurationOutput) ToBatchOutputConfigurationOutputWithContext(ctx context.Context) BatchOutputConfigurationOutput {
	return o
}

func (o BatchOutputConfigurationOutput) ToBatchOutputConfigurationPtrOutput() BatchOutputConfigurationPtrOutput {
	return o.ToBatchOutputConfigurationPtrOutputWithContext(context.Background())
}

func (o BatchOutputConfigurationOutput) ToBatchOutputConfigurationPtrOutputWithContext(ctx context.Context) BatchOutputConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BatchOutputConfiguration) *BatchOutputConfiguration {
		return &v
	}).(BatchOutputConfigurationPtrOutput)
}

func (o BatchOutputConfigurationOutput) AppendRowFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchOutputConfiguration) *string { return v.AppendRowFileName }).(pulumi.StringPtrOutput)
}

func (o BatchOutputConfigurationOutput) OutputAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchOutputConfiguration) *string { return v.OutputAction }).(pulumi.StringPtrOutput)
}

type BatchOutputConfigurationPtrOutput struct{ *pulumi.OutputState }

func (BatchOutputConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchOutputConfiguration)(nil)).Elem()
}

func (o BatchOutputConfigurationPtrOutput) ToBatchOutputConfigurationPtrOutput() BatchOutputConfigurationPtrOutput {
	return o
}

func (o BatchOutputConfigurationPtrOutput) ToBatchOutputConfigurationPtrOutputWithContext(ctx context.Context) BatchOutputConfigurationPtrOutput {
	return o
}

func (o BatchOutputConfigurationPtrOutput) Elem() BatchOutputConfigurationOutput {
	return o.ApplyT(func(v *BatchOutputConfiguration) BatchOutputConfiguration {
		if v != nil {
			return *v
		}
		var ret BatchOutputConfiguration
		return ret
	}).(BatchOutputConfigurationOutput)
}

func (o BatchOutputConfigurationPtrOutput) AppendRowFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchOutputConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AppendRowFileName
	}).(pulumi.StringPtrOutput)
}

func (o BatchOutputConfigurationPtrOutput) OutputAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchOutputConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.OutputAction
	}).(pulumi.StringPtrOutput)
}

type BatchOutputConfigurationResponse struct {
	AppendRowFileName *string `pulumi:"appendRowFileName"`
	OutputAction      *string `pulumi:"outputAction"`
}

type BatchOutputConfigurationResponseOutput struct{ *pulumi.OutputState }

func (BatchOutputConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchOutputConfigurationResponse)(nil)).Elem()
}

func (o BatchOutputConfigurationResponseOutput) ToBatchOutputConfigurationResponseOutput() BatchOutputConfigurationResponseOutput {
	return o
}

func (o BatchOutputConfigurationResponseOutput) ToBatchOutputConfigurationResponseOutputWithContext(ctx context.Context) BatchOutputConfigurationResponseOutput {
	return o
}

func (o BatchOutputConfigurationResponseOutput) AppendRowFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchOutputConfigurationResponse) *string { return v.AppendRowFileName }).(pulumi.StringPtrOutput)
}

func (o BatchOutputConfigurationResponseOutput) OutputAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchOutputConfigurationResponse) *string { return v.OutputAction }).(pulumi.StringPtrOutput)
}

type BatchOutputConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (BatchOutputConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchOutputConfigurationResponse)(nil)).Elem()
}

func (o BatchOutputConfigurationResponsePtrOutput) ToBatchOutputConfigurationResponsePtrOutput() BatchOutputConfigurationResponsePtrOutput {
	return o
}

func (o BatchOutputConfigurationResponsePtrOutput) ToBatchOutputConfigurationResponsePtrOutputWithContext(ctx context.Context) BatchOutputConfigurationResponsePtrOutput {
	return o
}

func (o BatchOutputConfigurationResponsePtrOutput) Elem() BatchOutputConfigurationResponseOutput {
	return o.ApplyT(func(v *BatchOutputConfigurationResponse) BatchOutputConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret BatchOutputConfigurationResponse
		return ret
	}).(BatchOutputConfigurationResponseOutput)
}

func (o BatchOutputConfigurationResponsePtrOutput) AppendRowFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchOutputConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppendRowFileName
	}).(pulumi.StringPtrOutput)
}

func (o BatchOutputConfigurationResponsePtrOutput) OutputAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchOutputConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.OutputAction
	}).(pulumi.StringPtrOutput)
}

type BatchRetrySettings struct {
	MaxRetries *int    `pulumi:"maxRetries"`
	Timeout    *string `pulumi:"timeout"`
}





type BatchRetrySettingsInput interface {
	pulumi.Input

	ToBatchRetrySettingsOutput() BatchRetrySettingsOutput
	ToBatchRetrySettingsOutputWithContext(context.Context) BatchRetrySettingsOutput
}

type BatchRetrySettingsArgs struct {
	MaxRetries pulumi.IntPtrInput    `pulumi:"maxRetries"`
	Timeout    pulumi.StringPtrInput `pulumi:"timeout"`
}

func (BatchRetrySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchRetrySettings)(nil)).Elem()
}

func (i BatchRetrySettingsArgs) ToBatchRetrySettingsOutput() BatchRetrySettingsOutput {
	return i.ToBatchRetrySettingsOutputWithContext(context.Background())
}

func (i BatchRetrySettingsArgs) ToBatchRetrySettingsOutputWithContext(ctx context.Context) BatchRetrySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchRetrySettingsOutput)
}

func (i BatchRetrySettingsArgs) ToBatchRetrySettingsPtrOutput() BatchRetrySettingsPtrOutput {
	return i.ToBatchRetrySettingsPtrOutputWithContext(context.Background())
}

func (i BatchRetrySettingsArgs) ToBatchRetrySettingsPtrOutputWithContext(ctx context.Context) BatchRetrySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchRetrySettingsOutput).ToBatchRetrySettingsPtrOutputWithContext(ctx)
}









type BatchRetrySettingsPtrInput interface {
	pulumi.Input

	ToBatchRetrySettingsPtrOutput() BatchRetrySettingsPtrOutput
	ToBatchRetrySettingsPtrOutputWithContext(context.Context) BatchRetrySettingsPtrOutput
}

type batchRetrySettingsPtrType BatchRetrySettingsArgs

func BatchRetrySettingsPtr(v *BatchRetrySettingsArgs) BatchRetrySettingsPtrInput {
	return (*batchRetrySettingsPtrType)(v)
}

func (*batchRetrySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchRetrySettings)(nil)).Elem()
}

func (i *batchRetrySettingsPtrType) ToBatchRetrySettingsPtrOutput() BatchRetrySettingsPtrOutput {
	return i.ToBatchRetrySettingsPtrOutputWithContext(context.Background())
}

func (i *batchRetrySettingsPtrType) ToBatchRetrySettingsPtrOutputWithContext(ctx context.Context) BatchRetrySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchRetrySettingsPtrOutput)
}

type BatchRetrySettingsOutput struct{ *pulumi.OutputState }

func (BatchRetrySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchRetrySettings)(nil)).Elem()
}

func (o BatchRetrySettingsOutput) ToBatchRetrySettingsOutput() BatchRetrySettingsOutput {
	return o
}

func (o BatchRetrySettingsOutput) ToBatchRetrySettingsOutputWithContext(ctx context.Context) BatchRetrySettingsOutput {
	return o
}

func (o BatchRetrySettingsOutput) ToBatchRetrySettingsPtrOutput() BatchRetrySettingsPtrOutput {
	return o.ToBatchRetrySettingsPtrOutputWithContext(context.Background())
}

func (o BatchRetrySettingsOutput) ToBatchRetrySettingsPtrOutputWithContext(ctx context.Context) BatchRetrySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BatchRetrySettings) *BatchRetrySettings {
		return &v
	}).(BatchRetrySettingsPtrOutput)
}

func (o BatchRetrySettingsOutput) MaxRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BatchRetrySettings) *int { return v.MaxRetries }).(pulumi.IntPtrOutput)
}

func (o BatchRetrySettingsOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchRetrySettings) *string { return v.Timeout }).(pulumi.StringPtrOutput)
}

type BatchRetrySettingsPtrOutput struct{ *pulumi.OutputState }

func (BatchRetrySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchRetrySettings)(nil)).Elem()
}

func (o BatchRetrySettingsPtrOutput) ToBatchRetrySettingsPtrOutput() BatchRetrySettingsPtrOutput {
	return o
}

func (o BatchRetrySettingsPtrOutput) ToBatchRetrySettingsPtrOutputWithContext(ctx context.Context) BatchRetrySettingsPtrOutput {
	return o
}

func (o BatchRetrySettingsPtrOutput) Elem() BatchRetrySettingsOutput {
	return o.ApplyT(func(v *BatchRetrySettings) BatchRetrySettings {
		if v != nil {
			return *v
		}
		var ret BatchRetrySettings
		return ret
	}).(BatchRetrySettingsOutput)
}

func (o BatchRetrySettingsPtrOutput) MaxRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BatchRetrySettings) *int {
		if v == nil {
			return nil
		}
		return v.MaxRetries
	}).(pulumi.IntPtrOutput)
}

func (o BatchRetrySettingsPtrOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchRetrySettings) *string {
		if v == nil {
			return nil
		}
		return v.Timeout
	}).(pulumi.StringPtrOutput)
}

type BatchRetrySettingsResponse struct {
	MaxRetries *int    `pulumi:"maxRetries"`
	Timeout    *string `pulumi:"timeout"`
}

type BatchRetrySettingsResponseOutput struct{ *pulumi.OutputState }

func (BatchRetrySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchRetrySettingsResponse)(nil)).Elem()
}

func (o BatchRetrySettingsResponseOutput) ToBatchRetrySettingsResponseOutput() BatchRetrySettingsResponseOutput {
	return o
}

func (o BatchRetrySettingsResponseOutput) ToBatchRetrySettingsResponseOutputWithContext(ctx context.Context) BatchRetrySettingsResponseOutput {
	return o
}

func (o BatchRetrySettingsResponseOutput) MaxRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BatchRetrySettingsResponse) *int { return v.MaxRetries }).(pulumi.IntPtrOutput)
}

func (o BatchRetrySettingsResponseOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchRetrySettingsResponse) *string { return v.Timeout }).(pulumi.StringPtrOutput)
}

type BatchRetrySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (BatchRetrySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchRetrySettingsResponse)(nil)).Elem()
}

func (o BatchRetrySettingsResponsePtrOutput) ToBatchRetrySettingsResponsePtrOutput() BatchRetrySettingsResponsePtrOutput {
	return o
}

func (o BatchRetrySettingsResponsePtrOutput) ToBatchRetrySettingsResponsePtrOutputWithContext(ctx context.Context) BatchRetrySettingsResponsePtrOutput {
	return o
}

func (o BatchRetrySettingsResponsePtrOutput) Elem() BatchRetrySettingsResponseOutput {
	return o.ApplyT(func(v *BatchRetrySettingsResponse) BatchRetrySettingsResponse {
		if v != nil {
			return *v
		}
		var ret BatchRetrySettingsResponse
		return ret
	}).(BatchRetrySettingsResponseOutput)
}

func (o BatchRetrySettingsResponsePtrOutput) MaxRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BatchRetrySettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxRetries
	}).(pulumi.IntPtrOutput)
}

func (o BatchRetrySettingsResponsePtrOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchRetrySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Timeout
	}).(pulumi.StringPtrOutput)
}

type CertificateDatastoreCredentials struct {
	AuthorityUrl    *string                      `pulumi:"authorityUrl"`
	ClientId        string                       `pulumi:"clientId"`
	CredentialsType string                       `pulumi:"credentialsType"`
	ResourceUri     *string                      `pulumi:"resourceUri"`
	Secrets         *CertificateDatastoreSecrets `pulumi:"secrets"`
	TenantId        string                       `pulumi:"tenantId"`
	Thumbprint      string                       `pulumi:"thumbprint"`
}

type CertificateDatastoreCredentialsResponse struct {
	AuthorityUrl    *string `pulumi:"authorityUrl"`
	ClientId        string  `pulumi:"clientId"`
	CredentialsType string  `pulumi:"credentialsType"`
	ResourceUri     *string `pulumi:"resourceUri"`
	TenantId        string  `pulumi:"tenantId"`
	Thumbprint      string  `pulumi:"thumbprint"`
}

type CertificateDatastoreSecrets struct {
	Certificate *string `pulumi:"certificate"`
	SecretsType string  `pulumi:"secretsType"`
}

type CodeConfiguration struct {
	CodeId        *string `pulumi:"codeId"`
	ScoringScript string  `pulumi:"scoringScript"`
}





type CodeConfigurationInput interface {
	pulumi.Input

	ToCodeConfigurationOutput() CodeConfigurationOutput
	ToCodeConfigurationOutputWithContext(context.Context) CodeConfigurationOutput
}

type CodeConfigurationArgs struct {
	CodeId        pulumi.StringPtrInput `pulumi:"codeId"`
	ScoringScript pulumi.StringInput    `pulumi:"scoringScript"`
}

func (CodeConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeConfiguration)(nil)).Elem()
}

func (i CodeConfigurationArgs) ToCodeConfigurationOutput() CodeConfigurationOutput {
	return i.ToCodeConfigurationOutputWithContext(context.Background())
}

func (i CodeConfigurationArgs) ToCodeConfigurationOutputWithContext(ctx context.Context) CodeConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeConfigurationOutput)
}

func (i CodeConfigurationArgs) ToCodeConfigurationPtrOutput() CodeConfigurationPtrOutput {
	return i.ToCodeConfigurationPtrOutputWithContext(context.Background())
}

func (i CodeConfigurationArgs) ToCodeConfigurationPtrOutputWithContext(ctx context.Context) CodeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeConfigurationOutput).ToCodeConfigurationPtrOutputWithContext(ctx)
}









type CodeConfigurationPtrInput interface {
	pulumi.Input

	ToCodeConfigurationPtrOutput() CodeConfigurationPtrOutput
	ToCodeConfigurationPtrOutputWithContext(context.Context) CodeConfigurationPtrOutput
}

type codeConfigurationPtrType CodeConfigurationArgs

func CodeConfigurationPtr(v *CodeConfigurationArgs) CodeConfigurationPtrInput {
	return (*codeConfigurationPtrType)(v)
}

func (*codeConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeConfiguration)(nil)).Elem()
}

func (i *codeConfigurationPtrType) ToCodeConfigurationPtrOutput() CodeConfigurationPtrOutput {
	return i.ToCodeConfigurationPtrOutputWithContext(context.Background())
}

func (i *codeConfigurationPtrType) ToCodeConfigurationPtrOutputWithContext(ctx context.Context) CodeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeConfigurationPtrOutput)
}

type CodeConfigurationOutput struct{ *pulumi.OutputState }

func (CodeConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeConfiguration)(nil)).Elem()
}

func (o CodeConfigurationOutput) ToCodeConfigurationOutput() CodeConfigurationOutput {
	return o
}

func (o CodeConfigurationOutput) ToCodeConfigurationOutputWithContext(ctx context.Context) CodeConfigurationOutput {
	return o
}

func (o CodeConfigurationOutput) ToCodeConfigurationPtrOutput() CodeConfigurationPtrOutput {
	return o.ToCodeConfigurationPtrOutputWithContext(context.Background())
}

func (o CodeConfigurationOutput) ToCodeConfigurationPtrOutputWithContext(ctx context.Context) CodeConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CodeConfiguration) *CodeConfiguration {
		return &v
	}).(CodeConfigurationPtrOutput)
}

func (o CodeConfigurationOutput) CodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeConfiguration) *string { return v.CodeId }).(pulumi.StringPtrOutput)
}

func (o CodeConfigurationOutput) ScoringScript() pulumi.StringOutput {
	return o.ApplyT(func(v CodeConfiguration) string { return v.ScoringScript }).(pulumi.StringOutput)
}

type CodeConfigurationPtrOutput struct{ *pulumi.OutputState }

func (CodeConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeConfiguration)(nil)).Elem()
}

func (o CodeConfigurationPtrOutput) ToCodeConfigurationPtrOutput() CodeConfigurationPtrOutput {
	return o
}

func (o CodeConfigurationPtrOutput) ToCodeConfigurationPtrOutputWithContext(ctx context.Context) CodeConfigurationPtrOutput {
	return o
}

func (o CodeConfigurationPtrOutput) Elem() CodeConfigurationOutput {
	return o.ApplyT(func(v *CodeConfiguration) CodeConfiguration {
		if v != nil {
			return *v
		}
		var ret CodeConfiguration
		return ret
	}).(CodeConfigurationOutput)
}

func (o CodeConfigurationPtrOutput) CodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CodeId
	}).(pulumi.StringPtrOutput)
}

func (o CodeConfigurationPtrOutput) ScoringScript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.ScoringScript
	}).(pulumi.StringPtrOutput)
}

type CodeConfigurationResponse struct {
	CodeId        *string `pulumi:"codeId"`
	ScoringScript string  `pulumi:"scoringScript"`
}

type CodeConfigurationResponseOutput struct{ *pulumi.OutputState }

func (CodeConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeConfigurationResponse)(nil)).Elem()
}

func (o CodeConfigurationResponseOutput) ToCodeConfigurationResponseOutput() CodeConfigurationResponseOutput {
	return o
}

func (o CodeConfigurationResponseOutput) ToCodeConfigurationResponseOutputWithContext(ctx context.Context) CodeConfigurationResponseOutput {
	return o
}

func (o CodeConfigurationResponseOutput) CodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeConfigurationResponse) *string { return v.CodeId }).(pulumi.StringPtrOutput)
}

func (o CodeConfigurationResponseOutput) ScoringScript() pulumi.StringOutput {
	return o.ApplyT(func(v CodeConfigurationResponse) string { return v.ScoringScript }).(pulumi.StringOutput)
}

type CodeConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (CodeConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeConfigurationResponse)(nil)).Elem()
}

func (o CodeConfigurationResponsePtrOutput) ToCodeConfigurationResponsePtrOutput() CodeConfigurationResponsePtrOutput {
	return o
}

func (o CodeConfigurationResponsePtrOutput) ToCodeConfigurationResponsePtrOutputWithContext(ctx context.Context) CodeConfigurationResponsePtrOutput {
	return o
}

func (o CodeConfigurationResponsePtrOutput) Elem() CodeConfigurationResponseOutput {
	return o.ApplyT(func(v *CodeConfigurationResponse) CodeConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret CodeConfigurationResponse
		return ret
	}).(CodeConfigurationResponseOutput)
}

func (o CodeConfigurationResponsePtrOutput) CodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CodeId
	}).(pulumi.StringPtrOutput)
}

func (o CodeConfigurationResponsePtrOutput) ScoringScript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ScoringScript
	}).(pulumi.StringPtrOutput)
}

type CodeContainerType struct {
	Description *string           `pulumi:"description"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}





type CodeContainerTypeInput interface {
	pulumi.Input

	ToCodeContainerTypeOutput() CodeContainerTypeOutput
	ToCodeContainerTypeOutputWithContext(context.Context) CodeContainerTypeOutput
}

type CodeContainerTypeArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}

func (CodeContainerTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeContainerType)(nil)).Elem()
}

func (i CodeContainerTypeArgs) ToCodeContainerTypeOutput() CodeContainerTypeOutput {
	return i.ToCodeContainerTypeOutputWithContext(context.Background())
}

func (i CodeContainerTypeArgs) ToCodeContainerTypeOutputWithContext(ctx context.Context) CodeContainerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeContainerTypeOutput)
}

type CodeContainerTypeOutput struct{ *pulumi.OutputState }

func (CodeContainerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeContainerType)(nil)).Elem()
}

func (o CodeContainerTypeOutput) ToCodeContainerTypeOutput() CodeContainerTypeOutput {
	return o
}

func (o CodeContainerTypeOutput) ToCodeContainerTypeOutputWithContext(ctx context.Context) CodeContainerTypeOutput {
	return o
}

func (o CodeContainerTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeContainerType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CodeContainerTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeContainerType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CodeContainerTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeContainerType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type CodeContainerResponse struct {
	Description *string           `pulumi:"description"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}

type CodeContainerResponseOutput struct{ *pulumi.OutputState }

func (CodeContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeContainerResponse)(nil)).Elem()
}

func (o CodeContainerResponseOutput) ToCodeContainerResponseOutput() CodeContainerResponseOutput {
	return o
}

func (o CodeContainerResponseOutput) ToCodeContainerResponseOutputWithContext(ctx context.Context) CodeContainerResponseOutput {
	return o
}

func (o CodeContainerResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeContainerResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CodeContainerResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeContainerResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CodeContainerResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeContainerResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type CodeVersionType struct {
	DatastoreId *string           `pulumi:"datastoreId"`
	Description *string           `pulumi:"description"`
	IsAnonymous *bool             `pulumi:"isAnonymous"`
	Path        string            `pulumi:"path"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}





type CodeVersionTypeInput interface {
	pulumi.Input

	ToCodeVersionTypeOutput() CodeVersionTypeOutput
	ToCodeVersionTypeOutputWithContext(context.Context) CodeVersionTypeOutput
}

type CodeVersionTypeArgs struct {
	DatastoreId pulumi.StringPtrInput `pulumi:"datastoreId"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	IsAnonymous pulumi.BoolPtrInput   `pulumi:"isAnonymous"`
	Path        pulumi.StringInput    `pulumi:"path"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}

func (CodeVersionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeVersionType)(nil)).Elem()
}

func (i CodeVersionTypeArgs) ToCodeVersionTypeOutput() CodeVersionTypeOutput {
	return i.ToCodeVersionTypeOutputWithContext(context.Background())
}

func (i CodeVersionTypeArgs) ToCodeVersionTypeOutputWithContext(ctx context.Context) CodeVersionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeVersionTypeOutput)
}

type CodeVersionTypeOutput struct{ *pulumi.OutputState }

func (CodeVersionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeVersionType)(nil)).Elem()
}

func (o CodeVersionTypeOutput) ToCodeVersionTypeOutput() CodeVersionTypeOutput {
	return o
}

func (o CodeVersionTypeOutput) ToCodeVersionTypeOutputWithContext(ctx context.Context) CodeVersionTypeOutput {
	return o
}

func (o CodeVersionTypeOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeVersionType) *string { return v.DatastoreId }).(pulumi.StringPtrOutput)
}

func (o CodeVersionTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeVersionType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CodeVersionTypeOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodeVersionType) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o CodeVersionTypeOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v CodeVersionType) string { return v.Path }).(pulumi.StringOutput)
}

func (o CodeVersionTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeVersionType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CodeVersionTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeVersionType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type CodeVersionResponse struct {
	DatastoreId *string           `pulumi:"datastoreId"`
	Description *string           `pulumi:"description"`
	IsAnonymous *bool             `pulumi:"isAnonymous"`
	Path        string            `pulumi:"path"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}

type CodeVersionResponseOutput struct{ *pulumi.OutputState }

func (CodeVersionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodeVersionResponse)(nil)).Elem()
}

func (o CodeVersionResponseOutput) ToCodeVersionResponseOutput() CodeVersionResponseOutput {
	return o
}

func (o CodeVersionResponseOutput) ToCodeVersionResponseOutputWithContext(ctx context.Context) CodeVersionResponseOutput {
	return o
}

func (o CodeVersionResponseOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeVersionResponse) *string { return v.DatastoreId }).(pulumi.StringPtrOutput)
}

func (o CodeVersionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeVersionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CodeVersionResponseOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodeVersionResponse) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o CodeVersionResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v CodeVersionResponse) string { return v.Path }).(pulumi.StringOutput)
}

func (o CodeVersionResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeVersionResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CodeVersionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeVersionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type CommandJob struct {
	CodeId               *string                      `pulumi:"codeId"`
	Command              string                       `pulumi:"command"`
	Compute              ComputeConfiguration         `pulumi:"compute"`
	Description          *string                      `pulumi:"description"`
	Distribution         interface{}                  `pulumi:"distribution"`
	EnvironmentId        *string                      `pulumi:"environmentId"`
	EnvironmentVariables map[string]string            `pulumi:"environmentVariables"`
	ExperimentName       *string                      `pulumi:"experimentName"`
	Identity             interface{}                  `pulumi:"identity"`
	InputDataBindings    map[string]InputDataBinding  `pulumi:"inputDataBindings"`
	JobType              string                       `pulumi:"jobType"`
	OutputDataBindings   map[string]OutputDataBinding `pulumi:"outputDataBindings"`
	Priority             *int                         `pulumi:"priority"`
	Properties           map[string]string            `pulumi:"properties"`
	Tags                 map[string]string            `pulumi:"tags"`
	Timeout              *string                      `pulumi:"timeout"`
}

type CommandJobResponse struct {
	CodeId               *string                              `pulumi:"codeId"`
	Command              string                               `pulumi:"command"`
	Compute              ComputeConfigurationResponse         `pulumi:"compute"`
	Description          *string                              `pulumi:"description"`
	Distribution         interface{}                          `pulumi:"distribution"`
	EnvironmentId        *string                              `pulumi:"environmentId"`
	EnvironmentVariables map[string]string                    `pulumi:"environmentVariables"`
	ExperimentName       *string                              `pulumi:"experimentName"`
	Identity             interface{}                          `pulumi:"identity"`
	InputDataBindings    map[string]InputDataBindingResponse  `pulumi:"inputDataBindings"`
	InteractionEndpoints map[string]JobEndpointResponse       `pulumi:"interactionEndpoints"`
	JobType              string                               `pulumi:"jobType"`
	Output               JobOutputResponse                    `pulumi:"output"`
	OutputDataBindings   map[string]OutputDataBindingResponse `pulumi:"outputDataBindings"`
	Parameters           interface{}                          `pulumi:"parameters"`
	Priority             *int                                 `pulumi:"priority"`
	Properties           map[string]string                    `pulumi:"properties"`
	ProvisioningState    string                               `pulumi:"provisioningState"`
	Status               string                               `pulumi:"status"`
	Tags                 map[string]string                    `pulumi:"tags"`
	Timeout              *string                              `pulumi:"timeout"`
}

type ComputeConfiguration struct {
	InstanceCount *int              `pulumi:"instanceCount"`
	InstanceType  *string           `pulumi:"instanceType"`
	IsLocal       *bool             `pulumi:"isLocal"`
	Location      *string           `pulumi:"location"`
	Properties    map[string]string `pulumi:"properties"`
	Target        *string           `pulumi:"target"`
}





type ComputeConfigurationInput interface {
	pulumi.Input

	ToComputeConfigurationOutput() ComputeConfigurationOutput
	ToComputeConfigurationOutputWithContext(context.Context) ComputeConfigurationOutput
}

type ComputeConfigurationArgs struct {
	InstanceCount pulumi.IntPtrInput    `pulumi:"instanceCount"`
	InstanceType  pulumi.StringPtrInput `pulumi:"instanceType"`
	IsLocal       pulumi.BoolPtrInput   `pulumi:"isLocal"`
	Location      pulumi.StringPtrInput `pulumi:"location"`
	Properties    pulumi.StringMapInput `pulumi:"properties"`
	Target        pulumi.StringPtrInput `pulumi:"target"`
}

func (ComputeConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeConfiguration)(nil)).Elem()
}

func (i ComputeConfigurationArgs) ToComputeConfigurationOutput() ComputeConfigurationOutput {
	return i.ToComputeConfigurationOutputWithContext(context.Background())
}

func (i ComputeConfigurationArgs) ToComputeConfigurationOutputWithContext(ctx context.Context) ComputeConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeConfigurationOutput)
}

func (i ComputeConfigurationArgs) ToComputeConfigurationPtrOutput() ComputeConfigurationPtrOutput {
	return i.ToComputeConfigurationPtrOutputWithContext(context.Background())
}

func (i ComputeConfigurationArgs) ToComputeConfigurationPtrOutputWithContext(ctx context.Context) ComputeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeConfigurationOutput).ToComputeConfigurationPtrOutputWithContext(ctx)
}









type ComputeConfigurationPtrInput interface {
	pulumi.Input

	ToComputeConfigurationPtrOutput() ComputeConfigurationPtrOutput
	ToComputeConfigurationPtrOutputWithContext(context.Context) ComputeConfigurationPtrOutput
}

type computeConfigurationPtrType ComputeConfigurationArgs

func ComputeConfigurationPtr(v *ComputeConfigurationArgs) ComputeConfigurationPtrInput {
	return (*computeConfigurationPtrType)(v)
}

func (*computeConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeConfiguration)(nil)).Elem()
}

func (i *computeConfigurationPtrType) ToComputeConfigurationPtrOutput() ComputeConfigurationPtrOutput {
	return i.ToComputeConfigurationPtrOutputWithContext(context.Background())
}

func (i *computeConfigurationPtrType) ToComputeConfigurationPtrOutputWithContext(ctx context.Context) ComputeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeConfigurationPtrOutput)
}

type ComputeConfigurationOutput struct{ *pulumi.OutputState }

func (ComputeConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeConfiguration)(nil)).Elem()
}

func (o ComputeConfigurationOutput) ToComputeConfigurationOutput() ComputeConfigurationOutput {
	return o
}

func (o ComputeConfigurationOutput) ToComputeConfigurationOutputWithContext(ctx context.Context) ComputeConfigurationOutput {
	return o
}

func (o ComputeConfigurationOutput) ToComputeConfigurationPtrOutput() ComputeConfigurationPtrOutput {
	return o.ToComputeConfigurationPtrOutputWithContext(context.Background())
}

func (o ComputeConfigurationOutput) ToComputeConfigurationPtrOutputWithContext(ctx context.Context) ComputeConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeConfiguration) *ComputeConfiguration {
		return &v
	}).(ComputeConfigurationPtrOutput)
}

func (o ComputeConfigurationOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ComputeConfiguration) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

func (o ComputeConfigurationOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeConfiguration) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o ComputeConfigurationOutput) IsLocal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComputeConfiguration) *bool { return v.IsLocal }).(pulumi.BoolPtrOutput)
}

func (o ComputeConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeConfiguration) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ComputeConfigurationOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ComputeConfiguration) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ComputeConfigurationOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeConfiguration) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ComputeConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ComputeConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeConfiguration)(nil)).Elem()
}

func (o ComputeConfigurationPtrOutput) ToComputeConfigurationPtrOutput() ComputeConfigurationPtrOutput {
	return o
}

func (o ComputeConfigurationPtrOutput) ToComputeConfigurationPtrOutputWithContext(ctx context.Context) ComputeConfigurationPtrOutput {
	return o
}

func (o ComputeConfigurationPtrOutput) Elem() ComputeConfigurationOutput {
	return o.ApplyT(func(v *ComputeConfiguration) ComputeConfiguration {
		if v != nil {
			return *v
		}
		var ret ComputeConfiguration
		return ret
	}).(ComputeConfigurationOutput)
}

func (o ComputeConfigurationPtrOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ComputeConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.InstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o ComputeConfigurationPtrOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.InstanceType
	}).(pulumi.StringPtrOutput)
}

func (o ComputeConfigurationPtrOutput) IsLocal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.IsLocal
	}).(pulumi.BoolPtrOutput)
}

func (o ComputeConfigurationPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ComputeConfigurationPtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeConfiguration) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o ComputeConfigurationPtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type ComputeConfigurationResponse struct {
	InstanceCount *int              `pulumi:"instanceCount"`
	InstanceType  *string           `pulumi:"instanceType"`
	IsLocal       *bool             `pulumi:"isLocal"`
	Location      *string           `pulumi:"location"`
	Properties    map[string]string `pulumi:"properties"`
	Target        *string           `pulumi:"target"`
}

type ComputeConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ComputeConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeConfigurationResponse)(nil)).Elem()
}

func (o ComputeConfigurationResponseOutput) ToComputeConfigurationResponseOutput() ComputeConfigurationResponseOutput {
	return o
}

func (o ComputeConfigurationResponseOutput) ToComputeConfigurationResponseOutputWithContext(ctx context.Context) ComputeConfigurationResponseOutput {
	return o
}

func (o ComputeConfigurationResponseOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ComputeConfigurationResponse) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

func (o ComputeConfigurationResponseOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeConfigurationResponse) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o ComputeConfigurationResponseOutput) IsLocal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComputeConfigurationResponse) *bool { return v.IsLocal }).(pulumi.BoolPtrOutput)
}

func (o ComputeConfigurationResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeConfigurationResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ComputeConfigurationResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ComputeConfigurationResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ComputeConfigurationResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeConfigurationResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ComputeConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ComputeConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeConfigurationResponse)(nil)).Elem()
}

func (o ComputeConfigurationResponsePtrOutput) ToComputeConfigurationResponsePtrOutput() ComputeConfigurationResponsePtrOutput {
	return o
}

func (o ComputeConfigurationResponsePtrOutput) ToComputeConfigurationResponsePtrOutputWithContext(ctx context.Context) ComputeConfigurationResponsePtrOutput {
	return o
}

func (o ComputeConfigurationResponsePtrOutput) Elem() ComputeConfigurationResponseOutput {
	return o.ApplyT(func(v *ComputeConfigurationResponse) ComputeConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ComputeConfigurationResponse
		return ret
	}).(ComputeConfigurationResponseOutput)
}

func (o ComputeConfigurationResponsePtrOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ComputeConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.InstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o ComputeConfigurationResponsePtrOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.InstanceType
	}).(pulumi.StringPtrOutput)
}

func (o ComputeConfigurationResponsePtrOutput) IsLocal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsLocal
	}).(pulumi.BoolPtrOutput)
}

func (o ComputeConfigurationResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ComputeConfigurationResponsePtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeConfigurationResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o ComputeConfigurationResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type ComputeInstance struct {
	ComputeLocation  *string                    `pulumi:"computeLocation"`
	ComputeType      string                     `pulumi:"computeType"`
	Description      *string                    `pulumi:"description"`
	DisableLocalAuth *bool                      `pulumi:"disableLocalAuth"`
	Properties       *ComputeInstanceProperties `pulumi:"properties"`
	ResourceId       *string                    `pulumi:"resourceId"`
}


func (val *ComputeInstance) Defaults() *ComputeInstance {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type ComputeInstanceApplicationResponse struct {
	DisplayName *string `pulumi:"displayName"`
	EndpointUri *string `pulumi:"endpointUri"`
}

type ComputeInstanceConnectivityEndpointsResponse struct {
	PrivateIpAddress string `pulumi:"privateIpAddress"`
	PublicIpAddress  string `pulumi:"publicIpAddress"`
}

type ComputeInstanceCreatedByResponse struct {
	UserId    string `pulumi:"userId"`
	UserName  string `pulumi:"userName"`
	UserOrgId string `pulumi:"userOrgId"`
}

type ComputeInstanceLastOperationResponse struct {
	OperationName   *string `pulumi:"operationName"`
	OperationStatus *string `pulumi:"operationStatus"`
	OperationTime   *string `pulumi:"operationTime"`
}

type ComputeInstanceProperties struct {
	ApplicationSharingPolicy         *string                          `pulumi:"applicationSharingPolicy"`
	ComputeInstanceAuthorizationType *string                          `pulumi:"computeInstanceAuthorizationType"`
	EnableNodePublicIp               *bool                            `pulumi:"enableNodePublicIp"`
	PersonalComputeInstanceSettings  *PersonalComputeInstanceSettings `pulumi:"personalComputeInstanceSettings"`
	Schedules                        *ComputeSchedules                `pulumi:"schedules"`
	SetupScripts                     *SetupScripts                    `pulumi:"setupScripts"`
	SshSettings                      *ComputeInstanceSshSettings      `pulumi:"sshSettings"`
	Subnet                           *ResourceId                      `pulumi:"subnet"`
	VmSize                           *string                          `pulumi:"vmSize"`
}


func (val *ComputeInstanceProperties) Defaults() *ComputeInstanceProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ApplicationSharingPolicy) {
		applicationSharingPolicy_ := "Shared"
		tmp.ApplicationSharingPolicy = &applicationSharingPolicy_
	}
	if isZero(tmp.ComputeInstanceAuthorizationType) {
		computeInstanceAuthorizationType_ := "personal"
		tmp.ComputeInstanceAuthorizationType = &computeInstanceAuthorizationType_
	}
	if isZero(tmp.EnableNodePublicIp) {
		enableNodePublicIp_ := true
		tmp.EnableNodePublicIp = &enableNodePublicIp_
	}
	tmp.SshSettings = tmp.SshSettings.Defaults()

	return &tmp
}

type ComputeInstanceResponse struct {
	ComputeLocation    *string                            `pulumi:"computeLocation"`
	ComputeType        string                             `pulumi:"computeType"`
	CreatedOn          string                             `pulumi:"createdOn"`
	Description        *string                            `pulumi:"description"`
	DisableLocalAuth   *bool                              `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                               `pulumi:"isAttachedCompute"`
	ModifiedOn         string                             `pulumi:"modifiedOn"`
	Properties         *ComputeInstanceResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse            `pulumi:"provisioningErrors"`
	ProvisioningState  string                             `pulumi:"provisioningState"`
	ResourceId         *string                            `pulumi:"resourceId"`
}


func (val *ComputeInstanceResponse) Defaults() *ComputeInstanceResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type ComputeInstanceResponseProperties struct {
	ApplicationSharingPolicy         *string                                      `pulumi:"applicationSharingPolicy"`
	Applications                     []ComputeInstanceApplicationResponse         `pulumi:"applications"`
	ComputeInstanceAuthorizationType *string                                      `pulumi:"computeInstanceAuthorizationType"`
	ConnectivityEndpoints            ComputeInstanceConnectivityEndpointsResponse `pulumi:"connectivityEndpoints"`
	CreatedBy                        ComputeInstanceCreatedByResponse             `pulumi:"createdBy"`
	EnableNodePublicIp               *bool                                        `pulumi:"enableNodePublicIp"`
	Errors                           []ErrorResponseResponse                      `pulumi:"errors"`
	LastOperation                    ComputeInstanceLastOperationResponse         `pulumi:"lastOperation"`
	PersonalComputeInstanceSettings  *PersonalComputeInstanceSettingsResponse     `pulumi:"personalComputeInstanceSettings"`
	Schedules                        *ComputeSchedulesResponse                    `pulumi:"schedules"`
	SetupScripts                     *SetupScriptsResponse                        `pulumi:"setupScripts"`
	SshSettings                      *ComputeInstanceSshSettingsResponse          `pulumi:"sshSettings"`
	State                            string                                       `pulumi:"state"`
	Subnet                           *ResourceIdResponse                          `pulumi:"subnet"`
	VmSize                           *string                                      `pulumi:"vmSize"`
}


func (val *ComputeInstanceResponseProperties) Defaults() *ComputeInstanceResponseProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ApplicationSharingPolicy) {
		applicationSharingPolicy_ := "Shared"
		tmp.ApplicationSharingPolicy = &applicationSharingPolicy_
	}
	if isZero(tmp.ComputeInstanceAuthorizationType) {
		computeInstanceAuthorizationType_ := "personal"
		tmp.ComputeInstanceAuthorizationType = &computeInstanceAuthorizationType_
	}
	if isZero(tmp.EnableNodePublicIp) {
		enableNodePublicIp_ := true
		tmp.EnableNodePublicIp = &enableNodePublicIp_
	}
	tmp.SshSettings = tmp.SshSettings.Defaults()

	return &tmp
}

type ComputeInstanceSshSettings struct {
	AdminPublicKey  *string `pulumi:"adminPublicKey"`
	SshPublicAccess *string `pulumi:"sshPublicAccess"`
}


func (val *ComputeInstanceSshSettings) Defaults() *ComputeInstanceSshSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SshPublicAccess) {
		sshPublicAccess_ := "Disabled"
		tmp.SshPublicAccess = &sshPublicAccess_
	}
	return &tmp
}

type ComputeInstanceSshSettingsResponse struct {
	AdminPublicKey  *string `pulumi:"adminPublicKey"`
	AdminUserName   string  `pulumi:"adminUserName"`
	SshPort         int     `pulumi:"sshPort"`
	SshPublicAccess *string `pulumi:"sshPublicAccess"`
}


func (val *ComputeInstanceSshSettingsResponse) Defaults() *ComputeInstanceSshSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SshPublicAccess) {
		sshPublicAccess_ := "Disabled"
		tmp.SshPublicAccess = &sshPublicAccess_
	}
	return &tmp
}

type ComputeSchedules struct {
	ComputeStartStop []ComputeStartStopSchedule `pulumi:"computeStartStop"`
}

type ComputeSchedulesResponse struct {
	ComputeStartStop []ComputeStartStopScheduleResponse `pulumi:"computeStartStop"`
}

type ComputeStartStopSchedule struct {
	Action      *string     `pulumi:"action"`
	Cron        *Cron       `pulumi:"cron"`
	Recurrence  *Recurrence `pulumi:"recurrence"`
	Status      *string     `pulumi:"status"`
	TriggerType *string     `pulumi:"triggerType"`
}

type ComputeStartStopScheduleResponse struct {
	Action             *string             `pulumi:"action"`
	Cron               *CronResponse       `pulumi:"cron"`
	Id                 string              `pulumi:"id"`
	ProvisioningStatus string              `pulumi:"provisioningStatus"`
	Recurrence         *RecurrenceResponse `pulumi:"recurrence"`
	Status             *string             `pulumi:"status"`
	TriggerType        *string             `pulumi:"triggerType"`
}

type ContainerResourceRequirements struct {
	Cpu             *float64 `pulumi:"cpu"`
	CpuLimit        *float64 `pulumi:"cpuLimit"`
	Fpga            *int     `pulumi:"fpga"`
	Gpu             *int     `pulumi:"gpu"`
	MemoryInGB      *float64 `pulumi:"memoryInGB"`
	MemoryInGBLimit *float64 `pulumi:"memoryInGBLimit"`
}

type ContainerResourceRequirementsResponse struct {
	Cpu             *float64 `pulumi:"cpu"`
	CpuLimit        *float64 `pulumi:"cpuLimit"`
	Fpga            *int     `pulumi:"fpga"`
	Gpu             *int     `pulumi:"gpu"`
	MemoryInGB      *float64 `pulumi:"memoryInGB"`
	MemoryInGBLimit *float64 `pulumi:"memoryInGBLimit"`
}

type CosmosDbSettings struct {
	CollectionsThroughput *int `pulumi:"collectionsThroughput"`
}





type CosmosDbSettingsInput interface {
	pulumi.Input

	ToCosmosDbSettingsOutput() CosmosDbSettingsOutput
	ToCosmosDbSettingsOutputWithContext(context.Context) CosmosDbSettingsOutput
}

type CosmosDbSettingsArgs struct {
	CollectionsThroughput pulumi.IntPtrInput `pulumi:"collectionsThroughput"`
}

func (CosmosDbSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CosmosDbSettings)(nil)).Elem()
}

func (i CosmosDbSettingsArgs) ToCosmosDbSettingsOutput() CosmosDbSettingsOutput {
	return i.ToCosmosDbSettingsOutputWithContext(context.Background())
}

func (i CosmosDbSettingsArgs) ToCosmosDbSettingsOutputWithContext(ctx context.Context) CosmosDbSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosmosDbSettingsOutput)
}

func (i CosmosDbSettingsArgs) ToCosmosDbSettingsPtrOutput() CosmosDbSettingsPtrOutput {
	return i.ToCosmosDbSettingsPtrOutputWithContext(context.Background())
}

func (i CosmosDbSettingsArgs) ToCosmosDbSettingsPtrOutputWithContext(ctx context.Context) CosmosDbSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosmosDbSettingsOutput).ToCosmosDbSettingsPtrOutputWithContext(ctx)
}









type CosmosDbSettingsPtrInput interface {
	pulumi.Input

	ToCosmosDbSettingsPtrOutput() CosmosDbSettingsPtrOutput
	ToCosmosDbSettingsPtrOutputWithContext(context.Context) CosmosDbSettingsPtrOutput
}

type cosmosDbSettingsPtrType CosmosDbSettingsArgs

func CosmosDbSettingsPtr(v *CosmosDbSettingsArgs) CosmosDbSettingsPtrInput {
	return (*cosmosDbSettingsPtrType)(v)
}

func (*cosmosDbSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbSettings)(nil)).Elem()
}

func (i *cosmosDbSettingsPtrType) ToCosmosDbSettingsPtrOutput() CosmosDbSettingsPtrOutput {
	return i.ToCosmosDbSettingsPtrOutputWithContext(context.Background())
}

func (i *cosmosDbSettingsPtrType) ToCosmosDbSettingsPtrOutputWithContext(ctx context.Context) CosmosDbSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosmosDbSettingsPtrOutput)
}

type CosmosDbSettingsOutput struct{ *pulumi.OutputState }

func (CosmosDbSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CosmosDbSettings)(nil)).Elem()
}

func (o CosmosDbSettingsOutput) ToCosmosDbSettingsOutput() CosmosDbSettingsOutput {
	return o
}

func (o CosmosDbSettingsOutput) ToCosmosDbSettingsOutputWithContext(ctx context.Context) CosmosDbSettingsOutput {
	return o
}

func (o CosmosDbSettingsOutput) ToCosmosDbSettingsPtrOutput() CosmosDbSettingsPtrOutput {
	return o.ToCosmosDbSettingsPtrOutputWithContext(context.Background())
}

func (o CosmosDbSettingsOutput) ToCosmosDbSettingsPtrOutputWithContext(ctx context.Context) CosmosDbSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CosmosDbSettings) *CosmosDbSettings {
		return &v
	}).(CosmosDbSettingsPtrOutput)
}

func (o CosmosDbSettingsOutput) CollectionsThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CosmosDbSettings) *int { return v.CollectionsThroughput }).(pulumi.IntPtrOutput)
}

type CosmosDbSettingsPtrOutput struct{ *pulumi.OutputState }

func (CosmosDbSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbSettings)(nil)).Elem()
}

func (o CosmosDbSettingsPtrOutput) ToCosmosDbSettingsPtrOutput() CosmosDbSettingsPtrOutput {
	return o
}

func (o CosmosDbSettingsPtrOutput) ToCosmosDbSettingsPtrOutputWithContext(ctx context.Context) CosmosDbSettingsPtrOutput {
	return o
}

func (o CosmosDbSettingsPtrOutput) Elem() CosmosDbSettingsOutput {
	return o.ApplyT(func(v *CosmosDbSettings) CosmosDbSettings {
		if v != nil {
			return *v
		}
		var ret CosmosDbSettings
		return ret
	}).(CosmosDbSettingsOutput)
}

func (o CosmosDbSettingsPtrOutput) CollectionsThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CosmosDbSettings) *int {
		if v == nil {
			return nil
		}
		return v.CollectionsThroughput
	}).(pulumi.IntPtrOutput)
}

type CosmosDbSettingsResponse struct {
	CollectionsThroughput *int `pulumi:"collectionsThroughput"`
}

type CosmosDbSettingsResponseOutput struct{ *pulumi.OutputState }

func (CosmosDbSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CosmosDbSettingsResponse)(nil)).Elem()
}

func (o CosmosDbSettingsResponseOutput) ToCosmosDbSettingsResponseOutput() CosmosDbSettingsResponseOutput {
	return o
}

func (o CosmosDbSettingsResponseOutput) ToCosmosDbSettingsResponseOutputWithContext(ctx context.Context) CosmosDbSettingsResponseOutput {
	return o
}

func (o CosmosDbSettingsResponseOutput) CollectionsThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CosmosDbSettingsResponse) *int { return v.CollectionsThroughput }).(pulumi.IntPtrOutput)
}

type CosmosDbSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CosmosDbSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbSettingsResponse)(nil)).Elem()
}

func (o CosmosDbSettingsResponsePtrOutput) ToCosmosDbSettingsResponsePtrOutput() CosmosDbSettingsResponsePtrOutput {
	return o
}

func (o CosmosDbSettingsResponsePtrOutput) ToCosmosDbSettingsResponsePtrOutputWithContext(ctx context.Context) CosmosDbSettingsResponsePtrOutput {
	return o
}

func (o CosmosDbSettingsResponsePtrOutput) Elem() CosmosDbSettingsResponseOutput {
	return o.ApplyT(func(v *CosmosDbSettingsResponse) CosmosDbSettingsResponse {
		if v != nil {
			return *v
		}
		var ret CosmosDbSettingsResponse
		return ret
	}).(CosmosDbSettingsResponseOutput)
}

func (o CosmosDbSettingsResponsePtrOutput) CollectionsThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CosmosDbSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.CollectionsThroughput
	}).(pulumi.IntPtrOutput)
}

type Cron struct {
	Expression *string `pulumi:"expression"`
	StartTime  *string `pulumi:"startTime"`
	TimeZone   *string `pulumi:"timeZone"`
}

type CronResponse struct {
	Expression *string `pulumi:"expression"`
	StartTime  *string `pulumi:"startTime"`
	TimeZone   *string `pulumi:"timeZone"`
}

type DataContainerType struct {
	Description *string           `pulumi:"description"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}





type DataContainerTypeInput interface {
	pulumi.Input

	ToDataContainerTypeOutput() DataContainerTypeOutput
	ToDataContainerTypeOutputWithContext(context.Context) DataContainerTypeOutput
}

type DataContainerTypeArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}

func (DataContainerTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataContainerType)(nil)).Elem()
}

func (i DataContainerTypeArgs) ToDataContainerTypeOutput() DataContainerTypeOutput {
	return i.ToDataContainerTypeOutputWithContext(context.Background())
}

func (i DataContainerTypeArgs) ToDataContainerTypeOutputWithContext(ctx context.Context) DataContainerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataContainerTypeOutput)
}

type DataContainerTypeOutput struct{ *pulumi.OutputState }

func (DataContainerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataContainerType)(nil)).Elem()
}

func (o DataContainerTypeOutput) ToDataContainerTypeOutput() DataContainerTypeOutput {
	return o
}

func (o DataContainerTypeOutput) ToDataContainerTypeOutputWithContext(ctx context.Context) DataContainerTypeOutput {
	return o
}

func (o DataContainerTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataContainerType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataContainerTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataContainerType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o DataContainerTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataContainerType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type DataContainerResponse struct {
	Description *string           `pulumi:"description"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}

type DataContainerResponseOutput struct{ *pulumi.OutputState }

func (DataContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataContainerResponse)(nil)).Elem()
}

func (o DataContainerResponseOutput) ToDataContainerResponseOutput() DataContainerResponseOutput {
	return o
}

func (o DataContainerResponseOutput) ToDataContainerResponseOutputWithContext(ctx context.Context) DataContainerResponseOutput {
	return o
}

func (o DataContainerResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataContainerResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataContainerResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataContainerResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o DataContainerResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataContainerResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type DataFactory struct {
	ComputeLocation  *string `pulumi:"computeLocation"`
	ComputeType      string  `pulumi:"computeType"`
	Description      *string `pulumi:"description"`
	DisableLocalAuth *bool   `pulumi:"disableLocalAuth"`
	ResourceId       *string `pulumi:"resourceId"`
}

type DataFactoryResponse struct {
	ComputeLocation    *string                 `pulumi:"computeLocation"`
	ComputeType        string                  `pulumi:"computeType"`
	CreatedOn          string                  `pulumi:"createdOn"`
	Description        *string                 `pulumi:"description"`
	DisableLocalAuth   *bool                   `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                    `pulumi:"isAttachedCompute"`
	ModifiedOn         string                  `pulumi:"modifiedOn"`
	ProvisioningErrors []ErrorResponseResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                  `pulumi:"provisioningState"`
	ResourceId         *string                 `pulumi:"resourceId"`
}

type DataLakeAnalytics struct {
	ComputeLocation  *string                      `pulumi:"computeLocation"`
	ComputeType      string                       `pulumi:"computeType"`
	Description      *string                      `pulumi:"description"`
	DisableLocalAuth *bool                        `pulumi:"disableLocalAuth"`
	Properties       *DataLakeAnalyticsProperties `pulumi:"properties"`
	ResourceId       *string                      `pulumi:"resourceId"`
}

type DataLakeAnalyticsProperties struct {
	DataLakeStoreAccountName *string `pulumi:"dataLakeStoreAccountName"`
}

type DataLakeAnalyticsResponse struct {
	ComputeLocation    *string                              `pulumi:"computeLocation"`
	ComputeType        string                               `pulumi:"computeType"`
	CreatedOn          string                               `pulumi:"createdOn"`
	Description        *string                              `pulumi:"description"`
	DisableLocalAuth   *bool                                `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                                 `pulumi:"isAttachedCompute"`
	ModifiedOn         string                               `pulumi:"modifiedOn"`
	Properties         *DataLakeAnalyticsResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse              `pulumi:"provisioningErrors"`
	ProvisioningState  string                               `pulumi:"provisioningState"`
	ResourceId         *string                              `pulumi:"resourceId"`
}

type DataLakeAnalyticsResponseProperties struct {
	DataLakeStoreAccountName *string `pulumi:"dataLakeStoreAccountName"`
}

type DataPathAssetReference struct {
	DatastoreId   *string `pulumi:"datastoreId"`
	Path          *string `pulumi:"path"`
	ReferenceType string  `pulumi:"referenceType"`
}

type DataPathAssetReferenceResponse struct {
	DatastoreId   *string `pulumi:"datastoreId"`
	Path          *string `pulumi:"path"`
	ReferenceType string  `pulumi:"referenceType"`
}

type DataVersionType struct {
	DatasetType *string           `pulumi:"datasetType"`
	DatastoreId *string           `pulumi:"datastoreId"`
	Description *string           `pulumi:"description"`
	IsAnonymous *bool             `pulumi:"isAnonymous"`
	Path        string            `pulumi:"path"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}





type DataVersionTypeInput interface {
	pulumi.Input

	ToDataVersionTypeOutput() DataVersionTypeOutput
	ToDataVersionTypeOutputWithContext(context.Context) DataVersionTypeOutput
}

type DataVersionTypeArgs struct {
	DatasetType pulumi.StringPtrInput `pulumi:"datasetType"`
	DatastoreId pulumi.StringPtrInput `pulumi:"datastoreId"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	IsAnonymous pulumi.BoolPtrInput   `pulumi:"isAnonymous"`
	Path        pulumi.StringInput    `pulumi:"path"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}

func (DataVersionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataVersionType)(nil)).Elem()
}

func (i DataVersionTypeArgs) ToDataVersionTypeOutput() DataVersionTypeOutput {
	return i.ToDataVersionTypeOutputWithContext(context.Background())
}

func (i DataVersionTypeArgs) ToDataVersionTypeOutputWithContext(ctx context.Context) DataVersionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataVersionTypeOutput)
}

type DataVersionTypeOutput struct{ *pulumi.OutputState }

func (DataVersionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataVersionType)(nil)).Elem()
}

func (o DataVersionTypeOutput) ToDataVersionTypeOutput() DataVersionTypeOutput {
	return o
}

func (o DataVersionTypeOutput) ToDataVersionTypeOutputWithContext(ctx context.Context) DataVersionTypeOutput {
	return o
}

func (o DataVersionTypeOutput) DatasetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataVersionType) *string { return v.DatasetType }).(pulumi.StringPtrOutput)
}

func (o DataVersionTypeOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataVersionType) *string { return v.DatastoreId }).(pulumi.StringPtrOutput)
}

func (o DataVersionTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataVersionType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataVersionTypeOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataVersionType) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o DataVersionTypeOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v DataVersionType) string { return v.Path }).(pulumi.StringOutput)
}

func (o DataVersionTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataVersionType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o DataVersionTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataVersionType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type DataVersionResponse struct {
	DatasetType *string           `pulumi:"datasetType"`
	DatastoreId *string           `pulumi:"datastoreId"`
	Description *string           `pulumi:"description"`
	IsAnonymous *bool             `pulumi:"isAnonymous"`
	Path        string            `pulumi:"path"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}

type DataVersionResponseOutput struct{ *pulumi.OutputState }

func (DataVersionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataVersionResponse)(nil)).Elem()
}

func (o DataVersionResponseOutput) ToDataVersionResponseOutput() DataVersionResponseOutput {
	return o
}

func (o DataVersionResponseOutput) ToDataVersionResponseOutputWithContext(ctx context.Context) DataVersionResponseOutput {
	return o
}

func (o DataVersionResponseOutput) DatasetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataVersionResponse) *string { return v.DatasetType }).(pulumi.StringPtrOutput)
}

func (o DataVersionResponseOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataVersionResponse) *string { return v.DatastoreId }).(pulumi.StringPtrOutput)
}

func (o DataVersionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataVersionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataVersionResponseOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataVersionResponse) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o DataVersionResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v DataVersionResponse) string { return v.Path }).(pulumi.StringOutput)
}

func (o DataVersionResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataVersionResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o DataVersionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataVersionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type Databricks struct {
	ComputeLocation  *string               `pulumi:"computeLocation"`
	ComputeType      string                `pulumi:"computeType"`
	Description      *string               `pulumi:"description"`
	DisableLocalAuth *bool                 `pulumi:"disableLocalAuth"`
	Properties       *DatabricksProperties `pulumi:"properties"`
	ResourceId       *string               `pulumi:"resourceId"`
}

type DatabricksProperties struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
	WorkspaceUrl          *string `pulumi:"workspaceUrl"`
}

type DatabricksResponse struct {
	ComputeLocation    *string                       `pulumi:"computeLocation"`
	ComputeType        string                        `pulumi:"computeType"`
	CreatedOn          string                        `pulumi:"createdOn"`
	Description        *string                       `pulumi:"description"`
	DisableLocalAuth   *bool                         `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                          `pulumi:"isAttachedCompute"`
	ModifiedOn         string                        `pulumi:"modifiedOn"`
	Properties         *DatabricksResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse       `pulumi:"provisioningErrors"`
	ProvisioningState  string                        `pulumi:"provisioningState"`
	ResourceId         *string                       `pulumi:"resourceId"`
}

type DatabricksResponseProperties struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
	WorkspaceUrl          *string `pulumi:"workspaceUrl"`
}

type DatastoreProperties struct {
	Contents    interface{}       `pulumi:"contents"`
	Description *string           `pulumi:"description"`
	IsDefault   *bool             `pulumi:"isDefault"`
	LinkedInfo  *LinkedInfo       `pulumi:"linkedInfo"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}





type DatastorePropertiesInput interface {
	pulumi.Input

	ToDatastorePropertiesOutput() DatastorePropertiesOutput
	ToDatastorePropertiesOutputWithContext(context.Context) DatastorePropertiesOutput
}

type DatastorePropertiesArgs struct {
	Contents    pulumi.Input          `pulumi:"contents"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	IsDefault   pulumi.BoolPtrInput   `pulumi:"isDefault"`
	LinkedInfo  LinkedInfoPtrInput    `pulumi:"linkedInfo"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}

func (DatastorePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatastoreProperties)(nil)).Elem()
}

func (i DatastorePropertiesArgs) ToDatastorePropertiesOutput() DatastorePropertiesOutput {
	return i.ToDatastorePropertiesOutputWithContext(context.Background())
}

func (i DatastorePropertiesArgs) ToDatastorePropertiesOutputWithContext(ctx context.Context) DatastorePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastorePropertiesOutput)
}

type DatastorePropertiesOutput struct{ *pulumi.OutputState }

func (DatastorePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatastoreProperties)(nil)).Elem()
}

func (o DatastorePropertiesOutput) ToDatastorePropertiesOutput() DatastorePropertiesOutput {
	return o
}

func (o DatastorePropertiesOutput) ToDatastorePropertiesOutputWithContext(ctx context.Context) DatastorePropertiesOutput {
	return o
}

func (o DatastorePropertiesOutput) Contents() pulumi.AnyOutput {
	return o.ApplyT(func(v DatastoreProperties) interface{} { return v.Contents }).(pulumi.AnyOutput)
}

func (o DatastorePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatastoreProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DatastorePropertiesOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DatastoreProperties) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o DatastorePropertiesOutput) LinkedInfo() LinkedInfoPtrOutput {
	return o.ApplyT(func(v DatastoreProperties) *LinkedInfo { return v.LinkedInfo }).(LinkedInfoPtrOutput)
}

func (o DatastorePropertiesOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v DatastoreProperties) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o DatastorePropertiesOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DatastoreProperties) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type DatastorePropertiesResponse struct {
	Contents         interface{}         `pulumi:"contents"`
	Description      *string             `pulumi:"description"`
	HasBeenValidated bool                `pulumi:"hasBeenValidated"`
	IsDefault        *bool               `pulumi:"isDefault"`
	LinkedInfo       *LinkedInfoResponse `pulumi:"linkedInfo"`
	Properties       map[string]string   `pulumi:"properties"`
	Tags             map[string]string   `pulumi:"tags"`
}

type DatastorePropertiesResponseOutput struct{ *pulumi.OutputState }

func (DatastorePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatastorePropertiesResponse)(nil)).Elem()
}

func (o DatastorePropertiesResponseOutput) ToDatastorePropertiesResponseOutput() DatastorePropertiesResponseOutput {
	return o
}

func (o DatastorePropertiesResponseOutput) ToDatastorePropertiesResponseOutputWithContext(ctx context.Context) DatastorePropertiesResponseOutput {
	return o
}

func (o DatastorePropertiesResponseOutput) Contents() pulumi.AnyOutput {
	return o.ApplyT(func(v DatastorePropertiesResponse) interface{} { return v.Contents }).(pulumi.AnyOutput)
}

func (o DatastorePropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatastorePropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DatastorePropertiesResponseOutput) HasBeenValidated() pulumi.BoolOutput {
	return o.ApplyT(func(v DatastorePropertiesResponse) bool { return v.HasBeenValidated }).(pulumi.BoolOutput)
}

func (o DatastorePropertiesResponseOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DatastorePropertiesResponse) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o DatastorePropertiesResponseOutput) LinkedInfo() LinkedInfoResponsePtrOutput {
	return o.ApplyT(func(v DatastorePropertiesResponse) *LinkedInfoResponse { return v.LinkedInfo }).(LinkedInfoResponsePtrOutput)
}

func (o DatastorePropertiesResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v DatastorePropertiesResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o DatastorePropertiesResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DatastorePropertiesResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type DockerBuild struct {
	Context                 *string              `pulumi:"context"`
	DockerSpecificationType string               `pulumi:"dockerSpecificationType"`
	Dockerfile              string               `pulumi:"dockerfile"`
	Platform                *DockerImagePlatform `pulumi:"platform"`
}

type DockerBuildResponse struct {
	Context                 *string                      `pulumi:"context"`
	DockerSpecificationType string                       `pulumi:"dockerSpecificationType"`
	Dockerfile              string                       `pulumi:"dockerfile"`
	Platform                *DockerImagePlatformResponse `pulumi:"platform"`
}

type DockerImage struct {
	DockerImageUri          string               `pulumi:"dockerImageUri"`
	DockerSpecificationType string               `pulumi:"dockerSpecificationType"`
	Platform                *DockerImagePlatform `pulumi:"platform"`
}

type DockerImagePlatform struct {
	OperatingSystemType *string `pulumi:"operatingSystemType"`
}

type DockerImagePlatformResponse struct {
	OperatingSystemType *string `pulumi:"operatingSystemType"`
}

type DockerImageResponse struct {
	DockerImageUri          string                       `pulumi:"dockerImageUri"`
	DockerSpecificationType string                       `pulumi:"dockerSpecificationType"`
	Platform                *DockerImagePlatformResponse `pulumi:"platform"`
}

type EncryptionProperty struct {
	Identity           *IdentityForCmk    `pulumi:"identity"`
	KeyVaultProperties KeyVaultProperties `pulumi:"keyVaultProperties"`
	Status             string             `pulumi:"status"`
}





type EncryptionPropertyInput interface {
	pulumi.Input

	ToEncryptionPropertyOutput() EncryptionPropertyOutput
	ToEncryptionPropertyOutputWithContext(context.Context) EncryptionPropertyOutput
}

type EncryptionPropertyArgs struct {
	Identity           IdentityForCmkPtrInput  `pulumi:"identity"`
	KeyVaultProperties KeyVaultPropertiesInput `pulumi:"keyVaultProperties"`
	Status             pulumi.StringInput      `pulumi:"status"`
}

func (EncryptionPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperty)(nil)).Elem()
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyOutput() EncryptionPropertyOutput {
	return i.ToEncryptionPropertyOutputWithContext(context.Background())
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyOutputWithContext(ctx context.Context) EncryptionPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyOutput)
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return i.ToEncryptionPropertyPtrOutputWithContext(context.Background())
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyOutput).ToEncryptionPropertyPtrOutputWithContext(ctx)
}









type EncryptionPropertyPtrInput interface {
	pulumi.Input

	ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput
	ToEncryptionPropertyPtrOutputWithContext(context.Context) EncryptionPropertyPtrOutput
}

type encryptionPropertyPtrType EncryptionPropertyArgs

func EncryptionPropertyPtr(v *EncryptionPropertyArgs) EncryptionPropertyPtrInput {
	return (*encryptionPropertyPtrType)(v)
}

func (*encryptionPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperty)(nil)).Elem()
}

func (i *encryptionPropertyPtrType) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return i.ToEncryptionPropertyPtrOutputWithContext(context.Background())
}

func (i *encryptionPropertyPtrType) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyPtrOutput)
}

type EncryptionPropertyOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperty)(nil)).Elem()
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyOutput() EncryptionPropertyOutput {
	return o
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyOutputWithContext(ctx context.Context) EncryptionPropertyOutput {
	return o
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return o.ToEncryptionPropertyPtrOutputWithContext(context.Background())
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionProperty) *EncryptionProperty {
		return &v
	}).(EncryptionPropertyPtrOutput)
}

func (o EncryptionPropertyOutput) Identity() IdentityForCmkPtrOutput {
	return o.ApplyT(func(v EncryptionProperty) *IdentityForCmk { return v.Identity }).(IdentityForCmkPtrOutput)
}

func (o EncryptionPropertyOutput) KeyVaultProperties() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v EncryptionProperty) KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesOutput)
}

func (o EncryptionPropertyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionProperty) string { return v.Status }).(pulumi.StringOutput)
}

type EncryptionPropertyPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperty)(nil)).Elem()
}

func (o EncryptionPropertyPtrOutput) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return o
}

func (o EncryptionPropertyPtrOutput) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return o
}

func (o EncryptionPropertyPtrOutput) Elem() EncryptionPropertyOutput {
	return o.ApplyT(func(v *EncryptionProperty) EncryptionProperty {
		if v != nil {
			return *v
		}
		var ret EncryptionProperty
		return ret
	}).(EncryptionPropertyOutput)
}

func (o EncryptionPropertyPtrOutput) Identity() IdentityForCmkPtrOutput {
	return o.ApplyT(func(v *EncryptionProperty) *IdentityForCmk {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(IdentityForCmkPtrOutput)
}

func (o EncryptionPropertyPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *EncryptionProperty) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return &v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

func (o EncryptionPropertyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type EncryptionPropertyResponse struct {
	Identity           *IdentityForCmkResponse    `pulumi:"identity"`
	KeyVaultProperties KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Status             string                     `pulumi:"status"`
}

type EncryptionPropertyResponseOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertyResponse)(nil)).Elem()
}

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponseOutput() EncryptionPropertyResponseOutput {
	return o
}

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponseOutputWithContext(ctx context.Context) EncryptionPropertyResponseOutput {
	return o
}

func (o EncryptionPropertyResponseOutput) Identity() IdentityForCmkResponsePtrOutput {
	return o.ApplyT(func(v EncryptionPropertyResponse) *IdentityForCmkResponse { return v.Identity }).(IdentityForCmkResponsePtrOutput)
}

func (o EncryptionPropertyResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v EncryptionPropertyResponse) KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponseOutput)
}

func (o EncryptionPropertyResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionPropertyResponse) string { return v.Status }).(pulumi.StringOutput)
}

type EncryptionPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertyResponse)(nil)).Elem()
}

func (o EncryptionPropertyResponsePtrOutput) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return o
}

func (o EncryptionPropertyResponsePtrOutput) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return o
}

func (o EncryptionPropertyResponsePtrOutput) Elem() EncryptionPropertyResponseOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) EncryptionPropertyResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertyResponse
		return ret
	}).(EncryptionPropertyResponseOutput)
}

func (o EncryptionPropertyResponsePtrOutput) Identity() IdentityForCmkResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) *IdentityForCmkResponse {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(IdentityForCmkResponsePtrOutput)
}

func (o EncryptionPropertyResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionPropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type EndpointAuthKeys struct {
	PrimaryKey   *string `pulumi:"primaryKey"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}





type EndpointAuthKeysInput interface {
	pulumi.Input

	ToEndpointAuthKeysOutput() EndpointAuthKeysOutput
	ToEndpointAuthKeysOutputWithContext(context.Context) EndpointAuthKeysOutput
}

type EndpointAuthKeysArgs struct {
	PrimaryKey   pulumi.StringPtrInput `pulumi:"primaryKey"`
	SecondaryKey pulumi.StringPtrInput `pulumi:"secondaryKey"`
}

func (EndpointAuthKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAuthKeys)(nil)).Elem()
}

func (i EndpointAuthKeysArgs) ToEndpointAuthKeysOutput() EndpointAuthKeysOutput {
	return i.ToEndpointAuthKeysOutputWithContext(context.Background())
}

func (i EndpointAuthKeysArgs) ToEndpointAuthKeysOutputWithContext(ctx context.Context) EndpointAuthKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAuthKeysOutput)
}

func (i EndpointAuthKeysArgs) ToEndpointAuthKeysPtrOutput() EndpointAuthKeysPtrOutput {
	return i.ToEndpointAuthKeysPtrOutputWithContext(context.Background())
}

func (i EndpointAuthKeysArgs) ToEndpointAuthKeysPtrOutputWithContext(ctx context.Context) EndpointAuthKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAuthKeysOutput).ToEndpointAuthKeysPtrOutputWithContext(ctx)
}









type EndpointAuthKeysPtrInput interface {
	pulumi.Input

	ToEndpointAuthKeysPtrOutput() EndpointAuthKeysPtrOutput
	ToEndpointAuthKeysPtrOutputWithContext(context.Context) EndpointAuthKeysPtrOutput
}

type endpointAuthKeysPtrType EndpointAuthKeysArgs

func EndpointAuthKeysPtr(v *EndpointAuthKeysArgs) EndpointAuthKeysPtrInput {
	return (*endpointAuthKeysPtrType)(v)
}

func (*endpointAuthKeysPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAuthKeys)(nil)).Elem()
}

func (i *endpointAuthKeysPtrType) ToEndpointAuthKeysPtrOutput() EndpointAuthKeysPtrOutput {
	return i.ToEndpointAuthKeysPtrOutputWithContext(context.Background())
}

func (i *endpointAuthKeysPtrType) ToEndpointAuthKeysPtrOutputWithContext(ctx context.Context) EndpointAuthKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAuthKeysPtrOutput)
}

type EndpointAuthKeysOutput struct{ *pulumi.OutputState }

func (EndpointAuthKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAuthKeys)(nil)).Elem()
}

func (o EndpointAuthKeysOutput) ToEndpointAuthKeysOutput() EndpointAuthKeysOutput {
	return o
}

func (o EndpointAuthKeysOutput) ToEndpointAuthKeysOutputWithContext(ctx context.Context) EndpointAuthKeysOutput {
	return o
}

func (o EndpointAuthKeysOutput) ToEndpointAuthKeysPtrOutput() EndpointAuthKeysPtrOutput {
	return o.ToEndpointAuthKeysPtrOutputWithContext(context.Background())
}

func (o EndpointAuthKeysOutput) ToEndpointAuthKeysPtrOutputWithContext(ctx context.Context) EndpointAuthKeysPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointAuthKeys) *EndpointAuthKeys {
		return &v
	}).(EndpointAuthKeysPtrOutput)
}

func (o EndpointAuthKeysOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointAuthKeys) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o EndpointAuthKeysOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointAuthKeys) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

type EndpointAuthKeysPtrOutput struct{ *pulumi.OutputState }

func (EndpointAuthKeysPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAuthKeys)(nil)).Elem()
}

func (o EndpointAuthKeysPtrOutput) ToEndpointAuthKeysPtrOutput() EndpointAuthKeysPtrOutput {
	return o
}

func (o EndpointAuthKeysPtrOutput) ToEndpointAuthKeysPtrOutputWithContext(ctx context.Context) EndpointAuthKeysPtrOutput {
	return o
}

func (o EndpointAuthKeysPtrOutput) Elem() EndpointAuthKeysOutput {
	return o.ApplyT(func(v *EndpointAuthKeys) EndpointAuthKeys {
		if v != nil {
			return *v
		}
		var ret EndpointAuthKeys
		return ret
	}).(EndpointAuthKeysOutput)
}

func (o EndpointAuthKeysPtrOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointAuthKeys) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryKey
	}).(pulumi.StringPtrOutput)
}

func (o EndpointAuthKeysPtrOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointAuthKeys) *string {
		if v == nil {
			return nil
		}
		return v.SecondaryKey
	}).(pulumi.StringPtrOutput)
}

type EnvironmentContainerType struct {
	Description *string           `pulumi:"description"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}





type EnvironmentContainerTypeInput interface {
	pulumi.Input

	ToEnvironmentContainerTypeOutput() EnvironmentContainerTypeOutput
	ToEnvironmentContainerTypeOutputWithContext(context.Context) EnvironmentContainerTypeOutput
}

type EnvironmentContainerTypeArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}

func (EnvironmentContainerTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentContainerType)(nil)).Elem()
}

func (i EnvironmentContainerTypeArgs) ToEnvironmentContainerTypeOutput() EnvironmentContainerTypeOutput {
	return i.ToEnvironmentContainerTypeOutputWithContext(context.Background())
}

func (i EnvironmentContainerTypeArgs) ToEnvironmentContainerTypeOutputWithContext(ctx context.Context) EnvironmentContainerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentContainerTypeOutput)
}

type EnvironmentContainerTypeOutput struct{ *pulumi.OutputState }

func (EnvironmentContainerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentContainerType)(nil)).Elem()
}

func (o EnvironmentContainerTypeOutput) ToEnvironmentContainerTypeOutput() EnvironmentContainerTypeOutput {
	return o
}

func (o EnvironmentContainerTypeOutput) ToEnvironmentContainerTypeOutputWithContext(ctx context.Context) EnvironmentContainerTypeOutput {
	return o
}

func (o EnvironmentContainerTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentContainerType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnvironmentContainerTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentContainerType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o EnvironmentContainerTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentContainerType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type EnvironmentContainerResponse struct {
	Description *string           `pulumi:"description"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}

type EnvironmentContainerResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentContainerResponse)(nil)).Elem()
}

func (o EnvironmentContainerResponseOutput) ToEnvironmentContainerResponseOutput() EnvironmentContainerResponseOutput {
	return o
}

func (o EnvironmentContainerResponseOutput) ToEnvironmentContainerResponseOutputWithContext(ctx context.Context) EnvironmentContainerResponseOutput {
	return o
}

func (o EnvironmentContainerResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentContainerResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnvironmentContainerResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentContainerResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o EnvironmentContainerResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentContainerResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type EnvironmentSpecificationVersionType struct {
	CondaFile                    *string                       `pulumi:"condaFile"`
	Description                  *string                       `pulumi:"description"`
	Docker                       interface{}                   `pulumi:"docker"`
	InferenceContainerProperties *InferenceContainerProperties `pulumi:"inferenceContainerProperties"`
	IsAnonymous                  *bool                         `pulumi:"isAnonymous"`
	Properties                   map[string]string             `pulumi:"properties"`
	Tags                         map[string]string             `pulumi:"tags"`
}





type EnvironmentSpecificationVersionTypeInput interface {
	pulumi.Input

	ToEnvironmentSpecificationVersionTypeOutput() EnvironmentSpecificationVersionTypeOutput
	ToEnvironmentSpecificationVersionTypeOutputWithContext(context.Context) EnvironmentSpecificationVersionTypeOutput
}

type EnvironmentSpecificationVersionTypeArgs struct {
	CondaFile                    pulumi.StringPtrInput                `pulumi:"condaFile"`
	Description                  pulumi.StringPtrInput                `pulumi:"description"`
	Docker                       pulumi.Input                         `pulumi:"docker"`
	InferenceContainerProperties InferenceContainerPropertiesPtrInput `pulumi:"inferenceContainerProperties"`
	IsAnonymous                  pulumi.BoolPtrInput                  `pulumi:"isAnonymous"`
	Properties                   pulumi.StringMapInput                `pulumi:"properties"`
	Tags                         pulumi.StringMapInput                `pulumi:"tags"`
}

func (EnvironmentSpecificationVersionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSpecificationVersionType)(nil)).Elem()
}

func (i EnvironmentSpecificationVersionTypeArgs) ToEnvironmentSpecificationVersionTypeOutput() EnvironmentSpecificationVersionTypeOutput {
	return i.ToEnvironmentSpecificationVersionTypeOutputWithContext(context.Background())
}

func (i EnvironmentSpecificationVersionTypeArgs) ToEnvironmentSpecificationVersionTypeOutputWithContext(ctx context.Context) EnvironmentSpecificationVersionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSpecificationVersionTypeOutput)
}

type EnvironmentSpecificationVersionTypeOutput struct{ *pulumi.OutputState }

func (EnvironmentSpecificationVersionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSpecificationVersionType)(nil)).Elem()
}

func (o EnvironmentSpecificationVersionTypeOutput) ToEnvironmentSpecificationVersionTypeOutput() EnvironmentSpecificationVersionTypeOutput {
	return o
}

func (o EnvironmentSpecificationVersionTypeOutput) ToEnvironmentSpecificationVersionTypeOutputWithContext(ctx context.Context) EnvironmentSpecificationVersionTypeOutput {
	return o
}

func (o EnvironmentSpecificationVersionTypeOutput) CondaFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionType) *string { return v.CondaFile }).(pulumi.StringPtrOutput)
}

func (o EnvironmentSpecificationVersionTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnvironmentSpecificationVersionTypeOutput) Docker() pulumi.AnyOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionType) interface{} { return v.Docker }).(pulumi.AnyOutput)
}

func (o EnvironmentSpecificationVersionTypeOutput) InferenceContainerProperties() InferenceContainerPropertiesPtrOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionType) *InferenceContainerProperties {
		return v.InferenceContainerProperties
	}).(InferenceContainerPropertiesPtrOutput)
}

func (o EnvironmentSpecificationVersionTypeOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionType) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o EnvironmentSpecificationVersionTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o EnvironmentSpecificationVersionTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type EnvironmentSpecificationVersionResponse struct {
	CondaFile                    *string                               `pulumi:"condaFile"`
	Description                  *string                               `pulumi:"description"`
	Docker                       interface{}                           `pulumi:"docker"`
	EnvironmentSpecificationType string                                `pulumi:"environmentSpecificationType"`
	InferenceContainerProperties *InferenceContainerPropertiesResponse `pulumi:"inferenceContainerProperties"`
	IsAnonymous                  *bool                                 `pulumi:"isAnonymous"`
	Properties                   map[string]string                     `pulumi:"properties"`
	Tags                         map[string]string                     `pulumi:"tags"`
}

type EnvironmentSpecificationVersionResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentSpecificationVersionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSpecificationVersionResponse)(nil)).Elem()
}

func (o EnvironmentSpecificationVersionResponseOutput) ToEnvironmentSpecificationVersionResponseOutput() EnvironmentSpecificationVersionResponseOutput {
	return o
}

func (o EnvironmentSpecificationVersionResponseOutput) ToEnvironmentSpecificationVersionResponseOutputWithContext(ctx context.Context) EnvironmentSpecificationVersionResponseOutput {
	return o
}

func (o EnvironmentSpecificationVersionResponseOutput) CondaFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionResponse) *string { return v.CondaFile }).(pulumi.StringPtrOutput)
}

func (o EnvironmentSpecificationVersionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnvironmentSpecificationVersionResponseOutput) Docker() pulumi.AnyOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionResponse) interface{} { return v.Docker }).(pulumi.AnyOutput)
}

func (o EnvironmentSpecificationVersionResponseOutput) EnvironmentSpecificationType() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionResponse) string { return v.EnvironmentSpecificationType }).(pulumi.StringOutput)
}

func (o EnvironmentSpecificationVersionResponseOutput) InferenceContainerProperties() InferenceContainerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionResponse) *InferenceContainerPropertiesResponse {
		return v.InferenceContainerProperties
	}).(InferenceContainerPropertiesResponsePtrOutput)
}

func (o EnvironmentSpecificationVersionResponseOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionResponse) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o EnvironmentSpecificationVersionResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o EnvironmentSpecificationVersionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentSpecificationVersionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}

type ErrorDetailResponse struct {
	AdditionalInfo []ErrorAdditionalInfoResponse `pulumi:"additionalInfo"`
	Code           string                        `pulumi:"code"`
	Details        []ErrorDetailResponse         `pulumi:"details"`
	Message        string                        `pulumi:"message"`
	Target         string                        `pulumi:"target"`
}

type ErrorResponseResponse struct {
	Error *ErrorDetailResponse `pulumi:"error"`
}

type FlavorData struct {
	Data map[string]string `pulumi:"data"`
}





type FlavorDataInput interface {
	pulumi.Input

	ToFlavorDataOutput() FlavorDataOutput
	ToFlavorDataOutputWithContext(context.Context) FlavorDataOutput
}

type FlavorDataArgs struct {
	Data pulumi.StringMapInput `pulumi:"data"`
}

func (FlavorDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FlavorData)(nil)).Elem()
}

func (i FlavorDataArgs) ToFlavorDataOutput() FlavorDataOutput {
	return i.ToFlavorDataOutputWithContext(context.Background())
}

func (i FlavorDataArgs) ToFlavorDataOutputWithContext(ctx context.Context) FlavorDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlavorDataOutput)
}





type FlavorDataMapInput interface {
	pulumi.Input

	ToFlavorDataMapOutput() FlavorDataMapOutput
	ToFlavorDataMapOutputWithContext(context.Context) FlavorDataMapOutput
}

type FlavorDataMap map[string]FlavorDataInput

func (FlavorDataMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FlavorData)(nil)).Elem()
}

func (i FlavorDataMap) ToFlavorDataMapOutput() FlavorDataMapOutput {
	return i.ToFlavorDataMapOutputWithContext(context.Background())
}

func (i FlavorDataMap) ToFlavorDataMapOutputWithContext(ctx context.Context) FlavorDataMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlavorDataMapOutput)
}

type FlavorDataOutput struct{ *pulumi.OutputState }

func (FlavorDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlavorData)(nil)).Elem()
}

func (o FlavorDataOutput) ToFlavorDataOutput() FlavorDataOutput {
	return o
}

func (o FlavorDataOutput) ToFlavorDataOutputWithContext(ctx context.Context) FlavorDataOutput {
	return o
}

func (o FlavorDataOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v FlavorData) map[string]string { return v.Data }).(pulumi.StringMapOutput)
}

type FlavorDataMapOutput struct{ *pulumi.OutputState }

func (FlavorDataMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FlavorData)(nil)).Elem()
}

func (o FlavorDataMapOutput) ToFlavorDataMapOutput() FlavorDataMapOutput {
	return o
}

func (o FlavorDataMapOutput) ToFlavorDataMapOutputWithContext(ctx context.Context) FlavorDataMapOutput {
	return o
}

func (o FlavorDataMapOutput) MapIndex(k pulumi.StringInput) FlavorDataOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FlavorData {
		return vs[0].(map[string]FlavorData)[vs[1].(string)]
	}).(FlavorDataOutput)
}

type FlavorDataResponse struct {
	Data map[string]string `pulumi:"data"`
}

type FlavorDataResponseOutput struct{ *pulumi.OutputState }

func (FlavorDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlavorDataResponse)(nil)).Elem()
}

func (o FlavorDataResponseOutput) ToFlavorDataResponseOutput() FlavorDataResponseOutput {
	return o
}

func (o FlavorDataResponseOutput) ToFlavorDataResponseOutputWithContext(ctx context.Context) FlavorDataResponseOutput {
	return o
}

func (o FlavorDataResponseOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v FlavorDataResponse) map[string]string { return v.Data }).(pulumi.StringMapOutput)
}

type FlavorDataResponseMapOutput struct{ *pulumi.OutputState }

func (FlavorDataResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FlavorDataResponse)(nil)).Elem()
}

func (o FlavorDataResponseMapOutput) ToFlavorDataResponseMapOutput() FlavorDataResponseMapOutput {
	return o
}

func (o FlavorDataResponseMapOutput) ToFlavorDataResponseMapOutputWithContext(ctx context.Context) FlavorDataResponseMapOutput {
	return o
}

func (o FlavorDataResponseMapOutput) MapIndex(k pulumi.StringInput) FlavorDataResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FlavorDataResponse {
		return vs[0].(map[string]FlavorDataResponse)[vs[1].(string)]
	}).(FlavorDataResponseOutput)
}

type GlusterFsContents struct {
	ContentsType  string `pulumi:"contentsType"`
	ServerAddress string `pulumi:"serverAddress"`
	VolumeName    string `pulumi:"volumeName"`
}

type GlusterFsContentsResponse struct {
	ContentsType  string `pulumi:"contentsType"`
	ServerAddress string `pulumi:"serverAddress"`
	VolumeName    string `pulumi:"volumeName"`
}

type HDInsight struct {
	ComputeLocation  *string              `pulumi:"computeLocation"`
	ComputeType      string               `pulumi:"computeType"`
	Description      *string              `pulumi:"description"`
	DisableLocalAuth *bool                `pulumi:"disableLocalAuth"`
	Properties       *HDInsightProperties `pulumi:"properties"`
	ResourceId       *string              `pulumi:"resourceId"`
}

type HDInsightProperties struct {
	Address              *string                       `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	SshPort              *int                          `pulumi:"sshPort"`
}

type HDInsightResponse struct {
	ComputeLocation    *string                      `pulumi:"computeLocation"`
	ComputeType        string                       `pulumi:"computeType"`
	CreatedOn          string                       `pulumi:"createdOn"`
	Description        *string                      `pulumi:"description"`
	DisableLocalAuth   *bool                        `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                         `pulumi:"isAttachedCompute"`
	ModifiedOn         string                       `pulumi:"modifiedOn"`
	Properties         *HDInsightResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse      `pulumi:"provisioningErrors"`
	ProvisioningState  string                       `pulumi:"provisioningState"`
	ResourceId         *string                      `pulumi:"resourceId"`
}

type HDInsightResponseProperties struct {
	Address              *string                               `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	SshPort              *int                                  `pulumi:"sshPort"`
}

type IdAssetReference struct {
	AssetId       string `pulumi:"assetId"`
	ReferenceType string `pulumi:"referenceType"`
}

type IdAssetReferenceResponse struct {
	AssetId       string `pulumi:"assetId"`
	ReferenceType string `pulumi:"referenceType"`
}

type Identity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityForCmk struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type IdentityForCmkInput interface {
	pulumi.Input

	ToIdentityForCmkOutput() IdentityForCmkOutput
	ToIdentityForCmkOutputWithContext(context.Context) IdentityForCmkOutput
}

type IdentityForCmkArgs struct {
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (IdentityForCmkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityForCmk)(nil)).Elem()
}

func (i IdentityForCmkArgs) ToIdentityForCmkOutput() IdentityForCmkOutput {
	return i.ToIdentityForCmkOutputWithContext(context.Background())
}

func (i IdentityForCmkArgs) ToIdentityForCmkOutputWithContext(ctx context.Context) IdentityForCmkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityForCmkOutput)
}

func (i IdentityForCmkArgs) ToIdentityForCmkPtrOutput() IdentityForCmkPtrOutput {
	return i.ToIdentityForCmkPtrOutputWithContext(context.Background())
}

func (i IdentityForCmkArgs) ToIdentityForCmkPtrOutputWithContext(ctx context.Context) IdentityForCmkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityForCmkOutput).ToIdentityForCmkPtrOutputWithContext(ctx)
}









type IdentityForCmkPtrInput interface {
	pulumi.Input

	ToIdentityForCmkPtrOutput() IdentityForCmkPtrOutput
	ToIdentityForCmkPtrOutputWithContext(context.Context) IdentityForCmkPtrOutput
}

type identityForCmkPtrType IdentityForCmkArgs

func IdentityForCmkPtr(v *IdentityForCmkArgs) IdentityForCmkPtrInput {
	return (*identityForCmkPtrType)(v)
}

func (*identityForCmkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityForCmk)(nil)).Elem()
}

func (i *identityForCmkPtrType) ToIdentityForCmkPtrOutput() IdentityForCmkPtrOutput {
	return i.ToIdentityForCmkPtrOutputWithContext(context.Background())
}

func (i *identityForCmkPtrType) ToIdentityForCmkPtrOutputWithContext(ctx context.Context) IdentityForCmkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityForCmkPtrOutput)
}

type IdentityForCmkOutput struct{ *pulumi.OutputState }

func (IdentityForCmkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityForCmk)(nil)).Elem()
}

func (o IdentityForCmkOutput) ToIdentityForCmkOutput() IdentityForCmkOutput {
	return o
}

func (o IdentityForCmkOutput) ToIdentityForCmkOutputWithContext(ctx context.Context) IdentityForCmkOutput {
	return o
}

func (o IdentityForCmkOutput) ToIdentityForCmkPtrOutput() IdentityForCmkPtrOutput {
	return o.ToIdentityForCmkPtrOutputWithContext(context.Background())
}

func (o IdentityForCmkOutput) ToIdentityForCmkPtrOutputWithContext(ctx context.Context) IdentityForCmkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityForCmk) *IdentityForCmk {
		return &v
	}).(IdentityForCmkPtrOutput)
}

func (o IdentityForCmkOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityForCmk) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type IdentityForCmkPtrOutput struct{ *pulumi.OutputState }

func (IdentityForCmkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityForCmk)(nil)).Elem()
}

func (o IdentityForCmkPtrOutput) ToIdentityForCmkPtrOutput() IdentityForCmkPtrOutput {
	return o
}

func (o IdentityForCmkPtrOutput) ToIdentityForCmkPtrOutputWithContext(ctx context.Context) IdentityForCmkPtrOutput {
	return o
}

func (o IdentityForCmkPtrOutput) Elem() IdentityForCmkOutput {
	return o.ApplyT(func(v *IdentityForCmk) IdentityForCmk {
		if v != nil {
			return *v
		}
		var ret IdentityForCmk
		return ret
	}).(IdentityForCmkOutput)
}

func (o IdentityForCmkPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityForCmk) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type IdentityForCmkResponse struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}

type IdentityForCmkResponseOutput struct{ *pulumi.OutputState }

func (IdentityForCmkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityForCmkResponse)(nil)).Elem()
}

func (o IdentityForCmkResponseOutput) ToIdentityForCmkResponseOutput() IdentityForCmkResponseOutput {
	return o
}

func (o IdentityForCmkResponseOutput) ToIdentityForCmkResponseOutputWithContext(ctx context.Context) IdentityForCmkResponseOutput {
	return o
}

func (o IdentityForCmkResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityForCmkResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type IdentityForCmkResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityForCmkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityForCmkResponse)(nil)).Elem()
}

func (o IdentityForCmkResponsePtrOutput) ToIdentityForCmkResponsePtrOutput() IdentityForCmkResponsePtrOutput {
	return o
}

func (o IdentityForCmkResponsePtrOutput) ToIdentityForCmkResponsePtrOutputWithContext(ctx context.Context) IdentityForCmkResponsePtrOutput {
	return o
}

func (o IdentityForCmkResponsePtrOutput) Elem() IdentityForCmkResponseOutput {
	return o.ApplyT(func(v *IdentityForCmkResponse) IdentityForCmkResponse {
		if v != nil {
			return *v
		}
		var ret IdentityForCmkResponse
		return ret
	}).(IdentityForCmkResponseOutput)
}

func (o IdentityForCmkResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityForCmkResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]UserAssignedIdentityResponse { return v.UserAssignedIdentities }).(UserAssignedIdentityResponseMapOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type InferenceContainerProperties struct {
	LivenessRoute  *Route `pulumi:"livenessRoute"`
	ReadinessRoute *Route `pulumi:"readinessRoute"`
	ScoringRoute   *Route `pulumi:"scoringRoute"`
}





type InferenceContainerPropertiesInput interface {
	pulumi.Input

	ToInferenceContainerPropertiesOutput() InferenceContainerPropertiesOutput
	ToInferenceContainerPropertiesOutputWithContext(context.Context) InferenceContainerPropertiesOutput
}

type InferenceContainerPropertiesArgs struct {
	LivenessRoute  RoutePtrInput `pulumi:"livenessRoute"`
	ReadinessRoute RoutePtrInput `pulumi:"readinessRoute"`
	ScoringRoute   RoutePtrInput `pulumi:"scoringRoute"`
}

func (InferenceContainerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InferenceContainerProperties)(nil)).Elem()
}

func (i InferenceContainerPropertiesArgs) ToInferenceContainerPropertiesOutput() InferenceContainerPropertiesOutput {
	return i.ToInferenceContainerPropertiesOutputWithContext(context.Background())
}

func (i InferenceContainerPropertiesArgs) ToInferenceContainerPropertiesOutputWithContext(ctx context.Context) InferenceContainerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceContainerPropertiesOutput)
}

func (i InferenceContainerPropertiesArgs) ToInferenceContainerPropertiesPtrOutput() InferenceContainerPropertiesPtrOutput {
	return i.ToInferenceContainerPropertiesPtrOutputWithContext(context.Background())
}

func (i InferenceContainerPropertiesArgs) ToInferenceContainerPropertiesPtrOutputWithContext(ctx context.Context) InferenceContainerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceContainerPropertiesOutput).ToInferenceContainerPropertiesPtrOutputWithContext(ctx)
}









type InferenceContainerPropertiesPtrInput interface {
	pulumi.Input

	ToInferenceContainerPropertiesPtrOutput() InferenceContainerPropertiesPtrOutput
	ToInferenceContainerPropertiesPtrOutputWithContext(context.Context) InferenceContainerPropertiesPtrOutput
}

type inferenceContainerPropertiesPtrType InferenceContainerPropertiesArgs

func InferenceContainerPropertiesPtr(v *InferenceContainerPropertiesArgs) InferenceContainerPropertiesPtrInput {
	return (*inferenceContainerPropertiesPtrType)(v)
}

func (*inferenceContainerPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceContainerProperties)(nil)).Elem()
}

func (i *inferenceContainerPropertiesPtrType) ToInferenceContainerPropertiesPtrOutput() InferenceContainerPropertiesPtrOutput {
	return i.ToInferenceContainerPropertiesPtrOutputWithContext(context.Background())
}

func (i *inferenceContainerPropertiesPtrType) ToInferenceContainerPropertiesPtrOutputWithContext(ctx context.Context) InferenceContainerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceContainerPropertiesPtrOutput)
}

type InferenceContainerPropertiesOutput struct{ *pulumi.OutputState }

func (InferenceContainerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InferenceContainerProperties)(nil)).Elem()
}

func (o InferenceContainerPropertiesOutput) ToInferenceContainerPropertiesOutput() InferenceContainerPropertiesOutput {
	return o
}

func (o InferenceContainerPropertiesOutput) ToInferenceContainerPropertiesOutputWithContext(ctx context.Context) InferenceContainerPropertiesOutput {
	return o
}

func (o InferenceContainerPropertiesOutput) ToInferenceContainerPropertiesPtrOutput() InferenceContainerPropertiesPtrOutput {
	return o.ToInferenceContainerPropertiesPtrOutputWithContext(context.Background())
}

func (o InferenceContainerPropertiesOutput) ToInferenceContainerPropertiesPtrOutputWithContext(ctx context.Context) InferenceContainerPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InferenceContainerProperties) *InferenceContainerProperties {
		return &v
	}).(InferenceContainerPropertiesPtrOutput)
}

func (o InferenceContainerPropertiesOutput) LivenessRoute() RoutePtrOutput {
	return o.ApplyT(func(v InferenceContainerProperties) *Route { return v.LivenessRoute }).(RoutePtrOutput)
}

func (o InferenceContainerPropertiesOutput) ReadinessRoute() RoutePtrOutput {
	return o.ApplyT(func(v InferenceContainerProperties) *Route { return v.ReadinessRoute }).(RoutePtrOutput)
}

func (o InferenceContainerPropertiesOutput) ScoringRoute() RoutePtrOutput {
	return o.ApplyT(func(v InferenceContainerProperties) *Route { return v.ScoringRoute }).(RoutePtrOutput)
}

type InferenceContainerPropertiesPtrOutput struct{ *pulumi.OutputState }

func (InferenceContainerPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceContainerProperties)(nil)).Elem()
}

func (o InferenceContainerPropertiesPtrOutput) ToInferenceContainerPropertiesPtrOutput() InferenceContainerPropertiesPtrOutput {
	return o
}

func (o InferenceContainerPropertiesPtrOutput) ToInferenceContainerPropertiesPtrOutputWithContext(ctx context.Context) InferenceContainerPropertiesPtrOutput {
	return o
}

func (o InferenceContainerPropertiesPtrOutput) Elem() InferenceContainerPropertiesOutput {
	return o.ApplyT(func(v *InferenceContainerProperties) InferenceContainerProperties {
		if v != nil {
			return *v
		}
		var ret InferenceContainerProperties
		return ret
	}).(InferenceContainerPropertiesOutput)
}

func (o InferenceContainerPropertiesPtrOutput) LivenessRoute() RoutePtrOutput {
	return o.ApplyT(func(v *InferenceContainerProperties) *Route {
		if v == nil {
			return nil
		}
		return v.LivenessRoute
	}).(RoutePtrOutput)
}

func (o InferenceContainerPropertiesPtrOutput) ReadinessRoute() RoutePtrOutput {
	return o.ApplyT(func(v *InferenceContainerProperties) *Route {
		if v == nil {
			return nil
		}
		return v.ReadinessRoute
	}).(RoutePtrOutput)
}

func (o InferenceContainerPropertiesPtrOutput) ScoringRoute() RoutePtrOutput {
	return o.ApplyT(func(v *InferenceContainerProperties) *Route {
		if v == nil {
			return nil
		}
		return v.ScoringRoute
	}).(RoutePtrOutput)
}

type InferenceContainerPropertiesResponse struct {
	LivenessRoute  *RouteResponse `pulumi:"livenessRoute"`
	ReadinessRoute *RouteResponse `pulumi:"readinessRoute"`
	ScoringRoute   *RouteResponse `pulumi:"scoringRoute"`
}

type InferenceContainerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (InferenceContainerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InferenceContainerPropertiesResponse)(nil)).Elem()
}

func (o InferenceContainerPropertiesResponseOutput) ToInferenceContainerPropertiesResponseOutput() InferenceContainerPropertiesResponseOutput {
	return o
}

func (o InferenceContainerPropertiesResponseOutput) ToInferenceContainerPropertiesResponseOutputWithContext(ctx context.Context) InferenceContainerPropertiesResponseOutput {
	return o
}

func (o InferenceContainerPropertiesResponseOutput) LivenessRoute() RouteResponsePtrOutput {
	return o.ApplyT(func(v InferenceContainerPropertiesResponse) *RouteResponse { return v.LivenessRoute }).(RouteResponsePtrOutput)
}

func (o InferenceContainerPropertiesResponseOutput) ReadinessRoute() RouteResponsePtrOutput {
	return o.ApplyT(func(v InferenceContainerPropertiesResponse) *RouteResponse { return v.ReadinessRoute }).(RouteResponsePtrOutput)
}

func (o InferenceContainerPropertiesResponseOutput) ScoringRoute() RouteResponsePtrOutput {
	return o.ApplyT(func(v InferenceContainerPropertiesResponse) *RouteResponse { return v.ScoringRoute }).(RouteResponsePtrOutput)
}

type InferenceContainerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (InferenceContainerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceContainerPropertiesResponse)(nil)).Elem()
}

func (o InferenceContainerPropertiesResponsePtrOutput) ToInferenceContainerPropertiesResponsePtrOutput() InferenceContainerPropertiesResponsePtrOutput {
	return o
}

func (o InferenceContainerPropertiesResponsePtrOutput) ToInferenceContainerPropertiesResponsePtrOutputWithContext(ctx context.Context) InferenceContainerPropertiesResponsePtrOutput {
	return o
}

func (o InferenceContainerPropertiesResponsePtrOutput) Elem() InferenceContainerPropertiesResponseOutput {
	return o.ApplyT(func(v *InferenceContainerPropertiesResponse) InferenceContainerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret InferenceContainerPropertiesResponse
		return ret
	}).(InferenceContainerPropertiesResponseOutput)
}

func (o InferenceContainerPropertiesResponsePtrOutput) LivenessRoute() RouteResponsePtrOutput {
	return o.ApplyT(func(v *InferenceContainerPropertiesResponse) *RouteResponse {
		if v == nil {
			return nil
		}
		return v.LivenessRoute
	}).(RouteResponsePtrOutput)
}

func (o InferenceContainerPropertiesResponsePtrOutput) ReadinessRoute() RouteResponsePtrOutput {
	return o.ApplyT(func(v *InferenceContainerPropertiesResponse) *RouteResponse {
		if v == nil {
			return nil
		}
		return v.ReadinessRoute
	}).(RouteResponsePtrOutput)
}

func (o InferenceContainerPropertiesResponsePtrOutput) ScoringRoute() RouteResponsePtrOutput {
	return o.ApplyT(func(v *InferenceContainerPropertiesResponse) *RouteResponse {
		if v == nil {
			return nil
		}
		return v.ScoringRoute
	}).(RouteResponsePtrOutput)
}

type InputDataBinding struct {
	DataId        *string `pulumi:"dataId"`
	Mode          *string `pulumi:"mode"`
	PathOnCompute *string `pulumi:"pathOnCompute"`
}

type InputDataBindingResponse struct {
	DataId        *string `pulumi:"dataId"`
	Mode          *string `pulumi:"mode"`
	PathOnCompute *string `pulumi:"pathOnCompute"`
}

type JobEndpointResponse struct {
	Endpoint        *string           `pulumi:"endpoint"`
	JobEndpointType *string           `pulumi:"jobEndpointType"`
	Port            *int              `pulumi:"port"`
	Properties      map[string]string `pulumi:"properties"`
}

type JobEndpointResponseOutput struct{ *pulumi.OutputState }

func (JobEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobEndpointResponse)(nil)).Elem()
}

func (o JobEndpointResponseOutput) ToJobEndpointResponseOutput() JobEndpointResponseOutput {
	return o
}

func (o JobEndpointResponseOutput) ToJobEndpointResponseOutputWithContext(ctx context.Context) JobEndpointResponseOutput {
	return o
}

func (o JobEndpointResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobEndpointResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o JobEndpointResponseOutput) JobEndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobEndpointResponse) *string { return v.JobEndpointType }).(pulumi.StringPtrOutput)
}

func (o JobEndpointResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobEndpointResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o JobEndpointResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v JobEndpointResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

type JobEndpointResponseMapOutput struct{ *pulumi.OutputState }

func (JobEndpointResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]JobEndpointResponse)(nil)).Elem()
}

func (o JobEndpointResponseMapOutput) ToJobEndpointResponseMapOutput() JobEndpointResponseMapOutput {
	return o
}

func (o JobEndpointResponseMapOutput) ToJobEndpointResponseMapOutputWithContext(ctx context.Context) JobEndpointResponseMapOutput {
	return o
}

func (o JobEndpointResponseMapOutput) MapIndex(k pulumi.StringInput) JobEndpointResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) JobEndpointResponse {
		return vs[0].(map[string]JobEndpointResponse)[vs[1].(string)]
	}).(JobEndpointResponseOutput)
}

type JobOutputResponse struct {
	DatastoreId string `pulumi:"datastoreId"`
	Path        string `pulumi:"path"`
}

type K8sOnlineDeployment struct {
	AppInsightsEnabled            *bool                          `pulumi:"appInsightsEnabled"`
	CodeConfiguration             *CodeConfiguration             `pulumi:"codeConfiguration"`
	ContainerResourceRequirements *ContainerResourceRequirements `pulumi:"containerResourceRequirements"`
	Description                   *string                        `pulumi:"description"`
	EndpointComputeType           string                         `pulumi:"endpointComputeType"`
	EnvironmentId                 *string                        `pulumi:"environmentId"`
	EnvironmentVariables          map[string]string              `pulumi:"environmentVariables"`
	LivenessProbe                 *ProbeSettings                 `pulumi:"livenessProbe"`
	Model                         interface{}                    `pulumi:"model"`
	Properties                    map[string]string              `pulumi:"properties"`
	RequestSettings               *OnlineRequestSettings         `pulumi:"requestSettings"`
	ScaleSettings                 interface{}                    `pulumi:"scaleSettings"`
}

type K8sOnlineDeploymentResponse struct {
	AppInsightsEnabled            *bool                                  `pulumi:"appInsightsEnabled"`
	CodeConfiguration             *CodeConfigurationResponse             `pulumi:"codeConfiguration"`
	ContainerResourceRequirements *ContainerResourceRequirementsResponse `pulumi:"containerResourceRequirements"`
	Description                   *string                                `pulumi:"description"`
	EndpointComputeType           string                                 `pulumi:"endpointComputeType"`
	EnvironmentId                 *string                                `pulumi:"environmentId"`
	EnvironmentVariables          map[string]string                      `pulumi:"environmentVariables"`
	LivenessProbe                 *ProbeSettingsResponse                 `pulumi:"livenessProbe"`
	Model                         interface{}                            `pulumi:"model"`
	Properties                    map[string]string                      `pulumi:"properties"`
	ProvisioningState             string                                 `pulumi:"provisioningState"`
	RequestSettings               *OnlineRequestSettingsResponse         `pulumi:"requestSettings"`
	ScaleSettings                 interface{}                            `pulumi:"scaleSettings"`
}

type KeyVaultProperties struct {
	IdentityClientId *string `pulumi:"identityClientId"`
	KeyIdentifier    string  `pulumi:"keyIdentifier"`
	KeyVaultArmId    string  `pulumi:"keyVaultArmId"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	IdentityClientId pulumi.StringPtrInput `pulumi:"identityClientId"`
	KeyIdentifier    pulumi.StringInput    `pulumi:"keyIdentifier"`
	KeyVaultArmId    pulumi.StringInput    `pulumi:"keyVaultArmId"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultProperties) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultArmId() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultProperties) string { return v.KeyVaultArmId }).(pulumi.StringOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultProperties
		return ret
	}).(KeyVaultPropertiesOutput)
}

func (o KeyVaultPropertiesPtrOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.IdentityClientId
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVaultArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultArmId
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	IdentityClientId *string `pulumi:"identityClientId"`
	KeyIdentifier    string  `pulumi:"keyIdentifier"`
	KeyVaultArmId    string  `pulumi:"keyVaultArmId"`
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultArmId() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.KeyVaultArmId }).(pulumi.StringOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultPropertiesResponse
		return ret
	}).(KeyVaultPropertiesResponseOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IdentityClientId
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVaultArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultArmId
	}).(pulumi.StringPtrOutput)
}

type LabelCategory struct {
	AllowMultiSelect *bool                 `pulumi:"allowMultiSelect"`
	Classes          map[string]LabelClass `pulumi:"classes"`
	DisplayName      *string               `pulumi:"displayName"`
}





type LabelCategoryInput interface {
	pulumi.Input

	ToLabelCategoryOutput() LabelCategoryOutput
	ToLabelCategoryOutputWithContext(context.Context) LabelCategoryOutput
}

type LabelCategoryArgs struct {
	AllowMultiSelect pulumi.BoolPtrInput   `pulumi:"allowMultiSelect"`
	Classes          LabelClassMapInput    `pulumi:"classes"`
	DisplayName      pulumi.StringPtrInput `pulumi:"displayName"`
}

func (LabelCategoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelCategory)(nil)).Elem()
}

func (i LabelCategoryArgs) ToLabelCategoryOutput() LabelCategoryOutput {
	return i.ToLabelCategoryOutputWithContext(context.Background())
}

func (i LabelCategoryArgs) ToLabelCategoryOutputWithContext(ctx context.Context) LabelCategoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelCategoryOutput)
}





type LabelCategoryMapInput interface {
	pulumi.Input

	ToLabelCategoryMapOutput() LabelCategoryMapOutput
	ToLabelCategoryMapOutputWithContext(context.Context) LabelCategoryMapOutput
}

type LabelCategoryMap map[string]LabelCategoryInput

func (LabelCategoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LabelCategory)(nil)).Elem()
}

func (i LabelCategoryMap) ToLabelCategoryMapOutput() LabelCategoryMapOutput {
	return i.ToLabelCategoryMapOutputWithContext(context.Background())
}

func (i LabelCategoryMap) ToLabelCategoryMapOutputWithContext(ctx context.Context) LabelCategoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelCategoryMapOutput)
}

type LabelCategoryOutput struct{ *pulumi.OutputState }

func (LabelCategoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelCategory)(nil)).Elem()
}

func (o LabelCategoryOutput) ToLabelCategoryOutput() LabelCategoryOutput {
	return o
}

func (o LabelCategoryOutput) ToLabelCategoryOutputWithContext(ctx context.Context) LabelCategoryOutput {
	return o
}

func (o LabelCategoryOutput) AllowMultiSelect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabelCategory) *bool { return v.AllowMultiSelect }).(pulumi.BoolPtrOutput)
}

func (o LabelCategoryOutput) Classes() LabelClassMapOutput {
	return o.ApplyT(func(v LabelCategory) map[string]LabelClass { return v.Classes }).(LabelClassMapOutput)
}

func (o LabelCategoryOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelCategory) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

type LabelCategoryMapOutput struct{ *pulumi.OutputState }

func (LabelCategoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LabelCategory)(nil)).Elem()
}

func (o LabelCategoryMapOutput) ToLabelCategoryMapOutput() LabelCategoryMapOutput {
	return o
}

func (o LabelCategoryMapOutput) ToLabelCategoryMapOutputWithContext(ctx context.Context) LabelCategoryMapOutput {
	return o
}

func (o LabelCategoryMapOutput) MapIndex(k pulumi.StringInput) LabelCategoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LabelCategory {
		return vs[0].(map[string]LabelCategory)[vs[1].(string)]
	}).(LabelCategoryOutput)
}

type LabelCategoryResponse struct {
	AllowMultiSelect *bool                         `pulumi:"allowMultiSelect"`
	Classes          map[string]LabelClassResponse `pulumi:"classes"`
	DisplayName      *string                       `pulumi:"displayName"`
}

type LabelCategoryResponseOutput struct{ *pulumi.OutputState }

func (LabelCategoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelCategoryResponse)(nil)).Elem()
}

func (o LabelCategoryResponseOutput) ToLabelCategoryResponseOutput() LabelCategoryResponseOutput {
	return o
}

func (o LabelCategoryResponseOutput) ToLabelCategoryResponseOutputWithContext(ctx context.Context) LabelCategoryResponseOutput {
	return o
}

func (o LabelCategoryResponseOutput) AllowMultiSelect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabelCategoryResponse) *bool { return v.AllowMultiSelect }).(pulumi.BoolPtrOutput)
}

func (o LabelCategoryResponseOutput) Classes() LabelClassResponseMapOutput {
	return o.ApplyT(func(v LabelCategoryResponse) map[string]LabelClassResponse { return v.Classes }).(LabelClassResponseMapOutput)
}

func (o LabelCategoryResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelCategoryResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

type LabelCategoryResponseMapOutput struct{ *pulumi.OutputState }

func (LabelCategoryResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LabelCategoryResponse)(nil)).Elem()
}

func (o LabelCategoryResponseMapOutput) ToLabelCategoryResponseMapOutput() LabelCategoryResponseMapOutput {
	return o
}

func (o LabelCategoryResponseMapOutput) ToLabelCategoryResponseMapOutputWithContext(ctx context.Context) LabelCategoryResponseMapOutput {
	return o
}

func (o LabelCategoryResponseMapOutput) MapIndex(k pulumi.StringInput) LabelCategoryResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LabelCategoryResponse {
		return vs[0].(map[string]LabelCategoryResponse)[vs[1].(string)]
	}).(LabelCategoryResponseOutput)
}

type LabelClass struct {
	DisplayName *string               `pulumi:"displayName"`
	Subclasses  map[string]LabelClass `pulumi:"subclasses"`
}





type LabelClassInput interface {
	pulumi.Input

	ToLabelClassOutput() LabelClassOutput
	ToLabelClassOutputWithContext(context.Context) LabelClassOutput
}

type LabelClassArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Subclasses  LabelClassMapInput    `pulumi:"subclasses"`
}

func (LabelClassArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelClass)(nil)).Elem()
}

func (i LabelClassArgs) ToLabelClassOutput() LabelClassOutput {
	return i.ToLabelClassOutputWithContext(context.Background())
}

func (i LabelClassArgs) ToLabelClassOutputWithContext(ctx context.Context) LabelClassOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelClassOutput)
}





type LabelClassMapInput interface {
	pulumi.Input

	ToLabelClassMapOutput() LabelClassMapOutput
	ToLabelClassMapOutputWithContext(context.Context) LabelClassMapOutput
}

type LabelClassMap map[string]LabelClassInput

func (LabelClassMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LabelClass)(nil)).Elem()
}

func (i LabelClassMap) ToLabelClassMapOutput() LabelClassMapOutput {
	return i.ToLabelClassMapOutputWithContext(context.Background())
}

func (i LabelClassMap) ToLabelClassMapOutputWithContext(ctx context.Context) LabelClassMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelClassMapOutput)
}

type LabelClassOutput struct{ *pulumi.OutputState }

func (LabelClassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelClass)(nil)).Elem()
}

func (o LabelClassOutput) ToLabelClassOutput() LabelClassOutput {
	return o
}

func (o LabelClassOutput) ToLabelClassOutputWithContext(ctx context.Context) LabelClassOutput {
	return o
}

func (o LabelClassOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelClass) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LabelClassOutput) Subclasses() LabelClassMapOutput {
	return o.ApplyT(func(v LabelClass) map[string]LabelClass { return v.Subclasses }).(LabelClassMapOutput)
}

type LabelClassMapOutput struct{ *pulumi.OutputState }

func (LabelClassMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LabelClass)(nil)).Elem()
}

func (o LabelClassMapOutput) ToLabelClassMapOutput() LabelClassMapOutput {
	return o
}

func (o LabelClassMapOutput) ToLabelClassMapOutputWithContext(ctx context.Context) LabelClassMapOutput {
	return o
}

func (o LabelClassMapOutput) MapIndex(k pulumi.StringInput) LabelClassOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LabelClass {
		return vs[0].(map[string]LabelClass)[vs[1].(string)]
	}).(LabelClassOutput)
}

type LabelClassResponse struct {
	DisplayName *string                       `pulumi:"displayName"`
	Subclasses  map[string]LabelClassResponse `pulumi:"subclasses"`
}

type LabelClassResponseOutput struct{ *pulumi.OutputState }

func (LabelClassResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelClassResponse)(nil)).Elem()
}

func (o LabelClassResponseOutput) ToLabelClassResponseOutput() LabelClassResponseOutput {
	return o
}

func (o LabelClassResponseOutput) ToLabelClassResponseOutputWithContext(ctx context.Context) LabelClassResponseOutput {
	return o
}

func (o LabelClassResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelClassResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LabelClassResponseOutput) Subclasses() LabelClassResponseMapOutput {
	return o.ApplyT(func(v LabelClassResponse) map[string]LabelClassResponse { return v.Subclasses }).(LabelClassResponseMapOutput)
}

type LabelClassResponseMapOutput struct{ *pulumi.OutputState }

func (LabelClassResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LabelClassResponse)(nil)).Elem()
}

func (o LabelClassResponseMapOutput) ToLabelClassResponseMapOutput() LabelClassResponseMapOutput {
	return o
}

func (o LabelClassResponseMapOutput) ToLabelClassResponseMapOutputWithContext(ctx context.Context) LabelClassResponseMapOutput {
	return o
}

func (o LabelClassResponseMapOutput) MapIndex(k pulumi.StringInput) LabelClassResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LabelClassResponse {
		return vs[0].(map[string]LabelClassResponse)[vs[1].(string)]
	}).(LabelClassResponseOutput)
}

type LabelingDatasetConfiguration struct {
	AssetName                        *string `pulumi:"assetName"`
	DatasetVersion                   *string `pulumi:"datasetVersion"`
	IncrementalDatasetRefreshEnabled *bool   `pulumi:"incrementalDatasetRefreshEnabled"`
}





type LabelingDatasetConfigurationInput interface {
	pulumi.Input

	ToLabelingDatasetConfigurationOutput() LabelingDatasetConfigurationOutput
	ToLabelingDatasetConfigurationOutputWithContext(context.Context) LabelingDatasetConfigurationOutput
}

type LabelingDatasetConfigurationArgs struct {
	AssetName                        pulumi.StringPtrInput `pulumi:"assetName"`
	DatasetVersion                   pulumi.StringPtrInput `pulumi:"datasetVersion"`
	IncrementalDatasetRefreshEnabled pulumi.BoolPtrInput   `pulumi:"incrementalDatasetRefreshEnabled"`
}

func (LabelingDatasetConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingDatasetConfiguration)(nil)).Elem()
}

func (i LabelingDatasetConfigurationArgs) ToLabelingDatasetConfigurationOutput() LabelingDatasetConfigurationOutput {
	return i.ToLabelingDatasetConfigurationOutputWithContext(context.Background())
}

func (i LabelingDatasetConfigurationArgs) ToLabelingDatasetConfigurationOutputWithContext(ctx context.Context) LabelingDatasetConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingDatasetConfigurationOutput)
}

func (i LabelingDatasetConfigurationArgs) ToLabelingDatasetConfigurationPtrOutput() LabelingDatasetConfigurationPtrOutput {
	return i.ToLabelingDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i LabelingDatasetConfigurationArgs) ToLabelingDatasetConfigurationPtrOutputWithContext(ctx context.Context) LabelingDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingDatasetConfigurationOutput).ToLabelingDatasetConfigurationPtrOutputWithContext(ctx)
}









type LabelingDatasetConfigurationPtrInput interface {
	pulumi.Input

	ToLabelingDatasetConfigurationPtrOutput() LabelingDatasetConfigurationPtrOutput
	ToLabelingDatasetConfigurationPtrOutputWithContext(context.Context) LabelingDatasetConfigurationPtrOutput
}

type labelingDatasetConfigurationPtrType LabelingDatasetConfigurationArgs

func LabelingDatasetConfigurationPtr(v *LabelingDatasetConfigurationArgs) LabelingDatasetConfigurationPtrInput {
	return (*labelingDatasetConfigurationPtrType)(v)
}

func (*labelingDatasetConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingDatasetConfiguration)(nil)).Elem()
}

func (i *labelingDatasetConfigurationPtrType) ToLabelingDatasetConfigurationPtrOutput() LabelingDatasetConfigurationPtrOutput {
	return i.ToLabelingDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i *labelingDatasetConfigurationPtrType) ToLabelingDatasetConfigurationPtrOutputWithContext(ctx context.Context) LabelingDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingDatasetConfigurationPtrOutput)
}

type LabelingDatasetConfigurationOutput struct{ *pulumi.OutputState }

func (LabelingDatasetConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingDatasetConfiguration)(nil)).Elem()
}

func (o LabelingDatasetConfigurationOutput) ToLabelingDatasetConfigurationOutput() LabelingDatasetConfigurationOutput {
	return o
}

func (o LabelingDatasetConfigurationOutput) ToLabelingDatasetConfigurationOutputWithContext(ctx context.Context) LabelingDatasetConfigurationOutput {
	return o
}

func (o LabelingDatasetConfigurationOutput) ToLabelingDatasetConfigurationPtrOutput() LabelingDatasetConfigurationPtrOutput {
	return o.ToLabelingDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (o LabelingDatasetConfigurationOutput) ToLabelingDatasetConfigurationPtrOutputWithContext(ctx context.Context) LabelingDatasetConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabelingDatasetConfiguration) *LabelingDatasetConfiguration {
		return &v
	}).(LabelingDatasetConfigurationPtrOutput)
}

func (o LabelingDatasetConfigurationOutput) AssetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelingDatasetConfiguration) *string { return v.AssetName }).(pulumi.StringPtrOutput)
}

func (o LabelingDatasetConfigurationOutput) DatasetVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelingDatasetConfiguration) *string { return v.DatasetVersion }).(pulumi.StringPtrOutput)
}

func (o LabelingDatasetConfigurationOutput) IncrementalDatasetRefreshEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabelingDatasetConfiguration) *bool { return v.IncrementalDatasetRefreshEnabled }).(pulumi.BoolPtrOutput)
}

type LabelingDatasetConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LabelingDatasetConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingDatasetConfiguration)(nil)).Elem()
}

func (o LabelingDatasetConfigurationPtrOutput) ToLabelingDatasetConfigurationPtrOutput() LabelingDatasetConfigurationPtrOutput {
	return o
}

func (o LabelingDatasetConfigurationPtrOutput) ToLabelingDatasetConfigurationPtrOutputWithContext(ctx context.Context) LabelingDatasetConfigurationPtrOutput {
	return o
}

func (o LabelingDatasetConfigurationPtrOutput) Elem() LabelingDatasetConfigurationOutput {
	return o.ApplyT(func(v *LabelingDatasetConfiguration) LabelingDatasetConfiguration {
		if v != nil {
			return *v
		}
		var ret LabelingDatasetConfiguration
		return ret
	}).(LabelingDatasetConfigurationOutput)
}

func (o LabelingDatasetConfigurationPtrOutput) AssetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelingDatasetConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AssetName
	}).(pulumi.StringPtrOutput)
}

func (o LabelingDatasetConfigurationPtrOutput) DatasetVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelingDatasetConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.DatasetVersion
	}).(pulumi.StringPtrOutput)
}

func (o LabelingDatasetConfigurationPtrOutput) IncrementalDatasetRefreshEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabelingDatasetConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.IncrementalDatasetRefreshEnabled
	}).(pulumi.BoolPtrOutput)
}

type LabelingDatasetConfigurationResponse struct {
	AssetName                        *string `pulumi:"assetName"`
	DatasetVersion                   *string `pulumi:"datasetVersion"`
	IncrementalDatasetRefreshEnabled *bool   `pulumi:"incrementalDatasetRefreshEnabled"`
}

type LabelingDatasetConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LabelingDatasetConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingDatasetConfigurationResponse)(nil)).Elem()
}

func (o LabelingDatasetConfigurationResponseOutput) ToLabelingDatasetConfigurationResponseOutput() LabelingDatasetConfigurationResponseOutput {
	return o
}

func (o LabelingDatasetConfigurationResponseOutput) ToLabelingDatasetConfigurationResponseOutputWithContext(ctx context.Context) LabelingDatasetConfigurationResponseOutput {
	return o
}

func (o LabelingDatasetConfigurationResponseOutput) AssetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelingDatasetConfigurationResponse) *string { return v.AssetName }).(pulumi.StringPtrOutput)
}

func (o LabelingDatasetConfigurationResponseOutput) DatasetVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelingDatasetConfigurationResponse) *string { return v.DatasetVersion }).(pulumi.StringPtrOutput)
}

func (o LabelingDatasetConfigurationResponseOutput) IncrementalDatasetRefreshEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabelingDatasetConfigurationResponse) *bool { return v.IncrementalDatasetRefreshEnabled }).(pulumi.BoolPtrOutput)
}

type LabelingDatasetConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LabelingDatasetConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingDatasetConfigurationResponse)(nil)).Elem()
}

func (o LabelingDatasetConfigurationResponsePtrOutput) ToLabelingDatasetConfigurationResponsePtrOutput() LabelingDatasetConfigurationResponsePtrOutput {
	return o
}

func (o LabelingDatasetConfigurationResponsePtrOutput) ToLabelingDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) LabelingDatasetConfigurationResponsePtrOutput {
	return o
}

func (o LabelingDatasetConfigurationResponsePtrOutput) Elem() LabelingDatasetConfigurationResponseOutput {
	return o.ApplyT(func(v *LabelingDatasetConfigurationResponse) LabelingDatasetConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LabelingDatasetConfigurationResponse
		return ret
	}).(LabelingDatasetConfigurationResponseOutput)
}

func (o LabelingDatasetConfigurationResponsePtrOutput) AssetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelingDatasetConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AssetName
	}).(pulumi.StringPtrOutput)
}

func (o LabelingDatasetConfigurationResponsePtrOutput) DatasetVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelingDatasetConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DatasetVersion
	}).(pulumi.StringPtrOutput)
}

func (o LabelingDatasetConfigurationResponsePtrOutput) IncrementalDatasetRefreshEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabelingDatasetConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IncrementalDatasetRefreshEnabled
	}).(pulumi.BoolPtrOutput)
}

type LabelingJobType struct {
	DatasetConfiguration       *LabelingDatasetConfiguration `pulumi:"datasetConfiguration"`
	Description                *string                       `pulumi:"description"`
	JobInstructions            *LabelingJobInstructions      `pulumi:"jobInstructions"`
	JobType                    string                        `pulumi:"jobType"`
	LabelCategories            map[string]LabelCategory      `pulumi:"labelCategories"`
	LabelingJobMediaProperties interface{}                   `pulumi:"labelingJobMediaProperties"`
	MlAssistConfiguration      *MLAssistConfiguration        `pulumi:"mlAssistConfiguration"`
	Properties                 map[string]string             `pulumi:"properties"`
	Tags                       map[string]string             `pulumi:"tags"`
}





type LabelingJobTypeInput interface {
	pulumi.Input

	ToLabelingJobTypeOutput() LabelingJobTypeOutput
	ToLabelingJobTypeOutputWithContext(context.Context) LabelingJobTypeOutput
}

type LabelingJobTypeArgs struct {
	DatasetConfiguration       LabelingDatasetConfigurationPtrInput `pulumi:"datasetConfiguration"`
	Description                pulumi.StringPtrInput                `pulumi:"description"`
	JobInstructions            LabelingJobInstructionsPtrInput      `pulumi:"jobInstructions"`
	JobType                    pulumi.StringInput                   `pulumi:"jobType"`
	LabelCategories            LabelCategoryMapInput                `pulumi:"labelCategories"`
	LabelingJobMediaProperties pulumi.Input                         `pulumi:"labelingJobMediaProperties"`
	MlAssistConfiguration      MLAssistConfigurationPtrInput        `pulumi:"mlAssistConfiguration"`
	Properties                 pulumi.StringMapInput                `pulumi:"properties"`
	Tags                       pulumi.StringMapInput                `pulumi:"tags"`
}

func (LabelingJobTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobType)(nil)).Elem()
}

func (i LabelingJobTypeArgs) ToLabelingJobTypeOutput() LabelingJobTypeOutput {
	return i.ToLabelingJobTypeOutputWithContext(context.Background())
}

func (i LabelingJobTypeArgs) ToLabelingJobTypeOutputWithContext(ctx context.Context) LabelingJobTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobTypeOutput)
}

type LabelingJobTypeOutput struct{ *pulumi.OutputState }

func (LabelingJobTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobType)(nil)).Elem()
}

func (o LabelingJobTypeOutput) ToLabelingJobTypeOutput() LabelingJobTypeOutput {
	return o
}

func (o LabelingJobTypeOutput) ToLabelingJobTypeOutputWithContext(ctx context.Context) LabelingJobTypeOutput {
	return o
}

func (o LabelingJobTypeOutput) DatasetConfiguration() LabelingDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v LabelingJobType) *LabelingDatasetConfiguration { return v.DatasetConfiguration }).(LabelingDatasetConfigurationPtrOutput)
}

func (o LabelingJobTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelingJobType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LabelingJobTypeOutput) JobInstructions() LabelingJobInstructionsPtrOutput {
	return o.ApplyT(func(v LabelingJobType) *LabelingJobInstructions { return v.JobInstructions }).(LabelingJobInstructionsPtrOutput)
}

func (o LabelingJobTypeOutput) JobType() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingJobType) string { return v.JobType }).(pulumi.StringOutput)
}

func (o LabelingJobTypeOutput) LabelCategories() LabelCategoryMapOutput {
	return o.ApplyT(func(v LabelingJobType) map[string]LabelCategory { return v.LabelCategories }).(LabelCategoryMapOutput)
}

func (o LabelingJobTypeOutput) LabelingJobMediaProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LabelingJobType) interface{} { return v.LabelingJobMediaProperties }).(pulumi.AnyOutput)
}

func (o LabelingJobTypeOutput) MlAssistConfiguration() MLAssistConfigurationPtrOutput {
	return o.ApplyT(func(v LabelingJobType) *MLAssistConfiguration { return v.MlAssistConfiguration }).(MLAssistConfigurationPtrOutput)
}

func (o LabelingJobTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabelingJobType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o LabelingJobTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabelingJobType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type LabelingJobImageProperties struct {
	AnnotationType *string `pulumi:"annotationType"`
	MediaType      string  `pulumi:"mediaType"`
}

type LabelingJobImagePropertiesResponse struct {
	AnnotationType *string `pulumi:"annotationType"`
	MediaType      string  `pulumi:"mediaType"`
}

type LabelingJobInstructions struct {
	Uri *string `pulumi:"uri"`
}





type LabelingJobInstructionsInput interface {
	pulumi.Input

	ToLabelingJobInstructionsOutput() LabelingJobInstructionsOutput
	ToLabelingJobInstructionsOutputWithContext(context.Context) LabelingJobInstructionsOutput
}

type LabelingJobInstructionsArgs struct {
	Uri pulumi.StringPtrInput `pulumi:"uri"`
}

func (LabelingJobInstructionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobInstructions)(nil)).Elem()
}

func (i LabelingJobInstructionsArgs) ToLabelingJobInstructionsOutput() LabelingJobInstructionsOutput {
	return i.ToLabelingJobInstructionsOutputWithContext(context.Background())
}

func (i LabelingJobInstructionsArgs) ToLabelingJobInstructionsOutputWithContext(ctx context.Context) LabelingJobInstructionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobInstructionsOutput)
}

func (i LabelingJobInstructionsArgs) ToLabelingJobInstructionsPtrOutput() LabelingJobInstructionsPtrOutput {
	return i.ToLabelingJobInstructionsPtrOutputWithContext(context.Background())
}

func (i LabelingJobInstructionsArgs) ToLabelingJobInstructionsPtrOutputWithContext(ctx context.Context) LabelingJobInstructionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobInstructionsOutput).ToLabelingJobInstructionsPtrOutputWithContext(ctx)
}









type LabelingJobInstructionsPtrInput interface {
	pulumi.Input

	ToLabelingJobInstructionsPtrOutput() LabelingJobInstructionsPtrOutput
	ToLabelingJobInstructionsPtrOutputWithContext(context.Context) LabelingJobInstructionsPtrOutput
}

type labelingJobInstructionsPtrType LabelingJobInstructionsArgs

func LabelingJobInstructionsPtr(v *LabelingJobInstructionsArgs) LabelingJobInstructionsPtrInput {
	return (*labelingJobInstructionsPtrType)(v)
}

func (*labelingJobInstructionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingJobInstructions)(nil)).Elem()
}

func (i *labelingJobInstructionsPtrType) ToLabelingJobInstructionsPtrOutput() LabelingJobInstructionsPtrOutput {
	return i.ToLabelingJobInstructionsPtrOutputWithContext(context.Background())
}

func (i *labelingJobInstructionsPtrType) ToLabelingJobInstructionsPtrOutputWithContext(ctx context.Context) LabelingJobInstructionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobInstructionsPtrOutput)
}

type LabelingJobInstructionsOutput struct{ *pulumi.OutputState }

func (LabelingJobInstructionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobInstructions)(nil)).Elem()
}

func (o LabelingJobInstructionsOutput) ToLabelingJobInstructionsOutput() LabelingJobInstructionsOutput {
	return o
}

func (o LabelingJobInstructionsOutput) ToLabelingJobInstructionsOutputWithContext(ctx context.Context) LabelingJobInstructionsOutput {
	return o
}

func (o LabelingJobInstructionsOutput) ToLabelingJobInstructionsPtrOutput() LabelingJobInstructionsPtrOutput {
	return o.ToLabelingJobInstructionsPtrOutputWithContext(context.Background())
}

func (o LabelingJobInstructionsOutput) ToLabelingJobInstructionsPtrOutputWithContext(ctx context.Context) LabelingJobInstructionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabelingJobInstructions) *LabelingJobInstructions {
		return &v
	}).(LabelingJobInstructionsPtrOutput)
}

func (o LabelingJobInstructionsOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelingJobInstructions) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type LabelingJobInstructionsPtrOutput struct{ *pulumi.OutputState }

func (LabelingJobInstructionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingJobInstructions)(nil)).Elem()
}

func (o LabelingJobInstructionsPtrOutput) ToLabelingJobInstructionsPtrOutput() LabelingJobInstructionsPtrOutput {
	return o
}

func (o LabelingJobInstructionsPtrOutput) ToLabelingJobInstructionsPtrOutputWithContext(ctx context.Context) LabelingJobInstructionsPtrOutput {
	return o
}

func (o LabelingJobInstructionsPtrOutput) Elem() LabelingJobInstructionsOutput {
	return o.ApplyT(func(v *LabelingJobInstructions) LabelingJobInstructions {
		if v != nil {
			return *v
		}
		var ret LabelingJobInstructions
		return ret
	}).(LabelingJobInstructionsOutput)
}

func (o LabelingJobInstructionsPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelingJobInstructions) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type LabelingJobInstructionsResponse struct {
	Uri *string `pulumi:"uri"`
}

type LabelingJobInstructionsResponseOutput struct{ *pulumi.OutputState }

func (LabelingJobInstructionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobInstructionsResponse)(nil)).Elem()
}

func (o LabelingJobInstructionsResponseOutput) ToLabelingJobInstructionsResponseOutput() LabelingJobInstructionsResponseOutput {
	return o
}

func (o LabelingJobInstructionsResponseOutput) ToLabelingJobInstructionsResponseOutputWithContext(ctx context.Context) LabelingJobInstructionsResponseOutput {
	return o
}

func (o LabelingJobInstructionsResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelingJobInstructionsResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type LabelingJobInstructionsResponsePtrOutput struct{ *pulumi.OutputState }

func (LabelingJobInstructionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingJobInstructionsResponse)(nil)).Elem()
}

func (o LabelingJobInstructionsResponsePtrOutput) ToLabelingJobInstructionsResponsePtrOutput() LabelingJobInstructionsResponsePtrOutput {
	return o
}

func (o LabelingJobInstructionsResponsePtrOutput) ToLabelingJobInstructionsResponsePtrOutputWithContext(ctx context.Context) LabelingJobInstructionsResponsePtrOutput {
	return o
}

func (o LabelingJobInstructionsResponsePtrOutput) Elem() LabelingJobInstructionsResponseOutput {
	return o.ApplyT(func(v *LabelingJobInstructionsResponse) LabelingJobInstructionsResponse {
		if v != nil {
			return *v
		}
		var ret LabelingJobInstructionsResponse
		return ret
	}).(LabelingJobInstructionsResponseOutput)
}

func (o LabelingJobInstructionsResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelingJobInstructionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type LabelingJobResponse struct {
	CreatedTimeUtc             string                                `pulumi:"createdTimeUtc"`
	DatasetConfiguration       *LabelingDatasetConfigurationResponse `pulumi:"datasetConfiguration"`
	Description                *string                               `pulumi:"description"`
	InteractionEndpoints       map[string]JobEndpointResponse        `pulumi:"interactionEndpoints"`
	JobInstructions            *LabelingJobInstructionsResponse      `pulumi:"jobInstructions"`
	JobType                    string                                `pulumi:"jobType"`
	LabelCategories            map[string]LabelCategoryResponse      `pulumi:"labelCategories"`
	LabelingJobMediaProperties interface{}                           `pulumi:"labelingJobMediaProperties"`
	MlAssistConfiguration      *MLAssistConfigurationResponse        `pulumi:"mlAssistConfiguration"`
	ProgressMetrics            ProgressMetricsResponse               `pulumi:"progressMetrics"`
	ProjectId                  string                                `pulumi:"projectId"`
	Properties                 map[string]string                     `pulumi:"properties"`
	ProvisioningState          string                                `pulumi:"provisioningState"`
	Status                     string                                `pulumi:"status"`
	StatusMessages             []StatusMessageResponse               `pulumi:"statusMessages"`
	Tags                       map[string]string                     `pulumi:"tags"`
}

type LabelingJobResponseOutput struct{ *pulumi.OutputState }

func (LabelingJobResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobResponse)(nil)).Elem()
}

func (o LabelingJobResponseOutput) ToLabelingJobResponseOutput() LabelingJobResponseOutput {
	return o
}

func (o LabelingJobResponseOutput) ToLabelingJobResponseOutputWithContext(ctx context.Context) LabelingJobResponseOutput {
	return o
}

func (o LabelingJobResponseOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingJobResponse) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o LabelingJobResponseOutput) DatasetConfiguration() LabelingDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LabelingJobResponse) *LabelingDatasetConfigurationResponse { return v.DatasetConfiguration }).(LabelingDatasetConfigurationResponsePtrOutput)
}

func (o LabelingJobResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelingJobResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LabelingJobResponseOutput) InteractionEndpoints() JobEndpointResponseMapOutput {
	return o.ApplyT(func(v LabelingJobResponse) map[string]JobEndpointResponse { return v.InteractionEndpoints }).(JobEndpointResponseMapOutput)
}

func (o LabelingJobResponseOutput) JobInstructions() LabelingJobInstructionsResponsePtrOutput {
	return o.ApplyT(func(v LabelingJobResponse) *LabelingJobInstructionsResponse { return v.JobInstructions }).(LabelingJobInstructionsResponsePtrOutput)
}

func (o LabelingJobResponseOutput) JobType() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingJobResponse) string { return v.JobType }).(pulumi.StringOutput)
}

func (o LabelingJobResponseOutput) LabelCategories() LabelCategoryResponseMapOutput {
	return o.ApplyT(func(v LabelingJobResponse) map[string]LabelCategoryResponse { return v.LabelCategories }).(LabelCategoryResponseMapOutput)
}

func (o LabelingJobResponseOutput) LabelingJobMediaProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LabelingJobResponse) interface{} { return v.LabelingJobMediaProperties }).(pulumi.AnyOutput)
}

func (o LabelingJobResponseOutput) MlAssistConfiguration() MLAssistConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LabelingJobResponse) *MLAssistConfigurationResponse { return v.MlAssistConfiguration }).(MLAssistConfigurationResponsePtrOutput)
}

func (o LabelingJobResponseOutput) ProgressMetrics() ProgressMetricsResponseOutput {
	return o.ApplyT(func(v LabelingJobResponse) ProgressMetricsResponse { return v.ProgressMetrics }).(ProgressMetricsResponseOutput)
}

func (o LabelingJobResponseOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingJobResponse) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LabelingJobResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabelingJobResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o LabelingJobResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingJobResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LabelingJobResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingJobResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o LabelingJobResponseOutput) StatusMessages() StatusMessageResponseArrayOutput {
	return o.ApplyT(func(v LabelingJobResponse) []StatusMessageResponse { return v.StatusMessages }).(StatusMessageResponseArrayOutput)
}

func (o LabelingJobResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabelingJobResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type LabelingJobTextProperties struct {
	AnnotationType *string `pulumi:"annotationType"`
	MediaType      string  `pulumi:"mediaType"`
}

type LabelingJobTextPropertiesResponse struct {
	AnnotationType *string `pulumi:"annotationType"`
	MediaType      string  `pulumi:"mediaType"`
}

type LinkedInfo struct {
	LinkedId           *string `pulumi:"linkedId"`
	LinkedResourceName *string `pulumi:"linkedResourceName"`
	Origin             *string `pulumi:"origin"`
}





type LinkedInfoInput interface {
	pulumi.Input

	ToLinkedInfoOutput() LinkedInfoOutput
	ToLinkedInfoOutputWithContext(context.Context) LinkedInfoOutput
}

type LinkedInfoArgs struct {
	LinkedId           pulumi.StringPtrInput `pulumi:"linkedId"`
	LinkedResourceName pulumi.StringPtrInput `pulumi:"linkedResourceName"`
	Origin             pulumi.StringPtrInput `pulumi:"origin"`
}

func (LinkedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedInfo)(nil)).Elem()
}

func (i LinkedInfoArgs) ToLinkedInfoOutput() LinkedInfoOutput {
	return i.ToLinkedInfoOutputWithContext(context.Background())
}

func (i LinkedInfoArgs) ToLinkedInfoOutputWithContext(ctx context.Context) LinkedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedInfoOutput)
}

func (i LinkedInfoArgs) ToLinkedInfoPtrOutput() LinkedInfoPtrOutput {
	return i.ToLinkedInfoPtrOutputWithContext(context.Background())
}

func (i LinkedInfoArgs) ToLinkedInfoPtrOutputWithContext(ctx context.Context) LinkedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedInfoOutput).ToLinkedInfoPtrOutputWithContext(ctx)
}









type LinkedInfoPtrInput interface {
	pulumi.Input

	ToLinkedInfoPtrOutput() LinkedInfoPtrOutput
	ToLinkedInfoPtrOutputWithContext(context.Context) LinkedInfoPtrOutput
}

type linkedInfoPtrType LinkedInfoArgs

func LinkedInfoPtr(v *LinkedInfoArgs) LinkedInfoPtrInput {
	return (*linkedInfoPtrType)(v)
}

func (*linkedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedInfo)(nil)).Elem()
}

func (i *linkedInfoPtrType) ToLinkedInfoPtrOutput() LinkedInfoPtrOutput {
	return i.ToLinkedInfoPtrOutputWithContext(context.Background())
}

func (i *linkedInfoPtrType) ToLinkedInfoPtrOutputWithContext(ctx context.Context) LinkedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedInfoPtrOutput)
}

type LinkedInfoOutput struct{ *pulumi.OutputState }

func (LinkedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedInfo)(nil)).Elem()
}

func (o LinkedInfoOutput) ToLinkedInfoOutput() LinkedInfoOutput {
	return o
}

func (o LinkedInfoOutput) ToLinkedInfoOutputWithContext(ctx context.Context) LinkedInfoOutput {
	return o
}

func (o LinkedInfoOutput) ToLinkedInfoPtrOutput() LinkedInfoPtrOutput {
	return o.ToLinkedInfoPtrOutputWithContext(context.Background())
}

func (o LinkedInfoOutput) ToLinkedInfoPtrOutputWithContext(ctx context.Context) LinkedInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinkedInfo) *LinkedInfo {
		return &v
	}).(LinkedInfoPtrOutput)
}

func (o LinkedInfoOutput) LinkedId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedInfo) *string { return v.LinkedId }).(pulumi.StringPtrOutput)
}

func (o LinkedInfoOutput) LinkedResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedInfo) *string { return v.LinkedResourceName }).(pulumi.StringPtrOutput)
}

func (o LinkedInfoOutput) Origin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedInfo) *string { return v.Origin }).(pulumi.StringPtrOutput)
}

type LinkedInfoPtrOutput struct{ *pulumi.OutputState }

func (LinkedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedInfo)(nil)).Elem()
}

func (o LinkedInfoPtrOutput) ToLinkedInfoPtrOutput() LinkedInfoPtrOutput {
	return o
}

func (o LinkedInfoPtrOutput) ToLinkedInfoPtrOutputWithContext(ctx context.Context) LinkedInfoPtrOutput {
	return o
}

func (o LinkedInfoPtrOutput) Elem() LinkedInfoOutput {
	return o.ApplyT(func(v *LinkedInfo) LinkedInfo {
		if v != nil {
			return *v
		}
		var ret LinkedInfo
		return ret
	}).(LinkedInfoOutput)
}

func (o LinkedInfoPtrOutput) LinkedId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedInfo) *string {
		if v == nil {
			return nil
		}
		return v.LinkedId
	}).(pulumi.StringPtrOutput)
}

func (o LinkedInfoPtrOutput) LinkedResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedInfo) *string {
		if v == nil {
			return nil
		}
		return v.LinkedResourceName
	}).(pulumi.StringPtrOutput)
}

func (o LinkedInfoPtrOutput) Origin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedInfo) *string {
		if v == nil {
			return nil
		}
		return v.Origin
	}).(pulumi.StringPtrOutput)
}

type LinkedInfoResponse struct {
	LinkedId           *string `pulumi:"linkedId"`
	LinkedResourceName *string `pulumi:"linkedResourceName"`
	Origin             *string `pulumi:"origin"`
}

type LinkedInfoResponseOutput struct{ *pulumi.OutputState }

func (LinkedInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedInfoResponse)(nil)).Elem()
}

func (o LinkedInfoResponseOutput) ToLinkedInfoResponseOutput() LinkedInfoResponseOutput {
	return o
}

func (o LinkedInfoResponseOutput) ToLinkedInfoResponseOutputWithContext(ctx context.Context) LinkedInfoResponseOutput {
	return o
}

func (o LinkedInfoResponseOutput) LinkedId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedInfoResponse) *string { return v.LinkedId }).(pulumi.StringPtrOutput)
}

func (o LinkedInfoResponseOutput) LinkedResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedInfoResponse) *string { return v.LinkedResourceName }).(pulumi.StringPtrOutput)
}

func (o LinkedInfoResponseOutput) Origin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedInfoResponse) *string { return v.Origin }).(pulumi.StringPtrOutput)
}

type LinkedInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (LinkedInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedInfoResponse)(nil)).Elem()
}

func (o LinkedInfoResponsePtrOutput) ToLinkedInfoResponsePtrOutput() LinkedInfoResponsePtrOutput {
	return o
}

func (o LinkedInfoResponsePtrOutput) ToLinkedInfoResponsePtrOutputWithContext(ctx context.Context) LinkedInfoResponsePtrOutput {
	return o
}

func (o LinkedInfoResponsePtrOutput) Elem() LinkedInfoResponseOutput {
	return o.ApplyT(func(v *LinkedInfoResponse) LinkedInfoResponse {
		if v != nil {
			return *v
		}
		var ret LinkedInfoResponse
		return ret
	}).(LinkedInfoResponseOutput)
}

func (o LinkedInfoResponsePtrOutput) LinkedId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LinkedId
	}).(pulumi.StringPtrOutput)
}

func (o LinkedInfoResponsePtrOutput) LinkedResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LinkedResourceName
	}).(pulumi.StringPtrOutput)
}

func (o LinkedInfoResponsePtrOutput) Origin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Origin
	}).(pulumi.StringPtrOutput)
}

type ListNotebookKeysResultResponse struct {
	PrimaryAccessKey   string `pulumi:"primaryAccessKey"`
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
}

type MLAssistConfiguration struct {
	InferencingComputeBinding *ComputeConfiguration `pulumi:"inferencingComputeBinding"`
	MlAssistEnabled           *bool                 `pulumi:"mlAssistEnabled"`
	TrainingComputeBinding    *ComputeConfiguration `pulumi:"trainingComputeBinding"`
}





type MLAssistConfigurationInput interface {
	pulumi.Input

	ToMLAssistConfigurationOutput() MLAssistConfigurationOutput
	ToMLAssistConfigurationOutputWithContext(context.Context) MLAssistConfigurationOutput
}

type MLAssistConfigurationArgs struct {
	InferencingComputeBinding ComputeConfigurationPtrInput `pulumi:"inferencingComputeBinding"`
	MlAssistEnabled           pulumi.BoolPtrInput          `pulumi:"mlAssistEnabled"`
	TrainingComputeBinding    ComputeConfigurationPtrInput `pulumi:"trainingComputeBinding"`
}

func (MLAssistConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MLAssistConfiguration)(nil)).Elem()
}

func (i MLAssistConfigurationArgs) ToMLAssistConfigurationOutput() MLAssistConfigurationOutput {
	return i.ToMLAssistConfigurationOutputWithContext(context.Background())
}

func (i MLAssistConfigurationArgs) ToMLAssistConfigurationOutputWithContext(ctx context.Context) MLAssistConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLAssistConfigurationOutput)
}

func (i MLAssistConfigurationArgs) ToMLAssistConfigurationPtrOutput() MLAssistConfigurationPtrOutput {
	return i.ToMLAssistConfigurationPtrOutputWithContext(context.Background())
}

func (i MLAssistConfigurationArgs) ToMLAssistConfigurationPtrOutputWithContext(ctx context.Context) MLAssistConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLAssistConfigurationOutput).ToMLAssistConfigurationPtrOutputWithContext(ctx)
}









type MLAssistConfigurationPtrInput interface {
	pulumi.Input

	ToMLAssistConfigurationPtrOutput() MLAssistConfigurationPtrOutput
	ToMLAssistConfigurationPtrOutputWithContext(context.Context) MLAssistConfigurationPtrOutput
}

type mlassistConfigurationPtrType MLAssistConfigurationArgs

func MLAssistConfigurationPtr(v *MLAssistConfigurationArgs) MLAssistConfigurationPtrInput {
	return (*mlassistConfigurationPtrType)(v)
}

func (*mlassistConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MLAssistConfiguration)(nil)).Elem()
}

func (i *mlassistConfigurationPtrType) ToMLAssistConfigurationPtrOutput() MLAssistConfigurationPtrOutput {
	return i.ToMLAssistConfigurationPtrOutputWithContext(context.Background())
}

func (i *mlassistConfigurationPtrType) ToMLAssistConfigurationPtrOutputWithContext(ctx context.Context) MLAssistConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLAssistConfigurationPtrOutput)
}

type MLAssistConfigurationOutput struct{ *pulumi.OutputState }

func (MLAssistConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MLAssistConfiguration)(nil)).Elem()
}

func (o MLAssistConfigurationOutput) ToMLAssistConfigurationOutput() MLAssistConfigurationOutput {
	return o
}

func (o MLAssistConfigurationOutput) ToMLAssistConfigurationOutputWithContext(ctx context.Context) MLAssistConfigurationOutput {
	return o
}

func (o MLAssistConfigurationOutput) ToMLAssistConfigurationPtrOutput() MLAssistConfigurationPtrOutput {
	return o.ToMLAssistConfigurationPtrOutputWithContext(context.Background())
}

func (o MLAssistConfigurationOutput) ToMLAssistConfigurationPtrOutputWithContext(ctx context.Context) MLAssistConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MLAssistConfiguration) *MLAssistConfiguration {
		return &v
	}).(MLAssistConfigurationPtrOutput)
}

func (o MLAssistConfigurationOutput) InferencingComputeBinding() ComputeConfigurationPtrOutput {
	return o.ApplyT(func(v MLAssistConfiguration) *ComputeConfiguration { return v.InferencingComputeBinding }).(ComputeConfigurationPtrOutput)
}

func (o MLAssistConfigurationOutput) MlAssistEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MLAssistConfiguration) *bool { return v.MlAssistEnabled }).(pulumi.BoolPtrOutput)
}

func (o MLAssistConfigurationOutput) TrainingComputeBinding() ComputeConfigurationPtrOutput {
	return o.ApplyT(func(v MLAssistConfiguration) *ComputeConfiguration { return v.TrainingComputeBinding }).(ComputeConfigurationPtrOutput)
}

type MLAssistConfigurationPtrOutput struct{ *pulumi.OutputState }

func (MLAssistConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MLAssistConfiguration)(nil)).Elem()
}

func (o MLAssistConfigurationPtrOutput) ToMLAssistConfigurationPtrOutput() MLAssistConfigurationPtrOutput {
	return o
}

func (o MLAssistConfigurationPtrOutput) ToMLAssistConfigurationPtrOutputWithContext(ctx context.Context) MLAssistConfigurationPtrOutput {
	return o
}

func (o MLAssistConfigurationPtrOutput) Elem() MLAssistConfigurationOutput {
	return o.ApplyT(func(v *MLAssistConfiguration) MLAssistConfiguration {
		if v != nil {
			return *v
		}
		var ret MLAssistConfiguration
		return ret
	}).(MLAssistConfigurationOutput)
}

func (o MLAssistConfigurationPtrOutput) InferencingComputeBinding() ComputeConfigurationPtrOutput {
	return o.ApplyT(func(v *MLAssistConfiguration) *ComputeConfiguration {
		if v == nil {
			return nil
		}
		return v.InferencingComputeBinding
	}).(ComputeConfigurationPtrOutput)
}

func (o MLAssistConfigurationPtrOutput) MlAssistEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MLAssistConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.MlAssistEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o MLAssistConfigurationPtrOutput) TrainingComputeBinding() ComputeConfigurationPtrOutput {
	return o.ApplyT(func(v *MLAssistConfiguration) *ComputeConfiguration {
		if v == nil {
			return nil
		}
		return v.TrainingComputeBinding
	}).(ComputeConfigurationPtrOutput)
}

type MLAssistConfigurationResponse struct {
	InferencingComputeBinding *ComputeConfigurationResponse `pulumi:"inferencingComputeBinding"`
	MlAssistEnabled           *bool                         `pulumi:"mlAssistEnabled"`
	TrainingComputeBinding    *ComputeConfigurationResponse `pulumi:"trainingComputeBinding"`
}

type MLAssistConfigurationResponseOutput struct{ *pulumi.OutputState }

func (MLAssistConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MLAssistConfigurationResponse)(nil)).Elem()
}

func (o MLAssistConfigurationResponseOutput) ToMLAssistConfigurationResponseOutput() MLAssistConfigurationResponseOutput {
	return o
}

func (o MLAssistConfigurationResponseOutput) ToMLAssistConfigurationResponseOutputWithContext(ctx context.Context) MLAssistConfigurationResponseOutput {
	return o
}

func (o MLAssistConfigurationResponseOutput) InferencingComputeBinding() ComputeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v MLAssistConfigurationResponse) *ComputeConfigurationResponse {
		return v.InferencingComputeBinding
	}).(ComputeConfigurationResponsePtrOutput)
}

func (o MLAssistConfigurationResponseOutput) MlAssistEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MLAssistConfigurationResponse) *bool { return v.MlAssistEnabled }).(pulumi.BoolPtrOutput)
}

func (o MLAssistConfigurationResponseOutput) TrainingComputeBinding() ComputeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v MLAssistConfigurationResponse) *ComputeConfigurationResponse { return v.TrainingComputeBinding }).(ComputeConfigurationResponsePtrOutput)
}

type MLAssistConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (MLAssistConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MLAssistConfigurationResponse)(nil)).Elem()
}

func (o MLAssistConfigurationResponsePtrOutput) ToMLAssistConfigurationResponsePtrOutput() MLAssistConfigurationResponsePtrOutput {
	return o
}

func (o MLAssistConfigurationResponsePtrOutput) ToMLAssistConfigurationResponsePtrOutputWithContext(ctx context.Context) MLAssistConfigurationResponsePtrOutput {
	return o
}

func (o MLAssistConfigurationResponsePtrOutput) Elem() MLAssistConfigurationResponseOutput {
	return o.ApplyT(func(v *MLAssistConfigurationResponse) MLAssistConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret MLAssistConfigurationResponse
		return ret
	}).(MLAssistConfigurationResponseOutput)
}

func (o MLAssistConfigurationResponsePtrOutput) InferencingComputeBinding() ComputeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *MLAssistConfigurationResponse) *ComputeConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.InferencingComputeBinding
	}).(ComputeConfigurationResponsePtrOutput)
}

func (o MLAssistConfigurationResponsePtrOutput) MlAssistEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MLAssistConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.MlAssistEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o MLAssistConfigurationResponsePtrOutput) TrainingComputeBinding() ComputeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *MLAssistConfigurationResponse) *ComputeConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.TrainingComputeBinding
	}).(ComputeConfigurationResponsePtrOutput)
}

type ManagedIdentity struct {
	ClientId     *string `pulumi:"clientId"`
	IdentityType string  `pulumi:"identityType"`
	ObjectId     *string `pulumi:"objectId"`
	ResourceId   *string `pulumi:"resourceId"`
}

type ManagedIdentityResponse struct {
	ClientId     *string `pulumi:"clientId"`
	IdentityType string  `pulumi:"identityType"`
	ObjectId     *string `pulumi:"objectId"`
	ResourceId   *string `pulumi:"resourceId"`
}

type ManagedOnlineDeployment struct {
	AppInsightsEnabled   *bool                  `pulumi:"appInsightsEnabled"`
	CodeConfiguration    *CodeConfiguration     `pulumi:"codeConfiguration"`
	Description          *string                `pulumi:"description"`
	EndpointComputeType  string                 `pulumi:"endpointComputeType"`
	EnvironmentId        *string                `pulumi:"environmentId"`
	EnvironmentVariables map[string]string      `pulumi:"environmentVariables"`
	InstanceType         *string                `pulumi:"instanceType"`
	LivenessProbe        *ProbeSettings         `pulumi:"livenessProbe"`
	Model                interface{}            `pulumi:"model"`
	Properties           map[string]string      `pulumi:"properties"`
	ReadinessProbe       *ProbeSettings         `pulumi:"readinessProbe"`
	RequestSettings      *OnlineRequestSettings `pulumi:"requestSettings"`
	ScaleSettings        interface{}            `pulumi:"scaleSettings"`
}

type ManagedOnlineDeploymentResponse struct {
	AppInsightsEnabled   *bool                          `pulumi:"appInsightsEnabled"`
	CodeConfiguration    *CodeConfigurationResponse     `pulumi:"codeConfiguration"`
	Description          *string                        `pulumi:"description"`
	EndpointComputeType  string                         `pulumi:"endpointComputeType"`
	EnvironmentId        *string                        `pulumi:"environmentId"`
	EnvironmentVariables map[string]string              `pulumi:"environmentVariables"`
	InstanceType         *string                        `pulumi:"instanceType"`
	LivenessProbe        *ProbeSettingsResponse         `pulumi:"livenessProbe"`
	Model                interface{}                    `pulumi:"model"`
	Properties           map[string]string              `pulumi:"properties"`
	ProvisioningState    string                         `pulumi:"provisioningState"`
	ReadinessProbe       *ProbeSettingsResponse         `pulumi:"readinessProbe"`
	RequestSettings      *OnlineRequestSettingsResponse `pulumi:"requestSettings"`
	ScaleSettings        interface{}                    `pulumi:"scaleSettings"`
}

type ManualScaleSettings struct {
	InstanceCount *int   `pulumi:"instanceCount"`
	MaxInstances  *int   `pulumi:"maxInstances"`
	MinInstances  *int   `pulumi:"minInstances"`
	ScaleType     string `pulumi:"scaleType"`
}

type ManualScaleSettingsResponse struct {
	InstanceCount *int   `pulumi:"instanceCount"`
	MaxInstances  *int   `pulumi:"maxInstances"`
	MinInstances  *int   `pulumi:"minInstances"`
	ScaleType     string `pulumi:"scaleType"`
}

type MedianStoppingPolicy struct {
	DelayEvaluation    *int   `pulumi:"delayEvaluation"`
	EvaluationInterval *int   `pulumi:"evaluationInterval"`
	PolicyType         string `pulumi:"policyType"`
}

type MedianStoppingPolicyResponse struct {
	DelayEvaluation    *int   `pulumi:"delayEvaluation"`
	EvaluationInterval *int   `pulumi:"evaluationInterval"`
	PolicyType         string `pulumi:"policyType"`
}

type ModelContainerType struct {
	Description *string           `pulumi:"description"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}





type ModelContainerTypeInput interface {
	pulumi.Input

	ToModelContainerTypeOutput() ModelContainerTypeOutput
	ToModelContainerTypeOutputWithContext(context.Context) ModelContainerTypeOutput
}

type ModelContainerTypeArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}

func (ModelContainerTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelContainerType)(nil)).Elem()
}

func (i ModelContainerTypeArgs) ToModelContainerTypeOutput() ModelContainerTypeOutput {
	return i.ToModelContainerTypeOutputWithContext(context.Background())
}

func (i ModelContainerTypeArgs) ToModelContainerTypeOutputWithContext(ctx context.Context) ModelContainerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelContainerTypeOutput)
}

type ModelContainerTypeOutput struct{ *pulumi.OutputState }

func (ModelContainerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelContainerType)(nil)).Elem()
}

func (o ModelContainerTypeOutput) ToModelContainerTypeOutput() ModelContainerTypeOutput {
	return o
}

func (o ModelContainerTypeOutput) ToModelContainerTypeOutputWithContext(ctx context.Context) ModelContainerTypeOutput {
	return o
}

func (o ModelContainerTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelContainerType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ModelContainerTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelContainerType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ModelContainerTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelContainerType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ModelContainerResponse struct {
	Description *string           `pulumi:"description"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}

type ModelContainerResponseOutput struct{ *pulumi.OutputState }

func (ModelContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelContainerResponse)(nil)).Elem()
}

func (o ModelContainerResponseOutput) ToModelContainerResponseOutput() ModelContainerResponseOutput {
	return o
}

func (o ModelContainerResponseOutput) ToModelContainerResponseOutputWithContext(ctx context.Context) ModelContainerResponseOutput {
	return o
}

func (o ModelContainerResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelContainerResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ModelContainerResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelContainerResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ModelContainerResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelContainerResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ModelVersionType struct {
	DatastoreId *string               `pulumi:"datastoreId"`
	Description *string               `pulumi:"description"`
	Flavors     map[string]FlavorData `pulumi:"flavors"`
	IsAnonymous *bool                 `pulumi:"isAnonymous"`
	Path        string                `pulumi:"path"`
	Properties  map[string]string     `pulumi:"properties"`
	Tags        map[string]string     `pulumi:"tags"`
}





type ModelVersionTypeInput interface {
	pulumi.Input

	ToModelVersionTypeOutput() ModelVersionTypeOutput
	ToModelVersionTypeOutputWithContext(context.Context) ModelVersionTypeOutput
}

type ModelVersionTypeArgs struct {
	DatastoreId pulumi.StringPtrInput `pulumi:"datastoreId"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	Flavors     FlavorDataMapInput    `pulumi:"flavors"`
	IsAnonymous pulumi.BoolPtrInput   `pulumi:"isAnonymous"`
	Path        pulumi.StringInput    `pulumi:"path"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}

func (ModelVersionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelVersionType)(nil)).Elem()
}

func (i ModelVersionTypeArgs) ToModelVersionTypeOutput() ModelVersionTypeOutput {
	return i.ToModelVersionTypeOutputWithContext(context.Background())
}

func (i ModelVersionTypeArgs) ToModelVersionTypeOutputWithContext(ctx context.Context) ModelVersionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelVersionTypeOutput)
}

type ModelVersionTypeOutput struct{ *pulumi.OutputState }

func (ModelVersionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelVersionType)(nil)).Elem()
}

func (o ModelVersionTypeOutput) ToModelVersionTypeOutput() ModelVersionTypeOutput {
	return o
}

func (o ModelVersionTypeOutput) ToModelVersionTypeOutputWithContext(ctx context.Context) ModelVersionTypeOutput {
	return o
}

func (o ModelVersionTypeOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionType) *string { return v.DatastoreId }).(pulumi.StringPtrOutput)
}

func (o ModelVersionTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ModelVersionTypeOutput) Flavors() FlavorDataMapOutput {
	return o.ApplyT(func(v ModelVersionType) map[string]FlavorData { return v.Flavors }).(FlavorDataMapOutput)
}

func (o ModelVersionTypeOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelVersionType) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o ModelVersionTypeOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v ModelVersionType) string { return v.Path }).(pulumi.StringOutput)
}

func (o ModelVersionTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelVersionType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ModelVersionTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelVersionType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ModelVersionResponse struct {
	DatastoreId *string                       `pulumi:"datastoreId"`
	Description *string                       `pulumi:"description"`
	Flavors     map[string]FlavorDataResponse `pulumi:"flavors"`
	IsAnonymous *bool                         `pulumi:"isAnonymous"`
	Path        string                        `pulumi:"path"`
	Properties  map[string]string             `pulumi:"properties"`
	Tags        map[string]string             `pulumi:"tags"`
}

type ModelVersionResponseOutput struct{ *pulumi.OutputState }

func (ModelVersionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelVersionResponse)(nil)).Elem()
}

func (o ModelVersionResponseOutput) ToModelVersionResponseOutput() ModelVersionResponseOutput {
	return o
}

func (o ModelVersionResponseOutput) ToModelVersionResponseOutputWithContext(ctx context.Context) ModelVersionResponseOutput {
	return o
}

func (o ModelVersionResponseOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionResponse) *string { return v.DatastoreId }).(pulumi.StringPtrOutput)
}

func (o ModelVersionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ModelVersionResponseOutput) Flavors() FlavorDataResponseMapOutput {
	return o.ApplyT(func(v ModelVersionResponse) map[string]FlavorDataResponse { return v.Flavors }).(FlavorDataResponseMapOutput)
}

func (o ModelVersionResponseOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelVersionResponse) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o ModelVersionResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v ModelVersionResponse) string { return v.Path }).(pulumi.StringOutput)
}

func (o ModelVersionResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelVersionResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ModelVersionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelVersionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type Mpi struct {
	DistributionType        string `pulumi:"distributionType"`
	ProcessCountPerInstance *int   `pulumi:"processCountPerInstance"`
}

type MpiResponse struct {
	DistributionType        string `pulumi:"distributionType"`
	ProcessCountPerInstance *int   `pulumi:"processCountPerInstance"`
}

type NodeStateCountsResponse struct {
	IdleNodeCount      int `pulumi:"idleNodeCount"`
	LeavingNodeCount   int `pulumi:"leavingNodeCount"`
	PreemptedNodeCount int `pulumi:"preemptedNodeCount"`
	PreparingNodeCount int `pulumi:"preparingNodeCount"`
	RunningNodeCount   int `pulumi:"runningNodeCount"`
	UnusableNodeCount  int `pulumi:"unusableNodeCount"`
}

type NoneDatastoreCredentials struct {
	CredentialsType string                `pulumi:"credentialsType"`
	Secrets         *NoneDatastoreSecrets `pulumi:"secrets"`
}

type NoneDatastoreCredentialsResponse struct {
	CredentialsType string `pulumi:"credentialsType"`
}

type NoneDatastoreSecrets struct {
	SecretsType string `pulumi:"secretsType"`
}

type NotebookPreparationErrorResponse struct {
	ErrorMessage *string `pulumi:"errorMessage"`
	StatusCode   *int    `pulumi:"statusCode"`
}

type NotebookPreparationErrorResponseOutput struct{ *pulumi.OutputState }

func (NotebookPreparationErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookPreparationErrorResponse)(nil)).Elem()
}

func (o NotebookPreparationErrorResponseOutput) ToNotebookPreparationErrorResponseOutput() NotebookPreparationErrorResponseOutput {
	return o
}

func (o NotebookPreparationErrorResponseOutput) ToNotebookPreparationErrorResponseOutputWithContext(ctx context.Context) NotebookPreparationErrorResponseOutput {
	return o
}

func (o NotebookPreparationErrorResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookPreparationErrorResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o NotebookPreparationErrorResponseOutput) StatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NotebookPreparationErrorResponse) *int { return v.StatusCode }).(pulumi.IntPtrOutput)
}

type NotebookPreparationErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (NotebookPreparationErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookPreparationErrorResponse)(nil)).Elem()
}

func (o NotebookPreparationErrorResponsePtrOutput) ToNotebookPreparationErrorResponsePtrOutput() NotebookPreparationErrorResponsePtrOutput {
	return o
}

func (o NotebookPreparationErrorResponsePtrOutput) ToNotebookPreparationErrorResponsePtrOutputWithContext(ctx context.Context) NotebookPreparationErrorResponsePtrOutput {
	return o
}

func (o NotebookPreparationErrorResponsePtrOutput) Elem() NotebookPreparationErrorResponseOutput {
	return o.ApplyT(func(v *NotebookPreparationErrorResponse) NotebookPreparationErrorResponse {
		if v != nil {
			return *v
		}
		var ret NotebookPreparationErrorResponse
		return ret
	}).(NotebookPreparationErrorResponseOutput)
}

func (o NotebookPreparationErrorResponsePtrOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookPreparationErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.ErrorMessage
	}).(pulumi.StringPtrOutput)
}

func (o NotebookPreparationErrorResponsePtrOutput) StatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NotebookPreparationErrorResponse) *int {
		if v == nil {
			return nil
		}
		return v.StatusCode
	}).(pulumi.IntPtrOutput)
}

type NotebookResourceInfoResponse struct {
	Fqdn                     *string                           `pulumi:"fqdn"`
	NotebookPreparationError *NotebookPreparationErrorResponse `pulumi:"notebookPreparationError"`
	ResourceId               *string                           `pulumi:"resourceId"`
}

type NotebookResourceInfoResponseOutput struct{ *pulumi.OutputState }

func (NotebookResourceInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookResourceInfoResponse)(nil)).Elem()
}

func (o NotebookResourceInfoResponseOutput) ToNotebookResourceInfoResponseOutput() NotebookResourceInfoResponseOutput {
	return o
}

func (o NotebookResourceInfoResponseOutput) ToNotebookResourceInfoResponseOutputWithContext(ctx context.Context) NotebookResourceInfoResponseOutput {
	return o
}

func (o NotebookResourceInfoResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceInfoResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceInfoResponseOutput) NotebookPreparationError() NotebookPreparationErrorResponsePtrOutput {
	return o.ApplyT(func(v NotebookResourceInfoResponse) *NotebookPreparationErrorResponse {
		return v.NotebookPreparationError
	}).(NotebookPreparationErrorResponsePtrOutput)
}

func (o NotebookResourceInfoResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceInfoResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type Objective struct {
	Goal          string `pulumi:"goal"`
	PrimaryMetric string `pulumi:"primaryMetric"`
}

type ObjectiveResponse struct {
	Goal          string `pulumi:"goal"`
	PrimaryMetric string `pulumi:"primaryMetric"`
}

type OnlineEndpointType struct {
	AuthMode    string            `pulumi:"authMode"`
	Description *string           `pulumi:"description"`
	Keys        *EndpointAuthKeys `pulumi:"keys"`
	Properties  map[string]string `pulumi:"properties"`
	Target      *string           `pulumi:"target"`
	Traffic     map[string]int    `pulumi:"traffic"`
}





type OnlineEndpointTypeInput interface {
	pulumi.Input

	ToOnlineEndpointTypeOutput() OnlineEndpointTypeOutput
	ToOnlineEndpointTypeOutputWithContext(context.Context) OnlineEndpointTypeOutput
}

type OnlineEndpointTypeArgs struct {
	AuthMode    pulumi.StringInput       `pulumi:"authMode"`
	Description pulumi.StringPtrInput    `pulumi:"description"`
	Keys        EndpointAuthKeysPtrInput `pulumi:"keys"`
	Properties  pulumi.StringMapInput    `pulumi:"properties"`
	Target      pulumi.StringPtrInput    `pulumi:"target"`
	Traffic     pulumi.IntMapInput       `pulumi:"traffic"`
}

func (OnlineEndpointTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnlineEndpointType)(nil)).Elem()
}

func (i OnlineEndpointTypeArgs) ToOnlineEndpointTypeOutput() OnlineEndpointTypeOutput {
	return i.ToOnlineEndpointTypeOutputWithContext(context.Background())
}

func (i OnlineEndpointTypeArgs) ToOnlineEndpointTypeOutputWithContext(ctx context.Context) OnlineEndpointTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnlineEndpointTypeOutput)
}

type OnlineEndpointTypeOutput struct{ *pulumi.OutputState }

func (OnlineEndpointTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnlineEndpointType)(nil)).Elem()
}

func (o OnlineEndpointTypeOutput) ToOnlineEndpointTypeOutput() OnlineEndpointTypeOutput {
	return o
}

func (o OnlineEndpointTypeOutput) ToOnlineEndpointTypeOutputWithContext(ctx context.Context) OnlineEndpointTypeOutput {
	return o
}

func (o OnlineEndpointTypeOutput) AuthMode() pulumi.StringOutput {
	return o.ApplyT(func(v OnlineEndpointType) string { return v.AuthMode }).(pulumi.StringOutput)
}

func (o OnlineEndpointTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnlineEndpointType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o OnlineEndpointTypeOutput) Keys() EndpointAuthKeysPtrOutput {
	return o.ApplyT(func(v OnlineEndpointType) *EndpointAuthKeys { return v.Keys }).(EndpointAuthKeysPtrOutput)
}

func (o OnlineEndpointTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v OnlineEndpointType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o OnlineEndpointTypeOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnlineEndpointType) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o OnlineEndpointTypeOutput) Traffic() pulumi.IntMapOutput {
	return o.ApplyT(func(v OnlineEndpointType) map[string]int { return v.Traffic }).(pulumi.IntMapOutput)
}

type OnlineEndpointResponse struct {
	AuthMode          string            `pulumi:"authMode"`
	Description       *string           `pulumi:"description"`
	Properties        map[string]string `pulumi:"properties"`
	ProvisioningState string            `pulumi:"provisioningState"`
	ScoringUri        string            `pulumi:"scoringUri"`
	SwaggerUri        string            `pulumi:"swaggerUri"`
	Target            *string           `pulumi:"target"`
	Traffic           map[string]int    `pulumi:"traffic"`
}

type OnlineEndpointResponseOutput struct{ *pulumi.OutputState }

func (OnlineEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnlineEndpointResponse)(nil)).Elem()
}

func (o OnlineEndpointResponseOutput) ToOnlineEndpointResponseOutput() OnlineEndpointResponseOutput {
	return o
}

func (o OnlineEndpointResponseOutput) ToOnlineEndpointResponseOutputWithContext(ctx context.Context) OnlineEndpointResponseOutput {
	return o
}

func (o OnlineEndpointResponseOutput) AuthMode() pulumi.StringOutput {
	return o.ApplyT(func(v OnlineEndpointResponse) string { return v.AuthMode }).(pulumi.StringOutput)
}

func (o OnlineEndpointResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnlineEndpointResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o OnlineEndpointResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v OnlineEndpointResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o OnlineEndpointResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v OnlineEndpointResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o OnlineEndpointResponseOutput) ScoringUri() pulumi.StringOutput {
	return o.ApplyT(func(v OnlineEndpointResponse) string { return v.ScoringUri }).(pulumi.StringOutput)
}

func (o OnlineEndpointResponseOutput) SwaggerUri() pulumi.StringOutput {
	return o.ApplyT(func(v OnlineEndpointResponse) string { return v.SwaggerUri }).(pulumi.StringOutput)
}

func (o OnlineEndpointResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnlineEndpointResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o OnlineEndpointResponseOutput) Traffic() pulumi.IntMapOutput {
	return o.ApplyT(func(v OnlineEndpointResponse) map[string]int { return v.Traffic }).(pulumi.IntMapOutput)
}

type OnlineRequestSettings struct {
	MaxConcurrentRequestsPerInstance *int    `pulumi:"maxConcurrentRequestsPerInstance"`
	MaxQueueWait                     *string `pulumi:"maxQueueWait"`
	RequestTimeout                   *string `pulumi:"requestTimeout"`
}

type OnlineRequestSettingsResponse struct {
	MaxConcurrentRequestsPerInstance *int    `pulumi:"maxConcurrentRequestsPerInstance"`
	MaxQueueWait                     *string `pulumi:"maxQueueWait"`
	RequestTimeout                   *string `pulumi:"requestTimeout"`
}

type OutputDataBinding struct {
	DatastoreId     *string `pulumi:"datastoreId"`
	Mode            *string `pulumi:"mode"`
	PathOnCompute   *string `pulumi:"pathOnCompute"`
	PathOnDatastore *string `pulumi:"pathOnDatastore"`
}

type OutputDataBindingResponse struct {
	DatastoreId     *string `pulumi:"datastoreId"`
	Mode            *string `pulumi:"mode"`
	PathOnCompute   *string `pulumi:"pathOnCompute"`
	PathOnDatastore *string `pulumi:"pathOnDatastore"`
}

type OutputPathAssetReference struct {
	JobId         *string `pulumi:"jobId"`
	Path          *string `pulumi:"path"`
	ReferenceType string  `pulumi:"referenceType"`
}

type OutputPathAssetReferenceResponse struct {
	JobId         *string `pulumi:"jobId"`
	Path          *string `pulumi:"path"`
	ReferenceType string  `pulumi:"referenceType"`
}

type PasswordResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type PersonalComputeInstanceSettings struct {
	AssignedUser *AssignedUser `pulumi:"assignedUser"`
}

type PersonalComputeInstanceSettingsResponse struct {
	AssignedUser *AssignedUserResponse `pulumi:"assignedUser"`
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Identity                          *IdentityResponse                         `pulumi:"identity"`
	Location                          *string                                   `pulumi:"location"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Sku                               *SkuResponse                              `pulumi:"sku"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Tags                              map[string]string                         `pulumi:"tags"`
	Type                              string                                    `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id          string `pulumi:"id"`
	SubnetArmId string `pulumi:"subnetArmId"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointResponseOutput) SubnetArmId() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.SubnetArmId }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointResponsePtrOutput) SubnetArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SubnetArmId
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ProbeSettings struct {
	FailureThreshold *int    `pulumi:"failureThreshold"`
	InitialDelay     *string `pulumi:"initialDelay"`
	Period           *string `pulumi:"period"`
	SuccessThreshold *int    `pulumi:"successThreshold"`
	Timeout          *string `pulumi:"timeout"`
}

type ProbeSettingsResponse struct {
	FailureThreshold *int    `pulumi:"failureThreshold"`
	InitialDelay     *string `pulumi:"initialDelay"`
	Period           *string `pulumi:"period"`
	SuccessThreshold *int    `pulumi:"successThreshold"`
	Timeout          *string `pulumi:"timeout"`
}

type ProgressMetricsResponse struct {
	CompletedDatapointCount           float64 `pulumi:"completedDatapointCount"`
	IncrementalDatasetLastRefreshTime string  `pulumi:"incrementalDatasetLastRefreshTime"`
	SkippedDatapointCount             float64 `pulumi:"skippedDatapointCount"`
	TotalDatapointCount               float64 `pulumi:"totalDatapointCount"`
}

type ProgressMetricsResponseOutput struct{ *pulumi.OutputState }

func (ProgressMetricsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProgressMetricsResponse)(nil)).Elem()
}

func (o ProgressMetricsResponseOutput) ToProgressMetricsResponseOutput() ProgressMetricsResponseOutput {
	return o
}

func (o ProgressMetricsResponseOutput) ToProgressMetricsResponseOutputWithContext(ctx context.Context) ProgressMetricsResponseOutput {
	return o
}

func (o ProgressMetricsResponseOutput) CompletedDatapointCount() pulumi.Float64Output {
	return o.ApplyT(func(v ProgressMetricsResponse) float64 { return v.CompletedDatapointCount }).(pulumi.Float64Output)
}

func (o ProgressMetricsResponseOutput) IncrementalDatasetLastRefreshTime() pulumi.StringOutput {
	return o.ApplyT(func(v ProgressMetricsResponse) string { return v.IncrementalDatasetLastRefreshTime }).(pulumi.StringOutput)
}

func (o ProgressMetricsResponseOutput) SkippedDatapointCount() pulumi.Float64Output {
	return o.ApplyT(func(v ProgressMetricsResponse) float64 { return v.SkippedDatapointCount }).(pulumi.Float64Output)
}

func (o ProgressMetricsResponseOutput) TotalDatapointCount() pulumi.Float64Output {
	return o.ApplyT(func(v ProgressMetricsResponse) float64 { return v.TotalDatapointCount }).(pulumi.Float64Output)
}

type PyTorch struct {
	DistributionType string `pulumi:"distributionType"`
	ProcessCount     *int   `pulumi:"processCount"`
}

type PyTorchResponse struct {
	DistributionType string `pulumi:"distributionType"`
	ProcessCount     *int   `pulumi:"processCount"`
}

type Recurrence struct {
	Frequency *string             `pulumi:"frequency"`
	Interval  *int                `pulumi:"interval"`
	Schedule  *RecurrenceSchedule `pulumi:"schedule"`
	StartTime *string             `pulumi:"startTime"`
	TimeZone  *string             `pulumi:"timeZone"`
}

type RecurrenceResponse struct {
	Frequency *string                     `pulumi:"frequency"`
	Interval  *int                        `pulumi:"interval"`
	Schedule  *RecurrenceScheduleResponse `pulumi:"schedule"`
	StartTime *string                     `pulumi:"startTime"`
	TimeZone  *string                     `pulumi:"timeZone"`
}

type RecurrenceSchedule struct {
	Hours    []int        `pulumi:"hours"`
	Minutes  []int        `pulumi:"minutes"`
	WeekDays []DaysOfWeek `pulumi:"weekDays"`
}

type RecurrenceScheduleResponse struct {
	Hours    []int    `pulumi:"hours"`
	Minutes  []int    `pulumi:"minutes"`
	WeekDays []string `pulumi:"weekDays"`
}

type RegistryListCredentialsResultResponse struct {
	Location  string             `pulumi:"location"`
	Passwords []PasswordResponse `pulumi:"passwords"`
	Username  string             `pulumi:"username"`
}

type ResourceId struct {
	Id string `pulumi:"id"`
}

type ResourceIdResponse struct {
	Id string `pulumi:"id"`
}

type ResourceIdentity struct {
	Type                   *string                             `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityMeta `pulumi:"userAssignedIdentities"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	Type                   pulumi.StringPtrInput            `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityMetaMapInput `pulumi:"userAssignedIdentities"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}









type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

func (o ResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityOutput) UserAssignedIdentities() UserAssignedIdentityMetaMapOutput {
	return o.ApplyT(func(v ResourceIdentity) map[string]UserAssignedIdentityMeta { return v.UserAssignedIdentities }).(UserAssignedIdentityMetaMapOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

func (o ResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityPtrOutput) UserAssignedIdentities() UserAssignedIdentityMetaMapOutput {
	return o.ApplyT(func(v *ResourceIdentity) map[string]UserAssignedIdentityMeta {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityMetaMapOutput)
}

type ResourceIdentityResponse struct {
	PrincipalId            string                                      `pulumi:"principalId"`
	TenantId               string                                      `pulumi:"tenantId"`
	Type                   *string                                     `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityMetaResponse `pulumi:"userAssignedIdentities"`
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityMetaResponseMapOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) map[string]UserAssignedIdentityMetaResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityMetaResponseMapOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

func (o ResourceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityMetaResponseMapOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) map[string]UserAssignedIdentityMetaResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityMetaResponseMapOutput)
}

type Route struct {
	Path string `pulumi:"path"`
	Port int    `pulumi:"port"`
}





type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(context.Context) RouteOutput
}

type RouteArgs struct {
	Path pulumi.StringInput `pulumi:"path"`
	Port pulumi.IntInput    `pulumi:"port"`
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil)).Elem()
}

func (i RouteArgs) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i RouteArgs) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

func (i RouteArgs) ToRoutePtrOutput() RoutePtrOutput {
	return i.ToRoutePtrOutputWithContext(context.Background())
}

func (i RouteArgs) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput).ToRoutePtrOutputWithContext(ctx)
}









type RoutePtrInput interface {
	pulumi.Input

	ToRoutePtrOutput() RoutePtrOutput
	ToRoutePtrOutputWithContext(context.Context) RoutePtrOutput
}

type routePtrType RouteArgs

func RoutePtr(v *RouteArgs) RoutePtrInput {
	return (*routePtrType)(v)
}

func (*routePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *routePtrType) ToRoutePtrOutput() RoutePtrOutput {
	return i.ToRoutePtrOutputWithContext(context.Background())
}

func (i *routePtrType) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePtrOutput)
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func (o RouteOutput) ToRoutePtrOutput() RoutePtrOutput {
	return o.ToRoutePtrOutputWithContext(context.Background())
}

func (o RouteOutput) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Route) *Route {
		return &v
	}).(RoutePtrOutput)
}

func (o RouteOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v Route) string { return v.Path }).(pulumi.StringOutput)
}

func (o RouteOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v Route) int { return v.Port }).(pulumi.IntOutput)
}

type RoutePtrOutput struct{ *pulumi.OutputState }

func (RoutePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RoutePtrOutput) ToRoutePtrOutput() RoutePtrOutput {
	return o
}

func (o RoutePtrOutput) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return o
}

func (o RoutePtrOutput) Elem() RouteOutput {
	return o.ApplyT(func(v *Route) Route {
		if v != nil {
			return *v
		}
		var ret Route
		return ret
	}).(RouteOutput)
}

func (o RoutePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) *string {
		if v == nil {
			return nil
		}
		return &v.Path
	}).(pulumi.StringPtrOutput)
}

func (o RoutePtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Route) *int {
		if v == nil {
			return nil
		}
		return &v.Port
	}).(pulumi.IntPtrOutput)
}

type RouteResponse struct {
	Path string `pulumi:"path"`
	Port int    `pulumi:"port"`
}

type RouteResponseOutput struct{ *pulumi.OutputState }

func (RouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteResponse)(nil)).Elem()
}

func (o RouteResponseOutput) ToRouteResponseOutput() RouteResponseOutput {
	return o
}

func (o RouteResponseOutput) ToRouteResponseOutputWithContext(ctx context.Context) RouteResponseOutput {
	return o
}

func (o RouteResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v RouteResponse) string { return v.Path }).(pulumi.StringOutput)
}

func (o RouteResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v RouteResponse) int { return v.Port }).(pulumi.IntOutput)
}

type RouteResponsePtrOutput struct{ *pulumi.OutputState }

func (RouteResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteResponse)(nil)).Elem()
}

func (o RouteResponsePtrOutput) ToRouteResponsePtrOutput() RouteResponsePtrOutput {
	return o
}

func (o RouteResponsePtrOutput) ToRouteResponsePtrOutputWithContext(ctx context.Context) RouteResponsePtrOutput {
	return o
}

func (o RouteResponsePtrOutput) Elem() RouteResponseOutput {
	return o.ApplyT(func(v *RouteResponse) RouteResponse {
		if v != nil {
			return *v
		}
		var ret RouteResponse
		return ret
	}).(RouteResponseOutput)
}

func (o RouteResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Path
	}).(pulumi.StringPtrOutput)
}

func (o RouteResponsePtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RouteResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Port
	}).(pulumi.IntPtrOutput)
}

type SasDatastoreCredentials struct {
	CredentialsType string               `pulumi:"credentialsType"`
	Secrets         *SasDatastoreSecrets `pulumi:"secrets"`
}

type SasDatastoreCredentialsResponse struct {
	CredentialsType string `pulumi:"credentialsType"`
}

type SasDatastoreSecrets struct {
	SasToken    *string `pulumi:"sasToken"`
	SecretsType string  `pulumi:"secretsType"`
}

type ScaleSettings struct {
	MaxNodeCount                int     `pulumi:"maxNodeCount"`
	MinNodeCount                *int    `pulumi:"minNodeCount"`
	NodeIdleTimeBeforeScaleDown *string `pulumi:"nodeIdleTimeBeforeScaleDown"`
}


func (val *ScaleSettings) Defaults() *ScaleSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MinNodeCount) {
		minNodeCount_ := 0
		tmp.MinNodeCount = &minNodeCount_
	}
	return &tmp
}

type ScaleSettingsResponse struct {
	MaxNodeCount                int     `pulumi:"maxNodeCount"`
	MinNodeCount                *int    `pulumi:"minNodeCount"`
	NodeIdleTimeBeforeScaleDown *string `pulumi:"nodeIdleTimeBeforeScaleDown"`
}


func (val *ScaleSettingsResponse) Defaults() *ScaleSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MinNodeCount) {
		minNodeCount_ := 0
		tmp.MinNodeCount = &minNodeCount_
	}
	return &tmp
}

type ScriptReference struct {
	ScriptArguments *string `pulumi:"scriptArguments"`
	ScriptData      *string `pulumi:"scriptData"`
	ScriptSource    *string `pulumi:"scriptSource"`
	Timeout         *string `pulumi:"timeout"`
}

type ScriptReferenceResponse struct {
	ScriptArguments *string `pulumi:"scriptArguments"`
	ScriptData      *string `pulumi:"scriptData"`
	ScriptSource    *string `pulumi:"scriptSource"`
	Timeout         *string `pulumi:"timeout"`
}

type ScriptsToExecute struct {
	CreationScript *ScriptReference `pulumi:"creationScript"`
	StartupScript  *ScriptReference `pulumi:"startupScript"`
}

type ScriptsToExecuteResponse struct {
	CreationScript *ScriptReferenceResponse `pulumi:"creationScript"`
	StartupScript  *ScriptReferenceResponse `pulumi:"startupScript"`
}

type ServiceManagedResourcesSettings struct {
	CosmosDb *CosmosDbSettings `pulumi:"cosmosDb"`
}





type ServiceManagedResourcesSettingsInput interface {
	pulumi.Input

	ToServiceManagedResourcesSettingsOutput() ServiceManagedResourcesSettingsOutput
	ToServiceManagedResourcesSettingsOutputWithContext(context.Context) ServiceManagedResourcesSettingsOutput
}

type ServiceManagedResourcesSettingsArgs struct {
	CosmosDb CosmosDbSettingsPtrInput `pulumi:"cosmosDb"`
}

func (ServiceManagedResourcesSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedResourcesSettings)(nil)).Elem()
}

func (i ServiceManagedResourcesSettingsArgs) ToServiceManagedResourcesSettingsOutput() ServiceManagedResourcesSettingsOutput {
	return i.ToServiceManagedResourcesSettingsOutputWithContext(context.Background())
}

func (i ServiceManagedResourcesSettingsArgs) ToServiceManagedResourcesSettingsOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedResourcesSettingsOutput)
}

func (i ServiceManagedResourcesSettingsArgs) ToServiceManagedResourcesSettingsPtrOutput() ServiceManagedResourcesSettingsPtrOutput {
	return i.ToServiceManagedResourcesSettingsPtrOutputWithContext(context.Background())
}

func (i ServiceManagedResourcesSettingsArgs) ToServiceManagedResourcesSettingsPtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedResourcesSettingsOutput).ToServiceManagedResourcesSettingsPtrOutputWithContext(ctx)
}









type ServiceManagedResourcesSettingsPtrInput interface {
	pulumi.Input

	ToServiceManagedResourcesSettingsPtrOutput() ServiceManagedResourcesSettingsPtrOutput
	ToServiceManagedResourcesSettingsPtrOutputWithContext(context.Context) ServiceManagedResourcesSettingsPtrOutput
}

type serviceManagedResourcesSettingsPtrType ServiceManagedResourcesSettingsArgs

func ServiceManagedResourcesSettingsPtr(v *ServiceManagedResourcesSettingsArgs) ServiceManagedResourcesSettingsPtrInput {
	return (*serviceManagedResourcesSettingsPtrType)(v)
}

func (*serviceManagedResourcesSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceManagedResourcesSettings)(nil)).Elem()
}

func (i *serviceManagedResourcesSettingsPtrType) ToServiceManagedResourcesSettingsPtrOutput() ServiceManagedResourcesSettingsPtrOutput {
	return i.ToServiceManagedResourcesSettingsPtrOutputWithContext(context.Background())
}

func (i *serviceManagedResourcesSettingsPtrType) ToServiceManagedResourcesSettingsPtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedResourcesSettingsPtrOutput)
}

type ServiceManagedResourcesSettingsOutput struct{ *pulumi.OutputState }

func (ServiceManagedResourcesSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedResourcesSettings)(nil)).Elem()
}

func (o ServiceManagedResourcesSettingsOutput) ToServiceManagedResourcesSettingsOutput() ServiceManagedResourcesSettingsOutput {
	return o
}

func (o ServiceManagedResourcesSettingsOutput) ToServiceManagedResourcesSettingsOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsOutput {
	return o
}

func (o ServiceManagedResourcesSettingsOutput) ToServiceManagedResourcesSettingsPtrOutput() ServiceManagedResourcesSettingsPtrOutput {
	return o.ToServiceManagedResourcesSettingsPtrOutputWithContext(context.Background())
}

func (o ServiceManagedResourcesSettingsOutput) ToServiceManagedResourcesSettingsPtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceManagedResourcesSettings) *ServiceManagedResourcesSettings {
		return &v
	}).(ServiceManagedResourcesSettingsPtrOutput)
}

func (o ServiceManagedResourcesSettingsOutput) CosmosDb() CosmosDbSettingsPtrOutput {
	return o.ApplyT(func(v ServiceManagedResourcesSettings) *CosmosDbSettings { return v.CosmosDb }).(CosmosDbSettingsPtrOutput)
}

type ServiceManagedResourcesSettingsPtrOutput struct{ *pulumi.OutputState }

func (ServiceManagedResourcesSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceManagedResourcesSettings)(nil)).Elem()
}

func (o ServiceManagedResourcesSettingsPtrOutput) ToServiceManagedResourcesSettingsPtrOutput() ServiceManagedResourcesSettingsPtrOutput {
	return o
}

func (o ServiceManagedResourcesSettingsPtrOutput) ToServiceManagedResourcesSettingsPtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsPtrOutput {
	return o
}

func (o ServiceManagedResourcesSettingsPtrOutput) Elem() ServiceManagedResourcesSettingsOutput {
	return o.ApplyT(func(v *ServiceManagedResourcesSettings) ServiceManagedResourcesSettings {
		if v != nil {
			return *v
		}
		var ret ServiceManagedResourcesSettings
		return ret
	}).(ServiceManagedResourcesSettingsOutput)
}

func (o ServiceManagedResourcesSettingsPtrOutput) CosmosDb() CosmosDbSettingsPtrOutput {
	return o.ApplyT(func(v *ServiceManagedResourcesSettings) *CosmosDbSettings {
		if v == nil {
			return nil
		}
		return v.CosmosDb
	}).(CosmosDbSettingsPtrOutput)
}

type ServiceManagedResourcesSettingsResponse struct {
	CosmosDb *CosmosDbSettingsResponse `pulumi:"cosmosDb"`
}

type ServiceManagedResourcesSettingsResponseOutput struct{ *pulumi.OutputState }

func (ServiceManagedResourcesSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedResourcesSettingsResponse)(nil)).Elem()
}

func (o ServiceManagedResourcesSettingsResponseOutput) ToServiceManagedResourcesSettingsResponseOutput() ServiceManagedResourcesSettingsResponseOutput {
	return o
}

func (o ServiceManagedResourcesSettingsResponseOutput) ToServiceManagedResourcesSettingsResponseOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsResponseOutput {
	return o
}

func (o ServiceManagedResourcesSettingsResponseOutput) CosmosDb() CosmosDbSettingsResponsePtrOutput {
	return o.ApplyT(func(v ServiceManagedResourcesSettingsResponse) *CosmosDbSettingsResponse { return v.CosmosDb }).(CosmosDbSettingsResponsePtrOutput)
}

type ServiceManagedResourcesSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceManagedResourcesSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceManagedResourcesSettingsResponse)(nil)).Elem()
}

func (o ServiceManagedResourcesSettingsResponsePtrOutput) ToServiceManagedResourcesSettingsResponsePtrOutput() ServiceManagedResourcesSettingsResponsePtrOutput {
	return o
}

func (o ServiceManagedResourcesSettingsResponsePtrOutput) ToServiceManagedResourcesSettingsResponsePtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsResponsePtrOutput {
	return o
}

func (o ServiceManagedResourcesSettingsResponsePtrOutput) Elem() ServiceManagedResourcesSettingsResponseOutput {
	return o.ApplyT(func(v *ServiceManagedResourcesSettingsResponse) ServiceManagedResourcesSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ServiceManagedResourcesSettingsResponse
		return ret
	}).(ServiceManagedResourcesSettingsResponseOutput)
}

func (o ServiceManagedResourcesSettingsResponsePtrOutput) CosmosDb() CosmosDbSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ServiceManagedResourcesSettingsResponse) *CosmosDbSettingsResponse {
		if v == nil {
			return nil
		}
		return v.CosmosDb
	}).(CosmosDbSettingsResponsePtrOutput)
}

type ServicePrincipalDatastoreCredentials struct {
	AuthorityUrl    *string                           `pulumi:"authorityUrl"`
	ClientId        string                            `pulumi:"clientId"`
	CredentialsType string                            `pulumi:"credentialsType"`
	ResourceUri     *string                           `pulumi:"resourceUri"`
	Secrets         *ServicePrincipalDatastoreSecrets `pulumi:"secrets"`
	TenantId        string                            `pulumi:"tenantId"`
}

type ServicePrincipalDatastoreCredentialsResponse struct {
	AuthorityUrl    *string `pulumi:"authorityUrl"`
	ClientId        string  `pulumi:"clientId"`
	CredentialsType string  `pulumi:"credentialsType"`
	ResourceUri     *string `pulumi:"resourceUri"`
	TenantId        string  `pulumi:"tenantId"`
}

type ServicePrincipalDatastoreSecrets struct {
	ClientSecret *string `pulumi:"clientSecret"`
	SecretsType  string  `pulumi:"secretsType"`
}

type SetupScripts struct {
	Scripts *ScriptsToExecute `pulumi:"scripts"`
}

type SetupScriptsResponse struct {
	Scripts *ScriptsToExecuteResponse `pulumi:"scripts"`
}

type SharedPrivateLinkResource struct {
	GroupId               *string `pulumi:"groupId"`
	Name                  *string `pulumi:"name"`
	PrivateLinkResourceId *string `pulumi:"privateLinkResourceId"`
	RequestMessage        *string `pulumi:"requestMessage"`
	Status                *string `pulumi:"status"`
}





type SharedPrivateLinkResourceInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceOutput() SharedPrivateLinkResourceOutput
	ToSharedPrivateLinkResourceOutputWithContext(context.Context) SharedPrivateLinkResourceOutput
}

type SharedPrivateLinkResourceArgs struct {
	GroupId               pulumi.StringPtrInput `pulumi:"groupId"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	PrivateLinkResourceId pulumi.StringPtrInput `pulumi:"privateLinkResourceId"`
	RequestMessage        pulumi.StringPtrInput `pulumi:"requestMessage"`
	Status                pulumi.StringPtrInput `pulumi:"status"`
}

func (SharedPrivateLinkResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResource)(nil)).Elem()
}

func (i SharedPrivateLinkResourceArgs) ToSharedPrivateLinkResourceOutput() SharedPrivateLinkResourceOutput {
	return i.ToSharedPrivateLinkResourceOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceArgs) ToSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SharedPrivateLinkResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceOutput)
}





type SharedPrivateLinkResourceArrayInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceArrayOutput() SharedPrivateLinkResourceArrayOutput
	ToSharedPrivateLinkResourceArrayOutputWithContext(context.Context) SharedPrivateLinkResourceArrayOutput
}

type SharedPrivateLinkResourceArray []SharedPrivateLinkResourceInput

func (SharedPrivateLinkResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResource)(nil)).Elem()
}

func (i SharedPrivateLinkResourceArray) ToSharedPrivateLinkResourceArrayOutput() SharedPrivateLinkResourceArrayOutput {
	return i.ToSharedPrivateLinkResourceArrayOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceArray) ToSharedPrivateLinkResourceArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceArrayOutput)
}

type SharedPrivateLinkResourceOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResource)(nil)).Elem()
}

func (o SharedPrivateLinkResourceOutput) ToSharedPrivateLinkResourceOutput() SharedPrivateLinkResourceOutput {
	return o
}

func (o SharedPrivateLinkResourceOutput) ToSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SharedPrivateLinkResourceOutput {
	return o
}

func (o SharedPrivateLinkResourceOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResource) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResource) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResource) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResource) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourceArrayOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResource)(nil)).Elem()
}

func (o SharedPrivateLinkResourceArrayOutput) ToSharedPrivateLinkResourceArrayOutput() SharedPrivateLinkResourceArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceArrayOutput) ToSharedPrivateLinkResourceArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceArrayOutput) Index(i pulumi.IntInput) SharedPrivateLinkResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedPrivateLinkResource {
		return vs[0].([]SharedPrivateLinkResource)[vs[1].(int)]
	}).(SharedPrivateLinkResourceOutput)
}

type SharedPrivateLinkResourceResponse struct {
	GroupId               *string `pulumi:"groupId"`
	Name                  *string `pulumi:"name"`
	PrivateLinkResourceId *string `pulumi:"privateLinkResourceId"`
	RequestMessage        *string `pulumi:"requestMessage"`
	Status                *string `pulumi:"status"`
}

type SharedPrivateLinkResourceResponseOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourceResponseOutput) ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseOutput) ToSharedPrivateLinkResourceResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourceResponseArrayOutput) ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseArrayOutput) ToSharedPrivateLinkResourceResponseArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseArrayOutput) Index(i pulumi.IntInput) SharedPrivateLinkResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedPrivateLinkResourceResponse {
		return vs[0].([]SharedPrivateLinkResourceResponse)[vs[1].(int)]
	}).(SharedPrivateLinkResourceResponseOutput)
}

type Sku struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SqlAdminDatastoreCredentials struct {
	CredentialsType string                    `pulumi:"credentialsType"`
	Secrets         *SqlAdminDatastoreSecrets `pulumi:"secrets"`
	UserId          string                    `pulumi:"userId"`
}

type SqlAdminDatastoreCredentialsResponse struct {
	CredentialsType string `pulumi:"credentialsType"`
	UserId          string `pulumi:"userId"`
}

type SqlAdminDatastoreSecrets struct {
	Password    *string `pulumi:"password"`
	SecretsType string  `pulumi:"secretsType"`
}

type SslConfiguration struct {
	Cert                    *string `pulumi:"cert"`
	Cname                   *string `pulumi:"cname"`
	Key                     *string `pulumi:"key"`
	LeafDomainLabel         *string `pulumi:"leafDomainLabel"`
	OverwriteExistingDomain *bool   `pulumi:"overwriteExistingDomain"`
	Status                  *string `pulumi:"status"`
}

type SslConfigurationResponse struct {
	Cert                    *string `pulumi:"cert"`
	Cname                   *string `pulumi:"cname"`
	Key                     *string `pulumi:"key"`
	LeafDomainLabel         *string `pulumi:"leafDomainLabel"`
	OverwriteExistingDomain *bool   `pulumi:"overwriteExistingDomain"`
	Status                  *string `pulumi:"status"`
}

type StatusMessageResponse struct {
	Code           string `pulumi:"code"`
	CreatedTimeUtc string `pulumi:"createdTimeUtc"`
	Level          string `pulumi:"level"`
	Message        string `pulumi:"message"`
}

type StatusMessageResponseOutput struct{ *pulumi.OutputState }

func (StatusMessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusMessageResponse)(nil)).Elem()
}

func (o StatusMessageResponseOutput) ToStatusMessageResponseOutput() StatusMessageResponseOutput {
	return o
}

func (o StatusMessageResponseOutput) ToStatusMessageResponseOutputWithContext(ctx context.Context) StatusMessageResponseOutput {
	return o
}

func (o StatusMessageResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v StatusMessageResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o StatusMessageResponseOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v StatusMessageResponse) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o StatusMessageResponseOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v StatusMessageResponse) string { return v.Level }).(pulumi.StringOutput)
}

func (o StatusMessageResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v StatusMessageResponse) string { return v.Message }).(pulumi.StringOutput)
}

type StatusMessageResponseArrayOutput struct{ *pulumi.OutputState }

func (StatusMessageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusMessageResponse)(nil)).Elem()
}

func (o StatusMessageResponseArrayOutput) ToStatusMessageResponseArrayOutput() StatusMessageResponseArrayOutput {
	return o
}

func (o StatusMessageResponseArrayOutput) ToStatusMessageResponseArrayOutputWithContext(ctx context.Context) StatusMessageResponseArrayOutput {
	return o
}

func (o StatusMessageResponseArrayOutput) Index(i pulumi.IntInput) StatusMessageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusMessageResponse {
		return vs[0].([]StatusMessageResponse)[vs[1].(int)]
	}).(StatusMessageResponseOutput)
}

type SweepJob struct {
	Algorithm           string                 `pulumi:"algorithm"`
	Compute             ComputeConfiguration   `pulumi:"compute"`
	Description         *string                `pulumi:"description"`
	EarlyTermination    interface{}            `pulumi:"earlyTermination"`
	ExperimentName      *string                `pulumi:"experimentName"`
	Identity            interface{}            `pulumi:"identity"`
	JobType             string                 `pulumi:"jobType"`
	MaxConcurrentTrials *int                   `pulumi:"maxConcurrentTrials"`
	MaxTotalTrials      *int                   `pulumi:"maxTotalTrials"`
	Objective           Objective              `pulumi:"objective"`
	Priority            *int                   `pulumi:"priority"`
	Properties          map[string]string      `pulumi:"properties"`
	SearchSpace         map[string]interface{} `pulumi:"searchSpace"`
	Tags                map[string]string      `pulumi:"tags"`
	Timeout             *string                `pulumi:"timeout"`
	Trial               *TrialComponent        `pulumi:"trial"`
}

type SweepJobResponse struct {
	Algorithm            string                         `pulumi:"algorithm"`
	Compute              ComputeConfigurationResponse   `pulumi:"compute"`
	Description          *string                        `pulumi:"description"`
	EarlyTermination     interface{}                    `pulumi:"earlyTermination"`
	ExperimentName       *string                        `pulumi:"experimentName"`
	Identity             interface{}                    `pulumi:"identity"`
	InteractionEndpoints map[string]JobEndpointResponse `pulumi:"interactionEndpoints"`
	JobType              string                         `pulumi:"jobType"`
	MaxConcurrentTrials  *int                           `pulumi:"maxConcurrentTrials"`
	MaxTotalTrials       *int                           `pulumi:"maxTotalTrials"`
	Objective            ObjectiveResponse              `pulumi:"objective"`
	Output               JobOutputResponse              `pulumi:"output"`
	Priority             *int                           `pulumi:"priority"`
	Properties           map[string]string              `pulumi:"properties"`
	ProvisioningState    string                         `pulumi:"provisioningState"`
	SearchSpace          map[string]interface{}         `pulumi:"searchSpace"`
	Status               string                         `pulumi:"status"`
	Tags                 map[string]string              `pulumi:"tags"`
	Timeout              *string                        `pulumi:"timeout"`
	Trial                *TrialComponentResponse        `pulumi:"trial"`
}

type SynapseSpark struct {
	ComputeLocation  *string                               `pulumi:"computeLocation"`
	ComputeType      string                                `pulumi:"computeType"`
	Description      *string                               `pulumi:"description"`
	DisableLocalAuth *bool                                 `pulumi:"disableLocalAuth"`
	Properties       *SynapseSparkPoolPropertiesProperties `pulumi:"properties"`
	ResourceId       *string                               `pulumi:"resourceId"`
}

type SynapseSparkPoolPropertiesProperties struct {
	AutoPauseProperties *AutoPauseProperties `pulumi:"autoPauseProperties"`
	AutoScaleProperties *AutoScaleProperties `pulumi:"autoScaleProperties"`
	NodeCount           *int                 `pulumi:"nodeCount"`
	NodeSize            *string              `pulumi:"nodeSize"`
	NodeSizeFamily      *string              `pulumi:"nodeSizeFamily"`
	PoolName            *string              `pulumi:"poolName"`
	ResourceGroup       *string              `pulumi:"resourceGroup"`
	SparkVersion        *string              `pulumi:"sparkVersion"`
	SubscriptionId      *string              `pulumi:"subscriptionId"`
	WorkspaceName       *string              `pulumi:"workspaceName"`
}

type SynapseSparkPoolPropertiesResponseProperties struct {
	AutoPauseProperties *AutoPausePropertiesResponse `pulumi:"autoPauseProperties"`
	AutoScaleProperties *AutoScalePropertiesResponse `pulumi:"autoScaleProperties"`
	NodeCount           *int                         `pulumi:"nodeCount"`
	NodeSize            *string                      `pulumi:"nodeSize"`
	NodeSizeFamily      *string                      `pulumi:"nodeSizeFamily"`
	PoolName            *string                      `pulumi:"poolName"`
	ResourceGroup       *string                      `pulumi:"resourceGroup"`
	SparkVersion        *string                      `pulumi:"sparkVersion"`
	SubscriptionId      *string                      `pulumi:"subscriptionId"`
	WorkspaceName       *string                      `pulumi:"workspaceName"`
}

type SynapseSparkResponse struct {
	ComputeLocation    *string                                       `pulumi:"computeLocation"`
	ComputeType        string                                        `pulumi:"computeType"`
	CreatedOn          string                                        `pulumi:"createdOn"`
	Description        *string                                       `pulumi:"description"`
	DisableLocalAuth   *bool                                         `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                                          `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                        `pulumi:"modifiedOn"`
	Properties         *SynapseSparkPoolPropertiesResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse                       `pulumi:"provisioningErrors"`
	ProvisioningState  string                                        `pulumi:"provisioningState"`
	ResourceId         *string                                       `pulumi:"resourceId"`
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

type SystemServiceResponse struct {
	PublicIpAddress   string `pulumi:"publicIpAddress"`
	SystemServiceType string `pulumi:"systemServiceType"`
	Version           string `pulumi:"version"`
}

type TensorFlow struct {
	DistributionType     string `pulumi:"distributionType"`
	ParameterServerCount *int   `pulumi:"parameterServerCount"`
	WorkerCount          *int   `pulumi:"workerCount"`
}

type TensorFlowResponse struct {
	DistributionType     string `pulumi:"distributionType"`
	ParameterServerCount *int   `pulumi:"parameterServerCount"`
	WorkerCount          *int   `pulumi:"workerCount"`
}

type TrialComponent struct {
	CodeId               *string                      `pulumi:"codeId"`
	Command              string                       `pulumi:"command"`
	Distribution         interface{}                  `pulumi:"distribution"`
	EnvironmentId        *string                      `pulumi:"environmentId"`
	EnvironmentVariables map[string]string            `pulumi:"environmentVariables"`
	InputDataBindings    map[string]InputDataBinding  `pulumi:"inputDataBindings"`
	OutputDataBindings   map[string]OutputDataBinding `pulumi:"outputDataBindings"`
	Timeout              *string                      `pulumi:"timeout"`
}

type TrialComponentResponse struct {
	CodeId               *string                              `pulumi:"codeId"`
	Command              string                               `pulumi:"command"`
	Distribution         interface{}                          `pulumi:"distribution"`
	EnvironmentId        *string                              `pulumi:"environmentId"`
	EnvironmentVariables map[string]string                    `pulumi:"environmentVariables"`
	InputDataBindings    map[string]InputDataBindingResponse  `pulumi:"inputDataBindings"`
	OutputDataBindings   map[string]OutputDataBindingResponse `pulumi:"outputDataBindings"`
	Timeout              *string                              `pulumi:"timeout"`
}

type TruncationSelectionPolicy struct {
	DelayEvaluation      *int   `pulumi:"delayEvaluation"`
	EvaluationInterval   *int   `pulumi:"evaluationInterval"`
	PolicyType           string `pulumi:"policyType"`
	TruncationPercentage *int   `pulumi:"truncationPercentage"`
}

type TruncationSelectionPolicyResponse struct {
	DelayEvaluation      *int   `pulumi:"delayEvaluation"`
	EvaluationInterval   *int   `pulumi:"evaluationInterval"`
	PolicyType           string `pulumi:"policyType"`
	TruncationPercentage *int   `pulumi:"truncationPercentage"`
}

type UserAccountCredentials struct {
	AdminUserName         string  `pulumi:"adminUserName"`
	AdminUserPassword     *string `pulumi:"adminUserPassword"`
	AdminUserSshPublicKey *string `pulumi:"adminUserSshPublicKey"`
}

type UserAccountCredentialsResponse struct {
	AdminUserName         string  `pulumi:"adminUserName"`
	AdminUserPassword     *string `pulumi:"adminUserPassword"`
	AdminUserSshPublicKey *string `pulumi:"adminUserSshPublicKey"`
}

type UserAssignedIdentityMeta struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}





type UserAssignedIdentityMetaInput interface {
	pulumi.Input

	ToUserAssignedIdentityMetaOutput() UserAssignedIdentityMetaOutput
	ToUserAssignedIdentityMetaOutputWithContext(context.Context) UserAssignedIdentityMetaOutput
}

type UserAssignedIdentityMetaArgs struct {
	ClientId    pulumi.StringPtrInput `pulumi:"clientId"`
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (UserAssignedIdentityMetaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityMeta)(nil)).Elem()
}

func (i UserAssignedIdentityMetaArgs) ToUserAssignedIdentityMetaOutput() UserAssignedIdentityMetaOutput {
	return i.ToUserAssignedIdentityMetaOutputWithContext(context.Background())
}

func (i UserAssignedIdentityMetaArgs) ToUserAssignedIdentityMetaOutputWithContext(ctx context.Context) UserAssignedIdentityMetaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityMetaOutput)
}





type UserAssignedIdentityMetaMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityMetaMapOutput() UserAssignedIdentityMetaMapOutput
	ToUserAssignedIdentityMetaMapOutputWithContext(context.Context) UserAssignedIdentityMetaMapOutput
}

type UserAssignedIdentityMetaMap map[string]UserAssignedIdentityMetaInput

func (UserAssignedIdentityMetaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityMeta)(nil)).Elem()
}

func (i UserAssignedIdentityMetaMap) ToUserAssignedIdentityMetaMapOutput() UserAssignedIdentityMetaMapOutput {
	return i.ToUserAssignedIdentityMetaMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityMetaMap) ToUserAssignedIdentityMetaMapOutputWithContext(ctx context.Context) UserAssignedIdentityMetaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityMetaMapOutput)
}

type UserAssignedIdentityMetaOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityMetaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityMeta)(nil)).Elem()
}

func (o UserAssignedIdentityMetaOutput) ToUserAssignedIdentityMetaOutput() UserAssignedIdentityMetaOutput {
	return o
}

func (o UserAssignedIdentityMetaOutput) ToUserAssignedIdentityMetaOutputWithContext(ctx context.Context) UserAssignedIdentityMetaOutput {
	return o
}

func (o UserAssignedIdentityMetaOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityMeta) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserAssignedIdentityMetaOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityMeta) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityMetaMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityMetaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityMeta)(nil)).Elem()
}

func (o UserAssignedIdentityMetaMapOutput) ToUserAssignedIdentityMetaMapOutput() UserAssignedIdentityMetaMapOutput {
	return o
}

func (o UserAssignedIdentityMetaMapOutput) ToUserAssignedIdentityMetaMapOutputWithContext(ctx context.Context) UserAssignedIdentityMetaMapOutput {
	return o
}

func (o UserAssignedIdentityMetaMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityMetaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityMeta {
		return vs[0].(map[string]UserAssignedIdentityMeta)[vs[1].(string)]
	}).(UserAssignedIdentityMetaOutput)
}

type UserAssignedIdentityMetaResponse struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}

type UserAssignedIdentityMetaResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityMetaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityMetaResponse)(nil)).Elem()
}

func (o UserAssignedIdentityMetaResponseOutput) ToUserAssignedIdentityMetaResponseOutput() UserAssignedIdentityMetaResponseOutput {
	return o
}

func (o UserAssignedIdentityMetaResponseOutput) ToUserAssignedIdentityMetaResponseOutputWithContext(ctx context.Context) UserAssignedIdentityMetaResponseOutput {
	return o
}

func (o UserAssignedIdentityMetaResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityMetaResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserAssignedIdentityMetaResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityMetaResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityMetaResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityMetaResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityMetaResponse)(nil)).Elem()
}

func (o UserAssignedIdentityMetaResponseMapOutput) ToUserAssignedIdentityMetaResponseMapOutput() UserAssignedIdentityMetaResponseMapOutput {
	return o
}

func (o UserAssignedIdentityMetaResponseMapOutput) ToUserAssignedIdentityMetaResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityMetaResponseMapOutput {
	return o
}

func (o UserAssignedIdentityMetaResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityMetaResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityMetaResponse {
		return vs[0].(map[string]UserAssignedIdentityMetaResponse)[vs[1].(string)]
	}).(UserAssignedIdentityMetaResponseOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type VirtualMachine struct {
	ComputeLocation  *string                   `pulumi:"computeLocation"`
	ComputeType      string                    `pulumi:"computeType"`
	Description      *string                   `pulumi:"description"`
	DisableLocalAuth *bool                     `pulumi:"disableLocalAuth"`
	Properties       *VirtualMachineProperties `pulumi:"properties"`
	ResourceId       *string                   `pulumi:"resourceId"`
}

type VirtualMachineImage struct {
	Id string `pulumi:"id"`
}

type VirtualMachineImageResponse struct {
	Id string `pulumi:"id"`
}

type VirtualMachineProperties struct {
	Address                   *string                       `pulumi:"address"`
	AdministratorAccount      *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	IsNotebookInstanceCompute *bool                         `pulumi:"isNotebookInstanceCompute"`
	SshPort                   *int                          `pulumi:"sshPort"`
	VirtualMachineSize        *string                       `pulumi:"virtualMachineSize"`
}

type VirtualMachineResponse struct {
	ComputeLocation    *string                           `pulumi:"computeLocation"`
	ComputeType        string                            `pulumi:"computeType"`
	CreatedOn          string                            `pulumi:"createdOn"`
	Description        *string                           `pulumi:"description"`
	DisableLocalAuth   *bool                             `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                              `pulumi:"isAttachedCompute"`
	ModifiedOn         string                            `pulumi:"modifiedOn"`
	Properties         *VirtualMachineResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse           `pulumi:"provisioningErrors"`
	ProvisioningState  string                            `pulumi:"provisioningState"`
	ResourceId         *string                           `pulumi:"resourceId"`
}

type VirtualMachineResponseProperties struct {
	Address                   *string                               `pulumi:"address"`
	AdministratorAccount      *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	IsNotebookInstanceCompute *bool                                 `pulumi:"isNotebookInstanceCompute"`
	SshPort                   *int                                  `pulumi:"sshPort"`
	VirtualMachineSize        *string                               `pulumi:"virtualMachineSize"`
}

type VirtualMachineSshCredentials struct {
	Password       *string `pulumi:"password"`
	PrivateKeyData *string `pulumi:"privateKeyData"`
	PublicKeyData  *string `pulumi:"publicKeyData"`
	Username       *string `pulumi:"username"`
}

type VirtualMachineSshCredentialsResponse struct {
	Password       *string `pulumi:"password"`
	PrivateKeyData *string `pulumi:"privateKeyData"`
	PublicKeyData  *string `pulumi:"publicKeyData"`
	Username       *string `pulumi:"username"`
}

func init() {
	pulumi.RegisterOutputType(BatchDeploymentTypeOutput{})
	pulumi.RegisterOutputType(BatchDeploymentResponseOutput{})
	pulumi.RegisterOutputType(BatchEndpointTypeOutput{})
	pulumi.RegisterOutputType(BatchEndpointResponseOutput{})
	pulumi.RegisterOutputType(BatchOutputConfigurationOutput{})
	pulumi.RegisterOutputType(BatchOutputConfigurationPtrOutput{})
	pulumi.RegisterOutputType(BatchOutputConfigurationResponseOutput{})
	pulumi.RegisterOutputType(BatchOutputConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(BatchRetrySettingsOutput{})
	pulumi.RegisterOutputType(BatchRetrySettingsPtrOutput{})
	pulumi.RegisterOutputType(BatchRetrySettingsResponseOutput{})
	pulumi.RegisterOutputType(BatchRetrySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CodeConfigurationOutput{})
	pulumi.RegisterOutputType(CodeConfigurationPtrOutput{})
	pulumi.RegisterOutputType(CodeConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CodeConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(CodeContainerTypeOutput{})
	pulumi.RegisterOutputType(CodeContainerResponseOutput{})
	pulumi.RegisterOutputType(CodeVersionTypeOutput{})
	pulumi.RegisterOutputType(CodeVersionResponseOutput{})
	pulumi.RegisterOutputType(ComputeConfigurationOutput{})
	pulumi.RegisterOutputType(ComputeConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ComputeConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ComputeConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsPtrOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsResponseOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(DataContainerTypeOutput{})
	pulumi.RegisterOutputType(DataContainerResponseOutput{})
	pulumi.RegisterOutputType(DataVersionTypeOutput{})
	pulumi.RegisterOutputType(DataVersionResponseOutput{})
	pulumi.RegisterOutputType(DatastorePropertiesOutput{})
	pulumi.RegisterOutputType(DatastorePropertiesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointAuthKeysOutput{})
	pulumi.RegisterOutputType(EndpointAuthKeysPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentContainerTypeOutput{})
	pulumi.RegisterOutputType(EnvironmentContainerResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentSpecificationVersionTypeOutput{})
	pulumi.RegisterOutputType(EnvironmentSpecificationVersionResponseOutput{})
	pulumi.RegisterOutputType(FlavorDataOutput{})
	pulumi.RegisterOutputType(FlavorDataMapOutput{})
	pulumi.RegisterOutputType(FlavorDataResponseOutput{})
	pulumi.RegisterOutputType(FlavorDataResponseMapOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityForCmkOutput{})
	pulumi.RegisterOutputType(IdentityForCmkPtrOutput{})
	pulumi.RegisterOutputType(IdentityForCmkResponseOutput{})
	pulumi.RegisterOutputType(IdentityForCmkResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(JobEndpointResponseOutput{})
	pulumi.RegisterOutputType(JobEndpointResponseMapOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LabelCategoryOutput{})
	pulumi.RegisterOutputType(LabelCategoryMapOutput{})
	pulumi.RegisterOutputType(LabelCategoryResponseOutput{})
	pulumi.RegisterOutputType(LabelCategoryResponseMapOutput{})
	pulumi.RegisterOutputType(LabelClassOutput{})
	pulumi.RegisterOutputType(LabelClassMapOutput{})
	pulumi.RegisterOutputType(LabelClassResponseOutput{})
	pulumi.RegisterOutputType(LabelClassResponseMapOutput{})
	pulumi.RegisterOutputType(LabelingDatasetConfigurationOutput{})
	pulumi.RegisterOutputType(LabelingDatasetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LabelingDatasetConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LabelingDatasetConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(LabelingJobTypeOutput{})
	pulumi.RegisterOutputType(LabelingJobInstructionsOutput{})
	pulumi.RegisterOutputType(LabelingJobInstructionsPtrOutput{})
	pulumi.RegisterOutputType(LabelingJobInstructionsResponseOutput{})
	pulumi.RegisterOutputType(LabelingJobInstructionsResponsePtrOutput{})
	pulumi.RegisterOutputType(LabelingJobResponseOutput{})
	pulumi.RegisterOutputType(LinkedInfoOutput{})
	pulumi.RegisterOutputType(LinkedInfoPtrOutput{})
	pulumi.RegisterOutputType(LinkedInfoResponseOutput{})
	pulumi.RegisterOutputType(LinkedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationPtrOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationResponseOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ModelContainerTypeOutput{})
	pulumi.RegisterOutputType(ModelContainerResponseOutput{})
	pulumi.RegisterOutputType(ModelVersionTypeOutput{})
	pulumi.RegisterOutputType(ModelVersionResponseOutput{})
	pulumi.RegisterOutputType(NotebookPreparationErrorResponseOutput{})
	pulumi.RegisterOutputType(NotebookPreparationErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(NotebookResourceInfoResponseOutput{})
	pulumi.RegisterOutputType(OnlineEndpointTypeOutput{})
	pulumi.RegisterOutputType(OnlineEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ProgressMetricsResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(RouteOutput{})
	pulumi.RegisterOutputType(RoutePtrOutput{})
	pulumi.RegisterOutputType(RouteResponseOutput{})
	pulumi.RegisterOutputType(RouteResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceManagedResourcesSettingsOutput{})
	pulumi.RegisterOutputType(ServiceManagedResourcesSettingsPtrOutput{})
	pulumi.RegisterOutputType(ServiceManagedResourcesSettingsResponseOutput{})
	pulumi.RegisterOutputType(ServiceManagedResourcesSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceArrayOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StatusMessageResponseOutput{})
	pulumi.RegisterOutputType(StatusMessageResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityMetaOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityMetaMapOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityMetaResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityMetaResponseMapOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}
