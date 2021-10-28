


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountCredentialDetailsResponse struct {
	AccountConnectionString string                           `pulumi:"accountConnectionString"`
	AccountName             string                           `pulumi:"accountName"`
	DataAccountType         string                           `pulumi:"dataAccountType"`
	ShareCredentialDetails  []ShareCredentialDetailsResponse `pulumi:"shareCredentialDetails"`
}





type AccountCredentialDetailsResponseInput interface {
	pulumi.Input

	ToAccountCredentialDetailsResponseOutput() AccountCredentialDetailsResponseOutput
	ToAccountCredentialDetailsResponseOutputWithContext(context.Context) AccountCredentialDetailsResponseOutput
}

type AccountCredentialDetailsResponseArgs struct {
	AccountConnectionString pulumi.StringInput                       `pulumi:"accountConnectionString"`
	AccountName             pulumi.StringInput                       `pulumi:"accountName"`
	DataAccountType         pulumi.StringInput                       `pulumi:"dataAccountType"`
	ShareCredentialDetails  ShareCredentialDetailsResponseArrayInput `pulumi:"shareCredentialDetails"`
}

func (AccountCredentialDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountCredentialDetailsResponse)(nil)).Elem()
}

func (i AccountCredentialDetailsResponseArgs) ToAccountCredentialDetailsResponseOutput() AccountCredentialDetailsResponseOutput {
	return i.ToAccountCredentialDetailsResponseOutputWithContext(context.Background())
}

func (i AccountCredentialDetailsResponseArgs) ToAccountCredentialDetailsResponseOutputWithContext(ctx context.Context) AccountCredentialDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountCredentialDetailsResponseOutput)
}





type AccountCredentialDetailsResponseArrayInput interface {
	pulumi.Input

	ToAccountCredentialDetailsResponseArrayOutput() AccountCredentialDetailsResponseArrayOutput
	ToAccountCredentialDetailsResponseArrayOutputWithContext(context.Context) AccountCredentialDetailsResponseArrayOutput
}

type AccountCredentialDetailsResponseArray []AccountCredentialDetailsResponseInput

func (AccountCredentialDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccountCredentialDetailsResponse)(nil)).Elem()
}

func (i AccountCredentialDetailsResponseArray) ToAccountCredentialDetailsResponseArrayOutput() AccountCredentialDetailsResponseArrayOutput {
	return i.ToAccountCredentialDetailsResponseArrayOutputWithContext(context.Background())
}

func (i AccountCredentialDetailsResponseArray) ToAccountCredentialDetailsResponseArrayOutputWithContext(ctx context.Context) AccountCredentialDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountCredentialDetailsResponseArrayOutput)
}

type AccountCredentialDetailsResponseOutput struct{ *pulumi.OutputState }

func (AccountCredentialDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountCredentialDetailsResponse)(nil)).Elem()
}

func (o AccountCredentialDetailsResponseOutput) ToAccountCredentialDetailsResponseOutput() AccountCredentialDetailsResponseOutput {
	return o
}

func (o AccountCredentialDetailsResponseOutput) ToAccountCredentialDetailsResponseOutputWithContext(ctx context.Context) AccountCredentialDetailsResponseOutput {
	return o
}

func (o AccountCredentialDetailsResponseOutput) AccountConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v AccountCredentialDetailsResponse) string { return v.AccountConnectionString }).(pulumi.StringOutput)
}

func (o AccountCredentialDetailsResponseOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v AccountCredentialDetailsResponse) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o AccountCredentialDetailsResponseOutput) DataAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v AccountCredentialDetailsResponse) string { return v.DataAccountType }).(pulumi.StringOutput)
}

func (o AccountCredentialDetailsResponseOutput) ShareCredentialDetails() ShareCredentialDetailsResponseArrayOutput {
	return o.ApplyT(func(v AccountCredentialDetailsResponse) []ShareCredentialDetailsResponse {
		return v.ShareCredentialDetails
	}).(ShareCredentialDetailsResponseArrayOutput)
}

type AccountCredentialDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (AccountCredentialDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccountCredentialDetailsResponse)(nil)).Elem()
}

func (o AccountCredentialDetailsResponseArrayOutput) ToAccountCredentialDetailsResponseArrayOutput() AccountCredentialDetailsResponseArrayOutput {
	return o
}

func (o AccountCredentialDetailsResponseArrayOutput) ToAccountCredentialDetailsResponseArrayOutputWithContext(ctx context.Context) AccountCredentialDetailsResponseArrayOutput {
	return o
}

func (o AccountCredentialDetailsResponseArrayOutput) Index(i pulumi.IntInput) AccountCredentialDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccountCredentialDetailsResponse {
		return vs[0].([]AccountCredentialDetailsResponse)[vs[1].(int)]
	}).(AccountCredentialDetailsResponseOutput)
}

type AdditionalErrorInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type *string     `pulumi:"type"`
}





type AdditionalErrorInfoResponseInput interface {
	pulumi.Input

	ToAdditionalErrorInfoResponseOutput() AdditionalErrorInfoResponseOutput
	ToAdditionalErrorInfoResponseOutputWithContext(context.Context) AdditionalErrorInfoResponseOutput
}

type AdditionalErrorInfoResponseArgs struct {
	Info pulumi.Input          `pulumi:"info"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (AdditionalErrorInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalErrorInfoResponse)(nil)).Elem()
}

func (i AdditionalErrorInfoResponseArgs) ToAdditionalErrorInfoResponseOutput() AdditionalErrorInfoResponseOutput {
	return i.ToAdditionalErrorInfoResponseOutputWithContext(context.Background())
}

func (i AdditionalErrorInfoResponseArgs) ToAdditionalErrorInfoResponseOutputWithContext(ctx context.Context) AdditionalErrorInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalErrorInfoResponseOutput)
}





type AdditionalErrorInfoResponseArrayInput interface {
	pulumi.Input

	ToAdditionalErrorInfoResponseArrayOutput() AdditionalErrorInfoResponseArrayOutput
	ToAdditionalErrorInfoResponseArrayOutputWithContext(context.Context) AdditionalErrorInfoResponseArrayOutput
}

type AdditionalErrorInfoResponseArray []AdditionalErrorInfoResponseInput

func (AdditionalErrorInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalErrorInfoResponse)(nil)).Elem()
}

func (i AdditionalErrorInfoResponseArray) ToAdditionalErrorInfoResponseArrayOutput() AdditionalErrorInfoResponseArrayOutput {
	return i.ToAdditionalErrorInfoResponseArrayOutputWithContext(context.Background())
}

func (i AdditionalErrorInfoResponseArray) ToAdditionalErrorInfoResponseArrayOutputWithContext(ctx context.Context) AdditionalErrorInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalErrorInfoResponseArrayOutput)
}

type AdditionalErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (AdditionalErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalErrorInfoResponse)(nil)).Elem()
}

func (o AdditionalErrorInfoResponseOutput) ToAdditionalErrorInfoResponseOutput() AdditionalErrorInfoResponseOutput {
	return o
}

func (o AdditionalErrorInfoResponseOutput) ToAdditionalErrorInfoResponseOutputWithContext(ctx context.Context) AdditionalErrorInfoResponseOutput {
	return o
}

func (o AdditionalErrorInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v AdditionalErrorInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o AdditionalErrorInfoResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalErrorInfoResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AdditionalErrorInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (AdditionalErrorInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalErrorInfoResponse)(nil)).Elem()
}

func (o AdditionalErrorInfoResponseArrayOutput) ToAdditionalErrorInfoResponseArrayOutput() AdditionalErrorInfoResponseArrayOutput {
	return o
}

func (o AdditionalErrorInfoResponseArrayOutput) ToAdditionalErrorInfoResponseArrayOutputWithContext(ctx context.Context) AdditionalErrorInfoResponseArrayOutput {
	return o
}

func (o AdditionalErrorInfoResponseArrayOutput) Index(i pulumi.IntInput) AdditionalErrorInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalErrorInfoResponse {
		return vs[0].([]AdditionalErrorInfoResponse)[vs[1].(int)]
	}).(AdditionalErrorInfoResponseOutput)
}

type ApplianceNetworkConfigurationResponse struct {
	MacAddress string `pulumi:"macAddress"`
	Name       string `pulumi:"name"`
}





type ApplianceNetworkConfigurationResponseInput interface {
	pulumi.Input

	ToApplianceNetworkConfigurationResponseOutput() ApplianceNetworkConfigurationResponseOutput
	ToApplianceNetworkConfigurationResponseOutputWithContext(context.Context) ApplianceNetworkConfigurationResponseOutput
}

type ApplianceNetworkConfigurationResponseArgs struct {
	MacAddress pulumi.StringInput `pulumi:"macAddress"`
	Name       pulumi.StringInput `pulumi:"name"`
}

func (ApplianceNetworkConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplianceNetworkConfigurationResponse)(nil)).Elem()
}

func (i ApplianceNetworkConfigurationResponseArgs) ToApplianceNetworkConfigurationResponseOutput() ApplianceNetworkConfigurationResponseOutput {
	return i.ToApplianceNetworkConfigurationResponseOutputWithContext(context.Background())
}

func (i ApplianceNetworkConfigurationResponseArgs) ToApplianceNetworkConfigurationResponseOutputWithContext(ctx context.Context) ApplianceNetworkConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceNetworkConfigurationResponseOutput)
}





type ApplianceNetworkConfigurationResponseArrayInput interface {
	pulumi.Input

	ToApplianceNetworkConfigurationResponseArrayOutput() ApplianceNetworkConfigurationResponseArrayOutput
	ToApplianceNetworkConfigurationResponseArrayOutputWithContext(context.Context) ApplianceNetworkConfigurationResponseArrayOutput
}

type ApplianceNetworkConfigurationResponseArray []ApplianceNetworkConfigurationResponseInput

func (ApplianceNetworkConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplianceNetworkConfigurationResponse)(nil)).Elem()
}

func (i ApplianceNetworkConfigurationResponseArray) ToApplianceNetworkConfigurationResponseArrayOutput() ApplianceNetworkConfigurationResponseArrayOutput {
	return i.ToApplianceNetworkConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i ApplianceNetworkConfigurationResponseArray) ToApplianceNetworkConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplianceNetworkConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceNetworkConfigurationResponseArrayOutput)
}

type ApplianceNetworkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ApplianceNetworkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplianceNetworkConfigurationResponse)(nil)).Elem()
}

func (o ApplianceNetworkConfigurationResponseOutput) ToApplianceNetworkConfigurationResponseOutput() ApplianceNetworkConfigurationResponseOutput {
	return o
}

func (o ApplianceNetworkConfigurationResponseOutput) ToApplianceNetworkConfigurationResponseOutputWithContext(ctx context.Context) ApplianceNetworkConfigurationResponseOutput {
	return o
}

func (o ApplianceNetworkConfigurationResponseOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v ApplianceNetworkConfigurationResponse) string { return v.MacAddress }).(pulumi.StringOutput)
}

func (o ApplianceNetworkConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplianceNetworkConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ApplianceNetworkConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplianceNetworkConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplianceNetworkConfigurationResponse)(nil)).Elem()
}

func (o ApplianceNetworkConfigurationResponseArrayOutput) ToApplianceNetworkConfigurationResponseArrayOutput() ApplianceNetworkConfigurationResponseArrayOutput {
	return o
}

func (o ApplianceNetworkConfigurationResponseArrayOutput) ToApplianceNetworkConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplianceNetworkConfigurationResponseArrayOutput {
	return o
}

func (o ApplianceNetworkConfigurationResponseArrayOutput) Index(i pulumi.IntInput) ApplianceNetworkConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplianceNetworkConfigurationResponse {
		return vs[0].([]ApplianceNetworkConfigurationResponse)[vs[1].(int)]
	}).(ApplianceNetworkConfigurationResponseOutput)
}

type AzureFileFilterDetails struct {
	FilePathList   []string `pulumi:"filePathList"`
	FilePrefixList []string `pulumi:"filePrefixList"`
	FileShareList  []string `pulumi:"fileShareList"`
}





type AzureFileFilterDetailsInput interface {
	pulumi.Input

	ToAzureFileFilterDetailsOutput() AzureFileFilterDetailsOutput
	ToAzureFileFilterDetailsOutputWithContext(context.Context) AzureFileFilterDetailsOutput
}

type AzureFileFilterDetailsArgs struct {
	FilePathList   pulumi.StringArrayInput `pulumi:"filePathList"`
	FilePrefixList pulumi.StringArrayInput `pulumi:"filePrefixList"`
	FileShareList  pulumi.StringArrayInput `pulumi:"fileShareList"`
}

func (AzureFileFilterDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileFilterDetails)(nil)).Elem()
}

func (i AzureFileFilterDetailsArgs) ToAzureFileFilterDetailsOutput() AzureFileFilterDetailsOutput {
	return i.ToAzureFileFilterDetailsOutputWithContext(context.Background())
}

func (i AzureFileFilterDetailsArgs) ToAzureFileFilterDetailsOutputWithContext(ctx context.Context) AzureFileFilterDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileFilterDetailsOutput)
}

func (i AzureFileFilterDetailsArgs) ToAzureFileFilterDetailsPtrOutput() AzureFileFilterDetailsPtrOutput {
	return i.ToAzureFileFilterDetailsPtrOutputWithContext(context.Background())
}

func (i AzureFileFilterDetailsArgs) ToAzureFileFilterDetailsPtrOutputWithContext(ctx context.Context) AzureFileFilterDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileFilterDetailsOutput).ToAzureFileFilterDetailsPtrOutputWithContext(ctx)
}









type AzureFileFilterDetailsPtrInput interface {
	pulumi.Input

	ToAzureFileFilterDetailsPtrOutput() AzureFileFilterDetailsPtrOutput
	ToAzureFileFilterDetailsPtrOutputWithContext(context.Context) AzureFileFilterDetailsPtrOutput
}

type azureFileFilterDetailsPtrType AzureFileFilterDetailsArgs

func AzureFileFilterDetailsPtr(v *AzureFileFilterDetailsArgs) AzureFileFilterDetailsPtrInput {
	return (*azureFileFilterDetailsPtrType)(v)
}

func (*azureFileFilterDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileFilterDetails)(nil)).Elem()
}

func (i *azureFileFilterDetailsPtrType) ToAzureFileFilterDetailsPtrOutput() AzureFileFilterDetailsPtrOutput {
	return i.ToAzureFileFilterDetailsPtrOutputWithContext(context.Background())
}

func (i *azureFileFilterDetailsPtrType) ToAzureFileFilterDetailsPtrOutputWithContext(ctx context.Context) AzureFileFilterDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileFilterDetailsPtrOutput)
}

type AzureFileFilterDetailsOutput struct{ *pulumi.OutputState }

func (AzureFileFilterDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileFilterDetails)(nil)).Elem()
}

func (o AzureFileFilterDetailsOutput) ToAzureFileFilterDetailsOutput() AzureFileFilterDetailsOutput {
	return o
}

func (o AzureFileFilterDetailsOutput) ToAzureFileFilterDetailsOutputWithContext(ctx context.Context) AzureFileFilterDetailsOutput {
	return o
}

func (o AzureFileFilterDetailsOutput) ToAzureFileFilterDetailsPtrOutput() AzureFileFilterDetailsPtrOutput {
	return o.ToAzureFileFilterDetailsPtrOutputWithContext(context.Background())
}

func (o AzureFileFilterDetailsOutput) ToAzureFileFilterDetailsPtrOutputWithContext(ctx context.Context) AzureFileFilterDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFileFilterDetails) *AzureFileFilterDetails {
		return &v
	}).(AzureFileFilterDetailsPtrOutput)
}

func (o AzureFileFilterDetailsOutput) FilePathList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureFileFilterDetails) []string { return v.FilePathList }).(pulumi.StringArrayOutput)
}

func (o AzureFileFilterDetailsOutput) FilePrefixList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureFileFilterDetails) []string { return v.FilePrefixList }).(pulumi.StringArrayOutput)
}

func (o AzureFileFilterDetailsOutput) FileShareList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureFileFilterDetails) []string { return v.FileShareList }).(pulumi.StringArrayOutput)
}

type AzureFileFilterDetailsPtrOutput struct{ *pulumi.OutputState }

func (AzureFileFilterDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileFilterDetails)(nil)).Elem()
}

func (o AzureFileFilterDetailsPtrOutput) ToAzureFileFilterDetailsPtrOutput() AzureFileFilterDetailsPtrOutput {
	return o
}

func (o AzureFileFilterDetailsPtrOutput) ToAzureFileFilterDetailsPtrOutputWithContext(ctx context.Context) AzureFileFilterDetailsPtrOutput {
	return o
}

func (o AzureFileFilterDetailsPtrOutput) Elem() AzureFileFilterDetailsOutput {
	return o.ApplyT(func(v *AzureFileFilterDetails) AzureFileFilterDetails {
		if v != nil {
			return *v
		}
		var ret AzureFileFilterDetails
		return ret
	}).(AzureFileFilterDetailsOutput)
}

func (o AzureFileFilterDetailsPtrOutput) FilePathList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureFileFilterDetails) []string {
		if v == nil {
			return nil
		}
		return v.FilePathList
	}).(pulumi.StringArrayOutput)
}

func (o AzureFileFilterDetailsPtrOutput) FilePrefixList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureFileFilterDetails) []string {
		if v == nil {
			return nil
		}
		return v.FilePrefixList
	}).(pulumi.StringArrayOutput)
}

func (o AzureFileFilterDetailsPtrOutput) FileShareList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureFileFilterDetails) []string {
		if v == nil {
			return nil
		}
		return v.FileShareList
	}).(pulumi.StringArrayOutput)
}

type AzureFileFilterDetailsResponse struct {
	FilePathList   []string `pulumi:"filePathList"`
	FilePrefixList []string `pulumi:"filePrefixList"`
	FileShareList  []string `pulumi:"fileShareList"`
}





type AzureFileFilterDetailsResponseInput interface {
	pulumi.Input

	ToAzureFileFilterDetailsResponseOutput() AzureFileFilterDetailsResponseOutput
	ToAzureFileFilterDetailsResponseOutputWithContext(context.Context) AzureFileFilterDetailsResponseOutput
}

type AzureFileFilterDetailsResponseArgs struct {
	FilePathList   pulumi.StringArrayInput `pulumi:"filePathList"`
	FilePrefixList pulumi.StringArrayInput `pulumi:"filePrefixList"`
	FileShareList  pulumi.StringArrayInput `pulumi:"fileShareList"`
}

func (AzureFileFilterDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileFilterDetailsResponse)(nil)).Elem()
}

func (i AzureFileFilterDetailsResponseArgs) ToAzureFileFilterDetailsResponseOutput() AzureFileFilterDetailsResponseOutput {
	return i.ToAzureFileFilterDetailsResponseOutputWithContext(context.Background())
}

func (i AzureFileFilterDetailsResponseArgs) ToAzureFileFilterDetailsResponseOutputWithContext(ctx context.Context) AzureFileFilterDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileFilterDetailsResponseOutput)
}

func (i AzureFileFilterDetailsResponseArgs) ToAzureFileFilterDetailsResponsePtrOutput() AzureFileFilterDetailsResponsePtrOutput {
	return i.ToAzureFileFilterDetailsResponsePtrOutputWithContext(context.Background())
}

func (i AzureFileFilterDetailsResponseArgs) ToAzureFileFilterDetailsResponsePtrOutputWithContext(ctx context.Context) AzureFileFilterDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileFilterDetailsResponseOutput).ToAzureFileFilterDetailsResponsePtrOutputWithContext(ctx)
}









type AzureFileFilterDetailsResponsePtrInput interface {
	pulumi.Input

	ToAzureFileFilterDetailsResponsePtrOutput() AzureFileFilterDetailsResponsePtrOutput
	ToAzureFileFilterDetailsResponsePtrOutputWithContext(context.Context) AzureFileFilterDetailsResponsePtrOutput
}

type azureFileFilterDetailsResponsePtrType AzureFileFilterDetailsResponseArgs

func AzureFileFilterDetailsResponsePtr(v *AzureFileFilterDetailsResponseArgs) AzureFileFilterDetailsResponsePtrInput {
	return (*azureFileFilterDetailsResponsePtrType)(v)
}

func (*azureFileFilterDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileFilterDetailsResponse)(nil)).Elem()
}

func (i *azureFileFilterDetailsResponsePtrType) ToAzureFileFilterDetailsResponsePtrOutput() AzureFileFilterDetailsResponsePtrOutput {
	return i.ToAzureFileFilterDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *azureFileFilterDetailsResponsePtrType) ToAzureFileFilterDetailsResponsePtrOutputWithContext(ctx context.Context) AzureFileFilterDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileFilterDetailsResponsePtrOutput)
}

type AzureFileFilterDetailsResponseOutput struct{ *pulumi.OutputState }

func (AzureFileFilterDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileFilterDetailsResponse)(nil)).Elem()
}

func (o AzureFileFilterDetailsResponseOutput) ToAzureFileFilterDetailsResponseOutput() AzureFileFilterDetailsResponseOutput {
	return o
}

func (o AzureFileFilterDetailsResponseOutput) ToAzureFileFilterDetailsResponseOutputWithContext(ctx context.Context) AzureFileFilterDetailsResponseOutput {
	return o
}

func (o AzureFileFilterDetailsResponseOutput) ToAzureFileFilterDetailsResponsePtrOutput() AzureFileFilterDetailsResponsePtrOutput {
	return o.ToAzureFileFilterDetailsResponsePtrOutputWithContext(context.Background())
}

func (o AzureFileFilterDetailsResponseOutput) ToAzureFileFilterDetailsResponsePtrOutputWithContext(ctx context.Context) AzureFileFilterDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFileFilterDetailsResponse) *AzureFileFilterDetailsResponse {
		return &v
	}).(AzureFileFilterDetailsResponsePtrOutput)
}

func (o AzureFileFilterDetailsResponseOutput) FilePathList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureFileFilterDetailsResponse) []string { return v.FilePathList }).(pulumi.StringArrayOutput)
}

func (o AzureFileFilterDetailsResponseOutput) FilePrefixList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureFileFilterDetailsResponse) []string { return v.FilePrefixList }).(pulumi.StringArrayOutput)
}

func (o AzureFileFilterDetailsResponseOutput) FileShareList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureFileFilterDetailsResponse) []string { return v.FileShareList }).(pulumi.StringArrayOutput)
}

type AzureFileFilterDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureFileFilterDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileFilterDetailsResponse)(nil)).Elem()
}

func (o AzureFileFilterDetailsResponsePtrOutput) ToAzureFileFilterDetailsResponsePtrOutput() AzureFileFilterDetailsResponsePtrOutput {
	return o
}

func (o AzureFileFilterDetailsResponsePtrOutput) ToAzureFileFilterDetailsResponsePtrOutputWithContext(ctx context.Context) AzureFileFilterDetailsResponsePtrOutput {
	return o
}

func (o AzureFileFilterDetailsResponsePtrOutput) Elem() AzureFileFilterDetailsResponseOutput {
	return o.ApplyT(func(v *AzureFileFilterDetailsResponse) AzureFileFilterDetailsResponse {
		if v != nil {
			return *v
		}
		var ret AzureFileFilterDetailsResponse
		return ret
	}).(AzureFileFilterDetailsResponseOutput)
}

func (o AzureFileFilterDetailsResponsePtrOutput) FilePathList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureFileFilterDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.FilePathList
	}).(pulumi.StringArrayOutput)
}

func (o AzureFileFilterDetailsResponsePtrOutput) FilePrefixList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureFileFilterDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.FilePrefixList
	}).(pulumi.StringArrayOutput)
}

func (o AzureFileFilterDetailsResponsePtrOutput) FileShareList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureFileFilterDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.FileShareList
	}).(pulumi.StringArrayOutput)
}

type BlobFilterDetails struct {
	BlobPathList   []string `pulumi:"blobPathList"`
	BlobPrefixList []string `pulumi:"blobPrefixList"`
	ContainerList  []string `pulumi:"containerList"`
}





type BlobFilterDetailsInput interface {
	pulumi.Input

	ToBlobFilterDetailsOutput() BlobFilterDetailsOutput
	ToBlobFilterDetailsOutputWithContext(context.Context) BlobFilterDetailsOutput
}

type BlobFilterDetailsArgs struct {
	BlobPathList   pulumi.StringArrayInput `pulumi:"blobPathList"`
	BlobPrefixList pulumi.StringArrayInput `pulumi:"blobPrefixList"`
	ContainerList  pulumi.StringArrayInput `pulumi:"containerList"`
}

func (BlobFilterDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobFilterDetails)(nil)).Elem()
}

func (i BlobFilterDetailsArgs) ToBlobFilterDetailsOutput() BlobFilterDetailsOutput {
	return i.ToBlobFilterDetailsOutputWithContext(context.Background())
}

func (i BlobFilterDetailsArgs) ToBlobFilterDetailsOutputWithContext(ctx context.Context) BlobFilterDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobFilterDetailsOutput)
}

func (i BlobFilterDetailsArgs) ToBlobFilterDetailsPtrOutput() BlobFilterDetailsPtrOutput {
	return i.ToBlobFilterDetailsPtrOutputWithContext(context.Background())
}

func (i BlobFilterDetailsArgs) ToBlobFilterDetailsPtrOutputWithContext(ctx context.Context) BlobFilterDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobFilterDetailsOutput).ToBlobFilterDetailsPtrOutputWithContext(ctx)
}









type BlobFilterDetailsPtrInput interface {
	pulumi.Input

	ToBlobFilterDetailsPtrOutput() BlobFilterDetailsPtrOutput
	ToBlobFilterDetailsPtrOutputWithContext(context.Context) BlobFilterDetailsPtrOutput
}

type blobFilterDetailsPtrType BlobFilterDetailsArgs

func BlobFilterDetailsPtr(v *BlobFilterDetailsArgs) BlobFilterDetailsPtrInput {
	return (*blobFilterDetailsPtrType)(v)
}

func (*blobFilterDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobFilterDetails)(nil)).Elem()
}

func (i *blobFilterDetailsPtrType) ToBlobFilterDetailsPtrOutput() BlobFilterDetailsPtrOutput {
	return i.ToBlobFilterDetailsPtrOutputWithContext(context.Background())
}

func (i *blobFilterDetailsPtrType) ToBlobFilterDetailsPtrOutputWithContext(ctx context.Context) BlobFilterDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobFilterDetailsPtrOutput)
}

type BlobFilterDetailsOutput struct{ *pulumi.OutputState }

func (BlobFilterDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobFilterDetails)(nil)).Elem()
}

func (o BlobFilterDetailsOutput) ToBlobFilterDetailsOutput() BlobFilterDetailsOutput {
	return o
}

func (o BlobFilterDetailsOutput) ToBlobFilterDetailsOutputWithContext(ctx context.Context) BlobFilterDetailsOutput {
	return o
}

func (o BlobFilterDetailsOutput) ToBlobFilterDetailsPtrOutput() BlobFilterDetailsPtrOutput {
	return o.ToBlobFilterDetailsPtrOutputWithContext(context.Background())
}

func (o BlobFilterDetailsOutput) ToBlobFilterDetailsPtrOutputWithContext(ctx context.Context) BlobFilterDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobFilterDetails) *BlobFilterDetails {
		return &v
	}).(BlobFilterDetailsPtrOutput)
}

func (o BlobFilterDetailsOutput) BlobPathList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BlobFilterDetails) []string { return v.BlobPathList }).(pulumi.StringArrayOutput)
}

func (o BlobFilterDetailsOutput) BlobPrefixList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BlobFilterDetails) []string { return v.BlobPrefixList }).(pulumi.StringArrayOutput)
}

func (o BlobFilterDetailsOutput) ContainerList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BlobFilterDetails) []string { return v.ContainerList }).(pulumi.StringArrayOutput)
}

type BlobFilterDetailsPtrOutput struct{ *pulumi.OutputState }

func (BlobFilterDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobFilterDetails)(nil)).Elem()
}

func (o BlobFilterDetailsPtrOutput) ToBlobFilterDetailsPtrOutput() BlobFilterDetailsPtrOutput {
	return o
}

func (o BlobFilterDetailsPtrOutput) ToBlobFilterDetailsPtrOutputWithContext(ctx context.Context) BlobFilterDetailsPtrOutput {
	return o
}

func (o BlobFilterDetailsPtrOutput) Elem() BlobFilterDetailsOutput {
	return o.ApplyT(func(v *BlobFilterDetails) BlobFilterDetails {
		if v != nil {
			return *v
		}
		var ret BlobFilterDetails
		return ret
	}).(BlobFilterDetailsOutput)
}

func (o BlobFilterDetailsPtrOutput) BlobPathList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BlobFilterDetails) []string {
		if v == nil {
			return nil
		}
		return v.BlobPathList
	}).(pulumi.StringArrayOutput)
}

func (o BlobFilterDetailsPtrOutput) BlobPrefixList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BlobFilterDetails) []string {
		if v == nil {
			return nil
		}
		return v.BlobPrefixList
	}).(pulumi.StringArrayOutput)
}

func (o BlobFilterDetailsPtrOutput) ContainerList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BlobFilterDetails) []string {
		if v == nil {
			return nil
		}
		return v.ContainerList
	}).(pulumi.StringArrayOutput)
}

type BlobFilterDetailsResponse struct {
	BlobPathList   []string `pulumi:"blobPathList"`
	BlobPrefixList []string `pulumi:"blobPrefixList"`
	ContainerList  []string `pulumi:"containerList"`
}





type BlobFilterDetailsResponseInput interface {
	pulumi.Input

	ToBlobFilterDetailsResponseOutput() BlobFilterDetailsResponseOutput
	ToBlobFilterDetailsResponseOutputWithContext(context.Context) BlobFilterDetailsResponseOutput
}

type BlobFilterDetailsResponseArgs struct {
	BlobPathList   pulumi.StringArrayInput `pulumi:"blobPathList"`
	BlobPrefixList pulumi.StringArrayInput `pulumi:"blobPrefixList"`
	ContainerList  pulumi.StringArrayInput `pulumi:"containerList"`
}

func (BlobFilterDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobFilterDetailsResponse)(nil)).Elem()
}

func (i BlobFilterDetailsResponseArgs) ToBlobFilterDetailsResponseOutput() BlobFilterDetailsResponseOutput {
	return i.ToBlobFilterDetailsResponseOutputWithContext(context.Background())
}

func (i BlobFilterDetailsResponseArgs) ToBlobFilterDetailsResponseOutputWithContext(ctx context.Context) BlobFilterDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobFilterDetailsResponseOutput)
}

func (i BlobFilterDetailsResponseArgs) ToBlobFilterDetailsResponsePtrOutput() BlobFilterDetailsResponsePtrOutput {
	return i.ToBlobFilterDetailsResponsePtrOutputWithContext(context.Background())
}

func (i BlobFilterDetailsResponseArgs) ToBlobFilterDetailsResponsePtrOutputWithContext(ctx context.Context) BlobFilterDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobFilterDetailsResponseOutput).ToBlobFilterDetailsResponsePtrOutputWithContext(ctx)
}









type BlobFilterDetailsResponsePtrInput interface {
	pulumi.Input

	ToBlobFilterDetailsResponsePtrOutput() BlobFilterDetailsResponsePtrOutput
	ToBlobFilterDetailsResponsePtrOutputWithContext(context.Context) BlobFilterDetailsResponsePtrOutput
}

type blobFilterDetailsResponsePtrType BlobFilterDetailsResponseArgs

func BlobFilterDetailsResponsePtr(v *BlobFilterDetailsResponseArgs) BlobFilterDetailsResponsePtrInput {
	return (*blobFilterDetailsResponsePtrType)(v)
}

func (*blobFilterDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobFilterDetailsResponse)(nil)).Elem()
}

func (i *blobFilterDetailsResponsePtrType) ToBlobFilterDetailsResponsePtrOutput() BlobFilterDetailsResponsePtrOutput {
	return i.ToBlobFilterDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *blobFilterDetailsResponsePtrType) ToBlobFilterDetailsResponsePtrOutputWithContext(ctx context.Context) BlobFilterDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobFilterDetailsResponsePtrOutput)
}

type BlobFilterDetailsResponseOutput struct{ *pulumi.OutputState }

func (BlobFilterDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobFilterDetailsResponse)(nil)).Elem()
}

func (o BlobFilterDetailsResponseOutput) ToBlobFilterDetailsResponseOutput() BlobFilterDetailsResponseOutput {
	return o
}

func (o BlobFilterDetailsResponseOutput) ToBlobFilterDetailsResponseOutputWithContext(ctx context.Context) BlobFilterDetailsResponseOutput {
	return o
}

func (o BlobFilterDetailsResponseOutput) ToBlobFilterDetailsResponsePtrOutput() BlobFilterDetailsResponsePtrOutput {
	return o.ToBlobFilterDetailsResponsePtrOutputWithContext(context.Background())
}

func (o BlobFilterDetailsResponseOutput) ToBlobFilterDetailsResponsePtrOutputWithContext(ctx context.Context) BlobFilterDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobFilterDetailsResponse) *BlobFilterDetailsResponse {
		return &v
	}).(BlobFilterDetailsResponsePtrOutput)
}

func (o BlobFilterDetailsResponseOutput) BlobPathList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BlobFilterDetailsResponse) []string { return v.BlobPathList }).(pulumi.StringArrayOutput)
}

func (o BlobFilterDetailsResponseOutput) BlobPrefixList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BlobFilterDetailsResponse) []string { return v.BlobPrefixList }).(pulumi.StringArrayOutput)
}

func (o BlobFilterDetailsResponseOutput) ContainerList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BlobFilterDetailsResponse) []string { return v.ContainerList }).(pulumi.StringArrayOutput)
}

type BlobFilterDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (BlobFilterDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobFilterDetailsResponse)(nil)).Elem()
}

func (o BlobFilterDetailsResponsePtrOutput) ToBlobFilterDetailsResponsePtrOutput() BlobFilterDetailsResponsePtrOutput {
	return o
}

func (o BlobFilterDetailsResponsePtrOutput) ToBlobFilterDetailsResponsePtrOutputWithContext(ctx context.Context) BlobFilterDetailsResponsePtrOutput {
	return o
}

func (o BlobFilterDetailsResponsePtrOutput) Elem() BlobFilterDetailsResponseOutput {
	return o.ApplyT(func(v *BlobFilterDetailsResponse) BlobFilterDetailsResponse {
		if v != nil {
			return *v
		}
		var ret BlobFilterDetailsResponse
		return ret
	}).(BlobFilterDetailsResponseOutput)
}

func (o BlobFilterDetailsResponsePtrOutput) BlobPathList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BlobFilterDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.BlobPathList
	}).(pulumi.StringArrayOutput)
}

func (o BlobFilterDetailsResponsePtrOutput) BlobPrefixList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BlobFilterDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.BlobPrefixList
	}).(pulumi.StringArrayOutput)
}

func (o BlobFilterDetailsResponsePtrOutput) ContainerList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BlobFilterDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.ContainerList
	}).(pulumi.StringArrayOutput)
}

type CloudErrorResponse struct {
	AdditionalInfo []AdditionalErrorInfoResponse `pulumi:"additionalInfo"`
	Code           *string                       `pulumi:"code"`
	Details        []CloudErrorResponse          `pulumi:"details"`
	Message        *string                       `pulumi:"message"`
	Target         *string                       `pulumi:"target"`
}





type CloudErrorResponseInput interface {
	pulumi.Input

	ToCloudErrorResponseOutput() CloudErrorResponseOutput
	ToCloudErrorResponseOutputWithContext(context.Context) CloudErrorResponseOutput
}

type CloudErrorResponseArgs struct {
	AdditionalInfo AdditionalErrorInfoResponseArrayInput `pulumi:"additionalInfo"`
	Code           pulumi.StringPtrInput                 `pulumi:"code"`
	Details        CloudErrorResponseArrayInput          `pulumi:"details"`
	Message        pulumi.StringPtrInput                 `pulumi:"message"`
	Target         pulumi.StringPtrInput                 `pulumi:"target"`
}

func (CloudErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudErrorResponse)(nil)).Elem()
}

func (i CloudErrorResponseArgs) ToCloudErrorResponseOutput() CloudErrorResponseOutput {
	return i.ToCloudErrorResponseOutputWithContext(context.Background())
}

func (i CloudErrorResponseArgs) ToCloudErrorResponseOutputWithContext(ctx context.Context) CloudErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorResponseOutput)
}

func (i CloudErrorResponseArgs) ToCloudErrorResponsePtrOutput() CloudErrorResponsePtrOutput {
	return i.ToCloudErrorResponsePtrOutputWithContext(context.Background())
}

func (i CloudErrorResponseArgs) ToCloudErrorResponsePtrOutputWithContext(ctx context.Context) CloudErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorResponseOutput).ToCloudErrorResponsePtrOutputWithContext(ctx)
}









type CloudErrorResponsePtrInput interface {
	pulumi.Input

	ToCloudErrorResponsePtrOutput() CloudErrorResponsePtrOutput
	ToCloudErrorResponsePtrOutputWithContext(context.Context) CloudErrorResponsePtrOutput
}

type cloudErrorResponsePtrType CloudErrorResponseArgs

func CloudErrorResponsePtr(v *CloudErrorResponseArgs) CloudErrorResponsePtrInput {
	return (*cloudErrorResponsePtrType)(v)
}

func (*cloudErrorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudErrorResponse)(nil)).Elem()
}

func (i *cloudErrorResponsePtrType) ToCloudErrorResponsePtrOutput() CloudErrorResponsePtrOutput {
	return i.ToCloudErrorResponsePtrOutputWithContext(context.Background())
}

func (i *cloudErrorResponsePtrType) ToCloudErrorResponsePtrOutputWithContext(ctx context.Context) CloudErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorResponsePtrOutput)
}





type CloudErrorResponseArrayInput interface {
	pulumi.Input

	ToCloudErrorResponseArrayOutput() CloudErrorResponseArrayOutput
	ToCloudErrorResponseArrayOutputWithContext(context.Context) CloudErrorResponseArrayOutput
}

type CloudErrorResponseArray []CloudErrorResponseInput

func (CloudErrorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudErrorResponse)(nil)).Elem()
}

func (i CloudErrorResponseArray) ToCloudErrorResponseArrayOutput() CloudErrorResponseArrayOutput {
	return i.ToCloudErrorResponseArrayOutputWithContext(context.Background())
}

func (i CloudErrorResponseArray) ToCloudErrorResponseArrayOutputWithContext(ctx context.Context) CloudErrorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudErrorResponseArrayOutput)
}

type CloudErrorResponseOutput struct{ *pulumi.OutputState }

func (CloudErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudErrorResponse)(nil)).Elem()
}

func (o CloudErrorResponseOutput) ToCloudErrorResponseOutput() CloudErrorResponseOutput {
	return o
}

func (o CloudErrorResponseOutput) ToCloudErrorResponseOutputWithContext(ctx context.Context) CloudErrorResponseOutput {
	return o
}

func (o CloudErrorResponseOutput) ToCloudErrorResponsePtrOutput() CloudErrorResponsePtrOutput {
	return o.ToCloudErrorResponsePtrOutputWithContext(context.Background())
}

func (o CloudErrorResponseOutput) ToCloudErrorResponsePtrOutputWithContext(ctx context.Context) CloudErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudErrorResponse) *CloudErrorResponse {
		return &v
	}).(CloudErrorResponsePtrOutput)
}

func (o CloudErrorResponseOutput) AdditionalInfo() AdditionalErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v CloudErrorResponse) []AdditionalErrorInfoResponse { return v.AdditionalInfo }).(AdditionalErrorInfoResponseArrayOutput)
}

func (o CloudErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o CloudErrorResponseOutput) Details() CloudErrorResponseArrayOutput {
	return o.ApplyT(func(v CloudErrorResponse) []CloudErrorResponse { return v.Details }).(CloudErrorResponseArrayOutput)
}

func (o CloudErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o CloudErrorResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudErrorResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type CloudErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudErrorResponse)(nil)).Elem()
}

func (o CloudErrorResponsePtrOutput) ToCloudErrorResponsePtrOutput() CloudErrorResponsePtrOutput {
	return o
}

func (o CloudErrorResponsePtrOutput) ToCloudErrorResponsePtrOutputWithContext(ctx context.Context) CloudErrorResponsePtrOutput {
	return o
}

func (o CloudErrorResponsePtrOutput) Elem() CloudErrorResponseOutput {
	return o.ApplyT(func(v *CloudErrorResponse) CloudErrorResponse {
		if v != nil {
			return *v
		}
		var ret CloudErrorResponse
		return ret
	}).(CloudErrorResponseOutput)
}

func (o CloudErrorResponsePtrOutput) AdditionalInfo() AdditionalErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v *CloudErrorResponse) []AdditionalErrorInfoResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalInfo
	}).(AdditionalErrorInfoResponseArrayOutput)
}

func (o CloudErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o CloudErrorResponsePtrOutput) Details() CloudErrorResponseArrayOutput {
	return o.ApplyT(func(v *CloudErrorResponse) []CloudErrorResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(CloudErrorResponseArrayOutput)
}

func (o CloudErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o CloudErrorResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type CloudErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (CloudErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudErrorResponse)(nil)).Elem()
}

func (o CloudErrorResponseArrayOutput) ToCloudErrorResponseArrayOutput() CloudErrorResponseArrayOutput {
	return o
}

func (o CloudErrorResponseArrayOutput) ToCloudErrorResponseArrayOutputWithContext(ctx context.Context) CloudErrorResponseArrayOutput {
	return o
}

func (o CloudErrorResponseArrayOutput) Index(i pulumi.IntInput) CloudErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudErrorResponse {
		return vs[0].([]CloudErrorResponse)[vs[1].(int)]
	}).(CloudErrorResponseOutput)
}

type ContactDetails struct {
	ContactName            string                   `pulumi:"contactName"`
	EmailList              []string                 `pulumi:"emailList"`
	Mobile                 *string                  `pulumi:"mobile"`
	NotificationPreference []NotificationPreference `pulumi:"notificationPreference"`
	Phone                  string                   `pulumi:"phone"`
	PhoneExtension         *string                  `pulumi:"phoneExtension"`
}





type ContactDetailsInput interface {
	pulumi.Input

	ToContactDetailsOutput() ContactDetailsOutput
	ToContactDetailsOutputWithContext(context.Context) ContactDetailsOutput
}

type ContactDetailsArgs struct {
	ContactName            pulumi.StringInput               `pulumi:"contactName"`
	EmailList              pulumi.StringArrayInput          `pulumi:"emailList"`
	Mobile                 pulumi.StringPtrInput            `pulumi:"mobile"`
	NotificationPreference NotificationPreferenceArrayInput `pulumi:"notificationPreference"`
	Phone                  pulumi.StringInput               `pulumi:"phone"`
	PhoneExtension         pulumi.StringPtrInput            `pulumi:"phoneExtension"`
}

func (ContactDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetails)(nil)).Elem()
}

func (i ContactDetailsArgs) ToContactDetailsOutput() ContactDetailsOutput {
	return i.ToContactDetailsOutputWithContext(context.Background())
}

func (i ContactDetailsArgs) ToContactDetailsOutputWithContext(ctx context.Context) ContactDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsOutput)
}

type ContactDetailsOutput struct{ *pulumi.OutputState }

func (ContactDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetails)(nil)).Elem()
}

func (o ContactDetailsOutput) ToContactDetailsOutput() ContactDetailsOutput {
	return o
}

func (o ContactDetailsOutput) ToContactDetailsOutputWithContext(ctx context.Context) ContactDetailsOutput {
	return o
}

func (o ContactDetailsOutput) ContactName() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetails) string { return v.ContactName }).(pulumi.StringOutput)
}

func (o ContactDetailsOutput) EmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContactDetails) []string { return v.EmailList }).(pulumi.StringArrayOutput)
}

func (o ContactDetailsOutput) Mobile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetails) *string { return v.Mobile }).(pulumi.StringPtrOutput)
}

func (o ContactDetailsOutput) NotificationPreference() NotificationPreferenceArrayOutput {
	return o.ApplyT(func(v ContactDetails) []NotificationPreference { return v.NotificationPreference }).(NotificationPreferenceArrayOutput)
}

func (o ContactDetailsOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetails) string { return v.Phone }).(pulumi.StringOutput)
}

func (o ContactDetailsOutput) PhoneExtension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetails) *string { return v.PhoneExtension }).(pulumi.StringPtrOutput)
}

type ContactDetailsResponse struct {
	ContactName            string                           `pulumi:"contactName"`
	EmailList              []string                         `pulumi:"emailList"`
	Mobile                 *string                          `pulumi:"mobile"`
	NotificationPreference []NotificationPreferenceResponse `pulumi:"notificationPreference"`
	Phone                  string                           `pulumi:"phone"`
	PhoneExtension         *string                          `pulumi:"phoneExtension"`
}





type ContactDetailsResponseInput interface {
	pulumi.Input

	ToContactDetailsResponseOutput() ContactDetailsResponseOutput
	ToContactDetailsResponseOutputWithContext(context.Context) ContactDetailsResponseOutput
}

type ContactDetailsResponseArgs struct {
	ContactName            pulumi.StringInput                       `pulumi:"contactName"`
	EmailList              pulumi.StringArrayInput                  `pulumi:"emailList"`
	Mobile                 pulumi.StringPtrInput                    `pulumi:"mobile"`
	NotificationPreference NotificationPreferenceResponseArrayInput `pulumi:"notificationPreference"`
	Phone                  pulumi.StringInput                       `pulumi:"phone"`
	PhoneExtension         pulumi.StringPtrInput                    `pulumi:"phoneExtension"`
}

func (ContactDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetailsResponse)(nil)).Elem()
}

func (i ContactDetailsResponseArgs) ToContactDetailsResponseOutput() ContactDetailsResponseOutput {
	return i.ToContactDetailsResponseOutputWithContext(context.Background())
}

func (i ContactDetailsResponseArgs) ToContactDetailsResponseOutputWithContext(ctx context.Context) ContactDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactDetailsResponseOutput)
}

type ContactDetailsResponseOutput struct{ *pulumi.OutputState }

func (ContactDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactDetailsResponse)(nil)).Elem()
}

func (o ContactDetailsResponseOutput) ToContactDetailsResponseOutput() ContactDetailsResponseOutput {
	return o
}

func (o ContactDetailsResponseOutput) ToContactDetailsResponseOutputWithContext(ctx context.Context) ContactDetailsResponseOutput {
	return o
}

func (o ContactDetailsResponseOutput) ContactName() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetailsResponse) string { return v.ContactName }).(pulumi.StringOutput)
}

func (o ContactDetailsResponseOutput) EmailList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContactDetailsResponse) []string { return v.EmailList }).(pulumi.StringArrayOutput)
}

func (o ContactDetailsResponseOutput) Mobile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetailsResponse) *string { return v.Mobile }).(pulumi.StringPtrOutput)
}

func (o ContactDetailsResponseOutput) NotificationPreference() NotificationPreferenceResponseArrayOutput {
	return o.ApplyT(func(v ContactDetailsResponse) []NotificationPreferenceResponse { return v.NotificationPreference }).(NotificationPreferenceResponseArrayOutput)
}

func (o ContactDetailsResponseOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ContactDetailsResponse) string { return v.Phone }).(pulumi.StringOutput)
}

func (o ContactDetailsResponseOutput) PhoneExtension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactDetailsResponse) *string { return v.PhoneExtension }).(pulumi.StringPtrOutput)
}

type CopyProgressResponse struct {
	AccountId                   string  `pulumi:"accountId"`
	BytesProcessed              float64 `pulumi:"bytesProcessed"`
	DataAccountType             string  `pulumi:"dataAccountType"`
	DirectoriesErroredOut       float64 `pulumi:"directoriesErroredOut"`
	FilesErroredOut             float64 `pulumi:"filesErroredOut"`
	FilesProcessed              float64 `pulumi:"filesProcessed"`
	InvalidDirectoriesProcessed float64 `pulumi:"invalidDirectoriesProcessed"`
	InvalidFileBytesUploaded    float64 `pulumi:"invalidFileBytesUploaded"`
	InvalidFilesProcessed       float64 `pulumi:"invalidFilesProcessed"`
	IsEnumerationInProgress     bool    `pulumi:"isEnumerationInProgress"`
	RenamedContainerCount       float64 `pulumi:"renamedContainerCount"`
	StorageAccountName          string  `pulumi:"storageAccountName"`
	TotalBytesToProcess         float64 `pulumi:"totalBytesToProcess"`
	TotalFilesToProcess         float64 `pulumi:"totalFilesToProcess"`
	TransferType                string  `pulumi:"transferType"`
}





type CopyProgressResponseInput interface {
	pulumi.Input

	ToCopyProgressResponseOutput() CopyProgressResponseOutput
	ToCopyProgressResponseOutputWithContext(context.Context) CopyProgressResponseOutput
}

type CopyProgressResponseArgs struct {
	AccountId                   pulumi.StringInput  `pulumi:"accountId"`
	BytesProcessed              pulumi.Float64Input `pulumi:"bytesProcessed"`
	DataAccountType             pulumi.StringInput  `pulumi:"dataAccountType"`
	DirectoriesErroredOut       pulumi.Float64Input `pulumi:"directoriesErroredOut"`
	FilesErroredOut             pulumi.Float64Input `pulumi:"filesErroredOut"`
	FilesProcessed              pulumi.Float64Input `pulumi:"filesProcessed"`
	InvalidDirectoriesProcessed pulumi.Float64Input `pulumi:"invalidDirectoriesProcessed"`
	InvalidFileBytesUploaded    pulumi.Float64Input `pulumi:"invalidFileBytesUploaded"`
	InvalidFilesProcessed       pulumi.Float64Input `pulumi:"invalidFilesProcessed"`
	IsEnumerationInProgress     pulumi.BoolInput    `pulumi:"isEnumerationInProgress"`
	RenamedContainerCount       pulumi.Float64Input `pulumi:"renamedContainerCount"`
	StorageAccountName          pulumi.StringInput  `pulumi:"storageAccountName"`
	TotalBytesToProcess         pulumi.Float64Input `pulumi:"totalBytesToProcess"`
	TotalFilesToProcess         pulumi.Float64Input `pulumi:"totalFilesToProcess"`
	TransferType                pulumi.StringInput  `pulumi:"transferType"`
}

func (CopyProgressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyProgressResponse)(nil)).Elem()
}

func (i CopyProgressResponseArgs) ToCopyProgressResponseOutput() CopyProgressResponseOutput {
	return i.ToCopyProgressResponseOutputWithContext(context.Background())
}

func (i CopyProgressResponseArgs) ToCopyProgressResponseOutputWithContext(ctx context.Context) CopyProgressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyProgressResponseOutput)
}





type CopyProgressResponseArrayInput interface {
	pulumi.Input

	ToCopyProgressResponseArrayOutput() CopyProgressResponseArrayOutput
	ToCopyProgressResponseArrayOutputWithContext(context.Context) CopyProgressResponseArrayOutput
}

type CopyProgressResponseArray []CopyProgressResponseInput

func (CopyProgressResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CopyProgressResponse)(nil)).Elem()
}

func (i CopyProgressResponseArray) ToCopyProgressResponseArrayOutput() CopyProgressResponseArrayOutput {
	return i.ToCopyProgressResponseArrayOutputWithContext(context.Background())
}

func (i CopyProgressResponseArray) ToCopyProgressResponseArrayOutputWithContext(ctx context.Context) CopyProgressResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyProgressResponseArrayOutput)
}

type CopyProgressResponseOutput struct{ *pulumi.OutputState }

func (CopyProgressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyProgressResponse)(nil)).Elem()
}

func (o CopyProgressResponseOutput) ToCopyProgressResponseOutput() CopyProgressResponseOutput {
	return o
}

func (o CopyProgressResponseOutput) ToCopyProgressResponseOutputWithContext(ctx context.Context) CopyProgressResponseOutput {
	return o
}

func (o CopyProgressResponseOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v CopyProgressResponse) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o CopyProgressResponseOutput) BytesProcessed() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.BytesProcessed }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) DataAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v CopyProgressResponse) string { return v.DataAccountType }).(pulumi.StringOutput)
}

func (o CopyProgressResponseOutput) DirectoriesErroredOut() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.DirectoriesErroredOut }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) FilesErroredOut() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.FilesErroredOut }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) FilesProcessed() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.FilesProcessed }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) InvalidDirectoriesProcessed() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.InvalidDirectoriesProcessed }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) InvalidFileBytesUploaded() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.InvalidFileBytesUploaded }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) InvalidFilesProcessed() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.InvalidFilesProcessed }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) IsEnumerationInProgress() pulumi.BoolOutput {
	return o.ApplyT(func(v CopyProgressResponse) bool { return v.IsEnumerationInProgress }).(pulumi.BoolOutput)
}

func (o CopyProgressResponseOutput) RenamedContainerCount() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.RenamedContainerCount }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v CopyProgressResponse) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o CopyProgressResponseOutput) TotalBytesToProcess() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.TotalBytesToProcess }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) TotalFilesToProcess() pulumi.Float64Output {
	return o.ApplyT(func(v CopyProgressResponse) float64 { return v.TotalFilesToProcess }).(pulumi.Float64Output)
}

func (o CopyProgressResponseOutput) TransferType() pulumi.StringOutput {
	return o.ApplyT(func(v CopyProgressResponse) string { return v.TransferType }).(pulumi.StringOutput)
}

type CopyProgressResponseArrayOutput struct{ *pulumi.OutputState }

func (CopyProgressResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CopyProgressResponse)(nil)).Elem()
}

func (o CopyProgressResponseArrayOutput) ToCopyProgressResponseArrayOutput() CopyProgressResponseArrayOutput {
	return o
}

func (o CopyProgressResponseArrayOutput) ToCopyProgressResponseArrayOutputWithContext(ctx context.Context) CopyProgressResponseArrayOutput {
	return o
}

func (o CopyProgressResponseArrayOutput) Index(i pulumi.IntInput) CopyProgressResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CopyProgressResponse {
		return vs[0].([]CopyProgressResponse)[vs[1].(int)]
	}).(CopyProgressResponseOutput)
}

type CustomerDiskJobSecretsResponse struct {
	CarrierAccountNumber string                       `pulumi:"carrierAccountNumber"`
	DcAccessSecurityCode DcAccessSecurityCodeResponse `pulumi:"dcAccessSecurityCode"`
	DiskSecrets          []DiskSecretResponse         `pulumi:"diskSecrets"`
	Error                CloudErrorResponse           `pulumi:"error"`
	JobSecretsType       string                       `pulumi:"jobSecretsType"`
}





type CustomerDiskJobSecretsResponseInput interface {
	pulumi.Input

	ToCustomerDiskJobSecretsResponseOutput() CustomerDiskJobSecretsResponseOutput
	ToCustomerDiskJobSecretsResponseOutputWithContext(context.Context) CustomerDiskJobSecretsResponseOutput
}

type CustomerDiskJobSecretsResponseArgs struct {
	CarrierAccountNumber pulumi.StringInput                `pulumi:"carrierAccountNumber"`
	DcAccessSecurityCode DcAccessSecurityCodeResponseInput `pulumi:"dcAccessSecurityCode"`
	DiskSecrets          DiskSecretResponseArrayInput      `pulumi:"diskSecrets"`
	Error                CloudErrorResponseInput           `pulumi:"error"`
	JobSecretsType       pulumi.StringInput                `pulumi:"jobSecretsType"`
}

func (CustomerDiskJobSecretsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerDiskJobSecretsResponse)(nil)).Elem()
}

func (i CustomerDiskJobSecretsResponseArgs) ToCustomerDiskJobSecretsResponseOutput() CustomerDiskJobSecretsResponseOutput {
	return i.ToCustomerDiskJobSecretsResponseOutputWithContext(context.Background())
}

func (i CustomerDiskJobSecretsResponseArgs) ToCustomerDiskJobSecretsResponseOutputWithContext(ctx context.Context) CustomerDiskJobSecretsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerDiskJobSecretsResponseOutput)
}

type CustomerDiskJobSecretsResponseOutput struct{ *pulumi.OutputState }

func (CustomerDiskJobSecretsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerDiskJobSecretsResponse)(nil)).Elem()
}

func (o CustomerDiskJobSecretsResponseOutput) ToCustomerDiskJobSecretsResponseOutput() CustomerDiskJobSecretsResponseOutput {
	return o
}

func (o CustomerDiskJobSecretsResponseOutput) ToCustomerDiskJobSecretsResponseOutputWithContext(ctx context.Context) CustomerDiskJobSecretsResponseOutput {
	return o
}

func (o CustomerDiskJobSecretsResponseOutput) CarrierAccountNumber() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerDiskJobSecretsResponse) string { return v.CarrierAccountNumber }).(pulumi.StringOutput)
}

func (o CustomerDiskJobSecretsResponseOutput) DcAccessSecurityCode() DcAccessSecurityCodeResponseOutput {
	return o.ApplyT(func(v CustomerDiskJobSecretsResponse) DcAccessSecurityCodeResponse { return v.DcAccessSecurityCode }).(DcAccessSecurityCodeResponseOutput)
}

func (o CustomerDiskJobSecretsResponseOutput) DiskSecrets() DiskSecretResponseArrayOutput {
	return o.ApplyT(func(v CustomerDiskJobSecretsResponse) []DiskSecretResponse { return v.DiskSecrets }).(DiskSecretResponseArrayOutput)
}

func (o CustomerDiskJobSecretsResponseOutput) Error() CloudErrorResponseOutput {
	return o.ApplyT(func(v CustomerDiskJobSecretsResponse) CloudErrorResponse { return v.Error }).(CloudErrorResponseOutput)
}

func (o CustomerDiskJobSecretsResponseOutput) JobSecretsType() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerDiskJobSecretsResponse) string { return v.JobSecretsType }).(pulumi.StringOutput)
}

type DataBoxAccountCopyLogDetailsResponse struct {
	AccountName        string `pulumi:"accountName"`
	CopyLogDetailsType string `pulumi:"copyLogDetailsType"`
	CopyLogLink        string `pulumi:"copyLogLink"`
	CopyVerboseLogLink string `pulumi:"copyVerboseLogLink"`
}





type DataBoxAccountCopyLogDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxAccountCopyLogDetailsResponseOutput() DataBoxAccountCopyLogDetailsResponseOutput
	ToDataBoxAccountCopyLogDetailsResponseOutputWithContext(context.Context) DataBoxAccountCopyLogDetailsResponseOutput
}

type DataBoxAccountCopyLogDetailsResponseArgs struct {
	AccountName        pulumi.StringInput `pulumi:"accountName"`
	CopyLogDetailsType pulumi.StringInput `pulumi:"copyLogDetailsType"`
	CopyLogLink        pulumi.StringInput `pulumi:"copyLogLink"`
	CopyVerboseLogLink pulumi.StringInput `pulumi:"copyVerboseLogLink"`
}

func (DataBoxAccountCopyLogDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxAccountCopyLogDetailsResponse)(nil)).Elem()
}

func (i DataBoxAccountCopyLogDetailsResponseArgs) ToDataBoxAccountCopyLogDetailsResponseOutput() DataBoxAccountCopyLogDetailsResponseOutput {
	return i.ToDataBoxAccountCopyLogDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxAccountCopyLogDetailsResponseArgs) ToDataBoxAccountCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxAccountCopyLogDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxAccountCopyLogDetailsResponseOutput)
}

type DataBoxAccountCopyLogDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxAccountCopyLogDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxAccountCopyLogDetailsResponse)(nil)).Elem()
}

func (o DataBoxAccountCopyLogDetailsResponseOutput) ToDataBoxAccountCopyLogDetailsResponseOutput() DataBoxAccountCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxAccountCopyLogDetailsResponseOutput) ToDataBoxAccountCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxAccountCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxAccountCopyLogDetailsResponseOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxAccountCopyLogDetailsResponse) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o DataBoxAccountCopyLogDetailsResponseOutput) CopyLogDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxAccountCopyLogDetailsResponse) string { return v.CopyLogDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxAccountCopyLogDetailsResponseOutput) CopyLogLink() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxAccountCopyLogDetailsResponse) string { return v.CopyLogLink }).(pulumi.StringOutput)
}

func (o DataBoxAccountCopyLogDetailsResponseOutput) CopyVerboseLogLink() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxAccountCopyLogDetailsResponse) string { return v.CopyVerboseLogLink }).(pulumi.StringOutput)
}

type DataBoxCustomerDiskCopyLogDetailsResponse struct {
	CopyLogDetailsType string `pulumi:"copyLogDetailsType"`
	ErrorLogLink       string `pulumi:"errorLogLink"`
	SerialNumber       string `pulumi:"serialNumber"`
	VerboseLogLink     string `pulumi:"verboseLogLink"`
}





type DataBoxCustomerDiskCopyLogDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxCustomerDiskCopyLogDetailsResponseOutput() DataBoxCustomerDiskCopyLogDetailsResponseOutput
	ToDataBoxCustomerDiskCopyLogDetailsResponseOutputWithContext(context.Context) DataBoxCustomerDiskCopyLogDetailsResponseOutput
}

type DataBoxCustomerDiskCopyLogDetailsResponseArgs struct {
	CopyLogDetailsType pulumi.StringInput `pulumi:"copyLogDetailsType"`
	ErrorLogLink       pulumi.StringInput `pulumi:"errorLogLink"`
	SerialNumber       pulumi.StringInput `pulumi:"serialNumber"`
	VerboseLogLink     pulumi.StringInput `pulumi:"verboseLogLink"`
}

func (DataBoxCustomerDiskCopyLogDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxCustomerDiskCopyLogDetailsResponse)(nil)).Elem()
}

func (i DataBoxCustomerDiskCopyLogDetailsResponseArgs) ToDataBoxCustomerDiskCopyLogDetailsResponseOutput() DataBoxCustomerDiskCopyLogDetailsResponseOutput {
	return i.ToDataBoxCustomerDiskCopyLogDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxCustomerDiskCopyLogDetailsResponseArgs) ToDataBoxCustomerDiskCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxCustomerDiskCopyLogDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxCustomerDiskCopyLogDetailsResponseOutput)
}

type DataBoxCustomerDiskCopyLogDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxCustomerDiskCopyLogDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxCustomerDiskCopyLogDetailsResponse)(nil)).Elem()
}

func (o DataBoxCustomerDiskCopyLogDetailsResponseOutput) ToDataBoxCustomerDiskCopyLogDetailsResponseOutput() DataBoxCustomerDiskCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxCustomerDiskCopyLogDetailsResponseOutput) ToDataBoxCustomerDiskCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxCustomerDiskCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxCustomerDiskCopyLogDetailsResponseOutput) CopyLogDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyLogDetailsResponse) string { return v.CopyLogDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskCopyLogDetailsResponseOutput) ErrorLogLink() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyLogDetailsResponse) string { return v.ErrorLogLink }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskCopyLogDetailsResponseOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyLogDetailsResponse) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskCopyLogDetailsResponseOutput) VerboseLogLink() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyLogDetailsResponse) string { return v.VerboseLogLink }).(pulumi.StringOutput)
}

type DataBoxCustomerDiskCopyProgressResponse struct {
	AccountId                   string  `pulumi:"accountId"`
	BytesProcessed              float64 `pulumi:"bytesProcessed"`
	CopyStatus                  string  `pulumi:"copyStatus"`
	DataAccountType             string  `pulumi:"dataAccountType"`
	DirectoriesErroredOut       float64 `pulumi:"directoriesErroredOut"`
	FilesErroredOut             float64 `pulumi:"filesErroredOut"`
	FilesProcessed              float64 `pulumi:"filesProcessed"`
	InvalidDirectoriesProcessed float64 `pulumi:"invalidDirectoriesProcessed"`
	InvalidFileBytesUploaded    float64 `pulumi:"invalidFileBytesUploaded"`
	InvalidFilesProcessed       float64 `pulumi:"invalidFilesProcessed"`
	IsEnumerationInProgress     bool    `pulumi:"isEnumerationInProgress"`
	RenamedContainerCount       float64 `pulumi:"renamedContainerCount"`
	SerialNumber                string  `pulumi:"serialNumber"`
	StorageAccountName          string  `pulumi:"storageAccountName"`
	TotalBytesToProcess         float64 `pulumi:"totalBytesToProcess"`
	TotalFilesToProcess         float64 `pulumi:"totalFilesToProcess"`
	TransferType                string  `pulumi:"transferType"`
}





type DataBoxCustomerDiskCopyProgressResponseInput interface {
	pulumi.Input

	ToDataBoxCustomerDiskCopyProgressResponseOutput() DataBoxCustomerDiskCopyProgressResponseOutput
	ToDataBoxCustomerDiskCopyProgressResponseOutputWithContext(context.Context) DataBoxCustomerDiskCopyProgressResponseOutput
}

type DataBoxCustomerDiskCopyProgressResponseArgs struct {
	AccountId                   pulumi.StringInput  `pulumi:"accountId"`
	BytesProcessed              pulumi.Float64Input `pulumi:"bytesProcessed"`
	CopyStatus                  pulumi.StringInput  `pulumi:"copyStatus"`
	DataAccountType             pulumi.StringInput  `pulumi:"dataAccountType"`
	DirectoriesErroredOut       pulumi.Float64Input `pulumi:"directoriesErroredOut"`
	FilesErroredOut             pulumi.Float64Input `pulumi:"filesErroredOut"`
	FilesProcessed              pulumi.Float64Input `pulumi:"filesProcessed"`
	InvalidDirectoriesProcessed pulumi.Float64Input `pulumi:"invalidDirectoriesProcessed"`
	InvalidFileBytesUploaded    pulumi.Float64Input `pulumi:"invalidFileBytesUploaded"`
	InvalidFilesProcessed       pulumi.Float64Input `pulumi:"invalidFilesProcessed"`
	IsEnumerationInProgress     pulumi.BoolInput    `pulumi:"isEnumerationInProgress"`
	RenamedContainerCount       pulumi.Float64Input `pulumi:"renamedContainerCount"`
	SerialNumber                pulumi.StringInput  `pulumi:"serialNumber"`
	StorageAccountName          pulumi.StringInput  `pulumi:"storageAccountName"`
	TotalBytesToProcess         pulumi.Float64Input `pulumi:"totalBytesToProcess"`
	TotalFilesToProcess         pulumi.Float64Input `pulumi:"totalFilesToProcess"`
	TransferType                pulumi.StringInput  `pulumi:"transferType"`
}

func (DataBoxCustomerDiskCopyProgressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxCustomerDiskCopyProgressResponse)(nil)).Elem()
}

func (i DataBoxCustomerDiskCopyProgressResponseArgs) ToDataBoxCustomerDiskCopyProgressResponseOutput() DataBoxCustomerDiskCopyProgressResponseOutput {
	return i.ToDataBoxCustomerDiskCopyProgressResponseOutputWithContext(context.Background())
}

func (i DataBoxCustomerDiskCopyProgressResponseArgs) ToDataBoxCustomerDiskCopyProgressResponseOutputWithContext(ctx context.Context) DataBoxCustomerDiskCopyProgressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxCustomerDiskCopyProgressResponseOutput)
}





type DataBoxCustomerDiskCopyProgressResponseArrayInput interface {
	pulumi.Input

	ToDataBoxCustomerDiskCopyProgressResponseArrayOutput() DataBoxCustomerDiskCopyProgressResponseArrayOutput
	ToDataBoxCustomerDiskCopyProgressResponseArrayOutputWithContext(context.Context) DataBoxCustomerDiskCopyProgressResponseArrayOutput
}

type DataBoxCustomerDiskCopyProgressResponseArray []DataBoxCustomerDiskCopyProgressResponseInput

func (DataBoxCustomerDiskCopyProgressResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxCustomerDiskCopyProgressResponse)(nil)).Elem()
}

func (i DataBoxCustomerDiskCopyProgressResponseArray) ToDataBoxCustomerDiskCopyProgressResponseArrayOutput() DataBoxCustomerDiskCopyProgressResponseArrayOutput {
	return i.ToDataBoxCustomerDiskCopyProgressResponseArrayOutputWithContext(context.Background())
}

func (i DataBoxCustomerDiskCopyProgressResponseArray) ToDataBoxCustomerDiskCopyProgressResponseArrayOutputWithContext(ctx context.Context) DataBoxCustomerDiskCopyProgressResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxCustomerDiskCopyProgressResponseArrayOutput)
}

type DataBoxCustomerDiskCopyProgressResponseOutput struct{ *pulumi.OutputState }

func (DataBoxCustomerDiskCopyProgressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxCustomerDiskCopyProgressResponse)(nil)).Elem()
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) ToDataBoxCustomerDiskCopyProgressResponseOutput() DataBoxCustomerDiskCopyProgressResponseOutput {
	return o
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) ToDataBoxCustomerDiskCopyProgressResponseOutputWithContext(ctx context.Context) DataBoxCustomerDiskCopyProgressResponseOutput {
	return o
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) BytesProcessed() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) float64 { return v.BytesProcessed }).(pulumi.Float64Output)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) CopyStatus() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) string { return v.CopyStatus }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) DataAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) string { return v.DataAccountType }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) DirectoriesErroredOut() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) float64 { return v.DirectoriesErroredOut }).(pulumi.Float64Output)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) FilesErroredOut() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) float64 { return v.FilesErroredOut }).(pulumi.Float64Output)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) FilesProcessed() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) float64 { return v.FilesProcessed }).(pulumi.Float64Output)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) InvalidDirectoriesProcessed() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) float64 { return v.InvalidDirectoriesProcessed }).(pulumi.Float64Output)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) InvalidFileBytesUploaded() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) float64 { return v.InvalidFileBytesUploaded }).(pulumi.Float64Output)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) InvalidFilesProcessed() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) float64 { return v.InvalidFilesProcessed }).(pulumi.Float64Output)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) IsEnumerationInProgress() pulumi.BoolOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) bool { return v.IsEnumerationInProgress }).(pulumi.BoolOutput)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) RenamedContainerCount() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) float64 { return v.RenamedContainerCount }).(pulumi.Float64Output)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) TotalBytesToProcess() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) float64 { return v.TotalBytesToProcess }).(pulumi.Float64Output)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) TotalFilesToProcess() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) float64 { return v.TotalFilesToProcess }).(pulumi.Float64Output)
}

func (o DataBoxCustomerDiskCopyProgressResponseOutput) TransferType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskCopyProgressResponse) string { return v.TransferType }).(pulumi.StringOutput)
}

type DataBoxCustomerDiskCopyProgressResponseArrayOutput struct{ *pulumi.OutputState }

func (DataBoxCustomerDiskCopyProgressResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxCustomerDiskCopyProgressResponse)(nil)).Elem()
}

func (o DataBoxCustomerDiskCopyProgressResponseArrayOutput) ToDataBoxCustomerDiskCopyProgressResponseArrayOutput() DataBoxCustomerDiskCopyProgressResponseArrayOutput {
	return o
}

func (o DataBoxCustomerDiskCopyProgressResponseArrayOutput) ToDataBoxCustomerDiskCopyProgressResponseArrayOutputWithContext(ctx context.Context) DataBoxCustomerDiskCopyProgressResponseArrayOutput {
	return o
}

func (o DataBoxCustomerDiskCopyProgressResponseArrayOutput) Index(i pulumi.IntInput) DataBoxCustomerDiskCopyProgressResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataBoxCustomerDiskCopyProgressResponse {
		return vs[0].([]DataBoxCustomerDiskCopyProgressResponse)[vs[1].(int)]
	}).(DataBoxCustomerDiskCopyProgressResponseOutput)
}

type DataBoxCustomerDiskJobDetails struct {
	ContactDetails                 ContactDetails               `pulumi:"contactDetails"`
	DataExportDetails              []DataExportDetails          `pulumi:"dataExportDetails"`
	DataImportDetails              []DataImportDetails          `pulumi:"dataImportDetails"`
	EnableManifestBackup           *bool                        `pulumi:"enableManifestBackup"`
	ExpectedDataSizeInTeraBytes    *int                         `pulumi:"expectedDataSizeInTeraBytes"`
	ImportDiskDetailsCollection    map[string]ImportDiskDetails `pulumi:"importDiskDetailsCollection"`
	JobDetailsType                 string                       `pulumi:"jobDetailsType"`
	KeyEncryptionKey               *KeyEncryptionKey            `pulumi:"keyEncryptionKey"`
	Preferences                    *Preferences                 `pulumi:"preferences"`
	ReturnToCustomerPackageDetails PackageCarrierDetails        `pulumi:"returnToCustomerPackageDetails"`
	ShippingAddress                *ShippingAddress             `pulumi:"shippingAddress"`
}





type DataBoxCustomerDiskJobDetailsInput interface {
	pulumi.Input

	ToDataBoxCustomerDiskJobDetailsOutput() DataBoxCustomerDiskJobDetailsOutput
	ToDataBoxCustomerDiskJobDetailsOutputWithContext(context.Context) DataBoxCustomerDiskJobDetailsOutput
}

type DataBoxCustomerDiskJobDetailsArgs struct {
	ContactDetails                 ContactDetailsInput         `pulumi:"contactDetails"`
	DataExportDetails              DataExportDetailsArrayInput `pulumi:"dataExportDetails"`
	DataImportDetails              DataImportDetailsArrayInput `pulumi:"dataImportDetails"`
	EnableManifestBackup           pulumi.BoolPtrInput         `pulumi:"enableManifestBackup"`
	ExpectedDataSizeInTeraBytes    pulumi.IntPtrInput          `pulumi:"expectedDataSizeInTeraBytes"`
	ImportDiskDetailsCollection    ImportDiskDetailsMapInput   `pulumi:"importDiskDetailsCollection"`
	JobDetailsType                 pulumi.StringInput          `pulumi:"jobDetailsType"`
	KeyEncryptionKey               KeyEncryptionKeyPtrInput    `pulumi:"keyEncryptionKey"`
	Preferences                    PreferencesPtrInput         `pulumi:"preferences"`
	ReturnToCustomerPackageDetails PackageCarrierDetailsInput  `pulumi:"returnToCustomerPackageDetails"`
	ShippingAddress                ShippingAddressPtrInput     `pulumi:"shippingAddress"`
}

func (DataBoxCustomerDiskJobDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxCustomerDiskJobDetails)(nil)).Elem()
}

func (i DataBoxCustomerDiskJobDetailsArgs) ToDataBoxCustomerDiskJobDetailsOutput() DataBoxCustomerDiskJobDetailsOutput {
	return i.ToDataBoxCustomerDiskJobDetailsOutputWithContext(context.Background())
}

func (i DataBoxCustomerDiskJobDetailsArgs) ToDataBoxCustomerDiskJobDetailsOutputWithContext(ctx context.Context) DataBoxCustomerDiskJobDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxCustomerDiskJobDetailsOutput)
}

type DataBoxCustomerDiskJobDetailsOutput struct{ *pulumi.OutputState }

func (DataBoxCustomerDiskJobDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxCustomerDiskJobDetails)(nil)).Elem()
}

func (o DataBoxCustomerDiskJobDetailsOutput) ToDataBoxCustomerDiskJobDetailsOutput() DataBoxCustomerDiskJobDetailsOutput {
	return o
}

func (o DataBoxCustomerDiskJobDetailsOutput) ToDataBoxCustomerDiskJobDetailsOutputWithContext(ctx context.Context) DataBoxCustomerDiskJobDetailsOutput {
	return o
}

func (o DataBoxCustomerDiskJobDetailsOutput) ContactDetails() ContactDetailsOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetails) ContactDetails { return v.ContactDetails }).(ContactDetailsOutput)
}

func (o DataBoxCustomerDiskJobDetailsOutput) DataExportDetails() DataExportDetailsArrayOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetails) []DataExportDetails { return v.DataExportDetails }).(DataExportDetailsArrayOutput)
}

func (o DataBoxCustomerDiskJobDetailsOutput) DataImportDetails() DataImportDetailsArrayOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetails) []DataImportDetails { return v.DataImportDetails }).(DataImportDetailsArrayOutput)
}

func (o DataBoxCustomerDiskJobDetailsOutput) EnableManifestBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetails) *bool { return v.EnableManifestBackup }).(pulumi.BoolPtrOutput)
}

func (o DataBoxCustomerDiskJobDetailsOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetails) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxCustomerDiskJobDetailsOutput) ImportDiskDetailsCollection() ImportDiskDetailsMapOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetails) map[string]ImportDiskDetails {
		return v.ImportDiskDetailsCollection
	}).(ImportDiskDetailsMapOutput)
}

func (o DataBoxCustomerDiskJobDetailsOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetails) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskJobDetailsOutput) KeyEncryptionKey() KeyEncryptionKeyPtrOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetails) *KeyEncryptionKey { return v.KeyEncryptionKey }).(KeyEncryptionKeyPtrOutput)
}

func (o DataBoxCustomerDiskJobDetailsOutput) Preferences() PreferencesPtrOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetails) *Preferences { return v.Preferences }).(PreferencesPtrOutput)
}

func (o DataBoxCustomerDiskJobDetailsOutput) ReturnToCustomerPackageDetails() PackageCarrierDetailsOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetails) PackageCarrierDetails { return v.ReturnToCustomerPackageDetails }).(PackageCarrierDetailsOutput)
}

func (o DataBoxCustomerDiskJobDetailsOutput) ShippingAddress() ShippingAddressPtrOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetails) *ShippingAddress { return v.ShippingAddress }).(ShippingAddressPtrOutput)
}

type DataBoxCustomerDiskJobDetailsResponse struct {
	Actions                        []string                                  `pulumi:"actions"`
	ChainOfCustodySasKey           string                                    `pulumi:"chainOfCustodySasKey"`
	ContactDetails                 ContactDetailsResponse                    `pulumi:"contactDetails"`
	CopyLogDetails                 []interface{}                             `pulumi:"copyLogDetails"`
	CopyProgress                   []DataBoxCustomerDiskCopyProgressResponse `pulumi:"copyProgress"`
	DataCenterCode                 string                                    `pulumi:"dataCenterCode"`
	DataExportDetails              []DataExportDetailsResponse               `pulumi:"dataExportDetails"`
	DataImportDetails              []DataImportDetailsResponse               `pulumi:"dataImportDetails"`
	DatacenterAddress              interface{}                               `pulumi:"datacenterAddress"`
	DeliverToDcPackageDetails      PackageCarrierInfoResponse                `pulumi:"deliverToDcPackageDetails"`
	DeliveryPackage                PackageShippingDetailsResponse            `pulumi:"deliveryPackage"`
	EnableManifestBackup           *bool                                     `pulumi:"enableManifestBackup"`
	ExpectedDataSizeInTeraBytes    *int                                      `pulumi:"expectedDataSizeInTeraBytes"`
	ExportDiskDetailsCollection    map[string]ExportDiskDetailsResponse      `pulumi:"exportDiskDetailsCollection"`
	ImportDiskDetailsCollection    map[string]ImportDiskDetailsResponse      `pulumi:"importDiskDetailsCollection"`
	JobDetailsType                 string                                    `pulumi:"jobDetailsType"`
	JobStages                      []JobStagesResponse                       `pulumi:"jobStages"`
	KeyEncryptionKey               *KeyEncryptionKeyResponse                 `pulumi:"keyEncryptionKey"`
	LastMitigationActionOnJob      LastMitigationActionOnJobResponse         `pulumi:"lastMitigationActionOnJob"`
	Preferences                    *PreferencesResponse                      `pulumi:"preferences"`
	ReturnPackage                  PackageShippingDetailsResponse            `pulumi:"returnPackage"`
	ReturnToCustomerPackageDetails PackageCarrierDetailsResponse             `pulumi:"returnToCustomerPackageDetails"`
	ReverseShipmentLabelSasKey     string                                    `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress                *ShippingAddressResponse                  `pulumi:"shippingAddress"`
}





type DataBoxCustomerDiskJobDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxCustomerDiskJobDetailsResponseOutput() DataBoxCustomerDiskJobDetailsResponseOutput
	ToDataBoxCustomerDiskJobDetailsResponseOutputWithContext(context.Context) DataBoxCustomerDiskJobDetailsResponseOutput
}

type DataBoxCustomerDiskJobDetailsResponseArgs struct {
	Actions                        pulumi.StringArrayInput                           `pulumi:"actions"`
	ChainOfCustodySasKey           pulumi.StringInput                                `pulumi:"chainOfCustodySasKey"`
	ContactDetails                 ContactDetailsResponseInput                       `pulumi:"contactDetails"`
	CopyLogDetails                 pulumi.ArrayInput                                 `pulumi:"copyLogDetails"`
	CopyProgress                   DataBoxCustomerDiskCopyProgressResponseArrayInput `pulumi:"copyProgress"`
	DataCenterCode                 pulumi.StringInput                                `pulumi:"dataCenterCode"`
	DataExportDetails              DataExportDetailsResponseArrayInput               `pulumi:"dataExportDetails"`
	DataImportDetails              DataImportDetailsResponseArrayInput               `pulumi:"dataImportDetails"`
	DatacenterAddress              pulumi.Input                                      `pulumi:"datacenterAddress"`
	DeliverToDcPackageDetails      PackageCarrierInfoResponseInput                   `pulumi:"deliverToDcPackageDetails"`
	DeliveryPackage                PackageShippingDetailsResponseInput               `pulumi:"deliveryPackage"`
	EnableManifestBackup           pulumi.BoolPtrInput                               `pulumi:"enableManifestBackup"`
	ExpectedDataSizeInTeraBytes    pulumi.IntPtrInput                                `pulumi:"expectedDataSizeInTeraBytes"`
	ExportDiskDetailsCollection    ExportDiskDetailsResponseMapInput                 `pulumi:"exportDiskDetailsCollection"`
	ImportDiskDetailsCollection    ImportDiskDetailsResponseMapInput                 `pulumi:"importDiskDetailsCollection"`
	JobDetailsType                 pulumi.StringInput                                `pulumi:"jobDetailsType"`
	JobStages                      JobStagesResponseArrayInput                       `pulumi:"jobStages"`
	KeyEncryptionKey               KeyEncryptionKeyResponsePtrInput                  `pulumi:"keyEncryptionKey"`
	LastMitigationActionOnJob      LastMitigationActionOnJobResponseInput            `pulumi:"lastMitigationActionOnJob"`
	Preferences                    PreferencesResponsePtrInput                       `pulumi:"preferences"`
	ReturnPackage                  PackageShippingDetailsResponseInput               `pulumi:"returnPackage"`
	ReturnToCustomerPackageDetails PackageCarrierDetailsResponseInput                `pulumi:"returnToCustomerPackageDetails"`
	ReverseShipmentLabelSasKey     pulumi.StringInput                                `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress                ShippingAddressResponsePtrInput                   `pulumi:"shippingAddress"`
}

func (DataBoxCustomerDiskJobDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxCustomerDiskJobDetailsResponse)(nil)).Elem()
}

func (i DataBoxCustomerDiskJobDetailsResponseArgs) ToDataBoxCustomerDiskJobDetailsResponseOutput() DataBoxCustomerDiskJobDetailsResponseOutput {
	return i.ToDataBoxCustomerDiskJobDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxCustomerDiskJobDetailsResponseArgs) ToDataBoxCustomerDiskJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxCustomerDiskJobDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxCustomerDiskJobDetailsResponseOutput)
}

type DataBoxCustomerDiskJobDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxCustomerDiskJobDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxCustomerDiskJobDetailsResponse)(nil)).Elem()
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) ToDataBoxCustomerDiskJobDetailsResponseOutput() DataBoxCustomerDiskJobDetailsResponseOutput {
	return o
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) ToDataBoxCustomerDiskJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxCustomerDiskJobDetailsResponseOutput {
	return o
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) ChainOfCustodySasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) string { return v.ChainOfCustodySasKey }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) ContactDetailsResponse { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) CopyLogDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) []interface{} { return v.CopyLogDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) CopyProgress() DataBoxCustomerDiskCopyProgressResponseArrayOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) []DataBoxCustomerDiskCopyProgressResponse {
		return v.CopyProgress
	}).(DataBoxCustomerDiskCopyProgressResponseArrayOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) DataCenterCode() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) string { return v.DataCenterCode }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) DataExportDetails() DataExportDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) []DataExportDetailsResponse { return v.DataExportDetails }).(DataExportDetailsResponseArrayOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) DataImportDetails() DataImportDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) []DataImportDetailsResponse { return v.DataImportDetails }).(DataImportDetailsResponseArrayOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) DatacenterAddress() pulumi.AnyOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) interface{} { return v.DatacenterAddress }).(pulumi.AnyOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) DeliverToDcPackageDetails() PackageCarrierInfoResponseOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) PackageCarrierInfoResponse {
		return v.DeliverToDcPackageDetails
	}).(PackageCarrierInfoResponseOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) DeliveryPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) PackageShippingDetailsResponse { return v.DeliveryPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) EnableManifestBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) *bool { return v.EnableManifestBackup }).(pulumi.BoolPtrOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) ExportDiskDetailsCollection() ExportDiskDetailsResponseMapOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) map[string]ExportDiskDetailsResponse {
		return v.ExportDiskDetailsCollection
	}).(ExportDiskDetailsResponseMapOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) ImportDiskDetailsCollection() ImportDiskDetailsResponseMapOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) map[string]ImportDiskDetailsResponse {
		return v.ImportDiskDetailsCollection
	}).(ImportDiskDetailsResponseMapOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) JobStages() JobStagesResponseArrayOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) []JobStagesResponse { return v.JobStages }).(JobStagesResponseArrayOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) KeyEncryptionKey() KeyEncryptionKeyResponsePtrOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) *KeyEncryptionKeyResponse { return v.KeyEncryptionKey }).(KeyEncryptionKeyResponsePtrOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) LastMitigationActionOnJob() LastMitigationActionOnJobResponseOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) LastMitigationActionOnJobResponse {
		return v.LastMitigationActionOnJob
	}).(LastMitigationActionOnJobResponseOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) Preferences() PreferencesResponsePtrOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) *PreferencesResponse { return v.Preferences }).(PreferencesResponsePtrOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) ReturnPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) PackageShippingDetailsResponse { return v.ReturnPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) ReturnToCustomerPackageDetails() PackageCarrierDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) PackageCarrierDetailsResponse {
		return v.ReturnToCustomerPackageDetails
	}).(PackageCarrierDetailsResponseOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) ReverseShipmentLabelSasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) string { return v.ReverseShipmentLabelSasKey }).(pulumi.StringOutput)
}

func (o DataBoxCustomerDiskJobDetailsResponseOutput) ShippingAddress() ShippingAddressResponsePtrOutput {
	return o.ApplyT(func(v DataBoxCustomerDiskJobDetailsResponse) *ShippingAddressResponse { return v.ShippingAddress }).(ShippingAddressResponsePtrOutput)
}

type DataBoxDiskCopyLogDetailsResponse struct {
	CopyLogDetailsType string `pulumi:"copyLogDetailsType"`
	DiskSerialNumber   string `pulumi:"diskSerialNumber"`
	ErrorLogLink       string `pulumi:"errorLogLink"`
	VerboseLogLink     string `pulumi:"verboseLogLink"`
}





type DataBoxDiskCopyLogDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxDiskCopyLogDetailsResponseOutput() DataBoxDiskCopyLogDetailsResponseOutput
	ToDataBoxDiskCopyLogDetailsResponseOutputWithContext(context.Context) DataBoxDiskCopyLogDetailsResponseOutput
}

type DataBoxDiskCopyLogDetailsResponseArgs struct {
	CopyLogDetailsType pulumi.StringInput `pulumi:"copyLogDetailsType"`
	DiskSerialNumber   pulumi.StringInput `pulumi:"diskSerialNumber"`
	ErrorLogLink       pulumi.StringInput `pulumi:"errorLogLink"`
	VerboseLogLink     pulumi.StringInput `pulumi:"verboseLogLink"`
}

func (DataBoxDiskCopyLogDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskCopyLogDetailsResponse)(nil)).Elem()
}

func (i DataBoxDiskCopyLogDetailsResponseArgs) ToDataBoxDiskCopyLogDetailsResponseOutput() DataBoxDiskCopyLogDetailsResponseOutput {
	return i.ToDataBoxDiskCopyLogDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxDiskCopyLogDetailsResponseArgs) ToDataBoxDiskCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxDiskCopyLogDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskCopyLogDetailsResponseOutput)
}

type DataBoxDiskCopyLogDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxDiskCopyLogDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskCopyLogDetailsResponse)(nil)).Elem()
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) ToDataBoxDiskCopyLogDetailsResponseOutput() DataBoxDiskCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) ToDataBoxDiskCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxDiskCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) CopyLogDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyLogDetailsResponse) string { return v.CopyLogDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) DiskSerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyLogDetailsResponse) string { return v.DiskSerialNumber }).(pulumi.StringOutput)
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) ErrorLogLink() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyLogDetailsResponse) string { return v.ErrorLogLink }).(pulumi.StringOutput)
}

func (o DataBoxDiskCopyLogDetailsResponseOutput) VerboseLogLink() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyLogDetailsResponse) string { return v.VerboseLogLink }).(pulumi.StringOutput)
}

type DataBoxDiskCopyProgressResponse struct {
	BytesCopied     float64 `pulumi:"bytesCopied"`
	PercentComplete int     `pulumi:"percentComplete"`
	SerialNumber    string  `pulumi:"serialNumber"`
	Status          string  `pulumi:"status"`
}





type DataBoxDiskCopyProgressResponseInput interface {
	pulumi.Input

	ToDataBoxDiskCopyProgressResponseOutput() DataBoxDiskCopyProgressResponseOutput
	ToDataBoxDiskCopyProgressResponseOutputWithContext(context.Context) DataBoxDiskCopyProgressResponseOutput
}

type DataBoxDiskCopyProgressResponseArgs struct {
	BytesCopied     pulumi.Float64Input `pulumi:"bytesCopied"`
	PercentComplete pulumi.IntInput     `pulumi:"percentComplete"`
	SerialNumber    pulumi.StringInput  `pulumi:"serialNumber"`
	Status          pulumi.StringInput  `pulumi:"status"`
}

func (DataBoxDiskCopyProgressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskCopyProgressResponse)(nil)).Elem()
}

func (i DataBoxDiskCopyProgressResponseArgs) ToDataBoxDiskCopyProgressResponseOutput() DataBoxDiskCopyProgressResponseOutput {
	return i.ToDataBoxDiskCopyProgressResponseOutputWithContext(context.Background())
}

func (i DataBoxDiskCopyProgressResponseArgs) ToDataBoxDiskCopyProgressResponseOutputWithContext(ctx context.Context) DataBoxDiskCopyProgressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskCopyProgressResponseOutput)
}





type DataBoxDiskCopyProgressResponseArrayInput interface {
	pulumi.Input

	ToDataBoxDiskCopyProgressResponseArrayOutput() DataBoxDiskCopyProgressResponseArrayOutput
	ToDataBoxDiskCopyProgressResponseArrayOutputWithContext(context.Context) DataBoxDiskCopyProgressResponseArrayOutput
}

type DataBoxDiskCopyProgressResponseArray []DataBoxDiskCopyProgressResponseInput

func (DataBoxDiskCopyProgressResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxDiskCopyProgressResponse)(nil)).Elem()
}

func (i DataBoxDiskCopyProgressResponseArray) ToDataBoxDiskCopyProgressResponseArrayOutput() DataBoxDiskCopyProgressResponseArrayOutput {
	return i.ToDataBoxDiskCopyProgressResponseArrayOutputWithContext(context.Background())
}

func (i DataBoxDiskCopyProgressResponseArray) ToDataBoxDiskCopyProgressResponseArrayOutputWithContext(ctx context.Context) DataBoxDiskCopyProgressResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskCopyProgressResponseArrayOutput)
}

type DataBoxDiskCopyProgressResponseOutput struct{ *pulumi.OutputState }

func (DataBoxDiskCopyProgressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskCopyProgressResponse)(nil)).Elem()
}

func (o DataBoxDiskCopyProgressResponseOutput) ToDataBoxDiskCopyProgressResponseOutput() DataBoxDiskCopyProgressResponseOutput {
	return o
}

func (o DataBoxDiskCopyProgressResponseOutput) ToDataBoxDiskCopyProgressResponseOutputWithContext(ctx context.Context) DataBoxDiskCopyProgressResponseOutput {
	return o
}

func (o DataBoxDiskCopyProgressResponseOutput) BytesCopied() pulumi.Float64Output {
	return o.ApplyT(func(v DataBoxDiskCopyProgressResponse) float64 { return v.BytesCopied }).(pulumi.Float64Output)
}

func (o DataBoxDiskCopyProgressResponseOutput) PercentComplete() pulumi.IntOutput {
	return o.ApplyT(func(v DataBoxDiskCopyProgressResponse) int { return v.PercentComplete }).(pulumi.IntOutput)
}

func (o DataBoxDiskCopyProgressResponseOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyProgressResponse) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o DataBoxDiskCopyProgressResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskCopyProgressResponse) string { return v.Status }).(pulumi.StringOutput)
}

type DataBoxDiskCopyProgressResponseArrayOutput struct{ *pulumi.OutputState }

func (DataBoxDiskCopyProgressResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxDiskCopyProgressResponse)(nil)).Elem()
}

func (o DataBoxDiskCopyProgressResponseArrayOutput) ToDataBoxDiskCopyProgressResponseArrayOutput() DataBoxDiskCopyProgressResponseArrayOutput {
	return o
}

func (o DataBoxDiskCopyProgressResponseArrayOutput) ToDataBoxDiskCopyProgressResponseArrayOutputWithContext(ctx context.Context) DataBoxDiskCopyProgressResponseArrayOutput {
	return o
}

func (o DataBoxDiskCopyProgressResponseArrayOutput) Index(i pulumi.IntInput) DataBoxDiskCopyProgressResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataBoxDiskCopyProgressResponse {
		return vs[0].([]DataBoxDiskCopyProgressResponse)[vs[1].(int)]
	}).(DataBoxDiskCopyProgressResponseOutput)
}

type DataBoxDiskJobDetails struct {
	ContactDetails              ContactDetails      `pulumi:"contactDetails"`
	DataExportDetails           []DataExportDetails `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetails `pulumi:"dataImportDetails"`
	ExpectedDataSizeInTeraBytes *int                `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string              `pulumi:"jobDetailsType"`
	KeyEncryptionKey            *KeyEncryptionKey   `pulumi:"keyEncryptionKey"`
	Passkey                     *string             `pulumi:"passkey"`
	Preferences                 *Preferences        `pulumi:"preferences"`
	PreferredDisks              map[string]int      `pulumi:"preferredDisks"`
	ShippingAddress             *ShippingAddress    `pulumi:"shippingAddress"`
}





type DataBoxDiskJobDetailsInput interface {
	pulumi.Input

	ToDataBoxDiskJobDetailsOutput() DataBoxDiskJobDetailsOutput
	ToDataBoxDiskJobDetailsOutputWithContext(context.Context) DataBoxDiskJobDetailsOutput
}

type DataBoxDiskJobDetailsArgs struct {
	ContactDetails              ContactDetailsInput         `pulumi:"contactDetails"`
	DataExportDetails           DataExportDetailsArrayInput `pulumi:"dataExportDetails"`
	DataImportDetails           DataImportDetailsArrayInput `pulumi:"dataImportDetails"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput          `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput          `pulumi:"jobDetailsType"`
	KeyEncryptionKey            KeyEncryptionKeyPtrInput    `pulumi:"keyEncryptionKey"`
	Passkey                     pulumi.StringPtrInput       `pulumi:"passkey"`
	Preferences                 PreferencesPtrInput         `pulumi:"preferences"`
	PreferredDisks              pulumi.IntMapInput          `pulumi:"preferredDisks"`
	ShippingAddress             ShippingAddressPtrInput     `pulumi:"shippingAddress"`
}

func (DataBoxDiskJobDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobDetails)(nil)).Elem()
}

func (i DataBoxDiskJobDetailsArgs) ToDataBoxDiskJobDetailsOutput() DataBoxDiskJobDetailsOutput {
	return i.ToDataBoxDiskJobDetailsOutputWithContext(context.Background())
}

func (i DataBoxDiskJobDetailsArgs) ToDataBoxDiskJobDetailsOutputWithContext(ctx context.Context) DataBoxDiskJobDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskJobDetailsOutput)
}

type DataBoxDiskJobDetailsOutput struct{ *pulumi.OutputState }

func (DataBoxDiskJobDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobDetails)(nil)).Elem()
}

func (o DataBoxDiskJobDetailsOutput) ToDataBoxDiskJobDetailsOutput() DataBoxDiskJobDetailsOutput {
	return o
}

func (o DataBoxDiskJobDetailsOutput) ToDataBoxDiskJobDetailsOutputWithContext(ctx context.Context) DataBoxDiskJobDetailsOutput {
	return o
}

func (o DataBoxDiskJobDetailsOutput) ContactDetails() ContactDetailsOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) ContactDetails { return v.ContactDetails }).(ContactDetailsOutput)
}

func (o DataBoxDiskJobDetailsOutput) DataExportDetails() DataExportDetailsArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) []DataExportDetails { return v.DataExportDetails }).(DataExportDetailsArrayOutput)
}

func (o DataBoxDiskJobDetailsOutput) DataImportDetails() DataImportDetailsArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) []DataImportDetails { return v.DataImportDetails }).(DataImportDetailsArrayOutput)
}

func (o DataBoxDiskJobDetailsOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxDiskJobDetailsOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxDiskJobDetailsOutput) KeyEncryptionKey() KeyEncryptionKeyPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) *KeyEncryptionKey { return v.KeyEncryptionKey }).(KeyEncryptionKeyPtrOutput)
}

func (o DataBoxDiskJobDetailsOutput) Passkey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) *string { return v.Passkey }).(pulumi.StringPtrOutput)
}

func (o DataBoxDiskJobDetailsOutput) Preferences() PreferencesPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) *Preferences { return v.Preferences }).(PreferencesPtrOutput)
}

func (o DataBoxDiskJobDetailsOutput) PreferredDisks() pulumi.IntMapOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) map[string]int { return v.PreferredDisks }).(pulumi.IntMapOutput)
}

func (o DataBoxDiskJobDetailsOutput) ShippingAddress() ShippingAddressPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetails) *ShippingAddress { return v.ShippingAddress }).(ShippingAddressPtrOutput)
}

type DataBoxDiskJobDetailsResponse struct {
	Actions                     []string                          `pulumi:"actions"`
	ChainOfCustodySasKey        string                            `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponse            `pulumi:"contactDetails"`
	CopyLogDetails              []interface{}                     `pulumi:"copyLogDetails"`
	CopyProgress                []DataBoxDiskCopyProgressResponse `pulumi:"copyProgress"`
	DataCenterCode              string                            `pulumi:"dataCenterCode"`
	DataExportDetails           []DataExportDetailsResponse       `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetailsResponse       `pulumi:"dataImportDetails"`
	DatacenterAddress           interface{}                       `pulumi:"datacenterAddress"`
	DeliveryPackage             PackageShippingDetailsResponse    `pulumi:"deliveryPackage"`
	DisksAndSizeDetails         map[string]int                    `pulumi:"disksAndSizeDetails"`
	ExpectedDataSizeInTeraBytes *int                              `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string                            `pulumi:"jobDetailsType"`
	JobStages                   []JobStagesResponse               `pulumi:"jobStages"`
	KeyEncryptionKey            *KeyEncryptionKeyResponse         `pulumi:"keyEncryptionKey"`
	LastMitigationActionOnJob   LastMitigationActionOnJobResponse `pulumi:"lastMitigationActionOnJob"`
	Passkey                     *string                           `pulumi:"passkey"`
	Preferences                 *PreferencesResponse              `pulumi:"preferences"`
	PreferredDisks              map[string]int                    `pulumi:"preferredDisks"`
	ReturnPackage               PackageShippingDetailsResponse    `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  string                            `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             *ShippingAddressResponse          `pulumi:"shippingAddress"`
}





type DataBoxDiskJobDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxDiskJobDetailsResponseOutput() DataBoxDiskJobDetailsResponseOutput
	ToDataBoxDiskJobDetailsResponseOutputWithContext(context.Context) DataBoxDiskJobDetailsResponseOutput
}

type DataBoxDiskJobDetailsResponseArgs struct {
	Actions                     pulumi.StringArrayInput                   `pulumi:"actions"`
	ChainOfCustodySasKey        pulumi.StringInput                        `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponseInput               `pulumi:"contactDetails"`
	CopyLogDetails              pulumi.ArrayInput                         `pulumi:"copyLogDetails"`
	CopyProgress                DataBoxDiskCopyProgressResponseArrayInput `pulumi:"copyProgress"`
	DataCenterCode              pulumi.StringInput                        `pulumi:"dataCenterCode"`
	DataExportDetails           DataExportDetailsResponseArrayInput       `pulumi:"dataExportDetails"`
	DataImportDetails           DataImportDetailsResponseArrayInput       `pulumi:"dataImportDetails"`
	DatacenterAddress           pulumi.Input                              `pulumi:"datacenterAddress"`
	DeliveryPackage             PackageShippingDetailsResponseInput       `pulumi:"deliveryPackage"`
	DisksAndSizeDetails         pulumi.IntMapInput                        `pulumi:"disksAndSizeDetails"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput                        `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput                        `pulumi:"jobDetailsType"`
	JobStages                   JobStagesResponseArrayInput               `pulumi:"jobStages"`
	KeyEncryptionKey            KeyEncryptionKeyResponsePtrInput          `pulumi:"keyEncryptionKey"`
	LastMitigationActionOnJob   LastMitigationActionOnJobResponseInput    `pulumi:"lastMitigationActionOnJob"`
	Passkey                     pulumi.StringPtrInput                     `pulumi:"passkey"`
	Preferences                 PreferencesResponsePtrInput               `pulumi:"preferences"`
	PreferredDisks              pulumi.IntMapInput                        `pulumi:"preferredDisks"`
	ReturnPackage               PackageShippingDetailsResponseInput       `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  pulumi.StringInput                        `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             ShippingAddressResponsePtrInput           `pulumi:"shippingAddress"`
}

func (DataBoxDiskJobDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobDetailsResponse)(nil)).Elem()
}

func (i DataBoxDiskJobDetailsResponseArgs) ToDataBoxDiskJobDetailsResponseOutput() DataBoxDiskJobDetailsResponseOutput {
	return i.ToDataBoxDiskJobDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxDiskJobDetailsResponseArgs) ToDataBoxDiskJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxDiskJobDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskJobDetailsResponseOutput)
}

type DataBoxDiskJobDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxDiskJobDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobDetailsResponse)(nil)).Elem()
}

func (o DataBoxDiskJobDetailsResponseOutput) ToDataBoxDiskJobDetailsResponseOutput() DataBoxDiskJobDetailsResponseOutput {
	return o
}

func (o DataBoxDiskJobDetailsResponseOutput) ToDataBoxDiskJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxDiskJobDetailsResponseOutput {
	return o
}

func (o DataBoxDiskJobDetailsResponseOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ChainOfCustodySasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) string { return v.ChainOfCustodySasKey }).(pulumi.StringOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) ContactDetailsResponse { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) CopyLogDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) []interface{} { return v.CopyLogDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) CopyProgress() DataBoxDiskCopyProgressResponseArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) []DataBoxDiskCopyProgressResponse { return v.CopyProgress }).(DataBoxDiskCopyProgressResponseArrayOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) DataCenterCode() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) string { return v.DataCenterCode }).(pulumi.StringOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) DataExportDetails() DataExportDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) []DataExportDetailsResponse { return v.DataExportDetails }).(DataExportDetailsResponseArrayOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) DataImportDetails() DataImportDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) []DataImportDetailsResponse { return v.DataImportDetails }).(DataImportDetailsResponseArrayOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) DatacenterAddress() pulumi.AnyOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) interface{} { return v.DatacenterAddress }).(pulumi.AnyOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) DeliveryPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) PackageShippingDetailsResponse { return v.DeliveryPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) DisksAndSizeDetails() pulumi.IntMapOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) map[string]int { return v.DisksAndSizeDetails }).(pulumi.IntMapOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) JobStages() JobStagesResponseArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) []JobStagesResponse { return v.JobStages }).(JobStagesResponseArrayOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) KeyEncryptionKey() KeyEncryptionKeyResponsePtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) *KeyEncryptionKeyResponse { return v.KeyEncryptionKey }).(KeyEncryptionKeyResponsePtrOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) LastMitigationActionOnJob() LastMitigationActionOnJobResponseOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) LastMitigationActionOnJobResponse {
		return v.LastMitigationActionOnJob
	}).(LastMitigationActionOnJobResponseOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) Passkey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) *string { return v.Passkey }).(pulumi.StringPtrOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) Preferences() PreferencesResponsePtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) *PreferencesResponse { return v.Preferences }).(PreferencesResponsePtrOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) PreferredDisks() pulumi.IntMapOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) map[string]int { return v.PreferredDisks }).(pulumi.IntMapOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ReturnPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) PackageShippingDetailsResponse { return v.ReturnPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ReverseShipmentLabelSasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) string { return v.ReverseShipmentLabelSasKey }).(pulumi.StringOutput)
}

func (o DataBoxDiskJobDetailsResponseOutput) ShippingAddress() ShippingAddressResponsePtrOutput {
	return o.ApplyT(func(v DataBoxDiskJobDetailsResponse) *ShippingAddressResponse { return v.ShippingAddress }).(ShippingAddressResponsePtrOutput)
}

type DataBoxDiskJobSecretsResponse struct {
	DcAccessSecurityCode DcAccessSecurityCodeResponse `pulumi:"dcAccessSecurityCode"`
	DiskSecrets          []DiskSecretResponse         `pulumi:"diskSecrets"`
	Error                CloudErrorResponse           `pulumi:"error"`
	IsPasskeyUserDefined bool                         `pulumi:"isPasskeyUserDefined"`
	JobSecretsType       string                       `pulumi:"jobSecretsType"`
	PassKey              string                       `pulumi:"passKey"`
}





type DataBoxDiskJobSecretsResponseInput interface {
	pulumi.Input

	ToDataBoxDiskJobSecretsResponseOutput() DataBoxDiskJobSecretsResponseOutput
	ToDataBoxDiskJobSecretsResponseOutputWithContext(context.Context) DataBoxDiskJobSecretsResponseOutput
}

type DataBoxDiskJobSecretsResponseArgs struct {
	DcAccessSecurityCode DcAccessSecurityCodeResponseInput `pulumi:"dcAccessSecurityCode"`
	DiskSecrets          DiskSecretResponseArrayInput      `pulumi:"diskSecrets"`
	Error                CloudErrorResponseInput           `pulumi:"error"`
	IsPasskeyUserDefined pulumi.BoolInput                  `pulumi:"isPasskeyUserDefined"`
	JobSecretsType       pulumi.StringInput                `pulumi:"jobSecretsType"`
	PassKey              pulumi.StringInput                `pulumi:"passKey"`
}

func (DataBoxDiskJobSecretsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobSecretsResponse)(nil)).Elem()
}

func (i DataBoxDiskJobSecretsResponseArgs) ToDataBoxDiskJobSecretsResponseOutput() DataBoxDiskJobSecretsResponseOutput {
	return i.ToDataBoxDiskJobSecretsResponseOutputWithContext(context.Background())
}

func (i DataBoxDiskJobSecretsResponseArgs) ToDataBoxDiskJobSecretsResponseOutputWithContext(ctx context.Context) DataBoxDiskJobSecretsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxDiskJobSecretsResponseOutput)
}

type DataBoxDiskJobSecretsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxDiskJobSecretsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxDiskJobSecretsResponse)(nil)).Elem()
}

func (o DataBoxDiskJobSecretsResponseOutput) ToDataBoxDiskJobSecretsResponseOutput() DataBoxDiskJobSecretsResponseOutput {
	return o
}

func (o DataBoxDiskJobSecretsResponseOutput) ToDataBoxDiskJobSecretsResponseOutputWithContext(ctx context.Context) DataBoxDiskJobSecretsResponseOutput {
	return o
}

func (o DataBoxDiskJobSecretsResponseOutput) DcAccessSecurityCode() DcAccessSecurityCodeResponseOutput {
	return o.ApplyT(func(v DataBoxDiskJobSecretsResponse) DcAccessSecurityCodeResponse { return v.DcAccessSecurityCode }).(DcAccessSecurityCodeResponseOutput)
}

func (o DataBoxDiskJobSecretsResponseOutput) DiskSecrets() DiskSecretResponseArrayOutput {
	return o.ApplyT(func(v DataBoxDiskJobSecretsResponse) []DiskSecretResponse { return v.DiskSecrets }).(DiskSecretResponseArrayOutput)
}

func (o DataBoxDiskJobSecretsResponseOutput) Error() CloudErrorResponseOutput {
	return o.ApplyT(func(v DataBoxDiskJobSecretsResponse) CloudErrorResponse { return v.Error }).(CloudErrorResponseOutput)
}

func (o DataBoxDiskJobSecretsResponseOutput) IsPasskeyUserDefined() pulumi.BoolOutput {
	return o.ApplyT(func(v DataBoxDiskJobSecretsResponse) bool { return v.IsPasskeyUserDefined }).(pulumi.BoolOutput)
}

func (o DataBoxDiskJobSecretsResponseOutput) JobSecretsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobSecretsResponse) string { return v.JobSecretsType }).(pulumi.StringOutput)
}

func (o DataBoxDiskJobSecretsResponseOutput) PassKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxDiskJobSecretsResponse) string { return v.PassKey }).(pulumi.StringOutput)
}

type DataBoxHeavyAccountCopyLogDetailsResponse struct {
	AccountName        string   `pulumi:"accountName"`
	CopyLogDetailsType string   `pulumi:"copyLogDetailsType"`
	CopyLogLink        []string `pulumi:"copyLogLink"`
	CopyVerboseLogLink []string `pulumi:"copyVerboseLogLink"`
}





type DataBoxHeavyAccountCopyLogDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxHeavyAccountCopyLogDetailsResponseOutput() DataBoxHeavyAccountCopyLogDetailsResponseOutput
	ToDataBoxHeavyAccountCopyLogDetailsResponseOutputWithContext(context.Context) DataBoxHeavyAccountCopyLogDetailsResponseOutput
}

type DataBoxHeavyAccountCopyLogDetailsResponseArgs struct {
	AccountName        pulumi.StringInput      `pulumi:"accountName"`
	CopyLogDetailsType pulumi.StringInput      `pulumi:"copyLogDetailsType"`
	CopyLogLink        pulumi.StringArrayInput `pulumi:"copyLogLink"`
	CopyVerboseLogLink pulumi.StringArrayInput `pulumi:"copyVerboseLogLink"`
}

func (DataBoxHeavyAccountCopyLogDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyAccountCopyLogDetailsResponse)(nil)).Elem()
}

func (i DataBoxHeavyAccountCopyLogDetailsResponseArgs) ToDataBoxHeavyAccountCopyLogDetailsResponseOutput() DataBoxHeavyAccountCopyLogDetailsResponseOutput {
	return i.ToDataBoxHeavyAccountCopyLogDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxHeavyAccountCopyLogDetailsResponseArgs) ToDataBoxHeavyAccountCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxHeavyAccountCopyLogDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavyAccountCopyLogDetailsResponseOutput)
}

type DataBoxHeavyAccountCopyLogDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxHeavyAccountCopyLogDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyAccountCopyLogDetailsResponse)(nil)).Elem()
}

func (o DataBoxHeavyAccountCopyLogDetailsResponseOutput) ToDataBoxHeavyAccountCopyLogDetailsResponseOutput() DataBoxHeavyAccountCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxHeavyAccountCopyLogDetailsResponseOutput) ToDataBoxHeavyAccountCopyLogDetailsResponseOutputWithContext(ctx context.Context) DataBoxHeavyAccountCopyLogDetailsResponseOutput {
	return o
}

func (o DataBoxHeavyAccountCopyLogDetailsResponseOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyAccountCopyLogDetailsResponse) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o DataBoxHeavyAccountCopyLogDetailsResponseOutput) CopyLogDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyAccountCopyLogDetailsResponse) string { return v.CopyLogDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxHeavyAccountCopyLogDetailsResponseOutput) CopyLogLink() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyAccountCopyLogDetailsResponse) []string { return v.CopyLogLink }).(pulumi.StringArrayOutput)
}

func (o DataBoxHeavyAccountCopyLogDetailsResponseOutput) CopyVerboseLogLink() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyAccountCopyLogDetailsResponse) []string { return v.CopyVerboseLogLink }).(pulumi.StringArrayOutput)
}

type DataBoxHeavyJobDetails struct {
	ContactDetails              ContactDetails      `pulumi:"contactDetails"`
	DataExportDetails           []DataExportDetails `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetails `pulumi:"dataImportDetails"`
	DevicePassword              *string             `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes *int                `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string              `pulumi:"jobDetailsType"`
	KeyEncryptionKey            *KeyEncryptionKey   `pulumi:"keyEncryptionKey"`
	Preferences                 *Preferences        `pulumi:"preferences"`
	ShippingAddress             *ShippingAddress    `pulumi:"shippingAddress"`
}





type DataBoxHeavyJobDetailsInput interface {
	pulumi.Input

	ToDataBoxHeavyJobDetailsOutput() DataBoxHeavyJobDetailsOutput
	ToDataBoxHeavyJobDetailsOutputWithContext(context.Context) DataBoxHeavyJobDetailsOutput
}

type DataBoxHeavyJobDetailsArgs struct {
	ContactDetails              ContactDetailsInput         `pulumi:"contactDetails"`
	DataExportDetails           DataExportDetailsArrayInput `pulumi:"dataExportDetails"`
	DataImportDetails           DataImportDetailsArrayInput `pulumi:"dataImportDetails"`
	DevicePassword              pulumi.StringPtrInput       `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput          `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput          `pulumi:"jobDetailsType"`
	KeyEncryptionKey            KeyEncryptionKeyPtrInput    `pulumi:"keyEncryptionKey"`
	Preferences                 PreferencesPtrInput         `pulumi:"preferences"`
	ShippingAddress             ShippingAddressPtrInput     `pulumi:"shippingAddress"`
}

func (DataBoxHeavyJobDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobDetails)(nil)).Elem()
}

func (i DataBoxHeavyJobDetailsArgs) ToDataBoxHeavyJobDetailsOutput() DataBoxHeavyJobDetailsOutput {
	return i.ToDataBoxHeavyJobDetailsOutputWithContext(context.Background())
}

func (i DataBoxHeavyJobDetailsArgs) ToDataBoxHeavyJobDetailsOutputWithContext(ctx context.Context) DataBoxHeavyJobDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavyJobDetailsOutput)
}

type DataBoxHeavyJobDetailsOutput struct{ *pulumi.OutputState }

func (DataBoxHeavyJobDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobDetails)(nil)).Elem()
}

func (o DataBoxHeavyJobDetailsOutput) ToDataBoxHeavyJobDetailsOutput() DataBoxHeavyJobDetailsOutput {
	return o
}

func (o DataBoxHeavyJobDetailsOutput) ToDataBoxHeavyJobDetailsOutputWithContext(ctx context.Context) DataBoxHeavyJobDetailsOutput {
	return o
}

func (o DataBoxHeavyJobDetailsOutput) ContactDetails() ContactDetailsOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) ContactDetails { return v.ContactDetails }).(ContactDetailsOutput)
}

func (o DataBoxHeavyJobDetailsOutput) DataExportDetails() DataExportDetailsArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) []DataExportDetails { return v.DataExportDetails }).(DataExportDetailsArrayOutput)
}

func (o DataBoxHeavyJobDetailsOutput) DataImportDetails() DataImportDetailsArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) []DataImportDetails { return v.DataImportDetails }).(DataImportDetailsArrayOutput)
}

func (o DataBoxHeavyJobDetailsOutput) DevicePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) *string { return v.DevicePassword }).(pulumi.StringPtrOutput)
}

func (o DataBoxHeavyJobDetailsOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxHeavyJobDetailsOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxHeavyJobDetailsOutput) KeyEncryptionKey() KeyEncryptionKeyPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) *KeyEncryptionKey { return v.KeyEncryptionKey }).(KeyEncryptionKeyPtrOutput)
}

func (o DataBoxHeavyJobDetailsOutput) Preferences() PreferencesPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) *Preferences { return v.Preferences }).(PreferencesPtrOutput)
}

func (o DataBoxHeavyJobDetailsOutput) ShippingAddress() ShippingAddressPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetails) *ShippingAddress { return v.ShippingAddress }).(ShippingAddressPtrOutput)
}

type DataBoxHeavyJobDetailsResponse struct {
	Actions                     []string                          `pulumi:"actions"`
	ChainOfCustodySasKey        string                            `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponse            `pulumi:"contactDetails"`
	CopyLogDetails              []interface{}                     `pulumi:"copyLogDetails"`
	CopyProgress                []CopyProgressResponse            `pulumi:"copyProgress"`
	DataCenterCode              string                            `pulumi:"dataCenterCode"`
	DataExportDetails           []DataExportDetailsResponse       `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetailsResponse       `pulumi:"dataImportDetails"`
	DatacenterAddress           interface{}                       `pulumi:"datacenterAddress"`
	DeliveryPackage             PackageShippingDetailsResponse    `pulumi:"deliveryPackage"`
	DevicePassword              *string                           `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes *int                              `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string                            `pulumi:"jobDetailsType"`
	JobStages                   []JobStagesResponse               `pulumi:"jobStages"`
	KeyEncryptionKey            *KeyEncryptionKeyResponse         `pulumi:"keyEncryptionKey"`
	LastMitigationActionOnJob   LastMitigationActionOnJobResponse `pulumi:"lastMitigationActionOnJob"`
	Preferences                 *PreferencesResponse              `pulumi:"preferences"`
	ReturnPackage               PackageShippingDetailsResponse    `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  string                            `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             *ShippingAddressResponse          `pulumi:"shippingAddress"`
}





type DataBoxHeavyJobDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxHeavyJobDetailsResponseOutput() DataBoxHeavyJobDetailsResponseOutput
	ToDataBoxHeavyJobDetailsResponseOutputWithContext(context.Context) DataBoxHeavyJobDetailsResponseOutput
}

type DataBoxHeavyJobDetailsResponseArgs struct {
	Actions                     pulumi.StringArrayInput                `pulumi:"actions"`
	ChainOfCustodySasKey        pulumi.StringInput                     `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponseInput            `pulumi:"contactDetails"`
	CopyLogDetails              pulumi.ArrayInput                      `pulumi:"copyLogDetails"`
	CopyProgress                CopyProgressResponseArrayInput         `pulumi:"copyProgress"`
	DataCenterCode              pulumi.StringInput                     `pulumi:"dataCenterCode"`
	DataExportDetails           DataExportDetailsResponseArrayInput    `pulumi:"dataExportDetails"`
	DataImportDetails           DataImportDetailsResponseArrayInput    `pulumi:"dataImportDetails"`
	DatacenterAddress           pulumi.Input                           `pulumi:"datacenterAddress"`
	DeliveryPackage             PackageShippingDetailsResponseInput    `pulumi:"deliveryPackage"`
	DevicePassword              pulumi.StringPtrInput                  `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput                     `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput                     `pulumi:"jobDetailsType"`
	JobStages                   JobStagesResponseArrayInput            `pulumi:"jobStages"`
	KeyEncryptionKey            KeyEncryptionKeyResponsePtrInput       `pulumi:"keyEncryptionKey"`
	LastMitigationActionOnJob   LastMitigationActionOnJobResponseInput `pulumi:"lastMitigationActionOnJob"`
	Preferences                 PreferencesResponsePtrInput            `pulumi:"preferences"`
	ReturnPackage               PackageShippingDetailsResponseInput    `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  pulumi.StringInput                     `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             ShippingAddressResponsePtrInput        `pulumi:"shippingAddress"`
}

func (DataBoxHeavyJobDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobDetailsResponse)(nil)).Elem()
}

func (i DataBoxHeavyJobDetailsResponseArgs) ToDataBoxHeavyJobDetailsResponseOutput() DataBoxHeavyJobDetailsResponseOutput {
	return i.ToDataBoxHeavyJobDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxHeavyJobDetailsResponseArgs) ToDataBoxHeavyJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxHeavyJobDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavyJobDetailsResponseOutput)
}

type DataBoxHeavyJobDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxHeavyJobDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobDetailsResponse)(nil)).Elem()
}

func (o DataBoxHeavyJobDetailsResponseOutput) ToDataBoxHeavyJobDetailsResponseOutput() DataBoxHeavyJobDetailsResponseOutput {
	return o
}

func (o DataBoxHeavyJobDetailsResponseOutput) ToDataBoxHeavyJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxHeavyJobDetailsResponseOutput {
	return o
}

func (o DataBoxHeavyJobDetailsResponseOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ChainOfCustodySasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) string { return v.ChainOfCustodySasKey }).(pulumi.StringOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) ContactDetailsResponse { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) CopyLogDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) []interface{} { return v.CopyLogDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) CopyProgress() CopyProgressResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) []CopyProgressResponse { return v.CopyProgress }).(CopyProgressResponseArrayOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) DataCenterCode() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) string { return v.DataCenterCode }).(pulumi.StringOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) DataExportDetails() DataExportDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) []DataExportDetailsResponse { return v.DataExportDetails }).(DataExportDetailsResponseArrayOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) DataImportDetails() DataImportDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) []DataImportDetailsResponse { return v.DataImportDetails }).(DataImportDetailsResponseArrayOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) DatacenterAddress() pulumi.AnyOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) interface{} { return v.DatacenterAddress }).(pulumi.AnyOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) DeliveryPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) PackageShippingDetailsResponse { return v.DeliveryPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) DevicePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) *string { return v.DevicePassword }).(pulumi.StringPtrOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) JobStages() JobStagesResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) []JobStagesResponse { return v.JobStages }).(JobStagesResponseArrayOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) KeyEncryptionKey() KeyEncryptionKeyResponsePtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) *KeyEncryptionKeyResponse { return v.KeyEncryptionKey }).(KeyEncryptionKeyResponsePtrOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) LastMitigationActionOnJob() LastMitigationActionOnJobResponseOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) LastMitigationActionOnJobResponse {
		return v.LastMitigationActionOnJob
	}).(LastMitigationActionOnJobResponseOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) Preferences() PreferencesResponsePtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) *PreferencesResponse { return v.Preferences }).(PreferencesResponsePtrOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ReturnPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) PackageShippingDetailsResponse { return v.ReturnPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ReverseShipmentLabelSasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) string { return v.ReverseShipmentLabelSasKey }).(pulumi.StringOutput)
}

func (o DataBoxHeavyJobDetailsResponseOutput) ShippingAddress() ShippingAddressResponsePtrOutput {
	return o.ApplyT(func(v DataBoxHeavyJobDetailsResponse) *ShippingAddressResponse { return v.ShippingAddress }).(ShippingAddressResponsePtrOutput)
}

type DataBoxHeavyJobSecretsResponse struct {
	CabinetPodSecrets    []DataBoxHeavySecretResponse `pulumi:"cabinetPodSecrets"`
	DcAccessSecurityCode DcAccessSecurityCodeResponse `pulumi:"dcAccessSecurityCode"`
	Error                CloudErrorResponse           `pulumi:"error"`
	JobSecretsType       string                       `pulumi:"jobSecretsType"`
}





type DataBoxHeavyJobSecretsResponseInput interface {
	pulumi.Input

	ToDataBoxHeavyJobSecretsResponseOutput() DataBoxHeavyJobSecretsResponseOutput
	ToDataBoxHeavyJobSecretsResponseOutputWithContext(context.Context) DataBoxHeavyJobSecretsResponseOutput
}

type DataBoxHeavyJobSecretsResponseArgs struct {
	CabinetPodSecrets    DataBoxHeavySecretResponseArrayInput `pulumi:"cabinetPodSecrets"`
	DcAccessSecurityCode DcAccessSecurityCodeResponseInput    `pulumi:"dcAccessSecurityCode"`
	Error                CloudErrorResponseInput              `pulumi:"error"`
	JobSecretsType       pulumi.StringInput                   `pulumi:"jobSecretsType"`
}

func (DataBoxHeavyJobSecretsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobSecretsResponse)(nil)).Elem()
}

func (i DataBoxHeavyJobSecretsResponseArgs) ToDataBoxHeavyJobSecretsResponseOutput() DataBoxHeavyJobSecretsResponseOutput {
	return i.ToDataBoxHeavyJobSecretsResponseOutputWithContext(context.Background())
}

func (i DataBoxHeavyJobSecretsResponseArgs) ToDataBoxHeavyJobSecretsResponseOutputWithContext(ctx context.Context) DataBoxHeavyJobSecretsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavyJobSecretsResponseOutput)
}

type DataBoxHeavyJobSecretsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxHeavyJobSecretsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavyJobSecretsResponse)(nil)).Elem()
}

func (o DataBoxHeavyJobSecretsResponseOutput) ToDataBoxHeavyJobSecretsResponseOutput() DataBoxHeavyJobSecretsResponseOutput {
	return o
}

func (o DataBoxHeavyJobSecretsResponseOutput) ToDataBoxHeavyJobSecretsResponseOutputWithContext(ctx context.Context) DataBoxHeavyJobSecretsResponseOutput {
	return o
}

func (o DataBoxHeavyJobSecretsResponseOutput) CabinetPodSecrets() DataBoxHeavySecretResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavyJobSecretsResponse) []DataBoxHeavySecretResponse { return v.CabinetPodSecrets }).(DataBoxHeavySecretResponseArrayOutput)
}

func (o DataBoxHeavyJobSecretsResponseOutput) DcAccessSecurityCode() DcAccessSecurityCodeResponseOutput {
	return o.ApplyT(func(v DataBoxHeavyJobSecretsResponse) DcAccessSecurityCodeResponse { return v.DcAccessSecurityCode }).(DcAccessSecurityCodeResponseOutput)
}

func (o DataBoxHeavyJobSecretsResponseOutput) Error() CloudErrorResponseOutput {
	return o.ApplyT(func(v DataBoxHeavyJobSecretsResponse) CloudErrorResponse { return v.Error }).(CloudErrorResponseOutput)
}

func (o DataBoxHeavyJobSecretsResponseOutput) JobSecretsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavyJobSecretsResponse) string { return v.JobSecretsType }).(pulumi.StringOutput)
}

type DataBoxHeavySecretResponse struct {
	AccountCredentialDetails    []AccountCredentialDetailsResponse      `pulumi:"accountCredentialDetails"`
	DevicePassword              string                                  `pulumi:"devicePassword"`
	DeviceSerialNumber          string                                  `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey string                                  `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       []ApplianceNetworkConfigurationResponse `pulumi:"networkConfigurations"`
}





type DataBoxHeavySecretResponseInput interface {
	pulumi.Input

	ToDataBoxHeavySecretResponseOutput() DataBoxHeavySecretResponseOutput
	ToDataBoxHeavySecretResponseOutputWithContext(context.Context) DataBoxHeavySecretResponseOutput
}

type DataBoxHeavySecretResponseArgs struct {
	AccountCredentialDetails    AccountCredentialDetailsResponseArrayInput      `pulumi:"accountCredentialDetails"`
	DevicePassword              pulumi.StringInput                              `pulumi:"devicePassword"`
	DeviceSerialNumber          pulumi.StringInput                              `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey pulumi.StringInput                              `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       ApplianceNetworkConfigurationResponseArrayInput `pulumi:"networkConfigurations"`
}

func (DataBoxHeavySecretResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavySecretResponse)(nil)).Elem()
}

func (i DataBoxHeavySecretResponseArgs) ToDataBoxHeavySecretResponseOutput() DataBoxHeavySecretResponseOutput {
	return i.ToDataBoxHeavySecretResponseOutputWithContext(context.Background())
}

func (i DataBoxHeavySecretResponseArgs) ToDataBoxHeavySecretResponseOutputWithContext(ctx context.Context) DataBoxHeavySecretResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavySecretResponseOutput)
}





type DataBoxHeavySecretResponseArrayInput interface {
	pulumi.Input

	ToDataBoxHeavySecretResponseArrayOutput() DataBoxHeavySecretResponseArrayOutput
	ToDataBoxHeavySecretResponseArrayOutputWithContext(context.Context) DataBoxHeavySecretResponseArrayOutput
}

type DataBoxHeavySecretResponseArray []DataBoxHeavySecretResponseInput

func (DataBoxHeavySecretResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxHeavySecretResponse)(nil)).Elem()
}

func (i DataBoxHeavySecretResponseArray) ToDataBoxHeavySecretResponseArrayOutput() DataBoxHeavySecretResponseArrayOutput {
	return i.ToDataBoxHeavySecretResponseArrayOutputWithContext(context.Background())
}

func (i DataBoxHeavySecretResponseArray) ToDataBoxHeavySecretResponseArrayOutputWithContext(ctx context.Context) DataBoxHeavySecretResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxHeavySecretResponseArrayOutput)
}

type DataBoxHeavySecretResponseOutput struct{ *pulumi.OutputState }

func (DataBoxHeavySecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxHeavySecretResponse)(nil)).Elem()
}

func (o DataBoxHeavySecretResponseOutput) ToDataBoxHeavySecretResponseOutput() DataBoxHeavySecretResponseOutput {
	return o
}

func (o DataBoxHeavySecretResponseOutput) ToDataBoxHeavySecretResponseOutputWithContext(ctx context.Context) DataBoxHeavySecretResponseOutput {
	return o
}

func (o DataBoxHeavySecretResponseOutput) AccountCredentialDetails() AccountCredentialDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavySecretResponse) []AccountCredentialDetailsResponse {
		return v.AccountCredentialDetails
	}).(AccountCredentialDetailsResponseArrayOutput)
}

func (o DataBoxHeavySecretResponseOutput) DevicePassword() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavySecretResponse) string { return v.DevicePassword }).(pulumi.StringOutput)
}

func (o DataBoxHeavySecretResponseOutput) DeviceSerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavySecretResponse) string { return v.DeviceSerialNumber }).(pulumi.StringOutput)
}

func (o DataBoxHeavySecretResponseOutput) EncodedValidationCertPubKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxHeavySecretResponse) string { return v.EncodedValidationCertPubKey }).(pulumi.StringOutput)
}

func (o DataBoxHeavySecretResponseOutput) NetworkConfigurations() ApplianceNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v DataBoxHeavySecretResponse) []ApplianceNetworkConfigurationResponse {
		return v.NetworkConfigurations
	}).(ApplianceNetworkConfigurationResponseArrayOutput)
}

type DataBoxHeavySecretResponseArrayOutput struct{ *pulumi.OutputState }

func (DataBoxHeavySecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxHeavySecretResponse)(nil)).Elem()
}

func (o DataBoxHeavySecretResponseArrayOutput) ToDataBoxHeavySecretResponseArrayOutput() DataBoxHeavySecretResponseArrayOutput {
	return o
}

func (o DataBoxHeavySecretResponseArrayOutput) ToDataBoxHeavySecretResponseArrayOutputWithContext(ctx context.Context) DataBoxHeavySecretResponseArrayOutput {
	return o
}

func (o DataBoxHeavySecretResponseArrayOutput) Index(i pulumi.IntInput) DataBoxHeavySecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataBoxHeavySecretResponse {
		return vs[0].([]DataBoxHeavySecretResponse)[vs[1].(int)]
	}).(DataBoxHeavySecretResponseOutput)
}

type DataBoxJobDetails struct {
	ContactDetails              ContactDetails      `pulumi:"contactDetails"`
	DataExportDetails           []DataExportDetails `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetails `pulumi:"dataImportDetails"`
	DevicePassword              *string             `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes *int                `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string              `pulumi:"jobDetailsType"`
	KeyEncryptionKey            *KeyEncryptionKey   `pulumi:"keyEncryptionKey"`
	Preferences                 *Preferences        `pulumi:"preferences"`
	ShippingAddress             *ShippingAddress    `pulumi:"shippingAddress"`
}





type DataBoxJobDetailsInput interface {
	pulumi.Input

	ToDataBoxJobDetailsOutput() DataBoxJobDetailsOutput
	ToDataBoxJobDetailsOutputWithContext(context.Context) DataBoxJobDetailsOutput
}

type DataBoxJobDetailsArgs struct {
	ContactDetails              ContactDetailsInput         `pulumi:"contactDetails"`
	DataExportDetails           DataExportDetailsArrayInput `pulumi:"dataExportDetails"`
	DataImportDetails           DataImportDetailsArrayInput `pulumi:"dataImportDetails"`
	DevicePassword              pulumi.StringPtrInput       `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput          `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput          `pulumi:"jobDetailsType"`
	KeyEncryptionKey            KeyEncryptionKeyPtrInput    `pulumi:"keyEncryptionKey"`
	Preferences                 PreferencesPtrInput         `pulumi:"preferences"`
	ShippingAddress             ShippingAddressPtrInput     `pulumi:"shippingAddress"`
}

func (DataBoxJobDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxJobDetails)(nil)).Elem()
}

func (i DataBoxJobDetailsArgs) ToDataBoxJobDetailsOutput() DataBoxJobDetailsOutput {
	return i.ToDataBoxJobDetailsOutputWithContext(context.Background())
}

func (i DataBoxJobDetailsArgs) ToDataBoxJobDetailsOutputWithContext(ctx context.Context) DataBoxJobDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxJobDetailsOutput)
}

type DataBoxJobDetailsOutput struct{ *pulumi.OutputState }

func (DataBoxJobDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxJobDetails)(nil)).Elem()
}

func (o DataBoxJobDetailsOutput) ToDataBoxJobDetailsOutput() DataBoxJobDetailsOutput {
	return o
}

func (o DataBoxJobDetailsOutput) ToDataBoxJobDetailsOutputWithContext(ctx context.Context) DataBoxJobDetailsOutput {
	return o
}

func (o DataBoxJobDetailsOutput) ContactDetails() ContactDetailsOutput {
	return o.ApplyT(func(v DataBoxJobDetails) ContactDetails { return v.ContactDetails }).(ContactDetailsOutput)
}

func (o DataBoxJobDetailsOutput) DataExportDetails() DataExportDetailsArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetails) []DataExportDetails { return v.DataExportDetails }).(DataExportDetailsArrayOutput)
}

func (o DataBoxJobDetailsOutput) DataImportDetails() DataImportDetailsArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetails) []DataImportDetails { return v.DataImportDetails }).(DataImportDetailsArrayOutput)
}

func (o DataBoxJobDetailsOutput) DevicePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetails) *string { return v.DevicePassword }).(pulumi.StringPtrOutput)
}

func (o DataBoxJobDetailsOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetails) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxJobDetailsOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxJobDetails) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxJobDetailsOutput) KeyEncryptionKey() KeyEncryptionKeyPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetails) *KeyEncryptionKey { return v.KeyEncryptionKey }).(KeyEncryptionKeyPtrOutput)
}

func (o DataBoxJobDetailsOutput) Preferences() PreferencesPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetails) *Preferences { return v.Preferences }).(PreferencesPtrOutput)
}

func (o DataBoxJobDetailsOutput) ShippingAddress() ShippingAddressPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetails) *ShippingAddress { return v.ShippingAddress }).(ShippingAddressPtrOutput)
}

type DataBoxJobDetailsResponse struct {
	Actions                     []string                          `pulumi:"actions"`
	ChainOfCustodySasKey        string                            `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponse            `pulumi:"contactDetails"`
	CopyLogDetails              []interface{}                     `pulumi:"copyLogDetails"`
	CopyProgress                []CopyProgressResponse            `pulumi:"copyProgress"`
	DataCenterCode              string                            `pulumi:"dataCenterCode"`
	DataExportDetails           []DataExportDetailsResponse       `pulumi:"dataExportDetails"`
	DataImportDetails           []DataImportDetailsResponse       `pulumi:"dataImportDetails"`
	DatacenterAddress           interface{}                       `pulumi:"datacenterAddress"`
	DeliveryPackage             PackageShippingDetailsResponse    `pulumi:"deliveryPackage"`
	DevicePassword              *string                           `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes *int                              `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              string                            `pulumi:"jobDetailsType"`
	JobStages                   []JobStagesResponse               `pulumi:"jobStages"`
	KeyEncryptionKey            *KeyEncryptionKeyResponse         `pulumi:"keyEncryptionKey"`
	LastMitigationActionOnJob   LastMitigationActionOnJobResponse `pulumi:"lastMitigationActionOnJob"`
	Preferences                 *PreferencesResponse              `pulumi:"preferences"`
	ReturnPackage               PackageShippingDetailsResponse    `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  string                            `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             *ShippingAddressResponse          `pulumi:"shippingAddress"`
}





type DataBoxJobDetailsResponseInput interface {
	pulumi.Input

	ToDataBoxJobDetailsResponseOutput() DataBoxJobDetailsResponseOutput
	ToDataBoxJobDetailsResponseOutputWithContext(context.Context) DataBoxJobDetailsResponseOutput
}

type DataBoxJobDetailsResponseArgs struct {
	Actions                     pulumi.StringArrayInput                `pulumi:"actions"`
	ChainOfCustodySasKey        pulumi.StringInput                     `pulumi:"chainOfCustodySasKey"`
	ContactDetails              ContactDetailsResponseInput            `pulumi:"contactDetails"`
	CopyLogDetails              pulumi.ArrayInput                      `pulumi:"copyLogDetails"`
	CopyProgress                CopyProgressResponseArrayInput         `pulumi:"copyProgress"`
	DataCenterCode              pulumi.StringInput                     `pulumi:"dataCenterCode"`
	DataExportDetails           DataExportDetailsResponseArrayInput    `pulumi:"dataExportDetails"`
	DataImportDetails           DataImportDetailsResponseArrayInput    `pulumi:"dataImportDetails"`
	DatacenterAddress           pulumi.Input                           `pulumi:"datacenterAddress"`
	DeliveryPackage             PackageShippingDetailsResponseInput    `pulumi:"deliveryPackage"`
	DevicePassword              pulumi.StringPtrInput                  `pulumi:"devicePassword"`
	ExpectedDataSizeInTeraBytes pulumi.IntPtrInput                     `pulumi:"expectedDataSizeInTeraBytes"`
	JobDetailsType              pulumi.StringInput                     `pulumi:"jobDetailsType"`
	JobStages                   JobStagesResponseArrayInput            `pulumi:"jobStages"`
	KeyEncryptionKey            KeyEncryptionKeyResponsePtrInput       `pulumi:"keyEncryptionKey"`
	LastMitigationActionOnJob   LastMitigationActionOnJobResponseInput `pulumi:"lastMitigationActionOnJob"`
	Preferences                 PreferencesResponsePtrInput            `pulumi:"preferences"`
	ReturnPackage               PackageShippingDetailsResponseInput    `pulumi:"returnPackage"`
	ReverseShipmentLabelSasKey  pulumi.StringInput                     `pulumi:"reverseShipmentLabelSasKey"`
	ShippingAddress             ShippingAddressResponsePtrInput        `pulumi:"shippingAddress"`
}

func (DataBoxJobDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxJobDetailsResponse)(nil)).Elem()
}

func (i DataBoxJobDetailsResponseArgs) ToDataBoxJobDetailsResponseOutput() DataBoxJobDetailsResponseOutput {
	return i.ToDataBoxJobDetailsResponseOutputWithContext(context.Background())
}

func (i DataBoxJobDetailsResponseArgs) ToDataBoxJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxJobDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxJobDetailsResponseOutput)
}

type DataBoxJobDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataBoxJobDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxJobDetailsResponse)(nil)).Elem()
}

func (o DataBoxJobDetailsResponseOutput) ToDataBoxJobDetailsResponseOutput() DataBoxJobDetailsResponseOutput {
	return o
}

func (o DataBoxJobDetailsResponseOutput) ToDataBoxJobDetailsResponseOutputWithContext(ctx context.Context) DataBoxJobDetailsResponseOutput {
	return o
}

func (o DataBoxJobDetailsResponseOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o DataBoxJobDetailsResponseOutput) ChainOfCustodySasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) string { return v.ChainOfCustodySasKey }).(pulumi.StringOutput)
}

func (o DataBoxJobDetailsResponseOutput) ContactDetails() ContactDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) ContactDetailsResponse { return v.ContactDetails }).(ContactDetailsResponseOutput)
}

func (o DataBoxJobDetailsResponseOutput) CopyLogDetails() pulumi.ArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) []interface{} { return v.CopyLogDetails }).(pulumi.ArrayOutput)
}

func (o DataBoxJobDetailsResponseOutput) CopyProgress() CopyProgressResponseArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) []CopyProgressResponse { return v.CopyProgress }).(CopyProgressResponseArrayOutput)
}

func (o DataBoxJobDetailsResponseOutput) DataCenterCode() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) string { return v.DataCenterCode }).(pulumi.StringOutput)
}

func (o DataBoxJobDetailsResponseOutput) DataExportDetails() DataExportDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) []DataExportDetailsResponse { return v.DataExportDetails }).(DataExportDetailsResponseArrayOutput)
}

func (o DataBoxJobDetailsResponseOutput) DataImportDetails() DataImportDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) []DataImportDetailsResponse { return v.DataImportDetails }).(DataImportDetailsResponseArrayOutput)
}

func (o DataBoxJobDetailsResponseOutput) DatacenterAddress() pulumi.AnyOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) interface{} { return v.DatacenterAddress }).(pulumi.AnyOutput)
}

func (o DataBoxJobDetailsResponseOutput) DeliveryPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) PackageShippingDetailsResponse { return v.DeliveryPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxJobDetailsResponseOutput) DevicePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) *string { return v.DevicePassword }).(pulumi.StringPtrOutput)
}

func (o DataBoxJobDetailsResponseOutput) ExpectedDataSizeInTeraBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) *int { return v.ExpectedDataSizeInTeraBytes }).(pulumi.IntPtrOutput)
}

func (o DataBoxJobDetailsResponseOutput) JobDetailsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) string { return v.JobDetailsType }).(pulumi.StringOutput)
}

func (o DataBoxJobDetailsResponseOutput) JobStages() JobStagesResponseArrayOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) []JobStagesResponse { return v.JobStages }).(JobStagesResponseArrayOutput)
}

func (o DataBoxJobDetailsResponseOutput) KeyEncryptionKey() KeyEncryptionKeyResponsePtrOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) *KeyEncryptionKeyResponse { return v.KeyEncryptionKey }).(KeyEncryptionKeyResponsePtrOutput)
}

func (o DataBoxJobDetailsResponseOutput) LastMitigationActionOnJob() LastMitigationActionOnJobResponseOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) LastMitigationActionOnJobResponse {
		return v.LastMitigationActionOnJob
	}).(LastMitigationActionOnJobResponseOutput)
}

func (o DataBoxJobDetailsResponseOutput) Preferences() PreferencesResponsePtrOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) *PreferencesResponse { return v.Preferences }).(PreferencesResponsePtrOutput)
}

func (o DataBoxJobDetailsResponseOutput) ReturnPackage() PackageShippingDetailsResponseOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) PackageShippingDetailsResponse { return v.ReturnPackage }).(PackageShippingDetailsResponseOutput)
}

func (o DataBoxJobDetailsResponseOutput) ReverseShipmentLabelSasKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) string { return v.ReverseShipmentLabelSasKey }).(pulumi.StringOutput)
}

func (o DataBoxJobDetailsResponseOutput) ShippingAddress() ShippingAddressResponsePtrOutput {
	return o.ApplyT(func(v DataBoxJobDetailsResponse) *ShippingAddressResponse { return v.ShippingAddress }).(ShippingAddressResponsePtrOutput)
}

type DataBoxSecretResponse struct {
	AccountCredentialDetails    []AccountCredentialDetailsResponse      `pulumi:"accountCredentialDetails"`
	DevicePassword              string                                  `pulumi:"devicePassword"`
	DeviceSerialNumber          string                                  `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey string                                  `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       []ApplianceNetworkConfigurationResponse `pulumi:"networkConfigurations"`
}





type DataBoxSecretResponseInput interface {
	pulumi.Input

	ToDataBoxSecretResponseOutput() DataBoxSecretResponseOutput
	ToDataBoxSecretResponseOutputWithContext(context.Context) DataBoxSecretResponseOutput
}

type DataBoxSecretResponseArgs struct {
	AccountCredentialDetails    AccountCredentialDetailsResponseArrayInput      `pulumi:"accountCredentialDetails"`
	DevicePassword              pulumi.StringInput                              `pulumi:"devicePassword"`
	DeviceSerialNumber          pulumi.StringInput                              `pulumi:"deviceSerialNumber"`
	EncodedValidationCertPubKey pulumi.StringInput                              `pulumi:"encodedValidationCertPubKey"`
	NetworkConfigurations       ApplianceNetworkConfigurationResponseArrayInput `pulumi:"networkConfigurations"`
}

func (DataBoxSecretResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxSecretResponse)(nil)).Elem()
}

func (i DataBoxSecretResponseArgs) ToDataBoxSecretResponseOutput() DataBoxSecretResponseOutput {
	return i.ToDataBoxSecretResponseOutputWithContext(context.Background())
}

func (i DataBoxSecretResponseArgs) ToDataBoxSecretResponseOutputWithContext(ctx context.Context) DataBoxSecretResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxSecretResponseOutput)
}





type DataBoxSecretResponseArrayInput interface {
	pulumi.Input

	ToDataBoxSecretResponseArrayOutput() DataBoxSecretResponseArrayOutput
	ToDataBoxSecretResponseArrayOutputWithContext(context.Context) DataBoxSecretResponseArrayOutput
}

type DataBoxSecretResponseArray []DataBoxSecretResponseInput

func (DataBoxSecretResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxSecretResponse)(nil)).Elem()
}

func (i DataBoxSecretResponseArray) ToDataBoxSecretResponseArrayOutput() DataBoxSecretResponseArrayOutput {
	return i.ToDataBoxSecretResponseArrayOutputWithContext(context.Background())
}

func (i DataBoxSecretResponseArray) ToDataBoxSecretResponseArrayOutputWithContext(ctx context.Context) DataBoxSecretResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataBoxSecretResponseArrayOutput)
}

type DataBoxSecretResponseOutput struct{ *pulumi.OutputState }

func (DataBoxSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBoxSecretResponse)(nil)).Elem()
}

func (o DataBoxSecretResponseOutput) ToDataBoxSecretResponseOutput() DataBoxSecretResponseOutput {
	return o
}

func (o DataBoxSecretResponseOutput) ToDataBoxSecretResponseOutputWithContext(ctx context.Context) DataBoxSecretResponseOutput {
	return o
}

func (o DataBoxSecretResponseOutput) AccountCredentialDetails() AccountCredentialDetailsResponseArrayOutput {
	return o.ApplyT(func(v DataBoxSecretResponse) []AccountCredentialDetailsResponse { return v.AccountCredentialDetails }).(AccountCredentialDetailsResponseArrayOutput)
}

func (o DataBoxSecretResponseOutput) DevicePassword() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxSecretResponse) string { return v.DevicePassword }).(pulumi.StringOutput)
}

func (o DataBoxSecretResponseOutput) DeviceSerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxSecretResponse) string { return v.DeviceSerialNumber }).(pulumi.StringOutput)
}

func (o DataBoxSecretResponseOutput) EncodedValidationCertPubKey() pulumi.StringOutput {
	return o.ApplyT(func(v DataBoxSecretResponse) string { return v.EncodedValidationCertPubKey }).(pulumi.StringOutput)
}

func (o DataBoxSecretResponseOutput) NetworkConfigurations() ApplianceNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v DataBoxSecretResponse) []ApplianceNetworkConfigurationResponse { return v.NetworkConfigurations }).(ApplianceNetworkConfigurationResponseArrayOutput)
}

type DataBoxSecretResponseArrayOutput struct{ *pulumi.OutputState }

func (DataBoxSecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataBoxSecretResponse)(nil)).Elem()
}

func (o DataBoxSecretResponseArrayOutput) ToDataBoxSecretResponseArrayOutput() DataBoxSecretResponseArrayOutput {
	return o
}

func (o DataBoxSecretResponseArrayOutput) ToDataBoxSecretResponseArrayOutputWithContext(ctx context.Context) DataBoxSecretResponseArrayOutput {
	return o
}

func (o DataBoxSecretResponseArrayOutput) Index(i pulumi.IntInput) DataBoxSecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataBoxSecretResponse {
		return vs[0].([]DataBoxSecretResponse)[vs[1].(int)]
	}).(DataBoxSecretResponseOutput)
}

type DataExportDetails struct {
	AccountDetails        interface{}           `pulumi:"accountDetails"`
	LogCollectionLevel    *string               `pulumi:"logCollectionLevel"`
	TransferConfiguration TransferConfiguration `pulumi:"transferConfiguration"`
}





type DataExportDetailsInput interface {
	pulumi.Input

	ToDataExportDetailsOutput() DataExportDetailsOutput
	ToDataExportDetailsOutputWithContext(context.Context) DataExportDetailsOutput
}

type DataExportDetailsArgs struct {
	AccountDetails        pulumi.Input               `pulumi:"accountDetails"`
	LogCollectionLevel    pulumi.StringPtrInput      `pulumi:"logCollectionLevel"`
	TransferConfiguration TransferConfigurationInput `pulumi:"transferConfiguration"`
}

func (DataExportDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataExportDetails)(nil)).Elem()
}

func (i DataExportDetailsArgs) ToDataExportDetailsOutput() DataExportDetailsOutput {
	return i.ToDataExportDetailsOutputWithContext(context.Background())
}

func (i DataExportDetailsArgs) ToDataExportDetailsOutputWithContext(ctx context.Context) DataExportDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExportDetailsOutput)
}





type DataExportDetailsArrayInput interface {
	pulumi.Input

	ToDataExportDetailsArrayOutput() DataExportDetailsArrayOutput
	ToDataExportDetailsArrayOutputWithContext(context.Context) DataExportDetailsArrayOutput
}

type DataExportDetailsArray []DataExportDetailsInput

func (DataExportDetailsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataExportDetails)(nil)).Elem()
}

func (i DataExportDetailsArray) ToDataExportDetailsArrayOutput() DataExportDetailsArrayOutput {
	return i.ToDataExportDetailsArrayOutputWithContext(context.Background())
}

func (i DataExportDetailsArray) ToDataExportDetailsArrayOutputWithContext(ctx context.Context) DataExportDetailsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExportDetailsArrayOutput)
}

type DataExportDetailsOutput struct{ *pulumi.OutputState }

func (DataExportDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataExportDetails)(nil)).Elem()
}

func (o DataExportDetailsOutput) ToDataExportDetailsOutput() DataExportDetailsOutput {
	return o
}

func (o DataExportDetailsOutput) ToDataExportDetailsOutputWithContext(ctx context.Context) DataExportDetailsOutput {
	return o
}

func (o DataExportDetailsOutput) AccountDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v DataExportDetails) interface{} { return v.AccountDetails }).(pulumi.AnyOutput)
}

func (o DataExportDetailsOutput) LogCollectionLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataExportDetails) *string { return v.LogCollectionLevel }).(pulumi.StringPtrOutput)
}

func (o DataExportDetailsOutput) TransferConfiguration() TransferConfigurationOutput {
	return o.ApplyT(func(v DataExportDetails) TransferConfiguration { return v.TransferConfiguration }).(TransferConfigurationOutput)
}

type DataExportDetailsArrayOutput struct{ *pulumi.OutputState }

func (DataExportDetailsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataExportDetails)(nil)).Elem()
}

func (o DataExportDetailsArrayOutput) ToDataExportDetailsArrayOutput() DataExportDetailsArrayOutput {
	return o
}

func (o DataExportDetailsArrayOutput) ToDataExportDetailsArrayOutputWithContext(ctx context.Context) DataExportDetailsArrayOutput {
	return o
}

func (o DataExportDetailsArrayOutput) Index(i pulumi.IntInput) DataExportDetailsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataExportDetails {
		return vs[0].([]DataExportDetails)[vs[1].(int)]
	}).(DataExportDetailsOutput)
}

type DataExportDetailsResponse struct {
	AccountDetails        interface{}                   `pulumi:"accountDetails"`
	LogCollectionLevel    *string                       `pulumi:"logCollectionLevel"`
	TransferConfiguration TransferConfigurationResponse `pulumi:"transferConfiguration"`
}





type DataExportDetailsResponseInput interface {
	pulumi.Input

	ToDataExportDetailsResponseOutput() DataExportDetailsResponseOutput
	ToDataExportDetailsResponseOutputWithContext(context.Context) DataExportDetailsResponseOutput
}

type DataExportDetailsResponseArgs struct {
	AccountDetails        pulumi.Input                       `pulumi:"accountDetails"`
	LogCollectionLevel    pulumi.StringPtrInput              `pulumi:"logCollectionLevel"`
	TransferConfiguration TransferConfigurationResponseInput `pulumi:"transferConfiguration"`
}

func (DataExportDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataExportDetailsResponse)(nil)).Elem()
}

func (i DataExportDetailsResponseArgs) ToDataExportDetailsResponseOutput() DataExportDetailsResponseOutput {
	return i.ToDataExportDetailsResponseOutputWithContext(context.Background())
}

func (i DataExportDetailsResponseArgs) ToDataExportDetailsResponseOutputWithContext(ctx context.Context) DataExportDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExportDetailsResponseOutput)
}





type DataExportDetailsResponseArrayInput interface {
	pulumi.Input

	ToDataExportDetailsResponseArrayOutput() DataExportDetailsResponseArrayOutput
	ToDataExportDetailsResponseArrayOutputWithContext(context.Context) DataExportDetailsResponseArrayOutput
}

type DataExportDetailsResponseArray []DataExportDetailsResponseInput

func (DataExportDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataExportDetailsResponse)(nil)).Elem()
}

func (i DataExportDetailsResponseArray) ToDataExportDetailsResponseArrayOutput() DataExportDetailsResponseArrayOutput {
	return i.ToDataExportDetailsResponseArrayOutputWithContext(context.Background())
}

func (i DataExportDetailsResponseArray) ToDataExportDetailsResponseArrayOutputWithContext(ctx context.Context) DataExportDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExportDetailsResponseArrayOutput)
}

type DataExportDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataExportDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataExportDetailsResponse)(nil)).Elem()
}

func (o DataExportDetailsResponseOutput) ToDataExportDetailsResponseOutput() DataExportDetailsResponseOutput {
	return o
}

func (o DataExportDetailsResponseOutput) ToDataExportDetailsResponseOutputWithContext(ctx context.Context) DataExportDetailsResponseOutput {
	return o
}

func (o DataExportDetailsResponseOutput) AccountDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v DataExportDetailsResponse) interface{} { return v.AccountDetails }).(pulumi.AnyOutput)
}

func (o DataExportDetailsResponseOutput) LogCollectionLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataExportDetailsResponse) *string { return v.LogCollectionLevel }).(pulumi.StringPtrOutput)
}

func (o DataExportDetailsResponseOutput) TransferConfiguration() TransferConfigurationResponseOutput {
	return o.ApplyT(func(v DataExportDetailsResponse) TransferConfigurationResponse { return v.TransferConfiguration }).(TransferConfigurationResponseOutput)
}

type DataExportDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (DataExportDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataExportDetailsResponse)(nil)).Elem()
}

func (o DataExportDetailsResponseArrayOutput) ToDataExportDetailsResponseArrayOutput() DataExportDetailsResponseArrayOutput {
	return o
}

func (o DataExportDetailsResponseArrayOutput) ToDataExportDetailsResponseArrayOutputWithContext(ctx context.Context) DataExportDetailsResponseArrayOutput {
	return o
}

func (o DataExportDetailsResponseArrayOutput) Index(i pulumi.IntInput) DataExportDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataExportDetailsResponse {
		return vs[0].([]DataExportDetailsResponse)[vs[1].(int)]
	}).(DataExportDetailsResponseOutput)
}

type DataImportDetails struct {
	AccountDetails     interface{} `pulumi:"accountDetails"`
	LogCollectionLevel *string     `pulumi:"logCollectionLevel"`
}





type DataImportDetailsInput interface {
	pulumi.Input

	ToDataImportDetailsOutput() DataImportDetailsOutput
	ToDataImportDetailsOutputWithContext(context.Context) DataImportDetailsOutput
}

type DataImportDetailsArgs struct {
	AccountDetails     pulumi.Input          `pulumi:"accountDetails"`
	LogCollectionLevel pulumi.StringPtrInput `pulumi:"logCollectionLevel"`
}

func (DataImportDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataImportDetails)(nil)).Elem()
}

func (i DataImportDetailsArgs) ToDataImportDetailsOutput() DataImportDetailsOutput {
	return i.ToDataImportDetailsOutputWithContext(context.Background())
}

func (i DataImportDetailsArgs) ToDataImportDetailsOutputWithContext(ctx context.Context) DataImportDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataImportDetailsOutput)
}





type DataImportDetailsArrayInput interface {
	pulumi.Input

	ToDataImportDetailsArrayOutput() DataImportDetailsArrayOutput
	ToDataImportDetailsArrayOutputWithContext(context.Context) DataImportDetailsArrayOutput
}

type DataImportDetailsArray []DataImportDetailsInput

func (DataImportDetailsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataImportDetails)(nil)).Elem()
}

func (i DataImportDetailsArray) ToDataImportDetailsArrayOutput() DataImportDetailsArrayOutput {
	return i.ToDataImportDetailsArrayOutputWithContext(context.Background())
}

func (i DataImportDetailsArray) ToDataImportDetailsArrayOutputWithContext(ctx context.Context) DataImportDetailsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataImportDetailsArrayOutput)
}

type DataImportDetailsOutput struct{ *pulumi.OutputState }

func (DataImportDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataImportDetails)(nil)).Elem()
}

func (o DataImportDetailsOutput) ToDataImportDetailsOutput() DataImportDetailsOutput {
	return o
}

func (o DataImportDetailsOutput) ToDataImportDetailsOutputWithContext(ctx context.Context) DataImportDetailsOutput {
	return o
}

func (o DataImportDetailsOutput) AccountDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v DataImportDetails) interface{} { return v.AccountDetails }).(pulumi.AnyOutput)
}

func (o DataImportDetailsOutput) LogCollectionLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataImportDetails) *string { return v.LogCollectionLevel }).(pulumi.StringPtrOutput)
}

type DataImportDetailsArrayOutput struct{ *pulumi.OutputState }

func (DataImportDetailsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataImportDetails)(nil)).Elem()
}

func (o DataImportDetailsArrayOutput) ToDataImportDetailsArrayOutput() DataImportDetailsArrayOutput {
	return o
}

func (o DataImportDetailsArrayOutput) ToDataImportDetailsArrayOutputWithContext(ctx context.Context) DataImportDetailsArrayOutput {
	return o
}

func (o DataImportDetailsArrayOutput) Index(i pulumi.IntInput) DataImportDetailsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataImportDetails {
		return vs[0].([]DataImportDetails)[vs[1].(int)]
	}).(DataImportDetailsOutput)
}

type DataImportDetailsResponse struct {
	AccountDetails     interface{} `pulumi:"accountDetails"`
	LogCollectionLevel *string     `pulumi:"logCollectionLevel"`
}





type DataImportDetailsResponseInput interface {
	pulumi.Input

	ToDataImportDetailsResponseOutput() DataImportDetailsResponseOutput
	ToDataImportDetailsResponseOutputWithContext(context.Context) DataImportDetailsResponseOutput
}

type DataImportDetailsResponseArgs struct {
	AccountDetails     pulumi.Input          `pulumi:"accountDetails"`
	LogCollectionLevel pulumi.StringPtrInput `pulumi:"logCollectionLevel"`
}

func (DataImportDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataImportDetailsResponse)(nil)).Elem()
}

func (i DataImportDetailsResponseArgs) ToDataImportDetailsResponseOutput() DataImportDetailsResponseOutput {
	return i.ToDataImportDetailsResponseOutputWithContext(context.Background())
}

func (i DataImportDetailsResponseArgs) ToDataImportDetailsResponseOutputWithContext(ctx context.Context) DataImportDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataImportDetailsResponseOutput)
}





type DataImportDetailsResponseArrayInput interface {
	pulumi.Input

	ToDataImportDetailsResponseArrayOutput() DataImportDetailsResponseArrayOutput
	ToDataImportDetailsResponseArrayOutputWithContext(context.Context) DataImportDetailsResponseArrayOutput
}

type DataImportDetailsResponseArray []DataImportDetailsResponseInput

func (DataImportDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataImportDetailsResponse)(nil)).Elem()
}

func (i DataImportDetailsResponseArray) ToDataImportDetailsResponseArrayOutput() DataImportDetailsResponseArrayOutput {
	return i.ToDataImportDetailsResponseArrayOutputWithContext(context.Background())
}

func (i DataImportDetailsResponseArray) ToDataImportDetailsResponseArrayOutputWithContext(ctx context.Context) DataImportDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataImportDetailsResponseArrayOutput)
}

type DataImportDetailsResponseOutput struct{ *pulumi.OutputState }

func (DataImportDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataImportDetailsResponse)(nil)).Elem()
}

func (o DataImportDetailsResponseOutput) ToDataImportDetailsResponseOutput() DataImportDetailsResponseOutput {
	return o
}

func (o DataImportDetailsResponseOutput) ToDataImportDetailsResponseOutputWithContext(ctx context.Context) DataImportDetailsResponseOutput {
	return o
}

func (o DataImportDetailsResponseOutput) AccountDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v DataImportDetailsResponse) interface{} { return v.AccountDetails }).(pulumi.AnyOutput)
}

func (o DataImportDetailsResponseOutput) LogCollectionLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataImportDetailsResponse) *string { return v.LogCollectionLevel }).(pulumi.StringPtrOutput)
}

type DataImportDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (DataImportDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataImportDetailsResponse)(nil)).Elem()
}

func (o DataImportDetailsResponseArrayOutput) ToDataImportDetailsResponseArrayOutput() DataImportDetailsResponseArrayOutput {
	return o
}

func (o DataImportDetailsResponseArrayOutput) ToDataImportDetailsResponseArrayOutputWithContext(ctx context.Context) DataImportDetailsResponseArrayOutput {
	return o
}

func (o DataImportDetailsResponseArrayOutput) Index(i pulumi.IntInput) DataImportDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataImportDetailsResponse {
		return vs[0].([]DataImportDetailsResponse)[vs[1].(int)]
	}).(DataImportDetailsResponseOutput)
}

type DataboxJobSecretsResponse struct {
	DcAccessSecurityCode DcAccessSecurityCodeResponse `pulumi:"dcAccessSecurityCode"`
	Error                CloudErrorResponse           `pulumi:"error"`
	JobSecretsType       string                       `pulumi:"jobSecretsType"`
	PodSecrets           []DataBoxSecretResponse      `pulumi:"podSecrets"`
}





type DataboxJobSecretsResponseInput interface {
	pulumi.Input

	ToDataboxJobSecretsResponseOutput() DataboxJobSecretsResponseOutput
	ToDataboxJobSecretsResponseOutputWithContext(context.Context) DataboxJobSecretsResponseOutput
}

type DataboxJobSecretsResponseArgs struct {
	DcAccessSecurityCode DcAccessSecurityCodeResponseInput `pulumi:"dcAccessSecurityCode"`
	Error                CloudErrorResponseInput           `pulumi:"error"`
	JobSecretsType       pulumi.StringInput                `pulumi:"jobSecretsType"`
	PodSecrets           DataBoxSecretResponseArrayInput   `pulumi:"podSecrets"`
}

func (DataboxJobSecretsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataboxJobSecretsResponse)(nil)).Elem()
}

func (i DataboxJobSecretsResponseArgs) ToDataboxJobSecretsResponseOutput() DataboxJobSecretsResponseOutput {
	return i.ToDataboxJobSecretsResponseOutputWithContext(context.Background())
}

func (i DataboxJobSecretsResponseArgs) ToDataboxJobSecretsResponseOutputWithContext(ctx context.Context) DataboxJobSecretsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataboxJobSecretsResponseOutput)
}

type DataboxJobSecretsResponseOutput struct{ *pulumi.OutputState }

func (DataboxJobSecretsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataboxJobSecretsResponse)(nil)).Elem()
}

func (o DataboxJobSecretsResponseOutput) ToDataboxJobSecretsResponseOutput() DataboxJobSecretsResponseOutput {
	return o
}

func (o DataboxJobSecretsResponseOutput) ToDataboxJobSecretsResponseOutputWithContext(ctx context.Context) DataboxJobSecretsResponseOutput {
	return o
}

func (o DataboxJobSecretsResponseOutput) DcAccessSecurityCode() DcAccessSecurityCodeResponseOutput {
	return o.ApplyT(func(v DataboxJobSecretsResponse) DcAccessSecurityCodeResponse { return v.DcAccessSecurityCode }).(DcAccessSecurityCodeResponseOutput)
}

func (o DataboxJobSecretsResponseOutput) Error() CloudErrorResponseOutput {
	return o.ApplyT(func(v DataboxJobSecretsResponse) CloudErrorResponse { return v.Error }).(CloudErrorResponseOutput)
}

func (o DataboxJobSecretsResponseOutput) JobSecretsType() pulumi.StringOutput {
	return o.ApplyT(func(v DataboxJobSecretsResponse) string { return v.JobSecretsType }).(pulumi.StringOutput)
}

func (o DataboxJobSecretsResponseOutput) PodSecrets() DataBoxSecretResponseArrayOutput {
	return o.ApplyT(func(v DataboxJobSecretsResponse) []DataBoxSecretResponse { return v.PodSecrets }).(DataBoxSecretResponseArrayOutput)
}

type DatacenterAddressInstructionResponseResponse struct {
	CommunicationInstruction           string   `pulumi:"communicationInstruction"`
	DataCenterAzureLocation            string   `pulumi:"dataCenterAzureLocation"`
	DatacenterAddressType              string   `pulumi:"datacenterAddressType"`
	SupportedCarriersForReturnShipment []string `pulumi:"supportedCarriersForReturnShipment"`
}





type DatacenterAddressInstructionResponseResponseInput interface {
	pulumi.Input

	ToDatacenterAddressInstructionResponseResponseOutput() DatacenterAddressInstructionResponseResponseOutput
	ToDatacenterAddressInstructionResponseResponseOutputWithContext(context.Context) DatacenterAddressInstructionResponseResponseOutput
}

type DatacenterAddressInstructionResponseResponseArgs struct {
	CommunicationInstruction           pulumi.StringInput      `pulumi:"communicationInstruction"`
	DataCenterAzureLocation            pulumi.StringInput      `pulumi:"dataCenterAzureLocation"`
	DatacenterAddressType              pulumi.StringInput      `pulumi:"datacenterAddressType"`
	SupportedCarriersForReturnShipment pulumi.StringArrayInput `pulumi:"supportedCarriersForReturnShipment"`
}

func (DatacenterAddressInstructionResponseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatacenterAddressInstructionResponseResponse)(nil)).Elem()
}

func (i DatacenterAddressInstructionResponseResponseArgs) ToDatacenterAddressInstructionResponseResponseOutput() DatacenterAddressInstructionResponseResponseOutput {
	return i.ToDatacenterAddressInstructionResponseResponseOutputWithContext(context.Background())
}

func (i DatacenterAddressInstructionResponseResponseArgs) ToDatacenterAddressInstructionResponseResponseOutputWithContext(ctx context.Context) DatacenterAddressInstructionResponseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatacenterAddressInstructionResponseResponseOutput)
}

type DatacenterAddressInstructionResponseResponseOutput struct{ *pulumi.OutputState }

func (DatacenterAddressInstructionResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatacenterAddressInstructionResponseResponse)(nil)).Elem()
}

func (o DatacenterAddressInstructionResponseResponseOutput) ToDatacenterAddressInstructionResponseResponseOutput() DatacenterAddressInstructionResponseResponseOutput {
	return o
}

func (o DatacenterAddressInstructionResponseResponseOutput) ToDatacenterAddressInstructionResponseResponseOutputWithContext(ctx context.Context) DatacenterAddressInstructionResponseResponseOutput {
	return o
}

func (o DatacenterAddressInstructionResponseResponseOutput) CommunicationInstruction() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressInstructionResponseResponse) string { return v.CommunicationInstruction }).(pulumi.StringOutput)
}

func (o DatacenterAddressInstructionResponseResponseOutput) DataCenterAzureLocation() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressInstructionResponseResponse) string { return v.DataCenterAzureLocation }).(pulumi.StringOutput)
}

func (o DatacenterAddressInstructionResponseResponseOutput) DatacenterAddressType() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressInstructionResponseResponse) string { return v.DatacenterAddressType }).(pulumi.StringOutput)
}

func (o DatacenterAddressInstructionResponseResponseOutput) SupportedCarriersForReturnShipment() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatacenterAddressInstructionResponseResponse) []string {
		return v.SupportedCarriersForReturnShipment
	}).(pulumi.StringArrayOutput)
}

type DatacenterAddressLocationResponseResponse struct {
	AdditionalShippingInformation      string   `pulumi:"additionalShippingInformation"`
	AddressType                        string   `pulumi:"addressType"`
	City                               string   `pulumi:"city"`
	Company                            string   `pulumi:"company"`
	ContactPersonName                  string   `pulumi:"contactPersonName"`
	Country                            string   `pulumi:"country"`
	DataCenterAzureLocation            string   `pulumi:"dataCenterAzureLocation"`
	DatacenterAddressType              string   `pulumi:"datacenterAddressType"`
	Phone                              string   `pulumi:"phone"`
	PhoneExtension                     string   `pulumi:"phoneExtension"`
	State                              string   `pulumi:"state"`
	Street1                            string   `pulumi:"street1"`
	Street2                            string   `pulumi:"street2"`
	Street3                            string   `pulumi:"street3"`
	SupportedCarriersForReturnShipment []string `pulumi:"supportedCarriersForReturnShipment"`
	Zip                                string   `pulumi:"zip"`
}





type DatacenterAddressLocationResponseResponseInput interface {
	pulumi.Input

	ToDatacenterAddressLocationResponseResponseOutput() DatacenterAddressLocationResponseResponseOutput
	ToDatacenterAddressLocationResponseResponseOutputWithContext(context.Context) DatacenterAddressLocationResponseResponseOutput
}

type DatacenterAddressLocationResponseResponseArgs struct {
	AdditionalShippingInformation      pulumi.StringInput      `pulumi:"additionalShippingInformation"`
	AddressType                        pulumi.StringInput      `pulumi:"addressType"`
	City                               pulumi.StringInput      `pulumi:"city"`
	Company                            pulumi.StringInput      `pulumi:"company"`
	ContactPersonName                  pulumi.StringInput      `pulumi:"contactPersonName"`
	Country                            pulumi.StringInput      `pulumi:"country"`
	DataCenterAzureLocation            pulumi.StringInput      `pulumi:"dataCenterAzureLocation"`
	DatacenterAddressType              pulumi.StringInput      `pulumi:"datacenterAddressType"`
	Phone                              pulumi.StringInput      `pulumi:"phone"`
	PhoneExtension                     pulumi.StringInput      `pulumi:"phoneExtension"`
	State                              pulumi.StringInput      `pulumi:"state"`
	Street1                            pulumi.StringInput      `pulumi:"street1"`
	Street2                            pulumi.StringInput      `pulumi:"street2"`
	Street3                            pulumi.StringInput      `pulumi:"street3"`
	SupportedCarriersForReturnShipment pulumi.StringArrayInput `pulumi:"supportedCarriersForReturnShipment"`
	Zip                                pulumi.StringInput      `pulumi:"zip"`
}

func (DatacenterAddressLocationResponseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatacenterAddressLocationResponseResponse)(nil)).Elem()
}

func (i DatacenterAddressLocationResponseResponseArgs) ToDatacenterAddressLocationResponseResponseOutput() DatacenterAddressLocationResponseResponseOutput {
	return i.ToDatacenterAddressLocationResponseResponseOutputWithContext(context.Background())
}

func (i DatacenterAddressLocationResponseResponseArgs) ToDatacenterAddressLocationResponseResponseOutputWithContext(ctx context.Context) DatacenterAddressLocationResponseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatacenterAddressLocationResponseResponseOutput)
}

type DatacenterAddressLocationResponseResponseOutput struct{ *pulumi.OutputState }

func (DatacenterAddressLocationResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatacenterAddressLocationResponseResponse)(nil)).Elem()
}

func (o DatacenterAddressLocationResponseResponseOutput) ToDatacenterAddressLocationResponseResponseOutput() DatacenterAddressLocationResponseResponseOutput {
	return o
}

func (o DatacenterAddressLocationResponseResponseOutput) ToDatacenterAddressLocationResponseResponseOutputWithContext(ctx context.Context) DatacenterAddressLocationResponseResponseOutput {
	return o
}

func (o DatacenterAddressLocationResponseResponseOutput) AdditionalShippingInformation() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.AdditionalShippingInformation }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.AddressType }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.City }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) Company() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.Company }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) ContactPersonName() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.ContactPersonName }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.Country }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) DataCenterAzureLocation() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.DataCenterAzureLocation }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) DatacenterAddressType() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.DatacenterAddressType }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.Phone }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) PhoneExtension() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.PhoneExtension }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) Street1() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.Street1 }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) Street2() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.Street2 }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) Street3() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.Street3 }).(pulumi.StringOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) SupportedCarriersForReturnShipment() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) []string {
		return v.SupportedCarriersForReturnShipment
	}).(pulumi.StringArrayOutput)
}

func (o DatacenterAddressLocationResponseResponseOutput) Zip() pulumi.StringOutput {
	return o.ApplyT(func(v DatacenterAddressLocationResponseResponse) string { return v.Zip }).(pulumi.StringOutput)
}

type DcAccessSecurityCodeResponse struct {
	ForwardDCAccessCode *string `pulumi:"forwardDCAccessCode"`
	ReverseDCAccessCode *string `pulumi:"reverseDCAccessCode"`
}





type DcAccessSecurityCodeResponseInput interface {
	pulumi.Input

	ToDcAccessSecurityCodeResponseOutput() DcAccessSecurityCodeResponseOutput
	ToDcAccessSecurityCodeResponseOutputWithContext(context.Context) DcAccessSecurityCodeResponseOutput
}

type DcAccessSecurityCodeResponseArgs struct {
	ForwardDCAccessCode pulumi.StringPtrInput `pulumi:"forwardDCAccessCode"`
	ReverseDCAccessCode pulumi.StringPtrInput `pulumi:"reverseDCAccessCode"`
}

func (DcAccessSecurityCodeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DcAccessSecurityCodeResponse)(nil)).Elem()
}

func (i DcAccessSecurityCodeResponseArgs) ToDcAccessSecurityCodeResponseOutput() DcAccessSecurityCodeResponseOutput {
	return i.ToDcAccessSecurityCodeResponseOutputWithContext(context.Background())
}

func (i DcAccessSecurityCodeResponseArgs) ToDcAccessSecurityCodeResponseOutputWithContext(ctx context.Context) DcAccessSecurityCodeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DcAccessSecurityCodeResponseOutput)
}

type DcAccessSecurityCodeResponseOutput struct{ *pulumi.OutputState }

func (DcAccessSecurityCodeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DcAccessSecurityCodeResponse)(nil)).Elem()
}

func (o DcAccessSecurityCodeResponseOutput) ToDcAccessSecurityCodeResponseOutput() DcAccessSecurityCodeResponseOutput {
	return o
}

func (o DcAccessSecurityCodeResponseOutput) ToDcAccessSecurityCodeResponseOutputWithContext(ctx context.Context) DcAccessSecurityCodeResponseOutput {
	return o
}

func (o DcAccessSecurityCodeResponseOutput) ForwardDCAccessCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DcAccessSecurityCodeResponse) *string { return v.ForwardDCAccessCode }).(pulumi.StringPtrOutput)
}

func (o DcAccessSecurityCodeResponseOutput) ReverseDCAccessCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DcAccessSecurityCodeResponse) *string { return v.ReverseDCAccessCode }).(pulumi.StringPtrOutput)
}

type DiskSecretResponse struct {
	BitLockerKey     string `pulumi:"bitLockerKey"`
	DiskSerialNumber string `pulumi:"diskSerialNumber"`
}





type DiskSecretResponseInput interface {
	pulumi.Input

	ToDiskSecretResponseOutput() DiskSecretResponseOutput
	ToDiskSecretResponseOutputWithContext(context.Context) DiskSecretResponseOutput
}

type DiskSecretResponseArgs struct {
	BitLockerKey     pulumi.StringInput `pulumi:"bitLockerKey"`
	DiskSerialNumber pulumi.StringInput `pulumi:"diskSerialNumber"`
}

func (DiskSecretResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSecretResponse)(nil)).Elem()
}

func (i DiskSecretResponseArgs) ToDiskSecretResponseOutput() DiskSecretResponseOutput {
	return i.ToDiskSecretResponseOutputWithContext(context.Background())
}

func (i DiskSecretResponseArgs) ToDiskSecretResponseOutputWithContext(ctx context.Context) DiskSecretResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSecretResponseOutput)
}





type DiskSecretResponseArrayInput interface {
	pulumi.Input

	ToDiskSecretResponseArrayOutput() DiskSecretResponseArrayOutput
	ToDiskSecretResponseArrayOutputWithContext(context.Context) DiskSecretResponseArrayOutput
}

type DiskSecretResponseArray []DiskSecretResponseInput

func (DiskSecretResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskSecretResponse)(nil)).Elem()
}

func (i DiskSecretResponseArray) ToDiskSecretResponseArrayOutput() DiskSecretResponseArrayOutput {
	return i.ToDiskSecretResponseArrayOutputWithContext(context.Background())
}

func (i DiskSecretResponseArray) ToDiskSecretResponseArrayOutputWithContext(ctx context.Context) DiskSecretResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSecretResponseArrayOutput)
}

type DiskSecretResponseOutput struct{ *pulumi.OutputState }

func (DiskSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSecretResponse)(nil)).Elem()
}

func (o DiskSecretResponseOutput) ToDiskSecretResponseOutput() DiskSecretResponseOutput {
	return o
}

func (o DiskSecretResponseOutput) ToDiskSecretResponseOutputWithContext(ctx context.Context) DiskSecretResponseOutput {
	return o
}

func (o DiskSecretResponseOutput) BitLockerKey() pulumi.StringOutput {
	return o.ApplyT(func(v DiskSecretResponse) string { return v.BitLockerKey }).(pulumi.StringOutput)
}

func (o DiskSecretResponseOutput) DiskSerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v DiskSecretResponse) string { return v.DiskSerialNumber }).(pulumi.StringOutput)
}

type DiskSecretResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskSecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskSecretResponse)(nil)).Elem()
}

func (o DiskSecretResponseArrayOutput) ToDiskSecretResponseArrayOutput() DiskSecretResponseArrayOutput {
	return o
}

func (o DiskSecretResponseArrayOutput) ToDiskSecretResponseArrayOutputWithContext(ctx context.Context) DiskSecretResponseArrayOutput {
	return o
}

func (o DiskSecretResponseArrayOutput) Index(i pulumi.IntInput) DiskSecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskSecretResponse {
		return vs[0].([]DiskSecretResponse)[vs[1].(int)]
	}).(DiskSecretResponseOutput)
}

type EncryptionPreferences struct {
	DoubleEncryption *string `pulumi:"doubleEncryption"`
}





type EncryptionPreferencesInput interface {
	pulumi.Input

	ToEncryptionPreferencesOutput() EncryptionPreferencesOutput
	ToEncryptionPreferencesOutputWithContext(context.Context) EncryptionPreferencesOutput
}

type EncryptionPreferencesArgs struct {
	DoubleEncryption pulumi.StringPtrInput `pulumi:"doubleEncryption"`
}

func (EncryptionPreferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPreferences)(nil)).Elem()
}

func (i EncryptionPreferencesArgs) ToEncryptionPreferencesOutput() EncryptionPreferencesOutput {
	return i.ToEncryptionPreferencesOutputWithContext(context.Background())
}

func (i EncryptionPreferencesArgs) ToEncryptionPreferencesOutputWithContext(ctx context.Context) EncryptionPreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesOutput)
}

func (i EncryptionPreferencesArgs) ToEncryptionPreferencesPtrOutput() EncryptionPreferencesPtrOutput {
	return i.ToEncryptionPreferencesPtrOutputWithContext(context.Background())
}

func (i EncryptionPreferencesArgs) ToEncryptionPreferencesPtrOutputWithContext(ctx context.Context) EncryptionPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesOutput).ToEncryptionPreferencesPtrOutputWithContext(ctx)
}









type EncryptionPreferencesPtrInput interface {
	pulumi.Input

	ToEncryptionPreferencesPtrOutput() EncryptionPreferencesPtrOutput
	ToEncryptionPreferencesPtrOutputWithContext(context.Context) EncryptionPreferencesPtrOutput
}

type encryptionPreferencesPtrType EncryptionPreferencesArgs

func EncryptionPreferencesPtr(v *EncryptionPreferencesArgs) EncryptionPreferencesPtrInput {
	return (*encryptionPreferencesPtrType)(v)
}

func (*encryptionPreferencesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPreferences)(nil)).Elem()
}

func (i *encryptionPreferencesPtrType) ToEncryptionPreferencesPtrOutput() EncryptionPreferencesPtrOutput {
	return i.ToEncryptionPreferencesPtrOutputWithContext(context.Background())
}

func (i *encryptionPreferencesPtrType) ToEncryptionPreferencesPtrOutputWithContext(ctx context.Context) EncryptionPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesPtrOutput)
}

type EncryptionPreferencesOutput struct{ *pulumi.OutputState }

func (EncryptionPreferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPreferences)(nil)).Elem()
}

func (o EncryptionPreferencesOutput) ToEncryptionPreferencesOutput() EncryptionPreferencesOutput {
	return o
}

func (o EncryptionPreferencesOutput) ToEncryptionPreferencesOutputWithContext(ctx context.Context) EncryptionPreferencesOutput {
	return o
}

func (o EncryptionPreferencesOutput) ToEncryptionPreferencesPtrOutput() EncryptionPreferencesPtrOutput {
	return o.ToEncryptionPreferencesPtrOutputWithContext(context.Background())
}

func (o EncryptionPreferencesOutput) ToEncryptionPreferencesPtrOutputWithContext(ctx context.Context) EncryptionPreferencesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPreferences) *EncryptionPreferences {
		return &v
	}).(EncryptionPreferencesPtrOutput)
}

func (o EncryptionPreferencesOutput) DoubleEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionPreferences) *string { return v.DoubleEncryption }).(pulumi.StringPtrOutput)
}

type EncryptionPreferencesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPreferencesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPreferences)(nil)).Elem()
}

func (o EncryptionPreferencesPtrOutput) ToEncryptionPreferencesPtrOutput() EncryptionPreferencesPtrOutput {
	return o
}

func (o EncryptionPreferencesPtrOutput) ToEncryptionPreferencesPtrOutputWithContext(ctx context.Context) EncryptionPreferencesPtrOutput {
	return o
}

func (o EncryptionPreferencesPtrOutput) Elem() EncryptionPreferencesOutput {
	return o.ApplyT(func(v *EncryptionPreferences) EncryptionPreferences {
		if v != nil {
			return *v
		}
		var ret EncryptionPreferences
		return ret
	}).(EncryptionPreferencesOutput)
}

func (o EncryptionPreferencesPtrOutput) DoubleEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPreferences) *string {
		if v == nil {
			return nil
		}
		return v.DoubleEncryption
	}).(pulumi.StringPtrOutput)
}

type EncryptionPreferencesResponse struct {
	DoubleEncryption *string `pulumi:"doubleEncryption"`
}





type EncryptionPreferencesResponseInput interface {
	pulumi.Input

	ToEncryptionPreferencesResponseOutput() EncryptionPreferencesResponseOutput
	ToEncryptionPreferencesResponseOutputWithContext(context.Context) EncryptionPreferencesResponseOutput
}

type EncryptionPreferencesResponseArgs struct {
	DoubleEncryption pulumi.StringPtrInput `pulumi:"doubleEncryption"`
}

func (EncryptionPreferencesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPreferencesResponse)(nil)).Elem()
}

func (i EncryptionPreferencesResponseArgs) ToEncryptionPreferencesResponseOutput() EncryptionPreferencesResponseOutput {
	return i.ToEncryptionPreferencesResponseOutputWithContext(context.Background())
}

func (i EncryptionPreferencesResponseArgs) ToEncryptionPreferencesResponseOutputWithContext(ctx context.Context) EncryptionPreferencesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesResponseOutput)
}

func (i EncryptionPreferencesResponseArgs) ToEncryptionPreferencesResponsePtrOutput() EncryptionPreferencesResponsePtrOutput {
	return i.ToEncryptionPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionPreferencesResponseArgs) ToEncryptionPreferencesResponsePtrOutputWithContext(ctx context.Context) EncryptionPreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesResponseOutput).ToEncryptionPreferencesResponsePtrOutputWithContext(ctx)
}









type EncryptionPreferencesResponsePtrInput interface {
	pulumi.Input

	ToEncryptionPreferencesResponsePtrOutput() EncryptionPreferencesResponsePtrOutput
	ToEncryptionPreferencesResponsePtrOutputWithContext(context.Context) EncryptionPreferencesResponsePtrOutput
}

type encryptionPreferencesResponsePtrType EncryptionPreferencesResponseArgs

func EncryptionPreferencesResponsePtr(v *EncryptionPreferencesResponseArgs) EncryptionPreferencesResponsePtrInput {
	return (*encryptionPreferencesResponsePtrType)(v)
}

func (*encryptionPreferencesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPreferencesResponse)(nil)).Elem()
}

func (i *encryptionPreferencesResponsePtrType) ToEncryptionPreferencesResponsePtrOutput() EncryptionPreferencesResponsePtrOutput {
	return i.ToEncryptionPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionPreferencesResponsePtrType) ToEncryptionPreferencesResponsePtrOutputWithContext(ctx context.Context) EncryptionPreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPreferencesResponsePtrOutput)
}

type EncryptionPreferencesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionPreferencesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPreferencesResponse)(nil)).Elem()
}

func (o EncryptionPreferencesResponseOutput) ToEncryptionPreferencesResponseOutput() EncryptionPreferencesResponseOutput {
	return o
}

func (o EncryptionPreferencesResponseOutput) ToEncryptionPreferencesResponseOutputWithContext(ctx context.Context) EncryptionPreferencesResponseOutput {
	return o
}

func (o EncryptionPreferencesResponseOutput) ToEncryptionPreferencesResponsePtrOutput() EncryptionPreferencesResponsePtrOutput {
	return o.ToEncryptionPreferencesResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionPreferencesResponseOutput) ToEncryptionPreferencesResponsePtrOutputWithContext(ctx context.Context) EncryptionPreferencesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPreferencesResponse) *EncryptionPreferencesResponse {
		return &v
	}).(EncryptionPreferencesResponsePtrOutput)
}

func (o EncryptionPreferencesResponseOutput) DoubleEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionPreferencesResponse) *string { return v.DoubleEncryption }).(pulumi.StringPtrOutput)
}

type EncryptionPreferencesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionPreferencesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPreferencesResponse)(nil)).Elem()
}

func (o EncryptionPreferencesResponsePtrOutput) ToEncryptionPreferencesResponsePtrOutput() EncryptionPreferencesResponsePtrOutput {
	return o
}

func (o EncryptionPreferencesResponsePtrOutput) ToEncryptionPreferencesResponsePtrOutputWithContext(ctx context.Context) EncryptionPreferencesResponsePtrOutput {
	return o
}

func (o EncryptionPreferencesResponsePtrOutput) Elem() EncryptionPreferencesResponseOutput {
	return o.ApplyT(func(v *EncryptionPreferencesResponse) EncryptionPreferencesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionPreferencesResponse
		return ret
	}).(EncryptionPreferencesResponseOutput)
}

func (o EncryptionPreferencesResponsePtrOutput) DoubleEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPreferencesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DoubleEncryption
	}).(pulumi.StringPtrOutput)
}

type ExportDiskDetailsResponse struct {
	BackupManifestCloudPath string `pulumi:"backupManifestCloudPath"`
	ManifestFile            string `pulumi:"manifestFile"`
	ManifestHash            string `pulumi:"manifestHash"`
}





type ExportDiskDetailsResponseInput interface {
	pulumi.Input

	ToExportDiskDetailsResponseOutput() ExportDiskDetailsResponseOutput
	ToExportDiskDetailsResponseOutputWithContext(context.Context) ExportDiskDetailsResponseOutput
}

type ExportDiskDetailsResponseArgs struct {
	BackupManifestCloudPath pulumi.StringInput `pulumi:"backupManifestCloudPath"`
	ManifestFile            pulumi.StringInput `pulumi:"manifestFile"`
	ManifestHash            pulumi.StringInput `pulumi:"manifestHash"`
}

func (ExportDiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDiskDetailsResponse)(nil)).Elem()
}

func (i ExportDiskDetailsResponseArgs) ToExportDiskDetailsResponseOutput() ExportDiskDetailsResponseOutput {
	return i.ToExportDiskDetailsResponseOutputWithContext(context.Background())
}

func (i ExportDiskDetailsResponseArgs) ToExportDiskDetailsResponseOutputWithContext(ctx context.Context) ExportDiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDiskDetailsResponseOutput)
}





type ExportDiskDetailsResponseMapInput interface {
	pulumi.Input

	ToExportDiskDetailsResponseMapOutput() ExportDiskDetailsResponseMapOutput
	ToExportDiskDetailsResponseMapOutputWithContext(context.Context) ExportDiskDetailsResponseMapOutput
}

type ExportDiskDetailsResponseMap map[string]ExportDiskDetailsResponseInput

func (ExportDiskDetailsResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ExportDiskDetailsResponse)(nil)).Elem()
}

func (i ExportDiskDetailsResponseMap) ToExportDiskDetailsResponseMapOutput() ExportDiskDetailsResponseMapOutput {
	return i.ToExportDiskDetailsResponseMapOutputWithContext(context.Background())
}

func (i ExportDiskDetailsResponseMap) ToExportDiskDetailsResponseMapOutputWithContext(ctx context.Context) ExportDiskDetailsResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDiskDetailsResponseMapOutput)
}

type ExportDiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (ExportDiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDiskDetailsResponse)(nil)).Elem()
}

func (o ExportDiskDetailsResponseOutput) ToExportDiskDetailsResponseOutput() ExportDiskDetailsResponseOutput {
	return o
}

func (o ExportDiskDetailsResponseOutput) ToExportDiskDetailsResponseOutputWithContext(ctx context.Context) ExportDiskDetailsResponseOutput {
	return o
}

func (o ExportDiskDetailsResponseOutput) BackupManifestCloudPath() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDiskDetailsResponse) string { return v.BackupManifestCloudPath }).(pulumi.StringOutput)
}

func (o ExportDiskDetailsResponseOutput) ManifestFile() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDiskDetailsResponse) string { return v.ManifestFile }).(pulumi.StringOutput)
}

func (o ExportDiskDetailsResponseOutput) ManifestHash() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDiskDetailsResponse) string { return v.ManifestHash }).(pulumi.StringOutput)
}

type ExportDiskDetailsResponseMapOutput struct{ *pulumi.OutputState }

func (ExportDiskDetailsResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ExportDiskDetailsResponse)(nil)).Elem()
}

func (o ExportDiskDetailsResponseMapOutput) ToExportDiskDetailsResponseMapOutput() ExportDiskDetailsResponseMapOutput {
	return o
}

func (o ExportDiskDetailsResponseMapOutput) ToExportDiskDetailsResponseMapOutputWithContext(ctx context.Context) ExportDiskDetailsResponseMapOutput {
	return o
}

func (o ExportDiskDetailsResponseMapOutput) MapIndex(k pulumi.StringInput) ExportDiskDetailsResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ExportDiskDetailsResponse {
		return vs[0].(map[string]ExportDiskDetailsResponse)[vs[1].(string)]
	}).(ExportDiskDetailsResponseOutput)
}

type FilterFileDetails struct {
	FilterFilePath string `pulumi:"filterFilePath"`
	FilterFileType string `pulumi:"filterFileType"`
}





type FilterFileDetailsInput interface {
	pulumi.Input

	ToFilterFileDetailsOutput() FilterFileDetailsOutput
	ToFilterFileDetailsOutputWithContext(context.Context) FilterFileDetailsOutput
}

type FilterFileDetailsArgs struct {
	FilterFilePath pulumi.StringInput `pulumi:"filterFilePath"`
	FilterFileType pulumi.StringInput `pulumi:"filterFileType"`
}

func (FilterFileDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFileDetails)(nil)).Elem()
}

func (i FilterFileDetailsArgs) ToFilterFileDetailsOutput() FilterFileDetailsOutput {
	return i.ToFilterFileDetailsOutputWithContext(context.Background())
}

func (i FilterFileDetailsArgs) ToFilterFileDetailsOutputWithContext(ctx context.Context) FilterFileDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFileDetailsOutput)
}





type FilterFileDetailsArrayInput interface {
	pulumi.Input

	ToFilterFileDetailsArrayOutput() FilterFileDetailsArrayOutput
	ToFilterFileDetailsArrayOutputWithContext(context.Context) FilterFileDetailsArrayOutput
}

type FilterFileDetailsArray []FilterFileDetailsInput

func (FilterFileDetailsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterFileDetails)(nil)).Elem()
}

func (i FilterFileDetailsArray) ToFilterFileDetailsArrayOutput() FilterFileDetailsArrayOutput {
	return i.ToFilterFileDetailsArrayOutputWithContext(context.Background())
}

func (i FilterFileDetailsArray) ToFilterFileDetailsArrayOutputWithContext(ctx context.Context) FilterFileDetailsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFileDetailsArrayOutput)
}

type FilterFileDetailsOutput struct{ *pulumi.OutputState }

func (FilterFileDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFileDetails)(nil)).Elem()
}

func (o FilterFileDetailsOutput) ToFilterFileDetailsOutput() FilterFileDetailsOutput {
	return o
}

func (o FilterFileDetailsOutput) ToFilterFileDetailsOutputWithContext(ctx context.Context) FilterFileDetailsOutput {
	return o
}

func (o FilterFileDetailsOutput) FilterFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v FilterFileDetails) string { return v.FilterFilePath }).(pulumi.StringOutput)
}

func (o FilterFileDetailsOutput) FilterFileType() pulumi.StringOutput {
	return o.ApplyT(func(v FilterFileDetails) string { return v.FilterFileType }).(pulumi.StringOutput)
}

type FilterFileDetailsArrayOutput struct{ *pulumi.OutputState }

func (FilterFileDetailsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterFileDetails)(nil)).Elem()
}

func (o FilterFileDetailsArrayOutput) ToFilterFileDetailsArrayOutput() FilterFileDetailsArrayOutput {
	return o
}

func (o FilterFileDetailsArrayOutput) ToFilterFileDetailsArrayOutputWithContext(ctx context.Context) FilterFileDetailsArrayOutput {
	return o
}

func (o FilterFileDetailsArrayOutput) Index(i pulumi.IntInput) FilterFileDetailsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilterFileDetails {
		return vs[0].([]FilterFileDetails)[vs[1].(int)]
	}).(FilterFileDetailsOutput)
}

type FilterFileDetailsResponse struct {
	FilterFilePath string `pulumi:"filterFilePath"`
	FilterFileType string `pulumi:"filterFileType"`
}





type FilterFileDetailsResponseInput interface {
	pulumi.Input

	ToFilterFileDetailsResponseOutput() FilterFileDetailsResponseOutput
	ToFilterFileDetailsResponseOutputWithContext(context.Context) FilterFileDetailsResponseOutput
}

type FilterFileDetailsResponseArgs struct {
	FilterFilePath pulumi.StringInput `pulumi:"filterFilePath"`
	FilterFileType pulumi.StringInput `pulumi:"filterFileType"`
}

func (FilterFileDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFileDetailsResponse)(nil)).Elem()
}

func (i FilterFileDetailsResponseArgs) ToFilterFileDetailsResponseOutput() FilterFileDetailsResponseOutput {
	return i.ToFilterFileDetailsResponseOutputWithContext(context.Background())
}

func (i FilterFileDetailsResponseArgs) ToFilterFileDetailsResponseOutputWithContext(ctx context.Context) FilterFileDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFileDetailsResponseOutput)
}





type FilterFileDetailsResponseArrayInput interface {
	pulumi.Input

	ToFilterFileDetailsResponseArrayOutput() FilterFileDetailsResponseArrayOutput
	ToFilterFileDetailsResponseArrayOutputWithContext(context.Context) FilterFileDetailsResponseArrayOutput
}

type FilterFileDetailsResponseArray []FilterFileDetailsResponseInput

func (FilterFileDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterFileDetailsResponse)(nil)).Elem()
}

func (i FilterFileDetailsResponseArray) ToFilterFileDetailsResponseArrayOutput() FilterFileDetailsResponseArrayOutput {
	return i.ToFilterFileDetailsResponseArrayOutputWithContext(context.Background())
}

func (i FilterFileDetailsResponseArray) ToFilterFileDetailsResponseArrayOutputWithContext(ctx context.Context) FilterFileDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFileDetailsResponseArrayOutput)
}

type FilterFileDetailsResponseOutput struct{ *pulumi.OutputState }

func (FilterFileDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFileDetailsResponse)(nil)).Elem()
}

func (o FilterFileDetailsResponseOutput) ToFilterFileDetailsResponseOutput() FilterFileDetailsResponseOutput {
	return o
}

func (o FilterFileDetailsResponseOutput) ToFilterFileDetailsResponseOutputWithContext(ctx context.Context) FilterFileDetailsResponseOutput {
	return o
}

func (o FilterFileDetailsResponseOutput) FilterFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v FilterFileDetailsResponse) string { return v.FilterFilePath }).(pulumi.StringOutput)
}

func (o FilterFileDetailsResponseOutput) FilterFileType() pulumi.StringOutput {
	return o.ApplyT(func(v FilterFileDetailsResponse) string { return v.FilterFileType }).(pulumi.StringOutput)
}

type FilterFileDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (FilterFileDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterFileDetailsResponse)(nil)).Elem()
}

func (o FilterFileDetailsResponseArrayOutput) ToFilterFileDetailsResponseArrayOutput() FilterFileDetailsResponseArrayOutput {
	return o
}

func (o FilterFileDetailsResponseArrayOutput) ToFilterFileDetailsResponseArrayOutputWithContext(ctx context.Context) FilterFileDetailsResponseArrayOutput {
	return o
}

func (o FilterFileDetailsResponseArrayOutput) Index(i pulumi.IntInput) FilterFileDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilterFileDetailsResponse {
		return vs[0].([]FilterFileDetailsResponse)[vs[1].(int)]
	}).(FilterFileDetailsResponseOutput)
}

type IdentityProperties struct {
	Type         *string                 `pulumi:"type"`
	UserAssigned *UserAssignedProperties `pulumi:"userAssigned"`
}





type IdentityPropertiesInput interface {
	pulumi.Input

	ToIdentityPropertiesOutput() IdentityPropertiesOutput
	ToIdentityPropertiesOutputWithContext(context.Context) IdentityPropertiesOutput
}

type IdentityPropertiesArgs struct {
	Type         pulumi.StringPtrInput          `pulumi:"type"`
	UserAssigned UserAssignedPropertiesPtrInput `pulumi:"userAssigned"`
}

func (IdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return i.ToIdentityPropertiesOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput)
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput).ToIdentityPropertiesPtrOutputWithContext(ctx)
}









type IdentityPropertiesPtrInput interface {
	pulumi.Input

	ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput
	ToIdentityPropertiesPtrOutputWithContext(context.Context) IdentityPropertiesPtrOutput
}

type identityPropertiesPtrType IdentityPropertiesArgs

func IdentityPropertiesPtr(v *IdentityPropertiesArgs) IdentityPropertiesPtrInput {
	return (*identityPropertiesPtrType)(v)
}

func (*identityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesPtrOutput)
}

type IdentityPropertiesOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProperties) *IdentityProperties {
		return &v
	}).(IdentityPropertiesPtrOutput)
}

func (o IdentityPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesOutput) UserAssigned() UserAssignedPropertiesPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *UserAssignedProperties { return v.UserAssigned }).(UserAssignedPropertiesPtrOutput)
}

type IdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) Elem() IdentityPropertiesOutput {
	return o.ApplyT(func(v *IdentityProperties) IdentityProperties {
		if v != nil {
			return *v
		}
		var ret IdentityProperties
		return ret
	}).(IdentityPropertiesOutput)
}

func (o IdentityPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesPtrOutput) UserAssigned() UserAssignedPropertiesPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *UserAssignedProperties {
		if v == nil {
			return nil
		}
		return v.UserAssigned
	}).(UserAssignedPropertiesPtrOutput)
}

type IdentityPropertiesResponse struct {
	Type         *string                         `pulumi:"type"`
	UserAssigned *UserAssignedPropertiesResponse `pulumi:"userAssigned"`
}





type IdentityPropertiesResponseInput interface {
	pulumi.Input

	ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput
	ToIdentityPropertiesResponseOutputWithContext(context.Context) IdentityPropertiesResponseOutput
}

type IdentityPropertiesResponseArgs struct {
	Type         pulumi.StringPtrInput                  `pulumi:"type"`
	UserAssigned UserAssignedPropertiesResponsePtrInput `pulumi:"userAssigned"`
}

func (IdentityPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return i.ToIdentityPropertiesResponseOutputWithContext(context.Background())
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponseOutput)
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return i.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponseOutput).ToIdentityPropertiesResponsePtrOutputWithContext(ctx)
}









type IdentityPropertiesResponsePtrInput interface {
	pulumi.Input

	ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput
	ToIdentityPropertiesResponsePtrOutputWithContext(context.Context) IdentityPropertiesResponsePtrOutput
}

type identityPropertiesResponsePtrType IdentityPropertiesResponseArgs

func IdentityPropertiesResponsePtr(v *IdentityPropertiesResponseArgs) IdentityPropertiesResponsePtrInput {
	return (*identityPropertiesResponsePtrType)(v)
}

func (*identityPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (i *identityPropertiesResponsePtrType) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return i.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *identityPropertiesResponsePtrType) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponsePtrOutput)
}

type IdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityPropertiesResponse) *IdentityPropertiesResponse {
		return &v
	}).(IdentityPropertiesResponsePtrOutput)
}

func (o IdentityPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponseOutput) UserAssigned() UserAssignedPropertiesResponsePtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *UserAssignedPropertiesResponse { return v.UserAssigned }).(UserAssignedPropertiesResponsePtrOutput)
}

type IdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) Elem() IdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) IdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret IdentityPropertiesResponse
		return ret
	}).(IdentityPropertiesResponseOutput)
}

func (o IdentityPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) UserAssigned() UserAssignedPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *UserAssignedPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.UserAssigned
	}).(UserAssignedPropertiesResponsePtrOutput)
}

type ImportDiskDetails struct {
	BitLockerKey string `pulumi:"bitLockerKey"`
	ManifestFile string `pulumi:"manifestFile"`
	ManifestHash string `pulumi:"manifestHash"`
}





type ImportDiskDetailsInput interface {
	pulumi.Input

	ToImportDiskDetailsOutput() ImportDiskDetailsOutput
	ToImportDiskDetailsOutputWithContext(context.Context) ImportDiskDetailsOutput
}

type ImportDiskDetailsArgs struct {
	BitLockerKey pulumi.StringInput `pulumi:"bitLockerKey"`
	ManifestFile pulumi.StringInput `pulumi:"manifestFile"`
	ManifestHash pulumi.StringInput `pulumi:"manifestHash"`
}

func (ImportDiskDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportDiskDetails)(nil)).Elem()
}

func (i ImportDiskDetailsArgs) ToImportDiskDetailsOutput() ImportDiskDetailsOutput {
	return i.ToImportDiskDetailsOutputWithContext(context.Background())
}

func (i ImportDiskDetailsArgs) ToImportDiskDetailsOutputWithContext(ctx context.Context) ImportDiskDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportDiskDetailsOutput)
}





type ImportDiskDetailsMapInput interface {
	pulumi.Input

	ToImportDiskDetailsMapOutput() ImportDiskDetailsMapOutput
	ToImportDiskDetailsMapOutputWithContext(context.Context) ImportDiskDetailsMapOutput
}

type ImportDiskDetailsMap map[string]ImportDiskDetailsInput

func (ImportDiskDetailsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ImportDiskDetails)(nil)).Elem()
}

func (i ImportDiskDetailsMap) ToImportDiskDetailsMapOutput() ImportDiskDetailsMapOutput {
	return i.ToImportDiskDetailsMapOutputWithContext(context.Background())
}

func (i ImportDiskDetailsMap) ToImportDiskDetailsMapOutputWithContext(ctx context.Context) ImportDiskDetailsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportDiskDetailsMapOutput)
}

type ImportDiskDetailsOutput struct{ *pulumi.OutputState }

func (ImportDiskDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportDiskDetails)(nil)).Elem()
}

func (o ImportDiskDetailsOutput) ToImportDiskDetailsOutput() ImportDiskDetailsOutput {
	return o
}

func (o ImportDiskDetailsOutput) ToImportDiskDetailsOutputWithContext(ctx context.Context) ImportDiskDetailsOutput {
	return o
}

func (o ImportDiskDetailsOutput) BitLockerKey() pulumi.StringOutput {
	return o.ApplyT(func(v ImportDiskDetails) string { return v.BitLockerKey }).(pulumi.StringOutput)
}

func (o ImportDiskDetailsOutput) ManifestFile() pulumi.StringOutput {
	return o.ApplyT(func(v ImportDiskDetails) string { return v.ManifestFile }).(pulumi.StringOutput)
}

func (o ImportDiskDetailsOutput) ManifestHash() pulumi.StringOutput {
	return o.ApplyT(func(v ImportDiskDetails) string { return v.ManifestHash }).(pulumi.StringOutput)
}

type ImportDiskDetailsMapOutput struct{ *pulumi.OutputState }

func (ImportDiskDetailsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ImportDiskDetails)(nil)).Elem()
}

func (o ImportDiskDetailsMapOutput) ToImportDiskDetailsMapOutput() ImportDiskDetailsMapOutput {
	return o
}

func (o ImportDiskDetailsMapOutput) ToImportDiskDetailsMapOutputWithContext(ctx context.Context) ImportDiskDetailsMapOutput {
	return o
}

func (o ImportDiskDetailsMapOutput) MapIndex(k pulumi.StringInput) ImportDiskDetailsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ImportDiskDetails {
		return vs[0].(map[string]ImportDiskDetails)[vs[1].(string)]
	}).(ImportDiskDetailsOutput)
}

type ImportDiskDetailsResponse struct {
	BackupManifestCloudPath string `pulumi:"backupManifestCloudPath"`
	BitLockerKey            string `pulumi:"bitLockerKey"`
	ManifestFile            string `pulumi:"manifestFile"`
	ManifestHash            string `pulumi:"manifestHash"`
}





type ImportDiskDetailsResponseInput interface {
	pulumi.Input

	ToImportDiskDetailsResponseOutput() ImportDiskDetailsResponseOutput
	ToImportDiskDetailsResponseOutputWithContext(context.Context) ImportDiskDetailsResponseOutput
}

type ImportDiskDetailsResponseArgs struct {
	BackupManifestCloudPath pulumi.StringInput `pulumi:"backupManifestCloudPath"`
	BitLockerKey            pulumi.StringInput `pulumi:"bitLockerKey"`
	ManifestFile            pulumi.StringInput `pulumi:"manifestFile"`
	ManifestHash            pulumi.StringInput `pulumi:"manifestHash"`
}

func (ImportDiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportDiskDetailsResponse)(nil)).Elem()
}

func (i ImportDiskDetailsResponseArgs) ToImportDiskDetailsResponseOutput() ImportDiskDetailsResponseOutput {
	return i.ToImportDiskDetailsResponseOutputWithContext(context.Background())
}

func (i ImportDiskDetailsResponseArgs) ToImportDiskDetailsResponseOutputWithContext(ctx context.Context) ImportDiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportDiskDetailsResponseOutput)
}





type ImportDiskDetailsResponseMapInput interface {
	pulumi.Input

	ToImportDiskDetailsResponseMapOutput() ImportDiskDetailsResponseMapOutput
	ToImportDiskDetailsResponseMapOutputWithContext(context.Context) ImportDiskDetailsResponseMapOutput
}

type ImportDiskDetailsResponseMap map[string]ImportDiskDetailsResponseInput

func (ImportDiskDetailsResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ImportDiskDetailsResponse)(nil)).Elem()
}

func (i ImportDiskDetailsResponseMap) ToImportDiskDetailsResponseMapOutput() ImportDiskDetailsResponseMapOutput {
	return i.ToImportDiskDetailsResponseMapOutputWithContext(context.Background())
}

func (i ImportDiskDetailsResponseMap) ToImportDiskDetailsResponseMapOutputWithContext(ctx context.Context) ImportDiskDetailsResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportDiskDetailsResponseMapOutput)
}

type ImportDiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (ImportDiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportDiskDetailsResponse)(nil)).Elem()
}

func (o ImportDiskDetailsResponseOutput) ToImportDiskDetailsResponseOutput() ImportDiskDetailsResponseOutput {
	return o
}

func (o ImportDiskDetailsResponseOutput) ToImportDiskDetailsResponseOutputWithContext(ctx context.Context) ImportDiskDetailsResponseOutput {
	return o
}

func (o ImportDiskDetailsResponseOutput) BackupManifestCloudPath() pulumi.StringOutput {
	return o.ApplyT(func(v ImportDiskDetailsResponse) string { return v.BackupManifestCloudPath }).(pulumi.StringOutput)
}

func (o ImportDiskDetailsResponseOutput) BitLockerKey() pulumi.StringOutput {
	return o.ApplyT(func(v ImportDiskDetailsResponse) string { return v.BitLockerKey }).(pulumi.StringOutput)
}

func (o ImportDiskDetailsResponseOutput) ManifestFile() pulumi.StringOutput {
	return o.ApplyT(func(v ImportDiskDetailsResponse) string { return v.ManifestFile }).(pulumi.StringOutput)
}

func (o ImportDiskDetailsResponseOutput) ManifestHash() pulumi.StringOutput {
	return o.ApplyT(func(v ImportDiskDetailsResponse) string { return v.ManifestHash }).(pulumi.StringOutput)
}

type ImportDiskDetailsResponseMapOutput struct{ *pulumi.OutputState }

func (ImportDiskDetailsResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ImportDiskDetailsResponse)(nil)).Elem()
}

func (o ImportDiskDetailsResponseMapOutput) ToImportDiskDetailsResponseMapOutput() ImportDiskDetailsResponseMapOutput {
	return o
}

func (o ImportDiskDetailsResponseMapOutput) ToImportDiskDetailsResponseMapOutputWithContext(ctx context.Context) ImportDiskDetailsResponseMapOutput {
	return o
}

func (o ImportDiskDetailsResponseMapOutput) MapIndex(k pulumi.StringInput) ImportDiskDetailsResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ImportDiskDetailsResponse {
		return vs[0].(map[string]ImportDiskDetailsResponse)[vs[1].(string)]
	}).(ImportDiskDetailsResponseOutput)
}

type JobDeliveryInfo struct {
	ScheduledDateTime *string `pulumi:"scheduledDateTime"`
}





type JobDeliveryInfoInput interface {
	pulumi.Input

	ToJobDeliveryInfoOutput() JobDeliveryInfoOutput
	ToJobDeliveryInfoOutputWithContext(context.Context) JobDeliveryInfoOutput
}

type JobDeliveryInfoArgs struct {
	ScheduledDateTime pulumi.StringPtrInput `pulumi:"scheduledDateTime"`
}

func (JobDeliveryInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDeliveryInfo)(nil)).Elem()
}

func (i JobDeliveryInfoArgs) ToJobDeliveryInfoOutput() JobDeliveryInfoOutput {
	return i.ToJobDeliveryInfoOutputWithContext(context.Background())
}

func (i JobDeliveryInfoArgs) ToJobDeliveryInfoOutputWithContext(ctx context.Context) JobDeliveryInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoOutput)
}

func (i JobDeliveryInfoArgs) ToJobDeliveryInfoPtrOutput() JobDeliveryInfoPtrOutput {
	return i.ToJobDeliveryInfoPtrOutputWithContext(context.Background())
}

func (i JobDeliveryInfoArgs) ToJobDeliveryInfoPtrOutputWithContext(ctx context.Context) JobDeliveryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoOutput).ToJobDeliveryInfoPtrOutputWithContext(ctx)
}









type JobDeliveryInfoPtrInput interface {
	pulumi.Input

	ToJobDeliveryInfoPtrOutput() JobDeliveryInfoPtrOutput
	ToJobDeliveryInfoPtrOutputWithContext(context.Context) JobDeliveryInfoPtrOutput
}

type jobDeliveryInfoPtrType JobDeliveryInfoArgs

func JobDeliveryInfoPtr(v *JobDeliveryInfoArgs) JobDeliveryInfoPtrInput {
	return (*jobDeliveryInfoPtrType)(v)
}

func (*jobDeliveryInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDeliveryInfo)(nil)).Elem()
}

func (i *jobDeliveryInfoPtrType) ToJobDeliveryInfoPtrOutput() JobDeliveryInfoPtrOutput {
	return i.ToJobDeliveryInfoPtrOutputWithContext(context.Background())
}

func (i *jobDeliveryInfoPtrType) ToJobDeliveryInfoPtrOutputWithContext(ctx context.Context) JobDeliveryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoPtrOutput)
}

type JobDeliveryInfoOutput struct{ *pulumi.OutputState }

func (JobDeliveryInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDeliveryInfo)(nil)).Elem()
}

func (o JobDeliveryInfoOutput) ToJobDeliveryInfoOutput() JobDeliveryInfoOutput {
	return o
}

func (o JobDeliveryInfoOutput) ToJobDeliveryInfoOutputWithContext(ctx context.Context) JobDeliveryInfoOutput {
	return o
}

func (o JobDeliveryInfoOutput) ToJobDeliveryInfoPtrOutput() JobDeliveryInfoPtrOutput {
	return o.ToJobDeliveryInfoPtrOutputWithContext(context.Background())
}

func (o JobDeliveryInfoOutput) ToJobDeliveryInfoPtrOutputWithContext(ctx context.Context) JobDeliveryInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobDeliveryInfo) *JobDeliveryInfo {
		return &v
	}).(JobDeliveryInfoPtrOutput)
}

func (o JobDeliveryInfoOutput) ScheduledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDeliveryInfo) *string { return v.ScheduledDateTime }).(pulumi.StringPtrOutput)
}

type JobDeliveryInfoPtrOutput struct{ *pulumi.OutputState }

func (JobDeliveryInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDeliveryInfo)(nil)).Elem()
}

func (o JobDeliveryInfoPtrOutput) ToJobDeliveryInfoPtrOutput() JobDeliveryInfoPtrOutput {
	return o
}

func (o JobDeliveryInfoPtrOutput) ToJobDeliveryInfoPtrOutputWithContext(ctx context.Context) JobDeliveryInfoPtrOutput {
	return o
}

func (o JobDeliveryInfoPtrOutput) Elem() JobDeliveryInfoOutput {
	return o.ApplyT(func(v *JobDeliveryInfo) JobDeliveryInfo {
		if v != nil {
			return *v
		}
		var ret JobDeliveryInfo
		return ret
	}).(JobDeliveryInfoOutput)
}

func (o JobDeliveryInfoPtrOutput) ScheduledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDeliveryInfo) *string {
		if v == nil {
			return nil
		}
		return v.ScheduledDateTime
	}).(pulumi.StringPtrOutput)
}

type JobDeliveryInfoResponse struct {
	ScheduledDateTime *string `pulumi:"scheduledDateTime"`
}





type JobDeliveryInfoResponseInput interface {
	pulumi.Input

	ToJobDeliveryInfoResponseOutput() JobDeliveryInfoResponseOutput
	ToJobDeliveryInfoResponseOutputWithContext(context.Context) JobDeliveryInfoResponseOutput
}

type JobDeliveryInfoResponseArgs struct {
	ScheduledDateTime pulumi.StringPtrInput `pulumi:"scheduledDateTime"`
}

func (JobDeliveryInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDeliveryInfoResponse)(nil)).Elem()
}

func (i JobDeliveryInfoResponseArgs) ToJobDeliveryInfoResponseOutput() JobDeliveryInfoResponseOutput {
	return i.ToJobDeliveryInfoResponseOutputWithContext(context.Background())
}

func (i JobDeliveryInfoResponseArgs) ToJobDeliveryInfoResponseOutputWithContext(ctx context.Context) JobDeliveryInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoResponseOutput)
}

func (i JobDeliveryInfoResponseArgs) ToJobDeliveryInfoResponsePtrOutput() JobDeliveryInfoResponsePtrOutput {
	return i.ToJobDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (i JobDeliveryInfoResponseArgs) ToJobDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) JobDeliveryInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoResponseOutput).ToJobDeliveryInfoResponsePtrOutputWithContext(ctx)
}









type JobDeliveryInfoResponsePtrInput interface {
	pulumi.Input

	ToJobDeliveryInfoResponsePtrOutput() JobDeliveryInfoResponsePtrOutput
	ToJobDeliveryInfoResponsePtrOutputWithContext(context.Context) JobDeliveryInfoResponsePtrOutput
}

type jobDeliveryInfoResponsePtrType JobDeliveryInfoResponseArgs

func JobDeliveryInfoResponsePtr(v *JobDeliveryInfoResponseArgs) JobDeliveryInfoResponsePtrInput {
	return (*jobDeliveryInfoResponsePtrType)(v)
}

func (*jobDeliveryInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDeliveryInfoResponse)(nil)).Elem()
}

func (i *jobDeliveryInfoResponsePtrType) ToJobDeliveryInfoResponsePtrOutput() JobDeliveryInfoResponsePtrOutput {
	return i.ToJobDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (i *jobDeliveryInfoResponsePtrType) ToJobDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) JobDeliveryInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDeliveryInfoResponsePtrOutput)
}

type JobDeliveryInfoResponseOutput struct{ *pulumi.OutputState }

func (JobDeliveryInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDeliveryInfoResponse)(nil)).Elem()
}

func (o JobDeliveryInfoResponseOutput) ToJobDeliveryInfoResponseOutput() JobDeliveryInfoResponseOutput {
	return o
}

func (o JobDeliveryInfoResponseOutput) ToJobDeliveryInfoResponseOutputWithContext(ctx context.Context) JobDeliveryInfoResponseOutput {
	return o
}

func (o JobDeliveryInfoResponseOutput) ToJobDeliveryInfoResponsePtrOutput() JobDeliveryInfoResponsePtrOutput {
	return o.ToJobDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (o JobDeliveryInfoResponseOutput) ToJobDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) JobDeliveryInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobDeliveryInfoResponse) *JobDeliveryInfoResponse {
		return &v
	}).(JobDeliveryInfoResponsePtrOutput)
}

func (o JobDeliveryInfoResponseOutput) ScheduledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobDeliveryInfoResponse) *string { return v.ScheduledDateTime }).(pulumi.StringPtrOutput)
}

type JobDeliveryInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (JobDeliveryInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDeliveryInfoResponse)(nil)).Elem()
}

func (o JobDeliveryInfoResponsePtrOutput) ToJobDeliveryInfoResponsePtrOutput() JobDeliveryInfoResponsePtrOutput {
	return o
}

func (o JobDeliveryInfoResponsePtrOutput) ToJobDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) JobDeliveryInfoResponsePtrOutput {
	return o
}

func (o JobDeliveryInfoResponsePtrOutput) Elem() JobDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *JobDeliveryInfoResponse) JobDeliveryInfoResponse {
		if v != nil {
			return *v
		}
		var ret JobDeliveryInfoResponse
		return ret
	}).(JobDeliveryInfoResponseOutput)
}

func (o JobDeliveryInfoResponsePtrOutput) ScheduledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDeliveryInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScheduledDateTime
	}).(pulumi.StringPtrOutput)
}

type JobStagesResponse struct {
	DisplayName     string      `pulumi:"displayName"`
	JobStageDetails interface{} `pulumi:"jobStageDetails"`
	StageName       string      `pulumi:"stageName"`
	StageStatus     string      `pulumi:"stageStatus"`
	StageTime       string      `pulumi:"stageTime"`
}





type JobStagesResponseInput interface {
	pulumi.Input

	ToJobStagesResponseOutput() JobStagesResponseOutput
	ToJobStagesResponseOutputWithContext(context.Context) JobStagesResponseOutput
}

type JobStagesResponseArgs struct {
	DisplayName     pulumi.StringInput `pulumi:"displayName"`
	JobStageDetails pulumi.Input       `pulumi:"jobStageDetails"`
	StageName       pulumi.StringInput `pulumi:"stageName"`
	StageStatus     pulumi.StringInput `pulumi:"stageStatus"`
	StageTime       pulumi.StringInput `pulumi:"stageTime"`
}

func (JobStagesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStagesResponse)(nil)).Elem()
}

func (i JobStagesResponseArgs) ToJobStagesResponseOutput() JobStagesResponseOutput {
	return i.ToJobStagesResponseOutputWithContext(context.Background())
}

func (i JobStagesResponseArgs) ToJobStagesResponseOutputWithContext(ctx context.Context) JobStagesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStagesResponseOutput)
}





type JobStagesResponseArrayInput interface {
	pulumi.Input

	ToJobStagesResponseArrayOutput() JobStagesResponseArrayOutput
	ToJobStagesResponseArrayOutputWithContext(context.Context) JobStagesResponseArrayOutput
}

type JobStagesResponseArray []JobStagesResponseInput

func (JobStagesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobStagesResponse)(nil)).Elem()
}

func (i JobStagesResponseArray) ToJobStagesResponseArrayOutput() JobStagesResponseArrayOutput {
	return i.ToJobStagesResponseArrayOutputWithContext(context.Background())
}

func (i JobStagesResponseArray) ToJobStagesResponseArrayOutputWithContext(ctx context.Context) JobStagesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStagesResponseArrayOutput)
}

type JobStagesResponseOutput struct{ *pulumi.OutputState }

func (JobStagesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStagesResponse)(nil)).Elem()
}

func (o JobStagesResponseOutput) ToJobStagesResponseOutput() JobStagesResponseOutput {
	return o
}

func (o JobStagesResponseOutput) ToJobStagesResponseOutputWithContext(ctx context.Context) JobStagesResponseOutput {
	return o
}

func (o JobStagesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStagesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o JobStagesResponseOutput) JobStageDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v JobStagesResponse) interface{} { return v.JobStageDetails }).(pulumi.AnyOutput)
}

func (o JobStagesResponseOutput) StageName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStagesResponse) string { return v.StageName }).(pulumi.StringOutput)
}

func (o JobStagesResponseOutput) StageStatus() pulumi.StringOutput {
	return o.ApplyT(func(v JobStagesResponse) string { return v.StageStatus }).(pulumi.StringOutput)
}

func (o JobStagesResponseOutput) StageTime() pulumi.StringOutput {
	return o.ApplyT(func(v JobStagesResponse) string { return v.StageTime }).(pulumi.StringOutput)
}

type JobStagesResponseArrayOutput struct{ *pulumi.OutputState }

func (JobStagesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobStagesResponse)(nil)).Elem()
}

func (o JobStagesResponseArrayOutput) ToJobStagesResponseArrayOutput() JobStagesResponseArrayOutput {
	return o
}

func (o JobStagesResponseArrayOutput) ToJobStagesResponseArrayOutputWithContext(ctx context.Context) JobStagesResponseArrayOutput {
	return o
}

func (o JobStagesResponseArrayOutput) Index(i pulumi.IntInput) JobStagesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobStagesResponse {
		return vs[0].([]JobStagesResponse)[vs[1].(int)]
	}).(JobStagesResponseOutput)
}

type KeyEncryptionKey struct {
	IdentityProperties *IdentityProperties `pulumi:"identityProperties"`
	KekType            string              `pulumi:"kekType"`
	KekUrl             *string             `pulumi:"kekUrl"`
	KekVaultResourceID *string             `pulumi:"kekVaultResourceID"`
}





type KeyEncryptionKeyInput interface {
	pulumi.Input

	ToKeyEncryptionKeyOutput() KeyEncryptionKeyOutput
	ToKeyEncryptionKeyOutputWithContext(context.Context) KeyEncryptionKeyOutput
}

type KeyEncryptionKeyArgs struct {
	IdentityProperties IdentityPropertiesPtrInput `pulumi:"identityProperties"`
	KekType            pulumi.StringInput         `pulumi:"kekType"`
	KekUrl             pulumi.StringPtrInput      `pulumi:"kekUrl"`
	KekVaultResourceID pulumi.StringPtrInput      `pulumi:"kekVaultResourceID"`
}

func (KeyEncryptionKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyEncryptionKey)(nil)).Elem()
}

func (i KeyEncryptionKeyArgs) ToKeyEncryptionKeyOutput() KeyEncryptionKeyOutput {
	return i.ToKeyEncryptionKeyOutputWithContext(context.Background())
}

func (i KeyEncryptionKeyArgs) ToKeyEncryptionKeyOutputWithContext(ctx context.Context) KeyEncryptionKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyEncryptionKeyOutput)
}

func (i KeyEncryptionKeyArgs) ToKeyEncryptionKeyPtrOutput() KeyEncryptionKeyPtrOutput {
	return i.ToKeyEncryptionKeyPtrOutputWithContext(context.Background())
}

func (i KeyEncryptionKeyArgs) ToKeyEncryptionKeyPtrOutputWithContext(ctx context.Context) KeyEncryptionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyEncryptionKeyOutput).ToKeyEncryptionKeyPtrOutputWithContext(ctx)
}









type KeyEncryptionKeyPtrInput interface {
	pulumi.Input

	ToKeyEncryptionKeyPtrOutput() KeyEncryptionKeyPtrOutput
	ToKeyEncryptionKeyPtrOutputWithContext(context.Context) KeyEncryptionKeyPtrOutput
}

type keyEncryptionKeyPtrType KeyEncryptionKeyArgs

func KeyEncryptionKeyPtr(v *KeyEncryptionKeyArgs) KeyEncryptionKeyPtrInput {
	return (*keyEncryptionKeyPtrType)(v)
}

func (*keyEncryptionKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyEncryptionKey)(nil)).Elem()
}

func (i *keyEncryptionKeyPtrType) ToKeyEncryptionKeyPtrOutput() KeyEncryptionKeyPtrOutput {
	return i.ToKeyEncryptionKeyPtrOutputWithContext(context.Background())
}

func (i *keyEncryptionKeyPtrType) ToKeyEncryptionKeyPtrOutputWithContext(ctx context.Context) KeyEncryptionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyEncryptionKeyPtrOutput)
}

type KeyEncryptionKeyOutput struct{ *pulumi.OutputState }

func (KeyEncryptionKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyEncryptionKey)(nil)).Elem()
}

func (o KeyEncryptionKeyOutput) ToKeyEncryptionKeyOutput() KeyEncryptionKeyOutput {
	return o
}

func (o KeyEncryptionKeyOutput) ToKeyEncryptionKeyOutputWithContext(ctx context.Context) KeyEncryptionKeyOutput {
	return o
}

func (o KeyEncryptionKeyOutput) ToKeyEncryptionKeyPtrOutput() KeyEncryptionKeyPtrOutput {
	return o.ToKeyEncryptionKeyPtrOutputWithContext(context.Background())
}

func (o KeyEncryptionKeyOutput) ToKeyEncryptionKeyPtrOutputWithContext(ctx context.Context) KeyEncryptionKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyEncryptionKey) *KeyEncryptionKey {
		return &v
	}).(KeyEncryptionKeyPtrOutput)
}

func (o KeyEncryptionKeyOutput) IdentityProperties() IdentityPropertiesPtrOutput {
	return o.ApplyT(func(v KeyEncryptionKey) *IdentityProperties { return v.IdentityProperties }).(IdentityPropertiesPtrOutput)
}

func (o KeyEncryptionKeyOutput) KekType() pulumi.StringOutput {
	return o.ApplyT(func(v KeyEncryptionKey) string { return v.KekType }).(pulumi.StringOutput)
}

func (o KeyEncryptionKeyOutput) KekUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyEncryptionKey) *string { return v.KekUrl }).(pulumi.StringPtrOutput)
}

func (o KeyEncryptionKeyOutput) KekVaultResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyEncryptionKey) *string { return v.KekVaultResourceID }).(pulumi.StringPtrOutput)
}

type KeyEncryptionKeyPtrOutput struct{ *pulumi.OutputState }

func (KeyEncryptionKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyEncryptionKey)(nil)).Elem()
}

func (o KeyEncryptionKeyPtrOutput) ToKeyEncryptionKeyPtrOutput() KeyEncryptionKeyPtrOutput {
	return o
}

func (o KeyEncryptionKeyPtrOutput) ToKeyEncryptionKeyPtrOutputWithContext(ctx context.Context) KeyEncryptionKeyPtrOutput {
	return o
}

func (o KeyEncryptionKeyPtrOutput) Elem() KeyEncryptionKeyOutput {
	return o.ApplyT(func(v *KeyEncryptionKey) KeyEncryptionKey {
		if v != nil {
			return *v
		}
		var ret KeyEncryptionKey
		return ret
	}).(KeyEncryptionKeyOutput)
}

func (o KeyEncryptionKeyPtrOutput) IdentityProperties() IdentityPropertiesPtrOutput {
	return o.ApplyT(func(v *KeyEncryptionKey) *IdentityProperties {
		if v == nil {
			return nil
		}
		return v.IdentityProperties
	}).(IdentityPropertiesPtrOutput)
}

func (o KeyEncryptionKeyPtrOutput) KekType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyEncryptionKey) *string {
		if v == nil {
			return nil
		}
		return &v.KekType
	}).(pulumi.StringPtrOutput)
}

func (o KeyEncryptionKeyPtrOutput) KekUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyEncryptionKey) *string {
		if v == nil {
			return nil
		}
		return v.KekUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyEncryptionKeyPtrOutput) KekVaultResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyEncryptionKey) *string {
		if v == nil {
			return nil
		}
		return v.KekVaultResourceID
	}).(pulumi.StringPtrOutput)
}

type KeyEncryptionKeyResponse struct {
	IdentityProperties *IdentityPropertiesResponse `pulumi:"identityProperties"`
	KekType            string                      `pulumi:"kekType"`
	KekUrl             *string                     `pulumi:"kekUrl"`
	KekVaultResourceID *string                     `pulumi:"kekVaultResourceID"`
}





type KeyEncryptionKeyResponseInput interface {
	pulumi.Input

	ToKeyEncryptionKeyResponseOutput() KeyEncryptionKeyResponseOutput
	ToKeyEncryptionKeyResponseOutputWithContext(context.Context) KeyEncryptionKeyResponseOutput
}

type KeyEncryptionKeyResponseArgs struct {
	IdentityProperties IdentityPropertiesResponsePtrInput `pulumi:"identityProperties"`
	KekType            pulumi.StringInput                 `pulumi:"kekType"`
	KekUrl             pulumi.StringPtrInput              `pulumi:"kekUrl"`
	KekVaultResourceID pulumi.StringPtrInput              `pulumi:"kekVaultResourceID"`
}

func (KeyEncryptionKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyEncryptionKeyResponse)(nil)).Elem()
}

func (i KeyEncryptionKeyResponseArgs) ToKeyEncryptionKeyResponseOutput() KeyEncryptionKeyResponseOutput {
	return i.ToKeyEncryptionKeyResponseOutputWithContext(context.Background())
}

func (i KeyEncryptionKeyResponseArgs) ToKeyEncryptionKeyResponseOutputWithContext(ctx context.Context) KeyEncryptionKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyEncryptionKeyResponseOutput)
}

func (i KeyEncryptionKeyResponseArgs) ToKeyEncryptionKeyResponsePtrOutput() KeyEncryptionKeyResponsePtrOutput {
	return i.ToKeyEncryptionKeyResponsePtrOutputWithContext(context.Background())
}

func (i KeyEncryptionKeyResponseArgs) ToKeyEncryptionKeyResponsePtrOutputWithContext(ctx context.Context) KeyEncryptionKeyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyEncryptionKeyResponseOutput).ToKeyEncryptionKeyResponsePtrOutputWithContext(ctx)
}









type KeyEncryptionKeyResponsePtrInput interface {
	pulumi.Input

	ToKeyEncryptionKeyResponsePtrOutput() KeyEncryptionKeyResponsePtrOutput
	ToKeyEncryptionKeyResponsePtrOutputWithContext(context.Context) KeyEncryptionKeyResponsePtrOutput
}

type keyEncryptionKeyResponsePtrType KeyEncryptionKeyResponseArgs

func KeyEncryptionKeyResponsePtr(v *KeyEncryptionKeyResponseArgs) KeyEncryptionKeyResponsePtrInput {
	return (*keyEncryptionKeyResponsePtrType)(v)
}

func (*keyEncryptionKeyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyEncryptionKeyResponse)(nil)).Elem()
}

func (i *keyEncryptionKeyResponsePtrType) ToKeyEncryptionKeyResponsePtrOutput() KeyEncryptionKeyResponsePtrOutput {
	return i.ToKeyEncryptionKeyResponsePtrOutputWithContext(context.Background())
}

func (i *keyEncryptionKeyResponsePtrType) ToKeyEncryptionKeyResponsePtrOutputWithContext(ctx context.Context) KeyEncryptionKeyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyEncryptionKeyResponsePtrOutput)
}

type KeyEncryptionKeyResponseOutput struct{ *pulumi.OutputState }

func (KeyEncryptionKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyEncryptionKeyResponse)(nil)).Elem()
}

func (o KeyEncryptionKeyResponseOutput) ToKeyEncryptionKeyResponseOutput() KeyEncryptionKeyResponseOutput {
	return o
}

func (o KeyEncryptionKeyResponseOutput) ToKeyEncryptionKeyResponseOutputWithContext(ctx context.Context) KeyEncryptionKeyResponseOutput {
	return o
}

func (o KeyEncryptionKeyResponseOutput) ToKeyEncryptionKeyResponsePtrOutput() KeyEncryptionKeyResponsePtrOutput {
	return o.ToKeyEncryptionKeyResponsePtrOutputWithContext(context.Background())
}

func (o KeyEncryptionKeyResponseOutput) ToKeyEncryptionKeyResponsePtrOutputWithContext(ctx context.Context) KeyEncryptionKeyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyEncryptionKeyResponse) *KeyEncryptionKeyResponse {
		return &v
	}).(KeyEncryptionKeyResponsePtrOutput)
}

func (o KeyEncryptionKeyResponseOutput) IdentityProperties() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v KeyEncryptionKeyResponse) *IdentityPropertiesResponse { return v.IdentityProperties }).(IdentityPropertiesResponsePtrOutput)
}

func (o KeyEncryptionKeyResponseOutput) KekType() pulumi.StringOutput {
	return o.ApplyT(func(v KeyEncryptionKeyResponse) string { return v.KekType }).(pulumi.StringOutput)
}

func (o KeyEncryptionKeyResponseOutput) KekUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyEncryptionKeyResponse) *string { return v.KekUrl }).(pulumi.StringPtrOutput)
}

func (o KeyEncryptionKeyResponseOutput) KekVaultResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyEncryptionKeyResponse) *string { return v.KekVaultResourceID }).(pulumi.StringPtrOutput)
}

type KeyEncryptionKeyResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyEncryptionKeyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyEncryptionKeyResponse)(nil)).Elem()
}

func (o KeyEncryptionKeyResponsePtrOutput) ToKeyEncryptionKeyResponsePtrOutput() KeyEncryptionKeyResponsePtrOutput {
	return o
}

func (o KeyEncryptionKeyResponsePtrOutput) ToKeyEncryptionKeyResponsePtrOutputWithContext(ctx context.Context) KeyEncryptionKeyResponsePtrOutput {
	return o
}

func (o KeyEncryptionKeyResponsePtrOutput) Elem() KeyEncryptionKeyResponseOutput {
	return o.ApplyT(func(v *KeyEncryptionKeyResponse) KeyEncryptionKeyResponse {
		if v != nil {
			return *v
		}
		var ret KeyEncryptionKeyResponse
		return ret
	}).(KeyEncryptionKeyResponseOutput)
}

func (o KeyEncryptionKeyResponsePtrOutput) IdentityProperties() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *KeyEncryptionKeyResponse) *IdentityPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.IdentityProperties
	}).(IdentityPropertiesResponsePtrOutput)
}

func (o KeyEncryptionKeyResponsePtrOutput) KekType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyEncryptionKeyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KekType
	}).(pulumi.StringPtrOutput)
}

func (o KeyEncryptionKeyResponsePtrOutput) KekUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyEncryptionKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.KekUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyEncryptionKeyResponsePtrOutput) KekVaultResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyEncryptionKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.KekVaultResourceID
	}).(pulumi.StringPtrOutput)
}

type LastMitigationActionOnJobResponse struct {
	ActionDateTimeInUtc   *string `pulumi:"actionDateTimeInUtc"`
	CustomerResolution    *string `pulumi:"customerResolution"`
	IsPerformedByCustomer *bool   `pulumi:"isPerformedByCustomer"`
}





type LastMitigationActionOnJobResponseInput interface {
	pulumi.Input

	ToLastMitigationActionOnJobResponseOutput() LastMitigationActionOnJobResponseOutput
	ToLastMitigationActionOnJobResponseOutputWithContext(context.Context) LastMitigationActionOnJobResponseOutput
}

type LastMitigationActionOnJobResponseArgs struct {
	ActionDateTimeInUtc   pulumi.StringPtrInput `pulumi:"actionDateTimeInUtc"`
	CustomerResolution    pulumi.StringPtrInput `pulumi:"customerResolution"`
	IsPerformedByCustomer pulumi.BoolPtrInput   `pulumi:"isPerformedByCustomer"`
}

func (LastMitigationActionOnJobResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LastMitigationActionOnJobResponse)(nil)).Elem()
}

func (i LastMitigationActionOnJobResponseArgs) ToLastMitigationActionOnJobResponseOutput() LastMitigationActionOnJobResponseOutput {
	return i.ToLastMitigationActionOnJobResponseOutputWithContext(context.Background())
}

func (i LastMitigationActionOnJobResponseArgs) ToLastMitigationActionOnJobResponseOutputWithContext(ctx context.Context) LastMitigationActionOnJobResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LastMitigationActionOnJobResponseOutput)
}

type LastMitigationActionOnJobResponseOutput struct{ *pulumi.OutputState }

func (LastMitigationActionOnJobResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LastMitigationActionOnJobResponse)(nil)).Elem()
}

func (o LastMitigationActionOnJobResponseOutput) ToLastMitigationActionOnJobResponseOutput() LastMitigationActionOnJobResponseOutput {
	return o
}

func (o LastMitigationActionOnJobResponseOutput) ToLastMitigationActionOnJobResponseOutputWithContext(ctx context.Context) LastMitigationActionOnJobResponseOutput {
	return o
}

func (o LastMitigationActionOnJobResponseOutput) ActionDateTimeInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LastMitigationActionOnJobResponse) *string { return v.ActionDateTimeInUtc }).(pulumi.StringPtrOutput)
}

func (o LastMitigationActionOnJobResponseOutput) CustomerResolution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LastMitigationActionOnJobResponse) *string { return v.CustomerResolution }).(pulumi.StringPtrOutput)
}

func (o LastMitigationActionOnJobResponseOutput) IsPerformedByCustomer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LastMitigationActionOnJobResponse) *bool { return v.IsPerformedByCustomer }).(pulumi.BoolPtrOutput)
}

type ManagedDiskDetails struct {
	DataAccountType         string  `pulumi:"dataAccountType"`
	ResourceGroupId         string  `pulumi:"resourceGroupId"`
	SharePassword           *string `pulumi:"sharePassword"`
	StagingStorageAccountId string  `pulumi:"stagingStorageAccountId"`
}





type ManagedDiskDetailsInput interface {
	pulumi.Input

	ToManagedDiskDetailsOutput() ManagedDiskDetailsOutput
	ToManagedDiskDetailsOutputWithContext(context.Context) ManagedDiskDetailsOutput
}

type ManagedDiskDetailsArgs struct {
	DataAccountType         pulumi.StringInput    `pulumi:"dataAccountType"`
	ResourceGroupId         pulumi.StringInput    `pulumi:"resourceGroupId"`
	SharePassword           pulumi.StringPtrInput `pulumi:"sharePassword"`
	StagingStorageAccountId pulumi.StringInput    `pulumi:"stagingStorageAccountId"`
}

func (ManagedDiskDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDiskDetails)(nil)).Elem()
}

func (i ManagedDiskDetailsArgs) ToManagedDiskDetailsOutput() ManagedDiskDetailsOutput {
	return i.ToManagedDiskDetailsOutputWithContext(context.Background())
}

func (i ManagedDiskDetailsArgs) ToManagedDiskDetailsOutputWithContext(ctx context.Context) ManagedDiskDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskDetailsOutput)
}

type ManagedDiskDetailsOutput struct{ *pulumi.OutputState }

func (ManagedDiskDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDiskDetails)(nil)).Elem()
}

func (o ManagedDiskDetailsOutput) ToManagedDiskDetailsOutput() ManagedDiskDetailsOutput {
	return o
}

func (o ManagedDiskDetailsOutput) ToManagedDiskDetailsOutputWithContext(ctx context.Context) ManagedDiskDetailsOutput {
	return o
}

func (o ManagedDiskDetailsOutput) DataAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedDiskDetails) string { return v.DataAccountType }).(pulumi.StringOutput)
}

func (o ManagedDiskDetailsOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedDiskDetails) string { return v.ResourceGroupId }).(pulumi.StringOutput)
}

func (o ManagedDiskDetailsOutput) SharePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedDiskDetails) *string { return v.SharePassword }).(pulumi.StringPtrOutput)
}

func (o ManagedDiskDetailsOutput) StagingStorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedDiskDetails) string { return v.StagingStorageAccountId }).(pulumi.StringOutput)
}

type ManagedDiskDetailsResponse struct {
	DataAccountType         string `pulumi:"dataAccountType"`
	ResourceGroupId         string `pulumi:"resourceGroupId"`
	StagingStorageAccountId string `pulumi:"stagingStorageAccountId"`
}





type ManagedDiskDetailsResponseInput interface {
	pulumi.Input

	ToManagedDiskDetailsResponseOutput() ManagedDiskDetailsResponseOutput
	ToManagedDiskDetailsResponseOutputWithContext(context.Context) ManagedDiskDetailsResponseOutput
}

type ManagedDiskDetailsResponseArgs struct {
	DataAccountType         pulumi.StringInput `pulumi:"dataAccountType"`
	ResourceGroupId         pulumi.StringInput `pulumi:"resourceGroupId"`
	StagingStorageAccountId pulumi.StringInput `pulumi:"stagingStorageAccountId"`
}

func (ManagedDiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDiskDetailsResponse)(nil)).Elem()
}

func (i ManagedDiskDetailsResponseArgs) ToManagedDiskDetailsResponseOutput() ManagedDiskDetailsResponseOutput {
	return i.ToManagedDiskDetailsResponseOutputWithContext(context.Background())
}

func (i ManagedDiskDetailsResponseArgs) ToManagedDiskDetailsResponseOutputWithContext(ctx context.Context) ManagedDiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskDetailsResponseOutput)
}

type ManagedDiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (ManagedDiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDiskDetailsResponse)(nil)).Elem()
}

func (o ManagedDiskDetailsResponseOutput) ToManagedDiskDetailsResponseOutput() ManagedDiskDetailsResponseOutput {
	return o
}

func (o ManagedDiskDetailsResponseOutput) ToManagedDiskDetailsResponseOutputWithContext(ctx context.Context) ManagedDiskDetailsResponseOutput {
	return o
}

func (o ManagedDiskDetailsResponseOutput) DataAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedDiskDetailsResponse) string { return v.DataAccountType }).(pulumi.StringOutput)
}

func (o ManagedDiskDetailsResponseOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedDiskDetailsResponse) string { return v.ResourceGroupId }).(pulumi.StringOutput)
}

func (o ManagedDiskDetailsResponseOutput) StagingStorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedDiskDetailsResponse) string { return v.StagingStorageAccountId }).(pulumi.StringOutput)
}

type NotificationPreference struct {
	SendNotification bool   `pulumi:"sendNotification"`
	StageName        string `pulumi:"stageName"`
}





type NotificationPreferenceInput interface {
	pulumi.Input

	ToNotificationPreferenceOutput() NotificationPreferenceOutput
	ToNotificationPreferenceOutputWithContext(context.Context) NotificationPreferenceOutput
}

type NotificationPreferenceArgs struct {
	SendNotification pulumi.BoolInput   `pulumi:"sendNotification"`
	StageName        pulumi.StringInput `pulumi:"stageName"`
}

func (NotificationPreferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationPreference)(nil)).Elem()
}

func (i NotificationPreferenceArgs) ToNotificationPreferenceOutput() NotificationPreferenceOutput {
	return i.ToNotificationPreferenceOutputWithContext(context.Background())
}

func (i NotificationPreferenceArgs) ToNotificationPreferenceOutputWithContext(ctx context.Context) NotificationPreferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPreferenceOutput)
}





type NotificationPreferenceArrayInput interface {
	pulumi.Input

	ToNotificationPreferenceArrayOutput() NotificationPreferenceArrayOutput
	ToNotificationPreferenceArrayOutputWithContext(context.Context) NotificationPreferenceArrayOutput
}

type NotificationPreferenceArray []NotificationPreferenceInput

func (NotificationPreferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationPreference)(nil)).Elem()
}

func (i NotificationPreferenceArray) ToNotificationPreferenceArrayOutput() NotificationPreferenceArrayOutput {
	return i.ToNotificationPreferenceArrayOutputWithContext(context.Background())
}

func (i NotificationPreferenceArray) ToNotificationPreferenceArrayOutputWithContext(ctx context.Context) NotificationPreferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPreferenceArrayOutput)
}

type NotificationPreferenceOutput struct{ *pulumi.OutputState }

func (NotificationPreferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationPreference)(nil)).Elem()
}

func (o NotificationPreferenceOutput) ToNotificationPreferenceOutput() NotificationPreferenceOutput {
	return o
}

func (o NotificationPreferenceOutput) ToNotificationPreferenceOutputWithContext(ctx context.Context) NotificationPreferenceOutput {
	return o
}

func (o NotificationPreferenceOutput) SendNotification() pulumi.BoolOutput {
	return o.ApplyT(func(v NotificationPreference) bool { return v.SendNotification }).(pulumi.BoolOutput)
}

func (o NotificationPreferenceOutput) StageName() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationPreference) string { return v.StageName }).(pulumi.StringOutput)
}

type NotificationPreferenceArrayOutput struct{ *pulumi.OutputState }

func (NotificationPreferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationPreference)(nil)).Elem()
}

func (o NotificationPreferenceArrayOutput) ToNotificationPreferenceArrayOutput() NotificationPreferenceArrayOutput {
	return o
}

func (o NotificationPreferenceArrayOutput) ToNotificationPreferenceArrayOutputWithContext(ctx context.Context) NotificationPreferenceArrayOutput {
	return o
}

func (o NotificationPreferenceArrayOutput) Index(i pulumi.IntInput) NotificationPreferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationPreference {
		return vs[0].([]NotificationPreference)[vs[1].(int)]
	}).(NotificationPreferenceOutput)
}

type NotificationPreferenceResponse struct {
	SendNotification bool   `pulumi:"sendNotification"`
	StageName        string `pulumi:"stageName"`
}





type NotificationPreferenceResponseInput interface {
	pulumi.Input

	ToNotificationPreferenceResponseOutput() NotificationPreferenceResponseOutput
	ToNotificationPreferenceResponseOutputWithContext(context.Context) NotificationPreferenceResponseOutput
}

type NotificationPreferenceResponseArgs struct {
	SendNotification pulumi.BoolInput   `pulumi:"sendNotification"`
	StageName        pulumi.StringInput `pulumi:"stageName"`
}

func (NotificationPreferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationPreferenceResponse)(nil)).Elem()
}

func (i NotificationPreferenceResponseArgs) ToNotificationPreferenceResponseOutput() NotificationPreferenceResponseOutput {
	return i.ToNotificationPreferenceResponseOutputWithContext(context.Background())
}

func (i NotificationPreferenceResponseArgs) ToNotificationPreferenceResponseOutputWithContext(ctx context.Context) NotificationPreferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPreferenceResponseOutput)
}





type NotificationPreferenceResponseArrayInput interface {
	pulumi.Input

	ToNotificationPreferenceResponseArrayOutput() NotificationPreferenceResponseArrayOutput
	ToNotificationPreferenceResponseArrayOutputWithContext(context.Context) NotificationPreferenceResponseArrayOutput
}

type NotificationPreferenceResponseArray []NotificationPreferenceResponseInput

func (NotificationPreferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationPreferenceResponse)(nil)).Elem()
}

func (i NotificationPreferenceResponseArray) ToNotificationPreferenceResponseArrayOutput() NotificationPreferenceResponseArrayOutput {
	return i.ToNotificationPreferenceResponseArrayOutputWithContext(context.Background())
}

func (i NotificationPreferenceResponseArray) ToNotificationPreferenceResponseArrayOutputWithContext(ctx context.Context) NotificationPreferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPreferenceResponseArrayOutput)
}

type NotificationPreferenceResponseOutput struct{ *pulumi.OutputState }

func (NotificationPreferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationPreferenceResponse)(nil)).Elem()
}

func (o NotificationPreferenceResponseOutput) ToNotificationPreferenceResponseOutput() NotificationPreferenceResponseOutput {
	return o
}

func (o NotificationPreferenceResponseOutput) ToNotificationPreferenceResponseOutputWithContext(ctx context.Context) NotificationPreferenceResponseOutput {
	return o
}

func (o NotificationPreferenceResponseOutput) SendNotification() pulumi.BoolOutput {
	return o.ApplyT(func(v NotificationPreferenceResponse) bool { return v.SendNotification }).(pulumi.BoolOutput)
}

func (o NotificationPreferenceResponseOutput) StageName() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationPreferenceResponse) string { return v.StageName }).(pulumi.StringOutput)
}

type NotificationPreferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (NotificationPreferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationPreferenceResponse)(nil)).Elem()
}

func (o NotificationPreferenceResponseArrayOutput) ToNotificationPreferenceResponseArrayOutput() NotificationPreferenceResponseArrayOutput {
	return o
}

func (o NotificationPreferenceResponseArrayOutput) ToNotificationPreferenceResponseArrayOutputWithContext(ctx context.Context) NotificationPreferenceResponseArrayOutput {
	return o
}

func (o NotificationPreferenceResponseArrayOutput) Index(i pulumi.IntInput) NotificationPreferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationPreferenceResponse {
		return vs[0].([]NotificationPreferenceResponse)[vs[1].(int)]
	}).(NotificationPreferenceResponseOutput)
}

type PackageCarrierDetails struct {
	CarrierAccountNumber *string `pulumi:"carrierAccountNumber"`
	CarrierName          *string `pulumi:"carrierName"`
	TrackingId           *string `pulumi:"trackingId"`
}





type PackageCarrierDetailsInput interface {
	pulumi.Input

	ToPackageCarrierDetailsOutput() PackageCarrierDetailsOutput
	ToPackageCarrierDetailsOutputWithContext(context.Context) PackageCarrierDetailsOutput
}

type PackageCarrierDetailsArgs struct {
	CarrierAccountNumber pulumi.StringPtrInput `pulumi:"carrierAccountNumber"`
	CarrierName          pulumi.StringPtrInput `pulumi:"carrierName"`
	TrackingId           pulumi.StringPtrInput `pulumi:"trackingId"`
}

func (PackageCarrierDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageCarrierDetails)(nil)).Elem()
}

func (i PackageCarrierDetailsArgs) ToPackageCarrierDetailsOutput() PackageCarrierDetailsOutput {
	return i.ToPackageCarrierDetailsOutputWithContext(context.Background())
}

func (i PackageCarrierDetailsArgs) ToPackageCarrierDetailsOutputWithContext(ctx context.Context) PackageCarrierDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageCarrierDetailsOutput)
}

type PackageCarrierDetailsOutput struct{ *pulumi.OutputState }

func (PackageCarrierDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageCarrierDetails)(nil)).Elem()
}

func (o PackageCarrierDetailsOutput) ToPackageCarrierDetailsOutput() PackageCarrierDetailsOutput {
	return o
}

func (o PackageCarrierDetailsOutput) ToPackageCarrierDetailsOutputWithContext(ctx context.Context) PackageCarrierDetailsOutput {
	return o
}

func (o PackageCarrierDetailsOutput) CarrierAccountNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageCarrierDetails) *string { return v.CarrierAccountNumber }).(pulumi.StringPtrOutput)
}

func (o PackageCarrierDetailsOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageCarrierDetails) *string { return v.CarrierName }).(pulumi.StringPtrOutput)
}

func (o PackageCarrierDetailsOutput) TrackingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageCarrierDetails) *string { return v.TrackingId }).(pulumi.StringPtrOutput)
}

type PackageCarrierDetailsResponse struct {
	CarrierAccountNumber *string `pulumi:"carrierAccountNumber"`
	CarrierName          *string `pulumi:"carrierName"`
	TrackingId           *string `pulumi:"trackingId"`
}





type PackageCarrierDetailsResponseInput interface {
	pulumi.Input

	ToPackageCarrierDetailsResponseOutput() PackageCarrierDetailsResponseOutput
	ToPackageCarrierDetailsResponseOutputWithContext(context.Context) PackageCarrierDetailsResponseOutput
}

type PackageCarrierDetailsResponseArgs struct {
	CarrierAccountNumber pulumi.StringPtrInput `pulumi:"carrierAccountNumber"`
	CarrierName          pulumi.StringPtrInput `pulumi:"carrierName"`
	TrackingId           pulumi.StringPtrInput `pulumi:"trackingId"`
}

func (PackageCarrierDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageCarrierDetailsResponse)(nil)).Elem()
}

func (i PackageCarrierDetailsResponseArgs) ToPackageCarrierDetailsResponseOutput() PackageCarrierDetailsResponseOutput {
	return i.ToPackageCarrierDetailsResponseOutputWithContext(context.Background())
}

func (i PackageCarrierDetailsResponseArgs) ToPackageCarrierDetailsResponseOutputWithContext(ctx context.Context) PackageCarrierDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageCarrierDetailsResponseOutput)
}

type PackageCarrierDetailsResponseOutput struct{ *pulumi.OutputState }

func (PackageCarrierDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageCarrierDetailsResponse)(nil)).Elem()
}

func (o PackageCarrierDetailsResponseOutput) ToPackageCarrierDetailsResponseOutput() PackageCarrierDetailsResponseOutput {
	return o
}

func (o PackageCarrierDetailsResponseOutput) ToPackageCarrierDetailsResponseOutputWithContext(ctx context.Context) PackageCarrierDetailsResponseOutput {
	return o
}

func (o PackageCarrierDetailsResponseOutput) CarrierAccountNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageCarrierDetailsResponse) *string { return v.CarrierAccountNumber }).(pulumi.StringPtrOutput)
}

func (o PackageCarrierDetailsResponseOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageCarrierDetailsResponse) *string { return v.CarrierName }).(pulumi.StringPtrOutput)
}

func (o PackageCarrierDetailsResponseOutput) TrackingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageCarrierDetailsResponse) *string { return v.TrackingId }).(pulumi.StringPtrOutput)
}

type PackageCarrierInfoResponse struct {
	CarrierName *string `pulumi:"carrierName"`
	TrackingId  *string `pulumi:"trackingId"`
}





type PackageCarrierInfoResponseInput interface {
	pulumi.Input

	ToPackageCarrierInfoResponseOutput() PackageCarrierInfoResponseOutput
	ToPackageCarrierInfoResponseOutputWithContext(context.Context) PackageCarrierInfoResponseOutput
}

type PackageCarrierInfoResponseArgs struct {
	CarrierName pulumi.StringPtrInput `pulumi:"carrierName"`
	TrackingId  pulumi.StringPtrInput `pulumi:"trackingId"`
}

func (PackageCarrierInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageCarrierInfoResponse)(nil)).Elem()
}

func (i PackageCarrierInfoResponseArgs) ToPackageCarrierInfoResponseOutput() PackageCarrierInfoResponseOutput {
	return i.ToPackageCarrierInfoResponseOutputWithContext(context.Background())
}

func (i PackageCarrierInfoResponseArgs) ToPackageCarrierInfoResponseOutputWithContext(ctx context.Context) PackageCarrierInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageCarrierInfoResponseOutput)
}

type PackageCarrierInfoResponseOutput struct{ *pulumi.OutputState }

func (PackageCarrierInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageCarrierInfoResponse)(nil)).Elem()
}

func (o PackageCarrierInfoResponseOutput) ToPackageCarrierInfoResponseOutput() PackageCarrierInfoResponseOutput {
	return o
}

func (o PackageCarrierInfoResponseOutput) ToPackageCarrierInfoResponseOutputWithContext(ctx context.Context) PackageCarrierInfoResponseOutput {
	return o
}

func (o PackageCarrierInfoResponseOutput) CarrierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageCarrierInfoResponse) *string { return v.CarrierName }).(pulumi.StringPtrOutput)
}

func (o PackageCarrierInfoResponseOutput) TrackingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PackageCarrierInfoResponse) *string { return v.TrackingId }).(pulumi.StringPtrOutput)
}

type PackageShippingDetailsResponse struct {
	CarrierName string `pulumi:"carrierName"`
	TrackingId  string `pulumi:"trackingId"`
	TrackingUrl string `pulumi:"trackingUrl"`
}





type PackageShippingDetailsResponseInput interface {
	pulumi.Input

	ToPackageShippingDetailsResponseOutput() PackageShippingDetailsResponseOutput
	ToPackageShippingDetailsResponseOutputWithContext(context.Context) PackageShippingDetailsResponseOutput
}

type PackageShippingDetailsResponseArgs struct {
	CarrierName pulumi.StringInput `pulumi:"carrierName"`
	TrackingId  pulumi.StringInput `pulumi:"trackingId"`
	TrackingUrl pulumi.StringInput `pulumi:"trackingUrl"`
}

func (PackageShippingDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageShippingDetailsResponse)(nil)).Elem()
}

func (i PackageShippingDetailsResponseArgs) ToPackageShippingDetailsResponseOutput() PackageShippingDetailsResponseOutput {
	return i.ToPackageShippingDetailsResponseOutputWithContext(context.Background())
}

func (i PackageShippingDetailsResponseArgs) ToPackageShippingDetailsResponseOutputWithContext(ctx context.Context) PackageShippingDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageShippingDetailsResponseOutput)
}

type PackageShippingDetailsResponseOutput struct{ *pulumi.OutputState }

func (PackageShippingDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackageShippingDetailsResponse)(nil)).Elem()
}

func (o PackageShippingDetailsResponseOutput) ToPackageShippingDetailsResponseOutput() PackageShippingDetailsResponseOutput {
	return o
}

func (o PackageShippingDetailsResponseOutput) ToPackageShippingDetailsResponseOutputWithContext(ctx context.Context) PackageShippingDetailsResponseOutput {
	return o
}

func (o PackageShippingDetailsResponseOutput) CarrierName() pulumi.StringOutput {
	return o.ApplyT(func(v PackageShippingDetailsResponse) string { return v.CarrierName }).(pulumi.StringOutput)
}

func (o PackageShippingDetailsResponseOutput) TrackingId() pulumi.StringOutput {
	return o.ApplyT(func(v PackageShippingDetailsResponse) string { return v.TrackingId }).(pulumi.StringOutput)
}

func (o PackageShippingDetailsResponseOutput) TrackingUrl() pulumi.StringOutput {
	return o.ApplyT(func(v PackageShippingDetailsResponse) string { return v.TrackingUrl }).(pulumi.StringOutput)
}

type Preferences struct {
	EncryptionPreferences     *EncryptionPreferences `pulumi:"encryptionPreferences"`
	PreferredDataCenterRegion []string               `pulumi:"preferredDataCenterRegion"`
	TransportPreferences      *TransportPreferences  `pulumi:"transportPreferences"`
}





type PreferencesInput interface {
	pulumi.Input

	ToPreferencesOutput() PreferencesOutput
	ToPreferencesOutputWithContext(context.Context) PreferencesOutput
}

type PreferencesArgs struct {
	EncryptionPreferences     EncryptionPreferencesPtrInput `pulumi:"encryptionPreferences"`
	PreferredDataCenterRegion pulumi.StringArrayInput       `pulumi:"preferredDataCenterRegion"`
	TransportPreferences      TransportPreferencesPtrInput  `pulumi:"transportPreferences"`
}

func (PreferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Preferences)(nil)).Elem()
}

func (i PreferencesArgs) ToPreferencesOutput() PreferencesOutput {
	return i.ToPreferencesOutputWithContext(context.Background())
}

func (i PreferencesArgs) ToPreferencesOutputWithContext(ctx context.Context) PreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesOutput)
}

func (i PreferencesArgs) ToPreferencesPtrOutput() PreferencesPtrOutput {
	return i.ToPreferencesPtrOutputWithContext(context.Background())
}

func (i PreferencesArgs) ToPreferencesPtrOutputWithContext(ctx context.Context) PreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesOutput).ToPreferencesPtrOutputWithContext(ctx)
}









type PreferencesPtrInput interface {
	pulumi.Input

	ToPreferencesPtrOutput() PreferencesPtrOutput
	ToPreferencesPtrOutputWithContext(context.Context) PreferencesPtrOutput
}

type preferencesPtrType PreferencesArgs

func PreferencesPtr(v *PreferencesArgs) PreferencesPtrInput {
	return (*preferencesPtrType)(v)
}

func (*preferencesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Preferences)(nil)).Elem()
}

func (i *preferencesPtrType) ToPreferencesPtrOutput() PreferencesPtrOutput {
	return i.ToPreferencesPtrOutputWithContext(context.Background())
}

func (i *preferencesPtrType) ToPreferencesPtrOutputWithContext(ctx context.Context) PreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesPtrOutput)
}

type PreferencesOutput struct{ *pulumi.OutputState }

func (PreferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Preferences)(nil)).Elem()
}

func (o PreferencesOutput) ToPreferencesOutput() PreferencesOutput {
	return o
}

func (o PreferencesOutput) ToPreferencesOutputWithContext(ctx context.Context) PreferencesOutput {
	return o
}

func (o PreferencesOutput) ToPreferencesPtrOutput() PreferencesPtrOutput {
	return o.ToPreferencesPtrOutputWithContext(context.Background())
}

func (o PreferencesOutput) ToPreferencesPtrOutputWithContext(ctx context.Context) PreferencesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Preferences) *Preferences {
		return &v
	}).(PreferencesPtrOutput)
}

func (o PreferencesOutput) EncryptionPreferences() EncryptionPreferencesPtrOutput {
	return o.ApplyT(func(v Preferences) *EncryptionPreferences { return v.EncryptionPreferences }).(EncryptionPreferencesPtrOutput)
}

func (o PreferencesOutput) PreferredDataCenterRegion() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Preferences) []string { return v.PreferredDataCenterRegion }).(pulumi.StringArrayOutput)
}

func (o PreferencesOutput) TransportPreferences() TransportPreferencesPtrOutput {
	return o.ApplyT(func(v Preferences) *TransportPreferences { return v.TransportPreferences }).(TransportPreferencesPtrOutput)
}

type PreferencesPtrOutput struct{ *pulumi.OutputState }

func (PreferencesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Preferences)(nil)).Elem()
}

func (o PreferencesPtrOutput) ToPreferencesPtrOutput() PreferencesPtrOutput {
	return o
}

func (o PreferencesPtrOutput) ToPreferencesPtrOutputWithContext(ctx context.Context) PreferencesPtrOutput {
	return o
}

func (o PreferencesPtrOutput) Elem() PreferencesOutput {
	return o.ApplyT(func(v *Preferences) Preferences {
		if v != nil {
			return *v
		}
		var ret Preferences
		return ret
	}).(PreferencesOutput)
}

func (o PreferencesPtrOutput) EncryptionPreferences() EncryptionPreferencesPtrOutput {
	return o.ApplyT(func(v *Preferences) *EncryptionPreferences {
		if v == nil {
			return nil
		}
		return v.EncryptionPreferences
	}).(EncryptionPreferencesPtrOutput)
}

func (o PreferencesPtrOutput) PreferredDataCenterRegion() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Preferences) []string {
		if v == nil {
			return nil
		}
		return v.PreferredDataCenterRegion
	}).(pulumi.StringArrayOutput)
}

func (o PreferencesPtrOutput) TransportPreferences() TransportPreferencesPtrOutput {
	return o.ApplyT(func(v *Preferences) *TransportPreferences {
		if v == nil {
			return nil
		}
		return v.TransportPreferences
	}).(TransportPreferencesPtrOutput)
}

type PreferencesResponse struct {
	EncryptionPreferences     *EncryptionPreferencesResponse `pulumi:"encryptionPreferences"`
	PreferredDataCenterRegion []string                       `pulumi:"preferredDataCenterRegion"`
	TransportPreferences      *TransportPreferencesResponse  `pulumi:"transportPreferences"`
}





type PreferencesResponseInput interface {
	pulumi.Input

	ToPreferencesResponseOutput() PreferencesResponseOutput
	ToPreferencesResponseOutputWithContext(context.Context) PreferencesResponseOutput
}

type PreferencesResponseArgs struct {
	EncryptionPreferences     EncryptionPreferencesResponsePtrInput `pulumi:"encryptionPreferences"`
	PreferredDataCenterRegion pulumi.StringArrayInput               `pulumi:"preferredDataCenterRegion"`
	TransportPreferences      TransportPreferencesResponsePtrInput  `pulumi:"transportPreferences"`
}

func (PreferencesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PreferencesResponse)(nil)).Elem()
}

func (i PreferencesResponseArgs) ToPreferencesResponseOutput() PreferencesResponseOutput {
	return i.ToPreferencesResponseOutputWithContext(context.Background())
}

func (i PreferencesResponseArgs) ToPreferencesResponseOutputWithContext(ctx context.Context) PreferencesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesResponseOutput)
}

func (i PreferencesResponseArgs) ToPreferencesResponsePtrOutput() PreferencesResponsePtrOutput {
	return i.ToPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i PreferencesResponseArgs) ToPreferencesResponsePtrOutputWithContext(ctx context.Context) PreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesResponseOutput).ToPreferencesResponsePtrOutputWithContext(ctx)
}









type PreferencesResponsePtrInput interface {
	pulumi.Input

	ToPreferencesResponsePtrOutput() PreferencesResponsePtrOutput
	ToPreferencesResponsePtrOutputWithContext(context.Context) PreferencesResponsePtrOutput
}

type preferencesResponsePtrType PreferencesResponseArgs

func PreferencesResponsePtr(v *PreferencesResponseArgs) PreferencesResponsePtrInput {
	return (*preferencesResponsePtrType)(v)
}

func (*preferencesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PreferencesResponse)(nil)).Elem()
}

func (i *preferencesResponsePtrType) ToPreferencesResponsePtrOutput() PreferencesResponsePtrOutput {
	return i.ToPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i *preferencesResponsePtrType) ToPreferencesResponsePtrOutputWithContext(ctx context.Context) PreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreferencesResponsePtrOutput)
}

type PreferencesResponseOutput struct{ *pulumi.OutputState }

func (PreferencesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PreferencesResponse)(nil)).Elem()
}

func (o PreferencesResponseOutput) ToPreferencesResponseOutput() PreferencesResponseOutput {
	return o
}

func (o PreferencesResponseOutput) ToPreferencesResponseOutputWithContext(ctx context.Context) PreferencesResponseOutput {
	return o
}

func (o PreferencesResponseOutput) ToPreferencesResponsePtrOutput() PreferencesResponsePtrOutput {
	return o.ToPreferencesResponsePtrOutputWithContext(context.Background())
}

func (o PreferencesResponseOutput) ToPreferencesResponsePtrOutputWithContext(ctx context.Context) PreferencesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PreferencesResponse) *PreferencesResponse {
		return &v
	}).(PreferencesResponsePtrOutput)
}

func (o PreferencesResponseOutput) EncryptionPreferences() EncryptionPreferencesResponsePtrOutput {
	return o.ApplyT(func(v PreferencesResponse) *EncryptionPreferencesResponse { return v.EncryptionPreferences }).(EncryptionPreferencesResponsePtrOutput)
}

func (o PreferencesResponseOutput) PreferredDataCenterRegion() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PreferencesResponse) []string { return v.PreferredDataCenterRegion }).(pulumi.StringArrayOutput)
}

func (o PreferencesResponseOutput) TransportPreferences() TransportPreferencesResponsePtrOutput {
	return o.ApplyT(func(v PreferencesResponse) *TransportPreferencesResponse { return v.TransportPreferences }).(TransportPreferencesResponsePtrOutput)
}

type PreferencesResponsePtrOutput struct{ *pulumi.OutputState }

func (PreferencesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PreferencesResponse)(nil)).Elem()
}

func (o PreferencesResponsePtrOutput) ToPreferencesResponsePtrOutput() PreferencesResponsePtrOutput {
	return o
}

func (o PreferencesResponsePtrOutput) ToPreferencesResponsePtrOutputWithContext(ctx context.Context) PreferencesResponsePtrOutput {
	return o
}

func (o PreferencesResponsePtrOutput) Elem() PreferencesResponseOutput {
	return o.ApplyT(func(v *PreferencesResponse) PreferencesResponse {
		if v != nil {
			return *v
		}
		var ret PreferencesResponse
		return ret
	}).(PreferencesResponseOutput)
}

func (o PreferencesResponsePtrOutput) EncryptionPreferences() EncryptionPreferencesResponsePtrOutput {
	return o.ApplyT(func(v *PreferencesResponse) *EncryptionPreferencesResponse {
		if v == nil {
			return nil
		}
		return v.EncryptionPreferences
	}).(EncryptionPreferencesResponsePtrOutput)
}

func (o PreferencesResponsePtrOutput) PreferredDataCenterRegion() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PreferencesResponse) []string {
		if v == nil {
			return nil
		}
		return v.PreferredDataCenterRegion
	}).(pulumi.StringArrayOutput)
}

func (o PreferencesResponsePtrOutput) TransportPreferences() TransportPreferencesResponsePtrOutput {
	return o.ApplyT(func(v *PreferencesResponse) *TransportPreferencesResponse {
		if v == nil {
			return nil
		}
		return v.TransportPreferences
	}).(TransportPreferencesResponsePtrOutput)
}

type ResourceIdentity struct {
	Type                   *string                `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	Type                   pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput       `pulumi:"userAssignedIdentities"`
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

func (o ResourceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ResourceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
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

func (o ResourceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ResourceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ResourceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}





type ResourceIdentityResponseInput interface {
	pulumi.Input

	ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput
	ToResourceIdentityResponseOutputWithContext(context.Context) ResourceIdentityResponseOutput
}

type ResourceIdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                   `pulumi:"principalId"`
	TenantId               pulumi.StringInput                   `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (ResourceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return i.ToResourceIdentityResponseOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput)
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput).ToResourceIdentityResponsePtrOutputWithContext(ctx)
}









type ResourceIdentityResponsePtrInput interface {
	pulumi.Input

	ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput
	ToResourceIdentityResponsePtrOutputWithContext(context.Context) ResourceIdentityResponsePtrOutput
}

type resourceIdentityResponsePtrType ResourceIdentityResponseArgs

func ResourceIdentityResponsePtr(v *ResourceIdentityResponseArgs) ResourceIdentityResponsePtrInput {
	return (*resourceIdentityResponsePtrType)(v)
}

func (*resourceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponsePtrOutput)
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

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityResponse) *ResourceIdentityResponse {
		return &v
	}).(ResourceIdentityResponsePtrOutput)
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

func (o ResourceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
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

func (o ResourceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ShareCredentialDetailsResponse struct {
	Password                 string   `pulumi:"password"`
	ShareName                string   `pulumi:"shareName"`
	ShareType                string   `pulumi:"shareType"`
	SupportedAccessProtocols []string `pulumi:"supportedAccessProtocols"`
	UserName                 string   `pulumi:"userName"`
}





type ShareCredentialDetailsResponseInput interface {
	pulumi.Input

	ToShareCredentialDetailsResponseOutput() ShareCredentialDetailsResponseOutput
	ToShareCredentialDetailsResponseOutputWithContext(context.Context) ShareCredentialDetailsResponseOutput
}

type ShareCredentialDetailsResponseArgs struct {
	Password                 pulumi.StringInput      `pulumi:"password"`
	ShareName                pulumi.StringInput      `pulumi:"shareName"`
	ShareType                pulumi.StringInput      `pulumi:"shareType"`
	SupportedAccessProtocols pulumi.StringArrayInput `pulumi:"supportedAccessProtocols"`
	UserName                 pulumi.StringInput      `pulumi:"userName"`
}

func (ShareCredentialDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareCredentialDetailsResponse)(nil)).Elem()
}

func (i ShareCredentialDetailsResponseArgs) ToShareCredentialDetailsResponseOutput() ShareCredentialDetailsResponseOutput {
	return i.ToShareCredentialDetailsResponseOutputWithContext(context.Background())
}

func (i ShareCredentialDetailsResponseArgs) ToShareCredentialDetailsResponseOutputWithContext(ctx context.Context) ShareCredentialDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareCredentialDetailsResponseOutput)
}





type ShareCredentialDetailsResponseArrayInput interface {
	pulumi.Input

	ToShareCredentialDetailsResponseArrayOutput() ShareCredentialDetailsResponseArrayOutput
	ToShareCredentialDetailsResponseArrayOutputWithContext(context.Context) ShareCredentialDetailsResponseArrayOutput
}

type ShareCredentialDetailsResponseArray []ShareCredentialDetailsResponseInput

func (ShareCredentialDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareCredentialDetailsResponse)(nil)).Elem()
}

func (i ShareCredentialDetailsResponseArray) ToShareCredentialDetailsResponseArrayOutput() ShareCredentialDetailsResponseArrayOutput {
	return i.ToShareCredentialDetailsResponseArrayOutputWithContext(context.Background())
}

func (i ShareCredentialDetailsResponseArray) ToShareCredentialDetailsResponseArrayOutputWithContext(ctx context.Context) ShareCredentialDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareCredentialDetailsResponseArrayOutput)
}

type ShareCredentialDetailsResponseOutput struct{ *pulumi.OutputState }

func (ShareCredentialDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareCredentialDetailsResponse)(nil)).Elem()
}

func (o ShareCredentialDetailsResponseOutput) ToShareCredentialDetailsResponseOutput() ShareCredentialDetailsResponseOutput {
	return o
}

func (o ShareCredentialDetailsResponseOutput) ToShareCredentialDetailsResponseOutputWithContext(ctx context.Context) ShareCredentialDetailsResponseOutput {
	return o
}

func (o ShareCredentialDetailsResponseOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v ShareCredentialDetailsResponse) string { return v.Password }).(pulumi.StringOutput)
}

func (o ShareCredentialDetailsResponseOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v ShareCredentialDetailsResponse) string { return v.ShareName }).(pulumi.StringOutput)
}

func (o ShareCredentialDetailsResponseOutput) ShareType() pulumi.StringOutput {
	return o.ApplyT(func(v ShareCredentialDetailsResponse) string { return v.ShareType }).(pulumi.StringOutput)
}

func (o ShareCredentialDetailsResponseOutput) SupportedAccessProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ShareCredentialDetailsResponse) []string { return v.SupportedAccessProtocols }).(pulumi.StringArrayOutput)
}

func (o ShareCredentialDetailsResponseOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v ShareCredentialDetailsResponse) string { return v.UserName }).(pulumi.StringOutput)
}

type ShareCredentialDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (ShareCredentialDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareCredentialDetailsResponse)(nil)).Elem()
}

func (o ShareCredentialDetailsResponseArrayOutput) ToShareCredentialDetailsResponseArrayOutput() ShareCredentialDetailsResponseArrayOutput {
	return o
}

func (o ShareCredentialDetailsResponseArrayOutput) ToShareCredentialDetailsResponseArrayOutputWithContext(ctx context.Context) ShareCredentialDetailsResponseArrayOutput {
	return o
}

func (o ShareCredentialDetailsResponseArrayOutput) Index(i pulumi.IntInput) ShareCredentialDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ShareCredentialDetailsResponse {
		return vs[0].([]ShareCredentialDetailsResponse)[vs[1].(int)]
	}).(ShareCredentialDetailsResponseOutput)
}

type ShippingAddress struct {
	AddressType     *string `pulumi:"addressType"`
	City            *string `pulumi:"city"`
	CompanyName     *string `pulumi:"companyName"`
	Country         string  `pulumi:"country"`
	PostalCode      *string `pulumi:"postalCode"`
	StateOrProvince *string `pulumi:"stateOrProvince"`
	StreetAddress1  string  `pulumi:"streetAddress1"`
	StreetAddress2  *string `pulumi:"streetAddress2"`
	StreetAddress3  *string `pulumi:"streetAddress3"`
	ZipExtendedCode *string `pulumi:"zipExtendedCode"`
}





type ShippingAddressInput interface {
	pulumi.Input

	ToShippingAddressOutput() ShippingAddressOutput
	ToShippingAddressOutputWithContext(context.Context) ShippingAddressOutput
}

type ShippingAddressArgs struct {
	AddressType     pulumi.StringPtrInput `pulumi:"addressType"`
	City            pulumi.StringPtrInput `pulumi:"city"`
	CompanyName     pulumi.StringPtrInput `pulumi:"companyName"`
	Country         pulumi.StringInput    `pulumi:"country"`
	PostalCode      pulumi.StringPtrInput `pulumi:"postalCode"`
	StateOrProvince pulumi.StringPtrInput `pulumi:"stateOrProvince"`
	StreetAddress1  pulumi.StringInput    `pulumi:"streetAddress1"`
	StreetAddress2  pulumi.StringPtrInput `pulumi:"streetAddress2"`
	StreetAddress3  pulumi.StringPtrInput `pulumi:"streetAddress3"`
	ZipExtendedCode pulumi.StringPtrInput `pulumi:"zipExtendedCode"`
}

func (ShippingAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingAddress)(nil)).Elem()
}

func (i ShippingAddressArgs) ToShippingAddressOutput() ShippingAddressOutput {
	return i.ToShippingAddressOutputWithContext(context.Background())
}

func (i ShippingAddressArgs) ToShippingAddressOutputWithContext(ctx context.Context) ShippingAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressOutput)
}

func (i ShippingAddressArgs) ToShippingAddressPtrOutput() ShippingAddressPtrOutput {
	return i.ToShippingAddressPtrOutputWithContext(context.Background())
}

func (i ShippingAddressArgs) ToShippingAddressPtrOutputWithContext(ctx context.Context) ShippingAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressOutput).ToShippingAddressPtrOutputWithContext(ctx)
}









type ShippingAddressPtrInput interface {
	pulumi.Input

	ToShippingAddressPtrOutput() ShippingAddressPtrOutput
	ToShippingAddressPtrOutputWithContext(context.Context) ShippingAddressPtrOutput
}

type shippingAddressPtrType ShippingAddressArgs

func ShippingAddressPtr(v *ShippingAddressArgs) ShippingAddressPtrInput {
	return (*shippingAddressPtrType)(v)
}

func (*shippingAddressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingAddress)(nil)).Elem()
}

func (i *shippingAddressPtrType) ToShippingAddressPtrOutput() ShippingAddressPtrOutput {
	return i.ToShippingAddressPtrOutputWithContext(context.Background())
}

func (i *shippingAddressPtrType) ToShippingAddressPtrOutputWithContext(ctx context.Context) ShippingAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressPtrOutput)
}

type ShippingAddressOutput struct{ *pulumi.OutputState }

func (ShippingAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingAddress)(nil)).Elem()
}

func (o ShippingAddressOutput) ToShippingAddressOutput() ShippingAddressOutput {
	return o
}

func (o ShippingAddressOutput) ToShippingAddressOutputWithContext(ctx context.Context) ShippingAddressOutput {
	return o
}

func (o ShippingAddressOutput) ToShippingAddressPtrOutput() ShippingAddressPtrOutput {
	return o.ToShippingAddressPtrOutputWithContext(context.Background())
}

func (o ShippingAddressOutput) ToShippingAddressPtrOutputWithContext(ctx context.Context) ShippingAddressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShippingAddress) *ShippingAddress {
		return &v
	}).(ShippingAddressPtrOutput)
}

func (o ShippingAddressOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.AddressType }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.CompanyName }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingAddress) string { return v.Country }).(pulumi.StringOutput)
}

func (o ShippingAddressOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.PostalCode }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.StateOrProvince }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) StreetAddress1() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingAddress) string { return v.StreetAddress1 }).(pulumi.StringOutput)
}

func (o ShippingAddressOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.StreetAddress2 }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) StreetAddress3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.StreetAddress3 }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressOutput) ZipExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddress) *string { return v.ZipExtendedCode }).(pulumi.StringPtrOutput)
}

type ShippingAddressPtrOutput struct{ *pulumi.OutputState }

func (ShippingAddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingAddress)(nil)).Elem()
}

func (o ShippingAddressPtrOutput) ToShippingAddressPtrOutput() ShippingAddressPtrOutput {
	return o
}

func (o ShippingAddressPtrOutput) ToShippingAddressPtrOutputWithContext(ctx context.Context) ShippingAddressPtrOutput {
	return o
}

func (o ShippingAddressPtrOutput) Elem() ShippingAddressOutput {
	return o.ApplyT(func(v *ShippingAddress) ShippingAddress {
		if v != nil {
			return *v
		}
		var ret ShippingAddress
		return ret
	}).(ShippingAddressOutput)
}

func (o ShippingAddressPtrOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.AddressType
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.CompanyName
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return &v.Country
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.PostalCode
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.StateOrProvince
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) StreetAddress1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return &v.StreetAddress1
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress2
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) StreetAddress3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress3
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressPtrOutput) ZipExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddress) *string {
		if v == nil {
			return nil
		}
		return v.ZipExtendedCode
	}).(pulumi.StringPtrOutput)
}

type ShippingAddressResponse struct {
	AddressType     *string `pulumi:"addressType"`
	City            *string `pulumi:"city"`
	CompanyName     *string `pulumi:"companyName"`
	Country         string  `pulumi:"country"`
	PostalCode      *string `pulumi:"postalCode"`
	StateOrProvince *string `pulumi:"stateOrProvince"`
	StreetAddress1  string  `pulumi:"streetAddress1"`
	StreetAddress2  *string `pulumi:"streetAddress2"`
	StreetAddress3  *string `pulumi:"streetAddress3"`
	ZipExtendedCode *string `pulumi:"zipExtendedCode"`
}





type ShippingAddressResponseInput interface {
	pulumi.Input

	ToShippingAddressResponseOutput() ShippingAddressResponseOutput
	ToShippingAddressResponseOutputWithContext(context.Context) ShippingAddressResponseOutput
}

type ShippingAddressResponseArgs struct {
	AddressType     pulumi.StringPtrInput `pulumi:"addressType"`
	City            pulumi.StringPtrInput `pulumi:"city"`
	CompanyName     pulumi.StringPtrInput `pulumi:"companyName"`
	Country         pulumi.StringInput    `pulumi:"country"`
	PostalCode      pulumi.StringPtrInput `pulumi:"postalCode"`
	StateOrProvince pulumi.StringPtrInput `pulumi:"stateOrProvince"`
	StreetAddress1  pulumi.StringInput    `pulumi:"streetAddress1"`
	StreetAddress2  pulumi.StringPtrInput `pulumi:"streetAddress2"`
	StreetAddress3  pulumi.StringPtrInput `pulumi:"streetAddress3"`
	ZipExtendedCode pulumi.StringPtrInput `pulumi:"zipExtendedCode"`
}

func (ShippingAddressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingAddressResponse)(nil)).Elem()
}

func (i ShippingAddressResponseArgs) ToShippingAddressResponseOutput() ShippingAddressResponseOutput {
	return i.ToShippingAddressResponseOutputWithContext(context.Background())
}

func (i ShippingAddressResponseArgs) ToShippingAddressResponseOutputWithContext(ctx context.Context) ShippingAddressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressResponseOutput)
}

func (i ShippingAddressResponseArgs) ToShippingAddressResponsePtrOutput() ShippingAddressResponsePtrOutput {
	return i.ToShippingAddressResponsePtrOutputWithContext(context.Background())
}

func (i ShippingAddressResponseArgs) ToShippingAddressResponsePtrOutputWithContext(ctx context.Context) ShippingAddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressResponseOutput).ToShippingAddressResponsePtrOutputWithContext(ctx)
}









type ShippingAddressResponsePtrInput interface {
	pulumi.Input

	ToShippingAddressResponsePtrOutput() ShippingAddressResponsePtrOutput
	ToShippingAddressResponsePtrOutputWithContext(context.Context) ShippingAddressResponsePtrOutput
}

type shippingAddressResponsePtrType ShippingAddressResponseArgs

func ShippingAddressResponsePtr(v *ShippingAddressResponseArgs) ShippingAddressResponsePtrInput {
	return (*shippingAddressResponsePtrType)(v)
}

func (*shippingAddressResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingAddressResponse)(nil)).Elem()
}

func (i *shippingAddressResponsePtrType) ToShippingAddressResponsePtrOutput() ShippingAddressResponsePtrOutput {
	return i.ToShippingAddressResponsePtrOutputWithContext(context.Background())
}

func (i *shippingAddressResponsePtrType) ToShippingAddressResponsePtrOutputWithContext(ctx context.Context) ShippingAddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingAddressResponsePtrOutput)
}

type ShippingAddressResponseOutput struct{ *pulumi.OutputState }

func (ShippingAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShippingAddressResponse)(nil)).Elem()
}

func (o ShippingAddressResponseOutput) ToShippingAddressResponseOutput() ShippingAddressResponseOutput {
	return o
}

func (o ShippingAddressResponseOutput) ToShippingAddressResponseOutputWithContext(ctx context.Context) ShippingAddressResponseOutput {
	return o
}

func (o ShippingAddressResponseOutput) ToShippingAddressResponsePtrOutput() ShippingAddressResponsePtrOutput {
	return o.ToShippingAddressResponsePtrOutputWithContext(context.Background())
}

func (o ShippingAddressResponseOutput) ToShippingAddressResponsePtrOutputWithContext(ctx context.Context) ShippingAddressResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ShippingAddressResponse) *ShippingAddressResponse {
		return &v
	}).(ShippingAddressResponsePtrOutput)
}

func (o ShippingAddressResponseOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.AddressType }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.CompanyName }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingAddressResponse) string { return v.Country }).(pulumi.StringOutput)
}

func (o ShippingAddressResponseOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.PostalCode }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.StateOrProvince }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) StreetAddress1() pulumi.StringOutput {
	return o.ApplyT(func(v ShippingAddressResponse) string { return v.StreetAddress1 }).(pulumi.StringOutput)
}

func (o ShippingAddressResponseOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.StreetAddress2 }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) StreetAddress3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.StreetAddress3 }).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponseOutput) ZipExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShippingAddressResponse) *string { return v.ZipExtendedCode }).(pulumi.StringPtrOutput)
}

type ShippingAddressResponsePtrOutput struct{ *pulumi.OutputState }

func (ShippingAddressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingAddressResponse)(nil)).Elem()
}

func (o ShippingAddressResponsePtrOutput) ToShippingAddressResponsePtrOutput() ShippingAddressResponsePtrOutput {
	return o
}

func (o ShippingAddressResponsePtrOutput) ToShippingAddressResponsePtrOutputWithContext(ctx context.Context) ShippingAddressResponsePtrOutput {
	return o
}

func (o ShippingAddressResponsePtrOutput) Elem() ShippingAddressResponseOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) ShippingAddressResponse {
		if v != nil {
			return *v
		}
		var ret ShippingAddressResponse
		return ret
	}).(ShippingAddressResponseOutput)
}

func (o ShippingAddressResponsePtrOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.AddressType
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) CompanyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.CompanyName
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Country
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.PostalCode
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) StateOrProvince() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.StateOrProvince
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) StreetAddress1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StreetAddress1
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) StreetAddress2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress2
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) StreetAddress3() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.StreetAddress3
	}).(pulumi.StringPtrOutput)
}

func (o ShippingAddressResponsePtrOutput) ZipExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.ZipExtendedCode
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	DisplayName *string `pulumi:"displayName"`
	Family      *string `pulumi:"family"`
	Name        string  `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Family      pulumi.StringPtrInput `pulumi:"family"`
	Name        pulumi.StringInput    `pulumi:"name"`
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

func (o SkuOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
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

type SkuResponse struct {
	DisplayName *string `pulumi:"displayName"`
	Family      *string `pulumi:"family"`
	Name        string  `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Family      pulumi.StringPtrInput `pulumi:"family"`
	Name        pulumi.StringInput    `pulumi:"name"`
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

func (o SkuResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
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

func (o SkuResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
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

type StorageAccountDetails struct {
	DataAccountType  string  `pulumi:"dataAccountType"`
	SharePassword    *string `pulumi:"sharePassword"`
	StorageAccountId string  `pulumi:"storageAccountId"`
}





type StorageAccountDetailsInput interface {
	pulumi.Input

	ToStorageAccountDetailsOutput() StorageAccountDetailsOutput
	ToStorageAccountDetailsOutputWithContext(context.Context) StorageAccountDetailsOutput
}

type StorageAccountDetailsArgs struct {
	DataAccountType  pulumi.StringInput    `pulumi:"dataAccountType"`
	SharePassword    pulumi.StringPtrInput `pulumi:"sharePassword"`
	StorageAccountId pulumi.StringInput    `pulumi:"storageAccountId"`
}

func (StorageAccountDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountDetails)(nil)).Elem()
}

func (i StorageAccountDetailsArgs) ToStorageAccountDetailsOutput() StorageAccountDetailsOutput {
	return i.ToStorageAccountDetailsOutputWithContext(context.Background())
}

func (i StorageAccountDetailsArgs) ToStorageAccountDetailsOutputWithContext(ctx context.Context) StorageAccountDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountDetailsOutput)
}

type StorageAccountDetailsOutput struct{ *pulumi.OutputState }

func (StorageAccountDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountDetails)(nil)).Elem()
}

func (o StorageAccountDetailsOutput) ToStorageAccountDetailsOutput() StorageAccountDetailsOutput {
	return o
}

func (o StorageAccountDetailsOutput) ToStorageAccountDetailsOutputWithContext(ctx context.Context) StorageAccountDetailsOutput {
	return o
}

func (o StorageAccountDetailsOutput) DataAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountDetails) string { return v.DataAccountType }).(pulumi.StringOutput)
}

func (o StorageAccountDetailsOutput) SharePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountDetails) *string { return v.SharePassword }).(pulumi.StringPtrOutput)
}

func (o StorageAccountDetailsOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountDetails) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type StorageAccountDetailsResponse struct {
	DataAccountType  string `pulumi:"dataAccountType"`
	StorageAccountId string `pulumi:"storageAccountId"`
}





type StorageAccountDetailsResponseInput interface {
	pulumi.Input

	ToStorageAccountDetailsResponseOutput() StorageAccountDetailsResponseOutput
	ToStorageAccountDetailsResponseOutputWithContext(context.Context) StorageAccountDetailsResponseOutput
}

type StorageAccountDetailsResponseArgs struct {
	DataAccountType  pulumi.StringInput `pulumi:"dataAccountType"`
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
}

func (StorageAccountDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountDetailsResponse)(nil)).Elem()
}

func (i StorageAccountDetailsResponseArgs) ToStorageAccountDetailsResponseOutput() StorageAccountDetailsResponseOutput {
	return i.ToStorageAccountDetailsResponseOutputWithContext(context.Background())
}

func (i StorageAccountDetailsResponseArgs) ToStorageAccountDetailsResponseOutputWithContext(ctx context.Context) StorageAccountDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountDetailsResponseOutput)
}

type StorageAccountDetailsResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountDetailsResponse)(nil)).Elem()
}

func (o StorageAccountDetailsResponseOutput) ToStorageAccountDetailsResponseOutput() StorageAccountDetailsResponseOutput {
	return o
}

func (o StorageAccountDetailsResponseOutput) ToStorageAccountDetailsResponseOutputWithContext(ctx context.Context) StorageAccountDetailsResponseOutput {
	return o
}

func (o StorageAccountDetailsResponseOutput) DataAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountDetailsResponse) string { return v.DataAccountType }).(pulumi.StringOutput)
}

func (o StorageAccountDetailsResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountDetailsResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type SystemDataResponse struct {
	CreatedAt          string `pulumi:"createdAt"`
	CreatedBy          string `pulumi:"createdBy"`
	CreatedByType      string `pulumi:"createdByType"`
	LastModifiedAt     string `pulumi:"lastModifiedAt"`
	LastModifiedBy     string `pulumi:"lastModifiedBy"`
	LastModifiedByType string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringInput `pulumi:"lastModifiedByType"`
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

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.CreatedByType }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemDataResponse) string { return v.LastModifiedByType }).(pulumi.StringOutput)
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
		return &v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type TransferAllDetails struct {
	DataAccountType  string `pulumi:"dataAccountType"`
	TransferAllBlobs *bool  `pulumi:"transferAllBlobs"`
	TransferAllFiles *bool  `pulumi:"transferAllFiles"`
}





type TransferAllDetailsInput interface {
	pulumi.Input

	ToTransferAllDetailsOutput() TransferAllDetailsOutput
	ToTransferAllDetailsOutputWithContext(context.Context) TransferAllDetailsOutput
}

type TransferAllDetailsArgs struct {
	DataAccountType  pulumi.StringInput  `pulumi:"dataAccountType"`
	TransferAllBlobs pulumi.BoolPtrInput `pulumi:"transferAllBlobs"`
	TransferAllFiles pulumi.BoolPtrInput `pulumi:"transferAllFiles"`
}

func (TransferAllDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferAllDetails)(nil)).Elem()
}

func (i TransferAllDetailsArgs) ToTransferAllDetailsOutput() TransferAllDetailsOutput {
	return i.ToTransferAllDetailsOutputWithContext(context.Background())
}

func (i TransferAllDetailsArgs) ToTransferAllDetailsOutputWithContext(ctx context.Context) TransferAllDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAllDetailsOutput)
}

func (i TransferAllDetailsArgs) ToTransferAllDetailsPtrOutput() TransferAllDetailsPtrOutput {
	return i.ToTransferAllDetailsPtrOutputWithContext(context.Background())
}

func (i TransferAllDetailsArgs) ToTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferAllDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAllDetailsOutput).ToTransferAllDetailsPtrOutputWithContext(ctx)
}









type TransferAllDetailsPtrInput interface {
	pulumi.Input

	ToTransferAllDetailsPtrOutput() TransferAllDetailsPtrOutput
	ToTransferAllDetailsPtrOutputWithContext(context.Context) TransferAllDetailsPtrOutput
}

type transferAllDetailsPtrType TransferAllDetailsArgs

func TransferAllDetailsPtr(v *TransferAllDetailsArgs) TransferAllDetailsPtrInput {
	return (*transferAllDetailsPtrType)(v)
}

func (*transferAllDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferAllDetails)(nil)).Elem()
}

func (i *transferAllDetailsPtrType) ToTransferAllDetailsPtrOutput() TransferAllDetailsPtrOutput {
	return i.ToTransferAllDetailsPtrOutputWithContext(context.Background())
}

func (i *transferAllDetailsPtrType) ToTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferAllDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAllDetailsPtrOutput)
}

type TransferAllDetailsOutput struct{ *pulumi.OutputState }

func (TransferAllDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferAllDetails)(nil)).Elem()
}

func (o TransferAllDetailsOutput) ToTransferAllDetailsOutput() TransferAllDetailsOutput {
	return o
}

func (o TransferAllDetailsOutput) ToTransferAllDetailsOutputWithContext(ctx context.Context) TransferAllDetailsOutput {
	return o
}

func (o TransferAllDetailsOutput) ToTransferAllDetailsPtrOutput() TransferAllDetailsPtrOutput {
	return o.ToTransferAllDetailsPtrOutputWithContext(context.Background())
}

func (o TransferAllDetailsOutput) ToTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferAllDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransferAllDetails) *TransferAllDetails {
		return &v
	}).(TransferAllDetailsPtrOutput)
}

func (o TransferAllDetailsOutput) DataAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v TransferAllDetails) string { return v.DataAccountType }).(pulumi.StringOutput)
}

func (o TransferAllDetailsOutput) TransferAllBlobs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TransferAllDetails) *bool { return v.TransferAllBlobs }).(pulumi.BoolPtrOutput)
}

func (o TransferAllDetailsOutput) TransferAllFiles() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TransferAllDetails) *bool { return v.TransferAllFiles }).(pulumi.BoolPtrOutput)
}

type TransferAllDetailsPtrOutput struct{ *pulumi.OutputState }

func (TransferAllDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferAllDetails)(nil)).Elem()
}

func (o TransferAllDetailsPtrOutput) ToTransferAllDetailsPtrOutput() TransferAllDetailsPtrOutput {
	return o
}

func (o TransferAllDetailsPtrOutput) ToTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferAllDetailsPtrOutput {
	return o
}

func (o TransferAllDetailsPtrOutput) Elem() TransferAllDetailsOutput {
	return o.ApplyT(func(v *TransferAllDetails) TransferAllDetails {
		if v != nil {
			return *v
		}
		var ret TransferAllDetails
		return ret
	}).(TransferAllDetailsOutput)
}

func (o TransferAllDetailsPtrOutput) DataAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransferAllDetails) *string {
		if v == nil {
			return nil
		}
		return &v.DataAccountType
	}).(pulumi.StringPtrOutput)
}

func (o TransferAllDetailsPtrOutput) TransferAllBlobs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransferAllDetails) *bool {
		if v == nil {
			return nil
		}
		return v.TransferAllBlobs
	}).(pulumi.BoolPtrOutput)
}

func (o TransferAllDetailsPtrOutput) TransferAllFiles() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransferAllDetails) *bool {
		if v == nil {
			return nil
		}
		return v.TransferAllFiles
	}).(pulumi.BoolPtrOutput)
}

type TransferAllDetailsResponse struct {
	DataAccountType  string `pulumi:"dataAccountType"`
	TransferAllBlobs *bool  `pulumi:"transferAllBlobs"`
	TransferAllFiles *bool  `pulumi:"transferAllFiles"`
}





type TransferAllDetailsResponseInput interface {
	pulumi.Input

	ToTransferAllDetailsResponseOutput() TransferAllDetailsResponseOutput
	ToTransferAllDetailsResponseOutputWithContext(context.Context) TransferAllDetailsResponseOutput
}

type TransferAllDetailsResponseArgs struct {
	DataAccountType  pulumi.StringInput  `pulumi:"dataAccountType"`
	TransferAllBlobs pulumi.BoolPtrInput `pulumi:"transferAllBlobs"`
	TransferAllFiles pulumi.BoolPtrInput `pulumi:"transferAllFiles"`
}

func (TransferAllDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferAllDetailsResponse)(nil)).Elem()
}

func (i TransferAllDetailsResponseArgs) ToTransferAllDetailsResponseOutput() TransferAllDetailsResponseOutput {
	return i.ToTransferAllDetailsResponseOutputWithContext(context.Background())
}

func (i TransferAllDetailsResponseArgs) ToTransferAllDetailsResponseOutputWithContext(ctx context.Context) TransferAllDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAllDetailsResponseOutput)
}

func (i TransferAllDetailsResponseArgs) ToTransferAllDetailsResponsePtrOutput() TransferAllDetailsResponsePtrOutput {
	return i.ToTransferAllDetailsResponsePtrOutputWithContext(context.Background())
}

func (i TransferAllDetailsResponseArgs) ToTransferAllDetailsResponsePtrOutputWithContext(ctx context.Context) TransferAllDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAllDetailsResponseOutput).ToTransferAllDetailsResponsePtrOutputWithContext(ctx)
}









type TransferAllDetailsResponsePtrInput interface {
	pulumi.Input

	ToTransferAllDetailsResponsePtrOutput() TransferAllDetailsResponsePtrOutput
	ToTransferAllDetailsResponsePtrOutputWithContext(context.Context) TransferAllDetailsResponsePtrOutput
}

type transferAllDetailsResponsePtrType TransferAllDetailsResponseArgs

func TransferAllDetailsResponsePtr(v *TransferAllDetailsResponseArgs) TransferAllDetailsResponsePtrInput {
	return (*transferAllDetailsResponsePtrType)(v)
}

func (*transferAllDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferAllDetailsResponse)(nil)).Elem()
}

func (i *transferAllDetailsResponsePtrType) ToTransferAllDetailsResponsePtrOutput() TransferAllDetailsResponsePtrOutput {
	return i.ToTransferAllDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *transferAllDetailsResponsePtrType) ToTransferAllDetailsResponsePtrOutputWithContext(ctx context.Context) TransferAllDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAllDetailsResponsePtrOutput)
}

type TransferAllDetailsResponseOutput struct{ *pulumi.OutputState }

func (TransferAllDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferAllDetailsResponse)(nil)).Elem()
}

func (o TransferAllDetailsResponseOutput) ToTransferAllDetailsResponseOutput() TransferAllDetailsResponseOutput {
	return o
}

func (o TransferAllDetailsResponseOutput) ToTransferAllDetailsResponseOutputWithContext(ctx context.Context) TransferAllDetailsResponseOutput {
	return o
}

func (o TransferAllDetailsResponseOutput) ToTransferAllDetailsResponsePtrOutput() TransferAllDetailsResponsePtrOutput {
	return o.ToTransferAllDetailsResponsePtrOutputWithContext(context.Background())
}

func (o TransferAllDetailsResponseOutput) ToTransferAllDetailsResponsePtrOutputWithContext(ctx context.Context) TransferAllDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransferAllDetailsResponse) *TransferAllDetailsResponse {
		return &v
	}).(TransferAllDetailsResponsePtrOutput)
}

func (o TransferAllDetailsResponseOutput) DataAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v TransferAllDetailsResponse) string { return v.DataAccountType }).(pulumi.StringOutput)
}

func (o TransferAllDetailsResponseOutput) TransferAllBlobs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TransferAllDetailsResponse) *bool { return v.TransferAllBlobs }).(pulumi.BoolPtrOutput)
}

func (o TransferAllDetailsResponseOutput) TransferAllFiles() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TransferAllDetailsResponse) *bool { return v.TransferAllFiles }).(pulumi.BoolPtrOutput)
}

type TransferAllDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (TransferAllDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferAllDetailsResponse)(nil)).Elem()
}

func (o TransferAllDetailsResponsePtrOutput) ToTransferAllDetailsResponsePtrOutput() TransferAllDetailsResponsePtrOutput {
	return o
}

func (o TransferAllDetailsResponsePtrOutput) ToTransferAllDetailsResponsePtrOutputWithContext(ctx context.Context) TransferAllDetailsResponsePtrOutput {
	return o
}

func (o TransferAllDetailsResponsePtrOutput) Elem() TransferAllDetailsResponseOutput {
	return o.ApplyT(func(v *TransferAllDetailsResponse) TransferAllDetailsResponse {
		if v != nil {
			return *v
		}
		var ret TransferAllDetailsResponse
		return ret
	}).(TransferAllDetailsResponseOutput)
}

func (o TransferAllDetailsResponsePtrOutput) DataAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransferAllDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataAccountType
	}).(pulumi.StringPtrOutput)
}

func (o TransferAllDetailsResponsePtrOutput) TransferAllBlobs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransferAllDetailsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.TransferAllBlobs
	}).(pulumi.BoolPtrOutput)
}

func (o TransferAllDetailsResponsePtrOutput) TransferAllFiles() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransferAllDetailsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.TransferAllFiles
	}).(pulumi.BoolPtrOutput)
}

type TransferConfiguration struct {
	TransferAllDetails        *TransferConfigurationTransferAllDetails    `pulumi:"transferAllDetails"`
	TransferConfigurationType string                                      `pulumi:"transferConfigurationType"`
	TransferFilterDetails     *TransferConfigurationTransferFilterDetails `pulumi:"transferFilterDetails"`
}





type TransferConfigurationInput interface {
	pulumi.Input

	ToTransferConfigurationOutput() TransferConfigurationOutput
	ToTransferConfigurationOutputWithContext(context.Context) TransferConfigurationOutput
}

type TransferConfigurationArgs struct {
	TransferAllDetails        TransferConfigurationTransferAllDetailsPtrInput    `pulumi:"transferAllDetails"`
	TransferConfigurationType pulumi.StringInput                                 `pulumi:"transferConfigurationType"`
	TransferFilterDetails     TransferConfigurationTransferFilterDetailsPtrInput `pulumi:"transferFilterDetails"`
}

func (TransferConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfiguration)(nil)).Elem()
}

func (i TransferConfigurationArgs) ToTransferConfigurationOutput() TransferConfigurationOutput {
	return i.ToTransferConfigurationOutputWithContext(context.Background())
}

func (i TransferConfigurationArgs) ToTransferConfigurationOutputWithContext(ctx context.Context) TransferConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationOutput)
}

type TransferConfigurationOutput struct{ *pulumi.OutputState }

func (TransferConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfiguration)(nil)).Elem()
}

func (o TransferConfigurationOutput) ToTransferConfigurationOutput() TransferConfigurationOutput {
	return o
}

func (o TransferConfigurationOutput) ToTransferConfigurationOutputWithContext(ctx context.Context) TransferConfigurationOutput {
	return o
}

func (o TransferConfigurationOutput) TransferAllDetails() TransferConfigurationTransferAllDetailsPtrOutput {
	return o.ApplyT(func(v TransferConfiguration) *TransferConfigurationTransferAllDetails { return v.TransferAllDetails }).(TransferConfigurationTransferAllDetailsPtrOutput)
}

func (o TransferConfigurationOutput) TransferConfigurationType() pulumi.StringOutput {
	return o.ApplyT(func(v TransferConfiguration) string { return v.TransferConfigurationType }).(pulumi.StringOutput)
}

func (o TransferConfigurationOutput) TransferFilterDetails() TransferConfigurationTransferFilterDetailsPtrOutput {
	return o.ApplyT(func(v TransferConfiguration) *TransferConfigurationTransferFilterDetails {
		return v.TransferFilterDetails
	}).(TransferConfigurationTransferFilterDetailsPtrOutput)
}

type TransferConfigurationResponse struct {
	TransferAllDetails        *TransferConfigurationResponseTransferAllDetails    `pulumi:"transferAllDetails"`
	TransferConfigurationType string                                              `pulumi:"transferConfigurationType"`
	TransferFilterDetails     *TransferConfigurationResponseTransferFilterDetails `pulumi:"transferFilterDetails"`
}





type TransferConfigurationResponseInput interface {
	pulumi.Input

	ToTransferConfigurationResponseOutput() TransferConfigurationResponseOutput
	ToTransferConfigurationResponseOutputWithContext(context.Context) TransferConfigurationResponseOutput
}

type TransferConfigurationResponseArgs struct {
	TransferAllDetails        TransferConfigurationResponseTransferAllDetailsPtrInput    `pulumi:"transferAllDetails"`
	TransferConfigurationType pulumi.StringInput                                         `pulumi:"transferConfigurationType"`
	TransferFilterDetails     TransferConfigurationResponseTransferFilterDetailsPtrInput `pulumi:"transferFilterDetails"`
}

func (TransferConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfigurationResponse)(nil)).Elem()
}

func (i TransferConfigurationResponseArgs) ToTransferConfigurationResponseOutput() TransferConfigurationResponseOutput {
	return i.ToTransferConfigurationResponseOutputWithContext(context.Background())
}

func (i TransferConfigurationResponseArgs) ToTransferConfigurationResponseOutputWithContext(ctx context.Context) TransferConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationResponseOutput)
}

type TransferConfigurationResponseOutput struct{ *pulumi.OutputState }

func (TransferConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfigurationResponse)(nil)).Elem()
}

func (o TransferConfigurationResponseOutput) ToTransferConfigurationResponseOutput() TransferConfigurationResponseOutput {
	return o
}

func (o TransferConfigurationResponseOutput) ToTransferConfigurationResponseOutputWithContext(ctx context.Context) TransferConfigurationResponseOutput {
	return o
}

func (o TransferConfigurationResponseOutput) TransferAllDetails() TransferConfigurationResponseTransferAllDetailsPtrOutput {
	return o.ApplyT(func(v TransferConfigurationResponse) *TransferConfigurationResponseTransferAllDetails {
		return v.TransferAllDetails
	}).(TransferConfigurationResponseTransferAllDetailsPtrOutput)
}

func (o TransferConfigurationResponseOutput) TransferConfigurationType() pulumi.StringOutput {
	return o.ApplyT(func(v TransferConfigurationResponse) string { return v.TransferConfigurationType }).(pulumi.StringOutput)
}

func (o TransferConfigurationResponseOutput) TransferFilterDetails() TransferConfigurationResponseTransferFilterDetailsPtrOutput {
	return o.ApplyT(func(v TransferConfigurationResponse) *TransferConfigurationResponseTransferFilterDetails {
		return v.TransferFilterDetails
	}).(TransferConfigurationResponseTransferFilterDetailsPtrOutput)
}

type TransferConfigurationResponseTransferAllDetails struct {
	Include *TransferAllDetailsResponse `pulumi:"include"`
}





type TransferConfigurationResponseTransferAllDetailsInput interface {
	pulumi.Input

	ToTransferConfigurationResponseTransferAllDetailsOutput() TransferConfigurationResponseTransferAllDetailsOutput
	ToTransferConfigurationResponseTransferAllDetailsOutputWithContext(context.Context) TransferConfigurationResponseTransferAllDetailsOutput
}

type TransferConfigurationResponseTransferAllDetailsArgs struct {
	Include TransferAllDetailsResponsePtrInput `pulumi:"include"`
}

func (TransferConfigurationResponseTransferAllDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfigurationResponseTransferAllDetails)(nil)).Elem()
}

func (i TransferConfigurationResponseTransferAllDetailsArgs) ToTransferConfigurationResponseTransferAllDetailsOutput() TransferConfigurationResponseTransferAllDetailsOutput {
	return i.ToTransferConfigurationResponseTransferAllDetailsOutputWithContext(context.Background())
}

func (i TransferConfigurationResponseTransferAllDetailsArgs) ToTransferConfigurationResponseTransferAllDetailsOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferAllDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationResponseTransferAllDetailsOutput)
}

func (i TransferConfigurationResponseTransferAllDetailsArgs) ToTransferConfigurationResponseTransferAllDetailsPtrOutput() TransferConfigurationResponseTransferAllDetailsPtrOutput {
	return i.ToTransferConfigurationResponseTransferAllDetailsPtrOutputWithContext(context.Background())
}

func (i TransferConfigurationResponseTransferAllDetailsArgs) ToTransferConfigurationResponseTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferAllDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationResponseTransferAllDetailsOutput).ToTransferConfigurationResponseTransferAllDetailsPtrOutputWithContext(ctx)
}









type TransferConfigurationResponseTransferAllDetailsPtrInput interface {
	pulumi.Input

	ToTransferConfigurationResponseTransferAllDetailsPtrOutput() TransferConfigurationResponseTransferAllDetailsPtrOutput
	ToTransferConfigurationResponseTransferAllDetailsPtrOutputWithContext(context.Context) TransferConfigurationResponseTransferAllDetailsPtrOutput
}

type transferConfigurationResponseTransferAllDetailsPtrType TransferConfigurationResponseTransferAllDetailsArgs

func TransferConfigurationResponseTransferAllDetailsPtr(v *TransferConfigurationResponseTransferAllDetailsArgs) TransferConfigurationResponseTransferAllDetailsPtrInput {
	return (*transferConfigurationResponseTransferAllDetailsPtrType)(v)
}

func (*transferConfigurationResponseTransferAllDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferConfigurationResponseTransferAllDetails)(nil)).Elem()
}

func (i *transferConfigurationResponseTransferAllDetailsPtrType) ToTransferConfigurationResponseTransferAllDetailsPtrOutput() TransferConfigurationResponseTransferAllDetailsPtrOutput {
	return i.ToTransferConfigurationResponseTransferAllDetailsPtrOutputWithContext(context.Background())
}

func (i *transferConfigurationResponseTransferAllDetailsPtrType) ToTransferConfigurationResponseTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferAllDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationResponseTransferAllDetailsPtrOutput)
}

type TransferConfigurationResponseTransferAllDetailsOutput struct{ *pulumi.OutputState }

func (TransferConfigurationResponseTransferAllDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfigurationResponseTransferAllDetails)(nil)).Elem()
}

func (o TransferConfigurationResponseTransferAllDetailsOutput) ToTransferConfigurationResponseTransferAllDetailsOutput() TransferConfigurationResponseTransferAllDetailsOutput {
	return o
}

func (o TransferConfigurationResponseTransferAllDetailsOutput) ToTransferConfigurationResponseTransferAllDetailsOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferAllDetailsOutput {
	return o
}

func (o TransferConfigurationResponseTransferAllDetailsOutput) ToTransferConfigurationResponseTransferAllDetailsPtrOutput() TransferConfigurationResponseTransferAllDetailsPtrOutput {
	return o.ToTransferConfigurationResponseTransferAllDetailsPtrOutputWithContext(context.Background())
}

func (o TransferConfigurationResponseTransferAllDetailsOutput) ToTransferConfigurationResponseTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferAllDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransferConfigurationResponseTransferAllDetails) *TransferConfigurationResponseTransferAllDetails {
		return &v
	}).(TransferConfigurationResponseTransferAllDetailsPtrOutput)
}

func (o TransferConfigurationResponseTransferAllDetailsOutput) Include() TransferAllDetailsResponsePtrOutput {
	return o.ApplyT(func(v TransferConfigurationResponseTransferAllDetails) *TransferAllDetailsResponse { return v.Include }).(TransferAllDetailsResponsePtrOutput)
}

type TransferConfigurationResponseTransferAllDetailsPtrOutput struct{ *pulumi.OutputState }

func (TransferConfigurationResponseTransferAllDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferConfigurationResponseTransferAllDetails)(nil)).Elem()
}

func (o TransferConfigurationResponseTransferAllDetailsPtrOutput) ToTransferConfigurationResponseTransferAllDetailsPtrOutput() TransferConfigurationResponseTransferAllDetailsPtrOutput {
	return o
}

func (o TransferConfigurationResponseTransferAllDetailsPtrOutput) ToTransferConfigurationResponseTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferAllDetailsPtrOutput {
	return o
}

func (o TransferConfigurationResponseTransferAllDetailsPtrOutput) Elem() TransferConfigurationResponseTransferAllDetailsOutput {
	return o.ApplyT(func(v *TransferConfigurationResponseTransferAllDetails) TransferConfigurationResponseTransferAllDetails {
		if v != nil {
			return *v
		}
		var ret TransferConfigurationResponseTransferAllDetails
		return ret
	}).(TransferConfigurationResponseTransferAllDetailsOutput)
}

func (o TransferConfigurationResponseTransferAllDetailsPtrOutput) Include() TransferAllDetailsResponsePtrOutput {
	return o.ApplyT(func(v *TransferConfigurationResponseTransferAllDetails) *TransferAllDetailsResponse {
		if v == nil {
			return nil
		}
		return v.Include
	}).(TransferAllDetailsResponsePtrOutput)
}

type TransferConfigurationResponseTransferFilterDetails struct {
	Include *TransferFilterDetailsResponse `pulumi:"include"`
}





type TransferConfigurationResponseTransferFilterDetailsInput interface {
	pulumi.Input

	ToTransferConfigurationResponseTransferFilterDetailsOutput() TransferConfigurationResponseTransferFilterDetailsOutput
	ToTransferConfigurationResponseTransferFilterDetailsOutputWithContext(context.Context) TransferConfigurationResponseTransferFilterDetailsOutput
}

type TransferConfigurationResponseTransferFilterDetailsArgs struct {
	Include TransferFilterDetailsResponsePtrInput `pulumi:"include"`
}

func (TransferConfigurationResponseTransferFilterDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfigurationResponseTransferFilterDetails)(nil)).Elem()
}

func (i TransferConfigurationResponseTransferFilterDetailsArgs) ToTransferConfigurationResponseTransferFilterDetailsOutput() TransferConfigurationResponseTransferFilterDetailsOutput {
	return i.ToTransferConfigurationResponseTransferFilterDetailsOutputWithContext(context.Background())
}

func (i TransferConfigurationResponseTransferFilterDetailsArgs) ToTransferConfigurationResponseTransferFilterDetailsOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferFilterDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationResponseTransferFilterDetailsOutput)
}

func (i TransferConfigurationResponseTransferFilterDetailsArgs) ToTransferConfigurationResponseTransferFilterDetailsPtrOutput() TransferConfigurationResponseTransferFilterDetailsPtrOutput {
	return i.ToTransferConfigurationResponseTransferFilterDetailsPtrOutputWithContext(context.Background())
}

func (i TransferConfigurationResponseTransferFilterDetailsArgs) ToTransferConfigurationResponseTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferFilterDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationResponseTransferFilterDetailsOutput).ToTransferConfigurationResponseTransferFilterDetailsPtrOutputWithContext(ctx)
}









type TransferConfigurationResponseTransferFilterDetailsPtrInput interface {
	pulumi.Input

	ToTransferConfigurationResponseTransferFilterDetailsPtrOutput() TransferConfigurationResponseTransferFilterDetailsPtrOutput
	ToTransferConfigurationResponseTransferFilterDetailsPtrOutputWithContext(context.Context) TransferConfigurationResponseTransferFilterDetailsPtrOutput
}

type transferConfigurationResponseTransferFilterDetailsPtrType TransferConfigurationResponseTransferFilterDetailsArgs

func TransferConfigurationResponseTransferFilterDetailsPtr(v *TransferConfigurationResponseTransferFilterDetailsArgs) TransferConfigurationResponseTransferFilterDetailsPtrInput {
	return (*transferConfigurationResponseTransferFilterDetailsPtrType)(v)
}

func (*transferConfigurationResponseTransferFilterDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferConfigurationResponseTransferFilterDetails)(nil)).Elem()
}

func (i *transferConfigurationResponseTransferFilterDetailsPtrType) ToTransferConfigurationResponseTransferFilterDetailsPtrOutput() TransferConfigurationResponseTransferFilterDetailsPtrOutput {
	return i.ToTransferConfigurationResponseTransferFilterDetailsPtrOutputWithContext(context.Background())
}

func (i *transferConfigurationResponseTransferFilterDetailsPtrType) ToTransferConfigurationResponseTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferFilterDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationResponseTransferFilterDetailsPtrOutput)
}

type TransferConfigurationResponseTransferFilterDetailsOutput struct{ *pulumi.OutputState }

func (TransferConfigurationResponseTransferFilterDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfigurationResponseTransferFilterDetails)(nil)).Elem()
}

func (o TransferConfigurationResponseTransferFilterDetailsOutput) ToTransferConfigurationResponseTransferFilterDetailsOutput() TransferConfigurationResponseTransferFilterDetailsOutput {
	return o
}

func (o TransferConfigurationResponseTransferFilterDetailsOutput) ToTransferConfigurationResponseTransferFilterDetailsOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferFilterDetailsOutput {
	return o
}

func (o TransferConfigurationResponseTransferFilterDetailsOutput) ToTransferConfigurationResponseTransferFilterDetailsPtrOutput() TransferConfigurationResponseTransferFilterDetailsPtrOutput {
	return o.ToTransferConfigurationResponseTransferFilterDetailsPtrOutputWithContext(context.Background())
}

func (o TransferConfigurationResponseTransferFilterDetailsOutput) ToTransferConfigurationResponseTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferFilterDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransferConfigurationResponseTransferFilterDetails) *TransferConfigurationResponseTransferFilterDetails {
		return &v
	}).(TransferConfigurationResponseTransferFilterDetailsPtrOutput)
}

func (o TransferConfigurationResponseTransferFilterDetailsOutput) Include() TransferFilterDetailsResponsePtrOutput {
	return o.ApplyT(func(v TransferConfigurationResponseTransferFilterDetails) *TransferFilterDetailsResponse {
		return v.Include
	}).(TransferFilterDetailsResponsePtrOutput)
}

type TransferConfigurationResponseTransferFilterDetailsPtrOutput struct{ *pulumi.OutputState }

func (TransferConfigurationResponseTransferFilterDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferConfigurationResponseTransferFilterDetails)(nil)).Elem()
}

func (o TransferConfigurationResponseTransferFilterDetailsPtrOutput) ToTransferConfigurationResponseTransferFilterDetailsPtrOutput() TransferConfigurationResponseTransferFilterDetailsPtrOutput {
	return o
}

func (o TransferConfigurationResponseTransferFilterDetailsPtrOutput) ToTransferConfigurationResponseTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationResponseTransferFilterDetailsPtrOutput {
	return o
}

func (o TransferConfigurationResponseTransferFilterDetailsPtrOutput) Elem() TransferConfigurationResponseTransferFilterDetailsOutput {
	return o.ApplyT(func(v *TransferConfigurationResponseTransferFilterDetails) TransferConfigurationResponseTransferFilterDetails {
		if v != nil {
			return *v
		}
		var ret TransferConfigurationResponseTransferFilterDetails
		return ret
	}).(TransferConfigurationResponseTransferFilterDetailsOutput)
}

func (o TransferConfigurationResponseTransferFilterDetailsPtrOutput) Include() TransferFilterDetailsResponsePtrOutput {
	return o.ApplyT(func(v *TransferConfigurationResponseTransferFilterDetails) *TransferFilterDetailsResponse {
		if v == nil {
			return nil
		}
		return v.Include
	}).(TransferFilterDetailsResponsePtrOutput)
}

type TransferConfigurationTransferAllDetails struct {
	Include *TransferAllDetails `pulumi:"include"`
}





type TransferConfigurationTransferAllDetailsInput interface {
	pulumi.Input

	ToTransferConfigurationTransferAllDetailsOutput() TransferConfigurationTransferAllDetailsOutput
	ToTransferConfigurationTransferAllDetailsOutputWithContext(context.Context) TransferConfigurationTransferAllDetailsOutput
}

type TransferConfigurationTransferAllDetailsArgs struct {
	Include TransferAllDetailsPtrInput `pulumi:"include"`
}

func (TransferConfigurationTransferAllDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfigurationTransferAllDetails)(nil)).Elem()
}

func (i TransferConfigurationTransferAllDetailsArgs) ToTransferConfigurationTransferAllDetailsOutput() TransferConfigurationTransferAllDetailsOutput {
	return i.ToTransferConfigurationTransferAllDetailsOutputWithContext(context.Background())
}

func (i TransferConfigurationTransferAllDetailsArgs) ToTransferConfigurationTransferAllDetailsOutputWithContext(ctx context.Context) TransferConfigurationTransferAllDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationTransferAllDetailsOutput)
}

func (i TransferConfigurationTransferAllDetailsArgs) ToTransferConfigurationTransferAllDetailsPtrOutput() TransferConfigurationTransferAllDetailsPtrOutput {
	return i.ToTransferConfigurationTransferAllDetailsPtrOutputWithContext(context.Background())
}

func (i TransferConfigurationTransferAllDetailsArgs) ToTransferConfigurationTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationTransferAllDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationTransferAllDetailsOutput).ToTransferConfigurationTransferAllDetailsPtrOutputWithContext(ctx)
}









type TransferConfigurationTransferAllDetailsPtrInput interface {
	pulumi.Input

	ToTransferConfigurationTransferAllDetailsPtrOutput() TransferConfigurationTransferAllDetailsPtrOutput
	ToTransferConfigurationTransferAllDetailsPtrOutputWithContext(context.Context) TransferConfigurationTransferAllDetailsPtrOutput
}

type transferConfigurationTransferAllDetailsPtrType TransferConfigurationTransferAllDetailsArgs

func TransferConfigurationTransferAllDetailsPtr(v *TransferConfigurationTransferAllDetailsArgs) TransferConfigurationTransferAllDetailsPtrInput {
	return (*transferConfigurationTransferAllDetailsPtrType)(v)
}

func (*transferConfigurationTransferAllDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferConfigurationTransferAllDetails)(nil)).Elem()
}

func (i *transferConfigurationTransferAllDetailsPtrType) ToTransferConfigurationTransferAllDetailsPtrOutput() TransferConfigurationTransferAllDetailsPtrOutput {
	return i.ToTransferConfigurationTransferAllDetailsPtrOutputWithContext(context.Background())
}

func (i *transferConfigurationTransferAllDetailsPtrType) ToTransferConfigurationTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationTransferAllDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationTransferAllDetailsPtrOutput)
}

type TransferConfigurationTransferAllDetailsOutput struct{ *pulumi.OutputState }

func (TransferConfigurationTransferAllDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfigurationTransferAllDetails)(nil)).Elem()
}

func (o TransferConfigurationTransferAllDetailsOutput) ToTransferConfigurationTransferAllDetailsOutput() TransferConfigurationTransferAllDetailsOutput {
	return o
}

func (o TransferConfigurationTransferAllDetailsOutput) ToTransferConfigurationTransferAllDetailsOutputWithContext(ctx context.Context) TransferConfigurationTransferAllDetailsOutput {
	return o
}

func (o TransferConfigurationTransferAllDetailsOutput) ToTransferConfigurationTransferAllDetailsPtrOutput() TransferConfigurationTransferAllDetailsPtrOutput {
	return o.ToTransferConfigurationTransferAllDetailsPtrOutputWithContext(context.Background())
}

func (o TransferConfigurationTransferAllDetailsOutput) ToTransferConfigurationTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationTransferAllDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransferConfigurationTransferAllDetails) *TransferConfigurationTransferAllDetails {
		return &v
	}).(TransferConfigurationTransferAllDetailsPtrOutput)
}

func (o TransferConfigurationTransferAllDetailsOutput) Include() TransferAllDetailsPtrOutput {
	return o.ApplyT(func(v TransferConfigurationTransferAllDetails) *TransferAllDetails { return v.Include }).(TransferAllDetailsPtrOutput)
}

type TransferConfigurationTransferAllDetailsPtrOutput struct{ *pulumi.OutputState }

func (TransferConfigurationTransferAllDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferConfigurationTransferAllDetails)(nil)).Elem()
}

func (o TransferConfigurationTransferAllDetailsPtrOutput) ToTransferConfigurationTransferAllDetailsPtrOutput() TransferConfigurationTransferAllDetailsPtrOutput {
	return o
}

func (o TransferConfigurationTransferAllDetailsPtrOutput) ToTransferConfigurationTransferAllDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationTransferAllDetailsPtrOutput {
	return o
}

func (o TransferConfigurationTransferAllDetailsPtrOutput) Elem() TransferConfigurationTransferAllDetailsOutput {
	return o.ApplyT(func(v *TransferConfigurationTransferAllDetails) TransferConfigurationTransferAllDetails {
		if v != nil {
			return *v
		}
		var ret TransferConfigurationTransferAllDetails
		return ret
	}).(TransferConfigurationTransferAllDetailsOutput)
}

func (o TransferConfigurationTransferAllDetailsPtrOutput) Include() TransferAllDetailsPtrOutput {
	return o.ApplyT(func(v *TransferConfigurationTransferAllDetails) *TransferAllDetails {
		if v == nil {
			return nil
		}
		return v.Include
	}).(TransferAllDetailsPtrOutput)
}

type TransferConfigurationTransferFilterDetails struct {
	Include *TransferFilterDetails `pulumi:"include"`
}





type TransferConfigurationTransferFilterDetailsInput interface {
	pulumi.Input

	ToTransferConfigurationTransferFilterDetailsOutput() TransferConfigurationTransferFilterDetailsOutput
	ToTransferConfigurationTransferFilterDetailsOutputWithContext(context.Context) TransferConfigurationTransferFilterDetailsOutput
}

type TransferConfigurationTransferFilterDetailsArgs struct {
	Include TransferFilterDetailsPtrInput `pulumi:"include"`
}

func (TransferConfigurationTransferFilterDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfigurationTransferFilterDetails)(nil)).Elem()
}

func (i TransferConfigurationTransferFilterDetailsArgs) ToTransferConfigurationTransferFilterDetailsOutput() TransferConfigurationTransferFilterDetailsOutput {
	return i.ToTransferConfigurationTransferFilterDetailsOutputWithContext(context.Background())
}

func (i TransferConfigurationTransferFilterDetailsArgs) ToTransferConfigurationTransferFilterDetailsOutputWithContext(ctx context.Context) TransferConfigurationTransferFilterDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationTransferFilterDetailsOutput)
}

func (i TransferConfigurationTransferFilterDetailsArgs) ToTransferConfigurationTransferFilterDetailsPtrOutput() TransferConfigurationTransferFilterDetailsPtrOutput {
	return i.ToTransferConfigurationTransferFilterDetailsPtrOutputWithContext(context.Background())
}

func (i TransferConfigurationTransferFilterDetailsArgs) ToTransferConfigurationTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationTransferFilterDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationTransferFilterDetailsOutput).ToTransferConfigurationTransferFilterDetailsPtrOutputWithContext(ctx)
}









type TransferConfigurationTransferFilterDetailsPtrInput interface {
	pulumi.Input

	ToTransferConfigurationTransferFilterDetailsPtrOutput() TransferConfigurationTransferFilterDetailsPtrOutput
	ToTransferConfigurationTransferFilterDetailsPtrOutputWithContext(context.Context) TransferConfigurationTransferFilterDetailsPtrOutput
}

type transferConfigurationTransferFilterDetailsPtrType TransferConfigurationTransferFilterDetailsArgs

func TransferConfigurationTransferFilterDetailsPtr(v *TransferConfigurationTransferFilterDetailsArgs) TransferConfigurationTransferFilterDetailsPtrInput {
	return (*transferConfigurationTransferFilterDetailsPtrType)(v)
}

func (*transferConfigurationTransferFilterDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferConfigurationTransferFilterDetails)(nil)).Elem()
}

func (i *transferConfigurationTransferFilterDetailsPtrType) ToTransferConfigurationTransferFilterDetailsPtrOutput() TransferConfigurationTransferFilterDetailsPtrOutput {
	return i.ToTransferConfigurationTransferFilterDetailsPtrOutputWithContext(context.Background())
}

func (i *transferConfigurationTransferFilterDetailsPtrType) ToTransferConfigurationTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationTransferFilterDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigurationTransferFilterDetailsPtrOutput)
}

type TransferConfigurationTransferFilterDetailsOutput struct{ *pulumi.OutputState }

func (TransferConfigurationTransferFilterDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferConfigurationTransferFilterDetails)(nil)).Elem()
}

func (o TransferConfigurationTransferFilterDetailsOutput) ToTransferConfigurationTransferFilterDetailsOutput() TransferConfigurationTransferFilterDetailsOutput {
	return o
}

func (o TransferConfigurationTransferFilterDetailsOutput) ToTransferConfigurationTransferFilterDetailsOutputWithContext(ctx context.Context) TransferConfigurationTransferFilterDetailsOutput {
	return o
}

func (o TransferConfigurationTransferFilterDetailsOutput) ToTransferConfigurationTransferFilterDetailsPtrOutput() TransferConfigurationTransferFilterDetailsPtrOutput {
	return o.ToTransferConfigurationTransferFilterDetailsPtrOutputWithContext(context.Background())
}

func (o TransferConfigurationTransferFilterDetailsOutput) ToTransferConfigurationTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationTransferFilterDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransferConfigurationTransferFilterDetails) *TransferConfigurationTransferFilterDetails {
		return &v
	}).(TransferConfigurationTransferFilterDetailsPtrOutput)
}

func (o TransferConfigurationTransferFilterDetailsOutput) Include() TransferFilterDetailsPtrOutput {
	return o.ApplyT(func(v TransferConfigurationTransferFilterDetails) *TransferFilterDetails { return v.Include }).(TransferFilterDetailsPtrOutput)
}

type TransferConfigurationTransferFilterDetailsPtrOutput struct{ *pulumi.OutputState }

func (TransferConfigurationTransferFilterDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferConfigurationTransferFilterDetails)(nil)).Elem()
}

func (o TransferConfigurationTransferFilterDetailsPtrOutput) ToTransferConfigurationTransferFilterDetailsPtrOutput() TransferConfigurationTransferFilterDetailsPtrOutput {
	return o
}

func (o TransferConfigurationTransferFilterDetailsPtrOutput) ToTransferConfigurationTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferConfigurationTransferFilterDetailsPtrOutput {
	return o
}

func (o TransferConfigurationTransferFilterDetailsPtrOutput) Elem() TransferConfigurationTransferFilterDetailsOutput {
	return o.ApplyT(func(v *TransferConfigurationTransferFilterDetails) TransferConfigurationTransferFilterDetails {
		if v != nil {
			return *v
		}
		var ret TransferConfigurationTransferFilterDetails
		return ret
	}).(TransferConfigurationTransferFilterDetailsOutput)
}

func (o TransferConfigurationTransferFilterDetailsPtrOutput) Include() TransferFilterDetailsPtrOutput {
	return o.ApplyT(func(v *TransferConfigurationTransferFilterDetails) *TransferFilterDetails {
		if v == nil {
			return nil
		}
		return v.Include
	}).(TransferFilterDetailsPtrOutput)
}

type TransferFilterDetails struct {
	AzureFileFilterDetails *AzureFileFilterDetails `pulumi:"azureFileFilterDetails"`
	BlobFilterDetails      *BlobFilterDetails      `pulumi:"blobFilterDetails"`
	DataAccountType        string                  `pulumi:"dataAccountType"`
	FilterFileDetails      []FilterFileDetails     `pulumi:"filterFileDetails"`
}





type TransferFilterDetailsInput interface {
	pulumi.Input

	ToTransferFilterDetailsOutput() TransferFilterDetailsOutput
	ToTransferFilterDetailsOutputWithContext(context.Context) TransferFilterDetailsOutput
}

type TransferFilterDetailsArgs struct {
	AzureFileFilterDetails AzureFileFilterDetailsPtrInput `pulumi:"azureFileFilterDetails"`
	BlobFilterDetails      BlobFilterDetailsPtrInput      `pulumi:"blobFilterDetails"`
	DataAccountType        pulumi.StringInput             `pulumi:"dataAccountType"`
	FilterFileDetails      FilterFileDetailsArrayInput    `pulumi:"filterFileDetails"`
}

func (TransferFilterDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferFilterDetails)(nil)).Elem()
}

func (i TransferFilterDetailsArgs) ToTransferFilterDetailsOutput() TransferFilterDetailsOutput {
	return i.ToTransferFilterDetailsOutputWithContext(context.Background())
}

func (i TransferFilterDetailsArgs) ToTransferFilterDetailsOutputWithContext(ctx context.Context) TransferFilterDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferFilterDetailsOutput)
}

func (i TransferFilterDetailsArgs) ToTransferFilterDetailsPtrOutput() TransferFilterDetailsPtrOutput {
	return i.ToTransferFilterDetailsPtrOutputWithContext(context.Background())
}

func (i TransferFilterDetailsArgs) ToTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferFilterDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferFilterDetailsOutput).ToTransferFilterDetailsPtrOutputWithContext(ctx)
}









type TransferFilterDetailsPtrInput interface {
	pulumi.Input

	ToTransferFilterDetailsPtrOutput() TransferFilterDetailsPtrOutput
	ToTransferFilterDetailsPtrOutputWithContext(context.Context) TransferFilterDetailsPtrOutput
}

type transferFilterDetailsPtrType TransferFilterDetailsArgs

func TransferFilterDetailsPtr(v *TransferFilterDetailsArgs) TransferFilterDetailsPtrInput {
	return (*transferFilterDetailsPtrType)(v)
}

func (*transferFilterDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferFilterDetails)(nil)).Elem()
}

func (i *transferFilterDetailsPtrType) ToTransferFilterDetailsPtrOutput() TransferFilterDetailsPtrOutput {
	return i.ToTransferFilterDetailsPtrOutputWithContext(context.Background())
}

func (i *transferFilterDetailsPtrType) ToTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferFilterDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferFilterDetailsPtrOutput)
}

type TransferFilterDetailsOutput struct{ *pulumi.OutputState }

func (TransferFilterDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferFilterDetails)(nil)).Elem()
}

func (o TransferFilterDetailsOutput) ToTransferFilterDetailsOutput() TransferFilterDetailsOutput {
	return o
}

func (o TransferFilterDetailsOutput) ToTransferFilterDetailsOutputWithContext(ctx context.Context) TransferFilterDetailsOutput {
	return o
}

func (o TransferFilterDetailsOutput) ToTransferFilterDetailsPtrOutput() TransferFilterDetailsPtrOutput {
	return o.ToTransferFilterDetailsPtrOutputWithContext(context.Background())
}

func (o TransferFilterDetailsOutput) ToTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferFilterDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransferFilterDetails) *TransferFilterDetails {
		return &v
	}).(TransferFilterDetailsPtrOutput)
}

func (o TransferFilterDetailsOutput) AzureFileFilterDetails() AzureFileFilterDetailsPtrOutput {
	return o.ApplyT(func(v TransferFilterDetails) *AzureFileFilterDetails { return v.AzureFileFilterDetails }).(AzureFileFilterDetailsPtrOutput)
}

func (o TransferFilterDetailsOutput) BlobFilterDetails() BlobFilterDetailsPtrOutput {
	return o.ApplyT(func(v TransferFilterDetails) *BlobFilterDetails { return v.BlobFilterDetails }).(BlobFilterDetailsPtrOutput)
}

func (o TransferFilterDetailsOutput) DataAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v TransferFilterDetails) string { return v.DataAccountType }).(pulumi.StringOutput)
}

func (o TransferFilterDetailsOutput) FilterFileDetails() FilterFileDetailsArrayOutput {
	return o.ApplyT(func(v TransferFilterDetails) []FilterFileDetails { return v.FilterFileDetails }).(FilterFileDetailsArrayOutput)
}

type TransferFilterDetailsPtrOutput struct{ *pulumi.OutputState }

func (TransferFilterDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferFilterDetails)(nil)).Elem()
}

func (o TransferFilterDetailsPtrOutput) ToTransferFilterDetailsPtrOutput() TransferFilterDetailsPtrOutput {
	return o
}

func (o TransferFilterDetailsPtrOutput) ToTransferFilterDetailsPtrOutputWithContext(ctx context.Context) TransferFilterDetailsPtrOutput {
	return o
}

func (o TransferFilterDetailsPtrOutput) Elem() TransferFilterDetailsOutput {
	return o.ApplyT(func(v *TransferFilterDetails) TransferFilterDetails {
		if v != nil {
			return *v
		}
		var ret TransferFilterDetails
		return ret
	}).(TransferFilterDetailsOutput)
}

func (o TransferFilterDetailsPtrOutput) AzureFileFilterDetails() AzureFileFilterDetailsPtrOutput {
	return o.ApplyT(func(v *TransferFilterDetails) *AzureFileFilterDetails {
		if v == nil {
			return nil
		}
		return v.AzureFileFilterDetails
	}).(AzureFileFilterDetailsPtrOutput)
}

func (o TransferFilterDetailsPtrOutput) BlobFilterDetails() BlobFilterDetailsPtrOutput {
	return o.ApplyT(func(v *TransferFilterDetails) *BlobFilterDetails {
		if v == nil {
			return nil
		}
		return v.BlobFilterDetails
	}).(BlobFilterDetailsPtrOutput)
}

func (o TransferFilterDetailsPtrOutput) DataAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransferFilterDetails) *string {
		if v == nil {
			return nil
		}
		return &v.DataAccountType
	}).(pulumi.StringPtrOutput)
}

func (o TransferFilterDetailsPtrOutput) FilterFileDetails() FilterFileDetailsArrayOutput {
	return o.ApplyT(func(v *TransferFilterDetails) []FilterFileDetails {
		if v == nil {
			return nil
		}
		return v.FilterFileDetails
	}).(FilterFileDetailsArrayOutput)
}

type TransferFilterDetailsResponse struct {
	AzureFileFilterDetails *AzureFileFilterDetailsResponse `pulumi:"azureFileFilterDetails"`
	BlobFilterDetails      *BlobFilterDetailsResponse      `pulumi:"blobFilterDetails"`
	DataAccountType        string                          `pulumi:"dataAccountType"`
	FilterFileDetails      []FilterFileDetailsResponse     `pulumi:"filterFileDetails"`
}





type TransferFilterDetailsResponseInput interface {
	pulumi.Input

	ToTransferFilterDetailsResponseOutput() TransferFilterDetailsResponseOutput
	ToTransferFilterDetailsResponseOutputWithContext(context.Context) TransferFilterDetailsResponseOutput
}

type TransferFilterDetailsResponseArgs struct {
	AzureFileFilterDetails AzureFileFilterDetailsResponsePtrInput `pulumi:"azureFileFilterDetails"`
	BlobFilterDetails      BlobFilterDetailsResponsePtrInput      `pulumi:"blobFilterDetails"`
	DataAccountType        pulumi.StringInput                     `pulumi:"dataAccountType"`
	FilterFileDetails      FilterFileDetailsResponseArrayInput    `pulumi:"filterFileDetails"`
}

func (TransferFilterDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferFilterDetailsResponse)(nil)).Elem()
}

func (i TransferFilterDetailsResponseArgs) ToTransferFilterDetailsResponseOutput() TransferFilterDetailsResponseOutput {
	return i.ToTransferFilterDetailsResponseOutputWithContext(context.Background())
}

func (i TransferFilterDetailsResponseArgs) ToTransferFilterDetailsResponseOutputWithContext(ctx context.Context) TransferFilterDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferFilterDetailsResponseOutput)
}

func (i TransferFilterDetailsResponseArgs) ToTransferFilterDetailsResponsePtrOutput() TransferFilterDetailsResponsePtrOutput {
	return i.ToTransferFilterDetailsResponsePtrOutputWithContext(context.Background())
}

func (i TransferFilterDetailsResponseArgs) ToTransferFilterDetailsResponsePtrOutputWithContext(ctx context.Context) TransferFilterDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferFilterDetailsResponseOutput).ToTransferFilterDetailsResponsePtrOutputWithContext(ctx)
}









type TransferFilterDetailsResponsePtrInput interface {
	pulumi.Input

	ToTransferFilterDetailsResponsePtrOutput() TransferFilterDetailsResponsePtrOutput
	ToTransferFilterDetailsResponsePtrOutputWithContext(context.Context) TransferFilterDetailsResponsePtrOutput
}

type transferFilterDetailsResponsePtrType TransferFilterDetailsResponseArgs

func TransferFilterDetailsResponsePtr(v *TransferFilterDetailsResponseArgs) TransferFilterDetailsResponsePtrInput {
	return (*transferFilterDetailsResponsePtrType)(v)
}

func (*transferFilterDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferFilterDetailsResponse)(nil)).Elem()
}

func (i *transferFilterDetailsResponsePtrType) ToTransferFilterDetailsResponsePtrOutput() TransferFilterDetailsResponsePtrOutput {
	return i.ToTransferFilterDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *transferFilterDetailsResponsePtrType) ToTransferFilterDetailsResponsePtrOutputWithContext(ctx context.Context) TransferFilterDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferFilterDetailsResponsePtrOutput)
}

type TransferFilterDetailsResponseOutput struct{ *pulumi.OutputState }

func (TransferFilterDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferFilterDetailsResponse)(nil)).Elem()
}

func (o TransferFilterDetailsResponseOutput) ToTransferFilterDetailsResponseOutput() TransferFilterDetailsResponseOutput {
	return o
}

func (o TransferFilterDetailsResponseOutput) ToTransferFilterDetailsResponseOutputWithContext(ctx context.Context) TransferFilterDetailsResponseOutput {
	return o
}

func (o TransferFilterDetailsResponseOutput) ToTransferFilterDetailsResponsePtrOutput() TransferFilterDetailsResponsePtrOutput {
	return o.ToTransferFilterDetailsResponsePtrOutputWithContext(context.Background())
}

func (o TransferFilterDetailsResponseOutput) ToTransferFilterDetailsResponsePtrOutputWithContext(ctx context.Context) TransferFilterDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransferFilterDetailsResponse) *TransferFilterDetailsResponse {
		return &v
	}).(TransferFilterDetailsResponsePtrOutput)
}

func (o TransferFilterDetailsResponseOutput) AzureFileFilterDetails() AzureFileFilterDetailsResponsePtrOutput {
	return o.ApplyT(func(v TransferFilterDetailsResponse) *AzureFileFilterDetailsResponse { return v.AzureFileFilterDetails }).(AzureFileFilterDetailsResponsePtrOutput)
}

func (o TransferFilterDetailsResponseOutput) BlobFilterDetails() BlobFilterDetailsResponsePtrOutput {
	return o.ApplyT(func(v TransferFilterDetailsResponse) *BlobFilterDetailsResponse { return v.BlobFilterDetails }).(BlobFilterDetailsResponsePtrOutput)
}

func (o TransferFilterDetailsResponseOutput) DataAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v TransferFilterDetailsResponse) string { return v.DataAccountType }).(pulumi.StringOutput)
}

func (o TransferFilterDetailsResponseOutput) FilterFileDetails() FilterFileDetailsResponseArrayOutput {
	return o.ApplyT(func(v TransferFilterDetailsResponse) []FilterFileDetailsResponse { return v.FilterFileDetails }).(FilterFileDetailsResponseArrayOutput)
}

type TransferFilterDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (TransferFilterDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferFilterDetailsResponse)(nil)).Elem()
}

func (o TransferFilterDetailsResponsePtrOutput) ToTransferFilterDetailsResponsePtrOutput() TransferFilterDetailsResponsePtrOutput {
	return o
}

func (o TransferFilterDetailsResponsePtrOutput) ToTransferFilterDetailsResponsePtrOutputWithContext(ctx context.Context) TransferFilterDetailsResponsePtrOutput {
	return o
}

func (o TransferFilterDetailsResponsePtrOutput) Elem() TransferFilterDetailsResponseOutput {
	return o.ApplyT(func(v *TransferFilterDetailsResponse) TransferFilterDetailsResponse {
		if v != nil {
			return *v
		}
		var ret TransferFilterDetailsResponse
		return ret
	}).(TransferFilterDetailsResponseOutput)
}

func (o TransferFilterDetailsResponsePtrOutput) AzureFileFilterDetails() AzureFileFilterDetailsResponsePtrOutput {
	return o.ApplyT(func(v *TransferFilterDetailsResponse) *AzureFileFilterDetailsResponse {
		if v == nil {
			return nil
		}
		return v.AzureFileFilterDetails
	}).(AzureFileFilterDetailsResponsePtrOutput)
}

func (o TransferFilterDetailsResponsePtrOutput) BlobFilterDetails() BlobFilterDetailsResponsePtrOutput {
	return o.ApplyT(func(v *TransferFilterDetailsResponse) *BlobFilterDetailsResponse {
		if v == nil {
			return nil
		}
		return v.BlobFilterDetails
	}).(BlobFilterDetailsResponsePtrOutput)
}

func (o TransferFilterDetailsResponsePtrOutput) DataAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransferFilterDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataAccountType
	}).(pulumi.StringPtrOutput)
}

func (o TransferFilterDetailsResponsePtrOutput) FilterFileDetails() FilterFileDetailsResponseArrayOutput {
	return o.ApplyT(func(v *TransferFilterDetailsResponse) []FilterFileDetailsResponse {
		if v == nil {
			return nil
		}
		return v.FilterFileDetails
	}).(FilterFileDetailsResponseArrayOutput)
}

type TransportPreferences struct {
	PreferredShipmentType string `pulumi:"preferredShipmentType"`
}





type TransportPreferencesInput interface {
	pulumi.Input

	ToTransportPreferencesOutput() TransportPreferencesOutput
	ToTransportPreferencesOutputWithContext(context.Context) TransportPreferencesOutput
}

type TransportPreferencesArgs struct {
	PreferredShipmentType pulumi.StringInput `pulumi:"preferredShipmentType"`
}

func (TransportPreferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportPreferences)(nil)).Elem()
}

func (i TransportPreferencesArgs) ToTransportPreferencesOutput() TransportPreferencesOutput {
	return i.ToTransportPreferencesOutputWithContext(context.Background())
}

func (i TransportPreferencesArgs) ToTransportPreferencesOutputWithContext(ctx context.Context) TransportPreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesOutput)
}

func (i TransportPreferencesArgs) ToTransportPreferencesPtrOutput() TransportPreferencesPtrOutput {
	return i.ToTransportPreferencesPtrOutputWithContext(context.Background())
}

func (i TransportPreferencesArgs) ToTransportPreferencesPtrOutputWithContext(ctx context.Context) TransportPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesOutput).ToTransportPreferencesPtrOutputWithContext(ctx)
}









type TransportPreferencesPtrInput interface {
	pulumi.Input

	ToTransportPreferencesPtrOutput() TransportPreferencesPtrOutput
	ToTransportPreferencesPtrOutputWithContext(context.Context) TransportPreferencesPtrOutput
}

type transportPreferencesPtrType TransportPreferencesArgs

func TransportPreferencesPtr(v *TransportPreferencesArgs) TransportPreferencesPtrInput {
	return (*transportPreferencesPtrType)(v)
}

func (*transportPreferencesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransportPreferences)(nil)).Elem()
}

func (i *transportPreferencesPtrType) ToTransportPreferencesPtrOutput() TransportPreferencesPtrOutput {
	return i.ToTransportPreferencesPtrOutputWithContext(context.Background())
}

func (i *transportPreferencesPtrType) ToTransportPreferencesPtrOutputWithContext(ctx context.Context) TransportPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesPtrOutput)
}

type TransportPreferencesOutput struct{ *pulumi.OutputState }

func (TransportPreferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportPreferences)(nil)).Elem()
}

func (o TransportPreferencesOutput) ToTransportPreferencesOutput() TransportPreferencesOutput {
	return o
}

func (o TransportPreferencesOutput) ToTransportPreferencesOutputWithContext(ctx context.Context) TransportPreferencesOutput {
	return o
}

func (o TransportPreferencesOutput) ToTransportPreferencesPtrOutput() TransportPreferencesPtrOutput {
	return o.ToTransportPreferencesPtrOutputWithContext(context.Background())
}

func (o TransportPreferencesOutput) ToTransportPreferencesPtrOutputWithContext(ctx context.Context) TransportPreferencesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransportPreferences) *TransportPreferences {
		return &v
	}).(TransportPreferencesPtrOutput)
}

func (o TransportPreferencesOutput) PreferredShipmentType() pulumi.StringOutput {
	return o.ApplyT(func(v TransportPreferences) string { return v.PreferredShipmentType }).(pulumi.StringOutput)
}

type TransportPreferencesPtrOutput struct{ *pulumi.OutputState }

func (TransportPreferencesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransportPreferences)(nil)).Elem()
}

func (o TransportPreferencesPtrOutput) ToTransportPreferencesPtrOutput() TransportPreferencesPtrOutput {
	return o
}

func (o TransportPreferencesPtrOutput) ToTransportPreferencesPtrOutputWithContext(ctx context.Context) TransportPreferencesPtrOutput {
	return o
}

func (o TransportPreferencesPtrOutput) Elem() TransportPreferencesOutput {
	return o.ApplyT(func(v *TransportPreferences) TransportPreferences {
		if v != nil {
			return *v
		}
		var ret TransportPreferences
		return ret
	}).(TransportPreferencesOutput)
}

func (o TransportPreferencesPtrOutput) PreferredShipmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransportPreferences) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredShipmentType
	}).(pulumi.StringPtrOutput)
}

type TransportPreferencesResponse struct {
	PreferredShipmentType string `pulumi:"preferredShipmentType"`
}





type TransportPreferencesResponseInput interface {
	pulumi.Input

	ToTransportPreferencesResponseOutput() TransportPreferencesResponseOutput
	ToTransportPreferencesResponseOutputWithContext(context.Context) TransportPreferencesResponseOutput
}

type TransportPreferencesResponseArgs struct {
	PreferredShipmentType pulumi.StringInput `pulumi:"preferredShipmentType"`
}

func (TransportPreferencesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportPreferencesResponse)(nil)).Elem()
}

func (i TransportPreferencesResponseArgs) ToTransportPreferencesResponseOutput() TransportPreferencesResponseOutput {
	return i.ToTransportPreferencesResponseOutputWithContext(context.Background())
}

func (i TransportPreferencesResponseArgs) ToTransportPreferencesResponseOutputWithContext(ctx context.Context) TransportPreferencesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesResponseOutput)
}

func (i TransportPreferencesResponseArgs) ToTransportPreferencesResponsePtrOutput() TransportPreferencesResponsePtrOutput {
	return i.ToTransportPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i TransportPreferencesResponseArgs) ToTransportPreferencesResponsePtrOutputWithContext(ctx context.Context) TransportPreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesResponseOutput).ToTransportPreferencesResponsePtrOutputWithContext(ctx)
}









type TransportPreferencesResponsePtrInput interface {
	pulumi.Input

	ToTransportPreferencesResponsePtrOutput() TransportPreferencesResponsePtrOutput
	ToTransportPreferencesResponsePtrOutputWithContext(context.Context) TransportPreferencesResponsePtrOutput
}

type transportPreferencesResponsePtrType TransportPreferencesResponseArgs

func TransportPreferencesResponsePtr(v *TransportPreferencesResponseArgs) TransportPreferencesResponsePtrInput {
	return (*transportPreferencesResponsePtrType)(v)
}

func (*transportPreferencesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransportPreferencesResponse)(nil)).Elem()
}

func (i *transportPreferencesResponsePtrType) ToTransportPreferencesResponsePtrOutput() TransportPreferencesResponsePtrOutput {
	return i.ToTransportPreferencesResponsePtrOutputWithContext(context.Background())
}

func (i *transportPreferencesResponsePtrType) ToTransportPreferencesResponsePtrOutputWithContext(ctx context.Context) TransportPreferencesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportPreferencesResponsePtrOutput)
}

type TransportPreferencesResponseOutput struct{ *pulumi.OutputState }

func (TransportPreferencesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportPreferencesResponse)(nil)).Elem()
}

func (o TransportPreferencesResponseOutput) ToTransportPreferencesResponseOutput() TransportPreferencesResponseOutput {
	return o
}

func (o TransportPreferencesResponseOutput) ToTransportPreferencesResponseOutputWithContext(ctx context.Context) TransportPreferencesResponseOutput {
	return o
}

func (o TransportPreferencesResponseOutput) ToTransportPreferencesResponsePtrOutput() TransportPreferencesResponsePtrOutput {
	return o.ToTransportPreferencesResponsePtrOutputWithContext(context.Background())
}

func (o TransportPreferencesResponseOutput) ToTransportPreferencesResponsePtrOutputWithContext(ctx context.Context) TransportPreferencesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransportPreferencesResponse) *TransportPreferencesResponse {
		return &v
	}).(TransportPreferencesResponsePtrOutput)
}

func (o TransportPreferencesResponseOutput) PreferredShipmentType() pulumi.StringOutput {
	return o.ApplyT(func(v TransportPreferencesResponse) string { return v.PreferredShipmentType }).(pulumi.StringOutput)
}

type TransportPreferencesResponsePtrOutput struct{ *pulumi.OutputState }

func (TransportPreferencesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransportPreferencesResponse)(nil)).Elem()
}

func (o TransportPreferencesResponsePtrOutput) ToTransportPreferencesResponsePtrOutput() TransportPreferencesResponsePtrOutput {
	return o
}

func (o TransportPreferencesResponsePtrOutput) ToTransportPreferencesResponsePtrOutputWithContext(ctx context.Context) TransportPreferencesResponsePtrOutput {
	return o
}

func (o TransportPreferencesResponsePtrOutput) Elem() TransportPreferencesResponseOutput {
	return o.ApplyT(func(v *TransportPreferencesResponse) TransportPreferencesResponse {
		if v != nil {
			return *v
		}
		var ret TransportPreferencesResponse
		return ret
	}).(TransportPreferencesResponseOutput)
}

func (o TransportPreferencesResponsePtrOutput) PreferredShipmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransportPreferencesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredShipmentType
	}).(pulumi.StringPtrOutput)
}

type UnencryptedCredentialsResponse struct {
	JobName    string      `pulumi:"jobName"`
	JobSecrets interface{} `pulumi:"jobSecrets"`
}





type UnencryptedCredentialsResponseInput interface {
	pulumi.Input

	ToUnencryptedCredentialsResponseOutput() UnencryptedCredentialsResponseOutput
	ToUnencryptedCredentialsResponseOutputWithContext(context.Context) UnencryptedCredentialsResponseOutput
}

type UnencryptedCredentialsResponseArgs struct {
	JobName    pulumi.StringInput `pulumi:"jobName"`
	JobSecrets pulumi.Input       `pulumi:"jobSecrets"`
}

func (UnencryptedCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UnencryptedCredentialsResponse)(nil)).Elem()
}

func (i UnencryptedCredentialsResponseArgs) ToUnencryptedCredentialsResponseOutput() UnencryptedCredentialsResponseOutput {
	return i.ToUnencryptedCredentialsResponseOutputWithContext(context.Background())
}

func (i UnencryptedCredentialsResponseArgs) ToUnencryptedCredentialsResponseOutputWithContext(ctx context.Context) UnencryptedCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnencryptedCredentialsResponseOutput)
}





type UnencryptedCredentialsResponseArrayInput interface {
	pulumi.Input

	ToUnencryptedCredentialsResponseArrayOutput() UnencryptedCredentialsResponseArrayOutput
	ToUnencryptedCredentialsResponseArrayOutputWithContext(context.Context) UnencryptedCredentialsResponseArrayOutput
}

type UnencryptedCredentialsResponseArray []UnencryptedCredentialsResponseInput

func (UnencryptedCredentialsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UnencryptedCredentialsResponse)(nil)).Elem()
}

func (i UnencryptedCredentialsResponseArray) ToUnencryptedCredentialsResponseArrayOutput() UnencryptedCredentialsResponseArrayOutput {
	return i.ToUnencryptedCredentialsResponseArrayOutputWithContext(context.Background())
}

func (i UnencryptedCredentialsResponseArray) ToUnencryptedCredentialsResponseArrayOutputWithContext(ctx context.Context) UnencryptedCredentialsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnencryptedCredentialsResponseArrayOutput)
}

type UnencryptedCredentialsResponseOutput struct{ *pulumi.OutputState }

func (UnencryptedCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnencryptedCredentialsResponse)(nil)).Elem()
}

func (o UnencryptedCredentialsResponseOutput) ToUnencryptedCredentialsResponseOutput() UnencryptedCredentialsResponseOutput {
	return o
}

func (o UnencryptedCredentialsResponseOutput) ToUnencryptedCredentialsResponseOutputWithContext(ctx context.Context) UnencryptedCredentialsResponseOutput {
	return o
}

func (o UnencryptedCredentialsResponseOutput) JobName() pulumi.StringOutput {
	return o.ApplyT(func(v UnencryptedCredentialsResponse) string { return v.JobName }).(pulumi.StringOutput)
}

func (o UnencryptedCredentialsResponseOutput) JobSecrets() pulumi.AnyOutput {
	return o.ApplyT(func(v UnencryptedCredentialsResponse) interface{} { return v.JobSecrets }).(pulumi.AnyOutput)
}

type UnencryptedCredentialsResponseArrayOutput struct{ *pulumi.OutputState }

func (UnencryptedCredentialsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UnencryptedCredentialsResponse)(nil)).Elem()
}

func (o UnencryptedCredentialsResponseArrayOutput) ToUnencryptedCredentialsResponseArrayOutput() UnencryptedCredentialsResponseArrayOutput {
	return o
}

func (o UnencryptedCredentialsResponseArrayOutput) ToUnencryptedCredentialsResponseArrayOutputWithContext(ctx context.Context) UnencryptedCredentialsResponseArrayOutput {
	return o
}

func (o UnencryptedCredentialsResponseArrayOutput) Index(i pulumi.IntInput) UnencryptedCredentialsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UnencryptedCredentialsResponse {
		return vs[0].([]UnencryptedCredentialsResponse)[vs[1].(int)]
	}).(UnencryptedCredentialsResponseOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type UserAssignedIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput
	ToUserAssignedIdentityResponseOutputWithContext(context.Context) UserAssignedIdentityResponseOutput
}

type UserAssignedIdentityResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
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

type UserAssignedProperties struct {
	ResourceId *string `pulumi:"resourceId"`
}





type UserAssignedPropertiesInput interface {
	pulumi.Input

	ToUserAssignedPropertiesOutput() UserAssignedPropertiesOutput
	ToUserAssignedPropertiesOutputWithContext(context.Context) UserAssignedPropertiesOutput
}

type UserAssignedPropertiesArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (UserAssignedPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedProperties)(nil)).Elem()
}

func (i UserAssignedPropertiesArgs) ToUserAssignedPropertiesOutput() UserAssignedPropertiesOutput {
	return i.ToUserAssignedPropertiesOutputWithContext(context.Background())
}

func (i UserAssignedPropertiesArgs) ToUserAssignedPropertiesOutputWithContext(ctx context.Context) UserAssignedPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedPropertiesOutput)
}

func (i UserAssignedPropertiesArgs) ToUserAssignedPropertiesPtrOutput() UserAssignedPropertiesPtrOutput {
	return i.ToUserAssignedPropertiesPtrOutputWithContext(context.Background())
}

func (i UserAssignedPropertiesArgs) ToUserAssignedPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedPropertiesOutput).ToUserAssignedPropertiesPtrOutputWithContext(ctx)
}









type UserAssignedPropertiesPtrInput interface {
	pulumi.Input

	ToUserAssignedPropertiesPtrOutput() UserAssignedPropertiesPtrOutput
	ToUserAssignedPropertiesPtrOutputWithContext(context.Context) UserAssignedPropertiesPtrOutput
}

type userAssignedPropertiesPtrType UserAssignedPropertiesArgs

func UserAssignedPropertiesPtr(v *UserAssignedPropertiesArgs) UserAssignedPropertiesPtrInput {
	return (*userAssignedPropertiesPtrType)(v)
}

func (*userAssignedPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedProperties)(nil)).Elem()
}

func (i *userAssignedPropertiesPtrType) ToUserAssignedPropertiesPtrOutput() UserAssignedPropertiesPtrOutput {
	return i.ToUserAssignedPropertiesPtrOutputWithContext(context.Background())
}

func (i *userAssignedPropertiesPtrType) ToUserAssignedPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedPropertiesPtrOutput)
}

type UserAssignedPropertiesOutput struct{ *pulumi.OutputState }

func (UserAssignedPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedProperties)(nil)).Elem()
}

func (o UserAssignedPropertiesOutput) ToUserAssignedPropertiesOutput() UserAssignedPropertiesOutput {
	return o
}

func (o UserAssignedPropertiesOutput) ToUserAssignedPropertiesOutputWithContext(ctx context.Context) UserAssignedPropertiesOutput {
	return o
}

func (o UserAssignedPropertiesOutput) ToUserAssignedPropertiesPtrOutput() UserAssignedPropertiesPtrOutput {
	return o.ToUserAssignedPropertiesPtrOutputWithContext(context.Background())
}

func (o UserAssignedPropertiesOutput) ToUserAssignedPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserAssignedProperties) *UserAssignedProperties {
		return &v
	}).(UserAssignedPropertiesPtrOutput)
}

func (o UserAssignedPropertiesOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedProperties) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type UserAssignedPropertiesPtrOutput struct{ *pulumi.OutputState }

func (UserAssignedPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedProperties)(nil)).Elem()
}

func (o UserAssignedPropertiesPtrOutput) ToUserAssignedPropertiesPtrOutput() UserAssignedPropertiesPtrOutput {
	return o
}

func (o UserAssignedPropertiesPtrOutput) ToUserAssignedPropertiesPtrOutputWithContext(ctx context.Context) UserAssignedPropertiesPtrOutput {
	return o
}

func (o UserAssignedPropertiesPtrOutput) Elem() UserAssignedPropertiesOutput {
	return o.ApplyT(func(v *UserAssignedProperties) UserAssignedProperties {
		if v != nil {
			return *v
		}
		var ret UserAssignedProperties
		return ret
	}).(UserAssignedPropertiesOutput)
}

func (o UserAssignedPropertiesPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAssignedProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type UserAssignedPropertiesResponse struct {
	ResourceId *string `pulumi:"resourceId"`
}





type UserAssignedPropertiesResponseInput interface {
	pulumi.Input

	ToUserAssignedPropertiesResponseOutput() UserAssignedPropertiesResponseOutput
	ToUserAssignedPropertiesResponseOutputWithContext(context.Context) UserAssignedPropertiesResponseOutput
}

type UserAssignedPropertiesResponseArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (UserAssignedPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedPropertiesResponse)(nil)).Elem()
}

func (i UserAssignedPropertiesResponseArgs) ToUserAssignedPropertiesResponseOutput() UserAssignedPropertiesResponseOutput {
	return i.ToUserAssignedPropertiesResponseOutputWithContext(context.Background())
}

func (i UserAssignedPropertiesResponseArgs) ToUserAssignedPropertiesResponseOutputWithContext(ctx context.Context) UserAssignedPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedPropertiesResponseOutput)
}

func (i UserAssignedPropertiesResponseArgs) ToUserAssignedPropertiesResponsePtrOutput() UserAssignedPropertiesResponsePtrOutput {
	return i.ToUserAssignedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i UserAssignedPropertiesResponseArgs) ToUserAssignedPropertiesResponsePtrOutputWithContext(ctx context.Context) UserAssignedPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedPropertiesResponseOutput).ToUserAssignedPropertiesResponsePtrOutputWithContext(ctx)
}









type UserAssignedPropertiesResponsePtrInput interface {
	pulumi.Input

	ToUserAssignedPropertiesResponsePtrOutput() UserAssignedPropertiesResponsePtrOutput
	ToUserAssignedPropertiesResponsePtrOutputWithContext(context.Context) UserAssignedPropertiesResponsePtrOutput
}

type userAssignedPropertiesResponsePtrType UserAssignedPropertiesResponseArgs

func UserAssignedPropertiesResponsePtr(v *UserAssignedPropertiesResponseArgs) UserAssignedPropertiesResponsePtrInput {
	return (*userAssignedPropertiesResponsePtrType)(v)
}

func (*userAssignedPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedPropertiesResponse)(nil)).Elem()
}

func (i *userAssignedPropertiesResponsePtrType) ToUserAssignedPropertiesResponsePtrOutput() UserAssignedPropertiesResponsePtrOutput {
	return i.ToUserAssignedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *userAssignedPropertiesResponsePtrType) ToUserAssignedPropertiesResponsePtrOutputWithContext(ctx context.Context) UserAssignedPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedPropertiesResponsePtrOutput)
}

type UserAssignedPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedPropertiesResponse)(nil)).Elem()
}

func (o UserAssignedPropertiesResponseOutput) ToUserAssignedPropertiesResponseOutput() UserAssignedPropertiesResponseOutput {
	return o
}

func (o UserAssignedPropertiesResponseOutput) ToUserAssignedPropertiesResponseOutputWithContext(ctx context.Context) UserAssignedPropertiesResponseOutput {
	return o
}

func (o UserAssignedPropertiesResponseOutput) ToUserAssignedPropertiesResponsePtrOutput() UserAssignedPropertiesResponsePtrOutput {
	return o.ToUserAssignedPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o UserAssignedPropertiesResponseOutput) ToUserAssignedPropertiesResponsePtrOutputWithContext(ctx context.Context) UserAssignedPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserAssignedPropertiesResponse) *UserAssignedPropertiesResponse {
		return &v
	}).(UserAssignedPropertiesResponsePtrOutput)
}

func (o UserAssignedPropertiesResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedPropertiesResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type UserAssignedPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (UserAssignedPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAssignedPropertiesResponse)(nil)).Elem()
}

func (o UserAssignedPropertiesResponsePtrOutput) ToUserAssignedPropertiesResponsePtrOutput() UserAssignedPropertiesResponsePtrOutput {
	return o
}

func (o UserAssignedPropertiesResponsePtrOutput) ToUserAssignedPropertiesResponsePtrOutputWithContext(ctx context.Context) UserAssignedPropertiesResponsePtrOutput {
	return o
}

func (o UserAssignedPropertiesResponsePtrOutput) Elem() UserAssignedPropertiesResponseOutput {
	return o.ApplyT(func(v *UserAssignedPropertiesResponse) UserAssignedPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret UserAssignedPropertiesResponse
		return ret
	}).(UserAssignedPropertiesResponseOutput)
}

func (o UserAssignedPropertiesResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAssignedPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountCredentialDetailsResponseOutput{})
	pulumi.RegisterOutputType(AccountCredentialDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(AdditionalErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(AdditionalErrorInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplianceNetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ApplianceNetworkConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureFileFilterDetailsOutput{})
	pulumi.RegisterOutputType(AzureFileFilterDetailsPtrOutput{})
	pulumi.RegisterOutputType(AzureFileFilterDetailsResponseOutput{})
	pulumi.RegisterOutputType(AzureFileFilterDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(BlobFilterDetailsOutput{})
	pulumi.RegisterOutputType(BlobFilterDetailsPtrOutput{})
	pulumi.RegisterOutputType(BlobFilterDetailsResponseOutput{})
	pulumi.RegisterOutputType(BlobFilterDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudErrorResponseOutput{})
	pulumi.RegisterOutputType(CloudErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ContactDetailsOutput{})
	pulumi.RegisterOutputType(ContactDetailsResponseOutput{})
	pulumi.RegisterOutputType(CopyProgressResponseOutput{})
	pulumi.RegisterOutputType(CopyProgressResponseArrayOutput{})
	pulumi.RegisterOutputType(CustomerDiskJobSecretsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxAccountCopyLogDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxCustomerDiskCopyLogDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxCustomerDiskCopyProgressResponseOutput{})
	pulumi.RegisterOutputType(DataBoxCustomerDiskCopyProgressResponseArrayOutput{})
	pulumi.RegisterOutputType(DataBoxCustomerDiskJobDetailsOutput{})
	pulumi.RegisterOutputType(DataBoxCustomerDiskJobDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxDiskCopyLogDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxDiskCopyProgressResponseOutput{})
	pulumi.RegisterOutputType(DataBoxDiskCopyProgressResponseArrayOutput{})
	pulumi.RegisterOutputType(DataBoxDiskJobDetailsOutput{})
	pulumi.RegisterOutputType(DataBoxDiskJobDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxDiskJobSecretsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxHeavyAccountCopyLogDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxHeavyJobDetailsOutput{})
	pulumi.RegisterOutputType(DataBoxHeavyJobDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxHeavyJobSecretsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxHeavySecretResponseOutput{})
	pulumi.RegisterOutputType(DataBoxHeavySecretResponseArrayOutput{})
	pulumi.RegisterOutputType(DataBoxJobDetailsOutput{})
	pulumi.RegisterOutputType(DataBoxJobDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataBoxSecretResponseOutput{})
	pulumi.RegisterOutputType(DataBoxSecretResponseArrayOutput{})
	pulumi.RegisterOutputType(DataExportDetailsOutput{})
	pulumi.RegisterOutputType(DataExportDetailsArrayOutput{})
	pulumi.RegisterOutputType(DataExportDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataExportDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(DataImportDetailsOutput{})
	pulumi.RegisterOutputType(DataImportDetailsArrayOutput{})
	pulumi.RegisterOutputType(DataImportDetailsResponseOutput{})
	pulumi.RegisterOutputType(DataImportDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(DataboxJobSecretsResponseOutput{})
	pulumi.RegisterOutputType(DatacenterAddressInstructionResponseResponseOutput{})
	pulumi.RegisterOutputType(DatacenterAddressLocationResponseResponseOutput{})
	pulumi.RegisterOutputType(DcAccessSecurityCodeResponseOutput{})
	pulumi.RegisterOutputType(DiskSecretResponseOutput{})
	pulumi.RegisterOutputType(DiskSecretResponseArrayOutput{})
	pulumi.RegisterOutputType(EncryptionPreferencesOutput{})
	pulumi.RegisterOutputType(EncryptionPreferencesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPreferencesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(ExportDiskDetailsResponseMapOutput{})
	pulumi.RegisterOutputType(FilterFileDetailsOutput{})
	pulumi.RegisterOutputType(FilterFileDetailsArrayOutput{})
	pulumi.RegisterOutputType(FilterFileDetailsResponseOutput{})
	pulumi.RegisterOutputType(FilterFileDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ImportDiskDetailsOutput{})
	pulumi.RegisterOutputType(ImportDiskDetailsMapOutput{})
	pulumi.RegisterOutputType(ImportDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(ImportDiskDetailsResponseMapOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoPtrOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoResponseOutput{})
	pulumi.RegisterOutputType(JobDeliveryInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStagesResponseOutput{})
	pulumi.RegisterOutputType(JobStagesResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyEncryptionKeyOutput{})
	pulumi.RegisterOutputType(KeyEncryptionKeyPtrOutput{})
	pulumi.RegisterOutputType(KeyEncryptionKeyResponseOutput{})
	pulumi.RegisterOutputType(KeyEncryptionKeyResponsePtrOutput{})
	pulumi.RegisterOutputType(LastMitigationActionOnJobResponseOutput{})
	pulumi.RegisterOutputType(ManagedDiskDetailsOutput{})
	pulumi.RegisterOutputType(ManagedDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceArrayOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceResponseOutput{})
	pulumi.RegisterOutputType(NotificationPreferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(PackageCarrierDetailsOutput{})
	pulumi.RegisterOutputType(PackageCarrierDetailsResponseOutput{})
	pulumi.RegisterOutputType(PackageCarrierInfoResponseOutput{})
	pulumi.RegisterOutputType(PackageShippingDetailsResponseOutput{})
	pulumi.RegisterOutputType(PreferencesOutput{})
	pulumi.RegisterOutputType(PreferencesPtrOutput{})
	pulumi.RegisterOutputType(PreferencesResponseOutput{})
	pulumi.RegisterOutputType(PreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ShareCredentialDetailsResponseOutput{})
	pulumi.RegisterOutputType(ShareCredentialDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(ShippingAddressOutput{})
	pulumi.RegisterOutputType(ShippingAddressPtrOutput{})
	pulumi.RegisterOutputType(ShippingAddressResponseOutput{})
	pulumi.RegisterOutputType(ShippingAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountDetailsOutput{})
	pulumi.RegisterOutputType(StorageAccountDetailsResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TransferAllDetailsOutput{})
	pulumi.RegisterOutputType(TransferAllDetailsPtrOutput{})
	pulumi.RegisterOutputType(TransferAllDetailsResponseOutput{})
	pulumi.RegisterOutputType(TransferAllDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(TransferConfigurationOutput{})
	pulumi.RegisterOutputType(TransferConfigurationResponseOutput{})
	pulumi.RegisterOutputType(TransferConfigurationResponseTransferAllDetailsOutput{})
	pulumi.RegisterOutputType(TransferConfigurationResponseTransferAllDetailsPtrOutput{})
	pulumi.RegisterOutputType(TransferConfigurationResponseTransferFilterDetailsOutput{})
	pulumi.RegisterOutputType(TransferConfigurationResponseTransferFilterDetailsPtrOutput{})
	pulumi.RegisterOutputType(TransferConfigurationTransferAllDetailsOutput{})
	pulumi.RegisterOutputType(TransferConfigurationTransferAllDetailsPtrOutput{})
	pulumi.RegisterOutputType(TransferConfigurationTransferFilterDetailsOutput{})
	pulumi.RegisterOutputType(TransferConfigurationTransferFilterDetailsPtrOutput{})
	pulumi.RegisterOutputType(TransferFilterDetailsOutput{})
	pulumi.RegisterOutputType(TransferFilterDetailsPtrOutput{})
	pulumi.RegisterOutputType(TransferFilterDetailsResponseOutput{})
	pulumi.RegisterOutputType(TransferFilterDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(TransportPreferencesOutput{})
	pulumi.RegisterOutputType(TransportPreferencesPtrOutput{})
	pulumi.RegisterOutputType(TransportPreferencesResponseOutput{})
	pulumi.RegisterOutputType(TransportPreferencesResponsePtrOutput{})
	pulumi.RegisterOutputType(UnencryptedCredentialsResponseOutput{})
	pulumi.RegisterOutputType(UnencryptedCredentialsResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(UserAssignedPropertiesOutput{})
	pulumi.RegisterOutputType(UserAssignedPropertiesPtrOutput{})
	pulumi.RegisterOutputType(UserAssignedPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedPropertiesResponsePtrOutput{})
}
