


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AsymmetricEncryptedSecret struct {
	EncryptionAlgorithm             EncryptionAlgorithm `pulumi:"encryptionAlgorithm"`
	EncryptionCertificateThumbprint *string             `pulumi:"encryptionCertificateThumbprint"`
	Value                           string              `pulumi:"value"`
}





type AsymmetricEncryptedSecretInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretOutput() AsymmetricEncryptedSecretOutput
	ToAsymmetricEncryptedSecretOutputWithContext(context.Context) AsymmetricEncryptedSecretOutput
}

type AsymmetricEncryptedSecretArgs struct {
	EncryptionAlgorithm             EncryptionAlgorithmInput `pulumi:"encryptionAlgorithm"`
	EncryptionCertificateThumbprint pulumi.StringPtrInput    `pulumi:"encryptionCertificateThumbprint"`
	Value                           pulumi.StringInput       `pulumi:"value"`
}

func (AsymmetricEncryptedSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecret)(nil)).Elem()
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretOutput() AsymmetricEncryptedSecretOutput {
	return i.ToAsymmetricEncryptedSecretOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretOutput)
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return i.ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretOutput).ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx)
}









type AsymmetricEncryptedSecretPtrInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput
	ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Context) AsymmetricEncryptedSecretPtrOutput
}

type asymmetricEncryptedSecretPtrType AsymmetricEncryptedSecretArgs

func AsymmetricEncryptedSecretPtr(v *AsymmetricEncryptedSecretArgs) AsymmetricEncryptedSecretPtrInput {
	return (*asymmetricEncryptedSecretPtrType)(v)
}

func (*asymmetricEncryptedSecretPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecret)(nil)).Elem()
}

func (i *asymmetricEncryptedSecretPtrType) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return i.ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Background())
}

func (i *asymmetricEncryptedSecretPtrType) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretPtrOutput)
}

type AsymmetricEncryptedSecretOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecret)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretOutput() AsymmetricEncryptedSecretOutput {
	return o
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretOutput {
	return o
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return o.ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Background())
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AsymmetricEncryptedSecret) *AsymmetricEncryptedSecret {
		return &v
	}).(AsymmetricEncryptedSecretPtrOutput)
}

func (o AsymmetricEncryptedSecretOutput) EncryptionAlgorithm() EncryptionAlgorithmOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecret) EncryptionAlgorithm { return v.EncryptionAlgorithm }).(EncryptionAlgorithmOutput)
}

func (o AsymmetricEncryptedSecretOutput) EncryptionCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecret) *string { return v.EncryptionCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecret) string { return v.Value }).(pulumi.StringOutput)
}

type AsymmetricEncryptedSecretPtrOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecret)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretPtrOutput) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretPtrOutput) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretPtrOutput) Elem() AsymmetricEncryptedSecretOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) AsymmetricEncryptedSecret {
		if v != nil {
			return *v
		}
		var ret AsymmetricEncryptedSecret
		return ret
	}).(AsymmetricEncryptedSecretOutput)
}

func (o AsymmetricEncryptedSecretPtrOutput) EncryptionAlgorithm() EncryptionAlgorithmPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) *EncryptionAlgorithm {
		if v == nil {
			return nil
		}
		return &v.EncryptionAlgorithm
	}).(EncryptionAlgorithmPtrOutput)
}

func (o AsymmetricEncryptedSecretPtrOutput) EncryptionCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionCertificateThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type AsymmetricEncryptedSecretResponse struct {
	EncryptionAlgorithm             string  `pulumi:"encryptionAlgorithm"`
	EncryptionCertificateThumbprint *string `pulumi:"encryptionCertificateThumbprint"`
	Value                           string  `pulumi:"value"`
}





type AsymmetricEncryptedSecretResponseInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretResponseOutput() AsymmetricEncryptedSecretResponseOutput
	ToAsymmetricEncryptedSecretResponseOutputWithContext(context.Context) AsymmetricEncryptedSecretResponseOutput
}

type AsymmetricEncryptedSecretResponseArgs struct {
	EncryptionAlgorithm             pulumi.StringInput    `pulumi:"encryptionAlgorithm"`
	EncryptionCertificateThumbprint pulumi.StringPtrInput `pulumi:"encryptionCertificateThumbprint"`
	Value                           pulumi.StringInput    `pulumi:"value"`
}

func (AsymmetricEncryptedSecretResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponseOutput() AsymmetricEncryptedSecretResponseOutput {
	return i.ToAsymmetricEncryptedSecretResponseOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponseOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretResponseOutput)
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return i.ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretResponseOutput).ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx)
}









type AsymmetricEncryptedSecretResponsePtrInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput
	ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Context) AsymmetricEncryptedSecretResponsePtrOutput
}

type asymmetricEncryptedSecretResponsePtrType AsymmetricEncryptedSecretResponseArgs

func AsymmetricEncryptedSecretResponsePtr(v *AsymmetricEncryptedSecretResponseArgs) AsymmetricEncryptedSecretResponsePtrInput {
	return (*asymmetricEncryptedSecretResponsePtrType)(v)
}

func (*asymmetricEncryptedSecretResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (i *asymmetricEncryptedSecretResponsePtrType) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return i.ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Background())
}

func (i *asymmetricEncryptedSecretResponsePtrType) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretResponsePtrOutput)
}

type AsymmetricEncryptedSecretResponseOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponseOutput() AsymmetricEncryptedSecretResponseOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponseOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponseOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Background())
}

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AsymmetricEncryptedSecretResponse) *AsymmetricEncryptedSecretResponse {
		return &v
	}).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o AsymmetricEncryptedSecretResponseOutput) EncryptionAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecretResponse) string { return v.EncryptionAlgorithm }).(pulumi.StringOutput)
}

func (o AsymmetricEncryptedSecretResponseOutput) EncryptionCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecretResponse) *string { return v.EncryptionCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecretResponse) string { return v.Value }).(pulumi.StringOutput)
}

type AsymmetricEncryptedSecretResponsePtrOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) Elem() AsymmetricEncryptedSecretResponseOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) AsymmetricEncryptedSecretResponse {
		if v != nil {
			return *v
		}
		var ret AsymmetricEncryptedSecretResponse
		return ret
	}).(AsymmetricEncryptedSecretResponseOutput)
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EncryptionAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) EncryptionCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionCertificateThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type ManagerIntrinsicSettings struct {
	Type ManagerType `pulumi:"type"`
}





type ManagerIntrinsicSettingsInput interface {
	pulumi.Input

	ToManagerIntrinsicSettingsOutput() ManagerIntrinsicSettingsOutput
	ToManagerIntrinsicSettingsOutputWithContext(context.Context) ManagerIntrinsicSettingsOutput
}

type ManagerIntrinsicSettingsArgs struct {
	Type ManagerTypeInput `pulumi:"type"`
}

func (ManagerIntrinsicSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerIntrinsicSettings)(nil)).Elem()
}

func (i ManagerIntrinsicSettingsArgs) ToManagerIntrinsicSettingsOutput() ManagerIntrinsicSettingsOutput {
	return i.ToManagerIntrinsicSettingsOutputWithContext(context.Background())
}

func (i ManagerIntrinsicSettingsArgs) ToManagerIntrinsicSettingsOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsOutput)
}

func (i ManagerIntrinsicSettingsArgs) ToManagerIntrinsicSettingsPtrOutput() ManagerIntrinsicSettingsPtrOutput {
	return i.ToManagerIntrinsicSettingsPtrOutputWithContext(context.Background())
}

func (i ManagerIntrinsicSettingsArgs) ToManagerIntrinsicSettingsPtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsOutput).ToManagerIntrinsicSettingsPtrOutputWithContext(ctx)
}









type ManagerIntrinsicSettingsPtrInput interface {
	pulumi.Input

	ToManagerIntrinsicSettingsPtrOutput() ManagerIntrinsicSettingsPtrOutput
	ToManagerIntrinsicSettingsPtrOutputWithContext(context.Context) ManagerIntrinsicSettingsPtrOutput
}

type managerIntrinsicSettingsPtrType ManagerIntrinsicSettingsArgs

func ManagerIntrinsicSettingsPtr(v *ManagerIntrinsicSettingsArgs) ManagerIntrinsicSettingsPtrInput {
	return (*managerIntrinsicSettingsPtrType)(v)
}

