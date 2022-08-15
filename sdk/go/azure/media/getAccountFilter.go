


package media

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccountFilter(ctx *pulumi.Context, args *LookupAccountFilterArgs, opts ...pulumi.InvokeOption) (*LookupAccountFilterResult, error) {
	var rv LookupAccountFilterResult
	err := ctx.Invoke("azure-native:media:getAccountFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountFilterArgs struct {
	AccountName       string `pulumi:"accountName"`
	FilterName        string `pulumi:"filterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountFilterResult struct {
	FirstQuality          *FirstQualityResponse          `pulumi:"firstQuality"`
	Id                    string                         `pulumi:"id"`
	Name                  string                         `pulumi:"name"`
	PresentationTimeRange *PresentationTimeRangeResponse `pulumi:"presentationTimeRange"`
	SystemData            SystemDataResponse             `pulumi:"systemData"`
	Tracks                []FilterTrackSelectionResponse `pulumi:"tracks"`
	Type                  string                         `pulumi:"type"`
}

func LookupAccountFilterOutput(ctx *pulumi.Context, args LookupAccountFilterOutputArgs, opts ...pulumi.InvokeOption) LookupAccountFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountFilterResult, error) {
			args := v.(LookupAccountFilterArgs)
			r, err := LookupAccountFilter(ctx, &args, opts...)
			var s LookupAccountFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountFilterResultOutput)
}

type LookupAccountFilterOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	FilterName        pulumi.StringInput `pulumi:"filterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccountFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountFilterArgs)(nil)).Elem()
}


type LookupAccountFilterResultOutput struct{ *pulumi.OutputState }

func (LookupAccountFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountFilterResult)(nil)).Elem()
}

func (o LookupAccountFilterResultOutput) ToLookupAccountFilterResultOutput() LookupAccountFilterResultOutput {
	return o
}

func (o LookupAccountFilterResultOutput) ToLookupAccountFilterResultOutputWithContext(ctx context.Context) LookupAccountFilterResultOutput {
	return o
}

func (o LookupAccountFilterResultOutput) FirstQuality() FirstQualityResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountFilterResult) *FirstQualityResponse { return v.FirstQuality }).(FirstQualityResponsePtrOutput)
}

func (o LookupAccountFilterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountFilterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccountFilterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountFilterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccountFilterResultOutput) PresentationTimeRange() PresentationTimeRangeResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountFilterResult) *PresentationTimeRangeResponse { return v.PresentationTimeRange }).(PresentationTimeRangeResponsePtrOutput)
}

func (o LookupAccountFilterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAccountFilterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAccountFilterResultOutput) Tracks() FilterTrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountFilterResult) []FilterTrackSelectionResponse { return v.Tracks }).(FilterTrackSelectionResponseArrayOutput)
}

func (o LookupAccountFilterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountFilterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountFilterResultOutput{})
}
