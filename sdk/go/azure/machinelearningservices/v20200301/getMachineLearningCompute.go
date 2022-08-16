


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachineLearningCompute(ctx *pulumi.Context, args *LookupMachineLearningComputeArgs, opts ...pulumi.InvokeOption) (*LookupMachineLearningComputeResult, error) {
	var rv LookupMachineLearningComputeResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200301:getMachineLearningCompute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineLearningComputeArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMachineLearningComputeResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupMachineLearningComputeOutput(ctx *pulumi.Context, args LookupMachineLearningComputeOutputArgs, opts ...pulumi.InvokeOption) LookupMachineLearningComputeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachineLearningComputeResult, error) {
			args := v.(LookupMachineLearningComputeArgs)
			r, err := LookupMachineLearningCompute(ctx, &args, opts...)
			var s LookupMachineLearningComputeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachineLearningComputeResultOutput)
}

type LookupMachineLearningComputeOutputArgs struct {
	ComputeName       pulumi.StringInput `pulumi:"computeName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMachineLearningComputeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineLearningComputeArgs)(nil)).Elem()
}


type LookupMachineLearningComputeResultOutput struct{ *pulumi.OutputState }

func (LookupMachineLearningComputeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineLearningComputeResult)(nil)).Elem()
}

func (o LookupMachineLearningComputeResultOutput) ToLookupMachineLearningComputeResultOutput() LookupMachineLearningComputeResultOutput {
	return o
}

func (o LookupMachineLearningComputeResultOutput) ToLookupMachineLearningComputeResultOutputWithContext(ctx context.Context) LookupMachineLearningComputeResultOutput {
	return o
}

func (o LookupMachineLearningComputeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningComputeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMachineLearningComputeResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineLearningComputeResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupMachineLearningComputeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineLearningComputeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMachineLearningComputeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningComputeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachineLearningComputeResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMachineLearningComputeResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupMachineLearningComputeResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineLearningComputeResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupMachineLearningComputeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMachineLearningComputeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMachineLearningComputeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningComputeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineLearningComputeResultOutput{})
}
