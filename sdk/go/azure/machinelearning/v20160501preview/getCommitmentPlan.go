


package v20160501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCommitmentPlan(ctx *pulumi.Context, args *LookupCommitmentPlanArgs, opts ...pulumi.InvokeOption) (*LookupCommitmentPlanResult, error) {
	var rv LookupCommitmentPlanResult
	err := ctx.Invoke("azure-native:machinelearning/v20160501preview:getCommitmentPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCommitmentPlanArgs struct {
	CommitmentPlanName string `pulumi:"commitmentPlanName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupCommitmentPlanResult struct {
	Etag       *string                          `pulumi:"etag"`
	Id         string                           `pulumi:"id"`
	Location   string                           `pulumi:"location"`
	Name       string                           `pulumi:"name"`
	Properties CommitmentPlanPropertiesResponse `pulumi:"properties"`
	Sku        *ResourceSkuResponse             `pulumi:"sku"`
	Tags       map[string]string                `pulumi:"tags"`
	Type       string                           `pulumi:"type"`
}

func LookupCommitmentPlanOutput(ctx *pulumi.Context, args LookupCommitmentPlanOutputArgs, opts ...pulumi.InvokeOption) LookupCommitmentPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCommitmentPlanResult, error) {
			args := v.(LookupCommitmentPlanArgs)
			r, err := LookupCommitmentPlan(ctx, &args, opts...)
			var s LookupCommitmentPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCommitmentPlanResultOutput)
}

type LookupCommitmentPlanOutputArgs struct {
	CommitmentPlanName pulumi.StringInput `pulumi:"commitmentPlanName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCommitmentPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommitmentPlanArgs)(nil)).Elem()
}


type LookupCommitmentPlanResultOutput struct{ *pulumi.OutputState }

func (LookupCommitmentPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommitmentPlanResult)(nil)).Elem()
}

func (o LookupCommitmentPlanResultOutput) ToLookupCommitmentPlanResultOutput() LookupCommitmentPlanResultOutput {
	return o
}

func (o LookupCommitmentPlanResultOutput) ToLookupCommitmentPlanResultOutputWithContext(ctx context.Context) LookupCommitmentPlanResultOutput {
	return o
}

func (o LookupCommitmentPlanResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupCommitmentPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCommitmentPlanResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCommitmentPlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCommitmentPlanResultOutput) Properties() CommitmentPlanPropertiesResponseOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) CommitmentPlanPropertiesResponse { return v.Properties }).(CommitmentPlanPropertiesResponseOutput)
}

func (o LookupCommitmentPlanResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

func (o LookupCommitmentPlanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCommitmentPlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCommitmentPlanResultOutput{})
}
