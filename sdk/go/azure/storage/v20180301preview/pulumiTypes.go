


package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomDomain struct {
	Name             string `pulumi:"name"`
	UseSubDomainName *bool  `pulumi:"useSubDomainName"`
}





type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(context.Context) CustomDomainOutput
}

type CustomDomainArgs struct {
	Name             pulumi.StringInput  `pulumi:"name"`
	UseSubDomainName pulumi.BoolPtrInput `pulumi:"useSubDomainName"`
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil)).Elem()
}

func (i CustomDomainArgs) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i CustomDomainArgs) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

func (i CustomDomainArgs) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i CustomDomainArgs) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput).ToCustomDomainPtrOutputWithContext(ctx)
}









type CustomDomainPtrInput interface {
	pulumi.Input

	ToCustomDomainPtrOutput() CustomDomainPtrOutput
	ToCustomDomainPtrOutputWithContext(context.Context) CustomDomainPtrOutput
}

type customDomainPtrType CustomDomainArgs

func CustomDomainPtr(v *CustomDomainArgs) CustomDomainPtrInput {
	return (*customDomainPtrType)(v)
}

func (*customDomainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (i *customDomainPtrType) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i *customDomainPtrType) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPtrOutput)
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil)).Elem()
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return o.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (o CustomDomainOutput) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDomain) *CustomDomain {
		return &v
	}).(CustomDomainPtrOutput)
}

func (o CustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomain) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomain) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
}

type CustomDomainPtrOutput struct{ *pulumi.OutputState }

func (CustomDomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (o CustomDomainPtrOutput) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return o
}

func (o CustomDomainPtrOutput) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return o
}

func (o CustomDomainPtrOutput) Elem() CustomDomainOutput {
	return o.ApplyT(func(v *CustomDomain) CustomDomain {
		if v != nil {
			return *v
		}
		var ret CustomDomain
		return ret
	}).(CustomDomainOutput)
}

func (o CustomDomainPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomain) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainPtrOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomDomain) *bool {
		if v == nil {
			return nil
		}
		return v.UseSubDomainName
	}).(pulumi.BoolPtrOutput)
}

type CustomDomainResponse struct {
	Name             string `pulumi:"name"`
	UseSubDomainName *bool  `pulumi:"useSubDomainName"`
}





type CustomDomainResponseInput interface {
	pulumi.Input

	ToCustomDomainResponseOutput() CustomDomainResponseOutput
	ToCustomDomainResponseOutputWithContext(context.Context) CustomDomainResponseOutput
}

type CustomDomainResponseArgs struct {
	Name             pulumi.StringInput  `pulumi:"name"`
	UseSubDomainName pulumi.BoolPtrInput `pulumi:"useSubDomainName"`
}

func (CustomDomainResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainResponse)(nil)).Elem()
}

func (i CustomDomainResponseArgs) ToCustomDomainResponseOutput() CustomDomainResponseOutput {
	return i.ToCustomDomainResponseOutputWithContext(context.Background())
}

func (i CustomDomainResponseArgs) ToCustomDomainResponseOutputWithContext(ctx context.Context) CustomDomainResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainResponseOutput)
}

func (i CustomDomainResponseArgs) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return i.ToCustomDomainResponsePtrOutputWithContext(context.Background())
}

func (i CustomDomainResponseArgs) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainResponseOutput).ToCustomDomainResponsePtrOutputWithContext(ctx)
}









type CustomDomainResponsePtrInput interface {
	pulumi.Input

	ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput
	ToCustomDomainResponsePtrOutputWithContext(context.Context) CustomDomainResponsePtrOutput
}

type customDomainResponsePtrType CustomDomainResponseArgs

func CustomDomainResponsePtr(v *CustomDomainResponseArgs) CustomDomainResponsePtrInput {
	return (*customDomainResponsePtrType)(v)
}

func (*customDomainResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainResponse)(nil)).Elem()
}

func (i *customDomainResponsePtrType) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return i.ToCustomDomainResponsePtrOutputWithContext(context.Background())
}

func (i *customDomainResponsePtrType) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainResponsePtrOutput)
}

type CustomDomainResponseOutput struct{ *pulumi.OutputState }

func (CustomDomainResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutput() CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutputWithContext(ctx context.Context) CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return o.ToCustomDomainResponsePtrOutputWithContext(context.Background())
}

func (o CustomDomainResponseOutput) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDomainResponse) *CustomDomainResponse {
		return &v
	}).(CustomDomainResponsePtrOutput)
}

func (o CustomDomainResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomainResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDomainResponseOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
}

type CustomDomainResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomDomainResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponsePtrOutput) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return o
}

func (o CustomDomainResponsePtrOutput) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return o
}

func (o CustomDomainResponsePtrOutput) Elem() CustomDomainResponseOutput {
	return o.ApplyT(func(v *CustomDomainResponse) CustomDomainResponse {
		if v != nil {
			return *v
		}
		var ret CustomDomainResponse
		return ret
	}).(CustomDomainResponseOutput)
}

func (o CustomDomainResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainResponsePtrOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomDomainResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseSubDomainName
	}).(pulumi.BoolPtrOutput)
}

type Encryption struct {
	KeySource          string              `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	Services           *EncryptionServices `pulumi:"services"`
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	KeySource          pulumi.StringInput         `pulumi:"keySource"`
	KeyVaultProperties KeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
	Services           EncryptionServicesPtrInput `pulumi:"services"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) KeySource() pulumi.StringOutput {
	return o.ApplyT(func(v Encryption) string { return v.KeySource }).(pulumi.StringOutput)
}

func (o EncryptionOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v Encryption) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

