


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSkusNestedResourceTypeThird(ctx *pulumi.Context, args *LookupSkusNestedResourceTypeThirdArgs, opts ...pulumi.InvokeOption) (*LookupSkusNestedResourceTypeThirdResult, error) {
	var rv LookupSkusNestedResourceTypeThirdResult
	err := ctx.Invoke("azure-native:providerhub/v20210901preview:getSkusNestedResourceTypeThird", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSkusNestedResourceTypeThirdArgs struct {
	NestedResourceTypeFirst  string `pulumi:"nestedResourceTypeFirst"`
	NestedResourceTypeSecond string `pulumi:"nestedResourceTypeSecond"`
	NestedResourceTypeThird  string `pulumi:"nestedResourceTypeThird"`
	ProviderNamespace        string `pulumi:"providerNamespace"`
	ResourceType             string `pulumi:"resourceType"`
	Sku                      string `pulumi:"sku"`
}

type LookupSkusNestedResourceTypeThirdResult struct {
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties SkuResourceResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Type       string                        `pulumi:"type"`
}

func LookupSkusNestedResourceTypeThirdOutput(ctx *pulumi.Context, args LookupSkusNestedResourceTypeThirdOutputArgs, opts ...pulumi.InvokeOption) LookupSkusNestedResourceTypeThirdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSkusNestedResourceTypeThirdResult, error) {
			args := v.(LookupSkusNestedResourceTypeThirdArgs)
			r, err := LookupSkusNestedResourceTypeThird(ctx, &args, opts...)
			var s LookupSkusNestedResourceTypeThirdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSkusNestedResourceTypeThirdResultOutput)
}

type LookupSkusNestedResourceTypeThirdOutputArgs struct {
	NestedResourceTypeFirst  pulumi.StringInput `pulumi:"nestedResourceTypeFirst"`
	NestedResourceTypeSecond pulumi.StringInput `pulumi:"nestedResourceTypeSecond"`
	NestedResourceTypeThird  pulumi.StringInput `pulumi:"nestedResourceTypeThird"`
	ProviderNamespace        pulumi.StringInput `pulumi:"providerNamespace"`
	ResourceType             pulumi.StringInput `pulumi:"resourceType"`
	Sku                      pulumi.StringInput `pulumi:"sku"`
}

func (LookupSkusNestedResourceTypeThirdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusNestedResourceTypeThirdArgs)(nil)).Elem()
}

type LookupSkusNestedResourceTypeThirdResultOutput struct{ *pulumi.OutputState }

func (LookupSkusNestedResourceTypeThirdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusNestedResourceTypeThirdResult)(nil)).Elem()
}

func (o LookupSkusNestedResourceTypeThirdResultOutput) ToLookupSkusNestedResourceTypeThirdResultOutput() LookupSkusNestedResourceTypeThirdResultOutput {
	return o
}

func (o LookupSkusNestedResourceTypeThirdResultOutput) ToLookupSkusNestedResourceTypeThirdResultOutputWithContext(ctx context.Context) LookupSkusNestedResourceTypeThirdResultOutput {
	return o
}

func (o LookupSkusNestedResourceTypeThirdResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeThirdResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSkusNestedResourceTypeThirdResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeThirdResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSkusNestedResourceTypeThirdResultOutput) Properties() SkuResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeThirdResult) SkuResourceResponseProperties { return v.Properties }).(SkuResourceResponsePropertiesOutput)
}

func (o LookupSkusNestedResourceTypeThirdResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeThirdResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSkusNestedResourceTypeThirdResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeThirdResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSkusNestedResourceTypeThirdResultOutput{})
}
