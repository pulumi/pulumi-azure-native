


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSkus(ctx *pulumi.Context, args *LookupSkusArgs, opts ...pulumi.InvokeOption) (*LookupSkusResult, error) {
	var rv LookupSkusResult
	err := ctx.Invoke("azure-native:providerhub/v20210601preview:getSkus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSkusArgs struct {
	ProviderNamespace string `pulumi:"providerNamespace"`
	ResourceType      string `pulumi:"resourceType"`
	Sku               string `pulumi:"sku"`
}

type LookupSkusResult struct {
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties SkuResourceResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Type       string                        `pulumi:"type"`
}

func LookupSkusOutput(ctx *pulumi.Context, args LookupSkusOutputArgs, opts ...pulumi.InvokeOption) LookupSkusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSkusResult, error) {
			args := v.(LookupSkusArgs)
			r, err := LookupSkus(ctx, &args, opts...)
			var s LookupSkusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSkusResultOutput)
}

type LookupSkusOutputArgs struct {
	ProviderNamespace pulumi.StringInput `pulumi:"providerNamespace"`
	ResourceType      pulumi.StringInput `pulumi:"resourceType"`
	Sku               pulumi.StringInput `pulumi:"sku"`
}

func (LookupSkusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusArgs)(nil)).Elem()
}

type LookupSkusResultOutput struct{ *pulumi.OutputState }

func (LookupSkusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusResult)(nil)).Elem()
}

func (o LookupSkusResultOutput) ToLookupSkusResultOutput() LookupSkusResultOutput {
	return o
}

func (o LookupSkusResultOutput) ToLookupSkusResultOutputWithContext(ctx context.Context) LookupSkusResultOutput {
	return o
}

func (o LookupSkusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSkusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSkusResultOutput) Properties() SkuResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupSkusResult) SkuResourceResponseProperties { return v.Properties }).(SkuResourceResponsePropertiesOutput)
}

func (o LookupSkusResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSkusResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSkusResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSkusResultOutput{})
}