func (*managerIntrinsicSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerIntrinsicSettings)(nil)).Elem()
}

func (i *managerIntrinsicSettingsPtrType) ToManagerIntrinsicSettingsPtrOutput() ManagerIntrinsicSettingsPtrOutput {
	return i.ToManagerIntrinsicSettingsPtrOutputWithContext(context.Background())
}

func (i *managerIntrinsicSettingsPtrType) ToManagerIntrinsicSettingsPtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsPtrOutput)
}

type ManagerIntrinsicSettingsOutput struct{ *pulumi.OutputState }

func (ManagerIntrinsicSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerIntrinsicSettings)(nil)).Elem()
}

func (o ManagerIntrinsicSettingsOutput) ToManagerIntrinsicSettingsOutput() ManagerIntrinsicSettingsOutput {
	return o
}

func (o ManagerIntrinsicSettingsOutput) ToManagerIntrinsicSettingsOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsOutput {
	return o
}

func (o ManagerIntrinsicSettingsOutput) ToManagerIntrinsicSettingsPtrOutput() ManagerIntrinsicSettingsPtrOutput {
	return o.ToManagerIntrinsicSettingsPtrOutputWithContext(context.Background())
}

func (o ManagerIntrinsicSettingsOutput) ToManagerIntrinsicSettingsPtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagerIntrinsicSettings) *ManagerIntrinsicSettings {
		return &v
	}).(ManagerIntrinsicSettingsPtrOutput)
}

func (o ManagerIntrinsicSettingsOutput) Type() ManagerTypeOutput {
	return o.ApplyT(func(v ManagerIntrinsicSettings) ManagerType { return v.Type }).(ManagerTypeOutput)
}

type ManagerIntrinsicSettingsPtrOutput struct{ *pulumi.OutputState }

func (ManagerIntrinsicSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerIntrinsicSettings)(nil)).Elem()
}

func (o ManagerIntrinsicSettingsPtrOutput) ToManagerIntrinsicSettingsPtrOutput() ManagerIntrinsicSettingsPtrOutput {
	return o
}

func (o ManagerIntrinsicSettingsPtrOutput) ToManagerIntrinsicSettingsPtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsPtrOutput {
	return o
}

func (o ManagerIntrinsicSettingsPtrOutput) Elem() ManagerIntrinsicSettingsOutput {
	return o.ApplyT(func(v *ManagerIntrinsicSettings) ManagerIntrinsicSettings {
		if v != nil {
			return *v
		}
		var ret ManagerIntrinsicSettings
		return ret
	}).(ManagerIntrinsicSettingsOutput)
}

func (o ManagerIntrinsicSettingsPtrOutput) Type() ManagerTypePtrOutput {
	return o.ApplyT(func(v *ManagerIntrinsicSettings) *ManagerType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(ManagerTypePtrOutput)
}

type ManagerIntrinsicSettingsResponse struct {
	Type string `pulumi:"type"`
}





type ManagerIntrinsicSettingsResponseInput interface {
	pulumi.Input

	ToManagerIntrinsicSettingsResponseOutput() ManagerIntrinsicSettingsResponseOutput
	ToManagerIntrinsicSettingsResponseOutputWithContext(context.Context) ManagerIntrinsicSettingsResponseOutput
}

type ManagerIntrinsicSettingsResponseArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (ManagerIntrinsicSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerIntrinsicSettingsResponse)(nil)).Elem()
}

func (i ManagerIntrinsicSettingsResponseArgs) ToManagerIntrinsicSettingsResponseOutput() ManagerIntrinsicSettingsResponseOutput {
	return i.ToManagerIntrinsicSettingsResponseOutputWithContext(context.Background())
}

func (i ManagerIntrinsicSettingsResponseArgs) ToManagerIntrinsicSettingsResponseOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsResponseOutput)
}

func (i ManagerIntrinsicSettingsResponseArgs) ToManagerIntrinsicSettingsResponsePtrOutput() ManagerIntrinsicSettingsResponsePtrOutput {
	return i.ToManagerIntrinsicSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ManagerIntrinsicSettingsResponseArgs) ToManagerIntrinsicSettingsResponsePtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsResponseOutput).ToManagerIntrinsicSettingsResponsePtrOutputWithContext(ctx)
}









