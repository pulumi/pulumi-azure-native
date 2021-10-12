


package videoanalyzer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountEncryption struct {
	Identity           *ResourceIdentity   `pulumi:"identity"`
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	Type               string              `pulumi:"type"`
}





type AccountEncryptionInput interface {
	pulumi.Input

	ToAccountEncryptionOutput() AccountEncryptionOutput
	ToAccountEncryptionOutputWithContext(context.Context) AccountEncryptionOutput
}

type AccountEncryptionArgs struct {
	Identity           ResourceIdentityPtrInput   `pulumi:"identity"`
	KeyVaultProperties KeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
	Type               pulumi.StringInput         `pulumi:"type"`
}

func (AccountEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryption)(nil)).Elem()
}

func (i AccountEncryptionArgs) ToAccountEncryptionOutput() AccountEncryptionOutput {
	return i.ToAccountEncryptionOutputWithContext(context.Background())
}

func (i AccountEncryptionArgs) ToAccountEncryptionOutputWithContext(ctx context.Context) AccountEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionOutput)
}

func (i AccountEncryptionArgs) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return i.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (i AccountEncryptionArgs) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionOutput).ToAccountEncryptionPtrOutputWithContext(ctx)
}









type AccountEncryptionPtrInput interface {
	pulumi.Input

	ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput
	ToAccountEncryptionPtrOutputWithContext(context.Context) AccountEncryptionPtrOutput
}

type accountEncryptionPtrType AccountEncryptionArgs

func AccountEncryptionPtr(v *AccountEncryptionArgs) AccountEncryptionPtrInput {
	return (*accountEncryptionPtrType)(v)
}

func (*accountEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryption)(nil)).Elem()
}

func (i *accountEncryptionPtrType) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return i.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (i *accountEncryptionPtrType) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionPtrOutput)
}

type AccountEncryptionOutput struct{ *pulumi.OutputState }

func (AccountEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryption)(nil)).Elem()
}

func (o AccountEncryptionOutput) ToAccountEncryptionOutput() AccountEncryptionOutput {
	return o
}

func (o AccountEncryptionOutput) ToAccountEncryptionOutputWithContext(ctx context.Context) AccountEncryptionOutput {
	return o
}

func (o AccountEncryptionOutput) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return o.ToAccountEncryptionPtrOutputWithContext(context.Background())
}

func (o AccountEncryptionOutput) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountEncryption) *AccountEncryption {
		return &v
	}).(AccountEncryptionPtrOutput)
}

func (o AccountEncryptionOutput) Identity() ResourceIdentityPtrOutput {
	return o.ApplyT(func(v AccountEncryption) *ResourceIdentity { return v.Identity }).(ResourceIdentityPtrOutput)
}

func (o AccountEncryptionOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v AccountEncryption) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

func (o AccountEncryptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AccountEncryption) string { return v.Type }).(pulumi.StringOutput)
}

type AccountEncryptionPtrOutput struct{ *pulumi.OutputState }

func (AccountEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryption)(nil)).Elem()
}

func (o AccountEncryptionPtrOutput) ToAccountEncryptionPtrOutput() AccountEncryptionPtrOutput {
	return o
}

func (o AccountEncryptionPtrOutput) ToAccountEncryptionPtrOutputWithContext(ctx context.Context) AccountEncryptionPtrOutput {
	return o
}

func (o AccountEncryptionPtrOutput) Elem() AccountEncryptionOutput {
	return o.ApplyT(func(v *AccountEncryption) AccountEncryption {
		if v != nil {
			return *v
		}
		var ret AccountEncryption
		return ret
	}).(AccountEncryptionOutput)
}

func (o AccountEncryptionPtrOutput) Identity() ResourceIdentityPtrOutput {
	return o.ApplyT(func(v *AccountEncryption) *ResourceIdentity {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(ResourceIdentityPtrOutput)
}

func (o AccountEncryptionPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *AccountEncryption) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

func (o AccountEncryptionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountEncryption) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type AccountEncryptionResponse struct {
	Identity           *ResourceIdentityResponse   `pulumi:"identity"`
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Status             string                      `pulumi:"status"`
	Type               string                      `pulumi:"type"`
}





type AccountEncryptionResponseInput interface {
	pulumi.Input

	ToAccountEncryptionResponseOutput() AccountEncryptionResponseOutput
	ToAccountEncryptionResponseOutputWithContext(context.Context) AccountEncryptionResponseOutput
}

type AccountEncryptionResponseArgs struct {
	Identity           ResourceIdentityResponsePtrInput   `pulumi:"identity"`
	KeyVaultProperties KeyVaultPropertiesResponsePtrInput `pulumi:"keyVaultProperties"`
	Status             pulumi.StringInput                 `pulumi:"status"`
	Type               pulumi.StringInput                 `pulumi:"type"`
}

func (AccountEncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryptionResponse)(nil)).Elem()
}

func (i AccountEncryptionResponseArgs) ToAccountEncryptionResponseOutput() AccountEncryptionResponseOutput {
	return i.ToAccountEncryptionResponseOutputWithContext(context.Background())
}

func (i AccountEncryptionResponseArgs) ToAccountEncryptionResponseOutputWithContext(ctx context.Context) AccountEncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionResponseOutput)
}

func (i AccountEncryptionResponseArgs) ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput {
	return i.ToAccountEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i AccountEncryptionResponseArgs) ToAccountEncryptionResponsePtrOutputWithContext(ctx context.Context) AccountEncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionResponseOutput).ToAccountEncryptionResponsePtrOutputWithContext(ctx)
}









type AccountEncryptionResponsePtrInput interface {
	pulumi.Input

	ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput
	ToAccountEncryptionResponsePtrOutputWithContext(context.Context) AccountEncryptionResponsePtrOutput
}

type accountEncryptionResponsePtrType AccountEncryptionResponseArgs

func AccountEncryptionResponsePtr(v *AccountEncryptionResponseArgs) AccountEncryptionResponsePtrInput {
	return (*accountEncryptionResponsePtrType)(v)
}

func (*accountEncryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryptionResponse)(nil)).Elem()
}

func (i *accountEncryptionResponsePtrType) ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput {
	return i.ToAccountEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *accountEncryptionResponsePtrType) ToAccountEncryptionResponsePtrOutputWithContext(ctx context.Context) AccountEncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountEncryptionResponsePtrOutput)
}

type AccountEncryptionResponseOutput struct{ *pulumi.OutputState }

func (AccountEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryptionResponse)(nil)).Elem()
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponseOutput() AccountEncryptionResponseOutput {
	return o
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponseOutputWithContext(ctx context.Context) AccountEncryptionResponseOutput {
	return o
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput {
	return o.ToAccountEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o AccountEncryptionResponseOutput) ToAccountEncryptionResponsePtrOutputWithContext(ctx context.Context) AccountEncryptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountEncryptionResponse) *AccountEncryptionResponse {
		return &v
	}).(AccountEncryptionResponsePtrOutput)
}

func (o AccountEncryptionResponseOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o AccountEncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

func (o AccountEncryptionResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o AccountEncryptionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AccountEncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (AccountEncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryptionResponse)(nil)).Elem()
}

func (o AccountEncryptionResponsePtrOutput) ToAccountEncryptionResponsePtrOutput() AccountEncryptionResponsePtrOutput {
	return o
}

func (o AccountEncryptionResponsePtrOutput) ToAccountEncryptionResponsePtrOutputWithContext(ctx context.Context) AccountEncryptionResponsePtrOutput {
	return o
}

func (o AccountEncryptionResponsePtrOutput) Elem() AccountEncryptionResponseOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) AccountEncryptionResponse {
		if v != nil {
			return *v
		}
		var ret AccountEncryptionResponse
		return ret
	}).(AccountEncryptionResponseOutput)
}

func (o AccountEncryptionResponsePtrOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *ResourceIdentityResponse {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(ResourceIdentityResponsePtrOutput)
}

func (o AccountEncryptionResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o AccountEncryptionResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o AccountEncryptionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type AudioEncoderAac struct {
	BitrateKbps *string `pulumi:"bitrateKbps"`
	Type        string  `pulumi:"type"`
}





type AudioEncoderAacInput interface {
	pulumi.Input

	ToAudioEncoderAacOutput() AudioEncoderAacOutput
	ToAudioEncoderAacOutputWithContext(context.Context) AudioEncoderAacOutput
}

type AudioEncoderAacArgs struct {
	BitrateKbps pulumi.StringPtrInput `pulumi:"bitrateKbps"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (AudioEncoderAacArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioEncoderAac)(nil)).Elem()
}

func (i AudioEncoderAacArgs) ToAudioEncoderAacOutput() AudioEncoderAacOutput {
	return i.ToAudioEncoderAacOutputWithContext(context.Background())
}

func (i AudioEncoderAacArgs) ToAudioEncoderAacOutputWithContext(ctx context.Context) AudioEncoderAacOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioEncoderAacOutput)
}

func (i AudioEncoderAacArgs) ToAudioEncoderAacPtrOutput() AudioEncoderAacPtrOutput {
	return i.ToAudioEncoderAacPtrOutputWithContext(context.Background())
}

func (i AudioEncoderAacArgs) ToAudioEncoderAacPtrOutputWithContext(ctx context.Context) AudioEncoderAacPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioEncoderAacOutput).ToAudioEncoderAacPtrOutputWithContext(ctx)
}









type AudioEncoderAacPtrInput interface {
	pulumi.Input

	ToAudioEncoderAacPtrOutput() AudioEncoderAacPtrOutput
	ToAudioEncoderAacPtrOutputWithContext(context.Context) AudioEncoderAacPtrOutput
}

type audioEncoderAacPtrType AudioEncoderAacArgs

func AudioEncoderAacPtr(v *AudioEncoderAacArgs) AudioEncoderAacPtrInput {
	return (*audioEncoderAacPtrType)(v)
}

func (*audioEncoderAacPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AudioEncoderAac)(nil)).Elem()
}

func (i *audioEncoderAacPtrType) ToAudioEncoderAacPtrOutput() AudioEncoderAacPtrOutput {
	return i.ToAudioEncoderAacPtrOutputWithContext(context.Background())
}

func (i *audioEncoderAacPtrType) ToAudioEncoderAacPtrOutputWithContext(ctx context.Context) AudioEncoderAacPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioEncoderAacPtrOutput)
}

type AudioEncoderAacOutput struct{ *pulumi.OutputState }

func (AudioEncoderAacOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioEncoderAac)(nil)).Elem()
}

func (o AudioEncoderAacOutput) ToAudioEncoderAacOutput() AudioEncoderAacOutput {
	return o
}

func (o AudioEncoderAacOutput) ToAudioEncoderAacOutputWithContext(ctx context.Context) AudioEncoderAacOutput {
	return o
}

func (o AudioEncoderAacOutput) ToAudioEncoderAacPtrOutput() AudioEncoderAacPtrOutput {
	return o.ToAudioEncoderAacPtrOutputWithContext(context.Background())
}

func (o AudioEncoderAacOutput) ToAudioEncoderAacPtrOutputWithContext(ctx context.Context) AudioEncoderAacPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AudioEncoderAac) *AudioEncoderAac {
		return &v
	}).(AudioEncoderAacPtrOutput)
}

func (o AudioEncoderAacOutput) BitrateKbps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioEncoderAac) *string { return v.BitrateKbps }).(pulumi.StringPtrOutput)
}

func (o AudioEncoderAacOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AudioEncoderAac) string { return v.Type }).(pulumi.StringOutput)
}

type AudioEncoderAacPtrOutput struct{ *pulumi.OutputState }

func (AudioEncoderAacPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AudioEncoderAac)(nil)).Elem()
}

func (o AudioEncoderAacPtrOutput) ToAudioEncoderAacPtrOutput() AudioEncoderAacPtrOutput {
	return o
}

func (o AudioEncoderAacPtrOutput) ToAudioEncoderAacPtrOutputWithContext(ctx context.Context) AudioEncoderAacPtrOutput {
	return o
}

func (o AudioEncoderAacPtrOutput) Elem() AudioEncoderAacOutput {
	return o.ApplyT(func(v *AudioEncoderAac) AudioEncoderAac {
		if v != nil {
			return *v
		}
		var ret AudioEncoderAac
		return ret
	}).(AudioEncoderAacOutput)
}

func (o AudioEncoderAacPtrOutput) BitrateKbps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AudioEncoderAac) *string {
		if v == nil {
			return nil
		}
		return v.BitrateKbps
	}).(pulumi.StringPtrOutput)
}

func (o AudioEncoderAacPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AudioEncoderAac) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type AudioEncoderAacResponse struct {
	BitrateKbps *string `pulumi:"bitrateKbps"`
	Type        string  `pulumi:"type"`
}





type AudioEncoderAacResponseInput interface {
	pulumi.Input

	ToAudioEncoderAacResponseOutput() AudioEncoderAacResponseOutput
	ToAudioEncoderAacResponseOutputWithContext(context.Context) AudioEncoderAacResponseOutput
}

type AudioEncoderAacResponseArgs struct {
	BitrateKbps pulumi.StringPtrInput `pulumi:"bitrateKbps"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (AudioEncoderAacResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioEncoderAacResponse)(nil)).Elem()
}

func (i AudioEncoderAacResponseArgs) ToAudioEncoderAacResponseOutput() AudioEncoderAacResponseOutput {
	return i.ToAudioEncoderAacResponseOutputWithContext(context.Background())
}

func (i AudioEncoderAacResponseArgs) ToAudioEncoderAacResponseOutputWithContext(ctx context.Context) AudioEncoderAacResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioEncoderAacResponseOutput)
}

func (i AudioEncoderAacResponseArgs) ToAudioEncoderAacResponsePtrOutput() AudioEncoderAacResponsePtrOutput {
	return i.ToAudioEncoderAacResponsePtrOutputWithContext(context.Background())
}

func (i AudioEncoderAacResponseArgs) ToAudioEncoderAacResponsePtrOutputWithContext(ctx context.Context) AudioEncoderAacResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioEncoderAacResponseOutput).ToAudioEncoderAacResponsePtrOutputWithContext(ctx)
}









type AudioEncoderAacResponsePtrInput interface {
	pulumi.Input

	ToAudioEncoderAacResponsePtrOutput() AudioEncoderAacResponsePtrOutput
	ToAudioEncoderAacResponsePtrOutputWithContext(context.Context) AudioEncoderAacResponsePtrOutput
}

type audioEncoderAacResponsePtrType AudioEncoderAacResponseArgs

func AudioEncoderAacResponsePtr(v *AudioEncoderAacResponseArgs) AudioEncoderAacResponsePtrInput {
	return (*audioEncoderAacResponsePtrType)(v)
}

func (*audioEncoderAacResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AudioEncoderAacResponse)(nil)).Elem()
}

func (i *audioEncoderAacResponsePtrType) ToAudioEncoderAacResponsePtrOutput() AudioEncoderAacResponsePtrOutput {
	return i.ToAudioEncoderAacResponsePtrOutputWithContext(context.Background())
}

func (i *audioEncoderAacResponsePtrType) ToAudioEncoderAacResponsePtrOutputWithContext(ctx context.Context) AudioEncoderAacResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioEncoderAacResponsePtrOutput)
}

type AudioEncoderAacResponseOutput struct{ *pulumi.OutputState }

func (AudioEncoderAacResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioEncoderAacResponse)(nil)).Elem()
}

func (o AudioEncoderAacResponseOutput) ToAudioEncoderAacResponseOutput() AudioEncoderAacResponseOutput {
	return o
}

func (o AudioEncoderAacResponseOutput) ToAudioEncoderAacResponseOutputWithContext(ctx context.Context) AudioEncoderAacResponseOutput {
	return o
}

func (o AudioEncoderAacResponseOutput) ToAudioEncoderAacResponsePtrOutput() AudioEncoderAacResponsePtrOutput {
	return o.ToAudioEncoderAacResponsePtrOutputWithContext(context.Background())
}

func (o AudioEncoderAacResponseOutput) ToAudioEncoderAacResponsePtrOutputWithContext(ctx context.Context) AudioEncoderAacResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AudioEncoderAacResponse) *AudioEncoderAacResponse {
		return &v
	}).(AudioEncoderAacResponsePtrOutput)
}

func (o AudioEncoderAacResponseOutput) BitrateKbps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioEncoderAacResponse) *string { return v.BitrateKbps }).(pulumi.StringPtrOutput)
}

func (o AudioEncoderAacResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AudioEncoderAacResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AudioEncoderAacResponsePtrOutput struct{ *pulumi.OutputState }

func (AudioEncoderAacResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AudioEncoderAacResponse)(nil)).Elem()
}

func (o AudioEncoderAacResponsePtrOutput) ToAudioEncoderAacResponsePtrOutput() AudioEncoderAacResponsePtrOutput {
	return o
}

func (o AudioEncoderAacResponsePtrOutput) ToAudioEncoderAacResponsePtrOutputWithContext(ctx context.Context) AudioEncoderAacResponsePtrOutput {
	return o
}

func (o AudioEncoderAacResponsePtrOutput) Elem() AudioEncoderAacResponseOutput {
	return o.ApplyT(func(v *AudioEncoderAacResponse) AudioEncoderAacResponse {
		if v != nil {
			return *v
		}
		var ret AudioEncoderAacResponse
		return ret
	}).(AudioEncoderAacResponseOutput)
}

func (o AudioEncoderAacResponsePtrOutput) BitrateKbps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AudioEncoderAacResponse) *string {
		if v == nil {
			return nil
		}
		return v.BitrateKbps
	}).(pulumi.StringPtrOutput)
}

func (o AudioEncoderAacResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AudioEncoderAacResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type EccTokenKey struct {
	Alg  string `pulumi:"alg"`
	Kid  string `pulumi:"kid"`
	Type string `pulumi:"type"`
	X    string `pulumi:"x"`
	Y    string `pulumi:"y"`
}





type EccTokenKeyInput interface {
	pulumi.Input

	ToEccTokenKeyOutput() EccTokenKeyOutput
	ToEccTokenKeyOutputWithContext(context.Context) EccTokenKeyOutput
}

type EccTokenKeyArgs struct {
	Alg  pulumi.StringInput `pulumi:"alg"`
	Kid  pulumi.StringInput `pulumi:"kid"`
	Type pulumi.StringInput `pulumi:"type"`
	X    pulumi.StringInput `pulumi:"x"`
	Y    pulumi.StringInput `pulumi:"y"`
}

func (EccTokenKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EccTokenKey)(nil)).Elem()
}

func (i EccTokenKeyArgs) ToEccTokenKeyOutput() EccTokenKeyOutput {
	return i.ToEccTokenKeyOutputWithContext(context.Background())
}

func (i EccTokenKeyArgs) ToEccTokenKeyOutputWithContext(ctx context.Context) EccTokenKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EccTokenKeyOutput)
}

type EccTokenKeyOutput struct{ *pulumi.OutputState }

func (EccTokenKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EccTokenKey)(nil)).Elem()
}

func (o EccTokenKeyOutput) ToEccTokenKeyOutput() EccTokenKeyOutput {
	return o
}

func (o EccTokenKeyOutput) ToEccTokenKeyOutputWithContext(ctx context.Context) EccTokenKeyOutput {
	return o
}

func (o EccTokenKeyOutput) Alg() pulumi.StringOutput {
	return o.ApplyT(func(v EccTokenKey) string { return v.Alg }).(pulumi.StringOutput)
}

func (o EccTokenKeyOutput) Kid() pulumi.StringOutput {
	return o.ApplyT(func(v EccTokenKey) string { return v.Kid }).(pulumi.StringOutput)
}

func (o EccTokenKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EccTokenKey) string { return v.Type }).(pulumi.StringOutput)
}

func (o EccTokenKeyOutput) X() pulumi.StringOutput {
	return o.ApplyT(func(v EccTokenKey) string { return v.X }).(pulumi.StringOutput)
}

func (o EccTokenKeyOutput) Y() pulumi.StringOutput {
	return o.ApplyT(func(v EccTokenKey) string { return v.Y }).(pulumi.StringOutput)
}

type EccTokenKeyResponse struct {
	Alg  string `pulumi:"alg"`
	Kid  string `pulumi:"kid"`
	Type string `pulumi:"type"`
	X    string `pulumi:"x"`
	Y    string `pulumi:"y"`
}





type EccTokenKeyResponseInput interface {
	pulumi.Input

	ToEccTokenKeyResponseOutput() EccTokenKeyResponseOutput
	ToEccTokenKeyResponseOutputWithContext(context.Context) EccTokenKeyResponseOutput
}

type EccTokenKeyResponseArgs struct {
	Alg  pulumi.StringInput `pulumi:"alg"`
	Kid  pulumi.StringInput `pulumi:"kid"`
	Type pulumi.StringInput `pulumi:"type"`
	X    pulumi.StringInput `pulumi:"x"`
	Y    pulumi.StringInput `pulumi:"y"`
}

func (EccTokenKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EccTokenKeyResponse)(nil)).Elem()
}

func (i EccTokenKeyResponseArgs) ToEccTokenKeyResponseOutput() EccTokenKeyResponseOutput {
	return i.ToEccTokenKeyResponseOutputWithContext(context.Background())
}

func (i EccTokenKeyResponseArgs) ToEccTokenKeyResponseOutputWithContext(ctx context.Context) EccTokenKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EccTokenKeyResponseOutput)
}

type EccTokenKeyResponseOutput struct{ *pulumi.OutputState }

func (EccTokenKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EccTokenKeyResponse)(nil)).Elem()
}

func (o EccTokenKeyResponseOutput) ToEccTokenKeyResponseOutput() EccTokenKeyResponseOutput {
	return o
}

func (o EccTokenKeyResponseOutput) ToEccTokenKeyResponseOutputWithContext(ctx context.Context) EccTokenKeyResponseOutput {
	return o
}

func (o EccTokenKeyResponseOutput) Alg() pulumi.StringOutput {
	return o.ApplyT(func(v EccTokenKeyResponse) string { return v.Alg }).(pulumi.StringOutput)
}

func (o EccTokenKeyResponseOutput) Kid() pulumi.StringOutput {
	return o.ApplyT(func(v EccTokenKeyResponse) string { return v.Kid }).(pulumi.StringOutput)
}

func (o EccTokenKeyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EccTokenKeyResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o EccTokenKeyResponseOutput) X() pulumi.StringOutput {
	return o.ApplyT(func(v EccTokenKeyResponse) string { return v.X }).(pulumi.StringOutput)
}

func (o EccTokenKeyResponseOutput) Y() pulumi.StringOutput {
	return o.ApplyT(func(v EccTokenKeyResponse) string { return v.Y }).(pulumi.StringOutput)
}

type EncoderCustomPreset struct {
	AudioEncoder *AudioEncoderAac  `pulumi:"audioEncoder"`
	Type         string            `pulumi:"type"`
	VideoEncoder *VideoEncoderH264 `pulumi:"videoEncoder"`
}





type EncoderCustomPresetInput interface {
	pulumi.Input

	ToEncoderCustomPresetOutput() EncoderCustomPresetOutput
	ToEncoderCustomPresetOutputWithContext(context.Context) EncoderCustomPresetOutput
}

type EncoderCustomPresetArgs struct {
	AudioEncoder AudioEncoderAacPtrInput  `pulumi:"audioEncoder"`
	Type         pulumi.StringInput       `pulumi:"type"`
	VideoEncoder VideoEncoderH264PtrInput `pulumi:"videoEncoder"`
}

func (EncoderCustomPresetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderCustomPreset)(nil)).Elem()
}

func (i EncoderCustomPresetArgs) ToEncoderCustomPresetOutput() EncoderCustomPresetOutput {
	return i.ToEncoderCustomPresetOutputWithContext(context.Background())
}

func (i EncoderCustomPresetArgs) ToEncoderCustomPresetOutputWithContext(ctx context.Context) EncoderCustomPresetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncoderCustomPresetOutput)
}

type EncoderCustomPresetOutput struct{ *pulumi.OutputState }

func (EncoderCustomPresetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderCustomPreset)(nil)).Elem()
}

func (o EncoderCustomPresetOutput) ToEncoderCustomPresetOutput() EncoderCustomPresetOutput {
	return o
}

func (o EncoderCustomPresetOutput) ToEncoderCustomPresetOutputWithContext(ctx context.Context) EncoderCustomPresetOutput {
	return o
}

func (o EncoderCustomPresetOutput) AudioEncoder() AudioEncoderAacPtrOutput {
	return o.ApplyT(func(v EncoderCustomPreset) *AudioEncoderAac { return v.AudioEncoder }).(AudioEncoderAacPtrOutput)
}

func (o EncoderCustomPresetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderCustomPreset) string { return v.Type }).(pulumi.StringOutput)
}

func (o EncoderCustomPresetOutput) VideoEncoder() VideoEncoderH264PtrOutput {
	return o.ApplyT(func(v EncoderCustomPreset) *VideoEncoderH264 { return v.VideoEncoder }).(VideoEncoderH264PtrOutput)
}

type EncoderCustomPresetResponse struct {
	AudioEncoder *AudioEncoderAacResponse  `pulumi:"audioEncoder"`
	Type         string                    `pulumi:"type"`
	VideoEncoder *VideoEncoderH264Response `pulumi:"videoEncoder"`
}





type EncoderCustomPresetResponseInput interface {
	pulumi.Input

	ToEncoderCustomPresetResponseOutput() EncoderCustomPresetResponseOutput
	ToEncoderCustomPresetResponseOutputWithContext(context.Context) EncoderCustomPresetResponseOutput
}

type EncoderCustomPresetResponseArgs struct {
	AudioEncoder AudioEncoderAacResponsePtrInput  `pulumi:"audioEncoder"`
	Type         pulumi.StringInput               `pulumi:"type"`
	VideoEncoder VideoEncoderH264ResponsePtrInput `pulumi:"videoEncoder"`
}

func (EncoderCustomPresetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderCustomPresetResponse)(nil)).Elem()
}

func (i EncoderCustomPresetResponseArgs) ToEncoderCustomPresetResponseOutput() EncoderCustomPresetResponseOutput {
	return i.ToEncoderCustomPresetResponseOutputWithContext(context.Background())
}

func (i EncoderCustomPresetResponseArgs) ToEncoderCustomPresetResponseOutputWithContext(ctx context.Context) EncoderCustomPresetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncoderCustomPresetResponseOutput)
}

type EncoderCustomPresetResponseOutput struct{ *pulumi.OutputState }

func (EncoderCustomPresetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderCustomPresetResponse)(nil)).Elem()
}

func (o EncoderCustomPresetResponseOutput) ToEncoderCustomPresetResponseOutput() EncoderCustomPresetResponseOutput {
	return o
}

func (o EncoderCustomPresetResponseOutput) ToEncoderCustomPresetResponseOutputWithContext(ctx context.Context) EncoderCustomPresetResponseOutput {
	return o
}

func (o EncoderCustomPresetResponseOutput) AudioEncoder() AudioEncoderAacResponsePtrOutput {
	return o.ApplyT(func(v EncoderCustomPresetResponse) *AudioEncoderAacResponse { return v.AudioEncoder }).(AudioEncoderAacResponsePtrOutput)
}

func (o EncoderCustomPresetResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderCustomPresetResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o EncoderCustomPresetResponseOutput) VideoEncoder() VideoEncoderH264ResponsePtrOutput {
	return o.ApplyT(func(v EncoderCustomPresetResponse) *VideoEncoderH264Response { return v.VideoEncoder }).(VideoEncoderH264ResponsePtrOutput)
}

type EncoderProcessor struct {
	Inputs []NodeInput `pulumi:"inputs"`
	Name   string      `pulumi:"name"`
	Preset interface{} `pulumi:"preset"`
	Type   string      `pulumi:"type"`
}





type EncoderProcessorInput interface {
	pulumi.Input

	ToEncoderProcessorOutput() EncoderProcessorOutput
	ToEncoderProcessorOutputWithContext(context.Context) EncoderProcessorOutput
}

type EncoderProcessorArgs struct {
	Inputs NodeInputArrayInput `pulumi:"inputs"`
	Name   pulumi.StringInput  `pulumi:"name"`
	Preset pulumi.Input        `pulumi:"preset"`
	Type   pulumi.StringInput  `pulumi:"type"`
}

func (EncoderProcessorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderProcessor)(nil)).Elem()
}

func (i EncoderProcessorArgs) ToEncoderProcessorOutput() EncoderProcessorOutput {
	return i.ToEncoderProcessorOutputWithContext(context.Background())
}

func (i EncoderProcessorArgs) ToEncoderProcessorOutputWithContext(ctx context.Context) EncoderProcessorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncoderProcessorOutput)
}





type EncoderProcessorArrayInput interface {
	pulumi.Input

	ToEncoderProcessorArrayOutput() EncoderProcessorArrayOutput
	ToEncoderProcessorArrayOutputWithContext(context.Context) EncoderProcessorArrayOutput
}

type EncoderProcessorArray []EncoderProcessorInput

func (EncoderProcessorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EncoderProcessor)(nil)).Elem()
}

func (i EncoderProcessorArray) ToEncoderProcessorArrayOutput() EncoderProcessorArrayOutput {
	return i.ToEncoderProcessorArrayOutputWithContext(context.Background())
}

func (i EncoderProcessorArray) ToEncoderProcessorArrayOutputWithContext(ctx context.Context) EncoderProcessorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncoderProcessorArrayOutput)
}

type EncoderProcessorOutput struct{ *pulumi.OutputState }

func (EncoderProcessorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderProcessor)(nil)).Elem()
}

func (o EncoderProcessorOutput) ToEncoderProcessorOutput() EncoderProcessorOutput {
	return o
}

func (o EncoderProcessorOutput) ToEncoderProcessorOutputWithContext(ctx context.Context) EncoderProcessorOutput {
	return o
}

func (o EncoderProcessorOutput) Inputs() NodeInputArrayOutput {
	return o.ApplyT(func(v EncoderProcessor) []NodeInput { return v.Inputs }).(NodeInputArrayOutput)
}

func (o EncoderProcessorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderProcessor) string { return v.Name }).(pulumi.StringOutput)
}

func (o EncoderProcessorOutput) Preset() pulumi.AnyOutput {
	return o.ApplyT(func(v EncoderProcessor) interface{} { return v.Preset }).(pulumi.AnyOutput)
}

func (o EncoderProcessorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderProcessor) string { return v.Type }).(pulumi.StringOutput)
}

type EncoderProcessorArrayOutput struct{ *pulumi.OutputState }

func (EncoderProcessorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EncoderProcessor)(nil)).Elem()
}

func (o EncoderProcessorArrayOutput) ToEncoderProcessorArrayOutput() EncoderProcessorArrayOutput {
	return o
}

func (o EncoderProcessorArrayOutput) ToEncoderProcessorArrayOutputWithContext(ctx context.Context) EncoderProcessorArrayOutput {
	return o
}

func (o EncoderProcessorArrayOutput) Index(i pulumi.IntInput) EncoderProcessorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EncoderProcessor {
		return vs[0].([]EncoderProcessor)[vs[1].(int)]
	}).(EncoderProcessorOutput)
}

type EncoderProcessorResponse struct {
	Inputs []NodeInputResponse `pulumi:"inputs"`
	Name   string              `pulumi:"name"`
	Preset interface{}         `pulumi:"preset"`
	Type   string              `pulumi:"type"`
}





