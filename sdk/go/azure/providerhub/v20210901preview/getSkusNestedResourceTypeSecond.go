


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSkusNestedResourceTypeSecond(ctx *pulumi.Context, args *LookupSkusNestedResourceTypeSecondArgs, opts ...pulumi.InvokeOption) (*LookupSkusNestedResourceTypeSecondResult, error) {
	var rv LookupSkusNestedResourceTypeSecondResult
	err := ctx.Invoke("azure-native:providerhub/v20210901preview:getSkusNestedResourceTypeSecond", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSkusNestedResourceTypeSecondArgs struct {
	NestedResourceTypeFirst  string `pulumi:"nestedResourceTypeFirst"`
	NestedResourceTypeSecond string `pulumi:"nestedResourceTypeSecond"`
	ProviderNamespace        string `pulumi:"providerNamespace"`
	ResourceType             string `pulumi:"resourceType"`
	Sku                      string `pulumi:"sku"`
}

type LookupSkusNestedResourceTypeSecondResult struct {
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties SkuResourceResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Type       string                        `pulumi:"type"`
}

func LookupSkusNestedResourceTypeSecondOutput(ctx *pulumi.Context, args LookupSkusNestedResourceTypeSecondOutputArgs, opts ...pulumi.InvokeOption) LookupSkusNestedResourceTypeSecondResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSkusNestedResourceTypeSecondResult, error) {
			args := v.(LookupSkusNestedResourceTypeSecondArgs)
			r, err := LookupSkusNestedResourceTypeSecond(ctx, &args, opts...)
			var s LookupSkusNestedResourceTypeSecondResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSkusNestedResourceTypeSecondResultOutput)
}

type LookupSkusNestedResourceTypeSecondOutputArgs struct {
	NestedResourceTypeFirst  pulumi.StringInput `pulumi:"nestedResourceTypeFirst"`
	NestedResourceTypeSecond pulumi.StringInput `pulumi:"nestedResourceTypeSecond"`
	ProviderNamespace        pulumi.StringInput `pulumi:"providerNamespace"`
	ResourceType             pulumi.StringInput `pulumi:"resourceType"`
	Sku                      pulumi.StringInput `pulumi:"sku"`
}

func (LookupSkusNestedResourceTypeSecondOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusNestedResourceTypeSecondArgs)(nil)).Elem()
}

type LookupSkusNestedResourceTypeSecondResultOutput struct{ *pulumi.OutputState }

func (LookupSkusNestedResourceTypeSecondResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusNestedResourceTypeSecondResult)(nil)).Elem()
}

func (o LookupSkusNestedResourceTypeSecondResultOutput) ToLookupSkusNestedResourceTypeSecondResultOutput() LookupSkusNestedResourceTypeSecondResultOutput {
	return o
}

func (o LookupSkusNestedResourceTypeSecondResultOutput) ToLookupSkusNestedResourceTypeSecondResultOutputWithContext(ctx context.Context) LookupSkusNestedResourceTypeSecondResultOutput {
	return o
}

func (o LookupSkusNestedResourceTypeSecondResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeSecondResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSkusNestedResourceTypeSecondResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeSecondResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSkusNestedResourceTypeSecondResultOutput) Properties() SkuResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeSecondResult) SkuResourceResponseProperties { return v.Properties }).(SkuResourceResponsePropertiesOutput)
}

func (o LookupSkusNestedResourceTypeSecondResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeSecondResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSkusNestedResourceTypeSecondResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeSecondResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSkusNestedResourceTypeSecondResultOutput{})
}