type ManagerIntrinsicSettingsResponsePtrInput interface {
	pulumi.Input

	ToManagerIntrinsicSettingsResponsePtrOutput() ManagerIntrinsicSettingsResponsePtrOutput
	ToManagerIntrinsicSettingsResponsePtrOutputWithContext(context.Context) ManagerIntrinsicSettingsResponsePtrOutput
}

type managerIntrinsicSettingsResponsePtrType ManagerIntrinsicSettingsResponseArgs

func ManagerIntrinsicSettingsResponsePtr(v *ManagerIntrinsicSettingsResponseArgs) ManagerIntrinsicSettingsResponsePtrInput {
	return (*managerIntrinsicSettingsResponsePtrType)(v)
}

func (*managerIntrinsicSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerIntrinsicSettingsResponse)(nil)).Elem()
}

func (i *managerIntrinsicSettingsResponsePtrType) ToManagerIntrinsicSettingsResponsePtrOutput() ManagerIntrinsicSettingsResponsePtrOutput {
	return i.ToManagerIntrinsicSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *managerIntrinsicSettingsResponsePtrType) ToManagerIntrinsicSettingsResponsePtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsResponsePtrOutput)
}

type ManagerIntrinsicSettingsResponseOutput struct{ *pulumi.OutputState }

func (ManagerIntrinsicSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerIntrinsicSettingsResponse)(nil)).Elem()
}

func (o ManagerIntrinsicSettingsResponseOutput) ToManagerIntrinsicSettingsResponseOutput() ManagerIntrinsicSettingsResponseOutput {
	return o
}

func (o ManagerIntrinsicSettingsResponseOutput) ToManagerIntrinsicSettingsResponseOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponseOutput {
	return o
}

func (o ManagerIntrinsicSettingsResponseOutput) ToManagerIntrinsicSettingsResponsePtrOutput() ManagerIntrinsicSettingsResponsePtrOutput {
	return o.ToManagerIntrinsicSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ManagerIntrinsicSettingsResponseOutput) ToManagerIntrinsicSettingsResponsePtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagerIntrinsicSettingsResponse) *ManagerIntrinsicSettingsResponse {
		return &v
	}).(ManagerIntrinsicSettingsResponsePtrOutput)
}

func (o ManagerIntrinsicSettingsResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagerIntrinsicSettingsResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ManagerIntrinsicSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagerIntrinsicSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerIntrinsicSettingsResponse)(nil)).Elem()
}

func (o ManagerIntrinsicSettingsResponsePtrOutput) ToManagerIntrinsicSettingsResponsePtrOutput() ManagerIntrinsicSettingsResponsePtrOutput {
	return o
}

func (o ManagerIntrinsicSettingsResponsePtrOutput) ToManagerIntrinsicSettingsResponsePtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponsePtrOutput {
	return o
}

func (o ManagerIntrinsicSettingsResponsePtrOutput) Elem() ManagerIntrinsicSettingsResponseOutput {
	return o.ApplyT(func(v *ManagerIntrinsicSettingsResponse) ManagerIntrinsicSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ManagerIntrinsicSettingsResponse
		return ret
	}).(ManagerIntrinsicSettingsResponseOutput)
}

func (o ManagerIntrinsicSettingsResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagerIntrinsicSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ManagerSku struct {
	Name ManagerSkuType `pulumi:"name"`
}





type ManagerSkuInput interface {
	pulumi.Input

	ToManagerSkuOutput() ManagerSkuOutput
	ToManagerSkuOutputWithContext(context.Context) ManagerSkuOutput
}

type ManagerSkuArgs struct {
	Name ManagerSkuTypeInput `pulumi:"name"`
}

func (ManagerSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerSku)(nil)).Elem()
}

func (i ManagerSkuArgs) ToManagerSkuOutput() ManagerSkuOutput {
	return i.ToManagerSkuOutputWithContext(context.Background())
}

func (i ManagerSkuArgs) ToManagerSkuOutputWithContext(ctx context.Context) ManagerSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuOutput)
}

