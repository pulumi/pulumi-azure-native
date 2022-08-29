


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEndpointVariant(ctx *pulumi.Context, args *LookupEndpointVariantArgs, opts ...pulumi.InvokeOption) (*LookupEndpointVariantResult, error) {
	var rv LookupEndpointVariantResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200901preview:getEndpointVariant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointVariantArgs struct {
	Expand            *bool  `pulumi:"expand"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupEndpointVariantResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupEndpointVariantOutput(ctx *pulumi.Context, args LookupEndpointVariantOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointVariantResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointVariantResult, error) {
			args := v.(LookupEndpointVariantArgs)
			r, err := LookupEndpointVariant(ctx, &args, opts...)
			var s LookupEndpointVariantResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointVariantResultOutput)
}

type LookupEndpointVariantOutputArgs struct {
	Expand            pulumi.BoolPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput  `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput  `pulumi:"serviceName"`
	WorkspaceName     pulumi.StringInput  `pulumi:"workspaceName"`
}

func (LookupEndpointVariantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointVariantArgs)(nil)).Elem()
}


type LookupEndpointVariantResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointVariantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointVariantResult)(nil)).Elem()
}

func (o LookupEndpointVariantResultOutput) ToLookupEndpointVariantResultOutput() LookupEndpointVariantResultOutput {
	return o
}

func (o LookupEndpointVariantResultOutput) ToLookupEndpointVariantResultOutputWithContext(ctx context.Context) LookupEndpointVariantResultOutput {
	return o
}

func (o LookupEndpointVariantResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEndpointVariantResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupEndpointVariantResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointVariantResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEndpointVariantResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupEndpointVariantResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupEndpointVariantResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEndpointVariantResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointVariantResultOutput{})
}
