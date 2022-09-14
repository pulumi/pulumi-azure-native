


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAssetContainerSas(ctx *pulumi.Context, args *ListAssetContainerSasArgs, opts ...pulumi.InvokeOption) (*ListAssetContainerSasResult, error) {
	var rv ListAssetContainerSasResult
	err := ctx.Invoke("azure-native:media/v20200501:listAssetContainerSas", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAssetContainerSasArgs struct {
	AccountName       string  `pulumi:"accountName"`
	AssetName         string  `pulumi:"assetName"`
	ExpiryTime        *string `pulumi:"expiryTime"`
	Permissions       *string `pulumi:"permissions"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type ListAssetContainerSasResult struct {
	AssetContainerSasUrls []string `pulumi:"assetContainerSasUrls"`
}

func ListAssetContainerSasOutput(ctx *pulumi.Context, args ListAssetContainerSasOutputArgs, opts ...pulumi.InvokeOption) ListAssetContainerSasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAssetContainerSasResult, error) {
			args := v.(ListAssetContainerSasArgs)
			r, err := ListAssetContainerSas(ctx, &args, opts...)
			var s ListAssetContainerSasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAssetContainerSasResultOutput)
}

type ListAssetContainerSasOutputArgs struct {
	AccountName       pulumi.StringInput    `pulumi:"accountName"`
	AssetName         pulumi.StringInput    `pulumi:"assetName"`
	ExpiryTime        pulumi.StringPtrInput `pulumi:"expiryTime"`
	Permissions       pulumi.StringPtrInput `pulumi:"permissions"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListAssetContainerSasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAssetContainerSasArgs)(nil)).Elem()
}


type ListAssetContainerSasResultOutput struct{ *pulumi.OutputState }

func (ListAssetContainerSasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAssetContainerSasResult)(nil)).Elem()
}

func (o ListAssetContainerSasResultOutput) ToListAssetContainerSasResultOutput() ListAssetContainerSasResultOutput {
	return o
}

func (o ListAssetContainerSasResultOutput) ToListAssetContainerSasResultOutputWithContext(ctx context.Context) ListAssetContainerSasResultOutput {
	return o
}

func (o ListAssetContainerSasResultOutput) AssetContainerSasUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListAssetContainerSasResult) []string { return v.AssetContainerSasUrls }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAssetContainerSasResultOutput{})
}