func (i ManagerSkuArgs) ToManagerSkuPtrOutput() ManagerSkuPtrOutput {
	return i.ToManagerSkuPtrOutputWithContext(context.Background())
}

func (i ManagerSkuArgs) ToManagerSkuPtrOutputWithContext(ctx context.Context) ManagerSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuOutput).ToManagerSkuPtrOutputWithContext(ctx)
}









type ManagerSkuPtrInput interface {
	pulumi.Input

	ToManagerSkuPtrOutput() ManagerSkuPtrOutput
	ToManagerSkuPtrOutputWithContext(context.Context) ManagerSkuPtrOutput
}

type managerSkuPtrType ManagerSkuArgs

func ManagerSkuPtr(v *ManagerSkuArgs) ManagerSkuPtrInput {
	return (*managerSkuPtrType)(v)
}

func (*managerSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerSku)(nil)).Elem()
}

func (i *managerSkuPtrType) ToManagerSkuPtrOutput() ManagerSkuPtrOutput {
	return i.ToManagerSkuPtrOutputWithContext(context.Background())
}

func (i *managerSkuPtrType) ToManagerSkuPtrOutputWithContext(ctx context.Context) ManagerSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuPtrOutput)
}

type ManagerSkuOutput struct{ *pulumi.OutputState }

func (ManagerSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerSku)(nil)).Elem()
}

func (o ManagerSkuOutput) ToManagerSkuOutput() ManagerSkuOutput {
	return o
}

func (o ManagerSkuOutput) ToManagerSkuOutputWithContext(ctx context.Context) ManagerSkuOutput {
	return o
}

func (o ManagerSkuOutput) ToManagerSkuPtrOutput() ManagerSkuPtrOutput {
	return o.ToManagerSkuPtrOutputWithContext(context.Background())
}

func (o ManagerSkuOutput) ToManagerSkuPtrOutputWithContext(ctx context.Context) ManagerSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagerSku) *ManagerSku {
		return &v
	}).(ManagerSkuPtrOutput)
}

func (o ManagerSkuOutput) Name() ManagerSkuTypeOutput {
	return o.ApplyT(func(v ManagerSku) ManagerSkuType { return v.Name }).(ManagerSkuTypeOutput)
}

type ManagerSkuPtrOutput struct{ *pulumi.OutputState }

func (ManagerSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerSku)(nil)).Elem()
}

func (o ManagerSkuPtrOutput) ToManagerSkuPtrOutput() ManagerSkuPtrOutput {
	return o
}

func (o ManagerSkuPtrOutput) ToManagerSkuPtrOutputWithContext(ctx context.Context) ManagerSkuPtrOutput {
	return o
}

func (o ManagerSkuPtrOutput) Elem() ManagerSkuOutput {
	return o.ApplyT(func(v *ManagerSku) ManagerSku {
		if v != nil {
			return *v
		}
		var ret ManagerSku
		return ret
	}).(ManagerSkuOutput)
}

func (o ManagerSkuPtrOutput) Name() ManagerSkuTypePtrOutput {
	return o.ApplyT(func(v *ManagerSku) *ManagerSkuType {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(ManagerSkuTypePtrOutput)
}

type ManagerSkuResponse struct {
	Name string `pulumi:"name"`
}





type ManagerSkuResponseInput interface {
	pulumi.Input

	ToManagerSkuResponseOutput() ManagerSkuResponseOutput
	ToManagerSkuResponseOutputWithContext(context.Context) ManagerSkuResponseOutput
}

type ManagerSkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (ManagerSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerSkuResponse)(nil)).Elem()
}

func (i ManagerSkuResponseArgs) ToManagerSkuResponseOutput() ManagerSkuResponseOutput {
	return i.ToManagerSkuResponseOutputWithContext(context.Background())
}

func (i ManagerSkuResponseArgs) ToManagerSkuResponseOutputWithContext(ctx context.Context) ManagerSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuResponseOutput)
}

func (i ManagerSkuResponseArgs) ToManagerSkuResponsePtrOutput() ManagerSkuResponsePtrOutput {
	return i.ToManagerSkuResponsePtrOutputWithContext(context.Background())
}