type EncoderProcessorResponseInput interface {
	pulumi.Input

	ToEncoderProcessorResponseOutput() EncoderProcessorResponseOutput
	ToEncoderProcessorResponseOutputWithContext(context.Context) EncoderProcessorResponseOutput
}

type EncoderProcessorResponseArgs struct {
	Inputs NodeInputResponseArrayInput `pulumi:"inputs"`
	Name   pulumi.StringInput          `pulumi:"name"`
	Preset pulumi.Input                `pulumi:"preset"`
	Type   pulumi.StringInput          `pulumi:"type"`
}

func (EncoderProcessorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderProcessorResponse)(nil)).Elem()
}

func (i EncoderProcessorResponseArgs) ToEncoderProcessorResponseOutput() EncoderProcessorResponseOutput {
	return i.ToEncoderProcessorResponseOutputWithContext(context.Background())
}

func (i EncoderProcessorResponseArgs) ToEncoderProcessorResponseOutputWithContext(ctx context.Context) EncoderProcessorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncoderProcessorResponseOutput)
}





type EncoderProcessorResponseArrayInput interface {
	pulumi.Input

	ToEncoderProcessorResponseArrayOutput() EncoderProcessorResponseArrayOutput
	ToEncoderProcessorResponseArrayOutputWithContext(context.Context) EncoderProcessorResponseArrayOutput
}

type EncoderProcessorResponseArray []EncoderProcessorResponseInput

func (EncoderProcessorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EncoderProcessorResponse)(nil)).Elem()
}

func (i EncoderProcessorResponseArray) ToEncoderProcessorResponseArrayOutput() EncoderProcessorResponseArrayOutput {
	return i.ToEncoderProcessorResponseArrayOutputWithContext(context.Background())
}

func (i EncoderProcessorResponseArray) ToEncoderProcessorResponseArrayOutputWithContext(ctx context.Context) EncoderProcessorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncoderProcessorResponseArrayOutput)
}

type EncoderProcessorResponseOutput struct{ *pulumi.OutputState }

func (EncoderProcessorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderProcessorResponse)(nil)).Elem()
}

func (o EncoderProcessorResponseOutput) ToEncoderProcessorResponseOutput() EncoderProcessorResponseOutput {
	return o
}

func (o EncoderProcessorResponseOutput) ToEncoderProcessorResponseOutputWithContext(ctx context.Context) EncoderProcessorResponseOutput {
	return o
}

func (o EncoderProcessorResponseOutput) Inputs() NodeInputResponseArrayOutput {
	return o.ApplyT(func(v EncoderProcessorResponse) []NodeInputResponse { return v.Inputs }).(NodeInputResponseArrayOutput)
}

func (o EncoderProcessorResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderProcessorResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o EncoderProcessorResponseOutput) Preset() pulumi.AnyOutput {
	return o.ApplyT(func(v EncoderProcessorResponse) interface{} { return v.Preset }).(pulumi.AnyOutput)
}

func (o EncoderProcessorResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderProcessorResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EncoderProcessorResponseArrayOutput struct{ *pulumi.OutputState }

func (EncoderProcessorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EncoderProcessorResponse)(nil)).Elem()
}

func (o EncoderProcessorResponseArrayOutput) ToEncoderProcessorResponseArrayOutput() EncoderProcessorResponseArrayOutput {
	return o
}

func (o EncoderProcessorResponseArrayOutput) ToEncoderProcessorResponseArrayOutputWithContext(ctx context.Context) EncoderProcessorResponseArrayOutput {
	return o
}

func (o EncoderProcessorResponseArrayOutput) Index(i pulumi.IntInput) EncoderProcessorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EncoderProcessorResponse {
		return vs[0].([]EncoderProcessorResponse)[vs[1].(int)]
	}).(EncoderProcessorResponseOutput)
}

type EncoderSystemPreset struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type EncoderSystemPresetInput interface {
	pulumi.Input

	ToEncoderSystemPresetOutput() EncoderSystemPresetOutput
	ToEncoderSystemPresetOutputWithContext(context.Context) EncoderSystemPresetOutput
}

type EncoderSystemPresetArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (EncoderSystemPresetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderSystemPreset)(nil)).Elem()
}

func (i EncoderSystemPresetArgs) ToEncoderSystemPresetOutput() EncoderSystemPresetOutput {
	return i.ToEncoderSystemPresetOutputWithContext(context.Background())
}

func (i EncoderSystemPresetArgs) ToEncoderSystemPresetOutputWithContext(ctx context.Context) EncoderSystemPresetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncoderSystemPresetOutput)
}

type EncoderSystemPresetOutput struct{ *pulumi.OutputState }

func (EncoderSystemPresetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderSystemPreset)(nil)).Elem()
}

func (o EncoderSystemPresetOutput) ToEncoderSystemPresetOutput() EncoderSystemPresetOutput {
	return o
}

func (o EncoderSystemPresetOutput) ToEncoderSystemPresetOutputWithContext(ctx context.Context) EncoderSystemPresetOutput {
	return o
}

func (o EncoderSystemPresetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderSystemPreset) string { return v.Name }).(pulumi.StringOutput)
}

func (o EncoderSystemPresetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderSystemPreset) string { return v.Type }).(pulumi.StringOutput)
}

type EncoderSystemPresetResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type EncoderSystemPresetResponseInput interface {
	pulumi.Input

	ToEncoderSystemPresetResponseOutput() EncoderSystemPresetResponseOutput
	ToEncoderSystemPresetResponseOutputWithContext(context.Context) EncoderSystemPresetResponseOutput
}

type EncoderSystemPresetResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (EncoderSystemPresetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderSystemPresetResponse)(nil)).Elem()
}

func (i EncoderSystemPresetResponseArgs) ToEncoderSystemPresetResponseOutput() EncoderSystemPresetResponseOutput {
	return i.ToEncoderSystemPresetResponseOutputWithContext(context.Background())
}

func (i EncoderSystemPresetResponseArgs) ToEncoderSystemPresetResponseOutputWithContext(ctx context.Context) EncoderSystemPresetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncoderSystemPresetResponseOutput)
}

type EncoderSystemPresetResponseOutput struct{ *pulumi.OutputState }

func (EncoderSystemPresetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderSystemPresetResponse)(nil)).Elem()
}

func (o EncoderSystemPresetResponseOutput) ToEncoderSystemPresetResponseOutput() EncoderSystemPresetResponseOutput {
	return o
}

func (o EncoderSystemPresetResponseOutput) ToEncoderSystemPresetResponseOutputWithContext(ctx context.Context) EncoderSystemPresetResponseOutput {
	return o
}

func (o EncoderSystemPresetResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderSystemPresetResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o EncoderSystemPresetResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncoderSystemPresetResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EndpointResponse struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
	Type        string  `pulumi:"type"`
}





type EndpointResponseInput interface {
	pulumi.Input

	ToEndpointResponseOutput() EndpointResponseOutput
	ToEndpointResponseOutputWithContext(context.Context) EndpointResponseOutput
}

type EndpointResponseArgs struct {
	EndpointUrl pulumi.StringPtrInput `pulumi:"endpointUrl"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (EndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointResponse)(nil)).Elem()
}

func (i EndpointResponseArgs) ToEndpointResponseOutput() EndpointResponseOutput {
	return i.ToEndpointResponseOutputWithContext(context.Background())
}

func (i EndpointResponseArgs) ToEndpointResponseOutputWithContext(ctx context.Context) EndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointResponseOutput)
}





type EndpointResponseArrayInput interface {
	pulumi.Input

	ToEndpointResponseArrayOutput() EndpointResponseArrayOutput
	ToEndpointResponseArrayOutputWithContext(context.Context) EndpointResponseArrayOutput
}

type EndpointResponseArray []EndpointResponseInput

func (EndpointResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointResponse)(nil)).Elem()
}

func (i EndpointResponseArray) ToEndpointResponseArrayOutput() EndpointResponseArrayOutput {
	return i.ToEndpointResponseArrayOutputWithContext(context.Background())
}

func (i EndpointResponseArray) ToEndpointResponseArrayOutputWithContext(ctx context.Context) EndpointResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointResponseArrayOutput)
}

type EndpointResponseOutput struct{ *pulumi.OutputState }

func (EndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointResponse)(nil)).Elem()
}

func (o EndpointResponseOutput) ToEndpointResponseOutput() EndpointResponseOutput {
	return o
}

func (o EndpointResponseOutput) ToEndpointResponseOutputWithContext(ctx context.Context) EndpointResponseOutput {
	return o
}

func (o EndpointResponseOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (EndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointResponse)(nil)).Elem()
}

func (o EndpointResponseArrayOutput) ToEndpointResponseArrayOutput() EndpointResponseArrayOutput {
	return o
}

func (o EndpointResponseArrayOutput) ToEndpointResponseArrayOutputWithContext(ctx context.Context) EndpointResponseArrayOutput {
	return o
}

func (o EndpointResponseArrayOutput) Index(i pulumi.IntInput) EndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointResponse {
		return vs[0].([]EndpointResponse)[vs[1].(int)]
	}).(EndpointResponseOutput)
}

type JwtAuthentication struct {
	Audiences []string      `pulumi:"audiences"`
	Claims    []TokenClaim  `pulumi:"claims"`
	Issuers   []string      `pulumi:"issuers"`
	Keys      []interface{} `pulumi:"keys"`
	Type      string        `pulumi:"type"`
}





type JwtAuthenticationInput interface {
	pulumi.Input

	ToJwtAuthenticationOutput() JwtAuthenticationOutput
	ToJwtAuthenticationOutputWithContext(context.Context) JwtAuthenticationOutput
}

type JwtAuthenticationArgs struct {
	Audiences pulumi.StringArrayInput `pulumi:"audiences"`
	Claims    TokenClaimArrayInput    `pulumi:"claims"`
	Issuers   pulumi.StringArrayInput `pulumi:"issuers"`
	Keys      pulumi.ArrayInput       `pulumi:"keys"`
	Type      pulumi.StringInput      `pulumi:"type"`
}

func (JwtAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtAuthentication)(nil)).Elem()
}

func (i JwtAuthenticationArgs) ToJwtAuthenticationOutput() JwtAuthenticationOutput {
	return i.ToJwtAuthenticationOutputWithContext(context.Background())
}

func (i JwtAuthenticationArgs) ToJwtAuthenticationOutputWithContext(ctx context.Context) JwtAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtAuthenticationOutput)
}

func (i JwtAuthenticationArgs) ToJwtAuthenticationPtrOutput() JwtAuthenticationPtrOutput {
	return i.ToJwtAuthenticationPtrOutputWithContext(context.Background())
}

func (i JwtAuthenticationArgs) ToJwtAuthenticationPtrOutputWithContext(ctx context.Context) JwtAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtAuthenticationOutput).ToJwtAuthenticationPtrOutputWithContext(ctx)
}









type JwtAuthenticationPtrInput interface {
	pulumi.Input

	ToJwtAuthenticationPtrOutput() JwtAuthenticationPtrOutput
	ToJwtAuthenticationPtrOutputWithContext(context.Context) JwtAuthenticationPtrOutput
}

type jwtAuthenticationPtrType JwtAuthenticationArgs

func JwtAuthenticationPtr(v *JwtAuthenticationArgs) JwtAuthenticationPtrInput {
	return (*jwtAuthenticationPtrType)(v)
}

func (*jwtAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtAuthentication)(nil)).Elem()
}

func (i *jwtAuthenticationPtrType) ToJwtAuthenticationPtrOutput() JwtAuthenticationPtrOutput {
	return i.ToJwtAuthenticationPtrOutputWithContext(context.Background())
}

func (i *jwtAuthenticationPtrType) ToJwtAuthenticationPtrOutputWithContext(ctx context.Context) JwtAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtAuthenticationPtrOutput)
}

type JwtAuthenticationOutput struct{ *pulumi.OutputState }

func (JwtAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtAuthentication)(nil)).Elem()
}

func (o JwtAuthenticationOutput) ToJwtAuthenticationOutput() JwtAuthenticationOutput {
	return o
}

func (o JwtAuthenticationOutput) ToJwtAuthenticationOutputWithContext(ctx context.Context) JwtAuthenticationOutput {
	return o
}

func (o JwtAuthenticationOutput) ToJwtAuthenticationPtrOutput() JwtAuthenticationPtrOutput {
	return o.ToJwtAuthenticationPtrOutputWithContext(context.Background())
}

func (o JwtAuthenticationOutput) ToJwtAuthenticationPtrOutputWithContext(ctx context.Context) JwtAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JwtAuthentication) *JwtAuthentication {
		return &v
	}).(JwtAuthenticationPtrOutput)
}

func (o JwtAuthenticationOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtAuthentication) []string { return v.Audiences }).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationOutput) Claims() TokenClaimArrayOutput {
	return o.ApplyT(func(v JwtAuthentication) []TokenClaim { return v.Claims }).(TokenClaimArrayOutput)
}

func (o JwtAuthenticationOutput) Issuers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtAuthentication) []string { return v.Issuers }).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationOutput) Keys() pulumi.ArrayOutput {
	return o.ApplyT(func(v JwtAuthentication) []interface{} { return v.Keys }).(pulumi.ArrayOutput)
}

func (o JwtAuthenticationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JwtAuthentication) string { return v.Type }).(pulumi.StringOutput)
}

type JwtAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (JwtAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtAuthentication)(nil)).Elem()
}

func (o JwtAuthenticationPtrOutput) ToJwtAuthenticationPtrOutput() JwtAuthenticationPtrOutput {
	return o
}

func (o JwtAuthenticationPtrOutput) ToJwtAuthenticationPtrOutputWithContext(ctx context.Context) JwtAuthenticationPtrOutput {
	return o
}

func (o JwtAuthenticationPtrOutput) Elem() JwtAuthenticationOutput {
	return o.ApplyT(func(v *JwtAuthentication) JwtAuthentication {
		if v != nil {
			return *v
		}
		var ret JwtAuthentication
		return ret
	}).(JwtAuthenticationOutput)
}

func (o JwtAuthenticationPtrOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtAuthentication) []string {
		if v == nil {
			return nil
		}
		return v.Audiences
	}).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationPtrOutput) Claims() TokenClaimArrayOutput {
	return o.ApplyT(func(v *JwtAuthentication) []TokenClaim {
		if v == nil {
			return nil
		}
		return v.Claims
	}).(TokenClaimArrayOutput)
}

func (o JwtAuthenticationPtrOutput) Issuers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtAuthentication) []string {
		if v == nil {
			return nil
		}
		return v.Issuers
	}).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationPtrOutput) Keys() pulumi.ArrayOutput {
	return o.ApplyT(func(v *JwtAuthentication) []interface{} {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(pulumi.ArrayOutput)
}

func (o JwtAuthenticationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JwtAuthentication) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type JwtAuthenticationResponse struct {
	Audiences []string             `pulumi:"audiences"`
	Claims    []TokenClaimResponse `pulumi:"claims"`
	Issuers   []string             `pulumi:"issuers"`
	Keys      []interface{}        `pulumi:"keys"`
	Type      string               `pulumi:"type"`
}





type JwtAuthenticationResponseInput interface {
	pulumi.Input

	ToJwtAuthenticationResponseOutput() JwtAuthenticationResponseOutput
	ToJwtAuthenticationResponseOutputWithContext(context.Context) JwtAuthenticationResponseOutput
}

type JwtAuthenticationResponseArgs struct {
	Audiences pulumi.StringArrayInput      `pulumi:"audiences"`
	Claims    TokenClaimResponseArrayInput `pulumi:"claims"`
	Issuers   pulumi.StringArrayInput      `pulumi:"issuers"`
	Keys      pulumi.ArrayInput            `pulumi:"keys"`
	Type      pulumi.StringInput           `pulumi:"type"`
}

func (JwtAuthenticationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtAuthenticationResponse)(nil)).Elem()
}

func (i JwtAuthenticationResponseArgs) ToJwtAuthenticationResponseOutput() JwtAuthenticationResponseOutput {
	return i.ToJwtAuthenticationResponseOutputWithContext(context.Background())
}

func (i JwtAuthenticationResponseArgs) ToJwtAuthenticationResponseOutputWithContext(ctx context.Context) JwtAuthenticationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtAuthenticationResponseOutput)
}

func (i JwtAuthenticationResponseArgs) ToJwtAuthenticationResponsePtrOutput() JwtAuthenticationResponsePtrOutput {
	return i.ToJwtAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i JwtAuthenticationResponseArgs) ToJwtAuthenticationResponsePtrOutputWithContext(ctx context.Context) JwtAuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtAuthenticationResponseOutput).ToJwtAuthenticationResponsePtrOutputWithContext(ctx)
}









type JwtAuthenticationResponsePtrInput interface {
	pulumi.Input

	ToJwtAuthenticationResponsePtrOutput() JwtAuthenticationResponsePtrOutput
	ToJwtAuthenticationResponsePtrOutputWithContext(context.Context) JwtAuthenticationResponsePtrOutput
}

type jwtAuthenticationResponsePtrType JwtAuthenticationResponseArgs

func JwtAuthenticationResponsePtr(v *JwtAuthenticationResponseArgs) JwtAuthenticationResponsePtrInput {
	return (*jwtAuthenticationResponsePtrType)(v)
}

func (*jwtAuthenticationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtAuthenticationResponse)(nil)).Elem()
}

func (i *jwtAuthenticationResponsePtrType) ToJwtAuthenticationResponsePtrOutput() JwtAuthenticationResponsePtrOutput {
	return i.ToJwtAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (i *jwtAuthenticationResponsePtrType) ToJwtAuthenticationResponsePtrOutputWithContext(ctx context.Context) JwtAuthenticationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtAuthenticationResponsePtrOutput)
}

type JwtAuthenticationResponseOutput struct{ *pulumi.OutputState }

func (JwtAuthenticationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtAuthenticationResponse)(nil)).Elem()
}

func (o JwtAuthenticationResponseOutput) ToJwtAuthenticationResponseOutput() JwtAuthenticationResponseOutput {
	return o
}

func (o JwtAuthenticationResponseOutput) ToJwtAuthenticationResponseOutputWithContext(ctx context.Context) JwtAuthenticationResponseOutput {
	return o
}

func (o JwtAuthenticationResponseOutput) ToJwtAuthenticationResponsePtrOutput() JwtAuthenticationResponsePtrOutput {
	return o.ToJwtAuthenticationResponsePtrOutputWithContext(context.Background())
}

func (o JwtAuthenticationResponseOutput) ToJwtAuthenticationResponsePtrOutputWithContext(ctx context.Context) JwtAuthenticationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JwtAuthenticationResponse) *JwtAuthenticationResponse {
		return &v
	}).(JwtAuthenticationResponsePtrOutput)
}

func (o JwtAuthenticationResponseOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtAuthenticationResponse) []string { return v.Audiences }).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationResponseOutput) Claims() TokenClaimResponseArrayOutput {
	return o.ApplyT(func(v JwtAuthenticationResponse) []TokenClaimResponse { return v.Claims }).(TokenClaimResponseArrayOutput)
}

func (o JwtAuthenticationResponseOutput) Issuers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtAuthenticationResponse) []string { return v.Issuers }).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationResponseOutput) Keys() pulumi.ArrayOutput {
	return o.ApplyT(func(v JwtAuthenticationResponse) []interface{} { return v.Keys }).(pulumi.ArrayOutput)
}

func (o JwtAuthenticationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JwtAuthenticationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type JwtAuthenticationResponsePtrOutput struct{ *pulumi.OutputState }

func (JwtAuthenticationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtAuthenticationResponse)(nil)).Elem()
}

func (o JwtAuthenticationResponsePtrOutput) ToJwtAuthenticationResponsePtrOutput() JwtAuthenticationResponsePtrOutput {
	return o
}

func (o JwtAuthenticationResponsePtrOutput) ToJwtAuthenticationResponsePtrOutputWithContext(ctx context.Context) JwtAuthenticationResponsePtrOutput {
	return o
}

func (o JwtAuthenticationResponsePtrOutput) Elem() JwtAuthenticationResponseOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) JwtAuthenticationResponse {
		if v != nil {
			return *v
		}
		var ret JwtAuthenticationResponse
		return ret
	}).(JwtAuthenticationResponseOutput)
}

func (o JwtAuthenticationResponsePtrOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Audiences
	}).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationResponsePtrOutput) Claims() TokenClaimResponseArrayOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) []TokenClaimResponse {
		if v == nil {
			return nil
		}
		return v.Claims
	}).(TokenClaimResponseArrayOutput)
}

func (o JwtAuthenticationResponsePtrOutput) Issuers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Issuers
	}).(pulumi.StringArrayOutput)
}

func (o JwtAuthenticationResponsePtrOutput) Keys() pulumi.ArrayOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) []interface{} {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(pulumi.ArrayOutput)
}

func (o JwtAuthenticationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JwtAuthenticationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type KeyVaultProperties struct {
	KeyIdentifier string `pulumi:"keyIdentifier"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyIdentifier pulumi.StringInput `pulumi:"keyIdentifier"`
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

func (o KeyVaultPropertiesOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultProperties) string { return v.KeyIdentifier }).(pulumi.StringOutput)
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

func (o KeyVaultPropertiesPtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	CurrentKeyIdentifier string `pulumi:"currentKeyIdentifier"`
	KeyIdentifier        string `pulumi:"keyIdentifier"`
}





type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

type KeyVaultPropertiesResponseArgs struct {
	CurrentKeyIdentifier pulumi.StringInput `pulumi:"currentKeyIdentifier"`
	KeyIdentifier        pulumi.StringInput `pulumi:"keyIdentifier"`
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

func (o KeyVaultPropertiesResponseOutput) CurrentKeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.CurrentKeyIdentifier }).(pulumi.StringOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.KeyIdentifier }).(pulumi.StringOutput)
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

func (o KeyVaultPropertiesResponsePtrOutput) CurrentKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CurrentKeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

type NodeInput struct {
	NodeName string `pulumi:"nodeName"`
}





type NodeInputInput interface {
	pulumi.Input

	ToNodeInputOutput() NodeInputOutput
	ToNodeInputOutputWithContext(context.Context) NodeInputOutput
}

type NodeInputArgs struct {
	NodeName pulumi.StringInput `pulumi:"nodeName"`
}

func (NodeInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeInput)(nil)).Elem()
}

func (i NodeInputArgs) ToNodeInputOutput() NodeInputOutput {
	return i.ToNodeInputOutputWithContext(context.Background())
}

func (i NodeInputArgs) ToNodeInputOutputWithContext(ctx context.Context) NodeInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeInputOutput)
}





type NodeInputArrayInput interface {
	pulumi.Input

	ToNodeInputArrayOutput() NodeInputArrayOutput
	ToNodeInputArrayOutputWithContext(context.Context) NodeInputArrayOutput
}

type NodeInputArray []NodeInputInput

func (NodeInputArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeInput)(nil)).Elem()
}

func (i NodeInputArray) ToNodeInputArrayOutput() NodeInputArrayOutput {
	return i.ToNodeInputArrayOutputWithContext(context.Background())
}

func (i NodeInputArray) ToNodeInputArrayOutputWithContext(ctx context.Context) NodeInputArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeInputArrayOutput)
}

type NodeInputOutput struct{ *pulumi.OutputState }

func (NodeInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeInput)(nil)).Elem()
}

func (o NodeInputOutput) ToNodeInputOutput() NodeInputOutput {
	return o
}

func (o NodeInputOutput) ToNodeInputOutputWithContext(ctx context.Context) NodeInputOutput {
	return o
}

func (o NodeInputOutput) NodeName() pulumi.StringOutput {
	return o.ApplyT(func(v NodeInput) string { return v.NodeName }).(pulumi.StringOutput)
}

type NodeInputArrayOutput struct{ *pulumi.OutputState }

func (NodeInputArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeInput)(nil)).Elem()
}

func (o NodeInputArrayOutput) ToNodeInputArrayOutput() NodeInputArrayOutput {
	return o
}

func (o NodeInputArrayOutput) ToNodeInputArrayOutputWithContext(ctx context.Context) NodeInputArrayOutput {
	return o
}

func (o NodeInputArrayOutput) Index(i pulumi.IntInput) NodeInputOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeInput {
		return vs[0].([]NodeInput)[vs[1].(int)]
	}).(NodeInputOutput)
}

type NodeInputResponse struct {
	NodeName string `pulumi:"nodeName"`
}





type NodeInputResponseInput interface {
	pulumi.Input

	ToNodeInputResponseOutput() NodeInputResponseOutput
	ToNodeInputResponseOutputWithContext(context.Context) NodeInputResponseOutput
}

type NodeInputResponseArgs struct {
	NodeName pulumi.StringInput `pulumi:"nodeName"`
}

func (NodeInputResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeInputResponse)(nil)).Elem()
}

func (i NodeInputResponseArgs) ToNodeInputResponseOutput() NodeInputResponseOutput {
	return i.ToNodeInputResponseOutputWithContext(context.Background())
}

func (i NodeInputResponseArgs) ToNodeInputResponseOutputWithContext(ctx context.Context) NodeInputResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeInputResponseOutput)
}





type NodeInputResponseArrayInput interface {
	pulumi.Input

	ToNodeInputResponseArrayOutput() NodeInputResponseArrayOutput
	ToNodeInputResponseArrayOutputWithContext(context.Context) NodeInputResponseArrayOutput
}

type NodeInputResponseArray []NodeInputResponseInput

func (NodeInputResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeInputResponse)(nil)).Elem()
}

func (i NodeInputResponseArray) ToNodeInputResponseArrayOutput() NodeInputResponseArrayOutput {
	return i.ToNodeInputResponseArrayOutputWithContext(context.Background())
}

func (i NodeInputResponseArray) ToNodeInputResponseArrayOutputWithContext(ctx context.Context) NodeInputResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeInputResponseArrayOutput)
}

type NodeInputResponseOutput struct{ *pulumi.OutputState }

func (NodeInputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeInputResponse)(nil)).Elem()
}

func (o NodeInputResponseOutput) ToNodeInputResponseOutput() NodeInputResponseOutput {
	return o
}

func (o NodeInputResponseOutput) ToNodeInputResponseOutputWithContext(ctx context.Context) NodeInputResponseOutput {
	return o
}

func (o NodeInputResponseOutput) NodeName() pulumi.StringOutput {
	return o.ApplyT(func(v NodeInputResponse) string { return v.NodeName }).(pulumi.StringOutput)
}

type NodeInputResponseArrayOutput struct{ *pulumi.OutputState }

func (NodeInputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeInputResponse)(nil)).Elem()
}

func (o NodeInputResponseArrayOutput) ToNodeInputResponseArrayOutput() NodeInputResponseArrayOutput {
	return o
}

func (o NodeInputResponseArrayOutput) ToNodeInputResponseArrayOutputWithContext(ctx context.Context) NodeInputResponseArrayOutput {
	return o
}

func (o NodeInputResponseArrayOutput) Index(i pulumi.IntInput) NodeInputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeInputResponse {
		return vs[0].([]NodeInputResponse)[vs[1].(int)]
	}).(NodeInputResponseOutput)
}

type ParameterDeclaration struct {
	Default     *string `pulumi:"default"`
	Description *string `pulumi:"description"`
	Name        string  `pulumi:"name"`
	Type        string  `pulumi:"type"`
}





type ParameterDeclarationInput interface {
	pulumi.Input

	ToParameterDeclarationOutput() ParameterDeclarationOutput
	ToParameterDeclarationOutputWithContext(context.Context) ParameterDeclarationOutput
}

type ParameterDeclarationArgs struct {
	Default     pulumi.StringPtrInput `pulumi:"default"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringInput    `pulumi:"name"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (ParameterDeclarationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDeclaration)(nil)).Elem()
}

func (i ParameterDeclarationArgs) ToParameterDeclarationOutput() ParameterDeclarationOutput {
	return i.ToParameterDeclarationOutputWithContext(context.Background())
}

func (i ParameterDeclarationArgs) ToParameterDeclarationOutputWithContext(ctx context.Context) ParameterDeclarationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDeclarationOutput)
}





type ParameterDeclarationArrayInput interface {
	pulumi.Input

	ToParameterDeclarationArrayOutput() ParameterDeclarationArrayOutput
	ToParameterDeclarationArrayOutputWithContext(context.Context) ParameterDeclarationArrayOutput
}

type ParameterDeclarationArray []ParameterDeclarationInput

func (ParameterDeclarationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDeclaration)(nil)).Elem()
}

func (i ParameterDeclarationArray) ToParameterDeclarationArrayOutput() ParameterDeclarationArrayOutput {
	return i.ToParameterDeclarationArrayOutputWithContext(context.Background())
}

func (i ParameterDeclarationArray) ToParameterDeclarationArrayOutputWithContext(ctx context.Context) ParameterDeclarationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDeclarationArrayOutput)
}

type ParameterDeclarationOutput struct{ *pulumi.OutputState }

func (ParameterDeclarationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDeclaration)(nil)).Elem()
}

func (o ParameterDeclarationOutput) ToParameterDeclarationOutput() ParameterDeclarationOutput {
	return o
}

func (o ParameterDeclarationOutput) ToParameterDeclarationOutputWithContext(ctx context.Context) ParameterDeclarationOutput {
	return o
}

func (o ParameterDeclarationOutput) Default() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDeclaration) *string { return v.Default }).(pulumi.StringPtrOutput)
}

func (o ParameterDeclarationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDeclaration) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDeclarationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDeclaration) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterDeclarationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDeclaration) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterDeclarationArrayOutput struct{ *pulumi.OutputState }

func (ParameterDeclarationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDeclaration)(nil)).Elem()
}

func (o ParameterDeclarationArrayOutput) ToParameterDeclarationArrayOutput() ParameterDeclarationArrayOutput {
	return o
}

func (o ParameterDeclarationArrayOutput) ToParameterDeclarationArrayOutputWithContext(ctx context.Context) ParameterDeclarationArrayOutput {
	return o
}

func (o ParameterDeclarationArrayOutput) Index(i pulumi.IntInput) ParameterDeclarationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterDeclaration {
		return vs[0].([]ParameterDeclaration)[vs[1].(int)]
	}).(ParameterDeclarationOutput)
}

type ParameterDeclarationResponse struct {
	Default     *string `pulumi:"default"`
	Description *string `pulumi:"description"`
	Name        string  `pulumi:"name"`
	Type        string  `pulumi:"type"`
}





type ParameterDeclarationResponseInput interface {
	pulumi.Input

	ToParameterDeclarationResponseOutput() ParameterDeclarationResponseOutput
	ToParameterDeclarationResponseOutputWithContext(context.Context) ParameterDeclarationResponseOutput
}

type ParameterDeclarationResponseArgs struct {
	Default     pulumi.StringPtrInput `pulumi:"default"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringInput    `pulumi:"name"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (ParameterDeclarationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDeclarationResponse)(nil)).Elem()
}

func (i ParameterDeclarationResponseArgs) ToParameterDeclarationResponseOutput() ParameterDeclarationResponseOutput {
	return i.ToParameterDeclarationResponseOutputWithContext(context.Background())
}

func (i ParameterDeclarationResponseArgs) ToParameterDeclarationResponseOutputWithContext(ctx context.Context) ParameterDeclarationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDeclarationResponseOutput)
}





