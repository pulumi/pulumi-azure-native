


package v20190601

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

type BlobInventoryPolicySchemaResponse struct {
	Destination string                            `pulumi:"destination"`
	Enabled     bool                              `pulumi:"enabled"`
	Rules       []BlobInventoryPolicyRuleResponse `pulumi:"rules"`
	Type        string                            `pulumi:"type"`
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

type BlobRestoreParametersResponse struct {
	BlobRanges    []BlobRestoreRangeResponse `pulumi:"blobRanges"`
	TimeToRestore string                     `pulumi:"timeToRestore"`
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

func (o BlobRestoreParametersResponseOutput) BlobRanges() BlobRestoreRangeResponseArrayOutput {
	return o.ApplyT(func(v BlobRestoreParametersResponse) []BlobRestoreRangeResponse { return v.BlobRanges }).(BlobRestoreRangeResponseArrayOutput)
}

func (o BlobRestoreParametersResponseOutput) TimeToRestore() pulumi.StringOutput {
	return o.ApplyT(func(v BlobRestoreParametersResponse) string { return v.TimeToRestore }).(pulumi.StringOutput)
}

type BlobRestoreRangeResponse struct {
	EndRange   string `pulumi:"endRange"`
	StartRange string `pulumi:"startRange"`
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

func (o CustomDomainResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomainResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDomainResponseOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
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

type GeoReplicationStatsResponse struct {
	CanFailover  bool   `pulumi:"canFailover"`
	LastSyncTime string `pulumi:"lastSyncTime"`
	Status       string `pulumi:"status"`
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

func (o GeoReplicationStatsResponseOutput) CanFailover() pulumi.BoolOutput {
	return o.ApplyT(func(v GeoReplicationStatsResponse) bool { return v.CanFailover }).(pulumi.BoolOutput)
}

func (o GeoReplicationStatsResponseOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v GeoReplicationStatsResponse) string { return v.LastSyncTime }).(pulumi.StringOutput)
}

func (o GeoReplicationStatsResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GeoReplicationStatsResponse) string { return v.Status }).(pulumi.StringOutput)
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

func (o LegalHoldPropertiesResponseOutput) HasLegalHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LegalHoldPropertiesResponse) bool { return v.HasLegalHold }).(pulumi.BoolOutput)
}

func (o LegalHoldPropertiesResponseOutput) Tags() TagPropertyResponseArrayOutput {
	return o.ApplyT(func(v LegalHoldPropertiesResponse) []TagPropertyResponse { return v.Tags }).(TagPropertyResponseArrayOutput)
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

func (o ManagementPolicySchemaOutput) Rules() ManagementPolicyRuleArrayOutput {
	return o.ApplyT(func(v ManagementPolicySchema) []ManagementPolicyRule { return v.Rules }).(ManagementPolicyRuleArrayOutput)
}

type ManagementPolicySchemaResponse struct {
	Rules []ManagementPolicyRuleResponse `pulumi:"rules"`
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

func (o ManagementPolicySchemaResponseOutput) Rules() ManagementPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v ManagementPolicySchemaResponse) []ManagementPolicyRuleResponse { return v.Rules }).(ManagementPolicyRuleResponseArrayOutput)
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

type NetworkRuleSet struct {
	Bypass              *string              `pulumi:"bypass"`
	DefaultAction       DefaultAction        `pulumi:"defaultAction"`
	IpRules             []IPRule             `pulumi:"ipRules"`
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

func (o NetworkRuleSetResponseOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponseOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) string { return v.DefaultAction }).(pulumi.StringOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IPRuleResponse { return v.IpRules }).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponseOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
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

func (o PrivateLinkServiceConnectionStateOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         *string `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
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

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
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

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type StorageAccountInternetEndpointsResponse struct {
	Blob string `pulumi:"blob"`
	Dfs  string `pulumi:"dfs"`
	File string `pulumi:"file"`
	Web  string `pulumi:"web"`
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

type StorageAccountMicrosoftEndpointsResponse struct {
	Blob  string `pulumi:"blob"`
	Dfs   string `pulumi:"dfs"`
	File  string `pulumi:"file"`
	Queue string `pulumi:"queue"`
	Table string `pulumi:"table"`
	Web   string `pulumi:"web"`
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
	pulumi.RegisterOutputType(BlobInventoryPolicySchemaResponseOutput{})
	pulumi.RegisterOutputType(BlobRestoreParametersResponseOutput{})
	pulumi.RegisterOutputType(BlobRestoreRangeResponseOutput{})
	pulumi.RegisterOutputType(BlobRestoreRangeResponseArrayOutput{})
	pulumi.RegisterOutputType(BlobRestoreStatusResponseOutput{})
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
	pulumi.RegisterOutputType(GeoReplicationStatsResponseOutput{})
	pulumi.RegisterOutputType(IPRuleOutput{})
	pulumi.RegisterOutputType(IPRuleArrayOutput{})
	pulumi.RegisterOutputType(IPRuleResponseOutput{})
	pulumi.RegisterOutputType(IPRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ImmutabilityPolicyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LastAccessTimeTrackingPolicyOutput{})
	pulumi.RegisterOutputType(LastAccessTimeTrackingPolicyPtrOutput{})
	pulumi.RegisterOutputType(LastAccessTimeTrackingPolicyResponseOutput{})
	pulumi.RegisterOutputType(LastAccessTimeTrackingPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(LegalHoldPropertiesResponseOutput{})
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
	pulumi.RegisterOutputType(ManagementPolicySchemaResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotPtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicySnapShotResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyVersionOutput{})
	pulumi.RegisterOutputType(ManagementPolicyVersionPtrOutput{})
	pulumi.RegisterOutputType(ManagementPolicyVersionResponseOutput{})
	pulumi.RegisterOutputType(ManagementPolicyVersionResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
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
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(RestorePolicyPropertiesOutput{})
	pulumi.RegisterOutputType(RestorePolicyPropertiesPtrOutput{})
	pulumi.RegisterOutputType(RestorePolicyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RestorePolicyPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(RoutingPreferenceOutput{})
	pulumi.RegisterOutputType(RoutingPreferencePtrOutput{})
	pulumi.RegisterOutputType(RoutingPreferenceResponseOutput{})
	pulumi.RegisterOutputType(RoutingPreferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountInternetEndpointsResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountInternetEndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountMicrosoftEndpointsResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountMicrosoftEndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
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