func (i ManagerSkuResponseArgs) ToManagerSkuResponsePtrOutputWithContext(ctx context.Context) ManagerSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuResponseOutput).ToManagerSkuResponsePtrOutputWithContext(ctx)
}









type ManagerSkuResponsePtrInput interface {
	pulumi.Input

	ToManagerSkuResponsePtrOutput() ManagerSkuResponsePtrOutput
	ToManagerSkuResponsePtrOutputWithContext(context.Context) ManagerSkuResponsePtrOutput
}

type managerSkuResponsePtrType ManagerSkuResponseArgs

func ManagerSkuResponsePtr(v *ManagerSkuResponseArgs) ManagerSkuResponsePtrInput {
	return (*managerSkuResponsePtrType)(v)
}

func (*managerSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerSkuResponse)(nil)).Elem()
}

func (i *managerSkuResponsePtrType) ToManagerSkuResponsePtrOutput() ManagerSkuResponsePtrOutput {
	return i.ToManagerSkuResponsePtrOutputWithContext(context.Background())
}

func (i *managerSkuResponsePtrType) ToManagerSkuResponsePtrOutputWithContext(ctx context.Context) ManagerSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuResponsePtrOutput)
}

type ManagerSkuResponseOutput struct{ *pulumi.OutputState }

func (ManagerSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerSkuResponse)(nil)).Elem()
}

func (o ManagerSkuResponseOutput) ToManagerSkuResponseOutput() ManagerSkuResponseOutput {
	return o
}

func (o ManagerSkuResponseOutput) ToManagerSkuResponseOutputWithContext(ctx context.Context) ManagerSkuResponseOutput {
	return o
}

func (o ManagerSkuResponseOutput) ToManagerSkuResponsePtrOutput() ManagerSkuResponsePtrOutput {
	return o.ToManagerSkuResponsePtrOutputWithContext(context.Background())
}

func (o ManagerSkuResponseOutput) ToManagerSkuResponsePtrOutputWithContext(ctx context.Context) ManagerSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagerSkuResponse) *ManagerSkuResponse {
		return &v
	}).(ManagerSkuResponsePtrOutput)
}

func (o ManagerSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagerSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ManagerSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagerSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerSkuResponse)(nil)).Elem()
}

func (o ManagerSkuResponsePtrOutput) ToManagerSkuResponsePtrOutput() ManagerSkuResponsePtrOutput {
	return o
}

func (o ManagerSkuResponsePtrOutput) ToManagerSkuResponsePtrOutputWithContext(ctx context.Context) ManagerSkuResponsePtrOutput {
	return o
}

func (o ManagerSkuResponsePtrOutput) Elem() ManagerSkuResponseOutput {
	return o.ApplyT(func(v *ManagerSkuResponse) ManagerSkuResponse {
		if v != nil {
			return *v
		}
		var ret ManagerSkuResponse
		return ret
	}).(ManagerSkuResponseOutput)
}

func (o ManagerSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagerSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type Time struct {
	Hour   int `pulumi:"hour"`
	Minute int `pulumi:"minute"`
}





type TimeInput interface {
	pulumi.Input

	ToTimeOutput() TimeOutput
	ToTimeOutputWithContext(context.Context) TimeOutput
}

type TimeArgs struct {
	Hour   pulumi.IntInput `pulumi:"hour"`
	Minute pulumi.IntInput `pulumi:"minute"`
}

func (TimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Time)(nil)).Elem()
}

func (i TimeArgs) ToTimeOutput() TimeOutput {
	return i.ToTimeOutputWithContext(context.Background())
}

func (i TimeArgs) ToTimeOutputWithContext(ctx context.Context) TimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeOutput)
}

func (i TimeArgs) ToTimePtrOutput() TimePtrOutput {
	return i.ToTimePtrOutputWithContext(context.Background())
}

func (i TimeArgs) ToTimePtrOutputWithContext(ctx context.Context) TimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeOutput).ToTimePtrOutputWithContext(ctx)
}









type TimePtrInput interface {
	pulumi.Input

	ToTimePtrOutput() TimePtrOutput
	ToTimePtrOutputWithContext(context.Context) TimePtrOutput
}

type timePtrType TimeArgs

func TimePtr(v *TimeArgs) TimePtrInput {
	return (*timePtrType)(v)
}

