


package v20200501preview

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
	AgentVMSize                *string                     `pulumi:"agentVMSize"`
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
	AgentVMSize                *string                             `pulumi:"agentVMSize"`
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
	RemoteLoginPortPublicAccess *string                 `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings               *ScaleSettings          `pulumi:"scaleSettings"`
	Subnet                      *ResourceId             `pulumi:"subnet"`
	UserAccountCredentials      *UserAccountCredentials `pulumi:"userAccountCredentials"`
	VmPriority                  *string                 `pulumi:"vmPriority"`
	VmSize                      *string                 `pulumi:"vmSize"`
}


func (val *AmlComputeProperties) Defaults() *AmlComputeProperties {
	if val == nil {
		return nil
	}
	tmp := *val
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
	Errors                        []MachineLearningServiceErrorResponse `pulumi:"errors"`
	NodeStateCounts               NodeStateCountsResponse               `pulumi:"nodeStateCounts"`
	RemoteLoginPortPublicAccess   *string                               `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings                 *ScaleSettingsResponse                `pulumi:"scaleSettings"`
	Subnet                        *ResourceIdResponse                   `pulumi:"subnet"`
	TargetNodeCount               int                                   `pulumi:"targetNodeCount"`
	UserAccountCredentials        *UserAccountCredentialsResponse       `pulumi:"userAccountCredentials"`
	VmPriority                    *string                               `pulumi:"vmPriority"`
	VmSize                        *string                               `pulumi:"vmSize"`
}


func (val *AmlComputeResponseProperties) Defaults() *AmlComputeResponseProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RemoteLoginPortPublicAccess) {
		remoteLoginPortPublicAccess_ := "NotSpecified"
		tmp.RemoteLoginPortPublicAccess = &remoteLoginPortPublicAccess_
	}
	tmp.ScaleSettings = tmp.ScaleSettings.Defaults()

	return &tmp
}

type AzureDataLakeSectionResponse struct {
	AuthorityUrl                  *string `pulumi:"authorityUrl"`
	Certificate                   *string `pulumi:"certificate"`
	ClientId                      *string `pulumi:"clientId"`
	ClientSecret                  *string `pulumi:"clientSecret"`
	CredentialType                *string `pulumi:"credentialType"`
	IsCertAuth                    *bool   `pulumi:"isCertAuth"`
	ResourceGroup                 *string `pulumi:"resourceGroup"`
	ResourceUri                   *string `pulumi:"resourceUri"`
	ServiceDataAccessAuthIdentity *string `pulumi:"serviceDataAccessAuthIdentity"`
	StoreName                     *string `pulumi:"storeName"`
	SubscriptionId                *string `pulumi:"subscriptionId"`
	TenantId                      *string `pulumi:"tenantId"`
	Thumbprint                    *string `pulumi:"thumbprint"`
}

type AzureDataLakeSectionResponseOutput struct{ *pulumi.OutputState }

func (AzureDataLakeSectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDataLakeSectionResponse)(nil)).Elem()
}

func (o AzureDataLakeSectionResponseOutput) ToAzureDataLakeSectionResponseOutput() AzureDataLakeSectionResponseOutput {
	return o
}

func (o AzureDataLakeSectionResponseOutput) ToAzureDataLakeSectionResponseOutputWithContext(ctx context.Context) AzureDataLakeSectionResponseOutput {
	return o
}

func (o AzureDataLakeSectionResponseOutput) AuthorityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.AuthorityUrl }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) CredentialType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.CredentialType }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) IsCertAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *bool { return v.IsCertAuth }).(pulumi.BoolPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.ServiceDataAccessAuthIdentity }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.StoreName }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeSectionResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type AzureDataLakeSectionResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureDataLakeSectionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureDataLakeSectionResponse)(nil)).Elem()
}

func (o AzureDataLakeSectionResponsePtrOutput) ToAzureDataLakeSectionResponsePtrOutput() AzureDataLakeSectionResponsePtrOutput {
	return o
}

func (o AzureDataLakeSectionResponsePtrOutput) ToAzureDataLakeSectionResponsePtrOutputWithContext(ctx context.Context) AzureDataLakeSectionResponsePtrOutput {
	return o
}

func (o AzureDataLakeSectionResponsePtrOutput) Elem() AzureDataLakeSectionResponseOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) AzureDataLakeSectionResponse {
		if v != nil {
			return *v
		}
		var ret AzureDataLakeSectionResponse
		return ret
	}).(AzureDataLakeSectionResponseOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) AuthorityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorityUrl
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Certificate
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) CredentialType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.CredentialType
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) IsCertAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCertAuth
	}).(pulumi.BoolPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceUri
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceDataAccessAuthIdentity
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) StoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.StoreName
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeSectionResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDataLakeSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type AzureMySqlSectionResponse struct {
	AuthorityUrl                  *string `pulumi:"authorityUrl"`
	Certificate                   *string `pulumi:"certificate"`
	ClientId                      *string `pulumi:"clientId"`
	ClientSecret                  *string `pulumi:"clientSecret"`
	CredentialType                *string `pulumi:"credentialType"`
	DatabaseName                  *string `pulumi:"databaseName"`
	Endpoint                      *string `pulumi:"endpoint"`
	IsCertAuth                    *bool   `pulumi:"isCertAuth"`
	PortNumber                    *string `pulumi:"portNumber"`
	ResourceGroup                 *string `pulumi:"resourceGroup"`
	ResourceUri                   *string `pulumi:"resourceUri"`
	ServerName                    *string `pulumi:"serverName"`
	ServiceDataAccessAuthIdentity *string `pulumi:"serviceDataAccessAuthIdentity"`
	SubscriptionId                *string `pulumi:"subscriptionId"`
	TenantId                      *string `pulumi:"tenantId"`
	Thumbprint                    *string `pulumi:"thumbprint"`
	UserId                        *string `pulumi:"userId"`
	UserPassword                  *string `pulumi:"userPassword"`
}

type AzureMySqlSectionResponseOutput struct{ *pulumi.OutputState }

func (AzureMySqlSectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMySqlSectionResponse)(nil)).Elem()
}

func (o AzureMySqlSectionResponseOutput) ToAzureMySqlSectionResponseOutput() AzureMySqlSectionResponseOutput {
	return o
}

func (o AzureMySqlSectionResponseOutput) ToAzureMySqlSectionResponseOutputWithContext(ctx context.Context) AzureMySqlSectionResponseOutput {
	return o
}

func (o AzureMySqlSectionResponseOutput) AuthorityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.AuthorityUrl }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) CredentialType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.CredentialType }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) IsCertAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *bool { return v.IsCertAuth }).(pulumi.BoolPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) PortNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.PortNumber }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.ServiceDataAccessAuthIdentity }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.UserId }).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponseOutput) UserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMySqlSectionResponse) *string { return v.UserPassword }).(pulumi.StringPtrOutput)
}

type AzureMySqlSectionResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureMySqlSectionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMySqlSectionResponse)(nil)).Elem()
}

func (o AzureMySqlSectionResponsePtrOutput) ToAzureMySqlSectionResponsePtrOutput() AzureMySqlSectionResponsePtrOutput {
	return o
}

func (o AzureMySqlSectionResponsePtrOutput) ToAzureMySqlSectionResponsePtrOutputWithContext(ctx context.Context) AzureMySqlSectionResponsePtrOutput {
	return o
}

func (o AzureMySqlSectionResponsePtrOutput) Elem() AzureMySqlSectionResponseOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) AzureMySqlSectionResponse {
		if v != nil {
			return *v
		}
		var ret AzureMySqlSectionResponse
		return ret
	}).(AzureMySqlSectionResponseOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) AuthorityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorityUrl
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Certificate
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) CredentialType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.CredentialType
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.DatabaseName
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) IsCertAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCertAuth
	}).(pulumi.BoolPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) PortNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.PortNumber
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceUri
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerName
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceDataAccessAuthIdentity
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserId
	}).(pulumi.StringPtrOutput)
}

func (o AzureMySqlSectionResponsePtrOutput) UserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMySqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserPassword
	}).(pulumi.StringPtrOutput)
}

type AzurePostgreSqlSectionResponse struct {
	AuthorityUrl                  *string `pulumi:"authorityUrl"`
	Certificate                   *string `pulumi:"certificate"`
	ClientId                      *string `pulumi:"clientId"`
	ClientSecret                  *string `pulumi:"clientSecret"`
	CredentialType                *string `pulumi:"credentialType"`
	DatabaseName                  *string `pulumi:"databaseName"`
	EnableSsl                     *bool   `pulumi:"enableSsl"`
	Endpoint                      *string `pulumi:"endpoint"`
	IsCertAuth                    *bool   `pulumi:"isCertAuth"`
	PortNumber                    *string `pulumi:"portNumber"`
	ResourceGroup                 *string `pulumi:"resourceGroup"`
	ResourceUri                   *string `pulumi:"resourceUri"`
	ServerName                    *string `pulumi:"serverName"`
	ServiceDataAccessAuthIdentity *string `pulumi:"serviceDataAccessAuthIdentity"`
	SubscriptionId                *string `pulumi:"subscriptionId"`
	TenantId                      *string `pulumi:"tenantId"`
	Thumbprint                    *string `pulumi:"thumbprint"`
	UserId                        *string `pulumi:"userId"`
	UserPassword                  *string `pulumi:"userPassword"`
}

type AzurePostgreSqlSectionResponseOutput struct{ *pulumi.OutputState }

func (AzurePostgreSqlSectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzurePostgreSqlSectionResponse)(nil)).Elem()
}

func (o AzurePostgreSqlSectionResponseOutput) ToAzurePostgreSqlSectionResponseOutput() AzurePostgreSqlSectionResponseOutput {
	return o
}

func (o AzurePostgreSqlSectionResponseOutput) ToAzurePostgreSqlSectionResponseOutputWithContext(ctx context.Context) AzurePostgreSqlSectionResponseOutput {
	return o
}

func (o AzurePostgreSqlSectionResponseOutput) AuthorityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.AuthorityUrl }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) CredentialType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.CredentialType }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) EnableSsl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *bool { return v.EnableSsl }).(pulumi.BoolPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) IsCertAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *bool { return v.IsCertAuth }).(pulumi.BoolPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) PortNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.PortNumber }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.ServiceDataAccessAuthIdentity }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.UserId }).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponseOutput) UserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePostgreSqlSectionResponse) *string { return v.UserPassword }).(pulumi.StringPtrOutput)
}

type AzurePostgreSqlSectionResponsePtrOutput struct{ *pulumi.OutputState }

func (AzurePostgreSqlSectionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzurePostgreSqlSectionResponse)(nil)).Elem()
}

func (o AzurePostgreSqlSectionResponsePtrOutput) ToAzurePostgreSqlSectionResponsePtrOutput() AzurePostgreSqlSectionResponsePtrOutput {
	return o
}

func (o AzurePostgreSqlSectionResponsePtrOutput) ToAzurePostgreSqlSectionResponsePtrOutputWithContext(ctx context.Context) AzurePostgreSqlSectionResponsePtrOutput {
	return o
}

func (o AzurePostgreSqlSectionResponsePtrOutput) Elem() AzurePostgreSqlSectionResponseOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) AzurePostgreSqlSectionResponse {
		if v != nil {
			return *v
		}
		var ret AzurePostgreSqlSectionResponse
		return ret
	}).(AzurePostgreSqlSectionResponseOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) AuthorityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorityUrl
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Certificate
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) CredentialType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.CredentialType
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.DatabaseName
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) EnableSsl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableSsl
	}).(pulumi.BoolPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) IsCertAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCertAuth
	}).(pulumi.BoolPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) PortNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.PortNumber
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceUri
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerName
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceDataAccessAuthIdentity
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserId
	}).(pulumi.StringPtrOutput)
}

func (o AzurePostgreSqlSectionResponsePtrOutput) UserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePostgreSqlSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserPassword
	}).(pulumi.StringPtrOutput)
}

