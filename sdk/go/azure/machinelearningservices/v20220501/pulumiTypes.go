


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AKS struct {
	ComputeType      string               `pulumi:"computeType"`
	Description      *string              `pulumi:"description"`
	DisableLocalAuth *bool                `pulumi:"disableLocalAuth"`
	Properties       *AKSSchemaProperties `pulumi:"properties"`
	ResourceId       *string              `pulumi:"resourceId"`
}


func (val *AKS) Defaults() *AKS {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type AKSResponse struct {
	ComputeLocation    string                       `pulumi:"computeLocation"`
	ComputeType        string                       `pulumi:"computeType"`
	CreatedOn          string                       `pulumi:"createdOn"`
	Description        *string                      `pulumi:"description"`
	DisableLocalAuth   *bool                        `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                         `pulumi:"isAttachedCompute"`
	ModifiedOn         string                       `pulumi:"modifiedOn"`
	Properties         *AKSSchemaResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse      `pulumi:"provisioningErrors"`
	ProvisioningState  string                       `pulumi:"provisioningState"`
	ResourceId         *string                      `pulumi:"resourceId"`
}


func (val *AKSResponse) Defaults() *AKSResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type AKSSchemaProperties struct {
	AgentCount                 *int                        `pulumi:"agentCount"`
	AgentVmSize                *string                     `pulumi:"agentVmSize"`
	AksNetworkingConfiguration *AksNetworkingConfiguration `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                *string                     `pulumi:"clusterFqdn"`
	ClusterPurpose             *string                     `pulumi:"clusterPurpose"`
	LoadBalancerSubnet         *string                     `pulumi:"loadBalancerSubnet"`
	LoadBalancerType           *string                     `pulumi:"loadBalancerType"`
	SslConfiguration           *SslConfiguration           `pulumi:"sslConfiguration"`
}


func (val *AKSSchemaProperties) Defaults() *AKSSchemaProperties {
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

type AKSSchemaResponseProperties struct {
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


func (val *AKSSchemaResponseProperties) Defaults() *AKSSchemaResponseProperties {
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
	CredentialsType string                     `pulumi:"credentialsType"`
	Secrets         AccountKeyDatastoreSecrets `pulumi:"secrets"`
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

type AmlComputeNodeInformationResponseOutput struct{ *pulumi.OutputState }

func (AmlComputeNodeInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeNodeInformationResponse)(nil)).Elem()
}

func (o AmlComputeNodeInformationResponseOutput) ToAmlComputeNodeInformationResponseOutput() AmlComputeNodeInformationResponseOutput {
	return o
}

func (o AmlComputeNodeInformationResponseOutput) ToAmlComputeNodeInformationResponseOutputWithContext(ctx context.Context) AmlComputeNodeInformationResponseOutput {
	return o
}

func (o AmlComputeNodeInformationResponseOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) string { return v.NodeId }).(pulumi.StringOutput)
}

func (o AmlComputeNodeInformationResponseOutput) NodeState() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) string { return v.NodeState }).(pulumi.StringOutput)
}

func (o AmlComputeNodeInformationResponseOutput) Port() pulumi.Float64Output {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) float64 { return v.Port }).(pulumi.Float64Output)
}

func (o AmlComputeNodeInformationResponseOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) string { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

func (o AmlComputeNodeInformationResponseOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) string { return v.PublicIpAddress }).(pulumi.StringOutput)
}

func (o AmlComputeNodeInformationResponseOutput) RunId() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) string { return v.RunId }).(pulumi.StringOutput)
}

type AmlComputeNodeInformationResponseArrayOutput struct{ *pulumi.OutputState }

func (AmlComputeNodeInformationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AmlComputeNodeInformationResponse)(nil)).Elem()
}

func (o AmlComputeNodeInformationResponseArrayOutput) ToAmlComputeNodeInformationResponseArrayOutput() AmlComputeNodeInformationResponseArrayOutput {
	return o
}

func (o AmlComputeNodeInformationResponseArrayOutput) ToAmlComputeNodeInformationResponseArrayOutputWithContext(ctx context.Context) AmlComputeNodeInformationResponseArrayOutput {
	return o
}

func (o AmlComputeNodeInformationResponseArrayOutput) Index(i pulumi.IntInput) AmlComputeNodeInformationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AmlComputeNodeInformationResponse {
		return vs[0].([]AmlComputeNodeInformationResponse)[vs[1].(int)]
	}).(AmlComputeNodeInformationResponseOutput)
}

type AmlComputeProperties struct {
	EnableNodePublicIp          *bool                   `pulumi:"enableNodePublicIp"`
	IsolatedNetwork             *bool                   `pulumi:"isolatedNetwork"`
	OsType                      *string                 `pulumi:"osType"`
	PropertyBag                 interface{}             `pulumi:"propertyBag"`
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

type AmlComputePropertiesResponse struct {
	AllocationState               string                          `pulumi:"allocationState"`
	AllocationStateTransitionTime string                          `pulumi:"allocationStateTransitionTime"`
	CurrentNodeCount              int                             `pulumi:"currentNodeCount"`
	EnableNodePublicIp            *bool                           `pulumi:"enableNodePublicIp"`
	Errors                        []ErrorResponseResponse         `pulumi:"errors"`
	IsolatedNetwork               *bool                           `pulumi:"isolatedNetwork"`
	NodeStateCounts               NodeStateCountsResponse         `pulumi:"nodeStateCounts"`
	OsType                        *string                         `pulumi:"osType"`
	PropertyBag                   interface{}                     `pulumi:"propertyBag"`
	RemoteLoginPortPublicAccess   *string                         `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings                 *ScaleSettingsResponse          `pulumi:"scaleSettings"`
	Subnet                        *ResourceIdResponse             `pulumi:"subnet"`
	TargetNodeCount               int                             `pulumi:"targetNodeCount"`
	UserAccountCredentials        *UserAccountCredentialsResponse `pulumi:"userAccountCredentials"`
	VirtualMachineImage           *VirtualMachineImageResponse    `pulumi:"virtualMachineImage"`
	VmPriority                    *string                         `pulumi:"vmPriority"`
	VmSize                        *string                         `pulumi:"vmSize"`
}


func (val *AmlComputePropertiesResponse) Defaults() *AmlComputePropertiesResponse {
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
	ComputeLocation    string                        `pulumi:"computeLocation"`
	ComputeType        string                        `pulumi:"computeType"`
	CreatedOn          string                        `pulumi:"createdOn"`
	Description        *string                       `pulumi:"description"`
	DisableLocalAuth   *bool                         `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                          `pulumi:"isAttachedCompute"`
	ModifiedOn         string                        `pulumi:"modifiedOn"`
	Properties         *AmlComputePropertiesResponse `pulumi:"properties"`
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

type AzureBlobDatastore struct {
	AccountName                   *string           `pulumi:"accountName"`
	ContainerName                 *string           `pulumi:"containerName"`
	Credentials                   interface{}       `pulumi:"credentials"`
	DatastoreType                 string            `pulumi:"datastoreType"`
	Description                   *string           `pulumi:"description"`
	Endpoint                      *string           `pulumi:"endpoint"`
	Properties                    map[string]string `pulumi:"properties"`
	Protocol                      *string           `pulumi:"protocol"`
	ServiceDataAccessAuthIdentity *string           `pulumi:"serviceDataAccessAuthIdentity"`
	Tags                          map[string]string `pulumi:"tags"`
}


func (val *AzureBlobDatastore) Defaults() *AzureBlobDatastore {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceDataAccessAuthIdentity) {
		serviceDataAccessAuthIdentity_ := "None"
		tmp.ServiceDataAccessAuthIdentity = &serviceDataAccessAuthIdentity_
	}
	return &tmp
}

type AzureBlobDatastoreResponse struct {
	AccountName                   *string           `pulumi:"accountName"`
	ContainerName                 *string           `pulumi:"containerName"`
	Credentials                   interface{}       `pulumi:"credentials"`
	DatastoreType                 string            `pulumi:"datastoreType"`
	Description                   *string           `pulumi:"description"`
	Endpoint                      *string           `pulumi:"endpoint"`
	IsDefault                     bool              `pulumi:"isDefault"`
	Properties                    map[string]string `pulumi:"properties"`
	Protocol                      *string           `pulumi:"protocol"`
	ServiceDataAccessAuthIdentity *string           `pulumi:"serviceDataAccessAuthIdentity"`
	Tags                          map[string]string `pulumi:"tags"`
}


func (val *AzureBlobDatastoreResponse) Defaults() *AzureBlobDatastoreResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceDataAccessAuthIdentity) {
		serviceDataAccessAuthIdentity_ := "None"
		tmp.ServiceDataAccessAuthIdentity = &serviceDataAccessAuthIdentity_
	}
	return &tmp
}

type AzureDataLakeGen1Datastore struct {
	Credentials                   interface{}       `pulumi:"credentials"`
	DatastoreType                 string            `pulumi:"datastoreType"`
	Description                   *string           `pulumi:"description"`
	Properties                    map[string]string `pulumi:"properties"`
	ServiceDataAccessAuthIdentity *string           `pulumi:"serviceDataAccessAuthIdentity"`
	StoreName                     string            `pulumi:"storeName"`
	Tags                          map[string]string `pulumi:"tags"`
}


func (val *AzureDataLakeGen1Datastore) Defaults() *AzureDataLakeGen1Datastore {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceDataAccessAuthIdentity) {
		serviceDataAccessAuthIdentity_ := "None"
		tmp.ServiceDataAccessAuthIdentity = &serviceDataAccessAuthIdentity_
	}
	return &tmp
}

type AzureDataLakeGen1DatastoreResponse struct {
	Credentials                   interface{}       `pulumi:"credentials"`
	DatastoreType                 string            `pulumi:"datastoreType"`
	Description                   *string           `pulumi:"description"`
	IsDefault                     bool              `pulumi:"isDefault"`
	Properties                    map[string]string `pulumi:"properties"`
	ServiceDataAccessAuthIdentity *string           `pulumi:"serviceDataAccessAuthIdentity"`
	StoreName                     string            `pulumi:"storeName"`
	Tags                          map[string]string `pulumi:"tags"`
}


func (val *AzureDataLakeGen1DatastoreResponse) Defaults() *AzureDataLakeGen1DatastoreResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceDataAccessAuthIdentity) {
		serviceDataAccessAuthIdentity_ := "None"
		tmp.ServiceDataAccessAuthIdentity = &serviceDataAccessAuthIdentity_
	}
	return &tmp
}

type AzureDataLakeGen2Datastore struct {
	AccountName                   string            `pulumi:"accountName"`
	Credentials                   interface{}       `pulumi:"credentials"`
	DatastoreType                 string            `pulumi:"datastoreType"`
	Description                   *string           `pulumi:"description"`
	Endpoint                      *string           `pulumi:"endpoint"`
	Filesystem                    string            `pulumi:"filesystem"`
	Properties                    map[string]string `pulumi:"properties"`
	Protocol                      *string           `pulumi:"protocol"`
	ServiceDataAccessAuthIdentity *string           `pulumi:"serviceDataAccessAuthIdentity"`
	Tags                          map[string]string `pulumi:"tags"`
}


func (val *AzureDataLakeGen2Datastore) Defaults() *AzureDataLakeGen2Datastore {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceDataAccessAuthIdentity) {
		serviceDataAccessAuthIdentity_ := "None"
		tmp.ServiceDataAccessAuthIdentity = &serviceDataAccessAuthIdentity_
	}
	return &tmp
}

type AzureDataLakeGen2DatastoreResponse struct {
	AccountName                   string            `pulumi:"accountName"`
	Credentials                   interface{}       `pulumi:"credentials"`
	DatastoreType                 string            `pulumi:"datastoreType"`
	Description                   *string           `pulumi:"description"`
	Endpoint                      *string           `pulumi:"endpoint"`
	Filesystem                    string            `pulumi:"filesystem"`
	IsDefault                     bool              `pulumi:"isDefault"`
	Properties                    map[string]string `pulumi:"properties"`
	Protocol                      *string           `pulumi:"protocol"`
	ServiceDataAccessAuthIdentity *string           `pulumi:"serviceDataAccessAuthIdentity"`
	Tags                          map[string]string `pulumi:"tags"`
}


func (val *AzureDataLakeGen2DatastoreResponse) Defaults() *AzureDataLakeGen2DatastoreResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceDataAccessAuthIdentity) {
		serviceDataAccessAuthIdentity_ := "None"
		tmp.ServiceDataAccessAuthIdentity = &serviceDataAccessAuthIdentity_
	}
	return &tmp
}

type AzureFileDatastore struct {
	AccountName                   string            `pulumi:"accountName"`
	Credentials                   interface{}       `pulumi:"credentials"`
	DatastoreType                 string            `pulumi:"datastoreType"`
	Description                   *string           `pulumi:"description"`
	Endpoint                      *string           `pulumi:"endpoint"`
	FileShareName                 string            `pulumi:"fileShareName"`
	Properties                    map[string]string `pulumi:"properties"`
	Protocol                      *string           `pulumi:"protocol"`
	ServiceDataAccessAuthIdentity *string           `pulumi:"serviceDataAccessAuthIdentity"`
	Tags                          map[string]string `pulumi:"tags"`
}


func (val *AzureFileDatastore) Defaults() *AzureFileDatastore {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceDataAccessAuthIdentity) {
		serviceDataAccessAuthIdentity_ := "None"
		tmp.ServiceDataAccessAuthIdentity = &serviceDataAccessAuthIdentity_
	}
	return &tmp
}

type AzureFileDatastoreResponse struct {
	AccountName                   string            `pulumi:"accountName"`
	Credentials                   interface{}       `pulumi:"credentials"`
	DatastoreType                 string            `pulumi:"datastoreType"`
	Description                   *string           `pulumi:"description"`
	Endpoint                      *string           `pulumi:"endpoint"`
	FileShareName                 string            `pulumi:"fileShareName"`
	IsDefault                     bool              `pulumi:"isDefault"`
	Properties                    map[string]string `pulumi:"properties"`
	Protocol                      *string           `pulumi:"protocol"`
	ServiceDataAccessAuthIdentity *string           `pulumi:"serviceDataAccessAuthIdentity"`
	Tags                          map[string]string `pulumi:"tags"`
}


func (val *AzureFileDatastoreResponse) Defaults() *AzureFileDatastoreResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceDataAccessAuthIdentity) {
		serviceDataAccessAuthIdentity_ := "None"
		tmp.ServiceDataAccessAuthIdentity = &serviceDataAccessAuthIdentity_
	}
	return &tmp
}

type BanditPolicy struct {
	DelayEvaluation    *int     `pulumi:"delayEvaluation"`
	EvaluationInterval *int     `pulumi:"evaluationInterval"`
	PolicyType         string   `pulumi:"policyType"`
	SlackAmount        *float64 `pulumi:"slackAmount"`
	SlackFactor        *float64 `pulumi:"slackFactor"`
}


func (val *BanditPolicy) Defaults() *BanditPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DelayEvaluation) {
		delayEvaluation_ := 0
		tmp.DelayEvaluation = &delayEvaluation_
	}
	if isZero(tmp.EvaluationInterval) {
		evaluationInterval_ := 0
		tmp.EvaluationInterval = &evaluationInterval_
	}
	if isZero(tmp.SlackAmount) {
		slackAmount_ := 0.0
		tmp.SlackAmount = &slackAmount_
	}
	if isZero(tmp.SlackFactor) {
		slackFactor_ := 0.0
		tmp.SlackFactor = &slackFactor_
	}
	return &tmp
}

type BanditPolicyResponse struct {
	DelayEvaluation    *int     `pulumi:"delayEvaluation"`
	EvaluationInterval *int     `pulumi:"evaluationInterval"`
	PolicyType         string   `pulumi:"policyType"`
	SlackAmount        *float64 `pulumi:"slackAmount"`
	SlackFactor        *float64 `pulumi:"slackFactor"`
}


func (val *BanditPolicyResponse) Defaults() *BanditPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DelayEvaluation) {
		delayEvaluation_ := 0
		tmp.DelayEvaluation = &delayEvaluation_
	}
	if isZero(tmp.EvaluationInterval) {
		evaluationInterval_ := 0
		tmp.EvaluationInterval = &evaluationInterval_
	}
	if isZero(tmp.SlackAmount) {
		slackAmount_ := 0.0
		tmp.SlackAmount = &slackAmount_
	}
	if isZero(tmp.SlackFactor) {
		slackFactor_ := 0.0
		tmp.SlackFactor = &slackFactor_
	}
	return &tmp
}

type BatchDeploymentType struct {
	CodeConfiguration         *CodeConfiguration     `pulumi:"codeConfiguration"`
	Compute                   *string                `pulumi:"compute"`
	Description               *string                `pulumi:"description"`
	EnvironmentId             *string                `pulumi:"environmentId"`
	EnvironmentVariables      map[string]string      `pulumi:"environmentVariables"`
	ErrorThreshold            *int                   `pulumi:"errorThreshold"`
	LoggingLevel              *string                `pulumi:"loggingLevel"`
	MaxConcurrencyPerInstance *int                   `pulumi:"maxConcurrencyPerInstance"`
	MiniBatchSize             *float64               `pulumi:"miniBatchSize"`
	Model                     interface{}            `pulumi:"model"`
	OutputAction              *string                `pulumi:"outputAction"`
	OutputFileName            *string                `pulumi:"outputFileName"`
	Properties                map[string]string      `pulumi:"properties"`
	Resources                 *ResourceConfiguration `pulumi:"resources"`
	RetrySettings             *BatchRetrySettings    `pulumi:"retrySettings"`
}


func (val *BatchDeploymentType) Defaults() *BatchDeploymentType {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ErrorThreshold) {
		errorThreshold_ := -1
		tmp.ErrorThreshold = &errorThreshold_
	}
	if isZero(tmp.LoggingLevel) {
		loggingLevel_ := "Info"
		tmp.LoggingLevel = &loggingLevel_
	}
	if isZero(tmp.MaxConcurrencyPerInstance) {
		maxConcurrencyPerInstance_ := 1
		tmp.MaxConcurrencyPerInstance = &maxConcurrencyPerInstance_
	}
	if isZero(tmp.MiniBatchSize) {
		miniBatchSize_ := 10.0
		tmp.MiniBatchSize = &miniBatchSize_
	}
	if isZero(tmp.OutputAction) {
		outputAction_ := "AppendRow"
		tmp.OutputAction = &outputAction_
	}
	if isZero(tmp.OutputFileName) {
		outputFileName_ := "predictions.csv"
		tmp.OutputFileName = &outputFileName_
	}
	tmp.Resources = tmp.Resources.Defaults()

	tmp.RetrySettings = tmp.RetrySettings.Defaults()

	return &tmp
}





type BatchDeploymentTypeInput interface {
	pulumi.Input

	ToBatchDeploymentTypeOutput() BatchDeploymentTypeOutput
	ToBatchDeploymentTypeOutputWithContext(context.Context) BatchDeploymentTypeOutput
}