type ParameterDeclarationResponseArrayInput interface {
	pulumi.Input

	ToParameterDeclarationResponseArrayOutput() ParameterDeclarationResponseArrayOutput
	ToParameterDeclarationResponseArrayOutputWithContext(context.Context) ParameterDeclarationResponseArrayOutput
}

type ParameterDeclarationResponseArray []ParameterDeclarationResponseInput

func (ParameterDeclarationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDeclarationResponse)(nil)).Elem()
}

func (i ParameterDeclarationResponseArray) ToParameterDeclarationResponseArrayOutput() ParameterDeclarationResponseArrayOutput {
	return i.ToParameterDeclarationResponseArrayOutputWithContext(context.Background())
}

func (i ParameterDeclarationResponseArray) ToParameterDeclarationResponseArrayOutputWithContext(ctx context.Context) ParameterDeclarationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDeclarationResponseArrayOutput)
}

type ParameterDeclarationResponseOutput struct{ *pulumi.OutputState }

func (ParameterDeclarationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDeclarationResponse)(nil)).Elem()
}

func (o ParameterDeclarationResponseOutput) ToParameterDeclarationResponseOutput() ParameterDeclarationResponseOutput {
	return o
}

func (o ParameterDeclarationResponseOutput) ToParameterDeclarationResponseOutputWithContext(ctx context.Context) ParameterDeclarationResponseOutput {
	return o
}

func (o ParameterDeclarationResponseOutput) Default() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDeclarationResponse) *string { return v.Default }).(pulumi.StringPtrOutput)
}

func (o ParameterDeclarationResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDeclarationResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDeclarationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDeclarationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterDeclarationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDeclarationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterDeclarationResponseArrayOutput struct{ *pulumi.OutputState }

func (ParameterDeclarationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDeclarationResponse)(nil)).Elem()
}

func (o ParameterDeclarationResponseArrayOutput) ToParameterDeclarationResponseArrayOutput() ParameterDeclarationResponseArrayOutput {
	return o
}

func (o ParameterDeclarationResponseArrayOutput) ToParameterDeclarationResponseArrayOutputWithContext(ctx context.Context) ParameterDeclarationResponseArrayOutput {
	return o
}

func (o ParameterDeclarationResponseArrayOutput) Index(i pulumi.IntInput) ParameterDeclarationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterDeclarationResponse {
		return vs[0].([]ParameterDeclarationResponse)[vs[1].(int)]
	}).(ParameterDeclarationResponseOutput)
}

type ParameterDefinition struct {
	Name  string  `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type ParameterDefinitionInput interface {
	pulumi.Input

	ToParameterDefinitionOutput() ParameterDefinitionOutput
	ToParameterDefinitionOutputWithContext(context.Context) ParameterDefinitionOutput
}

type ParameterDefinitionArgs struct {
	Name  pulumi.StringInput    `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ParameterDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinition)(nil)).Elem()
}

func (i ParameterDefinitionArgs) ToParameterDefinitionOutput() ParameterDefinitionOutput {
	return i.ToParameterDefinitionOutputWithContext(context.Background())
}

func (i ParameterDefinitionArgs) ToParameterDefinitionOutputWithContext(ctx context.Context) ParameterDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionOutput)
}





type ParameterDefinitionArrayInput interface {
	pulumi.Input

	ToParameterDefinitionArrayOutput() ParameterDefinitionArrayOutput
	ToParameterDefinitionArrayOutputWithContext(context.Context) ParameterDefinitionArrayOutput
}

type ParameterDefinitionArray []ParameterDefinitionInput

func (ParameterDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDefinition)(nil)).Elem()
}

func (i ParameterDefinitionArray) ToParameterDefinitionArrayOutput() ParameterDefinitionArrayOutput {
	return i.ToParameterDefinitionArrayOutputWithContext(context.Background())
}

func (i ParameterDefinitionArray) ToParameterDefinitionArrayOutputWithContext(ctx context.Context) ParameterDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionArrayOutput)
}

type ParameterDefinitionOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinition)(nil)).Elem()
}

func (o ParameterDefinitionOutput) ToParameterDefinitionOutput() ParameterDefinitionOutput {
	return o
}

func (o ParameterDefinitionOutput) ToParameterDefinitionOutputWithContext(ctx context.Context) ParameterDefinitionOutput {
	return o
}

func (o ParameterDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDefinition) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterDefinitionOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinition) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionArrayOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDefinition)(nil)).Elem()
}

func (o ParameterDefinitionArrayOutput) ToParameterDefinitionArrayOutput() ParameterDefinitionArrayOutput {
	return o
}

func (o ParameterDefinitionArrayOutput) ToParameterDefinitionArrayOutputWithContext(ctx context.Context) ParameterDefinitionArrayOutput {
	return o
}

func (o ParameterDefinitionArrayOutput) Index(i pulumi.IntInput) ParameterDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterDefinition {
		return vs[0].([]ParameterDefinition)[vs[1].(int)]
	}).(ParameterDefinitionOutput)
}

type ParameterDefinitionResponse struct {
	Name  string  `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type ParameterDefinitionResponseInput interface {
	pulumi.Input

	ToParameterDefinitionResponseOutput() ParameterDefinitionResponseOutput
	ToParameterDefinitionResponseOutputWithContext(context.Context) ParameterDefinitionResponseOutput
}

type ParameterDefinitionResponseArgs struct {
	Name  pulumi.StringInput    `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ParameterDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionResponse)(nil)).Elem()
}

func (i ParameterDefinitionResponseArgs) ToParameterDefinitionResponseOutput() ParameterDefinitionResponseOutput {
	return i.ToParameterDefinitionResponseOutputWithContext(context.Background())
}

func (i ParameterDefinitionResponseArgs) ToParameterDefinitionResponseOutputWithContext(ctx context.Context) ParameterDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionResponseOutput)
}





type ParameterDefinitionResponseArrayInput interface {
	pulumi.Input

	ToParameterDefinitionResponseArrayOutput() ParameterDefinitionResponseArrayOutput
	ToParameterDefinitionResponseArrayOutputWithContext(context.Context) ParameterDefinitionResponseArrayOutput
}

type ParameterDefinitionResponseArray []ParameterDefinitionResponseInput

func (ParameterDefinitionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDefinitionResponse)(nil)).Elem()
}

func (i ParameterDefinitionResponseArray) ToParameterDefinitionResponseArrayOutput() ParameterDefinitionResponseArrayOutput {
	return i.ToParameterDefinitionResponseArrayOutputWithContext(context.Background())
}

func (i ParameterDefinitionResponseArray) ToParameterDefinitionResponseArrayOutputWithContext(ctx context.Context) ParameterDefinitionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionResponseArrayOutput)
}

type ParameterDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionResponse)(nil)).Elem()
}

func (o ParameterDefinitionResponseOutput) ToParameterDefinitionResponseOutput() ParameterDefinitionResponseOutput {
	return o
}

func (o ParameterDefinitionResponseOutput) ToParameterDefinitionResponseOutputWithContext(ctx context.Context) ParameterDefinitionResponseOutput {
	return o
}

func (o ParameterDefinitionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterDefinitionResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterDefinitionResponse)(nil)).Elem()
}

func (o ParameterDefinitionResponseArrayOutput) ToParameterDefinitionResponseArrayOutput() ParameterDefinitionResponseArrayOutput {
	return o
}

func (o ParameterDefinitionResponseArrayOutput) ToParameterDefinitionResponseArrayOutputWithContext(ctx context.Context) ParameterDefinitionResponseArrayOutput {
	return o
}

func (o ParameterDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ParameterDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterDefinitionResponse {
		return vs[0].([]ParameterDefinitionResponse)[vs[1].(int)]
	}).(ParameterDefinitionResponseOutput)
}

type PemCertificateList struct {
	Certificates []string `pulumi:"certificates"`
	Type         string   `pulumi:"type"`
}





type PemCertificateListInput interface {
	pulumi.Input

	ToPemCertificateListOutput() PemCertificateListOutput
	ToPemCertificateListOutputWithContext(context.Context) PemCertificateListOutput
}

type PemCertificateListArgs struct {
	Certificates pulumi.StringArrayInput `pulumi:"certificates"`
	Type         pulumi.StringInput      `pulumi:"type"`
}

func (PemCertificateListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PemCertificateList)(nil)).Elem()
}

func (i PemCertificateListArgs) ToPemCertificateListOutput() PemCertificateListOutput {
	return i.ToPemCertificateListOutputWithContext(context.Background())
}

func (i PemCertificateListArgs) ToPemCertificateListOutputWithContext(ctx context.Context) PemCertificateListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PemCertificateListOutput)
}

func (i PemCertificateListArgs) ToPemCertificateListPtrOutput() PemCertificateListPtrOutput {
	return i.ToPemCertificateListPtrOutputWithContext(context.Background())
}

func (i PemCertificateListArgs) ToPemCertificateListPtrOutputWithContext(ctx context.Context) PemCertificateListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PemCertificateListOutput).ToPemCertificateListPtrOutputWithContext(ctx)
}









type PemCertificateListPtrInput interface {
	pulumi.Input

	ToPemCertificateListPtrOutput() PemCertificateListPtrOutput
	ToPemCertificateListPtrOutputWithContext(context.Context) PemCertificateListPtrOutput
}

type pemCertificateListPtrType PemCertificateListArgs

func PemCertificateListPtr(v *PemCertificateListArgs) PemCertificateListPtrInput {
	return (*pemCertificateListPtrType)(v)
}

func (*pemCertificateListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PemCertificateList)(nil)).Elem()
}

func (i *pemCertificateListPtrType) ToPemCertificateListPtrOutput() PemCertificateListPtrOutput {
	return i.ToPemCertificateListPtrOutputWithContext(context.Background())
}

func (i *pemCertificateListPtrType) ToPemCertificateListPtrOutputWithContext(ctx context.Context) PemCertificateListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PemCertificateListPtrOutput)
}

type PemCertificateListOutput struct{ *pulumi.OutputState }

func (PemCertificateListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PemCertificateList)(nil)).Elem()
}

func (o PemCertificateListOutput) ToPemCertificateListOutput() PemCertificateListOutput {
	return o
}

func (o PemCertificateListOutput) ToPemCertificateListOutputWithContext(ctx context.Context) PemCertificateListOutput {
	return o
}

func (o PemCertificateListOutput) ToPemCertificateListPtrOutput() PemCertificateListPtrOutput {
	return o.ToPemCertificateListPtrOutputWithContext(context.Background())
}

func (o PemCertificateListOutput) ToPemCertificateListPtrOutputWithContext(ctx context.Context) PemCertificateListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PemCertificateList) *PemCertificateList {
		return &v
	}).(PemCertificateListPtrOutput)
}

func (o PemCertificateListOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PemCertificateList) []string { return v.Certificates }).(pulumi.StringArrayOutput)
}

func (o PemCertificateListOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PemCertificateList) string { return v.Type }).(pulumi.StringOutput)
}

type PemCertificateListPtrOutput struct{ *pulumi.OutputState }

func (PemCertificateListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PemCertificateList)(nil)).Elem()
}

func (o PemCertificateListPtrOutput) ToPemCertificateListPtrOutput() PemCertificateListPtrOutput {
	return o
}

func (o PemCertificateListPtrOutput) ToPemCertificateListPtrOutputWithContext(ctx context.Context) PemCertificateListPtrOutput {
	return o
}

func (o PemCertificateListPtrOutput) Elem() PemCertificateListOutput {
	return o.ApplyT(func(v *PemCertificateList) PemCertificateList {
		if v != nil {
			return *v
		}
		var ret PemCertificateList
		return ret
	}).(PemCertificateListOutput)
}

func (o PemCertificateListPtrOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PemCertificateList) []string {
		if v == nil {
			return nil
		}
		return v.Certificates
	}).(pulumi.StringArrayOutput)
}

func (o PemCertificateListPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PemCertificateList) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type PemCertificateListResponse struct {
	Certificates []string `pulumi:"certificates"`
	Type         string   `pulumi:"type"`
}





type PemCertificateListResponseInput interface {
	pulumi.Input

	ToPemCertificateListResponseOutput() PemCertificateListResponseOutput
	ToPemCertificateListResponseOutputWithContext(context.Context) PemCertificateListResponseOutput
}

type PemCertificateListResponseArgs struct {
	Certificates pulumi.StringArrayInput `pulumi:"certificates"`
	Type         pulumi.StringInput      `pulumi:"type"`
}

func (PemCertificateListResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PemCertificateListResponse)(nil)).Elem()
}

func (i PemCertificateListResponseArgs) ToPemCertificateListResponseOutput() PemCertificateListResponseOutput {
	return i.ToPemCertificateListResponseOutputWithContext(context.Background())
}

func (i PemCertificateListResponseArgs) ToPemCertificateListResponseOutputWithContext(ctx context.Context) PemCertificateListResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PemCertificateListResponseOutput)
}

func (i PemCertificateListResponseArgs) ToPemCertificateListResponsePtrOutput() PemCertificateListResponsePtrOutput {
	return i.ToPemCertificateListResponsePtrOutputWithContext(context.Background())
}

func (i PemCertificateListResponseArgs) ToPemCertificateListResponsePtrOutputWithContext(ctx context.Context) PemCertificateListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PemCertificateListResponseOutput).ToPemCertificateListResponsePtrOutputWithContext(ctx)
}









type PemCertificateListResponsePtrInput interface {
	pulumi.Input

	ToPemCertificateListResponsePtrOutput() PemCertificateListResponsePtrOutput
	ToPemCertificateListResponsePtrOutputWithContext(context.Context) PemCertificateListResponsePtrOutput
}

type pemCertificateListResponsePtrType PemCertificateListResponseArgs

func PemCertificateListResponsePtr(v *PemCertificateListResponseArgs) PemCertificateListResponsePtrInput {
	return (*pemCertificateListResponsePtrType)(v)
}

func (*pemCertificateListResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PemCertificateListResponse)(nil)).Elem()
}

func (i *pemCertificateListResponsePtrType) ToPemCertificateListResponsePtrOutput() PemCertificateListResponsePtrOutput {
	return i.ToPemCertificateListResponsePtrOutputWithContext(context.Background())
}

func (i *pemCertificateListResponsePtrType) ToPemCertificateListResponsePtrOutputWithContext(ctx context.Context) PemCertificateListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PemCertificateListResponsePtrOutput)
}

type PemCertificateListResponseOutput struct{ *pulumi.OutputState }

func (PemCertificateListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PemCertificateListResponse)(nil)).Elem()
}

func (o PemCertificateListResponseOutput) ToPemCertificateListResponseOutput() PemCertificateListResponseOutput {
	return o
}

func (o PemCertificateListResponseOutput) ToPemCertificateListResponseOutputWithContext(ctx context.Context) PemCertificateListResponseOutput {
	return o
}

func (o PemCertificateListResponseOutput) ToPemCertificateListResponsePtrOutput() PemCertificateListResponsePtrOutput {
	return o.ToPemCertificateListResponsePtrOutputWithContext(context.Background())
}

func (o PemCertificateListResponseOutput) ToPemCertificateListResponsePtrOutputWithContext(ctx context.Context) PemCertificateListResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PemCertificateListResponse) *PemCertificateListResponse {
		return &v
	}).(PemCertificateListResponsePtrOutput)
}

func (o PemCertificateListResponseOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PemCertificateListResponse) []string { return v.Certificates }).(pulumi.StringArrayOutput)
}

func (o PemCertificateListResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PemCertificateListResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PemCertificateListResponsePtrOutput struct{ *pulumi.OutputState }

func (PemCertificateListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PemCertificateListResponse)(nil)).Elem()
}

func (o PemCertificateListResponsePtrOutput) ToPemCertificateListResponsePtrOutput() PemCertificateListResponsePtrOutput {
	return o
}

func (o PemCertificateListResponsePtrOutput) ToPemCertificateListResponsePtrOutputWithContext(ctx context.Context) PemCertificateListResponsePtrOutput {
	return o
}

func (o PemCertificateListResponsePtrOutput) Elem() PemCertificateListResponseOutput {
	return o.ApplyT(func(v *PemCertificateListResponse) PemCertificateListResponse {
		if v != nil {
			return *v
		}
		var ret PemCertificateListResponse
		return ret
	}).(PemCertificateListResponseOutput)
}

func (o PemCertificateListResponsePtrOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PemCertificateListResponse) []string {
		if v == nil {
			return nil
		}
		return v.Certificates
	}).(pulumi.StringArrayOutput)
}

func (o PemCertificateListResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PemCertificateListResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type PipelineJobErrorResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}





type PipelineJobErrorResponseInput interface {
	pulumi.Input

	ToPipelineJobErrorResponseOutput() PipelineJobErrorResponseOutput
	ToPipelineJobErrorResponseOutputWithContext(context.Context) PipelineJobErrorResponseOutput
}

type PipelineJobErrorResponseArgs struct {
	Code    pulumi.StringPtrInput `pulumi:"code"`
	Message pulumi.StringPtrInput `pulumi:"message"`
}

func (PipelineJobErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineJobErrorResponse)(nil)).Elem()
}

func (i PipelineJobErrorResponseArgs) ToPipelineJobErrorResponseOutput() PipelineJobErrorResponseOutput {
	return i.ToPipelineJobErrorResponseOutputWithContext(context.Background())
}

func (i PipelineJobErrorResponseArgs) ToPipelineJobErrorResponseOutputWithContext(ctx context.Context) PipelineJobErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineJobErrorResponseOutput)
}

func (i PipelineJobErrorResponseArgs) ToPipelineJobErrorResponsePtrOutput() PipelineJobErrorResponsePtrOutput {
	return i.ToPipelineJobErrorResponsePtrOutputWithContext(context.Background())
}

func (i PipelineJobErrorResponseArgs) ToPipelineJobErrorResponsePtrOutputWithContext(ctx context.Context) PipelineJobErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineJobErrorResponseOutput).ToPipelineJobErrorResponsePtrOutputWithContext(ctx)
}









type PipelineJobErrorResponsePtrInput interface {
	pulumi.Input

	ToPipelineJobErrorResponsePtrOutput() PipelineJobErrorResponsePtrOutput
	ToPipelineJobErrorResponsePtrOutputWithContext(context.Context) PipelineJobErrorResponsePtrOutput
}

type pipelineJobErrorResponsePtrType PipelineJobErrorResponseArgs

func PipelineJobErrorResponsePtr(v *PipelineJobErrorResponseArgs) PipelineJobErrorResponsePtrInput {
	return (*pipelineJobErrorResponsePtrType)(v)
}

func (*pipelineJobErrorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineJobErrorResponse)(nil)).Elem()
}

func (i *pipelineJobErrorResponsePtrType) ToPipelineJobErrorResponsePtrOutput() PipelineJobErrorResponsePtrOutput {
	return i.ToPipelineJobErrorResponsePtrOutputWithContext(context.Background())
}

func (i *pipelineJobErrorResponsePtrType) ToPipelineJobErrorResponsePtrOutputWithContext(ctx context.Context) PipelineJobErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineJobErrorResponsePtrOutput)
}

type PipelineJobErrorResponseOutput struct{ *pulumi.OutputState }

func (PipelineJobErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineJobErrorResponse)(nil)).Elem()
}

func (o PipelineJobErrorResponseOutput) ToPipelineJobErrorResponseOutput() PipelineJobErrorResponseOutput {
	return o
}

func (o PipelineJobErrorResponseOutput) ToPipelineJobErrorResponseOutputWithContext(ctx context.Context) PipelineJobErrorResponseOutput {
	return o
}

func (o PipelineJobErrorResponseOutput) ToPipelineJobErrorResponsePtrOutput() PipelineJobErrorResponsePtrOutput {
	return o.ToPipelineJobErrorResponsePtrOutputWithContext(context.Background())
}

func (o PipelineJobErrorResponseOutput) ToPipelineJobErrorResponsePtrOutputWithContext(ctx context.Context) PipelineJobErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineJobErrorResponse) *PipelineJobErrorResponse {
		return &v
	}).(PipelineJobErrorResponsePtrOutput)
}

func (o PipelineJobErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineJobErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o PipelineJobErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineJobErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type PipelineJobErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineJobErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineJobErrorResponse)(nil)).Elem()
}

func (o PipelineJobErrorResponsePtrOutput) ToPipelineJobErrorResponsePtrOutput() PipelineJobErrorResponsePtrOutput {
	return o
}

func (o PipelineJobErrorResponsePtrOutput) ToPipelineJobErrorResponsePtrOutputWithContext(ctx context.Context) PipelineJobErrorResponsePtrOutput {
	return o
}

func (o PipelineJobErrorResponsePtrOutput) Elem() PipelineJobErrorResponseOutput {
	return o.ApplyT(func(v *PipelineJobErrorResponse) PipelineJobErrorResponse {
		if v != nil {
			return *v
		}
		var ret PipelineJobErrorResponse
		return ret
	}).(PipelineJobErrorResponseOutput)
}

func (o PipelineJobErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineJobErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o PipelineJobErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineJobErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentity struct {
	UserAssignedIdentity string `pulumi:"userAssignedIdentity"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	UserAssignedIdentity pulumi.StringInput `pulumi:"userAssignedIdentity"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}









type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

func (o ResourceIdentityOutput) UserAssignedIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentity) string { return v.UserAssignedIdentity }).(pulumi.StringOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

func (o ResourceIdentityPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponse struct {
	UserAssignedIdentity string `pulumi:"userAssignedIdentity"`
}





type ResourceIdentityResponseInput interface {
	pulumi.Input

	ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput
	ToResourceIdentityResponseOutputWithContext(context.Context) ResourceIdentityResponseOutput
}

type ResourceIdentityResponseArgs struct {
	UserAssignedIdentity pulumi.StringInput `pulumi:"userAssignedIdentity"`
}

func (ResourceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return i.ToResourceIdentityResponseOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput)
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput).ToResourceIdentityResponsePtrOutputWithContext(ctx)
}









type ResourceIdentityResponsePtrInput interface {
	pulumi.Input

	ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput
	ToResourceIdentityResponsePtrOutputWithContext(context.Context) ResourceIdentityResponsePtrOutput
}

type resourceIdentityResponsePtrType ResourceIdentityResponseArgs

func ResourceIdentityResponsePtr(v *ResourceIdentityResponseArgs) ResourceIdentityResponsePtrInput {
	return (*resourceIdentityResponsePtrType)(v)
}

func (*resourceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponsePtrOutput)
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityResponse) *ResourceIdentityResponse {
		return &v
	}).(ResourceIdentityResponsePtrOutput)
}

func (o ResourceIdentityResponseOutput) UserAssignedIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.UserAssignedIdentity }).(pulumi.StringOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

func (o ResourceIdentityResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type RsaTokenKey struct {
	Alg  string `pulumi:"alg"`
	E    string `pulumi:"e"`
	Kid  string `pulumi:"kid"`
	N    string `pulumi:"n"`
	Type string `pulumi:"type"`
}





type RsaTokenKeyInput interface {
	pulumi.Input

	ToRsaTokenKeyOutput() RsaTokenKeyOutput
	ToRsaTokenKeyOutputWithContext(context.Context) RsaTokenKeyOutput
}

type RsaTokenKeyArgs struct {
	Alg  pulumi.StringInput `pulumi:"alg"`
	E    pulumi.StringInput `pulumi:"e"`
	Kid  pulumi.StringInput `pulumi:"kid"`
	N    pulumi.StringInput `pulumi:"n"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (RsaTokenKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RsaTokenKey)(nil)).Elem()
}

func (i RsaTokenKeyArgs) ToRsaTokenKeyOutput() RsaTokenKeyOutput {
	return i.ToRsaTokenKeyOutputWithContext(context.Background())
}

func (i RsaTokenKeyArgs) ToRsaTokenKeyOutputWithContext(ctx context.Context) RsaTokenKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RsaTokenKeyOutput)
}

type RsaTokenKeyOutput struct{ *pulumi.OutputState }

func (RsaTokenKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RsaTokenKey)(nil)).Elem()
}

func (o RsaTokenKeyOutput) ToRsaTokenKeyOutput() RsaTokenKeyOutput {
	return o
}

func (o RsaTokenKeyOutput) ToRsaTokenKeyOutputWithContext(ctx context.Context) RsaTokenKeyOutput {
	return o
}

func (o RsaTokenKeyOutput) Alg() pulumi.StringOutput {
	return o.ApplyT(func(v RsaTokenKey) string { return v.Alg }).(pulumi.StringOutput)
}

func (o RsaTokenKeyOutput) E() pulumi.StringOutput {
	return o.ApplyT(func(v RsaTokenKey) string { return v.E }).(pulumi.StringOutput)
}

func (o RsaTokenKeyOutput) Kid() pulumi.StringOutput {
	return o.ApplyT(func(v RsaTokenKey) string { return v.Kid }).(pulumi.StringOutput)
}

func (o RsaTokenKeyOutput) N() pulumi.StringOutput {
	return o.ApplyT(func(v RsaTokenKey) string { return v.N }).(pulumi.StringOutput)
}

func (o RsaTokenKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RsaTokenKey) string { return v.Type }).(pulumi.StringOutput)
}

type RsaTokenKeyResponse struct {
	Alg  string `pulumi:"alg"`
	E    string `pulumi:"e"`
	Kid  string `pulumi:"kid"`
	N    string `pulumi:"n"`
	Type string `pulumi:"type"`
}





type RsaTokenKeyResponseInput interface {
	pulumi.Input

	ToRsaTokenKeyResponseOutput() RsaTokenKeyResponseOutput
	ToRsaTokenKeyResponseOutputWithContext(context.Context) RsaTokenKeyResponseOutput
}

type RsaTokenKeyResponseArgs struct {
	Alg  pulumi.StringInput `pulumi:"alg"`
	E    pulumi.StringInput `pulumi:"e"`
	Kid  pulumi.StringInput `pulumi:"kid"`
	N    pulumi.StringInput `pulumi:"n"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (RsaTokenKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RsaTokenKeyResponse)(nil)).Elem()
}

func (i RsaTokenKeyResponseArgs) ToRsaTokenKeyResponseOutput() RsaTokenKeyResponseOutput {
	return i.ToRsaTokenKeyResponseOutputWithContext(context.Background())
}

func (i RsaTokenKeyResponseArgs) ToRsaTokenKeyResponseOutputWithContext(ctx context.Context) RsaTokenKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RsaTokenKeyResponseOutput)
}

type RsaTokenKeyResponseOutput struct{ *pulumi.OutputState }

func (RsaTokenKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RsaTokenKeyResponse)(nil)).Elem()
}

func (o RsaTokenKeyResponseOutput) ToRsaTokenKeyResponseOutput() RsaTokenKeyResponseOutput {
	return o
}

func (o RsaTokenKeyResponseOutput) ToRsaTokenKeyResponseOutputWithContext(ctx context.Context) RsaTokenKeyResponseOutput {
	return o
}

func (o RsaTokenKeyResponseOutput) Alg() pulumi.StringOutput {
	return o.ApplyT(func(v RsaTokenKeyResponse) string { return v.Alg }).(pulumi.StringOutput)
}

func (o RsaTokenKeyResponseOutput) E() pulumi.StringOutput {
	return o.ApplyT(func(v RsaTokenKeyResponse) string { return v.E }).(pulumi.StringOutput)
}

func (o RsaTokenKeyResponseOutput) Kid() pulumi.StringOutput {
	return o.ApplyT(func(v RsaTokenKeyResponse) string { return v.Kid }).(pulumi.StringOutput)
}

func (o RsaTokenKeyResponseOutput) N() pulumi.StringOutput {
	return o.ApplyT(func(v RsaTokenKeyResponse) string { return v.N }).(pulumi.StringOutput)
}

func (o RsaTokenKeyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RsaTokenKeyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RtspSource struct {
	Endpoint  interface{} `pulumi:"endpoint"`
	Name      string      `pulumi:"name"`
	Transport *string     `pulumi:"transport"`
	Type      string      `pulumi:"type"`
}





type RtspSourceInput interface {
	pulumi.Input

	ToRtspSourceOutput() RtspSourceOutput
	ToRtspSourceOutputWithContext(context.Context) RtspSourceOutput
}

type RtspSourceArgs struct {
	Endpoint  pulumi.Input          `pulumi:"endpoint"`
	Name      pulumi.StringInput    `pulumi:"name"`
	Transport pulumi.StringPtrInput `pulumi:"transport"`
	Type      pulumi.StringInput    `pulumi:"type"`
}

func (RtspSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RtspSource)(nil)).Elem()
}

func (i RtspSourceArgs) ToRtspSourceOutput() RtspSourceOutput {
	return i.ToRtspSourceOutputWithContext(context.Background())
}

func (i RtspSourceArgs) ToRtspSourceOutputWithContext(ctx context.Context) RtspSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RtspSourceOutput)
}

type RtspSourceOutput struct{ *pulumi.OutputState }

func (RtspSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RtspSource)(nil)).Elem()
}

func (o RtspSourceOutput) ToRtspSourceOutput() RtspSourceOutput {
	return o
}

func (o RtspSourceOutput) ToRtspSourceOutputWithContext(ctx context.Context) RtspSourceOutput {
	return o
}

func (o RtspSourceOutput) Endpoint() pulumi.AnyOutput {
	return o.ApplyT(func(v RtspSource) interface{} { return v.Endpoint }).(pulumi.AnyOutput)
}

func (o RtspSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RtspSource) string { return v.Name }).(pulumi.StringOutput)
}

func (o RtspSourceOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RtspSource) *string { return v.Transport }).(pulumi.StringPtrOutput)
}

func (o RtspSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RtspSource) string { return v.Type }).(pulumi.StringOutput)
}

type RtspSourceResponse struct {
	Endpoint  interface{} `pulumi:"endpoint"`
	Name      string      `pulumi:"name"`
	Transport *string     `pulumi:"transport"`
	Type      string      `pulumi:"type"`
}





type RtspSourceResponseInput interface {
	pulumi.Input

	ToRtspSourceResponseOutput() RtspSourceResponseOutput
	ToRtspSourceResponseOutputWithContext(context.Context) RtspSourceResponseOutput
}

type RtspSourceResponseArgs struct {
	Endpoint  pulumi.Input          `pulumi:"endpoint"`
	Name      pulumi.StringInput    `pulumi:"name"`
	Transport pulumi.StringPtrInput `pulumi:"transport"`
	Type      pulumi.StringInput    `pulumi:"type"`
}

func (RtspSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RtspSourceResponse)(nil)).Elem()
}

func (i RtspSourceResponseArgs) ToRtspSourceResponseOutput() RtspSourceResponseOutput {
	return i.ToRtspSourceResponseOutputWithContext(context.Background())
}

func (i RtspSourceResponseArgs) ToRtspSourceResponseOutputWithContext(ctx context.Context) RtspSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RtspSourceResponseOutput)
}

type RtspSourceResponseOutput struct{ *pulumi.OutputState }

func (RtspSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RtspSourceResponse)(nil)).Elem()
}

func (o RtspSourceResponseOutput) ToRtspSourceResponseOutput() RtspSourceResponseOutput {
	return o
}

func (o RtspSourceResponseOutput) ToRtspSourceResponseOutputWithContext(ctx context.Context) RtspSourceResponseOutput {
	return o
}

