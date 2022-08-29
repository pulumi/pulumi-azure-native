


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupMachineLearningService(ctx *pulumi.Context, args *LookupMachineLearningServiceArgs, opts ...pulumi.InvokeOption) (*LookupMachineLearningServiceResult, error) {
	var rv LookupMachineLearningServiceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200901preview:getMachineLearningService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineLearningServiceArgs struct {
	Expand            *bool  `pulumi:"expand"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMachineLearningServiceResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupMachineLearningServiceOutput(ctx *pulumi.Context, args LookupMachineLearningServiceOutputArgs, opts ...pulumi.InvokeOption) LookupMachineLearningServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachineLearningServiceResult, error) {
			args := v.(LookupMachineLearningServiceArgs)
			r, err := LookupMachineLearningService(ctx, &args, opts...)
			var s LookupMachineLearningServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachineLearningServiceResultOutput)
}

type LookupMachineLearningServiceOutputArgs struct {
	Expand            pulumi.BoolPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput  `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput  `pulumi:"serviceName"`
	WorkspaceName     pulumi.StringInput  `pulumi:"workspaceName"`
}

func (LookupMachineLearningServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineLearningServiceArgs)(nil)).Elem()
}


type LookupMachineLearningServiceResultOutput struct{ *pulumi.OutputState }

func (LookupMachineLearningServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineLearningServiceResult)(nil)).Elem()
}

func (o LookupMachineLearningServiceResultOutput) ToLookupMachineLearningServiceResultOutput() LookupMachineLearningServiceResultOutput {
	return o
}

func (o LookupMachineLearningServiceResultOutput) ToLookupMachineLearningServiceResultOutputWithContext(ctx context.Context) LookupMachineLearningServiceResultOutput {
	return o
}

func (o LookupMachineLearningServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMachineLearningServiceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineLearningServiceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupMachineLearningServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineLearningServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMachineLearningServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachineLearningServiceResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMachineLearningServiceResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupMachineLearningServiceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineLearningServiceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupMachineLearningServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMachineLearningServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMachineLearningServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineLearningServiceResultOutput{})
}