type BatchDeploymentTypeArgs struct {
	CodeConfiguration         CodeConfigurationPtrInput     `pulumi:"codeConfiguration"`
	Compute                   pulumi.StringPtrInput         `pulumi:"compute"`
	Description               pulumi.StringPtrInput         `pulumi:"description"`
	EnvironmentId             pulumi.StringPtrInput         `pulumi:"environmentId"`
	EnvironmentVariables      pulumi.StringMapInput         `pulumi:"environmentVariables"`
	ErrorThreshold            pulumi.IntPtrInput            `pulumi:"errorThreshold"`
	LoggingLevel              pulumi.StringPtrInput         `pulumi:"loggingLevel"`
	MaxConcurrencyPerInstance pulumi.IntPtrInput            `pulumi:"maxConcurrencyPerInstance"`
	MiniBatchSize             pulumi.Float64PtrInput        `pulumi:"miniBatchSize"`
	Model                     pulumi.Input                  `pulumi:"model"`
	OutputAction              pulumi.StringPtrInput         `pulumi:"outputAction"`
	OutputFileName            pulumi.StringPtrInput         `pulumi:"outputFileName"`
	Properties                pulumi.StringMapInput         `pulumi:"properties"`
	Resources                 ResourceConfigurationPtrInput `pulumi:"resources"`
	RetrySettings             BatchRetrySettingsPtrInput    `pulumi:"retrySettings"`
}


func (val *BatchDeploymentTypeArgs) Defaults() *BatchDeploymentTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ErrorThreshold) {
		tmp.ErrorThreshold = pulumi.IntPtr(-1)
	}
	if isZero(tmp.LoggingLevel) {
		tmp.LoggingLevel = pulumi.StringPtr("Info")
	}
	if isZero(tmp.MaxConcurrencyPerInstance) {
		tmp.MaxConcurrencyPerInstance = pulumi.IntPtr(1)
	}
	if isZero(tmp.MiniBatchSize) {
		tmp.MiniBatchSize = pulumi.Float64Ptr(10.0)
	}
	if isZero(tmp.OutputAction) {
		tmp.OutputAction = pulumi.StringPtr("AppendRow")
	}
	if isZero(tmp.OutputFileName) {
		tmp.OutputFileName = pulumi.StringPtr("predictions.csv")
	}

	return &tmp
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

func (o BatchDeploymentTypeOutput) Compute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *string { return v.Compute }).(pulumi.StringPtrOutput)
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

func (o BatchDeploymentTypeOutput) MaxConcurrencyPerInstance() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *int { return v.MaxConcurrencyPerInstance }).(pulumi.IntPtrOutput)
}

func (o BatchDeploymentTypeOutput) MiniBatchSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *float64 { return v.MiniBatchSize }).(pulumi.Float64PtrOutput)
}

func (o BatchDeploymentTypeOutput) Model() pulumi.AnyOutput {
	return o.ApplyT(func(v BatchDeploymentType) interface{} { return v.Model }).(pulumi.AnyOutput)
}

func (o BatchDeploymentTypeOutput) OutputAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *string { return v.OutputAction }).(pulumi.StringPtrOutput)
}

func (o BatchDeploymentTypeOutput) OutputFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *string { return v.OutputFileName }).(pulumi.StringPtrOutput)
}

func (o BatchDeploymentTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v BatchDeploymentType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o BatchDeploymentTypeOutput) Resources() ResourceConfigurationPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *ResourceConfiguration { return v.Resources }).(ResourceConfigurationPtrOutput)
}

func (o BatchDeploymentTypeOutput) RetrySettings() BatchRetrySettingsPtrOutput {
	return o.ApplyT(func(v BatchDeploymentType) *BatchRetrySettings { return v.RetrySettings }).(BatchRetrySettingsPtrOutput)
}

type BatchDeploymentResponse struct {
	CodeConfiguration         *CodeConfigurationResponse     `pulumi:"codeConfiguration"`
	Compute                   *string                        `pulumi:"compute"`
	Description               *string                        `pulumi:"description"`
	EnvironmentId             *string                        `pulumi:"environmentId"`
	EnvironmentVariables      map[string]string              `pulumi:"environmentVariables"`
	ErrorThreshold            *int                           `pulumi:"errorThreshold"`
	LoggingLevel              *string                        `pulumi:"loggingLevel"`
	MaxConcurrencyPerInstance *int                           `pulumi:"maxConcurrencyPerInstance"`
	MiniBatchSize             *float64                       `pulumi:"miniBatchSize"`
	Model                     interface{}                    `pulumi:"model"`
	OutputAction              *string                        `pulumi:"outputAction"`
	OutputFileName            *string                        `pulumi:"outputFileName"`
	Properties                map[string]string              `pulumi:"properties"`
	ProvisioningState         string                         `pulumi:"provisioningState"`
	Resources                 *ResourceConfigurationResponse `pulumi:"resources"`
	RetrySettings             *BatchRetrySettingsResponse    `pulumi:"retrySettings"`
}


func (val *BatchDeploymentResponse) Defaults() *BatchDeploymentResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ErrorThreshold) {
		errorThreshold_ := -1
		tmp.ErrorThreshold = &errorThreshold_
	}
	if isZero(tmp.LoggingLevel) {
		loggingLevel_ := "Info"
		tmp.LoggingLevel = &loggingLevel_
	}
	if isZero(tmp.MaxConcurrencyPerInstance) {
		maxConcurrencyPerInstance_ := 1
		tmp.MaxConcurrencyPerInstance = &maxConcurrencyPerInstance_
	}
	if isZero(tmp.MiniBatchSize) {
		miniBatchSize_ := 10.0
		tmp.MiniBatchSize = &miniBatchSize_
	}
	if isZero(tmp.OutputAction) {
		outputAction_ := "AppendRow"
		tmp.OutputAction = &outputAction_
	}
	if isZero(tmp.OutputFileName) {
		outputFileName_ := "predictions.csv"
		tmp.OutputFileName = &outputFileName_
	}
	tmp.Resources = tmp.Resources.Defaults()

	tmp.RetrySettings = tmp.RetrySettings.Defaults()

	return &tmp
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

func (o BatchDeploymentResponseOutput) Compute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *string { return v.Compute }).(pulumi.StringPtrOutput)
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

func (o BatchDeploymentResponseOutput) MaxConcurrencyPerInstance() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *int { return v.MaxConcurrencyPerInstance }).(pulumi.IntPtrOutput)
}

func (o BatchDeploymentResponseOutput) MiniBatchSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *float64 { return v.MiniBatchSize }).(pulumi.Float64PtrOutput)
}

func (o BatchDeploymentResponseOutput) Model() pulumi.AnyOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) interface{} { return v.Model }).(pulumi.AnyOutput)
}

func (o BatchDeploymentResponseOutput) OutputAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *string { return v.OutputAction }).(pulumi.StringPtrOutput)
}

func (o BatchDeploymentResponseOutput) OutputFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *string { return v.OutputFileName }).(pulumi.StringPtrOutput)
}

func (o BatchDeploymentResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o BatchDeploymentResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o BatchDeploymentResponseOutput) Resources() ResourceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *ResourceConfigurationResponse { return v.Resources }).(ResourceConfigurationResponsePtrOutput)
}

func (o BatchDeploymentResponseOutput) RetrySettings() BatchRetrySettingsResponsePtrOutput {
	return o.ApplyT(func(v BatchDeploymentResponse) *BatchRetrySettingsResponse { return v.RetrySettings }).(BatchRetrySettingsResponsePtrOutput)
}

type BatchEndpointType struct {
	AuthMode    string                 `pulumi:"authMode"`
	Defaults    *BatchEndpointDefaults `pulumi:"defaults"`
	Description *string                `pulumi:"description"`
	Keys        *EndpointAuthKeys      `pulumi:"keys"`
	Properties  map[string]string      `pulumi:"properties"`
}





type BatchEndpointTypeInput interface {
	pulumi.Input

	ToBatchEndpointTypeOutput() BatchEndpointTypeOutput
	ToBatchEndpointTypeOutputWithContext(context.Context) BatchEndpointTypeOutput
}

type BatchEndpointTypeArgs struct {
	AuthMode    pulumi.StringInput            `pulumi:"authMode"`
	Defaults    BatchEndpointDefaultsPtrInput `pulumi:"defaults"`
	Description pulumi.StringPtrInput         `pulumi:"description"`
	Keys        EndpointAuthKeysPtrInput      `pulumi:"keys"`
	Properties  pulumi.StringMapInput         `pulumi:"properties"`
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

func (o BatchEndpointTypeOutput) AuthMode() pulumi.StringOutput {
	return o.ApplyT(func(v BatchEndpointType) string { return v.AuthMode }).(pulumi.StringOutput)
}

func (o BatchEndpointTypeOutput) Defaults() BatchEndpointDefaultsPtrOutput {
	return o.ApplyT(func(v BatchEndpointType) *BatchEndpointDefaults { return v.Defaults }).(BatchEndpointDefaultsPtrOutput)
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

type BatchEndpointDefaults struct {
	DeploymentName *string `pulumi:"deploymentName"`
}





type BatchEndpointDefaultsInput interface {
	pulumi.Input

	ToBatchEndpointDefaultsOutput() BatchEndpointDefaultsOutput
	ToBatchEndpointDefaultsOutputWithContext(context.Context) BatchEndpointDefaultsOutput
}

type BatchEndpointDefaultsArgs struct {
	DeploymentName pulumi.StringPtrInput `pulumi:"deploymentName"`
}

func (BatchEndpointDefaultsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchEndpointDefaults)(nil)).Elem()
}

func (i BatchEndpointDefaultsArgs) ToBatchEndpointDefaultsOutput() BatchEndpointDefaultsOutput {
	return i.ToBatchEndpointDefaultsOutputWithContext(context.Background())
}

func (i BatchEndpointDefaultsArgs) ToBatchEndpointDefaultsOutputWithContext(ctx context.Context) BatchEndpointDefaultsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchEndpointDefaultsOutput)
}

func (i BatchEndpointDefaultsArgs) ToBatchEndpointDefaultsPtrOutput() BatchEndpointDefaultsPtrOutput {
	return i.ToBatchEndpointDefaultsPtrOutputWithContext(context.Background())
}

func (i BatchEndpointDefaultsArgs) ToBatchEndpointDefaultsPtrOutputWithContext(ctx context.Context) BatchEndpointDefaultsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchEndpointDefaultsOutput).ToBatchEndpointDefaultsPtrOutputWithContext(ctx)
}









type BatchEndpointDefaultsPtrInput interface {
	pulumi.Input

	ToBatchEndpointDefaultsPtrOutput() BatchEndpointDefaultsPtrOutput
	ToBatchEndpointDefaultsPtrOutputWithContext(context.Context) BatchEndpointDefaultsPtrOutput
}

type batchEndpointDefaultsPtrType BatchEndpointDefaultsArgs

func BatchEndpointDefaultsPtr(v *BatchEndpointDefaultsArgs) BatchEndpointDefaultsPtrInput {
	return (*batchEndpointDefaultsPtrType)(v)
}

func (*batchEndpointDefaultsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchEndpointDefaults)(nil)).Elem()
}

func (i *batchEndpointDefaultsPtrType) ToBatchEndpointDefaultsPtrOutput() BatchEndpointDefaultsPtrOutput {
	return i.ToBatchEndpointDefaultsPtrOutputWithContext(context.Background())
}

func (i *batchEndpointDefaultsPtrType) ToBatchEndpointDefaultsPtrOutputWithContext(ctx context.Context) BatchEndpointDefaultsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchEndpointDefaultsPtrOutput)
}

type BatchEndpointDefaultsOutput struct{ *pulumi.OutputState }

func (BatchEndpointDefaultsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchEndpointDefaults)(nil)).Elem()
}

func (o BatchEndpointDefaultsOutput) ToBatchEndpointDefaultsOutput() BatchEndpointDefaultsOutput {
	return o
}

func (o BatchEndpointDefaultsOutput) ToBatchEndpointDefaultsOutputWithContext(ctx context.Context) BatchEndpointDefaultsOutput {
	return o
}

func (o BatchEndpointDefaultsOutput) ToBatchEndpointDefaultsPtrOutput() BatchEndpointDefaultsPtrOutput {
	return o.ToBatchEndpointDefaultsPtrOutputWithContext(context.Background())
}

func (o BatchEndpointDefaultsOutput) ToBatchEndpointDefaultsPtrOutputWithContext(ctx context.Context) BatchEndpointDefaultsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BatchEndpointDefaults) *BatchEndpointDefaults {
		return &v
	}).(BatchEndpointDefaultsPtrOutput)
}

func (o BatchEndpointDefaultsOutput) DeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchEndpointDefaults) *string { return v.DeploymentName }).(pulumi.StringPtrOutput)
}

type BatchEndpointDefaultsPtrOutput struct{ *pulumi.OutputState }

func (BatchEndpointDefaultsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchEndpointDefaults)(nil)).Elem()
}

func (o BatchEndpointDefaultsPtrOutput) ToBatchEndpointDefaultsPtrOutput() BatchEndpointDefaultsPtrOutput {
	return o
}

func (o BatchEndpointDefaultsPtrOutput) ToBatchEndpointDefaultsPtrOutputWithContext(ctx context.Context) BatchEndpointDefaultsPtrOutput {
	return o
}

func (o BatchEndpointDefaultsPtrOutput) Elem() BatchEndpointDefaultsOutput {
	return o.ApplyT(func(v *BatchEndpointDefaults) BatchEndpointDefaults {
		if v != nil {
			return *v
		}
		var ret BatchEndpointDefaults
		return ret
	}).(BatchEndpointDefaultsOutput)
}

func (o BatchEndpointDefaultsPtrOutput) DeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchEndpointDefaults) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentName
	}).(pulumi.StringPtrOutput)
}

type BatchEndpointDefaultsResponse struct {
	DeploymentName *string `pulumi:"deploymentName"`
}

type BatchEndpointDefaultsResponseOutput struct{ *pulumi.OutputState }

func (BatchEndpointDefaultsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchEndpointDefaultsResponse)(nil)).Elem()
}

func (o BatchEndpointDefaultsResponseOutput) ToBatchEndpointDefaultsResponseOutput() BatchEndpointDefaultsResponseOutput {
	return o
}

func (o BatchEndpointDefaultsResponseOutput) ToBatchEndpointDefaultsResponseOutputWithContext(ctx context.Context) BatchEndpointDefaultsResponseOutput {
	return o
}

func (o BatchEndpointDefaultsResponseOutput) DeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchEndpointDefaultsResponse) *string { return v.DeploymentName }).(pulumi.StringPtrOutput)
}

type BatchEndpointDefaultsResponsePtrOutput struct{ *pulumi.OutputState }

func (BatchEndpointDefaultsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchEndpointDefaultsResponse)(nil)).Elem()
}

func (o BatchEndpointDefaultsResponsePtrOutput) ToBatchEndpointDefaultsResponsePtrOutput() BatchEndpointDefaultsResponsePtrOutput {
	return o
}

func (o BatchEndpointDefaultsResponsePtrOutput) ToBatchEndpointDefaultsResponsePtrOutputWithContext(ctx context.Context) BatchEndpointDefaultsResponsePtrOutput {
	return o
}

func (o BatchEndpointDefaultsResponsePtrOutput) Elem() BatchEndpointDefaultsResponseOutput {
	return o.ApplyT(func(v *BatchEndpointDefaultsResponse) BatchEndpointDefaultsResponse {
		if v != nil {
			return *v
		}
		var ret BatchEndpointDefaultsResponse
		return ret
	}).(BatchEndpointDefaultsResponseOutput)
}

func (o BatchEndpointDefaultsResponsePtrOutput) DeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchEndpointDefaultsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentName
	}).(pulumi.StringPtrOutput)
}

type BatchEndpointResponse struct {
	AuthMode          string                         `pulumi:"authMode"`
	Defaults          *BatchEndpointDefaultsResponse `pulumi:"defaults"`
	Description       *string                        `pulumi:"description"`
	Properties        map[string]string              `pulumi:"properties"`
	ProvisioningState string                         `pulumi:"provisioningState"`
	ScoringUri        string                         `pulumi:"scoringUri"`
	SwaggerUri        string                         `pulumi:"swaggerUri"`
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

func (o BatchEndpointResponseOutput) AuthMode() pulumi.StringOutput {
	return o.ApplyT(func(v BatchEndpointResponse) string { return v.AuthMode }).(pulumi.StringOutput)
}

func (o BatchEndpointResponseOutput) Defaults() BatchEndpointDefaultsResponsePtrOutput {
	return o.ApplyT(func(v BatchEndpointResponse) *BatchEndpointDefaultsResponse { return v.Defaults }).(BatchEndpointDefaultsResponsePtrOutput)
}

func (o BatchEndpointResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchEndpointResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BatchEndpointResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v BatchEndpointResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o BatchEndpointResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v BatchEndpointResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o BatchEndpointResponseOutput) ScoringUri() pulumi.StringOutput {
	return o.ApplyT(func(v BatchEndpointResponse) string { return v.ScoringUri }).(pulumi.StringOutput)
}

func (o BatchEndpointResponseOutput) SwaggerUri() pulumi.StringOutput {
	return o.ApplyT(func(v BatchEndpointResponse) string { return v.SwaggerUri }).(pulumi.StringOutput)
}

type BatchRetrySettings struct {
	MaxRetries *int    `pulumi:"maxRetries"`
	Timeout    *string `pulumi:"timeout"`
}


func (val *BatchRetrySettings) Defaults() *BatchRetrySettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxRetries) {
		maxRetries_ := 3
		tmp.MaxRetries = &maxRetries_
	}
	if isZero(tmp.Timeout) {
		timeout_ := "PT30S"
		tmp.Timeout = &timeout_
	}
	return &tmp
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


func (val *BatchRetrySettingsArgs) Defaults() *BatchRetrySettingsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxRetries) {
		tmp.MaxRetries = pulumi.IntPtr(3)
	}
	if isZero(tmp.Timeout) {
		tmp.Timeout = pulumi.StringPtr("PT30S")
	}
	return &tmp
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


func (val *BatchRetrySettingsResponse) Defaults() *BatchRetrySettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxRetries) {
		maxRetries_ := 3
		tmp.MaxRetries = &maxRetries_
	}
	if isZero(tmp.Timeout) {
		timeout_ := "PT30S"
		tmp.Timeout = &timeout_
	}
	return &tmp
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

type BayesianSamplingAlgorithm struct {
	SamplingAlgorithmType string `pulumi:"samplingAlgorithmType"`
}

type BayesianSamplingAlgorithmResponse struct {
	SamplingAlgorithmType string `pulumi:"samplingAlgorithmType"`
}

type BuildContext struct {
	ContextUri     string  `pulumi:"contextUri"`
	DockerfilePath *string `pulumi:"dockerfilePath"`
}


func (val *BuildContext) Defaults() *BuildContext {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DockerfilePath) {
		dockerfilePath_ := "Dockerfile"
		tmp.DockerfilePath = &dockerfilePath_
	}
	return &tmp
}





type BuildContextInput interface {
	pulumi.Input

	ToBuildContextOutput() BuildContextOutput
	ToBuildContextOutputWithContext(context.Context) BuildContextOutput
}

type BuildContextArgs struct {
	ContextUri     pulumi.StringInput    `pulumi:"contextUri"`
	DockerfilePath pulumi.StringPtrInput `pulumi:"dockerfilePath"`
}