type AzureSqlDatabaseSectionResponse struct {
	AuthorityUrl                  *string `pulumi:"authorityUrl"`
	Certificate                   *string `pulumi:"certificate"`
	ClientId                      *string `pulumi:"clientId"`
	ClientSecret                  *string `pulumi:"clientSecret"`
	CredentialType                *string `pulumi:"credentialType"`
	DatabaseName                  *string `pulumi:"databaseName"`
	Endpoint                      *string `pulumi:"endpoint"`
	IsCertAuth                    *bool   `pulumi:"isCertAuth"`
	PortNumber                    *string `pulumi:"portNumber"`
	ResourceGroup                 *string `pulumi:"resourceGroup"`
	ResourceUri                   *string `pulumi:"resourceUri"`
	ServerName                    *string `pulumi:"serverName"`
	ServiceDataAccessAuthIdentity *string `pulumi:"serviceDataAccessAuthIdentity"`
	SubscriptionId                *string `pulumi:"subscriptionId"`
	TenantId                      *string `pulumi:"tenantId"`
	Thumbprint                    *string `pulumi:"thumbprint"`
	UserId                        *string `pulumi:"userId"`
	UserPassword                  *string `pulumi:"userPassword"`
}

type AzureSqlDatabaseSectionResponseOutput struct{ *pulumi.OutputState }

func (AzureSqlDatabaseSectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlDatabaseSectionResponse)(nil)).Elem()
}

func (o AzureSqlDatabaseSectionResponseOutput) ToAzureSqlDatabaseSectionResponseOutput() AzureSqlDatabaseSectionResponseOutput {
	return o
}

func (o AzureSqlDatabaseSectionResponseOutput) ToAzureSqlDatabaseSectionResponseOutputWithContext(ctx context.Context) AzureSqlDatabaseSectionResponseOutput {
	return o
}

func (o AzureSqlDatabaseSectionResponseOutput) AuthorityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.AuthorityUrl }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) CredentialType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.CredentialType }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) IsCertAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *bool { return v.IsCertAuth }).(pulumi.BoolPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) PortNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.PortNumber }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.ServiceDataAccessAuthIdentity }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.UserId }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponseOutput) UserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseSectionResponse) *string { return v.UserPassword }).(pulumi.StringPtrOutput)
}

type AzureSqlDatabaseSectionResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureSqlDatabaseSectionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSqlDatabaseSectionResponse)(nil)).Elem()
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) ToAzureSqlDatabaseSectionResponsePtrOutput() AzureSqlDatabaseSectionResponsePtrOutput {
	return o
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) ToAzureSqlDatabaseSectionResponsePtrOutputWithContext(ctx context.Context) AzureSqlDatabaseSectionResponsePtrOutput {
	return o
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) Elem() AzureSqlDatabaseSectionResponseOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) AzureSqlDatabaseSectionResponse {
		if v != nil {
			return *v
		}
		var ret AzureSqlDatabaseSectionResponse
		return ret
	}).(AzureSqlDatabaseSectionResponseOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) AuthorityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorityUrl
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Certificate
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) CredentialType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.CredentialType
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.DatabaseName
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) IsCertAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCertAuth
	}).(pulumi.BoolPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) PortNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.PortNumber
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceUri
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerName
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceDataAccessAuthIdentity
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserId
	}).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseSectionResponsePtrOutput) UserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSqlDatabaseSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserPassword
	}).(pulumi.StringPtrOutput)
}

type AzureStorageSectionResponse struct {
	AccountKey                           *string                    `pulumi:"accountKey"`
	AccountName                          *string                    `pulumi:"accountName"`
	AreWorkspaceManagedIdentitiesAllowed *bool                      `pulumi:"areWorkspaceManagedIdentitiesAllowed"`
	BlobCacheTimeout                     *int                       `pulumi:"blobCacheTimeout"`
	ClientCredentials                    *ClientCredentialsResponse `pulumi:"clientCredentials"`
	ContainerName                        *string                    `pulumi:"containerName"`
	Credential                           *string                    `pulumi:"credential"`
	CredentialType                       *string                    `pulumi:"credentialType"`
	Endpoint                             *string                    `pulumi:"endpoint"`
	IsSas                                *bool                      `pulumi:"isSas"`
	Protocol                             *string                    `pulumi:"protocol"`
	ResourceGroup                        *string                    `pulumi:"resourceGroup"`
	SasToken                             *string                    `pulumi:"sasToken"`
	ServiceDataAccessAuthIdentity        *string                    `pulumi:"serviceDataAccessAuthIdentity"`
	SubscriptionId                       *string                    `pulumi:"subscriptionId"`
}

type AzureStorageSectionResponseOutput struct{ *pulumi.OutputState }

func (AzureStorageSectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageSectionResponse)(nil)).Elem()
}

func (o AzureStorageSectionResponseOutput) ToAzureStorageSectionResponseOutput() AzureStorageSectionResponseOutput {
	return o
}

func (o AzureStorageSectionResponseOutput) ToAzureStorageSectionResponseOutputWithContext(ctx context.Context) AzureStorageSectionResponseOutput {
	return o
}

func (o AzureStorageSectionResponseOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponseOutput) AreWorkspaceManagedIdentitiesAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *bool { return v.AreWorkspaceManagedIdentitiesAllowed }).(pulumi.BoolPtrOutput)
}

func (o AzureStorageSectionResponseOutput) BlobCacheTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *int { return v.BlobCacheTimeout }).(pulumi.IntPtrOutput)
}

func (o AzureStorageSectionResponseOutput) ClientCredentials() ClientCredentialsResponsePtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *ClientCredentialsResponse { return v.ClientCredentials }).(ClientCredentialsResponsePtrOutput)
}

func (o AzureStorageSectionResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponseOutput) Credential() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *string { return v.Credential }).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponseOutput) CredentialType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *string { return v.CredentialType }).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponseOutput) IsSas() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *bool { return v.IsSas }).(pulumi.BoolPtrOutput)
}

func (o AzureStorageSectionResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponseOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *string { return v.SasToken }).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponseOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *string { return v.ServiceDataAccessAuthIdentity }).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageSectionResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type AzureStorageSectionResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureStorageSectionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStorageSectionResponse)(nil)).Elem()
}

func (o AzureStorageSectionResponsePtrOutput) ToAzureStorageSectionResponsePtrOutput() AzureStorageSectionResponsePtrOutput {
	return o
}

func (o AzureStorageSectionResponsePtrOutput) ToAzureStorageSectionResponsePtrOutputWithContext(ctx context.Context) AzureStorageSectionResponsePtrOutput {
	return o
}

func (o AzureStorageSectionResponsePtrOutput) Elem() AzureStorageSectionResponseOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) AzureStorageSectionResponse {
		if v != nil {
			return *v
		}
		var ret AzureStorageSectionResponse
		return ret
	}).(AzureStorageSectionResponseOutput)
}

func (o AzureStorageSectionResponsePtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) AreWorkspaceManagedIdentitiesAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AreWorkspaceManagedIdentitiesAllowed
	}).(pulumi.BoolPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) BlobCacheTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *int {
		if v == nil {
			return nil
		}
		return v.BlobCacheTimeout
	}).(pulumi.IntPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) ClientCredentials() ClientCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *ClientCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.ClientCredentials
	}).(ClientCredentialsResponsePtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContainerName
	}).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) Credential() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Credential
	}).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) CredentialType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.CredentialType
	}).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) IsSas() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsSas
	}).(pulumi.BoolPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Protocol
	}).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasToken
	}).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceDataAccessAuthIdentity
	}).(pulumi.StringPtrOutput)
}

func (o AzureStorageSectionResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStorageSectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

type ClientCredentialsResponse struct {
	AuthorityUrl                  *string `pulumi:"authorityUrl"`
	Certificate                   *string `pulumi:"certificate"`
	ClientId                      *string `pulumi:"clientId"`
	ClientSecret                  *string `pulumi:"clientSecret"`
	IsCertAuth                    *bool   `pulumi:"isCertAuth"`
	ResourceGroup                 *string `pulumi:"resourceGroup"`
	ResourceUri                   *string `pulumi:"resourceUri"`
	ServiceDataAccessAuthIdentity *string `pulumi:"serviceDataAccessAuthIdentity"`
	SubscriptionId                *string `pulumi:"subscriptionId"`
	TenantId                      *string `pulumi:"tenantId"`
	Thumbprint                    *string `pulumi:"thumbprint"`
}

type ClientCredentialsResponseOutput struct{ *pulumi.OutputState }

func (ClientCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCredentialsResponse)(nil)).Elem()
}

func (o ClientCredentialsResponseOutput) ToClientCredentialsResponseOutput() ClientCredentialsResponseOutput {
	return o
}

func (o ClientCredentialsResponseOutput) ToClientCredentialsResponseOutputWithContext(ctx context.Context) ClientCredentialsResponseOutput {
	return o
}

func (o ClientCredentialsResponseOutput) AuthorityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCredentialsResponse) *string { return v.AuthorityUrl }).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponseOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCredentialsResponse) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCredentialsResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCredentialsResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponseOutput) IsCertAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClientCredentialsResponse) *bool { return v.IsCertAuth }).(pulumi.BoolPtrOutput)
}

func (o ClientCredentialsResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCredentialsResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponseOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCredentialsResponse) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponseOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCredentialsResponse) *string { return v.ServiceDataAccessAuthIdentity }).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCredentialsResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCredentialsResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCredentialsResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ClientCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (ClientCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientCredentialsResponse)(nil)).Elem()
}

func (o ClientCredentialsResponsePtrOutput) ToClientCredentialsResponsePtrOutput() ClientCredentialsResponsePtrOutput {
	return o
}

func (o ClientCredentialsResponsePtrOutput) ToClientCredentialsResponsePtrOutputWithContext(ctx context.Context) ClientCredentialsResponsePtrOutput {
	return o
}

func (o ClientCredentialsResponsePtrOutput) Elem() ClientCredentialsResponseOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) ClientCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret ClientCredentialsResponse
		return ret
	}).(ClientCredentialsResponseOutput)
}

func (o ClientCredentialsResponsePtrOutput) AuthorityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorityUrl
	}).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponsePtrOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Certificate
	}).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponsePtrOutput) IsCertAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCertAuth
	}).(pulumi.BoolPtrOutput)
}

func (o ClientCredentialsResponsePtrOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponsePtrOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceUri
	}).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponsePtrOutput) ServiceDataAccessAuthIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceDataAccessAuthIdentity
	}).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ClientCredentialsResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
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

type DatasetCreateRequestDataPath struct {
	DatastoreName *string `pulumi:"datastoreName"`
	RelativePath  *string `pulumi:"relativePath"`
}





type DatasetCreateRequestDataPathInput interface {
	pulumi.Input

	ToDatasetCreateRequestDataPathOutput() DatasetCreateRequestDataPathOutput
	ToDatasetCreateRequestDataPathOutputWithContext(context.Context) DatasetCreateRequestDataPathOutput
}

type DatasetCreateRequestDataPathArgs struct {
	DatastoreName pulumi.StringPtrInput `pulumi:"datastoreName"`
	RelativePath  pulumi.StringPtrInput `pulumi:"relativePath"`
}

func (DatasetCreateRequestDataPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestDataPath)(nil)).Elem()
}

func (i DatasetCreateRequestDataPathArgs) ToDatasetCreateRequestDataPathOutput() DatasetCreateRequestDataPathOutput {
	return i.ToDatasetCreateRequestDataPathOutputWithContext(context.Background())
}

func (i DatasetCreateRequestDataPathArgs) ToDatasetCreateRequestDataPathOutputWithContext(ctx context.Context) DatasetCreateRequestDataPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestDataPathOutput)
}

