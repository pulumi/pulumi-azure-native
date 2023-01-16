


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMarketplaceImageSasTokenSASToken(ctx *pulumi.Context, args *GetMarketplaceImageSasTokenSASTokenArgs, opts ...pulumi.InvokeOption) (*GetMarketplaceImageSasTokenSASTokenResult, error) {
	var rv GetMarketplaceImageSasTokenSASTokenResult
	err := ctx.Invoke("azure-native:databoxedge/v20221201preview:getMarketplaceImageSasTokenSASToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetMarketplaceImageSasTokenSASTokenArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	OfferName         string `pulumi:"offerName"`
	PublisherName     string `pulumi:"publisherName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SkuName           string `pulumi:"skuName"`
	VersionName       string `pulumi:"versionName"`
}

type GetMarketplaceImageSasTokenSASTokenResult struct {
	SasUri *string `pulumi:"sasUri"`
	Status *string `pulumi:"status"`
}

func GetMarketplaceImageSasTokenSASTokenOutput(ctx *pulumi.Context, args GetMarketplaceImageSasTokenSASTokenOutputArgs, opts ...pulumi.InvokeOption) GetMarketplaceImageSasTokenSASTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMarketplaceImageSasTokenSASTokenResult, error) {
			args := v.(GetMarketplaceImageSasTokenSASTokenArgs)
			r, err := GetMarketplaceImageSasTokenSASToken(ctx, &args, opts...)
			var s GetMarketplaceImageSasTokenSASTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMarketplaceImageSasTokenSASTokenResultOutput)
}

type GetMarketplaceImageSasTokenSASTokenOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	OfferName         pulumi.StringInput `pulumi:"offerName"`
	PublisherName     pulumi.StringInput `pulumi:"publisherName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SkuName           pulumi.StringInput `pulumi:"skuName"`
	VersionName       pulumi.StringInput `pulumi:"versionName"`
}

func (GetMarketplaceImageSasTokenSASTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMarketplaceImageSasTokenSASTokenArgs)(nil)).Elem()
}

type GetMarketplaceImageSasTokenSASTokenResultOutput struct{ *pulumi.OutputState }

func (GetMarketplaceImageSasTokenSASTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMarketplaceImageSasTokenSASTokenResult)(nil)).Elem()
}

func (o GetMarketplaceImageSasTokenSASTokenResultOutput) ToGetMarketplaceImageSasTokenSASTokenResultOutput() GetMarketplaceImageSasTokenSASTokenResultOutput {
	return o
}

func (o GetMarketplaceImageSasTokenSASTokenResultOutput) ToGetMarketplaceImageSasTokenSASTokenResultOutputWithContext(ctx context.Context) GetMarketplaceImageSasTokenSASTokenResultOutput {
	return o
}

func (o GetMarketplaceImageSasTokenSASTokenResultOutput) SasUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMarketplaceImageSasTokenSASTokenResult) *string { return v.SasUri }).(pulumi.StringPtrOutput)
}

func (o GetMarketplaceImageSasTokenSASTokenResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMarketplaceImageSasTokenSASTokenResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMarketplaceImageSasTokenSASTokenResultOutput{})
}