func (val *BuildContextArgs) Defaults() *BuildContextArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DockerfilePath) {
		tmp.DockerfilePath = pulumi.StringPtr("Dockerfile")
	}
	return &tmp
}
func (BuildContextArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildContext)(nil)).Elem()
}

func (i BuildContextArgs) ToBuildContextOutput() BuildContextOutput {
	return i.ToBuildContextOutputWithContext(context.Background())
}

func (i BuildContextArgs) ToBuildContextOutputWithContext(ctx context.Context) BuildContextOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildContextOutput)
}

func (i BuildContextArgs) ToBuildContextPtrOutput() BuildContextPtrOutput {
	return i.ToBuildContextPtrOutputWithContext(context.Background())
}

func (i BuildContextArgs) ToBuildContextPtrOutputWithContext(ctx context.Context) BuildContextPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildContextOutput).ToBuildContextPtrOutputWithContext(ctx)
}









type BuildContextPtrInput interface {
	pulumi.Input

	ToBuildContextPtrOutput() BuildContextPtrOutput
	ToBuildContextPtrOutputWithContext(context.Context) BuildContextPtrOutput
}

type buildContextPtrType BuildContextArgs

func BuildContextPtr(v *BuildContextArgs) BuildContextPtrInput {
	return (*buildContextPtrType)(v)
}

func (*buildContextPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildContext)(nil)).Elem()
}

func (i *buildContextPtrType) ToBuildContextPtrOutput() BuildContextPtrOutput {
	return i.ToBuildContextPtrOutputWithContext(context.Background())
}

func (i *buildContextPtrType) ToBuildContextPtrOutputWithContext(ctx context.Context) BuildContextPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuildContextPtrOutput)
}

type BuildContextOutput struct{ *pulumi.OutputState }

func (BuildContextOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildContext)(nil)).Elem()
}

func (o BuildContextOutput) ToBuildContextOutput() BuildContextOutput {
	return o
}

func (o BuildContextOutput) ToBuildContextOutputWithContext(ctx context.Context) BuildContextOutput {
	return o
}

func (o BuildContextOutput) ToBuildContextPtrOutput() BuildContextPtrOutput {
	return o.ToBuildContextPtrOutputWithContext(context.Background())
}

func (o BuildContextOutput) ToBuildContextPtrOutputWithContext(ctx context.Context) BuildContextPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BuildContext) *BuildContext {
		return &v
	}).(BuildContextPtrOutput)
}

func (o BuildContextOutput) ContextUri() pulumi.StringOutput {
	return o.ApplyT(func(v BuildContext) string { return v.ContextUri }).(pulumi.StringOutput)
}

func (o BuildContextOutput) DockerfilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildContext) *string { return v.DockerfilePath }).(pulumi.StringPtrOutput)
}

type BuildContextPtrOutput struct{ *pulumi.OutputState }

func (BuildContextPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildContext)(nil)).Elem()
}

func (o BuildContextPtrOutput) ToBuildContextPtrOutput() BuildContextPtrOutput {
	return o
}

func (o BuildContextPtrOutput) ToBuildContextPtrOutputWithContext(ctx context.Context) BuildContextPtrOutput {
	return o
}

func (o BuildContextPtrOutput) Elem() BuildContextOutput {
	return o.ApplyT(func(v *BuildContext) BuildContext {
		if v != nil {
			return *v
		}
		var ret BuildContext
		return ret
	}).(BuildContextOutput)
}

func (o BuildContextPtrOutput) ContextUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildContext) *string {
		if v == nil {
			return nil
		}
		return &v.ContextUri
	}).(pulumi.StringPtrOutput)
}

func (o BuildContextPtrOutput) DockerfilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildContext) *string {
		if v == nil {
			return nil
		}
		return v.DockerfilePath
	}).(pulumi.StringPtrOutput)
}

type BuildContextResponse struct {
	ContextUri     string  `pulumi:"contextUri"`
	DockerfilePath *string `pulumi:"dockerfilePath"`
}


func (val *BuildContextResponse) Defaults() *BuildContextResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DockerfilePath) {
		dockerfilePath_ := "Dockerfile"
		tmp.DockerfilePath = &dockerfilePath_
	}
	return &tmp
}

type BuildContextResponseOutput struct{ *pulumi.OutputState }

func (BuildContextResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildContextResponse)(nil)).Elem()
}

func (o BuildContextResponseOutput) ToBuildContextResponseOutput() BuildContextResponseOutput {
	return o
}

func (o BuildContextResponseOutput) ToBuildContextResponseOutputWithContext(ctx context.Context) BuildContextResponseOutput {
	return o
}

func (o BuildContextResponseOutput) ContextUri() pulumi.StringOutput {
	return o.ApplyT(func(v BuildContextResponse) string { return v.ContextUri }).(pulumi.StringOutput)
}

func (o BuildContextResponseOutput) DockerfilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BuildContextResponse) *string { return v.DockerfilePath }).(pulumi.StringPtrOutput)
}

type BuildContextResponsePtrOutput struct{ *pulumi.OutputState }

func (BuildContextResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuildContextResponse)(nil)).Elem()
}

func (o BuildContextResponsePtrOutput) ToBuildContextResponsePtrOutput() BuildContextResponsePtrOutput {
	return o
}

func (o BuildContextResponsePtrOutput) ToBuildContextResponsePtrOutputWithContext(ctx context.Context) BuildContextResponsePtrOutput {
	return o
}

func (o BuildContextResponsePtrOutput) Elem() BuildContextResponseOutput {
	return o.ApplyT(func(v *BuildContextResponse) BuildContextResponse {
		if v != nil {
			return *v
		}
		var ret BuildContextResponse
		return ret
	}).(BuildContextResponseOutput)
}

func (o BuildContextResponsePtrOutput) ContextUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildContextResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ContextUri
	}).(pulumi.StringPtrOutput)
}

func (o BuildContextResponsePtrOutput) DockerfilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BuildContextResponse) *string {
		if v == nil {
			return nil
		}
		return v.DockerfilePath
	}).(pulumi.StringPtrOutput)
}

type CertificateDatastoreCredentials struct {
	AuthorityUrl    *string                     `pulumi:"authorityUrl"`
	ClientId        string                      `pulumi:"clientId"`
	CredentialsType string                      `pulumi:"credentialsType"`
	ResourceUrl     *string                     `pulumi:"resourceUrl"`
	Secrets         CertificateDatastoreSecrets `pulumi:"secrets"`
	TenantId        string                      `pulumi:"tenantId"`
	Thumbprint      string                      `pulumi:"thumbprint"`
}

type CertificateDatastoreCredentialsResponse struct {
	AuthorityUrl    *string `pulumi:"authorityUrl"`
	ClientId        string  `pulumi:"clientId"`
	CredentialsType string  `pulumi:"credentialsType"`
	ResourceUrl     *string `pulumi:"resourceUrl"`
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
	IsArchived  *bool             `pulumi:"isArchived"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}


func (val *CodeContainerType) Defaults() *CodeContainerType {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}





type CodeContainerTypeInput interface {
	pulumi.Input

	ToCodeContainerTypeOutput() CodeContainerTypeOutput
	ToCodeContainerTypeOutputWithContext(context.Context) CodeContainerTypeOutput
}

type CodeContainerTypeArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	IsArchived  pulumi.BoolPtrInput   `pulumi:"isArchived"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}


func (val *CodeContainerTypeArgs) Defaults() *CodeContainerTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		tmp.IsArchived = pulumi.BoolPtr(false)
	}
	return &tmp
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

func (o CodeContainerTypeOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodeContainerType) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o CodeContainerTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeContainerType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CodeContainerTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeContainerType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type CodeContainerResponse struct {
	Description   *string           `pulumi:"description"`
	IsArchived    *bool             `pulumi:"isArchived"`
	LatestVersion string            `pulumi:"latestVersion"`
	NextVersion   string            `pulumi:"nextVersion"`
	Properties    map[string]string `pulumi:"properties"`
	Tags          map[string]string `pulumi:"tags"`
}


func (val *CodeContainerResponse) Defaults() *CodeContainerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
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

func (o CodeContainerResponseOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodeContainerResponse) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o CodeContainerResponseOutput) LatestVersion() pulumi.StringOutput {
	return o.ApplyT(func(v CodeContainerResponse) string { return v.LatestVersion }).(pulumi.StringOutput)
}

func (o CodeContainerResponseOutput) NextVersion() pulumi.StringOutput {
	return o.ApplyT(func(v CodeContainerResponse) string { return v.NextVersion }).(pulumi.StringOutput)
}

func (o CodeContainerResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeContainerResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CodeContainerResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeContainerResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type CodeVersionType struct {
	CodeUri     *string           `pulumi:"codeUri"`
	Description *string           `pulumi:"description"`
	IsAnonymous *bool             `pulumi:"isAnonymous"`
	IsArchived  *bool             `pulumi:"isArchived"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}


func (val *CodeVersionType) Defaults() *CodeVersionType {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}





type CodeVersionTypeInput interface {
	pulumi.Input

	ToCodeVersionTypeOutput() CodeVersionTypeOutput
	ToCodeVersionTypeOutputWithContext(context.Context) CodeVersionTypeOutput
}

type CodeVersionTypeArgs struct {
	CodeUri     pulumi.StringPtrInput `pulumi:"codeUri"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	IsAnonymous pulumi.BoolPtrInput   `pulumi:"isAnonymous"`
	IsArchived  pulumi.BoolPtrInput   `pulumi:"isArchived"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}


func (val *CodeVersionTypeArgs) Defaults() *CodeVersionTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		tmp.IsAnonymous = pulumi.BoolPtr(false)
	}
	if isZero(tmp.IsArchived) {
		tmp.IsArchived = pulumi.BoolPtr(false)
	}
	return &tmp
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

func (o CodeVersionTypeOutput) CodeUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeVersionType) *string { return v.CodeUri }).(pulumi.StringPtrOutput)
}

func (o CodeVersionTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeVersionType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CodeVersionTypeOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodeVersionType) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o CodeVersionTypeOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodeVersionType) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o CodeVersionTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeVersionType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CodeVersionTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeVersionType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type CodeVersionResponse struct {
	CodeUri     *string           `pulumi:"codeUri"`
	Description *string           `pulumi:"description"`
	IsAnonymous *bool             `pulumi:"isAnonymous"`
	IsArchived  *bool             `pulumi:"isArchived"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}


func (val *CodeVersionResponse) Defaults() *CodeVersionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
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

func (o CodeVersionResponseOutput) CodeUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeVersionResponse) *string { return v.CodeUri }).(pulumi.StringPtrOutput)
}

func (o CodeVersionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodeVersionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CodeVersionResponseOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodeVersionResponse) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o CodeVersionResponseOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodeVersionResponse) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o CodeVersionResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeVersionResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CodeVersionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v CodeVersionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type CommandJob struct {
	CodeId               *string                `pulumi:"codeId"`
	Command              string                 `pulumi:"command"`
	ComputeId            *string                `pulumi:"computeId"`
	Description          *string                `pulumi:"description"`
	DisplayName          *string                `pulumi:"displayName"`
	Distribution         interface{}            `pulumi:"distribution"`
	EnvironmentId        string                 `pulumi:"environmentId"`
	EnvironmentVariables map[string]string      `pulumi:"environmentVariables"`
	ExperimentName       *string                `pulumi:"experimentName"`
	Identity             interface{}            `pulumi:"identity"`
	Inputs               map[string]interface{} `pulumi:"inputs"`
	IsArchived           *bool                  `pulumi:"isArchived"`
	JobType              string                 `pulumi:"jobType"`
	Limits               *CommandJobLimits      `pulumi:"limits"`
	Outputs              map[string]interface{} `pulumi:"outputs"`
	Properties           map[string]string      `pulumi:"properties"`
	Resources            *ResourceConfiguration `pulumi:"resources"`
	Services             map[string]JobService  `pulumi:"services"`
	Tags                 map[string]string      `pulumi:"tags"`
}


func (val *CommandJob) Defaults() *CommandJob {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExperimentName) {
		experimentName_ := "Default"
		tmp.ExperimentName = &experimentName_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	tmp.Resources = tmp.Resources.Defaults()

	return &tmp
}

type CommandJobLimits struct {
	JobLimitsType string  `pulumi:"jobLimitsType"`
	Timeout       *string `pulumi:"timeout"`
}

type CommandJobLimitsResponse struct {
	JobLimitsType string  `pulumi:"jobLimitsType"`
	Timeout       *string `pulumi:"timeout"`
}

type CommandJobResponse struct {
	CodeId               *string                        `pulumi:"codeId"`
	Command              string                         `pulumi:"command"`
	ComputeId            *string                        `pulumi:"computeId"`
	Description          *string                        `pulumi:"description"`
	DisplayName          *string                        `pulumi:"displayName"`
	Distribution         interface{}                    `pulumi:"distribution"`
	EnvironmentId        string                         `pulumi:"environmentId"`
	EnvironmentVariables map[string]string              `pulumi:"environmentVariables"`
	ExperimentName       *string                        `pulumi:"experimentName"`
	Identity             interface{}                    `pulumi:"identity"`
	Inputs               map[string]interface{}         `pulumi:"inputs"`
	IsArchived           *bool                          `pulumi:"isArchived"`
	JobType              string                         `pulumi:"jobType"`
	Limits               *CommandJobLimitsResponse      `pulumi:"limits"`
	Outputs              map[string]interface{}         `pulumi:"outputs"`
	Parameters           interface{}                    `pulumi:"parameters"`
	Properties           map[string]string              `pulumi:"properties"`
	Resources            *ResourceConfigurationResponse `pulumi:"resources"`
	Services             map[string]JobServiceResponse  `pulumi:"services"`
	Status               string                         `pulumi:"status"`
	Tags                 map[string]string              `pulumi:"tags"`
}


func (val *CommandJobResponse) Defaults() *CommandJobResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExperimentName) {
		experimentName_ := "Default"
		tmp.ExperimentName = &experimentName_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	tmp.Resources = tmp.Resources.Defaults()

	return &tmp
}

type ComponentContainerType struct {
	Description *string           `pulumi:"description"`
	IsArchived  *bool             `pulumi:"isArchived"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}


func (val *ComponentContainerType) Defaults() *ComponentContainerType {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}





type ComponentContainerTypeInput interface {
	pulumi.Input

	ToComponentContainerTypeOutput() ComponentContainerTypeOutput
	ToComponentContainerTypeOutputWithContext(context.Context) ComponentContainerTypeOutput
}

type ComponentContainerTypeArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	IsArchived  pulumi.BoolPtrInput   `pulumi:"isArchived"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}


func (val *ComponentContainerTypeArgs) Defaults() *ComponentContainerTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		tmp.IsArchived = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (ComponentContainerTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentContainerType)(nil)).Elem()
}

func (i ComponentContainerTypeArgs) ToComponentContainerTypeOutput() ComponentContainerTypeOutput {
	return i.ToComponentContainerTypeOutputWithContext(context.Background())
}

func (i ComponentContainerTypeArgs) ToComponentContainerTypeOutputWithContext(ctx context.Context) ComponentContainerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentContainerTypeOutput)
}

type ComponentContainerTypeOutput struct{ *pulumi.OutputState }

func (ComponentContainerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentContainerType)(nil)).Elem()
}

func (o ComponentContainerTypeOutput) ToComponentContainerTypeOutput() ComponentContainerTypeOutput {
	return o
}

func (o ComponentContainerTypeOutput) ToComponentContainerTypeOutputWithContext(ctx context.Context) ComponentContainerTypeOutput {
	return o
}

func (o ComponentContainerTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComponentContainerType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ComponentContainerTypeOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComponentContainerType) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o ComponentContainerTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ComponentContainerType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ComponentContainerTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ComponentContainerType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ComponentContainerResponse struct {
	Description   *string           `pulumi:"description"`
	IsArchived    *bool             `pulumi:"isArchived"`
	LatestVersion string            `pulumi:"latestVersion"`
	NextVersion   string            `pulumi:"nextVersion"`
	Properties    map[string]string `pulumi:"properties"`
	Tags          map[string]string `pulumi:"tags"`
}


func (val *ComponentContainerResponse) Defaults() *ComponentContainerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}

type ComponentContainerResponseOutput struct{ *pulumi.OutputState }

func (ComponentContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentContainerResponse)(nil)).Elem()
}

func (o ComponentContainerResponseOutput) ToComponentContainerResponseOutput() ComponentContainerResponseOutput {
	return o
}

func (o ComponentContainerResponseOutput) ToComponentContainerResponseOutputWithContext(ctx context.Context) ComponentContainerResponseOutput {
	return o
}

func (o ComponentContainerResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComponentContainerResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ComponentContainerResponseOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComponentContainerResponse) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o ComponentContainerResponseOutput) LatestVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ComponentContainerResponse) string { return v.LatestVersion }).(pulumi.StringOutput)
}

func (o ComponentContainerResponseOutput) NextVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ComponentContainerResponse) string { return v.NextVersion }).(pulumi.StringOutput)
}

func (o ComponentContainerResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ComponentContainerResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ComponentContainerResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ComponentContainerResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ComponentVersionType struct {
	ComponentSpec interface{}       `pulumi:"componentSpec"`
	Description   *string           `pulumi:"description"`
	IsAnonymous   *bool             `pulumi:"isAnonymous"`
	IsArchived    *bool             `pulumi:"isArchived"`
	Properties    map[string]string `pulumi:"properties"`
	Tags          map[string]string `pulumi:"tags"`
}


func (val *ComponentVersionType) Defaults() *ComponentVersionType {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}





type ComponentVersionTypeInput interface {
	pulumi.Input

	ToComponentVersionTypeOutput() ComponentVersionTypeOutput
	ToComponentVersionTypeOutputWithContext(context.Context) ComponentVersionTypeOutput
}

type ComponentVersionTypeArgs struct {
	ComponentSpec pulumi.Input          `pulumi:"componentSpec"`
	Description   pulumi.StringPtrInput `pulumi:"description"`
	IsAnonymous   pulumi.BoolPtrInput   `pulumi:"isAnonymous"`
	IsArchived    pulumi.BoolPtrInput   `pulumi:"isArchived"`
	Properties    pulumi.StringMapInput `pulumi:"properties"`
	Tags          pulumi.StringMapInput `pulumi:"tags"`
}


func (val *ComponentVersionTypeArgs) Defaults() *ComponentVersionTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		tmp.IsAnonymous = pulumi.BoolPtr(false)
	}
	if isZero(tmp.IsArchived) {
		tmp.IsArchived = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (ComponentVersionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentVersionType)(nil)).Elem()
}

func (i ComponentVersionTypeArgs) ToComponentVersionTypeOutput() ComponentVersionTypeOutput {
	return i.ToComponentVersionTypeOutputWithContext(context.Background())
}

func (i ComponentVersionTypeArgs) ToComponentVersionTypeOutputWithContext(ctx context.Context) ComponentVersionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentVersionTypeOutput)
}

type ComponentVersionTypeOutput struct{ *pulumi.OutputState }