func (i DatasetCreateRequestDataPathArgs) ToDatasetCreateRequestDataPathPtrOutput() DatasetCreateRequestDataPathPtrOutput {
	return i.ToDatasetCreateRequestDataPathPtrOutputWithContext(context.Background())
}

func (i DatasetCreateRequestDataPathArgs) ToDatasetCreateRequestDataPathPtrOutputWithContext(ctx context.Context) DatasetCreateRequestDataPathPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestDataPathOutput).ToDatasetCreateRequestDataPathPtrOutputWithContext(ctx)
}









type DatasetCreateRequestDataPathPtrInput interface {
	pulumi.Input

	ToDatasetCreateRequestDataPathPtrOutput() DatasetCreateRequestDataPathPtrOutput
	ToDatasetCreateRequestDataPathPtrOutputWithContext(context.Context) DatasetCreateRequestDataPathPtrOutput
}

type datasetCreateRequestDataPathPtrType DatasetCreateRequestDataPathArgs

func DatasetCreateRequestDataPathPtr(v *DatasetCreateRequestDataPathArgs) DatasetCreateRequestDataPathPtrInput {
	return (*datasetCreateRequestDataPathPtrType)(v)
}

func (*datasetCreateRequestDataPathPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetCreateRequestDataPath)(nil)).Elem()
}

func (i *datasetCreateRequestDataPathPtrType) ToDatasetCreateRequestDataPathPtrOutput() DatasetCreateRequestDataPathPtrOutput {
	return i.ToDatasetCreateRequestDataPathPtrOutputWithContext(context.Background())
}

func (i *datasetCreateRequestDataPathPtrType) ToDatasetCreateRequestDataPathPtrOutputWithContext(ctx context.Context) DatasetCreateRequestDataPathPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestDataPathPtrOutput)
}

type DatasetCreateRequestDataPathOutput struct{ *pulumi.OutputState }

func (DatasetCreateRequestDataPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestDataPath)(nil)).Elem()
}

func (o DatasetCreateRequestDataPathOutput) ToDatasetCreateRequestDataPathOutput() DatasetCreateRequestDataPathOutput {
	return o
}

func (o DatasetCreateRequestDataPathOutput) ToDatasetCreateRequestDataPathOutputWithContext(ctx context.Context) DatasetCreateRequestDataPathOutput {
	return o
}

func (o DatasetCreateRequestDataPathOutput) ToDatasetCreateRequestDataPathPtrOutput() DatasetCreateRequestDataPathPtrOutput {
	return o.ToDatasetCreateRequestDataPathPtrOutputWithContext(context.Background())
}

func (o DatasetCreateRequestDataPathOutput) ToDatasetCreateRequestDataPathPtrOutputWithContext(ctx context.Context) DatasetCreateRequestDataPathPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasetCreateRequestDataPath) *DatasetCreateRequestDataPath {
		return &v
	}).(DatasetCreateRequestDataPathPtrOutput)
}

func (o DatasetCreateRequestDataPathOutput) DatastoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestDataPath) *string { return v.DatastoreName }).(pulumi.StringPtrOutput)
}

func (o DatasetCreateRequestDataPathOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestDataPath) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

type DatasetCreateRequestDataPathPtrOutput struct{ *pulumi.OutputState }

func (DatasetCreateRequestDataPathPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetCreateRequestDataPath)(nil)).Elem()
}

func (o DatasetCreateRequestDataPathPtrOutput) ToDatasetCreateRequestDataPathPtrOutput() DatasetCreateRequestDataPathPtrOutput {
	return o
}

func (o DatasetCreateRequestDataPathPtrOutput) ToDatasetCreateRequestDataPathPtrOutputWithContext(ctx context.Context) DatasetCreateRequestDataPathPtrOutput {
	return o
}

func (o DatasetCreateRequestDataPathPtrOutput) Elem() DatasetCreateRequestDataPathOutput {
	return o.ApplyT(func(v *DatasetCreateRequestDataPath) DatasetCreateRequestDataPath {
		if v != nil {
			return *v
		}
		var ret DatasetCreateRequestDataPath
		return ret
	}).(DatasetCreateRequestDataPathOutput)
}

func (o DatasetCreateRequestDataPathPtrOutput) DatastoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetCreateRequestDataPath) *string {
		if v == nil {
			return nil
		}
		return v.DatastoreName
	}).(pulumi.StringPtrOutput)
}

func (o DatasetCreateRequestDataPathPtrOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetCreateRequestDataPath) *string {
		if v == nil {
			return nil
		}
		return v.RelativePath
	}).(pulumi.StringPtrOutput)
}

type DatasetCreateRequestParameters struct {
	Header          *string                    `pulumi:"header"`
	IncludePath     *bool                      `pulumi:"includePath"`
	PartitionFormat *string                    `pulumi:"partitionFormat"`
	Path            *DatasetCreateRequestPath  `pulumi:"path"`
	Query           *DatasetCreateRequestQuery `pulumi:"query"`
	Separator       *string                    `pulumi:"separator"`
	SourceType      *string                    `pulumi:"sourceType"`
}


func (val *DatasetCreateRequestParameters) Defaults() *DatasetCreateRequestParameters {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IncludePath) {
		includePath_ := false
		tmp.IncludePath = &includePath_
	}
	return &tmp
}





type DatasetCreateRequestParametersInput interface {
	pulumi.Input

	ToDatasetCreateRequestParametersOutput() DatasetCreateRequestParametersOutput
	ToDatasetCreateRequestParametersOutputWithContext(context.Context) DatasetCreateRequestParametersOutput
}

type DatasetCreateRequestParametersArgs struct {
	Header          pulumi.StringPtrInput             `pulumi:"header"`
	IncludePath     pulumi.BoolPtrInput               `pulumi:"includePath"`
	PartitionFormat pulumi.StringPtrInput             `pulumi:"partitionFormat"`
	Path            DatasetCreateRequestPathPtrInput  `pulumi:"path"`
	Query           DatasetCreateRequestQueryPtrInput `pulumi:"query"`
	Separator       pulumi.StringPtrInput             `pulumi:"separator"`
	SourceType      pulumi.StringPtrInput             `pulumi:"sourceType"`
}

func (DatasetCreateRequestParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestParameters)(nil)).Elem()
}

func (i DatasetCreateRequestParametersArgs) ToDatasetCreateRequestParametersOutput() DatasetCreateRequestParametersOutput {
	return i.ToDatasetCreateRequestParametersOutputWithContext(context.Background())
}

func (i DatasetCreateRequestParametersArgs) ToDatasetCreateRequestParametersOutputWithContext(ctx context.Context) DatasetCreateRequestParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestParametersOutput)
}

type DatasetCreateRequestParametersOutput struct{ *pulumi.OutputState }

func (DatasetCreateRequestParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestParameters)(nil)).Elem()
}

func (o DatasetCreateRequestParametersOutput) ToDatasetCreateRequestParametersOutput() DatasetCreateRequestParametersOutput {
	return o
}

func (o DatasetCreateRequestParametersOutput) ToDatasetCreateRequestParametersOutputWithContext(ctx context.Context) DatasetCreateRequestParametersOutput {
	return o
}

func (o DatasetCreateRequestParametersOutput) Header() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestParameters) *string { return v.Header }).(pulumi.StringPtrOutput)
}

func (o DatasetCreateRequestParametersOutput) IncludePath() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestParameters) *bool { return v.IncludePath }).(pulumi.BoolPtrOutput)
}

func (o DatasetCreateRequestParametersOutput) PartitionFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestParameters) *string { return v.PartitionFormat }).(pulumi.StringPtrOutput)
}

func (o DatasetCreateRequestParametersOutput) Path() DatasetCreateRequestPathPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestParameters) *DatasetCreateRequestPath { return v.Path }).(DatasetCreateRequestPathPtrOutput)
}

func (o DatasetCreateRequestParametersOutput) Query() DatasetCreateRequestQueryPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestParameters) *DatasetCreateRequestQuery { return v.Query }).(DatasetCreateRequestQueryPtrOutput)
}

func (o DatasetCreateRequestParametersOutput) Separator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestParameters) *string { return v.Separator }).(pulumi.StringPtrOutput)
}

func (o DatasetCreateRequestParametersOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestParameters) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

type DatasetCreateRequestPath struct {
	DataPath *DatasetCreateRequestDataPath `pulumi:"dataPath"`
	HttpUrl  *string                       `pulumi:"httpUrl"`
}





type DatasetCreateRequestPathInput interface {
	pulumi.Input

	ToDatasetCreateRequestPathOutput() DatasetCreateRequestPathOutput
	ToDatasetCreateRequestPathOutputWithContext(context.Context) DatasetCreateRequestPathOutput
}

type DatasetCreateRequestPathArgs struct {
	DataPath DatasetCreateRequestDataPathPtrInput `pulumi:"dataPath"`
	HttpUrl  pulumi.StringPtrInput                `pulumi:"httpUrl"`
}

func (DatasetCreateRequestPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestPath)(nil)).Elem()
}

func (i DatasetCreateRequestPathArgs) ToDatasetCreateRequestPathOutput() DatasetCreateRequestPathOutput {
	return i.ToDatasetCreateRequestPathOutputWithContext(context.Background())
}

func (i DatasetCreateRequestPathArgs) ToDatasetCreateRequestPathOutputWithContext(ctx context.Context) DatasetCreateRequestPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestPathOutput)
}

func (i DatasetCreateRequestPathArgs) ToDatasetCreateRequestPathPtrOutput() DatasetCreateRequestPathPtrOutput {
	return i.ToDatasetCreateRequestPathPtrOutputWithContext(context.Background())
}

func (i DatasetCreateRequestPathArgs) ToDatasetCreateRequestPathPtrOutputWithContext(ctx context.Context) DatasetCreateRequestPathPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestPathOutput).ToDatasetCreateRequestPathPtrOutputWithContext(ctx)
}









type DatasetCreateRequestPathPtrInput interface {
	pulumi.Input

	ToDatasetCreateRequestPathPtrOutput() DatasetCreateRequestPathPtrOutput
	ToDatasetCreateRequestPathPtrOutputWithContext(context.Context) DatasetCreateRequestPathPtrOutput
}

type datasetCreateRequestPathPtrType DatasetCreateRequestPathArgs

func DatasetCreateRequestPathPtr(v *DatasetCreateRequestPathArgs) DatasetCreateRequestPathPtrInput {
	return (*datasetCreateRequestPathPtrType)(v)
}

func (*datasetCreateRequestPathPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetCreateRequestPath)(nil)).Elem()
}

func (i *datasetCreateRequestPathPtrType) ToDatasetCreateRequestPathPtrOutput() DatasetCreateRequestPathPtrOutput {
	return i.ToDatasetCreateRequestPathPtrOutputWithContext(context.Background())
}

func (i *datasetCreateRequestPathPtrType) ToDatasetCreateRequestPathPtrOutputWithContext(ctx context.Context) DatasetCreateRequestPathPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestPathPtrOutput)
}

type DatasetCreateRequestPathOutput struct{ *pulumi.OutputState }

func (DatasetCreateRequestPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestPath)(nil)).Elem()
}

func (o DatasetCreateRequestPathOutput) ToDatasetCreateRequestPathOutput() DatasetCreateRequestPathOutput {
	return o
}

func (o DatasetCreateRequestPathOutput) ToDatasetCreateRequestPathOutputWithContext(ctx context.Context) DatasetCreateRequestPathOutput {
	return o
}

func (o DatasetCreateRequestPathOutput) ToDatasetCreateRequestPathPtrOutput() DatasetCreateRequestPathPtrOutput {
	return o.ToDatasetCreateRequestPathPtrOutputWithContext(context.Background())
}

func (o DatasetCreateRequestPathOutput) ToDatasetCreateRequestPathPtrOutputWithContext(ctx context.Context) DatasetCreateRequestPathPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasetCreateRequestPath) *DatasetCreateRequestPath {
		return &v
	}).(DatasetCreateRequestPathPtrOutput)
}