func (*timePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Time)(nil)).Elem()
}

func (i *timePtrType) ToTimePtrOutput() TimePtrOutput {
	return i.ToTimePtrOutputWithContext(context.Background())
}

func (i *timePtrType) ToTimePtrOutputWithContext(ctx context.Context) TimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimePtrOutput)
}

type TimeOutput struct{ *pulumi.OutputState }

func (TimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Time)(nil)).Elem()
}

func (o TimeOutput) ToTimeOutput() TimeOutput {
	return o
}

func (o TimeOutput) ToTimeOutputWithContext(ctx context.Context) TimeOutput {
	return o
}

func (o TimeOutput) ToTimePtrOutput() TimePtrOutput {
	return o.ToTimePtrOutputWithContext(context.Background())
}

func (o TimeOutput) ToTimePtrOutputWithContext(ctx context.Context) TimePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Time) *Time {
		return &v
	}).(TimePtrOutput)
}

func (o TimeOutput) Hour() pulumi.IntOutput {
	return o.ApplyT(func(v Time) int { return v.Hour }).(pulumi.IntOutput)
}

func (o TimeOutput) Minute() pulumi.IntOutput {
	return o.ApplyT(func(v Time) int { return v.Minute }).(pulumi.IntOutput)
}

type TimePtrOutput struct{ *pulumi.OutputState }

func (TimePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Time)(nil)).Elem()
}

func (o TimePtrOutput) ToTimePtrOutput() TimePtrOutput {
	return o
}

func (o TimePtrOutput) ToTimePtrOutputWithContext(ctx context.Context) TimePtrOutput {
	return o
}

func (o TimePtrOutput) Elem() TimeOutput {
	return o.ApplyT(func(v *Time) Time {
		if v != nil {
			return *v
		}
		var ret Time
		return ret
	}).(TimeOutput)
}

func (o TimePtrOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Time) *int {
		if v == nil {
			return nil
		}
		return &v.Hour
	}).(pulumi.IntPtrOutput)
}

func (o TimePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Time) *int {
		if v == nil {
			return nil
		}
		return &v.Minute
	}).(pulumi.IntPtrOutput)
}

type TimeResponse struct {
	Hour   int `pulumi:"hour"`
	Minute int `pulumi:"minute"`
}





type TimeResponseInput interface {
	pulumi.Input

	ToTimeResponseOutput() TimeResponseOutput
	ToTimeResponseOutputWithContext(context.Context) TimeResponseOutput
}

type TimeResponseArgs struct {
	Hour   pulumi.IntInput `pulumi:"hour"`
	Minute pulumi.IntInput `pulumi:"minute"`
}

func (TimeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeResponse)(nil)).Elem()
}

func (i TimeResponseArgs) ToTimeResponseOutput() TimeResponseOutput {
	return i.ToTimeResponseOutputWithContext(context.Background())
}

func (i TimeResponseArgs) ToTimeResponseOutputWithContext(ctx context.Context) TimeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeResponseOutput)
}

func (i TimeResponseArgs) ToTimeResponsePtrOutput() TimeResponsePtrOutput {
	return i.ToTimeResponsePtrOutputWithContext(context.Background())
}

func (i TimeResponseArgs) ToTimeResponsePtrOutputWithContext(ctx context.Context) TimeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeResponseOutput).ToTimeResponsePtrOutputWithContext(ctx)
}









type TimeResponsePtrInput interface {
	pulumi.Input

	ToTimeResponsePtrOutput() TimeResponsePtrOutput
	ToTimeResponsePtrOutputWithContext(context.Context) TimeResponsePtrOutput
}

type timeResponsePtrType TimeResponseArgs

func TimeResponsePtr(v *TimeResponseArgs) TimeResponsePtrInput {
	return (*timeResponsePtrType)(v)
}

func (*timeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeResponse)(nil)).Elem()
}

func (i *timeResponsePtrType) ToTimeResponsePtrOutput() TimeResponsePtrOutput {
	return i.ToTimeResponsePtrOutputWithContext(context.Background())
}

func (i *timeResponsePtrType) ToTimeResponsePtrOutputWithContext(ctx context.Context) TimeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeResponsePtrOutput)
}