func (ComponentVersionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentVersionType)(nil)).Elem()
}

func (o ComponentVersionTypeOutput) ToComponentVersionTypeOutput() ComponentVersionTypeOutput {
	return o
}

func (o ComponentVersionTypeOutput) ToComponentVersionTypeOutputWithContext(ctx context.Context) ComponentVersionTypeOutput {
	return o
}

func (o ComponentVersionTypeOutput) ComponentSpec() pulumi.AnyOutput {
	return o.ApplyT(func(v ComponentVersionType) interface{} { return v.ComponentSpec }).(pulumi.AnyOutput)
}

func (o ComponentVersionTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComponentVersionType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ComponentVersionTypeOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComponentVersionType) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o ComponentVersionTypeOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComponentVersionType) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o ComponentVersionTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ComponentVersionType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ComponentVersionTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ComponentVersionType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ComponentVersionResponse struct {
	ComponentSpec interface{}       `pulumi:"componentSpec"`
	Description   *string           `pulumi:"description"`
	IsAnonymous   *bool             `pulumi:"isAnonymous"`
	IsArchived    *bool             `pulumi:"isArchived"`
	Properties    map[string]string `pulumi:"properties"`
	Tags          map[string]string `pulumi:"tags"`
}


func (val *ComponentVersionResponse) Defaults() *ComponentVersionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}

type ComponentVersionResponseOutput struct{ *pulumi.OutputState }

func (ComponentVersionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentVersionResponse)(nil)).Elem()
}

func (o ComponentVersionResponseOutput) ToComponentVersionResponseOutput() ComponentVersionResponseOutput {
	return o
}

func (o ComponentVersionResponseOutput) ToComponentVersionResponseOutputWithContext(ctx context.Context) ComponentVersionResponseOutput {
	return o
}

func (o ComponentVersionResponseOutput) ComponentSpec() pulumi.AnyOutput {
	return o.ApplyT(func(v ComponentVersionResponse) interface{} { return v.ComponentSpec }).(pulumi.AnyOutput)
}

func (o ComponentVersionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComponentVersionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ComponentVersionResponseOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComponentVersionResponse) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o ComponentVersionResponseOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComponentVersionResponse) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o ComponentVersionResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ComponentVersionResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ComponentVersionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ComponentVersionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ComputeInstance struct {
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

type ComputeInstanceContainerResponse struct {
	Autosave    *string                                 `pulumi:"autosave"`
	Environment *ComputeInstanceEnvironmentInfoResponse `pulumi:"environment"`
	Gpu         *string                                 `pulumi:"gpu"`
	Name        *string                                 `pulumi:"name"`
	Network     *string                                 `pulumi:"network"`
	Services    []interface{}                           `pulumi:"services"`
}

type ComputeInstanceCreatedByResponse struct {
	UserId    string `pulumi:"userId"`
	UserName  string `pulumi:"userName"`
	UserOrgId string `pulumi:"userOrgId"`
}

type ComputeInstanceDataDiskResponse struct {
	Caching            *string `pulumi:"caching"`
	DiskSizeGB         *int    `pulumi:"diskSizeGB"`
	Lun                *int    `pulumi:"lun"`
	StorageAccountType *string `pulumi:"storageAccountType"`
}


func (val *ComputeInstanceDataDiskResponse) Defaults() *ComputeInstanceDataDiskResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.StorageAccountType) {
		storageAccountType_ := "Standard_LRS"
		tmp.StorageAccountType = &storageAccountType_
	}
	return &tmp
}

type ComputeInstanceDataMountResponse struct {
	CreatedBy   *string `pulumi:"createdBy"`
	Error       *string `pulumi:"error"`
	MountAction *string `pulumi:"mountAction"`
	MountName   *string `pulumi:"mountName"`
	MountPath   *string `pulumi:"mountPath"`
	MountState  *string `pulumi:"mountState"`
	MountedOn   *string `pulumi:"mountedOn"`
	Source      *string `pulumi:"source"`
	SourceType  *string `pulumi:"sourceType"`
}

type ComputeInstanceEnvironmentInfoResponse struct {
	Name    *string `pulumi:"name"`
	Version *string `pulumi:"version"`
}

type ComputeInstanceLastOperationResponse struct {
	OperationName    *string `pulumi:"operationName"`
	OperationStatus  *string `pulumi:"operationStatus"`
	OperationTime    *string `pulumi:"operationTime"`
	OperationTrigger *string `pulumi:"operationTrigger"`
}

type ComputeInstanceProperties struct {
	ApplicationSharingPolicy         *string                          `pulumi:"applicationSharingPolicy"`
	ComputeInstanceAuthorizationType *string                          `pulumi:"computeInstanceAuthorizationType"`
	EnableNodePublicIp               *bool                            `pulumi:"enableNodePublicIp"`
	PersonalComputeInstanceSettings  *PersonalComputeInstanceSettings `pulumi:"personalComputeInstanceSettings"`
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
	tmp.SshSettings = tmp.SshSettings.Defaults()

	return &tmp
}

type ComputeInstancePropertiesResponse struct {
	ApplicationSharingPolicy         *string                                      `pulumi:"applicationSharingPolicy"`
	Applications                     []ComputeInstanceApplicationResponse         `pulumi:"applications"`
	ComputeInstanceAuthorizationType *string                                      `pulumi:"computeInstanceAuthorizationType"`
	ConnectivityEndpoints            ComputeInstanceConnectivityEndpointsResponse `pulumi:"connectivityEndpoints"`
	Containers                       []ComputeInstanceContainerResponse           `pulumi:"containers"`
	CreatedBy                        ComputeInstanceCreatedByResponse             `pulumi:"createdBy"`
	DataDisks                        []ComputeInstanceDataDiskResponse            `pulumi:"dataDisks"`
	DataMounts                       []ComputeInstanceDataMountResponse           `pulumi:"dataMounts"`
	EnableNodePublicIp               *bool                                        `pulumi:"enableNodePublicIp"`
	Errors                           []ErrorResponseResponse                      `pulumi:"errors"`
	LastOperation                    ComputeInstanceLastOperationResponse         `pulumi:"lastOperation"`
	PersonalComputeInstanceSettings  *PersonalComputeInstanceSettingsResponse     `pulumi:"personalComputeInstanceSettings"`
	Schedules                        ComputeSchedulesResponse                     `pulumi:"schedules"`
	SetupScripts                     *SetupScriptsResponse                        `pulumi:"setupScripts"`
	SshSettings                      *ComputeInstanceSshSettingsResponse          `pulumi:"sshSettings"`
	State                            string                                       `pulumi:"state"`
	Subnet                           *ResourceIdResponse                          `pulumi:"subnet"`
	Versions                         ComputeInstanceVersionResponse               `pulumi:"versions"`
	VmSize                           *string                                      `pulumi:"vmSize"`
}


func (val *ComputeInstancePropertiesResponse) Defaults() *ComputeInstancePropertiesResponse {
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
	tmp.SshSettings = tmp.SshSettings.Defaults()

	return &tmp
}

type ComputeInstanceResponse struct {
	ComputeLocation    string                             `pulumi:"computeLocation"`
	ComputeType        string                             `pulumi:"computeType"`
	CreatedOn          string                             `pulumi:"createdOn"`
	Description        *string                            `pulumi:"description"`
	DisableLocalAuth   *bool                              `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                               `pulumi:"isAttachedCompute"`
	ModifiedOn         string                             `pulumi:"modifiedOn"`
	Properties         *ComputeInstancePropertiesResponse `pulumi:"properties"`
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

type ComputeInstanceVersionResponse struct {
	Runtime *string `pulumi:"runtime"`
}

type ComputeSchedulesResponse struct {
	ComputeStartStop []ComputeStartStopScheduleResponse `pulumi:"computeStartStop"`
}

type ComputeStartStopScheduleResponse struct {
	Action             *string               `pulumi:"action"`
	Id                 string                `pulumi:"id"`
	ProvisioningStatus string                `pulumi:"provisioningStatus"`
	Schedule           *ScheduleBaseResponse `pulumi:"schedule"`
}

type ContainerResourceRequirements struct {
	ContainerResourceLimits   *ContainerResourceSettings `pulumi:"containerResourceLimits"`
	ContainerResourceRequests *ContainerResourceSettings `pulumi:"containerResourceRequests"`
}

type ContainerResourceRequirementsResponse struct {
	ContainerResourceLimits   *ContainerResourceSettingsResponse `pulumi:"containerResourceLimits"`
	ContainerResourceRequests *ContainerResourceSettingsResponse `pulumi:"containerResourceRequests"`
}

type ContainerResourceSettings struct {
	Cpu    *string `pulumi:"cpu"`
	Gpu    *string `pulumi:"gpu"`
	Memory *string `pulumi:"memory"`
}

type ContainerResourceSettingsResponse struct {
	Cpu    *string `pulumi:"cpu"`
	Gpu    *string `pulumi:"gpu"`
	Memory *string `pulumi:"memory"`
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

type CustomModelJobInput struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *CustomModelJobInput) Defaults() *CustomModelJobInput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type CustomModelJobInputResponse struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *CustomModelJobInputResponse) Defaults() *CustomModelJobInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type CustomModelJobOutput struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *CustomModelJobOutput) Defaults() *CustomModelJobOutput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type CustomModelJobOutputResponse struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *CustomModelJobOutputResponse) Defaults() *CustomModelJobOutputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type DataContainerType struct {
	DataType    string            `pulumi:"dataType"`
	Description *string           `pulumi:"description"`
	IsArchived  *bool             `pulumi:"isArchived"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}


func (val *DataContainerType) Defaults() *DataContainerType {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}





type DataContainerTypeInput interface {
	pulumi.Input

	ToDataContainerTypeOutput() DataContainerTypeOutput
	ToDataContainerTypeOutputWithContext(context.Context) DataContainerTypeOutput
}

type DataContainerTypeArgs struct {
	DataType    pulumi.StringInput    `pulumi:"dataType"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	IsArchived  pulumi.BoolPtrInput   `pulumi:"isArchived"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}


func (val *DataContainerTypeArgs) Defaults() *DataContainerTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		tmp.IsArchived = pulumi.BoolPtr(false)
	}
	return &tmp
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

func (o DataContainerTypeOutput) DataType() pulumi.StringOutput {
	return o.ApplyT(func(v DataContainerType) string { return v.DataType }).(pulumi.StringOutput)
}

func (o DataContainerTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataContainerType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataContainerTypeOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataContainerType) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o DataContainerTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataContainerType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o DataContainerTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataContainerType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type DataContainerResponse struct {
	DataType      string            `pulumi:"dataType"`
	Description   *string           `pulumi:"description"`
	IsArchived    *bool             `pulumi:"isArchived"`
	LatestVersion string            `pulumi:"latestVersion"`
	NextVersion   string            `pulumi:"nextVersion"`
	Properties    map[string]string `pulumi:"properties"`
	Tags          map[string]string `pulumi:"tags"`
}


func (val *DataContainerResponse) Defaults() *DataContainerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
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

func (o DataContainerResponseOutput) DataType() pulumi.StringOutput {
	return o.ApplyT(func(v DataContainerResponse) string { return v.DataType }).(pulumi.StringOutput)
}

func (o DataContainerResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataContainerResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataContainerResponseOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataContainerResponse) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o DataContainerResponseOutput) LatestVersion() pulumi.StringOutput {
	return o.ApplyT(func(v DataContainerResponse) string { return v.LatestVersion }).(pulumi.StringOutput)
}

func (o DataContainerResponseOutput) NextVersion() pulumi.StringOutput {
	return o.ApplyT(func(v DataContainerResponse) string { return v.NextVersion }).(pulumi.StringOutput)
}

func (o DataContainerResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataContainerResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o DataContainerResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DataContainerResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type DataFactory struct {
	ComputeType      string  `pulumi:"computeType"`
	Description      *string `pulumi:"description"`
	DisableLocalAuth *bool   `pulumi:"disableLocalAuth"`
	ResourceId       *string `pulumi:"resourceId"`
}

type DataFactoryResponse struct {
	ComputeLocation    string                  `pulumi:"computeLocation"`
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
	ComputeType      string                             `pulumi:"computeType"`
	Description      *string                            `pulumi:"description"`
	DisableLocalAuth *bool                              `pulumi:"disableLocalAuth"`
	Properties       *DataLakeAnalyticsSchemaProperties `pulumi:"properties"`
	ResourceId       *string                            `pulumi:"resourceId"`
}

type DataLakeAnalyticsResponse struct {
	ComputeLocation    string                                     `pulumi:"computeLocation"`
	ComputeType        string                                     `pulumi:"computeType"`
	CreatedOn          string                                     `pulumi:"createdOn"`
	Description        *string                                    `pulumi:"description"`
	DisableLocalAuth   *bool                                      `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                                       `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                     `pulumi:"modifiedOn"`
	Properties         *DataLakeAnalyticsSchemaResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse                    `pulumi:"provisioningErrors"`
	ProvisioningState  string                                     `pulumi:"provisioningState"`
	ResourceId         *string                                    `pulumi:"resourceId"`
}

type DataLakeAnalyticsSchemaProperties struct {
	DataLakeStoreAccountName *string `pulumi:"dataLakeStoreAccountName"`
}

type DataLakeAnalyticsSchemaResponseProperties struct {
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

type Databricks struct {
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

type DatabricksPropertiesResponse struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
	WorkspaceUrl          *string `pulumi:"workspaceUrl"`
}

type DatabricksResponse struct {
	ComputeLocation    string                        `pulumi:"computeLocation"`
	ComputeType        string                        `pulumi:"computeType"`
	CreatedOn          string                        `pulumi:"createdOn"`
	Description        *string                       `pulumi:"description"`
	DisableLocalAuth   *bool                         `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                          `pulumi:"isAttachedCompute"`
	ModifiedOn         string                        `pulumi:"modifiedOn"`
	Properties         *DatabricksPropertiesResponse `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse       `pulumi:"provisioningErrors"`
	ProvisioningState  string                        `pulumi:"provisioningState"`
	ResourceId         *string                       `pulumi:"resourceId"`
}

type DefaultScaleSettings struct {
	ScaleType string `pulumi:"scaleType"`
}

type DefaultScaleSettingsResponse struct {
	ScaleType string `pulumi:"scaleType"`
}

type EncryptionKeyVaultProperties struct {
	IdentityClientId *string `pulumi:"identityClientId"`
	KeyIdentifier    string  `pulumi:"keyIdentifier"`
	KeyVaultArmId    string  `pulumi:"keyVaultArmId"`
}





type EncryptionKeyVaultPropertiesInput interface {
	pulumi.Input

	ToEncryptionKeyVaultPropertiesOutput() EncryptionKeyVaultPropertiesOutput
	ToEncryptionKeyVaultPropertiesOutputWithContext(context.Context) EncryptionKeyVaultPropertiesOutput
}

type EncryptionKeyVaultPropertiesArgs struct {
	IdentityClientId pulumi.StringPtrInput `pulumi:"identityClientId"`
	KeyIdentifier    pulumi.StringInput    `pulumi:"keyIdentifier"`
	KeyVaultArmId    pulumi.StringInput    `pulumi:"keyVaultArmId"`
}

func (EncryptionKeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyVaultProperties)(nil)).Elem()
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesOutput() EncryptionKeyVaultPropertiesOutput {
	return i.ToEncryptionKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesOutput)
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return i.ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesOutput).ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type EncryptionKeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput
	ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Context) EncryptionKeyVaultPropertiesPtrOutput
}

type encryptionKeyVaultPropertiesPtrType EncryptionKeyVaultPropertiesArgs

func EncryptionKeyVaultPropertiesPtr(v *EncryptionKeyVaultPropertiesArgs) EncryptionKeyVaultPropertiesPtrInput {
	return (*encryptionKeyVaultPropertiesPtrType)(v)
}

func (*encryptionKeyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyVaultProperties)(nil)).Elem()
}

func (i *encryptionKeyVaultPropertiesPtrType) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return i.ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *encryptionKeyVaultPropertiesPtrType) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesPtrOutput)
}

type EncryptionKeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyVaultProperties)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesOutput() EncryptionKeyVaultPropertiesOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return o.ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionKeyVaultProperties) *EncryptionKeyVaultProperties {
		return &v
	}).(EncryptionKeyVaultPropertiesPtrOutput)
}

func (o EncryptionKeyVaultPropertiesOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultProperties) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionKeyVaultProperties) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

func (o EncryptionKeyVaultPropertiesOutput) KeyVaultArmId() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionKeyVaultProperties) string { return v.KeyVaultArmId }).(pulumi.StringOutput)
}

type EncryptionKeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyVaultProperties)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesPtrOutput) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesPtrOutput) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesPtrOutput) Elem() EncryptionKeyVaultPropertiesOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) EncryptionKeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionKeyVaultProperties
		return ret
	}).(EncryptionKeyVaultPropertiesOutput)
}

func (o EncryptionKeyVaultPropertiesPtrOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.IdentityClientId
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesPtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesPtrOutput) KeyVaultArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultArmId
	}).(pulumi.StringPtrOutput)
}

type EncryptionKeyVaultPropertiesResponse struct {
	IdentityClientId *string `pulumi:"identityClientId"`
	KeyIdentifier    string  `pulumi:"keyIdentifier"`
	KeyVaultArmId    string  `pulumi:"keyVaultArmId"`
}

type EncryptionKeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesResponseOutput) ToEncryptionKeyVaultPropertiesResponseOutput() EncryptionKeyVaultPropertiesResponseOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponseOutput) ToEncryptionKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesResponseOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponseOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) KeyVaultArmId() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) string { return v.KeyVaultArmId }).(pulumi.StringOutput)
}

type EncryptionKeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) ToEncryptionKeyVaultPropertiesResponsePtrOutput() EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) ToEncryptionKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) Elem() EncryptionKeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) EncryptionKeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionKeyVaultPropertiesResponse
		return ret
	}).(EncryptionKeyVaultPropertiesResponseOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IdentityClientId
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) KeyVaultArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultArmId
	}).(pulumi.StringPtrOutput)
}

type EncryptionProperty struct {
	Identity           *IdentityForCmk              `pulumi:"identity"`
	KeyVaultProperties EncryptionKeyVaultProperties `pulumi:"keyVaultProperties"`
	Status             string                       `pulumi:"status"`
}





type EncryptionPropertyInput interface {
	pulumi.Input

	ToEncryptionPropertyOutput() EncryptionPropertyOutput
	ToEncryptionPropertyOutputWithContext(context.Context) EncryptionPropertyOutput
}