func (o EncryptionOutput) Services() EncryptionServicesPtrOutput {
	return o.ApplyT(func(v Encryption) *EncryptionServices { return v.Services }).(EncryptionServicesPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return &v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *Encryption) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

func (o EncryptionPtrOutput) Services() EncryptionServicesPtrOutput {
	return o.ApplyT(func(v *Encryption) *EncryptionServices {
		if v == nil {
			return nil
		}
		return v.Services
	}).(EncryptionServicesPtrOutput)
}

type EncryptionResponse struct {
	KeySource          string                      `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Services           *EncryptionServicesResponse `pulumi:"services"`
}





type EncryptionResponseInput interface {
	pulumi.Input

	ToEncryptionResponseOutput() EncryptionResponseOutput
	ToEncryptionResponseOutputWithContext(context.Context) EncryptionResponseOutput
}

type EncryptionResponseArgs struct {
	KeySource          pulumi.StringInput                 `pulumi:"keySource"`
	KeyVaultProperties KeyVaultPropertiesResponsePtrInput `pulumi:"keyVaultProperties"`
	Services           EncryptionServicesResponsePtrInput `pulumi:"services"`
}

func (EncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return i.ToEncryptionResponseOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput)
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput).ToEncryptionResponsePtrOutputWithContext(ctx)
}









type EncryptionResponsePtrInput interface {
	pulumi.Input

	ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput
	ToEncryptionResponsePtrOutputWithContext(context.Context) EncryptionResponsePtrOutput
}

type encryptionResponsePtrType EncryptionResponseArgs

func EncryptionResponsePtr(v *EncryptionResponseArgs) EncryptionResponsePtrInput {
	return (*encryptionResponsePtrType)(v)
}

func (*encryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponsePtrOutput)
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionResponse) *EncryptionResponse {
		return &v
	}).(EncryptionResponsePtrOutput)
}

func (o EncryptionResponseOutput) KeySource() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionResponse) string { return v.KeySource }).(pulumi.StringOutput)
}

func (o EncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionResponseOutput) Services() EncryptionServicesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *EncryptionServicesResponse { return v.Services }).(EncryptionServicesResponsePtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionResponse
		return ret
	}).(EncryptionResponseOutput)
}

func (o EncryptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionResponsePtrOutput) Services() EncryptionServicesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *EncryptionServicesResponse {
		if v == nil {
			return nil
		}
		return v.Services
	}).(EncryptionServicesResponsePtrOutput)
}

type EncryptionService struct {
	Enabled *bool `pulumi:"enabled"`
}





type EncryptionServiceInput interface {
	pulumi.Input

	ToEncryptionServiceOutput() EncryptionServiceOutput
	ToEncryptionServiceOutputWithContext(context.Context) EncryptionServiceOutput
}

type EncryptionServiceArgs struct {
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (EncryptionServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionService)(nil)).Elem()
}

func (i EncryptionServiceArgs) ToEncryptionServiceOutput() EncryptionServiceOutput {
	return i.ToEncryptionServiceOutputWithContext(context.Background())
}

func (i EncryptionServiceArgs) ToEncryptionServiceOutputWithContext(ctx context.Context) EncryptionServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceOutput)
}

func (i EncryptionServiceArgs) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return i.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (i EncryptionServiceArgs) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceOutput).ToEncryptionServicePtrOutputWithContext(ctx)
}









type EncryptionServicePtrInput interface {
	pulumi.Input

	ToEncryptionServicePtrOutput() EncryptionServicePtrOutput
	ToEncryptionServicePtrOutputWithContext(context.Context) EncryptionServicePtrOutput
}

type encryptionServicePtrType EncryptionServiceArgs

func EncryptionServicePtr(v *EncryptionServiceArgs) EncryptionServicePtrInput {
	return (*encryptionServicePtrType)(v)
}

func (*encryptionServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionService)(nil)).Elem()
}

func (i *encryptionServicePtrType) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return i.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (i *encryptionServicePtrType) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicePtrOutput)
}

type EncryptionServiceOutput struct{ *pulumi.OutputState }

func (EncryptionServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionService)(nil)).Elem()
}

func (o EncryptionServiceOutput) ToEncryptionServiceOutput() EncryptionServiceOutput {
	return o
}

func (o EncryptionServiceOutput) ToEncryptionServiceOutputWithContext(ctx context.Context) EncryptionServiceOutput {
	return o
}

func (o EncryptionServiceOutput) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return o.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (o EncryptionServiceOutput) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionService) *EncryptionService {
		return &v
	}).(EncryptionServicePtrOutput)
}

func (o EncryptionServiceOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionService) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type EncryptionServicePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionService)(nil)).Elem()
}

func (o EncryptionServicePtrOutput) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return o
}

func (o EncryptionServicePtrOutput) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return o
}

func (o EncryptionServicePtrOutput) Elem() EncryptionServiceOutput {
	return o.ApplyT(func(v *EncryptionService) EncryptionService {
		if v != nil {
			return *v
		}
		var ret EncryptionService
		return ret
	}).(EncryptionServiceOutput)
}

func (o EncryptionServicePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionService) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type EncryptionServiceResponse struct {
	Enabled         *bool  `pulumi:"enabled"`
	LastEnabledTime string `pulumi:"lastEnabledTime"`
}





type EncryptionServiceResponseInput interface {
	pulumi.Input

	ToEncryptionServiceResponseOutput() EncryptionServiceResponseOutput
	ToEncryptionServiceResponseOutputWithContext(context.Context) EncryptionServiceResponseOutput
}

type EncryptionServiceResponseArgs struct {
	Enabled         pulumi.BoolPtrInput `pulumi:"enabled"`
	LastEnabledTime pulumi.StringInput  `pulumi:"lastEnabledTime"`
}

func (EncryptionServiceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServiceResponse)(nil)).Elem()
}

func (i EncryptionServiceResponseArgs) ToEncryptionServiceResponseOutput() EncryptionServiceResponseOutput {
	return i.ToEncryptionServiceResponseOutputWithContext(context.Background())
}

func (i EncryptionServiceResponseArgs) ToEncryptionServiceResponseOutputWithContext(ctx context.Context) EncryptionServiceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceResponseOutput)
}

func (i EncryptionServiceResponseArgs) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return i.ToEncryptionServiceResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionServiceResponseArgs) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceResponseOutput).ToEncryptionServiceResponsePtrOutputWithContext(ctx)
}









type EncryptionServiceResponsePtrInput interface {
	pulumi.Input

	ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput
	ToEncryptionServiceResponsePtrOutputWithContext(context.Context) EncryptionServiceResponsePtrOutput
}

type encryptionServiceResponsePtrType EncryptionServiceResponseArgs

func EncryptionServiceResponsePtr(v *EncryptionServiceResponseArgs) EncryptionServiceResponsePtrInput {
	return (*encryptionServiceResponsePtrType)(v)
}

func (*encryptionServiceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServiceResponse)(nil)).Elem()
}

func (i *encryptionServiceResponsePtrType) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return i.ToEncryptionServiceResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionServiceResponsePtrType) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceResponsePtrOutput)
}

type EncryptionServiceResponseOutput struct{ *pulumi.OutputState }

func (EncryptionServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServiceResponse)(nil)).Elem()
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponseOutput() EncryptionServiceResponseOutput {
	return o
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponseOutputWithContext(ctx context.Context) EncryptionServiceResponseOutput {
	return o
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return o.ToEncryptionServiceResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionServiceResponse) *EncryptionServiceResponse {
		return &v
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServiceResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o EncryptionServiceResponseOutput) LastEnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) string { return v.LastEnabledTime }).(pulumi.StringOutput)
}

type EncryptionServiceResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServiceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServiceResponse)(nil)).Elem()
}

func (o EncryptionServiceResponsePtrOutput) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return o
}

func (o EncryptionServiceResponsePtrOutput) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return o
}

func (o EncryptionServiceResponsePtrOutput) Elem() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) EncryptionServiceResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionServiceResponse
		return ret
	}).(EncryptionServiceResponseOutput)
}

func (o EncryptionServiceResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o EncryptionServiceResponsePtrOutput) LastEnabledTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastEnabledTime
	}).(pulumi.StringPtrOutput)
}

type EncryptionServices struct {
	Blob *EncryptionService `pulumi:"blob"`
	File *EncryptionService `pulumi:"file"`
}





type EncryptionServicesInput interface {
	pulumi.Input

	ToEncryptionServicesOutput() EncryptionServicesOutput
	ToEncryptionServicesOutputWithContext(context.Context) EncryptionServicesOutput
}

type EncryptionServicesArgs struct {
	Blob EncryptionServicePtrInput `pulumi:"blob"`
	File EncryptionServicePtrInput `pulumi:"file"`
}

func (EncryptionServicesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServices)(nil)).Elem()
}

func (i EncryptionServicesArgs) ToEncryptionServicesOutput() EncryptionServicesOutput {
	return i.ToEncryptionServicesOutputWithContext(context.Background())
}

func (i EncryptionServicesArgs) ToEncryptionServicesOutputWithContext(ctx context.Context) EncryptionServicesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesOutput)
}

func (i EncryptionServicesArgs) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return i.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (i EncryptionServicesArgs) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesOutput).ToEncryptionServicesPtrOutputWithContext(ctx)
}









type EncryptionServicesPtrInput interface {
	pulumi.Input

	ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput
	ToEncryptionServicesPtrOutputWithContext(context.Context) EncryptionServicesPtrOutput
}

type encryptionServicesPtrType EncryptionServicesArgs

func EncryptionServicesPtr(v *EncryptionServicesArgs) EncryptionServicesPtrInput {
	return (*encryptionServicesPtrType)(v)
}

func (*encryptionServicesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServices)(nil)).Elem()
}

func (i *encryptionServicesPtrType) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return i.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (i *encryptionServicesPtrType) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesPtrOutput)
}

type EncryptionServicesOutput struct{ *pulumi.OutputState }

func (EncryptionServicesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServices)(nil)).Elem()
}

func (o EncryptionServicesOutput) ToEncryptionServicesOutput() EncryptionServicesOutput {
	return o
}

func (o EncryptionServicesOutput) ToEncryptionServicesOutputWithContext(ctx context.Context) EncryptionServicesOutput {
	return o
}

func (o EncryptionServicesOutput) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return o.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (o EncryptionServicesOutput) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionServices) *EncryptionServices {
		return &v
	}).(EncryptionServicesPtrOutput)
}

func (o EncryptionServicesOutput) Blob() EncryptionServicePtrOutput {
	return o.ApplyT(func(v EncryptionServices) *EncryptionService { return v.Blob }).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesOutput) File() EncryptionServicePtrOutput {
	return o.ApplyT(func(v EncryptionServices) *EncryptionService { return v.File }).(EncryptionServicePtrOutput)
}

type EncryptionServicesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServices)(nil)).Elem()
}

func (o EncryptionServicesPtrOutput) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return o
}

func (o EncryptionServicesPtrOutput) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return o
}

func (o EncryptionServicesPtrOutput) Elem() EncryptionServicesOutput {
	return o.ApplyT(func(v *EncryptionServices) EncryptionServices {
		if v != nil {
			return *v
		}
		var ret EncryptionServices
		return ret
	}).(EncryptionServicesOutput)
}

func (o EncryptionServicesPtrOutput) Blob() EncryptionServicePtrOutput {
	return o.ApplyT(func(v *EncryptionServices) *EncryptionService {
		if v == nil {
			return nil
		}
		return v.Blob
	}).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesPtrOutput) File() EncryptionServicePtrOutput {
	return o.ApplyT(func(v *EncryptionServices) *EncryptionService {
		if v == nil {
			return nil
		}
		return v.File
	}).(EncryptionServicePtrOutput)
}

type EncryptionServicesResponse struct {
	Blob  *EncryptionServiceResponse `pulumi:"blob"`
	File  *EncryptionServiceResponse `pulumi:"file"`
	Queue EncryptionServiceResponse  `pulumi:"queue"`
	Table EncryptionServiceResponse  `pulumi:"table"`
}





type EncryptionServicesResponseInput interface {
	pulumi.Input

	ToEncryptionServicesResponseOutput() EncryptionServicesResponseOutput
	ToEncryptionServicesResponseOutputWithContext(context.Context) EncryptionServicesResponseOutput
}

type EncryptionServicesResponseArgs struct {
	Blob  EncryptionServiceResponsePtrInput `pulumi:"blob"`
	File  EncryptionServiceResponsePtrInput `pulumi:"file"`
	Queue EncryptionServiceResponseInput    `pulumi:"queue"`
	Table EncryptionServiceResponseInput    `pulumi:"table"`
}

func (EncryptionServicesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServicesResponse)(nil)).Elem()
}

func (i EncryptionServicesResponseArgs) ToEncryptionServicesResponseOutput() EncryptionServicesResponseOutput {
	return i.ToEncryptionServicesResponseOutputWithContext(context.Background())
}

func (i EncryptionServicesResponseArgs) ToEncryptionServicesResponseOutputWithContext(ctx context.Context) EncryptionServicesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesResponseOutput)
}

func (i EncryptionServicesResponseArgs) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return i.ToEncryptionServicesResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionServicesResponseArgs) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesResponseOutput).ToEncryptionServicesResponsePtrOutputWithContext(ctx)
}









type EncryptionServicesResponsePtrInput interface {
	pulumi.Input

	ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput
	ToEncryptionServicesResponsePtrOutputWithContext(context.Context) EncryptionServicesResponsePtrOutput
}

type encryptionServicesResponsePtrType EncryptionServicesResponseArgs

func EncryptionServicesResponsePtr(v *EncryptionServicesResponseArgs) EncryptionServicesResponsePtrInput {
	return (*encryptionServicesResponsePtrType)(v)
}

func (*encryptionServicesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServicesResponse)(nil)).Elem()
}

func (i *encryptionServicesResponsePtrType) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return i.ToEncryptionServicesResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionServicesResponsePtrType) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesResponsePtrOutput)
}

type EncryptionServicesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionServicesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServicesResponse)(nil)).Elem()
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponseOutput() EncryptionServicesResponseOutput {
	return o
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponseOutputWithContext(ctx context.Context) EncryptionServicesResponseOutput {
	return o
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return o.ToEncryptionServicesResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionServicesResponse) *EncryptionServicesResponse {
		return &v
	}).(EncryptionServicesResponsePtrOutput)
}

func (o EncryptionServicesResponseOutput) Blob() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.Blob }).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponseOutput) File() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.File }).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponseOutput) Queue() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) EncryptionServiceResponse { return v.Queue }).(EncryptionServiceResponseOutput)
}

func (o EncryptionServicesResponseOutput) Table() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) EncryptionServiceResponse { return v.Table }).(EncryptionServiceResponseOutput)
}

type EncryptionServicesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServicesResponse)(nil)).Elem()
}

func (o EncryptionServicesResponsePtrOutput) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return o
}

func (o EncryptionServicesResponsePtrOutput) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return o
}

func (o EncryptionServicesResponsePtrOutput) Elem() EncryptionServicesResponseOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) EncryptionServicesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionServicesResponse
		return ret
	}).(EncryptionServicesResponseOutput)
}

func (o EncryptionServicesResponsePtrOutput) Blob() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.Blob
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) File() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.File
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) Queue() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return &v.Queue
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) Table() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return &v.Table
	}).(EncryptionServiceResponsePtrOutput)
}

type EndpointsResponse struct {
	Blob  string `pulumi:"blob"`
	Dfs   string `pulumi:"dfs"`
	File  string `pulumi:"file"`
	Queue string `pulumi:"queue"`
	Table string `pulumi:"table"`
	Web   string `pulumi:"web"`
}





type EndpointsResponseInput interface {
	pulumi.Input

	ToEndpointsResponseOutput() EndpointsResponseOutput
	ToEndpointsResponseOutputWithContext(context.Context) EndpointsResponseOutput
}

type EndpointsResponseArgs struct {
	Blob  pulumi.StringInput `pulumi:"blob"`
	Dfs   pulumi.StringInput `pulumi:"dfs"`
	File  pulumi.StringInput `pulumi:"file"`
	Queue pulumi.StringInput `pulumi:"queue"`
	Table pulumi.StringInput `pulumi:"table"`
	Web   pulumi.StringInput `pulumi:"web"`
}

func (EndpointsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (i EndpointsResponseArgs) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return i.ToEndpointsResponseOutputWithContext(context.Background())
}

func (i EndpointsResponseArgs) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponseOutput)
}

func (i EndpointsResponseArgs) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return i.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i EndpointsResponseArgs) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponseOutput).ToEndpointsResponsePtrOutputWithContext(ctx)
}









type EndpointsResponsePtrInput interface {
	pulumi.Input

	ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput
	ToEndpointsResponsePtrOutputWithContext(context.Context) EndpointsResponsePtrOutput
}

type endpointsResponsePtrType EndpointsResponseArgs

func EndpointsResponsePtr(v *EndpointsResponseArgs) EndpointsResponsePtrInput {
	return (*endpointsResponsePtrType)(v)
}

func (*endpointsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsResponse)(nil)).Elem()
}

func (i *endpointsResponsePtrType) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return i.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i *endpointsResponsePtrType) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsResponsePtrOutput)
}

type EndpointsResponseOutput struct{ *pulumi.OutputState }

func (EndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return o.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (o EndpointsResponseOutput) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointsResponse) *EndpointsResponse {
		return &v
	}).(EndpointsResponsePtrOutput)
}

func (o EndpointsResponseOutput) Blob() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Blob }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Dfs() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Dfs }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.File }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Queue() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Queue }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Table() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Table }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Web() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Web }).(pulumi.StringOutput)
}

type EndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (EndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) Elem() EndpointsResponseOutput {
	return o.ApplyT(func(v *EndpointsResponse) EndpointsResponse {
		if v != nil {
			return *v
		}
		var ret EndpointsResponse
		return ret
	}).(EndpointsResponseOutput)
}

func (o EndpointsResponsePtrOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Blob
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Dfs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Dfs
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.File
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Queue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Queue
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Table
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Web() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Web
	}).(pulumi.StringPtrOutput)
}

type IPRule struct {
	Action           *Action `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}





type IPRuleInput interface {
	pulumi.Input

	ToIPRuleOutput() IPRuleOutput
	ToIPRuleOutputWithContext(context.Context) IPRuleOutput
}

type IPRuleArgs struct {
	Action           ActionPtrInput     `pulumi:"action"`
	IPAddressOrRange pulumi.StringInput `pulumi:"iPAddressOrRange"`
}

func (IPRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (i IPRuleArgs) ToIPRuleOutput() IPRuleOutput {
	return i.ToIPRuleOutputWithContext(context.Background())
}

func (i IPRuleArgs) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleOutput)
}





