


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ACIServiceCreateRequestDataCollection struct {
	EventHubEnabled *bool `pulumi:"eventHubEnabled"`
	StorageEnabled  *bool `pulumi:"storageEnabled"`
}





type ACIServiceCreateRequestDataCollectionInput interface {
	pulumi.Input

	ToACIServiceCreateRequestDataCollectionOutput() ACIServiceCreateRequestDataCollectionOutput
	ToACIServiceCreateRequestDataCollectionOutputWithContext(context.Context) ACIServiceCreateRequestDataCollectionOutput
}

type ACIServiceCreateRequestDataCollectionArgs struct {
	EventHubEnabled pulumi.BoolPtrInput `pulumi:"eventHubEnabled"`
	StorageEnabled  pulumi.BoolPtrInput `pulumi:"storageEnabled"`
}

func (ACIServiceCreateRequestDataCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceCreateRequestDataCollection)(nil)).Elem()
}

func (i ACIServiceCreateRequestDataCollectionArgs) ToACIServiceCreateRequestDataCollectionOutput() ACIServiceCreateRequestDataCollectionOutput {
	return i.ToACIServiceCreateRequestDataCollectionOutputWithContext(context.Background())
}

func (i ACIServiceCreateRequestDataCollectionArgs) ToACIServiceCreateRequestDataCollectionOutputWithContext(ctx context.Context) ACIServiceCreateRequestDataCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceCreateRequestDataCollectionOutput)
}

func (i ACIServiceCreateRequestDataCollectionArgs) ToACIServiceCreateRequestDataCollectionPtrOutput() ACIServiceCreateRequestDataCollectionPtrOutput {
	return i.ToACIServiceCreateRequestDataCollectionPtrOutputWithContext(context.Background())
}

func (i ACIServiceCreateRequestDataCollectionArgs) ToACIServiceCreateRequestDataCollectionPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestDataCollectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceCreateRequestDataCollectionOutput).ToACIServiceCreateRequestDataCollectionPtrOutputWithContext(ctx)
}









type ACIServiceCreateRequestDataCollectionPtrInput interface {
	pulumi.Input

	ToACIServiceCreateRequestDataCollectionPtrOutput() ACIServiceCreateRequestDataCollectionPtrOutput
	ToACIServiceCreateRequestDataCollectionPtrOutputWithContext(context.Context) ACIServiceCreateRequestDataCollectionPtrOutput
}

type aciserviceCreateRequestDataCollectionPtrType ACIServiceCreateRequestDataCollectionArgs

func ACIServiceCreateRequestDataCollectionPtr(v *ACIServiceCreateRequestDataCollectionArgs) ACIServiceCreateRequestDataCollectionPtrInput {
	return (*aciserviceCreateRequestDataCollectionPtrType)(v)
}

func (*aciserviceCreateRequestDataCollectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceCreateRequestDataCollection)(nil)).Elem()
}

func (i *aciserviceCreateRequestDataCollectionPtrType) ToACIServiceCreateRequestDataCollectionPtrOutput() ACIServiceCreateRequestDataCollectionPtrOutput {
	return i.ToACIServiceCreateRequestDataCollectionPtrOutputWithContext(context.Background())
}

func (i *aciserviceCreateRequestDataCollectionPtrType) ToACIServiceCreateRequestDataCollectionPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestDataCollectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceCreateRequestDataCollectionPtrOutput)
}

type ACIServiceCreateRequestDataCollectionOutput struct{ *pulumi.OutputState }

func (ACIServiceCreateRequestDataCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceCreateRequestDataCollection)(nil)).Elem()
}

func (o ACIServiceCreateRequestDataCollectionOutput) ToACIServiceCreateRequestDataCollectionOutput() ACIServiceCreateRequestDataCollectionOutput {
	return o
}

func (o ACIServiceCreateRequestDataCollectionOutput) ToACIServiceCreateRequestDataCollectionOutputWithContext(ctx context.Context) ACIServiceCreateRequestDataCollectionOutput {
	return o
}

func (o ACIServiceCreateRequestDataCollectionOutput) ToACIServiceCreateRequestDataCollectionPtrOutput() ACIServiceCreateRequestDataCollectionPtrOutput {
	return o.ToACIServiceCreateRequestDataCollectionPtrOutputWithContext(context.Background())
}

func (o ACIServiceCreateRequestDataCollectionOutput) ToACIServiceCreateRequestDataCollectionPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestDataCollectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ACIServiceCreateRequestDataCollection) *ACIServiceCreateRequestDataCollection {
		return &v
	}).(ACIServiceCreateRequestDataCollectionPtrOutput)
}

func (o ACIServiceCreateRequestDataCollectionOutput) EventHubEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ACIServiceCreateRequestDataCollection) *bool { return v.EventHubEnabled }).(pulumi.BoolPtrOutput)
}

func (o ACIServiceCreateRequestDataCollectionOutput) StorageEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ACIServiceCreateRequestDataCollection) *bool { return v.StorageEnabled }).(pulumi.BoolPtrOutput)
}

type ACIServiceCreateRequestDataCollectionPtrOutput struct{ *pulumi.OutputState }

func (ACIServiceCreateRequestDataCollectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceCreateRequestDataCollection)(nil)).Elem()
}

func (o ACIServiceCreateRequestDataCollectionPtrOutput) ToACIServiceCreateRequestDataCollectionPtrOutput() ACIServiceCreateRequestDataCollectionPtrOutput {
	return o
}

func (o ACIServiceCreateRequestDataCollectionPtrOutput) ToACIServiceCreateRequestDataCollectionPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestDataCollectionPtrOutput {
	return o
}

func (o ACIServiceCreateRequestDataCollectionPtrOutput) Elem() ACIServiceCreateRequestDataCollectionOutput {
	return o.ApplyT(func(v *ACIServiceCreateRequestDataCollection) ACIServiceCreateRequestDataCollection {
		if v != nil {
			return *v
		}
		var ret ACIServiceCreateRequestDataCollection
		return ret
	}).(ACIServiceCreateRequestDataCollectionOutput)
}

func (o ACIServiceCreateRequestDataCollectionPtrOutput) EventHubEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ACIServiceCreateRequestDataCollection) *bool {
		if v == nil {
			return nil
		}
		return v.EventHubEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ACIServiceCreateRequestDataCollectionPtrOutput) StorageEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ACIServiceCreateRequestDataCollection) *bool {
		if v == nil {
			return nil
		}
		return v.StorageEnabled
	}).(pulumi.BoolPtrOutput)
}

type ACIServiceCreateRequestEncryptionProperties struct {
	KeyName      string `pulumi:"keyName"`
	KeyVersion   string `pulumi:"keyVersion"`
	VaultBaseUrl string `pulumi:"vaultBaseUrl"`
}





type ACIServiceCreateRequestEncryptionPropertiesInput interface {
	pulumi.Input

	ToACIServiceCreateRequestEncryptionPropertiesOutput() ACIServiceCreateRequestEncryptionPropertiesOutput
	ToACIServiceCreateRequestEncryptionPropertiesOutputWithContext(context.Context) ACIServiceCreateRequestEncryptionPropertiesOutput
}

type ACIServiceCreateRequestEncryptionPropertiesArgs struct {
	KeyName      pulumi.StringInput `pulumi:"keyName"`
	KeyVersion   pulumi.StringInput `pulumi:"keyVersion"`
	VaultBaseUrl pulumi.StringInput `pulumi:"vaultBaseUrl"`
}

func (ACIServiceCreateRequestEncryptionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceCreateRequestEncryptionProperties)(nil)).Elem()
}

func (i ACIServiceCreateRequestEncryptionPropertiesArgs) ToACIServiceCreateRequestEncryptionPropertiesOutput() ACIServiceCreateRequestEncryptionPropertiesOutput {
	return i.ToACIServiceCreateRequestEncryptionPropertiesOutputWithContext(context.Background())
}

func (i ACIServiceCreateRequestEncryptionPropertiesArgs) ToACIServiceCreateRequestEncryptionPropertiesOutputWithContext(ctx context.Context) ACIServiceCreateRequestEncryptionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceCreateRequestEncryptionPropertiesOutput)
}

func (i ACIServiceCreateRequestEncryptionPropertiesArgs) ToACIServiceCreateRequestEncryptionPropertiesPtrOutput() ACIServiceCreateRequestEncryptionPropertiesPtrOutput {
	return i.ToACIServiceCreateRequestEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i ACIServiceCreateRequestEncryptionPropertiesArgs) ToACIServiceCreateRequestEncryptionPropertiesPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestEncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceCreateRequestEncryptionPropertiesOutput).ToACIServiceCreateRequestEncryptionPropertiesPtrOutputWithContext(ctx)
}









type ACIServiceCreateRequestEncryptionPropertiesPtrInput interface {
	pulumi.Input

	ToACIServiceCreateRequestEncryptionPropertiesPtrOutput() ACIServiceCreateRequestEncryptionPropertiesPtrOutput
	ToACIServiceCreateRequestEncryptionPropertiesPtrOutputWithContext(context.Context) ACIServiceCreateRequestEncryptionPropertiesPtrOutput
}

type aciserviceCreateRequestEncryptionPropertiesPtrType ACIServiceCreateRequestEncryptionPropertiesArgs

func ACIServiceCreateRequestEncryptionPropertiesPtr(v *ACIServiceCreateRequestEncryptionPropertiesArgs) ACIServiceCreateRequestEncryptionPropertiesPtrInput {
	return (*aciserviceCreateRequestEncryptionPropertiesPtrType)(v)
}

func (*aciserviceCreateRequestEncryptionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceCreateRequestEncryptionProperties)(nil)).Elem()
}

func (i *aciserviceCreateRequestEncryptionPropertiesPtrType) ToACIServiceCreateRequestEncryptionPropertiesPtrOutput() ACIServiceCreateRequestEncryptionPropertiesPtrOutput {
	return i.ToACIServiceCreateRequestEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i *aciserviceCreateRequestEncryptionPropertiesPtrType) ToACIServiceCreateRequestEncryptionPropertiesPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestEncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceCreateRequestEncryptionPropertiesPtrOutput)
}

type ACIServiceCreateRequestEncryptionPropertiesOutput struct{ *pulumi.OutputState }

func (ACIServiceCreateRequestEncryptionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceCreateRequestEncryptionProperties)(nil)).Elem()
}

func (o ACIServiceCreateRequestEncryptionPropertiesOutput) ToACIServiceCreateRequestEncryptionPropertiesOutput() ACIServiceCreateRequestEncryptionPropertiesOutput {
	return o
}

func (o ACIServiceCreateRequestEncryptionPropertiesOutput) ToACIServiceCreateRequestEncryptionPropertiesOutputWithContext(ctx context.Context) ACIServiceCreateRequestEncryptionPropertiesOutput {
	return o
}

func (o ACIServiceCreateRequestEncryptionPropertiesOutput) ToACIServiceCreateRequestEncryptionPropertiesPtrOutput() ACIServiceCreateRequestEncryptionPropertiesPtrOutput {
	return o.ToACIServiceCreateRequestEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (o ACIServiceCreateRequestEncryptionPropertiesOutput) ToACIServiceCreateRequestEncryptionPropertiesPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestEncryptionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ACIServiceCreateRequestEncryptionProperties) *ACIServiceCreateRequestEncryptionProperties {
		return &v
	}).(ACIServiceCreateRequestEncryptionPropertiesPtrOutput)
}

func (o ACIServiceCreateRequestEncryptionPropertiesOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ACIServiceCreateRequestEncryptionProperties) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o ACIServiceCreateRequestEncryptionPropertiesOutput) KeyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ACIServiceCreateRequestEncryptionProperties) string { return v.KeyVersion }).(pulumi.StringOutput)
}

func (o ACIServiceCreateRequestEncryptionPropertiesOutput) VaultBaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ACIServiceCreateRequestEncryptionProperties) string { return v.VaultBaseUrl }).(pulumi.StringOutput)
}

type ACIServiceCreateRequestEncryptionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ACIServiceCreateRequestEncryptionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceCreateRequestEncryptionProperties)(nil)).Elem()
}

func (o ACIServiceCreateRequestEncryptionPropertiesPtrOutput) ToACIServiceCreateRequestEncryptionPropertiesPtrOutput() ACIServiceCreateRequestEncryptionPropertiesPtrOutput {
	return o
}

func (o ACIServiceCreateRequestEncryptionPropertiesPtrOutput) ToACIServiceCreateRequestEncryptionPropertiesPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestEncryptionPropertiesPtrOutput {
	return o
}

func (o ACIServiceCreateRequestEncryptionPropertiesPtrOutput) Elem() ACIServiceCreateRequestEncryptionPropertiesOutput {
	return o.ApplyT(func(v *ACIServiceCreateRequestEncryptionProperties) ACIServiceCreateRequestEncryptionProperties {
		if v != nil {
			return *v
		}
		var ret ACIServiceCreateRequestEncryptionProperties
		return ret
	}).(ACIServiceCreateRequestEncryptionPropertiesOutput)
}

func (o ACIServiceCreateRequestEncryptionPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIServiceCreateRequestEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o ACIServiceCreateRequestEncryptionPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIServiceCreateRequestEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o ACIServiceCreateRequestEncryptionPropertiesPtrOutput) VaultBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIServiceCreateRequestEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.VaultBaseUrl
	}).(pulumi.StringPtrOutput)
}

type ACIServiceCreateRequestVnetConfiguration struct {
	SubnetName *string `pulumi:"subnetName"`
	VnetName   *string `pulumi:"vnetName"`
}





type ACIServiceCreateRequestVnetConfigurationInput interface {
	pulumi.Input

	ToACIServiceCreateRequestVnetConfigurationOutput() ACIServiceCreateRequestVnetConfigurationOutput
	ToACIServiceCreateRequestVnetConfigurationOutputWithContext(context.Context) ACIServiceCreateRequestVnetConfigurationOutput
}

type ACIServiceCreateRequestVnetConfigurationArgs struct {
	SubnetName pulumi.StringPtrInput `pulumi:"subnetName"`
	VnetName   pulumi.StringPtrInput `pulumi:"vnetName"`
}

func (ACIServiceCreateRequestVnetConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceCreateRequestVnetConfiguration)(nil)).Elem()
}

func (i ACIServiceCreateRequestVnetConfigurationArgs) ToACIServiceCreateRequestVnetConfigurationOutput() ACIServiceCreateRequestVnetConfigurationOutput {
	return i.ToACIServiceCreateRequestVnetConfigurationOutputWithContext(context.Background())
}

func (i ACIServiceCreateRequestVnetConfigurationArgs) ToACIServiceCreateRequestVnetConfigurationOutputWithContext(ctx context.Context) ACIServiceCreateRequestVnetConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceCreateRequestVnetConfigurationOutput)
}

func (i ACIServiceCreateRequestVnetConfigurationArgs) ToACIServiceCreateRequestVnetConfigurationPtrOutput() ACIServiceCreateRequestVnetConfigurationPtrOutput {
	return i.ToACIServiceCreateRequestVnetConfigurationPtrOutputWithContext(context.Background())
}

func (i ACIServiceCreateRequestVnetConfigurationArgs) ToACIServiceCreateRequestVnetConfigurationPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestVnetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceCreateRequestVnetConfigurationOutput).ToACIServiceCreateRequestVnetConfigurationPtrOutputWithContext(ctx)
}









type ACIServiceCreateRequestVnetConfigurationPtrInput interface {
	pulumi.Input

	ToACIServiceCreateRequestVnetConfigurationPtrOutput() ACIServiceCreateRequestVnetConfigurationPtrOutput
	ToACIServiceCreateRequestVnetConfigurationPtrOutputWithContext(context.Context) ACIServiceCreateRequestVnetConfigurationPtrOutput
}

type aciserviceCreateRequestVnetConfigurationPtrType ACIServiceCreateRequestVnetConfigurationArgs

func ACIServiceCreateRequestVnetConfigurationPtr(v *ACIServiceCreateRequestVnetConfigurationArgs) ACIServiceCreateRequestVnetConfigurationPtrInput {
	return (*aciserviceCreateRequestVnetConfigurationPtrType)(v)
}

func (*aciserviceCreateRequestVnetConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceCreateRequestVnetConfiguration)(nil)).Elem()
}

func (i *aciserviceCreateRequestVnetConfigurationPtrType) ToACIServiceCreateRequestVnetConfigurationPtrOutput() ACIServiceCreateRequestVnetConfigurationPtrOutput {
	return i.ToACIServiceCreateRequestVnetConfigurationPtrOutputWithContext(context.Background())
}

func (i *aciserviceCreateRequestVnetConfigurationPtrType) ToACIServiceCreateRequestVnetConfigurationPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestVnetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceCreateRequestVnetConfigurationPtrOutput)
}

type ACIServiceCreateRequestVnetConfigurationOutput struct{ *pulumi.OutputState }

func (ACIServiceCreateRequestVnetConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceCreateRequestVnetConfiguration)(nil)).Elem()
}

func (o ACIServiceCreateRequestVnetConfigurationOutput) ToACIServiceCreateRequestVnetConfigurationOutput() ACIServiceCreateRequestVnetConfigurationOutput {
	return o
}

func (o ACIServiceCreateRequestVnetConfigurationOutput) ToACIServiceCreateRequestVnetConfigurationOutputWithContext(ctx context.Context) ACIServiceCreateRequestVnetConfigurationOutput {
	return o
}

func (o ACIServiceCreateRequestVnetConfigurationOutput) ToACIServiceCreateRequestVnetConfigurationPtrOutput() ACIServiceCreateRequestVnetConfigurationPtrOutput {
	return o.ToACIServiceCreateRequestVnetConfigurationPtrOutputWithContext(context.Background())
}

func (o ACIServiceCreateRequestVnetConfigurationOutput) ToACIServiceCreateRequestVnetConfigurationPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestVnetConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ACIServiceCreateRequestVnetConfiguration) *ACIServiceCreateRequestVnetConfiguration {
		return &v
	}).(ACIServiceCreateRequestVnetConfigurationPtrOutput)
}

func (o ACIServiceCreateRequestVnetConfigurationOutput) SubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceCreateRequestVnetConfiguration) *string { return v.SubnetName }).(pulumi.StringPtrOutput)
}

func (o ACIServiceCreateRequestVnetConfigurationOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceCreateRequestVnetConfiguration) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

type ACIServiceCreateRequestVnetConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ACIServiceCreateRequestVnetConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceCreateRequestVnetConfiguration)(nil)).Elem()
}

func (o ACIServiceCreateRequestVnetConfigurationPtrOutput) ToACIServiceCreateRequestVnetConfigurationPtrOutput() ACIServiceCreateRequestVnetConfigurationPtrOutput {
	return o
}

func (o ACIServiceCreateRequestVnetConfigurationPtrOutput) ToACIServiceCreateRequestVnetConfigurationPtrOutputWithContext(ctx context.Context) ACIServiceCreateRequestVnetConfigurationPtrOutput {
	return o
}

func (o ACIServiceCreateRequestVnetConfigurationPtrOutput) Elem() ACIServiceCreateRequestVnetConfigurationOutput {
	return o.ApplyT(func(v *ACIServiceCreateRequestVnetConfiguration) ACIServiceCreateRequestVnetConfiguration {
		if v != nil {
			return *v
		}
		var ret ACIServiceCreateRequestVnetConfiguration
		return ret
	}).(ACIServiceCreateRequestVnetConfigurationOutput)
}

func (o ACIServiceCreateRequestVnetConfigurationPtrOutput) SubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIServiceCreateRequestVnetConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SubnetName
	}).(pulumi.StringPtrOutput)
}

func (o ACIServiceCreateRequestVnetConfigurationPtrOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIServiceCreateRequestVnetConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.VnetName
	}).(pulumi.StringPtrOutput)
}

type ACIServiceResponseResponse struct {
	AppInsightsEnabled            *bool                                              `pulumi:"appInsightsEnabled"`
	AuthEnabled                   *bool                                              `pulumi:"authEnabled"`
	Cname                         *string                                            `pulumi:"cname"`
	ComputeType                   string                                             `pulumi:"computeType"`
	ContainerResourceRequirements *ContainerResourceRequirementsResponse             `pulumi:"containerResourceRequirements"`
	DataCollection                *ACIServiceResponseResponseDataCollection          `pulumi:"dataCollection"`
	DeploymentType                *string                                            `pulumi:"deploymentType"`
	Description                   *string                                            `pulumi:"description"`
	EncryptionProperties          *ACIServiceResponseResponseEncryptionProperties    `pulumi:"encryptionProperties"`
	EnvironmentImageRequest       *ACIServiceResponseResponseEnvironmentImageRequest `pulumi:"environmentImageRequest"`
	Error                         ServiceResponseBaseResponseError                   `pulumi:"error"`
	KvTags                        map[string]string                                  `pulumi:"kvTags"`
	Location                      *string                                            `pulumi:"location"`
	ModelConfigMap                map[string]interface{}                             `pulumi:"modelConfigMap"`
	Models                        []ModelResponse                                    `pulumi:"models"`
	Properties                    map[string]string                                  `pulumi:"properties"`
	PublicFqdn                    *string                                            `pulumi:"publicFqdn"`
	PublicIp                      *string                                            `pulumi:"publicIp"`
	ScoringUri                    string                                             `pulumi:"scoringUri"`
	SslCertificate                *string                                            `pulumi:"sslCertificate"`
	SslEnabled                    *bool                                              `pulumi:"sslEnabled"`
	SslKey                        *string                                            `pulumi:"sslKey"`
	State                         string                                             `pulumi:"state"`
	SwaggerUri                    string                                             `pulumi:"swaggerUri"`
	VnetConfiguration             *ACIServiceResponseResponseVnetConfiguration       `pulumi:"vnetConfiguration"`
}

type ACIServiceResponseResponseDataCollection struct {
	EventHubEnabled *bool `pulumi:"eventHubEnabled"`
	StorageEnabled  *bool `pulumi:"storageEnabled"`
}

type ACIServiceResponseResponseEncryptionProperties struct {
	KeyName      string `pulumi:"keyName"`
	KeyVersion   string `pulumi:"keyVersion"`
	VaultBaseUrl string `pulumi:"vaultBaseUrl"`
}

type ACIServiceResponseResponseEnvironmentImageRequest struct {
	Assets               []ImageAssetResponse                                  `pulumi:"assets"`
	DriverProgram        *string                                               `pulumi:"driverProgram"`
	Environment          *EnvironmentImageResponseResponseEnvironment          `pulumi:"environment"`
	EnvironmentReference *EnvironmentImageResponseResponseEnvironmentReference `pulumi:"environmentReference"`
	ModelIds             []string                                              `pulumi:"modelIds"`
	Models               []ModelResponse                                       `pulumi:"models"`
}

type ACIServiceResponseResponseVnetConfiguration struct {
	SubnetName *string `pulumi:"subnetName"`
	VnetName   *string `pulumi:"vnetName"`
}

type AKS struct {
	ComputeLocation *string        `pulumi:"computeLocation"`
	ComputeType     string         `pulumi:"computeType"`
	Description     *string        `pulumi:"description"`
	Properties      *AKSProperties `pulumi:"properties"`
	ResourceId      *string        `pulumi:"resourceId"`
}