func (o RtspSourceResponseOutput) Endpoint() pulumi.AnyOutput {
	return o.ApplyT(func(v RtspSourceResponse) interface{} { return v.Endpoint }).(pulumi.AnyOutput)
}

func (o RtspSourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RtspSourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RtspSourceResponseOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RtspSourceResponse) *string { return v.Transport }).(pulumi.StringPtrOutput)
}

func (o RtspSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RtspSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SecureIotDeviceRemoteTunnel struct {
	DeviceId   string `pulumi:"deviceId"`
	IotHubName string `pulumi:"iotHubName"`
	Type       string `pulumi:"type"`
}





type SecureIotDeviceRemoteTunnelInput interface {
	pulumi.Input

	ToSecureIotDeviceRemoteTunnelOutput() SecureIotDeviceRemoteTunnelOutput
	ToSecureIotDeviceRemoteTunnelOutputWithContext(context.Context) SecureIotDeviceRemoteTunnelOutput
}

type SecureIotDeviceRemoteTunnelArgs struct {
	DeviceId   pulumi.StringInput `pulumi:"deviceId"`
	IotHubName pulumi.StringInput `pulumi:"iotHubName"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (SecureIotDeviceRemoteTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecureIotDeviceRemoteTunnel)(nil)).Elem()
}

func (i SecureIotDeviceRemoteTunnelArgs) ToSecureIotDeviceRemoteTunnelOutput() SecureIotDeviceRemoteTunnelOutput {
	return i.ToSecureIotDeviceRemoteTunnelOutputWithContext(context.Background())
}

func (i SecureIotDeviceRemoteTunnelArgs) ToSecureIotDeviceRemoteTunnelOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecureIotDeviceRemoteTunnelOutput)
}

func (i SecureIotDeviceRemoteTunnelArgs) ToSecureIotDeviceRemoteTunnelPtrOutput() SecureIotDeviceRemoteTunnelPtrOutput {
	return i.ToSecureIotDeviceRemoteTunnelPtrOutputWithContext(context.Background())
}

func (i SecureIotDeviceRemoteTunnelArgs) ToSecureIotDeviceRemoteTunnelPtrOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecureIotDeviceRemoteTunnelOutput).ToSecureIotDeviceRemoteTunnelPtrOutputWithContext(ctx)
}









type SecureIotDeviceRemoteTunnelPtrInput interface {
	pulumi.Input

	ToSecureIotDeviceRemoteTunnelPtrOutput() SecureIotDeviceRemoteTunnelPtrOutput
	ToSecureIotDeviceRemoteTunnelPtrOutputWithContext(context.Context) SecureIotDeviceRemoteTunnelPtrOutput
}

type secureIotDeviceRemoteTunnelPtrType SecureIotDeviceRemoteTunnelArgs

func SecureIotDeviceRemoteTunnelPtr(v *SecureIotDeviceRemoteTunnelArgs) SecureIotDeviceRemoteTunnelPtrInput {
	return (*secureIotDeviceRemoteTunnelPtrType)(v)
}

func (*secureIotDeviceRemoteTunnelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureIotDeviceRemoteTunnel)(nil)).Elem()
}

func (i *secureIotDeviceRemoteTunnelPtrType) ToSecureIotDeviceRemoteTunnelPtrOutput() SecureIotDeviceRemoteTunnelPtrOutput {
	return i.ToSecureIotDeviceRemoteTunnelPtrOutputWithContext(context.Background())
}

func (i *secureIotDeviceRemoteTunnelPtrType) ToSecureIotDeviceRemoteTunnelPtrOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecureIotDeviceRemoteTunnelPtrOutput)
}

type SecureIotDeviceRemoteTunnelOutput struct{ *pulumi.OutputState }

func (SecureIotDeviceRemoteTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecureIotDeviceRemoteTunnel)(nil)).Elem()
}

func (o SecureIotDeviceRemoteTunnelOutput) ToSecureIotDeviceRemoteTunnelOutput() SecureIotDeviceRemoteTunnelOutput {
	return o
}

func (o SecureIotDeviceRemoteTunnelOutput) ToSecureIotDeviceRemoteTunnelOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelOutput {
	return o
}

func (o SecureIotDeviceRemoteTunnelOutput) ToSecureIotDeviceRemoteTunnelPtrOutput() SecureIotDeviceRemoteTunnelPtrOutput {
	return o.ToSecureIotDeviceRemoteTunnelPtrOutputWithContext(context.Background())
}

func (o SecureIotDeviceRemoteTunnelOutput) ToSecureIotDeviceRemoteTunnelPtrOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecureIotDeviceRemoteTunnel) *SecureIotDeviceRemoteTunnel {
		return &v
	}).(SecureIotDeviceRemoteTunnelPtrOutput)
}

func (o SecureIotDeviceRemoteTunnelOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v SecureIotDeviceRemoteTunnel) string { return v.DeviceId }).(pulumi.StringOutput)
}

func (o SecureIotDeviceRemoteTunnelOutput) IotHubName() pulumi.StringOutput {
	return o.ApplyT(func(v SecureIotDeviceRemoteTunnel) string { return v.IotHubName }).(pulumi.StringOutput)
}

func (o SecureIotDeviceRemoteTunnelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SecureIotDeviceRemoteTunnel) string { return v.Type }).(pulumi.StringOutput)
}

type SecureIotDeviceRemoteTunnelPtrOutput struct{ *pulumi.OutputState }

func (SecureIotDeviceRemoteTunnelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureIotDeviceRemoteTunnel)(nil)).Elem()
}

func (o SecureIotDeviceRemoteTunnelPtrOutput) ToSecureIotDeviceRemoteTunnelPtrOutput() SecureIotDeviceRemoteTunnelPtrOutput {
	return o
}

func (o SecureIotDeviceRemoteTunnelPtrOutput) ToSecureIotDeviceRemoteTunnelPtrOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelPtrOutput {
	return o
}

func (o SecureIotDeviceRemoteTunnelPtrOutput) Elem() SecureIotDeviceRemoteTunnelOutput {
	return o.ApplyT(func(v *SecureIotDeviceRemoteTunnel) SecureIotDeviceRemoteTunnel {
		if v != nil {
			return *v
		}
		var ret SecureIotDeviceRemoteTunnel
		return ret
	}).(SecureIotDeviceRemoteTunnelOutput)
}

func (o SecureIotDeviceRemoteTunnelPtrOutput) DeviceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecureIotDeviceRemoteTunnel) *string {
		if v == nil {
			return nil
		}
		return &v.DeviceId
	}).(pulumi.StringPtrOutput)
}

func (o SecureIotDeviceRemoteTunnelPtrOutput) IotHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecureIotDeviceRemoteTunnel) *string {
		if v == nil {
			return nil
		}
		return &v.IotHubName
	}).(pulumi.StringPtrOutput)
}

func (o SecureIotDeviceRemoteTunnelPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecureIotDeviceRemoteTunnel) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type SecureIotDeviceRemoteTunnelResponse struct {
	DeviceId   string `pulumi:"deviceId"`
	IotHubName string `pulumi:"iotHubName"`
	Type       string `pulumi:"type"`
}





type SecureIotDeviceRemoteTunnelResponseInput interface {
	pulumi.Input

	ToSecureIotDeviceRemoteTunnelResponseOutput() SecureIotDeviceRemoteTunnelResponseOutput
	ToSecureIotDeviceRemoteTunnelResponseOutputWithContext(context.Context) SecureIotDeviceRemoteTunnelResponseOutput
}

type SecureIotDeviceRemoteTunnelResponseArgs struct {
	DeviceId   pulumi.StringInput `pulumi:"deviceId"`
	IotHubName pulumi.StringInput `pulumi:"iotHubName"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (SecureIotDeviceRemoteTunnelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecureIotDeviceRemoteTunnelResponse)(nil)).Elem()
}

func (i SecureIotDeviceRemoteTunnelResponseArgs) ToSecureIotDeviceRemoteTunnelResponseOutput() SecureIotDeviceRemoteTunnelResponseOutput {
	return i.ToSecureIotDeviceRemoteTunnelResponseOutputWithContext(context.Background())
}

func (i SecureIotDeviceRemoteTunnelResponseArgs) ToSecureIotDeviceRemoteTunnelResponseOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecureIotDeviceRemoteTunnelResponseOutput)
}

func (i SecureIotDeviceRemoteTunnelResponseArgs) ToSecureIotDeviceRemoteTunnelResponsePtrOutput() SecureIotDeviceRemoteTunnelResponsePtrOutput {
	return i.ToSecureIotDeviceRemoteTunnelResponsePtrOutputWithContext(context.Background())
}

func (i SecureIotDeviceRemoteTunnelResponseArgs) ToSecureIotDeviceRemoteTunnelResponsePtrOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecureIotDeviceRemoteTunnelResponseOutput).ToSecureIotDeviceRemoteTunnelResponsePtrOutputWithContext(ctx)
}









type SecureIotDeviceRemoteTunnelResponsePtrInput interface {
	pulumi.Input

	ToSecureIotDeviceRemoteTunnelResponsePtrOutput() SecureIotDeviceRemoteTunnelResponsePtrOutput
	ToSecureIotDeviceRemoteTunnelResponsePtrOutputWithContext(context.Context) SecureIotDeviceRemoteTunnelResponsePtrOutput
}

type secureIotDeviceRemoteTunnelResponsePtrType SecureIotDeviceRemoteTunnelResponseArgs

func SecureIotDeviceRemoteTunnelResponsePtr(v *SecureIotDeviceRemoteTunnelResponseArgs) SecureIotDeviceRemoteTunnelResponsePtrInput {
	return (*secureIotDeviceRemoteTunnelResponsePtrType)(v)
}

func (*secureIotDeviceRemoteTunnelResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureIotDeviceRemoteTunnelResponse)(nil)).Elem()
}

func (i *secureIotDeviceRemoteTunnelResponsePtrType) ToSecureIotDeviceRemoteTunnelResponsePtrOutput() SecureIotDeviceRemoteTunnelResponsePtrOutput {
	return i.ToSecureIotDeviceRemoteTunnelResponsePtrOutputWithContext(context.Background())
}

func (i *secureIotDeviceRemoteTunnelResponsePtrType) ToSecureIotDeviceRemoteTunnelResponsePtrOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecureIotDeviceRemoteTunnelResponsePtrOutput)
}

type SecureIotDeviceRemoteTunnelResponseOutput struct{ *pulumi.OutputState }

func (SecureIotDeviceRemoteTunnelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecureIotDeviceRemoteTunnelResponse)(nil)).Elem()
}

func (o SecureIotDeviceRemoteTunnelResponseOutput) ToSecureIotDeviceRemoteTunnelResponseOutput() SecureIotDeviceRemoteTunnelResponseOutput {
	return o
}

func (o SecureIotDeviceRemoteTunnelResponseOutput) ToSecureIotDeviceRemoteTunnelResponseOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelResponseOutput {
	return o
}

func (o SecureIotDeviceRemoteTunnelResponseOutput) ToSecureIotDeviceRemoteTunnelResponsePtrOutput() SecureIotDeviceRemoteTunnelResponsePtrOutput {
	return o.ToSecureIotDeviceRemoteTunnelResponsePtrOutputWithContext(context.Background())
}

func (o SecureIotDeviceRemoteTunnelResponseOutput) ToSecureIotDeviceRemoteTunnelResponsePtrOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecureIotDeviceRemoteTunnelResponse) *SecureIotDeviceRemoteTunnelResponse {
		return &v
	}).(SecureIotDeviceRemoteTunnelResponsePtrOutput)
}

func (o SecureIotDeviceRemoteTunnelResponseOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v SecureIotDeviceRemoteTunnelResponse) string { return v.DeviceId }).(pulumi.StringOutput)
}

func (o SecureIotDeviceRemoteTunnelResponseOutput) IotHubName() pulumi.StringOutput {
	return o.ApplyT(func(v SecureIotDeviceRemoteTunnelResponse) string { return v.IotHubName }).(pulumi.StringOutput)
}

func (o SecureIotDeviceRemoteTunnelResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SecureIotDeviceRemoteTunnelResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SecureIotDeviceRemoteTunnelResponsePtrOutput struct{ *pulumi.OutputState }

func (SecureIotDeviceRemoteTunnelResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecureIotDeviceRemoteTunnelResponse)(nil)).Elem()
}

func (o SecureIotDeviceRemoteTunnelResponsePtrOutput) ToSecureIotDeviceRemoteTunnelResponsePtrOutput() SecureIotDeviceRemoteTunnelResponsePtrOutput {
	return o
}

func (o SecureIotDeviceRemoteTunnelResponsePtrOutput) ToSecureIotDeviceRemoteTunnelResponsePtrOutputWithContext(ctx context.Context) SecureIotDeviceRemoteTunnelResponsePtrOutput {
	return o
}

func (o SecureIotDeviceRemoteTunnelResponsePtrOutput) Elem() SecureIotDeviceRemoteTunnelResponseOutput {
	return o.ApplyT(func(v *SecureIotDeviceRemoteTunnelResponse) SecureIotDeviceRemoteTunnelResponse {
		if v != nil {
			return *v
		}
		var ret SecureIotDeviceRemoteTunnelResponse
		return ret
	}).(SecureIotDeviceRemoteTunnelResponseOutput)
}

func (o SecureIotDeviceRemoteTunnelResponsePtrOutput) DeviceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecureIotDeviceRemoteTunnelResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DeviceId
	}).(pulumi.StringPtrOutput)
}

func (o SecureIotDeviceRemoteTunnelResponsePtrOutput) IotHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecureIotDeviceRemoteTunnelResponse) *string {
		if v == nil {
			return nil
		}
		return &v.IotHubName
	}).(pulumi.StringPtrOutput)
}

func (o SecureIotDeviceRemoteTunnelResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecureIotDeviceRemoteTunnelResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
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

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
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

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
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

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
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

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type StorageAccount struct {
	Id       *string           `pulumi:"id"`
	Identity *ResourceIdentity `pulumi:"identity"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	Id       pulumi.StringPtrInput    `pulumi:"id"`
	Identity ResourceIdentityPtrInput `pulumi:"identity"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}





type StorageAccountArrayInput interface {
	pulumi.Input

	ToStorageAccountArrayOutput() StorageAccountArrayOutput
	ToStorageAccountArrayOutputWithContext(context.Context) StorageAccountArrayOutput
}

type StorageAccountArray []StorageAccountInput

func (StorageAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (i StorageAccountArray) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return i.ToStorageAccountArrayOutputWithContext(context.Background())
}

func (i StorageAccountArray) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountArrayOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Identity() ResourceIdentityPtrOutput {
	return o.ApplyT(func(v StorageAccount) *ResourceIdentity { return v.Identity }).(ResourceIdentityPtrOutput)
}

type StorageAccountArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) Index(i pulumi.IntInput) StorageAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccount {
		return vs[0].([]StorageAccount)[vs[1].(int)]
	}).(StorageAccountOutput)
}

type StorageAccountResponse struct {
	Id       *string                   `pulumi:"id"`
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	Status   string                    `pulumi:"status"`
}





type StorageAccountResponseInput interface {
	pulumi.Input

	ToStorageAccountResponseOutput() StorageAccountResponseOutput
	ToStorageAccountResponseOutputWithContext(context.Context) StorageAccountResponseOutput
}

type StorageAccountResponseArgs struct {
	Id       pulumi.StringPtrInput            `pulumi:"id"`
	Identity ResourceIdentityResponsePtrInput `pulumi:"identity"`
	Status   pulumi.StringInput               `pulumi:"status"`
}

func (StorageAccountResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (i StorageAccountResponseArgs) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return i.ToStorageAccountResponseOutputWithContext(context.Background())
}

func (i StorageAccountResponseArgs) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountResponseOutput)
}





type StorageAccountResponseArrayInput interface {
	pulumi.Input

	ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput
	ToStorageAccountResponseArrayOutputWithContext(context.Context) StorageAccountResponseArrayOutput
}

type StorageAccountResponseArray []StorageAccountResponseInput

func (StorageAccountResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountResponse)(nil)).Elem()
}

func (i StorageAccountResponseArray) ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput {
	return i.ToStorageAccountResponseArrayOutputWithContext(context.Background())
}

func (i StorageAccountResponseArray) ToStorageAccountResponseArrayOutputWithContext(ctx context.Context) StorageAccountResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountResponseArrayOutput)
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o StorageAccountResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Status }).(pulumi.StringOutput)
}

type StorageAccountResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutputWithContext(ctx context.Context) StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountResponse {
		return vs[0].([]StorageAccountResponse)[vs[1].(int)]
	}).(StorageAccountResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type TlsEndpoint struct {
	Credentials         *UsernamePasswordCredentials `pulumi:"credentials"`
	TrustedCertificates *PemCertificateList          `pulumi:"trustedCertificates"`
	Tunnel              *SecureIotDeviceRemoteTunnel `pulumi:"tunnel"`
	Type                string                       `pulumi:"type"`
	Url                 string                       `pulumi:"url"`
	ValidationOptions   *TlsValidationOptions        `pulumi:"validationOptions"`
}





type TlsEndpointInput interface {
	pulumi.Input

	ToTlsEndpointOutput() TlsEndpointOutput
	ToTlsEndpointOutputWithContext(context.Context) TlsEndpointOutput
}

type TlsEndpointArgs struct {
	Credentials         UsernamePasswordCredentialsPtrInput `pulumi:"credentials"`
	TrustedCertificates PemCertificateListPtrInput          `pulumi:"trustedCertificates"`
	Tunnel              SecureIotDeviceRemoteTunnelPtrInput `pulumi:"tunnel"`
	Type                pulumi.StringInput                  `pulumi:"type"`
	Url                 pulumi.StringInput                  `pulumi:"url"`
	ValidationOptions   TlsValidationOptionsPtrInput        `pulumi:"validationOptions"`
}

func (TlsEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsEndpoint)(nil)).Elem()
}

func (i TlsEndpointArgs) ToTlsEndpointOutput() TlsEndpointOutput {
	return i.ToTlsEndpointOutputWithContext(context.Background())
}

func (i TlsEndpointArgs) ToTlsEndpointOutputWithContext(ctx context.Context) TlsEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsEndpointOutput)
}

type TlsEndpointOutput struct{ *pulumi.OutputState }

func (TlsEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsEndpoint)(nil)).Elem()
}

func (o TlsEndpointOutput) ToTlsEndpointOutput() TlsEndpointOutput {
	return o
}

func (o TlsEndpointOutput) ToTlsEndpointOutputWithContext(ctx context.Context) TlsEndpointOutput {
	return o
}

func (o TlsEndpointOutput) Credentials() UsernamePasswordCredentialsPtrOutput {
	return o.ApplyT(func(v TlsEndpoint) *UsernamePasswordCredentials { return v.Credentials }).(UsernamePasswordCredentialsPtrOutput)
}

func (o TlsEndpointOutput) TrustedCertificates() PemCertificateListPtrOutput {
	return o.ApplyT(func(v TlsEndpoint) *PemCertificateList { return v.TrustedCertificates }).(PemCertificateListPtrOutput)
}

func (o TlsEndpointOutput) Tunnel() SecureIotDeviceRemoteTunnelPtrOutput {
	return o.ApplyT(func(v TlsEndpoint) *SecureIotDeviceRemoteTunnel { return v.Tunnel }).(SecureIotDeviceRemoteTunnelPtrOutput)
}

func (o TlsEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TlsEndpoint) string { return v.Type }).(pulumi.StringOutput)
}

func (o TlsEndpointOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v TlsEndpoint) string { return v.Url }).(pulumi.StringOutput)
}

func (o TlsEndpointOutput) ValidationOptions() TlsValidationOptionsPtrOutput {
	return o.ApplyT(func(v TlsEndpoint) *TlsValidationOptions { return v.ValidationOptions }).(TlsValidationOptionsPtrOutput)
}

type TlsEndpointResponse struct {
	Credentials         *UsernamePasswordCredentialsResponse `pulumi:"credentials"`
	TrustedCertificates *PemCertificateListResponse          `pulumi:"trustedCertificates"`
	Tunnel              *SecureIotDeviceRemoteTunnelResponse `pulumi:"tunnel"`
	Type                string                               `pulumi:"type"`
	Url                 string                               `pulumi:"url"`
	ValidationOptions   *TlsValidationOptionsResponse        `pulumi:"validationOptions"`
}





type TlsEndpointResponseInput interface {
	pulumi.Input

	ToTlsEndpointResponseOutput() TlsEndpointResponseOutput
	ToTlsEndpointResponseOutputWithContext(context.Context) TlsEndpointResponseOutput
}

type TlsEndpointResponseArgs struct {
	Credentials         UsernamePasswordCredentialsResponsePtrInput `pulumi:"credentials"`
	TrustedCertificates PemCertificateListResponsePtrInput          `pulumi:"trustedCertificates"`
	Tunnel              SecureIotDeviceRemoteTunnelResponsePtrInput `pulumi:"tunnel"`
	Type                pulumi.StringInput                          `pulumi:"type"`
	Url                 pulumi.StringInput                          `pulumi:"url"`
	ValidationOptions   TlsValidationOptionsResponsePtrInput        `pulumi:"validationOptions"`
}

func (TlsEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsEndpointResponse)(nil)).Elem()
}

func (i TlsEndpointResponseArgs) ToTlsEndpointResponseOutput() TlsEndpointResponseOutput {
	return i.ToTlsEndpointResponseOutputWithContext(context.Background())
}

func (i TlsEndpointResponseArgs) ToTlsEndpointResponseOutputWithContext(ctx context.Context) TlsEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsEndpointResponseOutput)
}

type TlsEndpointResponseOutput struct{ *pulumi.OutputState }

func (TlsEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsEndpointResponse)(nil)).Elem()
}

func (o TlsEndpointResponseOutput) ToTlsEndpointResponseOutput() TlsEndpointResponseOutput {
	return o
}

func (o TlsEndpointResponseOutput) ToTlsEndpointResponseOutputWithContext(ctx context.Context) TlsEndpointResponseOutput {
	return o
}

func (o TlsEndpointResponseOutput) Credentials() UsernamePasswordCredentialsResponsePtrOutput {
	return o.ApplyT(func(v TlsEndpointResponse) *UsernamePasswordCredentialsResponse { return v.Credentials }).(UsernamePasswordCredentialsResponsePtrOutput)
}

func (o TlsEndpointResponseOutput) TrustedCertificates() PemCertificateListResponsePtrOutput {
	return o.ApplyT(func(v TlsEndpointResponse) *PemCertificateListResponse { return v.TrustedCertificates }).(PemCertificateListResponsePtrOutput)
}

func (o TlsEndpointResponseOutput) Tunnel() SecureIotDeviceRemoteTunnelResponsePtrOutput {
	return o.ApplyT(func(v TlsEndpointResponse) *SecureIotDeviceRemoteTunnelResponse { return v.Tunnel }).(SecureIotDeviceRemoteTunnelResponsePtrOutput)
}

func (o TlsEndpointResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TlsEndpointResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o TlsEndpointResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v TlsEndpointResponse) string { return v.Url }).(pulumi.StringOutput)
}

func (o TlsEndpointResponseOutput) ValidationOptions() TlsValidationOptionsResponsePtrOutput {
	return o.ApplyT(func(v TlsEndpointResponse) *TlsValidationOptionsResponse { return v.ValidationOptions }).(TlsValidationOptionsResponsePtrOutput)
}

type TlsValidationOptions struct {
	IgnoreHostname  *string `pulumi:"ignoreHostname"`
	IgnoreSignature *string `pulumi:"ignoreSignature"`
}





type TlsValidationOptionsInput interface {
	pulumi.Input

	ToTlsValidationOptionsOutput() TlsValidationOptionsOutput
	ToTlsValidationOptionsOutputWithContext(context.Context) TlsValidationOptionsOutput
}

type TlsValidationOptionsArgs struct {
	IgnoreHostname  pulumi.StringPtrInput `pulumi:"ignoreHostname"`
	IgnoreSignature pulumi.StringPtrInput `pulumi:"ignoreSignature"`
}

func (TlsValidationOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsValidationOptions)(nil)).Elem()
}

func (i TlsValidationOptionsArgs) ToTlsValidationOptionsOutput() TlsValidationOptionsOutput {
	return i.ToTlsValidationOptionsOutputWithContext(context.Background())
}

func (i TlsValidationOptionsArgs) ToTlsValidationOptionsOutputWithContext(ctx context.Context) TlsValidationOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsValidationOptionsOutput)
}

func (i TlsValidationOptionsArgs) ToTlsValidationOptionsPtrOutput() TlsValidationOptionsPtrOutput {
	return i.ToTlsValidationOptionsPtrOutputWithContext(context.Background())
}

func (i TlsValidationOptionsArgs) ToTlsValidationOptionsPtrOutputWithContext(ctx context.Context) TlsValidationOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsValidationOptionsOutput).ToTlsValidationOptionsPtrOutputWithContext(ctx)
}









type TlsValidationOptionsPtrInput interface {
	pulumi.Input

	ToTlsValidationOptionsPtrOutput() TlsValidationOptionsPtrOutput
	ToTlsValidationOptionsPtrOutputWithContext(context.Context) TlsValidationOptionsPtrOutput
}

type tlsValidationOptionsPtrType TlsValidationOptionsArgs

func TlsValidationOptionsPtr(v *TlsValidationOptionsArgs) TlsValidationOptionsPtrInput {
	return (*tlsValidationOptionsPtrType)(v)
}

func (*tlsValidationOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsValidationOptions)(nil)).Elem()
}

func (i *tlsValidationOptionsPtrType) ToTlsValidationOptionsPtrOutput() TlsValidationOptionsPtrOutput {
	return i.ToTlsValidationOptionsPtrOutputWithContext(context.Background())
}

func (i *tlsValidationOptionsPtrType) ToTlsValidationOptionsPtrOutputWithContext(ctx context.Context) TlsValidationOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsValidationOptionsPtrOutput)
}

type TlsValidationOptionsOutput struct{ *pulumi.OutputState }

func (TlsValidationOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsValidationOptions)(nil)).Elem()
}

func (o TlsValidationOptionsOutput) ToTlsValidationOptionsOutput() TlsValidationOptionsOutput {
	return o
}

func (o TlsValidationOptionsOutput) ToTlsValidationOptionsOutputWithContext(ctx context.Context) TlsValidationOptionsOutput {
	return o
}

func (o TlsValidationOptionsOutput) ToTlsValidationOptionsPtrOutput() TlsValidationOptionsPtrOutput {
	return o.ToTlsValidationOptionsPtrOutputWithContext(context.Background())
}

func (o TlsValidationOptionsOutput) ToTlsValidationOptionsPtrOutputWithContext(ctx context.Context) TlsValidationOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TlsValidationOptions) *TlsValidationOptions {
		return &v
	}).(TlsValidationOptionsPtrOutput)
}

func (o TlsValidationOptionsOutput) IgnoreHostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TlsValidationOptions) *string { return v.IgnoreHostname }).(pulumi.StringPtrOutput)
}

func (o TlsValidationOptionsOutput) IgnoreSignature() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TlsValidationOptions) *string { return v.IgnoreSignature }).(pulumi.StringPtrOutput)
}

type TlsValidationOptionsPtrOutput struct{ *pulumi.OutputState }

func (TlsValidationOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsValidationOptions)(nil)).Elem()
}

func (o TlsValidationOptionsPtrOutput) ToTlsValidationOptionsPtrOutput() TlsValidationOptionsPtrOutput {
	return o
}

func (o TlsValidationOptionsPtrOutput) ToTlsValidationOptionsPtrOutputWithContext(ctx context.Context) TlsValidationOptionsPtrOutput {
	return o
}

func (o TlsValidationOptionsPtrOutput) Elem() TlsValidationOptionsOutput {
	return o.ApplyT(func(v *TlsValidationOptions) TlsValidationOptions {
		if v != nil {
			return *v
		}
		var ret TlsValidationOptions
		return ret
	}).(TlsValidationOptionsOutput)
}

func (o TlsValidationOptionsPtrOutput) IgnoreHostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsValidationOptions) *string {
		if v == nil {
			return nil
		}
		return v.IgnoreHostname
	}).(pulumi.StringPtrOutput)
}

func (o TlsValidationOptionsPtrOutput) IgnoreSignature() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsValidationOptions) *string {
		if v == nil {
			return nil
		}
		return v.IgnoreSignature
	}).(pulumi.StringPtrOutput)
}

type TlsValidationOptionsResponse struct {
	IgnoreHostname  *string `pulumi:"ignoreHostname"`
	IgnoreSignature *string `pulumi:"ignoreSignature"`
}





type TlsValidationOptionsResponseInput interface {
	pulumi.Input

	ToTlsValidationOptionsResponseOutput() TlsValidationOptionsResponseOutput
	ToTlsValidationOptionsResponseOutputWithContext(context.Context) TlsValidationOptionsResponseOutput
}

type TlsValidationOptionsResponseArgs struct {
	IgnoreHostname  pulumi.StringPtrInput `pulumi:"ignoreHostname"`
	IgnoreSignature pulumi.StringPtrInput `pulumi:"ignoreSignature"`
}

func (TlsValidationOptionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsValidationOptionsResponse)(nil)).Elem()
}

func (i TlsValidationOptionsResponseArgs) ToTlsValidationOptionsResponseOutput() TlsValidationOptionsResponseOutput {
	return i.ToTlsValidationOptionsResponseOutputWithContext(context.Background())
}

func (i TlsValidationOptionsResponseArgs) ToTlsValidationOptionsResponseOutputWithContext(ctx context.Context) TlsValidationOptionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsValidationOptionsResponseOutput)
}

func (i TlsValidationOptionsResponseArgs) ToTlsValidationOptionsResponsePtrOutput() TlsValidationOptionsResponsePtrOutput {
	return i.ToTlsValidationOptionsResponsePtrOutputWithContext(context.Background())
}

func (i TlsValidationOptionsResponseArgs) ToTlsValidationOptionsResponsePtrOutputWithContext(ctx context.Context) TlsValidationOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsValidationOptionsResponseOutput).ToTlsValidationOptionsResponsePtrOutputWithContext(ctx)
}









type TlsValidationOptionsResponsePtrInput interface {
	pulumi.Input

	ToTlsValidationOptionsResponsePtrOutput() TlsValidationOptionsResponsePtrOutput
	ToTlsValidationOptionsResponsePtrOutputWithContext(context.Context) TlsValidationOptionsResponsePtrOutput
}

type tlsValidationOptionsResponsePtrType TlsValidationOptionsResponseArgs

func TlsValidationOptionsResponsePtr(v *TlsValidationOptionsResponseArgs) TlsValidationOptionsResponsePtrInput {
	return (*tlsValidationOptionsResponsePtrType)(v)
}

func (*tlsValidationOptionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsValidationOptionsResponse)(nil)).Elem()
}

func (i *tlsValidationOptionsResponsePtrType) ToTlsValidationOptionsResponsePtrOutput() TlsValidationOptionsResponsePtrOutput {
	return i.ToTlsValidationOptionsResponsePtrOutputWithContext(context.Background())
}

func (i *tlsValidationOptionsResponsePtrType) ToTlsValidationOptionsResponsePtrOutputWithContext(ctx context.Context) TlsValidationOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsValidationOptionsResponsePtrOutput)
}

type TlsValidationOptionsResponseOutput struct{ *pulumi.OutputState }