type IPRuleArrayInput interface {
	pulumi.Input

	ToIPRuleArrayOutput() IPRuleArrayOutput
	ToIPRuleArrayOutputWithContext(context.Context) IPRuleArrayOutput
}

type IPRuleArray []IPRuleInput

func (IPRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (i IPRuleArray) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return i.ToIPRuleArrayOutputWithContext(context.Background())
}

func (i IPRuleArray) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleArrayOutput)
}

type IPRuleOutput struct{ *pulumi.OutputState }

func (IPRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (o IPRuleOutput) ToIPRuleOutput() IPRuleOutput {
	return o
}

func (o IPRuleOutput) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return o
}

func (o IPRuleOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v IPRule) *Action { return v.Action }).(ActionPtrOutput)
}

func (o IPRuleOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRule) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleArrayOutput struct{ *pulumi.OutputState }

func (IPRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) Index(i pulumi.IntInput) IPRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRule {
		return vs[0].([]IPRule)[vs[1].(int)]
	}).(IPRuleOutput)
}

type IPRuleResponse struct {
	Action           *string `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}





type IPRuleResponseInput interface {
	pulumi.Input

	ToIPRuleResponseOutput() IPRuleResponseOutput
	ToIPRuleResponseOutputWithContext(context.Context) IPRuleResponseOutput
}

type IPRuleResponseArgs struct {
	Action           pulumi.StringPtrInput `pulumi:"action"`
	IPAddressOrRange pulumi.StringInput    `pulumi:"iPAddressOrRange"`
}

func (IPRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRuleResponse)(nil)).Elem()
}

func (i IPRuleResponseArgs) ToIPRuleResponseOutput() IPRuleResponseOutput {
	return i.ToIPRuleResponseOutputWithContext(context.Background())
}

func (i IPRuleResponseArgs) ToIPRuleResponseOutputWithContext(ctx context.Context) IPRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleResponseOutput)
}





type IPRuleResponseArrayInput interface {
	pulumi.Input

	ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput
	ToIPRuleResponseArrayOutputWithContext(context.Context) IPRuleResponseArrayOutput
}

type IPRuleResponseArray []IPRuleResponseInput

func (IPRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRuleResponse)(nil)).Elem()
}

func (i IPRuleResponseArray) ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput {
	return i.ToIPRuleResponseArrayOutputWithContext(context.Background())
}

func (i IPRuleResponseArray) ToIPRuleResponseArrayOutputWithContext(ctx context.Context) IPRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleResponseArrayOutput)
}

type IPRuleResponseOutput struct{ *pulumi.OutputState }

func (IPRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutput() IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutputWithContext(ctx context.Context) IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IPRuleResponseOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRuleResponse) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IPRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutputWithContext(ctx context.Context) IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) Index(i pulumi.IntInput) IPRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRuleResponse {
		return vs[0].([]IPRuleResponse)[vs[1].(int)]
	}).(IPRuleResponseOutput)
}

type Identity struct {
	Type IdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type IdentityTypeInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() IdentityTypeOutput {
	return o.ApplyT(func(v Identity) IdentityType { return v.Type }).(IdentityTypeOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() IdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *IdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(IdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ImmutabilityPolicyPropertiesResponse struct {
	Etag                                  string                          `pulumi:"etag"`
	ImmutabilityPeriodSinceCreationInDays int                             `pulumi:"immutabilityPeriodSinceCreationInDays"`
	State                                 string                          `pulumi:"state"`
	UpdateHistory                         []UpdateHistoryPropertyResponse `pulumi:"updateHistory"`
}





type ImmutabilityPolicyPropertiesResponseInput interface {
	pulumi.Input

	ToImmutabilityPolicyPropertiesResponseOutput() ImmutabilityPolicyPropertiesResponseOutput
	ToImmutabilityPolicyPropertiesResponseOutputWithContext(context.Context) ImmutabilityPolicyPropertiesResponseOutput
}

type ImmutabilityPolicyPropertiesResponseArgs struct {
	Etag                                  pulumi.StringInput                      `pulumi:"etag"`
	ImmutabilityPeriodSinceCreationInDays pulumi.IntInput                         `pulumi:"immutabilityPeriodSinceCreationInDays"`
	State                                 pulumi.StringInput                      `pulumi:"state"`
	UpdateHistory                         UpdateHistoryPropertyResponseArrayInput `pulumi:"updateHistory"`
}

func (ImmutabilityPolicyPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImmutabilityPolicyPropertiesResponse)(nil)).Elem()
}

func (i ImmutabilityPolicyPropertiesResponseArgs) ToImmutabilityPolicyPropertiesResponseOutput() ImmutabilityPolicyPropertiesResponseOutput {
	return i.ToImmutabilityPolicyPropertiesResponseOutputWithContext(context.Background())
}

func (i ImmutabilityPolicyPropertiesResponseArgs) ToImmutabilityPolicyPropertiesResponseOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilityPolicyPropertiesResponseOutput)
}

func (i ImmutabilityPolicyPropertiesResponseArgs) ToImmutabilityPolicyPropertiesResponsePtrOutput() ImmutabilityPolicyPropertiesResponsePtrOutput {
	return i.ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ImmutabilityPolicyPropertiesResponseArgs) ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilityPolicyPropertiesResponseOutput).ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(ctx)
}









type ImmutabilityPolicyPropertiesResponsePtrInput interface {
	pulumi.Input

	ToImmutabilityPolicyPropertiesResponsePtrOutput() ImmutabilityPolicyPropertiesResponsePtrOutput
	ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(context.Context) ImmutabilityPolicyPropertiesResponsePtrOutput
}

type immutabilityPolicyPropertiesResponsePtrType ImmutabilityPolicyPropertiesResponseArgs

func ImmutabilityPolicyPropertiesResponsePtr(v *ImmutabilityPolicyPropertiesResponseArgs) ImmutabilityPolicyPropertiesResponsePtrInput {
	return (*immutabilityPolicyPropertiesResponsePtrType)(v)
}

func (*immutabilityPolicyPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutabilityPolicyPropertiesResponse)(nil)).Elem()
}

func (i *immutabilityPolicyPropertiesResponsePtrType) ToImmutabilityPolicyPropertiesResponsePtrOutput() ImmutabilityPolicyPropertiesResponsePtrOutput {
	return i.ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *immutabilityPolicyPropertiesResponsePtrType) ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilityPolicyPropertiesResponsePtrOutput)
}

type ImmutabilityPolicyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ImmutabilityPolicyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImmutabilityPolicyPropertiesResponse)(nil)).Elem()
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ToImmutabilityPolicyPropertiesResponseOutput() ImmutabilityPolicyPropertiesResponseOutput {
	return o
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ToImmutabilityPolicyPropertiesResponseOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponseOutput {
	return o
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ToImmutabilityPolicyPropertiesResponsePtrOutput() ImmutabilityPolicyPropertiesResponsePtrOutput {
	return o.ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImmutabilityPolicyPropertiesResponse) *ImmutabilityPolicyPropertiesResponse {
		return &v
	}).(ImmutabilityPolicyPropertiesResponsePtrOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) ImmutabilityPeriodSinceCreationInDays() pulumi.IntOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) int { return v.ImmutabilityPeriodSinceCreationInDays }).(pulumi.IntOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o ImmutabilityPolicyPropertiesResponseOutput) UpdateHistory() UpdateHistoryPropertyResponseArrayOutput {
	return o.ApplyT(func(v ImmutabilityPolicyPropertiesResponse) []UpdateHistoryPropertyResponse { return v.UpdateHistory }).(UpdateHistoryPropertyResponseArrayOutput)
}

type ImmutabilityPolicyPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ImmutabilityPolicyPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutabilityPolicyPropertiesResponse)(nil)).Elem()
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) ToImmutabilityPolicyPropertiesResponsePtrOutput() ImmutabilityPolicyPropertiesResponsePtrOutput {
	return o
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) ToImmutabilityPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) ImmutabilityPolicyPropertiesResponsePtrOutput {
	return o
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) Elem() ImmutabilityPolicyPropertiesResponseOutput {
	return o.ApplyT(func(v *ImmutabilityPolicyPropertiesResponse) ImmutabilityPolicyPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ImmutabilityPolicyPropertiesResponse
		return ret
	}).(ImmutabilityPolicyPropertiesResponseOutput)
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutabilityPolicyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) ImmutabilityPeriodSinceCreationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImmutabilityPolicyPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ImmutabilityPeriodSinceCreationInDays
	}).(pulumi.IntPtrOutput)
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutabilityPolicyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

func (o ImmutabilityPolicyPropertiesResponsePtrOutput) UpdateHistory() UpdateHistoryPropertyResponseArrayOutput {
	return o.ApplyT(func(v *ImmutabilityPolicyPropertiesResponse) []UpdateHistoryPropertyResponse {
		if v == nil {
			return nil
		}
		return v.UpdateHistory
	}).(UpdateHistoryPropertyResponseArrayOutput)
}

type KeyVaultProperties struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultProperties
		return ret
	}).(KeyVaultPropertiesOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}





type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

type KeyVaultPropertiesResponseArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
}

func (KeyVaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return i.ToKeyVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput)
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput).ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx)
}









type KeyVaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput
	ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Context) KeyVaultPropertiesResponsePtrOutput
}

type keyVaultPropertiesResponsePtrType KeyVaultPropertiesResponseArgs

func KeyVaultPropertiesResponsePtr(v *KeyVaultPropertiesResponseArgs) KeyVaultPropertiesResponsePtrInput {
	return (*keyVaultPropertiesResponsePtrType)(v)
}

func (*keyVaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponsePtrOutput)
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultPropertiesResponse) *KeyVaultPropertiesResponse {
		return &v
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultPropertiesResponse
		return ret
	}).(KeyVaultPropertiesResponseOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type LegalHoldPropertiesResponse struct {
	HasLegalHold bool                  `pulumi:"hasLegalHold"`
	Tags         []TagPropertyResponse `pulumi:"tags"`
}





type LegalHoldPropertiesResponseInput interface {
	pulumi.Input

	ToLegalHoldPropertiesResponseOutput() LegalHoldPropertiesResponseOutput
	ToLegalHoldPropertiesResponseOutputWithContext(context.Context) LegalHoldPropertiesResponseOutput
}

type LegalHoldPropertiesResponseArgs struct {
	HasLegalHold pulumi.BoolInput              `pulumi:"hasLegalHold"`
	Tags         TagPropertyResponseArrayInput `pulumi:"tags"`
}

func (LegalHoldPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LegalHoldPropertiesResponse)(nil)).Elem()
}

func (i LegalHoldPropertiesResponseArgs) ToLegalHoldPropertiesResponseOutput() LegalHoldPropertiesResponseOutput {
	return i.ToLegalHoldPropertiesResponseOutputWithContext(context.Background())
}

func (i LegalHoldPropertiesResponseArgs) ToLegalHoldPropertiesResponseOutputWithContext(ctx context.Context) LegalHoldPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegalHoldPropertiesResponseOutput)
}

func (i LegalHoldPropertiesResponseArgs) ToLegalHoldPropertiesResponsePtrOutput() LegalHoldPropertiesResponsePtrOutput {
	return i.ToLegalHoldPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i LegalHoldPropertiesResponseArgs) ToLegalHoldPropertiesResponsePtrOutputWithContext(ctx context.Context) LegalHoldPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegalHoldPropertiesResponseOutput).ToLegalHoldPropertiesResponsePtrOutputWithContext(ctx)
}









type LegalHoldPropertiesResponsePtrInput interface {
	pulumi.Input

	ToLegalHoldPropertiesResponsePtrOutput() LegalHoldPropertiesResponsePtrOutput
	ToLegalHoldPropertiesResponsePtrOutputWithContext(context.Context) LegalHoldPropertiesResponsePtrOutput
}

type legalHoldPropertiesResponsePtrType LegalHoldPropertiesResponseArgs

func LegalHoldPropertiesResponsePtr(v *LegalHoldPropertiesResponseArgs) LegalHoldPropertiesResponsePtrInput {
	return (*legalHoldPropertiesResponsePtrType)(v)
}

func (*legalHoldPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LegalHoldPropertiesResponse)(nil)).Elem()
}

func (i *legalHoldPropertiesResponsePtrType) ToLegalHoldPropertiesResponsePtrOutput() LegalHoldPropertiesResponsePtrOutput {
	return i.ToLegalHoldPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *legalHoldPropertiesResponsePtrType) ToLegalHoldPropertiesResponsePtrOutputWithContext(ctx context.Context) LegalHoldPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegalHoldPropertiesResponsePtrOutput)
}

type LegalHoldPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LegalHoldPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LegalHoldPropertiesResponse)(nil)).Elem()
}

func (o LegalHoldPropertiesResponseOutput) ToLegalHoldPropertiesResponseOutput() LegalHoldPropertiesResponseOutput {
	return o
}

func (o LegalHoldPropertiesResponseOutput) ToLegalHoldPropertiesResponseOutputWithContext(ctx context.Context) LegalHoldPropertiesResponseOutput {
	return o
}

func (o LegalHoldPropertiesResponseOutput) ToLegalHoldPropertiesResponsePtrOutput() LegalHoldPropertiesResponsePtrOutput {
	return o.ToLegalHoldPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o LegalHoldPropertiesResponseOutput) ToLegalHoldPropertiesResponsePtrOutputWithContext(ctx context.Context) LegalHoldPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LegalHoldPropertiesResponse) *LegalHoldPropertiesResponse {
		return &v
	}).(LegalHoldPropertiesResponsePtrOutput)
}

func (o LegalHoldPropertiesResponseOutput) HasLegalHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LegalHoldPropertiesResponse) bool { return v.HasLegalHold }).(pulumi.BoolOutput)
}

func (o LegalHoldPropertiesResponseOutput) Tags() TagPropertyResponseArrayOutput {
	return o.ApplyT(func(v LegalHoldPropertiesResponse) []TagPropertyResponse { return v.Tags }).(TagPropertyResponseArrayOutput)
}

type LegalHoldPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LegalHoldPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LegalHoldPropertiesResponse)(nil)).Elem()
}

func (o LegalHoldPropertiesResponsePtrOutput) ToLegalHoldPropertiesResponsePtrOutput() LegalHoldPropertiesResponsePtrOutput {
	return o
}

func (o LegalHoldPropertiesResponsePtrOutput) ToLegalHoldPropertiesResponsePtrOutputWithContext(ctx context.Context) LegalHoldPropertiesResponsePtrOutput {
	return o
}

func (o LegalHoldPropertiesResponsePtrOutput) Elem() LegalHoldPropertiesResponseOutput {
	return o.ApplyT(func(v *LegalHoldPropertiesResponse) LegalHoldPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LegalHoldPropertiesResponse
		return ret
	}).(LegalHoldPropertiesResponseOutput)
}

func (o LegalHoldPropertiesResponsePtrOutput) HasLegalHold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LegalHoldPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.HasLegalHold
	}).(pulumi.BoolPtrOutput)
}

func (o LegalHoldPropertiesResponsePtrOutput) Tags() TagPropertyResponseArrayOutput {
	return o.ApplyT(func(v *LegalHoldPropertiesResponse) []TagPropertyResponse {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(TagPropertyResponseArrayOutput)
}

type NetworkRuleSet struct {
	Bypass              *string              `pulumi:"bypass"`
	DefaultAction       DefaultAction        `pulumi:"defaultAction"`
	IpRules             []IPRule             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	Bypass              pulumi.StringPtrInput        `pulumi:"bypass"`
	DefaultAction       DefaultActionInput           `pulumi:"defaultAction"`
	IpRules             IPRuleArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSet) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetOutput) DefaultAction() DefaultActionOutput {
	return o.ApplyT(func(v NetworkRuleSet) DefaultAction { return v.DefaultAction }).(DefaultActionOutput)
}

func (o NetworkRuleSetOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []IPRule { return v.IpRules }).(IPRuleArrayOutput)
}

func (o NetworkRuleSetOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []VirtualNetworkRule { return v.VirtualNetworkRules }).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return v.Bypass
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPtrOutput) DefaultAction() DefaultActionPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *DefaultAction {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(DefaultActionPtrOutput)
}

func (o NetworkRuleSetPtrOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []IPRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleArrayOutput)
}

func (o NetworkRuleSetPtrOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []VirtualNetworkRule {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	Bypass              *string                      `pulumi:"bypass"`
	DefaultAction       string                       `pulumi:"defaultAction"`
	IpRules             []IPRuleResponse             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}





type NetworkRuleSetResponseInput interface {
	pulumi.Input

	ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput
	ToNetworkRuleSetResponseOutputWithContext(context.Context) NetworkRuleSetResponseOutput
}

type NetworkRuleSetResponseArgs struct {
	Bypass              pulumi.StringPtrInput                `pulumi:"bypass"`
	DefaultAction       pulumi.StringInput                   `pulumi:"defaultAction"`
	IpRules             IPRuleResponseArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleResponseArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return i.ToNetworkRuleSetResponseOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput)
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput).ToNetworkRuleSetResponsePtrOutputWithContext(ctx)
}









type NetworkRuleSetResponsePtrInput interface {
	pulumi.Input

	ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput
	ToNetworkRuleSetResponsePtrOutputWithContext(context.Context) NetworkRuleSetResponsePtrOutput
}

type networkRuleSetResponsePtrType NetworkRuleSetResponseArgs

func NetworkRuleSetResponsePtr(v *NetworkRuleSetResponseArgs) NetworkRuleSetResponsePtrInput {
	return (*networkRuleSetResponsePtrType)(v)
}

func (*networkRuleSetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponsePtrOutput)
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSetResponse) *NetworkRuleSetResponse {
		return &v
	}).(NetworkRuleSetResponsePtrOutput)
}

func (o NetworkRuleSetResponseOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponseOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) string { return v.DefaultAction }).(pulumi.StringOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IPRuleResponse { return v.IpRules }).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponseOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

type NetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) Elem() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) NetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetResponse
		return ret
	}).(NetworkRuleSetResponseOutput)
}

func (o NetworkRuleSetResponsePtrOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Bypass
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []IPRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponsePtrOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []VirtualNetworkRuleResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleResponseArrayOutput)
}

type Restriction struct {
	ReasonCode *string `pulumi:"reasonCode"`
}





type RestrictionInput interface {
	pulumi.Input

	ToRestrictionOutput() RestrictionOutput
	ToRestrictionOutputWithContext(context.Context) RestrictionOutput
}

type RestrictionArgs struct {
	ReasonCode pulumi.StringPtrInput `pulumi:"reasonCode"`
}

func (RestrictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Restriction)(nil)).Elem()
}

func (i RestrictionArgs) ToRestrictionOutput() RestrictionOutput {
	return i.ToRestrictionOutputWithContext(context.Background())
}

func (i RestrictionArgs) ToRestrictionOutputWithContext(ctx context.Context) RestrictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestrictionOutput)
}





type RestrictionArrayInput interface {
	pulumi.Input

	ToRestrictionArrayOutput() RestrictionArrayOutput
	ToRestrictionArrayOutputWithContext(context.Context) RestrictionArrayOutput
}

type RestrictionArray []RestrictionInput

func (RestrictionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Restriction)(nil)).Elem()
}

func (i RestrictionArray) ToRestrictionArrayOutput() RestrictionArrayOutput {
	return i.ToRestrictionArrayOutputWithContext(context.Background())
}

func (i RestrictionArray) ToRestrictionArrayOutputWithContext(ctx context.Context) RestrictionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestrictionArrayOutput)
}

type RestrictionOutput struct{ *pulumi.OutputState }

func (RestrictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Restriction)(nil)).Elem()
}

func (o RestrictionOutput) ToRestrictionOutput() RestrictionOutput {
	return o
}

func (o RestrictionOutput) ToRestrictionOutputWithContext(ctx context.Context) RestrictionOutput {
	return o
}

func (o RestrictionOutput) ReasonCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Restriction) *string { return v.ReasonCode }).(pulumi.StringPtrOutput)
}

type RestrictionArrayOutput struct{ *pulumi.OutputState }

func (RestrictionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Restriction)(nil)).Elem()
}

func (o RestrictionArrayOutput) ToRestrictionArrayOutput() RestrictionArrayOutput {
	return o
}

func (o RestrictionArrayOutput) ToRestrictionArrayOutputWithContext(ctx context.Context) RestrictionArrayOutput {
	return o
}

func (o RestrictionArrayOutput) Index(i pulumi.IntInput) RestrictionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Restriction {
		return vs[0].([]Restriction)[vs[1].(int)]
	}).(RestrictionOutput)
}

type RestrictionResponse struct {
	ReasonCode *string  `pulumi:"reasonCode"`
	Type       string   `pulumi:"type"`
	Values     []string `pulumi:"values"`
}





type RestrictionResponseInput interface {
	pulumi.Input

	ToRestrictionResponseOutput() RestrictionResponseOutput
	ToRestrictionResponseOutputWithContext(context.Context) RestrictionResponseOutput
}

type RestrictionResponseArgs struct {
	ReasonCode pulumi.StringPtrInput   `pulumi:"reasonCode"`
	Type       pulumi.StringInput      `pulumi:"type"`
	Values     pulumi.StringArrayInput `pulumi:"values"`
}

func (RestrictionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RestrictionResponse)(nil)).Elem()
}

func (i RestrictionResponseArgs) ToRestrictionResponseOutput() RestrictionResponseOutput {
	return i.ToRestrictionResponseOutputWithContext(context.Background())
}

func (i RestrictionResponseArgs) ToRestrictionResponseOutputWithContext(ctx context.Context) RestrictionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestrictionResponseOutput)
}





type RestrictionResponseArrayInput interface {
	pulumi.Input

	ToRestrictionResponseArrayOutput() RestrictionResponseArrayOutput
	ToRestrictionResponseArrayOutputWithContext(context.Context) RestrictionResponseArrayOutput
}

type RestrictionResponseArray []RestrictionResponseInput

func (RestrictionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RestrictionResponse)(nil)).Elem()
}

func (i RestrictionResponseArray) ToRestrictionResponseArrayOutput() RestrictionResponseArrayOutput {
	return i.ToRestrictionResponseArrayOutputWithContext(context.Background())
}

func (i RestrictionResponseArray) ToRestrictionResponseArrayOutputWithContext(ctx context.Context) RestrictionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestrictionResponseArrayOutput)
}

type RestrictionResponseOutput struct{ *pulumi.OutputState }

func (RestrictionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestrictionResponse)(nil)).Elem()
}

func (o RestrictionResponseOutput) ToRestrictionResponseOutput() RestrictionResponseOutput {
	return o
}

func (o RestrictionResponseOutput) ToRestrictionResponseOutputWithContext(ctx context.Context) RestrictionResponseOutput {
	return o
}

func (o RestrictionResponseOutput) ReasonCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestrictionResponse) *string { return v.ReasonCode }).(pulumi.StringPtrOutput)
}

func (o RestrictionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RestrictionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o RestrictionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RestrictionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type RestrictionResponseArrayOutput struct{ *pulumi.OutputState }

func (RestrictionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RestrictionResponse)(nil)).Elem()
}

func (o RestrictionResponseArrayOutput) ToRestrictionResponseArrayOutput() RestrictionResponseArrayOutput {
	return o
}

func (o RestrictionResponseArrayOutput) ToRestrictionResponseArrayOutputWithContext(ctx context.Context) RestrictionResponseArrayOutput {
	return o
}

func (o RestrictionResponseArrayOutput) Index(i pulumi.IntInput) RestrictionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RestrictionResponse {
		return vs[0].([]RestrictionResponse)[vs[1].(int)]
	}).(RestrictionResponseOutput)
}

type SKUCapabilityResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type SKUCapabilityResponseInput interface {
	pulumi.Input

	ToSKUCapabilityResponseOutput() SKUCapabilityResponseOutput
	ToSKUCapabilityResponseOutputWithContext(context.Context) SKUCapabilityResponseOutput
}

type SKUCapabilityResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (SKUCapabilityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SKUCapabilityResponse)(nil)).Elem()
}

func (i SKUCapabilityResponseArgs) ToSKUCapabilityResponseOutput() SKUCapabilityResponseOutput {
	return i.ToSKUCapabilityResponseOutputWithContext(context.Background())
}

func (i SKUCapabilityResponseArgs) ToSKUCapabilityResponseOutputWithContext(ctx context.Context) SKUCapabilityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SKUCapabilityResponseOutput)
}





type SKUCapabilityResponseArrayInput interface {
	pulumi.Input

	ToSKUCapabilityResponseArrayOutput() SKUCapabilityResponseArrayOutput
	ToSKUCapabilityResponseArrayOutputWithContext(context.Context) SKUCapabilityResponseArrayOutput
}

type SKUCapabilityResponseArray []SKUCapabilityResponseInput

func (SKUCapabilityResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SKUCapabilityResponse)(nil)).Elem()
}

func (i SKUCapabilityResponseArray) ToSKUCapabilityResponseArrayOutput() SKUCapabilityResponseArrayOutput {
	return i.ToSKUCapabilityResponseArrayOutputWithContext(context.Background())
}

func (i SKUCapabilityResponseArray) ToSKUCapabilityResponseArrayOutputWithContext(ctx context.Context) SKUCapabilityResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SKUCapabilityResponseArrayOutput)
}

type SKUCapabilityResponseOutput struct{ *pulumi.OutputState }

func (SKUCapabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SKUCapabilityResponse)(nil)).Elem()
}

func (o SKUCapabilityResponseOutput) ToSKUCapabilityResponseOutput() SKUCapabilityResponseOutput {
	return o
}

func (o SKUCapabilityResponseOutput) ToSKUCapabilityResponseOutputWithContext(ctx context.Context) SKUCapabilityResponseOutput {
	return o
}

func (o SKUCapabilityResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SKUCapabilityResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SKUCapabilityResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SKUCapabilityResponse) string { return v.Value }).(pulumi.StringOutput)
}

type SKUCapabilityResponseArrayOutput struct{ *pulumi.OutputState }

func (SKUCapabilityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SKUCapabilityResponse)(nil)).Elem()
}

func (o SKUCapabilityResponseArrayOutput) ToSKUCapabilityResponseArrayOutput() SKUCapabilityResponseArrayOutput {
	return o
}

func (o SKUCapabilityResponseArrayOutput) ToSKUCapabilityResponseArrayOutputWithContext(ctx context.Context) SKUCapabilityResponseArrayOutput {
	return o
}

func (o SKUCapabilityResponseArrayOutput) Index(i pulumi.IntInput) SKUCapabilityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SKUCapabilityResponse {
		return vs[0].([]SKUCapabilityResponse)[vs[1].(int)]
	}).(SKUCapabilityResponseOutput)
}

type Sku struct {
	Name         SkuName       `pulumi:"name"`
	Restrictions []Restriction `pulumi:"restrictions"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name         SkuNameInput          `pulumi:"name"`
	Restrictions RestrictionArrayInput `pulumi:"restrictions"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() SkuNameOutput {
	return o.ApplyT(func(v Sku) SkuName { return v.Name }).(SkuNameOutput)
}

func (o SkuOutput) Restrictions() RestrictionArrayOutput {
	return o.ApplyT(func(v Sku) []Restriction { return v.Restrictions }).(RestrictionArrayOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v *Sku) *SkuName {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(SkuNamePtrOutput)
}

func (o SkuPtrOutput) Restrictions() RestrictionArrayOutput {
	return o.ApplyT(func(v *Sku) []Restriction {
		if v == nil {
			return nil
		}
		return v.Restrictions
	}).(RestrictionArrayOutput)
}

type SkuResponse struct {
	Capabilities []SKUCapabilityResponse `pulumi:"capabilities"`
	Kind         string                  `pulumi:"kind"`
	Locations    []string                `pulumi:"locations"`
	Name         string                  `pulumi:"name"`
	ResourceType string                  `pulumi:"resourceType"`
	Restrictions []RestrictionResponse   `pulumi:"restrictions"`
	Tier         string                  `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capabilities SKUCapabilityResponseArrayInput `pulumi:"capabilities"`
	Kind         pulumi.StringInput              `pulumi:"kind"`
	Locations    pulumi.StringArrayInput         `pulumi:"locations"`
	Name         pulumi.StringInput              `pulumi:"name"`
	ResourceType pulumi.StringInput              `pulumi:"resourceType"`
	Restrictions RestrictionResponseArrayInput   `pulumi:"restrictions"`
	Tier         pulumi.StringInput              `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Capabilities() SKUCapabilityResponseArrayOutput {
	return o.ApplyT(func(v SkuResponse) []SKUCapabilityResponse { return v.Capabilities }).(SKUCapabilityResponseArrayOutput)
}

func (o SkuResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SkuResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Restrictions() RestrictionResponseArrayOutput {
	return o.ApplyT(func(v SkuResponse) []RestrictionResponse { return v.Restrictions }).(RestrictionResponseArrayOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capabilities() SKUCapabilityResponseArrayOutput {
	return o.ApplyT(func(v *SkuResponse) []SKUCapabilityResponse {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(SKUCapabilityResponseArrayOutput)
}

func (o SkuResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SkuResponse) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceType
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Restrictions() RestrictionResponseArrayOutput {
	return o.ApplyT(func(v *SkuResponse) []RestrictionResponse {
		if v == nil {
			return nil
		}
		return v.Restrictions
	}).(RestrictionResponseArrayOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type StorageAccountKeyResponse struct {
	KeyName     string `pulumi:"keyName"`
	Permissions string `pulumi:"permissions"`
	Value       string `pulumi:"value"`
}





type StorageAccountKeyResponseInput interface {
	pulumi.Input

	ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput
	ToStorageAccountKeyResponseOutputWithContext(context.Context) StorageAccountKeyResponseOutput
}

type StorageAccountKeyResponseArgs struct {
	KeyName     pulumi.StringInput `pulumi:"keyName"`
	Permissions pulumi.StringInput `pulumi:"permissions"`
	Value       pulumi.StringInput `pulumi:"value"`
}

func (StorageAccountKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountKeyResponse)(nil)).Elem()
}

func (i StorageAccountKeyResponseArgs) ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput {
	return i.ToStorageAccountKeyResponseOutputWithContext(context.Background())
}

func (i StorageAccountKeyResponseArgs) ToStorageAccountKeyResponseOutputWithContext(ctx context.Context) StorageAccountKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountKeyResponseOutput)
}





type StorageAccountKeyResponseArrayInput interface {
	pulumi.Input

	ToStorageAccountKeyResponseArrayOutput() StorageAccountKeyResponseArrayOutput
	ToStorageAccountKeyResponseArrayOutputWithContext(context.Context) StorageAccountKeyResponseArrayOutput
}

type StorageAccountKeyResponseArray []StorageAccountKeyResponseInput

func (StorageAccountKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountKeyResponse)(nil)).Elem()
}

func (i StorageAccountKeyResponseArray) ToStorageAccountKeyResponseArrayOutput() StorageAccountKeyResponseArrayOutput {
	return i.ToStorageAccountKeyResponseArrayOutputWithContext(context.Background())
}

func (i StorageAccountKeyResponseArray) ToStorageAccountKeyResponseArrayOutputWithContext(ctx context.Context) StorageAccountKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountKeyResponseArrayOutput)
}

type StorageAccountKeyResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountKeyResponse)(nil)).Elem()
}

func (o StorageAccountKeyResponseOutput) ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput {
	return o
}

func (o StorageAccountKeyResponseOutput) ToStorageAccountKeyResponseOutputWithContext(ctx context.Context) StorageAccountKeyResponseOutput {
	return o
}

func (o StorageAccountKeyResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o StorageAccountKeyResponseOutput) Permissions() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.Permissions }).(pulumi.StringOutput)
}

func (o StorageAccountKeyResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.Value }).(pulumi.StringOutput)
}

type StorageAccountKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountKeyResponse)(nil)).Elem()
}

func (o StorageAccountKeyResponseArrayOutput) ToStorageAccountKeyResponseArrayOutput() StorageAccountKeyResponseArrayOutput {
	return o
}

func (o StorageAccountKeyResponseArrayOutput) ToStorageAccountKeyResponseArrayOutputWithContext(ctx context.Context) StorageAccountKeyResponseArrayOutput {
	return o
}

func (o StorageAccountKeyResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountKeyResponse {
		return vs[0].([]StorageAccountKeyResponse)[vs[1].(int)]
	}).(StorageAccountKeyResponseOutput)
}

type TagPropertyResponse struct {
	ObjectIdentifier string `pulumi:"objectIdentifier"`
	Tag              string `pulumi:"tag"`
	TenantId         string `pulumi:"tenantId"`
	Timestamp        string `pulumi:"timestamp"`
	Upn              string `pulumi:"upn"`
}





type TagPropertyResponseInput interface {
	pulumi.Input

	ToTagPropertyResponseOutput() TagPropertyResponseOutput
	ToTagPropertyResponseOutputWithContext(context.Context) TagPropertyResponseOutput
}

type TagPropertyResponseArgs struct {
	ObjectIdentifier pulumi.StringInput `pulumi:"objectIdentifier"`
	Tag              pulumi.StringInput `pulumi:"tag"`
	TenantId         pulumi.StringInput `pulumi:"tenantId"`
	Timestamp        pulumi.StringInput `pulumi:"timestamp"`
	Upn              pulumi.StringInput `pulumi:"upn"`
}

func (TagPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagPropertyResponse)(nil)).Elem()
}

func (i TagPropertyResponseArgs) ToTagPropertyResponseOutput() TagPropertyResponseOutput {
	return i.ToTagPropertyResponseOutputWithContext(context.Background())
}

func (i TagPropertyResponseArgs) ToTagPropertyResponseOutputWithContext(ctx context.Context) TagPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagPropertyResponseOutput)
}





type TagPropertyResponseArrayInput interface {
	pulumi.Input

	ToTagPropertyResponseArrayOutput() TagPropertyResponseArrayOutput
	ToTagPropertyResponseArrayOutputWithContext(context.Context) TagPropertyResponseArrayOutput
}

type TagPropertyResponseArray []TagPropertyResponseInput

func (TagPropertyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagPropertyResponse)(nil)).Elem()
}

func (i TagPropertyResponseArray) ToTagPropertyResponseArrayOutput() TagPropertyResponseArrayOutput {
	return i.ToTagPropertyResponseArrayOutputWithContext(context.Background())
}

func (i TagPropertyResponseArray) ToTagPropertyResponseArrayOutputWithContext(ctx context.Context) TagPropertyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagPropertyResponseArrayOutput)
}

type TagPropertyResponseOutput struct{ *pulumi.OutputState }

func (TagPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagPropertyResponse)(nil)).Elem()
}

func (o TagPropertyResponseOutput) ToTagPropertyResponseOutput() TagPropertyResponseOutput {
	return o
}

func (o TagPropertyResponseOutput) ToTagPropertyResponseOutputWithContext(ctx context.Context) TagPropertyResponseOutput {
	return o
}

func (o TagPropertyResponseOutput) ObjectIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.ObjectIdentifier }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.Tag }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

func (o TagPropertyResponseOutput) Upn() pulumi.StringOutput {
	return o.ApplyT(func(v TagPropertyResponse) string { return v.Upn }).(pulumi.StringOutput)
}

type TagPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (TagPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagPropertyResponse)(nil)).Elem()
}

func (o TagPropertyResponseArrayOutput) ToTagPropertyResponseArrayOutput() TagPropertyResponseArrayOutput {
	return o
}

func (o TagPropertyResponseArrayOutput) ToTagPropertyResponseArrayOutputWithContext(ctx context.Context) TagPropertyResponseArrayOutput {
	return o
}

func (o TagPropertyResponseArrayOutput) Index(i pulumi.IntInput) TagPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagPropertyResponse {
		return vs[0].([]TagPropertyResponse)[vs[1].(int)]
	}).(TagPropertyResponseOutput)
}

type UpdateHistoryPropertyResponse struct {
	ImmutabilityPeriodSinceCreationInDays int    `pulumi:"immutabilityPeriodSinceCreationInDays"`
	ObjectIdentifier                      string `pulumi:"objectIdentifier"`
	TenantId                              string `pulumi:"tenantId"`
	Timestamp                             string `pulumi:"timestamp"`
	Update                                string `pulumi:"update"`
	Upn                                   string `pulumi:"upn"`
}





type UpdateHistoryPropertyResponseInput interface {
	pulumi.Input

	ToUpdateHistoryPropertyResponseOutput() UpdateHistoryPropertyResponseOutput
	ToUpdateHistoryPropertyResponseOutputWithContext(context.Context) UpdateHistoryPropertyResponseOutput
}

type UpdateHistoryPropertyResponseArgs struct {
	ImmutabilityPeriodSinceCreationInDays pulumi.IntInput    `pulumi:"immutabilityPeriodSinceCreationInDays"`
	ObjectIdentifier                      pulumi.StringInput `pulumi:"objectIdentifier"`
	TenantId                              pulumi.StringInput `pulumi:"tenantId"`
	Timestamp                             pulumi.StringInput `pulumi:"timestamp"`
	Update                                pulumi.StringInput `pulumi:"update"`
	Upn                                   pulumi.StringInput `pulumi:"upn"`
}

func (UpdateHistoryPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateHistoryPropertyResponse)(nil)).Elem()
}

func (i UpdateHistoryPropertyResponseArgs) ToUpdateHistoryPropertyResponseOutput() UpdateHistoryPropertyResponseOutput {
	return i.ToUpdateHistoryPropertyResponseOutputWithContext(context.Background())
}

func (i UpdateHistoryPropertyResponseArgs) ToUpdateHistoryPropertyResponseOutputWithContext(ctx context.Context) UpdateHistoryPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateHistoryPropertyResponseOutput)
}





type UpdateHistoryPropertyResponseArrayInput interface {
	pulumi.Input

	ToUpdateHistoryPropertyResponseArrayOutput() UpdateHistoryPropertyResponseArrayOutput
	ToUpdateHistoryPropertyResponseArrayOutputWithContext(context.Context) UpdateHistoryPropertyResponseArrayOutput
}

type UpdateHistoryPropertyResponseArray []UpdateHistoryPropertyResponseInput

func (UpdateHistoryPropertyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UpdateHistoryPropertyResponse)(nil)).Elem()
}

func (i UpdateHistoryPropertyResponseArray) ToUpdateHistoryPropertyResponseArrayOutput() UpdateHistoryPropertyResponseArrayOutput {
	return i.ToUpdateHistoryPropertyResponseArrayOutputWithContext(context.Background())
}

func (i UpdateHistoryPropertyResponseArray) ToUpdateHistoryPropertyResponseArrayOutputWithContext(ctx context.Context) UpdateHistoryPropertyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateHistoryPropertyResponseArrayOutput)
}

type UpdateHistoryPropertyResponseOutput struct{ *pulumi.OutputState }

func (UpdateHistoryPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateHistoryPropertyResponse)(nil)).Elem()
}

func (o UpdateHistoryPropertyResponseOutput) ToUpdateHistoryPropertyResponseOutput() UpdateHistoryPropertyResponseOutput {
	return o
}

func (o UpdateHistoryPropertyResponseOutput) ToUpdateHistoryPropertyResponseOutputWithContext(ctx context.Context) UpdateHistoryPropertyResponseOutput {
	return o
}

func (o UpdateHistoryPropertyResponseOutput) ImmutabilityPeriodSinceCreationInDays() pulumi.IntOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) int { return v.ImmutabilityPeriodSinceCreationInDays }).(pulumi.IntOutput)
}

func (o UpdateHistoryPropertyResponseOutput) ObjectIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.ObjectIdentifier }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) Update() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.Update }).(pulumi.StringOutput)
}

func (o UpdateHistoryPropertyResponseOutput) Upn() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateHistoryPropertyResponse) string { return v.Upn }).(pulumi.StringOutput)
}

type UpdateHistoryPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (UpdateHistoryPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UpdateHistoryPropertyResponse)(nil)).Elem()
}

func (o UpdateHistoryPropertyResponseArrayOutput) ToUpdateHistoryPropertyResponseArrayOutput() UpdateHistoryPropertyResponseArrayOutput {
	return o
}

func (o UpdateHistoryPropertyResponseArrayOutput) ToUpdateHistoryPropertyResponseArrayOutputWithContext(ctx context.Context) UpdateHistoryPropertyResponseArrayOutput {
	return o
}

func (o UpdateHistoryPropertyResponseArrayOutput) Index(i pulumi.IntInput) UpdateHistoryPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UpdateHistoryPropertyResponse {
		return vs[0].([]UpdateHistoryPropertyResponse)[vs[1].(int)]
	}).(UpdateHistoryPropertyResponseOutput)
}

type VirtualNetworkRule struct {
	Action                   *Action `pulumi:"action"`
	State                    *State  `pulumi:"state"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Action                   ActionPtrInput     `pulumi:"action"`
	State                    StatePtrInput      `pulumi:"state"`
	VirtualNetworkResourceId pulumi.StringInput `pulumi:"virtualNetworkResourceId"`
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *Action { return v.Action }).(ActionPtrOutput)
}