type AKSProperties struct {
	AgentCount                 *int                        `pulumi:"agentCount"`
	AgentVmSize                *string                     `pulumi:"agentVmSize"`
	AksNetworkingConfiguration *AksNetworkingConfiguration `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                *string                     `pulumi:"clusterFqdn"`
	SslConfiguration           *SslConfiguration           `pulumi:"sslConfiguration"`
}

type AKSReplicaStatusResponseError struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
}

type AKSResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *AKSResponseProperties                `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type AKSResponseProperties struct {
	AgentCount                 *int                                `pulumi:"agentCount"`
	AgentVmSize                *string                             `pulumi:"agentVmSize"`
	AksNetworkingConfiguration *AksNetworkingConfigurationResponse `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                *string                             `pulumi:"clusterFqdn"`
	SslConfiguration           *SslConfigurationResponse           `pulumi:"sslConfiguration"`
	SystemServices             []SystemServiceResponse             `pulumi:"systemServices"`
}

type AKSServiceCreateRequestAutoScaler struct {
	AutoscaleEnabled       *bool `pulumi:"autoscaleEnabled"`
	MaxReplicas            *int  `pulumi:"maxReplicas"`
	MinReplicas            *int  `pulumi:"minReplicas"`
	RefreshPeriodInSeconds *int  `pulumi:"refreshPeriodInSeconds"`
	TargetUtilization      *int  `pulumi:"targetUtilization"`
}





type AKSServiceCreateRequestAutoScalerInput interface {
	pulumi.Input

	ToAKSServiceCreateRequestAutoScalerOutput() AKSServiceCreateRequestAutoScalerOutput
	ToAKSServiceCreateRequestAutoScalerOutputWithContext(context.Context) AKSServiceCreateRequestAutoScalerOutput
}

type AKSServiceCreateRequestAutoScalerArgs struct {
	AutoscaleEnabled       pulumi.BoolPtrInput `pulumi:"autoscaleEnabled"`
	MaxReplicas            pulumi.IntPtrInput  `pulumi:"maxReplicas"`
	MinReplicas            pulumi.IntPtrInput  `pulumi:"minReplicas"`
	RefreshPeriodInSeconds pulumi.IntPtrInput  `pulumi:"refreshPeriodInSeconds"`
	TargetUtilization      pulumi.IntPtrInput  `pulumi:"targetUtilization"`
}

func (AKSServiceCreateRequestAutoScalerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceCreateRequestAutoScaler)(nil)).Elem()
}

func (i AKSServiceCreateRequestAutoScalerArgs) ToAKSServiceCreateRequestAutoScalerOutput() AKSServiceCreateRequestAutoScalerOutput {
	return i.ToAKSServiceCreateRequestAutoScalerOutputWithContext(context.Background())
}

func (i AKSServiceCreateRequestAutoScalerArgs) ToAKSServiceCreateRequestAutoScalerOutputWithContext(ctx context.Context) AKSServiceCreateRequestAutoScalerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceCreateRequestAutoScalerOutput)
}

func (i AKSServiceCreateRequestAutoScalerArgs) ToAKSServiceCreateRequestAutoScalerPtrOutput() AKSServiceCreateRequestAutoScalerPtrOutput {
	return i.ToAKSServiceCreateRequestAutoScalerPtrOutputWithContext(context.Background())
}

func (i AKSServiceCreateRequestAutoScalerArgs) ToAKSServiceCreateRequestAutoScalerPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestAutoScalerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceCreateRequestAutoScalerOutput).ToAKSServiceCreateRequestAutoScalerPtrOutputWithContext(ctx)
}









type AKSServiceCreateRequestAutoScalerPtrInput interface {
	pulumi.Input

	ToAKSServiceCreateRequestAutoScalerPtrOutput() AKSServiceCreateRequestAutoScalerPtrOutput
	ToAKSServiceCreateRequestAutoScalerPtrOutputWithContext(context.Context) AKSServiceCreateRequestAutoScalerPtrOutput
}

type aksserviceCreateRequestAutoScalerPtrType AKSServiceCreateRequestAutoScalerArgs

func AKSServiceCreateRequestAutoScalerPtr(v *AKSServiceCreateRequestAutoScalerArgs) AKSServiceCreateRequestAutoScalerPtrInput {
	return (*aksserviceCreateRequestAutoScalerPtrType)(v)
}

func (*aksserviceCreateRequestAutoScalerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceCreateRequestAutoScaler)(nil)).Elem()
}

func (i *aksserviceCreateRequestAutoScalerPtrType) ToAKSServiceCreateRequestAutoScalerPtrOutput() AKSServiceCreateRequestAutoScalerPtrOutput {
	return i.ToAKSServiceCreateRequestAutoScalerPtrOutputWithContext(context.Background())
}

func (i *aksserviceCreateRequestAutoScalerPtrType) ToAKSServiceCreateRequestAutoScalerPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestAutoScalerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceCreateRequestAutoScalerPtrOutput)
}

type AKSServiceCreateRequestAutoScalerOutput struct{ *pulumi.OutputState }

func (AKSServiceCreateRequestAutoScalerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceCreateRequestAutoScaler)(nil)).Elem()
}

func (o AKSServiceCreateRequestAutoScalerOutput) ToAKSServiceCreateRequestAutoScalerOutput() AKSServiceCreateRequestAutoScalerOutput {
	return o
}

func (o AKSServiceCreateRequestAutoScalerOutput) ToAKSServiceCreateRequestAutoScalerOutputWithContext(ctx context.Context) AKSServiceCreateRequestAutoScalerOutput {
	return o
}

func (o AKSServiceCreateRequestAutoScalerOutput) ToAKSServiceCreateRequestAutoScalerPtrOutput() AKSServiceCreateRequestAutoScalerPtrOutput {
	return o.ToAKSServiceCreateRequestAutoScalerPtrOutputWithContext(context.Background())
}

func (o AKSServiceCreateRequestAutoScalerOutput) ToAKSServiceCreateRequestAutoScalerPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestAutoScalerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSServiceCreateRequestAutoScaler) *AKSServiceCreateRequestAutoScaler {
		return &v
	}).(AKSServiceCreateRequestAutoScalerPtrOutput)
}

func (o AKSServiceCreateRequestAutoScalerOutput) AutoscaleEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestAutoScaler) *bool { return v.AutoscaleEnabled }).(pulumi.BoolPtrOutput)
}

func (o AKSServiceCreateRequestAutoScalerOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestAutoScaler) *int { return v.MaxReplicas }).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestAutoScalerOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestAutoScaler) *int { return v.MinReplicas }).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestAutoScalerOutput) RefreshPeriodInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestAutoScaler) *int { return v.RefreshPeriodInSeconds }).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestAutoScalerOutput) TargetUtilization() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestAutoScaler) *int { return v.TargetUtilization }).(pulumi.IntPtrOutput)
}

type AKSServiceCreateRequestAutoScalerPtrOutput struct{ *pulumi.OutputState }

func (AKSServiceCreateRequestAutoScalerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceCreateRequestAutoScaler)(nil)).Elem()
}

func (o AKSServiceCreateRequestAutoScalerPtrOutput) ToAKSServiceCreateRequestAutoScalerPtrOutput() AKSServiceCreateRequestAutoScalerPtrOutput {
	return o
}

func (o AKSServiceCreateRequestAutoScalerPtrOutput) ToAKSServiceCreateRequestAutoScalerPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestAutoScalerPtrOutput {
	return o
}

func (o AKSServiceCreateRequestAutoScalerPtrOutput) Elem() AKSServiceCreateRequestAutoScalerOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestAutoScaler) AKSServiceCreateRequestAutoScaler {
		if v != nil {
			return *v
		}
		var ret AKSServiceCreateRequestAutoScaler
		return ret
	}).(AKSServiceCreateRequestAutoScalerOutput)
}

func (o AKSServiceCreateRequestAutoScalerPtrOutput) AutoscaleEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestAutoScaler) *bool {
		if v == nil {
			return nil
		}
		return v.AutoscaleEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AKSServiceCreateRequestAutoScalerPtrOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestAutoScaler) *int {
		if v == nil {
			return nil
		}
		return v.MaxReplicas
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestAutoScalerPtrOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestAutoScaler) *int {
		if v == nil {
			return nil
		}
		return v.MinReplicas
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestAutoScalerPtrOutput) RefreshPeriodInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestAutoScaler) *int {
		if v == nil {
			return nil
		}
		return v.RefreshPeriodInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestAutoScalerPtrOutput) TargetUtilization() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestAutoScaler) *int {
		if v == nil {
			return nil
		}
		return v.TargetUtilization
	}).(pulumi.IntPtrOutput)
}

type AKSServiceCreateRequestDataCollection struct {
	EventHubEnabled *bool `pulumi:"eventHubEnabled"`
	StorageEnabled  *bool `pulumi:"storageEnabled"`
}





type AKSServiceCreateRequestDataCollectionInput interface {
	pulumi.Input

	ToAKSServiceCreateRequestDataCollectionOutput() AKSServiceCreateRequestDataCollectionOutput
	ToAKSServiceCreateRequestDataCollectionOutputWithContext(context.Context) AKSServiceCreateRequestDataCollectionOutput
}

type AKSServiceCreateRequestDataCollectionArgs struct {
	EventHubEnabled pulumi.BoolPtrInput `pulumi:"eventHubEnabled"`
	StorageEnabled  pulumi.BoolPtrInput `pulumi:"storageEnabled"`
}

func (AKSServiceCreateRequestDataCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceCreateRequestDataCollection)(nil)).Elem()
}

func (i AKSServiceCreateRequestDataCollectionArgs) ToAKSServiceCreateRequestDataCollectionOutput() AKSServiceCreateRequestDataCollectionOutput {
	return i.ToAKSServiceCreateRequestDataCollectionOutputWithContext(context.Background())
}

func (i AKSServiceCreateRequestDataCollectionArgs) ToAKSServiceCreateRequestDataCollectionOutputWithContext(ctx context.Context) AKSServiceCreateRequestDataCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceCreateRequestDataCollectionOutput)
}

func (i AKSServiceCreateRequestDataCollectionArgs) ToAKSServiceCreateRequestDataCollectionPtrOutput() AKSServiceCreateRequestDataCollectionPtrOutput {
	return i.ToAKSServiceCreateRequestDataCollectionPtrOutputWithContext(context.Background())
}

func (i AKSServiceCreateRequestDataCollectionArgs) ToAKSServiceCreateRequestDataCollectionPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestDataCollectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceCreateRequestDataCollectionOutput).ToAKSServiceCreateRequestDataCollectionPtrOutputWithContext(ctx)
}









type AKSServiceCreateRequestDataCollectionPtrInput interface {
	pulumi.Input

	ToAKSServiceCreateRequestDataCollectionPtrOutput() AKSServiceCreateRequestDataCollectionPtrOutput
	ToAKSServiceCreateRequestDataCollectionPtrOutputWithContext(context.Context) AKSServiceCreateRequestDataCollectionPtrOutput
}

type aksserviceCreateRequestDataCollectionPtrType AKSServiceCreateRequestDataCollectionArgs

func AKSServiceCreateRequestDataCollectionPtr(v *AKSServiceCreateRequestDataCollectionArgs) AKSServiceCreateRequestDataCollectionPtrInput {
	return (*aksserviceCreateRequestDataCollectionPtrType)(v)
}

func (*aksserviceCreateRequestDataCollectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceCreateRequestDataCollection)(nil)).Elem()
}

func (i *aksserviceCreateRequestDataCollectionPtrType) ToAKSServiceCreateRequestDataCollectionPtrOutput() AKSServiceCreateRequestDataCollectionPtrOutput {
	return i.ToAKSServiceCreateRequestDataCollectionPtrOutputWithContext(context.Background())
}

func (i *aksserviceCreateRequestDataCollectionPtrType) ToAKSServiceCreateRequestDataCollectionPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestDataCollectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceCreateRequestDataCollectionPtrOutput)
}

type AKSServiceCreateRequestDataCollectionOutput struct{ *pulumi.OutputState }

func (AKSServiceCreateRequestDataCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceCreateRequestDataCollection)(nil)).Elem()
}

func (o AKSServiceCreateRequestDataCollectionOutput) ToAKSServiceCreateRequestDataCollectionOutput() AKSServiceCreateRequestDataCollectionOutput {
	return o
}

func (o AKSServiceCreateRequestDataCollectionOutput) ToAKSServiceCreateRequestDataCollectionOutputWithContext(ctx context.Context) AKSServiceCreateRequestDataCollectionOutput {
	return o
}

func (o AKSServiceCreateRequestDataCollectionOutput) ToAKSServiceCreateRequestDataCollectionPtrOutput() AKSServiceCreateRequestDataCollectionPtrOutput {
	return o.ToAKSServiceCreateRequestDataCollectionPtrOutputWithContext(context.Background())
}

func (o AKSServiceCreateRequestDataCollectionOutput) ToAKSServiceCreateRequestDataCollectionPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestDataCollectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSServiceCreateRequestDataCollection) *AKSServiceCreateRequestDataCollection {
		return &v
	}).(AKSServiceCreateRequestDataCollectionPtrOutput)
}

func (o AKSServiceCreateRequestDataCollectionOutput) EventHubEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestDataCollection) *bool { return v.EventHubEnabled }).(pulumi.BoolPtrOutput)
}

func (o AKSServiceCreateRequestDataCollectionOutput) StorageEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestDataCollection) *bool { return v.StorageEnabled }).(pulumi.BoolPtrOutput)
}

type AKSServiceCreateRequestDataCollectionPtrOutput struct{ *pulumi.OutputState }

func (AKSServiceCreateRequestDataCollectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceCreateRequestDataCollection)(nil)).Elem()
}

func (o AKSServiceCreateRequestDataCollectionPtrOutput) ToAKSServiceCreateRequestDataCollectionPtrOutput() AKSServiceCreateRequestDataCollectionPtrOutput {
	return o
}

func (o AKSServiceCreateRequestDataCollectionPtrOutput) ToAKSServiceCreateRequestDataCollectionPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestDataCollectionPtrOutput {
	return o
}

func (o AKSServiceCreateRequestDataCollectionPtrOutput) Elem() AKSServiceCreateRequestDataCollectionOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestDataCollection) AKSServiceCreateRequestDataCollection {
		if v != nil {
			return *v
		}
		var ret AKSServiceCreateRequestDataCollection
		return ret
	}).(AKSServiceCreateRequestDataCollectionOutput)
}

func (o AKSServiceCreateRequestDataCollectionPtrOutput) EventHubEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestDataCollection) *bool {
		if v == nil {
			return nil
		}
		return v.EventHubEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AKSServiceCreateRequestDataCollectionPtrOutput) StorageEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestDataCollection) *bool {
		if v == nil {
			return nil
		}
		return v.StorageEnabled
	}).(pulumi.BoolPtrOutput)
}

type AKSServiceCreateRequestLivenessProbeRequirements struct {
	FailureThreshold    *int `pulumi:"failureThreshold"`
	InitialDelaySeconds *int `pulumi:"initialDelaySeconds"`
	PeriodSeconds       *int `pulumi:"periodSeconds"`
	SuccessThreshold    *int `pulumi:"successThreshold"`
	TimeoutSeconds      *int `pulumi:"timeoutSeconds"`
}





type AKSServiceCreateRequestLivenessProbeRequirementsInput interface {
	pulumi.Input

	ToAKSServiceCreateRequestLivenessProbeRequirementsOutput() AKSServiceCreateRequestLivenessProbeRequirementsOutput
	ToAKSServiceCreateRequestLivenessProbeRequirementsOutputWithContext(context.Context) AKSServiceCreateRequestLivenessProbeRequirementsOutput
}

type AKSServiceCreateRequestLivenessProbeRequirementsArgs struct {
	FailureThreshold    pulumi.IntPtrInput `pulumi:"failureThreshold"`
	InitialDelaySeconds pulumi.IntPtrInput `pulumi:"initialDelaySeconds"`
	PeriodSeconds       pulumi.IntPtrInput `pulumi:"periodSeconds"`
	SuccessThreshold    pulumi.IntPtrInput `pulumi:"successThreshold"`
	TimeoutSeconds      pulumi.IntPtrInput `pulumi:"timeoutSeconds"`
}

func (AKSServiceCreateRequestLivenessProbeRequirementsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceCreateRequestLivenessProbeRequirements)(nil)).Elem()
}

func (i AKSServiceCreateRequestLivenessProbeRequirementsArgs) ToAKSServiceCreateRequestLivenessProbeRequirementsOutput() AKSServiceCreateRequestLivenessProbeRequirementsOutput {
	return i.ToAKSServiceCreateRequestLivenessProbeRequirementsOutputWithContext(context.Background())
}

func (i AKSServiceCreateRequestLivenessProbeRequirementsArgs) ToAKSServiceCreateRequestLivenessProbeRequirementsOutputWithContext(ctx context.Context) AKSServiceCreateRequestLivenessProbeRequirementsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceCreateRequestLivenessProbeRequirementsOutput)
}

func (i AKSServiceCreateRequestLivenessProbeRequirementsArgs) ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutput() AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput {
	return i.ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutputWithContext(context.Background())
}

func (i AKSServiceCreateRequestLivenessProbeRequirementsArgs) ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceCreateRequestLivenessProbeRequirementsOutput).ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutputWithContext(ctx)
}









type AKSServiceCreateRequestLivenessProbeRequirementsPtrInput interface {
	pulumi.Input

	ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutput() AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput
	ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutputWithContext(context.Context) AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput
}

type aksserviceCreateRequestLivenessProbeRequirementsPtrType AKSServiceCreateRequestLivenessProbeRequirementsArgs

func AKSServiceCreateRequestLivenessProbeRequirementsPtr(v *AKSServiceCreateRequestLivenessProbeRequirementsArgs) AKSServiceCreateRequestLivenessProbeRequirementsPtrInput {
	return (*aksserviceCreateRequestLivenessProbeRequirementsPtrType)(v)
}

func (*aksserviceCreateRequestLivenessProbeRequirementsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceCreateRequestLivenessProbeRequirements)(nil)).Elem()
}

func (i *aksserviceCreateRequestLivenessProbeRequirementsPtrType) ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutput() AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput {
	return i.ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutputWithContext(context.Background())
}

func (i *aksserviceCreateRequestLivenessProbeRequirementsPtrType) ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput)
}

type AKSServiceCreateRequestLivenessProbeRequirementsOutput struct{ *pulumi.OutputState }

func (AKSServiceCreateRequestLivenessProbeRequirementsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceCreateRequestLivenessProbeRequirements)(nil)).Elem()
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsOutput) ToAKSServiceCreateRequestLivenessProbeRequirementsOutput() AKSServiceCreateRequestLivenessProbeRequirementsOutput {
	return o
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsOutput) ToAKSServiceCreateRequestLivenessProbeRequirementsOutputWithContext(ctx context.Context) AKSServiceCreateRequestLivenessProbeRequirementsOutput {
	return o
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsOutput) ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutput() AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput {
	return o.ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutputWithContext(context.Background())
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsOutput) ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSServiceCreateRequestLivenessProbeRequirements) *AKSServiceCreateRequestLivenessProbeRequirements {
		return &v
	}).(AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput)
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestLivenessProbeRequirements) *int { return v.FailureThreshold }).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsOutput) InitialDelaySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestLivenessProbeRequirements) *int { return v.InitialDelaySeconds }).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsOutput) PeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestLivenessProbeRequirements) *int { return v.PeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsOutput) SuccessThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestLivenessProbeRequirements) *int { return v.SuccessThreshold }).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceCreateRequestLivenessProbeRequirements) *int { return v.TimeoutSeconds }).(pulumi.IntPtrOutput)
}

type AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput struct{ *pulumi.OutputState }

func (AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceCreateRequestLivenessProbeRequirements)(nil)).Elem()
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput) ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutput() AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput {
	return o
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput) ToAKSServiceCreateRequestLivenessProbeRequirementsPtrOutputWithContext(ctx context.Context) AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput {
	return o
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput) Elem() AKSServiceCreateRequestLivenessProbeRequirementsOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestLivenessProbeRequirements) AKSServiceCreateRequestLivenessProbeRequirements {
		if v != nil {
			return *v
		}
		var ret AKSServiceCreateRequestLivenessProbeRequirements
		return ret
	}).(AKSServiceCreateRequestLivenessProbeRequirementsOutput)
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestLivenessProbeRequirements) *int {
		if v == nil {
			return nil
		}
		return v.FailureThreshold
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput) InitialDelaySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestLivenessProbeRequirements) *int {
		if v == nil {
			return nil
		}
		return v.InitialDelaySeconds
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput) PeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestLivenessProbeRequirements) *int {
		if v == nil {
			return nil
		}
		return v.PeriodSeconds
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput) SuccessThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestLivenessProbeRequirements) *int {
		if v == nil {
			return nil
		}
		return v.SuccessThreshold
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceCreateRequestLivenessProbeRequirements) *int {
		if v == nil {
			return nil
		}
		return v.TimeoutSeconds
	}).(pulumi.IntPtrOutput)
}

type AKSServiceResponseResponse struct {
	AadAuthEnabled                    *bool                                                `pulumi:"aadAuthEnabled"`
	AppInsightsEnabled                *bool                                                `pulumi:"appInsightsEnabled"`
	AuthEnabled                       *bool                                                `pulumi:"authEnabled"`
	AutoScaler                        *AKSServiceResponseResponseAutoScaler                `pulumi:"autoScaler"`
	ComputeName                       *string                                              `pulumi:"computeName"`
	ComputeType                       string                                               `pulumi:"computeType"`
	ContainerResourceRequirements     *ContainerResourceRequirementsResponse               `pulumi:"containerResourceRequirements"`
	DataCollection                    *AKSServiceResponseResponseDataCollection            `pulumi:"dataCollection"`
	DeploymentStatus                  AKSServiceResponseResponseDeploymentStatus           `pulumi:"deploymentStatus"`
	DeploymentType                    *string                                              `pulumi:"deploymentType"`
	Description                       *string                                              `pulumi:"description"`
	EnvironmentImageRequest           *AKSServiceResponseResponseEnvironmentImageRequest   `pulumi:"environmentImageRequest"`
	Error                             ServiceResponseBaseResponseError                     `pulumi:"error"`
	IsDefault                         *bool                                                `pulumi:"isDefault"`
	KvTags                            map[string]string                                    `pulumi:"kvTags"`
	LivenessProbeRequirements         *AKSServiceResponseResponseLivenessProbeRequirements `pulumi:"livenessProbeRequirements"`
	MaxConcurrentRequestsPerContainer *int                                                 `pulumi:"maxConcurrentRequestsPerContainer"`
	MaxQueueWaitMs                    *int                                                 `pulumi:"maxQueueWaitMs"`
	ModelConfigMap                    map[string]interface{}                               `pulumi:"modelConfigMap"`
	Models                            []ModelResponse                                      `pulumi:"models"`
	Namespace                         *string                                              `pulumi:"namespace"`
	NumReplicas                       *int                                                 `pulumi:"numReplicas"`
	Properties                        map[string]string                                    `pulumi:"properties"`
	ScoringTimeoutMs                  *int                                                 `pulumi:"scoringTimeoutMs"`
	ScoringUri                        string                                               `pulumi:"scoringUri"`
	State                             string                                               `pulumi:"state"`
	SwaggerUri                        string                                               `pulumi:"swaggerUri"`
	TrafficPercentile                 *float64                                             `pulumi:"trafficPercentile"`
	Type                              *string                                              `pulumi:"type"`
}

type AKSServiceResponseResponseAutoScaler struct {
	AutoscaleEnabled       *bool `pulumi:"autoscaleEnabled"`
	MaxReplicas            *int  `pulumi:"maxReplicas"`
	MinReplicas            *int  `pulumi:"minReplicas"`
	RefreshPeriodInSeconds *int  `pulumi:"refreshPeriodInSeconds"`
	TargetUtilization      *int  `pulumi:"targetUtilization"`
}

type AKSServiceResponseResponseDataCollection struct {
	EventHubEnabled *bool `pulumi:"eventHubEnabled"`
	StorageEnabled  *bool `pulumi:"storageEnabled"`
}

type AKSServiceResponseResponseDeploymentStatus struct {
	AvailableReplicas *int                           `pulumi:"availableReplicas"`
	DesiredReplicas   *int                           `pulumi:"desiredReplicas"`
	Error             *AKSReplicaStatusResponseError `pulumi:"error"`
	UpdatedReplicas   *int                           `pulumi:"updatedReplicas"`
}

type AKSServiceResponseResponseEnvironmentImageRequest struct {
	Assets               []ImageAssetResponse                                  `pulumi:"assets"`
	DriverProgram        *string                                               `pulumi:"driverProgram"`
	Environment          *EnvironmentImageResponseResponseEnvironment          `pulumi:"environment"`
	EnvironmentReference *EnvironmentImageResponseResponseEnvironmentReference `pulumi:"environmentReference"`
	ModelIds             []string                                              `pulumi:"modelIds"`
	Models               []ModelResponse                                       `pulumi:"models"`
}

type AKSServiceResponseResponseLivenessProbeRequirements struct {
	FailureThreshold    *int `pulumi:"failureThreshold"`
	InitialDelaySeconds *int `pulumi:"initialDelaySeconds"`
	PeriodSeconds       *int `pulumi:"periodSeconds"`
	SuccessThreshold    *int `pulumi:"successThreshold"`
	TimeoutSeconds      *int `pulumi:"timeoutSeconds"`
}

type AKSVariantResponseResponse struct {
	ComputeType       string                           `pulumi:"computeType"`
	DeploymentType    *string                          `pulumi:"deploymentType"`
	Description       *string                          `pulumi:"description"`
	Error             ServiceResponseBaseResponseError `pulumi:"error"`
	IsDefault         *bool                            `pulumi:"isDefault"`
	KvTags            map[string]string                `pulumi:"kvTags"`
	Properties        map[string]string                `pulumi:"properties"`
	State             string                           `pulumi:"state"`
	TrafficPercentile *float64                         `pulumi:"trafficPercentile"`
	Type              *string                          `pulumi:"type"`
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
	ComputeLocation *string               `pulumi:"computeLocation"`
	ComputeType     string                `pulumi:"computeType"`
	Description     *string               `pulumi:"description"`
	Properties      *AmlComputeProperties `pulumi:"properties"`
	ResourceId      *string               `pulumi:"resourceId"`
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
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *AmlComputeResponseProperties         `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
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
	AllocationState               string                                `pulumi:"allocationState"`
	AllocationStateTransitionTime string                                `pulumi:"allocationStateTransitionTime"`
	CurrentNodeCount              int                                   `pulumi:"currentNodeCount"`
	EnableNodePublicIp            *bool                                 `pulumi:"enableNodePublicIp"`
	Errors                        []MachineLearningServiceErrorResponse `pulumi:"errors"`
	IsolatedNetwork               *bool                                 `pulumi:"isolatedNetwork"`
	NodeStateCounts               NodeStateCountsResponse               `pulumi:"nodeStateCounts"`
	OsType                        *string                               `pulumi:"osType"`
	RemoteLoginPortPublicAccess   *string                               `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings                 *ScaleSettingsResponse                `pulumi:"scaleSettings"`
	Subnet                        *ResourceIdResponse                   `pulumi:"subnet"`
	TargetNodeCount               int                                   `pulumi:"targetNodeCount"`
	UserAccountCredentials        *UserAccountCredentialsResponse       `pulumi:"userAccountCredentials"`
	VirtualMachineImage           *VirtualMachineImageResponse          `pulumi:"virtualMachineImage"`
	VmPriority                    *string                               `pulumi:"vmPriority"`
	VmSize                        *string                               `pulumi:"vmSize"`
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

type AssignedUser struct {
	ObjectId string `pulumi:"objectId"`
	TenantId string `pulumi:"tenantId"`
}

type AssignedUserResponse struct {
	ObjectId string `pulumi:"objectId"`
	TenantId string `pulumi:"tenantId"`
}

type ComputeBinding struct {
	ComputeId *string `pulumi:"computeId"`
	NodeCount *int    `pulumi:"nodeCount"`
}





type ComputeBindingInput interface {
	pulumi.Input

	ToComputeBindingOutput() ComputeBindingOutput
	ToComputeBindingOutputWithContext(context.Context) ComputeBindingOutput
}

type ComputeBindingArgs struct {
	ComputeId pulumi.StringPtrInput `pulumi:"computeId"`
	NodeCount pulumi.IntPtrInput    `pulumi:"nodeCount"`
}

func (ComputeBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeBinding)(nil)).Elem()
}

func (i ComputeBindingArgs) ToComputeBindingOutput() ComputeBindingOutput {
	return i.ToComputeBindingOutputWithContext(context.Background())
}

func (i ComputeBindingArgs) ToComputeBindingOutputWithContext(ctx context.Context) ComputeBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeBindingOutput)
}

func (i ComputeBindingArgs) ToComputeBindingPtrOutput() ComputeBindingPtrOutput {
	return i.ToComputeBindingPtrOutputWithContext(context.Background())
}

func (i ComputeBindingArgs) ToComputeBindingPtrOutputWithContext(ctx context.Context) ComputeBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeBindingOutput).ToComputeBindingPtrOutputWithContext(ctx)
}









type ComputeBindingPtrInput interface {
	pulumi.Input

	ToComputeBindingPtrOutput() ComputeBindingPtrOutput
	ToComputeBindingPtrOutputWithContext(context.Context) ComputeBindingPtrOutput
}

type computeBindingPtrType ComputeBindingArgs

func ComputeBindingPtr(v *ComputeBindingArgs) ComputeBindingPtrInput {
	return (*computeBindingPtrType)(v)
}

func (*computeBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeBinding)(nil)).Elem()
}

func (i *computeBindingPtrType) ToComputeBindingPtrOutput() ComputeBindingPtrOutput {
	return i.ToComputeBindingPtrOutputWithContext(context.Background())
}

func (i *computeBindingPtrType) ToComputeBindingPtrOutputWithContext(ctx context.Context) ComputeBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeBindingPtrOutput)
}

type ComputeBindingOutput struct{ *pulumi.OutputState }

func (ComputeBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeBinding)(nil)).Elem()
}

func (o ComputeBindingOutput) ToComputeBindingOutput() ComputeBindingOutput {
	return o
}

func (o ComputeBindingOutput) ToComputeBindingOutputWithContext(ctx context.Context) ComputeBindingOutput {
	return o
}

func (o ComputeBindingOutput) ToComputeBindingPtrOutput() ComputeBindingPtrOutput {
	return o.ToComputeBindingPtrOutputWithContext(context.Background())
}

func (o ComputeBindingOutput) ToComputeBindingPtrOutputWithContext(ctx context.Context) ComputeBindingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeBinding) *ComputeBinding {
		return &v
	}).(ComputeBindingPtrOutput)
}

func (o ComputeBindingOutput) ComputeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeBinding) *string { return v.ComputeId }).(pulumi.StringPtrOutput)
}

func (o ComputeBindingOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ComputeBinding) *int { return v.NodeCount }).(pulumi.IntPtrOutput)
}

type ComputeBindingPtrOutput struct{ *pulumi.OutputState }

func (ComputeBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeBinding)(nil)).Elem()
}

func (o ComputeBindingPtrOutput) ToComputeBindingPtrOutput() ComputeBindingPtrOutput {
	return o
}

func (o ComputeBindingPtrOutput) ToComputeBindingPtrOutputWithContext(ctx context.Context) ComputeBindingPtrOutput {
	return o
}

func (o ComputeBindingPtrOutput) Elem() ComputeBindingOutput {
	return o.ApplyT(func(v *ComputeBinding) ComputeBinding {
		if v != nil {
			return *v
		}
		var ret ComputeBinding
		return ret
	}).(ComputeBindingOutput)
}

func (o ComputeBindingPtrOutput) ComputeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeBinding) *string {
		if v == nil {
			return nil
		}
		return v.ComputeId
	}).(pulumi.StringPtrOutput)
}

func (o ComputeBindingPtrOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ComputeBinding) *int {
		if v == nil {
			return nil
		}
		return v.NodeCount
	}).(pulumi.IntPtrOutput)
}

type ComputeBindingResponse struct {
	ComputeId *string `pulumi:"computeId"`
	NodeCount *int    `pulumi:"nodeCount"`
}

type ComputeBindingResponseOutput struct{ *pulumi.OutputState }

func (ComputeBindingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeBindingResponse)(nil)).Elem()
}

func (o ComputeBindingResponseOutput) ToComputeBindingResponseOutput() ComputeBindingResponseOutput {
	return o
}

func (o ComputeBindingResponseOutput) ToComputeBindingResponseOutputWithContext(ctx context.Context) ComputeBindingResponseOutput {
	return o
}

func (o ComputeBindingResponseOutput) ComputeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeBindingResponse) *string { return v.ComputeId }).(pulumi.StringPtrOutput)
}

func (o ComputeBindingResponseOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ComputeBindingResponse) *int { return v.NodeCount }).(pulumi.IntPtrOutput)
}

type ComputeBindingResponsePtrOutput struct{ *pulumi.OutputState }

func (ComputeBindingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeBindingResponse)(nil)).Elem()
}

func (o ComputeBindingResponsePtrOutput) ToComputeBindingResponsePtrOutput() ComputeBindingResponsePtrOutput {
	return o
}

func (o ComputeBindingResponsePtrOutput) ToComputeBindingResponsePtrOutputWithContext(ctx context.Context) ComputeBindingResponsePtrOutput {
	return o
}

func (o ComputeBindingResponsePtrOutput) Elem() ComputeBindingResponseOutput {
	return o.ApplyT(func(v *ComputeBindingResponse) ComputeBindingResponse {
		if v != nil {
			return *v
		}
		var ret ComputeBindingResponse
		return ret
	}).(ComputeBindingResponseOutput)
}

func (o ComputeBindingResponsePtrOutput) ComputeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeBindingResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputeId
	}).(pulumi.StringPtrOutput)
}

func (o ComputeBindingResponsePtrOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ComputeBindingResponse) *int {
		if v == nil {
			return nil
		}
		return v.NodeCount
	}).(pulumi.IntPtrOutput)
}

type ComputeInstance struct {
	ComputeLocation *string                    `pulumi:"computeLocation"`
	ComputeType     string                     `pulumi:"computeType"`
	Description     *string                    `pulumi:"description"`
	Properties      *ComputeInstanceProperties `pulumi:"properties"`
	ResourceId      *string                    `pulumi:"resourceId"`
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

type ComputeInstanceResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *ComputeInstanceResponseProperties    `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
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
	Errors                           []MachineLearningServiceErrorResponse        `pulumi:"errors"`
	LastOperation                    ComputeInstanceLastOperationResponse         `pulumi:"lastOperation"`
	PersonalComputeInstanceSettings  *PersonalComputeInstanceSettingsResponse     `pulumi:"personalComputeInstanceSettings"`
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

type ContainerResourceRequirements struct {
	Cpu        *float64 `pulumi:"cpu"`
	Fpga       *int     `pulumi:"fpga"`
	Gpu        *int     `pulumi:"gpu"`
	MemoryInGB *float64 `pulumi:"memoryInGB"`
}





type ContainerResourceRequirementsInput interface {
	pulumi.Input

	ToContainerResourceRequirementsOutput() ContainerResourceRequirementsOutput
	ToContainerResourceRequirementsOutputWithContext(context.Context) ContainerResourceRequirementsOutput
}

type ContainerResourceRequirementsArgs struct {
	Cpu        pulumi.Float64PtrInput `pulumi:"cpu"`
	Fpga       pulumi.IntPtrInput     `pulumi:"fpga"`
	Gpu        pulumi.IntPtrInput     `pulumi:"gpu"`
	MemoryInGB pulumi.Float64PtrInput `pulumi:"memoryInGB"`
}

func (ContainerResourceRequirementsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResourceRequirements)(nil)).Elem()
}