type EncryptionPropertyArgs struct {
	Identity           IdentityForCmkPtrInput            `pulumi:"identity"`
	KeyVaultProperties EncryptionKeyVaultPropertiesInput `pulumi:"keyVaultProperties"`
	Status             pulumi.StringInput                `pulumi:"status"`
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

func (o EncryptionPropertyOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesOutput {
	return o.ApplyT(func(v EncryptionProperty) EncryptionKeyVaultProperties { return v.KeyVaultProperties }).(EncryptionKeyVaultPropertiesOutput)
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

func (o EncryptionPropertyPtrOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *EncryptionProperty) *EncryptionKeyVaultProperties {
		if v == nil {
			return nil
		}
		return &v.KeyVaultProperties
	}).(EncryptionKeyVaultPropertiesPtrOutput)
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
	Identity           *IdentityForCmkResponse              `pulumi:"identity"`
	KeyVaultProperties EncryptionKeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Status             string                               `pulumi:"status"`
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

func (o EncryptionPropertyResponseOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v EncryptionPropertyResponse) EncryptionKeyVaultPropertiesResponse { return v.KeyVaultProperties }).(EncryptionKeyVaultPropertiesResponseOutput)
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

func (o EncryptionPropertyResponsePtrOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) *EncryptionKeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.KeyVaultProperties
	}).(EncryptionKeyVaultPropertiesResponsePtrOutput)
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
	IsArchived  *bool             `pulumi:"isArchived"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}


func (val *EnvironmentContainerType) Defaults() *EnvironmentContainerType {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}





type EnvironmentContainerTypeInput interface {
	pulumi.Input

	ToEnvironmentContainerTypeOutput() EnvironmentContainerTypeOutput
	ToEnvironmentContainerTypeOutputWithContext(context.Context) EnvironmentContainerTypeOutput
}

type EnvironmentContainerTypeArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	IsArchived  pulumi.BoolPtrInput   `pulumi:"isArchived"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}


func (val *EnvironmentContainerTypeArgs) Defaults() *EnvironmentContainerTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		tmp.IsArchived = pulumi.BoolPtr(false)
	}
	return &tmp
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

func (o EnvironmentContainerTypeOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnvironmentContainerType) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o EnvironmentContainerTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentContainerType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o EnvironmentContainerTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentContainerType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type EnvironmentContainerResponse struct {
	Description   *string           `pulumi:"description"`
	IsArchived    *bool             `pulumi:"isArchived"`
	LatestVersion string            `pulumi:"latestVersion"`
	NextVersion   string            `pulumi:"nextVersion"`
	Properties    map[string]string `pulumi:"properties"`
	Tags          map[string]string `pulumi:"tags"`
}


func (val *EnvironmentContainerResponse) Defaults() *EnvironmentContainerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
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

func (o EnvironmentContainerResponseOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnvironmentContainerResponse) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o EnvironmentContainerResponseOutput) LatestVersion() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentContainerResponse) string { return v.LatestVersion }).(pulumi.StringOutput)
}

func (o EnvironmentContainerResponseOutput) NextVersion() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentContainerResponse) string { return v.NextVersion }).(pulumi.StringOutput)
}

func (o EnvironmentContainerResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentContainerResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o EnvironmentContainerResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentContainerResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type EnvironmentVersionType struct {
	Build           *BuildContext                 `pulumi:"build"`
	CondaFile       *string                       `pulumi:"condaFile"`
	Description     *string                       `pulumi:"description"`
	Image           *string                       `pulumi:"image"`
	InferenceConfig *InferenceContainerProperties `pulumi:"inferenceConfig"`
	IsAnonymous     *bool                         `pulumi:"isAnonymous"`
	IsArchived      *bool                         `pulumi:"isArchived"`
	OsType          *string                       `pulumi:"osType"`
	Properties      map[string]string             `pulumi:"properties"`
	Tags            map[string]string             `pulumi:"tags"`
}


func (val *EnvironmentVersionType) Defaults() *EnvironmentVersionType {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Build = tmp.Build.Defaults()

	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	if isZero(tmp.OsType) {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	return &tmp
}





type EnvironmentVersionTypeInput interface {
	pulumi.Input

	ToEnvironmentVersionTypeOutput() EnvironmentVersionTypeOutput
	ToEnvironmentVersionTypeOutputWithContext(context.Context) EnvironmentVersionTypeOutput
}

type EnvironmentVersionTypeArgs struct {
	Build           BuildContextPtrInput                 `pulumi:"build"`
	CondaFile       pulumi.StringPtrInput                `pulumi:"condaFile"`
	Description     pulumi.StringPtrInput                `pulumi:"description"`
	Image           pulumi.StringPtrInput                `pulumi:"image"`
	InferenceConfig InferenceContainerPropertiesPtrInput `pulumi:"inferenceConfig"`
	IsAnonymous     pulumi.BoolPtrInput                  `pulumi:"isAnonymous"`
	IsArchived      pulumi.BoolPtrInput                  `pulumi:"isArchived"`
	OsType          pulumi.StringPtrInput                `pulumi:"osType"`
	Properties      pulumi.StringMapInput                `pulumi:"properties"`
	Tags            pulumi.StringMapInput                `pulumi:"tags"`
}


func (val *EnvironmentVersionTypeArgs) Defaults() *EnvironmentVersionTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	if isZero(tmp.IsAnonymous) {
		tmp.IsAnonymous = pulumi.BoolPtr(false)
	}
	if isZero(tmp.IsArchived) {
		tmp.IsArchived = pulumi.BoolPtr(false)
	}
	if isZero(tmp.OsType) {
		tmp.OsType = pulumi.StringPtr("Linux")
	}
	return &tmp
}
func (EnvironmentVersionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVersionType)(nil)).Elem()
}

func (i EnvironmentVersionTypeArgs) ToEnvironmentVersionTypeOutput() EnvironmentVersionTypeOutput {
	return i.ToEnvironmentVersionTypeOutputWithContext(context.Background())
}

func (i EnvironmentVersionTypeArgs) ToEnvironmentVersionTypeOutputWithContext(ctx context.Context) EnvironmentVersionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVersionTypeOutput)
}

type EnvironmentVersionTypeOutput struct{ *pulumi.OutputState }

func (EnvironmentVersionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVersionType)(nil)).Elem()
}

func (o EnvironmentVersionTypeOutput) ToEnvironmentVersionTypeOutput() EnvironmentVersionTypeOutput {
	return o
}

func (o EnvironmentVersionTypeOutput) ToEnvironmentVersionTypeOutputWithContext(ctx context.Context) EnvironmentVersionTypeOutput {
	return o
}

func (o EnvironmentVersionTypeOutput) Build() BuildContextPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionType) *BuildContext { return v.Build }).(BuildContextPtrOutput)
}

func (o EnvironmentVersionTypeOutput) CondaFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionType) *string { return v.CondaFile }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVersionTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVersionTypeOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionType) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVersionTypeOutput) InferenceConfig() InferenceContainerPropertiesPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionType) *InferenceContainerProperties { return v.InferenceConfig }).(InferenceContainerPropertiesPtrOutput)
}

func (o EnvironmentVersionTypeOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionType) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o EnvironmentVersionTypeOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionType) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o EnvironmentVersionTypeOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionType) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVersionTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentVersionType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o EnvironmentVersionTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentVersionType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type EnvironmentVersionResponse struct {
	Build           *BuildContextResponse                 `pulumi:"build"`
	CondaFile       *string                               `pulumi:"condaFile"`
	Description     *string                               `pulumi:"description"`
	EnvironmentType string                                `pulumi:"environmentType"`
	Image           *string                               `pulumi:"image"`
	InferenceConfig *InferenceContainerPropertiesResponse `pulumi:"inferenceConfig"`
	IsAnonymous     *bool                                 `pulumi:"isAnonymous"`
	IsArchived      *bool                                 `pulumi:"isArchived"`
	OsType          *string                               `pulumi:"osType"`
	Properties      map[string]string                     `pulumi:"properties"`
	Tags            map[string]string                     `pulumi:"tags"`
}


func (val *EnvironmentVersionResponse) Defaults() *EnvironmentVersionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Build = tmp.Build.Defaults()

	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	if isZero(tmp.OsType) {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	return &tmp
}

type EnvironmentVersionResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentVersionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVersionResponse)(nil)).Elem()
}

func (o EnvironmentVersionResponseOutput) ToEnvironmentVersionResponseOutput() EnvironmentVersionResponseOutput {
	return o
}

func (o EnvironmentVersionResponseOutput) ToEnvironmentVersionResponseOutputWithContext(ctx context.Context) EnvironmentVersionResponseOutput {
	return o
}

func (o EnvironmentVersionResponseOutput) Build() BuildContextResponsePtrOutput {
	return o.ApplyT(func(v EnvironmentVersionResponse) *BuildContextResponse { return v.Build }).(BuildContextResponsePtrOutput)
}

func (o EnvironmentVersionResponseOutput) CondaFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionResponse) *string { return v.CondaFile }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVersionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVersionResponseOutput) EnvironmentType() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentVersionResponse) string { return v.EnvironmentType }).(pulumi.StringOutput)
}

func (o EnvironmentVersionResponseOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionResponse) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVersionResponseOutput) InferenceConfig() InferenceContainerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EnvironmentVersionResponse) *InferenceContainerPropertiesResponse { return v.InferenceConfig }).(InferenceContainerPropertiesResponsePtrOutput)
}

func (o EnvironmentVersionResponseOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionResponse) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o EnvironmentVersionResponseOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionResponse) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o EnvironmentVersionResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVersionResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVersionResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentVersionResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o EnvironmentVersionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentVersionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
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

type GridSamplingAlgorithm struct {
	SamplingAlgorithmType string `pulumi:"samplingAlgorithmType"`
}

type GridSamplingAlgorithmResponse struct {
	SamplingAlgorithmType string `pulumi:"samplingAlgorithmType"`
}

type HDInsight struct {
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

type HDInsightPropertiesResponse struct {
	Address              *string                               `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	SshPort              *int                                  `pulumi:"sshPort"`
}

type HDInsightResponse struct {
	ComputeLocation    string                       `pulumi:"computeLocation"`
	ComputeType        string                       `pulumi:"computeType"`
	CreatedOn          string                       `pulumi:"createdOn"`
	Description        *string                      `pulumi:"description"`
	DisableLocalAuth   *bool                        `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                         `pulumi:"isAttachedCompute"`
	ModifiedOn         string                       `pulumi:"modifiedOn"`
	Properties         *HDInsightPropertiesResponse `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse      `pulumi:"provisioningErrors"`
	ProvisioningState  string                       `pulumi:"provisioningState"`
	ResourceId         *string                      `pulumi:"resourceId"`
}

type IdAssetReference struct {
	AssetId       string `pulumi:"assetId"`
	ReferenceType string `pulumi:"referenceType"`
}

type IdAssetReferenceResponse struct {
	AssetId       string `pulumi:"assetId"`
	ReferenceType string `pulumi:"referenceType"`
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

type InstanceTypeSchema struct {
	NodeSelector map[string]string            `pulumi:"nodeSelector"`
	Resources    *InstanceTypeSchemaResources `pulumi:"resources"`
}

type InstanceTypeSchemaResources struct {
	Limits   map[string]string `pulumi:"limits"`
	Requests map[string]string `pulumi:"requests"`
}

type InstanceTypeSchemaResponse struct {
	NodeSelector map[string]string                    `pulumi:"nodeSelector"`
	Resources    *InstanceTypeSchemaResponseResources `pulumi:"resources"`
}

type InstanceTypeSchemaResponseResources struct {
	Limits   map[string]string `pulumi:"limits"`
	Requests map[string]string `pulumi:"requests"`
}

type JobService struct {
	Endpoint       *string           `pulumi:"endpoint"`
	JobServiceType *string           `pulumi:"jobServiceType"`
	Port           *int              `pulumi:"port"`
	Properties     map[string]string `pulumi:"properties"`
}

type JobServiceResponse struct {
	Endpoint       *string           `pulumi:"endpoint"`
	ErrorMessage   string            `pulumi:"errorMessage"`
	JobServiceType *string           `pulumi:"jobServiceType"`
	Port           *int              `pulumi:"port"`
	Properties     map[string]string `pulumi:"properties"`
	Status         string            `pulumi:"status"`
}

type Kubernetes struct {
	ComputeType      string                `pulumi:"computeType"`
	Description      *string               `pulumi:"description"`
	DisableLocalAuth *bool                 `pulumi:"disableLocalAuth"`
	Properties       *KubernetesProperties `pulumi:"properties"`
	ResourceId       *string               `pulumi:"resourceId"`
}


func (val *Kubernetes) Defaults() *Kubernetes {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type KubernetesOnlineDeployment struct {
	AppInsightsEnabled            *bool                          `pulumi:"appInsightsEnabled"`
	CodeConfiguration             *CodeConfiguration             `pulumi:"codeConfiguration"`
	ContainerResourceRequirements *ContainerResourceRequirements `pulumi:"containerResourceRequirements"`
	Description                   *string                        `pulumi:"description"`
	EndpointComputeType           string                         `pulumi:"endpointComputeType"`
	EnvironmentId                 *string                        `pulumi:"environmentId"`
	EnvironmentVariables          map[string]string              `pulumi:"environmentVariables"`
	InstanceType                  *string                        `pulumi:"instanceType"`
	LivenessProbe                 *ProbeSettings                 `pulumi:"livenessProbe"`
	Model                         *string                        `pulumi:"model"`
	ModelMountPath                *string                        `pulumi:"modelMountPath"`
	Properties                    map[string]string              `pulumi:"properties"`
	ReadinessProbe                *ProbeSettings                 `pulumi:"readinessProbe"`
	RequestSettings               *OnlineRequestSettings         `pulumi:"requestSettings"`
	ScaleSettings                 interface{}                    `pulumi:"scaleSettings"`
}


func (val *KubernetesOnlineDeployment) Defaults() *KubernetesOnlineDeployment {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AppInsightsEnabled) {
		appInsightsEnabled_ := false
		tmp.AppInsightsEnabled = &appInsightsEnabled_
	}
	tmp.LivenessProbe = tmp.LivenessProbe.Defaults()

	tmp.ReadinessProbe = tmp.ReadinessProbe.Defaults()

	tmp.RequestSettings = tmp.RequestSettings.Defaults()

	return &tmp
}

type KubernetesOnlineDeploymentResponse struct {
	AppInsightsEnabled            *bool                                  `pulumi:"appInsightsEnabled"`
	CodeConfiguration             *CodeConfigurationResponse             `pulumi:"codeConfiguration"`
	ContainerResourceRequirements *ContainerResourceRequirementsResponse `pulumi:"containerResourceRequirements"`
	Description                   *string                                `pulumi:"description"`
	EndpointComputeType           string                                 `pulumi:"endpointComputeType"`
	EnvironmentId                 *string                                `pulumi:"environmentId"`
	EnvironmentVariables          map[string]string                      `pulumi:"environmentVariables"`
	InstanceType                  *string                                `pulumi:"instanceType"`
	LivenessProbe                 *ProbeSettingsResponse                 `pulumi:"livenessProbe"`
	Model                         *string                                `pulumi:"model"`
	ModelMountPath                *string                                `pulumi:"modelMountPath"`
	Properties                    map[string]string                      `pulumi:"properties"`
	ProvisioningState             string                                 `pulumi:"provisioningState"`
	ReadinessProbe                *ProbeSettingsResponse                 `pulumi:"readinessProbe"`
	RequestSettings               *OnlineRequestSettingsResponse         `pulumi:"requestSettings"`
	ScaleSettings                 interface{}                            `pulumi:"scaleSettings"`
}


func (val *KubernetesOnlineDeploymentResponse) Defaults() *KubernetesOnlineDeploymentResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AppInsightsEnabled) {
		appInsightsEnabled_ := false
		tmp.AppInsightsEnabled = &appInsightsEnabled_
	}
	tmp.LivenessProbe = tmp.LivenessProbe.Defaults()

	tmp.ReadinessProbe = tmp.ReadinessProbe.Defaults()

	tmp.RequestSettings = tmp.RequestSettings.Defaults()

	return &tmp
}

type KubernetesProperties struct {
	DefaultInstanceType           *string                       `pulumi:"defaultInstanceType"`
	ExtensionInstanceReleaseTrain *string                       `pulumi:"extensionInstanceReleaseTrain"`
	ExtensionPrincipalId          *string                       `pulumi:"extensionPrincipalId"`
	InstanceTypes                 map[string]InstanceTypeSchema `pulumi:"instanceTypes"`
	Namespace                     *string                       `pulumi:"namespace"`
	RelayConnectionString         *string                       `pulumi:"relayConnectionString"`
	ServiceBusConnectionString    *string                       `pulumi:"serviceBusConnectionString"`
	VcName                        *string                       `pulumi:"vcName"`
}


func (val *KubernetesProperties) Defaults() *KubernetesProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Namespace) {
		namespace_ := "default"
		tmp.Namespace = &namespace_
	}
	return &tmp
}

type KubernetesPropertiesResponse struct {
	DefaultInstanceType           *string                               `pulumi:"defaultInstanceType"`
	ExtensionInstanceReleaseTrain *string                               `pulumi:"extensionInstanceReleaseTrain"`
	ExtensionPrincipalId          *string                               `pulumi:"extensionPrincipalId"`
	InstanceTypes                 map[string]InstanceTypeSchemaResponse `pulumi:"instanceTypes"`
	Namespace                     *string                               `pulumi:"namespace"`
	RelayConnectionString         *string                               `pulumi:"relayConnectionString"`
	ServiceBusConnectionString    *string                               `pulumi:"serviceBusConnectionString"`
	VcName                        *string                               `pulumi:"vcName"`
}


func (val *KubernetesPropertiesResponse) Defaults() *KubernetesPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Namespace) {
		namespace_ := "default"
		tmp.Namespace = &namespace_
	}
	return &tmp
}

type KubernetesResponse struct {
	ComputeLocation    string                        `pulumi:"computeLocation"`
	ComputeType        string                        `pulumi:"computeType"`
	CreatedOn          string                        `pulumi:"createdOn"`
	Description        *string                       `pulumi:"description"`
	DisableLocalAuth   *bool                         `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                          `pulumi:"isAttachedCompute"`
	ModifiedOn         string                        `pulumi:"modifiedOn"`
	Properties         *KubernetesPropertiesResponse `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse       `pulumi:"provisioningErrors"`
	ProvisioningState  string                        `pulumi:"provisioningState"`
	ResourceId         *string                       `pulumi:"resourceId"`
}


func (val *KubernetesResponse) Defaults() *KubernetesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type ListNotebookKeysResultResponse struct {
	PrimaryAccessKey   string `pulumi:"primaryAccessKey"`
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
}

type ListNotebookKeysResultResponseOutput struct{ *pulumi.OutputState }

func (ListNotebookKeysResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNotebookKeysResultResponse)(nil)).Elem()
}

func (o ListNotebookKeysResultResponseOutput) ToListNotebookKeysResultResponseOutput() ListNotebookKeysResultResponseOutput {
	return o
}

func (o ListNotebookKeysResultResponseOutput) ToListNotebookKeysResultResponseOutputWithContext(ctx context.Context) ListNotebookKeysResultResponseOutput {
	return o
}

func (o ListNotebookKeysResultResponseOutput) PrimaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListNotebookKeysResultResponse) string { return v.PrimaryAccessKey }).(pulumi.StringOutput)
}

func (o ListNotebookKeysResultResponseOutput) SecondaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListNotebookKeysResultResponse) string { return v.SecondaryAccessKey }).(pulumi.StringOutput)
}

type LiteralJobInput struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Value        string  `pulumi:"value"`
}

type LiteralJobInputResponse struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Value        string  `pulumi:"value"`
}

type MLFlowModelJobInput struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *MLFlowModelJobInput) Defaults() *MLFlowModelJobInput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type MLFlowModelJobInputResponse struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *MLFlowModelJobInputResponse) Defaults() *MLFlowModelJobInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type MLFlowModelJobOutput struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *MLFlowModelJobOutput) Defaults() *MLFlowModelJobOutput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type MLFlowModelJobOutputResponse struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *MLFlowModelJobOutputResponse) Defaults() *MLFlowModelJobOutputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type MLTableData struct {
	DataType       string            `pulumi:"dataType"`
	DataUri        string            `pulumi:"dataUri"`
	Description    *string           `pulumi:"description"`
	IsAnonymous    *bool             `pulumi:"isAnonymous"`
	IsArchived     *bool             `pulumi:"isArchived"`
	Properties     map[string]string `pulumi:"properties"`
	ReferencedUris []string          `pulumi:"referencedUris"`
	Tags           map[string]string `pulumi:"tags"`
}


func (val *MLTableData) Defaults() *MLTableData {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}

type MLTableDataResponse struct {
	DataType       string            `pulumi:"dataType"`
	DataUri        string            `pulumi:"dataUri"`
	Description    *string           `pulumi:"description"`
	IsAnonymous    *bool             `pulumi:"isAnonymous"`
	IsArchived     *bool             `pulumi:"isArchived"`
	Properties     map[string]string `pulumi:"properties"`
	ReferencedUris []string          `pulumi:"referencedUris"`
	Tags           map[string]string `pulumi:"tags"`
}


func (val *MLTableDataResponse) Defaults() *MLTableDataResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}

type MLTableJobInput struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *MLTableJobInput) Defaults() *MLTableJobInput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type MLTableJobInputResponse struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *MLTableJobInputResponse) Defaults() *MLTableJobInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type MLTableJobOutput struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *MLTableJobOutput) Defaults() *MLTableJobOutput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type MLTableJobOutputResponse struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *MLTableJobOutputResponse) Defaults() *MLTableJobOutputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type ManagedIdentity struct {
	ClientId     *string `pulumi:"clientId"`
	IdentityType string  `pulumi:"identityType"`
	ObjectId     *string `pulumi:"objectId"`
	ResourceId   *string `pulumi:"resourceId"`
}

type ManagedIdentityAuthTypeWorkspaceConnectionProperties struct {
	AuthType    string                              `pulumi:"authType"`
	Category    *string                             `pulumi:"category"`
	Credentials *WorkspaceConnectionManagedIdentity `pulumi:"credentials"`
	Target      *string                             `pulumi:"target"`
	Value       *string                             `pulumi:"value"`
	ValueFormat *string                             `pulumi:"valueFormat"`
}

type ManagedIdentityAuthTypeWorkspaceConnectionPropertiesResponse struct {
	AuthType    string                                      `pulumi:"authType"`
	Category    *string                                     `pulumi:"category"`
	Credentials *WorkspaceConnectionManagedIdentityResponse `pulumi:"credentials"`
	Target      *string                                     `pulumi:"target"`
	Value       *string                                     `pulumi:"value"`
	ValueFormat *string                                     `pulumi:"valueFormat"`
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
	Model                *string                `pulumi:"model"`
	ModelMountPath       *string                `pulumi:"modelMountPath"`
	Properties           map[string]string      `pulumi:"properties"`
	ReadinessProbe       *ProbeSettings         `pulumi:"readinessProbe"`
	RequestSettings      *OnlineRequestSettings `pulumi:"requestSettings"`
	ScaleSettings        interface{}            `pulumi:"scaleSettings"`
}


func (val *ManagedOnlineDeployment) Defaults() *ManagedOnlineDeployment {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AppInsightsEnabled) {
		appInsightsEnabled_ := false
		tmp.AppInsightsEnabled = &appInsightsEnabled_
	}
	tmp.LivenessProbe = tmp.LivenessProbe.Defaults()

	tmp.ReadinessProbe = tmp.ReadinessProbe.Defaults()

	tmp.RequestSettings = tmp.RequestSettings.Defaults()

	return &tmp
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
	Model                *string                        `pulumi:"model"`
	ModelMountPath       *string                        `pulumi:"modelMountPath"`
	Properties           map[string]string              `pulumi:"properties"`
	ProvisioningState    string                         `pulumi:"provisioningState"`
	ReadinessProbe       *ProbeSettingsResponse         `pulumi:"readinessProbe"`
	RequestSettings      *OnlineRequestSettingsResponse `pulumi:"requestSettings"`
	ScaleSettings        interface{}                    `pulumi:"scaleSettings"`
}


func (val *ManagedOnlineDeploymentResponse) Defaults() *ManagedOnlineDeploymentResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AppInsightsEnabled) {
		appInsightsEnabled_ := false
		tmp.AppInsightsEnabled = &appInsightsEnabled_
	}
	tmp.LivenessProbe = tmp.LivenessProbe.Defaults()

	tmp.ReadinessProbe = tmp.ReadinessProbe.Defaults()

	tmp.RequestSettings = tmp.RequestSettings.Defaults()

	return &tmp
}

type ManagedServiceIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type MedianStoppingPolicy struct {
	DelayEvaluation    *int   `pulumi:"delayEvaluation"`
	EvaluationInterval *int   `pulumi:"evaluationInterval"`
	PolicyType         string `pulumi:"policyType"`
}


func (val *MedianStoppingPolicy) Defaults() *MedianStoppingPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DelayEvaluation) {
		delayEvaluation_ := 0
		tmp.DelayEvaluation = &delayEvaluation_
	}
	if isZero(tmp.EvaluationInterval) {
		evaluationInterval_ := 0
		tmp.EvaluationInterval = &evaluationInterval_
	}
	return &tmp
}

type MedianStoppingPolicyResponse struct {
	DelayEvaluation    *int   `pulumi:"delayEvaluation"`
	EvaluationInterval *int   `pulumi:"evaluationInterval"`
	PolicyType         string `pulumi:"policyType"`
}


func (val *MedianStoppingPolicyResponse) Defaults() *MedianStoppingPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DelayEvaluation) {
		delayEvaluation_ := 0
		tmp.DelayEvaluation = &delayEvaluation_
	}
	if isZero(tmp.EvaluationInterval) {
		evaluationInterval_ := 0
		tmp.EvaluationInterval = &evaluationInterval_
	}
	return &tmp
}

type ModelContainerType struct {
	Description *string           `pulumi:"description"`
	IsArchived  *bool             `pulumi:"isArchived"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}


func (val *ModelContainerType) Defaults() *ModelContainerType {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}





type ModelContainerTypeInput interface {
	pulumi.Input

	ToModelContainerTypeOutput() ModelContainerTypeOutput
	ToModelContainerTypeOutputWithContext(context.Context) ModelContainerTypeOutput
}

type ModelContainerTypeArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	IsArchived  pulumi.BoolPtrInput   `pulumi:"isArchived"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}


func (val *ModelContainerTypeArgs) Defaults() *ModelContainerTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		tmp.IsArchived = pulumi.BoolPtr(false)
	}
	return &tmp
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

func (o ModelContainerTypeOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelContainerType) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o ModelContainerTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelContainerType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ModelContainerTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelContainerType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ModelContainerResponse struct {
	Description   *string           `pulumi:"description"`
	IsArchived    *bool             `pulumi:"isArchived"`
	LatestVersion string            `pulumi:"latestVersion"`
	NextVersion   string            `pulumi:"nextVersion"`
	Properties    map[string]string `pulumi:"properties"`
	Tags          map[string]string `pulumi:"tags"`
}


func (val *ModelContainerResponse) Defaults() *ModelContainerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
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

func (o ModelContainerResponseOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelContainerResponse) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o ModelContainerResponseOutput) LatestVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ModelContainerResponse) string { return v.LatestVersion }).(pulumi.StringOutput)
}

func (o ModelContainerResponseOutput) NextVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ModelContainerResponse) string { return v.NextVersion }).(pulumi.StringOutput)
}

func (o ModelContainerResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelContainerResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ModelContainerResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelContainerResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ModelVersionType struct {
	Description *string               `pulumi:"description"`
	Flavors     map[string]FlavorData `pulumi:"flavors"`
	IsAnonymous *bool                 `pulumi:"isAnonymous"`
	IsArchived  *bool                 `pulumi:"isArchived"`
	JobName     *string               `pulumi:"jobName"`
	ModelType   *string               `pulumi:"modelType"`
	ModelUri    *string               `pulumi:"modelUri"`
	Properties  map[string]string     `pulumi:"properties"`
	Tags        map[string]string     `pulumi:"tags"`
}


func (val *ModelVersionType) Defaults() *ModelVersionType {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}





type ModelVersionTypeInput interface {
	pulumi.Input

	ToModelVersionTypeOutput() ModelVersionTypeOutput
	ToModelVersionTypeOutputWithContext(context.Context) ModelVersionTypeOutput
}

type ModelVersionTypeArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Flavors     FlavorDataMapInput    `pulumi:"flavors"`
	IsAnonymous pulumi.BoolPtrInput   `pulumi:"isAnonymous"`
	IsArchived  pulumi.BoolPtrInput   `pulumi:"isArchived"`
	JobName     pulumi.StringPtrInput `pulumi:"jobName"`
	ModelType   pulumi.StringPtrInput `pulumi:"modelType"`
	ModelUri    pulumi.StringPtrInput `pulumi:"modelUri"`
	Properties  pulumi.StringMapInput `pulumi:"properties"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}


