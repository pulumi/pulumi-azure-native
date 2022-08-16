


package azureactivedirectory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type B2CResourceSKU struct {
	Name *B2CResourceSKUName `pulumi:"name"`
	Tier *B2CResourceSKUTier `pulumi:"tier"`
}





type B2CResourceSKUInput interface {
	pulumi.Input

	ToB2CResourceSKUOutput() B2CResourceSKUOutput
	ToB2CResourceSKUOutputWithContext(context.Context) B2CResourceSKUOutput
}

type B2CResourceSKUArgs struct {
	Name B2CResourceSKUNamePtrInput `pulumi:"name"`
	Tier B2CResourceSKUTierPtrInput `pulumi:"tier"`
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

func (o B2CResourceSKUOutput) Name() B2CResourceSKUNamePtrOutput {
	return o.ApplyT(func(v B2CResourceSKU) *B2CResourceSKUName { return v.Name }).(B2CResourceSKUNamePtrOutput)
}

func (o B2CResourceSKUOutput) Tier() B2CResourceSKUTierPtrOutput {
	return o.ApplyT(func(v B2CResourceSKU) *B2CResourceSKUTier { return v.Tier }).(B2CResourceSKUTierPtrOutput)
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

type CreateTenantRequestBodyProperties struct {
	CountryCode *string `pulumi:"countryCode"`
	DisplayName *string `pulumi:"displayName"`
}





type CreateTenantRequestBodyPropertiesInput interface {
	pulumi.Input

	ToCreateTenantRequestBodyPropertiesOutput() CreateTenantRequestBodyPropertiesOutput
	ToCreateTenantRequestBodyPropertiesOutputWithContext(context.Context) CreateTenantRequestBodyPropertiesOutput
}

type CreateTenantRequestBodyPropertiesArgs struct {
	CountryCode pulumi.StringPtrInput `pulumi:"countryCode"`
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
}

func (CreateTenantRequestBodyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateTenantRequestBodyProperties)(nil)).Elem()
}

func (i CreateTenantRequestBodyPropertiesArgs) ToCreateTenantRequestBodyPropertiesOutput() CreateTenantRequestBodyPropertiesOutput {
	return i.ToCreateTenantRequestBodyPropertiesOutputWithContext(context.Background())
}

func (i CreateTenantRequestBodyPropertiesArgs) ToCreateTenantRequestBodyPropertiesOutputWithContext(ctx context.Context) CreateTenantRequestBodyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateTenantRequestBodyPropertiesOutput)
}

type CreateTenantRequestBodyPropertiesOutput struct{ *pulumi.OutputState }

func (CreateTenantRequestBodyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateTenantRequestBodyProperties)(nil)).Elem()
}

func (o CreateTenantRequestBodyPropertiesOutput) ToCreateTenantRequestBodyPropertiesOutput() CreateTenantRequestBodyPropertiesOutput {
	return o
}

func (o CreateTenantRequestBodyPropertiesOutput) ToCreateTenantRequestBodyPropertiesOutputWithContext(ctx context.Context) CreateTenantRequestBodyPropertiesOutput {
	return o
}

func (o CreateTenantRequestBodyPropertiesOutput) CountryCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateTenantRequestBodyProperties) *string { return v.CountryCode }).(pulumi.StringPtrOutput)
}

func (o CreateTenantRequestBodyPropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateTenantRequestBodyProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(B2CResourceSKUOutput{})
	pulumi.RegisterOutputType(B2CResourceSKUResponseOutput{})
	pulumi.RegisterOutputType(B2CTenantResourcePropertiesResponseBillingConfigOutput{})
	pulumi.RegisterOutputType(B2CTenantResourcePropertiesResponseBillingConfigPtrOutput{})
	pulumi.RegisterOutputType(CreateTenantRequestBodyPropertiesOutput{})
}
