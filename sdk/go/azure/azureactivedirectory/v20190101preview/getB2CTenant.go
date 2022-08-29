


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupB2CTenant(ctx *pulumi.Context, args *LookupB2CTenantArgs, opts ...pulumi.InvokeOption) (*LookupB2CTenantResult, error) {
	var rv LookupB2CTenantResult
	err := ctx.Invoke("azure-native:azureactivedirectory/v20190101preview:getB2CTenant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupB2CTenantArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}

type LookupB2CTenantResult struct {
	BillingConfig *B2CTenantResourcePropertiesResponseBillingConfig `pulumi:"billingConfig"`
	Id            string                                            `pulumi:"id"`
	Location      string                                            `pulumi:"location"`
	Name          string                                            `pulumi:"name"`
	Sku           B2CResourceSKUResponse                            `pulumi:"sku"`
	Tags          map[string]string                                 `pulumi:"tags"`
	TenantId      *string                                           `pulumi:"tenantId"`
	Type          string                                            `pulumi:"type"`
}

func LookupB2CTenantOutput(ctx *pulumi.Context, args LookupB2CTenantOutputArgs, opts ...pulumi.InvokeOption) LookupB2CTenantResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupB2CTenantResult, error) {
			args := v.(LookupB2CTenantArgs)
			r, err := LookupB2CTenant(ctx, &args, opts...)
			var s LookupB2CTenantResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupB2CTenantResultOutput)
}

type LookupB2CTenantOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupB2CTenantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupB2CTenantArgs)(nil)).Elem()
}

type LookupB2CTenantResultOutput struct{ *pulumi.OutputState }

func (LookupB2CTenantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupB2CTenantResult)(nil)).Elem()
}

func (o LookupB2CTenantResultOutput) ToLookupB2CTenantResultOutput() LookupB2CTenantResultOutput {
	return o
}

func (o LookupB2CTenantResultOutput) ToLookupB2CTenantResultOutputWithContext(ctx context.Context) LookupB2CTenantResultOutput {
	return o
}

func (o LookupB2CTenantResultOutput) BillingConfig() B2CTenantResourcePropertiesResponseBillingConfigPtrOutput {
	return o.ApplyT(func(v LookupB2CTenantResult) *B2CTenantResourcePropertiesResponseBillingConfig {
		return v.BillingConfig
	}).(B2CTenantResourcePropertiesResponseBillingConfigPtrOutput)
}

func (o LookupB2CTenantResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupB2CTenantResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupB2CTenantResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupB2CTenantResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupB2CTenantResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupB2CTenantResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupB2CTenantResultOutput) Sku() B2CResourceSKUResponseOutput {
	return o.ApplyT(func(v LookupB2CTenantResult) B2CResourceSKUResponse { return v.Sku }).(B2CResourceSKUResponseOutput)
}

func (o LookupB2CTenantResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupB2CTenantResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupB2CTenantResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupB2CTenantResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupB2CTenantResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupB2CTenantResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupB2CTenantResultOutput{})
}
