


package v20200515preview

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





type ACIServiceResponseResponseInput interface {
	pulumi.Input

	ToACIServiceResponseResponseOutput() ACIServiceResponseResponseOutput
	ToACIServiceResponseResponseOutputWithContext(context.Context) ACIServiceResponseResponseOutput
}

type ACIServiceResponseResponseArgs struct {
	AppInsightsEnabled            pulumi.BoolPtrInput                                       `pulumi:"appInsightsEnabled"`
	AuthEnabled                   pulumi.BoolPtrInput                                       `pulumi:"authEnabled"`
	Cname                         pulumi.StringPtrInput                                     `pulumi:"cname"`
	ComputeType                   pulumi.StringInput                                        `pulumi:"computeType"`
	ContainerResourceRequirements ContainerResourceRequirementsResponsePtrInput             `pulumi:"containerResourceRequirements"`
	DataCollection                ACIServiceResponseResponseDataCollectionPtrInput          `pulumi:"dataCollection"`
	DeploymentType                pulumi.StringPtrInput                                     `pulumi:"deploymentType"`
	Description                   pulumi.StringPtrInput                                     `pulumi:"description"`
	EncryptionProperties          ACIServiceResponseResponseEncryptionPropertiesPtrInput    `pulumi:"encryptionProperties"`
	EnvironmentImageRequest       ACIServiceResponseResponseEnvironmentImageRequestPtrInput `pulumi:"environmentImageRequest"`
	Error                         ServiceResponseBaseResponseErrorInput                     `pulumi:"error"`
	KvTags                        pulumi.StringMapInput                                     `pulumi:"kvTags"`
	Location                      pulumi.StringPtrInput                                     `pulumi:"location"`
	ModelConfigMap                pulumi.MapInput                                           `pulumi:"modelConfigMap"`
	Models                        ModelResponseArrayInput                                   `pulumi:"models"`
	Properties                    pulumi.StringMapInput                                     `pulumi:"properties"`
	PublicFqdn                    pulumi.StringPtrInput                                     `pulumi:"publicFqdn"`
	PublicIp                      pulumi.StringPtrInput                                     `pulumi:"publicIp"`
	ScoringUri                    pulumi.StringInput                                        `pulumi:"scoringUri"`
	SslCertificate                pulumi.StringPtrInput                                     `pulumi:"sslCertificate"`
	SslEnabled                    pulumi.BoolPtrInput                                       `pulumi:"sslEnabled"`
	SslKey                        pulumi.StringPtrInput                                     `pulumi:"sslKey"`
	State                         pulumi.StringInput                                        `pulumi:"state"`
	SwaggerUri                    pulumi.StringInput                                        `pulumi:"swaggerUri"`
	VnetConfiguration             ACIServiceResponseResponseVnetConfigurationPtrInput       `pulumi:"vnetConfiguration"`
}

func (ACIServiceResponseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceResponseResponse)(nil)).Elem()
}

func (i ACIServiceResponseResponseArgs) ToACIServiceResponseResponseOutput() ACIServiceResponseResponseOutput {
	return i.ToACIServiceResponseResponseOutputWithContext(context.Background())
}

func (i ACIServiceResponseResponseArgs) ToACIServiceResponseResponseOutputWithContext(ctx context.Context) ACIServiceResponseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseOutput)
}

type ACIServiceResponseResponseOutput struct{ *pulumi.OutputState }

func (ACIServiceResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceResponseResponse)(nil)).Elem()
}

func (o ACIServiceResponseResponseOutput) ToACIServiceResponseResponseOutput() ACIServiceResponseResponseOutput {
	return o
}

func (o ACIServiceResponseResponseOutput) ToACIServiceResponseResponseOutputWithContext(ctx context.Context) ACIServiceResponseResponseOutput {
	return o
}

func (o ACIServiceResponseResponseOutput) AppInsightsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *bool { return v.AppInsightsEnabled }).(pulumi.BoolPtrOutput)
}

func (o ACIServiceResponseResponseOutput) AuthEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *bool { return v.AuthEnabled }).(pulumi.BoolPtrOutput)
}

func (o ACIServiceResponseResponseOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *string { return v.Cname }).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o ACIServiceResponseResponseOutput) ContainerResourceRequirements() ContainerResourceRequirementsResponsePtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *ContainerResourceRequirementsResponse {
		return v.ContainerResourceRequirements
	}).(ContainerResourceRequirementsResponsePtrOutput)
}

func (o ACIServiceResponseResponseOutput) DataCollection() ACIServiceResponseResponseDataCollectionPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *ACIServiceResponseResponseDataCollection { return v.DataCollection }).(ACIServiceResponseResponseDataCollectionPtrOutput)
}

func (o ACIServiceResponseResponseOutput) DeploymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *string { return v.DeploymentType }).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseOutput) EncryptionProperties() ACIServiceResponseResponseEncryptionPropertiesPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *ACIServiceResponseResponseEncryptionProperties {
		return v.EncryptionProperties
	}).(ACIServiceResponseResponseEncryptionPropertiesPtrOutput)
}

func (o ACIServiceResponseResponseOutput) EnvironmentImageRequest() ACIServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *ACIServiceResponseResponseEnvironmentImageRequest {
		return v.EnvironmentImageRequest
	}).(ACIServiceResponseResponseEnvironmentImageRequestPtrOutput)
}

func (o ACIServiceResponseResponseOutput) Error() ServiceResponseBaseResponseErrorOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) ServiceResponseBaseResponseError { return v.Error }).(ServiceResponseBaseResponseErrorOutput)
}

func (o ACIServiceResponseResponseOutput) KvTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) map[string]string { return v.KvTags }).(pulumi.StringMapOutput)
}

func (o ACIServiceResponseResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseOutput) ModelConfigMap() pulumi.MapOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) map[string]interface{} { return v.ModelConfigMap }).(pulumi.MapOutput)
}

func (o ACIServiceResponseResponseOutput) Models() ModelResponseArrayOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) []ModelResponse { return v.Models }).(ModelResponseArrayOutput)
}

func (o ACIServiceResponseResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ACIServiceResponseResponseOutput) PublicFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *string { return v.PublicFqdn }).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseOutput) PublicIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *string { return v.PublicIp }).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseOutput) ScoringUri() pulumi.StringOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) string { return v.ScoringUri }).(pulumi.StringOutput)
}

func (o ACIServiceResponseResponseOutput) SslCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *string { return v.SslCertificate }).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseOutput) SslEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *bool { return v.SslEnabled }).(pulumi.BoolPtrOutput)
}

func (o ACIServiceResponseResponseOutput) SslKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *string { return v.SslKey }).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o ACIServiceResponseResponseOutput) SwaggerUri() pulumi.StringOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) string { return v.SwaggerUri }).(pulumi.StringOutput)
}

func (o ACIServiceResponseResponseOutput) VnetConfiguration() ACIServiceResponseResponseVnetConfigurationPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponse) *ACIServiceResponseResponseVnetConfiguration {
		return v.VnetConfiguration
	}).(ACIServiceResponseResponseVnetConfigurationPtrOutput)
}

type ACIServiceResponseResponseDataCollection struct {
	EventHubEnabled *bool `pulumi:"eventHubEnabled"`
	StorageEnabled  *bool `pulumi:"storageEnabled"`
}





type ACIServiceResponseResponseDataCollectionInput interface {
	pulumi.Input

	ToACIServiceResponseResponseDataCollectionOutput() ACIServiceResponseResponseDataCollectionOutput
	ToACIServiceResponseResponseDataCollectionOutputWithContext(context.Context) ACIServiceResponseResponseDataCollectionOutput
}

type ACIServiceResponseResponseDataCollectionArgs struct {
	EventHubEnabled pulumi.BoolPtrInput `pulumi:"eventHubEnabled"`
	StorageEnabled  pulumi.BoolPtrInput `pulumi:"storageEnabled"`
}

func (ACIServiceResponseResponseDataCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceResponseResponseDataCollection)(nil)).Elem()
}

func (i ACIServiceResponseResponseDataCollectionArgs) ToACIServiceResponseResponseDataCollectionOutput() ACIServiceResponseResponseDataCollectionOutput {
	return i.ToACIServiceResponseResponseDataCollectionOutputWithContext(context.Background())
}

func (i ACIServiceResponseResponseDataCollectionArgs) ToACIServiceResponseResponseDataCollectionOutputWithContext(ctx context.Context) ACIServiceResponseResponseDataCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseDataCollectionOutput)
}

func (i ACIServiceResponseResponseDataCollectionArgs) ToACIServiceResponseResponseDataCollectionPtrOutput() ACIServiceResponseResponseDataCollectionPtrOutput {
	return i.ToACIServiceResponseResponseDataCollectionPtrOutputWithContext(context.Background())
}

func (i ACIServiceResponseResponseDataCollectionArgs) ToACIServiceResponseResponseDataCollectionPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseDataCollectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseDataCollectionOutput).ToACIServiceResponseResponseDataCollectionPtrOutputWithContext(ctx)
}









type ACIServiceResponseResponseDataCollectionPtrInput interface {
	pulumi.Input

	ToACIServiceResponseResponseDataCollectionPtrOutput() ACIServiceResponseResponseDataCollectionPtrOutput
	ToACIServiceResponseResponseDataCollectionPtrOutputWithContext(context.Context) ACIServiceResponseResponseDataCollectionPtrOutput
}

type aciserviceResponseResponseDataCollectionPtrType ACIServiceResponseResponseDataCollectionArgs

func ACIServiceResponseResponseDataCollectionPtr(v *ACIServiceResponseResponseDataCollectionArgs) ACIServiceResponseResponseDataCollectionPtrInput {
	return (*aciserviceResponseResponseDataCollectionPtrType)(v)
}

func (*aciserviceResponseResponseDataCollectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceResponseResponseDataCollection)(nil)).Elem()
}

func (i *aciserviceResponseResponseDataCollectionPtrType) ToACIServiceResponseResponseDataCollectionPtrOutput() ACIServiceResponseResponseDataCollectionPtrOutput {
	return i.ToACIServiceResponseResponseDataCollectionPtrOutputWithContext(context.Background())
}

func (i *aciserviceResponseResponseDataCollectionPtrType) ToACIServiceResponseResponseDataCollectionPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseDataCollectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseDataCollectionPtrOutput)
}

type ACIServiceResponseResponseDataCollectionOutput struct{ *pulumi.OutputState }

func (ACIServiceResponseResponseDataCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceResponseResponseDataCollection)(nil)).Elem()
}

func (o ACIServiceResponseResponseDataCollectionOutput) ToACIServiceResponseResponseDataCollectionOutput() ACIServiceResponseResponseDataCollectionOutput {
	return o
}

func (o ACIServiceResponseResponseDataCollectionOutput) ToACIServiceResponseResponseDataCollectionOutputWithContext(ctx context.Context) ACIServiceResponseResponseDataCollectionOutput {
	return o
}

func (o ACIServiceResponseResponseDataCollectionOutput) ToACIServiceResponseResponseDataCollectionPtrOutput() ACIServiceResponseResponseDataCollectionPtrOutput {
	return o.ToACIServiceResponseResponseDataCollectionPtrOutputWithContext(context.Background())
}

func (o ACIServiceResponseResponseDataCollectionOutput) ToACIServiceResponseResponseDataCollectionPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseDataCollectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ACIServiceResponseResponseDataCollection) *ACIServiceResponseResponseDataCollection {
		return &v
	}).(ACIServiceResponseResponseDataCollectionPtrOutput)
}

func (o ACIServiceResponseResponseDataCollectionOutput) EventHubEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseDataCollection) *bool { return v.EventHubEnabled }).(pulumi.BoolPtrOutput)
}

func (o ACIServiceResponseResponseDataCollectionOutput) StorageEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseDataCollection) *bool { return v.StorageEnabled }).(pulumi.BoolPtrOutput)
}

type ACIServiceResponseResponseDataCollectionPtrOutput struct{ *pulumi.OutputState }

func (ACIServiceResponseResponseDataCollectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceResponseResponseDataCollection)(nil)).Elem()
}

func (o ACIServiceResponseResponseDataCollectionPtrOutput) ToACIServiceResponseResponseDataCollectionPtrOutput() ACIServiceResponseResponseDataCollectionPtrOutput {
	return o
}

func (o ACIServiceResponseResponseDataCollectionPtrOutput) ToACIServiceResponseResponseDataCollectionPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseDataCollectionPtrOutput {
	return o
}

func (o ACIServiceResponseResponseDataCollectionPtrOutput) Elem() ACIServiceResponseResponseDataCollectionOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseDataCollection) ACIServiceResponseResponseDataCollection {
		if v != nil {
			return *v
		}
		var ret ACIServiceResponseResponseDataCollection
		return ret
	}).(ACIServiceResponseResponseDataCollectionOutput)
}

func (o ACIServiceResponseResponseDataCollectionPtrOutput) EventHubEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseDataCollection) *bool {
		if v == nil {
			return nil
		}
		return v.EventHubEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ACIServiceResponseResponseDataCollectionPtrOutput) StorageEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseDataCollection) *bool {
		if v == nil {
			return nil
		}
		return v.StorageEnabled
	}).(pulumi.BoolPtrOutput)
}

type ACIServiceResponseResponseEncryptionProperties struct {
	KeyName      string `pulumi:"keyName"`
	KeyVersion   string `pulumi:"keyVersion"`
	VaultBaseUrl string `pulumi:"vaultBaseUrl"`
}





type ACIServiceResponseResponseEncryptionPropertiesInput interface {
	pulumi.Input

	ToACIServiceResponseResponseEncryptionPropertiesOutput() ACIServiceResponseResponseEncryptionPropertiesOutput
	ToACIServiceResponseResponseEncryptionPropertiesOutputWithContext(context.Context) ACIServiceResponseResponseEncryptionPropertiesOutput
}

type ACIServiceResponseResponseEncryptionPropertiesArgs struct {
	KeyName      pulumi.StringInput `pulumi:"keyName"`
	KeyVersion   pulumi.StringInput `pulumi:"keyVersion"`
	VaultBaseUrl pulumi.StringInput `pulumi:"vaultBaseUrl"`
}

func (ACIServiceResponseResponseEncryptionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceResponseResponseEncryptionProperties)(nil)).Elem()
}

func (i ACIServiceResponseResponseEncryptionPropertiesArgs) ToACIServiceResponseResponseEncryptionPropertiesOutput() ACIServiceResponseResponseEncryptionPropertiesOutput {
	return i.ToACIServiceResponseResponseEncryptionPropertiesOutputWithContext(context.Background())
}

func (i ACIServiceResponseResponseEncryptionPropertiesArgs) ToACIServiceResponseResponseEncryptionPropertiesOutputWithContext(ctx context.Context) ACIServiceResponseResponseEncryptionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseEncryptionPropertiesOutput)
}

func (i ACIServiceResponseResponseEncryptionPropertiesArgs) ToACIServiceResponseResponseEncryptionPropertiesPtrOutput() ACIServiceResponseResponseEncryptionPropertiesPtrOutput {
	return i.ToACIServiceResponseResponseEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i ACIServiceResponseResponseEncryptionPropertiesArgs) ToACIServiceResponseResponseEncryptionPropertiesPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseEncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseEncryptionPropertiesOutput).ToACIServiceResponseResponseEncryptionPropertiesPtrOutputWithContext(ctx)
}









type ACIServiceResponseResponseEncryptionPropertiesPtrInput interface {
	pulumi.Input

	ToACIServiceResponseResponseEncryptionPropertiesPtrOutput() ACIServiceResponseResponseEncryptionPropertiesPtrOutput
	ToACIServiceResponseResponseEncryptionPropertiesPtrOutputWithContext(context.Context) ACIServiceResponseResponseEncryptionPropertiesPtrOutput
}

type aciserviceResponseResponseEncryptionPropertiesPtrType ACIServiceResponseResponseEncryptionPropertiesArgs

func ACIServiceResponseResponseEncryptionPropertiesPtr(v *ACIServiceResponseResponseEncryptionPropertiesArgs) ACIServiceResponseResponseEncryptionPropertiesPtrInput {
	return (*aciserviceResponseResponseEncryptionPropertiesPtrType)(v)
}

func (*aciserviceResponseResponseEncryptionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceResponseResponseEncryptionProperties)(nil)).Elem()
}

func (i *aciserviceResponseResponseEncryptionPropertiesPtrType) ToACIServiceResponseResponseEncryptionPropertiesPtrOutput() ACIServiceResponseResponseEncryptionPropertiesPtrOutput {
	return i.ToACIServiceResponseResponseEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i *aciserviceResponseResponseEncryptionPropertiesPtrType) ToACIServiceResponseResponseEncryptionPropertiesPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseEncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseEncryptionPropertiesPtrOutput)
}

type ACIServiceResponseResponseEncryptionPropertiesOutput struct{ *pulumi.OutputState }

func (ACIServiceResponseResponseEncryptionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceResponseResponseEncryptionProperties)(nil)).Elem()
}

func (o ACIServiceResponseResponseEncryptionPropertiesOutput) ToACIServiceResponseResponseEncryptionPropertiesOutput() ACIServiceResponseResponseEncryptionPropertiesOutput {
	return o
}

func (o ACIServiceResponseResponseEncryptionPropertiesOutput) ToACIServiceResponseResponseEncryptionPropertiesOutputWithContext(ctx context.Context) ACIServiceResponseResponseEncryptionPropertiesOutput {
	return o
}

func (o ACIServiceResponseResponseEncryptionPropertiesOutput) ToACIServiceResponseResponseEncryptionPropertiesPtrOutput() ACIServiceResponseResponseEncryptionPropertiesPtrOutput {
	return o.ToACIServiceResponseResponseEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (o ACIServiceResponseResponseEncryptionPropertiesOutput) ToACIServiceResponseResponseEncryptionPropertiesPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseEncryptionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ACIServiceResponseResponseEncryptionProperties) *ACIServiceResponseResponseEncryptionProperties {
		return &v
	}).(ACIServiceResponseResponseEncryptionPropertiesPtrOutput)
}

func (o ACIServiceResponseResponseEncryptionPropertiesOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseEncryptionProperties) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o ACIServiceResponseResponseEncryptionPropertiesOutput) KeyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseEncryptionProperties) string { return v.KeyVersion }).(pulumi.StringOutput)
}

func (o ACIServiceResponseResponseEncryptionPropertiesOutput) VaultBaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseEncryptionProperties) string { return v.VaultBaseUrl }).(pulumi.StringOutput)
}

type ACIServiceResponseResponseEncryptionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ACIServiceResponseResponseEncryptionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceResponseResponseEncryptionProperties)(nil)).Elem()
}

func (o ACIServiceResponseResponseEncryptionPropertiesPtrOutput) ToACIServiceResponseResponseEncryptionPropertiesPtrOutput() ACIServiceResponseResponseEncryptionPropertiesPtrOutput {
	return o
}

func (o ACIServiceResponseResponseEncryptionPropertiesPtrOutput) ToACIServiceResponseResponseEncryptionPropertiesPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseEncryptionPropertiesPtrOutput {
	return o
}

func (o ACIServiceResponseResponseEncryptionPropertiesPtrOutput) Elem() ACIServiceResponseResponseEncryptionPropertiesOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseEncryptionProperties) ACIServiceResponseResponseEncryptionProperties {
		if v != nil {
			return *v
		}
		var ret ACIServiceResponseResponseEncryptionProperties
		return ret
	}).(ACIServiceResponseResponseEncryptionPropertiesOutput)
}

func (o ACIServiceResponseResponseEncryptionPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseEncryptionPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseEncryptionPropertiesPtrOutput) VaultBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseEncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.VaultBaseUrl
	}).(pulumi.StringPtrOutput)
}

type ACIServiceResponseResponseEnvironmentImageRequest struct {
	Assets               []ImageAssetResponse                                  `pulumi:"assets"`
	DriverProgram        *string                                               `pulumi:"driverProgram"`
	Environment          *EnvironmentImageResponseResponseEnvironment          `pulumi:"environment"`
	EnvironmentReference *EnvironmentImageResponseResponseEnvironmentReference `pulumi:"environmentReference"`
	ModelIds             []string                                              `pulumi:"modelIds"`
	Models               []ModelResponse                                       `pulumi:"models"`
}





type ACIServiceResponseResponseEnvironmentImageRequestInput interface {
	pulumi.Input

	ToACIServiceResponseResponseEnvironmentImageRequestOutput() ACIServiceResponseResponseEnvironmentImageRequestOutput
	ToACIServiceResponseResponseEnvironmentImageRequestOutputWithContext(context.Context) ACIServiceResponseResponseEnvironmentImageRequestOutput
}

type ACIServiceResponseResponseEnvironmentImageRequestArgs struct {
	Assets               ImageAssetResponseArrayInput                                 `pulumi:"assets"`
	DriverProgram        pulumi.StringPtrInput                                        `pulumi:"driverProgram"`
	Environment          EnvironmentImageResponseResponseEnvironmentPtrInput          `pulumi:"environment"`
	EnvironmentReference EnvironmentImageResponseResponseEnvironmentReferencePtrInput `pulumi:"environmentReference"`
	ModelIds             pulumi.StringArrayInput                                      `pulumi:"modelIds"`
	Models               ModelResponseArrayInput                                      `pulumi:"models"`
}

func (ACIServiceResponseResponseEnvironmentImageRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceResponseResponseEnvironmentImageRequest)(nil)).Elem()
}

func (i ACIServiceResponseResponseEnvironmentImageRequestArgs) ToACIServiceResponseResponseEnvironmentImageRequestOutput() ACIServiceResponseResponseEnvironmentImageRequestOutput {
	return i.ToACIServiceResponseResponseEnvironmentImageRequestOutputWithContext(context.Background())
}

func (i ACIServiceResponseResponseEnvironmentImageRequestArgs) ToACIServiceResponseResponseEnvironmentImageRequestOutputWithContext(ctx context.Context) ACIServiceResponseResponseEnvironmentImageRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseEnvironmentImageRequestOutput)
}

func (i ACIServiceResponseResponseEnvironmentImageRequestArgs) ToACIServiceResponseResponseEnvironmentImageRequestPtrOutput() ACIServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return i.ToACIServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(context.Background())
}

func (i ACIServiceResponseResponseEnvironmentImageRequestArgs) ToACIServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseEnvironmentImageRequestOutput).ToACIServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(ctx)
}









type ACIServiceResponseResponseEnvironmentImageRequestPtrInput interface {
	pulumi.Input

	ToACIServiceResponseResponseEnvironmentImageRequestPtrOutput() ACIServiceResponseResponseEnvironmentImageRequestPtrOutput
	ToACIServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(context.Context) ACIServiceResponseResponseEnvironmentImageRequestPtrOutput
}

type aciserviceResponseResponseEnvironmentImageRequestPtrType ACIServiceResponseResponseEnvironmentImageRequestArgs

func ACIServiceResponseResponseEnvironmentImageRequestPtr(v *ACIServiceResponseResponseEnvironmentImageRequestArgs) ACIServiceResponseResponseEnvironmentImageRequestPtrInput {
	return (*aciserviceResponseResponseEnvironmentImageRequestPtrType)(v)
}

func (*aciserviceResponseResponseEnvironmentImageRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceResponseResponseEnvironmentImageRequest)(nil)).Elem()
}

func (i *aciserviceResponseResponseEnvironmentImageRequestPtrType) ToACIServiceResponseResponseEnvironmentImageRequestPtrOutput() ACIServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return i.ToACIServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(context.Background())
}

func (i *aciserviceResponseResponseEnvironmentImageRequestPtrType) ToACIServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseEnvironmentImageRequestPtrOutput)
}

type ACIServiceResponseResponseEnvironmentImageRequestOutput struct{ *pulumi.OutputState }

func (ACIServiceResponseResponseEnvironmentImageRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceResponseResponseEnvironmentImageRequest)(nil)).Elem()
}

func (o ACIServiceResponseResponseEnvironmentImageRequestOutput) ToACIServiceResponseResponseEnvironmentImageRequestOutput() ACIServiceResponseResponseEnvironmentImageRequestOutput {
	return o
}

func (o ACIServiceResponseResponseEnvironmentImageRequestOutput) ToACIServiceResponseResponseEnvironmentImageRequestOutputWithContext(ctx context.Context) ACIServiceResponseResponseEnvironmentImageRequestOutput {
	return o
}

func (o ACIServiceResponseResponseEnvironmentImageRequestOutput) ToACIServiceResponseResponseEnvironmentImageRequestPtrOutput() ACIServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return o.ToACIServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(context.Background())
}

func (o ACIServiceResponseResponseEnvironmentImageRequestOutput) ToACIServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ACIServiceResponseResponseEnvironmentImageRequest) *ACIServiceResponseResponseEnvironmentImageRequest {
		return &v
	}).(ACIServiceResponseResponseEnvironmentImageRequestPtrOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestOutput) Assets() ImageAssetResponseArrayOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseEnvironmentImageRequest) []ImageAssetResponse { return v.Assets }).(ImageAssetResponseArrayOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestOutput) DriverProgram() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseEnvironmentImageRequest) *string { return v.DriverProgram }).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestOutput) Environment() EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseEnvironmentImageRequest) *EnvironmentImageResponseResponseEnvironment {
		return v.Environment
	}).(EnvironmentImageResponseResponseEnvironmentPtrOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestOutput) EnvironmentReference() EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseEnvironmentImageRequest) *EnvironmentImageResponseResponseEnvironmentReference {
		return v.EnvironmentReference
	}).(EnvironmentImageResponseResponseEnvironmentReferencePtrOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestOutput) ModelIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseEnvironmentImageRequest) []string { return v.ModelIds }).(pulumi.StringArrayOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestOutput) Models() ModelResponseArrayOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseEnvironmentImageRequest) []ModelResponse { return v.Models }).(ModelResponseArrayOutput)
}

type ACIServiceResponseResponseEnvironmentImageRequestPtrOutput struct{ *pulumi.OutputState }

func (ACIServiceResponseResponseEnvironmentImageRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceResponseResponseEnvironmentImageRequest)(nil)).Elem()
}

func (o ACIServiceResponseResponseEnvironmentImageRequestPtrOutput) ToACIServiceResponseResponseEnvironmentImageRequestPtrOutput() ACIServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return o
}

func (o ACIServiceResponseResponseEnvironmentImageRequestPtrOutput) ToACIServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return o
}

func (o ACIServiceResponseResponseEnvironmentImageRequestPtrOutput) Elem() ACIServiceResponseResponseEnvironmentImageRequestOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseEnvironmentImageRequest) ACIServiceResponseResponseEnvironmentImageRequest {
		if v != nil {
			return *v
		}
		var ret ACIServiceResponseResponseEnvironmentImageRequest
		return ret
	}).(ACIServiceResponseResponseEnvironmentImageRequestOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestPtrOutput) Assets() ImageAssetResponseArrayOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseEnvironmentImageRequest) []ImageAssetResponse {
		if v == nil {
			return nil
		}
		return v.Assets
	}).(ImageAssetResponseArrayOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestPtrOutput) DriverProgram() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseEnvironmentImageRequest) *string {
		if v == nil {
			return nil
		}
		return v.DriverProgram
	}).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestPtrOutput) Environment() EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseEnvironmentImageRequest) *EnvironmentImageResponseResponseEnvironment {
		if v == nil {
			return nil
		}
		return v.Environment
	}).(EnvironmentImageResponseResponseEnvironmentPtrOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestPtrOutput) EnvironmentReference() EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseEnvironmentImageRequest) *EnvironmentImageResponseResponseEnvironmentReference {
		if v == nil {
			return nil
		}
		return v.EnvironmentReference
	}).(EnvironmentImageResponseResponseEnvironmentReferencePtrOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestPtrOutput) ModelIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseEnvironmentImageRequest) []string {
		if v == nil {
			return nil
		}
		return v.ModelIds
	}).(pulumi.StringArrayOutput)
}

func (o ACIServiceResponseResponseEnvironmentImageRequestPtrOutput) Models() ModelResponseArrayOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseEnvironmentImageRequest) []ModelResponse {
		if v == nil {
			return nil
		}
		return v.Models
	}).(ModelResponseArrayOutput)
}

type ACIServiceResponseResponseVnetConfiguration struct {
	SubnetName *string `pulumi:"subnetName"`
	VnetName   *string `pulumi:"vnetName"`
}





type ACIServiceResponseResponseVnetConfigurationInput interface {
	pulumi.Input

	ToACIServiceResponseResponseVnetConfigurationOutput() ACIServiceResponseResponseVnetConfigurationOutput
	ToACIServiceResponseResponseVnetConfigurationOutputWithContext(context.Context) ACIServiceResponseResponseVnetConfigurationOutput
}

type ACIServiceResponseResponseVnetConfigurationArgs struct {
	SubnetName pulumi.StringPtrInput `pulumi:"subnetName"`
	VnetName   pulumi.StringPtrInput `pulumi:"vnetName"`
}

func (ACIServiceResponseResponseVnetConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceResponseResponseVnetConfiguration)(nil)).Elem()
}

func (i ACIServiceResponseResponseVnetConfigurationArgs) ToACIServiceResponseResponseVnetConfigurationOutput() ACIServiceResponseResponseVnetConfigurationOutput {
	return i.ToACIServiceResponseResponseVnetConfigurationOutputWithContext(context.Background())
}

func (i ACIServiceResponseResponseVnetConfigurationArgs) ToACIServiceResponseResponseVnetConfigurationOutputWithContext(ctx context.Context) ACIServiceResponseResponseVnetConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseVnetConfigurationOutput)
}

func (i ACIServiceResponseResponseVnetConfigurationArgs) ToACIServiceResponseResponseVnetConfigurationPtrOutput() ACIServiceResponseResponseVnetConfigurationPtrOutput {
	return i.ToACIServiceResponseResponseVnetConfigurationPtrOutputWithContext(context.Background())
}

func (i ACIServiceResponseResponseVnetConfigurationArgs) ToACIServiceResponseResponseVnetConfigurationPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseVnetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseVnetConfigurationOutput).ToACIServiceResponseResponseVnetConfigurationPtrOutputWithContext(ctx)
}









type ACIServiceResponseResponseVnetConfigurationPtrInput interface {
	pulumi.Input

	ToACIServiceResponseResponseVnetConfigurationPtrOutput() ACIServiceResponseResponseVnetConfigurationPtrOutput
	ToACIServiceResponseResponseVnetConfigurationPtrOutputWithContext(context.Context) ACIServiceResponseResponseVnetConfigurationPtrOutput
}

type aciserviceResponseResponseVnetConfigurationPtrType ACIServiceResponseResponseVnetConfigurationArgs

func ACIServiceResponseResponseVnetConfigurationPtr(v *ACIServiceResponseResponseVnetConfigurationArgs) ACIServiceResponseResponseVnetConfigurationPtrInput {
	return (*aciserviceResponseResponseVnetConfigurationPtrType)(v)
}

func (*aciserviceResponseResponseVnetConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceResponseResponseVnetConfiguration)(nil)).Elem()
}

func (i *aciserviceResponseResponseVnetConfigurationPtrType) ToACIServiceResponseResponseVnetConfigurationPtrOutput() ACIServiceResponseResponseVnetConfigurationPtrOutput {
	return i.ToACIServiceResponseResponseVnetConfigurationPtrOutputWithContext(context.Background())
}

func (i *aciserviceResponseResponseVnetConfigurationPtrType) ToACIServiceResponseResponseVnetConfigurationPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseVnetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACIServiceResponseResponseVnetConfigurationPtrOutput)
}

type ACIServiceResponseResponseVnetConfigurationOutput struct{ *pulumi.OutputState }

func (ACIServiceResponseResponseVnetConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACIServiceResponseResponseVnetConfiguration)(nil)).Elem()
}

func (o ACIServiceResponseResponseVnetConfigurationOutput) ToACIServiceResponseResponseVnetConfigurationOutput() ACIServiceResponseResponseVnetConfigurationOutput {
	return o
}

func (o ACIServiceResponseResponseVnetConfigurationOutput) ToACIServiceResponseResponseVnetConfigurationOutputWithContext(ctx context.Context) ACIServiceResponseResponseVnetConfigurationOutput {
	return o
}

func (o ACIServiceResponseResponseVnetConfigurationOutput) ToACIServiceResponseResponseVnetConfigurationPtrOutput() ACIServiceResponseResponseVnetConfigurationPtrOutput {
	return o.ToACIServiceResponseResponseVnetConfigurationPtrOutputWithContext(context.Background())
}

func (o ACIServiceResponseResponseVnetConfigurationOutput) ToACIServiceResponseResponseVnetConfigurationPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseVnetConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ACIServiceResponseResponseVnetConfiguration) *ACIServiceResponseResponseVnetConfiguration {
		return &v
	}).(ACIServiceResponseResponseVnetConfigurationPtrOutput)
}

func (o ACIServiceResponseResponseVnetConfigurationOutput) SubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseVnetConfiguration) *string { return v.SubnetName }).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseVnetConfigurationOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACIServiceResponseResponseVnetConfiguration) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

type ACIServiceResponseResponseVnetConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ACIServiceResponseResponseVnetConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACIServiceResponseResponseVnetConfiguration)(nil)).Elem()
}

func (o ACIServiceResponseResponseVnetConfigurationPtrOutput) ToACIServiceResponseResponseVnetConfigurationPtrOutput() ACIServiceResponseResponseVnetConfigurationPtrOutput {
	return o
}

func (o ACIServiceResponseResponseVnetConfigurationPtrOutput) ToACIServiceResponseResponseVnetConfigurationPtrOutputWithContext(ctx context.Context) ACIServiceResponseResponseVnetConfigurationPtrOutput {
	return o
}

func (o ACIServiceResponseResponseVnetConfigurationPtrOutput) Elem() ACIServiceResponseResponseVnetConfigurationOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseVnetConfiguration) ACIServiceResponseResponseVnetConfiguration {
		if v != nil {
			return *v
		}
		var ret ACIServiceResponseResponseVnetConfiguration
		return ret
	}).(ACIServiceResponseResponseVnetConfigurationOutput)
}

func (o ACIServiceResponseResponseVnetConfigurationPtrOutput) SubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseVnetConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SubnetName
	}).(pulumi.StringPtrOutput)
}

func (o ACIServiceResponseResponseVnetConfigurationPtrOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACIServiceResponseResponseVnetConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.VnetName
	}).(pulumi.StringPtrOutput)
}

type AKS struct {
	ComputeLocation *string        `pulumi:"computeLocation"`
	ComputeType     string         `pulumi:"computeType"`
	Description     *string        `pulumi:"description"`
	Properties      *AKSProperties `pulumi:"properties"`
	ResourceId      *string        `pulumi:"resourceId"`
}





type AKSInput interface {
	pulumi.Input

	ToAKSOutput() AKSOutput
	ToAKSOutputWithContext(context.Context) AKSOutput
}

type AKSArgs struct {
	ComputeLocation pulumi.StringPtrInput `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput    `pulumi:"computeType"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Properties      AKSPropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (AKSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKS)(nil)).Elem()
}

func (i AKSArgs) ToAKSOutput() AKSOutput {
	return i.ToAKSOutputWithContext(context.Background())
}

func (i AKSArgs) ToAKSOutputWithContext(ctx context.Context) AKSOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSOutput)
}

type AKSOutput struct{ *pulumi.OutputState }

func (AKSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKS)(nil)).Elem()
}

func (o AKSOutput) ToAKSOutput() AKSOutput {
	return o
}

func (o AKSOutput) ToAKSOutputWithContext(ctx context.Context) AKSOutput {
	return o
}

func (o AKSOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKS) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o AKSOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v AKS) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o AKSOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKS) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AKSOutput) Properties() AKSPropertiesPtrOutput {
	return o.ApplyT(func(v AKS) *AKSProperties { return v.Properties }).(AKSPropertiesPtrOutput)
}

func (o AKSOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKS) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type AKSProperties struct {
	AgentCount                 *int                        `pulumi:"agentCount"`
	AgentVMSize                *string                     `pulumi:"agentVMSize"`
	AksNetworkingConfiguration *AksNetworkingConfiguration `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                *string                     `pulumi:"clusterFqdn"`
	SslConfiguration           *SslConfiguration           `pulumi:"sslConfiguration"`
}





type AKSPropertiesInput interface {
	pulumi.Input

	ToAKSPropertiesOutput() AKSPropertiesOutput
	ToAKSPropertiesOutputWithContext(context.Context) AKSPropertiesOutput
}

type AKSPropertiesArgs struct {
	AgentCount                 pulumi.IntPtrInput                 `pulumi:"agentCount"`
	AgentVMSize                pulumi.StringPtrInput              `pulumi:"agentVMSize"`
	AksNetworkingConfiguration AksNetworkingConfigurationPtrInput `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                pulumi.StringPtrInput              `pulumi:"clusterFqdn"`
	SslConfiguration           SslConfigurationPtrInput           `pulumi:"sslConfiguration"`
}

func (AKSPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSProperties)(nil)).Elem()
}

func (i AKSPropertiesArgs) ToAKSPropertiesOutput() AKSPropertiesOutput {
	return i.ToAKSPropertiesOutputWithContext(context.Background())
}

func (i AKSPropertiesArgs) ToAKSPropertiesOutputWithContext(ctx context.Context) AKSPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSPropertiesOutput)
}

func (i AKSPropertiesArgs) ToAKSPropertiesPtrOutput() AKSPropertiesPtrOutput {
	return i.ToAKSPropertiesPtrOutputWithContext(context.Background())
}

func (i AKSPropertiesArgs) ToAKSPropertiesPtrOutputWithContext(ctx context.Context) AKSPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSPropertiesOutput).ToAKSPropertiesPtrOutputWithContext(ctx)
}









type AKSPropertiesPtrInput interface {
	pulumi.Input

	ToAKSPropertiesPtrOutput() AKSPropertiesPtrOutput
	ToAKSPropertiesPtrOutputWithContext(context.Context) AKSPropertiesPtrOutput
}

type akspropertiesPtrType AKSPropertiesArgs

func AKSPropertiesPtr(v *AKSPropertiesArgs) AKSPropertiesPtrInput {
	return (*akspropertiesPtrType)(v)
}

func (*akspropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSProperties)(nil)).Elem()
}

func (i *akspropertiesPtrType) ToAKSPropertiesPtrOutput() AKSPropertiesPtrOutput {
	return i.ToAKSPropertiesPtrOutputWithContext(context.Background())
}

func (i *akspropertiesPtrType) ToAKSPropertiesPtrOutputWithContext(ctx context.Context) AKSPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSPropertiesPtrOutput)
}

type AKSPropertiesOutput struct{ *pulumi.OutputState }

func (AKSPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSProperties)(nil)).Elem()
}

func (o AKSPropertiesOutput) ToAKSPropertiesOutput() AKSPropertiesOutput {
	return o
}

func (o AKSPropertiesOutput) ToAKSPropertiesOutputWithContext(ctx context.Context) AKSPropertiesOutput {
	return o
}

func (o AKSPropertiesOutput) ToAKSPropertiesPtrOutput() AKSPropertiesPtrOutput {
	return o.ToAKSPropertiesPtrOutputWithContext(context.Background())
}

func (o AKSPropertiesOutput) ToAKSPropertiesPtrOutputWithContext(ctx context.Context) AKSPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSProperties) *AKSProperties {
		return &v
	}).(AKSPropertiesPtrOutput)
}

func (o AKSPropertiesOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSProperties) *int { return v.AgentCount }).(pulumi.IntPtrOutput)
}

func (o AKSPropertiesOutput) AgentVMSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSProperties) *string { return v.AgentVMSize }).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesOutput) AksNetworkingConfiguration() AksNetworkingConfigurationPtrOutput {
	return o.ApplyT(func(v AKSProperties) *AksNetworkingConfiguration { return v.AksNetworkingConfiguration }).(AksNetworkingConfigurationPtrOutput)
}

func (o AKSPropertiesOutput) ClusterFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSProperties) *string { return v.ClusterFqdn }).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesOutput) SslConfiguration() SslConfigurationPtrOutput {
	return o.ApplyT(func(v AKSProperties) *SslConfiguration { return v.SslConfiguration }).(SslConfigurationPtrOutput)
}

type AKSPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AKSPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSProperties)(nil)).Elem()
}

func (o AKSPropertiesPtrOutput) ToAKSPropertiesPtrOutput() AKSPropertiesPtrOutput {
	return o
}

func (o AKSPropertiesPtrOutput) ToAKSPropertiesPtrOutputWithContext(ctx context.Context) AKSPropertiesPtrOutput {
	return o
}

func (o AKSPropertiesPtrOutput) Elem() AKSPropertiesOutput {
	return o.ApplyT(func(v *AKSProperties) AKSProperties {
		if v != nil {
			return *v
		}
		var ret AKSProperties
		return ret
	}).(AKSPropertiesOutput)
}

func (o AKSPropertiesPtrOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *int {
		if v == nil {
			return nil
		}
		return v.AgentCount
	}).(pulumi.IntPtrOutput)
}

func (o AKSPropertiesPtrOutput) AgentVMSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *string {
		if v == nil {
			return nil
		}
		return v.AgentVMSize
	}).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesPtrOutput) AksNetworkingConfiguration() AksNetworkingConfigurationPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *AksNetworkingConfiguration {
		if v == nil {
			return nil
		}
		return v.AksNetworkingConfiguration
	}).(AksNetworkingConfigurationPtrOutput)
}

func (o AKSPropertiesPtrOutput) ClusterFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClusterFqdn
	}).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesPtrOutput) SslConfiguration() SslConfigurationPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *SslConfiguration {
		if v == nil {
			return nil
		}
		return v.SslConfiguration
	}).(SslConfigurationPtrOutput)
}

type AKSReplicaStatusResponseError struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
}





type AKSReplicaStatusResponseErrorInput interface {
	pulumi.Input

	ToAKSReplicaStatusResponseErrorOutput() AKSReplicaStatusResponseErrorOutput
	ToAKSReplicaStatusResponseErrorOutputWithContext(context.Context) AKSReplicaStatusResponseErrorOutput
}

type AKSReplicaStatusResponseErrorArgs struct {
	Code    pulumi.StringInput            `pulumi:"code"`
	Details ErrorDetailResponseArrayInput `pulumi:"details"`
	Message pulumi.StringInput            `pulumi:"message"`
}

func (AKSReplicaStatusResponseErrorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSReplicaStatusResponseError)(nil)).Elem()
}

func (i AKSReplicaStatusResponseErrorArgs) ToAKSReplicaStatusResponseErrorOutput() AKSReplicaStatusResponseErrorOutput {
	return i.ToAKSReplicaStatusResponseErrorOutputWithContext(context.Background())
}

func (i AKSReplicaStatusResponseErrorArgs) ToAKSReplicaStatusResponseErrorOutputWithContext(ctx context.Context) AKSReplicaStatusResponseErrorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSReplicaStatusResponseErrorOutput)
}

func (i AKSReplicaStatusResponseErrorArgs) ToAKSReplicaStatusResponseErrorPtrOutput() AKSReplicaStatusResponseErrorPtrOutput {
	return i.ToAKSReplicaStatusResponseErrorPtrOutputWithContext(context.Background())
}

func (i AKSReplicaStatusResponseErrorArgs) ToAKSReplicaStatusResponseErrorPtrOutputWithContext(ctx context.Context) AKSReplicaStatusResponseErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSReplicaStatusResponseErrorOutput).ToAKSReplicaStatusResponseErrorPtrOutputWithContext(ctx)
}









type AKSReplicaStatusResponseErrorPtrInput interface {
	pulumi.Input

	ToAKSReplicaStatusResponseErrorPtrOutput() AKSReplicaStatusResponseErrorPtrOutput
	ToAKSReplicaStatusResponseErrorPtrOutputWithContext(context.Context) AKSReplicaStatusResponseErrorPtrOutput
}

type aksreplicaStatusResponseErrorPtrType AKSReplicaStatusResponseErrorArgs

func AKSReplicaStatusResponseErrorPtr(v *AKSReplicaStatusResponseErrorArgs) AKSReplicaStatusResponseErrorPtrInput {
	return (*aksreplicaStatusResponseErrorPtrType)(v)
}

func (*aksreplicaStatusResponseErrorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSReplicaStatusResponseError)(nil)).Elem()
}

func (i *aksreplicaStatusResponseErrorPtrType) ToAKSReplicaStatusResponseErrorPtrOutput() AKSReplicaStatusResponseErrorPtrOutput {
	return i.ToAKSReplicaStatusResponseErrorPtrOutputWithContext(context.Background())
}

func (i *aksreplicaStatusResponseErrorPtrType) ToAKSReplicaStatusResponseErrorPtrOutputWithContext(ctx context.Context) AKSReplicaStatusResponseErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSReplicaStatusResponseErrorPtrOutput)
}

type AKSReplicaStatusResponseErrorOutput struct{ *pulumi.OutputState }

func (AKSReplicaStatusResponseErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSReplicaStatusResponseError)(nil)).Elem()
}

func (o AKSReplicaStatusResponseErrorOutput) ToAKSReplicaStatusResponseErrorOutput() AKSReplicaStatusResponseErrorOutput {
	return o
}

func (o AKSReplicaStatusResponseErrorOutput) ToAKSReplicaStatusResponseErrorOutputWithContext(ctx context.Context) AKSReplicaStatusResponseErrorOutput {
	return o
}

func (o AKSReplicaStatusResponseErrorOutput) ToAKSReplicaStatusResponseErrorPtrOutput() AKSReplicaStatusResponseErrorPtrOutput {
	return o.ToAKSReplicaStatusResponseErrorPtrOutputWithContext(context.Background())
}

func (o AKSReplicaStatusResponseErrorOutput) ToAKSReplicaStatusResponseErrorPtrOutputWithContext(ctx context.Context) AKSReplicaStatusResponseErrorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSReplicaStatusResponseError) *AKSReplicaStatusResponseError {
		return &v
	}).(AKSReplicaStatusResponseErrorPtrOutput)
}

func (o AKSReplicaStatusResponseErrorOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AKSReplicaStatusResponseError) string { return v.Code }).(pulumi.StringOutput)
}

func (o AKSReplicaStatusResponseErrorOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v AKSReplicaStatusResponseError) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o AKSReplicaStatusResponseErrorOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v AKSReplicaStatusResponseError) string { return v.Message }).(pulumi.StringOutput)
}

type AKSReplicaStatusResponseErrorPtrOutput struct{ *pulumi.OutputState }

func (AKSReplicaStatusResponseErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSReplicaStatusResponseError)(nil)).Elem()
}

func (o AKSReplicaStatusResponseErrorPtrOutput) ToAKSReplicaStatusResponseErrorPtrOutput() AKSReplicaStatusResponseErrorPtrOutput {
	return o
}

func (o AKSReplicaStatusResponseErrorPtrOutput) ToAKSReplicaStatusResponseErrorPtrOutputWithContext(ctx context.Context) AKSReplicaStatusResponseErrorPtrOutput {
	return o
}

func (o AKSReplicaStatusResponseErrorPtrOutput) Elem() AKSReplicaStatusResponseErrorOutput {
	return o.ApplyT(func(v *AKSReplicaStatusResponseError) AKSReplicaStatusResponseError {
		if v != nil {
			return *v
		}
		var ret AKSReplicaStatusResponseError
		return ret
	}).(AKSReplicaStatusResponseErrorOutput)
}

func (o AKSReplicaStatusResponseErrorPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSReplicaStatusResponseError) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o AKSReplicaStatusResponseErrorPtrOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *AKSReplicaStatusResponseError) []ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ErrorDetailResponseArrayOutput)
}

func (o AKSReplicaStatusResponseErrorPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSReplicaStatusResponseError) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
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





type AKSResponseInput interface {
	pulumi.Input

	ToAKSResponseOutput() AKSResponseOutput
	ToAKSResponseOutputWithContext(context.Context) AKSResponseOutput
}

type AKSResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         AKSResponsePropertiesPtrInput                 `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (AKSResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSResponse)(nil)).Elem()
}

func (i AKSResponseArgs) ToAKSResponseOutput() AKSResponseOutput {
	return i.ToAKSResponseOutputWithContext(context.Background())
}

func (i AKSResponseArgs) ToAKSResponseOutputWithContext(ctx context.Context) AKSResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSResponseOutput)
}

type AKSResponseOutput struct{ *pulumi.OutputState }

func (AKSResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSResponse)(nil)).Elem()
}

func (o AKSResponseOutput) ToAKSResponseOutput() AKSResponseOutput {
	return o
}

func (o AKSResponseOutput) ToAKSResponseOutputWithContext(ctx context.Context) AKSResponseOutput {
	return o
}

func (o AKSResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o AKSResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v AKSResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o AKSResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v AKSResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o AKSResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AKSResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v AKSResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o AKSResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v AKSResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o AKSResponseOutput) Properties() AKSResponsePropertiesPtrOutput {
	return o.ApplyT(func(v AKSResponse) *AKSResponseProperties { return v.Properties }).(AKSResponsePropertiesPtrOutput)
}

func (o AKSResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v AKSResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o AKSResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AKSResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AKSResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type AKSResponseProperties struct {
	AgentCount                 *int                                `pulumi:"agentCount"`
	AgentVMSize                *string                             `pulumi:"agentVMSize"`
	AksNetworkingConfiguration *AksNetworkingConfigurationResponse `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                *string                             `pulumi:"clusterFqdn"`
	SslConfiguration           *SslConfigurationResponse           `pulumi:"sslConfiguration"`
	SystemServices             []SystemServiceResponse             `pulumi:"systemServices"`
}





type AKSResponsePropertiesInput interface {
	pulumi.Input

	ToAKSResponsePropertiesOutput() AKSResponsePropertiesOutput
	ToAKSResponsePropertiesOutputWithContext(context.Context) AKSResponsePropertiesOutput
}

type AKSResponsePropertiesArgs struct {
	AgentCount                 pulumi.IntPtrInput                         `pulumi:"agentCount"`
	AgentVMSize                pulumi.StringPtrInput                      `pulumi:"agentVMSize"`
	AksNetworkingConfiguration AksNetworkingConfigurationResponsePtrInput `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                pulumi.StringPtrInput                      `pulumi:"clusterFqdn"`
	SslConfiguration           SslConfigurationResponsePtrInput           `pulumi:"sslConfiguration"`
	SystemServices             SystemServiceResponseArrayInput            `pulumi:"systemServices"`
}

func (AKSResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSResponseProperties)(nil)).Elem()
}

func (i AKSResponsePropertiesArgs) ToAKSResponsePropertiesOutput() AKSResponsePropertiesOutput {
	return i.ToAKSResponsePropertiesOutputWithContext(context.Background())
}

func (i AKSResponsePropertiesArgs) ToAKSResponsePropertiesOutputWithContext(ctx context.Context) AKSResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSResponsePropertiesOutput)
}

func (i AKSResponsePropertiesArgs) ToAKSResponsePropertiesPtrOutput() AKSResponsePropertiesPtrOutput {
	return i.ToAKSResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i AKSResponsePropertiesArgs) ToAKSResponsePropertiesPtrOutputWithContext(ctx context.Context) AKSResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSResponsePropertiesOutput).ToAKSResponsePropertiesPtrOutputWithContext(ctx)
}









type AKSResponsePropertiesPtrInput interface {
	pulumi.Input

	ToAKSResponsePropertiesPtrOutput() AKSResponsePropertiesPtrOutput
	ToAKSResponsePropertiesPtrOutputWithContext(context.Context) AKSResponsePropertiesPtrOutput
}

type aksresponsePropertiesPtrType AKSResponsePropertiesArgs

func AKSResponsePropertiesPtr(v *AKSResponsePropertiesArgs) AKSResponsePropertiesPtrInput {
	return (*aksresponsePropertiesPtrType)(v)
}

func (*aksresponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSResponseProperties)(nil)).Elem()
}

func (i *aksresponsePropertiesPtrType) ToAKSResponsePropertiesPtrOutput() AKSResponsePropertiesPtrOutput {
	return i.ToAKSResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *aksresponsePropertiesPtrType) ToAKSResponsePropertiesPtrOutputWithContext(ctx context.Context) AKSResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSResponsePropertiesPtrOutput)
}

type AKSResponsePropertiesOutput struct{ *pulumi.OutputState }

func (AKSResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSResponseProperties)(nil)).Elem()
}

func (o AKSResponsePropertiesOutput) ToAKSResponsePropertiesOutput() AKSResponsePropertiesOutput {
	return o
}

func (o AKSResponsePropertiesOutput) ToAKSResponsePropertiesOutputWithContext(ctx context.Context) AKSResponsePropertiesOutput {
	return o
}

func (o AKSResponsePropertiesOutput) ToAKSResponsePropertiesPtrOutput() AKSResponsePropertiesPtrOutput {
	return o.ToAKSResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o AKSResponsePropertiesOutput) ToAKSResponsePropertiesPtrOutputWithContext(ctx context.Context) AKSResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSResponseProperties) *AKSResponseProperties {
		return &v
	}).(AKSResponsePropertiesPtrOutput)
}

func (o AKSResponsePropertiesOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *int { return v.AgentCount }).(pulumi.IntPtrOutput)
}

func (o AKSResponsePropertiesOutput) AgentVMSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *string { return v.AgentVMSize }).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesOutput) AksNetworkingConfiguration() AksNetworkingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *AksNetworkingConfigurationResponse { return v.AksNetworkingConfiguration }).(AksNetworkingConfigurationResponsePtrOutput)
}

func (o AKSResponsePropertiesOutput) ClusterFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *string { return v.ClusterFqdn }).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesOutput) SslConfiguration() SslConfigurationResponsePtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *SslConfigurationResponse { return v.SslConfiguration }).(SslConfigurationResponsePtrOutput)
}

func (o AKSResponsePropertiesOutput) SystemServices() SystemServiceResponseArrayOutput {
	return o.ApplyT(func(v AKSResponseProperties) []SystemServiceResponse { return v.SystemServices }).(SystemServiceResponseArrayOutput)
}

type AKSResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AKSResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSResponseProperties)(nil)).Elem()
}

func (o AKSResponsePropertiesPtrOutput) ToAKSResponsePropertiesPtrOutput() AKSResponsePropertiesPtrOutput {
	return o
}

func (o AKSResponsePropertiesPtrOutput) ToAKSResponsePropertiesPtrOutputWithContext(ctx context.Context) AKSResponsePropertiesPtrOutput {
	return o
}

func (o AKSResponsePropertiesPtrOutput) Elem() AKSResponsePropertiesOutput {
	return o.ApplyT(func(v *AKSResponseProperties) AKSResponseProperties {
		if v != nil {
			return *v
		}
		var ret AKSResponseProperties
		return ret
	}).(AKSResponsePropertiesOutput)
}

func (o AKSResponsePropertiesPtrOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *int {
		if v == nil {
			return nil
		}
		return v.AgentCount
	}).(pulumi.IntPtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) AgentVMSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.AgentVMSize
	}).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) AksNetworkingConfiguration() AksNetworkingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *AksNetworkingConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.AksNetworkingConfiguration
	}).(AksNetworkingConfigurationResponsePtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) ClusterFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClusterFqdn
	}).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) SslConfiguration() SslConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *SslConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.SslConfiguration
	}).(SslConfigurationResponsePtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) SystemServices() SystemServiceResponseArrayOutput {
	return o.ApplyT(func(v *AKSResponseProperties) []SystemServiceResponse {
		if v == nil {
			return nil
		}
		return v.SystemServices
	}).(SystemServiceResponseArrayOutput)
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





type AKSServiceResponseResponseInput interface {
	pulumi.Input

	ToAKSServiceResponseResponseOutput() AKSServiceResponseResponseOutput
	ToAKSServiceResponseResponseOutputWithContext(context.Context) AKSServiceResponseResponseOutput
}

type AKSServiceResponseResponseArgs struct {
	AadAuthEnabled                    pulumi.BoolPtrInput                                         `pulumi:"aadAuthEnabled"`
	AppInsightsEnabled                pulumi.BoolPtrInput                                         `pulumi:"appInsightsEnabled"`
	AuthEnabled                       pulumi.BoolPtrInput                                         `pulumi:"authEnabled"`
	AutoScaler                        AKSServiceResponseResponseAutoScalerPtrInput                `pulumi:"autoScaler"`
	ComputeName                       pulumi.StringPtrInput                                       `pulumi:"computeName"`
	ComputeType                       pulumi.StringInput                                          `pulumi:"computeType"`
	ContainerResourceRequirements     ContainerResourceRequirementsResponsePtrInput               `pulumi:"containerResourceRequirements"`
	DataCollection                    AKSServiceResponseResponseDataCollectionPtrInput            `pulumi:"dataCollection"`
	DeploymentStatus                  AKSServiceResponseResponseDeploymentStatusInput             `pulumi:"deploymentStatus"`
	DeploymentType                    pulumi.StringPtrInput                                       `pulumi:"deploymentType"`
	Description                       pulumi.StringPtrInput                                       `pulumi:"description"`
	EnvironmentImageRequest           AKSServiceResponseResponseEnvironmentImageRequestPtrInput   `pulumi:"environmentImageRequest"`
	Error                             ServiceResponseBaseResponseErrorInput                       `pulumi:"error"`
	IsDefault                         pulumi.BoolPtrInput                                         `pulumi:"isDefault"`
	KvTags                            pulumi.StringMapInput                                       `pulumi:"kvTags"`
	LivenessProbeRequirements         AKSServiceResponseResponseLivenessProbeRequirementsPtrInput `pulumi:"livenessProbeRequirements"`
	MaxConcurrentRequestsPerContainer pulumi.IntPtrInput                                          `pulumi:"maxConcurrentRequestsPerContainer"`
	MaxQueueWaitMs                    pulumi.IntPtrInput                                          `pulumi:"maxQueueWaitMs"`
	ModelConfigMap                    pulumi.MapInput                                             `pulumi:"modelConfigMap"`
	Models                            ModelResponseArrayInput                                     `pulumi:"models"`
	Namespace                         pulumi.StringPtrInput                                       `pulumi:"namespace"`
	NumReplicas                       pulumi.IntPtrInput                                          `pulumi:"numReplicas"`
	Properties                        pulumi.StringMapInput                                       `pulumi:"properties"`
	ScoringTimeoutMs                  pulumi.IntPtrInput                                          `pulumi:"scoringTimeoutMs"`
	ScoringUri                        pulumi.StringInput                                          `pulumi:"scoringUri"`
	State                             pulumi.StringInput                                          `pulumi:"state"`
	SwaggerUri                        pulumi.StringInput                                          `pulumi:"swaggerUri"`
	TrafficPercentile                 pulumi.Float64PtrInput                                      `pulumi:"trafficPercentile"`
	Type                              pulumi.StringPtrInput                                       `pulumi:"type"`
}

func (AKSServiceResponseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponse)(nil)).Elem()
}

func (i AKSServiceResponseResponseArgs) ToAKSServiceResponseResponseOutput() AKSServiceResponseResponseOutput {
	return i.ToAKSServiceResponseResponseOutputWithContext(context.Background())
}

func (i AKSServiceResponseResponseArgs) ToAKSServiceResponseResponseOutputWithContext(ctx context.Context) AKSServiceResponseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseOutput)
}

type AKSServiceResponseResponseOutput struct{ *pulumi.OutputState }

func (AKSServiceResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponse)(nil)).Elem()
}

func (o AKSServiceResponseResponseOutput) ToAKSServiceResponseResponseOutput() AKSServiceResponseResponseOutput {
	return o
}

func (o AKSServiceResponseResponseOutput) ToAKSServiceResponseResponseOutputWithContext(ctx context.Context) AKSServiceResponseResponseOutput {
	return o
}

func (o AKSServiceResponseResponseOutput) AadAuthEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *bool { return v.AadAuthEnabled }).(pulumi.BoolPtrOutput)
}

func (o AKSServiceResponseResponseOutput) AppInsightsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *bool { return v.AppInsightsEnabled }).(pulumi.BoolPtrOutput)
}

func (o AKSServiceResponseResponseOutput) AuthEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *bool { return v.AuthEnabled }).(pulumi.BoolPtrOutput)
}

func (o AKSServiceResponseResponseOutput) AutoScaler() AKSServiceResponseResponseAutoScalerPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *AKSServiceResponseResponseAutoScaler { return v.AutoScaler }).(AKSServiceResponseResponseAutoScalerPtrOutput)
}

func (o AKSServiceResponseResponseOutput) ComputeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *string { return v.ComputeName }).(pulumi.StringPtrOutput)
}

func (o AKSServiceResponseResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o AKSServiceResponseResponseOutput) ContainerResourceRequirements() ContainerResourceRequirementsResponsePtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *ContainerResourceRequirementsResponse {
		return v.ContainerResourceRequirements
	}).(ContainerResourceRequirementsResponsePtrOutput)
}

func (o AKSServiceResponseResponseOutput) DataCollection() AKSServiceResponseResponseDataCollectionPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *AKSServiceResponseResponseDataCollection { return v.DataCollection }).(AKSServiceResponseResponseDataCollectionPtrOutput)
}

func (o AKSServiceResponseResponseOutput) DeploymentStatus() AKSServiceResponseResponseDeploymentStatusOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) AKSServiceResponseResponseDeploymentStatus {
		return v.DeploymentStatus
	}).(AKSServiceResponseResponseDeploymentStatusOutput)
}

func (o AKSServiceResponseResponseOutput) DeploymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *string { return v.DeploymentType }).(pulumi.StringPtrOutput)
}

func (o AKSServiceResponseResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AKSServiceResponseResponseOutput) EnvironmentImageRequest() AKSServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *AKSServiceResponseResponseEnvironmentImageRequest {
		return v.EnvironmentImageRequest
	}).(AKSServiceResponseResponseEnvironmentImageRequestPtrOutput)
}

func (o AKSServiceResponseResponseOutput) Error() ServiceResponseBaseResponseErrorOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) ServiceResponseBaseResponseError { return v.Error }).(ServiceResponseBaseResponseErrorOutput)
}

func (o AKSServiceResponseResponseOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o AKSServiceResponseResponseOutput) KvTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) map[string]string { return v.KvTags }).(pulumi.StringMapOutput)
}

func (o AKSServiceResponseResponseOutput) LivenessProbeRequirements() AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *AKSServiceResponseResponseLivenessProbeRequirements {
		return v.LivenessProbeRequirements
	}).(AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput)
}

func (o AKSServiceResponseResponseOutput) MaxConcurrentRequestsPerContainer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *int { return v.MaxConcurrentRequestsPerContainer }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseOutput) MaxQueueWaitMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *int { return v.MaxQueueWaitMs }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseOutput) ModelConfigMap() pulumi.MapOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) map[string]interface{} { return v.ModelConfigMap }).(pulumi.MapOutput)
}

func (o AKSServiceResponseResponseOutput) Models() ModelResponseArrayOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) []ModelResponse { return v.Models }).(ModelResponseArrayOutput)
}

func (o AKSServiceResponseResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o AKSServiceResponseResponseOutput) NumReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *int { return v.NumReplicas }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o AKSServiceResponseResponseOutput) ScoringTimeoutMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *int { return v.ScoringTimeoutMs }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseOutput) ScoringUri() pulumi.StringOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) string { return v.ScoringUri }).(pulumi.StringOutput)
}

func (o AKSServiceResponseResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o AKSServiceResponseResponseOutput) SwaggerUri() pulumi.StringOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) string { return v.SwaggerUri }).(pulumi.StringOutput)
}

func (o AKSServiceResponseResponseOutput) TrafficPercentile() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *float64 { return v.TrafficPercentile }).(pulumi.Float64PtrOutput)
}

func (o AKSServiceResponseResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AKSServiceResponseResponseAutoScaler struct {
	AutoscaleEnabled       *bool `pulumi:"autoscaleEnabled"`
	MaxReplicas            *int  `pulumi:"maxReplicas"`
	MinReplicas            *int  `pulumi:"minReplicas"`
	RefreshPeriodInSeconds *int  `pulumi:"refreshPeriodInSeconds"`
	TargetUtilization      *int  `pulumi:"targetUtilization"`
}





type AKSServiceResponseResponseAutoScalerInput interface {
	pulumi.Input

	ToAKSServiceResponseResponseAutoScalerOutput() AKSServiceResponseResponseAutoScalerOutput
	ToAKSServiceResponseResponseAutoScalerOutputWithContext(context.Context) AKSServiceResponseResponseAutoScalerOutput
}

type AKSServiceResponseResponseAutoScalerArgs struct {
	AutoscaleEnabled       pulumi.BoolPtrInput `pulumi:"autoscaleEnabled"`
	MaxReplicas            pulumi.IntPtrInput  `pulumi:"maxReplicas"`
	MinReplicas            pulumi.IntPtrInput  `pulumi:"minReplicas"`
	RefreshPeriodInSeconds pulumi.IntPtrInput  `pulumi:"refreshPeriodInSeconds"`
	TargetUtilization      pulumi.IntPtrInput  `pulumi:"targetUtilization"`
}

func (AKSServiceResponseResponseAutoScalerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponseAutoScaler)(nil)).Elem()
}

func (i AKSServiceResponseResponseAutoScalerArgs) ToAKSServiceResponseResponseAutoScalerOutput() AKSServiceResponseResponseAutoScalerOutput {
	return i.ToAKSServiceResponseResponseAutoScalerOutputWithContext(context.Background())
}

func (i AKSServiceResponseResponseAutoScalerArgs) ToAKSServiceResponseResponseAutoScalerOutputWithContext(ctx context.Context) AKSServiceResponseResponseAutoScalerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseAutoScalerOutput)
}

func (i AKSServiceResponseResponseAutoScalerArgs) ToAKSServiceResponseResponseAutoScalerPtrOutput() AKSServiceResponseResponseAutoScalerPtrOutput {
	return i.ToAKSServiceResponseResponseAutoScalerPtrOutputWithContext(context.Background())
}

func (i AKSServiceResponseResponseAutoScalerArgs) ToAKSServiceResponseResponseAutoScalerPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseAutoScalerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseAutoScalerOutput).ToAKSServiceResponseResponseAutoScalerPtrOutputWithContext(ctx)
}









type AKSServiceResponseResponseAutoScalerPtrInput interface {
	pulumi.Input

	ToAKSServiceResponseResponseAutoScalerPtrOutput() AKSServiceResponseResponseAutoScalerPtrOutput
	ToAKSServiceResponseResponseAutoScalerPtrOutputWithContext(context.Context) AKSServiceResponseResponseAutoScalerPtrOutput
}

type aksserviceResponseResponseAutoScalerPtrType AKSServiceResponseResponseAutoScalerArgs

func AKSServiceResponseResponseAutoScalerPtr(v *AKSServiceResponseResponseAutoScalerArgs) AKSServiceResponseResponseAutoScalerPtrInput {
	return (*aksserviceResponseResponseAutoScalerPtrType)(v)
}

func (*aksserviceResponseResponseAutoScalerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceResponseResponseAutoScaler)(nil)).Elem()
}

func (i *aksserviceResponseResponseAutoScalerPtrType) ToAKSServiceResponseResponseAutoScalerPtrOutput() AKSServiceResponseResponseAutoScalerPtrOutput {
	return i.ToAKSServiceResponseResponseAutoScalerPtrOutputWithContext(context.Background())
}

func (i *aksserviceResponseResponseAutoScalerPtrType) ToAKSServiceResponseResponseAutoScalerPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseAutoScalerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseAutoScalerPtrOutput)
}

type AKSServiceResponseResponseAutoScalerOutput struct{ *pulumi.OutputState }

func (AKSServiceResponseResponseAutoScalerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponseAutoScaler)(nil)).Elem()
}

func (o AKSServiceResponseResponseAutoScalerOutput) ToAKSServiceResponseResponseAutoScalerOutput() AKSServiceResponseResponseAutoScalerOutput {
	return o
}

func (o AKSServiceResponseResponseAutoScalerOutput) ToAKSServiceResponseResponseAutoScalerOutputWithContext(ctx context.Context) AKSServiceResponseResponseAutoScalerOutput {
	return o
}

func (o AKSServiceResponseResponseAutoScalerOutput) ToAKSServiceResponseResponseAutoScalerPtrOutput() AKSServiceResponseResponseAutoScalerPtrOutput {
	return o.ToAKSServiceResponseResponseAutoScalerPtrOutputWithContext(context.Background())
}

func (o AKSServiceResponseResponseAutoScalerOutput) ToAKSServiceResponseResponseAutoScalerPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseAutoScalerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSServiceResponseResponseAutoScaler) *AKSServiceResponseResponseAutoScaler {
		return &v
	}).(AKSServiceResponseResponseAutoScalerPtrOutput)
}

func (o AKSServiceResponseResponseAutoScalerOutput) AutoscaleEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseAutoScaler) *bool { return v.AutoscaleEnabled }).(pulumi.BoolPtrOutput)
}

func (o AKSServiceResponseResponseAutoScalerOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseAutoScaler) *int { return v.MaxReplicas }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseAutoScalerOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseAutoScaler) *int { return v.MinReplicas }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseAutoScalerOutput) RefreshPeriodInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseAutoScaler) *int { return v.RefreshPeriodInSeconds }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseAutoScalerOutput) TargetUtilization() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseAutoScaler) *int { return v.TargetUtilization }).(pulumi.IntPtrOutput)
}

type AKSServiceResponseResponseAutoScalerPtrOutput struct{ *pulumi.OutputState }

func (AKSServiceResponseResponseAutoScalerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceResponseResponseAutoScaler)(nil)).Elem()
}

func (o AKSServiceResponseResponseAutoScalerPtrOutput) ToAKSServiceResponseResponseAutoScalerPtrOutput() AKSServiceResponseResponseAutoScalerPtrOutput {
	return o
}

func (o AKSServiceResponseResponseAutoScalerPtrOutput) ToAKSServiceResponseResponseAutoScalerPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseAutoScalerPtrOutput {
	return o
}

func (o AKSServiceResponseResponseAutoScalerPtrOutput) Elem() AKSServiceResponseResponseAutoScalerOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseAutoScaler) AKSServiceResponseResponseAutoScaler {
		if v != nil {
			return *v
		}
		var ret AKSServiceResponseResponseAutoScaler
		return ret
	}).(AKSServiceResponseResponseAutoScalerOutput)
}

func (o AKSServiceResponseResponseAutoScalerPtrOutput) AutoscaleEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseAutoScaler) *bool {
		if v == nil {
			return nil
		}
		return v.AutoscaleEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AKSServiceResponseResponseAutoScalerPtrOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseAutoScaler) *int {
		if v == nil {
			return nil
		}
		return v.MaxReplicas
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseAutoScalerPtrOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseAutoScaler) *int {
		if v == nil {
			return nil
		}
		return v.MinReplicas
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseAutoScalerPtrOutput) RefreshPeriodInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseAutoScaler) *int {
		if v == nil {
			return nil
		}
		return v.RefreshPeriodInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseAutoScalerPtrOutput) TargetUtilization() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseAutoScaler) *int {
		if v == nil {
			return nil
		}
		return v.TargetUtilization
	}).(pulumi.IntPtrOutput)
}

type AKSServiceResponseResponseDataCollection struct {
	EventHubEnabled *bool `pulumi:"eventHubEnabled"`
	StorageEnabled  *bool `pulumi:"storageEnabled"`
}





type AKSServiceResponseResponseDataCollectionInput interface {
	pulumi.Input

	ToAKSServiceResponseResponseDataCollectionOutput() AKSServiceResponseResponseDataCollectionOutput
	ToAKSServiceResponseResponseDataCollectionOutputWithContext(context.Context) AKSServiceResponseResponseDataCollectionOutput
}

type AKSServiceResponseResponseDataCollectionArgs struct {
	EventHubEnabled pulumi.BoolPtrInput `pulumi:"eventHubEnabled"`
	StorageEnabled  pulumi.BoolPtrInput `pulumi:"storageEnabled"`
}

func (AKSServiceResponseResponseDataCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponseDataCollection)(nil)).Elem()
}

func (i AKSServiceResponseResponseDataCollectionArgs) ToAKSServiceResponseResponseDataCollectionOutput() AKSServiceResponseResponseDataCollectionOutput {
	return i.ToAKSServiceResponseResponseDataCollectionOutputWithContext(context.Background())
}

func (i AKSServiceResponseResponseDataCollectionArgs) ToAKSServiceResponseResponseDataCollectionOutputWithContext(ctx context.Context) AKSServiceResponseResponseDataCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseDataCollectionOutput)
}

func (i AKSServiceResponseResponseDataCollectionArgs) ToAKSServiceResponseResponseDataCollectionPtrOutput() AKSServiceResponseResponseDataCollectionPtrOutput {
	return i.ToAKSServiceResponseResponseDataCollectionPtrOutputWithContext(context.Background())
}

func (i AKSServiceResponseResponseDataCollectionArgs) ToAKSServiceResponseResponseDataCollectionPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseDataCollectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseDataCollectionOutput).ToAKSServiceResponseResponseDataCollectionPtrOutputWithContext(ctx)
}









type AKSServiceResponseResponseDataCollectionPtrInput interface {
	pulumi.Input

	ToAKSServiceResponseResponseDataCollectionPtrOutput() AKSServiceResponseResponseDataCollectionPtrOutput
	ToAKSServiceResponseResponseDataCollectionPtrOutputWithContext(context.Context) AKSServiceResponseResponseDataCollectionPtrOutput
}

type aksserviceResponseResponseDataCollectionPtrType AKSServiceResponseResponseDataCollectionArgs

func AKSServiceResponseResponseDataCollectionPtr(v *AKSServiceResponseResponseDataCollectionArgs) AKSServiceResponseResponseDataCollectionPtrInput {
	return (*aksserviceResponseResponseDataCollectionPtrType)(v)
}

func (*aksserviceResponseResponseDataCollectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceResponseResponseDataCollection)(nil)).Elem()
}

func (i *aksserviceResponseResponseDataCollectionPtrType) ToAKSServiceResponseResponseDataCollectionPtrOutput() AKSServiceResponseResponseDataCollectionPtrOutput {
	return i.ToAKSServiceResponseResponseDataCollectionPtrOutputWithContext(context.Background())
}

func (i *aksserviceResponseResponseDataCollectionPtrType) ToAKSServiceResponseResponseDataCollectionPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseDataCollectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseDataCollectionPtrOutput)
}

type AKSServiceResponseResponseDataCollectionOutput struct{ *pulumi.OutputState }

func (AKSServiceResponseResponseDataCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponseDataCollection)(nil)).Elem()
}

func (o AKSServiceResponseResponseDataCollectionOutput) ToAKSServiceResponseResponseDataCollectionOutput() AKSServiceResponseResponseDataCollectionOutput {
	return o
}

func (o AKSServiceResponseResponseDataCollectionOutput) ToAKSServiceResponseResponseDataCollectionOutputWithContext(ctx context.Context) AKSServiceResponseResponseDataCollectionOutput {
	return o
}

func (o AKSServiceResponseResponseDataCollectionOutput) ToAKSServiceResponseResponseDataCollectionPtrOutput() AKSServiceResponseResponseDataCollectionPtrOutput {
	return o.ToAKSServiceResponseResponseDataCollectionPtrOutputWithContext(context.Background())
}

func (o AKSServiceResponseResponseDataCollectionOutput) ToAKSServiceResponseResponseDataCollectionPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseDataCollectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSServiceResponseResponseDataCollection) *AKSServiceResponseResponseDataCollection {
		return &v
	}).(AKSServiceResponseResponseDataCollectionPtrOutput)
}

func (o AKSServiceResponseResponseDataCollectionOutput) EventHubEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseDataCollection) *bool { return v.EventHubEnabled }).(pulumi.BoolPtrOutput)
}

func (o AKSServiceResponseResponseDataCollectionOutput) StorageEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseDataCollection) *bool { return v.StorageEnabled }).(pulumi.BoolPtrOutput)
}

type AKSServiceResponseResponseDataCollectionPtrOutput struct{ *pulumi.OutputState }

func (AKSServiceResponseResponseDataCollectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceResponseResponseDataCollection)(nil)).Elem()
}

func (o AKSServiceResponseResponseDataCollectionPtrOutput) ToAKSServiceResponseResponseDataCollectionPtrOutput() AKSServiceResponseResponseDataCollectionPtrOutput {
	return o
}

func (o AKSServiceResponseResponseDataCollectionPtrOutput) ToAKSServiceResponseResponseDataCollectionPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseDataCollectionPtrOutput {
	return o
}

func (o AKSServiceResponseResponseDataCollectionPtrOutput) Elem() AKSServiceResponseResponseDataCollectionOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseDataCollection) AKSServiceResponseResponseDataCollection {
		if v != nil {
			return *v
		}
		var ret AKSServiceResponseResponseDataCollection
		return ret
	}).(AKSServiceResponseResponseDataCollectionOutput)
}

func (o AKSServiceResponseResponseDataCollectionPtrOutput) EventHubEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseDataCollection) *bool {
		if v == nil {
			return nil
		}
		return v.EventHubEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AKSServiceResponseResponseDataCollectionPtrOutput) StorageEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseDataCollection) *bool {
		if v == nil {
			return nil
		}
		return v.StorageEnabled
	}).(pulumi.BoolPtrOutput)
}

type AKSServiceResponseResponseDeploymentStatus struct {
	AvailableReplicas *int                           `pulumi:"availableReplicas"`
	DesiredReplicas   *int                           `pulumi:"desiredReplicas"`
	Error             *AKSReplicaStatusResponseError `pulumi:"error"`
	UpdatedReplicas   *int                           `pulumi:"updatedReplicas"`
}





type AKSServiceResponseResponseDeploymentStatusInput interface {
	pulumi.Input

	ToAKSServiceResponseResponseDeploymentStatusOutput() AKSServiceResponseResponseDeploymentStatusOutput
	ToAKSServiceResponseResponseDeploymentStatusOutputWithContext(context.Context) AKSServiceResponseResponseDeploymentStatusOutput
}

type AKSServiceResponseResponseDeploymentStatusArgs struct {
	AvailableReplicas pulumi.IntPtrInput                    `pulumi:"availableReplicas"`
	DesiredReplicas   pulumi.IntPtrInput                    `pulumi:"desiredReplicas"`
	Error             AKSReplicaStatusResponseErrorPtrInput `pulumi:"error"`
	UpdatedReplicas   pulumi.IntPtrInput                    `pulumi:"updatedReplicas"`
}

func (AKSServiceResponseResponseDeploymentStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponseDeploymentStatus)(nil)).Elem()
}

func (i AKSServiceResponseResponseDeploymentStatusArgs) ToAKSServiceResponseResponseDeploymentStatusOutput() AKSServiceResponseResponseDeploymentStatusOutput {
	return i.ToAKSServiceResponseResponseDeploymentStatusOutputWithContext(context.Background())
}

func (i AKSServiceResponseResponseDeploymentStatusArgs) ToAKSServiceResponseResponseDeploymentStatusOutputWithContext(ctx context.Context) AKSServiceResponseResponseDeploymentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseDeploymentStatusOutput)
}

type AKSServiceResponseResponseDeploymentStatusOutput struct{ *pulumi.OutputState }

func (AKSServiceResponseResponseDeploymentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponseDeploymentStatus)(nil)).Elem()
}

func (o AKSServiceResponseResponseDeploymentStatusOutput) ToAKSServiceResponseResponseDeploymentStatusOutput() AKSServiceResponseResponseDeploymentStatusOutput {
	return o
}

func (o AKSServiceResponseResponseDeploymentStatusOutput) ToAKSServiceResponseResponseDeploymentStatusOutputWithContext(ctx context.Context) AKSServiceResponseResponseDeploymentStatusOutput {
	return o
}

func (o AKSServiceResponseResponseDeploymentStatusOutput) AvailableReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseDeploymentStatus) *int { return v.AvailableReplicas }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseDeploymentStatusOutput) DesiredReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseDeploymentStatus) *int { return v.DesiredReplicas }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseDeploymentStatusOutput) Error() AKSReplicaStatusResponseErrorPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseDeploymentStatus) *AKSReplicaStatusResponseError { return v.Error }).(AKSReplicaStatusResponseErrorPtrOutput)
}

func (o AKSServiceResponseResponseDeploymentStatusOutput) UpdatedReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseDeploymentStatus) *int { return v.UpdatedReplicas }).(pulumi.IntPtrOutput)
}

type AKSServiceResponseResponseEnvironmentImageRequest struct {
	Assets               []ImageAssetResponse                                  `pulumi:"assets"`
	DriverProgram        *string                                               `pulumi:"driverProgram"`
	Environment          *EnvironmentImageResponseResponseEnvironment          `pulumi:"environment"`
	EnvironmentReference *EnvironmentImageResponseResponseEnvironmentReference `pulumi:"environmentReference"`
	ModelIds             []string                                              `pulumi:"modelIds"`
	Models               []ModelResponse                                       `pulumi:"models"`
}





type AKSServiceResponseResponseEnvironmentImageRequestInput interface {
	pulumi.Input

	ToAKSServiceResponseResponseEnvironmentImageRequestOutput() AKSServiceResponseResponseEnvironmentImageRequestOutput
	ToAKSServiceResponseResponseEnvironmentImageRequestOutputWithContext(context.Context) AKSServiceResponseResponseEnvironmentImageRequestOutput
}

type AKSServiceResponseResponseEnvironmentImageRequestArgs struct {
	Assets               ImageAssetResponseArrayInput                                 `pulumi:"assets"`
	DriverProgram        pulumi.StringPtrInput                                        `pulumi:"driverProgram"`
	Environment          EnvironmentImageResponseResponseEnvironmentPtrInput          `pulumi:"environment"`
	EnvironmentReference EnvironmentImageResponseResponseEnvironmentReferencePtrInput `pulumi:"environmentReference"`
	ModelIds             pulumi.StringArrayInput                                      `pulumi:"modelIds"`
	Models               ModelResponseArrayInput                                      `pulumi:"models"`
}

func (AKSServiceResponseResponseEnvironmentImageRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponseEnvironmentImageRequest)(nil)).Elem()
}

func (i AKSServiceResponseResponseEnvironmentImageRequestArgs) ToAKSServiceResponseResponseEnvironmentImageRequestOutput() AKSServiceResponseResponseEnvironmentImageRequestOutput {
	return i.ToAKSServiceResponseResponseEnvironmentImageRequestOutputWithContext(context.Background())
}

func (i AKSServiceResponseResponseEnvironmentImageRequestArgs) ToAKSServiceResponseResponseEnvironmentImageRequestOutputWithContext(ctx context.Context) AKSServiceResponseResponseEnvironmentImageRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseEnvironmentImageRequestOutput)
}

func (i AKSServiceResponseResponseEnvironmentImageRequestArgs) ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutput() AKSServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return i.ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(context.Background())
}

func (i AKSServiceResponseResponseEnvironmentImageRequestArgs) ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseEnvironmentImageRequestOutput).ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(ctx)
}









type AKSServiceResponseResponseEnvironmentImageRequestPtrInput interface {
	pulumi.Input

	ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutput() AKSServiceResponseResponseEnvironmentImageRequestPtrOutput
	ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(context.Context) AKSServiceResponseResponseEnvironmentImageRequestPtrOutput
}

type aksserviceResponseResponseEnvironmentImageRequestPtrType AKSServiceResponseResponseEnvironmentImageRequestArgs

func AKSServiceResponseResponseEnvironmentImageRequestPtr(v *AKSServiceResponseResponseEnvironmentImageRequestArgs) AKSServiceResponseResponseEnvironmentImageRequestPtrInput {
	return (*aksserviceResponseResponseEnvironmentImageRequestPtrType)(v)
}

func (*aksserviceResponseResponseEnvironmentImageRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceResponseResponseEnvironmentImageRequest)(nil)).Elem()
}

func (i *aksserviceResponseResponseEnvironmentImageRequestPtrType) ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutput() AKSServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return i.ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(context.Background())
}

func (i *aksserviceResponseResponseEnvironmentImageRequestPtrType) ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseEnvironmentImageRequestPtrOutput)
}

type AKSServiceResponseResponseEnvironmentImageRequestOutput struct{ *pulumi.OutputState }

func (AKSServiceResponseResponseEnvironmentImageRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponseEnvironmentImageRequest)(nil)).Elem()
}

func (o AKSServiceResponseResponseEnvironmentImageRequestOutput) ToAKSServiceResponseResponseEnvironmentImageRequestOutput() AKSServiceResponseResponseEnvironmentImageRequestOutput {
	return o
}

func (o AKSServiceResponseResponseEnvironmentImageRequestOutput) ToAKSServiceResponseResponseEnvironmentImageRequestOutputWithContext(ctx context.Context) AKSServiceResponseResponseEnvironmentImageRequestOutput {
	return o
}

func (o AKSServiceResponseResponseEnvironmentImageRequestOutput) ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutput() AKSServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return o.ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(context.Background())
}

func (o AKSServiceResponseResponseEnvironmentImageRequestOutput) ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSServiceResponseResponseEnvironmentImageRequest) *AKSServiceResponseResponseEnvironmentImageRequest {
		return &v
	}).(AKSServiceResponseResponseEnvironmentImageRequestPtrOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestOutput) Assets() ImageAssetResponseArrayOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseEnvironmentImageRequest) []ImageAssetResponse { return v.Assets }).(ImageAssetResponseArrayOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestOutput) DriverProgram() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseEnvironmentImageRequest) *string { return v.DriverProgram }).(pulumi.StringPtrOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestOutput) Environment() EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseEnvironmentImageRequest) *EnvironmentImageResponseResponseEnvironment {
		return v.Environment
	}).(EnvironmentImageResponseResponseEnvironmentPtrOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestOutput) EnvironmentReference() EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseEnvironmentImageRequest) *EnvironmentImageResponseResponseEnvironmentReference {
		return v.EnvironmentReference
	}).(EnvironmentImageResponseResponseEnvironmentReferencePtrOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestOutput) ModelIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseEnvironmentImageRequest) []string { return v.ModelIds }).(pulumi.StringArrayOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestOutput) Models() ModelResponseArrayOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseEnvironmentImageRequest) []ModelResponse { return v.Models }).(ModelResponseArrayOutput)
}

type AKSServiceResponseResponseEnvironmentImageRequestPtrOutput struct{ *pulumi.OutputState }

func (AKSServiceResponseResponseEnvironmentImageRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceResponseResponseEnvironmentImageRequest)(nil)).Elem()
}

func (o AKSServiceResponseResponseEnvironmentImageRequestPtrOutput) ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutput() AKSServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return o
}

func (o AKSServiceResponseResponseEnvironmentImageRequestPtrOutput) ToAKSServiceResponseResponseEnvironmentImageRequestPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseEnvironmentImageRequestPtrOutput {
	return o
}

func (o AKSServiceResponseResponseEnvironmentImageRequestPtrOutput) Elem() AKSServiceResponseResponseEnvironmentImageRequestOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseEnvironmentImageRequest) AKSServiceResponseResponseEnvironmentImageRequest {
		if v != nil {
			return *v
		}
		var ret AKSServiceResponseResponseEnvironmentImageRequest
		return ret
	}).(AKSServiceResponseResponseEnvironmentImageRequestOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestPtrOutput) Assets() ImageAssetResponseArrayOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseEnvironmentImageRequest) []ImageAssetResponse {
		if v == nil {
			return nil
		}
		return v.Assets
	}).(ImageAssetResponseArrayOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestPtrOutput) DriverProgram() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseEnvironmentImageRequest) *string {
		if v == nil {
			return nil
		}
		return v.DriverProgram
	}).(pulumi.StringPtrOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestPtrOutput) Environment() EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseEnvironmentImageRequest) *EnvironmentImageResponseResponseEnvironment {
		if v == nil {
			return nil
		}
		return v.Environment
	}).(EnvironmentImageResponseResponseEnvironmentPtrOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestPtrOutput) EnvironmentReference() EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseEnvironmentImageRequest) *EnvironmentImageResponseResponseEnvironmentReference {
		if v == nil {
			return nil
		}
		return v.EnvironmentReference
	}).(EnvironmentImageResponseResponseEnvironmentReferencePtrOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestPtrOutput) ModelIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseEnvironmentImageRequest) []string {
		if v == nil {
			return nil
		}
		return v.ModelIds
	}).(pulumi.StringArrayOutput)
}

func (o AKSServiceResponseResponseEnvironmentImageRequestPtrOutput) Models() ModelResponseArrayOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseEnvironmentImageRequest) []ModelResponse {
		if v == nil {
			return nil
		}
		return v.Models
	}).(ModelResponseArrayOutput)
}

type AKSServiceResponseResponseLivenessProbeRequirements struct {
	FailureThreshold    *int `pulumi:"failureThreshold"`
	InitialDelaySeconds *int `pulumi:"initialDelaySeconds"`
	PeriodSeconds       *int `pulumi:"periodSeconds"`
	SuccessThreshold    *int `pulumi:"successThreshold"`
	TimeoutSeconds      *int `pulumi:"timeoutSeconds"`
}





type AKSServiceResponseResponseLivenessProbeRequirementsInput interface {
	pulumi.Input

	ToAKSServiceResponseResponseLivenessProbeRequirementsOutput() AKSServiceResponseResponseLivenessProbeRequirementsOutput
	ToAKSServiceResponseResponseLivenessProbeRequirementsOutputWithContext(context.Context) AKSServiceResponseResponseLivenessProbeRequirementsOutput
}

type AKSServiceResponseResponseLivenessProbeRequirementsArgs struct {
	FailureThreshold    pulumi.IntPtrInput `pulumi:"failureThreshold"`
	InitialDelaySeconds pulumi.IntPtrInput `pulumi:"initialDelaySeconds"`
	PeriodSeconds       pulumi.IntPtrInput `pulumi:"periodSeconds"`
	SuccessThreshold    pulumi.IntPtrInput `pulumi:"successThreshold"`
	TimeoutSeconds      pulumi.IntPtrInput `pulumi:"timeoutSeconds"`
}

func (AKSServiceResponseResponseLivenessProbeRequirementsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponseLivenessProbeRequirements)(nil)).Elem()
}

func (i AKSServiceResponseResponseLivenessProbeRequirementsArgs) ToAKSServiceResponseResponseLivenessProbeRequirementsOutput() AKSServiceResponseResponseLivenessProbeRequirementsOutput {
	return i.ToAKSServiceResponseResponseLivenessProbeRequirementsOutputWithContext(context.Background())
}

func (i AKSServiceResponseResponseLivenessProbeRequirementsArgs) ToAKSServiceResponseResponseLivenessProbeRequirementsOutputWithContext(ctx context.Context) AKSServiceResponseResponseLivenessProbeRequirementsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseLivenessProbeRequirementsOutput)
}

func (i AKSServiceResponseResponseLivenessProbeRequirementsArgs) ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutput() AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput {
	return i.ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutputWithContext(context.Background())
}

func (i AKSServiceResponseResponseLivenessProbeRequirementsArgs) ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseLivenessProbeRequirementsOutput).ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutputWithContext(ctx)
}









type AKSServiceResponseResponseLivenessProbeRequirementsPtrInput interface {
	pulumi.Input

	ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutput() AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput
	ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutputWithContext(context.Context) AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput
}

type aksserviceResponseResponseLivenessProbeRequirementsPtrType AKSServiceResponseResponseLivenessProbeRequirementsArgs

func AKSServiceResponseResponseLivenessProbeRequirementsPtr(v *AKSServiceResponseResponseLivenessProbeRequirementsArgs) AKSServiceResponseResponseLivenessProbeRequirementsPtrInput {
	return (*aksserviceResponseResponseLivenessProbeRequirementsPtrType)(v)
}

func (*aksserviceResponseResponseLivenessProbeRequirementsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceResponseResponseLivenessProbeRequirements)(nil)).Elem()
}

func (i *aksserviceResponseResponseLivenessProbeRequirementsPtrType) ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutput() AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput {
	return i.ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutputWithContext(context.Background())
}

func (i *aksserviceResponseResponseLivenessProbeRequirementsPtrType) ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput)
}

type AKSServiceResponseResponseLivenessProbeRequirementsOutput struct{ *pulumi.OutputState }

func (AKSServiceResponseResponseLivenessProbeRequirementsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSServiceResponseResponseLivenessProbeRequirements)(nil)).Elem()
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsOutput) ToAKSServiceResponseResponseLivenessProbeRequirementsOutput() AKSServiceResponseResponseLivenessProbeRequirementsOutput {
	return o
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsOutput) ToAKSServiceResponseResponseLivenessProbeRequirementsOutputWithContext(ctx context.Context) AKSServiceResponseResponseLivenessProbeRequirementsOutput {
	return o
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsOutput) ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutput() AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput {
	return o.ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutputWithContext(context.Background())
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsOutput) ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSServiceResponseResponseLivenessProbeRequirements) *AKSServiceResponseResponseLivenessProbeRequirements {
		return &v
	}).(AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput)
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseLivenessProbeRequirements) *int { return v.FailureThreshold }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsOutput) InitialDelaySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseLivenessProbeRequirements) *int { return v.InitialDelaySeconds }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsOutput) PeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseLivenessProbeRequirements) *int { return v.PeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsOutput) SuccessThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseLivenessProbeRequirements) *int { return v.SuccessThreshold }).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSServiceResponseResponseLivenessProbeRequirements) *int { return v.TimeoutSeconds }).(pulumi.IntPtrOutput)
}

type AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput struct{ *pulumi.OutputState }

func (AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSServiceResponseResponseLivenessProbeRequirements)(nil)).Elem()
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput) ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutput() AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput {
	return o
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput) ToAKSServiceResponseResponseLivenessProbeRequirementsPtrOutputWithContext(ctx context.Context) AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput {
	return o
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput) Elem() AKSServiceResponseResponseLivenessProbeRequirementsOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseLivenessProbeRequirements) AKSServiceResponseResponseLivenessProbeRequirements {
		if v != nil {
			return *v
		}
		var ret AKSServiceResponseResponseLivenessProbeRequirements
		return ret
	}).(AKSServiceResponseResponseLivenessProbeRequirementsOutput)
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseLivenessProbeRequirements) *int {
		if v == nil {
			return nil
		}
		return v.FailureThreshold
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput) InitialDelaySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseLivenessProbeRequirements) *int {
		if v == nil {
			return nil
		}
		return v.InitialDelaySeconds
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput) PeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseLivenessProbeRequirements) *int {
		if v == nil {
			return nil
		}
		return v.PeriodSeconds
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput) SuccessThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseLivenessProbeRequirements) *int {
		if v == nil {
			return nil
		}
		return v.SuccessThreshold
	}).(pulumi.IntPtrOutput)
}

func (o AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSServiceResponseResponseLivenessProbeRequirements) *int {
		if v == nil {
			return nil
		}
		return v.TimeoutSeconds
	}).(pulumi.IntPtrOutput)
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





type AKSVariantResponseResponseInput interface {
	pulumi.Input

	ToAKSVariantResponseResponseOutput() AKSVariantResponseResponseOutput
	ToAKSVariantResponseResponseOutputWithContext(context.Context) AKSVariantResponseResponseOutput
}

type AKSVariantResponseResponseArgs struct {
	ComputeType       pulumi.StringInput                    `pulumi:"computeType"`
	DeploymentType    pulumi.StringPtrInput                 `pulumi:"deploymentType"`
	Description       pulumi.StringPtrInput                 `pulumi:"description"`
	Error             ServiceResponseBaseResponseErrorInput `pulumi:"error"`
	IsDefault         pulumi.BoolPtrInput                   `pulumi:"isDefault"`
	KvTags            pulumi.StringMapInput                 `pulumi:"kvTags"`
	Properties        pulumi.StringMapInput                 `pulumi:"properties"`
	State             pulumi.StringInput                    `pulumi:"state"`
	TrafficPercentile pulumi.Float64PtrInput                `pulumi:"trafficPercentile"`
	Type              pulumi.StringPtrInput                 `pulumi:"type"`
}

func (AKSVariantResponseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSVariantResponseResponse)(nil)).Elem()
}

func (i AKSVariantResponseResponseArgs) ToAKSVariantResponseResponseOutput() AKSVariantResponseResponseOutput {
	return i.ToAKSVariantResponseResponseOutputWithContext(context.Background())
}

func (i AKSVariantResponseResponseArgs) ToAKSVariantResponseResponseOutputWithContext(ctx context.Context) AKSVariantResponseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSVariantResponseResponseOutput)
}

type AKSVariantResponseResponseOutput struct{ *pulumi.OutputState }

func (AKSVariantResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSVariantResponseResponse)(nil)).Elem()
}

func (o AKSVariantResponseResponseOutput) ToAKSVariantResponseResponseOutput() AKSVariantResponseResponseOutput {
	return o
}

func (o AKSVariantResponseResponseOutput) ToAKSVariantResponseResponseOutputWithContext(ctx context.Context) AKSVariantResponseResponseOutput {
	return o
}

func (o AKSVariantResponseResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v AKSVariantResponseResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o AKSVariantResponseResponseOutput) DeploymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSVariantResponseResponse) *string { return v.DeploymentType }).(pulumi.StringPtrOutput)
}

func (o AKSVariantResponseResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSVariantResponseResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AKSVariantResponseResponseOutput) Error() ServiceResponseBaseResponseErrorOutput {
	return o.ApplyT(func(v AKSVariantResponseResponse) ServiceResponseBaseResponseError { return v.Error }).(ServiceResponseBaseResponseErrorOutput)
}

func (o AKSVariantResponseResponseOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSVariantResponseResponse) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o AKSVariantResponseResponseOutput) KvTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v AKSVariantResponseResponse) map[string]string { return v.KvTags }).(pulumi.StringMapOutput)
}

func (o AKSVariantResponseResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v AKSVariantResponseResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o AKSVariantResponseResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AKSVariantResponseResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o AKSVariantResponseResponseOutput) TrafficPercentile() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AKSVariantResponseResponse) *float64 { return v.TrafficPercentile }).(pulumi.Float64PtrOutput)
}

func (o AKSVariantResponseResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSVariantResponseResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AksNetworkingConfiguration struct {
	DnsServiceIP     *string `pulumi:"dnsServiceIP"`
	DockerBridgeCidr *string `pulumi:"dockerBridgeCidr"`
	ServiceCidr      *string `pulumi:"serviceCidr"`
	SubnetId         *string `pulumi:"subnetId"`
}





type AksNetworkingConfigurationInput interface {
	pulumi.Input

	ToAksNetworkingConfigurationOutput() AksNetworkingConfigurationOutput
	ToAksNetworkingConfigurationOutputWithContext(context.Context) AksNetworkingConfigurationOutput
}

type AksNetworkingConfigurationArgs struct {
	DnsServiceIP     pulumi.StringPtrInput `pulumi:"dnsServiceIP"`
	DockerBridgeCidr pulumi.StringPtrInput `pulumi:"dockerBridgeCidr"`
	ServiceCidr      pulumi.StringPtrInput `pulumi:"serviceCidr"`
	SubnetId         pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (AksNetworkingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AksNetworkingConfiguration)(nil)).Elem()
}

func (i AksNetworkingConfigurationArgs) ToAksNetworkingConfigurationOutput() AksNetworkingConfigurationOutput {
	return i.ToAksNetworkingConfigurationOutputWithContext(context.Background())
}

func (i AksNetworkingConfigurationArgs) ToAksNetworkingConfigurationOutputWithContext(ctx context.Context) AksNetworkingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationOutput)
}

func (i AksNetworkingConfigurationArgs) ToAksNetworkingConfigurationPtrOutput() AksNetworkingConfigurationPtrOutput {
	return i.ToAksNetworkingConfigurationPtrOutputWithContext(context.Background())
}

func (i AksNetworkingConfigurationArgs) ToAksNetworkingConfigurationPtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationOutput).ToAksNetworkingConfigurationPtrOutputWithContext(ctx)
}









type AksNetworkingConfigurationPtrInput interface {
	pulumi.Input

	ToAksNetworkingConfigurationPtrOutput() AksNetworkingConfigurationPtrOutput
	ToAksNetworkingConfigurationPtrOutputWithContext(context.Context) AksNetworkingConfigurationPtrOutput
}

type aksNetworkingConfigurationPtrType AksNetworkingConfigurationArgs

func AksNetworkingConfigurationPtr(v *AksNetworkingConfigurationArgs) AksNetworkingConfigurationPtrInput {
	return (*aksNetworkingConfigurationPtrType)(v)
}

func (*aksNetworkingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AksNetworkingConfiguration)(nil)).Elem()
}

func (i *aksNetworkingConfigurationPtrType) ToAksNetworkingConfigurationPtrOutput() AksNetworkingConfigurationPtrOutput {
	return i.ToAksNetworkingConfigurationPtrOutputWithContext(context.Background())
}

func (i *aksNetworkingConfigurationPtrType) ToAksNetworkingConfigurationPtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationPtrOutput)
}

type AksNetworkingConfigurationOutput struct{ *pulumi.OutputState }

func (AksNetworkingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AksNetworkingConfiguration)(nil)).Elem()
}

func (o AksNetworkingConfigurationOutput) ToAksNetworkingConfigurationOutput() AksNetworkingConfigurationOutput {
	return o
}

func (o AksNetworkingConfigurationOutput) ToAksNetworkingConfigurationOutputWithContext(ctx context.Context) AksNetworkingConfigurationOutput {
	return o
}

func (o AksNetworkingConfigurationOutput) ToAksNetworkingConfigurationPtrOutput() AksNetworkingConfigurationPtrOutput {
	return o.ToAksNetworkingConfigurationPtrOutputWithContext(context.Background())
}

func (o AksNetworkingConfigurationOutput) ToAksNetworkingConfigurationPtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AksNetworkingConfiguration) *AksNetworkingConfiguration {
		return &v
	}).(AksNetworkingConfigurationPtrOutput)
}

func (o AksNetworkingConfigurationOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfiguration) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfiguration) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfiguration) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfiguration) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type AksNetworkingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AksNetworkingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AksNetworkingConfiguration)(nil)).Elem()
}

func (o AksNetworkingConfigurationPtrOutput) ToAksNetworkingConfigurationPtrOutput() AksNetworkingConfigurationPtrOutput {
	return o
}

func (o AksNetworkingConfigurationPtrOutput) ToAksNetworkingConfigurationPtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationPtrOutput {
	return o
}

func (o AksNetworkingConfigurationPtrOutput) Elem() AksNetworkingConfigurationOutput {
	return o.ApplyT(func(v *AksNetworkingConfiguration) AksNetworkingConfiguration {
		if v != nil {
			return *v
		}
		var ret AksNetworkingConfiguration
		return ret
	}).(AksNetworkingConfigurationOutput)
}

func (o AksNetworkingConfigurationPtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationPtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationPtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationPtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type AksNetworkingConfigurationResponse struct {
	DnsServiceIP     *string `pulumi:"dnsServiceIP"`
	DockerBridgeCidr *string `pulumi:"dockerBridgeCidr"`
	ServiceCidr      *string `pulumi:"serviceCidr"`
	SubnetId         *string `pulumi:"subnetId"`
}





type AksNetworkingConfigurationResponseInput interface {
	pulumi.Input

	ToAksNetworkingConfigurationResponseOutput() AksNetworkingConfigurationResponseOutput
	ToAksNetworkingConfigurationResponseOutputWithContext(context.Context) AksNetworkingConfigurationResponseOutput
}

type AksNetworkingConfigurationResponseArgs struct {
	DnsServiceIP     pulumi.StringPtrInput `pulumi:"dnsServiceIP"`
	DockerBridgeCidr pulumi.StringPtrInput `pulumi:"dockerBridgeCidr"`
	ServiceCidr      pulumi.StringPtrInput `pulumi:"serviceCidr"`
	SubnetId         pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (AksNetworkingConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AksNetworkingConfigurationResponse)(nil)).Elem()
}

func (i AksNetworkingConfigurationResponseArgs) ToAksNetworkingConfigurationResponseOutput() AksNetworkingConfigurationResponseOutput {
	return i.ToAksNetworkingConfigurationResponseOutputWithContext(context.Background())
}

func (i AksNetworkingConfigurationResponseArgs) ToAksNetworkingConfigurationResponseOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationResponseOutput)
}

func (i AksNetworkingConfigurationResponseArgs) ToAksNetworkingConfigurationResponsePtrOutput() AksNetworkingConfigurationResponsePtrOutput {
	return i.ToAksNetworkingConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i AksNetworkingConfigurationResponseArgs) ToAksNetworkingConfigurationResponsePtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationResponseOutput).ToAksNetworkingConfigurationResponsePtrOutputWithContext(ctx)
}









type AksNetworkingConfigurationResponsePtrInput interface {
	pulumi.Input

	ToAksNetworkingConfigurationResponsePtrOutput() AksNetworkingConfigurationResponsePtrOutput
	ToAksNetworkingConfigurationResponsePtrOutputWithContext(context.Context) AksNetworkingConfigurationResponsePtrOutput
}

type aksNetworkingConfigurationResponsePtrType AksNetworkingConfigurationResponseArgs

func AksNetworkingConfigurationResponsePtr(v *AksNetworkingConfigurationResponseArgs) AksNetworkingConfigurationResponsePtrInput {
	return (*aksNetworkingConfigurationResponsePtrType)(v)
}

func (*aksNetworkingConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AksNetworkingConfigurationResponse)(nil)).Elem()
}

func (i *aksNetworkingConfigurationResponsePtrType) ToAksNetworkingConfigurationResponsePtrOutput() AksNetworkingConfigurationResponsePtrOutput {
	return i.ToAksNetworkingConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *aksNetworkingConfigurationResponsePtrType) ToAksNetworkingConfigurationResponsePtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationResponsePtrOutput)
}

type AksNetworkingConfigurationResponseOutput struct{ *pulumi.OutputState }

func (AksNetworkingConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AksNetworkingConfigurationResponse)(nil)).Elem()
}

func (o AksNetworkingConfigurationResponseOutput) ToAksNetworkingConfigurationResponseOutput() AksNetworkingConfigurationResponseOutput {
	return o
}

func (o AksNetworkingConfigurationResponseOutput) ToAksNetworkingConfigurationResponseOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponseOutput {
	return o
}

func (o AksNetworkingConfigurationResponseOutput) ToAksNetworkingConfigurationResponsePtrOutput() AksNetworkingConfigurationResponsePtrOutput {
	return o.ToAksNetworkingConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o AksNetworkingConfigurationResponseOutput) ToAksNetworkingConfigurationResponsePtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AksNetworkingConfigurationResponse) *AksNetworkingConfigurationResponse {
		return &v
	}).(AksNetworkingConfigurationResponsePtrOutput)
}

func (o AksNetworkingConfigurationResponseOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfigurationResponse) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponseOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfigurationResponse) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponseOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfigurationResponse) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponseOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfigurationResponse) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type AksNetworkingConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (AksNetworkingConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AksNetworkingConfigurationResponse)(nil)).Elem()
}

func (o AksNetworkingConfigurationResponsePtrOutput) ToAksNetworkingConfigurationResponsePtrOutput() AksNetworkingConfigurationResponsePtrOutput {
	return o
}

func (o AksNetworkingConfigurationResponsePtrOutput) ToAksNetworkingConfigurationResponsePtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponsePtrOutput {
	return o
}

func (o AksNetworkingConfigurationResponsePtrOutput) Elem() AksNetworkingConfigurationResponseOutput {
	return o.ApplyT(func(v *AksNetworkingConfigurationResponse) AksNetworkingConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret AksNetworkingConfigurationResponse
		return ret
	}).(AksNetworkingConfigurationResponseOutput)
}

func (o AksNetworkingConfigurationResponsePtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponsePtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponsePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponsePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type AmlCompute struct {
	ComputeLocation *string               `pulumi:"computeLocation"`
	ComputeType     string                `pulumi:"computeType"`
	Description     *string               `pulumi:"description"`
	Properties      *AmlComputeProperties `pulumi:"properties"`
	ResourceId      *string               `pulumi:"resourceId"`
}





type AmlComputeInput interface {
	pulumi.Input

	ToAmlComputeOutput() AmlComputeOutput
	ToAmlComputeOutputWithContext(context.Context) AmlComputeOutput
}

type AmlComputeArgs struct {
	ComputeLocation pulumi.StringPtrInput        `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput           `pulumi:"computeType"`
	Description     pulumi.StringPtrInput        `pulumi:"description"`
	Properties      AmlComputePropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput        `pulumi:"resourceId"`
}

func (AmlComputeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlCompute)(nil)).Elem()
}

func (i AmlComputeArgs) ToAmlComputeOutput() AmlComputeOutput {
	return i.ToAmlComputeOutputWithContext(context.Background())
}

func (i AmlComputeArgs) ToAmlComputeOutputWithContext(ctx context.Context) AmlComputeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeOutput)
}

type AmlComputeOutput struct{ *pulumi.OutputState }

func (AmlComputeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlCompute)(nil)).Elem()
}

func (o AmlComputeOutput) ToAmlComputeOutput() AmlComputeOutput {
	return o
}

func (o AmlComputeOutput) ToAmlComputeOutputWithContext(ctx context.Context) AmlComputeOutput {
	return o
}

func (o AmlComputeOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlCompute) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o AmlComputeOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v AmlCompute) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o AmlComputeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlCompute) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AmlComputeOutput) Properties() AmlComputePropertiesPtrOutput {
	return o.ApplyT(func(v AmlCompute) *AmlComputeProperties { return v.Properties }).(AmlComputePropertiesPtrOutput)
}

func (o AmlComputeOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlCompute) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type AmlComputeNodeInformationResponse struct {
	NodeId           string  `pulumi:"nodeId"`
	NodeState        string  `pulumi:"nodeState"`
	Port             float64 `pulumi:"port"`
	PrivateIpAddress string  `pulumi:"privateIpAddress"`
	PublicIpAddress  string  `pulumi:"publicIpAddress"`
	RunId            string  `pulumi:"runId"`
}





type AmlComputeNodeInformationResponseInput interface {
	pulumi.Input

	ToAmlComputeNodeInformationResponseOutput() AmlComputeNodeInformationResponseOutput
	ToAmlComputeNodeInformationResponseOutputWithContext(context.Context) AmlComputeNodeInformationResponseOutput
}

type AmlComputeNodeInformationResponseArgs struct {
	NodeId           pulumi.StringInput  `pulumi:"nodeId"`
	NodeState        pulumi.StringInput  `pulumi:"nodeState"`
	Port             pulumi.Float64Input `pulumi:"port"`
	PrivateIpAddress pulumi.StringInput  `pulumi:"privateIpAddress"`
	PublicIpAddress  pulumi.StringInput  `pulumi:"publicIpAddress"`
	RunId            pulumi.StringInput  `pulumi:"runId"`
}

func (AmlComputeNodeInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeNodeInformationResponse)(nil)).Elem()
}

func (i AmlComputeNodeInformationResponseArgs) ToAmlComputeNodeInformationResponseOutput() AmlComputeNodeInformationResponseOutput {
	return i.ToAmlComputeNodeInformationResponseOutputWithContext(context.Background())
}

func (i AmlComputeNodeInformationResponseArgs) ToAmlComputeNodeInformationResponseOutputWithContext(ctx context.Context) AmlComputeNodeInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeNodeInformationResponseOutput)
}





type AmlComputeNodeInformationResponseArrayInput interface {
	pulumi.Input

	ToAmlComputeNodeInformationResponseArrayOutput() AmlComputeNodeInformationResponseArrayOutput
	ToAmlComputeNodeInformationResponseArrayOutputWithContext(context.Context) AmlComputeNodeInformationResponseArrayOutput
}

type AmlComputeNodeInformationResponseArray []AmlComputeNodeInformationResponseInput

func (AmlComputeNodeInformationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AmlComputeNodeInformationResponse)(nil)).Elem()
}

func (i AmlComputeNodeInformationResponseArray) ToAmlComputeNodeInformationResponseArrayOutput() AmlComputeNodeInformationResponseArrayOutput {
	return i.ToAmlComputeNodeInformationResponseArrayOutputWithContext(context.Background())
}

func (i AmlComputeNodeInformationResponseArray) ToAmlComputeNodeInformationResponseArrayOutputWithContext(ctx context.Context) AmlComputeNodeInformationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeNodeInformationResponseArrayOutput)
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
	RemoteLoginPortPublicAccess *string                 `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings               *ScaleSettings          `pulumi:"scaleSettings"`
	Subnet                      *ResourceId             `pulumi:"subnet"`
	UserAccountCredentials      *UserAccountCredentials `pulumi:"userAccountCredentials"`
	VirtualMachineImage         *VirtualMachineImage    `pulumi:"virtualMachineImage"`
	VmPriority                  *string                 `pulumi:"vmPriority"`
	VmSize                      *string                 `pulumi:"vmSize"`
}





type AmlComputePropertiesInput interface {
	pulumi.Input

	ToAmlComputePropertiesOutput() AmlComputePropertiesOutput
	ToAmlComputePropertiesOutputWithContext(context.Context) AmlComputePropertiesOutput
}

type AmlComputePropertiesArgs struct {
	EnableNodePublicIp          pulumi.BoolPtrInput            `pulumi:"enableNodePublicIp"`
	IsolatedNetwork             pulumi.BoolPtrInput            `pulumi:"isolatedNetwork"`
	OsType                      pulumi.StringPtrInput          `pulumi:"osType"`
	RemoteLoginPortPublicAccess pulumi.StringPtrInput          `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings               ScaleSettingsPtrInput          `pulumi:"scaleSettings"`
	Subnet                      ResourceIdPtrInput             `pulumi:"subnet"`
	UserAccountCredentials      UserAccountCredentialsPtrInput `pulumi:"userAccountCredentials"`
	VirtualMachineImage         VirtualMachineImagePtrInput    `pulumi:"virtualMachineImage"`
	VmPriority                  pulumi.StringPtrInput          `pulumi:"vmPriority"`
	VmSize                      pulumi.StringPtrInput          `pulumi:"vmSize"`
}

func (AmlComputePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeProperties)(nil)).Elem()
}

func (i AmlComputePropertiesArgs) ToAmlComputePropertiesOutput() AmlComputePropertiesOutput {
	return i.ToAmlComputePropertiesOutputWithContext(context.Background())
}

func (i AmlComputePropertiesArgs) ToAmlComputePropertiesOutputWithContext(ctx context.Context) AmlComputePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputePropertiesOutput)
}

func (i AmlComputePropertiesArgs) ToAmlComputePropertiesPtrOutput() AmlComputePropertiesPtrOutput {
	return i.ToAmlComputePropertiesPtrOutputWithContext(context.Background())
}

func (i AmlComputePropertiesArgs) ToAmlComputePropertiesPtrOutputWithContext(ctx context.Context) AmlComputePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputePropertiesOutput).ToAmlComputePropertiesPtrOutputWithContext(ctx)
}









type AmlComputePropertiesPtrInput interface {
	pulumi.Input

	ToAmlComputePropertiesPtrOutput() AmlComputePropertiesPtrOutput
	ToAmlComputePropertiesPtrOutputWithContext(context.Context) AmlComputePropertiesPtrOutput
}

type amlComputePropertiesPtrType AmlComputePropertiesArgs

func AmlComputePropertiesPtr(v *AmlComputePropertiesArgs) AmlComputePropertiesPtrInput {
	return (*amlComputePropertiesPtrType)(v)
}

func (*amlComputePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AmlComputeProperties)(nil)).Elem()
}

func (i *amlComputePropertiesPtrType) ToAmlComputePropertiesPtrOutput() AmlComputePropertiesPtrOutput {
	return i.ToAmlComputePropertiesPtrOutputWithContext(context.Background())
}

func (i *amlComputePropertiesPtrType) ToAmlComputePropertiesPtrOutputWithContext(ctx context.Context) AmlComputePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputePropertiesPtrOutput)
}

type AmlComputePropertiesOutput struct{ *pulumi.OutputState }

func (AmlComputePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeProperties)(nil)).Elem()
}

func (o AmlComputePropertiesOutput) ToAmlComputePropertiesOutput() AmlComputePropertiesOutput {
	return o
}

func (o AmlComputePropertiesOutput) ToAmlComputePropertiesOutputWithContext(ctx context.Context) AmlComputePropertiesOutput {
	return o
}

func (o AmlComputePropertiesOutput) ToAmlComputePropertiesPtrOutput() AmlComputePropertiesPtrOutput {
	return o.ToAmlComputePropertiesPtrOutputWithContext(context.Background())
}

func (o AmlComputePropertiesOutput) ToAmlComputePropertiesPtrOutputWithContext(ctx context.Context) AmlComputePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AmlComputeProperties) *AmlComputeProperties {
		return &v
	}).(AmlComputePropertiesPtrOutput)
}

func (o AmlComputePropertiesOutput) EnableNodePublicIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *bool { return v.EnableNodePublicIp }).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesOutput) IsolatedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *bool { return v.IsolatedNetwork }).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesOutput) RemoteLoginPortPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *string { return v.RemoteLoginPortPublicAccess }).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesOutput) ScaleSettings() ScaleSettingsPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *ScaleSettings { return v.ScaleSettings }).(ScaleSettingsPtrOutput)
}

func (o AmlComputePropertiesOutput) Subnet() ResourceIdPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *ResourceId { return v.Subnet }).(ResourceIdPtrOutput)
}

func (o AmlComputePropertiesOutput) UserAccountCredentials() UserAccountCredentialsPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *UserAccountCredentials { return v.UserAccountCredentials }).(UserAccountCredentialsPtrOutput)
}

func (o AmlComputePropertiesOutput) VirtualMachineImage() VirtualMachineImagePtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *VirtualMachineImage { return v.VirtualMachineImage }).(VirtualMachineImagePtrOutput)
}

func (o AmlComputePropertiesOutput) VmPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *string { return v.VmPriority }).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type AmlComputePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AmlComputePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AmlComputeProperties)(nil)).Elem()
}

func (o AmlComputePropertiesPtrOutput) ToAmlComputePropertiesPtrOutput() AmlComputePropertiesPtrOutput {
	return o
}

func (o AmlComputePropertiesPtrOutput) ToAmlComputePropertiesPtrOutputWithContext(ctx context.Context) AmlComputePropertiesPtrOutput {
	return o
}

func (o AmlComputePropertiesPtrOutput) Elem() AmlComputePropertiesOutput {
	return o.ApplyT(func(v *AmlComputeProperties) AmlComputeProperties {
		if v != nil {
			return *v
		}
		var ret AmlComputeProperties
		return ret
	}).(AmlComputePropertiesOutput)
}

func (o AmlComputePropertiesPtrOutput) EnableNodePublicIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNodePublicIp
	}).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) IsolatedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsolatedNetwork
	}).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) RemoteLoginPortPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *string {
		if v == nil {
			return nil
		}
		return v.RemoteLoginPortPublicAccess
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) ScaleSettings() ScaleSettingsPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *ScaleSettings {
		if v == nil {
			return nil
		}
		return v.ScaleSettings
	}).(ScaleSettingsPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) Subnet() ResourceIdPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *ResourceId {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(ResourceIdPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) UserAccountCredentials() UserAccountCredentialsPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *UserAccountCredentials {
		if v == nil {
			return nil
		}
		return v.UserAccountCredentials
	}).(UserAccountCredentialsPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) VirtualMachineImage() VirtualMachineImagePtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *VirtualMachineImage {
		if v == nil {
			return nil
		}
		return v.VirtualMachineImage
	}).(VirtualMachineImagePtrOutput)
}

func (o AmlComputePropertiesPtrOutput) VmPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *string {
		if v == nil {
			return nil
		}
		return v.VmPriority
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
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





type AmlComputeResponseInput interface {
	pulumi.Input

	ToAmlComputeResponseOutput() AmlComputeResponseOutput
	ToAmlComputeResponseOutputWithContext(context.Context) AmlComputeResponseOutput
}

type AmlComputeResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         AmlComputeResponsePropertiesPtrInput          `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (AmlComputeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeResponse)(nil)).Elem()
}

func (i AmlComputeResponseArgs) ToAmlComputeResponseOutput() AmlComputeResponseOutput {
	return i.ToAmlComputeResponseOutputWithContext(context.Background())
}

func (i AmlComputeResponseArgs) ToAmlComputeResponseOutputWithContext(ctx context.Context) AmlComputeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeResponseOutput)
}

type AmlComputeResponseOutput struct{ *pulumi.OutputState }

func (AmlComputeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeResponse)(nil)).Elem()
}

func (o AmlComputeResponseOutput) ToAmlComputeResponseOutput() AmlComputeResponseOutput {
	return o
}

func (o AmlComputeResponseOutput) ToAmlComputeResponseOutputWithContext(ctx context.Context) AmlComputeResponseOutput {
	return o
}

func (o AmlComputeResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o AmlComputeResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o AmlComputeResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v AmlComputeResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o AmlComputeResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o AmlComputeResponseOutput) Properties() AmlComputeResponsePropertiesPtrOutput {
	return o.ApplyT(func(v AmlComputeResponse) *AmlComputeResponseProperties { return v.Properties }).(AmlComputeResponsePropertiesPtrOutput)
}

func (o AmlComputeResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v AmlComputeResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o AmlComputeResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AmlComputeResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
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





type AmlComputeResponsePropertiesInput interface {
	pulumi.Input

	ToAmlComputeResponsePropertiesOutput() AmlComputeResponsePropertiesOutput
	ToAmlComputeResponsePropertiesOutputWithContext(context.Context) AmlComputeResponsePropertiesOutput
}

type AmlComputeResponsePropertiesArgs struct {
	AllocationState               pulumi.StringInput                            `pulumi:"allocationState"`
	AllocationStateTransitionTime pulumi.StringInput                            `pulumi:"allocationStateTransitionTime"`
	CurrentNodeCount              pulumi.IntInput                               `pulumi:"currentNodeCount"`
	EnableNodePublicIp            pulumi.BoolPtrInput                           `pulumi:"enableNodePublicIp"`
	Errors                        MachineLearningServiceErrorResponseArrayInput `pulumi:"errors"`
	IsolatedNetwork               pulumi.BoolPtrInput                           `pulumi:"isolatedNetwork"`
	NodeStateCounts               NodeStateCountsResponseInput                  `pulumi:"nodeStateCounts"`
	OsType                        pulumi.StringPtrInput                         `pulumi:"osType"`
	RemoteLoginPortPublicAccess   pulumi.StringPtrInput                         `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings                 ScaleSettingsResponsePtrInput                 `pulumi:"scaleSettings"`
	Subnet                        ResourceIdResponsePtrInput                    `pulumi:"subnet"`
	TargetNodeCount               pulumi.IntInput                               `pulumi:"targetNodeCount"`
	UserAccountCredentials        UserAccountCredentialsResponsePtrInput        `pulumi:"userAccountCredentials"`
	VirtualMachineImage           VirtualMachineImageResponsePtrInput           `pulumi:"virtualMachineImage"`
	VmPriority                    pulumi.StringPtrInput                         `pulumi:"vmPriority"`
	VmSize                        pulumi.StringPtrInput                         `pulumi:"vmSize"`
}

func (AmlComputeResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeResponseProperties)(nil)).Elem()
}

func (i AmlComputeResponsePropertiesArgs) ToAmlComputeResponsePropertiesOutput() AmlComputeResponsePropertiesOutput {
	return i.ToAmlComputeResponsePropertiesOutputWithContext(context.Background())
}

func (i AmlComputeResponsePropertiesArgs) ToAmlComputeResponsePropertiesOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeResponsePropertiesOutput)
}

func (i AmlComputeResponsePropertiesArgs) ToAmlComputeResponsePropertiesPtrOutput() AmlComputeResponsePropertiesPtrOutput {
	return i.ToAmlComputeResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i AmlComputeResponsePropertiesArgs) ToAmlComputeResponsePropertiesPtrOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeResponsePropertiesOutput).ToAmlComputeResponsePropertiesPtrOutputWithContext(ctx)
}









type AmlComputeResponsePropertiesPtrInput interface {
	pulumi.Input

	ToAmlComputeResponsePropertiesPtrOutput() AmlComputeResponsePropertiesPtrOutput
	ToAmlComputeResponsePropertiesPtrOutputWithContext(context.Context) AmlComputeResponsePropertiesPtrOutput
}

type amlComputeResponsePropertiesPtrType AmlComputeResponsePropertiesArgs

func AmlComputeResponsePropertiesPtr(v *AmlComputeResponsePropertiesArgs) AmlComputeResponsePropertiesPtrInput {
	return (*amlComputeResponsePropertiesPtrType)(v)
}

func (*amlComputeResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AmlComputeResponseProperties)(nil)).Elem()
}

func (i *amlComputeResponsePropertiesPtrType) ToAmlComputeResponsePropertiesPtrOutput() AmlComputeResponsePropertiesPtrOutput {
	return i.ToAmlComputeResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *amlComputeResponsePropertiesPtrType) ToAmlComputeResponsePropertiesPtrOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeResponsePropertiesPtrOutput)
}

type AmlComputeResponsePropertiesOutput struct{ *pulumi.OutputState }

func (AmlComputeResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeResponseProperties)(nil)).Elem()
}

func (o AmlComputeResponsePropertiesOutput) ToAmlComputeResponsePropertiesOutput() AmlComputeResponsePropertiesOutput {
	return o
}

func (o AmlComputeResponsePropertiesOutput) ToAmlComputeResponsePropertiesOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesOutput {
	return o
}

func (o AmlComputeResponsePropertiesOutput) ToAmlComputeResponsePropertiesPtrOutput() AmlComputeResponsePropertiesPtrOutput {
	return o.ToAmlComputeResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o AmlComputeResponsePropertiesOutput) ToAmlComputeResponsePropertiesPtrOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AmlComputeResponseProperties) *AmlComputeResponseProperties {
		return &v
	}).(AmlComputeResponsePropertiesPtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) AllocationState() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) string { return v.AllocationState }).(pulumi.StringOutput)
}

func (o AmlComputeResponsePropertiesOutput) AllocationStateTransitionTime() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) string { return v.AllocationStateTransitionTime }).(pulumi.StringOutput)
}

func (o AmlComputeResponsePropertiesOutput) CurrentNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) int { return v.CurrentNodeCount }).(pulumi.IntOutput)
}

func (o AmlComputeResponsePropertiesOutput) EnableNodePublicIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *bool { return v.EnableNodePublicIp }).(pulumi.BoolPtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) Errors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) []MachineLearningServiceErrorResponse { return v.Errors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o AmlComputeResponsePropertiesOutput) IsolatedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *bool { return v.IsolatedNetwork }).(pulumi.BoolPtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) NodeStateCounts() NodeStateCountsResponseOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) NodeStateCountsResponse { return v.NodeStateCounts }).(NodeStateCountsResponseOutput)
}

func (o AmlComputeResponsePropertiesOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) RemoteLoginPortPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *string { return v.RemoteLoginPortPublicAccess }).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) ScaleSettings() ScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *ScaleSettingsResponse { return v.ScaleSettings }).(ScaleSettingsResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) Subnet() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *ResourceIdResponse { return v.Subnet }).(ResourceIdResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) TargetNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) int { return v.TargetNodeCount }).(pulumi.IntOutput)
}

func (o AmlComputeResponsePropertiesOutput) UserAccountCredentials() UserAccountCredentialsResponsePtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *UserAccountCredentialsResponse { return v.UserAccountCredentials }).(UserAccountCredentialsResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) VirtualMachineImage() VirtualMachineImageResponsePtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *VirtualMachineImageResponse { return v.VirtualMachineImage }).(VirtualMachineImageResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) VmPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *string { return v.VmPriority }).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type AmlComputeResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AmlComputeResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AmlComputeResponseProperties)(nil)).Elem()
}

func (o AmlComputeResponsePropertiesPtrOutput) ToAmlComputeResponsePropertiesPtrOutput() AmlComputeResponsePropertiesPtrOutput {
	return o
}

func (o AmlComputeResponsePropertiesPtrOutput) ToAmlComputeResponsePropertiesPtrOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesPtrOutput {
	return o
}

func (o AmlComputeResponsePropertiesPtrOutput) Elem() AmlComputeResponsePropertiesOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) AmlComputeResponseProperties {
		if v != nil {
			return *v
		}
		var ret AmlComputeResponseProperties
		return ret
	}).(AmlComputeResponsePropertiesOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) AllocationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AllocationState
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) AllocationStateTransitionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AllocationStateTransitionTime
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) CurrentNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *int {
		if v == nil {
			return nil
		}
		return &v.CurrentNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) EnableNodePublicIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNodePublicIp
	}).(pulumi.BoolPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) Errors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) []MachineLearningServiceErrorResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) IsolatedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsolatedNetwork
	}).(pulumi.BoolPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) NodeStateCounts() NodeStateCountsResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *NodeStateCountsResponse {
		if v == nil {
			return nil
		}
		return &v.NodeStateCounts
	}).(NodeStateCountsResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) RemoteLoginPortPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.RemoteLoginPortPublicAccess
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) ScaleSettings() ScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *ScaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ScaleSettings
	}).(ScaleSettingsResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) Subnet() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *ResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(ResourceIdResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) TargetNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *int {
		if v == nil {
			return nil
		}
		return &v.TargetNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) UserAccountCredentials() UserAccountCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *UserAccountCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.UserAccountCredentials
	}).(UserAccountCredentialsResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) VirtualMachineImage() VirtualMachineImageResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *VirtualMachineImageResponse {
		if v == nil {
			return nil
		}
		return v.VirtualMachineImage
	}).(VirtualMachineImageResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) VmPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.VmPriority
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
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





type ContainerResourceRequirementsResponseInput interface {
	pulumi.Input

	ToContainerResourceRequirementsResponseOutput() ContainerResourceRequirementsResponseOutput
	ToContainerResourceRequirementsResponseOutputWithContext(context.Context) ContainerResourceRequirementsResponseOutput
}

type ContainerResourceRequirementsResponseArgs struct {
	Cpu        pulumi.Float64PtrInput `pulumi:"cpu"`
	Fpga       pulumi.IntPtrInput     `pulumi:"fpga"`
	Gpu        pulumi.IntPtrInput     `pulumi:"gpu"`
	MemoryInGB pulumi.Float64PtrInput `pulumi:"memoryInGB"`
}

func (ContainerResourceRequirementsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResourceRequirementsResponse)(nil)).Elem()
}

func (i ContainerResourceRequirementsResponseArgs) ToContainerResourceRequirementsResponseOutput() ContainerResourceRequirementsResponseOutput {
	return i.ToContainerResourceRequirementsResponseOutputWithContext(context.Background())
}

func (i ContainerResourceRequirementsResponseArgs) ToContainerResourceRequirementsResponseOutputWithContext(ctx context.Context) ContainerResourceRequirementsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourceRequirementsResponseOutput)
}

func (i ContainerResourceRequirementsResponseArgs) ToContainerResourceRequirementsResponsePtrOutput() ContainerResourceRequirementsResponsePtrOutput {
	return i.ToContainerResourceRequirementsResponsePtrOutputWithContext(context.Background())
}

func (i ContainerResourceRequirementsResponseArgs) ToContainerResourceRequirementsResponsePtrOutputWithContext(ctx context.Context) ContainerResourceRequirementsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourceRequirementsResponseOutput).ToContainerResourceRequirementsResponsePtrOutputWithContext(ctx)
}









type ContainerResourceRequirementsResponsePtrInput interface {
	pulumi.Input

	ToContainerResourceRequirementsResponsePtrOutput() ContainerResourceRequirementsResponsePtrOutput
	ToContainerResourceRequirementsResponsePtrOutputWithContext(context.Context) ContainerResourceRequirementsResponsePtrOutput
}

type containerResourceRequirementsResponsePtrType ContainerResourceRequirementsResponseArgs

func ContainerResourceRequirementsResponsePtr(v *ContainerResourceRequirementsResponseArgs) ContainerResourceRequirementsResponsePtrInput {
	return (*containerResourceRequirementsResponsePtrType)(v)
}

func (*containerResourceRequirementsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerResourceRequirementsResponse)(nil)).Elem()
}

func (i *containerResourceRequirementsResponsePtrType) ToContainerResourceRequirementsResponsePtrOutput() ContainerResourceRequirementsResponsePtrOutput {
	return i.ToContainerResourceRequirementsResponsePtrOutputWithContext(context.Background())
}

func (i *containerResourceRequirementsResponsePtrType) ToContainerResourceRequirementsResponsePtrOutputWithContext(ctx context.Context) ContainerResourceRequirementsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourceRequirementsResponsePtrOutput)
}

type ContainerResourceRequirementsResponseOutput struct{ *pulumi.OutputState }

func (ContainerResourceRequirementsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResourceRequirementsResponse)(nil)).Elem()
}

func (o ContainerResourceRequirementsResponseOutput) ToContainerResourceRequirementsResponseOutput() ContainerResourceRequirementsResponseOutput {
	return o
}

func (o ContainerResourceRequirementsResponseOutput) ToContainerResourceRequirementsResponseOutputWithContext(ctx context.Context) ContainerResourceRequirementsResponseOutput {
	return o
}

func (o ContainerResourceRequirementsResponseOutput) ToContainerResourceRequirementsResponsePtrOutput() ContainerResourceRequirementsResponsePtrOutput {
	return o.ToContainerResourceRequirementsResponsePtrOutputWithContext(context.Background())
}

func (o ContainerResourceRequirementsResponseOutput) ToContainerResourceRequirementsResponsePtrOutputWithContext(ctx context.Context) ContainerResourceRequirementsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerResourceRequirementsResponse) *ContainerResourceRequirementsResponse {
		return &v
	}).(ContainerResourceRequirementsResponsePtrOutput)
}

func (o ContainerResourceRequirementsResponseOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerResourceRequirementsResponse) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ContainerResourceRequirementsResponseOutput) Fpga() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerResourceRequirementsResponse) *int { return v.Fpga }).(pulumi.IntPtrOutput)
}

func (o ContainerResourceRequirementsResponseOutput) Gpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerResourceRequirementsResponse) *int { return v.Gpu }).(pulumi.IntPtrOutput)
}

func (o ContainerResourceRequirementsResponseOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerResourceRequirementsResponse) *float64 { return v.MemoryInGB }).(pulumi.Float64PtrOutput)
}

type ContainerResourceRequirementsResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerResourceRequirementsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerResourceRequirementsResponse)(nil)).Elem()
}

func (o ContainerResourceRequirementsResponsePtrOutput) ToContainerResourceRequirementsResponsePtrOutput() ContainerResourceRequirementsResponsePtrOutput {
	return o
}

func (o ContainerResourceRequirementsResponsePtrOutput) ToContainerResourceRequirementsResponsePtrOutputWithContext(ctx context.Context) ContainerResourceRequirementsResponsePtrOutput {
	return o
}

func (o ContainerResourceRequirementsResponsePtrOutput) Elem() ContainerResourceRequirementsResponseOutput {
	return o.ApplyT(func(v *ContainerResourceRequirementsResponse) ContainerResourceRequirementsResponse {
		if v != nil {
			return *v
		}
		var ret ContainerResourceRequirementsResponse
		return ret
	}).(ContainerResourceRequirementsResponseOutput)
}

func (o ContainerResourceRequirementsResponsePtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContainerResourceRequirementsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ContainerResourceRequirementsResponsePtrOutput) Fpga() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerResourceRequirementsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Fpga
	}).(pulumi.IntPtrOutput)
}

func (o ContainerResourceRequirementsResponsePtrOutput) Gpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerResourceRequirementsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Gpu
	}).(pulumi.IntPtrOutput)
}

func (o ContainerResourceRequirementsResponsePtrOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContainerResourceRequirementsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.Float64PtrOutput)
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





type DataFactoryInput interface {
	pulumi.Input

	ToDataFactoryOutput() DataFactoryOutput
	ToDataFactoryOutputWithContext(context.Context) DataFactoryOutput
}

type DataFactoryArgs struct {
	ComputeLocation pulumi.StringPtrInput `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput    `pulumi:"computeType"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	ResourceId      pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (DataFactoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFactory)(nil)).Elem()
}

func (i DataFactoryArgs) ToDataFactoryOutput() DataFactoryOutput {
	return i.ToDataFactoryOutputWithContext(context.Background())
}

func (i DataFactoryArgs) ToDataFactoryOutputWithContext(ctx context.Context) DataFactoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFactoryOutput)
}

type DataFactoryOutput struct{ *pulumi.OutputState }

func (DataFactoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFactory)(nil)).Elem()
}

func (o DataFactoryOutput) ToDataFactoryOutput() DataFactoryOutput {
	return o
}

func (o DataFactoryOutput) ToDataFactoryOutputWithContext(ctx context.Context) DataFactoryOutput {
	return o
}

func (o DataFactoryOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactory) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DataFactoryOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactory) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DataFactoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactory) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataFactoryOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactory) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
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





type DataFactoryResponseInput interface {
	pulumi.Input

	ToDataFactoryResponseOutput() DataFactoryResponseOutput
	ToDataFactoryResponseOutputWithContext(context.Context) DataFactoryResponseOutput
}

type DataFactoryResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (DataFactoryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFactoryResponse)(nil)).Elem()
}

func (i DataFactoryResponseArgs) ToDataFactoryResponseOutput() DataFactoryResponseOutput {
	return i.ToDataFactoryResponseOutputWithContext(context.Background())
}

func (i DataFactoryResponseArgs) ToDataFactoryResponseOutputWithContext(ctx context.Context) DataFactoryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFactoryResponseOutput)
}

type DataFactoryResponseOutput struct{ *pulumi.OutputState }

func (DataFactoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFactoryResponse)(nil)).Elem()
}

func (o DataFactoryResponseOutput) ToDataFactoryResponseOutput() DataFactoryResponseOutput {
	return o
}

func (o DataFactoryResponseOutput) ToDataFactoryResponseOutputWithContext(ctx context.Context) DataFactoryResponseOutput {
	return o
}

func (o DataFactoryResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactoryResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DataFactoryResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactoryResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DataFactoryResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactoryResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o DataFactoryResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactoryResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataFactoryResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v DataFactoryResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o DataFactoryResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactoryResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o DataFactoryResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v DataFactoryResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o DataFactoryResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactoryResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DataFactoryResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactoryResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DataLakeAnalytics struct {
	ComputeLocation *string                      `pulumi:"computeLocation"`
	ComputeType     string                       `pulumi:"computeType"`
	Description     *string                      `pulumi:"description"`
	Properties      *DataLakeAnalyticsProperties `pulumi:"properties"`
	ResourceId      *string                      `pulumi:"resourceId"`
}





type DataLakeAnalyticsInput interface {
	pulumi.Input

	ToDataLakeAnalyticsOutput() DataLakeAnalyticsOutput
	ToDataLakeAnalyticsOutputWithContext(context.Context) DataLakeAnalyticsOutput
}

type DataLakeAnalyticsArgs struct {
	ComputeLocation pulumi.StringPtrInput               `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput                  `pulumi:"computeType"`
	Description     pulumi.StringPtrInput               `pulumi:"description"`
	Properties      DataLakeAnalyticsPropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput               `pulumi:"resourceId"`
}

func (DataLakeAnalyticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalytics)(nil)).Elem()
}

func (i DataLakeAnalyticsArgs) ToDataLakeAnalyticsOutput() DataLakeAnalyticsOutput {
	return i.ToDataLakeAnalyticsOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsArgs) ToDataLakeAnalyticsOutputWithContext(ctx context.Context) DataLakeAnalyticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsOutput)
}

type DataLakeAnalyticsOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalytics)(nil)).Elem()
}

func (o DataLakeAnalyticsOutput) ToDataLakeAnalyticsOutput() DataLakeAnalyticsOutput {
	return o
}

func (o DataLakeAnalyticsOutput) ToDataLakeAnalyticsOutputWithContext(ctx context.Context) DataLakeAnalyticsOutput {
	return o
}

func (o DataLakeAnalyticsOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalytics) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DataLakeAnalyticsOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeAnalytics) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DataLakeAnalyticsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalytics) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataLakeAnalyticsOutput) Properties() DataLakeAnalyticsPropertiesPtrOutput {
	return o.ApplyT(func(v DataLakeAnalytics) *DataLakeAnalyticsProperties { return v.Properties }).(DataLakeAnalyticsPropertiesPtrOutput)
}

func (o DataLakeAnalyticsOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalytics) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DataLakeAnalyticsProperties struct {
	DataLakeStoreAccountName *string `pulumi:"dataLakeStoreAccountName"`
}





type DataLakeAnalyticsPropertiesInput interface {
	pulumi.Input

	ToDataLakeAnalyticsPropertiesOutput() DataLakeAnalyticsPropertiesOutput
	ToDataLakeAnalyticsPropertiesOutputWithContext(context.Context) DataLakeAnalyticsPropertiesOutput
}

type DataLakeAnalyticsPropertiesArgs struct {
	DataLakeStoreAccountName pulumi.StringPtrInput `pulumi:"dataLakeStoreAccountName"`
}

func (DataLakeAnalyticsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsProperties)(nil)).Elem()
}

func (i DataLakeAnalyticsPropertiesArgs) ToDataLakeAnalyticsPropertiesOutput() DataLakeAnalyticsPropertiesOutput {
	return i.ToDataLakeAnalyticsPropertiesOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsPropertiesArgs) ToDataLakeAnalyticsPropertiesOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsPropertiesOutput)
}

func (i DataLakeAnalyticsPropertiesArgs) ToDataLakeAnalyticsPropertiesPtrOutput() DataLakeAnalyticsPropertiesPtrOutput {
	return i.ToDataLakeAnalyticsPropertiesPtrOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsPropertiesArgs) ToDataLakeAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsPropertiesOutput).ToDataLakeAnalyticsPropertiesPtrOutputWithContext(ctx)
}









type DataLakeAnalyticsPropertiesPtrInput interface {
	pulumi.Input

	ToDataLakeAnalyticsPropertiesPtrOutput() DataLakeAnalyticsPropertiesPtrOutput
	ToDataLakeAnalyticsPropertiesPtrOutputWithContext(context.Context) DataLakeAnalyticsPropertiesPtrOutput
}

type dataLakeAnalyticsPropertiesPtrType DataLakeAnalyticsPropertiesArgs

func DataLakeAnalyticsPropertiesPtr(v *DataLakeAnalyticsPropertiesArgs) DataLakeAnalyticsPropertiesPtrInput {
	return (*dataLakeAnalyticsPropertiesPtrType)(v)
}

func (*dataLakeAnalyticsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeAnalyticsProperties)(nil)).Elem()
}

func (i *dataLakeAnalyticsPropertiesPtrType) ToDataLakeAnalyticsPropertiesPtrOutput() DataLakeAnalyticsPropertiesPtrOutput {
	return i.ToDataLakeAnalyticsPropertiesPtrOutputWithContext(context.Background())
}

func (i *dataLakeAnalyticsPropertiesPtrType) ToDataLakeAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsPropertiesPtrOutput)
}

type DataLakeAnalyticsPropertiesOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsProperties)(nil)).Elem()
}

func (o DataLakeAnalyticsPropertiesOutput) ToDataLakeAnalyticsPropertiesOutput() DataLakeAnalyticsPropertiesOutput {
	return o
}

func (o DataLakeAnalyticsPropertiesOutput) ToDataLakeAnalyticsPropertiesOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesOutput {
	return o
}

func (o DataLakeAnalyticsPropertiesOutput) ToDataLakeAnalyticsPropertiesPtrOutput() DataLakeAnalyticsPropertiesPtrOutput {
	return o.ToDataLakeAnalyticsPropertiesPtrOutputWithContext(context.Background())
}

func (o DataLakeAnalyticsPropertiesOutput) ToDataLakeAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataLakeAnalyticsProperties) *DataLakeAnalyticsProperties {
		return &v
	}).(DataLakeAnalyticsPropertiesPtrOutput)
}

func (o DataLakeAnalyticsPropertiesOutput) DataLakeStoreAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsProperties) *string { return v.DataLakeStoreAccountName }).(pulumi.StringPtrOutput)
}

type DataLakeAnalyticsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeAnalyticsProperties)(nil)).Elem()
}

func (o DataLakeAnalyticsPropertiesPtrOutput) ToDataLakeAnalyticsPropertiesPtrOutput() DataLakeAnalyticsPropertiesPtrOutput {
	return o
}

func (o DataLakeAnalyticsPropertiesPtrOutput) ToDataLakeAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesPtrOutput {
	return o
}

func (o DataLakeAnalyticsPropertiesPtrOutput) Elem() DataLakeAnalyticsPropertiesOutput {
	return o.ApplyT(func(v *DataLakeAnalyticsProperties) DataLakeAnalyticsProperties {
		if v != nil {
			return *v
		}
		var ret DataLakeAnalyticsProperties
		return ret
	}).(DataLakeAnalyticsPropertiesOutput)
}

func (o DataLakeAnalyticsPropertiesPtrOutput) DataLakeStoreAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLakeAnalyticsProperties) *string {
		if v == nil {
			return nil
		}
		return v.DataLakeStoreAccountName
	}).(pulumi.StringPtrOutput)
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





type DataLakeAnalyticsResponseInput interface {
	pulumi.Input

	ToDataLakeAnalyticsResponseOutput() DataLakeAnalyticsResponseOutput
	ToDataLakeAnalyticsResponseOutputWithContext(context.Context) DataLakeAnalyticsResponseOutput
}

type DataLakeAnalyticsResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         DataLakeAnalyticsResponsePropertiesPtrInput   `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (DataLakeAnalyticsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsResponse)(nil)).Elem()
}

func (i DataLakeAnalyticsResponseArgs) ToDataLakeAnalyticsResponseOutput() DataLakeAnalyticsResponseOutput {
	return i.ToDataLakeAnalyticsResponseOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsResponseArgs) ToDataLakeAnalyticsResponseOutputWithContext(ctx context.Context) DataLakeAnalyticsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsResponseOutput)
}

type DataLakeAnalyticsResponseOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsResponse)(nil)).Elem()
}

func (o DataLakeAnalyticsResponseOutput) ToDataLakeAnalyticsResponseOutput() DataLakeAnalyticsResponseOutput {
	return o
}

func (o DataLakeAnalyticsResponseOutput) ToDataLakeAnalyticsResponseOutputWithContext(ctx context.Context) DataLakeAnalyticsResponseOutput {
	return o
}

func (o DataLakeAnalyticsResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DataLakeAnalyticsResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DataLakeAnalyticsResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o DataLakeAnalyticsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataLakeAnalyticsResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o DataLakeAnalyticsResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o DataLakeAnalyticsResponseOutput) Properties() DataLakeAnalyticsResponsePropertiesPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) *DataLakeAnalyticsResponseProperties { return v.Properties }).(DataLakeAnalyticsResponsePropertiesPtrOutput)
}

func (o DataLakeAnalyticsResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o DataLakeAnalyticsResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DataLakeAnalyticsResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DataLakeAnalyticsResponseProperties struct {
	DataLakeStoreAccountName *string `pulumi:"dataLakeStoreAccountName"`
}





type DataLakeAnalyticsResponsePropertiesInput interface {
	pulumi.Input

	ToDataLakeAnalyticsResponsePropertiesOutput() DataLakeAnalyticsResponsePropertiesOutput
	ToDataLakeAnalyticsResponsePropertiesOutputWithContext(context.Context) DataLakeAnalyticsResponsePropertiesOutput
}

type DataLakeAnalyticsResponsePropertiesArgs struct {
	DataLakeStoreAccountName pulumi.StringPtrInput `pulumi:"dataLakeStoreAccountName"`
}

func (DataLakeAnalyticsResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsResponseProperties)(nil)).Elem()
}

func (i DataLakeAnalyticsResponsePropertiesArgs) ToDataLakeAnalyticsResponsePropertiesOutput() DataLakeAnalyticsResponsePropertiesOutput {
	return i.ToDataLakeAnalyticsResponsePropertiesOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsResponsePropertiesArgs) ToDataLakeAnalyticsResponsePropertiesOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsResponsePropertiesOutput)
}

func (i DataLakeAnalyticsResponsePropertiesArgs) ToDataLakeAnalyticsResponsePropertiesPtrOutput() DataLakeAnalyticsResponsePropertiesPtrOutput {
	return i.ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsResponsePropertiesArgs) ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsResponsePropertiesOutput).ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(ctx)
}









type DataLakeAnalyticsResponsePropertiesPtrInput interface {
	pulumi.Input

	ToDataLakeAnalyticsResponsePropertiesPtrOutput() DataLakeAnalyticsResponsePropertiesPtrOutput
	ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(context.Context) DataLakeAnalyticsResponsePropertiesPtrOutput
}

type dataLakeAnalyticsResponsePropertiesPtrType DataLakeAnalyticsResponsePropertiesArgs

func DataLakeAnalyticsResponsePropertiesPtr(v *DataLakeAnalyticsResponsePropertiesArgs) DataLakeAnalyticsResponsePropertiesPtrInput {
	return (*dataLakeAnalyticsResponsePropertiesPtrType)(v)
}

func (*dataLakeAnalyticsResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeAnalyticsResponseProperties)(nil)).Elem()
}

func (i *dataLakeAnalyticsResponsePropertiesPtrType) ToDataLakeAnalyticsResponsePropertiesPtrOutput() DataLakeAnalyticsResponsePropertiesPtrOutput {
	return i.ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *dataLakeAnalyticsResponsePropertiesPtrType) ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsResponsePropertiesPtrOutput)
}

type DataLakeAnalyticsResponsePropertiesOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsResponseProperties)(nil)).Elem()
}

func (o DataLakeAnalyticsResponsePropertiesOutput) ToDataLakeAnalyticsResponsePropertiesOutput() DataLakeAnalyticsResponsePropertiesOutput {
	return o
}

func (o DataLakeAnalyticsResponsePropertiesOutput) ToDataLakeAnalyticsResponsePropertiesOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesOutput {
	return o
}

func (o DataLakeAnalyticsResponsePropertiesOutput) ToDataLakeAnalyticsResponsePropertiesPtrOutput() DataLakeAnalyticsResponsePropertiesPtrOutput {
	return o.ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o DataLakeAnalyticsResponsePropertiesOutput) ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataLakeAnalyticsResponseProperties) *DataLakeAnalyticsResponseProperties {
		return &v
	}).(DataLakeAnalyticsResponsePropertiesPtrOutput)
}

func (o DataLakeAnalyticsResponsePropertiesOutput) DataLakeStoreAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponseProperties) *string { return v.DataLakeStoreAccountName }).(pulumi.StringPtrOutput)
}

type DataLakeAnalyticsResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeAnalyticsResponseProperties)(nil)).Elem()
}

func (o DataLakeAnalyticsResponsePropertiesPtrOutput) ToDataLakeAnalyticsResponsePropertiesPtrOutput() DataLakeAnalyticsResponsePropertiesPtrOutput {
	return o
}

func (o DataLakeAnalyticsResponsePropertiesPtrOutput) ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesPtrOutput {
	return o
}

func (o DataLakeAnalyticsResponsePropertiesPtrOutput) Elem() DataLakeAnalyticsResponsePropertiesOutput {
	return o.ApplyT(func(v *DataLakeAnalyticsResponseProperties) DataLakeAnalyticsResponseProperties {
		if v != nil {
			return *v
		}
		var ret DataLakeAnalyticsResponseProperties
		return ret
	}).(DataLakeAnalyticsResponsePropertiesOutput)
}

func (o DataLakeAnalyticsResponsePropertiesPtrOutput) DataLakeStoreAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLakeAnalyticsResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.DataLakeStoreAccountName
	}).(pulumi.StringPtrOutput)
}

type Databricks struct {
	ComputeLocation *string               `pulumi:"computeLocation"`
	ComputeType     string                `pulumi:"computeType"`
	Description     *string               `pulumi:"description"`
	Properties      *DatabricksProperties `pulumi:"properties"`
	ResourceId      *string               `pulumi:"resourceId"`
}





type DatabricksInput interface {
	pulumi.Input

	ToDatabricksOutput() DatabricksOutput
	ToDatabricksOutputWithContext(context.Context) DatabricksOutput
}

type DatabricksArgs struct {
	ComputeLocation pulumi.StringPtrInput        `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput           `pulumi:"computeType"`
	Description     pulumi.StringPtrInput        `pulumi:"description"`
	Properties      DatabricksPropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput        `pulumi:"resourceId"`
}

func (DatabricksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Databricks)(nil)).Elem()
}

func (i DatabricksArgs) ToDatabricksOutput() DatabricksOutput {
	return i.ToDatabricksOutputWithContext(context.Background())
}

func (i DatabricksArgs) ToDatabricksOutputWithContext(ctx context.Context) DatabricksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksOutput)
}

type DatabricksOutput struct{ *pulumi.OutputState }

func (DatabricksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Databricks)(nil)).Elem()
}

func (o DatabricksOutput) ToDatabricksOutput() DatabricksOutput {
	return o
}

func (o DatabricksOutput) ToDatabricksOutputWithContext(ctx context.Context) DatabricksOutput {
	return o
}

func (o DatabricksOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Databricks) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DatabricksOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v Databricks) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DatabricksOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Databricks) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DatabricksOutput) Properties() DatabricksPropertiesPtrOutput {
	return o.ApplyT(func(v Databricks) *DatabricksProperties { return v.Properties }).(DatabricksPropertiesPtrOutput)
}

func (o DatabricksOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Databricks) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DatabricksProperties struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
}





type DatabricksPropertiesInput interface {
	pulumi.Input

	ToDatabricksPropertiesOutput() DatabricksPropertiesOutput
	ToDatabricksPropertiesOutputWithContext(context.Context) DatabricksPropertiesOutput
}

type DatabricksPropertiesArgs struct {
	DatabricksAccessToken pulumi.StringPtrInput `pulumi:"databricksAccessToken"`
}

func (DatabricksPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksProperties)(nil)).Elem()
}

func (i DatabricksPropertiesArgs) ToDatabricksPropertiesOutput() DatabricksPropertiesOutput {
	return i.ToDatabricksPropertiesOutputWithContext(context.Background())
}

func (i DatabricksPropertiesArgs) ToDatabricksPropertiesOutputWithContext(ctx context.Context) DatabricksPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksPropertiesOutput)
}

func (i DatabricksPropertiesArgs) ToDatabricksPropertiesPtrOutput() DatabricksPropertiesPtrOutput {
	return i.ToDatabricksPropertiesPtrOutputWithContext(context.Background())
}

func (i DatabricksPropertiesArgs) ToDatabricksPropertiesPtrOutputWithContext(ctx context.Context) DatabricksPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksPropertiesOutput).ToDatabricksPropertiesPtrOutputWithContext(ctx)
}









type DatabricksPropertiesPtrInput interface {
	pulumi.Input

	ToDatabricksPropertiesPtrOutput() DatabricksPropertiesPtrOutput
	ToDatabricksPropertiesPtrOutputWithContext(context.Context) DatabricksPropertiesPtrOutput
}

type databricksPropertiesPtrType DatabricksPropertiesArgs

func DatabricksPropertiesPtr(v *DatabricksPropertiesArgs) DatabricksPropertiesPtrInput {
	return (*databricksPropertiesPtrType)(v)
}

func (*databricksPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksProperties)(nil)).Elem()
}

func (i *databricksPropertiesPtrType) ToDatabricksPropertiesPtrOutput() DatabricksPropertiesPtrOutput {
	return i.ToDatabricksPropertiesPtrOutputWithContext(context.Background())
}

func (i *databricksPropertiesPtrType) ToDatabricksPropertiesPtrOutputWithContext(ctx context.Context) DatabricksPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksPropertiesPtrOutput)
}

type DatabricksPropertiesOutput struct{ *pulumi.OutputState }

func (DatabricksPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksProperties)(nil)).Elem()
}

func (o DatabricksPropertiesOutput) ToDatabricksPropertiesOutput() DatabricksPropertiesOutput {
	return o
}

func (o DatabricksPropertiesOutput) ToDatabricksPropertiesOutputWithContext(ctx context.Context) DatabricksPropertiesOutput {
	return o
}

func (o DatabricksPropertiesOutput) ToDatabricksPropertiesPtrOutput() DatabricksPropertiesPtrOutput {
	return o.ToDatabricksPropertiesPtrOutputWithContext(context.Background())
}

func (o DatabricksPropertiesOutput) ToDatabricksPropertiesPtrOutputWithContext(ctx context.Context) DatabricksPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabricksProperties) *DatabricksProperties {
		return &v
	}).(DatabricksPropertiesPtrOutput)
}

func (o DatabricksPropertiesOutput) DatabricksAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksProperties) *string { return v.DatabricksAccessToken }).(pulumi.StringPtrOutput)
}

type DatabricksPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DatabricksPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksProperties)(nil)).Elem()
}

func (o DatabricksPropertiesPtrOutput) ToDatabricksPropertiesPtrOutput() DatabricksPropertiesPtrOutput {
	return o
}

func (o DatabricksPropertiesPtrOutput) ToDatabricksPropertiesPtrOutputWithContext(ctx context.Context) DatabricksPropertiesPtrOutput {
	return o
}

func (o DatabricksPropertiesPtrOutput) Elem() DatabricksPropertiesOutput {
	return o.ApplyT(func(v *DatabricksProperties) DatabricksProperties {
		if v != nil {
			return *v
		}
		var ret DatabricksProperties
		return ret
	}).(DatabricksPropertiesOutput)
}

func (o DatabricksPropertiesPtrOutput) DatabricksAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabricksProperties) *string {
		if v == nil {
			return nil
		}
		return v.DatabricksAccessToken
	}).(pulumi.StringPtrOutput)
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





type DatabricksResponseInput interface {
	pulumi.Input

	ToDatabricksResponseOutput() DatabricksResponseOutput
	ToDatabricksResponseOutputWithContext(context.Context) DatabricksResponseOutput
}

type DatabricksResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         DatabricksResponsePropertiesPtrInput          `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (DatabricksResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksResponse)(nil)).Elem()
}

func (i DatabricksResponseArgs) ToDatabricksResponseOutput() DatabricksResponseOutput {
	return i.ToDatabricksResponseOutputWithContext(context.Background())
}

func (i DatabricksResponseArgs) ToDatabricksResponseOutputWithContext(ctx context.Context) DatabricksResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksResponseOutput)
}

type DatabricksResponseOutput struct{ *pulumi.OutputState }

func (DatabricksResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksResponse)(nil)).Elem()
}

func (o DatabricksResponseOutput) ToDatabricksResponseOutput() DatabricksResponseOutput {
	return o
}

func (o DatabricksResponseOutput) ToDatabricksResponseOutputWithContext(ctx context.Context) DatabricksResponseOutput {
	return o
}

func (o DatabricksResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DatabricksResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v DatabricksResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DatabricksResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DatabricksResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o DatabricksResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DatabricksResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v DatabricksResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o DatabricksResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DatabricksResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o DatabricksResponseOutput) Properties() DatabricksResponsePropertiesPtrOutput {
	return o.ApplyT(func(v DatabricksResponse) *DatabricksResponseProperties { return v.Properties }).(DatabricksResponsePropertiesPtrOutput)
}

func (o DatabricksResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v DatabricksResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o DatabricksResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DatabricksResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DatabricksResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DatabricksResponseProperties struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
}





type DatabricksResponsePropertiesInput interface {
	pulumi.Input

	ToDatabricksResponsePropertiesOutput() DatabricksResponsePropertiesOutput
	ToDatabricksResponsePropertiesOutputWithContext(context.Context) DatabricksResponsePropertiesOutput
}

type DatabricksResponsePropertiesArgs struct {
	DatabricksAccessToken pulumi.StringPtrInput `pulumi:"databricksAccessToken"`
}

func (DatabricksResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksResponseProperties)(nil)).Elem()
}

func (i DatabricksResponsePropertiesArgs) ToDatabricksResponsePropertiesOutput() DatabricksResponsePropertiesOutput {
	return i.ToDatabricksResponsePropertiesOutputWithContext(context.Background())
}

func (i DatabricksResponsePropertiesArgs) ToDatabricksResponsePropertiesOutputWithContext(ctx context.Context) DatabricksResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksResponsePropertiesOutput)
}

func (i DatabricksResponsePropertiesArgs) ToDatabricksResponsePropertiesPtrOutput() DatabricksResponsePropertiesPtrOutput {
	return i.ToDatabricksResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i DatabricksResponsePropertiesArgs) ToDatabricksResponsePropertiesPtrOutputWithContext(ctx context.Context) DatabricksResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksResponsePropertiesOutput).ToDatabricksResponsePropertiesPtrOutputWithContext(ctx)
}









type DatabricksResponsePropertiesPtrInput interface {
	pulumi.Input

	ToDatabricksResponsePropertiesPtrOutput() DatabricksResponsePropertiesPtrOutput
	ToDatabricksResponsePropertiesPtrOutputWithContext(context.Context) DatabricksResponsePropertiesPtrOutput
}

type databricksResponsePropertiesPtrType DatabricksResponsePropertiesArgs

func DatabricksResponsePropertiesPtr(v *DatabricksResponsePropertiesArgs) DatabricksResponsePropertiesPtrInput {
	return (*databricksResponsePropertiesPtrType)(v)
}

func (*databricksResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksResponseProperties)(nil)).Elem()
}

func (i *databricksResponsePropertiesPtrType) ToDatabricksResponsePropertiesPtrOutput() DatabricksResponsePropertiesPtrOutput {
	return i.ToDatabricksResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *databricksResponsePropertiesPtrType) ToDatabricksResponsePropertiesPtrOutputWithContext(ctx context.Context) DatabricksResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksResponsePropertiesPtrOutput)
}

type DatabricksResponsePropertiesOutput struct{ *pulumi.OutputState }

func (DatabricksResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksResponseProperties)(nil)).Elem()
}

func (o DatabricksResponsePropertiesOutput) ToDatabricksResponsePropertiesOutput() DatabricksResponsePropertiesOutput {
	return o
}

func (o DatabricksResponsePropertiesOutput) ToDatabricksResponsePropertiesOutputWithContext(ctx context.Context) DatabricksResponsePropertiesOutput {
	return o
}

func (o DatabricksResponsePropertiesOutput) ToDatabricksResponsePropertiesPtrOutput() DatabricksResponsePropertiesPtrOutput {
	return o.ToDatabricksResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o DatabricksResponsePropertiesOutput) ToDatabricksResponsePropertiesPtrOutputWithContext(ctx context.Context) DatabricksResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabricksResponseProperties) *DatabricksResponseProperties {
		return &v
	}).(DatabricksResponsePropertiesPtrOutput)
}

func (o DatabricksResponsePropertiesOutput) DatabricksAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksResponseProperties) *string { return v.DatabricksAccessToken }).(pulumi.StringPtrOutput)
}

type DatabricksResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (DatabricksResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksResponseProperties)(nil)).Elem()
}

func (o DatabricksResponsePropertiesPtrOutput) ToDatabricksResponsePropertiesPtrOutput() DatabricksResponsePropertiesPtrOutput {
	return o
}

func (o DatabricksResponsePropertiesPtrOutput) ToDatabricksResponsePropertiesPtrOutputWithContext(ctx context.Context) DatabricksResponsePropertiesPtrOutput {
	return o
}

func (o DatabricksResponsePropertiesPtrOutput) Elem() DatabricksResponsePropertiesOutput {
	return o.ApplyT(func(v *DatabricksResponseProperties) DatabricksResponseProperties {
		if v != nil {
			return *v
		}
		var ret DatabricksResponseProperties
		return ret
	}).(DatabricksResponsePropertiesOutput)
}

func (o DatabricksResponsePropertiesPtrOutput) DatabricksAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabricksResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.DatabricksAccessToken
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





type DatasetReferenceResponseInput interface {
	pulumi.Input

	ToDatasetReferenceResponseOutput() DatasetReferenceResponseOutput
	ToDatasetReferenceResponseOutputWithContext(context.Context) DatasetReferenceResponseOutput
}

type DatasetReferenceResponseArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DatasetReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetReferenceResponse)(nil)).Elem()
}

func (i DatasetReferenceResponseArgs) ToDatasetReferenceResponseOutput() DatasetReferenceResponseOutput {
	return i.ToDatasetReferenceResponseOutputWithContext(context.Background())
}

func (i DatasetReferenceResponseArgs) ToDatasetReferenceResponseOutputWithContext(ctx context.Context) DatasetReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetReferenceResponseOutput)
}





type DatasetReferenceResponseArrayInput interface {
	pulumi.Input

	ToDatasetReferenceResponseArrayOutput() DatasetReferenceResponseArrayOutput
	ToDatasetReferenceResponseArrayOutputWithContext(context.Context) DatasetReferenceResponseArrayOutput
}

type DatasetReferenceResponseArray []DatasetReferenceResponseInput

func (DatasetReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatasetReferenceResponse)(nil)).Elem()
}

func (i DatasetReferenceResponseArray) ToDatasetReferenceResponseArrayOutput() DatasetReferenceResponseArrayOutput {
	return i.ToDatasetReferenceResponseArrayOutputWithContext(context.Background())
}

func (i DatasetReferenceResponseArray) ToDatasetReferenceResponseArrayOutputWithContext(ctx context.Context) DatasetReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetReferenceResponseArrayOutput)
}

type DatasetReferenceResponseOutput struct{ *pulumi.OutputState }

func (DatasetReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetReferenceResponse)(nil)).Elem()
}

func (o DatasetReferenceResponseOutput) ToDatasetReferenceResponseOutput() DatasetReferenceResponseOutput {
	return o
}

func (o DatasetReferenceResponseOutput) ToDatasetReferenceResponseOutputWithContext(ctx context.Context) DatasetReferenceResponseOutput {
	return o
}

func (o DatasetReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o DatasetReferenceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatasetReferenceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DatasetReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (DatasetReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatasetReferenceResponse)(nil)).Elem()
}

func (o DatasetReferenceResponseArrayOutput) ToDatasetReferenceResponseArrayOutput() DatasetReferenceResponseArrayOutput {
	return o
}

func (o DatasetReferenceResponseArrayOutput) ToDatasetReferenceResponseArrayOutputWithContext(ctx context.Context) DatasetReferenceResponseArrayOutput {
	return o
}

func (o DatasetReferenceResponseArrayOutput) Index(i pulumi.IntInput) DatasetReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatasetReferenceResponse {
		return vs[0].([]DatasetReferenceResponse)[vs[1].(int)]
	}).(DatasetReferenceResponseOutput)
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





type EncryptionPropertyResponseInput interface {
	pulumi.Input

	ToEncryptionPropertyResponseOutput() EncryptionPropertyResponseOutput
	ToEncryptionPropertyResponseOutputWithContext(context.Context) EncryptionPropertyResponseOutput
}

type EncryptionPropertyResponseArgs struct {
	KeyVaultProperties KeyVaultPropertiesResponseInput `pulumi:"keyVaultProperties"`
	Status             pulumi.StringInput              `pulumi:"status"`
}

func (EncryptionPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertyResponse)(nil)).Elem()
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponseOutput() EncryptionPropertyResponseOutput {
	return i.ToEncryptionPropertyResponseOutputWithContext(context.Background())
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponseOutputWithContext(ctx context.Context) EncryptionPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyResponseOutput)
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return i.ToEncryptionPropertyResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyResponseOutput).ToEncryptionPropertyResponsePtrOutputWithContext(ctx)
}









type EncryptionPropertyResponsePtrInput interface {
	pulumi.Input

	ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput
	ToEncryptionPropertyResponsePtrOutputWithContext(context.Context) EncryptionPropertyResponsePtrOutput
}

type encryptionPropertyResponsePtrType EncryptionPropertyResponseArgs

func EncryptionPropertyResponsePtr(v *EncryptionPropertyResponseArgs) EncryptionPropertyResponsePtrInput {
	return (*encryptionPropertyResponsePtrType)(v)
}

func (*encryptionPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertyResponse)(nil)).Elem()
}

func (i *encryptionPropertyResponsePtrType) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return i.ToEncryptionPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionPropertyResponsePtrType) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyResponsePtrOutput)
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

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return o.ToEncryptionPropertyResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPropertyResponse) *EncryptionPropertyResponse {
		return &v
	}).(EncryptionPropertyResponsePtrOutput)
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





type EnvironmentImageResponseResponseEnvironmentInput interface {
	pulumi.Input

	ToEnvironmentImageResponseResponseEnvironmentOutput() EnvironmentImageResponseResponseEnvironmentOutput
	ToEnvironmentImageResponseResponseEnvironmentOutputWithContext(context.Context) EnvironmentImageResponseResponseEnvironmentOutput
}

type EnvironmentImageResponseResponseEnvironmentArgs struct {
	Docker                  ModelEnvironmentDefinitionResponseResponseDockerPtrInput `pulumi:"docker"`
	EnvironmentVariables    pulumi.StringMapInput                                    `pulumi:"environmentVariables"`
	InferencingStackVersion pulumi.StringPtrInput                                    `pulumi:"inferencingStackVersion"`
	Name                    pulumi.StringPtrInput                                    `pulumi:"name"`
	Python                  ModelEnvironmentDefinitionResponseResponsePythonPtrInput `pulumi:"python"`
	R                       ModelEnvironmentDefinitionResponseResponseRPtrInput      `pulumi:"r"`
	Spark                   ModelEnvironmentDefinitionResponseResponseSparkPtrInput  `pulumi:"spark"`
	Version                 pulumi.StringPtrInput                                    `pulumi:"version"`
}

func (EnvironmentImageResponseResponseEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentImageResponseResponseEnvironment)(nil)).Elem()
}

func (i EnvironmentImageResponseResponseEnvironmentArgs) ToEnvironmentImageResponseResponseEnvironmentOutput() EnvironmentImageResponseResponseEnvironmentOutput {
	return i.ToEnvironmentImageResponseResponseEnvironmentOutputWithContext(context.Background())
}

func (i EnvironmentImageResponseResponseEnvironmentArgs) ToEnvironmentImageResponseResponseEnvironmentOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageResponseResponseEnvironmentOutput)
}

func (i EnvironmentImageResponseResponseEnvironmentArgs) ToEnvironmentImageResponseResponseEnvironmentPtrOutput() EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return i.ToEnvironmentImageResponseResponseEnvironmentPtrOutputWithContext(context.Background())
}

func (i EnvironmentImageResponseResponseEnvironmentArgs) ToEnvironmentImageResponseResponseEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageResponseResponseEnvironmentOutput).ToEnvironmentImageResponseResponseEnvironmentPtrOutputWithContext(ctx)
}









type EnvironmentImageResponseResponseEnvironmentPtrInput interface {
	pulumi.Input

	ToEnvironmentImageResponseResponseEnvironmentPtrOutput() EnvironmentImageResponseResponseEnvironmentPtrOutput
	ToEnvironmentImageResponseResponseEnvironmentPtrOutputWithContext(context.Context) EnvironmentImageResponseResponseEnvironmentPtrOutput
}

type environmentImageResponseResponseEnvironmentPtrType EnvironmentImageResponseResponseEnvironmentArgs

func EnvironmentImageResponseResponseEnvironmentPtr(v *EnvironmentImageResponseResponseEnvironmentArgs) EnvironmentImageResponseResponseEnvironmentPtrInput {
	return (*environmentImageResponseResponseEnvironmentPtrType)(v)
}

func (*environmentImageResponseResponseEnvironmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentImageResponseResponseEnvironment)(nil)).Elem()
}

func (i *environmentImageResponseResponseEnvironmentPtrType) ToEnvironmentImageResponseResponseEnvironmentPtrOutput() EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return i.ToEnvironmentImageResponseResponseEnvironmentPtrOutputWithContext(context.Background())
}

func (i *environmentImageResponseResponseEnvironmentPtrType) ToEnvironmentImageResponseResponseEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageResponseResponseEnvironmentPtrOutput)
}

type EnvironmentImageResponseResponseEnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentImageResponseResponseEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentImageResponseResponseEnvironment)(nil)).Elem()
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) ToEnvironmentImageResponseResponseEnvironmentOutput() EnvironmentImageResponseResponseEnvironmentOutput {
	return o
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) ToEnvironmentImageResponseResponseEnvironmentOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentOutput {
	return o
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) ToEnvironmentImageResponseResponseEnvironmentPtrOutput() EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return o.ToEnvironmentImageResponseResponseEnvironmentPtrOutputWithContext(context.Background())
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) ToEnvironmentImageResponseResponseEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentImageResponseResponseEnvironment) *EnvironmentImageResponseResponseEnvironment {
		return &v
	}).(EnvironmentImageResponseResponseEnvironmentPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) Docker() ModelEnvironmentDefinitionResponseResponseDockerPtrOutput {
	return o.ApplyT(func(v EnvironmentImageResponseResponseEnvironment) *ModelEnvironmentDefinitionResponseResponseDocker {
		return v.Docker
	}).(ModelEnvironmentDefinitionResponseResponseDockerPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v EnvironmentImageResponseResponseEnvironment) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) InferencingStackVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentImageResponseResponseEnvironment) *string { return v.InferencingStackVersion }).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentImageResponseResponseEnvironment) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) Python() ModelEnvironmentDefinitionResponseResponsePythonPtrOutput {
	return o.ApplyT(func(v EnvironmentImageResponseResponseEnvironment) *ModelEnvironmentDefinitionResponseResponsePython {
		return v.Python
	}).(ModelEnvironmentDefinitionResponseResponsePythonPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) R() ModelEnvironmentDefinitionResponseResponseRPtrOutput {
	return o.ApplyT(func(v EnvironmentImageResponseResponseEnvironment) *ModelEnvironmentDefinitionResponseResponseR {
		return v.R
	}).(ModelEnvironmentDefinitionResponseResponseRPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) Spark() ModelEnvironmentDefinitionResponseResponseSparkPtrOutput {
	return o.ApplyT(func(v EnvironmentImageResponseResponseEnvironment) *ModelEnvironmentDefinitionResponseResponseSpark {
		return v.Spark
	}).(ModelEnvironmentDefinitionResponseResponseSparkPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentImageResponseResponseEnvironment) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type EnvironmentImageResponseResponseEnvironmentPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentImageResponseResponseEnvironmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentImageResponseResponseEnvironment)(nil)).Elem()
}

func (o EnvironmentImageResponseResponseEnvironmentPtrOutput) ToEnvironmentImageResponseResponseEnvironmentPtrOutput() EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return o
}

func (o EnvironmentImageResponseResponseEnvironmentPtrOutput) ToEnvironmentImageResponseResponseEnvironmentPtrOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentPtrOutput {
	return o
}

func (o EnvironmentImageResponseResponseEnvironmentPtrOutput) Elem() EnvironmentImageResponseResponseEnvironmentOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironment) EnvironmentImageResponseResponseEnvironment {
		if v != nil {
			return *v
		}
		var ret EnvironmentImageResponseResponseEnvironment
		return ret
	}).(EnvironmentImageResponseResponseEnvironmentOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentPtrOutput) Docker() ModelEnvironmentDefinitionResponseResponseDockerPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironment) *ModelEnvironmentDefinitionResponseResponseDocker {
		if v == nil {
			return nil
		}
		return v.Docker
	}).(ModelEnvironmentDefinitionResponseResponseDockerPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentPtrOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironment) map[string]string {
		if v == nil {
			return nil
		}
		return v.EnvironmentVariables
	}).(pulumi.StringMapOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentPtrOutput) InferencingStackVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironment) *string {
		if v == nil {
			return nil
		}
		return v.InferencingStackVersion
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironment) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentPtrOutput) Python() ModelEnvironmentDefinitionResponseResponsePythonPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironment) *ModelEnvironmentDefinitionResponseResponsePython {
		if v == nil {
			return nil
		}
		return v.Python
	}).(ModelEnvironmentDefinitionResponseResponsePythonPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentPtrOutput) R() ModelEnvironmentDefinitionResponseResponseRPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironment) *ModelEnvironmentDefinitionResponseResponseR {
		if v == nil {
			return nil
		}
		return v.R
	}).(ModelEnvironmentDefinitionResponseResponseRPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentPtrOutput) Spark() ModelEnvironmentDefinitionResponseResponseSparkPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironment) *ModelEnvironmentDefinitionResponseResponseSpark {
		if v == nil {
			return nil
		}
		return v.Spark
	}).(ModelEnvironmentDefinitionResponseResponseSparkPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironment) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type EnvironmentImageResponseResponseEnvironmentReference struct {
	Name    *string `pulumi:"name"`
	Version *string `pulumi:"version"`
}





type EnvironmentImageResponseResponseEnvironmentReferenceInput interface {
	pulumi.Input

	ToEnvironmentImageResponseResponseEnvironmentReferenceOutput() EnvironmentImageResponseResponseEnvironmentReferenceOutput
	ToEnvironmentImageResponseResponseEnvironmentReferenceOutputWithContext(context.Context) EnvironmentImageResponseResponseEnvironmentReferenceOutput
}

type EnvironmentImageResponseResponseEnvironmentReferenceArgs struct {
	Name    pulumi.StringPtrInput `pulumi:"name"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (EnvironmentImageResponseResponseEnvironmentReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentImageResponseResponseEnvironmentReference)(nil)).Elem()
}

func (i EnvironmentImageResponseResponseEnvironmentReferenceArgs) ToEnvironmentImageResponseResponseEnvironmentReferenceOutput() EnvironmentImageResponseResponseEnvironmentReferenceOutput {
	return i.ToEnvironmentImageResponseResponseEnvironmentReferenceOutputWithContext(context.Background())
}

func (i EnvironmentImageResponseResponseEnvironmentReferenceArgs) ToEnvironmentImageResponseResponseEnvironmentReferenceOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageResponseResponseEnvironmentReferenceOutput)
}

func (i EnvironmentImageResponseResponseEnvironmentReferenceArgs) ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutput() EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return i.ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutputWithContext(context.Background())
}

func (i EnvironmentImageResponseResponseEnvironmentReferenceArgs) ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageResponseResponseEnvironmentReferenceOutput).ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutputWithContext(ctx)
}









type EnvironmentImageResponseResponseEnvironmentReferencePtrInput interface {
	pulumi.Input

	ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutput() EnvironmentImageResponseResponseEnvironmentReferencePtrOutput
	ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutputWithContext(context.Context) EnvironmentImageResponseResponseEnvironmentReferencePtrOutput
}

type environmentImageResponseResponseEnvironmentReferencePtrType EnvironmentImageResponseResponseEnvironmentReferenceArgs

func EnvironmentImageResponseResponseEnvironmentReferencePtr(v *EnvironmentImageResponseResponseEnvironmentReferenceArgs) EnvironmentImageResponseResponseEnvironmentReferencePtrInput {
	return (*environmentImageResponseResponseEnvironmentReferencePtrType)(v)
}

func (*environmentImageResponseResponseEnvironmentReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentImageResponseResponseEnvironmentReference)(nil)).Elem()
}

func (i *environmentImageResponseResponseEnvironmentReferencePtrType) ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutput() EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return i.ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutputWithContext(context.Background())
}

func (i *environmentImageResponseResponseEnvironmentReferencePtrType) ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentImageResponseResponseEnvironmentReferencePtrOutput)
}

type EnvironmentImageResponseResponseEnvironmentReferenceOutput struct{ *pulumi.OutputState }

func (EnvironmentImageResponseResponseEnvironmentReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentImageResponseResponseEnvironmentReference)(nil)).Elem()
}

func (o EnvironmentImageResponseResponseEnvironmentReferenceOutput) ToEnvironmentImageResponseResponseEnvironmentReferenceOutput() EnvironmentImageResponseResponseEnvironmentReferenceOutput {
	return o
}

func (o EnvironmentImageResponseResponseEnvironmentReferenceOutput) ToEnvironmentImageResponseResponseEnvironmentReferenceOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentReferenceOutput {
	return o
}

func (o EnvironmentImageResponseResponseEnvironmentReferenceOutput) ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutput() EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return o.ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutputWithContext(context.Background())
}

func (o EnvironmentImageResponseResponseEnvironmentReferenceOutput) ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentImageResponseResponseEnvironmentReference) *EnvironmentImageResponseResponseEnvironmentReference {
		return &v
	}).(EnvironmentImageResponseResponseEnvironmentReferencePtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentImageResponseResponseEnvironmentReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentReferenceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentImageResponseResponseEnvironmentReference) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type EnvironmentImageResponseResponseEnvironmentReferencePtrOutput struct{ *pulumi.OutputState }

func (EnvironmentImageResponseResponseEnvironmentReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentImageResponseResponseEnvironmentReference)(nil)).Elem()
}

func (o EnvironmentImageResponseResponseEnvironmentReferencePtrOutput) ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutput() EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return o
}

func (o EnvironmentImageResponseResponseEnvironmentReferencePtrOutput) ToEnvironmentImageResponseResponseEnvironmentReferencePtrOutputWithContext(ctx context.Context) EnvironmentImageResponseResponseEnvironmentReferencePtrOutput {
	return o
}

func (o EnvironmentImageResponseResponseEnvironmentReferencePtrOutput) Elem() EnvironmentImageResponseResponseEnvironmentReferenceOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironmentReference) EnvironmentImageResponseResponseEnvironmentReference {
		if v != nil {
			return *v
		}
		var ret EnvironmentImageResponseResponseEnvironmentReference
		return ret
	}).(EnvironmentImageResponseResponseEnvironmentReferenceOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironmentReference) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentImageResponseResponseEnvironmentReferencePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentImageResponseResponseEnvironmentReference) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ErrorDetailResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}





type ErrorDetailResponseInput interface {
	pulumi.Input

	ToErrorDetailResponseOutput() ErrorDetailResponseOutput
	ToErrorDetailResponseOutputWithContext(context.Context) ErrorDetailResponseOutput
}

type ErrorDetailResponseArgs struct {
	Code    pulumi.StringInput `pulumi:"code"`
	Message pulumi.StringInput `pulumi:"message"`
}

func (ErrorDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return i.ToErrorDetailResponseOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput)
}





type ErrorDetailResponseArrayInput interface {
	pulumi.Input

	ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput
	ToErrorDetailResponseArrayOutputWithContext(context.Context) ErrorDetailResponseArrayOutput
}

type ErrorDetailResponseArray []ErrorDetailResponseInput

func (ErrorDetailResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return i.ToErrorDetailResponseArrayOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseArrayOutput)
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

type ErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) ErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetailResponse {
		return vs[0].([]ErrorDetailResponse)[vs[1].(int)]
	}).(ErrorDetailResponseOutput)
}

type ErrorResponseResponse struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
}





type ErrorResponseResponseInput interface {
	pulumi.Input

	ToErrorResponseResponseOutput() ErrorResponseResponseOutput
	ToErrorResponseResponseOutputWithContext(context.Context) ErrorResponseResponseOutput
}

type ErrorResponseResponseArgs struct {
	Code    pulumi.StringInput            `pulumi:"code"`
	Details ErrorDetailResponseArrayInput `pulumi:"details"`
	Message pulumi.StringInput            `pulumi:"message"`
}

func (ErrorResponseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponseResponse)(nil)).Elem()
}

func (i ErrorResponseResponseArgs) ToErrorResponseResponseOutput() ErrorResponseResponseOutput {
	return i.ToErrorResponseResponseOutputWithContext(context.Background())
}

func (i ErrorResponseResponseArgs) ToErrorResponseResponseOutputWithContext(ctx context.Context) ErrorResponseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseResponseOutput)
}

type ErrorResponseResponseOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutput() ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutputWithContext(ctx context.Context) ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorResponseResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorResponseResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorResponseResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Message }).(pulumi.StringOutput)
}

type HDInsight struct {
	ComputeLocation *string              `pulumi:"computeLocation"`
	ComputeType     string               `pulumi:"computeType"`
	Description     *string              `pulumi:"description"`
	Properties      *HDInsightProperties `pulumi:"properties"`
	ResourceId      *string              `pulumi:"resourceId"`
}





type HDInsightInput interface {
	pulumi.Input

	ToHDInsightOutput() HDInsightOutput
	ToHDInsightOutputWithContext(context.Context) HDInsightOutput
}

type HDInsightArgs struct {
	ComputeLocation pulumi.StringPtrInput       `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput          `pulumi:"computeType"`
	Description     pulumi.StringPtrInput       `pulumi:"description"`
	Properties      HDInsightPropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput       `pulumi:"resourceId"`
}

func (HDInsightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsight)(nil)).Elem()
}

func (i HDInsightArgs) ToHDInsightOutput() HDInsightOutput {
	return i.ToHDInsightOutputWithContext(context.Background())
}

func (i HDInsightArgs) ToHDInsightOutputWithContext(ctx context.Context) HDInsightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightOutput)
}

type HDInsightOutput struct{ *pulumi.OutputState }

func (HDInsightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsight)(nil)).Elem()
}

func (o HDInsightOutput) ToHDInsightOutput() HDInsightOutput {
	return o
}

func (o HDInsightOutput) ToHDInsightOutputWithContext(ctx context.Context) HDInsightOutput {
	return o
}

func (o HDInsightOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsight) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o HDInsightOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsight) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o HDInsightOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsight) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o HDInsightOutput) Properties() HDInsightPropertiesPtrOutput {
	return o.ApplyT(func(v HDInsight) *HDInsightProperties { return v.Properties }).(HDInsightPropertiesPtrOutput)
}

func (o HDInsightOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsight) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type HDInsightProperties struct {
	Address              *string                       `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	SshPort              *int                          `pulumi:"sshPort"`
}





type HDInsightPropertiesInput interface {
	pulumi.Input

	ToHDInsightPropertiesOutput() HDInsightPropertiesOutput
	ToHDInsightPropertiesOutputWithContext(context.Context) HDInsightPropertiesOutput
}

type HDInsightPropertiesArgs struct {
	Address              pulumi.StringPtrInput                `pulumi:"address"`
	AdministratorAccount VirtualMachineSshCredentialsPtrInput `pulumi:"administratorAccount"`
	SshPort              pulumi.IntPtrInput                   `pulumi:"sshPort"`
}

func (HDInsightPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightProperties)(nil)).Elem()
}

func (i HDInsightPropertiesArgs) ToHDInsightPropertiesOutput() HDInsightPropertiesOutput {
	return i.ToHDInsightPropertiesOutputWithContext(context.Background())
}

func (i HDInsightPropertiesArgs) ToHDInsightPropertiesOutputWithContext(ctx context.Context) HDInsightPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightPropertiesOutput)
}

func (i HDInsightPropertiesArgs) ToHDInsightPropertiesPtrOutput() HDInsightPropertiesPtrOutput {
	return i.ToHDInsightPropertiesPtrOutputWithContext(context.Background())
}

func (i HDInsightPropertiesArgs) ToHDInsightPropertiesPtrOutputWithContext(ctx context.Context) HDInsightPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightPropertiesOutput).ToHDInsightPropertiesPtrOutputWithContext(ctx)
}









type HDInsightPropertiesPtrInput interface {
	pulumi.Input

	ToHDInsightPropertiesPtrOutput() HDInsightPropertiesPtrOutput
	ToHDInsightPropertiesPtrOutputWithContext(context.Context) HDInsightPropertiesPtrOutput
}

type hdinsightPropertiesPtrType HDInsightPropertiesArgs

func HDInsightPropertiesPtr(v *HDInsightPropertiesArgs) HDInsightPropertiesPtrInput {
	return (*hdinsightPropertiesPtrType)(v)
}

func (*hdinsightPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HDInsightProperties)(nil)).Elem()
}

func (i *hdinsightPropertiesPtrType) ToHDInsightPropertiesPtrOutput() HDInsightPropertiesPtrOutput {
	return i.ToHDInsightPropertiesPtrOutputWithContext(context.Background())
}

func (i *hdinsightPropertiesPtrType) ToHDInsightPropertiesPtrOutputWithContext(ctx context.Context) HDInsightPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightPropertiesPtrOutput)
}

type HDInsightPropertiesOutput struct{ *pulumi.OutputState }

func (HDInsightPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightProperties)(nil)).Elem()
}

func (o HDInsightPropertiesOutput) ToHDInsightPropertiesOutput() HDInsightPropertiesOutput {
	return o
}

func (o HDInsightPropertiesOutput) ToHDInsightPropertiesOutputWithContext(ctx context.Context) HDInsightPropertiesOutput {
	return o
}

func (o HDInsightPropertiesOutput) ToHDInsightPropertiesPtrOutput() HDInsightPropertiesPtrOutput {
	return o.ToHDInsightPropertiesPtrOutputWithContext(context.Background())
}

func (o HDInsightPropertiesOutput) ToHDInsightPropertiesPtrOutputWithContext(ctx context.Context) HDInsightPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HDInsightProperties) *HDInsightProperties {
		return &v
	}).(HDInsightPropertiesPtrOutput)
}

func (o HDInsightPropertiesOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightProperties) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o HDInsightPropertiesOutput) AdministratorAccount() VirtualMachineSshCredentialsPtrOutput {
	return o.ApplyT(func(v HDInsightProperties) *VirtualMachineSshCredentials { return v.AdministratorAccount }).(VirtualMachineSshCredentialsPtrOutput)
}

func (o HDInsightPropertiesOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HDInsightProperties) *int { return v.SshPort }).(pulumi.IntPtrOutput)
}

type HDInsightPropertiesPtrOutput struct{ *pulumi.OutputState }

func (HDInsightPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HDInsightProperties)(nil)).Elem()
}

func (o HDInsightPropertiesPtrOutput) ToHDInsightPropertiesPtrOutput() HDInsightPropertiesPtrOutput {
	return o
}

func (o HDInsightPropertiesPtrOutput) ToHDInsightPropertiesPtrOutputWithContext(ctx context.Context) HDInsightPropertiesPtrOutput {
	return o
}

func (o HDInsightPropertiesPtrOutput) Elem() HDInsightPropertiesOutput {
	return o.ApplyT(func(v *HDInsightProperties) HDInsightProperties {
		if v != nil {
			return *v
		}
		var ret HDInsightProperties
		return ret
	}).(HDInsightPropertiesOutput)
}

func (o HDInsightPropertiesPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HDInsightProperties) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o HDInsightPropertiesPtrOutput) AdministratorAccount() VirtualMachineSshCredentialsPtrOutput {
	return o.ApplyT(func(v *HDInsightProperties) *VirtualMachineSshCredentials {
		if v == nil {
			return nil
		}
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsPtrOutput)
}

func (o HDInsightPropertiesPtrOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HDInsightProperties) *int {
		if v == nil {
			return nil
		}
		return v.SshPort
	}).(pulumi.IntPtrOutput)
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





type HDInsightResponseInput interface {
	pulumi.Input

	ToHDInsightResponseOutput() HDInsightResponseOutput
	ToHDInsightResponseOutputWithContext(context.Context) HDInsightResponseOutput
}

type HDInsightResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         HDInsightResponsePropertiesPtrInput           `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (HDInsightResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightResponse)(nil)).Elem()
}

func (i HDInsightResponseArgs) ToHDInsightResponseOutput() HDInsightResponseOutput {
	return i.ToHDInsightResponseOutputWithContext(context.Background())
}

func (i HDInsightResponseArgs) ToHDInsightResponseOutputWithContext(ctx context.Context) HDInsightResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightResponseOutput)
}

type HDInsightResponseOutput struct{ *pulumi.OutputState }

func (HDInsightResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightResponse)(nil)).Elem()
}

func (o HDInsightResponseOutput) ToHDInsightResponseOutput() HDInsightResponseOutput {
	return o
}

func (o HDInsightResponseOutput) ToHDInsightResponseOutputWithContext(ctx context.Context) HDInsightResponseOutput {
	return o
}

func (o HDInsightResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o HDInsightResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsightResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o HDInsightResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsightResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o HDInsightResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o HDInsightResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v HDInsightResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o HDInsightResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsightResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o HDInsightResponseOutput) Properties() HDInsightResponsePropertiesPtrOutput {
	return o.ApplyT(func(v HDInsightResponse) *HDInsightResponseProperties { return v.Properties }).(HDInsightResponsePropertiesPtrOutput)
}

func (o HDInsightResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v HDInsightResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o HDInsightResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsightResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o HDInsightResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type HDInsightResponseProperties struct {
	Address              *string                               `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	SshPort              *int                                  `pulumi:"sshPort"`
}





type HDInsightResponsePropertiesInput interface {
	pulumi.Input

	ToHDInsightResponsePropertiesOutput() HDInsightResponsePropertiesOutput
	ToHDInsightResponsePropertiesOutputWithContext(context.Context) HDInsightResponsePropertiesOutput
}

type HDInsightResponsePropertiesArgs struct {
	Address              pulumi.StringPtrInput                        `pulumi:"address"`
	AdministratorAccount VirtualMachineSshCredentialsResponsePtrInput `pulumi:"administratorAccount"`
	SshPort              pulumi.IntPtrInput                           `pulumi:"sshPort"`
}

func (HDInsightResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightResponseProperties)(nil)).Elem()
}

func (i HDInsightResponsePropertiesArgs) ToHDInsightResponsePropertiesOutput() HDInsightResponsePropertiesOutput {
	return i.ToHDInsightResponsePropertiesOutputWithContext(context.Background())
}

func (i HDInsightResponsePropertiesArgs) ToHDInsightResponsePropertiesOutputWithContext(ctx context.Context) HDInsightResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightResponsePropertiesOutput)
}

func (i HDInsightResponsePropertiesArgs) ToHDInsightResponsePropertiesPtrOutput() HDInsightResponsePropertiesPtrOutput {
	return i.ToHDInsightResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i HDInsightResponsePropertiesArgs) ToHDInsightResponsePropertiesPtrOutputWithContext(ctx context.Context) HDInsightResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightResponsePropertiesOutput).ToHDInsightResponsePropertiesPtrOutputWithContext(ctx)
}









type HDInsightResponsePropertiesPtrInput interface {
	pulumi.Input

	ToHDInsightResponsePropertiesPtrOutput() HDInsightResponsePropertiesPtrOutput
	ToHDInsightResponsePropertiesPtrOutputWithContext(context.Context) HDInsightResponsePropertiesPtrOutput
}

type hdinsightResponsePropertiesPtrType HDInsightResponsePropertiesArgs

func HDInsightResponsePropertiesPtr(v *HDInsightResponsePropertiesArgs) HDInsightResponsePropertiesPtrInput {
	return (*hdinsightResponsePropertiesPtrType)(v)
}

func (*hdinsightResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HDInsightResponseProperties)(nil)).Elem()
}

func (i *hdinsightResponsePropertiesPtrType) ToHDInsightResponsePropertiesPtrOutput() HDInsightResponsePropertiesPtrOutput {
	return i.ToHDInsightResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *hdinsightResponsePropertiesPtrType) ToHDInsightResponsePropertiesPtrOutputWithContext(ctx context.Context) HDInsightResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightResponsePropertiesPtrOutput)
}

type HDInsightResponsePropertiesOutput struct{ *pulumi.OutputState }

func (HDInsightResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightResponseProperties)(nil)).Elem()
}

func (o HDInsightResponsePropertiesOutput) ToHDInsightResponsePropertiesOutput() HDInsightResponsePropertiesOutput {
	return o
}

func (o HDInsightResponsePropertiesOutput) ToHDInsightResponsePropertiesOutputWithContext(ctx context.Context) HDInsightResponsePropertiesOutput {
	return o
}

func (o HDInsightResponsePropertiesOutput) ToHDInsightResponsePropertiesPtrOutput() HDInsightResponsePropertiesPtrOutput {
	return o.ToHDInsightResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o HDInsightResponsePropertiesOutput) ToHDInsightResponsePropertiesPtrOutputWithContext(ctx context.Context) HDInsightResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HDInsightResponseProperties) *HDInsightResponseProperties {
		return &v
	}).(HDInsightResponsePropertiesPtrOutput)
}

func (o HDInsightResponsePropertiesOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightResponseProperties) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o HDInsightResponsePropertiesOutput) AdministratorAccount() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyT(func(v HDInsightResponseProperties) *VirtualMachineSshCredentialsResponse {
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o HDInsightResponsePropertiesOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HDInsightResponseProperties) *int { return v.SshPort }).(pulumi.IntPtrOutput)
}

type HDInsightResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (HDInsightResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HDInsightResponseProperties)(nil)).Elem()
}

func (o HDInsightResponsePropertiesPtrOutput) ToHDInsightResponsePropertiesPtrOutput() HDInsightResponsePropertiesPtrOutput {
	return o
}

func (o HDInsightResponsePropertiesPtrOutput) ToHDInsightResponsePropertiesPtrOutputWithContext(ctx context.Context) HDInsightResponsePropertiesPtrOutput {
	return o
}

func (o HDInsightResponsePropertiesPtrOutput) Elem() HDInsightResponsePropertiesOutput {
	return o.ApplyT(func(v *HDInsightResponseProperties) HDInsightResponseProperties {
		if v != nil {
			return *v
		}
		var ret HDInsightResponseProperties
		return ret
	}).(HDInsightResponsePropertiesOutput)
}

func (o HDInsightResponsePropertiesPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HDInsightResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o HDInsightResponsePropertiesPtrOutput) AdministratorAccount() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *HDInsightResponseProperties) *VirtualMachineSshCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o HDInsightResponsePropertiesPtrOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HDInsightResponseProperties) *int {
		if v == nil {
			return nil
		}
		return v.SshPort
	}).(pulumi.IntPtrOutput)
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





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                   `pulumi:"principalId"`
	TenantId               pulumi.StringInput                   `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
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

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
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





type ImageAssetResponseInput interface {
	pulumi.Input

	ToImageAssetResponseOutput() ImageAssetResponseOutput
	ToImageAssetResponseOutputWithContext(context.Context) ImageAssetResponseOutput
}

type ImageAssetResponseArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	MimeType pulumi.StringPtrInput `pulumi:"mimeType"`
	Unpack   pulumi.BoolPtrInput   `pulumi:"unpack"`
	Url      pulumi.StringPtrInput `pulumi:"url"`
}

func (ImageAssetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageAssetResponse)(nil)).Elem()
}

func (i ImageAssetResponseArgs) ToImageAssetResponseOutput() ImageAssetResponseOutput {
	return i.ToImageAssetResponseOutputWithContext(context.Background())
}

func (i ImageAssetResponseArgs) ToImageAssetResponseOutputWithContext(ctx context.Context) ImageAssetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageAssetResponseOutput)
}





type ImageAssetResponseArrayInput interface {
	pulumi.Input

	ToImageAssetResponseArrayOutput() ImageAssetResponseArrayOutput
	ToImageAssetResponseArrayOutputWithContext(context.Context) ImageAssetResponseArrayOutput
}

type ImageAssetResponseArray []ImageAssetResponseInput

func (ImageAssetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageAssetResponse)(nil)).Elem()
}

func (i ImageAssetResponseArray) ToImageAssetResponseArrayOutput() ImageAssetResponseArrayOutput {
	return i.ToImageAssetResponseArrayOutputWithContext(context.Background())
}

func (i ImageAssetResponseArray) ToImageAssetResponseArrayOutputWithContext(ctx context.Context) ImageAssetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageAssetResponseArrayOutput)
}

type ImageAssetResponseOutput struct{ *pulumi.OutputState }

func (ImageAssetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageAssetResponse)(nil)).Elem()
}

func (o ImageAssetResponseOutput) ToImageAssetResponseOutput() ImageAssetResponseOutput {
	return o
}

func (o ImageAssetResponseOutput) ToImageAssetResponseOutputWithContext(ctx context.Context) ImageAssetResponseOutput {
	return o
}

func (o ImageAssetResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageAssetResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ImageAssetResponseOutput) MimeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageAssetResponse) *string { return v.MimeType }).(pulumi.StringPtrOutput)
}

func (o ImageAssetResponseOutput) Unpack() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageAssetResponse) *bool { return v.Unpack }).(pulumi.BoolPtrOutput)
}

func (o ImageAssetResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageAssetResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ImageAssetResponseArrayOutput struct{ *pulumi.OutputState }

func (ImageAssetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageAssetResponse)(nil)).Elem()
}

func (o ImageAssetResponseArrayOutput) ToImageAssetResponseArrayOutput() ImageAssetResponseArrayOutput {
	return o
}

func (o ImageAssetResponseArrayOutput) ToImageAssetResponseArrayOutputWithContext(ctx context.Context) ImageAssetResponseArrayOutput {
	return o
}

func (o ImageAssetResponseArrayOutput) Index(i pulumi.IntInput) ImageAssetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageAssetResponse {
		return vs[0].([]ImageAssetResponse)[vs[1].(int)]
	}).(ImageAssetResponseOutput)
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





type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

type KeyVaultPropertiesResponseArgs struct {
	IdentityClientId pulumi.StringPtrInput `pulumi:"identityClientId"`
	KeyIdentifier    pulumi.StringInput    `pulumi:"keyIdentifier"`
	KeyVaultArmId    pulumi.StringInput    `pulumi:"keyVaultArmId"`
}

func (KeyVaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return i.ToKeyVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput)
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput).ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx)
}









type KeyVaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput
	ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Context) KeyVaultPropertiesResponsePtrOutput
}

type keyVaultPropertiesResponsePtrType KeyVaultPropertiesResponseArgs

func KeyVaultPropertiesResponsePtr(v *KeyVaultPropertiesResponseArgs) KeyVaultPropertiesResponsePtrInput {
	return (*keyVaultPropertiesResponsePtrType)(v)
}

func (*keyVaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponsePtrOutput)
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

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultPropertiesResponse) *KeyVaultPropertiesResponse {
		return &v
	}).(KeyVaultPropertiesResponsePtrOutput)
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





type LinkedWorkspacePropsResponseInput interface {
	pulumi.Input

	ToLinkedWorkspacePropsResponseOutput() LinkedWorkspacePropsResponseOutput
	ToLinkedWorkspacePropsResponseOutputWithContext(context.Context) LinkedWorkspacePropsResponseOutput
}

type LinkedWorkspacePropsResponseArgs struct {
	LinkedWorkspaceResourceId      pulumi.StringPtrInput `pulumi:"linkedWorkspaceResourceId"`
	UserAssignedIdentityResourceId pulumi.StringPtrInput `pulumi:"userAssignedIdentityResourceId"`
}

func (LinkedWorkspacePropsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedWorkspacePropsResponse)(nil)).Elem()
}

func (i LinkedWorkspacePropsResponseArgs) ToLinkedWorkspacePropsResponseOutput() LinkedWorkspacePropsResponseOutput {
	return i.ToLinkedWorkspacePropsResponseOutputWithContext(context.Background())
}

func (i LinkedWorkspacePropsResponseArgs) ToLinkedWorkspacePropsResponseOutputWithContext(ctx context.Context) LinkedWorkspacePropsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedWorkspacePropsResponseOutput)
}

func (i LinkedWorkspacePropsResponseArgs) ToLinkedWorkspacePropsResponsePtrOutput() LinkedWorkspacePropsResponsePtrOutput {
	return i.ToLinkedWorkspacePropsResponsePtrOutputWithContext(context.Background())
}

func (i LinkedWorkspacePropsResponseArgs) ToLinkedWorkspacePropsResponsePtrOutputWithContext(ctx context.Context) LinkedWorkspacePropsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedWorkspacePropsResponseOutput).ToLinkedWorkspacePropsResponsePtrOutputWithContext(ctx)
}









type LinkedWorkspacePropsResponsePtrInput interface {
	pulumi.Input

	ToLinkedWorkspacePropsResponsePtrOutput() LinkedWorkspacePropsResponsePtrOutput
	ToLinkedWorkspacePropsResponsePtrOutputWithContext(context.Context) LinkedWorkspacePropsResponsePtrOutput
}

type linkedWorkspacePropsResponsePtrType LinkedWorkspacePropsResponseArgs

func LinkedWorkspacePropsResponsePtr(v *LinkedWorkspacePropsResponseArgs) LinkedWorkspacePropsResponsePtrInput {
	return (*linkedWorkspacePropsResponsePtrType)(v)
}

func (*linkedWorkspacePropsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedWorkspacePropsResponse)(nil)).Elem()
}

func (i *linkedWorkspacePropsResponsePtrType) ToLinkedWorkspacePropsResponsePtrOutput() LinkedWorkspacePropsResponsePtrOutput {
	return i.ToLinkedWorkspacePropsResponsePtrOutputWithContext(context.Background())
}

func (i *linkedWorkspacePropsResponsePtrType) ToLinkedWorkspacePropsResponsePtrOutputWithContext(ctx context.Context) LinkedWorkspacePropsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedWorkspacePropsResponsePtrOutput)
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

func (o LinkedWorkspacePropsResponseOutput) ToLinkedWorkspacePropsResponsePtrOutput() LinkedWorkspacePropsResponsePtrOutput {
	return o.ToLinkedWorkspacePropsResponsePtrOutputWithContext(context.Background())
}

func (o LinkedWorkspacePropsResponseOutput) ToLinkedWorkspacePropsResponsePtrOutputWithContext(ctx context.Context) LinkedWorkspacePropsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinkedWorkspacePropsResponse) *LinkedWorkspacePropsResponse {
		return &v
	}).(LinkedWorkspacePropsResponsePtrOutput)
}

func (o LinkedWorkspacePropsResponseOutput) LinkedWorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedWorkspacePropsResponse) *string { return v.LinkedWorkspaceResourceId }).(pulumi.StringPtrOutput)
}

func (o LinkedWorkspacePropsResponseOutput) UserAssignedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkedWorkspacePropsResponse) *string { return v.UserAssignedIdentityResourceId }).(pulumi.StringPtrOutput)
}

type LinkedWorkspacePropsResponsePtrOutput struct{ *pulumi.OutputState }

func (LinkedWorkspacePropsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedWorkspacePropsResponse)(nil)).Elem()
}

func (o LinkedWorkspacePropsResponsePtrOutput) ToLinkedWorkspacePropsResponsePtrOutput() LinkedWorkspacePropsResponsePtrOutput {
	return o
}

func (o LinkedWorkspacePropsResponsePtrOutput) ToLinkedWorkspacePropsResponsePtrOutputWithContext(ctx context.Context) LinkedWorkspacePropsResponsePtrOutput {
	return o
}

func (o LinkedWorkspacePropsResponsePtrOutput) Elem() LinkedWorkspacePropsResponseOutput {
	return o.ApplyT(func(v *LinkedWorkspacePropsResponse) LinkedWorkspacePropsResponse {
		if v != nil {
			return *v
		}
		var ret LinkedWorkspacePropsResponse
		return ret
	}).(LinkedWorkspacePropsResponseOutput)
}

func (o LinkedWorkspacePropsResponsePtrOutput) LinkedWorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedWorkspacePropsResponse) *string {
		if v == nil {
			return nil
		}
		return v.LinkedWorkspaceResourceId
	}).(pulumi.StringPtrOutput)
}

func (o LinkedWorkspacePropsResponsePtrOutput) UserAssignedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedWorkspacePropsResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentityResourceId
	}).(pulumi.StringPtrOutput)
}

type MachineLearningServiceErrorResponse struct {
	Error ErrorResponseResponse `pulumi:"error"`
}





type MachineLearningServiceErrorResponseInput interface {
	pulumi.Input

	ToMachineLearningServiceErrorResponseOutput() MachineLearningServiceErrorResponseOutput
	ToMachineLearningServiceErrorResponseOutputWithContext(context.Context) MachineLearningServiceErrorResponseOutput
}

type MachineLearningServiceErrorResponseArgs struct {
	Error ErrorResponseResponseInput `pulumi:"error"`
}

func (MachineLearningServiceErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningServiceErrorResponse)(nil)).Elem()
}

func (i MachineLearningServiceErrorResponseArgs) ToMachineLearningServiceErrorResponseOutput() MachineLearningServiceErrorResponseOutput {
	return i.ToMachineLearningServiceErrorResponseOutputWithContext(context.Background())
}

func (i MachineLearningServiceErrorResponseArgs) ToMachineLearningServiceErrorResponseOutputWithContext(ctx context.Context) MachineLearningServiceErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningServiceErrorResponseOutput)
}





type MachineLearningServiceErrorResponseArrayInput interface {
	pulumi.Input

	ToMachineLearningServiceErrorResponseArrayOutput() MachineLearningServiceErrorResponseArrayOutput
	ToMachineLearningServiceErrorResponseArrayOutputWithContext(context.Context) MachineLearningServiceErrorResponseArrayOutput
}

type MachineLearningServiceErrorResponseArray []MachineLearningServiceErrorResponseInput

func (MachineLearningServiceErrorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineLearningServiceErrorResponse)(nil)).Elem()
}

func (i MachineLearningServiceErrorResponseArray) ToMachineLearningServiceErrorResponseArrayOutput() MachineLearningServiceErrorResponseArrayOutput {
	return i.ToMachineLearningServiceErrorResponseArrayOutputWithContext(context.Background())
}

func (i MachineLearningServiceErrorResponseArray) ToMachineLearningServiceErrorResponseArrayOutputWithContext(ctx context.Context) MachineLearningServiceErrorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningServiceErrorResponseArrayOutput)
}

type MachineLearningServiceErrorResponseOutput struct{ *pulumi.OutputState }

func (MachineLearningServiceErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningServiceErrorResponse)(nil)).Elem()
}

func (o MachineLearningServiceErrorResponseOutput) ToMachineLearningServiceErrorResponseOutput() MachineLearningServiceErrorResponseOutput {
	return o
}

func (o MachineLearningServiceErrorResponseOutput) ToMachineLearningServiceErrorResponseOutputWithContext(ctx context.Context) MachineLearningServiceErrorResponseOutput {
	return o
}

func (o MachineLearningServiceErrorResponseOutput) Error() ErrorResponseResponseOutput {
	return o.ApplyT(func(v MachineLearningServiceErrorResponse) ErrorResponseResponse { return v.Error }).(ErrorResponseResponseOutput)
}

type MachineLearningServiceErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (MachineLearningServiceErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineLearningServiceErrorResponse)(nil)).Elem()
}

func (o MachineLearningServiceErrorResponseArrayOutput) ToMachineLearningServiceErrorResponseArrayOutput() MachineLearningServiceErrorResponseArrayOutput {
	return o
}

func (o MachineLearningServiceErrorResponseArrayOutput) ToMachineLearningServiceErrorResponseArrayOutputWithContext(ctx context.Context) MachineLearningServiceErrorResponseArrayOutput {
	return o
}

func (o MachineLearningServiceErrorResponseArrayOutput) Index(i pulumi.IntInput) MachineLearningServiceErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineLearningServiceErrorResponse {
		return vs[0].([]MachineLearningServiceErrorResponse)[vs[1].(int)]
	}).(MachineLearningServiceErrorResponseOutput)
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





type ModelDockerSectionResponseResponseBaseImageRegistryInput interface {
	pulumi.Input

	ToModelDockerSectionResponseResponseBaseImageRegistryOutput() ModelDockerSectionResponseResponseBaseImageRegistryOutput
	ToModelDockerSectionResponseResponseBaseImageRegistryOutputWithContext(context.Context) ModelDockerSectionResponseResponseBaseImageRegistryOutput
}

type ModelDockerSectionResponseResponseBaseImageRegistryArgs struct {
	Address pulumi.StringPtrInput `pulumi:"address"`
}

func (ModelDockerSectionResponseResponseBaseImageRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelDockerSectionResponseResponseBaseImageRegistry)(nil)).Elem()
}

func (i ModelDockerSectionResponseResponseBaseImageRegistryArgs) ToModelDockerSectionResponseResponseBaseImageRegistryOutput() ModelDockerSectionResponseResponseBaseImageRegistryOutput {
	return i.ToModelDockerSectionResponseResponseBaseImageRegistryOutputWithContext(context.Background())
}

func (i ModelDockerSectionResponseResponseBaseImageRegistryArgs) ToModelDockerSectionResponseResponseBaseImageRegistryOutputWithContext(ctx context.Context) ModelDockerSectionResponseResponseBaseImageRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelDockerSectionResponseResponseBaseImageRegistryOutput)
}

func (i ModelDockerSectionResponseResponseBaseImageRegistryArgs) ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutput() ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput {
	return i.ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutputWithContext(context.Background())
}

func (i ModelDockerSectionResponseResponseBaseImageRegistryArgs) ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutputWithContext(ctx context.Context) ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelDockerSectionResponseResponseBaseImageRegistryOutput).ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutputWithContext(ctx)
}









type ModelDockerSectionResponseResponseBaseImageRegistryPtrInput interface {
	pulumi.Input

	ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutput() ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput
	ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutputWithContext(context.Context) ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput
}

type modelDockerSectionResponseResponseBaseImageRegistryPtrType ModelDockerSectionResponseResponseBaseImageRegistryArgs

func ModelDockerSectionResponseResponseBaseImageRegistryPtr(v *ModelDockerSectionResponseResponseBaseImageRegistryArgs) ModelDockerSectionResponseResponseBaseImageRegistryPtrInput {
	return (*modelDockerSectionResponseResponseBaseImageRegistryPtrType)(v)
}

func (*modelDockerSectionResponseResponseBaseImageRegistryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelDockerSectionResponseResponseBaseImageRegistry)(nil)).Elem()
}

func (i *modelDockerSectionResponseResponseBaseImageRegistryPtrType) ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutput() ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput {
	return i.ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutputWithContext(context.Background())
}

func (i *modelDockerSectionResponseResponseBaseImageRegistryPtrType) ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutputWithContext(ctx context.Context) ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput)
}

type ModelDockerSectionResponseResponseBaseImageRegistryOutput struct{ *pulumi.OutputState }

func (ModelDockerSectionResponseResponseBaseImageRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelDockerSectionResponseResponseBaseImageRegistry)(nil)).Elem()
}

func (o ModelDockerSectionResponseResponseBaseImageRegistryOutput) ToModelDockerSectionResponseResponseBaseImageRegistryOutput() ModelDockerSectionResponseResponseBaseImageRegistryOutput {
	return o
}

func (o ModelDockerSectionResponseResponseBaseImageRegistryOutput) ToModelDockerSectionResponseResponseBaseImageRegistryOutputWithContext(ctx context.Context) ModelDockerSectionResponseResponseBaseImageRegistryOutput {
	return o
}

func (o ModelDockerSectionResponseResponseBaseImageRegistryOutput) ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutput() ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput {
	return o.ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutputWithContext(context.Background())
}

func (o ModelDockerSectionResponseResponseBaseImageRegistryOutput) ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutputWithContext(ctx context.Context) ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelDockerSectionResponseResponseBaseImageRegistry) *ModelDockerSectionResponseResponseBaseImageRegistry {
		return &v
	}).(ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput)
}

func (o ModelDockerSectionResponseResponseBaseImageRegistryOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelDockerSectionResponseResponseBaseImageRegistry) *string { return v.Address }).(pulumi.StringPtrOutput)
}

type ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput struct{ *pulumi.OutputState }

func (ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelDockerSectionResponseResponseBaseImageRegistry)(nil)).Elem()
}

func (o ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput) ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutput() ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput {
	return o
}

func (o ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput) ToModelDockerSectionResponseResponseBaseImageRegistryPtrOutputWithContext(ctx context.Context) ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput {
	return o
}

func (o ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput) Elem() ModelDockerSectionResponseResponseBaseImageRegistryOutput {
	return o.ApplyT(func(v *ModelDockerSectionResponseResponseBaseImageRegistry) ModelDockerSectionResponseResponseBaseImageRegistry {
		if v != nil {
			return *v
		}
		var ret ModelDockerSectionResponseResponseBaseImageRegistry
		return ret
	}).(ModelDockerSectionResponseResponseBaseImageRegistryOutput)
}

func (o ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelDockerSectionResponseResponseBaseImageRegistry) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
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





type ModelEnvironmentDefinitionResponseResponseDockerInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionResponseResponseDockerOutput() ModelEnvironmentDefinitionResponseResponseDockerOutput
	ToModelEnvironmentDefinitionResponseResponseDockerOutputWithContext(context.Context) ModelEnvironmentDefinitionResponseResponseDockerOutput
}

type ModelEnvironmentDefinitionResponseResponseDockerArgs struct {
	BaseDockerfile    pulumi.StringPtrInput                                       `pulumi:"baseDockerfile"`
	BaseImage         pulumi.StringPtrInput                                       `pulumi:"baseImage"`
	BaseImageRegistry ModelDockerSectionResponseResponseBaseImageRegistryPtrInput `pulumi:"baseImageRegistry"`
}

func (ModelEnvironmentDefinitionResponseResponseDockerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseDocker)(nil)).Elem()
}

func (i ModelEnvironmentDefinitionResponseResponseDockerArgs) ToModelEnvironmentDefinitionResponseResponseDockerOutput() ModelEnvironmentDefinitionResponseResponseDockerOutput {
	return i.ToModelEnvironmentDefinitionResponseResponseDockerOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionResponseResponseDockerArgs) ToModelEnvironmentDefinitionResponseResponseDockerOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseDockerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponseDockerOutput)
}

func (i ModelEnvironmentDefinitionResponseResponseDockerArgs) ToModelEnvironmentDefinitionResponseResponseDockerPtrOutput() ModelEnvironmentDefinitionResponseResponseDockerPtrOutput {
	return i.ToModelEnvironmentDefinitionResponseResponseDockerPtrOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionResponseResponseDockerArgs) ToModelEnvironmentDefinitionResponseResponseDockerPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseDockerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponseDockerOutput).ToModelEnvironmentDefinitionResponseResponseDockerPtrOutputWithContext(ctx)
}









type ModelEnvironmentDefinitionResponseResponseDockerPtrInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionResponseResponseDockerPtrOutput() ModelEnvironmentDefinitionResponseResponseDockerPtrOutput
	ToModelEnvironmentDefinitionResponseResponseDockerPtrOutputWithContext(context.Context) ModelEnvironmentDefinitionResponseResponseDockerPtrOutput
}

type modelEnvironmentDefinitionResponseResponseDockerPtrType ModelEnvironmentDefinitionResponseResponseDockerArgs

func ModelEnvironmentDefinitionResponseResponseDockerPtr(v *ModelEnvironmentDefinitionResponseResponseDockerArgs) ModelEnvironmentDefinitionResponseResponseDockerPtrInput {
	return (*modelEnvironmentDefinitionResponseResponseDockerPtrType)(v)
}

func (*modelEnvironmentDefinitionResponseResponseDockerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionResponseResponseDocker)(nil)).Elem()
}

func (i *modelEnvironmentDefinitionResponseResponseDockerPtrType) ToModelEnvironmentDefinitionResponseResponseDockerPtrOutput() ModelEnvironmentDefinitionResponseResponseDockerPtrOutput {
	return i.ToModelEnvironmentDefinitionResponseResponseDockerPtrOutputWithContext(context.Background())
}

func (i *modelEnvironmentDefinitionResponseResponseDockerPtrType) ToModelEnvironmentDefinitionResponseResponseDockerPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseDockerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponseDockerPtrOutput)
}

type ModelEnvironmentDefinitionResponseResponseDockerOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionResponseResponseDockerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseDocker)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionResponseResponseDockerOutput) ToModelEnvironmentDefinitionResponseResponseDockerOutput() ModelEnvironmentDefinitionResponseResponseDockerOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseDockerOutput) ToModelEnvironmentDefinitionResponseResponseDockerOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseDockerOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseDockerOutput) ToModelEnvironmentDefinitionResponseResponseDockerPtrOutput() ModelEnvironmentDefinitionResponseResponseDockerPtrOutput {
	return o.ToModelEnvironmentDefinitionResponseResponseDockerPtrOutputWithContext(context.Background())
}

func (o ModelEnvironmentDefinitionResponseResponseDockerOutput) ToModelEnvironmentDefinitionResponseResponseDockerPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseDockerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelEnvironmentDefinitionResponseResponseDocker) *ModelEnvironmentDefinitionResponseResponseDocker {
		return &v
	}).(ModelEnvironmentDefinitionResponseResponseDockerPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseDockerOutput) BaseDockerfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseDocker) *string { return v.BaseDockerfile }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseDockerOutput) BaseImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseDocker) *string { return v.BaseImage }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseDockerOutput) BaseImageRegistry() ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseDocker) *ModelDockerSectionResponseResponseBaseImageRegistry {
		return v.BaseImageRegistry
	}).(ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput)
}

type ModelEnvironmentDefinitionResponseResponseDockerPtrOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionResponseResponseDockerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionResponseResponseDocker)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionResponseResponseDockerPtrOutput) ToModelEnvironmentDefinitionResponseResponseDockerPtrOutput() ModelEnvironmentDefinitionResponseResponseDockerPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseDockerPtrOutput) ToModelEnvironmentDefinitionResponseResponseDockerPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseDockerPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseDockerPtrOutput) Elem() ModelEnvironmentDefinitionResponseResponseDockerOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseDocker) ModelEnvironmentDefinitionResponseResponseDocker {
		if v != nil {
			return *v
		}
		var ret ModelEnvironmentDefinitionResponseResponseDocker
		return ret
	}).(ModelEnvironmentDefinitionResponseResponseDockerOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseDockerPtrOutput) BaseDockerfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseDocker) *string {
		if v == nil {
			return nil
		}
		return v.BaseDockerfile
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseDockerPtrOutput) BaseImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseDocker) *string {
		if v == nil {
			return nil
		}
		return v.BaseImage
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseDockerPtrOutput) BaseImageRegistry() ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseDocker) *ModelDockerSectionResponseResponseBaseImageRegistry {
		if v == nil {
			return nil
		}
		return v.BaseImageRegistry
	}).(ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput)
}

type ModelEnvironmentDefinitionResponseResponsePython struct {
	BaseCondaEnvironment    *string     `pulumi:"baseCondaEnvironment"`
	CondaDependencies       interface{} `pulumi:"condaDependencies"`
	InterpreterPath         *string     `pulumi:"interpreterPath"`
	UserManagedDependencies *bool       `pulumi:"userManagedDependencies"`
}





type ModelEnvironmentDefinitionResponseResponsePythonInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionResponseResponsePythonOutput() ModelEnvironmentDefinitionResponseResponsePythonOutput
	ToModelEnvironmentDefinitionResponseResponsePythonOutputWithContext(context.Context) ModelEnvironmentDefinitionResponseResponsePythonOutput
}

type ModelEnvironmentDefinitionResponseResponsePythonArgs struct {
	BaseCondaEnvironment    pulumi.StringPtrInput `pulumi:"baseCondaEnvironment"`
	CondaDependencies       pulumi.Input          `pulumi:"condaDependencies"`
	InterpreterPath         pulumi.StringPtrInput `pulumi:"interpreterPath"`
	UserManagedDependencies pulumi.BoolPtrInput   `pulumi:"userManagedDependencies"`
}

func (ModelEnvironmentDefinitionResponseResponsePythonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponsePython)(nil)).Elem()
}

func (i ModelEnvironmentDefinitionResponseResponsePythonArgs) ToModelEnvironmentDefinitionResponseResponsePythonOutput() ModelEnvironmentDefinitionResponseResponsePythonOutput {
	return i.ToModelEnvironmentDefinitionResponseResponsePythonOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionResponseResponsePythonArgs) ToModelEnvironmentDefinitionResponseResponsePythonOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponsePythonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponsePythonOutput)
}

func (i ModelEnvironmentDefinitionResponseResponsePythonArgs) ToModelEnvironmentDefinitionResponseResponsePythonPtrOutput() ModelEnvironmentDefinitionResponseResponsePythonPtrOutput {
	return i.ToModelEnvironmentDefinitionResponseResponsePythonPtrOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionResponseResponsePythonArgs) ToModelEnvironmentDefinitionResponseResponsePythonPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponsePythonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponsePythonOutput).ToModelEnvironmentDefinitionResponseResponsePythonPtrOutputWithContext(ctx)
}









type ModelEnvironmentDefinitionResponseResponsePythonPtrInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionResponseResponsePythonPtrOutput() ModelEnvironmentDefinitionResponseResponsePythonPtrOutput
	ToModelEnvironmentDefinitionResponseResponsePythonPtrOutputWithContext(context.Context) ModelEnvironmentDefinitionResponseResponsePythonPtrOutput
}

type modelEnvironmentDefinitionResponseResponsePythonPtrType ModelEnvironmentDefinitionResponseResponsePythonArgs

func ModelEnvironmentDefinitionResponseResponsePythonPtr(v *ModelEnvironmentDefinitionResponseResponsePythonArgs) ModelEnvironmentDefinitionResponseResponsePythonPtrInput {
	return (*modelEnvironmentDefinitionResponseResponsePythonPtrType)(v)
}

func (*modelEnvironmentDefinitionResponseResponsePythonPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionResponseResponsePython)(nil)).Elem()
}

func (i *modelEnvironmentDefinitionResponseResponsePythonPtrType) ToModelEnvironmentDefinitionResponseResponsePythonPtrOutput() ModelEnvironmentDefinitionResponseResponsePythonPtrOutput {
	return i.ToModelEnvironmentDefinitionResponseResponsePythonPtrOutputWithContext(context.Background())
}

func (i *modelEnvironmentDefinitionResponseResponsePythonPtrType) ToModelEnvironmentDefinitionResponseResponsePythonPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponsePythonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponsePythonPtrOutput)
}

type ModelEnvironmentDefinitionResponseResponsePythonOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionResponseResponsePythonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponsePython)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionResponseResponsePythonOutput) ToModelEnvironmentDefinitionResponseResponsePythonOutput() ModelEnvironmentDefinitionResponseResponsePythonOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponsePythonOutput) ToModelEnvironmentDefinitionResponseResponsePythonOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponsePythonOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponsePythonOutput) ToModelEnvironmentDefinitionResponseResponsePythonPtrOutput() ModelEnvironmentDefinitionResponseResponsePythonPtrOutput {
	return o.ToModelEnvironmentDefinitionResponseResponsePythonPtrOutputWithContext(context.Background())
}

func (o ModelEnvironmentDefinitionResponseResponsePythonOutput) ToModelEnvironmentDefinitionResponseResponsePythonPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponsePythonPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelEnvironmentDefinitionResponseResponsePython) *ModelEnvironmentDefinitionResponseResponsePython {
		return &v
	}).(ModelEnvironmentDefinitionResponseResponsePythonPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponsePythonOutput) BaseCondaEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponsePython) *string { return v.BaseCondaEnvironment }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponsePythonOutput) CondaDependencies() pulumi.AnyOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponsePython) interface{} { return v.CondaDependencies }).(pulumi.AnyOutput)
}

func (o ModelEnvironmentDefinitionResponseResponsePythonOutput) InterpreterPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponsePython) *string { return v.InterpreterPath }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponsePythonOutput) UserManagedDependencies() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponsePython) *bool { return v.UserManagedDependencies }).(pulumi.BoolPtrOutput)
}

type ModelEnvironmentDefinitionResponseResponsePythonPtrOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionResponseResponsePythonPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionResponseResponsePython)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionResponseResponsePythonPtrOutput) ToModelEnvironmentDefinitionResponseResponsePythonPtrOutput() ModelEnvironmentDefinitionResponseResponsePythonPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponsePythonPtrOutput) ToModelEnvironmentDefinitionResponseResponsePythonPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponsePythonPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponsePythonPtrOutput) Elem() ModelEnvironmentDefinitionResponseResponsePythonOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponsePython) ModelEnvironmentDefinitionResponseResponsePython {
		if v != nil {
			return *v
		}
		var ret ModelEnvironmentDefinitionResponseResponsePython
		return ret
	}).(ModelEnvironmentDefinitionResponseResponsePythonOutput)
}

func (o ModelEnvironmentDefinitionResponseResponsePythonPtrOutput) BaseCondaEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponsePython) *string {
		if v == nil {
			return nil
		}
		return v.BaseCondaEnvironment
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponsePythonPtrOutput) CondaDependencies() pulumi.AnyOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponsePython) interface{} {
		if v == nil {
			return nil
		}
		return v.CondaDependencies
	}).(pulumi.AnyOutput)
}

func (o ModelEnvironmentDefinitionResponseResponsePythonPtrOutput) InterpreterPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponsePython) *string {
		if v == nil {
			return nil
		}
		return v.InterpreterPath
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponsePythonPtrOutput) UserManagedDependencies() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponsePython) *bool {
		if v == nil {
			return nil
		}
		return v.UserManagedDependencies
	}).(pulumi.BoolPtrOutput)
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





type ModelEnvironmentDefinitionResponseResponseRInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionResponseResponseROutput() ModelEnvironmentDefinitionResponseResponseROutput
	ToModelEnvironmentDefinitionResponseResponseROutputWithContext(context.Context) ModelEnvironmentDefinitionResponseResponseROutput
}

type ModelEnvironmentDefinitionResponseResponseRArgs struct {
	BioConductorPackages pulumi.StringArrayInput                  `pulumi:"bioConductorPackages"`
	CranPackages         RCranPackageResponseArrayInput           `pulumi:"cranPackages"`
	CustomUrlPackages    pulumi.StringArrayInput                  `pulumi:"customUrlPackages"`
	GitHubPackages       RGitHubPackageResponseResponseArrayInput `pulumi:"gitHubPackages"`
	RVersion             pulumi.StringPtrInput                    `pulumi:"rVersion"`
	RscriptPath          pulumi.StringPtrInput                    `pulumi:"rscriptPath"`
	SnapshotDate         pulumi.StringPtrInput                    `pulumi:"snapshotDate"`
	UserManaged          pulumi.BoolPtrInput                      `pulumi:"userManaged"`
}

func (ModelEnvironmentDefinitionResponseResponseRArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseR)(nil)).Elem()
}

func (i ModelEnvironmentDefinitionResponseResponseRArgs) ToModelEnvironmentDefinitionResponseResponseROutput() ModelEnvironmentDefinitionResponseResponseROutput {
	return i.ToModelEnvironmentDefinitionResponseResponseROutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionResponseResponseRArgs) ToModelEnvironmentDefinitionResponseResponseROutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseROutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponseROutput)
}

func (i ModelEnvironmentDefinitionResponseResponseRArgs) ToModelEnvironmentDefinitionResponseResponseRPtrOutput() ModelEnvironmentDefinitionResponseResponseRPtrOutput {
	return i.ToModelEnvironmentDefinitionResponseResponseRPtrOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionResponseResponseRArgs) ToModelEnvironmentDefinitionResponseResponseRPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseRPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponseROutput).ToModelEnvironmentDefinitionResponseResponseRPtrOutputWithContext(ctx)
}









type ModelEnvironmentDefinitionResponseResponseRPtrInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionResponseResponseRPtrOutput() ModelEnvironmentDefinitionResponseResponseRPtrOutput
	ToModelEnvironmentDefinitionResponseResponseRPtrOutputWithContext(context.Context) ModelEnvironmentDefinitionResponseResponseRPtrOutput
}

type modelEnvironmentDefinitionResponseResponseRPtrType ModelEnvironmentDefinitionResponseResponseRArgs

func ModelEnvironmentDefinitionResponseResponseRPtr(v *ModelEnvironmentDefinitionResponseResponseRArgs) ModelEnvironmentDefinitionResponseResponseRPtrInput {
	return (*modelEnvironmentDefinitionResponseResponseRPtrType)(v)
}

func (*modelEnvironmentDefinitionResponseResponseRPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionResponseResponseR)(nil)).Elem()
}

func (i *modelEnvironmentDefinitionResponseResponseRPtrType) ToModelEnvironmentDefinitionResponseResponseRPtrOutput() ModelEnvironmentDefinitionResponseResponseRPtrOutput {
	return i.ToModelEnvironmentDefinitionResponseResponseRPtrOutputWithContext(context.Background())
}

func (i *modelEnvironmentDefinitionResponseResponseRPtrType) ToModelEnvironmentDefinitionResponseResponseRPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseRPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponseRPtrOutput)
}

type ModelEnvironmentDefinitionResponseResponseROutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionResponseResponseROutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseR)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) ToModelEnvironmentDefinitionResponseResponseROutput() ModelEnvironmentDefinitionResponseResponseROutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) ToModelEnvironmentDefinitionResponseResponseROutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseROutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) ToModelEnvironmentDefinitionResponseResponseRPtrOutput() ModelEnvironmentDefinitionResponseResponseRPtrOutput {
	return o.ToModelEnvironmentDefinitionResponseResponseRPtrOutputWithContext(context.Background())
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) ToModelEnvironmentDefinitionResponseResponseRPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseRPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelEnvironmentDefinitionResponseResponseR) *ModelEnvironmentDefinitionResponseResponseR {
		return &v
	}).(ModelEnvironmentDefinitionResponseResponseRPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) BioConductorPackages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseR) []string { return v.BioConductorPackages }).(pulumi.StringArrayOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) CranPackages() RCranPackageResponseArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseR) []RCranPackageResponse { return v.CranPackages }).(RCranPackageResponseArrayOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) CustomUrlPackages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseR) []string { return v.CustomUrlPackages }).(pulumi.StringArrayOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) GitHubPackages() RGitHubPackageResponseResponseArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseR) []RGitHubPackageResponseResponse {
		return v.GitHubPackages
	}).(RGitHubPackageResponseResponseArrayOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) RVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseR) *string { return v.RVersion }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) RscriptPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseR) *string { return v.RscriptPath }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) SnapshotDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseR) *string { return v.SnapshotDate }).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseROutput) UserManaged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseR) *bool { return v.UserManaged }).(pulumi.BoolPtrOutput)
}

type ModelEnvironmentDefinitionResponseResponseRPtrOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionResponseResponseRPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionResponseResponseR)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionResponseResponseRPtrOutput) ToModelEnvironmentDefinitionResponseResponseRPtrOutput() ModelEnvironmentDefinitionResponseResponseRPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseRPtrOutput) ToModelEnvironmentDefinitionResponseResponseRPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseRPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseRPtrOutput) Elem() ModelEnvironmentDefinitionResponseResponseROutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseR) ModelEnvironmentDefinitionResponseResponseR {
		if v != nil {
			return *v
		}
		var ret ModelEnvironmentDefinitionResponseResponseR
		return ret
	}).(ModelEnvironmentDefinitionResponseResponseROutput)
}

func (o ModelEnvironmentDefinitionResponseResponseRPtrOutput) BioConductorPackages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseR) []string {
		if v == nil {
			return nil
		}
		return v.BioConductorPackages
	}).(pulumi.StringArrayOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseRPtrOutput) CranPackages() RCranPackageResponseArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseR) []RCranPackageResponse {
		if v == nil {
			return nil
		}
		return v.CranPackages
	}).(RCranPackageResponseArrayOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseRPtrOutput) CustomUrlPackages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseR) []string {
		if v == nil {
			return nil
		}
		return v.CustomUrlPackages
	}).(pulumi.StringArrayOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseRPtrOutput) GitHubPackages() RGitHubPackageResponseResponseArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseR) []RGitHubPackageResponseResponse {
		if v == nil {
			return nil
		}
		return v.GitHubPackages
	}).(RGitHubPackageResponseResponseArrayOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseRPtrOutput) RVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseR) *string {
		if v == nil {
			return nil
		}
		return v.RVersion
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseRPtrOutput) RscriptPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseR) *string {
		if v == nil {
			return nil
		}
		return v.RscriptPath
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseRPtrOutput) SnapshotDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseR) *string {
		if v == nil {
			return nil
		}
		return v.SnapshotDate
	}).(pulumi.StringPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseRPtrOutput) UserManaged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseR) *bool {
		if v == nil {
			return nil
		}
		return v.UserManaged
	}).(pulumi.BoolPtrOutput)
}

type ModelEnvironmentDefinitionResponseResponseSpark struct {
	Packages         []SparkMavenPackageResponse `pulumi:"packages"`
	PrecachePackages *bool                       `pulumi:"precachePackages"`
	Repositories     []string                    `pulumi:"repositories"`
}





type ModelEnvironmentDefinitionResponseResponseSparkInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionResponseResponseSparkOutput() ModelEnvironmentDefinitionResponseResponseSparkOutput
	ToModelEnvironmentDefinitionResponseResponseSparkOutputWithContext(context.Context) ModelEnvironmentDefinitionResponseResponseSparkOutput
}

type ModelEnvironmentDefinitionResponseResponseSparkArgs struct {
	Packages         SparkMavenPackageResponseArrayInput `pulumi:"packages"`
	PrecachePackages pulumi.BoolPtrInput                 `pulumi:"precachePackages"`
	Repositories     pulumi.StringArrayInput             `pulumi:"repositories"`
}

func (ModelEnvironmentDefinitionResponseResponseSparkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseSpark)(nil)).Elem()
}

func (i ModelEnvironmentDefinitionResponseResponseSparkArgs) ToModelEnvironmentDefinitionResponseResponseSparkOutput() ModelEnvironmentDefinitionResponseResponseSparkOutput {
	return i.ToModelEnvironmentDefinitionResponseResponseSparkOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionResponseResponseSparkArgs) ToModelEnvironmentDefinitionResponseResponseSparkOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseSparkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponseSparkOutput)
}

func (i ModelEnvironmentDefinitionResponseResponseSparkArgs) ToModelEnvironmentDefinitionResponseResponseSparkPtrOutput() ModelEnvironmentDefinitionResponseResponseSparkPtrOutput {
	return i.ToModelEnvironmentDefinitionResponseResponseSparkPtrOutputWithContext(context.Background())
}

func (i ModelEnvironmentDefinitionResponseResponseSparkArgs) ToModelEnvironmentDefinitionResponseResponseSparkPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseSparkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponseSparkOutput).ToModelEnvironmentDefinitionResponseResponseSparkPtrOutputWithContext(ctx)
}









type ModelEnvironmentDefinitionResponseResponseSparkPtrInput interface {
	pulumi.Input

	ToModelEnvironmentDefinitionResponseResponseSparkPtrOutput() ModelEnvironmentDefinitionResponseResponseSparkPtrOutput
	ToModelEnvironmentDefinitionResponseResponseSparkPtrOutputWithContext(context.Context) ModelEnvironmentDefinitionResponseResponseSparkPtrOutput
}

type modelEnvironmentDefinitionResponseResponseSparkPtrType ModelEnvironmentDefinitionResponseResponseSparkArgs

func ModelEnvironmentDefinitionResponseResponseSparkPtr(v *ModelEnvironmentDefinitionResponseResponseSparkArgs) ModelEnvironmentDefinitionResponseResponseSparkPtrInput {
	return (*modelEnvironmentDefinitionResponseResponseSparkPtrType)(v)
}

func (*modelEnvironmentDefinitionResponseResponseSparkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionResponseResponseSpark)(nil)).Elem()
}

func (i *modelEnvironmentDefinitionResponseResponseSparkPtrType) ToModelEnvironmentDefinitionResponseResponseSparkPtrOutput() ModelEnvironmentDefinitionResponseResponseSparkPtrOutput {
	return i.ToModelEnvironmentDefinitionResponseResponseSparkPtrOutputWithContext(context.Background())
}

func (i *modelEnvironmentDefinitionResponseResponseSparkPtrType) ToModelEnvironmentDefinitionResponseResponseSparkPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseSparkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelEnvironmentDefinitionResponseResponseSparkPtrOutput)
}

type ModelEnvironmentDefinitionResponseResponseSparkOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionResponseResponseSparkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseSpark)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionResponseResponseSparkOutput) ToModelEnvironmentDefinitionResponseResponseSparkOutput() ModelEnvironmentDefinitionResponseResponseSparkOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseSparkOutput) ToModelEnvironmentDefinitionResponseResponseSparkOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseSparkOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseSparkOutput) ToModelEnvironmentDefinitionResponseResponseSparkPtrOutput() ModelEnvironmentDefinitionResponseResponseSparkPtrOutput {
	return o.ToModelEnvironmentDefinitionResponseResponseSparkPtrOutputWithContext(context.Background())
}

func (o ModelEnvironmentDefinitionResponseResponseSparkOutput) ToModelEnvironmentDefinitionResponseResponseSparkPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseSparkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModelEnvironmentDefinitionResponseResponseSpark) *ModelEnvironmentDefinitionResponseResponseSpark {
		return &v
	}).(ModelEnvironmentDefinitionResponseResponseSparkPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseSparkOutput) Packages() SparkMavenPackageResponseArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseSpark) []SparkMavenPackageResponse { return v.Packages }).(SparkMavenPackageResponseArrayOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseSparkOutput) PrecachePackages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseSpark) *bool { return v.PrecachePackages }).(pulumi.BoolPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseSparkOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ModelEnvironmentDefinitionResponseResponseSpark) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

type ModelEnvironmentDefinitionResponseResponseSparkPtrOutput struct{ *pulumi.OutputState }

func (ModelEnvironmentDefinitionResponseResponseSparkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelEnvironmentDefinitionResponseResponseSpark)(nil)).Elem()
}

func (o ModelEnvironmentDefinitionResponseResponseSparkPtrOutput) ToModelEnvironmentDefinitionResponseResponseSparkPtrOutput() ModelEnvironmentDefinitionResponseResponseSparkPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseSparkPtrOutput) ToModelEnvironmentDefinitionResponseResponseSparkPtrOutputWithContext(ctx context.Context) ModelEnvironmentDefinitionResponseResponseSparkPtrOutput {
	return o
}

func (o ModelEnvironmentDefinitionResponseResponseSparkPtrOutput) Elem() ModelEnvironmentDefinitionResponseResponseSparkOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseSpark) ModelEnvironmentDefinitionResponseResponseSpark {
		if v != nil {
			return *v
		}
		var ret ModelEnvironmentDefinitionResponseResponseSpark
		return ret
	}).(ModelEnvironmentDefinitionResponseResponseSparkOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseSparkPtrOutput) Packages() SparkMavenPackageResponseArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseSpark) []SparkMavenPackageResponse {
		if v == nil {
			return nil
		}
		return v.Packages
	}).(SparkMavenPackageResponseArrayOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseSparkPtrOutput) PrecachePackages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseSpark) *bool {
		if v == nil {
			return nil
		}
		return v.PrecachePackages
	}).(pulumi.BoolPtrOutput)
}

func (o ModelEnvironmentDefinitionResponseResponseSparkPtrOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ModelEnvironmentDefinitionResponseResponseSpark) []string {
		if v == nil {
			return nil
		}
		return v.Repositories
	}).(pulumi.StringArrayOutput)
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





type ModelResponseInput interface {
	pulumi.Input

	ToModelResponseOutput() ModelResponseOutput
	ToModelResponseOutputWithContext(context.Context) ModelResponseOutput
}

type ModelResponseArgs struct {
	CreatedTime          pulumi.StringPtrInput                         `pulumi:"createdTime"`
	Datasets             DatasetReferenceResponseArrayInput            `pulumi:"datasets"`
	DerivedModelIds      pulumi.StringArrayInput                       `pulumi:"derivedModelIds"`
	Description          pulumi.StringPtrInput                         `pulumi:"description"`
	ExperimentName       pulumi.StringPtrInput                         `pulumi:"experimentName"`
	Framework            pulumi.StringPtrInput                         `pulumi:"framework"`
	FrameworkVersion     pulumi.StringPtrInput                         `pulumi:"frameworkVersion"`
	Id                   pulumi.StringPtrInput                         `pulumi:"id"`
	KvTags               pulumi.StringMapInput                         `pulumi:"kvTags"`
	MimeType             pulumi.StringInput                            `pulumi:"mimeType"`
	ModifiedTime         pulumi.StringPtrInput                         `pulumi:"modifiedTime"`
	Name                 pulumi.StringInput                            `pulumi:"name"`
	ParentModelId        pulumi.StringPtrInput                         `pulumi:"parentModelId"`
	Properties           pulumi.StringMapInput                         `pulumi:"properties"`
	ResourceRequirements ContainerResourceRequirementsResponsePtrInput `pulumi:"resourceRequirements"`
	RunId                pulumi.StringPtrInput                         `pulumi:"runId"`
	SampleInputData      pulumi.StringPtrInput                         `pulumi:"sampleInputData"`
	SampleOutputData     pulumi.StringPtrInput                         `pulumi:"sampleOutputData"`
	Unpack               pulumi.BoolPtrInput                           `pulumi:"unpack"`
	Url                  pulumi.StringInput                            `pulumi:"url"`
	Version              pulumi.Float64PtrInput                        `pulumi:"version"`
}

func (ModelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelResponse)(nil)).Elem()
}

func (i ModelResponseArgs) ToModelResponseOutput() ModelResponseOutput {
	return i.ToModelResponseOutputWithContext(context.Background())
}

func (i ModelResponseArgs) ToModelResponseOutputWithContext(ctx context.Context) ModelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelResponseOutput)
}





type ModelResponseArrayInput interface {
	pulumi.Input

	ToModelResponseArrayOutput() ModelResponseArrayOutput
	ToModelResponseArrayOutputWithContext(context.Context) ModelResponseArrayOutput
}

type ModelResponseArray []ModelResponseInput

func (ModelResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ModelResponse)(nil)).Elem()
}

func (i ModelResponseArray) ToModelResponseArrayOutput() ModelResponseArrayOutput {
	return i.ToModelResponseArrayOutputWithContext(context.Background())
}

func (i ModelResponseArray) ToModelResponseArrayOutputWithContext(ctx context.Context) ModelResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelResponseArrayOutput)
}

type ModelResponseOutput struct{ *pulumi.OutputState }

func (ModelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelResponse)(nil)).Elem()
}

func (o ModelResponseOutput) ToModelResponseOutput() ModelResponseOutput {
	return o
}

func (o ModelResponseOutput) ToModelResponseOutputWithContext(ctx context.Context) ModelResponseOutput {
	return o
}

func (o ModelResponseOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelResponse) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o ModelResponseOutput) Datasets() DatasetReferenceResponseArrayOutput {
	return o.ApplyT(func(v ModelResponse) []DatasetReferenceResponse { return v.Datasets }).(DatasetReferenceResponseArrayOutput)
}

func (o ModelResponseOutput) DerivedModelIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ModelResponse) []string { return v.DerivedModelIds }).(pulumi.StringArrayOutput)
}

func (o ModelResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ModelResponseOutput) ExperimentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelResponse) *string { return v.ExperimentName }).(pulumi.StringPtrOutput)
}

func (o ModelResponseOutput) Framework() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelResponse) *string { return v.Framework }).(pulumi.StringPtrOutput)
}

func (o ModelResponseOutput) FrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelResponse) *string { return v.FrameworkVersion }).(pulumi.StringPtrOutput)
}

func (o ModelResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ModelResponseOutput) KvTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelResponse) map[string]string { return v.KvTags }).(pulumi.StringMapOutput)
}

func (o ModelResponseOutput) MimeType() pulumi.StringOutput {
	return o.ApplyT(func(v ModelResponse) string { return v.MimeType }).(pulumi.StringOutput)
}

func (o ModelResponseOutput) ModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelResponse) *string { return v.ModifiedTime }).(pulumi.StringPtrOutput)
}

func (o ModelResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ModelResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ModelResponseOutput) ParentModelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelResponse) *string { return v.ParentModelId }).(pulumi.StringPtrOutput)
}

func (o ModelResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ModelResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ModelResponseOutput) ResourceRequirements() ContainerResourceRequirementsResponsePtrOutput {
	return o.ApplyT(func(v ModelResponse) *ContainerResourceRequirementsResponse { return v.ResourceRequirements }).(ContainerResourceRequirementsResponsePtrOutput)
}

func (o ModelResponseOutput) RunId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelResponse) *string { return v.RunId }).(pulumi.StringPtrOutput)
}

func (o ModelResponseOutput) SampleInputData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelResponse) *string { return v.SampleInputData }).(pulumi.StringPtrOutput)
}

func (o ModelResponseOutput) SampleOutputData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModelResponse) *string { return v.SampleOutputData }).(pulumi.StringPtrOutput)
}

func (o ModelResponseOutput) Unpack() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ModelResponse) *bool { return v.Unpack }).(pulumi.BoolPtrOutput)
}

func (o ModelResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v ModelResponse) string { return v.Url }).(pulumi.StringOutput)
}

func (o ModelResponseOutput) Version() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ModelResponse) *float64 { return v.Version }).(pulumi.Float64PtrOutput)
}

type ModelResponseArrayOutput struct{ *pulumi.OutputState }

func (ModelResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ModelResponse)(nil)).Elem()
}

func (o ModelResponseArrayOutput) ToModelResponseArrayOutput() ModelResponseArrayOutput {
	return o
}

func (o ModelResponseArrayOutput) ToModelResponseArrayOutputWithContext(ctx context.Context) ModelResponseArrayOutput {
	return o
}

func (o ModelResponseArrayOutput) Index(i pulumi.IntInput) ModelResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ModelResponse {
		return vs[0].([]ModelResponse)[vs[1].(int)]
	}).(ModelResponseOutput)
}

type NodeStateCountsResponse struct {
	IdleNodeCount      int `pulumi:"idleNodeCount"`
	LeavingNodeCount   int `pulumi:"leavingNodeCount"`
	PreemptedNodeCount int `pulumi:"preemptedNodeCount"`
	PreparingNodeCount int `pulumi:"preparingNodeCount"`
	RunningNodeCount   int `pulumi:"runningNodeCount"`
	UnusableNodeCount  int `pulumi:"unusableNodeCount"`
}





type NodeStateCountsResponseInput interface {
	pulumi.Input

	ToNodeStateCountsResponseOutput() NodeStateCountsResponseOutput
	ToNodeStateCountsResponseOutputWithContext(context.Context) NodeStateCountsResponseOutput
}

type NodeStateCountsResponseArgs struct {
	IdleNodeCount      pulumi.IntInput `pulumi:"idleNodeCount"`
	LeavingNodeCount   pulumi.IntInput `pulumi:"leavingNodeCount"`
	PreemptedNodeCount pulumi.IntInput `pulumi:"preemptedNodeCount"`
	PreparingNodeCount pulumi.IntInput `pulumi:"preparingNodeCount"`
	RunningNodeCount   pulumi.IntInput `pulumi:"runningNodeCount"`
	UnusableNodeCount  pulumi.IntInput `pulumi:"unusableNodeCount"`
}

func (NodeStateCountsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeStateCountsResponse)(nil)).Elem()
}

func (i NodeStateCountsResponseArgs) ToNodeStateCountsResponseOutput() NodeStateCountsResponseOutput {
	return i.ToNodeStateCountsResponseOutputWithContext(context.Background())
}

func (i NodeStateCountsResponseArgs) ToNodeStateCountsResponseOutputWithContext(ctx context.Context) NodeStateCountsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeStateCountsResponseOutput)
}

func (i NodeStateCountsResponseArgs) ToNodeStateCountsResponsePtrOutput() NodeStateCountsResponsePtrOutput {
	return i.ToNodeStateCountsResponsePtrOutputWithContext(context.Background())
}

func (i NodeStateCountsResponseArgs) ToNodeStateCountsResponsePtrOutputWithContext(ctx context.Context) NodeStateCountsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeStateCountsResponseOutput).ToNodeStateCountsResponsePtrOutputWithContext(ctx)
}









type NodeStateCountsResponsePtrInput interface {
	pulumi.Input

	ToNodeStateCountsResponsePtrOutput() NodeStateCountsResponsePtrOutput
	ToNodeStateCountsResponsePtrOutputWithContext(context.Context) NodeStateCountsResponsePtrOutput
}

type nodeStateCountsResponsePtrType NodeStateCountsResponseArgs

func NodeStateCountsResponsePtr(v *NodeStateCountsResponseArgs) NodeStateCountsResponsePtrInput {
	return (*nodeStateCountsResponsePtrType)(v)
}

func (*nodeStateCountsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeStateCountsResponse)(nil)).Elem()
}

func (i *nodeStateCountsResponsePtrType) ToNodeStateCountsResponsePtrOutput() NodeStateCountsResponsePtrOutput {
	return i.ToNodeStateCountsResponsePtrOutputWithContext(context.Background())
}

func (i *nodeStateCountsResponsePtrType) ToNodeStateCountsResponsePtrOutputWithContext(ctx context.Context) NodeStateCountsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeStateCountsResponsePtrOutput)
}

type NodeStateCountsResponseOutput struct{ *pulumi.OutputState }

func (NodeStateCountsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeStateCountsResponse)(nil)).Elem()
}

func (o NodeStateCountsResponseOutput) ToNodeStateCountsResponseOutput() NodeStateCountsResponseOutput {
	return o
}

func (o NodeStateCountsResponseOutput) ToNodeStateCountsResponseOutputWithContext(ctx context.Context) NodeStateCountsResponseOutput {
	return o
}

func (o NodeStateCountsResponseOutput) ToNodeStateCountsResponsePtrOutput() NodeStateCountsResponsePtrOutput {
	return o.ToNodeStateCountsResponsePtrOutputWithContext(context.Background())
}

func (o NodeStateCountsResponseOutput) ToNodeStateCountsResponsePtrOutputWithContext(ctx context.Context) NodeStateCountsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NodeStateCountsResponse) *NodeStateCountsResponse {
		return &v
	}).(NodeStateCountsResponsePtrOutput)
}

func (o NodeStateCountsResponseOutput) IdleNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.IdleNodeCount }).(pulumi.IntOutput)
}

func (o NodeStateCountsResponseOutput) LeavingNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.LeavingNodeCount }).(pulumi.IntOutput)
}

func (o NodeStateCountsResponseOutput) PreemptedNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.PreemptedNodeCount }).(pulumi.IntOutput)
}

func (o NodeStateCountsResponseOutput) PreparingNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.PreparingNodeCount }).(pulumi.IntOutput)
}

func (o NodeStateCountsResponseOutput) RunningNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.RunningNodeCount }).(pulumi.IntOutput)
}

func (o NodeStateCountsResponseOutput) UnusableNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.UnusableNodeCount }).(pulumi.IntOutput)
}

type NodeStateCountsResponsePtrOutput struct{ *pulumi.OutputState }

func (NodeStateCountsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeStateCountsResponse)(nil)).Elem()
}

func (o NodeStateCountsResponsePtrOutput) ToNodeStateCountsResponsePtrOutput() NodeStateCountsResponsePtrOutput {
	return o
}

func (o NodeStateCountsResponsePtrOutput) ToNodeStateCountsResponsePtrOutputWithContext(ctx context.Context) NodeStateCountsResponsePtrOutput {
	return o
}

func (o NodeStateCountsResponsePtrOutput) Elem() NodeStateCountsResponseOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) NodeStateCountsResponse {
		if v != nil {
			return *v
		}
		var ret NodeStateCountsResponse
		return ret
	}).(NodeStateCountsResponseOutput)
}

func (o NodeStateCountsResponsePtrOutput) IdleNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.IdleNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o NodeStateCountsResponsePtrOutput) LeavingNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.LeavingNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o NodeStateCountsResponsePtrOutput) PreemptedNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.PreemptedNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o NodeStateCountsResponsePtrOutput) PreparingNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.PreparingNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o NodeStateCountsResponsePtrOutput) RunningNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.RunningNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o NodeStateCountsResponsePtrOutput) UnusableNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.UnusableNodeCount
	}).(pulumi.IntPtrOutput)
}

type PasswordResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type PasswordResponseInput interface {
	pulumi.Input

	ToPasswordResponseOutput() PasswordResponseOutput
	ToPasswordResponseOutputWithContext(context.Context) PasswordResponseOutput
}

type PasswordResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (PasswordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PasswordResponse)(nil)).Elem()
}

func (i PasswordResponseArgs) ToPasswordResponseOutput() PasswordResponseOutput {
	return i.ToPasswordResponseOutputWithContext(context.Background())
}

func (i PasswordResponseArgs) ToPasswordResponseOutputWithContext(ctx context.Context) PasswordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordResponseOutput)
}





type PasswordResponseArrayInput interface {
	pulumi.Input

	ToPasswordResponseArrayOutput() PasswordResponseArrayOutput
	ToPasswordResponseArrayOutputWithContext(context.Context) PasswordResponseArrayOutput
}

type PasswordResponseArray []PasswordResponseInput

func (PasswordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PasswordResponse)(nil)).Elem()
}

func (i PasswordResponseArray) ToPasswordResponseArrayOutput() PasswordResponseArrayOutput {
	return i.ToPasswordResponseArrayOutputWithContext(context.Background())
}

func (i PasswordResponseArray) ToPasswordResponseArrayOutputWithContext(ctx context.Context) PasswordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordResponseArrayOutput)
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





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                             `pulumi:"id"`
	Identity                          IdentityResponsePtrInput                       `pulumi:"identity"`
	Location                          pulumi.StringPtrInput                          `pulumi:"location"`
	Name                              pulumi.StringInput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                             `pulumi:"provisioningState"`
	Sku                               SkuResponsePtrInput                            `pulumi:"sku"`
	Tags                              pulumi.StringMapInput                          `pulumi:"tags"`
	Type                              pulumi.StringInput                             `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
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





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
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

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
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

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
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

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
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

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
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

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
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





type RCranPackageResponseInput interface {
	pulumi.Input

	ToRCranPackageResponseOutput() RCranPackageResponseOutput
	ToRCranPackageResponseOutputWithContext(context.Context) RCranPackageResponseOutput
}

type RCranPackageResponseArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Repository pulumi.StringPtrInput `pulumi:"repository"`
}

func (RCranPackageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RCranPackageResponse)(nil)).Elem()
}

func (i RCranPackageResponseArgs) ToRCranPackageResponseOutput() RCranPackageResponseOutput {
	return i.ToRCranPackageResponseOutputWithContext(context.Background())
}

func (i RCranPackageResponseArgs) ToRCranPackageResponseOutputWithContext(ctx context.Context) RCranPackageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RCranPackageResponseOutput)
}





type RCranPackageResponseArrayInput interface {
	pulumi.Input

	ToRCranPackageResponseArrayOutput() RCranPackageResponseArrayOutput
	ToRCranPackageResponseArrayOutputWithContext(context.Context) RCranPackageResponseArrayOutput
}

type RCranPackageResponseArray []RCranPackageResponseInput

func (RCranPackageResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RCranPackageResponse)(nil)).Elem()
}

func (i RCranPackageResponseArray) ToRCranPackageResponseArrayOutput() RCranPackageResponseArrayOutput {
	return i.ToRCranPackageResponseArrayOutputWithContext(context.Background())
}

func (i RCranPackageResponseArray) ToRCranPackageResponseArrayOutputWithContext(ctx context.Context) RCranPackageResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RCranPackageResponseArrayOutput)
}

type RCranPackageResponseOutput struct{ *pulumi.OutputState }

func (RCranPackageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RCranPackageResponse)(nil)).Elem()
}

func (o RCranPackageResponseOutput) ToRCranPackageResponseOutput() RCranPackageResponseOutput {
	return o
}

func (o RCranPackageResponseOutput) ToRCranPackageResponseOutputWithContext(ctx context.Context) RCranPackageResponseOutput {
	return o
}

func (o RCranPackageResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RCranPackageResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RCranPackageResponseOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RCranPackageResponse) *string { return v.Repository }).(pulumi.StringPtrOutput)
}

type RCranPackageResponseArrayOutput struct{ *pulumi.OutputState }

func (RCranPackageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RCranPackageResponse)(nil)).Elem()
}

func (o RCranPackageResponseArrayOutput) ToRCranPackageResponseArrayOutput() RCranPackageResponseArrayOutput {
	return o
}

func (o RCranPackageResponseArrayOutput) ToRCranPackageResponseArrayOutputWithContext(ctx context.Context) RCranPackageResponseArrayOutput {
	return o
}

func (o RCranPackageResponseArrayOutput) Index(i pulumi.IntInput) RCranPackageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RCranPackageResponse {
		return vs[0].([]RCranPackageResponse)[vs[1].(int)]
	}).(RCranPackageResponseOutput)
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





type RGitHubPackageResponseResponseInput interface {
	pulumi.Input

	ToRGitHubPackageResponseResponseOutput() RGitHubPackageResponseResponseOutput
	ToRGitHubPackageResponseResponseOutputWithContext(context.Context) RGitHubPackageResponseResponseOutput
}

type RGitHubPackageResponseResponseArgs struct {
	Repository pulumi.StringPtrInput `pulumi:"repository"`
}

func (RGitHubPackageResponseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RGitHubPackageResponseResponse)(nil)).Elem()
}

func (i RGitHubPackageResponseResponseArgs) ToRGitHubPackageResponseResponseOutput() RGitHubPackageResponseResponseOutput {
	return i.ToRGitHubPackageResponseResponseOutputWithContext(context.Background())
}

func (i RGitHubPackageResponseResponseArgs) ToRGitHubPackageResponseResponseOutputWithContext(ctx context.Context) RGitHubPackageResponseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RGitHubPackageResponseResponseOutput)
}





type RGitHubPackageResponseResponseArrayInput interface {
	pulumi.Input

	ToRGitHubPackageResponseResponseArrayOutput() RGitHubPackageResponseResponseArrayOutput
	ToRGitHubPackageResponseResponseArrayOutputWithContext(context.Context) RGitHubPackageResponseResponseArrayOutput
}

type RGitHubPackageResponseResponseArray []RGitHubPackageResponseResponseInput

func (RGitHubPackageResponseResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RGitHubPackageResponseResponse)(nil)).Elem()
}

func (i RGitHubPackageResponseResponseArray) ToRGitHubPackageResponseResponseArrayOutput() RGitHubPackageResponseResponseArrayOutput {
	return i.ToRGitHubPackageResponseResponseArrayOutputWithContext(context.Background())
}

func (i RGitHubPackageResponseResponseArray) ToRGitHubPackageResponseResponseArrayOutputWithContext(ctx context.Context) RGitHubPackageResponseResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RGitHubPackageResponseResponseArrayOutput)
}

type RGitHubPackageResponseResponseOutput struct{ *pulumi.OutputState }

func (RGitHubPackageResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RGitHubPackageResponseResponse)(nil)).Elem()
}

func (o RGitHubPackageResponseResponseOutput) ToRGitHubPackageResponseResponseOutput() RGitHubPackageResponseResponseOutput {
	return o
}

func (o RGitHubPackageResponseResponseOutput) ToRGitHubPackageResponseResponseOutputWithContext(ctx context.Context) RGitHubPackageResponseResponseOutput {
	return o
}

func (o RGitHubPackageResponseResponseOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RGitHubPackageResponseResponse) *string { return v.Repository }).(pulumi.StringPtrOutput)
}

type RGitHubPackageResponseResponseArrayOutput struct{ *pulumi.OutputState }

func (RGitHubPackageResponseResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RGitHubPackageResponseResponse)(nil)).Elem()
}

func (o RGitHubPackageResponseResponseArrayOutput) ToRGitHubPackageResponseResponseArrayOutput() RGitHubPackageResponseResponseArrayOutput {
	return o
}

func (o RGitHubPackageResponseResponseArrayOutput) ToRGitHubPackageResponseResponseArrayOutputWithContext(ctx context.Context) RGitHubPackageResponseResponseArrayOutput {
	return o
}

func (o RGitHubPackageResponseResponseArrayOutput) Index(i pulumi.IntInput) RGitHubPackageResponseResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RGitHubPackageResponseResponse {
		return vs[0].([]RGitHubPackageResponseResponse)[vs[1].(int)]
	}).(RGitHubPackageResponseResponseOutput)
}

type RegistryListCredentialsResultResponse struct {
	Location  string             `pulumi:"location"`
	Passwords []PasswordResponse `pulumi:"passwords"`
	Username  string             `pulumi:"username"`
}





type RegistryListCredentialsResultResponseInput interface {
	pulumi.Input

	ToRegistryListCredentialsResultResponseOutput() RegistryListCredentialsResultResponseOutput
	ToRegistryListCredentialsResultResponseOutputWithContext(context.Context) RegistryListCredentialsResultResponseOutput
}

type RegistryListCredentialsResultResponseArgs struct {
	Location  pulumi.StringInput         `pulumi:"location"`
	Passwords PasswordResponseArrayInput `pulumi:"passwords"`
	Username  pulumi.StringInput         `pulumi:"username"`
}

func (RegistryListCredentialsResultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryListCredentialsResultResponse)(nil)).Elem()
}

func (i RegistryListCredentialsResultResponseArgs) ToRegistryListCredentialsResultResponseOutput() RegistryListCredentialsResultResponseOutput {
	return i.ToRegistryListCredentialsResultResponseOutputWithContext(context.Background())
}

func (i RegistryListCredentialsResultResponseArgs) ToRegistryListCredentialsResultResponseOutputWithContext(ctx context.Context) RegistryListCredentialsResultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryListCredentialsResultResponseOutput)
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

type ResourceId struct {
	Id string `pulumi:"id"`
}





type ResourceIdInput interface {
	pulumi.Input

	ToResourceIdOutput() ResourceIdOutput
	ToResourceIdOutputWithContext(context.Context) ResourceIdOutput
}

type ResourceIdArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (ResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceId)(nil)).Elem()
}

func (i ResourceIdArgs) ToResourceIdOutput() ResourceIdOutput {
	return i.ToResourceIdOutputWithContext(context.Background())
}

func (i ResourceIdArgs) ToResourceIdOutputWithContext(ctx context.Context) ResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdOutput)
}

func (i ResourceIdArgs) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return i.ToResourceIdPtrOutputWithContext(context.Background())
}

func (i ResourceIdArgs) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdOutput).ToResourceIdPtrOutputWithContext(ctx)
}









type ResourceIdPtrInput interface {
	pulumi.Input

	ToResourceIdPtrOutput() ResourceIdPtrOutput
	ToResourceIdPtrOutputWithContext(context.Context) ResourceIdPtrOutput
}

type resourceIdPtrType ResourceIdArgs

func ResourceIdPtr(v *ResourceIdArgs) ResourceIdPtrInput {
	return (*resourceIdPtrType)(v)
}

func (*resourceIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceId)(nil)).Elem()
}

func (i *resourceIdPtrType) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return i.ToResourceIdPtrOutputWithContext(context.Background())
}

func (i *resourceIdPtrType) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdPtrOutput)
}

type ResourceIdOutput struct{ *pulumi.OutputState }

func (ResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceId)(nil)).Elem()
}

func (o ResourceIdOutput) ToResourceIdOutput() ResourceIdOutput {
	return o
}

func (o ResourceIdOutput) ToResourceIdOutputWithContext(ctx context.Context) ResourceIdOutput {
	return o
}

func (o ResourceIdOutput) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return o.ToResourceIdPtrOutputWithContext(context.Background())
}

func (o ResourceIdOutput) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceId) *ResourceId {
		return &v
	}).(ResourceIdPtrOutput)
}

func (o ResourceIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceId) string { return v.Id }).(pulumi.StringOutput)
}

type ResourceIdPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceId)(nil)).Elem()
}

func (o ResourceIdPtrOutput) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return o
}

func (o ResourceIdPtrOutput) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return o
}

func (o ResourceIdPtrOutput) Elem() ResourceIdOutput {
	return o.ApplyT(func(v *ResourceId) ResourceId {
		if v != nil {
			return *v
		}
		var ret ResourceId
		return ret
	}).(ResourceIdOutput)
}

func (o ResourceIdPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceId) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourceIdResponse struct {
	Id string `pulumi:"id"`
}





type ResourceIdResponseInput interface {
	pulumi.Input

	ToResourceIdResponseOutput() ResourceIdResponseOutput
	ToResourceIdResponseOutputWithContext(context.Context) ResourceIdResponseOutput
}

type ResourceIdResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (ResourceIdResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdResponse)(nil)).Elem()
}

func (i ResourceIdResponseArgs) ToResourceIdResponseOutput() ResourceIdResponseOutput {
	return i.ToResourceIdResponseOutputWithContext(context.Background())
}

func (i ResourceIdResponseArgs) ToResourceIdResponseOutputWithContext(ctx context.Context) ResourceIdResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdResponseOutput)
}

func (i ResourceIdResponseArgs) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return i.ToResourceIdResponsePtrOutputWithContext(context.Background())
}

func (i ResourceIdResponseArgs) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdResponseOutput).ToResourceIdResponsePtrOutputWithContext(ctx)
}









type ResourceIdResponsePtrInput interface {
	pulumi.Input

	ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput
	ToResourceIdResponsePtrOutputWithContext(context.Context) ResourceIdResponsePtrOutput
}

type resourceIdResponsePtrType ResourceIdResponseArgs

func ResourceIdResponsePtr(v *ResourceIdResponseArgs) ResourceIdResponsePtrInput {
	return (*resourceIdResponsePtrType)(v)
}

func (*resourceIdResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdResponse)(nil)).Elem()
}

func (i *resourceIdResponsePtrType) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return i.ToResourceIdResponsePtrOutputWithContext(context.Background())
}

func (i *resourceIdResponsePtrType) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdResponsePtrOutput)
}

type ResourceIdResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdResponse)(nil)).Elem()
}

func (o ResourceIdResponseOutput) ToResourceIdResponseOutput() ResourceIdResponseOutput {
	return o
}

func (o ResourceIdResponseOutput) ToResourceIdResponseOutputWithContext(ctx context.Context) ResourceIdResponseOutput {
	return o
}

func (o ResourceIdResponseOutput) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return o.ToResourceIdResponsePtrOutputWithContext(context.Background())
}

func (o ResourceIdResponseOutput) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdResponse) *ResourceIdResponse {
		return &v
	}).(ResourceIdResponsePtrOutput)
}

func (o ResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ResourceIdResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdResponse)(nil)).Elem()
}

func (o ResourceIdResponsePtrOutput) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return o
}

func (o ResourceIdResponsePtrOutput) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return o
}

func (o ResourceIdResponsePtrOutput) Elem() ResourceIdResponseOutput {
	return o.ApplyT(func(v *ResourceIdResponse) ResourceIdResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdResponse
		return ret
	}).(ResourceIdResponseOutput)
}

func (o ResourceIdResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ScaleSettings struct {
	MaxNodeCount                int     `pulumi:"maxNodeCount"`
	MinNodeCount                *int    `pulumi:"minNodeCount"`
	NodeIdleTimeBeforeScaleDown *string `pulumi:"nodeIdleTimeBeforeScaleDown"`
}





type ScaleSettingsInput interface {
	pulumi.Input

	ToScaleSettingsOutput() ScaleSettingsOutput
	ToScaleSettingsOutputWithContext(context.Context) ScaleSettingsOutput
}

type ScaleSettingsArgs struct {
	MaxNodeCount                pulumi.IntInput       `pulumi:"maxNodeCount"`
	MinNodeCount                pulumi.IntPtrInput    `pulumi:"minNodeCount"`
	NodeIdleTimeBeforeScaleDown pulumi.StringPtrInput `pulumi:"nodeIdleTimeBeforeScaleDown"`
}

func (ScaleSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSettings)(nil)).Elem()
}

func (i ScaleSettingsArgs) ToScaleSettingsOutput() ScaleSettingsOutput {
	return i.ToScaleSettingsOutputWithContext(context.Background())
}

func (i ScaleSettingsArgs) ToScaleSettingsOutputWithContext(ctx context.Context) ScaleSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsOutput)
}

func (i ScaleSettingsArgs) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return i.ToScaleSettingsPtrOutputWithContext(context.Background())
}

func (i ScaleSettingsArgs) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsOutput).ToScaleSettingsPtrOutputWithContext(ctx)
}









type ScaleSettingsPtrInput interface {
	pulumi.Input

	ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput
	ToScaleSettingsPtrOutputWithContext(context.Context) ScaleSettingsPtrOutput
}

type scaleSettingsPtrType ScaleSettingsArgs

func ScaleSettingsPtr(v *ScaleSettingsArgs) ScaleSettingsPtrInput {
	return (*scaleSettingsPtrType)(v)
}

func (*scaleSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSettings)(nil)).Elem()
}

func (i *scaleSettingsPtrType) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return i.ToScaleSettingsPtrOutputWithContext(context.Background())
}

func (i *scaleSettingsPtrType) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsPtrOutput)
}

type ScaleSettingsOutput struct{ *pulumi.OutputState }

func (ScaleSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSettings)(nil)).Elem()
}

func (o ScaleSettingsOutput) ToScaleSettingsOutput() ScaleSettingsOutput {
	return o
}

func (o ScaleSettingsOutput) ToScaleSettingsOutputWithContext(ctx context.Context) ScaleSettingsOutput {
	return o
}

func (o ScaleSettingsOutput) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return o.ToScaleSettingsPtrOutputWithContext(context.Background())
}

func (o ScaleSettingsOutput) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScaleSettings) *ScaleSettings {
		return &v
	}).(ScaleSettingsPtrOutput)
}

func (o ScaleSettingsOutput) MaxNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v ScaleSettings) int { return v.MaxNodeCount }).(pulumi.IntOutput)
}

func (o ScaleSettingsOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScaleSettings) *int { return v.MinNodeCount }).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsOutput) NodeIdleTimeBeforeScaleDown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleSettings) *string { return v.NodeIdleTimeBeforeScaleDown }).(pulumi.StringPtrOutput)
}

type ScaleSettingsPtrOutput struct{ *pulumi.OutputState }

func (ScaleSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSettings)(nil)).Elem()
}

func (o ScaleSettingsPtrOutput) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return o
}

func (o ScaleSettingsPtrOutput) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return o
}

func (o ScaleSettingsPtrOutput) Elem() ScaleSettingsOutput {
	return o.ApplyT(func(v *ScaleSettings) ScaleSettings {
		if v != nil {
			return *v
		}
		var ret ScaleSettings
		return ret
	}).(ScaleSettingsOutput)
}

func (o ScaleSettingsPtrOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleSettings) *int {
		if v == nil {
			return nil
		}
		return &v.MaxNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsPtrOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleSettings) *int {
		if v == nil {
			return nil
		}
		return v.MinNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsPtrOutput) NodeIdleTimeBeforeScaleDown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScaleSettings) *string {
		if v == nil {
			return nil
		}
		return v.NodeIdleTimeBeforeScaleDown
	}).(pulumi.StringPtrOutput)
}

type ScaleSettingsResponse struct {
	MaxNodeCount                int     `pulumi:"maxNodeCount"`
	MinNodeCount                *int    `pulumi:"minNodeCount"`
	NodeIdleTimeBeforeScaleDown *string `pulumi:"nodeIdleTimeBeforeScaleDown"`
}





type ScaleSettingsResponseInput interface {
	pulumi.Input

	ToScaleSettingsResponseOutput() ScaleSettingsResponseOutput
	ToScaleSettingsResponseOutputWithContext(context.Context) ScaleSettingsResponseOutput
}

type ScaleSettingsResponseArgs struct {
	MaxNodeCount                pulumi.IntInput       `pulumi:"maxNodeCount"`
	MinNodeCount                pulumi.IntPtrInput    `pulumi:"minNodeCount"`
	NodeIdleTimeBeforeScaleDown pulumi.StringPtrInput `pulumi:"nodeIdleTimeBeforeScaleDown"`
}

func (ScaleSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSettingsResponse)(nil)).Elem()
}

func (i ScaleSettingsResponseArgs) ToScaleSettingsResponseOutput() ScaleSettingsResponseOutput {
	return i.ToScaleSettingsResponseOutputWithContext(context.Background())
}

func (i ScaleSettingsResponseArgs) ToScaleSettingsResponseOutputWithContext(ctx context.Context) ScaleSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsResponseOutput)
}

func (i ScaleSettingsResponseArgs) ToScaleSettingsResponsePtrOutput() ScaleSettingsResponsePtrOutput {
	return i.ToScaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ScaleSettingsResponseArgs) ToScaleSettingsResponsePtrOutputWithContext(ctx context.Context) ScaleSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsResponseOutput).ToScaleSettingsResponsePtrOutputWithContext(ctx)
}









type ScaleSettingsResponsePtrInput interface {
	pulumi.Input

	ToScaleSettingsResponsePtrOutput() ScaleSettingsResponsePtrOutput
	ToScaleSettingsResponsePtrOutputWithContext(context.Context) ScaleSettingsResponsePtrOutput
}

type scaleSettingsResponsePtrType ScaleSettingsResponseArgs

func ScaleSettingsResponsePtr(v *ScaleSettingsResponseArgs) ScaleSettingsResponsePtrInput {
	return (*scaleSettingsResponsePtrType)(v)
}

func (*scaleSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSettingsResponse)(nil)).Elem()
}

func (i *scaleSettingsResponsePtrType) ToScaleSettingsResponsePtrOutput() ScaleSettingsResponsePtrOutput {
	return i.ToScaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *scaleSettingsResponsePtrType) ToScaleSettingsResponsePtrOutputWithContext(ctx context.Context) ScaleSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsResponsePtrOutput)
}

type ScaleSettingsResponseOutput struct{ *pulumi.OutputState }

func (ScaleSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSettingsResponse)(nil)).Elem()
}

func (o ScaleSettingsResponseOutput) ToScaleSettingsResponseOutput() ScaleSettingsResponseOutput {
	return o
}

func (o ScaleSettingsResponseOutput) ToScaleSettingsResponseOutputWithContext(ctx context.Context) ScaleSettingsResponseOutput {
	return o
}

func (o ScaleSettingsResponseOutput) ToScaleSettingsResponsePtrOutput() ScaleSettingsResponsePtrOutput {
	return o.ToScaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ScaleSettingsResponseOutput) ToScaleSettingsResponsePtrOutputWithContext(ctx context.Context) ScaleSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScaleSettingsResponse) *ScaleSettingsResponse {
		return &v
	}).(ScaleSettingsResponsePtrOutput)
}

func (o ScaleSettingsResponseOutput) MaxNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v ScaleSettingsResponse) int { return v.MaxNodeCount }).(pulumi.IntOutput)
}

func (o ScaleSettingsResponseOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScaleSettingsResponse) *int { return v.MinNodeCount }).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsResponseOutput) NodeIdleTimeBeforeScaleDown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleSettingsResponse) *string { return v.NodeIdleTimeBeforeScaleDown }).(pulumi.StringPtrOutput)
}

type ScaleSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ScaleSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSettingsResponse)(nil)).Elem()
}

func (o ScaleSettingsResponsePtrOutput) ToScaleSettingsResponsePtrOutput() ScaleSettingsResponsePtrOutput {
	return o
}

func (o ScaleSettingsResponsePtrOutput) ToScaleSettingsResponsePtrOutputWithContext(ctx context.Context) ScaleSettingsResponsePtrOutput {
	return o
}

func (o ScaleSettingsResponsePtrOutput) Elem() ScaleSettingsResponseOutput {
	return o.ApplyT(func(v *ScaleSettingsResponse) ScaleSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ScaleSettingsResponse
		return ret
	}).(ScaleSettingsResponseOutput)
}

func (o ScaleSettingsResponsePtrOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsResponsePtrOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsResponsePtrOutput) NodeIdleTimeBeforeScaleDown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScaleSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.NodeIdleTimeBeforeScaleDown
	}).(pulumi.StringPtrOutput)
}

type ServiceResponseBaseResponseError struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
}





type ServiceResponseBaseResponseErrorInput interface {
	pulumi.Input

	ToServiceResponseBaseResponseErrorOutput() ServiceResponseBaseResponseErrorOutput
	ToServiceResponseBaseResponseErrorOutputWithContext(context.Context) ServiceResponseBaseResponseErrorOutput
}

type ServiceResponseBaseResponseErrorArgs struct {
	Code    pulumi.StringInput            `pulumi:"code"`
	Details ErrorDetailResponseArrayInput `pulumi:"details"`
	Message pulumi.StringInput            `pulumi:"message"`
}

func (ServiceResponseBaseResponseErrorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResponseBaseResponseError)(nil)).Elem()
}

func (i ServiceResponseBaseResponseErrorArgs) ToServiceResponseBaseResponseErrorOutput() ServiceResponseBaseResponseErrorOutput {
	return i.ToServiceResponseBaseResponseErrorOutputWithContext(context.Background())
}

func (i ServiceResponseBaseResponseErrorArgs) ToServiceResponseBaseResponseErrorOutputWithContext(ctx context.Context) ServiceResponseBaseResponseErrorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceResponseBaseResponseErrorOutput)
}

type ServiceResponseBaseResponseErrorOutput struct{ *pulumi.OutputState }

func (ServiceResponseBaseResponseErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResponseBaseResponseError)(nil)).Elem()
}

func (o ServiceResponseBaseResponseErrorOutput) ToServiceResponseBaseResponseErrorOutput() ServiceResponseBaseResponseErrorOutput {
	return o
}

func (o ServiceResponseBaseResponseErrorOutput) ToServiceResponseBaseResponseErrorOutputWithContext(ctx context.Context) ServiceResponseBaseResponseErrorOutput {
	return o
}

func (o ServiceResponseBaseResponseErrorOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResponseBaseResponseError) string { return v.Code }).(pulumi.StringOutput)
}

func (o ServiceResponseBaseResponseErrorOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ServiceResponseBaseResponseError) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ServiceResponseBaseResponseErrorOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResponseBaseResponseError) string { return v.Message }).(pulumi.StringOutput)
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





type SharedPrivateLinkResourceResponseInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput
	ToSharedPrivateLinkResourceResponseOutputWithContext(context.Context) SharedPrivateLinkResourceResponseOutput
}

type SharedPrivateLinkResourceResponseArgs struct {
	GroupId               pulumi.StringPtrInput `pulumi:"groupId"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	PrivateLinkResourceId pulumi.StringPtrInput `pulumi:"privateLinkResourceId"`
	RequestMessage        pulumi.StringPtrInput `pulumi:"requestMessage"`
	Status                pulumi.StringPtrInput `pulumi:"status"`
}

func (SharedPrivateLinkResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (i SharedPrivateLinkResourceResponseArgs) ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput {
	return i.ToSharedPrivateLinkResourceResponseOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceResponseArgs) ToSharedPrivateLinkResourceResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceResponseOutput)
}





type SharedPrivateLinkResourceResponseArrayInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput
	ToSharedPrivateLinkResourceResponseArrayOutputWithContext(context.Context) SharedPrivateLinkResourceResponseArrayOutput
}

type SharedPrivateLinkResourceResponseArray []SharedPrivateLinkResourceResponseInput

func (SharedPrivateLinkResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (i SharedPrivateLinkResourceResponseArray) ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput {
	return i.ToSharedPrivateLinkResourceResponseArrayOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceResponseArray) ToSharedPrivateLinkResourceResponseArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceResponseArrayOutput)
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





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
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

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
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





type SparkMavenPackageResponseInput interface {
	pulumi.Input

	ToSparkMavenPackageResponseOutput() SparkMavenPackageResponseOutput
	ToSparkMavenPackageResponseOutputWithContext(context.Context) SparkMavenPackageResponseOutput
}

type SparkMavenPackageResponseArgs struct {
	Artifact pulumi.StringPtrInput `pulumi:"artifact"`
	Group    pulumi.StringPtrInput `pulumi:"group"`
	Version  pulumi.StringPtrInput `pulumi:"version"`
}

func (SparkMavenPackageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkMavenPackageResponse)(nil)).Elem()
}

func (i SparkMavenPackageResponseArgs) ToSparkMavenPackageResponseOutput() SparkMavenPackageResponseOutput {
	return i.ToSparkMavenPackageResponseOutputWithContext(context.Background())
}

func (i SparkMavenPackageResponseArgs) ToSparkMavenPackageResponseOutputWithContext(ctx context.Context) SparkMavenPackageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkMavenPackageResponseOutput)
}





type SparkMavenPackageResponseArrayInput interface {
	pulumi.Input

	ToSparkMavenPackageResponseArrayOutput() SparkMavenPackageResponseArrayOutput
	ToSparkMavenPackageResponseArrayOutputWithContext(context.Context) SparkMavenPackageResponseArrayOutput
}

type SparkMavenPackageResponseArray []SparkMavenPackageResponseInput

func (SparkMavenPackageResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SparkMavenPackageResponse)(nil)).Elem()
}

func (i SparkMavenPackageResponseArray) ToSparkMavenPackageResponseArrayOutput() SparkMavenPackageResponseArrayOutput {
	return i.ToSparkMavenPackageResponseArrayOutputWithContext(context.Background())
}

func (i SparkMavenPackageResponseArray) ToSparkMavenPackageResponseArrayOutputWithContext(ctx context.Context) SparkMavenPackageResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SparkMavenPackageResponseArrayOutput)
}

type SparkMavenPackageResponseOutput struct{ *pulumi.OutputState }

func (SparkMavenPackageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkMavenPackageResponse)(nil)).Elem()
}

func (o SparkMavenPackageResponseOutput) ToSparkMavenPackageResponseOutput() SparkMavenPackageResponseOutput {
	return o
}

func (o SparkMavenPackageResponseOutput) ToSparkMavenPackageResponseOutputWithContext(ctx context.Context) SparkMavenPackageResponseOutput {
	return o
}

func (o SparkMavenPackageResponseOutput) Artifact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkMavenPackageResponse) *string { return v.Artifact }).(pulumi.StringPtrOutput)
}

func (o SparkMavenPackageResponseOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkMavenPackageResponse) *string { return v.Group }).(pulumi.StringPtrOutput)
}

func (o SparkMavenPackageResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SparkMavenPackageResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type SparkMavenPackageResponseArrayOutput struct{ *pulumi.OutputState }

func (SparkMavenPackageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SparkMavenPackageResponse)(nil)).Elem()
}

func (o SparkMavenPackageResponseArrayOutput) ToSparkMavenPackageResponseArrayOutput() SparkMavenPackageResponseArrayOutput {
	return o
}

func (o SparkMavenPackageResponseArrayOutput) ToSparkMavenPackageResponseArrayOutputWithContext(ctx context.Context) SparkMavenPackageResponseArrayOutput {
	return o
}

func (o SparkMavenPackageResponseArrayOutput) Index(i pulumi.IntInput) SparkMavenPackageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SparkMavenPackageResponse {
		return vs[0].([]SparkMavenPackageResponse)[vs[1].(int)]
	}).(SparkMavenPackageResponseOutput)
}

type SslConfiguration struct {
	Cert   *string `pulumi:"cert"`
	Cname  *string `pulumi:"cname"`
	Key    *string `pulumi:"key"`
	Status *string `pulumi:"status"`
}





type SslConfigurationInput interface {
	pulumi.Input

	ToSslConfigurationOutput() SslConfigurationOutput
	ToSslConfigurationOutputWithContext(context.Context) SslConfigurationOutput
}

type SslConfigurationArgs struct {
	Cert   pulumi.StringPtrInput `pulumi:"cert"`
	Cname  pulumi.StringPtrInput `pulumi:"cname"`
	Key    pulumi.StringPtrInput `pulumi:"key"`
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (SslConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfiguration)(nil)).Elem()
}

func (i SslConfigurationArgs) ToSslConfigurationOutput() SslConfigurationOutput {
	return i.ToSslConfigurationOutputWithContext(context.Background())
}

func (i SslConfigurationArgs) ToSslConfigurationOutputWithContext(ctx context.Context) SslConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationOutput)
}

func (i SslConfigurationArgs) ToSslConfigurationPtrOutput() SslConfigurationPtrOutput {
	return i.ToSslConfigurationPtrOutputWithContext(context.Background())
}

func (i SslConfigurationArgs) ToSslConfigurationPtrOutputWithContext(ctx context.Context) SslConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationOutput).ToSslConfigurationPtrOutputWithContext(ctx)
}









type SslConfigurationPtrInput interface {
	pulumi.Input

	ToSslConfigurationPtrOutput() SslConfigurationPtrOutput
	ToSslConfigurationPtrOutputWithContext(context.Context) SslConfigurationPtrOutput
}

type sslConfigurationPtrType SslConfigurationArgs

func SslConfigurationPtr(v *SslConfigurationArgs) SslConfigurationPtrInput {
	return (*sslConfigurationPtrType)(v)
}

func (*sslConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfiguration)(nil)).Elem()
}

func (i *sslConfigurationPtrType) ToSslConfigurationPtrOutput() SslConfigurationPtrOutput {
	return i.ToSslConfigurationPtrOutputWithContext(context.Background())
}

func (i *sslConfigurationPtrType) ToSslConfigurationPtrOutputWithContext(ctx context.Context) SslConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationPtrOutput)
}

type SslConfigurationOutput struct{ *pulumi.OutputState }

func (SslConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfiguration)(nil)).Elem()
}

func (o SslConfigurationOutput) ToSslConfigurationOutput() SslConfigurationOutput {
	return o
}

func (o SslConfigurationOutput) ToSslConfigurationOutputWithContext(ctx context.Context) SslConfigurationOutput {
	return o
}

func (o SslConfigurationOutput) ToSslConfigurationPtrOutput() SslConfigurationPtrOutput {
	return o.ToSslConfigurationPtrOutputWithContext(context.Background())
}

func (o SslConfigurationOutput) ToSslConfigurationPtrOutputWithContext(ctx context.Context) SslConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslConfiguration) *SslConfiguration {
		return &v
	}).(SslConfigurationPtrOutput)
}

func (o SslConfigurationOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfiguration) *string { return v.Cert }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfiguration) *string { return v.Cname }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfiguration) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfiguration) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SslConfigurationPtrOutput struct{ *pulumi.OutputState }

func (SslConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfiguration)(nil)).Elem()
}

func (o SslConfigurationPtrOutput) ToSslConfigurationPtrOutput() SslConfigurationPtrOutput {
	return o
}

func (o SslConfigurationPtrOutput) ToSslConfigurationPtrOutputWithContext(ctx context.Context) SslConfigurationPtrOutput {
	return o
}

func (o SslConfigurationPtrOutput) Elem() SslConfigurationOutput {
	return o.ApplyT(func(v *SslConfiguration) SslConfiguration {
		if v != nil {
			return *v
		}
		var ret SslConfiguration
		return ret
	}).(SslConfigurationOutput)
}

func (o SslConfigurationPtrOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Cert
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationPtrOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Cname
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type SslConfigurationResponse struct {
	Cert   *string `pulumi:"cert"`
	Cname  *string `pulumi:"cname"`
	Key    *string `pulumi:"key"`
	Status *string `pulumi:"status"`
}





type SslConfigurationResponseInput interface {
	pulumi.Input

	ToSslConfigurationResponseOutput() SslConfigurationResponseOutput
	ToSslConfigurationResponseOutputWithContext(context.Context) SslConfigurationResponseOutput
}

type SslConfigurationResponseArgs struct {
	Cert   pulumi.StringPtrInput `pulumi:"cert"`
	Cname  pulumi.StringPtrInput `pulumi:"cname"`
	Key    pulumi.StringPtrInput `pulumi:"key"`
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (SslConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigurationResponse)(nil)).Elem()
}

func (i SslConfigurationResponseArgs) ToSslConfigurationResponseOutput() SslConfigurationResponseOutput {
	return i.ToSslConfigurationResponseOutputWithContext(context.Background())
}

func (i SslConfigurationResponseArgs) ToSslConfigurationResponseOutputWithContext(ctx context.Context) SslConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationResponseOutput)
}

func (i SslConfigurationResponseArgs) ToSslConfigurationResponsePtrOutput() SslConfigurationResponsePtrOutput {
	return i.ToSslConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i SslConfigurationResponseArgs) ToSslConfigurationResponsePtrOutputWithContext(ctx context.Context) SslConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationResponseOutput).ToSslConfigurationResponsePtrOutputWithContext(ctx)
}









type SslConfigurationResponsePtrInput interface {
	pulumi.Input

	ToSslConfigurationResponsePtrOutput() SslConfigurationResponsePtrOutput
	ToSslConfigurationResponsePtrOutputWithContext(context.Context) SslConfigurationResponsePtrOutput
}

type sslConfigurationResponsePtrType SslConfigurationResponseArgs

func SslConfigurationResponsePtr(v *SslConfigurationResponseArgs) SslConfigurationResponsePtrInput {
	return (*sslConfigurationResponsePtrType)(v)
}

func (*sslConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfigurationResponse)(nil)).Elem()
}

func (i *sslConfigurationResponsePtrType) ToSslConfigurationResponsePtrOutput() SslConfigurationResponsePtrOutput {
	return i.ToSslConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *sslConfigurationResponsePtrType) ToSslConfigurationResponsePtrOutputWithContext(ctx context.Context) SslConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationResponsePtrOutput)
}

type SslConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SslConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigurationResponse)(nil)).Elem()
}

func (o SslConfigurationResponseOutput) ToSslConfigurationResponseOutput() SslConfigurationResponseOutput {
	return o
}

func (o SslConfigurationResponseOutput) ToSslConfigurationResponseOutputWithContext(ctx context.Context) SslConfigurationResponseOutput {
	return o
}

func (o SslConfigurationResponseOutput) ToSslConfigurationResponsePtrOutput() SslConfigurationResponsePtrOutput {
	return o.ToSslConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o SslConfigurationResponseOutput) ToSslConfigurationResponsePtrOutputWithContext(ctx context.Context) SslConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslConfigurationResponse) *SslConfigurationResponse {
		return &v
	}).(SslConfigurationResponsePtrOutput)
}

func (o SslConfigurationResponseOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfigurationResponse) *string { return v.Cert }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponseOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfigurationResponse) *string { return v.Cname }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfigurationResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfigurationResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SslConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (SslConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfigurationResponse)(nil)).Elem()
}

func (o SslConfigurationResponsePtrOutput) ToSslConfigurationResponsePtrOutput() SslConfigurationResponsePtrOutput {
	return o
}

func (o SslConfigurationResponsePtrOutput) ToSslConfigurationResponsePtrOutputWithContext(ctx context.Context) SslConfigurationResponsePtrOutput {
	return o
}

func (o SslConfigurationResponsePtrOutput) Elem() SslConfigurationResponseOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) SslConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret SslConfigurationResponse
		return ret
	}).(SslConfigurationResponseOutput)
}

func (o SslConfigurationResponsePtrOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Cert
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponsePtrOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Cname
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponsePtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type SystemServiceResponse struct {
	PublicIpAddress   string `pulumi:"publicIpAddress"`
	SystemServiceType string `pulumi:"systemServiceType"`
	Version           string `pulumi:"version"`
}





type SystemServiceResponseInput interface {
	pulumi.Input

	ToSystemServiceResponseOutput() SystemServiceResponseOutput
	ToSystemServiceResponseOutputWithContext(context.Context) SystemServiceResponseOutput
}

type SystemServiceResponseArgs struct {
	PublicIpAddress   pulumi.StringInput `pulumi:"publicIpAddress"`
	SystemServiceType pulumi.StringInput `pulumi:"systemServiceType"`
	Version           pulumi.StringInput `pulumi:"version"`
}

func (SystemServiceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemServiceResponse)(nil)).Elem()
}

func (i SystemServiceResponseArgs) ToSystemServiceResponseOutput() SystemServiceResponseOutput {
	return i.ToSystemServiceResponseOutputWithContext(context.Background())
}

func (i SystemServiceResponseArgs) ToSystemServiceResponseOutputWithContext(ctx context.Context) SystemServiceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemServiceResponseOutput)
}





type SystemServiceResponseArrayInput interface {
	pulumi.Input

	ToSystemServiceResponseArrayOutput() SystemServiceResponseArrayOutput
	ToSystemServiceResponseArrayOutputWithContext(context.Context) SystemServiceResponseArrayOutput
}

type SystemServiceResponseArray []SystemServiceResponseInput

func (SystemServiceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemServiceResponse)(nil)).Elem()
}

func (i SystemServiceResponseArray) ToSystemServiceResponseArrayOutput() SystemServiceResponseArrayOutput {
	return i.ToSystemServiceResponseArrayOutputWithContext(context.Background())
}

func (i SystemServiceResponseArray) ToSystemServiceResponseArrayOutputWithContext(ctx context.Context) SystemServiceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemServiceResponseArrayOutput)
}

type SystemServiceResponseOutput struct{ *pulumi.OutputState }

func (SystemServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemServiceResponse)(nil)).Elem()
}

func (o SystemServiceResponseOutput) ToSystemServiceResponseOutput() SystemServiceResponseOutput {
	return o
}

func (o SystemServiceResponseOutput) ToSystemServiceResponseOutputWithContext(ctx context.Context) SystemServiceResponseOutput {
	return o
}

func (o SystemServiceResponseOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v SystemServiceResponse) string { return v.PublicIpAddress }).(pulumi.StringOutput)
}

func (o SystemServiceResponseOutput) SystemServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemServiceResponse) string { return v.SystemServiceType }).(pulumi.StringOutput)
}

func (o SystemServiceResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v SystemServiceResponse) string { return v.Version }).(pulumi.StringOutput)
}

type SystemServiceResponseArrayOutput struct{ *pulumi.OutputState }

func (SystemServiceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemServiceResponse)(nil)).Elem()
}

func (o SystemServiceResponseArrayOutput) ToSystemServiceResponseArrayOutput() SystemServiceResponseArrayOutput {
	return o
}

func (o SystemServiceResponseArrayOutput) ToSystemServiceResponseArrayOutputWithContext(ctx context.Context) SystemServiceResponseArrayOutput {
	return o
}

func (o SystemServiceResponseArrayOutput) Index(i pulumi.IntInput) SystemServiceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemServiceResponse {
		return vs[0].([]SystemServiceResponse)[vs[1].(int)]
	}).(SystemServiceResponseOutput)
}

type UserAccountCredentials struct {
	AdminUserName         string  `pulumi:"adminUserName"`
	AdminUserPassword     *string `pulumi:"adminUserPassword"`
	AdminUserSshPublicKey *string `pulumi:"adminUserSshPublicKey"`
}





type UserAccountCredentialsInput interface {
	pulumi.Input

	ToUserAccountCredentialsOutput() UserAccountCredentialsOutput
	ToUserAccountCredentialsOutputWithContext(context.Context) UserAccountCredentialsOutput
}

type UserAccountCredentialsArgs struct {
	AdminUserName         pulumi.StringInput    `pulumi:"adminUserName"`
	AdminUserPassword     pulumi.StringPtrInput `pulumi:"adminUserPassword"`
	AdminUserSshPublicKey pulumi.StringPtrInput `pulumi:"adminUserSshPublicKey"`
}

func (UserAccountCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccountCredentials)(nil)).Elem()
}

func (i UserAccountCredentialsArgs) ToUserAccountCredentialsOutput() UserAccountCredentialsOutput {
	return i.ToUserAccountCredentialsOutputWithContext(context.Background())
}

func (i UserAccountCredentialsArgs) ToUserAccountCredentialsOutputWithContext(ctx context.Context) UserAccountCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsOutput)
}

func (i UserAccountCredentialsArgs) ToUserAccountCredentialsPtrOutput() UserAccountCredentialsPtrOutput {
	return i.ToUserAccountCredentialsPtrOutputWithContext(context.Background())
}

func (i UserAccountCredentialsArgs) ToUserAccountCredentialsPtrOutputWithContext(ctx context.Context) UserAccountCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsOutput).ToUserAccountCredentialsPtrOutputWithContext(ctx)
}









type UserAccountCredentialsPtrInput interface {
	pulumi.Input

	ToUserAccountCredentialsPtrOutput() UserAccountCredentialsPtrOutput
	ToUserAccountCredentialsPtrOutputWithContext(context.Context) UserAccountCredentialsPtrOutput
}

type userAccountCredentialsPtrType UserAccountCredentialsArgs

func UserAccountCredentialsPtr(v *UserAccountCredentialsArgs) UserAccountCredentialsPtrInput {
	return (*userAccountCredentialsPtrType)(v)
}

func (*userAccountCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAccountCredentials)(nil)).Elem()
}

func (i *userAccountCredentialsPtrType) ToUserAccountCredentialsPtrOutput() UserAccountCredentialsPtrOutput {
	return i.ToUserAccountCredentialsPtrOutputWithContext(context.Background())
}

func (i *userAccountCredentialsPtrType) ToUserAccountCredentialsPtrOutputWithContext(ctx context.Context) UserAccountCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsPtrOutput)
}

type UserAccountCredentialsOutput struct{ *pulumi.OutputState }

func (UserAccountCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccountCredentials)(nil)).Elem()
}

func (o UserAccountCredentialsOutput) ToUserAccountCredentialsOutput() UserAccountCredentialsOutput {
	return o
}

func (o UserAccountCredentialsOutput) ToUserAccountCredentialsOutputWithContext(ctx context.Context) UserAccountCredentialsOutput {
	return o
}

func (o UserAccountCredentialsOutput) ToUserAccountCredentialsPtrOutput() UserAccountCredentialsPtrOutput {
	return o.ToUserAccountCredentialsPtrOutputWithContext(context.Background())
}

func (o UserAccountCredentialsOutput) ToUserAccountCredentialsPtrOutputWithContext(ctx context.Context) UserAccountCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserAccountCredentials) *UserAccountCredentials {
		return &v
	}).(UserAccountCredentialsPtrOutput)
}

func (o UserAccountCredentialsOutput) AdminUserName() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccountCredentials) string { return v.AdminUserName }).(pulumi.StringOutput)
}

func (o UserAccountCredentialsOutput) AdminUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccountCredentials) *string { return v.AdminUserPassword }).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsOutput) AdminUserSshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccountCredentials) *string { return v.AdminUserSshPublicKey }).(pulumi.StringPtrOutput)
}

type UserAccountCredentialsPtrOutput struct{ *pulumi.OutputState }

func (UserAccountCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAccountCredentials)(nil)).Elem()
}

func (o UserAccountCredentialsPtrOutput) ToUserAccountCredentialsPtrOutput() UserAccountCredentialsPtrOutput {
	return o
}

func (o UserAccountCredentialsPtrOutput) ToUserAccountCredentialsPtrOutputWithContext(ctx context.Context) UserAccountCredentialsPtrOutput {
	return o
}

func (o UserAccountCredentialsPtrOutput) Elem() UserAccountCredentialsOutput {
	return o.ApplyT(func(v *UserAccountCredentials) UserAccountCredentials {
		if v != nil {
			return *v
		}
		var ret UserAccountCredentials
		return ret
	}).(UserAccountCredentialsOutput)
}

func (o UserAccountCredentialsPtrOutput) AdminUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUserName
	}).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsPtrOutput) AdminUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentials) *string {
		if v == nil {
			return nil
		}
		return v.AdminUserPassword
	}).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsPtrOutput) AdminUserSshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentials) *string {
		if v == nil {
			return nil
		}
		return v.AdminUserSshPublicKey
	}).(pulumi.StringPtrOutput)
}

type UserAccountCredentialsResponse struct {
	AdminUserName         string  `pulumi:"adminUserName"`
	AdminUserPassword     *string `pulumi:"adminUserPassword"`
	AdminUserSshPublicKey *string `pulumi:"adminUserSshPublicKey"`
}





type UserAccountCredentialsResponseInput interface {
	pulumi.Input

	ToUserAccountCredentialsResponseOutput() UserAccountCredentialsResponseOutput
	ToUserAccountCredentialsResponseOutputWithContext(context.Context) UserAccountCredentialsResponseOutput
}

type UserAccountCredentialsResponseArgs struct {
	AdminUserName         pulumi.StringInput    `pulumi:"adminUserName"`
	AdminUserPassword     pulumi.StringPtrInput `pulumi:"adminUserPassword"`
	AdminUserSshPublicKey pulumi.StringPtrInput `pulumi:"adminUserSshPublicKey"`
}

func (UserAccountCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccountCredentialsResponse)(nil)).Elem()
}

func (i UserAccountCredentialsResponseArgs) ToUserAccountCredentialsResponseOutput() UserAccountCredentialsResponseOutput {
	return i.ToUserAccountCredentialsResponseOutputWithContext(context.Background())
}

func (i UserAccountCredentialsResponseArgs) ToUserAccountCredentialsResponseOutputWithContext(ctx context.Context) UserAccountCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsResponseOutput)
}

func (i UserAccountCredentialsResponseArgs) ToUserAccountCredentialsResponsePtrOutput() UserAccountCredentialsResponsePtrOutput {
	return i.ToUserAccountCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i UserAccountCredentialsResponseArgs) ToUserAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) UserAccountCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsResponseOutput).ToUserAccountCredentialsResponsePtrOutputWithContext(ctx)
}









type UserAccountCredentialsResponsePtrInput interface {
	pulumi.Input

	ToUserAccountCredentialsResponsePtrOutput() UserAccountCredentialsResponsePtrOutput
	ToUserAccountCredentialsResponsePtrOutputWithContext(context.Context) UserAccountCredentialsResponsePtrOutput
}

type userAccountCredentialsResponsePtrType UserAccountCredentialsResponseArgs

func UserAccountCredentialsResponsePtr(v *UserAccountCredentialsResponseArgs) UserAccountCredentialsResponsePtrInput {
	return (*userAccountCredentialsResponsePtrType)(v)
}

func (*userAccountCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAccountCredentialsResponse)(nil)).Elem()
}

func (i *userAccountCredentialsResponsePtrType) ToUserAccountCredentialsResponsePtrOutput() UserAccountCredentialsResponsePtrOutput {
	return i.ToUserAccountCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *userAccountCredentialsResponsePtrType) ToUserAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) UserAccountCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsResponsePtrOutput)
}

type UserAccountCredentialsResponseOutput struct{ *pulumi.OutputState }

func (UserAccountCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccountCredentialsResponse)(nil)).Elem()
}

func (o UserAccountCredentialsResponseOutput) ToUserAccountCredentialsResponseOutput() UserAccountCredentialsResponseOutput {
	return o
}

func (o UserAccountCredentialsResponseOutput) ToUserAccountCredentialsResponseOutputWithContext(ctx context.Context) UserAccountCredentialsResponseOutput {
	return o
}

func (o UserAccountCredentialsResponseOutput) ToUserAccountCredentialsResponsePtrOutput() UserAccountCredentialsResponsePtrOutput {
	return o.ToUserAccountCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o UserAccountCredentialsResponseOutput) ToUserAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) UserAccountCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserAccountCredentialsResponse) *UserAccountCredentialsResponse {
		return &v
	}).(UserAccountCredentialsResponsePtrOutput)
}

func (o UserAccountCredentialsResponseOutput) AdminUserName() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccountCredentialsResponse) string { return v.AdminUserName }).(pulumi.StringOutput)
}

func (o UserAccountCredentialsResponseOutput) AdminUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccountCredentialsResponse) *string { return v.AdminUserPassword }).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsResponseOutput) AdminUserSshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccountCredentialsResponse) *string { return v.AdminUserSshPublicKey }).(pulumi.StringPtrOutput)
}

type UserAccountCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (UserAccountCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAccountCredentialsResponse)(nil)).Elem()
}

func (o UserAccountCredentialsResponsePtrOutput) ToUserAccountCredentialsResponsePtrOutput() UserAccountCredentialsResponsePtrOutput {
	return o
}

func (o UserAccountCredentialsResponsePtrOutput) ToUserAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) UserAccountCredentialsResponsePtrOutput {
	return o
}

func (o UserAccountCredentialsResponsePtrOutput) Elem() UserAccountCredentialsResponseOutput {
	return o.ApplyT(func(v *UserAccountCredentialsResponse) UserAccountCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret UserAccountCredentialsResponse
		return ret
	}).(UserAccountCredentialsResponseOutput)
}

func (o UserAccountCredentialsResponsePtrOutput) AdminUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUserName
	}).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsResponsePtrOutput) AdminUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUserPassword
	}).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsResponsePtrOutput) AdminUserSshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUserSshPublicKey
	}).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
}





type UserAssignedIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput
	ToUserAssignedIdentityResponseOutputWithContext(context.Context) UserAssignedIdentityResponseOutput
}

type UserAssignedIdentityResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
}

func (UserAssignedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return i.ToUserAssignedIdentityResponseOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseOutput)
}





type UserAssignedIdentityResponseMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput
	ToUserAssignedIdentityResponseMapOutputWithContext(context.Context) UserAssignedIdentityResponseMapOutput
}

type UserAssignedIdentityResponseMap map[string]UserAssignedIdentityResponseInput

func (UserAssignedIdentityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return i.ToUserAssignedIdentityResponseMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseMapOutput)
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





type VirtualMachineInput interface {
	pulumi.Input

	ToVirtualMachineOutput() VirtualMachineOutput
	ToVirtualMachineOutputWithContext(context.Context) VirtualMachineOutput
}

type VirtualMachineArgs struct {
	ComputeLocation pulumi.StringPtrInput            `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput               `pulumi:"computeType"`
	Description     pulumi.StringPtrInput            `pulumi:"description"`
	Properties      VirtualMachinePropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput            `pulumi:"resourceId"`
}

func (VirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachine)(nil)).Elem()
}

func (i VirtualMachineArgs) ToVirtualMachineOutput() VirtualMachineOutput {
	return i.ToVirtualMachineOutputWithContext(context.Background())
}

func (i VirtualMachineArgs) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineOutput)
}

type VirtualMachineOutput struct{ *pulumi.OutputState }

func (VirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachine)(nil)).Elem()
}

func (o VirtualMachineOutput) ToVirtualMachineOutput() VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachine) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachine) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachine) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Properties() VirtualMachinePropertiesPtrOutput {
	return o.ApplyT(func(v VirtualMachine) *VirtualMachineProperties { return v.Properties }).(VirtualMachinePropertiesPtrOutput)
}

func (o VirtualMachineOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachine) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type VirtualMachineImage struct {
	Id string `pulumi:"id"`
}





type VirtualMachineImageInput interface {
	pulumi.Input

	ToVirtualMachineImageOutput() VirtualMachineImageOutput
	ToVirtualMachineImageOutputWithContext(context.Context) VirtualMachineImageOutput
}

type VirtualMachineImageArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (VirtualMachineImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineImage)(nil)).Elem()
}

func (i VirtualMachineImageArgs) ToVirtualMachineImageOutput() VirtualMachineImageOutput {
	return i.ToVirtualMachineImageOutputWithContext(context.Background())
}

func (i VirtualMachineImageArgs) ToVirtualMachineImageOutputWithContext(ctx context.Context) VirtualMachineImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageOutput)
}

func (i VirtualMachineImageArgs) ToVirtualMachineImagePtrOutput() VirtualMachineImagePtrOutput {
	return i.ToVirtualMachineImagePtrOutputWithContext(context.Background())
}

func (i VirtualMachineImageArgs) ToVirtualMachineImagePtrOutputWithContext(ctx context.Context) VirtualMachineImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageOutput).ToVirtualMachineImagePtrOutputWithContext(ctx)
}









type VirtualMachineImagePtrInput interface {
	pulumi.Input

	ToVirtualMachineImagePtrOutput() VirtualMachineImagePtrOutput
	ToVirtualMachineImagePtrOutputWithContext(context.Context) VirtualMachineImagePtrOutput
}

type virtualMachineImagePtrType VirtualMachineImageArgs

func VirtualMachineImagePtr(v *VirtualMachineImageArgs) VirtualMachineImagePtrInput {
	return (*virtualMachineImagePtrType)(v)
}

func (*virtualMachineImagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImage)(nil)).Elem()
}

func (i *virtualMachineImagePtrType) ToVirtualMachineImagePtrOutput() VirtualMachineImagePtrOutput {
	return i.ToVirtualMachineImagePtrOutputWithContext(context.Background())
}

func (i *virtualMachineImagePtrType) ToVirtualMachineImagePtrOutputWithContext(ctx context.Context) VirtualMachineImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImagePtrOutput)
}

type VirtualMachineImageOutput struct{ *pulumi.OutputState }

func (VirtualMachineImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineImage)(nil)).Elem()
}

func (o VirtualMachineImageOutput) ToVirtualMachineImageOutput() VirtualMachineImageOutput {
	return o
}

func (o VirtualMachineImageOutput) ToVirtualMachineImageOutputWithContext(ctx context.Context) VirtualMachineImageOutput {
	return o
}

func (o VirtualMachineImageOutput) ToVirtualMachineImagePtrOutput() VirtualMachineImagePtrOutput {
	return o.ToVirtualMachineImagePtrOutputWithContext(context.Background())
}

func (o VirtualMachineImageOutput) ToVirtualMachineImagePtrOutputWithContext(ctx context.Context) VirtualMachineImagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineImage) *VirtualMachineImage {
		return &v
	}).(VirtualMachineImagePtrOutput)
}

func (o VirtualMachineImageOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineImage) string { return v.Id }).(pulumi.StringOutput)
}

type VirtualMachineImagePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineImagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImage)(nil)).Elem()
}

func (o VirtualMachineImagePtrOutput) ToVirtualMachineImagePtrOutput() VirtualMachineImagePtrOutput {
	return o
}

func (o VirtualMachineImagePtrOutput) ToVirtualMachineImagePtrOutputWithContext(ctx context.Context) VirtualMachineImagePtrOutput {
	return o
}

func (o VirtualMachineImagePtrOutput) Elem() VirtualMachineImageOutput {
	return o.ApplyT(func(v *VirtualMachineImage) VirtualMachineImage {
		if v != nil {
			return *v
		}
		var ret VirtualMachineImage
		return ret
	}).(VirtualMachineImageOutput)
}

func (o VirtualMachineImagePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineImage) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineImageResponse struct {
	Id string `pulumi:"id"`
}





type VirtualMachineImageResponseInput interface {
	pulumi.Input

	ToVirtualMachineImageResponseOutput() VirtualMachineImageResponseOutput
	ToVirtualMachineImageResponseOutputWithContext(context.Context) VirtualMachineImageResponseOutput
}

type VirtualMachineImageResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (VirtualMachineImageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineImageResponse)(nil)).Elem()
}

func (i VirtualMachineImageResponseArgs) ToVirtualMachineImageResponseOutput() VirtualMachineImageResponseOutput {
	return i.ToVirtualMachineImageResponseOutputWithContext(context.Background())
}

func (i VirtualMachineImageResponseArgs) ToVirtualMachineImageResponseOutputWithContext(ctx context.Context) VirtualMachineImageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageResponseOutput)
}

func (i VirtualMachineImageResponseArgs) ToVirtualMachineImageResponsePtrOutput() VirtualMachineImageResponsePtrOutput {
	return i.ToVirtualMachineImageResponsePtrOutputWithContext(context.Background())
}

func (i VirtualMachineImageResponseArgs) ToVirtualMachineImageResponsePtrOutputWithContext(ctx context.Context) VirtualMachineImageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageResponseOutput).ToVirtualMachineImageResponsePtrOutputWithContext(ctx)
}









type VirtualMachineImageResponsePtrInput interface {
	pulumi.Input

	ToVirtualMachineImageResponsePtrOutput() VirtualMachineImageResponsePtrOutput
	ToVirtualMachineImageResponsePtrOutputWithContext(context.Context) VirtualMachineImageResponsePtrOutput
}

type virtualMachineImageResponsePtrType VirtualMachineImageResponseArgs

func VirtualMachineImageResponsePtr(v *VirtualMachineImageResponseArgs) VirtualMachineImageResponsePtrInput {
	return (*virtualMachineImageResponsePtrType)(v)
}

func (*virtualMachineImageResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImageResponse)(nil)).Elem()
}

func (i *virtualMachineImageResponsePtrType) ToVirtualMachineImageResponsePtrOutput() VirtualMachineImageResponsePtrOutput {
	return i.ToVirtualMachineImageResponsePtrOutputWithContext(context.Background())
}

func (i *virtualMachineImageResponsePtrType) ToVirtualMachineImageResponsePtrOutputWithContext(ctx context.Context) VirtualMachineImageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageResponsePtrOutput)
}

type VirtualMachineImageResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineImageResponse)(nil)).Elem()
}

func (o VirtualMachineImageResponseOutput) ToVirtualMachineImageResponseOutput() VirtualMachineImageResponseOutput {
	return o
}

func (o VirtualMachineImageResponseOutput) ToVirtualMachineImageResponseOutputWithContext(ctx context.Context) VirtualMachineImageResponseOutput {
	return o
}

func (o VirtualMachineImageResponseOutput) ToVirtualMachineImageResponsePtrOutput() VirtualMachineImageResponsePtrOutput {
	return o.ToVirtualMachineImageResponsePtrOutputWithContext(context.Background())
}

func (o VirtualMachineImageResponseOutput) ToVirtualMachineImageResponsePtrOutputWithContext(ctx context.Context) VirtualMachineImageResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineImageResponse) *VirtualMachineImageResponse {
		return &v
	}).(VirtualMachineImageResponsePtrOutput)
}

func (o VirtualMachineImageResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineImageResponse) string { return v.Id }).(pulumi.StringOutput)
}

type VirtualMachineImageResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineImageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImageResponse)(nil)).Elem()
}

func (o VirtualMachineImageResponsePtrOutput) ToVirtualMachineImageResponsePtrOutput() VirtualMachineImageResponsePtrOutput {
	return o
}

func (o VirtualMachineImageResponsePtrOutput) ToVirtualMachineImageResponsePtrOutputWithContext(ctx context.Context) VirtualMachineImageResponsePtrOutput {
	return o
}

func (o VirtualMachineImageResponsePtrOutput) Elem() VirtualMachineImageResponseOutput {
	return o.ApplyT(func(v *VirtualMachineImageResponse) VirtualMachineImageResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineImageResponse
		return ret
	}).(VirtualMachineImageResponseOutput)
}

func (o VirtualMachineImageResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineImageResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineProperties struct {
	Address              *string                       `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	SshPort              *int                          `pulumi:"sshPort"`
	VirtualMachineSize   *string                       `pulumi:"virtualMachineSize"`
}





type VirtualMachinePropertiesInput interface {
	pulumi.Input

	ToVirtualMachinePropertiesOutput() VirtualMachinePropertiesOutput
	ToVirtualMachinePropertiesOutputWithContext(context.Context) VirtualMachinePropertiesOutput
}

type VirtualMachinePropertiesArgs struct {
	Address              pulumi.StringPtrInput                `pulumi:"address"`
	AdministratorAccount VirtualMachineSshCredentialsPtrInput `pulumi:"administratorAccount"`
	SshPort              pulumi.IntPtrInput                   `pulumi:"sshPort"`
	VirtualMachineSize   pulumi.StringPtrInput                `pulumi:"virtualMachineSize"`
}

func (VirtualMachinePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineProperties)(nil)).Elem()
}

func (i VirtualMachinePropertiesArgs) ToVirtualMachinePropertiesOutput() VirtualMachinePropertiesOutput {
	return i.ToVirtualMachinePropertiesOutputWithContext(context.Background())
}

func (i VirtualMachinePropertiesArgs) ToVirtualMachinePropertiesOutputWithContext(ctx context.Context) VirtualMachinePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePropertiesOutput)
}

func (i VirtualMachinePropertiesArgs) ToVirtualMachinePropertiesPtrOutput() VirtualMachinePropertiesPtrOutput {
	return i.ToVirtualMachinePropertiesPtrOutputWithContext(context.Background())
}

func (i VirtualMachinePropertiesArgs) ToVirtualMachinePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachinePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePropertiesOutput).ToVirtualMachinePropertiesPtrOutputWithContext(ctx)
}









type VirtualMachinePropertiesPtrInput interface {
	pulumi.Input

	ToVirtualMachinePropertiesPtrOutput() VirtualMachinePropertiesPtrOutput
	ToVirtualMachinePropertiesPtrOutputWithContext(context.Context) VirtualMachinePropertiesPtrOutput
}

type virtualMachinePropertiesPtrType VirtualMachinePropertiesArgs

func VirtualMachinePropertiesPtr(v *VirtualMachinePropertiesArgs) VirtualMachinePropertiesPtrInput {
	return (*virtualMachinePropertiesPtrType)(v)
}

func (*virtualMachinePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineProperties)(nil)).Elem()
}

func (i *virtualMachinePropertiesPtrType) ToVirtualMachinePropertiesPtrOutput() VirtualMachinePropertiesPtrOutput {
	return i.ToVirtualMachinePropertiesPtrOutputWithContext(context.Background())
}

func (i *virtualMachinePropertiesPtrType) ToVirtualMachinePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachinePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePropertiesPtrOutput)
}

type VirtualMachinePropertiesOutput struct{ *pulumi.OutputState }

func (VirtualMachinePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineProperties)(nil)).Elem()
}

func (o VirtualMachinePropertiesOutput) ToVirtualMachinePropertiesOutput() VirtualMachinePropertiesOutput {
	return o
}

func (o VirtualMachinePropertiesOutput) ToVirtualMachinePropertiesOutputWithContext(ctx context.Context) VirtualMachinePropertiesOutput {
	return o
}

func (o VirtualMachinePropertiesOutput) ToVirtualMachinePropertiesPtrOutput() VirtualMachinePropertiesPtrOutput {
	return o.ToVirtualMachinePropertiesPtrOutputWithContext(context.Background())
}

func (o VirtualMachinePropertiesOutput) ToVirtualMachinePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachinePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineProperties) *VirtualMachineProperties {
		return &v
	}).(VirtualMachinePropertiesPtrOutput)
}

func (o VirtualMachinePropertiesOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineProperties) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePropertiesOutput) AdministratorAccount() VirtualMachineSshCredentialsPtrOutput {
	return o.ApplyT(func(v VirtualMachineProperties) *VirtualMachineSshCredentials { return v.AdministratorAccount }).(VirtualMachineSshCredentialsPtrOutput)
}

func (o VirtualMachinePropertiesOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineProperties) *int { return v.SshPort }).(pulumi.IntPtrOutput)
}

func (o VirtualMachinePropertiesOutput) VirtualMachineSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineProperties) *string { return v.VirtualMachineSize }).(pulumi.StringPtrOutput)
}

type VirtualMachinePropertiesPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachinePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineProperties)(nil)).Elem()
}

func (o VirtualMachinePropertiesPtrOutput) ToVirtualMachinePropertiesPtrOutput() VirtualMachinePropertiesPtrOutput {
	return o
}

func (o VirtualMachinePropertiesPtrOutput) ToVirtualMachinePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachinePropertiesPtrOutput {
	return o
}

func (o VirtualMachinePropertiesPtrOutput) Elem() VirtualMachinePropertiesOutput {
	return o.ApplyT(func(v *VirtualMachineProperties) VirtualMachineProperties {
		if v != nil {
			return *v
		}
		var ret VirtualMachineProperties
		return ret
	}).(VirtualMachinePropertiesOutput)
}

func (o VirtualMachinePropertiesPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProperties) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePropertiesPtrOutput) AdministratorAccount() VirtualMachineSshCredentialsPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProperties) *VirtualMachineSshCredentials {
		if v == nil {
			return nil
		}
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsPtrOutput)
}

func (o VirtualMachinePropertiesPtrOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProperties) *int {
		if v == nil {
			return nil
		}
		return v.SshPort
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachinePropertiesPtrOutput) VirtualMachineSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProperties) *string {
		if v == nil {
			return nil
		}
		return v.VirtualMachineSize
	}).(pulumi.StringPtrOutput)
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





type VirtualMachineResponseInput interface {
	pulumi.Input

	ToVirtualMachineResponseOutput() VirtualMachineResponseOutput
	ToVirtualMachineResponseOutputWithContext(context.Context) VirtualMachineResponseOutput
}

type VirtualMachineResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         VirtualMachineResponsePropertiesPtrInput      `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (VirtualMachineResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResponse)(nil)).Elem()
}

func (i VirtualMachineResponseArgs) ToVirtualMachineResponseOutput() VirtualMachineResponseOutput {
	return i.ToVirtualMachineResponseOutputWithContext(context.Background())
}

func (i VirtualMachineResponseArgs) ToVirtualMachineResponseOutputWithContext(ctx context.Context) VirtualMachineResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResponseOutput)
}

type VirtualMachineResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResponse)(nil)).Elem()
}

func (o VirtualMachineResponseOutput) ToVirtualMachineResponseOutput() VirtualMachineResponseOutput {
	return o
}

func (o VirtualMachineResponseOutput) ToVirtualMachineResponseOutputWithContext(ctx context.Context) VirtualMachineResponseOutput {
	return o
}

func (o VirtualMachineResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o VirtualMachineResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o VirtualMachineResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v VirtualMachineResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o VirtualMachineResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o VirtualMachineResponseOutput) Properties() VirtualMachineResponsePropertiesPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponse) *VirtualMachineResponseProperties { return v.Properties }).(VirtualMachineResponsePropertiesPtrOutput)
}

func (o VirtualMachineResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o VirtualMachineResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type VirtualMachineResponseProperties struct {
	Address              *string                               `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	SshPort              *int                                  `pulumi:"sshPort"`
	VirtualMachineSize   *string                               `pulumi:"virtualMachineSize"`
}





type VirtualMachineResponsePropertiesInput interface {
	pulumi.Input

	ToVirtualMachineResponsePropertiesOutput() VirtualMachineResponsePropertiesOutput
	ToVirtualMachineResponsePropertiesOutputWithContext(context.Context) VirtualMachineResponsePropertiesOutput
}

type VirtualMachineResponsePropertiesArgs struct {
	Address              pulumi.StringPtrInput                        `pulumi:"address"`
	AdministratorAccount VirtualMachineSshCredentialsResponsePtrInput `pulumi:"administratorAccount"`
	SshPort              pulumi.IntPtrInput                           `pulumi:"sshPort"`
	VirtualMachineSize   pulumi.StringPtrInput                        `pulumi:"virtualMachineSize"`
}

func (VirtualMachineResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResponseProperties)(nil)).Elem()
}

func (i VirtualMachineResponsePropertiesArgs) ToVirtualMachineResponsePropertiesOutput() VirtualMachineResponsePropertiesOutput {
	return i.ToVirtualMachineResponsePropertiesOutputWithContext(context.Background())
}

func (i VirtualMachineResponsePropertiesArgs) ToVirtualMachineResponsePropertiesOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResponsePropertiesOutput)
}

func (i VirtualMachineResponsePropertiesArgs) ToVirtualMachineResponsePropertiesPtrOutput() VirtualMachineResponsePropertiesPtrOutput {
	return i.ToVirtualMachineResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i VirtualMachineResponsePropertiesArgs) ToVirtualMachineResponsePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResponsePropertiesOutput).ToVirtualMachineResponsePropertiesPtrOutputWithContext(ctx)
}









type VirtualMachineResponsePropertiesPtrInput interface {
	pulumi.Input

	ToVirtualMachineResponsePropertiesPtrOutput() VirtualMachineResponsePropertiesPtrOutput
	ToVirtualMachineResponsePropertiesPtrOutputWithContext(context.Context) VirtualMachineResponsePropertiesPtrOutput
}

type virtualMachineResponsePropertiesPtrType VirtualMachineResponsePropertiesArgs

func VirtualMachineResponsePropertiesPtr(v *VirtualMachineResponsePropertiesArgs) VirtualMachineResponsePropertiesPtrInput {
	return (*virtualMachineResponsePropertiesPtrType)(v)
}

func (*virtualMachineResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineResponseProperties)(nil)).Elem()
}

func (i *virtualMachineResponsePropertiesPtrType) ToVirtualMachineResponsePropertiesPtrOutput() VirtualMachineResponsePropertiesPtrOutput {
	return i.ToVirtualMachineResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *virtualMachineResponsePropertiesPtrType) ToVirtualMachineResponsePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResponsePropertiesPtrOutput)
}

type VirtualMachineResponsePropertiesOutput struct{ *pulumi.OutputState }

func (VirtualMachineResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResponseProperties)(nil)).Elem()
}

func (o VirtualMachineResponsePropertiesOutput) ToVirtualMachineResponsePropertiesOutput() VirtualMachineResponsePropertiesOutput {
	return o
}

func (o VirtualMachineResponsePropertiesOutput) ToVirtualMachineResponsePropertiesOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesOutput {
	return o
}

func (o VirtualMachineResponsePropertiesOutput) ToVirtualMachineResponsePropertiesPtrOutput() VirtualMachineResponsePropertiesPtrOutput {
	return o.ToVirtualMachineResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o VirtualMachineResponsePropertiesOutput) ToVirtualMachineResponsePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineResponseProperties) *VirtualMachineResponseProperties {
		return &v
	}).(VirtualMachineResponsePropertiesPtrOutput)
}

func (o VirtualMachineResponsePropertiesOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponseProperties) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResponsePropertiesOutput) AdministratorAccount() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineResponseProperties) *VirtualMachineSshCredentialsResponse {
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o VirtualMachineResponsePropertiesOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponseProperties) *int { return v.SshPort }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineResponsePropertiesOutput) VirtualMachineSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponseProperties) *string { return v.VirtualMachineSize }).(pulumi.StringPtrOutput)
}

type VirtualMachineResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineResponseProperties)(nil)).Elem()
}

func (o VirtualMachineResponsePropertiesPtrOutput) ToVirtualMachineResponsePropertiesPtrOutput() VirtualMachineResponsePropertiesPtrOutput {
	return o
}

func (o VirtualMachineResponsePropertiesPtrOutput) ToVirtualMachineResponsePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesPtrOutput {
	return o
}

func (o VirtualMachineResponsePropertiesPtrOutput) Elem() VirtualMachineResponsePropertiesOutput {
	return o.ApplyT(func(v *VirtualMachineResponseProperties) VirtualMachineResponseProperties {
		if v != nil {
			return *v
		}
		var ret VirtualMachineResponseProperties
		return ret
	}).(VirtualMachineResponsePropertiesOutput)
}

func (o VirtualMachineResponsePropertiesPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResponsePropertiesPtrOutput) AdministratorAccount() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineResponseProperties) *VirtualMachineSshCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o VirtualMachineResponsePropertiesPtrOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResponseProperties) *int {
		if v == nil {
			return nil
		}
		return v.SshPort
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachineResponsePropertiesPtrOutput) VirtualMachineSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.VirtualMachineSize
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineSshCredentials struct {
	Password       *string `pulumi:"password"`
	PrivateKeyData *string `pulumi:"privateKeyData"`
	PublicKeyData  *string `pulumi:"publicKeyData"`
	Username       *string `pulumi:"username"`
}





type VirtualMachineSshCredentialsInput interface {
	pulumi.Input

	ToVirtualMachineSshCredentialsOutput() VirtualMachineSshCredentialsOutput
	ToVirtualMachineSshCredentialsOutputWithContext(context.Context) VirtualMachineSshCredentialsOutput
}

type VirtualMachineSshCredentialsArgs struct {
	Password       pulumi.StringPtrInput `pulumi:"password"`
	PrivateKeyData pulumi.StringPtrInput `pulumi:"privateKeyData"`
	PublicKeyData  pulumi.StringPtrInput `pulumi:"publicKeyData"`
	Username       pulumi.StringPtrInput `pulumi:"username"`
}

func (VirtualMachineSshCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSshCredentials)(nil)).Elem()
}

func (i VirtualMachineSshCredentialsArgs) ToVirtualMachineSshCredentialsOutput() VirtualMachineSshCredentialsOutput {
	return i.ToVirtualMachineSshCredentialsOutputWithContext(context.Background())
}

func (i VirtualMachineSshCredentialsArgs) ToVirtualMachineSshCredentialsOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsOutput)
}

func (i VirtualMachineSshCredentialsArgs) ToVirtualMachineSshCredentialsPtrOutput() VirtualMachineSshCredentialsPtrOutput {
	return i.ToVirtualMachineSshCredentialsPtrOutputWithContext(context.Background())
}

func (i VirtualMachineSshCredentialsArgs) ToVirtualMachineSshCredentialsPtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsOutput).ToVirtualMachineSshCredentialsPtrOutputWithContext(ctx)
}









type VirtualMachineSshCredentialsPtrInput interface {
	pulumi.Input

	ToVirtualMachineSshCredentialsPtrOutput() VirtualMachineSshCredentialsPtrOutput
	ToVirtualMachineSshCredentialsPtrOutputWithContext(context.Context) VirtualMachineSshCredentialsPtrOutput
}

type virtualMachineSshCredentialsPtrType VirtualMachineSshCredentialsArgs

func VirtualMachineSshCredentialsPtr(v *VirtualMachineSshCredentialsArgs) VirtualMachineSshCredentialsPtrInput {
	return (*virtualMachineSshCredentialsPtrType)(v)
}

func (*virtualMachineSshCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSshCredentials)(nil)).Elem()
}

func (i *virtualMachineSshCredentialsPtrType) ToVirtualMachineSshCredentialsPtrOutput() VirtualMachineSshCredentialsPtrOutput {
	return i.ToVirtualMachineSshCredentialsPtrOutputWithContext(context.Background())
}

func (i *virtualMachineSshCredentialsPtrType) ToVirtualMachineSshCredentialsPtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsPtrOutput)
}

type VirtualMachineSshCredentialsOutput struct{ *pulumi.OutputState }

func (VirtualMachineSshCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSshCredentials)(nil)).Elem()
}

func (o VirtualMachineSshCredentialsOutput) ToVirtualMachineSshCredentialsOutput() VirtualMachineSshCredentialsOutput {
	return o
}

func (o VirtualMachineSshCredentialsOutput) ToVirtualMachineSshCredentialsOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsOutput {
	return o
}

func (o VirtualMachineSshCredentialsOutput) ToVirtualMachineSshCredentialsPtrOutput() VirtualMachineSshCredentialsPtrOutput {
	return o.ToVirtualMachineSshCredentialsPtrOutputWithContext(context.Background())
}

func (o VirtualMachineSshCredentialsOutput) ToVirtualMachineSshCredentialsPtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineSshCredentials) *VirtualMachineSshCredentials {
		return &v
	}).(VirtualMachineSshCredentialsPtrOutput)
}

func (o VirtualMachineSshCredentialsOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentials) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsOutput) PrivateKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentials) *string { return v.PrivateKeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsOutput) PublicKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentials) *string { return v.PublicKeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentials) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type VirtualMachineSshCredentialsPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineSshCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSshCredentials)(nil)).Elem()
}

func (o VirtualMachineSshCredentialsPtrOutput) ToVirtualMachineSshCredentialsPtrOutput() VirtualMachineSshCredentialsPtrOutput {
	return o
}

func (o VirtualMachineSshCredentialsPtrOutput) ToVirtualMachineSshCredentialsPtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsPtrOutput {
	return o
}

func (o VirtualMachineSshCredentialsPtrOutput) Elem() VirtualMachineSshCredentialsOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentials) VirtualMachineSshCredentials {
		if v != nil {
			return *v
		}
		var ret VirtualMachineSshCredentials
		return ret
	}).(VirtualMachineSshCredentialsOutput)
}

func (o VirtualMachineSshCredentialsPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentials) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsPtrOutput) PrivateKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentials) *string {
		if v == nil {
			return nil
		}
		return v.PrivateKeyData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsPtrOutput) PublicKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentials) *string {
		if v == nil {
			return nil
		}
		return v.PublicKeyData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentials) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineSshCredentialsResponse struct {
	Password       *string `pulumi:"password"`
	PrivateKeyData *string `pulumi:"privateKeyData"`
	PublicKeyData  *string `pulumi:"publicKeyData"`
	Username       *string `pulumi:"username"`
}





type VirtualMachineSshCredentialsResponseInput interface {
	pulumi.Input

	ToVirtualMachineSshCredentialsResponseOutput() VirtualMachineSshCredentialsResponseOutput
	ToVirtualMachineSshCredentialsResponseOutputWithContext(context.Context) VirtualMachineSshCredentialsResponseOutput
}

type VirtualMachineSshCredentialsResponseArgs struct {
	Password       pulumi.StringPtrInput `pulumi:"password"`
	PrivateKeyData pulumi.StringPtrInput `pulumi:"privateKeyData"`
	PublicKeyData  pulumi.StringPtrInput `pulumi:"publicKeyData"`
	Username       pulumi.StringPtrInput `pulumi:"username"`
}

func (VirtualMachineSshCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSshCredentialsResponse)(nil)).Elem()
}

func (i VirtualMachineSshCredentialsResponseArgs) ToVirtualMachineSshCredentialsResponseOutput() VirtualMachineSshCredentialsResponseOutput {
	return i.ToVirtualMachineSshCredentialsResponseOutputWithContext(context.Background())
}

func (i VirtualMachineSshCredentialsResponseArgs) ToVirtualMachineSshCredentialsResponseOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsResponseOutput)
}

func (i VirtualMachineSshCredentialsResponseArgs) ToVirtualMachineSshCredentialsResponsePtrOutput() VirtualMachineSshCredentialsResponsePtrOutput {
	return i.ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i VirtualMachineSshCredentialsResponseArgs) ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsResponseOutput).ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(ctx)
}









type VirtualMachineSshCredentialsResponsePtrInput interface {
	pulumi.Input

	ToVirtualMachineSshCredentialsResponsePtrOutput() VirtualMachineSshCredentialsResponsePtrOutput
	ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(context.Context) VirtualMachineSshCredentialsResponsePtrOutput
}

type virtualMachineSshCredentialsResponsePtrType VirtualMachineSshCredentialsResponseArgs

func VirtualMachineSshCredentialsResponsePtr(v *VirtualMachineSshCredentialsResponseArgs) VirtualMachineSshCredentialsResponsePtrInput {
	return (*virtualMachineSshCredentialsResponsePtrType)(v)
}

func (*virtualMachineSshCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSshCredentialsResponse)(nil)).Elem()
}

func (i *virtualMachineSshCredentialsResponsePtrType) ToVirtualMachineSshCredentialsResponsePtrOutput() VirtualMachineSshCredentialsResponsePtrOutput {
	return i.ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *virtualMachineSshCredentialsResponsePtrType) ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsResponsePtrOutput)
}

type VirtualMachineSshCredentialsResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineSshCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSshCredentialsResponse)(nil)).Elem()
}

func (o VirtualMachineSshCredentialsResponseOutput) ToVirtualMachineSshCredentialsResponseOutput() VirtualMachineSshCredentialsResponseOutput {
	return o
}

func (o VirtualMachineSshCredentialsResponseOutput) ToVirtualMachineSshCredentialsResponseOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponseOutput {
	return o
}

func (o VirtualMachineSshCredentialsResponseOutput) ToVirtualMachineSshCredentialsResponsePtrOutput() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o VirtualMachineSshCredentialsResponseOutput) ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineSshCredentialsResponse) *VirtualMachineSshCredentialsResponse {
		return &v
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o VirtualMachineSshCredentialsResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentialsResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponseOutput) PrivateKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentialsResponse) *string { return v.PrivateKeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponseOutput) PublicKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentialsResponse) *string { return v.PublicKeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentialsResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type VirtualMachineSshCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineSshCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSshCredentialsResponse)(nil)).Elem()
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) ToVirtualMachineSshCredentialsResponsePtrOutput() VirtualMachineSshCredentialsResponsePtrOutput {
	return o
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponsePtrOutput {
	return o
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) Elem() VirtualMachineSshCredentialsResponseOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentialsResponse) VirtualMachineSshCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineSshCredentialsResponse
		return ret
	}).(VirtualMachineSshCredentialsResponseOutput)
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) PrivateKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateKeyData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) PublicKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicKeyData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceCreateRequestDataCollectionInput)(nil)).Elem(), ACIServiceCreateRequestDataCollectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceCreateRequestDataCollectionPtrInput)(nil)).Elem(), ACIServiceCreateRequestDataCollectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceCreateRequestEncryptionPropertiesInput)(nil)).Elem(), ACIServiceCreateRequestEncryptionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceCreateRequestEncryptionPropertiesPtrInput)(nil)).Elem(), ACIServiceCreateRequestEncryptionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceCreateRequestVnetConfigurationInput)(nil)).Elem(), ACIServiceCreateRequestVnetConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceCreateRequestVnetConfigurationPtrInput)(nil)).Elem(), ACIServiceCreateRequestVnetConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceResponseResponseInput)(nil)).Elem(), ACIServiceResponseResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceResponseResponseDataCollectionInput)(nil)).Elem(), ACIServiceResponseResponseDataCollectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceResponseResponseDataCollectionPtrInput)(nil)).Elem(), ACIServiceResponseResponseDataCollectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceResponseResponseEncryptionPropertiesInput)(nil)).Elem(), ACIServiceResponseResponseEncryptionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceResponseResponseEncryptionPropertiesPtrInput)(nil)).Elem(), ACIServiceResponseResponseEncryptionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceResponseResponseEnvironmentImageRequestInput)(nil)).Elem(), ACIServiceResponseResponseEnvironmentImageRequestArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceResponseResponseEnvironmentImageRequestPtrInput)(nil)).Elem(), ACIServiceResponseResponseEnvironmentImageRequestArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceResponseResponseVnetConfigurationInput)(nil)).Elem(), ACIServiceResponseResponseVnetConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACIServiceResponseResponseVnetConfigurationPtrInput)(nil)).Elem(), ACIServiceResponseResponseVnetConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSInput)(nil)).Elem(), AKSArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSPropertiesInput)(nil)).Elem(), AKSPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSPropertiesPtrInput)(nil)).Elem(), AKSPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSReplicaStatusResponseErrorInput)(nil)).Elem(), AKSReplicaStatusResponseErrorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSReplicaStatusResponseErrorPtrInput)(nil)).Elem(), AKSReplicaStatusResponseErrorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSResponseInput)(nil)).Elem(), AKSResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSResponsePropertiesInput)(nil)).Elem(), AKSResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSResponsePropertiesPtrInput)(nil)).Elem(), AKSResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceCreateRequestAutoScalerInput)(nil)).Elem(), AKSServiceCreateRequestAutoScalerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceCreateRequestAutoScalerPtrInput)(nil)).Elem(), AKSServiceCreateRequestAutoScalerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceCreateRequestDataCollectionInput)(nil)).Elem(), AKSServiceCreateRequestDataCollectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceCreateRequestDataCollectionPtrInput)(nil)).Elem(), AKSServiceCreateRequestDataCollectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceCreateRequestLivenessProbeRequirementsInput)(nil)).Elem(), AKSServiceCreateRequestLivenessProbeRequirementsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceCreateRequestLivenessProbeRequirementsPtrInput)(nil)).Elem(), AKSServiceCreateRequestLivenessProbeRequirementsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceResponseResponseInput)(nil)).Elem(), AKSServiceResponseResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceResponseResponseAutoScalerInput)(nil)).Elem(), AKSServiceResponseResponseAutoScalerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceResponseResponseAutoScalerPtrInput)(nil)).Elem(), AKSServiceResponseResponseAutoScalerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceResponseResponseDataCollectionInput)(nil)).Elem(), AKSServiceResponseResponseDataCollectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceResponseResponseDataCollectionPtrInput)(nil)).Elem(), AKSServiceResponseResponseDataCollectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceResponseResponseDeploymentStatusInput)(nil)).Elem(), AKSServiceResponseResponseDeploymentStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceResponseResponseEnvironmentImageRequestInput)(nil)).Elem(), AKSServiceResponseResponseEnvironmentImageRequestArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceResponseResponseEnvironmentImageRequestPtrInput)(nil)).Elem(), AKSServiceResponseResponseEnvironmentImageRequestArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceResponseResponseLivenessProbeRequirementsInput)(nil)).Elem(), AKSServiceResponseResponseLivenessProbeRequirementsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSServiceResponseResponseLivenessProbeRequirementsPtrInput)(nil)).Elem(), AKSServiceResponseResponseLivenessProbeRequirementsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AKSVariantResponseResponseInput)(nil)).Elem(), AKSVariantResponseResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AksNetworkingConfigurationInput)(nil)).Elem(), AksNetworkingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AksNetworkingConfigurationPtrInput)(nil)).Elem(), AksNetworkingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AksNetworkingConfigurationResponseInput)(nil)).Elem(), AksNetworkingConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AksNetworkingConfigurationResponsePtrInput)(nil)).Elem(), AksNetworkingConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmlComputeInput)(nil)).Elem(), AmlComputeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmlComputeNodeInformationResponseInput)(nil)).Elem(), AmlComputeNodeInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmlComputeNodeInformationResponseArrayInput)(nil)).Elem(), AmlComputeNodeInformationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmlComputePropertiesInput)(nil)).Elem(), AmlComputePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmlComputePropertiesPtrInput)(nil)).Elem(), AmlComputePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmlComputeResponseInput)(nil)).Elem(), AmlComputeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmlComputeResponsePropertiesInput)(nil)).Elem(), AmlComputeResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmlComputeResponsePropertiesPtrInput)(nil)).Elem(), AmlComputeResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerResourceRequirementsInput)(nil)).Elem(), ContainerResourceRequirementsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerResourceRequirementsPtrInput)(nil)).Elem(), ContainerResourceRequirementsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerResourceRequirementsResponseInput)(nil)).Elem(), ContainerResourceRequirementsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerResourceRequirementsResponsePtrInput)(nil)).Elem(), ContainerResourceRequirementsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateServiceRequestEnvironmentImageRequestInput)(nil)).Elem(), CreateServiceRequestEnvironmentImageRequestArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateServiceRequestEnvironmentImageRequestPtrInput)(nil)).Elem(), CreateServiceRequestEnvironmentImageRequestArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateServiceRequestKeysInput)(nil)).Elem(), CreateServiceRequestKeysArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateServiceRequestKeysPtrInput)(nil)).Elem(), CreateServiceRequestKeysArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataFactoryInput)(nil)).Elem(), DataFactoryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataFactoryResponseInput)(nil)).Elem(), DataFactoryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLakeAnalyticsInput)(nil)).Elem(), DataLakeAnalyticsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLakeAnalyticsPropertiesInput)(nil)).Elem(), DataLakeAnalyticsPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLakeAnalyticsPropertiesPtrInput)(nil)).Elem(), DataLakeAnalyticsPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLakeAnalyticsResponseInput)(nil)).Elem(), DataLakeAnalyticsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLakeAnalyticsResponsePropertiesInput)(nil)).Elem(), DataLakeAnalyticsResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLakeAnalyticsResponsePropertiesPtrInput)(nil)).Elem(), DataLakeAnalyticsResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabricksInput)(nil)).Elem(), DatabricksArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabricksPropertiesInput)(nil)).Elem(), DatabricksPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabricksPropertiesPtrInput)(nil)).Elem(), DatabricksPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabricksResponseInput)(nil)).Elem(), DatabricksResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabricksResponsePropertiesInput)(nil)).Elem(), DatabricksResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabricksResponsePropertiesPtrInput)(nil)).Elem(), DatabricksResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetReferenceInput)(nil)).Elem(), DatasetReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetReferenceArrayInput)(nil)).Elem(), DatasetReferenceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetReferenceResponseInput)(nil)).Elem(), DatasetReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetReferenceResponseArrayInput)(nil)).Elem(), DatasetReferenceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPropertyInput)(nil)).Elem(), EncryptionPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPropertyPtrInput)(nil)).Elem(), EncryptionPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPropertyResponseInput)(nil)).Elem(), EncryptionPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPropertyResponsePtrInput)(nil)).Elem(), EncryptionPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentImageRequestEnvironmentInput)(nil)).Elem(), EnvironmentImageRequestEnvironmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentImageRequestEnvironmentPtrInput)(nil)).Elem(), EnvironmentImageRequestEnvironmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentImageRequestEnvironmentReferenceInput)(nil)).Elem(), EnvironmentImageRequestEnvironmentReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentImageRequestEnvironmentReferencePtrInput)(nil)).Elem(), EnvironmentImageRequestEnvironmentReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentImageResponseResponseEnvironmentInput)(nil)).Elem(), EnvironmentImageResponseResponseEnvironmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentImageResponseResponseEnvironmentPtrInput)(nil)).Elem(), EnvironmentImageResponseResponseEnvironmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentImageResponseResponseEnvironmentReferenceInput)(nil)).Elem(), EnvironmentImageResponseResponseEnvironmentReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentImageResponseResponseEnvironmentReferencePtrInput)(nil)).Elem(), EnvironmentImageResponseResponseEnvironmentReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorDetailResponseInput)(nil)).Elem(), ErrorDetailResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorDetailResponseArrayInput)(nil)).Elem(), ErrorDetailResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ErrorResponseResponseInput)(nil)).Elem(), ErrorResponseResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HDInsightInput)(nil)).Elem(), HDInsightArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HDInsightPropertiesInput)(nil)).Elem(), HDInsightPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HDInsightPropertiesPtrInput)(nil)).Elem(), HDInsightPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HDInsightResponseInput)(nil)).Elem(), HDInsightResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HDInsightResponsePropertiesInput)(nil)).Elem(), HDInsightResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HDInsightResponsePropertiesPtrInput)(nil)).Elem(), HDInsightResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPtrInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponseInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponsePtrInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageAssetInput)(nil)).Elem(), ImageAssetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageAssetArrayInput)(nil)).Elem(), ImageAssetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageAssetResponseInput)(nil)).Elem(), ImageAssetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageAssetResponseArrayInput)(nil)).Elem(), ImageAssetResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesInput)(nil)).Elem(), KeyVaultPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesPtrInput)(nil)).Elem(), KeyVaultPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesResponseInput)(nil)).Elem(), KeyVaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesResponsePtrInput)(nil)).Elem(), KeyVaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedWorkspacePropsInput)(nil)).Elem(), LinkedWorkspacePropsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedWorkspacePropsPtrInput)(nil)).Elem(), LinkedWorkspacePropsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedWorkspacePropsResponseInput)(nil)).Elem(), LinkedWorkspacePropsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedWorkspacePropsResponsePtrInput)(nil)).Elem(), LinkedWorkspacePropsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineLearningServiceErrorResponseInput)(nil)).Elem(), MachineLearningServiceErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineLearningServiceErrorResponseArrayInput)(nil)).Elem(), MachineLearningServiceErrorResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelInput)(nil)).Elem(), ModelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelArrayInput)(nil)).Elem(), ModelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelDockerSectionBaseImageRegistryInput)(nil)).Elem(), ModelDockerSectionBaseImageRegistryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelDockerSectionBaseImageRegistryPtrInput)(nil)).Elem(), ModelDockerSectionBaseImageRegistryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelDockerSectionResponseResponseBaseImageRegistryInput)(nil)).Elem(), ModelDockerSectionResponseResponseBaseImageRegistryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelDockerSectionResponseResponseBaseImageRegistryPtrInput)(nil)).Elem(), ModelDockerSectionResponseResponseBaseImageRegistryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionDockerInput)(nil)).Elem(), ModelEnvironmentDefinitionDockerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionDockerPtrInput)(nil)).Elem(), ModelEnvironmentDefinitionDockerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionPythonInput)(nil)).Elem(), ModelEnvironmentDefinitionPythonArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionPythonPtrInput)(nil)).Elem(), ModelEnvironmentDefinitionPythonArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionRInput)(nil)).Elem(), ModelEnvironmentDefinitionRArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionRPtrInput)(nil)).Elem(), ModelEnvironmentDefinitionRArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseDockerInput)(nil)).Elem(), ModelEnvironmentDefinitionResponseResponseDockerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseDockerPtrInput)(nil)).Elem(), ModelEnvironmentDefinitionResponseResponseDockerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponsePythonInput)(nil)).Elem(), ModelEnvironmentDefinitionResponseResponsePythonArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponsePythonPtrInput)(nil)).Elem(), ModelEnvironmentDefinitionResponseResponsePythonArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseRInput)(nil)).Elem(), ModelEnvironmentDefinitionResponseResponseRArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseRPtrInput)(nil)).Elem(), ModelEnvironmentDefinitionResponseResponseRArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseSparkInput)(nil)).Elem(), ModelEnvironmentDefinitionResponseResponseSparkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionResponseResponseSparkPtrInput)(nil)).Elem(), ModelEnvironmentDefinitionResponseResponseSparkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionSparkInput)(nil)).Elem(), ModelEnvironmentDefinitionSparkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelEnvironmentDefinitionSparkPtrInput)(nil)).Elem(), ModelEnvironmentDefinitionSparkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelResponseInput)(nil)).Elem(), ModelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModelResponseArrayInput)(nil)).Elem(), ModelResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeStateCountsResponseInput)(nil)).Elem(), NodeStateCountsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeStateCountsResponsePtrInput)(nil)).Elem(), NodeStateCountsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordResponseInput)(nil)).Elem(), PasswordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordResponseArrayInput)(nil)).Elem(), PasswordResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseInput)(nil)).Elem(), PrivateEndpointConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseArrayInput)(nil)).Elem(), PrivateEndpointConnectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponseInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponsePtrInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponseInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponsePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RCranPackageInput)(nil)).Elem(), RCranPackageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RCranPackageArrayInput)(nil)).Elem(), RCranPackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RCranPackageResponseInput)(nil)).Elem(), RCranPackageResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RCranPackageResponseArrayInput)(nil)).Elem(), RCranPackageResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RGitHubPackageInput)(nil)).Elem(), RGitHubPackageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RGitHubPackageArrayInput)(nil)).Elem(), RGitHubPackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RGitHubPackageResponseResponseInput)(nil)).Elem(), RGitHubPackageResponseResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RGitHubPackageResponseResponseArrayInput)(nil)).Elem(), RGitHubPackageResponseResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryListCredentialsResultResponseInput)(nil)).Elem(), RegistryListCredentialsResultResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdInput)(nil)).Elem(), ResourceIdArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdPtrInput)(nil)).Elem(), ResourceIdArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdResponseInput)(nil)).Elem(), ResourceIdResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdResponsePtrInput)(nil)).Elem(), ResourceIdResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleSettingsInput)(nil)).Elem(), ScaleSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleSettingsPtrInput)(nil)).Elem(), ScaleSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleSettingsResponseInput)(nil)).Elem(), ScaleSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleSettingsResponsePtrInput)(nil)).Elem(), ScaleSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceResponseBaseResponseErrorInput)(nil)).Elem(), ServiceResponseBaseResponseErrorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourceInput)(nil)).Elem(), SharedPrivateLinkResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourceArrayInput)(nil)).Elem(), SharedPrivateLinkResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourceResponseInput)(nil)).Elem(), SharedPrivateLinkResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedPrivateLinkResourceResponseArrayInput)(nil)).Elem(), SharedPrivateLinkResourceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SparkMavenPackageInput)(nil)).Elem(), SparkMavenPackageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SparkMavenPackageArrayInput)(nil)).Elem(), SparkMavenPackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SparkMavenPackageResponseInput)(nil)).Elem(), SparkMavenPackageResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SparkMavenPackageResponseArrayInput)(nil)).Elem(), SparkMavenPackageResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigurationInput)(nil)).Elem(), SslConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigurationPtrInput)(nil)).Elem(), SslConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigurationResponseInput)(nil)).Elem(), SslConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigurationResponsePtrInput)(nil)).Elem(), SslConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemServiceResponseInput)(nil)).Elem(), SystemServiceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemServiceResponseArrayInput)(nil)).Elem(), SystemServiceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAccountCredentialsInput)(nil)).Elem(), UserAccountCredentialsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAccountCredentialsPtrInput)(nil)).Elem(), UserAccountCredentialsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAccountCredentialsResponseInput)(nil)).Elem(), UserAccountCredentialsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAccountCredentialsResponsePtrInput)(nil)).Elem(), UserAccountCredentialsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedIdentityResponseInput)(nil)).Elem(), UserAssignedIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedIdentityResponseMapInput)(nil)).Elem(), UserAssignedIdentityResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineInput)(nil)).Elem(), VirtualMachineArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineImageInput)(nil)).Elem(), VirtualMachineImageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineImagePtrInput)(nil)).Elem(), VirtualMachineImageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineImageResponseInput)(nil)).Elem(), VirtualMachineImageResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineImageResponsePtrInput)(nil)).Elem(), VirtualMachineImageResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachinePropertiesInput)(nil)).Elem(), VirtualMachinePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachinePropertiesPtrInput)(nil)).Elem(), VirtualMachinePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineResponseInput)(nil)).Elem(), VirtualMachineResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineResponsePropertiesInput)(nil)).Elem(), VirtualMachineResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineResponsePropertiesPtrInput)(nil)).Elem(), VirtualMachineResponsePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineSshCredentialsInput)(nil)).Elem(), VirtualMachineSshCredentialsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineSshCredentialsPtrInput)(nil)).Elem(), VirtualMachineSshCredentialsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineSshCredentialsResponseInput)(nil)).Elem(), VirtualMachineSshCredentialsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineSshCredentialsResponsePtrInput)(nil)).Elem(), VirtualMachineSshCredentialsResponseArgs{})
	pulumi.RegisterOutputType(ACIServiceCreateRequestDataCollectionOutput{})
	pulumi.RegisterOutputType(ACIServiceCreateRequestDataCollectionPtrOutput{})
	pulumi.RegisterOutputType(ACIServiceCreateRequestEncryptionPropertiesOutput{})
	pulumi.RegisterOutputType(ACIServiceCreateRequestEncryptionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ACIServiceCreateRequestVnetConfigurationOutput{})
	pulumi.RegisterOutputType(ACIServiceCreateRequestVnetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ACIServiceResponseResponseOutput{})
	pulumi.RegisterOutputType(ACIServiceResponseResponseDataCollectionOutput{})
	pulumi.RegisterOutputType(ACIServiceResponseResponseDataCollectionPtrOutput{})
	pulumi.RegisterOutputType(ACIServiceResponseResponseEncryptionPropertiesOutput{})
	pulumi.RegisterOutputType(ACIServiceResponseResponseEncryptionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ACIServiceResponseResponseEnvironmentImageRequestOutput{})
	pulumi.RegisterOutputType(ACIServiceResponseResponseEnvironmentImageRequestPtrOutput{})
	pulumi.RegisterOutputType(ACIServiceResponseResponseVnetConfigurationOutput{})
	pulumi.RegisterOutputType(ACIServiceResponseResponseVnetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AKSOutput{})
	pulumi.RegisterOutputType(AKSPropertiesOutput{})
	pulumi.RegisterOutputType(AKSPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AKSReplicaStatusResponseErrorOutput{})
	pulumi.RegisterOutputType(AKSReplicaStatusResponseErrorPtrOutput{})
	pulumi.RegisterOutputType(AKSResponseOutput{})
	pulumi.RegisterOutputType(AKSResponsePropertiesOutput{})
	pulumi.RegisterOutputType(AKSResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestAutoScalerOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestAutoScalerPtrOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestDataCollectionOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestDataCollectionPtrOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestLivenessProbeRequirementsOutput{})
	pulumi.RegisterOutputType(AKSServiceCreateRequestLivenessProbeRequirementsPtrOutput{})
	pulumi.RegisterOutputType(AKSServiceResponseResponseOutput{})
	pulumi.RegisterOutputType(AKSServiceResponseResponseAutoScalerOutput{})
	pulumi.RegisterOutputType(AKSServiceResponseResponseAutoScalerPtrOutput{})
	pulumi.RegisterOutputType(AKSServiceResponseResponseDataCollectionOutput{})
	pulumi.RegisterOutputType(AKSServiceResponseResponseDataCollectionPtrOutput{})
	pulumi.RegisterOutputType(AKSServiceResponseResponseDeploymentStatusOutput{})
	pulumi.RegisterOutputType(AKSServiceResponseResponseEnvironmentImageRequestOutput{})
	pulumi.RegisterOutputType(AKSServiceResponseResponseEnvironmentImageRequestPtrOutput{})
	pulumi.RegisterOutputType(AKSServiceResponseResponseLivenessProbeRequirementsOutput{})
	pulumi.RegisterOutputType(AKSServiceResponseResponseLivenessProbeRequirementsPtrOutput{})
	pulumi.RegisterOutputType(AKSVariantResponseResponseOutput{})
	pulumi.RegisterOutputType(AksNetworkingConfigurationOutput{})
	pulumi.RegisterOutputType(AksNetworkingConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AksNetworkingConfigurationResponseOutput{})
	pulumi.RegisterOutputType(AksNetworkingConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(AmlComputeOutput{})
	pulumi.RegisterOutputType(AmlComputeNodeInformationResponseOutput{})
	pulumi.RegisterOutputType(AmlComputeNodeInformationResponseArrayOutput{})
	pulumi.RegisterOutputType(AmlComputePropertiesOutput{})
	pulumi.RegisterOutputType(AmlComputePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AmlComputeResponseOutput{})
	pulumi.RegisterOutputType(AmlComputeResponsePropertiesOutput{})
	pulumi.RegisterOutputType(AmlComputeResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ContainerResourceRequirementsOutput{})
	pulumi.RegisterOutputType(ContainerResourceRequirementsPtrOutput{})
	pulumi.RegisterOutputType(ContainerResourceRequirementsResponseOutput{})
	pulumi.RegisterOutputType(ContainerResourceRequirementsResponsePtrOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestEnvironmentImageRequestOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestEnvironmentImageRequestPtrOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestKeysOutput{})
	pulumi.RegisterOutputType(CreateServiceRequestKeysPtrOutput{})
	pulumi.RegisterOutputType(DataFactoryOutput{})
	pulumi.RegisterOutputType(DataFactoryResponseOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsPropertiesOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsResponseOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsResponsePropertiesOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(DatabricksOutput{})
	pulumi.RegisterOutputType(DatabricksPropertiesOutput{})
	pulumi.RegisterOutputType(DatabricksPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DatabricksResponseOutput{})
	pulumi.RegisterOutputType(DatabricksResponsePropertiesOutput{})
	pulumi.RegisterOutputType(DatabricksResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(DatasetReferenceOutput{})
	pulumi.RegisterOutputType(DatasetReferenceArrayOutput{})
	pulumi.RegisterOutputType(DatasetReferenceResponseOutput{})
	pulumi.RegisterOutputType(DatasetReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentReferenceOutput{})
	pulumi.RegisterOutputType(EnvironmentImageRequestEnvironmentReferencePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentImageResponseResponseEnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentImageResponseResponseEnvironmentPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentImageResponseResponseEnvironmentReferenceOutput{})
	pulumi.RegisterOutputType(EnvironmentImageResponseResponseEnvironmentReferencePtrOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseOutput{})
	pulumi.RegisterOutputType(HDInsightOutput{})
	pulumi.RegisterOutputType(HDInsightPropertiesOutput{})
	pulumi.RegisterOutputType(HDInsightPropertiesPtrOutput{})
	pulumi.RegisterOutputType(HDInsightResponseOutput{})
	pulumi.RegisterOutputType(HDInsightResponsePropertiesOutput{})
	pulumi.RegisterOutputType(HDInsightResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageAssetOutput{})
	pulumi.RegisterOutputType(ImageAssetArrayOutput{})
	pulumi.RegisterOutputType(ImageAssetResponseOutput{})
	pulumi.RegisterOutputType(ImageAssetResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LinkedWorkspacePropsOutput{})
	pulumi.RegisterOutputType(LinkedWorkspacePropsPtrOutput{})
	pulumi.RegisterOutputType(LinkedWorkspacePropsResponseOutput{})
	pulumi.RegisterOutputType(LinkedWorkspacePropsResponsePtrOutput{})
	pulumi.RegisterOutputType(MachineLearningServiceErrorResponseOutput{})
	pulumi.RegisterOutputType(MachineLearningServiceErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ModelOutput{})
	pulumi.RegisterOutputType(ModelArrayOutput{})
	pulumi.RegisterOutputType(ModelDockerSectionBaseImageRegistryOutput{})
	pulumi.RegisterOutputType(ModelDockerSectionBaseImageRegistryPtrOutput{})
	pulumi.RegisterOutputType(ModelDockerSectionResponseResponseBaseImageRegistryOutput{})
	pulumi.RegisterOutputType(ModelDockerSectionResponseResponseBaseImageRegistryPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionDockerOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionDockerPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionPythonOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionPythonPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionROutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionRPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionResponseResponseDockerOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionResponseResponseDockerPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionResponseResponsePythonOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionResponseResponsePythonPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionResponseResponseROutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionResponseResponseRPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionResponseResponseSparkOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionResponseResponseSparkPtrOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionSparkOutput{})
	pulumi.RegisterOutputType(ModelEnvironmentDefinitionSparkPtrOutput{})
	pulumi.RegisterOutputType(ModelResponseOutput{})
	pulumi.RegisterOutputType(ModelResponseArrayOutput{})
	pulumi.RegisterOutputType(NodeStateCountsResponseOutput{})
	pulumi.RegisterOutputType(NodeStateCountsResponsePtrOutput{})
	pulumi.RegisterOutputType(PasswordResponseOutput{})
	pulumi.RegisterOutputType(PasswordResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(RCranPackageOutput{})
	pulumi.RegisterOutputType(RCranPackageArrayOutput{})
	pulumi.RegisterOutputType(RCranPackageResponseOutput{})
	pulumi.RegisterOutputType(RCranPackageResponseArrayOutput{})
	pulumi.RegisterOutputType(RGitHubPackageOutput{})
	pulumi.RegisterOutputType(RGitHubPackageArrayOutput{})
	pulumi.RegisterOutputType(RGitHubPackageResponseResponseOutput{})
	pulumi.RegisterOutputType(RGitHubPackageResponseResponseArrayOutput{})
	pulumi.RegisterOutputType(RegistryListCredentialsResultResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdOutput{})
	pulumi.RegisterOutputType(ResourceIdPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(ScaleSettingsOutput{})
	pulumi.RegisterOutputType(ScaleSettingsPtrOutput{})
	pulumi.RegisterOutputType(ScaleSettingsResponseOutput{})
	pulumi.RegisterOutputType(ScaleSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceResponseBaseResponseErrorOutput{})
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
	pulumi.RegisterOutputType(SparkMavenPackageResponseOutput{})
	pulumi.RegisterOutputType(SparkMavenPackageResponseArrayOutput{})
	pulumi.RegisterOutputType(SslConfigurationOutput{})
	pulumi.RegisterOutputType(SslConfigurationPtrOutput{})
	pulumi.RegisterOutputType(SslConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SslConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemServiceResponseOutput{})
	pulumi.RegisterOutputType(SystemServiceResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsPtrOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsResponseOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VirtualMachineOutput{})
	pulumi.RegisterOutputType(VirtualMachineImageOutput{})
	pulumi.RegisterOutputType(VirtualMachineImagePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineImageResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineImageResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachinePropertiesOutput{})
	pulumi.RegisterOutputType(VirtualMachinePropertiesPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineResponsePropertiesOutput{})
	pulumi.RegisterOutputType(VirtualMachineResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineSshCredentialsOutput{})
	pulumi.RegisterOutputType(VirtualMachineSshCredentialsPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineSshCredentialsResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineSshCredentialsResponsePtrOutput{})
}
