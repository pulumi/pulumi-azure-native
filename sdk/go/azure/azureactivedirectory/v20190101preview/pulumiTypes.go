


package v20190101preview

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

func (i B2CResourceSKUArgs) ToB2CResourceSKUPtrOutput() B2CResourceSKUPtrOutput {
	return i.ToB2CResourceSKUPtrOutputWithContext(context.Background())
}

func (i B2CResourceSKUArgs) ToB2CResourceSKUPtrOutputWithContext(ctx context.Context) B2CResourceSKUPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2CResourceSKUOutput).ToB2CResourceSKUPtrOutputWithContext(ctx)
}









type B2CResourceSKUPtrInput interface {
	pulumi.Input

	ToB2CResourceSKUPtrOutput() B2CResourceSKUPtrOutput
	ToB2CResourceSKUPtrOutputWithContext(context.Context) B2CResourceSKUPtrOutput
}

type b2cresourceSKUPtrType B2CResourceSKUArgs

func B2CResourceSKUPtr(v *B2CResourceSKUArgs) B2CResourceSKUPtrInput {
	return (*b2cresourceSKUPtrType)(v)
}

func (*b2cresourceSKUPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CResourceSKU)(nil)).Elem()
}

func (i *b2cresourceSKUPtrType) ToB2CResourceSKUPtrOutput() B2CResourceSKUPtrOutput {
	return i.ToB2CResourceSKUPtrOutputWithContext(context.Background())
}

func (i *b2cresourceSKUPtrType) ToB2CResourceSKUPtrOutputWithContext(ctx context.Context) B2CResourceSKUPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2CResourceSKUPtrOutput)
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

func (o B2CResourceSKUOutput) ToB2CResourceSKUPtrOutput() B2CResourceSKUPtrOutput {
	return o.ToB2CResourceSKUPtrOutputWithContext(context.Background())
}

func (o B2CResourceSKUOutput) ToB2CResourceSKUPtrOutputWithContext(ctx context.Context) B2CResourceSKUPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v B2CResourceSKU) *B2CResourceSKU {
		return &v
	}).(B2CResourceSKUPtrOutput)
}

func (o B2CResourceSKUOutput) Name() B2CResourceSKUNamePtrOutput {
	return o.ApplyT(func(v B2CResourceSKU) *B2CResourceSKUName { return v.Name }).(B2CResourceSKUNamePtrOutput)
}

func (o B2CResourceSKUOutput) Tier() B2CResourceSKUTierPtrOutput {
	return o.ApplyT(func(v B2CResourceSKU) *B2CResourceSKUTier { return v.Tier }).(B2CResourceSKUTierPtrOutput)
}

type B2CResourceSKUPtrOutput struct{ *pulumi.OutputState }

func (B2CResourceSKUPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CResourceSKU)(nil)).Elem()
}

func (o B2CResourceSKUPtrOutput) ToB2CResourceSKUPtrOutput() B2CResourceSKUPtrOutput {
	return o
}

func (o B2CResourceSKUPtrOutput) ToB2CResourceSKUPtrOutputWithContext(ctx context.Context) B2CResourceSKUPtrOutput {
	return o
}

func (o B2CResourceSKUPtrOutput) Elem() B2CResourceSKUOutput {
	return o.ApplyT(func(v *B2CResourceSKU) B2CResourceSKU {
		if v != nil {
			return *v
		}
		var ret B2CResourceSKU
		return ret
	}).(B2CResourceSKUOutput)
}

func (o B2CResourceSKUPtrOutput) Name() B2CResourceSKUNamePtrOutput {
	return o.ApplyT(func(v *B2CResourceSKU) *B2CResourceSKUName {
		if v == nil {
			return nil
		}
		return v.Name
	}).(B2CResourceSKUNamePtrOutput)
}

func (o B2CResourceSKUPtrOutput) Tier() B2CResourceSKUTierPtrOutput {
	return o.ApplyT(func(v *B2CResourceSKU) *B2CResourceSKUTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(B2CResourceSKUTierPtrOutput)
}

type B2CResourceSKUResponse struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type B2CResourceSKUResponseInput interface {
	pulumi.Input

	ToB2CResourceSKUResponseOutput() B2CResourceSKUResponseOutput
	ToB2CResourceSKUResponseOutputWithContext(context.Context) B2CResourceSKUResponseOutput
}

type B2CResourceSKUResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (B2CResourceSKUResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*B2CResourceSKUResponse)(nil)).Elem()
}

func (i B2CResourceSKUResponseArgs) ToB2CResourceSKUResponseOutput() B2CResourceSKUResponseOutput {
	return i.ToB2CResourceSKUResponseOutputWithContext(context.Background())
}