func (o DatasetCreateRequestPathOutput) DataPath() DatasetCreateRequestDataPathPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestPath) *DatasetCreateRequestDataPath { return v.DataPath }).(DatasetCreateRequestDataPathPtrOutput)
}

func (o DatasetCreateRequestPathOutput) HttpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestPath) *string { return v.HttpUrl }).(pulumi.StringPtrOutput)
}

type DatasetCreateRequestPathPtrOutput struct{ *pulumi.OutputState }

func (DatasetCreateRequestPathPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetCreateRequestPath)(nil)).Elem()
}

func (o DatasetCreateRequestPathPtrOutput) ToDatasetCreateRequestPathPtrOutput() DatasetCreateRequestPathPtrOutput {
	return o
}

func (o DatasetCreateRequestPathPtrOutput) ToDatasetCreateRequestPathPtrOutputWithContext(ctx context.Context) DatasetCreateRequestPathPtrOutput {
	return o
}

func (o DatasetCreateRequestPathPtrOutput) Elem() DatasetCreateRequestPathOutput {
	return o.ApplyT(func(v *DatasetCreateRequestPath) DatasetCreateRequestPath {
		if v != nil {
			return *v
		}
		var ret DatasetCreateRequestPath
		return ret
	}).(DatasetCreateRequestPathOutput)
}

func (o DatasetCreateRequestPathPtrOutput) DataPath() DatasetCreateRequestDataPathPtrOutput {
	return o.ApplyT(func(v *DatasetCreateRequestPath) *DatasetCreateRequestDataPath {
		if v == nil {
			return nil
		}
		return v.DataPath
	}).(DatasetCreateRequestDataPathPtrOutput)
}

func (o DatasetCreateRequestPathPtrOutput) HttpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetCreateRequestPath) *string {
		if v == nil {
			return nil
		}
		return v.HttpUrl
	}).(pulumi.StringPtrOutput)
}

type DatasetCreateRequestQuery struct {
	DatastoreName *string `pulumi:"datastoreName"`
	Query         *string `pulumi:"query"`
}





type DatasetCreateRequestQueryInput interface {
	pulumi.Input

	ToDatasetCreateRequestQueryOutput() DatasetCreateRequestQueryOutput
	ToDatasetCreateRequestQueryOutputWithContext(context.Context) DatasetCreateRequestQueryOutput
}

type DatasetCreateRequestQueryArgs struct {
	DatastoreName pulumi.StringPtrInput `pulumi:"datastoreName"`
	Query         pulumi.StringPtrInput `pulumi:"query"`
}

func (DatasetCreateRequestQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestQuery)(nil)).Elem()
}

func (i DatasetCreateRequestQueryArgs) ToDatasetCreateRequestQueryOutput() DatasetCreateRequestQueryOutput {
	return i.ToDatasetCreateRequestQueryOutputWithContext(context.Background())
}

func (i DatasetCreateRequestQueryArgs) ToDatasetCreateRequestQueryOutputWithContext(ctx context.Context) DatasetCreateRequestQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestQueryOutput)
}

func (i DatasetCreateRequestQueryArgs) ToDatasetCreateRequestQueryPtrOutput() DatasetCreateRequestQueryPtrOutput {
	return i.ToDatasetCreateRequestQueryPtrOutputWithContext(context.Background())
}

func (i DatasetCreateRequestQueryArgs) ToDatasetCreateRequestQueryPtrOutputWithContext(ctx context.Context) DatasetCreateRequestQueryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestQueryOutput).ToDatasetCreateRequestQueryPtrOutputWithContext(ctx)
}









type DatasetCreateRequestQueryPtrInput interface {
	pulumi.Input

	ToDatasetCreateRequestQueryPtrOutput() DatasetCreateRequestQueryPtrOutput
	ToDatasetCreateRequestQueryPtrOutputWithContext(context.Context) DatasetCreateRequestQueryPtrOutput
}

type datasetCreateRequestQueryPtrType DatasetCreateRequestQueryArgs

func DatasetCreateRequestQueryPtr(v *DatasetCreateRequestQueryArgs) DatasetCreateRequestQueryPtrInput {
	return (*datasetCreateRequestQueryPtrType)(v)
}

func (*datasetCreateRequestQueryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetCreateRequestQuery)(nil)).Elem()
}

func (i *datasetCreateRequestQueryPtrType) ToDatasetCreateRequestQueryPtrOutput() DatasetCreateRequestQueryPtrOutput {
	return i.ToDatasetCreateRequestQueryPtrOutputWithContext(context.Background())
}

func (i *datasetCreateRequestQueryPtrType) ToDatasetCreateRequestQueryPtrOutputWithContext(ctx context.Context) DatasetCreateRequestQueryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestQueryPtrOutput)
}

type DatasetCreateRequestQueryOutput struct{ *pulumi.OutputState }

func (DatasetCreateRequestQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestQuery)(nil)).Elem()
}

func (o DatasetCreateRequestQueryOutput) ToDatasetCreateRequestQueryOutput() DatasetCreateRequestQueryOutput {
	return o
}

func (o DatasetCreateRequestQueryOutput) ToDatasetCreateRequestQueryOutputWithContext(ctx context.Context) DatasetCreateRequestQueryOutput {
	return o
}

func (o DatasetCreateRequestQueryOutput) ToDatasetCreateRequestQueryPtrOutput() DatasetCreateRequestQueryPtrOutput {
	return o.ToDatasetCreateRequestQueryPtrOutputWithContext(context.Background())
}

func (o DatasetCreateRequestQueryOutput) ToDatasetCreateRequestQueryPtrOutputWithContext(ctx context.Context) DatasetCreateRequestQueryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasetCreateRequestQuery) *DatasetCreateRequestQuery {
		return &v
	}).(DatasetCreateRequestQueryPtrOutput)
}

func (o DatasetCreateRequestQueryOutput) DatastoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestQuery) *string { return v.DatastoreName }).(pulumi.StringPtrOutput)
}

func (o DatasetCreateRequestQueryOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestQuery) *string { return v.Query }).(pulumi.StringPtrOutput)
}

type DatasetCreateRequestQueryPtrOutput struct{ *pulumi.OutputState }

func (DatasetCreateRequestQueryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetCreateRequestQuery)(nil)).Elem()
}

func (o DatasetCreateRequestQueryPtrOutput) ToDatasetCreateRequestQueryPtrOutput() DatasetCreateRequestQueryPtrOutput {
	return o
}

func (o DatasetCreateRequestQueryPtrOutput) ToDatasetCreateRequestQueryPtrOutputWithContext(ctx context.Context) DatasetCreateRequestQueryPtrOutput {
	return o
}

func (o DatasetCreateRequestQueryPtrOutput) Elem() DatasetCreateRequestQueryOutput {
	return o.ApplyT(func(v *DatasetCreateRequestQuery) DatasetCreateRequestQuery {
		if v != nil {
			return *v
		}
		var ret DatasetCreateRequestQuery
		return ret
	}).(DatasetCreateRequestQueryOutput)
}

func (o DatasetCreateRequestQueryPtrOutput) DatastoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetCreateRequestQuery) *string {
		if v == nil {
			return nil
		}
		return v.DatastoreName
	}).(pulumi.StringPtrOutput)
}

func (o DatasetCreateRequestQueryPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetCreateRequestQuery) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

type DatasetCreateRequestRegistration struct {
	Description *string           `pulumi:"description"`
	Name        *string           `pulumi:"name"`
	Tags        map[string]string `pulumi:"tags"`
}





type DatasetCreateRequestRegistrationInput interface {
	pulumi.Input

	ToDatasetCreateRequestRegistrationOutput() DatasetCreateRequestRegistrationOutput
	ToDatasetCreateRequestRegistrationOutputWithContext(context.Context) DatasetCreateRequestRegistrationOutput
}

type DatasetCreateRequestRegistrationArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	Tags        pulumi.StringMapInput `pulumi:"tags"`
}

func (DatasetCreateRequestRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestRegistration)(nil)).Elem()
}

func (i DatasetCreateRequestRegistrationArgs) ToDatasetCreateRequestRegistrationOutput() DatasetCreateRequestRegistrationOutput {
	return i.ToDatasetCreateRequestRegistrationOutputWithContext(context.Background())
}

func (i DatasetCreateRequestRegistrationArgs) ToDatasetCreateRequestRegistrationOutputWithContext(ctx context.Context) DatasetCreateRequestRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestRegistrationOutput)
}

type DatasetCreateRequestRegistrationOutput struct{ *pulumi.OutputState }

func (DatasetCreateRequestRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestRegistration)(nil)).Elem()
}

func (o DatasetCreateRequestRegistrationOutput) ToDatasetCreateRequestRegistrationOutput() DatasetCreateRequestRegistrationOutput {
	return o
}

func (o DatasetCreateRequestRegistrationOutput) ToDatasetCreateRequestRegistrationOutputWithContext(ctx context.Context) DatasetCreateRequestRegistrationOutput {
	return o
}

func (o DatasetCreateRequestRegistrationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestRegistration) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DatasetCreateRequestRegistrationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestRegistration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DatasetCreateRequestRegistrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DatasetCreateRequestRegistration) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type DatasetCreateRequestTimeSeries struct {
	CoarseGrainTimestamp *string `pulumi:"coarseGrainTimestamp"`
	FineGrainTimestamp   *string `pulumi:"fineGrainTimestamp"`
}





type DatasetCreateRequestTimeSeriesInput interface {
	pulumi.Input

	ToDatasetCreateRequestTimeSeriesOutput() DatasetCreateRequestTimeSeriesOutput
	ToDatasetCreateRequestTimeSeriesOutputWithContext(context.Context) DatasetCreateRequestTimeSeriesOutput
}

type DatasetCreateRequestTimeSeriesArgs struct {
	CoarseGrainTimestamp pulumi.StringPtrInput `pulumi:"coarseGrainTimestamp"`
	FineGrainTimestamp   pulumi.StringPtrInput `pulumi:"fineGrainTimestamp"`
}

func (DatasetCreateRequestTimeSeriesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestTimeSeries)(nil)).Elem()
}

func (i DatasetCreateRequestTimeSeriesArgs) ToDatasetCreateRequestTimeSeriesOutput() DatasetCreateRequestTimeSeriesOutput {
	return i.ToDatasetCreateRequestTimeSeriesOutputWithContext(context.Background())
}

func (i DatasetCreateRequestTimeSeriesArgs) ToDatasetCreateRequestTimeSeriesOutputWithContext(ctx context.Context) DatasetCreateRequestTimeSeriesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestTimeSeriesOutput)
}

func (i DatasetCreateRequestTimeSeriesArgs) ToDatasetCreateRequestTimeSeriesPtrOutput() DatasetCreateRequestTimeSeriesPtrOutput {
	return i.ToDatasetCreateRequestTimeSeriesPtrOutputWithContext(context.Background())
}

func (i DatasetCreateRequestTimeSeriesArgs) ToDatasetCreateRequestTimeSeriesPtrOutputWithContext(ctx context.Context) DatasetCreateRequestTimeSeriesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestTimeSeriesOutput).ToDatasetCreateRequestTimeSeriesPtrOutputWithContext(ctx)
}









type DatasetCreateRequestTimeSeriesPtrInput interface {
	pulumi.Input

	ToDatasetCreateRequestTimeSeriesPtrOutput() DatasetCreateRequestTimeSeriesPtrOutput
	ToDatasetCreateRequestTimeSeriesPtrOutputWithContext(context.Context) DatasetCreateRequestTimeSeriesPtrOutput
}

type datasetCreateRequestTimeSeriesPtrType DatasetCreateRequestTimeSeriesArgs

func DatasetCreateRequestTimeSeriesPtr(v *DatasetCreateRequestTimeSeriesArgs) DatasetCreateRequestTimeSeriesPtrInput {
	return (*datasetCreateRequestTimeSeriesPtrType)(v)
}

func (*datasetCreateRequestTimeSeriesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetCreateRequestTimeSeries)(nil)).Elem()
}

func (i *datasetCreateRequestTimeSeriesPtrType) ToDatasetCreateRequestTimeSeriesPtrOutput() DatasetCreateRequestTimeSeriesPtrOutput {
	return i.ToDatasetCreateRequestTimeSeriesPtrOutputWithContext(context.Background())
}

func (i *datasetCreateRequestTimeSeriesPtrType) ToDatasetCreateRequestTimeSeriesPtrOutputWithContext(ctx context.Context) DatasetCreateRequestTimeSeriesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetCreateRequestTimeSeriesPtrOutput)
}

type DatasetCreateRequestTimeSeriesOutput struct{ *pulumi.OutputState }

func (DatasetCreateRequestTimeSeriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetCreateRequestTimeSeries)(nil)).Elem()
}

func (o DatasetCreateRequestTimeSeriesOutput) ToDatasetCreateRequestTimeSeriesOutput() DatasetCreateRequestTimeSeriesOutput {
	return o
}

func (o DatasetCreateRequestTimeSeriesOutput) ToDatasetCreateRequestTimeSeriesOutputWithContext(ctx context.Context) DatasetCreateRequestTimeSeriesOutput {
	return o
}

func (o DatasetCreateRequestTimeSeriesOutput) ToDatasetCreateRequestTimeSeriesPtrOutput() DatasetCreateRequestTimeSeriesPtrOutput {
	return o.ToDatasetCreateRequestTimeSeriesPtrOutputWithContext(context.Background())
}

func (o DatasetCreateRequestTimeSeriesOutput) ToDatasetCreateRequestTimeSeriesPtrOutputWithContext(ctx context.Context) DatasetCreateRequestTimeSeriesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasetCreateRequestTimeSeries) *DatasetCreateRequestTimeSeries {
		return &v
	}).(DatasetCreateRequestTimeSeriesPtrOutput)
}

func (o DatasetCreateRequestTimeSeriesOutput) CoarseGrainTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestTimeSeries) *string { return v.CoarseGrainTimestamp }).(pulumi.StringPtrOutput)
}

func (o DatasetCreateRequestTimeSeriesOutput) FineGrainTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetCreateRequestTimeSeries) *string { return v.FineGrainTimestamp }).(pulumi.StringPtrOutput)
}

type DatasetCreateRequestTimeSeriesPtrOutput struct{ *pulumi.OutputState }

func (DatasetCreateRequestTimeSeriesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetCreateRequestTimeSeries)(nil)).Elem()
}

func (o DatasetCreateRequestTimeSeriesPtrOutput) ToDatasetCreateRequestTimeSeriesPtrOutput() DatasetCreateRequestTimeSeriesPtrOutput {
	return o
}

func (o DatasetCreateRequestTimeSeriesPtrOutput) ToDatasetCreateRequestTimeSeriesPtrOutputWithContext(ctx context.Context) DatasetCreateRequestTimeSeriesPtrOutput {
	return o
}

func (o DatasetCreateRequestTimeSeriesPtrOutput) Elem() DatasetCreateRequestTimeSeriesOutput {
	return o.ApplyT(func(v *DatasetCreateRequestTimeSeries) DatasetCreateRequestTimeSeries {
		if v != nil {
			return *v
		}
		var ret DatasetCreateRequestTimeSeries
		return ret
	}).(DatasetCreateRequestTimeSeriesOutput)
}

func (o DatasetCreateRequestTimeSeriesPtrOutput) CoarseGrainTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetCreateRequestTimeSeries) *string {
		if v == nil {
			return nil
		}
		return v.CoarseGrainTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o DatasetCreateRequestTimeSeriesPtrOutput) FineGrainTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetCreateRequestTimeSeries) *string {
		if v == nil {
			return nil
		}
		return v.FineGrainTimestamp
	}).(pulumi.StringPtrOutput)
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

type DatasetResponse struct {
	CreatedTime    string                 `pulumi:"createdTime"`
	DatasetId      string                 `pulumi:"datasetId"`
	DatasetState   *DatasetStateResponse  `pulumi:"datasetState"`
	DatasetType    string                 `pulumi:"datasetType"`
	DefaultCompute string                 `pulumi:"defaultCompute"`
	Description    string                 `pulumi:"description"`
	Etag           string                 `pulumi:"etag"`
	IsVisible      bool                   `pulumi:"isVisible"`
	Latest         *DatasetResponseLatest `pulumi:"latest"`
	ModifiedTime   string                 `pulumi:"modifiedTime"`
	Name           string                 `pulumi:"name"`
	Tags           map[string]string      `pulumi:"tags"`
}

type DatasetResponseOutput struct{ *pulumi.OutputState }

func (DatasetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetResponse)(nil)).Elem()
}

func (o DatasetResponseOutput) ToDatasetResponseOutput() DatasetResponseOutput {
	return o
}

func (o DatasetResponseOutput) ToDatasetResponseOutputWithContext(ctx context.Context) DatasetResponseOutput {
	return o
}

func (o DatasetResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o DatasetResponseOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponse) string { return v.DatasetId }).(pulumi.StringOutput)
}

func (o DatasetResponseOutput) DatasetState() DatasetStateResponsePtrOutput {
	return o.ApplyT(func(v DatasetResponse) *DatasetStateResponse { return v.DatasetState }).(DatasetStateResponsePtrOutput)
}

func (o DatasetResponseOutput) DatasetType() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponse) string { return v.DatasetType }).(pulumi.StringOutput)
}

func (o DatasetResponseOutput) DefaultCompute() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponse) string { return v.DefaultCompute }).(pulumi.StringOutput)
}

func (o DatasetResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o DatasetResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o DatasetResponseOutput) IsVisible() pulumi.BoolOutput {
	return o.ApplyT(func(v DatasetResponse) bool { return v.IsVisible }).(pulumi.BoolOutput)
}

func (o DatasetResponseOutput) Latest() DatasetResponseLatestPtrOutput {
	return o.ApplyT(func(v DatasetResponse) *DatasetResponseLatest { return v.Latest }).(DatasetResponseLatestPtrOutput)
}

func (o DatasetResponseOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponse) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o DatasetResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DatasetResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DatasetResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type DatasetResponseDataPath struct {
	AdditionalProperties       map[string]interface{}      `pulumi:"additionalProperties"`
	AzureFilePath              string                      `pulumi:"azureFilePath"`
	DatastoreName              string                      `pulumi:"datastoreName"`
	HttpUrl                    string                      `pulumi:"httpUrl"`
	PartitionFormat            string                      `pulumi:"partitionFormat"`
	PartitionFormatIgnoreError bool                        `pulumi:"partitionFormatIgnoreError"`
	Paths                      []string                    `pulumi:"paths"`
	RelativePath               string                      `pulumi:"relativePath"`
	SqlDataPath                *DatasetResponseSqlDataPath `pulumi:"sqlDataPath"`
}

type DatasetResponseDataPathOutput struct{ *pulumi.OutputState }

func (DatasetResponseDataPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetResponseDataPath)(nil)).Elem()
}

func (o DatasetResponseDataPathOutput) ToDatasetResponseDataPathOutput() DatasetResponseDataPathOutput {
	return o
}

func (o DatasetResponseDataPathOutput) ToDatasetResponseDataPathOutputWithContext(ctx context.Context) DatasetResponseDataPathOutput {
	return o
}

func (o DatasetResponseDataPathOutput) AdditionalProperties() pulumi.MapOutput {
	return o.ApplyT(func(v DatasetResponseDataPath) map[string]interface{} { return v.AdditionalProperties }).(pulumi.MapOutput)
}

func (o DatasetResponseDataPathOutput) AzureFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseDataPath) string { return v.AzureFilePath }).(pulumi.StringOutput)
}

func (o DatasetResponseDataPathOutput) DatastoreName() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseDataPath) string { return v.DatastoreName }).(pulumi.StringOutput)
}

func (o DatasetResponseDataPathOutput) HttpUrl() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseDataPath) string { return v.HttpUrl }).(pulumi.StringOutput)
}

func (o DatasetResponseDataPathOutput) PartitionFormat() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseDataPath) string { return v.PartitionFormat }).(pulumi.StringOutput)
}

func (o DatasetResponseDataPathOutput) PartitionFormatIgnoreError() pulumi.BoolOutput {
	return o.ApplyT(func(v DatasetResponseDataPath) bool { return v.PartitionFormatIgnoreError }).(pulumi.BoolOutput)
}

func (o DatasetResponseDataPathOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatasetResponseDataPath) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o DatasetResponseDataPathOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseDataPath) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o DatasetResponseDataPathOutput) SqlDataPath() DatasetResponseSqlDataPathPtrOutput {
	return o.ApplyT(func(v DatasetResponseDataPath) *DatasetResponseSqlDataPath { return v.SqlDataPath }).(DatasetResponseSqlDataPathPtrOutput)
}

type DatasetResponseDataPathPtrOutput struct{ *pulumi.OutputState }

func (DatasetResponseDataPathPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetResponseDataPath)(nil)).Elem()
}

func (o DatasetResponseDataPathPtrOutput) ToDatasetResponseDataPathPtrOutput() DatasetResponseDataPathPtrOutput {
	return o
}

func (o DatasetResponseDataPathPtrOutput) ToDatasetResponseDataPathPtrOutputWithContext(ctx context.Context) DatasetResponseDataPathPtrOutput {
	return o
}

func (o DatasetResponseDataPathPtrOutput) Elem() DatasetResponseDataPathOutput {
	return o.ApplyT(func(v *DatasetResponseDataPath) DatasetResponseDataPath {
		if v != nil {
			return *v
		}
		var ret DatasetResponseDataPath
		return ret
	}).(DatasetResponseDataPathOutput)
}

func (o DatasetResponseDataPathPtrOutput) AdditionalProperties() pulumi.MapOutput {
	return o.ApplyT(func(v *DatasetResponseDataPath) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.AdditionalProperties
	}).(pulumi.MapOutput)
}

func (o DatasetResponseDataPathPtrOutput) AzureFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseDataPath) *string {
		if v == nil {
			return nil
		}
		return &v.AzureFilePath
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseDataPathPtrOutput) DatastoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseDataPath) *string {
		if v == nil {
			return nil
		}
		return &v.DatastoreName
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseDataPathPtrOutput) HttpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseDataPath) *string {
		if v == nil {
			return nil
		}
		return &v.HttpUrl
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseDataPathPtrOutput) PartitionFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseDataPath) *string {
		if v == nil {
			return nil
		}
		return &v.PartitionFormat
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseDataPathPtrOutput) PartitionFormatIgnoreError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatasetResponseDataPath) *bool {
		if v == nil {
			return nil
		}
		return &v.PartitionFormatIgnoreError
	}).(pulumi.BoolPtrOutput)
}

func (o DatasetResponseDataPathPtrOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatasetResponseDataPath) []string {
		if v == nil {
			return nil
		}
		return v.Paths
	}).(pulumi.StringArrayOutput)
}

func (o DatasetResponseDataPathPtrOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseDataPath) *string {
		if v == nil {
			return nil
		}
		return &v.RelativePath
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseDataPathPtrOutput) SqlDataPath() DatasetResponseSqlDataPathPtrOutput {
	return o.ApplyT(func(v *DatasetResponseDataPath) *DatasetResponseSqlDataPath {
		if v == nil {
			return nil
		}
		return v.SqlDataPath
	}).(DatasetResponseSqlDataPathPtrOutput)
}