func (i ContainerResourceRequirementsArgs) ToContainerResourceRequirementsOutput() ContainerResourceRequirementsOutput {
	return i.ToContainerResourceRequirementsOutputWithContext(context.Background())
}

func (i ContainerResourceRequirementsArgs) ToContainerResourceRequirementsOutputWithContext(ctx context.Context) ContainerResourceRequirementsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourceRequirementsOutput)
}

func (i ContainerResourceRequirementsArgs) ToContainerResourceRequirementsPtrOutput() ContainerResourceRequirementsPtrOutput {
	return i.ToContainerResourceRequirementsPtrOutputWithContext(context.Background())
}

func (i ContainerResourceRequirementsArgs) ToContainerResourceRequirementsPtrOutputWithContext(ctx context.Context) ContainerResourceRequirementsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourceRequirementsOutput).ToContainerResourceRequirementsPtrOutputWithContext(ctx)
}









type ContainerResourceRequirementsPtrInput interface {
	pulumi.Input

	ToContainerResourceRequirementsPtrOutput() ContainerResourceRequirementsPtrOutput
	ToContainerResourceRequirementsPtrOutputWithContext(context.Context) ContainerResourceRequirementsPtrOutput
}

type containerResourceRequirementsPtrType ContainerResourceRequirementsArgs

func ContainerResourceRequirementsPtr(v *ContainerResourceRequirementsArgs) ContainerResourceRequirementsPtrInput {
	return (*containerResourceRequirementsPtrType)(v)
}

func (*containerResourceRequirementsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerResourceRequirements)(nil)).Elem()
}

func (i *containerResourceRequirementsPtrType) ToContainerResourceRequirementsPtrOutput() ContainerResourceRequirementsPtrOutput {
	return i.ToContainerResourceRequirementsPtrOutputWithContext(context.Background())
}

func (i *containerResourceRequirementsPtrType) ToContainerResourceRequirementsPtrOutputWithContext(ctx context.Context) ContainerResourceRequirementsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourceRequirementsPtrOutput)
}

type ContainerResourceRequirementsOutput struct{ *pulumi.OutputState }

func (ContainerResourceRequirementsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResourceRequirements)(nil)).Elem()
}

func (o ContainerResourceRequirementsOutput) ToContainerResourceRequirementsOutput() ContainerResourceRequirementsOutput {
	return o
}

func (o ContainerResourceRequirementsOutput) ToContainerResourceRequirementsOutputWithContext(ctx context.Context) ContainerResourceRequirementsOutput {
	return o
}

func (o ContainerResourceRequirementsOutput) ToContainerResourceRequirementsPtrOutput() ContainerResourceRequirementsPtrOutput {
	return o.ToContainerResourceRequirementsPtrOutputWithContext(context.Background())
}

func (o ContainerResourceRequirementsOutput) ToContainerResourceRequirementsPtrOutputWithContext(ctx context.Context) ContainerResourceRequirementsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerResourceRequirements) *ContainerResourceRequirements {
		return &v
	}).(ContainerResourceRequirementsPtrOutput)
}

func (o ContainerResourceRequirementsOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerResourceRequirements) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ContainerResourceRequirementsOutput) Fpga() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerResourceRequirements) *int { return v.Fpga }).(pulumi.IntPtrOutput)
}

func (o ContainerResourceRequirementsOutput) Gpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerResourceRequirements) *int { return v.Gpu }).(pulumi.IntPtrOutput)
}

func (o ContainerResourceRequirementsOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerResourceRequirements) *float64 { return v.MemoryInGB }).(pulumi.Float64PtrOutput)
}

type ContainerResourceRequirementsPtrOutput struct{ *pulumi.OutputState }

func (ContainerResourceRequirementsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerResourceRequirements)(nil)).Elem()
}

func (o ContainerResourceRequirementsPtrOutput) ToContainerResourceRequirementsPtrOutput() ContainerResourceRequirementsPtrOutput {
	return o
}

func (o ContainerResourceRequirementsPtrOutput) ToContainerResourceRequirementsPtrOutputWithContext(ctx context.Context) ContainerResourceRequirementsPtrOutput {
	return o
}

func (o ContainerResourceRequirementsPtrOutput) Elem() ContainerResourceRequirementsOutput {
	return o.ApplyT(func(v *ContainerResourceRequirements) ContainerResourceRequirements {
		if v != nil {
			return *v
		}
		var ret ContainerResourceRequirements
		return ret
	}).(ContainerResourceRequirementsOutput)
}

func (o ContainerResourceRequirementsPtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContainerResourceRequirements) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ContainerResourceRequirementsPtrOutput) Fpga() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerResourceRequirements) *int {
		if v == nil {
			return nil
		}
		return v.Fpga
	}).(pulumi.IntPtrOutput)
}

func (o ContainerResourceRequirementsPtrOutput) Gpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerResourceRequirements) *int {
		if v == nil {
			return nil
		}
		return v.Gpu
	}).(pulumi.IntPtrOutput)
}

func (o ContainerResourceRequirementsPtrOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContainerResourceRequirements) *float64 {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.Float64PtrOutput)
}

type ContainerResourceRequirementsResponse struct {
	Cpu        *float64 `pulumi:"cpu"`
	Fpga       *int     `pulumi:"fpga"`
	Gpu        *int     `pulumi:"gpu"`
	MemoryInGB *float64 `pulumi:"memoryInGB"`
}

type CreateServiceRequestEnvironmentImageRequest struct {
	Assets               []ImageAsset                                 `pulumi:"assets"`
	DriverProgram        *string                                      `pulumi:"driverProgram"`
	Environment          *EnvironmentImageRequestEnvironment          `pulumi:"environment"`
	EnvironmentReference *EnvironmentImageRequestEnvironmentReference `pulumi:"environmentReference"`
	ModelIds             []string                                     `pulumi:"modelIds"`
	Models               []Model                                      `pulumi:"models"`
}





type CreateServiceRequestEnvironmentImageRequestInput interface {
	pulumi.Input

	ToCreateServiceRequestEnvironmentImageRequestOutput() CreateServiceRequestEnvironmentImageRequestOutput
	ToCreateServiceRequestEnvironmentImageRequestOutputWithContext(context.Context) CreateServiceRequestEnvironmentImageRequestOutput
}

type CreateServiceRequestEnvironmentImageRequestArgs struct {
	Assets               ImageAssetArrayInput                                `pulumi:"assets"`
	DriverProgram        pulumi.StringPtrInput                               `pulumi:"driverProgram"`
	Environment          EnvironmentImageRequestEnvironmentPtrInput          `pulumi:"environment"`
	EnvironmentReference EnvironmentImageRequestEnvironmentReferencePtrInput `pulumi:"environmentReference"`
	ModelIds             pulumi.StringArrayInput                             `pulumi:"modelIds"`
	Models               ModelArrayInput                                     `pulumi:"models"`
}

func (CreateServiceRequestEnvironmentImageRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateServiceRequestEnvironmentImageRequest)(nil)).Elem()
}

func (i CreateServiceRequestEnvironmentImageRequestArgs) ToCreateServiceRequestEnvironmentImageRequestOutput() CreateServiceRequestEnvironmentImageRequestOutput {
	return i.ToCreateServiceRequestEnvironmentImageRequestOutputWithContext(context.Background())
}

func (i CreateServiceRequestEnvironmentImageRequestArgs) ToCreateServiceRequestEnvironmentImageRequestOutputWithContext(ctx context.Context) CreateServiceRequestEnvironmentImageRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateServiceRequestEnvironmentImageRequestOutput)
}

func (i CreateServiceRequestEnvironmentImageRequestArgs) ToCreateServiceRequestEnvironmentImageRequestPtrOutput() CreateServiceRequestEnvironmentImageRequestPtrOutput {
	return i.ToCreateServiceRequestEnvironmentImageRequestPtrOutputWithContext(context.Background())
}

func (i CreateServiceRequestEnvironmentImageRequestArgs) ToCreateServiceRequestEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) CreateServiceRequestEnvironmentImageRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateServiceRequestEnvironmentImageRequestOutput).ToCreateServiceRequestEnvironmentImageRequestPtrOutputWithContext(ctx)
}









type CreateServiceRequestEnvironmentImageRequestPtrInput interface {
	pulumi.Input

	ToCreateServiceRequestEnvironmentImageRequestPtrOutput() CreateServiceRequestEnvironmentImageRequestPtrOutput
	ToCreateServiceRequestEnvironmentImageRequestPtrOutputWithContext(context.Context) CreateServiceRequestEnvironmentImageRequestPtrOutput
}

type createServiceRequestEnvironmentImageRequestPtrType CreateServiceRequestEnvironmentImageRequestArgs

func CreateServiceRequestEnvironmentImageRequestPtr(v *CreateServiceRequestEnvironmentImageRequestArgs) CreateServiceRequestEnvironmentImageRequestPtrInput {
	return (*createServiceRequestEnvironmentImageRequestPtrType)(v)
}

func (*createServiceRequestEnvironmentImageRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateServiceRequestEnvironmentImageRequest)(nil)).Elem()
}

func (i *createServiceRequestEnvironmentImageRequestPtrType) ToCreateServiceRequestEnvironmentImageRequestPtrOutput() CreateServiceRequestEnvironmentImageRequestPtrOutput {
	return i.ToCreateServiceRequestEnvironmentImageRequestPtrOutputWithContext(context.Background())
}

func (i *createServiceRequestEnvironmentImageRequestPtrType) ToCreateServiceRequestEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) CreateServiceRequestEnvironmentImageRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateServiceRequestEnvironmentImageRequestPtrOutput)
}

type CreateServiceRequestEnvironmentImageRequestOutput struct{ *pulumi.OutputState }

func (CreateServiceRequestEnvironmentImageRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateServiceRequestEnvironmentImageRequest)(nil)).Elem()
}

func (o CreateServiceRequestEnvironmentImageRequestOutput) ToCreateServiceRequestEnvironmentImageRequestOutput() CreateServiceRequestEnvironmentImageRequestOutput {
	return o
}

func (o CreateServiceRequestEnvironmentImageRequestOutput) ToCreateServiceRequestEnvironmentImageRequestOutputWithContext(ctx context.Context) CreateServiceRequestEnvironmentImageRequestOutput {
	return o
}

func (o CreateServiceRequestEnvironmentImageRequestOutput) ToCreateServiceRequestEnvironmentImageRequestPtrOutput() CreateServiceRequestEnvironmentImageRequestPtrOutput {
	return o.ToCreateServiceRequestEnvironmentImageRequestPtrOutputWithContext(context.Background())
}

func (o CreateServiceRequestEnvironmentImageRequestOutput) ToCreateServiceRequestEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) CreateServiceRequestEnvironmentImageRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateServiceRequestEnvironmentImageRequest) *CreateServiceRequestEnvironmentImageRequest {
		return &v
	}).(CreateServiceRequestEnvironmentImageRequestPtrOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestOutput) Assets() ImageAssetArrayOutput {
	return o.ApplyT(func(v CreateServiceRequestEnvironmentImageRequest) []ImageAsset { return v.Assets }).(ImageAssetArrayOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestOutput) DriverProgram() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateServiceRequestEnvironmentImageRequest) *string { return v.DriverProgram }).(pulumi.StringPtrOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestOutput) Environment() EnvironmentImageRequestEnvironmentPtrOutput {
	return o.ApplyT(func(v CreateServiceRequestEnvironmentImageRequest) *EnvironmentImageRequestEnvironment {
		return v.Environment
	}).(EnvironmentImageRequestEnvironmentPtrOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestOutput) EnvironmentReference() EnvironmentImageRequestEnvironmentReferencePtrOutput {
	return o.ApplyT(func(v CreateServiceRequestEnvironmentImageRequest) *EnvironmentImageRequestEnvironmentReference {
		return v.EnvironmentReference
	}).(EnvironmentImageRequestEnvironmentReferencePtrOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestOutput) ModelIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CreateServiceRequestEnvironmentImageRequest) []string { return v.ModelIds }).(pulumi.StringArrayOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestOutput) Models() ModelArrayOutput {
	return o.ApplyT(func(v CreateServiceRequestEnvironmentImageRequest) []Model { return v.Models }).(ModelArrayOutput)
}

type CreateServiceRequestEnvironmentImageRequestPtrOutput struct{ *pulumi.OutputState }

func (CreateServiceRequestEnvironmentImageRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateServiceRequestEnvironmentImageRequest)(nil)).Elem()
}

func (o CreateServiceRequestEnvironmentImageRequestPtrOutput) ToCreateServiceRequestEnvironmentImageRequestPtrOutput() CreateServiceRequestEnvironmentImageRequestPtrOutput {
	return o
}

func (o CreateServiceRequestEnvironmentImageRequestPtrOutput) ToCreateServiceRequestEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) CreateServiceRequestEnvironmentImageRequestPtrOutput {
	return o
}