func (val *ModelVersionTypeArgs) Defaults() *ModelVersionTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		tmp.IsAnonymous = pulumi.BoolPtr(false)
	}
	if isZero(tmp.IsArchived) {
		tmp.IsArchived = pulumi.BoolPtr(false)
	}
	return &tmp
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

func (o ModelVersionTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ModelVersionTypeOutput) Flavors() FlavorDataMapOutput {
	return o.ApplyT(func(v ModelVersionType) map[string]FlavorData { return v.Flavors }).(FlavorDataMapOutput)
}

func (o ModelVersionTypeOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelVersionType) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o ModelVersionTypeOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelVersionType) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o ModelVersionTypeOutput) JobName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionType) *string { return v.JobName }).(pulumi.StringPtrOutput)
}

func (o ModelVersionTypeOutput) ModelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionType) *string { return v.ModelType }).(pulumi.StringPtrOutput)
}

func (o ModelVersionTypeOutput) ModelUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionType) *string { return v.ModelUri }).(pulumi.StringPtrOutput)
}

func (o ModelVersionTypeOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelVersionType) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ModelVersionTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelVersionType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ModelVersionResponse struct {
	Description *string                       `pulumi:"description"`
	Flavors     map[string]FlavorDataResponse `pulumi:"flavors"`
	IsAnonymous *bool                         `pulumi:"isAnonymous"`
	IsArchived  *bool                         `pulumi:"isArchived"`
	JobName     *string                       `pulumi:"jobName"`
	ModelType   *string                       `pulumi:"modelType"`
	ModelUri    *string                       `pulumi:"modelUri"`
	Properties  map[string]string             `pulumi:"properties"`
	Tags        map[string]string             `pulumi:"tags"`
}


func (val *ModelVersionResponse) Defaults() *ModelVersionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
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

func (o ModelVersionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ModelVersionResponseOutput) Flavors() FlavorDataResponseMapOutput {
	return o.ApplyT(func(v ModelVersionResponse) map[string]FlavorDataResponse { return v.Flavors }).(FlavorDataResponseMapOutput)
}

func (o ModelVersionResponseOutput) IsAnonymous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelVersionResponse) *bool { return v.IsAnonymous }).(pulumi.BoolPtrOutput)
}

func (o ModelVersionResponseOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelVersionResponse) *bool { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

func (o ModelVersionResponseOutput) JobName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionResponse) *string { return v.JobName }).(pulumi.StringPtrOutput)
}

func (o ModelVersionResponseOutput) ModelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionResponse) *string { return v.ModelType }).(pulumi.StringPtrOutput)
}

func (o ModelVersionResponseOutput) ModelUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelVersionResponse) *string { return v.ModelUri }).(pulumi.StringPtrOutput)
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

type NoneAuthTypeWorkspaceConnectionProperties struct {
	AuthType    string  `pulumi:"authType"`
	Category    *string `pulumi:"category"`
	Target      *string `pulumi:"target"`
	Value       *string `pulumi:"value"`
	ValueFormat *string `pulumi:"valueFormat"`
}

type NoneAuthTypeWorkspaceConnectionPropertiesResponse struct {
	AuthType    string  `pulumi:"authType"`
	Category    *string `pulumi:"category"`
	Target      *string `pulumi:"target"`
	Value       *string `pulumi:"value"`
	ValueFormat *string `pulumi:"valueFormat"`
}

type NoneDatastoreCredentials struct {
	CredentialsType string `pulumi:"credentialsType"`
}

type NoneDatastoreCredentialsResponse struct {
	CredentialsType string `pulumi:"credentialsType"`
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
	Compute     *string           `pulumi:"compute"`
	Description *string           `pulumi:"description"`
	Keys        *EndpointAuthKeys `pulumi:"keys"`
	Properties  map[string]string `pulumi:"properties"`
	Traffic     map[string]int    `pulumi:"traffic"`
}





type OnlineEndpointTypeInput interface {
	pulumi.Input

	ToOnlineEndpointTypeOutput() OnlineEndpointTypeOutput
	ToOnlineEndpointTypeOutputWithContext(context.Context) OnlineEndpointTypeOutput
}

type OnlineEndpointTypeArgs struct {
	AuthMode    pulumi.StringInput       `pulumi:"authMode"`
	Compute     pulumi.StringPtrInput    `pulumi:"compute"`
	Description pulumi.StringPtrInput    `pulumi:"description"`
	Keys        EndpointAuthKeysPtrInput `pulumi:"keys"`
	Properties  pulumi.StringMapInput    `pulumi:"properties"`
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

func (o OnlineEndpointTypeOutput) Compute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnlineEndpointType) *string { return v.Compute }).(pulumi.StringPtrOutput)
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

func (o OnlineEndpointTypeOutput) Traffic() pulumi.IntMapOutput {
	return o.ApplyT(func(v OnlineEndpointType) map[string]int { return v.Traffic }).(pulumi.IntMapOutput)
}

type OnlineEndpointResponse struct {
	AuthMode          string            `pulumi:"authMode"`
	Compute           *string           `pulumi:"compute"`
	Description       *string           `pulumi:"description"`
	Properties        map[string]string `pulumi:"properties"`
	ProvisioningState string            `pulumi:"provisioningState"`
	ScoringUri        string            `pulumi:"scoringUri"`
	SwaggerUri        string            `pulumi:"swaggerUri"`
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

func (o OnlineEndpointResponseOutput) Compute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OnlineEndpointResponse) *string { return v.Compute }).(pulumi.StringPtrOutput)
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

func (o OnlineEndpointResponseOutput) Traffic() pulumi.IntMapOutput {
	return o.ApplyT(func(v OnlineEndpointResponse) map[string]int { return v.Traffic }).(pulumi.IntMapOutput)
}

type OnlineRequestSettings struct {
	MaxConcurrentRequestsPerInstance *int    `pulumi:"maxConcurrentRequestsPerInstance"`
	MaxQueueWait                     *string `pulumi:"maxQueueWait"`
	RequestTimeout                   *string `pulumi:"requestTimeout"`
}


func (val *OnlineRequestSettings) Defaults() *OnlineRequestSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxConcurrentRequestsPerInstance) {
		maxConcurrentRequestsPerInstance_ := 1
		tmp.MaxConcurrentRequestsPerInstance = &maxConcurrentRequestsPerInstance_
	}
	if isZero(tmp.MaxQueueWait) {
		maxQueueWait_ := "PT0.5S"
		tmp.MaxQueueWait = &maxQueueWait_
	}
	if isZero(tmp.RequestTimeout) {
		requestTimeout_ := "PT5S"
		tmp.RequestTimeout = &requestTimeout_
	}
	return &tmp
}

type OnlineRequestSettingsResponse struct {
	MaxConcurrentRequestsPerInstance *int    `pulumi:"maxConcurrentRequestsPerInstance"`
	MaxQueueWait                     *string `pulumi:"maxQueueWait"`
	RequestTimeout                   *string `pulumi:"requestTimeout"`
}


func (val *OnlineRequestSettingsResponse) Defaults() *OnlineRequestSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxConcurrentRequestsPerInstance) {
		maxConcurrentRequestsPerInstance_ := 1
		tmp.MaxConcurrentRequestsPerInstance = &maxConcurrentRequestsPerInstance_
	}
	if isZero(tmp.MaxQueueWait) {
		maxQueueWait_ := "PT0.5S"
		tmp.MaxQueueWait = &maxQueueWait_
	}
	if isZero(tmp.RequestTimeout) {
		requestTimeout_ := "PT5S"
		tmp.RequestTimeout = &requestTimeout_
	}
	return &tmp
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

type PATAuthTypeWorkspaceConnectionProperties struct {
	AuthType    string                                  `pulumi:"authType"`
	Category    *string                                 `pulumi:"category"`
	Credentials *WorkspaceConnectionPersonalAccessToken `pulumi:"credentials"`
	Target      *string                                 `pulumi:"target"`
	Value       *string                                 `pulumi:"value"`
	ValueFormat *string                                 `pulumi:"valueFormat"`
}

type PATAuthTypeWorkspaceConnectionPropertiesResponse struct {
	AuthType    string                                          `pulumi:"authType"`
	Category    *string                                         `pulumi:"category"`
	Credentials *WorkspaceConnectionPersonalAccessTokenResponse `pulumi:"credentials"`
	Target      *string                                         `pulumi:"target"`
	Value       *string                                         `pulumi:"value"`
	ValueFormat *string                                         `pulumi:"valueFormat"`
}

type PasswordResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type PasswordResponseOutput struct{ *pulumi.OutputState }

func (PasswordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PasswordResponse)(nil)).Elem()
}

func (o PasswordResponseOutput) ToPasswordResponseOutput() PasswordResponseOutput {
	return o
}

func (o PasswordResponseOutput) ToPasswordResponseOutputWithContext(ctx context.Context) PasswordResponseOutput {
	return o
}

func (o PasswordResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PasswordResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PasswordResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v PasswordResponse) string { return v.Value }).(pulumi.StringOutput)
}

type PasswordResponseArrayOutput struct{ *pulumi.OutputState }

func (PasswordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PasswordResponse)(nil)).Elem()
}

func (o PasswordResponseArrayOutput) ToPasswordResponseArrayOutput() PasswordResponseArrayOutput {
	return o
}

func (o PasswordResponseArrayOutput) ToPasswordResponseArrayOutputWithContext(ctx context.Context) PasswordResponseArrayOutput {
	return o
}

func (o PasswordResponseArrayOutput) Index(i pulumi.IntInput) PasswordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PasswordResponse {
		return vs[0].([]PasswordResponse)[vs[1].(int)]
	}).(PasswordResponseOutput)
}

type PersonalComputeInstanceSettings struct {
	AssignedUser *AssignedUser `pulumi:"assignedUser"`
}

