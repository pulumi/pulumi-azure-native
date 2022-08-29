


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type B2CResourceSKU struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type B2CResourceSKUInput interface {
	pulumi.Input

	ToB2CResourceSKUOutput() B2CResourceSKUOutput
	ToB2CResourceSKUOutputWithContext(context.Context) B2CResourceSKUOutput
}

type B2CResourceSKUArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (B2CResourceSKUArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*B2CResourceSKU)(nil)).Elem()
}

func (i B2CResourceSKUArgs) ToB2CResourceSKUOutput() B2CResourceSKUOutput {
	return i.ToB2CResourceSKUOutputWithContext(context.Background())
}

func (i B2CResourceSKUArgs) ToB2CResourceSKUOutputWithContext(ctx context.Context) B2CResourceSKUOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2CResourceSKUOutput)
}

type B2CResourceSKUOutput struct{ *pulumi.OutputState }

func (B2CResourceSKUOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*B2CResourceSKU)(nil)).Elem()
}

func (o B2CResourceSKUOutput) ToB2CResourceSKUOutput() B2CResourceSKUOutput {
	return o
}

func (o B2CResourceSKUOutput) ToB2CResourceSKUOutputWithContext(ctx context.Context) B2CResourceSKUOutput {
	return o
}

func (o B2CResourceSKUOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v B2CResourceSKU) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o B2CResourceSKUOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v B2CResourceSKU) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type B2CResourceSKUResponse struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}

type B2CResourceSKUResponseOutput struct{ *pulumi.OutputState }

func (B2CResourceSKUResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*B2CResourceSKUResponse)(nil)).Elem()
}

func (o B2CResourceSKUResponseOutput) ToB2CResourceSKUResponseOutput() B2CResourceSKUResponseOutput {
	return o
}

func (o B2CResourceSKUResponseOutput) ToB2CResourceSKUResponseOutputWithContext(ctx context.Context) B2CResourceSKUResponseOutput {
	return o
}

func (o B2CResourceSKUResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v B2CResourceSKUResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o B2CResourceSKUResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v B2CResourceSKUResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type B2CTenantResourcePropertiesResponseBillingConfig struct {
	BillingType           *string `pulumi:"billingType"`
	EffectiveStartDateUtc string  `pulumi:"effectiveStartDateUtc"`
}

type B2CTenantResourcePropertiesResponseBillingConfigOutput struct{ *pulumi.OutputState }

func (B2CTenantResourcePropertiesResponseBillingConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*B2CTenantResourcePropertiesResponseBillingConfig)(nil)).Elem()
}

func (o B2CTenantResourcePropertiesResponseBillingConfigOutput) ToB2CTenantResourcePropertiesResponseBillingConfigOutput() B2CTenantResourcePropertiesResponseBillingConfigOutput {
	return o
}

func (o B2CTenantResourcePropertiesResponseBillingConfigOutput) ToB2CTenantResourcePropertiesResponseBillingConfigOutputWithContext(ctx context.Context) B2CTenantResourcePropertiesResponseBillingConfigOutput {
	return o
}

func (o B2CTenantResourcePropertiesResponseBillingConfigOutput) BillingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v B2CTenantResourcePropertiesResponseBillingConfig) *string { return v.BillingType }).(pulumi.StringPtrOutput)
}

func (o B2CTenantResourcePropertiesResponseBillingConfigOutput) EffectiveStartDateUtc() pulumi.StringOutput {
	return o.ApplyT(func(v B2CTenantResourcePropertiesResponseBillingConfig) string { return v.EffectiveStartDateUtc }).(pulumi.StringOutput)
}

type B2CTenantResourcePropertiesResponseBillingConfigPtrOutput struct{ *pulumi.OutputState }

func (B2CTenantResourcePropertiesResponseBillingConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CTenantResourcePropertiesResponseBillingConfig)(nil)).Elem()
}

func (o B2CTenantResourcePropertiesResponseBillingConfigPtrOutput) ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutput() B2CTenantResourcePropertiesResponseBillingConfigPtrOutput {
	return o
}

func (o B2CTenantResourcePropertiesResponseBillingConfigPtrOutput) ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutputWithContext(ctx context.Context) B2CTenantResourcePropertiesResponseBillingConfigPtrOutput {
	return o
}

func (o B2CTenantResourcePropertiesResponseBillingConfigPtrOutput) Elem() B2CTenantResourcePropertiesResponseBillingConfigOutput {
	return o.ApplyT(func(v *B2CTenantResourcePropertiesResponseBillingConfig) B2CTenantResourcePropertiesResponseBillingConfig {
		if v != nil {
			return *v
		}
		var ret B2CTenantResourcePropertiesResponseBillingConfig
		return ret
	}).(B2CTenantResourcePropertiesResponseBillingConfigOutput)
}

func (o B2CTenantResourcePropertiesResponseBillingConfigPtrOutput) BillingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *B2CTenantResourcePropertiesResponseBillingConfig) *string {
		if v == nil {
			return nil
		}
		return v.BillingType
	}).(pulumi.StringPtrOutput)
}

func (o B2CTenantResourcePropertiesResponseBillingConfigPtrOutput) EffectiveStartDateUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *B2CTenantResourcePropertiesResponseBillingConfig) *string {
		if v == nil {
			return nil
		}
		return &v.EffectiveStartDateUtc
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(B2CResourceSKUOutput{})
	pulumi.RegisterOutputType(B2CResourceSKUResponseOutput{})
	pulumi.RegisterOutputType(B2CTenantResourcePropertiesResponseBillingConfigOutput{})
	pulumi.RegisterOutputType(B2CTenantResourcePropertiesResponseBillingConfigPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
