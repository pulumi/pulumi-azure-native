


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSkusNestedResourceTypeFirst(ctx *pulumi.Context, args *LookupSkusNestedResourceTypeFirstArgs, opts ...pulumi.InvokeOption) (*LookupSkusNestedResourceTypeFirstResult, error) {
	var rv LookupSkusNestedResourceTypeFirstResult
	err := ctx.Invoke("azure-native:providerhub/v20210601preview:getSkusNestedResourceTypeFirst", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSkusNestedResourceTypeFirstArgs struct {
	NestedResourceTypeFirst string `pulumi:"nestedResourceTypeFirst"`
	ProviderNamespace       string `pulumi:"providerNamespace"`
	ResourceType            string `pulumi:"resourceType"`
	Sku                     string `pulumi:"sku"`
}

type LookupSkusNestedResourceTypeFirstResult struct {
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties SkuResourceResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Type       string                        `pulumi:"type"`
}

func LookupSkusNestedResourceTypeFirstOutput(ctx *pulumi.Context, args LookupSkusNestedResourceTypeFirstOutputArgs, opts ...pulumi.InvokeOption) LookupSkusNestedResourceTypeFirstResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSkusNestedResourceTypeFirstResult, error) {
			args := v.(LookupSkusNestedResourceTypeFirstArgs)
			r, err := LookupSkusNestedResourceTypeFirst(ctx, &args, opts...)
			var s LookupSkusNestedResourceTypeFirstResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSkusNestedResourceTypeFirstResultOutput)
}

type LookupSkusNestedResourceTypeFirstOutputArgs struct {
	NestedResourceTypeFirst pulumi.StringInput `pulumi:"nestedResourceTypeFirst"`
	ProviderNamespace       pulumi.StringInput `pulumi:"providerNamespace"`
	ResourceType            pulumi.StringInput `pulumi:"resourceType"`
	Sku                     pulumi.StringInput `pulumi:"sku"`
}

func (LookupSkusNestedResourceTypeFirstOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusNestedResourceTypeFirstArgs)(nil)).Elem()
}

type LookupSkusNestedResourceTypeFirstResultOutput struct{ *pulumi.OutputState }

func (LookupSkusNestedResourceTypeFirstResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusNestedResourceTypeFirstResult)(nil)).Elem()
}

func (o LookupSkusNestedResourceTypeFirstResultOutput) ToLookupSkusNestedResourceTypeFirstResultOutput() LookupSkusNestedResourceTypeFirstResultOutput {
	return o
}

func (o LookupSkusNestedResourceTypeFirstResultOutput) ToLookupSkusNestedResourceTypeFirstResultOutputWithContext(ctx context.Context) LookupSkusNestedResourceTypeFirstResultOutput {
	return o
}

func (o LookupSkusNestedResourceTypeFirstResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeFirstResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSkusNestedResourceTypeFirstResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeFirstResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSkusNestedResourceTypeFirstResultOutput) Properties() SkuResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeFirstResult) SkuResourceResponseProperties { return v.Properties }).(SkuResourceResponsePropertiesOutput)
}

func (o LookupSkusNestedResourceTypeFirstResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeFirstResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSkusNestedResourceTypeFirstResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeFirstResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSkusNestedResourceTypeFirstResultOutput{})
}
