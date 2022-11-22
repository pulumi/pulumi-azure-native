


package v20180701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssetFilter(ctx *pulumi.Context, args *LookupAssetFilterArgs, opts ...pulumi.InvokeOption) (*LookupAssetFilterResult, error) {
	var rv LookupAssetFilterResult
	err := ctx.Invoke("azure-native:media/v20180701:getAssetFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssetFilterArgs struct {
	AccountName       string `pulumi:"accountName"`
	AssetName         string `pulumi:"assetName"`
	FilterName        string `pulumi:"filterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAssetFilterResult struct {
	FirstQuality          *FirstQualityResponse          `pulumi:"firstQuality"`
	Id                    string                         `pulumi:"id"`
	Name                  string                         `pulumi:"name"`
	PresentationTimeRange *PresentationTimeRangeResponse `pulumi:"presentationTimeRange"`
	Tracks                []FilterTrackSelectionResponse `pulumi:"tracks"`
	Type                  string                         `pulumi:"type"`
}

func LookupAssetFilterOutput(ctx *pulumi.Context, args LookupAssetFilterOutputArgs, opts ...pulumi.InvokeOption) LookupAssetFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssetFilterResult, error) {
			args := v.(LookupAssetFilterArgs)
			r, err := LookupAssetFilter(ctx, &args, opts...)
			var s LookupAssetFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssetFilterResultOutput)
}

type LookupAssetFilterOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	AssetName         pulumi.StringInput `pulumi:"assetName"`
	FilterName        pulumi.StringInput `pulumi:"filterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAssetFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetFilterArgs)(nil)).Elem()
}


type LookupAssetFilterResultOutput struct{ *pulumi.OutputState }

func (LookupAssetFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetFilterResult)(nil)).Elem()
}

func (o LookupAssetFilterResultOutput) ToLookupAssetFilterResultOutput() LookupAssetFilterResultOutput {
	return o
}

func (o LookupAssetFilterResultOutput) ToLookupAssetFilterResultOutputWithContext(ctx context.Context) LookupAssetFilterResultOutput {
	return o
}

func (o LookupAssetFilterResultOutput) FirstQuality() FirstQualityResponsePtrOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) *FirstQualityResponse { return v.FirstQuality }).(FirstQualityResponsePtrOutput)
}

func (o LookupAssetFilterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAssetFilterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAssetFilterResultOutput) PresentationTimeRange() PresentationTimeRangeResponsePtrOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) *PresentationTimeRangeResponse { return v.PresentationTimeRange }).(PresentationTimeRangeResponsePtrOutput)
}

func (o LookupAssetFilterResultOutput) Tracks() FilterTrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) []FilterTrackSelectionResponse { return v.Tracks }).(FilterTrackSelectionResponseArrayOutput)
}

func (o LookupAssetFilterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssetFilterResultOutput{})
}