func (o VirtualNetworkRuleOutput) State() StatePtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *State { return v.State }).(StatePtrOutput)
}

func (o VirtualNetworkRuleOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRule) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Action                   *string `pulumi:"action"`
	State                    *string `pulumi:"state"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}





type VirtualNetworkRuleResponseInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput
	ToVirtualNetworkRuleResponseOutputWithContext(context.Context) VirtualNetworkRuleResponseOutput
}

type VirtualNetworkRuleResponseArgs struct {
	Action                   pulumi.StringPtrInput `pulumi:"action"`
	State                    pulumi.StringPtrInput `pulumi:"state"`
	VirtualNetworkResourceId pulumi.StringInput    `pulumi:"virtualNetworkResourceId"`
}

func (VirtualNetworkRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return i.ToVirtualNetworkRuleResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseOutput)
}





type VirtualNetworkRuleResponseArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput
	ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Context) VirtualNetworkRuleResponseArrayOutput
}

type VirtualNetworkRuleResponseArray []VirtualNetworkRuleResponseInput

func (VirtualNetworkRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return i.ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseArrayOutput)
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomDomainOutput{})
	pulumi.RegisterOutputType(CustomDomainPtrOutput{})
	pulumi.RegisterOutputType(CustomDomainResponseOutput{})
	pulumi.RegisterOutputType(CustomDomainResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServiceOutput{})
	pulumi.RegisterOutputType(EncryptionServicePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServiceResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServiceResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServicesOutput{})
	pulumi.RegisterOutputType(EncryptionServicesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionServicesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServicesResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointsResponseOutput{})
	pulumi.RegisterOutputType(EndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(IPRuleOutput{})
	pulumi.RegisterOutputType(IPRuleArrayOutput{})
	pulumi.RegisterOutputType(IPRuleResponseOutput{})
	pulumi.RegisterOutputType(IPRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ImmutabilityPolicyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ImmutabilityPolicyPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LegalHoldPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LegalHoldPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(RestrictionOutput{})
	pulumi.RegisterOutputType(RestrictionArrayOutput{})
	pulumi.RegisterOutputType(RestrictionResponseOutput{})
	pulumi.RegisterOutputType(RestrictionResponseArrayOutput{})
	pulumi.RegisterOutputType(SKUCapabilityResponseOutput{})
	pulumi.RegisterOutputType(SKUCapabilityResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountKeyResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(TagPropertyResponseOutput{})
	pulumi.RegisterOutputType(TagPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(UpdateHistoryPropertyResponseOutput{})
	pulumi.RegisterOutputType(UpdateHistoryPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