func (o CreateServiceRequestEnvironmentImageRequestPtrOutput) Elem() CreateServiceRequestEnvironmentImageRequestOutput {
	return o.ApplyT(func(v *CreateServiceRequestEnvironmentImageRequest) CreateServiceRequestEnvironmentImageRequest {
		if v != nil {
			return *v
		}
		var ret CreateServiceRequestEnvironmentImageRequest
		return ret
	}).(CreateServiceRequestEnvironmentImageRequestOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestPtrOutput) Assets() ImageAssetArrayOutput {
	return o.ApplyT(func(v *CreateServiceRequestEnvironmentImageRequest) []ImageAsset {
		if v == nil {
			return nil
		}
		return v.Assets
	}).(ImageAssetArrayOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestPtrOutput) DriverProgram() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateServiceRequestEnvironmentImageRequest) *string {
		if v == nil {
			return nil
		}
		return v.DriverProgram
	}).(pulumi.StringPtrOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestPtrOutput) Environment() EnvironmentImageRequestEnvironmentPtrOutput {
	return o.ApplyT(func(v *CreateServiceRequestEnvironmentImageRequest) *EnvironmentImageRequestEnvironment {
		if v == nil {
			return nil
		}
		return v.Environment
	}).(EnvironmentImageRequestEnvironmentPtrOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestPtrOutput) EnvironmentReference() EnvironmentImageRequestEnvironmentReferencePtrOutput {
	return o.ApplyT(func(v *CreateServiceRequestEnvironmentImageRequest) *EnvironmentImageRequestEnvironmentReference {
		if v == nil {
			return nil
		}
		return v.EnvironmentReference
	}).(EnvironmentImageRequestEnvironmentReferencePtrOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestPtrOutput) ModelIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CreateServiceRequestEnvironmentImageRequest) []string {
		if v == nil {
			return nil
		}
		return v.ModelIds
	}).(pulumi.StringArrayOutput)
}

func (o CreateServiceRequestEnvironmentImageRequestPtrOutput) Models() ModelArrayOutput {
	return o.ApplyT(func(v *CreateServiceRequestEnvironmentImageRequest) []Model {
		if v == nil {
			return nil
		}
		return v.Models
	}).(ModelArrayOutput)
}

type CreateServiceRequestKeys struct {
	PrimaryKey   *string `pulumi:"primaryKey"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}





type CreateServiceRequestKeysInput interface {
	pulumi.Input

	ToCreateServiceRequestKeysOutput() CreateServiceRequestKeysOutput
	ToCreateServiceRequestKeysOutputWithContext(context.Context) CreateServiceRequestKeysOutput
}

type CreateServiceRequestKeysArgs struct {
	PrimaryKey   pulumi.StringPtrInput `pulumi:"primaryKey"`
	SecondaryKey pulumi.StringPtrInput `pulumi:"secondaryKey"`
}

func (CreateServiceRequestKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateServiceRequestKeys)(nil)).Elem()
}

func (i CreateServiceRequestKeysArgs) ToCreateServiceRequestKeysOutput() CreateServiceRequestKeysOutput {
	return i.ToCreateServiceRequestKeysOutputWithContext(context.Background())
}

func (i CreateServiceRequestKeysArgs) ToCreateServiceRequestKeysOutputWithContext(ctx context.Context) CreateServiceRequestKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateServiceRequestKeysOutput)
}

func (i CreateServiceRequestKeysArgs) ToCreateServiceRequestKeysPtrOutput() CreateServiceRequestKeysPtrOutput {
	return i.ToCreateServiceRequestKeysPtrOutputWithContext(context.Background())
}

func (i CreateServiceRequestKeysArgs) ToCreateServiceRequestKeysPtrOutputWithContext(ctx context.Context) CreateServiceRequestKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateServiceRequestKeysOutput).ToCreateServiceRequestKeysPtrOutputWithContext(ctx)
}









type CreateServiceRequestKeysPtrInput interface {
	pulumi.Input

	ToCreateServiceRequestKeysPtrOutput() CreateServiceRequestKeysPtrOutput
	ToCreateServiceRequestKeysPtrOutputWithContext(context.Context) CreateServiceRequestKeysPtrOutput
}

type createServiceRequestKeysPtrType CreateServiceRequestKeysArgs

func CreateServiceRequestKeysPtr(v *CreateServiceRequestKeysArgs) CreateServiceRequestKeysPtrInput {
	return (*createServiceRequestKeysPtrType)(v)
}

func (*createServiceRequestKeysPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateServiceRequestKeys)(nil)).Elem()
}

func (i *createServiceRequestKeysPtrType) ToCreateServiceRequestKeysPtrOutput() CreateServiceRequestKeysPtrOutput {
	return i.ToCreateServiceRequestKeysPtrOutputWithContext(context.Background())
}

func (i *createServiceRequestKeysPtrType) ToCreateServiceRequestKeysPtrOutputWithContext(ctx context.Context) CreateServiceRequestKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateServiceRequestKeysPtrOutput)
}

type CreateServiceRequestKeysOutput struct{ *pulumi.OutputState }

func (CreateServiceRequestKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateServiceRequestKeys)(nil)).Elem()
}

func (o CreateServiceRequestKeysOutput) ToCreateServiceRequestKeysOutput() CreateServiceRequestKeysOutput {
	return o
}

func (o CreateServiceRequestKeysOutput) ToCreateServiceRequestKeysOutputWithContext(ctx context.Context) CreateServiceRequestKeysOutput {
	return o
}

func (o CreateServiceRequestKeysOutput) ToCreateServiceRequestKeysPtrOutput() CreateServiceRequestKeysPtrOutput {
	return o.ToCreateServiceRequestKeysPtrOutputWithContext(context.Background())
}

func (o CreateServiceRequestKeysOutput) ToCreateServiceRequestKeysPtrOutputWithContext(ctx context.Context) CreateServiceRequestKeysPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateServiceRequestKeys) *CreateServiceRequestKeys {
		return &v
	}).(CreateServiceRequestKeysPtrOutput)
}

func (o CreateServiceRequestKeysOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateServiceRequestKeys) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o CreateServiceRequestKeysOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateServiceRequestKeys) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

type CreateServiceRequestKeysPtrOutput struct{ *pulumi.OutputState }

func (CreateServiceRequestKeysPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateServiceRequestKeys)(nil)).Elem()
}

func (o CreateServiceRequestKeysPtrOutput) ToCreateServiceRequestKeysPtrOutput() CreateServiceRequestKeysPtrOutput {
	return o
}

func (o CreateServiceRequestKeysPtrOutput) ToCreateServiceRequestKeysPtrOutputWithContext(ctx context.Context) CreateServiceRequestKeysPtrOutput {
	return o
}

func (o CreateServiceRequestKeysPtrOutput) Elem() CreateServiceRequestKeysOutput {
	return o.ApplyT(func(v *CreateServiceRequestKeys) CreateServiceRequestKeys {
		if v != nil {
			return *v
		}
		var ret CreateServiceRequestKeys
		return ret
	}).(CreateServiceRequestKeysOutput)
}

func (o CreateServiceRequestKeysPtrOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateServiceRequestKeys) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryKey
	}).(pulumi.StringPtrOutput)
}

func (o CreateServiceRequestKeysPtrOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateServiceRequestKeys) *string {
		if v == nil {
			return nil
		}
		return v.SecondaryKey
	}).(pulumi.StringPtrOutput)
}

type DataFactory struct {
	ComputeLocation *string `pulumi:"computeLocation"`
	ComputeType     string  `pulumi:"computeType"`
	Description     *string `pulumi:"description"`
	ResourceId      *string `pulumi:"resourceId"`
}

type DataFactoryResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type DataLakeAnalytics struct {
	ComputeLocation *string                      `pulumi:"computeLocation"`
	ComputeType     string                       `pulumi:"computeType"`
	Description     *string                      `pulumi:"description"`
	Properties      *DataLakeAnalyticsProperties `pulumi:"properties"`
	ResourceId      *string                      `pulumi:"resourceId"`
}

type DataLakeAnalyticsProperties struct {
	DataLakeStoreAccountName *string `pulumi:"dataLakeStoreAccountName"`
}

type DataLakeAnalyticsResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *DataLakeAnalyticsResponseProperties  `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type DataLakeAnalyticsResponseProperties struct {
	DataLakeStoreAccountName *string `pulumi:"dataLakeStoreAccountName"`
}

type Databricks struct {
	ComputeLocation *string               `pulumi:"computeLocation"`
	ComputeType     string                `pulumi:"computeType"`
	Description     *string               `pulumi:"description"`
	Properties      *DatabricksProperties `pulumi:"properties"`
	ResourceId      *string               `pulumi:"resourceId"`
}

type DatabricksProperties struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
}

type DatabricksResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *DatabricksResponseProperties         `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type DatabricksResponseProperties struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
}

type DatasetReference struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}





type DatasetReferenceInput interface {
	pulumi.Input

	ToDatasetReferenceOutput() DatasetReferenceOutput
	ToDatasetReferenceOutputWithContext(context.Context) DatasetReferenceOutput
}

type DatasetReferenceArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DatasetReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetReference)(nil)).Elem()
}

func (i DatasetReferenceArgs) ToDatasetReferenceOutput() DatasetReferenceOutput {
	return i.ToDatasetReferenceOutputWithContext(context.Background())
}

func (i DatasetReferenceArgs) ToDatasetReferenceOutputWithContext(ctx context.Context) DatasetReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetReferenceOutput)
}





type DatasetReferenceArrayInput interface {
	pulumi.Input

	ToDatasetReferenceArrayOutput() DatasetReferenceArrayOutput
	ToDatasetReferenceArrayOutputWithContext(context.Context) DatasetReferenceArrayOutput
}

type DatasetReferenceArray []DatasetReferenceInput

func (DatasetReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatasetReference)(nil)).Elem()
}

func (i DatasetReferenceArray) ToDatasetReferenceArrayOutput() DatasetReferenceArrayOutput {
	return i.ToDatasetReferenceArrayOutputWithContext(context.Background())
}

func (i DatasetReferenceArray) ToDatasetReferenceArrayOutputWithContext(ctx context.Context) DatasetReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetReferenceArrayOutput)
}

type DatasetReferenceOutput struct{ *pulumi.OutputState }

func (DatasetReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetReference)(nil)).Elem()
}

func (o DatasetReferenceOutput) ToDatasetReferenceOutput() DatasetReferenceOutput {
	return o
}

func (o DatasetReferenceOutput) ToDatasetReferenceOutputWithContext(ctx context.Context) DatasetReferenceOutput {
	return o
}

func (o DatasetReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o DatasetReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DatasetReferenceArrayOutput struct{ *pulumi.OutputState }

func (DatasetReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatasetReference)(nil)).Elem()
}

func (o DatasetReferenceArrayOutput) ToDatasetReferenceArrayOutput() DatasetReferenceArrayOutput {
	return o
}

func (o DatasetReferenceArrayOutput) ToDatasetReferenceArrayOutputWithContext(ctx context.Context) DatasetReferenceArrayOutput {
	return o
}

func (o DatasetReferenceArrayOutput) Index(i pulumi.IntInput) DatasetReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatasetReference {
		return vs[0].([]DatasetReference)[vs[1].(int)]
	}).(DatasetReferenceOutput)
}

type DatasetReferenceResponse struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}

type EncryptionProperty struct {
	KeyVaultProperties KeyVaultProperties `pulumi:"keyVaultProperties"`
	Status             string             `pulumi:"status"`
}





type EncryptionPropertyInput interface {
	pulumi.Input

	ToEncryptionPropertyOutput() EncryptionPropertyOutput
	ToEncryptionPropertyOutputWithContext(context.Context) EncryptionPropertyOutput
}

type EncryptionPropertyArgs struct {
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

type EnvironmentImageRequestEnvironment struct {
	Docker                  *ModelEnvironmentDefinitionDocker `pulumi:"docker"`
	EnvironmentVariables    map[string]string                 `pulumi:"environmentVariables"`
	InferencingStackVersion *string                           `pulumi:"inferencingStackVersion"`
	Name                    *string                           `pulumi:"name"`
	Python                  *ModelEnvironmentDefinitionPython `pulumi:"python"`
	R                       *ModelEnvironmentDefinitionR      `pulumi:"r"`
	Spark                   *ModelEnvironmentDefinitionSpark  `pulumi:"spark"`
	Version                 *string                           `pulumi:"version"`
}





type EnvironmentImageRequestEnvironmentInput interface {
	pulumi.Input

	ToEnvironmentImageRequestEnvironmentOutput() EnvironmentImageRequestEnvironmentOutput
	ToEnvironmentImageRequestEnvironmentOutputWithContext(context.Context) EnvironmentImageRequestEnvironmentOutput
}

type EnvironmentImageRequestEnvironmentArgs struct {
	Docker                  ModelEnvironmentDefinitionDockerPtrInput `pulumi:"docker"`
	EnvironmentVariables    pulumi.StringMapInput                    `pulumi:"environmentVariables"`
	InferencingStackVersion pulumi.StringPtrInput                    `pulumi:"inferencingStackVersion"`
	Name                    pulumi.StringPtrInput                    `pulumi:"name"`
	Python                  ModelEnvironmentDefinitionPythonPtrInput `pulumi:"python"`
	R                       ModelEnvironmentDefinitionRPtrInput      `pulumi:"r"`
	Spark                   ModelEnvironmentDefinitionSparkPtrInput  `pulumi:"spark"`
	Version                 pulumi.StringPtrInput                    `pulumi:"version"`
}

func (EnvironmentImageRequestEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentImageRequestEnvironment)(nil)).Elem()
}

func (i EnvironmentImageRequestEnvironmentArgs) ToEnvironmentImageRequestEnvironmentOutput() EnvironmentImageRequestEnvironmentOutput {
	return i.ToEnvironmentImageRequestEnvironmentOutputWithContext(context.Background())
}

func (i EnvironmentImageRequestEnvironmentArgs) ToEnvironmentImageRequestEnvironmentOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageRequestEnvironmentOutput)
}

func (i EnvironmentImageRequestEnvironmentArgs) ToEnvironmentImageRequestEnvironmentPtrOutput() EnvironmentImageRequestEnvironmentPtrOutput {
	return i.ToEnvironmentImageRequestEnvironmentPtrOutputWithContext(context.Background())
}

func (i EnvironmentImageRequestEnvironmentArgs) ToEnvironmentImageRequestEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageRequestEnvironmentOutput).ToEnvironmentImageRequestEnvironmentPtrOutputWithContext(ctx)
}









type EnvironmentImageRequestEnvironmentPtrInput interface {
	pulumi.Input

	ToEnvironmentImageRequestEnvironmentPtrOutput() EnvironmentImageRequestEnvironmentPtrOutput
	ToEnvironmentImageRequestEnvironmentPtrOutputWithContext(context.Context) EnvironmentImageRequestEnvironmentPtrOutput
}

type environmentImageRequestEnvironmentPtrType EnvironmentImageRequestEnvironmentArgs

func EnvironmentImageRequestEnvironmentPtr(v *EnvironmentImageRequestEnvironmentArgs) EnvironmentImageRequestEnvironmentPtrInput {
	return (*environmentImageRequestEnvironmentPtrType)(v)
}

func (*environmentImageRequestEnvironmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentImageRequestEnvironment)(nil)).Elem()
}

func (i *environmentImageRequestEnvironmentPtrType) ToEnvironmentImageRequestEnvironmentPtrOutput() EnvironmentImageRequestEnvironmentPtrOutput {
	return i.ToEnvironmentImageRequestEnvironmentPtrOutputWithContext(context.Background())
}

func (i *environmentImageRequestEnvironmentPtrType) ToEnvironmentImageRequestEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageRequestEnvironmentPtrOutput)
}

type EnvironmentImageRequestEnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentImageRequestEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentImageRequestEnvironment)(nil)).Elem()
}

func (o EnvironmentImageRequestEnvironmentOutput) ToEnvironmentImageRequestEnvironmentOutput() EnvironmentImageRequestEnvironmentOutput {
	return o
}

func (o EnvironmentImageRequestEnvironmentOutput) ToEnvironmentImageRequestEnvironmentOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentOutput {
	return o
}

func (o EnvironmentImageRequestEnvironmentOutput) ToEnvironmentImageRequestEnvironmentPtrOutput() EnvironmentImageRequestEnvironmentPtrOutput {
	return o.ToEnvironmentImageRequestEnvironmentPtrOutputWithContext(context.Background())
}

func (o EnvironmentImageRequestEnvironmentOutput) ToEnvironmentImageRequestEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentImageRequestEnvironment) *EnvironmentImageRequestEnvironment {
		return &v
	}).(EnvironmentImageRequestEnvironmentPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentOutput) Docker() ModelEnvironmentDefinitionDockerPtrOutput {
	return o.ApplyT(func(v EnvironmentImageRequestEnvironment) *ModelEnvironmentDefinitionDocker { return v.Docker }).(ModelEnvironmentDefinitionDockerPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentImageRequestEnvironment) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o EnvironmentImageRequestEnvironmentOutput) InferencingStackVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentImageRequestEnvironment) *string { return v.InferencingStackVersion }).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentImageRequestEnvironment) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentOutput) Python() ModelEnvironmentDefinitionPythonPtrOutput {
	return o.ApplyT(func(v EnvironmentImageRequestEnvironment) *ModelEnvironmentDefinitionPython { return v.Python }).(ModelEnvironmentDefinitionPythonPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentOutput) R() ModelEnvironmentDefinitionRPtrOutput {
	return o.ApplyT(func(v EnvironmentImageRequestEnvironment) *ModelEnvironmentDefinitionR { return v.R }).(ModelEnvironmentDefinitionRPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentOutput) Spark() ModelEnvironmentDefinitionSparkPtrOutput {
	return o.ApplyT(func(v EnvironmentImageRequestEnvironment) *ModelEnvironmentDefinitionSpark { return v.Spark }).(ModelEnvironmentDefinitionSparkPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentImageRequestEnvironment) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type EnvironmentImageRequestEnvironmentPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentImageRequestEnvironmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentImageRequestEnvironment)(nil)).Elem()
}

func (o EnvironmentImageRequestEnvironmentPtrOutput) ToEnvironmentImageRequestEnvironmentPtrOutput() EnvironmentImageRequestEnvironmentPtrOutput {
	return o
}

func (o EnvironmentImageRequestEnvironmentPtrOutput) ToEnvironmentImageRequestEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentPtrOutput {
	return o
}

func (o EnvironmentImageRequestEnvironmentPtrOutput) Elem() EnvironmentImageRequestEnvironmentOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironment) EnvironmentImageRequestEnvironment {
		if v != nil {
			return *v
		}
		var ret EnvironmentImageRequestEnvironment
		return ret
	}).(EnvironmentImageRequestEnvironmentOutput)
}

func (o EnvironmentImageRequestEnvironmentPtrOutput) Docker() ModelEnvironmentDefinitionDockerPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironment) *ModelEnvironmentDefinitionDocker {
		if v == nil {
			return nil
		}
		return v.Docker
	}).(ModelEnvironmentDefinitionDockerPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentPtrOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironment) map[string]string {
		if v == nil {
			return nil
		}
		return v.EnvironmentVariables
	}).(pulumi.StringMapOutput)
}

func (o EnvironmentImageRequestEnvironmentPtrOutput) InferencingStackVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironment) *string {
		if v == nil {
			return nil
		}
		return v.InferencingStackVersion
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironment) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentPtrOutput) Python() ModelEnvironmentDefinitionPythonPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironment) *ModelEnvironmentDefinitionPython {
		if v == nil {
			return nil
		}
		return v.Python
	}).(ModelEnvironmentDefinitionPythonPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentPtrOutput) R() ModelEnvironmentDefinitionRPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironment) *ModelEnvironmentDefinitionR {
		if v == nil {
			return nil
		}
		return v.R
	}).(ModelEnvironmentDefinitionRPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentPtrOutput) Spark() ModelEnvironmentDefinitionSparkPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironment) *ModelEnvironmentDefinitionSpark {
		if v == nil {
			return nil
		}
		return v.Spark
	}).(ModelEnvironmentDefinitionSparkPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironment) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type EnvironmentImageRequestEnvironmentReference struct {
	Name    *string `pulumi:"name"`
	Version *string `pulumi:"version"`
}





type EnvironmentImageRequestEnvironmentReferenceInput interface {
	pulumi.Input

	ToEnvironmentImageRequestEnvironmentReferenceOutput() EnvironmentImageRequestEnvironmentReferenceOutput
	ToEnvironmentImageRequestEnvironmentReferenceOutputWithContext(context.Context) EnvironmentImageRequestEnvironmentReferenceOutput
}

type EnvironmentImageRequestEnvironmentReferenceArgs struct {
	Name    pulumi.StringPtrInput `pulumi:"name"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (EnvironmentImageRequestEnvironmentReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentImageRequestEnvironmentReference)(nil)).Elem()
}

func (i EnvironmentImageRequestEnvironmentReferenceArgs) ToEnvironmentImageRequestEnvironmentReferenceOutput() EnvironmentImageRequestEnvironmentReferenceOutput {
	return i.ToEnvironmentImageRequestEnvironmentReferenceOutputWithContext(context.Background())
}

func (i EnvironmentImageRequestEnvironmentReferenceArgs) ToEnvironmentImageRequestEnvironmentReferenceOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageRequestEnvironmentReferenceOutput)
}

func (i EnvironmentImageRequestEnvironmentReferenceArgs) ToEnvironmentImageRequestEnvironmentReferencePtrOutput() EnvironmentImageRequestEnvironmentReferencePtrOutput {
	return i.ToEnvironmentImageRequestEnvironmentReferencePtrOutputWithContext(context.Background())
}

func (i EnvironmentImageRequestEnvironmentReferenceArgs) ToEnvironmentImageRequestEnvironmentReferencePtrOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageRequestEnvironmentReferenceOutput).ToEnvironmentImageRequestEnvironmentReferencePtrOutputWithContext(ctx)
}









type EnvironmentImageRequestEnvironmentReferencePtrInput interface {
	pulumi.Input

	ToEnvironmentImageRequestEnvironmentReferencePtrOutput() EnvironmentImageRequestEnvironmentReferencePtrOutput
	ToEnvironmentImageRequestEnvironmentReferencePtrOutputWithContext(context.Context) EnvironmentImageRequestEnvironmentReferencePtrOutput
}

type environmentImageRequestEnvironmentReferencePtrType EnvironmentImageRequestEnvironmentReferenceArgs

func EnvironmentImageRequestEnvironmentReferencePtr(v *EnvironmentImageRequestEnvironmentReferenceArgs) EnvironmentImageRequestEnvironmentReferencePtrInput {
	return (*environmentImageRequestEnvironmentReferencePtrType)(v)
}

func (*environmentImageRequestEnvironmentReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentImageRequestEnvironmentReference)(nil)).Elem()
}

func (i *environmentImageRequestEnvironmentReferencePtrType) ToEnvironmentImageRequestEnvironmentReferencePtrOutput() EnvironmentImageRequestEnvironmentReferencePtrOutput {
	return i.ToEnvironmentImageRequestEnvironmentReferencePtrOutputWithContext(context.Background())
}

func (i *environmentImageRequestEnvironmentReferencePtrType) ToEnvironmentImageRequestEnvironmentReferencePtrOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageRequestEnvironmentReferencePtrOutput)
}

type EnvironmentImageRequestEnvironmentReferenceOutput struct{ *pulumi.OutputState }

func (EnvironmentImageRequestEnvironmentReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentImageRequestEnvironmentReference)(nil)).Elem()
}

func (o EnvironmentImageRequestEnvironmentReferenceOutput) ToEnvironmentImageRequestEnvironmentReferenceOutput() EnvironmentImageRequestEnvironmentReferenceOutput {
	return o
}

func (o EnvironmentImageRequestEnvironmentReferenceOutput) ToEnvironmentImageRequestEnvironmentReferenceOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentReferenceOutput {
	return o
}

func (o EnvironmentImageRequestEnvironmentReferenceOutput) ToEnvironmentImageRequestEnvironmentReferencePtrOutput() EnvironmentImageRequestEnvironmentReferencePtrOutput {
	return o.ToEnvironmentImageRequestEnvironmentReferencePtrOutputWithContext(context.Background())
}

func (o EnvironmentImageRequestEnvironmentReferenceOutput) ToEnvironmentImageRequestEnvironmentReferencePtrOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentImageRequestEnvironmentReference) *EnvironmentImageRequestEnvironmentReference {
		return &v
	}).(EnvironmentImageRequestEnvironmentReferencePtrOutput)
}

func (o EnvironmentImageRequestEnvironmentReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentImageRequestEnvironmentReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentReferenceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentImageRequestEnvironmentReference) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type EnvironmentImageRequestEnvironmentReferencePtrOutput struct{ *pulumi.OutputState }

func (EnvironmentImageRequestEnvironmentReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentImageRequestEnvironmentReference)(nil)).Elem()
}

func (o EnvironmentImageRequestEnvironmentReferencePtrOutput) ToEnvironmentImageRequestEnvironmentReferencePtrOutput() EnvironmentImageRequestEnvironmentReferencePtrOutput {
	return o
}

func (o EnvironmentImageRequestEnvironmentReferencePtrOutput) ToEnvironmentImageRequestEnvironmentReferencePtrOutputWithContext(ctx context.Context) EnvironmentImageRequestEnvironmentReferencePtrOutput {
	return o
}

func (o EnvironmentImageRequestEnvironmentReferencePtrOutput) Elem() EnvironmentImageRequestEnvironmentReferenceOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironmentReference) EnvironmentImageRequestEnvironmentReference {
		if v != nil {
			return *v
		}
		var ret EnvironmentImageRequestEnvironmentReference
		return ret
	}).(EnvironmentImageRequestEnvironmentReferenceOutput)
}

func (o EnvironmentImageRequestEnvironmentReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironmentReference) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageRequestEnvironmentReferencePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageRequestEnvironmentReference) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type EnvironmentImageResponseResponseEnvironment struct {
	Docker                  *ModelEnvironmentDefinitionResponseResponseDocker `pulumi:"docker"`
	EnvironmentVariables    map[string]string                                 `pulumi:"environmentVariables"`
	InferencingStackVersion *string                                           `pulumi:"inferencingStackVersion"`
	Name                    *string                                           `pulumi:"name"`
	Python                  *ModelEnvironmentDefinitionResponseResponsePython `pulumi:"python"`
	R                       *ModelEnvironmentDefinitionResponseResponseR      `pulumi:"r"`
	Spark                   *ModelEnvironmentDefinitionResponseResponseSpark  `pulumi:"spark"`
	Version                 *string                                           `pulumi:"version"`
}

type EnvironmentImageResponseResponseEnvironmentReference struct {
	Name    *string `pulumi:"name"`
	Version *string `pulumi:"version"`
}

type ErrorDetailResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}

type ErrorResponseResponse struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
}

type HDInsight struct {
	ComputeLocation *string              `pulumi:"computeLocation"`
	ComputeType     string               `pulumi:"computeType"`
	Description     *string              `pulumi:"description"`
	Properties      *HDInsightProperties `pulumi:"properties"`
	ResourceId      *string              `pulumi:"resourceId"`
}

