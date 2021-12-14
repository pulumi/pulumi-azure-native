


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActiveDirectoryProperties struct {
	AzureStorageSid   string `pulumi:"azureStorageSid"`
	DomainGuid        string `pulumi:"domainGuid"`
	DomainName        string `pulumi:"domainName"`
	DomainSid         string `pulumi:"domainSid"`
	ForestName        string `pulumi:"forestName"`
	NetBiosDomainName string `pulumi:"netBiosDomainName"`
}





type ActiveDirectoryPropertiesInput interface {
	pulumi.Input

	ToActiveDirectoryPropertiesOutput() ActiveDirectoryPropertiesOutput
	ToActiveDirectoryPropertiesOutputWithContext(context.Context) ActiveDirectoryPropertiesOutput
}

type ActiveDirectoryPropertiesArgs struct {
	AzureStorageSid   pulumi.StringInput `pulumi:"azureStorageSid"`
	DomainGuid        pulumi.StringInput `pulumi:"domainGuid"`
	DomainName        pulumi.StringInput `pulumi:"domainName"`
	DomainSid         pulumi.StringInput `pulumi:"domainSid"`
	ForestName        pulumi.StringInput `pulumi:"forestName"`
	NetBiosDomainName pulumi.StringInput `pulumi:"netBiosDomainName"`
}

func (ActiveDirectoryPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryProperties)(nil)).Elem()
}

func (i ActiveDirectoryPropertiesArgs) ToActiveDirectoryPropertiesOutput() ActiveDirectoryPropertiesOutput {
	return i.ToActiveDirectoryPropertiesOutputWithContext(context.Background())
}

func (i ActiveDirectoryPropertiesArgs) ToActiveDirectoryPropertiesOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryPropertiesOutput)
}

func (i ActiveDirectoryPropertiesArgs) ToActiveDirectoryPropertiesPtrOutput() ActiveDirectoryPropertiesPtrOutput {
	return i.ToActiveDirectoryPropertiesPtrOutputWithContext(context.Background())
}

func (i ActiveDirectoryPropertiesArgs) ToActiveDirectoryPropertiesPtrOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryPropertiesOutput).ToActiveDirectoryPropertiesPtrOutputWithContext(ctx)
}









type ActiveDirectoryPropertiesPtrInput interface {
	pulumi.Input

	ToActiveDirectoryPropertiesPtrOutput() ActiveDirectoryPropertiesPtrOutput
	ToActiveDirectoryPropertiesPtrOutputWithContext(context.Context) ActiveDirectoryPropertiesPtrOutput
}

type activeDirectoryPropertiesPtrType ActiveDirectoryPropertiesArgs

func ActiveDirectoryPropertiesPtr(v *ActiveDirectoryPropertiesArgs) ActiveDirectoryPropertiesPtrInput {
	return (*activeDirectoryPropertiesPtrType)(v)
}

func (*activeDirectoryPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryProperties)(nil)).Elem()
}

func (i *activeDirectoryPropertiesPtrType) ToActiveDirectoryPropertiesPtrOutput() ActiveDirectoryPropertiesPtrOutput {
	return i.ToActiveDirectoryPropertiesPtrOutputWithContext(context.Background())
}

func (i *activeDirectoryPropertiesPtrType) ToActiveDirectoryPropertiesPtrOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryPropertiesPtrOutput)
}

type ActiveDirectoryPropertiesOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryProperties)(nil)).Elem()
}

func (o ActiveDirectoryPropertiesOutput) ToActiveDirectoryPropertiesOutput() ActiveDirectoryPropertiesOutput {
	return o
}

func (o ActiveDirectoryPropertiesOutput) ToActiveDirectoryPropertiesOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesOutput {
	return o
}

func (o ActiveDirectoryPropertiesOutput) ToActiveDirectoryPropertiesPtrOutput() ActiveDirectoryPropertiesPtrOutput {
	return o.ToActiveDirectoryPropertiesPtrOutputWithContext(context.Background())
}

func (o ActiveDirectoryPropertiesOutput) ToActiveDirectoryPropertiesPtrOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActiveDirectoryProperties) *ActiveDirectoryProperties {
		return &v
	}).(ActiveDirectoryPropertiesPtrOutput)
}

func (o ActiveDirectoryPropertiesOutput) AzureStorageSid() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryProperties) string { return v.AzureStorageSid }).(pulumi.StringOutput)
}

func (o ActiveDirectoryPropertiesOutput) DomainGuid() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryProperties) string { return v.DomainGuid }).(pulumi.StringOutput)
}

func (o ActiveDirectoryPropertiesOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryProperties) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o ActiveDirectoryPropertiesOutput) DomainSid() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryProperties) string { return v.DomainSid }).(pulumi.StringOutput)
}

func (o ActiveDirectoryPropertiesOutput) ForestName() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryProperties) string { return v.ForestName }).(pulumi.StringOutput)
}

func (o ActiveDirectoryPropertiesOutput) NetBiosDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryProperties) string { return v.NetBiosDomainName }).(pulumi.StringOutput)
}

type ActiveDirectoryPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryProperties)(nil)).Elem()
}

func (o ActiveDirectoryPropertiesPtrOutput) ToActiveDirectoryPropertiesPtrOutput() ActiveDirectoryPropertiesPtrOutput {
	return o
}

func (o ActiveDirectoryPropertiesPtrOutput) ToActiveDirectoryPropertiesPtrOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesPtrOutput {
	return o
}

func (o ActiveDirectoryPropertiesPtrOutput) Elem() ActiveDirectoryPropertiesOutput {
	return o.ApplyT(func(v *ActiveDirectoryProperties) ActiveDirectoryProperties {
		if v != nil {
			return *v
		}
		var ret ActiveDirectoryProperties
		return ret
	}).(ActiveDirectoryPropertiesOutput)
}

func (o ActiveDirectoryPropertiesPtrOutput) AzureStorageSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AzureStorageSid
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryPropertiesPtrOutput) DomainGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DomainGuid
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryPropertiesPtrOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DomainName
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryPropertiesPtrOutput) DomainSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DomainSid
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryPropertiesPtrOutput) ForestName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ForestName
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryPropertiesPtrOutput) NetBiosDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryProperties) *string {
		if v == nil {
			return nil
		}
		return &v.NetBiosDomainName
	}).(pulumi.StringPtrOutput)
}

type ActiveDirectoryPropertiesResponse struct {
	AzureStorageSid   string `pulumi:"azureStorageSid"`
	DomainGuid        string `pulumi:"domainGuid"`
	DomainName        string `pulumi:"domainName"`
	DomainSid         string `pulumi:"domainSid"`
	ForestName        string `pulumi:"forestName"`
	NetBiosDomainName string `pulumi:"netBiosDomainName"`
}





type ActiveDirectoryPropertiesResponseInput interface {
	pulumi.Input

	ToActiveDirectoryPropertiesResponseOutput() ActiveDirectoryPropertiesResponseOutput
	ToActiveDirectoryPropertiesResponseOutputWithContext(context.Context) ActiveDirectoryPropertiesResponseOutput
}

type ActiveDirectoryPropertiesResponseArgs struct {
	AzureStorageSid   pulumi.StringInput `pulumi:"azureStorageSid"`
	DomainGuid        pulumi.StringInput `pulumi:"domainGuid"`
	DomainName        pulumi.StringInput `pulumi:"domainName"`
	DomainSid         pulumi.StringInput `pulumi:"domainSid"`
	ForestName        pulumi.StringInput `pulumi:"forestName"`
	NetBiosDomainName pulumi.StringInput `pulumi:"netBiosDomainName"`
}

func (ActiveDirectoryPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryPropertiesResponse)(nil)).Elem()
}

func (i ActiveDirectoryPropertiesResponseArgs) ToActiveDirectoryPropertiesResponseOutput() ActiveDirectoryPropertiesResponseOutput {
	return i.ToActiveDirectoryPropertiesResponseOutputWithContext(context.Background())
}

func (i ActiveDirectoryPropertiesResponseArgs) ToActiveDirectoryPropertiesResponseOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryPropertiesResponseOutput)
}

func (i ActiveDirectoryPropertiesResponseArgs) ToActiveDirectoryPropertiesResponsePtrOutput() ActiveDirectoryPropertiesResponsePtrOutput {
	return i.ToActiveDirectoryPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ActiveDirectoryPropertiesResponseArgs) ToActiveDirectoryPropertiesResponsePtrOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryPropertiesResponseOutput).ToActiveDirectoryPropertiesResponsePtrOutputWithContext(ctx)
}









type ActiveDirectoryPropertiesResponsePtrInput interface {
	pulumi.Input

	ToActiveDirectoryPropertiesResponsePtrOutput() ActiveDirectoryPropertiesResponsePtrOutput
	ToActiveDirectoryPropertiesResponsePtrOutputWithContext(context.Context) ActiveDirectoryPropertiesResponsePtrOutput
}

type activeDirectoryPropertiesResponsePtrType ActiveDirectoryPropertiesResponseArgs

func ActiveDirectoryPropertiesResponsePtr(v *ActiveDirectoryPropertiesResponseArgs) ActiveDirectoryPropertiesResponsePtrInput {
	return (*activeDirectoryPropertiesResponsePtrType)(v)
}

func (*activeDirectoryPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryPropertiesResponse)(nil)).Elem()
}

func (i *activeDirectoryPropertiesResponsePtrType) ToActiveDirectoryPropertiesResponsePtrOutput() ActiveDirectoryPropertiesResponsePtrOutput {
	return i.ToActiveDirectoryPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *activeDirectoryPropertiesResponsePtrType) ToActiveDirectoryPropertiesResponsePtrOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryPropertiesResponsePtrOutput)
}

type ActiveDirectoryPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActiveDirectoryPropertiesResponse)(nil)).Elem()
}

func (o ActiveDirectoryPropertiesResponseOutput) ToActiveDirectoryPropertiesResponseOutput() ActiveDirectoryPropertiesResponseOutput {
	return o
}

func (o ActiveDirectoryPropertiesResponseOutput) ToActiveDirectoryPropertiesResponseOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesResponseOutput {
	return o
}

func (o ActiveDirectoryPropertiesResponseOutput) ToActiveDirectoryPropertiesResponsePtrOutput() ActiveDirectoryPropertiesResponsePtrOutput {
	return o.ToActiveDirectoryPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ActiveDirectoryPropertiesResponseOutput) ToActiveDirectoryPropertiesResponsePtrOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActiveDirectoryPropertiesResponse) *ActiveDirectoryPropertiesResponse {
		return &v
	}).(ActiveDirectoryPropertiesResponsePtrOutput)
}

func (o ActiveDirectoryPropertiesResponseOutput) AzureStorageSid() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryPropertiesResponse) string { return v.AzureStorageSid }).(pulumi.StringOutput)
}

func (o ActiveDirectoryPropertiesResponseOutput) DomainGuid() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryPropertiesResponse) string { return v.DomainGuid }).(pulumi.StringOutput)
}

func (o ActiveDirectoryPropertiesResponseOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryPropertiesResponse) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o ActiveDirectoryPropertiesResponseOutput) DomainSid() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryPropertiesResponse) string { return v.DomainSid }).(pulumi.StringOutput)
}

func (o ActiveDirectoryPropertiesResponseOutput) ForestName() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryPropertiesResponse) string { return v.ForestName }).(pulumi.StringOutput)
}

func (o ActiveDirectoryPropertiesResponseOutput) NetBiosDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ActiveDirectoryPropertiesResponse) string { return v.NetBiosDomainName }).(pulumi.StringOutput)
}

type ActiveDirectoryPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryPropertiesResponse)(nil)).Elem()
}

func (o ActiveDirectoryPropertiesResponsePtrOutput) ToActiveDirectoryPropertiesResponsePtrOutput() ActiveDirectoryPropertiesResponsePtrOutput {
	return o
}

func (o ActiveDirectoryPropertiesResponsePtrOutput) ToActiveDirectoryPropertiesResponsePtrOutputWithContext(ctx context.Context) ActiveDirectoryPropertiesResponsePtrOutput {
	return o
}

func (o ActiveDirectoryPropertiesResponsePtrOutput) Elem() ActiveDirectoryPropertiesResponseOutput {
	return o.ApplyT(func(v *ActiveDirectoryPropertiesResponse) ActiveDirectoryPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ActiveDirectoryPropertiesResponse
		return ret
	}).(ActiveDirectoryPropertiesResponseOutput)
}

func (o ActiveDirectoryPropertiesResponsePtrOutput) AzureStorageSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AzureStorageSid
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryPropertiesResponsePtrOutput) DomainGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DomainGuid
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryPropertiesResponsePtrOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DomainName
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryPropertiesResponsePtrOutput) DomainSid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DomainSid
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryPropertiesResponsePtrOutput) ForestName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ForestName
	}).(pulumi.StringPtrOutput)
}

func (o ActiveDirectoryPropertiesResponsePtrOutput) NetBiosDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActiveDirectoryPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NetBiosDomainName
	}).(pulumi.StringPtrOutput)
}

type AzureFilesIdentityBasedAuthentication struct {
	ActiveDirectoryProperties *ActiveDirectoryProperties `pulumi:"activeDirectoryProperties"`
	DirectoryServiceOptions   string                     `pulumi:"directoryServiceOptions"`
}





type AzureFilesIdentityBasedAuthenticationInput interface {
	pulumi.Input

	ToAzureFilesIdentityBasedAuthenticationOutput() AzureFilesIdentityBasedAuthenticationOutput
	ToAzureFilesIdentityBasedAuthenticationOutputWithContext(context.Context) AzureFilesIdentityBasedAuthenticationOutput
}

type AzureFilesIdentityBasedAuthenticationArgs struct {
	ActiveDirectoryProperties ActiveDirectoryPropertiesPtrInput `pulumi:"activeDirectoryProperties"`
	DirectoryServiceOptions   pulumi.StringInput                `pulumi:"directoryServiceOptions"`
}

func (AzureFilesIdentityBasedAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFilesIdentityBasedAuthentication)(nil)).Elem()
}

func (i AzureFilesIdentityBasedAuthenticationArgs) ToAzureFilesIdentityBasedAuthenticationOutput() AzureFilesIdentityBasedAuthenticationOutput {
	return i.ToAzureFilesIdentityBasedAuthenticationOutputWithContext(context.Background())
}

func (i AzureFilesIdentityBasedAuthenticationArgs) ToAzureFilesIdentityBasedAuthenticationOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFilesIdentityBasedAuthenticationOutput)
}

func (i AzureFilesIdentityBasedAuthenticationArgs) ToAzureFilesIdentityBasedAuthenticationPtrOutput() AzureFilesIdentityBasedAuthenticationPtrOutput {
	return i.ToAzureFilesIdentityBasedAuthenticationPtrOutputWithContext(context.Background())
}

func (i AzureFilesIdentityBasedAuthenticationArgs) ToAzureFilesIdentityBasedAuthenticationPtrOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFilesIdentityBasedAuthenticationOutput).ToAzureFilesIdentityBasedAuthenticationPtrOutputWithContext(ctx)
}









type AzureFilesIdentityBasedAuthenticationPtrInput interface {
	pulumi.Input

	ToAzureFilesIdentityBasedAuthenticationPtrOutput() AzureFilesIdentityBasedAuthenticationPtrOutput
	ToAzureFilesIdentityBasedAuthenticationPtrOutputWithContext(context.Context) AzureFilesIdentityBasedAuthenticationPtrOutput
}

type azureFilesIdentityBasedAuthenticationPtrType AzureFilesIdentityBasedAuthenticationArgs

func AzureFilesIdentityBasedAuthenticationPtr(v *AzureFilesIdentityBasedAuthenticationArgs) AzureFilesIdentityBasedAuthenticationPtrInput {
	return (*azureFilesIdentityBasedAuthenticationPtrType)(v)
}

func (*azureFilesIdentityBasedAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFilesIdentityBasedAuthentication)(nil)).Elem()
}

func (i *azureFilesIdentityBasedAuthenticationPtrType) ToAzureFilesIdentityBasedAuthenticationPtrOutput() AzureFilesIdentityBasedAuthenticationPtrOutput {
	return i.ToAzureFilesIdentityBasedAuthenticationPtrOutputWithContext(context.Background())
}

func (i *azureFilesIdentityBasedAuthenticationPtrType) ToAzureFilesIdentityBasedAuthenticationPtrOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFilesIdentityBasedAuthenticationPtrOutput)
}

type AzureFilesIdentityBasedAuthenticationOutput struct{ *pulumi.OutputState }

func (AzureFilesIdentityBasedAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFilesIdentityBasedAuthentication)(nil)).Elem()
}

func (o AzureFilesIdentityBasedAuthenticationOutput) ToAzureFilesIdentityBasedAuthenticationOutput() AzureFilesIdentityBasedAuthenticationOutput {
	return o
}

func (o AzureFilesIdentityBasedAuthenticationOutput) ToAzureFilesIdentityBasedAuthenticationOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationOutput {
	return o
}

func (o AzureFilesIdentityBasedAuthenticationOutput) ToAzureFilesIdentityBasedAuthenticationPtrOutput() AzureFilesIdentityBasedAuthenticationPtrOutput {
	return o.ToAzureFilesIdentityBasedAuthenticationPtrOutputWithContext(context.Background())
}

func (o AzureFilesIdentityBasedAuthenticationOutput) ToAzureFilesIdentityBasedAuthenticationPtrOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFilesIdentityBasedAuthentication) *AzureFilesIdentityBasedAuthentication {
		return &v
	}).(AzureFilesIdentityBasedAuthenticationPtrOutput)
}

func (o AzureFilesIdentityBasedAuthenticationOutput) ActiveDirectoryProperties() ActiveDirectoryPropertiesPtrOutput {
	return o.ApplyT(func(v AzureFilesIdentityBasedAuthentication) *ActiveDirectoryProperties {
		return v.ActiveDirectoryProperties
	}).(ActiveDirectoryPropertiesPtrOutput)
}

func (o AzureFilesIdentityBasedAuthenticationOutput) DirectoryServiceOptions() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFilesIdentityBasedAuthentication) string { return v.DirectoryServiceOptions }).(pulumi.StringOutput)
}

type AzureFilesIdentityBasedAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (AzureFilesIdentityBasedAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFilesIdentityBasedAuthentication)(nil)).Elem()
}

func (o AzureFilesIdentityBasedAuthenticationPtrOutput) ToAzureFilesIdentityBasedAuthenticationPtrOutput() AzureFilesIdentityBasedAuthenticationPtrOutput {
	return o
}

func (o AzureFilesIdentityBasedAuthenticationPtrOutput) ToAzureFilesIdentityBasedAuthenticationPtrOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationPtrOutput {
	return o
}

func (o AzureFilesIdentityBasedAuthenticationPtrOutput) Elem() AzureFilesIdentityBasedAuthenticationOutput {
	return o.ApplyT(func(v *AzureFilesIdentityBasedAuthentication) AzureFilesIdentityBasedAuthentication {
		if v != nil {
			return *v
		}
		var ret AzureFilesIdentityBasedAuthentication
		return ret
	}).(AzureFilesIdentityBasedAuthenticationOutput)
}

func (o AzureFilesIdentityBasedAuthenticationPtrOutput) ActiveDirectoryProperties() ActiveDirectoryPropertiesPtrOutput {
	return o.ApplyT(func(v *AzureFilesIdentityBasedAuthentication) *ActiveDirectoryProperties {
		if v == nil {
			return nil
		}
		return v.ActiveDirectoryProperties
	}).(ActiveDirectoryPropertiesPtrOutput)
}

func (o AzureFilesIdentityBasedAuthenticationPtrOutput) DirectoryServiceOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFilesIdentityBasedAuthentication) *string {
		if v == nil {
			return nil
		}
		return &v.DirectoryServiceOptions
	}).(pulumi.StringPtrOutput)
}

type AzureFilesIdentityBasedAuthenticationResponse struct {
	ActiveDirectoryProperties *ActiveDirectoryPropertiesResponse `pulumi:"activeDirectoryProperties"`
	DirectoryServiceOptions   string                             `pulumi:"directoryServiceOptions"`
}





type AzureFilesIdentityBasedAuthenticationResponseInput interface {
	pulumi.Input

	ToAzureFilesIdentityBasedAuthenticationResponseOutput() AzureFilesIdentityBasedAuthenticationResponseOutput
	ToAzureFilesIdentityBasedAuthenticationResponseOutputWithContext(context.Context) AzureFilesIdentityBasedAuthenticationResponseOutput
}

type AzureFilesIdentityBasedAuthenticationResponseArgs struct {
	ActiveDirectoryProperties ActiveDirectoryPropertiesResponsePtrInput `pulumi:"activeDirectoryProperties"`
	DirectoryServiceOptions   pulumi.StringInput                        `pulumi:"directoryServiceOptions"`
}

func (AzureFilesIdentityBasedAuthenticationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFilesIdentityBasedAuthenticationResponse)(nil)).Elem()
}

func (i AzureFilesIdentityBasedAuthenticationResponseArgs) ToAzureFilesIdentityBasedAuthenticationResponseOutput() AzureFilesIdentityBasedAuthenticationResponseOutput {
	return i.ToAzureFilesIdentityBasedAuthenticationResponseOutputWithContext(context.Background())
}

func (i AzureFilesIdentityBasedAuthenticationResponseArgs) ToAzureFilesIdentityBasedAuthenticationResponseOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFilesIdentityBasedAuthenticationResponseOutput)
}

func (i AzureFilesIdentityBasedAuthenticationResponseArgs) ToAzureFilesIdentityBasedAuthenticationResponsePtrOutput() AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return i.ToAzureFilesIdentityBasedAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i AzureFilesIdentityBasedAuthenticationResponseArgs) ToAzureFilesIdentityBasedAuthenticationResponsePtrOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFilesIdentityBasedAuthenticationResponseOutput).ToAzureFilesIdentityBasedAuthenticationResponsePtrOutputWithContext(ctx)
}









type AzureFilesIdentityBasedAuthenticationResponsePtrInput interface {
	pulumi.Input

	ToAzureFilesIdentityBasedAuthenticationResponsePtrOutput() AzureFilesIdentityBasedAuthenticationResponsePtrOutput
	ToAzureFilesIdentityBasedAuthenticationResponsePtrOutputWithContext(context.Context) AzureFilesIdentityBasedAuthenticationResponsePtrOutput
}

type azureFilesIdentityBasedAuthenticationResponsePtrType AzureFilesIdentityBasedAuthenticationResponseArgs

func AzureFilesIdentityBasedAuthenticationResponsePtr(v *AzureFilesIdentityBasedAuthenticationResponseArgs) AzureFilesIdentityBasedAuthenticationResponsePtrInput {
	return (*azureFilesIdentityBasedAuthenticationResponsePtrType)(v)
}

func (*azureFilesIdentityBasedAuthenticationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFilesIdentityBasedAuthenticationResponse)(nil)).Elem()
}

func (i *azureFilesIdentityBasedAuthenticationResponsePtrType) ToAzureFilesIdentityBasedAuthenticationResponsePtrOutput() AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return i.ToAzureFilesIdentityBasedAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i *azureFilesIdentityBasedAuthenticationResponsePtrType) ToAzureFilesIdentityBasedAuthenticationResponsePtrOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFilesIdentityBasedAuthenticationResponsePtrOutput)
}

type AzureFilesIdentityBasedAuthenticationResponseOutput struct{ *pulumi.OutputState }

func (AzureFilesIdentityBasedAuthenticationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFilesIdentityBasedAuthenticationResponse)(nil)).Elem()
}

func (o AzureFilesIdentityBasedAuthenticationResponseOutput) ToAzureFilesIdentityBasedAuthenticationResponseOutput() AzureFilesIdentityBasedAuthenticationResponseOutput {
	return o
}

func (o AzureFilesIdentityBasedAuthenticationResponseOutput) ToAzureFilesIdentityBasedAuthenticationResponseOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationResponseOutput {
	return o
}

func (o AzureFilesIdentityBasedAuthenticationResponseOutput) ToAzureFilesIdentityBasedAuthenticationResponsePtrOutput() AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return o.ToAzureFilesIdentityBasedAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (o AzureFilesIdentityBasedAuthenticationResponseOutput) ToAzureFilesIdentityBasedAuthenticationResponsePtrOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFilesIdentityBasedAuthenticationResponse) *AzureFilesIdentityBasedAuthenticationResponse {
		return &v
	}).(AzureFilesIdentityBasedAuthenticationResponsePtrOutput)
}

func (o AzureFilesIdentityBasedAuthenticationResponseOutput) ActiveDirectoryProperties() ActiveDirectoryPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AzureFilesIdentityBasedAuthenticationResponse) *ActiveDirectoryPropertiesResponse {
		return v.ActiveDirectoryProperties
	}).(ActiveDirectoryPropertiesResponsePtrOutput)
}

func (o AzureFilesIdentityBasedAuthenticationResponseOutput) DirectoryServiceOptions() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFilesIdentityBasedAuthenticationResponse) string { return v.DirectoryServiceOptions }).(pulumi.StringOutput)
}

type AzureFilesIdentityBasedAuthenticationResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureFilesIdentityBasedAuthenticationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFilesIdentityBasedAuthenticationResponse)(nil)).Elem()
}

func (o AzureFilesIdentityBasedAuthenticationResponsePtrOutput) ToAzureFilesIdentityBasedAuthenticationResponsePtrOutput() AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return o
}

func (o AzureFilesIdentityBasedAuthenticationResponsePtrOutput) ToAzureFilesIdentityBasedAuthenticationResponsePtrOutputWithContext(ctx context.Context) AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return o
}

func (o AzureFilesIdentityBasedAuthenticationResponsePtrOutput) Elem() AzureFilesIdentityBasedAuthenticationResponseOutput {
	return o.ApplyT(func(v *AzureFilesIdentityBasedAuthenticationResponse) AzureFilesIdentityBasedAuthenticationResponse {
		if v != nil {
			return *v
		}
		var ret AzureFilesIdentityBasedAuthenticationResponse
		return ret
	}).(AzureFilesIdentityBasedAuthenticationResponseOutput)
}

func (o AzureFilesIdentityBasedAuthenticationResponsePtrOutput) ActiveDirectoryProperties() ActiveDirectoryPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AzureFilesIdentityBasedAuthenticationResponse) *ActiveDirectoryPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ActiveDirectoryProperties
	}).(ActiveDirectoryPropertiesResponsePtrOutput)
}

func (o AzureFilesIdentityBasedAuthenticationResponsePtrOutput) DirectoryServiceOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFilesIdentityBasedAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DirectoryServiceOptions
	}).(pulumi.StringPtrOutput)
}

type BlobInventoryPolicyDefinition struct {
	Filters BlobInventoryPolicyFilter `pulumi:"filters"`
}





type BlobInventoryPolicyDefinitionInput interface {
	pulumi.Input

	ToBlobInventoryPolicyDefinitionOutput() BlobInventoryPolicyDefinitionOutput
	ToBlobInventoryPolicyDefinitionOutputWithContext(context.Context) BlobInventoryPolicyDefinitionOutput
}

type BlobInventoryPolicyDefinitionArgs struct {
	Filters BlobInventoryPolicyFilterInput `pulumi:"filters"`
}

func (BlobInventoryPolicyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyDefinition)(nil)).Elem()
}

func (i BlobInventoryPolicyDefinitionArgs) ToBlobInventoryPolicyDefinitionOutput() BlobInventoryPolicyDefinitionOutput {
	return i.ToBlobInventoryPolicyDefinitionOutputWithContext(context.Background())
}

func (i BlobInventoryPolicyDefinitionArgs) ToBlobInventoryPolicyDefinitionOutputWithContext(ctx context.Context) BlobInventoryPolicyDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyDefinitionOutput)
}

type BlobInventoryPolicyDefinitionOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicyDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyDefinition)(nil)).Elem()
}

func (o BlobInventoryPolicyDefinitionOutput) ToBlobInventoryPolicyDefinitionOutput() BlobInventoryPolicyDefinitionOutput {
	return o
}

func (o BlobInventoryPolicyDefinitionOutput) ToBlobInventoryPolicyDefinitionOutputWithContext(ctx context.Context) BlobInventoryPolicyDefinitionOutput {
	return o
}

func (o BlobInventoryPolicyDefinitionOutput) Filters() BlobInventoryPolicyFilterOutput {
	return o.ApplyT(func(v BlobInventoryPolicyDefinition) BlobInventoryPolicyFilter { return v.Filters }).(BlobInventoryPolicyFilterOutput)
}

type BlobInventoryPolicyDefinitionResponse struct {
	Filters BlobInventoryPolicyFilterResponse `pulumi:"filters"`
}





type BlobInventoryPolicyDefinitionResponseInput interface {
	pulumi.Input

	ToBlobInventoryPolicyDefinitionResponseOutput() BlobInventoryPolicyDefinitionResponseOutput
	ToBlobInventoryPolicyDefinitionResponseOutputWithContext(context.Context) BlobInventoryPolicyDefinitionResponseOutput
}

type BlobInventoryPolicyDefinitionResponseArgs struct {
	Filters BlobInventoryPolicyFilterResponseInput `pulumi:"filters"`
}

func (BlobInventoryPolicyDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyDefinitionResponse)(nil)).Elem()
}

func (i BlobInventoryPolicyDefinitionResponseArgs) ToBlobInventoryPolicyDefinitionResponseOutput() BlobInventoryPolicyDefinitionResponseOutput {
	return i.ToBlobInventoryPolicyDefinitionResponseOutputWithContext(context.Background())
}

func (i BlobInventoryPolicyDefinitionResponseArgs) ToBlobInventoryPolicyDefinitionResponseOutputWithContext(ctx context.Context) BlobInventoryPolicyDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyDefinitionResponseOutput)
}

type BlobInventoryPolicyDefinitionResponseOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicyDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyDefinitionResponse)(nil)).Elem()
}

func (o BlobInventoryPolicyDefinitionResponseOutput) ToBlobInventoryPolicyDefinitionResponseOutput() BlobInventoryPolicyDefinitionResponseOutput {
	return o
}

func (o BlobInventoryPolicyDefinitionResponseOutput) ToBlobInventoryPolicyDefinitionResponseOutputWithContext(ctx context.Context) BlobInventoryPolicyDefinitionResponseOutput {
	return o
}

func (o BlobInventoryPolicyDefinitionResponseOutput) Filters() BlobInventoryPolicyFilterResponseOutput {
	return o.ApplyT(func(v BlobInventoryPolicyDefinitionResponse) BlobInventoryPolicyFilterResponse { return v.Filters }).(BlobInventoryPolicyFilterResponseOutput)
}

type BlobInventoryPolicyFilter struct {
	BlobTypes           []string `pulumi:"blobTypes"`
	IncludeBlobVersions *bool    `pulumi:"includeBlobVersions"`
	IncludeSnapshots    *bool    `pulumi:"includeSnapshots"`
	PrefixMatch         []string `pulumi:"prefixMatch"`
}





type BlobInventoryPolicyFilterInput interface {
	pulumi.Input

	ToBlobInventoryPolicyFilterOutput() BlobInventoryPolicyFilterOutput
	ToBlobInventoryPolicyFilterOutputWithContext(context.Context) BlobInventoryPolicyFilterOutput
}

type BlobInventoryPolicyFilterArgs struct {
	BlobTypes           pulumi.StringArrayInput `pulumi:"blobTypes"`
	IncludeBlobVersions pulumi.BoolPtrInput     `pulumi:"includeBlobVersions"`
	IncludeSnapshots    pulumi.BoolPtrInput     `pulumi:"includeSnapshots"`
	PrefixMatch         pulumi.StringArrayInput `pulumi:"prefixMatch"`
}

func (BlobInventoryPolicyFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyFilter)(nil)).Elem()
}

func (i BlobInventoryPolicyFilterArgs) ToBlobInventoryPolicyFilterOutput() BlobInventoryPolicyFilterOutput {
	return i.ToBlobInventoryPolicyFilterOutputWithContext(context.Background())
}

func (i BlobInventoryPolicyFilterArgs) ToBlobInventoryPolicyFilterOutputWithContext(ctx context.Context) BlobInventoryPolicyFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyFilterOutput)
}

type BlobInventoryPolicyFilterOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicyFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyFilter)(nil)).Elem()
}

func (o BlobInventoryPolicyFilterOutput) ToBlobInventoryPolicyFilterOutput() BlobInventoryPolicyFilterOutput {
	return o
}

func (o BlobInventoryPolicyFilterOutput) ToBlobInventoryPolicyFilterOutputWithContext(ctx context.Context) BlobInventoryPolicyFilterOutput {
	return o
}

func (o BlobInventoryPolicyFilterOutput) BlobTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BlobInventoryPolicyFilter) []string { return v.BlobTypes }).(pulumi.StringArrayOutput)
}

func (o BlobInventoryPolicyFilterOutput) IncludeBlobVersions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BlobInventoryPolicyFilter) *bool { return v.IncludeBlobVersions }).(pulumi.BoolPtrOutput)
}

func (o BlobInventoryPolicyFilterOutput) IncludeSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BlobInventoryPolicyFilter) *bool { return v.IncludeSnapshots }).(pulumi.BoolPtrOutput)
}

func (o BlobInventoryPolicyFilterOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BlobInventoryPolicyFilter) []string { return v.PrefixMatch }).(pulumi.StringArrayOutput)
}

type BlobInventoryPolicyFilterResponse struct {
	BlobTypes           []string `pulumi:"blobTypes"`
	IncludeBlobVersions *bool    `pulumi:"includeBlobVersions"`
	IncludeSnapshots    *bool    `pulumi:"includeSnapshots"`
	PrefixMatch         []string `pulumi:"prefixMatch"`
}





type BlobInventoryPolicyFilterResponseInput interface {
	pulumi.Input

	ToBlobInventoryPolicyFilterResponseOutput() BlobInventoryPolicyFilterResponseOutput
	ToBlobInventoryPolicyFilterResponseOutputWithContext(context.Context) BlobInventoryPolicyFilterResponseOutput
}

type BlobInventoryPolicyFilterResponseArgs struct {
	BlobTypes           pulumi.StringArrayInput `pulumi:"blobTypes"`
	IncludeBlobVersions pulumi.BoolPtrInput     `pulumi:"includeBlobVersions"`
	IncludeSnapshots    pulumi.BoolPtrInput     `pulumi:"includeSnapshots"`
	PrefixMatch         pulumi.StringArrayInput `pulumi:"prefixMatch"`
}

func (BlobInventoryPolicyFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyFilterResponse)(nil)).Elem()
}

func (i BlobInventoryPolicyFilterResponseArgs) ToBlobInventoryPolicyFilterResponseOutput() BlobInventoryPolicyFilterResponseOutput {
	return i.ToBlobInventoryPolicyFilterResponseOutputWithContext(context.Background())
}

func (i BlobInventoryPolicyFilterResponseArgs) ToBlobInventoryPolicyFilterResponseOutputWithContext(ctx context.Context) BlobInventoryPolicyFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyFilterResponseOutput)
}

type BlobInventoryPolicyFilterResponseOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicyFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyFilterResponse)(nil)).Elem()
}

func (o BlobInventoryPolicyFilterResponseOutput) ToBlobInventoryPolicyFilterResponseOutput() BlobInventoryPolicyFilterResponseOutput {
	return o
}

func (o BlobInventoryPolicyFilterResponseOutput) ToBlobInventoryPolicyFilterResponseOutputWithContext(ctx context.Context) BlobInventoryPolicyFilterResponseOutput {
	return o
}

func (o BlobInventoryPolicyFilterResponseOutput) BlobTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BlobInventoryPolicyFilterResponse) []string { return v.BlobTypes }).(pulumi.StringArrayOutput)
}

func (o BlobInventoryPolicyFilterResponseOutput) IncludeBlobVersions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BlobInventoryPolicyFilterResponse) *bool { return v.IncludeBlobVersions }).(pulumi.BoolPtrOutput)
}

func (o BlobInventoryPolicyFilterResponseOutput) IncludeSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BlobInventoryPolicyFilterResponse) *bool { return v.IncludeSnapshots }).(pulumi.BoolPtrOutput)
}

func (o BlobInventoryPolicyFilterResponseOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BlobInventoryPolicyFilterResponse) []string { return v.PrefixMatch }).(pulumi.StringArrayOutput)
}

type BlobInventoryPolicyRule struct {
	Definition BlobInventoryPolicyDefinition `pulumi:"definition"`
	Enabled    bool                          `pulumi:"enabled"`
	Name       string                        `pulumi:"name"`
}





type BlobInventoryPolicyRuleInput interface {
	pulumi.Input

	ToBlobInventoryPolicyRuleOutput() BlobInventoryPolicyRuleOutput
	ToBlobInventoryPolicyRuleOutputWithContext(context.Context) BlobInventoryPolicyRuleOutput
}

type BlobInventoryPolicyRuleArgs struct {
	Definition BlobInventoryPolicyDefinitionInput `pulumi:"definition"`
	Enabled    pulumi.BoolInput                   `pulumi:"enabled"`
	Name       pulumi.StringInput                 `pulumi:"name"`
}

func (BlobInventoryPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyRule)(nil)).Elem()
}

func (i BlobInventoryPolicyRuleArgs) ToBlobInventoryPolicyRuleOutput() BlobInventoryPolicyRuleOutput {
	return i.ToBlobInventoryPolicyRuleOutputWithContext(context.Background())
}

func (i BlobInventoryPolicyRuleArgs) ToBlobInventoryPolicyRuleOutputWithContext(ctx context.Context) BlobInventoryPolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyRuleOutput)
}





type BlobInventoryPolicyRuleArrayInput interface {
	pulumi.Input

	ToBlobInventoryPolicyRuleArrayOutput() BlobInventoryPolicyRuleArrayOutput
	ToBlobInventoryPolicyRuleArrayOutputWithContext(context.Context) BlobInventoryPolicyRuleArrayOutput
}

type BlobInventoryPolicyRuleArray []BlobInventoryPolicyRuleInput

func (BlobInventoryPolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BlobInventoryPolicyRule)(nil)).Elem()
}

func (i BlobInventoryPolicyRuleArray) ToBlobInventoryPolicyRuleArrayOutput() BlobInventoryPolicyRuleArrayOutput {
	return i.ToBlobInventoryPolicyRuleArrayOutputWithContext(context.Background())
}

func (i BlobInventoryPolicyRuleArray) ToBlobInventoryPolicyRuleArrayOutputWithContext(ctx context.Context) BlobInventoryPolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyRuleArrayOutput)
}

type BlobInventoryPolicyRuleOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyRule)(nil)).Elem()
}

func (o BlobInventoryPolicyRuleOutput) ToBlobInventoryPolicyRuleOutput() BlobInventoryPolicyRuleOutput {
	return o
}

func (o BlobInventoryPolicyRuleOutput) ToBlobInventoryPolicyRuleOutputWithContext(ctx context.Context) BlobInventoryPolicyRuleOutput {
	return o
}

func (o BlobInventoryPolicyRuleOutput) Definition() BlobInventoryPolicyDefinitionOutput {
	return o.ApplyT(func(v BlobInventoryPolicyRule) BlobInventoryPolicyDefinition { return v.Definition }).(BlobInventoryPolicyDefinitionOutput)
}

func (o BlobInventoryPolicyRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v BlobInventoryPolicyRule) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o BlobInventoryPolicyRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BlobInventoryPolicyRule) string { return v.Name }).(pulumi.StringOutput)
}

type BlobInventoryPolicyRuleArrayOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BlobInventoryPolicyRule)(nil)).Elem()
}

func (o BlobInventoryPolicyRuleArrayOutput) ToBlobInventoryPolicyRuleArrayOutput() BlobInventoryPolicyRuleArrayOutput {
	return o
}

func (o BlobInventoryPolicyRuleArrayOutput) ToBlobInventoryPolicyRuleArrayOutputWithContext(ctx context.Context) BlobInventoryPolicyRuleArrayOutput {
	return o
}

func (o BlobInventoryPolicyRuleArrayOutput) Index(i pulumi.IntInput) BlobInventoryPolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BlobInventoryPolicyRule {
		return vs[0].([]BlobInventoryPolicyRule)[vs[1].(int)]
	}).(BlobInventoryPolicyRuleOutput)
}

type BlobInventoryPolicyRuleResponse struct {
	Definition BlobInventoryPolicyDefinitionResponse `pulumi:"definition"`
	Enabled    bool                                  `pulumi:"enabled"`
	Name       string                                `pulumi:"name"`
}





type BlobInventoryPolicyRuleResponseInput interface {
	pulumi.Input

	ToBlobInventoryPolicyRuleResponseOutput() BlobInventoryPolicyRuleResponseOutput
	ToBlobInventoryPolicyRuleResponseOutputWithContext(context.Context) BlobInventoryPolicyRuleResponseOutput
}

type BlobInventoryPolicyRuleResponseArgs struct {
	Definition BlobInventoryPolicyDefinitionResponseInput `pulumi:"definition"`
	Enabled    pulumi.BoolInput                           `pulumi:"enabled"`
	Name       pulumi.StringInput                         `pulumi:"name"`
}

func (BlobInventoryPolicyRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyRuleResponse)(nil)).Elem()
}

func (i BlobInventoryPolicyRuleResponseArgs) ToBlobInventoryPolicyRuleResponseOutput() BlobInventoryPolicyRuleResponseOutput {
	return i.ToBlobInventoryPolicyRuleResponseOutputWithContext(context.Background())
}

func (i BlobInventoryPolicyRuleResponseArgs) ToBlobInventoryPolicyRuleResponseOutputWithContext(ctx context.Context) BlobInventoryPolicyRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyRuleResponseOutput)
}





type BlobInventoryPolicyRuleResponseArrayInput interface {
	pulumi.Input

	ToBlobInventoryPolicyRuleResponseArrayOutput() BlobInventoryPolicyRuleResponseArrayOutput
	ToBlobInventoryPolicyRuleResponseArrayOutputWithContext(context.Context) BlobInventoryPolicyRuleResponseArrayOutput
}

type BlobInventoryPolicyRuleResponseArray []BlobInventoryPolicyRuleResponseInput

func (BlobInventoryPolicyRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BlobInventoryPolicyRuleResponse)(nil)).Elem()
}

func (i BlobInventoryPolicyRuleResponseArray) ToBlobInventoryPolicyRuleResponseArrayOutput() BlobInventoryPolicyRuleResponseArrayOutput {
	return i.ToBlobInventoryPolicyRuleResponseArrayOutputWithContext(context.Background())
}

func (i BlobInventoryPolicyRuleResponseArray) ToBlobInventoryPolicyRuleResponseArrayOutputWithContext(ctx context.Context) BlobInventoryPolicyRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyRuleResponseArrayOutput)
}

type BlobInventoryPolicyRuleResponseOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicyRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicyRuleResponse)(nil)).Elem()
}

func (o BlobInventoryPolicyRuleResponseOutput) ToBlobInventoryPolicyRuleResponseOutput() BlobInventoryPolicyRuleResponseOutput {
	return o
}

func (o BlobInventoryPolicyRuleResponseOutput) ToBlobInventoryPolicyRuleResponseOutputWithContext(ctx context.Context) BlobInventoryPolicyRuleResponseOutput {
	return o
}

func (o BlobInventoryPolicyRuleResponseOutput) Definition() BlobInventoryPolicyDefinitionResponseOutput {
	return o.ApplyT(func(v BlobInventoryPolicyRuleResponse) BlobInventoryPolicyDefinitionResponse { return v.Definition }).(BlobInventoryPolicyDefinitionResponseOutput)
}

func (o BlobInventoryPolicyRuleResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v BlobInventoryPolicyRuleResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o BlobInventoryPolicyRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BlobInventoryPolicyRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

type BlobInventoryPolicyRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicyRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BlobInventoryPolicyRuleResponse)(nil)).Elem()
}

func (o BlobInventoryPolicyRuleResponseArrayOutput) ToBlobInventoryPolicyRuleResponseArrayOutput() BlobInventoryPolicyRuleResponseArrayOutput {
	return o
}

func (o BlobInventoryPolicyRuleResponseArrayOutput) ToBlobInventoryPolicyRuleResponseArrayOutputWithContext(ctx context.Context) BlobInventoryPolicyRuleResponseArrayOutput {
	return o
}

func (o BlobInventoryPolicyRuleResponseArrayOutput) Index(i pulumi.IntInput) BlobInventoryPolicyRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BlobInventoryPolicyRuleResponse {
		return vs[0].([]BlobInventoryPolicyRuleResponse)[vs[1].(int)]
	}).(BlobInventoryPolicyRuleResponseOutput)
}

type BlobInventoryPolicySchema struct {
	Destination string                    `pulumi:"destination"`
	Enabled     bool                      `pulumi:"enabled"`
	Rules       []BlobInventoryPolicyRule `pulumi:"rules"`
	Type        string                    `pulumi:"type"`
}





type BlobInventoryPolicySchemaInput interface {
	pulumi.Input

	ToBlobInventoryPolicySchemaOutput() BlobInventoryPolicySchemaOutput
	ToBlobInventoryPolicySchemaOutputWithContext(context.Context) BlobInventoryPolicySchemaOutput
}

type BlobInventoryPolicySchemaArgs struct {
	Destination pulumi.StringInput                `pulumi:"destination"`
	Enabled     pulumi.BoolInput                  `pulumi:"enabled"`
	Rules       BlobInventoryPolicyRuleArrayInput `pulumi:"rules"`
	Type        pulumi.StringInput                `pulumi:"type"`
}

func (BlobInventoryPolicySchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicySchema)(nil)).Elem()
}

func (i BlobInventoryPolicySchemaArgs) ToBlobInventoryPolicySchemaOutput() BlobInventoryPolicySchemaOutput {
	return i.ToBlobInventoryPolicySchemaOutputWithContext(context.Background())
}

func (i BlobInventoryPolicySchemaArgs) ToBlobInventoryPolicySchemaOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicySchemaOutput)
}

func (i BlobInventoryPolicySchemaArgs) ToBlobInventoryPolicySchemaPtrOutput() BlobInventoryPolicySchemaPtrOutput {
	return i.ToBlobInventoryPolicySchemaPtrOutputWithContext(context.Background())
}

func (i BlobInventoryPolicySchemaArgs) ToBlobInventoryPolicySchemaPtrOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicySchemaOutput).ToBlobInventoryPolicySchemaPtrOutputWithContext(ctx)
}









type BlobInventoryPolicySchemaPtrInput interface {
	pulumi.Input

	ToBlobInventoryPolicySchemaPtrOutput() BlobInventoryPolicySchemaPtrOutput
	ToBlobInventoryPolicySchemaPtrOutputWithContext(context.Context) BlobInventoryPolicySchemaPtrOutput
}

type blobInventoryPolicySchemaPtrType BlobInventoryPolicySchemaArgs

func BlobInventoryPolicySchemaPtr(v *BlobInventoryPolicySchemaArgs) BlobInventoryPolicySchemaPtrInput {
	return (*blobInventoryPolicySchemaPtrType)(v)
}

func (*blobInventoryPolicySchemaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobInventoryPolicySchema)(nil)).Elem()
}

func (i *blobInventoryPolicySchemaPtrType) ToBlobInventoryPolicySchemaPtrOutput() BlobInventoryPolicySchemaPtrOutput {
	return i.ToBlobInventoryPolicySchemaPtrOutputWithContext(context.Background())
}

func (i *blobInventoryPolicySchemaPtrType) ToBlobInventoryPolicySchemaPtrOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicySchemaPtrOutput)
}

type BlobInventoryPolicySchemaOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicySchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicySchema)(nil)).Elem()
}

func (o BlobInventoryPolicySchemaOutput) ToBlobInventoryPolicySchemaOutput() BlobInventoryPolicySchemaOutput {
	return o
}

func (o BlobInventoryPolicySchemaOutput) ToBlobInventoryPolicySchemaOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaOutput {
	return o
}

func (o BlobInventoryPolicySchemaOutput) ToBlobInventoryPolicySchemaPtrOutput() BlobInventoryPolicySchemaPtrOutput {
	return o.ToBlobInventoryPolicySchemaPtrOutputWithContext(context.Background())
}

func (o BlobInventoryPolicySchemaOutput) ToBlobInventoryPolicySchemaPtrOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobInventoryPolicySchema) *BlobInventoryPolicySchema {
		return &v
	}).(BlobInventoryPolicySchemaPtrOutput)
}

func (o BlobInventoryPolicySchemaOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v BlobInventoryPolicySchema) string { return v.Destination }).(pulumi.StringOutput)
}

func (o BlobInventoryPolicySchemaOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v BlobInventoryPolicySchema) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o BlobInventoryPolicySchemaOutput) Rules() BlobInventoryPolicyRuleArrayOutput {
	return o.ApplyT(func(v BlobInventoryPolicySchema) []BlobInventoryPolicyRule { return v.Rules }).(BlobInventoryPolicyRuleArrayOutput)
}

func (o BlobInventoryPolicySchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BlobInventoryPolicySchema) string { return v.Type }).(pulumi.StringOutput)
}

type BlobInventoryPolicySchemaPtrOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicySchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobInventoryPolicySchema)(nil)).Elem()
}

func (o BlobInventoryPolicySchemaPtrOutput) ToBlobInventoryPolicySchemaPtrOutput() BlobInventoryPolicySchemaPtrOutput {
	return o
}

func (o BlobInventoryPolicySchemaPtrOutput) ToBlobInventoryPolicySchemaPtrOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaPtrOutput {
	return o
}

func (o BlobInventoryPolicySchemaPtrOutput) Elem() BlobInventoryPolicySchemaOutput {
	return o.ApplyT(func(v *BlobInventoryPolicySchema) BlobInventoryPolicySchema {
		if v != nil {
			return *v
		}
		var ret BlobInventoryPolicySchema
		return ret
	}).(BlobInventoryPolicySchemaOutput)
}

func (o BlobInventoryPolicySchemaPtrOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobInventoryPolicySchema) *string {
		if v == nil {
			return nil
		}
		return &v.Destination
	}).(pulumi.StringPtrOutput)
}

func (o BlobInventoryPolicySchemaPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlobInventoryPolicySchema) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o BlobInventoryPolicySchemaPtrOutput) Rules() BlobInventoryPolicyRuleArrayOutput {
	return o.ApplyT(func(v *BlobInventoryPolicySchema) []BlobInventoryPolicyRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(BlobInventoryPolicyRuleArrayOutput)
}

func (o BlobInventoryPolicySchemaPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobInventoryPolicySchema) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type BlobInventoryPolicySchemaResponse struct {
	Destination string                            `pulumi:"destination"`
	Enabled     bool                              `pulumi:"enabled"`
	Rules       []BlobInventoryPolicyRuleResponse `pulumi:"rules"`
	Type        string                            `pulumi:"type"`
}





type BlobInventoryPolicySchemaResponseInput interface {
	pulumi.Input

	ToBlobInventoryPolicySchemaResponseOutput() BlobInventoryPolicySchemaResponseOutput
	ToBlobInventoryPolicySchemaResponseOutputWithContext(context.Context) BlobInventoryPolicySchemaResponseOutput
}

type BlobInventoryPolicySchemaResponseArgs struct {
	Destination pulumi.StringInput                        `pulumi:"destination"`
	Enabled     pulumi.BoolInput                          `pulumi:"enabled"`
	Rules       BlobInventoryPolicyRuleResponseArrayInput `pulumi:"rules"`
	Type        pulumi.StringInput                        `pulumi:"type"`
}

func (BlobInventoryPolicySchemaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicySchemaResponse)(nil)).Elem()
}

func (i BlobInventoryPolicySchemaResponseArgs) ToBlobInventoryPolicySchemaResponseOutput() BlobInventoryPolicySchemaResponseOutput {
	return i.ToBlobInventoryPolicySchemaResponseOutputWithContext(context.Background())
}

func (i BlobInventoryPolicySchemaResponseArgs) ToBlobInventoryPolicySchemaResponseOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicySchemaResponseOutput)
}

func (i BlobInventoryPolicySchemaResponseArgs) ToBlobInventoryPolicySchemaResponsePtrOutput() BlobInventoryPolicySchemaResponsePtrOutput {
	return i.ToBlobInventoryPolicySchemaResponsePtrOutputWithContext(context.Background())
}

func (i BlobInventoryPolicySchemaResponseArgs) ToBlobInventoryPolicySchemaResponsePtrOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicySchemaResponseOutput).ToBlobInventoryPolicySchemaResponsePtrOutputWithContext(ctx)
}









type BlobInventoryPolicySchemaResponsePtrInput interface {
	pulumi.Input

	ToBlobInventoryPolicySchemaResponsePtrOutput() BlobInventoryPolicySchemaResponsePtrOutput
	ToBlobInventoryPolicySchemaResponsePtrOutputWithContext(context.Context) BlobInventoryPolicySchemaResponsePtrOutput
}

type blobInventoryPolicySchemaResponsePtrType BlobInventoryPolicySchemaResponseArgs

func BlobInventoryPolicySchemaResponsePtr(v *BlobInventoryPolicySchemaResponseArgs) BlobInventoryPolicySchemaResponsePtrInput {
	return (*blobInventoryPolicySchemaResponsePtrType)(v)
}

func (*blobInventoryPolicySchemaResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobInventoryPolicySchemaResponse)(nil)).Elem()
}

func (i *blobInventoryPolicySchemaResponsePtrType) ToBlobInventoryPolicySchemaResponsePtrOutput() BlobInventoryPolicySchemaResponsePtrOutput {
	return i.ToBlobInventoryPolicySchemaResponsePtrOutputWithContext(context.Background())
}

func (i *blobInventoryPolicySchemaResponsePtrType) ToBlobInventoryPolicySchemaResponsePtrOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicySchemaResponsePtrOutput)
}

type BlobInventoryPolicySchemaResponseOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicySchemaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicySchemaResponse)(nil)).Elem()
}

func (o BlobInventoryPolicySchemaResponseOutput) ToBlobInventoryPolicySchemaResponseOutput() BlobInventoryPolicySchemaResponseOutput {
	return o
}

func (o BlobInventoryPolicySchemaResponseOutput) ToBlobInventoryPolicySchemaResponseOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaResponseOutput {
	return o
}

func (o BlobInventoryPolicySchemaResponseOutput) ToBlobInventoryPolicySchemaResponsePtrOutput() BlobInventoryPolicySchemaResponsePtrOutput {
	return o.ToBlobInventoryPolicySchemaResponsePtrOutputWithContext(context.Background())
}

func (o BlobInventoryPolicySchemaResponseOutput) ToBlobInventoryPolicySchemaResponsePtrOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobInventoryPolicySchemaResponse) *BlobInventoryPolicySchemaResponse {
		return &v
	}).(BlobInventoryPolicySchemaResponsePtrOutput)
}

func (o BlobInventoryPolicySchemaResponseOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v BlobInventoryPolicySchemaResponse) string { return v.Destination }).(pulumi.StringOutput)
}

func (o BlobInventoryPolicySchemaResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v BlobInventoryPolicySchemaResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o BlobInventoryPolicySchemaResponseOutput) Rules() BlobInventoryPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v BlobInventoryPolicySchemaResponse) []BlobInventoryPolicyRuleResponse { return v.Rules }).(BlobInventoryPolicyRuleResponseArrayOutput)
}

func (o BlobInventoryPolicySchemaResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BlobInventoryPolicySchemaResponse) string { return v.Type }).(pulumi.StringOutput)
}

type BlobInventoryPolicySchemaResponsePtrOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicySchemaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobInventoryPolicySchemaResponse)(nil)).Elem()
}

func (o BlobInventoryPolicySchemaResponsePtrOutput) ToBlobInventoryPolicySchemaResponsePtrOutput() BlobInventoryPolicySchemaResponsePtrOutput {
	return o
}

func (o BlobInventoryPolicySchemaResponsePtrOutput) ToBlobInventoryPolicySchemaResponsePtrOutputWithContext(ctx context.Context) BlobInventoryPolicySchemaResponsePtrOutput {
	return o
}

func (o BlobInventoryPolicySchemaResponsePtrOutput) Elem() BlobInventoryPolicySchemaResponseOutput {
	return o.ApplyT(func(v *BlobInventoryPolicySchemaResponse) BlobInventoryPolicySchemaResponse {
		if v != nil {
			return *v
		}
		var ret BlobInventoryPolicySchemaResponse
		return ret
	}).(BlobInventoryPolicySchemaResponseOutput)
}

func (o BlobInventoryPolicySchemaResponsePtrOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobInventoryPolicySchemaResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Destination
	}).(pulumi.StringPtrOutput)
}

func (o BlobInventoryPolicySchemaResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlobInventoryPolicySchemaResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o BlobInventoryPolicySchemaResponsePtrOutput) Rules() BlobInventoryPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v *BlobInventoryPolicySchemaResponse) []BlobInventoryPolicyRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(BlobInventoryPolicyRuleResponseArrayOutput)
}

func (o BlobInventoryPolicySchemaResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobInventoryPolicySchemaResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type BlobRestoreParametersResponse struct {
	BlobRanges    []BlobRestoreRangeResponse `pulumi:"blobRanges"`
	TimeToRestore string                     `pulumi:"timeToRestore"`
}





type BlobRestoreParametersResponseInput interface {
	pulumi.Input

	ToBlobRestoreParametersResponseOutput() BlobRestoreParametersResponseOutput
	ToBlobRestoreParametersResponseOutputWithContext(context.Context) BlobRestoreParametersResponseOutput
}

type BlobRestoreParametersResponseArgs struct {
	BlobRanges    BlobRestoreRangeResponseArrayInput `pulumi:"blobRanges"`
	TimeToRestore pulumi.StringInput                 `pulumi:"timeToRestore"`
}

func (BlobRestoreParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobRestoreParametersResponse)(nil)).Elem()
}

func (i BlobRestoreParametersResponseArgs) ToBlobRestoreParametersResponseOutput() BlobRestoreParametersResponseOutput {
	return i.ToBlobRestoreParametersResponseOutputWithContext(context.Background())
}

func (i BlobRestoreParametersResponseArgs) ToBlobRestoreParametersResponseOutputWithContext(ctx context.Context) BlobRestoreParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobRestoreParametersResponseOutput)
}

func (i BlobRestoreParametersResponseArgs) ToBlobRestoreParametersResponsePtrOutput() BlobRestoreParametersResponsePtrOutput {
	return i.ToBlobRestoreParametersResponsePtrOutputWithContext(context.Background())
}

func (i BlobRestoreParametersResponseArgs) ToBlobRestoreParametersResponsePtrOutputWithContext(ctx context.Context) BlobRestoreParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobRestoreParametersResponseOutput).ToBlobRestoreParametersResponsePtrOutputWithContext(ctx)
}









type BlobRestoreParametersResponsePtrInput interface {
	pulumi.Input

	ToBlobRestoreParametersResponsePtrOutput() BlobRestoreParametersResponsePtrOutput
	ToBlobRestoreParametersResponsePtrOutputWithContext(context.Context) BlobRestoreParametersResponsePtrOutput
}

type blobRestoreParametersResponsePtrType BlobRestoreParametersResponseArgs

func BlobRestoreParametersResponsePtr(v *BlobRestoreParametersResponseArgs) BlobRestoreParametersResponsePtrInput {
	return (*blobRestoreParametersResponsePtrType)(v)
}

func (*blobRestoreParametersResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobRestoreParametersResponse)(nil)).Elem()
}

func (i *blobRestoreParametersResponsePtrType) ToBlobRestoreParametersResponsePtrOutput() BlobRestoreParametersResponsePtrOutput {
	return i.ToBlobRestoreParametersResponsePtrOutputWithContext(context.Background())
}

func (i *blobRestoreParametersResponsePtrType) ToBlobRestoreParametersResponsePtrOutputWithContext(ctx context.Context) BlobRestoreParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobRestoreParametersResponsePtrOutput)
}

type BlobRestoreParametersResponseOutput struct{ *pulumi.OutputState }

func (BlobRestoreParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobRestoreParametersResponse)(nil)).Elem()
}

func (o BlobRestoreParametersResponseOutput) ToBlobRestoreParametersResponseOutput() BlobRestoreParametersResponseOutput {
	return o
}

func (o BlobRestoreParametersResponseOutput) ToBlobRestoreParametersResponseOutputWithContext(ctx context.Context) BlobRestoreParametersResponseOutput {
	return o
}

func (o BlobRestoreParametersResponseOutput) ToBlobRestoreParametersResponsePtrOutput() BlobRestoreParametersResponsePtrOutput {
	return o.ToBlobRestoreParametersResponsePtrOutputWithContext(context.Background())
}

func (o BlobRestoreParametersResponseOutput) ToBlobRestoreParametersResponsePtrOutputWithContext(ctx context.Context) BlobRestoreParametersResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobRestoreParametersResponse) *BlobRestoreParametersResponse {
		return &v
	}).(BlobRestoreParametersResponsePtrOutput)
}

func (o BlobRestoreParametersResponseOutput) BlobRanges() BlobRestoreRangeResponseArrayOutput {
	return o.ApplyT(func(v BlobRestoreParametersResponse) []BlobRestoreRangeResponse { return v.BlobRanges }).(BlobRestoreRangeResponseArrayOutput)
}

func (o BlobRestoreParametersResponseOutput) TimeToRestore() pulumi.StringOutput {
	return o.ApplyT(func(v BlobRestoreParametersResponse) string { return v.TimeToRestore }).(pulumi.StringOutput)
}

type BlobRestoreParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (BlobRestoreParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobRestoreParametersResponse)(nil)).Elem()
}

func (o BlobRestoreParametersResponsePtrOutput) ToBlobRestoreParametersResponsePtrOutput() BlobRestoreParametersResponsePtrOutput {
	return o
}

func (o BlobRestoreParametersResponsePtrOutput) ToBlobRestoreParametersResponsePtrOutputWithContext(ctx context.Context) BlobRestoreParametersResponsePtrOutput {
	return o
}

func (o BlobRestoreParametersResponsePtrOutput) Elem() BlobRestoreParametersResponseOutput {
	return o.ApplyT(func(v *BlobRestoreParametersResponse) BlobRestoreParametersResponse {
		if v != nil {
			return *v
		}
		var ret BlobRestoreParametersResponse
		return ret
	}).(BlobRestoreParametersResponseOutput)
}

func (o BlobRestoreParametersResponsePtrOutput) BlobRanges() BlobRestoreRangeResponseArrayOutput {
	return o.ApplyT(func(v *BlobRestoreParametersResponse) []BlobRestoreRangeResponse {
		if v == nil {
			return nil
		}
		return v.BlobRanges
	}).(BlobRestoreRangeResponseArrayOutput)
}

func (o BlobRestoreParametersResponsePtrOutput) TimeToRestore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobRestoreParametersResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TimeToRestore
	}).(pulumi.StringPtrOutput)
}

type BlobRestoreRangeResponse struct {
	EndRange   string `pulumi:"endRange"`
	StartRange string `pulumi:"startRange"`
}





type BlobRestoreRangeResponseInput interface {
	pulumi.Input

	ToBlobRestoreRangeResponseOutput() BlobRestoreRangeResponseOutput
	ToBlobRestoreRangeResponseOutputWithContext(context.Context) BlobRestoreRangeResponseOutput
}

type BlobRestoreRangeResponseArgs struct {
	EndRange   pulumi.StringInput `pulumi:"endRange"`
	StartRange pulumi.StringInput `pulumi:"startRange"`
}

func (BlobRestoreRangeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobRestoreRangeResponse)(nil)).Elem()
}

func (i BlobRestoreRangeResponseArgs) ToBlobRestoreRangeResponseOutput() BlobRestoreRangeResponseOutput {
	return i.ToBlobRestoreRangeResponseOutputWithContext(context.Background())
}

func (i BlobRestoreRangeResponseArgs) ToBlobRestoreRangeResponseOutputWithContext(ctx context.Context) BlobRestoreRangeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobRestoreRangeResponseOutput)
}





type BlobRestoreRangeResponseArrayInput interface {
	pulumi.Input

	ToBlobRestoreRangeResponseArrayOutput() BlobRestoreRangeResponseArrayOutput
	ToBlobRestoreRangeResponseArrayOutputWithContext(context.Context) BlobRestoreRangeResponseArrayOutput
}

type BlobRestoreRangeResponseArray []BlobRestoreRangeResponseInput

func (BlobRestoreRangeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BlobRestoreRangeResponse)(nil)).Elem()
}

func (i BlobRestoreRangeResponseArray) ToBlobRestoreRangeResponseArrayOutput() BlobRestoreRangeResponseArrayOutput {
	return i.ToBlobRestoreRangeResponseArrayOutputWithContext(context.Background())
}

func (i BlobRestoreRangeResponseArray) ToBlobRestoreRangeResponseArrayOutputWithContext(ctx context.Context) BlobRestoreRangeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobRestoreRangeResponseArrayOutput)
}

type BlobRestoreRangeResponseOutput struct{ *pulumi.OutputState }

func (BlobRestoreRangeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobRestoreRangeResponse)(nil)).Elem()
}

func (o BlobRestoreRangeResponseOutput) ToBlobRestoreRangeResponseOutput() BlobRestoreRangeResponseOutput {
	return o
}

func (o BlobRestoreRangeResponseOutput) ToBlobRestoreRangeResponseOutputWithContext(ctx context.Context) BlobRestoreRangeResponseOutput {
	return o
}

func (o BlobRestoreRangeResponseOutput) EndRange() pulumi.StringOutput {
	return o.ApplyT(func(v BlobRestoreRangeResponse) string { return v.EndRange }).(pulumi.StringOutput)
}

func (o BlobRestoreRangeResponseOutput) StartRange() pulumi.StringOutput {
	return o.ApplyT(func(v BlobRestoreRangeResponse) string { return v.StartRange }).(pulumi.StringOutput)
}

type BlobRestoreRangeResponseArrayOutput struct{ *pulumi.OutputState }

func (BlobRestoreRangeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BlobRestoreRangeResponse)(nil)).Elem()
}

func (o BlobRestoreRangeResponseArrayOutput) ToBlobRestoreRangeResponseArrayOutput() BlobRestoreRangeResponseArrayOutput {
	return o
}

func (o BlobRestoreRangeResponseArrayOutput) ToBlobRestoreRangeResponseArrayOutputWithContext(ctx context.Context) BlobRestoreRangeResponseArrayOutput {
	return o
}

func (o BlobRestoreRangeResponseArrayOutput) Index(i pulumi.IntInput) BlobRestoreRangeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BlobRestoreRangeResponse {
		return vs[0].([]BlobRestoreRangeResponse)[vs[1].(int)]
	}).(BlobRestoreRangeResponseOutput)
}

type BlobRestoreStatusResponse struct {
	FailureReason string                        `pulumi:"failureReason"`
	Parameters    BlobRestoreParametersResponse `pulumi:"parameters"`
	RestoreId     string                        `pulumi:"restoreId"`
	Status        string                        `pulumi:"status"`
}





type BlobRestoreStatusResponseInput interface {
	pulumi.Input

	ToBlobRestoreStatusResponseOutput() BlobRestoreStatusResponseOutput
	ToBlobRestoreStatusResponseOutputWithContext(context.Context) BlobRestoreStatusResponseOutput
}

type BlobRestoreStatusResponseArgs struct {
	FailureReason pulumi.StringInput                 `pulumi:"failureReason"`
	Parameters    BlobRestoreParametersResponseInput `pulumi:"parameters"`
	RestoreId     pulumi.StringInput                 `pulumi:"restoreId"`
	Status        pulumi.StringInput                 `pulumi:"status"`
}

func (BlobRestoreStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobRestoreStatusResponse)(nil)).Elem()
}

func (i BlobRestoreStatusResponseArgs) ToBlobRestoreStatusResponseOutput() BlobRestoreStatusResponseOutput {
	return i.ToBlobRestoreStatusResponseOutputWithContext(context.Background())
}

func (i BlobRestoreStatusResponseArgs) ToBlobRestoreStatusResponseOutputWithContext(ctx context.Context) BlobRestoreStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobRestoreStatusResponseOutput)
}

func (i BlobRestoreStatusResponseArgs) ToBlobRestoreStatusResponsePtrOutput() BlobRestoreStatusResponsePtrOutput {
	return i.ToBlobRestoreStatusResponsePtrOutputWithContext(context.Background())
}

func (i BlobRestoreStatusResponseArgs) ToBlobRestoreStatusResponsePtrOutputWithContext(ctx context.Context) BlobRestoreStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobRestoreStatusResponseOutput).ToBlobRestoreStatusResponsePtrOutputWithContext(ctx)
}









type BlobRestoreStatusResponsePtrInput interface {
	pulumi.Input

	ToBlobRestoreStatusResponsePtrOutput() BlobRestoreStatusResponsePtrOutput
	ToBlobRestoreStatusResponsePtrOutputWithContext(context.Context) BlobRestoreStatusResponsePtrOutput
}

type blobRestoreStatusResponsePtrType BlobRestoreStatusResponseArgs

func BlobRestoreStatusResponsePtr(v *BlobRestoreStatusResponseArgs) BlobRestoreStatusResponsePtrInput {
	return (*blobRestoreStatusResponsePtrType)(v)
}

func (*blobRestoreStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobRestoreStatusResponse)(nil)).Elem()
}

func (i *blobRestoreStatusResponsePtrType) ToBlobRestoreStatusResponsePtrOutput() BlobRestoreStatusResponsePtrOutput {
	return i.ToBlobRestoreStatusResponsePtrOutputWithContext(context.Background())
}

func (i *blobRestoreStatusResponsePtrType) ToBlobRestoreStatusResponsePtrOutputWithContext(ctx context.Context) BlobRestoreStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobRestoreStatusResponsePtrOutput)
}

type BlobRestoreStatusResponseOutput struct{ *pulumi.OutputState }

func (BlobRestoreStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobRestoreStatusResponse)(nil)).Elem()
}

func (o BlobRestoreStatusResponseOutput) ToBlobRestoreStatusResponseOutput() BlobRestoreStatusResponseOutput {
	return o
}

func (o BlobRestoreStatusResponseOutput) ToBlobRestoreStatusResponseOutputWithContext(ctx context.Context) BlobRestoreStatusResponseOutput {
	return o
}

func (o BlobRestoreStatusResponseOutput) ToBlobRestoreStatusResponsePtrOutput() BlobRestoreStatusResponsePtrOutput {
	return o.ToBlobRestoreStatusResponsePtrOutputWithContext(context.Background())
}

func (o BlobRestoreStatusResponseOutput) ToBlobRestoreStatusResponsePtrOutputWithContext(ctx context.Context) BlobRestoreStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobRestoreStatusResponse) *BlobRestoreStatusResponse {
		return &v
	}).(BlobRestoreStatusResponsePtrOutput)
}

func (o BlobRestoreStatusResponseOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v BlobRestoreStatusResponse) string { return v.FailureReason }).(pulumi.StringOutput)
}

func (o BlobRestoreStatusResponseOutput) Parameters() BlobRestoreParametersResponseOutput {
	return o.ApplyT(func(v BlobRestoreStatusResponse) BlobRestoreParametersResponse { return v.Parameters }).(BlobRestoreParametersResponseOutput)
}

func (o BlobRestoreStatusResponseOutput) RestoreId() pulumi.StringOutput {
	return o.ApplyT(func(v BlobRestoreStatusResponse) string { return v.RestoreId }).(pulumi.StringOutput)
}

func (o BlobRestoreStatusResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v BlobRestoreStatusResponse) string { return v.Status }).(pulumi.StringOutput)
}

type BlobRestoreStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (BlobRestoreStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobRestoreStatusResponse)(nil)).Elem()
}

func (o BlobRestoreStatusResponsePtrOutput) ToBlobRestoreStatusResponsePtrOutput() BlobRestoreStatusResponsePtrOutput {
	return o
}

func (o BlobRestoreStatusResponsePtrOutput) ToBlobRestoreStatusResponsePtrOutputWithContext(ctx context.Context) BlobRestoreStatusResponsePtrOutput {
	return o
}

func (o BlobRestoreStatusResponsePtrOutput) Elem() BlobRestoreStatusResponseOutput {
	return o.ApplyT(func(v *BlobRestoreStatusResponse) BlobRestoreStatusResponse {
		if v != nil {
			return *v
		}
		var ret BlobRestoreStatusResponse
		return ret
	}).(BlobRestoreStatusResponseOutput)
}

func (o BlobRestoreStatusResponsePtrOutput) FailureReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobRestoreStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FailureReason
	}).(pulumi.StringPtrOutput)
}

func (o BlobRestoreStatusResponsePtrOutput) Parameters() BlobRestoreParametersResponsePtrOutput {
	return o.ApplyT(func(v *BlobRestoreStatusResponse) *BlobRestoreParametersResponse {
		if v == nil {
			return nil
		}
		return &v.Parameters
	}).(BlobRestoreParametersResponsePtrOutput)
}

func (o BlobRestoreStatusResponsePtrOutput) RestoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobRestoreStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RestoreId
	}).(pulumi.StringPtrOutput)
}

func (o BlobRestoreStatusResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobRestoreStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type ChangeFeed struct {
	Enabled         *bool `pulumi:"enabled"`
	RetentionInDays *int  `pulumi:"retentionInDays"`
}





type ChangeFeedInput interface {
	pulumi.Input

	ToChangeFeedOutput() ChangeFeedOutput
	ToChangeFeedOutputWithContext(context.Context) ChangeFeedOutput
}

type ChangeFeedArgs struct {
	Enabled         pulumi.BoolPtrInput `pulumi:"enabled"`
	RetentionInDays pulumi.IntPtrInput  `pulumi:"retentionInDays"`
}

func (ChangeFeedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ChangeFeed)(nil)).Elem()
}

func (i ChangeFeedArgs) ToChangeFeedOutput() ChangeFeedOutput {
	return i.ToChangeFeedOutputWithContext(context.Background())
}

func (i ChangeFeedArgs) ToChangeFeedOutputWithContext(ctx context.Context) ChangeFeedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChangeFeedOutput)
}

func (i ChangeFeedArgs) ToChangeFeedPtrOutput() ChangeFeedPtrOutput {
	return i.ToChangeFeedPtrOutputWithContext(context.Background())
}

func (i ChangeFeedArgs) ToChangeFeedPtrOutputWithContext(ctx context.Context) ChangeFeedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChangeFeedOutput).ToChangeFeedPtrOutputWithContext(ctx)
}









type ChangeFeedPtrInput interface {
	pulumi.Input

	ToChangeFeedPtrOutput() ChangeFeedPtrOutput
	ToChangeFeedPtrOutputWithContext(context.Context) ChangeFeedPtrOutput
}

type changeFeedPtrType ChangeFeedArgs

func ChangeFeedPtr(v *ChangeFeedArgs) ChangeFeedPtrInput {
	return (*changeFeedPtrType)(v)
}

func (*changeFeedPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ChangeFeed)(nil)).Elem()
}

func (i *changeFeedPtrType) ToChangeFeedPtrOutput() ChangeFeedPtrOutput {
	return i.ToChangeFeedPtrOutputWithContext(context.Background())
}

func (i *changeFeedPtrType) ToChangeFeedPtrOutputWithContext(ctx context.Context) ChangeFeedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChangeFeedPtrOutput)
}

type ChangeFeedOutput struct{ *pulumi.OutputState }

func (ChangeFeedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChangeFeed)(nil)).Elem()
}

func (o ChangeFeedOutput) ToChangeFeedOutput() ChangeFeedOutput {
	return o
}

func (o ChangeFeedOutput) ToChangeFeedOutputWithContext(ctx context.Context) ChangeFeedOutput {
	return o
}

func (o ChangeFeedOutput) ToChangeFeedPtrOutput() ChangeFeedPtrOutput {
	return o.ToChangeFeedPtrOutputWithContext(context.Background())
}

func (o ChangeFeedOutput) ToChangeFeedPtrOutputWithContext(ctx context.Context) ChangeFeedPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ChangeFeed) *ChangeFeed {
		return &v
	}).(ChangeFeedPtrOutput)
}

func (o ChangeFeedOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ChangeFeed) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ChangeFeedOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ChangeFeed) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

type ChangeFeedPtrOutput struct{ *pulumi.OutputState }

func (ChangeFeedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChangeFeed)(nil)).Elem()
}

func (o ChangeFeedPtrOutput) ToChangeFeedPtrOutput() ChangeFeedPtrOutput {
	return o
}

func (o ChangeFeedPtrOutput) ToChangeFeedPtrOutputWithContext(ctx context.Context) ChangeFeedPtrOutput {
	return o
}

func (o ChangeFeedPtrOutput) Elem() ChangeFeedOutput {
	return o.ApplyT(func(v *ChangeFeed) ChangeFeed {
		if v != nil {
			return *v
		}
		var ret ChangeFeed
		return ret
	}).(ChangeFeedOutput)
}

func (o ChangeFeedPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ChangeFeed) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ChangeFeedPtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ChangeFeed) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

type ChangeFeedResponse struct {
	Enabled         *bool `pulumi:"enabled"`
	RetentionInDays *int  `pulumi:"retentionInDays"`
}





type ChangeFeedResponseInput interface {
	pulumi.Input

	ToChangeFeedResponseOutput() ChangeFeedResponseOutput
	ToChangeFeedResponseOutputWithContext(context.Context) ChangeFeedResponseOutput
}

type ChangeFeedResponseArgs struct {
	Enabled         pulumi.BoolPtrInput `pulumi:"enabled"`
	RetentionInDays pulumi.IntPtrInput  `pulumi:"retentionInDays"`
}

func (ChangeFeedResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ChangeFeedResponse)(nil)).Elem()
}

func (i ChangeFeedResponseArgs) ToChangeFeedResponseOutput() ChangeFeedResponseOutput {
	return i.ToChangeFeedResponseOutputWithContext(context.Background())
}

func (i ChangeFeedResponseArgs) ToChangeFeedResponseOutputWithContext(ctx context.Context) ChangeFeedResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChangeFeedResponseOutput)
}

func (i ChangeFeedResponseArgs) ToChangeFeedResponsePtrOutput() ChangeFeedResponsePtrOutput {
	return i.ToChangeFeedResponsePtrOutputWithContext(context.Background())
}

func (i ChangeFeedResponseArgs) ToChangeFeedResponsePtrOutputWithContext(ctx context.Context) ChangeFeedResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChangeFeedResponseOutput).ToChangeFeedResponsePtrOutputWithContext(ctx)
}









type ChangeFeedResponsePtrInput interface {
	pulumi.Input

	ToChangeFeedResponsePtrOutput() ChangeFeedResponsePtrOutput
	ToChangeFeedResponsePtrOutputWithContext(context.Context) ChangeFeedResponsePtrOutput
}

type changeFeedResponsePtrType ChangeFeedResponseArgs

func ChangeFeedResponsePtr(v *ChangeFeedResponseArgs) ChangeFeedResponsePtrInput {
	return (*changeFeedResponsePtrType)(v)
}

func (*changeFeedResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ChangeFeedResponse)(nil)).Elem()
}

func (i *changeFeedResponsePtrType) ToChangeFeedResponsePtrOutput() ChangeFeedResponsePtrOutput {
	return i.ToChangeFeedResponsePtrOutputWithContext(context.Background())
}

func (i *changeFeedResponsePtrType) ToChangeFeedResponsePtrOutputWithContext(ctx context.Context) ChangeFeedResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChangeFeedResponsePtrOutput)
}

type ChangeFeedResponseOutput struct{ *pulumi.OutputState }

func (ChangeFeedResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChangeFeedResponse)(nil)).Elem()
}

func (o ChangeFeedResponseOutput) ToChangeFeedResponseOutput() ChangeFeedResponseOutput {
	return o
}

func (o ChangeFeedResponseOutput) ToChangeFeedResponseOutputWithContext(ctx context.Context) ChangeFeedResponseOutput {
	return o
}

func (o ChangeFeedResponseOutput) ToChangeFeedResponsePtrOutput() ChangeFeedResponsePtrOutput {
	return o.ToChangeFeedResponsePtrOutputWithContext(context.Background())
}

func (o ChangeFeedResponseOutput) ToChangeFeedResponsePtrOutputWithContext(ctx context.Context) ChangeFeedResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ChangeFeedResponse) *ChangeFeedResponse {
		return &v
	}).(ChangeFeedResponsePtrOutput)
}

func (o ChangeFeedResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ChangeFeedResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ChangeFeedResponseOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ChangeFeedResponse) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

type ChangeFeedResponsePtrOutput struct{ *pulumi.OutputState }

func (ChangeFeedResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChangeFeedResponse)(nil)).Elem()
}

func (o ChangeFeedResponsePtrOutput) ToChangeFeedResponsePtrOutput() ChangeFeedResponsePtrOutput {
	return o
}

func (o ChangeFeedResponsePtrOutput) ToChangeFeedResponsePtrOutputWithContext(ctx context.Context) ChangeFeedResponsePtrOutput {
	return o
}

func (o ChangeFeedResponsePtrOutput) Elem() ChangeFeedResponseOutput {
	return o.ApplyT(func(v *ChangeFeedResponse) ChangeFeedResponse {
		if v != nil {
			return *v
		}
		var ret ChangeFeedResponse
		return ret
	}).(ChangeFeedResponseOutput)
}

func (o ChangeFeedResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ChangeFeedResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ChangeFeedResponsePtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ChangeFeedResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

type CorsRule struct {
	AllowedHeaders  []string `pulumi:"allowedHeaders"`
	AllowedMethods  []string `pulumi:"allowedMethods"`
	AllowedOrigins  []string `pulumi:"allowedOrigins"`
	ExposedHeaders  []string `pulumi:"exposedHeaders"`
	MaxAgeInSeconds int      `pulumi:"maxAgeInSeconds"`
}





type CorsRuleInput interface {
	pulumi.Input

	ToCorsRuleOutput() CorsRuleOutput
	ToCorsRuleOutputWithContext(context.Context) CorsRuleOutput
}

type CorsRuleArgs struct {
	AllowedHeaders  pulumi.StringArrayInput `pulumi:"allowedHeaders"`
	AllowedMethods  pulumi.StringArrayInput `pulumi:"allowedMethods"`
	AllowedOrigins  pulumi.StringArrayInput `pulumi:"allowedOrigins"`
	ExposedHeaders  pulumi.StringArrayInput `pulumi:"exposedHeaders"`
	MaxAgeInSeconds pulumi.IntInput         `pulumi:"maxAgeInSeconds"`
}

func (CorsRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRule)(nil)).Elem()
}

func (i CorsRuleArgs) ToCorsRuleOutput() CorsRuleOutput {
	return i.ToCorsRuleOutputWithContext(context.Background())
}

func (i CorsRuleArgs) ToCorsRuleOutputWithContext(ctx context.Context) CorsRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRuleOutput)
}





type CorsRuleArrayInput interface {
	pulumi.Input

	ToCorsRuleArrayOutput() CorsRuleArrayOutput
	ToCorsRuleArrayOutputWithContext(context.Context) CorsRuleArrayOutput
}

type CorsRuleArray []CorsRuleInput

func (CorsRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsRule)(nil)).Elem()
}

func (i CorsRuleArray) ToCorsRuleArrayOutput() CorsRuleArrayOutput {
	return i.ToCorsRuleArrayOutputWithContext(context.Background())
}

func (i CorsRuleArray) ToCorsRuleArrayOutputWithContext(ctx context.Context) CorsRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRuleArrayOutput)
}

type CorsRuleOutput struct{ *pulumi.OutputState }

func (CorsRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRule)(nil)).Elem()
}

func (o CorsRuleOutput) ToCorsRuleOutput() CorsRuleOutput {
	return o
}

func (o CorsRuleOutput) ToCorsRuleOutputWithContext(ctx context.Context) CorsRuleOutput {
	return o
}

func (o CorsRuleOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.AllowedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.AllowedMethods }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) ExposedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRule) []string { return v.ExposedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleOutput) MaxAgeInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v CorsRule) int { return v.MaxAgeInSeconds }).(pulumi.IntOutput)
}

type CorsRuleArrayOutput struct{ *pulumi.OutputState }

func (CorsRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsRule)(nil)).Elem()
}

func (o CorsRuleArrayOutput) ToCorsRuleArrayOutput() CorsRuleArrayOutput {
	return o
}

func (o CorsRuleArrayOutput) ToCorsRuleArrayOutputWithContext(ctx context.Context) CorsRuleArrayOutput {
	return o
}

func (o CorsRuleArrayOutput) Index(i pulumi.IntInput) CorsRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CorsRule {
		return vs[0].([]CorsRule)[vs[1].(int)]
	}).(CorsRuleOutput)
}

type CorsRuleResponse struct {
	AllowedHeaders  []string `pulumi:"allowedHeaders"`
	AllowedMethods  []string `pulumi:"allowedMethods"`
	AllowedOrigins  []string `pulumi:"allowedOrigins"`
	ExposedHeaders  []string `pulumi:"exposedHeaders"`
	MaxAgeInSeconds int      `pulumi:"maxAgeInSeconds"`
}





type CorsRuleResponseInput interface {
	pulumi.Input

	ToCorsRuleResponseOutput() CorsRuleResponseOutput
	ToCorsRuleResponseOutputWithContext(context.Context) CorsRuleResponseOutput
}

type CorsRuleResponseArgs struct {
	AllowedHeaders  pulumi.StringArrayInput `pulumi:"allowedHeaders"`
	AllowedMethods  pulumi.StringArrayInput `pulumi:"allowedMethods"`
	AllowedOrigins  pulumi.StringArrayInput `pulumi:"allowedOrigins"`
	ExposedHeaders  pulumi.StringArrayInput `pulumi:"exposedHeaders"`
	MaxAgeInSeconds pulumi.IntInput         `pulumi:"maxAgeInSeconds"`
}

func (CorsRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRuleResponse)(nil)).Elem()
}

func (i CorsRuleResponseArgs) ToCorsRuleResponseOutput() CorsRuleResponseOutput {
	return i.ToCorsRuleResponseOutputWithContext(context.Background())
}

func (i CorsRuleResponseArgs) ToCorsRuleResponseOutputWithContext(ctx context.Context) CorsRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRuleResponseOutput)
}





type CorsRuleResponseArrayInput interface {
	pulumi.Input

	ToCorsRuleResponseArrayOutput() CorsRuleResponseArrayOutput
	ToCorsRuleResponseArrayOutputWithContext(context.Context) CorsRuleResponseArrayOutput
}

type CorsRuleResponseArray []CorsRuleResponseInput

func (CorsRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsRuleResponse)(nil)).Elem()
}

func (i CorsRuleResponseArray) ToCorsRuleResponseArrayOutput() CorsRuleResponseArrayOutput {
	return i.ToCorsRuleResponseArrayOutputWithContext(context.Background())
}

func (i CorsRuleResponseArray) ToCorsRuleResponseArrayOutputWithContext(ctx context.Context) CorsRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRuleResponseArrayOutput)
}

type CorsRuleResponseOutput struct{ *pulumi.OutputState }

func (CorsRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRuleResponse)(nil)).Elem()
}

func (o CorsRuleResponseOutput) ToCorsRuleResponseOutput() CorsRuleResponseOutput {
	return o
}

func (o CorsRuleResponseOutput) ToCorsRuleResponseOutputWithContext(ctx context.Context) CorsRuleResponseOutput {
	return o
}

func (o CorsRuleResponseOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.AllowedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.AllowedMethods }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) ExposedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsRuleResponse) []string { return v.ExposedHeaders }).(pulumi.StringArrayOutput)
}

func (o CorsRuleResponseOutput) MaxAgeInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v CorsRuleResponse) int { return v.MaxAgeInSeconds }).(pulumi.IntOutput)
}

type CorsRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (CorsRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsRuleResponse)(nil)).Elem()
}

func (o CorsRuleResponseArrayOutput) ToCorsRuleResponseArrayOutput() CorsRuleResponseArrayOutput {
	return o
}

func (o CorsRuleResponseArrayOutput) ToCorsRuleResponseArrayOutputWithContext(ctx context.Context) CorsRuleResponseArrayOutput {
	return o
}

func (o CorsRuleResponseArrayOutput) Index(i pulumi.IntInput) CorsRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CorsRuleResponse {
		return vs[0].([]CorsRuleResponse)[vs[1].(int)]
	}).(CorsRuleResponseOutput)
}

type CorsRules struct {
	CorsRules []CorsRule `pulumi:"corsRules"`
}





type CorsRulesInput interface {
	pulumi.Input

	ToCorsRulesOutput() CorsRulesOutput
	ToCorsRulesOutputWithContext(context.Context) CorsRulesOutput
}

type CorsRulesArgs struct {
	CorsRules CorsRuleArrayInput `pulumi:"corsRules"`
}

func (CorsRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRules)(nil)).Elem()
}

func (i CorsRulesArgs) ToCorsRulesOutput() CorsRulesOutput {
	return i.ToCorsRulesOutputWithContext(context.Background())
}

func (i CorsRulesArgs) ToCorsRulesOutputWithContext(ctx context.Context) CorsRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesOutput)
}

func (i CorsRulesArgs) ToCorsRulesPtrOutput() CorsRulesPtrOutput {
	return i.ToCorsRulesPtrOutputWithContext(context.Background())
}

func (i CorsRulesArgs) ToCorsRulesPtrOutputWithContext(ctx context.Context) CorsRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesOutput).ToCorsRulesPtrOutputWithContext(ctx)
}









type CorsRulesPtrInput interface {
	pulumi.Input

	ToCorsRulesPtrOutput() CorsRulesPtrOutput
	ToCorsRulesPtrOutputWithContext(context.Context) CorsRulesPtrOutput
}

type corsRulesPtrType CorsRulesArgs

func CorsRulesPtr(v *CorsRulesArgs) CorsRulesPtrInput {
	return (*corsRulesPtrType)(v)
}

func (*corsRulesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsRules)(nil)).Elem()
}

func (i *corsRulesPtrType) ToCorsRulesPtrOutput() CorsRulesPtrOutput {
	return i.ToCorsRulesPtrOutputWithContext(context.Background())
}

func (i *corsRulesPtrType) ToCorsRulesPtrOutputWithContext(ctx context.Context) CorsRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesPtrOutput)
}

type CorsRulesOutput struct{ *pulumi.OutputState }

func (CorsRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRules)(nil)).Elem()
}

func (o CorsRulesOutput) ToCorsRulesOutput() CorsRulesOutput {
	return o
}

func (o CorsRulesOutput) ToCorsRulesOutputWithContext(ctx context.Context) CorsRulesOutput {
	return o
}

func (o CorsRulesOutput) ToCorsRulesPtrOutput() CorsRulesPtrOutput {
	return o.ToCorsRulesPtrOutputWithContext(context.Background())
}

func (o CorsRulesOutput) ToCorsRulesPtrOutputWithContext(ctx context.Context) CorsRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CorsRules) *CorsRules {
		return &v
	}).(CorsRulesPtrOutput)
}

func (o CorsRulesOutput) CorsRules() CorsRuleArrayOutput {
	return o.ApplyT(func(v CorsRules) []CorsRule { return v.CorsRules }).(CorsRuleArrayOutput)
}

type CorsRulesPtrOutput struct{ *pulumi.OutputState }

func (CorsRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsRules)(nil)).Elem()
}

func (o CorsRulesPtrOutput) ToCorsRulesPtrOutput() CorsRulesPtrOutput {
	return o
}

func (o CorsRulesPtrOutput) ToCorsRulesPtrOutputWithContext(ctx context.Context) CorsRulesPtrOutput {
	return o
}

func (o CorsRulesPtrOutput) Elem() CorsRulesOutput {
	return o.ApplyT(func(v *CorsRules) CorsRules {
		if v != nil {
			return *v
		}
		var ret CorsRules
		return ret
	}).(CorsRulesOutput)
}

func (o CorsRulesPtrOutput) CorsRules() CorsRuleArrayOutput {
	return o.ApplyT(func(v *CorsRules) []CorsRule {
		if v == nil {
			return nil
		}
		return v.CorsRules
	}).(CorsRuleArrayOutput)
}

type CorsRulesResponse struct {
	CorsRules []CorsRuleResponse `pulumi:"corsRules"`
}





type CorsRulesResponseInput interface {
	pulumi.Input

	ToCorsRulesResponseOutput() CorsRulesResponseOutput
	ToCorsRulesResponseOutputWithContext(context.Context) CorsRulesResponseOutput
}

type CorsRulesResponseArgs struct {
	CorsRules CorsRuleResponseArrayInput `pulumi:"corsRules"`
}

func (CorsRulesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRulesResponse)(nil)).Elem()
}

func (i CorsRulesResponseArgs) ToCorsRulesResponseOutput() CorsRulesResponseOutput {
	return i.ToCorsRulesResponseOutputWithContext(context.Background())
}

func (i CorsRulesResponseArgs) ToCorsRulesResponseOutputWithContext(ctx context.Context) CorsRulesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesResponseOutput)
}

func (i CorsRulesResponseArgs) ToCorsRulesResponsePtrOutput() CorsRulesResponsePtrOutput {
	return i.ToCorsRulesResponsePtrOutputWithContext(context.Background())
}

func (i CorsRulesResponseArgs) ToCorsRulesResponsePtrOutputWithContext(ctx context.Context) CorsRulesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesResponseOutput).ToCorsRulesResponsePtrOutputWithContext(ctx)
}









type CorsRulesResponsePtrInput interface {
	pulumi.Input

	ToCorsRulesResponsePtrOutput() CorsRulesResponsePtrOutput
	ToCorsRulesResponsePtrOutputWithContext(context.Context) CorsRulesResponsePtrOutput
}

type corsRulesResponsePtrType CorsRulesResponseArgs

func CorsRulesResponsePtr(v *CorsRulesResponseArgs) CorsRulesResponsePtrInput {
	return (*corsRulesResponsePtrType)(v)
}

func (*corsRulesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsRulesResponse)(nil)).Elem()
}

func (i *corsRulesResponsePtrType) ToCorsRulesResponsePtrOutput() CorsRulesResponsePtrOutput {
	return i.ToCorsRulesResponsePtrOutputWithContext(context.Background())
}

func (i *corsRulesResponsePtrType) ToCorsRulesResponsePtrOutputWithContext(ctx context.Context) CorsRulesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsRulesResponsePtrOutput)
}

type CorsRulesResponseOutput struct{ *pulumi.OutputState }

func (CorsRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsRulesResponse)(nil)).Elem()
}

func (o CorsRulesResponseOutput) ToCorsRulesResponseOutput() CorsRulesResponseOutput {
	return o
}

func (o CorsRulesResponseOutput) ToCorsRulesResponseOutputWithContext(ctx context.Context) CorsRulesResponseOutput {
	return o
}

func (o CorsRulesResponseOutput) ToCorsRulesResponsePtrOutput() CorsRulesResponsePtrOutput {
	return o.ToCorsRulesResponsePtrOutputWithContext(context.Background())
}

func (o CorsRulesResponseOutput) ToCorsRulesResponsePtrOutputWithContext(ctx context.Context) CorsRulesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CorsRulesResponse) *CorsRulesResponse {
		return &v
	}).(CorsRulesResponsePtrOutput)
}

func (o CorsRulesResponseOutput) CorsRules() CorsRuleResponseArrayOutput {
	return o.ApplyT(func(v CorsRulesResponse) []CorsRuleResponse { return v.CorsRules }).(CorsRuleResponseArrayOutput)
}

type CorsRulesResponsePtrOutput struct{ *pulumi.OutputState }

func (CorsRulesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsRulesResponse)(nil)).Elem()
}

func (o CorsRulesResponsePtrOutput) ToCorsRulesResponsePtrOutput() CorsRulesResponsePtrOutput {
	return o
}

func (o CorsRulesResponsePtrOutput) ToCorsRulesResponsePtrOutputWithContext(ctx context.Context) CorsRulesResponsePtrOutput {
	return o
}

func (o CorsRulesResponsePtrOutput) Elem() CorsRulesResponseOutput {
	return o.ApplyT(func(v *CorsRulesResponse) CorsRulesResponse {
		if v != nil {
			return *v
		}
		var ret CorsRulesResponse
		return ret
	}).(CorsRulesResponseOutput)
}

func (o CorsRulesResponsePtrOutput) CorsRules() CorsRuleResponseArrayOutput {
	return o.ApplyT(func(v *CorsRulesResponse) []CorsRuleResponse {
		if v == nil {
			return nil
		}
		return v.CorsRules
	}).(CorsRuleResponseArrayOutput)
}

type CustomDomain struct {
	Name             string `pulumi:"name"`
	UseSubDomainName *bool  `pulumi:"useSubDomainName"`
}





type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(context.Context) CustomDomainOutput
}

type CustomDomainArgs struct {
	Name             pulumi.StringInput  `pulumi:"name"`
	UseSubDomainName pulumi.BoolPtrInput `pulumi:"useSubDomainName"`
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil)).Elem()
}

func (i CustomDomainArgs) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i CustomDomainArgs) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

func (i CustomDomainArgs) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i CustomDomainArgs) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput).ToCustomDomainPtrOutputWithContext(ctx)
}









type CustomDomainPtrInput interface {
	pulumi.Input

	ToCustomDomainPtrOutput() CustomDomainPtrOutput
	ToCustomDomainPtrOutputWithContext(context.Context) CustomDomainPtrOutput
}

type customDomainPtrType CustomDomainArgs

func CustomDomainPtr(v *CustomDomainArgs) CustomDomainPtrInput {
	return (*customDomainPtrType)(v)
}

func (*customDomainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (i *customDomainPtrType) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i *customDomainPtrType) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPtrOutput)
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil)).Elem()
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return o.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (o CustomDomainOutput) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDomain) *CustomDomain {
		return &v
	}).(CustomDomainPtrOutput)
}

func (o CustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomain) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomain) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
}

type CustomDomainPtrOutput struct{ *pulumi.OutputState }

func (CustomDomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (o CustomDomainPtrOutput) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return o
}

func (o CustomDomainPtrOutput) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return o
}

func (o CustomDomainPtrOutput) Elem() CustomDomainOutput {
	return o.ApplyT(func(v *CustomDomain) CustomDomain {
		if v != nil {
			return *v
		}
		var ret CustomDomain
		return ret
	}).(CustomDomainOutput)
}

func (o CustomDomainPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomain) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainPtrOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomDomain) *bool {
		if v == nil {
			return nil
		}
		return v.UseSubDomainName
	}).(pulumi.BoolPtrOutput)
}

type CustomDomainResponse struct {
	Name             string `pulumi:"name"`
	UseSubDomainName *bool  `pulumi:"useSubDomainName"`
}





type CustomDomainResponseInput interface {
	pulumi.Input

	ToCustomDomainResponseOutput() CustomDomainResponseOutput
	ToCustomDomainResponseOutputWithContext(context.Context) CustomDomainResponseOutput
}

type CustomDomainResponseArgs struct {
	Name             pulumi.StringInput  `pulumi:"name"`
	UseSubDomainName pulumi.BoolPtrInput `pulumi:"useSubDomainName"`
}

func (CustomDomainResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainResponse)(nil)).Elem()
}

func (i CustomDomainResponseArgs) ToCustomDomainResponseOutput() CustomDomainResponseOutput {
	return i.ToCustomDomainResponseOutputWithContext(context.Background())
}

func (i CustomDomainResponseArgs) ToCustomDomainResponseOutputWithContext(ctx context.Context) CustomDomainResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainResponseOutput)
}

func (i CustomDomainResponseArgs) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return i.ToCustomDomainResponsePtrOutputWithContext(context.Background())
}

func (i CustomDomainResponseArgs) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainResponseOutput).ToCustomDomainResponsePtrOutputWithContext(ctx)
}









type CustomDomainResponsePtrInput interface {
	pulumi.Input

	ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput
	ToCustomDomainResponsePtrOutputWithContext(context.Context) CustomDomainResponsePtrOutput
}

type customDomainResponsePtrType CustomDomainResponseArgs

func CustomDomainResponsePtr(v *CustomDomainResponseArgs) CustomDomainResponsePtrInput {
	return (*customDomainResponsePtrType)(v)
}

func (*customDomainResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainResponse)(nil)).Elem()
}

func (i *customDomainResponsePtrType) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return i.ToCustomDomainResponsePtrOutputWithContext(context.Background())
}

func (i *customDomainResponsePtrType) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainResponsePtrOutput)
}

type CustomDomainResponseOutput struct{ *pulumi.OutputState }

func (CustomDomainResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutput() CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutputWithContext(ctx context.Context) CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return o.ToCustomDomainResponsePtrOutputWithContext(context.Background())
}

func (o CustomDomainResponseOutput) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDomainResponse) *CustomDomainResponse {
		return &v
	}).(CustomDomainResponsePtrOutput)
}

func (o CustomDomainResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomainResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDomainResponseOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
}

type CustomDomainResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomDomainResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponsePtrOutput) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return o
}

func (o CustomDomainResponsePtrOutput) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return o
}

func (o CustomDomainResponsePtrOutput) Elem() CustomDomainResponseOutput {
	return o.ApplyT(func(v *CustomDomainResponse) CustomDomainResponse {
		if v != nil {
			return *v
		}
		var ret CustomDomainResponse
		return ret
	}).(CustomDomainResponseOutput)
}

func (o CustomDomainResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainResponsePtrOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomDomainResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseSubDomainName
	}).(pulumi.BoolPtrOutput)
}

type DateAfterCreation struct {
	DaysAfterCreationGreaterThan float64 `pulumi:"daysAfterCreationGreaterThan"`
}





type DateAfterCreationInput interface {
	pulumi.Input

	ToDateAfterCreationOutput() DateAfterCreationOutput
	ToDateAfterCreationOutputWithContext(context.Context) DateAfterCreationOutput
}

type DateAfterCreationArgs struct {
	DaysAfterCreationGreaterThan pulumi.Float64Input `pulumi:"daysAfterCreationGreaterThan"`
}

func (DateAfterCreationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterCreation)(nil)).Elem()
}

func (i DateAfterCreationArgs) ToDateAfterCreationOutput() DateAfterCreationOutput {
	return i.ToDateAfterCreationOutputWithContext(context.Background())
}

func (i DateAfterCreationArgs) ToDateAfterCreationOutputWithContext(ctx context.Context) DateAfterCreationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterCreationOutput)
}

func (i DateAfterCreationArgs) ToDateAfterCreationPtrOutput() DateAfterCreationPtrOutput {
	return i.ToDateAfterCreationPtrOutputWithContext(context.Background())
}

func (i DateAfterCreationArgs) ToDateAfterCreationPtrOutputWithContext(ctx context.Context) DateAfterCreationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterCreationOutput).ToDateAfterCreationPtrOutputWithContext(ctx)
}









type DateAfterCreationPtrInput interface {
	pulumi.Input

	ToDateAfterCreationPtrOutput() DateAfterCreationPtrOutput
	ToDateAfterCreationPtrOutputWithContext(context.Context) DateAfterCreationPtrOutput
}

type dateAfterCreationPtrType DateAfterCreationArgs

func DateAfterCreationPtr(v *DateAfterCreationArgs) DateAfterCreationPtrInput {
	return (*dateAfterCreationPtrType)(v)
}

func (*dateAfterCreationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterCreation)(nil)).Elem()
}

func (i *dateAfterCreationPtrType) ToDateAfterCreationPtrOutput() DateAfterCreationPtrOutput {
	return i.ToDateAfterCreationPtrOutputWithContext(context.Background())
}

func (i *dateAfterCreationPtrType) ToDateAfterCreationPtrOutputWithContext(ctx context.Context) DateAfterCreationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterCreationPtrOutput)
}

type DateAfterCreationOutput struct{ *pulumi.OutputState }

func (DateAfterCreationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterCreation)(nil)).Elem()
}

func (o DateAfterCreationOutput) ToDateAfterCreationOutput() DateAfterCreationOutput {
	return o
}

func (o DateAfterCreationOutput) ToDateAfterCreationOutputWithContext(ctx context.Context) DateAfterCreationOutput {
	return o
}

func (o DateAfterCreationOutput) ToDateAfterCreationPtrOutput() DateAfterCreationPtrOutput {
	return o.ToDateAfterCreationPtrOutputWithContext(context.Background())
}

func (o DateAfterCreationOutput) ToDateAfterCreationPtrOutputWithContext(ctx context.Context) DateAfterCreationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DateAfterCreation) *DateAfterCreation {
		return &v
	}).(DateAfterCreationPtrOutput)
}

func (o DateAfterCreationOutput) DaysAfterCreationGreaterThan() pulumi.Float64Output {
	return o.ApplyT(func(v DateAfterCreation) float64 { return v.DaysAfterCreationGreaterThan }).(pulumi.Float64Output)
}

type DateAfterCreationPtrOutput struct{ *pulumi.OutputState }

func (DateAfterCreationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterCreation)(nil)).Elem()
}

func (o DateAfterCreationPtrOutput) ToDateAfterCreationPtrOutput() DateAfterCreationPtrOutput {
	return o
}

func (o DateAfterCreationPtrOutput) ToDateAfterCreationPtrOutputWithContext(ctx context.Context) DateAfterCreationPtrOutput {
	return o
}

func (o DateAfterCreationPtrOutput) Elem() DateAfterCreationOutput {
	return o.ApplyT(func(v *DateAfterCreation) DateAfterCreation {
		if v != nil {
			return *v
		}
		var ret DateAfterCreation
		return ret
	}).(DateAfterCreationOutput)
}

func (o DateAfterCreationPtrOutput) DaysAfterCreationGreaterThan() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DateAfterCreation) *float64 {
		if v == nil {
			return nil
		}
		return &v.DaysAfterCreationGreaterThan
	}).(pulumi.Float64PtrOutput)
}

type DateAfterCreationResponse struct {
	DaysAfterCreationGreaterThan float64 `pulumi:"daysAfterCreationGreaterThan"`
}





type DateAfterCreationResponseInput interface {
	pulumi.Input

	ToDateAfterCreationResponseOutput() DateAfterCreationResponseOutput
	ToDateAfterCreationResponseOutputWithContext(context.Context) DateAfterCreationResponseOutput
}

type DateAfterCreationResponseArgs struct {
	DaysAfterCreationGreaterThan pulumi.Float64Input `pulumi:"daysAfterCreationGreaterThan"`
}

func (DateAfterCreationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterCreationResponse)(nil)).Elem()
}

func (i DateAfterCreationResponseArgs) ToDateAfterCreationResponseOutput() DateAfterCreationResponseOutput {
	return i.ToDateAfterCreationResponseOutputWithContext(context.Background())
}

func (i DateAfterCreationResponseArgs) ToDateAfterCreationResponseOutputWithContext(ctx context.Context) DateAfterCreationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterCreationResponseOutput)
}

func (i DateAfterCreationResponseArgs) ToDateAfterCreationResponsePtrOutput() DateAfterCreationResponsePtrOutput {
	return i.ToDateAfterCreationResponsePtrOutputWithContext(context.Background())
}

func (i DateAfterCreationResponseArgs) ToDateAfterCreationResponsePtrOutputWithContext(ctx context.Context) DateAfterCreationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterCreationResponseOutput).ToDateAfterCreationResponsePtrOutputWithContext(ctx)
}









type DateAfterCreationResponsePtrInput interface {
	pulumi.Input

	ToDateAfterCreationResponsePtrOutput() DateAfterCreationResponsePtrOutput
	ToDateAfterCreationResponsePtrOutputWithContext(context.Context) DateAfterCreationResponsePtrOutput
}

type dateAfterCreationResponsePtrType DateAfterCreationResponseArgs

func DateAfterCreationResponsePtr(v *DateAfterCreationResponseArgs) DateAfterCreationResponsePtrInput {
	return (*dateAfterCreationResponsePtrType)(v)
}

func (*dateAfterCreationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterCreationResponse)(nil)).Elem()
}

func (i *dateAfterCreationResponsePtrType) ToDateAfterCreationResponsePtrOutput() DateAfterCreationResponsePtrOutput {
	return i.ToDateAfterCreationResponsePtrOutputWithContext(context.Background())
}

func (i *dateAfterCreationResponsePtrType) ToDateAfterCreationResponsePtrOutputWithContext(ctx context.Context) DateAfterCreationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterCreationResponsePtrOutput)
}

type DateAfterCreationResponseOutput struct{ *pulumi.OutputState }

func (DateAfterCreationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterCreationResponse)(nil)).Elem()
}

func (o DateAfterCreationResponseOutput) ToDateAfterCreationResponseOutput() DateAfterCreationResponseOutput {
	return o
}

func (o DateAfterCreationResponseOutput) ToDateAfterCreationResponseOutputWithContext(ctx context.Context) DateAfterCreationResponseOutput {
	return o
}

func (o DateAfterCreationResponseOutput) ToDateAfterCreationResponsePtrOutput() DateAfterCreationResponsePtrOutput {
	return o.ToDateAfterCreationResponsePtrOutputWithContext(context.Background())
}

func (o DateAfterCreationResponseOutput) ToDateAfterCreationResponsePtrOutputWithContext(ctx context.Context) DateAfterCreationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DateAfterCreationResponse) *DateAfterCreationResponse {
		return &v
	}).(DateAfterCreationResponsePtrOutput)
}

func (o DateAfterCreationResponseOutput) DaysAfterCreationGreaterThan() pulumi.Float64Output {
	return o.ApplyT(func(v DateAfterCreationResponse) float64 { return v.DaysAfterCreationGreaterThan }).(pulumi.Float64Output)
}

type DateAfterCreationResponsePtrOutput struct{ *pulumi.OutputState }

func (DateAfterCreationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterCreationResponse)(nil)).Elem()
}

func (o DateAfterCreationResponsePtrOutput) ToDateAfterCreationResponsePtrOutput() DateAfterCreationResponsePtrOutput {
	return o
}

func (o DateAfterCreationResponsePtrOutput) ToDateAfterCreationResponsePtrOutputWithContext(ctx context.Context) DateAfterCreationResponsePtrOutput {
	return o
}

func (o DateAfterCreationResponsePtrOutput) Elem() DateAfterCreationResponseOutput {
	return o.ApplyT(func(v *DateAfterCreationResponse) DateAfterCreationResponse {
		if v != nil {
			return *v
		}
		var ret DateAfterCreationResponse
		return ret
	}).(DateAfterCreationResponseOutput)
}

func (o DateAfterCreationResponsePtrOutput) DaysAfterCreationGreaterThan() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DateAfterCreationResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.DaysAfterCreationGreaterThan
	}).(pulumi.Float64PtrOutput)
}

type DateAfterModification struct {
	DaysAfterLastAccessTimeGreaterThan *float64 `pulumi:"daysAfterLastAccessTimeGreaterThan"`
	DaysAfterModificationGreaterThan   *float64 `pulumi:"daysAfterModificationGreaterThan"`
}





type DateAfterModificationInput interface {
	pulumi.Input

	ToDateAfterModificationOutput() DateAfterModificationOutput
	ToDateAfterModificationOutputWithContext(context.Context) DateAfterModificationOutput
}

type DateAfterModificationArgs struct {
	DaysAfterLastAccessTimeGreaterThan pulumi.Float64PtrInput `pulumi:"daysAfterLastAccessTimeGreaterThan"`
	DaysAfterModificationGreaterThan   pulumi.Float64PtrInput `pulumi:"daysAfterModificationGreaterThan"`
}

func (DateAfterModificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterModification)(nil)).Elem()
}

func (i DateAfterModificationArgs) ToDateAfterModificationOutput() DateAfterModificationOutput {
	return i.ToDateAfterModificationOutputWithContext(context.Background())
}

func (i DateAfterModificationArgs) ToDateAfterModificationOutputWithContext(ctx context.Context) DateAfterModificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterModificationOutput)
}

func (i DateAfterModificationArgs) ToDateAfterModificationPtrOutput() DateAfterModificationPtrOutput {
	return i.ToDateAfterModificationPtrOutputWithContext(context.Background())
}

func (i DateAfterModificationArgs) ToDateAfterModificationPtrOutputWithContext(ctx context.Context) DateAfterModificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterModificationOutput).ToDateAfterModificationPtrOutputWithContext(ctx)
}









type DateAfterModificationPtrInput interface {
	pulumi.Input

	ToDateAfterModificationPtrOutput() DateAfterModificationPtrOutput
	ToDateAfterModificationPtrOutputWithContext(context.Context) DateAfterModificationPtrOutput
}

type dateAfterModificationPtrType DateAfterModificationArgs

func DateAfterModificationPtr(v *DateAfterModificationArgs) DateAfterModificationPtrInput {
	return (*dateAfterModificationPtrType)(v)
}

func (*dateAfterModificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterModification)(nil)).Elem()
}

func (i *dateAfterModificationPtrType) ToDateAfterModificationPtrOutput() DateAfterModificationPtrOutput {
	return i.ToDateAfterModificationPtrOutputWithContext(context.Background())
}

func (i *dateAfterModificationPtrType) ToDateAfterModificationPtrOutputWithContext(ctx context.Context) DateAfterModificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterModificationPtrOutput)
}

type DateAfterModificationOutput struct{ *pulumi.OutputState }

func (DateAfterModificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterModification)(nil)).Elem()
}

func (o DateAfterModificationOutput) ToDateAfterModificationOutput() DateAfterModificationOutput {
	return o
}

func (o DateAfterModificationOutput) ToDateAfterModificationOutputWithContext(ctx context.Context) DateAfterModificationOutput {
	return o
}

func (o DateAfterModificationOutput) ToDateAfterModificationPtrOutput() DateAfterModificationPtrOutput {
	return o.ToDateAfterModificationPtrOutputWithContext(context.Background())
}

func (o DateAfterModificationOutput) ToDateAfterModificationPtrOutputWithContext(ctx context.Context) DateAfterModificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DateAfterModification) *DateAfterModification {
		return &v
	}).(DateAfterModificationPtrOutput)
}

func (o DateAfterModificationOutput) DaysAfterLastAccessTimeGreaterThan() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DateAfterModification) *float64 { return v.DaysAfterLastAccessTimeGreaterThan }).(pulumi.Float64PtrOutput)
}

func (o DateAfterModificationOutput) DaysAfterModificationGreaterThan() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DateAfterModification) *float64 { return v.DaysAfterModificationGreaterThan }).(pulumi.Float64PtrOutput)
}

type DateAfterModificationPtrOutput struct{ *pulumi.OutputState }

func (DateAfterModificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterModification)(nil)).Elem()
}

func (o DateAfterModificationPtrOutput) ToDateAfterModificationPtrOutput() DateAfterModificationPtrOutput {
	return o
}

func (o DateAfterModificationPtrOutput) ToDateAfterModificationPtrOutputWithContext(ctx context.Context) DateAfterModificationPtrOutput {
	return o
}

func (o DateAfterModificationPtrOutput) Elem() DateAfterModificationOutput {
	return o.ApplyT(func(v *DateAfterModification) DateAfterModification {
		if v != nil {
			return *v
		}
		var ret DateAfterModification
		return ret
	}).(DateAfterModificationOutput)
}

func (o DateAfterModificationPtrOutput) DaysAfterLastAccessTimeGreaterThan() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DateAfterModification) *float64 {
		if v == nil {
			return nil
		}
		return v.DaysAfterLastAccessTimeGreaterThan
	}).(pulumi.Float64PtrOutput)
}

func (o DateAfterModificationPtrOutput) DaysAfterModificationGreaterThan() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DateAfterModification) *float64 {
		if v == nil {
			return nil
		}
		return v.DaysAfterModificationGreaterThan
	}).(pulumi.Float64PtrOutput)
}

type DateAfterModificationResponse struct {
	DaysAfterLastAccessTimeGreaterThan *float64 `pulumi:"daysAfterLastAccessTimeGreaterThan"`
	DaysAfterModificationGreaterThan   *float64 `pulumi:"daysAfterModificationGreaterThan"`
}





type DateAfterModificationResponseInput interface {
	pulumi.Input

	ToDateAfterModificationResponseOutput() DateAfterModificationResponseOutput
	ToDateAfterModificationResponseOutputWithContext(context.Context) DateAfterModificationResponseOutput
}

type DateAfterModificationResponseArgs struct {
	DaysAfterLastAccessTimeGreaterThan pulumi.Float64PtrInput `pulumi:"daysAfterLastAccessTimeGreaterThan"`
	DaysAfterModificationGreaterThan   pulumi.Float64PtrInput `pulumi:"daysAfterModificationGreaterThan"`
}

func (DateAfterModificationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterModificationResponse)(nil)).Elem()
}

func (i DateAfterModificationResponseArgs) ToDateAfterModificationResponseOutput() DateAfterModificationResponseOutput {
	return i.ToDateAfterModificationResponseOutputWithContext(context.Background())
}

func (i DateAfterModificationResponseArgs) ToDateAfterModificationResponseOutputWithContext(ctx context.Context) DateAfterModificationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterModificationResponseOutput)
}

func (i DateAfterModificationResponseArgs) ToDateAfterModificationResponsePtrOutput() DateAfterModificationResponsePtrOutput {
	return i.ToDateAfterModificationResponsePtrOutputWithContext(context.Background())
}

func (i DateAfterModificationResponseArgs) ToDateAfterModificationResponsePtrOutputWithContext(ctx context.Context) DateAfterModificationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterModificationResponseOutput).ToDateAfterModificationResponsePtrOutputWithContext(ctx)
}









type DateAfterModificationResponsePtrInput interface {
	pulumi.Input

	ToDateAfterModificationResponsePtrOutput() DateAfterModificationResponsePtrOutput
	ToDateAfterModificationResponsePtrOutputWithContext(context.Context) DateAfterModificationResponsePtrOutput
}

type dateAfterModificationResponsePtrType DateAfterModificationResponseArgs

func DateAfterModificationResponsePtr(v *DateAfterModificationResponseArgs) DateAfterModificationResponsePtrInput {
	return (*dateAfterModificationResponsePtrType)(v)
}

func (*dateAfterModificationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterModificationResponse)(nil)).Elem()
}

func (i *dateAfterModificationResponsePtrType) ToDateAfterModificationResponsePtrOutput() DateAfterModificationResponsePtrOutput {
	return i.ToDateAfterModificationResponsePtrOutputWithContext(context.Background())
}

func (i *dateAfterModificationResponsePtrType) ToDateAfterModificationResponsePtrOutputWithContext(ctx context.Context) DateAfterModificationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DateAfterModificationResponsePtrOutput)
}

type DateAfterModificationResponseOutput struct{ *pulumi.OutputState }

func (DateAfterModificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DateAfterModificationResponse)(nil)).Elem()
}

func (o DateAfterModificationResponseOutput) ToDateAfterModificationResponseOutput() DateAfterModificationResponseOutput {
	return o
}

func (o DateAfterModificationResponseOutput) ToDateAfterModificationResponseOutputWithContext(ctx context.Context) DateAfterModificationResponseOutput {
	return o
}

func (o DateAfterModificationResponseOutput) ToDateAfterModificationResponsePtrOutput() DateAfterModificationResponsePtrOutput {
	return o.ToDateAfterModificationResponsePtrOutputWithContext(context.Background())
}

func (o DateAfterModificationResponseOutput) ToDateAfterModificationResponsePtrOutputWithContext(ctx context.Context) DateAfterModificationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DateAfterModificationResponse) *DateAfterModificationResponse {
		return &v
	}).(DateAfterModificationResponsePtrOutput)
}

func (o DateAfterModificationResponseOutput) DaysAfterLastAccessTimeGreaterThan() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DateAfterModificationResponse) *float64 { return v.DaysAfterLastAccessTimeGreaterThan }).(pulumi.Float64PtrOutput)
}

func (o DateAfterModificationResponseOutput) DaysAfterModificationGreaterThan() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DateAfterModificationResponse) *float64 { return v.DaysAfterModificationGreaterThan }).(pulumi.Float64PtrOutput)
}

type DateAfterModificationResponsePtrOutput struct{ *pulumi.OutputState }

func (DateAfterModificationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DateAfterModificationResponse)(nil)).Elem()
}

func (o DateAfterModificationResponsePtrOutput) ToDateAfterModificationResponsePtrOutput() DateAfterModificationResponsePtrOutput {
	return o
}

func (o DateAfterModificationResponsePtrOutput) ToDateAfterModificationResponsePtrOutputWithContext(ctx context.Context) DateAfterModificationResponsePtrOutput {
	return o
}

func (o DateAfterModificationResponsePtrOutput) Elem() DateAfterModificationResponseOutput {
	return o.ApplyT(func(v *DateAfterModificationResponse) DateAfterModificationResponse {
		if v != nil {
			return *v
		}
		var ret DateAfterModificationResponse
		return ret
	}).(DateAfterModificationResponseOutput)
}

func (o DateAfterModificationResponsePtrOutput) DaysAfterLastAccessTimeGreaterThan() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DateAfterModificationResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.DaysAfterLastAccessTimeGreaterThan
	}).(pulumi.Float64PtrOutput)
}

func (o DateAfterModificationResponsePtrOutput) DaysAfterModificationGreaterThan() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DateAfterModificationResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.DaysAfterModificationGreaterThan
	}).(pulumi.Float64PtrOutput)
}

type DeleteRetentionPolicy struct {
	Days    *int  `pulumi:"days"`
	Enabled *bool `pulumi:"enabled"`
}





type DeleteRetentionPolicyInput interface {
	pulumi.Input

	ToDeleteRetentionPolicyOutput() DeleteRetentionPolicyOutput
	ToDeleteRetentionPolicyOutputWithContext(context.Context) DeleteRetentionPolicyOutput
}

type DeleteRetentionPolicyArgs struct {
	Days    pulumi.IntPtrInput  `pulumi:"days"`
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (DeleteRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteRetentionPolicy)(nil)).Elem()
}

func (i DeleteRetentionPolicyArgs) ToDeleteRetentionPolicyOutput() DeleteRetentionPolicyOutput {
	return i.ToDeleteRetentionPolicyOutputWithContext(context.Background())
}

func (i DeleteRetentionPolicyArgs) ToDeleteRetentionPolicyOutputWithContext(ctx context.Context) DeleteRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeleteRetentionPolicyOutput)
}

func (i DeleteRetentionPolicyArgs) ToDeleteRetentionPolicyPtrOutput() DeleteRetentionPolicyPtrOutput {
	return i.ToDeleteRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i DeleteRetentionPolicyArgs) ToDeleteRetentionPolicyPtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeleteRetentionPolicyOutput).ToDeleteRetentionPolicyPtrOutputWithContext(ctx)
}









type DeleteRetentionPolicyPtrInput interface {
	pulumi.Input

	ToDeleteRetentionPolicyPtrOutput() DeleteRetentionPolicyPtrOutput
	ToDeleteRetentionPolicyPtrOutputWithContext(context.Context) DeleteRetentionPolicyPtrOutput
}

type deleteRetentionPolicyPtrType DeleteRetentionPolicyArgs

func DeleteRetentionPolicyPtr(v *DeleteRetentionPolicyArgs) DeleteRetentionPolicyPtrInput {
	return (*deleteRetentionPolicyPtrType)(v)
}

func (*deleteRetentionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeleteRetentionPolicy)(nil)).Elem()
}

func (i *deleteRetentionPolicyPtrType) ToDeleteRetentionPolicyPtrOutput() DeleteRetentionPolicyPtrOutput {
	return i.ToDeleteRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i *deleteRetentionPolicyPtrType) ToDeleteRetentionPolicyPtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeleteRetentionPolicyPtrOutput)
}

type DeleteRetentionPolicyOutput struct{ *pulumi.OutputState }

func (DeleteRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteRetentionPolicy)(nil)).Elem()
}

func (o DeleteRetentionPolicyOutput) ToDeleteRetentionPolicyOutput() DeleteRetentionPolicyOutput {
	return o
}

func (o DeleteRetentionPolicyOutput) ToDeleteRetentionPolicyOutputWithContext(ctx context.Context) DeleteRetentionPolicyOutput {
	return o
}

func (o DeleteRetentionPolicyOutput) ToDeleteRetentionPolicyPtrOutput() DeleteRetentionPolicyPtrOutput {
	return o.ToDeleteRetentionPolicyPtrOutputWithContext(context.Background())
}

func (o DeleteRetentionPolicyOutput) ToDeleteRetentionPolicyPtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeleteRetentionPolicy) *DeleteRetentionPolicy {
		return &v
	}).(DeleteRetentionPolicyPtrOutput)
}

func (o DeleteRetentionPolicyOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeleteRetentionPolicy) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o DeleteRetentionPolicyOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DeleteRetentionPolicy) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type DeleteRetentionPolicyPtrOutput struct{ *pulumi.OutputState }

func (DeleteRetentionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeleteRetentionPolicy)(nil)).Elem()
}

func (o DeleteRetentionPolicyPtrOutput) ToDeleteRetentionPolicyPtrOutput() DeleteRetentionPolicyPtrOutput {
	return o
}

func (o DeleteRetentionPolicyPtrOutput) ToDeleteRetentionPolicyPtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyPtrOutput {
	return o
}

func (o DeleteRetentionPolicyPtrOutput) Elem() DeleteRetentionPolicyOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicy) DeleteRetentionPolicy {
		if v != nil {
			return *v
		}
		var ret DeleteRetentionPolicy
		return ret
	}).(DeleteRetentionPolicyOutput)
}

func (o DeleteRetentionPolicyPtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicy) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o DeleteRetentionPolicyPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type DeleteRetentionPolicyResponse struct {
	Days    *int  `pulumi:"days"`
	Enabled *bool `pulumi:"enabled"`
}





type DeleteRetentionPolicyResponseInput interface {
	pulumi.Input

	ToDeleteRetentionPolicyResponseOutput() DeleteRetentionPolicyResponseOutput
	ToDeleteRetentionPolicyResponseOutputWithContext(context.Context) DeleteRetentionPolicyResponseOutput
}

type DeleteRetentionPolicyResponseArgs struct {
	Days    pulumi.IntPtrInput  `pulumi:"days"`
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (DeleteRetentionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteRetentionPolicyResponse)(nil)).Elem()
}

func (i DeleteRetentionPolicyResponseArgs) ToDeleteRetentionPolicyResponseOutput() DeleteRetentionPolicyResponseOutput {
	return i.ToDeleteRetentionPolicyResponseOutputWithContext(context.Background())
}

func (i DeleteRetentionPolicyResponseArgs) ToDeleteRetentionPolicyResponseOutputWithContext(ctx context.Context) DeleteRetentionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeleteRetentionPolicyResponseOutput)
}

func (i DeleteRetentionPolicyResponseArgs) ToDeleteRetentionPolicyResponsePtrOutput() DeleteRetentionPolicyResponsePtrOutput {
	return i.ToDeleteRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i DeleteRetentionPolicyResponseArgs) ToDeleteRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeleteRetentionPolicyResponseOutput).ToDeleteRetentionPolicyResponsePtrOutputWithContext(ctx)
}









type DeleteRetentionPolicyResponsePtrInput interface {
	pulumi.Input

	ToDeleteRetentionPolicyResponsePtrOutput() DeleteRetentionPolicyResponsePtrOutput
	ToDeleteRetentionPolicyResponsePtrOutputWithContext(context.Context) DeleteRetentionPolicyResponsePtrOutput
}

type deleteRetentionPolicyResponsePtrType DeleteRetentionPolicyResponseArgs

func DeleteRetentionPolicyResponsePtr(v *DeleteRetentionPolicyResponseArgs) DeleteRetentionPolicyResponsePtrInput {
	return (*deleteRetentionPolicyResponsePtrType)(v)
}

func (*deleteRetentionPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeleteRetentionPolicyResponse)(nil)).Elem()
}

func (i *deleteRetentionPolicyResponsePtrType) ToDeleteRetentionPolicyResponsePtrOutput() DeleteRetentionPolicyResponsePtrOutput {
	return i.ToDeleteRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *deleteRetentionPolicyResponsePtrType) ToDeleteRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeleteRetentionPolicyResponsePtrOutput)
}

type DeleteRetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (DeleteRetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteRetentionPolicyResponse)(nil)).Elem()
}

func (o DeleteRetentionPolicyResponseOutput) ToDeleteRetentionPolicyResponseOutput() DeleteRetentionPolicyResponseOutput {
	return o
}

func (o DeleteRetentionPolicyResponseOutput) ToDeleteRetentionPolicyResponseOutputWithContext(ctx context.Context) DeleteRetentionPolicyResponseOutput {
	return o
}

func (o DeleteRetentionPolicyResponseOutput) ToDeleteRetentionPolicyResponsePtrOutput() DeleteRetentionPolicyResponsePtrOutput {
	return o.ToDeleteRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (o DeleteRetentionPolicyResponseOutput) ToDeleteRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeleteRetentionPolicyResponse) *DeleteRetentionPolicyResponse {
		return &v
	}).(DeleteRetentionPolicyResponsePtrOutput)
}

func (o DeleteRetentionPolicyResponseOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeleteRetentionPolicyResponse) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o DeleteRetentionPolicyResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DeleteRetentionPolicyResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type DeleteRetentionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (DeleteRetentionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeleteRetentionPolicyResponse)(nil)).Elem()
}

func (o DeleteRetentionPolicyResponsePtrOutput) ToDeleteRetentionPolicyResponsePtrOutput() DeleteRetentionPolicyResponsePtrOutput {
	return o
}

func (o DeleteRetentionPolicyResponsePtrOutput) ToDeleteRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) DeleteRetentionPolicyResponsePtrOutput {
	return o
}

func (o DeleteRetentionPolicyResponsePtrOutput) Elem() DeleteRetentionPolicyResponseOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicyResponse) DeleteRetentionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret DeleteRetentionPolicyResponse
		return ret
	}).(DeleteRetentionPolicyResponseOutput)
}

func (o DeleteRetentionPolicyResponsePtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o DeleteRetentionPolicyResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DeleteRetentionPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type Encryption struct {
	KeySource                       string              `pulumi:"keySource"`
	KeyVaultProperties              *KeyVaultProperties `pulumi:"keyVaultProperties"`
	RequireInfrastructureEncryption *bool               `pulumi:"requireInfrastructureEncryption"`
	Services                        *EncryptionServices `pulumi:"services"`
}


func (val *Encryption) Defaults() *Encryption {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		tmp.KeySource = "Microsoft.Storage"
	}
	return &tmp
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	KeySource                       pulumi.StringInput         `pulumi:"keySource"`
	KeyVaultProperties              KeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
	RequireInfrastructureEncryption pulumi.BoolPtrInput        `pulumi:"requireInfrastructureEncryption"`
	Services                        EncryptionServicesPtrInput `pulumi:"services"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) KeySource() pulumi.StringOutput {
	return o.ApplyT(func(v Encryption) string { return v.KeySource }).(pulumi.StringOutput)
}

func (o EncryptionOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v Encryption) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

func (o EncryptionOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Encryption) *bool { return v.RequireInfrastructureEncryption }).(pulumi.BoolPtrOutput)
}

func (o EncryptionOutput) Services() EncryptionServicesPtrOutput {
	return o.ApplyT(func(v Encryption) *EncryptionServices { return v.Services }).(EncryptionServicesPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return &v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *Encryption) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

func (o EncryptionPtrOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Encryption) *bool {
		if v == nil {
			return nil
		}
		return v.RequireInfrastructureEncryption
	}).(pulumi.BoolPtrOutput)
}

func (o EncryptionPtrOutput) Services() EncryptionServicesPtrOutput {
	return o.ApplyT(func(v *Encryption) *EncryptionServices {
		if v == nil {
			return nil
		}
		return v.Services
	}).(EncryptionServicesPtrOutput)
}

type EncryptionResponse struct {
	KeySource                       string                      `pulumi:"keySource"`
	KeyVaultProperties              *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	RequireInfrastructureEncryption *bool                       `pulumi:"requireInfrastructureEncryption"`
	Services                        *EncryptionServicesResponse `pulumi:"services"`
}


func (val *EncryptionResponse) Defaults() *EncryptionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		tmp.KeySource = "Microsoft.Storage"
	}
	return &tmp
}





type EncryptionResponseInput interface {
	pulumi.Input

	ToEncryptionResponseOutput() EncryptionResponseOutput
	ToEncryptionResponseOutputWithContext(context.Context) EncryptionResponseOutput
}

type EncryptionResponseArgs struct {
	KeySource                       pulumi.StringInput                 `pulumi:"keySource"`
	KeyVaultProperties              KeyVaultPropertiesResponsePtrInput `pulumi:"keyVaultProperties"`
	RequireInfrastructureEncryption pulumi.BoolPtrInput                `pulumi:"requireInfrastructureEncryption"`
	Services                        EncryptionServicesResponsePtrInput `pulumi:"services"`
}

func (EncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return i.ToEncryptionResponseOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput)
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput).ToEncryptionResponsePtrOutputWithContext(ctx)
}









type EncryptionResponsePtrInput interface {
	pulumi.Input

	ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput
	ToEncryptionResponsePtrOutputWithContext(context.Context) EncryptionResponsePtrOutput
}

type encryptionResponsePtrType EncryptionResponseArgs

func EncryptionResponsePtr(v *EncryptionResponseArgs) EncryptionResponsePtrInput {
	return (*encryptionResponsePtrType)(v)
}

func (*encryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponsePtrOutput)
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionResponse) *EncryptionResponse {
		return &v
	}).(EncryptionResponsePtrOutput)
}

func (o EncryptionResponseOutput) KeySource() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionResponse) string { return v.KeySource }).(pulumi.StringOutput)
}

func (o EncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionResponseOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *bool { return v.RequireInfrastructureEncryption }).(pulumi.BoolPtrOutput)
}

func (o EncryptionResponseOutput) Services() EncryptionServicesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *EncryptionServicesResponse { return v.Services }).(EncryptionServicesResponsePtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionResponse
		return ret
	}).(EncryptionResponseOutput)
}

func (o EncryptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionResponsePtrOutput) RequireInfrastructureEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequireInfrastructureEncryption
	}).(pulumi.BoolPtrOutput)
}

func (o EncryptionResponsePtrOutput) Services() EncryptionServicesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *EncryptionServicesResponse {
		if v == nil {
			return nil
		}
		return v.Services
	}).(EncryptionServicesResponsePtrOutput)
}

type EncryptionScopeKeyVaultProperties struct {
	KeyUri *string `pulumi:"keyUri"`
}





type EncryptionScopeKeyVaultPropertiesInput interface {
	pulumi.Input

	ToEncryptionScopeKeyVaultPropertiesOutput() EncryptionScopeKeyVaultPropertiesOutput
	ToEncryptionScopeKeyVaultPropertiesOutputWithContext(context.Context) EncryptionScopeKeyVaultPropertiesOutput
}

type EncryptionScopeKeyVaultPropertiesArgs struct {
	KeyUri pulumi.StringPtrInput `pulumi:"keyUri"`
}

func (EncryptionScopeKeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionScopeKeyVaultProperties)(nil)).Elem()
}

func (i EncryptionScopeKeyVaultPropertiesArgs) ToEncryptionScopeKeyVaultPropertiesOutput() EncryptionScopeKeyVaultPropertiesOutput {
	return i.ToEncryptionScopeKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i EncryptionScopeKeyVaultPropertiesArgs) ToEncryptionScopeKeyVaultPropertiesOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionScopeKeyVaultPropertiesOutput)
}

func (i EncryptionScopeKeyVaultPropertiesArgs) ToEncryptionScopeKeyVaultPropertiesPtrOutput() EncryptionScopeKeyVaultPropertiesPtrOutput {
	return i.ToEncryptionScopeKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i EncryptionScopeKeyVaultPropertiesArgs) ToEncryptionScopeKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionScopeKeyVaultPropertiesOutput).ToEncryptionScopeKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type EncryptionScopeKeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToEncryptionScopeKeyVaultPropertiesPtrOutput() EncryptionScopeKeyVaultPropertiesPtrOutput
	ToEncryptionScopeKeyVaultPropertiesPtrOutputWithContext(context.Context) EncryptionScopeKeyVaultPropertiesPtrOutput
}

type encryptionScopeKeyVaultPropertiesPtrType EncryptionScopeKeyVaultPropertiesArgs

func EncryptionScopeKeyVaultPropertiesPtr(v *EncryptionScopeKeyVaultPropertiesArgs) EncryptionScopeKeyVaultPropertiesPtrInput {
	return (*encryptionScopeKeyVaultPropertiesPtrType)(v)
}

func (*encryptionScopeKeyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionScopeKeyVaultProperties)(nil)).Elem()
}

func (i *encryptionScopeKeyVaultPropertiesPtrType) ToEncryptionScopeKeyVaultPropertiesPtrOutput() EncryptionScopeKeyVaultPropertiesPtrOutput {
	return i.ToEncryptionScopeKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *encryptionScopeKeyVaultPropertiesPtrType) ToEncryptionScopeKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionScopeKeyVaultPropertiesPtrOutput)
}

type EncryptionScopeKeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionScopeKeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionScopeKeyVaultProperties)(nil)).Elem()
}

func (o EncryptionScopeKeyVaultPropertiesOutput) ToEncryptionScopeKeyVaultPropertiesOutput() EncryptionScopeKeyVaultPropertiesOutput {
	return o
}

func (o EncryptionScopeKeyVaultPropertiesOutput) ToEncryptionScopeKeyVaultPropertiesOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesOutput {
	return o
}

func (o EncryptionScopeKeyVaultPropertiesOutput) ToEncryptionScopeKeyVaultPropertiesPtrOutput() EncryptionScopeKeyVaultPropertiesPtrOutput {
	return o.ToEncryptionScopeKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o EncryptionScopeKeyVaultPropertiesOutput) ToEncryptionScopeKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionScopeKeyVaultProperties) *EncryptionScopeKeyVaultProperties {
		return &v
	}).(EncryptionScopeKeyVaultPropertiesPtrOutput)
}

func (o EncryptionScopeKeyVaultPropertiesOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionScopeKeyVaultProperties) *string { return v.KeyUri }).(pulumi.StringPtrOutput)
}

type EncryptionScopeKeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionScopeKeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionScopeKeyVaultProperties)(nil)).Elem()
}

func (o EncryptionScopeKeyVaultPropertiesPtrOutput) ToEncryptionScopeKeyVaultPropertiesPtrOutput() EncryptionScopeKeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionScopeKeyVaultPropertiesPtrOutput) ToEncryptionScopeKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionScopeKeyVaultPropertiesPtrOutput) Elem() EncryptionScopeKeyVaultPropertiesOutput {
	return o.ApplyT(func(v *EncryptionScopeKeyVaultProperties) EncryptionScopeKeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionScopeKeyVaultProperties
		return ret
	}).(EncryptionScopeKeyVaultPropertiesOutput)
}

func (o EncryptionScopeKeyVaultPropertiesPtrOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionScopeKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyUri
	}).(pulumi.StringPtrOutput)
}

type EncryptionScopeKeyVaultPropertiesResponse struct {
	KeyUri *string `pulumi:"keyUri"`
}





type EncryptionScopeKeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToEncryptionScopeKeyVaultPropertiesResponseOutput() EncryptionScopeKeyVaultPropertiesResponseOutput
	ToEncryptionScopeKeyVaultPropertiesResponseOutputWithContext(context.Context) EncryptionScopeKeyVaultPropertiesResponseOutput
}

type EncryptionScopeKeyVaultPropertiesResponseArgs struct {
	KeyUri pulumi.StringPtrInput `pulumi:"keyUri"`
}

func (EncryptionScopeKeyVaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionScopeKeyVaultPropertiesResponse)(nil)).Elem()
}

func (i EncryptionScopeKeyVaultPropertiesResponseArgs) ToEncryptionScopeKeyVaultPropertiesResponseOutput() EncryptionScopeKeyVaultPropertiesResponseOutput {
	return i.ToEncryptionScopeKeyVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i EncryptionScopeKeyVaultPropertiesResponseArgs) ToEncryptionScopeKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionScopeKeyVaultPropertiesResponseOutput)
}

func (i EncryptionScopeKeyVaultPropertiesResponseArgs) ToEncryptionScopeKeyVaultPropertiesResponsePtrOutput() EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
	return i.ToEncryptionScopeKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionScopeKeyVaultPropertiesResponseArgs) ToEncryptionScopeKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionScopeKeyVaultPropertiesResponseOutput).ToEncryptionScopeKeyVaultPropertiesResponsePtrOutputWithContext(ctx)
}









type EncryptionScopeKeyVaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToEncryptionScopeKeyVaultPropertiesResponsePtrOutput() EncryptionScopeKeyVaultPropertiesResponsePtrOutput
	ToEncryptionScopeKeyVaultPropertiesResponsePtrOutputWithContext(context.Context) EncryptionScopeKeyVaultPropertiesResponsePtrOutput
}

type encryptionScopeKeyVaultPropertiesResponsePtrType EncryptionScopeKeyVaultPropertiesResponseArgs

func EncryptionScopeKeyVaultPropertiesResponsePtr(v *EncryptionScopeKeyVaultPropertiesResponseArgs) EncryptionScopeKeyVaultPropertiesResponsePtrInput {
	return (*encryptionScopeKeyVaultPropertiesResponsePtrType)(v)
}

func (*encryptionScopeKeyVaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionScopeKeyVaultPropertiesResponse)(nil)).Elem()
}

func (i *encryptionScopeKeyVaultPropertiesResponsePtrType) ToEncryptionScopeKeyVaultPropertiesResponsePtrOutput() EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
	return i.ToEncryptionScopeKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionScopeKeyVaultPropertiesResponsePtrType) ToEncryptionScopeKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionScopeKeyVaultPropertiesResponsePtrOutput)
}

type EncryptionScopeKeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionScopeKeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionScopeKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o EncryptionScopeKeyVaultPropertiesResponseOutput) ToEncryptionScopeKeyVaultPropertiesResponseOutput() EncryptionScopeKeyVaultPropertiesResponseOutput {
	return o
}

func (o EncryptionScopeKeyVaultPropertiesResponseOutput) ToEncryptionScopeKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesResponseOutput {
	return o
}

func (o EncryptionScopeKeyVaultPropertiesResponseOutput) ToEncryptionScopeKeyVaultPropertiesResponsePtrOutput() EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
	return o.ToEncryptionScopeKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionScopeKeyVaultPropertiesResponseOutput) ToEncryptionScopeKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionScopeKeyVaultPropertiesResponse) *EncryptionScopeKeyVaultPropertiesResponse {
		return &v
	}).(EncryptionScopeKeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionScopeKeyVaultPropertiesResponseOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionScopeKeyVaultPropertiesResponse) *string { return v.KeyUri }).(pulumi.StringPtrOutput)
}

type EncryptionScopeKeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionScopeKeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionScopeKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o EncryptionScopeKeyVaultPropertiesResponsePtrOutput) ToEncryptionScopeKeyVaultPropertiesResponsePtrOutput() EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionScopeKeyVaultPropertiesResponsePtrOutput) ToEncryptionScopeKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionScopeKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionScopeKeyVaultPropertiesResponsePtrOutput) Elem() EncryptionScopeKeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *EncryptionScopeKeyVaultPropertiesResponse) EncryptionScopeKeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionScopeKeyVaultPropertiesResponse
		return ret
	}).(EncryptionScopeKeyVaultPropertiesResponseOutput)
}

func (o EncryptionScopeKeyVaultPropertiesResponsePtrOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionScopeKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyUri
	}).(pulumi.StringPtrOutput)
}

type EncryptionService struct {
	Enabled *bool   `pulumi:"enabled"`
	KeyType *string `pulumi:"keyType"`
}





type EncryptionServiceInput interface {
	pulumi.Input

	ToEncryptionServiceOutput() EncryptionServiceOutput
	ToEncryptionServiceOutputWithContext(context.Context) EncryptionServiceOutput
}

type EncryptionServiceArgs struct {
	Enabled pulumi.BoolPtrInput   `pulumi:"enabled"`
	KeyType pulumi.StringPtrInput `pulumi:"keyType"`
}

func (EncryptionServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionService)(nil)).Elem()
}

func (i EncryptionServiceArgs) ToEncryptionServiceOutput() EncryptionServiceOutput {
	return i.ToEncryptionServiceOutputWithContext(context.Background())
}

func (i EncryptionServiceArgs) ToEncryptionServiceOutputWithContext(ctx context.Context) EncryptionServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceOutput)
}

func (i EncryptionServiceArgs) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return i.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (i EncryptionServiceArgs) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceOutput).ToEncryptionServicePtrOutputWithContext(ctx)
}









type EncryptionServicePtrInput interface {
	pulumi.Input

	ToEncryptionServicePtrOutput() EncryptionServicePtrOutput
	ToEncryptionServicePtrOutputWithContext(context.Context) EncryptionServicePtrOutput
}

type encryptionServicePtrType EncryptionServiceArgs

func EncryptionServicePtr(v *EncryptionServiceArgs) EncryptionServicePtrInput {
	return (*encryptionServicePtrType)(v)
}

func (*encryptionServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionService)(nil)).Elem()
}

func (i *encryptionServicePtrType) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return i.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (i *encryptionServicePtrType) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicePtrOutput)
}

type EncryptionServiceOutput struct{ *pulumi.OutputState }

func (EncryptionServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionService)(nil)).Elem()
}

func (o EncryptionServiceOutput) ToEncryptionServiceOutput() EncryptionServiceOutput {
	return o
}

func (o EncryptionServiceOutput) ToEncryptionServiceOutputWithContext(ctx context.Context) EncryptionServiceOutput {
	return o
}

func (o EncryptionServiceOutput) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return o.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (o EncryptionServiceOutput) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionService) *EncryptionService {
		return &v
	}).(EncryptionServicePtrOutput)
}

func (o EncryptionServiceOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionService) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o EncryptionServiceOutput) KeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionService) *string { return v.KeyType }).(pulumi.StringPtrOutput)
}

type EncryptionServicePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionService)(nil)).Elem()
}

func (o EncryptionServicePtrOutput) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return o
}

func (o EncryptionServicePtrOutput) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return o
}

func (o EncryptionServicePtrOutput) Elem() EncryptionServiceOutput {
	return o.ApplyT(func(v *EncryptionService) EncryptionService {
		if v != nil {
			return *v
		}
		var ret EncryptionService
		return ret
	}).(EncryptionServiceOutput)
}

func (o EncryptionServicePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionService) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o EncryptionServicePtrOutput) KeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionService) *string {
		if v == nil {
			return nil
		}
		return v.KeyType
	}).(pulumi.StringPtrOutput)
}

type EncryptionServiceResponse struct {
	Enabled         *bool   `pulumi:"enabled"`
	KeyType         *string `pulumi:"keyType"`
	LastEnabledTime string  `pulumi:"lastEnabledTime"`
}





type EncryptionServiceResponseInput interface {
	pulumi.Input

	ToEncryptionServiceResponseOutput() EncryptionServiceResponseOutput
	ToEncryptionServiceResponseOutputWithContext(context.Context) EncryptionServiceResponseOutput
}

type EncryptionServiceResponseArgs struct {
	Enabled         pulumi.BoolPtrInput   `pulumi:"enabled"`
	KeyType         pulumi.StringPtrInput `pulumi:"keyType"`
	LastEnabledTime pulumi.StringInput    `pulumi:"lastEnabledTime"`
}

func (EncryptionServiceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServiceResponse)(nil)).Elem()
}

func (i EncryptionServiceResponseArgs) ToEncryptionServiceResponseOutput() EncryptionServiceResponseOutput {
	return i.ToEncryptionServiceResponseOutputWithContext(context.Background())
}

func (i EncryptionServiceResponseArgs) ToEncryptionServiceResponseOutputWithContext(ctx context.Context) EncryptionServiceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceResponseOutput)
}

func (i EncryptionServiceResponseArgs) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return i.ToEncryptionServiceResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionServiceResponseArgs) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceResponseOutput).ToEncryptionServiceResponsePtrOutputWithContext(ctx)
}









type EncryptionServiceResponsePtrInput interface {
	pulumi.Input

	ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput
	ToEncryptionServiceResponsePtrOutputWithContext(context.Context) EncryptionServiceResponsePtrOutput
}

type encryptionServiceResponsePtrType EncryptionServiceResponseArgs

func EncryptionServiceResponsePtr(v *EncryptionServiceResponseArgs) EncryptionServiceResponsePtrInput {
	return (*encryptionServiceResponsePtrType)(v)
}

func (*encryptionServiceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServiceResponse)(nil)).Elem()
}

func (i *encryptionServiceResponsePtrType) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return i.ToEncryptionServiceResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionServiceResponsePtrType) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceResponsePtrOutput)
}

type EncryptionServiceResponseOutput struct{ *pulumi.OutputState }

func (EncryptionServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServiceResponse)(nil)).Elem()
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponseOutput() EncryptionServiceResponseOutput {
	return o
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponseOutputWithContext(ctx context.Context) EncryptionServiceResponseOutput {
	return o
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return o.ToEncryptionServiceResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionServiceResponse) *EncryptionServiceResponse {
		return &v
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServiceResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o EncryptionServiceResponseOutput) KeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) *string { return v.KeyType }).(pulumi.StringPtrOutput)
}

func (o EncryptionServiceResponseOutput) LastEnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) string { return v.LastEnabledTime }).(pulumi.StringOutput)
}

type EncryptionServiceResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServiceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServiceResponse)(nil)).Elem()
}

func (o EncryptionServiceResponsePtrOutput) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return o
}

func (o EncryptionServiceResponsePtrOutput) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return o
}

func (o EncryptionServiceResponsePtrOutput) Elem() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) EncryptionServiceResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionServiceResponse
		return ret
	}).(EncryptionServiceResponseOutput)
}

func (o EncryptionServiceResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o EncryptionServiceResponsePtrOutput) KeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyType
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionServiceResponsePtrOutput) LastEnabledTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastEnabledTime
	}).(pulumi.StringPtrOutput)
}

type EncryptionServices struct {
	Blob  *EncryptionService `pulumi:"blob"`
	File  *EncryptionService `pulumi:"file"`
	Queue *EncryptionService `pulumi:"queue"`
	Table *EncryptionService `pulumi:"table"`
}





type EncryptionServicesInput interface {
	pulumi.Input

	ToEncryptionServicesOutput() EncryptionServicesOutput
	ToEncryptionServicesOutputWithContext(context.Context) EncryptionServicesOutput
}

type EncryptionServicesArgs struct {
	Blob  EncryptionServicePtrInput `pulumi:"blob"`
	File  EncryptionServicePtrInput `pulumi:"file"`
	Queue EncryptionServicePtrInput `pulumi:"queue"`
	Table EncryptionServicePtrInput `pulumi:"table"`
}

func (EncryptionServicesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServices)(nil)).Elem()
}

func (i EncryptionServicesArgs) ToEncryptionServicesOutput() EncryptionServicesOutput {
	return i.ToEncryptionServicesOutputWithContext(context.Background())
}

func (i EncryptionServicesArgs) ToEncryptionServicesOutputWithContext(ctx context.Context) EncryptionServicesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesOutput)
}

func (i EncryptionServicesArgs) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return i.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (i EncryptionServicesArgs) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesOutput).ToEncryptionServicesPtrOutputWithContext(ctx)
}









type EncryptionServicesPtrInput interface {
	pulumi.Input

	ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput
	ToEncryptionServicesPtrOutputWithContext(context.Context) EncryptionServicesPtrOutput
}

type encryptionServicesPtrType EncryptionServicesArgs

func EncryptionServicesPtr(v *EncryptionServicesArgs) EncryptionServicesPtrInput {
	return (*encryptionServicesPtrType)(v)
}

func (*encryptionServicesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServices)(nil)).Elem()
}

func (i *encryptionServicesPtrType) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return i.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (i *encryptionServicesPtrType) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesPtrOutput)
}

type EncryptionServicesOutput struct{ *pulumi.OutputState }

func (EncryptionServicesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServices)(nil)).Elem()
}

func (o EncryptionServicesOutput) ToEncryptionServicesOutput() EncryptionServicesOutput {
	return o
}

func (o EncryptionServicesOutput) ToEncryptionServicesOutputWithContext(ctx context.Context) EncryptionServicesOutput {
	return o
}

func (o EncryptionServicesOutput) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return o.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (o EncryptionServicesOutput) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionServices) *EncryptionServices {
		return &v
	}).(EncryptionServicesPtrOutput)
}

func (o EncryptionServicesOutput) Blob() EncryptionServicePtrOutput {
	return o.ApplyT(func(v EncryptionServices) *EncryptionService { return v.Blob }).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesOutput) File() EncryptionServicePtrOutput {
	return o.ApplyT(func(v EncryptionServices) *EncryptionService { return v.File }).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesOutput) Queue() EncryptionServicePtrOutput {
	return o.ApplyT(func(v EncryptionServices) *EncryptionService { return v.Queue }).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesOutput) Table() EncryptionServicePtrOutput {
	return o.ApplyT(func(v EncryptionServices) *EncryptionService { return v.Table }).(EncryptionServicePtrOutput)
}

type EncryptionServicesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServices)(nil)).Elem()
}

func (o EncryptionServicesPtrOutput) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return o
}

func (o EncryptionServicesPtrOutput) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return o
}

func (o EncryptionServicesPtrOutput) Elem() EncryptionServicesOutput {
	return o.ApplyT(func(v *EncryptionServices) EncryptionServices {
		if v != nil {
			return *v
		}
		var ret EncryptionServices
		return ret
	}).(EncryptionServicesOutput)
}

func (o EncryptionServicesPtrOutput) Blob() EncryptionServicePtrOutput {
	return o.ApplyT(func(v *EncryptionServices) *EncryptionService {
		if v == nil {
			return nil
		}
		return v.Blob
	}).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesPtrOutput) File() EncryptionServicePtrOutput {
	return o.ApplyT(func(v *EncryptionServices) *EncryptionService {
		if v == nil {
			return nil
		}
		return v.File
	}).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesPtrOutput) Queue() EncryptionServicePtrOutput {
	return o.ApplyT(func(v *EncryptionServices) *EncryptionService {
		if v == nil {
			return nil
		}
		return v.Queue
	}).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesPtrOutput) Table() EncryptionServicePtrOutput {
	return o.ApplyT(func(v *EncryptionServices) *EncryptionService {
		if v == nil {
			return nil
		}
		return v.Table
	}).(EncryptionServicePtrOutput)
}

type EncryptionServicesResponse struct {
	Blob  *EncryptionServiceResponse `pulumi:"blob"`
	File  *EncryptionServiceResponse `pulumi:"file"`
	Queue *EncryptionServiceResponse `pulumi:"queue"`
	Table *EncryptionServiceResponse `pulumi:"table"`
}





type EncryptionServicesResponseInput interface {
	pulumi.Input

	ToEncryptionServicesResponseOutput() EncryptionServicesResponseOutput
	ToEncryptionServicesResponseOutputWithContext(context.Context) EncryptionServicesResponseOutput
}

type EncryptionServicesResponseArgs struct {
	Blob  EncryptionServiceResponsePtrInput `pulumi:"blob"`
	File  EncryptionServiceResponsePtrInput `pulumi:"file"`
	Queue EncryptionServiceResponsePtrInput `pulumi:"queue"`
	Table EncryptionServiceResponsePtrInput `pulumi:"table"`
}

func (EncryptionServicesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServicesResponse)(nil)).Elem()
}

func (i EncryptionServicesResponseArgs) ToEncryptionServicesResponseOutput() EncryptionServicesResponseOutput {
	return i.ToEncryptionServicesResponseOutputWithContext(context.Background())
}

func (i EncryptionServicesResponseArgs) ToEncryptionServicesResponseOutputWithContext(ctx context.Context) EncryptionServicesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesResponseOutput)
}

func (i EncryptionServicesResponseArgs) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return i.ToEncryptionServicesResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionServicesResponseArgs) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesResponseOutput).ToEncryptionServicesResponsePtrOutputWithContext(ctx)
}









type EncryptionServicesResponsePtrInput interface {
	pulumi.Input

	ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput
	ToEncryptionServicesResponsePtrOutputWithContext(context.Context) EncryptionServicesResponsePtrOutput
}

type encryptionServicesResponsePtrType EncryptionServicesResponseArgs

func EncryptionServicesResponsePtr(v *EncryptionServicesResponseArgs) EncryptionServicesResponsePtrInput {
	return (*encryptionServicesResponsePtrType)(v)
}

func (*encryptionServicesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServicesResponse)(nil)).Elem()
}

func (i *encryptionServicesResponsePtrType) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return i.ToEncryptionServicesResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionServicesResponsePtrType) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesResponsePtrOutput)
}

type EncryptionServicesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionServicesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServicesResponse)(nil)).Elem()
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponseOutput() EncryptionServicesResponseOutput {
	return o
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponseOutputWithContext(ctx context.Context) EncryptionServicesResponseOutput {
	return o
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return o.ToEncryptionServicesResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionServicesResponse) *EncryptionServicesResponse {
		return &v
	}).(EncryptionServicesResponsePtrOutput)
}

func (o EncryptionServicesResponseOutput) Blob() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.Blob }).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponseOutput) File() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.File }).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponseOutput) Queue() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.Queue }).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponseOutput) Table() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.Table }).(EncryptionServiceResponsePtrOutput)
}

type EncryptionServicesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServicesResponse)(nil)).Elem()
}

func (o EncryptionServicesResponsePtrOutput) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return o
}

func (o EncryptionServicesResponsePtrOutput) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return o
}

func (o EncryptionServicesResponsePtrOutput) Elem() EncryptionServicesResponseOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) EncryptionServicesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionServicesResponse
		return ret
	}).(EncryptionServicesResponseOutput)
}

func (o EncryptionServicesResponsePtrOutput) Blob() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.Blob
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) File() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.File
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) Queue() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.Queue
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) Table() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.Table
	}).(EncryptionServiceResponsePtrOutput)
}

type EndpointsResponse struct {
	Blob               string                                    `pulumi:"blob"`
	Dfs                string                                    `pulumi:"dfs"`
	File               string                                    `pulumi:"file"`
	InternetEndpoints  *StorageAccountInternetEndpointsResponse  `pulumi:"internetEndpoints"`
	MicrosoftEndpoints *StorageAccountMicrosoftEndpointsResponse `pulumi:"microsoftEndpoints"`
	Queue              string                                    `pulumi:"queue"`
	Table              string                                    `pulumi:"table"`
	Web                string                                    `pulumi:"web"`
}





type EndpointsResponseInput interface {
	pulumi.Input

	ToEndpointsResponseOutput() EndpointsResponseOutput
	ToEndpointsResponseOutputWithContext(context.Context) EndpointsResponseOutput
}

type EndpointsResponseArgs struct {
	Blob               pulumi.StringInput                               `pulumi:"blob"`
	Dfs                pulumi.StringInput                               `pulumi:"dfs"`
	File               pulumi.StringInput                               `pulumi:"file"`
	InternetEndpoints  StorageAccountInternetEndpointsResponsePtrInput  `pulumi:"internetEndpoints"`
	MicrosoftEndpoints StorageAccountMicrosoftEndpointsResponsePtrInput `pulumi:"microsoftEndpoints"`
	Queue              pulumi.StringInput                               `pulumi:"queue"`
	Table              pulumi.StringInput                               `pulumi:"table"`
	Web                pulumi.StringInput                               `pulumi:"web"`
}

func (EndpointsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (i EndpointsResponseArgs) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return i.ToEndpointsResponseOutputWithContext(context.Background())
}

func (i EndpointsResponseArgs) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponseOutput)
}

func (i EndpointsResponseArgs) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return i.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i EndpointsResponseArgs) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponseOutput).ToEndpointsResponsePtrOutputWithContext(ctx)
}









type EndpointsResponsePtrInput interface {
	pulumi.Input

	ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput
	ToEndpointsResponsePtrOutputWithContext(context.Context) EndpointsResponsePtrOutput
}

type endpointsResponsePtrType EndpointsResponseArgs

func EndpointsResponsePtr(v *EndpointsResponseArgs) EndpointsResponsePtrInput {
	return (*endpointsResponsePtrType)(v)
}

func (*endpointsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsResponse)(nil)).Elem()
}

func (i *endpointsResponsePtrType) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return i.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i *endpointsResponsePtrType) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponsePtrOutput)
}

type EndpointsResponseOutput struct{ *pulumi.OutputState }

func (EndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return o.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (o EndpointsResponseOutput) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointsResponse) *EndpointsResponse {
		return &v
	}).(EndpointsResponsePtrOutput)
}

func (o EndpointsResponseOutput) Blob() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Blob }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Dfs() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Dfs }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.File }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) InternetEndpoints() StorageAccountInternetEndpointsResponsePtrOutput {
	return o.ApplyT(func(v EndpointsResponse) *StorageAccountInternetEndpointsResponse { return v.InternetEndpoints }).(StorageAccountInternetEndpointsResponsePtrOutput)
}

func (o EndpointsResponseOutput) MicrosoftEndpoints() StorageAccountMicrosoftEndpointsResponsePtrOutput {
	return o.ApplyT(func(v EndpointsResponse) *StorageAccountMicrosoftEndpointsResponse { return v.MicrosoftEndpoints }).(StorageAccountMicrosoftEndpointsResponsePtrOutput)
}

func (o EndpointsResponseOutput) Queue() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Queue }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Table() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Table }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Web() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Web }).(pulumi.StringOutput)
}

type EndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (EndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) Elem() EndpointsResponseOutput {
	return o.ApplyT(func(v *EndpointsResponse) EndpointsResponse {
		if v != nil {
			return *v
		}
		var ret EndpointsResponse
		return ret
	}).(EndpointsResponseOutput)
}

func (o EndpointsResponsePtrOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Blob
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Dfs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Dfs
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.File
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) InternetEndpoints() StorageAccountInternetEndpointsResponsePtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *StorageAccountInternetEndpointsResponse {
		if v == nil {
			return nil
		}
		return v.InternetEndpoints
	}).(StorageAccountInternetEndpointsResponsePtrOutput)
}

func (o EndpointsResponsePtrOutput) MicrosoftEndpoints() StorageAccountMicrosoftEndpointsResponsePtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *StorageAccountMicrosoftEndpointsResponse {
		if v == nil {
			return nil
		}
		return v.MicrosoftEndpoints
	}).(StorageAccountMicrosoftEndpointsResponsePtrOutput)
}

func (o EndpointsResponsePtrOutput) Queue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Queue
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Table
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Web() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Web
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (i ExtendedLocationArgs) ToExtendedLocationOutput() ExtendedLocationOutput {
	return i.ToExtendedLocationOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput)
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput).ToExtendedLocationPtrOutputWithContext(ctx)
}









type ExtendedLocationPtrInput interface {
	pulumi.Input

	ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput
	ToExtendedLocationPtrOutputWithContext(context.Context) ExtendedLocationPtrOutput
}

type extendedLocationPtrType ExtendedLocationArgs

func ExtendedLocationPtr(v *ExtendedLocationArgs) ExtendedLocationPtrInput {
	return (*extendedLocationPtrType)(v)
}

func (*extendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationPtrOutput)
}

type ExtendedLocationOutput struct{ *pulumi.OutputState }

func (ExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationOutput) ToExtendedLocationOutput() ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocation) *ExtendedLocation {
		return &v
	}).(ExtendedLocationPtrOutput)
}

func (o ExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) Elem() ExtendedLocationOutput {
	return o.ApplyT(func(v *ExtendedLocation) ExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ExtendedLocation
		return ret
	}).(ExtendedLocationOutput)
}

func (o ExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationResponseInput interface {
	pulumi.Input

	ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput
	ToExtendedLocationResponseOutputWithContext(context.Context) ExtendedLocationResponseOutput
}

type ExtendedLocationResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return i.ToExtendedLocationResponseOutputWithContext(context.Background())
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponseOutput)
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return i.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (i ExtendedLocationResponseArgs) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponseOutput).ToExtendedLocationResponsePtrOutputWithContext(ctx)
}









type ExtendedLocationResponsePtrInput interface {
	pulumi.Input

	ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput
	ToExtendedLocationResponsePtrOutputWithContext(context.Context) ExtendedLocationResponsePtrOutput
}

type extendedLocationResponsePtrType ExtendedLocationResponseArgs

func ExtendedLocationResponsePtr(v *ExtendedLocationResponseArgs) ExtendedLocationResponsePtrInput {
	return (*extendedLocationResponsePtrType)(v)
}

func (*extendedLocationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (i *extendedLocationResponsePtrType) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return i.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (i *extendedLocationResponsePtrType) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationResponsePtrOutput)
}

type ExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o.ToExtendedLocationResponsePtrOutputWithContext(context.Background())
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocationResponse) *ExtendedLocationResponse {
		return &v
	}).(ExtendedLocationResponsePtrOutput)
}

func (o ExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) Elem() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) ExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationResponse
		return ret
	}).(ExtendedLocationResponseOutput)
}

func (o ExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type GeoReplicationStatsResponse struct {
	CanFailover  bool   `pulumi:"canFailover"`
	LastSyncTime string `pulumi:"lastSyncTime"`
	Status       string `pulumi:"status"`
}





type GeoReplicationStatsResponseInput interface {
	pulumi.Input

	ToGeoReplicationStatsResponseOutput() GeoReplicationStatsResponseOutput
	ToGeoReplicationStatsResponseOutputWithContext(context.Context) GeoReplicationStatsResponseOutput
}

type GeoReplicationStatsResponseArgs struct {
	CanFailover  pulumi.BoolInput   `pulumi:"canFailover"`
	LastSyncTime pulumi.StringInput `pulumi:"lastSyncTime"`
	Status       pulumi.StringInput `pulumi:"status"`
}

func (GeoReplicationStatsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoReplicationStatsResponse)(nil)).Elem()
}

func (i GeoReplicationStatsResponseArgs) ToGeoReplicationStatsResponseOutput() GeoReplicationStatsResponseOutput {
	return i.ToGeoReplicationStatsResponseOutputWithContext(context.Background())
}

func (i GeoReplicationStatsResponseArgs) ToGeoReplicationStatsResponseOutputWithContext(ctx context.Context) GeoReplicationStatsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoReplicationStatsResponseOutput)
}

func (i GeoReplicationStatsResponseArgs) ToGeoReplicationStatsResponsePtrOutput() GeoReplicationStatsResponsePtrOutput {
	return i.ToGeoReplicationStatsResponsePtrOutputWithContext(context.Background())
}

func (i GeoReplicationStatsResponseArgs) ToGeoReplicationStatsResponsePtrOutputWithContext(ctx context.Context) GeoReplicationStatsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoReplicationStatsResponseOutput).ToGeoReplicationStatsResponsePtrOutputWithContext(ctx)
}









type GeoReplicationStatsResponsePtrInput interface {
	pulumi.Input

	ToGeoReplicationStatsResponsePtrOutput() GeoReplicationStatsResponsePtrOutput
	ToGeoReplicationStatsResponsePtrOutputWithContext(context.Context) GeoReplicationStatsResponsePtrOutput
}

type geoReplicationStatsResponsePtrType GeoReplicationStatsResponseArgs

func GeoReplicationStatsResponsePtr(v *GeoReplicationStatsResponseArgs) GeoReplicationStatsResponsePtrInput {
	return (*geoReplicationStatsResponsePtrType)(v)
}

func (*geoReplicationStatsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoReplicationStatsResponse)(nil)).Elem()
}

func (i *geoReplicationStatsResponsePtrType) ToGeoReplicationStatsResponsePtrOutput() GeoReplicationStatsResponsePtrOutput {
	return i.ToGeoReplicationStatsResponsePtrOutputWithContext(context.Background())
}

func (i *geoReplicationStatsResponsePtrType) ToGeoReplicationStatsResponsePtrOutputWithContext(ctx context.Context) GeoReplicationStatsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoReplicationStatsResponsePtrOutput)
}

type GeoReplicationStatsResponseOutput struct{ *pulumi.OutputState }

func (GeoReplicationStatsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoReplicationStatsResponse)(nil)).Elem()
}

func (o GeoReplicationStatsResponseOutput) ToGeoReplicationStatsResponseOutput() GeoReplicationStatsResponseOutput {
	return o
}

func (o GeoReplicationStatsResponseOutput) ToGeoReplicationStatsResponseOutputWithContext(ctx context.Context) GeoReplicationStatsResponseOutput {
	return o
}

func (o GeoReplicationStatsResponseOutput) ToGeoReplicationStatsResponsePtrOutput() GeoReplicationStatsResponsePtrOutput {
	return o.ToGeoReplicationStatsResponsePtrOutputWithContext(context.Background())
}

func (o GeoReplicationStatsResponseOutput) ToGeoReplicationStatsResponsePtrOutputWithContext(ctx context.Context) GeoReplicationStatsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GeoReplicationStatsResponse) *GeoReplicationStatsResponse {
		return &v
	}).(GeoReplicationStatsResponsePtrOutput)
}

func (o GeoReplicationStatsResponseOutput) CanFailover() pulumi.BoolOutput {
	return o.ApplyT(func(v GeoReplicationStatsResponse) bool { return v.CanFailover }).(pulumi.BoolOutput)
}

func (o GeoReplicationStatsResponseOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v GeoReplicationStatsResponse) string { return v.LastSyncTime }).(pulumi.StringOutput)
}

func (o GeoReplicationStatsResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GeoReplicationStatsResponse) string { return v.Status }).(pulumi.StringOutput)
}

type GeoReplicationStatsResponsePtrOutput struct{ *pulumi.OutputState }

func (GeoReplicationStatsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoReplicationStatsResponse)(nil)).Elem()
}

func (o GeoReplicationStatsResponsePtrOutput) ToGeoReplicationStatsResponsePtrOutput() GeoReplicationStatsResponsePtrOutput {
	return o
}

func (o GeoReplicationStatsResponsePtrOutput) ToGeoReplicationStatsResponsePtrOutputWithContext(ctx context.Context) GeoReplicationStatsResponsePtrOutput {
	return o
}

func (o GeoReplicationStatsResponsePtrOutput) Elem() GeoReplicationStatsResponseOutput {
	return o.ApplyT(func(v *GeoReplicationStatsResponse) GeoReplicationStatsResponse {
		if v != nil {
			return *v
		}
		var ret GeoReplicationStatsResponse
		return ret
	}).(GeoReplicationStatsResponseOutput)
}

func (o GeoReplicationStatsResponsePtrOutput) CanFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GeoReplicationStatsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CanFailover
	}).(pulumi.BoolPtrOutput)
}

func (o GeoReplicationStatsResponsePtrOutput) LastSyncTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeoReplicationStatsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSyncTime
	}).(pulumi.StringPtrOutput)
}

func (o GeoReplicationStatsResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeoReplicationStatsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type IPRule struct {
	Action           *Action `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}


func (val *IPRule) Defaults() *IPRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := Action("Allow")
		tmp.Action = &action_
	}
	return &tmp
}





type IPRuleInput interface {
	pulumi.Input

	ToIPRuleOutput() IPRuleOutput
	ToIPRuleOutputWithContext(context.Context) IPRuleOutput
}

type IPRuleArgs struct {
	Action           ActionPtrInput     `pulumi:"action"`
	IPAddressOrRange pulumi.StringInput `pulumi:"iPAddressOrRange"`
}

func (IPRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (i IPRuleArgs) ToIPRuleOutput() IPRuleOutput {
	return i.ToIPRuleOutputWithContext(context.Background())
}

func (i IPRuleArgs) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleOutput)
}





type IPRuleArrayInput interface {
	pulumi.Input

	ToIPRuleArrayOutput() IPRuleArrayOutput
	ToIPRuleArrayOutputWithContext(context.Context) IPRuleArrayOutput
}

type IPRuleArray []IPRuleInput

func (IPRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (i IPRuleArray) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return i.ToIPRuleArrayOutputWithContext(context.Background())
}

func (i IPRuleArray) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleArrayOutput)
}

type IPRuleOutput struct{ *pulumi.OutputState }

func (IPRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (o IPRuleOutput) ToIPRuleOutput() IPRuleOutput {
	return o
}

func (o IPRuleOutput) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return o
}

func (o IPRuleOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v IPRule) *Action { return v.Action }).(ActionPtrOutput)
}

func (o IPRuleOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRule) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleArrayOutput struct{ *pulumi.OutputState }

func (IPRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) Index(i pulumi.IntInput) IPRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRule {
		return vs[0].([]IPRule)[vs[1].(int)]
	}).(IPRuleOutput)
}

type IPRuleResponse struct {
	Action           *string `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}


func (val *IPRuleResponse) Defaults() *IPRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type IPRuleResponseInput interface {
	pulumi.Input

	ToIPRuleResponseOutput() IPRuleResponseOutput
	ToIPRuleResponseOutputWithContext(context.Context) IPRuleResponseOutput
}

type IPRuleResponseArgs struct {
	Action           pulumi.StringPtrInput `pulumi:"action"`
	IPAddressOrRange pulumi.StringInput    `pulumi:"iPAddressOrRange"`
}

func (IPRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRuleResponse)(nil)).Elem()
}

func (i IPRuleResponseArgs) ToIPRuleResponseOutput() IPRuleResponseOutput {
	return i.ToIPRuleResponseOutputWithContext(context.Background())
}

func (i IPRuleResponseArgs) ToIPRuleResponseOutputWithContext(ctx context.Context) IPRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleResponseOutput)
}





type IPRuleResponseArrayInput interface {
	pulumi.Input

	ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput
	ToIPRuleResponseArrayOutputWithContext(context.Context) IPRuleResponseArrayOutput
}

type IPRuleResponseArray []IPRuleResponseInput

func (IPRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRuleResponse)(nil)).Elem()
}

func (i IPRuleResponseArray) ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput {
	return i.ToIPRuleResponseArrayOutputWithContext(context.Background())
}

func (i IPRuleResponseArray) ToIPRuleResponseArrayOutputWithContext(ctx context.Context) IPRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleResponseArrayOutput)
}

type IPRuleResponseOutput struct{ *pulumi.OutputState }

func (IPRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutput() IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutputWithContext(ctx context.Context) IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IPRuleResponseOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRuleResponse) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IPRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutputWithContext(ctx context.Context) IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) Index(i pulumi.IntInput) IPRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRuleResponse {
		return vs[0].([]IPRuleResponse)[vs[1].(int)]
	}).(IPRuleResponseOutput)
}

type Identity struct {
	Type IdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type IdentityTypeInput `pulumi:"type"`
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

func (o IdentityOutput) Type() IdentityTypeOutput {
	return o.ApplyT(func(v Identity) IdentityType { return v.Type }).(IdentityTypeOutput)
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

func (o IdentityPtrOutput) Type() IdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *IdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(IdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
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

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
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
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ImmutabilityPolicyPropertiesResponse struct {
	AllowProtectedAppendWrites            *bool                           `pulumi:"allowProtectedAppendWrites"`
	Etag                                  string                          `pulumi:"etag"`
	ImmutabilityPeriodSinceCreationInDays *int                            `pulumi:"immutabilityPeriodSinceCreationInDays"`
	State                                 string                          `pulumi:"state"`
	UpdateHistory                         []UpdateHistoryPropertyResponse `pulumi:"updateHistory"`
}





type ImmutabilityPolicyPropertiesResponseInput interface {
	pulumi.Input

	ToImmutabilityPolicyPropertiesResponseOutput() ImmutabilityPolicyPropertiesResponseOutput
	ToImmutabilityPolicyPropertiesResponseOutputWithContext(context.Context) ImmutabilityPolicyPropertiesResponseOutput
}

type ImmutabilityPolicyPropertiesResponseArgs struct {
	AllowProtectedAppendWrites            pulumi.BoolPtrInput                     `pulumi:"allowProtectedAppendWrites"`
	Etag                                  pulumi.StringInput                      `pulumi:"etag"`
	ImmutabilityPeriodSinceCreationInDays pulumi.IntPtrInput                      `pulumi:"immutabilityPeriodSinceCreationInDays"`
	State                                 pulumi.StringInput                      `pulumi:"state"`
	UpdateHistory                         UpdateHistoryPropertyResponseArrayInput `pulumi:"updateHistory"`
}

func (ImmutabilityPolicyPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImmutabilityPolicyPropertiesResponse)(nil)).Elem()
}

func (i ImmutabilityPolicyPropertiesResponseArgs) ToImmutabilityPolicyPropertiesResponseOutput() ImmutabilityPolicyPropertiesResponseOutput {
	return i.ToImmutabilityPolicyPropertiesResponseOutputWithContext(context.Background())
}

func (i ImmutabilityPolicyPropertiesResponseArgs) ToImmutabilityPolicyPropertiesResponseOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilityPolicyPropertiesResponseOutput)
}

func (i ImmutabilityPolicyPropertiesResponseArgs) ToImmutabilityPolicyPropertiesResponsePtrOutput() ImmutabilityPolicyPropertiesResponsePtrOutput {
	return i.ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ImmutabilityPolicyPropertiesResponseArgs) ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilityPolicyPropertiesResponseOutput).ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(ctx)
}









type ImmutabilityPolicyPropertiesResponsePtrInput interface {
	pulumi.Input

	ToImmutabilityPolicyPropertiesResponsePtrOutput() ImmutabilityPolicyPropertiesResponsePtrOutput
	ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(context.Context) ImmutabilityPolicyPropertiesResponsePtrOutput
}

type immutabilityPolicyPropertiesResponsePtrType ImmutabilityPolicyPropertiesResponseArgs

func ImmutabilityPolicyPropertiesResponsePtr(v *ImmutabilityPolicyPropertiesResponseArgs) ImmutabilityPolicyPropertiesResponsePtrInput {
	return (*immutabilityPolicyPropertiesResponsePtrType)(v)
}

func (*immutabilityPolicyPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutabilityPolicyPropertiesResponse)(nil)).Elem()
}

func (i *immutabilityPolicyPropertiesResponsePtrType) ToImmutabilityPolicyPropertiesResponsePtrOutput() ImmutabilityPolicyPropertiesResponsePtrOutput {
	return i.ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *immutabilityPolicyPropertiesResponsePtrType) ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilityPolicyPropertiesResponsePtrOutput)
}

type ImmutabilityPolicyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ImmutabilityPolicyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImmutabilityPolicyPropertiesResponse)(nil)).Elem()
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ToImmutabilityPolicyPropertiesResponseOutput() ImmutabilityPolicyPropertiesResponseOutput {
	return o
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ToImmutabilityPolicyPropertiesResponseOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponseOutput {
	return o
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ToImmutabilityPolicyPropertiesResponsePtrOutput() ImmutabilityPolicyPropertiesResponsePtrOutput {
	return o.ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImmutabilityPolicyPropertiesResponse) *ImmutabilityPolicyPropertiesResponse {
		return &v
	}).(ImmutabilityPolicyPropertiesResponsePtrOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) AllowProtectedAppendWrites() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) *bool { return v.AllowProtectedAppendWrites }).(pulumi.BoolPtrOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ImmutabilityPeriodSinceCreationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) *int { return v.ImmutabilityPeriodSinceCreationInDays }).(pulumi.IntPtrOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) UpdateHistory() UpdateHistoryPropertyResponseArrayOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) []UpdateHistoryPropertyResponse { return v.UpdateHistory }).(UpdateHistoryPropertyResponseArrayOutput)
}

type ImmutabilityPolicyPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ImmutabilityPolicyPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutabilityPolicyPropertiesResponse)(nil)).Elem()
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) ToImmutabilityPolicyPropertiesResponsePtrOutput() ImmutabilityPolicyPropertiesResponsePtrOutput {
	return o
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponsePtrOutput {
	return o
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) Elem() ImmutabilityPolicyPropertiesResponseOutput {
	return o.ApplyT(func(v *ImmutabilityPolicyPropertiesResponse) ImmutabilityPolicyPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ImmutabilityPolicyPropertiesResponse
		return ret
	}).(ImmutabilityPolicyPropertiesResponseOutput)
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) AllowProtectedAppendWrites() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ImmutabilityPolicyPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowProtectedAppendWrites
	}).(pulumi.BoolPtrOutput)
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutabilityPolicyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) ImmutabilityPeriodSinceCreationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImmutabilityPolicyPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.ImmutabilityPeriodSinceCreationInDays
	}).(pulumi.IntPtrOutput)
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutabilityPolicyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) UpdateHistory() UpdateHistoryPropertyResponseArrayOutput {
	return o.ApplyT(func(v *ImmutabilityPolicyPropertiesResponse) []UpdateHistoryPropertyResponse {
		if v == nil {
			return nil
		}
		return v.UpdateHistory
	}).(UpdateHistoryPropertyResponseArrayOutput)
}

type KeyVaultProperties struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
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

func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
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

func (o KeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	CurrentVersionedKeyIdentifier string  `pulumi:"currentVersionedKeyIdentifier"`
	KeyName                       *string `pulumi:"keyName"`
	KeyVaultUri                   *string `pulumi:"keyVaultUri"`
	KeyVersion                    *string `pulumi:"keyVersion"`
	LastKeyRotationTimestamp      string  `pulumi:"lastKeyRotationTimestamp"`
}





type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

type KeyVaultPropertiesResponseArgs struct {
	CurrentVersionedKeyIdentifier pulumi.StringInput    `pulumi:"currentVersionedKeyIdentifier"`
	KeyName                       pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUri                   pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion                    pulumi.StringPtrInput `pulumi:"keyVersion"`
	LastKeyRotationTimestamp      pulumi.StringInput    `pulumi:"lastKeyRotationTimestamp"`
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

func (o KeyVaultPropertiesResponseOutput) CurrentVersionedKeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.CurrentVersionedKeyIdentifier }).(pulumi.StringOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) LastKeyRotationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.LastKeyRotationTimestamp }).(pulumi.StringOutput)
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

func (o KeyVaultPropertiesResponsePtrOutput) CurrentVersionedKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CurrentVersionedKeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) LastKeyRotationTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastKeyRotationTimestamp
	}).(pulumi.StringPtrOutput)
}

type LastAccessTimeTrackingPolicy struct {
	BlobType                  []string `pulumi:"blobType"`
	Enable                    bool     `pulumi:"enable"`
	Name                      *string  `pulumi:"name"`
	TrackingGranularityInDays *int     `pulumi:"trackingGranularityInDays"`
}





type LastAccessTimeTrackingPolicyInput interface {
	pulumi.Input

	ToLastAccessTimeTrackingPolicyOutput() LastAccessTimeTrackingPolicyOutput
	ToLastAccessTimeTrackingPolicyOutputWithContext(context.Context) LastAccessTimeTrackingPolicyOutput
}

type LastAccessTimeTrackingPolicyArgs struct {
	BlobType                  pulumi.StringArrayInput `pulumi:"blobType"`
	Enable                    pulumi.BoolInput        `pulumi:"enable"`
	Name                      pulumi.StringPtrInput   `pulumi:"name"`
	TrackingGranularityInDays pulumi.IntPtrInput      `pulumi:"trackingGranularityInDays"`
}

func (LastAccessTimeTrackingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LastAccessTimeTrackingPolicy)(nil)).Elem()
}

func (i LastAccessTimeTrackingPolicyArgs) ToLastAccessTimeTrackingPolicyOutput() LastAccessTimeTrackingPolicyOutput {
	return i.ToLastAccessTimeTrackingPolicyOutputWithContext(context.Background())
}

func (i LastAccessTimeTrackingPolicyArgs) ToLastAccessTimeTrackingPolicyOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LastAccessTimeTrackingPolicyOutput)
}

func (i LastAccessTimeTrackingPolicyArgs) ToLastAccessTimeTrackingPolicyPtrOutput() LastAccessTimeTrackingPolicyPtrOutput {
	return i.ToLastAccessTimeTrackingPolicyPtrOutputWithContext(context.Background())
}

func (i LastAccessTimeTrackingPolicyArgs) ToLastAccessTimeTrackingPolicyPtrOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LastAccessTimeTrackingPolicyOutput).ToLastAccessTimeTrackingPolicyPtrOutputWithContext(ctx)
}









type LastAccessTimeTrackingPolicyPtrInput interface {
	pulumi.Input

	ToLastAccessTimeTrackingPolicyPtrOutput() LastAccessTimeTrackingPolicyPtrOutput
	ToLastAccessTimeTrackingPolicyPtrOutputWithContext(context.Context) LastAccessTimeTrackingPolicyPtrOutput
}

type lastAccessTimeTrackingPolicyPtrType LastAccessTimeTrackingPolicyArgs

func LastAccessTimeTrackingPolicyPtr(v *LastAccessTimeTrackingPolicyArgs) LastAccessTimeTrackingPolicyPtrInput {
	return (*lastAccessTimeTrackingPolicyPtrType)(v)
}

func (*lastAccessTimeTrackingPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LastAccessTimeTrackingPolicy)(nil)).Elem()
}

func (i *lastAccessTimeTrackingPolicyPtrType) ToLastAccessTimeTrackingPolicyPtrOutput() LastAccessTimeTrackingPolicyPtrOutput {
	return i.ToLastAccessTimeTrackingPolicyPtrOutputWithContext(context.Background())
}

func (i *lastAccessTimeTrackingPolicyPtrType) ToLastAccessTimeTrackingPolicyPtrOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LastAccessTimeTrackingPolicyPtrOutput)
}

type LastAccessTimeTrackingPolicyOutput struct{ *pulumi.OutputState }

func (LastAccessTimeTrackingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LastAccessTimeTrackingPolicy)(nil)).Elem()
}

func (o LastAccessTimeTrackingPolicyOutput) ToLastAccessTimeTrackingPolicyOutput() LastAccessTimeTrackingPolicyOutput {
	return o
}

func (o LastAccessTimeTrackingPolicyOutput) ToLastAccessTimeTrackingPolicyOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyOutput {
	return o
}

func (o LastAccessTimeTrackingPolicyOutput) ToLastAccessTimeTrackingPolicyPtrOutput() LastAccessTimeTrackingPolicyPtrOutput {
	return o.ToLastAccessTimeTrackingPolicyPtrOutputWithContext(context.Background())
}

func (o LastAccessTimeTrackingPolicyOutput) ToLastAccessTimeTrackingPolicyPtrOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LastAccessTimeTrackingPolicy) *LastAccessTimeTrackingPolicy {
		return &v
	}).(LastAccessTimeTrackingPolicyPtrOutput)
}

func (o LastAccessTimeTrackingPolicyOutput) BlobType() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LastAccessTimeTrackingPolicy) []string { return v.BlobType }).(pulumi.StringArrayOutput)
}

func (o LastAccessTimeTrackingPolicyOutput) Enable() pulumi.BoolOutput {
	return o.ApplyT(func(v LastAccessTimeTrackingPolicy) bool { return v.Enable }).(pulumi.BoolOutput)
}

func (o LastAccessTimeTrackingPolicyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LastAccessTimeTrackingPolicy) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LastAccessTimeTrackingPolicyOutput) TrackingGranularityInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LastAccessTimeTrackingPolicy) *int { return v.TrackingGranularityInDays }).(pulumi.IntPtrOutput)
}

type LastAccessTimeTrackingPolicyPtrOutput struct{ *pulumi.OutputState }

func (LastAccessTimeTrackingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LastAccessTimeTrackingPolicy)(nil)).Elem()
}

func (o LastAccessTimeTrackingPolicyPtrOutput) ToLastAccessTimeTrackingPolicyPtrOutput() LastAccessTimeTrackingPolicyPtrOutput {
	return o
}

func (o LastAccessTimeTrackingPolicyPtrOutput) ToLastAccessTimeTrackingPolicyPtrOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyPtrOutput {
	return o
}

func (o LastAccessTimeTrackingPolicyPtrOutput) Elem() LastAccessTimeTrackingPolicyOutput {
	return o.ApplyT(func(v *LastAccessTimeTrackingPolicy) LastAccessTimeTrackingPolicy {
		if v != nil {
			return *v
		}
		var ret LastAccessTimeTrackingPolicy
		return ret
	}).(LastAccessTimeTrackingPolicyOutput)
}

func (o LastAccessTimeTrackingPolicyPtrOutput) BlobType() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LastAccessTimeTrackingPolicy) []string {
		if v == nil {
			return nil
		}
		return v.BlobType
	}).(pulumi.StringArrayOutput)
}

func (o LastAccessTimeTrackingPolicyPtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LastAccessTimeTrackingPolicy) *bool {
		if v == nil {
			return nil
		}
		return &v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o LastAccessTimeTrackingPolicyPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LastAccessTimeTrackingPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o LastAccessTimeTrackingPolicyPtrOutput) TrackingGranularityInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LastAccessTimeTrackingPolicy) *int {
		if v == nil {
			return nil
		}
		return v.TrackingGranularityInDays
	}).(pulumi.IntPtrOutput)
}

type LastAccessTimeTrackingPolicyResponse struct {
	BlobType                  []string `pulumi:"blobType"`
	Enable                    bool     `pulumi:"enable"`
	Name                      *string  `pulumi:"name"`
	TrackingGranularityInDays *int     `pulumi:"trackingGranularityInDays"`
}





type LastAccessTimeTrackingPolicyResponseInput interface {
	pulumi.Input

	ToLastAccessTimeTrackingPolicyResponseOutput() LastAccessTimeTrackingPolicyResponseOutput
	ToLastAccessTimeTrackingPolicyResponseOutputWithContext(context.Context) LastAccessTimeTrackingPolicyResponseOutput
}

type LastAccessTimeTrackingPolicyResponseArgs struct {
	BlobType                  pulumi.StringArrayInput `pulumi:"blobType"`
	Enable                    pulumi.BoolInput        `pulumi:"enable"`
	Name                      pulumi.StringPtrInput   `pulumi:"name"`
	TrackingGranularityInDays pulumi.IntPtrInput      `pulumi:"trackingGranularityInDays"`
}

func (LastAccessTimeTrackingPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LastAccessTimeTrackingPolicyResponse)(nil)).Elem()
}

func (i LastAccessTimeTrackingPolicyResponseArgs) ToLastAccessTimeTrackingPolicyResponseOutput() LastAccessTimeTrackingPolicyResponseOutput {
	return i.ToLastAccessTimeTrackingPolicyResponseOutputWithContext(context.Background())
}

func (i LastAccessTimeTrackingPolicyResponseArgs) ToLastAccessTimeTrackingPolicyResponseOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LastAccessTimeTrackingPolicyResponseOutput)
}

func (i LastAccessTimeTrackingPolicyResponseArgs) ToLastAccessTimeTrackingPolicyResponsePtrOutput() LastAccessTimeTrackingPolicyResponsePtrOutput {
	return i.ToLastAccessTimeTrackingPolicyResponsePtrOutputWithContext(context.Background())
}

func (i LastAccessTimeTrackingPolicyResponseArgs) ToLastAccessTimeTrackingPolicyResponsePtrOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LastAccessTimeTrackingPolicyResponseOutput).ToLastAccessTimeTrackingPolicyResponsePtrOutputWithContext(ctx)
}









type LastAccessTimeTrackingPolicyResponsePtrInput interface {
	pulumi.Input

	ToLastAccessTimeTrackingPolicyResponsePtrOutput() LastAccessTimeTrackingPolicyResponsePtrOutput
	ToLastAccessTimeTrackingPolicyResponsePtrOutputWithContext(context.Context) LastAccessTimeTrackingPolicyResponsePtrOutput
}

type lastAccessTimeTrackingPolicyResponsePtrType LastAccessTimeTrackingPolicyResponseArgs

func LastAccessTimeTrackingPolicyResponsePtr(v *LastAccessTimeTrackingPolicyResponseArgs) LastAccessTimeTrackingPolicyResponsePtrInput {
	return (*lastAccessTimeTrackingPolicyResponsePtrType)(v)
}

func (*lastAccessTimeTrackingPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LastAccessTimeTrackingPolicyResponse)(nil)).Elem()
}

func (i *lastAccessTimeTrackingPolicyResponsePtrType) ToLastAccessTimeTrackingPolicyResponsePtrOutput() LastAccessTimeTrackingPolicyResponsePtrOutput {
	return i.ToLastAccessTimeTrackingPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *lastAccessTimeTrackingPolicyResponsePtrType) ToLastAccessTimeTrackingPolicyResponsePtrOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LastAccessTimeTrackingPolicyResponsePtrOutput)
}

type LastAccessTimeTrackingPolicyResponseOutput struct{ *pulumi.OutputState }

func (LastAccessTimeTrackingPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LastAccessTimeTrackingPolicyResponse)(nil)).Elem()
}

func (o LastAccessTimeTrackingPolicyResponseOutput) ToLastAccessTimeTrackingPolicyResponseOutput() LastAccessTimeTrackingPolicyResponseOutput {
	return o
}

func (o LastAccessTimeTrackingPolicyResponseOutput) ToLastAccessTimeTrackingPolicyResponseOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyResponseOutput {
	return o
}

func (o LastAccessTimeTrackingPolicyResponseOutput) ToLastAccessTimeTrackingPolicyResponsePtrOutput() LastAccessTimeTrackingPolicyResponsePtrOutput {
	return o.ToLastAccessTimeTrackingPolicyResponsePtrOutputWithContext(context.Background())
}

func (o LastAccessTimeTrackingPolicyResponseOutput) ToLastAccessTimeTrackingPolicyResponsePtrOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LastAccessTimeTrackingPolicyResponse) *LastAccessTimeTrackingPolicyResponse {
		return &v
	}).(LastAccessTimeTrackingPolicyResponsePtrOutput)
}

func (o LastAccessTimeTrackingPolicyResponseOutput) BlobType() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LastAccessTimeTrackingPolicyResponse) []string { return v.BlobType }).(pulumi.StringArrayOutput)
}

func (o LastAccessTimeTrackingPolicyResponseOutput) Enable() pulumi.BoolOutput {
	return o.ApplyT(func(v LastAccessTimeTrackingPolicyResponse) bool { return v.Enable }).(pulumi.BoolOutput)
}

func (o LastAccessTimeTrackingPolicyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LastAccessTimeTrackingPolicyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LastAccessTimeTrackingPolicyResponseOutput) TrackingGranularityInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LastAccessTimeTrackingPolicyResponse) *int { return v.TrackingGranularityInDays }).(pulumi.IntPtrOutput)
}

type LastAccessTimeTrackingPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (LastAccessTimeTrackingPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LastAccessTimeTrackingPolicyResponse)(nil)).Elem()
}

func (o LastAccessTimeTrackingPolicyResponsePtrOutput) ToLastAccessTimeTrackingPolicyResponsePtrOutput() LastAccessTimeTrackingPolicyResponsePtrOutput {
	return o
}

func (o LastAccessTimeTrackingPolicyResponsePtrOutput) ToLastAccessTimeTrackingPolicyResponsePtrOutputWithContext(ctx context.Context) LastAccessTimeTrackingPolicyResponsePtrOutput {
	return o
}

func (o LastAccessTimeTrackingPolicyResponsePtrOutput) Elem() LastAccessTimeTrackingPolicyResponseOutput {
	return o.ApplyT(func(v *LastAccessTimeTrackingPolicyResponse) LastAccessTimeTrackingPolicyResponse {
		if v != nil {
			return *v
		}
		var ret LastAccessTimeTrackingPolicyResponse
		return ret
	}).(LastAccessTimeTrackingPolicyResponseOutput)
}

func (o LastAccessTimeTrackingPolicyResponsePtrOutput) BlobType() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LastAccessTimeTrackingPolicyResponse) []string {
		if v == nil {
			return nil
		}
		return v.BlobType
	}).(pulumi.StringArrayOutput)
}

func (o LastAccessTimeTrackingPolicyResponsePtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LastAccessTimeTrackingPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o LastAccessTimeTrackingPolicyResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LastAccessTimeTrackingPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o LastAccessTimeTrackingPolicyResponsePtrOutput) TrackingGranularityInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LastAccessTimeTrackingPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.TrackingGranularityInDays
	}).(pulumi.IntPtrOutput)
}

type LegalHoldPropertiesResponse struct {
	HasLegalHold bool                  `pulumi:"hasLegalHold"`
	Tags         []TagPropertyResponse `pulumi:"tags"`
}





type LegalHoldPropertiesResponseInput interface {
	pulumi.Input

	ToLegalHoldPropertiesResponseOutput() LegalHoldPropertiesResponseOutput
	ToLegalHoldPropertiesResponseOutputWithContext(context.Context) LegalHoldPropertiesResponseOutput
}

type LegalHoldPropertiesResponseArgs struct {
	HasLegalHold pulumi.BoolInput              `pulumi:"hasLegalHold"`
	Tags         TagPropertyResponseArrayInput `pulumi:"tags"`
}

func (LegalHoldPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LegalHoldPropertiesResponse)(nil)).Elem()
}

func (i LegalHoldPropertiesResponseArgs) ToLegalHoldPropertiesResponseOutput() LegalHoldPropertiesResponseOutput {
	return i.ToLegalHoldPropertiesResponseOutputWithContext(context.Background())
}

func (i LegalHoldPropertiesResponseArgs) ToLegalHoldPropertiesResponseOutputWithContext(ctx context.Context) LegalHoldPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegalHoldPropertiesResponseOutput)
}

func (i LegalHoldPropertiesResponseArgs) ToLegalHoldPropertiesResponsePtrOutput() LegalHoldPropertiesResponsePtrOutput {
	return i.ToLegalHoldPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i LegalHoldPropertiesResponseArgs) ToLegalHoldPropertiesResponsePtrOutputWithContext(ctx context.Context) LegalHoldPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegalHoldPropertiesResponseOutput).ToLegalHoldPropertiesResponsePtrOutputWithContext(ctx)
}









type LegalHoldPropertiesResponsePtrInput interface {
	pulumi.Input

	ToLegalHoldPropertiesResponsePtrOutput() LegalHoldPropertiesResponsePtrOutput
	ToLegalHoldPropertiesResponsePtrOutputWithContext(context.Context) LegalHoldPropertiesResponsePtrOutput
}

type legalHoldPropertiesResponsePtrType LegalHoldPropertiesResponseArgs

func LegalHoldPropertiesResponsePtr(v *LegalHoldPropertiesResponseArgs) LegalHoldPropertiesResponsePtrInput {
	return (*legalHoldPropertiesResponsePtrType)(v)
}

func (*legalHoldPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LegalHoldPropertiesResponse)(nil)).Elem()
}

func (i *legalHoldPropertiesResponsePtrType) ToLegalHoldPropertiesResponsePtrOutput() LegalHoldPropertiesResponsePtrOutput {
	return i.ToLegalHoldPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *legalHoldPropertiesResponsePtrType) ToLegalHoldPropertiesResponsePtrOutputWithContext(ctx context.Context) LegalHoldPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegalHoldPropertiesResponsePtrOutput)
}

type LegalHoldPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LegalHoldPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LegalHoldPropertiesResponse)(nil)).Elem()
}

func (o LegalHoldPropertiesResponseOutput) ToLegalHoldPropertiesResponseOutput() LegalHoldPropertiesResponseOutput {
	return o
}

func (o LegalHoldPropertiesResponseOutput) ToLegalHoldPropertiesResponseOutputWithContext(ctx context.Context) LegalHoldPropertiesResponseOutput {
	return o
}

func (o LegalHoldPropertiesResponseOutput) ToLegalHoldPropertiesResponsePtrOutput() LegalHoldPropertiesResponsePtrOutput {
	return o.ToLegalHoldPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o LegalHoldPropertiesResponseOutput) ToLegalHoldPropertiesResponsePtrOutputWithContext(ctx context.Context) LegalHoldPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LegalHoldPropertiesResponse) *LegalHoldPropertiesResponse {
		return &v
	}).(LegalHoldPropertiesResponsePtrOutput)
}

func (o LegalHoldPropertiesResponseOutput) HasLegalHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LegalHoldPropertiesResponse) bool { return v.HasLegalHold }).(pulumi.BoolOutput)
}

func (o LegalHoldPropertiesResponseOutput) Tags() TagPropertyResponseArrayOutput {
	return o.ApplyT(func(v LegalHoldPropertiesResponse) []TagPropertyResponse { return v.Tags }).(TagPropertyResponseArrayOutput)
}

type LegalHoldPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LegalHoldPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LegalHoldPropertiesResponse)(nil)).Elem()
}

func (o LegalHoldPropertiesResponsePtrOutput) ToLegalHoldPropertiesResponsePtrOutput() LegalHoldPropertiesResponsePtrOutput {
	return o
}

func (o LegalHoldPropertiesResponsePtrOutput) ToLegalHoldPropertiesResponsePtrOutputWithContext(ctx context.Context) LegalHoldPropertiesResponsePtrOutput {
	return o
}

func (o LegalHoldPropertiesResponsePtrOutput) Elem() LegalHoldPropertiesResponseOutput {
	return o.ApplyT(func(v *LegalHoldPropertiesResponse) LegalHoldPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LegalHoldPropertiesResponse
		return ret
	}).(LegalHoldPropertiesResponseOutput)
}

func (o LegalHoldPropertiesResponsePtrOutput) HasLegalHold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LegalHoldPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.HasLegalHold
	}).(pulumi.BoolPtrOutput)
}

func (o LegalHoldPropertiesResponsePtrOutput) Tags() TagPropertyResponseArrayOutput {
	return o.ApplyT(func(v *LegalHoldPropertiesResponse) []TagPropertyResponse {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(TagPropertyResponseArrayOutput)
}

type ManagementPolicyAction struct {
	BaseBlob *ManagementPolicyBaseBlob `pulumi:"baseBlob"`
	Snapshot *ManagementPolicySnapShot `pulumi:"snapshot"`
	Version  *ManagementPolicyVersion  `pulumi:"version"`
}





type ManagementPolicyActionInput interface {
	pulumi.Input

	ToManagementPolicyActionOutput() ManagementPolicyActionOutput
	ToManagementPolicyActionOutputWithContext(context.Context) ManagementPolicyActionOutput
}

type ManagementPolicyActionArgs struct {
	BaseBlob ManagementPolicyBaseBlobPtrInput `pulumi:"baseBlob"`
	Snapshot ManagementPolicySnapShotPtrInput `pulumi:"snapshot"`
	Version  ManagementPolicyVersionPtrInput  `pulumi:"version"`
}

func (ManagementPolicyActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyAction)(nil)).Elem()
}

func (i ManagementPolicyActionArgs) ToManagementPolicyActionOutput() ManagementPolicyActionOutput {
	return i.ToManagementPolicyActionOutputWithContext(context.Background())
}

func (i ManagementPolicyActionArgs) ToManagementPolicyActionOutputWithContext(ctx context.Context) ManagementPolicyActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyActionOutput)
}

type ManagementPolicyActionOutput struct{ *pulumi.OutputState }

func (ManagementPolicyActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyAction)(nil)).Elem()
}

func (o ManagementPolicyActionOutput) ToManagementPolicyActionOutput() ManagementPolicyActionOutput {
	return o
}

func (o ManagementPolicyActionOutput) ToManagementPolicyActionOutputWithContext(ctx context.Context) ManagementPolicyActionOutput {
	return o
}

func (o ManagementPolicyActionOutput) BaseBlob() ManagementPolicyBaseBlobPtrOutput {
	return o.ApplyT(func(v ManagementPolicyAction) *ManagementPolicyBaseBlob { return v.BaseBlob }).(ManagementPolicyBaseBlobPtrOutput)
}

func (o ManagementPolicyActionOutput) Snapshot() ManagementPolicySnapShotPtrOutput {
	return o.ApplyT(func(v ManagementPolicyAction) *ManagementPolicySnapShot { return v.Snapshot }).(ManagementPolicySnapShotPtrOutput)
}

func (o ManagementPolicyActionOutput) Version() ManagementPolicyVersionPtrOutput {
	return o.ApplyT(func(v ManagementPolicyAction) *ManagementPolicyVersion { return v.Version }).(ManagementPolicyVersionPtrOutput)
}

type ManagementPolicyActionResponse struct {
	BaseBlob *ManagementPolicyBaseBlobResponse `pulumi:"baseBlob"`
	Snapshot *ManagementPolicySnapShotResponse `pulumi:"snapshot"`
	Version  *ManagementPolicyVersionResponse  `pulumi:"version"`
}





type ManagementPolicyActionResponseInput interface {
	pulumi.Input

	ToManagementPolicyActionResponseOutput() ManagementPolicyActionResponseOutput
	ToManagementPolicyActionResponseOutputWithContext(context.Context) ManagementPolicyActionResponseOutput
}

type ManagementPolicyActionResponseArgs struct {
	BaseBlob ManagementPolicyBaseBlobResponsePtrInput `pulumi:"baseBlob"`
	Snapshot ManagementPolicySnapShotResponsePtrInput `pulumi:"snapshot"`
	Version  ManagementPolicyVersionResponsePtrInput  `pulumi:"version"`
}

func (ManagementPolicyActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyActionResponse)(nil)).Elem()
}

func (i ManagementPolicyActionResponseArgs) ToManagementPolicyActionResponseOutput() ManagementPolicyActionResponseOutput {
	return i.ToManagementPolicyActionResponseOutputWithContext(context.Background())
}

func (i ManagementPolicyActionResponseArgs) ToManagementPolicyActionResponseOutputWithContext(ctx context.Context) ManagementPolicyActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyActionResponseOutput)
}

type ManagementPolicyActionResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicyActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyActionResponse)(nil)).Elem()
}

func (o ManagementPolicyActionResponseOutput) ToManagementPolicyActionResponseOutput() ManagementPolicyActionResponseOutput {
	return o
}

func (o ManagementPolicyActionResponseOutput) ToManagementPolicyActionResponseOutputWithContext(ctx context.Context) ManagementPolicyActionResponseOutput {
	return o
}

func (o ManagementPolicyActionResponseOutput) BaseBlob() ManagementPolicyBaseBlobResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyActionResponse) *ManagementPolicyBaseBlobResponse { return v.BaseBlob }).(ManagementPolicyBaseBlobResponsePtrOutput)
}

func (o ManagementPolicyActionResponseOutput) Snapshot() ManagementPolicySnapShotResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyActionResponse) *ManagementPolicySnapShotResponse { return v.Snapshot }).(ManagementPolicySnapShotResponsePtrOutput)
}

func (o ManagementPolicyActionResponseOutput) Version() ManagementPolicyVersionResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyActionResponse) *ManagementPolicyVersionResponse { return v.Version }).(ManagementPolicyVersionResponsePtrOutput)
}

type ManagementPolicyBaseBlob struct {
	Delete                      *DateAfterModification `pulumi:"delete"`
	EnableAutoTierToHotFromCool *bool                  `pulumi:"enableAutoTierToHotFromCool"`
	TierToArchive               *DateAfterModification `pulumi:"tierToArchive"`
	TierToCool                  *DateAfterModification `pulumi:"tierToCool"`
}





type ManagementPolicyBaseBlobInput interface {
	pulumi.Input

	ToManagementPolicyBaseBlobOutput() ManagementPolicyBaseBlobOutput
	ToManagementPolicyBaseBlobOutputWithContext(context.Context) ManagementPolicyBaseBlobOutput
}

type ManagementPolicyBaseBlobArgs struct {
	Delete                      DateAfterModificationPtrInput `pulumi:"delete"`
	EnableAutoTierToHotFromCool pulumi.BoolPtrInput           `pulumi:"enableAutoTierToHotFromCool"`
	TierToArchive               DateAfterModificationPtrInput `pulumi:"tierToArchive"`
	TierToCool                  DateAfterModificationPtrInput `pulumi:"tierToCool"`
}

func (ManagementPolicyBaseBlobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyBaseBlob)(nil)).Elem()
}

func (i ManagementPolicyBaseBlobArgs) ToManagementPolicyBaseBlobOutput() ManagementPolicyBaseBlobOutput {
	return i.ToManagementPolicyBaseBlobOutputWithContext(context.Background())
}

func (i ManagementPolicyBaseBlobArgs) ToManagementPolicyBaseBlobOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyBaseBlobOutput)
}

func (i ManagementPolicyBaseBlobArgs) ToManagementPolicyBaseBlobPtrOutput() ManagementPolicyBaseBlobPtrOutput {
	return i.ToManagementPolicyBaseBlobPtrOutputWithContext(context.Background())
}

func (i ManagementPolicyBaseBlobArgs) ToManagementPolicyBaseBlobPtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyBaseBlobOutput).ToManagementPolicyBaseBlobPtrOutputWithContext(ctx)
}









type ManagementPolicyBaseBlobPtrInput interface {
	pulumi.Input

	ToManagementPolicyBaseBlobPtrOutput() ManagementPolicyBaseBlobPtrOutput
	ToManagementPolicyBaseBlobPtrOutputWithContext(context.Context) ManagementPolicyBaseBlobPtrOutput
}

type managementPolicyBaseBlobPtrType ManagementPolicyBaseBlobArgs

func ManagementPolicyBaseBlobPtr(v *ManagementPolicyBaseBlobArgs) ManagementPolicyBaseBlobPtrInput {
	return (*managementPolicyBaseBlobPtrType)(v)
}

func (*managementPolicyBaseBlobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyBaseBlob)(nil)).Elem()
}

func (i *managementPolicyBaseBlobPtrType) ToManagementPolicyBaseBlobPtrOutput() ManagementPolicyBaseBlobPtrOutput {
	return i.ToManagementPolicyBaseBlobPtrOutputWithContext(context.Background())
}

func (i *managementPolicyBaseBlobPtrType) ToManagementPolicyBaseBlobPtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyBaseBlobPtrOutput)
}

type ManagementPolicyBaseBlobOutput struct{ *pulumi.OutputState }

func (ManagementPolicyBaseBlobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyBaseBlob)(nil)).Elem()
}

func (o ManagementPolicyBaseBlobOutput) ToManagementPolicyBaseBlobOutput() ManagementPolicyBaseBlobOutput {
	return o
}

func (o ManagementPolicyBaseBlobOutput) ToManagementPolicyBaseBlobOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobOutput {
	return o
}

func (o ManagementPolicyBaseBlobOutput) ToManagementPolicyBaseBlobPtrOutput() ManagementPolicyBaseBlobPtrOutput {
	return o.ToManagementPolicyBaseBlobPtrOutputWithContext(context.Background())
}

func (o ManagementPolicyBaseBlobOutput) ToManagementPolicyBaseBlobPtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicyBaseBlob) *ManagementPolicyBaseBlob {
		return &v
	}).(ManagementPolicyBaseBlobPtrOutput)
}

func (o ManagementPolicyBaseBlobOutput) Delete() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlob) *DateAfterModification { return v.Delete }).(DateAfterModificationPtrOutput)
}

func (o ManagementPolicyBaseBlobOutput) EnableAutoTierToHotFromCool() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlob) *bool { return v.EnableAutoTierToHotFromCool }).(pulumi.BoolPtrOutput)
}

func (o ManagementPolicyBaseBlobOutput) TierToArchive() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlob) *DateAfterModification { return v.TierToArchive }).(DateAfterModificationPtrOutput)
}

func (o ManagementPolicyBaseBlobOutput) TierToCool() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlob) *DateAfterModification { return v.TierToCool }).(DateAfterModificationPtrOutput)
}

type ManagementPolicyBaseBlobPtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicyBaseBlobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyBaseBlob)(nil)).Elem()
}

func (o ManagementPolicyBaseBlobPtrOutput) ToManagementPolicyBaseBlobPtrOutput() ManagementPolicyBaseBlobPtrOutput {
	return o
}

func (o ManagementPolicyBaseBlobPtrOutput) ToManagementPolicyBaseBlobPtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobPtrOutput {
	return o
}

func (o ManagementPolicyBaseBlobPtrOutput) Elem() ManagementPolicyBaseBlobOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlob) ManagementPolicyBaseBlob {
		if v != nil {
			return *v
		}
		var ret ManagementPolicyBaseBlob
		return ret
	}).(ManagementPolicyBaseBlobOutput)
}

func (o ManagementPolicyBaseBlobPtrOutput) Delete() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlob) *DateAfterModification {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(DateAfterModificationPtrOutput)
}

func (o ManagementPolicyBaseBlobPtrOutput) EnableAutoTierToHotFromCool() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlob) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutoTierToHotFromCool
	}).(pulumi.BoolPtrOutput)
}

func (o ManagementPolicyBaseBlobPtrOutput) TierToArchive() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlob) *DateAfterModification {
		if v == nil {
			return nil
		}
		return v.TierToArchive
	}).(DateAfterModificationPtrOutput)
}

func (o ManagementPolicyBaseBlobPtrOutput) TierToCool() DateAfterModificationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlob) *DateAfterModification {
		if v == nil {
			return nil
		}
		return v.TierToCool
	}).(DateAfterModificationPtrOutput)
}

type ManagementPolicyBaseBlobResponse struct {
	Delete                      *DateAfterModificationResponse `pulumi:"delete"`
	EnableAutoTierToHotFromCool *bool                          `pulumi:"enableAutoTierToHotFromCool"`
	TierToArchive               *DateAfterModificationResponse `pulumi:"tierToArchive"`
	TierToCool                  *DateAfterModificationResponse `pulumi:"tierToCool"`
}





type ManagementPolicyBaseBlobResponseInput interface {
	pulumi.Input

	ToManagementPolicyBaseBlobResponseOutput() ManagementPolicyBaseBlobResponseOutput
	ToManagementPolicyBaseBlobResponseOutputWithContext(context.Context) ManagementPolicyBaseBlobResponseOutput
}

type ManagementPolicyBaseBlobResponseArgs struct {
	Delete                      DateAfterModificationResponsePtrInput `pulumi:"delete"`
	EnableAutoTierToHotFromCool pulumi.BoolPtrInput                   `pulumi:"enableAutoTierToHotFromCool"`
	TierToArchive               DateAfterModificationResponsePtrInput `pulumi:"tierToArchive"`
	TierToCool                  DateAfterModificationResponsePtrInput `pulumi:"tierToCool"`
}

func (ManagementPolicyBaseBlobResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyBaseBlobResponse)(nil)).Elem()
}

func (i ManagementPolicyBaseBlobResponseArgs) ToManagementPolicyBaseBlobResponseOutput() ManagementPolicyBaseBlobResponseOutput {
	return i.ToManagementPolicyBaseBlobResponseOutputWithContext(context.Background())
}

func (i ManagementPolicyBaseBlobResponseArgs) ToManagementPolicyBaseBlobResponseOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyBaseBlobResponseOutput)
}

func (i ManagementPolicyBaseBlobResponseArgs) ToManagementPolicyBaseBlobResponsePtrOutput() ManagementPolicyBaseBlobResponsePtrOutput {
	return i.ToManagementPolicyBaseBlobResponsePtrOutputWithContext(context.Background())
}

func (i ManagementPolicyBaseBlobResponseArgs) ToManagementPolicyBaseBlobResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyBaseBlobResponseOutput).ToManagementPolicyBaseBlobResponsePtrOutputWithContext(ctx)
}









type ManagementPolicyBaseBlobResponsePtrInput interface {
	pulumi.Input

	ToManagementPolicyBaseBlobResponsePtrOutput() ManagementPolicyBaseBlobResponsePtrOutput
	ToManagementPolicyBaseBlobResponsePtrOutputWithContext(context.Context) ManagementPolicyBaseBlobResponsePtrOutput
}

type managementPolicyBaseBlobResponsePtrType ManagementPolicyBaseBlobResponseArgs

func ManagementPolicyBaseBlobResponsePtr(v *ManagementPolicyBaseBlobResponseArgs) ManagementPolicyBaseBlobResponsePtrInput {
	return (*managementPolicyBaseBlobResponsePtrType)(v)
}

func (*managementPolicyBaseBlobResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyBaseBlobResponse)(nil)).Elem()
}

func (i *managementPolicyBaseBlobResponsePtrType) ToManagementPolicyBaseBlobResponsePtrOutput() ManagementPolicyBaseBlobResponsePtrOutput {
	return i.ToManagementPolicyBaseBlobResponsePtrOutputWithContext(context.Background())
}

func (i *managementPolicyBaseBlobResponsePtrType) ToManagementPolicyBaseBlobResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyBaseBlobResponsePtrOutput)
}

type ManagementPolicyBaseBlobResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicyBaseBlobResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyBaseBlobResponse)(nil)).Elem()
}

func (o ManagementPolicyBaseBlobResponseOutput) ToManagementPolicyBaseBlobResponseOutput() ManagementPolicyBaseBlobResponseOutput {
	return o
}

func (o ManagementPolicyBaseBlobResponseOutput) ToManagementPolicyBaseBlobResponseOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobResponseOutput {
	return o
}

func (o ManagementPolicyBaseBlobResponseOutput) ToManagementPolicyBaseBlobResponsePtrOutput() ManagementPolicyBaseBlobResponsePtrOutput {
	return o.ToManagementPolicyBaseBlobResponsePtrOutputWithContext(context.Background())
}

func (o ManagementPolicyBaseBlobResponseOutput) ToManagementPolicyBaseBlobResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicyBaseBlobResponse) *ManagementPolicyBaseBlobResponse {
		return &v
	}).(ManagementPolicyBaseBlobResponsePtrOutput)
}

func (o ManagementPolicyBaseBlobResponseOutput) Delete() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse { return v.Delete }).(DateAfterModificationResponsePtrOutput)
}

func (o ManagementPolicyBaseBlobResponseOutput) EnableAutoTierToHotFromCool() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlobResponse) *bool { return v.EnableAutoTierToHotFromCool }).(pulumi.BoolPtrOutput)
}

func (o ManagementPolicyBaseBlobResponseOutput) TierToArchive() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse { return v.TierToArchive }).(DateAfterModificationResponsePtrOutput)
}

func (o ManagementPolicyBaseBlobResponseOutput) TierToCool() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse { return v.TierToCool }).(DateAfterModificationResponsePtrOutput)
}

type ManagementPolicyBaseBlobResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicyBaseBlobResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyBaseBlobResponse)(nil)).Elem()
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) ToManagementPolicyBaseBlobResponsePtrOutput() ManagementPolicyBaseBlobResponsePtrOutput {
	return o
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) ToManagementPolicyBaseBlobResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyBaseBlobResponsePtrOutput {
	return o
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) Elem() ManagementPolicyBaseBlobResponseOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlobResponse) ManagementPolicyBaseBlobResponse {
		if v != nil {
			return *v
		}
		var ret ManagementPolicyBaseBlobResponse
		return ret
	}).(ManagementPolicyBaseBlobResponseOutput)
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) Delete() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(DateAfterModificationResponsePtrOutput)
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) EnableAutoTierToHotFromCool() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlobResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutoTierToHotFromCool
	}).(pulumi.BoolPtrOutput)
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) TierToArchive() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse {
		if v == nil {
			return nil
		}
		return v.TierToArchive
	}).(DateAfterModificationResponsePtrOutput)
}

func (o ManagementPolicyBaseBlobResponsePtrOutput) TierToCool() DateAfterModificationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicyBaseBlobResponse) *DateAfterModificationResponse {
		if v == nil {
			return nil
		}
		return v.TierToCool
	}).(DateAfterModificationResponsePtrOutput)
}

type ManagementPolicyDefinition struct {
	Actions ManagementPolicyAction  `pulumi:"actions"`
	Filters *ManagementPolicyFilter `pulumi:"filters"`
}





type ManagementPolicyDefinitionInput interface {
	pulumi.Input

	ToManagementPolicyDefinitionOutput() ManagementPolicyDefinitionOutput
	ToManagementPolicyDefinitionOutputWithContext(context.Context) ManagementPolicyDefinitionOutput
}

type ManagementPolicyDefinitionArgs struct {
	Actions ManagementPolicyActionInput    `pulumi:"actions"`
	Filters ManagementPolicyFilterPtrInput `pulumi:"filters"`
}

func (ManagementPolicyDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyDefinition)(nil)).Elem()
}

func (i ManagementPolicyDefinitionArgs) ToManagementPolicyDefinitionOutput() ManagementPolicyDefinitionOutput {
	return i.ToManagementPolicyDefinitionOutputWithContext(context.Background())
}

func (i ManagementPolicyDefinitionArgs) ToManagementPolicyDefinitionOutputWithContext(ctx context.Context) ManagementPolicyDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyDefinitionOutput)
}

type ManagementPolicyDefinitionOutput struct{ *pulumi.OutputState }

func (ManagementPolicyDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyDefinition)(nil)).Elem()
}

func (o ManagementPolicyDefinitionOutput) ToManagementPolicyDefinitionOutput() ManagementPolicyDefinitionOutput {
	return o
}

func (o ManagementPolicyDefinitionOutput) ToManagementPolicyDefinitionOutputWithContext(ctx context.Context) ManagementPolicyDefinitionOutput {
	return o
}

func (o ManagementPolicyDefinitionOutput) Actions() ManagementPolicyActionOutput {
	return o.ApplyT(func(v ManagementPolicyDefinition) ManagementPolicyAction { return v.Actions }).(ManagementPolicyActionOutput)
}

func (o ManagementPolicyDefinitionOutput) Filters() ManagementPolicyFilterPtrOutput {
	return o.ApplyT(func(v ManagementPolicyDefinition) *ManagementPolicyFilter { return v.Filters }).(ManagementPolicyFilterPtrOutput)
}

type ManagementPolicyDefinitionResponse struct {
	Actions ManagementPolicyActionResponse  `pulumi:"actions"`
	Filters *ManagementPolicyFilterResponse `pulumi:"filters"`
}





type ManagementPolicyDefinitionResponseInput interface {
	pulumi.Input

	ToManagementPolicyDefinitionResponseOutput() ManagementPolicyDefinitionResponseOutput
	ToManagementPolicyDefinitionResponseOutputWithContext(context.Context) ManagementPolicyDefinitionResponseOutput
}

type ManagementPolicyDefinitionResponseArgs struct {
	Actions ManagementPolicyActionResponseInput    `pulumi:"actions"`
	Filters ManagementPolicyFilterResponsePtrInput `pulumi:"filters"`
}

func (ManagementPolicyDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyDefinitionResponse)(nil)).Elem()
}

func (i ManagementPolicyDefinitionResponseArgs) ToManagementPolicyDefinitionResponseOutput() ManagementPolicyDefinitionResponseOutput {
	return i.ToManagementPolicyDefinitionResponseOutputWithContext(context.Background())
}

func (i ManagementPolicyDefinitionResponseArgs) ToManagementPolicyDefinitionResponseOutputWithContext(ctx context.Context) ManagementPolicyDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyDefinitionResponseOutput)
}

type ManagementPolicyDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicyDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyDefinitionResponse)(nil)).Elem()
}

func (o ManagementPolicyDefinitionResponseOutput) ToManagementPolicyDefinitionResponseOutput() ManagementPolicyDefinitionResponseOutput {
	return o
}

func (o ManagementPolicyDefinitionResponseOutput) ToManagementPolicyDefinitionResponseOutputWithContext(ctx context.Context) ManagementPolicyDefinitionResponseOutput {
	return o
}

func (o ManagementPolicyDefinitionResponseOutput) Actions() ManagementPolicyActionResponseOutput {
	return o.ApplyT(func(v ManagementPolicyDefinitionResponse) ManagementPolicyActionResponse { return v.Actions }).(ManagementPolicyActionResponseOutput)
}

func (o ManagementPolicyDefinitionResponseOutput) Filters() ManagementPolicyFilterResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyDefinitionResponse) *ManagementPolicyFilterResponse { return v.Filters }).(ManagementPolicyFilterResponsePtrOutput)
}

type ManagementPolicyFilter struct {
	BlobIndexMatch []TagFilter `pulumi:"blobIndexMatch"`
	BlobTypes      []string    `pulumi:"blobTypes"`
	PrefixMatch    []string    `pulumi:"prefixMatch"`
}





type ManagementPolicyFilterInput interface {
	pulumi.Input

	ToManagementPolicyFilterOutput() ManagementPolicyFilterOutput
	ToManagementPolicyFilterOutputWithContext(context.Context) ManagementPolicyFilterOutput
}

type ManagementPolicyFilterArgs struct {
	BlobIndexMatch TagFilterArrayInput     `pulumi:"blobIndexMatch"`
	BlobTypes      pulumi.StringArrayInput `pulumi:"blobTypes"`
	PrefixMatch    pulumi.StringArrayInput `pulumi:"prefixMatch"`
}

func (ManagementPolicyFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyFilter)(nil)).Elem()
}

func (i ManagementPolicyFilterArgs) ToManagementPolicyFilterOutput() ManagementPolicyFilterOutput {
	return i.ToManagementPolicyFilterOutputWithContext(context.Background())
}

func (i ManagementPolicyFilterArgs) ToManagementPolicyFilterOutputWithContext(ctx context.Context) ManagementPolicyFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyFilterOutput)
}

func (i ManagementPolicyFilterArgs) ToManagementPolicyFilterPtrOutput() ManagementPolicyFilterPtrOutput {
	return i.ToManagementPolicyFilterPtrOutputWithContext(context.Background())
}

func (i ManagementPolicyFilterArgs) ToManagementPolicyFilterPtrOutputWithContext(ctx context.Context) ManagementPolicyFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyFilterOutput).ToManagementPolicyFilterPtrOutputWithContext(ctx)
}









type ManagementPolicyFilterPtrInput interface {
	pulumi.Input

	ToManagementPolicyFilterPtrOutput() ManagementPolicyFilterPtrOutput
	ToManagementPolicyFilterPtrOutputWithContext(context.Context) ManagementPolicyFilterPtrOutput
}

type managementPolicyFilterPtrType ManagementPolicyFilterArgs

func ManagementPolicyFilterPtr(v *ManagementPolicyFilterArgs) ManagementPolicyFilterPtrInput {
	return (*managementPolicyFilterPtrType)(v)
}

func (*managementPolicyFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyFilter)(nil)).Elem()
}

func (i *managementPolicyFilterPtrType) ToManagementPolicyFilterPtrOutput() ManagementPolicyFilterPtrOutput {
	return i.ToManagementPolicyFilterPtrOutputWithContext(context.Background())
}

func (i *managementPolicyFilterPtrType) ToManagementPolicyFilterPtrOutputWithContext(ctx context.Context) ManagementPolicyFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyFilterPtrOutput)
}

type ManagementPolicyFilterOutput struct{ *pulumi.OutputState }

func (ManagementPolicyFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyFilter)(nil)).Elem()
}

func (o ManagementPolicyFilterOutput) ToManagementPolicyFilterOutput() ManagementPolicyFilterOutput {
	return o
}

func (o ManagementPolicyFilterOutput) ToManagementPolicyFilterOutputWithContext(ctx context.Context) ManagementPolicyFilterOutput {
	return o
}

func (o ManagementPolicyFilterOutput) ToManagementPolicyFilterPtrOutput() ManagementPolicyFilterPtrOutput {
	return o.ToManagementPolicyFilterPtrOutputWithContext(context.Background())
}

func (o ManagementPolicyFilterOutput) ToManagementPolicyFilterPtrOutputWithContext(ctx context.Context) ManagementPolicyFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicyFilter) *ManagementPolicyFilter {
		return &v
	}).(ManagementPolicyFilterPtrOutput)
}

func (o ManagementPolicyFilterOutput) BlobIndexMatch() TagFilterArrayOutput {
	return o.ApplyT(func(v ManagementPolicyFilter) []TagFilter { return v.BlobIndexMatch }).(TagFilterArrayOutput)
}

func (o ManagementPolicyFilterOutput) BlobTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementPolicyFilter) []string { return v.BlobTypes }).(pulumi.StringArrayOutput)
}

func (o ManagementPolicyFilterOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementPolicyFilter) []string { return v.PrefixMatch }).(pulumi.StringArrayOutput)
}

type ManagementPolicyFilterPtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicyFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyFilter)(nil)).Elem()
}

func (o ManagementPolicyFilterPtrOutput) ToManagementPolicyFilterPtrOutput() ManagementPolicyFilterPtrOutput {
	return o
}

func (o ManagementPolicyFilterPtrOutput) ToManagementPolicyFilterPtrOutputWithContext(ctx context.Context) ManagementPolicyFilterPtrOutput {
	return o
}

func (o ManagementPolicyFilterPtrOutput) Elem() ManagementPolicyFilterOutput {
	return o.ApplyT(func(v *ManagementPolicyFilter) ManagementPolicyFilter {
		if v != nil {
			return *v
		}
		var ret ManagementPolicyFilter
		return ret
	}).(ManagementPolicyFilterOutput)
}

func (o ManagementPolicyFilterPtrOutput) BlobIndexMatch() TagFilterArrayOutput {
	return o.ApplyT(func(v *ManagementPolicyFilter) []TagFilter {
		if v == nil {
			return nil
		}
		return v.BlobIndexMatch
	}).(TagFilterArrayOutput)
}

func (o ManagementPolicyFilterPtrOutput) BlobTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementPolicyFilter) []string {
		if v == nil {
			return nil
		}
		return v.BlobTypes
	}).(pulumi.StringArrayOutput)
}

func (o ManagementPolicyFilterPtrOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementPolicyFilter) []string {
		if v == nil {
			return nil
		}
		return v.PrefixMatch
	}).(pulumi.StringArrayOutput)
}

type ManagementPolicyFilterResponse struct {
	BlobIndexMatch []TagFilterResponse `pulumi:"blobIndexMatch"`
	BlobTypes      []string            `pulumi:"blobTypes"`
	PrefixMatch    []string            `pulumi:"prefixMatch"`
}





type ManagementPolicyFilterResponseInput interface {
	pulumi.Input

	ToManagementPolicyFilterResponseOutput() ManagementPolicyFilterResponseOutput
	ToManagementPolicyFilterResponseOutputWithContext(context.Context) ManagementPolicyFilterResponseOutput
}

type ManagementPolicyFilterResponseArgs struct {
	BlobIndexMatch TagFilterResponseArrayInput `pulumi:"blobIndexMatch"`
	BlobTypes      pulumi.StringArrayInput     `pulumi:"blobTypes"`
	PrefixMatch    pulumi.StringArrayInput     `pulumi:"prefixMatch"`
}

func (ManagementPolicyFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyFilterResponse)(nil)).Elem()
}

func (i ManagementPolicyFilterResponseArgs) ToManagementPolicyFilterResponseOutput() ManagementPolicyFilterResponseOutput {
	return i.ToManagementPolicyFilterResponseOutputWithContext(context.Background())
}

func (i ManagementPolicyFilterResponseArgs) ToManagementPolicyFilterResponseOutputWithContext(ctx context.Context) ManagementPolicyFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyFilterResponseOutput)
}

func (i ManagementPolicyFilterResponseArgs) ToManagementPolicyFilterResponsePtrOutput() ManagementPolicyFilterResponsePtrOutput {
	return i.ToManagementPolicyFilterResponsePtrOutputWithContext(context.Background())
}

func (i ManagementPolicyFilterResponseArgs) ToManagementPolicyFilterResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyFilterResponseOutput).ToManagementPolicyFilterResponsePtrOutputWithContext(ctx)
}









type ManagementPolicyFilterResponsePtrInput interface {
	pulumi.Input

	ToManagementPolicyFilterResponsePtrOutput() ManagementPolicyFilterResponsePtrOutput
	ToManagementPolicyFilterResponsePtrOutputWithContext(context.Context) ManagementPolicyFilterResponsePtrOutput
}

type managementPolicyFilterResponsePtrType ManagementPolicyFilterResponseArgs

func ManagementPolicyFilterResponsePtr(v *ManagementPolicyFilterResponseArgs) ManagementPolicyFilterResponsePtrInput {
	return (*managementPolicyFilterResponsePtrType)(v)
}

func (*managementPolicyFilterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyFilterResponse)(nil)).Elem()
}

func (i *managementPolicyFilterResponsePtrType) ToManagementPolicyFilterResponsePtrOutput() ManagementPolicyFilterResponsePtrOutput {
	return i.ToManagementPolicyFilterResponsePtrOutputWithContext(context.Background())
}

func (i *managementPolicyFilterResponsePtrType) ToManagementPolicyFilterResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyFilterResponsePtrOutput)
}

type ManagementPolicyFilterResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicyFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyFilterResponse)(nil)).Elem()
}

func (o ManagementPolicyFilterResponseOutput) ToManagementPolicyFilterResponseOutput() ManagementPolicyFilterResponseOutput {
	return o
}

func (o ManagementPolicyFilterResponseOutput) ToManagementPolicyFilterResponseOutputWithContext(ctx context.Context) ManagementPolicyFilterResponseOutput {
	return o
}

func (o ManagementPolicyFilterResponseOutput) ToManagementPolicyFilterResponsePtrOutput() ManagementPolicyFilterResponsePtrOutput {
	return o.ToManagementPolicyFilterResponsePtrOutputWithContext(context.Background())
}

func (o ManagementPolicyFilterResponseOutput) ToManagementPolicyFilterResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyFilterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicyFilterResponse) *ManagementPolicyFilterResponse {
		return &v
	}).(ManagementPolicyFilterResponsePtrOutput)
}

func (o ManagementPolicyFilterResponseOutput) BlobIndexMatch() TagFilterResponseArrayOutput {
	return o.ApplyT(func(v ManagementPolicyFilterResponse) []TagFilterResponse { return v.BlobIndexMatch }).(TagFilterResponseArrayOutput)
}

func (o ManagementPolicyFilterResponseOutput) BlobTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementPolicyFilterResponse) []string { return v.BlobTypes }).(pulumi.StringArrayOutput)
}

func (o ManagementPolicyFilterResponseOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementPolicyFilterResponse) []string { return v.PrefixMatch }).(pulumi.StringArrayOutput)
}

type ManagementPolicyFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicyFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyFilterResponse)(nil)).Elem()
}

func (o ManagementPolicyFilterResponsePtrOutput) ToManagementPolicyFilterResponsePtrOutput() ManagementPolicyFilterResponsePtrOutput {
	return o
}

func (o ManagementPolicyFilterResponsePtrOutput) ToManagementPolicyFilterResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyFilterResponsePtrOutput {
	return o
}

func (o ManagementPolicyFilterResponsePtrOutput) Elem() ManagementPolicyFilterResponseOutput {
	return o.ApplyT(func(v *ManagementPolicyFilterResponse) ManagementPolicyFilterResponse {
		if v != nil {
			return *v
		}
		var ret ManagementPolicyFilterResponse
		return ret
	}).(ManagementPolicyFilterResponseOutput)
}

func (o ManagementPolicyFilterResponsePtrOutput) BlobIndexMatch() TagFilterResponseArrayOutput {
	return o.ApplyT(func(v *ManagementPolicyFilterResponse) []TagFilterResponse {
		if v == nil {
			return nil
		}
		return v.BlobIndexMatch
	}).(TagFilterResponseArrayOutput)
}

func (o ManagementPolicyFilterResponsePtrOutput) BlobTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementPolicyFilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.BlobTypes
	}).(pulumi.StringArrayOutput)
}

func (o ManagementPolicyFilterResponsePtrOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementPolicyFilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.PrefixMatch
	}).(pulumi.StringArrayOutput)
}

type ManagementPolicyRule struct {
	Definition ManagementPolicyDefinition `pulumi:"definition"`
	Enabled    *bool                      `pulumi:"enabled"`
	Name       string                     `pulumi:"name"`
	Type       string                     `pulumi:"type"`
}





type ManagementPolicyRuleInput interface {
	pulumi.Input

	ToManagementPolicyRuleOutput() ManagementPolicyRuleOutput
	ToManagementPolicyRuleOutputWithContext(context.Context) ManagementPolicyRuleOutput
}

type ManagementPolicyRuleArgs struct {
	Definition ManagementPolicyDefinitionInput `pulumi:"definition"`
	Enabled    pulumi.BoolPtrInput             `pulumi:"enabled"`
	Name       pulumi.StringInput              `pulumi:"name"`
	Type       pulumi.StringInput              `pulumi:"type"`
}

func (ManagementPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyRule)(nil)).Elem()
}

func (i ManagementPolicyRuleArgs) ToManagementPolicyRuleOutput() ManagementPolicyRuleOutput {
	return i.ToManagementPolicyRuleOutputWithContext(context.Background())
}

func (i ManagementPolicyRuleArgs) ToManagementPolicyRuleOutputWithContext(ctx context.Context) ManagementPolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyRuleOutput)
}





type ManagementPolicyRuleArrayInput interface {
	pulumi.Input

	ToManagementPolicyRuleArrayOutput() ManagementPolicyRuleArrayOutput
	ToManagementPolicyRuleArrayOutputWithContext(context.Context) ManagementPolicyRuleArrayOutput
}

type ManagementPolicyRuleArray []ManagementPolicyRuleInput

func (ManagementPolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementPolicyRule)(nil)).Elem()
}

func (i ManagementPolicyRuleArray) ToManagementPolicyRuleArrayOutput() ManagementPolicyRuleArrayOutput {
	return i.ToManagementPolicyRuleArrayOutputWithContext(context.Background())
}

func (i ManagementPolicyRuleArray) ToManagementPolicyRuleArrayOutputWithContext(ctx context.Context) ManagementPolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyRuleArrayOutput)
}

type ManagementPolicyRuleOutput struct{ *pulumi.OutputState }

func (ManagementPolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyRule)(nil)).Elem()
}

func (o ManagementPolicyRuleOutput) ToManagementPolicyRuleOutput() ManagementPolicyRuleOutput {
	return o
}

func (o ManagementPolicyRuleOutput) ToManagementPolicyRuleOutputWithContext(ctx context.Context) ManagementPolicyRuleOutput {
	return o
}

func (o ManagementPolicyRuleOutput) Definition() ManagementPolicyDefinitionOutput {
	return o.ApplyT(func(v ManagementPolicyRule) ManagementPolicyDefinition { return v.Definition }).(ManagementPolicyDefinitionOutput)
}

func (o ManagementPolicyRuleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagementPolicyRule) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ManagementPolicyRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementPolicyRule) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementPolicyRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementPolicyRule) string { return v.Type }).(pulumi.StringOutput)
}

type ManagementPolicyRuleArrayOutput struct{ *pulumi.OutputState }

func (ManagementPolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementPolicyRule)(nil)).Elem()
}

func (o ManagementPolicyRuleArrayOutput) ToManagementPolicyRuleArrayOutput() ManagementPolicyRuleArrayOutput {
	return o
}

func (o ManagementPolicyRuleArrayOutput) ToManagementPolicyRuleArrayOutputWithContext(ctx context.Context) ManagementPolicyRuleArrayOutput {
	return o
}

func (o ManagementPolicyRuleArrayOutput) Index(i pulumi.IntInput) ManagementPolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementPolicyRule {
		return vs[0].([]ManagementPolicyRule)[vs[1].(int)]
	}).(ManagementPolicyRuleOutput)
}

type ManagementPolicyRuleResponse struct {
	Definition ManagementPolicyDefinitionResponse `pulumi:"definition"`
	Enabled    *bool                              `pulumi:"enabled"`
	Name       string                             `pulumi:"name"`
	Type       string                             `pulumi:"type"`
}





type ManagementPolicyRuleResponseInput interface {
	pulumi.Input

	ToManagementPolicyRuleResponseOutput() ManagementPolicyRuleResponseOutput
	ToManagementPolicyRuleResponseOutputWithContext(context.Context) ManagementPolicyRuleResponseOutput
}

type ManagementPolicyRuleResponseArgs struct {
	Definition ManagementPolicyDefinitionResponseInput `pulumi:"definition"`
	Enabled    pulumi.BoolPtrInput                     `pulumi:"enabled"`
	Name       pulumi.StringInput                      `pulumi:"name"`
	Type       pulumi.StringInput                      `pulumi:"type"`
}

func (ManagementPolicyRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyRuleResponse)(nil)).Elem()
}

func (i ManagementPolicyRuleResponseArgs) ToManagementPolicyRuleResponseOutput() ManagementPolicyRuleResponseOutput {
	return i.ToManagementPolicyRuleResponseOutputWithContext(context.Background())
}

func (i ManagementPolicyRuleResponseArgs) ToManagementPolicyRuleResponseOutputWithContext(ctx context.Context) ManagementPolicyRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyRuleResponseOutput)
}





type ManagementPolicyRuleResponseArrayInput interface {
	pulumi.Input

	ToManagementPolicyRuleResponseArrayOutput() ManagementPolicyRuleResponseArrayOutput
	ToManagementPolicyRuleResponseArrayOutputWithContext(context.Context) ManagementPolicyRuleResponseArrayOutput
}

type ManagementPolicyRuleResponseArray []ManagementPolicyRuleResponseInput

func (ManagementPolicyRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementPolicyRuleResponse)(nil)).Elem()
}

func (i ManagementPolicyRuleResponseArray) ToManagementPolicyRuleResponseArrayOutput() ManagementPolicyRuleResponseArrayOutput {
	return i.ToManagementPolicyRuleResponseArrayOutputWithContext(context.Background())
}

func (i ManagementPolicyRuleResponseArray) ToManagementPolicyRuleResponseArrayOutputWithContext(ctx context.Context) ManagementPolicyRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyRuleResponseArrayOutput)
}

type ManagementPolicyRuleResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicyRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyRuleResponse)(nil)).Elem()
}

func (o ManagementPolicyRuleResponseOutput) ToManagementPolicyRuleResponseOutput() ManagementPolicyRuleResponseOutput {
	return o
}

func (o ManagementPolicyRuleResponseOutput) ToManagementPolicyRuleResponseOutputWithContext(ctx context.Context) ManagementPolicyRuleResponseOutput {
	return o
}

func (o ManagementPolicyRuleResponseOutput) Definition() ManagementPolicyDefinitionResponseOutput {
	return o.ApplyT(func(v ManagementPolicyRuleResponse) ManagementPolicyDefinitionResponse { return v.Definition }).(ManagementPolicyDefinitionResponseOutput)
}

func (o ManagementPolicyRuleResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagementPolicyRuleResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ManagementPolicyRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementPolicyRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementPolicyRuleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementPolicyRuleResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ManagementPolicyRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagementPolicyRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementPolicyRuleResponse)(nil)).Elem()
}

func (o ManagementPolicyRuleResponseArrayOutput) ToManagementPolicyRuleResponseArrayOutput() ManagementPolicyRuleResponseArrayOutput {
	return o
}

func (o ManagementPolicyRuleResponseArrayOutput) ToManagementPolicyRuleResponseArrayOutputWithContext(ctx context.Context) ManagementPolicyRuleResponseArrayOutput {
	return o
}

func (o ManagementPolicyRuleResponseArrayOutput) Index(i pulumi.IntInput) ManagementPolicyRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementPolicyRuleResponse {
		return vs[0].([]ManagementPolicyRuleResponse)[vs[1].(int)]
	}).(ManagementPolicyRuleResponseOutput)
}

type ManagementPolicySchema struct {
	Rules []ManagementPolicyRule `pulumi:"rules"`
}





type ManagementPolicySchemaInput interface {
	pulumi.Input

	ToManagementPolicySchemaOutput() ManagementPolicySchemaOutput
	ToManagementPolicySchemaOutputWithContext(context.Context) ManagementPolicySchemaOutput
}

type ManagementPolicySchemaArgs struct {
	Rules ManagementPolicyRuleArrayInput `pulumi:"rules"`
}

func (ManagementPolicySchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySchema)(nil)).Elem()
}

func (i ManagementPolicySchemaArgs) ToManagementPolicySchemaOutput() ManagementPolicySchemaOutput {
	return i.ToManagementPolicySchemaOutputWithContext(context.Background())
}

func (i ManagementPolicySchemaArgs) ToManagementPolicySchemaOutputWithContext(ctx context.Context) ManagementPolicySchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySchemaOutput)
}

func (i ManagementPolicySchemaArgs) ToManagementPolicySchemaPtrOutput() ManagementPolicySchemaPtrOutput {
	return i.ToManagementPolicySchemaPtrOutputWithContext(context.Background())
}

func (i ManagementPolicySchemaArgs) ToManagementPolicySchemaPtrOutputWithContext(ctx context.Context) ManagementPolicySchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySchemaOutput).ToManagementPolicySchemaPtrOutputWithContext(ctx)
}









type ManagementPolicySchemaPtrInput interface {
	pulumi.Input

	ToManagementPolicySchemaPtrOutput() ManagementPolicySchemaPtrOutput
	ToManagementPolicySchemaPtrOutputWithContext(context.Context) ManagementPolicySchemaPtrOutput
}

type managementPolicySchemaPtrType ManagementPolicySchemaArgs

func ManagementPolicySchemaPtr(v *ManagementPolicySchemaArgs) ManagementPolicySchemaPtrInput {
	return (*managementPolicySchemaPtrType)(v)
}

func (*managementPolicySchemaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicySchema)(nil)).Elem()
}

func (i *managementPolicySchemaPtrType) ToManagementPolicySchemaPtrOutput() ManagementPolicySchemaPtrOutput {
	return i.ToManagementPolicySchemaPtrOutputWithContext(context.Background())
}

func (i *managementPolicySchemaPtrType) ToManagementPolicySchemaPtrOutputWithContext(ctx context.Context) ManagementPolicySchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySchemaPtrOutput)
}

type ManagementPolicySchemaOutput struct{ *pulumi.OutputState }

func (ManagementPolicySchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySchema)(nil)).Elem()
}

func (o ManagementPolicySchemaOutput) ToManagementPolicySchemaOutput() ManagementPolicySchemaOutput {
	return o
}

func (o ManagementPolicySchemaOutput) ToManagementPolicySchemaOutputWithContext(ctx context.Context) ManagementPolicySchemaOutput {
	return o
}

func (o ManagementPolicySchemaOutput) ToManagementPolicySchemaPtrOutput() ManagementPolicySchemaPtrOutput {
	return o.ToManagementPolicySchemaPtrOutputWithContext(context.Background())
}

func (o ManagementPolicySchemaOutput) ToManagementPolicySchemaPtrOutputWithContext(ctx context.Context) ManagementPolicySchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicySchema) *ManagementPolicySchema {
		return &v
	}).(ManagementPolicySchemaPtrOutput)
}

func (o ManagementPolicySchemaOutput) Rules() ManagementPolicyRuleArrayOutput {
	return o.ApplyT(func(v ManagementPolicySchema) []ManagementPolicyRule { return v.Rules }).(ManagementPolicyRuleArrayOutput)
}

type ManagementPolicySchemaPtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicySchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicySchema)(nil)).Elem()
}

func (o ManagementPolicySchemaPtrOutput) ToManagementPolicySchemaPtrOutput() ManagementPolicySchemaPtrOutput {
	return o
}

func (o ManagementPolicySchemaPtrOutput) ToManagementPolicySchemaPtrOutputWithContext(ctx context.Context) ManagementPolicySchemaPtrOutput {
	return o
}

func (o ManagementPolicySchemaPtrOutput) Elem() ManagementPolicySchemaOutput {
	return o.ApplyT(func(v *ManagementPolicySchema) ManagementPolicySchema {
		if v != nil {
			return *v
		}
		var ret ManagementPolicySchema
		return ret
	}).(ManagementPolicySchemaOutput)
}

func (o ManagementPolicySchemaPtrOutput) Rules() ManagementPolicyRuleArrayOutput {
	return o.ApplyT(func(v *ManagementPolicySchema) []ManagementPolicyRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ManagementPolicyRuleArrayOutput)
}

type ManagementPolicySchemaResponse struct {
	Rules []ManagementPolicyRuleResponse `pulumi:"rules"`
}





type ManagementPolicySchemaResponseInput interface {
	pulumi.Input

	ToManagementPolicySchemaResponseOutput() ManagementPolicySchemaResponseOutput
	ToManagementPolicySchemaResponseOutputWithContext(context.Context) ManagementPolicySchemaResponseOutput
}

type ManagementPolicySchemaResponseArgs struct {
	Rules ManagementPolicyRuleResponseArrayInput `pulumi:"rules"`
}

func (ManagementPolicySchemaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySchemaResponse)(nil)).Elem()
}

func (i ManagementPolicySchemaResponseArgs) ToManagementPolicySchemaResponseOutput() ManagementPolicySchemaResponseOutput {
	return i.ToManagementPolicySchemaResponseOutputWithContext(context.Background())
}

func (i ManagementPolicySchemaResponseArgs) ToManagementPolicySchemaResponseOutputWithContext(ctx context.Context) ManagementPolicySchemaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySchemaResponseOutput)
}

func (i ManagementPolicySchemaResponseArgs) ToManagementPolicySchemaResponsePtrOutput() ManagementPolicySchemaResponsePtrOutput {
	return i.ToManagementPolicySchemaResponsePtrOutputWithContext(context.Background())
}

func (i ManagementPolicySchemaResponseArgs) ToManagementPolicySchemaResponsePtrOutputWithContext(ctx context.Context) ManagementPolicySchemaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySchemaResponseOutput).ToManagementPolicySchemaResponsePtrOutputWithContext(ctx)
}









type ManagementPolicySchemaResponsePtrInput interface {
	pulumi.Input

	ToManagementPolicySchemaResponsePtrOutput() ManagementPolicySchemaResponsePtrOutput
	ToManagementPolicySchemaResponsePtrOutputWithContext(context.Context) ManagementPolicySchemaResponsePtrOutput
}

type managementPolicySchemaResponsePtrType ManagementPolicySchemaResponseArgs

func ManagementPolicySchemaResponsePtr(v *ManagementPolicySchemaResponseArgs) ManagementPolicySchemaResponsePtrInput {
	return (*managementPolicySchemaResponsePtrType)(v)
}

func (*managementPolicySchemaResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicySchemaResponse)(nil)).Elem()
}

func (i *managementPolicySchemaResponsePtrType) ToManagementPolicySchemaResponsePtrOutput() ManagementPolicySchemaResponsePtrOutput {
	return i.ToManagementPolicySchemaResponsePtrOutputWithContext(context.Background())
}

func (i *managementPolicySchemaResponsePtrType) ToManagementPolicySchemaResponsePtrOutputWithContext(ctx context.Context) ManagementPolicySchemaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySchemaResponsePtrOutput)
}

type ManagementPolicySchemaResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicySchemaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySchemaResponse)(nil)).Elem()
}

func (o ManagementPolicySchemaResponseOutput) ToManagementPolicySchemaResponseOutput() ManagementPolicySchemaResponseOutput {
	return o
}

func (o ManagementPolicySchemaResponseOutput) ToManagementPolicySchemaResponseOutputWithContext(ctx context.Context) ManagementPolicySchemaResponseOutput {
	return o
}

func (o ManagementPolicySchemaResponseOutput) ToManagementPolicySchemaResponsePtrOutput() ManagementPolicySchemaResponsePtrOutput {
	return o.ToManagementPolicySchemaResponsePtrOutputWithContext(context.Background())
}

func (o ManagementPolicySchemaResponseOutput) ToManagementPolicySchemaResponsePtrOutputWithContext(ctx context.Context) ManagementPolicySchemaResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicySchemaResponse) *ManagementPolicySchemaResponse {
		return &v
	}).(ManagementPolicySchemaResponsePtrOutput)
}

func (o ManagementPolicySchemaResponseOutput) Rules() ManagementPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v ManagementPolicySchemaResponse) []ManagementPolicyRuleResponse { return v.Rules }).(ManagementPolicyRuleResponseArrayOutput)
}

type ManagementPolicySchemaResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicySchemaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicySchemaResponse)(nil)).Elem()
}

func (o ManagementPolicySchemaResponsePtrOutput) ToManagementPolicySchemaResponsePtrOutput() ManagementPolicySchemaResponsePtrOutput {
	return o
}

func (o ManagementPolicySchemaResponsePtrOutput) ToManagementPolicySchemaResponsePtrOutputWithContext(ctx context.Context) ManagementPolicySchemaResponsePtrOutput {
	return o
}

func (o ManagementPolicySchemaResponsePtrOutput) Elem() ManagementPolicySchemaResponseOutput {
	return o.ApplyT(func(v *ManagementPolicySchemaResponse) ManagementPolicySchemaResponse {
		if v != nil {
			return *v
		}
		var ret ManagementPolicySchemaResponse
		return ret
	}).(ManagementPolicySchemaResponseOutput)
}

func (o ManagementPolicySchemaResponsePtrOutput) Rules() ManagementPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v *ManagementPolicySchemaResponse) []ManagementPolicyRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ManagementPolicyRuleResponseArrayOutput)
}

type ManagementPolicySnapShot struct {
	Delete        *DateAfterCreation `pulumi:"delete"`
	TierToArchive *DateAfterCreation `pulumi:"tierToArchive"`
	TierToCool    *DateAfterCreation `pulumi:"tierToCool"`
}





type ManagementPolicySnapShotInput interface {
	pulumi.Input

	ToManagementPolicySnapShotOutput() ManagementPolicySnapShotOutput
	ToManagementPolicySnapShotOutputWithContext(context.Context) ManagementPolicySnapShotOutput
}

type ManagementPolicySnapShotArgs struct {
	Delete        DateAfterCreationPtrInput `pulumi:"delete"`
	TierToArchive DateAfterCreationPtrInput `pulumi:"tierToArchive"`
	TierToCool    DateAfterCreationPtrInput `pulumi:"tierToCool"`
}

func (ManagementPolicySnapShotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySnapShot)(nil)).Elem()
}

func (i ManagementPolicySnapShotArgs) ToManagementPolicySnapShotOutput() ManagementPolicySnapShotOutput {
	return i.ToManagementPolicySnapShotOutputWithContext(context.Background())
}

func (i ManagementPolicySnapShotArgs) ToManagementPolicySnapShotOutputWithContext(ctx context.Context) ManagementPolicySnapShotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySnapShotOutput)
}

func (i ManagementPolicySnapShotArgs) ToManagementPolicySnapShotPtrOutput() ManagementPolicySnapShotPtrOutput {
	return i.ToManagementPolicySnapShotPtrOutputWithContext(context.Background())
}

func (i ManagementPolicySnapShotArgs) ToManagementPolicySnapShotPtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySnapShotOutput).ToManagementPolicySnapShotPtrOutputWithContext(ctx)
}









type ManagementPolicySnapShotPtrInput interface {
	pulumi.Input

	ToManagementPolicySnapShotPtrOutput() ManagementPolicySnapShotPtrOutput
	ToManagementPolicySnapShotPtrOutputWithContext(context.Context) ManagementPolicySnapShotPtrOutput
}

type managementPolicySnapShotPtrType ManagementPolicySnapShotArgs

func ManagementPolicySnapShotPtr(v *ManagementPolicySnapShotArgs) ManagementPolicySnapShotPtrInput {
	return (*managementPolicySnapShotPtrType)(v)
}

func (*managementPolicySnapShotPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicySnapShot)(nil)).Elem()
}

func (i *managementPolicySnapShotPtrType) ToManagementPolicySnapShotPtrOutput() ManagementPolicySnapShotPtrOutput {
	return i.ToManagementPolicySnapShotPtrOutputWithContext(context.Background())
}

func (i *managementPolicySnapShotPtrType) ToManagementPolicySnapShotPtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySnapShotPtrOutput)
}

type ManagementPolicySnapShotOutput struct{ *pulumi.OutputState }

func (ManagementPolicySnapShotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySnapShot)(nil)).Elem()
}

func (o ManagementPolicySnapShotOutput) ToManagementPolicySnapShotOutput() ManagementPolicySnapShotOutput {
	return o
}

func (o ManagementPolicySnapShotOutput) ToManagementPolicySnapShotOutputWithContext(ctx context.Context) ManagementPolicySnapShotOutput {
	return o
}

func (o ManagementPolicySnapShotOutput) ToManagementPolicySnapShotPtrOutput() ManagementPolicySnapShotPtrOutput {
	return o.ToManagementPolicySnapShotPtrOutputWithContext(context.Background())
}

func (o ManagementPolicySnapShotOutput) ToManagementPolicySnapShotPtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicySnapShot) *ManagementPolicySnapShot {
		return &v
	}).(ManagementPolicySnapShotPtrOutput)
}

func (o ManagementPolicySnapShotOutput) Delete() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v ManagementPolicySnapShot) *DateAfterCreation { return v.Delete }).(DateAfterCreationPtrOutput)
}

func (o ManagementPolicySnapShotOutput) TierToArchive() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v ManagementPolicySnapShot) *DateAfterCreation { return v.TierToArchive }).(DateAfterCreationPtrOutput)
}

func (o ManagementPolicySnapShotOutput) TierToCool() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v ManagementPolicySnapShot) *DateAfterCreation { return v.TierToCool }).(DateAfterCreationPtrOutput)
}

type ManagementPolicySnapShotPtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicySnapShotPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicySnapShot)(nil)).Elem()
}

func (o ManagementPolicySnapShotPtrOutput) ToManagementPolicySnapShotPtrOutput() ManagementPolicySnapShotPtrOutput {
	return o
}

func (o ManagementPolicySnapShotPtrOutput) ToManagementPolicySnapShotPtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotPtrOutput {
	return o
}

func (o ManagementPolicySnapShotPtrOutput) Elem() ManagementPolicySnapShotOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShot) ManagementPolicySnapShot {
		if v != nil {
			return *v
		}
		var ret ManagementPolicySnapShot
		return ret
	}).(ManagementPolicySnapShotOutput)
}

func (o ManagementPolicySnapShotPtrOutput) Delete() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShot) *DateAfterCreation {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(DateAfterCreationPtrOutput)
}

func (o ManagementPolicySnapShotPtrOutput) TierToArchive() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShot) *DateAfterCreation {
		if v == nil {
			return nil
		}
		return v.TierToArchive
	}).(DateAfterCreationPtrOutput)
}

func (o ManagementPolicySnapShotPtrOutput) TierToCool() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShot) *DateAfterCreation {
		if v == nil {
			return nil
		}
		return v.TierToCool
	}).(DateAfterCreationPtrOutput)
}

type ManagementPolicySnapShotResponse struct {
	Delete        *DateAfterCreationResponse `pulumi:"delete"`
	TierToArchive *DateAfterCreationResponse `pulumi:"tierToArchive"`
	TierToCool    *DateAfterCreationResponse `pulumi:"tierToCool"`
}





type ManagementPolicySnapShotResponseInput interface {
	pulumi.Input

	ToManagementPolicySnapShotResponseOutput() ManagementPolicySnapShotResponseOutput
	ToManagementPolicySnapShotResponseOutputWithContext(context.Context) ManagementPolicySnapShotResponseOutput
}

type ManagementPolicySnapShotResponseArgs struct {
	Delete        DateAfterCreationResponsePtrInput `pulumi:"delete"`
	TierToArchive DateAfterCreationResponsePtrInput `pulumi:"tierToArchive"`
	TierToCool    DateAfterCreationResponsePtrInput `pulumi:"tierToCool"`
}

func (ManagementPolicySnapShotResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySnapShotResponse)(nil)).Elem()
}

func (i ManagementPolicySnapShotResponseArgs) ToManagementPolicySnapShotResponseOutput() ManagementPolicySnapShotResponseOutput {
	return i.ToManagementPolicySnapShotResponseOutputWithContext(context.Background())
}

func (i ManagementPolicySnapShotResponseArgs) ToManagementPolicySnapShotResponseOutputWithContext(ctx context.Context) ManagementPolicySnapShotResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySnapShotResponseOutput)
}

func (i ManagementPolicySnapShotResponseArgs) ToManagementPolicySnapShotResponsePtrOutput() ManagementPolicySnapShotResponsePtrOutput {
	return i.ToManagementPolicySnapShotResponsePtrOutputWithContext(context.Background())
}

func (i ManagementPolicySnapShotResponseArgs) ToManagementPolicySnapShotResponsePtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySnapShotResponseOutput).ToManagementPolicySnapShotResponsePtrOutputWithContext(ctx)
}









type ManagementPolicySnapShotResponsePtrInput interface {
	pulumi.Input

	ToManagementPolicySnapShotResponsePtrOutput() ManagementPolicySnapShotResponsePtrOutput
	ToManagementPolicySnapShotResponsePtrOutputWithContext(context.Context) ManagementPolicySnapShotResponsePtrOutput
}

type managementPolicySnapShotResponsePtrType ManagementPolicySnapShotResponseArgs

func ManagementPolicySnapShotResponsePtr(v *ManagementPolicySnapShotResponseArgs) ManagementPolicySnapShotResponsePtrInput {
	return (*managementPolicySnapShotResponsePtrType)(v)
}

func (*managementPolicySnapShotResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicySnapShotResponse)(nil)).Elem()
}

func (i *managementPolicySnapShotResponsePtrType) ToManagementPolicySnapShotResponsePtrOutput() ManagementPolicySnapShotResponsePtrOutput {
	return i.ToManagementPolicySnapShotResponsePtrOutputWithContext(context.Background())
}

func (i *managementPolicySnapShotResponsePtrType) ToManagementPolicySnapShotResponsePtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicySnapShotResponsePtrOutput)
}

type ManagementPolicySnapShotResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicySnapShotResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicySnapShotResponse)(nil)).Elem()
}

func (o ManagementPolicySnapShotResponseOutput) ToManagementPolicySnapShotResponseOutput() ManagementPolicySnapShotResponseOutput {
	return o
}

func (o ManagementPolicySnapShotResponseOutput) ToManagementPolicySnapShotResponseOutputWithContext(ctx context.Context) ManagementPolicySnapShotResponseOutput {
	return o
}

func (o ManagementPolicySnapShotResponseOutput) ToManagementPolicySnapShotResponsePtrOutput() ManagementPolicySnapShotResponsePtrOutput {
	return o.ToManagementPolicySnapShotResponsePtrOutputWithContext(context.Background())
}

func (o ManagementPolicySnapShotResponseOutput) ToManagementPolicySnapShotResponsePtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicySnapShotResponse) *ManagementPolicySnapShotResponse {
		return &v
	}).(ManagementPolicySnapShotResponsePtrOutput)
}

func (o ManagementPolicySnapShotResponseOutput) Delete() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicySnapShotResponse) *DateAfterCreationResponse { return v.Delete }).(DateAfterCreationResponsePtrOutput)
}

func (o ManagementPolicySnapShotResponseOutput) TierToArchive() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicySnapShotResponse) *DateAfterCreationResponse { return v.TierToArchive }).(DateAfterCreationResponsePtrOutput)
}

func (o ManagementPolicySnapShotResponseOutput) TierToCool() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicySnapShotResponse) *DateAfterCreationResponse { return v.TierToCool }).(DateAfterCreationResponsePtrOutput)
}

type ManagementPolicySnapShotResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicySnapShotResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicySnapShotResponse)(nil)).Elem()
}

func (o ManagementPolicySnapShotResponsePtrOutput) ToManagementPolicySnapShotResponsePtrOutput() ManagementPolicySnapShotResponsePtrOutput {
	return o
}

func (o ManagementPolicySnapShotResponsePtrOutput) ToManagementPolicySnapShotResponsePtrOutputWithContext(ctx context.Context) ManagementPolicySnapShotResponsePtrOutput {
	return o
}

func (o ManagementPolicySnapShotResponsePtrOutput) Elem() ManagementPolicySnapShotResponseOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShotResponse) ManagementPolicySnapShotResponse {
		if v != nil {
			return *v
		}
		var ret ManagementPolicySnapShotResponse
		return ret
	}).(ManagementPolicySnapShotResponseOutput)
}

func (o ManagementPolicySnapShotResponsePtrOutput) Delete() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShotResponse) *DateAfterCreationResponse {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(DateAfterCreationResponsePtrOutput)
}

func (o ManagementPolicySnapShotResponsePtrOutput) TierToArchive() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShotResponse) *DateAfterCreationResponse {
		if v == nil {
			return nil
		}
		return v.TierToArchive
	}).(DateAfterCreationResponsePtrOutput)
}

func (o ManagementPolicySnapShotResponsePtrOutput) TierToCool() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicySnapShotResponse) *DateAfterCreationResponse {
		if v == nil {
			return nil
		}
		return v.TierToCool
	}).(DateAfterCreationResponsePtrOutput)
}

type ManagementPolicyVersion struct {
	Delete        *DateAfterCreation `pulumi:"delete"`
	TierToArchive *DateAfterCreation `pulumi:"tierToArchive"`
	TierToCool    *DateAfterCreation `pulumi:"tierToCool"`
}





type ManagementPolicyVersionInput interface {
	pulumi.Input

	ToManagementPolicyVersionOutput() ManagementPolicyVersionOutput
	ToManagementPolicyVersionOutputWithContext(context.Context) ManagementPolicyVersionOutput
}

type ManagementPolicyVersionArgs struct {
	Delete        DateAfterCreationPtrInput `pulumi:"delete"`
	TierToArchive DateAfterCreationPtrInput `pulumi:"tierToArchive"`
	TierToCool    DateAfterCreationPtrInput `pulumi:"tierToCool"`
}

func (ManagementPolicyVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyVersion)(nil)).Elem()
}

func (i ManagementPolicyVersionArgs) ToManagementPolicyVersionOutput() ManagementPolicyVersionOutput {
	return i.ToManagementPolicyVersionOutputWithContext(context.Background())
}

func (i ManagementPolicyVersionArgs) ToManagementPolicyVersionOutputWithContext(ctx context.Context) ManagementPolicyVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyVersionOutput)
}

func (i ManagementPolicyVersionArgs) ToManagementPolicyVersionPtrOutput() ManagementPolicyVersionPtrOutput {
	return i.ToManagementPolicyVersionPtrOutputWithContext(context.Background())
}

func (i ManagementPolicyVersionArgs) ToManagementPolicyVersionPtrOutputWithContext(ctx context.Context) ManagementPolicyVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyVersionOutput).ToManagementPolicyVersionPtrOutputWithContext(ctx)
}









type ManagementPolicyVersionPtrInput interface {
	pulumi.Input

	ToManagementPolicyVersionPtrOutput() ManagementPolicyVersionPtrOutput
	ToManagementPolicyVersionPtrOutputWithContext(context.Context) ManagementPolicyVersionPtrOutput
}

type managementPolicyVersionPtrType ManagementPolicyVersionArgs

func ManagementPolicyVersionPtr(v *ManagementPolicyVersionArgs) ManagementPolicyVersionPtrInput {
	return (*managementPolicyVersionPtrType)(v)
}

func (*managementPolicyVersionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyVersion)(nil)).Elem()
}

func (i *managementPolicyVersionPtrType) ToManagementPolicyVersionPtrOutput() ManagementPolicyVersionPtrOutput {
	return i.ToManagementPolicyVersionPtrOutputWithContext(context.Background())
}

func (i *managementPolicyVersionPtrType) ToManagementPolicyVersionPtrOutputWithContext(ctx context.Context) ManagementPolicyVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyVersionPtrOutput)
}

type ManagementPolicyVersionOutput struct{ *pulumi.OutputState }

func (ManagementPolicyVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyVersion)(nil)).Elem()
}

func (o ManagementPolicyVersionOutput) ToManagementPolicyVersionOutput() ManagementPolicyVersionOutput {
	return o
}

func (o ManagementPolicyVersionOutput) ToManagementPolicyVersionOutputWithContext(ctx context.Context) ManagementPolicyVersionOutput {
	return o
}

func (o ManagementPolicyVersionOutput) ToManagementPolicyVersionPtrOutput() ManagementPolicyVersionPtrOutput {
	return o.ToManagementPolicyVersionPtrOutputWithContext(context.Background())
}

func (o ManagementPolicyVersionOutput) ToManagementPolicyVersionPtrOutputWithContext(ctx context.Context) ManagementPolicyVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicyVersion) *ManagementPolicyVersion {
		return &v
	}).(ManagementPolicyVersionPtrOutput)
}

func (o ManagementPolicyVersionOutput) Delete() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v ManagementPolicyVersion) *DateAfterCreation { return v.Delete }).(DateAfterCreationPtrOutput)
}

func (o ManagementPolicyVersionOutput) TierToArchive() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v ManagementPolicyVersion) *DateAfterCreation { return v.TierToArchive }).(DateAfterCreationPtrOutput)
}

func (o ManagementPolicyVersionOutput) TierToCool() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v ManagementPolicyVersion) *DateAfterCreation { return v.TierToCool }).(DateAfterCreationPtrOutput)
}

type ManagementPolicyVersionPtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicyVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyVersion)(nil)).Elem()
}

func (o ManagementPolicyVersionPtrOutput) ToManagementPolicyVersionPtrOutput() ManagementPolicyVersionPtrOutput {
	return o
}

func (o ManagementPolicyVersionPtrOutput) ToManagementPolicyVersionPtrOutputWithContext(ctx context.Context) ManagementPolicyVersionPtrOutput {
	return o
}

func (o ManagementPolicyVersionPtrOutput) Elem() ManagementPolicyVersionOutput {
	return o.ApplyT(func(v *ManagementPolicyVersion) ManagementPolicyVersion {
		if v != nil {
			return *v
		}
		var ret ManagementPolicyVersion
		return ret
	}).(ManagementPolicyVersionOutput)
}

func (o ManagementPolicyVersionPtrOutput) Delete() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicyVersion) *DateAfterCreation {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(DateAfterCreationPtrOutput)
}

func (o ManagementPolicyVersionPtrOutput) TierToArchive() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicyVersion) *DateAfterCreation {
		if v == nil {
			return nil
		}
		return v.TierToArchive
	}).(DateAfterCreationPtrOutput)
}

func (o ManagementPolicyVersionPtrOutput) TierToCool() DateAfterCreationPtrOutput {
	return o.ApplyT(func(v *ManagementPolicyVersion) *DateAfterCreation {
		if v == nil {
			return nil
		}
		return v.TierToCool
	}).(DateAfterCreationPtrOutput)
}

type ManagementPolicyVersionResponse struct {
	Delete        *DateAfterCreationResponse `pulumi:"delete"`
	TierToArchive *DateAfterCreationResponse `pulumi:"tierToArchive"`
	TierToCool    *DateAfterCreationResponse `pulumi:"tierToCool"`
}





type ManagementPolicyVersionResponseInput interface {
	pulumi.Input

	ToManagementPolicyVersionResponseOutput() ManagementPolicyVersionResponseOutput
	ToManagementPolicyVersionResponseOutputWithContext(context.Context) ManagementPolicyVersionResponseOutput
}

type ManagementPolicyVersionResponseArgs struct {
	Delete        DateAfterCreationResponsePtrInput `pulumi:"delete"`
	TierToArchive DateAfterCreationResponsePtrInput `pulumi:"tierToArchive"`
	TierToCool    DateAfterCreationResponsePtrInput `pulumi:"tierToCool"`
}

func (ManagementPolicyVersionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyVersionResponse)(nil)).Elem()
}

func (i ManagementPolicyVersionResponseArgs) ToManagementPolicyVersionResponseOutput() ManagementPolicyVersionResponseOutput {
	return i.ToManagementPolicyVersionResponseOutputWithContext(context.Background())
}

func (i ManagementPolicyVersionResponseArgs) ToManagementPolicyVersionResponseOutputWithContext(ctx context.Context) ManagementPolicyVersionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyVersionResponseOutput)
}

func (i ManagementPolicyVersionResponseArgs) ToManagementPolicyVersionResponsePtrOutput() ManagementPolicyVersionResponsePtrOutput {
	return i.ToManagementPolicyVersionResponsePtrOutputWithContext(context.Background())
}

func (i ManagementPolicyVersionResponseArgs) ToManagementPolicyVersionResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyVersionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyVersionResponseOutput).ToManagementPolicyVersionResponsePtrOutputWithContext(ctx)
}









type ManagementPolicyVersionResponsePtrInput interface {
	pulumi.Input

	ToManagementPolicyVersionResponsePtrOutput() ManagementPolicyVersionResponsePtrOutput
	ToManagementPolicyVersionResponsePtrOutputWithContext(context.Context) ManagementPolicyVersionResponsePtrOutput
}

type managementPolicyVersionResponsePtrType ManagementPolicyVersionResponseArgs

func ManagementPolicyVersionResponsePtr(v *ManagementPolicyVersionResponseArgs) ManagementPolicyVersionResponsePtrInput {
	return (*managementPolicyVersionResponsePtrType)(v)
}

func (*managementPolicyVersionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyVersionResponse)(nil)).Elem()
}

func (i *managementPolicyVersionResponsePtrType) ToManagementPolicyVersionResponsePtrOutput() ManagementPolicyVersionResponsePtrOutput {
	return i.ToManagementPolicyVersionResponsePtrOutputWithContext(context.Background())
}

func (i *managementPolicyVersionResponsePtrType) ToManagementPolicyVersionResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyVersionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementPolicyVersionResponsePtrOutput)
}

type ManagementPolicyVersionResponseOutput struct{ *pulumi.OutputState }

func (ManagementPolicyVersionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementPolicyVersionResponse)(nil)).Elem()
}

func (o ManagementPolicyVersionResponseOutput) ToManagementPolicyVersionResponseOutput() ManagementPolicyVersionResponseOutput {
	return o
}

func (o ManagementPolicyVersionResponseOutput) ToManagementPolicyVersionResponseOutputWithContext(ctx context.Context) ManagementPolicyVersionResponseOutput {
	return o
}

func (o ManagementPolicyVersionResponseOutput) ToManagementPolicyVersionResponsePtrOutput() ManagementPolicyVersionResponsePtrOutput {
	return o.ToManagementPolicyVersionResponsePtrOutputWithContext(context.Background())
}

func (o ManagementPolicyVersionResponseOutput) ToManagementPolicyVersionResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyVersionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementPolicyVersionResponse) *ManagementPolicyVersionResponse {
		return &v
	}).(ManagementPolicyVersionResponsePtrOutput)
}

func (o ManagementPolicyVersionResponseOutput) Delete() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyVersionResponse) *DateAfterCreationResponse { return v.Delete }).(DateAfterCreationResponsePtrOutput)
}

func (o ManagementPolicyVersionResponseOutput) TierToArchive() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyVersionResponse) *DateAfterCreationResponse { return v.TierToArchive }).(DateAfterCreationResponsePtrOutput)
}

func (o ManagementPolicyVersionResponseOutput) TierToCool() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v ManagementPolicyVersionResponse) *DateAfterCreationResponse { return v.TierToCool }).(DateAfterCreationResponsePtrOutput)
}

type ManagementPolicyVersionResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementPolicyVersionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementPolicyVersionResponse)(nil)).Elem()
}

func (o ManagementPolicyVersionResponsePtrOutput) ToManagementPolicyVersionResponsePtrOutput() ManagementPolicyVersionResponsePtrOutput {
	return o
}

func (o ManagementPolicyVersionResponsePtrOutput) ToManagementPolicyVersionResponsePtrOutputWithContext(ctx context.Context) ManagementPolicyVersionResponsePtrOutput {
	return o
}

func (o ManagementPolicyVersionResponsePtrOutput) Elem() ManagementPolicyVersionResponseOutput {
	return o.ApplyT(func(v *ManagementPolicyVersionResponse) ManagementPolicyVersionResponse {
		if v != nil {
			return *v
		}
		var ret ManagementPolicyVersionResponse
		return ret
	}).(ManagementPolicyVersionResponseOutput)
}

func (o ManagementPolicyVersionResponsePtrOutput) Delete() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicyVersionResponse) *DateAfterCreationResponse {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(DateAfterCreationResponsePtrOutput)
}

func (o ManagementPolicyVersionResponsePtrOutput) TierToArchive() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicyVersionResponse) *DateAfterCreationResponse {
		if v == nil {
			return nil
		}
		return v.TierToArchive
	}).(DateAfterCreationResponsePtrOutput)
}

func (o ManagementPolicyVersionResponsePtrOutput) TierToCool() DateAfterCreationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementPolicyVersionResponse) *DateAfterCreationResponse {
		if v == nil {
			return nil
		}
		return v.TierToCool
	}).(DateAfterCreationResponsePtrOutput)
}

type Multichannel struct {
	Enabled *bool `pulumi:"enabled"`
}





type MultichannelInput interface {
	pulumi.Input

	ToMultichannelOutput() MultichannelOutput
	ToMultichannelOutputWithContext(context.Context) MultichannelOutput
}

type MultichannelArgs struct {
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (MultichannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Multichannel)(nil)).Elem()
}

func (i MultichannelArgs) ToMultichannelOutput() MultichannelOutput {
	return i.ToMultichannelOutputWithContext(context.Background())
}

func (i MultichannelArgs) ToMultichannelOutputWithContext(ctx context.Context) MultichannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultichannelOutput)
}

func (i MultichannelArgs) ToMultichannelPtrOutput() MultichannelPtrOutput {
	return i.ToMultichannelPtrOutputWithContext(context.Background())
}

func (i MultichannelArgs) ToMultichannelPtrOutputWithContext(ctx context.Context) MultichannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultichannelOutput).ToMultichannelPtrOutputWithContext(ctx)
}









type MultichannelPtrInput interface {
	pulumi.Input

	ToMultichannelPtrOutput() MultichannelPtrOutput
	ToMultichannelPtrOutputWithContext(context.Context) MultichannelPtrOutput
}

type multichannelPtrType MultichannelArgs

func MultichannelPtr(v *MultichannelArgs) MultichannelPtrInput {
	return (*multichannelPtrType)(v)
}

func (*multichannelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Multichannel)(nil)).Elem()
}

func (i *multichannelPtrType) ToMultichannelPtrOutput() MultichannelPtrOutput {
	return i.ToMultichannelPtrOutputWithContext(context.Background())
}

func (i *multichannelPtrType) ToMultichannelPtrOutputWithContext(ctx context.Context) MultichannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultichannelPtrOutput)
}

type MultichannelOutput struct{ *pulumi.OutputState }

func (MultichannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Multichannel)(nil)).Elem()
}

func (o MultichannelOutput) ToMultichannelOutput() MultichannelOutput {
	return o
}

func (o MultichannelOutput) ToMultichannelOutputWithContext(ctx context.Context) MultichannelOutput {
	return o
}

func (o MultichannelOutput) ToMultichannelPtrOutput() MultichannelPtrOutput {
	return o.ToMultichannelPtrOutputWithContext(context.Background())
}

func (o MultichannelOutput) ToMultichannelPtrOutputWithContext(ctx context.Context) MultichannelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Multichannel) *Multichannel {
		return &v
	}).(MultichannelPtrOutput)
}

func (o MultichannelOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Multichannel) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type MultichannelPtrOutput struct{ *pulumi.OutputState }

func (MultichannelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Multichannel)(nil)).Elem()
}

func (o MultichannelPtrOutput) ToMultichannelPtrOutput() MultichannelPtrOutput {
	return o
}

func (o MultichannelPtrOutput) ToMultichannelPtrOutputWithContext(ctx context.Context) MultichannelPtrOutput {
	return o
}

func (o MultichannelPtrOutput) Elem() MultichannelOutput {
	return o.ApplyT(func(v *Multichannel) Multichannel {
		if v != nil {
			return *v
		}
		var ret Multichannel
		return ret
	}).(MultichannelOutput)
}

func (o MultichannelPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Multichannel) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type MultichannelResponse struct {
	Enabled *bool `pulumi:"enabled"`
}





type MultichannelResponseInput interface {
	pulumi.Input

	ToMultichannelResponseOutput() MultichannelResponseOutput
	ToMultichannelResponseOutputWithContext(context.Context) MultichannelResponseOutput
}

type MultichannelResponseArgs struct {
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (MultichannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MultichannelResponse)(nil)).Elem()
}

func (i MultichannelResponseArgs) ToMultichannelResponseOutput() MultichannelResponseOutput {
	return i.ToMultichannelResponseOutputWithContext(context.Background())
}

func (i MultichannelResponseArgs) ToMultichannelResponseOutputWithContext(ctx context.Context) MultichannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultichannelResponseOutput)
}

func (i MultichannelResponseArgs) ToMultichannelResponsePtrOutput() MultichannelResponsePtrOutput {
	return i.ToMultichannelResponsePtrOutputWithContext(context.Background())
}

func (i MultichannelResponseArgs) ToMultichannelResponsePtrOutputWithContext(ctx context.Context) MultichannelResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultichannelResponseOutput).ToMultichannelResponsePtrOutputWithContext(ctx)
}









type MultichannelResponsePtrInput interface {
	pulumi.Input

	ToMultichannelResponsePtrOutput() MultichannelResponsePtrOutput
	ToMultichannelResponsePtrOutputWithContext(context.Context) MultichannelResponsePtrOutput
}

type multichannelResponsePtrType MultichannelResponseArgs

func MultichannelResponsePtr(v *MultichannelResponseArgs) MultichannelResponsePtrInput {
	return (*multichannelResponsePtrType)(v)
}

func (*multichannelResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MultichannelResponse)(nil)).Elem()
}

func (i *multichannelResponsePtrType) ToMultichannelResponsePtrOutput() MultichannelResponsePtrOutput {
	return i.ToMultichannelResponsePtrOutputWithContext(context.Background())
}

func (i *multichannelResponsePtrType) ToMultichannelResponsePtrOutputWithContext(ctx context.Context) MultichannelResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultichannelResponsePtrOutput)
}

type MultichannelResponseOutput struct{ *pulumi.OutputState }

func (MultichannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MultichannelResponse)(nil)).Elem()
}

func (o MultichannelResponseOutput) ToMultichannelResponseOutput() MultichannelResponseOutput {
	return o
}

func (o MultichannelResponseOutput) ToMultichannelResponseOutputWithContext(ctx context.Context) MultichannelResponseOutput {
	return o
}

func (o MultichannelResponseOutput) ToMultichannelResponsePtrOutput() MultichannelResponsePtrOutput {
	return o.ToMultichannelResponsePtrOutputWithContext(context.Background())
}

func (o MultichannelResponseOutput) ToMultichannelResponsePtrOutputWithContext(ctx context.Context) MultichannelResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MultichannelResponse) *MultichannelResponse {
		return &v
	}).(MultichannelResponsePtrOutput)
}

func (o MultichannelResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MultichannelResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type MultichannelResponsePtrOutput struct{ *pulumi.OutputState }

func (MultichannelResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MultichannelResponse)(nil)).Elem()
}

func (o MultichannelResponsePtrOutput) ToMultichannelResponsePtrOutput() MultichannelResponsePtrOutput {
	return o
}

func (o MultichannelResponsePtrOutput) ToMultichannelResponsePtrOutputWithContext(ctx context.Context) MultichannelResponsePtrOutput {
	return o
}

func (o MultichannelResponsePtrOutput) Elem() MultichannelResponseOutput {
	return o.ApplyT(func(v *MultichannelResponse) MultichannelResponse {
		if v != nil {
			return *v
		}
		var ret MultichannelResponse
		return ret
	}).(MultichannelResponseOutput)
}

func (o MultichannelResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MultichannelResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type NetworkRuleSet struct {
	Bypass              *string              `pulumi:"bypass"`
	DefaultAction       DefaultAction        `pulumi:"defaultAction"`
	IpRules             []IPRule             `pulumi:"ipRules"`
	ResourceAccessRules []ResourceAccessRule `pulumi:"resourceAccessRules"`
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}


func (val *NetworkRuleSet) Defaults() *NetworkRuleSet {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Bypass) {
		bypass_ := "AzureServices"
		tmp.Bypass = &bypass_
	}
	if isZero(tmp.DefaultAction) {
		tmp.DefaultAction = DefaultAction("Allow")
	}
	return &tmp
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	Bypass              pulumi.StringPtrInput        `pulumi:"bypass"`
	DefaultAction       DefaultActionInput           `pulumi:"defaultAction"`
	IpRules             IPRuleArrayInput             `pulumi:"ipRules"`
	ResourceAccessRules ResourceAccessRuleArrayInput `pulumi:"resourceAccessRules"`
	VirtualNetworkRules VirtualNetworkRuleArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSet) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetOutput) DefaultAction() DefaultActionOutput {
	return o.ApplyT(func(v NetworkRuleSet) DefaultAction { return v.DefaultAction }).(DefaultActionOutput)
}

func (o NetworkRuleSetOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []IPRule { return v.IpRules }).(IPRuleArrayOutput)
}

func (o NetworkRuleSetOutput) ResourceAccessRules() ResourceAccessRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []ResourceAccessRule { return v.ResourceAccessRules }).(ResourceAccessRuleArrayOutput)
}

func (o NetworkRuleSetOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []VirtualNetworkRule { return v.VirtualNetworkRules }).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return v.Bypass
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPtrOutput) DefaultAction() DefaultActionPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *DefaultAction {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(DefaultActionPtrOutput)
}

func (o NetworkRuleSetPtrOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []IPRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleArrayOutput)
}

func (o NetworkRuleSetPtrOutput) ResourceAccessRules() ResourceAccessRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []ResourceAccessRule {
		if v == nil {
			return nil
		}
		return v.ResourceAccessRules
	}).(ResourceAccessRuleArrayOutput)
}

func (o NetworkRuleSetPtrOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []VirtualNetworkRule {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	Bypass              *string                      `pulumi:"bypass"`
	DefaultAction       string                       `pulumi:"defaultAction"`
	IpRules             []IPRuleResponse             `pulumi:"ipRules"`
	ResourceAccessRules []ResourceAccessRuleResponse `pulumi:"resourceAccessRules"`
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}


func (val *NetworkRuleSetResponse) Defaults() *NetworkRuleSetResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Bypass) {
		bypass_ := "AzureServices"
		tmp.Bypass = &bypass_
	}
	if isZero(tmp.DefaultAction) {
		tmp.DefaultAction = "Allow"
	}
	return &tmp
}





type NetworkRuleSetResponseInput interface {
	pulumi.Input

	ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput
	ToNetworkRuleSetResponseOutputWithContext(context.Context) NetworkRuleSetResponseOutput
}

type NetworkRuleSetResponseArgs struct {
	Bypass              pulumi.StringPtrInput                `pulumi:"bypass"`
	DefaultAction       pulumi.StringInput                   `pulumi:"defaultAction"`
	IpRules             IPRuleResponseArrayInput             `pulumi:"ipRules"`
	ResourceAccessRules ResourceAccessRuleResponseArrayInput `pulumi:"resourceAccessRules"`
	VirtualNetworkRules VirtualNetworkRuleResponseArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return i.ToNetworkRuleSetResponseOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput)
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput).ToNetworkRuleSetResponsePtrOutputWithContext(ctx)
}









type NetworkRuleSetResponsePtrInput interface {
	pulumi.Input

	ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput
	ToNetworkRuleSetResponsePtrOutputWithContext(context.Context) NetworkRuleSetResponsePtrOutput
}

type networkRuleSetResponsePtrType NetworkRuleSetResponseArgs

func NetworkRuleSetResponsePtr(v *NetworkRuleSetResponseArgs) NetworkRuleSetResponsePtrInput {
	return (*networkRuleSetResponsePtrType)(v)
}

func (*networkRuleSetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponsePtrOutput)
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSetResponse) *NetworkRuleSetResponse {
		return &v
	}).(NetworkRuleSetResponsePtrOutput)
}

func (o NetworkRuleSetResponseOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponseOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) string { return v.DefaultAction }).(pulumi.StringOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IPRuleResponse { return v.IpRules }).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponseOutput) ResourceAccessRules() ResourceAccessRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []ResourceAccessRuleResponse { return v.ResourceAccessRules }).(ResourceAccessRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponseOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

type NetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) Elem() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) NetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetResponse
		return ret
	}).(NetworkRuleSetResponseOutput)
}

func (o NetworkRuleSetResponsePtrOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Bypass
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []IPRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponsePtrOutput) ResourceAccessRules() ResourceAccessRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []ResourceAccessRuleResponse {
		if v == nil {
			return nil
		}
		return v.ResourceAccessRules
	}).(ResourceAccessRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponsePtrOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []VirtualNetworkRuleResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleResponseArrayOutput)
}

type ObjectReplicationPolicyFilter struct {
	MinCreationTime *string  `pulumi:"minCreationTime"`
	PrefixMatch     []string `pulumi:"prefixMatch"`
}





type ObjectReplicationPolicyFilterInput interface {
	pulumi.Input

	ToObjectReplicationPolicyFilterOutput() ObjectReplicationPolicyFilterOutput
	ToObjectReplicationPolicyFilterOutputWithContext(context.Context) ObjectReplicationPolicyFilterOutput
}

type ObjectReplicationPolicyFilterArgs struct {
	MinCreationTime pulumi.StringPtrInput   `pulumi:"minCreationTime"`
	PrefixMatch     pulumi.StringArrayInput `pulumi:"prefixMatch"`
}

func (ObjectReplicationPolicyFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectReplicationPolicyFilter)(nil)).Elem()
}

func (i ObjectReplicationPolicyFilterArgs) ToObjectReplicationPolicyFilterOutput() ObjectReplicationPolicyFilterOutput {
	return i.ToObjectReplicationPolicyFilterOutputWithContext(context.Background())
}

func (i ObjectReplicationPolicyFilterArgs) ToObjectReplicationPolicyFilterOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyFilterOutput)
}

func (i ObjectReplicationPolicyFilterArgs) ToObjectReplicationPolicyFilterPtrOutput() ObjectReplicationPolicyFilterPtrOutput {
	return i.ToObjectReplicationPolicyFilterPtrOutputWithContext(context.Background())
}

func (i ObjectReplicationPolicyFilterArgs) ToObjectReplicationPolicyFilterPtrOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyFilterOutput).ToObjectReplicationPolicyFilterPtrOutputWithContext(ctx)
}









type ObjectReplicationPolicyFilterPtrInput interface {
	pulumi.Input

	ToObjectReplicationPolicyFilterPtrOutput() ObjectReplicationPolicyFilterPtrOutput
	ToObjectReplicationPolicyFilterPtrOutputWithContext(context.Context) ObjectReplicationPolicyFilterPtrOutput
}

type objectReplicationPolicyFilterPtrType ObjectReplicationPolicyFilterArgs

func ObjectReplicationPolicyFilterPtr(v *ObjectReplicationPolicyFilterArgs) ObjectReplicationPolicyFilterPtrInput {
	return (*objectReplicationPolicyFilterPtrType)(v)
}

func (*objectReplicationPolicyFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectReplicationPolicyFilter)(nil)).Elem()
}

func (i *objectReplicationPolicyFilterPtrType) ToObjectReplicationPolicyFilterPtrOutput() ObjectReplicationPolicyFilterPtrOutput {
	return i.ToObjectReplicationPolicyFilterPtrOutputWithContext(context.Background())
}

func (i *objectReplicationPolicyFilterPtrType) ToObjectReplicationPolicyFilterPtrOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyFilterPtrOutput)
}

type ObjectReplicationPolicyFilterOutput struct{ *pulumi.OutputState }

func (ObjectReplicationPolicyFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectReplicationPolicyFilter)(nil)).Elem()
}

func (o ObjectReplicationPolicyFilterOutput) ToObjectReplicationPolicyFilterOutput() ObjectReplicationPolicyFilterOutput {
	return o
}

func (o ObjectReplicationPolicyFilterOutput) ToObjectReplicationPolicyFilterOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterOutput {
	return o
}

func (o ObjectReplicationPolicyFilterOutput) ToObjectReplicationPolicyFilterPtrOutput() ObjectReplicationPolicyFilterPtrOutput {
	return o.ToObjectReplicationPolicyFilterPtrOutputWithContext(context.Background())
}

func (o ObjectReplicationPolicyFilterOutput) ToObjectReplicationPolicyFilterPtrOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ObjectReplicationPolicyFilter) *ObjectReplicationPolicyFilter {
		return &v
	}).(ObjectReplicationPolicyFilterPtrOutput)
}

func (o ObjectReplicationPolicyFilterOutput) MinCreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyFilter) *string { return v.MinCreationTime }).(pulumi.StringPtrOutput)
}

func (o ObjectReplicationPolicyFilterOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyFilter) []string { return v.PrefixMatch }).(pulumi.StringArrayOutput)
}

type ObjectReplicationPolicyFilterPtrOutput struct{ *pulumi.OutputState }

func (ObjectReplicationPolicyFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectReplicationPolicyFilter)(nil)).Elem()
}

func (o ObjectReplicationPolicyFilterPtrOutput) ToObjectReplicationPolicyFilterPtrOutput() ObjectReplicationPolicyFilterPtrOutput {
	return o
}

func (o ObjectReplicationPolicyFilterPtrOutput) ToObjectReplicationPolicyFilterPtrOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterPtrOutput {
	return o
}

func (o ObjectReplicationPolicyFilterPtrOutput) Elem() ObjectReplicationPolicyFilterOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicyFilter) ObjectReplicationPolicyFilter {
		if v != nil {
			return *v
		}
		var ret ObjectReplicationPolicyFilter
		return ret
	}).(ObjectReplicationPolicyFilterOutput)
}

func (o ObjectReplicationPolicyFilterPtrOutput) MinCreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicyFilter) *string {
		if v == nil {
			return nil
		}
		return v.MinCreationTime
	}).(pulumi.StringPtrOutput)
}

func (o ObjectReplicationPolicyFilterPtrOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicyFilter) []string {
		if v == nil {
			return nil
		}
		return v.PrefixMatch
	}).(pulumi.StringArrayOutput)
}

type ObjectReplicationPolicyFilterResponse struct {
	MinCreationTime *string  `pulumi:"minCreationTime"`
	PrefixMatch     []string `pulumi:"prefixMatch"`
}





type ObjectReplicationPolicyFilterResponseInput interface {
	pulumi.Input

	ToObjectReplicationPolicyFilterResponseOutput() ObjectReplicationPolicyFilterResponseOutput
	ToObjectReplicationPolicyFilterResponseOutputWithContext(context.Context) ObjectReplicationPolicyFilterResponseOutput
}

type ObjectReplicationPolicyFilterResponseArgs struct {
	MinCreationTime pulumi.StringPtrInput   `pulumi:"minCreationTime"`
	PrefixMatch     pulumi.StringArrayInput `pulumi:"prefixMatch"`
}

func (ObjectReplicationPolicyFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectReplicationPolicyFilterResponse)(nil)).Elem()
}

func (i ObjectReplicationPolicyFilterResponseArgs) ToObjectReplicationPolicyFilterResponseOutput() ObjectReplicationPolicyFilterResponseOutput {
	return i.ToObjectReplicationPolicyFilterResponseOutputWithContext(context.Background())
}

func (i ObjectReplicationPolicyFilterResponseArgs) ToObjectReplicationPolicyFilterResponseOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyFilterResponseOutput)
}

func (i ObjectReplicationPolicyFilterResponseArgs) ToObjectReplicationPolicyFilterResponsePtrOutput() ObjectReplicationPolicyFilterResponsePtrOutput {
	return i.ToObjectReplicationPolicyFilterResponsePtrOutputWithContext(context.Background())
}

func (i ObjectReplicationPolicyFilterResponseArgs) ToObjectReplicationPolicyFilterResponsePtrOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyFilterResponseOutput).ToObjectReplicationPolicyFilterResponsePtrOutputWithContext(ctx)
}









type ObjectReplicationPolicyFilterResponsePtrInput interface {
	pulumi.Input

	ToObjectReplicationPolicyFilterResponsePtrOutput() ObjectReplicationPolicyFilterResponsePtrOutput
	ToObjectReplicationPolicyFilterResponsePtrOutputWithContext(context.Context) ObjectReplicationPolicyFilterResponsePtrOutput
}

type objectReplicationPolicyFilterResponsePtrType ObjectReplicationPolicyFilterResponseArgs

func ObjectReplicationPolicyFilterResponsePtr(v *ObjectReplicationPolicyFilterResponseArgs) ObjectReplicationPolicyFilterResponsePtrInput {
	return (*objectReplicationPolicyFilterResponsePtrType)(v)
}

func (*objectReplicationPolicyFilterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectReplicationPolicyFilterResponse)(nil)).Elem()
}

func (i *objectReplicationPolicyFilterResponsePtrType) ToObjectReplicationPolicyFilterResponsePtrOutput() ObjectReplicationPolicyFilterResponsePtrOutput {
	return i.ToObjectReplicationPolicyFilterResponsePtrOutputWithContext(context.Background())
}

func (i *objectReplicationPolicyFilterResponsePtrType) ToObjectReplicationPolicyFilterResponsePtrOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyFilterResponsePtrOutput)
}

type ObjectReplicationPolicyFilterResponseOutput struct{ *pulumi.OutputState }

func (ObjectReplicationPolicyFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectReplicationPolicyFilterResponse)(nil)).Elem()
}

func (o ObjectReplicationPolicyFilterResponseOutput) ToObjectReplicationPolicyFilterResponseOutput() ObjectReplicationPolicyFilterResponseOutput {
	return o
}

func (o ObjectReplicationPolicyFilterResponseOutput) ToObjectReplicationPolicyFilterResponseOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterResponseOutput {
	return o
}

func (o ObjectReplicationPolicyFilterResponseOutput) ToObjectReplicationPolicyFilterResponsePtrOutput() ObjectReplicationPolicyFilterResponsePtrOutput {
	return o.ToObjectReplicationPolicyFilterResponsePtrOutputWithContext(context.Background())
}

func (o ObjectReplicationPolicyFilterResponseOutput) ToObjectReplicationPolicyFilterResponsePtrOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ObjectReplicationPolicyFilterResponse) *ObjectReplicationPolicyFilterResponse {
		return &v
	}).(ObjectReplicationPolicyFilterResponsePtrOutput)
}

func (o ObjectReplicationPolicyFilterResponseOutput) MinCreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyFilterResponse) *string { return v.MinCreationTime }).(pulumi.StringPtrOutput)
}

func (o ObjectReplicationPolicyFilterResponseOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyFilterResponse) []string { return v.PrefixMatch }).(pulumi.StringArrayOutput)
}

type ObjectReplicationPolicyFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (ObjectReplicationPolicyFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectReplicationPolicyFilterResponse)(nil)).Elem()
}

func (o ObjectReplicationPolicyFilterResponsePtrOutput) ToObjectReplicationPolicyFilterResponsePtrOutput() ObjectReplicationPolicyFilterResponsePtrOutput {
	return o
}

func (o ObjectReplicationPolicyFilterResponsePtrOutput) ToObjectReplicationPolicyFilterResponsePtrOutputWithContext(ctx context.Context) ObjectReplicationPolicyFilterResponsePtrOutput {
	return o
}

func (o ObjectReplicationPolicyFilterResponsePtrOutput) Elem() ObjectReplicationPolicyFilterResponseOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicyFilterResponse) ObjectReplicationPolicyFilterResponse {
		if v != nil {
			return *v
		}
		var ret ObjectReplicationPolicyFilterResponse
		return ret
	}).(ObjectReplicationPolicyFilterResponseOutput)
}

func (o ObjectReplicationPolicyFilterResponsePtrOutput) MinCreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicyFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.MinCreationTime
	}).(pulumi.StringPtrOutput)
}

func (o ObjectReplicationPolicyFilterResponsePtrOutput) PrefixMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ObjectReplicationPolicyFilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.PrefixMatch
	}).(pulumi.StringArrayOutput)
}

type ObjectReplicationPolicyRule struct {
	DestinationContainer string                         `pulumi:"destinationContainer"`
	Filters              *ObjectReplicationPolicyFilter `pulumi:"filters"`
	RuleId               *string                        `pulumi:"ruleId"`
	SourceContainer      string                         `pulumi:"sourceContainer"`
}





type ObjectReplicationPolicyRuleInput interface {
	pulumi.Input

	ToObjectReplicationPolicyRuleOutput() ObjectReplicationPolicyRuleOutput
	ToObjectReplicationPolicyRuleOutputWithContext(context.Context) ObjectReplicationPolicyRuleOutput
}

type ObjectReplicationPolicyRuleArgs struct {
	DestinationContainer pulumi.StringInput                    `pulumi:"destinationContainer"`
	Filters              ObjectReplicationPolicyFilterPtrInput `pulumi:"filters"`
	RuleId               pulumi.StringPtrInput                 `pulumi:"ruleId"`
	SourceContainer      pulumi.StringInput                    `pulumi:"sourceContainer"`
}

func (ObjectReplicationPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectReplicationPolicyRule)(nil)).Elem()
}

func (i ObjectReplicationPolicyRuleArgs) ToObjectReplicationPolicyRuleOutput() ObjectReplicationPolicyRuleOutput {
	return i.ToObjectReplicationPolicyRuleOutputWithContext(context.Background())
}

func (i ObjectReplicationPolicyRuleArgs) ToObjectReplicationPolicyRuleOutputWithContext(ctx context.Context) ObjectReplicationPolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyRuleOutput)
}





type ObjectReplicationPolicyRuleArrayInput interface {
	pulumi.Input

	ToObjectReplicationPolicyRuleArrayOutput() ObjectReplicationPolicyRuleArrayOutput
	ToObjectReplicationPolicyRuleArrayOutputWithContext(context.Context) ObjectReplicationPolicyRuleArrayOutput
}

type ObjectReplicationPolicyRuleArray []ObjectReplicationPolicyRuleInput

func (ObjectReplicationPolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ObjectReplicationPolicyRule)(nil)).Elem()
}

func (i ObjectReplicationPolicyRuleArray) ToObjectReplicationPolicyRuleArrayOutput() ObjectReplicationPolicyRuleArrayOutput {
	return i.ToObjectReplicationPolicyRuleArrayOutputWithContext(context.Background())
}

func (i ObjectReplicationPolicyRuleArray) ToObjectReplicationPolicyRuleArrayOutputWithContext(ctx context.Context) ObjectReplicationPolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyRuleArrayOutput)
}

type ObjectReplicationPolicyRuleOutput struct{ *pulumi.OutputState }

func (ObjectReplicationPolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectReplicationPolicyRule)(nil)).Elem()
}

func (o ObjectReplicationPolicyRuleOutput) ToObjectReplicationPolicyRuleOutput() ObjectReplicationPolicyRuleOutput {
	return o
}

func (o ObjectReplicationPolicyRuleOutput) ToObjectReplicationPolicyRuleOutputWithContext(ctx context.Context) ObjectReplicationPolicyRuleOutput {
	return o
}

func (o ObjectReplicationPolicyRuleOutput) DestinationContainer() pulumi.StringOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyRule) string { return v.DestinationContainer }).(pulumi.StringOutput)
}

func (o ObjectReplicationPolicyRuleOutput) Filters() ObjectReplicationPolicyFilterPtrOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyRule) *ObjectReplicationPolicyFilter { return v.Filters }).(ObjectReplicationPolicyFilterPtrOutput)
}

func (o ObjectReplicationPolicyRuleOutput) RuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyRule) *string { return v.RuleId }).(pulumi.StringPtrOutput)
}

func (o ObjectReplicationPolicyRuleOutput) SourceContainer() pulumi.StringOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyRule) string { return v.SourceContainer }).(pulumi.StringOutput)
}

type ObjectReplicationPolicyRuleArrayOutput struct{ *pulumi.OutputState }

func (ObjectReplicationPolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ObjectReplicationPolicyRule)(nil)).Elem()
}

func (o ObjectReplicationPolicyRuleArrayOutput) ToObjectReplicationPolicyRuleArrayOutput() ObjectReplicationPolicyRuleArrayOutput {
	return o
}

func (o ObjectReplicationPolicyRuleArrayOutput) ToObjectReplicationPolicyRuleArrayOutputWithContext(ctx context.Context) ObjectReplicationPolicyRuleArrayOutput {
	return o
}

func (o ObjectReplicationPolicyRuleArrayOutput) Index(i pulumi.IntInput) ObjectReplicationPolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ObjectReplicationPolicyRule {
		return vs[0].([]ObjectReplicationPolicyRule)[vs[1].(int)]
	}).(ObjectReplicationPolicyRuleOutput)
}

type ObjectReplicationPolicyRuleResponse struct {
	DestinationContainer string                                 `pulumi:"destinationContainer"`
	Filters              *ObjectReplicationPolicyFilterResponse `pulumi:"filters"`
	RuleId               *string                                `pulumi:"ruleId"`
	SourceContainer      string                                 `pulumi:"sourceContainer"`
}





type ObjectReplicationPolicyRuleResponseInput interface {
	pulumi.Input

	ToObjectReplicationPolicyRuleResponseOutput() ObjectReplicationPolicyRuleResponseOutput
	ToObjectReplicationPolicyRuleResponseOutputWithContext(context.Context) ObjectReplicationPolicyRuleResponseOutput
}

type ObjectReplicationPolicyRuleResponseArgs struct {
	DestinationContainer pulumi.StringInput                            `pulumi:"destinationContainer"`
	Filters              ObjectReplicationPolicyFilterResponsePtrInput `pulumi:"filters"`
	RuleId               pulumi.StringPtrInput                         `pulumi:"ruleId"`
	SourceContainer      pulumi.StringInput                            `pulumi:"sourceContainer"`
}

func (ObjectReplicationPolicyRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectReplicationPolicyRuleResponse)(nil)).Elem()
}

func (i ObjectReplicationPolicyRuleResponseArgs) ToObjectReplicationPolicyRuleResponseOutput() ObjectReplicationPolicyRuleResponseOutput {
	return i.ToObjectReplicationPolicyRuleResponseOutputWithContext(context.Background())
}

func (i ObjectReplicationPolicyRuleResponseArgs) ToObjectReplicationPolicyRuleResponseOutputWithContext(ctx context.Context) ObjectReplicationPolicyRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyRuleResponseOutput)
}





type ObjectReplicationPolicyRuleResponseArrayInput interface {
	pulumi.Input

	ToObjectReplicationPolicyRuleResponseArrayOutput() ObjectReplicationPolicyRuleResponseArrayOutput
	ToObjectReplicationPolicyRuleResponseArrayOutputWithContext(context.Context) ObjectReplicationPolicyRuleResponseArrayOutput
}

type ObjectReplicationPolicyRuleResponseArray []ObjectReplicationPolicyRuleResponseInput

func (ObjectReplicationPolicyRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ObjectReplicationPolicyRuleResponse)(nil)).Elem()
}

func (i ObjectReplicationPolicyRuleResponseArray) ToObjectReplicationPolicyRuleResponseArrayOutput() ObjectReplicationPolicyRuleResponseArrayOutput {
	return i.ToObjectReplicationPolicyRuleResponseArrayOutputWithContext(context.Background())
}

func (i ObjectReplicationPolicyRuleResponseArray) ToObjectReplicationPolicyRuleResponseArrayOutputWithContext(ctx context.Context) ObjectReplicationPolicyRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectReplicationPolicyRuleResponseArrayOutput)
}

type ObjectReplicationPolicyRuleResponseOutput struct{ *pulumi.OutputState }

func (ObjectReplicationPolicyRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectReplicationPolicyRuleResponse)(nil)).Elem()
}

func (o ObjectReplicationPolicyRuleResponseOutput) ToObjectReplicationPolicyRuleResponseOutput() ObjectReplicationPolicyRuleResponseOutput {
	return o
}

func (o ObjectReplicationPolicyRuleResponseOutput) ToObjectReplicationPolicyRuleResponseOutputWithContext(ctx context.Context) ObjectReplicationPolicyRuleResponseOutput {
	return o
}

func (o ObjectReplicationPolicyRuleResponseOutput) DestinationContainer() pulumi.StringOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyRuleResponse) string { return v.DestinationContainer }).(pulumi.StringOutput)
}

func (o ObjectReplicationPolicyRuleResponseOutput) Filters() ObjectReplicationPolicyFilterResponsePtrOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyRuleResponse) *ObjectReplicationPolicyFilterResponse { return v.Filters }).(ObjectReplicationPolicyFilterResponsePtrOutput)
}

func (o ObjectReplicationPolicyRuleResponseOutput) RuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyRuleResponse) *string { return v.RuleId }).(pulumi.StringPtrOutput)
}

func (o ObjectReplicationPolicyRuleResponseOutput) SourceContainer() pulumi.StringOutput {
	return o.ApplyT(func(v ObjectReplicationPolicyRuleResponse) string { return v.SourceContainer }).(pulumi.StringOutput)
}

type ObjectReplicationPolicyRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ObjectReplicationPolicyRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ObjectReplicationPolicyRuleResponse)(nil)).Elem()
}

func (o ObjectReplicationPolicyRuleResponseArrayOutput) ToObjectReplicationPolicyRuleResponseArrayOutput() ObjectReplicationPolicyRuleResponseArrayOutput {
	return o
}

func (o ObjectReplicationPolicyRuleResponseArrayOutput) ToObjectReplicationPolicyRuleResponseArrayOutputWithContext(ctx context.Context) ObjectReplicationPolicyRuleResponseArrayOutput {
	return o
}

func (o ObjectReplicationPolicyRuleResponseArrayOutput) Index(i pulumi.IntInput) ObjectReplicationPolicyRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ObjectReplicationPolicyRuleResponse {
		return vs[0].([]ObjectReplicationPolicyRuleResponse)[vs[1].(int)]
	}).(ObjectReplicationPolicyRuleResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                             `pulumi:"id"`
	Name                              pulumi.StringInput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                             `pulumi:"provisioningState"`
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
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionRequired pulumi.StringPtrInput `pulumi:"actionRequired"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	Status         pulumi.StringPtrInput `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
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

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionRequired
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
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionRequired pulumi.StringPtrInput `pulumi:"actionRequired"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	Status         pulumi.StringPtrInput `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
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

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionRequired
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

type ProtocolSettings struct {
	Smb *SmbSetting `pulumi:"smb"`
}





type ProtocolSettingsInput interface {
	pulumi.Input

	ToProtocolSettingsOutput() ProtocolSettingsOutput
	ToProtocolSettingsOutputWithContext(context.Context) ProtocolSettingsOutput
}

type ProtocolSettingsArgs struct {
	Smb SmbSettingPtrInput `pulumi:"smb"`
}

func (ProtocolSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtocolSettings)(nil)).Elem()
}

func (i ProtocolSettingsArgs) ToProtocolSettingsOutput() ProtocolSettingsOutput {
	return i.ToProtocolSettingsOutputWithContext(context.Background())
}

func (i ProtocolSettingsArgs) ToProtocolSettingsOutputWithContext(ctx context.Context) ProtocolSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtocolSettingsOutput)
}

func (i ProtocolSettingsArgs) ToProtocolSettingsPtrOutput() ProtocolSettingsPtrOutput {
	return i.ToProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i ProtocolSettingsArgs) ToProtocolSettingsPtrOutputWithContext(ctx context.Context) ProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtocolSettingsOutput).ToProtocolSettingsPtrOutputWithContext(ctx)
}









type ProtocolSettingsPtrInput interface {
	pulumi.Input

	ToProtocolSettingsPtrOutput() ProtocolSettingsPtrOutput
	ToProtocolSettingsPtrOutputWithContext(context.Context) ProtocolSettingsPtrOutput
}

type protocolSettingsPtrType ProtocolSettingsArgs

func ProtocolSettingsPtr(v *ProtocolSettingsArgs) ProtocolSettingsPtrInput {
	return (*protocolSettingsPtrType)(v)
}

func (*protocolSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtocolSettings)(nil)).Elem()
}

func (i *protocolSettingsPtrType) ToProtocolSettingsPtrOutput() ProtocolSettingsPtrOutput {
	return i.ToProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i *protocolSettingsPtrType) ToProtocolSettingsPtrOutputWithContext(ctx context.Context) ProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtocolSettingsPtrOutput)
}

type ProtocolSettingsOutput struct{ *pulumi.OutputState }

func (ProtocolSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtocolSettings)(nil)).Elem()
}

func (o ProtocolSettingsOutput) ToProtocolSettingsOutput() ProtocolSettingsOutput {
	return o
}

func (o ProtocolSettingsOutput) ToProtocolSettingsOutputWithContext(ctx context.Context) ProtocolSettingsOutput {
	return o
}

func (o ProtocolSettingsOutput) ToProtocolSettingsPtrOutput() ProtocolSettingsPtrOutput {
	return o.ToProtocolSettingsPtrOutputWithContext(context.Background())
}

func (o ProtocolSettingsOutput) ToProtocolSettingsPtrOutputWithContext(ctx context.Context) ProtocolSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtocolSettings) *ProtocolSettings {
		return &v
	}).(ProtocolSettingsPtrOutput)
}

func (o ProtocolSettingsOutput) Smb() SmbSettingPtrOutput {
	return o.ApplyT(func(v ProtocolSettings) *SmbSetting { return v.Smb }).(SmbSettingPtrOutput)
}

type ProtocolSettingsPtrOutput struct{ *pulumi.OutputState }

func (ProtocolSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtocolSettings)(nil)).Elem()
}

func (o ProtocolSettingsPtrOutput) ToProtocolSettingsPtrOutput() ProtocolSettingsPtrOutput {
	return o
}

func (o ProtocolSettingsPtrOutput) ToProtocolSettingsPtrOutputWithContext(ctx context.Context) ProtocolSettingsPtrOutput {
	return o
}

func (o ProtocolSettingsPtrOutput) Elem() ProtocolSettingsOutput {
	return o.ApplyT(func(v *ProtocolSettings) ProtocolSettings {
		if v != nil {
			return *v
		}
		var ret ProtocolSettings
		return ret
	}).(ProtocolSettingsOutput)
}

func (o ProtocolSettingsPtrOutput) Smb() SmbSettingPtrOutput {
	return o.ApplyT(func(v *ProtocolSettings) *SmbSetting {
		if v == nil {
			return nil
		}
		return v.Smb
	}).(SmbSettingPtrOutput)
}

type ProtocolSettingsResponse struct {
	Smb *SmbSettingResponse `pulumi:"smb"`
}





type ProtocolSettingsResponseInput interface {
	pulumi.Input

	ToProtocolSettingsResponseOutput() ProtocolSettingsResponseOutput
	ToProtocolSettingsResponseOutputWithContext(context.Context) ProtocolSettingsResponseOutput
}

type ProtocolSettingsResponseArgs struct {
	Smb SmbSettingResponsePtrInput `pulumi:"smb"`
}

func (ProtocolSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtocolSettingsResponse)(nil)).Elem()
}

func (i ProtocolSettingsResponseArgs) ToProtocolSettingsResponseOutput() ProtocolSettingsResponseOutput {
	return i.ToProtocolSettingsResponseOutputWithContext(context.Background())
}

func (i ProtocolSettingsResponseArgs) ToProtocolSettingsResponseOutputWithContext(ctx context.Context) ProtocolSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtocolSettingsResponseOutput)
}

func (i ProtocolSettingsResponseArgs) ToProtocolSettingsResponsePtrOutput() ProtocolSettingsResponsePtrOutput {
	return i.ToProtocolSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ProtocolSettingsResponseArgs) ToProtocolSettingsResponsePtrOutputWithContext(ctx context.Context) ProtocolSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtocolSettingsResponseOutput).ToProtocolSettingsResponsePtrOutputWithContext(ctx)
}









type ProtocolSettingsResponsePtrInput interface {
	pulumi.Input

	ToProtocolSettingsResponsePtrOutput() ProtocolSettingsResponsePtrOutput
	ToProtocolSettingsResponsePtrOutputWithContext(context.Context) ProtocolSettingsResponsePtrOutput
}

type protocolSettingsResponsePtrType ProtocolSettingsResponseArgs

func ProtocolSettingsResponsePtr(v *ProtocolSettingsResponseArgs) ProtocolSettingsResponsePtrInput {
	return (*protocolSettingsResponsePtrType)(v)
}

func (*protocolSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtocolSettingsResponse)(nil)).Elem()
}

func (i *protocolSettingsResponsePtrType) ToProtocolSettingsResponsePtrOutput() ProtocolSettingsResponsePtrOutput {
	return i.ToProtocolSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *protocolSettingsResponsePtrType) ToProtocolSettingsResponsePtrOutputWithContext(ctx context.Context) ProtocolSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtocolSettingsResponsePtrOutput)
}

type ProtocolSettingsResponseOutput struct{ *pulumi.OutputState }

func (ProtocolSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtocolSettingsResponse)(nil)).Elem()
}

func (o ProtocolSettingsResponseOutput) ToProtocolSettingsResponseOutput() ProtocolSettingsResponseOutput {
	return o
}

func (o ProtocolSettingsResponseOutput) ToProtocolSettingsResponseOutputWithContext(ctx context.Context) ProtocolSettingsResponseOutput {
	return o
}

func (o ProtocolSettingsResponseOutput) ToProtocolSettingsResponsePtrOutput() ProtocolSettingsResponsePtrOutput {
	return o.ToProtocolSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ProtocolSettingsResponseOutput) ToProtocolSettingsResponsePtrOutputWithContext(ctx context.Context) ProtocolSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtocolSettingsResponse) *ProtocolSettingsResponse {
		return &v
	}).(ProtocolSettingsResponsePtrOutput)
}

func (o ProtocolSettingsResponseOutput) Smb() SmbSettingResponsePtrOutput {
	return o.ApplyT(func(v ProtocolSettingsResponse) *SmbSettingResponse { return v.Smb }).(SmbSettingResponsePtrOutput)
}

type ProtocolSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ProtocolSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtocolSettingsResponse)(nil)).Elem()
}

func (o ProtocolSettingsResponsePtrOutput) ToProtocolSettingsResponsePtrOutput() ProtocolSettingsResponsePtrOutput {
	return o
}

func (o ProtocolSettingsResponsePtrOutput) ToProtocolSettingsResponsePtrOutputWithContext(ctx context.Context) ProtocolSettingsResponsePtrOutput {
	return o
}

func (o ProtocolSettingsResponsePtrOutput) Elem() ProtocolSettingsResponseOutput {
	return o.ApplyT(func(v *ProtocolSettingsResponse) ProtocolSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ProtocolSettingsResponse
		return ret
	}).(ProtocolSettingsResponseOutput)
}

func (o ProtocolSettingsResponsePtrOutput) Smb() SmbSettingResponsePtrOutput {
	return o.ApplyT(func(v *ProtocolSettingsResponse) *SmbSettingResponse {
		if v == nil {
			return nil
		}
		return v.Smb
	}).(SmbSettingResponsePtrOutput)
}

type ResourceAccessRule struct {
	ResourceId *string `pulumi:"resourceId"`
	TenantId   *string `pulumi:"tenantId"`
}





type ResourceAccessRuleInput interface {
	pulumi.Input

	ToResourceAccessRuleOutput() ResourceAccessRuleOutput
	ToResourceAccessRuleOutputWithContext(context.Context) ResourceAccessRuleOutput
}

type ResourceAccessRuleArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	TenantId   pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (ResourceAccessRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceAccessRule)(nil)).Elem()
}

func (i ResourceAccessRuleArgs) ToResourceAccessRuleOutput() ResourceAccessRuleOutput {
	return i.ToResourceAccessRuleOutputWithContext(context.Background())
}

func (i ResourceAccessRuleArgs) ToResourceAccessRuleOutputWithContext(ctx context.Context) ResourceAccessRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceAccessRuleOutput)
}





type ResourceAccessRuleArrayInput interface {
	pulumi.Input

	ToResourceAccessRuleArrayOutput() ResourceAccessRuleArrayOutput
	ToResourceAccessRuleArrayOutputWithContext(context.Context) ResourceAccessRuleArrayOutput
}

type ResourceAccessRuleArray []ResourceAccessRuleInput

func (ResourceAccessRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceAccessRule)(nil)).Elem()
}

func (i ResourceAccessRuleArray) ToResourceAccessRuleArrayOutput() ResourceAccessRuleArrayOutput {
	return i.ToResourceAccessRuleArrayOutputWithContext(context.Background())
}

func (i ResourceAccessRuleArray) ToResourceAccessRuleArrayOutputWithContext(ctx context.Context) ResourceAccessRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceAccessRuleArrayOutput)
}

type ResourceAccessRuleOutput struct{ *pulumi.OutputState }

func (ResourceAccessRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceAccessRule)(nil)).Elem()
}

func (o ResourceAccessRuleOutput) ToResourceAccessRuleOutput() ResourceAccessRuleOutput {
	return o
}

func (o ResourceAccessRuleOutput) ToResourceAccessRuleOutputWithContext(ctx context.Context) ResourceAccessRuleOutput {
	return o
}

func (o ResourceAccessRuleOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceAccessRule) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o ResourceAccessRuleOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceAccessRule) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ResourceAccessRuleArrayOutput struct{ *pulumi.OutputState }

func (ResourceAccessRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceAccessRule)(nil)).Elem()
}

func (o ResourceAccessRuleArrayOutput) ToResourceAccessRuleArrayOutput() ResourceAccessRuleArrayOutput {
	return o
}

func (o ResourceAccessRuleArrayOutput) ToResourceAccessRuleArrayOutputWithContext(ctx context.Context) ResourceAccessRuleArrayOutput {
	return o
}

func (o ResourceAccessRuleArrayOutput) Index(i pulumi.IntInput) ResourceAccessRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceAccessRule {
		return vs[0].([]ResourceAccessRule)[vs[1].(int)]
	}).(ResourceAccessRuleOutput)
}

type ResourceAccessRuleResponse struct {
	ResourceId *string `pulumi:"resourceId"`
	TenantId   *string `pulumi:"tenantId"`
}





type ResourceAccessRuleResponseInput interface {
	pulumi.Input

	ToResourceAccessRuleResponseOutput() ResourceAccessRuleResponseOutput
	ToResourceAccessRuleResponseOutputWithContext(context.Context) ResourceAccessRuleResponseOutput
}

type ResourceAccessRuleResponseArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	TenantId   pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (ResourceAccessRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceAccessRuleResponse)(nil)).Elem()
}

func (i ResourceAccessRuleResponseArgs) ToResourceAccessRuleResponseOutput() ResourceAccessRuleResponseOutput {
	return i.ToResourceAccessRuleResponseOutputWithContext(context.Background())
}

func (i ResourceAccessRuleResponseArgs) ToResourceAccessRuleResponseOutputWithContext(ctx context.Context) ResourceAccessRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceAccessRuleResponseOutput)
}





type ResourceAccessRuleResponseArrayInput interface {
	pulumi.Input

	ToResourceAccessRuleResponseArrayOutput() ResourceAccessRuleResponseArrayOutput
	ToResourceAccessRuleResponseArrayOutputWithContext(context.Context) ResourceAccessRuleResponseArrayOutput
}

type ResourceAccessRuleResponseArray []ResourceAccessRuleResponseInput

func (ResourceAccessRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceAccessRuleResponse)(nil)).Elem()
}

func (i ResourceAccessRuleResponseArray) ToResourceAccessRuleResponseArrayOutput() ResourceAccessRuleResponseArrayOutput {
	return i.ToResourceAccessRuleResponseArrayOutputWithContext(context.Background())
}

func (i ResourceAccessRuleResponseArray) ToResourceAccessRuleResponseArrayOutputWithContext(ctx context.Context) ResourceAccessRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceAccessRuleResponseArrayOutput)
}

type ResourceAccessRuleResponseOutput struct{ *pulumi.OutputState }

func (ResourceAccessRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceAccessRuleResponse)(nil)).Elem()
}

func (o ResourceAccessRuleResponseOutput) ToResourceAccessRuleResponseOutput() ResourceAccessRuleResponseOutput {
	return o
}

func (o ResourceAccessRuleResponseOutput) ToResourceAccessRuleResponseOutputWithContext(ctx context.Context) ResourceAccessRuleResponseOutput {
	return o
}

func (o ResourceAccessRuleResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceAccessRuleResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o ResourceAccessRuleResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceAccessRuleResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ResourceAccessRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceAccessRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceAccessRuleResponse)(nil)).Elem()
}

func (o ResourceAccessRuleResponseArrayOutput) ToResourceAccessRuleResponseArrayOutput() ResourceAccessRuleResponseArrayOutput {
	return o
}

func (o ResourceAccessRuleResponseArrayOutput) ToResourceAccessRuleResponseArrayOutputWithContext(ctx context.Context) ResourceAccessRuleResponseArrayOutput {
	return o
}

func (o ResourceAccessRuleResponseArrayOutput) Index(i pulumi.IntInput) ResourceAccessRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceAccessRuleResponse {
		return vs[0].([]ResourceAccessRuleResponse)[vs[1].(int)]
	}).(ResourceAccessRuleResponseOutput)
}

type RestorePolicyProperties struct {
	Days    *int `pulumi:"days"`
	Enabled bool `pulumi:"enabled"`
}





type RestorePolicyPropertiesInput interface {
	pulumi.Input

	ToRestorePolicyPropertiesOutput() RestorePolicyPropertiesOutput
	ToRestorePolicyPropertiesOutputWithContext(context.Context) RestorePolicyPropertiesOutput
}

type RestorePolicyPropertiesArgs struct {
	Days    pulumi.IntPtrInput `pulumi:"days"`
	Enabled pulumi.BoolInput   `pulumi:"enabled"`
}

func (RestorePolicyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePolicyProperties)(nil)).Elem()
}

func (i RestorePolicyPropertiesArgs) ToRestorePolicyPropertiesOutput() RestorePolicyPropertiesOutput {
	return i.ToRestorePolicyPropertiesOutputWithContext(context.Background())
}

func (i RestorePolicyPropertiesArgs) ToRestorePolicyPropertiesOutputWithContext(ctx context.Context) RestorePolicyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePolicyPropertiesOutput)
}

func (i RestorePolicyPropertiesArgs) ToRestorePolicyPropertiesPtrOutput() RestorePolicyPropertiesPtrOutput {
	return i.ToRestorePolicyPropertiesPtrOutputWithContext(context.Background())
}

func (i RestorePolicyPropertiesArgs) ToRestorePolicyPropertiesPtrOutputWithContext(ctx context.Context) RestorePolicyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePolicyPropertiesOutput).ToRestorePolicyPropertiesPtrOutputWithContext(ctx)
}









type RestorePolicyPropertiesPtrInput interface {
	pulumi.Input

	ToRestorePolicyPropertiesPtrOutput() RestorePolicyPropertiesPtrOutput
	ToRestorePolicyPropertiesPtrOutputWithContext(context.Context) RestorePolicyPropertiesPtrOutput
}

type restorePolicyPropertiesPtrType RestorePolicyPropertiesArgs

func RestorePolicyPropertiesPtr(v *RestorePolicyPropertiesArgs) RestorePolicyPropertiesPtrInput {
	return (*restorePolicyPropertiesPtrType)(v)
}

func (*restorePolicyPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePolicyProperties)(nil)).Elem()
}

func (i *restorePolicyPropertiesPtrType) ToRestorePolicyPropertiesPtrOutput() RestorePolicyPropertiesPtrOutput {
	return i.ToRestorePolicyPropertiesPtrOutputWithContext(context.Background())
}

func (i *restorePolicyPropertiesPtrType) ToRestorePolicyPropertiesPtrOutputWithContext(ctx context.Context) RestorePolicyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePolicyPropertiesPtrOutput)
}

type RestorePolicyPropertiesOutput struct{ *pulumi.OutputState }

func (RestorePolicyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePolicyProperties)(nil)).Elem()
}

func (o RestorePolicyPropertiesOutput) ToRestorePolicyPropertiesOutput() RestorePolicyPropertiesOutput {
	return o
}

func (o RestorePolicyPropertiesOutput) ToRestorePolicyPropertiesOutputWithContext(ctx context.Context) RestorePolicyPropertiesOutput {
	return o
}

func (o RestorePolicyPropertiesOutput) ToRestorePolicyPropertiesPtrOutput() RestorePolicyPropertiesPtrOutput {
	return o.ToRestorePolicyPropertiesPtrOutputWithContext(context.Background())
}

func (o RestorePolicyPropertiesOutput) ToRestorePolicyPropertiesPtrOutputWithContext(ctx context.Context) RestorePolicyPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestorePolicyProperties) *RestorePolicyProperties {
		return &v
	}).(RestorePolicyPropertiesPtrOutput)
}

func (o RestorePolicyPropertiesOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RestorePolicyProperties) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o RestorePolicyPropertiesOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RestorePolicyProperties) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type RestorePolicyPropertiesPtrOutput struct{ *pulumi.OutputState }

func (RestorePolicyPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePolicyProperties)(nil)).Elem()
}

func (o RestorePolicyPropertiesPtrOutput) ToRestorePolicyPropertiesPtrOutput() RestorePolicyPropertiesPtrOutput {
	return o
}

func (o RestorePolicyPropertiesPtrOutput) ToRestorePolicyPropertiesPtrOutputWithContext(ctx context.Context) RestorePolicyPropertiesPtrOutput {
	return o
}

func (o RestorePolicyPropertiesPtrOutput) Elem() RestorePolicyPropertiesOutput {
	return o.ApplyT(func(v *RestorePolicyProperties) RestorePolicyProperties {
		if v != nil {
			return *v
		}
		var ret RestorePolicyProperties
		return ret
	}).(RestorePolicyPropertiesOutput)
}

func (o RestorePolicyPropertiesPtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RestorePolicyProperties) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RestorePolicyPropertiesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RestorePolicyProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type RestorePolicyPropertiesResponse struct {
	Days            *int   `pulumi:"days"`
	Enabled         bool   `pulumi:"enabled"`
	LastEnabledTime string `pulumi:"lastEnabledTime"`
	MinRestoreTime  string `pulumi:"minRestoreTime"`
}





type RestorePolicyPropertiesResponseInput interface {
	pulumi.Input

	ToRestorePolicyPropertiesResponseOutput() RestorePolicyPropertiesResponseOutput
	ToRestorePolicyPropertiesResponseOutputWithContext(context.Context) RestorePolicyPropertiesResponseOutput
}

type RestorePolicyPropertiesResponseArgs struct {
	Days            pulumi.IntPtrInput `pulumi:"days"`
	Enabled         pulumi.BoolInput   `pulumi:"enabled"`
	LastEnabledTime pulumi.StringInput `pulumi:"lastEnabledTime"`
	MinRestoreTime  pulumi.StringInput `pulumi:"minRestoreTime"`
}

func (RestorePolicyPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePolicyPropertiesResponse)(nil)).Elem()
}

func (i RestorePolicyPropertiesResponseArgs) ToRestorePolicyPropertiesResponseOutput() RestorePolicyPropertiesResponseOutput {
	return i.ToRestorePolicyPropertiesResponseOutputWithContext(context.Background())
}

func (i RestorePolicyPropertiesResponseArgs) ToRestorePolicyPropertiesResponseOutputWithContext(ctx context.Context) RestorePolicyPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePolicyPropertiesResponseOutput)
}

func (i RestorePolicyPropertiesResponseArgs) ToRestorePolicyPropertiesResponsePtrOutput() RestorePolicyPropertiesResponsePtrOutput {
	return i.ToRestorePolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i RestorePolicyPropertiesResponseArgs) ToRestorePolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) RestorePolicyPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePolicyPropertiesResponseOutput).ToRestorePolicyPropertiesResponsePtrOutputWithContext(ctx)
}









type RestorePolicyPropertiesResponsePtrInput interface {
	pulumi.Input

	ToRestorePolicyPropertiesResponsePtrOutput() RestorePolicyPropertiesResponsePtrOutput
	ToRestorePolicyPropertiesResponsePtrOutputWithContext(context.Context) RestorePolicyPropertiesResponsePtrOutput
}

type restorePolicyPropertiesResponsePtrType RestorePolicyPropertiesResponseArgs

func RestorePolicyPropertiesResponsePtr(v *RestorePolicyPropertiesResponseArgs) RestorePolicyPropertiesResponsePtrInput {
	return (*restorePolicyPropertiesResponsePtrType)(v)
}

func (*restorePolicyPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePolicyPropertiesResponse)(nil)).Elem()
}

func (i *restorePolicyPropertiesResponsePtrType) ToRestorePolicyPropertiesResponsePtrOutput() RestorePolicyPropertiesResponsePtrOutput {
	return i.ToRestorePolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *restorePolicyPropertiesResponsePtrType) ToRestorePolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) RestorePolicyPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePolicyPropertiesResponsePtrOutput)
}

type RestorePolicyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RestorePolicyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePolicyPropertiesResponse)(nil)).Elem()
}

func (o RestorePolicyPropertiesResponseOutput) ToRestorePolicyPropertiesResponseOutput() RestorePolicyPropertiesResponseOutput {
	return o
}

func (o RestorePolicyPropertiesResponseOutput) ToRestorePolicyPropertiesResponseOutputWithContext(ctx context.Context) RestorePolicyPropertiesResponseOutput {
	return o
}

func (o RestorePolicyPropertiesResponseOutput) ToRestorePolicyPropertiesResponsePtrOutput() RestorePolicyPropertiesResponsePtrOutput {
	return o.ToRestorePolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o RestorePolicyPropertiesResponseOutput) ToRestorePolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) RestorePolicyPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestorePolicyPropertiesResponse) *RestorePolicyPropertiesResponse {
		return &v
	}).(RestorePolicyPropertiesResponsePtrOutput)
}

func (o RestorePolicyPropertiesResponseOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RestorePolicyPropertiesResponse) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o RestorePolicyPropertiesResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RestorePolicyPropertiesResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o RestorePolicyPropertiesResponseOutput) LastEnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v RestorePolicyPropertiesResponse) string { return v.LastEnabledTime }).(pulumi.StringOutput)
}

func (o RestorePolicyPropertiesResponseOutput) MinRestoreTime() pulumi.StringOutput {
	return o.ApplyT(func(v RestorePolicyPropertiesResponse) string { return v.MinRestoreTime }).(pulumi.StringOutput)
}

type RestorePolicyPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (RestorePolicyPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePolicyPropertiesResponse)(nil)).Elem()
}

func (o RestorePolicyPropertiesResponsePtrOutput) ToRestorePolicyPropertiesResponsePtrOutput() RestorePolicyPropertiesResponsePtrOutput {
	return o
}

func (o RestorePolicyPropertiesResponsePtrOutput) ToRestorePolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) RestorePolicyPropertiesResponsePtrOutput {
	return o
}

func (o RestorePolicyPropertiesResponsePtrOutput) Elem() RestorePolicyPropertiesResponseOutput {
	return o.ApplyT(func(v *RestorePolicyPropertiesResponse) RestorePolicyPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret RestorePolicyPropertiesResponse
		return ret
	}).(RestorePolicyPropertiesResponseOutput)
}

func (o RestorePolicyPropertiesResponsePtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RestorePolicyPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RestorePolicyPropertiesResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RestorePolicyPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o RestorePolicyPropertiesResponsePtrOutput) LastEnabledTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestorePolicyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastEnabledTime
	}).(pulumi.StringPtrOutput)
}

func (o RestorePolicyPropertiesResponsePtrOutput) MinRestoreTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestorePolicyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MinRestoreTime
	}).(pulumi.StringPtrOutput)
}

type RoutingPreference struct {
	PublishInternetEndpoints  *bool   `pulumi:"publishInternetEndpoints"`
	PublishMicrosoftEndpoints *bool   `pulumi:"publishMicrosoftEndpoints"`
	RoutingChoice             *string `pulumi:"routingChoice"`
}





type RoutingPreferenceInput interface {
	pulumi.Input

	ToRoutingPreferenceOutput() RoutingPreferenceOutput
	ToRoutingPreferenceOutputWithContext(context.Context) RoutingPreferenceOutput
}

type RoutingPreferenceArgs struct {
	PublishInternetEndpoints  pulumi.BoolPtrInput   `pulumi:"publishInternetEndpoints"`
	PublishMicrosoftEndpoints pulumi.BoolPtrInput   `pulumi:"publishMicrosoftEndpoints"`
	RoutingChoice             pulumi.StringPtrInput `pulumi:"routingChoice"`
}

func (RoutingPreferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingPreference)(nil)).Elem()
}

func (i RoutingPreferenceArgs) ToRoutingPreferenceOutput() RoutingPreferenceOutput {
	return i.ToRoutingPreferenceOutputWithContext(context.Background())
}

func (i RoutingPreferenceArgs) ToRoutingPreferenceOutputWithContext(ctx context.Context) RoutingPreferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPreferenceOutput)
}

func (i RoutingPreferenceArgs) ToRoutingPreferencePtrOutput() RoutingPreferencePtrOutput {
	return i.ToRoutingPreferencePtrOutputWithContext(context.Background())
}

func (i RoutingPreferenceArgs) ToRoutingPreferencePtrOutputWithContext(ctx context.Context) RoutingPreferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPreferenceOutput).ToRoutingPreferencePtrOutputWithContext(ctx)
}









type RoutingPreferencePtrInput interface {
	pulumi.Input

	ToRoutingPreferencePtrOutput() RoutingPreferencePtrOutput
	ToRoutingPreferencePtrOutputWithContext(context.Context) RoutingPreferencePtrOutput
}

type routingPreferencePtrType RoutingPreferenceArgs

func RoutingPreferencePtr(v *RoutingPreferenceArgs) RoutingPreferencePtrInput {
	return (*routingPreferencePtrType)(v)
}

func (*routingPreferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingPreference)(nil)).Elem()
}

func (i *routingPreferencePtrType) ToRoutingPreferencePtrOutput() RoutingPreferencePtrOutput {
	return i.ToRoutingPreferencePtrOutputWithContext(context.Background())
}

func (i *routingPreferencePtrType) ToRoutingPreferencePtrOutputWithContext(ctx context.Context) RoutingPreferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPreferencePtrOutput)
}

type RoutingPreferenceOutput struct{ *pulumi.OutputState }

func (RoutingPreferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingPreference)(nil)).Elem()
}

func (o RoutingPreferenceOutput) ToRoutingPreferenceOutput() RoutingPreferenceOutput {
	return o
}

func (o RoutingPreferenceOutput) ToRoutingPreferenceOutputWithContext(ctx context.Context) RoutingPreferenceOutput {
	return o
}

func (o RoutingPreferenceOutput) ToRoutingPreferencePtrOutput() RoutingPreferencePtrOutput {
	return o.ToRoutingPreferencePtrOutputWithContext(context.Background())
}

func (o RoutingPreferenceOutput) ToRoutingPreferencePtrOutputWithContext(ctx context.Context) RoutingPreferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingPreference) *RoutingPreference {
		return &v
	}).(RoutingPreferencePtrOutput)
}

func (o RoutingPreferenceOutput) PublishInternetEndpoints() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RoutingPreference) *bool { return v.PublishInternetEndpoints }).(pulumi.BoolPtrOutput)
}

func (o RoutingPreferenceOutput) PublishMicrosoftEndpoints() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RoutingPreference) *bool { return v.PublishMicrosoftEndpoints }).(pulumi.BoolPtrOutput)
}

func (o RoutingPreferenceOutput) RoutingChoice() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingPreference) *string { return v.RoutingChoice }).(pulumi.StringPtrOutput)
}

type RoutingPreferencePtrOutput struct{ *pulumi.OutputState }

func (RoutingPreferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingPreference)(nil)).Elem()
}

func (o RoutingPreferencePtrOutput) ToRoutingPreferencePtrOutput() RoutingPreferencePtrOutput {
	return o
}

func (o RoutingPreferencePtrOutput) ToRoutingPreferencePtrOutputWithContext(ctx context.Context) RoutingPreferencePtrOutput {
	return o
}

func (o RoutingPreferencePtrOutput) Elem() RoutingPreferenceOutput {
	return o.ApplyT(func(v *RoutingPreference) RoutingPreference {
		if v != nil {
			return *v
		}
		var ret RoutingPreference
		return ret
	}).(RoutingPreferenceOutput)
}

func (o RoutingPreferencePtrOutput) PublishInternetEndpoints() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RoutingPreference) *bool {
		if v == nil {
			return nil
		}
		return v.PublishInternetEndpoints
	}).(pulumi.BoolPtrOutput)
}

func (o RoutingPreferencePtrOutput) PublishMicrosoftEndpoints() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RoutingPreference) *bool {
		if v == nil {
			return nil
		}
		return v.PublishMicrosoftEndpoints
	}).(pulumi.BoolPtrOutput)
}

func (o RoutingPreferencePtrOutput) RoutingChoice() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoutingPreference) *string {
		if v == nil {
			return nil
		}
		return v.RoutingChoice
	}).(pulumi.StringPtrOutput)
}

type RoutingPreferenceResponse struct {
	PublishInternetEndpoints  *bool   `pulumi:"publishInternetEndpoints"`
	PublishMicrosoftEndpoints *bool   `pulumi:"publishMicrosoftEndpoints"`
	RoutingChoice             *string `pulumi:"routingChoice"`
}





type RoutingPreferenceResponseInput interface {
	pulumi.Input

	ToRoutingPreferenceResponseOutput() RoutingPreferenceResponseOutput
	ToRoutingPreferenceResponseOutputWithContext(context.Context) RoutingPreferenceResponseOutput
}

type RoutingPreferenceResponseArgs struct {
	PublishInternetEndpoints  pulumi.BoolPtrInput   `pulumi:"publishInternetEndpoints"`
	PublishMicrosoftEndpoints pulumi.BoolPtrInput   `pulumi:"publishMicrosoftEndpoints"`
	RoutingChoice             pulumi.StringPtrInput `pulumi:"routingChoice"`
}

func (RoutingPreferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingPreferenceResponse)(nil)).Elem()
}

func (i RoutingPreferenceResponseArgs) ToRoutingPreferenceResponseOutput() RoutingPreferenceResponseOutput {
	return i.ToRoutingPreferenceResponseOutputWithContext(context.Background())
}

func (i RoutingPreferenceResponseArgs) ToRoutingPreferenceResponseOutputWithContext(ctx context.Context) RoutingPreferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPreferenceResponseOutput)
}

func (i RoutingPreferenceResponseArgs) ToRoutingPreferenceResponsePtrOutput() RoutingPreferenceResponsePtrOutput {
	return i.ToRoutingPreferenceResponsePtrOutputWithContext(context.Background())
}

func (i RoutingPreferenceResponseArgs) ToRoutingPreferenceResponsePtrOutputWithContext(ctx context.Context) RoutingPreferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPreferenceResponseOutput).ToRoutingPreferenceResponsePtrOutputWithContext(ctx)
}









type RoutingPreferenceResponsePtrInput interface {
	pulumi.Input

	ToRoutingPreferenceResponsePtrOutput() RoutingPreferenceResponsePtrOutput
	ToRoutingPreferenceResponsePtrOutputWithContext(context.Context) RoutingPreferenceResponsePtrOutput
}

type routingPreferenceResponsePtrType RoutingPreferenceResponseArgs

func RoutingPreferenceResponsePtr(v *RoutingPreferenceResponseArgs) RoutingPreferenceResponsePtrInput {
	return (*routingPreferenceResponsePtrType)(v)
}

func (*routingPreferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingPreferenceResponse)(nil)).Elem()
}

func (i *routingPreferenceResponsePtrType) ToRoutingPreferenceResponsePtrOutput() RoutingPreferenceResponsePtrOutput {
	return i.ToRoutingPreferenceResponsePtrOutputWithContext(context.Background())
}

func (i *routingPreferenceResponsePtrType) ToRoutingPreferenceResponsePtrOutputWithContext(ctx context.Context) RoutingPreferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPreferenceResponsePtrOutput)
}

type RoutingPreferenceResponseOutput struct{ *pulumi.OutputState }

func (RoutingPreferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingPreferenceResponse)(nil)).Elem()
}

func (o RoutingPreferenceResponseOutput) ToRoutingPreferenceResponseOutput() RoutingPreferenceResponseOutput {
	return o
}

func (o RoutingPreferenceResponseOutput) ToRoutingPreferenceResponseOutputWithContext(ctx context.Context) RoutingPreferenceResponseOutput {
	return o
}

func (o RoutingPreferenceResponseOutput) ToRoutingPreferenceResponsePtrOutput() RoutingPreferenceResponsePtrOutput {
	return o.ToRoutingPreferenceResponsePtrOutputWithContext(context.Background())
}

func (o RoutingPreferenceResponseOutput) ToRoutingPreferenceResponsePtrOutputWithContext(ctx context.Context) RoutingPreferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingPreferenceResponse) *RoutingPreferenceResponse {
		return &v
	}).(RoutingPreferenceResponsePtrOutput)
}

func (o RoutingPreferenceResponseOutput) PublishInternetEndpoints() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RoutingPreferenceResponse) *bool { return v.PublishInternetEndpoints }).(pulumi.BoolPtrOutput)
}

func (o RoutingPreferenceResponseOutput) PublishMicrosoftEndpoints() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RoutingPreferenceResponse) *bool { return v.PublishMicrosoftEndpoints }).(pulumi.BoolPtrOutput)
}

func (o RoutingPreferenceResponseOutput) RoutingChoice() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingPreferenceResponse) *string { return v.RoutingChoice }).(pulumi.StringPtrOutput)
}

type RoutingPreferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (RoutingPreferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingPreferenceResponse)(nil)).Elem()
}

func (o RoutingPreferenceResponsePtrOutput) ToRoutingPreferenceResponsePtrOutput() RoutingPreferenceResponsePtrOutput {
	return o
}

func (o RoutingPreferenceResponsePtrOutput) ToRoutingPreferenceResponsePtrOutputWithContext(ctx context.Context) RoutingPreferenceResponsePtrOutput {
	return o
}

func (o RoutingPreferenceResponsePtrOutput) Elem() RoutingPreferenceResponseOutput {
	return o.ApplyT(func(v *RoutingPreferenceResponse) RoutingPreferenceResponse {
		if v != nil {
			return *v
		}
		var ret RoutingPreferenceResponse
		return ret
	}).(RoutingPreferenceResponseOutput)
}

func (o RoutingPreferenceResponsePtrOutput) PublishInternetEndpoints() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RoutingPreferenceResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PublishInternetEndpoints
	}).(pulumi.BoolPtrOutput)
}

func (o RoutingPreferenceResponsePtrOutput) PublishMicrosoftEndpoints() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RoutingPreferenceResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PublishMicrosoftEndpoints
	}).(pulumi.BoolPtrOutput)
}

func (o RoutingPreferenceResponsePtrOutput) RoutingChoice() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoutingPreferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.RoutingChoice
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
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

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
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
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
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

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
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
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type SmbSetting struct {
	Multichannel *Multichannel `pulumi:"multichannel"`
}





type SmbSettingInput interface {
	pulumi.Input

	ToSmbSettingOutput() SmbSettingOutput
	ToSmbSettingOutputWithContext(context.Context) SmbSettingOutput
}

type SmbSettingArgs struct {
	Multichannel MultichannelPtrInput `pulumi:"multichannel"`
}

func (SmbSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SmbSetting)(nil)).Elem()
}

func (i SmbSettingArgs) ToSmbSettingOutput() SmbSettingOutput {
	return i.ToSmbSettingOutputWithContext(context.Background())
}

func (i SmbSettingArgs) ToSmbSettingOutputWithContext(ctx context.Context) SmbSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmbSettingOutput)
}

func (i SmbSettingArgs) ToSmbSettingPtrOutput() SmbSettingPtrOutput {
	return i.ToSmbSettingPtrOutputWithContext(context.Background())
}

func (i SmbSettingArgs) ToSmbSettingPtrOutputWithContext(ctx context.Context) SmbSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmbSettingOutput).ToSmbSettingPtrOutputWithContext(ctx)
}









type SmbSettingPtrInput interface {
	pulumi.Input

	ToSmbSettingPtrOutput() SmbSettingPtrOutput
	ToSmbSettingPtrOutputWithContext(context.Context) SmbSettingPtrOutput
}

type smbSettingPtrType SmbSettingArgs

func SmbSettingPtr(v *SmbSettingArgs) SmbSettingPtrInput {
	return (*smbSettingPtrType)(v)
}

func (*smbSettingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SmbSetting)(nil)).Elem()
}

func (i *smbSettingPtrType) ToSmbSettingPtrOutput() SmbSettingPtrOutput {
	return i.ToSmbSettingPtrOutputWithContext(context.Background())
}

func (i *smbSettingPtrType) ToSmbSettingPtrOutputWithContext(ctx context.Context) SmbSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmbSettingPtrOutput)
}

type SmbSettingOutput struct{ *pulumi.OutputState }

func (SmbSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmbSetting)(nil)).Elem()
}

func (o SmbSettingOutput) ToSmbSettingOutput() SmbSettingOutput {
	return o
}

func (o SmbSettingOutput) ToSmbSettingOutputWithContext(ctx context.Context) SmbSettingOutput {
	return o
}

func (o SmbSettingOutput) ToSmbSettingPtrOutput() SmbSettingPtrOutput {
	return o.ToSmbSettingPtrOutputWithContext(context.Background())
}

func (o SmbSettingOutput) ToSmbSettingPtrOutputWithContext(ctx context.Context) SmbSettingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SmbSetting) *SmbSetting {
		return &v
	}).(SmbSettingPtrOutput)
}

func (o SmbSettingOutput) Multichannel() MultichannelPtrOutput {
	return o.ApplyT(func(v SmbSetting) *Multichannel { return v.Multichannel }).(MultichannelPtrOutput)
}

type SmbSettingPtrOutput struct{ *pulumi.OutputState }

func (SmbSettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmbSetting)(nil)).Elem()
}

func (o SmbSettingPtrOutput) ToSmbSettingPtrOutput() SmbSettingPtrOutput {
	return o
}

func (o SmbSettingPtrOutput) ToSmbSettingPtrOutputWithContext(ctx context.Context) SmbSettingPtrOutput {
	return o
}

func (o SmbSettingPtrOutput) Elem() SmbSettingOutput {
	return o.ApplyT(func(v *SmbSetting) SmbSetting {
		if v != nil {
			return *v
		}
		var ret SmbSetting
		return ret
	}).(SmbSettingOutput)
}

func (o SmbSettingPtrOutput) Multichannel() MultichannelPtrOutput {
	return o.ApplyT(func(v *SmbSetting) *Multichannel {
		if v == nil {
			return nil
		}
		return v.Multichannel
	}).(MultichannelPtrOutput)
}

type SmbSettingResponse struct {
	Multichannel *MultichannelResponse `pulumi:"multichannel"`
}





type SmbSettingResponseInput interface {
	pulumi.Input

	ToSmbSettingResponseOutput() SmbSettingResponseOutput
	ToSmbSettingResponseOutputWithContext(context.Context) SmbSettingResponseOutput
}

type SmbSettingResponseArgs struct {
	Multichannel MultichannelResponsePtrInput `pulumi:"multichannel"`
}

func (SmbSettingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SmbSettingResponse)(nil)).Elem()
}

func (i SmbSettingResponseArgs) ToSmbSettingResponseOutput() SmbSettingResponseOutput {
	return i.ToSmbSettingResponseOutputWithContext(context.Background())
}

func (i SmbSettingResponseArgs) ToSmbSettingResponseOutputWithContext(ctx context.Context) SmbSettingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmbSettingResponseOutput)
}

func (i SmbSettingResponseArgs) ToSmbSettingResponsePtrOutput() SmbSettingResponsePtrOutput {
	return i.ToSmbSettingResponsePtrOutputWithContext(context.Background())
}

func (i SmbSettingResponseArgs) ToSmbSettingResponsePtrOutputWithContext(ctx context.Context) SmbSettingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmbSettingResponseOutput).ToSmbSettingResponsePtrOutputWithContext(ctx)
}









type SmbSettingResponsePtrInput interface {
	pulumi.Input

	ToSmbSettingResponsePtrOutput() SmbSettingResponsePtrOutput
	ToSmbSettingResponsePtrOutputWithContext(context.Context) SmbSettingResponsePtrOutput
}

type smbSettingResponsePtrType SmbSettingResponseArgs

func SmbSettingResponsePtr(v *SmbSettingResponseArgs) SmbSettingResponsePtrInput {
	return (*smbSettingResponsePtrType)(v)
}

func (*smbSettingResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SmbSettingResponse)(nil)).Elem()
}

func (i *smbSettingResponsePtrType) ToSmbSettingResponsePtrOutput() SmbSettingResponsePtrOutput {
	return i.ToSmbSettingResponsePtrOutputWithContext(context.Background())
}

func (i *smbSettingResponsePtrType) ToSmbSettingResponsePtrOutputWithContext(ctx context.Context) SmbSettingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmbSettingResponsePtrOutput)
}

type SmbSettingResponseOutput struct{ *pulumi.OutputState }

func (SmbSettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmbSettingResponse)(nil)).Elem()
}

func (o SmbSettingResponseOutput) ToSmbSettingResponseOutput() SmbSettingResponseOutput {
	return o
}

func (o SmbSettingResponseOutput) ToSmbSettingResponseOutputWithContext(ctx context.Context) SmbSettingResponseOutput {
	return o
}

func (o SmbSettingResponseOutput) ToSmbSettingResponsePtrOutput() SmbSettingResponsePtrOutput {
	return o.ToSmbSettingResponsePtrOutputWithContext(context.Background())
}

func (o SmbSettingResponseOutput) ToSmbSettingResponsePtrOutputWithContext(ctx context.Context) SmbSettingResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SmbSettingResponse) *SmbSettingResponse {
		return &v
	}).(SmbSettingResponsePtrOutput)
}

func (o SmbSettingResponseOutput) Multichannel() MultichannelResponsePtrOutput {
	return o.ApplyT(func(v SmbSettingResponse) *MultichannelResponse { return v.Multichannel }).(MultichannelResponsePtrOutput)
}

type SmbSettingResponsePtrOutput struct{ *pulumi.OutputState }

func (SmbSettingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmbSettingResponse)(nil)).Elem()
}

func (o SmbSettingResponsePtrOutput) ToSmbSettingResponsePtrOutput() SmbSettingResponsePtrOutput {
	return o
}

func (o SmbSettingResponsePtrOutput) ToSmbSettingResponsePtrOutputWithContext(ctx context.Context) SmbSettingResponsePtrOutput {
	return o
}

func (o SmbSettingResponsePtrOutput) Elem() SmbSettingResponseOutput {
	return o.ApplyT(func(v *SmbSettingResponse) SmbSettingResponse {
		if v != nil {
			return *v
		}
		var ret SmbSettingResponse
		return ret
	}).(SmbSettingResponseOutput)
}

func (o SmbSettingResponsePtrOutput) Multichannel() MultichannelResponsePtrOutput {
	return o.ApplyT(func(v *SmbSettingResponse) *MultichannelResponse {
		if v == nil {
			return nil
		}
		return v.Multichannel
	}).(MultichannelResponsePtrOutput)
}

type StorageAccountInternetEndpointsResponse struct {
	Blob string `pulumi:"blob"`
	Dfs  string `pulumi:"dfs"`
	File string `pulumi:"file"`
	Web  string `pulumi:"web"`
}





type StorageAccountInternetEndpointsResponseInput interface {
	pulumi.Input

	ToStorageAccountInternetEndpointsResponseOutput() StorageAccountInternetEndpointsResponseOutput
	ToStorageAccountInternetEndpointsResponseOutputWithContext(context.Context) StorageAccountInternetEndpointsResponseOutput
}

type StorageAccountInternetEndpointsResponseArgs struct {
	Blob pulumi.StringInput `pulumi:"blob"`
	Dfs  pulumi.StringInput `pulumi:"dfs"`
	File pulumi.StringInput `pulumi:"file"`
	Web  pulumi.StringInput `pulumi:"web"`
}

func (StorageAccountInternetEndpointsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountInternetEndpointsResponse)(nil)).Elem()
}

func (i StorageAccountInternetEndpointsResponseArgs) ToStorageAccountInternetEndpointsResponseOutput() StorageAccountInternetEndpointsResponseOutput {
	return i.ToStorageAccountInternetEndpointsResponseOutputWithContext(context.Background())
}

func (i StorageAccountInternetEndpointsResponseArgs) ToStorageAccountInternetEndpointsResponseOutputWithContext(ctx context.Context) StorageAccountInternetEndpointsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountInternetEndpointsResponseOutput)
}

func (i StorageAccountInternetEndpointsResponseArgs) ToStorageAccountInternetEndpointsResponsePtrOutput() StorageAccountInternetEndpointsResponsePtrOutput {
	return i.ToStorageAccountInternetEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i StorageAccountInternetEndpointsResponseArgs) ToStorageAccountInternetEndpointsResponsePtrOutputWithContext(ctx context.Context) StorageAccountInternetEndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountInternetEndpointsResponseOutput).ToStorageAccountInternetEndpointsResponsePtrOutputWithContext(ctx)
}









type StorageAccountInternetEndpointsResponsePtrInput interface {
	pulumi.Input

	ToStorageAccountInternetEndpointsResponsePtrOutput() StorageAccountInternetEndpointsResponsePtrOutput
	ToStorageAccountInternetEndpointsResponsePtrOutputWithContext(context.Context) StorageAccountInternetEndpointsResponsePtrOutput
}

type storageAccountInternetEndpointsResponsePtrType StorageAccountInternetEndpointsResponseArgs

func StorageAccountInternetEndpointsResponsePtr(v *StorageAccountInternetEndpointsResponseArgs) StorageAccountInternetEndpointsResponsePtrInput {
	return (*storageAccountInternetEndpointsResponsePtrType)(v)
}

func (*storageAccountInternetEndpointsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountInternetEndpointsResponse)(nil)).Elem()
}

func (i *storageAccountInternetEndpointsResponsePtrType) ToStorageAccountInternetEndpointsResponsePtrOutput() StorageAccountInternetEndpointsResponsePtrOutput {
	return i.ToStorageAccountInternetEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i *storageAccountInternetEndpointsResponsePtrType) ToStorageAccountInternetEndpointsResponsePtrOutputWithContext(ctx context.Context) StorageAccountInternetEndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountInternetEndpointsResponsePtrOutput)
}

type StorageAccountInternetEndpointsResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountInternetEndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountInternetEndpointsResponse)(nil)).Elem()
}

func (o StorageAccountInternetEndpointsResponseOutput) ToStorageAccountInternetEndpointsResponseOutput() StorageAccountInternetEndpointsResponseOutput {
	return o
}

func (o StorageAccountInternetEndpointsResponseOutput) ToStorageAccountInternetEndpointsResponseOutputWithContext(ctx context.Context) StorageAccountInternetEndpointsResponseOutput {
	return o
}

func (o StorageAccountInternetEndpointsResponseOutput) ToStorageAccountInternetEndpointsResponsePtrOutput() StorageAccountInternetEndpointsResponsePtrOutput {
	return o.ToStorageAccountInternetEndpointsResponsePtrOutputWithContext(context.Background())
}

func (o StorageAccountInternetEndpointsResponseOutput) ToStorageAccountInternetEndpointsResponsePtrOutputWithContext(ctx context.Context) StorageAccountInternetEndpointsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountInternetEndpointsResponse) *StorageAccountInternetEndpointsResponse {
		return &v
	}).(StorageAccountInternetEndpointsResponsePtrOutput)
}

func (o StorageAccountInternetEndpointsResponseOutput) Blob() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountInternetEndpointsResponse) string { return v.Blob }).(pulumi.StringOutput)
}

func (o StorageAccountInternetEndpointsResponseOutput) Dfs() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountInternetEndpointsResponse) string { return v.Dfs }).(pulumi.StringOutput)
}

func (o StorageAccountInternetEndpointsResponseOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountInternetEndpointsResponse) string { return v.File }).(pulumi.StringOutput)
}

func (o StorageAccountInternetEndpointsResponseOutput) Web() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountInternetEndpointsResponse) string { return v.Web }).(pulumi.StringOutput)
}

type StorageAccountInternetEndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountInternetEndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountInternetEndpointsResponse)(nil)).Elem()
}

func (o StorageAccountInternetEndpointsResponsePtrOutput) ToStorageAccountInternetEndpointsResponsePtrOutput() StorageAccountInternetEndpointsResponsePtrOutput {
	return o
}

func (o StorageAccountInternetEndpointsResponsePtrOutput) ToStorageAccountInternetEndpointsResponsePtrOutputWithContext(ctx context.Context) StorageAccountInternetEndpointsResponsePtrOutput {
	return o
}

func (o StorageAccountInternetEndpointsResponsePtrOutput) Elem() StorageAccountInternetEndpointsResponseOutput {
	return o.ApplyT(func(v *StorageAccountInternetEndpointsResponse) StorageAccountInternetEndpointsResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountInternetEndpointsResponse
		return ret
	}).(StorageAccountInternetEndpointsResponseOutput)
}

func (o StorageAccountInternetEndpointsResponsePtrOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountInternetEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Blob
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountInternetEndpointsResponsePtrOutput) Dfs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountInternetEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Dfs
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountInternetEndpointsResponsePtrOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountInternetEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.File
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountInternetEndpointsResponsePtrOutput) Web() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountInternetEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Web
	}).(pulumi.StringPtrOutput)
}

type StorageAccountKeyResponse struct {
	KeyName     string `pulumi:"keyName"`
	Permissions string `pulumi:"permissions"`
	Value       string `pulumi:"value"`
}





type StorageAccountKeyResponseInput interface {
	pulumi.Input

	ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput
	ToStorageAccountKeyResponseOutputWithContext(context.Context) StorageAccountKeyResponseOutput
}

type StorageAccountKeyResponseArgs struct {
	KeyName     pulumi.StringInput `pulumi:"keyName"`
	Permissions pulumi.StringInput `pulumi:"permissions"`
	Value       pulumi.StringInput `pulumi:"value"`
}

func (StorageAccountKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountKeyResponse)(nil)).Elem()
}

func (i StorageAccountKeyResponseArgs) ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput {
	return i.ToStorageAccountKeyResponseOutputWithContext(context.Background())
}

func (i StorageAccountKeyResponseArgs) ToStorageAccountKeyResponseOutputWithContext(ctx context.Context) StorageAccountKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountKeyResponseOutput)
}





type StorageAccountKeyResponseArrayInput interface {
	pulumi.Input

	ToStorageAccountKeyResponseArrayOutput() StorageAccountKeyResponseArrayOutput
	ToStorageAccountKeyResponseArrayOutputWithContext(context.Context) StorageAccountKeyResponseArrayOutput
}

type StorageAccountKeyResponseArray []StorageAccountKeyResponseInput

func (StorageAccountKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountKeyResponse)(nil)).Elem()
}

func (i StorageAccountKeyResponseArray) ToStorageAccountKeyResponseArrayOutput() StorageAccountKeyResponseArrayOutput {
	return i.ToStorageAccountKeyResponseArrayOutputWithContext(context.Background())
}

func (i StorageAccountKeyResponseArray) ToStorageAccountKeyResponseArrayOutputWithContext(ctx context.Context) StorageAccountKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountKeyResponseArrayOutput)
}

type StorageAccountKeyResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountKeyResponse)(nil)).Elem()
}

func (o StorageAccountKeyResponseOutput) ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput {
	return o
}

func (o StorageAccountKeyResponseOutput) ToStorageAccountKeyResponseOutputWithContext(ctx context.Context) StorageAccountKeyResponseOutput {
	return o
}

func (o StorageAccountKeyResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o StorageAccountKeyResponseOutput) Permissions() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.Permissions }).(pulumi.StringOutput)
}

func (o StorageAccountKeyResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.Value }).(pulumi.StringOutput)
}

type StorageAccountKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountKeyResponse)(nil)).Elem()
}

func (o StorageAccountKeyResponseArrayOutput) ToStorageAccountKeyResponseArrayOutput() StorageAccountKeyResponseArrayOutput {
	return o
}

func (o StorageAccountKeyResponseArrayOutput) ToStorageAccountKeyResponseArrayOutputWithContext(ctx context.Context) StorageAccountKeyResponseArrayOutput {
	return o
}

func (o StorageAccountKeyResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountKeyResponse {
		return vs[0].([]StorageAccountKeyResponse)[vs[1].(int)]
	}).(StorageAccountKeyResponseOutput)
}

type StorageAccountMicrosoftEndpointsResponse struct {
	Blob  string `pulumi:"blob"`
	Dfs   string `pulumi:"dfs"`
	File  string `pulumi:"file"`
	Queue string `pulumi:"queue"`
	Table string `pulumi:"table"`
	Web   string `pulumi:"web"`
}





type StorageAccountMicrosoftEndpointsResponseInput interface {
	pulumi.Input

	ToStorageAccountMicrosoftEndpointsResponseOutput() StorageAccountMicrosoftEndpointsResponseOutput
	ToStorageAccountMicrosoftEndpointsResponseOutputWithContext(context.Context) StorageAccountMicrosoftEndpointsResponseOutput
}

type StorageAccountMicrosoftEndpointsResponseArgs struct {
	Blob  pulumi.StringInput `pulumi:"blob"`
	Dfs   pulumi.StringInput `pulumi:"dfs"`
	File  pulumi.StringInput `pulumi:"file"`
	Queue pulumi.StringInput `pulumi:"queue"`
	Table pulumi.StringInput `pulumi:"table"`
	Web   pulumi.StringInput `pulumi:"web"`
}

func (StorageAccountMicrosoftEndpointsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountMicrosoftEndpointsResponse)(nil)).Elem()
}

func (i StorageAccountMicrosoftEndpointsResponseArgs) ToStorageAccountMicrosoftEndpointsResponseOutput() StorageAccountMicrosoftEndpointsResponseOutput {
	return i.ToStorageAccountMicrosoftEndpointsResponseOutputWithContext(context.Background())
}

func (i StorageAccountMicrosoftEndpointsResponseArgs) ToStorageAccountMicrosoftEndpointsResponseOutputWithContext(ctx context.Context) StorageAccountMicrosoftEndpointsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountMicrosoftEndpointsResponseOutput)
}

func (i StorageAccountMicrosoftEndpointsResponseArgs) ToStorageAccountMicrosoftEndpointsResponsePtrOutput() StorageAccountMicrosoftEndpointsResponsePtrOutput {
	return i.ToStorageAccountMicrosoftEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i StorageAccountMicrosoftEndpointsResponseArgs) ToStorageAccountMicrosoftEndpointsResponsePtrOutputWithContext(ctx context.Context) StorageAccountMicrosoftEndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountMicrosoftEndpointsResponseOutput).ToStorageAccountMicrosoftEndpointsResponsePtrOutputWithContext(ctx)
}









type StorageAccountMicrosoftEndpointsResponsePtrInput interface {
	pulumi.Input

	ToStorageAccountMicrosoftEndpointsResponsePtrOutput() StorageAccountMicrosoftEndpointsResponsePtrOutput
	ToStorageAccountMicrosoftEndpointsResponsePtrOutputWithContext(context.Context) StorageAccountMicrosoftEndpointsResponsePtrOutput
}

type storageAccountMicrosoftEndpointsResponsePtrType StorageAccountMicrosoftEndpointsResponseArgs

func StorageAccountMicrosoftEndpointsResponsePtr(v *StorageAccountMicrosoftEndpointsResponseArgs) StorageAccountMicrosoftEndpointsResponsePtrInput {
	return (*storageAccountMicrosoftEndpointsResponsePtrType)(v)
}

func (*storageAccountMicrosoftEndpointsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountMicrosoftEndpointsResponse)(nil)).Elem()
}

func (i *storageAccountMicrosoftEndpointsResponsePtrType) ToStorageAccountMicrosoftEndpointsResponsePtrOutput() StorageAccountMicrosoftEndpointsResponsePtrOutput {
	return i.ToStorageAccountMicrosoftEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i *storageAccountMicrosoftEndpointsResponsePtrType) ToStorageAccountMicrosoftEndpointsResponsePtrOutputWithContext(ctx context.Context) StorageAccountMicrosoftEndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountMicrosoftEndpointsResponsePtrOutput)
}

type StorageAccountMicrosoftEndpointsResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountMicrosoftEndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountMicrosoftEndpointsResponse)(nil)).Elem()
}

func (o StorageAccountMicrosoftEndpointsResponseOutput) ToStorageAccountMicrosoftEndpointsResponseOutput() StorageAccountMicrosoftEndpointsResponseOutput {
	return o
}

func (o StorageAccountMicrosoftEndpointsResponseOutput) ToStorageAccountMicrosoftEndpointsResponseOutputWithContext(ctx context.Context) StorageAccountMicrosoftEndpointsResponseOutput {
	return o
}

func (o StorageAccountMicrosoftEndpointsResponseOutput) ToStorageAccountMicrosoftEndpointsResponsePtrOutput() StorageAccountMicrosoftEndpointsResponsePtrOutput {
	return o.ToStorageAccountMicrosoftEndpointsResponsePtrOutputWithContext(context.Background())
}

func (o StorageAccountMicrosoftEndpointsResponseOutput) ToStorageAccountMicrosoftEndpointsResponsePtrOutputWithContext(ctx context.Context) StorageAccountMicrosoftEndpointsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountMicrosoftEndpointsResponse) *StorageAccountMicrosoftEndpointsResponse {
		return &v
	}).(StorageAccountMicrosoftEndpointsResponsePtrOutput)
}

func (o StorageAccountMicrosoftEndpointsResponseOutput) Blob() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountMicrosoftEndpointsResponse) string { return v.Blob }).(pulumi.StringOutput)
}

func (o StorageAccountMicrosoftEndpointsResponseOutput) Dfs() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountMicrosoftEndpointsResponse) string { return v.Dfs }).(pulumi.StringOutput)
}

func (o StorageAccountMicrosoftEndpointsResponseOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountMicrosoftEndpointsResponse) string { return v.File }).(pulumi.StringOutput)
}

func (o StorageAccountMicrosoftEndpointsResponseOutput) Queue() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountMicrosoftEndpointsResponse) string { return v.Queue }).(pulumi.StringOutput)
}

func (o StorageAccountMicrosoftEndpointsResponseOutput) Table() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountMicrosoftEndpointsResponse) string { return v.Table }).(pulumi.StringOutput)
}

func (o StorageAccountMicrosoftEndpointsResponseOutput) Web() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountMicrosoftEndpointsResponse) string { return v.Web }).(pulumi.StringOutput)
}

type StorageAccountMicrosoftEndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountMicrosoftEndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountMicrosoftEndpointsResponse)(nil)).Elem()
}

func (o StorageAccountMicrosoftEndpointsResponsePtrOutput) ToStorageAccountMicrosoftEndpointsResponsePtrOutput() StorageAccountMicrosoftEndpointsResponsePtrOutput {
	return o
}

func (o StorageAccountMicrosoftEndpointsResponsePtrOutput) ToStorageAccountMicrosoftEndpointsResponsePtrOutputWithContext(ctx context.Context) StorageAccountMicrosoftEndpointsResponsePtrOutput {
	return o
}

func (o StorageAccountMicrosoftEndpointsResponsePtrOutput) Elem() StorageAccountMicrosoftEndpointsResponseOutput {
	return o.ApplyT(func(v *StorageAccountMicrosoftEndpointsResponse) StorageAccountMicrosoftEndpointsResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountMicrosoftEndpointsResponse
		return ret
	}).(StorageAccountMicrosoftEndpointsResponseOutput)
}

func (o StorageAccountMicrosoftEndpointsResponsePtrOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountMicrosoftEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Blob
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountMicrosoftEndpointsResponsePtrOutput) Dfs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountMicrosoftEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Dfs
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountMicrosoftEndpointsResponsePtrOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountMicrosoftEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.File
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountMicrosoftEndpointsResponsePtrOutput) Queue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountMicrosoftEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Queue
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountMicrosoftEndpointsResponsePtrOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountMicrosoftEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Table
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountMicrosoftEndpointsResponsePtrOutput) Web() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountMicrosoftEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Web
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type TagFilter struct {
	Name  string `pulumi:"name"`
	Op    string `pulumi:"op"`
	Value string `pulumi:"value"`
}





type TagFilterInput interface {
	pulumi.Input

	ToTagFilterOutput() TagFilterOutput
	ToTagFilterOutputWithContext(context.Context) TagFilterOutput
}

type TagFilterArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Op    pulumi.StringInput `pulumi:"op"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TagFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagFilter)(nil)).Elem()
}

func (i TagFilterArgs) ToTagFilterOutput() TagFilterOutput {
	return i.ToTagFilterOutputWithContext(context.Background())
}

func (i TagFilterArgs) ToTagFilterOutputWithContext(ctx context.Context) TagFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagFilterOutput)
}





type TagFilterArrayInput interface {
	pulumi.Input

	ToTagFilterArrayOutput() TagFilterArrayOutput
	ToTagFilterArrayOutputWithContext(context.Context) TagFilterArrayOutput
}

type TagFilterArray []TagFilterInput

func (TagFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagFilter)(nil)).Elem()
}

func (i TagFilterArray) ToTagFilterArrayOutput() TagFilterArrayOutput {
	return i.ToTagFilterArrayOutputWithContext(context.Background())
}

func (i TagFilterArray) ToTagFilterArrayOutputWithContext(ctx context.Context) TagFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagFilterArrayOutput)
}

type TagFilterOutput struct{ *pulumi.OutputState }

func (TagFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagFilter)(nil)).Elem()
}

func (o TagFilterOutput) ToTagFilterOutput() TagFilterOutput {
	return o
}

func (o TagFilterOutput) ToTagFilterOutputWithContext(ctx context.Context) TagFilterOutput {
	return o
}

func (o TagFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TagFilter) string { return v.Name }).(pulumi.StringOutput)
}

func (o TagFilterOutput) Op() pulumi.StringOutput {
	return o.ApplyT(func(v TagFilter) string { return v.Op }).(pulumi.StringOutput)
}

func (o TagFilterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TagFilter) string { return v.Value }).(pulumi.StringOutput)
}

type TagFilterArrayOutput struct{ *pulumi.OutputState }

func (TagFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagFilter)(nil)).Elem()
}

func (o TagFilterArrayOutput) ToTagFilterArrayOutput() TagFilterArrayOutput {
	return o
}

func (o TagFilterArrayOutput) ToTagFilterArrayOutputWithContext(ctx context.Context) TagFilterArrayOutput {
	return o
}

func (o TagFilterArrayOutput) Index(i pulumi.IntInput) TagFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagFilter {
		return vs[0].([]TagFilter)[vs[1].(int)]
	}).(TagFilterOutput)
}

type TagFilterResponse struct {
	Name  string `pulumi:"name"`
	Op    string `pulumi:"op"`
	Value string `pulumi:"value"`
}





type TagFilterResponseInput interface {
	pulumi.Input

	ToTagFilterResponseOutput() TagFilterResponseOutput
	ToTagFilterResponseOutputWithContext(context.Context) TagFilterResponseOutput
}

type TagFilterResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Op    pulumi.StringInput `pulumi:"op"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TagFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagFilterResponse)(nil)).Elem()
}

func (i TagFilterResponseArgs) ToTagFilterResponseOutput() TagFilterResponseOutput {
	return i.ToTagFilterResponseOutputWithContext(context.Background())
}

func (i TagFilterResponseArgs) ToTagFilterResponseOutputWithContext(ctx context.Context) TagFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagFilterResponseOutput)
}





type TagFilterResponseArrayInput interface {
	pulumi.Input

	ToTagFilterResponseArrayOutput() TagFilterResponseArrayOutput
	ToTagFilterResponseArrayOutputWithContext(context.Context) TagFilterResponseArrayOutput
}

type TagFilterResponseArray []TagFilterResponseInput

func (TagFilterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagFilterResponse)(nil)).Elem()
}

func (i TagFilterResponseArray) ToTagFilterResponseArrayOutput() TagFilterResponseArrayOutput {
	return i.ToTagFilterResponseArrayOutputWithContext(context.Background())
}

func (i TagFilterResponseArray) ToTagFilterResponseArrayOutputWithContext(ctx context.Context) TagFilterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagFilterResponseArrayOutput)
}

type TagFilterResponseOutput struct{ *pulumi.OutputState }

func (TagFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagFilterResponse)(nil)).Elem()
}

func (o TagFilterResponseOutput) ToTagFilterResponseOutput() TagFilterResponseOutput {
	return o
}

func (o TagFilterResponseOutput) ToTagFilterResponseOutputWithContext(ctx context.Context) TagFilterResponseOutput {
	return o
}

func (o TagFilterResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TagFilterResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TagFilterResponseOutput) Op() pulumi.StringOutput {
	return o.ApplyT(func(v TagFilterResponse) string { return v.Op }).(pulumi.StringOutput)
}

func (o TagFilterResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TagFilterResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TagFilterResponseArrayOutput struct{ *pulumi.OutputState }

func (TagFilterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagFilterResponse)(nil)).Elem()
}

func (o TagFilterResponseArrayOutput) ToTagFilterResponseArrayOutput() TagFilterResponseArrayOutput {
	return o
}

func (o TagFilterResponseArrayOutput) ToTagFilterResponseArrayOutputWithContext(ctx context.Context) TagFilterResponseArrayOutput {
	return o
}

func (o TagFilterResponseArrayOutput) Index(i pulumi.IntInput) TagFilterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagFilterResponse {
		return vs[0].([]TagFilterResponse)[vs[1].(int)]
	}).(TagFilterResponseOutput)
}

type TagPropertyResponse struct {
	ObjectIdentifier string `pulumi:"objectIdentifier"`
	Tag              string `pulumi:"tag"`
	TenantId         string `pulumi:"tenantId"`
	Timestamp        string `pulumi:"timestamp"`
	Upn              string `pulumi:"upn"`
}





type TagPropertyResponseInput interface {
	pulumi.Input

	ToTagPropertyResponseOutput() TagPropertyResponseOutput
	ToTagPropertyResponseOutputWithContext(context.Context) TagPropertyResponseOutput
}

type TagPropertyResponseArgs struct {
	ObjectIdentifier pulumi.StringInput `pulumi:"objectIdentifier"`
	Tag              pulumi.StringInput `pulumi:"tag"`
	TenantId         pulumi.StringInput `pulumi:"tenantId"`
	Timestamp        pulumi.StringInput `pulumi:"timestamp"`
	Upn              pulumi.StringInput `pulumi:"upn"`
}

func (TagPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagPropertyResponse)(nil)).Elem()
}

func (i TagPropertyResponseArgs) ToTagPropertyResponseOutput() TagPropertyResponseOutput {
	return i.ToTagPropertyResponseOutputWithContext(context.Background())
}

func (i TagPropertyResponseArgs) ToTagPropertyResponseOutputWithContext(ctx context.Context) TagPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagPropertyResponseOutput)
}





type TagPropertyResponseArrayInput interface {
	pulumi.Input

	ToTagPropertyResponseArrayOutput() TagPropertyResponseArrayOutput
	ToTagPropertyResponseArrayOutputWithContext(context.Context) TagPropertyResponseArrayOutput
}

type TagPropertyResponseArray []TagPropertyResponseInput

func (TagPropertyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagPropertyResponse)(nil)).Elem()
}

func (i TagPropertyResponseArray) ToTagPropertyResponseArrayOutput() TagPropertyResponseArrayOutput {
	return i.ToTagPropertyResponseArrayOutputWithContext(context.Background())
}

func (i TagPropertyResponseArray) ToTagPropertyResponseArrayOutputWithContext(ctx context.Context) TagPropertyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagPropertyResponseArrayOutput)
}

type TagPropertyResponseOutput struct{ *pulumi.OutputState }

func (TagPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagPropertyResponse)(nil)).Elem()
}

func (o TagPropertyResponseOutput) ToTagPropertyResponseOutput() TagPropertyResponseOutput {
	return o
}

func (o TagPropertyResponseOutput) ToTagPropertyResponseOutputWithContext(ctx context.Context) TagPropertyResponseOutput {
	return o
}

func (o TagPropertyResponseOutput) ObjectIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.ObjectIdentifier }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.Tag }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) Upn() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.Upn }).(pulumi.StringOutput)
}

type TagPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (TagPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagPropertyResponse)(nil)).Elem()
}

func (o TagPropertyResponseArrayOutput) ToTagPropertyResponseArrayOutput() TagPropertyResponseArrayOutput {
	return o
}

func (o TagPropertyResponseArrayOutput) ToTagPropertyResponseArrayOutputWithContext(ctx context.Context) TagPropertyResponseArrayOutput {
	return o
}

func (o TagPropertyResponseArrayOutput) Index(i pulumi.IntInput) TagPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagPropertyResponse {
		return vs[0].([]TagPropertyResponse)[vs[1].(int)]
	}).(TagPropertyResponseOutput)
}

type UpdateHistoryPropertyResponse struct {
	ImmutabilityPeriodSinceCreationInDays int    `pulumi:"immutabilityPeriodSinceCreationInDays"`
	ObjectIdentifier                      string `pulumi:"objectIdentifier"`
	TenantId                              string `pulumi:"tenantId"`
	Timestamp                             string `pulumi:"timestamp"`
	Update                                string `pulumi:"update"`
	Upn                                   string `pulumi:"upn"`
}





type UpdateHistoryPropertyResponseInput interface {
	pulumi.Input

	ToUpdateHistoryPropertyResponseOutput() UpdateHistoryPropertyResponseOutput
	ToUpdateHistoryPropertyResponseOutputWithContext(context.Context) UpdateHistoryPropertyResponseOutput
}

type UpdateHistoryPropertyResponseArgs struct {
	ImmutabilityPeriodSinceCreationInDays pulumi.IntInput    `pulumi:"immutabilityPeriodSinceCreationInDays"`
	ObjectIdentifier                      pulumi.StringInput `pulumi:"objectIdentifier"`
	TenantId                              pulumi.StringInput `pulumi:"tenantId"`
	Timestamp                             pulumi.StringInput `pulumi:"timestamp"`
	Update                                pulumi.StringInput `pulumi:"update"`
	Upn                                   pulumi.StringInput `pulumi:"upn"`
}

func (UpdateHistoryPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateHistoryPropertyResponse)(nil)).Elem()
}

func (i UpdateHistoryPropertyResponseArgs) ToUpdateHistoryPropertyResponseOutput() UpdateHistoryPropertyResponseOutput {
	return i.ToUpdateHistoryPropertyResponseOutputWithContext(context.Background())
}

func (i UpdateHistoryPropertyResponseArgs) ToUpdateHistoryPropertyResponseOutputWithContext(ctx context.Context) UpdateHistoryPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateHistoryPropertyResponseOutput)
}





type UpdateHistoryPropertyResponseArrayInput interface {
	pulumi.Input

	ToUpdateHistoryPropertyResponseArrayOutput() UpdateHistoryPropertyResponseArrayOutput
	ToUpdateHistoryPropertyResponseArrayOutputWithContext(context.Context) UpdateHistoryPropertyResponseArrayOutput
}

type UpdateHistoryPropertyResponseArray []UpdateHistoryPropertyResponseInput

func (UpdateHistoryPropertyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UpdateHistoryPropertyResponse)(nil)).Elem()
}

func (i UpdateHistoryPropertyResponseArray) ToUpdateHistoryPropertyResponseArrayOutput() UpdateHistoryPropertyResponseArrayOutput {
	return i.ToUpdateHistoryPropertyResponseArrayOutputWithContext(context.Background())
}

func (i UpdateHistoryPropertyResponseArray) ToUpdateHistoryPropertyResponseArrayOutputWithContext(ctx context.Context) UpdateHistoryPropertyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateHistoryPropertyResponseArrayOutput)
}

type UpdateHistoryPropertyResponseOutput struct{ *pulumi.OutputState }

func (UpdateHistoryPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateHistoryPropertyResponse)(nil)).Elem()
}

func (o UpdateHistoryPropertyResponseOutput) ToUpdateHistoryPropertyResponseOutput() UpdateHistoryPropertyResponseOutput {
	return o
}

func (o UpdateHistoryPropertyResponseOutput) ToUpdateHistoryPropertyResponseOutputWithContext(ctx context.Context) UpdateHistoryPropertyResponseOutput {
	return o
}

func (o UpdateHistoryPropertyResponseOutput) ImmutabilityPeriodSinceCreationInDays() pulumi.IntOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) int { return v.ImmutabilityPeriodSinceCreationInDays }).(pulumi.IntOutput)
}

func (o UpdateHistoryPropertyResponseOutput) ObjectIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.ObjectIdentifier }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) Update() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.Update }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) Upn() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.Upn }).(pulumi.StringOutput)
}

type UpdateHistoryPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (UpdateHistoryPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UpdateHistoryPropertyResponse)(nil)).Elem()
}

func (o UpdateHistoryPropertyResponseArrayOutput) ToUpdateHistoryPropertyResponseArrayOutput() UpdateHistoryPropertyResponseArrayOutput {
	return o
}

func (o UpdateHistoryPropertyResponseArrayOutput) ToUpdateHistoryPropertyResponseArrayOutputWithContext(ctx context.Context) UpdateHistoryPropertyResponseArrayOutput {
	return o
}

func (o UpdateHistoryPropertyResponseArrayOutput) Index(i pulumi.IntInput) UpdateHistoryPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UpdateHistoryPropertyResponse {
		return vs[0].([]UpdateHistoryPropertyResponse)[vs[1].(int)]
	}).(UpdateHistoryPropertyResponseOutput)
}

type VirtualNetworkRule struct {
	Action                   *Action `pulumi:"action"`
	State                    *State  `pulumi:"state"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRule) Defaults() *VirtualNetworkRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := Action("Allow")
		tmp.Action = &action_
	}
	return &tmp
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Action                   ActionPtrInput     `pulumi:"action"`
	State                    StatePtrInput      `pulumi:"state"`
	VirtualNetworkResourceId pulumi.StringInput `pulumi:"virtualNetworkResourceId"`
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *Action { return v.Action }).(ActionPtrOutput)
}

func (o VirtualNetworkRuleOutput) State() StatePtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *State { return v.State }).(StatePtrOutput)
}

func (o VirtualNetworkRuleOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRule) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Action                   *string `pulumi:"action"`
	State                    *string `pulumi:"state"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRuleResponse) Defaults() *VirtualNetworkRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type VirtualNetworkRuleResponseInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput
	ToVirtualNetworkRuleResponseOutputWithContext(context.Context) VirtualNetworkRuleResponseOutput
}

type VirtualNetworkRuleResponseArgs struct {
	Action                   pulumi.StringPtrInput `pulumi:"action"`
	State                    pulumi.StringPtrInput `pulumi:"state"`
	VirtualNetworkResourceId pulumi.StringInput    `pulumi:"virtualNetworkResourceId"`
}

func (VirtualNetworkRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return i.ToVirtualNetworkRuleResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseOutput)
}





type VirtualNetworkRuleResponseArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput
	ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Context) VirtualNetworkRuleResponseArrayOutput
}

type VirtualNetworkRuleResponseArray []VirtualNetworkRuleResponseInput

func (VirtualNetworkRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return i.ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseArrayOutput)
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveDirectoryPropertiesOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ActiveDirectoryPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureFilesIdentityBasedAuthenticationOutput{})
	pulumi.RegisterOutputType(AzureFilesIdentityBasedAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(AzureFilesIdentityBasedAuthenticationResponseOutput{})
	pulumi.RegisterOutputType(AzureFilesIdentityBasedAuthenticationResponsePtrOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicyDefinitionOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicyDefinitionResponseOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicyFilterOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicyFilterResponseOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicyRuleOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicyRuleResponseOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicyRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicySchemaOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicySchemaPtrOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicySchemaResponseOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicySchemaResponsePtrOutput{})
	pulumi.RegisterOutputType(BlobRestoreParametersResponseOutput{})
	pulumi.RegisterOutputType(BlobRestoreParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(BlobRestoreRangeResponseOutput{})
	pulumi.RegisterOutputType(BlobRestoreRangeResponseArrayOutput{})
	pulumi.RegisterOutputType(BlobRestoreStatusResponseOutput{})
	pulumi.RegisterOutputType(BlobRestoreStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ChangeFeedOutput{})
	pulumi.RegisterOutputType(ChangeFeedPtrOutput{})
	pulumi.RegisterOutputType(ChangeFeedResponseOutput{})
	pulumi.RegisterOutputType(ChangeFeedResponsePtrOutput{})
	pulumi.RegisterOutputType(CorsRuleOutput{})
	pulumi.RegisterOutputType(CorsRuleArrayOutput{})
	pulumi.RegisterOutputType(CorsRuleResponseOutput{})
	pulumi.RegisterOutputType(CorsRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(CorsRulesOutput{})
	pulumi.RegisterOutputType(CorsRulesPtrOutput{})
	pulumi.RegisterOutputType(CorsRulesResponseOutput{})
	pulumi.RegisterOutputType(CorsRulesResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomDomainOutput{})
	pulumi.RegisterOutputType(CustomDomainPtrOutput{})
	pulumi.RegisterOutputType(CustomDomainResponseOutput{})
	pulumi.RegisterOutputType(CustomDomainResponsePtrOutput{})
	pulumi.RegisterOutputType(DateAfterCreationOutput{})
	pulumi.RegisterOutputType(DateAfterCreationPtrOutput{})
	pulumi.RegisterOutputType(DateAfterCreationResponseOutput{})
	pulumi.RegisterOutputType(DateAfterCreationResponsePtrOutput{})
	pulumi.RegisterOutputType(DateAfterModificationOutput{})
	pulumi.RegisterOutputType(DateAfterModificationPtrOutput{})
	pulumi.RegisterOutputType(DateAfterModificationResponseOutput{})
	pulumi.RegisterOutputType(DateAfterModificationResponsePtrOutput{})
	pulumi.RegisterOutputType(DeleteRetentionPolicyOutput{})
	pulumi.RegisterOutputType(DeleteRetentionPolicyPtrOutput{})
	pulumi.RegisterOutputType(DeleteRetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(DeleteRetentionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionScopeKeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionScopeKeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionScopeKeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionScopeKeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServiceOutput{})
	pulumi.RegisterOutputType(EncryptionServicePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServiceResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServiceResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServicesOutput{})
	pulumi.RegisterOutputType(EncryptionServicesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionServicesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServicesResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointsResponseOutput{})
	pulumi.RegisterOutputType(EndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(GeoReplicationStatsResponseOutput{})
	pulumi.RegisterOutputType(GeoReplicationStatsResponsePtrOutput{})
	pulumi.RegisterOutputType(IPRuleOutput{})
	pulumi.RegisterOutputType(IPRuleArrayOutput{})
	pulumi.RegisterOutputType(IPRuleResponseOutput{})
	pulumi.RegisterOutputType(IPRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ImmutabilityPolicyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ImmutabilityPolicyPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LastAccessTimeTrackingPolicyOutput{})
	pulumi.RegisterOutputType(LastAccessTimeTrackingPolicyPtrOutput{})
	pulumi.RegisterOutputType(LastAccessTimeTrackingPolicyResponseOutput{})
	pulumi.RegisterOutputType(LastAccessTimeTrackingPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(LegalHoldPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LegalHoldPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyActionOutput{})
	pulumi.RegisterOutputType(ManagementPolicyActionResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyBaseBlobOutput{})
	pulumi.RegisterOutputType(ManagementPolicyBaseBlobPtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyBaseBlobResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyBaseBlobResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyDefinitionOutput{})
	pulumi.RegisterOutputType(ManagementPolicyDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyFilterOutput{})
	pulumi.RegisterOutputType(ManagementPolicyFilterPtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyFilterResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyRuleOutput{})
	pulumi.RegisterOutputType(ManagementPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(ManagementPolicyRuleResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementPolicySchemaOutput{})
	pulumi.RegisterOutputType(ManagementPolicySchemaPtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicySchemaResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicySchemaResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotPtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyVersionOutput{})
	pulumi.RegisterOutputType(ManagementPolicyVersionPtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyVersionResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyVersionResponsePtrOutput{})
	pulumi.RegisterOutputType(MultichannelOutput{})
	pulumi.RegisterOutputType(MultichannelPtrOutput{})
	pulumi.RegisterOutputType(MultichannelResponseOutput{})
	pulumi.RegisterOutputType(MultichannelResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(ObjectReplicationPolicyFilterOutput{})
	pulumi.RegisterOutputType(ObjectReplicationPolicyFilterPtrOutput{})
	pulumi.RegisterOutputType(ObjectReplicationPolicyFilterResponseOutput{})
	pulumi.RegisterOutputType(ObjectReplicationPolicyFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(ObjectReplicationPolicyRuleOutput{})
	pulumi.RegisterOutputType(ObjectReplicationPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(ObjectReplicationPolicyRuleResponseOutput{})
	pulumi.RegisterOutputType(ObjectReplicationPolicyRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ProtocolSettingsOutput{})
	pulumi.RegisterOutputType(ProtocolSettingsPtrOutput{})
	pulumi.RegisterOutputType(ProtocolSettingsResponseOutput{})
	pulumi.RegisterOutputType(ProtocolSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceAccessRuleOutput{})
	pulumi.RegisterOutputType(ResourceAccessRuleArrayOutput{})
	pulumi.RegisterOutputType(ResourceAccessRuleResponseOutput{})
	pulumi.RegisterOutputType(ResourceAccessRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(RestorePolicyPropertiesOutput{})
	pulumi.RegisterOutputType(RestorePolicyPropertiesPtrOutput{})
	pulumi.RegisterOutputType(RestorePolicyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RestorePolicyPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(RoutingPreferenceOutput{})
	pulumi.RegisterOutputType(RoutingPreferencePtrOutput{})
	pulumi.RegisterOutputType(RoutingPreferenceResponseOutput{})
	pulumi.RegisterOutputType(RoutingPreferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SmbSettingOutput{})
	pulumi.RegisterOutputType(SmbSettingPtrOutput{})
	pulumi.RegisterOutputType(SmbSettingResponseOutput{})
	pulumi.RegisterOutputType(SmbSettingResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountInternetEndpointsResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountInternetEndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountKeyResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountMicrosoftEndpointsResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountMicrosoftEndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TagFilterOutput{})
	pulumi.RegisterOutputType(TagFilterArrayOutput{})
	pulumi.RegisterOutputType(TagFilterResponseOutput{})
	pulumi.RegisterOutputType(TagFilterResponseArrayOutput{})
	pulumi.RegisterOutputType(TagPropertyResponseOutput{})
	pulumi.RegisterOutputType(TagPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(UpdateHistoryPropertyResponseOutput{})
	pulumi.RegisterOutputType(UpdateHistoryPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
