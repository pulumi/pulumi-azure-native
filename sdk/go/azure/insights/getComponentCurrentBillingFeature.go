


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComponentCurrentBillingFeature(ctx *pulumi.Context, args *LookupComponentCurrentBillingFeatureArgs, opts ...pulumi.InvokeOption) (*LookupComponentCurrentBillingFeatureResult, error) {
	var rv LookupComponentCurrentBillingFeatureResult
	err := ctx.Invoke("azure-native:insights:getComponentCurrentBillingFeature", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComponentCurrentBillingFeatureArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupComponentCurrentBillingFeatureResult struct {
	CurrentBillingFeatures []string                                           `pulumi:"currentBillingFeatures"`
	DataVolumeCap          *ApplicationInsightsComponentDataVolumeCapResponse `pulumi:"dataVolumeCap"`
}

func LookupComponentCurrentBillingFeatureOutput(ctx *pulumi.Context, args LookupComponentCurrentBillingFeatureOutputArgs, opts ...pulumi.InvokeOption) LookupComponentCurrentBillingFeatureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComponentCurrentBillingFeatureResult, error) {
			args := v.(LookupComponentCurrentBillingFeatureArgs)
			r, err := LookupComponentCurrentBillingFeature(ctx, &args, opts...)
			var s LookupComponentCurrentBillingFeatureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComponentCurrentBillingFeatureResultOutput)
}

type LookupComponentCurrentBillingFeatureOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupComponentCurrentBillingFeatureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentCurrentBillingFeatureArgs)(nil)).Elem()
}


type LookupComponentCurrentBillingFeatureResultOutput struct{ *pulumi.OutputState }

func (LookupComponentCurrentBillingFeatureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentCurrentBillingFeatureResult)(nil)).Elem()
}

func (o LookupComponentCurrentBillingFeatureResultOutput) ToLookupComponentCurrentBillingFeatureResultOutput() LookupComponentCurrentBillingFeatureResultOutput {
	return o
}

func (o LookupComponentCurrentBillingFeatureResultOutput) ToLookupComponentCurrentBillingFeatureResultOutputWithContext(ctx context.Context) LookupComponentCurrentBillingFeatureResultOutput {
	return o
}

func (o LookupComponentCurrentBillingFeatureResultOutput) CurrentBillingFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComponentCurrentBillingFeatureResult) []string { return v.CurrentBillingFeatures }).(pulumi.StringArrayOutput)
}

func (o LookupComponentCurrentBillingFeatureResultOutput) DataVolumeCap() ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return o.ApplyT(func(v LookupComponentCurrentBillingFeatureResult) *ApplicationInsightsComponentDataVolumeCapResponse {
		return v.DataVolumeCap
	}).(ApplicationInsightsComponentDataVolumeCapResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComponentCurrentBillingFeatureResultOutput{})
}