type PersonalComputeInstanceSettingsResponse struct {
	AssignedUser *AssignedUserResponse `pulumi:"assignedUser"`
}

type PipelineJob struct {
	ComputeId      *string                `pulumi:"computeId"`
	Description    *string                `pulumi:"description"`
	DisplayName    *string                `pulumi:"displayName"`
	ExperimentName *string                `pulumi:"experimentName"`
	Identity       interface{}            `pulumi:"identity"`
	Inputs         map[string]interface{} `pulumi:"inputs"`
	IsArchived     *bool                  `pulumi:"isArchived"`
	JobType        string                 `pulumi:"jobType"`
	Jobs           map[string]interface{} `pulumi:"jobs"`
	Outputs        map[string]interface{} `pulumi:"outputs"`
	Properties     map[string]string      `pulumi:"properties"`
	Services       map[string]JobService  `pulumi:"services"`
	Settings       interface{}            `pulumi:"settings"`
	Tags           map[string]string      `pulumi:"tags"`
}


func (val *PipelineJob) Defaults() *PipelineJob {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExperimentName) {
		experimentName_ := "Default"
		tmp.ExperimentName = &experimentName_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}

type PipelineJobResponse struct {
	ComputeId      *string                       `pulumi:"computeId"`
	Description    *string                       `pulumi:"description"`
	DisplayName    *string                       `pulumi:"displayName"`
	ExperimentName *string                       `pulumi:"experimentName"`
	Identity       interface{}                   `pulumi:"identity"`
	Inputs         map[string]interface{}        `pulumi:"inputs"`
	IsArchived     *bool                         `pulumi:"isArchived"`
	JobType        string                        `pulumi:"jobType"`
	Jobs           map[string]interface{}        `pulumi:"jobs"`
	Outputs        map[string]interface{}        `pulumi:"outputs"`
	Properties     map[string]string             `pulumi:"properties"`
	Services       map[string]JobServiceResponse `pulumi:"services"`
	Settings       interface{}                   `pulumi:"settings"`
	Status         string                        `pulumi:"status"`
	Tags           map[string]string             `pulumi:"tags"`
}


func (val *PipelineJobResponse) Defaults() *PipelineJobResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExperimentName) {
		experimentName_ := "Default"
		tmp.ExperimentName = &experimentName_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Identity                          *ManagedServiceIdentityResponse           `pulumi:"identity"`
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

func (o PrivateEndpointConnectionResponseOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
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


func (val *ProbeSettings) Defaults() *ProbeSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.FailureThreshold) {
		failureThreshold_ := 30
		tmp.FailureThreshold = &failureThreshold_
	}
	if isZero(tmp.Period) {
		period_ := "PT10S"
		tmp.Period = &period_
	}
	if isZero(tmp.SuccessThreshold) {
		successThreshold_ := 1
		tmp.SuccessThreshold = &successThreshold_
	}
	if isZero(tmp.Timeout) {
		timeout_ := "PT2S"
		tmp.Timeout = &timeout_
	}
	return &tmp
}

type ProbeSettingsResponse struct {
	FailureThreshold *int    `pulumi:"failureThreshold"`
	InitialDelay     *string `pulumi:"initialDelay"`
	Period           *string `pulumi:"period"`
	SuccessThreshold *int    `pulumi:"successThreshold"`
	Timeout          *string `pulumi:"timeout"`
}


func (val *ProbeSettingsResponse) Defaults() *ProbeSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.FailureThreshold) {
		failureThreshold_ := 30
		tmp.FailureThreshold = &failureThreshold_
	}
	if isZero(tmp.Period) {
		period_ := "PT10S"
		tmp.Period = &period_
	}
	if isZero(tmp.SuccessThreshold) {
		successThreshold_ := 1
		tmp.SuccessThreshold = &successThreshold_
	}
	if isZero(tmp.Timeout) {
		timeout_ := "PT2S"
		tmp.Timeout = &timeout_
	}
	return &tmp
}

type PyTorch struct {
	DistributionType        string `pulumi:"distributionType"`
	ProcessCountPerInstance *int   `pulumi:"processCountPerInstance"`
}

type PyTorchResponse struct {
	DistributionType        string `pulumi:"distributionType"`
	ProcessCountPerInstance *int   `pulumi:"processCountPerInstance"`
}

type RandomSamplingAlgorithm struct {
	Rule                  *string `pulumi:"rule"`
	SamplingAlgorithmType string  `pulumi:"samplingAlgorithmType"`
	Seed                  *int    `pulumi:"seed"`
}


func (val *RandomSamplingAlgorithm) Defaults() *RandomSamplingAlgorithm {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Rule) {
		rule_ := "Random"
		tmp.Rule = &rule_
	}
	return &tmp
}

type RandomSamplingAlgorithmResponse struct {
	Rule                  *string `pulumi:"rule"`
	SamplingAlgorithmType string  `pulumi:"samplingAlgorithmType"`
	Seed                  *int    `pulumi:"seed"`
}


func (val *RandomSamplingAlgorithmResponse) Defaults() *RandomSamplingAlgorithmResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Rule) {
		rule_ := "Random"
		tmp.Rule = &rule_
	}
	return &tmp
}

type RegistryListCredentialsResultResponse struct {
	Location  string             `pulumi:"location"`
	Passwords []PasswordResponse `pulumi:"passwords"`
	Username  string             `pulumi:"username"`
}

type RegistryListCredentialsResultResponseOutput struct{ *pulumi.OutputState }

func (RegistryListCredentialsResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryListCredentialsResultResponse)(nil)).Elem()
}

func (o RegistryListCredentialsResultResponseOutput) ToRegistryListCredentialsResultResponseOutput() RegistryListCredentialsResultResponseOutput {
	return o
}

func (o RegistryListCredentialsResultResponseOutput) ToRegistryListCredentialsResultResponseOutputWithContext(ctx context.Context) RegistryListCredentialsResultResponseOutput {
	return o
}

func (o RegistryListCredentialsResultResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v RegistryListCredentialsResultResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o RegistryListCredentialsResultResponseOutput) Passwords() PasswordResponseArrayOutput {
	return o.ApplyT(func(v RegistryListCredentialsResultResponse) []PasswordResponse { return v.Passwords }).(PasswordResponseArrayOutput)
}

func (o RegistryListCredentialsResultResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v RegistryListCredentialsResultResponse) string { return v.Username }).(pulumi.StringOutput)
}

type ResourceConfiguration struct {
	InstanceCount *int                   `pulumi:"instanceCount"`
	InstanceType  *string                `pulumi:"instanceType"`
	Properties    map[string]interface{} `pulumi:"properties"`
}


func (val *ResourceConfiguration) Defaults() *ResourceConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.InstanceCount) {
		instanceCount_ := 1
		tmp.InstanceCount = &instanceCount_
	}
	return &tmp
}





type ResourceConfigurationInput interface {
	pulumi.Input

	ToResourceConfigurationOutput() ResourceConfigurationOutput
	ToResourceConfigurationOutputWithContext(context.Context) ResourceConfigurationOutput
}

type ResourceConfigurationArgs struct {
	InstanceCount pulumi.IntPtrInput    `pulumi:"instanceCount"`
	InstanceType  pulumi.StringPtrInput `pulumi:"instanceType"`
	Properties    pulumi.MapInput       `pulumi:"properties"`
}


func (val *ResourceConfigurationArgs) Defaults() *ResourceConfigurationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.InstanceCount) {
		tmp.InstanceCount = pulumi.IntPtr(1)
	}
	return &tmp
}
func (ResourceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceConfiguration)(nil)).Elem()
}

func (i ResourceConfigurationArgs) ToResourceConfigurationOutput() ResourceConfigurationOutput {
	return i.ToResourceConfigurationOutputWithContext(context.Background())
}

func (i ResourceConfigurationArgs) ToResourceConfigurationOutputWithContext(ctx context.Context) ResourceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceConfigurationOutput)
}

func (i ResourceConfigurationArgs) ToResourceConfigurationPtrOutput() ResourceConfigurationPtrOutput {
	return i.ToResourceConfigurationPtrOutputWithContext(context.Background())
}

func (i ResourceConfigurationArgs) ToResourceConfigurationPtrOutputWithContext(ctx context.Context) ResourceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceConfigurationOutput).ToResourceConfigurationPtrOutputWithContext(ctx)
}









type ResourceConfigurationPtrInput interface {
	pulumi.Input

	ToResourceConfigurationPtrOutput() ResourceConfigurationPtrOutput
	ToResourceConfigurationPtrOutputWithContext(context.Context) ResourceConfigurationPtrOutput
}

type resourceConfigurationPtrType ResourceConfigurationArgs

func ResourceConfigurationPtr(v *ResourceConfigurationArgs) ResourceConfigurationPtrInput {
	return (*resourceConfigurationPtrType)(v)
}

func (*resourceConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceConfiguration)(nil)).Elem()
}

func (i *resourceConfigurationPtrType) ToResourceConfigurationPtrOutput() ResourceConfigurationPtrOutput {
	return i.ToResourceConfigurationPtrOutputWithContext(context.Background())
}

func (i *resourceConfigurationPtrType) ToResourceConfigurationPtrOutputWithContext(ctx context.Context) ResourceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceConfigurationPtrOutput)
}

type ResourceConfigurationOutput struct{ *pulumi.OutputState }

func (ResourceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceConfiguration)(nil)).Elem()
}

func (o ResourceConfigurationOutput) ToResourceConfigurationOutput() ResourceConfigurationOutput {
	return o
}

func (o ResourceConfigurationOutput) ToResourceConfigurationOutputWithContext(ctx context.Context) ResourceConfigurationOutput {
	return o
}

func (o ResourceConfigurationOutput) ToResourceConfigurationPtrOutput() ResourceConfigurationPtrOutput {
	return o.ToResourceConfigurationPtrOutputWithContext(context.Background())
}

func (o ResourceConfigurationOutput) ToResourceConfigurationPtrOutputWithContext(ctx context.Context) ResourceConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceConfiguration) *ResourceConfiguration {
		return &v
	}).(ResourceConfigurationPtrOutput)
}

func (o ResourceConfigurationOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceConfiguration) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

func (o ResourceConfigurationOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceConfiguration) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o ResourceConfigurationOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func(v ResourceConfiguration) map[string]interface{} { return v.Properties }).(pulumi.MapOutput)
}

type ResourceConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ResourceConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceConfiguration)(nil)).Elem()
}

func (o ResourceConfigurationPtrOutput) ToResourceConfigurationPtrOutput() ResourceConfigurationPtrOutput {
	return o
}

func (o ResourceConfigurationPtrOutput) ToResourceConfigurationPtrOutputWithContext(ctx context.Context) ResourceConfigurationPtrOutput {
	return o
}

func (o ResourceConfigurationPtrOutput) Elem() ResourceConfigurationOutput {
	return o.ApplyT(func(v *ResourceConfiguration) ResourceConfiguration {
		if v != nil {
			return *v
		}
		var ret ResourceConfiguration
		return ret
	}).(ResourceConfigurationOutput)
}

func (o ResourceConfigurationPtrOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.InstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o ResourceConfigurationPtrOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.InstanceType
	}).(pulumi.StringPtrOutput)
}

func (o ResourceConfigurationPtrOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func(v *ResourceConfiguration) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.MapOutput)
}

type ResourceConfigurationResponse struct {
	InstanceCount *int                   `pulumi:"instanceCount"`
	InstanceType  *string                `pulumi:"instanceType"`
	Properties    map[string]interface{} `pulumi:"properties"`
}


func (val *ResourceConfigurationResponse) Defaults() *ResourceConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.InstanceCount) {
		instanceCount_ := 1
		tmp.InstanceCount = &instanceCount_
	}
	return &tmp
}

type ResourceConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ResourceConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceConfigurationResponse)(nil)).Elem()
}

func (o ResourceConfigurationResponseOutput) ToResourceConfigurationResponseOutput() ResourceConfigurationResponseOutput {
	return o
}

func (o ResourceConfigurationResponseOutput) ToResourceConfigurationResponseOutputWithContext(ctx context.Context) ResourceConfigurationResponseOutput {
	return o
}

func (o ResourceConfigurationResponseOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceConfigurationResponse) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

func (o ResourceConfigurationResponseOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceConfigurationResponse) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o ResourceConfigurationResponseOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func(v ResourceConfigurationResponse) map[string]interface{} { return v.Properties }).(pulumi.MapOutput)
}

type ResourceConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceConfigurationResponse)(nil)).Elem()
}

func (o ResourceConfigurationResponsePtrOutput) ToResourceConfigurationResponsePtrOutput() ResourceConfigurationResponsePtrOutput {
	return o
}

func (o ResourceConfigurationResponsePtrOutput) ToResourceConfigurationResponsePtrOutputWithContext(ctx context.Context) ResourceConfigurationResponsePtrOutput {
	return o
}

func (o ResourceConfigurationResponsePtrOutput) Elem() ResourceConfigurationResponseOutput {
	return o.ApplyT(func(v *ResourceConfigurationResponse) ResourceConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ResourceConfigurationResponse
		return ret
	}).(ResourceConfigurationResponseOutput)
}

func (o ResourceConfigurationResponsePtrOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.InstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o ResourceConfigurationResponsePtrOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.InstanceType
	}).(pulumi.StringPtrOutput)
}

func (o ResourceConfigurationResponsePtrOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func(v *ResourceConfigurationResponse) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.MapOutput)
}

type ResourceId struct {
	Id string `pulumi:"id"`
}