func (TlsValidationOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsValidationOptionsResponse)(nil)).Elem()
}

func (o TlsValidationOptionsResponseOutput) ToTlsValidationOptionsResponseOutput() TlsValidationOptionsResponseOutput {
	return o
}

func (o TlsValidationOptionsResponseOutput) ToTlsValidationOptionsResponseOutputWithContext(ctx context.Context) TlsValidationOptionsResponseOutput {
	return o
}

func (o TlsValidationOptionsResponseOutput) ToTlsValidationOptionsResponsePtrOutput() TlsValidationOptionsResponsePtrOutput {
	return o.ToTlsValidationOptionsResponsePtrOutputWithContext(context.Background())
}

func (o TlsValidationOptionsResponseOutput) ToTlsValidationOptionsResponsePtrOutputWithContext(ctx context.Context) TlsValidationOptionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TlsValidationOptionsResponse) *TlsValidationOptionsResponse {
		return &v
	}).(TlsValidationOptionsResponsePtrOutput)
}

func (o TlsValidationOptionsResponseOutput) IgnoreHostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TlsValidationOptionsResponse) *string { return v.IgnoreHostname }).(pulumi.StringPtrOutput)
}

func (o TlsValidationOptionsResponseOutput) IgnoreSignature() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TlsValidationOptionsResponse) *string { return v.IgnoreSignature }).(pulumi.StringPtrOutput)
}

type TlsValidationOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (TlsValidationOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsValidationOptionsResponse)(nil)).Elem()
}

func (o TlsValidationOptionsResponsePtrOutput) ToTlsValidationOptionsResponsePtrOutput() TlsValidationOptionsResponsePtrOutput {
	return o
}

func (o TlsValidationOptionsResponsePtrOutput) ToTlsValidationOptionsResponsePtrOutputWithContext(ctx context.Context) TlsValidationOptionsResponsePtrOutput {
	return o
}

func (o TlsValidationOptionsResponsePtrOutput) Elem() TlsValidationOptionsResponseOutput {
	return o.ApplyT(func(v *TlsValidationOptionsResponse) TlsValidationOptionsResponse {
		if v != nil {
			return *v
		}
		var ret TlsValidationOptionsResponse
		return ret
	}).(TlsValidationOptionsResponseOutput)
}

func (o TlsValidationOptionsResponsePtrOutput) IgnoreHostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsValidationOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.IgnoreHostname
	}).(pulumi.StringPtrOutput)
}

func (o TlsValidationOptionsResponsePtrOutput) IgnoreSignature() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsValidationOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.IgnoreSignature
	}).(pulumi.StringPtrOutput)
}

type TokenClaim struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type TokenClaimInput interface {
	pulumi.Input

	ToTokenClaimOutput() TokenClaimOutput
	ToTokenClaimOutputWithContext(context.Context) TokenClaimOutput
}

type TokenClaimArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TokenClaimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenClaim)(nil)).Elem()
}

func (i TokenClaimArgs) ToTokenClaimOutput() TokenClaimOutput {
	return i.ToTokenClaimOutputWithContext(context.Background())
}

func (i TokenClaimArgs) ToTokenClaimOutputWithContext(ctx context.Context) TokenClaimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenClaimOutput)
}





type TokenClaimArrayInput interface {
	pulumi.Input

	ToTokenClaimArrayOutput() TokenClaimArrayOutput
	ToTokenClaimArrayOutputWithContext(context.Context) TokenClaimArrayOutput
}

type TokenClaimArray []TokenClaimInput

func (TokenClaimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenClaim)(nil)).Elem()
}

func (i TokenClaimArray) ToTokenClaimArrayOutput() TokenClaimArrayOutput {
	return i.ToTokenClaimArrayOutputWithContext(context.Background())
}

func (i TokenClaimArray) ToTokenClaimArrayOutputWithContext(ctx context.Context) TokenClaimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenClaimArrayOutput)
}

type TokenClaimOutput struct{ *pulumi.OutputState }

func (TokenClaimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenClaim)(nil)).Elem()
}

func (o TokenClaimOutput) ToTokenClaimOutput() TokenClaimOutput {
	return o
}

func (o TokenClaimOutput) ToTokenClaimOutputWithContext(ctx context.Context) TokenClaimOutput {
	return o
}

func (o TokenClaimOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TokenClaim) string { return v.Name }).(pulumi.StringOutput)
}

func (o TokenClaimOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TokenClaim) string { return v.Value }).(pulumi.StringOutput)
}

type TokenClaimArrayOutput struct{ *pulumi.OutputState }

func (TokenClaimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenClaim)(nil)).Elem()
}

func (o TokenClaimArrayOutput) ToTokenClaimArrayOutput() TokenClaimArrayOutput {
	return o
}

func (o TokenClaimArrayOutput) ToTokenClaimArrayOutputWithContext(ctx context.Context) TokenClaimArrayOutput {
	return o
}

func (o TokenClaimArrayOutput) Index(i pulumi.IntInput) TokenClaimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenClaim {
		return vs[0].([]TokenClaim)[vs[1].(int)]
	}).(TokenClaimOutput)
}

type TokenClaimResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type TokenClaimResponseInput interface {
	pulumi.Input

	ToTokenClaimResponseOutput() TokenClaimResponseOutput
	ToTokenClaimResponseOutputWithContext(context.Context) TokenClaimResponseOutput
}

type TokenClaimResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TokenClaimResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenClaimResponse)(nil)).Elem()
}

func (i TokenClaimResponseArgs) ToTokenClaimResponseOutput() TokenClaimResponseOutput {
	return i.ToTokenClaimResponseOutputWithContext(context.Background())
}

func (i TokenClaimResponseArgs) ToTokenClaimResponseOutputWithContext(ctx context.Context) TokenClaimResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenClaimResponseOutput)
}





type TokenClaimResponseArrayInput interface {
	pulumi.Input

	ToTokenClaimResponseArrayOutput() TokenClaimResponseArrayOutput
	ToTokenClaimResponseArrayOutputWithContext(context.Context) TokenClaimResponseArrayOutput
}

type TokenClaimResponseArray []TokenClaimResponseInput

func (TokenClaimResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenClaimResponse)(nil)).Elem()
}

func (i TokenClaimResponseArray) ToTokenClaimResponseArrayOutput() TokenClaimResponseArrayOutput {
	return i.ToTokenClaimResponseArrayOutputWithContext(context.Background())
}

func (i TokenClaimResponseArray) ToTokenClaimResponseArrayOutputWithContext(ctx context.Context) TokenClaimResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenClaimResponseArrayOutput)
}

type TokenClaimResponseOutput struct{ *pulumi.OutputState }

func (TokenClaimResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenClaimResponse)(nil)).Elem()
}

func (o TokenClaimResponseOutput) ToTokenClaimResponseOutput() TokenClaimResponseOutput {
	return o
}

func (o TokenClaimResponseOutput) ToTokenClaimResponseOutputWithContext(ctx context.Context) TokenClaimResponseOutput {
	return o
}

func (o TokenClaimResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TokenClaimResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TokenClaimResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TokenClaimResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TokenClaimResponseArrayOutput struct{ *pulumi.OutputState }

func (TokenClaimResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenClaimResponse)(nil)).Elem()
}

func (o TokenClaimResponseArrayOutput) ToTokenClaimResponseArrayOutput() TokenClaimResponseArrayOutput {
	return o
}

func (o TokenClaimResponseArrayOutput) ToTokenClaimResponseArrayOutputWithContext(ctx context.Context) TokenClaimResponseArrayOutput {
	return o
}

func (o TokenClaimResponseArrayOutput) Index(i pulumi.IntInput) TokenClaimResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenClaimResponse {
		return vs[0].([]TokenClaimResponse)[vs[1].(int)]
	}).(TokenClaimResponseOutput)
}

type UnsecuredEndpoint struct {
	Credentials *UsernamePasswordCredentials `pulumi:"credentials"`
	Tunnel      *SecureIotDeviceRemoteTunnel `pulumi:"tunnel"`
	Type        string                       `pulumi:"type"`
	Url         string                       `pulumi:"url"`
}





type UnsecuredEndpointInput interface {
	pulumi.Input

	ToUnsecuredEndpointOutput() UnsecuredEndpointOutput
	ToUnsecuredEndpointOutputWithContext(context.Context) UnsecuredEndpointOutput
}

type UnsecuredEndpointArgs struct {
	Credentials UsernamePasswordCredentialsPtrInput `pulumi:"credentials"`
	Tunnel      SecureIotDeviceRemoteTunnelPtrInput `pulumi:"tunnel"`
	Type        pulumi.StringInput                  `pulumi:"type"`
	Url         pulumi.StringInput                  `pulumi:"url"`
}

func (UnsecuredEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UnsecuredEndpoint)(nil)).Elem()
}

func (i UnsecuredEndpointArgs) ToUnsecuredEndpointOutput() UnsecuredEndpointOutput {
	return i.ToUnsecuredEndpointOutputWithContext(context.Background())
}

func (i UnsecuredEndpointArgs) ToUnsecuredEndpointOutputWithContext(ctx context.Context) UnsecuredEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnsecuredEndpointOutput)
}

type UnsecuredEndpointOutput struct{ *pulumi.OutputState }

func (UnsecuredEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnsecuredEndpoint)(nil)).Elem()
}

func (o UnsecuredEndpointOutput) ToUnsecuredEndpointOutput() UnsecuredEndpointOutput {
	return o
}

func (o UnsecuredEndpointOutput) ToUnsecuredEndpointOutputWithContext(ctx context.Context) UnsecuredEndpointOutput {
	return o
}

func (o UnsecuredEndpointOutput) Credentials() UsernamePasswordCredentialsPtrOutput {
	return o.ApplyT(func(v UnsecuredEndpoint) *UsernamePasswordCredentials { return v.Credentials }).(UsernamePasswordCredentialsPtrOutput)
}

func (o UnsecuredEndpointOutput) Tunnel() SecureIotDeviceRemoteTunnelPtrOutput {
	return o.ApplyT(func(v UnsecuredEndpoint) *SecureIotDeviceRemoteTunnel { return v.Tunnel }).(SecureIotDeviceRemoteTunnelPtrOutput)
}

func (o UnsecuredEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v UnsecuredEndpoint) string { return v.Type }).(pulumi.StringOutput)
}

func (o UnsecuredEndpointOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v UnsecuredEndpoint) string { return v.Url }).(pulumi.StringOutput)
}

type UnsecuredEndpointResponse struct {
	Credentials *UsernamePasswordCredentialsResponse `pulumi:"credentials"`
	Tunnel      *SecureIotDeviceRemoteTunnelResponse `pulumi:"tunnel"`
	Type        string                               `pulumi:"type"`
	Url         string                               `pulumi:"url"`
}





type UnsecuredEndpointResponseInput interface {
	pulumi.Input

	ToUnsecuredEndpointResponseOutput() UnsecuredEndpointResponseOutput
	ToUnsecuredEndpointResponseOutputWithContext(context.Context) UnsecuredEndpointResponseOutput
}

type UnsecuredEndpointResponseArgs struct {
	Credentials UsernamePasswordCredentialsResponsePtrInput `pulumi:"credentials"`
	Tunnel      SecureIotDeviceRemoteTunnelResponsePtrInput `pulumi:"tunnel"`
	Type        pulumi.StringInput                          `pulumi:"type"`
	Url         pulumi.StringInput                          `pulumi:"url"`
}

func (UnsecuredEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UnsecuredEndpointResponse)(nil)).Elem()
}

func (i UnsecuredEndpointResponseArgs) ToUnsecuredEndpointResponseOutput() UnsecuredEndpointResponseOutput {
	return i.ToUnsecuredEndpointResponseOutputWithContext(context.Background())
}

func (i UnsecuredEndpointResponseArgs) ToUnsecuredEndpointResponseOutputWithContext(ctx context.Context) UnsecuredEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnsecuredEndpointResponseOutput)
}

type UnsecuredEndpointResponseOutput struct{ *pulumi.OutputState }

func (UnsecuredEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnsecuredEndpointResponse)(nil)).Elem()
}

func (o UnsecuredEndpointResponseOutput) ToUnsecuredEndpointResponseOutput() UnsecuredEndpointResponseOutput {
	return o
}

func (o UnsecuredEndpointResponseOutput) ToUnsecuredEndpointResponseOutputWithContext(ctx context.Context) UnsecuredEndpointResponseOutput {
	return o
}

func (o UnsecuredEndpointResponseOutput) Credentials() UsernamePasswordCredentialsResponsePtrOutput {
	return o.ApplyT(func(v UnsecuredEndpointResponse) *UsernamePasswordCredentialsResponse { return v.Credentials }).(UsernamePasswordCredentialsResponsePtrOutput)
}

func (o UnsecuredEndpointResponseOutput) Tunnel() SecureIotDeviceRemoteTunnelResponsePtrOutput {
	return o.ApplyT(func(v UnsecuredEndpointResponse) *SecureIotDeviceRemoteTunnelResponse { return v.Tunnel }).(SecureIotDeviceRemoteTunnelResponsePtrOutput)
}

func (o UnsecuredEndpointResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v UnsecuredEndpointResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o UnsecuredEndpointResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v UnsecuredEndpointResponse) string { return v.Url }).(pulumi.StringOutput)
}

type UserAssignedManagedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type UserAssignedManagedIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedManagedIdentityResponseOutput() UserAssignedManagedIdentityResponseOutput
	ToUserAssignedManagedIdentityResponseOutputWithContext(context.Context) UserAssignedManagedIdentityResponseOutput
}

type UserAssignedManagedIdentityResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (UserAssignedManagedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedManagedIdentityResponseArgs) ToUserAssignedManagedIdentityResponseOutput() UserAssignedManagedIdentityResponseOutput {
	return i.ToUserAssignedManagedIdentityResponseOutputWithContext(context.Background())
}

func (i UserAssignedManagedIdentityResponseArgs) ToUserAssignedManagedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedManagedIdentityResponseOutput)
}





type UserAssignedManagedIdentityResponseMapInput interface {
	pulumi.Input

	ToUserAssignedManagedIdentityResponseMapOutput() UserAssignedManagedIdentityResponseMapOutput
	ToUserAssignedManagedIdentityResponseMapOutputWithContext(context.Context) UserAssignedManagedIdentityResponseMapOutput
}

type UserAssignedManagedIdentityResponseMap map[string]UserAssignedManagedIdentityResponseInput

func (UserAssignedManagedIdentityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedManagedIdentityResponseMap) ToUserAssignedManagedIdentityResponseMapOutput() UserAssignedManagedIdentityResponseMapOutput {
	return i.ToUserAssignedManagedIdentityResponseMapOutputWithContext(context.Background())
}

func (i UserAssignedManagedIdentityResponseMap) ToUserAssignedManagedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedManagedIdentityResponseMapOutput)
}

type UserAssignedManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedManagedIdentityResponseOutput) ToUserAssignedManagedIdentityResponseOutput() UserAssignedManagedIdentityResponseOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseOutput) ToUserAssignedManagedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedManagedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedManagedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedManagedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedManagedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedManagedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedManagedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedManagedIdentityResponseMapOutput) ToUserAssignedManagedIdentityResponseMapOutput() UserAssignedManagedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseMapOutput) ToUserAssignedManagedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedManagedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedManagedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedManagedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedManagedIdentityResponse {
		return vs[0].(map[string]UserAssignedManagedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedManagedIdentityResponseOutput)
}

type UsernamePasswordCredentials struct {
	Password string `pulumi:"password"`
	Type     string `pulumi:"type"`
	Username string `pulumi:"username"`
}





type UsernamePasswordCredentialsInput interface {
	pulumi.Input

	ToUsernamePasswordCredentialsOutput() UsernamePasswordCredentialsOutput
	ToUsernamePasswordCredentialsOutputWithContext(context.Context) UsernamePasswordCredentialsOutput
}

type UsernamePasswordCredentialsArgs struct {
	Password pulumi.StringInput `pulumi:"password"`
	Type     pulumi.StringInput `pulumi:"type"`
	Username pulumi.StringInput `pulumi:"username"`
}

func (UsernamePasswordCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UsernamePasswordCredentials)(nil)).Elem()
}

func (i UsernamePasswordCredentialsArgs) ToUsernamePasswordCredentialsOutput() UsernamePasswordCredentialsOutput {
	return i.ToUsernamePasswordCredentialsOutputWithContext(context.Background())
}

func (i UsernamePasswordCredentialsArgs) ToUsernamePasswordCredentialsOutputWithContext(ctx context.Context) UsernamePasswordCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsernamePasswordCredentialsOutput)
}

func (i UsernamePasswordCredentialsArgs) ToUsernamePasswordCredentialsPtrOutput() UsernamePasswordCredentialsPtrOutput {
	return i.ToUsernamePasswordCredentialsPtrOutputWithContext(context.Background())
}

func (i UsernamePasswordCredentialsArgs) ToUsernamePasswordCredentialsPtrOutputWithContext(ctx context.Context) UsernamePasswordCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsernamePasswordCredentialsOutput).ToUsernamePasswordCredentialsPtrOutputWithContext(ctx)
}









type UsernamePasswordCredentialsPtrInput interface {
	pulumi.Input

	ToUsernamePasswordCredentialsPtrOutput() UsernamePasswordCredentialsPtrOutput
	ToUsernamePasswordCredentialsPtrOutputWithContext(context.Context) UsernamePasswordCredentialsPtrOutput
}

type usernamePasswordCredentialsPtrType UsernamePasswordCredentialsArgs

func UsernamePasswordCredentialsPtr(v *UsernamePasswordCredentialsArgs) UsernamePasswordCredentialsPtrInput {
	return (*usernamePasswordCredentialsPtrType)(v)
}

func (*usernamePasswordCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UsernamePasswordCredentials)(nil)).Elem()
}

func (i *usernamePasswordCredentialsPtrType) ToUsernamePasswordCredentialsPtrOutput() UsernamePasswordCredentialsPtrOutput {
	return i.ToUsernamePasswordCredentialsPtrOutputWithContext(context.Background())
}

func (i *usernamePasswordCredentialsPtrType) ToUsernamePasswordCredentialsPtrOutputWithContext(ctx context.Context) UsernamePasswordCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsernamePasswordCredentialsPtrOutput)
}

type UsernamePasswordCredentialsOutput struct{ *pulumi.OutputState }

func (UsernamePasswordCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UsernamePasswordCredentials)(nil)).Elem()
}

func (o UsernamePasswordCredentialsOutput) ToUsernamePasswordCredentialsOutput() UsernamePasswordCredentialsOutput {
	return o
}

func (o UsernamePasswordCredentialsOutput) ToUsernamePasswordCredentialsOutputWithContext(ctx context.Context) UsernamePasswordCredentialsOutput {
	return o
}

func (o UsernamePasswordCredentialsOutput) ToUsernamePasswordCredentialsPtrOutput() UsernamePasswordCredentialsPtrOutput {
	return o.ToUsernamePasswordCredentialsPtrOutputWithContext(context.Background())
}

func (o UsernamePasswordCredentialsOutput) ToUsernamePasswordCredentialsPtrOutputWithContext(ctx context.Context) UsernamePasswordCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UsernamePasswordCredentials) *UsernamePasswordCredentials {
		return &v
	}).(UsernamePasswordCredentialsPtrOutput)
}

func (o UsernamePasswordCredentialsOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v UsernamePasswordCredentials) string { return v.Password }).(pulumi.StringOutput)
}

func (o UsernamePasswordCredentialsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v UsernamePasswordCredentials) string { return v.Type }).(pulumi.StringOutput)
}

func (o UsernamePasswordCredentialsOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v UsernamePasswordCredentials) string { return v.Username }).(pulumi.StringOutput)
}

type UsernamePasswordCredentialsPtrOutput struct{ *pulumi.OutputState }

func (UsernamePasswordCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UsernamePasswordCredentials)(nil)).Elem()
}

func (o UsernamePasswordCredentialsPtrOutput) ToUsernamePasswordCredentialsPtrOutput() UsernamePasswordCredentialsPtrOutput {
	return o
}

func (o UsernamePasswordCredentialsPtrOutput) ToUsernamePasswordCredentialsPtrOutputWithContext(ctx context.Context) UsernamePasswordCredentialsPtrOutput {
	return o
}

func (o UsernamePasswordCredentialsPtrOutput) Elem() UsernamePasswordCredentialsOutput {
	return o.ApplyT(func(v *UsernamePasswordCredentials) UsernamePasswordCredentials {
		if v != nil {
			return *v
		}
		var ret UsernamePasswordCredentials
		return ret
	}).(UsernamePasswordCredentialsOutput)
}

func (o UsernamePasswordCredentialsPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UsernamePasswordCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o UsernamePasswordCredentialsPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UsernamePasswordCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o UsernamePasswordCredentialsPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UsernamePasswordCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type UsernamePasswordCredentialsResponse struct {
	Password string `pulumi:"password"`
	Type     string `pulumi:"type"`
	Username string `pulumi:"username"`
}





type UsernamePasswordCredentialsResponseInput interface {
	pulumi.Input

	ToUsernamePasswordCredentialsResponseOutput() UsernamePasswordCredentialsResponseOutput
	ToUsernamePasswordCredentialsResponseOutputWithContext(context.Context) UsernamePasswordCredentialsResponseOutput
}

type UsernamePasswordCredentialsResponseArgs struct {
	Password pulumi.StringInput `pulumi:"password"`
	Type     pulumi.StringInput `pulumi:"type"`
	Username pulumi.StringInput `pulumi:"username"`
}

func (UsernamePasswordCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UsernamePasswordCredentialsResponse)(nil)).Elem()
}

func (i UsernamePasswordCredentialsResponseArgs) ToUsernamePasswordCredentialsResponseOutput() UsernamePasswordCredentialsResponseOutput {
	return i.ToUsernamePasswordCredentialsResponseOutputWithContext(context.Background())
}

func (i UsernamePasswordCredentialsResponseArgs) ToUsernamePasswordCredentialsResponseOutputWithContext(ctx context.Context) UsernamePasswordCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsernamePasswordCredentialsResponseOutput)
}

func (i UsernamePasswordCredentialsResponseArgs) ToUsernamePasswordCredentialsResponsePtrOutput() UsernamePasswordCredentialsResponsePtrOutput {
	return i.ToUsernamePasswordCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i UsernamePasswordCredentialsResponseArgs) ToUsernamePasswordCredentialsResponsePtrOutputWithContext(ctx context.Context) UsernamePasswordCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsernamePasswordCredentialsResponseOutput).ToUsernamePasswordCredentialsResponsePtrOutputWithContext(ctx)
}









type UsernamePasswordCredentialsResponsePtrInput interface {
	pulumi.Input

	ToUsernamePasswordCredentialsResponsePtrOutput() UsernamePasswordCredentialsResponsePtrOutput
	ToUsernamePasswordCredentialsResponsePtrOutputWithContext(context.Context) UsernamePasswordCredentialsResponsePtrOutput
}

type usernamePasswordCredentialsResponsePtrType UsernamePasswordCredentialsResponseArgs

func UsernamePasswordCredentialsResponsePtr(v *UsernamePasswordCredentialsResponseArgs) UsernamePasswordCredentialsResponsePtrInput {
	return (*usernamePasswordCredentialsResponsePtrType)(v)
}

func (*usernamePasswordCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UsernamePasswordCredentialsResponse)(nil)).Elem()
}

func (i *usernamePasswordCredentialsResponsePtrType) ToUsernamePasswordCredentialsResponsePtrOutput() UsernamePasswordCredentialsResponsePtrOutput {
	return i.ToUsernamePasswordCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *usernamePasswordCredentialsResponsePtrType) ToUsernamePasswordCredentialsResponsePtrOutputWithContext(ctx context.Context) UsernamePasswordCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsernamePasswordCredentialsResponsePtrOutput)
}

type UsernamePasswordCredentialsResponseOutput struct{ *pulumi.OutputState }

func (UsernamePasswordCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UsernamePasswordCredentialsResponse)(nil)).Elem()
}

func (o UsernamePasswordCredentialsResponseOutput) ToUsernamePasswordCredentialsResponseOutput() UsernamePasswordCredentialsResponseOutput {
	return o
}

func (o UsernamePasswordCredentialsResponseOutput) ToUsernamePasswordCredentialsResponseOutputWithContext(ctx context.Context) UsernamePasswordCredentialsResponseOutput {
	return o
}

func (o UsernamePasswordCredentialsResponseOutput) ToUsernamePasswordCredentialsResponsePtrOutput() UsernamePasswordCredentialsResponsePtrOutput {
	return o.ToUsernamePasswordCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o UsernamePasswordCredentialsResponseOutput) ToUsernamePasswordCredentialsResponsePtrOutputWithContext(ctx context.Context) UsernamePasswordCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UsernamePasswordCredentialsResponse) *UsernamePasswordCredentialsResponse {
		return &v
	}).(UsernamePasswordCredentialsResponsePtrOutput)
}

func (o UsernamePasswordCredentialsResponseOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v UsernamePasswordCredentialsResponse) string { return v.Password }).(pulumi.StringOutput)
}

func (o UsernamePasswordCredentialsResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v UsernamePasswordCredentialsResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o UsernamePasswordCredentialsResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v UsernamePasswordCredentialsResponse) string { return v.Username }).(pulumi.StringOutput)
}

type UsernamePasswordCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (UsernamePasswordCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UsernamePasswordCredentialsResponse)(nil)).Elem()
}

func (o UsernamePasswordCredentialsResponsePtrOutput) ToUsernamePasswordCredentialsResponsePtrOutput() UsernamePasswordCredentialsResponsePtrOutput {
	return o
}

func (o UsernamePasswordCredentialsResponsePtrOutput) ToUsernamePasswordCredentialsResponsePtrOutputWithContext(ctx context.Context) UsernamePasswordCredentialsResponsePtrOutput {
	return o
}

func (o UsernamePasswordCredentialsResponsePtrOutput) Elem() UsernamePasswordCredentialsResponseOutput {
	return o.ApplyT(func(v *UsernamePasswordCredentialsResponse) UsernamePasswordCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret UsernamePasswordCredentialsResponse
		return ret
	}).(UsernamePasswordCredentialsResponseOutput)
}

func (o UsernamePasswordCredentialsResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UsernamePasswordCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o UsernamePasswordCredentialsResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UsernamePasswordCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o UsernamePasswordCredentialsResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UsernamePasswordCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type VideoAnalyzerIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type VideoAnalyzerIdentityInput interface {
	pulumi.Input

	ToVideoAnalyzerIdentityOutput() VideoAnalyzerIdentityOutput
	ToVideoAnalyzerIdentityOutputWithContext(context.Context) VideoAnalyzerIdentityOutput
}

type VideoAnalyzerIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (VideoAnalyzerIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoAnalyzerIdentity)(nil)).Elem()
}

func (i VideoAnalyzerIdentityArgs) ToVideoAnalyzerIdentityOutput() VideoAnalyzerIdentityOutput {
	return i.ToVideoAnalyzerIdentityOutputWithContext(context.Background())
}

func (i VideoAnalyzerIdentityArgs) ToVideoAnalyzerIdentityOutputWithContext(ctx context.Context) VideoAnalyzerIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerIdentityOutput)
}

func (i VideoAnalyzerIdentityArgs) ToVideoAnalyzerIdentityPtrOutput() VideoAnalyzerIdentityPtrOutput {
	return i.ToVideoAnalyzerIdentityPtrOutputWithContext(context.Background())
}

func (i VideoAnalyzerIdentityArgs) ToVideoAnalyzerIdentityPtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerIdentityOutput).ToVideoAnalyzerIdentityPtrOutputWithContext(ctx)
}









type VideoAnalyzerIdentityPtrInput interface {
	pulumi.Input

	ToVideoAnalyzerIdentityPtrOutput() VideoAnalyzerIdentityPtrOutput
	ToVideoAnalyzerIdentityPtrOutputWithContext(context.Context) VideoAnalyzerIdentityPtrOutput
}

type videoAnalyzerIdentityPtrType VideoAnalyzerIdentityArgs

func VideoAnalyzerIdentityPtr(v *VideoAnalyzerIdentityArgs) VideoAnalyzerIdentityPtrInput {
	return (*videoAnalyzerIdentityPtrType)(v)
}

func (*videoAnalyzerIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoAnalyzerIdentity)(nil)).Elem()
}

func (i *videoAnalyzerIdentityPtrType) ToVideoAnalyzerIdentityPtrOutput() VideoAnalyzerIdentityPtrOutput {
	return i.ToVideoAnalyzerIdentityPtrOutputWithContext(context.Background())
}

func (i *videoAnalyzerIdentityPtrType) ToVideoAnalyzerIdentityPtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerIdentityPtrOutput)
}

type VideoAnalyzerIdentityOutput struct{ *pulumi.OutputState }

func (VideoAnalyzerIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoAnalyzerIdentity)(nil)).Elem()
}

func (o VideoAnalyzerIdentityOutput) ToVideoAnalyzerIdentityOutput() VideoAnalyzerIdentityOutput {
	return o
}

func (o VideoAnalyzerIdentityOutput) ToVideoAnalyzerIdentityOutputWithContext(ctx context.Context) VideoAnalyzerIdentityOutput {
	return o
}

func (o VideoAnalyzerIdentityOutput) ToVideoAnalyzerIdentityPtrOutput() VideoAnalyzerIdentityPtrOutput {
	return o.ToVideoAnalyzerIdentityPtrOutputWithContext(context.Background())
}

func (o VideoAnalyzerIdentityOutput) ToVideoAnalyzerIdentityPtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoAnalyzerIdentity) *VideoAnalyzerIdentity {
		return &v
	}).(VideoAnalyzerIdentityPtrOutput)
}

func (o VideoAnalyzerIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoAnalyzerIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o VideoAnalyzerIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v VideoAnalyzerIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type VideoAnalyzerIdentityPtrOutput struct{ *pulumi.OutputState }

func (VideoAnalyzerIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoAnalyzerIdentity)(nil)).Elem()
}

func (o VideoAnalyzerIdentityPtrOutput) ToVideoAnalyzerIdentityPtrOutput() VideoAnalyzerIdentityPtrOutput {
	return o
}

func (o VideoAnalyzerIdentityPtrOutput) ToVideoAnalyzerIdentityPtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityPtrOutput {
	return o
}

func (o VideoAnalyzerIdentityPtrOutput) Elem() VideoAnalyzerIdentityOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentity) VideoAnalyzerIdentity {
		if v != nil {
			return *v
		}
		var ret VideoAnalyzerIdentity
		return ret
	}).(VideoAnalyzerIdentityOutput)
}

