


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCompute(ctx *pulumi.Context, args *LookupComputeArgs, opts ...pulumi.InvokeOption) (*LookupComputeResult, error) {
	var rv LookupComputeResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:getCompute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComputeArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupComputeResult struct {
	Id         string                          `pulumi:"id"`
	Identity   *ManagedServiceIdentityResponse `pulumi:"identity"`
	Location   *string                         `pulumi:"location"`
	Name       string                          `pulumi:"name"`
	Properties interface{}                     `pulumi:"properties"`
	Sku        *SkuResponse                    `pulumi:"sku"`
	SystemData SystemDataResponse              `pulumi:"systemData"`
	Tags       map[string]string               `pulumi:"tags"`
	Type       string                          `pulumi:"type"`
}

func LookupComputeOutput(ctx *pulumi.Context, args LookupComputeOutputArgs, opts ...pulumi.InvokeOption) LookupComputeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComputeResult, error) {
			args := v.(LookupComputeArgs)
			r, err := LookupCompute(ctx, &args, opts...)
			var s LookupComputeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComputeResultOutput)
}

type LookupComputeOutputArgs struct {
	ComputeName       pulumi.StringInput `pulumi:"computeName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupComputeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeArgs)(nil)).Elem()
}


type LookupComputeResultOutput struct{ *pulumi.OutputState }

func (LookupComputeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeResult)(nil)).Elem()
}

func (o LookupComputeResultOutput) ToLookupComputeResultOutput() LookupComputeResultOutput {
	return o
}

func (o LookupComputeResultOutput) ToLookupComputeResultOutputWithContext(ctx context.Context) LookupComputeResultOutput {
	return o
}

func (o LookupComputeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupComputeResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupComputeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupComputeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupComputeResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupComputeResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupComputeResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupComputeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupComputeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupComputeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComputeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupComputeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeResultOutput{})
}
