


package v20200515preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAKSService(ctx *pulumi.Context, args *LookupAKSServiceArgs, opts ...pulumi.InvokeOption) (*LookupAKSServiceResult, error) {
	var rv LookupAKSServiceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200515preview:getAKSService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAKSServiceArgs struct {
	Expand            *bool  `pulumi:"expand"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAKSServiceResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupAKSServiceOutput(ctx *pulumi.Context, args LookupAKSServiceOutputArgs, opts ...pulumi.InvokeOption) LookupAKSServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAKSServiceResult, error) {
			args := v.(LookupAKSServiceArgs)
			r, err := LookupAKSService(ctx, &args, opts...)
			var s LookupAKSServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAKSServiceResultOutput)
}

type LookupAKSServiceOutputArgs struct {
	Expand            pulumi.BoolPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput  `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput  `pulumi:"serviceName"`
	WorkspaceName     pulumi.StringInput  `pulumi:"workspaceName"`
}

func (LookupAKSServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAKSServiceArgs)(nil)).Elem()
}


type LookupAKSServiceResultOutput struct{ *pulumi.OutputState }

func (LookupAKSServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAKSServiceResult)(nil)).Elem()
}

func (o LookupAKSServiceResultOutput) ToLookupAKSServiceResultOutput() LookupAKSServiceResultOutput {
	return o
}

func (o LookupAKSServiceResultOutput) ToLookupAKSServiceResultOutputWithContext(ctx context.Context) LookupAKSServiceResultOutput {
	return o
}

func (o LookupAKSServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAKSServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAKSServiceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAKSServiceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupAKSServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAKSServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAKSServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAKSServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAKSServiceResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAKSServiceResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupAKSServiceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupAKSServiceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupAKSServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAKSServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAKSServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAKSServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAKSServiceResultOutput{})
}