func (o VideoAnalyzerIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VideoAnalyzerIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type VideoAnalyzerIdentityResponse struct {
	Type                   string                                         `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedManagedIdentityResponse `pulumi:"userAssignedIdentities"`
}





type VideoAnalyzerIdentityResponseInput interface {
	pulumi.Input

	ToVideoAnalyzerIdentityResponseOutput() VideoAnalyzerIdentityResponseOutput
	ToVideoAnalyzerIdentityResponseOutputWithContext(context.Context) VideoAnalyzerIdentityResponseOutput
}

type VideoAnalyzerIdentityResponseArgs struct {
	Type                   pulumi.StringInput                          `pulumi:"type"`
	UserAssignedIdentities UserAssignedManagedIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (VideoAnalyzerIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoAnalyzerIdentityResponse)(nil)).Elem()
}

func (i VideoAnalyzerIdentityResponseArgs) ToVideoAnalyzerIdentityResponseOutput() VideoAnalyzerIdentityResponseOutput {
	return i.ToVideoAnalyzerIdentityResponseOutputWithContext(context.Background())
}

func (i VideoAnalyzerIdentityResponseArgs) ToVideoAnalyzerIdentityResponseOutputWithContext(ctx context.Context) VideoAnalyzerIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerIdentityResponseOutput)
}

func (i VideoAnalyzerIdentityResponseArgs) ToVideoAnalyzerIdentityResponsePtrOutput() VideoAnalyzerIdentityResponsePtrOutput {
	return i.ToVideoAnalyzerIdentityResponsePtrOutputWithContext(context.Background())
}

func (i VideoAnalyzerIdentityResponseArgs) ToVideoAnalyzerIdentityResponsePtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerIdentityResponseOutput).ToVideoAnalyzerIdentityResponsePtrOutputWithContext(ctx)
}









type VideoAnalyzerIdentityResponsePtrInput interface {
	pulumi.Input

	ToVideoAnalyzerIdentityResponsePtrOutput() VideoAnalyzerIdentityResponsePtrOutput
	ToVideoAnalyzerIdentityResponsePtrOutputWithContext(context.Context) VideoAnalyzerIdentityResponsePtrOutput
}

type videoAnalyzerIdentityResponsePtrType VideoAnalyzerIdentityResponseArgs

func VideoAnalyzerIdentityResponsePtr(v *VideoAnalyzerIdentityResponseArgs) VideoAnalyzerIdentityResponsePtrInput {
	return (*videoAnalyzerIdentityResponsePtrType)(v)
}

func (*videoAnalyzerIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoAnalyzerIdentityResponse)(nil)).Elem()
}

func (i *videoAnalyzerIdentityResponsePtrType) ToVideoAnalyzerIdentityResponsePtrOutput() VideoAnalyzerIdentityResponsePtrOutput {
	return i.ToVideoAnalyzerIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *videoAnalyzerIdentityResponsePtrType) ToVideoAnalyzerIdentityResponsePtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerIdentityResponsePtrOutput)
}

type VideoAnalyzerIdentityResponseOutput struct{ *pulumi.OutputState }

func (VideoAnalyzerIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoAnalyzerIdentityResponse)(nil)).Elem()
}

func (o VideoAnalyzerIdentityResponseOutput) ToVideoAnalyzerIdentityResponseOutput() VideoAnalyzerIdentityResponseOutput {
	return o
}

func (o VideoAnalyzerIdentityResponseOutput) ToVideoAnalyzerIdentityResponseOutputWithContext(ctx context.Context) VideoAnalyzerIdentityResponseOutput {
	return o
}

func (o VideoAnalyzerIdentityResponseOutput) ToVideoAnalyzerIdentityResponsePtrOutput() VideoAnalyzerIdentityResponsePtrOutput {
	return o.ToVideoAnalyzerIdentityResponsePtrOutputWithContext(context.Background())
}

func (o VideoAnalyzerIdentityResponseOutput) ToVideoAnalyzerIdentityResponsePtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoAnalyzerIdentityResponse) *VideoAnalyzerIdentityResponse {
		return &v
	}).(VideoAnalyzerIdentityResponsePtrOutput)
}

func (o VideoAnalyzerIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoAnalyzerIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VideoAnalyzerIdentityResponseOutput) UserAssignedIdentities() UserAssignedManagedIdentityResponseMapOutput {
	return o.ApplyT(func(v VideoAnalyzerIdentityResponse) map[string]UserAssignedManagedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedManagedIdentityResponseMapOutput)
}

type VideoAnalyzerIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoAnalyzerIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoAnalyzerIdentityResponse)(nil)).Elem()
}

func (o VideoAnalyzerIdentityResponsePtrOutput) ToVideoAnalyzerIdentityResponsePtrOutput() VideoAnalyzerIdentityResponsePtrOutput {
	return o
}

func (o VideoAnalyzerIdentityResponsePtrOutput) ToVideoAnalyzerIdentityResponsePtrOutputWithContext(ctx context.Context) VideoAnalyzerIdentityResponsePtrOutput {
	return o
}

func (o VideoAnalyzerIdentityResponsePtrOutput) Elem() VideoAnalyzerIdentityResponseOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentityResponse) VideoAnalyzerIdentityResponse {
		if v != nil {
			return *v
		}
		var ret VideoAnalyzerIdentityResponse
		return ret
	}).(VideoAnalyzerIdentityResponseOutput)
}

func (o VideoAnalyzerIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VideoAnalyzerIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedManagedIdentityResponseMapOutput {
	return o.ApplyT(func(v *VideoAnalyzerIdentityResponse) map[string]UserAssignedManagedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedManagedIdentityResponseMapOutput)
}

type VideoCreationProperties struct {
	Description     *string `pulumi:"description"`
	RetentionPeriod *string `pulumi:"retentionPeriod"`
	SegmentLength   *string `pulumi:"segmentLength"`
	Title           *string `pulumi:"title"`
}





type VideoCreationPropertiesInput interface {
	pulumi.Input

	ToVideoCreationPropertiesOutput() VideoCreationPropertiesOutput
	ToVideoCreationPropertiesOutputWithContext(context.Context) VideoCreationPropertiesOutput
}

type VideoCreationPropertiesArgs struct {
	Description     pulumi.StringPtrInput `pulumi:"description"`
	RetentionPeriod pulumi.StringPtrInput `pulumi:"retentionPeriod"`
	SegmentLength   pulumi.StringPtrInput `pulumi:"segmentLength"`
	Title           pulumi.StringPtrInput `pulumi:"title"`
}

func (VideoCreationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoCreationProperties)(nil)).Elem()
}

func (i VideoCreationPropertiesArgs) ToVideoCreationPropertiesOutput() VideoCreationPropertiesOutput {
	return i.ToVideoCreationPropertiesOutputWithContext(context.Background())
}

func (i VideoCreationPropertiesArgs) ToVideoCreationPropertiesOutputWithContext(ctx context.Context) VideoCreationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoCreationPropertiesOutput)
}

func (i VideoCreationPropertiesArgs) ToVideoCreationPropertiesPtrOutput() VideoCreationPropertiesPtrOutput {
	return i.ToVideoCreationPropertiesPtrOutputWithContext(context.Background())
}

func (i VideoCreationPropertiesArgs) ToVideoCreationPropertiesPtrOutputWithContext(ctx context.Context) VideoCreationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoCreationPropertiesOutput).ToVideoCreationPropertiesPtrOutputWithContext(ctx)
}









type VideoCreationPropertiesPtrInput interface {
	pulumi.Input

	ToVideoCreationPropertiesPtrOutput() VideoCreationPropertiesPtrOutput
	ToVideoCreationPropertiesPtrOutputWithContext(context.Context) VideoCreationPropertiesPtrOutput
}

type videoCreationPropertiesPtrType VideoCreationPropertiesArgs

func VideoCreationPropertiesPtr(v *VideoCreationPropertiesArgs) VideoCreationPropertiesPtrInput {
	return (*videoCreationPropertiesPtrType)(v)
}

func (*videoCreationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoCreationProperties)(nil)).Elem()
}

func (i *videoCreationPropertiesPtrType) ToVideoCreationPropertiesPtrOutput() VideoCreationPropertiesPtrOutput {
	return i.ToVideoCreationPropertiesPtrOutputWithContext(context.Background())
}

func (i *videoCreationPropertiesPtrType) ToVideoCreationPropertiesPtrOutputWithContext(ctx context.Context) VideoCreationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoCreationPropertiesPtrOutput)
}

type VideoCreationPropertiesOutput struct{ *pulumi.OutputState }

func (VideoCreationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoCreationProperties)(nil)).Elem()
}

func (o VideoCreationPropertiesOutput) ToVideoCreationPropertiesOutput() VideoCreationPropertiesOutput {
	return o
}

func (o VideoCreationPropertiesOutput) ToVideoCreationPropertiesOutputWithContext(ctx context.Context) VideoCreationPropertiesOutput {
	return o
}

func (o VideoCreationPropertiesOutput) ToVideoCreationPropertiesPtrOutput() VideoCreationPropertiesPtrOutput {
	return o.ToVideoCreationPropertiesPtrOutputWithContext(context.Background())
}

func (o VideoCreationPropertiesOutput) ToVideoCreationPropertiesPtrOutputWithContext(ctx context.Context) VideoCreationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoCreationProperties) *VideoCreationProperties {
		return &v
	}).(VideoCreationPropertiesPtrOutput)
}

func (o VideoCreationPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationProperties) *string { return v.RetentionPeriod }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationProperties) *string { return v.SegmentLength }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationProperties) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type VideoCreationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VideoCreationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoCreationProperties)(nil)).Elem()
}

func (o VideoCreationPropertiesPtrOutput) ToVideoCreationPropertiesPtrOutput() VideoCreationPropertiesPtrOutput {
	return o
}

func (o VideoCreationPropertiesPtrOutput) ToVideoCreationPropertiesPtrOutputWithContext(ctx context.Context) VideoCreationPropertiesPtrOutput {
	return o
}

func (o VideoCreationPropertiesPtrOutput) Elem() VideoCreationPropertiesOutput {
	return o.ApplyT(func(v *VideoCreationProperties) VideoCreationProperties {
		if v != nil {
			return *v
		}
		var ret VideoCreationProperties
		return ret
	}).(VideoCreationPropertiesOutput)
}

func (o VideoCreationPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesPtrOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.RetentionPeriod
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesPtrOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.SegmentLength
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationProperties) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

type VideoCreationPropertiesResponse struct {
	Description     *string `pulumi:"description"`
	RetentionPeriod *string `pulumi:"retentionPeriod"`
	SegmentLength   *string `pulumi:"segmentLength"`
	Title           *string `pulumi:"title"`
}





type VideoCreationPropertiesResponseInput interface {
	pulumi.Input

	ToVideoCreationPropertiesResponseOutput() VideoCreationPropertiesResponseOutput
	ToVideoCreationPropertiesResponseOutputWithContext(context.Context) VideoCreationPropertiesResponseOutput
}

type VideoCreationPropertiesResponseArgs struct {
	Description     pulumi.StringPtrInput `pulumi:"description"`
	RetentionPeriod pulumi.StringPtrInput `pulumi:"retentionPeriod"`
	SegmentLength   pulumi.StringPtrInput `pulumi:"segmentLength"`
	Title           pulumi.StringPtrInput `pulumi:"title"`
}

func (VideoCreationPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoCreationPropertiesResponse)(nil)).Elem()
}

func (i VideoCreationPropertiesResponseArgs) ToVideoCreationPropertiesResponseOutput() VideoCreationPropertiesResponseOutput {
	return i.ToVideoCreationPropertiesResponseOutputWithContext(context.Background())
}

func (i VideoCreationPropertiesResponseArgs) ToVideoCreationPropertiesResponseOutputWithContext(ctx context.Context) VideoCreationPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoCreationPropertiesResponseOutput)
}

func (i VideoCreationPropertiesResponseArgs) ToVideoCreationPropertiesResponsePtrOutput() VideoCreationPropertiesResponsePtrOutput {
	return i.ToVideoCreationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i VideoCreationPropertiesResponseArgs) ToVideoCreationPropertiesResponsePtrOutputWithContext(ctx context.Context) VideoCreationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoCreationPropertiesResponseOutput).ToVideoCreationPropertiesResponsePtrOutputWithContext(ctx)
}









type VideoCreationPropertiesResponsePtrInput interface {
	pulumi.Input

	ToVideoCreationPropertiesResponsePtrOutput() VideoCreationPropertiesResponsePtrOutput
	ToVideoCreationPropertiesResponsePtrOutputWithContext(context.Context) VideoCreationPropertiesResponsePtrOutput
}

type videoCreationPropertiesResponsePtrType VideoCreationPropertiesResponseArgs

func VideoCreationPropertiesResponsePtr(v *VideoCreationPropertiesResponseArgs) VideoCreationPropertiesResponsePtrInput {
	return (*videoCreationPropertiesResponsePtrType)(v)
}

func (*videoCreationPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoCreationPropertiesResponse)(nil)).Elem()
}

func (i *videoCreationPropertiesResponsePtrType) ToVideoCreationPropertiesResponsePtrOutput() VideoCreationPropertiesResponsePtrOutput {
	return i.ToVideoCreationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *videoCreationPropertiesResponsePtrType) ToVideoCreationPropertiesResponsePtrOutputWithContext(ctx context.Context) VideoCreationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoCreationPropertiesResponsePtrOutput)
}

type VideoCreationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VideoCreationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoCreationPropertiesResponse)(nil)).Elem()
}

func (o VideoCreationPropertiesResponseOutput) ToVideoCreationPropertiesResponseOutput() VideoCreationPropertiesResponseOutput {
	return o
}

func (o VideoCreationPropertiesResponseOutput) ToVideoCreationPropertiesResponseOutputWithContext(ctx context.Context) VideoCreationPropertiesResponseOutput {
	return o
}

func (o VideoCreationPropertiesResponseOutput) ToVideoCreationPropertiesResponsePtrOutput() VideoCreationPropertiesResponsePtrOutput {
	return o.ToVideoCreationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o VideoCreationPropertiesResponseOutput) ToVideoCreationPropertiesResponsePtrOutputWithContext(ctx context.Context) VideoCreationPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoCreationPropertiesResponse) *VideoCreationPropertiesResponse {
		return &v
	}).(VideoCreationPropertiesResponsePtrOutput)
}

func (o VideoCreationPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponseOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationPropertiesResponse) *string { return v.RetentionPeriod }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponseOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationPropertiesResponse) *string { return v.SegmentLength }).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponseOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoCreationPropertiesResponse) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type VideoCreationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoCreationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoCreationPropertiesResponse)(nil)).Elem()
}

func (o VideoCreationPropertiesResponsePtrOutput) ToVideoCreationPropertiesResponsePtrOutput() VideoCreationPropertiesResponsePtrOutput {
	return o
}

func (o VideoCreationPropertiesResponsePtrOutput) ToVideoCreationPropertiesResponsePtrOutputWithContext(ctx context.Context) VideoCreationPropertiesResponsePtrOutput {
	return o
}

func (o VideoCreationPropertiesResponsePtrOutput) Elem() VideoCreationPropertiesResponseOutput {
	return o.ApplyT(func(v *VideoCreationPropertiesResponse) VideoCreationPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VideoCreationPropertiesResponse
		return ret
	}).(VideoCreationPropertiesResponseOutput)
}

func (o VideoCreationPropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponsePtrOutput) RetentionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RetentionPeriod
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponsePtrOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SegmentLength
	}).(pulumi.StringPtrOutput)
}

func (o VideoCreationPropertiesResponsePtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoCreationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

type VideoEncoderH264 struct {
	BitrateKbps *string     `pulumi:"bitrateKbps"`
	FrameRate   *string     `pulumi:"frameRate"`
	Scale       *VideoScale `pulumi:"scale"`
	Type        string      `pulumi:"type"`
}





type VideoEncoderH264Input interface {
	pulumi.Input

	ToVideoEncoderH264Output() VideoEncoderH264Output
	ToVideoEncoderH264OutputWithContext(context.Context) VideoEncoderH264Output
}

type VideoEncoderH264Args struct {
	BitrateKbps pulumi.StringPtrInput `pulumi:"bitrateKbps"`
	FrameRate   pulumi.StringPtrInput `pulumi:"frameRate"`
	Scale       VideoScalePtrInput    `pulumi:"scale"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (VideoEncoderH264Args) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoEncoderH264)(nil)).Elem()
}

func (i VideoEncoderH264Args) ToVideoEncoderH264Output() VideoEncoderH264Output {
	return i.ToVideoEncoderH264OutputWithContext(context.Background())
}

func (i VideoEncoderH264Args) ToVideoEncoderH264OutputWithContext(ctx context.Context) VideoEncoderH264Output {
	return pulumi.ToOutputWithContext(ctx, i).(VideoEncoderH264Output)
}

func (i VideoEncoderH264Args) ToVideoEncoderH264PtrOutput() VideoEncoderH264PtrOutput {
	return i.ToVideoEncoderH264PtrOutputWithContext(context.Background())
}

func (i VideoEncoderH264Args) ToVideoEncoderH264PtrOutputWithContext(ctx context.Context) VideoEncoderH264PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoEncoderH264Output).ToVideoEncoderH264PtrOutputWithContext(ctx)
}









type VideoEncoderH264PtrInput interface {
	pulumi.Input

	ToVideoEncoderH264PtrOutput() VideoEncoderH264PtrOutput
	ToVideoEncoderH264PtrOutputWithContext(context.Context) VideoEncoderH264PtrOutput
}

type videoEncoderH264PtrType VideoEncoderH264Args

func VideoEncoderH264Ptr(v *VideoEncoderH264Args) VideoEncoderH264PtrInput {
	return (*videoEncoderH264PtrType)(v)
}

func (*videoEncoderH264PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoEncoderH264)(nil)).Elem()
}

func (i *videoEncoderH264PtrType) ToVideoEncoderH264PtrOutput() VideoEncoderH264PtrOutput {
	return i.ToVideoEncoderH264PtrOutputWithContext(context.Background())
}

func (i *videoEncoderH264PtrType) ToVideoEncoderH264PtrOutputWithContext(ctx context.Context) VideoEncoderH264PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoEncoderH264PtrOutput)
}

type VideoEncoderH264Output struct{ *pulumi.OutputState }

func (VideoEncoderH264Output) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoEncoderH264)(nil)).Elem()
}

func (o VideoEncoderH264Output) ToVideoEncoderH264Output() VideoEncoderH264Output {
	return o
}

func (o VideoEncoderH264Output) ToVideoEncoderH264OutputWithContext(ctx context.Context) VideoEncoderH264Output {
	return o
}

func (o VideoEncoderH264Output) ToVideoEncoderH264PtrOutput() VideoEncoderH264PtrOutput {
	return o.ToVideoEncoderH264PtrOutputWithContext(context.Background())
}

func (o VideoEncoderH264Output) ToVideoEncoderH264PtrOutputWithContext(ctx context.Context) VideoEncoderH264PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoEncoderH264) *VideoEncoderH264 {
		return &v
	}).(VideoEncoderH264PtrOutput)
}

func (o VideoEncoderH264Output) BitrateKbps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoEncoderH264) *string { return v.BitrateKbps }).(pulumi.StringPtrOutput)
}

func (o VideoEncoderH264Output) FrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoEncoderH264) *string { return v.FrameRate }).(pulumi.StringPtrOutput)
}

func (o VideoEncoderH264Output) Scale() VideoScalePtrOutput {
	return o.ApplyT(func(v VideoEncoderH264) *VideoScale { return v.Scale }).(VideoScalePtrOutput)
}

func (o VideoEncoderH264Output) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoEncoderH264) string { return v.Type }).(pulumi.StringOutput)
}

type VideoEncoderH264PtrOutput struct{ *pulumi.OutputState }

func (VideoEncoderH264PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoEncoderH264)(nil)).Elem()
}

func (o VideoEncoderH264PtrOutput) ToVideoEncoderH264PtrOutput() VideoEncoderH264PtrOutput {
	return o
}

func (o VideoEncoderH264PtrOutput) ToVideoEncoderH264PtrOutputWithContext(ctx context.Context) VideoEncoderH264PtrOutput {
	return o
}

func (o VideoEncoderH264PtrOutput) Elem() VideoEncoderH264Output {
	return o.ApplyT(func(v *VideoEncoderH264) VideoEncoderH264 {
		if v != nil {
			return *v
		}
		var ret VideoEncoderH264
		return ret
	}).(VideoEncoderH264Output)
}

func (o VideoEncoderH264PtrOutput) BitrateKbps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoEncoderH264) *string {
		if v == nil {
			return nil
		}
		return v.BitrateKbps
	}).(pulumi.StringPtrOutput)
}

func (o VideoEncoderH264PtrOutput) FrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoEncoderH264) *string {
		if v == nil {
			return nil
		}
		return v.FrameRate
	}).(pulumi.StringPtrOutput)
}

func (o VideoEncoderH264PtrOutput) Scale() VideoScalePtrOutput {
	return o.ApplyT(func(v *VideoEncoderH264) *VideoScale {
		if v == nil {
			return nil
		}
		return v.Scale
	}).(VideoScalePtrOutput)
}

func (o VideoEncoderH264PtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoEncoderH264) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type VideoEncoderH264Response struct {
	BitrateKbps *string             `pulumi:"bitrateKbps"`
	FrameRate   *string             `pulumi:"frameRate"`
	Scale       *VideoScaleResponse `pulumi:"scale"`
	Type        string              `pulumi:"type"`
}





type VideoEncoderH264ResponseInput interface {
	pulumi.Input

	ToVideoEncoderH264ResponseOutput() VideoEncoderH264ResponseOutput
	ToVideoEncoderH264ResponseOutputWithContext(context.Context) VideoEncoderH264ResponseOutput
}

type VideoEncoderH264ResponseArgs struct {
	BitrateKbps pulumi.StringPtrInput      `pulumi:"bitrateKbps"`
	FrameRate   pulumi.StringPtrInput      `pulumi:"frameRate"`
	Scale       VideoScaleResponsePtrInput `pulumi:"scale"`
	Type        pulumi.StringInput         `pulumi:"type"`
}

func (VideoEncoderH264ResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoEncoderH264Response)(nil)).Elem()
}

func (i VideoEncoderH264ResponseArgs) ToVideoEncoderH264ResponseOutput() VideoEncoderH264ResponseOutput {
	return i.ToVideoEncoderH264ResponseOutputWithContext(context.Background())
}

func (i VideoEncoderH264ResponseArgs) ToVideoEncoderH264ResponseOutputWithContext(ctx context.Context) VideoEncoderH264ResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoEncoderH264ResponseOutput)
}

func (i VideoEncoderH264ResponseArgs) ToVideoEncoderH264ResponsePtrOutput() VideoEncoderH264ResponsePtrOutput {
	return i.ToVideoEncoderH264ResponsePtrOutputWithContext(context.Background())
}

func (i VideoEncoderH264ResponseArgs) ToVideoEncoderH264ResponsePtrOutputWithContext(ctx context.Context) VideoEncoderH264ResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoEncoderH264ResponseOutput).ToVideoEncoderH264ResponsePtrOutputWithContext(ctx)
}









type VideoEncoderH264ResponsePtrInput interface {
	pulumi.Input

	ToVideoEncoderH264ResponsePtrOutput() VideoEncoderH264ResponsePtrOutput
	ToVideoEncoderH264ResponsePtrOutputWithContext(context.Context) VideoEncoderH264ResponsePtrOutput
}

type videoEncoderH264ResponsePtrType VideoEncoderH264ResponseArgs

func VideoEncoderH264ResponsePtr(v *VideoEncoderH264ResponseArgs) VideoEncoderH264ResponsePtrInput {
	return (*videoEncoderH264ResponsePtrType)(v)
}

func (*videoEncoderH264ResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoEncoderH264Response)(nil)).Elem()
}

func (i *videoEncoderH264ResponsePtrType) ToVideoEncoderH264ResponsePtrOutput() VideoEncoderH264ResponsePtrOutput {
	return i.ToVideoEncoderH264ResponsePtrOutputWithContext(context.Background())
}

func (i *videoEncoderH264ResponsePtrType) ToVideoEncoderH264ResponsePtrOutputWithContext(ctx context.Context) VideoEncoderH264ResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoEncoderH264ResponsePtrOutput)
}

type VideoEncoderH264ResponseOutput struct{ *pulumi.OutputState }

func (VideoEncoderH264ResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoEncoderH264Response)(nil)).Elem()
}

func (o VideoEncoderH264ResponseOutput) ToVideoEncoderH264ResponseOutput() VideoEncoderH264ResponseOutput {
	return o
}

func (o VideoEncoderH264ResponseOutput) ToVideoEncoderH264ResponseOutputWithContext(ctx context.Context) VideoEncoderH264ResponseOutput {
	return o
}

func (o VideoEncoderH264ResponseOutput) ToVideoEncoderH264ResponsePtrOutput() VideoEncoderH264ResponsePtrOutput {
	return o.ToVideoEncoderH264ResponsePtrOutputWithContext(context.Background())
}

func (o VideoEncoderH264ResponseOutput) ToVideoEncoderH264ResponsePtrOutputWithContext(ctx context.Context) VideoEncoderH264ResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoEncoderH264Response) *VideoEncoderH264Response {
		return &v
	}).(VideoEncoderH264ResponsePtrOutput)
}

func (o VideoEncoderH264ResponseOutput) BitrateKbps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoEncoderH264Response) *string { return v.BitrateKbps }).(pulumi.StringPtrOutput)
}

func (o VideoEncoderH264ResponseOutput) FrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoEncoderH264Response) *string { return v.FrameRate }).(pulumi.StringPtrOutput)
}

func (o VideoEncoderH264ResponseOutput) Scale() VideoScaleResponsePtrOutput {
	return o.ApplyT(func(v VideoEncoderH264Response) *VideoScaleResponse { return v.Scale }).(VideoScaleResponsePtrOutput)
}

func (o VideoEncoderH264ResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoEncoderH264Response) string { return v.Type }).(pulumi.StringOutput)
}

type VideoEncoderH264ResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoEncoderH264ResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoEncoderH264Response)(nil)).Elem()
}

func (o VideoEncoderH264ResponsePtrOutput) ToVideoEncoderH264ResponsePtrOutput() VideoEncoderH264ResponsePtrOutput {
	return o
}

func (o VideoEncoderH264ResponsePtrOutput) ToVideoEncoderH264ResponsePtrOutputWithContext(ctx context.Context) VideoEncoderH264ResponsePtrOutput {
	return o
}

func (o VideoEncoderH264ResponsePtrOutput) Elem() VideoEncoderH264ResponseOutput {
	return o.ApplyT(func(v *VideoEncoderH264Response) VideoEncoderH264Response {
		if v != nil {
			return *v
		}
		var ret VideoEncoderH264Response
		return ret
	}).(VideoEncoderH264ResponseOutput)
}

func (o VideoEncoderH264ResponsePtrOutput) BitrateKbps() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoEncoderH264Response) *string {
		if v == nil {
			return nil
		}
		return v.BitrateKbps
	}).(pulumi.StringPtrOutput)
}

func (o VideoEncoderH264ResponsePtrOutput) FrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoEncoderH264Response) *string {
		if v == nil {
			return nil
		}
		return v.FrameRate
	}).(pulumi.StringPtrOutput)
}

func (o VideoEncoderH264ResponsePtrOutput) Scale() VideoScaleResponsePtrOutput {
	return o.ApplyT(func(v *VideoEncoderH264Response) *VideoScaleResponse {
		if v == nil {
			return nil
		}
		return v.Scale
	}).(VideoScaleResponsePtrOutput)
}

func (o VideoEncoderH264ResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoEncoderH264Response) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type VideoFlagsResponse struct {
	CanStream   bool `pulumi:"canStream"`
	HasData     bool `pulumi:"hasData"`
	IsRecording bool `pulumi:"isRecording"`
}





type VideoFlagsResponseInput interface {
	pulumi.Input

	ToVideoFlagsResponseOutput() VideoFlagsResponseOutput
	ToVideoFlagsResponseOutputWithContext(context.Context) VideoFlagsResponseOutput
}

type VideoFlagsResponseArgs struct {
	CanStream   pulumi.BoolInput `pulumi:"canStream"`
	HasData     pulumi.BoolInput `pulumi:"hasData"`
	IsRecording pulumi.BoolInput `pulumi:"isRecording"`
}

func (VideoFlagsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoFlagsResponse)(nil)).Elem()
}

func (i VideoFlagsResponseArgs) ToVideoFlagsResponseOutput() VideoFlagsResponseOutput {
	return i.ToVideoFlagsResponseOutputWithContext(context.Background())
}

func (i VideoFlagsResponseArgs) ToVideoFlagsResponseOutputWithContext(ctx context.Context) VideoFlagsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoFlagsResponseOutput)
}

func (i VideoFlagsResponseArgs) ToVideoFlagsResponsePtrOutput() VideoFlagsResponsePtrOutput {
	return i.ToVideoFlagsResponsePtrOutputWithContext(context.Background())
}

func (i VideoFlagsResponseArgs) ToVideoFlagsResponsePtrOutputWithContext(ctx context.Context) VideoFlagsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoFlagsResponseOutput).ToVideoFlagsResponsePtrOutputWithContext(ctx)
}









type VideoFlagsResponsePtrInput interface {
	pulumi.Input

	ToVideoFlagsResponsePtrOutput() VideoFlagsResponsePtrOutput
	ToVideoFlagsResponsePtrOutputWithContext(context.Context) VideoFlagsResponsePtrOutput
}

type videoFlagsResponsePtrType VideoFlagsResponseArgs

func VideoFlagsResponsePtr(v *VideoFlagsResponseArgs) VideoFlagsResponsePtrInput {
	return (*videoFlagsResponsePtrType)(v)
}

func (*videoFlagsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoFlagsResponse)(nil)).Elem()
}

func (i *videoFlagsResponsePtrType) ToVideoFlagsResponsePtrOutput() VideoFlagsResponsePtrOutput {
	return i.ToVideoFlagsResponsePtrOutputWithContext(context.Background())
}

func (i *videoFlagsResponsePtrType) ToVideoFlagsResponsePtrOutputWithContext(ctx context.Context) VideoFlagsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoFlagsResponsePtrOutput)
}

type VideoFlagsResponseOutput struct{ *pulumi.OutputState }

func (VideoFlagsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoFlagsResponse)(nil)).Elem()
}

func (o VideoFlagsResponseOutput) ToVideoFlagsResponseOutput() VideoFlagsResponseOutput {
	return o
}

func (o VideoFlagsResponseOutput) ToVideoFlagsResponseOutputWithContext(ctx context.Context) VideoFlagsResponseOutput {
	return o
}

func (o VideoFlagsResponseOutput) ToVideoFlagsResponsePtrOutput() VideoFlagsResponsePtrOutput {
	return o.ToVideoFlagsResponsePtrOutputWithContext(context.Background())
}

func (o VideoFlagsResponseOutput) ToVideoFlagsResponsePtrOutputWithContext(ctx context.Context) VideoFlagsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoFlagsResponse) *VideoFlagsResponse {
		return &v
	}).(VideoFlagsResponsePtrOutput)
}

func (o VideoFlagsResponseOutput) CanStream() pulumi.BoolOutput {
	return o.ApplyT(func(v VideoFlagsResponse) bool { return v.CanStream }).(pulumi.BoolOutput)
}

func (o VideoFlagsResponseOutput) HasData() pulumi.BoolOutput {
	return o.ApplyT(func(v VideoFlagsResponse) bool { return v.HasData }).(pulumi.BoolOutput)
}

func (o VideoFlagsResponseOutput) IsRecording() pulumi.BoolOutput {
	return o.ApplyT(func(v VideoFlagsResponse) bool { return v.IsRecording }).(pulumi.BoolOutput)
}

type VideoFlagsResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoFlagsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoFlagsResponse)(nil)).Elem()
}

func (o VideoFlagsResponsePtrOutput) ToVideoFlagsResponsePtrOutput() VideoFlagsResponsePtrOutput {
	return o
}