type HDInsightProperties struct {
	Address              *string                       `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	SshPort              *int                          `pulumi:"sshPort"`
}

type HDInsightResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *HDInsightResponseProperties          `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type HDInsightResponseProperties struct {
	Address              *string                               `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	SshPort              *int                                  `pulumi:"sshPort"`
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

type ImageAsset struct {
	Id       *string `pulumi:"id"`
	MimeType *string `pulumi:"mimeType"`
	Unpack   *bool   `pulumi:"unpack"`
	Url      *string `pulumi:"url"`
}





type ImageAssetInput interface {
	pulumi.Input

	ToImageAssetOutput() ImageAssetOutput
	ToImageAssetOutputWithContext(context.Context) ImageAssetOutput
}

type ImageAssetArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	MimeType pulumi.StringPtrInput `pulumi:"mimeType"`
	Unpack   pulumi.BoolPtrInput   `pulumi:"unpack"`
	Url      pulumi.StringPtrInput `pulumi:"url"`
}

func (ImageAssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageAsset)(nil)).Elem()
}

func (i ImageAssetArgs) ToImageAssetOutput() ImageAssetOutput {
	return i.ToImageAssetOutputWithContext(context.Background())
}

func (i ImageAssetArgs) ToImageAssetOutputWithContext(ctx context.Context) ImageAssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageAssetOutput)
}





type ImageAssetArrayInput interface {
	pulumi.Input

	ToImageAssetArrayOutput() ImageAssetArrayOutput
	ToImageAssetArrayOutputWithContext(context.Context) ImageAssetArrayOutput
}

type ImageAssetArray []ImageAssetInput

func (ImageAssetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageAsset)(nil)).Elem()
}

func (i ImageAssetArray) ToImageAssetArrayOutput() ImageAssetArrayOutput {
	return i.ToImageAssetArrayOutputWithContext(context.Background())
}

func (i ImageAssetArray) ToImageAssetArrayOutputWithContext(ctx context.Context) ImageAssetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageAssetArrayOutput)
}

type ImageAssetOutput struct{ *pulumi.OutputState }

func (ImageAssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageAsset)(nil)).Elem()
}

func (o ImageAssetOutput) ToImageAssetOutput() ImageAssetOutput {
	return o
}

func (o ImageAssetOutput) ToImageAssetOutputWithContext(ctx context.Context) ImageAssetOutput {
	return o
}

func (o ImageAssetOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageAsset) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ImageAssetOutput) MimeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageAsset) *string { return v.MimeType }).(pulumi.StringPtrOutput)
}

func (o ImageAssetOutput) Unpack() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageAsset) *bool { return v.Unpack }).(pulumi.BoolPtrOutput)
}

func (o ImageAssetOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageAsset) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ImageAssetArrayOutput struct{ *pulumi.OutputState }

func (ImageAssetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageAsset)(nil)).Elem()
}

func (o ImageAssetArrayOutput) ToImageAssetArrayOutput() ImageAssetArrayOutput {
	return o
}

func (o ImageAssetArrayOutput) ToImageAssetArrayOutputWithContext(ctx context.Context) ImageAssetArrayOutput {
	return o
}

func (o ImageAssetArrayOutput) Index(i pulumi.IntInput) ImageAssetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageAsset {
		return vs[0].([]ImageAsset)[vs[1].(int)]
	}).(ImageAssetOutput)
}

type ImageAssetResponse struct {
	Id       *string `pulumi:"id"`
	MimeType *string `pulumi:"mimeType"`
	Unpack   *bool   `pulumi:"unpack"`
	Url      *string `pulumi:"url"`
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
	AssetName                       string `pulumi:"assetName"`
	DatasetVersion                  string `pulumi:"datasetVersion"`
	EnableIncrementalDatasetRefresh *bool  `pulumi:"enableIncrementalDatasetRefresh"`
}





type LabelingDatasetConfigurationInput interface {
	pulumi.Input

	ToLabelingDatasetConfigurationOutput() LabelingDatasetConfigurationOutput
	ToLabelingDatasetConfigurationOutputWithContext(context.Context) LabelingDatasetConfigurationOutput
}

type LabelingDatasetConfigurationArgs struct {
	AssetName                       pulumi.StringInput  `pulumi:"assetName"`
	DatasetVersion                  pulumi.StringInput  `pulumi:"datasetVersion"`
	EnableIncrementalDatasetRefresh pulumi.BoolPtrInput `pulumi:"enableIncrementalDatasetRefresh"`
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

func (o LabelingDatasetConfigurationOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingDatasetConfiguration) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o LabelingDatasetConfigurationOutput) DatasetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingDatasetConfiguration) string { return v.DatasetVersion }).(pulumi.StringOutput)
}

func (o LabelingDatasetConfigurationOutput) EnableIncrementalDatasetRefresh() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabelingDatasetConfiguration) *bool { return v.EnableIncrementalDatasetRefresh }).(pulumi.BoolPtrOutput)
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
		return &v.AssetName
	}).(pulumi.StringPtrOutput)
}

func (o LabelingDatasetConfigurationPtrOutput) DatasetVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelingDatasetConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.DatasetVersion
	}).(pulumi.StringPtrOutput)
}

func (o LabelingDatasetConfigurationPtrOutput) EnableIncrementalDatasetRefresh() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabelingDatasetConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.EnableIncrementalDatasetRefresh
	}).(pulumi.BoolPtrOutput)
}

type LabelingDatasetConfigurationResponse struct {
	AssetName                       string `pulumi:"assetName"`
	DatasetVersion                  string `pulumi:"datasetVersion"`
	EnableIncrementalDatasetRefresh *bool  `pulumi:"enableIncrementalDatasetRefresh"`
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

func (o LabelingDatasetConfigurationResponseOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingDatasetConfigurationResponse) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o LabelingDatasetConfigurationResponseOutput) DatasetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingDatasetConfigurationResponse) string { return v.DatasetVersion }).(pulumi.StringOutput)
}

func (o LabelingDatasetConfigurationResponseOutput) EnableIncrementalDatasetRefresh() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabelingDatasetConfigurationResponse) *bool { return v.EnableIncrementalDatasetRefresh }).(pulumi.BoolPtrOutput)
}

type LabelingJobImageProperties struct {
	AnnotationType *string `pulumi:"annotationType"`
	MediaType      string  `pulumi:"mediaType"`
}





type LabelingJobImagePropertiesInput interface {
	pulumi.Input

	ToLabelingJobImagePropertiesOutput() LabelingJobImagePropertiesOutput
	ToLabelingJobImagePropertiesOutputWithContext(context.Context) LabelingJobImagePropertiesOutput
}

type LabelingJobImagePropertiesArgs struct {
	AnnotationType pulumi.StringPtrInput `pulumi:"annotationType"`
	MediaType      pulumi.StringInput    `pulumi:"mediaType"`
}

func (LabelingJobImagePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobImageProperties)(nil)).Elem()
}

func (i LabelingJobImagePropertiesArgs) ToLabelingJobImagePropertiesOutput() LabelingJobImagePropertiesOutput {
	return i.ToLabelingJobImagePropertiesOutputWithContext(context.Background())
}

func (i LabelingJobImagePropertiesArgs) ToLabelingJobImagePropertiesOutputWithContext(ctx context.Context) LabelingJobImagePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobImagePropertiesOutput)
}

func (i LabelingJobImagePropertiesArgs) ToLabelingJobImagePropertiesPtrOutput() LabelingJobImagePropertiesPtrOutput {
	return i.ToLabelingJobImagePropertiesPtrOutputWithContext(context.Background())
}

func (i LabelingJobImagePropertiesArgs) ToLabelingJobImagePropertiesPtrOutputWithContext(ctx context.Context) LabelingJobImagePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobImagePropertiesOutput).ToLabelingJobImagePropertiesPtrOutputWithContext(ctx)
}









type LabelingJobImagePropertiesPtrInput interface {
	pulumi.Input

	ToLabelingJobImagePropertiesPtrOutput() LabelingJobImagePropertiesPtrOutput
	ToLabelingJobImagePropertiesPtrOutputWithContext(context.Context) LabelingJobImagePropertiesPtrOutput
}

type labelingJobImagePropertiesPtrType LabelingJobImagePropertiesArgs

func LabelingJobImagePropertiesPtr(v *LabelingJobImagePropertiesArgs) LabelingJobImagePropertiesPtrInput {
	return (*labelingJobImagePropertiesPtrType)(v)
}

func (*labelingJobImagePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingJobImageProperties)(nil)).Elem()
}

func (i *labelingJobImagePropertiesPtrType) ToLabelingJobImagePropertiesPtrOutput() LabelingJobImagePropertiesPtrOutput {
	return i.ToLabelingJobImagePropertiesPtrOutputWithContext(context.Background())
}

func (i *labelingJobImagePropertiesPtrType) ToLabelingJobImagePropertiesPtrOutputWithContext(ctx context.Context) LabelingJobImagePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobImagePropertiesPtrOutput)
}

type LabelingJobImagePropertiesOutput struct{ *pulumi.OutputState }

func (LabelingJobImagePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobImageProperties)(nil)).Elem()
}

func (o LabelingJobImagePropertiesOutput) ToLabelingJobImagePropertiesOutput() LabelingJobImagePropertiesOutput {
	return o
}

func (o LabelingJobImagePropertiesOutput) ToLabelingJobImagePropertiesOutputWithContext(ctx context.Context) LabelingJobImagePropertiesOutput {
	return o
}

func (o LabelingJobImagePropertiesOutput) ToLabelingJobImagePropertiesPtrOutput() LabelingJobImagePropertiesPtrOutput {
	return o.ToLabelingJobImagePropertiesPtrOutputWithContext(context.Background())
}

func (o LabelingJobImagePropertiesOutput) ToLabelingJobImagePropertiesPtrOutputWithContext(ctx context.Context) LabelingJobImagePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabelingJobImageProperties) *LabelingJobImageProperties {
		return &v
	}).(LabelingJobImagePropertiesPtrOutput)
}

func (o LabelingJobImagePropertiesOutput) AnnotationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelingJobImageProperties) *string { return v.AnnotationType }).(pulumi.StringPtrOutput)
}

func (o LabelingJobImagePropertiesOutput) MediaType() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingJobImageProperties) string { return v.MediaType }).(pulumi.StringOutput)
}

type LabelingJobImagePropertiesPtrOutput struct{ *pulumi.OutputState }

func (LabelingJobImagePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingJobImageProperties)(nil)).Elem()
}

func (o LabelingJobImagePropertiesPtrOutput) ToLabelingJobImagePropertiesPtrOutput() LabelingJobImagePropertiesPtrOutput {
	return o
}

func (o LabelingJobImagePropertiesPtrOutput) ToLabelingJobImagePropertiesPtrOutputWithContext(ctx context.Context) LabelingJobImagePropertiesPtrOutput {
	return o
}

func (o LabelingJobImagePropertiesPtrOutput) Elem() LabelingJobImagePropertiesOutput {
	return o.ApplyT(func(v *LabelingJobImageProperties) LabelingJobImageProperties {
		if v != nil {
			return *v
		}
		var ret LabelingJobImageProperties
		return ret
	}).(LabelingJobImagePropertiesOutput)
}

func (o LabelingJobImagePropertiesPtrOutput) AnnotationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelingJobImageProperties) *string {
		if v == nil {
			return nil
		}
		return v.AnnotationType
	}).(pulumi.StringPtrOutput)
}

func (o LabelingJobImagePropertiesPtrOutput) MediaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelingJobImageProperties) *string {
		if v == nil {
			return nil
		}
		return &v.MediaType
	}).(pulumi.StringPtrOutput)
}

type LabelingJobImagePropertiesResponse struct {
	AnnotationType *string `pulumi:"annotationType"`
	MediaType      string  `pulumi:"mediaType"`
}

type LabelingJobImagePropertiesResponseOutput struct{ *pulumi.OutputState }

func (LabelingJobImagePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobImagePropertiesResponse)(nil)).Elem()
}

func (o LabelingJobImagePropertiesResponseOutput) ToLabelingJobImagePropertiesResponseOutput() LabelingJobImagePropertiesResponseOutput {
	return o
}

func (o LabelingJobImagePropertiesResponseOutput) ToLabelingJobImagePropertiesResponseOutputWithContext(ctx context.Context) LabelingJobImagePropertiesResponseOutput {
	return o
}

func (o LabelingJobImagePropertiesResponseOutput) AnnotationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabelingJobImagePropertiesResponse) *string { return v.AnnotationType }).(pulumi.StringPtrOutput)
}

func (o LabelingJobImagePropertiesResponseOutput) MediaType() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingJobImagePropertiesResponse) string { return v.MediaType }).(pulumi.StringOutput)
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

type LabelingJobProperties struct {
	DatasetConfiguration       LabelingDatasetConfiguration `pulumi:"datasetConfiguration"`
	JobInstructions            LabelingJobInstructions      `pulumi:"jobInstructions"`
	LabelCategories            map[string]LabelCategory     `pulumi:"labelCategories"`
	LabelingJobMediaProperties LabelingJobImageProperties   `pulumi:"labelingJobMediaProperties"`
	MlAssistConfiguration      *MLAssistConfiguration       `pulumi:"mlAssistConfiguration"`
	Properties                 map[string]string            `pulumi:"properties"`
	Tags                       map[string]string            `pulumi:"tags"`
}





type LabelingJobPropertiesInput interface {
	pulumi.Input

	ToLabelingJobPropertiesOutput() LabelingJobPropertiesOutput
	ToLabelingJobPropertiesOutputWithContext(context.Context) LabelingJobPropertiesOutput
}

type LabelingJobPropertiesArgs struct {
	DatasetConfiguration       LabelingDatasetConfigurationInput `pulumi:"datasetConfiguration"`
	JobInstructions            LabelingJobInstructionsInput      `pulumi:"jobInstructions"`
	LabelCategories            LabelCategoryMapInput             `pulumi:"labelCategories"`
	LabelingJobMediaProperties LabelingJobImagePropertiesInput   `pulumi:"labelingJobMediaProperties"`
	MlAssistConfiguration      MLAssistConfigurationPtrInput     `pulumi:"mlAssistConfiguration"`
	Properties                 pulumi.StringMapInput             `pulumi:"properties"`
	Tags                       pulumi.StringMapInput             `pulumi:"tags"`
}

func (LabelingJobPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobProperties)(nil)).Elem()
}

func (i LabelingJobPropertiesArgs) ToLabelingJobPropertiesOutput() LabelingJobPropertiesOutput {
	return i.ToLabelingJobPropertiesOutputWithContext(context.Background())
}

func (i LabelingJobPropertiesArgs) ToLabelingJobPropertiesOutputWithContext(ctx context.Context) LabelingJobPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobPropertiesOutput)
}

func (i LabelingJobPropertiesArgs) ToLabelingJobPropertiesPtrOutput() LabelingJobPropertiesPtrOutput {
	return i.ToLabelingJobPropertiesPtrOutputWithContext(context.Background())
}

func (i LabelingJobPropertiesArgs) ToLabelingJobPropertiesPtrOutputWithContext(ctx context.Context) LabelingJobPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobPropertiesOutput).ToLabelingJobPropertiesPtrOutputWithContext(ctx)
}









type LabelingJobPropertiesPtrInput interface {
	pulumi.Input

	ToLabelingJobPropertiesPtrOutput() LabelingJobPropertiesPtrOutput
	ToLabelingJobPropertiesPtrOutputWithContext(context.Context) LabelingJobPropertiesPtrOutput
}

type labelingJobPropertiesPtrType LabelingJobPropertiesArgs

func LabelingJobPropertiesPtr(v *LabelingJobPropertiesArgs) LabelingJobPropertiesPtrInput {
	return (*labelingJobPropertiesPtrType)(v)
}

func (*labelingJobPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingJobProperties)(nil)).Elem()
}

func (i *labelingJobPropertiesPtrType) ToLabelingJobPropertiesPtrOutput() LabelingJobPropertiesPtrOutput {
	return i.ToLabelingJobPropertiesPtrOutputWithContext(context.Background())
}

func (i *labelingJobPropertiesPtrType) ToLabelingJobPropertiesPtrOutputWithContext(ctx context.Context) LabelingJobPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobPropertiesPtrOutput)
}

type LabelingJobPropertiesOutput struct{ *pulumi.OutputState }

func (LabelingJobPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobProperties)(nil)).Elem()
}

func (o LabelingJobPropertiesOutput) ToLabelingJobPropertiesOutput() LabelingJobPropertiesOutput {
	return o
}

func (o LabelingJobPropertiesOutput) ToLabelingJobPropertiesOutputWithContext(ctx context.Context) LabelingJobPropertiesOutput {
	return o
}

func (o LabelingJobPropertiesOutput) ToLabelingJobPropertiesPtrOutput() LabelingJobPropertiesPtrOutput {
	return o.ToLabelingJobPropertiesPtrOutputWithContext(context.Background())
}

func (o LabelingJobPropertiesOutput) ToLabelingJobPropertiesPtrOutputWithContext(ctx context.Context) LabelingJobPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabelingJobProperties) *LabelingJobProperties {
		return &v
	}).(LabelingJobPropertiesPtrOutput)
}

func (o LabelingJobPropertiesOutput) DatasetConfiguration() LabelingDatasetConfigurationOutput {
	return o.ApplyT(func(v LabelingJobProperties) LabelingDatasetConfiguration { return v.DatasetConfiguration }).(LabelingDatasetConfigurationOutput)
}

func (o LabelingJobPropertiesOutput) JobInstructions() LabelingJobInstructionsOutput {
	return o.ApplyT(func(v LabelingJobProperties) LabelingJobInstructions { return v.JobInstructions }).(LabelingJobInstructionsOutput)
}

func (o LabelingJobPropertiesOutput) LabelCategories() LabelCategoryMapOutput {
	return o.ApplyT(func(v LabelingJobProperties) map[string]LabelCategory { return v.LabelCategories }).(LabelCategoryMapOutput)
}

func (o LabelingJobPropertiesOutput) LabelingJobMediaProperties() LabelingJobImagePropertiesOutput {
	return o.ApplyT(func(v LabelingJobProperties) LabelingJobImageProperties { return v.LabelingJobMediaProperties }).(LabelingJobImagePropertiesOutput)
}

func (o LabelingJobPropertiesOutput) MlAssistConfiguration() MLAssistConfigurationPtrOutput {
	return o.ApplyT(func(v LabelingJobProperties) *MLAssistConfiguration { return v.MlAssistConfiguration }).(MLAssistConfigurationPtrOutput)
}

func (o LabelingJobPropertiesOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabelingJobProperties) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o LabelingJobPropertiesOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabelingJobProperties) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type LabelingJobPropertiesPtrOutput struct{ *pulumi.OutputState }

func (LabelingJobPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingJobProperties)(nil)).Elem()
}

func (o LabelingJobPropertiesPtrOutput) ToLabelingJobPropertiesPtrOutput() LabelingJobPropertiesPtrOutput {
	return o
}

func (o LabelingJobPropertiesPtrOutput) ToLabelingJobPropertiesPtrOutputWithContext(ctx context.Context) LabelingJobPropertiesPtrOutput {
	return o
}

func (o LabelingJobPropertiesPtrOutput) Elem() LabelingJobPropertiesOutput {
	return o.ApplyT(func(v *LabelingJobProperties) LabelingJobProperties {
		if v != nil {
			return *v
		}
		var ret LabelingJobProperties
		return ret
	}).(LabelingJobPropertiesOutput)
}

func (o LabelingJobPropertiesPtrOutput) DatasetConfiguration() LabelingDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v *LabelingJobProperties) *LabelingDatasetConfiguration {
		if v == nil {
			return nil
		}
		return &v.DatasetConfiguration
	}).(LabelingDatasetConfigurationPtrOutput)
}

func (o LabelingJobPropertiesPtrOutput) JobInstructions() LabelingJobInstructionsPtrOutput {
	return o.ApplyT(func(v *LabelingJobProperties) *LabelingJobInstructions {
		if v == nil {
			return nil
		}
		return &v.JobInstructions
	}).(LabelingJobInstructionsPtrOutput)
}

func (o LabelingJobPropertiesPtrOutput) LabelCategories() LabelCategoryMapOutput {
	return o.ApplyT(func(v *LabelingJobProperties) map[string]LabelCategory {
		if v == nil {
			return nil
		}
		return v.LabelCategories
	}).(LabelCategoryMapOutput)
}

func (o LabelingJobPropertiesPtrOutput) LabelingJobMediaProperties() LabelingJobImagePropertiesPtrOutput {
	return o.ApplyT(func(v *LabelingJobProperties) *LabelingJobImageProperties {
		if v == nil {
			return nil
		}
		return &v.LabelingJobMediaProperties
	}).(LabelingJobImagePropertiesPtrOutput)
}

func (o LabelingJobPropertiesPtrOutput) MlAssistConfiguration() MLAssistConfigurationPtrOutput {
	return o.ApplyT(func(v *LabelingJobProperties) *MLAssistConfiguration {
		if v == nil {
			return nil
		}
		return v.MlAssistConfiguration
	}).(MLAssistConfigurationPtrOutput)
}

func (o LabelingJobPropertiesPtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LabelingJobProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o LabelingJobPropertiesPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LabelingJobProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type LabelingJobPropertiesResponse struct {
	CreatedTimeUtc             string                               `pulumi:"createdTimeUtc"`
	DatasetConfiguration       LabelingDatasetConfigurationResponse `pulumi:"datasetConfiguration"`
	JobInstructions            LabelingJobInstructionsResponse      `pulumi:"jobInstructions"`
	LabelCategories            map[string]LabelCategoryResponse     `pulumi:"labelCategories"`
	LabelingJobMediaProperties LabelingJobImagePropertiesResponse   `pulumi:"labelingJobMediaProperties"`
	MlAssistConfiguration      *MLAssistConfigurationResponse       `pulumi:"mlAssistConfiguration"`
	ProgressMetrics            ProgressMetricsResponse              `pulumi:"progressMetrics"`
	ProjectId                  string                               `pulumi:"projectId"`
	Properties                 map[string]string                    `pulumi:"properties"`
	Status                     string                               `pulumi:"status"`
	StatusMessages             []StatusMessageResponse              `pulumi:"statusMessages"`
	Tags                       map[string]string                    `pulumi:"tags"`
}

type LabelingJobPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LabelingJobPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJobPropertiesResponse)(nil)).Elem()
}

func (o LabelingJobPropertiesResponseOutput) ToLabelingJobPropertiesResponseOutput() LabelingJobPropertiesResponseOutput {
	return o
}

func (o LabelingJobPropertiesResponseOutput) ToLabelingJobPropertiesResponseOutputWithContext(ctx context.Context) LabelingJobPropertiesResponseOutput {
	return o
}

func (o LabelingJobPropertiesResponseOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o LabelingJobPropertiesResponseOutput) DatasetConfiguration() LabelingDatasetConfigurationResponseOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) LabelingDatasetConfigurationResponse {
		return v.DatasetConfiguration
	}).(LabelingDatasetConfigurationResponseOutput)
}

func (o LabelingJobPropertiesResponseOutput) JobInstructions() LabelingJobInstructionsResponseOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) LabelingJobInstructionsResponse { return v.JobInstructions }).(LabelingJobInstructionsResponseOutput)
}

func (o LabelingJobPropertiesResponseOutput) LabelCategories() LabelCategoryResponseMapOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) map[string]LabelCategoryResponse { return v.LabelCategories }).(LabelCategoryResponseMapOutput)
}

func (o LabelingJobPropertiesResponseOutput) LabelingJobMediaProperties() LabelingJobImagePropertiesResponseOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) LabelingJobImagePropertiesResponse {
		return v.LabelingJobMediaProperties
	}).(LabelingJobImagePropertiesResponseOutput)
}

func (o LabelingJobPropertiesResponseOutput) MlAssistConfiguration() MLAssistConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) *MLAssistConfigurationResponse { return v.MlAssistConfiguration }).(MLAssistConfigurationResponsePtrOutput)
}

func (o LabelingJobPropertiesResponseOutput) ProgressMetrics() ProgressMetricsResponseOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) ProgressMetricsResponse { return v.ProgressMetrics }).(ProgressMetricsResponseOutput)
}

func (o LabelingJobPropertiesResponseOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LabelingJobPropertiesResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o LabelingJobPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o LabelingJobPropertiesResponseOutput) StatusMessages() StatusMessageResponseArrayOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) []StatusMessageResponse { return v.StatusMessages }).(StatusMessageResponseArrayOutput)
}

func (o LabelingJobPropertiesResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabelingJobPropertiesResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type LinkedServiceProps struct {
	CreatedTime             *string                `pulumi:"createdTime"`
	LinkType                *LinkedServiceLinkType `pulumi:"linkType"`
	LinkedServiceResourceId string                 `pulumi:"linkedServiceResourceId"`
	ModifiedTime            *string                `pulumi:"modifiedTime"`
}





type LinkedServicePropsInput interface {
	pulumi.Input

	ToLinkedServicePropsOutput() LinkedServicePropsOutput
	ToLinkedServicePropsOutputWithContext(context.Context) LinkedServicePropsOutput
}

type LinkedServicePropsArgs struct {
	CreatedTime             pulumi.StringPtrInput         `pulumi:"createdTime"`
	LinkType                LinkedServiceLinkTypePtrInput `pulumi:"linkType"`
	LinkedServiceResourceId pulumi.StringInput            `pulumi:"linkedServiceResourceId"`
	ModifiedTime            pulumi.StringPtrInput         `pulumi:"modifiedTime"`
}

func (LinkedServicePropsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceProps)(nil)).Elem()
}

func (i LinkedServicePropsArgs) ToLinkedServicePropsOutput() LinkedServicePropsOutput {
	return i.ToLinkedServicePropsOutputWithContext(context.Background())
}

func (i LinkedServicePropsArgs) ToLinkedServicePropsOutputWithContext(ctx context.Context) LinkedServicePropsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServicePropsOutput)
}

func (i LinkedServicePropsArgs) ToLinkedServicePropsPtrOutput() LinkedServicePropsPtrOutput {
	return i.ToLinkedServicePropsPtrOutputWithContext(context.Background())
}

func (i LinkedServicePropsArgs) ToLinkedServicePropsPtrOutputWithContext(ctx context.Context) LinkedServicePropsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServicePropsOutput).ToLinkedServicePropsPtrOutputWithContext(ctx)
}









type LinkedServicePropsPtrInput interface {
	pulumi.Input

	ToLinkedServicePropsPtrOutput() LinkedServicePropsPtrOutput
	ToLinkedServicePropsPtrOutputWithContext(context.Context) LinkedServicePropsPtrOutput
}

type linkedServicePropsPtrType LinkedServicePropsArgs

func LinkedServicePropsPtr(v *LinkedServicePropsArgs) LinkedServicePropsPtrInput {
	return (*linkedServicePropsPtrType)(v)
}

func (*linkedServicePropsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceProps)(nil)).Elem()
}

func (i *linkedServicePropsPtrType) ToLinkedServicePropsPtrOutput() LinkedServicePropsPtrOutput {
	return i.ToLinkedServicePropsPtrOutputWithContext(context.Background())
}

func (i *linkedServicePropsPtrType) ToLinkedServicePropsPtrOutputWithContext(ctx context.Context) LinkedServicePropsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServicePropsPtrOutput)
}

type LinkedServicePropsOutput struct{ *pulumi.OutputState }

func (LinkedServicePropsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceProps)(nil)).Elem()
}

func (o LinkedServicePropsOutput) ToLinkedServicePropsOutput() LinkedServicePropsOutput {
	return o
}

func (o LinkedServicePropsOutput) ToLinkedServicePropsOutputWithContext(ctx context.Context) LinkedServicePropsOutput {
	return o
}

func (o LinkedServicePropsOutput) ToLinkedServicePropsPtrOutput() LinkedServicePropsPtrOutput {
	return o.ToLinkedServicePropsPtrOutputWithContext(context.Background())
}

func (o LinkedServicePropsOutput) ToLinkedServicePropsPtrOutputWithContext(ctx context.Context) LinkedServicePropsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinkedServiceProps) *LinkedServiceProps {
		return &v
	}).(LinkedServicePropsPtrOutput)
}

func (o LinkedServicePropsOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedServiceProps) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o LinkedServicePropsOutput) LinkType() LinkedServiceLinkTypePtrOutput {
	return o.ApplyT(func(v LinkedServiceProps) *LinkedServiceLinkType { return v.LinkType }).(LinkedServiceLinkTypePtrOutput)
}

func (o LinkedServicePropsOutput) LinkedServiceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedServiceProps) string { return v.LinkedServiceResourceId }).(pulumi.StringOutput)
}

func (o LinkedServicePropsOutput) ModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedServiceProps) *string { return v.ModifiedTime }).(pulumi.StringPtrOutput)
}

type LinkedServicePropsPtrOutput struct{ *pulumi.OutputState }

func (LinkedServicePropsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceProps)(nil)).Elem()
}

func (o LinkedServicePropsPtrOutput) ToLinkedServicePropsPtrOutput() LinkedServicePropsPtrOutput {
	return o
}

func (o LinkedServicePropsPtrOutput) ToLinkedServicePropsPtrOutputWithContext(ctx context.Context) LinkedServicePropsPtrOutput {
	return o
}

func (o LinkedServicePropsPtrOutput) Elem() LinkedServicePropsOutput {
	return o.ApplyT(func(v *LinkedServiceProps) LinkedServiceProps {
		if v != nil {
			return *v
		}
		var ret LinkedServiceProps
		return ret
	}).(LinkedServicePropsOutput)
}

func (o LinkedServicePropsPtrOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedServiceProps) *string {
		if v == nil {
			return nil
		}
		return v.CreatedTime
	}).(pulumi.StringPtrOutput)
}

func (o LinkedServicePropsPtrOutput) LinkType() LinkedServiceLinkTypePtrOutput {
	return o.ApplyT(func(v *LinkedServiceProps) *LinkedServiceLinkType {
		if v == nil {
			return nil
		}
		return v.LinkType
	}).(LinkedServiceLinkTypePtrOutput)
}

func (o LinkedServicePropsPtrOutput) LinkedServiceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedServiceProps) *string {
		if v == nil {
			return nil
		}
		return &v.LinkedServiceResourceId
	}).(pulumi.StringPtrOutput)
}

func (o LinkedServicePropsPtrOutput) ModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedServiceProps) *string {
		if v == nil {
			return nil
		}
		return v.ModifiedTime
	}).(pulumi.StringPtrOutput)
}

type LinkedServicePropsResponse struct {
	CreatedTime             *string `pulumi:"createdTime"`
	LinkType                *string `pulumi:"linkType"`
	LinkedServiceResourceId string  `pulumi:"linkedServiceResourceId"`
	ModifiedTime            *string `pulumi:"modifiedTime"`
}

type LinkedServicePropsResponseOutput struct{ *pulumi.OutputState }

func (LinkedServicePropsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServicePropsResponse)(nil)).Elem()
}

func (o LinkedServicePropsResponseOutput) ToLinkedServicePropsResponseOutput() LinkedServicePropsResponseOutput {
	return o
}

func (o LinkedServicePropsResponseOutput) ToLinkedServicePropsResponseOutputWithContext(ctx context.Context) LinkedServicePropsResponseOutput {
	return o
}

func (o LinkedServicePropsResponseOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedServicePropsResponse) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o LinkedServicePropsResponseOutput) LinkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedServicePropsResponse) *string { return v.LinkType }).(pulumi.StringPtrOutput)
}

func (o LinkedServicePropsResponseOutput) LinkedServiceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedServicePropsResponse) string { return v.LinkedServiceResourceId }).(pulumi.StringOutput)
}

func (o LinkedServicePropsResponseOutput) ModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedServicePropsResponse) *string { return v.ModifiedTime }).(pulumi.StringPtrOutput)
}

type MLAssistConfiguration struct {
	InferencingComputeBinding ComputeBinding `pulumi:"inferencingComputeBinding"`
	MlAssistEnabled           *bool          `pulumi:"mlAssistEnabled"`
	ModelNamePrefix           string         `pulumi:"modelNamePrefix"`
	PrelabelAccuracyThreshold *float64       `pulumi:"prelabelAccuracyThreshold"`
	TrainingComputeBinding    ComputeBinding `pulumi:"trainingComputeBinding"`
}





type MLAssistConfigurationInput interface {
	pulumi.Input

	ToMLAssistConfigurationOutput() MLAssistConfigurationOutput
	ToMLAssistConfigurationOutputWithContext(context.Context) MLAssistConfigurationOutput
}

type MLAssistConfigurationArgs struct {
	InferencingComputeBinding ComputeBindingInput    `pulumi:"inferencingComputeBinding"`
	MlAssistEnabled           pulumi.BoolPtrInput    `pulumi:"mlAssistEnabled"`
	ModelNamePrefix           pulumi.StringInput     `pulumi:"modelNamePrefix"`
	PrelabelAccuracyThreshold pulumi.Float64PtrInput `pulumi:"prelabelAccuracyThreshold"`
	TrainingComputeBinding    ComputeBindingInput    `pulumi:"trainingComputeBinding"`
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

func (o MLAssistConfigurationOutput) InferencingComputeBinding() ComputeBindingOutput {
	return o.ApplyT(func(v MLAssistConfiguration) ComputeBinding { return v.InferencingComputeBinding }).(ComputeBindingOutput)
}

func (o MLAssistConfigurationOutput) MlAssistEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MLAssistConfiguration) *bool { return v.MlAssistEnabled }).(pulumi.BoolPtrOutput)
}

func (o MLAssistConfigurationOutput) ModelNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v MLAssistConfiguration) string { return v.ModelNamePrefix }).(pulumi.StringOutput)
}

func (o MLAssistConfigurationOutput) PrelabelAccuracyThreshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MLAssistConfiguration) *float64 { return v.PrelabelAccuracyThreshold }).(pulumi.Float64PtrOutput)
}

func (o MLAssistConfigurationOutput) TrainingComputeBinding() ComputeBindingOutput {
	return o.ApplyT(func(v MLAssistConfiguration) ComputeBinding { return v.TrainingComputeBinding }).(ComputeBindingOutput)
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

func (o MLAssistConfigurationPtrOutput) InferencingComputeBinding() ComputeBindingPtrOutput {
	return o.ApplyT(func(v *MLAssistConfiguration) *ComputeBinding {
		if v == nil {
			return nil
		}
		return &v.InferencingComputeBinding
	}).(ComputeBindingPtrOutput)
}

func (o MLAssistConfigurationPtrOutput) MlAssistEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MLAssistConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.MlAssistEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o MLAssistConfigurationPtrOutput) ModelNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MLAssistConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.ModelNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o MLAssistConfigurationPtrOutput) PrelabelAccuracyThreshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MLAssistConfiguration) *float64 {
		if v == nil {
			return nil
		}
		return v.PrelabelAccuracyThreshold
	}).(pulumi.Float64PtrOutput)
}

func (o MLAssistConfigurationPtrOutput) TrainingComputeBinding() ComputeBindingPtrOutput {
	return o.ApplyT(func(v *MLAssistConfiguration) *ComputeBinding {
		if v == nil {
			return nil
		}
		return &v.TrainingComputeBinding
	}).(ComputeBindingPtrOutput)
}

type MLAssistConfigurationResponse struct {
	InferencingComputeBinding ComputeBindingResponse `pulumi:"inferencingComputeBinding"`
	MlAssistEnabled           *bool                  `pulumi:"mlAssistEnabled"`
	ModelNamePrefix           string                 `pulumi:"modelNamePrefix"`
	PrelabelAccuracyThreshold *float64               `pulumi:"prelabelAccuracyThreshold"`
	TrainingComputeBinding    ComputeBindingResponse `pulumi:"trainingComputeBinding"`
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

func (o MLAssistConfigurationResponseOutput) InferencingComputeBinding() ComputeBindingResponseOutput {
	return o.ApplyT(func(v MLAssistConfigurationResponse) ComputeBindingResponse { return v.InferencingComputeBinding }).(ComputeBindingResponseOutput)
}

func (o MLAssistConfigurationResponseOutput) MlAssistEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MLAssistConfigurationResponse) *bool { return v.MlAssistEnabled }).(pulumi.BoolPtrOutput)
}

func (o MLAssistConfigurationResponseOutput) ModelNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v MLAssistConfigurationResponse) string { return v.ModelNamePrefix }).(pulumi.StringOutput)
}

func (o MLAssistConfigurationResponseOutput) PrelabelAccuracyThreshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MLAssistConfigurationResponse) *float64 { return v.PrelabelAccuracyThreshold }).(pulumi.Float64PtrOutput)
}

func (o MLAssistConfigurationResponseOutput) TrainingComputeBinding() ComputeBindingResponseOutput {
	return o.ApplyT(func(v MLAssistConfigurationResponse) ComputeBindingResponse { return v.TrainingComputeBinding }).(ComputeBindingResponseOutput)
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

func (o MLAssistConfigurationResponsePtrOutput) InferencingComputeBinding() ComputeBindingResponsePtrOutput {
	return o.ApplyT(func(v *MLAssistConfigurationResponse) *ComputeBindingResponse {
		if v == nil {
			return nil
		}
		return &v.InferencingComputeBinding
	}).(ComputeBindingResponsePtrOutput)
}

func (o MLAssistConfigurationResponsePtrOutput) MlAssistEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MLAssistConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.MlAssistEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o MLAssistConfigurationResponsePtrOutput) ModelNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MLAssistConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ModelNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o MLAssistConfigurationResponsePtrOutput) PrelabelAccuracyThreshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MLAssistConfigurationResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.PrelabelAccuracyThreshold
	}).(pulumi.Float64PtrOutput)
}

func (o MLAssistConfigurationResponsePtrOutput) TrainingComputeBinding() ComputeBindingResponsePtrOutput {
	return o.ApplyT(func(v *MLAssistConfigurationResponse) *ComputeBindingResponse {
		if v == nil {
			return nil
		}
		return &v.TrainingComputeBinding
	}).(ComputeBindingResponsePtrOutput)
}

type MachineLearningServiceErrorResponse struct {
	Error ErrorResponseResponse `pulumi:"error"`
}

type Model struct {
	CreatedTime          *string                        `pulumi:"createdTime"`
	Datasets             []DatasetReference             `pulumi:"datasets"`
	DerivedModelIds      []string                       `pulumi:"derivedModelIds"`
	Description          *string                        `pulumi:"description"`
	ExperimentName       *string                        `pulumi:"experimentName"`
	Framework            *string                        `pulumi:"framework"`
	FrameworkVersion     *string                        `pulumi:"frameworkVersion"`
	Id                   *string                        `pulumi:"id"`
	KvTags               map[string]string              `pulumi:"kvTags"`
	MimeType             string                         `pulumi:"mimeType"`
	ModifiedTime         *string                        `pulumi:"modifiedTime"`
	Name                 string                         `pulumi:"name"`
	ParentModelId        *string                        `pulumi:"parentModelId"`
	Properties           map[string]string              `pulumi:"properties"`
	ResourceRequirements *ContainerResourceRequirements `pulumi:"resourceRequirements"`
	RunId                *string                        `pulumi:"runId"`
	SampleInputData      *string                        `pulumi:"sampleInputData"`
	SampleOutputData     *string                        `pulumi:"sampleOutputData"`
	Unpack               *bool                          `pulumi:"unpack"`
	Url                  string                         `pulumi:"url"`
	Version              *float64                       `pulumi:"version"`
}





type ModelInput interface {
	pulumi.Input

	ToModelOutput() ModelOutput
	ToModelOutputWithContext(context.Context) ModelOutput
}

type ModelArgs struct {
	CreatedTime          pulumi.StringPtrInput                 `pulumi:"createdTime"`
	Datasets             DatasetReferenceArrayInput            `pulumi:"datasets"`
	DerivedModelIds      pulumi.StringArrayInput               `pulumi:"derivedModelIds"`
	Description          pulumi.StringPtrInput                 `pulumi:"description"`
	ExperimentName       pulumi.StringPtrInput                 `pulumi:"experimentName"`
	Framework            pulumi.StringPtrInput                 `pulumi:"framework"`
	FrameworkVersion     pulumi.StringPtrInput                 `pulumi:"frameworkVersion"`
	Id                   pulumi.StringPtrInput                 `pulumi:"id"`
	KvTags               pulumi.StringMapInput                 `pulumi:"kvTags"`
	MimeType             pulumi.StringInput                    `pulumi:"mimeType"`
	ModifiedTime         pulumi.StringPtrInput                 `pulumi:"modifiedTime"`
	Name                 pulumi.StringInput                    `pulumi:"name"`
	ParentModelId        pulumi.StringPtrInput                 `pulumi:"parentModelId"`
	Properties           pulumi.StringMapInput                 `pulumi:"properties"`
	ResourceRequirements ContainerResourceRequirementsPtrInput `pulumi:"resourceRequirements"`
	RunId                pulumi.StringPtrInput                 `pulumi:"runId"`
	SampleInputData      pulumi.StringPtrInput                 `pulumi:"sampleInputData"`
	SampleOutputData     pulumi.StringPtrInput                 `pulumi:"sampleOutputData"`
	Unpack               pulumi.BoolPtrInput                   `pulumi:"unpack"`
	Url                  pulumi.StringInput                    `pulumi:"url"`
	Version              pulumi.Float64PtrInput                `pulumi:"version"`
}

func (ModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Model)(nil)).Elem()
}

func (i ModelArgs) ToModelOutput() ModelOutput {
	return i.ToModelOutputWithContext(context.Background())
}

func (i ModelArgs) ToModelOutputWithContext(ctx context.Context) ModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelOutput)
}





type ModelArrayInput interface {
	pulumi.Input

	ToModelArrayOutput() ModelArrayOutput
	ToModelArrayOutputWithContext(context.Context) ModelArrayOutput
}

type ModelArray []ModelInput

func (ModelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Model)(nil)).Elem()
}

func (i ModelArray) ToModelArrayOutput() ModelArrayOutput {
	return i.ToModelArrayOutputWithContext(context.Background())
}

func (i ModelArray) ToModelArrayOutputWithContext(ctx context.Context) ModelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelArrayOutput)
}

type ModelOutput struct{ *pulumi.OutputState }

func (ModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Model)(nil)).Elem()
}

func (o ModelOutput) ToModelOutput() ModelOutput {
	return o
}

func (o ModelOutput) ToModelOutputWithContext(ctx context.Context) ModelOutput {
	return o
}

func (o ModelOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Model) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) Datasets() DatasetReferenceArrayOutput {
	return o.ApplyT(func(v Model) []DatasetReference { return v.Datasets }).(DatasetReferenceArrayOutput)
}

func (o ModelOutput) DerivedModelIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Model) []string { return v.DerivedModelIds }).(pulumi.StringArrayOutput)
}

func (o ModelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Model) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) ExperimentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Model) *string { return v.ExperimentName }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) Framework() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Model) *string { return v.Framework }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) FrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Model) *string { return v.FrameworkVersion }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Model) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) KvTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v Model) map[string]string { return v.KvTags }).(pulumi.StringMapOutput)
}

func (o ModelOutput) MimeType() pulumi.StringOutput {
	return o.ApplyT(func(v Model) string { return v.MimeType }).(pulumi.StringOutput)
}

func (o ModelOutput) ModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Model) *string { return v.ModifiedTime }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Model) string { return v.Name }).(pulumi.StringOutput)
}

func (o ModelOutput) ParentModelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Model) *string { return v.ParentModelId }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v Model) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ModelOutput) ResourceRequirements() ContainerResourceRequirementsPtrOutput {
	return o.ApplyT(func(v Model) *ContainerResourceRequirements { return v.ResourceRequirements }).(ContainerResourceRequirementsPtrOutput)
}

func (o ModelOutput) RunId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Model) *string { return v.RunId }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) SampleInputData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Model) *string { return v.SampleInputData }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) SampleOutputData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Model) *string { return v.SampleOutputData }).(pulumi.StringPtrOutput)
}

func (o ModelOutput) Unpack() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Model) *bool { return v.Unpack }).(pulumi.BoolPtrOutput)
}

func (o ModelOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v Model) string { return v.Url }).(pulumi.StringOutput)
}

func (o ModelOutput) Version() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v Model) *float64 { return v.Version }).(pulumi.Float64PtrOutput)
}

type ModelArrayOutput struct{ *pulumi.OutputState }

func (ModelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Model)(nil)).Elem()
}

func (o ModelArrayOutput) ToModelArrayOutput() ModelArrayOutput {
	return o
}

func (o ModelArrayOutput) ToModelArrayOutputWithContext(ctx context.Context) ModelArrayOutput {
	return o
}

func (o ModelArrayOutput) Index(i pulumi.IntInput) ModelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Model {
		return vs[0].([]Model)[vs[1].(int)]
	}).(ModelOutput)
}

type ModelDockerSectionBaseImageRegistry struct {
	Address  *string `pulumi:"address"`
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}





type ModelDockerSectionBaseImageRegistryInput interface {
	pulumi.Input

	ToModelDockerSectionBaseImageRegistryOutput() ModelDockerSectionBaseImageRegistryOutput
	ToModelDockerSectionBaseImageRegistryOutputWithContext(context.Context) ModelDockerSectionBaseImageRegistryOutput
}

type ModelDockerSectionBaseImageRegistryArgs struct {
	Address  pulumi.StringPtrInput `pulumi:"address"`
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (ModelDockerSectionBaseImageRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelDockerSectionBaseImageRegistry)(nil)).Elem()
}

func (i ModelDockerSectionBaseImageRegistryArgs) ToModelDockerSectionBaseImageRegistryOutput() ModelDockerSectionBaseImageRegistryOutput {
	return i.ToModelDockerSectionBaseImageRegistryOutputWithContext(context.Background())
}

func (i ModelDockerSectionBaseImageRegistryArgs) ToModelDockerSectionBaseImageRegistryOutputWithContext(ctx context.Context) ModelDockerSectionBaseImageRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelDockerSectionBaseImageRegistryOutput)
}

func (i ModelDockerSectionBaseImageRegistryArgs) ToModelDockerSectionBaseImageRegistryPtrOutput() ModelDockerSectionBaseImageRegistryPtrOutput {
	return i.ToModelDockerSectionBaseImageRegistryPtrOutputWithContext(context.Background())
}

func (i ModelDockerSectionBaseImageRegistryArgs) ToModelDockerSectionBaseImageRegistryPtrOutputWithContext(ctx context.Context) ModelDockerSectionBaseImageRegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelDockerSectionBaseImageRegistryOutput).ToModelDockerSectionBaseImageRegistryPtrOutputWithContext(ctx)
}









type ModelDockerSectionBaseImageRegistryPtrInput interface {
	pulumi.Input

	ToModelDockerSectionBaseImageRegistryPtrOutput() ModelDockerSectionBaseImageRegistryPtrOutput
	ToModelDockerSectionBaseImageRegistryPtrOutputWithContext(context.Context) ModelDockerSectionBaseImageRegistryPtrOutput
}

type modelDockerSectionBaseImageRegistryPtrType ModelDockerSectionBaseImageRegistryArgs

func ModelDockerSectionBaseImageRegistryPtr(v *ModelDockerSectionBaseImageRegistryArgs) ModelDockerSectionBaseImageRegistryPtrInput {
	return (*modelDockerSectionBaseImageRegistryPtrType)(v)
}

func (*modelDockerSectionBaseImageRegistryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelDockerSectionBaseImageRegistry)(nil)).Elem()
}

func (i *modelDockerSectionBaseImageRegistryPtrType) ToModelDockerSectionBaseImageRegistryPtrOutput() ModelDockerSectionBaseImageRegistryPtrOutput {
	return i.ToModelDockerSectionBaseImageRegistryPtrOutputWithContext(context.Background())
}

func (i *modelDockerSectionBaseImageRegistryPtrType) ToModelDockerSectionBaseImageRegistryPtrOutputWithContext(ctx context.Context) ModelDockerSectionBaseImageRegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelDockerSectionBaseImageRegistryPtrOutput)
}

type ModelDockerSectionBaseImageRegistryOutput struct{ *pulumi.OutputState }

func (ModelDockerSectionBaseImageRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelDockerSectionBaseImageRegistry)(nil)).Elem()
}

func (o ModelDockerSectionBaseImageRegistryOutput) ToModelDockerSectionBaseImageRegistryOutput() ModelDockerSectionBaseImageRegistryOutput {
	return o
}

func (o ModelDockerSectionBaseImageRegistryOutput) ToModelDockerSectionBaseImageRegistryOutputWithContext(ctx context.Context) ModelDockerSectionBaseImageRegistryOutput {
	return o
}

func (o ModelDockerSectionBaseImageRegistryOutput) ToModelDockerSectionBaseImageRegistryPtrOutput() ModelDockerSectionBaseImageRegistryPtrOutput {
	return o.ToModelDockerSectionBaseImageRegistryPtrOutputWithContext(context.Background())
}

func (o ModelDockerSectionBaseImageRegistryOutput) ToModelDockerSectionBaseImageRegistryPtrOutputWithContext(ctx context.Context) ModelDockerSectionBaseImageRegistryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelDockerSectionBaseImageRegistry) *ModelDockerSectionBaseImageRegistry {
		return &v
	}).(ModelDockerSectionBaseImageRegistryPtrOutput)
}

func (o ModelDockerSectionBaseImageRegistryOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelDockerSectionBaseImageRegistry) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o ModelDockerSectionBaseImageRegistryOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelDockerSectionBaseImageRegistry) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ModelDockerSectionBaseImageRegistryOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelDockerSectionBaseImageRegistry) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ModelDockerSectionBaseImageRegistryPtrOutput struct{ *pulumi.OutputState }

func (ModelDockerSectionBaseImageRegistryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelDockerSectionBaseImageRegistry)(nil)).Elem()
}

func (o ModelDockerSectionBaseImageRegistryPtrOutput) ToModelDockerSectionBaseImageRegistryPtrOutput() ModelDockerSectionBaseImageRegistryPtrOutput {
	return o
}

func (o ModelDockerSectionBaseImageRegistryPtrOutput) ToModelDockerSectionBaseImageRegistryPtrOutputWithContext(ctx context.Context) ModelDockerSectionBaseImageRegistryPtrOutput {
	return o
}

func (o ModelDockerSectionBaseImageRegistryPtrOutput) Elem() ModelDockerSectionBaseImageRegistryOutput {
	return o.ApplyT(func(v *ModelDockerSectionBaseImageRegistry) ModelDockerSectionBaseImageRegistry {
		if v != nil {
			return *v
		}
		var ret ModelDockerSectionBaseImageRegistry
		return ret
	}).(ModelDockerSectionBaseImageRegistryOutput)
}

func (o ModelDockerSectionBaseImageRegistryPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelDockerSectionBaseImageRegistry) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o ModelDockerSectionBaseImageRegistryPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelDockerSectionBaseImageRegistry) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ModelDockerSectionBaseImageRegistryPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelDockerSectionBaseImageRegistry) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type ModelDockerSectionResponseResponseBaseImageRegistry struct {
	Address *string `pulumi:"address"`
}

type ModelEnvironmentDefinitionDocker struct {
	BaseDockerfile    *string                              `pulumi:"baseDockerfile"`
	BaseImage         *string                              `pulumi:"baseImage"`
	BaseImageRegistry *ModelDockerSectionBaseImageRegistry `pulumi:"baseImageRegistry"`
}





type ModelEnvironmentDefinitionDockerInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionDockerOutput() ModelEnvironmentDefinitionDockerOutput
	ToModelEnvironmentDefinitionDockerOutputWithContext(context.Context) ModelEnvironmentDefinitionDockerOutput
}

type ModelEnvironmentDefinitionDockerArgs struct {
	BaseDockerfile    pulumi.StringPtrInput                       `pulumi:"baseDockerfile"`
	BaseImage         pulumi.StringPtrInput                       `pulumi:"baseImage"`
	BaseImageRegistry ModelDockerSectionBaseImageRegistryPtrInput `pulumi:"baseImageRegistry"`
}

func (ModelEnvironmentDefinitionDockerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionDocker)(nil)).Elem()
}

func (i ModelEnvironmentDefinitionDockerArgs) ToModelEnvironmentDefinitionDockerOutput() ModelEnvironmentDefinitionDockerOutput {
	return i.ToModelEnvironmentDefinitionDockerOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionDockerArgs) ToModelEnvironmentDefinitionDockerOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionDockerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionDockerOutput)
}

func (i ModelEnvironmentDefinitionDockerArgs) ToModelEnvironmentDefinitionDockerPtrOutput() ModelEnvironmentDefinitionDockerPtrOutput {
	return i.ToModelEnvironmentDefinitionDockerPtrOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionDockerArgs) ToModelEnvironmentDefinitionDockerPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionDockerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionDockerOutput).ToModelEnvironmentDefinitionDockerPtrOutputWithContext(ctx)
}









type ModelEnvironmentDefinitionDockerPtrInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionDockerPtrOutput() ModelEnvironmentDefinitionDockerPtrOutput
	ToModelEnvironmentDefinitionDockerPtrOutputWithContext(context.Context) ModelEnvironmentDefinitionDockerPtrOutput
}

type modelEnvironmentDefinitionDockerPtrType ModelEnvironmentDefinitionDockerArgs

func ModelEnvironmentDefinitionDockerPtr(v *ModelEnvironmentDefinitionDockerArgs) ModelEnvironmentDefinitionDockerPtrInput {
	return (*modelEnvironmentDefinitionDockerPtrType)(v)
}

func (*modelEnvironmentDefinitionDockerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionDocker)(nil)).Elem()
}

func (i *modelEnvironmentDefinitionDockerPtrType) ToModelEnvironmentDefinitionDockerPtrOutput() ModelEnvironmentDefinitionDockerPtrOutput {
	return i.ToModelEnvironmentDefinitionDockerPtrOutputWithContext(context.Background())
}

func (i *modelEnvironmentDefinitionDockerPtrType) ToModelEnvironmentDefinitionDockerPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionDockerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionDockerPtrOutput)
}

type ModelEnvironmentDefinitionDockerOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionDockerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionDocker)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionDockerOutput) ToModelEnvironmentDefinitionDockerOutput() ModelEnvironmentDefinitionDockerOutput {
	return o
}

func (o ModelEnvironmentDefinitionDockerOutput) ToModelEnvironmentDefinitionDockerOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionDockerOutput {
	return o
}

func (o ModelEnvironmentDefinitionDockerOutput) ToModelEnvironmentDefinitionDockerPtrOutput() ModelEnvironmentDefinitionDockerPtrOutput {
	return o.ToModelEnvironmentDefinitionDockerPtrOutputWithContext(context.Background())
}

func (o ModelEnvironmentDefinitionDockerOutput) ToModelEnvironmentDefinitionDockerPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionDockerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelEnvironmentDefinitionDocker) *ModelEnvironmentDefinitionDocker {
		return &v
	}).(ModelEnvironmentDefinitionDockerPtrOutput)
}

func (o ModelEnvironmentDefinitionDockerOutput) BaseDockerfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionDocker) *string { return v.BaseDockerfile }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionDockerOutput) BaseImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionDocker) *string { return v.BaseImage }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionDockerOutput) BaseImageRegistry() ModelDockerSectionBaseImageRegistryPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionDocker) *ModelDockerSectionBaseImageRegistry {
		return v.BaseImageRegistry
	}).(ModelDockerSectionBaseImageRegistryPtrOutput)
}

type ModelEnvironmentDefinitionDockerPtrOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionDockerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionDocker)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionDockerPtrOutput) ToModelEnvironmentDefinitionDockerPtrOutput() ModelEnvironmentDefinitionDockerPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionDockerPtrOutput) ToModelEnvironmentDefinitionDockerPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionDockerPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionDockerPtrOutput) Elem() ModelEnvironmentDefinitionDockerOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionDocker) ModelEnvironmentDefinitionDocker {
		if v != nil {
			return *v
		}
		var ret ModelEnvironmentDefinitionDocker
		return ret
	}).(ModelEnvironmentDefinitionDockerOutput)
}

func (o ModelEnvironmentDefinitionDockerPtrOutput) BaseDockerfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionDocker) *string {
		if v == nil {
			return nil
		}
		return v.BaseDockerfile
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionDockerPtrOutput) BaseImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionDocker) *string {
		if v == nil {
			return nil
		}
		return v.BaseImage
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionDockerPtrOutput) BaseImageRegistry() ModelDockerSectionBaseImageRegistryPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionDocker) *ModelDockerSectionBaseImageRegistry {
		if v == nil {
			return nil
		}
		return v.BaseImageRegistry
	}).(ModelDockerSectionBaseImageRegistryPtrOutput)
}

type ModelEnvironmentDefinitionPython struct {
	BaseCondaEnvironment    *string     `pulumi:"baseCondaEnvironment"`
	CondaDependencies       interface{} `pulumi:"condaDependencies"`
	InterpreterPath         *string     `pulumi:"interpreterPath"`
	UserManagedDependencies *bool       `pulumi:"userManagedDependencies"`
}





type ModelEnvironmentDefinitionPythonInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionPythonOutput() ModelEnvironmentDefinitionPythonOutput
	ToModelEnvironmentDefinitionPythonOutputWithContext(context.Context) ModelEnvironmentDefinitionPythonOutput
}

type ModelEnvironmentDefinitionPythonArgs struct {
	BaseCondaEnvironment    pulumi.StringPtrInput `pulumi:"baseCondaEnvironment"`
	CondaDependencies       pulumi.Input          `pulumi:"condaDependencies"`
	InterpreterPath         pulumi.StringPtrInput `pulumi:"interpreterPath"`
	UserManagedDependencies pulumi.BoolPtrInput   `pulumi:"userManagedDependencies"`
}

func (ModelEnvironmentDefinitionPythonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionPython)(nil)).Elem()
}

func (i ModelEnvironmentDefinitionPythonArgs) ToModelEnvironmentDefinitionPythonOutput() ModelEnvironmentDefinitionPythonOutput {
	return i.ToModelEnvironmentDefinitionPythonOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionPythonArgs) ToModelEnvironmentDefinitionPythonOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionPythonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionPythonOutput)
}

func (i ModelEnvironmentDefinitionPythonArgs) ToModelEnvironmentDefinitionPythonPtrOutput() ModelEnvironmentDefinitionPythonPtrOutput {
	return i.ToModelEnvironmentDefinitionPythonPtrOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionPythonArgs) ToModelEnvironmentDefinitionPythonPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionPythonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionPythonOutput).ToModelEnvironmentDefinitionPythonPtrOutputWithContext(ctx)
}









type ModelEnvironmentDefinitionPythonPtrInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionPythonPtrOutput() ModelEnvironmentDefinitionPythonPtrOutput
	ToModelEnvironmentDefinitionPythonPtrOutputWithContext(context.Context) ModelEnvironmentDefinitionPythonPtrOutput
}

type modelEnvironmentDefinitionPythonPtrType ModelEnvironmentDefinitionPythonArgs

func ModelEnvironmentDefinitionPythonPtr(v *ModelEnvironmentDefinitionPythonArgs) ModelEnvironmentDefinitionPythonPtrInput {
	return (*modelEnvironmentDefinitionPythonPtrType)(v)
}

func (*modelEnvironmentDefinitionPythonPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionPython)(nil)).Elem()
}

func (i *modelEnvironmentDefinitionPythonPtrType) ToModelEnvironmentDefinitionPythonPtrOutput() ModelEnvironmentDefinitionPythonPtrOutput {
	return i.ToModelEnvironmentDefinitionPythonPtrOutputWithContext(context.Background())
}

func (i *modelEnvironmentDefinitionPythonPtrType) ToModelEnvironmentDefinitionPythonPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionPythonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionPythonPtrOutput)
}

type ModelEnvironmentDefinitionPythonOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionPythonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionPython)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionPythonOutput) ToModelEnvironmentDefinitionPythonOutput() ModelEnvironmentDefinitionPythonOutput {
	return o
}

func (o ModelEnvironmentDefinitionPythonOutput) ToModelEnvironmentDefinitionPythonOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionPythonOutput {
	return o
}

func (o ModelEnvironmentDefinitionPythonOutput) ToModelEnvironmentDefinitionPythonPtrOutput() ModelEnvironmentDefinitionPythonPtrOutput {
	return o.ToModelEnvironmentDefinitionPythonPtrOutputWithContext(context.Background())
}

func (o ModelEnvironmentDefinitionPythonOutput) ToModelEnvironmentDefinitionPythonPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionPythonPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelEnvironmentDefinitionPython) *ModelEnvironmentDefinitionPython {
		return &v
	}).(ModelEnvironmentDefinitionPythonPtrOutput)
}

func (o ModelEnvironmentDefinitionPythonOutput) BaseCondaEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionPython) *string { return v.BaseCondaEnvironment }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionPythonOutput) CondaDependencies() pulumi.AnyOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionPython) interface{} { return v.CondaDependencies }).(pulumi.AnyOutput)
}

func (o ModelEnvironmentDefinitionPythonOutput) InterpreterPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionPython) *string { return v.InterpreterPath }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionPythonOutput) UserManagedDependencies() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionPython) *bool { return v.UserManagedDependencies }).(pulumi.BoolPtrOutput)
}

type ModelEnvironmentDefinitionPythonPtrOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionPythonPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionPython)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionPythonPtrOutput) ToModelEnvironmentDefinitionPythonPtrOutput() ModelEnvironmentDefinitionPythonPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionPythonPtrOutput) ToModelEnvironmentDefinitionPythonPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionPythonPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionPythonPtrOutput) Elem() ModelEnvironmentDefinitionPythonOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionPython) ModelEnvironmentDefinitionPython {
		if v != nil {
			return *v
		}
		var ret ModelEnvironmentDefinitionPython
		return ret
	}).(ModelEnvironmentDefinitionPythonOutput)
}

func (o ModelEnvironmentDefinitionPythonPtrOutput) BaseCondaEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionPython) *string {
		if v == nil {
			return nil
		}
		return v.BaseCondaEnvironment
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionPythonPtrOutput) CondaDependencies() pulumi.AnyOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionPython) interface{} {
		if v == nil {
			return nil
		}
		return v.CondaDependencies
	}).(pulumi.AnyOutput)
}

func (o ModelEnvironmentDefinitionPythonPtrOutput) InterpreterPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionPython) *string {
		if v == nil {
			return nil
		}
		return v.InterpreterPath
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionPythonPtrOutput) UserManagedDependencies() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionPython) *bool {
		if v == nil {
			return nil
		}
		return v.UserManagedDependencies
	}).(pulumi.BoolPtrOutput)
}

type ModelEnvironmentDefinitionR struct {
	BioConductorPackages []string         `pulumi:"bioConductorPackages"`
	CranPackages         []RCranPackage   `pulumi:"cranPackages"`
	CustomUrlPackages    []string         `pulumi:"customUrlPackages"`
	GitHubPackages       []RGitHubPackage `pulumi:"gitHubPackages"`
	RVersion             *string          `pulumi:"rVersion"`
	RscriptPath          *string          `pulumi:"rscriptPath"`
	SnapshotDate         *string          `pulumi:"snapshotDate"`
	UserManaged          *bool            `pulumi:"userManaged"`
}





type ModelEnvironmentDefinitionRInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionROutput() ModelEnvironmentDefinitionROutput
	ToModelEnvironmentDefinitionROutputWithContext(context.Context) ModelEnvironmentDefinitionROutput
}

type ModelEnvironmentDefinitionRArgs struct {
	BioConductorPackages pulumi.StringArrayInput  `pulumi:"bioConductorPackages"`
	CranPackages         RCranPackageArrayInput   `pulumi:"cranPackages"`
	CustomUrlPackages    pulumi.StringArrayInput  `pulumi:"customUrlPackages"`
	GitHubPackages       RGitHubPackageArrayInput `pulumi:"gitHubPackages"`
	RVersion             pulumi.StringPtrInput    `pulumi:"rVersion"`
	RscriptPath          pulumi.StringPtrInput    `pulumi:"rscriptPath"`
	SnapshotDate         pulumi.StringPtrInput    `pulumi:"snapshotDate"`
	UserManaged          pulumi.BoolPtrInput      `pulumi:"userManaged"`
}

func (ModelEnvironmentDefinitionRArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionR)(nil)).Elem()
}

func (i ModelEnvironmentDefinitionRArgs) ToModelEnvironmentDefinitionROutput() ModelEnvironmentDefinitionROutput {
	return i.ToModelEnvironmentDefinitionROutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionRArgs) ToModelEnvironmentDefinitionROutputWithContext(ctx context.Context) ModelEnvironmentDefinitionROutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionROutput)
}

func (i ModelEnvironmentDefinitionRArgs) ToModelEnvironmentDefinitionRPtrOutput() ModelEnvironmentDefinitionRPtrOutput {
	return i.ToModelEnvironmentDefinitionRPtrOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionRArgs) ToModelEnvironmentDefinitionRPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionRPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionROutput).ToModelEnvironmentDefinitionRPtrOutputWithContext(ctx)
}









type ModelEnvironmentDefinitionRPtrInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionRPtrOutput() ModelEnvironmentDefinitionRPtrOutput
	ToModelEnvironmentDefinitionRPtrOutputWithContext(context.Context) ModelEnvironmentDefinitionRPtrOutput
}

type modelEnvironmentDefinitionRPtrType ModelEnvironmentDefinitionRArgs

func ModelEnvironmentDefinitionRPtr(v *ModelEnvironmentDefinitionRArgs) ModelEnvironmentDefinitionRPtrInput {
	return (*modelEnvironmentDefinitionRPtrType)(v)
}

func (*modelEnvironmentDefinitionRPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionR)(nil)).Elem()
}

func (i *modelEnvironmentDefinitionRPtrType) ToModelEnvironmentDefinitionRPtrOutput() ModelEnvironmentDefinitionRPtrOutput {
	return i.ToModelEnvironmentDefinitionRPtrOutputWithContext(context.Background())
}

func (i *modelEnvironmentDefinitionRPtrType) ToModelEnvironmentDefinitionRPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionRPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionRPtrOutput)
}

type ModelEnvironmentDefinitionROutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionROutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionR)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionROutput) ToModelEnvironmentDefinitionROutput() ModelEnvironmentDefinitionROutput {
	return o
}

func (o ModelEnvironmentDefinitionROutput) ToModelEnvironmentDefinitionROutputWithContext(ctx context.Context) ModelEnvironmentDefinitionROutput {
	return o
}

func (o ModelEnvironmentDefinitionROutput) ToModelEnvironmentDefinitionRPtrOutput() ModelEnvironmentDefinitionRPtrOutput {
	return o.ToModelEnvironmentDefinitionRPtrOutputWithContext(context.Background())
}

func (o ModelEnvironmentDefinitionROutput) ToModelEnvironmentDefinitionRPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionRPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelEnvironmentDefinitionR) *ModelEnvironmentDefinitionR {
		return &v
	}).(ModelEnvironmentDefinitionRPtrOutput)
}

func (o ModelEnvironmentDefinitionROutput) BioConductorPackages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionR) []string { return v.BioConductorPackages }).(pulumi.StringArrayOutput)
}

func (o ModelEnvironmentDefinitionROutput) CranPackages() RCranPackageArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionR) []RCranPackage { return v.CranPackages }).(RCranPackageArrayOutput)
}

func (o ModelEnvironmentDefinitionROutput) CustomUrlPackages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionR) []string { return v.CustomUrlPackages }).(pulumi.StringArrayOutput)
}

func (o ModelEnvironmentDefinitionROutput) GitHubPackages() RGitHubPackageArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionR) []RGitHubPackage { return v.GitHubPackages }).(RGitHubPackageArrayOutput)
}

func (o ModelEnvironmentDefinitionROutput) RVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionR) *string { return v.RVersion }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionROutput) RscriptPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionR) *string { return v.RscriptPath }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionROutput) SnapshotDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionR) *string { return v.SnapshotDate }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionROutput) UserManaged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionR) *bool { return v.UserManaged }).(pulumi.BoolPtrOutput)
}

type ModelEnvironmentDefinitionRPtrOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionRPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionR)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionRPtrOutput) ToModelEnvironmentDefinitionRPtrOutput() ModelEnvironmentDefinitionRPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionRPtrOutput) ToModelEnvironmentDefinitionRPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionRPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionRPtrOutput) Elem() ModelEnvironmentDefinitionROutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionR) ModelEnvironmentDefinitionR {
		if v != nil {
			return *v
		}
		var ret ModelEnvironmentDefinitionR
		return ret
	}).(ModelEnvironmentDefinitionROutput)
}

func (o ModelEnvironmentDefinitionRPtrOutput) BioConductorPackages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionR) []string {
		if v == nil {
			return nil
		}
		return v.BioConductorPackages
	}).(pulumi.StringArrayOutput)
}

func (o ModelEnvironmentDefinitionRPtrOutput) CranPackages() RCranPackageArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionR) []RCranPackage {
		if v == nil {
			return nil
		}
		return v.CranPackages
	}).(RCranPackageArrayOutput)
}

func (o ModelEnvironmentDefinitionRPtrOutput) CustomUrlPackages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionR) []string {
		if v == nil {
			return nil
		}
		return v.CustomUrlPackages
	}).(pulumi.StringArrayOutput)
}

func (o ModelEnvironmentDefinitionRPtrOutput) GitHubPackages() RGitHubPackageArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionR) []RGitHubPackage {
		if v == nil {
			return nil
		}
		return v.GitHubPackages
	}).(RGitHubPackageArrayOutput)
}

func (o ModelEnvironmentDefinitionRPtrOutput) RVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionR) *string {
		if v == nil {
			return nil
		}
		return v.RVersion
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionRPtrOutput) RscriptPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionR) *string {
		if v == nil {
			return nil
		}
		return v.RscriptPath
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionRPtrOutput) SnapshotDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionR) *string {
		if v == nil {
			return nil
		}
		return v.SnapshotDate
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionRPtrOutput) UserManaged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionR) *bool {
		if v == nil {
			return nil
		}
		return v.UserManaged
	}).(pulumi.BoolPtrOutput)
}

type ModelEnvironmentDefinitionResponseResponseDocker struct {
	BaseDockerfile    *string                                              `pulumi:"baseDockerfile"`
	BaseImage         *string                                              `pulumi:"baseImage"`
	BaseImageRegistry *ModelDockerSectionResponseResponseBaseImageRegistry `pulumi:"baseImageRegistry"`
}

type ModelEnvironmentDefinitionResponseResponsePython struct {
	BaseCondaEnvironment    *string     `pulumi:"baseCondaEnvironment"`
	CondaDependencies       interface{} `pulumi:"condaDependencies"`
	InterpreterPath         *string     `pulumi:"interpreterPath"`
	UserManagedDependencies *bool       `pulumi:"userManagedDependencies"`
}

type ModelEnvironmentDefinitionResponseResponseR struct {
	BioConductorPackages []string                         `pulumi:"bioConductorPackages"`
	CranPackages         []RCranPackageResponse           `pulumi:"cranPackages"`
	CustomUrlPackages    []string                         `pulumi:"customUrlPackages"`
	GitHubPackages       []RGitHubPackageResponseResponse `pulumi:"gitHubPackages"`
	RVersion             *string                          `pulumi:"rVersion"`
	RscriptPath          *string                          `pulumi:"rscriptPath"`
	SnapshotDate         *string                          `pulumi:"snapshotDate"`
	UserManaged          *bool                            `pulumi:"userManaged"`
}

type ModelEnvironmentDefinitionResponseResponseSpark struct {
	Packages         []SparkMavenPackageResponse `pulumi:"packages"`
	PrecachePackages *bool                       `pulumi:"precachePackages"`
	Repositories     []string                    `pulumi:"repositories"`
}

type ModelEnvironmentDefinitionSpark struct {
	Packages         []SparkMavenPackage `pulumi:"packages"`
	PrecachePackages *bool               `pulumi:"precachePackages"`
	Repositories     []string            `pulumi:"repositories"`
}





type ModelEnvironmentDefinitionSparkInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionSparkOutput() ModelEnvironmentDefinitionSparkOutput
	ToModelEnvironmentDefinitionSparkOutputWithContext(context.Context) ModelEnvironmentDefinitionSparkOutput
}

type ModelEnvironmentDefinitionSparkArgs struct {
	Packages         SparkMavenPackageArrayInput `pulumi:"packages"`
	PrecachePackages pulumi.BoolPtrInput         `pulumi:"precachePackages"`
	Repositories     pulumi.StringArrayInput     `pulumi:"repositories"`
}

func (ModelEnvironmentDefinitionSparkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionSpark)(nil)).Elem()
}

func (i ModelEnvironmentDefinitionSparkArgs) ToModelEnvironmentDefinitionSparkOutput() ModelEnvironmentDefinitionSparkOutput {
	return i.ToModelEnvironmentDefinitionSparkOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionSparkArgs) ToModelEnvironmentDefinitionSparkOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionSparkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionSparkOutput)
}

func (i ModelEnvironmentDefinitionSparkArgs) ToModelEnvironmentDefinitionSparkPtrOutput() ModelEnvironmentDefinitionSparkPtrOutput {
	return i.ToModelEnvironmentDefinitionSparkPtrOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionSparkArgs) ToModelEnvironmentDefinitionSparkPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionSparkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionSparkOutput).ToModelEnvironmentDefinitionSparkPtrOutputWithContext(ctx)
}









type ModelEnvironmentDefinitionSparkPtrInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionSparkPtrOutput() ModelEnvironmentDefinitionSparkPtrOutput
	ToModelEnvironmentDefinitionSparkPtrOutputWithContext(context.Context) ModelEnvironmentDefinitionSparkPtrOutput
}

type modelEnvironmentDefinitionSparkPtrType ModelEnvironmentDefinitionSparkArgs

func ModelEnvironmentDefinitionSparkPtr(v *ModelEnvironmentDefinitionSparkArgs) ModelEnvironmentDefinitionSparkPtrInput {
	return (*modelEnvironmentDefinitionSparkPtrType)(v)
}

func (*modelEnvironmentDefinitionSparkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionSpark)(nil)).Elem()
}

func (i *modelEnvironmentDefinitionSparkPtrType) ToModelEnvironmentDefinitionSparkPtrOutput() ModelEnvironmentDefinitionSparkPtrOutput {
	return i.ToModelEnvironmentDefinitionSparkPtrOutputWithContext(context.Background())
}

func (i *modelEnvironmentDefinitionSparkPtrType) ToModelEnvironmentDefinitionSparkPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionSparkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionSparkPtrOutput)
}

type ModelEnvironmentDefinitionSparkOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionSparkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionSpark)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionSparkOutput) ToModelEnvironmentDefinitionSparkOutput() ModelEnvironmentDefinitionSparkOutput {
	return o
}

func (o ModelEnvironmentDefinitionSparkOutput) ToModelEnvironmentDefinitionSparkOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionSparkOutput {
	return o
}

func (o ModelEnvironmentDefinitionSparkOutput) ToModelEnvironmentDefinitionSparkPtrOutput() ModelEnvironmentDefinitionSparkPtrOutput {
	return o.ToModelEnvironmentDefinitionSparkPtrOutputWithContext(context.Background())
}

func (o ModelEnvironmentDefinitionSparkOutput) ToModelEnvironmentDefinitionSparkPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionSparkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelEnvironmentDefinitionSpark) *ModelEnvironmentDefinitionSpark {
		return &v
	}).(ModelEnvironmentDefinitionSparkPtrOutput)
}

func (o ModelEnvironmentDefinitionSparkOutput) Packages() SparkMavenPackageArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionSpark) []SparkMavenPackage { return v.Packages }).(SparkMavenPackageArrayOutput)
}

func (o ModelEnvironmentDefinitionSparkOutput) PrecachePackages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionSpark) *bool { return v.PrecachePackages }).(pulumi.BoolPtrOutput)
}

func (o ModelEnvironmentDefinitionSparkOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionSpark) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

type ModelEnvironmentDefinitionSparkPtrOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionSparkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionSpark)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionSparkPtrOutput) ToModelEnvironmentDefinitionSparkPtrOutput() ModelEnvironmentDefinitionSparkPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionSparkPtrOutput) ToModelEnvironmentDefinitionSparkPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionSparkPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionSparkPtrOutput) Elem() ModelEnvironmentDefinitionSparkOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionSpark) ModelEnvironmentDefinitionSpark {
		if v != nil {
			return *v
		}
		var ret ModelEnvironmentDefinitionSpark
		return ret
	}).(ModelEnvironmentDefinitionSparkOutput)
}

func (o ModelEnvironmentDefinitionSparkPtrOutput) Packages() SparkMavenPackageArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionSpark) []SparkMavenPackage {
		if v == nil {
			return nil
		}
		return v.Packages
	}).(SparkMavenPackageArrayOutput)
}

func (o ModelEnvironmentDefinitionSparkPtrOutput) PrecachePackages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionSpark) *bool {
		if v == nil {
			return nil
		}
		return v.PrecachePackages
	}).(pulumi.BoolPtrOutput)
}

func (o ModelEnvironmentDefinitionSparkPtrOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionSpark) []string {
		if v == nil {
			return nil
		}
		return v.Repositories
	}).(pulumi.StringArrayOutput)
}

type ModelResponse struct {
	CreatedTime          *string                                `pulumi:"createdTime"`
	Datasets             []DatasetReferenceResponse             `pulumi:"datasets"`
	DerivedModelIds      []string                               `pulumi:"derivedModelIds"`
	Description          *string                                `pulumi:"description"`
	ExperimentName       *string                                `pulumi:"experimentName"`
	Framework            *string                                `pulumi:"framework"`
	FrameworkVersion     *string                                `pulumi:"frameworkVersion"`
	Id                   *string                                `pulumi:"id"`
	KvTags               map[string]string                      `pulumi:"kvTags"`
	MimeType             string                                 `pulumi:"mimeType"`
	ModifiedTime         *string                                `pulumi:"modifiedTime"`
	Name                 string                                 `pulumi:"name"`
	ParentModelId        *string                                `pulumi:"parentModelId"`
	Properties           map[string]string                      `pulumi:"properties"`
	ResourceRequirements *ContainerResourceRequirementsResponse `pulumi:"resourceRequirements"`
	RunId                *string                                `pulumi:"runId"`
	SampleInputData      *string                                `pulumi:"sampleInputData"`
	SampleOutputData     *string                                `pulumi:"sampleOutputData"`
	Unpack               *bool                                  `pulumi:"unpack"`
	Url                  string                                 `pulumi:"url"`
	Version              *float64                               `pulumi:"version"`
}

type NodeStateCountsResponse struct {
	IdleNodeCount      int `pulumi:"idleNodeCount"`
	LeavingNodeCount   int `pulumi:"leavingNodeCount"`
	PreemptedNodeCount int `pulumi:"preemptedNodeCount"`
	PreparingNodeCount int `pulumi:"preparingNodeCount"`
	RunningNodeCount   int `pulumi:"runningNodeCount"`
	UnusableNodeCount  int `pulumi:"unusableNodeCount"`
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
	Id string `pulumi:"id"`
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

type RCranPackage struct {
	Name       *string `pulumi:"name"`
	Repository *string `pulumi:"repository"`
}





type RCranPackageInput interface {
	pulumi.Input

	ToRCranPackageOutput() RCranPackageOutput
	ToRCranPackageOutputWithContext(context.Context) RCranPackageOutput
}

type RCranPackageArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Repository pulumi.StringPtrInput `pulumi:"repository"`
}

func (RCranPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RCranPackage)(nil)).Elem()
}

func (i RCranPackageArgs) ToRCranPackageOutput() RCranPackageOutput {
	return i.ToRCranPackageOutputWithContext(context.Background())
}

func (i RCranPackageArgs) ToRCranPackageOutputWithContext(ctx context.Context) RCranPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RCranPackageOutput)
}





type RCranPackageArrayInput interface {
	pulumi.Input

	ToRCranPackageArrayOutput() RCranPackageArrayOutput
	ToRCranPackageArrayOutputWithContext(context.Context) RCranPackageArrayOutput
}

type RCranPackageArray []RCranPackageInput

func (RCranPackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RCranPackage)(nil)).Elem()
}

func (i RCranPackageArray) ToRCranPackageArrayOutput() RCranPackageArrayOutput {
	return i.ToRCranPackageArrayOutputWithContext(context.Background())
}

func (i RCranPackageArray) ToRCranPackageArrayOutputWithContext(ctx context.Context) RCranPackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RCranPackageArrayOutput)
}

type RCranPackageOutput struct{ *pulumi.OutputState }

func (RCranPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RCranPackage)(nil)).Elem()
}

func (o RCranPackageOutput) ToRCranPackageOutput() RCranPackageOutput {
	return o
}

func (o RCranPackageOutput) ToRCranPackageOutputWithContext(ctx context.Context) RCranPackageOutput {
	return o
}

func (o RCranPackageOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RCranPackage) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RCranPackageOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RCranPackage) *string { return v.Repository }).(pulumi.StringPtrOutput)
}

type RCranPackageArrayOutput struct{ *pulumi.OutputState }

func (RCranPackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RCranPackage)(nil)).Elem()
}

func (o RCranPackageArrayOutput) ToRCranPackageArrayOutput() RCranPackageArrayOutput {
	return o
}

func (o RCranPackageArrayOutput) ToRCranPackageArrayOutputWithContext(ctx context.Context) RCranPackageArrayOutput {
	return o
}

func (o RCranPackageArrayOutput) Index(i pulumi.IntInput) RCranPackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RCranPackage {
		return vs[0].([]RCranPackage)[vs[1].(int)]
	}).(RCranPackageOutput)
}

type RCranPackageResponse struct {
	Name       *string `pulumi:"name"`
	Repository *string `pulumi:"repository"`
}

type RGitHubPackage struct {
	AuthToken  *string `pulumi:"authToken"`
	Repository *string `pulumi:"repository"`
}





type RGitHubPackageInput interface {
	pulumi.Input

	ToRGitHubPackageOutput() RGitHubPackageOutput
	ToRGitHubPackageOutputWithContext(context.Context) RGitHubPackageOutput
}

type RGitHubPackageArgs struct {
	AuthToken  pulumi.StringPtrInput `pulumi:"authToken"`
	Repository pulumi.StringPtrInput `pulumi:"repository"`
}

func (RGitHubPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RGitHubPackage)(nil)).Elem()
}

func (i RGitHubPackageArgs) ToRGitHubPackageOutput() RGitHubPackageOutput {
	return i.ToRGitHubPackageOutputWithContext(context.Background())
}

func (i RGitHubPackageArgs) ToRGitHubPackageOutputWithContext(ctx context.Context) RGitHubPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RGitHubPackageOutput)
}





type RGitHubPackageArrayInput interface {
	pulumi.Input

	ToRGitHubPackageArrayOutput() RGitHubPackageArrayOutput
	ToRGitHubPackageArrayOutputWithContext(context.Context) RGitHubPackageArrayOutput
}

type RGitHubPackageArray []RGitHubPackageInput

func (RGitHubPackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RGitHubPackage)(nil)).Elem()
}

func (i RGitHubPackageArray) ToRGitHubPackageArrayOutput() RGitHubPackageArrayOutput {
	return i.ToRGitHubPackageArrayOutputWithContext(context.Background())
}

func (i RGitHubPackageArray) ToRGitHubPackageArrayOutputWithContext(ctx context.Context) RGitHubPackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RGitHubPackageArrayOutput)
}

type RGitHubPackageOutput struct{ *pulumi.OutputState }

func (RGitHubPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RGitHubPackage)(nil)).Elem()
}

func (o RGitHubPackageOutput) ToRGitHubPackageOutput() RGitHubPackageOutput {
	return o
}

func (o RGitHubPackageOutput) ToRGitHubPackageOutputWithContext(ctx context.Context) RGitHubPackageOutput {
	return o
}

func (o RGitHubPackageOutput) AuthToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RGitHubPackage) *string { return v.AuthToken }).(pulumi.StringPtrOutput)
}

func (o RGitHubPackageOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RGitHubPackage) *string { return v.Repository }).(pulumi.StringPtrOutput)
}

type RGitHubPackageArrayOutput struct{ *pulumi.OutputState }

func (RGitHubPackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RGitHubPackage)(nil)).Elem()
}

func (o RGitHubPackageArrayOutput) ToRGitHubPackageArrayOutput() RGitHubPackageArrayOutput {
	return o
}

func (o RGitHubPackageArrayOutput) ToRGitHubPackageArrayOutputWithContext(ctx context.Context) RGitHubPackageArrayOutput {
	return o
}

func (o RGitHubPackageArrayOutput) Index(i pulumi.IntInput) RGitHubPackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RGitHubPackage {
		return vs[0].([]RGitHubPackage)[vs[1].(int)]
	}).(RGitHubPackageOutput)
}

type RGitHubPackageResponseResponse struct {
	Repository *string `pulumi:"repository"`
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

type ServiceResponseBaseResponseError struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
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

type SparkMavenPackage struct {
	Artifact *string `pulumi:"artifact"`
	Group    *string `pulumi:"group"`
	Version  *string `pulumi:"version"`
}





type SparkMavenPackageInput interface {
	pulumi.Input

	ToSparkMavenPackageOutput() SparkMavenPackageOutput
	ToSparkMavenPackageOutputWithContext(context.Context) SparkMavenPackageOutput
}

type SparkMavenPackageArgs struct {
	Artifact pulumi.StringPtrInput `pulumi:"artifact"`
	Group    pulumi.StringPtrInput `pulumi:"group"`
	Version  pulumi.StringPtrInput `pulumi:"version"`
}

func (SparkMavenPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkMavenPackage)(nil)).Elem()
}

func (i SparkMavenPackageArgs) ToSparkMavenPackageOutput() SparkMavenPackageOutput {
	return i.ToSparkMavenPackageOutputWithContext(context.Background())
}

func (i SparkMavenPackageArgs) ToSparkMavenPackageOutputWithContext(ctx context.Context) SparkMavenPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkMavenPackageOutput)
}





type SparkMavenPackageArrayInput interface {
	pulumi.Input

	ToSparkMavenPackageArrayOutput() SparkMavenPackageArrayOutput
	ToSparkMavenPackageArrayOutputWithContext(context.Context) SparkMavenPackageArrayOutput
}

type SparkMavenPackageArray []SparkMavenPackageInput

func (SparkMavenPackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SparkMavenPackage)(nil)).Elem()
}

func (i SparkMavenPackageArray) ToSparkMavenPackageArrayOutput() SparkMavenPackageArrayOutput {
	return i.ToSparkMavenPackageArrayOutputWithContext(context.Background())
}

func (i SparkMavenPackageArray) ToSparkMavenPackageArrayOutputWithContext(ctx context.Context) SparkMavenPackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkMavenPackageArrayOutput)
}

type SparkMavenPackageOutput struct{ *pulumi.OutputState }

func (SparkMavenPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkMavenPackage)(nil)).Elem()
}

func (o SparkMavenPackageOutput) ToSparkMavenPackageOutput() SparkMavenPackageOutput {
	return o
}

func (o SparkMavenPackageOutput) ToSparkMavenPackageOutputWithContext(ctx context.Context) SparkMavenPackageOutput {
	return o
}

func (o SparkMavenPackageOutput) Artifact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkMavenPackage) *string { return v.Artifact }).(pulumi.StringPtrOutput)
}

func (o SparkMavenPackageOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkMavenPackage) *string { return v.Group }).(pulumi.StringPtrOutput)
}

func (o SparkMavenPackageOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkMavenPackage) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type SparkMavenPackageArrayOutput struct{ *pulumi.OutputState }

func (SparkMavenPackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SparkMavenPackage)(nil)).Elem()
}

func (o SparkMavenPackageArrayOutput) ToSparkMavenPackageArrayOutput() SparkMavenPackageArrayOutput {
	return o
}

func (o SparkMavenPackageArrayOutput) ToSparkMavenPackageArrayOutputWithContext(ctx context.Context) SparkMavenPackageArrayOutput {
	return o
}

func (o SparkMavenPackageArrayOutput) Index(i pulumi.IntInput) SparkMavenPackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SparkMavenPackage {
		return vs[0].([]SparkMavenPackage)[vs[1].(int)]
	}).(SparkMavenPackageOutput)
}

type SparkMavenPackageResponse struct {
	Artifact *string `pulumi:"artifact"`
	Group    *string `pulumi:"group"`
	Version  *string `pulumi:"version"`
}

type SslConfiguration struct {
	Cert   *string `pulumi:"cert"`
	Cname  *string `pulumi:"cname"`
	Key    *string `pulumi:"key"`
	Status *string `pulumi:"status"`
}

type SslConfigurationResponse struct {
	Cert   *string `pulumi:"cert"`
	Cname  *string `pulumi:"cname"`
	Key    *string `pulumi:"key"`
	Status *string `pulumi:"status"`
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
	ComputeLocation *string                   `pulumi:"computeLocation"`
	ComputeType     string                    `pulumi:"computeType"`
	Description     *string                   `pulumi:"description"`
	Properties      *VirtualMachineProperties `pulumi:"properties"`
	ResourceId      *string                   `pulumi:"resourceId"`
}

type VirtualMachineImage struct {
	Id string `pulumi:"id"`
}

type VirtualMachineImageResponse struct {
	Id string `pulumi:"id"`
}

type VirtualMachineProperties struct {
	Address              *string                       `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	SshPort              *int                          `pulumi:"sshPort"`
	VirtualMachineSize   *string                       `pulumi:"virtualMachineSize"`
}

type VirtualMachineResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *VirtualMachineResponseProperties     `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type VirtualMachineResponseProperties struct {
	Address              *string                               `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	SshPort              *int                                  `pulumi:"sshPort"`
	VirtualMachineSize   *string                               `pulumi:"virtualMachineSize"`
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
	pulumi.RegisterOutputType(ACIServiceCreateRequestDataCollectionOutput{})
	pulumi.RegisterOutputType(ACIServiceCreateRequestDataCollectionPtrOutput{})
	pulumi.RegisterOutputType(ACIServiceCreateRequestEncryptionPropertiesOutput{})
	pulumi.RegisterOutputType(ACIServiceCreateRequestEncryptionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ACIServiceCreateRequestVnetConfigurationOutput{})
	pulumi.RegisterOutputType(ACIServiceCreateRequestVnetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestAutoScalerOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestAutoScalerPtrOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestDataCollectionOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestDataCollectionPtrOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestLivenessProbeRequirementsOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput{})
	pulumi.RegisterOutputType(ComputeBindingOutput{})
	pulumi.RegisterOutputType(ComputeBindingPtrOutput{})
	pulumi.RegisterOutputType(ComputeBindingResponseOutput{})
	pulumi.RegisterOutputType(ComputeBindingResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerResourceRequirementsOutput{})
	pulumi.RegisterOutputType(ContainerResourceRequirementsPtrOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestEnvironmentImageRequestOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestEnvironmentImageRequestPtrOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestKeysOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestKeysPtrOutput{})
	pulumi.RegisterOutputType(DatasetReferenceOutput{})
	pulumi.RegisterOutputType(DatasetReferenceArrayOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentReferenceOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentReferencePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageAssetOutput{})
	pulumi.RegisterOutputType(ImageAssetArrayOutput{})
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
	pulumi.RegisterOutputType(LabelingJobImagePropertiesOutput{})
	pulumi.RegisterOutputType(LabelingJobImagePropertiesPtrOutput{})
	pulumi.RegisterOutputType(LabelingJobImagePropertiesResponseOutput{})
	pulumi.RegisterOutputType(LabelingJobInstructionsOutput{})
	pulumi.RegisterOutputType(LabelingJobInstructionsPtrOutput{})
	pulumi.RegisterOutputType(LabelingJobInstructionsResponseOutput{})
	pulumi.RegisterOutputType(LabelingJobPropertiesOutput{})
	pulumi.RegisterOutputType(LabelingJobPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LabelingJobPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LinkedServicePropsOutput{})
	pulumi.RegisterOutputType(LinkedServicePropsPtrOutput{})
	pulumi.RegisterOutputType(LinkedServicePropsResponseOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationPtrOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationResponseOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ModelOutput{})
	pulumi.RegisterOutputType(ModelArrayOutput{})
	pulumi.RegisterOutputType(ModelDockerSectionBaseImageRegistryOutput{})
	pulumi.RegisterOutputType(ModelDockerSectionBaseImageRegistryPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionDockerOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionDockerPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionPythonOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionPythonPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionROutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionRPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionSparkOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionSparkPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ProgressMetricsResponseOutput{})
	pulumi.RegisterOutputType(RCranPackageOutput{})
	pulumi.RegisterOutputType(RCranPackageArrayOutput{})
	pulumi.RegisterOutputType(RGitHubPackageOutput{})
	pulumi.RegisterOutputType(RGitHubPackageArrayOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceArrayOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SparkMavenPackageOutput{})
	pulumi.RegisterOutputType(SparkMavenPackageArrayOutput{})
	pulumi.RegisterOutputType(StatusMessageResponseOutput{})
	pulumi.RegisterOutputType(StatusMessageResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}