func (i B2CResourceSKUResponseArgs) ToB2CResourceSKUResponseOutputWithContext(ctx context.Context) B2CResourceSKUResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2CResourceSKUResponseOutput)
}

func (i B2CResourceSKUResponseArgs) ToB2CResourceSKUResponsePtrOutput() B2CResourceSKUResponsePtrOutput {
	return i.ToB2CResourceSKUResponsePtrOutputWithContext(context.Background())
}

func (i B2CResourceSKUResponseArgs) ToB2CResourceSKUResponsePtrOutputWithContext(ctx context.Context) B2CResourceSKUResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2CResourceSKUResponseOutput).ToB2CResourceSKUResponsePtrOutputWithContext(ctx)
}









type B2CResourceSKUResponsePtrInput interface {
	pulumi.Input

	ToB2CResourceSKUResponsePtrOutput() B2CResourceSKUResponsePtrOutput
	ToB2CResourceSKUResponsePtrOutputWithContext(context.Context) B2CResourceSKUResponsePtrOutput
}

type b2cresourceSKUResponsePtrType B2CResourceSKUResponseArgs

func B2CResourceSKUResponsePtr(v *B2CResourceSKUResponseArgs) B2CResourceSKUResponsePtrInput {
	return (*b2cresourceSKUResponsePtrType)(v)
}

func (*b2cresourceSKUResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CResourceSKUResponse)(nil)).Elem()
}

func (i *b2cresourceSKUResponsePtrType) ToB2CResourceSKUResponsePtrOutput() B2CResourceSKUResponsePtrOutput {
	return i.ToB2CResourceSKUResponsePtrOutputWithContext(context.Background())
}

func (i *b2cresourceSKUResponsePtrType) ToB2CResourceSKUResponsePtrOutputWithContext(ctx context.Context) B2CResourceSKUResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2CResourceSKUResponsePtrOutput)
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

func (o B2CResourceSKUResponseOutput) ToB2CResourceSKUResponsePtrOutput() B2CResourceSKUResponsePtrOutput {
	return o.ToB2CResourceSKUResponsePtrOutputWithContext(context.Background())
}

func (o B2CResourceSKUResponseOutput) ToB2CResourceSKUResponsePtrOutputWithContext(ctx context.Context) B2CResourceSKUResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v B2CResourceSKUResponse) *B2CResourceSKUResponse {
		return &v
	}).(B2CResourceSKUResponsePtrOutput)
}

func (o B2CResourceSKUResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v B2CResourceSKUResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o B2CResourceSKUResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v B2CResourceSKUResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type B2CResourceSKUResponsePtrOutput struct{ *pulumi.OutputState }

func (B2CResourceSKUResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CResourceSKUResponse)(nil)).Elem()
}

func (o B2CResourceSKUResponsePtrOutput) ToB2CResourceSKUResponsePtrOutput() B2CResourceSKUResponsePtrOutput {
	return o
}

func (o B2CResourceSKUResponsePtrOutput) ToB2CResourceSKUResponsePtrOutputWithContext(ctx context.Context) B2CResourceSKUResponsePtrOutput {
	return o
}

func (o B2CResourceSKUResponsePtrOutput) Elem() B2CResourceSKUResponseOutput {
	return o.ApplyT(func(v *B2CResourceSKUResponse) B2CResourceSKUResponse {
		if v != nil {
			return *v
		}
		var ret B2CResourceSKUResponse
		return ret
	}).(B2CResourceSKUResponseOutput)
}

func (o B2CResourceSKUResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *B2CResourceSKUResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o B2CResourceSKUResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *B2CResourceSKUResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type B2CTenantResourcePropertiesResponseBillingConfig struct {
	BillingType           *string `pulumi:"billingType"`
	EffectiveStartDateUtc string  `pulumi:"effectiveStartDateUtc"`
}





type B2CTenantResourcePropertiesResponseBillingConfigInput interface {
	pulumi.Input

	ToB2CTenantResourcePropertiesResponseBillingConfigOutput() B2CTenantResourcePropertiesResponseBillingConfigOutput
	ToB2CTenantResourcePropertiesResponseBillingConfigOutputWithContext(context.Context) B2CTenantResourcePropertiesResponseBillingConfigOutput
}

type B2CTenantResourcePropertiesResponseBillingConfigArgs struct {
	BillingType           pulumi.StringPtrInput `pulumi:"billingType"`
	EffectiveStartDateUtc pulumi.StringInput    `pulumi:"effectiveStartDateUtc"`
}

func (B2CTenantResourcePropertiesResponseBillingConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*B2CTenantResourcePropertiesResponseBillingConfig)(nil)).Elem()
}