func (o VideoFlagsResponsePtrOutput) ToVideoFlagsResponsePtrOutputWithContext(ctx context.Context) VideoFlagsResponsePtrOutput {
	return o
}

func (o VideoFlagsResponsePtrOutput) Elem() VideoFlagsResponseOutput {
	return o.ApplyT(func(v *VideoFlagsResponse) VideoFlagsResponse {
		if v != nil {
			return *v
		}
		var ret VideoFlagsResponse
		return ret
	}).(VideoFlagsResponseOutput)
}

func (o VideoFlagsResponsePtrOutput) CanStream() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VideoFlagsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CanStream
	}).(pulumi.BoolPtrOutput)
}

func (o VideoFlagsResponsePtrOutput) HasData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VideoFlagsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.HasData
	}).(pulumi.BoolPtrOutput)
}

func (o VideoFlagsResponsePtrOutput) IsRecording() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VideoFlagsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsRecording
	}).(pulumi.BoolPtrOutput)
}

type VideoMediaInfoResponse struct {
	SegmentLength string `pulumi:"segmentLength"`
}





type VideoMediaInfoResponseInput interface {
	pulumi.Input

	ToVideoMediaInfoResponseOutput() VideoMediaInfoResponseOutput
	ToVideoMediaInfoResponseOutputWithContext(context.Context) VideoMediaInfoResponseOutput
}

type VideoMediaInfoResponseArgs struct {
	SegmentLength pulumi.StringInput `pulumi:"segmentLength"`
}

func (VideoMediaInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoMediaInfoResponse)(nil)).Elem()
}

func (i VideoMediaInfoResponseArgs) ToVideoMediaInfoResponseOutput() VideoMediaInfoResponseOutput {
	return i.ToVideoMediaInfoResponseOutputWithContext(context.Background())
}

func (i VideoMediaInfoResponseArgs) ToVideoMediaInfoResponseOutputWithContext(ctx context.Context) VideoMediaInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoMediaInfoResponseOutput)
}

func (i VideoMediaInfoResponseArgs) ToVideoMediaInfoResponsePtrOutput() VideoMediaInfoResponsePtrOutput {
	return i.ToVideoMediaInfoResponsePtrOutputWithContext(context.Background())
}

func (i VideoMediaInfoResponseArgs) ToVideoMediaInfoResponsePtrOutputWithContext(ctx context.Context) VideoMediaInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoMediaInfoResponseOutput).ToVideoMediaInfoResponsePtrOutputWithContext(ctx)
}









type VideoMediaInfoResponsePtrInput interface {
	pulumi.Input

	ToVideoMediaInfoResponsePtrOutput() VideoMediaInfoResponsePtrOutput
	ToVideoMediaInfoResponsePtrOutputWithContext(context.Context) VideoMediaInfoResponsePtrOutput
}

type videoMediaInfoResponsePtrType VideoMediaInfoResponseArgs

func VideoMediaInfoResponsePtr(v *VideoMediaInfoResponseArgs) VideoMediaInfoResponsePtrInput {
	return (*videoMediaInfoResponsePtrType)(v)
}

func (*videoMediaInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoMediaInfoResponse)(nil)).Elem()
}

func (i *videoMediaInfoResponsePtrType) ToVideoMediaInfoResponsePtrOutput() VideoMediaInfoResponsePtrOutput {
	return i.ToVideoMediaInfoResponsePtrOutputWithContext(context.Background())
}

func (i *videoMediaInfoResponsePtrType) ToVideoMediaInfoResponsePtrOutputWithContext(ctx context.Context) VideoMediaInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoMediaInfoResponsePtrOutput)
}

type VideoMediaInfoResponseOutput struct{ *pulumi.OutputState }

func (VideoMediaInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoMediaInfoResponse)(nil)).Elem()
}

func (o VideoMediaInfoResponseOutput) ToVideoMediaInfoResponseOutput() VideoMediaInfoResponseOutput {
	return o
}

func (o VideoMediaInfoResponseOutput) ToVideoMediaInfoResponseOutputWithContext(ctx context.Context) VideoMediaInfoResponseOutput {
	return o
}

func (o VideoMediaInfoResponseOutput) ToVideoMediaInfoResponsePtrOutput() VideoMediaInfoResponsePtrOutput {
	return o.ToVideoMediaInfoResponsePtrOutputWithContext(context.Background())
}

func (o VideoMediaInfoResponseOutput) ToVideoMediaInfoResponsePtrOutputWithContext(ctx context.Context) VideoMediaInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoMediaInfoResponse) *VideoMediaInfoResponse {
		return &v
	}).(VideoMediaInfoResponsePtrOutput)
}

func (o VideoMediaInfoResponseOutput) SegmentLength() pulumi.StringOutput {
	return o.ApplyT(func(v VideoMediaInfoResponse) string { return v.SegmentLength }).(pulumi.StringOutput)
}

type VideoMediaInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoMediaInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoMediaInfoResponse)(nil)).Elem()
}

func (o VideoMediaInfoResponsePtrOutput) ToVideoMediaInfoResponsePtrOutput() VideoMediaInfoResponsePtrOutput {
	return o
}

func (o VideoMediaInfoResponsePtrOutput) ToVideoMediaInfoResponsePtrOutputWithContext(ctx context.Context) VideoMediaInfoResponsePtrOutput {
	return o
}

func (o VideoMediaInfoResponsePtrOutput) Elem() VideoMediaInfoResponseOutput {
	return o.ApplyT(func(v *VideoMediaInfoResponse) VideoMediaInfoResponse {
		if v != nil {
			return *v
		}
		var ret VideoMediaInfoResponse
		return ret
	}).(VideoMediaInfoResponseOutput)
}

func (o VideoMediaInfoResponsePtrOutput) SegmentLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoMediaInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SegmentLength
	}).(pulumi.StringPtrOutput)
}

type VideoPublishingOptions struct {
	DisableArchive        *string `pulumi:"disableArchive"`
	DisableRtspPublishing *string `pulumi:"disableRtspPublishing"`
}





type VideoPublishingOptionsInput interface {
	pulumi.Input

	ToVideoPublishingOptionsOutput() VideoPublishingOptionsOutput
	ToVideoPublishingOptionsOutputWithContext(context.Context) VideoPublishingOptionsOutput
}

type VideoPublishingOptionsArgs struct {
	DisableArchive        pulumi.StringPtrInput `pulumi:"disableArchive"`
	DisableRtspPublishing pulumi.StringPtrInput `pulumi:"disableRtspPublishing"`
}

func (VideoPublishingOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoPublishingOptions)(nil)).Elem()
}

func (i VideoPublishingOptionsArgs) ToVideoPublishingOptionsOutput() VideoPublishingOptionsOutput {
	return i.ToVideoPublishingOptionsOutputWithContext(context.Background())
}

func (i VideoPublishingOptionsArgs) ToVideoPublishingOptionsOutputWithContext(ctx context.Context) VideoPublishingOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoPublishingOptionsOutput)
}

func (i VideoPublishingOptionsArgs) ToVideoPublishingOptionsPtrOutput() VideoPublishingOptionsPtrOutput {
	return i.ToVideoPublishingOptionsPtrOutputWithContext(context.Background())
}

func (i VideoPublishingOptionsArgs) ToVideoPublishingOptionsPtrOutputWithContext(ctx context.Context) VideoPublishingOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoPublishingOptionsOutput).ToVideoPublishingOptionsPtrOutputWithContext(ctx)
}









type VideoPublishingOptionsPtrInput interface {
	pulumi.Input

	ToVideoPublishingOptionsPtrOutput() VideoPublishingOptionsPtrOutput
	ToVideoPublishingOptionsPtrOutputWithContext(context.Context) VideoPublishingOptionsPtrOutput
}

type videoPublishingOptionsPtrType VideoPublishingOptionsArgs

func VideoPublishingOptionsPtr(v *VideoPublishingOptionsArgs) VideoPublishingOptionsPtrInput {
	return (*videoPublishingOptionsPtrType)(v)
}

func (*videoPublishingOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoPublishingOptions)(nil)).Elem()
}

func (i *videoPublishingOptionsPtrType) ToVideoPublishingOptionsPtrOutput() VideoPublishingOptionsPtrOutput {
	return i.ToVideoPublishingOptionsPtrOutputWithContext(context.Background())
}

func (i *videoPublishingOptionsPtrType) ToVideoPublishingOptionsPtrOutputWithContext(ctx context.Context) VideoPublishingOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoPublishingOptionsPtrOutput)
}

type VideoPublishingOptionsOutput struct{ *pulumi.OutputState }

func (VideoPublishingOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoPublishingOptions)(nil)).Elem()
}

func (o VideoPublishingOptionsOutput) ToVideoPublishingOptionsOutput() VideoPublishingOptionsOutput {
	return o
}

func (o VideoPublishingOptionsOutput) ToVideoPublishingOptionsOutputWithContext(ctx context.Context) VideoPublishingOptionsOutput {
	return o
}

func (o VideoPublishingOptionsOutput) ToVideoPublishingOptionsPtrOutput() VideoPublishingOptionsPtrOutput {
	return o.ToVideoPublishingOptionsPtrOutputWithContext(context.Background())
}

func (o VideoPublishingOptionsOutput) ToVideoPublishingOptionsPtrOutputWithContext(ctx context.Context) VideoPublishingOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoPublishingOptions) *VideoPublishingOptions {
		return &v
	}).(VideoPublishingOptionsPtrOutput)
}

func (o VideoPublishingOptionsOutput) DisableArchive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoPublishingOptions) *string { return v.DisableArchive }).(pulumi.StringPtrOutput)
}

func (o VideoPublishingOptionsOutput) DisableRtspPublishing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoPublishingOptions) *string { return v.DisableRtspPublishing }).(pulumi.StringPtrOutput)
}

type VideoPublishingOptionsPtrOutput struct{ *pulumi.OutputState }

func (VideoPublishingOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoPublishingOptions)(nil)).Elem()
}

func (o VideoPublishingOptionsPtrOutput) ToVideoPublishingOptionsPtrOutput() VideoPublishingOptionsPtrOutput {
	return o
}

func (o VideoPublishingOptionsPtrOutput) ToVideoPublishingOptionsPtrOutputWithContext(ctx context.Context) VideoPublishingOptionsPtrOutput {
	return o
}

func (o VideoPublishingOptionsPtrOutput) Elem() VideoPublishingOptionsOutput {
	return o.ApplyT(func(v *VideoPublishingOptions) VideoPublishingOptions {
		if v != nil {
			return *v
		}
		var ret VideoPublishingOptions
		return ret
	}).(VideoPublishingOptionsOutput)
}

func (o VideoPublishingOptionsPtrOutput) DisableArchive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoPublishingOptions) *string {
		if v == nil {
			return nil
		}
		return v.DisableArchive
	}).(pulumi.StringPtrOutput)
}

func (o VideoPublishingOptionsPtrOutput) DisableRtspPublishing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoPublishingOptions) *string {
		if v == nil {
			return nil
		}
		return v.DisableRtspPublishing
	}).(pulumi.StringPtrOutput)
}

type VideoPublishingOptionsResponse struct {
	DisableArchive        *string `pulumi:"disableArchive"`
	DisableRtspPublishing *string `pulumi:"disableRtspPublishing"`
}





type VideoPublishingOptionsResponseInput interface {
	pulumi.Input

	ToVideoPublishingOptionsResponseOutput() VideoPublishingOptionsResponseOutput
	ToVideoPublishingOptionsResponseOutputWithContext(context.Context) VideoPublishingOptionsResponseOutput
}

type VideoPublishingOptionsResponseArgs struct {
	DisableArchive        pulumi.StringPtrInput `pulumi:"disableArchive"`
	DisableRtspPublishing pulumi.StringPtrInput `pulumi:"disableRtspPublishing"`
}

func (VideoPublishingOptionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoPublishingOptionsResponse)(nil)).Elem()
}

func (i VideoPublishingOptionsResponseArgs) ToVideoPublishingOptionsResponseOutput() VideoPublishingOptionsResponseOutput {
	return i.ToVideoPublishingOptionsResponseOutputWithContext(context.Background())
}

func (i VideoPublishingOptionsResponseArgs) ToVideoPublishingOptionsResponseOutputWithContext(ctx context.Context) VideoPublishingOptionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoPublishingOptionsResponseOutput)
}

func (i VideoPublishingOptionsResponseArgs) ToVideoPublishingOptionsResponsePtrOutput() VideoPublishingOptionsResponsePtrOutput {
	return i.ToVideoPublishingOptionsResponsePtrOutputWithContext(context.Background())
}

func (i VideoPublishingOptionsResponseArgs) ToVideoPublishingOptionsResponsePtrOutputWithContext(ctx context.Context) VideoPublishingOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoPublishingOptionsResponseOutput).ToVideoPublishingOptionsResponsePtrOutputWithContext(ctx)
}









type VideoPublishingOptionsResponsePtrInput interface {
	pulumi.Input

	ToVideoPublishingOptionsResponsePtrOutput() VideoPublishingOptionsResponsePtrOutput
	ToVideoPublishingOptionsResponsePtrOutputWithContext(context.Context) VideoPublishingOptionsResponsePtrOutput
}

type videoPublishingOptionsResponsePtrType VideoPublishingOptionsResponseArgs

func VideoPublishingOptionsResponsePtr(v *VideoPublishingOptionsResponseArgs) VideoPublishingOptionsResponsePtrInput {
	return (*videoPublishingOptionsResponsePtrType)(v)
}

func (*videoPublishingOptionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoPublishingOptionsResponse)(nil)).Elem()
}

func (i *videoPublishingOptionsResponsePtrType) ToVideoPublishingOptionsResponsePtrOutput() VideoPublishingOptionsResponsePtrOutput {
	return i.ToVideoPublishingOptionsResponsePtrOutputWithContext(context.Background())
}

func (i *videoPublishingOptionsResponsePtrType) ToVideoPublishingOptionsResponsePtrOutputWithContext(ctx context.Context) VideoPublishingOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoPublishingOptionsResponsePtrOutput)
}

type VideoPublishingOptionsResponseOutput struct{ *pulumi.OutputState }

func (VideoPublishingOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoPublishingOptionsResponse)(nil)).Elem()
}

func (o VideoPublishingOptionsResponseOutput) ToVideoPublishingOptionsResponseOutput() VideoPublishingOptionsResponseOutput {
	return o
}

func (o VideoPublishingOptionsResponseOutput) ToVideoPublishingOptionsResponseOutputWithContext(ctx context.Context) VideoPublishingOptionsResponseOutput {
	return o
}

func (o VideoPublishingOptionsResponseOutput) ToVideoPublishingOptionsResponsePtrOutput() VideoPublishingOptionsResponsePtrOutput {
	return o.ToVideoPublishingOptionsResponsePtrOutputWithContext(context.Background())
}

func (o VideoPublishingOptionsResponseOutput) ToVideoPublishingOptionsResponsePtrOutputWithContext(ctx context.Context) VideoPublishingOptionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoPublishingOptionsResponse) *VideoPublishingOptionsResponse {
		return &v
	}).(VideoPublishingOptionsResponsePtrOutput)
}

func (o VideoPublishingOptionsResponseOutput) DisableArchive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoPublishingOptionsResponse) *string { return v.DisableArchive }).(pulumi.StringPtrOutput)
}

func (o VideoPublishingOptionsResponseOutput) DisableRtspPublishing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoPublishingOptionsResponse) *string { return v.DisableRtspPublishing }).(pulumi.StringPtrOutput)
}

type VideoPublishingOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoPublishingOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoPublishingOptionsResponse)(nil)).Elem()
}

func (o VideoPublishingOptionsResponsePtrOutput) ToVideoPublishingOptionsResponsePtrOutput() VideoPublishingOptionsResponsePtrOutput {
	return o
}

func (o VideoPublishingOptionsResponsePtrOutput) ToVideoPublishingOptionsResponsePtrOutputWithContext(ctx context.Context) VideoPublishingOptionsResponsePtrOutput {
	return o
}

func (o VideoPublishingOptionsResponsePtrOutput) Elem() VideoPublishingOptionsResponseOutput {
	return o.ApplyT(func(v *VideoPublishingOptionsResponse) VideoPublishingOptionsResponse {
		if v != nil {
			return *v
		}
		var ret VideoPublishingOptionsResponse
		return ret
	}).(VideoPublishingOptionsResponseOutput)
}

func (o VideoPublishingOptionsResponsePtrOutput) DisableArchive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoPublishingOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisableArchive
	}).(pulumi.StringPtrOutput)
}

func (o VideoPublishingOptionsResponsePtrOutput) DisableRtspPublishing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoPublishingOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisableRtspPublishing
	}).(pulumi.StringPtrOutput)
}

type VideoScale struct {
	Height *string `pulumi:"height"`
	Mode   *string `pulumi:"mode"`
	Width  *string `pulumi:"width"`
}





type VideoScaleInput interface {
	pulumi.Input

	ToVideoScaleOutput() VideoScaleOutput
	ToVideoScaleOutputWithContext(context.Context) VideoScaleOutput
}

type VideoScaleArgs struct {
	Height pulumi.StringPtrInput `pulumi:"height"`
	Mode   pulumi.StringPtrInput `pulumi:"mode"`
	Width  pulumi.StringPtrInput `pulumi:"width"`
}

func (VideoScaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoScale)(nil)).Elem()
}

func (i VideoScaleArgs) ToVideoScaleOutput() VideoScaleOutput {
	return i.ToVideoScaleOutputWithContext(context.Background())
}

func (i VideoScaleArgs) ToVideoScaleOutputWithContext(ctx context.Context) VideoScaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoScaleOutput)
}

func (i VideoScaleArgs) ToVideoScalePtrOutput() VideoScalePtrOutput {
	return i.ToVideoScalePtrOutputWithContext(context.Background())
}

func (i VideoScaleArgs) ToVideoScalePtrOutputWithContext(ctx context.Context) VideoScalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoScaleOutput).ToVideoScalePtrOutputWithContext(ctx)
}









type VideoScalePtrInput interface {
	pulumi.Input

	ToVideoScalePtrOutput() VideoScalePtrOutput
	ToVideoScalePtrOutputWithContext(context.Context) VideoScalePtrOutput
}

type videoScalePtrType VideoScaleArgs

func VideoScalePtr(v *VideoScaleArgs) VideoScalePtrInput {
	return (*videoScalePtrType)(v)
}

func (*videoScalePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoScale)(nil)).Elem()
}

func (i *videoScalePtrType) ToVideoScalePtrOutput() VideoScalePtrOutput {
	return i.ToVideoScalePtrOutputWithContext(context.Background())
}

func (i *videoScalePtrType) ToVideoScalePtrOutputWithContext(ctx context.Context) VideoScalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoScalePtrOutput)
}

type VideoScaleOutput struct{ *pulumi.OutputState }

func (VideoScaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoScale)(nil)).Elem()
}

func (o VideoScaleOutput) ToVideoScaleOutput() VideoScaleOutput {
	return o
}

func (o VideoScaleOutput) ToVideoScaleOutputWithContext(ctx context.Context) VideoScaleOutput {
	return o
}

func (o VideoScaleOutput) ToVideoScalePtrOutput() VideoScalePtrOutput {
	return o.ToVideoScalePtrOutputWithContext(context.Background())
}

func (o VideoScaleOutput) ToVideoScalePtrOutputWithContext(ctx context.Context) VideoScalePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoScale) *VideoScale {
		return &v
	}).(VideoScalePtrOutput)
}

func (o VideoScaleOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoScale) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o VideoScaleOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoScale) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o VideoScaleOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoScale) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type VideoScalePtrOutput struct{ *pulumi.OutputState }

func (VideoScalePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoScale)(nil)).Elem()
}

func (o VideoScalePtrOutput) ToVideoScalePtrOutput() VideoScalePtrOutput {
	return o
}

func (o VideoScalePtrOutput) ToVideoScalePtrOutputWithContext(ctx context.Context) VideoScalePtrOutput {
	return o
}

func (o VideoScalePtrOutput) Elem() VideoScaleOutput {
	return o.ApplyT(func(v *VideoScale) VideoScale {
		if v != nil {
			return *v
		}
		var ret VideoScale
		return ret
	}).(VideoScaleOutput)
}

func (o VideoScalePtrOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoScale) *string {
		if v == nil {
			return nil
		}
		return v.Height
	}).(pulumi.StringPtrOutput)
}

func (o VideoScalePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoScale) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o VideoScalePtrOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoScale) *string {
		if v == nil {
			return nil
		}
		return v.Width
	}).(pulumi.StringPtrOutput)
}

type VideoScaleResponse struct {
	Height *string `pulumi:"height"`
	Mode   *string `pulumi:"mode"`
	Width  *string `pulumi:"width"`
}





type VideoScaleResponseInput interface {
	pulumi.Input

	ToVideoScaleResponseOutput() VideoScaleResponseOutput
	ToVideoScaleResponseOutputWithContext(context.Context) VideoScaleResponseOutput
}

type VideoScaleResponseArgs struct {
	Height pulumi.StringPtrInput `pulumi:"height"`
	Mode   pulumi.StringPtrInput `pulumi:"mode"`
	Width  pulumi.StringPtrInput `pulumi:"width"`
}

func (VideoScaleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoScaleResponse)(nil)).Elem()
}

func (i VideoScaleResponseArgs) ToVideoScaleResponseOutput() VideoScaleResponseOutput {
	return i.ToVideoScaleResponseOutputWithContext(context.Background())
}

func (i VideoScaleResponseArgs) ToVideoScaleResponseOutputWithContext(ctx context.Context) VideoScaleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoScaleResponseOutput)
}

func (i VideoScaleResponseArgs) ToVideoScaleResponsePtrOutput() VideoScaleResponsePtrOutput {
	return i.ToVideoScaleResponsePtrOutputWithContext(context.Background())
}

func (i VideoScaleResponseArgs) ToVideoScaleResponsePtrOutputWithContext(ctx context.Context) VideoScaleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoScaleResponseOutput).ToVideoScaleResponsePtrOutputWithContext(ctx)
}









type VideoScaleResponsePtrInput interface {
	pulumi.Input

	ToVideoScaleResponsePtrOutput() VideoScaleResponsePtrOutput
	ToVideoScaleResponsePtrOutputWithContext(context.Context) VideoScaleResponsePtrOutput
}

type videoScaleResponsePtrType VideoScaleResponseArgs

func VideoScaleResponsePtr(v *VideoScaleResponseArgs) VideoScaleResponsePtrInput {
	return (*videoScaleResponsePtrType)(v)
}

func (*videoScaleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoScaleResponse)(nil)).Elem()
}

func (i *videoScaleResponsePtrType) ToVideoScaleResponsePtrOutput() VideoScaleResponsePtrOutput {
	return i.ToVideoScaleResponsePtrOutputWithContext(context.Background())
}

func (i *videoScaleResponsePtrType) ToVideoScaleResponsePtrOutputWithContext(ctx context.Context) VideoScaleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoScaleResponsePtrOutput)
}

type VideoScaleResponseOutput struct{ *pulumi.OutputState }

func (VideoScaleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoScaleResponse)(nil)).Elem()
}

func (o VideoScaleResponseOutput) ToVideoScaleResponseOutput() VideoScaleResponseOutput {
	return o
}

func (o VideoScaleResponseOutput) ToVideoScaleResponseOutputWithContext(ctx context.Context) VideoScaleResponseOutput {
	return o
}

func (o VideoScaleResponseOutput) ToVideoScaleResponsePtrOutput() VideoScaleResponsePtrOutput {
	return o.ToVideoScaleResponsePtrOutputWithContext(context.Background())
}

func (o VideoScaleResponseOutput) ToVideoScaleResponsePtrOutputWithContext(ctx context.Context) VideoScaleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoScaleResponse) *VideoScaleResponse {
		return &v
	}).(VideoScaleResponsePtrOutput)
}

func (o VideoScaleResponseOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoScaleResponse) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o VideoScaleResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoScaleResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o VideoScaleResponseOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoScaleResponse) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type VideoScaleResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoScaleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoScaleResponse)(nil)).Elem()
}

func (o VideoScaleResponsePtrOutput) ToVideoScaleResponsePtrOutput() VideoScaleResponsePtrOutput {
	return o
}

func (o VideoScaleResponsePtrOutput) ToVideoScaleResponsePtrOutputWithContext(ctx context.Context) VideoScaleResponsePtrOutput {
	return o
}

func (o VideoScaleResponsePtrOutput) Elem() VideoScaleResponseOutput {
	return o.ApplyT(func(v *VideoScaleResponse) VideoScaleResponse {
		if v != nil {
			return *v
		}
		var ret VideoScaleResponse
		return ret
	}).(VideoScaleResponseOutput)
}

func (o VideoScaleResponsePtrOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoScaleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Height
	}).(pulumi.StringPtrOutput)
}

func (o VideoScaleResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoScaleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o VideoScaleResponsePtrOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoScaleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Width
	}).(pulumi.StringPtrOutput)
}

type VideoSequenceAbsoluteTimeMarkers struct {
	Ranges string `pulumi:"ranges"`
	Type   string `pulumi:"type"`
}





type VideoSequenceAbsoluteTimeMarkersInput interface {
	pulumi.Input

	ToVideoSequenceAbsoluteTimeMarkersOutput() VideoSequenceAbsoluteTimeMarkersOutput
	ToVideoSequenceAbsoluteTimeMarkersOutputWithContext(context.Context) VideoSequenceAbsoluteTimeMarkersOutput
}

type VideoSequenceAbsoluteTimeMarkersArgs struct {
	Ranges pulumi.StringInput `pulumi:"ranges"`
	Type   pulumi.StringInput `pulumi:"type"`
}

func (VideoSequenceAbsoluteTimeMarkersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSequenceAbsoluteTimeMarkers)(nil)).Elem()
}

func (i VideoSequenceAbsoluteTimeMarkersArgs) ToVideoSequenceAbsoluteTimeMarkersOutput() VideoSequenceAbsoluteTimeMarkersOutput {
	return i.ToVideoSequenceAbsoluteTimeMarkersOutputWithContext(context.Background())
}

func (i VideoSequenceAbsoluteTimeMarkersArgs) ToVideoSequenceAbsoluteTimeMarkersOutputWithContext(ctx context.Context) VideoSequenceAbsoluteTimeMarkersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoSequenceAbsoluteTimeMarkersOutput)
}

type VideoSequenceAbsoluteTimeMarkersOutput struct{ *pulumi.OutputState }

func (VideoSequenceAbsoluteTimeMarkersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSequenceAbsoluteTimeMarkers)(nil)).Elem()
}

func (o VideoSequenceAbsoluteTimeMarkersOutput) ToVideoSequenceAbsoluteTimeMarkersOutput() VideoSequenceAbsoluteTimeMarkersOutput {
	return o
}

func (o VideoSequenceAbsoluteTimeMarkersOutput) ToVideoSequenceAbsoluteTimeMarkersOutputWithContext(ctx context.Context) VideoSequenceAbsoluteTimeMarkersOutput {
	return o
}

func (o VideoSequenceAbsoluteTimeMarkersOutput) Ranges() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSequenceAbsoluteTimeMarkers) string { return v.Ranges }).(pulumi.StringOutput)
}

func (o VideoSequenceAbsoluteTimeMarkersOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSequenceAbsoluteTimeMarkers) string { return v.Type }).(pulumi.StringOutput)
}

type VideoSequenceAbsoluteTimeMarkersResponse struct {
	Ranges string `pulumi:"ranges"`
	Type   string `pulumi:"type"`
}





type VideoSequenceAbsoluteTimeMarkersResponseInput interface {
	pulumi.Input

	ToVideoSequenceAbsoluteTimeMarkersResponseOutput() VideoSequenceAbsoluteTimeMarkersResponseOutput
	ToVideoSequenceAbsoluteTimeMarkersResponseOutputWithContext(context.Context) VideoSequenceAbsoluteTimeMarkersResponseOutput
}

type VideoSequenceAbsoluteTimeMarkersResponseArgs struct {
	Ranges pulumi.StringInput `pulumi:"ranges"`
	Type   pulumi.StringInput `pulumi:"type"`
}

func (VideoSequenceAbsoluteTimeMarkersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSequenceAbsoluteTimeMarkersResponse)(nil)).Elem()
}

func (i VideoSequenceAbsoluteTimeMarkersResponseArgs) ToVideoSequenceAbsoluteTimeMarkersResponseOutput() VideoSequenceAbsoluteTimeMarkersResponseOutput {
	return i.ToVideoSequenceAbsoluteTimeMarkersResponseOutputWithContext(context.Background())
}

func (i VideoSequenceAbsoluteTimeMarkersResponseArgs) ToVideoSequenceAbsoluteTimeMarkersResponseOutputWithContext(ctx context.Context) VideoSequenceAbsoluteTimeMarkersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoSequenceAbsoluteTimeMarkersResponseOutput)
}

type VideoSequenceAbsoluteTimeMarkersResponseOutput struct{ *pulumi.OutputState }

func (VideoSequenceAbsoluteTimeMarkersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSequenceAbsoluteTimeMarkersResponse)(nil)).Elem()
}

func (o VideoSequenceAbsoluteTimeMarkersResponseOutput) ToVideoSequenceAbsoluteTimeMarkersResponseOutput() VideoSequenceAbsoluteTimeMarkersResponseOutput {
	return o
}

func (o VideoSequenceAbsoluteTimeMarkersResponseOutput) ToVideoSequenceAbsoluteTimeMarkersResponseOutputWithContext(ctx context.Context) VideoSequenceAbsoluteTimeMarkersResponseOutput {
	return o
}

func (o VideoSequenceAbsoluteTimeMarkersResponseOutput) Ranges() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSequenceAbsoluteTimeMarkersResponse) string { return v.Ranges }).(pulumi.StringOutput)
}

func (o VideoSequenceAbsoluteTimeMarkersResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSequenceAbsoluteTimeMarkersResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VideoSink struct {
	Inputs                  []NodeInput              `pulumi:"inputs"`
	Name                    string                   `pulumi:"name"`
	Type                    string                   `pulumi:"type"`
	VideoCreationProperties *VideoCreationProperties `pulumi:"videoCreationProperties"`
	VideoName               string                   `pulumi:"videoName"`
	VideoPublishingOptions  *VideoPublishingOptions  `pulumi:"videoPublishingOptions"`
}





type VideoSinkInput interface {
	pulumi.Input

	ToVideoSinkOutput() VideoSinkOutput
	ToVideoSinkOutputWithContext(context.Context) VideoSinkOutput
}

type VideoSinkArgs struct {
	Inputs                  NodeInputArrayInput             `pulumi:"inputs"`
	Name                    pulumi.StringInput              `pulumi:"name"`
	Type                    pulumi.StringInput              `pulumi:"type"`
	VideoCreationProperties VideoCreationPropertiesPtrInput `pulumi:"videoCreationProperties"`
	VideoName               pulumi.StringInput              `pulumi:"videoName"`
	VideoPublishingOptions  VideoPublishingOptionsPtrInput  `pulumi:"videoPublishingOptions"`
}

func (VideoSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSink)(nil)).Elem()
}

func (i VideoSinkArgs) ToVideoSinkOutput() VideoSinkOutput {
	return i.ToVideoSinkOutputWithContext(context.Background())
}