type ResourceIdResponse struct {
	Id string `pulumi:"id"`
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

type SASAuthTypeWorkspaceConnectionProperties struct {
	AuthType    string                                    `pulumi:"authType"`
	Category    *string                                   `pulumi:"category"`
	Credentials *WorkspaceConnectionSharedAccessSignature `pulumi:"credentials"`
	Target      *string                                   `pulumi:"target"`
	Value       *string                                   `pulumi:"value"`
	ValueFormat *string                                   `pulumi:"valueFormat"`
}

type SASAuthTypeWorkspaceConnectionPropertiesResponse struct {
	AuthType    string                                            `pulumi:"authType"`
	Category    *string                                           `pulumi:"category"`
	Credentials *WorkspaceConnectionSharedAccessSignatureResponse `pulumi:"credentials"`
	Target      *string                                           `pulumi:"target"`
	Value       *string                                           `pulumi:"value"`
	ValueFormat *string                                           `pulumi:"valueFormat"`
}

type SasDatastoreCredentials struct {
	CredentialsType string              `pulumi:"credentialsType"`
	Secrets         SasDatastoreSecrets `pulumi:"secrets"`
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

type ScheduleBaseResponse struct {
	Id                 *string `pulumi:"id"`
	ProvisioningStatus *string `pulumi:"provisioningStatus"`
	Status             *string `pulumi:"status"`
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
	AuthorityUrl    *string                          `pulumi:"authorityUrl"`
	ClientId        string                           `pulumi:"clientId"`
	CredentialsType string                           `pulumi:"credentialsType"`
	ResourceUrl     *string                          `pulumi:"resourceUrl"`
	Secrets         ServicePrincipalDatastoreSecrets `pulumi:"secrets"`
	TenantId        string                           `pulumi:"tenantId"`
}

type ServicePrincipalDatastoreCredentialsResponse struct {
	AuthorityUrl    *string `pulumi:"authorityUrl"`
	ClientId        string  `pulumi:"clientId"`
	CredentialsType string  `pulumi:"credentialsType"`
	ResourceUrl     *string `pulumi:"resourceUrl"`
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
	Capacity *int     `pulumi:"capacity"`
	Family   *string  `pulumi:"family"`
	Name     string   `pulumi:"name"`
	Size     *string  `pulumi:"size"`
	Tier     *SkuTier `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     SkuTierPtrInput       `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v Sku) *SkuTier { return v.Tier }).(SkuTierPtrOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v *Sku) *SkuTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(SkuTierPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
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

type SweepJob struct {
	ComputeId         *string                `pulumi:"computeId"`
	Description       *string                `pulumi:"description"`
	DisplayName       *string                `pulumi:"displayName"`
	EarlyTermination  interface{}            `pulumi:"earlyTermination"`
	ExperimentName    *string                `pulumi:"experimentName"`
	Identity          interface{}            `pulumi:"identity"`
	Inputs            map[string]interface{} `pulumi:"inputs"`
	IsArchived        *bool                  `pulumi:"isArchived"`
	JobType           string                 `pulumi:"jobType"`
	Limits            *SweepJobLimits        `pulumi:"limits"`
	Objective         Objective              `pulumi:"objective"`
	Outputs           map[string]interface{} `pulumi:"outputs"`
	Properties        map[string]string      `pulumi:"properties"`
	SamplingAlgorithm interface{}            `pulumi:"samplingAlgorithm"`
	SearchSpace       interface{}            `pulumi:"searchSpace"`
	Services          map[string]JobService  `pulumi:"services"`
	Tags              map[string]string      `pulumi:"tags"`
	Trial             TrialComponent         `pulumi:"trial"`
}


func (val *SweepJob) Defaults() *SweepJob {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExperimentName) {
		experimentName_ := "Default"
		tmp.ExperimentName = &experimentName_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	tmp.Trial = *tmp.Trial.Defaults()

	return &tmp
}

type SweepJobLimits struct {
	JobLimitsType       string  `pulumi:"jobLimitsType"`
	MaxConcurrentTrials *int    `pulumi:"maxConcurrentTrials"`
	MaxTotalTrials      *int    `pulumi:"maxTotalTrials"`
	Timeout             *string `pulumi:"timeout"`
	TrialTimeout        *string `pulumi:"trialTimeout"`
}

type SweepJobLimitsResponse struct {
	JobLimitsType       string  `pulumi:"jobLimitsType"`
	MaxConcurrentTrials *int    `pulumi:"maxConcurrentTrials"`
	MaxTotalTrials      *int    `pulumi:"maxTotalTrials"`
	Timeout             *string `pulumi:"timeout"`
	TrialTimeout        *string `pulumi:"trialTimeout"`
}

type SweepJobResponse struct {
	ComputeId         *string                       `pulumi:"computeId"`
	Description       *string                       `pulumi:"description"`
	DisplayName       *string                       `pulumi:"displayName"`
	EarlyTermination  interface{}                   `pulumi:"earlyTermination"`
	ExperimentName    *string                       `pulumi:"experimentName"`
	Identity          interface{}                   `pulumi:"identity"`
	Inputs            map[string]interface{}        `pulumi:"inputs"`
	IsArchived        *bool                         `pulumi:"isArchived"`
	JobType           string                        `pulumi:"jobType"`
	Limits            *SweepJobLimitsResponse       `pulumi:"limits"`
	Objective         ObjectiveResponse             `pulumi:"objective"`
	Outputs           map[string]interface{}        `pulumi:"outputs"`
	Properties        map[string]string             `pulumi:"properties"`
	SamplingAlgorithm interface{}                   `pulumi:"samplingAlgorithm"`
	SearchSpace       interface{}                   `pulumi:"searchSpace"`
	Services          map[string]JobServiceResponse `pulumi:"services"`
	Status            string                        `pulumi:"status"`
	Tags              map[string]string             `pulumi:"tags"`
	Trial             TrialComponentResponse        `pulumi:"trial"`
}


func (val *SweepJobResponse) Defaults() *SweepJobResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExperimentName) {
		experimentName_ := "Default"
		tmp.ExperimentName = &experimentName_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	tmp.Trial = *tmp.Trial.Defaults()

	return &tmp
}

type SynapseSpark struct {
	ComputeType      string                  `pulumi:"computeType"`
	Description      *string                 `pulumi:"description"`
	DisableLocalAuth *bool                   `pulumi:"disableLocalAuth"`
	Properties       *SynapseSparkProperties `pulumi:"properties"`
	ResourceId       *string                 `pulumi:"resourceId"`
}

type SynapseSparkProperties struct {
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

type SynapseSparkResponse struct {
	ComputeLocation    string                          `pulumi:"computeLocation"`
	ComputeType        string                          `pulumi:"computeType"`
	CreatedOn          string                          `pulumi:"createdOn"`
	Description        *string                         `pulumi:"description"`
	DisableLocalAuth   *bool                           `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                            `pulumi:"isAttachedCompute"`
	ModifiedOn         string                          `pulumi:"modifiedOn"`
	Properties         *SynapseSparkResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse         `pulumi:"provisioningErrors"`
	ProvisioningState  string                          `pulumi:"provisioningState"`
	ResourceId         *string                         `pulumi:"resourceId"`
}

type SynapseSparkResponseProperties struct {
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

type TargetUtilizationScaleSettings struct {
	MaxInstances                *int    `pulumi:"maxInstances"`
	MinInstances                *int    `pulumi:"minInstances"`
	PollingInterval             *string `pulumi:"pollingInterval"`
	ScaleType                   string  `pulumi:"scaleType"`
	TargetUtilizationPercentage *int    `pulumi:"targetUtilizationPercentage"`
}


func (val *TargetUtilizationScaleSettings) Defaults() *TargetUtilizationScaleSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxInstances) {
		maxInstances_ := 1
		tmp.MaxInstances = &maxInstances_
	}
	if isZero(tmp.MinInstances) {
		minInstances_ := 1
		tmp.MinInstances = &minInstances_
	}
	if isZero(tmp.PollingInterval) {
		pollingInterval_ := "PT1S"
		tmp.PollingInterval = &pollingInterval_
	}
	if isZero(tmp.TargetUtilizationPercentage) {
		targetUtilizationPercentage_ := 70
		tmp.TargetUtilizationPercentage = &targetUtilizationPercentage_
	}
	return &tmp
}

type TargetUtilizationScaleSettingsResponse struct {
	MaxInstances                *int    `pulumi:"maxInstances"`
	MinInstances                *int    `pulumi:"minInstances"`
	PollingInterval             *string `pulumi:"pollingInterval"`
	ScaleType                   string  `pulumi:"scaleType"`
	TargetUtilizationPercentage *int    `pulumi:"targetUtilizationPercentage"`
}


func (val *TargetUtilizationScaleSettingsResponse) Defaults() *TargetUtilizationScaleSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxInstances) {
		maxInstances_ := 1
		tmp.MaxInstances = &maxInstances_
	}
	if isZero(tmp.MinInstances) {
		minInstances_ := 1
		tmp.MinInstances = &minInstances_
	}
	if isZero(tmp.PollingInterval) {
		pollingInterval_ := "PT1S"
		tmp.PollingInterval = &pollingInterval_
	}
	if isZero(tmp.TargetUtilizationPercentage) {
		targetUtilizationPercentage_ := 70
		tmp.TargetUtilizationPercentage = &targetUtilizationPercentage_
	}
	return &tmp
}

type TensorFlow struct {
	DistributionType     string `pulumi:"distributionType"`
	ParameterServerCount *int   `pulumi:"parameterServerCount"`
	WorkerCount          *int   `pulumi:"workerCount"`
}


func (val *TensorFlow) Defaults() *TensorFlow {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ParameterServerCount) {
		parameterServerCount_ := 0
		tmp.ParameterServerCount = &parameterServerCount_
	}
	return &tmp
}

type TensorFlowResponse struct {
	DistributionType     string `pulumi:"distributionType"`
	ParameterServerCount *int   `pulumi:"parameterServerCount"`
	WorkerCount          *int   `pulumi:"workerCount"`
}


func (val *TensorFlowResponse) Defaults() *TensorFlowResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ParameterServerCount) {
		parameterServerCount_ := 0
		tmp.ParameterServerCount = &parameterServerCount_
	}
	return &tmp
}

type TrialComponent struct {
	CodeId               *string                `pulumi:"codeId"`
	Command              string                 `pulumi:"command"`
	Distribution         interface{}            `pulumi:"distribution"`
	EnvironmentId        string                 `pulumi:"environmentId"`
	EnvironmentVariables map[string]string      `pulumi:"environmentVariables"`
	Resources            *ResourceConfiguration `pulumi:"resources"`
}


func (val *TrialComponent) Defaults() *TrialComponent {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Resources = tmp.Resources.Defaults()

	return &tmp
}

type TrialComponentResponse struct {
	CodeId               *string                        `pulumi:"codeId"`
	Command              string                         `pulumi:"command"`
	Distribution         interface{}                    `pulumi:"distribution"`
	EnvironmentId        string                         `pulumi:"environmentId"`
	EnvironmentVariables map[string]string              `pulumi:"environmentVariables"`
	Resources            *ResourceConfigurationResponse `pulumi:"resources"`
}


func (val *TrialComponentResponse) Defaults() *TrialComponentResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Resources = tmp.Resources.Defaults()

	return &tmp
}

type TritonModelJobInput struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *TritonModelJobInput) Defaults() *TritonModelJobInput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type TritonModelJobInputResponse struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *TritonModelJobInputResponse) Defaults() *TritonModelJobInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type TritonModelJobOutput struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *TritonModelJobOutput) Defaults() *TritonModelJobOutput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type TritonModelJobOutputResponse struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *TritonModelJobOutputResponse) Defaults() *TritonModelJobOutputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type TruncationSelectionPolicy struct {
	DelayEvaluation      *int   `pulumi:"delayEvaluation"`
	EvaluationInterval   *int   `pulumi:"evaluationInterval"`
	PolicyType           string `pulumi:"policyType"`
	TruncationPercentage *int   `pulumi:"truncationPercentage"`
}


func (val *TruncationSelectionPolicy) Defaults() *TruncationSelectionPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DelayEvaluation) {
		delayEvaluation_ := 0
		tmp.DelayEvaluation = &delayEvaluation_
	}
	if isZero(tmp.EvaluationInterval) {
		evaluationInterval_ := 0
		tmp.EvaluationInterval = &evaluationInterval_
	}
	if isZero(tmp.TruncationPercentage) {
		truncationPercentage_ := 0
		tmp.TruncationPercentage = &truncationPercentage_
	}
	return &tmp
}

type TruncationSelectionPolicyResponse struct {
	DelayEvaluation      *int   `pulumi:"delayEvaluation"`
	EvaluationInterval   *int   `pulumi:"evaluationInterval"`
	PolicyType           string `pulumi:"policyType"`
	TruncationPercentage *int   `pulumi:"truncationPercentage"`
}


func (val *TruncationSelectionPolicyResponse) Defaults() *TruncationSelectionPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DelayEvaluation) {
		delayEvaluation_ := 0
		tmp.DelayEvaluation = &delayEvaluation_
	}
	if isZero(tmp.EvaluationInterval) {
		evaluationInterval_ := 0
		tmp.EvaluationInterval = &evaluationInterval_
	}
	if isZero(tmp.TruncationPercentage) {
		truncationPercentage_ := 0
		tmp.TruncationPercentage = &truncationPercentage_
	}
	return &tmp
}

type UriFileDataVersion struct {
	DataType    string            `pulumi:"dataType"`
	DataUri     string            `pulumi:"dataUri"`
	Description *string           `pulumi:"description"`
	IsAnonymous *bool             `pulumi:"isAnonymous"`
	IsArchived  *bool             `pulumi:"isArchived"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}


func (val *UriFileDataVersion) Defaults() *UriFileDataVersion {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}

type UriFileDataVersionResponse struct {
	DataType    string            `pulumi:"dataType"`
	DataUri     string            `pulumi:"dataUri"`
	Description *string           `pulumi:"description"`
	IsAnonymous *bool             `pulumi:"isAnonymous"`
	IsArchived  *bool             `pulumi:"isArchived"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}


func (val *UriFileDataVersionResponse) Defaults() *UriFileDataVersionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}

type UriFileJobInput struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *UriFileJobInput) Defaults() *UriFileJobInput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type UriFileJobInputResponse struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *UriFileJobInputResponse) Defaults() *UriFileJobInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type UriFileJobOutput struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *UriFileJobOutput) Defaults() *UriFileJobOutput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type UriFileJobOutputResponse struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *UriFileJobOutputResponse) Defaults() *UriFileJobOutputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type UriFolderDataVersion struct {
	DataType    string            `pulumi:"dataType"`
	DataUri     string            `pulumi:"dataUri"`
	Description *string           `pulumi:"description"`
	IsAnonymous *bool             `pulumi:"isAnonymous"`
	IsArchived  *bool             `pulumi:"isArchived"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}


func (val *UriFolderDataVersion) Defaults() *UriFolderDataVersion {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}

type UriFolderDataVersionResponse struct {
	DataType    string            `pulumi:"dataType"`
	DataUri     string            `pulumi:"dataUri"`
	Description *string           `pulumi:"description"`
	IsAnonymous *bool             `pulumi:"isAnonymous"`
	IsArchived  *bool             `pulumi:"isArchived"`
	Properties  map[string]string `pulumi:"properties"`
	Tags        map[string]string `pulumi:"tags"`
}


func (val *UriFolderDataVersionResponse) Defaults() *UriFolderDataVersionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsAnonymous) {
		isAnonymous_ := false
		tmp.IsAnonymous = &isAnonymous_
	}
	if isZero(tmp.IsArchived) {
		isArchived_ := false
		tmp.IsArchived = &isArchived_
	}
	return &tmp
}

type UriFolderJobInput struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *UriFolderJobInput) Defaults() *UriFolderJobInput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type UriFolderJobInputResponse struct {
	Description  *string `pulumi:"description"`
	JobInputType string  `pulumi:"jobInputType"`
	Mode         *string `pulumi:"mode"`
	Uri          string  `pulumi:"uri"`
}


func (val *UriFolderJobInputResponse) Defaults() *UriFolderJobInputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadOnlyMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type UriFolderJobOutput struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *UriFolderJobOutput) Defaults() *UriFolderJobOutput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
}

type UriFolderJobOutputResponse struct {
	Description   *string `pulumi:"description"`
	JobOutputType string  `pulumi:"jobOutputType"`
	Mode          *string `pulumi:"mode"`
	Uri           *string `pulumi:"uri"`
}


func (val *UriFolderJobOutputResponse) Defaults() *UriFolderJobOutputResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "ReadWriteMount"
		tmp.Mode = &mode_
	}
	return &tmp
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

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
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

type UserIdentity struct {
	IdentityType string `pulumi:"identityType"`
}

type UserIdentityResponse struct {
	IdentityType string `pulumi:"identityType"`
}

type UsernamePasswordAuthTypeWorkspaceConnectionProperties struct {
	AuthType    string                               `pulumi:"authType"`
	Category    *string                              `pulumi:"category"`
	Credentials *WorkspaceConnectionUsernamePassword `pulumi:"credentials"`
	Target      *string                              `pulumi:"target"`
	Value       *string                              `pulumi:"value"`
	ValueFormat *string                              `pulumi:"valueFormat"`
}

type UsernamePasswordAuthTypeWorkspaceConnectionPropertiesResponse struct {
	AuthType    string                                       `pulumi:"authType"`
	Category    *string                                      `pulumi:"category"`
	Credentials *WorkspaceConnectionUsernamePasswordResponse `pulumi:"credentials"`
	Target      *string                                      `pulumi:"target"`
	Value       *string                                      `pulumi:"value"`
	ValueFormat *string                                      `pulumi:"valueFormat"`
}

type VirtualMachine struct {
	ComputeType      string                          `pulumi:"computeType"`
	Description      *string                         `pulumi:"description"`
	DisableLocalAuth *bool                           `pulumi:"disableLocalAuth"`
	Properties       *VirtualMachineSchemaProperties `pulumi:"properties"`
	ResourceId       *string                         `pulumi:"resourceId"`
}

type VirtualMachineImage struct {
	Id string `pulumi:"id"`
}

type VirtualMachineImageResponse struct {
	Id string `pulumi:"id"`
}

type VirtualMachineResponse struct {
	ComputeLocation    string                                  `pulumi:"computeLocation"`
	ComputeType        string                                  `pulumi:"computeType"`
	CreatedOn          string                                  `pulumi:"createdOn"`
	Description        *string                                 `pulumi:"description"`
	DisableLocalAuth   *bool                                   `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                                    `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                  `pulumi:"modifiedOn"`
	Properties         *VirtualMachineSchemaResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse                 `pulumi:"provisioningErrors"`
	ProvisioningState  string                                  `pulumi:"provisioningState"`
	ResourceId         *string                                 `pulumi:"resourceId"`
}

type VirtualMachineSchemaProperties struct {
	Address                   *string                       `pulumi:"address"`
	AdministratorAccount      *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	IsNotebookInstanceCompute *bool                         `pulumi:"isNotebookInstanceCompute"`
	NotebookServerPort        *int                          `pulumi:"notebookServerPort"`
	SshPort                   *int                          `pulumi:"sshPort"`
	VirtualMachineSize        *string                       `pulumi:"virtualMachineSize"`
}

type VirtualMachineSchemaResponseProperties struct {
	Address                   *string                               `pulumi:"address"`
	AdministratorAccount      *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	IsNotebookInstanceCompute *bool                                 `pulumi:"isNotebookInstanceCompute"`
	NotebookServerPort        *int                                  `pulumi:"notebookServerPort"`
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

type WorkspaceConnectionManagedIdentity struct {
	ClientId   *string `pulumi:"clientId"`
	ResourceId *string `pulumi:"resourceId"`
}

type WorkspaceConnectionManagedIdentityResponse struct {
	ClientId   *string `pulumi:"clientId"`
	ResourceId *string `pulumi:"resourceId"`
}

type WorkspaceConnectionPersonalAccessToken struct {
	Pat *string `pulumi:"pat"`
}

type WorkspaceConnectionPersonalAccessTokenResponse struct {
	Pat *string `pulumi:"pat"`
}

type WorkspaceConnectionSharedAccessSignature struct {
	Sas *string `pulumi:"sas"`
}

type WorkspaceConnectionSharedAccessSignatureResponse struct {
	Sas *string `pulumi:"sas"`
}

type WorkspaceConnectionUsernamePassword struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}

type WorkspaceConnectionUsernamePasswordResponse struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}

func init() {
	pulumi.RegisterOutputType(AmlComputeNodeInformationResponseOutput{})
	pulumi.RegisterOutputType(AmlComputeNodeInformationResponseArrayOutput{})
	pulumi.RegisterOutputType(BatchDeploymentTypeOutput{})
	pulumi.RegisterOutputType(BatchDeploymentResponseOutput{})
	pulumi.RegisterOutputType(BatchEndpointTypeOutput{})
	pulumi.RegisterOutputType(BatchEndpointDefaultsOutput{})
	pulumi.RegisterOutputType(BatchEndpointDefaultsPtrOutput{})
	pulumi.RegisterOutputType(BatchEndpointDefaultsResponseOutput{})
	pulumi.RegisterOutputType(BatchEndpointDefaultsResponsePtrOutput{})
	pulumi.RegisterOutputType(BatchEndpointResponseOutput{})
	pulumi.RegisterOutputType(BatchRetrySettingsOutput{})
	pulumi.RegisterOutputType(BatchRetrySettingsPtrOutput{})
	pulumi.RegisterOutputType(BatchRetrySettingsResponseOutput{})
	pulumi.RegisterOutputType(BatchRetrySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(BuildContextOutput{})
	pulumi.RegisterOutputType(BuildContextPtrOutput{})
	pulumi.RegisterOutputType(BuildContextResponseOutput{})
	pulumi.RegisterOutputType(BuildContextResponsePtrOutput{})
	pulumi.RegisterOutputType(CodeConfigurationOutput{})
	pulumi.RegisterOutputType(CodeConfigurationPtrOutput{})
	pulumi.RegisterOutputType(CodeConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CodeConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(CodeContainerTypeOutput{})
	pulumi.RegisterOutputType(CodeContainerResponseOutput{})
	pulumi.RegisterOutputType(CodeVersionTypeOutput{})
	pulumi.RegisterOutputType(CodeVersionResponseOutput{})
	pulumi.RegisterOutputType(ComponentContainerTypeOutput{})
	pulumi.RegisterOutputType(ComponentContainerResponseOutput{})
	pulumi.RegisterOutputType(ComponentVersionTypeOutput{})
	pulumi.RegisterOutputType(ComponentVersionResponseOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsPtrOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsResponseOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(DataContainerTypeOutput{})
	pulumi.RegisterOutputType(DataContainerResponseOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointAuthKeysOutput{})
	pulumi.RegisterOutputType(EndpointAuthKeysPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentContainerTypeOutput{})
	pulumi.RegisterOutputType(EnvironmentContainerResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentVersionTypeOutput{})
	pulumi.RegisterOutputType(EnvironmentVersionResponseOutput{})
	pulumi.RegisterOutputType(FlavorDataOutput{})
	pulumi.RegisterOutputType(FlavorDataMapOutput{})
	pulumi.RegisterOutputType(FlavorDataResponseOutput{})
	pulumi.RegisterOutputType(FlavorDataResponseMapOutput{})
	pulumi.RegisterOutputType(IdentityForCmkOutput{})
	pulumi.RegisterOutputType(IdentityForCmkPtrOutput{})
	pulumi.RegisterOutputType(IdentityForCmkResponseOutput{})
	pulumi.RegisterOutputType(IdentityForCmkResponsePtrOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ListNotebookKeysResultResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ModelContainerTypeOutput{})
	pulumi.RegisterOutputType(ModelContainerResponseOutput{})
	pulumi.RegisterOutputType(ModelVersionTypeOutput{})
	pulumi.RegisterOutputType(ModelVersionResponseOutput{})
	pulumi.RegisterOutputType(NotebookPreparationErrorResponseOutput{})
	pulumi.RegisterOutputType(NotebookPreparationErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(NotebookResourceInfoResponseOutput{})
	pulumi.RegisterOutputType(OnlineEndpointTypeOutput{})
	pulumi.RegisterOutputType(OnlineEndpointResponseOutput{})
	pulumi.RegisterOutputType(PasswordResponseOutput{})
	pulumi.RegisterOutputType(PasswordResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(RegistryListCredentialsResultResponseOutput{})
	pulumi.RegisterOutputType(ResourceConfigurationOutput{})
	pulumi.RegisterOutputType(ResourceConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ResourceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ResourceConfigurationResponsePtrOutput{})
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
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}