func (i B2CTenantResourcePropertiesResponseBillingConfigArgs) ToB2CTenantResourcePropertiesResponseBillingConfigOutput() B2CTenantResourcePropertiesResponseBillingConfigOutput {
	return i.ToB2CTenantResourcePropertiesResponseBillingConfigOutputWithContext(context.Background())
}

func (i B2CTenantResourcePropertiesResponseBillingConfigArgs) ToB2CTenantResourcePropertiesResponseBillingConfigOutputWithContext(ctx context.Context) B2CTenantResourcePropertiesResponseBillingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2CTenantResourcePropertiesResponseBillingConfigOutput)
}

func (i B2CTenantResourcePropertiesResponseBillingConfigArgs) ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutput() B2CTenantResourcePropertiesResponseBillingConfigPtrOutput {
	return i.ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutputWithContext(context.Background())
}

func (i B2CTenantResourcePropertiesResponseBillingConfigArgs) ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutputWithContext(ctx context.Context) B2CTenantResourcePropertiesResponseBillingConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2CTenantResourcePropertiesResponseBillingConfigOutput).ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutputWithContext(ctx)
}









type B2CTenantResourcePropertiesResponseBillingConfigPtrInput interface {
	pulumi.Input

	ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutput() B2CTenantResourcePropertiesResponseBillingConfigPtrOutput
	ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutputWithContext(context.Context) B2CTenantResourcePropertiesResponseBillingConfigPtrOutput
}

type b2ctenantResourcePropertiesResponseBillingConfigPtrType B2CTenantResourcePropertiesResponseBillingConfigArgs

func B2CTenantResourcePropertiesResponseBillingConfigPtr(v *B2CTenantResourcePropertiesResponseBillingConfigArgs) B2CTenantResourcePropertiesResponseBillingConfigPtrInput {
	return (*b2ctenantResourcePropertiesResponseBillingConfigPtrType)(v)
}

func (*b2ctenantResourcePropertiesResponseBillingConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CTenantResourcePropertiesResponseBillingConfig)(nil)).Elem()
}

func (i *b2ctenantResourcePropertiesResponseBillingConfigPtrType) ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutput() B2CTenantResourcePropertiesResponseBillingConfigPtrOutput {
	return i.ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutputWithContext(context.Background())
}

func (i *b2ctenantResourcePropertiesResponseBillingConfigPtrType) ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutputWithContext(ctx context.Context) B2CTenantResourcePropertiesResponseBillingConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2CTenantResourcePropertiesResponseBillingConfigPtrOutput)
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

func (o B2CTenantResourcePropertiesResponseBillingConfigOutput) ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutput() B2CTenantResourcePropertiesResponseBillingConfigPtrOutput {
	return o.ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutputWithContext(context.Background())
}

func (o B2CTenantResourcePropertiesResponseBillingConfigOutput) ToB2CTenantResourcePropertiesResponseBillingConfigPtrOutputWithContext(ctx context.Context) B2CTenantResourcePropertiesResponseBillingConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v B2CTenantResourcePropertiesResponseBillingConfig) *B2CTenantResourcePropertiesResponseBillingConfig {
		return &v
	}).(B2CTenantResourcePropertiesResponseBillingConfigPtrOutput)
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

func (i CreateTenantRequestBodyPropertiesArgs) ToCreateTenantRequestBodyPropertiesPtrOutput() CreateTenantRequestBodyPropertiesPtrOutput {
	return i.ToCreateTenantRequestBodyPropertiesPtrOutputWithContext(context.Background())
}

func (i CreateTenantRequestBodyPropertiesArgs) ToCreateTenantRequestBodyPropertiesPtrOutputWithContext(ctx context.Context) CreateTenantRequestBodyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateTenantRequestBodyPropertiesOutput).ToCreateTenantRequestBodyPropertiesPtrOutputWithContext(ctx)
}









type CreateTenantRequestBodyPropertiesPtrInput interface {
	pulumi.Input

	ToCreateTenantRequestBodyPropertiesPtrOutput() CreateTenantRequestBodyPropertiesPtrOutput
	ToCreateTenantRequestBodyPropertiesPtrOutputWithContext(context.Context) CreateTenantRequestBodyPropertiesPtrOutput
}

type createTenantRequestBodyPropertiesPtrType CreateTenantRequestBodyPropertiesArgs

func CreateTenantRequestBodyPropertiesPtr(v *CreateTenantRequestBodyPropertiesArgs) CreateTenantRequestBodyPropertiesPtrInput {
	return (*createTenantRequestBodyPropertiesPtrType)(v)
}