func (i VideoSinkArgs) ToVideoSinkOutputWithContext(ctx context.Context) VideoSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoSinkOutput)
}





type VideoSinkArrayInput interface {
	pulumi.Input

	ToVideoSinkArrayOutput() VideoSinkArrayOutput
	ToVideoSinkArrayOutputWithContext(context.Context) VideoSinkArrayOutput
}

type VideoSinkArray []VideoSinkInput

func (VideoSinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VideoSink)(nil)).Elem()
}

func (i VideoSinkArray) ToVideoSinkArrayOutput() VideoSinkArrayOutput {
	return i.ToVideoSinkArrayOutputWithContext(context.Background())
}

func (i VideoSinkArray) ToVideoSinkArrayOutputWithContext(ctx context.Context) VideoSinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoSinkArrayOutput)
}

type VideoSinkOutput struct{ *pulumi.OutputState }

func (VideoSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSink)(nil)).Elem()
}

func (o VideoSinkOutput) ToVideoSinkOutput() VideoSinkOutput {
	return o
}

func (o VideoSinkOutput) ToVideoSinkOutputWithContext(ctx context.Context) VideoSinkOutput {
	return o
}

func (o VideoSinkOutput) Inputs() NodeInputArrayOutput {
	return o.ApplyT(func(v VideoSink) []NodeInput { return v.Inputs }).(NodeInputArrayOutput)
}

func (o VideoSinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSink) string { return v.Name }).(pulumi.StringOutput)
}

func (o VideoSinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSink) string { return v.Type }).(pulumi.StringOutput)
}

func (o VideoSinkOutput) VideoCreationProperties() VideoCreationPropertiesPtrOutput {
	return o.ApplyT(func(v VideoSink) *VideoCreationProperties { return v.VideoCreationProperties }).(VideoCreationPropertiesPtrOutput)
}

func (o VideoSinkOutput) VideoName() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSink) string { return v.VideoName }).(pulumi.StringOutput)
}

func (o VideoSinkOutput) VideoPublishingOptions() VideoPublishingOptionsPtrOutput {
	return o.ApplyT(func(v VideoSink) *VideoPublishingOptions { return v.VideoPublishingOptions }).(VideoPublishingOptionsPtrOutput)
}

type VideoSinkArrayOutput struct{ *pulumi.OutputState }

func (VideoSinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VideoSink)(nil)).Elem()
}

func (o VideoSinkArrayOutput) ToVideoSinkArrayOutput() VideoSinkArrayOutput {
	return o
}

func (o VideoSinkArrayOutput) ToVideoSinkArrayOutputWithContext(ctx context.Context) VideoSinkArrayOutput {
	return o
}

func (o VideoSinkArrayOutput) Index(i pulumi.IntInput) VideoSinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VideoSink {
		return vs[0].([]VideoSink)[vs[1].(int)]
	}).(VideoSinkOutput)
}

type VideoSinkResponse struct {
	Inputs                  []NodeInputResponse              `pulumi:"inputs"`
	Name                    string                           `pulumi:"name"`
	Type                    string                           `pulumi:"type"`
	VideoCreationProperties *VideoCreationPropertiesResponse `pulumi:"videoCreationProperties"`
	VideoName               string                           `pulumi:"videoName"`
	VideoPublishingOptions  *VideoPublishingOptionsResponse  `pulumi:"videoPublishingOptions"`
}





type VideoSinkResponseInput interface {
	pulumi.Input

	ToVideoSinkResponseOutput() VideoSinkResponseOutput
	ToVideoSinkResponseOutputWithContext(context.Context) VideoSinkResponseOutput
}

type VideoSinkResponseArgs struct {
	Inputs                  NodeInputResponseArrayInput             `pulumi:"inputs"`
	Name                    pulumi.StringInput                      `pulumi:"name"`
	Type                    pulumi.StringInput                      `pulumi:"type"`
	VideoCreationProperties VideoCreationPropertiesResponsePtrInput `pulumi:"videoCreationProperties"`
	VideoName               pulumi.StringInput                      `pulumi:"videoName"`
	VideoPublishingOptions  VideoPublishingOptionsResponsePtrInput  `pulumi:"videoPublishingOptions"`
}

func (VideoSinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSinkResponse)(nil)).Elem()
}

func (i VideoSinkResponseArgs) ToVideoSinkResponseOutput() VideoSinkResponseOutput {
	return i.ToVideoSinkResponseOutputWithContext(context.Background())
}

func (i VideoSinkResponseArgs) ToVideoSinkResponseOutputWithContext(ctx context.Context) VideoSinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoSinkResponseOutput)
}





type VideoSinkResponseArrayInput interface {
	pulumi.Input

	ToVideoSinkResponseArrayOutput() VideoSinkResponseArrayOutput
	ToVideoSinkResponseArrayOutputWithContext(context.Context) VideoSinkResponseArrayOutput
}

type VideoSinkResponseArray []VideoSinkResponseInput

func (VideoSinkResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VideoSinkResponse)(nil)).Elem()
}

func (i VideoSinkResponseArray) ToVideoSinkResponseArrayOutput() VideoSinkResponseArrayOutput {
	return i.ToVideoSinkResponseArrayOutputWithContext(context.Background())
}

func (i VideoSinkResponseArray) ToVideoSinkResponseArrayOutputWithContext(ctx context.Context) VideoSinkResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoSinkResponseArrayOutput)
}

type VideoSinkResponseOutput struct{ *pulumi.OutputState }

func (VideoSinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSinkResponse)(nil)).Elem()
}

func (o VideoSinkResponseOutput) ToVideoSinkResponseOutput() VideoSinkResponseOutput {
	return o
}

func (o VideoSinkResponseOutput) ToVideoSinkResponseOutputWithContext(ctx context.Context) VideoSinkResponseOutput {
	return o
}

func (o VideoSinkResponseOutput) Inputs() NodeInputResponseArrayOutput {
	return o.ApplyT(func(v VideoSinkResponse) []NodeInputResponse { return v.Inputs }).(NodeInputResponseArrayOutput)
}

func (o VideoSinkResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSinkResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VideoSinkResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSinkResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VideoSinkResponseOutput) VideoCreationProperties() VideoCreationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VideoSinkResponse) *VideoCreationPropertiesResponse { return v.VideoCreationProperties }).(VideoCreationPropertiesResponsePtrOutput)
}

func (o VideoSinkResponseOutput) VideoName() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSinkResponse) string { return v.VideoName }).(pulumi.StringOutput)
}

func (o VideoSinkResponseOutput) VideoPublishingOptions() VideoPublishingOptionsResponsePtrOutput {
	return o.ApplyT(func(v VideoSinkResponse) *VideoPublishingOptionsResponse { return v.VideoPublishingOptions }).(VideoPublishingOptionsResponsePtrOutput)
}

type VideoSinkResponseArrayOutput struct{ *pulumi.OutputState }

func (VideoSinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VideoSinkResponse)(nil)).Elem()
}

func (o VideoSinkResponseArrayOutput) ToVideoSinkResponseArrayOutput() VideoSinkResponseArrayOutput {
	return o
}

func (o VideoSinkResponseArrayOutput) ToVideoSinkResponseArrayOutputWithContext(ctx context.Context) VideoSinkResponseArrayOutput {
	return o
}

func (o VideoSinkResponseArrayOutput) Index(i pulumi.IntInput) VideoSinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VideoSinkResponse {
		return vs[0].([]VideoSinkResponse)[vs[1].(int)]
	}).(VideoSinkResponseOutput)
}

type VideoSource struct {
	Name          string                           `pulumi:"name"`
	TimeSequences VideoSequenceAbsoluteTimeMarkers `pulumi:"timeSequences"`
	Type          string                           `pulumi:"type"`
	VideoName     string                           `pulumi:"videoName"`
}





type VideoSourceInput interface {
	pulumi.Input

	ToVideoSourceOutput() VideoSourceOutput
	ToVideoSourceOutputWithContext(context.Context) VideoSourceOutput
}

type VideoSourceArgs struct {
	Name          pulumi.StringInput                    `pulumi:"name"`
	TimeSequences VideoSequenceAbsoluteTimeMarkersInput `pulumi:"timeSequences"`
	Type          pulumi.StringInput                    `pulumi:"type"`
	VideoName     pulumi.StringInput                    `pulumi:"videoName"`
}

func (VideoSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSource)(nil)).Elem()
}

func (i VideoSourceArgs) ToVideoSourceOutput() VideoSourceOutput {
	return i.ToVideoSourceOutputWithContext(context.Background())
}

func (i VideoSourceArgs) ToVideoSourceOutputWithContext(ctx context.Context) VideoSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoSourceOutput)
}

type VideoSourceOutput struct{ *pulumi.OutputState }

func (VideoSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSource)(nil)).Elem()
}

func (o VideoSourceOutput) ToVideoSourceOutput() VideoSourceOutput {
	return o
}

func (o VideoSourceOutput) ToVideoSourceOutputWithContext(ctx context.Context) VideoSourceOutput {
	return o
}

func (o VideoSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSource) string { return v.Name }).(pulumi.StringOutput)
}

func (o VideoSourceOutput) TimeSequences() VideoSequenceAbsoluteTimeMarkersOutput {
	return o.ApplyT(func(v VideoSource) VideoSequenceAbsoluteTimeMarkers { return v.TimeSequences }).(VideoSequenceAbsoluteTimeMarkersOutput)
}

func (o VideoSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSource) string { return v.Type }).(pulumi.StringOutput)
}

func (o VideoSourceOutput) VideoName() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSource) string { return v.VideoName }).(pulumi.StringOutput)
}

type VideoSourceResponse struct {
	Name          string                                   `pulumi:"name"`
	TimeSequences VideoSequenceAbsoluteTimeMarkersResponse `pulumi:"timeSequences"`
	Type          string                                   `pulumi:"type"`
	VideoName     string                                   `pulumi:"videoName"`
}





type VideoSourceResponseInput interface {
	pulumi.Input

	ToVideoSourceResponseOutput() VideoSourceResponseOutput
	ToVideoSourceResponseOutputWithContext(context.Context) VideoSourceResponseOutput
}

type VideoSourceResponseArgs struct {
	Name          pulumi.StringInput                            `pulumi:"name"`
	TimeSequences VideoSequenceAbsoluteTimeMarkersResponseInput `pulumi:"timeSequences"`
	Type          pulumi.StringInput                            `pulumi:"type"`
	VideoName     pulumi.StringInput                            `pulumi:"videoName"`
}

func (VideoSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSourceResponse)(nil)).Elem()
}

func (i VideoSourceResponseArgs) ToVideoSourceResponseOutput() VideoSourceResponseOutput {
	return i.ToVideoSourceResponseOutputWithContext(context.Background())
}

func (i VideoSourceResponseArgs) ToVideoSourceResponseOutputWithContext(ctx context.Context) VideoSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoSourceResponseOutput)
}

type VideoSourceResponseOutput struct{ *pulumi.OutputState }

func (VideoSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoSourceResponse)(nil)).Elem()
}

func (o VideoSourceResponseOutput) ToVideoSourceResponseOutput() VideoSourceResponseOutput {
	return o
}

func (o VideoSourceResponseOutput) ToVideoSourceResponseOutputWithContext(ctx context.Context) VideoSourceResponseOutput {
	return o
}

func (o VideoSourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VideoSourceResponseOutput) TimeSequences() VideoSequenceAbsoluteTimeMarkersResponseOutput {
	return o.ApplyT(func(v VideoSourceResponse) VideoSequenceAbsoluteTimeMarkersResponse { return v.TimeSequences }).(VideoSequenceAbsoluteTimeMarkersResponseOutput)
}

func (o VideoSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VideoSourceResponseOutput) VideoName() pulumi.StringOutput {
	return o.ApplyT(func(v VideoSourceResponse) string { return v.VideoName }).(pulumi.StringOutput)
}

type VideoStreamingResponse struct {
	ArchiveBaseUrl *string `pulumi:"archiveBaseUrl"`
}





type VideoStreamingResponseInput interface {
	pulumi.Input

	ToVideoStreamingResponseOutput() VideoStreamingResponseOutput
	ToVideoStreamingResponseOutputWithContext(context.Context) VideoStreamingResponseOutput
}

type VideoStreamingResponseArgs struct {
	ArchiveBaseUrl pulumi.StringPtrInput `pulumi:"archiveBaseUrl"`
}

func (VideoStreamingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoStreamingResponse)(nil)).Elem()
}

func (i VideoStreamingResponseArgs) ToVideoStreamingResponseOutput() VideoStreamingResponseOutput {
	return i.ToVideoStreamingResponseOutputWithContext(context.Background())
}

func (i VideoStreamingResponseArgs) ToVideoStreamingResponseOutputWithContext(ctx context.Context) VideoStreamingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoStreamingResponseOutput)
}

func (i VideoStreamingResponseArgs) ToVideoStreamingResponsePtrOutput() VideoStreamingResponsePtrOutput {
	return i.ToVideoStreamingResponsePtrOutputWithContext(context.Background())
}

func (i VideoStreamingResponseArgs) ToVideoStreamingResponsePtrOutputWithContext(ctx context.Context) VideoStreamingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoStreamingResponseOutput).ToVideoStreamingResponsePtrOutputWithContext(ctx)
}









type VideoStreamingResponsePtrInput interface {
	pulumi.Input

	ToVideoStreamingResponsePtrOutput() VideoStreamingResponsePtrOutput
	ToVideoStreamingResponsePtrOutputWithContext(context.Context) VideoStreamingResponsePtrOutput
}

type videoStreamingResponsePtrType VideoStreamingResponseArgs

func VideoStreamingResponsePtr(v *VideoStreamingResponseArgs) VideoStreamingResponsePtrInput {
	return (*videoStreamingResponsePtrType)(v)
}

func (*videoStreamingResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoStreamingResponse)(nil)).Elem()
}

func (i *videoStreamingResponsePtrType) ToVideoStreamingResponsePtrOutput() VideoStreamingResponsePtrOutput {
	return i.ToVideoStreamingResponsePtrOutputWithContext(context.Background())
}

func (i *videoStreamingResponsePtrType) ToVideoStreamingResponsePtrOutputWithContext(ctx context.Context) VideoStreamingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoStreamingResponsePtrOutput)
}

type VideoStreamingResponseOutput struct{ *pulumi.OutputState }

func (VideoStreamingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoStreamingResponse)(nil)).Elem()
}

func (o VideoStreamingResponseOutput) ToVideoStreamingResponseOutput() VideoStreamingResponseOutput {
	return o
}

func (o VideoStreamingResponseOutput) ToVideoStreamingResponseOutputWithContext(ctx context.Context) VideoStreamingResponseOutput {
	return o
}

func (o VideoStreamingResponseOutput) ToVideoStreamingResponsePtrOutput() VideoStreamingResponsePtrOutput {
	return o.ToVideoStreamingResponsePtrOutputWithContext(context.Background())
}

func (o VideoStreamingResponseOutput) ToVideoStreamingResponsePtrOutputWithContext(ctx context.Context) VideoStreamingResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoStreamingResponse) *VideoStreamingResponse {
		return &v
	}).(VideoStreamingResponsePtrOutput)
}

func (o VideoStreamingResponseOutput) ArchiveBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoStreamingResponse) *string { return v.ArchiveBaseUrl }).(pulumi.StringPtrOutput)
}

type VideoStreamingResponsePtrOutput struct{ *pulumi.OutputState }

func (VideoStreamingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoStreamingResponse)(nil)).Elem()
}

func (o VideoStreamingResponsePtrOutput) ToVideoStreamingResponsePtrOutput() VideoStreamingResponsePtrOutput {
	return o
}

func (o VideoStreamingResponsePtrOutput) ToVideoStreamingResponsePtrOutputWithContext(ctx context.Context) VideoStreamingResponsePtrOutput {
	return o
}

func (o VideoStreamingResponsePtrOutput) Elem() VideoStreamingResponseOutput {
	return o.ApplyT(func(v *VideoStreamingResponse) VideoStreamingResponse {
		if v != nil {
			return *v
		}
		var ret VideoStreamingResponse
		return ret
	}).(VideoStreamingResponseOutput)
}

func (o VideoStreamingResponsePtrOutput) ArchiveBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoStreamingResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArchiveBaseUrl
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountEncryptionInput)(nil)).Elem(), AccountEncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountEncryptionPtrInput)(nil)).Elem(), AccountEncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountEncryptionResponseInput)(nil)).Elem(), AccountEncryptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountEncryptionResponsePtrInput)(nil)).Elem(), AccountEncryptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioEncoderAacInput)(nil)).Elem(), AudioEncoderAacArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioEncoderAacPtrInput)(nil)).Elem(), AudioEncoderAacArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioEncoderAacResponseInput)(nil)).Elem(), AudioEncoderAacResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioEncoderAacResponsePtrInput)(nil)).Elem(), AudioEncoderAacResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EccTokenKeyInput)(nil)).Elem(), EccTokenKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EccTokenKeyResponseInput)(nil)).Elem(), EccTokenKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncoderCustomPresetInput)(nil)).Elem(), EncoderCustomPresetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncoderCustomPresetResponseInput)(nil)).Elem(), EncoderCustomPresetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncoderProcessorInput)(nil)).Elem(), EncoderProcessorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncoderProcessorArrayInput)(nil)).Elem(), EncoderProcessorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncoderProcessorResponseInput)(nil)).Elem(), EncoderProcessorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncoderProcessorResponseArrayInput)(nil)).Elem(), EncoderProcessorResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncoderSystemPresetInput)(nil)).Elem(), EncoderSystemPresetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncoderSystemPresetResponseInput)(nil)).Elem(), EncoderSystemPresetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointResponseInput)(nil)).Elem(), EndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointResponseArrayInput)(nil)).Elem(), EndpointResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JwtAuthenticationInput)(nil)).Elem(), JwtAuthenticationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JwtAuthenticationPtrInput)(nil)).Elem(), JwtAuthenticationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JwtAuthenticationResponseInput)(nil)).Elem(), JwtAuthenticationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JwtAuthenticationResponsePtrInput)(nil)).Elem(), JwtAuthenticationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesInput)(nil)).Elem(), KeyVaultPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesPtrInput)(nil)).Elem(), KeyVaultPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesResponseInput)(nil)).Elem(), KeyVaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesResponsePtrInput)(nil)).Elem(), KeyVaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeInputInput)(nil)).Elem(), NodeInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeInputArrayInput)(nil)).Elem(), NodeInputArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeInputResponseInput)(nil)).Elem(), NodeInputResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeInputResponseArrayInput)(nil)).Elem(), NodeInputResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterDeclarationInput)(nil)).Elem(), ParameterDeclarationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterDeclarationArrayInput)(nil)).Elem(), ParameterDeclarationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterDeclarationResponseInput)(nil)).Elem(), ParameterDeclarationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterDeclarationResponseArrayInput)(nil)).Elem(), ParameterDeclarationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterDefinitionInput)(nil)).Elem(), ParameterDefinitionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterDefinitionArrayInput)(nil)).Elem(), ParameterDefinitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterDefinitionResponseInput)(nil)).Elem(), ParameterDefinitionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterDefinitionResponseArrayInput)(nil)).Elem(), ParameterDefinitionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PemCertificateListInput)(nil)).Elem(), PemCertificateListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PemCertificateListPtrInput)(nil)).Elem(), PemCertificateListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PemCertificateListResponseInput)(nil)).Elem(), PemCertificateListResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PemCertificateListResponsePtrInput)(nil)).Elem(), PemCertificateListResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineJobErrorResponseInput)(nil)).Elem(), PipelineJobErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineJobErrorResponsePtrInput)(nil)).Elem(), PipelineJobErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponseInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponsePtrInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponseInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponsePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityInput)(nil)).Elem(), ResourceIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityPtrInput)(nil)).Elem(), ResourceIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityResponseInput)(nil)).Elem(), ResourceIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityResponsePtrInput)(nil)).Elem(), ResourceIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RsaTokenKeyInput)(nil)).Elem(), RsaTokenKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RsaTokenKeyResponseInput)(nil)).Elem(), RsaTokenKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RtspSourceInput)(nil)).Elem(), RtspSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RtspSourceResponseInput)(nil)).Elem(), RtspSourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecureIotDeviceRemoteTunnelInput)(nil)).Elem(), SecureIotDeviceRemoteTunnelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecureIotDeviceRemoteTunnelPtrInput)(nil)).Elem(), SecureIotDeviceRemoteTunnelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecureIotDeviceRemoteTunnelResponseInput)(nil)).Elem(), SecureIotDeviceRemoteTunnelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecureIotDeviceRemoteTunnelResponsePtrInput)(nil)).Elem(), SecureIotDeviceRemoteTunnelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountInput)(nil)).Elem(), StorageAccountArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountArrayInput)(nil)).Elem(), StorageAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountResponseInput)(nil)).Elem(), StorageAccountResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountResponseArrayInput)(nil)).Elem(), StorageAccountResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsEndpointInput)(nil)).Elem(), TlsEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsEndpointResponseInput)(nil)).Elem(), TlsEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsValidationOptionsInput)(nil)).Elem(), TlsValidationOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsValidationOptionsPtrInput)(nil)).Elem(), TlsValidationOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsValidationOptionsResponseInput)(nil)).Elem(), TlsValidationOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsValidationOptionsResponsePtrInput)(nil)).Elem(), TlsValidationOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenClaimInput)(nil)).Elem(), TokenClaimArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenClaimArrayInput)(nil)).Elem(), TokenClaimArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenClaimResponseInput)(nil)).Elem(), TokenClaimResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenClaimResponseArrayInput)(nil)).Elem(), TokenClaimResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UnsecuredEndpointInput)(nil)).Elem(), UnsecuredEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UnsecuredEndpointResponseInput)(nil)).Elem(), UnsecuredEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedManagedIdentityResponseInput)(nil)).Elem(), UserAssignedManagedIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedManagedIdentityResponseMapInput)(nil)).Elem(), UserAssignedManagedIdentityResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*UsernamePasswordCredentialsInput)(nil)).Elem(), UsernamePasswordCredentialsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UsernamePasswordCredentialsPtrInput)(nil)).Elem(), UsernamePasswordCredentialsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UsernamePasswordCredentialsResponseInput)(nil)).Elem(), UsernamePasswordCredentialsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UsernamePasswordCredentialsResponsePtrInput)(nil)).Elem(), UsernamePasswordCredentialsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoAnalyzerIdentityInput)(nil)).Elem(), VideoAnalyzerIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoAnalyzerIdentityPtrInput)(nil)).Elem(), VideoAnalyzerIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoAnalyzerIdentityResponseInput)(nil)).Elem(), VideoAnalyzerIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoAnalyzerIdentityResponsePtrInput)(nil)).Elem(), VideoAnalyzerIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoCreationPropertiesInput)(nil)).Elem(), VideoCreationPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoCreationPropertiesPtrInput)(nil)).Elem(), VideoCreationPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoCreationPropertiesResponseInput)(nil)).Elem(), VideoCreationPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoCreationPropertiesResponsePtrInput)(nil)).Elem(), VideoCreationPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoEncoderH264Input)(nil)).Elem(), VideoEncoderH264Args{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoEncoderH264PtrInput)(nil)).Elem(), VideoEncoderH264Args{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoEncoderH264ResponseInput)(nil)).Elem(), VideoEncoderH264ResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoEncoderH264ResponsePtrInput)(nil)).Elem(), VideoEncoderH264ResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoFlagsResponseInput)(nil)).Elem(), VideoFlagsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoFlagsResponsePtrInput)(nil)).Elem(), VideoFlagsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoMediaInfoResponseInput)(nil)).Elem(), VideoMediaInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoMediaInfoResponsePtrInput)(nil)).Elem(), VideoMediaInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoPublishingOptionsInput)(nil)).Elem(), VideoPublishingOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoPublishingOptionsPtrInput)(nil)).Elem(), VideoPublishingOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoPublishingOptionsResponseInput)(nil)).Elem(), VideoPublishingOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoPublishingOptionsResponsePtrInput)(nil)).Elem(), VideoPublishingOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoScaleInput)(nil)).Elem(), VideoScaleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoScalePtrInput)(nil)).Elem(), VideoScaleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoScaleResponseInput)(nil)).Elem(), VideoScaleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoScaleResponsePtrInput)(nil)).Elem(), VideoScaleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoSequenceAbsoluteTimeMarkersInput)(nil)).Elem(), VideoSequenceAbsoluteTimeMarkersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoSequenceAbsoluteTimeMarkersResponseInput)(nil)).Elem(), VideoSequenceAbsoluteTimeMarkersResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoSinkInput)(nil)).Elem(), VideoSinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoSinkArrayInput)(nil)).Elem(), VideoSinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoSinkResponseInput)(nil)).Elem(), VideoSinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoSinkResponseArrayInput)(nil)).Elem(), VideoSinkResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoSourceInput)(nil)).Elem(), VideoSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoSourceResponseInput)(nil)).Elem(), VideoSourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoStreamingResponseInput)(nil)).Elem(), VideoStreamingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoStreamingResponsePtrInput)(nil)).Elem(), VideoStreamingResponseArgs{})
	pulumi.RegisterOutputType(AccountEncryptionOutput{})
	pulumi.RegisterOutputType(AccountEncryptionPtrOutput{})
	pulumi.RegisterOutputType(AccountEncryptionResponseOutput{})
	pulumi.RegisterOutputType(AccountEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(AudioEncoderAacOutput{})
	pulumi.RegisterOutputType(AudioEncoderAacPtrOutput{})
	pulumi.RegisterOutputType(AudioEncoderAacResponseOutput{})
	pulumi.RegisterOutputType(AudioEncoderAacResponsePtrOutput{})
	pulumi.RegisterOutputType(EccTokenKeyOutput{})
	pulumi.RegisterOutputType(EccTokenKeyResponseOutput{})
	pulumi.RegisterOutputType(EncoderCustomPresetOutput{})
	pulumi.RegisterOutputType(EncoderCustomPresetResponseOutput{})
	pulumi.RegisterOutputType(EncoderProcessorOutput{})
	pulumi.RegisterOutputType(EncoderProcessorArrayOutput{})
	pulumi.RegisterOutputType(EncoderProcessorResponseOutput{})
	pulumi.RegisterOutputType(EncoderProcessorResponseArrayOutput{})
	pulumi.RegisterOutputType(EncoderSystemPresetOutput{})
	pulumi.RegisterOutputType(EncoderSystemPresetResponseOutput{})
	pulumi.RegisterOutputType(EndpointResponseOutput{})
	pulumi.RegisterOutputType(EndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(JwtAuthenticationOutput{})
	pulumi.RegisterOutputType(JwtAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(JwtAuthenticationResponseOutput{})
	pulumi.RegisterOutputType(JwtAuthenticationResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NodeInputOutput{})
	pulumi.RegisterOutputType(NodeInputArrayOutput{})
	pulumi.RegisterOutputType(NodeInputResponseOutput{})
	pulumi.RegisterOutputType(NodeInputResponseArrayOutput{})
	pulumi.RegisterOutputType(ParameterDeclarationOutput{})
	pulumi.RegisterOutputType(ParameterDeclarationArrayOutput{})
	pulumi.RegisterOutputType(ParameterDeclarationResponseOutput{})
	pulumi.RegisterOutputType(ParameterDeclarationResponseArrayOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionArrayOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(PemCertificateListOutput{})
	pulumi.RegisterOutputType(PemCertificateListPtrOutput{})
	pulumi.RegisterOutputType(PemCertificateListResponseOutput{})
	pulumi.RegisterOutputType(PemCertificateListResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineJobErrorResponseOutput{})
	pulumi.RegisterOutputType(PipelineJobErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(RsaTokenKeyOutput{})
	pulumi.RegisterOutputType(RsaTokenKeyResponseOutput{})
	pulumi.RegisterOutputType(RtspSourceOutput{})
	pulumi.RegisterOutputType(RtspSourceResponseOutput{})
	pulumi.RegisterOutputType(SecureIotDeviceRemoteTunnelOutput{})
	pulumi.RegisterOutputType(SecureIotDeviceRemoteTunnelPtrOutput{})
	pulumi.RegisterOutputType(SecureIotDeviceRemoteTunnelResponseOutput{})
	pulumi.RegisterOutputType(SecureIotDeviceRemoteTunnelResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TlsEndpointOutput{})
	pulumi.RegisterOutputType(TlsEndpointResponseOutput{})
	pulumi.RegisterOutputType(TlsValidationOptionsOutput{})
	pulumi.RegisterOutputType(TlsValidationOptionsPtrOutput{})
	pulumi.RegisterOutputType(TlsValidationOptionsResponseOutput{})
	pulumi.RegisterOutputType(TlsValidationOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(TokenClaimOutput{})
	pulumi.RegisterOutputType(TokenClaimArrayOutput{})
	pulumi.RegisterOutputType(TokenClaimResponseOutput{})
	pulumi.RegisterOutputType(TokenClaimResponseArrayOutput{})
	pulumi.RegisterOutputType(UnsecuredEndpointOutput{})
	pulumi.RegisterOutputType(UnsecuredEndpointResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedManagedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(UsernamePasswordCredentialsOutput{})
	pulumi.RegisterOutputType(UsernamePasswordCredentialsPtrOutput{})
	pulumi.RegisterOutputType(UsernamePasswordCredentialsResponseOutput{})
	pulumi.RegisterOutputType(UsernamePasswordCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityPtrOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityResponseOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoCreationPropertiesOutput{})
	pulumi.RegisterOutputType(VideoCreationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VideoCreationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VideoCreationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoEncoderH264Output{})
	pulumi.RegisterOutputType(VideoEncoderH264PtrOutput{})
	pulumi.RegisterOutputType(VideoEncoderH264ResponseOutput{})
	pulumi.RegisterOutputType(VideoEncoderH264ResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoFlagsResponseOutput{})
	pulumi.RegisterOutputType(VideoFlagsResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoMediaInfoResponseOutput{})
	pulumi.RegisterOutputType(VideoMediaInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoPublishingOptionsOutput{})
	pulumi.RegisterOutputType(VideoPublishingOptionsPtrOutput{})
	pulumi.RegisterOutputType(VideoPublishingOptionsResponseOutput{})
	pulumi.RegisterOutputType(VideoPublishingOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoScaleOutput{})
	pulumi.RegisterOutputType(VideoScalePtrOutput{})
	pulumi.RegisterOutputType(VideoScaleResponseOutput{})
	pulumi.RegisterOutputType(VideoScaleResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoSequenceAbsoluteTimeMarkersOutput{})
	pulumi.RegisterOutputType(VideoSequenceAbsoluteTimeMarkersResponseOutput{})
	pulumi.RegisterOutputType(VideoSinkOutput{})
	pulumi.RegisterOutputType(VideoSinkArrayOutput{})
	pulumi.RegisterOutputType(VideoSinkResponseOutput{})
	pulumi.RegisterOutputType(VideoSinkResponseArrayOutput{})
	pulumi.RegisterOutputType(VideoSourceOutput{})
	pulumi.RegisterOutputType(VideoSourceResponseOutput{})
	pulumi.RegisterOutputType(VideoStreamingResponseOutput{})
	pulumi.RegisterOutputType(VideoStreamingResponsePtrOutput{})
}