type DatasetResponseLatest struct {
	CreatedBy                        *UserInfoResponse        `pulumi:"createdBy"`
	CreatedTime                      string                   `pulumi:"createdTime"`
	DataPath                         *DatasetResponseDataPath `pulumi:"dataPath"`
	Dataflow                         string                   `pulumi:"dataflow"`
	DatasetDefinitionState           *DatasetStateResponse    `pulumi:"datasetDefinitionState"`
	DatasetId                        string                   `pulumi:"datasetId"`
	Description                      string                   `pulumi:"description"`
	Etag                             string                   `pulumi:"etag"`
	FileType                         string                   `pulumi:"fileType"`
	ModifiedTime                     string                   `pulumi:"modifiedTime"`
	Notes                            string                   `pulumi:"notes"`
	PartitionFormatInPath            bool                     `pulumi:"partitionFormatInPath"`
	Properties                       map[string]interface{}   `pulumi:"properties"`
	SavedDatasetId                   string                   `pulumi:"savedDatasetId"`
	Tags                             map[string]string        `pulumi:"tags"`
	TelemetryInfo                    map[string]string        `pulumi:"telemetryInfo"`
	UseDescriptionTagsFromDefinition bool                     `pulumi:"useDescriptionTagsFromDefinition"`
	VersionId                        string                   `pulumi:"versionId"`
}

type DatasetResponseLatestOutput struct{ *pulumi.OutputState }

func (DatasetResponseLatestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetResponseLatest)(nil)).Elem()
}

func (o DatasetResponseLatestOutput) ToDatasetResponseLatestOutput() DatasetResponseLatestOutput {
	return o
}

func (o DatasetResponseLatestOutput) ToDatasetResponseLatestOutputWithContext(ctx context.Context) DatasetResponseLatestOutput {
	return o
}

func (o DatasetResponseLatestOutput) CreatedBy() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v DatasetResponseLatest) *UserInfoResponse { return v.CreatedBy }).(UserInfoResponsePtrOutput)
}

func (o DatasetResponseLatestOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseLatest) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o DatasetResponseLatestOutput) DataPath() DatasetResponseDataPathPtrOutput {
	return o.ApplyT(func(v DatasetResponseLatest) *DatasetResponseDataPath { return v.DataPath }).(DatasetResponseDataPathPtrOutput)
}

func (o DatasetResponseLatestOutput) Dataflow() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseLatest) string { return v.Dataflow }).(pulumi.StringOutput)
}

func (o DatasetResponseLatestOutput) DatasetDefinitionState() DatasetStateResponsePtrOutput {
	return o.ApplyT(func(v DatasetResponseLatest) *DatasetStateResponse { return v.DatasetDefinitionState }).(DatasetStateResponsePtrOutput)
}

func (o DatasetResponseLatestOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseLatest) string { return v.DatasetId }).(pulumi.StringOutput)
}

func (o DatasetResponseLatestOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseLatest) string { return v.Description }).(pulumi.StringOutput)
}

func (o DatasetResponseLatestOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseLatest) string { return v.Etag }).(pulumi.StringOutput)
}

func (o DatasetResponseLatestOutput) FileType() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseLatest) string { return v.FileType }).(pulumi.StringOutput)
}

func (o DatasetResponseLatestOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseLatest) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o DatasetResponseLatestOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseLatest) string { return v.Notes }).(pulumi.StringOutput)
}

func (o DatasetResponseLatestOutput) PartitionFormatInPath() pulumi.BoolOutput {
	return o.ApplyT(func(v DatasetResponseLatest) bool { return v.PartitionFormatInPath }).(pulumi.BoolOutput)
}

func (o DatasetResponseLatestOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func(v DatasetResponseLatest) map[string]interface{} { return v.Properties }).(pulumi.MapOutput)
}

func (o DatasetResponseLatestOutput) SavedDatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseLatest) string { return v.SavedDatasetId }).(pulumi.StringOutput)
}

func (o DatasetResponseLatestOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DatasetResponseLatest) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatasetResponseLatestOutput) TelemetryInfo() pulumi.StringMapOutput {
	return o.ApplyT(func(v DatasetResponseLatest) map[string]string { return v.TelemetryInfo }).(pulumi.StringMapOutput)
}

func (o DatasetResponseLatestOutput) UseDescriptionTagsFromDefinition() pulumi.BoolOutput {
	return o.ApplyT(func(v DatasetResponseLatest) bool { return v.UseDescriptionTagsFromDefinition }).(pulumi.BoolOutput)
}

func (o DatasetResponseLatestOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseLatest) string { return v.VersionId }).(pulumi.StringOutput)
}

type DatasetResponseLatestPtrOutput struct{ *pulumi.OutputState }

func (DatasetResponseLatestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetResponseLatest)(nil)).Elem()
}

func (o DatasetResponseLatestPtrOutput) ToDatasetResponseLatestPtrOutput() DatasetResponseLatestPtrOutput {
	return o
}

func (o DatasetResponseLatestPtrOutput) ToDatasetResponseLatestPtrOutputWithContext(ctx context.Context) DatasetResponseLatestPtrOutput {
	return o
}

func (o DatasetResponseLatestPtrOutput) Elem() DatasetResponseLatestOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) DatasetResponseLatest {
		if v != nil {
			return *v
		}
		var ret DatasetResponseLatest
		return ret
	}).(DatasetResponseLatestOutput)
}

func (o DatasetResponseLatestPtrOutput) CreatedBy() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *UserInfoResponse {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(UserInfoResponsePtrOutput)
}

func (o DatasetResponseLatestPtrOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedTime
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) DataPath() DatasetResponseDataPathPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *DatasetResponseDataPath {
		if v == nil {
			return nil
		}
		return v.DataPath
	}).(DatasetResponseDataPathPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) Dataflow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *string {
		if v == nil {
			return nil
		}
		return &v.Dataflow
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) DatasetDefinitionState() DatasetStateResponsePtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *DatasetStateResponse {
		if v == nil {
			return nil
		}
		return v.DatasetDefinitionState
	}).(DatasetStateResponsePtrOutput)
}

func (o DatasetResponseLatestPtrOutput) DatasetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *string {
		if v == nil {
			return nil
		}
		return &v.DatasetId
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) FileType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *string {
		if v == nil {
			return nil
		}
		return &v.FileType
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) ModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *string {
		if v == nil {
			return nil
		}
		return &v.ModifiedTime
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *string {
		if v == nil {
			return nil
		}
		return &v.Notes
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) PartitionFormatInPath() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *bool {
		if v == nil {
			return nil
		}
		return &v.PartitionFormatInPath
	}).(pulumi.BoolPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.MapOutput)
}

func (o DatasetResponseLatestPtrOutput) SavedDatasetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *string {
		if v == nil {
			return nil
		}
		return &v.SavedDatasetId
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o DatasetResponseLatestPtrOutput) TelemetryInfo() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) map[string]string {
		if v == nil {
			return nil
		}
		return v.TelemetryInfo
	}).(pulumi.StringMapOutput)
}

func (o DatasetResponseLatestPtrOutput) UseDescriptionTagsFromDefinition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *bool {
		if v == nil {
			return nil
		}
		return &v.UseDescriptionTagsFromDefinition
	}).(pulumi.BoolPtrOutput)
}

func (o DatasetResponseLatestPtrOutput) VersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseLatest) *string {
		if v == nil {
			return nil
		}
		return &v.VersionId
	}).(pulumi.StringPtrOutput)
}

type DatasetResponseSqlDataPath struct {
	QueryTimeout           float64 `pulumi:"queryTimeout"`
	SqlQuery               string  `pulumi:"sqlQuery"`
	SqlStoredProcedureName string  `pulumi:"sqlStoredProcedureName"`
	SqlTableName           string  `pulumi:"sqlTableName"`
}

type DatasetResponseSqlDataPathOutput struct{ *pulumi.OutputState }

func (DatasetResponseSqlDataPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetResponseSqlDataPath)(nil)).Elem()
}

func (o DatasetResponseSqlDataPathOutput) ToDatasetResponseSqlDataPathOutput() DatasetResponseSqlDataPathOutput {
	return o
}

func (o DatasetResponseSqlDataPathOutput) ToDatasetResponseSqlDataPathOutputWithContext(ctx context.Context) DatasetResponseSqlDataPathOutput {
	return o
}

func (o DatasetResponseSqlDataPathOutput) QueryTimeout() pulumi.Float64Output {
	return o.ApplyT(func(v DatasetResponseSqlDataPath) float64 { return v.QueryTimeout }).(pulumi.Float64Output)
}

func (o DatasetResponseSqlDataPathOutput) SqlQuery() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseSqlDataPath) string { return v.SqlQuery }).(pulumi.StringOutput)
}

func (o DatasetResponseSqlDataPathOutput) SqlStoredProcedureName() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseSqlDataPath) string { return v.SqlStoredProcedureName }).(pulumi.StringOutput)
}

func (o DatasetResponseSqlDataPathOutput) SqlTableName() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetResponseSqlDataPath) string { return v.SqlTableName }).(pulumi.StringOutput)
}

type DatasetResponseSqlDataPathPtrOutput struct{ *pulumi.OutputState }

func (DatasetResponseSqlDataPathPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetResponseSqlDataPath)(nil)).Elem()
}

func (o DatasetResponseSqlDataPathPtrOutput) ToDatasetResponseSqlDataPathPtrOutput() DatasetResponseSqlDataPathPtrOutput {
	return o
}

func (o DatasetResponseSqlDataPathPtrOutput) ToDatasetResponseSqlDataPathPtrOutputWithContext(ctx context.Context) DatasetResponseSqlDataPathPtrOutput {
	return o
}

func (o DatasetResponseSqlDataPathPtrOutput) Elem() DatasetResponseSqlDataPathOutput {
	return o.ApplyT(func(v *DatasetResponseSqlDataPath) DatasetResponseSqlDataPath {
		if v != nil {
			return *v
		}
		var ret DatasetResponseSqlDataPath
		return ret
	}).(DatasetResponseSqlDataPathOutput)
}

func (o DatasetResponseSqlDataPathPtrOutput) QueryTimeout() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DatasetResponseSqlDataPath) *float64 {
		if v == nil {
			return nil
		}
		return &v.QueryTimeout
	}).(pulumi.Float64PtrOutput)
}

func (o DatasetResponseSqlDataPathPtrOutput) SqlQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseSqlDataPath) *string {
		if v == nil {
			return nil
		}
		return &v.SqlQuery
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseSqlDataPathPtrOutput) SqlStoredProcedureName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseSqlDataPath) *string {
		if v == nil {
			return nil
		}
		return &v.SqlStoredProcedureName
	}).(pulumi.StringPtrOutput)
}

func (o DatasetResponseSqlDataPathPtrOutput) SqlTableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetResponseSqlDataPath) *string {
		if v == nil {
			return nil
		}
		return &v.SqlTableName
	}).(pulumi.StringPtrOutput)
}

type DatasetStateResponse struct {
	DeprecatedBy *DatasetStateResponseDeprecatedBy `pulumi:"deprecatedBy"`
	Etag         string                            `pulumi:"etag"`
	State        *string                           `pulumi:"state"`
}

type DatasetStateResponseOutput struct{ *pulumi.OutputState }

func (DatasetStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetStateResponse)(nil)).Elem()
}

func (o DatasetStateResponseOutput) ToDatasetStateResponseOutput() DatasetStateResponseOutput {
	return o
}

func (o DatasetStateResponseOutput) ToDatasetStateResponseOutputWithContext(ctx context.Context) DatasetStateResponseOutput {
	return o
}

func (o DatasetStateResponseOutput) DeprecatedBy() DatasetStateResponseDeprecatedByPtrOutput {
	return o.ApplyT(func(v DatasetStateResponse) *DatasetStateResponseDeprecatedBy { return v.DeprecatedBy }).(DatasetStateResponseDeprecatedByPtrOutput)
}

func (o DatasetStateResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetStateResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o DatasetStateResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetStateResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type DatasetStateResponsePtrOutput struct{ *pulumi.OutputState }

func (DatasetStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetStateResponse)(nil)).Elem()
}

func (o DatasetStateResponsePtrOutput) ToDatasetStateResponsePtrOutput() DatasetStateResponsePtrOutput {
	return o
}

