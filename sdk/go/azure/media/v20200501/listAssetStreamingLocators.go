


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAssetStreamingLocators(ctx *pulumi.Context, args *ListAssetStreamingLocatorsArgs, opts ...pulumi.InvokeOption) (*ListAssetStreamingLocatorsResult, error) {
	var rv ListAssetStreamingLocatorsResult
	err := ctx.Invoke("azure-native:media/v20200501:listAssetStreamingLocators", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAssetStreamingLocatorsArgs struct {
	AccountName       string `pulumi:"accountName"`
	AssetName         string `pulumi:"assetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAssetStreamingLocatorsResult struct {
	StreamingLocators []AssetStreamingLocatorResponse `pulumi:"streamingLocators"`
}

func ListAssetStreamingLocatorsOutput(ctx *pulumi.Context, args ListAssetStreamingLocatorsOutputArgs, opts ...pulumi.InvokeOption) ListAssetStreamingLocatorsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAssetStreamingLocatorsResult, error) {
			args := v.(ListAssetStreamingLocatorsArgs)
			r, err := ListAssetStreamingLocators(ctx, &args, opts...)
			var s ListAssetStreamingLocatorsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAssetStreamingLocatorsResultOutput)
}

type ListAssetStreamingLocatorsOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	AssetName         pulumi.StringInput `pulumi:"assetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAssetStreamingLocatorsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAssetStreamingLocatorsArgs)(nil)).Elem()
}


type ListAssetStreamingLocatorsResultOutput struct{ *pulumi.OutputState }

func (ListAssetStreamingLocatorsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAssetStreamingLocatorsResult)(nil)).Elem()
}

func (o ListAssetStreamingLocatorsResultOutput) ToListAssetStreamingLocatorsResultOutput() ListAssetStreamingLocatorsResultOutput {
	return o
}

func (o ListAssetStreamingLocatorsResultOutput) ToListAssetStreamingLocatorsResultOutputWithContext(ctx context.Context) ListAssetStreamingLocatorsResultOutput {
	return o
}

func (o ListAssetStreamingLocatorsResultOutput) StreamingLocators() AssetStreamingLocatorResponseArrayOutput {
	return o.ApplyT(func(v ListAssetStreamingLocatorsResult) []AssetStreamingLocatorResponse { return v.StreamingLocators }).(AssetStreamingLocatorResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAssetStreamingLocatorsResultOutput{})
}
