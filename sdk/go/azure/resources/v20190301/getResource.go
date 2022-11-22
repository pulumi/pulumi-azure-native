


package v20190301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupResource(ctx *pulumi.Context, args *LookupResourceArgs, opts ...pulumi.InvokeOption) (*LookupResourceResult, error) {
	var rv LookupResourceResult
	err := ctx.Invoke("azure-native:resources/v20190301:getResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceArgs struct {
	ParentResourcePath        string `pulumi:"parentResourcePath"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	ResourceName              string `pulumi:"resourceName"`
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	ResourceType              string `pulumi:"resourceType"`
}


type LookupResourceResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Kind       *string           `pulumi:"kind"`
	Location   *string           `pulumi:"location"`
	ManagedBy  *string           `pulumi:"managedBy"`
	Name       string            `pulumi:"name"`
	Plan       *PlanResponse     `pulumi:"plan"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupResourceOutput(ctx *pulumi.Context, args LookupResourceOutputArgs, opts ...pulumi.InvokeOption) LookupResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceResult, error) {
			args := v.(LookupResourceArgs)
			r, err := LookupResource(ctx, &args, opts...)
			var s LookupResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceResultOutput)
}

type LookupResourceOutputArgs struct {
	ParentResourcePath        pulumi.StringInput `pulumi:"parentResourcePath"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName              pulumi.StringInput `pulumi:"resourceName"`
	ResourceProviderNamespace pulumi.StringInput `pulumi:"resourceProviderNamespace"`
	ResourceType              pulumi.StringInput `pulumi:"resourceType"`
}

func (LookupResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceArgs)(nil)).Elem()
}


type LookupResourceResultOutput struct{ *pulumi.OutputState }

func (LookupResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceResult)(nil)).Elem()
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutput() LookupResourceResultOutput {
	return o
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutputWithContext(ctx context.Context) LookupResourceResultOutput {
	return o
}

func (o LookupResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResourceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupResourceResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupResourceResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o LookupResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupResourceResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

func (o LookupResourceResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupResourceResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupResourceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceResultOutput{})
}