func (o DatasetStateResponsePtrOutput) ToDatasetStateResponsePtrOutputWithContext(ctx context.Context) DatasetStateResponsePtrOutput {
	return o
}

func (o DatasetStateResponsePtrOutput) Elem() DatasetStateResponseOutput {
	return o.ApplyT(func(v *DatasetStateResponse) DatasetStateResponse {
		if v != nil {
			return *v
		}
		var ret DatasetStateResponse
		return ret
	}).(DatasetStateResponseOutput)
}

func (o DatasetStateResponsePtrOutput) DeprecatedBy() DatasetStateResponseDeprecatedByPtrOutput {
	return o.ApplyT(func(v *DatasetStateResponse) *DatasetStateResponseDeprecatedBy {
		if v == nil {
			return nil
		}
		return v.DeprecatedBy
	}).(DatasetStateResponseDeprecatedByPtrOutput)
}

func (o DatasetStateResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetStateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o DatasetStateResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type DatasetStateResponseDeprecatedBy struct {
	DatasetId         string  `pulumi:"datasetId"`
	DefinitionVersion *string `pulumi:"definitionVersion"`
}

type DatasetStateResponseDeprecatedByOutput struct{ *pulumi.OutputState }

func (DatasetStateResponseDeprecatedByOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetStateResponseDeprecatedBy)(nil)).Elem()
}

func (o DatasetStateResponseDeprecatedByOutput) ToDatasetStateResponseDeprecatedByOutput() DatasetStateResponseDeprecatedByOutput {
	return o
}

func (o DatasetStateResponseDeprecatedByOutput) ToDatasetStateResponseDeprecatedByOutputWithContext(ctx context.Context) DatasetStateResponseDeprecatedByOutput {
	return o
}

func (o DatasetStateResponseDeprecatedByOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v DatasetStateResponseDeprecatedBy) string { return v.DatasetId }).(pulumi.StringOutput)
}

func (o DatasetStateResponseDeprecatedByOutput) DefinitionVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetStateResponseDeprecatedBy) *string { return v.DefinitionVersion }).(pulumi.StringPtrOutput)
}

type DatasetStateResponseDeprecatedByPtrOutput struct{ *pulumi.OutputState }

func (DatasetStateResponseDeprecatedByPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetStateResponseDeprecatedBy)(nil)).Elem()
}

func (o DatasetStateResponseDeprecatedByPtrOutput) ToDatasetStateResponseDeprecatedByPtrOutput() DatasetStateResponseDeprecatedByPtrOutput {
	return o
}

func (o DatasetStateResponseDeprecatedByPtrOutput) ToDatasetStateResponseDeprecatedByPtrOutputWithContext(ctx context.Context) DatasetStateResponseDeprecatedByPtrOutput {
	return o
}

func (o DatasetStateResponseDeprecatedByPtrOutput) Elem() DatasetStateResponseDeprecatedByOutput {
	return o.ApplyT(func(v *DatasetStateResponseDeprecatedBy) DatasetStateResponseDeprecatedBy {
		if v != nil {
			return *v
		}
		var ret DatasetStateResponseDeprecatedBy
		return ret
	}).(DatasetStateResponseDeprecatedByOutput)
}

func (o DatasetStateResponseDeprecatedByPtrOutput) DatasetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetStateResponseDeprecatedBy) *string {
		if v == nil {
			return nil
		}
		return &v.DatasetId
	}).(pulumi.StringPtrOutput)
}

func (o DatasetStateResponseDeprecatedByPtrOutput) DefinitionVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatasetStateResponseDeprecatedBy) *string {
		if v == nil {
			return nil
		}
		return v.DefinitionVersion
	}).(pulumi.StringPtrOutput)
}

type DatastoreResponse struct {
	AzureDataLakeSection    *AzureDataLakeSectionResponse    `pulumi:"azureDataLakeSection"`
	AzureMySqlSection       *AzureMySqlSectionResponse       `pulumi:"azureMySqlSection"`
	AzurePostgreSqlSection  *AzurePostgreSqlSectionResponse  `pulumi:"azurePostgreSqlSection"`
	AzureSqlDatabaseSection *AzureSqlDatabaseSectionResponse `pulumi:"azureSqlDatabaseSection"`
	AzureStorageSection     *AzureStorageSectionResponse     `pulumi:"azureStorageSection"`
	CreatedBy               UserInfoResponse                 `pulumi:"createdBy"`
	CreatedTime             string                           `pulumi:"createdTime"`
	DataStoreType           *string                          `pulumi:"dataStoreType"`
	Description             *string                          `pulumi:"description"`
	GlusterFsSection        *GlusterFsSectionResponse        `pulumi:"glusterFsSection"`
	HasBeenValidated        *bool                            `pulumi:"hasBeenValidated"`
	LinkedInfo              *LinkedInfoResponse              `pulumi:"linkedInfo"`
	ModifiedBy              UserInfoResponse                 `pulumi:"modifiedBy"`
	ModifiedTime            string                           `pulumi:"modifiedTime"`
	Name                    *string                          `pulumi:"name"`
	Tags                    map[string]string                `pulumi:"tags"`
}


func (val *DatastoreResponse) Defaults() *DatastoreResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HasBeenValidated) {
		hasBeenValidated_ := false
		tmp.HasBeenValidated = &hasBeenValidated_
	}
	return &tmp
}

type DatastoreResponseOutput struct{ *pulumi.OutputState }

func (DatastoreResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatastoreResponse)(nil)).Elem()
}

func (o DatastoreResponseOutput) ToDatastoreResponseOutput() DatastoreResponseOutput {
	return o
}

func (o DatastoreResponseOutput) ToDatastoreResponseOutputWithContext(ctx context.Context) DatastoreResponseOutput {
	return o
}

func (o DatastoreResponseOutput) AzureDataLakeSection() AzureDataLakeSectionResponsePtrOutput {
	return o.ApplyT(func(v DatastoreResponse) *AzureDataLakeSectionResponse { return v.AzureDataLakeSection }).(AzureDataLakeSectionResponsePtrOutput)
}

func (o DatastoreResponseOutput) AzureMySqlSection() AzureMySqlSectionResponsePtrOutput {
	return o.ApplyT(func(v DatastoreResponse) *AzureMySqlSectionResponse { return v.AzureMySqlSection }).(AzureMySqlSectionResponsePtrOutput)
}

func (o DatastoreResponseOutput) AzurePostgreSqlSection() AzurePostgreSqlSectionResponsePtrOutput {
	return o.ApplyT(func(v DatastoreResponse) *AzurePostgreSqlSectionResponse { return v.AzurePostgreSqlSection }).(AzurePostgreSqlSectionResponsePtrOutput)
}

func (o DatastoreResponseOutput) AzureSqlDatabaseSection() AzureSqlDatabaseSectionResponsePtrOutput {
	return o.ApplyT(func(v DatastoreResponse) *AzureSqlDatabaseSectionResponse { return v.AzureSqlDatabaseSection }).(AzureSqlDatabaseSectionResponsePtrOutput)
}

func (o DatastoreResponseOutput) AzureStorageSection() AzureStorageSectionResponsePtrOutput {
	return o.ApplyT(func(v DatastoreResponse) *AzureStorageSectionResponse { return v.AzureStorageSection }).(AzureStorageSectionResponsePtrOutput)
}

func (o DatastoreResponseOutput) CreatedBy() UserInfoResponseOutput {
	return o.ApplyT(func(v DatastoreResponse) UserInfoResponse { return v.CreatedBy }).(UserInfoResponseOutput)
}

func (o DatastoreResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v DatastoreResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o DatastoreResponseOutput) DataStoreType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatastoreResponse) *string { return v.DataStoreType }).(pulumi.StringPtrOutput)
}

func (o DatastoreResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatastoreResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DatastoreResponseOutput) GlusterFsSection() GlusterFsSectionResponsePtrOutput {
	return o.ApplyT(func(v DatastoreResponse) *GlusterFsSectionResponse { return v.GlusterFsSection }).(GlusterFsSectionResponsePtrOutput)
}

func (o DatastoreResponseOutput) HasBeenValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DatastoreResponse) *bool { return v.HasBeenValidated }).(pulumi.BoolPtrOutput)
}

func (o DatastoreResponseOutput) LinkedInfo() LinkedInfoResponsePtrOutput {
	return o.ApplyT(func(v DatastoreResponse) *LinkedInfoResponse { return v.LinkedInfo }).(LinkedInfoResponsePtrOutput)
}

func (o DatastoreResponseOutput) ModifiedBy() UserInfoResponseOutput {
	return o.ApplyT(func(v DatastoreResponse) UserInfoResponse { return v.ModifiedBy }).(UserInfoResponseOutput)
}

func (o DatastoreResponseOutput) ModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v DatastoreResponse) string { return v.ModifiedTime }).(pulumi.StringOutput)
}

func (o DatastoreResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatastoreResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DatastoreResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v DatastoreResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
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

type GlusterFsSectionResponse struct {
	ServerAddress string `pulumi:"serverAddress"`
	VolumeName    string `pulumi:"volumeName"`
}

type GlusterFsSectionResponseOutput struct{ *pulumi.OutputState }

func (GlusterFsSectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlusterFsSectionResponse)(nil)).Elem()
}

func (o GlusterFsSectionResponseOutput) ToGlusterFsSectionResponseOutput() GlusterFsSectionResponseOutput {
	return o
}

func (o GlusterFsSectionResponseOutput) ToGlusterFsSectionResponseOutputWithContext(ctx context.Context) GlusterFsSectionResponseOutput {
	return o
}

func (o GlusterFsSectionResponseOutput) ServerAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GlusterFsSectionResponse) string { return v.ServerAddress }).(pulumi.StringOutput)
}

func (o GlusterFsSectionResponseOutput) VolumeName() pulumi.StringOutput {
	return o.ApplyT(func(v GlusterFsSectionResponse) string { return v.VolumeName }).(pulumi.StringOutput)
}

type GlusterFsSectionResponsePtrOutput struct{ *pulumi.OutputState }

func (GlusterFsSectionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlusterFsSectionResponse)(nil)).Elem()
}

func (o GlusterFsSectionResponsePtrOutput) ToGlusterFsSectionResponsePtrOutput() GlusterFsSectionResponsePtrOutput {
	return o
}

func (o GlusterFsSectionResponsePtrOutput) ToGlusterFsSectionResponsePtrOutputWithContext(ctx context.Context) GlusterFsSectionResponsePtrOutput {
	return o
}

func (o GlusterFsSectionResponsePtrOutput) Elem() GlusterFsSectionResponseOutput {
	return o.ApplyT(func(v *GlusterFsSectionResponse) GlusterFsSectionResponse {
		if v != nil {
			return *v
		}
		var ret GlusterFsSectionResponse
		return ret
	}).(GlusterFsSectionResponseOutput)
}

func (o GlusterFsSectionResponsePtrOutput) ServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlusterFsSectionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServerAddress
	}).(pulumi.StringPtrOutput)
}

func (o GlusterFsSectionResponsePtrOutput) VolumeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlusterFsSectionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VolumeName
	}).(pulumi.StringPtrOutput)
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

type LinkedWorkspaceProps struct {
	LinkedWorkspaceResourceId      *string `pulumi:"linkedWorkspaceResourceId"`
	UserAssignedIdentityResourceId *string `pulumi:"userAssignedIdentityResourceId"`
}





type LinkedWorkspacePropsInput interface {
	pulumi.Input

	ToLinkedWorkspacePropsOutput() LinkedWorkspacePropsOutput
	ToLinkedWorkspacePropsOutputWithContext(context.Context) LinkedWorkspacePropsOutput
}

type LinkedWorkspacePropsArgs struct {
	LinkedWorkspaceResourceId      pulumi.StringPtrInput `pulumi:"linkedWorkspaceResourceId"`
	UserAssignedIdentityResourceId pulumi.StringPtrInput `pulumi:"userAssignedIdentityResourceId"`
}

