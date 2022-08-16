


package v20151001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupComputePolicy(ctx *pulumi.Context, args *LookupComputePolicyArgs, opts ...pulumi.InvokeOption) (*LookupComputePolicyResult, error) {
	var rv LookupComputePolicyResult
	err := ctx.Invoke("azure-native:datalakeanalytics/v20151001preview:getComputePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComputePolicyArgs struct {
	AccountName       string `pulumi:"accountName"`
	ComputePolicyName string `pulumi:"computePolicyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupComputePolicyResult struct {
	Id                           string `pulumi:"id"`
	MaxDegreeOfParallelismPerJob int    `pulumi:"maxDegreeOfParallelismPerJob"`
	MinPriorityPerJob            int    `pulumi:"minPriorityPerJob"`
	Name                         string `pulumi:"name"`
	ObjectId                     string `pulumi:"objectId"`
	ObjectType                   string `pulumi:"objectType"`
	Type                         string `pulumi:"type"`
}

func LookupComputePolicyOutput(ctx *pulumi.Context, args LookupComputePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupComputePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComputePolicyResult, error) {
			args := v.(LookupComputePolicyArgs)
			r, err := LookupComputePolicy(ctx, &args, opts...)
			var s LookupComputePolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComputePolicyResultOutput)
}

type LookupComputePolicyOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ComputePolicyName pulumi.StringInput `pulumi:"computePolicyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupComputePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputePolicyArgs)(nil)).Elem()
}


type LookupComputePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupComputePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputePolicyResult)(nil)).Elem()
}

func (o LookupComputePolicyResultOutput) ToLookupComputePolicyResultOutput() LookupComputePolicyResultOutput {
	return o
}

func (o LookupComputePolicyResultOutput) ToLookupComputePolicyResultOutputWithContext(ctx context.Context) LookupComputePolicyResultOutput {
	return o
}

func (o LookupComputePolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputePolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputePolicyResultOutput) MaxDegreeOfParallelismPerJob() pulumi.IntOutput {
	return o.ApplyT(func(v LookupComputePolicyResult) int { return v.MaxDegreeOfParallelismPerJob }).(pulumi.IntOutput)
}

func (o LookupComputePolicyResultOutput) MinPriorityPerJob() pulumi.IntOutput {
	return o.ApplyT(func(v LookupComputePolicyResult) int { return v.MinPriorityPerJob }).(pulumi.IntOutput)
}

func (o LookupComputePolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputePolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputePolicyResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputePolicyResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o LookupComputePolicyResultOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputePolicyResult) string { return v.ObjectType }).(pulumi.StringOutput)
}

func (o LookupComputePolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputePolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputePolicyResultOutput{})
}