type TimeResponseOutput struct{ *pulumi.OutputState }

func (TimeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeResponse)(nil)).Elem()
}

func (o TimeResponseOutput) ToTimeResponseOutput() TimeResponseOutput {
	return o
}

func (o TimeResponseOutput) ToTimeResponseOutputWithContext(ctx context.Context) TimeResponseOutput {
	return o
}

func (o TimeResponseOutput) ToTimeResponsePtrOutput() TimeResponsePtrOutput {
	return o.ToTimeResponsePtrOutputWithContext(context.Background())
}

func (o TimeResponseOutput) ToTimeResponsePtrOutputWithContext(ctx context.Context) TimeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeResponse) *TimeResponse {
		return &v
	}).(TimeResponsePtrOutput)
}

func (o TimeResponseOutput) Hour() pulumi.IntOutput {
	return o.ApplyT(func(v TimeResponse) int { return v.Hour }).(pulumi.IntOutput)
}

func (o TimeResponseOutput) Minute() pulumi.IntOutput {
	return o.ApplyT(func(v TimeResponse) int { return v.Minute }).(pulumi.IntOutput)
}

type TimeResponsePtrOutput struct{ *pulumi.OutputState }

func (TimeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeResponse)(nil)).Elem()
}

func (o TimeResponsePtrOutput) ToTimeResponsePtrOutput() TimeResponsePtrOutput {
	return o
}

func (o TimeResponsePtrOutput) ToTimeResponsePtrOutputWithContext(ctx context.Context) TimeResponsePtrOutput {
	return o
}

func (o TimeResponsePtrOutput) Elem() TimeResponseOutput {
	return o.ApplyT(func(v *TimeResponse) TimeResponse {
		if v != nil {
			return *v
		}
		var ret TimeResponse
		return ret
	}).(TimeResponseOutput)
}

func (o TimeResponsePtrOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimeResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Hour
	}).(pulumi.IntPtrOutput)
}

func (o TimeResponsePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimeResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Minute
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AsymmetricEncryptedSecretInput)(nil)).Elem(), AsymmetricEncryptedSecretArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AsymmetricEncryptedSecretPtrInput)(nil)).Elem(), AsymmetricEncryptedSecretArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AsymmetricEncryptedSecretResponseInput)(nil)).Elem(), AsymmetricEncryptedSecretResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AsymmetricEncryptedSecretResponsePtrInput)(nil)).Elem(), AsymmetricEncryptedSecretResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagerIntrinsicSettingsInput)(nil)).Elem(), ManagerIntrinsicSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagerIntrinsicSettingsPtrInput)(nil)).Elem(), ManagerIntrinsicSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagerIntrinsicSettingsResponseInput)(nil)).Elem(), ManagerIntrinsicSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagerIntrinsicSettingsResponsePtrInput)(nil)).Elem(), ManagerIntrinsicSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagerSkuInput)(nil)).Elem(), ManagerSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagerSkuPtrInput)(nil)).Elem(), ManagerSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagerSkuResponseInput)(nil)).Elem(), ManagerSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagerSkuResponsePtrInput)(nil)).Elem(), ManagerSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimeInput)(nil)).Elem(), TimeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimePtrInput)(nil)).Elem(), TimeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimeResponseInput)(nil)).Elem(), TimeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimeResponsePtrInput)(nil)).Elem(), TimeResponseArgs{})
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretOutput{})
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretPtrOutput{})
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretResponseOutput{})
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagerIntrinsicSettingsOutput{})
	pulumi.RegisterOutputType(ManagerIntrinsicSettingsPtrOutput{})
	pulumi.RegisterOutputType(ManagerIntrinsicSettingsResponseOutput{})
	pulumi.RegisterOutputType(ManagerIntrinsicSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagerSkuOutput{})
	pulumi.RegisterOutputType(ManagerSkuPtrOutput{})
	pulumi.RegisterOutputType(ManagerSkuResponseOutput{})
	pulumi.RegisterOutputType(ManagerSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(TimeOutput{})
	pulumi.RegisterOutputType(TimePtrOutput{})
	pulumi.RegisterOutputType(TimeResponseOutput{})
	pulumi.RegisterOutputType(TimeResponsePtrOutput{})
}