func (LinkedWorkspacePropsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedWorkspaceProps)(nil)).Elem()
}

func (i LinkedWorkspacePropsArgs) ToLinkedWorkspacePropsOutput() LinkedWorkspacePropsOutput {
	return i.ToLinkedWorkspacePropsOutputWithContext(context.Background())
}

func (i LinkedWorkspacePropsArgs) ToLinkedWorkspacePropsOutputWithContext(ctx context.Context) LinkedWorkspacePropsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedWorkspacePropsOutput)
}

func (i LinkedWorkspacePropsArgs) ToLinkedWorkspacePropsPtrOutput() LinkedWorkspacePropsPtrOutput {
	return i.ToLinkedWorkspacePropsPtrOutputWithContext(context.Background())
}

func (i LinkedWorkspacePropsArgs) ToLinkedWorkspacePropsPtrOutputWithContext(ctx context.Context) LinkedWorkspacePropsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedWorkspacePropsOutput).ToLinkedWorkspacePropsPtrOutputWithContext(ctx)
}









type LinkedWorkspacePropsPtrInput interface {
	pulumi.Input

	ToLinkedWorkspacePropsPtrOutput() LinkedWorkspacePropsPtrOutput
	ToLinkedWorkspacePropsPtrOutputWithContext(context.Context) LinkedWorkspacePropsPtrOutput
}

type linkedWorkspacePropsPtrType LinkedWorkspacePropsArgs

func LinkedWorkspacePropsPtr(v *LinkedWorkspacePropsArgs) LinkedWorkspacePropsPtrInput {
	return (*linkedWorkspacePropsPtrType)(v)
}

func (*linkedWorkspacePropsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedWorkspaceProps)(nil)).Elem()
}

func (i *linkedWorkspacePropsPtrType) ToLinkedWorkspacePropsPtrOutput() LinkedWorkspacePropsPtrOutput {
	return i.ToLinkedWorkspacePropsPtrOutputWithContext(context.Background())
}

func (i *linkedWorkspacePropsPtrType) ToLinkedWorkspacePropsPtrOutputWithContext(ctx context.Context) LinkedWorkspacePropsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedWorkspacePropsPtrOutput)
}

type LinkedWorkspacePropsOutput struct{ *pulumi.OutputState }

func (LinkedWorkspacePropsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedWorkspaceProps)(nil)).Elem()
}

func (o LinkedWorkspacePropsOutput) ToLinkedWorkspacePropsOutput() LinkedWorkspacePropsOutput {
	return o
}

func (o LinkedWorkspacePropsOutput) ToLinkedWorkspacePropsOutputWithContext(ctx context.Context) LinkedWorkspacePropsOutput {
	return o
}

func (o LinkedWorkspacePropsOutput) ToLinkedWorkspacePropsPtrOutput() LinkedWorkspacePropsPtrOutput {
	return o.ToLinkedWorkspacePropsPtrOutputWithContext(context.Background())
}

func (o LinkedWorkspacePropsOutput) ToLinkedWorkspacePropsPtrOutputWithContext(ctx context.Context) LinkedWorkspacePropsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinkedWorkspaceProps) *LinkedWorkspaceProps {
		return &v
	}).(LinkedWorkspacePropsPtrOutput)
}

func (o LinkedWorkspacePropsOutput) LinkedWorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedWorkspaceProps) *string { return v.LinkedWorkspaceResourceId }).(pulumi.StringPtrOutput)
}

func (o LinkedWorkspacePropsOutput) UserAssignedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedWorkspaceProps) *string { return v.UserAssignedIdentityResourceId }).(pulumi.StringPtrOutput)
}

type LinkedWorkspacePropsPtrOutput struct{ *pulumi.OutputState }

func (LinkedWorkspacePropsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedWorkspaceProps)(nil)).Elem()
}

func (o LinkedWorkspacePropsPtrOutput) ToLinkedWorkspacePropsPtrOutput() LinkedWorkspacePropsPtrOutput {
	return o
}

func (o LinkedWorkspacePropsPtrOutput) ToLinkedWorkspacePropsPtrOutputWithContext(ctx context.Context) LinkedWorkspacePropsPtrOutput {
	return o
}

func (o LinkedWorkspacePropsPtrOutput) Elem() LinkedWorkspacePropsOutput {
	return o.ApplyT(func(v *LinkedWorkspaceProps) LinkedWorkspaceProps {
		if v != nil {
			return *v
		}
		var ret LinkedWorkspaceProps
		return ret
	}).(LinkedWorkspacePropsOutput)
}

func (o LinkedWorkspacePropsPtrOutput) LinkedWorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedWorkspaceProps) *string {
		if v == nil {
			return nil
		}
		return v.LinkedWorkspaceResourceId
	}).(pulumi.StringPtrOutput)
}

func (o LinkedWorkspacePropsPtrOutput) UserAssignedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedWorkspaceProps) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentityResourceId
	}).(pulumi.StringPtrOutput)
}

type LinkedWorkspacePropsResponse struct {
	LinkedWorkspaceResourceId      *string `pulumi:"linkedWorkspaceResourceId"`
	UserAssignedIdentityResourceId *string `pulumi:"userAssignedIdentityResourceId"`
}

type LinkedWorkspacePropsResponseOutput struct{ *pulumi.OutputState }

func (LinkedWorkspacePropsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedWorkspacePropsResponse)(nil)).Elem()
}

func (o LinkedWorkspacePropsResponseOutput) ToLinkedWorkspacePropsResponseOutput() LinkedWorkspacePropsResponseOutput {
	return o
}

func (o LinkedWorkspacePropsResponseOutput) ToLinkedWorkspacePropsResponseOutputWithContext(ctx context.Context) LinkedWorkspacePropsResponseOutput {
	return o
}

func (o LinkedWorkspacePropsResponseOutput) LinkedWorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedWorkspacePropsResponse) *string { return v.LinkedWorkspaceResourceId }).(pulumi.StringPtrOutput)
}

func (o LinkedWorkspacePropsResponseOutput) UserAssignedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedWorkspacePropsResponse) *string { return v.UserAssignedIdentityResourceId }).(pulumi.StringPtrOutput)
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

type ServiceResponseBaseResponseError struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
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

type UserInfoResponse struct {
	UserAltSecId *string `pulumi:"userAltSecId"`
	UserIdp      *string `pulumi:"userIdp"`
	UserIss      *string `pulumi:"userIss"`
	UserName     *string `pulumi:"userName"`
	UserObjectId *string `pulumi:"userObjectId"`
	UserPuId     *string `pulumi:"userPuId"`
	UserTenantId *string `pulumi:"userTenantId"`
}

type UserInfoResponseOutput struct{ *pulumi.OutputState }

func (UserInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutput() UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutputWithContext(ctx context.Context) UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) UserAltSecId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserAltSecId }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) UserIdp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserIdp }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) UserIss() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserIss }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) UserObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserObjectId }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) UserPuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserPuId }).(pulumi.StringPtrOutput)
}

func (o UserInfoResponseOutput) UserTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserTenantId }).(pulumi.StringPtrOutput)
}

type UserInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (UserInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) Elem() UserInfoResponseOutput {
	return o.ApplyT(func(v *UserInfoResponse) UserInfoResponse {
		if v != nil {
			return *v
		}
		var ret UserInfoResponse
		return ret
	}).(UserInfoResponseOutput)
}

func (o UserInfoResponsePtrOutput) UserAltSecId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAltSecId
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) UserIdp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserIdp
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) UserIss() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserIss
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) UserObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserObjectId
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) UserPuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserPuId
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) UserTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserTenantId
	}).(pulumi.StringPtrOutput)
}

type VirtualMachine struct {
	ComputeLocation *string                   `pulumi:"computeLocation"`
	ComputeType     string                    `pulumi:"computeType"`
	Description     *string                   `pulumi:"description"`
	Properties      *VirtualMachineProperties `pulumi:"properties"`
	ResourceId      *string                   `pulumi:"resourceId"`
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
	pulumi.RegisterOutputType(AzureDataLakeSectionResponseOutput{})
	pulumi.RegisterOutputType(AzureDataLakeSectionResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureMySqlSectionResponseOutput{})
	pulumi.RegisterOutputType(AzureMySqlSectionResponsePtrOutput{})
	pulumi.RegisterOutputType(AzurePostgreSqlSectionResponseOutput{})
	pulumi.RegisterOutputType(AzurePostgreSqlSectionResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureSqlDatabaseSectionResponseOutput{})
	pulumi.RegisterOutputType(AzureSqlDatabaseSectionResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureStorageSectionResponseOutput{})
	pulumi.RegisterOutputType(AzureStorageSectionResponsePtrOutput{})
	pulumi.RegisterOutputType(ClientCredentialsResponseOutput{})
	pulumi.RegisterOutputType(ClientCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerResourceRequirementsOutput{})
	pulumi.RegisterOutputType(ContainerResourceRequirementsPtrOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestEnvironmentImageRequestOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestEnvironmentImageRequestPtrOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestKeysOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestKeysPtrOutput{})
	pulumi.RegisterOutputType(DatasetCreateRequestDataPathOutput{})
	pulumi.RegisterOutputType(DatasetCreateRequestDataPathPtrOutput{})
	pulumi.RegisterOutputType(DatasetCreateRequestParametersOutput{})
	pulumi.RegisterOutputType(DatasetCreateRequestPathOutput{})
	pulumi.RegisterOutputType(DatasetCreateRequestPathPtrOutput{})
	pulumi.RegisterOutputType(DatasetCreateRequestQueryOutput{})
	pulumi.RegisterOutputType(DatasetCreateRequestQueryPtrOutput{})
	pulumi.RegisterOutputType(DatasetCreateRequestRegistrationOutput{})
	pulumi.RegisterOutputType(DatasetCreateRequestTimeSeriesOutput{})
	pulumi.RegisterOutputType(DatasetCreateRequestTimeSeriesPtrOutput{})
	pulumi.RegisterOutputType(DatasetReferenceOutput{})
	pulumi.RegisterOutputType(DatasetReferenceArrayOutput{})
	pulumi.RegisterOutputType(DatasetResponseOutput{})
	pulumi.RegisterOutputType(DatasetResponseDataPathOutput{})
	pulumi.RegisterOutputType(DatasetResponseDataPathPtrOutput{})
	pulumi.RegisterOutputType(DatasetResponseLatestOutput{})
	pulumi.RegisterOutputType(DatasetResponseLatestPtrOutput{})
	pulumi.RegisterOutputType(DatasetResponseSqlDataPathOutput{})
	pulumi.RegisterOutputType(DatasetResponseSqlDataPathPtrOutput{})
	pulumi.RegisterOutputType(DatasetStateResponseOutput{})
	pulumi.RegisterOutputType(DatasetStateResponsePtrOutput{})
	pulumi.RegisterOutputType(DatasetStateResponseDeprecatedByOutput{})
	pulumi.RegisterOutputType(DatasetStateResponseDeprecatedByPtrOutput{})
	pulumi.RegisterOutputType(DatastoreResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentReferenceOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentReferencePtrOutput{})
	pulumi.RegisterOutputType(GlusterFsSectionResponseOutput{})
	pulumi.RegisterOutputType(GlusterFsSectionResponsePtrOutput{})
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
	pulumi.RegisterOutputType(LinkedInfoResponseOutput{})
	pulumi.RegisterOutputType(LinkedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(LinkedWorkspacePropsOutput{})
	pulumi.RegisterOutputType(LinkedWorkspacePropsPtrOutput{})
	pulumi.RegisterOutputType(LinkedWorkspacePropsResponseOutput{})
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
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(UserInfoResponseOutput{})
	pulumi.RegisterOutputType(UserInfoResponsePtrOutput{})
}
