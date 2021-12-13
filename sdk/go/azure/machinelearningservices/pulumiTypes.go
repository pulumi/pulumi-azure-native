


package machinelearningservices

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
	return &tmp
}

type AKSReplicaStatusResponseError struct {
	Error ErrorResponseResponse `pulumi:"error"`
}

type AKSResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	Properties         *AKSResponseProperties                `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
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
	return &tmp
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
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
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
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
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
	Cpu             *float64 `pulumi:"cpu"`
	CpuLimit        *float64 `pulumi:"cpuLimit"`
	Fpga            *int     `pulumi:"fpga"`
	Gpu             *int     `pulumi:"gpu"`
	MemoryInGB      *float64 `pulumi:"memoryInGB"`
	MemoryInGBLimit *float64 `pulumi:"memoryInGBLimit"`
}





type ContainerResourceRequirementsInput interface {
	pulumi.Input

	ToContainerResourceRequirementsOutput() ContainerResourceRequirementsOutput
	ToContainerResourceRequirementsOutputWithContext(context.Context) ContainerResourceRequirementsOutput
}

type ContainerResourceRequirementsArgs struct {
	Cpu             pulumi.Float64PtrInput `pulumi:"cpu"`
	CpuLimit        pulumi.Float64PtrInput `pulumi:"cpuLimit"`
	Fpga            pulumi.IntPtrInput     `pulumi:"fpga"`
	Gpu             pulumi.IntPtrInput     `pulumi:"gpu"`
	MemoryInGB      pulumi.Float64PtrInput `pulumi:"memoryInGB"`
	MemoryInGBLimit pulumi.Float64PtrInput `pulumi:"memoryInGBLimit"`
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

func (o ContainerResourceRequirementsOutput) CpuLimit() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerResourceRequirements) *float64 { return v.CpuLimit }).(pulumi.Float64PtrOutput)
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

func (o ContainerResourceRequirementsOutput) MemoryInGBLimit() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerResourceRequirements) *float64 { return v.MemoryInGBLimit }).(pulumi.Float64PtrOutput)
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

func (o ContainerResourceRequirementsPtrOutput) CpuLimit() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContainerResourceRequirements) *float64 {
		if v == nil {
			return nil
		}
		return v.CpuLimit
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

func (o ContainerResourceRequirementsPtrOutput) MemoryInGBLimit() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContainerResourceRequirements) *float64 {
		if v == nil {
			return nil
		}
		return v.MemoryInGBLimit
	}).(pulumi.Float64PtrOutput)
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
	ComputeLocation *string `pulumi:"computeLocation"`
	ComputeType     string  `pulumi:"computeType"`
	Description     *string `pulumi:"description"`
	ResourceId      *string `pulumi:"resourceId"`
}

type DataFactoryResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
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
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	Properties         *DataLakeAnalyticsResponseProperties  `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
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
	ComputeLocation *string               `pulumi:"computeLocation"`
	ComputeType     string                `pulumi:"computeType"`
	Description     *string               `pulumi:"description"`
	Properties      *DatabricksProperties `pulumi:"properties"`
	ResourceId      *string               `pulumi:"resourceId"`
}

type DatabricksProperties struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
	WorkspaceUrl          *string `pulumi:"workspaceUrl"`
}

type DatabricksResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	Properties         *DatabricksResponseProperties         `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type DatabricksResponseProperties struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
	WorkspaceUrl          *string `pulumi:"workspaceUrl"`
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

type ErrorDetailResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}

type ErrorResponseResponse struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
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
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
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
	UserAssignedIdentity string `pulumi:"userAssignedIdentity"`
}





type IdentityForCmkInput interface {
	pulumi.Input

	ToIdentityForCmkOutput() IdentityForCmkOutput
	ToIdentityForCmkOutputWithContext(context.Context) IdentityForCmkOutput
}

type IdentityForCmkArgs struct {
	UserAssignedIdentity pulumi.StringInput `pulumi:"userAssignedIdentity"`
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

func (o IdentityForCmkOutput) UserAssignedIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityForCmk) string { return v.UserAssignedIdentity }).(pulumi.StringOutput)
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
		return &v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type IdentityForCmkResponse struct {
	UserAssignedIdentity string `pulumi:"userAssignedIdentity"`
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

func (o IdentityForCmkResponseOutput) UserAssignedIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityForCmkResponse) string { return v.UserAssignedIdentity }).(pulumi.StringOutput)
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
		return &v.UserAssignedIdentity
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

type ListNotebookKeysResultResponse struct {
	PrimaryAccessKey   string `pulumi:"primaryAccessKey"`
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
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

type ServiceResponseBaseResponseError struct {
	Error ErrorResponseResponse `pulumi:"error"`
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
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
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
	pulumi.RegisterOutputType(ClientCredentialsResponseOutput{})
	pulumi.RegisterOutputType(ClientCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(CodeConfigurationOutput{})
	pulumi.RegisterOutputType(CodeConfigurationPtrOutput{})
	pulumi.RegisterOutputType(CodeConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CodeConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(CodeContainerTypeOutput{})
	pulumi.RegisterOutputType(CodeContainerResponseOutput{})
	pulumi.RegisterOutputType(CodeVersionTypeOutput{})
	pulumi.RegisterOutputType(CodeVersionResponseOutput{})
	pulumi.RegisterOutputType(ComputeBindingOutput{})
	pulumi.RegisterOutputType(ComputeBindingPtrOutput{})
	pulumi.RegisterOutputType(ComputeBindingResponseOutput{})
	pulumi.RegisterOutputType(ComputeBindingResponsePtrOutput{})
	pulumi.RegisterOutputType(ComputeConfigurationOutput{})
	pulumi.RegisterOutputType(ComputeConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ComputeConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ComputeConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerResourceRequirementsOutput{})
	pulumi.RegisterOutputType(ContainerResourceRequirementsPtrOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsPtrOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsResponseOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestEnvironmentImageRequestOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestEnvironmentImageRequestPtrOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestKeysOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestKeysPtrOutput{})
	pulumi.RegisterOutputType(DataContainerTypeOutput{})
	pulumi.RegisterOutputType(DataContainerResponseOutput{})
	pulumi.RegisterOutputType(DataVersionTypeOutput{})
	pulumi.RegisterOutputType(DataVersionResponseOutput{})
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
	pulumi.RegisterOutputType(EndpointAuthKeysOutput{})
	pulumi.RegisterOutputType(EndpointAuthKeysPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentContainerTypeOutput{})
	pulumi.RegisterOutputType(EnvironmentContainerResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentReferenceOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentReferencePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentSpecificationVersionTypeOutput{})
	pulumi.RegisterOutputType(EnvironmentSpecificationVersionResponseOutput{})
	pulumi.RegisterOutputType(FlavorDataOutput{})
	pulumi.RegisterOutputType(FlavorDataMapOutput{})
	pulumi.RegisterOutputType(FlavorDataResponseOutput{})
	pulumi.RegisterOutputType(FlavorDataResponseMapOutput{})
	pulumi.RegisterOutputType(GlusterFsSectionResponseOutput{})
	pulumi.RegisterOutputType(GlusterFsSectionResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityForCmkOutput{})
	pulumi.RegisterOutputType(IdentityForCmkPtrOutput{})
	pulumi.RegisterOutputType(IdentityForCmkResponseOutput{})
	pulumi.RegisterOutputType(IdentityForCmkResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageAssetOutput{})
	pulumi.RegisterOutputType(ImageAssetArrayOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(InferenceContainerPropertiesResponsePtrOutput{})
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
	pulumi.RegisterOutputType(LinkedInfoResponseOutput{})
	pulumi.RegisterOutputType(LinkedInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(LinkedServicePropsOutput{})
	pulumi.RegisterOutputType(LinkedServicePropsPtrOutput{})
	pulumi.RegisterOutputType(LinkedServicePropsResponseOutput{})
	pulumi.RegisterOutputType(LinkedWorkspacePropsOutput{})
	pulumi.RegisterOutputType(LinkedWorkspacePropsPtrOutput{})
	pulumi.RegisterOutputType(LinkedWorkspacePropsResponseOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationPtrOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationResponseOutput{})
	pulumi.RegisterOutputType(MLAssistConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ModelOutput{})
	pulumi.RegisterOutputType(ModelArrayOutput{})
	pulumi.RegisterOutputType(ModelContainerTypeOutput{})
	pulumi.RegisterOutputType(ModelContainerResponseOutput{})
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
	pulumi.RegisterOutputType(RCranPackageOutput{})
	pulumi.RegisterOutputType(RCranPackageArrayOutput{})
	pulumi.RegisterOutputType(RGitHubPackageOutput{})
	pulumi.RegisterOutputType(RGitHubPackageArrayOutput{})
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
	pulumi.RegisterOutputType(SparkMavenPackageOutput{})
	pulumi.RegisterOutputType(SparkMavenPackageArrayOutput{})
	pulumi.RegisterOutputType(StatusMessageResponseOutput{})
	pulumi.RegisterOutputType(StatusMessageResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityMetaOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityMetaMapOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityMetaResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityMetaResponseMapOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(UserInfoResponseOutput{})
	pulumi.RegisterOutputType(UserInfoResponsePtrOutput{})
}