func (*createTenantRequestBodyPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateTenantRequestBodyProperties)(nil)).Elem()
}

func (i *createTenantRequestBodyPropertiesPtrType) ToCreateTenantRequestBodyPropertiesPtrOutput() CreateTenantRequestBodyPropertiesPtrOutput {
	return i.ToCreateTenantRequestBodyPropertiesPtrOutputWithContext(context.Background())
}

func (i *createTenantRequestBodyPropertiesPtrType) ToCreateTenantRequestBodyPropertiesPtrOutputWithContext(ctx context.Context) CreateTenantRequestBodyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateTenantRequestBodyPropertiesPtrOutput)
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

func (o CreateTenantRequestBodyPropertiesOutput) ToCreateTenantRequestBodyPropertiesPtrOutput() CreateTenantRequestBodyPropertiesPtrOutput {
	return o.ToCreateTenantRequestBodyPropertiesPtrOutputWithContext(context.Background())
}

func (o CreateTenantRequestBodyPropertiesOutput) ToCreateTenantRequestBodyPropertiesPtrOutputWithContext(ctx context.Context) CreateTenantRequestBodyPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateTenantRequestBodyProperties) *CreateTenantRequestBodyProperties {
		return &v
	}).(CreateTenantRequestBodyPropertiesPtrOutput)
}

func (o CreateTenantRequestBodyPropertiesOutput) CountryCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateTenantRequestBodyProperties) *string { return v.CountryCode }).(pulumi.StringPtrOutput)
}

func (o CreateTenantRequestBodyPropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateTenantRequestBodyProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

type CreateTenantRequestBodyPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CreateTenantRequestBodyPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateTenantRequestBodyProperties)(nil)).Elem()
}

func (o CreateTenantRequestBodyPropertiesPtrOutput) ToCreateTenantRequestBodyPropertiesPtrOutput() CreateTenantRequestBodyPropertiesPtrOutput {
	return o
}

func (o CreateTenantRequestBodyPropertiesPtrOutput) ToCreateTenantRequestBodyPropertiesPtrOutputWithContext(ctx context.Context) CreateTenantRequestBodyPropertiesPtrOutput {
	return o
}

func (o CreateTenantRequestBodyPropertiesPtrOutput) Elem() CreateTenantRequestBodyPropertiesOutput {
	return o.ApplyT(func(v *CreateTenantRequestBodyProperties) CreateTenantRequestBodyProperties {
		if v != nil {
			return *v
		}
		var ret CreateTenantRequestBodyProperties
		return ret
	}).(CreateTenantRequestBodyPropertiesOutput)
}

func (o CreateTenantRequestBodyPropertiesPtrOutput) CountryCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateTenantRequestBodyProperties) *string {
		if v == nil {
			return nil
		}
		return v.CountryCode
	}).(pulumi.StringPtrOutput)
}

func (o CreateTenantRequestBodyPropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateTenantRequestBodyProperties) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*B2CResourceSKUInput)(nil)).Elem(), B2CResourceSKUArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*B2CResourceSKUPtrInput)(nil)).Elem(), B2CResourceSKUArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*B2CResourceSKUResponseInput)(nil)).Elem(), B2CResourceSKUResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*B2CResourceSKUResponsePtrInput)(nil)).Elem(), B2CResourceSKUResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*B2CTenantResourcePropertiesResponseBillingConfigInput)(nil)).Elem(), B2CTenantResourcePropertiesResponseBillingConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*B2CTenantResourcePropertiesResponseBillingConfigPtrInput)(nil)).Elem(), B2CTenantResourcePropertiesResponseBillingConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateTenantRequestBodyPropertiesInput)(nil)).Elem(), CreateTenantRequestBodyPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateTenantRequestBodyPropertiesPtrInput)(nil)).Elem(), CreateTenantRequestBodyPropertiesArgs{})
	pulumi.RegisterOutputType(B2CResourceSKUOutput{})
	pulumi.RegisterOutputType(B2CResourceSKUPtrOutput{})
	pulumi.RegisterOutputType(B2CResourceSKUResponseOutput{})
	pulumi.RegisterOutputType(B2CResourceSKUResponsePtrOutput{})
	pulumi.RegisterOutputType(B2CTenantResourcePropertiesResponseBillingConfigOutput{})
	pulumi.RegisterOutputType(B2CTenantResourcePropertiesResponseBillingConfigPtrOutput{})
	pulumi.RegisterOutputType(CreateTenantRequestBodyPropertiesOutput{})
	pulumi.RegisterOutputType(CreateTenantRequestBodyPropertiesPtrOutput{})
}
