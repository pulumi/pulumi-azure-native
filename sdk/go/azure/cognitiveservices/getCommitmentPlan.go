


package cognitiveservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCommitmentPlan(ctx *pulumi.Context, args *LookupCommitmentPlanArgs, opts ...pulumi.InvokeOption) (*LookupCommitmentPlanResult, error) {
	var rv LookupCommitmentPlanResult
	err := ctx.Invoke("azure-native:cognitiveservices:getCommitmentPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCommitmentPlanArgs struct {
	AccountName        string `pulumi:"accountName"`
	CommitmentPlanName string `pulumi:"commitmentPlanName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupCommitmentPlanResult struct {
	Etag       string                           `pulumi:"etag"`
	Id         string                           `pulumi:"id"`
	Name       string                           `pulumi:"name"`
	Properties CommitmentPlanPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse               `pulumi:"systemData"`
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
	AccountName        pulumi.StringInput `pulumi:"accountName"`
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

func (o LookupCommitmentPlanResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupCommitmentPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCommitmentPlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCommitmentPlanResultOutput) Properties() CommitmentPlanPropertiesResponseOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) CommitmentPlanPropertiesResponse { return v.Properties }).(CommitmentPlanPropertiesResponseOutput)
}

func (o LookupCommitmentPlanResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCommitmentPlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCommitmentPlanResultOutput{})
}
