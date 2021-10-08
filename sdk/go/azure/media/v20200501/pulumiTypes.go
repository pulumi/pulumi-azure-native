


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AacAudio struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	Profile      *string `pulumi:"profile"`
	SamplingRate *int    `pulumi:"samplingRate"`
}





type AacAudioInput interface {
	pulumi.Input

	ToAacAudioOutput() AacAudioOutput
	ToAacAudioOutputWithContext(context.Context) AacAudioOutput
}

type AacAudioArgs struct {
	Bitrate      pulumi.IntPtrInput    `pulumi:"bitrate"`
	Channels     pulumi.IntPtrInput    `pulumi:"channels"`
	Label        pulumi.StringPtrInput `pulumi:"label"`
	OdataType    pulumi.StringInput    `pulumi:"odataType"`
	Profile      pulumi.StringPtrInput `pulumi:"profile"`
	SamplingRate pulumi.IntPtrInput    `pulumi:"samplingRate"`
}

func (AacAudioArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AacAudio)(nil)).Elem()
}

func (i AacAudioArgs) ToAacAudioOutput() AacAudioOutput {
	return i.ToAacAudioOutputWithContext(context.Background())
}

func (i AacAudioArgs) ToAacAudioOutputWithContext(ctx context.Context) AacAudioOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AacAudioOutput)
}

type AacAudioOutput struct{ *pulumi.OutputState }

func (AacAudioOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AacAudio)(nil)).Elem()
}

func (o AacAudioOutput) ToAacAudioOutput() AacAudioOutput {
	return o
}

func (o AacAudioOutput) ToAacAudioOutputWithContext(ctx context.Context) AacAudioOutput {
	return o
}

func (o AacAudioOutput) Bitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AacAudio) *int { return v.Bitrate }).(pulumi.IntPtrOutput)
}

func (o AacAudioOutput) Channels() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AacAudio) *int { return v.Channels }).(pulumi.IntPtrOutput)
}

func (o AacAudioOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AacAudio) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o AacAudioOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AacAudio) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o AacAudioOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AacAudio) *string { return v.Profile }).(pulumi.StringPtrOutput)
}

func (o AacAudioOutput) SamplingRate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AacAudio) *int { return v.SamplingRate }).(pulumi.IntPtrOutput)
}

type AacAudioResponse struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	Profile      *string `pulumi:"profile"`
	SamplingRate *int    `pulumi:"samplingRate"`
}





type AacAudioResponseInput interface {
	pulumi.Input

	ToAacAudioResponseOutput() AacAudioResponseOutput
	ToAacAudioResponseOutputWithContext(context.Context) AacAudioResponseOutput
}

type AacAudioResponseArgs struct {
	Bitrate      pulumi.IntPtrInput    `pulumi:"bitrate"`
	Channels     pulumi.IntPtrInput    `pulumi:"channels"`
	Label        pulumi.StringPtrInput `pulumi:"label"`
	OdataType    pulumi.StringInput    `pulumi:"odataType"`
	Profile      pulumi.StringPtrInput `pulumi:"profile"`
	SamplingRate pulumi.IntPtrInput    `pulumi:"samplingRate"`
}

func (AacAudioResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AacAudioResponse)(nil)).Elem()
}

func (i AacAudioResponseArgs) ToAacAudioResponseOutput() AacAudioResponseOutput {
	return i.ToAacAudioResponseOutputWithContext(context.Background())
}

func (i AacAudioResponseArgs) ToAacAudioResponseOutputWithContext(ctx context.Context) AacAudioResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AacAudioResponseOutput)
}

type AacAudioResponseOutput struct{ *pulumi.OutputState }

func (AacAudioResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AacAudioResponse)(nil)).Elem()
}

func (o AacAudioResponseOutput) ToAacAudioResponseOutput() AacAudioResponseOutput {
	return o
}

func (o AacAudioResponseOutput) ToAacAudioResponseOutputWithContext(ctx context.Context) AacAudioResponseOutput {
	return o
}

func (o AacAudioResponseOutput) Bitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AacAudioResponse) *int { return v.Bitrate }).(pulumi.IntPtrOutput)
}

func (o AacAudioResponseOutput) Channels() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AacAudioResponse) *int { return v.Channels }).(pulumi.IntPtrOutput)
}

func (o AacAudioResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AacAudioResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o AacAudioResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AacAudioResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o AacAudioResponseOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AacAudioResponse) *string { return v.Profile }).(pulumi.StringPtrOutput)
}

func (o AacAudioResponseOutput) SamplingRate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AacAudioResponse) *int { return v.SamplingRate }).(pulumi.IntPtrOutput)
}

type AbsoluteClipTime struct {
	OdataType string `pulumi:"odataType"`
	Time      string `pulumi:"time"`
}





type AbsoluteClipTimeInput interface {
	pulumi.Input

	ToAbsoluteClipTimeOutput() AbsoluteClipTimeOutput
	ToAbsoluteClipTimeOutputWithContext(context.Context) AbsoluteClipTimeOutput
}

type AbsoluteClipTimeArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
	Time      pulumi.StringInput `pulumi:"time"`
}

func (AbsoluteClipTimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AbsoluteClipTime)(nil)).Elem()
}

func (i AbsoluteClipTimeArgs) ToAbsoluteClipTimeOutput() AbsoluteClipTimeOutput {
	return i.ToAbsoluteClipTimeOutputWithContext(context.Background())
}

func (i AbsoluteClipTimeArgs) ToAbsoluteClipTimeOutputWithContext(ctx context.Context) AbsoluteClipTimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AbsoluteClipTimeOutput)
}

type AbsoluteClipTimeOutput struct{ *pulumi.OutputState }

func (AbsoluteClipTimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AbsoluteClipTime)(nil)).Elem()
}

func (o AbsoluteClipTimeOutput) ToAbsoluteClipTimeOutput() AbsoluteClipTimeOutput {
	return o
}

func (o AbsoluteClipTimeOutput) ToAbsoluteClipTimeOutputWithContext(ctx context.Context) AbsoluteClipTimeOutput {
	return o
}

func (o AbsoluteClipTimeOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AbsoluteClipTime) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o AbsoluteClipTimeOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v AbsoluteClipTime) string { return v.Time }).(pulumi.StringOutput)
}

type AbsoluteClipTimeResponse struct {
	OdataType string `pulumi:"odataType"`
	Time      string `pulumi:"time"`
}





type AbsoluteClipTimeResponseInput interface {
	pulumi.Input

	ToAbsoluteClipTimeResponseOutput() AbsoluteClipTimeResponseOutput
	ToAbsoluteClipTimeResponseOutputWithContext(context.Context) AbsoluteClipTimeResponseOutput
}

type AbsoluteClipTimeResponseArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
	Time      pulumi.StringInput `pulumi:"time"`
}

func (AbsoluteClipTimeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AbsoluteClipTimeResponse)(nil)).Elem()
}

func (i AbsoluteClipTimeResponseArgs) ToAbsoluteClipTimeResponseOutput() AbsoluteClipTimeResponseOutput {
	return i.ToAbsoluteClipTimeResponseOutputWithContext(context.Background())
}

func (i AbsoluteClipTimeResponseArgs) ToAbsoluteClipTimeResponseOutputWithContext(ctx context.Context) AbsoluteClipTimeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AbsoluteClipTimeResponseOutput)
}

type AbsoluteClipTimeResponseOutput struct{ *pulumi.OutputState }

func (AbsoluteClipTimeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AbsoluteClipTimeResponse)(nil)).Elem()
}

func (o AbsoluteClipTimeResponseOutput) ToAbsoluteClipTimeResponseOutput() AbsoluteClipTimeResponseOutput {
	return o
}

func (o AbsoluteClipTimeResponseOutput) ToAbsoluteClipTimeResponseOutputWithContext(ctx context.Context) AbsoluteClipTimeResponseOutput {
	return o
}

func (o AbsoluteClipTimeResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AbsoluteClipTimeResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o AbsoluteClipTimeResponseOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v AbsoluteClipTimeResponse) string { return v.Time }).(pulumi.StringOutput)
}

type AccountEncryption struct {
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	Type               string              `pulumi:"type"`
}





type AccountEncryptionInput interface {
	pulumi.Input

	ToAccountEncryptionOutput() AccountEncryptionOutput
	ToAccountEncryptionOutputWithContext(context.Context) AccountEncryptionOutput
}

type AccountEncryptionArgs struct {
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
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Type               string                      `pulumi:"type"`
}





type AccountEncryptionResponseInput interface {
	pulumi.Input

	ToAccountEncryptionResponseOutput() AccountEncryptionResponseOutput
	ToAccountEncryptionResponseOutputWithContext(context.Context) AccountEncryptionResponseOutput
}

type AccountEncryptionResponseArgs struct {
	KeyVaultProperties KeyVaultPropertiesResponsePtrInput `pulumi:"keyVaultProperties"`
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

func (o AccountEncryptionResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AccountEncryptionResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
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

func (o AccountEncryptionResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o AccountEncryptionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type AkamaiAccessControl struct {
	AkamaiSignatureHeaderAuthenticationKeyList []AkamaiSignatureHeaderAuthenticationKey `pulumi:"akamaiSignatureHeaderAuthenticationKeyList"`
}





type AkamaiAccessControlInput interface {
	pulumi.Input

	ToAkamaiAccessControlOutput() AkamaiAccessControlOutput
	ToAkamaiAccessControlOutputWithContext(context.Context) AkamaiAccessControlOutput
}

type AkamaiAccessControlArgs struct {
	AkamaiSignatureHeaderAuthenticationKeyList AkamaiSignatureHeaderAuthenticationKeyArrayInput `pulumi:"akamaiSignatureHeaderAuthenticationKeyList"`
}

func (AkamaiAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiAccessControl)(nil)).Elem()
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlOutput() AkamaiAccessControlOutput {
	return i.ToAkamaiAccessControlOutputWithContext(context.Background())
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlOutputWithContext(ctx context.Context) AkamaiAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlOutput)
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return i.ToAkamaiAccessControlPtrOutputWithContext(context.Background())
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlOutput).ToAkamaiAccessControlPtrOutputWithContext(ctx)
}









type AkamaiAccessControlPtrInput interface {
	pulumi.Input

	ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput
	ToAkamaiAccessControlPtrOutputWithContext(context.Context) AkamaiAccessControlPtrOutput
}

type akamaiAccessControlPtrType AkamaiAccessControlArgs

func AkamaiAccessControlPtr(v *AkamaiAccessControlArgs) AkamaiAccessControlPtrInput {
	return (*akamaiAccessControlPtrType)(v)
}

func (*akamaiAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AkamaiAccessControl)(nil)).Elem()
}

func (i *akamaiAccessControlPtrType) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return i.ToAkamaiAccessControlPtrOutputWithContext(context.Background())
}

func (i *akamaiAccessControlPtrType) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlPtrOutput)
}

type AkamaiAccessControlOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiAccessControl)(nil)).Elem()
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlOutput() AkamaiAccessControlOutput {
	return o
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlOutputWithContext(ctx context.Context) AkamaiAccessControlOutput {
	return o
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return o.ToAkamaiAccessControlPtrOutputWithContext(context.Background())
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AkamaiAccessControl) *AkamaiAccessControl {
		return &v
	}).(AkamaiAccessControlPtrOutput)
}

func (o AkamaiAccessControlOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o.ApplyT(func(v AkamaiAccessControl) []AkamaiSignatureHeaderAuthenticationKey {
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyArrayOutput)
}

type AkamaiAccessControlPtrOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AkamaiAccessControl)(nil)).Elem()
}

func (o AkamaiAccessControlPtrOutput) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return o
}

func (o AkamaiAccessControlPtrOutput) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return o
}

func (o AkamaiAccessControlPtrOutput) Elem() AkamaiAccessControlOutput {
	return o.ApplyT(func(v *AkamaiAccessControl) AkamaiAccessControl {
		if v != nil {
			return *v
		}
		var ret AkamaiAccessControl
		return ret
	}).(AkamaiAccessControlOutput)
}

func (o AkamaiAccessControlPtrOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o.ApplyT(func(v *AkamaiAccessControl) []AkamaiSignatureHeaderAuthenticationKey {
		if v == nil {
			return nil
		}
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyArrayOutput)
}

type AkamaiAccessControlResponse struct {
	AkamaiSignatureHeaderAuthenticationKeyList []AkamaiSignatureHeaderAuthenticationKeyResponse `pulumi:"akamaiSignatureHeaderAuthenticationKeyList"`
}





type AkamaiAccessControlResponseInput interface {
	pulumi.Input

	ToAkamaiAccessControlResponseOutput() AkamaiAccessControlResponseOutput
	ToAkamaiAccessControlResponseOutputWithContext(context.Context) AkamaiAccessControlResponseOutput
}

type AkamaiAccessControlResponseArgs struct {
	AkamaiSignatureHeaderAuthenticationKeyList AkamaiSignatureHeaderAuthenticationKeyResponseArrayInput `pulumi:"akamaiSignatureHeaderAuthenticationKeyList"`
}

func (AkamaiAccessControlResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiAccessControlResponse)(nil)).Elem()
}

func (i AkamaiAccessControlResponseArgs) ToAkamaiAccessControlResponseOutput() AkamaiAccessControlResponseOutput {
	return i.ToAkamaiAccessControlResponseOutputWithContext(context.Background())
}

func (i AkamaiAccessControlResponseArgs) ToAkamaiAccessControlResponseOutputWithContext(ctx context.Context) AkamaiAccessControlResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlResponseOutput)
}

func (i AkamaiAccessControlResponseArgs) ToAkamaiAccessControlResponsePtrOutput() AkamaiAccessControlResponsePtrOutput {
	return i.ToAkamaiAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i AkamaiAccessControlResponseArgs) ToAkamaiAccessControlResponsePtrOutputWithContext(ctx context.Context) AkamaiAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlResponseOutput).ToAkamaiAccessControlResponsePtrOutputWithContext(ctx)
}









type AkamaiAccessControlResponsePtrInput interface {
	pulumi.Input

	ToAkamaiAccessControlResponsePtrOutput() AkamaiAccessControlResponsePtrOutput
	ToAkamaiAccessControlResponsePtrOutputWithContext(context.Context) AkamaiAccessControlResponsePtrOutput
}

type akamaiAccessControlResponsePtrType AkamaiAccessControlResponseArgs

func AkamaiAccessControlResponsePtr(v *AkamaiAccessControlResponseArgs) AkamaiAccessControlResponsePtrInput {
	return (*akamaiAccessControlResponsePtrType)(v)
}

func (*akamaiAccessControlResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AkamaiAccessControlResponse)(nil)).Elem()
}

func (i *akamaiAccessControlResponsePtrType) ToAkamaiAccessControlResponsePtrOutput() AkamaiAccessControlResponsePtrOutput {
	return i.ToAkamaiAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i *akamaiAccessControlResponsePtrType) ToAkamaiAccessControlResponsePtrOutputWithContext(ctx context.Context) AkamaiAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlResponsePtrOutput)
}

type AkamaiAccessControlResponseOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiAccessControlResponse)(nil)).Elem()
}

func (o AkamaiAccessControlResponseOutput) ToAkamaiAccessControlResponseOutput() AkamaiAccessControlResponseOutput {
	return o
}

func (o AkamaiAccessControlResponseOutput) ToAkamaiAccessControlResponseOutputWithContext(ctx context.Context) AkamaiAccessControlResponseOutput {
	return o
}

func (o AkamaiAccessControlResponseOutput) ToAkamaiAccessControlResponsePtrOutput() AkamaiAccessControlResponsePtrOutput {
	return o.ToAkamaiAccessControlResponsePtrOutputWithContext(context.Background())
}

func (o AkamaiAccessControlResponseOutput) ToAkamaiAccessControlResponsePtrOutputWithContext(ctx context.Context) AkamaiAccessControlResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AkamaiAccessControlResponse) *AkamaiAccessControlResponse {
		return &v
	}).(AkamaiAccessControlResponsePtrOutput)
}

func (o AkamaiAccessControlResponseOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o.ApplyT(func(v AkamaiAccessControlResponse) []AkamaiSignatureHeaderAuthenticationKeyResponse {
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput)
}

type AkamaiAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AkamaiAccessControlResponse)(nil)).Elem()
}

func (o AkamaiAccessControlResponsePtrOutput) ToAkamaiAccessControlResponsePtrOutput() AkamaiAccessControlResponsePtrOutput {
	return o
}

func (o AkamaiAccessControlResponsePtrOutput) ToAkamaiAccessControlResponsePtrOutputWithContext(ctx context.Context) AkamaiAccessControlResponsePtrOutput {
	return o
}

func (o AkamaiAccessControlResponsePtrOutput) Elem() AkamaiAccessControlResponseOutput {
	return o.ApplyT(func(v *AkamaiAccessControlResponse) AkamaiAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret AkamaiAccessControlResponse
		return ret
	}).(AkamaiAccessControlResponseOutput)
}

func (o AkamaiAccessControlResponsePtrOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o.ApplyT(func(v *AkamaiAccessControlResponse) []AkamaiSignatureHeaderAuthenticationKeyResponse {
		if v == nil {
			return nil
		}
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput)
}

type AkamaiSignatureHeaderAuthenticationKey struct {
	Base64Key  *string `pulumi:"base64Key"`
	Expiration *string `pulumi:"expiration"`
	Identifier *string `pulumi:"identifier"`
}





type AkamaiSignatureHeaderAuthenticationKeyInput interface {
	pulumi.Input

	ToAkamaiSignatureHeaderAuthenticationKeyOutput() AkamaiSignatureHeaderAuthenticationKeyOutput
	ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(context.Context) AkamaiSignatureHeaderAuthenticationKeyOutput
}

type AkamaiSignatureHeaderAuthenticationKeyArgs struct {
	Base64Key  pulumi.StringPtrInput `pulumi:"base64Key"`
	Expiration pulumi.StringPtrInput `pulumi:"expiration"`
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
}

func (AkamaiSignatureHeaderAuthenticationKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (i AkamaiSignatureHeaderAuthenticationKeyArgs) ToAkamaiSignatureHeaderAuthenticationKeyOutput() AkamaiSignatureHeaderAuthenticationKeyOutput {
	return i.ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(context.Background())
}

func (i AkamaiSignatureHeaderAuthenticationKeyArgs) ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiSignatureHeaderAuthenticationKeyOutput)
}





type AkamaiSignatureHeaderAuthenticationKeyArrayInput interface {
	pulumi.Input

	ToAkamaiSignatureHeaderAuthenticationKeyArrayOutput() AkamaiSignatureHeaderAuthenticationKeyArrayOutput
	ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(context.Context) AkamaiSignatureHeaderAuthenticationKeyArrayOutput
}

type AkamaiSignatureHeaderAuthenticationKeyArray []AkamaiSignatureHeaderAuthenticationKeyInput

func (AkamaiSignatureHeaderAuthenticationKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (i AkamaiSignatureHeaderAuthenticationKeyArray) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutput() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return i.ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(context.Background())
}

func (i AkamaiSignatureHeaderAuthenticationKeyArray) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiSignatureHeaderAuthenticationKeyArrayOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) ToAkamaiSignatureHeaderAuthenticationKeyOutput() AkamaiSignatureHeaderAuthenticationKeyOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) Base64Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKey) *string { return v.Base64Key }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKey) *string { return v.Expiration }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKey) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyArrayOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutput() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyArrayOutput) Index(i pulumi.IntInput) AkamaiSignatureHeaderAuthenticationKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AkamaiSignatureHeaderAuthenticationKey {
		return vs[0].([]AkamaiSignatureHeaderAuthenticationKey)[vs[1].(int)]
	}).(AkamaiSignatureHeaderAuthenticationKeyOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyResponse struct {
	Base64Key  *string `pulumi:"base64Key"`
	Expiration *string `pulumi:"expiration"`
	Identifier *string `pulumi:"identifier"`
}





type AkamaiSignatureHeaderAuthenticationKeyResponseInput interface {
	pulumi.Input

	ToAkamaiSignatureHeaderAuthenticationKeyResponseOutput() AkamaiSignatureHeaderAuthenticationKeyResponseOutput
	ToAkamaiSignatureHeaderAuthenticationKeyResponseOutputWithContext(context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseOutput
}

type AkamaiSignatureHeaderAuthenticationKeyResponseArgs struct {
	Base64Key  pulumi.StringPtrInput `pulumi:"base64Key"`
	Expiration pulumi.StringPtrInput `pulumi:"expiration"`
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
}

func (AkamaiSignatureHeaderAuthenticationKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyResponse)(nil)).Elem()
}

func (i AkamaiSignatureHeaderAuthenticationKeyResponseArgs) ToAkamaiSignatureHeaderAuthenticationKeyResponseOutput() AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return i.ToAkamaiSignatureHeaderAuthenticationKeyResponseOutputWithContext(context.Background())
}

func (i AkamaiSignatureHeaderAuthenticationKeyResponseArgs) ToAkamaiSignatureHeaderAuthenticationKeyResponseOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiSignatureHeaderAuthenticationKeyResponseOutput)
}





type AkamaiSignatureHeaderAuthenticationKeyResponseArrayInput interface {
	pulumi.Input

	ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput
	ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutputWithContext(context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput
}

type AkamaiSignatureHeaderAuthenticationKeyResponseArray []AkamaiSignatureHeaderAuthenticationKeyResponseInput

func (AkamaiSignatureHeaderAuthenticationKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AkamaiSignatureHeaderAuthenticationKeyResponse)(nil)).Elem()
}

func (i AkamaiSignatureHeaderAuthenticationKeyResponseArray) ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return i.ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutputWithContext(context.Background())
}

func (i AkamaiSignatureHeaderAuthenticationKeyResponseArray) ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyResponseOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyResponse)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseOutput() AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) Base64Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKeyResponse) *string { return v.Base64Key }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKeyResponse) *string { return v.Expiration }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKeyResponse) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AkamaiSignatureHeaderAuthenticationKeyResponse)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) Index(i pulumi.IntInput) AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AkamaiSignatureHeaderAuthenticationKeyResponse {
		return vs[0].([]AkamaiSignatureHeaderAuthenticationKeyResponse)[vs[1].(int)]
	}).(AkamaiSignatureHeaderAuthenticationKeyResponseOutput)
}

type AssetFileEncryptionMetadataResponse struct {
	AssetFileId          string  `pulumi:"assetFileId"`
	AssetFileName        *string `pulumi:"assetFileName"`
	InitializationVector *string `pulumi:"initializationVector"`
}





type AssetFileEncryptionMetadataResponseInput interface {
	pulumi.Input

	ToAssetFileEncryptionMetadataResponseOutput() AssetFileEncryptionMetadataResponseOutput
	ToAssetFileEncryptionMetadataResponseOutputWithContext(context.Context) AssetFileEncryptionMetadataResponseOutput
}

type AssetFileEncryptionMetadataResponseArgs struct {
	AssetFileId          pulumi.StringInput    `pulumi:"assetFileId"`
	AssetFileName        pulumi.StringPtrInput `pulumi:"assetFileName"`
	InitializationVector pulumi.StringPtrInput `pulumi:"initializationVector"`
}

func (AssetFileEncryptionMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetFileEncryptionMetadataResponse)(nil)).Elem()
}

func (i AssetFileEncryptionMetadataResponseArgs) ToAssetFileEncryptionMetadataResponseOutput() AssetFileEncryptionMetadataResponseOutput {
	return i.ToAssetFileEncryptionMetadataResponseOutputWithContext(context.Background())
}

func (i AssetFileEncryptionMetadataResponseArgs) ToAssetFileEncryptionMetadataResponseOutputWithContext(ctx context.Context) AssetFileEncryptionMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetFileEncryptionMetadataResponseOutput)
}





type AssetFileEncryptionMetadataResponseArrayInput interface {
	pulumi.Input

	ToAssetFileEncryptionMetadataResponseArrayOutput() AssetFileEncryptionMetadataResponseArrayOutput
	ToAssetFileEncryptionMetadataResponseArrayOutputWithContext(context.Context) AssetFileEncryptionMetadataResponseArrayOutput
}

type AssetFileEncryptionMetadataResponseArray []AssetFileEncryptionMetadataResponseInput

func (AssetFileEncryptionMetadataResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssetFileEncryptionMetadataResponse)(nil)).Elem()
}

func (i AssetFileEncryptionMetadataResponseArray) ToAssetFileEncryptionMetadataResponseArrayOutput() AssetFileEncryptionMetadataResponseArrayOutput {
	return i.ToAssetFileEncryptionMetadataResponseArrayOutputWithContext(context.Background())
}

func (i AssetFileEncryptionMetadataResponseArray) ToAssetFileEncryptionMetadataResponseArrayOutputWithContext(ctx context.Context) AssetFileEncryptionMetadataResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetFileEncryptionMetadataResponseArrayOutput)
}

type AssetFileEncryptionMetadataResponseOutput struct{ *pulumi.OutputState }

func (AssetFileEncryptionMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetFileEncryptionMetadataResponse)(nil)).Elem()
}

func (o AssetFileEncryptionMetadataResponseOutput) ToAssetFileEncryptionMetadataResponseOutput() AssetFileEncryptionMetadataResponseOutput {
	return o
}

func (o AssetFileEncryptionMetadataResponseOutput) ToAssetFileEncryptionMetadataResponseOutputWithContext(ctx context.Context) AssetFileEncryptionMetadataResponseOutput {
	return o
}

func (o AssetFileEncryptionMetadataResponseOutput) AssetFileId() pulumi.StringOutput {
	return o.ApplyT(func(v AssetFileEncryptionMetadataResponse) string { return v.AssetFileId }).(pulumi.StringOutput)
}

func (o AssetFileEncryptionMetadataResponseOutput) AssetFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssetFileEncryptionMetadataResponse) *string { return v.AssetFileName }).(pulumi.StringPtrOutput)
}

func (o AssetFileEncryptionMetadataResponseOutput) InitializationVector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssetFileEncryptionMetadataResponse) *string { return v.InitializationVector }).(pulumi.StringPtrOutput)
}

type AssetFileEncryptionMetadataResponseArrayOutput struct{ *pulumi.OutputState }

func (AssetFileEncryptionMetadataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssetFileEncryptionMetadataResponse)(nil)).Elem()
}

func (o AssetFileEncryptionMetadataResponseArrayOutput) ToAssetFileEncryptionMetadataResponseArrayOutput() AssetFileEncryptionMetadataResponseArrayOutput {
	return o
}

func (o AssetFileEncryptionMetadataResponseArrayOutput) ToAssetFileEncryptionMetadataResponseArrayOutputWithContext(ctx context.Context) AssetFileEncryptionMetadataResponseArrayOutput {
	return o
}

func (o AssetFileEncryptionMetadataResponseArrayOutput) Index(i pulumi.IntInput) AssetFileEncryptionMetadataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AssetFileEncryptionMetadataResponse {
		return vs[0].([]AssetFileEncryptionMetadataResponse)[vs[1].(int)]
	}).(AssetFileEncryptionMetadataResponseOutput)
}

type AssetStreamingLocatorResponse struct {
	AssetName                   string `pulumi:"assetName"`
	Created                     string `pulumi:"created"`
	DefaultContentKeyPolicyName string `pulumi:"defaultContentKeyPolicyName"`
	EndTime                     string `pulumi:"endTime"`
	Name                        string `pulumi:"name"`
	StartTime                   string `pulumi:"startTime"`
	StreamingLocatorId          string `pulumi:"streamingLocatorId"`
	StreamingPolicyName         string `pulumi:"streamingPolicyName"`
}





type AssetStreamingLocatorResponseInput interface {
	pulumi.Input

	ToAssetStreamingLocatorResponseOutput() AssetStreamingLocatorResponseOutput
	ToAssetStreamingLocatorResponseOutputWithContext(context.Context) AssetStreamingLocatorResponseOutput
}

type AssetStreamingLocatorResponseArgs struct {
	AssetName                   pulumi.StringInput `pulumi:"assetName"`
	Created                     pulumi.StringInput `pulumi:"created"`
	DefaultContentKeyPolicyName pulumi.StringInput `pulumi:"defaultContentKeyPolicyName"`
	EndTime                     pulumi.StringInput `pulumi:"endTime"`
	Name                        pulumi.StringInput `pulumi:"name"`
	StartTime                   pulumi.StringInput `pulumi:"startTime"`
	StreamingLocatorId          pulumi.StringInput `pulumi:"streamingLocatorId"`
	StreamingPolicyName         pulumi.StringInput `pulumi:"streamingPolicyName"`
}

func (AssetStreamingLocatorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetStreamingLocatorResponse)(nil)).Elem()
}

func (i AssetStreamingLocatorResponseArgs) ToAssetStreamingLocatorResponseOutput() AssetStreamingLocatorResponseOutput {
	return i.ToAssetStreamingLocatorResponseOutputWithContext(context.Background())
}

func (i AssetStreamingLocatorResponseArgs) ToAssetStreamingLocatorResponseOutputWithContext(ctx context.Context) AssetStreamingLocatorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetStreamingLocatorResponseOutput)
}





type AssetStreamingLocatorResponseArrayInput interface {
	pulumi.Input

	ToAssetStreamingLocatorResponseArrayOutput() AssetStreamingLocatorResponseArrayOutput
	ToAssetStreamingLocatorResponseArrayOutputWithContext(context.Context) AssetStreamingLocatorResponseArrayOutput
}

type AssetStreamingLocatorResponseArray []AssetStreamingLocatorResponseInput

func (AssetStreamingLocatorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssetStreamingLocatorResponse)(nil)).Elem()
}

func (i AssetStreamingLocatorResponseArray) ToAssetStreamingLocatorResponseArrayOutput() AssetStreamingLocatorResponseArrayOutput {
	return i.ToAssetStreamingLocatorResponseArrayOutputWithContext(context.Background())
}

func (i AssetStreamingLocatorResponseArray) ToAssetStreamingLocatorResponseArrayOutputWithContext(ctx context.Context) AssetStreamingLocatorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetStreamingLocatorResponseArrayOutput)
}

type AssetStreamingLocatorResponseOutput struct{ *pulumi.OutputState }

func (AssetStreamingLocatorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetStreamingLocatorResponse)(nil)).Elem()
}

func (o AssetStreamingLocatorResponseOutput) ToAssetStreamingLocatorResponseOutput() AssetStreamingLocatorResponseOutput {
	return o
}

func (o AssetStreamingLocatorResponseOutput) ToAssetStreamingLocatorResponseOutputWithContext(ctx context.Context) AssetStreamingLocatorResponseOutput {
	return o
}

func (o AssetStreamingLocatorResponseOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v AssetStreamingLocatorResponse) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o AssetStreamingLocatorResponseOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v AssetStreamingLocatorResponse) string { return v.Created }).(pulumi.StringOutput)
}

func (o AssetStreamingLocatorResponseOutput) DefaultContentKeyPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v AssetStreamingLocatorResponse) string { return v.DefaultContentKeyPolicyName }).(pulumi.StringOutput)
}

func (o AssetStreamingLocatorResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v AssetStreamingLocatorResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o AssetStreamingLocatorResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AssetStreamingLocatorResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AssetStreamingLocatorResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v AssetStreamingLocatorResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o AssetStreamingLocatorResponseOutput) StreamingLocatorId() pulumi.StringOutput {
	return o.ApplyT(func(v AssetStreamingLocatorResponse) string { return v.StreamingLocatorId }).(pulumi.StringOutput)
}

func (o AssetStreamingLocatorResponseOutput) StreamingPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v AssetStreamingLocatorResponse) string { return v.StreamingPolicyName }).(pulumi.StringOutput)
}

type AssetStreamingLocatorResponseArrayOutput struct{ *pulumi.OutputState }

func (AssetStreamingLocatorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssetStreamingLocatorResponse)(nil)).Elem()
}

func (o AssetStreamingLocatorResponseArrayOutput) ToAssetStreamingLocatorResponseArrayOutput() AssetStreamingLocatorResponseArrayOutput {
	return o
}

func (o AssetStreamingLocatorResponseArrayOutput) ToAssetStreamingLocatorResponseArrayOutputWithContext(ctx context.Context) AssetStreamingLocatorResponseArrayOutput {
	return o
}

func (o AssetStreamingLocatorResponseArrayOutput) Index(i pulumi.IntInput) AssetStreamingLocatorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AssetStreamingLocatorResponse {
		return vs[0].([]AssetStreamingLocatorResponse)[vs[1].(int)]
	}).(AssetStreamingLocatorResponseOutput)
}

type Audio struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	SamplingRate *int    `pulumi:"samplingRate"`
}





type AudioInput interface {
	pulumi.Input

	ToAudioOutput() AudioOutput
	ToAudioOutputWithContext(context.Context) AudioOutput
}

type AudioArgs struct {
	Bitrate      pulumi.IntPtrInput    `pulumi:"bitrate"`
	Channels     pulumi.IntPtrInput    `pulumi:"channels"`
	Label        pulumi.StringPtrInput `pulumi:"label"`
	OdataType    pulumi.StringInput    `pulumi:"odataType"`
	SamplingRate pulumi.IntPtrInput    `pulumi:"samplingRate"`
}

func (AudioArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Audio)(nil)).Elem()
}

func (i AudioArgs) ToAudioOutput() AudioOutput {
	return i.ToAudioOutputWithContext(context.Background())
}

func (i AudioArgs) ToAudioOutputWithContext(ctx context.Context) AudioOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioOutput)
}

type AudioOutput struct{ *pulumi.OutputState }

func (AudioOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Audio)(nil)).Elem()
}

func (o AudioOutput) ToAudioOutput() AudioOutput {
	return o
}

func (o AudioOutput) ToAudioOutputWithContext(ctx context.Context) AudioOutput {
	return o
}

func (o AudioOutput) Bitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Audio) *int { return v.Bitrate }).(pulumi.IntPtrOutput)
}

func (o AudioOutput) Channels() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Audio) *int { return v.Channels }).(pulumi.IntPtrOutput)
}

func (o AudioOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Audio) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o AudioOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v Audio) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o AudioOutput) SamplingRate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Audio) *int { return v.SamplingRate }).(pulumi.IntPtrOutput)
}

type AudioAnalyzerPreset struct {
	AudioLanguage       *string           `pulumi:"audioLanguage"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
}





type AudioAnalyzerPresetInput interface {
	pulumi.Input

	ToAudioAnalyzerPresetOutput() AudioAnalyzerPresetOutput
	ToAudioAnalyzerPresetOutputWithContext(context.Context) AudioAnalyzerPresetOutput
}

type AudioAnalyzerPresetArgs struct {
	AudioLanguage       pulumi.StringPtrInput `pulumi:"audioLanguage"`
	ExperimentalOptions pulumi.StringMapInput `pulumi:"experimentalOptions"`
	Mode                pulumi.StringPtrInput `pulumi:"mode"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
}

func (AudioAnalyzerPresetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioAnalyzerPreset)(nil)).Elem()
}

func (i AudioAnalyzerPresetArgs) ToAudioAnalyzerPresetOutput() AudioAnalyzerPresetOutput {
	return i.ToAudioAnalyzerPresetOutputWithContext(context.Background())
}

func (i AudioAnalyzerPresetArgs) ToAudioAnalyzerPresetOutputWithContext(ctx context.Context) AudioAnalyzerPresetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioAnalyzerPresetOutput)
}

type AudioAnalyzerPresetOutput struct{ *pulumi.OutputState }

func (AudioAnalyzerPresetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioAnalyzerPreset)(nil)).Elem()
}

func (o AudioAnalyzerPresetOutput) ToAudioAnalyzerPresetOutput() AudioAnalyzerPresetOutput {
	return o
}

func (o AudioAnalyzerPresetOutput) ToAudioAnalyzerPresetOutputWithContext(ctx context.Context) AudioAnalyzerPresetOutput {
	return o
}

func (o AudioAnalyzerPresetOutput) AudioLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioAnalyzerPreset) *string { return v.AudioLanguage }).(pulumi.StringPtrOutput)
}

func (o AudioAnalyzerPresetOutput) ExperimentalOptions() pulumi.StringMapOutput {
	return o.ApplyT(func(v AudioAnalyzerPreset) map[string]string { return v.ExperimentalOptions }).(pulumi.StringMapOutput)
}

func (o AudioAnalyzerPresetOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioAnalyzerPreset) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o AudioAnalyzerPresetOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AudioAnalyzerPreset) string { return v.OdataType }).(pulumi.StringOutput)
}

type AudioAnalyzerPresetResponse struct {
	AudioLanguage       *string           `pulumi:"audioLanguage"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
}





type AudioAnalyzerPresetResponseInput interface {
	pulumi.Input

	ToAudioAnalyzerPresetResponseOutput() AudioAnalyzerPresetResponseOutput
	ToAudioAnalyzerPresetResponseOutputWithContext(context.Context) AudioAnalyzerPresetResponseOutput
}

type AudioAnalyzerPresetResponseArgs struct {
	AudioLanguage       pulumi.StringPtrInput `pulumi:"audioLanguage"`
	ExperimentalOptions pulumi.StringMapInput `pulumi:"experimentalOptions"`
	Mode                pulumi.StringPtrInput `pulumi:"mode"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
}

func (AudioAnalyzerPresetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioAnalyzerPresetResponse)(nil)).Elem()
}

func (i AudioAnalyzerPresetResponseArgs) ToAudioAnalyzerPresetResponseOutput() AudioAnalyzerPresetResponseOutput {
	return i.ToAudioAnalyzerPresetResponseOutputWithContext(context.Background())
}

func (i AudioAnalyzerPresetResponseArgs) ToAudioAnalyzerPresetResponseOutputWithContext(ctx context.Context) AudioAnalyzerPresetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioAnalyzerPresetResponseOutput)
}

type AudioAnalyzerPresetResponseOutput struct{ *pulumi.OutputState }

func (AudioAnalyzerPresetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioAnalyzerPresetResponse)(nil)).Elem()
}

func (o AudioAnalyzerPresetResponseOutput) ToAudioAnalyzerPresetResponseOutput() AudioAnalyzerPresetResponseOutput {
	return o
}

func (o AudioAnalyzerPresetResponseOutput) ToAudioAnalyzerPresetResponseOutputWithContext(ctx context.Context) AudioAnalyzerPresetResponseOutput {
	return o
}

func (o AudioAnalyzerPresetResponseOutput) AudioLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioAnalyzerPresetResponse) *string { return v.AudioLanguage }).(pulumi.StringPtrOutput)
}

func (o AudioAnalyzerPresetResponseOutput) ExperimentalOptions() pulumi.StringMapOutput {
	return o.ApplyT(func(v AudioAnalyzerPresetResponse) map[string]string { return v.ExperimentalOptions }).(pulumi.StringMapOutput)
}

func (o AudioAnalyzerPresetResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioAnalyzerPresetResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o AudioAnalyzerPresetResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AudioAnalyzerPresetResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type AudioOverlay struct {
	AudioGainLevel  *float64 `pulumi:"audioGainLevel"`
	End             *string  `pulumi:"end"`
	FadeInDuration  *string  `pulumi:"fadeInDuration"`
	FadeOutDuration *string  `pulumi:"fadeOutDuration"`
	InputLabel      string   `pulumi:"inputLabel"`
	OdataType       string   `pulumi:"odataType"`
	Start           *string  `pulumi:"start"`
}





type AudioOverlayInput interface {
	pulumi.Input

	ToAudioOverlayOutput() AudioOverlayOutput
	ToAudioOverlayOutputWithContext(context.Context) AudioOverlayOutput
}

type AudioOverlayArgs struct {
	AudioGainLevel  pulumi.Float64PtrInput `pulumi:"audioGainLevel"`
	End             pulumi.StringPtrInput  `pulumi:"end"`
	FadeInDuration  pulumi.StringPtrInput  `pulumi:"fadeInDuration"`
	FadeOutDuration pulumi.StringPtrInput  `pulumi:"fadeOutDuration"`
	InputLabel      pulumi.StringInput     `pulumi:"inputLabel"`
	OdataType       pulumi.StringInput     `pulumi:"odataType"`
	Start           pulumi.StringPtrInput  `pulumi:"start"`
}

func (AudioOverlayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioOverlay)(nil)).Elem()
}

func (i AudioOverlayArgs) ToAudioOverlayOutput() AudioOverlayOutput {
	return i.ToAudioOverlayOutputWithContext(context.Background())
}

func (i AudioOverlayArgs) ToAudioOverlayOutputWithContext(ctx context.Context) AudioOverlayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioOverlayOutput)
}

type AudioOverlayOutput struct{ *pulumi.OutputState }

func (AudioOverlayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioOverlay)(nil)).Elem()
}

func (o AudioOverlayOutput) ToAudioOverlayOutput() AudioOverlayOutput {
	return o
}

func (o AudioOverlayOutput) ToAudioOverlayOutputWithContext(ctx context.Context) AudioOverlayOutput {
	return o
}

func (o AudioOverlayOutput) AudioGainLevel() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AudioOverlay) *float64 { return v.AudioGainLevel }).(pulumi.Float64PtrOutput)
}

func (o AudioOverlayOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioOverlay) *string { return v.End }).(pulumi.StringPtrOutput)
}

func (o AudioOverlayOutput) FadeInDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioOverlay) *string { return v.FadeInDuration }).(pulumi.StringPtrOutput)
}

func (o AudioOverlayOutput) FadeOutDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioOverlay) *string { return v.FadeOutDuration }).(pulumi.StringPtrOutput)
}

func (o AudioOverlayOutput) InputLabel() pulumi.StringOutput {
	return o.ApplyT(func(v AudioOverlay) string { return v.InputLabel }).(pulumi.StringOutput)
}

func (o AudioOverlayOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AudioOverlay) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o AudioOverlayOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioOverlay) *string { return v.Start }).(pulumi.StringPtrOutput)
}

type AudioOverlayResponse struct {
	AudioGainLevel  *float64 `pulumi:"audioGainLevel"`
	End             *string  `pulumi:"end"`
	FadeInDuration  *string  `pulumi:"fadeInDuration"`
	FadeOutDuration *string  `pulumi:"fadeOutDuration"`
	InputLabel      string   `pulumi:"inputLabel"`
	OdataType       string   `pulumi:"odataType"`
	Start           *string  `pulumi:"start"`
}





type AudioOverlayResponseInput interface {
	pulumi.Input

	ToAudioOverlayResponseOutput() AudioOverlayResponseOutput
	ToAudioOverlayResponseOutputWithContext(context.Context) AudioOverlayResponseOutput
}

type AudioOverlayResponseArgs struct {
	AudioGainLevel  pulumi.Float64PtrInput `pulumi:"audioGainLevel"`
	End             pulumi.StringPtrInput  `pulumi:"end"`
	FadeInDuration  pulumi.StringPtrInput  `pulumi:"fadeInDuration"`
	FadeOutDuration pulumi.StringPtrInput  `pulumi:"fadeOutDuration"`
	InputLabel      pulumi.StringInput     `pulumi:"inputLabel"`
	OdataType       pulumi.StringInput     `pulumi:"odataType"`
	Start           pulumi.StringPtrInput  `pulumi:"start"`
}

func (AudioOverlayResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioOverlayResponse)(nil)).Elem()
}

func (i AudioOverlayResponseArgs) ToAudioOverlayResponseOutput() AudioOverlayResponseOutput {
	return i.ToAudioOverlayResponseOutputWithContext(context.Background())
}

func (i AudioOverlayResponseArgs) ToAudioOverlayResponseOutputWithContext(ctx context.Context) AudioOverlayResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioOverlayResponseOutput)
}

type AudioOverlayResponseOutput struct{ *pulumi.OutputState }

func (AudioOverlayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioOverlayResponse)(nil)).Elem()
}

func (o AudioOverlayResponseOutput) ToAudioOverlayResponseOutput() AudioOverlayResponseOutput {
	return o
}

func (o AudioOverlayResponseOutput) ToAudioOverlayResponseOutputWithContext(ctx context.Context) AudioOverlayResponseOutput {
	return o
}

func (o AudioOverlayResponseOutput) AudioGainLevel() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AudioOverlayResponse) *float64 { return v.AudioGainLevel }).(pulumi.Float64PtrOutput)
}

func (o AudioOverlayResponseOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioOverlayResponse) *string { return v.End }).(pulumi.StringPtrOutput)
}

func (o AudioOverlayResponseOutput) FadeInDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioOverlayResponse) *string { return v.FadeInDuration }).(pulumi.StringPtrOutput)
}

func (o AudioOverlayResponseOutput) FadeOutDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioOverlayResponse) *string { return v.FadeOutDuration }).(pulumi.StringPtrOutput)
}

func (o AudioOverlayResponseOutput) InputLabel() pulumi.StringOutput {
	return o.ApplyT(func(v AudioOverlayResponse) string { return v.InputLabel }).(pulumi.StringOutput)
}

func (o AudioOverlayResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AudioOverlayResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o AudioOverlayResponseOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioOverlayResponse) *string { return v.Start }).(pulumi.StringPtrOutput)
}

type AudioResponse struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	SamplingRate *int    `pulumi:"samplingRate"`
}





type AudioResponseInput interface {
	pulumi.Input

	ToAudioResponseOutput() AudioResponseOutput
	ToAudioResponseOutputWithContext(context.Context) AudioResponseOutput
}

type AudioResponseArgs struct {
	Bitrate      pulumi.IntPtrInput    `pulumi:"bitrate"`
	Channels     pulumi.IntPtrInput    `pulumi:"channels"`
	Label        pulumi.StringPtrInput `pulumi:"label"`
	OdataType    pulumi.StringInput    `pulumi:"odataType"`
	SamplingRate pulumi.IntPtrInput    `pulumi:"samplingRate"`
}

func (AudioResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioResponse)(nil)).Elem()
}

func (i AudioResponseArgs) ToAudioResponseOutput() AudioResponseOutput {
	return i.ToAudioResponseOutputWithContext(context.Background())
}

func (i AudioResponseArgs) ToAudioResponseOutputWithContext(ctx context.Context) AudioResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioResponseOutput)
}

type AudioResponseOutput struct{ *pulumi.OutputState }

func (AudioResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioResponse)(nil)).Elem()
}

func (o AudioResponseOutput) ToAudioResponseOutput() AudioResponseOutput {
	return o
}

func (o AudioResponseOutput) ToAudioResponseOutputWithContext(ctx context.Context) AudioResponseOutput {
	return o
}

func (o AudioResponseOutput) Bitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AudioResponse) *int { return v.Bitrate }).(pulumi.IntPtrOutput)
}

func (o AudioResponseOutput) Channels() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AudioResponse) *int { return v.Channels }).(pulumi.IntPtrOutput)
}

func (o AudioResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o AudioResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AudioResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o AudioResponseOutput) SamplingRate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AudioResponse) *int { return v.SamplingRate }).(pulumi.IntPtrOutput)
}

type AudioTrackDescriptor struct {
	ChannelMapping *string `pulumi:"channelMapping"`
	OdataType      string  `pulumi:"odataType"`
}





type AudioTrackDescriptorInput interface {
	pulumi.Input

	ToAudioTrackDescriptorOutput() AudioTrackDescriptorOutput
	ToAudioTrackDescriptorOutputWithContext(context.Context) AudioTrackDescriptorOutput
}

type AudioTrackDescriptorArgs struct {
	ChannelMapping pulumi.StringPtrInput `pulumi:"channelMapping"`
	OdataType      pulumi.StringInput    `pulumi:"odataType"`
}

func (AudioTrackDescriptorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioTrackDescriptor)(nil)).Elem()
}

func (i AudioTrackDescriptorArgs) ToAudioTrackDescriptorOutput() AudioTrackDescriptorOutput {
	return i.ToAudioTrackDescriptorOutputWithContext(context.Background())
}

func (i AudioTrackDescriptorArgs) ToAudioTrackDescriptorOutputWithContext(ctx context.Context) AudioTrackDescriptorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioTrackDescriptorOutput)
}

type AudioTrackDescriptorOutput struct{ *pulumi.OutputState }

func (AudioTrackDescriptorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioTrackDescriptor)(nil)).Elem()
}

func (o AudioTrackDescriptorOutput) ToAudioTrackDescriptorOutput() AudioTrackDescriptorOutput {
	return o
}

func (o AudioTrackDescriptorOutput) ToAudioTrackDescriptorOutputWithContext(ctx context.Context) AudioTrackDescriptorOutput {
	return o
}

func (o AudioTrackDescriptorOutput) ChannelMapping() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioTrackDescriptor) *string { return v.ChannelMapping }).(pulumi.StringPtrOutput)
}

func (o AudioTrackDescriptorOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AudioTrackDescriptor) string { return v.OdataType }).(pulumi.StringOutput)
}

type AudioTrackDescriptorResponse struct {
	ChannelMapping *string `pulumi:"channelMapping"`
	OdataType      string  `pulumi:"odataType"`
}





type AudioTrackDescriptorResponseInput interface {
	pulumi.Input

	ToAudioTrackDescriptorResponseOutput() AudioTrackDescriptorResponseOutput
	ToAudioTrackDescriptorResponseOutputWithContext(context.Context) AudioTrackDescriptorResponseOutput
}

type AudioTrackDescriptorResponseArgs struct {
	ChannelMapping pulumi.StringPtrInput `pulumi:"channelMapping"`
	OdataType      pulumi.StringInput    `pulumi:"odataType"`
}

func (AudioTrackDescriptorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioTrackDescriptorResponse)(nil)).Elem()
}

func (i AudioTrackDescriptorResponseArgs) ToAudioTrackDescriptorResponseOutput() AudioTrackDescriptorResponseOutput {
	return i.ToAudioTrackDescriptorResponseOutputWithContext(context.Background())
}

func (i AudioTrackDescriptorResponseArgs) ToAudioTrackDescriptorResponseOutputWithContext(ctx context.Context) AudioTrackDescriptorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudioTrackDescriptorResponseOutput)
}

type AudioTrackDescriptorResponseOutput struct{ *pulumi.OutputState }

func (AudioTrackDescriptorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AudioTrackDescriptorResponse)(nil)).Elem()
}

func (o AudioTrackDescriptorResponseOutput) ToAudioTrackDescriptorResponseOutput() AudioTrackDescriptorResponseOutput {
	return o
}

func (o AudioTrackDescriptorResponseOutput) ToAudioTrackDescriptorResponseOutputWithContext(ctx context.Context) AudioTrackDescriptorResponseOutput {
	return o
}

func (o AudioTrackDescriptorResponseOutput) ChannelMapping() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AudioTrackDescriptorResponse) *string { return v.ChannelMapping }).(pulumi.StringPtrOutput)
}

func (o AudioTrackDescriptorResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AudioTrackDescriptorResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type BuiltInStandardEncoderPreset struct {
	OdataType  string `pulumi:"odataType"`
	PresetName string `pulumi:"presetName"`
}





type BuiltInStandardEncoderPresetInput interface {
	pulumi.Input

	ToBuiltInStandardEncoderPresetOutput() BuiltInStandardEncoderPresetOutput
	ToBuiltInStandardEncoderPresetOutputWithContext(context.Context) BuiltInStandardEncoderPresetOutput
}

type BuiltInStandardEncoderPresetArgs struct {
	OdataType  pulumi.StringInput `pulumi:"odataType"`
	PresetName pulumi.StringInput `pulumi:"presetName"`
}

func (BuiltInStandardEncoderPresetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltInStandardEncoderPreset)(nil)).Elem()
}

func (i BuiltInStandardEncoderPresetArgs) ToBuiltInStandardEncoderPresetOutput() BuiltInStandardEncoderPresetOutput {
	return i.ToBuiltInStandardEncoderPresetOutputWithContext(context.Background())
}

func (i BuiltInStandardEncoderPresetArgs) ToBuiltInStandardEncoderPresetOutputWithContext(ctx context.Context) BuiltInStandardEncoderPresetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuiltInStandardEncoderPresetOutput)
}

type BuiltInStandardEncoderPresetOutput struct{ *pulumi.OutputState }

func (BuiltInStandardEncoderPresetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltInStandardEncoderPreset)(nil)).Elem()
}

func (o BuiltInStandardEncoderPresetOutput) ToBuiltInStandardEncoderPresetOutput() BuiltInStandardEncoderPresetOutput {
	return o
}

func (o BuiltInStandardEncoderPresetOutput) ToBuiltInStandardEncoderPresetOutputWithContext(ctx context.Context) BuiltInStandardEncoderPresetOutput {
	return o
}

func (o BuiltInStandardEncoderPresetOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v BuiltInStandardEncoderPreset) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o BuiltInStandardEncoderPresetOutput) PresetName() pulumi.StringOutput {
	return o.ApplyT(func(v BuiltInStandardEncoderPreset) string { return v.PresetName }).(pulumi.StringOutput)
}

type BuiltInStandardEncoderPresetResponse struct {
	OdataType  string `pulumi:"odataType"`
	PresetName string `pulumi:"presetName"`
}





type BuiltInStandardEncoderPresetResponseInput interface {
	pulumi.Input

	ToBuiltInStandardEncoderPresetResponseOutput() BuiltInStandardEncoderPresetResponseOutput
	ToBuiltInStandardEncoderPresetResponseOutputWithContext(context.Context) BuiltInStandardEncoderPresetResponseOutput
}

type BuiltInStandardEncoderPresetResponseArgs struct {
	OdataType  pulumi.StringInput `pulumi:"odataType"`
	PresetName pulumi.StringInput `pulumi:"presetName"`
}

func (BuiltInStandardEncoderPresetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltInStandardEncoderPresetResponse)(nil)).Elem()
}

func (i BuiltInStandardEncoderPresetResponseArgs) ToBuiltInStandardEncoderPresetResponseOutput() BuiltInStandardEncoderPresetResponseOutput {
	return i.ToBuiltInStandardEncoderPresetResponseOutputWithContext(context.Background())
}

func (i BuiltInStandardEncoderPresetResponseArgs) ToBuiltInStandardEncoderPresetResponseOutputWithContext(ctx context.Context) BuiltInStandardEncoderPresetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BuiltInStandardEncoderPresetResponseOutput)
}

type BuiltInStandardEncoderPresetResponseOutput struct{ *pulumi.OutputState }

func (BuiltInStandardEncoderPresetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltInStandardEncoderPresetResponse)(nil)).Elem()
}

func (o BuiltInStandardEncoderPresetResponseOutput) ToBuiltInStandardEncoderPresetResponseOutput() BuiltInStandardEncoderPresetResponseOutput {
	return o
}

func (o BuiltInStandardEncoderPresetResponseOutput) ToBuiltInStandardEncoderPresetResponseOutputWithContext(ctx context.Context) BuiltInStandardEncoderPresetResponseOutput {
	return o
}

func (o BuiltInStandardEncoderPresetResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v BuiltInStandardEncoderPresetResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o BuiltInStandardEncoderPresetResponseOutput) PresetName() pulumi.StringOutput {
	return o.ApplyT(func(v BuiltInStandardEncoderPresetResponse) string { return v.PresetName }).(pulumi.StringOutput)
}

type CbcsDrmConfiguration struct {
	FairPlay  *StreamingPolicyFairPlayConfiguration  `pulumi:"fairPlay"`
	PlayReady *StreamingPolicyPlayReadyConfiguration `pulumi:"playReady"`
	Widevine  *StreamingPolicyWidevineConfiguration  `pulumi:"widevine"`
}





type CbcsDrmConfigurationInput interface {
	pulumi.Input

	ToCbcsDrmConfigurationOutput() CbcsDrmConfigurationOutput
	ToCbcsDrmConfigurationOutputWithContext(context.Context) CbcsDrmConfigurationOutput
}

type CbcsDrmConfigurationArgs struct {
	FairPlay  StreamingPolicyFairPlayConfigurationPtrInput  `pulumi:"fairPlay"`
	PlayReady StreamingPolicyPlayReadyConfigurationPtrInput `pulumi:"playReady"`
	Widevine  StreamingPolicyWidevineConfigurationPtrInput  `pulumi:"widevine"`
}

func (CbcsDrmConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CbcsDrmConfiguration)(nil)).Elem()
}

func (i CbcsDrmConfigurationArgs) ToCbcsDrmConfigurationOutput() CbcsDrmConfigurationOutput {
	return i.ToCbcsDrmConfigurationOutputWithContext(context.Background())
}

func (i CbcsDrmConfigurationArgs) ToCbcsDrmConfigurationOutputWithContext(ctx context.Context) CbcsDrmConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CbcsDrmConfigurationOutput)
}

func (i CbcsDrmConfigurationArgs) ToCbcsDrmConfigurationPtrOutput() CbcsDrmConfigurationPtrOutput {
	return i.ToCbcsDrmConfigurationPtrOutputWithContext(context.Background())
}

func (i CbcsDrmConfigurationArgs) ToCbcsDrmConfigurationPtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CbcsDrmConfigurationOutput).ToCbcsDrmConfigurationPtrOutputWithContext(ctx)
}









type CbcsDrmConfigurationPtrInput interface {
	pulumi.Input

	ToCbcsDrmConfigurationPtrOutput() CbcsDrmConfigurationPtrOutput
	ToCbcsDrmConfigurationPtrOutputWithContext(context.Context) CbcsDrmConfigurationPtrOutput
}

type cbcsDrmConfigurationPtrType CbcsDrmConfigurationArgs

func CbcsDrmConfigurationPtr(v *CbcsDrmConfigurationArgs) CbcsDrmConfigurationPtrInput {
	return (*cbcsDrmConfigurationPtrType)(v)
}

func (*cbcsDrmConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CbcsDrmConfiguration)(nil)).Elem()
}

func (i *cbcsDrmConfigurationPtrType) ToCbcsDrmConfigurationPtrOutput() CbcsDrmConfigurationPtrOutput {
	return i.ToCbcsDrmConfigurationPtrOutputWithContext(context.Background())
}

func (i *cbcsDrmConfigurationPtrType) ToCbcsDrmConfigurationPtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CbcsDrmConfigurationPtrOutput)
}

type CbcsDrmConfigurationOutput struct{ *pulumi.OutputState }

func (CbcsDrmConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CbcsDrmConfiguration)(nil)).Elem()
}

func (o CbcsDrmConfigurationOutput) ToCbcsDrmConfigurationOutput() CbcsDrmConfigurationOutput {
	return o
}

func (o CbcsDrmConfigurationOutput) ToCbcsDrmConfigurationOutputWithContext(ctx context.Context) CbcsDrmConfigurationOutput {
	return o
}

func (o CbcsDrmConfigurationOutput) ToCbcsDrmConfigurationPtrOutput() CbcsDrmConfigurationPtrOutput {
	return o.ToCbcsDrmConfigurationPtrOutputWithContext(context.Background())
}

func (o CbcsDrmConfigurationOutput) ToCbcsDrmConfigurationPtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CbcsDrmConfiguration) *CbcsDrmConfiguration {
		return &v
	}).(CbcsDrmConfigurationPtrOutput)
}

func (o CbcsDrmConfigurationOutput) FairPlay() StreamingPolicyFairPlayConfigurationPtrOutput {
	return o.ApplyT(func(v CbcsDrmConfiguration) *StreamingPolicyFairPlayConfiguration { return v.FairPlay }).(StreamingPolicyFairPlayConfigurationPtrOutput)
}

func (o CbcsDrmConfigurationOutput) PlayReady() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ApplyT(func(v CbcsDrmConfiguration) *StreamingPolicyPlayReadyConfiguration { return v.PlayReady }).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

func (o CbcsDrmConfigurationOutput) Widevine() StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ApplyT(func(v CbcsDrmConfiguration) *StreamingPolicyWidevineConfiguration { return v.Widevine }).(StreamingPolicyWidevineConfigurationPtrOutput)
}

type CbcsDrmConfigurationPtrOutput struct{ *pulumi.OutputState }

func (CbcsDrmConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CbcsDrmConfiguration)(nil)).Elem()
}

func (o CbcsDrmConfigurationPtrOutput) ToCbcsDrmConfigurationPtrOutput() CbcsDrmConfigurationPtrOutput {
	return o
}

func (o CbcsDrmConfigurationPtrOutput) ToCbcsDrmConfigurationPtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationPtrOutput {
	return o
}

func (o CbcsDrmConfigurationPtrOutput) Elem() CbcsDrmConfigurationOutput {
	return o.ApplyT(func(v *CbcsDrmConfiguration) CbcsDrmConfiguration {
		if v != nil {
			return *v
		}
		var ret CbcsDrmConfiguration
		return ret
	}).(CbcsDrmConfigurationOutput)
}

func (o CbcsDrmConfigurationPtrOutput) FairPlay() StreamingPolicyFairPlayConfigurationPtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfiguration) *StreamingPolicyFairPlayConfiguration {
		if v == nil {
			return nil
		}
		return v.FairPlay
	}).(StreamingPolicyFairPlayConfigurationPtrOutput)
}

func (o CbcsDrmConfigurationPtrOutput) PlayReady() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfiguration) *StreamingPolicyPlayReadyConfiguration {
		if v == nil {
			return nil
		}
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

func (o CbcsDrmConfigurationPtrOutput) Widevine() StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfiguration) *StreamingPolicyWidevineConfiguration {
		if v == nil {
			return nil
		}
		return v.Widevine
	}).(StreamingPolicyWidevineConfigurationPtrOutput)
}

type CbcsDrmConfigurationResponse struct {
	FairPlay  *StreamingPolicyFairPlayConfigurationResponse  `pulumi:"fairPlay"`
	PlayReady *StreamingPolicyPlayReadyConfigurationResponse `pulumi:"playReady"`
	Widevine  *StreamingPolicyWidevineConfigurationResponse  `pulumi:"widevine"`
}





type CbcsDrmConfigurationResponseInput interface {
	pulumi.Input

	ToCbcsDrmConfigurationResponseOutput() CbcsDrmConfigurationResponseOutput
	ToCbcsDrmConfigurationResponseOutputWithContext(context.Context) CbcsDrmConfigurationResponseOutput
}

type CbcsDrmConfigurationResponseArgs struct {
	FairPlay  StreamingPolicyFairPlayConfigurationResponsePtrInput  `pulumi:"fairPlay"`
	PlayReady StreamingPolicyPlayReadyConfigurationResponsePtrInput `pulumi:"playReady"`
	Widevine  StreamingPolicyWidevineConfigurationResponsePtrInput  `pulumi:"widevine"`
}

func (CbcsDrmConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CbcsDrmConfigurationResponse)(nil)).Elem()
}

func (i CbcsDrmConfigurationResponseArgs) ToCbcsDrmConfigurationResponseOutput() CbcsDrmConfigurationResponseOutput {
	return i.ToCbcsDrmConfigurationResponseOutputWithContext(context.Background())
}

func (i CbcsDrmConfigurationResponseArgs) ToCbcsDrmConfigurationResponseOutputWithContext(ctx context.Context) CbcsDrmConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CbcsDrmConfigurationResponseOutput)
}

func (i CbcsDrmConfigurationResponseArgs) ToCbcsDrmConfigurationResponsePtrOutput() CbcsDrmConfigurationResponsePtrOutput {
	return i.ToCbcsDrmConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i CbcsDrmConfigurationResponseArgs) ToCbcsDrmConfigurationResponsePtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CbcsDrmConfigurationResponseOutput).ToCbcsDrmConfigurationResponsePtrOutputWithContext(ctx)
}









type CbcsDrmConfigurationResponsePtrInput interface {
	pulumi.Input

	ToCbcsDrmConfigurationResponsePtrOutput() CbcsDrmConfigurationResponsePtrOutput
	ToCbcsDrmConfigurationResponsePtrOutputWithContext(context.Context) CbcsDrmConfigurationResponsePtrOutput
}

type cbcsDrmConfigurationResponsePtrType CbcsDrmConfigurationResponseArgs

func CbcsDrmConfigurationResponsePtr(v *CbcsDrmConfigurationResponseArgs) CbcsDrmConfigurationResponsePtrInput {
	return (*cbcsDrmConfigurationResponsePtrType)(v)
}

func (*cbcsDrmConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CbcsDrmConfigurationResponse)(nil)).Elem()
}

func (i *cbcsDrmConfigurationResponsePtrType) ToCbcsDrmConfigurationResponsePtrOutput() CbcsDrmConfigurationResponsePtrOutput {
	return i.ToCbcsDrmConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *cbcsDrmConfigurationResponsePtrType) ToCbcsDrmConfigurationResponsePtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CbcsDrmConfigurationResponsePtrOutput)
}

type CbcsDrmConfigurationResponseOutput struct{ *pulumi.OutputState }

func (CbcsDrmConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CbcsDrmConfigurationResponse)(nil)).Elem()
}

func (o CbcsDrmConfigurationResponseOutput) ToCbcsDrmConfigurationResponseOutput() CbcsDrmConfigurationResponseOutput {
	return o
}

func (o CbcsDrmConfigurationResponseOutput) ToCbcsDrmConfigurationResponseOutputWithContext(ctx context.Context) CbcsDrmConfigurationResponseOutput {
	return o
}

func (o CbcsDrmConfigurationResponseOutput) ToCbcsDrmConfigurationResponsePtrOutput() CbcsDrmConfigurationResponsePtrOutput {
	return o.ToCbcsDrmConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o CbcsDrmConfigurationResponseOutput) ToCbcsDrmConfigurationResponsePtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CbcsDrmConfigurationResponse) *CbcsDrmConfigurationResponse {
		return &v
	}).(CbcsDrmConfigurationResponsePtrOutput)
}

func (o CbcsDrmConfigurationResponseOutput) FairPlay() StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CbcsDrmConfigurationResponse) *StreamingPolicyFairPlayConfigurationResponse { return v.FairPlay }).(StreamingPolicyFairPlayConfigurationResponsePtrOutput)
}

func (o CbcsDrmConfigurationResponseOutput) PlayReady() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CbcsDrmConfigurationResponse) *StreamingPolicyPlayReadyConfigurationResponse {
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationResponsePtrOutput)
}

func (o CbcsDrmConfigurationResponseOutput) Widevine() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CbcsDrmConfigurationResponse) *StreamingPolicyWidevineConfigurationResponse { return v.Widevine }).(StreamingPolicyWidevineConfigurationResponsePtrOutput)
}

type CbcsDrmConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (CbcsDrmConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CbcsDrmConfigurationResponse)(nil)).Elem()
}

func (o CbcsDrmConfigurationResponsePtrOutput) ToCbcsDrmConfigurationResponsePtrOutput() CbcsDrmConfigurationResponsePtrOutput {
	return o
}

func (o CbcsDrmConfigurationResponsePtrOutput) ToCbcsDrmConfigurationResponsePtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationResponsePtrOutput {
	return o
}

func (o CbcsDrmConfigurationResponsePtrOutput) Elem() CbcsDrmConfigurationResponseOutput {
	return o.ApplyT(func(v *CbcsDrmConfigurationResponse) CbcsDrmConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret CbcsDrmConfigurationResponse
		return ret
	}).(CbcsDrmConfigurationResponseOutput)
}

func (o CbcsDrmConfigurationResponsePtrOutput) FairPlay() StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfigurationResponse) *StreamingPolicyFairPlayConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.FairPlay
	}).(StreamingPolicyFairPlayConfigurationResponsePtrOutput)
}

func (o CbcsDrmConfigurationResponsePtrOutput) PlayReady() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfigurationResponse) *StreamingPolicyPlayReadyConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationResponsePtrOutput)
}

func (o CbcsDrmConfigurationResponsePtrOutput) Widevine() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfigurationResponse) *StreamingPolicyWidevineConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Widevine
	}).(StreamingPolicyWidevineConfigurationResponsePtrOutput)
}

type CencDrmConfiguration struct {
	PlayReady *StreamingPolicyPlayReadyConfiguration `pulumi:"playReady"`
	Widevine  *StreamingPolicyWidevineConfiguration  `pulumi:"widevine"`
}





type CencDrmConfigurationInput interface {
	pulumi.Input

	ToCencDrmConfigurationOutput() CencDrmConfigurationOutput
	ToCencDrmConfigurationOutputWithContext(context.Context) CencDrmConfigurationOutput
}

type CencDrmConfigurationArgs struct {
	PlayReady StreamingPolicyPlayReadyConfigurationPtrInput `pulumi:"playReady"`
	Widevine  StreamingPolicyWidevineConfigurationPtrInput  `pulumi:"widevine"`
}

func (CencDrmConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CencDrmConfiguration)(nil)).Elem()
}

func (i CencDrmConfigurationArgs) ToCencDrmConfigurationOutput() CencDrmConfigurationOutput {
	return i.ToCencDrmConfigurationOutputWithContext(context.Background())
}

func (i CencDrmConfigurationArgs) ToCencDrmConfigurationOutputWithContext(ctx context.Context) CencDrmConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CencDrmConfigurationOutput)
}

func (i CencDrmConfigurationArgs) ToCencDrmConfigurationPtrOutput() CencDrmConfigurationPtrOutput {
	return i.ToCencDrmConfigurationPtrOutputWithContext(context.Background())
}

func (i CencDrmConfigurationArgs) ToCencDrmConfigurationPtrOutputWithContext(ctx context.Context) CencDrmConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CencDrmConfigurationOutput).ToCencDrmConfigurationPtrOutputWithContext(ctx)
}









type CencDrmConfigurationPtrInput interface {
	pulumi.Input

	ToCencDrmConfigurationPtrOutput() CencDrmConfigurationPtrOutput
	ToCencDrmConfigurationPtrOutputWithContext(context.Context) CencDrmConfigurationPtrOutput
}

type cencDrmConfigurationPtrType CencDrmConfigurationArgs

func CencDrmConfigurationPtr(v *CencDrmConfigurationArgs) CencDrmConfigurationPtrInput {
	return (*cencDrmConfigurationPtrType)(v)
}

func (*cencDrmConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CencDrmConfiguration)(nil)).Elem()
}

func (i *cencDrmConfigurationPtrType) ToCencDrmConfigurationPtrOutput() CencDrmConfigurationPtrOutput {
	return i.ToCencDrmConfigurationPtrOutputWithContext(context.Background())
}

func (i *cencDrmConfigurationPtrType) ToCencDrmConfigurationPtrOutputWithContext(ctx context.Context) CencDrmConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CencDrmConfigurationPtrOutput)
}

type CencDrmConfigurationOutput struct{ *pulumi.OutputState }

func (CencDrmConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CencDrmConfiguration)(nil)).Elem()
}

func (o CencDrmConfigurationOutput) ToCencDrmConfigurationOutput() CencDrmConfigurationOutput {
	return o
}

func (o CencDrmConfigurationOutput) ToCencDrmConfigurationOutputWithContext(ctx context.Context) CencDrmConfigurationOutput {
	return o
}

func (o CencDrmConfigurationOutput) ToCencDrmConfigurationPtrOutput() CencDrmConfigurationPtrOutput {
	return o.ToCencDrmConfigurationPtrOutputWithContext(context.Background())
}

func (o CencDrmConfigurationOutput) ToCencDrmConfigurationPtrOutputWithContext(ctx context.Context) CencDrmConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CencDrmConfiguration) *CencDrmConfiguration {
		return &v
	}).(CencDrmConfigurationPtrOutput)
}

func (o CencDrmConfigurationOutput) PlayReady() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ApplyT(func(v CencDrmConfiguration) *StreamingPolicyPlayReadyConfiguration { return v.PlayReady }).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

func (o CencDrmConfigurationOutput) Widevine() StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ApplyT(func(v CencDrmConfiguration) *StreamingPolicyWidevineConfiguration { return v.Widevine }).(StreamingPolicyWidevineConfigurationPtrOutput)
}

type CencDrmConfigurationPtrOutput struct{ *pulumi.OutputState }

func (CencDrmConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CencDrmConfiguration)(nil)).Elem()
}

func (o CencDrmConfigurationPtrOutput) ToCencDrmConfigurationPtrOutput() CencDrmConfigurationPtrOutput {
	return o
}

func (o CencDrmConfigurationPtrOutput) ToCencDrmConfigurationPtrOutputWithContext(ctx context.Context) CencDrmConfigurationPtrOutput {
	return o
}

func (o CencDrmConfigurationPtrOutput) Elem() CencDrmConfigurationOutput {
	return o.ApplyT(func(v *CencDrmConfiguration) CencDrmConfiguration {
		if v != nil {
			return *v
		}
		var ret CencDrmConfiguration
		return ret
	}).(CencDrmConfigurationOutput)
}

func (o CencDrmConfigurationPtrOutput) PlayReady() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ApplyT(func(v *CencDrmConfiguration) *StreamingPolicyPlayReadyConfiguration {
		if v == nil {
			return nil
		}
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

func (o CencDrmConfigurationPtrOutput) Widevine() StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ApplyT(func(v *CencDrmConfiguration) *StreamingPolicyWidevineConfiguration {
		if v == nil {
			return nil
		}
		return v.Widevine
	}).(StreamingPolicyWidevineConfigurationPtrOutput)
}

type CencDrmConfigurationResponse struct {
	PlayReady *StreamingPolicyPlayReadyConfigurationResponse `pulumi:"playReady"`
	Widevine  *StreamingPolicyWidevineConfigurationResponse  `pulumi:"widevine"`
}





type CencDrmConfigurationResponseInput interface {
	pulumi.Input

	ToCencDrmConfigurationResponseOutput() CencDrmConfigurationResponseOutput
	ToCencDrmConfigurationResponseOutputWithContext(context.Context) CencDrmConfigurationResponseOutput
}

type CencDrmConfigurationResponseArgs struct {
	PlayReady StreamingPolicyPlayReadyConfigurationResponsePtrInput `pulumi:"playReady"`
	Widevine  StreamingPolicyWidevineConfigurationResponsePtrInput  `pulumi:"widevine"`
}

func (CencDrmConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CencDrmConfigurationResponse)(nil)).Elem()
}

func (i CencDrmConfigurationResponseArgs) ToCencDrmConfigurationResponseOutput() CencDrmConfigurationResponseOutput {
	return i.ToCencDrmConfigurationResponseOutputWithContext(context.Background())
}

func (i CencDrmConfigurationResponseArgs) ToCencDrmConfigurationResponseOutputWithContext(ctx context.Context) CencDrmConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CencDrmConfigurationResponseOutput)
}

func (i CencDrmConfigurationResponseArgs) ToCencDrmConfigurationResponsePtrOutput() CencDrmConfigurationResponsePtrOutput {
	return i.ToCencDrmConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i CencDrmConfigurationResponseArgs) ToCencDrmConfigurationResponsePtrOutputWithContext(ctx context.Context) CencDrmConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CencDrmConfigurationResponseOutput).ToCencDrmConfigurationResponsePtrOutputWithContext(ctx)
}









type CencDrmConfigurationResponsePtrInput interface {
	pulumi.Input

	ToCencDrmConfigurationResponsePtrOutput() CencDrmConfigurationResponsePtrOutput
	ToCencDrmConfigurationResponsePtrOutputWithContext(context.Context) CencDrmConfigurationResponsePtrOutput
}

type cencDrmConfigurationResponsePtrType CencDrmConfigurationResponseArgs

func CencDrmConfigurationResponsePtr(v *CencDrmConfigurationResponseArgs) CencDrmConfigurationResponsePtrInput {
	return (*cencDrmConfigurationResponsePtrType)(v)
}

func (*cencDrmConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CencDrmConfigurationResponse)(nil)).Elem()
}

func (i *cencDrmConfigurationResponsePtrType) ToCencDrmConfigurationResponsePtrOutput() CencDrmConfigurationResponsePtrOutput {
	return i.ToCencDrmConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *cencDrmConfigurationResponsePtrType) ToCencDrmConfigurationResponsePtrOutputWithContext(ctx context.Context) CencDrmConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CencDrmConfigurationResponsePtrOutput)
}

type CencDrmConfigurationResponseOutput struct{ *pulumi.OutputState }

func (CencDrmConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CencDrmConfigurationResponse)(nil)).Elem()
}

func (o CencDrmConfigurationResponseOutput) ToCencDrmConfigurationResponseOutput() CencDrmConfigurationResponseOutput {
	return o
}

func (o CencDrmConfigurationResponseOutput) ToCencDrmConfigurationResponseOutputWithContext(ctx context.Context) CencDrmConfigurationResponseOutput {
	return o
}

func (o CencDrmConfigurationResponseOutput) ToCencDrmConfigurationResponsePtrOutput() CencDrmConfigurationResponsePtrOutput {
	return o.ToCencDrmConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o CencDrmConfigurationResponseOutput) ToCencDrmConfigurationResponsePtrOutputWithContext(ctx context.Context) CencDrmConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CencDrmConfigurationResponse) *CencDrmConfigurationResponse {
		return &v
	}).(CencDrmConfigurationResponsePtrOutput)
}

func (o CencDrmConfigurationResponseOutput) PlayReady() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CencDrmConfigurationResponse) *StreamingPolicyPlayReadyConfigurationResponse {
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationResponsePtrOutput)
}

func (o CencDrmConfigurationResponseOutput) Widevine() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CencDrmConfigurationResponse) *StreamingPolicyWidevineConfigurationResponse { return v.Widevine }).(StreamingPolicyWidevineConfigurationResponsePtrOutput)
}

type CencDrmConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (CencDrmConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CencDrmConfigurationResponse)(nil)).Elem()
}

func (o CencDrmConfigurationResponsePtrOutput) ToCencDrmConfigurationResponsePtrOutput() CencDrmConfigurationResponsePtrOutput {
	return o
}

func (o CencDrmConfigurationResponsePtrOutput) ToCencDrmConfigurationResponsePtrOutputWithContext(ctx context.Context) CencDrmConfigurationResponsePtrOutput {
	return o
}

func (o CencDrmConfigurationResponsePtrOutput) Elem() CencDrmConfigurationResponseOutput {
	return o.ApplyT(func(v *CencDrmConfigurationResponse) CencDrmConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret CencDrmConfigurationResponse
		return ret
	}).(CencDrmConfigurationResponseOutput)
}

func (o CencDrmConfigurationResponsePtrOutput) PlayReady() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CencDrmConfigurationResponse) *StreamingPolicyPlayReadyConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationResponsePtrOutput)
}

func (o CencDrmConfigurationResponsePtrOutput) Widevine() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CencDrmConfigurationResponse) *StreamingPolicyWidevineConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Widevine
	}).(StreamingPolicyWidevineConfigurationResponsePtrOutput)
}

type CommonEncryptionCbcs struct {
	ClearTracks      []TrackSelection            `pulumi:"clearTracks"`
	ContentKeys      *StreamingPolicyContentKeys `pulumi:"contentKeys"`
	Drm              *CbcsDrmConfiguration       `pulumi:"drm"`
	EnabledProtocols *EnabledProtocols           `pulumi:"enabledProtocols"`
}





type CommonEncryptionCbcsInput interface {
	pulumi.Input

	ToCommonEncryptionCbcsOutput() CommonEncryptionCbcsOutput
	ToCommonEncryptionCbcsOutputWithContext(context.Context) CommonEncryptionCbcsOutput
}

type CommonEncryptionCbcsArgs struct {
	ClearTracks      TrackSelectionArrayInput           `pulumi:"clearTracks"`
	ContentKeys      StreamingPolicyContentKeysPtrInput `pulumi:"contentKeys"`
	Drm              CbcsDrmConfigurationPtrInput       `pulumi:"drm"`
	EnabledProtocols EnabledProtocolsPtrInput           `pulumi:"enabledProtocols"`
}

func (CommonEncryptionCbcsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCbcs)(nil)).Elem()
}

func (i CommonEncryptionCbcsArgs) ToCommonEncryptionCbcsOutput() CommonEncryptionCbcsOutput {
	return i.ToCommonEncryptionCbcsOutputWithContext(context.Background())
}

func (i CommonEncryptionCbcsArgs) ToCommonEncryptionCbcsOutputWithContext(ctx context.Context) CommonEncryptionCbcsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCbcsOutput)
}

func (i CommonEncryptionCbcsArgs) ToCommonEncryptionCbcsPtrOutput() CommonEncryptionCbcsPtrOutput {
	return i.ToCommonEncryptionCbcsPtrOutputWithContext(context.Background())
}

func (i CommonEncryptionCbcsArgs) ToCommonEncryptionCbcsPtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCbcsOutput).ToCommonEncryptionCbcsPtrOutputWithContext(ctx)
}









type CommonEncryptionCbcsPtrInput interface {
	pulumi.Input

	ToCommonEncryptionCbcsPtrOutput() CommonEncryptionCbcsPtrOutput
	ToCommonEncryptionCbcsPtrOutputWithContext(context.Context) CommonEncryptionCbcsPtrOutput
}

type commonEncryptionCbcsPtrType CommonEncryptionCbcsArgs

func CommonEncryptionCbcsPtr(v *CommonEncryptionCbcsArgs) CommonEncryptionCbcsPtrInput {
	return (*commonEncryptionCbcsPtrType)(v)
}

func (*commonEncryptionCbcsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCbcs)(nil)).Elem()
}

func (i *commonEncryptionCbcsPtrType) ToCommonEncryptionCbcsPtrOutput() CommonEncryptionCbcsPtrOutput {
	return i.ToCommonEncryptionCbcsPtrOutputWithContext(context.Background())
}

func (i *commonEncryptionCbcsPtrType) ToCommonEncryptionCbcsPtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCbcsPtrOutput)
}

type CommonEncryptionCbcsOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCbcsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCbcs)(nil)).Elem()
}

func (o CommonEncryptionCbcsOutput) ToCommonEncryptionCbcsOutput() CommonEncryptionCbcsOutput {
	return o
}

func (o CommonEncryptionCbcsOutput) ToCommonEncryptionCbcsOutputWithContext(ctx context.Context) CommonEncryptionCbcsOutput {
	return o
}

func (o CommonEncryptionCbcsOutput) ToCommonEncryptionCbcsPtrOutput() CommonEncryptionCbcsPtrOutput {
	return o.ToCommonEncryptionCbcsPtrOutputWithContext(context.Background())
}

func (o CommonEncryptionCbcsOutput) ToCommonEncryptionCbcsPtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommonEncryptionCbcs) *CommonEncryptionCbcs {
		return &v
	}).(CommonEncryptionCbcsPtrOutput)
}

func (o CommonEncryptionCbcsOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v CommonEncryptionCbcs) []TrackSelection { return v.ClearTracks }).(TrackSelectionArrayOutput)
}

func (o CommonEncryptionCbcsOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcs) *StreamingPolicyContentKeys { return v.ContentKeys }).(StreamingPolicyContentKeysPtrOutput)
}

func (o CommonEncryptionCbcsOutput) Drm() CbcsDrmConfigurationPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcs) *CbcsDrmConfiguration { return v.Drm }).(CbcsDrmConfigurationPtrOutput)
}

func (o CommonEncryptionCbcsOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcs) *EnabledProtocols { return v.EnabledProtocols }).(EnabledProtocolsPtrOutput)
}

type CommonEncryptionCbcsPtrOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCbcsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCbcs)(nil)).Elem()
}

func (o CommonEncryptionCbcsPtrOutput) ToCommonEncryptionCbcsPtrOutput() CommonEncryptionCbcsPtrOutput {
	return o
}

func (o CommonEncryptionCbcsPtrOutput) ToCommonEncryptionCbcsPtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsPtrOutput {
	return o
}

func (o CommonEncryptionCbcsPtrOutput) Elem() CommonEncryptionCbcsOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcs) CommonEncryptionCbcs {
		if v != nil {
			return *v
		}
		var ret CommonEncryptionCbcs
		return ret
	}).(CommonEncryptionCbcsOutput)
}

func (o CommonEncryptionCbcsPtrOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcs) []TrackSelection {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionArrayOutput)
}

func (o CommonEncryptionCbcsPtrOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcs) *StreamingPolicyContentKeys {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysPtrOutput)
}

func (o CommonEncryptionCbcsPtrOutput) Drm() CbcsDrmConfigurationPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcs) *CbcsDrmConfiguration {
		if v == nil {
			return nil
		}
		return v.Drm
	}).(CbcsDrmConfigurationPtrOutput)
}

func (o CommonEncryptionCbcsPtrOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcs) *EnabledProtocols {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsPtrOutput)
}

type CommonEncryptionCbcsResponse struct {
	ClearTracks      []TrackSelectionResponse            `pulumi:"clearTracks"`
	ContentKeys      *StreamingPolicyContentKeysResponse `pulumi:"contentKeys"`
	Drm              *CbcsDrmConfigurationResponse       `pulumi:"drm"`
	EnabledProtocols *EnabledProtocolsResponse           `pulumi:"enabledProtocols"`
}





type CommonEncryptionCbcsResponseInput interface {
	pulumi.Input

	ToCommonEncryptionCbcsResponseOutput() CommonEncryptionCbcsResponseOutput
	ToCommonEncryptionCbcsResponseOutputWithContext(context.Context) CommonEncryptionCbcsResponseOutput
}

type CommonEncryptionCbcsResponseArgs struct {
	ClearTracks      TrackSelectionResponseArrayInput           `pulumi:"clearTracks"`
	ContentKeys      StreamingPolicyContentKeysResponsePtrInput `pulumi:"contentKeys"`
	Drm              CbcsDrmConfigurationResponsePtrInput       `pulumi:"drm"`
	EnabledProtocols EnabledProtocolsResponsePtrInput           `pulumi:"enabledProtocols"`
}

func (CommonEncryptionCbcsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCbcsResponse)(nil)).Elem()
}

func (i CommonEncryptionCbcsResponseArgs) ToCommonEncryptionCbcsResponseOutput() CommonEncryptionCbcsResponseOutput {
	return i.ToCommonEncryptionCbcsResponseOutputWithContext(context.Background())
}

func (i CommonEncryptionCbcsResponseArgs) ToCommonEncryptionCbcsResponseOutputWithContext(ctx context.Context) CommonEncryptionCbcsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCbcsResponseOutput)
}

func (i CommonEncryptionCbcsResponseArgs) ToCommonEncryptionCbcsResponsePtrOutput() CommonEncryptionCbcsResponsePtrOutput {
	return i.ToCommonEncryptionCbcsResponsePtrOutputWithContext(context.Background())
}

func (i CommonEncryptionCbcsResponseArgs) ToCommonEncryptionCbcsResponsePtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCbcsResponseOutput).ToCommonEncryptionCbcsResponsePtrOutputWithContext(ctx)
}









type CommonEncryptionCbcsResponsePtrInput interface {
	pulumi.Input

	ToCommonEncryptionCbcsResponsePtrOutput() CommonEncryptionCbcsResponsePtrOutput
	ToCommonEncryptionCbcsResponsePtrOutputWithContext(context.Context) CommonEncryptionCbcsResponsePtrOutput
}

type commonEncryptionCbcsResponsePtrType CommonEncryptionCbcsResponseArgs

func CommonEncryptionCbcsResponsePtr(v *CommonEncryptionCbcsResponseArgs) CommonEncryptionCbcsResponsePtrInput {
	return (*commonEncryptionCbcsResponsePtrType)(v)
}

func (*commonEncryptionCbcsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCbcsResponse)(nil)).Elem()
}

func (i *commonEncryptionCbcsResponsePtrType) ToCommonEncryptionCbcsResponsePtrOutput() CommonEncryptionCbcsResponsePtrOutput {
	return i.ToCommonEncryptionCbcsResponsePtrOutputWithContext(context.Background())
}

func (i *commonEncryptionCbcsResponsePtrType) ToCommonEncryptionCbcsResponsePtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCbcsResponsePtrOutput)
}

type CommonEncryptionCbcsResponseOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCbcsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCbcsResponse)(nil)).Elem()
}

func (o CommonEncryptionCbcsResponseOutput) ToCommonEncryptionCbcsResponseOutput() CommonEncryptionCbcsResponseOutput {
	return o
}

func (o CommonEncryptionCbcsResponseOutput) ToCommonEncryptionCbcsResponseOutputWithContext(ctx context.Context) CommonEncryptionCbcsResponseOutput {
	return o
}

func (o CommonEncryptionCbcsResponseOutput) ToCommonEncryptionCbcsResponsePtrOutput() CommonEncryptionCbcsResponsePtrOutput {
	return o.ToCommonEncryptionCbcsResponsePtrOutputWithContext(context.Background())
}

func (o CommonEncryptionCbcsResponseOutput) ToCommonEncryptionCbcsResponsePtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommonEncryptionCbcsResponse) *CommonEncryptionCbcsResponse {
		return &v
	}).(CommonEncryptionCbcsResponsePtrOutput)
}

func (o CommonEncryptionCbcsResponseOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v CommonEncryptionCbcsResponse) []TrackSelectionResponse { return v.ClearTracks }).(TrackSelectionResponseArrayOutput)
}

func (o CommonEncryptionCbcsResponseOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcsResponse) *StreamingPolicyContentKeysResponse { return v.ContentKeys }).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o CommonEncryptionCbcsResponseOutput) Drm() CbcsDrmConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcsResponse) *CbcsDrmConfigurationResponse { return v.Drm }).(CbcsDrmConfigurationResponsePtrOutput)
}

func (o CommonEncryptionCbcsResponseOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcsResponse) *EnabledProtocolsResponse { return v.EnabledProtocols }).(EnabledProtocolsResponsePtrOutput)
}

type CommonEncryptionCbcsResponsePtrOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCbcsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCbcsResponse)(nil)).Elem()
}

func (o CommonEncryptionCbcsResponsePtrOutput) ToCommonEncryptionCbcsResponsePtrOutput() CommonEncryptionCbcsResponsePtrOutput {
	return o
}

func (o CommonEncryptionCbcsResponsePtrOutput) ToCommonEncryptionCbcsResponsePtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsResponsePtrOutput {
	return o
}

func (o CommonEncryptionCbcsResponsePtrOutput) Elem() CommonEncryptionCbcsResponseOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcsResponse) CommonEncryptionCbcsResponse {
		if v != nil {
			return *v
		}
		var ret CommonEncryptionCbcsResponse
		return ret
	}).(CommonEncryptionCbcsResponseOutput)
}

func (o CommonEncryptionCbcsResponsePtrOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcsResponse) []TrackSelectionResponse {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionResponseArrayOutput)
}

func (o CommonEncryptionCbcsResponsePtrOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcsResponse) *StreamingPolicyContentKeysResponse {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o CommonEncryptionCbcsResponsePtrOutput) Drm() CbcsDrmConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcsResponse) *CbcsDrmConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Drm
	}).(CbcsDrmConfigurationResponsePtrOutput)
}

func (o CommonEncryptionCbcsResponsePtrOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcsResponse) *EnabledProtocolsResponse {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsResponsePtrOutput)
}

type CommonEncryptionCenc struct {
	ClearTracks      []TrackSelection            `pulumi:"clearTracks"`
	ContentKeys      *StreamingPolicyContentKeys `pulumi:"contentKeys"`
	Drm              *CencDrmConfiguration       `pulumi:"drm"`
	EnabledProtocols *EnabledProtocols           `pulumi:"enabledProtocols"`
}





type CommonEncryptionCencInput interface {
	pulumi.Input

	ToCommonEncryptionCencOutput() CommonEncryptionCencOutput
	ToCommonEncryptionCencOutputWithContext(context.Context) CommonEncryptionCencOutput
}

type CommonEncryptionCencArgs struct {
	ClearTracks      TrackSelectionArrayInput           `pulumi:"clearTracks"`
	ContentKeys      StreamingPolicyContentKeysPtrInput `pulumi:"contentKeys"`
	Drm              CencDrmConfigurationPtrInput       `pulumi:"drm"`
	EnabledProtocols EnabledProtocolsPtrInput           `pulumi:"enabledProtocols"`
}

func (CommonEncryptionCencArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCenc)(nil)).Elem()
}

func (i CommonEncryptionCencArgs) ToCommonEncryptionCencOutput() CommonEncryptionCencOutput {
	return i.ToCommonEncryptionCencOutputWithContext(context.Background())
}

func (i CommonEncryptionCencArgs) ToCommonEncryptionCencOutputWithContext(ctx context.Context) CommonEncryptionCencOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCencOutput)
}

func (i CommonEncryptionCencArgs) ToCommonEncryptionCencPtrOutput() CommonEncryptionCencPtrOutput {
	return i.ToCommonEncryptionCencPtrOutputWithContext(context.Background())
}

func (i CommonEncryptionCencArgs) ToCommonEncryptionCencPtrOutputWithContext(ctx context.Context) CommonEncryptionCencPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCencOutput).ToCommonEncryptionCencPtrOutputWithContext(ctx)
}









type CommonEncryptionCencPtrInput interface {
	pulumi.Input

	ToCommonEncryptionCencPtrOutput() CommonEncryptionCencPtrOutput
	ToCommonEncryptionCencPtrOutputWithContext(context.Context) CommonEncryptionCencPtrOutput
}

type commonEncryptionCencPtrType CommonEncryptionCencArgs

func CommonEncryptionCencPtr(v *CommonEncryptionCencArgs) CommonEncryptionCencPtrInput {
	return (*commonEncryptionCencPtrType)(v)
}

func (*commonEncryptionCencPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCenc)(nil)).Elem()
}

func (i *commonEncryptionCencPtrType) ToCommonEncryptionCencPtrOutput() CommonEncryptionCencPtrOutput {
	return i.ToCommonEncryptionCencPtrOutputWithContext(context.Background())
}

func (i *commonEncryptionCencPtrType) ToCommonEncryptionCencPtrOutputWithContext(ctx context.Context) CommonEncryptionCencPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCencPtrOutput)
}

type CommonEncryptionCencOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCencOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCenc)(nil)).Elem()
}

func (o CommonEncryptionCencOutput) ToCommonEncryptionCencOutput() CommonEncryptionCencOutput {
	return o
}

func (o CommonEncryptionCencOutput) ToCommonEncryptionCencOutputWithContext(ctx context.Context) CommonEncryptionCencOutput {
	return o
}

func (o CommonEncryptionCencOutput) ToCommonEncryptionCencPtrOutput() CommonEncryptionCencPtrOutput {
	return o.ToCommonEncryptionCencPtrOutputWithContext(context.Background())
}

func (o CommonEncryptionCencOutput) ToCommonEncryptionCencPtrOutputWithContext(ctx context.Context) CommonEncryptionCencPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommonEncryptionCenc) *CommonEncryptionCenc {
		return &v
	}).(CommonEncryptionCencPtrOutput)
}

func (o CommonEncryptionCencOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v CommonEncryptionCenc) []TrackSelection { return v.ClearTracks }).(TrackSelectionArrayOutput)
}

func (o CommonEncryptionCencOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCenc) *StreamingPolicyContentKeys { return v.ContentKeys }).(StreamingPolicyContentKeysPtrOutput)
}

func (o CommonEncryptionCencOutput) Drm() CencDrmConfigurationPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCenc) *CencDrmConfiguration { return v.Drm }).(CencDrmConfigurationPtrOutput)
}

func (o CommonEncryptionCencOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCenc) *EnabledProtocols { return v.EnabledProtocols }).(EnabledProtocolsPtrOutput)
}

type CommonEncryptionCencPtrOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCencPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCenc)(nil)).Elem()
}

func (o CommonEncryptionCencPtrOutput) ToCommonEncryptionCencPtrOutput() CommonEncryptionCencPtrOutput {
	return o
}

func (o CommonEncryptionCencPtrOutput) ToCommonEncryptionCencPtrOutputWithContext(ctx context.Context) CommonEncryptionCencPtrOutput {
	return o
}

func (o CommonEncryptionCencPtrOutput) Elem() CommonEncryptionCencOutput {
	return o.ApplyT(func(v *CommonEncryptionCenc) CommonEncryptionCenc {
		if v != nil {
			return *v
		}
		var ret CommonEncryptionCenc
		return ret
	}).(CommonEncryptionCencOutput)
}

func (o CommonEncryptionCencPtrOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v *CommonEncryptionCenc) []TrackSelection {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionArrayOutput)
}

func (o CommonEncryptionCencPtrOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCenc) *StreamingPolicyContentKeys {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysPtrOutput)
}

func (o CommonEncryptionCencPtrOutput) Drm() CencDrmConfigurationPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCenc) *CencDrmConfiguration {
		if v == nil {
			return nil
		}
		return v.Drm
	}).(CencDrmConfigurationPtrOutput)
}

func (o CommonEncryptionCencPtrOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCenc) *EnabledProtocols {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsPtrOutput)
}

type CommonEncryptionCencResponse struct {
	ClearTracks      []TrackSelectionResponse            `pulumi:"clearTracks"`
	ContentKeys      *StreamingPolicyContentKeysResponse `pulumi:"contentKeys"`
	Drm              *CencDrmConfigurationResponse       `pulumi:"drm"`
	EnabledProtocols *EnabledProtocolsResponse           `pulumi:"enabledProtocols"`
}





type CommonEncryptionCencResponseInput interface {
	pulumi.Input

	ToCommonEncryptionCencResponseOutput() CommonEncryptionCencResponseOutput
	ToCommonEncryptionCencResponseOutputWithContext(context.Context) CommonEncryptionCencResponseOutput
}

type CommonEncryptionCencResponseArgs struct {
	ClearTracks      TrackSelectionResponseArrayInput           `pulumi:"clearTracks"`
	ContentKeys      StreamingPolicyContentKeysResponsePtrInput `pulumi:"contentKeys"`
	Drm              CencDrmConfigurationResponsePtrInput       `pulumi:"drm"`
	EnabledProtocols EnabledProtocolsResponsePtrInput           `pulumi:"enabledProtocols"`
}

func (CommonEncryptionCencResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCencResponse)(nil)).Elem()
}

func (i CommonEncryptionCencResponseArgs) ToCommonEncryptionCencResponseOutput() CommonEncryptionCencResponseOutput {
	return i.ToCommonEncryptionCencResponseOutputWithContext(context.Background())
}

func (i CommonEncryptionCencResponseArgs) ToCommonEncryptionCencResponseOutputWithContext(ctx context.Context) CommonEncryptionCencResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCencResponseOutput)
}

func (i CommonEncryptionCencResponseArgs) ToCommonEncryptionCencResponsePtrOutput() CommonEncryptionCencResponsePtrOutput {
	return i.ToCommonEncryptionCencResponsePtrOutputWithContext(context.Background())
}

func (i CommonEncryptionCencResponseArgs) ToCommonEncryptionCencResponsePtrOutputWithContext(ctx context.Context) CommonEncryptionCencResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCencResponseOutput).ToCommonEncryptionCencResponsePtrOutputWithContext(ctx)
}









type CommonEncryptionCencResponsePtrInput interface {
	pulumi.Input

	ToCommonEncryptionCencResponsePtrOutput() CommonEncryptionCencResponsePtrOutput
	ToCommonEncryptionCencResponsePtrOutputWithContext(context.Context) CommonEncryptionCencResponsePtrOutput
}

type commonEncryptionCencResponsePtrType CommonEncryptionCencResponseArgs

func CommonEncryptionCencResponsePtr(v *CommonEncryptionCencResponseArgs) CommonEncryptionCencResponsePtrInput {
	return (*commonEncryptionCencResponsePtrType)(v)
}

func (*commonEncryptionCencResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCencResponse)(nil)).Elem()
}

func (i *commonEncryptionCencResponsePtrType) ToCommonEncryptionCencResponsePtrOutput() CommonEncryptionCencResponsePtrOutput {
	return i.ToCommonEncryptionCencResponsePtrOutputWithContext(context.Background())
}

func (i *commonEncryptionCencResponsePtrType) ToCommonEncryptionCencResponsePtrOutputWithContext(ctx context.Context) CommonEncryptionCencResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCencResponsePtrOutput)
}

type CommonEncryptionCencResponseOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCencResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCencResponse)(nil)).Elem()
}

func (o CommonEncryptionCencResponseOutput) ToCommonEncryptionCencResponseOutput() CommonEncryptionCencResponseOutput {
	return o
}

func (o CommonEncryptionCencResponseOutput) ToCommonEncryptionCencResponseOutputWithContext(ctx context.Context) CommonEncryptionCencResponseOutput {
	return o
}

func (o CommonEncryptionCencResponseOutput) ToCommonEncryptionCencResponsePtrOutput() CommonEncryptionCencResponsePtrOutput {
	return o.ToCommonEncryptionCencResponsePtrOutputWithContext(context.Background())
}

func (o CommonEncryptionCencResponseOutput) ToCommonEncryptionCencResponsePtrOutputWithContext(ctx context.Context) CommonEncryptionCencResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommonEncryptionCencResponse) *CommonEncryptionCencResponse {
		return &v
	}).(CommonEncryptionCencResponsePtrOutput)
}

func (o CommonEncryptionCencResponseOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v CommonEncryptionCencResponse) []TrackSelectionResponse { return v.ClearTracks }).(TrackSelectionResponseArrayOutput)
}

func (o CommonEncryptionCencResponseOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCencResponse) *StreamingPolicyContentKeysResponse { return v.ContentKeys }).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o CommonEncryptionCencResponseOutput) Drm() CencDrmConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCencResponse) *CencDrmConfigurationResponse { return v.Drm }).(CencDrmConfigurationResponsePtrOutput)
}

func (o CommonEncryptionCencResponseOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCencResponse) *EnabledProtocolsResponse { return v.EnabledProtocols }).(EnabledProtocolsResponsePtrOutput)
}

type CommonEncryptionCencResponsePtrOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCencResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCencResponse)(nil)).Elem()
}

func (o CommonEncryptionCencResponsePtrOutput) ToCommonEncryptionCencResponsePtrOutput() CommonEncryptionCencResponsePtrOutput {
	return o
}

func (o CommonEncryptionCencResponsePtrOutput) ToCommonEncryptionCencResponsePtrOutputWithContext(ctx context.Context) CommonEncryptionCencResponsePtrOutput {
	return o
}

func (o CommonEncryptionCencResponsePtrOutput) Elem() CommonEncryptionCencResponseOutput {
	return o.ApplyT(func(v *CommonEncryptionCencResponse) CommonEncryptionCencResponse {
		if v != nil {
			return *v
		}
		var ret CommonEncryptionCencResponse
		return ret
	}).(CommonEncryptionCencResponseOutput)
}

func (o CommonEncryptionCencResponsePtrOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v *CommonEncryptionCencResponse) []TrackSelectionResponse {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionResponseArrayOutput)
}

func (o CommonEncryptionCencResponsePtrOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCencResponse) *StreamingPolicyContentKeysResponse {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o CommonEncryptionCencResponsePtrOutput) Drm() CencDrmConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCencResponse) *CencDrmConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Drm
	}).(CencDrmConfigurationResponsePtrOutput)
}

func (o CommonEncryptionCencResponsePtrOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCencResponse) *EnabledProtocolsResponse {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsResponsePtrOutput)
}

type ContentKeyPolicyClearKeyConfiguration struct {
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyClearKeyConfigurationInput interface {
	pulumi.Input

	ToContentKeyPolicyClearKeyConfigurationOutput() ContentKeyPolicyClearKeyConfigurationOutput
	ToContentKeyPolicyClearKeyConfigurationOutputWithContext(context.Context) ContentKeyPolicyClearKeyConfigurationOutput
}

type ContentKeyPolicyClearKeyConfigurationArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyClearKeyConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyClearKeyConfiguration)(nil)).Elem()
}

func (i ContentKeyPolicyClearKeyConfigurationArgs) ToContentKeyPolicyClearKeyConfigurationOutput() ContentKeyPolicyClearKeyConfigurationOutput {
	return i.ToContentKeyPolicyClearKeyConfigurationOutputWithContext(context.Background())
}

func (i ContentKeyPolicyClearKeyConfigurationArgs) ToContentKeyPolicyClearKeyConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyClearKeyConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyClearKeyConfigurationOutput)
}

type ContentKeyPolicyClearKeyConfigurationOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyClearKeyConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyClearKeyConfiguration)(nil)).Elem()
}

func (o ContentKeyPolicyClearKeyConfigurationOutput) ToContentKeyPolicyClearKeyConfigurationOutput() ContentKeyPolicyClearKeyConfigurationOutput {
	return o
}

func (o ContentKeyPolicyClearKeyConfigurationOutput) ToContentKeyPolicyClearKeyConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyClearKeyConfigurationOutput {
	return o
}

func (o ContentKeyPolicyClearKeyConfigurationOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyClearKeyConfiguration) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyClearKeyConfigurationResponse struct {
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyClearKeyConfigurationResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyClearKeyConfigurationResponseOutput() ContentKeyPolicyClearKeyConfigurationResponseOutput
	ToContentKeyPolicyClearKeyConfigurationResponseOutputWithContext(context.Context) ContentKeyPolicyClearKeyConfigurationResponseOutput
}

type ContentKeyPolicyClearKeyConfigurationResponseArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyClearKeyConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyClearKeyConfigurationResponse)(nil)).Elem()
}

func (i ContentKeyPolicyClearKeyConfigurationResponseArgs) ToContentKeyPolicyClearKeyConfigurationResponseOutput() ContentKeyPolicyClearKeyConfigurationResponseOutput {
	return i.ToContentKeyPolicyClearKeyConfigurationResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyClearKeyConfigurationResponseArgs) ToContentKeyPolicyClearKeyConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyClearKeyConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyClearKeyConfigurationResponseOutput)
}

type ContentKeyPolicyClearKeyConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyClearKeyConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyClearKeyConfigurationResponse)(nil)).Elem()
}

func (o ContentKeyPolicyClearKeyConfigurationResponseOutput) ToContentKeyPolicyClearKeyConfigurationResponseOutput() ContentKeyPolicyClearKeyConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyClearKeyConfigurationResponseOutput) ToContentKeyPolicyClearKeyConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyClearKeyConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyClearKeyConfigurationResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyClearKeyConfigurationResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyFairPlayConfiguration struct {
	Ask                        string                                              `pulumi:"ask"`
	FairPlayPfx                string                                              `pulumi:"fairPlayPfx"`
	FairPlayPfxPassword        string                                              `pulumi:"fairPlayPfxPassword"`
	OdataType                  string                                              `pulumi:"odataType"`
	OfflineRentalConfiguration *ContentKeyPolicyFairPlayOfflineRentalConfiguration `pulumi:"offlineRentalConfiguration"`
	RentalAndLeaseKeyType      string                                              `pulumi:"rentalAndLeaseKeyType"`
	RentalDuration             float64                                             `pulumi:"rentalDuration"`
}





type ContentKeyPolicyFairPlayConfigurationInput interface {
	pulumi.Input

	ToContentKeyPolicyFairPlayConfigurationOutput() ContentKeyPolicyFairPlayConfigurationOutput
	ToContentKeyPolicyFairPlayConfigurationOutputWithContext(context.Context) ContentKeyPolicyFairPlayConfigurationOutput
}

type ContentKeyPolicyFairPlayConfigurationArgs struct {
	Ask                        pulumi.StringInput                                         `pulumi:"ask"`
	FairPlayPfx                pulumi.StringInput                                         `pulumi:"fairPlayPfx"`
	FairPlayPfxPassword        pulumi.StringInput                                         `pulumi:"fairPlayPfxPassword"`
	OdataType                  pulumi.StringInput                                         `pulumi:"odataType"`
	OfflineRentalConfiguration ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrInput `pulumi:"offlineRentalConfiguration"`
	RentalAndLeaseKeyType      pulumi.StringInput                                         `pulumi:"rentalAndLeaseKeyType"`
	RentalDuration             pulumi.Float64Input                                        `pulumi:"rentalDuration"`
}

func (ContentKeyPolicyFairPlayConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyFairPlayConfiguration)(nil)).Elem()
}

func (i ContentKeyPolicyFairPlayConfigurationArgs) ToContentKeyPolicyFairPlayConfigurationOutput() ContentKeyPolicyFairPlayConfigurationOutput {
	return i.ToContentKeyPolicyFairPlayConfigurationOutputWithContext(context.Background())
}

func (i ContentKeyPolicyFairPlayConfigurationArgs) ToContentKeyPolicyFairPlayConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyFairPlayConfigurationOutput)
}

type ContentKeyPolicyFairPlayConfigurationOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyFairPlayConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyFairPlayConfiguration)(nil)).Elem()
}

func (o ContentKeyPolicyFairPlayConfigurationOutput) ToContentKeyPolicyFairPlayConfigurationOutput() ContentKeyPolicyFairPlayConfigurationOutput {
	return o
}

func (o ContentKeyPolicyFairPlayConfigurationOutput) ToContentKeyPolicyFairPlayConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayConfigurationOutput {
	return o
}

func (o ContentKeyPolicyFairPlayConfigurationOutput) Ask() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfiguration) string { return v.Ask }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationOutput) FairPlayPfx() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfiguration) string { return v.FairPlayPfx }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationOutput) FairPlayPfxPassword() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfiguration) string { return v.FairPlayPfxPassword }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfiguration) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationOutput) OfflineRentalConfiguration() ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfiguration) *ContentKeyPolicyFairPlayOfflineRentalConfiguration {
		return v.OfflineRentalConfiguration
	}).(ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationOutput) RentalAndLeaseKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfiguration) string { return v.RentalAndLeaseKeyType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationOutput) RentalDuration() pulumi.Float64Output {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfiguration) float64 { return v.RentalDuration }).(pulumi.Float64Output)
}

type ContentKeyPolicyFairPlayConfigurationResponse struct {
	Ask                        string                                                      `pulumi:"ask"`
	FairPlayPfx                string                                                      `pulumi:"fairPlayPfx"`
	FairPlayPfxPassword        string                                                      `pulumi:"fairPlayPfxPassword"`
	OdataType                  string                                                      `pulumi:"odataType"`
	OfflineRentalConfiguration *ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse `pulumi:"offlineRentalConfiguration"`
	RentalAndLeaseKeyType      string                                                      `pulumi:"rentalAndLeaseKeyType"`
	RentalDuration             float64                                                     `pulumi:"rentalDuration"`
}





type ContentKeyPolicyFairPlayConfigurationResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyFairPlayConfigurationResponseOutput() ContentKeyPolicyFairPlayConfigurationResponseOutput
	ToContentKeyPolicyFairPlayConfigurationResponseOutputWithContext(context.Context) ContentKeyPolicyFairPlayConfigurationResponseOutput
}

type ContentKeyPolicyFairPlayConfigurationResponseArgs struct {
	Ask                        pulumi.StringInput                                                 `pulumi:"ask"`
	FairPlayPfx                pulumi.StringInput                                                 `pulumi:"fairPlayPfx"`
	FairPlayPfxPassword        pulumi.StringInput                                                 `pulumi:"fairPlayPfxPassword"`
	OdataType                  pulumi.StringInput                                                 `pulumi:"odataType"`
	OfflineRentalConfiguration ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrInput `pulumi:"offlineRentalConfiguration"`
	RentalAndLeaseKeyType      pulumi.StringInput                                                 `pulumi:"rentalAndLeaseKeyType"`
	RentalDuration             pulumi.Float64Input                                                `pulumi:"rentalDuration"`
}

func (ContentKeyPolicyFairPlayConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyFairPlayConfigurationResponse)(nil)).Elem()
}

func (i ContentKeyPolicyFairPlayConfigurationResponseArgs) ToContentKeyPolicyFairPlayConfigurationResponseOutput() ContentKeyPolicyFairPlayConfigurationResponseOutput {
	return i.ToContentKeyPolicyFairPlayConfigurationResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyFairPlayConfigurationResponseArgs) ToContentKeyPolicyFairPlayConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyFairPlayConfigurationResponseOutput)
}

type ContentKeyPolicyFairPlayConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyFairPlayConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyFairPlayConfigurationResponse)(nil)).Elem()
}

func (o ContentKeyPolicyFairPlayConfigurationResponseOutput) ToContentKeyPolicyFairPlayConfigurationResponseOutput() ContentKeyPolicyFairPlayConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyFairPlayConfigurationResponseOutput) ToContentKeyPolicyFairPlayConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyFairPlayConfigurationResponseOutput) Ask() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfigurationResponse) string { return v.Ask }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationResponseOutput) FairPlayPfx() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfigurationResponse) string { return v.FairPlayPfx }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationResponseOutput) FairPlayPfxPassword() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfigurationResponse) string { return v.FairPlayPfxPassword }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfigurationResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationResponseOutput) OfflineRentalConfiguration() ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfigurationResponse) *ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse {
		return v.OfflineRentalConfiguration
	}).(ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationResponseOutput) RentalAndLeaseKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfigurationResponse) string { return v.RentalAndLeaseKeyType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyFairPlayConfigurationResponseOutput) RentalDuration() pulumi.Float64Output {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayConfigurationResponse) float64 { return v.RentalDuration }).(pulumi.Float64Output)
}

type ContentKeyPolicyFairPlayOfflineRentalConfiguration struct {
	PlaybackDurationSeconds float64 `pulumi:"playbackDurationSeconds"`
	StorageDurationSeconds  float64 `pulumi:"storageDurationSeconds"`
}





type ContentKeyPolicyFairPlayOfflineRentalConfigurationInput interface {
	pulumi.Input

	ToContentKeyPolicyFairPlayOfflineRentalConfigurationOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput
	ToContentKeyPolicyFairPlayOfflineRentalConfigurationOutputWithContext(context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput
}

type ContentKeyPolicyFairPlayOfflineRentalConfigurationArgs struct {
	PlaybackDurationSeconds pulumi.Float64Input `pulumi:"playbackDurationSeconds"`
	StorageDurationSeconds  pulumi.Float64Input `pulumi:"storageDurationSeconds"`
}

func (ContentKeyPolicyFairPlayOfflineRentalConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyFairPlayOfflineRentalConfiguration)(nil)).Elem()
}

func (i ContentKeyPolicyFairPlayOfflineRentalConfigurationArgs) ToContentKeyPolicyFairPlayOfflineRentalConfigurationOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput {
	return i.ToContentKeyPolicyFairPlayOfflineRentalConfigurationOutputWithContext(context.Background())
}

func (i ContentKeyPolicyFairPlayOfflineRentalConfigurationArgs) ToContentKeyPolicyFairPlayOfflineRentalConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput)
}

func (i ContentKeyPolicyFairPlayOfflineRentalConfigurationArgs) ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput {
	return i.ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutputWithContext(context.Background())
}

func (i ContentKeyPolicyFairPlayOfflineRentalConfigurationArgs) ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput).ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutputWithContext(ctx)
}









type ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrInput interface {
	pulumi.Input

	ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput
	ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutputWithContext(context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput
}

type contentKeyPolicyFairPlayOfflineRentalConfigurationPtrType ContentKeyPolicyFairPlayOfflineRentalConfigurationArgs

func ContentKeyPolicyFairPlayOfflineRentalConfigurationPtr(v *ContentKeyPolicyFairPlayOfflineRentalConfigurationArgs) ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrInput {
	return (*contentKeyPolicyFairPlayOfflineRentalConfigurationPtrType)(v)
}

func (*contentKeyPolicyFairPlayOfflineRentalConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyFairPlayOfflineRentalConfiguration)(nil)).Elem()
}

func (i *contentKeyPolicyFairPlayOfflineRentalConfigurationPtrType) ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput {
	return i.ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutputWithContext(context.Background())
}

func (i *contentKeyPolicyFairPlayOfflineRentalConfigurationPtrType) ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput)
}

type ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyFairPlayOfflineRentalConfiguration)(nil)).Elem()
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput {
	return o
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput {
	return o
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput {
	return o.ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentKeyPolicyFairPlayOfflineRentalConfiguration) *ContentKeyPolicyFairPlayOfflineRentalConfiguration {
		return &v
	}).(ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput)
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput) PlaybackDurationSeconds() pulumi.Float64Output {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayOfflineRentalConfiguration) float64 { return v.PlaybackDurationSeconds }).(pulumi.Float64Output)
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput) StorageDurationSeconds() pulumi.Float64Output {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayOfflineRentalConfiguration) float64 { return v.StorageDurationSeconds }).(pulumi.Float64Output)
}

type ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyFairPlayOfflineRentalConfiguration)(nil)).Elem()
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput {
	return o
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput {
	return o
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput) Elem() ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput {
	return o.ApplyT(func(v *ContentKeyPolicyFairPlayOfflineRentalConfiguration) ContentKeyPolicyFairPlayOfflineRentalConfiguration {
		if v != nil {
			return *v
		}
		var ret ContentKeyPolicyFairPlayOfflineRentalConfiguration
		return ret
	}).(ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput)
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput) PlaybackDurationSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyFairPlayOfflineRentalConfiguration) *float64 {
		if v == nil {
			return nil
		}
		return &v.PlaybackDurationSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput) StorageDurationSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyFairPlayOfflineRentalConfiguration) *float64 {
		if v == nil {
			return nil
		}
		return &v.StorageDurationSeconds
	}).(pulumi.Float64PtrOutput)
}

type ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse struct {
	PlaybackDurationSeconds float64 `pulumi:"playbackDurationSeconds"`
	StorageDurationSeconds  float64 `pulumi:"storageDurationSeconds"`
}





type ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput
	ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutputWithContext(context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput
}

type ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseArgs struct {
	PlaybackDurationSeconds pulumi.Float64Input `pulumi:"playbackDurationSeconds"`
	StorageDurationSeconds  pulumi.Float64Input `pulumi:"storageDurationSeconds"`
}

func (ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse)(nil)).Elem()
}

func (i ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseArgs) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput {
	return i.ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseArgs) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput)
}

func (i ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseArgs) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput {
	return i.ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseArgs) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput).ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutputWithContext(ctx)
}









type ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrInput interface {
	pulumi.Input

	ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput
	ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutputWithContext(context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput
}

type contentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrType ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseArgs

func ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtr(v *ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseArgs) ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrInput {
	return (*contentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrType)(v)
}

func (*contentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse)(nil)).Elem()
}

func (i *contentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrType) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput {
	return i.ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *contentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrType) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput)
}

type ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse)(nil)).Elem()
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput {
	return o.ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse) *ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse {
		return &v
	}).(ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput)
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput) PlaybackDurationSeconds() pulumi.Float64Output {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse) float64 {
		return v.PlaybackDurationSeconds
	}).(pulumi.Float64Output)
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput) StorageDurationSeconds() pulumi.Float64Output {
	return o.ApplyT(func(v ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse) float64 {
		return v.StorageDurationSeconds
	}).(pulumi.Float64Output)
}

type ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse)(nil)).Elem()
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput() ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput {
	return o
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput) ToContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput {
	return o
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput) Elem() ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput {
	return o.ApplyT(func(v *ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse) ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse
		return ret
	}).(ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput)
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput) PlaybackDurationSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.PlaybackDurationSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput) StorageDurationSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyFairPlayOfflineRentalConfigurationResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.StorageDurationSeconds
	}).(pulumi.Float64PtrOutput)
}

type ContentKeyPolicyOpenRestriction struct {
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyOpenRestrictionInput interface {
	pulumi.Input

	ToContentKeyPolicyOpenRestrictionOutput() ContentKeyPolicyOpenRestrictionOutput
	ToContentKeyPolicyOpenRestrictionOutputWithContext(context.Context) ContentKeyPolicyOpenRestrictionOutput
}

type ContentKeyPolicyOpenRestrictionArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyOpenRestrictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyOpenRestriction)(nil)).Elem()
}

func (i ContentKeyPolicyOpenRestrictionArgs) ToContentKeyPolicyOpenRestrictionOutput() ContentKeyPolicyOpenRestrictionOutput {
	return i.ToContentKeyPolicyOpenRestrictionOutputWithContext(context.Background())
}

func (i ContentKeyPolicyOpenRestrictionArgs) ToContentKeyPolicyOpenRestrictionOutputWithContext(ctx context.Context) ContentKeyPolicyOpenRestrictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyOpenRestrictionOutput)
}

type ContentKeyPolicyOpenRestrictionOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOpenRestrictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyOpenRestriction)(nil)).Elem()
}

func (o ContentKeyPolicyOpenRestrictionOutput) ToContentKeyPolicyOpenRestrictionOutput() ContentKeyPolicyOpenRestrictionOutput {
	return o
}

func (o ContentKeyPolicyOpenRestrictionOutput) ToContentKeyPolicyOpenRestrictionOutputWithContext(ctx context.Context) ContentKeyPolicyOpenRestrictionOutput {
	return o
}

func (o ContentKeyPolicyOpenRestrictionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyOpenRestriction) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyOpenRestrictionResponse struct {
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyOpenRestrictionResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyOpenRestrictionResponseOutput() ContentKeyPolicyOpenRestrictionResponseOutput
	ToContentKeyPolicyOpenRestrictionResponseOutputWithContext(context.Context) ContentKeyPolicyOpenRestrictionResponseOutput
}

type ContentKeyPolicyOpenRestrictionResponseArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyOpenRestrictionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyOpenRestrictionResponse)(nil)).Elem()
}

func (i ContentKeyPolicyOpenRestrictionResponseArgs) ToContentKeyPolicyOpenRestrictionResponseOutput() ContentKeyPolicyOpenRestrictionResponseOutput {
	return i.ToContentKeyPolicyOpenRestrictionResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyOpenRestrictionResponseArgs) ToContentKeyPolicyOpenRestrictionResponseOutputWithContext(ctx context.Context) ContentKeyPolicyOpenRestrictionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyOpenRestrictionResponseOutput)
}

type ContentKeyPolicyOpenRestrictionResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOpenRestrictionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyOpenRestrictionResponse)(nil)).Elem()
}

func (o ContentKeyPolicyOpenRestrictionResponseOutput) ToContentKeyPolicyOpenRestrictionResponseOutput() ContentKeyPolicyOpenRestrictionResponseOutput {
	return o
}

func (o ContentKeyPolicyOpenRestrictionResponseOutput) ToContentKeyPolicyOpenRestrictionResponseOutputWithContext(ctx context.Context) ContentKeyPolicyOpenRestrictionResponseOutput {
	return o
}

func (o ContentKeyPolicyOpenRestrictionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyOpenRestrictionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyOption struct {
	Configuration interface{} `pulumi:"configuration"`
	Name          *string     `pulumi:"name"`
	Restriction   interface{} `pulumi:"restriction"`
}





type ContentKeyPolicyOptionInput interface {
	pulumi.Input

	ToContentKeyPolicyOptionOutput() ContentKeyPolicyOptionOutput
	ToContentKeyPolicyOptionOutputWithContext(context.Context) ContentKeyPolicyOptionOutput
}

type ContentKeyPolicyOptionArgs struct {
	Configuration pulumi.Input          `pulumi:"configuration"`
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Restriction   pulumi.Input          `pulumi:"restriction"`
}

func (ContentKeyPolicyOptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyOption)(nil)).Elem()
}

func (i ContentKeyPolicyOptionArgs) ToContentKeyPolicyOptionOutput() ContentKeyPolicyOptionOutput {
	return i.ToContentKeyPolicyOptionOutputWithContext(context.Background())
}

func (i ContentKeyPolicyOptionArgs) ToContentKeyPolicyOptionOutputWithContext(ctx context.Context) ContentKeyPolicyOptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyOptionOutput)
}





type ContentKeyPolicyOptionArrayInput interface {
	pulumi.Input

	ToContentKeyPolicyOptionArrayOutput() ContentKeyPolicyOptionArrayOutput
	ToContentKeyPolicyOptionArrayOutputWithContext(context.Context) ContentKeyPolicyOptionArrayOutput
}

type ContentKeyPolicyOptionArray []ContentKeyPolicyOptionInput

func (ContentKeyPolicyOptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyOption)(nil)).Elem()
}

func (i ContentKeyPolicyOptionArray) ToContentKeyPolicyOptionArrayOutput() ContentKeyPolicyOptionArrayOutput {
	return i.ToContentKeyPolicyOptionArrayOutputWithContext(context.Background())
}

func (i ContentKeyPolicyOptionArray) ToContentKeyPolicyOptionArrayOutputWithContext(ctx context.Context) ContentKeyPolicyOptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyOptionArrayOutput)
}

type ContentKeyPolicyOptionOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyOption)(nil)).Elem()
}

func (o ContentKeyPolicyOptionOutput) ToContentKeyPolicyOptionOutput() ContentKeyPolicyOptionOutput {
	return o
}

func (o ContentKeyPolicyOptionOutput) ToContentKeyPolicyOptionOutputWithContext(ctx context.Context) ContentKeyPolicyOptionOutput {
	return o
}

func (o ContentKeyPolicyOptionOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyOption) interface{} { return v.Configuration }).(pulumi.AnyOutput)
}

func (o ContentKeyPolicyOptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyOption) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyOptionOutput) Restriction() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyOption) interface{} { return v.Restriction }).(pulumi.AnyOutput)
}

type ContentKeyPolicyOptionArrayOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyOption)(nil)).Elem()
}

func (o ContentKeyPolicyOptionArrayOutput) ToContentKeyPolicyOptionArrayOutput() ContentKeyPolicyOptionArrayOutput {
	return o
}

func (o ContentKeyPolicyOptionArrayOutput) ToContentKeyPolicyOptionArrayOutputWithContext(ctx context.Context) ContentKeyPolicyOptionArrayOutput {
	return o
}

func (o ContentKeyPolicyOptionArrayOutput) Index(i pulumi.IntInput) ContentKeyPolicyOptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContentKeyPolicyOption {
		return vs[0].([]ContentKeyPolicyOption)[vs[1].(int)]
	}).(ContentKeyPolicyOptionOutput)
}

type ContentKeyPolicyOptionResponse struct {
	Configuration  interface{} `pulumi:"configuration"`
	Name           *string     `pulumi:"name"`
	PolicyOptionId string      `pulumi:"policyOptionId"`
	Restriction    interface{} `pulumi:"restriction"`
}





type ContentKeyPolicyOptionResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyOptionResponseOutput() ContentKeyPolicyOptionResponseOutput
	ToContentKeyPolicyOptionResponseOutputWithContext(context.Context) ContentKeyPolicyOptionResponseOutput
}

type ContentKeyPolicyOptionResponseArgs struct {
	Configuration  pulumi.Input          `pulumi:"configuration"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
	PolicyOptionId pulumi.StringInput    `pulumi:"policyOptionId"`
	Restriction    pulumi.Input          `pulumi:"restriction"`
}

func (ContentKeyPolicyOptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyOptionResponse)(nil)).Elem()
}

func (i ContentKeyPolicyOptionResponseArgs) ToContentKeyPolicyOptionResponseOutput() ContentKeyPolicyOptionResponseOutput {
	return i.ToContentKeyPolicyOptionResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyOptionResponseArgs) ToContentKeyPolicyOptionResponseOutputWithContext(ctx context.Context) ContentKeyPolicyOptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyOptionResponseOutput)
}





type ContentKeyPolicyOptionResponseArrayInput interface {
	pulumi.Input

	ToContentKeyPolicyOptionResponseArrayOutput() ContentKeyPolicyOptionResponseArrayOutput
	ToContentKeyPolicyOptionResponseArrayOutputWithContext(context.Context) ContentKeyPolicyOptionResponseArrayOutput
}

type ContentKeyPolicyOptionResponseArray []ContentKeyPolicyOptionResponseInput

func (ContentKeyPolicyOptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyOptionResponse)(nil)).Elem()
}

func (i ContentKeyPolicyOptionResponseArray) ToContentKeyPolicyOptionResponseArrayOutput() ContentKeyPolicyOptionResponseArrayOutput {
	return i.ToContentKeyPolicyOptionResponseArrayOutputWithContext(context.Background())
}

func (i ContentKeyPolicyOptionResponseArray) ToContentKeyPolicyOptionResponseArrayOutputWithContext(ctx context.Context) ContentKeyPolicyOptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyOptionResponseArrayOutput)
}

type ContentKeyPolicyOptionResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyOptionResponse)(nil)).Elem()
}

func (o ContentKeyPolicyOptionResponseOutput) ToContentKeyPolicyOptionResponseOutput() ContentKeyPolicyOptionResponseOutput {
	return o
}

func (o ContentKeyPolicyOptionResponseOutput) ToContentKeyPolicyOptionResponseOutputWithContext(ctx context.Context) ContentKeyPolicyOptionResponseOutput {
	return o
}

func (o ContentKeyPolicyOptionResponseOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyOptionResponse) interface{} { return v.Configuration }).(pulumi.AnyOutput)
}

func (o ContentKeyPolicyOptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyOptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyOptionResponseOutput) PolicyOptionId() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyOptionResponse) string { return v.PolicyOptionId }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyOptionResponseOutput) Restriction() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyOptionResponse) interface{} { return v.Restriction }).(pulumi.AnyOutput)
}

type ContentKeyPolicyOptionResponseArrayOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyOptionResponse)(nil)).Elem()
}

func (o ContentKeyPolicyOptionResponseArrayOutput) ToContentKeyPolicyOptionResponseArrayOutput() ContentKeyPolicyOptionResponseArrayOutput {
	return o
}

func (o ContentKeyPolicyOptionResponseArrayOutput) ToContentKeyPolicyOptionResponseArrayOutputWithContext(ctx context.Context) ContentKeyPolicyOptionResponseArrayOutput {
	return o
}

func (o ContentKeyPolicyOptionResponseArrayOutput) Index(i pulumi.IntInput) ContentKeyPolicyOptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContentKeyPolicyOptionResponse {
		return vs[0].([]ContentKeyPolicyOptionResponse)[vs[1].(int)]
	}).(ContentKeyPolicyOptionResponseOutput)
}

type ContentKeyPolicyPlayReadyConfiguration struct {
	Licenses           []ContentKeyPolicyPlayReadyLicense `pulumi:"licenses"`
	OdataType          string                             `pulumi:"odataType"`
	ResponseCustomData *string                            `pulumi:"responseCustomData"`
}





type ContentKeyPolicyPlayReadyConfigurationInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyConfigurationOutput() ContentKeyPolicyPlayReadyConfigurationOutput
	ToContentKeyPolicyPlayReadyConfigurationOutputWithContext(context.Context) ContentKeyPolicyPlayReadyConfigurationOutput
}

type ContentKeyPolicyPlayReadyConfigurationArgs struct {
	Licenses           ContentKeyPolicyPlayReadyLicenseArrayInput `pulumi:"licenses"`
	OdataType          pulumi.StringInput                         `pulumi:"odataType"`
	ResponseCustomData pulumi.StringPtrInput                      `pulumi:"responseCustomData"`
}

func (ContentKeyPolicyPlayReadyConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyConfiguration)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyConfigurationArgs) ToContentKeyPolicyPlayReadyConfigurationOutput() ContentKeyPolicyPlayReadyConfigurationOutput {
	return i.ToContentKeyPolicyPlayReadyConfigurationOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyConfigurationArgs) ToContentKeyPolicyPlayReadyConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyConfigurationOutput)
}

type ContentKeyPolicyPlayReadyConfigurationOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyConfiguration)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyConfigurationOutput) ToContentKeyPolicyPlayReadyConfigurationOutput() ContentKeyPolicyPlayReadyConfigurationOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyConfigurationOutput) ToContentKeyPolicyPlayReadyConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyConfigurationOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyConfigurationOutput) Licenses() ContentKeyPolicyPlayReadyLicenseArrayOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyConfiguration) []ContentKeyPolicyPlayReadyLicense { return v.Licenses }).(ContentKeyPolicyPlayReadyLicenseArrayOutput)
}

func (o ContentKeyPolicyPlayReadyConfigurationOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyConfiguration) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyConfigurationOutput) ResponseCustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyConfiguration) *string { return v.ResponseCustomData }).(pulumi.StringPtrOutput)
}

type ContentKeyPolicyPlayReadyConfigurationResponse struct {
	Licenses           []ContentKeyPolicyPlayReadyLicenseResponse `pulumi:"licenses"`
	OdataType          string                                     `pulumi:"odataType"`
	ResponseCustomData *string                                    `pulumi:"responseCustomData"`
}





type ContentKeyPolicyPlayReadyConfigurationResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyConfigurationResponseOutput() ContentKeyPolicyPlayReadyConfigurationResponseOutput
	ToContentKeyPolicyPlayReadyConfigurationResponseOutputWithContext(context.Context) ContentKeyPolicyPlayReadyConfigurationResponseOutput
}

type ContentKeyPolicyPlayReadyConfigurationResponseArgs struct {
	Licenses           ContentKeyPolicyPlayReadyLicenseResponseArrayInput `pulumi:"licenses"`
	OdataType          pulumi.StringInput                                 `pulumi:"odataType"`
	ResponseCustomData pulumi.StringPtrInput                              `pulumi:"responseCustomData"`
}

func (ContentKeyPolicyPlayReadyConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyConfigurationResponse)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyConfigurationResponseArgs) ToContentKeyPolicyPlayReadyConfigurationResponseOutput() ContentKeyPolicyPlayReadyConfigurationResponseOutput {
	return i.ToContentKeyPolicyPlayReadyConfigurationResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyConfigurationResponseArgs) ToContentKeyPolicyPlayReadyConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyConfigurationResponseOutput)
}

type ContentKeyPolicyPlayReadyConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyConfigurationResponse)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyConfigurationResponseOutput) ToContentKeyPolicyPlayReadyConfigurationResponseOutput() ContentKeyPolicyPlayReadyConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyConfigurationResponseOutput) ToContentKeyPolicyPlayReadyConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyConfigurationResponseOutput) Licenses() ContentKeyPolicyPlayReadyLicenseResponseArrayOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyConfigurationResponse) []ContentKeyPolicyPlayReadyLicenseResponse {
		return v.Licenses
	}).(ContentKeyPolicyPlayReadyLicenseResponseArrayOutput)
}

func (o ContentKeyPolicyPlayReadyConfigurationResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyConfigurationResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyConfigurationResponseOutput) ResponseCustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyConfigurationResponse) *string { return v.ResponseCustomData }).(pulumi.StringPtrOutput)
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader struct {
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput
	ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutputWithContext(context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderArgs) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput {
	return i.ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderArgs) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput)
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponse struct {
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput
	ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutputWithContext(context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponse)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseArgs) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput {
	return i.ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseArgs) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput)
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponse)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier struct {
	KeyId     string `pulumi:"keyId"`
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput
	ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutputWithContext(context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierArgs struct {
	KeyId     pulumi.StringInput `pulumi:"keyId"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierArgs) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput {
	return i.ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierArgs) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput)
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier) string { return v.KeyId }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponse struct {
	KeyId     string `pulumi:"keyId"`
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput
	ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutputWithContext(context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseArgs struct {
	KeyId     pulumi.StringInput `pulumi:"keyId"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponse)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseArgs) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput {
	return i.ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseArgs) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput)
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponse)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput() ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput) ToContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponse) string { return v.KeyId }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponse) string {
		return v.OdataType
	}).(pulumi.StringOutput)
}

type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction struct {
	BestEffort        bool `pulumi:"bestEffort"`
	ConfigurationData int  `pulumi:"configurationData"`
}





type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput
	ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutputWithContext(context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput
}

type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionArgs struct {
	BestEffort        pulumi.BoolInput `pulumi:"bestEffort"`
	ConfigurationData pulumi.IntInput  `pulumi:"configurationData"`
}

func (ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionArgs) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput {
	return i.ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionArgs) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput)
}

func (i ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionArgs) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput {
	return i.ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionArgs) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput).ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutputWithContext(ctx)
}









type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput
	ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutputWithContext(context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput
}

type contentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrType ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionArgs

func ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtr(v *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionArgs) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrInput {
	return (*contentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrType)(v)
}

func (*contentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction)(nil)).Elem()
}

func (i *contentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrType) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput {
	return i.ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutputWithContext(context.Background())
}

func (i *contentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrType) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput)
}

type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput {
	return o.ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction) *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction {
		return &v
	}).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput)
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput) BestEffort() pulumi.BoolOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction) bool { return v.BestEffort }).(pulumi.BoolOutput)
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput) ConfigurationData() pulumi.IntOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction) int { return v.ConfigurationData }).(pulumi.IntOutput)
}

type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput) Elem() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction {
		if v != nil {
			return *v
		}
		var ret ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction
		return ret
	}).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput)
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput) BestEffort() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction) *bool {
		if v == nil {
			return nil
		}
		return &v.BestEffort
	}).(pulumi.BoolPtrOutput)
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput) ConfigurationData() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction) *int {
		if v == nil {
			return nil
		}
		return &v.ConfigurationData
	}).(pulumi.IntPtrOutput)
}

type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse struct {
	BestEffort        bool `pulumi:"bestEffort"`
	ConfigurationData int  `pulumi:"configurationData"`
}





type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput
	ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutputWithContext(context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput
}

type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseArgs struct {
	BestEffort        pulumi.BoolInput `pulumi:"bestEffort"`
	ConfigurationData pulumi.IntInput  `pulumi:"configurationData"`
}

func (ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseArgs) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput {
	return i.ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseArgs) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput)
}

func (i ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseArgs) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput {
	return i.ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseArgs) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput).ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutputWithContext(ctx)
}









type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput
	ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutputWithContext(context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput
}

type contentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrType ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseArgs

func ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtr(v *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseArgs) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrInput {
	return (*contentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrType)(v)
}

func (*contentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse)(nil)).Elem()
}

func (i *contentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrType) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput {
	return i.ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutputWithContext(context.Background())
}

func (i *contentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrType) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput)
}

type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput {
	return o.ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse) *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse {
		return &v
	}).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput)
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput) BestEffort() pulumi.BoolOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse) bool { return v.BestEffort }).(pulumi.BoolOutput)
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput) ConfigurationData() pulumi.IntOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse) int {
		return v.ConfigurationData
	}).(pulumi.IntOutput)
}

type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput) ToContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput) Elem() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse) ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse {
		if v != nil {
			return *v
		}
		var ret ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse
		return ret
	}).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput)
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput) BestEffort() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.BestEffort
	}).(pulumi.BoolPtrOutput)
}

func (o ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput) ConfigurationData() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ConfigurationData
	}).(pulumi.IntPtrOutput)
}

type ContentKeyPolicyPlayReadyLicense struct {
	AllowTestDevices       bool                                `pulumi:"allowTestDevices"`
	BeginDate              *string                             `pulumi:"beginDate"`
	ContentKeyLocation     interface{}                         `pulumi:"contentKeyLocation"`
	ContentType            string                              `pulumi:"contentType"`
	ExpirationDate         *string                             `pulumi:"expirationDate"`
	GracePeriod            *string                             `pulumi:"gracePeriod"`
	LicenseType            string                              `pulumi:"licenseType"`
	PlayRight              *ContentKeyPolicyPlayReadyPlayRight `pulumi:"playRight"`
	RelativeBeginDate      *string                             `pulumi:"relativeBeginDate"`
	RelativeExpirationDate *string                             `pulumi:"relativeExpirationDate"`
}





type ContentKeyPolicyPlayReadyLicenseInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyLicenseOutput() ContentKeyPolicyPlayReadyLicenseOutput
	ToContentKeyPolicyPlayReadyLicenseOutputWithContext(context.Context) ContentKeyPolicyPlayReadyLicenseOutput
}

type ContentKeyPolicyPlayReadyLicenseArgs struct {
	AllowTestDevices       pulumi.BoolInput                           `pulumi:"allowTestDevices"`
	BeginDate              pulumi.StringPtrInput                      `pulumi:"beginDate"`
	ContentKeyLocation     pulumi.Input                               `pulumi:"contentKeyLocation"`
	ContentType            pulumi.StringInput                         `pulumi:"contentType"`
	ExpirationDate         pulumi.StringPtrInput                      `pulumi:"expirationDate"`
	GracePeriod            pulumi.StringPtrInput                      `pulumi:"gracePeriod"`
	LicenseType            pulumi.StringInput                         `pulumi:"licenseType"`
	PlayRight              ContentKeyPolicyPlayReadyPlayRightPtrInput `pulumi:"playRight"`
	RelativeBeginDate      pulumi.StringPtrInput                      `pulumi:"relativeBeginDate"`
	RelativeExpirationDate pulumi.StringPtrInput                      `pulumi:"relativeExpirationDate"`
}

func (ContentKeyPolicyPlayReadyLicenseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyLicense)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyLicenseArgs) ToContentKeyPolicyPlayReadyLicenseOutput() ContentKeyPolicyPlayReadyLicenseOutput {
	return i.ToContentKeyPolicyPlayReadyLicenseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyLicenseArgs) ToContentKeyPolicyPlayReadyLicenseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyLicenseOutput)
}





type ContentKeyPolicyPlayReadyLicenseArrayInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyLicenseArrayOutput() ContentKeyPolicyPlayReadyLicenseArrayOutput
	ToContentKeyPolicyPlayReadyLicenseArrayOutputWithContext(context.Context) ContentKeyPolicyPlayReadyLicenseArrayOutput
}

type ContentKeyPolicyPlayReadyLicenseArray []ContentKeyPolicyPlayReadyLicenseInput

func (ContentKeyPolicyPlayReadyLicenseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyPlayReadyLicense)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyLicenseArray) ToContentKeyPolicyPlayReadyLicenseArrayOutput() ContentKeyPolicyPlayReadyLicenseArrayOutput {
	return i.ToContentKeyPolicyPlayReadyLicenseArrayOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyLicenseArray) ToContentKeyPolicyPlayReadyLicenseArrayOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyLicenseArrayOutput)
}

type ContentKeyPolicyPlayReadyLicenseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyLicenseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyLicense)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) ToContentKeyPolicyPlayReadyLicenseOutput() ContentKeyPolicyPlayReadyLicenseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) ToContentKeyPolicyPlayReadyLicenseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) AllowTestDevices() pulumi.BoolOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicense) bool { return v.AllowTestDevices }).(pulumi.BoolOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) BeginDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicense) *string { return v.BeginDate }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) ContentKeyLocation() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicense) interface{} { return v.ContentKeyLocation }).(pulumi.AnyOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicense) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicense) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) GracePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicense) *string { return v.GracePeriod }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicense) string { return v.LicenseType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) PlayRight() ContentKeyPolicyPlayReadyPlayRightPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicense) *ContentKeyPolicyPlayReadyPlayRight { return v.PlayRight }).(ContentKeyPolicyPlayReadyPlayRightPtrOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) RelativeBeginDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicense) *string { return v.RelativeBeginDate }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseOutput) RelativeExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicense) *string { return v.RelativeExpirationDate }).(pulumi.StringPtrOutput)
}

type ContentKeyPolicyPlayReadyLicenseArrayOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyLicenseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyPlayReadyLicense)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyLicenseArrayOutput) ToContentKeyPolicyPlayReadyLicenseArrayOutput() ContentKeyPolicyPlayReadyLicenseArrayOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseArrayOutput) ToContentKeyPolicyPlayReadyLicenseArrayOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseArrayOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseArrayOutput) Index(i pulumi.IntInput) ContentKeyPolicyPlayReadyLicenseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContentKeyPolicyPlayReadyLicense {
		return vs[0].([]ContentKeyPolicyPlayReadyLicense)[vs[1].(int)]
	}).(ContentKeyPolicyPlayReadyLicenseOutput)
}

type ContentKeyPolicyPlayReadyLicenseResponse struct {
	AllowTestDevices       bool                                        `pulumi:"allowTestDevices"`
	BeginDate              *string                                     `pulumi:"beginDate"`
	ContentKeyLocation     interface{}                                 `pulumi:"contentKeyLocation"`
	ContentType            string                                      `pulumi:"contentType"`
	ExpirationDate         *string                                     `pulumi:"expirationDate"`
	GracePeriod            *string                                     `pulumi:"gracePeriod"`
	LicenseType            string                                      `pulumi:"licenseType"`
	PlayRight              *ContentKeyPolicyPlayReadyPlayRightResponse `pulumi:"playRight"`
	RelativeBeginDate      *string                                     `pulumi:"relativeBeginDate"`
	RelativeExpirationDate *string                                     `pulumi:"relativeExpirationDate"`
}





type ContentKeyPolicyPlayReadyLicenseResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyLicenseResponseOutput() ContentKeyPolicyPlayReadyLicenseResponseOutput
	ToContentKeyPolicyPlayReadyLicenseResponseOutputWithContext(context.Context) ContentKeyPolicyPlayReadyLicenseResponseOutput
}

type ContentKeyPolicyPlayReadyLicenseResponseArgs struct {
	AllowTestDevices       pulumi.BoolInput                                   `pulumi:"allowTestDevices"`
	BeginDate              pulumi.StringPtrInput                              `pulumi:"beginDate"`
	ContentKeyLocation     pulumi.Input                                       `pulumi:"contentKeyLocation"`
	ContentType            pulumi.StringInput                                 `pulumi:"contentType"`
	ExpirationDate         pulumi.StringPtrInput                              `pulumi:"expirationDate"`
	GracePeriod            pulumi.StringPtrInput                              `pulumi:"gracePeriod"`
	LicenseType            pulumi.StringInput                                 `pulumi:"licenseType"`
	PlayRight              ContentKeyPolicyPlayReadyPlayRightResponsePtrInput `pulumi:"playRight"`
	RelativeBeginDate      pulumi.StringPtrInput                              `pulumi:"relativeBeginDate"`
	RelativeExpirationDate pulumi.StringPtrInput                              `pulumi:"relativeExpirationDate"`
}

func (ContentKeyPolicyPlayReadyLicenseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyLicenseResponse)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyLicenseResponseArgs) ToContentKeyPolicyPlayReadyLicenseResponseOutput() ContentKeyPolicyPlayReadyLicenseResponseOutput {
	return i.ToContentKeyPolicyPlayReadyLicenseResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyLicenseResponseArgs) ToContentKeyPolicyPlayReadyLicenseResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyLicenseResponseOutput)
}





type ContentKeyPolicyPlayReadyLicenseResponseArrayInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyLicenseResponseArrayOutput() ContentKeyPolicyPlayReadyLicenseResponseArrayOutput
	ToContentKeyPolicyPlayReadyLicenseResponseArrayOutputWithContext(context.Context) ContentKeyPolicyPlayReadyLicenseResponseArrayOutput
}

type ContentKeyPolicyPlayReadyLicenseResponseArray []ContentKeyPolicyPlayReadyLicenseResponseInput

func (ContentKeyPolicyPlayReadyLicenseResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyPlayReadyLicenseResponse)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyLicenseResponseArray) ToContentKeyPolicyPlayReadyLicenseResponseArrayOutput() ContentKeyPolicyPlayReadyLicenseResponseArrayOutput {
	return i.ToContentKeyPolicyPlayReadyLicenseResponseArrayOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyLicenseResponseArray) ToContentKeyPolicyPlayReadyLicenseResponseArrayOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyLicenseResponseArrayOutput)
}

type ContentKeyPolicyPlayReadyLicenseResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyLicenseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyLicenseResponse)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) ToContentKeyPolicyPlayReadyLicenseResponseOutput() ContentKeyPolicyPlayReadyLicenseResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) ToContentKeyPolicyPlayReadyLicenseResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) AllowTestDevices() pulumi.BoolOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicenseResponse) bool { return v.AllowTestDevices }).(pulumi.BoolOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) BeginDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicenseResponse) *string { return v.BeginDate }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) ContentKeyLocation() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicenseResponse) interface{} { return v.ContentKeyLocation }).(pulumi.AnyOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicenseResponse) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicenseResponse) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) GracePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicenseResponse) *string { return v.GracePeriod }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicenseResponse) string { return v.LicenseType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) PlayRight() ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicenseResponse) *ContentKeyPolicyPlayReadyPlayRightResponse {
		return v.PlayRight
	}).(ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) RelativeBeginDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicenseResponse) *string { return v.RelativeBeginDate }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseResponseOutput) RelativeExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyLicenseResponse) *string { return v.RelativeExpirationDate }).(pulumi.StringPtrOutput)
}

type ContentKeyPolicyPlayReadyLicenseResponseArrayOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyLicenseResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyPlayReadyLicenseResponse)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyLicenseResponseArrayOutput) ToContentKeyPolicyPlayReadyLicenseResponseArrayOutput() ContentKeyPolicyPlayReadyLicenseResponseArrayOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseResponseArrayOutput) ToContentKeyPolicyPlayReadyLicenseResponseArrayOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseResponseArrayOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseResponseArrayOutput) Index(i pulumi.IntInput) ContentKeyPolicyPlayReadyLicenseResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContentKeyPolicyPlayReadyLicenseResponse {
		return vs[0].([]ContentKeyPolicyPlayReadyLicenseResponse)[vs[1].(int)]
	}).(ContentKeyPolicyPlayReadyLicenseResponseOutput)
}

type ContentKeyPolicyPlayReadyPlayRight struct {
	AgcAndColorStripeRestriction                       *int                                                          `pulumi:"agcAndColorStripeRestriction"`
	AllowPassingVideoContentToUnknownOutput            string                                                        `pulumi:"allowPassingVideoContentToUnknownOutput"`
	AnalogVideoOpl                                     *int                                                          `pulumi:"analogVideoOpl"`
	CompressedDigitalAudioOpl                          *int                                                          `pulumi:"compressedDigitalAudioOpl"`
	CompressedDigitalVideoOpl                          *int                                                          `pulumi:"compressedDigitalVideoOpl"`
	DigitalVideoOnlyContentRestriction                 bool                                                          `pulumi:"digitalVideoOnlyContentRestriction"`
	ExplicitAnalogTelevisionOutputRestriction          *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction `pulumi:"explicitAnalogTelevisionOutputRestriction"`
	FirstPlayExpiration                                *string                                                       `pulumi:"firstPlayExpiration"`
	ImageConstraintForAnalogComponentVideoRestriction  bool                                                          `pulumi:"imageConstraintForAnalogComponentVideoRestriction"`
	ImageConstraintForAnalogComputerMonitorRestriction bool                                                          `pulumi:"imageConstraintForAnalogComputerMonitorRestriction"`
	ScmsRestriction                                    *int                                                          `pulumi:"scmsRestriction"`
	UncompressedDigitalAudioOpl                        *int                                                          `pulumi:"uncompressedDigitalAudioOpl"`
	UncompressedDigitalVideoOpl                        *int                                                          `pulumi:"uncompressedDigitalVideoOpl"`
}





type ContentKeyPolicyPlayReadyPlayRightInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyPlayRightOutput() ContentKeyPolicyPlayReadyPlayRightOutput
	ToContentKeyPolicyPlayReadyPlayRightOutputWithContext(context.Context) ContentKeyPolicyPlayReadyPlayRightOutput
}

type ContentKeyPolicyPlayReadyPlayRightArgs struct {
	AgcAndColorStripeRestriction                       pulumi.IntPtrInput                                                   `pulumi:"agcAndColorStripeRestriction"`
	AllowPassingVideoContentToUnknownOutput            pulumi.StringInput                                                   `pulumi:"allowPassingVideoContentToUnknownOutput"`
	AnalogVideoOpl                                     pulumi.IntPtrInput                                                   `pulumi:"analogVideoOpl"`
	CompressedDigitalAudioOpl                          pulumi.IntPtrInput                                                   `pulumi:"compressedDigitalAudioOpl"`
	CompressedDigitalVideoOpl                          pulumi.IntPtrInput                                                   `pulumi:"compressedDigitalVideoOpl"`
	DigitalVideoOnlyContentRestriction                 pulumi.BoolInput                                                     `pulumi:"digitalVideoOnlyContentRestriction"`
	ExplicitAnalogTelevisionOutputRestriction          ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrInput `pulumi:"explicitAnalogTelevisionOutputRestriction"`
	FirstPlayExpiration                                pulumi.StringPtrInput                                                `pulumi:"firstPlayExpiration"`
	ImageConstraintForAnalogComponentVideoRestriction  pulumi.BoolInput                                                     `pulumi:"imageConstraintForAnalogComponentVideoRestriction"`
	ImageConstraintForAnalogComputerMonitorRestriction pulumi.BoolInput                                                     `pulumi:"imageConstraintForAnalogComputerMonitorRestriction"`
	ScmsRestriction                                    pulumi.IntPtrInput                                                   `pulumi:"scmsRestriction"`
	UncompressedDigitalAudioOpl                        pulumi.IntPtrInput                                                   `pulumi:"uncompressedDigitalAudioOpl"`
	UncompressedDigitalVideoOpl                        pulumi.IntPtrInput                                                   `pulumi:"uncompressedDigitalVideoOpl"`
}

func (ContentKeyPolicyPlayReadyPlayRightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyPlayRight)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyPlayRightArgs) ToContentKeyPolicyPlayReadyPlayRightOutput() ContentKeyPolicyPlayReadyPlayRightOutput {
	return i.ToContentKeyPolicyPlayReadyPlayRightOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyPlayRightArgs) ToContentKeyPolicyPlayReadyPlayRightOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyPlayRightOutput)
}

func (i ContentKeyPolicyPlayReadyPlayRightArgs) ToContentKeyPolicyPlayReadyPlayRightPtrOutput() ContentKeyPolicyPlayReadyPlayRightPtrOutput {
	return i.ToContentKeyPolicyPlayReadyPlayRightPtrOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyPlayRightArgs) ToContentKeyPolicyPlayReadyPlayRightPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyPlayRightOutput).ToContentKeyPolicyPlayReadyPlayRightPtrOutputWithContext(ctx)
}









type ContentKeyPolicyPlayReadyPlayRightPtrInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyPlayRightPtrOutput() ContentKeyPolicyPlayReadyPlayRightPtrOutput
	ToContentKeyPolicyPlayReadyPlayRightPtrOutputWithContext(context.Context) ContentKeyPolicyPlayReadyPlayRightPtrOutput
}

type contentKeyPolicyPlayReadyPlayRightPtrType ContentKeyPolicyPlayReadyPlayRightArgs

func ContentKeyPolicyPlayReadyPlayRightPtr(v *ContentKeyPolicyPlayReadyPlayRightArgs) ContentKeyPolicyPlayReadyPlayRightPtrInput {
	return (*contentKeyPolicyPlayReadyPlayRightPtrType)(v)
}

func (*contentKeyPolicyPlayReadyPlayRightPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyPlayReadyPlayRight)(nil)).Elem()
}

func (i *contentKeyPolicyPlayReadyPlayRightPtrType) ToContentKeyPolicyPlayReadyPlayRightPtrOutput() ContentKeyPolicyPlayReadyPlayRightPtrOutput {
	return i.ToContentKeyPolicyPlayReadyPlayRightPtrOutputWithContext(context.Background())
}

func (i *contentKeyPolicyPlayReadyPlayRightPtrType) ToContentKeyPolicyPlayReadyPlayRightPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyPlayRightPtrOutput)
}

type ContentKeyPolicyPlayReadyPlayRightOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyPlayRightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyPlayRight)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) ToContentKeyPolicyPlayReadyPlayRightOutput() ContentKeyPolicyPlayReadyPlayRightOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) ToContentKeyPolicyPlayReadyPlayRightOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) ToContentKeyPolicyPlayReadyPlayRightPtrOutput() ContentKeyPolicyPlayReadyPlayRightPtrOutput {
	return o.ToContentKeyPolicyPlayReadyPlayRightPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) ToContentKeyPolicyPlayReadyPlayRightPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentKeyPolicyPlayReadyPlayRight) *ContentKeyPolicyPlayReadyPlayRight {
		return &v
	}).(ContentKeyPolicyPlayReadyPlayRightPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) AgcAndColorStripeRestriction() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) *int { return v.AgcAndColorStripeRestriction }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) AllowPassingVideoContentToUnknownOutput() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) string { return v.AllowPassingVideoContentToUnknownOutput }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) AnalogVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) *int { return v.AnalogVideoOpl }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) CompressedDigitalAudioOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) *int { return v.CompressedDigitalAudioOpl }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) CompressedDigitalVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) *int { return v.CompressedDigitalVideoOpl }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) DigitalVideoOnlyContentRestriction() pulumi.BoolOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) bool { return v.DigitalVideoOnlyContentRestriction }).(pulumi.BoolOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) ExplicitAnalogTelevisionOutputRestriction() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction {
		return v.ExplicitAnalogTelevisionOutputRestriction
	}).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) FirstPlayExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) *string { return v.FirstPlayExpiration }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) ImageConstraintForAnalogComponentVideoRestriction() pulumi.BoolOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) bool {
		return v.ImageConstraintForAnalogComponentVideoRestriction
	}).(pulumi.BoolOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) ImageConstraintForAnalogComputerMonitorRestriction() pulumi.BoolOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) bool {
		return v.ImageConstraintForAnalogComputerMonitorRestriction
	}).(pulumi.BoolOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) ScmsRestriction() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) *int { return v.ScmsRestriction }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) UncompressedDigitalAudioOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) *int { return v.UncompressedDigitalAudioOpl }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightOutput) UncompressedDigitalVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRight) *int { return v.UncompressedDigitalVideoOpl }).(pulumi.IntPtrOutput)
}

type ContentKeyPolicyPlayReadyPlayRightPtrOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyPlayRightPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyPlayReadyPlayRight)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) ToContentKeyPolicyPlayReadyPlayRightPtrOutput() ContentKeyPolicyPlayReadyPlayRightPtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) ToContentKeyPolicyPlayReadyPlayRightPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightPtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) Elem() ContentKeyPolicyPlayReadyPlayRightOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) ContentKeyPolicyPlayReadyPlayRight {
		if v != nil {
			return *v
		}
		var ret ContentKeyPolicyPlayReadyPlayRight
		return ret
	}).(ContentKeyPolicyPlayReadyPlayRightOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) AgcAndColorStripeRestriction() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *int {
		if v == nil {
			return nil
		}
		return v.AgcAndColorStripeRestriction
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) AllowPassingVideoContentToUnknownOutput() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *string {
		if v == nil {
			return nil
		}
		return &v.AllowPassingVideoContentToUnknownOutput
	}).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) AnalogVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *int {
		if v == nil {
			return nil
		}
		return v.AnalogVideoOpl
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) CompressedDigitalAudioOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *int {
		if v == nil {
			return nil
		}
		return v.CompressedDigitalAudioOpl
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) CompressedDigitalVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *int {
		if v == nil {
			return nil
		}
		return v.CompressedDigitalVideoOpl
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) DigitalVideoOnlyContentRestriction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *bool {
		if v == nil {
			return nil
		}
		return &v.DigitalVideoOnlyContentRestriction
	}).(pulumi.BoolPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) ExplicitAnalogTelevisionOutputRestriction() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction {
		if v == nil {
			return nil
		}
		return v.ExplicitAnalogTelevisionOutputRestriction
	}).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) FirstPlayExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *string {
		if v == nil {
			return nil
		}
		return v.FirstPlayExpiration
	}).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) ImageConstraintForAnalogComponentVideoRestriction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *bool {
		if v == nil {
			return nil
		}
		return &v.ImageConstraintForAnalogComponentVideoRestriction
	}).(pulumi.BoolPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) ImageConstraintForAnalogComputerMonitorRestriction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *bool {
		if v == nil {
			return nil
		}
		return &v.ImageConstraintForAnalogComputerMonitorRestriction
	}).(pulumi.BoolPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) ScmsRestriction() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *int {
		if v == nil {
			return nil
		}
		return v.ScmsRestriction
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) UncompressedDigitalAudioOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *int {
		if v == nil {
			return nil
		}
		return v.UncompressedDigitalAudioOpl
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightPtrOutput) UncompressedDigitalVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRight) *int {
		if v == nil {
			return nil
		}
		return v.UncompressedDigitalVideoOpl
	}).(pulumi.IntPtrOutput)
}

type ContentKeyPolicyPlayReadyPlayRightResponse struct {
	AgcAndColorStripeRestriction                       *int                                                                  `pulumi:"agcAndColorStripeRestriction"`
	AllowPassingVideoContentToUnknownOutput            string                                                                `pulumi:"allowPassingVideoContentToUnknownOutput"`
	AnalogVideoOpl                                     *int                                                                  `pulumi:"analogVideoOpl"`
	CompressedDigitalAudioOpl                          *int                                                                  `pulumi:"compressedDigitalAudioOpl"`
	CompressedDigitalVideoOpl                          *int                                                                  `pulumi:"compressedDigitalVideoOpl"`
	DigitalVideoOnlyContentRestriction                 bool                                                                  `pulumi:"digitalVideoOnlyContentRestriction"`
	ExplicitAnalogTelevisionOutputRestriction          *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse `pulumi:"explicitAnalogTelevisionOutputRestriction"`
	FirstPlayExpiration                                *string                                                               `pulumi:"firstPlayExpiration"`
	ImageConstraintForAnalogComponentVideoRestriction  bool                                                                  `pulumi:"imageConstraintForAnalogComponentVideoRestriction"`
	ImageConstraintForAnalogComputerMonitorRestriction bool                                                                  `pulumi:"imageConstraintForAnalogComputerMonitorRestriction"`
	ScmsRestriction                                    *int                                                                  `pulumi:"scmsRestriction"`
	UncompressedDigitalAudioOpl                        *int                                                                  `pulumi:"uncompressedDigitalAudioOpl"`
	UncompressedDigitalVideoOpl                        *int                                                                  `pulumi:"uncompressedDigitalVideoOpl"`
}





type ContentKeyPolicyPlayReadyPlayRightResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyPlayRightResponseOutput() ContentKeyPolicyPlayReadyPlayRightResponseOutput
	ToContentKeyPolicyPlayReadyPlayRightResponseOutputWithContext(context.Context) ContentKeyPolicyPlayReadyPlayRightResponseOutput
}

type ContentKeyPolicyPlayReadyPlayRightResponseArgs struct {
	AgcAndColorStripeRestriction                       pulumi.IntPtrInput                                                           `pulumi:"agcAndColorStripeRestriction"`
	AllowPassingVideoContentToUnknownOutput            pulumi.StringInput                                                           `pulumi:"allowPassingVideoContentToUnknownOutput"`
	AnalogVideoOpl                                     pulumi.IntPtrInput                                                           `pulumi:"analogVideoOpl"`
	CompressedDigitalAudioOpl                          pulumi.IntPtrInput                                                           `pulumi:"compressedDigitalAudioOpl"`
	CompressedDigitalVideoOpl                          pulumi.IntPtrInput                                                           `pulumi:"compressedDigitalVideoOpl"`
	DigitalVideoOnlyContentRestriction                 pulumi.BoolInput                                                             `pulumi:"digitalVideoOnlyContentRestriction"`
	ExplicitAnalogTelevisionOutputRestriction          ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrInput `pulumi:"explicitAnalogTelevisionOutputRestriction"`
	FirstPlayExpiration                                pulumi.StringPtrInput                                                        `pulumi:"firstPlayExpiration"`
	ImageConstraintForAnalogComponentVideoRestriction  pulumi.BoolInput                                                             `pulumi:"imageConstraintForAnalogComponentVideoRestriction"`
	ImageConstraintForAnalogComputerMonitorRestriction pulumi.BoolInput                                                             `pulumi:"imageConstraintForAnalogComputerMonitorRestriction"`
	ScmsRestriction                                    pulumi.IntPtrInput                                                           `pulumi:"scmsRestriction"`
	UncompressedDigitalAudioOpl                        pulumi.IntPtrInput                                                           `pulumi:"uncompressedDigitalAudioOpl"`
	UncompressedDigitalVideoOpl                        pulumi.IntPtrInput                                                           `pulumi:"uncompressedDigitalVideoOpl"`
}

func (ContentKeyPolicyPlayReadyPlayRightResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyPlayRightResponse)(nil)).Elem()
}

func (i ContentKeyPolicyPlayReadyPlayRightResponseArgs) ToContentKeyPolicyPlayReadyPlayRightResponseOutput() ContentKeyPolicyPlayReadyPlayRightResponseOutput {
	return i.ToContentKeyPolicyPlayReadyPlayRightResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyPlayRightResponseArgs) ToContentKeyPolicyPlayReadyPlayRightResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyPlayRightResponseOutput)
}

func (i ContentKeyPolicyPlayReadyPlayRightResponseArgs) ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutput() ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput {
	return i.ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutputWithContext(context.Background())
}

func (i ContentKeyPolicyPlayReadyPlayRightResponseArgs) ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyPlayRightResponseOutput).ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutputWithContext(ctx)
}









type ContentKeyPolicyPlayReadyPlayRightResponsePtrInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutput() ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput
	ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutputWithContext(context.Context) ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput
}

type contentKeyPolicyPlayReadyPlayRightResponsePtrType ContentKeyPolicyPlayReadyPlayRightResponseArgs

func ContentKeyPolicyPlayReadyPlayRightResponsePtr(v *ContentKeyPolicyPlayReadyPlayRightResponseArgs) ContentKeyPolicyPlayReadyPlayRightResponsePtrInput {
	return (*contentKeyPolicyPlayReadyPlayRightResponsePtrType)(v)
}

func (*contentKeyPolicyPlayReadyPlayRightResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyPlayReadyPlayRightResponse)(nil)).Elem()
}

func (i *contentKeyPolicyPlayReadyPlayRightResponsePtrType) ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutput() ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput {
	return i.ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutputWithContext(context.Background())
}

func (i *contentKeyPolicyPlayReadyPlayRightResponsePtrType) ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput)
}

type ContentKeyPolicyPlayReadyPlayRightResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyPlayRightResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyPlayRightResponse)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) ToContentKeyPolicyPlayReadyPlayRightResponseOutput() ContentKeyPolicyPlayReadyPlayRightResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) ToContentKeyPolicyPlayReadyPlayRightResponseOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightResponseOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutput() ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput {
	return o.ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentKeyPolicyPlayReadyPlayRightResponse) *ContentKeyPolicyPlayReadyPlayRightResponse {
		return &v
	}).(ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) AgcAndColorStripeRestriction() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) *int { return v.AgcAndColorStripeRestriction }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) AllowPassingVideoContentToUnknownOutput() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) string {
		return v.AllowPassingVideoContentToUnknownOutput
	}).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) AnalogVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) *int { return v.AnalogVideoOpl }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) CompressedDigitalAudioOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) *int { return v.CompressedDigitalAudioOpl }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) CompressedDigitalVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) *int { return v.CompressedDigitalVideoOpl }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) DigitalVideoOnlyContentRestriction() pulumi.BoolOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) bool { return v.DigitalVideoOnlyContentRestriction }).(pulumi.BoolOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) ExplicitAnalogTelevisionOutputRestriction() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse {
		return v.ExplicitAnalogTelevisionOutputRestriction
	}).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) FirstPlayExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) *string { return v.FirstPlayExpiration }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) ImageConstraintForAnalogComponentVideoRestriction() pulumi.BoolOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) bool {
		return v.ImageConstraintForAnalogComponentVideoRestriction
	}).(pulumi.BoolOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) ImageConstraintForAnalogComputerMonitorRestriction() pulumi.BoolOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) bool {
		return v.ImageConstraintForAnalogComputerMonitorRestriction
	}).(pulumi.BoolOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) ScmsRestriction() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) *int { return v.ScmsRestriction }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) UncompressedDigitalAudioOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) *int { return v.UncompressedDigitalAudioOpl }).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponseOutput) UncompressedDigitalVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyPlayReadyPlayRightResponse) *int { return v.UncompressedDigitalVideoOpl }).(pulumi.IntPtrOutput)
}

type ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyPlayReadyPlayRightResponse)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutput() ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) ToContentKeyPolicyPlayReadyPlayRightResponsePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) Elem() ContentKeyPolicyPlayReadyPlayRightResponseOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) ContentKeyPolicyPlayReadyPlayRightResponse {
		if v != nil {
			return *v
		}
		var ret ContentKeyPolicyPlayReadyPlayRightResponse
		return ret
	}).(ContentKeyPolicyPlayReadyPlayRightResponseOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) AgcAndColorStripeRestriction() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *int {
		if v == nil {
			return nil
		}
		return v.AgcAndColorStripeRestriction
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) AllowPassingVideoContentToUnknownOutput() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AllowPassingVideoContentToUnknownOutput
	}).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) AnalogVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *int {
		if v == nil {
			return nil
		}
		return v.AnalogVideoOpl
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) CompressedDigitalAudioOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *int {
		if v == nil {
			return nil
		}
		return v.CompressedDigitalAudioOpl
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) CompressedDigitalVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *int {
		if v == nil {
			return nil
		}
		return v.CompressedDigitalVideoOpl
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) DigitalVideoOnlyContentRestriction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.DigitalVideoOnlyContentRestriction
	}).(pulumi.BoolPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) ExplicitAnalogTelevisionOutputRestriction() ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse {
		if v == nil {
			return nil
		}
		return v.ExplicitAnalogTelevisionOutputRestriction
	}).(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) FirstPlayExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *string {
		if v == nil {
			return nil
		}
		return v.FirstPlayExpiration
	}).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) ImageConstraintForAnalogComponentVideoRestriction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ImageConstraintForAnalogComponentVideoRestriction
	}).(pulumi.BoolPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) ImageConstraintForAnalogComputerMonitorRestriction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ImageConstraintForAnalogComputerMonitorRestriction
	}).(pulumi.BoolPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) ScmsRestriction() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *int {
		if v == nil {
			return nil
		}
		return v.ScmsRestriction
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) UncompressedDigitalAudioOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *int {
		if v == nil {
			return nil
		}
		return v.UncompressedDigitalAudioOpl
	}).(pulumi.IntPtrOutput)
}

func (o ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput) UncompressedDigitalVideoOpl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyPlayRightResponse) *int {
		if v == nil {
			return nil
		}
		return v.UncompressedDigitalVideoOpl
	}).(pulumi.IntPtrOutput)
}

type ContentKeyPolicyRsaTokenKey struct {
	Exponent  string `pulumi:"exponent"`
	Modulus   string `pulumi:"modulus"`
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyRsaTokenKeyInput interface {
	pulumi.Input

	ToContentKeyPolicyRsaTokenKeyOutput() ContentKeyPolicyRsaTokenKeyOutput
	ToContentKeyPolicyRsaTokenKeyOutputWithContext(context.Context) ContentKeyPolicyRsaTokenKeyOutput
}

type ContentKeyPolicyRsaTokenKeyArgs struct {
	Exponent  pulumi.StringInput `pulumi:"exponent"`
	Modulus   pulumi.StringInput `pulumi:"modulus"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyRsaTokenKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyRsaTokenKey)(nil)).Elem()
}

func (i ContentKeyPolicyRsaTokenKeyArgs) ToContentKeyPolicyRsaTokenKeyOutput() ContentKeyPolicyRsaTokenKeyOutput {
	return i.ToContentKeyPolicyRsaTokenKeyOutputWithContext(context.Background())
}

func (i ContentKeyPolicyRsaTokenKeyArgs) ToContentKeyPolicyRsaTokenKeyOutputWithContext(ctx context.Context) ContentKeyPolicyRsaTokenKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyRsaTokenKeyOutput)
}

type ContentKeyPolicyRsaTokenKeyOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyRsaTokenKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyRsaTokenKey)(nil)).Elem()
}

func (o ContentKeyPolicyRsaTokenKeyOutput) ToContentKeyPolicyRsaTokenKeyOutput() ContentKeyPolicyRsaTokenKeyOutput {
	return o
}

func (o ContentKeyPolicyRsaTokenKeyOutput) ToContentKeyPolicyRsaTokenKeyOutputWithContext(ctx context.Context) ContentKeyPolicyRsaTokenKeyOutput {
	return o
}

func (o ContentKeyPolicyRsaTokenKeyOutput) Exponent() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyRsaTokenKey) string { return v.Exponent }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyRsaTokenKeyOutput) Modulus() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyRsaTokenKey) string { return v.Modulus }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyRsaTokenKeyOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyRsaTokenKey) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyRsaTokenKeyResponse struct {
	Exponent  string `pulumi:"exponent"`
	Modulus   string `pulumi:"modulus"`
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyRsaTokenKeyResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyRsaTokenKeyResponseOutput() ContentKeyPolicyRsaTokenKeyResponseOutput
	ToContentKeyPolicyRsaTokenKeyResponseOutputWithContext(context.Context) ContentKeyPolicyRsaTokenKeyResponseOutput
}

type ContentKeyPolicyRsaTokenKeyResponseArgs struct {
	Exponent  pulumi.StringInput `pulumi:"exponent"`
	Modulus   pulumi.StringInput `pulumi:"modulus"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyRsaTokenKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyRsaTokenKeyResponse)(nil)).Elem()
}

func (i ContentKeyPolicyRsaTokenKeyResponseArgs) ToContentKeyPolicyRsaTokenKeyResponseOutput() ContentKeyPolicyRsaTokenKeyResponseOutput {
	return i.ToContentKeyPolicyRsaTokenKeyResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyRsaTokenKeyResponseArgs) ToContentKeyPolicyRsaTokenKeyResponseOutputWithContext(ctx context.Context) ContentKeyPolicyRsaTokenKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyRsaTokenKeyResponseOutput)
}

type ContentKeyPolicyRsaTokenKeyResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyRsaTokenKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyRsaTokenKeyResponse)(nil)).Elem()
}

func (o ContentKeyPolicyRsaTokenKeyResponseOutput) ToContentKeyPolicyRsaTokenKeyResponseOutput() ContentKeyPolicyRsaTokenKeyResponseOutput {
	return o
}

func (o ContentKeyPolicyRsaTokenKeyResponseOutput) ToContentKeyPolicyRsaTokenKeyResponseOutputWithContext(ctx context.Context) ContentKeyPolicyRsaTokenKeyResponseOutput {
	return o
}

func (o ContentKeyPolicyRsaTokenKeyResponseOutput) Exponent() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyRsaTokenKeyResponse) string { return v.Exponent }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyRsaTokenKeyResponseOutput) Modulus() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyRsaTokenKeyResponse) string { return v.Modulus }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyRsaTokenKeyResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyRsaTokenKeyResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicySymmetricTokenKey struct {
	KeyValue  string `pulumi:"keyValue"`
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicySymmetricTokenKeyInput interface {
	pulumi.Input

	ToContentKeyPolicySymmetricTokenKeyOutput() ContentKeyPolicySymmetricTokenKeyOutput
	ToContentKeyPolicySymmetricTokenKeyOutputWithContext(context.Context) ContentKeyPolicySymmetricTokenKeyOutput
}

type ContentKeyPolicySymmetricTokenKeyArgs struct {
	KeyValue  pulumi.StringInput `pulumi:"keyValue"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicySymmetricTokenKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicySymmetricTokenKey)(nil)).Elem()
}

func (i ContentKeyPolicySymmetricTokenKeyArgs) ToContentKeyPolicySymmetricTokenKeyOutput() ContentKeyPolicySymmetricTokenKeyOutput {
	return i.ToContentKeyPolicySymmetricTokenKeyOutputWithContext(context.Background())
}

func (i ContentKeyPolicySymmetricTokenKeyArgs) ToContentKeyPolicySymmetricTokenKeyOutputWithContext(ctx context.Context) ContentKeyPolicySymmetricTokenKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicySymmetricTokenKeyOutput)
}

type ContentKeyPolicySymmetricTokenKeyOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicySymmetricTokenKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicySymmetricTokenKey)(nil)).Elem()
}

func (o ContentKeyPolicySymmetricTokenKeyOutput) ToContentKeyPolicySymmetricTokenKeyOutput() ContentKeyPolicySymmetricTokenKeyOutput {
	return o
}

func (o ContentKeyPolicySymmetricTokenKeyOutput) ToContentKeyPolicySymmetricTokenKeyOutputWithContext(ctx context.Context) ContentKeyPolicySymmetricTokenKeyOutput {
	return o
}

func (o ContentKeyPolicySymmetricTokenKeyOutput) KeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicySymmetricTokenKey) string { return v.KeyValue }).(pulumi.StringOutput)
}

func (o ContentKeyPolicySymmetricTokenKeyOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicySymmetricTokenKey) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicySymmetricTokenKeyResponse struct {
	KeyValue  string `pulumi:"keyValue"`
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicySymmetricTokenKeyResponseInput interface {
	pulumi.Input

	ToContentKeyPolicySymmetricTokenKeyResponseOutput() ContentKeyPolicySymmetricTokenKeyResponseOutput
	ToContentKeyPolicySymmetricTokenKeyResponseOutputWithContext(context.Context) ContentKeyPolicySymmetricTokenKeyResponseOutput
}

type ContentKeyPolicySymmetricTokenKeyResponseArgs struct {
	KeyValue  pulumi.StringInput `pulumi:"keyValue"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicySymmetricTokenKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicySymmetricTokenKeyResponse)(nil)).Elem()
}

func (i ContentKeyPolicySymmetricTokenKeyResponseArgs) ToContentKeyPolicySymmetricTokenKeyResponseOutput() ContentKeyPolicySymmetricTokenKeyResponseOutput {
	return i.ToContentKeyPolicySymmetricTokenKeyResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicySymmetricTokenKeyResponseArgs) ToContentKeyPolicySymmetricTokenKeyResponseOutputWithContext(ctx context.Context) ContentKeyPolicySymmetricTokenKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicySymmetricTokenKeyResponseOutput)
}

type ContentKeyPolicySymmetricTokenKeyResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicySymmetricTokenKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicySymmetricTokenKeyResponse)(nil)).Elem()
}

func (o ContentKeyPolicySymmetricTokenKeyResponseOutput) ToContentKeyPolicySymmetricTokenKeyResponseOutput() ContentKeyPolicySymmetricTokenKeyResponseOutput {
	return o
}

func (o ContentKeyPolicySymmetricTokenKeyResponseOutput) ToContentKeyPolicySymmetricTokenKeyResponseOutputWithContext(ctx context.Context) ContentKeyPolicySymmetricTokenKeyResponseOutput {
	return o
}

func (o ContentKeyPolicySymmetricTokenKeyResponseOutput) KeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicySymmetricTokenKeyResponse) string { return v.KeyValue }).(pulumi.StringOutput)
}

func (o ContentKeyPolicySymmetricTokenKeyResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicySymmetricTokenKeyResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyTokenClaim struct {
	ClaimType  *string `pulumi:"claimType"`
	ClaimValue *string `pulumi:"claimValue"`
}





type ContentKeyPolicyTokenClaimInput interface {
	pulumi.Input

	ToContentKeyPolicyTokenClaimOutput() ContentKeyPolicyTokenClaimOutput
	ToContentKeyPolicyTokenClaimOutputWithContext(context.Context) ContentKeyPolicyTokenClaimOutput
}

type ContentKeyPolicyTokenClaimArgs struct {
	ClaimType  pulumi.StringPtrInput `pulumi:"claimType"`
	ClaimValue pulumi.StringPtrInput `pulumi:"claimValue"`
}

func (ContentKeyPolicyTokenClaimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyTokenClaim)(nil)).Elem()
}

func (i ContentKeyPolicyTokenClaimArgs) ToContentKeyPolicyTokenClaimOutput() ContentKeyPolicyTokenClaimOutput {
	return i.ToContentKeyPolicyTokenClaimOutputWithContext(context.Background())
}

func (i ContentKeyPolicyTokenClaimArgs) ToContentKeyPolicyTokenClaimOutputWithContext(ctx context.Context) ContentKeyPolicyTokenClaimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyTokenClaimOutput)
}





type ContentKeyPolicyTokenClaimArrayInput interface {
	pulumi.Input

	ToContentKeyPolicyTokenClaimArrayOutput() ContentKeyPolicyTokenClaimArrayOutput
	ToContentKeyPolicyTokenClaimArrayOutputWithContext(context.Context) ContentKeyPolicyTokenClaimArrayOutput
}

type ContentKeyPolicyTokenClaimArray []ContentKeyPolicyTokenClaimInput

func (ContentKeyPolicyTokenClaimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyTokenClaim)(nil)).Elem()
}

func (i ContentKeyPolicyTokenClaimArray) ToContentKeyPolicyTokenClaimArrayOutput() ContentKeyPolicyTokenClaimArrayOutput {
	return i.ToContentKeyPolicyTokenClaimArrayOutputWithContext(context.Background())
}

func (i ContentKeyPolicyTokenClaimArray) ToContentKeyPolicyTokenClaimArrayOutputWithContext(ctx context.Context) ContentKeyPolicyTokenClaimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyTokenClaimArrayOutput)
}

type ContentKeyPolicyTokenClaimOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyTokenClaimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyTokenClaim)(nil)).Elem()
}

func (o ContentKeyPolicyTokenClaimOutput) ToContentKeyPolicyTokenClaimOutput() ContentKeyPolicyTokenClaimOutput {
	return o
}

func (o ContentKeyPolicyTokenClaimOutput) ToContentKeyPolicyTokenClaimOutputWithContext(ctx context.Context) ContentKeyPolicyTokenClaimOutput {
	return o
}

func (o ContentKeyPolicyTokenClaimOutput) ClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenClaim) *string { return v.ClaimType }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyTokenClaimOutput) ClaimValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenClaim) *string { return v.ClaimValue }).(pulumi.StringPtrOutput)
}

type ContentKeyPolicyTokenClaimArrayOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyTokenClaimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyTokenClaim)(nil)).Elem()
}

func (o ContentKeyPolicyTokenClaimArrayOutput) ToContentKeyPolicyTokenClaimArrayOutput() ContentKeyPolicyTokenClaimArrayOutput {
	return o
}

func (o ContentKeyPolicyTokenClaimArrayOutput) ToContentKeyPolicyTokenClaimArrayOutputWithContext(ctx context.Context) ContentKeyPolicyTokenClaimArrayOutput {
	return o
}

func (o ContentKeyPolicyTokenClaimArrayOutput) Index(i pulumi.IntInput) ContentKeyPolicyTokenClaimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContentKeyPolicyTokenClaim {
		return vs[0].([]ContentKeyPolicyTokenClaim)[vs[1].(int)]
	}).(ContentKeyPolicyTokenClaimOutput)
}

type ContentKeyPolicyTokenClaimResponse struct {
	ClaimType  *string `pulumi:"claimType"`
	ClaimValue *string `pulumi:"claimValue"`
}





type ContentKeyPolicyTokenClaimResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyTokenClaimResponseOutput() ContentKeyPolicyTokenClaimResponseOutput
	ToContentKeyPolicyTokenClaimResponseOutputWithContext(context.Context) ContentKeyPolicyTokenClaimResponseOutput
}

type ContentKeyPolicyTokenClaimResponseArgs struct {
	ClaimType  pulumi.StringPtrInput `pulumi:"claimType"`
	ClaimValue pulumi.StringPtrInput `pulumi:"claimValue"`
}

func (ContentKeyPolicyTokenClaimResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyTokenClaimResponse)(nil)).Elem()
}

func (i ContentKeyPolicyTokenClaimResponseArgs) ToContentKeyPolicyTokenClaimResponseOutput() ContentKeyPolicyTokenClaimResponseOutput {
	return i.ToContentKeyPolicyTokenClaimResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyTokenClaimResponseArgs) ToContentKeyPolicyTokenClaimResponseOutputWithContext(ctx context.Context) ContentKeyPolicyTokenClaimResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyTokenClaimResponseOutput)
}





type ContentKeyPolicyTokenClaimResponseArrayInput interface {
	pulumi.Input

	ToContentKeyPolicyTokenClaimResponseArrayOutput() ContentKeyPolicyTokenClaimResponseArrayOutput
	ToContentKeyPolicyTokenClaimResponseArrayOutputWithContext(context.Context) ContentKeyPolicyTokenClaimResponseArrayOutput
}

type ContentKeyPolicyTokenClaimResponseArray []ContentKeyPolicyTokenClaimResponseInput

func (ContentKeyPolicyTokenClaimResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyTokenClaimResponse)(nil)).Elem()
}

func (i ContentKeyPolicyTokenClaimResponseArray) ToContentKeyPolicyTokenClaimResponseArrayOutput() ContentKeyPolicyTokenClaimResponseArrayOutput {
	return i.ToContentKeyPolicyTokenClaimResponseArrayOutputWithContext(context.Background())
}

func (i ContentKeyPolicyTokenClaimResponseArray) ToContentKeyPolicyTokenClaimResponseArrayOutputWithContext(ctx context.Context) ContentKeyPolicyTokenClaimResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyTokenClaimResponseArrayOutput)
}

type ContentKeyPolicyTokenClaimResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyTokenClaimResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyTokenClaimResponse)(nil)).Elem()
}

func (o ContentKeyPolicyTokenClaimResponseOutput) ToContentKeyPolicyTokenClaimResponseOutput() ContentKeyPolicyTokenClaimResponseOutput {
	return o
}

func (o ContentKeyPolicyTokenClaimResponseOutput) ToContentKeyPolicyTokenClaimResponseOutputWithContext(ctx context.Context) ContentKeyPolicyTokenClaimResponseOutput {
	return o
}

func (o ContentKeyPolicyTokenClaimResponseOutput) ClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenClaimResponse) *string { return v.ClaimType }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyTokenClaimResponseOutput) ClaimValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenClaimResponse) *string { return v.ClaimValue }).(pulumi.StringPtrOutput)
}

type ContentKeyPolicyTokenClaimResponseArrayOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyTokenClaimResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyTokenClaimResponse)(nil)).Elem()
}

func (o ContentKeyPolicyTokenClaimResponseArrayOutput) ToContentKeyPolicyTokenClaimResponseArrayOutput() ContentKeyPolicyTokenClaimResponseArrayOutput {
	return o
}

func (o ContentKeyPolicyTokenClaimResponseArrayOutput) ToContentKeyPolicyTokenClaimResponseArrayOutputWithContext(ctx context.Context) ContentKeyPolicyTokenClaimResponseArrayOutput {
	return o
}

func (o ContentKeyPolicyTokenClaimResponseArrayOutput) Index(i pulumi.IntInput) ContentKeyPolicyTokenClaimResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContentKeyPolicyTokenClaimResponse {
		return vs[0].([]ContentKeyPolicyTokenClaimResponse)[vs[1].(int)]
	}).(ContentKeyPolicyTokenClaimResponseOutput)
}

type ContentKeyPolicyTokenRestriction struct {
	AlternateVerificationKeys      []interface{}                `pulumi:"alternateVerificationKeys"`
	Audience                       string                       `pulumi:"audience"`
	Issuer                         string                       `pulumi:"issuer"`
	OdataType                      string                       `pulumi:"odataType"`
	OpenIdConnectDiscoveryDocument *string                      `pulumi:"openIdConnectDiscoveryDocument"`
	PrimaryVerificationKey         interface{}                  `pulumi:"primaryVerificationKey"`
	RequiredClaims                 []ContentKeyPolicyTokenClaim `pulumi:"requiredClaims"`
	RestrictionTokenType           string                       `pulumi:"restrictionTokenType"`
}





type ContentKeyPolicyTokenRestrictionInput interface {
	pulumi.Input

	ToContentKeyPolicyTokenRestrictionOutput() ContentKeyPolicyTokenRestrictionOutput
	ToContentKeyPolicyTokenRestrictionOutputWithContext(context.Context) ContentKeyPolicyTokenRestrictionOutput
}

type ContentKeyPolicyTokenRestrictionArgs struct {
	AlternateVerificationKeys      pulumi.ArrayInput                    `pulumi:"alternateVerificationKeys"`
	Audience                       pulumi.StringInput                   `pulumi:"audience"`
	Issuer                         pulumi.StringInput                   `pulumi:"issuer"`
	OdataType                      pulumi.StringInput                   `pulumi:"odataType"`
	OpenIdConnectDiscoveryDocument pulumi.StringPtrInput                `pulumi:"openIdConnectDiscoveryDocument"`
	PrimaryVerificationKey         pulumi.Input                         `pulumi:"primaryVerificationKey"`
	RequiredClaims                 ContentKeyPolicyTokenClaimArrayInput `pulumi:"requiredClaims"`
	RestrictionTokenType           pulumi.StringInput                   `pulumi:"restrictionTokenType"`
}

func (ContentKeyPolicyTokenRestrictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyTokenRestriction)(nil)).Elem()
}

func (i ContentKeyPolicyTokenRestrictionArgs) ToContentKeyPolicyTokenRestrictionOutput() ContentKeyPolicyTokenRestrictionOutput {
	return i.ToContentKeyPolicyTokenRestrictionOutputWithContext(context.Background())
}

func (i ContentKeyPolicyTokenRestrictionArgs) ToContentKeyPolicyTokenRestrictionOutputWithContext(ctx context.Context) ContentKeyPolicyTokenRestrictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyTokenRestrictionOutput)
}

type ContentKeyPolicyTokenRestrictionOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyTokenRestrictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyTokenRestriction)(nil)).Elem()
}

func (o ContentKeyPolicyTokenRestrictionOutput) ToContentKeyPolicyTokenRestrictionOutput() ContentKeyPolicyTokenRestrictionOutput {
	return o
}

func (o ContentKeyPolicyTokenRestrictionOutput) ToContentKeyPolicyTokenRestrictionOutputWithContext(ctx context.Context) ContentKeyPolicyTokenRestrictionOutput {
	return o
}

func (o ContentKeyPolicyTokenRestrictionOutput) AlternateVerificationKeys() pulumi.ArrayOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestriction) []interface{} { return v.AlternateVerificationKeys }).(pulumi.ArrayOutput)
}

func (o ContentKeyPolicyTokenRestrictionOutput) Audience() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestriction) string { return v.Audience }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyTokenRestrictionOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestriction) string { return v.Issuer }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyTokenRestrictionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestriction) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyTokenRestrictionOutput) OpenIdConnectDiscoveryDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestriction) *string { return v.OpenIdConnectDiscoveryDocument }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyTokenRestrictionOutput) PrimaryVerificationKey() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestriction) interface{} { return v.PrimaryVerificationKey }).(pulumi.AnyOutput)
}

func (o ContentKeyPolicyTokenRestrictionOutput) RequiredClaims() ContentKeyPolicyTokenClaimArrayOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestriction) []ContentKeyPolicyTokenClaim { return v.RequiredClaims }).(ContentKeyPolicyTokenClaimArrayOutput)
}

func (o ContentKeyPolicyTokenRestrictionOutput) RestrictionTokenType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestriction) string { return v.RestrictionTokenType }).(pulumi.StringOutput)
}

type ContentKeyPolicyTokenRestrictionResponse struct {
	AlternateVerificationKeys      []interface{}                        `pulumi:"alternateVerificationKeys"`
	Audience                       string                               `pulumi:"audience"`
	Issuer                         string                               `pulumi:"issuer"`
	OdataType                      string                               `pulumi:"odataType"`
	OpenIdConnectDiscoveryDocument *string                              `pulumi:"openIdConnectDiscoveryDocument"`
	PrimaryVerificationKey         interface{}                          `pulumi:"primaryVerificationKey"`
	RequiredClaims                 []ContentKeyPolicyTokenClaimResponse `pulumi:"requiredClaims"`
	RestrictionTokenType           string                               `pulumi:"restrictionTokenType"`
}





type ContentKeyPolicyTokenRestrictionResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyTokenRestrictionResponseOutput() ContentKeyPolicyTokenRestrictionResponseOutput
	ToContentKeyPolicyTokenRestrictionResponseOutputWithContext(context.Context) ContentKeyPolicyTokenRestrictionResponseOutput
}

type ContentKeyPolicyTokenRestrictionResponseArgs struct {
	AlternateVerificationKeys      pulumi.ArrayInput                            `pulumi:"alternateVerificationKeys"`
	Audience                       pulumi.StringInput                           `pulumi:"audience"`
	Issuer                         pulumi.StringInput                           `pulumi:"issuer"`
	OdataType                      pulumi.StringInput                           `pulumi:"odataType"`
	OpenIdConnectDiscoveryDocument pulumi.StringPtrInput                        `pulumi:"openIdConnectDiscoveryDocument"`
	PrimaryVerificationKey         pulumi.Input                                 `pulumi:"primaryVerificationKey"`
	RequiredClaims                 ContentKeyPolicyTokenClaimResponseArrayInput `pulumi:"requiredClaims"`
	RestrictionTokenType           pulumi.StringInput                           `pulumi:"restrictionTokenType"`
}

func (ContentKeyPolicyTokenRestrictionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyTokenRestrictionResponse)(nil)).Elem()
}

func (i ContentKeyPolicyTokenRestrictionResponseArgs) ToContentKeyPolicyTokenRestrictionResponseOutput() ContentKeyPolicyTokenRestrictionResponseOutput {
	return i.ToContentKeyPolicyTokenRestrictionResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyTokenRestrictionResponseArgs) ToContentKeyPolicyTokenRestrictionResponseOutputWithContext(ctx context.Context) ContentKeyPolicyTokenRestrictionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyTokenRestrictionResponseOutput)
}

type ContentKeyPolicyTokenRestrictionResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyTokenRestrictionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyTokenRestrictionResponse)(nil)).Elem()
}

func (o ContentKeyPolicyTokenRestrictionResponseOutput) ToContentKeyPolicyTokenRestrictionResponseOutput() ContentKeyPolicyTokenRestrictionResponseOutput {
	return o
}

func (o ContentKeyPolicyTokenRestrictionResponseOutput) ToContentKeyPolicyTokenRestrictionResponseOutputWithContext(ctx context.Context) ContentKeyPolicyTokenRestrictionResponseOutput {
	return o
}

func (o ContentKeyPolicyTokenRestrictionResponseOutput) AlternateVerificationKeys() pulumi.ArrayOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestrictionResponse) []interface{} { return v.AlternateVerificationKeys }).(pulumi.ArrayOutput)
}

func (o ContentKeyPolicyTokenRestrictionResponseOutput) Audience() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestrictionResponse) string { return v.Audience }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyTokenRestrictionResponseOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestrictionResponse) string { return v.Issuer }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyTokenRestrictionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestrictionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyTokenRestrictionResponseOutput) OpenIdConnectDiscoveryDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestrictionResponse) *string { return v.OpenIdConnectDiscoveryDocument }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyTokenRestrictionResponseOutput) PrimaryVerificationKey() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestrictionResponse) interface{} { return v.PrimaryVerificationKey }).(pulumi.AnyOutput)
}

func (o ContentKeyPolicyTokenRestrictionResponseOutput) RequiredClaims() ContentKeyPolicyTokenClaimResponseArrayOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestrictionResponse) []ContentKeyPolicyTokenClaimResponse {
		return v.RequiredClaims
	}).(ContentKeyPolicyTokenClaimResponseArrayOutput)
}

func (o ContentKeyPolicyTokenRestrictionResponseOutput) RestrictionTokenType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyTokenRestrictionResponse) string { return v.RestrictionTokenType }).(pulumi.StringOutput)
}

type ContentKeyPolicyUnknownConfiguration struct {
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyUnknownConfigurationInput interface {
	pulumi.Input

	ToContentKeyPolicyUnknownConfigurationOutput() ContentKeyPolicyUnknownConfigurationOutput
	ToContentKeyPolicyUnknownConfigurationOutputWithContext(context.Context) ContentKeyPolicyUnknownConfigurationOutput
}

type ContentKeyPolicyUnknownConfigurationArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyUnknownConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyUnknownConfiguration)(nil)).Elem()
}

func (i ContentKeyPolicyUnknownConfigurationArgs) ToContentKeyPolicyUnknownConfigurationOutput() ContentKeyPolicyUnknownConfigurationOutput {
	return i.ToContentKeyPolicyUnknownConfigurationOutputWithContext(context.Background())
}

func (i ContentKeyPolicyUnknownConfigurationArgs) ToContentKeyPolicyUnknownConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyUnknownConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyUnknownConfigurationOutput)
}

type ContentKeyPolicyUnknownConfigurationOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyUnknownConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyUnknownConfiguration)(nil)).Elem()
}

func (o ContentKeyPolicyUnknownConfigurationOutput) ToContentKeyPolicyUnknownConfigurationOutput() ContentKeyPolicyUnknownConfigurationOutput {
	return o
}

func (o ContentKeyPolicyUnknownConfigurationOutput) ToContentKeyPolicyUnknownConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyUnknownConfigurationOutput {
	return o
}

func (o ContentKeyPolicyUnknownConfigurationOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyUnknownConfiguration) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyUnknownConfigurationResponse struct {
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyUnknownConfigurationResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyUnknownConfigurationResponseOutput() ContentKeyPolicyUnknownConfigurationResponseOutput
	ToContentKeyPolicyUnknownConfigurationResponseOutputWithContext(context.Context) ContentKeyPolicyUnknownConfigurationResponseOutput
}

type ContentKeyPolicyUnknownConfigurationResponseArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyUnknownConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyUnknownConfigurationResponse)(nil)).Elem()
}

func (i ContentKeyPolicyUnknownConfigurationResponseArgs) ToContentKeyPolicyUnknownConfigurationResponseOutput() ContentKeyPolicyUnknownConfigurationResponseOutput {
	return i.ToContentKeyPolicyUnknownConfigurationResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyUnknownConfigurationResponseArgs) ToContentKeyPolicyUnknownConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyUnknownConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyUnknownConfigurationResponseOutput)
}

type ContentKeyPolicyUnknownConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyUnknownConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyUnknownConfigurationResponse)(nil)).Elem()
}

func (o ContentKeyPolicyUnknownConfigurationResponseOutput) ToContentKeyPolicyUnknownConfigurationResponseOutput() ContentKeyPolicyUnknownConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyUnknownConfigurationResponseOutput) ToContentKeyPolicyUnknownConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyUnknownConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyUnknownConfigurationResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyUnknownConfigurationResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyUnknownRestriction struct {
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyUnknownRestrictionInput interface {
	pulumi.Input

	ToContentKeyPolicyUnknownRestrictionOutput() ContentKeyPolicyUnknownRestrictionOutput
	ToContentKeyPolicyUnknownRestrictionOutputWithContext(context.Context) ContentKeyPolicyUnknownRestrictionOutput
}

type ContentKeyPolicyUnknownRestrictionArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyUnknownRestrictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyUnknownRestriction)(nil)).Elem()
}

func (i ContentKeyPolicyUnknownRestrictionArgs) ToContentKeyPolicyUnknownRestrictionOutput() ContentKeyPolicyUnknownRestrictionOutput {
	return i.ToContentKeyPolicyUnknownRestrictionOutputWithContext(context.Background())
}

func (i ContentKeyPolicyUnknownRestrictionArgs) ToContentKeyPolicyUnknownRestrictionOutputWithContext(ctx context.Context) ContentKeyPolicyUnknownRestrictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyUnknownRestrictionOutput)
}

type ContentKeyPolicyUnknownRestrictionOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyUnknownRestrictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyUnknownRestriction)(nil)).Elem()
}

func (o ContentKeyPolicyUnknownRestrictionOutput) ToContentKeyPolicyUnknownRestrictionOutput() ContentKeyPolicyUnknownRestrictionOutput {
	return o
}

func (o ContentKeyPolicyUnknownRestrictionOutput) ToContentKeyPolicyUnknownRestrictionOutputWithContext(ctx context.Context) ContentKeyPolicyUnknownRestrictionOutput {
	return o
}

func (o ContentKeyPolicyUnknownRestrictionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyUnknownRestriction) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyUnknownRestrictionResponse struct {
	OdataType string `pulumi:"odataType"`
}





type ContentKeyPolicyUnknownRestrictionResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyUnknownRestrictionResponseOutput() ContentKeyPolicyUnknownRestrictionResponseOutput
	ToContentKeyPolicyUnknownRestrictionResponseOutputWithContext(context.Context) ContentKeyPolicyUnknownRestrictionResponseOutput
}

type ContentKeyPolicyUnknownRestrictionResponseArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (ContentKeyPolicyUnknownRestrictionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyUnknownRestrictionResponse)(nil)).Elem()
}

func (i ContentKeyPolicyUnknownRestrictionResponseArgs) ToContentKeyPolicyUnknownRestrictionResponseOutput() ContentKeyPolicyUnknownRestrictionResponseOutput {
	return i.ToContentKeyPolicyUnknownRestrictionResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyUnknownRestrictionResponseArgs) ToContentKeyPolicyUnknownRestrictionResponseOutputWithContext(ctx context.Context) ContentKeyPolicyUnknownRestrictionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyUnknownRestrictionResponseOutput)
}

type ContentKeyPolicyUnknownRestrictionResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyUnknownRestrictionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyUnknownRestrictionResponse)(nil)).Elem()
}

func (o ContentKeyPolicyUnknownRestrictionResponseOutput) ToContentKeyPolicyUnknownRestrictionResponseOutput() ContentKeyPolicyUnknownRestrictionResponseOutput {
	return o
}

func (o ContentKeyPolicyUnknownRestrictionResponseOutput) ToContentKeyPolicyUnknownRestrictionResponseOutputWithContext(ctx context.Context) ContentKeyPolicyUnknownRestrictionResponseOutput {
	return o
}

func (o ContentKeyPolicyUnknownRestrictionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyUnknownRestrictionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type ContentKeyPolicyWidevineConfiguration struct {
	OdataType        string `pulumi:"odataType"`
	WidevineTemplate string `pulumi:"widevineTemplate"`
}





type ContentKeyPolicyWidevineConfigurationInput interface {
	pulumi.Input

	ToContentKeyPolicyWidevineConfigurationOutput() ContentKeyPolicyWidevineConfigurationOutput
	ToContentKeyPolicyWidevineConfigurationOutputWithContext(context.Context) ContentKeyPolicyWidevineConfigurationOutput
}

type ContentKeyPolicyWidevineConfigurationArgs struct {
	OdataType        pulumi.StringInput `pulumi:"odataType"`
	WidevineTemplate pulumi.StringInput `pulumi:"widevineTemplate"`
}

func (ContentKeyPolicyWidevineConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyWidevineConfiguration)(nil)).Elem()
}

func (i ContentKeyPolicyWidevineConfigurationArgs) ToContentKeyPolicyWidevineConfigurationOutput() ContentKeyPolicyWidevineConfigurationOutput {
	return i.ToContentKeyPolicyWidevineConfigurationOutputWithContext(context.Background())
}

func (i ContentKeyPolicyWidevineConfigurationArgs) ToContentKeyPolicyWidevineConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyWidevineConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyWidevineConfigurationOutput)
}

type ContentKeyPolicyWidevineConfigurationOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyWidevineConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyWidevineConfiguration)(nil)).Elem()
}

func (o ContentKeyPolicyWidevineConfigurationOutput) ToContentKeyPolicyWidevineConfigurationOutput() ContentKeyPolicyWidevineConfigurationOutput {
	return o
}

func (o ContentKeyPolicyWidevineConfigurationOutput) ToContentKeyPolicyWidevineConfigurationOutputWithContext(ctx context.Context) ContentKeyPolicyWidevineConfigurationOutput {
	return o
}

func (o ContentKeyPolicyWidevineConfigurationOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyWidevineConfiguration) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyWidevineConfigurationOutput) WidevineTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyWidevineConfiguration) string { return v.WidevineTemplate }).(pulumi.StringOutput)
}

type ContentKeyPolicyWidevineConfigurationResponse struct {
	OdataType        string `pulumi:"odataType"`
	WidevineTemplate string `pulumi:"widevineTemplate"`
}





type ContentKeyPolicyWidevineConfigurationResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyWidevineConfigurationResponseOutput() ContentKeyPolicyWidevineConfigurationResponseOutput
	ToContentKeyPolicyWidevineConfigurationResponseOutputWithContext(context.Context) ContentKeyPolicyWidevineConfigurationResponseOutput
}

type ContentKeyPolicyWidevineConfigurationResponseArgs struct {
	OdataType        pulumi.StringInput `pulumi:"odataType"`
	WidevineTemplate pulumi.StringInput `pulumi:"widevineTemplate"`
}

func (ContentKeyPolicyWidevineConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyWidevineConfigurationResponse)(nil)).Elem()
}

func (i ContentKeyPolicyWidevineConfigurationResponseArgs) ToContentKeyPolicyWidevineConfigurationResponseOutput() ContentKeyPolicyWidevineConfigurationResponseOutput {
	return i.ToContentKeyPolicyWidevineConfigurationResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyWidevineConfigurationResponseArgs) ToContentKeyPolicyWidevineConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyWidevineConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyWidevineConfigurationResponseOutput)
}

type ContentKeyPolicyWidevineConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyWidevineConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyWidevineConfigurationResponse)(nil)).Elem()
}

func (o ContentKeyPolicyWidevineConfigurationResponseOutput) ToContentKeyPolicyWidevineConfigurationResponseOutput() ContentKeyPolicyWidevineConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyWidevineConfigurationResponseOutput) ToContentKeyPolicyWidevineConfigurationResponseOutputWithContext(ctx context.Context) ContentKeyPolicyWidevineConfigurationResponseOutput {
	return o
}

func (o ContentKeyPolicyWidevineConfigurationResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyWidevineConfigurationResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyWidevineConfigurationResponseOutput) WidevineTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyWidevineConfigurationResponse) string { return v.WidevineTemplate }).(pulumi.StringOutput)
}

type ContentKeyPolicyX509CertificateTokenKey struct {
	OdataType string `pulumi:"odataType"`
	RawBody   string `pulumi:"rawBody"`
}





type ContentKeyPolicyX509CertificateTokenKeyInput interface {
	pulumi.Input

	ToContentKeyPolicyX509CertificateTokenKeyOutput() ContentKeyPolicyX509CertificateTokenKeyOutput
	ToContentKeyPolicyX509CertificateTokenKeyOutputWithContext(context.Context) ContentKeyPolicyX509CertificateTokenKeyOutput
}

type ContentKeyPolicyX509CertificateTokenKeyArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
	RawBody   pulumi.StringInput `pulumi:"rawBody"`
}

func (ContentKeyPolicyX509CertificateTokenKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyX509CertificateTokenKey)(nil)).Elem()
}

func (i ContentKeyPolicyX509CertificateTokenKeyArgs) ToContentKeyPolicyX509CertificateTokenKeyOutput() ContentKeyPolicyX509CertificateTokenKeyOutput {
	return i.ToContentKeyPolicyX509CertificateTokenKeyOutputWithContext(context.Background())
}

func (i ContentKeyPolicyX509CertificateTokenKeyArgs) ToContentKeyPolicyX509CertificateTokenKeyOutputWithContext(ctx context.Context) ContentKeyPolicyX509CertificateTokenKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyX509CertificateTokenKeyOutput)
}

type ContentKeyPolicyX509CertificateTokenKeyOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyX509CertificateTokenKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyX509CertificateTokenKey)(nil)).Elem()
}

func (o ContentKeyPolicyX509CertificateTokenKeyOutput) ToContentKeyPolicyX509CertificateTokenKeyOutput() ContentKeyPolicyX509CertificateTokenKeyOutput {
	return o
}

func (o ContentKeyPolicyX509CertificateTokenKeyOutput) ToContentKeyPolicyX509CertificateTokenKeyOutputWithContext(ctx context.Context) ContentKeyPolicyX509CertificateTokenKeyOutput {
	return o
}

func (o ContentKeyPolicyX509CertificateTokenKeyOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyX509CertificateTokenKey) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyX509CertificateTokenKeyOutput) RawBody() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyX509CertificateTokenKey) string { return v.RawBody }).(pulumi.StringOutput)
}

type ContentKeyPolicyX509CertificateTokenKeyResponse struct {
	OdataType string `pulumi:"odataType"`
	RawBody   string `pulumi:"rawBody"`
}





type ContentKeyPolicyX509CertificateTokenKeyResponseInput interface {
	pulumi.Input

	ToContentKeyPolicyX509CertificateTokenKeyResponseOutput() ContentKeyPolicyX509CertificateTokenKeyResponseOutput
	ToContentKeyPolicyX509CertificateTokenKeyResponseOutputWithContext(context.Context) ContentKeyPolicyX509CertificateTokenKeyResponseOutput
}

type ContentKeyPolicyX509CertificateTokenKeyResponseArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
	RawBody   pulumi.StringInput `pulumi:"rawBody"`
}

func (ContentKeyPolicyX509CertificateTokenKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyX509CertificateTokenKeyResponse)(nil)).Elem()
}

func (i ContentKeyPolicyX509CertificateTokenKeyResponseArgs) ToContentKeyPolicyX509CertificateTokenKeyResponseOutput() ContentKeyPolicyX509CertificateTokenKeyResponseOutput {
	return i.ToContentKeyPolicyX509CertificateTokenKeyResponseOutputWithContext(context.Background())
}

func (i ContentKeyPolicyX509CertificateTokenKeyResponseArgs) ToContentKeyPolicyX509CertificateTokenKeyResponseOutputWithContext(ctx context.Context) ContentKeyPolicyX509CertificateTokenKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyX509CertificateTokenKeyResponseOutput)
}

type ContentKeyPolicyX509CertificateTokenKeyResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyX509CertificateTokenKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyX509CertificateTokenKeyResponse)(nil)).Elem()
}

func (o ContentKeyPolicyX509CertificateTokenKeyResponseOutput) ToContentKeyPolicyX509CertificateTokenKeyResponseOutput() ContentKeyPolicyX509CertificateTokenKeyResponseOutput {
	return o
}

func (o ContentKeyPolicyX509CertificateTokenKeyResponseOutput) ToContentKeyPolicyX509CertificateTokenKeyResponseOutputWithContext(ctx context.Context) ContentKeyPolicyX509CertificateTokenKeyResponseOutput {
	return o
}

func (o ContentKeyPolicyX509CertificateTokenKeyResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyX509CertificateTokenKeyResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyX509CertificateTokenKeyResponseOutput) RawBody() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyX509CertificateTokenKeyResponse) string { return v.RawBody }).(pulumi.StringOutput)
}

type CopyAudio struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}





type CopyAudioInput interface {
	pulumi.Input

	ToCopyAudioOutput() CopyAudioOutput
	ToCopyAudioOutputWithContext(context.Context) CopyAudioOutput
}

type CopyAudioArgs struct {
	Label     pulumi.StringPtrInput `pulumi:"label"`
	OdataType pulumi.StringInput    `pulumi:"odataType"`
}

func (CopyAudioArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyAudio)(nil)).Elem()
}

func (i CopyAudioArgs) ToCopyAudioOutput() CopyAudioOutput {
	return i.ToCopyAudioOutputWithContext(context.Background())
}

func (i CopyAudioArgs) ToCopyAudioOutputWithContext(ctx context.Context) CopyAudioOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyAudioOutput)
}

type CopyAudioOutput struct{ *pulumi.OutputState }

func (CopyAudioOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyAudio)(nil)).Elem()
}

func (o CopyAudioOutput) ToCopyAudioOutput() CopyAudioOutput {
	return o
}

func (o CopyAudioOutput) ToCopyAudioOutputWithContext(ctx context.Context) CopyAudioOutput {
	return o
}

func (o CopyAudioOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CopyAudio) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o CopyAudioOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v CopyAudio) string { return v.OdataType }).(pulumi.StringOutput)
}

type CopyAudioResponse struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}





type CopyAudioResponseInput interface {
	pulumi.Input

	ToCopyAudioResponseOutput() CopyAudioResponseOutput
	ToCopyAudioResponseOutputWithContext(context.Context) CopyAudioResponseOutput
}

type CopyAudioResponseArgs struct {
	Label     pulumi.StringPtrInput `pulumi:"label"`
	OdataType pulumi.StringInput    `pulumi:"odataType"`
}

func (CopyAudioResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyAudioResponse)(nil)).Elem()
}

func (i CopyAudioResponseArgs) ToCopyAudioResponseOutput() CopyAudioResponseOutput {
	return i.ToCopyAudioResponseOutputWithContext(context.Background())
}

func (i CopyAudioResponseArgs) ToCopyAudioResponseOutputWithContext(ctx context.Context) CopyAudioResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyAudioResponseOutput)
}

type CopyAudioResponseOutput struct{ *pulumi.OutputState }

func (CopyAudioResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyAudioResponse)(nil)).Elem()
}

func (o CopyAudioResponseOutput) ToCopyAudioResponseOutput() CopyAudioResponseOutput {
	return o
}

func (o CopyAudioResponseOutput) ToCopyAudioResponseOutputWithContext(ctx context.Context) CopyAudioResponseOutput {
	return o
}

func (o CopyAudioResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CopyAudioResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o CopyAudioResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v CopyAudioResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type CopyVideo struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}





type CopyVideoInput interface {
	pulumi.Input

	ToCopyVideoOutput() CopyVideoOutput
	ToCopyVideoOutputWithContext(context.Context) CopyVideoOutput
}

type CopyVideoArgs struct {
	Label     pulumi.StringPtrInput `pulumi:"label"`
	OdataType pulumi.StringInput    `pulumi:"odataType"`
}

func (CopyVideoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyVideo)(nil)).Elem()
}

func (i CopyVideoArgs) ToCopyVideoOutput() CopyVideoOutput {
	return i.ToCopyVideoOutputWithContext(context.Background())
}

func (i CopyVideoArgs) ToCopyVideoOutputWithContext(ctx context.Context) CopyVideoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyVideoOutput)
}

type CopyVideoOutput struct{ *pulumi.OutputState }

func (CopyVideoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyVideo)(nil)).Elem()
}

func (o CopyVideoOutput) ToCopyVideoOutput() CopyVideoOutput {
	return o
}

func (o CopyVideoOutput) ToCopyVideoOutputWithContext(ctx context.Context) CopyVideoOutput {
	return o
}

func (o CopyVideoOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CopyVideo) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o CopyVideoOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v CopyVideo) string { return v.OdataType }).(pulumi.StringOutput)
}

type CopyVideoResponse struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}





type CopyVideoResponseInput interface {
	pulumi.Input

	ToCopyVideoResponseOutput() CopyVideoResponseOutput
	ToCopyVideoResponseOutputWithContext(context.Context) CopyVideoResponseOutput
}

type CopyVideoResponseArgs struct {
	Label     pulumi.StringPtrInput `pulumi:"label"`
	OdataType pulumi.StringInput    `pulumi:"odataType"`
}

func (CopyVideoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyVideoResponse)(nil)).Elem()
}

func (i CopyVideoResponseArgs) ToCopyVideoResponseOutput() CopyVideoResponseOutput {
	return i.ToCopyVideoResponseOutputWithContext(context.Background())
}

func (i CopyVideoResponseArgs) ToCopyVideoResponseOutputWithContext(ctx context.Context) CopyVideoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CopyVideoResponseOutput)
}

type CopyVideoResponseOutput struct{ *pulumi.OutputState }

func (CopyVideoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CopyVideoResponse)(nil)).Elem()
}

func (o CopyVideoResponseOutput) ToCopyVideoResponseOutput() CopyVideoResponseOutput {
	return o
}

func (o CopyVideoResponseOutput) ToCopyVideoResponseOutputWithContext(ctx context.Context) CopyVideoResponseOutput {
	return o
}

func (o CopyVideoResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CopyVideoResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o CopyVideoResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v CopyVideoResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type CrossSiteAccessPolicies struct {
	ClientAccessPolicy *string `pulumi:"clientAccessPolicy"`
	CrossDomainPolicy  *string `pulumi:"crossDomainPolicy"`
}





type CrossSiteAccessPoliciesInput interface {
	pulumi.Input

	ToCrossSiteAccessPoliciesOutput() CrossSiteAccessPoliciesOutput
	ToCrossSiteAccessPoliciesOutputWithContext(context.Context) CrossSiteAccessPoliciesOutput
}

type CrossSiteAccessPoliciesArgs struct {
	ClientAccessPolicy pulumi.StringPtrInput `pulumi:"clientAccessPolicy"`
	CrossDomainPolicy  pulumi.StringPtrInput `pulumi:"crossDomainPolicy"`
}

func (CrossSiteAccessPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossSiteAccessPolicies)(nil)).Elem()
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesOutput() CrossSiteAccessPoliciesOutput {
	return i.ToCrossSiteAccessPoliciesOutputWithContext(context.Background())
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesOutput)
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return i.ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Background())
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesOutput).ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx)
}









type CrossSiteAccessPoliciesPtrInput interface {
	pulumi.Input

	ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput
	ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Context) CrossSiteAccessPoliciesPtrOutput
}

type crossSiteAccessPoliciesPtrType CrossSiteAccessPoliciesArgs

func CrossSiteAccessPoliciesPtr(v *CrossSiteAccessPoliciesArgs) CrossSiteAccessPoliciesPtrInput {
	return (*crossSiteAccessPoliciesPtrType)(v)
}

func (*crossSiteAccessPoliciesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossSiteAccessPolicies)(nil)).Elem()
}

func (i *crossSiteAccessPoliciesPtrType) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return i.ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Background())
}

func (i *crossSiteAccessPoliciesPtrType) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesPtrOutput)
}

type CrossSiteAccessPoliciesOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossSiteAccessPolicies)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesOutput() CrossSiteAccessPoliciesOutput {
	return o
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesOutput {
	return o
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return o.ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Background())
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CrossSiteAccessPolicies) *CrossSiteAccessPolicies {
		return &v
	}).(CrossSiteAccessPoliciesPtrOutput)
}

func (o CrossSiteAccessPoliciesOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPolicies) *string { return v.ClientAccessPolicy }).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPolicies) *string { return v.CrossDomainPolicy }).(pulumi.StringPtrOutput)
}

type CrossSiteAccessPoliciesPtrOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossSiteAccessPolicies)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesPtrOutput) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesPtrOutput) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesPtrOutput) Elem() CrossSiteAccessPoliciesOutput {
	return o.ApplyT(func(v *CrossSiteAccessPolicies) CrossSiteAccessPolicies {
		if v != nil {
			return *v
		}
		var ret CrossSiteAccessPolicies
		return ret
	}).(CrossSiteAccessPoliciesOutput)
}

func (o CrossSiteAccessPoliciesPtrOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPolicies) *string {
		if v == nil {
			return nil
		}
		return v.ClientAccessPolicy
	}).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesPtrOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPolicies) *string {
		if v == nil {
			return nil
		}
		return v.CrossDomainPolicy
	}).(pulumi.StringPtrOutput)
}

type CrossSiteAccessPoliciesResponse struct {
	ClientAccessPolicy *string `pulumi:"clientAccessPolicy"`
	CrossDomainPolicy  *string `pulumi:"crossDomainPolicy"`
}





type CrossSiteAccessPoliciesResponseInput interface {
	pulumi.Input

	ToCrossSiteAccessPoliciesResponseOutput() CrossSiteAccessPoliciesResponseOutput
	ToCrossSiteAccessPoliciesResponseOutputWithContext(context.Context) CrossSiteAccessPoliciesResponseOutput
}

type CrossSiteAccessPoliciesResponseArgs struct {
	ClientAccessPolicy pulumi.StringPtrInput `pulumi:"clientAccessPolicy"`
	CrossDomainPolicy  pulumi.StringPtrInput `pulumi:"crossDomainPolicy"`
}

func (CrossSiteAccessPoliciesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossSiteAccessPoliciesResponse)(nil)).Elem()
}

func (i CrossSiteAccessPoliciesResponseArgs) ToCrossSiteAccessPoliciesResponseOutput() CrossSiteAccessPoliciesResponseOutput {
	return i.ToCrossSiteAccessPoliciesResponseOutputWithContext(context.Background())
}

func (i CrossSiteAccessPoliciesResponseArgs) ToCrossSiteAccessPoliciesResponseOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesResponseOutput)
}

func (i CrossSiteAccessPoliciesResponseArgs) ToCrossSiteAccessPoliciesResponsePtrOutput() CrossSiteAccessPoliciesResponsePtrOutput {
	return i.ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(context.Background())
}

func (i CrossSiteAccessPoliciesResponseArgs) ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesResponseOutput).ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(ctx)
}









type CrossSiteAccessPoliciesResponsePtrInput interface {
	pulumi.Input

	ToCrossSiteAccessPoliciesResponsePtrOutput() CrossSiteAccessPoliciesResponsePtrOutput
	ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(context.Context) CrossSiteAccessPoliciesResponsePtrOutput
}

type crossSiteAccessPoliciesResponsePtrType CrossSiteAccessPoliciesResponseArgs

func CrossSiteAccessPoliciesResponsePtr(v *CrossSiteAccessPoliciesResponseArgs) CrossSiteAccessPoliciesResponsePtrInput {
	return (*crossSiteAccessPoliciesResponsePtrType)(v)
}

func (*crossSiteAccessPoliciesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossSiteAccessPoliciesResponse)(nil)).Elem()
}

func (i *crossSiteAccessPoliciesResponsePtrType) ToCrossSiteAccessPoliciesResponsePtrOutput() CrossSiteAccessPoliciesResponsePtrOutput {
	return i.ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(context.Background())
}

func (i *crossSiteAccessPoliciesResponsePtrType) ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesResponsePtrOutput)
}

type CrossSiteAccessPoliciesResponseOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossSiteAccessPoliciesResponse)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesResponseOutput) ToCrossSiteAccessPoliciesResponseOutput() CrossSiteAccessPoliciesResponseOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponseOutput) ToCrossSiteAccessPoliciesResponseOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponseOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponseOutput) ToCrossSiteAccessPoliciesResponsePtrOutput() CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(context.Background())
}

func (o CrossSiteAccessPoliciesResponseOutput) ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CrossSiteAccessPoliciesResponse) *CrossSiteAccessPoliciesResponse {
		return &v
	}).(CrossSiteAccessPoliciesResponsePtrOutput)
}

func (o CrossSiteAccessPoliciesResponseOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPoliciesResponse) *string { return v.ClientAccessPolicy }).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesResponseOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPoliciesResponse) *string { return v.CrossDomainPolicy }).(pulumi.StringPtrOutput)
}

type CrossSiteAccessPoliciesResponsePtrOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossSiteAccessPoliciesResponse)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) ToCrossSiteAccessPoliciesResponsePtrOutput() CrossSiteAccessPoliciesResponsePtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponsePtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) Elem() CrossSiteAccessPoliciesResponseOutput {
	return o.ApplyT(func(v *CrossSiteAccessPoliciesResponse) CrossSiteAccessPoliciesResponse {
		if v != nil {
			return *v
		}
		var ret CrossSiteAccessPoliciesResponse
		return ret
	}).(CrossSiteAccessPoliciesResponseOutput)
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientAccessPolicy
	}).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CrossDomainPolicy
	}).(pulumi.StringPtrOutput)
}

type DefaultKey struct {
	Label      *string `pulumi:"label"`
	PolicyName *string `pulumi:"policyName"`
}





type DefaultKeyInput interface {
	pulumi.Input

	ToDefaultKeyOutput() DefaultKeyOutput
	ToDefaultKeyOutputWithContext(context.Context) DefaultKeyOutput
}

type DefaultKeyArgs struct {
	Label      pulumi.StringPtrInput `pulumi:"label"`
	PolicyName pulumi.StringPtrInput `pulumi:"policyName"`
}

func (DefaultKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultKey)(nil)).Elem()
}

func (i DefaultKeyArgs) ToDefaultKeyOutput() DefaultKeyOutput {
	return i.ToDefaultKeyOutputWithContext(context.Background())
}

func (i DefaultKeyArgs) ToDefaultKeyOutputWithContext(ctx context.Context) DefaultKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultKeyOutput)
}

func (i DefaultKeyArgs) ToDefaultKeyPtrOutput() DefaultKeyPtrOutput {
	return i.ToDefaultKeyPtrOutputWithContext(context.Background())
}

func (i DefaultKeyArgs) ToDefaultKeyPtrOutputWithContext(ctx context.Context) DefaultKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultKeyOutput).ToDefaultKeyPtrOutputWithContext(ctx)
}









type DefaultKeyPtrInput interface {
	pulumi.Input

	ToDefaultKeyPtrOutput() DefaultKeyPtrOutput
	ToDefaultKeyPtrOutputWithContext(context.Context) DefaultKeyPtrOutput
}

type defaultKeyPtrType DefaultKeyArgs

func DefaultKeyPtr(v *DefaultKeyArgs) DefaultKeyPtrInput {
	return (*defaultKeyPtrType)(v)
}

func (*defaultKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultKey)(nil)).Elem()
}

func (i *defaultKeyPtrType) ToDefaultKeyPtrOutput() DefaultKeyPtrOutput {
	return i.ToDefaultKeyPtrOutputWithContext(context.Background())
}

func (i *defaultKeyPtrType) ToDefaultKeyPtrOutputWithContext(ctx context.Context) DefaultKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultKeyPtrOutput)
}

type DefaultKeyOutput struct{ *pulumi.OutputState }

func (DefaultKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultKey)(nil)).Elem()
}

func (o DefaultKeyOutput) ToDefaultKeyOutput() DefaultKeyOutput {
	return o
}

func (o DefaultKeyOutput) ToDefaultKeyOutputWithContext(ctx context.Context) DefaultKeyOutput {
	return o
}

func (o DefaultKeyOutput) ToDefaultKeyPtrOutput() DefaultKeyPtrOutput {
	return o.ToDefaultKeyPtrOutputWithContext(context.Background())
}

func (o DefaultKeyOutput) ToDefaultKeyPtrOutputWithContext(ctx context.Context) DefaultKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultKey) *DefaultKey {
		return &v
	}).(DefaultKeyPtrOutput)
}

func (o DefaultKeyOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefaultKey) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o DefaultKeyOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefaultKey) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

type DefaultKeyPtrOutput struct{ *pulumi.OutputState }

func (DefaultKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultKey)(nil)).Elem()
}

func (o DefaultKeyPtrOutput) ToDefaultKeyPtrOutput() DefaultKeyPtrOutput {
	return o
}

func (o DefaultKeyPtrOutput) ToDefaultKeyPtrOutputWithContext(ctx context.Context) DefaultKeyPtrOutput {
	return o
}

func (o DefaultKeyPtrOutput) Elem() DefaultKeyOutput {
	return o.ApplyT(func(v *DefaultKey) DefaultKey {
		if v != nil {
			return *v
		}
		var ret DefaultKey
		return ret
	}).(DefaultKeyOutput)
}

func (o DefaultKeyPtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultKey) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o DefaultKeyPtrOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultKey) *string {
		if v == nil {
			return nil
		}
		return v.PolicyName
	}).(pulumi.StringPtrOutput)
}

type DefaultKeyResponse struct {
	Label      *string `pulumi:"label"`
	PolicyName *string `pulumi:"policyName"`
}





type DefaultKeyResponseInput interface {
	pulumi.Input

	ToDefaultKeyResponseOutput() DefaultKeyResponseOutput
	ToDefaultKeyResponseOutputWithContext(context.Context) DefaultKeyResponseOutput
}

type DefaultKeyResponseArgs struct {
	Label      pulumi.StringPtrInput `pulumi:"label"`
	PolicyName pulumi.StringPtrInput `pulumi:"policyName"`
}

func (DefaultKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultKeyResponse)(nil)).Elem()
}

func (i DefaultKeyResponseArgs) ToDefaultKeyResponseOutput() DefaultKeyResponseOutput {
	return i.ToDefaultKeyResponseOutputWithContext(context.Background())
}

func (i DefaultKeyResponseArgs) ToDefaultKeyResponseOutputWithContext(ctx context.Context) DefaultKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultKeyResponseOutput)
}

func (i DefaultKeyResponseArgs) ToDefaultKeyResponsePtrOutput() DefaultKeyResponsePtrOutput {
	return i.ToDefaultKeyResponsePtrOutputWithContext(context.Background())
}

func (i DefaultKeyResponseArgs) ToDefaultKeyResponsePtrOutputWithContext(ctx context.Context) DefaultKeyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultKeyResponseOutput).ToDefaultKeyResponsePtrOutputWithContext(ctx)
}









type DefaultKeyResponsePtrInput interface {
	pulumi.Input

	ToDefaultKeyResponsePtrOutput() DefaultKeyResponsePtrOutput
	ToDefaultKeyResponsePtrOutputWithContext(context.Context) DefaultKeyResponsePtrOutput
}

type defaultKeyResponsePtrType DefaultKeyResponseArgs

func DefaultKeyResponsePtr(v *DefaultKeyResponseArgs) DefaultKeyResponsePtrInput {
	return (*defaultKeyResponsePtrType)(v)
}

func (*defaultKeyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultKeyResponse)(nil)).Elem()
}

func (i *defaultKeyResponsePtrType) ToDefaultKeyResponsePtrOutput() DefaultKeyResponsePtrOutput {
	return i.ToDefaultKeyResponsePtrOutputWithContext(context.Background())
}

func (i *defaultKeyResponsePtrType) ToDefaultKeyResponsePtrOutputWithContext(ctx context.Context) DefaultKeyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultKeyResponsePtrOutput)
}

type DefaultKeyResponseOutput struct{ *pulumi.OutputState }

func (DefaultKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultKeyResponse)(nil)).Elem()
}

func (o DefaultKeyResponseOutput) ToDefaultKeyResponseOutput() DefaultKeyResponseOutput {
	return o
}

func (o DefaultKeyResponseOutput) ToDefaultKeyResponseOutputWithContext(ctx context.Context) DefaultKeyResponseOutput {
	return o
}

func (o DefaultKeyResponseOutput) ToDefaultKeyResponsePtrOutput() DefaultKeyResponsePtrOutput {
	return o.ToDefaultKeyResponsePtrOutputWithContext(context.Background())
}

func (o DefaultKeyResponseOutput) ToDefaultKeyResponsePtrOutputWithContext(ctx context.Context) DefaultKeyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultKeyResponse) *DefaultKeyResponse {
		return &v
	}).(DefaultKeyResponsePtrOutput)
}

func (o DefaultKeyResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefaultKeyResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o DefaultKeyResponseOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefaultKeyResponse) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

type DefaultKeyResponsePtrOutput struct{ *pulumi.OutputState }

func (DefaultKeyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultKeyResponse)(nil)).Elem()
}

func (o DefaultKeyResponsePtrOutput) ToDefaultKeyResponsePtrOutput() DefaultKeyResponsePtrOutput {
	return o
}

func (o DefaultKeyResponsePtrOutput) ToDefaultKeyResponsePtrOutputWithContext(ctx context.Context) DefaultKeyResponsePtrOutput {
	return o
}

func (o DefaultKeyResponsePtrOutput) Elem() DefaultKeyResponseOutput {
	return o.ApplyT(func(v *DefaultKeyResponse) DefaultKeyResponse {
		if v != nil {
			return *v
		}
		var ret DefaultKeyResponse
		return ret
	}).(DefaultKeyResponseOutput)
}

func (o DefaultKeyResponsePtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o DefaultKeyResponsePtrOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyName
	}).(pulumi.StringPtrOutput)
}

type Deinterlace struct {
	Mode   *string `pulumi:"mode"`
	Parity *string `pulumi:"parity"`
}





type DeinterlaceInput interface {
	pulumi.Input

	ToDeinterlaceOutput() DeinterlaceOutput
	ToDeinterlaceOutputWithContext(context.Context) DeinterlaceOutput
}

type DeinterlaceArgs struct {
	Mode   pulumi.StringPtrInput `pulumi:"mode"`
	Parity pulumi.StringPtrInput `pulumi:"parity"`
}

func (DeinterlaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Deinterlace)(nil)).Elem()
}

func (i DeinterlaceArgs) ToDeinterlaceOutput() DeinterlaceOutput {
	return i.ToDeinterlaceOutputWithContext(context.Background())
}

func (i DeinterlaceArgs) ToDeinterlaceOutputWithContext(ctx context.Context) DeinterlaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeinterlaceOutput)
}

func (i DeinterlaceArgs) ToDeinterlacePtrOutput() DeinterlacePtrOutput {
	return i.ToDeinterlacePtrOutputWithContext(context.Background())
}

func (i DeinterlaceArgs) ToDeinterlacePtrOutputWithContext(ctx context.Context) DeinterlacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeinterlaceOutput).ToDeinterlacePtrOutputWithContext(ctx)
}









type DeinterlacePtrInput interface {
	pulumi.Input

	ToDeinterlacePtrOutput() DeinterlacePtrOutput
	ToDeinterlacePtrOutputWithContext(context.Context) DeinterlacePtrOutput
}

type deinterlacePtrType DeinterlaceArgs

func DeinterlacePtr(v *DeinterlaceArgs) DeinterlacePtrInput {
	return (*deinterlacePtrType)(v)
}

func (*deinterlacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Deinterlace)(nil)).Elem()
}

func (i *deinterlacePtrType) ToDeinterlacePtrOutput() DeinterlacePtrOutput {
	return i.ToDeinterlacePtrOutputWithContext(context.Background())
}

func (i *deinterlacePtrType) ToDeinterlacePtrOutputWithContext(ctx context.Context) DeinterlacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeinterlacePtrOutput)
}

type DeinterlaceOutput struct{ *pulumi.OutputState }

func (DeinterlaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Deinterlace)(nil)).Elem()
}

func (o DeinterlaceOutput) ToDeinterlaceOutput() DeinterlaceOutput {
	return o
}

func (o DeinterlaceOutput) ToDeinterlaceOutputWithContext(ctx context.Context) DeinterlaceOutput {
	return o
}

func (o DeinterlaceOutput) ToDeinterlacePtrOutput() DeinterlacePtrOutput {
	return o.ToDeinterlacePtrOutputWithContext(context.Background())
}

func (o DeinterlaceOutput) ToDeinterlacePtrOutputWithContext(ctx context.Context) DeinterlacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Deinterlace) *Deinterlace {
		return &v
	}).(DeinterlacePtrOutput)
}

func (o DeinterlaceOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Deinterlace) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o DeinterlaceOutput) Parity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Deinterlace) *string { return v.Parity }).(pulumi.StringPtrOutput)
}

type DeinterlacePtrOutput struct{ *pulumi.OutputState }

func (DeinterlacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deinterlace)(nil)).Elem()
}

func (o DeinterlacePtrOutput) ToDeinterlacePtrOutput() DeinterlacePtrOutput {
	return o
}

func (o DeinterlacePtrOutput) ToDeinterlacePtrOutputWithContext(ctx context.Context) DeinterlacePtrOutput {
	return o
}

func (o DeinterlacePtrOutput) Elem() DeinterlaceOutput {
	return o.ApplyT(func(v *Deinterlace) Deinterlace {
		if v != nil {
			return *v
		}
		var ret Deinterlace
		return ret
	}).(DeinterlaceOutput)
}

func (o DeinterlacePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deinterlace) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o DeinterlacePtrOutput) Parity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deinterlace) *string {
		if v == nil {
			return nil
		}
		return v.Parity
	}).(pulumi.StringPtrOutput)
}

type DeinterlaceResponse struct {
	Mode   *string `pulumi:"mode"`
	Parity *string `pulumi:"parity"`
}





type DeinterlaceResponseInput interface {
	pulumi.Input

	ToDeinterlaceResponseOutput() DeinterlaceResponseOutput
	ToDeinterlaceResponseOutputWithContext(context.Context) DeinterlaceResponseOutput
}

type DeinterlaceResponseArgs struct {
	Mode   pulumi.StringPtrInput `pulumi:"mode"`
	Parity pulumi.StringPtrInput `pulumi:"parity"`
}

func (DeinterlaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeinterlaceResponse)(nil)).Elem()
}

func (i DeinterlaceResponseArgs) ToDeinterlaceResponseOutput() DeinterlaceResponseOutput {
	return i.ToDeinterlaceResponseOutputWithContext(context.Background())
}

func (i DeinterlaceResponseArgs) ToDeinterlaceResponseOutputWithContext(ctx context.Context) DeinterlaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeinterlaceResponseOutput)
}

func (i DeinterlaceResponseArgs) ToDeinterlaceResponsePtrOutput() DeinterlaceResponsePtrOutput {
	return i.ToDeinterlaceResponsePtrOutputWithContext(context.Background())
}

func (i DeinterlaceResponseArgs) ToDeinterlaceResponsePtrOutputWithContext(ctx context.Context) DeinterlaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeinterlaceResponseOutput).ToDeinterlaceResponsePtrOutputWithContext(ctx)
}









type DeinterlaceResponsePtrInput interface {
	pulumi.Input

	ToDeinterlaceResponsePtrOutput() DeinterlaceResponsePtrOutput
	ToDeinterlaceResponsePtrOutputWithContext(context.Context) DeinterlaceResponsePtrOutput
}

type deinterlaceResponsePtrType DeinterlaceResponseArgs

func DeinterlaceResponsePtr(v *DeinterlaceResponseArgs) DeinterlaceResponsePtrInput {
	return (*deinterlaceResponsePtrType)(v)
}

func (*deinterlaceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeinterlaceResponse)(nil)).Elem()
}

func (i *deinterlaceResponsePtrType) ToDeinterlaceResponsePtrOutput() DeinterlaceResponsePtrOutput {
	return i.ToDeinterlaceResponsePtrOutputWithContext(context.Background())
}

func (i *deinterlaceResponsePtrType) ToDeinterlaceResponsePtrOutputWithContext(ctx context.Context) DeinterlaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeinterlaceResponsePtrOutput)
}

type DeinterlaceResponseOutput struct{ *pulumi.OutputState }

func (DeinterlaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeinterlaceResponse)(nil)).Elem()
}

func (o DeinterlaceResponseOutput) ToDeinterlaceResponseOutput() DeinterlaceResponseOutput {
	return o
}

func (o DeinterlaceResponseOutput) ToDeinterlaceResponseOutputWithContext(ctx context.Context) DeinterlaceResponseOutput {
	return o
}

func (o DeinterlaceResponseOutput) ToDeinterlaceResponsePtrOutput() DeinterlaceResponsePtrOutput {
	return o.ToDeinterlaceResponsePtrOutputWithContext(context.Background())
}

func (o DeinterlaceResponseOutput) ToDeinterlaceResponsePtrOutputWithContext(ctx context.Context) DeinterlaceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeinterlaceResponse) *DeinterlaceResponse {
		return &v
	}).(DeinterlaceResponsePtrOutput)
}

func (o DeinterlaceResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeinterlaceResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o DeinterlaceResponseOutput) Parity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeinterlaceResponse) *string { return v.Parity }).(pulumi.StringPtrOutput)
}

type DeinterlaceResponsePtrOutput struct{ *pulumi.OutputState }

func (DeinterlaceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeinterlaceResponse)(nil)).Elem()
}

func (o DeinterlaceResponsePtrOutput) ToDeinterlaceResponsePtrOutput() DeinterlaceResponsePtrOutput {
	return o
}

func (o DeinterlaceResponsePtrOutput) ToDeinterlaceResponsePtrOutputWithContext(ctx context.Context) DeinterlaceResponsePtrOutput {
	return o
}

func (o DeinterlaceResponsePtrOutput) Elem() DeinterlaceResponseOutput {
	return o.ApplyT(func(v *DeinterlaceResponse) DeinterlaceResponse {
		if v != nil {
			return *v
		}
		var ret DeinterlaceResponse
		return ret
	}).(DeinterlaceResponseOutput)
}

func (o DeinterlaceResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeinterlaceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o DeinterlaceResponsePtrOutput) Parity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeinterlaceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Parity
	}).(pulumi.StringPtrOutput)
}

type EdgeUsageDataCollectionPolicyResponse struct {
	DataCollectionFrequency           *string                        `pulumi:"dataCollectionFrequency"`
	DataReportingFrequency            *string                        `pulumi:"dataReportingFrequency"`
	EventHubDetails                   *EdgeUsageDataEventHubResponse `pulumi:"eventHubDetails"`
	MaxAllowedUnreportedUsageDuration *string                        `pulumi:"maxAllowedUnreportedUsageDuration"`
}





type EdgeUsageDataCollectionPolicyResponseInput interface {
	pulumi.Input

	ToEdgeUsageDataCollectionPolicyResponseOutput() EdgeUsageDataCollectionPolicyResponseOutput
	ToEdgeUsageDataCollectionPolicyResponseOutputWithContext(context.Context) EdgeUsageDataCollectionPolicyResponseOutput
}

type EdgeUsageDataCollectionPolicyResponseArgs struct {
	DataCollectionFrequency           pulumi.StringPtrInput                 `pulumi:"dataCollectionFrequency"`
	DataReportingFrequency            pulumi.StringPtrInput                 `pulumi:"dataReportingFrequency"`
	EventHubDetails                   EdgeUsageDataEventHubResponsePtrInput `pulumi:"eventHubDetails"`
	MaxAllowedUnreportedUsageDuration pulumi.StringPtrInput                 `pulumi:"maxAllowedUnreportedUsageDuration"`
}

func (EdgeUsageDataCollectionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeUsageDataCollectionPolicyResponse)(nil)).Elem()
}

func (i EdgeUsageDataCollectionPolicyResponseArgs) ToEdgeUsageDataCollectionPolicyResponseOutput() EdgeUsageDataCollectionPolicyResponseOutput {
	return i.ToEdgeUsageDataCollectionPolicyResponseOutputWithContext(context.Background())
}

func (i EdgeUsageDataCollectionPolicyResponseArgs) ToEdgeUsageDataCollectionPolicyResponseOutputWithContext(ctx context.Context) EdgeUsageDataCollectionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeUsageDataCollectionPolicyResponseOutput)
}

func (i EdgeUsageDataCollectionPolicyResponseArgs) ToEdgeUsageDataCollectionPolicyResponsePtrOutput() EdgeUsageDataCollectionPolicyResponsePtrOutput {
	return i.ToEdgeUsageDataCollectionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i EdgeUsageDataCollectionPolicyResponseArgs) ToEdgeUsageDataCollectionPolicyResponsePtrOutputWithContext(ctx context.Context) EdgeUsageDataCollectionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeUsageDataCollectionPolicyResponseOutput).ToEdgeUsageDataCollectionPolicyResponsePtrOutputWithContext(ctx)
}









type EdgeUsageDataCollectionPolicyResponsePtrInput interface {
	pulumi.Input

	ToEdgeUsageDataCollectionPolicyResponsePtrOutput() EdgeUsageDataCollectionPolicyResponsePtrOutput
	ToEdgeUsageDataCollectionPolicyResponsePtrOutputWithContext(context.Context) EdgeUsageDataCollectionPolicyResponsePtrOutput
}

type edgeUsageDataCollectionPolicyResponsePtrType EdgeUsageDataCollectionPolicyResponseArgs

func EdgeUsageDataCollectionPolicyResponsePtr(v *EdgeUsageDataCollectionPolicyResponseArgs) EdgeUsageDataCollectionPolicyResponsePtrInput {
	return (*edgeUsageDataCollectionPolicyResponsePtrType)(v)
}

func (*edgeUsageDataCollectionPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeUsageDataCollectionPolicyResponse)(nil)).Elem()
}

func (i *edgeUsageDataCollectionPolicyResponsePtrType) ToEdgeUsageDataCollectionPolicyResponsePtrOutput() EdgeUsageDataCollectionPolicyResponsePtrOutput {
	return i.ToEdgeUsageDataCollectionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *edgeUsageDataCollectionPolicyResponsePtrType) ToEdgeUsageDataCollectionPolicyResponsePtrOutputWithContext(ctx context.Context) EdgeUsageDataCollectionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeUsageDataCollectionPolicyResponsePtrOutput)
}

type EdgeUsageDataCollectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (EdgeUsageDataCollectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeUsageDataCollectionPolicyResponse)(nil)).Elem()
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) ToEdgeUsageDataCollectionPolicyResponseOutput() EdgeUsageDataCollectionPolicyResponseOutput {
	return o
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) ToEdgeUsageDataCollectionPolicyResponseOutputWithContext(ctx context.Context) EdgeUsageDataCollectionPolicyResponseOutput {
	return o
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) ToEdgeUsageDataCollectionPolicyResponsePtrOutput() EdgeUsageDataCollectionPolicyResponsePtrOutput {
	return o.ToEdgeUsageDataCollectionPolicyResponsePtrOutputWithContext(context.Background())
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) ToEdgeUsageDataCollectionPolicyResponsePtrOutputWithContext(ctx context.Context) EdgeUsageDataCollectionPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdgeUsageDataCollectionPolicyResponse) *EdgeUsageDataCollectionPolicyResponse {
		return &v
	}).(EdgeUsageDataCollectionPolicyResponsePtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) DataCollectionFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataCollectionPolicyResponse) *string { return v.DataCollectionFrequency }).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) DataReportingFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataCollectionPolicyResponse) *string { return v.DataReportingFrequency }).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) EventHubDetails() EdgeUsageDataEventHubResponsePtrOutput {
	return o.ApplyT(func(v EdgeUsageDataCollectionPolicyResponse) *EdgeUsageDataEventHubResponse { return v.EventHubDetails }).(EdgeUsageDataEventHubResponsePtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponseOutput) MaxAllowedUnreportedUsageDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataCollectionPolicyResponse) *string { return v.MaxAllowedUnreportedUsageDuration }).(pulumi.StringPtrOutput)
}

type EdgeUsageDataCollectionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (EdgeUsageDataCollectionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeUsageDataCollectionPolicyResponse)(nil)).Elem()
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) ToEdgeUsageDataCollectionPolicyResponsePtrOutput() EdgeUsageDataCollectionPolicyResponsePtrOutput {
	return o
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) ToEdgeUsageDataCollectionPolicyResponsePtrOutputWithContext(ctx context.Context) EdgeUsageDataCollectionPolicyResponsePtrOutput {
	return o
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) Elem() EdgeUsageDataCollectionPolicyResponseOutput {
	return o.ApplyT(func(v *EdgeUsageDataCollectionPolicyResponse) EdgeUsageDataCollectionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret EdgeUsageDataCollectionPolicyResponse
		return ret
	}).(EdgeUsageDataCollectionPolicyResponseOutput)
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) DataCollectionFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataCollectionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.DataCollectionFrequency
	}).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) DataReportingFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataCollectionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.DataReportingFrequency
	}).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) EventHubDetails() EdgeUsageDataEventHubResponsePtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataCollectionPolicyResponse) *EdgeUsageDataEventHubResponse {
		if v == nil {
			return nil
		}
		return v.EventHubDetails
	}).(EdgeUsageDataEventHubResponsePtrOutput)
}

func (o EdgeUsageDataCollectionPolicyResponsePtrOutput) MaxAllowedUnreportedUsageDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataCollectionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaxAllowedUnreportedUsageDuration
	}).(pulumi.StringPtrOutput)
}

type EdgeUsageDataEventHubResponse struct {
	Name      *string `pulumi:"name"`
	Namespace *string `pulumi:"namespace"`
	Token     *string `pulumi:"token"`
}





type EdgeUsageDataEventHubResponseInput interface {
	pulumi.Input

	ToEdgeUsageDataEventHubResponseOutput() EdgeUsageDataEventHubResponseOutput
	ToEdgeUsageDataEventHubResponseOutputWithContext(context.Context) EdgeUsageDataEventHubResponseOutput
}

type EdgeUsageDataEventHubResponseArgs struct {
	Name      pulumi.StringPtrInput `pulumi:"name"`
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	Token     pulumi.StringPtrInput `pulumi:"token"`
}

func (EdgeUsageDataEventHubResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeUsageDataEventHubResponse)(nil)).Elem()
}

func (i EdgeUsageDataEventHubResponseArgs) ToEdgeUsageDataEventHubResponseOutput() EdgeUsageDataEventHubResponseOutput {
	return i.ToEdgeUsageDataEventHubResponseOutputWithContext(context.Background())
}

func (i EdgeUsageDataEventHubResponseArgs) ToEdgeUsageDataEventHubResponseOutputWithContext(ctx context.Context) EdgeUsageDataEventHubResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeUsageDataEventHubResponseOutput)
}

func (i EdgeUsageDataEventHubResponseArgs) ToEdgeUsageDataEventHubResponsePtrOutput() EdgeUsageDataEventHubResponsePtrOutput {
	return i.ToEdgeUsageDataEventHubResponsePtrOutputWithContext(context.Background())
}

func (i EdgeUsageDataEventHubResponseArgs) ToEdgeUsageDataEventHubResponsePtrOutputWithContext(ctx context.Context) EdgeUsageDataEventHubResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeUsageDataEventHubResponseOutput).ToEdgeUsageDataEventHubResponsePtrOutputWithContext(ctx)
}









type EdgeUsageDataEventHubResponsePtrInput interface {
	pulumi.Input

	ToEdgeUsageDataEventHubResponsePtrOutput() EdgeUsageDataEventHubResponsePtrOutput
	ToEdgeUsageDataEventHubResponsePtrOutputWithContext(context.Context) EdgeUsageDataEventHubResponsePtrOutput
}

type edgeUsageDataEventHubResponsePtrType EdgeUsageDataEventHubResponseArgs

func EdgeUsageDataEventHubResponsePtr(v *EdgeUsageDataEventHubResponseArgs) EdgeUsageDataEventHubResponsePtrInput {
	return (*edgeUsageDataEventHubResponsePtrType)(v)
}

func (*edgeUsageDataEventHubResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeUsageDataEventHubResponse)(nil)).Elem()
}

func (i *edgeUsageDataEventHubResponsePtrType) ToEdgeUsageDataEventHubResponsePtrOutput() EdgeUsageDataEventHubResponsePtrOutput {
	return i.ToEdgeUsageDataEventHubResponsePtrOutputWithContext(context.Background())
}

func (i *edgeUsageDataEventHubResponsePtrType) ToEdgeUsageDataEventHubResponsePtrOutputWithContext(ctx context.Context) EdgeUsageDataEventHubResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeUsageDataEventHubResponsePtrOutput)
}

type EdgeUsageDataEventHubResponseOutput struct{ *pulumi.OutputState }

func (EdgeUsageDataEventHubResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeUsageDataEventHubResponse)(nil)).Elem()
}

func (o EdgeUsageDataEventHubResponseOutput) ToEdgeUsageDataEventHubResponseOutput() EdgeUsageDataEventHubResponseOutput {
	return o
}

func (o EdgeUsageDataEventHubResponseOutput) ToEdgeUsageDataEventHubResponseOutputWithContext(ctx context.Context) EdgeUsageDataEventHubResponseOutput {
	return o
}

func (o EdgeUsageDataEventHubResponseOutput) ToEdgeUsageDataEventHubResponsePtrOutput() EdgeUsageDataEventHubResponsePtrOutput {
	return o.ToEdgeUsageDataEventHubResponsePtrOutputWithContext(context.Background())
}

func (o EdgeUsageDataEventHubResponseOutput) ToEdgeUsageDataEventHubResponsePtrOutputWithContext(ctx context.Context) EdgeUsageDataEventHubResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdgeUsageDataEventHubResponse) *EdgeUsageDataEventHubResponse {
		return &v
	}).(EdgeUsageDataEventHubResponsePtrOutput)
}

func (o EdgeUsageDataEventHubResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataEventHubResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataEventHubResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataEventHubResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataEventHubResponseOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdgeUsageDataEventHubResponse) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type EdgeUsageDataEventHubResponsePtrOutput struct{ *pulumi.OutputState }

func (EdgeUsageDataEventHubResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeUsageDataEventHubResponse)(nil)).Elem()
}

func (o EdgeUsageDataEventHubResponsePtrOutput) ToEdgeUsageDataEventHubResponsePtrOutput() EdgeUsageDataEventHubResponsePtrOutput {
	return o
}

func (o EdgeUsageDataEventHubResponsePtrOutput) ToEdgeUsageDataEventHubResponsePtrOutputWithContext(ctx context.Context) EdgeUsageDataEventHubResponsePtrOutput {
	return o
}

func (o EdgeUsageDataEventHubResponsePtrOutput) Elem() EdgeUsageDataEventHubResponseOutput {
	return o.ApplyT(func(v *EdgeUsageDataEventHubResponse) EdgeUsageDataEventHubResponse {
		if v != nil {
			return *v
		}
		var ret EdgeUsageDataEventHubResponse
		return ret
	}).(EdgeUsageDataEventHubResponseOutput)
}

func (o EdgeUsageDataEventHubResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataEventHubResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataEventHubResponsePtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataEventHubResponse) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

func (o EdgeUsageDataEventHubResponsePtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeUsageDataEventHubResponse) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

type EnabledProtocols struct {
	Dash            bool `pulumi:"dash"`
	Download        bool `pulumi:"download"`
	Hls             bool `pulumi:"hls"`
	SmoothStreaming bool `pulumi:"smoothStreaming"`
}





type EnabledProtocolsInput interface {
	pulumi.Input

	ToEnabledProtocolsOutput() EnabledProtocolsOutput
	ToEnabledProtocolsOutputWithContext(context.Context) EnabledProtocolsOutput
}

type EnabledProtocolsArgs struct {
	Dash            pulumi.BoolInput `pulumi:"dash"`
	Download        pulumi.BoolInput `pulumi:"download"`
	Hls             pulumi.BoolInput `pulumi:"hls"`
	SmoothStreaming pulumi.BoolInput `pulumi:"smoothStreaming"`
}

func (EnabledProtocolsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledProtocols)(nil)).Elem()
}

func (i EnabledProtocolsArgs) ToEnabledProtocolsOutput() EnabledProtocolsOutput {
	return i.ToEnabledProtocolsOutputWithContext(context.Background())
}

func (i EnabledProtocolsArgs) ToEnabledProtocolsOutputWithContext(ctx context.Context) EnabledProtocolsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledProtocolsOutput)
}

func (i EnabledProtocolsArgs) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return i.ToEnabledProtocolsPtrOutputWithContext(context.Background())
}

func (i EnabledProtocolsArgs) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledProtocolsOutput).ToEnabledProtocolsPtrOutputWithContext(ctx)
}









type EnabledProtocolsPtrInput interface {
	pulumi.Input

	ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput
	ToEnabledProtocolsPtrOutputWithContext(context.Context) EnabledProtocolsPtrOutput
}

type enabledProtocolsPtrType EnabledProtocolsArgs

func EnabledProtocolsPtr(v *EnabledProtocolsArgs) EnabledProtocolsPtrInput {
	return (*enabledProtocolsPtrType)(v)
}

func (*enabledProtocolsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledProtocols)(nil)).Elem()
}

func (i *enabledProtocolsPtrType) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return i.ToEnabledProtocolsPtrOutputWithContext(context.Background())
}

func (i *enabledProtocolsPtrType) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledProtocolsPtrOutput)
}

type EnabledProtocolsOutput struct{ *pulumi.OutputState }

func (EnabledProtocolsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledProtocols)(nil)).Elem()
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsOutput() EnabledProtocolsOutput {
	return o
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsOutputWithContext(ctx context.Context) EnabledProtocolsOutput {
	return o
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return o.ToEnabledProtocolsPtrOutputWithContext(context.Background())
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnabledProtocols) *EnabledProtocols {
		return &v
	}).(EnabledProtocolsPtrOutput)
}

func (o EnabledProtocolsOutput) Dash() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocols) bool { return v.Dash }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsOutput) Download() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocols) bool { return v.Download }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsOutput) Hls() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocols) bool { return v.Hls }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsOutput) SmoothStreaming() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocols) bool { return v.SmoothStreaming }).(pulumi.BoolOutput)
}

type EnabledProtocolsPtrOutput struct{ *pulumi.OutputState }

func (EnabledProtocolsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledProtocols)(nil)).Elem()
}

func (o EnabledProtocolsPtrOutput) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return o
}

func (o EnabledProtocolsPtrOutput) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return o
}

func (o EnabledProtocolsPtrOutput) Elem() EnabledProtocolsOutput {
	return o.ApplyT(func(v *EnabledProtocols) EnabledProtocols {
		if v != nil {
			return *v
		}
		var ret EnabledProtocols
		return ret
	}).(EnabledProtocolsOutput)
}

func (o EnabledProtocolsPtrOutput) Dash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocols) *bool {
		if v == nil {
			return nil
		}
		return &v.Dash
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsPtrOutput) Download() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocols) *bool {
		if v == nil {
			return nil
		}
		return &v.Download
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsPtrOutput) Hls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocols) *bool {
		if v == nil {
			return nil
		}
		return &v.Hls
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsPtrOutput) SmoothStreaming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocols) *bool {
		if v == nil {
			return nil
		}
		return &v.SmoothStreaming
	}).(pulumi.BoolPtrOutput)
}

type EnabledProtocolsResponse struct {
	Dash            bool `pulumi:"dash"`
	Download        bool `pulumi:"download"`
	Hls             bool `pulumi:"hls"`
	SmoothStreaming bool `pulumi:"smoothStreaming"`
}





type EnabledProtocolsResponseInput interface {
	pulumi.Input

	ToEnabledProtocolsResponseOutput() EnabledProtocolsResponseOutput
	ToEnabledProtocolsResponseOutputWithContext(context.Context) EnabledProtocolsResponseOutput
}

type EnabledProtocolsResponseArgs struct {
	Dash            pulumi.BoolInput `pulumi:"dash"`
	Download        pulumi.BoolInput `pulumi:"download"`
	Hls             pulumi.BoolInput `pulumi:"hls"`
	SmoothStreaming pulumi.BoolInput `pulumi:"smoothStreaming"`
}

func (EnabledProtocolsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledProtocolsResponse)(nil)).Elem()
}

func (i EnabledProtocolsResponseArgs) ToEnabledProtocolsResponseOutput() EnabledProtocolsResponseOutput {
	return i.ToEnabledProtocolsResponseOutputWithContext(context.Background())
}

func (i EnabledProtocolsResponseArgs) ToEnabledProtocolsResponseOutputWithContext(ctx context.Context) EnabledProtocolsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledProtocolsResponseOutput)
}

func (i EnabledProtocolsResponseArgs) ToEnabledProtocolsResponsePtrOutput() EnabledProtocolsResponsePtrOutput {
	return i.ToEnabledProtocolsResponsePtrOutputWithContext(context.Background())
}

func (i EnabledProtocolsResponseArgs) ToEnabledProtocolsResponsePtrOutputWithContext(ctx context.Context) EnabledProtocolsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledProtocolsResponseOutput).ToEnabledProtocolsResponsePtrOutputWithContext(ctx)
}









type EnabledProtocolsResponsePtrInput interface {
	pulumi.Input

	ToEnabledProtocolsResponsePtrOutput() EnabledProtocolsResponsePtrOutput
	ToEnabledProtocolsResponsePtrOutputWithContext(context.Context) EnabledProtocolsResponsePtrOutput
}

type enabledProtocolsResponsePtrType EnabledProtocolsResponseArgs

func EnabledProtocolsResponsePtr(v *EnabledProtocolsResponseArgs) EnabledProtocolsResponsePtrInput {
	return (*enabledProtocolsResponsePtrType)(v)
}

func (*enabledProtocolsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledProtocolsResponse)(nil)).Elem()
}

func (i *enabledProtocolsResponsePtrType) ToEnabledProtocolsResponsePtrOutput() EnabledProtocolsResponsePtrOutput {
	return i.ToEnabledProtocolsResponsePtrOutputWithContext(context.Background())
}

func (i *enabledProtocolsResponsePtrType) ToEnabledProtocolsResponsePtrOutputWithContext(ctx context.Context) EnabledProtocolsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledProtocolsResponsePtrOutput)
}

type EnabledProtocolsResponseOutput struct{ *pulumi.OutputState }

func (EnabledProtocolsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledProtocolsResponse)(nil)).Elem()
}

func (o EnabledProtocolsResponseOutput) ToEnabledProtocolsResponseOutput() EnabledProtocolsResponseOutput {
	return o
}

func (o EnabledProtocolsResponseOutput) ToEnabledProtocolsResponseOutputWithContext(ctx context.Context) EnabledProtocolsResponseOutput {
	return o
}

func (o EnabledProtocolsResponseOutput) ToEnabledProtocolsResponsePtrOutput() EnabledProtocolsResponsePtrOutput {
	return o.ToEnabledProtocolsResponsePtrOutputWithContext(context.Background())
}

func (o EnabledProtocolsResponseOutput) ToEnabledProtocolsResponsePtrOutputWithContext(ctx context.Context) EnabledProtocolsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnabledProtocolsResponse) *EnabledProtocolsResponse {
		return &v
	}).(EnabledProtocolsResponsePtrOutput)
}

func (o EnabledProtocolsResponseOutput) Dash() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocolsResponse) bool { return v.Dash }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsResponseOutput) Download() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocolsResponse) bool { return v.Download }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsResponseOutput) Hls() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocolsResponse) bool { return v.Hls }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsResponseOutput) SmoothStreaming() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocolsResponse) bool { return v.SmoothStreaming }).(pulumi.BoolOutput)
}

type EnabledProtocolsResponsePtrOutput struct{ *pulumi.OutputState }

func (EnabledProtocolsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledProtocolsResponse)(nil)).Elem()
}

func (o EnabledProtocolsResponsePtrOutput) ToEnabledProtocolsResponsePtrOutput() EnabledProtocolsResponsePtrOutput {
	return o
}

func (o EnabledProtocolsResponsePtrOutput) ToEnabledProtocolsResponsePtrOutputWithContext(ctx context.Context) EnabledProtocolsResponsePtrOutput {
	return o
}

func (o EnabledProtocolsResponsePtrOutput) Elem() EnabledProtocolsResponseOutput {
	return o.ApplyT(func(v *EnabledProtocolsResponse) EnabledProtocolsResponse {
		if v != nil {
			return *v
		}
		var ret EnabledProtocolsResponse
		return ret
	}).(EnabledProtocolsResponseOutput)
}

func (o EnabledProtocolsResponsePtrOutput) Dash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocolsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Dash
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsResponsePtrOutput) Download() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocolsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Download
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsResponsePtrOutput) Hls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocolsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Hls
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsResponsePtrOutput) SmoothStreaming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocolsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SmoothStreaming
	}).(pulumi.BoolPtrOutput)
}

type EnvelopeEncryption struct {
	ClearTracks                     []TrackSelection            `pulumi:"clearTracks"`
	ContentKeys                     *StreamingPolicyContentKeys `pulumi:"contentKeys"`
	CustomKeyAcquisitionUrlTemplate *string                     `pulumi:"customKeyAcquisitionUrlTemplate"`
	EnabledProtocols                *EnabledProtocols           `pulumi:"enabledProtocols"`
}





type EnvelopeEncryptionInput interface {
	pulumi.Input

	ToEnvelopeEncryptionOutput() EnvelopeEncryptionOutput
	ToEnvelopeEncryptionOutputWithContext(context.Context) EnvelopeEncryptionOutput
}

type EnvelopeEncryptionArgs struct {
	ClearTracks                     TrackSelectionArrayInput           `pulumi:"clearTracks"`
	ContentKeys                     StreamingPolicyContentKeysPtrInput `pulumi:"contentKeys"`
	CustomKeyAcquisitionUrlTemplate pulumi.StringPtrInput              `pulumi:"customKeyAcquisitionUrlTemplate"`
	EnabledProtocols                EnabledProtocolsPtrInput           `pulumi:"enabledProtocols"`
}

func (EnvelopeEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvelopeEncryption)(nil)).Elem()
}

func (i EnvelopeEncryptionArgs) ToEnvelopeEncryptionOutput() EnvelopeEncryptionOutput {
	return i.ToEnvelopeEncryptionOutputWithContext(context.Background())
}

func (i EnvelopeEncryptionArgs) ToEnvelopeEncryptionOutputWithContext(ctx context.Context) EnvelopeEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvelopeEncryptionOutput)
}

func (i EnvelopeEncryptionArgs) ToEnvelopeEncryptionPtrOutput() EnvelopeEncryptionPtrOutput {
	return i.ToEnvelopeEncryptionPtrOutputWithContext(context.Background())
}

func (i EnvelopeEncryptionArgs) ToEnvelopeEncryptionPtrOutputWithContext(ctx context.Context) EnvelopeEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvelopeEncryptionOutput).ToEnvelopeEncryptionPtrOutputWithContext(ctx)
}









type EnvelopeEncryptionPtrInput interface {
	pulumi.Input

	ToEnvelopeEncryptionPtrOutput() EnvelopeEncryptionPtrOutput
	ToEnvelopeEncryptionPtrOutputWithContext(context.Context) EnvelopeEncryptionPtrOutput
}

type envelopeEncryptionPtrType EnvelopeEncryptionArgs

func EnvelopeEncryptionPtr(v *EnvelopeEncryptionArgs) EnvelopeEncryptionPtrInput {
	return (*envelopeEncryptionPtrType)(v)
}

func (*envelopeEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvelopeEncryption)(nil)).Elem()
}

func (i *envelopeEncryptionPtrType) ToEnvelopeEncryptionPtrOutput() EnvelopeEncryptionPtrOutput {
	return i.ToEnvelopeEncryptionPtrOutputWithContext(context.Background())
}

func (i *envelopeEncryptionPtrType) ToEnvelopeEncryptionPtrOutputWithContext(ctx context.Context) EnvelopeEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvelopeEncryptionPtrOutput)
}

type EnvelopeEncryptionOutput struct{ *pulumi.OutputState }

func (EnvelopeEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvelopeEncryption)(nil)).Elem()
}

func (o EnvelopeEncryptionOutput) ToEnvelopeEncryptionOutput() EnvelopeEncryptionOutput {
	return o
}

func (o EnvelopeEncryptionOutput) ToEnvelopeEncryptionOutputWithContext(ctx context.Context) EnvelopeEncryptionOutput {
	return o
}

func (o EnvelopeEncryptionOutput) ToEnvelopeEncryptionPtrOutput() EnvelopeEncryptionPtrOutput {
	return o.ToEnvelopeEncryptionPtrOutputWithContext(context.Background())
}

func (o EnvelopeEncryptionOutput) ToEnvelopeEncryptionPtrOutputWithContext(ctx context.Context) EnvelopeEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvelopeEncryption) *EnvelopeEncryption {
		return &v
	}).(EnvelopeEncryptionPtrOutput)
}

func (o EnvelopeEncryptionOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v EnvelopeEncryption) []TrackSelection { return v.ClearTracks }).(TrackSelectionArrayOutput)
}

func (o EnvelopeEncryptionOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v EnvelopeEncryption) *StreamingPolicyContentKeys { return v.ContentKeys }).(StreamingPolicyContentKeysPtrOutput)
}

func (o EnvelopeEncryptionOutput) CustomKeyAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvelopeEncryption) *string { return v.CustomKeyAcquisitionUrlTemplate }).(pulumi.StringPtrOutput)
}

func (o EnvelopeEncryptionOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v EnvelopeEncryption) *EnabledProtocols { return v.EnabledProtocols }).(EnabledProtocolsPtrOutput)
}

type EnvelopeEncryptionPtrOutput struct{ *pulumi.OutputState }

func (EnvelopeEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvelopeEncryption)(nil)).Elem()
}

func (o EnvelopeEncryptionPtrOutput) ToEnvelopeEncryptionPtrOutput() EnvelopeEncryptionPtrOutput {
	return o
}

func (o EnvelopeEncryptionPtrOutput) ToEnvelopeEncryptionPtrOutputWithContext(ctx context.Context) EnvelopeEncryptionPtrOutput {
	return o
}

func (o EnvelopeEncryptionPtrOutput) Elem() EnvelopeEncryptionOutput {
	return o.ApplyT(func(v *EnvelopeEncryption) EnvelopeEncryption {
		if v != nil {
			return *v
		}
		var ret EnvelopeEncryption
		return ret
	}).(EnvelopeEncryptionOutput)
}

func (o EnvelopeEncryptionPtrOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v *EnvelopeEncryption) []TrackSelection {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionArrayOutput)
}

func (o EnvelopeEncryptionPtrOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryption) *StreamingPolicyContentKeys {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysPtrOutput)
}

func (o EnvelopeEncryptionPtrOutput) CustomKeyAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryption) *string {
		if v == nil {
			return nil
		}
		return v.CustomKeyAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

func (o EnvelopeEncryptionPtrOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryption) *EnabledProtocols {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsPtrOutput)
}

type EnvelopeEncryptionResponse struct {
	ClearTracks                     []TrackSelectionResponse            `pulumi:"clearTracks"`
	ContentKeys                     *StreamingPolicyContentKeysResponse `pulumi:"contentKeys"`
	CustomKeyAcquisitionUrlTemplate *string                             `pulumi:"customKeyAcquisitionUrlTemplate"`
	EnabledProtocols                *EnabledProtocolsResponse           `pulumi:"enabledProtocols"`
}





type EnvelopeEncryptionResponseInput interface {
	pulumi.Input

	ToEnvelopeEncryptionResponseOutput() EnvelopeEncryptionResponseOutput
	ToEnvelopeEncryptionResponseOutputWithContext(context.Context) EnvelopeEncryptionResponseOutput
}

type EnvelopeEncryptionResponseArgs struct {
	ClearTracks                     TrackSelectionResponseArrayInput           `pulumi:"clearTracks"`
	ContentKeys                     StreamingPolicyContentKeysResponsePtrInput `pulumi:"contentKeys"`
	CustomKeyAcquisitionUrlTemplate pulumi.StringPtrInput                      `pulumi:"customKeyAcquisitionUrlTemplate"`
	EnabledProtocols                EnabledProtocolsResponsePtrInput           `pulumi:"enabledProtocols"`
}

func (EnvelopeEncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvelopeEncryptionResponse)(nil)).Elem()
}

func (i EnvelopeEncryptionResponseArgs) ToEnvelopeEncryptionResponseOutput() EnvelopeEncryptionResponseOutput {
	return i.ToEnvelopeEncryptionResponseOutputWithContext(context.Background())
}

func (i EnvelopeEncryptionResponseArgs) ToEnvelopeEncryptionResponseOutputWithContext(ctx context.Context) EnvelopeEncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvelopeEncryptionResponseOutput)
}

func (i EnvelopeEncryptionResponseArgs) ToEnvelopeEncryptionResponsePtrOutput() EnvelopeEncryptionResponsePtrOutput {
	return i.ToEnvelopeEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i EnvelopeEncryptionResponseArgs) ToEnvelopeEncryptionResponsePtrOutputWithContext(ctx context.Context) EnvelopeEncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvelopeEncryptionResponseOutput).ToEnvelopeEncryptionResponsePtrOutputWithContext(ctx)
}









type EnvelopeEncryptionResponsePtrInput interface {
	pulumi.Input

	ToEnvelopeEncryptionResponsePtrOutput() EnvelopeEncryptionResponsePtrOutput
	ToEnvelopeEncryptionResponsePtrOutputWithContext(context.Context) EnvelopeEncryptionResponsePtrOutput
}

type envelopeEncryptionResponsePtrType EnvelopeEncryptionResponseArgs

func EnvelopeEncryptionResponsePtr(v *EnvelopeEncryptionResponseArgs) EnvelopeEncryptionResponsePtrInput {
	return (*envelopeEncryptionResponsePtrType)(v)
}

func (*envelopeEncryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvelopeEncryptionResponse)(nil)).Elem()
}

func (i *envelopeEncryptionResponsePtrType) ToEnvelopeEncryptionResponsePtrOutput() EnvelopeEncryptionResponsePtrOutput {
	return i.ToEnvelopeEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *envelopeEncryptionResponsePtrType) ToEnvelopeEncryptionResponsePtrOutputWithContext(ctx context.Context) EnvelopeEncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvelopeEncryptionResponsePtrOutput)
}

type EnvelopeEncryptionResponseOutput struct{ *pulumi.OutputState }

func (EnvelopeEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvelopeEncryptionResponse)(nil)).Elem()
}

func (o EnvelopeEncryptionResponseOutput) ToEnvelopeEncryptionResponseOutput() EnvelopeEncryptionResponseOutput {
	return o
}

func (o EnvelopeEncryptionResponseOutput) ToEnvelopeEncryptionResponseOutputWithContext(ctx context.Context) EnvelopeEncryptionResponseOutput {
	return o
}

func (o EnvelopeEncryptionResponseOutput) ToEnvelopeEncryptionResponsePtrOutput() EnvelopeEncryptionResponsePtrOutput {
	return o.ToEnvelopeEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o EnvelopeEncryptionResponseOutput) ToEnvelopeEncryptionResponsePtrOutputWithContext(ctx context.Context) EnvelopeEncryptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvelopeEncryptionResponse) *EnvelopeEncryptionResponse {
		return &v
	}).(EnvelopeEncryptionResponsePtrOutput)
}

func (o EnvelopeEncryptionResponseOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v EnvelopeEncryptionResponse) []TrackSelectionResponse { return v.ClearTracks }).(TrackSelectionResponseArrayOutput)
}

func (o EnvelopeEncryptionResponseOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v EnvelopeEncryptionResponse) *StreamingPolicyContentKeysResponse { return v.ContentKeys }).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o EnvelopeEncryptionResponseOutput) CustomKeyAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvelopeEncryptionResponse) *string { return v.CustomKeyAcquisitionUrlTemplate }).(pulumi.StringPtrOutput)
}

func (o EnvelopeEncryptionResponseOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v EnvelopeEncryptionResponse) *EnabledProtocolsResponse { return v.EnabledProtocols }).(EnabledProtocolsResponsePtrOutput)
}

type EnvelopeEncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EnvelopeEncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvelopeEncryptionResponse)(nil)).Elem()
}

func (o EnvelopeEncryptionResponsePtrOutput) ToEnvelopeEncryptionResponsePtrOutput() EnvelopeEncryptionResponsePtrOutput {
	return o
}

func (o EnvelopeEncryptionResponsePtrOutput) ToEnvelopeEncryptionResponsePtrOutputWithContext(ctx context.Context) EnvelopeEncryptionResponsePtrOutput {
	return o
}

func (o EnvelopeEncryptionResponsePtrOutput) Elem() EnvelopeEncryptionResponseOutput {
	return o.ApplyT(func(v *EnvelopeEncryptionResponse) EnvelopeEncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EnvelopeEncryptionResponse
		return ret
	}).(EnvelopeEncryptionResponseOutput)
}

func (o EnvelopeEncryptionResponsePtrOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v *EnvelopeEncryptionResponse) []TrackSelectionResponse {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionResponseArrayOutput)
}

func (o EnvelopeEncryptionResponsePtrOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryptionResponse) *StreamingPolicyContentKeysResponse {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o EnvelopeEncryptionResponsePtrOutput) CustomKeyAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomKeyAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

func (o EnvelopeEncryptionResponsePtrOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryptionResponse) *EnabledProtocolsResponse {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsResponsePtrOutput)
}

type FaceDetectorPreset struct {
	BlurType            *string           `pulumi:"blurType"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
	Resolution          *string           `pulumi:"resolution"`
}





type FaceDetectorPresetInput interface {
	pulumi.Input

	ToFaceDetectorPresetOutput() FaceDetectorPresetOutput
	ToFaceDetectorPresetOutputWithContext(context.Context) FaceDetectorPresetOutput
}

type FaceDetectorPresetArgs struct {
	BlurType            pulumi.StringPtrInput `pulumi:"blurType"`
	ExperimentalOptions pulumi.StringMapInput `pulumi:"experimentalOptions"`
	Mode                pulumi.StringPtrInput `pulumi:"mode"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
	Resolution          pulumi.StringPtrInput `pulumi:"resolution"`
}

func (FaceDetectorPresetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FaceDetectorPreset)(nil)).Elem()
}

func (i FaceDetectorPresetArgs) ToFaceDetectorPresetOutput() FaceDetectorPresetOutput {
	return i.ToFaceDetectorPresetOutputWithContext(context.Background())
}

func (i FaceDetectorPresetArgs) ToFaceDetectorPresetOutputWithContext(ctx context.Context) FaceDetectorPresetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FaceDetectorPresetOutput)
}

type FaceDetectorPresetOutput struct{ *pulumi.OutputState }

func (FaceDetectorPresetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FaceDetectorPreset)(nil)).Elem()
}

func (o FaceDetectorPresetOutput) ToFaceDetectorPresetOutput() FaceDetectorPresetOutput {
	return o
}

func (o FaceDetectorPresetOutput) ToFaceDetectorPresetOutputWithContext(ctx context.Context) FaceDetectorPresetOutput {
	return o
}

func (o FaceDetectorPresetOutput) BlurType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FaceDetectorPreset) *string { return v.BlurType }).(pulumi.StringPtrOutput)
}

func (o FaceDetectorPresetOutput) ExperimentalOptions() pulumi.StringMapOutput {
	return o.ApplyT(func(v FaceDetectorPreset) map[string]string { return v.ExperimentalOptions }).(pulumi.StringMapOutput)
}

func (o FaceDetectorPresetOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FaceDetectorPreset) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o FaceDetectorPresetOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v FaceDetectorPreset) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o FaceDetectorPresetOutput) Resolution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FaceDetectorPreset) *string { return v.Resolution }).(pulumi.StringPtrOutput)
}

type FaceDetectorPresetResponse struct {
	BlurType            *string           `pulumi:"blurType"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
	Resolution          *string           `pulumi:"resolution"`
}





type FaceDetectorPresetResponseInput interface {
	pulumi.Input

	ToFaceDetectorPresetResponseOutput() FaceDetectorPresetResponseOutput
	ToFaceDetectorPresetResponseOutputWithContext(context.Context) FaceDetectorPresetResponseOutput
}

type FaceDetectorPresetResponseArgs struct {
	BlurType            pulumi.StringPtrInput `pulumi:"blurType"`
	ExperimentalOptions pulumi.StringMapInput `pulumi:"experimentalOptions"`
	Mode                pulumi.StringPtrInput `pulumi:"mode"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
	Resolution          pulumi.StringPtrInput `pulumi:"resolution"`
}

func (FaceDetectorPresetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FaceDetectorPresetResponse)(nil)).Elem()
}

func (i FaceDetectorPresetResponseArgs) ToFaceDetectorPresetResponseOutput() FaceDetectorPresetResponseOutput {
	return i.ToFaceDetectorPresetResponseOutputWithContext(context.Background())
}

func (i FaceDetectorPresetResponseArgs) ToFaceDetectorPresetResponseOutputWithContext(ctx context.Context) FaceDetectorPresetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FaceDetectorPresetResponseOutput)
}

type FaceDetectorPresetResponseOutput struct{ *pulumi.OutputState }

func (FaceDetectorPresetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FaceDetectorPresetResponse)(nil)).Elem()
}

func (o FaceDetectorPresetResponseOutput) ToFaceDetectorPresetResponseOutput() FaceDetectorPresetResponseOutput {
	return o
}

func (o FaceDetectorPresetResponseOutput) ToFaceDetectorPresetResponseOutputWithContext(ctx context.Context) FaceDetectorPresetResponseOutput {
	return o
}

func (o FaceDetectorPresetResponseOutput) BlurType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FaceDetectorPresetResponse) *string { return v.BlurType }).(pulumi.StringPtrOutput)
}

func (o FaceDetectorPresetResponseOutput) ExperimentalOptions() pulumi.StringMapOutput {
	return o.ApplyT(func(v FaceDetectorPresetResponse) map[string]string { return v.ExperimentalOptions }).(pulumi.StringMapOutput)
}

func (o FaceDetectorPresetResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FaceDetectorPresetResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o FaceDetectorPresetResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v FaceDetectorPresetResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o FaceDetectorPresetResponseOutput) Resolution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FaceDetectorPresetResponse) *string { return v.Resolution }).(pulumi.StringPtrOutput)
}

type FilterTrackPropertyCondition struct {
	Operation string `pulumi:"operation"`
	Property  string `pulumi:"property"`
	Value     string `pulumi:"value"`
}





type FilterTrackPropertyConditionInput interface {
	pulumi.Input

	ToFilterTrackPropertyConditionOutput() FilterTrackPropertyConditionOutput
	ToFilterTrackPropertyConditionOutputWithContext(context.Context) FilterTrackPropertyConditionOutput
}

type FilterTrackPropertyConditionArgs struct {
	Operation pulumi.StringInput `pulumi:"operation"`
	Property  pulumi.StringInput `pulumi:"property"`
	Value     pulumi.StringInput `pulumi:"value"`
}

func (FilterTrackPropertyConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterTrackPropertyCondition)(nil)).Elem()
}

func (i FilterTrackPropertyConditionArgs) ToFilterTrackPropertyConditionOutput() FilterTrackPropertyConditionOutput {
	return i.ToFilterTrackPropertyConditionOutputWithContext(context.Background())
}

func (i FilterTrackPropertyConditionArgs) ToFilterTrackPropertyConditionOutputWithContext(ctx context.Context) FilterTrackPropertyConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterTrackPropertyConditionOutput)
}





type FilterTrackPropertyConditionArrayInput interface {
	pulumi.Input

	ToFilterTrackPropertyConditionArrayOutput() FilterTrackPropertyConditionArrayOutput
	ToFilterTrackPropertyConditionArrayOutputWithContext(context.Context) FilterTrackPropertyConditionArrayOutput
}

type FilterTrackPropertyConditionArray []FilterTrackPropertyConditionInput

func (FilterTrackPropertyConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterTrackPropertyCondition)(nil)).Elem()
}

func (i FilterTrackPropertyConditionArray) ToFilterTrackPropertyConditionArrayOutput() FilterTrackPropertyConditionArrayOutput {
	return i.ToFilterTrackPropertyConditionArrayOutputWithContext(context.Background())
}

func (i FilterTrackPropertyConditionArray) ToFilterTrackPropertyConditionArrayOutputWithContext(ctx context.Context) FilterTrackPropertyConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterTrackPropertyConditionArrayOutput)
}

type FilterTrackPropertyConditionOutput struct{ *pulumi.OutputState }

func (FilterTrackPropertyConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterTrackPropertyCondition)(nil)).Elem()
}

func (o FilterTrackPropertyConditionOutput) ToFilterTrackPropertyConditionOutput() FilterTrackPropertyConditionOutput {
	return o
}

func (o FilterTrackPropertyConditionOutput) ToFilterTrackPropertyConditionOutputWithContext(ctx context.Context) FilterTrackPropertyConditionOutput {
	return o
}

func (o FilterTrackPropertyConditionOutput) Operation() pulumi.StringOutput {
	return o.ApplyT(func(v FilterTrackPropertyCondition) string { return v.Operation }).(pulumi.StringOutput)
}

func (o FilterTrackPropertyConditionOutput) Property() pulumi.StringOutput {
	return o.ApplyT(func(v FilterTrackPropertyCondition) string { return v.Property }).(pulumi.StringOutput)
}

func (o FilterTrackPropertyConditionOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v FilterTrackPropertyCondition) string { return v.Value }).(pulumi.StringOutput)
}

type FilterTrackPropertyConditionArrayOutput struct{ *pulumi.OutputState }

func (FilterTrackPropertyConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterTrackPropertyCondition)(nil)).Elem()
}

func (o FilterTrackPropertyConditionArrayOutput) ToFilterTrackPropertyConditionArrayOutput() FilterTrackPropertyConditionArrayOutput {
	return o
}

func (o FilterTrackPropertyConditionArrayOutput) ToFilterTrackPropertyConditionArrayOutputWithContext(ctx context.Context) FilterTrackPropertyConditionArrayOutput {
	return o
}

func (o FilterTrackPropertyConditionArrayOutput) Index(i pulumi.IntInput) FilterTrackPropertyConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilterTrackPropertyCondition {
		return vs[0].([]FilterTrackPropertyCondition)[vs[1].(int)]
	}).(FilterTrackPropertyConditionOutput)
}

type FilterTrackPropertyConditionResponse struct {
	Operation string `pulumi:"operation"`
	Property  string `pulumi:"property"`
	Value     string `pulumi:"value"`
}





type FilterTrackPropertyConditionResponseInput interface {
	pulumi.Input

	ToFilterTrackPropertyConditionResponseOutput() FilterTrackPropertyConditionResponseOutput
	ToFilterTrackPropertyConditionResponseOutputWithContext(context.Context) FilterTrackPropertyConditionResponseOutput
}

type FilterTrackPropertyConditionResponseArgs struct {
	Operation pulumi.StringInput `pulumi:"operation"`
	Property  pulumi.StringInput `pulumi:"property"`
	Value     pulumi.StringInput `pulumi:"value"`
}

func (FilterTrackPropertyConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterTrackPropertyConditionResponse)(nil)).Elem()
}

func (i FilterTrackPropertyConditionResponseArgs) ToFilterTrackPropertyConditionResponseOutput() FilterTrackPropertyConditionResponseOutput {
	return i.ToFilterTrackPropertyConditionResponseOutputWithContext(context.Background())
}

func (i FilterTrackPropertyConditionResponseArgs) ToFilterTrackPropertyConditionResponseOutputWithContext(ctx context.Context) FilterTrackPropertyConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterTrackPropertyConditionResponseOutput)
}





type FilterTrackPropertyConditionResponseArrayInput interface {
	pulumi.Input

	ToFilterTrackPropertyConditionResponseArrayOutput() FilterTrackPropertyConditionResponseArrayOutput
	ToFilterTrackPropertyConditionResponseArrayOutputWithContext(context.Context) FilterTrackPropertyConditionResponseArrayOutput
}

type FilterTrackPropertyConditionResponseArray []FilterTrackPropertyConditionResponseInput

func (FilterTrackPropertyConditionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterTrackPropertyConditionResponse)(nil)).Elem()
}

func (i FilterTrackPropertyConditionResponseArray) ToFilterTrackPropertyConditionResponseArrayOutput() FilterTrackPropertyConditionResponseArrayOutput {
	return i.ToFilterTrackPropertyConditionResponseArrayOutputWithContext(context.Background())
}

func (i FilterTrackPropertyConditionResponseArray) ToFilterTrackPropertyConditionResponseArrayOutputWithContext(ctx context.Context) FilterTrackPropertyConditionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterTrackPropertyConditionResponseArrayOutput)
}

type FilterTrackPropertyConditionResponseOutput struct{ *pulumi.OutputState }

func (FilterTrackPropertyConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterTrackPropertyConditionResponse)(nil)).Elem()
}

func (o FilterTrackPropertyConditionResponseOutput) ToFilterTrackPropertyConditionResponseOutput() FilterTrackPropertyConditionResponseOutput {
	return o
}

func (o FilterTrackPropertyConditionResponseOutput) ToFilterTrackPropertyConditionResponseOutputWithContext(ctx context.Context) FilterTrackPropertyConditionResponseOutput {
	return o
}

func (o FilterTrackPropertyConditionResponseOutput) Operation() pulumi.StringOutput {
	return o.ApplyT(func(v FilterTrackPropertyConditionResponse) string { return v.Operation }).(pulumi.StringOutput)
}

func (o FilterTrackPropertyConditionResponseOutput) Property() pulumi.StringOutput {
	return o.ApplyT(func(v FilterTrackPropertyConditionResponse) string { return v.Property }).(pulumi.StringOutput)
}

func (o FilterTrackPropertyConditionResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v FilterTrackPropertyConditionResponse) string { return v.Value }).(pulumi.StringOutput)
}

type FilterTrackPropertyConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (FilterTrackPropertyConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterTrackPropertyConditionResponse)(nil)).Elem()
}

func (o FilterTrackPropertyConditionResponseArrayOutput) ToFilterTrackPropertyConditionResponseArrayOutput() FilterTrackPropertyConditionResponseArrayOutput {
	return o
}

func (o FilterTrackPropertyConditionResponseArrayOutput) ToFilterTrackPropertyConditionResponseArrayOutputWithContext(ctx context.Context) FilterTrackPropertyConditionResponseArrayOutput {
	return o
}

func (o FilterTrackPropertyConditionResponseArrayOutput) Index(i pulumi.IntInput) FilterTrackPropertyConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilterTrackPropertyConditionResponse {
		return vs[0].([]FilterTrackPropertyConditionResponse)[vs[1].(int)]
	}).(FilterTrackPropertyConditionResponseOutput)
}

type FilterTrackSelection struct {
	TrackSelections []FilterTrackPropertyCondition `pulumi:"trackSelections"`
}





type FilterTrackSelectionInput interface {
	pulumi.Input

	ToFilterTrackSelectionOutput() FilterTrackSelectionOutput
	ToFilterTrackSelectionOutputWithContext(context.Context) FilterTrackSelectionOutput
}

type FilterTrackSelectionArgs struct {
	TrackSelections FilterTrackPropertyConditionArrayInput `pulumi:"trackSelections"`
}

func (FilterTrackSelectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterTrackSelection)(nil)).Elem()
}

func (i FilterTrackSelectionArgs) ToFilterTrackSelectionOutput() FilterTrackSelectionOutput {
	return i.ToFilterTrackSelectionOutputWithContext(context.Background())
}

func (i FilterTrackSelectionArgs) ToFilterTrackSelectionOutputWithContext(ctx context.Context) FilterTrackSelectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterTrackSelectionOutput)
}





type FilterTrackSelectionArrayInput interface {
	pulumi.Input

	ToFilterTrackSelectionArrayOutput() FilterTrackSelectionArrayOutput
	ToFilterTrackSelectionArrayOutputWithContext(context.Context) FilterTrackSelectionArrayOutput
}

type FilterTrackSelectionArray []FilterTrackSelectionInput

func (FilterTrackSelectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterTrackSelection)(nil)).Elem()
}

func (i FilterTrackSelectionArray) ToFilterTrackSelectionArrayOutput() FilterTrackSelectionArrayOutput {
	return i.ToFilterTrackSelectionArrayOutputWithContext(context.Background())
}

func (i FilterTrackSelectionArray) ToFilterTrackSelectionArrayOutputWithContext(ctx context.Context) FilterTrackSelectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterTrackSelectionArrayOutput)
}

type FilterTrackSelectionOutput struct{ *pulumi.OutputState }

func (FilterTrackSelectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterTrackSelection)(nil)).Elem()
}

func (o FilterTrackSelectionOutput) ToFilterTrackSelectionOutput() FilterTrackSelectionOutput {
	return o
}

func (o FilterTrackSelectionOutput) ToFilterTrackSelectionOutputWithContext(ctx context.Context) FilterTrackSelectionOutput {
	return o
}

func (o FilterTrackSelectionOutput) TrackSelections() FilterTrackPropertyConditionArrayOutput {
	return o.ApplyT(func(v FilterTrackSelection) []FilterTrackPropertyCondition { return v.TrackSelections }).(FilterTrackPropertyConditionArrayOutput)
}

type FilterTrackSelectionArrayOutput struct{ *pulumi.OutputState }

func (FilterTrackSelectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterTrackSelection)(nil)).Elem()
}

func (o FilterTrackSelectionArrayOutput) ToFilterTrackSelectionArrayOutput() FilterTrackSelectionArrayOutput {
	return o
}

func (o FilterTrackSelectionArrayOutput) ToFilterTrackSelectionArrayOutputWithContext(ctx context.Context) FilterTrackSelectionArrayOutput {
	return o
}

func (o FilterTrackSelectionArrayOutput) Index(i pulumi.IntInput) FilterTrackSelectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilterTrackSelection {
		return vs[0].([]FilterTrackSelection)[vs[1].(int)]
	}).(FilterTrackSelectionOutput)
}

type FilterTrackSelectionResponse struct {
	TrackSelections []FilterTrackPropertyConditionResponse `pulumi:"trackSelections"`
}





type FilterTrackSelectionResponseInput interface {
	pulumi.Input

	ToFilterTrackSelectionResponseOutput() FilterTrackSelectionResponseOutput
	ToFilterTrackSelectionResponseOutputWithContext(context.Context) FilterTrackSelectionResponseOutput
}

type FilterTrackSelectionResponseArgs struct {
	TrackSelections FilterTrackPropertyConditionResponseArrayInput `pulumi:"trackSelections"`
}

func (FilterTrackSelectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterTrackSelectionResponse)(nil)).Elem()
}

func (i FilterTrackSelectionResponseArgs) ToFilterTrackSelectionResponseOutput() FilterTrackSelectionResponseOutput {
	return i.ToFilterTrackSelectionResponseOutputWithContext(context.Background())
}

func (i FilterTrackSelectionResponseArgs) ToFilterTrackSelectionResponseOutputWithContext(ctx context.Context) FilterTrackSelectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterTrackSelectionResponseOutput)
}





type FilterTrackSelectionResponseArrayInput interface {
	pulumi.Input

	ToFilterTrackSelectionResponseArrayOutput() FilterTrackSelectionResponseArrayOutput
	ToFilterTrackSelectionResponseArrayOutputWithContext(context.Context) FilterTrackSelectionResponseArrayOutput
}

type FilterTrackSelectionResponseArray []FilterTrackSelectionResponseInput

func (FilterTrackSelectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterTrackSelectionResponse)(nil)).Elem()
}

func (i FilterTrackSelectionResponseArray) ToFilterTrackSelectionResponseArrayOutput() FilterTrackSelectionResponseArrayOutput {
	return i.ToFilterTrackSelectionResponseArrayOutputWithContext(context.Background())
}

func (i FilterTrackSelectionResponseArray) ToFilterTrackSelectionResponseArrayOutputWithContext(ctx context.Context) FilterTrackSelectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterTrackSelectionResponseArrayOutput)
}

type FilterTrackSelectionResponseOutput struct{ *pulumi.OutputState }

func (FilterTrackSelectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterTrackSelectionResponse)(nil)).Elem()
}

func (o FilterTrackSelectionResponseOutput) ToFilterTrackSelectionResponseOutput() FilterTrackSelectionResponseOutput {
	return o
}

func (o FilterTrackSelectionResponseOutput) ToFilterTrackSelectionResponseOutputWithContext(ctx context.Context) FilterTrackSelectionResponseOutput {
	return o
}

func (o FilterTrackSelectionResponseOutput) TrackSelections() FilterTrackPropertyConditionResponseArrayOutput {
	return o.ApplyT(func(v FilterTrackSelectionResponse) []FilterTrackPropertyConditionResponse { return v.TrackSelections }).(FilterTrackPropertyConditionResponseArrayOutput)
}

type FilterTrackSelectionResponseArrayOutput struct{ *pulumi.OutputState }

func (FilterTrackSelectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterTrackSelectionResponse)(nil)).Elem()
}

func (o FilterTrackSelectionResponseArrayOutput) ToFilterTrackSelectionResponseArrayOutput() FilterTrackSelectionResponseArrayOutput {
	return o
}

func (o FilterTrackSelectionResponseArrayOutput) ToFilterTrackSelectionResponseArrayOutputWithContext(ctx context.Context) FilterTrackSelectionResponseArrayOutput {
	return o
}

func (o FilterTrackSelectionResponseArrayOutput) Index(i pulumi.IntInput) FilterTrackSelectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilterTrackSelectionResponse {
		return vs[0].([]FilterTrackSelectionResponse)[vs[1].(int)]
	}).(FilterTrackSelectionResponseOutput)
}

type Filters struct {
	Crop        *Rectangle    `pulumi:"crop"`
	Deinterlace *Deinterlace  `pulumi:"deinterlace"`
	Overlays    []interface{} `pulumi:"overlays"`
	Rotation    *string       `pulumi:"rotation"`
}





type FiltersInput interface {
	pulumi.Input

	ToFiltersOutput() FiltersOutput
	ToFiltersOutputWithContext(context.Context) FiltersOutput
}

type FiltersArgs struct {
	Crop        RectanglePtrInput     `pulumi:"crop"`
	Deinterlace DeinterlacePtrInput   `pulumi:"deinterlace"`
	Overlays    pulumi.ArrayInput     `pulumi:"overlays"`
	Rotation    pulumi.StringPtrInput `pulumi:"rotation"`
}

func (FiltersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Filters)(nil)).Elem()
}

func (i FiltersArgs) ToFiltersOutput() FiltersOutput {
	return i.ToFiltersOutputWithContext(context.Background())
}

func (i FiltersArgs) ToFiltersOutputWithContext(ctx context.Context) FiltersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FiltersOutput)
}

func (i FiltersArgs) ToFiltersPtrOutput() FiltersPtrOutput {
	return i.ToFiltersPtrOutputWithContext(context.Background())
}

func (i FiltersArgs) ToFiltersPtrOutputWithContext(ctx context.Context) FiltersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FiltersOutput).ToFiltersPtrOutputWithContext(ctx)
}









type FiltersPtrInput interface {
	pulumi.Input

	ToFiltersPtrOutput() FiltersPtrOutput
	ToFiltersPtrOutputWithContext(context.Context) FiltersPtrOutput
}

type filtersPtrType FiltersArgs

func FiltersPtr(v *FiltersArgs) FiltersPtrInput {
	return (*filtersPtrType)(v)
}

func (*filtersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Filters)(nil)).Elem()
}

func (i *filtersPtrType) ToFiltersPtrOutput() FiltersPtrOutput {
	return i.ToFiltersPtrOutputWithContext(context.Background())
}

func (i *filtersPtrType) ToFiltersPtrOutputWithContext(ctx context.Context) FiltersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FiltersPtrOutput)
}

type FiltersOutput struct{ *pulumi.OutputState }

func (FiltersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Filters)(nil)).Elem()
}

func (o FiltersOutput) ToFiltersOutput() FiltersOutput {
	return o
}

func (o FiltersOutput) ToFiltersOutputWithContext(ctx context.Context) FiltersOutput {
	return o
}

func (o FiltersOutput) ToFiltersPtrOutput() FiltersPtrOutput {
	return o.ToFiltersPtrOutputWithContext(context.Background())
}

func (o FiltersOutput) ToFiltersPtrOutputWithContext(ctx context.Context) FiltersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Filters) *Filters {
		return &v
	}).(FiltersPtrOutput)
}

func (o FiltersOutput) Crop() RectanglePtrOutput {
	return o.ApplyT(func(v Filters) *Rectangle { return v.Crop }).(RectanglePtrOutput)
}

func (o FiltersOutput) Deinterlace() DeinterlacePtrOutput {
	return o.ApplyT(func(v Filters) *Deinterlace { return v.Deinterlace }).(DeinterlacePtrOutput)
}

func (o FiltersOutput) Overlays() pulumi.ArrayOutput {
	return o.ApplyT(func(v Filters) []interface{} { return v.Overlays }).(pulumi.ArrayOutput)
}

func (o FiltersOutput) Rotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Filters) *string { return v.Rotation }).(pulumi.StringPtrOutput)
}

type FiltersPtrOutput struct{ *pulumi.OutputState }

func (FiltersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Filters)(nil)).Elem()
}

func (o FiltersPtrOutput) ToFiltersPtrOutput() FiltersPtrOutput {
	return o
}

func (o FiltersPtrOutput) ToFiltersPtrOutputWithContext(ctx context.Context) FiltersPtrOutput {
	return o
}

func (o FiltersPtrOutput) Elem() FiltersOutput {
	return o.ApplyT(func(v *Filters) Filters {
		if v != nil {
			return *v
		}
		var ret Filters
		return ret
	}).(FiltersOutput)
}

func (o FiltersPtrOutput) Crop() RectanglePtrOutput {
	return o.ApplyT(func(v *Filters) *Rectangle {
		if v == nil {
			return nil
		}
		return v.Crop
	}).(RectanglePtrOutput)
}

func (o FiltersPtrOutput) Deinterlace() DeinterlacePtrOutput {
	return o.ApplyT(func(v *Filters) *Deinterlace {
		if v == nil {
			return nil
		}
		return v.Deinterlace
	}).(DeinterlacePtrOutput)
}

func (o FiltersPtrOutput) Overlays() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Filters) []interface{} {
		if v == nil {
			return nil
		}
		return v.Overlays
	}).(pulumi.ArrayOutput)
}

func (o FiltersPtrOutput) Rotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Filters) *string {
		if v == nil {
			return nil
		}
		return v.Rotation
	}).(pulumi.StringPtrOutput)
}

type FiltersResponse struct {
	Crop        *RectangleResponse   `pulumi:"crop"`
	Deinterlace *DeinterlaceResponse `pulumi:"deinterlace"`
	Overlays    []interface{}        `pulumi:"overlays"`
	Rotation    *string              `pulumi:"rotation"`
}





type FiltersResponseInput interface {
	pulumi.Input

	ToFiltersResponseOutput() FiltersResponseOutput
	ToFiltersResponseOutputWithContext(context.Context) FiltersResponseOutput
}

type FiltersResponseArgs struct {
	Crop        RectangleResponsePtrInput   `pulumi:"crop"`
	Deinterlace DeinterlaceResponsePtrInput `pulumi:"deinterlace"`
	Overlays    pulumi.ArrayInput           `pulumi:"overlays"`
	Rotation    pulumi.StringPtrInput       `pulumi:"rotation"`
}

func (FiltersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FiltersResponse)(nil)).Elem()
}

func (i FiltersResponseArgs) ToFiltersResponseOutput() FiltersResponseOutput {
	return i.ToFiltersResponseOutputWithContext(context.Background())
}

func (i FiltersResponseArgs) ToFiltersResponseOutputWithContext(ctx context.Context) FiltersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FiltersResponseOutput)
}

func (i FiltersResponseArgs) ToFiltersResponsePtrOutput() FiltersResponsePtrOutput {
	return i.ToFiltersResponsePtrOutputWithContext(context.Background())
}

func (i FiltersResponseArgs) ToFiltersResponsePtrOutputWithContext(ctx context.Context) FiltersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FiltersResponseOutput).ToFiltersResponsePtrOutputWithContext(ctx)
}









type FiltersResponsePtrInput interface {
	pulumi.Input

	ToFiltersResponsePtrOutput() FiltersResponsePtrOutput
	ToFiltersResponsePtrOutputWithContext(context.Context) FiltersResponsePtrOutput
}

type filtersResponsePtrType FiltersResponseArgs

func FiltersResponsePtr(v *FiltersResponseArgs) FiltersResponsePtrInput {
	return (*filtersResponsePtrType)(v)
}

func (*filtersResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FiltersResponse)(nil)).Elem()
}

func (i *filtersResponsePtrType) ToFiltersResponsePtrOutput() FiltersResponsePtrOutput {
	return i.ToFiltersResponsePtrOutputWithContext(context.Background())
}

func (i *filtersResponsePtrType) ToFiltersResponsePtrOutputWithContext(ctx context.Context) FiltersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FiltersResponsePtrOutput)
}

type FiltersResponseOutput struct{ *pulumi.OutputState }

func (FiltersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FiltersResponse)(nil)).Elem()
}

func (o FiltersResponseOutput) ToFiltersResponseOutput() FiltersResponseOutput {
	return o
}

func (o FiltersResponseOutput) ToFiltersResponseOutputWithContext(ctx context.Context) FiltersResponseOutput {
	return o
}

func (o FiltersResponseOutput) ToFiltersResponsePtrOutput() FiltersResponsePtrOutput {
	return o.ToFiltersResponsePtrOutputWithContext(context.Background())
}

func (o FiltersResponseOutput) ToFiltersResponsePtrOutputWithContext(ctx context.Context) FiltersResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FiltersResponse) *FiltersResponse {
		return &v
	}).(FiltersResponsePtrOutput)
}

func (o FiltersResponseOutput) Crop() RectangleResponsePtrOutput {
	return o.ApplyT(func(v FiltersResponse) *RectangleResponse { return v.Crop }).(RectangleResponsePtrOutput)
}

func (o FiltersResponseOutput) Deinterlace() DeinterlaceResponsePtrOutput {
	return o.ApplyT(func(v FiltersResponse) *DeinterlaceResponse { return v.Deinterlace }).(DeinterlaceResponsePtrOutput)
}

func (o FiltersResponseOutput) Overlays() pulumi.ArrayOutput {
	return o.ApplyT(func(v FiltersResponse) []interface{} { return v.Overlays }).(pulumi.ArrayOutput)
}

func (o FiltersResponseOutput) Rotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FiltersResponse) *string { return v.Rotation }).(pulumi.StringPtrOutput)
}

type FiltersResponsePtrOutput struct{ *pulumi.OutputState }

func (FiltersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FiltersResponse)(nil)).Elem()
}

func (o FiltersResponsePtrOutput) ToFiltersResponsePtrOutput() FiltersResponsePtrOutput {
	return o
}

func (o FiltersResponsePtrOutput) ToFiltersResponsePtrOutputWithContext(ctx context.Context) FiltersResponsePtrOutput {
	return o
}

func (o FiltersResponsePtrOutput) Elem() FiltersResponseOutput {
	return o.ApplyT(func(v *FiltersResponse) FiltersResponse {
		if v != nil {
			return *v
		}
		var ret FiltersResponse
		return ret
	}).(FiltersResponseOutput)
}

func (o FiltersResponsePtrOutput) Crop() RectangleResponsePtrOutput {
	return o.ApplyT(func(v *FiltersResponse) *RectangleResponse {
		if v == nil {
			return nil
		}
		return v.Crop
	}).(RectangleResponsePtrOutput)
}

func (o FiltersResponsePtrOutput) Deinterlace() DeinterlaceResponsePtrOutput {
	return o.ApplyT(func(v *FiltersResponse) *DeinterlaceResponse {
		if v == nil {
			return nil
		}
		return v.Deinterlace
	}).(DeinterlaceResponsePtrOutput)
}

func (o FiltersResponsePtrOutput) Overlays() pulumi.ArrayOutput {
	return o.ApplyT(func(v *FiltersResponse) []interface{} {
		if v == nil {
			return nil
		}
		return v.Overlays
	}).(pulumi.ArrayOutput)
}

func (o FiltersResponsePtrOutput) Rotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FiltersResponse) *string {
		if v == nil {
			return nil
		}
		return v.Rotation
	}).(pulumi.StringPtrOutput)
}

type FirstQuality struct {
	Bitrate int `pulumi:"bitrate"`
}





type FirstQualityInput interface {
	pulumi.Input

	ToFirstQualityOutput() FirstQualityOutput
	ToFirstQualityOutputWithContext(context.Context) FirstQualityOutput
}

type FirstQualityArgs struct {
	Bitrate pulumi.IntInput `pulumi:"bitrate"`
}

func (FirstQualityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirstQuality)(nil)).Elem()
}

func (i FirstQualityArgs) ToFirstQualityOutput() FirstQualityOutput {
	return i.ToFirstQualityOutputWithContext(context.Background())
}

func (i FirstQualityArgs) ToFirstQualityOutputWithContext(ctx context.Context) FirstQualityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirstQualityOutput)
}

func (i FirstQualityArgs) ToFirstQualityPtrOutput() FirstQualityPtrOutput {
	return i.ToFirstQualityPtrOutputWithContext(context.Background())
}

func (i FirstQualityArgs) ToFirstQualityPtrOutputWithContext(ctx context.Context) FirstQualityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirstQualityOutput).ToFirstQualityPtrOutputWithContext(ctx)
}









type FirstQualityPtrInput interface {
	pulumi.Input

	ToFirstQualityPtrOutput() FirstQualityPtrOutput
	ToFirstQualityPtrOutputWithContext(context.Context) FirstQualityPtrOutput
}

type firstQualityPtrType FirstQualityArgs

func FirstQualityPtr(v *FirstQualityArgs) FirstQualityPtrInput {
	return (*firstQualityPtrType)(v)
}

func (*firstQualityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirstQuality)(nil)).Elem()
}

func (i *firstQualityPtrType) ToFirstQualityPtrOutput() FirstQualityPtrOutput {
	return i.ToFirstQualityPtrOutputWithContext(context.Background())
}

func (i *firstQualityPtrType) ToFirstQualityPtrOutputWithContext(ctx context.Context) FirstQualityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirstQualityPtrOutput)
}

type FirstQualityOutput struct{ *pulumi.OutputState }

func (FirstQualityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirstQuality)(nil)).Elem()
}

func (o FirstQualityOutput) ToFirstQualityOutput() FirstQualityOutput {
	return o
}

func (o FirstQualityOutput) ToFirstQualityOutputWithContext(ctx context.Context) FirstQualityOutput {
	return o
}

func (o FirstQualityOutput) ToFirstQualityPtrOutput() FirstQualityPtrOutput {
	return o.ToFirstQualityPtrOutputWithContext(context.Background())
}

func (o FirstQualityOutput) ToFirstQualityPtrOutputWithContext(ctx context.Context) FirstQualityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirstQuality) *FirstQuality {
		return &v
	}).(FirstQualityPtrOutput)
}

func (o FirstQualityOutput) Bitrate() pulumi.IntOutput {
	return o.ApplyT(func(v FirstQuality) int { return v.Bitrate }).(pulumi.IntOutput)
}

type FirstQualityPtrOutput struct{ *pulumi.OutputState }

func (FirstQualityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirstQuality)(nil)).Elem()
}

func (o FirstQualityPtrOutput) ToFirstQualityPtrOutput() FirstQualityPtrOutput {
	return o
}

func (o FirstQualityPtrOutput) ToFirstQualityPtrOutputWithContext(ctx context.Context) FirstQualityPtrOutput {
	return o
}

func (o FirstQualityPtrOutput) Elem() FirstQualityOutput {
	return o.ApplyT(func(v *FirstQuality) FirstQuality {
		if v != nil {
			return *v
		}
		var ret FirstQuality
		return ret
	}).(FirstQualityOutput)
}

func (o FirstQualityPtrOutput) Bitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirstQuality) *int {
		if v == nil {
			return nil
		}
		return &v.Bitrate
	}).(pulumi.IntPtrOutput)
}

type FirstQualityResponse struct {
	Bitrate int `pulumi:"bitrate"`
}





type FirstQualityResponseInput interface {
	pulumi.Input

	ToFirstQualityResponseOutput() FirstQualityResponseOutput
	ToFirstQualityResponseOutputWithContext(context.Context) FirstQualityResponseOutput
}

type FirstQualityResponseArgs struct {
	Bitrate pulumi.IntInput `pulumi:"bitrate"`
}

func (FirstQualityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirstQualityResponse)(nil)).Elem()
}

func (i FirstQualityResponseArgs) ToFirstQualityResponseOutput() FirstQualityResponseOutput {
	return i.ToFirstQualityResponseOutputWithContext(context.Background())
}

func (i FirstQualityResponseArgs) ToFirstQualityResponseOutputWithContext(ctx context.Context) FirstQualityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirstQualityResponseOutput)
}

func (i FirstQualityResponseArgs) ToFirstQualityResponsePtrOutput() FirstQualityResponsePtrOutput {
	return i.ToFirstQualityResponsePtrOutputWithContext(context.Background())
}

func (i FirstQualityResponseArgs) ToFirstQualityResponsePtrOutputWithContext(ctx context.Context) FirstQualityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirstQualityResponseOutput).ToFirstQualityResponsePtrOutputWithContext(ctx)
}









type FirstQualityResponsePtrInput interface {
	pulumi.Input

	ToFirstQualityResponsePtrOutput() FirstQualityResponsePtrOutput
	ToFirstQualityResponsePtrOutputWithContext(context.Context) FirstQualityResponsePtrOutput
}

type firstQualityResponsePtrType FirstQualityResponseArgs

func FirstQualityResponsePtr(v *FirstQualityResponseArgs) FirstQualityResponsePtrInput {
	return (*firstQualityResponsePtrType)(v)
}

func (*firstQualityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirstQualityResponse)(nil)).Elem()
}

func (i *firstQualityResponsePtrType) ToFirstQualityResponsePtrOutput() FirstQualityResponsePtrOutput {
	return i.ToFirstQualityResponsePtrOutputWithContext(context.Background())
}

func (i *firstQualityResponsePtrType) ToFirstQualityResponsePtrOutputWithContext(ctx context.Context) FirstQualityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirstQualityResponsePtrOutput)
}

type FirstQualityResponseOutput struct{ *pulumi.OutputState }

func (FirstQualityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirstQualityResponse)(nil)).Elem()
}

func (o FirstQualityResponseOutput) ToFirstQualityResponseOutput() FirstQualityResponseOutput {
	return o
}

func (o FirstQualityResponseOutput) ToFirstQualityResponseOutputWithContext(ctx context.Context) FirstQualityResponseOutput {
	return o
}

func (o FirstQualityResponseOutput) ToFirstQualityResponsePtrOutput() FirstQualityResponsePtrOutput {
	return o.ToFirstQualityResponsePtrOutputWithContext(context.Background())
}

func (o FirstQualityResponseOutput) ToFirstQualityResponsePtrOutputWithContext(ctx context.Context) FirstQualityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirstQualityResponse) *FirstQualityResponse {
		return &v
	}).(FirstQualityResponsePtrOutput)
}

func (o FirstQualityResponseOutput) Bitrate() pulumi.IntOutput {
	return o.ApplyT(func(v FirstQualityResponse) int { return v.Bitrate }).(pulumi.IntOutput)
}

type FirstQualityResponsePtrOutput struct{ *pulumi.OutputState }

func (FirstQualityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirstQualityResponse)(nil)).Elem()
}

func (o FirstQualityResponsePtrOutput) ToFirstQualityResponsePtrOutput() FirstQualityResponsePtrOutput {
	return o
}

func (o FirstQualityResponsePtrOutput) ToFirstQualityResponsePtrOutputWithContext(ctx context.Context) FirstQualityResponsePtrOutput {
	return o
}

func (o FirstQualityResponsePtrOutput) Elem() FirstQualityResponseOutput {
	return o.ApplyT(func(v *FirstQualityResponse) FirstQualityResponse {
		if v != nil {
			return *v
		}
		var ret FirstQualityResponse
		return ret
	}).(FirstQualityResponseOutput)
}

func (o FirstQualityResponsePtrOutput) Bitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirstQualityResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Bitrate
	}).(pulumi.IntPtrOutput)
}

type FromAllInputFile struct {
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}





type FromAllInputFileInput interface {
	pulumi.Input

	ToFromAllInputFileOutput() FromAllInputFileOutput
	ToFromAllInputFileOutputWithContext(context.Context) FromAllInputFileOutput
}

type FromAllInputFileArgs struct {
	IncludedTracks pulumi.ArrayInput  `pulumi:"includedTracks"`
	OdataType      pulumi.StringInput `pulumi:"odataType"`
}

func (FromAllInputFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FromAllInputFile)(nil)).Elem()
}

func (i FromAllInputFileArgs) ToFromAllInputFileOutput() FromAllInputFileOutput {
	return i.ToFromAllInputFileOutputWithContext(context.Background())
}

func (i FromAllInputFileArgs) ToFromAllInputFileOutputWithContext(ctx context.Context) FromAllInputFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FromAllInputFileOutput)
}

type FromAllInputFileOutput struct{ *pulumi.OutputState }

func (FromAllInputFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FromAllInputFile)(nil)).Elem()
}

func (o FromAllInputFileOutput) ToFromAllInputFileOutput() FromAllInputFileOutput {
	return o
}

func (o FromAllInputFileOutput) ToFromAllInputFileOutputWithContext(ctx context.Context) FromAllInputFileOutput {
	return o
}

func (o FromAllInputFileOutput) IncludedTracks() pulumi.ArrayOutput {
	return o.ApplyT(func(v FromAllInputFile) []interface{} { return v.IncludedTracks }).(pulumi.ArrayOutput)
}

func (o FromAllInputFileOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v FromAllInputFile) string { return v.OdataType }).(pulumi.StringOutput)
}

type FromAllInputFileResponse struct {
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}





type FromAllInputFileResponseInput interface {
	pulumi.Input

	ToFromAllInputFileResponseOutput() FromAllInputFileResponseOutput
	ToFromAllInputFileResponseOutputWithContext(context.Context) FromAllInputFileResponseOutput
}

type FromAllInputFileResponseArgs struct {
	IncludedTracks pulumi.ArrayInput  `pulumi:"includedTracks"`
	OdataType      pulumi.StringInput `pulumi:"odataType"`
}

func (FromAllInputFileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FromAllInputFileResponse)(nil)).Elem()
}

func (i FromAllInputFileResponseArgs) ToFromAllInputFileResponseOutput() FromAllInputFileResponseOutput {
	return i.ToFromAllInputFileResponseOutputWithContext(context.Background())
}

func (i FromAllInputFileResponseArgs) ToFromAllInputFileResponseOutputWithContext(ctx context.Context) FromAllInputFileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FromAllInputFileResponseOutput)
}

type FromAllInputFileResponseOutput struct{ *pulumi.OutputState }

func (FromAllInputFileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FromAllInputFileResponse)(nil)).Elem()
}

func (o FromAllInputFileResponseOutput) ToFromAllInputFileResponseOutput() FromAllInputFileResponseOutput {
	return o
}

func (o FromAllInputFileResponseOutput) ToFromAllInputFileResponseOutputWithContext(ctx context.Context) FromAllInputFileResponseOutput {
	return o
}

func (o FromAllInputFileResponseOutput) IncludedTracks() pulumi.ArrayOutput {
	return o.ApplyT(func(v FromAllInputFileResponse) []interface{} { return v.IncludedTracks }).(pulumi.ArrayOutput)
}

func (o FromAllInputFileResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v FromAllInputFileResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type FromEachInputFile struct {
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}





type FromEachInputFileInput interface {
	pulumi.Input

	ToFromEachInputFileOutput() FromEachInputFileOutput
	ToFromEachInputFileOutputWithContext(context.Context) FromEachInputFileOutput
}

type FromEachInputFileArgs struct {
	IncludedTracks pulumi.ArrayInput  `pulumi:"includedTracks"`
	OdataType      pulumi.StringInput `pulumi:"odataType"`
}

func (FromEachInputFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FromEachInputFile)(nil)).Elem()
}

func (i FromEachInputFileArgs) ToFromEachInputFileOutput() FromEachInputFileOutput {
	return i.ToFromEachInputFileOutputWithContext(context.Background())
}

func (i FromEachInputFileArgs) ToFromEachInputFileOutputWithContext(ctx context.Context) FromEachInputFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FromEachInputFileOutput)
}

type FromEachInputFileOutput struct{ *pulumi.OutputState }

func (FromEachInputFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FromEachInputFile)(nil)).Elem()
}

func (o FromEachInputFileOutput) ToFromEachInputFileOutput() FromEachInputFileOutput {
	return o
}

func (o FromEachInputFileOutput) ToFromEachInputFileOutputWithContext(ctx context.Context) FromEachInputFileOutput {
	return o
}

func (o FromEachInputFileOutput) IncludedTracks() pulumi.ArrayOutput {
	return o.ApplyT(func(v FromEachInputFile) []interface{} { return v.IncludedTracks }).(pulumi.ArrayOutput)
}

func (o FromEachInputFileOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v FromEachInputFile) string { return v.OdataType }).(pulumi.StringOutput)
}

type FromEachInputFileResponse struct {
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}





type FromEachInputFileResponseInput interface {
	pulumi.Input

	ToFromEachInputFileResponseOutput() FromEachInputFileResponseOutput
	ToFromEachInputFileResponseOutputWithContext(context.Context) FromEachInputFileResponseOutput
}

type FromEachInputFileResponseArgs struct {
	IncludedTracks pulumi.ArrayInput  `pulumi:"includedTracks"`
	OdataType      pulumi.StringInput `pulumi:"odataType"`
}

func (FromEachInputFileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FromEachInputFileResponse)(nil)).Elem()
}

func (i FromEachInputFileResponseArgs) ToFromEachInputFileResponseOutput() FromEachInputFileResponseOutput {
	return i.ToFromEachInputFileResponseOutputWithContext(context.Background())
}

func (i FromEachInputFileResponseArgs) ToFromEachInputFileResponseOutputWithContext(ctx context.Context) FromEachInputFileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FromEachInputFileResponseOutput)
}

type FromEachInputFileResponseOutput struct{ *pulumi.OutputState }

func (FromEachInputFileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FromEachInputFileResponse)(nil)).Elem()
}

func (o FromEachInputFileResponseOutput) ToFromEachInputFileResponseOutput() FromEachInputFileResponseOutput {
	return o
}

func (o FromEachInputFileResponseOutput) ToFromEachInputFileResponseOutputWithContext(ctx context.Context) FromEachInputFileResponseOutput {
	return o
}

func (o FromEachInputFileResponseOutput) IncludedTracks() pulumi.ArrayOutput {
	return o.ApplyT(func(v FromEachInputFileResponse) []interface{} { return v.IncludedTracks }).(pulumi.ArrayOutput)
}

func (o FromEachInputFileResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v FromEachInputFileResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type H264Layer struct {
	AdaptiveBFrame  *bool   `pulumi:"adaptiveBFrame"`
	BFrames         *int    `pulumi:"bFrames"`
	Bitrate         int     `pulumi:"bitrate"`
	BufferWindow    *string `pulumi:"bufferWindow"`
	EntropyMode     *string `pulumi:"entropyMode"`
	FrameRate       *string `pulumi:"frameRate"`
	Height          *string `pulumi:"height"`
	Label           *string `pulumi:"label"`
	Level           *string `pulumi:"level"`
	MaxBitrate      *int    `pulumi:"maxBitrate"`
	OdataType       string  `pulumi:"odataType"`
	Profile         *string `pulumi:"profile"`
	ReferenceFrames *int    `pulumi:"referenceFrames"`
	Slices          *int    `pulumi:"slices"`
	Width           *string `pulumi:"width"`
}





type H264LayerInput interface {
	pulumi.Input

	ToH264LayerOutput() H264LayerOutput
	ToH264LayerOutputWithContext(context.Context) H264LayerOutput
}

type H264LayerArgs struct {
	AdaptiveBFrame  pulumi.BoolPtrInput   `pulumi:"adaptiveBFrame"`
	BFrames         pulumi.IntPtrInput    `pulumi:"bFrames"`
	Bitrate         pulumi.IntInput       `pulumi:"bitrate"`
	BufferWindow    pulumi.StringPtrInput `pulumi:"bufferWindow"`
	EntropyMode     pulumi.StringPtrInput `pulumi:"entropyMode"`
	FrameRate       pulumi.StringPtrInput `pulumi:"frameRate"`
	Height          pulumi.StringPtrInput `pulumi:"height"`
	Label           pulumi.StringPtrInput `pulumi:"label"`
	Level           pulumi.StringPtrInput `pulumi:"level"`
	MaxBitrate      pulumi.IntPtrInput    `pulumi:"maxBitrate"`
	OdataType       pulumi.StringInput    `pulumi:"odataType"`
	Profile         pulumi.StringPtrInput `pulumi:"profile"`
	ReferenceFrames pulumi.IntPtrInput    `pulumi:"referenceFrames"`
	Slices          pulumi.IntPtrInput    `pulumi:"slices"`
	Width           pulumi.StringPtrInput `pulumi:"width"`
}

func (H264LayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*H264Layer)(nil)).Elem()
}

func (i H264LayerArgs) ToH264LayerOutput() H264LayerOutput {
	return i.ToH264LayerOutputWithContext(context.Background())
}

func (i H264LayerArgs) ToH264LayerOutputWithContext(ctx context.Context) H264LayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H264LayerOutput)
}





type H264LayerArrayInput interface {
	pulumi.Input

	ToH264LayerArrayOutput() H264LayerArrayOutput
	ToH264LayerArrayOutputWithContext(context.Context) H264LayerArrayOutput
}

type H264LayerArray []H264LayerInput

func (H264LayerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]H264Layer)(nil)).Elem()
}

func (i H264LayerArray) ToH264LayerArrayOutput() H264LayerArrayOutput {
	return i.ToH264LayerArrayOutputWithContext(context.Background())
}

func (i H264LayerArray) ToH264LayerArrayOutputWithContext(ctx context.Context) H264LayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H264LayerArrayOutput)
}

type H264LayerOutput struct{ *pulumi.OutputState }

func (H264LayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*H264Layer)(nil)).Elem()
}

func (o H264LayerOutput) ToH264LayerOutput() H264LayerOutput {
	return o
}

func (o H264LayerOutput) ToH264LayerOutputWithContext(ctx context.Context) H264LayerOutput {
	return o
}

func (o H264LayerOutput) AdaptiveBFrame() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v H264Layer) *bool { return v.AdaptiveBFrame }).(pulumi.BoolPtrOutput)
}

func (o H264LayerOutput) BFrames() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H264Layer) *int { return v.BFrames }).(pulumi.IntPtrOutput)
}

func (o H264LayerOutput) Bitrate() pulumi.IntOutput {
	return o.ApplyT(func(v H264Layer) int { return v.Bitrate }).(pulumi.IntOutput)
}

func (o H264LayerOutput) BufferWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Layer) *string { return v.BufferWindow }).(pulumi.StringPtrOutput)
}

func (o H264LayerOutput) EntropyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Layer) *string { return v.EntropyMode }).(pulumi.StringPtrOutput)
}

func (o H264LayerOutput) FrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Layer) *string { return v.FrameRate }).(pulumi.StringPtrOutput)
}

func (o H264LayerOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Layer) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o H264LayerOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Layer) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o H264LayerOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Layer) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o H264LayerOutput) MaxBitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H264Layer) *int { return v.MaxBitrate }).(pulumi.IntPtrOutput)
}

func (o H264LayerOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v H264Layer) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o H264LayerOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Layer) *string { return v.Profile }).(pulumi.StringPtrOutput)
}

func (o H264LayerOutput) ReferenceFrames() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H264Layer) *int { return v.ReferenceFrames }).(pulumi.IntPtrOutput)
}

func (o H264LayerOutput) Slices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H264Layer) *int { return v.Slices }).(pulumi.IntPtrOutput)
}

func (o H264LayerOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Layer) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type H264LayerArrayOutput struct{ *pulumi.OutputState }

func (H264LayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]H264Layer)(nil)).Elem()
}

func (o H264LayerArrayOutput) ToH264LayerArrayOutput() H264LayerArrayOutput {
	return o
}

func (o H264LayerArrayOutput) ToH264LayerArrayOutputWithContext(ctx context.Context) H264LayerArrayOutput {
	return o
}

func (o H264LayerArrayOutput) Index(i pulumi.IntInput) H264LayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) H264Layer {
		return vs[0].([]H264Layer)[vs[1].(int)]
	}).(H264LayerOutput)
}

type H264LayerResponse struct {
	AdaptiveBFrame  *bool   `pulumi:"adaptiveBFrame"`
	BFrames         *int    `pulumi:"bFrames"`
	Bitrate         int     `pulumi:"bitrate"`
	BufferWindow    *string `pulumi:"bufferWindow"`
	EntropyMode     *string `pulumi:"entropyMode"`
	FrameRate       *string `pulumi:"frameRate"`
	Height          *string `pulumi:"height"`
	Label           *string `pulumi:"label"`
	Level           *string `pulumi:"level"`
	MaxBitrate      *int    `pulumi:"maxBitrate"`
	OdataType       string  `pulumi:"odataType"`
	Profile         *string `pulumi:"profile"`
	ReferenceFrames *int    `pulumi:"referenceFrames"`
	Slices          *int    `pulumi:"slices"`
	Width           *string `pulumi:"width"`
}





type H264LayerResponseInput interface {
	pulumi.Input

	ToH264LayerResponseOutput() H264LayerResponseOutput
	ToH264LayerResponseOutputWithContext(context.Context) H264LayerResponseOutput
}

type H264LayerResponseArgs struct {
	AdaptiveBFrame  pulumi.BoolPtrInput   `pulumi:"adaptiveBFrame"`
	BFrames         pulumi.IntPtrInput    `pulumi:"bFrames"`
	Bitrate         pulumi.IntInput       `pulumi:"bitrate"`
	BufferWindow    pulumi.StringPtrInput `pulumi:"bufferWindow"`
	EntropyMode     pulumi.StringPtrInput `pulumi:"entropyMode"`
	FrameRate       pulumi.StringPtrInput `pulumi:"frameRate"`
	Height          pulumi.StringPtrInput `pulumi:"height"`
	Label           pulumi.StringPtrInput `pulumi:"label"`
	Level           pulumi.StringPtrInput `pulumi:"level"`
	MaxBitrate      pulumi.IntPtrInput    `pulumi:"maxBitrate"`
	OdataType       pulumi.StringInput    `pulumi:"odataType"`
	Profile         pulumi.StringPtrInput `pulumi:"profile"`
	ReferenceFrames pulumi.IntPtrInput    `pulumi:"referenceFrames"`
	Slices          pulumi.IntPtrInput    `pulumi:"slices"`
	Width           pulumi.StringPtrInput `pulumi:"width"`
}

func (H264LayerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*H264LayerResponse)(nil)).Elem()
}

func (i H264LayerResponseArgs) ToH264LayerResponseOutput() H264LayerResponseOutput {
	return i.ToH264LayerResponseOutputWithContext(context.Background())
}

func (i H264LayerResponseArgs) ToH264LayerResponseOutputWithContext(ctx context.Context) H264LayerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H264LayerResponseOutput)
}





type H264LayerResponseArrayInput interface {
	pulumi.Input

	ToH264LayerResponseArrayOutput() H264LayerResponseArrayOutput
	ToH264LayerResponseArrayOutputWithContext(context.Context) H264LayerResponseArrayOutput
}

type H264LayerResponseArray []H264LayerResponseInput

func (H264LayerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]H264LayerResponse)(nil)).Elem()
}

func (i H264LayerResponseArray) ToH264LayerResponseArrayOutput() H264LayerResponseArrayOutput {
	return i.ToH264LayerResponseArrayOutputWithContext(context.Background())
}

func (i H264LayerResponseArray) ToH264LayerResponseArrayOutputWithContext(ctx context.Context) H264LayerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H264LayerResponseArrayOutput)
}

type H264LayerResponseOutput struct{ *pulumi.OutputState }

func (H264LayerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*H264LayerResponse)(nil)).Elem()
}

func (o H264LayerResponseOutput) ToH264LayerResponseOutput() H264LayerResponseOutput {
	return o
}

func (o H264LayerResponseOutput) ToH264LayerResponseOutputWithContext(ctx context.Context) H264LayerResponseOutput {
	return o
}

func (o H264LayerResponseOutput) AdaptiveBFrame() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *bool { return v.AdaptiveBFrame }).(pulumi.BoolPtrOutput)
}

func (o H264LayerResponseOutput) BFrames() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *int { return v.BFrames }).(pulumi.IntPtrOutput)
}

func (o H264LayerResponseOutput) Bitrate() pulumi.IntOutput {
	return o.ApplyT(func(v H264LayerResponse) int { return v.Bitrate }).(pulumi.IntOutput)
}

func (o H264LayerResponseOutput) BufferWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *string { return v.BufferWindow }).(pulumi.StringPtrOutput)
}

func (o H264LayerResponseOutput) EntropyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *string { return v.EntropyMode }).(pulumi.StringPtrOutput)
}

func (o H264LayerResponseOutput) FrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *string { return v.FrameRate }).(pulumi.StringPtrOutput)
}

func (o H264LayerResponseOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o H264LayerResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o H264LayerResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o H264LayerResponseOutput) MaxBitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *int { return v.MaxBitrate }).(pulumi.IntPtrOutput)
}

func (o H264LayerResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v H264LayerResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o H264LayerResponseOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *string { return v.Profile }).(pulumi.StringPtrOutput)
}

func (o H264LayerResponseOutput) ReferenceFrames() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *int { return v.ReferenceFrames }).(pulumi.IntPtrOutput)
}

func (o H264LayerResponseOutput) Slices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *int { return v.Slices }).(pulumi.IntPtrOutput)
}

func (o H264LayerResponseOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264LayerResponse) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type H264LayerResponseArrayOutput struct{ *pulumi.OutputState }

func (H264LayerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]H264LayerResponse)(nil)).Elem()
}

func (o H264LayerResponseArrayOutput) ToH264LayerResponseArrayOutput() H264LayerResponseArrayOutput {
	return o
}

func (o H264LayerResponseArrayOutput) ToH264LayerResponseArrayOutputWithContext(ctx context.Context) H264LayerResponseArrayOutput {
	return o
}

func (o H264LayerResponseArrayOutput) Index(i pulumi.IntInput) H264LayerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) H264LayerResponse {
		return vs[0].([]H264LayerResponse)[vs[1].(int)]
	}).(H264LayerResponseOutput)
}

type H264Video struct {
	Complexity           *string     `pulumi:"complexity"`
	KeyFrameInterval     *string     `pulumi:"keyFrameInterval"`
	Label                *string     `pulumi:"label"`
	Layers               []H264Layer `pulumi:"layers"`
	OdataType            string      `pulumi:"odataType"`
	SceneChangeDetection *bool       `pulumi:"sceneChangeDetection"`
	StretchMode          *string     `pulumi:"stretchMode"`
	SyncMode             *string     `pulumi:"syncMode"`
}





type H264VideoInput interface {
	pulumi.Input

	ToH264VideoOutput() H264VideoOutput
	ToH264VideoOutputWithContext(context.Context) H264VideoOutput
}

type H264VideoArgs struct {
	Complexity           pulumi.StringPtrInput `pulumi:"complexity"`
	KeyFrameInterval     pulumi.StringPtrInput `pulumi:"keyFrameInterval"`
	Label                pulumi.StringPtrInput `pulumi:"label"`
	Layers               H264LayerArrayInput   `pulumi:"layers"`
	OdataType            pulumi.StringInput    `pulumi:"odataType"`
	SceneChangeDetection pulumi.BoolPtrInput   `pulumi:"sceneChangeDetection"`
	StretchMode          pulumi.StringPtrInput `pulumi:"stretchMode"`
	SyncMode             pulumi.StringPtrInput `pulumi:"syncMode"`
}

func (H264VideoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*H264Video)(nil)).Elem()
}

func (i H264VideoArgs) ToH264VideoOutput() H264VideoOutput {
	return i.ToH264VideoOutputWithContext(context.Background())
}

func (i H264VideoArgs) ToH264VideoOutputWithContext(ctx context.Context) H264VideoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H264VideoOutput)
}

type H264VideoOutput struct{ *pulumi.OutputState }

func (H264VideoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*H264Video)(nil)).Elem()
}

func (o H264VideoOutput) ToH264VideoOutput() H264VideoOutput {
	return o
}

func (o H264VideoOutput) ToH264VideoOutputWithContext(ctx context.Context) H264VideoOutput {
	return o
}

func (o H264VideoOutput) Complexity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Video) *string { return v.Complexity }).(pulumi.StringPtrOutput)
}

func (o H264VideoOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Video) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o H264VideoOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Video) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o H264VideoOutput) Layers() H264LayerArrayOutput {
	return o.ApplyT(func(v H264Video) []H264Layer { return v.Layers }).(H264LayerArrayOutput)
}

func (o H264VideoOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v H264Video) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o H264VideoOutput) SceneChangeDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v H264Video) *bool { return v.SceneChangeDetection }).(pulumi.BoolPtrOutput)
}

func (o H264VideoOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Video) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o H264VideoOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264Video) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type H264VideoResponse struct {
	Complexity           *string             `pulumi:"complexity"`
	KeyFrameInterval     *string             `pulumi:"keyFrameInterval"`
	Label                *string             `pulumi:"label"`
	Layers               []H264LayerResponse `pulumi:"layers"`
	OdataType            string              `pulumi:"odataType"`
	SceneChangeDetection *bool               `pulumi:"sceneChangeDetection"`
	StretchMode          *string             `pulumi:"stretchMode"`
	SyncMode             *string             `pulumi:"syncMode"`
}





type H264VideoResponseInput interface {
	pulumi.Input

	ToH264VideoResponseOutput() H264VideoResponseOutput
	ToH264VideoResponseOutputWithContext(context.Context) H264VideoResponseOutput
}

type H264VideoResponseArgs struct {
	Complexity           pulumi.StringPtrInput       `pulumi:"complexity"`
	KeyFrameInterval     pulumi.StringPtrInput       `pulumi:"keyFrameInterval"`
	Label                pulumi.StringPtrInput       `pulumi:"label"`
	Layers               H264LayerResponseArrayInput `pulumi:"layers"`
	OdataType            pulumi.StringInput          `pulumi:"odataType"`
	SceneChangeDetection pulumi.BoolPtrInput         `pulumi:"sceneChangeDetection"`
	StretchMode          pulumi.StringPtrInput       `pulumi:"stretchMode"`
	SyncMode             pulumi.StringPtrInput       `pulumi:"syncMode"`
}

func (H264VideoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*H264VideoResponse)(nil)).Elem()
}

func (i H264VideoResponseArgs) ToH264VideoResponseOutput() H264VideoResponseOutput {
	return i.ToH264VideoResponseOutputWithContext(context.Background())
}

func (i H264VideoResponseArgs) ToH264VideoResponseOutputWithContext(ctx context.Context) H264VideoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H264VideoResponseOutput)
}

type H264VideoResponseOutput struct{ *pulumi.OutputState }

func (H264VideoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*H264VideoResponse)(nil)).Elem()
}

func (o H264VideoResponseOutput) ToH264VideoResponseOutput() H264VideoResponseOutput {
	return o
}

func (o H264VideoResponseOutput) ToH264VideoResponseOutputWithContext(ctx context.Context) H264VideoResponseOutput {
	return o
}

func (o H264VideoResponseOutput) Complexity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264VideoResponse) *string { return v.Complexity }).(pulumi.StringPtrOutput)
}

func (o H264VideoResponseOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264VideoResponse) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o H264VideoResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264VideoResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o H264VideoResponseOutput) Layers() H264LayerResponseArrayOutput {
	return o.ApplyT(func(v H264VideoResponse) []H264LayerResponse { return v.Layers }).(H264LayerResponseArrayOutput)
}

func (o H264VideoResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v H264VideoResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o H264VideoResponseOutput) SceneChangeDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v H264VideoResponse) *bool { return v.SceneChangeDetection }).(pulumi.BoolPtrOutput)
}

func (o H264VideoResponseOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264VideoResponse) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o H264VideoResponseOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H264VideoResponse) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type H265Layer struct {
	AdaptiveBFrame  *bool   `pulumi:"adaptiveBFrame"`
	BFrames         *int    `pulumi:"bFrames"`
	Bitrate         int     `pulumi:"bitrate"`
	BufferWindow    *string `pulumi:"bufferWindow"`
	FrameRate       *string `pulumi:"frameRate"`
	Height          *string `pulumi:"height"`
	Label           *string `pulumi:"label"`
	Level           *string `pulumi:"level"`
	MaxBitrate      *int    `pulumi:"maxBitrate"`
	OdataType       string  `pulumi:"odataType"`
	Profile         *string `pulumi:"profile"`
	ReferenceFrames *int    `pulumi:"referenceFrames"`
	Slices          *int    `pulumi:"slices"`
	Width           *string `pulumi:"width"`
}





type H265LayerInput interface {
	pulumi.Input

	ToH265LayerOutput() H265LayerOutput
	ToH265LayerOutputWithContext(context.Context) H265LayerOutput
}

type H265LayerArgs struct {
	AdaptiveBFrame  pulumi.BoolPtrInput   `pulumi:"adaptiveBFrame"`
	BFrames         pulumi.IntPtrInput    `pulumi:"bFrames"`
	Bitrate         pulumi.IntInput       `pulumi:"bitrate"`
	BufferWindow    pulumi.StringPtrInput `pulumi:"bufferWindow"`
	FrameRate       pulumi.StringPtrInput `pulumi:"frameRate"`
	Height          pulumi.StringPtrInput `pulumi:"height"`
	Label           pulumi.StringPtrInput `pulumi:"label"`
	Level           pulumi.StringPtrInput `pulumi:"level"`
	MaxBitrate      pulumi.IntPtrInput    `pulumi:"maxBitrate"`
	OdataType       pulumi.StringInput    `pulumi:"odataType"`
	Profile         pulumi.StringPtrInput `pulumi:"profile"`
	ReferenceFrames pulumi.IntPtrInput    `pulumi:"referenceFrames"`
	Slices          pulumi.IntPtrInput    `pulumi:"slices"`
	Width           pulumi.StringPtrInput `pulumi:"width"`
}

func (H265LayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*H265Layer)(nil)).Elem()
}

func (i H265LayerArgs) ToH265LayerOutput() H265LayerOutput {
	return i.ToH265LayerOutputWithContext(context.Background())
}

func (i H265LayerArgs) ToH265LayerOutputWithContext(ctx context.Context) H265LayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H265LayerOutput)
}





type H265LayerArrayInput interface {
	pulumi.Input

	ToH265LayerArrayOutput() H265LayerArrayOutput
	ToH265LayerArrayOutputWithContext(context.Context) H265LayerArrayOutput
}

type H265LayerArray []H265LayerInput

func (H265LayerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]H265Layer)(nil)).Elem()
}

func (i H265LayerArray) ToH265LayerArrayOutput() H265LayerArrayOutput {
	return i.ToH265LayerArrayOutputWithContext(context.Background())
}

func (i H265LayerArray) ToH265LayerArrayOutputWithContext(ctx context.Context) H265LayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H265LayerArrayOutput)
}

type H265LayerOutput struct{ *pulumi.OutputState }

func (H265LayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*H265Layer)(nil)).Elem()
}

func (o H265LayerOutput) ToH265LayerOutput() H265LayerOutput {
	return o
}

func (o H265LayerOutput) ToH265LayerOutputWithContext(ctx context.Context) H265LayerOutput {
	return o
}

func (o H265LayerOutput) AdaptiveBFrame() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v H265Layer) *bool { return v.AdaptiveBFrame }).(pulumi.BoolPtrOutput)
}

func (o H265LayerOutput) BFrames() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H265Layer) *int { return v.BFrames }).(pulumi.IntPtrOutput)
}

func (o H265LayerOutput) Bitrate() pulumi.IntOutput {
	return o.ApplyT(func(v H265Layer) int { return v.Bitrate }).(pulumi.IntOutput)
}

func (o H265LayerOutput) BufferWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Layer) *string { return v.BufferWindow }).(pulumi.StringPtrOutput)
}

func (o H265LayerOutput) FrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Layer) *string { return v.FrameRate }).(pulumi.StringPtrOutput)
}

func (o H265LayerOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Layer) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o H265LayerOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Layer) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o H265LayerOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Layer) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o H265LayerOutput) MaxBitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H265Layer) *int { return v.MaxBitrate }).(pulumi.IntPtrOutput)
}

func (o H265LayerOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v H265Layer) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o H265LayerOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Layer) *string { return v.Profile }).(pulumi.StringPtrOutput)
}

func (o H265LayerOutput) ReferenceFrames() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H265Layer) *int { return v.ReferenceFrames }).(pulumi.IntPtrOutput)
}

func (o H265LayerOutput) Slices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H265Layer) *int { return v.Slices }).(pulumi.IntPtrOutput)
}

func (o H265LayerOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Layer) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type H265LayerArrayOutput struct{ *pulumi.OutputState }

func (H265LayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]H265Layer)(nil)).Elem()
}

func (o H265LayerArrayOutput) ToH265LayerArrayOutput() H265LayerArrayOutput {
	return o
}

func (o H265LayerArrayOutput) ToH265LayerArrayOutputWithContext(ctx context.Context) H265LayerArrayOutput {
	return o
}

func (o H265LayerArrayOutput) Index(i pulumi.IntInput) H265LayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) H265Layer {
		return vs[0].([]H265Layer)[vs[1].(int)]
	}).(H265LayerOutput)
}

type H265LayerResponse struct {
	AdaptiveBFrame  *bool   `pulumi:"adaptiveBFrame"`
	BFrames         *int    `pulumi:"bFrames"`
	Bitrate         int     `pulumi:"bitrate"`
	BufferWindow    *string `pulumi:"bufferWindow"`
	FrameRate       *string `pulumi:"frameRate"`
	Height          *string `pulumi:"height"`
	Label           *string `pulumi:"label"`
	Level           *string `pulumi:"level"`
	MaxBitrate      *int    `pulumi:"maxBitrate"`
	OdataType       string  `pulumi:"odataType"`
	Profile         *string `pulumi:"profile"`
	ReferenceFrames *int    `pulumi:"referenceFrames"`
	Slices          *int    `pulumi:"slices"`
	Width           *string `pulumi:"width"`
}





type H265LayerResponseInput interface {
	pulumi.Input

	ToH265LayerResponseOutput() H265LayerResponseOutput
	ToH265LayerResponseOutputWithContext(context.Context) H265LayerResponseOutput
}

type H265LayerResponseArgs struct {
	AdaptiveBFrame  pulumi.BoolPtrInput   `pulumi:"adaptiveBFrame"`
	BFrames         pulumi.IntPtrInput    `pulumi:"bFrames"`
	Bitrate         pulumi.IntInput       `pulumi:"bitrate"`
	BufferWindow    pulumi.StringPtrInput `pulumi:"bufferWindow"`
	FrameRate       pulumi.StringPtrInput `pulumi:"frameRate"`
	Height          pulumi.StringPtrInput `pulumi:"height"`
	Label           pulumi.StringPtrInput `pulumi:"label"`
	Level           pulumi.StringPtrInput `pulumi:"level"`
	MaxBitrate      pulumi.IntPtrInput    `pulumi:"maxBitrate"`
	OdataType       pulumi.StringInput    `pulumi:"odataType"`
	Profile         pulumi.StringPtrInput `pulumi:"profile"`
	ReferenceFrames pulumi.IntPtrInput    `pulumi:"referenceFrames"`
	Slices          pulumi.IntPtrInput    `pulumi:"slices"`
	Width           pulumi.StringPtrInput `pulumi:"width"`
}

func (H265LayerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*H265LayerResponse)(nil)).Elem()
}

func (i H265LayerResponseArgs) ToH265LayerResponseOutput() H265LayerResponseOutput {
	return i.ToH265LayerResponseOutputWithContext(context.Background())
}

func (i H265LayerResponseArgs) ToH265LayerResponseOutputWithContext(ctx context.Context) H265LayerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H265LayerResponseOutput)
}





type H265LayerResponseArrayInput interface {
	pulumi.Input

	ToH265LayerResponseArrayOutput() H265LayerResponseArrayOutput
	ToH265LayerResponseArrayOutputWithContext(context.Context) H265LayerResponseArrayOutput
}

type H265LayerResponseArray []H265LayerResponseInput

func (H265LayerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]H265LayerResponse)(nil)).Elem()
}

func (i H265LayerResponseArray) ToH265LayerResponseArrayOutput() H265LayerResponseArrayOutput {
	return i.ToH265LayerResponseArrayOutputWithContext(context.Background())
}

func (i H265LayerResponseArray) ToH265LayerResponseArrayOutputWithContext(ctx context.Context) H265LayerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H265LayerResponseArrayOutput)
}

type H265LayerResponseOutput struct{ *pulumi.OutputState }

func (H265LayerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*H265LayerResponse)(nil)).Elem()
}

func (o H265LayerResponseOutput) ToH265LayerResponseOutput() H265LayerResponseOutput {
	return o
}

func (o H265LayerResponseOutput) ToH265LayerResponseOutputWithContext(ctx context.Context) H265LayerResponseOutput {
	return o
}

func (o H265LayerResponseOutput) AdaptiveBFrame() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *bool { return v.AdaptiveBFrame }).(pulumi.BoolPtrOutput)
}

func (o H265LayerResponseOutput) BFrames() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *int { return v.BFrames }).(pulumi.IntPtrOutput)
}

func (o H265LayerResponseOutput) Bitrate() pulumi.IntOutput {
	return o.ApplyT(func(v H265LayerResponse) int { return v.Bitrate }).(pulumi.IntOutput)
}

func (o H265LayerResponseOutput) BufferWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *string { return v.BufferWindow }).(pulumi.StringPtrOutput)
}

func (o H265LayerResponseOutput) FrameRate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *string { return v.FrameRate }).(pulumi.StringPtrOutput)
}

func (o H265LayerResponseOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o H265LayerResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o H265LayerResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o H265LayerResponseOutput) MaxBitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *int { return v.MaxBitrate }).(pulumi.IntPtrOutput)
}

func (o H265LayerResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v H265LayerResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o H265LayerResponseOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *string { return v.Profile }).(pulumi.StringPtrOutput)
}

func (o H265LayerResponseOutput) ReferenceFrames() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *int { return v.ReferenceFrames }).(pulumi.IntPtrOutput)
}

func (o H265LayerResponseOutput) Slices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *int { return v.Slices }).(pulumi.IntPtrOutput)
}

func (o H265LayerResponseOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265LayerResponse) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type H265LayerResponseArrayOutput struct{ *pulumi.OutputState }

func (H265LayerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]H265LayerResponse)(nil)).Elem()
}

func (o H265LayerResponseArrayOutput) ToH265LayerResponseArrayOutput() H265LayerResponseArrayOutput {
	return o
}

func (o H265LayerResponseArrayOutput) ToH265LayerResponseArrayOutputWithContext(ctx context.Context) H265LayerResponseArrayOutput {
	return o
}

func (o H265LayerResponseArrayOutput) Index(i pulumi.IntInput) H265LayerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) H265LayerResponse {
		return vs[0].([]H265LayerResponse)[vs[1].(int)]
	}).(H265LayerResponseOutput)
}

type H265Video struct {
	Complexity           *string     `pulumi:"complexity"`
	KeyFrameInterval     *string     `pulumi:"keyFrameInterval"`
	Label                *string     `pulumi:"label"`
	Layers               []H265Layer `pulumi:"layers"`
	OdataType            string      `pulumi:"odataType"`
	SceneChangeDetection *bool       `pulumi:"sceneChangeDetection"`
	StretchMode          *string     `pulumi:"stretchMode"`
	SyncMode             *string     `pulumi:"syncMode"`
}





type H265VideoInput interface {
	pulumi.Input

	ToH265VideoOutput() H265VideoOutput
	ToH265VideoOutputWithContext(context.Context) H265VideoOutput
}

type H265VideoArgs struct {
	Complexity           pulumi.StringPtrInput `pulumi:"complexity"`
	KeyFrameInterval     pulumi.StringPtrInput `pulumi:"keyFrameInterval"`
	Label                pulumi.StringPtrInput `pulumi:"label"`
	Layers               H265LayerArrayInput   `pulumi:"layers"`
	OdataType            pulumi.StringInput    `pulumi:"odataType"`
	SceneChangeDetection pulumi.BoolPtrInput   `pulumi:"sceneChangeDetection"`
	StretchMode          pulumi.StringPtrInput `pulumi:"stretchMode"`
	SyncMode             pulumi.StringPtrInput `pulumi:"syncMode"`
}

func (H265VideoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*H265Video)(nil)).Elem()
}

func (i H265VideoArgs) ToH265VideoOutput() H265VideoOutput {
	return i.ToH265VideoOutputWithContext(context.Background())
}

func (i H265VideoArgs) ToH265VideoOutputWithContext(ctx context.Context) H265VideoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H265VideoOutput)
}

type H265VideoOutput struct{ *pulumi.OutputState }

func (H265VideoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*H265Video)(nil)).Elem()
}

func (o H265VideoOutput) ToH265VideoOutput() H265VideoOutput {
	return o
}

func (o H265VideoOutput) ToH265VideoOutputWithContext(ctx context.Context) H265VideoOutput {
	return o
}

func (o H265VideoOutput) Complexity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Video) *string { return v.Complexity }).(pulumi.StringPtrOutput)
}

func (o H265VideoOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Video) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o H265VideoOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Video) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o H265VideoOutput) Layers() H265LayerArrayOutput {
	return o.ApplyT(func(v H265Video) []H265Layer { return v.Layers }).(H265LayerArrayOutput)
}

func (o H265VideoOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v H265Video) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o H265VideoOutput) SceneChangeDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v H265Video) *bool { return v.SceneChangeDetection }).(pulumi.BoolPtrOutput)
}

func (o H265VideoOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Video) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o H265VideoOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265Video) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type H265VideoResponse struct {
	Complexity           *string             `pulumi:"complexity"`
	KeyFrameInterval     *string             `pulumi:"keyFrameInterval"`
	Label                *string             `pulumi:"label"`
	Layers               []H265LayerResponse `pulumi:"layers"`
	OdataType            string              `pulumi:"odataType"`
	SceneChangeDetection *bool               `pulumi:"sceneChangeDetection"`
	StretchMode          *string             `pulumi:"stretchMode"`
	SyncMode             *string             `pulumi:"syncMode"`
}





type H265VideoResponseInput interface {
	pulumi.Input

	ToH265VideoResponseOutput() H265VideoResponseOutput
	ToH265VideoResponseOutputWithContext(context.Context) H265VideoResponseOutput
}

type H265VideoResponseArgs struct {
	Complexity           pulumi.StringPtrInput       `pulumi:"complexity"`
	KeyFrameInterval     pulumi.StringPtrInput       `pulumi:"keyFrameInterval"`
	Label                pulumi.StringPtrInput       `pulumi:"label"`
	Layers               H265LayerResponseArrayInput `pulumi:"layers"`
	OdataType            pulumi.StringInput          `pulumi:"odataType"`
	SceneChangeDetection pulumi.BoolPtrInput         `pulumi:"sceneChangeDetection"`
	StretchMode          pulumi.StringPtrInput       `pulumi:"stretchMode"`
	SyncMode             pulumi.StringPtrInput       `pulumi:"syncMode"`
}

func (H265VideoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*H265VideoResponse)(nil)).Elem()
}

func (i H265VideoResponseArgs) ToH265VideoResponseOutput() H265VideoResponseOutput {
	return i.ToH265VideoResponseOutputWithContext(context.Background())
}

func (i H265VideoResponseArgs) ToH265VideoResponseOutputWithContext(ctx context.Context) H265VideoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(H265VideoResponseOutput)
}

type H265VideoResponseOutput struct{ *pulumi.OutputState }

func (H265VideoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*H265VideoResponse)(nil)).Elem()
}

func (o H265VideoResponseOutput) ToH265VideoResponseOutput() H265VideoResponseOutput {
	return o
}

func (o H265VideoResponseOutput) ToH265VideoResponseOutputWithContext(ctx context.Context) H265VideoResponseOutput {
	return o
}

func (o H265VideoResponseOutput) Complexity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265VideoResponse) *string { return v.Complexity }).(pulumi.StringPtrOutput)
}

func (o H265VideoResponseOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265VideoResponse) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o H265VideoResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265VideoResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o H265VideoResponseOutput) Layers() H265LayerResponseArrayOutput {
	return o.ApplyT(func(v H265VideoResponse) []H265LayerResponse { return v.Layers }).(H265LayerResponseArrayOutput)
}

func (o H265VideoResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v H265VideoResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o H265VideoResponseOutput) SceneChangeDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v H265VideoResponse) *bool { return v.SceneChangeDetection }).(pulumi.BoolPtrOutput)
}

func (o H265VideoResponseOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265VideoResponse) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o H265VideoResponseOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v H265VideoResponse) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type Hls struct {
	FragmentsPerTsSegment *int `pulumi:"fragmentsPerTsSegment"`
}





type HlsInput interface {
	pulumi.Input

	ToHlsOutput() HlsOutput
	ToHlsOutputWithContext(context.Context) HlsOutput
}

type HlsArgs struct {
	FragmentsPerTsSegment pulumi.IntPtrInput `pulumi:"fragmentsPerTsSegment"`
}

func (HlsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Hls)(nil)).Elem()
}

func (i HlsArgs) ToHlsOutput() HlsOutput {
	return i.ToHlsOutputWithContext(context.Background())
}

func (i HlsArgs) ToHlsOutputWithContext(ctx context.Context) HlsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsOutput)
}

func (i HlsArgs) ToHlsPtrOutput() HlsPtrOutput {
	return i.ToHlsPtrOutputWithContext(context.Background())
}

func (i HlsArgs) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsOutput).ToHlsPtrOutputWithContext(ctx)
}









type HlsPtrInput interface {
	pulumi.Input

	ToHlsPtrOutput() HlsPtrOutput
	ToHlsPtrOutputWithContext(context.Context) HlsPtrOutput
}

type hlsPtrType HlsArgs

func HlsPtr(v *HlsArgs) HlsPtrInput {
	return (*hlsPtrType)(v)
}

func (*hlsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Hls)(nil)).Elem()
}

func (i *hlsPtrType) ToHlsPtrOutput() HlsPtrOutput {
	return i.ToHlsPtrOutputWithContext(context.Background())
}

func (i *hlsPtrType) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsPtrOutput)
}

type HlsOutput struct{ *pulumi.OutputState }

func (HlsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Hls)(nil)).Elem()
}

func (o HlsOutput) ToHlsOutput() HlsOutput {
	return o
}

func (o HlsOutput) ToHlsOutputWithContext(ctx context.Context) HlsOutput {
	return o
}

func (o HlsOutput) ToHlsPtrOutput() HlsPtrOutput {
	return o.ToHlsPtrOutputWithContext(context.Background())
}

func (o HlsOutput) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Hls) *Hls {
		return &v
	}).(HlsPtrOutput)
}

func (o HlsOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Hls) *int { return v.FragmentsPerTsSegment }).(pulumi.IntPtrOutput)
}

type HlsPtrOutput struct{ *pulumi.OutputState }

func (HlsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hls)(nil)).Elem()
}

func (o HlsPtrOutput) ToHlsPtrOutput() HlsPtrOutput {
	return o
}

func (o HlsPtrOutput) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return o
}

func (o HlsPtrOutput) Elem() HlsOutput {
	return o.ApplyT(func(v *Hls) Hls {
		if v != nil {
			return *v
		}
		var ret Hls
		return ret
	}).(HlsOutput)
}

func (o HlsPtrOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Hls) *int {
		if v == nil {
			return nil
		}
		return v.FragmentsPerTsSegment
	}).(pulumi.IntPtrOutput)
}

type HlsResponse struct {
	FragmentsPerTsSegment *int `pulumi:"fragmentsPerTsSegment"`
}





type HlsResponseInput interface {
	pulumi.Input

	ToHlsResponseOutput() HlsResponseOutput
	ToHlsResponseOutputWithContext(context.Context) HlsResponseOutput
}

type HlsResponseArgs struct {
	FragmentsPerTsSegment pulumi.IntPtrInput `pulumi:"fragmentsPerTsSegment"`
}

func (HlsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HlsResponse)(nil)).Elem()
}

func (i HlsResponseArgs) ToHlsResponseOutput() HlsResponseOutput {
	return i.ToHlsResponseOutputWithContext(context.Background())
}

func (i HlsResponseArgs) ToHlsResponseOutputWithContext(ctx context.Context) HlsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsResponseOutput)
}

func (i HlsResponseArgs) ToHlsResponsePtrOutput() HlsResponsePtrOutput {
	return i.ToHlsResponsePtrOutputWithContext(context.Background())
}

func (i HlsResponseArgs) ToHlsResponsePtrOutputWithContext(ctx context.Context) HlsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsResponseOutput).ToHlsResponsePtrOutputWithContext(ctx)
}









type HlsResponsePtrInput interface {
	pulumi.Input

	ToHlsResponsePtrOutput() HlsResponsePtrOutput
	ToHlsResponsePtrOutputWithContext(context.Context) HlsResponsePtrOutput
}

type hlsResponsePtrType HlsResponseArgs

func HlsResponsePtr(v *HlsResponseArgs) HlsResponsePtrInput {
	return (*hlsResponsePtrType)(v)
}

func (*hlsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HlsResponse)(nil)).Elem()
}

func (i *hlsResponsePtrType) ToHlsResponsePtrOutput() HlsResponsePtrOutput {
	return i.ToHlsResponsePtrOutputWithContext(context.Background())
}

func (i *hlsResponsePtrType) ToHlsResponsePtrOutputWithContext(ctx context.Context) HlsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsResponsePtrOutput)
}

type HlsResponseOutput struct{ *pulumi.OutputState }

func (HlsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HlsResponse)(nil)).Elem()
}

func (o HlsResponseOutput) ToHlsResponseOutput() HlsResponseOutput {
	return o
}

func (o HlsResponseOutput) ToHlsResponseOutputWithContext(ctx context.Context) HlsResponseOutput {
	return o
}

func (o HlsResponseOutput) ToHlsResponsePtrOutput() HlsResponsePtrOutput {
	return o.ToHlsResponsePtrOutputWithContext(context.Background())
}

func (o HlsResponseOutput) ToHlsResponsePtrOutputWithContext(ctx context.Context) HlsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HlsResponse) *HlsResponse {
		return &v
	}).(HlsResponsePtrOutput)
}

func (o HlsResponseOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HlsResponse) *int { return v.FragmentsPerTsSegment }).(pulumi.IntPtrOutput)
}

type HlsResponsePtrOutput struct{ *pulumi.OutputState }

func (HlsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HlsResponse)(nil)).Elem()
}

func (o HlsResponsePtrOutput) ToHlsResponsePtrOutput() HlsResponsePtrOutput {
	return o
}

func (o HlsResponsePtrOutput) ToHlsResponsePtrOutputWithContext(ctx context.Context) HlsResponsePtrOutput {
	return o
}

func (o HlsResponsePtrOutput) Elem() HlsResponseOutput {
	return o.ApplyT(func(v *HlsResponse) HlsResponse {
		if v != nil {
			return *v
		}
		var ret HlsResponse
		return ret
	}).(HlsResponseOutput)
}

func (o HlsResponsePtrOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HlsResponse) *int {
		if v == nil {
			return nil
		}
		return v.FragmentsPerTsSegment
	}).(pulumi.IntPtrOutput)
}

type IPAccessControl struct {
	Allow []IPRange `pulumi:"allow"`
}





type IPAccessControlInput interface {
	pulumi.Input

	ToIPAccessControlOutput() IPAccessControlOutput
	ToIPAccessControlOutputWithContext(context.Context) IPAccessControlOutput
}

type IPAccessControlArgs struct {
	Allow IPRangeArrayInput `pulumi:"allow"`
}

func (IPAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAccessControl)(nil)).Elem()
}

func (i IPAccessControlArgs) ToIPAccessControlOutput() IPAccessControlOutput {
	return i.ToIPAccessControlOutputWithContext(context.Background())
}

func (i IPAccessControlArgs) ToIPAccessControlOutputWithContext(ctx context.Context) IPAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlOutput)
}

func (i IPAccessControlArgs) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return i.ToIPAccessControlPtrOutputWithContext(context.Background())
}

func (i IPAccessControlArgs) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlOutput).ToIPAccessControlPtrOutputWithContext(ctx)
}









type IPAccessControlPtrInput interface {
	pulumi.Input

	ToIPAccessControlPtrOutput() IPAccessControlPtrOutput
	ToIPAccessControlPtrOutputWithContext(context.Context) IPAccessControlPtrOutput
}

type ipaccessControlPtrType IPAccessControlArgs

func IPAccessControlPtr(v *IPAccessControlArgs) IPAccessControlPtrInput {
	return (*ipaccessControlPtrType)(v)
}

func (*ipaccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAccessControl)(nil)).Elem()
}

func (i *ipaccessControlPtrType) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return i.ToIPAccessControlPtrOutputWithContext(context.Background())
}

func (i *ipaccessControlPtrType) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlPtrOutput)
}

type IPAccessControlOutput struct{ *pulumi.OutputState }

func (IPAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAccessControl)(nil)).Elem()
}

func (o IPAccessControlOutput) ToIPAccessControlOutput() IPAccessControlOutput {
	return o
}

func (o IPAccessControlOutput) ToIPAccessControlOutputWithContext(ctx context.Context) IPAccessControlOutput {
	return o
}

func (o IPAccessControlOutput) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return o.ToIPAccessControlPtrOutputWithContext(context.Background())
}

func (o IPAccessControlOutput) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPAccessControl) *IPAccessControl {
		return &v
	}).(IPAccessControlPtrOutput)
}

func (o IPAccessControlOutput) Allow() IPRangeArrayOutput {
	return o.ApplyT(func(v IPAccessControl) []IPRange { return v.Allow }).(IPRangeArrayOutput)
}

type IPAccessControlPtrOutput struct{ *pulumi.OutputState }

func (IPAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAccessControl)(nil)).Elem()
}

func (o IPAccessControlPtrOutput) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return o
}

func (o IPAccessControlPtrOutput) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return o
}

func (o IPAccessControlPtrOutput) Elem() IPAccessControlOutput {
	return o.ApplyT(func(v *IPAccessControl) IPAccessControl {
		if v != nil {
			return *v
		}
		var ret IPAccessControl
		return ret
	}).(IPAccessControlOutput)
}

func (o IPAccessControlPtrOutput) Allow() IPRangeArrayOutput {
	return o.ApplyT(func(v *IPAccessControl) []IPRange {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(IPRangeArrayOutput)
}

type IPAccessControlResponse struct {
	Allow []IPRangeResponse `pulumi:"allow"`
}





type IPAccessControlResponseInput interface {
	pulumi.Input

	ToIPAccessControlResponseOutput() IPAccessControlResponseOutput
	ToIPAccessControlResponseOutputWithContext(context.Context) IPAccessControlResponseOutput
}

type IPAccessControlResponseArgs struct {
	Allow IPRangeResponseArrayInput `pulumi:"allow"`
}

func (IPAccessControlResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAccessControlResponse)(nil)).Elem()
}

func (i IPAccessControlResponseArgs) ToIPAccessControlResponseOutput() IPAccessControlResponseOutput {
	return i.ToIPAccessControlResponseOutputWithContext(context.Background())
}

func (i IPAccessControlResponseArgs) ToIPAccessControlResponseOutputWithContext(ctx context.Context) IPAccessControlResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlResponseOutput)
}

func (i IPAccessControlResponseArgs) ToIPAccessControlResponsePtrOutput() IPAccessControlResponsePtrOutput {
	return i.ToIPAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i IPAccessControlResponseArgs) ToIPAccessControlResponsePtrOutputWithContext(ctx context.Context) IPAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlResponseOutput).ToIPAccessControlResponsePtrOutputWithContext(ctx)
}









type IPAccessControlResponsePtrInput interface {
	pulumi.Input

	ToIPAccessControlResponsePtrOutput() IPAccessControlResponsePtrOutput
	ToIPAccessControlResponsePtrOutputWithContext(context.Context) IPAccessControlResponsePtrOutput
}

type ipaccessControlResponsePtrType IPAccessControlResponseArgs

func IPAccessControlResponsePtr(v *IPAccessControlResponseArgs) IPAccessControlResponsePtrInput {
	return (*ipaccessControlResponsePtrType)(v)
}

func (*ipaccessControlResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAccessControlResponse)(nil)).Elem()
}

func (i *ipaccessControlResponsePtrType) ToIPAccessControlResponsePtrOutput() IPAccessControlResponsePtrOutput {
	return i.ToIPAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i *ipaccessControlResponsePtrType) ToIPAccessControlResponsePtrOutputWithContext(ctx context.Context) IPAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlResponsePtrOutput)
}

type IPAccessControlResponseOutput struct{ *pulumi.OutputState }

func (IPAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAccessControlResponse)(nil)).Elem()
}

func (o IPAccessControlResponseOutput) ToIPAccessControlResponseOutput() IPAccessControlResponseOutput {
	return o
}

func (o IPAccessControlResponseOutput) ToIPAccessControlResponseOutputWithContext(ctx context.Context) IPAccessControlResponseOutput {
	return o
}

func (o IPAccessControlResponseOutput) ToIPAccessControlResponsePtrOutput() IPAccessControlResponsePtrOutput {
	return o.ToIPAccessControlResponsePtrOutputWithContext(context.Background())
}

func (o IPAccessControlResponseOutput) ToIPAccessControlResponsePtrOutputWithContext(ctx context.Context) IPAccessControlResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPAccessControlResponse) *IPAccessControlResponse {
		return &v
	}).(IPAccessControlResponsePtrOutput)
}

func (o IPAccessControlResponseOutput) Allow() IPRangeResponseArrayOutput {
	return o.ApplyT(func(v IPAccessControlResponse) []IPRangeResponse { return v.Allow }).(IPRangeResponseArrayOutput)
}

type IPAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (IPAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAccessControlResponse)(nil)).Elem()
}

func (o IPAccessControlResponsePtrOutput) ToIPAccessControlResponsePtrOutput() IPAccessControlResponsePtrOutput {
	return o
}

func (o IPAccessControlResponsePtrOutput) ToIPAccessControlResponsePtrOutputWithContext(ctx context.Context) IPAccessControlResponsePtrOutput {
	return o
}

func (o IPAccessControlResponsePtrOutput) Elem() IPAccessControlResponseOutput {
	return o.ApplyT(func(v *IPAccessControlResponse) IPAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret IPAccessControlResponse
		return ret
	}).(IPAccessControlResponseOutput)
}

func (o IPAccessControlResponsePtrOutput) Allow() IPRangeResponseArrayOutput {
	return o.ApplyT(func(v *IPAccessControlResponse) []IPRangeResponse {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(IPRangeResponseArrayOutput)
}

type IPRange struct {
	Address            *string `pulumi:"address"`
	Name               *string `pulumi:"name"`
	SubnetPrefixLength *int    `pulumi:"subnetPrefixLength"`
}





type IPRangeInput interface {
	pulumi.Input

	ToIPRangeOutput() IPRangeOutput
	ToIPRangeOutputWithContext(context.Context) IPRangeOutput
}

type IPRangeArgs struct {
	Address            pulumi.StringPtrInput `pulumi:"address"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	SubnetPrefixLength pulumi.IntPtrInput    `pulumi:"subnetPrefixLength"`
}

func (IPRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRange)(nil)).Elem()
}

func (i IPRangeArgs) ToIPRangeOutput() IPRangeOutput {
	return i.ToIPRangeOutputWithContext(context.Background())
}

func (i IPRangeArgs) ToIPRangeOutputWithContext(ctx context.Context) IPRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRangeOutput)
}





type IPRangeArrayInput interface {
	pulumi.Input

	ToIPRangeArrayOutput() IPRangeArrayOutput
	ToIPRangeArrayOutputWithContext(context.Context) IPRangeArrayOutput
}

type IPRangeArray []IPRangeInput

func (IPRangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRange)(nil)).Elem()
}

func (i IPRangeArray) ToIPRangeArrayOutput() IPRangeArrayOutput {
	return i.ToIPRangeArrayOutputWithContext(context.Background())
}

func (i IPRangeArray) ToIPRangeArrayOutputWithContext(ctx context.Context) IPRangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRangeArrayOutput)
}

type IPRangeOutput struct{ *pulumi.OutputState }

func (IPRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRange)(nil)).Elem()
}

func (o IPRangeOutput) ToIPRangeOutput() IPRangeOutput {
	return o
}

func (o IPRangeOutput) ToIPRangeOutputWithContext(ctx context.Context) IPRangeOutput {
	return o
}

func (o IPRangeOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRange) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o IPRangeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRange) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IPRangeOutput) SubnetPrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IPRange) *int { return v.SubnetPrefixLength }).(pulumi.IntPtrOutput)
}

type IPRangeArrayOutput struct{ *pulumi.OutputState }

func (IPRangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRange)(nil)).Elem()
}

func (o IPRangeArrayOutput) ToIPRangeArrayOutput() IPRangeArrayOutput {
	return o
}

func (o IPRangeArrayOutput) ToIPRangeArrayOutputWithContext(ctx context.Context) IPRangeArrayOutput {
	return o
}

func (o IPRangeArrayOutput) Index(i pulumi.IntInput) IPRangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRange {
		return vs[0].([]IPRange)[vs[1].(int)]
	}).(IPRangeOutput)
}

type IPRangeResponse struct {
	Address            *string `pulumi:"address"`
	Name               *string `pulumi:"name"`
	SubnetPrefixLength *int    `pulumi:"subnetPrefixLength"`
}





type IPRangeResponseInput interface {
	pulumi.Input

	ToIPRangeResponseOutput() IPRangeResponseOutput
	ToIPRangeResponseOutputWithContext(context.Context) IPRangeResponseOutput
}

type IPRangeResponseArgs struct {
	Address            pulumi.StringPtrInput `pulumi:"address"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	SubnetPrefixLength pulumi.IntPtrInput    `pulumi:"subnetPrefixLength"`
}

func (IPRangeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRangeResponse)(nil)).Elem()
}

func (i IPRangeResponseArgs) ToIPRangeResponseOutput() IPRangeResponseOutput {
	return i.ToIPRangeResponseOutputWithContext(context.Background())
}

func (i IPRangeResponseArgs) ToIPRangeResponseOutputWithContext(ctx context.Context) IPRangeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRangeResponseOutput)
}





type IPRangeResponseArrayInput interface {
	pulumi.Input

	ToIPRangeResponseArrayOutput() IPRangeResponseArrayOutput
	ToIPRangeResponseArrayOutputWithContext(context.Context) IPRangeResponseArrayOutput
}

type IPRangeResponseArray []IPRangeResponseInput

func (IPRangeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRangeResponse)(nil)).Elem()
}

func (i IPRangeResponseArray) ToIPRangeResponseArrayOutput() IPRangeResponseArrayOutput {
	return i.ToIPRangeResponseArrayOutputWithContext(context.Background())
}

func (i IPRangeResponseArray) ToIPRangeResponseArrayOutputWithContext(ctx context.Context) IPRangeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRangeResponseArrayOutput)
}

type IPRangeResponseOutput struct{ *pulumi.OutputState }

func (IPRangeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRangeResponse)(nil)).Elem()
}

func (o IPRangeResponseOutput) ToIPRangeResponseOutput() IPRangeResponseOutput {
	return o
}

func (o IPRangeResponseOutput) ToIPRangeResponseOutputWithContext(ctx context.Context) IPRangeResponseOutput {
	return o
}

func (o IPRangeResponseOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRangeResponse) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o IPRangeResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRangeResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IPRangeResponseOutput) SubnetPrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IPRangeResponse) *int { return v.SubnetPrefixLength }).(pulumi.IntPtrOutput)
}

type IPRangeResponseArrayOutput struct{ *pulumi.OutputState }

func (IPRangeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRangeResponse)(nil)).Elem()
}

func (o IPRangeResponseArrayOutput) ToIPRangeResponseArrayOutput() IPRangeResponseArrayOutput {
	return o
}

func (o IPRangeResponseArrayOutput) ToIPRangeResponseArrayOutputWithContext(ctx context.Context) IPRangeResponseArrayOutput {
	return o
}

func (o IPRangeResponseArrayOutput) Index(i pulumi.IntInput) IPRangeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRangeResponse {
		return vs[0].([]IPRangeResponse)[vs[1].(int)]
	}).(IPRangeResponseOutput)
}

type Image struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	Range            *string `pulumi:"range"`
	Start            string  `pulumi:"start"`
	Step             *string `pulumi:"step"`
	StretchMode      *string `pulumi:"stretchMode"`
	SyncMode         *string `pulumi:"syncMode"`
}





type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(context.Context) ImageOutput
}

type ImageArgs struct {
	KeyFrameInterval pulumi.StringPtrInput `pulumi:"keyFrameInterval"`
	Label            pulumi.StringPtrInput `pulumi:"label"`
	OdataType        pulumi.StringInput    `pulumi:"odataType"`
	Range            pulumi.StringPtrInput `pulumi:"range"`
	Start            pulumi.StringInput    `pulumi:"start"`
	Step             pulumi.StringPtrInput `pulumi:"step"`
	StretchMode      pulumi.StringPtrInput `pulumi:"stretchMode"`
	SyncMode         pulumi.StringPtrInput `pulumi:"syncMode"`
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil)).Elem()
}

func (i ImageArgs) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i ImageArgs) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

func (o ImageOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Image) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o ImageOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Image) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o ImageOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v Image) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ImageOutput) Range() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Image) *string { return v.Range }).(pulumi.StringPtrOutput)
}

func (o ImageOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v Image) string { return v.Start }).(pulumi.StringOutput)
}

func (o ImageOutput) Step() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Image) *string { return v.Step }).(pulumi.StringPtrOutput)
}

func (o ImageOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Image) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o ImageOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Image) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type ImageFormat struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}





type ImageFormatInput interface {
	pulumi.Input

	ToImageFormatOutput() ImageFormatOutput
	ToImageFormatOutputWithContext(context.Context) ImageFormatOutput
}

type ImageFormatArgs struct {
	FilenamePattern pulumi.StringInput `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput `pulumi:"odataType"`
}

func (ImageFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageFormat)(nil)).Elem()
}

func (i ImageFormatArgs) ToImageFormatOutput() ImageFormatOutput {
	return i.ToImageFormatOutputWithContext(context.Background())
}

func (i ImageFormatArgs) ToImageFormatOutputWithContext(ctx context.Context) ImageFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageFormatOutput)
}

type ImageFormatOutput struct{ *pulumi.OutputState }

func (ImageFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageFormat)(nil)).Elem()
}

func (o ImageFormatOutput) ToImageFormatOutput() ImageFormatOutput {
	return o
}

func (o ImageFormatOutput) ToImageFormatOutputWithContext(ctx context.Context) ImageFormatOutput {
	return o
}

func (o ImageFormatOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v ImageFormat) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o ImageFormatOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ImageFormat) string { return v.OdataType }).(pulumi.StringOutput)
}

type ImageFormatResponse struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}





type ImageFormatResponseInput interface {
	pulumi.Input

	ToImageFormatResponseOutput() ImageFormatResponseOutput
	ToImageFormatResponseOutputWithContext(context.Context) ImageFormatResponseOutput
}

type ImageFormatResponseArgs struct {
	FilenamePattern pulumi.StringInput `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput `pulumi:"odataType"`
}

func (ImageFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageFormatResponse)(nil)).Elem()
}

func (i ImageFormatResponseArgs) ToImageFormatResponseOutput() ImageFormatResponseOutput {
	return i.ToImageFormatResponseOutputWithContext(context.Background())
}

func (i ImageFormatResponseArgs) ToImageFormatResponseOutputWithContext(ctx context.Context) ImageFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageFormatResponseOutput)
}

type ImageFormatResponseOutput struct{ *pulumi.OutputState }

func (ImageFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageFormatResponse)(nil)).Elem()
}

func (o ImageFormatResponseOutput) ToImageFormatResponseOutput() ImageFormatResponseOutput {
	return o
}

func (o ImageFormatResponseOutput) ToImageFormatResponseOutputWithContext(ctx context.Context) ImageFormatResponseOutput {
	return o
}

func (o ImageFormatResponseOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v ImageFormatResponse) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o ImageFormatResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ImageFormatResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type ImageResponse struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	Range            *string `pulumi:"range"`
	Start            string  `pulumi:"start"`
	Step             *string `pulumi:"step"`
	StretchMode      *string `pulumi:"stretchMode"`
	SyncMode         *string `pulumi:"syncMode"`
}





type ImageResponseInput interface {
	pulumi.Input

	ToImageResponseOutput() ImageResponseOutput
	ToImageResponseOutputWithContext(context.Context) ImageResponseOutput
}

type ImageResponseArgs struct {
	KeyFrameInterval pulumi.StringPtrInput `pulumi:"keyFrameInterval"`
	Label            pulumi.StringPtrInput `pulumi:"label"`
	OdataType        pulumi.StringInput    `pulumi:"odataType"`
	Range            pulumi.StringPtrInput `pulumi:"range"`
	Start            pulumi.StringInput    `pulumi:"start"`
	Step             pulumi.StringPtrInput `pulumi:"step"`
	StretchMode      pulumi.StringPtrInput `pulumi:"stretchMode"`
	SyncMode         pulumi.StringPtrInput `pulumi:"syncMode"`
}

func (ImageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageResponse)(nil)).Elem()
}

func (i ImageResponseArgs) ToImageResponseOutput() ImageResponseOutput {
	return i.ToImageResponseOutputWithContext(context.Background())
}

func (i ImageResponseArgs) ToImageResponseOutputWithContext(ctx context.Context) ImageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageResponseOutput)
}

type ImageResponseOutput struct{ *pulumi.OutputState }

func (ImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageResponse)(nil)).Elem()
}

func (o ImageResponseOutput) ToImageResponseOutput() ImageResponseOutput {
	return o
}

func (o ImageResponseOutput) ToImageResponseOutputWithContext(ctx context.Context) ImageResponseOutput {
	return o
}

func (o ImageResponseOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageResponse) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o ImageResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o ImageResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ImageResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ImageResponseOutput) Range() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageResponse) *string { return v.Range }).(pulumi.StringPtrOutput)
}

func (o ImageResponseOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v ImageResponse) string { return v.Start }).(pulumi.StringOutput)
}

func (o ImageResponseOutput) Step() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageResponse) *string { return v.Step }).(pulumi.StringPtrOutput)
}

func (o ImageResponseOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageResponse) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o ImageResponseOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageResponse) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type InputFile struct {
	Filename       *string       `pulumi:"filename"`
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}





type InputFileInput interface {
	pulumi.Input

	ToInputFileOutput() InputFileOutput
	ToInputFileOutputWithContext(context.Context) InputFileOutput
}

type InputFileArgs struct {
	Filename       pulumi.StringPtrInput `pulumi:"filename"`
	IncludedTracks pulumi.ArrayInput     `pulumi:"includedTracks"`
	OdataType      pulumi.StringInput    `pulumi:"odataType"`
}

func (InputFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputFile)(nil)).Elem()
}

func (i InputFileArgs) ToInputFileOutput() InputFileOutput {
	return i.ToInputFileOutputWithContext(context.Background())
}

func (i InputFileArgs) ToInputFileOutputWithContext(ctx context.Context) InputFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputFileOutput)
}

type InputFileOutput struct{ *pulumi.OutputState }

func (InputFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputFile)(nil)).Elem()
}

func (o InputFileOutput) ToInputFileOutput() InputFileOutput {
	return o
}

func (o InputFileOutput) ToInputFileOutputWithContext(ctx context.Context) InputFileOutput {
	return o
}

func (o InputFileOutput) Filename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputFile) *string { return v.Filename }).(pulumi.StringPtrOutput)
}

func (o InputFileOutput) IncludedTracks() pulumi.ArrayOutput {
	return o.ApplyT(func(v InputFile) []interface{} { return v.IncludedTracks }).(pulumi.ArrayOutput)
}

func (o InputFileOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v InputFile) string { return v.OdataType }).(pulumi.StringOutput)
}

type InputFileResponse struct {
	Filename       *string       `pulumi:"filename"`
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}





type InputFileResponseInput interface {
	pulumi.Input

	ToInputFileResponseOutput() InputFileResponseOutput
	ToInputFileResponseOutputWithContext(context.Context) InputFileResponseOutput
}

type InputFileResponseArgs struct {
	Filename       pulumi.StringPtrInput `pulumi:"filename"`
	IncludedTracks pulumi.ArrayInput     `pulumi:"includedTracks"`
	OdataType      pulumi.StringInput    `pulumi:"odataType"`
}

func (InputFileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputFileResponse)(nil)).Elem()
}

func (i InputFileResponseArgs) ToInputFileResponseOutput() InputFileResponseOutput {
	return i.ToInputFileResponseOutputWithContext(context.Background())
}

func (i InputFileResponseArgs) ToInputFileResponseOutputWithContext(ctx context.Context) InputFileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputFileResponseOutput)
}

type InputFileResponseOutput struct{ *pulumi.OutputState }

func (InputFileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputFileResponse)(nil)).Elem()
}

func (o InputFileResponseOutput) ToInputFileResponseOutput() InputFileResponseOutput {
	return o
}

func (o InputFileResponseOutput) ToInputFileResponseOutputWithContext(ctx context.Context) InputFileResponseOutput {
	return o
}

func (o InputFileResponseOutput) Filename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputFileResponse) *string { return v.Filename }).(pulumi.StringPtrOutput)
}

func (o InputFileResponseOutput) IncludedTracks() pulumi.ArrayOutput {
	return o.ApplyT(func(v InputFileResponse) []interface{} { return v.IncludedTracks }).(pulumi.ArrayOutput)
}

func (o InputFileResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v InputFileResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type JobErrorDetailResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}





type JobErrorDetailResponseInput interface {
	pulumi.Input

	ToJobErrorDetailResponseOutput() JobErrorDetailResponseOutput
	ToJobErrorDetailResponseOutputWithContext(context.Context) JobErrorDetailResponseOutput
}

type JobErrorDetailResponseArgs struct {
	Code    pulumi.StringInput `pulumi:"code"`
	Message pulumi.StringInput `pulumi:"message"`
}

func (JobErrorDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorDetailResponse)(nil)).Elem()
}

func (i JobErrorDetailResponseArgs) ToJobErrorDetailResponseOutput() JobErrorDetailResponseOutput {
	return i.ToJobErrorDetailResponseOutputWithContext(context.Background())
}

func (i JobErrorDetailResponseArgs) ToJobErrorDetailResponseOutputWithContext(ctx context.Context) JobErrorDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobErrorDetailResponseOutput)
}





type JobErrorDetailResponseArrayInput interface {
	pulumi.Input

	ToJobErrorDetailResponseArrayOutput() JobErrorDetailResponseArrayOutput
	ToJobErrorDetailResponseArrayOutputWithContext(context.Context) JobErrorDetailResponseArrayOutput
}

type JobErrorDetailResponseArray []JobErrorDetailResponseInput

func (JobErrorDetailResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobErrorDetailResponse)(nil)).Elem()
}

func (i JobErrorDetailResponseArray) ToJobErrorDetailResponseArrayOutput() JobErrorDetailResponseArrayOutput {
	return i.ToJobErrorDetailResponseArrayOutputWithContext(context.Background())
}

func (i JobErrorDetailResponseArray) ToJobErrorDetailResponseArrayOutputWithContext(ctx context.Context) JobErrorDetailResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobErrorDetailResponseArrayOutput)
}

type JobErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (JobErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorDetailResponse)(nil)).Elem()
}

func (o JobErrorDetailResponseOutput) ToJobErrorDetailResponseOutput() JobErrorDetailResponseOutput {
	return o
}

func (o JobErrorDetailResponseOutput) ToJobErrorDetailResponseOutputWithContext(ctx context.Context) JobErrorDetailResponseOutput {
	return o
}

func (o JobErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o JobErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

type JobErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (JobErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobErrorDetailResponse)(nil)).Elem()
}

func (o JobErrorDetailResponseArrayOutput) ToJobErrorDetailResponseArrayOutput() JobErrorDetailResponseArrayOutput {
	return o
}

func (o JobErrorDetailResponseArrayOutput) ToJobErrorDetailResponseArrayOutputWithContext(ctx context.Context) JobErrorDetailResponseArrayOutput {
	return o
}

func (o JobErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) JobErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobErrorDetailResponse {
		return vs[0].([]JobErrorDetailResponse)[vs[1].(int)]
	}).(JobErrorDetailResponseOutput)
}

type JobErrorResponse struct {
	Category string                   `pulumi:"category"`
	Code     string                   `pulumi:"code"`
	Details  []JobErrorDetailResponse `pulumi:"details"`
	Message  string                   `pulumi:"message"`
	Retry    string                   `pulumi:"retry"`
}





type JobErrorResponseInput interface {
	pulumi.Input

	ToJobErrorResponseOutput() JobErrorResponseOutput
	ToJobErrorResponseOutputWithContext(context.Context) JobErrorResponseOutput
}

type JobErrorResponseArgs struct {
	Category pulumi.StringInput               `pulumi:"category"`
	Code     pulumi.StringInput               `pulumi:"code"`
	Details  JobErrorDetailResponseArrayInput `pulumi:"details"`
	Message  pulumi.StringInput               `pulumi:"message"`
	Retry    pulumi.StringInput               `pulumi:"retry"`
}

func (JobErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorResponse)(nil)).Elem()
}

func (i JobErrorResponseArgs) ToJobErrorResponseOutput() JobErrorResponseOutput {
	return i.ToJobErrorResponseOutputWithContext(context.Background())
}

func (i JobErrorResponseArgs) ToJobErrorResponseOutputWithContext(ctx context.Context) JobErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobErrorResponseOutput)
}

type JobErrorResponseOutput struct{ *pulumi.OutputState }

func (JobErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorResponse)(nil)).Elem()
}

func (o JobErrorResponseOutput) ToJobErrorResponseOutput() JobErrorResponseOutput {
	return o
}

func (o JobErrorResponseOutput) ToJobErrorResponseOutputWithContext(ctx context.Context) JobErrorResponseOutput {
	return o
}

func (o JobErrorResponseOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Category }).(pulumi.StringOutput)
}

func (o JobErrorResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o JobErrorResponseOutput) Details() JobErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v JobErrorResponse) []JobErrorDetailResponse { return v.Details }).(JobErrorDetailResponseArrayOutput)
}

func (o JobErrorResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o JobErrorResponseOutput) Retry() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Retry }).(pulumi.StringOutput)
}

type JobInputAsset struct {
	AssetName        string        `pulumi:"assetName"`
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}





type JobInputAssetInput interface {
	pulumi.Input

	ToJobInputAssetOutput() JobInputAssetOutput
	ToJobInputAssetOutputWithContext(context.Context) JobInputAssetOutput
}

type JobInputAssetArgs struct {
	AssetName        pulumi.StringInput      `pulumi:"assetName"`
	End              pulumi.Input            `pulumi:"end"`
	Files            pulumi.StringArrayInput `pulumi:"files"`
	InputDefinitions pulumi.ArrayInput       `pulumi:"inputDefinitions"`
	Label            pulumi.StringPtrInput   `pulumi:"label"`
	OdataType        pulumi.StringInput      `pulumi:"odataType"`
	Start            pulumi.Input            `pulumi:"start"`
}

func (JobInputAssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputAsset)(nil)).Elem()
}

func (i JobInputAssetArgs) ToJobInputAssetOutput() JobInputAssetOutput {
	return i.ToJobInputAssetOutputWithContext(context.Background())
}

func (i JobInputAssetArgs) ToJobInputAssetOutputWithContext(ctx context.Context) JobInputAssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputAssetOutput)
}

type JobInputAssetOutput struct{ *pulumi.OutputState }

func (JobInputAssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputAsset)(nil)).Elem()
}

func (o JobInputAssetOutput) ToJobInputAssetOutput() JobInputAssetOutput {
	return o
}

func (o JobInputAssetOutput) ToJobInputAssetOutputWithContext(ctx context.Context) JobInputAssetOutput {
	return o
}

func (o JobInputAssetOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputAsset) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o JobInputAssetOutput) End() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputAsset) interface{} { return v.End }).(pulumi.AnyOutput)
}

func (o JobInputAssetOutput) Files() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JobInputAsset) []string { return v.Files }).(pulumi.StringArrayOutput)
}

func (o JobInputAssetOutput) InputDefinitions() pulumi.ArrayOutput {
	return o.ApplyT(func(v JobInputAsset) []interface{} { return v.InputDefinitions }).(pulumi.ArrayOutput)
}

func (o JobInputAssetOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobInputAsset) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JobInputAssetOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputAsset) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JobInputAssetOutput) Start() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputAsset) interface{} { return v.Start }).(pulumi.AnyOutput)
}

type JobInputAssetResponse struct {
	AssetName        string        `pulumi:"assetName"`
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}





type JobInputAssetResponseInput interface {
	pulumi.Input

	ToJobInputAssetResponseOutput() JobInputAssetResponseOutput
	ToJobInputAssetResponseOutputWithContext(context.Context) JobInputAssetResponseOutput
}

type JobInputAssetResponseArgs struct {
	AssetName        pulumi.StringInput      `pulumi:"assetName"`
	End              pulumi.Input            `pulumi:"end"`
	Files            pulumi.StringArrayInput `pulumi:"files"`
	InputDefinitions pulumi.ArrayInput       `pulumi:"inputDefinitions"`
	Label            pulumi.StringPtrInput   `pulumi:"label"`
	OdataType        pulumi.StringInput      `pulumi:"odataType"`
	Start            pulumi.Input            `pulumi:"start"`
}

func (JobInputAssetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputAssetResponse)(nil)).Elem()
}

func (i JobInputAssetResponseArgs) ToJobInputAssetResponseOutput() JobInputAssetResponseOutput {
	return i.ToJobInputAssetResponseOutputWithContext(context.Background())
}

func (i JobInputAssetResponseArgs) ToJobInputAssetResponseOutputWithContext(ctx context.Context) JobInputAssetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputAssetResponseOutput)
}

type JobInputAssetResponseOutput struct{ *pulumi.OutputState }

func (JobInputAssetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputAssetResponse)(nil)).Elem()
}

func (o JobInputAssetResponseOutput) ToJobInputAssetResponseOutput() JobInputAssetResponseOutput {
	return o
}

func (o JobInputAssetResponseOutput) ToJobInputAssetResponseOutputWithContext(ctx context.Context) JobInputAssetResponseOutput {
	return o
}

func (o JobInputAssetResponseOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputAssetResponse) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o JobInputAssetResponseOutput) End() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputAssetResponse) interface{} { return v.End }).(pulumi.AnyOutput)
}

func (o JobInputAssetResponseOutput) Files() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JobInputAssetResponse) []string { return v.Files }).(pulumi.StringArrayOutput)
}

func (o JobInputAssetResponseOutput) InputDefinitions() pulumi.ArrayOutput {
	return o.ApplyT(func(v JobInputAssetResponse) []interface{} { return v.InputDefinitions }).(pulumi.ArrayOutput)
}

func (o JobInputAssetResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobInputAssetResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JobInputAssetResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputAssetResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JobInputAssetResponseOutput) Start() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputAssetResponse) interface{} { return v.Start }).(pulumi.AnyOutput)
}

type JobInputClip struct {
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}





type JobInputClipInput interface {
	pulumi.Input

	ToJobInputClipOutput() JobInputClipOutput
	ToJobInputClipOutputWithContext(context.Context) JobInputClipOutput
}

type JobInputClipArgs struct {
	End              pulumi.Input            `pulumi:"end"`
	Files            pulumi.StringArrayInput `pulumi:"files"`
	InputDefinitions pulumi.ArrayInput       `pulumi:"inputDefinitions"`
	Label            pulumi.StringPtrInput   `pulumi:"label"`
	OdataType        pulumi.StringInput      `pulumi:"odataType"`
	Start            pulumi.Input            `pulumi:"start"`
}

func (JobInputClipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputClip)(nil)).Elem()
}

func (i JobInputClipArgs) ToJobInputClipOutput() JobInputClipOutput {
	return i.ToJobInputClipOutputWithContext(context.Background())
}

func (i JobInputClipArgs) ToJobInputClipOutputWithContext(ctx context.Context) JobInputClipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputClipOutput)
}





type JobInputClipArrayInput interface {
	pulumi.Input

	ToJobInputClipArrayOutput() JobInputClipArrayOutput
	ToJobInputClipArrayOutputWithContext(context.Context) JobInputClipArrayOutput
}

type JobInputClipArray []JobInputClipInput

func (JobInputClipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobInputClip)(nil)).Elem()
}

func (i JobInputClipArray) ToJobInputClipArrayOutput() JobInputClipArrayOutput {
	return i.ToJobInputClipArrayOutputWithContext(context.Background())
}

func (i JobInputClipArray) ToJobInputClipArrayOutputWithContext(ctx context.Context) JobInputClipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputClipArrayOutput)
}

type JobInputClipOutput struct{ *pulumi.OutputState }

func (JobInputClipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputClip)(nil)).Elem()
}

func (o JobInputClipOutput) ToJobInputClipOutput() JobInputClipOutput {
	return o
}

func (o JobInputClipOutput) ToJobInputClipOutputWithContext(ctx context.Context) JobInputClipOutput {
	return o
}

func (o JobInputClipOutput) End() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputClip) interface{} { return v.End }).(pulumi.AnyOutput)
}

func (o JobInputClipOutput) Files() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JobInputClip) []string { return v.Files }).(pulumi.StringArrayOutput)
}

func (o JobInputClipOutput) InputDefinitions() pulumi.ArrayOutput {
	return o.ApplyT(func(v JobInputClip) []interface{} { return v.InputDefinitions }).(pulumi.ArrayOutput)
}

func (o JobInputClipOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobInputClip) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JobInputClipOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputClip) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JobInputClipOutput) Start() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputClip) interface{} { return v.Start }).(pulumi.AnyOutput)
}

type JobInputClipArrayOutput struct{ *pulumi.OutputState }

func (JobInputClipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobInputClip)(nil)).Elem()
}

func (o JobInputClipArrayOutput) ToJobInputClipArrayOutput() JobInputClipArrayOutput {
	return o
}

func (o JobInputClipArrayOutput) ToJobInputClipArrayOutputWithContext(ctx context.Context) JobInputClipArrayOutput {
	return o
}

func (o JobInputClipArrayOutput) Index(i pulumi.IntInput) JobInputClipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobInputClip {
		return vs[0].([]JobInputClip)[vs[1].(int)]
	}).(JobInputClipOutput)
}

type JobInputClipResponse struct {
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}





type JobInputClipResponseInput interface {
	pulumi.Input

	ToJobInputClipResponseOutput() JobInputClipResponseOutput
	ToJobInputClipResponseOutputWithContext(context.Context) JobInputClipResponseOutput
}

type JobInputClipResponseArgs struct {
	End              pulumi.Input            `pulumi:"end"`
	Files            pulumi.StringArrayInput `pulumi:"files"`
	InputDefinitions pulumi.ArrayInput       `pulumi:"inputDefinitions"`
	Label            pulumi.StringPtrInput   `pulumi:"label"`
	OdataType        pulumi.StringInput      `pulumi:"odataType"`
	Start            pulumi.Input            `pulumi:"start"`
}

func (JobInputClipResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputClipResponse)(nil)).Elem()
}

func (i JobInputClipResponseArgs) ToJobInputClipResponseOutput() JobInputClipResponseOutput {
	return i.ToJobInputClipResponseOutputWithContext(context.Background())
}

func (i JobInputClipResponseArgs) ToJobInputClipResponseOutputWithContext(ctx context.Context) JobInputClipResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputClipResponseOutput)
}





type JobInputClipResponseArrayInput interface {
	pulumi.Input

	ToJobInputClipResponseArrayOutput() JobInputClipResponseArrayOutput
	ToJobInputClipResponseArrayOutputWithContext(context.Context) JobInputClipResponseArrayOutput
}

type JobInputClipResponseArray []JobInputClipResponseInput

func (JobInputClipResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobInputClipResponse)(nil)).Elem()
}

func (i JobInputClipResponseArray) ToJobInputClipResponseArrayOutput() JobInputClipResponseArrayOutput {
	return i.ToJobInputClipResponseArrayOutputWithContext(context.Background())
}

func (i JobInputClipResponseArray) ToJobInputClipResponseArrayOutputWithContext(ctx context.Context) JobInputClipResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputClipResponseArrayOutput)
}

type JobInputClipResponseOutput struct{ *pulumi.OutputState }

func (JobInputClipResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputClipResponse)(nil)).Elem()
}

func (o JobInputClipResponseOutput) ToJobInputClipResponseOutput() JobInputClipResponseOutput {
	return o
}

func (o JobInputClipResponseOutput) ToJobInputClipResponseOutputWithContext(ctx context.Context) JobInputClipResponseOutput {
	return o
}

func (o JobInputClipResponseOutput) End() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputClipResponse) interface{} { return v.End }).(pulumi.AnyOutput)
}

func (o JobInputClipResponseOutput) Files() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JobInputClipResponse) []string { return v.Files }).(pulumi.StringArrayOutput)
}

func (o JobInputClipResponseOutput) InputDefinitions() pulumi.ArrayOutput {
	return o.ApplyT(func(v JobInputClipResponse) []interface{} { return v.InputDefinitions }).(pulumi.ArrayOutput)
}

func (o JobInputClipResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobInputClipResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JobInputClipResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputClipResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JobInputClipResponseOutput) Start() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputClipResponse) interface{} { return v.Start }).(pulumi.AnyOutput)
}

type JobInputClipResponseArrayOutput struct{ *pulumi.OutputState }

func (JobInputClipResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobInputClipResponse)(nil)).Elem()
}

func (o JobInputClipResponseArrayOutput) ToJobInputClipResponseArrayOutput() JobInputClipResponseArrayOutput {
	return o
}

func (o JobInputClipResponseArrayOutput) ToJobInputClipResponseArrayOutputWithContext(ctx context.Context) JobInputClipResponseArrayOutput {
	return o
}

func (o JobInputClipResponseArrayOutput) Index(i pulumi.IntInput) JobInputClipResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobInputClipResponse {
		return vs[0].([]JobInputClipResponse)[vs[1].(int)]
	}).(JobInputClipResponseOutput)
}

type JobInputHttp struct {
	BaseUri          *string       `pulumi:"baseUri"`
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}





type JobInputHttpInput interface {
	pulumi.Input

	ToJobInputHttpOutput() JobInputHttpOutput
	ToJobInputHttpOutputWithContext(context.Context) JobInputHttpOutput
}

type JobInputHttpArgs struct {
	BaseUri          pulumi.StringPtrInput   `pulumi:"baseUri"`
	End              pulumi.Input            `pulumi:"end"`
	Files            pulumi.StringArrayInput `pulumi:"files"`
	InputDefinitions pulumi.ArrayInput       `pulumi:"inputDefinitions"`
	Label            pulumi.StringPtrInput   `pulumi:"label"`
	OdataType        pulumi.StringInput      `pulumi:"odataType"`
	Start            pulumi.Input            `pulumi:"start"`
}

func (JobInputHttpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputHttp)(nil)).Elem()
}

func (i JobInputHttpArgs) ToJobInputHttpOutput() JobInputHttpOutput {
	return i.ToJobInputHttpOutputWithContext(context.Background())
}

func (i JobInputHttpArgs) ToJobInputHttpOutputWithContext(ctx context.Context) JobInputHttpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputHttpOutput)
}

type JobInputHttpOutput struct{ *pulumi.OutputState }

func (JobInputHttpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputHttp)(nil)).Elem()
}

func (o JobInputHttpOutput) ToJobInputHttpOutput() JobInputHttpOutput {
	return o
}

func (o JobInputHttpOutput) ToJobInputHttpOutputWithContext(ctx context.Context) JobInputHttpOutput {
	return o
}

func (o JobInputHttpOutput) BaseUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobInputHttp) *string { return v.BaseUri }).(pulumi.StringPtrOutput)
}

func (o JobInputHttpOutput) End() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputHttp) interface{} { return v.End }).(pulumi.AnyOutput)
}

func (o JobInputHttpOutput) Files() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JobInputHttp) []string { return v.Files }).(pulumi.StringArrayOutput)
}

func (o JobInputHttpOutput) InputDefinitions() pulumi.ArrayOutput {
	return o.ApplyT(func(v JobInputHttp) []interface{} { return v.InputDefinitions }).(pulumi.ArrayOutput)
}

func (o JobInputHttpOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobInputHttp) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JobInputHttpOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputHttp) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JobInputHttpOutput) Start() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputHttp) interface{} { return v.Start }).(pulumi.AnyOutput)
}

type JobInputHttpResponse struct {
	BaseUri          *string       `pulumi:"baseUri"`
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}





type JobInputHttpResponseInput interface {
	pulumi.Input

	ToJobInputHttpResponseOutput() JobInputHttpResponseOutput
	ToJobInputHttpResponseOutputWithContext(context.Context) JobInputHttpResponseOutput
}

type JobInputHttpResponseArgs struct {
	BaseUri          pulumi.StringPtrInput   `pulumi:"baseUri"`
	End              pulumi.Input            `pulumi:"end"`
	Files            pulumi.StringArrayInput `pulumi:"files"`
	InputDefinitions pulumi.ArrayInput       `pulumi:"inputDefinitions"`
	Label            pulumi.StringPtrInput   `pulumi:"label"`
	OdataType        pulumi.StringInput      `pulumi:"odataType"`
	Start            pulumi.Input            `pulumi:"start"`
}

func (JobInputHttpResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputHttpResponse)(nil)).Elem()
}

func (i JobInputHttpResponseArgs) ToJobInputHttpResponseOutput() JobInputHttpResponseOutput {
	return i.ToJobInputHttpResponseOutputWithContext(context.Background())
}

func (i JobInputHttpResponseArgs) ToJobInputHttpResponseOutputWithContext(ctx context.Context) JobInputHttpResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputHttpResponseOutput)
}

type JobInputHttpResponseOutput struct{ *pulumi.OutputState }

func (JobInputHttpResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputHttpResponse)(nil)).Elem()
}

func (o JobInputHttpResponseOutput) ToJobInputHttpResponseOutput() JobInputHttpResponseOutput {
	return o
}

func (o JobInputHttpResponseOutput) ToJobInputHttpResponseOutputWithContext(ctx context.Context) JobInputHttpResponseOutput {
	return o
}

func (o JobInputHttpResponseOutput) BaseUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobInputHttpResponse) *string { return v.BaseUri }).(pulumi.StringPtrOutput)
}

func (o JobInputHttpResponseOutput) End() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputHttpResponse) interface{} { return v.End }).(pulumi.AnyOutput)
}

func (o JobInputHttpResponseOutput) Files() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JobInputHttpResponse) []string { return v.Files }).(pulumi.StringArrayOutput)
}

func (o JobInputHttpResponseOutput) InputDefinitions() pulumi.ArrayOutput {
	return o.ApplyT(func(v JobInputHttpResponse) []interface{} { return v.InputDefinitions }).(pulumi.ArrayOutput)
}

func (o JobInputHttpResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobInputHttpResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JobInputHttpResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputHttpResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JobInputHttpResponseOutput) Start() pulumi.AnyOutput {
	return o.ApplyT(func(v JobInputHttpResponse) interface{} { return v.Start }).(pulumi.AnyOutput)
}

type JobInputSequence struct {
	Inputs    []JobInputClip `pulumi:"inputs"`
	OdataType string         `pulumi:"odataType"`
}





type JobInputSequenceInput interface {
	pulumi.Input

	ToJobInputSequenceOutput() JobInputSequenceOutput
	ToJobInputSequenceOutputWithContext(context.Context) JobInputSequenceOutput
}

type JobInputSequenceArgs struct {
	Inputs    JobInputClipArrayInput `pulumi:"inputs"`
	OdataType pulumi.StringInput     `pulumi:"odataType"`
}

func (JobInputSequenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputSequence)(nil)).Elem()
}

func (i JobInputSequenceArgs) ToJobInputSequenceOutput() JobInputSequenceOutput {
	return i.ToJobInputSequenceOutputWithContext(context.Background())
}

func (i JobInputSequenceArgs) ToJobInputSequenceOutputWithContext(ctx context.Context) JobInputSequenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputSequenceOutput)
}

type JobInputSequenceOutput struct{ *pulumi.OutputState }

func (JobInputSequenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputSequence)(nil)).Elem()
}

func (o JobInputSequenceOutput) ToJobInputSequenceOutput() JobInputSequenceOutput {
	return o
}

func (o JobInputSequenceOutput) ToJobInputSequenceOutputWithContext(ctx context.Context) JobInputSequenceOutput {
	return o
}

func (o JobInputSequenceOutput) Inputs() JobInputClipArrayOutput {
	return o.ApplyT(func(v JobInputSequence) []JobInputClip { return v.Inputs }).(JobInputClipArrayOutput)
}

func (o JobInputSequenceOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputSequence) string { return v.OdataType }).(pulumi.StringOutput)
}

type JobInputSequenceResponse struct {
	Inputs    []JobInputClipResponse `pulumi:"inputs"`
	OdataType string                 `pulumi:"odataType"`
}





type JobInputSequenceResponseInput interface {
	pulumi.Input

	ToJobInputSequenceResponseOutput() JobInputSequenceResponseOutput
	ToJobInputSequenceResponseOutputWithContext(context.Context) JobInputSequenceResponseOutput
}

type JobInputSequenceResponseArgs struct {
	Inputs    JobInputClipResponseArrayInput `pulumi:"inputs"`
	OdataType pulumi.StringInput             `pulumi:"odataType"`
}

func (JobInputSequenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputSequenceResponse)(nil)).Elem()
}

func (i JobInputSequenceResponseArgs) ToJobInputSequenceResponseOutput() JobInputSequenceResponseOutput {
	return i.ToJobInputSequenceResponseOutputWithContext(context.Background())
}

func (i JobInputSequenceResponseArgs) ToJobInputSequenceResponseOutputWithContext(ctx context.Context) JobInputSequenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputSequenceResponseOutput)
}

type JobInputSequenceResponseOutput struct{ *pulumi.OutputState }

func (JobInputSequenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputSequenceResponse)(nil)).Elem()
}

func (o JobInputSequenceResponseOutput) ToJobInputSequenceResponseOutput() JobInputSequenceResponseOutput {
	return o
}

func (o JobInputSequenceResponseOutput) ToJobInputSequenceResponseOutputWithContext(ctx context.Context) JobInputSequenceResponseOutput {
	return o
}

func (o JobInputSequenceResponseOutput) Inputs() JobInputClipResponseArrayOutput {
	return o.ApplyT(func(v JobInputSequenceResponse) []JobInputClipResponse { return v.Inputs }).(JobInputClipResponseArrayOutput)
}

func (o JobInputSequenceResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputSequenceResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type JobInputs struct {
	Inputs    []interface{} `pulumi:"inputs"`
	OdataType string        `pulumi:"odataType"`
}





type JobInputsInput interface {
	pulumi.Input

	ToJobInputsOutput() JobInputsOutput
	ToJobInputsOutputWithContext(context.Context) JobInputsOutput
}

type JobInputsArgs struct {
	Inputs    pulumi.ArrayInput  `pulumi:"inputs"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (JobInputsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputs)(nil)).Elem()
}

func (i JobInputsArgs) ToJobInputsOutput() JobInputsOutput {
	return i.ToJobInputsOutputWithContext(context.Background())
}

func (i JobInputsArgs) ToJobInputsOutputWithContext(ctx context.Context) JobInputsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputsOutput)
}

type JobInputsOutput struct{ *pulumi.OutputState }

func (JobInputsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputs)(nil)).Elem()
}

func (o JobInputsOutput) ToJobInputsOutput() JobInputsOutput {
	return o
}

func (o JobInputsOutput) ToJobInputsOutputWithContext(ctx context.Context) JobInputsOutput {
	return o
}

func (o JobInputsOutput) Inputs() pulumi.ArrayOutput {
	return o.ApplyT(func(v JobInputs) []interface{} { return v.Inputs }).(pulumi.ArrayOutput)
}

func (o JobInputsOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputs) string { return v.OdataType }).(pulumi.StringOutput)
}

type JobInputsResponse struct {
	Inputs    []interface{} `pulumi:"inputs"`
	OdataType string        `pulumi:"odataType"`
}





type JobInputsResponseInput interface {
	pulumi.Input

	ToJobInputsResponseOutput() JobInputsResponseOutput
	ToJobInputsResponseOutputWithContext(context.Context) JobInputsResponseOutput
}

type JobInputsResponseArgs struct {
	Inputs    pulumi.ArrayInput  `pulumi:"inputs"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (JobInputsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputsResponse)(nil)).Elem()
}

func (i JobInputsResponseArgs) ToJobInputsResponseOutput() JobInputsResponseOutput {
	return i.ToJobInputsResponseOutputWithContext(context.Background())
}

func (i JobInputsResponseArgs) ToJobInputsResponseOutputWithContext(ctx context.Context) JobInputsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobInputsResponseOutput)
}

type JobInputsResponseOutput struct{ *pulumi.OutputState }

func (JobInputsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobInputsResponse)(nil)).Elem()
}

func (o JobInputsResponseOutput) ToJobInputsResponseOutput() JobInputsResponseOutput {
	return o
}

func (o JobInputsResponseOutput) ToJobInputsResponseOutputWithContext(ctx context.Context) JobInputsResponseOutput {
	return o
}

func (o JobInputsResponseOutput) Inputs() pulumi.ArrayOutput {
	return o.ApplyT(func(v JobInputsResponse) []interface{} { return v.Inputs }).(pulumi.ArrayOutput)
}

func (o JobInputsResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobInputsResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type JobOutputAsset struct {
	AssetName string  `pulumi:"assetName"`
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}





type JobOutputAssetInput interface {
	pulumi.Input

	ToJobOutputAssetOutput() JobOutputAssetOutput
	ToJobOutputAssetOutputWithContext(context.Context) JobOutputAssetOutput
}

type JobOutputAssetArgs struct {
	AssetName pulumi.StringInput    `pulumi:"assetName"`
	Label     pulumi.StringPtrInput `pulumi:"label"`
	OdataType pulumi.StringInput    `pulumi:"odataType"`
}

func (JobOutputAssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobOutputAsset)(nil)).Elem()
}

func (i JobOutputAssetArgs) ToJobOutputAssetOutput() JobOutputAssetOutput {
	return i.ToJobOutputAssetOutputWithContext(context.Background())
}

func (i JobOutputAssetArgs) ToJobOutputAssetOutputWithContext(ctx context.Context) JobOutputAssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutputAssetOutput)
}





type JobOutputAssetArrayInput interface {
	pulumi.Input

	ToJobOutputAssetArrayOutput() JobOutputAssetArrayOutput
	ToJobOutputAssetArrayOutputWithContext(context.Context) JobOutputAssetArrayOutput
}

type JobOutputAssetArray []JobOutputAssetInput

func (JobOutputAssetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobOutputAsset)(nil)).Elem()
}

func (i JobOutputAssetArray) ToJobOutputAssetArrayOutput() JobOutputAssetArrayOutput {
	return i.ToJobOutputAssetArrayOutputWithContext(context.Background())
}

func (i JobOutputAssetArray) ToJobOutputAssetArrayOutputWithContext(ctx context.Context) JobOutputAssetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutputAssetArrayOutput)
}

type JobOutputAssetOutput struct{ *pulumi.OutputState }

func (JobOutputAssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobOutputAsset)(nil)).Elem()
}

func (o JobOutputAssetOutput) ToJobOutputAssetOutput() JobOutputAssetOutput {
	return o
}

func (o JobOutputAssetOutput) ToJobOutputAssetOutputWithContext(ctx context.Context) JobOutputAssetOutput {
	return o
}

func (o JobOutputAssetOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAsset) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o JobOutputAssetOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobOutputAsset) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JobOutputAssetOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAsset) string { return v.OdataType }).(pulumi.StringOutput)
}

type JobOutputAssetArrayOutput struct{ *pulumi.OutputState }

func (JobOutputAssetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobOutputAsset)(nil)).Elem()
}

func (o JobOutputAssetArrayOutput) ToJobOutputAssetArrayOutput() JobOutputAssetArrayOutput {
	return o
}

func (o JobOutputAssetArrayOutput) ToJobOutputAssetArrayOutputWithContext(ctx context.Context) JobOutputAssetArrayOutput {
	return o
}

func (o JobOutputAssetArrayOutput) Index(i pulumi.IntInput) JobOutputAssetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobOutputAsset {
		return vs[0].([]JobOutputAsset)[vs[1].(int)]
	}).(JobOutputAssetOutput)
}

type JobOutputAssetResponse struct {
	AssetName string           `pulumi:"assetName"`
	EndTime   string           `pulumi:"endTime"`
	Error     JobErrorResponse `pulumi:"error"`
	Label     *string          `pulumi:"label"`
	OdataType string           `pulumi:"odataType"`
	Progress  int              `pulumi:"progress"`
	StartTime string           `pulumi:"startTime"`
	State     string           `pulumi:"state"`
}





type JobOutputAssetResponseInput interface {
	pulumi.Input

	ToJobOutputAssetResponseOutput() JobOutputAssetResponseOutput
	ToJobOutputAssetResponseOutputWithContext(context.Context) JobOutputAssetResponseOutput
}

type JobOutputAssetResponseArgs struct {
	AssetName pulumi.StringInput    `pulumi:"assetName"`
	EndTime   pulumi.StringInput    `pulumi:"endTime"`
	Error     JobErrorResponseInput `pulumi:"error"`
	Label     pulumi.StringPtrInput `pulumi:"label"`
	OdataType pulumi.StringInput    `pulumi:"odataType"`
	Progress  pulumi.IntInput       `pulumi:"progress"`
	StartTime pulumi.StringInput    `pulumi:"startTime"`
	State     pulumi.StringInput    `pulumi:"state"`
}

func (JobOutputAssetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobOutputAssetResponse)(nil)).Elem()
}

func (i JobOutputAssetResponseArgs) ToJobOutputAssetResponseOutput() JobOutputAssetResponseOutput {
	return i.ToJobOutputAssetResponseOutputWithContext(context.Background())
}

func (i JobOutputAssetResponseArgs) ToJobOutputAssetResponseOutputWithContext(ctx context.Context) JobOutputAssetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutputAssetResponseOutput)
}





type JobOutputAssetResponseArrayInput interface {
	pulumi.Input

	ToJobOutputAssetResponseArrayOutput() JobOutputAssetResponseArrayOutput
	ToJobOutputAssetResponseArrayOutputWithContext(context.Context) JobOutputAssetResponseArrayOutput
}

type JobOutputAssetResponseArray []JobOutputAssetResponseInput

func (JobOutputAssetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobOutputAssetResponse)(nil)).Elem()
}

func (i JobOutputAssetResponseArray) ToJobOutputAssetResponseArrayOutput() JobOutputAssetResponseArrayOutput {
	return i.ToJobOutputAssetResponseArrayOutputWithContext(context.Background())
}

func (i JobOutputAssetResponseArray) ToJobOutputAssetResponseArrayOutputWithContext(ctx context.Context) JobOutputAssetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutputAssetResponseArrayOutput)
}

type JobOutputAssetResponseOutput struct{ *pulumi.OutputState }

func (JobOutputAssetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobOutputAssetResponse)(nil)).Elem()
}

func (o JobOutputAssetResponseOutput) ToJobOutputAssetResponseOutput() JobOutputAssetResponseOutput {
	return o
}

func (o JobOutputAssetResponseOutput) ToJobOutputAssetResponseOutputWithContext(ctx context.Context) JobOutputAssetResponseOutput {
	return o
}

func (o JobOutputAssetResponseOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o JobOutputAssetResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o JobOutputAssetResponseOutput) Error() JobErrorResponseOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) JobErrorResponse { return v.Error }).(JobErrorResponseOutput)
}

func (o JobOutputAssetResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JobOutputAssetResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JobOutputAssetResponseOutput) Progress() pulumi.IntOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) int { return v.Progress }).(pulumi.IntOutput)
}

func (o JobOutputAssetResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o JobOutputAssetResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.State }).(pulumi.StringOutput)
}

type JobOutputAssetResponseArrayOutput struct{ *pulumi.OutputState }

func (JobOutputAssetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobOutputAssetResponse)(nil)).Elem()
}

func (o JobOutputAssetResponseArrayOutput) ToJobOutputAssetResponseArrayOutput() JobOutputAssetResponseArrayOutput {
	return o
}

func (o JobOutputAssetResponseArrayOutput) ToJobOutputAssetResponseArrayOutputWithContext(ctx context.Context) JobOutputAssetResponseArrayOutput {
	return o
}

func (o JobOutputAssetResponseArrayOutput) Index(i pulumi.IntInput) JobOutputAssetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobOutputAssetResponse {
		return vs[0].([]JobOutputAssetResponse)[vs[1].(int)]
	}).(JobOutputAssetResponseOutput)
}

type JpgFormat struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}





type JpgFormatInput interface {
	pulumi.Input

	ToJpgFormatOutput() JpgFormatOutput
	ToJpgFormatOutputWithContext(context.Context) JpgFormatOutput
}

type JpgFormatArgs struct {
	FilenamePattern pulumi.StringInput `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput `pulumi:"odataType"`
}

func (JpgFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgFormat)(nil)).Elem()
}

func (i JpgFormatArgs) ToJpgFormatOutput() JpgFormatOutput {
	return i.ToJpgFormatOutputWithContext(context.Background())
}

func (i JpgFormatArgs) ToJpgFormatOutputWithContext(ctx context.Context) JpgFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JpgFormatOutput)
}

type JpgFormatOutput struct{ *pulumi.OutputState }

func (JpgFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgFormat)(nil)).Elem()
}

func (o JpgFormatOutput) ToJpgFormatOutput() JpgFormatOutput {
	return o
}

func (o JpgFormatOutput) ToJpgFormatOutputWithContext(ctx context.Context) JpgFormatOutput {
	return o
}

func (o JpgFormatOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v JpgFormat) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o JpgFormatOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JpgFormat) string { return v.OdataType }).(pulumi.StringOutput)
}

type JpgFormatResponse struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}





type JpgFormatResponseInput interface {
	pulumi.Input

	ToJpgFormatResponseOutput() JpgFormatResponseOutput
	ToJpgFormatResponseOutputWithContext(context.Context) JpgFormatResponseOutput
}

type JpgFormatResponseArgs struct {
	FilenamePattern pulumi.StringInput `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput `pulumi:"odataType"`
}

func (JpgFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgFormatResponse)(nil)).Elem()
}

func (i JpgFormatResponseArgs) ToJpgFormatResponseOutput() JpgFormatResponseOutput {
	return i.ToJpgFormatResponseOutputWithContext(context.Background())
}

func (i JpgFormatResponseArgs) ToJpgFormatResponseOutputWithContext(ctx context.Context) JpgFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JpgFormatResponseOutput)
}

type JpgFormatResponseOutput struct{ *pulumi.OutputState }

func (JpgFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgFormatResponse)(nil)).Elem()
}

func (o JpgFormatResponseOutput) ToJpgFormatResponseOutput() JpgFormatResponseOutput {
	return o
}

func (o JpgFormatResponseOutput) ToJpgFormatResponseOutputWithContext(ctx context.Context) JpgFormatResponseOutput {
	return o
}

func (o JpgFormatResponseOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v JpgFormatResponse) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o JpgFormatResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JpgFormatResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type JpgImage struct {
	KeyFrameInterval *string    `pulumi:"keyFrameInterval"`
	Label            *string    `pulumi:"label"`
	Layers           []JpgLayer `pulumi:"layers"`
	OdataType        string     `pulumi:"odataType"`
	Range            *string    `pulumi:"range"`
	SpriteColumn     *int       `pulumi:"spriteColumn"`
	Start            string     `pulumi:"start"`
	Step             *string    `pulumi:"step"`
	StretchMode      *string    `pulumi:"stretchMode"`
	SyncMode         *string    `pulumi:"syncMode"`
}





type JpgImageInput interface {
	pulumi.Input

	ToJpgImageOutput() JpgImageOutput
	ToJpgImageOutputWithContext(context.Context) JpgImageOutput
}

type JpgImageArgs struct {
	KeyFrameInterval pulumi.StringPtrInput `pulumi:"keyFrameInterval"`
	Label            pulumi.StringPtrInput `pulumi:"label"`
	Layers           JpgLayerArrayInput    `pulumi:"layers"`
	OdataType        pulumi.StringInput    `pulumi:"odataType"`
	Range            pulumi.StringPtrInput `pulumi:"range"`
	SpriteColumn     pulumi.IntPtrInput    `pulumi:"spriteColumn"`
	Start            pulumi.StringInput    `pulumi:"start"`
	Step             pulumi.StringPtrInput `pulumi:"step"`
	StretchMode      pulumi.StringPtrInput `pulumi:"stretchMode"`
	SyncMode         pulumi.StringPtrInput `pulumi:"syncMode"`
}

func (JpgImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgImage)(nil)).Elem()
}

func (i JpgImageArgs) ToJpgImageOutput() JpgImageOutput {
	return i.ToJpgImageOutputWithContext(context.Background())
}

func (i JpgImageArgs) ToJpgImageOutputWithContext(ctx context.Context) JpgImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JpgImageOutput)
}

type JpgImageOutput struct{ *pulumi.OutputState }

func (JpgImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgImage)(nil)).Elem()
}

func (o JpgImageOutput) ToJpgImageOutput() JpgImageOutput {
	return o
}

func (o JpgImageOutput) ToJpgImageOutputWithContext(ctx context.Context) JpgImageOutput {
	return o
}

func (o JpgImageOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImage) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o JpgImageOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImage) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JpgImageOutput) Layers() JpgLayerArrayOutput {
	return o.ApplyT(func(v JpgImage) []JpgLayer { return v.Layers }).(JpgLayerArrayOutput)
}

func (o JpgImageOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JpgImage) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JpgImageOutput) Range() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImage) *string { return v.Range }).(pulumi.StringPtrOutput)
}

func (o JpgImageOutput) SpriteColumn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JpgImage) *int { return v.SpriteColumn }).(pulumi.IntPtrOutput)
}

func (o JpgImageOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v JpgImage) string { return v.Start }).(pulumi.StringOutput)
}

func (o JpgImageOutput) Step() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImage) *string { return v.Step }).(pulumi.StringPtrOutput)
}

func (o JpgImageOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImage) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o JpgImageOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImage) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type JpgImageResponse struct {
	KeyFrameInterval *string            `pulumi:"keyFrameInterval"`
	Label            *string            `pulumi:"label"`
	Layers           []JpgLayerResponse `pulumi:"layers"`
	OdataType        string             `pulumi:"odataType"`
	Range            *string            `pulumi:"range"`
	SpriteColumn     *int               `pulumi:"spriteColumn"`
	Start            string             `pulumi:"start"`
	Step             *string            `pulumi:"step"`
	StretchMode      *string            `pulumi:"stretchMode"`
	SyncMode         *string            `pulumi:"syncMode"`
}





type JpgImageResponseInput interface {
	pulumi.Input

	ToJpgImageResponseOutput() JpgImageResponseOutput
	ToJpgImageResponseOutputWithContext(context.Context) JpgImageResponseOutput
}

type JpgImageResponseArgs struct {
	KeyFrameInterval pulumi.StringPtrInput      `pulumi:"keyFrameInterval"`
	Label            pulumi.StringPtrInput      `pulumi:"label"`
	Layers           JpgLayerResponseArrayInput `pulumi:"layers"`
	OdataType        pulumi.StringInput         `pulumi:"odataType"`
	Range            pulumi.StringPtrInput      `pulumi:"range"`
	SpriteColumn     pulumi.IntPtrInput         `pulumi:"spriteColumn"`
	Start            pulumi.StringInput         `pulumi:"start"`
	Step             pulumi.StringPtrInput      `pulumi:"step"`
	StretchMode      pulumi.StringPtrInput      `pulumi:"stretchMode"`
	SyncMode         pulumi.StringPtrInput      `pulumi:"syncMode"`
}

func (JpgImageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgImageResponse)(nil)).Elem()
}

func (i JpgImageResponseArgs) ToJpgImageResponseOutput() JpgImageResponseOutput {
	return i.ToJpgImageResponseOutputWithContext(context.Background())
}

func (i JpgImageResponseArgs) ToJpgImageResponseOutputWithContext(ctx context.Context) JpgImageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JpgImageResponseOutput)
}

type JpgImageResponseOutput struct{ *pulumi.OutputState }

func (JpgImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgImageResponse)(nil)).Elem()
}

func (o JpgImageResponseOutput) ToJpgImageResponseOutput() JpgImageResponseOutput {
	return o
}

func (o JpgImageResponseOutput) ToJpgImageResponseOutputWithContext(ctx context.Context) JpgImageResponseOutput {
	return o
}

func (o JpgImageResponseOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImageResponse) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o JpgImageResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImageResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JpgImageResponseOutput) Layers() JpgLayerResponseArrayOutput {
	return o.ApplyT(func(v JpgImageResponse) []JpgLayerResponse { return v.Layers }).(JpgLayerResponseArrayOutput)
}

func (o JpgImageResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JpgImageResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JpgImageResponseOutput) Range() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImageResponse) *string { return v.Range }).(pulumi.StringPtrOutput)
}

func (o JpgImageResponseOutput) SpriteColumn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JpgImageResponse) *int { return v.SpriteColumn }).(pulumi.IntPtrOutput)
}

func (o JpgImageResponseOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v JpgImageResponse) string { return v.Start }).(pulumi.StringOutput)
}

func (o JpgImageResponseOutput) Step() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImageResponse) *string { return v.Step }).(pulumi.StringPtrOutput)
}

func (o JpgImageResponseOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImageResponse) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o JpgImageResponseOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgImageResponse) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type JpgLayer struct {
	Height    *string `pulumi:"height"`
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
	Quality   *int    `pulumi:"quality"`
	Width     *string `pulumi:"width"`
}





type JpgLayerInput interface {
	pulumi.Input

	ToJpgLayerOutput() JpgLayerOutput
	ToJpgLayerOutputWithContext(context.Context) JpgLayerOutput
}

type JpgLayerArgs struct {
	Height    pulumi.StringPtrInput `pulumi:"height"`
	Label     pulumi.StringPtrInput `pulumi:"label"`
	OdataType pulumi.StringInput    `pulumi:"odataType"`
	Quality   pulumi.IntPtrInput    `pulumi:"quality"`
	Width     pulumi.StringPtrInput `pulumi:"width"`
}

func (JpgLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgLayer)(nil)).Elem()
}

func (i JpgLayerArgs) ToJpgLayerOutput() JpgLayerOutput {
	return i.ToJpgLayerOutputWithContext(context.Background())
}

func (i JpgLayerArgs) ToJpgLayerOutputWithContext(ctx context.Context) JpgLayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JpgLayerOutput)
}





type JpgLayerArrayInput interface {
	pulumi.Input

	ToJpgLayerArrayOutput() JpgLayerArrayOutput
	ToJpgLayerArrayOutputWithContext(context.Context) JpgLayerArrayOutput
}

type JpgLayerArray []JpgLayerInput

func (JpgLayerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JpgLayer)(nil)).Elem()
}

func (i JpgLayerArray) ToJpgLayerArrayOutput() JpgLayerArrayOutput {
	return i.ToJpgLayerArrayOutputWithContext(context.Background())
}

func (i JpgLayerArray) ToJpgLayerArrayOutputWithContext(ctx context.Context) JpgLayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JpgLayerArrayOutput)
}

type JpgLayerOutput struct{ *pulumi.OutputState }

func (JpgLayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgLayer)(nil)).Elem()
}

func (o JpgLayerOutput) ToJpgLayerOutput() JpgLayerOutput {
	return o
}

func (o JpgLayerOutput) ToJpgLayerOutputWithContext(ctx context.Context) JpgLayerOutput {
	return o
}

func (o JpgLayerOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgLayer) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o JpgLayerOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgLayer) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JpgLayerOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JpgLayer) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JpgLayerOutput) Quality() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JpgLayer) *int { return v.Quality }).(pulumi.IntPtrOutput)
}

func (o JpgLayerOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgLayer) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type JpgLayerArrayOutput struct{ *pulumi.OutputState }

func (JpgLayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JpgLayer)(nil)).Elem()
}

func (o JpgLayerArrayOutput) ToJpgLayerArrayOutput() JpgLayerArrayOutput {
	return o
}

func (o JpgLayerArrayOutput) ToJpgLayerArrayOutputWithContext(ctx context.Context) JpgLayerArrayOutput {
	return o
}

func (o JpgLayerArrayOutput) Index(i pulumi.IntInput) JpgLayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JpgLayer {
		return vs[0].([]JpgLayer)[vs[1].(int)]
	}).(JpgLayerOutput)
}

type JpgLayerResponse struct {
	Height    *string `pulumi:"height"`
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
	Quality   *int    `pulumi:"quality"`
	Width     *string `pulumi:"width"`
}





type JpgLayerResponseInput interface {
	pulumi.Input

	ToJpgLayerResponseOutput() JpgLayerResponseOutput
	ToJpgLayerResponseOutputWithContext(context.Context) JpgLayerResponseOutput
}

type JpgLayerResponseArgs struct {
	Height    pulumi.StringPtrInput `pulumi:"height"`
	Label     pulumi.StringPtrInput `pulumi:"label"`
	OdataType pulumi.StringInput    `pulumi:"odataType"`
	Quality   pulumi.IntPtrInput    `pulumi:"quality"`
	Width     pulumi.StringPtrInput `pulumi:"width"`
}

func (JpgLayerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgLayerResponse)(nil)).Elem()
}

func (i JpgLayerResponseArgs) ToJpgLayerResponseOutput() JpgLayerResponseOutput {
	return i.ToJpgLayerResponseOutputWithContext(context.Background())
}

func (i JpgLayerResponseArgs) ToJpgLayerResponseOutputWithContext(ctx context.Context) JpgLayerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JpgLayerResponseOutput)
}





type JpgLayerResponseArrayInput interface {
	pulumi.Input

	ToJpgLayerResponseArrayOutput() JpgLayerResponseArrayOutput
	ToJpgLayerResponseArrayOutputWithContext(context.Context) JpgLayerResponseArrayOutput
}

type JpgLayerResponseArray []JpgLayerResponseInput

func (JpgLayerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JpgLayerResponse)(nil)).Elem()
}

func (i JpgLayerResponseArray) ToJpgLayerResponseArrayOutput() JpgLayerResponseArrayOutput {
	return i.ToJpgLayerResponseArrayOutputWithContext(context.Background())
}

func (i JpgLayerResponseArray) ToJpgLayerResponseArrayOutputWithContext(ctx context.Context) JpgLayerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JpgLayerResponseArrayOutput)
}

type JpgLayerResponseOutput struct{ *pulumi.OutputState }

func (JpgLayerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JpgLayerResponse)(nil)).Elem()
}

func (o JpgLayerResponseOutput) ToJpgLayerResponseOutput() JpgLayerResponseOutput {
	return o
}

func (o JpgLayerResponseOutput) ToJpgLayerResponseOutputWithContext(ctx context.Context) JpgLayerResponseOutput {
	return o
}

func (o JpgLayerResponseOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgLayerResponse) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o JpgLayerResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgLayerResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JpgLayerResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JpgLayerResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JpgLayerResponseOutput) Quality() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JpgLayerResponse) *int { return v.Quality }).(pulumi.IntPtrOutput)
}

func (o JpgLayerResponseOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JpgLayerResponse) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type JpgLayerResponseArrayOutput struct{ *pulumi.OutputState }

func (JpgLayerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JpgLayerResponse)(nil)).Elem()
}

func (o JpgLayerResponseArrayOutput) ToJpgLayerResponseArrayOutput() JpgLayerResponseArrayOutput {
	return o
}

func (o JpgLayerResponseArrayOutput) ToJpgLayerResponseArrayOutputWithContext(ctx context.Context) JpgLayerResponseArrayOutput {
	return o
}

func (o JpgLayerResponseArrayOutput) Index(i pulumi.IntInput) JpgLayerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JpgLayerResponse {
		return vs[0].([]JpgLayerResponse)[vs[1].(int)]
	}).(JpgLayerResponseOutput)
}

type KeyVaultProperties struct {
	KeyIdentifier *string `pulumi:"keyIdentifier"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyIdentifier pulumi.StringPtrInput `pulumi:"keyIdentifier"`
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

func (o KeyVaultPropertiesOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
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
		return v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	CurrentKeyIdentifier string  `pulumi:"currentKeyIdentifier"`
	KeyIdentifier        *string `pulumi:"keyIdentifier"`
}





type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

type KeyVaultPropertiesResponseArgs struct {
	CurrentKeyIdentifier pulumi.StringInput    `pulumi:"currentKeyIdentifier"`
	KeyIdentifier        pulumi.StringPtrInput `pulumi:"keyIdentifier"`
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

func (o KeyVaultPropertiesResponseOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
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
		return v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

type LiveEventEncoding struct {
	EncodingType     *string `pulumi:"encodingType"`
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	PresetName       *string `pulumi:"presetName"`
	StretchMode      *string `pulumi:"stretchMode"`
}





type LiveEventEncodingInput interface {
	pulumi.Input

	ToLiveEventEncodingOutput() LiveEventEncodingOutput
	ToLiveEventEncodingOutputWithContext(context.Context) LiveEventEncodingOutput
}

type LiveEventEncodingArgs struct {
	EncodingType     pulumi.StringPtrInput `pulumi:"encodingType"`
	KeyFrameInterval pulumi.StringPtrInput `pulumi:"keyFrameInterval"`
	PresetName       pulumi.StringPtrInput `pulumi:"presetName"`
	StretchMode      pulumi.StringPtrInput `pulumi:"stretchMode"`
}

func (LiveEventEncodingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncoding)(nil)).Elem()
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingOutput() LiveEventEncodingOutput {
	return i.ToLiveEventEncodingOutputWithContext(context.Background())
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingOutputWithContext(ctx context.Context) LiveEventEncodingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingOutput)
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return i.ToLiveEventEncodingPtrOutputWithContext(context.Background())
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingOutput).ToLiveEventEncodingPtrOutputWithContext(ctx)
}









type LiveEventEncodingPtrInput interface {
	pulumi.Input

	ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput
	ToLiveEventEncodingPtrOutputWithContext(context.Context) LiveEventEncodingPtrOutput
}

type liveEventEncodingPtrType LiveEventEncodingArgs

func LiveEventEncodingPtr(v *LiveEventEncodingArgs) LiveEventEncodingPtrInput {
	return (*liveEventEncodingPtrType)(v)
}

func (*liveEventEncodingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncoding)(nil)).Elem()
}

func (i *liveEventEncodingPtrType) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return i.ToLiveEventEncodingPtrOutputWithContext(context.Background())
}

func (i *liveEventEncodingPtrType) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingPtrOutput)
}

type LiveEventEncodingOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncoding)(nil)).Elem()
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingOutput() LiveEventEncodingOutput {
	return o
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingOutputWithContext(ctx context.Context) LiveEventEncodingOutput {
	return o
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return o.ToLiveEventEncodingPtrOutputWithContext(context.Background())
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventEncoding) *LiveEventEncoding {
		return &v
	}).(LiveEventEncodingPtrOutput)
}

func (o LiveEventEncodingOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncoding) *string { return v.EncodingType }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncoding) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncoding) *string { return v.PresetName }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncoding) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

type LiveEventEncodingPtrOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncoding)(nil)).Elem()
}

func (o LiveEventEncodingPtrOutput) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return o
}

func (o LiveEventEncodingPtrOutput) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return o
}

func (o LiveEventEncodingPtrOutput) Elem() LiveEventEncodingOutput {
	return o.ApplyT(func(v *LiveEventEncoding) LiveEventEncoding {
		if v != nil {
			return *v
		}
		var ret LiveEventEncoding
		return ret
	}).(LiveEventEncodingOutput)
}

func (o LiveEventEncodingPtrOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncoding) *string {
		if v == nil {
			return nil
		}
		return v.EncodingType
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingPtrOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncoding) *string {
		if v == nil {
			return nil
		}
		return v.KeyFrameInterval
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingPtrOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncoding) *string {
		if v == nil {
			return nil
		}
		return v.PresetName
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingPtrOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncoding) *string {
		if v == nil {
			return nil
		}
		return v.StretchMode
	}).(pulumi.StringPtrOutput)
}

type LiveEventEncodingResponse struct {
	EncodingType     *string `pulumi:"encodingType"`
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	PresetName       *string `pulumi:"presetName"`
	StretchMode      *string `pulumi:"stretchMode"`
}





type LiveEventEncodingResponseInput interface {
	pulumi.Input

	ToLiveEventEncodingResponseOutput() LiveEventEncodingResponseOutput
	ToLiveEventEncodingResponseOutputWithContext(context.Context) LiveEventEncodingResponseOutput
}

type LiveEventEncodingResponseArgs struct {
	EncodingType     pulumi.StringPtrInput `pulumi:"encodingType"`
	KeyFrameInterval pulumi.StringPtrInput `pulumi:"keyFrameInterval"`
	PresetName       pulumi.StringPtrInput `pulumi:"presetName"`
	StretchMode      pulumi.StringPtrInput `pulumi:"stretchMode"`
}

func (LiveEventEncodingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncodingResponse)(nil)).Elem()
}

func (i LiveEventEncodingResponseArgs) ToLiveEventEncodingResponseOutput() LiveEventEncodingResponseOutput {
	return i.ToLiveEventEncodingResponseOutputWithContext(context.Background())
}

func (i LiveEventEncodingResponseArgs) ToLiveEventEncodingResponseOutputWithContext(ctx context.Context) LiveEventEncodingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingResponseOutput)
}

func (i LiveEventEncodingResponseArgs) ToLiveEventEncodingResponsePtrOutput() LiveEventEncodingResponsePtrOutput {
	return i.ToLiveEventEncodingResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventEncodingResponseArgs) ToLiveEventEncodingResponsePtrOutputWithContext(ctx context.Context) LiveEventEncodingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingResponseOutput).ToLiveEventEncodingResponsePtrOutputWithContext(ctx)
}









type LiveEventEncodingResponsePtrInput interface {
	pulumi.Input

	ToLiveEventEncodingResponsePtrOutput() LiveEventEncodingResponsePtrOutput
	ToLiveEventEncodingResponsePtrOutputWithContext(context.Context) LiveEventEncodingResponsePtrOutput
}

type liveEventEncodingResponsePtrType LiveEventEncodingResponseArgs

func LiveEventEncodingResponsePtr(v *LiveEventEncodingResponseArgs) LiveEventEncodingResponsePtrInput {
	return (*liveEventEncodingResponsePtrType)(v)
}

func (*liveEventEncodingResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncodingResponse)(nil)).Elem()
}

func (i *liveEventEncodingResponsePtrType) ToLiveEventEncodingResponsePtrOutput() LiveEventEncodingResponsePtrOutput {
	return i.ToLiveEventEncodingResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventEncodingResponsePtrType) ToLiveEventEncodingResponsePtrOutputWithContext(ctx context.Context) LiveEventEncodingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingResponsePtrOutput)
}

type LiveEventEncodingResponseOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncodingResponse)(nil)).Elem()
}

func (o LiveEventEncodingResponseOutput) ToLiveEventEncodingResponseOutput() LiveEventEncodingResponseOutput {
	return o
}

func (o LiveEventEncodingResponseOutput) ToLiveEventEncodingResponseOutputWithContext(ctx context.Context) LiveEventEncodingResponseOutput {
	return o
}

func (o LiveEventEncodingResponseOutput) ToLiveEventEncodingResponsePtrOutput() LiveEventEncodingResponsePtrOutput {
	return o.ToLiveEventEncodingResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventEncodingResponseOutput) ToLiveEventEncodingResponsePtrOutputWithContext(ctx context.Context) LiveEventEncodingResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventEncodingResponse) *LiveEventEncodingResponse {
		return &v
	}).(LiveEventEncodingResponsePtrOutput)
}

func (o LiveEventEncodingResponseOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.EncodingType }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponseOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponseOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.PresetName }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponseOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

type LiveEventEncodingResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncodingResponse)(nil)).Elem()
}

func (o LiveEventEncodingResponsePtrOutput) ToLiveEventEncodingResponsePtrOutput() LiveEventEncodingResponsePtrOutput {
	return o
}

func (o LiveEventEncodingResponsePtrOutput) ToLiveEventEncodingResponsePtrOutputWithContext(ctx context.Context) LiveEventEncodingResponsePtrOutput {
	return o
}

func (o LiveEventEncodingResponsePtrOutput) Elem() LiveEventEncodingResponseOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) LiveEventEncodingResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventEncodingResponse
		return ret
	}).(LiveEventEncodingResponseOutput)
}

func (o LiveEventEncodingResponsePtrOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncodingType
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponsePtrOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyFrameInterval
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponsePtrOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) *string {
		if v == nil {
			return nil
		}
		return v.PresetName
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponsePtrOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) *string {
		if v == nil {
			return nil
		}
		return v.StretchMode
	}).(pulumi.StringPtrOutput)
}

type LiveEventEndpoint struct {
	Protocol *string `pulumi:"protocol"`
	Url      *string `pulumi:"url"`
}





type LiveEventEndpointInput interface {
	pulumi.Input

	ToLiveEventEndpointOutput() LiveEventEndpointOutput
	ToLiveEventEndpointOutputWithContext(context.Context) LiveEventEndpointOutput
}

type LiveEventEndpointArgs struct {
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	Url      pulumi.StringPtrInput `pulumi:"url"`
}

func (LiveEventEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEndpoint)(nil)).Elem()
}

func (i LiveEventEndpointArgs) ToLiveEventEndpointOutput() LiveEventEndpointOutput {
	return i.ToLiveEventEndpointOutputWithContext(context.Background())
}

func (i LiveEventEndpointArgs) ToLiveEventEndpointOutputWithContext(ctx context.Context) LiveEventEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEndpointOutput)
}





type LiveEventEndpointArrayInput interface {
	pulumi.Input

	ToLiveEventEndpointArrayOutput() LiveEventEndpointArrayOutput
	ToLiveEventEndpointArrayOutputWithContext(context.Context) LiveEventEndpointArrayOutput
}

type LiveEventEndpointArray []LiveEventEndpointInput

func (LiveEventEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventEndpoint)(nil)).Elem()
}

func (i LiveEventEndpointArray) ToLiveEventEndpointArrayOutput() LiveEventEndpointArrayOutput {
	return i.ToLiveEventEndpointArrayOutputWithContext(context.Background())
}

func (i LiveEventEndpointArray) ToLiveEventEndpointArrayOutputWithContext(ctx context.Context) LiveEventEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEndpointArrayOutput)
}

type LiveEventEndpointOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEndpoint)(nil)).Elem()
}

func (o LiveEventEndpointOutput) ToLiveEventEndpointOutput() LiveEventEndpointOutput {
	return o
}

func (o LiveEventEndpointOutput) ToLiveEventEndpointOutputWithContext(ctx context.Context) LiveEventEndpointOutput {
	return o
}

func (o LiveEventEndpointOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpoint) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o LiveEventEndpointOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpoint) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type LiveEventEndpointArrayOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventEndpoint)(nil)).Elem()
}

func (o LiveEventEndpointArrayOutput) ToLiveEventEndpointArrayOutput() LiveEventEndpointArrayOutput {
	return o
}

func (o LiveEventEndpointArrayOutput) ToLiveEventEndpointArrayOutputWithContext(ctx context.Context) LiveEventEndpointArrayOutput {
	return o
}

func (o LiveEventEndpointArrayOutput) Index(i pulumi.IntInput) LiveEventEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventEndpoint {
		return vs[0].([]LiveEventEndpoint)[vs[1].(int)]
	}).(LiveEventEndpointOutput)
}

type LiveEventEndpointResponse struct {
	Protocol *string `pulumi:"protocol"`
	Url      *string `pulumi:"url"`
}





type LiveEventEndpointResponseInput interface {
	pulumi.Input

	ToLiveEventEndpointResponseOutput() LiveEventEndpointResponseOutput
	ToLiveEventEndpointResponseOutputWithContext(context.Context) LiveEventEndpointResponseOutput
}

type LiveEventEndpointResponseArgs struct {
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	Url      pulumi.StringPtrInput `pulumi:"url"`
}

func (LiveEventEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEndpointResponse)(nil)).Elem()
}

func (i LiveEventEndpointResponseArgs) ToLiveEventEndpointResponseOutput() LiveEventEndpointResponseOutput {
	return i.ToLiveEventEndpointResponseOutputWithContext(context.Background())
}

func (i LiveEventEndpointResponseArgs) ToLiveEventEndpointResponseOutputWithContext(ctx context.Context) LiveEventEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEndpointResponseOutput)
}





type LiveEventEndpointResponseArrayInput interface {
	pulumi.Input

	ToLiveEventEndpointResponseArrayOutput() LiveEventEndpointResponseArrayOutput
	ToLiveEventEndpointResponseArrayOutputWithContext(context.Context) LiveEventEndpointResponseArrayOutput
}

type LiveEventEndpointResponseArray []LiveEventEndpointResponseInput

func (LiveEventEndpointResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventEndpointResponse)(nil)).Elem()
}

func (i LiveEventEndpointResponseArray) ToLiveEventEndpointResponseArrayOutput() LiveEventEndpointResponseArrayOutput {
	return i.ToLiveEventEndpointResponseArrayOutputWithContext(context.Background())
}

func (i LiveEventEndpointResponseArray) ToLiveEventEndpointResponseArrayOutputWithContext(ctx context.Context) LiveEventEndpointResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEndpointResponseArrayOutput)
}

type LiveEventEndpointResponseOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEndpointResponse)(nil)).Elem()
}

func (o LiveEventEndpointResponseOutput) ToLiveEventEndpointResponseOutput() LiveEventEndpointResponseOutput {
	return o
}

func (o LiveEventEndpointResponseOutput) ToLiveEventEndpointResponseOutputWithContext(ctx context.Context) LiveEventEndpointResponseOutput {
	return o
}

func (o LiveEventEndpointResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpointResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o LiveEventEndpointResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpointResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type LiveEventEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventEndpointResponse)(nil)).Elem()
}

func (o LiveEventEndpointResponseArrayOutput) ToLiveEventEndpointResponseArrayOutput() LiveEventEndpointResponseArrayOutput {
	return o
}

func (o LiveEventEndpointResponseArrayOutput) ToLiveEventEndpointResponseArrayOutputWithContext(ctx context.Context) LiveEventEndpointResponseArrayOutput {
	return o
}

func (o LiveEventEndpointResponseArrayOutput) Index(i pulumi.IntInput) LiveEventEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventEndpointResponse {
		return vs[0].([]LiveEventEndpointResponse)[vs[1].(int)]
	}).(LiveEventEndpointResponseOutput)
}

type LiveEventInputType struct {
	AccessControl            *LiveEventInputAccessControl `pulumi:"accessControl"`
	AccessToken              *string                      `pulumi:"accessToken"`
	Endpoints                []LiveEventEndpoint          `pulumi:"endpoints"`
	KeyFrameIntervalDuration *string                      `pulumi:"keyFrameIntervalDuration"`
	StreamingProtocol        string                       `pulumi:"streamingProtocol"`
}





type LiveEventInputTypeInput interface {
	pulumi.Input

	ToLiveEventInputTypeOutput() LiveEventInputTypeOutput
	ToLiveEventInputTypeOutputWithContext(context.Context) LiveEventInputTypeOutput
}

type LiveEventInputTypeArgs struct {
	AccessControl            LiveEventInputAccessControlPtrInput `pulumi:"accessControl"`
	AccessToken              pulumi.StringPtrInput               `pulumi:"accessToken"`
	Endpoints                LiveEventEndpointArrayInput         `pulumi:"endpoints"`
	KeyFrameIntervalDuration pulumi.StringPtrInput               `pulumi:"keyFrameIntervalDuration"`
	StreamingProtocol        pulumi.StringInput                  `pulumi:"streamingProtocol"`
}

func (LiveEventInputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputType)(nil)).Elem()
}

func (i LiveEventInputTypeArgs) ToLiveEventInputTypeOutput() LiveEventInputTypeOutput {
	return i.ToLiveEventInputTypeOutputWithContext(context.Background())
}

func (i LiveEventInputTypeArgs) ToLiveEventInputTypeOutputWithContext(ctx context.Context) LiveEventInputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTypeOutput)
}

func (i LiveEventInputTypeArgs) ToLiveEventInputTypePtrOutput() LiveEventInputTypePtrOutput {
	return i.ToLiveEventInputTypePtrOutputWithContext(context.Background())
}

func (i LiveEventInputTypeArgs) ToLiveEventInputTypePtrOutputWithContext(ctx context.Context) LiveEventInputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTypeOutput).ToLiveEventInputTypePtrOutputWithContext(ctx)
}









type LiveEventInputTypePtrInput interface {
	pulumi.Input

	ToLiveEventInputTypePtrOutput() LiveEventInputTypePtrOutput
	ToLiveEventInputTypePtrOutputWithContext(context.Context) LiveEventInputTypePtrOutput
}

type liveEventInputTypePtrType LiveEventInputTypeArgs

func LiveEventInputTypePtr(v *LiveEventInputTypeArgs) LiveEventInputTypePtrInput {
	return (*liveEventInputTypePtrType)(v)
}

func (*liveEventInputTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputType)(nil)).Elem()
}

func (i *liveEventInputTypePtrType) ToLiveEventInputTypePtrOutput() LiveEventInputTypePtrOutput {
	return i.ToLiveEventInputTypePtrOutputWithContext(context.Background())
}

func (i *liveEventInputTypePtrType) ToLiveEventInputTypePtrOutputWithContext(ctx context.Context) LiveEventInputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTypePtrOutput)
}

type LiveEventInputTypeOutput struct{ *pulumi.OutputState }

func (LiveEventInputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputType)(nil)).Elem()
}

func (o LiveEventInputTypeOutput) ToLiveEventInputTypeOutput() LiveEventInputTypeOutput {
	return o
}

func (o LiveEventInputTypeOutput) ToLiveEventInputTypeOutputWithContext(ctx context.Context) LiveEventInputTypeOutput {
	return o
}

func (o LiveEventInputTypeOutput) ToLiveEventInputTypePtrOutput() LiveEventInputTypePtrOutput {
	return o.ToLiveEventInputTypePtrOutputWithContext(context.Background())
}

func (o LiveEventInputTypeOutput) ToLiveEventInputTypePtrOutputWithContext(ctx context.Context) LiveEventInputTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventInputType) *LiveEventInputType {
		return &v
	}).(LiveEventInputTypePtrOutput)
}

func (o LiveEventInputTypeOutput) AccessControl() LiveEventInputAccessControlPtrOutput {
	return o.ApplyT(func(v LiveEventInputType) *LiveEventInputAccessControl { return v.AccessControl }).(LiveEventInputAccessControlPtrOutput)
}

func (o LiveEventInputTypeOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputType) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTypeOutput) Endpoints() LiveEventEndpointArrayOutput {
	return o.ApplyT(func(v LiveEventInputType) []LiveEventEndpoint { return v.Endpoints }).(LiveEventEndpointArrayOutput)
}

func (o LiveEventInputTypeOutput) KeyFrameIntervalDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputType) *string { return v.KeyFrameIntervalDuration }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTypeOutput) StreamingProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LiveEventInputType) string { return v.StreamingProtocol }).(pulumi.StringOutput)
}

type LiveEventInputTypePtrOutput struct{ *pulumi.OutputState }

func (LiveEventInputTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputType)(nil)).Elem()
}

func (o LiveEventInputTypePtrOutput) ToLiveEventInputTypePtrOutput() LiveEventInputTypePtrOutput {
	return o
}

func (o LiveEventInputTypePtrOutput) ToLiveEventInputTypePtrOutputWithContext(ctx context.Context) LiveEventInputTypePtrOutput {
	return o
}

func (o LiveEventInputTypePtrOutput) Elem() LiveEventInputTypeOutput {
	return o.ApplyT(func(v *LiveEventInputType) LiveEventInputType {
		if v != nil {
			return *v
		}
		var ret LiveEventInputType
		return ret
	}).(LiveEventInputTypeOutput)
}

func (o LiveEventInputTypePtrOutput) AccessControl() LiveEventInputAccessControlPtrOutput {
	return o.ApplyT(func(v *LiveEventInputType) *LiveEventInputAccessControl {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(LiveEventInputAccessControlPtrOutput)
}

func (o LiveEventInputTypePtrOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputType) *string {
		if v == nil {
			return nil
		}
		return v.AccessToken
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTypePtrOutput) Endpoints() LiveEventEndpointArrayOutput {
	return o.ApplyT(func(v *LiveEventInputType) []LiveEventEndpoint {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(LiveEventEndpointArrayOutput)
}

func (o LiveEventInputTypePtrOutput) KeyFrameIntervalDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputType) *string {
		if v == nil {
			return nil
		}
		return v.KeyFrameIntervalDuration
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTypePtrOutput) StreamingProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputType) *string {
		if v == nil {
			return nil
		}
		return &v.StreamingProtocol
	}).(pulumi.StringPtrOutput)
}

type LiveEventInputAccessControl struct {
	Ip *IPAccessControl `pulumi:"ip"`
}





type LiveEventInputAccessControlInput interface {
	pulumi.Input

	ToLiveEventInputAccessControlOutput() LiveEventInputAccessControlOutput
	ToLiveEventInputAccessControlOutputWithContext(context.Context) LiveEventInputAccessControlOutput
}

type LiveEventInputAccessControlArgs struct {
	Ip IPAccessControlPtrInput `pulumi:"ip"`
}

func (LiveEventInputAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputAccessControl)(nil)).Elem()
}

func (i LiveEventInputAccessControlArgs) ToLiveEventInputAccessControlOutput() LiveEventInputAccessControlOutput {
	return i.ToLiveEventInputAccessControlOutputWithContext(context.Background())
}

func (i LiveEventInputAccessControlArgs) ToLiveEventInputAccessControlOutputWithContext(ctx context.Context) LiveEventInputAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlOutput)
}

func (i LiveEventInputAccessControlArgs) ToLiveEventInputAccessControlPtrOutput() LiveEventInputAccessControlPtrOutput {
	return i.ToLiveEventInputAccessControlPtrOutputWithContext(context.Background())
}

func (i LiveEventInputAccessControlArgs) ToLiveEventInputAccessControlPtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlOutput).ToLiveEventInputAccessControlPtrOutputWithContext(ctx)
}









type LiveEventInputAccessControlPtrInput interface {
	pulumi.Input

	ToLiveEventInputAccessControlPtrOutput() LiveEventInputAccessControlPtrOutput
	ToLiveEventInputAccessControlPtrOutputWithContext(context.Context) LiveEventInputAccessControlPtrOutput
}

type liveEventInputAccessControlPtrType LiveEventInputAccessControlArgs

func LiveEventInputAccessControlPtr(v *LiveEventInputAccessControlArgs) LiveEventInputAccessControlPtrInput {
	return (*liveEventInputAccessControlPtrType)(v)
}

func (*liveEventInputAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputAccessControl)(nil)).Elem()
}

func (i *liveEventInputAccessControlPtrType) ToLiveEventInputAccessControlPtrOutput() LiveEventInputAccessControlPtrOutput {
	return i.ToLiveEventInputAccessControlPtrOutputWithContext(context.Background())
}

func (i *liveEventInputAccessControlPtrType) ToLiveEventInputAccessControlPtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlPtrOutput)
}

type LiveEventInputAccessControlOutput struct{ *pulumi.OutputState }

func (LiveEventInputAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputAccessControl)(nil)).Elem()
}

func (o LiveEventInputAccessControlOutput) ToLiveEventInputAccessControlOutput() LiveEventInputAccessControlOutput {
	return o
}

func (o LiveEventInputAccessControlOutput) ToLiveEventInputAccessControlOutputWithContext(ctx context.Context) LiveEventInputAccessControlOutput {
	return o
}

func (o LiveEventInputAccessControlOutput) ToLiveEventInputAccessControlPtrOutput() LiveEventInputAccessControlPtrOutput {
	return o.ToLiveEventInputAccessControlPtrOutputWithContext(context.Background())
}

func (o LiveEventInputAccessControlOutput) ToLiveEventInputAccessControlPtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventInputAccessControl) *LiveEventInputAccessControl {
		return &v
	}).(LiveEventInputAccessControlPtrOutput)
}

func (o LiveEventInputAccessControlOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v LiveEventInputAccessControl) *IPAccessControl { return v.Ip }).(IPAccessControlPtrOutput)
}

type LiveEventInputAccessControlPtrOutput struct{ *pulumi.OutputState }

func (LiveEventInputAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputAccessControl)(nil)).Elem()
}

func (o LiveEventInputAccessControlPtrOutput) ToLiveEventInputAccessControlPtrOutput() LiveEventInputAccessControlPtrOutput {
	return o
}

func (o LiveEventInputAccessControlPtrOutput) ToLiveEventInputAccessControlPtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlPtrOutput {
	return o
}

func (o LiveEventInputAccessControlPtrOutput) Elem() LiveEventInputAccessControlOutput {
	return o.ApplyT(func(v *LiveEventInputAccessControl) LiveEventInputAccessControl {
		if v != nil {
			return *v
		}
		var ret LiveEventInputAccessControl
		return ret
	}).(LiveEventInputAccessControlOutput)
}

func (o LiveEventInputAccessControlPtrOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v *LiveEventInputAccessControl) *IPAccessControl {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlPtrOutput)
}

type LiveEventInputAccessControlResponse struct {
	Ip *IPAccessControlResponse `pulumi:"ip"`
}





type LiveEventInputAccessControlResponseInput interface {
	pulumi.Input

	ToLiveEventInputAccessControlResponseOutput() LiveEventInputAccessControlResponseOutput
	ToLiveEventInputAccessControlResponseOutputWithContext(context.Context) LiveEventInputAccessControlResponseOutput
}

type LiveEventInputAccessControlResponseArgs struct {
	Ip IPAccessControlResponsePtrInput `pulumi:"ip"`
}

func (LiveEventInputAccessControlResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputAccessControlResponse)(nil)).Elem()
}

func (i LiveEventInputAccessControlResponseArgs) ToLiveEventInputAccessControlResponseOutput() LiveEventInputAccessControlResponseOutput {
	return i.ToLiveEventInputAccessControlResponseOutputWithContext(context.Background())
}

func (i LiveEventInputAccessControlResponseArgs) ToLiveEventInputAccessControlResponseOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlResponseOutput)
}

func (i LiveEventInputAccessControlResponseArgs) ToLiveEventInputAccessControlResponsePtrOutput() LiveEventInputAccessControlResponsePtrOutput {
	return i.ToLiveEventInputAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventInputAccessControlResponseArgs) ToLiveEventInputAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlResponseOutput).ToLiveEventInputAccessControlResponsePtrOutputWithContext(ctx)
}









type LiveEventInputAccessControlResponsePtrInput interface {
	pulumi.Input

	ToLiveEventInputAccessControlResponsePtrOutput() LiveEventInputAccessControlResponsePtrOutput
	ToLiveEventInputAccessControlResponsePtrOutputWithContext(context.Context) LiveEventInputAccessControlResponsePtrOutput
}

type liveEventInputAccessControlResponsePtrType LiveEventInputAccessControlResponseArgs

func LiveEventInputAccessControlResponsePtr(v *LiveEventInputAccessControlResponseArgs) LiveEventInputAccessControlResponsePtrInput {
	return (*liveEventInputAccessControlResponsePtrType)(v)
}

func (*liveEventInputAccessControlResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputAccessControlResponse)(nil)).Elem()
}

func (i *liveEventInputAccessControlResponsePtrType) ToLiveEventInputAccessControlResponsePtrOutput() LiveEventInputAccessControlResponsePtrOutput {
	return i.ToLiveEventInputAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventInputAccessControlResponsePtrType) ToLiveEventInputAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputAccessControlResponsePtrOutput)
}

type LiveEventInputAccessControlResponseOutput struct{ *pulumi.OutputState }

func (LiveEventInputAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputAccessControlResponse)(nil)).Elem()
}

func (o LiveEventInputAccessControlResponseOutput) ToLiveEventInputAccessControlResponseOutput() LiveEventInputAccessControlResponseOutput {
	return o
}

func (o LiveEventInputAccessControlResponseOutput) ToLiveEventInputAccessControlResponseOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponseOutput {
	return o
}

func (o LiveEventInputAccessControlResponseOutput) ToLiveEventInputAccessControlResponsePtrOutput() LiveEventInputAccessControlResponsePtrOutput {
	return o.ToLiveEventInputAccessControlResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventInputAccessControlResponseOutput) ToLiveEventInputAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventInputAccessControlResponse) *LiveEventInputAccessControlResponse {
		return &v
	}).(LiveEventInputAccessControlResponsePtrOutput)
}

func (o LiveEventInputAccessControlResponseOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v LiveEventInputAccessControlResponse) *IPAccessControlResponse { return v.Ip }).(IPAccessControlResponsePtrOutput)
}

type LiveEventInputAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventInputAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputAccessControlResponse)(nil)).Elem()
}

func (o LiveEventInputAccessControlResponsePtrOutput) ToLiveEventInputAccessControlResponsePtrOutput() LiveEventInputAccessControlResponsePtrOutput {
	return o
}

func (o LiveEventInputAccessControlResponsePtrOutput) ToLiveEventInputAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventInputAccessControlResponsePtrOutput {
	return o
}

func (o LiveEventInputAccessControlResponsePtrOutput) Elem() LiveEventInputAccessControlResponseOutput {
	return o.ApplyT(func(v *LiveEventInputAccessControlResponse) LiveEventInputAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventInputAccessControlResponse
		return ret
	}).(LiveEventInputAccessControlResponseOutput)
}

func (o LiveEventInputAccessControlResponsePtrOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *LiveEventInputAccessControlResponse) *IPAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlResponsePtrOutput)
}

type LiveEventInputResponse struct {
	AccessControl            *LiveEventInputAccessControlResponse `pulumi:"accessControl"`
	AccessToken              *string                              `pulumi:"accessToken"`
	Endpoints                []LiveEventEndpointResponse          `pulumi:"endpoints"`
	KeyFrameIntervalDuration *string                              `pulumi:"keyFrameIntervalDuration"`
	StreamingProtocol        string                               `pulumi:"streamingProtocol"`
}





type LiveEventInputResponseInput interface {
	pulumi.Input

	ToLiveEventInputResponseOutput() LiveEventInputResponseOutput
	ToLiveEventInputResponseOutputWithContext(context.Context) LiveEventInputResponseOutput
}

type LiveEventInputResponseArgs struct {
	AccessControl            LiveEventInputAccessControlResponsePtrInput `pulumi:"accessControl"`
	AccessToken              pulumi.StringPtrInput                       `pulumi:"accessToken"`
	Endpoints                LiveEventEndpointResponseArrayInput         `pulumi:"endpoints"`
	KeyFrameIntervalDuration pulumi.StringPtrInput                       `pulumi:"keyFrameIntervalDuration"`
	StreamingProtocol        pulumi.StringInput                          `pulumi:"streamingProtocol"`
}

func (LiveEventInputResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputResponse)(nil)).Elem()
}

func (i LiveEventInputResponseArgs) ToLiveEventInputResponseOutput() LiveEventInputResponseOutput {
	return i.ToLiveEventInputResponseOutputWithContext(context.Background())
}

func (i LiveEventInputResponseArgs) ToLiveEventInputResponseOutputWithContext(ctx context.Context) LiveEventInputResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputResponseOutput)
}

func (i LiveEventInputResponseArgs) ToLiveEventInputResponsePtrOutput() LiveEventInputResponsePtrOutput {
	return i.ToLiveEventInputResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventInputResponseArgs) ToLiveEventInputResponsePtrOutputWithContext(ctx context.Context) LiveEventInputResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputResponseOutput).ToLiveEventInputResponsePtrOutputWithContext(ctx)
}









type LiveEventInputResponsePtrInput interface {
	pulumi.Input

	ToLiveEventInputResponsePtrOutput() LiveEventInputResponsePtrOutput
	ToLiveEventInputResponsePtrOutputWithContext(context.Context) LiveEventInputResponsePtrOutput
}

type liveEventInputResponsePtrType LiveEventInputResponseArgs

func LiveEventInputResponsePtr(v *LiveEventInputResponseArgs) LiveEventInputResponsePtrInput {
	return (*liveEventInputResponsePtrType)(v)
}

func (*liveEventInputResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputResponse)(nil)).Elem()
}

func (i *liveEventInputResponsePtrType) ToLiveEventInputResponsePtrOutput() LiveEventInputResponsePtrOutput {
	return i.ToLiveEventInputResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventInputResponsePtrType) ToLiveEventInputResponsePtrOutputWithContext(ctx context.Context) LiveEventInputResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputResponsePtrOutput)
}

type LiveEventInputResponseOutput struct{ *pulumi.OutputState }

func (LiveEventInputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputResponse)(nil)).Elem()
}

func (o LiveEventInputResponseOutput) ToLiveEventInputResponseOutput() LiveEventInputResponseOutput {
	return o
}

func (o LiveEventInputResponseOutput) ToLiveEventInputResponseOutputWithContext(ctx context.Context) LiveEventInputResponseOutput {
	return o
}

func (o LiveEventInputResponseOutput) ToLiveEventInputResponsePtrOutput() LiveEventInputResponsePtrOutput {
	return o.ToLiveEventInputResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventInputResponseOutput) ToLiveEventInputResponsePtrOutputWithContext(ctx context.Context) LiveEventInputResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventInputResponse) *LiveEventInputResponse {
		return &v
	}).(LiveEventInputResponsePtrOutput)
}

func (o LiveEventInputResponseOutput) AccessControl() LiveEventInputAccessControlResponsePtrOutput {
	return o.ApplyT(func(v LiveEventInputResponse) *LiveEventInputAccessControlResponse { return v.AccessControl }).(LiveEventInputAccessControlResponsePtrOutput)
}

func (o LiveEventInputResponseOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputResponse) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputResponseOutput) Endpoints() LiveEventEndpointResponseArrayOutput {
	return o.ApplyT(func(v LiveEventInputResponse) []LiveEventEndpointResponse { return v.Endpoints }).(LiveEventEndpointResponseArrayOutput)
}

func (o LiveEventInputResponseOutput) KeyFrameIntervalDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputResponse) *string { return v.KeyFrameIntervalDuration }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputResponseOutput) StreamingProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LiveEventInputResponse) string { return v.StreamingProtocol }).(pulumi.StringOutput)
}

type LiveEventInputResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventInputResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputResponse)(nil)).Elem()
}

func (o LiveEventInputResponsePtrOutput) ToLiveEventInputResponsePtrOutput() LiveEventInputResponsePtrOutput {
	return o
}

func (o LiveEventInputResponsePtrOutput) ToLiveEventInputResponsePtrOutputWithContext(ctx context.Context) LiveEventInputResponsePtrOutput {
	return o
}

func (o LiveEventInputResponsePtrOutput) Elem() LiveEventInputResponseOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) LiveEventInputResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventInputResponse
		return ret
	}).(LiveEventInputResponseOutput)
}

func (o LiveEventInputResponsePtrOutput) AccessControl() LiveEventInputAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) *LiveEventInputAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(LiveEventInputAccessControlResponsePtrOutput)
}

func (o LiveEventInputResponsePtrOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccessToken
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventInputResponsePtrOutput) Endpoints() LiveEventEndpointResponseArrayOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) []LiveEventEndpointResponse {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(LiveEventEndpointResponseArrayOutput)
}

func (o LiveEventInputResponsePtrOutput) KeyFrameIntervalDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyFrameIntervalDuration
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventInputResponsePtrOutput) StreamingProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventInputResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StreamingProtocol
	}).(pulumi.StringPtrOutput)
}

type LiveEventInputTrackSelection struct {
	Operation *string `pulumi:"operation"`
	Property  *string `pulumi:"property"`
	Value     *string `pulumi:"value"`
}





type LiveEventInputTrackSelectionInput interface {
	pulumi.Input

	ToLiveEventInputTrackSelectionOutput() LiveEventInputTrackSelectionOutput
	ToLiveEventInputTrackSelectionOutputWithContext(context.Context) LiveEventInputTrackSelectionOutput
}

type LiveEventInputTrackSelectionArgs struct {
	Operation pulumi.StringPtrInput `pulumi:"operation"`
	Property  pulumi.StringPtrInput `pulumi:"property"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (LiveEventInputTrackSelectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputTrackSelection)(nil)).Elem()
}

func (i LiveEventInputTrackSelectionArgs) ToLiveEventInputTrackSelectionOutput() LiveEventInputTrackSelectionOutput {
	return i.ToLiveEventInputTrackSelectionOutputWithContext(context.Background())
}

func (i LiveEventInputTrackSelectionArgs) ToLiveEventInputTrackSelectionOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTrackSelectionOutput)
}





type LiveEventInputTrackSelectionArrayInput interface {
	pulumi.Input

	ToLiveEventInputTrackSelectionArrayOutput() LiveEventInputTrackSelectionArrayOutput
	ToLiveEventInputTrackSelectionArrayOutputWithContext(context.Context) LiveEventInputTrackSelectionArrayOutput
}

type LiveEventInputTrackSelectionArray []LiveEventInputTrackSelectionInput

func (LiveEventInputTrackSelectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventInputTrackSelection)(nil)).Elem()
}

func (i LiveEventInputTrackSelectionArray) ToLiveEventInputTrackSelectionArrayOutput() LiveEventInputTrackSelectionArrayOutput {
	return i.ToLiveEventInputTrackSelectionArrayOutputWithContext(context.Background())
}

func (i LiveEventInputTrackSelectionArray) ToLiveEventInputTrackSelectionArrayOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTrackSelectionArrayOutput)
}

type LiveEventInputTrackSelectionOutput struct{ *pulumi.OutputState }

func (LiveEventInputTrackSelectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputTrackSelection)(nil)).Elem()
}

func (o LiveEventInputTrackSelectionOutput) ToLiveEventInputTrackSelectionOutput() LiveEventInputTrackSelectionOutput {
	return o
}

func (o LiveEventInputTrackSelectionOutput) ToLiveEventInputTrackSelectionOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionOutput {
	return o
}

func (o LiveEventInputTrackSelectionOutput) Operation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelection) *string { return v.Operation }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTrackSelectionOutput) Property() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelection) *string { return v.Property }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTrackSelectionOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelection) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type LiveEventInputTrackSelectionArrayOutput struct{ *pulumi.OutputState }

func (LiveEventInputTrackSelectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventInputTrackSelection)(nil)).Elem()
}

func (o LiveEventInputTrackSelectionArrayOutput) ToLiveEventInputTrackSelectionArrayOutput() LiveEventInputTrackSelectionArrayOutput {
	return o
}

func (o LiveEventInputTrackSelectionArrayOutput) ToLiveEventInputTrackSelectionArrayOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionArrayOutput {
	return o
}

func (o LiveEventInputTrackSelectionArrayOutput) Index(i pulumi.IntInput) LiveEventInputTrackSelectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventInputTrackSelection {
		return vs[0].([]LiveEventInputTrackSelection)[vs[1].(int)]
	}).(LiveEventInputTrackSelectionOutput)
}

type LiveEventInputTrackSelectionResponse struct {
	Operation *string `pulumi:"operation"`
	Property  *string `pulumi:"property"`
	Value     *string `pulumi:"value"`
}





type LiveEventInputTrackSelectionResponseInput interface {
	pulumi.Input

	ToLiveEventInputTrackSelectionResponseOutput() LiveEventInputTrackSelectionResponseOutput
	ToLiveEventInputTrackSelectionResponseOutputWithContext(context.Context) LiveEventInputTrackSelectionResponseOutput
}

type LiveEventInputTrackSelectionResponseArgs struct {
	Operation pulumi.StringPtrInput `pulumi:"operation"`
	Property  pulumi.StringPtrInput `pulumi:"property"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (LiveEventInputTrackSelectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputTrackSelectionResponse)(nil)).Elem()
}

func (i LiveEventInputTrackSelectionResponseArgs) ToLiveEventInputTrackSelectionResponseOutput() LiveEventInputTrackSelectionResponseOutput {
	return i.ToLiveEventInputTrackSelectionResponseOutputWithContext(context.Background())
}

func (i LiveEventInputTrackSelectionResponseArgs) ToLiveEventInputTrackSelectionResponseOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTrackSelectionResponseOutput)
}





type LiveEventInputTrackSelectionResponseArrayInput interface {
	pulumi.Input

	ToLiveEventInputTrackSelectionResponseArrayOutput() LiveEventInputTrackSelectionResponseArrayOutput
	ToLiveEventInputTrackSelectionResponseArrayOutputWithContext(context.Context) LiveEventInputTrackSelectionResponseArrayOutput
}

type LiveEventInputTrackSelectionResponseArray []LiveEventInputTrackSelectionResponseInput

func (LiveEventInputTrackSelectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventInputTrackSelectionResponse)(nil)).Elem()
}

func (i LiveEventInputTrackSelectionResponseArray) ToLiveEventInputTrackSelectionResponseArrayOutput() LiveEventInputTrackSelectionResponseArrayOutput {
	return i.ToLiveEventInputTrackSelectionResponseArrayOutputWithContext(context.Background())
}

func (i LiveEventInputTrackSelectionResponseArray) ToLiveEventInputTrackSelectionResponseArrayOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTrackSelectionResponseArrayOutput)
}

type LiveEventInputTrackSelectionResponseOutput struct{ *pulumi.OutputState }

func (LiveEventInputTrackSelectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputTrackSelectionResponse)(nil)).Elem()
}

func (o LiveEventInputTrackSelectionResponseOutput) ToLiveEventInputTrackSelectionResponseOutput() LiveEventInputTrackSelectionResponseOutput {
	return o
}

func (o LiveEventInputTrackSelectionResponseOutput) ToLiveEventInputTrackSelectionResponseOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionResponseOutput {
	return o
}

func (o LiveEventInputTrackSelectionResponseOutput) Operation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelectionResponse) *string { return v.Operation }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTrackSelectionResponseOutput) Property() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelectionResponse) *string { return v.Property }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTrackSelectionResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputTrackSelectionResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type LiveEventInputTrackSelectionResponseArrayOutput struct{ *pulumi.OutputState }

func (LiveEventInputTrackSelectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventInputTrackSelectionResponse)(nil)).Elem()
}

func (o LiveEventInputTrackSelectionResponseArrayOutput) ToLiveEventInputTrackSelectionResponseArrayOutput() LiveEventInputTrackSelectionResponseArrayOutput {
	return o
}

func (o LiveEventInputTrackSelectionResponseArrayOutput) ToLiveEventInputTrackSelectionResponseArrayOutputWithContext(ctx context.Context) LiveEventInputTrackSelectionResponseArrayOutput {
	return o
}

func (o LiveEventInputTrackSelectionResponseArrayOutput) Index(i pulumi.IntInput) LiveEventInputTrackSelectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventInputTrackSelectionResponse {
		return vs[0].([]LiveEventInputTrackSelectionResponse)[vs[1].(int)]
	}).(LiveEventInputTrackSelectionResponseOutput)
}

type LiveEventOutputTranscriptionTrack struct {
	TrackName string `pulumi:"trackName"`
}





type LiveEventOutputTranscriptionTrackInput interface {
	pulumi.Input

	ToLiveEventOutputTranscriptionTrackOutput() LiveEventOutputTranscriptionTrackOutput
	ToLiveEventOutputTranscriptionTrackOutputWithContext(context.Context) LiveEventOutputTranscriptionTrackOutput
}

type LiveEventOutputTranscriptionTrackArgs struct {
	TrackName pulumi.StringInput `pulumi:"trackName"`
}

func (LiveEventOutputTranscriptionTrackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventOutputTranscriptionTrack)(nil)).Elem()
}

func (i LiveEventOutputTranscriptionTrackArgs) ToLiveEventOutputTranscriptionTrackOutput() LiveEventOutputTranscriptionTrackOutput {
	return i.ToLiveEventOutputTranscriptionTrackOutputWithContext(context.Background())
}

func (i LiveEventOutputTranscriptionTrackArgs) ToLiveEventOutputTranscriptionTrackOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackOutput)
}

func (i LiveEventOutputTranscriptionTrackArgs) ToLiveEventOutputTranscriptionTrackPtrOutput() LiveEventOutputTranscriptionTrackPtrOutput {
	return i.ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(context.Background())
}

func (i LiveEventOutputTranscriptionTrackArgs) ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackOutput).ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(ctx)
}









type LiveEventOutputTranscriptionTrackPtrInput interface {
	pulumi.Input

	ToLiveEventOutputTranscriptionTrackPtrOutput() LiveEventOutputTranscriptionTrackPtrOutput
	ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(context.Context) LiveEventOutputTranscriptionTrackPtrOutput
}

type liveEventOutputTranscriptionTrackPtrType LiveEventOutputTranscriptionTrackArgs

func LiveEventOutputTranscriptionTrackPtr(v *LiveEventOutputTranscriptionTrackArgs) LiveEventOutputTranscriptionTrackPtrInput {
	return (*liveEventOutputTranscriptionTrackPtrType)(v)
}

func (*liveEventOutputTranscriptionTrackPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventOutputTranscriptionTrack)(nil)).Elem()
}

func (i *liveEventOutputTranscriptionTrackPtrType) ToLiveEventOutputTranscriptionTrackPtrOutput() LiveEventOutputTranscriptionTrackPtrOutput {
	return i.ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(context.Background())
}

func (i *liveEventOutputTranscriptionTrackPtrType) ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackPtrOutput)
}

type LiveEventOutputTranscriptionTrackOutput struct{ *pulumi.OutputState }

func (LiveEventOutputTranscriptionTrackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventOutputTranscriptionTrack)(nil)).Elem()
}

func (o LiveEventOutputTranscriptionTrackOutput) ToLiveEventOutputTranscriptionTrackOutput() LiveEventOutputTranscriptionTrackOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackOutput) ToLiveEventOutputTranscriptionTrackOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackOutput) ToLiveEventOutputTranscriptionTrackPtrOutput() LiveEventOutputTranscriptionTrackPtrOutput {
	return o.ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(context.Background())
}

func (o LiveEventOutputTranscriptionTrackOutput) ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventOutputTranscriptionTrack) *LiveEventOutputTranscriptionTrack {
		return &v
	}).(LiveEventOutputTranscriptionTrackPtrOutput)
}

func (o LiveEventOutputTranscriptionTrackOutput) TrackName() pulumi.StringOutput {
	return o.ApplyT(func(v LiveEventOutputTranscriptionTrack) string { return v.TrackName }).(pulumi.StringOutput)
}

type LiveEventOutputTranscriptionTrackPtrOutput struct{ *pulumi.OutputState }

func (LiveEventOutputTranscriptionTrackPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventOutputTranscriptionTrack)(nil)).Elem()
}

func (o LiveEventOutputTranscriptionTrackPtrOutput) ToLiveEventOutputTranscriptionTrackPtrOutput() LiveEventOutputTranscriptionTrackPtrOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackPtrOutput) ToLiveEventOutputTranscriptionTrackPtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackPtrOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackPtrOutput) Elem() LiveEventOutputTranscriptionTrackOutput {
	return o.ApplyT(func(v *LiveEventOutputTranscriptionTrack) LiveEventOutputTranscriptionTrack {
		if v != nil {
			return *v
		}
		var ret LiveEventOutputTranscriptionTrack
		return ret
	}).(LiveEventOutputTranscriptionTrackOutput)
}

func (o LiveEventOutputTranscriptionTrackPtrOutput) TrackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventOutputTranscriptionTrack) *string {
		if v == nil {
			return nil
		}
		return &v.TrackName
	}).(pulumi.StringPtrOutput)
}

type LiveEventOutputTranscriptionTrackResponse struct {
	TrackName string `pulumi:"trackName"`
}





type LiveEventOutputTranscriptionTrackResponseInput interface {
	pulumi.Input

	ToLiveEventOutputTranscriptionTrackResponseOutput() LiveEventOutputTranscriptionTrackResponseOutput
	ToLiveEventOutputTranscriptionTrackResponseOutputWithContext(context.Context) LiveEventOutputTranscriptionTrackResponseOutput
}

type LiveEventOutputTranscriptionTrackResponseArgs struct {
	TrackName pulumi.StringInput `pulumi:"trackName"`
}

func (LiveEventOutputTranscriptionTrackResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventOutputTranscriptionTrackResponse)(nil)).Elem()
}

func (i LiveEventOutputTranscriptionTrackResponseArgs) ToLiveEventOutputTranscriptionTrackResponseOutput() LiveEventOutputTranscriptionTrackResponseOutput {
	return i.ToLiveEventOutputTranscriptionTrackResponseOutputWithContext(context.Background())
}

func (i LiveEventOutputTranscriptionTrackResponseArgs) ToLiveEventOutputTranscriptionTrackResponseOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackResponseOutput)
}

func (i LiveEventOutputTranscriptionTrackResponseArgs) ToLiveEventOutputTranscriptionTrackResponsePtrOutput() LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return i.ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventOutputTranscriptionTrackResponseArgs) ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackResponseOutput).ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(ctx)
}









type LiveEventOutputTranscriptionTrackResponsePtrInput interface {
	pulumi.Input

	ToLiveEventOutputTranscriptionTrackResponsePtrOutput() LiveEventOutputTranscriptionTrackResponsePtrOutput
	ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(context.Context) LiveEventOutputTranscriptionTrackResponsePtrOutput
}

type liveEventOutputTranscriptionTrackResponsePtrType LiveEventOutputTranscriptionTrackResponseArgs

func LiveEventOutputTranscriptionTrackResponsePtr(v *LiveEventOutputTranscriptionTrackResponseArgs) LiveEventOutputTranscriptionTrackResponsePtrInput {
	return (*liveEventOutputTranscriptionTrackResponsePtrType)(v)
}

func (*liveEventOutputTranscriptionTrackResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventOutputTranscriptionTrackResponse)(nil)).Elem()
}

func (i *liveEventOutputTranscriptionTrackResponsePtrType) ToLiveEventOutputTranscriptionTrackResponsePtrOutput() LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return i.ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventOutputTranscriptionTrackResponsePtrType) ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutputTranscriptionTrackResponsePtrOutput)
}

type LiveEventOutputTranscriptionTrackResponseOutput struct{ *pulumi.OutputState }

func (LiveEventOutputTranscriptionTrackResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventOutputTranscriptionTrackResponse)(nil)).Elem()
}

func (o LiveEventOutputTranscriptionTrackResponseOutput) ToLiveEventOutputTranscriptionTrackResponseOutput() LiveEventOutputTranscriptionTrackResponseOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackResponseOutput) ToLiveEventOutputTranscriptionTrackResponseOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponseOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackResponseOutput) ToLiveEventOutputTranscriptionTrackResponsePtrOutput() LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return o.ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventOutputTranscriptionTrackResponseOutput) ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventOutputTranscriptionTrackResponse) *LiveEventOutputTranscriptionTrackResponse {
		return &v
	}).(LiveEventOutputTranscriptionTrackResponsePtrOutput)
}

func (o LiveEventOutputTranscriptionTrackResponseOutput) TrackName() pulumi.StringOutput {
	return o.ApplyT(func(v LiveEventOutputTranscriptionTrackResponse) string { return v.TrackName }).(pulumi.StringOutput)
}

type LiveEventOutputTranscriptionTrackResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventOutputTranscriptionTrackResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventOutputTranscriptionTrackResponse)(nil)).Elem()
}

func (o LiveEventOutputTranscriptionTrackResponsePtrOutput) ToLiveEventOutputTranscriptionTrackResponsePtrOutput() LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackResponsePtrOutput) ToLiveEventOutputTranscriptionTrackResponsePtrOutputWithContext(ctx context.Context) LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return o
}

func (o LiveEventOutputTranscriptionTrackResponsePtrOutput) Elem() LiveEventOutputTranscriptionTrackResponseOutput {
	return o.ApplyT(func(v *LiveEventOutputTranscriptionTrackResponse) LiveEventOutputTranscriptionTrackResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventOutputTranscriptionTrackResponse
		return ret
	}).(LiveEventOutputTranscriptionTrackResponseOutput)
}

func (o LiveEventOutputTranscriptionTrackResponsePtrOutput) TrackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventOutputTranscriptionTrackResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TrackName
	}).(pulumi.StringPtrOutput)
}

type LiveEventPreview struct {
	AccessControl       *LiveEventPreviewAccessControl `pulumi:"accessControl"`
	AlternativeMediaId  *string                        `pulumi:"alternativeMediaId"`
	Endpoints           []LiveEventEndpoint            `pulumi:"endpoints"`
	PreviewLocator      *string                        `pulumi:"previewLocator"`
	StreamingPolicyName *string                        `pulumi:"streamingPolicyName"`
}





type LiveEventPreviewInput interface {
	pulumi.Input

	ToLiveEventPreviewOutput() LiveEventPreviewOutput
	ToLiveEventPreviewOutputWithContext(context.Context) LiveEventPreviewOutput
}

type LiveEventPreviewArgs struct {
	AccessControl       LiveEventPreviewAccessControlPtrInput `pulumi:"accessControl"`
	AlternativeMediaId  pulumi.StringPtrInput                 `pulumi:"alternativeMediaId"`
	Endpoints           LiveEventEndpointArrayInput           `pulumi:"endpoints"`
	PreviewLocator      pulumi.StringPtrInput                 `pulumi:"previewLocator"`
	StreamingPolicyName pulumi.StringPtrInput                 `pulumi:"streamingPolicyName"`
}

func (LiveEventPreviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreview)(nil)).Elem()
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewOutput() LiveEventPreviewOutput {
	return i.ToLiveEventPreviewOutputWithContext(context.Background())
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewOutputWithContext(ctx context.Context) LiveEventPreviewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewOutput)
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return i.ToLiveEventPreviewPtrOutputWithContext(context.Background())
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewOutput).ToLiveEventPreviewPtrOutputWithContext(ctx)
}









type LiveEventPreviewPtrInput interface {
	pulumi.Input

	ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput
	ToLiveEventPreviewPtrOutputWithContext(context.Context) LiveEventPreviewPtrOutput
}

type liveEventPreviewPtrType LiveEventPreviewArgs

func LiveEventPreviewPtr(v *LiveEventPreviewArgs) LiveEventPreviewPtrInput {
	return (*liveEventPreviewPtrType)(v)
}

func (*liveEventPreviewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreview)(nil)).Elem()
}

func (i *liveEventPreviewPtrType) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return i.ToLiveEventPreviewPtrOutputWithContext(context.Background())
}

func (i *liveEventPreviewPtrType) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewPtrOutput)
}

type LiveEventPreviewOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreview)(nil)).Elem()
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewOutput() LiveEventPreviewOutput {
	return o
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewOutputWithContext(ctx context.Context) LiveEventPreviewOutput {
	return o
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return o.ToLiveEventPreviewPtrOutputWithContext(context.Background())
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventPreview) *LiveEventPreview {
		return &v
	}).(LiveEventPreviewPtrOutput)
}

func (o LiveEventPreviewOutput) AccessControl() LiveEventPreviewAccessControlPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *LiveEventPreviewAccessControl { return v.AccessControl }).(LiveEventPreviewAccessControlPtrOutput)
}

func (o LiveEventPreviewOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *string { return v.AlternativeMediaId }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewOutput) Endpoints() LiveEventEndpointArrayOutput {
	return o.ApplyT(func(v LiveEventPreview) []LiveEventEndpoint { return v.Endpoints }).(LiveEventEndpointArrayOutput)
}

func (o LiveEventPreviewOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *string { return v.PreviewLocator }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *string { return v.StreamingPolicyName }).(pulumi.StringPtrOutput)
}

type LiveEventPreviewPtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreview)(nil)).Elem()
}

func (o LiveEventPreviewPtrOutput) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return o
}

func (o LiveEventPreviewPtrOutput) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return o
}

func (o LiveEventPreviewPtrOutput) Elem() LiveEventPreviewOutput {
	return o.ApplyT(func(v *LiveEventPreview) LiveEventPreview {
		if v != nil {
			return *v
		}
		var ret LiveEventPreview
		return ret
	}).(LiveEventPreviewOutput)
}

func (o LiveEventPreviewPtrOutput) AccessControl() LiveEventPreviewAccessControlPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *LiveEventPreviewAccessControl {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(LiveEventPreviewAccessControlPtrOutput)
}

func (o LiveEventPreviewPtrOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *string {
		if v == nil {
			return nil
		}
		return v.AlternativeMediaId
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewPtrOutput) Endpoints() LiveEventEndpointArrayOutput {
	return o.ApplyT(func(v *LiveEventPreview) []LiveEventEndpoint {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(LiveEventEndpointArrayOutput)
}

func (o LiveEventPreviewPtrOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *string {
		if v == nil {
			return nil
		}
		return v.PreviewLocator
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewPtrOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *string {
		if v == nil {
			return nil
		}
		return v.StreamingPolicyName
	}).(pulumi.StringPtrOutput)
}

type LiveEventPreviewAccessControl struct {
	Ip *IPAccessControl `pulumi:"ip"`
}





type LiveEventPreviewAccessControlInput interface {
	pulumi.Input

	ToLiveEventPreviewAccessControlOutput() LiveEventPreviewAccessControlOutput
	ToLiveEventPreviewAccessControlOutputWithContext(context.Context) LiveEventPreviewAccessControlOutput
}

type LiveEventPreviewAccessControlArgs struct {
	Ip IPAccessControlPtrInput `pulumi:"ip"`
}

func (LiveEventPreviewAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewAccessControl)(nil)).Elem()
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlOutput() LiveEventPreviewAccessControlOutput {
	return i.ToLiveEventPreviewAccessControlOutputWithContext(context.Background())
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlOutput)
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return i.ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Background())
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlOutput).ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx)
}









type LiveEventPreviewAccessControlPtrInput interface {
	pulumi.Input

	ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput
	ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Context) LiveEventPreviewAccessControlPtrOutput
}

type liveEventPreviewAccessControlPtrType LiveEventPreviewAccessControlArgs

func LiveEventPreviewAccessControlPtr(v *LiveEventPreviewAccessControlArgs) LiveEventPreviewAccessControlPtrInput {
	return (*liveEventPreviewAccessControlPtrType)(v)
}

func (*liveEventPreviewAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewAccessControl)(nil)).Elem()
}

func (i *liveEventPreviewAccessControlPtrType) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return i.ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Background())
}

func (i *liveEventPreviewAccessControlPtrType) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlPtrOutput)
}

type LiveEventPreviewAccessControlOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewAccessControl)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlOutput() LiveEventPreviewAccessControlOutput {
	return o
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlOutput {
	return o
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return o.ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Background())
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventPreviewAccessControl) *LiveEventPreviewAccessControl {
		return &v
	}).(LiveEventPreviewAccessControlPtrOutput)
}

func (o LiveEventPreviewAccessControlOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewAccessControl) *IPAccessControl { return v.Ip }).(IPAccessControlPtrOutput)
}

type LiveEventPreviewAccessControlPtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewAccessControl)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlPtrOutput) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlPtrOutput) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlPtrOutput) Elem() LiveEventPreviewAccessControlOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControl) LiveEventPreviewAccessControl {
		if v != nil {
			return *v
		}
		var ret LiveEventPreviewAccessControl
		return ret
	}).(LiveEventPreviewAccessControlOutput)
}

func (o LiveEventPreviewAccessControlPtrOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControl) *IPAccessControl {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlPtrOutput)
}

type LiveEventPreviewAccessControlResponse struct {
	Ip *IPAccessControlResponse `pulumi:"ip"`
}





type LiveEventPreviewAccessControlResponseInput interface {
	pulumi.Input

	ToLiveEventPreviewAccessControlResponseOutput() LiveEventPreviewAccessControlResponseOutput
	ToLiveEventPreviewAccessControlResponseOutputWithContext(context.Context) LiveEventPreviewAccessControlResponseOutput
}

type LiveEventPreviewAccessControlResponseArgs struct {
	Ip IPAccessControlResponsePtrInput `pulumi:"ip"`
}

func (LiveEventPreviewAccessControlResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewAccessControlResponse)(nil)).Elem()
}

func (i LiveEventPreviewAccessControlResponseArgs) ToLiveEventPreviewAccessControlResponseOutput() LiveEventPreviewAccessControlResponseOutput {
	return i.ToLiveEventPreviewAccessControlResponseOutputWithContext(context.Background())
}

func (i LiveEventPreviewAccessControlResponseArgs) ToLiveEventPreviewAccessControlResponseOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlResponseOutput)
}

func (i LiveEventPreviewAccessControlResponseArgs) ToLiveEventPreviewAccessControlResponsePtrOutput() LiveEventPreviewAccessControlResponsePtrOutput {
	return i.ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventPreviewAccessControlResponseArgs) ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlResponseOutput).ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(ctx)
}









type LiveEventPreviewAccessControlResponsePtrInput interface {
	pulumi.Input

	ToLiveEventPreviewAccessControlResponsePtrOutput() LiveEventPreviewAccessControlResponsePtrOutput
	ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(context.Context) LiveEventPreviewAccessControlResponsePtrOutput
}

type liveEventPreviewAccessControlResponsePtrType LiveEventPreviewAccessControlResponseArgs

func LiveEventPreviewAccessControlResponsePtr(v *LiveEventPreviewAccessControlResponseArgs) LiveEventPreviewAccessControlResponsePtrInput {
	return (*liveEventPreviewAccessControlResponsePtrType)(v)
}

func (*liveEventPreviewAccessControlResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewAccessControlResponse)(nil)).Elem()
}

func (i *liveEventPreviewAccessControlResponsePtrType) ToLiveEventPreviewAccessControlResponsePtrOutput() LiveEventPreviewAccessControlResponsePtrOutput {
	return i.ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventPreviewAccessControlResponsePtrType) ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlResponsePtrOutput)
}

type LiveEventPreviewAccessControlResponseOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewAccessControlResponse)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlResponseOutput) ToLiveEventPreviewAccessControlResponseOutput() LiveEventPreviewAccessControlResponseOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponseOutput) ToLiveEventPreviewAccessControlResponseOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponseOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponseOutput) ToLiveEventPreviewAccessControlResponsePtrOutput() LiveEventPreviewAccessControlResponsePtrOutput {
	return o.ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventPreviewAccessControlResponseOutput) ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventPreviewAccessControlResponse) *LiveEventPreviewAccessControlResponse {
		return &v
	}).(LiveEventPreviewAccessControlResponsePtrOutput)
}

func (o LiveEventPreviewAccessControlResponseOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v LiveEventPreviewAccessControlResponse) *IPAccessControlResponse { return v.Ip }).(IPAccessControlResponsePtrOutput)
}

type LiveEventPreviewAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewAccessControlResponse)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) ToLiveEventPreviewAccessControlResponsePtrOutput() LiveEventPreviewAccessControlResponsePtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponsePtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) Elem() LiveEventPreviewAccessControlResponseOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControlResponse) LiveEventPreviewAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventPreviewAccessControlResponse
		return ret
	}).(LiveEventPreviewAccessControlResponseOutput)
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControlResponse) *IPAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlResponsePtrOutput)
}

type LiveEventPreviewResponse struct {
	AccessControl       *LiveEventPreviewAccessControlResponse `pulumi:"accessControl"`
	AlternativeMediaId  *string                                `pulumi:"alternativeMediaId"`
	Endpoints           []LiveEventEndpointResponse            `pulumi:"endpoints"`
	PreviewLocator      *string                                `pulumi:"previewLocator"`
	StreamingPolicyName *string                                `pulumi:"streamingPolicyName"`
}





type LiveEventPreviewResponseInput interface {
	pulumi.Input

	ToLiveEventPreviewResponseOutput() LiveEventPreviewResponseOutput
	ToLiveEventPreviewResponseOutputWithContext(context.Context) LiveEventPreviewResponseOutput
}

type LiveEventPreviewResponseArgs struct {
	AccessControl       LiveEventPreviewAccessControlResponsePtrInput `pulumi:"accessControl"`
	AlternativeMediaId  pulumi.StringPtrInput                         `pulumi:"alternativeMediaId"`
	Endpoints           LiveEventEndpointResponseArrayInput           `pulumi:"endpoints"`
	PreviewLocator      pulumi.StringPtrInput                         `pulumi:"previewLocator"`
	StreamingPolicyName pulumi.StringPtrInput                         `pulumi:"streamingPolicyName"`
}

func (LiveEventPreviewResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewResponse)(nil)).Elem()
}

func (i LiveEventPreviewResponseArgs) ToLiveEventPreviewResponseOutput() LiveEventPreviewResponseOutput {
	return i.ToLiveEventPreviewResponseOutputWithContext(context.Background())
}

func (i LiveEventPreviewResponseArgs) ToLiveEventPreviewResponseOutputWithContext(ctx context.Context) LiveEventPreviewResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewResponseOutput)
}

func (i LiveEventPreviewResponseArgs) ToLiveEventPreviewResponsePtrOutput() LiveEventPreviewResponsePtrOutput {
	return i.ToLiveEventPreviewResponsePtrOutputWithContext(context.Background())
}

func (i LiveEventPreviewResponseArgs) ToLiveEventPreviewResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewResponseOutput).ToLiveEventPreviewResponsePtrOutputWithContext(ctx)
}









type LiveEventPreviewResponsePtrInput interface {
	pulumi.Input

	ToLiveEventPreviewResponsePtrOutput() LiveEventPreviewResponsePtrOutput
	ToLiveEventPreviewResponsePtrOutputWithContext(context.Context) LiveEventPreviewResponsePtrOutput
}

type liveEventPreviewResponsePtrType LiveEventPreviewResponseArgs

func LiveEventPreviewResponsePtr(v *LiveEventPreviewResponseArgs) LiveEventPreviewResponsePtrInput {
	return (*liveEventPreviewResponsePtrType)(v)
}

func (*liveEventPreviewResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewResponse)(nil)).Elem()
}

func (i *liveEventPreviewResponsePtrType) ToLiveEventPreviewResponsePtrOutput() LiveEventPreviewResponsePtrOutput {
	return i.ToLiveEventPreviewResponsePtrOutputWithContext(context.Background())
}

func (i *liveEventPreviewResponsePtrType) ToLiveEventPreviewResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewResponsePtrOutput)
}

type LiveEventPreviewResponseOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewResponse)(nil)).Elem()
}

func (o LiveEventPreviewResponseOutput) ToLiveEventPreviewResponseOutput() LiveEventPreviewResponseOutput {
	return o
}

func (o LiveEventPreviewResponseOutput) ToLiveEventPreviewResponseOutputWithContext(ctx context.Context) LiveEventPreviewResponseOutput {
	return o
}

func (o LiveEventPreviewResponseOutput) ToLiveEventPreviewResponsePtrOutput() LiveEventPreviewResponsePtrOutput {
	return o.ToLiveEventPreviewResponsePtrOutputWithContext(context.Background())
}

func (o LiveEventPreviewResponseOutput) ToLiveEventPreviewResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventPreviewResponse) *LiveEventPreviewResponse {
		return &v
	}).(LiveEventPreviewResponsePtrOutput)
}

func (o LiveEventPreviewResponseOutput) AccessControl() LiveEventPreviewAccessControlResponsePtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *LiveEventPreviewAccessControlResponse { return v.AccessControl }).(LiveEventPreviewAccessControlResponsePtrOutput)
}

func (o LiveEventPreviewResponseOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *string { return v.AlternativeMediaId }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponseOutput) Endpoints() LiveEventEndpointResponseArrayOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) []LiveEventEndpointResponse { return v.Endpoints }).(LiveEventEndpointResponseArrayOutput)
}

func (o LiveEventPreviewResponseOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *string { return v.PreviewLocator }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponseOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *string { return v.StreamingPolicyName }).(pulumi.StringPtrOutput)
}

type LiveEventPreviewResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewResponse)(nil)).Elem()
}

func (o LiveEventPreviewResponsePtrOutput) ToLiveEventPreviewResponsePtrOutput() LiveEventPreviewResponsePtrOutput {
	return o
}

func (o LiveEventPreviewResponsePtrOutput) ToLiveEventPreviewResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewResponsePtrOutput {
	return o
}

func (o LiveEventPreviewResponsePtrOutput) Elem() LiveEventPreviewResponseOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) LiveEventPreviewResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventPreviewResponse
		return ret
	}).(LiveEventPreviewResponseOutput)
}

func (o LiveEventPreviewResponsePtrOutput) AccessControl() LiveEventPreviewAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *LiveEventPreviewAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(LiveEventPreviewAccessControlResponsePtrOutput)
}

func (o LiveEventPreviewResponsePtrOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *string {
		if v == nil {
			return nil
		}
		return v.AlternativeMediaId
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponsePtrOutput) Endpoints() LiveEventEndpointResponseArrayOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) []LiveEventEndpointResponse {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(LiveEventEndpointResponseArrayOutput)
}

func (o LiveEventPreviewResponsePtrOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreviewLocator
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponsePtrOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *string {
		if v == nil {
			return nil
		}
		return v.StreamingPolicyName
	}).(pulumi.StringPtrOutput)
}

type LiveEventTranscription struct {
	InputTrackSelection      []LiveEventInputTrackSelection     `pulumi:"inputTrackSelection"`
	Language                 *string                            `pulumi:"language"`
	OutputTranscriptionTrack *LiveEventOutputTranscriptionTrack `pulumi:"outputTranscriptionTrack"`
}





type LiveEventTranscriptionInput interface {
	pulumi.Input

	ToLiveEventTranscriptionOutput() LiveEventTranscriptionOutput
	ToLiveEventTranscriptionOutputWithContext(context.Context) LiveEventTranscriptionOutput
}

type LiveEventTranscriptionArgs struct {
	InputTrackSelection      LiveEventInputTrackSelectionArrayInput    `pulumi:"inputTrackSelection"`
	Language                 pulumi.StringPtrInput                     `pulumi:"language"`
	OutputTranscriptionTrack LiveEventOutputTranscriptionTrackPtrInput `pulumi:"outputTranscriptionTrack"`
}

func (LiveEventTranscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventTranscription)(nil)).Elem()
}

func (i LiveEventTranscriptionArgs) ToLiveEventTranscriptionOutput() LiveEventTranscriptionOutput {
	return i.ToLiveEventTranscriptionOutputWithContext(context.Background())
}

func (i LiveEventTranscriptionArgs) ToLiveEventTranscriptionOutputWithContext(ctx context.Context) LiveEventTranscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventTranscriptionOutput)
}





type LiveEventTranscriptionArrayInput interface {
	pulumi.Input

	ToLiveEventTranscriptionArrayOutput() LiveEventTranscriptionArrayOutput
	ToLiveEventTranscriptionArrayOutputWithContext(context.Context) LiveEventTranscriptionArrayOutput
}

type LiveEventTranscriptionArray []LiveEventTranscriptionInput

func (LiveEventTranscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventTranscription)(nil)).Elem()
}

func (i LiveEventTranscriptionArray) ToLiveEventTranscriptionArrayOutput() LiveEventTranscriptionArrayOutput {
	return i.ToLiveEventTranscriptionArrayOutputWithContext(context.Background())
}

func (i LiveEventTranscriptionArray) ToLiveEventTranscriptionArrayOutputWithContext(ctx context.Context) LiveEventTranscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventTranscriptionArrayOutput)
}

type LiveEventTranscriptionOutput struct{ *pulumi.OutputState }

func (LiveEventTranscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventTranscription)(nil)).Elem()
}

func (o LiveEventTranscriptionOutput) ToLiveEventTranscriptionOutput() LiveEventTranscriptionOutput {
	return o
}

func (o LiveEventTranscriptionOutput) ToLiveEventTranscriptionOutputWithContext(ctx context.Context) LiveEventTranscriptionOutput {
	return o
}

func (o LiveEventTranscriptionOutput) InputTrackSelection() LiveEventInputTrackSelectionArrayOutput {
	return o.ApplyT(func(v LiveEventTranscription) []LiveEventInputTrackSelection { return v.InputTrackSelection }).(LiveEventInputTrackSelectionArrayOutput)
}

func (o LiveEventTranscriptionOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventTranscription) *string { return v.Language }).(pulumi.StringPtrOutput)
}

func (o LiveEventTranscriptionOutput) OutputTranscriptionTrack() LiveEventOutputTranscriptionTrackPtrOutput {
	return o.ApplyT(func(v LiveEventTranscription) *LiveEventOutputTranscriptionTrack { return v.OutputTranscriptionTrack }).(LiveEventOutputTranscriptionTrackPtrOutput)
}

type LiveEventTranscriptionArrayOutput struct{ *pulumi.OutputState }

func (LiveEventTranscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventTranscription)(nil)).Elem()
}

func (o LiveEventTranscriptionArrayOutput) ToLiveEventTranscriptionArrayOutput() LiveEventTranscriptionArrayOutput {
	return o
}

func (o LiveEventTranscriptionArrayOutput) ToLiveEventTranscriptionArrayOutputWithContext(ctx context.Context) LiveEventTranscriptionArrayOutput {
	return o
}

func (o LiveEventTranscriptionArrayOutput) Index(i pulumi.IntInput) LiveEventTranscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventTranscription {
		return vs[0].([]LiveEventTranscription)[vs[1].(int)]
	}).(LiveEventTranscriptionOutput)
}

type LiveEventTranscriptionResponse struct {
	InputTrackSelection      []LiveEventInputTrackSelectionResponse     `pulumi:"inputTrackSelection"`
	Language                 *string                                    `pulumi:"language"`
	OutputTranscriptionTrack *LiveEventOutputTranscriptionTrackResponse `pulumi:"outputTranscriptionTrack"`
}





type LiveEventTranscriptionResponseInput interface {
	pulumi.Input

	ToLiveEventTranscriptionResponseOutput() LiveEventTranscriptionResponseOutput
	ToLiveEventTranscriptionResponseOutputWithContext(context.Context) LiveEventTranscriptionResponseOutput
}

type LiveEventTranscriptionResponseArgs struct {
	InputTrackSelection      LiveEventInputTrackSelectionResponseArrayInput    `pulumi:"inputTrackSelection"`
	Language                 pulumi.StringPtrInput                             `pulumi:"language"`
	OutputTranscriptionTrack LiveEventOutputTranscriptionTrackResponsePtrInput `pulumi:"outputTranscriptionTrack"`
}

func (LiveEventTranscriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventTranscriptionResponse)(nil)).Elem()
}

func (i LiveEventTranscriptionResponseArgs) ToLiveEventTranscriptionResponseOutput() LiveEventTranscriptionResponseOutput {
	return i.ToLiveEventTranscriptionResponseOutputWithContext(context.Background())
}

func (i LiveEventTranscriptionResponseArgs) ToLiveEventTranscriptionResponseOutputWithContext(ctx context.Context) LiveEventTranscriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventTranscriptionResponseOutput)
}





type LiveEventTranscriptionResponseArrayInput interface {
	pulumi.Input

	ToLiveEventTranscriptionResponseArrayOutput() LiveEventTranscriptionResponseArrayOutput
	ToLiveEventTranscriptionResponseArrayOutputWithContext(context.Context) LiveEventTranscriptionResponseArrayOutput
}

type LiveEventTranscriptionResponseArray []LiveEventTranscriptionResponseInput

func (LiveEventTranscriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventTranscriptionResponse)(nil)).Elem()
}

func (i LiveEventTranscriptionResponseArray) ToLiveEventTranscriptionResponseArrayOutput() LiveEventTranscriptionResponseArrayOutput {
	return i.ToLiveEventTranscriptionResponseArrayOutputWithContext(context.Background())
}

func (i LiveEventTranscriptionResponseArray) ToLiveEventTranscriptionResponseArrayOutputWithContext(ctx context.Context) LiveEventTranscriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventTranscriptionResponseArrayOutput)
}

type LiveEventTranscriptionResponseOutput struct{ *pulumi.OutputState }

func (LiveEventTranscriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventTranscriptionResponse)(nil)).Elem()
}

func (o LiveEventTranscriptionResponseOutput) ToLiveEventTranscriptionResponseOutput() LiveEventTranscriptionResponseOutput {
	return o
}

func (o LiveEventTranscriptionResponseOutput) ToLiveEventTranscriptionResponseOutputWithContext(ctx context.Context) LiveEventTranscriptionResponseOutput {
	return o
}

func (o LiveEventTranscriptionResponseOutput) InputTrackSelection() LiveEventInputTrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v LiveEventTranscriptionResponse) []LiveEventInputTrackSelectionResponse {
		return v.InputTrackSelection
	}).(LiveEventInputTrackSelectionResponseArrayOutput)
}

func (o LiveEventTranscriptionResponseOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventTranscriptionResponse) *string { return v.Language }).(pulumi.StringPtrOutput)
}

func (o LiveEventTranscriptionResponseOutput) OutputTranscriptionTrack() LiveEventOutputTranscriptionTrackResponsePtrOutput {
	return o.ApplyT(func(v LiveEventTranscriptionResponse) *LiveEventOutputTranscriptionTrackResponse {
		return v.OutputTranscriptionTrack
	}).(LiveEventOutputTranscriptionTrackResponsePtrOutput)
}

type LiveEventTranscriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (LiveEventTranscriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventTranscriptionResponse)(nil)).Elem()
}

func (o LiveEventTranscriptionResponseArrayOutput) ToLiveEventTranscriptionResponseArrayOutput() LiveEventTranscriptionResponseArrayOutput {
	return o
}

func (o LiveEventTranscriptionResponseArrayOutput) ToLiveEventTranscriptionResponseArrayOutputWithContext(ctx context.Context) LiveEventTranscriptionResponseArrayOutput {
	return o
}

func (o LiveEventTranscriptionResponseArrayOutput) Index(i pulumi.IntInput) LiveEventTranscriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventTranscriptionResponse {
		return vs[0].([]LiveEventTranscriptionResponse)[vs[1].(int)]
	}).(LiveEventTranscriptionResponseOutput)
}

type MediaServiceIdentity struct {
	Type string `pulumi:"type"`
}





type MediaServiceIdentityInput interface {
	pulumi.Input

	ToMediaServiceIdentityOutput() MediaServiceIdentityOutput
	ToMediaServiceIdentityOutputWithContext(context.Context) MediaServiceIdentityOutput
}

type MediaServiceIdentityArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (MediaServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServiceIdentity)(nil)).Elem()
}

func (i MediaServiceIdentityArgs) ToMediaServiceIdentityOutput() MediaServiceIdentityOutput {
	return i.ToMediaServiceIdentityOutputWithContext(context.Background())
}

func (i MediaServiceIdentityArgs) ToMediaServiceIdentityOutputWithContext(ctx context.Context) MediaServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceIdentityOutput)
}

func (i MediaServiceIdentityArgs) ToMediaServiceIdentityPtrOutput() MediaServiceIdentityPtrOutput {
	return i.ToMediaServiceIdentityPtrOutputWithContext(context.Background())
}

func (i MediaServiceIdentityArgs) ToMediaServiceIdentityPtrOutputWithContext(ctx context.Context) MediaServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceIdentityOutput).ToMediaServiceIdentityPtrOutputWithContext(ctx)
}









type MediaServiceIdentityPtrInput interface {
	pulumi.Input

	ToMediaServiceIdentityPtrOutput() MediaServiceIdentityPtrOutput
	ToMediaServiceIdentityPtrOutputWithContext(context.Context) MediaServiceIdentityPtrOutput
}

type mediaServiceIdentityPtrType MediaServiceIdentityArgs

func MediaServiceIdentityPtr(v *MediaServiceIdentityArgs) MediaServiceIdentityPtrInput {
	return (*mediaServiceIdentityPtrType)(v)
}

func (*mediaServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServiceIdentity)(nil)).Elem()
}

func (i *mediaServiceIdentityPtrType) ToMediaServiceIdentityPtrOutput() MediaServiceIdentityPtrOutput {
	return i.ToMediaServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *mediaServiceIdentityPtrType) ToMediaServiceIdentityPtrOutputWithContext(ctx context.Context) MediaServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceIdentityPtrOutput)
}

type MediaServiceIdentityOutput struct{ *pulumi.OutputState }

func (MediaServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServiceIdentity)(nil)).Elem()
}

func (o MediaServiceIdentityOutput) ToMediaServiceIdentityOutput() MediaServiceIdentityOutput {
	return o
}

func (o MediaServiceIdentityOutput) ToMediaServiceIdentityOutputWithContext(ctx context.Context) MediaServiceIdentityOutput {
	return o
}

func (o MediaServiceIdentityOutput) ToMediaServiceIdentityPtrOutput() MediaServiceIdentityPtrOutput {
	return o.ToMediaServiceIdentityPtrOutputWithContext(context.Background())
}

func (o MediaServiceIdentityOutput) ToMediaServiceIdentityPtrOutputWithContext(ctx context.Context) MediaServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaServiceIdentity) *MediaServiceIdentity {
		return &v
	}).(MediaServiceIdentityPtrOutput)
}

func (o MediaServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MediaServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type MediaServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (MediaServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServiceIdentity)(nil)).Elem()
}

func (o MediaServiceIdentityPtrOutput) ToMediaServiceIdentityPtrOutput() MediaServiceIdentityPtrOutput {
	return o
}

func (o MediaServiceIdentityPtrOutput) ToMediaServiceIdentityPtrOutputWithContext(ctx context.Context) MediaServiceIdentityPtrOutput {
	return o
}

func (o MediaServiceIdentityPtrOutput) Elem() MediaServiceIdentityOutput {
	return o.ApplyT(func(v *MediaServiceIdentity) MediaServiceIdentity {
		if v != nil {
			return *v
		}
		var ret MediaServiceIdentity
		return ret
	}).(MediaServiceIdentityOutput)
}

func (o MediaServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type MediaServiceIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}





type MediaServiceIdentityResponseInput interface {
	pulumi.Input

	ToMediaServiceIdentityResponseOutput() MediaServiceIdentityResponseOutput
	ToMediaServiceIdentityResponseOutputWithContext(context.Context) MediaServiceIdentityResponseOutput
}

type MediaServiceIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (MediaServiceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServiceIdentityResponse)(nil)).Elem()
}

func (i MediaServiceIdentityResponseArgs) ToMediaServiceIdentityResponseOutput() MediaServiceIdentityResponseOutput {
	return i.ToMediaServiceIdentityResponseOutputWithContext(context.Background())
}

func (i MediaServiceIdentityResponseArgs) ToMediaServiceIdentityResponseOutputWithContext(ctx context.Context) MediaServiceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceIdentityResponseOutput)
}

func (i MediaServiceIdentityResponseArgs) ToMediaServiceIdentityResponsePtrOutput() MediaServiceIdentityResponsePtrOutput {
	return i.ToMediaServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i MediaServiceIdentityResponseArgs) ToMediaServiceIdentityResponsePtrOutputWithContext(ctx context.Context) MediaServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceIdentityResponseOutput).ToMediaServiceIdentityResponsePtrOutputWithContext(ctx)
}









type MediaServiceIdentityResponsePtrInput interface {
	pulumi.Input

	ToMediaServiceIdentityResponsePtrOutput() MediaServiceIdentityResponsePtrOutput
	ToMediaServiceIdentityResponsePtrOutputWithContext(context.Context) MediaServiceIdentityResponsePtrOutput
}

type mediaServiceIdentityResponsePtrType MediaServiceIdentityResponseArgs

func MediaServiceIdentityResponsePtr(v *MediaServiceIdentityResponseArgs) MediaServiceIdentityResponsePtrInput {
	return (*mediaServiceIdentityResponsePtrType)(v)
}

func (*mediaServiceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServiceIdentityResponse)(nil)).Elem()
}

func (i *mediaServiceIdentityResponsePtrType) ToMediaServiceIdentityResponsePtrOutput() MediaServiceIdentityResponsePtrOutput {
	return i.ToMediaServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *mediaServiceIdentityResponsePtrType) ToMediaServiceIdentityResponsePtrOutputWithContext(ctx context.Context) MediaServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceIdentityResponsePtrOutput)
}

type MediaServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (MediaServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServiceIdentityResponse)(nil)).Elem()
}

func (o MediaServiceIdentityResponseOutput) ToMediaServiceIdentityResponseOutput() MediaServiceIdentityResponseOutput {
	return o
}

func (o MediaServiceIdentityResponseOutput) ToMediaServiceIdentityResponseOutputWithContext(ctx context.Context) MediaServiceIdentityResponseOutput {
	return o
}

func (o MediaServiceIdentityResponseOutput) ToMediaServiceIdentityResponsePtrOutput() MediaServiceIdentityResponsePtrOutput {
	return o.ToMediaServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o MediaServiceIdentityResponseOutput) ToMediaServiceIdentityResponsePtrOutputWithContext(ctx context.Context) MediaServiceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaServiceIdentityResponse) *MediaServiceIdentityResponse {
		return &v
	}).(MediaServiceIdentityResponsePtrOutput)
}

func (o MediaServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v MediaServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o MediaServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v MediaServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o MediaServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MediaServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type MediaServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (MediaServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServiceIdentityResponse)(nil)).Elem()
}

func (o MediaServiceIdentityResponsePtrOutput) ToMediaServiceIdentityResponsePtrOutput() MediaServiceIdentityResponsePtrOutput {
	return o
}

func (o MediaServiceIdentityResponsePtrOutput) ToMediaServiceIdentityResponsePtrOutputWithContext(ctx context.Context) MediaServiceIdentityResponsePtrOutput {
	return o
}

func (o MediaServiceIdentityResponsePtrOutput) Elem() MediaServiceIdentityResponseOutput {
	return o.ApplyT(func(v *MediaServiceIdentityResponse) MediaServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret MediaServiceIdentityResponse
		return ret
	}).(MediaServiceIdentityResponseOutput)
}

func (o MediaServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o MediaServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o MediaServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type Mp4Format struct {
	FilenamePattern string       `pulumi:"filenamePattern"`
	OdataType       string       `pulumi:"odataType"`
	OutputFiles     []OutputFile `pulumi:"outputFiles"`
}





type Mp4FormatInput interface {
	pulumi.Input

	ToMp4FormatOutput() Mp4FormatOutput
	ToMp4FormatOutputWithContext(context.Context) Mp4FormatOutput
}

type Mp4FormatArgs struct {
	FilenamePattern pulumi.StringInput   `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput   `pulumi:"odataType"`
	OutputFiles     OutputFileArrayInput `pulumi:"outputFiles"`
}

func (Mp4FormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Mp4Format)(nil)).Elem()
}

func (i Mp4FormatArgs) ToMp4FormatOutput() Mp4FormatOutput {
	return i.ToMp4FormatOutputWithContext(context.Background())
}

func (i Mp4FormatArgs) ToMp4FormatOutputWithContext(ctx context.Context) Mp4FormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Mp4FormatOutput)
}

type Mp4FormatOutput struct{ *pulumi.OutputState }

func (Mp4FormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Mp4Format)(nil)).Elem()
}

func (o Mp4FormatOutput) ToMp4FormatOutput() Mp4FormatOutput {
	return o
}

func (o Mp4FormatOutput) ToMp4FormatOutputWithContext(ctx context.Context) Mp4FormatOutput {
	return o
}

func (o Mp4FormatOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v Mp4Format) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o Mp4FormatOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v Mp4Format) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o Mp4FormatOutput) OutputFiles() OutputFileArrayOutput {
	return o.ApplyT(func(v Mp4Format) []OutputFile { return v.OutputFiles }).(OutputFileArrayOutput)
}

type Mp4FormatResponse struct {
	FilenamePattern string               `pulumi:"filenamePattern"`
	OdataType       string               `pulumi:"odataType"`
	OutputFiles     []OutputFileResponse `pulumi:"outputFiles"`
}





type Mp4FormatResponseInput interface {
	pulumi.Input

	ToMp4FormatResponseOutput() Mp4FormatResponseOutput
	ToMp4FormatResponseOutputWithContext(context.Context) Mp4FormatResponseOutput
}

type Mp4FormatResponseArgs struct {
	FilenamePattern pulumi.StringInput           `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput           `pulumi:"odataType"`
	OutputFiles     OutputFileResponseArrayInput `pulumi:"outputFiles"`
}

func (Mp4FormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Mp4FormatResponse)(nil)).Elem()
}

func (i Mp4FormatResponseArgs) ToMp4FormatResponseOutput() Mp4FormatResponseOutput {
	return i.ToMp4FormatResponseOutputWithContext(context.Background())
}

func (i Mp4FormatResponseArgs) ToMp4FormatResponseOutputWithContext(ctx context.Context) Mp4FormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Mp4FormatResponseOutput)
}

type Mp4FormatResponseOutput struct{ *pulumi.OutputState }

func (Mp4FormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Mp4FormatResponse)(nil)).Elem()
}

func (o Mp4FormatResponseOutput) ToMp4FormatResponseOutput() Mp4FormatResponseOutput {
	return o
}

func (o Mp4FormatResponseOutput) ToMp4FormatResponseOutputWithContext(ctx context.Context) Mp4FormatResponseOutput {
	return o
}

func (o Mp4FormatResponseOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v Mp4FormatResponse) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o Mp4FormatResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v Mp4FormatResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o Mp4FormatResponseOutput) OutputFiles() OutputFileResponseArrayOutput {
	return o.ApplyT(func(v Mp4FormatResponse) []OutputFileResponse { return v.OutputFiles }).(OutputFileResponseArrayOutput)
}

type MultiBitrateFormat struct {
	FilenamePattern string       `pulumi:"filenamePattern"`
	OdataType       string       `pulumi:"odataType"`
	OutputFiles     []OutputFile `pulumi:"outputFiles"`
}





type MultiBitrateFormatInput interface {
	pulumi.Input

	ToMultiBitrateFormatOutput() MultiBitrateFormatOutput
	ToMultiBitrateFormatOutputWithContext(context.Context) MultiBitrateFormatOutput
}

type MultiBitrateFormatArgs struct {
	FilenamePattern pulumi.StringInput   `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput   `pulumi:"odataType"`
	OutputFiles     OutputFileArrayInput `pulumi:"outputFiles"`
}

func (MultiBitrateFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MultiBitrateFormat)(nil)).Elem()
}

func (i MultiBitrateFormatArgs) ToMultiBitrateFormatOutput() MultiBitrateFormatOutput {
	return i.ToMultiBitrateFormatOutputWithContext(context.Background())
}

func (i MultiBitrateFormatArgs) ToMultiBitrateFormatOutputWithContext(ctx context.Context) MultiBitrateFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiBitrateFormatOutput)
}

type MultiBitrateFormatOutput struct{ *pulumi.OutputState }

func (MultiBitrateFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MultiBitrateFormat)(nil)).Elem()
}

func (o MultiBitrateFormatOutput) ToMultiBitrateFormatOutput() MultiBitrateFormatOutput {
	return o
}

func (o MultiBitrateFormatOutput) ToMultiBitrateFormatOutputWithContext(ctx context.Context) MultiBitrateFormatOutput {
	return o
}

func (o MultiBitrateFormatOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v MultiBitrateFormat) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o MultiBitrateFormatOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MultiBitrateFormat) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MultiBitrateFormatOutput) OutputFiles() OutputFileArrayOutput {
	return o.ApplyT(func(v MultiBitrateFormat) []OutputFile { return v.OutputFiles }).(OutputFileArrayOutput)
}

type MultiBitrateFormatResponse struct {
	FilenamePattern string               `pulumi:"filenamePattern"`
	OdataType       string               `pulumi:"odataType"`
	OutputFiles     []OutputFileResponse `pulumi:"outputFiles"`
}





type MultiBitrateFormatResponseInput interface {
	pulumi.Input

	ToMultiBitrateFormatResponseOutput() MultiBitrateFormatResponseOutput
	ToMultiBitrateFormatResponseOutputWithContext(context.Context) MultiBitrateFormatResponseOutput
}

type MultiBitrateFormatResponseArgs struct {
	FilenamePattern pulumi.StringInput           `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput           `pulumi:"odataType"`
	OutputFiles     OutputFileResponseArrayInput `pulumi:"outputFiles"`
}

func (MultiBitrateFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MultiBitrateFormatResponse)(nil)).Elem()
}

func (i MultiBitrateFormatResponseArgs) ToMultiBitrateFormatResponseOutput() MultiBitrateFormatResponseOutput {
	return i.ToMultiBitrateFormatResponseOutputWithContext(context.Background())
}

func (i MultiBitrateFormatResponseArgs) ToMultiBitrateFormatResponseOutputWithContext(ctx context.Context) MultiBitrateFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiBitrateFormatResponseOutput)
}

type MultiBitrateFormatResponseOutput struct{ *pulumi.OutputState }

func (MultiBitrateFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MultiBitrateFormatResponse)(nil)).Elem()
}

func (o MultiBitrateFormatResponseOutput) ToMultiBitrateFormatResponseOutput() MultiBitrateFormatResponseOutput {
	return o
}

func (o MultiBitrateFormatResponseOutput) ToMultiBitrateFormatResponseOutputWithContext(ctx context.Context) MultiBitrateFormatResponseOutput {
	return o
}

func (o MultiBitrateFormatResponseOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v MultiBitrateFormatResponse) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o MultiBitrateFormatResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MultiBitrateFormatResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o MultiBitrateFormatResponseOutput) OutputFiles() OutputFileResponseArrayOutput {
	return o.ApplyT(func(v MultiBitrateFormatResponse) []OutputFileResponse { return v.OutputFiles }).(OutputFileResponseArrayOutput)
}

type NoEncryption struct {
	EnabledProtocols *EnabledProtocols `pulumi:"enabledProtocols"`
}





type NoEncryptionInput interface {
	pulumi.Input

	ToNoEncryptionOutput() NoEncryptionOutput
	ToNoEncryptionOutputWithContext(context.Context) NoEncryptionOutput
}

type NoEncryptionArgs struct {
	EnabledProtocols EnabledProtocolsPtrInput `pulumi:"enabledProtocols"`
}

func (NoEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NoEncryption)(nil)).Elem()
}

func (i NoEncryptionArgs) ToNoEncryptionOutput() NoEncryptionOutput {
	return i.ToNoEncryptionOutputWithContext(context.Background())
}

func (i NoEncryptionArgs) ToNoEncryptionOutputWithContext(ctx context.Context) NoEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoEncryptionOutput)
}

func (i NoEncryptionArgs) ToNoEncryptionPtrOutput() NoEncryptionPtrOutput {
	return i.ToNoEncryptionPtrOutputWithContext(context.Background())
}

func (i NoEncryptionArgs) ToNoEncryptionPtrOutputWithContext(ctx context.Context) NoEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoEncryptionOutput).ToNoEncryptionPtrOutputWithContext(ctx)
}









type NoEncryptionPtrInput interface {
	pulumi.Input

	ToNoEncryptionPtrOutput() NoEncryptionPtrOutput
	ToNoEncryptionPtrOutputWithContext(context.Context) NoEncryptionPtrOutput
}

type noEncryptionPtrType NoEncryptionArgs

func NoEncryptionPtr(v *NoEncryptionArgs) NoEncryptionPtrInput {
	return (*noEncryptionPtrType)(v)
}

func (*noEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NoEncryption)(nil)).Elem()
}

func (i *noEncryptionPtrType) ToNoEncryptionPtrOutput() NoEncryptionPtrOutput {
	return i.ToNoEncryptionPtrOutputWithContext(context.Background())
}

func (i *noEncryptionPtrType) ToNoEncryptionPtrOutputWithContext(ctx context.Context) NoEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoEncryptionPtrOutput)
}

type NoEncryptionOutput struct{ *pulumi.OutputState }

func (NoEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NoEncryption)(nil)).Elem()
}

func (o NoEncryptionOutput) ToNoEncryptionOutput() NoEncryptionOutput {
	return o
}

func (o NoEncryptionOutput) ToNoEncryptionOutputWithContext(ctx context.Context) NoEncryptionOutput {
	return o
}

func (o NoEncryptionOutput) ToNoEncryptionPtrOutput() NoEncryptionPtrOutput {
	return o.ToNoEncryptionPtrOutputWithContext(context.Background())
}

func (o NoEncryptionOutput) ToNoEncryptionPtrOutputWithContext(ctx context.Context) NoEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NoEncryption) *NoEncryption {
		return &v
	}).(NoEncryptionPtrOutput)
}

func (o NoEncryptionOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v NoEncryption) *EnabledProtocols { return v.EnabledProtocols }).(EnabledProtocolsPtrOutput)
}

type NoEncryptionPtrOutput struct{ *pulumi.OutputState }

func (NoEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NoEncryption)(nil)).Elem()
}

func (o NoEncryptionPtrOutput) ToNoEncryptionPtrOutput() NoEncryptionPtrOutput {
	return o
}

func (o NoEncryptionPtrOutput) ToNoEncryptionPtrOutputWithContext(ctx context.Context) NoEncryptionPtrOutput {
	return o
}

func (o NoEncryptionPtrOutput) Elem() NoEncryptionOutput {
	return o.ApplyT(func(v *NoEncryption) NoEncryption {
		if v != nil {
			return *v
		}
		var ret NoEncryption
		return ret
	}).(NoEncryptionOutput)
}

func (o NoEncryptionPtrOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v *NoEncryption) *EnabledProtocols {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsPtrOutput)
}

type NoEncryptionResponse struct {
	EnabledProtocols *EnabledProtocolsResponse `pulumi:"enabledProtocols"`
}





type NoEncryptionResponseInput interface {
	pulumi.Input

	ToNoEncryptionResponseOutput() NoEncryptionResponseOutput
	ToNoEncryptionResponseOutputWithContext(context.Context) NoEncryptionResponseOutput
}

type NoEncryptionResponseArgs struct {
	EnabledProtocols EnabledProtocolsResponsePtrInput `pulumi:"enabledProtocols"`
}

func (NoEncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NoEncryptionResponse)(nil)).Elem()
}

func (i NoEncryptionResponseArgs) ToNoEncryptionResponseOutput() NoEncryptionResponseOutput {
	return i.ToNoEncryptionResponseOutputWithContext(context.Background())
}

func (i NoEncryptionResponseArgs) ToNoEncryptionResponseOutputWithContext(ctx context.Context) NoEncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoEncryptionResponseOutput)
}

func (i NoEncryptionResponseArgs) ToNoEncryptionResponsePtrOutput() NoEncryptionResponsePtrOutput {
	return i.ToNoEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i NoEncryptionResponseArgs) ToNoEncryptionResponsePtrOutputWithContext(ctx context.Context) NoEncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoEncryptionResponseOutput).ToNoEncryptionResponsePtrOutputWithContext(ctx)
}









type NoEncryptionResponsePtrInput interface {
	pulumi.Input

	ToNoEncryptionResponsePtrOutput() NoEncryptionResponsePtrOutput
	ToNoEncryptionResponsePtrOutputWithContext(context.Context) NoEncryptionResponsePtrOutput
}

type noEncryptionResponsePtrType NoEncryptionResponseArgs

func NoEncryptionResponsePtr(v *NoEncryptionResponseArgs) NoEncryptionResponsePtrInput {
	return (*noEncryptionResponsePtrType)(v)
}

func (*noEncryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NoEncryptionResponse)(nil)).Elem()
}

func (i *noEncryptionResponsePtrType) ToNoEncryptionResponsePtrOutput() NoEncryptionResponsePtrOutput {
	return i.ToNoEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *noEncryptionResponsePtrType) ToNoEncryptionResponsePtrOutputWithContext(ctx context.Context) NoEncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoEncryptionResponsePtrOutput)
}

type NoEncryptionResponseOutput struct{ *pulumi.OutputState }

func (NoEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NoEncryptionResponse)(nil)).Elem()
}

func (o NoEncryptionResponseOutput) ToNoEncryptionResponseOutput() NoEncryptionResponseOutput {
	return o
}

func (o NoEncryptionResponseOutput) ToNoEncryptionResponseOutputWithContext(ctx context.Context) NoEncryptionResponseOutput {
	return o
}

func (o NoEncryptionResponseOutput) ToNoEncryptionResponsePtrOutput() NoEncryptionResponsePtrOutput {
	return o.ToNoEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o NoEncryptionResponseOutput) ToNoEncryptionResponsePtrOutputWithContext(ctx context.Context) NoEncryptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NoEncryptionResponse) *NoEncryptionResponse {
		return &v
	}).(NoEncryptionResponsePtrOutput)
}

func (o NoEncryptionResponseOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v NoEncryptionResponse) *EnabledProtocolsResponse { return v.EnabledProtocols }).(EnabledProtocolsResponsePtrOutput)
}

type NoEncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (NoEncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NoEncryptionResponse)(nil)).Elem()
}

func (o NoEncryptionResponsePtrOutput) ToNoEncryptionResponsePtrOutput() NoEncryptionResponsePtrOutput {
	return o
}

func (o NoEncryptionResponsePtrOutput) ToNoEncryptionResponsePtrOutputWithContext(ctx context.Context) NoEncryptionResponsePtrOutput {
	return o
}

func (o NoEncryptionResponsePtrOutput) Elem() NoEncryptionResponseOutput {
	return o.ApplyT(func(v *NoEncryptionResponse) NoEncryptionResponse {
		if v != nil {
			return *v
		}
		var ret NoEncryptionResponse
		return ret
	}).(NoEncryptionResponseOutput)
}

func (o NoEncryptionResponsePtrOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v *NoEncryptionResponse) *EnabledProtocolsResponse {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsResponsePtrOutput)
}

type OutputFile struct {
	Labels []string `pulumi:"labels"`
}





type OutputFileInput interface {
	pulumi.Input

	ToOutputFileOutput() OutputFileOutput
	ToOutputFileOutputWithContext(context.Context) OutputFileOutput
}

type OutputFileArgs struct {
	Labels pulumi.StringArrayInput `pulumi:"labels"`
}

func (OutputFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputFile)(nil)).Elem()
}

func (i OutputFileArgs) ToOutputFileOutput() OutputFileOutput {
	return i.ToOutputFileOutputWithContext(context.Background())
}

func (i OutputFileArgs) ToOutputFileOutputWithContext(ctx context.Context) OutputFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputFileOutput)
}





type OutputFileArrayInput interface {
	pulumi.Input

	ToOutputFileArrayOutput() OutputFileArrayOutput
	ToOutputFileArrayOutputWithContext(context.Context) OutputFileArrayOutput
}

type OutputFileArray []OutputFileInput

func (OutputFileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputFile)(nil)).Elem()
}

func (i OutputFileArray) ToOutputFileArrayOutput() OutputFileArrayOutput {
	return i.ToOutputFileArrayOutputWithContext(context.Background())
}

func (i OutputFileArray) ToOutputFileArrayOutputWithContext(ctx context.Context) OutputFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputFileArrayOutput)
}

type OutputFileOutput struct{ *pulumi.OutputState }

func (OutputFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputFile)(nil)).Elem()
}

func (o OutputFileOutput) ToOutputFileOutput() OutputFileOutput {
	return o
}

func (o OutputFileOutput) ToOutputFileOutputWithContext(ctx context.Context) OutputFileOutput {
	return o
}

func (o OutputFileOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OutputFile) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

type OutputFileArrayOutput struct{ *pulumi.OutputState }

func (OutputFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputFile)(nil)).Elem()
}

func (o OutputFileArrayOutput) ToOutputFileArrayOutput() OutputFileArrayOutput {
	return o
}

func (o OutputFileArrayOutput) ToOutputFileArrayOutputWithContext(ctx context.Context) OutputFileArrayOutput {
	return o
}

func (o OutputFileArrayOutput) Index(i pulumi.IntInput) OutputFileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputFile {
		return vs[0].([]OutputFile)[vs[1].(int)]
	}).(OutputFileOutput)
}

type OutputFileResponse struct {
	Labels []string `pulumi:"labels"`
}





type OutputFileResponseInput interface {
	pulumi.Input

	ToOutputFileResponseOutput() OutputFileResponseOutput
	ToOutputFileResponseOutputWithContext(context.Context) OutputFileResponseOutput
}

type OutputFileResponseArgs struct {
	Labels pulumi.StringArrayInput `pulumi:"labels"`
}

func (OutputFileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputFileResponse)(nil)).Elem()
}

func (i OutputFileResponseArgs) ToOutputFileResponseOutput() OutputFileResponseOutput {
	return i.ToOutputFileResponseOutputWithContext(context.Background())
}

func (i OutputFileResponseArgs) ToOutputFileResponseOutputWithContext(ctx context.Context) OutputFileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputFileResponseOutput)
}





type OutputFileResponseArrayInput interface {
	pulumi.Input

	ToOutputFileResponseArrayOutput() OutputFileResponseArrayOutput
	ToOutputFileResponseArrayOutputWithContext(context.Context) OutputFileResponseArrayOutput
}

type OutputFileResponseArray []OutputFileResponseInput

func (OutputFileResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputFileResponse)(nil)).Elem()
}

func (i OutputFileResponseArray) ToOutputFileResponseArrayOutput() OutputFileResponseArrayOutput {
	return i.ToOutputFileResponseArrayOutputWithContext(context.Background())
}

func (i OutputFileResponseArray) ToOutputFileResponseArrayOutputWithContext(ctx context.Context) OutputFileResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputFileResponseArrayOutput)
}

type OutputFileResponseOutput struct{ *pulumi.OutputState }

func (OutputFileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputFileResponse)(nil)).Elem()
}

func (o OutputFileResponseOutput) ToOutputFileResponseOutput() OutputFileResponseOutput {
	return o
}

func (o OutputFileResponseOutput) ToOutputFileResponseOutputWithContext(ctx context.Context) OutputFileResponseOutput {
	return o
}

func (o OutputFileResponseOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OutputFileResponse) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

type OutputFileResponseArrayOutput struct{ *pulumi.OutputState }

func (OutputFileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputFileResponse)(nil)).Elem()
}

func (o OutputFileResponseArrayOutput) ToOutputFileResponseArrayOutput() OutputFileResponseArrayOutput {
	return o
}

func (o OutputFileResponseArrayOutput) ToOutputFileResponseArrayOutputWithContext(ctx context.Context) OutputFileResponseArrayOutput {
	return o
}

func (o OutputFileResponseArrayOutput) Index(i pulumi.IntInput) OutputFileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputFileResponse {
		return vs[0].([]OutputFileResponse)[vs[1].(int)]
	}).(OutputFileResponseOutput)
}

type PngFormat struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}





type PngFormatInput interface {
	pulumi.Input

	ToPngFormatOutput() PngFormatOutput
	ToPngFormatOutputWithContext(context.Context) PngFormatOutput
}

type PngFormatArgs struct {
	FilenamePattern pulumi.StringInput `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput `pulumi:"odataType"`
}

func (PngFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PngFormat)(nil)).Elem()
}

func (i PngFormatArgs) ToPngFormatOutput() PngFormatOutput {
	return i.ToPngFormatOutputWithContext(context.Background())
}

func (i PngFormatArgs) ToPngFormatOutputWithContext(ctx context.Context) PngFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PngFormatOutput)
}

type PngFormatOutput struct{ *pulumi.OutputState }

func (PngFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PngFormat)(nil)).Elem()
}

func (o PngFormatOutput) ToPngFormatOutput() PngFormatOutput {
	return o
}

func (o PngFormatOutput) ToPngFormatOutputWithContext(ctx context.Context) PngFormatOutput {
	return o
}

func (o PngFormatOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v PngFormat) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o PngFormatOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v PngFormat) string { return v.OdataType }).(pulumi.StringOutput)
}

type PngFormatResponse struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}





type PngFormatResponseInput interface {
	pulumi.Input

	ToPngFormatResponseOutput() PngFormatResponseOutput
	ToPngFormatResponseOutputWithContext(context.Context) PngFormatResponseOutput
}

type PngFormatResponseArgs struct {
	FilenamePattern pulumi.StringInput `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput `pulumi:"odataType"`
}

func (PngFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PngFormatResponse)(nil)).Elem()
}

func (i PngFormatResponseArgs) ToPngFormatResponseOutput() PngFormatResponseOutput {
	return i.ToPngFormatResponseOutputWithContext(context.Background())
}

func (i PngFormatResponseArgs) ToPngFormatResponseOutputWithContext(ctx context.Context) PngFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PngFormatResponseOutput)
}

type PngFormatResponseOutput struct{ *pulumi.OutputState }

func (PngFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PngFormatResponse)(nil)).Elem()
}

func (o PngFormatResponseOutput) ToPngFormatResponseOutput() PngFormatResponseOutput {
	return o
}

func (o PngFormatResponseOutput) ToPngFormatResponseOutputWithContext(ctx context.Context) PngFormatResponseOutput {
	return o
}

func (o PngFormatResponseOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v PngFormatResponse) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o PngFormatResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v PngFormatResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type PngImage struct {
	KeyFrameInterval *string    `pulumi:"keyFrameInterval"`
	Label            *string    `pulumi:"label"`
	Layers           []PngLayer `pulumi:"layers"`
	OdataType        string     `pulumi:"odataType"`
	Range            *string    `pulumi:"range"`
	Start            string     `pulumi:"start"`
	Step             *string    `pulumi:"step"`
	StretchMode      *string    `pulumi:"stretchMode"`
	SyncMode         *string    `pulumi:"syncMode"`
}





type PngImageInput interface {
	pulumi.Input

	ToPngImageOutput() PngImageOutput
	ToPngImageOutputWithContext(context.Context) PngImageOutput
}

type PngImageArgs struct {
	KeyFrameInterval pulumi.StringPtrInput `pulumi:"keyFrameInterval"`
	Label            pulumi.StringPtrInput `pulumi:"label"`
	Layers           PngLayerArrayInput    `pulumi:"layers"`
	OdataType        pulumi.StringInput    `pulumi:"odataType"`
	Range            pulumi.StringPtrInput `pulumi:"range"`
	Start            pulumi.StringInput    `pulumi:"start"`
	Step             pulumi.StringPtrInput `pulumi:"step"`
	StretchMode      pulumi.StringPtrInput `pulumi:"stretchMode"`
	SyncMode         pulumi.StringPtrInput `pulumi:"syncMode"`
}

func (PngImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PngImage)(nil)).Elem()
}

func (i PngImageArgs) ToPngImageOutput() PngImageOutput {
	return i.ToPngImageOutputWithContext(context.Background())
}

func (i PngImageArgs) ToPngImageOutputWithContext(ctx context.Context) PngImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PngImageOutput)
}

type PngImageOutput struct{ *pulumi.OutputState }

func (PngImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PngImage)(nil)).Elem()
}

func (o PngImageOutput) ToPngImageOutput() PngImageOutput {
	return o
}

func (o PngImageOutput) ToPngImageOutputWithContext(ctx context.Context) PngImageOutput {
	return o
}

func (o PngImageOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImage) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o PngImageOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImage) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o PngImageOutput) Layers() PngLayerArrayOutput {
	return o.ApplyT(func(v PngImage) []PngLayer { return v.Layers }).(PngLayerArrayOutput)
}

func (o PngImageOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v PngImage) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o PngImageOutput) Range() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImage) *string { return v.Range }).(pulumi.StringPtrOutput)
}

func (o PngImageOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v PngImage) string { return v.Start }).(pulumi.StringOutput)
}

func (o PngImageOutput) Step() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImage) *string { return v.Step }).(pulumi.StringPtrOutput)
}

func (o PngImageOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImage) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o PngImageOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImage) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type PngImageResponse struct {
	KeyFrameInterval *string            `pulumi:"keyFrameInterval"`
	Label            *string            `pulumi:"label"`
	Layers           []PngLayerResponse `pulumi:"layers"`
	OdataType        string             `pulumi:"odataType"`
	Range            *string            `pulumi:"range"`
	Start            string             `pulumi:"start"`
	Step             *string            `pulumi:"step"`
	StretchMode      *string            `pulumi:"stretchMode"`
	SyncMode         *string            `pulumi:"syncMode"`
}





type PngImageResponseInput interface {
	pulumi.Input

	ToPngImageResponseOutput() PngImageResponseOutput
	ToPngImageResponseOutputWithContext(context.Context) PngImageResponseOutput
}

type PngImageResponseArgs struct {
	KeyFrameInterval pulumi.StringPtrInput      `pulumi:"keyFrameInterval"`
	Label            pulumi.StringPtrInput      `pulumi:"label"`
	Layers           PngLayerResponseArrayInput `pulumi:"layers"`
	OdataType        pulumi.StringInput         `pulumi:"odataType"`
	Range            pulumi.StringPtrInput      `pulumi:"range"`
	Start            pulumi.StringInput         `pulumi:"start"`
	Step             pulumi.StringPtrInput      `pulumi:"step"`
	StretchMode      pulumi.StringPtrInput      `pulumi:"stretchMode"`
	SyncMode         pulumi.StringPtrInput      `pulumi:"syncMode"`
}

func (PngImageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PngImageResponse)(nil)).Elem()
}

func (i PngImageResponseArgs) ToPngImageResponseOutput() PngImageResponseOutput {
	return i.ToPngImageResponseOutputWithContext(context.Background())
}

func (i PngImageResponseArgs) ToPngImageResponseOutputWithContext(ctx context.Context) PngImageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PngImageResponseOutput)
}

type PngImageResponseOutput struct{ *pulumi.OutputState }

func (PngImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PngImageResponse)(nil)).Elem()
}

func (o PngImageResponseOutput) ToPngImageResponseOutput() PngImageResponseOutput {
	return o
}

func (o PngImageResponseOutput) ToPngImageResponseOutputWithContext(ctx context.Context) PngImageResponseOutput {
	return o
}

func (o PngImageResponseOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImageResponse) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o PngImageResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImageResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o PngImageResponseOutput) Layers() PngLayerResponseArrayOutput {
	return o.ApplyT(func(v PngImageResponse) []PngLayerResponse { return v.Layers }).(PngLayerResponseArrayOutput)
}

func (o PngImageResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v PngImageResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o PngImageResponseOutput) Range() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImageResponse) *string { return v.Range }).(pulumi.StringPtrOutput)
}

func (o PngImageResponseOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v PngImageResponse) string { return v.Start }).(pulumi.StringOutput)
}

func (o PngImageResponseOutput) Step() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImageResponse) *string { return v.Step }).(pulumi.StringPtrOutput)
}

func (o PngImageResponseOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImageResponse) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o PngImageResponseOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngImageResponse) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type PngLayer struct {
	Height    *string `pulumi:"height"`
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
	Width     *string `pulumi:"width"`
}





type PngLayerInput interface {
	pulumi.Input

	ToPngLayerOutput() PngLayerOutput
	ToPngLayerOutputWithContext(context.Context) PngLayerOutput
}

type PngLayerArgs struct {
	Height    pulumi.StringPtrInput `pulumi:"height"`
	Label     pulumi.StringPtrInput `pulumi:"label"`
	OdataType pulumi.StringInput    `pulumi:"odataType"`
	Width     pulumi.StringPtrInput `pulumi:"width"`
}

func (PngLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PngLayer)(nil)).Elem()
}

func (i PngLayerArgs) ToPngLayerOutput() PngLayerOutput {
	return i.ToPngLayerOutputWithContext(context.Background())
}

func (i PngLayerArgs) ToPngLayerOutputWithContext(ctx context.Context) PngLayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PngLayerOutput)
}





type PngLayerArrayInput interface {
	pulumi.Input

	ToPngLayerArrayOutput() PngLayerArrayOutput
	ToPngLayerArrayOutputWithContext(context.Context) PngLayerArrayOutput
}

type PngLayerArray []PngLayerInput

func (PngLayerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PngLayer)(nil)).Elem()
}

func (i PngLayerArray) ToPngLayerArrayOutput() PngLayerArrayOutput {
	return i.ToPngLayerArrayOutputWithContext(context.Background())
}

func (i PngLayerArray) ToPngLayerArrayOutputWithContext(ctx context.Context) PngLayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PngLayerArrayOutput)
}

type PngLayerOutput struct{ *pulumi.OutputState }

func (PngLayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PngLayer)(nil)).Elem()
}

func (o PngLayerOutput) ToPngLayerOutput() PngLayerOutput {
	return o
}

func (o PngLayerOutput) ToPngLayerOutputWithContext(ctx context.Context) PngLayerOutput {
	return o
}

func (o PngLayerOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngLayer) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o PngLayerOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngLayer) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o PngLayerOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v PngLayer) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o PngLayerOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngLayer) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type PngLayerArrayOutput struct{ *pulumi.OutputState }

func (PngLayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PngLayer)(nil)).Elem()
}

func (o PngLayerArrayOutput) ToPngLayerArrayOutput() PngLayerArrayOutput {
	return o
}

func (o PngLayerArrayOutput) ToPngLayerArrayOutputWithContext(ctx context.Context) PngLayerArrayOutput {
	return o
}

func (o PngLayerArrayOutput) Index(i pulumi.IntInput) PngLayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PngLayer {
		return vs[0].([]PngLayer)[vs[1].(int)]
	}).(PngLayerOutput)
}

type PngLayerResponse struct {
	Height    *string `pulumi:"height"`
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
	Width     *string `pulumi:"width"`
}





type PngLayerResponseInput interface {
	pulumi.Input

	ToPngLayerResponseOutput() PngLayerResponseOutput
	ToPngLayerResponseOutputWithContext(context.Context) PngLayerResponseOutput
}

type PngLayerResponseArgs struct {
	Height    pulumi.StringPtrInput `pulumi:"height"`
	Label     pulumi.StringPtrInput `pulumi:"label"`
	OdataType pulumi.StringInput    `pulumi:"odataType"`
	Width     pulumi.StringPtrInput `pulumi:"width"`
}

func (PngLayerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PngLayerResponse)(nil)).Elem()
}

func (i PngLayerResponseArgs) ToPngLayerResponseOutput() PngLayerResponseOutput {
	return i.ToPngLayerResponseOutputWithContext(context.Background())
}

func (i PngLayerResponseArgs) ToPngLayerResponseOutputWithContext(ctx context.Context) PngLayerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PngLayerResponseOutput)
}





type PngLayerResponseArrayInput interface {
	pulumi.Input

	ToPngLayerResponseArrayOutput() PngLayerResponseArrayOutput
	ToPngLayerResponseArrayOutputWithContext(context.Context) PngLayerResponseArrayOutput
}

type PngLayerResponseArray []PngLayerResponseInput

func (PngLayerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PngLayerResponse)(nil)).Elem()
}

func (i PngLayerResponseArray) ToPngLayerResponseArrayOutput() PngLayerResponseArrayOutput {
	return i.ToPngLayerResponseArrayOutputWithContext(context.Background())
}

func (i PngLayerResponseArray) ToPngLayerResponseArrayOutputWithContext(ctx context.Context) PngLayerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PngLayerResponseArrayOutput)
}

type PngLayerResponseOutput struct{ *pulumi.OutputState }

func (PngLayerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PngLayerResponse)(nil)).Elem()
}

func (o PngLayerResponseOutput) ToPngLayerResponseOutput() PngLayerResponseOutput {
	return o
}

func (o PngLayerResponseOutput) ToPngLayerResponseOutputWithContext(ctx context.Context) PngLayerResponseOutput {
	return o
}

func (o PngLayerResponseOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngLayerResponse) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o PngLayerResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngLayerResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o PngLayerResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v PngLayerResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o PngLayerResponseOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PngLayerResponse) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type PngLayerResponseArrayOutput struct{ *pulumi.OutputState }

func (PngLayerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PngLayerResponse)(nil)).Elem()
}

func (o PngLayerResponseArrayOutput) ToPngLayerResponseArrayOutput() PngLayerResponseArrayOutput {
	return o
}

func (o PngLayerResponseArrayOutput) ToPngLayerResponseArrayOutputWithContext(ctx context.Context) PngLayerResponseArrayOutput {
	return o
}

func (o PngLayerResponseArrayOutput) Index(i pulumi.IntInput) PngLayerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PngLayerResponse {
		return vs[0].([]PngLayerResponse)[vs[1].(int)]
	}).(PngLayerResponseOutput)
}

type PresentationTimeRange struct {
	EndTimestamp               *float64 `pulumi:"endTimestamp"`
	ForceEndTimestamp          *bool    `pulumi:"forceEndTimestamp"`
	LiveBackoffDuration        *float64 `pulumi:"liveBackoffDuration"`
	PresentationWindowDuration *float64 `pulumi:"presentationWindowDuration"`
	StartTimestamp             *float64 `pulumi:"startTimestamp"`
	Timescale                  *float64 `pulumi:"timescale"`
}





type PresentationTimeRangeInput interface {
	pulumi.Input

	ToPresentationTimeRangeOutput() PresentationTimeRangeOutput
	ToPresentationTimeRangeOutputWithContext(context.Context) PresentationTimeRangeOutput
}

type PresentationTimeRangeArgs struct {
	EndTimestamp               pulumi.Float64PtrInput `pulumi:"endTimestamp"`
	ForceEndTimestamp          pulumi.BoolPtrInput    `pulumi:"forceEndTimestamp"`
	LiveBackoffDuration        pulumi.Float64PtrInput `pulumi:"liveBackoffDuration"`
	PresentationWindowDuration pulumi.Float64PtrInput `pulumi:"presentationWindowDuration"`
	StartTimestamp             pulumi.Float64PtrInput `pulumi:"startTimestamp"`
	Timescale                  pulumi.Float64PtrInput `pulumi:"timescale"`
}

func (PresentationTimeRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PresentationTimeRange)(nil)).Elem()
}

func (i PresentationTimeRangeArgs) ToPresentationTimeRangeOutput() PresentationTimeRangeOutput {
	return i.ToPresentationTimeRangeOutputWithContext(context.Background())
}

func (i PresentationTimeRangeArgs) ToPresentationTimeRangeOutputWithContext(ctx context.Context) PresentationTimeRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresentationTimeRangeOutput)
}

func (i PresentationTimeRangeArgs) ToPresentationTimeRangePtrOutput() PresentationTimeRangePtrOutput {
	return i.ToPresentationTimeRangePtrOutputWithContext(context.Background())
}

func (i PresentationTimeRangeArgs) ToPresentationTimeRangePtrOutputWithContext(ctx context.Context) PresentationTimeRangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresentationTimeRangeOutput).ToPresentationTimeRangePtrOutputWithContext(ctx)
}









type PresentationTimeRangePtrInput interface {
	pulumi.Input

	ToPresentationTimeRangePtrOutput() PresentationTimeRangePtrOutput
	ToPresentationTimeRangePtrOutputWithContext(context.Context) PresentationTimeRangePtrOutput
}

type presentationTimeRangePtrType PresentationTimeRangeArgs

func PresentationTimeRangePtr(v *PresentationTimeRangeArgs) PresentationTimeRangePtrInput {
	return (*presentationTimeRangePtrType)(v)
}

func (*presentationTimeRangePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PresentationTimeRange)(nil)).Elem()
}

func (i *presentationTimeRangePtrType) ToPresentationTimeRangePtrOutput() PresentationTimeRangePtrOutput {
	return i.ToPresentationTimeRangePtrOutputWithContext(context.Background())
}

func (i *presentationTimeRangePtrType) ToPresentationTimeRangePtrOutputWithContext(ctx context.Context) PresentationTimeRangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresentationTimeRangePtrOutput)
}

type PresentationTimeRangeOutput struct{ *pulumi.OutputState }

func (PresentationTimeRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PresentationTimeRange)(nil)).Elem()
}

func (o PresentationTimeRangeOutput) ToPresentationTimeRangeOutput() PresentationTimeRangeOutput {
	return o
}

func (o PresentationTimeRangeOutput) ToPresentationTimeRangeOutputWithContext(ctx context.Context) PresentationTimeRangeOutput {
	return o
}

func (o PresentationTimeRangeOutput) ToPresentationTimeRangePtrOutput() PresentationTimeRangePtrOutput {
	return o.ToPresentationTimeRangePtrOutputWithContext(context.Background())
}

func (o PresentationTimeRangeOutput) ToPresentationTimeRangePtrOutputWithContext(ctx context.Context) PresentationTimeRangePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PresentationTimeRange) *PresentationTimeRange {
		return &v
	}).(PresentationTimeRangePtrOutput)
}

func (o PresentationTimeRangeOutput) EndTimestamp() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v PresentationTimeRange) *float64 { return v.EndTimestamp }).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeOutput) ForceEndTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PresentationTimeRange) *bool { return v.ForceEndTimestamp }).(pulumi.BoolPtrOutput)
}

func (o PresentationTimeRangeOutput) LiveBackoffDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v PresentationTimeRange) *float64 { return v.LiveBackoffDuration }).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeOutput) PresentationWindowDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v PresentationTimeRange) *float64 { return v.PresentationWindowDuration }).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeOutput) StartTimestamp() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v PresentationTimeRange) *float64 { return v.StartTimestamp }).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeOutput) Timescale() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v PresentationTimeRange) *float64 { return v.Timescale }).(pulumi.Float64PtrOutput)
}

type PresentationTimeRangePtrOutput struct{ *pulumi.OutputState }

func (PresentationTimeRangePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PresentationTimeRange)(nil)).Elem()
}

func (o PresentationTimeRangePtrOutput) ToPresentationTimeRangePtrOutput() PresentationTimeRangePtrOutput {
	return o
}

func (o PresentationTimeRangePtrOutput) ToPresentationTimeRangePtrOutputWithContext(ctx context.Context) PresentationTimeRangePtrOutput {
	return o
}

func (o PresentationTimeRangePtrOutput) Elem() PresentationTimeRangeOutput {
	return o.ApplyT(func(v *PresentationTimeRange) PresentationTimeRange {
		if v != nil {
			return *v
		}
		var ret PresentationTimeRange
		return ret
	}).(PresentationTimeRangeOutput)
}

func (o PresentationTimeRangePtrOutput) EndTimestamp() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PresentationTimeRange) *float64 {
		if v == nil {
			return nil
		}
		return v.EndTimestamp
	}).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangePtrOutput) ForceEndTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PresentationTimeRange) *bool {
		if v == nil {
			return nil
		}
		return v.ForceEndTimestamp
	}).(pulumi.BoolPtrOutput)
}

func (o PresentationTimeRangePtrOutput) LiveBackoffDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PresentationTimeRange) *float64 {
		if v == nil {
			return nil
		}
		return v.LiveBackoffDuration
	}).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangePtrOutput) PresentationWindowDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PresentationTimeRange) *float64 {
		if v == nil {
			return nil
		}
		return v.PresentationWindowDuration
	}).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangePtrOutput) StartTimestamp() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PresentationTimeRange) *float64 {
		if v == nil {
			return nil
		}
		return v.StartTimestamp
	}).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangePtrOutput) Timescale() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PresentationTimeRange) *float64 {
		if v == nil {
			return nil
		}
		return v.Timescale
	}).(pulumi.Float64PtrOutput)
}

type PresentationTimeRangeResponse struct {
	EndTimestamp               *float64 `pulumi:"endTimestamp"`
	ForceEndTimestamp          *bool    `pulumi:"forceEndTimestamp"`
	LiveBackoffDuration        *float64 `pulumi:"liveBackoffDuration"`
	PresentationWindowDuration *float64 `pulumi:"presentationWindowDuration"`
	StartTimestamp             *float64 `pulumi:"startTimestamp"`
	Timescale                  *float64 `pulumi:"timescale"`
}





type PresentationTimeRangeResponseInput interface {
	pulumi.Input

	ToPresentationTimeRangeResponseOutput() PresentationTimeRangeResponseOutput
	ToPresentationTimeRangeResponseOutputWithContext(context.Context) PresentationTimeRangeResponseOutput
}

type PresentationTimeRangeResponseArgs struct {
	EndTimestamp               pulumi.Float64PtrInput `pulumi:"endTimestamp"`
	ForceEndTimestamp          pulumi.BoolPtrInput    `pulumi:"forceEndTimestamp"`
	LiveBackoffDuration        pulumi.Float64PtrInput `pulumi:"liveBackoffDuration"`
	PresentationWindowDuration pulumi.Float64PtrInput `pulumi:"presentationWindowDuration"`
	StartTimestamp             pulumi.Float64PtrInput `pulumi:"startTimestamp"`
	Timescale                  pulumi.Float64PtrInput `pulumi:"timescale"`
}

func (PresentationTimeRangeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PresentationTimeRangeResponse)(nil)).Elem()
}

func (i PresentationTimeRangeResponseArgs) ToPresentationTimeRangeResponseOutput() PresentationTimeRangeResponseOutput {
	return i.ToPresentationTimeRangeResponseOutputWithContext(context.Background())
}

func (i PresentationTimeRangeResponseArgs) ToPresentationTimeRangeResponseOutputWithContext(ctx context.Context) PresentationTimeRangeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresentationTimeRangeResponseOutput)
}

func (i PresentationTimeRangeResponseArgs) ToPresentationTimeRangeResponsePtrOutput() PresentationTimeRangeResponsePtrOutput {
	return i.ToPresentationTimeRangeResponsePtrOutputWithContext(context.Background())
}

func (i PresentationTimeRangeResponseArgs) ToPresentationTimeRangeResponsePtrOutputWithContext(ctx context.Context) PresentationTimeRangeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresentationTimeRangeResponseOutput).ToPresentationTimeRangeResponsePtrOutputWithContext(ctx)
}









type PresentationTimeRangeResponsePtrInput interface {
	pulumi.Input

	ToPresentationTimeRangeResponsePtrOutput() PresentationTimeRangeResponsePtrOutput
	ToPresentationTimeRangeResponsePtrOutputWithContext(context.Context) PresentationTimeRangeResponsePtrOutput
}

type presentationTimeRangeResponsePtrType PresentationTimeRangeResponseArgs

func PresentationTimeRangeResponsePtr(v *PresentationTimeRangeResponseArgs) PresentationTimeRangeResponsePtrInput {
	return (*presentationTimeRangeResponsePtrType)(v)
}

func (*presentationTimeRangeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PresentationTimeRangeResponse)(nil)).Elem()
}

func (i *presentationTimeRangeResponsePtrType) ToPresentationTimeRangeResponsePtrOutput() PresentationTimeRangeResponsePtrOutput {
	return i.ToPresentationTimeRangeResponsePtrOutputWithContext(context.Background())
}

func (i *presentationTimeRangeResponsePtrType) ToPresentationTimeRangeResponsePtrOutputWithContext(ctx context.Context) PresentationTimeRangeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresentationTimeRangeResponsePtrOutput)
}

type PresentationTimeRangeResponseOutput struct{ *pulumi.OutputState }

func (PresentationTimeRangeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PresentationTimeRangeResponse)(nil)).Elem()
}

func (o PresentationTimeRangeResponseOutput) ToPresentationTimeRangeResponseOutput() PresentationTimeRangeResponseOutput {
	return o
}

func (o PresentationTimeRangeResponseOutput) ToPresentationTimeRangeResponseOutputWithContext(ctx context.Context) PresentationTimeRangeResponseOutput {
	return o
}

func (o PresentationTimeRangeResponseOutput) ToPresentationTimeRangeResponsePtrOutput() PresentationTimeRangeResponsePtrOutput {
	return o.ToPresentationTimeRangeResponsePtrOutputWithContext(context.Background())
}

func (o PresentationTimeRangeResponseOutput) ToPresentationTimeRangeResponsePtrOutputWithContext(ctx context.Context) PresentationTimeRangeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PresentationTimeRangeResponse) *PresentationTimeRangeResponse {
		return &v
	}).(PresentationTimeRangeResponsePtrOutput)
}

func (o PresentationTimeRangeResponseOutput) EndTimestamp() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v PresentationTimeRangeResponse) *float64 { return v.EndTimestamp }).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeResponseOutput) ForceEndTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PresentationTimeRangeResponse) *bool { return v.ForceEndTimestamp }).(pulumi.BoolPtrOutput)
}

func (o PresentationTimeRangeResponseOutput) LiveBackoffDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v PresentationTimeRangeResponse) *float64 { return v.LiveBackoffDuration }).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeResponseOutput) PresentationWindowDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v PresentationTimeRangeResponse) *float64 { return v.PresentationWindowDuration }).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeResponseOutput) StartTimestamp() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v PresentationTimeRangeResponse) *float64 { return v.StartTimestamp }).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeResponseOutput) Timescale() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v PresentationTimeRangeResponse) *float64 { return v.Timescale }).(pulumi.Float64PtrOutput)
}

type PresentationTimeRangeResponsePtrOutput struct{ *pulumi.OutputState }

func (PresentationTimeRangeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PresentationTimeRangeResponse)(nil)).Elem()
}

func (o PresentationTimeRangeResponsePtrOutput) ToPresentationTimeRangeResponsePtrOutput() PresentationTimeRangeResponsePtrOutput {
	return o
}

func (o PresentationTimeRangeResponsePtrOutput) ToPresentationTimeRangeResponsePtrOutputWithContext(ctx context.Context) PresentationTimeRangeResponsePtrOutput {
	return o
}

func (o PresentationTimeRangeResponsePtrOutput) Elem() PresentationTimeRangeResponseOutput {
	return o.ApplyT(func(v *PresentationTimeRangeResponse) PresentationTimeRangeResponse {
		if v != nil {
			return *v
		}
		var ret PresentationTimeRangeResponse
		return ret
	}).(PresentationTimeRangeResponseOutput)
}

func (o PresentationTimeRangeResponsePtrOutput) EndTimestamp() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PresentationTimeRangeResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.EndTimestamp
	}).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeResponsePtrOutput) ForceEndTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PresentationTimeRangeResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ForceEndTimestamp
	}).(pulumi.BoolPtrOutput)
}

func (o PresentationTimeRangeResponsePtrOutput) LiveBackoffDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PresentationTimeRangeResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.LiveBackoffDuration
	}).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeResponsePtrOutput) PresentationWindowDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PresentationTimeRangeResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.PresentationWindowDuration
	}).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeResponsePtrOutput) StartTimestamp() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PresentationTimeRangeResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.StartTimestamp
	}).(pulumi.Float64PtrOutput)
}

func (o PresentationTimeRangeResponsePtrOutput) Timescale() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PresentationTimeRangeResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Timescale
	}).(pulumi.Float64PtrOutput)
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

type Rectangle struct {
	Height *string `pulumi:"height"`
	Left   *string `pulumi:"left"`
	Top    *string `pulumi:"top"`
	Width  *string `pulumi:"width"`
}





type RectangleInput interface {
	pulumi.Input

	ToRectangleOutput() RectangleOutput
	ToRectangleOutputWithContext(context.Context) RectangleOutput
}

type RectangleArgs struct {
	Height pulumi.StringPtrInput `pulumi:"height"`
	Left   pulumi.StringPtrInput `pulumi:"left"`
	Top    pulumi.StringPtrInput `pulumi:"top"`
	Width  pulumi.StringPtrInput `pulumi:"width"`
}

func (RectangleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Rectangle)(nil)).Elem()
}

func (i RectangleArgs) ToRectangleOutput() RectangleOutput {
	return i.ToRectangleOutputWithContext(context.Background())
}

func (i RectangleArgs) ToRectangleOutputWithContext(ctx context.Context) RectangleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RectangleOutput)
}

func (i RectangleArgs) ToRectanglePtrOutput() RectanglePtrOutput {
	return i.ToRectanglePtrOutputWithContext(context.Background())
}

func (i RectangleArgs) ToRectanglePtrOutputWithContext(ctx context.Context) RectanglePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RectangleOutput).ToRectanglePtrOutputWithContext(ctx)
}









type RectanglePtrInput interface {
	pulumi.Input

	ToRectanglePtrOutput() RectanglePtrOutput
	ToRectanglePtrOutputWithContext(context.Context) RectanglePtrOutput
}

type rectanglePtrType RectangleArgs

func RectanglePtr(v *RectangleArgs) RectanglePtrInput {
	return (*rectanglePtrType)(v)
}

func (*rectanglePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Rectangle)(nil)).Elem()
}

func (i *rectanglePtrType) ToRectanglePtrOutput() RectanglePtrOutput {
	return i.ToRectanglePtrOutputWithContext(context.Background())
}

func (i *rectanglePtrType) ToRectanglePtrOutputWithContext(ctx context.Context) RectanglePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RectanglePtrOutput)
}

type RectangleOutput struct{ *pulumi.OutputState }

func (RectangleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rectangle)(nil)).Elem()
}

func (o RectangleOutput) ToRectangleOutput() RectangleOutput {
	return o
}

func (o RectangleOutput) ToRectangleOutputWithContext(ctx context.Context) RectangleOutput {
	return o
}

func (o RectangleOutput) ToRectanglePtrOutput() RectanglePtrOutput {
	return o.ToRectanglePtrOutputWithContext(context.Background())
}

func (o RectangleOutput) ToRectanglePtrOutputWithContext(ctx context.Context) RectanglePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Rectangle) *Rectangle {
		return &v
	}).(RectanglePtrOutput)
}

func (o RectangleOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Rectangle) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o RectangleOutput) Left() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Rectangle) *string { return v.Left }).(pulumi.StringPtrOutput)
}

func (o RectangleOutput) Top() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Rectangle) *string { return v.Top }).(pulumi.StringPtrOutput)
}

func (o RectangleOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Rectangle) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type RectanglePtrOutput struct{ *pulumi.OutputState }

func (RectanglePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rectangle)(nil)).Elem()
}

func (o RectanglePtrOutput) ToRectanglePtrOutput() RectanglePtrOutput {
	return o
}

func (o RectanglePtrOutput) ToRectanglePtrOutputWithContext(ctx context.Context) RectanglePtrOutput {
	return o
}

func (o RectanglePtrOutput) Elem() RectangleOutput {
	return o.ApplyT(func(v *Rectangle) Rectangle {
		if v != nil {
			return *v
		}
		var ret Rectangle
		return ret
	}).(RectangleOutput)
}

func (o RectanglePtrOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rectangle) *string {
		if v == nil {
			return nil
		}
		return v.Height
	}).(pulumi.StringPtrOutput)
}

func (o RectanglePtrOutput) Left() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rectangle) *string {
		if v == nil {
			return nil
		}
		return v.Left
	}).(pulumi.StringPtrOutput)
}

func (o RectanglePtrOutput) Top() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rectangle) *string {
		if v == nil {
			return nil
		}
		return v.Top
	}).(pulumi.StringPtrOutput)
}

func (o RectanglePtrOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rectangle) *string {
		if v == nil {
			return nil
		}
		return v.Width
	}).(pulumi.StringPtrOutput)
}

type RectangleResponse struct {
	Height *string `pulumi:"height"`
	Left   *string `pulumi:"left"`
	Top    *string `pulumi:"top"`
	Width  *string `pulumi:"width"`
}





type RectangleResponseInput interface {
	pulumi.Input

	ToRectangleResponseOutput() RectangleResponseOutput
	ToRectangleResponseOutputWithContext(context.Context) RectangleResponseOutput
}

type RectangleResponseArgs struct {
	Height pulumi.StringPtrInput `pulumi:"height"`
	Left   pulumi.StringPtrInput `pulumi:"left"`
	Top    pulumi.StringPtrInput `pulumi:"top"`
	Width  pulumi.StringPtrInput `pulumi:"width"`
}

func (RectangleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RectangleResponse)(nil)).Elem()
}

func (i RectangleResponseArgs) ToRectangleResponseOutput() RectangleResponseOutput {
	return i.ToRectangleResponseOutputWithContext(context.Background())
}

func (i RectangleResponseArgs) ToRectangleResponseOutputWithContext(ctx context.Context) RectangleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RectangleResponseOutput)
}

func (i RectangleResponseArgs) ToRectangleResponsePtrOutput() RectangleResponsePtrOutput {
	return i.ToRectangleResponsePtrOutputWithContext(context.Background())
}

func (i RectangleResponseArgs) ToRectangleResponsePtrOutputWithContext(ctx context.Context) RectangleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RectangleResponseOutput).ToRectangleResponsePtrOutputWithContext(ctx)
}









type RectangleResponsePtrInput interface {
	pulumi.Input

	ToRectangleResponsePtrOutput() RectangleResponsePtrOutput
	ToRectangleResponsePtrOutputWithContext(context.Context) RectangleResponsePtrOutput
}

type rectangleResponsePtrType RectangleResponseArgs

func RectangleResponsePtr(v *RectangleResponseArgs) RectangleResponsePtrInput {
	return (*rectangleResponsePtrType)(v)
}

func (*rectangleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RectangleResponse)(nil)).Elem()
}

func (i *rectangleResponsePtrType) ToRectangleResponsePtrOutput() RectangleResponsePtrOutput {
	return i.ToRectangleResponsePtrOutputWithContext(context.Background())
}

func (i *rectangleResponsePtrType) ToRectangleResponsePtrOutputWithContext(ctx context.Context) RectangleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RectangleResponsePtrOutput)
}

type RectangleResponseOutput struct{ *pulumi.OutputState }

func (RectangleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RectangleResponse)(nil)).Elem()
}

func (o RectangleResponseOutput) ToRectangleResponseOutput() RectangleResponseOutput {
	return o
}

func (o RectangleResponseOutput) ToRectangleResponseOutputWithContext(ctx context.Context) RectangleResponseOutput {
	return o
}

func (o RectangleResponseOutput) ToRectangleResponsePtrOutput() RectangleResponsePtrOutput {
	return o.ToRectangleResponsePtrOutputWithContext(context.Background())
}

func (o RectangleResponseOutput) ToRectangleResponsePtrOutputWithContext(ctx context.Context) RectangleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RectangleResponse) *RectangleResponse {
		return &v
	}).(RectangleResponsePtrOutput)
}

func (o RectangleResponseOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RectangleResponse) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o RectangleResponseOutput) Left() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RectangleResponse) *string { return v.Left }).(pulumi.StringPtrOutput)
}

func (o RectangleResponseOutput) Top() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RectangleResponse) *string { return v.Top }).(pulumi.StringPtrOutput)
}

func (o RectangleResponseOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RectangleResponse) *string { return v.Width }).(pulumi.StringPtrOutput)
}

type RectangleResponsePtrOutput struct{ *pulumi.OutputState }

func (RectangleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RectangleResponse)(nil)).Elem()
}

func (o RectangleResponsePtrOutput) ToRectangleResponsePtrOutput() RectangleResponsePtrOutput {
	return o
}

func (o RectangleResponsePtrOutput) ToRectangleResponsePtrOutputWithContext(ctx context.Context) RectangleResponsePtrOutput {
	return o
}

func (o RectangleResponsePtrOutput) Elem() RectangleResponseOutput {
	return o.ApplyT(func(v *RectangleResponse) RectangleResponse {
		if v != nil {
			return *v
		}
		var ret RectangleResponse
		return ret
	}).(RectangleResponseOutput)
}

func (o RectangleResponsePtrOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RectangleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Height
	}).(pulumi.StringPtrOutput)
}

func (o RectangleResponsePtrOutput) Left() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RectangleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Left
	}).(pulumi.StringPtrOutput)
}

func (o RectangleResponsePtrOutput) Top() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RectangleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Top
	}).(pulumi.StringPtrOutput)
}

func (o RectangleResponsePtrOutput) Width() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RectangleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Width
	}).(pulumi.StringPtrOutput)
}

type SelectAudioTrackByAttribute struct {
	Attribute      string  `pulumi:"attribute"`
	ChannelMapping *string `pulumi:"channelMapping"`
	Filter         string  `pulumi:"filter"`
	FilterValue    *string `pulumi:"filterValue"`
	OdataType      string  `pulumi:"odataType"`
}





type SelectAudioTrackByAttributeInput interface {
	pulumi.Input

	ToSelectAudioTrackByAttributeOutput() SelectAudioTrackByAttributeOutput
	ToSelectAudioTrackByAttributeOutputWithContext(context.Context) SelectAudioTrackByAttributeOutput
}

type SelectAudioTrackByAttributeArgs struct {
	Attribute      pulumi.StringInput    `pulumi:"attribute"`
	ChannelMapping pulumi.StringPtrInput `pulumi:"channelMapping"`
	Filter         pulumi.StringInput    `pulumi:"filter"`
	FilterValue    pulumi.StringPtrInput `pulumi:"filterValue"`
	OdataType      pulumi.StringInput    `pulumi:"odataType"`
}

func (SelectAudioTrackByAttributeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectAudioTrackByAttribute)(nil)).Elem()
}

func (i SelectAudioTrackByAttributeArgs) ToSelectAudioTrackByAttributeOutput() SelectAudioTrackByAttributeOutput {
	return i.ToSelectAudioTrackByAttributeOutputWithContext(context.Background())
}

func (i SelectAudioTrackByAttributeArgs) ToSelectAudioTrackByAttributeOutputWithContext(ctx context.Context) SelectAudioTrackByAttributeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectAudioTrackByAttributeOutput)
}

type SelectAudioTrackByAttributeOutput struct{ *pulumi.OutputState }

func (SelectAudioTrackByAttributeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectAudioTrackByAttribute)(nil)).Elem()
}

func (o SelectAudioTrackByAttributeOutput) ToSelectAudioTrackByAttributeOutput() SelectAudioTrackByAttributeOutput {
	return o
}

func (o SelectAudioTrackByAttributeOutput) ToSelectAudioTrackByAttributeOutputWithContext(ctx context.Context) SelectAudioTrackByAttributeOutput {
	return o
}

func (o SelectAudioTrackByAttributeOutput) Attribute() pulumi.StringOutput {
	return o.ApplyT(func(v SelectAudioTrackByAttribute) string { return v.Attribute }).(pulumi.StringOutput)
}

func (o SelectAudioTrackByAttributeOutput) ChannelMapping() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelectAudioTrackByAttribute) *string { return v.ChannelMapping }).(pulumi.StringPtrOutput)
}

func (o SelectAudioTrackByAttributeOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v SelectAudioTrackByAttribute) string { return v.Filter }).(pulumi.StringOutput)
}

func (o SelectAudioTrackByAttributeOutput) FilterValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelectAudioTrackByAttribute) *string { return v.FilterValue }).(pulumi.StringPtrOutput)
}

func (o SelectAudioTrackByAttributeOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v SelectAudioTrackByAttribute) string { return v.OdataType }).(pulumi.StringOutput)
}

type SelectAudioTrackByAttributeResponse struct {
	Attribute      string  `pulumi:"attribute"`
	ChannelMapping *string `pulumi:"channelMapping"`
	Filter         string  `pulumi:"filter"`
	FilterValue    *string `pulumi:"filterValue"`
	OdataType      string  `pulumi:"odataType"`
}





type SelectAudioTrackByAttributeResponseInput interface {
	pulumi.Input

	ToSelectAudioTrackByAttributeResponseOutput() SelectAudioTrackByAttributeResponseOutput
	ToSelectAudioTrackByAttributeResponseOutputWithContext(context.Context) SelectAudioTrackByAttributeResponseOutput
}

type SelectAudioTrackByAttributeResponseArgs struct {
	Attribute      pulumi.StringInput    `pulumi:"attribute"`
	ChannelMapping pulumi.StringPtrInput `pulumi:"channelMapping"`
	Filter         pulumi.StringInput    `pulumi:"filter"`
	FilterValue    pulumi.StringPtrInput `pulumi:"filterValue"`
	OdataType      pulumi.StringInput    `pulumi:"odataType"`
}

func (SelectAudioTrackByAttributeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectAudioTrackByAttributeResponse)(nil)).Elem()
}

func (i SelectAudioTrackByAttributeResponseArgs) ToSelectAudioTrackByAttributeResponseOutput() SelectAudioTrackByAttributeResponseOutput {
	return i.ToSelectAudioTrackByAttributeResponseOutputWithContext(context.Background())
}

func (i SelectAudioTrackByAttributeResponseArgs) ToSelectAudioTrackByAttributeResponseOutputWithContext(ctx context.Context) SelectAudioTrackByAttributeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectAudioTrackByAttributeResponseOutput)
}

type SelectAudioTrackByAttributeResponseOutput struct{ *pulumi.OutputState }

func (SelectAudioTrackByAttributeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectAudioTrackByAttributeResponse)(nil)).Elem()
}

func (o SelectAudioTrackByAttributeResponseOutput) ToSelectAudioTrackByAttributeResponseOutput() SelectAudioTrackByAttributeResponseOutput {
	return o
}

func (o SelectAudioTrackByAttributeResponseOutput) ToSelectAudioTrackByAttributeResponseOutputWithContext(ctx context.Context) SelectAudioTrackByAttributeResponseOutput {
	return o
}

func (o SelectAudioTrackByAttributeResponseOutput) Attribute() pulumi.StringOutput {
	return o.ApplyT(func(v SelectAudioTrackByAttributeResponse) string { return v.Attribute }).(pulumi.StringOutput)
}

func (o SelectAudioTrackByAttributeResponseOutput) ChannelMapping() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelectAudioTrackByAttributeResponse) *string { return v.ChannelMapping }).(pulumi.StringPtrOutput)
}

func (o SelectAudioTrackByAttributeResponseOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v SelectAudioTrackByAttributeResponse) string { return v.Filter }).(pulumi.StringOutput)
}

func (o SelectAudioTrackByAttributeResponseOutput) FilterValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelectAudioTrackByAttributeResponse) *string { return v.FilterValue }).(pulumi.StringPtrOutput)
}

func (o SelectAudioTrackByAttributeResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v SelectAudioTrackByAttributeResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type SelectAudioTrackById struct {
	ChannelMapping *string `pulumi:"channelMapping"`
	OdataType      string  `pulumi:"odataType"`
	TrackId        float64 `pulumi:"trackId"`
}





type SelectAudioTrackByIdInput interface {
	pulumi.Input

	ToSelectAudioTrackByIdOutput() SelectAudioTrackByIdOutput
	ToSelectAudioTrackByIdOutputWithContext(context.Context) SelectAudioTrackByIdOutput
}

type SelectAudioTrackByIdArgs struct {
	ChannelMapping pulumi.StringPtrInput `pulumi:"channelMapping"`
	OdataType      pulumi.StringInput    `pulumi:"odataType"`
	TrackId        pulumi.Float64Input   `pulumi:"trackId"`
}

func (SelectAudioTrackByIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectAudioTrackById)(nil)).Elem()
}

func (i SelectAudioTrackByIdArgs) ToSelectAudioTrackByIdOutput() SelectAudioTrackByIdOutput {
	return i.ToSelectAudioTrackByIdOutputWithContext(context.Background())
}

func (i SelectAudioTrackByIdArgs) ToSelectAudioTrackByIdOutputWithContext(ctx context.Context) SelectAudioTrackByIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectAudioTrackByIdOutput)
}

type SelectAudioTrackByIdOutput struct{ *pulumi.OutputState }

func (SelectAudioTrackByIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectAudioTrackById)(nil)).Elem()
}

func (o SelectAudioTrackByIdOutput) ToSelectAudioTrackByIdOutput() SelectAudioTrackByIdOutput {
	return o
}

func (o SelectAudioTrackByIdOutput) ToSelectAudioTrackByIdOutputWithContext(ctx context.Context) SelectAudioTrackByIdOutput {
	return o
}

func (o SelectAudioTrackByIdOutput) ChannelMapping() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelectAudioTrackById) *string { return v.ChannelMapping }).(pulumi.StringPtrOutput)
}

func (o SelectAudioTrackByIdOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v SelectAudioTrackById) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o SelectAudioTrackByIdOutput) TrackId() pulumi.Float64Output {
	return o.ApplyT(func(v SelectAudioTrackById) float64 { return v.TrackId }).(pulumi.Float64Output)
}

type SelectAudioTrackByIdResponse struct {
	ChannelMapping *string `pulumi:"channelMapping"`
	OdataType      string  `pulumi:"odataType"`
	TrackId        float64 `pulumi:"trackId"`
}





type SelectAudioTrackByIdResponseInput interface {
	pulumi.Input

	ToSelectAudioTrackByIdResponseOutput() SelectAudioTrackByIdResponseOutput
	ToSelectAudioTrackByIdResponseOutputWithContext(context.Context) SelectAudioTrackByIdResponseOutput
}

type SelectAudioTrackByIdResponseArgs struct {
	ChannelMapping pulumi.StringPtrInput `pulumi:"channelMapping"`
	OdataType      pulumi.StringInput    `pulumi:"odataType"`
	TrackId        pulumi.Float64Input   `pulumi:"trackId"`
}

func (SelectAudioTrackByIdResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectAudioTrackByIdResponse)(nil)).Elem()
}

func (i SelectAudioTrackByIdResponseArgs) ToSelectAudioTrackByIdResponseOutput() SelectAudioTrackByIdResponseOutput {
	return i.ToSelectAudioTrackByIdResponseOutputWithContext(context.Background())
}

func (i SelectAudioTrackByIdResponseArgs) ToSelectAudioTrackByIdResponseOutputWithContext(ctx context.Context) SelectAudioTrackByIdResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectAudioTrackByIdResponseOutput)
}

type SelectAudioTrackByIdResponseOutput struct{ *pulumi.OutputState }

func (SelectAudioTrackByIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectAudioTrackByIdResponse)(nil)).Elem()
}

func (o SelectAudioTrackByIdResponseOutput) ToSelectAudioTrackByIdResponseOutput() SelectAudioTrackByIdResponseOutput {
	return o
}

func (o SelectAudioTrackByIdResponseOutput) ToSelectAudioTrackByIdResponseOutputWithContext(ctx context.Context) SelectAudioTrackByIdResponseOutput {
	return o
}

func (o SelectAudioTrackByIdResponseOutput) ChannelMapping() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelectAudioTrackByIdResponse) *string { return v.ChannelMapping }).(pulumi.StringPtrOutput)
}

func (o SelectAudioTrackByIdResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v SelectAudioTrackByIdResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o SelectAudioTrackByIdResponseOutput) TrackId() pulumi.Float64Output {
	return o.ApplyT(func(v SelectAudioTrackByIdResponse) float64 { return v.TrackId }).(pulumi.Float64Output)
}

type SelectVideoTrackByAttribute struct {
	Attribute   string  `pulumi:"attribute"`
	Filter      string  `pulumi:"filter"`
	FilterValue *string `pulumi:"filterValue"`
	OdataType   string  `pulumi:"odataType"`
}





type SelectVideoTrackByAttributeInput interface {
	pulumi.Input

	ToSelectVideoTrackByAttributeOutput() SelectVideoTrackByAttributeOutput
	ToSelectVideoTrackByAttributeOutputWithContext(context.Context) SelectVideoTrackByAttributeOutput
}

type SelectVideoTrackByAttributeArgs struct {
	Attribute   pulumi.StringInput    `pulumi:"attribute"`
	Filter      pulumi.StringInput    `pulumi:"filter"`
	FilterValue pulumi.StringPtrInput `pulumi:"filterValue"`
	OdataType   pulumi.StringInput    `pulumi:"odataType"`
}

func (SelectVideoTrackByAttributeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectVideoTrackByAttribute)(nil)).Elem()
}

func (i SelectVideoTrackByAttributeArgs) ToSelectVideoTrackByAttributeOutput() SelectVideoTrackByAttributeOutput {
	return i.ToSelectVideoTrackByAttributeOutputWithContext(context.Background())
}

func (i SelectVideoTrackByAttributeArgs) ToSelectVideoTrackByAttributeOutputWithContext(ctx context.Context) SelectVideoTrackByAttributeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectVideoTrackByAttributeOutput)
}

type SelectVideoTrackByAttributeOutput struct{ *pulumi.OutputState }

func (SelectVideoTrackByAttributeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectVideoTrackByAttribute)(nil)).Elem()
}

func (o SelectVideoTrackByAttributeOutput) ToSelectVideoTrackByAttributeOutput() SelectVideoTrackByAttributeOutput {
	return o
}

func (o SelectVideoTrackByAttributeOutput) ToSelectVideoTrackByAttributeOutputWithContext(ctx context.Context) SelectVideoTrackByAttributeOutput {
	return o
}

func (o SelectVideoTrackByAttributeOutput) Attribute() pulumi.StringOutput {
	return o.ApplyT(func(v SelectVideoTrackByAttribute) string { return v.Attribute }).(pulumi.StringOutput)
}

func (o SelectVideoTrackByAttributeOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v SelectVideoTrackByAttribute) string { return v.Filter }).(pulumi.StringOutput)
}

func (o SelectVideoTrackByAttributeOutput) FilterValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelectVideoTrackByAttribute) *string { return v.FilterValue }).(pulumi.StringPtrOutput)
}

func (o SelectVideoTrackByAttributeOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v SelectVideoTrackByAttribute) string { return v.OdataType }).(pulumi.StringOutput)
}

type SelectVideoTrackByAttributeResponse struct {
	Attribute   string  `pulumi:"attribute"`
	Filter      string  `pulumi:"filter"`
	FilterValue *string `pulumi:"filterValue"`
	OdataType   string  `pulumi:"odataType"`
}





type SelectVideoTrackByAttributeResponseInput interface {
	pulumi.Input

	ToSelectVideoTrackByAttributeResponseOutput() SelectVideoTrackByAttributeResponseOutput
	ToSelectVideoTrackByAttributeResponseOutputWithContext(context.Context) SelectVideoTrackByAttributeResponseOutput
}

type SelectVideoTrackByAttributeResponseArgs struct {
	Attribute   pulumi.StringInput    `pulumi:"attribute"`
	Filter      pulumi.StringInput    `pulumi:"filter"`
	FilterValue pulumi.StringPtrInput `pulumi:"filterValue"`
	OdataType   pulumi.StringInput    `pulumi:"odataType"`
}

func (SelectVideoTrackByAttributeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectVideoTrackByAttributeResponse)(nil)).Elem()
}

func (i SelectVideoTrackByAttributeResponseArgs) ToSelectVideoTrackByAttributeResponseOutput() SelectVideoTrackByAttributeResponseOutput {
	return i.ToSelectVideoTrackByAttributeResponseOutputWithContext(context.Background())
}

func (i SelectVideoTrackByAttributeResponseArgs) ToSelectVideoTrackByAttributeResponseOutputWithContext(ctx context.Context) SelectVideoTrackByAttributeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectVideoTrackByAttributeResponseOutput)
}

type SelectVideoTrackByAttributeResponseOutput struct{ *pulumi.OutputState }

func (SelectVideoTrackByAttributeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectVideoTrackByAttributeResponse)(nil)).Elem()
}

func (o SelectVideoTrackByAttributeResponseOutput) ToSelectVideoTrackByAttributeResponseOutput() SelectVideoTrackByAttributeResponseOutput {
	return o
}

func (o SelectVideoTrackByAttributeResponseOutput) ToSelectVideoTrackByAttributeResponseOutputWithContext(ctx context.Context) SelectVideoTrackByAttributeResponseOutput {
	return o
}

func (o SelectVideoTrackByAttributeResponseOutput) Attribute() pulumi.StringOutput {
	return o.ApplyT(func(v SelectVideoTrackByAttributeResponse) string { return v.Attribute }).(pulumi.StringOutput)
}

func (o SelectVideoTrackByAttributeResponseOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v SelectVideoTrackByAttributeResponse) string { return v.Filter }).(pulumi.StringOutput)
}

func (o SelectVideoTrackByAttributeResponseOutput) FilterValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelectVideoTrackByAttributeResponse) *string { return v.FilterValue }).(pulumi.StringPtrOutput)
}

func (o SelectVideoTrackByAttributeResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v SelectVideoTrackByAttributeResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type SelectVideoTrackById struct {
	OdataType string  `pulumi:"odataType"`
	TrackId   float64 `pulumi:"trackId"`
}





type SelectVideoTrackByIdInput interface {
	pulumi.Input

	ToSelectVideoTrackByIdOutput() SelectVideoTrackByIdOutput
	ToSelectVideoTrackByIdOutputWithContext(context.Context) SelectVideoTrackByIdOutput
}

type SelectVideoTrackByIdArgs struct {
	OdataType pulumi.StringInput  `pulumi:"odataType"`
	TrackId   pulumi.Float64Input `pulumi:"trackId"`
}

func (SelectVideoTrackByIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectVideoTrackById)(nil)).Elem()
}

func (i SelectVideoTrackByIdArgs) ToSelectVideoTrackByIdOutput() SelectVideoTrackByIdOutput {
	return i.ToSelectVideoTrackByIdOutputWithContext(context.Background())
}

func (i SelectVideoTrackByIdArgs) ToSelectVideoTrackByIdOutputWithContext(ctx context.Context) SelectVideoTrackByIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectVideoTrackByIdOutput)
}

type SelectVideoTrackByIdOutput struct{ *pulumi.OutputState }

func (SelectVideoTrackByIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectVideoTrackById)(nil)).Elem()
}

func (o SelectVideoTrackByIdOutput) ToSelectVideoTrackByIdOutput() SelectVideoTrackByIdOutput {
	return o
}

func (o SelectVideoTrackByIdOutput) ToSelectVideoTrackByIdOutputWithContext(ctx context.Context) SelectVideoTrackByIdOutput {
	return o
}

func (o SelectVideoTrackByIdOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v SelectVideoTrackById) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o SelectVideoTrackByIdOutput) TrackId() pulumi.Float64Output {
	return o.ApplyT(func(v SelectVideoTrackById) float64 { return v.TrackId }).(pulumi.Float64Output)
}

type SelectVideoTrackByIdResponse struct {
	OdataType string  `pulumi:"odataType"`
	TrackId   float64 `pulumi:"trackId"`
}





type SelectVideoTrackByIdResponseInput interface {
	pulumi.Input

	ToSelectVideoTrackByIdResponseOutput() SelectVideoTrackByIdResponseOutput
	ToSelectVideoTrackByIdResponseOutputWithContext(context.Context) SelectVideoTrackByIdResponseOutput
}

type SelectVideoTrackByIdResponseArgs struct {
	OdataType pulumi.StringInput  `pulumi:"odataType"`
	TrackId   pulumi.Float64Input `pulumi:"trackId"`
}

func (SelectVideoTrackByIdResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectVideoTrackByIdResponse)(nil)).Elem()
}

func (i SelectVideoTrackByIdResponseArgs) ToSelectVideoTrackByIdResponseOutput() SelectVideoTrackByIdResponseOutput {
	return i.ToSelectVideoTrackByIdResponseOutputWithContext(context.Background())
}

func (i SelectVideoTrackByIdResponseArgs) ToSelectVideoTrackByIdResponseOutputWithContext(ctx context.Context) SelectVideoTrackByIdResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectVideoTrackByIdResponseOutput)
}

type SelectVideoTrackByIdResponseOutput struct{ *pulumi.OutputState }

func (SelectVideoTrackByIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectVideoTrackByIdResponse)(nil)).Elem()
}

func (o SelectVideoTrackByIdResponseOutput) ToSelectVideoTrackByIdResponseOutput() SelectVideoTrackByIdResponseOutput {
	return o
}

func (o SelectVideoTrackByIdResponseOutput) ToSelectVideoTrackByIdResponseOutputWithContext(ctx context.Context) SelectVideoTrackByIdResponseOutput {
	return o
}

func (o SelectVideoTrackByIdResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v SelectVideoTrackByIdResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o SelectVideoTrackByIdResponseOutput) TrackId() pulumi.Float64Output {
	return o.ApplyT(func(v SelectVideoTrackByIdResponse) float64 { return v.TrackId }).(pulumi.Float64Output)
}

type StandardEncoderPreset struct {
	Codecs    []interface{} `pulumi:"codecs"`
	Filters   *Filters      `pulumi:"filters"`
	Formats   []interface{} `pulumi:"formats"`
	OdataType string        `pulumi:"odataType"`
}





type StandardEncoderPresetInput interface {
	pulumi.Input

	ToStandardEncoderPresetOutput() StandardEncoderPresetOutput
	ToStandardEncoderPresetOutputWithContext(context.Context) StandardEncoderPresetOutput
}

type StandardEncoderPresetArgs struct {
	Codecs    pulumi.ArrayInput  `pulumi:"codecs"`
	Filters   FiltersPtrInput    `pulumi:"filters"`
	Formats   pulumi.ArrayInput  `pulumi:"formats"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (StandardEncoderPresetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardEncoderPreset)(nil)).Elem()
}

func (i StandardEncoderPresetArgs) ToStandardEncoderPresetOutput() StandardEncoderPresetOutput {
	return i.ToStandardEncoderPresetOutputWithContext(context.Background())
}

func (i StandardEncoderPresetArgs) ToStandardEncoderPresetOutputWithContext(ctx context.Context) StandardEncoderPresetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardEncoderPresetOutput)
}

type StandardEncoderPresetOutput struct{ *pulumi.OutputState }

func (StandardEncoderPresetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardEncoderPreset)(nil)).Elem()
}

func (o StandardEncoderPresetOutput) ToStandardEncoderPresetOutput() StandardEncoderPresetOutput {
	return o
}

func (o StandardEncoderPresetOutput) ToStandardEncoderPresetOutputWithContext(ctx context.Context) StandardEncoderPresetOutput {
	return o
}

func (o StandardEncoderPresetOutput) Codecs() pulumi.ArrayOutput {
	return o.ApplyT(func(v StandardEncoderPreset) []interface{} { return v.Codecs }).(pulumi.ArrayOutput)
}

func (o StandardEncoderPresetOutput) Filters() FiltersPtrOutput {
	return o.ApplyT(func(v StandardEncoderPreset) *Filters { return v.Filters }).(FiltersPtrOutput)
}

func (o StandardEncoderPresetOutput) Formats() pulumi.ArrayOutput {
	return o.ApplyT(func(v StandardEncoderPreset) []interface{} { return v.Formats }).(pulumi.ArrayOutput)
}

func (o StandardEncoderPresetOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v StandardEncoderPreset) string { return v.OdataType }).(pulumi.StringOutput)
}

type StandardEncoderPresetResponse struct {
	Codecs    []interface{}    `pulumi:"codecs"`
	Filters   *FiltersResponse `pulumi:"filters"`
	Formats   []interface{}    `pulumi:"formats"`
	OdataType string           `pulumi:"odataType"`
}





type StandardEncoderPresetResponseInput interface {
	pulumi.Input

	ToStandardEncoderPresetResponseOutput() StandardEncoderPresetResponseOutput
	ToStandardEncoderPresetResponseOutputWithContext(context.Context) StandardEncoderPresetResponseOutput
}

type StandardEncoderPresetResponseArgs struct {
	Codecs    pulumi.ArrayInput       `pulumi:"codecs"`
	Filters   FiltersResponsePtrInput `pulumi:"filters"`
	Formats   pulumi.ArrayInput       `pulumi:"formats"`
	OdataType pulumi.StringInput      `pulumi:"odataType"`
}

func (StandardEncoderPresetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardEncoderPresetResponse)(nil)).Elem()
}

func (i StandardEncoderPresetResponseArgs) ToStandardEncoderPresetResponseOutput() StandardEncoderPresetResponseOutput {
	return i.ToStandardEncoderPresetResponseOutputWithContext(context.Background())
}

func (i StandardEncoderPresetResponseArgs) ToStandardEncoderPresetResponseOutputWithContext(ctx context.Context) StandardEncoderPresetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardEncoderPresetResponseOutput)
}

type StandardEncoderPresetResponseOutput struct{ *pulumi.OutputState }

func (StandardEncoderPresetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardEncoderPresetResponse)(nil)).Elem()
}

func (o StandardEncoderPresetResponseOutput) ToStandardEncoderPresetResponseOutput() StandardEncoderPresetResponseOutput {
	return o
}

func (o StandardEncoderPresetResponseOutput) ToStandardEncoderPresetResponseOutputWithContext(ctx context.Context) StandardEncoderPresetResponseOutput {
	return o
}

func (o StandardEncoderPresetResponseOutput) Codecs() pulumi.ArrayOutput {
	return o.ApplyT(func(v StandardEncoderPresetResponse) []interface{} { return v.Codecs }).(pulumi.ArrayOutput)
}

func (o StandardEncoderPresetResponseOutput) Filters() FiltersResponsePtrOutput {
	return o.ApplyT(func(v StandardEncoderPresetResponse) *FiltersResponse { return v.Filters }).(FiltersResponsePtrOutput)
}

func (o StandardEncoderPresetResponseOutput) Formats() pulumi.ArrayOutput {
	return o.ApplyT(func(v StandardEncoderPresetResponse) []interface{} { return v.Formats }).(pulumi.ArrayOutput)
}

func (o StandardEncoderPresetResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v StandardEncoderPresetResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type StorageAccount struct {
	Id   *string `pulumi:"id"`
	Type string  `pulumi:"type"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Type pulumi.StringInput    `pulumi:"type"`
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

func (o StorageAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Type }).(pulumi.StringOutput)
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
	Id   *string `pulumi:"id"`
	Type string  `pulumi:"type"`
}





type StorageAccountResponseInput interface {
	pulumi.Input

	ToStorageAccountResponseOutput() StorageAccountResponseOutput
	ToStorageAccountResponseOutputWithContext(context.Context) StorageAccountResponseOutput
}

type StorageAccountResponseArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Type pulumi.StringInput    `pulumi:"type"`
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

func (o StorageAccountResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Type }).(pulumi.StringOutput)
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

type StreamingEndpointAccessControl struct {
	Akamai *AkamaiAccessControl `pulumi:"akamai"`
	Ip     *IPAccessControl     `pulumi:"ip"`
}





type StreamingEndpointAccessControlInput interface {
	pulumi.Input

	ToStreamingEndpointAccessControlOutput() StreamingEndpointAccessControlOutput
	ToStreamingEndpointAccessControlOutputWithContext(context.Context) StreamingEndpointAccessControlOutput
}

type StreamingEndpointAccessControlArgs struct {
	Akamai AkamaiAccessControlPtrInput `pulumi:"akamai"`
	Ip     IPAccessControlPtrInput     `pulumi:"ip"`
}

func (StreamingEndpointAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpointAccessControl)(nil)).Elem()
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlOutput() StreamingEndpointAccessControlOutput {
	return i.ToStreamingEndpointAccessControlOutputWithContext(context.Background())
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlOutputWithContext(ctx context.Context) StreamingEndpointAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlOutput)
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return i.ToStreamingEndpointAccessControlPtrOutputWithContext(context.Background())
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlOutput).ToStreamingEndpointAccessControlPtrOutputWithContext(ctx)
}









type StreamingEndpointAccessControlPtrInput interface {
	pulumi.Input

	ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput
	ToStreamingEndpointAccessControlPtrOutputWithContext(context.Context) StreamingEndpointAccessControlPtrOutput
}

type streamingEndpointAccessControlPtrType StreamingEndpointAccessControlArgs

func StreamingEndpointAccessControlPtr(v *StreamingEndpointAccessControlArgs) StreamingEndpointAccessControlPtrInput {
	return (*streamingEndpointAccessControlPtrType)(v)
}

func (*streamingEndpointAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpointAccessControl)(nil)).Elem()
}

func (i *streamingEndpointAccessControlPtrType) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return i.ToStreamingEndpointAccessControlPtrOutputWithContext(context.Background())
}

func (i *streamingEndpointAccessControlPtrType) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlPtrOutput)
}

type StreamingEndpointAccessControlOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpointAccessControl)(nil)).Elem()
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlOutput() StreamingEndpointAccessControlOutput {
	return o
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlOutputWithContext(ctx context.Context) StreamingEndpointAccessControlOutput {
	return o
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return o.ToStreamingEndpointAccessControlPtrOutputWithContext(context.Background())
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingEndpointAccessControl) *StreamingEndpointAccessControl {
		return &v
	}).(StreamingEndpointAccessControlPtrOutput)
}

func (o StreamingEndpointAccessControlOutput) Akamai() AkamaiAccessControlPtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControl) *AkamaiAccessControl { return v.Akamai }).(AkamaiAccessControlPtrOutput)
}

func (o StreamingEndpointAccessControlOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControl) *IPAccessControl { return v.Ip }).(IPAccessControlPtrOutput)
}

type StreamingEndpointAccessControlPtrOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpointAccessControl)(nil)).Elem()
}

func (o StreamingEndpointAccessControlPtrOutput) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return o
}

func (o StreamingEndpointAccessControlPtrOutput) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return o
}

func (o StreamingEndpointAccessControlPtrOutput) Elem() StreamingEndpointAccessControlOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControl) StreamingEndpointAccessControl {
		if v != nil {
			return *v
		}
		var ret StreamingEndpointAccessControl
		return ret
	}).(StreamingEndpointAccessControlOutput)
}

func (o StreamingEndpointAccessControlPtrOutput) Akamai() AkamaiAccessControlPtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControl) *AkamaiAccessControl {
		if v == nil {
			return nil
		}
		return v.Akamai
	}).(AkamaiAccessControlPtrOutput)
}

func (o StreamingEndpointAccessControlPtrOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControl) *IPAccessControl {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlPtrOutput)
}

type StreamingEndpointAccessControlResponse struct {
	Akamai *AkamaiAccessControlResponse `pulumi:"akamai"`
	Ip     *IPAccessControlResponse     `pulumi:"ip"`
}





type StreamingEndpointAccessControlResponseInput interface {
	pulumi.Input

	ToStreamingEndpointAccessControlResponseOutput() StreamingEndpointAccessControlResponseOutput
	ToStreamingEndpointAccessControlResponseOutputWithContext(context.Context) StreamingEndpointAccessControlResponseOutput
}

type StreamingEndpointAccessControlResponseArgs struct {
	Akamai AkamaiAccessControlResponsePtrInput `pulumi:"akamai"`
	Ip     IPAccessControlResponsePtrInput     `pulumi:"ip"`
}

func (StreamingEndpointAccessControlResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpointAccessControlResponse)(nil)).Elem()
}

func (i StreamingEndpointAccessControlResponseArgs) ToStreamingEndpointAccessControlResponseOutput() StreamingEndpointAccessControlResponseOutput {
	return i.ToStreamingEndpointAccessControlResponseOutputWithContext(context.Background())
}

func (i StreamingEndpointAccessControlResponseArgs) ToStreamingEndpointAccessControlResponseOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlResponseOutput)
}

func (i StreamingEndpointAccessControlResponseArgs) ToStreamingEndpointAccessControlResponsePtrOutput() StreamingEndpointAccessControlResponsePtrOutput {
	return i.ToStreamingEndpointAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i StreamingEndpointAccessControlResponseArgs) ToStreamingEndpointAccessControlResponsePtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlResponseOutput).ToStreamingEndpointAccessControlResponsePtrOutputWithContext(ctx)
}









type StreamingEndpointAccessControlResponsePtrInput interface {
	pulumi.Input

	ToStreamingEndpointAccessControlResponsePtrOutput() StreamingEndpointAccessControlResponsePtrOutput
	ToStreamingEndpointAccessControlResponsePtrOutputWithContext(context.Context) StreamingEndpointAccessControlResponsePtrOutput
}

type streamingEndpointAccessControlResponsePtrType StreamingEndpointAccessControlResponseArgs

func StreamingEndpointAccessControlResponsePtr(v *StreamingEndpointAccessControlResponseArgs) StreamingEndpointAccessControlResponsePtrInput {
	return (*streamingEndpointAccessControlResponsePtrType)(v)
}

func (*streamingEndpointAccessControlResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpointAccessControlResponse)(nil)).Elem()
}

func (i *streamingEndpointAccessControlResponsePtrType) ToStreamingEndpointAccessControlResponsePtrOutput() StreamingEndpointAccessControlResponsePtrOutput {
	return i.ToStreamingEndpointAccessControlResponsePtrOutputWithContext(context.Background())
}

func (i *streamingEndpointAccessControlResponsePtrType) ToStreamingEndpointAccessControlResponsePtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlResponsePtrOutput)
}

type StreamingEndpointAccessControlResponseOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpointAccessControlResponse)(nil)).Elem()
}

func (o StreamingEndpointAccessControlResponseOutput) ToStreamingEndpointAccessControlResponseOutput() StreamingEndpointAccessControlResponseOutput {
	return o
}

func (o StreamingEndpointAccessControlResponseOutput) ToStreamingEndpointAccessControlResponseOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponseOutput {
	return o
}

func (o StreamingEndpointAccessControlResponseOutput) ToStreamingEndpointAccessControlResponsePtrOutput() StreamingEndpointAccessControlResponsePtrOutput {
	return o.ToStreamingEndpointAccessControlResponsePtrOutputWithContext(context.Background())
}

func (o StreamingEndpointAccessControlResponseOutput) ToStreamingEndpointAccessControlResponsePtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingEndpointAccessControlResponse) *StreamingEndpointAccessControlResponse {
		return &v
	}).(StreamingEndpointAccessControlResponsePtrOutput)
}

func (o StreamingEndpointAccessControlResponseOutput) Akamai() AkamaiAccessControlResponsePtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControlResponse) *AkamaiAccessControlResponse { return v.Akamai }).(AkamaiAccessControlResponsePtrOutput)
}

func (o StreamingEndpointAccessControlResponseOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControlResponse) *IPAccessControlResponse { return v.Ip }).(IPAccessControlResponsePtrOutput)
}

type StreamingEndpointAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpointAccessControlResponse)(nil)).Elem()
}

func (o StreamingEndpointAccessControlResponsePtrOutput) ToStreamingEndpointAccessControlResponsePtrOutput() StreamingEndpointAccessControlResponsePtrOutput {
	return o
}

func (o StreamingEndpointAccessControlResponsePtrOutput) ToStreamingEndpointAccessControlResponsePtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponsePtrOutput {
	return o
}

func (o StreamingEndpointAccessControlResponsePtrOutput) Elem() StreamingEndpointAccessControlResponseOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControlResponse) StreamingEndpointAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret StreamingEndpointAccessControlResponse
		return ret
	}).(StreamingEndpointAccessControlResponseOutput)
}

func (o StreamingEndpointAccessControlResponsePtrOutput) Akamai() AkamaiAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControlResponse) *AkamaiAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Akamai
	}).(AkamaiAccessControlResponsePtrOutput)
}

func (o StreamingEndpointAccessControlResponsePtrOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControlResponse) *IPAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlResponsePtrOutput)
}

type StreamingLocatorContentKey struct {
	Id                              string  `pulumi:"id"`
	LabelReferenceInStreamingPolicy *string `pulumi:"labelReferenceInStreamingPolicy"`
	Value                           *string `pulumi:"value"`
}





type StreamingLocatorContentKeyInput interface {
	pulumi.Input

	ToStreamingLocatorContentKeyOutput() StreamingLocatorContentKeyOutput
	ToStreamingLocatorContentKeyOutputWithContext(context.Context) StreamingLocatorContentKeyOutput
}

type StreamingLocatorContentKeyArgs struct {
	Id                              pulumi.StringInput    `pulumi:"id"`
	LabelReferenceInStreamingPolicy pulumi.StringPtrInput `pulumi:"labelReferenceInStreamingPolicy"`
	Value                           pulumi.StringPtrInput `pulumi:"value"`
}

func (StreamingLocatorContentKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingLocatorContentKey)(nil)).Elem()
}

func (i StreamingLocatorContentKeyArgs) ToStreamingLocatorContentKeyOutput() StreamingLocatorContentKeyOutput {
	return i.ToStreamingLocatorContentKeyOutputWithContext(context.Background())
}

func (i StreamingLocatorContentKeyArgs) ToStreamingLocatorContentKeyOutputWithContext(ctx context.Context) StreamingLocatorContentKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorContentKeyOutput)
}





type StreamingLocatorContentKeyArrayInput interface {
	pulumi.Input

	ToStreamingLocatorContentKeyArrayOutput() StreamingLocatorContentKeyArrayOutput
	ToStreamingLocatorContentKeyArrayOutputWithContext(context.Context) StreamingLocatorContentKeyArrayOutput
}

type StreamingLocatorContentKeyArray []StreamingLocatorContentKeyInput

func (StreamingLocatorContentKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingLocatorContentKey)(nil)).Elem()
}

func (i StreamingLocatorContentKeyArray) ToStreamingLocatorContentKeyArrayOutput() StreamingLocatorContentKeyArrayOutput {
	return i.ToStreamingLocatorContentKeyArrayOutputWithContext(context.Background())
}

func (i StreamingLocatorContentKeyArray) ToStreamingLocatorContentKeyArrayOutputWithContext(ctx context.Context) StreamingLocatorContentKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorContentKeyArrayOutput)
}

type StreamingLocatorContentKeyOutput struct{ *pulumi.OutputState }

func (StreamingLocatorContentKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingLocatorContentKey)(nil)).Elem()
}

func (o StreamingLocatorContentKeyOutput) ToStreamingLocatorContentKeyOutput() StreamingLocatorContentKeyOutput {
	return o
}

func (o StreamingLocatorContentKeyOutput) ToStreamingLocatorContentKeyOutputWithContext(ctx context.Context) StreamingLocatorContentKeyOutput {
	return o
}

func (o StreamingLocatorContentKeyOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingLocatorContentKey) string { return v.Id }).(pulumi.StringOutput)
}

func (o StreamingLocatorContentKeyOutput) LabelReferenceInStreamingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingLocatorContentKey) *string { return v.LabelReferenceInStreamingPolicy }).(pulumi.StringPtrOutput)
}

func (o StreamingLocatorContentKeyOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingLocatorContentKey) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type StreamingLocatorContentKeyArrayOutput struct{ *pulumi.OutputState }

func (StreamingLocatorContentKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingLocatorContentKey)(nil)).Elem()
}

func (o StreamingLocatorContentKeyArrayOutput) ToStreamingLocatorContentKeyArrayOutput() StreamingLocatorContentKeyArrayOutput {
	return o
}

func (o StreamingLocatorContentKeyArrayOutput) ToStreamingLocatorContentKeyArrayOutputWithContext(ctx context.Context) StreamingLocatorContentKeyArrayOutput {
	return o
}

func (o StreamingLocatorContentKeyArrayOutput) Index(i pulumi.IntInput) StreamingLocatorContentKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamingLocatorContentKey {
		return vs[0].([]StreamingLocatorContentKey)[vs[1].(int)]
	}).(StreamingLocatorContentKeyOutput)
}

type StreamingLocatorContentKeyResponse struct {
	Id                              string                   `pulumi:"id"`
	LabelReferenceInStreamingPolicy *string                  `pulumi:"labelReferenceInStreamingPolicy"`
	PolicyName                      string                   `pulumi:"policyName"`
	Tracks                          []TrackSelectionResponse `pulumi:"tracks"`
	Type                            string                   `pulumi:"type"`
	Value                           *string                  `pulumi:"value"`
}





type StreamingLocatorContentKeyResponseInput interface {
	pulumi.Input

	ToStreamingLocatorContentKeyResponseOutput() StreamingLocatorContentKeyResponseOutput
	ToStreamingLocatorContentKeyResponseOutputWithContext(context.Context) StreamingLocatorContentKeyResponseOutput
}

type StreamingLocatorContentKeyResponseArgs struct {
	Id                              pulumi.StringInput               `pulumi:"id"`
	LabelReferenceInStreamingPolicy pulumi.StringPtrInput            `pulumi:"labelReferenceInStreamingPolicy"`
	PolicyName                      pulumi.StringInput               `pulumi:"policyName"`
	Tracks                          TrackSelectionResponseArrayInput `pulumi:"tracks"`
	Type                            pulumi.StringInput               `pulumi:"type"`
	Value                           pulumi.StringPtrInput            `pulumi:"value"`
}

func (StreamingLocatorContentKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingLocatorContentKeyResponse)(nil)).Elem()
}

func (i StreamingLocatorContentKeyResponseArgs) ToStreamingLocatorContentKeyResponseOutput() StreamingLocatorContentKeyResponseOutput {
	return i.ToStreamingLocatorContentKeyResponseOutputWithContext(context.Background())
}

func (i StreamingLocatorContentKeyResponseArgs) ToStreamingLocatorContentKeyResponseOutputWithContext(ctx context.Context) StreamingLocatorContentKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorContentKeyResponseOutput)
}





type StreamingLocatorContentKeyResponseArrayInput interface {
	pulumi.Input

	ToStreamingLocatorContentKeyResponseArrayOutput() StreamingLocatorContentKeyResponseArrayOutput
	ToStreamingLocatorContentKeyResponseArrayOutputWithContext(context.Context) StreamingLocatorContentKeyResponseArrayOutput
}

type StreamingLocatorContentKeyResponseArray []StreamingLocatorContentKeyResponseInput

func (StreamingLocatorContentKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingLocatorContentKeyResponse)(nil)).Elem()
}

func (i StreamingLocatorContentKeyResponseArray) ToStreamingLocatorContentKeyResponseArrayOutput() StreamingLocatorContentKeyResponseArrayOutput {
	return i.ToStreamingLocatorContentKeyResponseArrayOutputWithContext(context.Background())
}

func (i StreamingLocatorContentKeyResponseArray) ToStreamingLocatorContentKeyResponseArrayOutputWithContext(ctx context.Context) StreamingLocatorContentKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorContentKeyResponseArrayOutput)
}

type StreamingLocatorContentKeyResponseOutput struct{ *pulumi.OutputState }

func (StreamingLocatorContentKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingLocatorContentKeyResponse)(nil)).Elem()
}

func (o StreamingLocatorContentKeyResponseOutput) ToStreamingLocatorContentKeyResponseOutput() StreamingLocatorContentKeyResponseOutput {
	return o
}

func (o StreamingLocatorContentKeyResponseOutput) ToStreamingLocatorContentKeyResponseOutputWithContext(ctx context.Context) StreamingLocatorContentKeyResponseOutput {
	return o
}

func (o StreamingLocatorContentKeyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StreamingLocatorContentKeyResponseOutput) LabelReferenceInStreamingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) *string { return v.LabelReferenceInStreamingPolicy }).(pulumi.StringPtrOutput)
}

func (o StreamingLocatorContentKeyResponseOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) string { return v.PolicyName }).(pulumi.StringOutput)
}

func (o StreamingLocatorContentKeyResponseOutput) Tracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) []TrackSelectionResponse { return v.Tracks }).(TrackSelectionResponseArrayOutput)
}

func (o StreamingLocatorContentKeyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o StreamingLocatorContentKeyResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type StreamingLocatorContentKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (StreamingLocatorContentKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingLocatorContentKeyResponse)(nil)).Elem()
}

func (o StreamingLocatorContentKeyResponseArrayOutput) ToStreamingLocatorContentKeyResponseArrayOutput() StreamingLocatorContentKeyResponseArrayOutput {
	return o
}

func (o StreamingLocatorContentKeyResponseArrayOutput) ToStreamingLocatorContentKeyResponseArrayOutputWithContext(ctx context.Context) StreamingLocatorContentKeyResponseArrayOutput {
	return o
}

func (o StreamingLocatorContentKeyResponseArrayOutput) Index(i pulumi.IntInput) StreamingLocatorContentKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamingLocatorContentKeyResponse {
		return vs[0].([]StreamingLocatorContentKeyResponse)[vs[1].(int)]
	}).(StreamingLocatorContentKeyResponseOutput)
}

type StreamingPathResponse struct {
	EncryptionScheme  string   `pulumi:"encryptionScheme"`
	Paths             []string `pulumi:"paths"`
	StreamingProtocol string   `pulumi:"streamingProtocol"`
}





type StreamingPathResponseInput interface {
	pulumi.Input

	ToStreamingPathResponseOutput() StreamingPathResponseOutput
	ToStreamingPathResponseOutputWithContext(context.Context) StreamingPathResponseOutput
}

type StreamingPathResponseArgs struct {
	EncryptionScheme  pulumi.StringInput      `pulumi:"encryptionScheme"`
	Paths             pulumi.StringArrayInput `pulumi:"paths"`
	StreamingProtocol pulumi.StringInput      `pulumi:"streamingProtocol"`
}

func (StreamingPathResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPathResponse)(nil)).Elem()
}

func (i StreamingPathResponseArgs) ToStreamingPathResponseOutput() StreamingPathResponseOutput {
	return i.ToStreamingPathResponseOutputWithContext(context.Background())
}

func (i StreamingPathResponseArgs) ToStreamingPathResponseOutputWithContext(ctx context.Context) StreamingPathResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPathResponseOutput)
}





type StreamingPathResponseArrayInput interface {
	pulumi.Input

	ToStreamingPathResponseArrayOutput() StreamingPathResponseArrayOutput
	ToStreamingPathResponseArrayOutputWithContext(context.Context) StreamingPathResponseArrayOutput
}

type StreamingPathResponseArray []StreamingPathResponseInput

func (StreamingPathResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingPathResponse)(nil)).Elem()
}

func (i StreamingPathResponseArray) ToStreamingPathResponseArrayOutput() StreamingPathResponseArrayOutput {
	return i.ToStreamingPathResponseArrayOutputWithContext(context.Background())
}

func (i StreamingPathResponseArray) ToStreamingPathResponseArrayOutputWithContext(ctx context.Context) StreamingPathResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPathResponseArrayOutput)
}

type StreamingPathResponseOutput struct{ *pulumi.OutputState }

func (StreamingPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPathResponse)(nil)).Elem()
}

func (o StreamingPathResponseOutput) ToStreamingPathResponseOutput() StreamingPathResponseOutput {
	return o
}

func (o StreamingPathResponseOutput) ToStreamingPathResponseOutputWithContext(ctx context.Context) StreamingPathResponseOutput {
	return o
}

func (o StreamingPathResponseOutput) EncryptionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingPathResponse) string { return v.EncryptionScheme }).(pulumi.StringOutput)
}

func (o StreamingPathResponseOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StreamingPathResponse) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o StreamingPathResponseOutput) StreamingProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingPathResponse) string { return v.StreamingProtocol }).(pulumi.StringOutput)
}

type StreamingPathResponseArrayOutput struct{ *pulumi.OutputState }

func (StreamingPathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingPathResponse)(nil)).Elem()
}

func (o StreamingPathResponseArrayOutput) ToStreamingPathResponseArrayOutput() StreamingPathResponseArrayOutput {
	return o
}

func (o StreamingPathResponseArrayOutput) ToStreamingPathResponseArrayOutputWithContext(ctx context.Context) StreamingPathResponseArrayOutput {
	return o
}

func (o StreamingPathResponseArrayOutput) Index(i pulumi.IntInput) StreamingPathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamingPathResponse {
		return vs[0].([]StreamingPathResponse)[vs[1].(int)]
	}).(StreamingPathResponseOutput)
}

type StreamingPolicyContentKey struct {
	Label      *string          `pulumi:"label"`
	PolicyName *string          `pulumi:"policyName"`
	Tracks     []TrackSelection `pulumi:"tracks"`
}





type StreamingPolicyContentKeyInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeyOutput() StreamingPolicyContentKeyOutput
	ToStreamingPolicyContentKeyOutputWithContext(context.Context) StreamingPolicyContentKeyOutput
}

type StreamingPolicyContentKeyArgs struct {
	Label      pulumi.StringPtrInput    `pulumi:"label"`
	PolicyName pulumi.StringPtrInput    `pulumi:"policyName"`
	Tracks     TrackSelectionArrayInput `pulumi:"tracks"`
}

func (StreamingPolicyContentKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKey)(nil)).Elem()
}

func (i StreamingPolicyContentKeyArgs) ToStreamingPolicyContentKeyOutput() StreamingPolicyContentKeyOutput {
	return i.ToStreamingPolicyContentKeyOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeyArgs) ToStreamingPolicyContentKeyOutputWithContext(ctx context.Context) StreamingPolicyContentKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeyOutput)
}





type StreamingPolicyContentKeyArrayInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeyArrayOutput() StreamingPolicyContentKeyArrayOutput
	ToStreamingPolicyContentKeyArrayOutputWithContext(context.Context) StreamingPolicyContentKeyArrayOutput
}

type StreamingPolicyContentKeyArray []StreamingPolicyContentKeyInput

func (StreamingPolicyContentKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingPolicyContentKey)(nil)).Elem()
}

func (i StreamingPolicyContentKeyArray) ToStreamingPolicyContentKeyArrayOutput() StreamingPolicyContentKeyArrayOutput {
	return i.ToStreamingPolicyContentKeyArrayOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeyArray) ToStreamingPolicyContentKeyArrayOutputWithContext(ctx context.Context) StreamingPolicyContentKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeyArrayOutput)
}

type StreamingPolicyContentKeyOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKey)(nil)).Elem()
}

func (o StreamingPolicyContentKeyOutput) ToStreamingPolicyContentKeyOutput() StreamingPolicyContentKeyOutput {
	return o
}

func (o StreamingPolicyContentKeyOutput) ToStreamingPolicyContentKeyOutputWithContext(ctx context.Context) StreamingPolicyContentKeyOutput {
	return o
}

func (o StreamingPolicyContentKeyOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKey) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyContentKeyOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKey) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyContentKeyOutput) Tracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v StreamingPolicyContentKey) []TrackSelection { return v.Tracks }).(TrackSelectionArrayOutput)
}

type StreamingPolicyContentKeyArrayOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingPolicyContentKey)(nil)).Elem()
}

func (o StreamingPolicyContentKeyArrayOutput) ToStreamingPolicyContentKeyArrayOutput() StreamingPolicyContentKeyArrayOutput {
	return o
}

func (o StreamingPolicyContentKeyArrayOutput) ToStreamingPolicyContentKeyArrayOutputWithContext(ctx context.Context) StreamingPolicyContentKeyArrayOutput {
	return o
}

func (o StreamingPolicyContentKeyArrayOutput) Index(i pulumi.IntInput) StreamingPolicyContentKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamingPolicyContentKey {
		return vs[0].([]StreamingPolicyContentKey)[vs[1].(int)]
	}).(StreamingPolicyContentKeyOutput)
}

type StreamingPolicyContentKeyResponse struct {
	Label      *string                  `pulumi:"label"`
	PolicyName *string                  `pulumi:"policyName"`
	Tracks     []TrackSelectionResponse `pulumi:"tracks"`
}





type StreamingPolicyContentKeyResponseInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeyResponseOutput() StreamingPolicyContentKeyResponseOutput
	ToStreamingPolicyContentKeyResponseOutputWithContext(context.Context) StreamingPolicyContentKeyResponseOutput
}

type StreamingPolicyContentKeyResponseArgs struct {
	Label      pulumi.StringPtrInput            `pulumi:"label"`
	PolicyName pulumi.StringPtrInput            `pulumi:"policyName"`
	Tracks     TrackSelectionResponseArrayInput `pulumi:"tracks"`
}

func (StreamingPolicyContentKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKeyResponse)(nil)).Elem()
}

func (i StreamingPolicyContentKeyResponseArgs) ToStreamingPolicyContentKeyResponseOutput() StreamingPolicyContentKeyResponseOutput {
	return i.ToStreamingPolicyContentKeyResponseOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeyResponseArgs) ToStreamingPolicyContentKeyResponseOutputWithContext(ctx context.Context) StreamingPolicyContentKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeyResponseOutput)
}





type StreamingPolicyContentKeyResponseArrayInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeyResponseArrayOutput() StreamingPolicyContentKeyResponseArrayOutput
	ToStreamingPolicyContentKeyResponseArrayOutputWithContext(context.Context) StreamingPolicyContentKeyResponseArrayOutput
}

type StreamingPolicyContentKeyResponseArray []StreamingPolicyContentKeyResponseInput

func (StreamingPolicyContentKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingPolicyContentKeyResponse)(nil)).Elem()
}

func (i StreamingPolicyContentKeyResponseArray) ToStreamingPolicyContentKeyResponseArrayOutput() StreamingPolicyContentKeyResponseArrayOutput {
	return i.ToStreamingPolicyContentKeyResponseArrayOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeyResponseArray) ToStreamingPolicyContentKeyResponseArrayOutputWithContext(ctx context.Context) StreamingPolicyContentKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeyResponseArrayOutput)
}

type StreamingPolicyContentKeyResponseOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKeyResponse)(nil)).Elem()
}

func (o StreamingPolicyContentKeyResponseOutput) ToStreamingPolicyContentKeyResponseOutput() StreamingPolicyContentKeyResponseOutput {
	return o
}

func (o StreamingPolicyContentKeyResponseOutput) ToStreamingPolicyContentKeyResponseOutputWithContext(ctx context.Context) StreamingPolicyContentKeyResponseOutput {
	return o
}

func (o StreamingPolicyContentKeyResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeyResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyContentKeyResponseOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeyResponse) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyContentKeyResponseOutput) Tracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeyResponse) []TrackSelectionResponse { return v.Tracks }).(TrackSelectionResponseArrayOutput)
}

type StreamingPolicyContentKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingPolicyContentKeyResponse)(nil)).Elem()
}

func (o StreamingPolicyContentKeyResponseArrayOutput) ToStreamingPolicyContentKeyResponseArrayOutput() StreamingPolicyContentKeyResponseArrayOutput {
	return o
}

func (o StreamingPolicyContentKeyResponseArrayOutput) ToStreamingPolicyContentKeyResponseArrayOutputWithContext(ctx context.Context) StreamingPolicyContentKeyResponseArrayOutput {
	return o
}

func (o StreamingPolicyContentKeyResponseArrayOutput) Index(i pulumi.IntInput) StreamingPolicyContentKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamingPolicyContentKeyResponse {
		return vs[0].([]StreamingPolicyContentKeyResponse)[vs[1].(int)]
	}).(StreamingPolicyContentKeyResponseOutput)
}

type StreamingPolicyContentKeys struct {
	DefaultKey         *DefaultKey                 `pulumi:"defaultKey"`
	KeyToTrackMappings []StreamingPolicyContentKey `pulumi:"keyToTrackMappings"`
}





type StreamingPolicyContentKeysInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeysOutput() StreamingPolicyContentKeysOutput
	ToStreamingPolicyContentKeysOutputWithContext(context.Context) StreamingPolicyContentKeysOutput
}

type StreamingPolicyContentKeysArgs struct {
	DefaultKey         DefaultKeyPtrInput                  `pulumi:"defaultKey"`
	KeyToTrackMappings StreamingPolicyContentKeyArrayInput `pulumi:"keyToTrackMappings"`
}

func (StreamingPolicyContentKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKeys)(nil)).Elem()
}

func (i StreamingPolicyContentKeysArgs) ToStreamingPolicyContentKeysOutput() StreamingPolicyContentKeysOutput {
	return i.ToStreamingPolicyContentKeysOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeysArgs) ToStreamingPolicyContentKeysOutputWithContext(ctx context.Context) StreamingPolicyContentKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeysOutput)
}

func (i StreamingPolicyContentKeysArgs) ToStreamingPolicyContentKeysPtrOutput() StreamingPolicyContentKeysPtrOutput {
	return i.ToStreamingPolicyContentKeysPtrOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeysArgs) ToStreamingPolicyContentKeysPtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeysOutput).ToStreamingPolicyContentKeysPtrOutputWithContext(ctx)
}









type StreamingPolicyContentKeysPtrInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeysPtrOutput() StreamingPolicyContentKeysPtrOutput
	ToStreamingPolicyContentKeysPtrOutputWithContext(context.Context) StreamingPolicyContentKeysPtrOutput
}

type streamingPolicyContentKeysPtrType StreamingPolicyContentKeysArgs

func StreamingPolicyContentKeysPtr(v *StreamingPolicyContentKeysArgs) StreamingPolicyContentKeysPtrInput {
	return (*streamingPolicyContentKeysPtrType)(v)
}

func (*streamingPolicyContentKeysPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyContentKeys)(nil)).Elem()
}

func (i *streamingPolicyContentKeysPtrType) ToStreamingPolicyContentKeysPtrOutput() StreamingPolicyContentKeysPtrOutput {
	return i.ToStreamingPolicyContentKeysPtrOutputWithContext(context.Background())
}

func (i *streamingPolicyContentKeysPtrType) ToStreamingPolicyContentKeysPtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeysPtrOutput)
}

type StreamingPolicyContentKeysOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKeys)(nil)).Elem()
}

func (o StreamingPolicyContentKeysOutput) ToStreamingPolicyContentKeysOutput() StreamingPolicyContentKeysOutput {
	return o
}

func (o StreamingPolicyContentKeysOutput) ToStreamingPolicyContentKeysOutputWithContext(ctx context.Context) StreamingPolicyContentKeysOutput {
	return o
}

func (o StreamingPolicyContentKeysOutput) ToStreamingPolicyContentKeysPtrOutput() StreamingPolicyContentKeysPtrOutput {
	return o.ToStreamingPolicyContentKeysPtrOutputWithContext(context.Background())
}

func (o StreamingPolicyContentKeysOutput) ToStreamingPolicyContentKeysPtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyContentKeys) *StreamingPolicyContentKeys {
		return &v
	}).(StreamingPolicyContentKeysPtrOutput)
}

func (o StreamingPolicyContentKeysOutput) DefaultKey() DefaultKeyPtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeys) *DefaultKey { return v.DefaultKey }).(DefaultKeyPtrOutput)
}

func (o StreamingPolicyContentKeysOutput) KeyToTrackMappings() StreamingPolicyContentKeyArrayOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeys) []StreamingPolicyContentKey { return v.KeyToTrackMappings }).(StreamingPolicyContentKeyArrayOutput)
}

type StreamingPolicyContentKeysPtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeysPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyContentKeys)(nil)).Elem()
}

func (o StreamingPolicyContentKeysPtrOutput) ToStreamingPolicyContentKeysPtrOutput() StreamingPolicyContentKeysPtrOutput {
	return o
}

func (o StreamingPolicyContentKeysPtrOutput) ToStreamingPolicyContentKeysPtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysPtrOutput {
	return o
}

func (o StreamingPolicyContentKeysPtrOutput) Elem() StreamingPolicyContentKeysOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeys) StreamingPolicyContentKeys {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyContentKeys
		return ret
	}).(StreamingPolicyContentKeysOutput)
}

func (o StreamingPolicyContentKeysPtrOutput) DefaultKey() DefaultKeyPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeys) *DefaultKey {
		if v == nil {
			return nil
		}
		return v.DefaultKey
	}).(DefaultKeyPtrOutput)
}

func (o StreamingPolicyContentKeysPtrOutput) KeyToTrackMappings() StreamingPolicyContentKeyArrayOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeys) []StreamingPolicyContentKey {
		if v == nil {
			return nil
		}
		return v.KeyToTrackMappings
	}).(StreamingPolicyContentKeyArrayOutput)
}

type StreamingPolicyContentKeysResponse struct {
	DefaultKey         *DefaultKeyResponse                 `pulumi:"defaultKey"`
	KeyToTrackMappings []StreamingPolicyContentKeyResponse `pulumi:"keyToTrackMappings"`
}





type StreamingPolicyContentKeysResponseInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeysResponseOutput() StreamingPolicyContentKeysResponseOutput
	ToStreamingPolicyContentKeysResponseOutputWithContext(context.Context) StreamingPolicyContentKeysResponseOutput
}

type StreamingPolicyContentKeysResponseArgs struct {
	DefaultKey         DefaultKeyResponsePtrInput                  `pulumi:"defaultKey"`
	KeyToTrackMappings StreamingPolicyContentKeyResponseArrayInput `pulumi:"keyToTrackMappings"`
}

func (StreamingPolicyContentKeysResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKeysResponse)(nil)).Elem()
}

func (i StreamingPolicyContentKeysResponseArgs) ToStreamingPolicyContentKeysResponseOutput() StreamingPolicyContentKeysResponseOutput {
	return i.ToStreamingPolicyContentKeysResponseOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeysResponseArgs) ToStreamingPolicyContentKeysResponseOutputWithContext(ctx context.Context) StreamingPolicyContentKeysResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeysResponseOutput)
}

func (i StreamingPolicyContentKeysResponseArgs) ToStreamingPolicyContentKeysResponsePtrOutput() StreamingPolicyContentKeysResponsePtrOutput {
	return i.ToStreamingPolicyContentKeysResponsePtrOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeysResponseArgs) ToStreamingPolicyContentKeysResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeysResponseOutput).ToStreamingPolicyContentKeysResponsePtrOutputWithContext(ctx)
}









type StreamingPolicyContentKeysResponsePtrInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeysResponsePtrOutput() StreamingPolicyContentKeysResponsePtrOutput
	ToStreamingPolicyContentKeysResponsePtrOutputWithContext(context.Context) StreamingPolicyContentKeysResponsePtrOutput
}

type streamingPolicyContentKeysResponsePtrType StreamingPolicyContentKeysResponseArgs

func StreamingPolicyContentKeysResponsePtr(v *StreamingPolicyContentKeysResponseArgs) StreamingPolicyContentKeysResponsePtrInput {
	return (*streamingPolicyContentKeysResponsePtrType)(v)
}

func (*streamingPolicyContentKeysResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyContentKeysResponse)(nil)).Elem()
}

func (i *streamingPolicyContentKeysResponsePtrType) ToStreamingPolicyContentKeysResponsePtrOutput() StreamingPolicyContentKeysResponsePtrOutput {
	return i.ToStreamingPolicyContentKeysResponsePtrOutputWithContext(context.Background())
}

func (i *streamingPolicyContentKeysResponsePtrType) ToStreamingPolicyContentKeysResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeysResponsePtrOutput)
}

type StreamingPolicyContentKeysResponseOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeysResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKeysResponse)(nil)).Elem()
}

func (o StreamingPolicyContentKeysResponseOutput) ToStreamingPolicyContentKeysResponseOutput() StreamingPolicyContentKeysResponseOutput {
	return o
}

func (o StreamingPolicyContentKeysResponseOutput) ToStreamingPolicyContentKeysResponseOutputWithContext(ctx context.Context) StreamingPolicyContentKeysResponseOutput {
	return o
}

func (o StreamingPolicyContentKeysResponseOutput) ToStreamingPolicyContentKeysResponsePtrOutput() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ToStreamingPolicyContentKeysResponsePtrOutputWithContext(context.Background())
}

func (o StreamingPolicyContentKeysResponseOutput) ToStreamingPolicyContentKeysResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyContentKeysResponse) *StreamingPolicyContentKeysResponse {
		return &v
	}).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o StreamingPolicyContentKeysResponseOutput) DefaultKey() DefaultKeyResponsePtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeysResponse) *DefaultKeyResponse { return v.DefaultKey }).(DefaultKeyResponsePtrOutput)
}

func (o StreamingPolicyContentKeysResponseOutput) KeyToTrackMappings() StreamingPolicyContentKeyResponseArrayOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeysResponse) []StreamingPolicyContentKeyResponse {
		return v.KeyToTrackMappings
	}).(StreamingPolicyContentKeyResponseArrayOutput)
}

type StreamingPolicyContentKeysResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeysResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyContentKeysResponse)(nil)).Elem()
}

func (o StreamingPolicyContentKeysResponsePtrOutput) ToStreamingPolicyContentKeysResponsePtrOutput() StreamingPolicyContentKeysResponsePtrOutput {
	return o
}

func (o StreamingPolicyContentKeysResponsePtrOutput) ToStreamingPolicyContentKeysResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysResponsePtrOutput {
	return o
}

func (o StreamingPolicyContentKeysResponsePtrOutput) Elem() StreamingPolicyContentKeysResponseOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeysResponse) StreamingPolicyContentKeysResponse {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyContentKeysResponse
		return ret
	}).(StreamingPolicyContentKeysResponseOutput)
}

func (o StreamingPolicyContentKeysResponsePtrOutput) DefaultKey() DefaultKeyResponsePtrOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeysResponse) *DefaultKeyResponse {
		if v == nil {
			return nil
		}
		return v.DefaultKey
	}).(DefaultKeyResponsePtrOutput)
}

func (o StreamingPolicyContentKeysResponsePtrOutput) KeyToTrackMappings() StreamingPolicyContentKeyResponseArrayOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeysResponse) []StreamingPolicyContentKeyResponse {
		if v == nil {
			return nil
		}
		return v.KeyToTrackMappings
	}).(StreamingPolicyContentKeyResponseArrayOutput)
}

type StreamingPolicyFairPlayConfiguration struct {
	AllowPersistentLicense              bool    `pulumi:"allowPersistentLicense"`
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
}





type StreamingPolicyFairPlayConfigurationInput interface {
	pulumi.Input

	ToStreamingPolicyFairPlayConfigurationOutput() StreamingPolicyFairPlayConfigurationOutput
	ToStreamingPolicyFairPlayConfigurationOutputWithContext(context.Context) StreamingPolicyFairPlayConfigurationOutput
}

type StreamingPolicyFairPlayConfigurationArgs struct {
	AllowPersistentLicense              pulumi.BoolInput      `pulumi:"allowPersistentLicense"`
	CustomLicenseAcquisitionUrlTemplate pulumi.StringPtrInput `pulumi:"customLicenseAcquisitionUrlTemplate"`
}

func (StreamingPolicyFairPlayConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyFairPlayConfiguration)(nil)).Elem()
}

func (i StreamingPolicyFairPlayConfigurationArgs) ToStreamingPolicyFairPlayConfigurationOutput() StreamingPolicyFairPlayConfigurationOutput {
	return i.ToStreamingPolicyFairPlayConfigurationOutputWithContext(context.Background())
}

func (i StreamingPolicyFairPlayConfigurationArgs) ToStreamingPolicyFairPlayConfigurationOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyFairPlayConfigurationOutput)
}

func (i StreamingPolicyFairPlayConfigurationArgs) ToStreamingPolicyFairPlayConfigurationPtrOutput() StreamingPolicyFairPlayConfigurationPtrOutput {
	return i.ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(context.Background())
}

func (i StreamingPolicyFairPlayConfigurationArgs) ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyFairPlayConfigurationOutput).ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(ctx)
}









type StreamingPolicyFairPlayConfigurationPtrInput interface {
	pulumi.Input

	ToStreamingPolicyFairPlayConfigurationPtrOutput() StreamingPolicyFairPlayConfigurationPtrOutput
	ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(context.Context) StreamingPolicyFairPlayConfigurationPtrOutput
}

type streamingPolicyFairPlayConfigurationPtrType StreamingPolicyFairPlayConfigurationArgs

func StreamingPolicyFairPlayConfigurationPtr(v *StreamingPolicyFairPlayConfigurationArgs) StreamingPolicyFairPlayConfigurationPtrInput {
	return (*streamingPolicyFairPlayConfigurationPtrType)(v)
}

func (*streamingPolicyFairPlayConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyFairPlayConfiguration)(nil)).Elem()
}

func (i *streamingPolicyFairPlayConfigurationPtrType) ToStreamingPolicyFairPlayConfigurationPtrOutput() StreamingPolicyFairPlayConfigurationPtrOutput {
	return i.ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(context.Background())
}

func (i *streamingPolicyFairPlayConfigurationPtrType) ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyFairPlayConfigurationPtrOutput)
}

type StreamingPolicyFairPlayConfigurationOutput struct{ *pulumi.OutputState }

func (StreamingPolicyFairPlayConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyFairPlayConfiguration)(nil)).Elem()
}

func (o StreamingPolicyFairPlayConfigurationOutput) ToStreamingPolicyFairPlayConfigurationOutput() StreamingPolicyFairPlayConfigurationOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationOutput) ToStreamingPolicyFairPlayConfigurationOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationOutput) ToStreamingPolicyFairPlayConfigurationPtrOutput() StreamingPolicyFairPlayConfigurationPtrOutput {
	return o.ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(context.Background())
}

func (o StreamingPolicyFairPlayConfigurationOutput) ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyFairPlayConfiguration) *StreamingPolicyFairPlayConfiguration {
		return &v
	}).(StreamingPolicyFairPlayConfigurationPtrOutput)
}

func (o StreamingPolicyFairPlayConfigurationOutput) AllowPersistentLicense() pulumi.BoolOutput {
	return o.ApplyT(func(v StreamingPolicyFairPlayConfiguration) bool { return v.AllowPersistentLicense }).(pulumi.BoolOutput)
}

func (o StreamingPolicyFairPlayConfigurationOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyFairPlayConfiguration) *string { return v.CustomLicenseAcquisitionUrlTemplate }).(pulumi.StringPtrOutput)
}

type StreamingPolicyFairPlayConfigurationPtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyFairPlayConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyFairPlayConfiguration)(nil)).Elem()
}

func (o StreamingPolicyFairPlayConfigurationPtrOutput) ToStreamingPolicyFairPlayConfigurationPtrOutput() StreamingPolicyFairPlayConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationPtrOutput) ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationPtrOutput) Elem() StreamingPolicyFairPlayConfigurationOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfiguration) StreamingPolicyFairPlayConfiguration {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyFairPlayConfiguration
		return ret
	}).(StreamingPolicyFairPlayConfigurationOutput)
}

func (o StreamingPolicyFairPlayConfigurationPtrOutput) AllowPersistentLicense() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfiguration) *bool {
		if v == nil {
			return nil
		}
		return &v.AllowPersistentLicense
	}).(pulumi.BoolPtrOutput)
}

func (o StreamingPolicyFairPlayConfigurationPtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyFairPlayConfigurationResponse struct {
	AllowPersistentLicense              bool    `pulumi:"allowPersistentLicense"`
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
}





type StreamingPolicyFairPlayConfigurationResponseInput interface {
	pulumi.Input

	ToStreamingPolicyFairPlayConfigurationResponseOutput() StreamingPolicyFairPlayConfigurationResponseOutput
	ToStreamingPolicyFairPlayConfigurationResponseOutputWithContext(context.Context) StreamingPolicyFairPlayConfigurationResponseOutput
}

type StreamingPolicyFairPlayConfigurationResponseArgs struct {
	AllowPersistentLicense              pulumi.BoolInput      `pulumi:"allowPersistentLicense"`
	CustomLicenseAcquisitionUrlTemplate pulumi.StringPtrInput `pulumi:"customLicenseAcquisitionUrlTemplate"`
}

func (StreamingPolicyFairPlayConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyFairPlayConfigurationResponse)(nil)).Elem()
}

func (i StreamingPolicyFairPlayConfigurationResponseArgs) ToStreamingPolicyFairPlayConfigurationResponseOutput() StreamingPolicyFairPlayConfigurationResponseOutput {
	return i.ToStreamingPolicyFairPlayConfigurationResponseOutputWithContext(context.Background())
}

func (i StreamingPolicyFairPlayConfigurationResponseArgs) ToStreamingPolicyFairPlayConfigurationResponseOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyFairPlayConfigurationResponseOutput)
}

func (i StreamingPolicyFairPlayConfigurationResponseArgs) ToStreamingPolicyFairPlayConfigurationResponsePtrOutput() StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return i.ToStreamingPolicyFairPlayConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i StreamingPolicyFairPlayConfigurationResponseArgs) ToStreamingPolicyFairPlayConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyFairPlayConfigurationResponseOutput).ToStreamingPolicyFairPlayConfigurationResponsePtrOutputWithContext(ctx)
}









type StreamingPolicyFairPlayConfigurationResponsePtrInput interface {
	pulumi.Input

	ToStreamingPolicyFairPlayConfigurationResponsePtrOutput() StreamingPolicyFairPlayConfigurationResponsePtrOutput
	ToStreamingPolicyFairPlayConfigurationResponsePtrOutputWithContext(context.Context) StreamingPolicyFairPlayConfigurationResponsePtrOutput
}

type streamingPolicyFairPlayConfigurationResponsePtrType StreamingPolicyFairPlayConfigurationResponseArgs

func StreamingPolicyFairPlayConfigurationResponsePtr(v *StreamingPolicyFairPlayConfigurationResponseArgs) StreamingPolicyFairPlayConfigurationResponsePtrInput {
	return (*streamingPolicyFairPlayConfigurationResponsePtrType)(v)
}

func (*streamingPolicyFairPlayConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyFairPlayConfigurationResponse)(nil)).Elem()
}

func (i *streamingPolicyFairPlayConfigurationResponsePtrType) ToStreamingPolicyFairPlayConfigurationResponsePtrOutput() StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return i.ToStreamingPolicyFairPlayConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *streamingPolicyFairPlayConfigurationResponsePtrType) ToStreamingPolicyFairPlayConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyFairPlayConfigurationResponsePtrOutput)
}

type StreamingPolicyFairPlayConfigurationResponseOutput struct{ *pulumi.OutputState }

func (StreamingPolicyFairPlayConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyFairPlayConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyFairPlayConfigurationResponseOutput) ToStreamingPolicyFairPlayConfigurationResponseOutput() StreamingPolicyFairPlayConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationResponseOutput) ToStreamingPolicyFairPlayConfigurationResponseOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationResponseOutput) ToStreamingPolicyFairPlayConfigurationResponsePtrOutput() StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return o.ToStreamingPolicyFairPlayConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o StreamingPolicyFairPlayConfigurationResponseOutput) ToStreamingPolicyFairPlayConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyFairPlayConfigurationResponse) *StreamingPolicyFairPlayConfigurationResponse {
		return &v
	}).(StreamingPolicyFairPlayConfigurationResponsePtrOutput)
}

func (o StreamingPolicyFairPlayConfigurationResponseOutput) AllowPersistentLicense() pulumi.BoolOutput {
	return o.ApplyT(func(v StreamingPolicyFairPlayConfigurationResponse) bool { return v.AllowPersistentLicense }).(pulumi.BoolOutput)
}

func (o StreamingPolicyFairPlayConfigurationResponseOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyFairPlayConfigurationResponse) *string {
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyFairPlayConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyFairPlayConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyFairPlayConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyFairPlayConfigurationResponsePtrOutput) ToStreamingPolicyFairPlayConfigurationResponsePtrOutput() StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationResponsePtrOutput) ToStreamingPolicyFairPlayConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationResponsePtrOutput) Elem() StreamingPolicyFairPlayConfigurationResponseOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfigurationResponse) StreamingPolicyFairPlayConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyFairPlayConfigurationResponse
		return ret
	}).(StreamingPolicyFairPlayConfigurationResponseOutput)
}

func (o StreamingPolicyFairPlayConfigurationResponsePtrOutput) AllowPersistentLicense() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.AllowPersistentLicense
	}).(pulumi.BoolPtrOutput)
}

func (o StreamingPolicyFairPlayConfigurationResponsePtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyPlayReadyConfiguration struct {
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
	PlayReadyCustomAttributes           *string `pulumi:"playReadyCustomAttributes"`
}





type StreamingPolicyPlayReadyConfigurationInput interface {
	pulumi.Input

	ToStreamingPolicyPlayReadyConfigurationOutput() StreamingPolicyPlayReadyConfigurationOutput
	ToStreamingPolicyPlayReadyConfigurationOutputWithContext(context.Context) StreamingPolicyPlayReadyConfigurationOutput
}

type StreamingPolicyPlayReadyConfigurationArgs struct {
	CustomLicenseAcquisitionUrlTemplate pulumi.StringPtrInput `pulumi:"customLicenseAcquisitionUrlTemplate"`
	PlayReadyCustomAttributes           pulumi.StringPtrInput `pulumi:"playReadyCustomAttributes"`
}

func (StreamingPolicyPlayReadyConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyPlayReadyConfiguration)(nil)).Elem()
}

func (i StreamingPolicyPlayReadyConfigurationArgs) ToStreamingPolicyPlayReadyConfigurationOutput() StreamingPolicyPlayReadyConfigurationOutput {
	return i.ToStreamingPolicyPlayReadyConfigurationOutputWithContext(context.Background())
}

func (i StreamingPolicyPlayReadyConfigurationArgs) ToStreamingPolicyPlayReadyConfigurationOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyPlayReadyConfigurationOutput)
}

func (i StreamingPolicyPlayReadyConfigurationArgs) ToStreamingPolicyPlayReadyConfigurationPtrOutput() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return i.ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(context.Background())
}

func (i StreamingPolicyPlayReadyConfigurationArgs) ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyPlayReadyConfigurationOutput).ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(ctx)
}









type StreamingPolicyPlayReadyConfigurationPtrInput interface {
	pulumi.Input

	ToStreamingPolicyPlayReadyConfigurationPtrOutput() StreamingPolicyPlayReadyConfigurationPtrOutput
	ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(context.Context) StreamingPolicyPlayReadyConfigurationPtrOutput
}

type streamingPolicyPlayReadyConfigurationPtrType StreamingPolicyPlayReadyConfigurationArgs

func StreamingPolicyPlayReadyConfigurationPtr(v *StreamingPolicyPlayReadyConfigurationArgs) StreamingPolicyPlayReadyConfigurationPtrInput {
	return (*streamingPolicyPlayReadyConfigurationPtrType)(v)
}

func (*streamingPolicyPlayReadyConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyPlayReadyConfiguration)(nil)).Elem()
}

func (i *streamingPolicyPlayReadyConfigurationPtrType) ToStreamingPolicyPlayReadyConfigurationPtrOutput() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return i.ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(context.Background())
}

func (i *streamingPolicyPlayReadyConfigurationPtrType) ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

type StreamingPolicyPlayReadyConfigurationOutput struct{ *pulumi.OutputState }

func (StreamingPolicyPlayReadyConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyPlayReadyConfiguration)(nil)).Elem()
}

func (o StreamingPolicyPlayReadyConfigurationOutput) ToStreamingPolicyPlayReadyConfigurationOutput() StreamingPolicyPlayReadyConfigurationOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationOutput) ToStreamingPolicyPlayReadyConfigurationOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationOutput) ToStreamingPolicyPlayReadyConfigurationPtrOutput() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(context.Background())
}

func (o StreamingPolicyPlayReadyConfigurationOutput) ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyPlayReadyConfiguration) *StreamingPolicyPlayReadyConfiguration {
		return &v
	}).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

func (o StreamingPolicyPlayReadyConfigurationOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyPlayReadyConfiguration) *string { return v.CustomLicenseAcquisitionUrlTemplate }).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyPlayReadyConfigurationOutput) PlayReadyCustomAttributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyPlayReadyConfiguration) *string { return v.PlayReadyCustomAttributes }).(pulumi.StringPtrOutput)
}

type StreamingPolicyPlayReadyConfigurationPtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyPlayReadyConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyPlayReadyConfiguration)(nil)).Elem()
}

func (o StreamingPolicyPlayReadyConfigurationPtrOutput) ToStreamingPolicyPlayReadyConfigurationPtrOutput() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationPtrOutput) ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationPtrOutput) Elem() StreamingPolicyPlayReadyConfigurationOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfiguration) StreamingPolicyPlayReadyConfiguration {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyPlayReadyConfiguration
		return ret
	}).(StreamingPolicyPlayReadyConfigurationOutput)
}

func (o StreamingPolicyPlayReadyConfigurationPtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyPlayReadyConfigurationPtrOutput) PlayReadyCustomAttributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PlayReadyCustomAttributes
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyPlayReadyConfigurationResponse struct {
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
	PlayReadyCustomAttributes           *string `pulumi:"playReadyCustomAttributes"`
}





type StreamingPolicyPlayReadyConfigurationResponseInput interface {
	pulumi.Input

	ToStreamingPolicyPlayReadyConfigurationResponseOutput() StreamingPolicyPlayReadyConfigurationResponseOutput
	ToStreamingPolicyPlayReadyConfigurationResponseOutputWithContext(context.Context) StreamingPolicyPlayReadyConfigurationResponseOutput
}

type StreamingPolicyPlayReadyConfigurationResponseArgs struct {
	CustomLicenseAcquisitionUrlTemplate pulumi.StringPtrInput `pulumi:"customLicenseAcquisitionUrlTemplate"`
	PlayReadyCustomAttributes           pulumi.StringPtrInput `pulumi:"playReadyCustomAttributes"`
}

func (StreamingPolicyPlayReadyConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyPlayReadyConfigurationResponse)(nil)).Elem()
}

func (i StreamingPolicyPlayReadyConfigurationResponseArgs) ToStreamingPolicyPlayReadyConfigurationResponseOutput() StreamingPolicyPlayReadyConfigurationResponseOutput {
	return i.ToStreamingPolicyPlayReadyConfigurationResponseOutputWithContext(context.Background())
}

func (i StreamingPolicyPlayReadyConfigurationResponseArgs) ToStreamingPolicyPlayReadyConfigurationResponseOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyPlayReadyConfigurationResponseOutput)
}

func (i StreamingPolicyPlayReadyConfigurationResponseArgs) ToStreamingPolicyPlayReadyConfigurationResponsePtrOutput() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return i.ToStreamingPolicyPlayReadyConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i StreamingPolicyPlayReadyConfigurationResponseArgs) ToStreamingPolicyPlayReadyConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyPlayReadyConfigurationResponseOutput).ToStreamingPolicyPlayReadyConfigurationResponsePtrOutputWithContext(ctx)
}









type StreamingPolicyPlayReadyConfigurationResponsePtrInput interface {
	pulumi.Input

	ToStreamingPolicyPlayReadyConfigurationResponsePtrOutput() StreamingPolicyPlayReadyConfigurationResponsePtrOutput
	ToStreamingPolicyPlayReadyConfigurationResponsePtrOutputWithContext(context.Context) StreamingPolicyPlayReadyConfigurationResponsePtrOutput
}

type streamingPolicyPlayReadyConfigurationResponsePtrType StreamingPolicyPlayReadyConfigurationResponseArgs

func StreamingPolicyPlayReadyConfigurationResponsePtr(v *StreamingPolicyPlayReadyConfigurationResponseArgs) StreamingPolicyPlayReadyConfigurationResponsePtrInput {
	return (*streamingPolicyPlayReadyConfigurationResponsePtrType)(v)
}

func (*streamingPolicyPlayReadyConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyPlayReadyConfigurationResponse)(nil)).Elem()
}

func (i *streamingPolicyPlayReadyConfigurationResponsePtrType) ToStreamingPolicyPlayReadyConfigurationResponsePtrOutput() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return i.ToStreamingPolicyPlayReadyConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *streamingPolicyPlayReadyConfigurationResponsePtrType) ToStreamingPolicyPlayReadyConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyPlayReadyConfigurationResponsePtrOutput)
}

type StreamingPolicyPlayReadyConfigurationResponseOutput struct{ *pulumi.OutputState }

func (StreamingPolicyPlayReadyConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyPlayReadyConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyPlayReadyConfigurationResponseOutput) ToStreamingPolicyPlayReadyConfigurationResponseOutput() StreamingPolicyPlayReadyConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationResponseOutput) ToStreamingPolicyPlayReadyConfigurationResponseOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationResponseOutput) ToStreamingPolicyPlayReadyConfigurationResponsePtrOutput() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o.ToStreamingPolicyPlayReadyConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o StreamingPolicyPlayReadyConfigurationResponseOutput) ToStreamingPolicyPlayReadyConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyPlayReadyConfigurationResponse) *StreamingPolicyPlayReadyConfigurationResponse {
		return &v
	}).(StreamingPolicyPlayReadyConfigurationResponsePtrOutput)
}

func (o StreamingPolicyPlayReadyConfigurationResponseOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyPlayReadyConfigurationResponse) *string {
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyPlayReadyConfigurationResponseOutput) PlayReadyCustomAttributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyPlayReadyConfigurationResponse) *string { return v.PlayReadyCustomAttributes }).(pulumi.StringPtrOutput)
}

type StreamingPolicyPlayReadyConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyPlayReadyConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyPlayReadyConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyPlayReadyConfigurationResponsePtrOutput) ToStreamingPolicyPlayReadyConfigurationResponsePtrOutput() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationResponsePtrOutput) ToStreamingPolicyPlayReadyConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationResponsePtrOutput) Elem() StreamingPolicyPlayReadyConfigurationResponseOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfigurationResponse) StreamingPolicyPlayReadyConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyPlayReadyConfigurationResponse
		return ret
	}).(StreamingPolicyPlayReadyConfigurationResponseOutput)
}

func (o StreamingPolicyPlayReadyConfigurationResponsePtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyPlayReadyConfigurationResponsePtrOutput) PlayReadyCustomAttributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PlayReadyCustomAttributes
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyWidevineConfiguration struct {
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
}





type StreamingPolicyWidevineConfigurationInput interface {
	pulumi.Input

	ToStreamingPolicyWidevineConfigurationOutput() StreamingPolicyWidevineConfigurationOutput
	ToStreamingPolicyWidevineConfigurationOutputWithContext(context.Context) StreamingPolicyWidevineConfigurationOutput
}

type StreamingPolicyWidevineConfigurationArgs struct {
	CustomLicenseAcquisitionUrlTemplate pulumi.StringPtrInput `pulumi:"customLicenseAcquisitionUrlTemplate"`
}

func (StreamingPolicyWidevineConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyWidevineConfiguration)(nil)).Elem()
}

func (i StreamingPolicyWidevineConfigurationArgs) ToStreamingPolicyWidevineConfigurationOutput() StreamingPolicyWidevineConfigurationOutput {
	return i.ToStreamingPolicyWidevineConfigurationOutputWithContext(context.Background())
}

func (i StreamingPolicyWidevineConfigurationArgs) ToStreamingPolicyWidevineConfigurationOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyWidevineConfigurationOutput)
}

func (i StreamingPolicyWidevineConfigurationArgs) ToStreamingPolicyWidevineConfigurationPtrOutput() StreamingPolicyWidevineConfigurationPtrOutput {
	return i.ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(context.Background())
}

func (i StreamingPolicyWidevineConfigurationArgs) ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyWidevineConfigurationOutput).ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(ctx)
}









type StreamingPolicyWidevineConfigurationPtrInput interface {
	pulumi.Input

	ToStreamingPolicyWidevineConfigurationPtrOutput() StreamingPolicyWidevineConfigurationPtrOutput
	ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(context.Context) StreamingPolicyWidevineConfigurationPtrOutput
}

type streamingPolicyWidevineConfigurationPtrType StreamingPolicyWidevineConfigurationArgs

func StreamingPolicyWidevineConfigurationPtr(v *StreamingPolicyWidevineConfigurationArgs) StreamingPolicyWidevineConfigurationPtrInput {
	return (*streamingPolicyWidevineConfigurationPtrType)(v)
}

func (*streamingPolicyWidevineConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyWidevineConfiguration)(nil)).Elem()
}

func (i *streamingPolicyWidevineConfigurationPtrType) ToStreamingPolicyWidevineConfigurationPtrOutput() StreamingPolicyWidevineConfigurationPtrOutput {
	return i.ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(context.Background())
}

func (i *streamingPolicyWidevineConfigurationPtrType) ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyWidevineConfigurationPtrOutput)
}

type StreamingPolicyWidevineConfigurationOutput struct{ *pulumi.OutputState }

func (StreamingPolicyWidevineConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyWidevineConfiguration)(nil)).Elem()
}

func (o StreamingPolicyWidevineConfigurationOutput) ToStreamingPolicyWidevineConfigurationOutput() StreamingPolicyWidevineConfigurationOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationOutput) ToStreamingPolicyWidevineConfigurationOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationOutput) ToStreamingPolicyWidevineConfigurationPtrOutput() StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(context.Background())
}

func (o StreamingPolicyWidevineConfigurationOutput) ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyWidevineConfiguration) *StreamingPolicyWidevineConfiguration {
		return &v
	}).(StreamingPolicyWidevineConfigurationPtrOutput)
}

func (o StreamingPolicyWidevineConfigurationOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyWidevineConfiguration) *string { return v.CustomLicenseAcquisitionUrlTemplate }).(pulumi.StringPtrOutput)
}

type StreamingPolicyWidevineConfigurationPtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyWidevineConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyWidevineConfiguration)(nil)).Elem()
}

func (o StreamingPolicyWidevineConfigurationPtrOutput) ToStreamingPolicyWidevineConfigurationPtrOutput() StreamingPolicyWidevineConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationPtrOutput) ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationPtrOutput) Elem() StreamingPolicyWidevineConfigurationOutput {
	return o.ApplyT(func(v *StreamingPolicyWidevineConfiguration) StreamingPolicyWidevineConfiguration {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyWidevineConfiguration
		return ret
	}).(StreamingPolicyWidevineConfigurationOutput)
}

func (o StreamingPolicyWidevineConfigurationPtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyWidevineConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyWidevineConfigurationResponse struct {
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
}





type StreamingPolicyWidevineConfigurationResponseInput interface {
	pulumi.Input

	ToStreamingPolicyWidevineConfigurationResponseOutput() StreamingPolicyWidevineConfigurationResponseOutput
	ToStreamingPolicyWidevineConfigurationResponseOutputWithContext(context.Context) StreamingPolicyWidevineConfigurationResponseOutput
}

type StreamingPolicyWidevineConfigurationResponseArgs struct {
	CustomLicenseAcquisitionUrlTemplate pulumi.StringPtrInput `pulumi:"customLicenseAcquisitionUrlTemplate"`
}

func (StreamingPolicyWidevineConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyWidevineConfigurationResponse)(nil)).Elem()
}

func (i StreamingPolicyWidevineConfigurationResponseArgs) ToStreamingPolicyWidevineConfigurationResponseOutput() StreamingPolicyWidevineConfigurationResponseOutput {
	return i.ToStreamingPolicyWidevineConfigurationResponseOutputWithContext(context.Background())
}

func (i StreamingPolicyWidevineConfigurationResponseArgs) ToStreamingPolicyWidevineConfigurationResponseOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyWidevineConfigurationResponseOutput)
}

func (i StreamingPolicyWidevineConfigurationResponseArgs) ToStreamingPolicyWidevineConfigurationResponsePtrOutput() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return i.ToStreamingPolicyWidevineConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i StreamingPolicyWidevineConfigurationResponseArgs) ToStreamingPolicyWidevineConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyWidevineConfigurationResponseOutput).ToStreamingPolicyWidevineConfigurationResponsePtrOutputWithContext(ctx)
}









type StreamingPolicyWidevineConfigurationResponsePtrInput interface {
	pulumi.Input

	ToStreamingPolicyWidevineConfigurationResponsePtrOutput() StreamingPolicyWidevineConfigurationResponsePtrOutput
	ToStreamingPolicyWidevineConfigurationResponsePtrOutputWithContext(context.Context) StreamingPolicyWidevineConfigurationResponsePtrOutput
}

type streamingPolicyWidevineConfigurationResponsePtrType StreamingPolicyWidevineConfigurationResponseArgs

func StreamingPolicyWidevineConfigurationResponsePtr(v *StreamingPolicyWidevineConfigurationResponseArgs) StreamingPolicyWidevineConfigurationResponsePtrInput {
	return (*streamingPolicyWidevineConfigurationResponsePtrType)(v)
}

func (*streamingPolicyWidevineConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyWidevineConfigurationResponse)(nil)).Elem()
}

func (i *streamingPolicyWidevineConfigurationResponsePtrType) ToStreamingPolicyWidevineConfigurationResponsePtrOutput() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return i.ToStreamingPolicyWidevineConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *streamingPolicyWidevineConfigurationResponsePtrType) ToStreamingPolicyWidevineConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyWidevineConfigurationResponsePtrOutput)
}

type StreamingPolicyWidevineConfigurationResponseOutput struct{ *pulumi.OutputState }

func (StreamingPolicyWidevineConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyWidevineConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyWidevineConfigurationResponseOutput) ToStreamingPolicyWidevineConfigurationResponseOutput() StreamingPolicyWidevineConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationResponseOutput) ToStreamingPolicyWidevineConfigurationResponseOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationResponseOutput) ToStreamingPolicyWidevineConfigurationResponsePtrOutput() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o.ToStreamingPolicyWidevineConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o StreamingPolicyWidevineConfigurationResponseOutput) ToStreamingPolicyWidevineConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyWidevineConfigurationResponse) *StreamingPolicyWidevineConfigurationResponse {
		return &v
	}).(StreamingPolicyWidevineConfigurationResponsePtrOutput)
}

func (o StreamingPolicyWidevineConfigurationResponseOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyWidevineConfigurationResponse) *string {
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyWidevineConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyWidevineConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyWidevineConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyWidevineConfigurationResponsePtrOutput) ToStreamingPolicyWidevineConfigurationResponsePtrOutput() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationResponsePtrOutput) ToStreamingPolicyWidevineConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationResponsePtrOutput) Elem() StreamingPolicyWidevineConfigurationResponseOutput {
	return o.ApplyT(func(v *StreamingPolicyWidevineConfigurationResponse) StreamingPolicyWidevineConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyWidevineConfigurationResponse
		return ret
	}).(StreamingPolicyWidevineConfigurationResponseOutput)
}

func (o StreamingPolicyWidevineConfigurationResponsePtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyWidevineConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
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

type TrackPropertyCondition struct {
	Operation string  `pulumi:"operation"`
	Property  string  `pulumi:"property"`
	Value     *string `pulumi:"value"`
}





type TrackPropertyConditionInput interface {
	pulumi.Input

	ToTrackPropertyConditionOutput() TrackPropertyConditionOutput
	ToTrackPropertyConditionOutputWithContext(context.Context) TrackPropertyConditionOutput
}

type TrackPropertyConditionArgs struct {
	Operation pulumi.StringInput    `pulumi:"operation"`
	Property  pulumi.StringInput    `pulumi:"property"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (TrackPropertyConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackPropertyCondition)(nil)).Elem()
}

func (i TrackPropertyConditionArgs) ToTrackPropertyConditionOutput() TrackPropertyConditionOutput {
	return i.ToTrackPropertyConditionOutputWithContext(context.Background())
}

func (i TrackPropertyConditionArgs) ToTrackPropertyConditionOutputWithContext(ctx context.Context) TrackPropertyConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackPropertyConditionOutput)
}





type TrackPropertyConditionArrayInput interface {
	pulumi.Input

	ToTrackPropertyConditionArrayOutput() TrackPropertyConditionArrayOutput
	ToTrackPropertyConditionArrayOutputWithContext(context.Context) TrackPropertyConditionArrayOutput
}

type TrackPropertyConditionArray []TrackPropertyConditionInput

func (TrackPropertyConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackPropertyCondition)(nil)).Elem()
}

func (i TrackPropertyConditionArray) ToTrackPropertyConditionArrayOutput() TrackPropertyConditionArrayOutput {
	return i.ToTrackPropertyConditionArrayOutputWithContext(context.Background())
}

func (i TrackPropertyConditionArray) ToTrackPropertyConditionArrayOutputWithContext(ctx context.Context) TrackPropertyConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackPropertyConditionArrayOutput)
}

type TrackPropertyConditionOutput struct{ *pulumi.OutputState }

func (TrackPropertyConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackPropertyCondition)(nil)).Elem()
}

func (o TrackPropertyConditionOutput) ToTrackPropertyConditionOutput() TrackPropertyConditionOutput {
	return o
}

func (o TrackPropertyConditionOutput) ToTrackPropertyConditionOutputWithContext(ctx context.Context) TrackPropertyConditionOutput {
	return o
}

func (o TrackPropertyConditionOutput) Operation() pulumi.StringOutput {
	return o.ApplyT(func(v TrackPropertyCondition) string { return v.Operation }).(pulumi.StringOutput)
}

func (o TrackPropertyConditionOutput) Property() pulumi.StringOutput {
	return o.ApplyT(func(v TrackPropertyCondition) string { return v.Property }).(pulumi.StringOutput)
}

func (o TrackPropertyConditionOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackPropertyCondition) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type TrackPropertyConditionArrayOutput struct{ *pulumi.OutputState }

func (TrackPropertyConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackPropertyCondition)(nil)).Elem()
}

func (o TrackPropertyConditionArrayOutput) ToTrackPropertyConditionArrayOutput() TrackPropertyConditionArrayOutput {
	return o
}

func (o TrackPropertyConditionArrayOutput) ToTrackPropertyConditionArrayOutputWithContext(ctx context.Context) TrackPropertyConditionArrayOutput {
	return o
}

func (o TrackPropertyConditionArrayOutput) Index(i pulumi.IntInput) TrackPropertyConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrackPropertyCondition {
		return vs[0].([]TrackPropertyCondition)[vs[1].(int)]
	}).(TrackPropertyConditionOutput)
}

type TrackPropertyConditionResponse struct {
	Operation string  `pulumi:"operation"`
	Property  string  `pulumi:"property"`
	Value     *string `pulumi:"value"`
}





type TrackPropertyConditionResponseInput interface {
	pulumi.Input

	ToTrackPropertyConditionResponseOutput() TrackPropertyConditionResponseOutput
	ToTrackPropertyConditionResponseOutputWithContext(context.Context) TrackPropertyConditionResponseOutput
}

type TrackPropertyConditionResponseArgs struct {
	Operation pulumi.StringInput    `pulumi:"operation"`
	Property  pulumi.StringInput    `pulumi:"property"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (TrackPropertyConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackPropertyConditionResponse)(nil)).Elem()
}

func (i TrackPropertyConditionResponseArgs) ToTrackPropertyConditionResponseOutput() TrackPropertyConditionResponseOutput {
	return i.ToTrackPropertyConditionResponseOutputWithContext(context.Background())
}

func (i TrackPropertyConditionResponseArgs) ToTrackPropertyConditionResponseOutputWithContext(ctx context.Context) TrackPropertyConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackPropertyConditionResponseOutput)
}





type TrackPropertyConditionResponseArrayInput interface {
	pulumi.Input

	ToTrackPropertyConditionResponseArrayOutput() TrackPropertyConditionResponseArrayOutput
	ToTrackPropertyConditionResponseArrayOutputWithContext(context.Context) TrackPropertyConditionResponseArrayOutput
}

type TrackPropertyConditionResponseArray []TrackPropertyConditionResponseInput

func (TrackPropertyConditionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackPropertyConditionResponse)(nil)).Elem()
}

func (i TrackPropertyConditionResponseArray) ToTrackPropertyConditionResponseArrayOutput() TrackPropertyConditionResponseArrayOutput {
	return i.ToTrackPropertyConditionResponseArrayOutputWithContext(context.Background())
}

func (i TrackPropertyConditionResponseArray) ToTrackPropertyConditionResponseArrayOutputWithContext(ctx context.Context) TrackPropertyConditionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackPropertyConditionResponseArrayOutput)
}

type TrackPropertyConditionResponseOutput struct{ *pulumi.OutputState }

func (TrackPropertyConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackPropertyConditionResponse)(nil)).Elem()
}

func (o TrackPropertyConditionResponseOutput) ToTrackPropertyConditionResponseOutput() TrackPropertyConditionResponseOutput {
	return o
}

func (o TrackPropertyConditionResponseOutput) ToTrackPropertyConditionResponseOutputWithContext(ctx context.Context) TrackPropertyConditionResponseOutput {
	return o
}

func (o TrackPropertyConditionResponseOutput) Operation() pulumi.StringOutput {
	return o.ApplyT(func(v TrackPropertyConditionResponse) string { return v.Operation }).(pulumi.StringOutput)
}

func (o TrackPropertyConditionResponseOutput) Property() pulumi.StringOutput {
	return o.ApplyT(func(v TrackPropertyConditionResponse) string { return v.Property }).(pulumi.StringOutput)
}

func (o TrackPropertyConditionResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackPropertyConditionResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type TrackPropertyConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (TrackPropertyConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackPropertyConditionResponse)(nil)).Elem()
}

func (o TrackPropertyConditionResponseArrayOutput) ToTrackPropertyConditionResponseArrayOutput() TrackPropertyConditionResponseArrayOutput {
	return o
}

func (o TrackPropertyConditionResponseArrayOutput) ToTrackPropertyConditionResponseArrayOutputWithContext(ctx context.Context) TrackPropertyConditionResponseArrayOutput {
	return o
}

func (o TrackPropertyConditionResponseArrayOutput) Index(i pulumi.IntInput) TrackPropertyConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrackPropertyConditionResponse {
		return vs[0].([]TrackPropertyConditionResponse)[vs[1].(int)]
	}).(TrackPropertyConditionResponseOutput)
}

type TrackSelection struct {
	TrackSelections []TrackPropertyCondition `pulumi:"trackSelections"`
}





type TrackSelectionInput interface {
	pulumi.Input

	ToTrackSelectionOutput() TrackSelectionOutput
	ToTrackSelectionOutputWithContext(context.Context) TrackSelectionOutput
}

type TrackSelectionArgs struct {
	TrackSelections TrackPropertyConditionArrayInput `pulumi:"trackSelections"`
}

func (TrackSelectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackSelection)(nil)).Elem()
}

func (i TrackSelectionArgs) ToTrackSelectionOutput() TrackSelectionOutput {
	return i.ToTrackSelectionOutputWithContext(context.Background())
}

func (i TrackSelectionArgs) ToTrackSelectionOutputWithContext(ctx context.Context) TrackSelectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackSelectionOutput)
}





type TrackSelectionArrayInput interface {
	pulumi.Input

	ToTrackSelectionArrayOutput() TrackSelectionArrayOutput
	ToTrackSelectionArrayOutputWithContext(context.Context) TrackSelectionArrayOutput
}

type TrackSelectionArray []TrackSelectionInput

func (TrackSelectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackSelection)(nil)).Elem()
}

func (i TrackSelectionArray) ToTrackSelectionArrayOutput() TrackSelectionArrayOutput {
	return i.ToTrackSelectionArrayOutputWithContext(context.Background())
}

func (i TrackSelectionArray) ToTrackSelectionArrayOutputWithContext(ctx context.Context) TrackSelectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackSelectionArrayOutput)
}

type TrackSelectionOutput struct{ *pulumi.OutputState }

func (TrackSelectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackSelection)(nil)).Elem()
}

func (o TrackSelectionOutput) ToTrackSelectionOutput() TrackSelectionOutput {
	return o
}

func (o TrackSelectionOutput) ToTrackSelectionOutputWithContext(ctx context.Context) TrackSelectionOutput {
	return o
}

func (o TrackSelectionOutput) TrackSelections() TrackPropertyConditionArrayOutput {
	return o.ApplyT(func(v TrackSelection) []TrackPropertyCondition { return v.TrackSelections }).(TrackPropertyConditionArrayOutput)
}

type TrackSelectionArrayOutput struct{ *pulumi.OutputState }

func (TrackSelectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackSelection)(nil)).Elem()
}

func (o TrackSelectionArrayOutput) ToTrackSelectionArrayOutput() TrackSelectionArrayOutput {
	return o
}

func (o TrackSelectionArrayOutput) ToTrackSelectionArrayOutputWithContext(ctx context.Context) TrackSelectionArrayOutput {
	return o
}

func (o TrackSelectionArrayOutput) Index(i pulumi.IntInput) TrackSelectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrackSelection {
		return vs[0].([]TrackSelection)[vs[1].(int)]
	}).(TrackSelectionOutput)
}

type TrackSelectionResponse struct {
	TrackSelections []TrackPropertyConditionResponse `pulumi:"trackSelections"`
}





type TrackSelectionResponseInput interface {
	pulumi.Input

	ToTrackSelectionResponseOutput() TrackSelectionResponseOutput
	ToTrackSelectionResponseOutputWithContext(context.Context) TrackSelectionResponseOutput
}

type TrackSelectionResponseArgs struct {
	TrackSelections TrackPropertyConditionResponseArrayInput `pulumi:"trackSelections"`
}

func (TrackSelectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackSelectionResponse)(nil)).Elem()
}

func (i TrackSelectionResponseArgs) ToTrackSelectionResponseOutput() TrackSelectionResponseOutput {
	return i.ToTrackSelectionResponseOutputWithContext(context.Background())
}

func (i TrackSelectionResponseArgs) ToTrackSelectionResponseOutputWithContext(ctx context.Context) TrackSelectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackSelectionResponseOutput)
}





type TrackSelectionResponseArrayInput interface {
	pulumi.Input

	ToTrackSelectionResponseArrayOutput() TrackSelectionResponseArrayOutput
	ToTrackSelectionResponseArrayOutputWithContext(context.Context) TrackSelectionResponseArrayOutput
}

type TrackSelectionResponseArray []TrackSelectionResponseInput

func (TrackSelectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackSelectionResponse)(nil)).Elem()
}

func (i TrackSelectionResponseArray) ToTrackSelectionResponseArrayOutput() TrackSelectionResponseArrayOutput {
	return i.ToTrackSelectionResponseArrayOutputWithContext(context.Background())
}

func (i TrackSelectionResponseArray) ToTrackSelectionResponseArrayOutputWithContext(ctx context.Context) TrackSelectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackSelectionResponseArrayOutput)
}

type TrackSelectionResponseOutput struct{ *pulumi.OutputState }

func (TrackSelectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackSelectionResponse)(nil)).Elem()
}

func (o TrackSelectionResponseOutput) ToTrackSelectionResponseOutput() TrackSelectionResponseOutput {
	return o
}

func (o TrackSelectionResponseOutput) ToTrackSelectionResponseOutputWithContext(ctx context.Context) TrackSelectionResponseOutput {
	return o
}

func (o TrackSelectionResponseOutput) TrackSelections() TrackPropertyConditionResponseArrayOutput {
	return o.ApplyT(func(v TrackSelectionResponse) []TrackPropertyConditionResponse { return v.TrackSelections }).(TrackPropertyConditionResponseArrayOutput)
}

type TrackSelectionResponseArrayOutput struct{ *pulumi.OutputState }

func (TrackSelectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackSelectionResponse)(nil)).Elem()
}

func (o TrackSelectionResponseArrayOutput) ToTrackSelectionResponseArrayOutput() TrackSelectionResponseArrayOutput {
	return o
}

func (o TrackSelectionResponseArrayOutput) ToTrackSelectionResponseArrayOutputWithContext(ctx context.Context) TrackSelectionResponseArrayOutput {
	return o
}

func (o TrackSelectionResponseArrayOutput) Index(i pulumi.IntInput) TrackSelectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrackSelectionResponse {
		return vs[0].([]TrackSelectionResponse)[vs[1].(int)]
	}).(TrackSelectionResponseOutput)
}

type TransformOutputType struct {
	OnError          *string     `pulumi:"onError"`
	Preset           interface{} `pulumi:"preset"`
	RelativePriority *string     `pulumi:"relativePriority"`
}





type TransformOutputTypeInput interface {
	pulumi.Input

	ToTransformOutputTypeOutput() TransformOutputTypeOutput
	ToTransformOutputTypeOutputWithContext(context.Context) TransformOutputTypeOutput
}

type TransformOutputTypeArgs struct {
	OnError          pulumi.StringPtrInput `pulumi:"onError"`
	Preset           pulumi.Input          `pulumi:"preset"`
	RelativePriority pulumi.StringPtrInput `pulumi:"relativePriority"`
}

func (TransformOutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformOutputType)(nil)).Elem()
}

func (i TransformOutputTypeArgs) ToTransformOutputTypeOutput() TransformOutputTypeOutput {
	return i.ToTransformOutputTypeOutputWithContext(context.Background())
}

func (i TransformOutputTypeArgs) ToTransformOutputTypeOutputWithContext(ctx context.Context) TransformOutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformOutputTypeOutput)
}





type TransformOutputTypeArrayInput interface {
	pulumi.Input

	ToTransformOutputTypeArrayOutput() TransformOutputTypeArrayOutput
	ToTransformOutputTypeArrayOutputWithContext(context.Context) TransformOutputTypeArrayOutput
}

type TransformOutputTypeArray []TransformOutputTypeInput

func (TransformOutputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransformOutputType)(nil)).Elem()
}

func (i TransformOutputTypeArray) ToTransformOutputTypeArrayOutput() TransformOutputTypeArrayOutput {
	return i.ToTransformOutputTypeArrayOutputWithContext(context.Background())
}

func (i TransformOutputTypeArray) ToTransformOutputTypeArrayOutputWithContext(ctx context.Context) TransformOutputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformOutputTypeArrayOutput)
}

type TransformOutputTypeOutput struct{ *pulumi.OutputState }

func (TransformOutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformOutputType)(nil)).Elem()
}

func (o TransformOutputTypeOutput) ToTransformOutputTypeOutput() TransformOutputTypeOutput {
	return o
}

func (o TransformOutputTypeOutput) ToTransformOutputTypeOutputWithContext(ctx context.Context) TransformOutputTypeOutput {
	return o
}

func (o TransformOutputTypeOutput) OnError() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputType) *string { return v.OnError }).(pulumi.StringPtrOutput)
}

func (o TransformOutputTypeOutput) Preset() pulumi.AnyOutput {
	return o.ApplyT(func(v TransformOutputType) interface{} { return v.Preset }).(pulumi.AnyOutput)
}

func (o TransformOutputTypeOutput) RelativePriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputType) *string { return v.RelativePriority }).(pulumi.StringPtrOutput)
}

type TransformOutputTypeArrayOutput struct{ *pulumi.OutputState }

func (TransformOutputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransformOutputType)(nil)).Elem()
}

func (o TransformOutputTypeArrayOutput) ToTransformOutputTypeArrayOutput() TransformOutputTypeArrayOutput {
	return o
}

func (o TransformOutputTypeArrayOutput) ToTransformOutputTypeArrayOutputWithContext(ctx context.Context) TransformOutputTypeArrayOutput {
	return o
}

func (o TransformOutputTypeArrayOutput) Index(i pulumi.IntInput) TransformOutputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TransformOutputType {
		return vs[0].([]TransformOutputType)[vs[1].(int)]
	}).(TransformOutputTypeOutput)
}

type TransformOutputResponse struct {
	OnError          *string     `pulumi:"onError"`
	Preset           interface{} `pulumi:"preset"`
	RelativePriority *string     `pulumi:"relativePriority"`
}





type TransformOutputResponseInput interface {
	pulumi.Input

	ToTransformOutputResponseOutput() TransformOutputResponseOutput
	ToTransformOutputResponseOutputWithContext(context.Context) TransformOutputResponseOutput
}

type TransformOutputResponseArgs struct {
	OnError          pulumi.StringPtrInput `pulumi:"onError"`
	Preset           pulumi.Input          `pulumi:"preset"`
	RelativePriority pulumi.StringPtrInput `pulumi:"relativePriority"`
}

func (TransformOutputResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformOutputResponse)(nil)).Elem()
}

func (i TransformOutputResponseArgs) ToTransformOutputResponseOutput() TransformOutputResponseOutput {
	return i.ToTransformOutputResponseOutputWithContext(context.Background())
}

func (i TransformOutputResponseArgs) ToTransformOutputResponseOutputWithContext(ctx context.Context) TransformOutputResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformOutputResponseOutput)
}





type TransformOutputResponseArrayInput interface {
	pulumi.Input

	ToTransformOutputResponseArrayOutput() TransformOutputResponseArrayOutput
	ToTransformOutputResponseArrayOutputWithContext(context.Context) TransformOutputResponseArrayOutput
}

type TransformOutputResponseArray []TransformOutputResponseInput

func (TransformOutputResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransformOutputResponse)(nil)).Elem()
}

func (i TransformOutputResponseArray) ToTransformOutputResponseArrayOutput() TransformOutputResponseArrayOutput {
	return i.ToTransformOutputResponseArrayOutputWithContext(context.Background())
}

func (i TransformOutputResponseArray) ToTransformOutputResponseArrayOutputWithContext(ctx context.Context) TransformOutputResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformOutputResponseArrayOutput)
}

type TransformOutputResponseOutput struct{ *pulumi.OutputState }

func (TransformOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformOutputResponse)(nil)).Elem()
}

func (o TransformOutputResponseOutput) ToTransformOutputResponseOutput() TransformOutputResponseOutput {
	return o
}

func (o TransformOutputResponseOutput) ToTransformOutputResponseOutputWithContext(ctx context.Context) TransformOutputResponseOutput {
	return o
}

func (o TransformOutputResponseOutput) OnError() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputResponse) *string { return v.OnError }).(pulumi.StringPtrOutput)
}

func (o TransformOutputResponseOutput) Preset() pulumi.AnyOutput {
	return o.ApplyT(func(v TransformOutputResponse) interface{} { return v.Preset }).(pulumi.AnyOutput)
}

func (o TransformOutputResponseOutput) RelativePriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputResponse) *string { return v.RelativePriority }).(pulumi.StringPtrOutput)
}

type TransformOutputResponseArrayOutput struct{ *pulumi.OutputState }

func (TransformOutputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransformOutputResponse)(nil)).Elem()
}

func (o TransformOutputResponseArrayOutput) ToTransformOutputResponseArrayOutput() TransformOutputResponseArrayOutput {
	return o
}

func (o TransformOutputResponseArrayOutput) ToTransformOutputResponseArrayOutputWithContext(ctx context.Context) TransformOutputResponseArrayOutput {
	return o
}

func (o TransformOutputResponseArrayOutput) Index(i pulumi.IntInput) TransformOutputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TransformOutputResponse {
		return vs[0].([]TransformOutputResponse)[vs[1].(int)]
	}).(TransformOutputResponseOutput)
}

type TransportStreamFormat struct {
	FilenamePattern string       `pulumi:"filenamePattern"`
	OdataType       string       `pulumi:"odataType"`
	OutputFiles     []OutputFile `pulumi:"outputFiles"`
}





type TransportStreamFormatInput interface {
	pulumi.Input

	ToTransportStreamFormatOutput() TransportStreamFormatOutput
	ToTransportStreamFormatOutputWithContext(context.Context) TransportStreamFormatOutput
}

type TransportStreamFormatArgs struct {
	FilenamePattern pulumi.StringInput   `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput   `pulumi:"odataType"`
	OutputFiles     OutputFileArrayInput `pulumi:"outputFiles"`
}

func (TransportStreamFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportStreamFormat)(nil)).Elem()
}

func (i TransportStreamFormatArgs) ToTransportStreamFormatOutput() TransportStreamFormatOutput {
	return i.ToTransportStreamFormatOutputWithContext(context.Background())
}

func (i TransportStreamFormatArgs) ToTransportStreamFormatOutputWithContext(ctx context.Context) TransportStreamFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportStreamFormatOutput)
}

type TransportStreamFormatOutput struct{ *pulumi.OutputState }

func (TransportStreamFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportStreamFormat)(nil)).Elem()
}

func (o TransportStreamFormatOutput) ToTransportStreamFormatOutput() TransportStreamFormatOutput {
	return o
}

func (o TransportStreamFormatOutput) ToTransportStreamFormatOutputWithContext(ctx context.Context) TransportStreamFormatOutput {
	return o
}

func (o TransportStreamFormatOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v TransportStreamFormat) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o TransportStreamFormatOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v TransportStreamFormat) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o TransportStreamFormatOutput) OutputFiles() OutputFileArrayOutput {
	return o.ApplyT(func(v TransportStreamFormat) []OutputFile { return v.OutputFiles }).(OutputFileArrayOutput)
}

type TransportStreamFormatResponse struct {
	FilenamePattern string               `pulumi:"filenamePattern"`
	OdataType       string               `pulumi:"odataType"`
	OutputFiles     []OutputFileResponse `pulumi:"outputFiles"`
}





type TransportStreamFormatResponseInput interface {
	pulumi.Input

	ToTransportStreamFormatResponseOutput() TransportStreamFormatResponseOutput
	ToTransportStreamFormatResponseOutputWithContext(context.Context) TransportStreamFormatResponseOutput
}

type TransportStreamFormatResponseArgs struct {
	FilenamePattern pulumi.StringInput           `pulumi:"filenamePattern"`
	OdataType       pulumi.StringInput           `pulumi:"odataType"`
	OutputFiles     OutputFileResponseArrayInput `pulumi:"outputFiles"`
}

func (TransportStreamFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportStreamFormatResponse)(nil)).Elem()
}

func (i TransportStreamFormatResponseArgs) ToTransportStreamFormatResponseOutput() TransportStreamFormatResponseOutput {
	return i.ToTransportStreamFormatResponseOutputWithContext(context.Background())
}

func (i TransportStreamFormatResponseArgs) ToTransportStreamFormatResponseOutputWithContext(ctx context.Context) TransportStreamFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransportStreamFormatResponseOutput)
}

type TransportStreamFormatResponseOutput struct{ *pulumi.OutputState }

func (TransportStreamFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportStreamFormatResponse)(nil)).Elem()
}

func (o TransportStreamFormatResponseOutput) ToTransportStreamFormatResponseOutput() TransportStreamFormatResponseOutput {
	return o
}

func (o TransportStreamFormatResponseOutput) ToTransportStreamFormatResponseOutputWithContext(ctx context.Context) TransportStreamFormatResponseOutput {
	return o
}

func (o TransportStreamFormatResponseOutput) FilenamePattern() pulumi.StringOutput {
	return o.ApplyT(func(v TransportStreamFormatResponse) string { return v.FilenamePattern }).(pulumi.StringOutput)
}

func (o TransportStreamFormatResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v TransportStreamFormatResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o TransportStreamFormatResponseOutput) OutputFiles() OutputFileResponseArrayOutput {
	return o.ApplyT(func(v TransportStreamFormatResponse) []OutputFileResponse { return v.OutputFiles }).(OutputFileResponseArrayOutput)
}

type UtcClipTime struct {
	OdataType string `pulumi:"odataType"`
	Time      string `pulumi:"time"`
}





type UtcClipTimeInput interface {
	pulumi.Input

	ToUtcClipTimeOutput() UtcClipTimeOutput
	ToUtcClipTimeOutputWithContext(context.Context) UtcClipTimeOutput
}

type UtcClipTimeArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
	Time      pulumi.StringInput `pulumi:"time"`
}

func (UtcClipTimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UtcClipTime)(nil)).Elem()
}

func (i UtcClipTimeArgs) ToUtcClipTimeOutput() UtcClipTimeOutput {
	return i.ToUtcClipTimeOutputWithContext(context.Background())
}

func (i UtcClipTimeArgs) ToUtcClipTimeOutputWithContext(ctx context.Context) UtcClipTimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UtcClipTimeOutput)
}

type UtcClipTimeOutput struct{ *pulumi.OutputState }

func (UtcClipTimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UtcClipTime)(nil)).Elem()
}

func (o UtcClipTimeOutput) ToUtcClipTimeOutput() UtcClipTimeOutput {
	return o
}

func (o UtcClipTimeOutput) ToUtcClipTimeOutputWithContext(ctx context.Context) UtcClipTimeOutput {
	return o
}

func (o UtcClipTimeOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UtcClipTime) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UtcClipTimeOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v UtcClipTime) string { return v.Time }).(pulumi.StringOutput)
}

type UtcClipTimeResponse struct {
	OdataType string `pulumi:"odataType"`
	Time      string `pulumi:"time"`
}





type UtcClipTimeResponseInput interface {
	pulumi.Input

	ToUtcClipTimeResponseOutput() UtcClipTimeResponseOutput
	ToUtcClipTimeResponseOutputWithContext(context.Context) UtcClipTimeResponseOutput
}

type UtcClipTimeResponseArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
	Time      pulumi.StringInput `pulumi:"time"`
}

func (UtcClipTimeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UtcClipTimeResponse)(nil)).Elem()
}

func (i UtcClipTimeResponseArgs) ToUtcClipTimeResponseOutput() UtcClipTimeResponseOutput {
	return i.ToUtcClipTimeResponseOutputWithContext(context.Background())
}

func (i UtcClipTimeResponseArgs) ToUtcClipTimeResponseOutputWithContext(ctx context.Context) UtcClipTimeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UtcClipTimeResponseOutput)
}

type UtcClipTimeResponseOutput struct{ *pulumi.OutputState }

func (UtcClipTimeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UtcClipTimeResponse)(nil)).Elem()
}

func (o UtcClipTimeResponseOutput) ToUtcClipTimeResponseOutput() UtcClipTimeResponseOutput {
	return o
}

func (o UtcClipTimeResponseOutput) ToUtcClipTimeResponseOutputWithContext(ctx context.Context) UtcClipTimeResponseOutput {
	return o
}

func (o UtcClipTimeResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v UtcClipTimeResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o UtcClipTimeResponseOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v UtcClipTimeResponse) string { return v.Time }).(pulumi.StringOutput)
}

type Video struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	StretchMode      *string `pulumi:"stretchMode"`
	SyncMode         *string `pulumi:"syncMode"`
}





type VideoInput interface {
	pulumi.Input

	ToVideoOutput() VideoOutput
	ToVideoOutputWithContext(context.Context) VideoOutput
}

type VideoArgs struct {
	KeyFrameInterval pulumi.StringPtrInput `pulumi:"keyFrameInterval"`
	Label            pulumi.StringPtrInput `pulumi:"label"`
	OdataType        pulumi.StringInput    `pulumi:"odataType"`
	StretchMode      pulumi.StringPtrInput `pulumi:"stretchMode"`
	SyncMode         pulumi.StringPtrInput `pulumi:"syncMode"`
}

func (VideoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Video)(nil)).Elem()
}

func (i VideoArgs) ToVideoOutput() VideoOutput {
	return i.ToVideoOutputWithContext(context.Background())
}

func (i VideoArgs) ToVideoOutputWithContext(ctx context.Context) VideoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoOutput)
}

type VideoOutput struct{ *pulumi.OutputState }

func (VideoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Video)(nil)).Elem()
}

func (o VideoOutput) ToVideoOutput() VideoOutput {
	return o
}

func (o VideoOutput) ToVideoOutputWithContext(ctx context.Context) VideoOutput {
	return o
}

func (o VideoOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Video) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o VideoOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Video) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o VideoOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v Video) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o VideoOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Video) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o VideoOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Video) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type VideoAnalyzerPreset struct {
	AudioLanguage       *string           `pulumi:"audioLanguage"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	InsightsToExtract   *string           `pulumi:"insightsToExtract"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
}





type VideoAnalyzerPresetInput interface {
	pulumi.Input

	ToVideoAnalyzerPresetOutput() VideoAnalyzerPresetOutput
	ToVideoAnalyzerPresetOutputWithContext(context.Context) VideoAnalyzerPresetOutput
}

type VideoAnalyzerPresetArgs struct {
	AudioLanguage       pulumi.StringPtrInput `pulumi:"audioLanguage"`
	ExperimentalOptions pulumi.StringMapInput `pulumi:"experimentalOptions"`
	InsightsToExtract   pulumi.StringPtrInput `pulumi:"insightsToExtract"`
	Mode                pulumi.StringPtrInput `pulumi:"mode"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
}

func (VideoAnalyzerPresetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoAnalyzerPreset)(nil)).Elem()
}

func (i VideoAnalyzerPresetArgs) ToVideoAnalyzerPresetOutput() VideoAnalyzerPresetOutput {
	return i.ToVideoAnalyzerPresetOutputWithContext(context.Background())
}

func (i VideoAnalyzerPresetArgs) ToVideoAnalyzerPresetOutputWithContext(ctx context.Context) VideoAnalyzerPresetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerPresetOutput)
}

type VideoAnalyzerPresetOutput struct{ *pulumi.OutputState }

func (VideoAnalyzerPresetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoAnalyzerPreset)(nil)).Elem()
}

func (o VideoAnalyzerPresetOutput) ToVideoAnalyzerPresetOutput() VideoAnalyzerPresetOutput {
	return o
}

func (o VideoAnalyzerPresetOutput) ToVideoAnalyzerPresetOutputWithContext(ctx context.Context) VideoAnalyzerPresetOutput {
	return o
}

func (o VideoAnalyzerPresetOutput) AudioLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoAnalyzerPreset) *string { return v.AudioLanguage }).(pulumi.StringPtrOutput)
}

func (o VideoAnalyzerPresetOutput) ExperimentalOptions() pulumi.StringMapOutput {
	return o.ApplyT(func(v VideoAnalyzerPreset) map[string]string { return v.ExperimentalOptions }).(pulumi.StringMapOutput)
}

func (o VideoAnalyzerPresetOutput) InsightsToExtract() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoAnalyzerPreset) *string { return v.InsightsToExtract }).(pulumi.StringPtrOutput)
}

func (o VideoAnalyzerPresetOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoAnalyzerPreset) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o VideoAnalyzerPresetOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v VideoAnalyzerPreset) string { return v.OdataType }).(pulumi.StringOutput)
}

type VideoAnalyzerPresetResponse struct {
	AudioLanguage       *string           `pulumi:"audioLanguage"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	InsightsToExtract   *string           `pulumi:"insightsToExtract"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
}





type VideoAnalyzerPresetResponseInput interface {
	pulumi.Input

	ToVideoAnalyzerPresetResponseOutput() VideoAnalyzerPresetResponseOutput
	ToVideoAnalyzerPresetResponseOutputWithContext(context.Context) VideoAnalyzerPresetResponseOutput
}

type VideoAnalyzerPresetResponseArgs struct {
	AudioLanguage       pulumi.StringPtrInput `pulumi:"audioLanguage"`
	ExperimentalOptions pulumi.StringMapInput `pulumi:"experimentalOptions"`
	InsightsToExtract   pulumi.StringPtrInput `pulumi:"insightsToExtract"`
	Mode                pulumi.StringPtrInput `pulumi:"mode"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
}

func (VideoAnalyzerPresetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoAnalyzerPresetResponse)(nil)).Elem()
}

func (i VideoAnalyzerPresetResponseArgs) ToVideoAnalyzerPresetResponseOutput() VideoAnalyzerPresetResponseOutput {
	return i.ToVideoAnalyzerPresetResponseOutputWithContext(context.Background())
}

func (i VideoAnalyzerPresetResponseArgs) ToVideoAnalyzerPresetResponseOutputWithContext(ctx context.Context) VideoAnalyzerPresetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoAnalyzerPresetResponseOutput)
}

type VideoAnalyzerPresetResponseOutput struct{ *pulumi.OutputState }

func (VideoAnalyzerPresetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoAnalyzerPresetResponse)(nil)).Elem()
}

func (o VideoAnalyzerPresetResponseOutput) ToVideoAnalyzerPresetResponseOutput() VideoAnalyzerPresetResponseOutput {
	return o
}

func (o VideoAnalyzerPresetResponseOutput) ToVideoAnalyzerPresetResponseOutputWithContext(ctx context.Context) VideoAnalyzerPresetResponseOutput {
	return o
}

func (o VideoAnalyzerPresetResponseOutput) AudioLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoAnalyzerPresetResponse) *string { return v.AudioLanguage }).(pulumi.StringPtrOutput)
}

func (o VideoAnalyzerPresetResponseOutput) ExperimentalOptions() pulumi.StringMapOutput {
	return o.ApplyT(func(v VideoAnalyzerPresetResponse) map[string]string { return v.ExperimentalOptions }).(pulumi.StringMapOutput)
}

func (o VideoAnalyzerPresetResponseOutput) InsightsToExtract() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoAnalyzerPresetResponse) *string { return v.InsightsToExtract }).(pulumi.StringPtrOutput)
}

func (o VideoAnalyzerPresetResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoAnalyzerPresetResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o VideoAnalyzerPresetResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v VideoAnalyzerPresetResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type VideoOverlay struct {
	AudioGainLevel  *float64   `pulumi:"audioGainLevel"`
	CropRectangle   *Rectangle `pulumi:"cropRectangle"`
	End             *string    `pulumi:"end"`
	FadeInDuration  *string    `pulumi:"fadeInDuration"`
	FadeOutDuration *string    `pulumi:"fadeOutDuration"`
	InputLabel      string     `pulumi:"inputLabel"`
	OdataType       string     `pulumi:"odataType"`
	Opacity         *float64   `pulumi:"opacity"`
	Position        *Rectangle `pulumi:"position"`
	Start           *string    `pulumi:"start"`
}





type VideoOverlayInput interface {
	pulumi.Input

	ToVideoOverlayOutput() VideoOverlayOutput
	ToVideoOverlayOutputWithContext(context.Context) VideoOverlayOutput
}

type VideoOverlayArgs struct {
	AudioGainLevel  pulumi.Float64PtrInput `pulumi:"audioGainLevel"`
	CropRectangle   RectanglePtrInput      `pulumi:"cropRectangle"`
	End             pulumi.StringPtrInput  `pulumi:"end"`
	FadeInDuration  pulumi.StringPtrInput  `pulumi:"fadeInDuration"`
	FadeOutDuration pulumi.StringPtrInput  `pulumi:"fadeOutDuration"`
	InputLabel      pulumi.StringInput     `pulumi:"inputLabel"`
	OdataType       pulumi.StringInput     `pulumi:"odataType"`
	Opacity         pulumi.Float64PtrInput `pulumi:"opacity"`
	Position        RectanglePtrInput      `pulumi:"position"`
	Start           pulumi.StringPtrInput  `pulumi:"start"`
}

func (VideoOverlayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoOverlay)(nil)).Elem()
}

func (i VideoOverlayArgs) ToVideoOverlayOutput() VideoOverlayOutput {
	return i.ToVideoOverlayOutputWithContext(context.Background())
}

func (i VideoOverlayArgs) ToVideoOverlayOutputWithContext(ctx context.Context) VideoOverlayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoOverlayOutput)
}

type VideoOverlayOutput struct{ *pulumi.OutputState }

func (VideoOverlayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoOverlay)(nil)).Elem()
}

func (o VideoOverlayOutput) ToVideoOverlayOutput() VideoOverlayOutput {
	return o
}

func (o VideoOverlayOutput) ToVideoOverlayOutputWithContext(ctx context.Context) VideoOverlayOutput {
	return o
}

func (o VideoOverlayOutput) AudioGainLevel() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VideoOverlay) *float64 { return v.AudioGainLevel }).(pulumi.Float64PtrOutput)
}

func (o VideoOverlayOutput) CropRectangle() RectanglePtrOutput {
	return o.ApplyT(func(v VideoOverlay) *Rectangle { return v.CropRectangle }).(RectanglePtrOutput)
}

func (o VideoOverlayOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoOverlay) *string { return v.End }).(pulumi.StringPtrOutput)
}

func (o VideoOverlayOutput) FadeInDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoOverlay) *string { return v.FadeInDuration }).(pulumi.StringPtrOutput)
}

func (o VideoOverlayOutput) FadeOutDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoOverlay) *string { return v.FadeOutDuration }).(pulumi.StringPtrOutput)
}

func (o VideoOverlayOutput) InputLabel() pulumi.StringOutput {
	return o.ApplyT(func(v VideoOverlay) string { return v.InputLabel }).(pulumi.StringOutput)
}

func (o VideoOverlayOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v VideoOverlay) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o VideoOverlayOutput) Opacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VideoOverlay) *float64 { return v.Opacity }).(pulumi.Float64PtrOutput)
}

func (o VideoOverlayOutput) Position() RectanglePtrOutput {
	return o.ApplyT(func(v VideoOverlay) *Rectangle { return v.Position }).(RectanglePtrOutput)
}

func (o VideoOverlayOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoOverlay) *string { return v.Start }).(pulumi.StringPtrOutput)
}

type VideoOverlayResponse struct {
	AudioGainLevel  *float64           `pulumi:"audioGainLevel"`
	CropRectangle   *RectangleResponse `pulumi:"cropRectangle"`
	End             *string            `pulumi:"end"`
	FadeInDuration  *string            `pulumi:"fadeInDuration"`
	FadeOutDuration *string            `pulumi:"fadeOutDuration"`
	InputLabel      string             `pulumi:"inputLabel"`
	OdataType       string             `pulumi:"odataType"`
	Opacity         *float64           `pulumi:"opacity"`
	Position        *RectangleResponse `pulumi:"position"`
	Start           *string            `pulumi:"start"`
}





type VideoOverlayResponseInput interface {
	pulumi.Input

	ToVideoOverlayResponseOutput() VideoOverlayResponseOutput
	ToVideoOverlayResponseOutputWithContext(context.Context) VideoOverlayResponseOutput
}

type VideoOverlayResponseArgs struct {
	AudioGainLevel  pulumi.Float64PtrInput    `pulumi:"audioGainLevel"`
	CropRectangle   RectangleResponsePtrInput `pulumi:"cropRectangle"`
	End             pulumi.StringPtrInput     `pulumi:"end"`
	FadeInDuration  pulumi.StringPtrInput     `pulumi:"fadeInDuration"`
	FadeOutDuration pulumi.StringPtrInput     `pulumi:"fadeOutDuration"`
	InputLabel      pulumi.StringInput        `pulumi:"inputLabel"`
	OdataType       pulumi.StringInput        `pulumi:"odataType"`
	Opacity         pulumi.Float64PtrInput    `pulumi:"opacity"`
	Position        RectangleResponsePtrInput `pulumi:"position"`
	Start           pulumi.StringPtrInput     `pulumi:"start"`
}

func (VideoOverlayResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoOverlayResponse)(nil)).Elem()
}

func (i VideoOverlayResponseArgs) ToVideoOverlayResponseOutput() VideoOverlayResponseOutput {
	return i.ToVideoOverlayResponseOutputWithContext(context.Background())
}

func (i VideoOverlayResponseArgs) ToVideoOverlayResponseOutputWithContext(ctx context.Context) VideoOverlayResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoOverlayResponseOutput)
}

type VideoOverlayResponseOutput struct{ *pulumi.OutputState }

func (VideoOverlayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoOverlayResponse)(nil)).Elem()
}

func (o VideoOverlayResponseOutput) ToVideoOverlayResponseOutput() VideoOverlayResponseOutput {
	return o
}

func (o VideoOverlayResponseOutput) ToVideoOverlayResponseOutputWithContext(ctx context.Context) VideoOverlayResponseOutput {
	return o
}

func (o VideoOverlayResponseOutput) AudioGainLevel() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VideoOverlayResponse) *float64 { return v.AudioGainLevel }).(pulumi.Float64PtrOutput)
}

func (o VideoOverlayResponseOutput) CropRectangle() RectangleResponsePtrOutput {
	return o.ApplyT(func(v VideoOverlayResponse) *RectangleResponse { return v.CropRectangle }).(RectangleResponsePtrOutput)
}

func (o VideoOverlayResponseOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoOverlayResponse) *string { return v.End }).(pulumi.StringPtrOutput)
}

func (o VideoOverlayResponseOutput) FadeInDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoOverlayResponse) *string { return v.FadeInDuration }).(pulumi.StringPtrOutput)
}

func (o VideoOverlayResponseOutput) FadeOutDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoOverlayResponse) *string { return v.FadeOutDuration }).(pulumi.StringPtrOutput)
}

func (o VideoOverlayResponseOutput) InputLabel() pulumi.StringOutput {
	return o.ApplyT(func(v VideoOverlayResponse) string { return v.InputLabel }).(pulumi.StringOutput)
}

func (o VideoOverlayResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v VideoOverlayResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o VideoOverlayResponseOutput) Opacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VideoOverlayResponse) *float64 { return v.Opacity }).(pulumi.Float64PtrOutput)
}

func (o VideoOverlayResponseOutput) Position() RectangleResponsePtrOutput {
	return o.ApplyT(func(v VideoOverlayResponse) *RectangleResponse { return v.Position }).(RectangleResponsePtrOutput)
}

func (o VideoOverlayResponseOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoOverlayResponse) *string { return v.Start }).(pulumi.StringPtrOutput)
}

type VideoResponse struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	StretchMode      *string `pulumi:"stretchMode"`
	SyncMode         *string `pulumi:"syncMode"`
}





type VideoResponseInput interface {
	pulumi.Input

	ToVideoResponseOutput() VideoResponseOutput
	ToVideoResponseOutputWithContext(context.Context) VideoResponseOutput
}

type VideoResponseArgs struct {
	KeyFrameInterval pulumi.StringPtrInput `pulumi:"keyFrameInterval"`
	Label            pulumi.StringPtrInput `pulumi:"label"`
	OdataType        pulumi.StringInput    `pulumi:"odataType"`
	StretchMode      pulumi.StringPtrInput `pulumi:"stretchMode"`
	SyncMode         pulumi.StringPtrInput `pulumi:"syncMode"`
}

func (VideoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoResponse)(nil)).Elem()
}

func (i VideoResponseArgs) ToVideoResponseOutput() VideoResponseOutput {
	return i.ToVideoResponseOutputWithContext(context.Background())
}

func (i VideoResponseArgs) ToVideoResponseOutputWithContext(ctx context.Context) VideoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoResponseOutput)
}

type VideoResponseOutput struct{ *pulumi.OutputState }

func (VideoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoResponse)(nil)).Elem()
}

func (o VideoResponseOutput) ToVideoResponseOutput() VideoResponseOutput {
	return o
}

func (o VideoResponseOutput) ToVideoResponseOutputWithContext(ctx context.Context) VideoResponseOutput {
	return o
}

func (o VideoResponseOutput) KeyFrameInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoResponse) *string { return v.KeyFrameInterval }).(pulumi.StringPtrOutput)
}

func (o VideoResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o VideoResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v VideoResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o VideoResponseOutput) StretchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoResponse) *string { return v.StretchMode }).(pulumi.StringPtrOutput)
}

func (o VideoResponseOutput) SyncMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoResponse) *string { return v.SyncMode }).(pulumi.StringPtrOutput)
}

type VideoTrackDescriptor struct {
	OdataType string `pulumi:"odataType"`
}





type VideoTrackDescriptorInput interface {
	pulumi.Input

	ToVideoTrackDescriptorOutput() VideoTrackDescriptorOutput
	ToVideoTrackDescriptorOutputWithContext(context.Context) VideoTrackDescriptorOutput
}

type VideoTrackDescriptorArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (VideoTrackDescriptorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoTrackDescriptor)(nil)).Elem()
}

func (i VideoTrackDescriptorArgs) ToVideoTrackDescriptorOutput() VideoTrackDescriptorOutput {
	return i.ToVideoTrackDescriptorOutputWithContext(context.Background())
}

func (i VideoTrackDescriptorArgs) ToVideoTrackDescriptorOutputWithContext(ctx context.Context) VideoTrackDescriptorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoTrackDescriptorOutput)
}

type VideoTrackDescriptorOutput struct{ *pulumi.OutputState }

func (VideoTrackDescriptorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoTrackDescriptor)(nil)).Elem()
}

func (o VideoTrackDescriptorOutput) ToVideoTrackDescriptorOutput() VideoTrackDescriptorOutput {
	return o
}

func (o VideoTrackDescriptorOutput) ToVideoTrackDescriptorOutputWithContext(ctx context.Context) VideoTrackDescriptorOutput {
	return o
}

func (o VideoTrackDescriptorOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v VideoTrackDescriptor) string { return v.OdataType }).(pulumi.StringOutput)
}

type VideoTrackDescriptorResponse struct {
	OdataType string `pulumi:"odataType"`
}





type VideoTrackDescriptorResponseInput interface {
	pulumi.Input

	ToVideoTrackDescriptorResponseOutput() VideoTrackDescriptorResponseOutput
	ToVideoTrackDescriptorResponseOutputWithContext(context.Context) VideoTrackDescriptorResponseOutput
}

type VideoTrackDescriptorResponseArgs struct {
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (VideoTrackDescriptorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoTrackDescriptorResponse)(nil)).Elem()
}

func (i VideoTrackDescriptorResponseArgs) ToVideoTrackDescriptorResponseOutput() VideoTrackDescriptorResponseOutput {
	return i.ToVideoTrackDescriptorResponseOutputWithContext(context.Background())
}

func (i VideoTrackDescriptorResponseArgs) ToVideoTrackDescriptorResponseOutputWithContext(ctx context.Context) VideoTrackDescriptorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoTrackDescriptorResponseOutput)
}

type VideoTrackDescriptorResponseOutput struct{ *pulumi.OutputState }

func (VideoTrackDescriptorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoTrackDescriptorResponse)(nil)).Elem()
}

func (o VideoTrackDescriptorResponseOutput) ToVideoTrackDescriptorResponseOutput() VideoTrackDescriptorResponseOutput {
	return o
}

func (o VideoTrackDescriptorResponseOutput) ToVideoTrackDescriptorResponseOutputWithContext(ctx context.Context) VideoTrackDescriptorResponseOutput {
	return o
}

func (o VideoTrackDescriptorResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v VideoTrackDescriptorResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AacAudioInput)(nil)).Elem(), AacAudioArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AacAudioResponseInput)(nil)).Elem(), AacAudioResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AbsoluteClipTimeInput)(nil)).Elem(), AbsoluteClipTimeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AbsoluteClipTimeResponseInput)(nil)).Elem(), AbsoluteClipTimeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountEncryptionInput)(nil)).Elem(), AccountEncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountEncryptionPtrInput)(nil)).Elem(), AccountEncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountEncryptionResponseInput)(nil)).Elem(), AccountEncryptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountEncryptionResponsePtrInput)(nil)).Elem(), AccountEncryptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiAccessControlInput)(nil)).Elem(), AkamaiAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiAccessControlPtrInput)(nil)).Elem(), AkamaiAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiAccessControlResponseInput)(nil)).Elem(), AkamaiAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiAccessControlResponsePtrInput)(nil)).Elem(), AkamaiAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyInput)(nil)).Elem(), AkamaiSignatureHeaderAuthenticationKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyArrayInput)(nil)).Elem(), AkamaiSignatureHeaderAuthenticationKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyResponseInput)(nil)).Elem(), AkamaiSignatureHeaderAuthenticationKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyResponseArrayInput)(nil)).Elem(), AkamaiSignatureHeaderAuthenticationKeyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssetFileEncryptionMetadataResponseInput)(nil)).Elem(), AssetFileEncryptionMetadataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssetFileEncryptionMetadataResponseArrayInput)(nil)).Elem(), AssetFileEncryptionMetadataResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssetStreamingLocatorResponseInput)(nil)).Elem(), AssetStreamingLocatorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssetStreamingLocatorResponseArrayInput)(nil)).Elem(), AssetStreamingLocatorResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioInput)(nil)).Elem(), AudioArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioAnalyzerPresetInput)(nil)).Elem(), AudioAnalyzerPresetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioAnalyzerPresetResponseInput)(nil)).Elem(), AudioAnalyzerPresetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioOverlayInput)(nil)).Elem(), AudioOverlayArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioOverlayResponseInput)(nil)).Elem(), AudioOverlayResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioResponseInput)(nil)).Elem(), AudioResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioTrackDescriptorInput)(nil)).Elem(), AudioTrackDescriptorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudioTrackDescriptorResponseInput)(nil)).Elem(), AudioTrackDescriptorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuiltInStandardEncoderPresetInput)(nil)).Elem(), BuiltInStandardEncoderPresetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BuiltInStandardEncoderPresetResponseInput)(nil)).Elem(), BuiltInStandardEncoderPresetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CbcsDrmConfigurationInput)(nil)).Elem(), CbcsDrmConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CbcsDrmConfigurationPtrInput)(nil)).Elem(), CbcsDrmConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CbcsDrmConfigurationResponseInput)(nil)).Elem(), CbcsDrmConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CbcsDrmConfigurationResponsePtrInput)(nil)).Elem(), CbcsDrmConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CencDrmConfigurationInput)(nil)).Elem(), CencDrmConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CencDrmConfigurationPtrInput)(nil)).Elem(), CencDrmConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CencDrmConfigurationResponseInput)(nil)).Elem(), CencDrmConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CencDrmConfigurationResponsePtrInput)(nil)).Elem(), CencDrmConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonEncryptionCbcsInput)(nil)).Elem(), CommonEncryptionCbcsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonEncryptionCbcsPtrInput)(nil)).Elem(), CommonEncryptionCbcsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonEncryptionCbcsResponseInput)(nil)).Elem(), CommonEncryptionCbcsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonEncryptionCbcsResponsePtrInput)(nil)).Elem(), CommonEncryptionCbcsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonEncryptionCencInput)(nil)).Elem(), CommonEncryptionCencArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonEncryptionCencPtrInput)(nil)).Elem(), CommonEncryptionCencArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonEncryptionCencResponseInput)(nil)).Elem(), CommonEncryptionCencResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonEncryptionCencResponsePtrInput)(nil)).Elem(), CommonEncryptionCencResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyClearKeyConfigurationInput)(nil)).Elem(), ContentKeyPolicyClearKeyConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyClearKeyConfigurationResponseInput)(nil)).Elem(), ContentKeyPolicyClearKeyConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyFairPlayConfigurationInput)(nil)).Elem(), ContentKeyPolicyFairPlayConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyFairPlayConfigurationResponseInput)(nil)).Elem(), ContentKeyPolicyFairPlayConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyFairPlayOfflineRentalConfigurationInput)(nil)).Elem(), ContentKeyPolicyFairPlayOfflineRentalConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrInput)(nil)).Elem(), ContentKeyPolicyFairPlayOfflineRentalConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseInput)(nil)).Elem(), ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrInput)(nil)).Elem(), ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyOpenRestrictionInput)(nil)).Elem(), ContentKeyPolicyOpenRestrictionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyOpenRestrictionResponseInput)(nil)).Elem(), ContentKeyPolicyOpenRestrictionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyOptionInput)(nil)).Elem(), ContentKeyPolicyOptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyOptionArrayInput)(nil)).Elem(), ContentKeyPolicyOptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyOptionResponseInput)(nil)).Elem(), ContentKeyPolicyOptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyOptionResponseArrayInput)(nil)).Elem(), ContentKeyPolicyOptionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyConfigurationInput)(nil)).Elem(), ContentKeyPolicyPlayReadyConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyConfigurationResponseInput)(nil)).Elem(), ContentKeyPolicyPlayReadyConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderInput)(nil)).Elem(), ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseInput)(nil)).Elem(), ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierInput)(nil)).Elem(), ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseInput)(nil)).Elem(), ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionInput)(nil)).Elem(), ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrInput)(nil)).Elem(), ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseInput)(nil)).Elem(), ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrInput)(nil)).Elem(), ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyLicenseInput)(nil)).Elem(), ContentKeyPolicyPlayReadyLicenseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyLicenseArrayInput)(nil)).Elem(), ContentKeyPolicyPlayReadyLicenseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyLicenseResponseInput)(nil)).Elem(), ContentKeyPolicyPlayReadyLicenseResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyLicenseResponseArrayInput)(nil)).Elem(), ContentKeyPolicyPlayReadyLicenseResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyPlayRightInput)(nil)).Elem(), ContentKeyPolicyPlayReadyPlayRightArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyPlayRightPtrInput)(nil)).Elem(), ContentKeyPolicyPlayReadyPlayRightArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyPlayRightResponseInput)(nil)).Elem(), ContentKeyPolicyPlayReadyPlayRightResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyPlayRightResponsePtrInput)(nil)).Elem(), ContentKeyPolicyPlayReadyPlayRightResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyRsaTokenKeyInput)(nil)).Elem(), ContentKeyPolicyRsaTokenKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyRsaTokenKeyResponseInput)(nil)).Elem(), ContentKeyPolicyRsaTokenKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicySymmetricTokenKeyInput)(nil)).Elem(), ContentKeyPolicySymmetricTokenKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicySymmetricTokenKeyResponseInput)(nil)).Elem(), ContentKeyPolicySymmetricTokenKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyTokenClaimInput)(nil)).Elem(), ContentKeyPolicyTokenClaimArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyTokenClaimArrayInput)(nil)).Elem(), ContentKeyPolicyTokenClaimArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyTokenClaimResponseInput)(nil)).Elem(), ContentKeyPolicyTokenClaimResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyTokenClaimResponseArrayInput)(nil)).Elem(), ContentKeyPolicyTokenClaimResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyTokenRestrictionInput)(nil)).Elem(), ContentKeyPolicyTokenRestrictionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyTokenRestrictionResponseInput)(nil)).Elem(), ContentKeyPolicyTokenRestrictionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyUnknownConfigurationInput)(nil)).Elem(), ContentKeyPolicyUnknownConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyUnknownConfigurationResponseInput)(nil)).Elem(), ContentKeyPolicyUnknownConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyUnknownRestrictionInput)(nil)).Elem(), ContentKeyPolicyUnknownRestrictionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyUnknownRestrictionResponseInput)(nil)).Elem(), ContentKeyPolicyUnknownRestrictionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyWidevineConfigurationInput)(nil)).Elem(), ContentKeyPolicyWidevineConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyWidevineConfigurationResponseInput)(nil)).Elem(), ContentKeyPolicyWidevineConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyX509CertificateTokenKeyInput)(nil)).Elem(), ContentKeyPolicyX509CertificateTokenKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyX509CertificateTokenKeyResponseInput)(nil)).Elem(), ContentKeyPolicyX509CertificateTokenKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CopyAudioInput)(nil)).Elem(), CopyAudioArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CopyAudioResponseInput)(nil)).Elem(), CopyAudioResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CopyVideoInput)(nil)).Elem(), CopyVideoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CopyVideoResponseInput)(nil)).Elem(), CopyVideoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CrossSiteAccessPoliciesInput)(nil)).Elem(), CrossSiteAccessPoliciesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CrossSiteAccessPoliciesPtrInput)(nil)).Elem(), CrossSiteAccessPoliciesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CrossSiteAccessPoliciesResponseInput)(nil)).Elem(), CrossSiteAccessPoliciesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CrossSiteAccessPoliciesResponsePtrInput)(nil)).Elem(), CrossSiteAccessPoliciesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultKeyInput)(nil)).Elem(), DefaultKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultKeyPtrInput)(nil)).Elem(), DefaultKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultKeyResponseInput)(nil)).Elem(), DefaultKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultKeyResponsePtrInput)(nil)).Elem(), DefaultKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeinterlaceInput)(nil)).Elem(), DeinterlaceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeinterlacePtrInput)(nil)).Elem(), DeinterlaceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeinterlaceResponseInput)(nil)).Elem(), DeinterlaceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeinterlaceResponsePtrInput)(nil)).Elem(), DeinterlaceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeUsageDataCollectionPolicyResponseInput)(nil)).Elem(), EdgeUsageDataCollectionPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeUsageDataCollectionPolicyResponsePtrInput)(nil)).Elem(), EdgeUsageDataCollectionPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeUsageDataEventHubResponseInput)(nil)).Elem(), EdgeUsageDataEventHubResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeUsageDataEventHubResponsePtrInput)(nil)).Elem(), EdgeUsageDataEventHubResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledProtocolsInput)(nil)).Elem(), EnabledProtocolsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledProtocolsPtrInput)(nil)).Elem(), EnabledProtocolsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledProtocolsResponseInput)(nil)).Elem(), EnabledProtocolsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledProtocolsResponsePtrInput)(nil)).Elem(), EnabledProtocolsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvelopeEncryptionInput)(nil)).Elem(), EnvelopeEncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvelopeEncryptionPtrInput)(nil)).Elem(), EnvelopeEncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvelopeEncryptionResponseInput)(nil)).Elem(), EnvelopeEncryptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvelopeEncryptionResponsePtrInput)(nil)).Elem(), EnvelopeEncryptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FaceDetectorPresetInput)(nil)).Elem(), FaceDetectorPresetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FaceDetectorPresetResponseInput)(nil)).Elem(), FaceDetectorPresetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterTrackPropertyConditionInput)(nil)).Elem(), FilterTrackPropertyConditionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterTrackPropertyConditionArrayInput)(nil)).Elem(), FilterTrackPropertyConditionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterTrackPropertyConditionResponseInput)(nil)).Elem(), FilterTrackPropertyConditionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterTrackPropertyConditionResponseArrayInput)(nil)).Elem(), FilterTrackPropertyConditionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterTrackSelectionInput)(nil)).Elem(), FilterTrackSelectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterTrackSelectionArrayInput)(nil)).Elem(), FilterTrackSelectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterTrackSelectionResponseInput)(nil)).Elem(), FilterTrackSelectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterTrackSelectionResponseArrayInput)(nil)).Elem(), FilterTrackSelectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FiltersInput)(nil)).Elem(), FiltersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FiltersPtrInput)(nil)).Elem(), FiltersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FiltersResponseInput)(nil)).Elem(), FiltersResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FiltersResponsePtrInput)(nil)).Elem(), FiltersResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirstQualityInput)(nil)).Elem(), FirstQualityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirstQualityPtrInput)(nil)).Elem(), FirstQualityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirstQualityResponseInput)(nil)).Elem(), FirstQualityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirstQualityResponsePtrInput)(nil)).Elem(), FirstQualityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FromAllInputFileInput)(nil)).Elem(), FromAllInputFileArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FromAllInputFileResponseInput)(nil)).Elem(), FromAllInputFileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FromEachInputFileInput)(nil)).Elem(), FromEachInputFileArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FromEachInputFileResponseInput)(nil)).Elem(), FromEachInputFileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*H264LayerInput)(nil)).Elem(), H264LayerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*H264LayerArrayInput)(nil)).Elem(), H264LayerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*H264LayerResponseInput)(nil)).Elem(), H264LayerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*H264LayerResponseArrayInput)(nil)).Elem(), H264LayerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*H264VideoInput)(nil)).Elem(), H264VideoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*H264VideoResponseInput)(nil)).Elem(), H264VideoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*H265LayerInput)(nil)).Elem(), H265LayerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*H265LayerArrayInput)(nil)).Elem(), H265LayerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*H265LayerResponseInput)(nil)).Elem(), H265LayerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*H265LayerResponseArrayInput)(nil)).Elem(), H265LayerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*H265VideoInput)(nil)).Elem(), H265VideoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*H265VideoResponseInput)(nil)).Elem(), H265VideoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HlsInput)(nil)).Elem(), HlsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HlsPtrInput)(nil)).Elem(), HlsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HlsResponseInput)(nil)).Elem(), HlsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HlsResponsePtrInput)(nil)).Elem(), HlsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPAccessControlInput)(nil)).Elem(), IPAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPAccessControlPtrInput)(nil)).Elem(), IPAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPAccessControlResponseInput)(nil)).Elem(), IPAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPAccessControlResponsePtrInput)(nil)).Elem(), IPAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRangeInput)(nil)).Elem(), IPRangeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRangeArrayInput)(nil)).Elem(), IPRangeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRangeResponseInput)(nil)).Elem(), IPRangeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRangeResponseArrayInput)(nil)).Elem(), IPRangeResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), ImageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageFormatInput)(nil)).Elem(), ImageFormatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageFormatResponseInput)(nil)).Elem(), ImageFormatResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageResponseInput)(nil)).Elem(), ImageResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputFileInput)(nil)).Elem(), InputFileArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputFileResponseInput)(nil)).Elem(), InputFileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobErrorDetailResponseInput)(nil)).Elem(), JobErrorDetailResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobErrorDetailResponseArrayInput)(nil)).Elem(), JobErrorDetailResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobErrorResponseInput)(nil)).Elem(), JobErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputAssetInput)(nil)).Elem(), JobInputAssetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputAssetResponseInput)(nil)).Elem(), JobInputAssetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputClipInput)(nil)).Elem(), JobInputClipArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputClipArrayInput)(nil)).Elem(), JobInputClipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputClipResponseInput)(nil)).Elem(), JobInputClipResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputClipResponseArrayInput)(nil)).Elem(), JobInputClipResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputHttpInput)(nil)).Elem(), JobInputHttpArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputHttpResponseInput)(nil)).Elem(), JobInputHttpResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputSequenceInput)(nil)).Elem(), JobInputSequenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputSequenceResponseInput)(nil)).Elem(), JobInputSequenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputsInput)(nil)).Elem(), JobInputsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobInputsResponseInput)(nil)).Elem(), JobInputsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobOutputAssetInput)(nil)).Elem(), JobOutputAssetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobOutputAssetArrayInput)(nil)).Elem(), JobOutputAssetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobOutputAssetResponseInput)(nil)).Elem(), JobOutputAssetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobOutputAssetResponseArrayInput)(nil)).Elem(), JobOutputAssetResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JpgFormatInput)(nil)).Elem(), JpgFormatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JpgFormatResponseInput)(nil)).Elem(), JpgFormatResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JpgImageInput)(nil)).Elem(), JpgImageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JpgImageResponseInput)(nil)).Elem(), JpgImageResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JpgLayerInput)(nil)).Elem(), JpgLayerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JpgLayerArrayInput)(nil)).Elem(), JpgLayerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JpgLayerResponseInput)(nil)).Elem(), JpgLayerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JpgLayerResponseArrayInput)(nil)).Elem(), JpgLayerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesInput)(nil)).Elem(), KeyVaultPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesPtrInput)(nil)).Elem(), KeyVaultPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesResponseInput)(nil)).Elem(), KeyVaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesResponsePtrInput)(nil)).Elem(), KeyVaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingInput)(nil)).Elem(), LiveEventEncodingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingPtrInput)(nil)).Elem(), LiveEventEncodingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingResponseInput)(nil)).Elem(), LiveEventEncodingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingResponsePtrInput)(nil)).Elem(), LiveEventEncodingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEndpointInput)(nil)).Elem(), LiveEventEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEndpointArrayInput)(nil)).Elem(), LiveEventEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEndpointResponseInput)(nil)).Elem(), LiveEventEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEndpointResponseArrayInput)(nil)).Elem(), LiveEventEndpointResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTypeInput)(nil)).Elem(), LiveEventInputTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTypePtrInput)(nil)).Elem(), LiveEventInputTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputAccessControlInput)(nil)).Elem(), LiveEventInputAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputAccessControlPtrInput)(nil)).Elem(), LiveEventInputAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputAccessControlResponseInput)(nil)).Elem(), LiveEventInputAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputAccessControlResponsePtrInput)(nil)).Elem(), LiveEventInputAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputResponseInput)(nil)).Elem(), LiveEventInputResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputResponsePtrInput)(nil)).Elem(), LiveEventInputResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTrackSelectionInput)(nil)).Elem(), LiveEventInputTrackSelectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTrackSelectionArrayInput)(nil)).Elem(), LiveEventInputTrackSelectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTrackSelectionResponseInput)(nil)).Elem(), LiveEventInputTrackSelectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputTrackSelectionResponseArrayInput)(nil)).Elem(), LiveEventInputTrackSelectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventOutputTranscriptionTrackInput)(nil)).Elem(), LiveEventOutputTranscriptionTrackArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventOutputTranscriptionTrackPtrInput)(nil)).Elem(), LiveEventOutputTranscriptionTrackArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventOutputTranscriptionTrackResponseInput)(nil)).Elem(), LiveEventOutputTranscriptionTrackResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventOutputTranscriptionTrackResponsePtrInput)(nil)).Elem(), LiveEventOutputTranscriptionTrackResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewInput)(nil)).Elem(), LiveEventPreviewArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewPtrInput)(nil)).Elem(), LiveEventPreviewArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewAccessControlInput)(nil)).Elem(), LiveEventPreviewAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewAccessControlPtrInput)(nil)).Elem(), LiveEventPreviewAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewAccessControlResponseInput)(nil)).Elem(), LiveEventPreviewAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewAccessControlResponsePtrInput)(nil)).Elem(), LiveEventPreviewAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewResponseInput)(nil)).Elem(), LiveEventPreviewResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventPreviewResponsePtrInput)(nil)).Elem(), LiveEventPreviewResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventTranscriptionInput)(nil)).Elem(), LiveEventTranscriptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventTranscriptionArrayInput)(nil)).Elem(), LiveEventTranscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventTranscriptionResponseInput)(nil)).Elem(), LiveEventTranscriptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventTranscriptionResponseArrayInput)(nil)).Elem(), LiveEventTranscriptionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaServiceIdentityInput)(nil)).Elem(), MediaServiceIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaServiceIdentityPtrInput)(nil)).Elem(), MediaServiceIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaServiceIdentityResponseInput)(nil)).Elem(), MediaServiceIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaServiceIdentityResponsePtrInput)(nil)).Elem(), MediaServiceIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Mp4FormatInput)(nil)).Elem(), Mp4FormatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Mp4FormatResponseInput)(nil)).Elem(), Mp4FormatResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MultiBitrateFormatInput)(nil)).Elem(), MultiBitrateFormatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MultiBitrateFormatResponseInput)(nil)).Elem(), MultiBitrateFormatResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NoEncryptionInput)(nil)).Elem(), NoEncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NoEncryptionPtrInput)(nil)).Elem(), NoEncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NoEncryptionResponseInput)(nil)).Elem(), NoEncryptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NoEncryptionResponsePtrInput)(nil)).Elem(), NoEncryptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputFileInput)(nil)).Elem(), OutputFileArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputFileArrayInput)(nil)).Elem(), OutputFileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputFileResponseInput)(nil)).Elem(), OutputFileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OutputFileResponseArrayInput)(nil)).Elem(), OutputFileResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PngFormatInput)(nil)).Elem(), PngFormatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PngFormatResponseInput)(nil)).Elem(), PngFormatResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PngImageInput)(nil)).Elem(), PngImageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PngImageResponseInput)(nil)).Elem(), PngImageResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PngLayerInput)(nil)).Elem(), PngLayerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PngLayerArrayInput)(nil)).Elem(), PngLayerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PngLayerResponseInput)(nil)).Elem(), PngLayerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PngLayerResponseArrayInput)(nil)).Elem(), PngLayerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PresentationTimeRangeInput)(nil)).Elem(), PresentationTimeRangeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PresentationTimeRangePtrInput)(nil)).Elem(), PresentationTimeRangeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PresentationTimeRangeResponseInput)(nil)).Elem(), PresentationTimeRangeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PresentationTimeRangeResponsePtrInput)(nil)).Elem(), PresentationTimeRangeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponseInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponsePtrInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponseInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponsePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RectangleInput)(nil)).Elem(), RectangleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RectanglePtrInput)(nil)).Elem(), RectangleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RectangleResponseInput)(nil)).Elem(), RectangleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RectangleResponsePtrInput)(nil)).Elem(), RectangleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelectAudioTrackByAttributeInput)(nil)).Elem(), SelectAudioTrackByAttributeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelectAudioTrackByAttributeResponseInput)(nil)).Elem(), SelectAudioTrackByAttributeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelectAudioTrackByIdInput)(nil)).Elem(), SelectAudioTrackByIdArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelectAudioTrackByIdResponseInput)(nil)).Elem(), SelectAudioTrackByIdResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelectVideoTrackByAttributeInput)(nil)).Elem(), SelectVideoTrackByAttributeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelectVideoTrackByAttributeResponseInput)(nil)).Elem(), SelectVideoTrackByAttributeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelectVideoTrackByIdInput)(nil)).Elem(), SelectVideoTrackByIdArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelectVideoTrackByIdResponseInput)(nil)).Elem(), SelectVideoTrackByIdResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StandardEncoderPresetInput)(nil)).Elem(), StandardEncoderPresetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StandardEncoderPresetResponseInput)(nil)).Elem(), StandardEncoderPresetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountInput)(nil)).Elem(), StorageAccountArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountArrayInput)(nil)).Elem(), StorageAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountResponseInput)(nil)).Elem(), StorageAccountResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountResponseArrayInput)(nil)).Elem(), StorageAccountResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingEndpointAccessControlInput)(nil)).Elem(), StreamingEndpointAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingEndpointAccessControlPtrInput)(nil)).Elem(), StreamingEndpointAccessControlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingEndpointAccessControlResponseInput)(nil)).Elem(), StreamingEndpointAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingEndpointAccessControlResponsePtrInput)(nil)).Elem(), StreamingEndpointAccessControlResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingLocatorContentKeyInput)(nil)).Elem(), StreamingLocatorContentKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingLocatorContentKeyArrayInput)(nil)).Elem(), StreamingLocatorContentKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingLocatorContentKeyResponseInput)(nil)).Elem(), StreamingLocatorContentKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingLocatorContentKeyResponseArrayInput)(nil)).Elem(), StreamingLocatorContentKeyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPathResponseInput)(nil)).Elem(), StreamingPathResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPathResponseArrayInput)(nil)).Elem(), StreamingPathResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyContentKeyInput)(nil)).Elem(), StreamingPolicyContentKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyContentKeyArrayInput)(nil)).Elem(), StreamingPolicyContentKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyContentKeyResponseInput)(nil)).Elem(), StreamingPolicyContentKeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyContentKeyResponseArrayInput)(nil)).Elem(), StreamingPolicyContentKeyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyContentKeysInput)(nil)).Elem(), StreamingPolicyContentKeysArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyContentKeysPtrInput)(nil)).Elem(), StreamingPolicyContentKeysArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyContentKeysResponseInput)(nil)).Elem(), StreamingPolicyContentKeysResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyContentKeysResponsePtrInput)(nil)).Elem(), StreamingPolicyContentKeysResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyFairPlayConfigurationInput)(nil)).Elem(), StreamingPolicyFairPlayConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyFairPlayConfigurationPtrInput)(nil)).Elem(), StreamingPolicyFairPlayConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyFairPlayConfigurationResponseInput)(nil)).Elem(), StreamingPolicyFairPlayConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyFairPlayConfigurationResponsePtrInput)(nil)).Elem(), StreamingPolicyFairPlayConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyPlayReadyConfigurationInput)(nil)).Elem(), StreamingPolicyPlayReadyConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyPlayReadyConfigurationPtrInput)(nil)).Elem(), StreamingPolicyPlayReadyConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyPlayReadyConfigurationResponseInput)(nil)).Elem(), StreamingPolicyPlayReadyConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyPlayReadyConfigurationResponsePtrInput)(nil)).Elem(), StreamingPolicyPlayReadyConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyWidevineConfigurationInput)(nil)).Elem(), StreamingPolicyWidevineConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyWidevineConfigurationPtrInput)(nil)).Elem(), StreamingPolicyWidevineConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyWidevineConfigurationResponseInput)(nil)).Elem(), StreamingPolicyWidevineConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingPolicyWidevineConfigurationResponsePtrInput)(nil)).Elem(), StreamingPolicyWidevineConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrackPropertyConditionInput)(nil)).Elem(), TrackPropertyConditionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrackPropertyConditionArrayInput)(nil)).Elem(), TrackPropertyConditionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrackPropertyConditionResponseInput)(nil)).Elem(), TrackPropertyConditionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrackPropertyConditionResponseArrayInput)(nil)).Elem(), TrackPropertyConditionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrackSelectionInput)(nil)).Elem(), TrackSelectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrackSelectionArrayInput)(nil)).Elem(), TrackSelectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrackSelectionResponseInput)(nil)).Elem(), TrackSelectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrackSelectionResponseArrayInput)(nil)).Elem(), TrackSelectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransformOutputTypeInput)(nil)).Elem(), TransformOutputTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransformOutputTypeArrayInput)(nil)).Elem(), TransformOutputTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransformOutputResponseInput)(nil)).Elem(), TransformOutputResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransformOutputResponseArrayInput)(nil)).Elem(), TransformOutputResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransportStreamFormatInput)(nil)).Elem(), TransportStreamFormatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransportStreamFormatResponseInput)(nil)).Elem(), TransportStreamFormatResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UtcClipTimeInput)(nil)).Elem(), UtcClipTimeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UtcClipTimeResponseInput)(nil)).Elem(), UtcClipTimeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoInput)(nil)).Elem(), VideoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoAnalyzerPresetInput)(nil)).Elem(), VideoAnalyzerPresetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoAnalyzerPresetResponseInput)(nil)).Elem(), VideoAnalyzerPresetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoOverlayInput)(nil)).Elem(), VideoOverlayArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoOverlayResponseInput)(nil)).Elem(), VideoOverlayResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoResponseInput)(nil)).Elem(), VideoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoTrackDescriptorInput)(nil)).Elem(), VideoTrackDescriptorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoTrackDescriptorResponseInput)(nil)).Elem(), VideoTrackDescriptorResponseArgs{})
	pulumi.RegisterOutputType(AacAudioOutput{})
	pulumi.RegisterOutputType(AacAudioResponseOutput{})
	pulumi.RegisterOutputType(AbsoluteClipTimeOutput{})
	pulumi.RegisterOutputType(AbsoluteClipTimeResponseOutput{})
	pulumi.RegisterOutputType(AccountEncryptionOutput{})
	pulumi.RegisterOutputType(AccountEncryptionPtrOutput{})
	pulumi.RegisterOutputType(AccountEncryptionResponseOutput{})
	pulumi.RegisterOutputType(AccountEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlPtrOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlResponseOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyArrayOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyResponseOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(AssetFileEncryptionMetadataResponseOutput{})
	pulumi.RegisterOutputType(AssetFileEncryptionMetadataResponseArrayOutput{})
	pulumi.RegisterOutputType(AssetStreamingLocatorResponseOutput{})
	pulumi.RegisterOutputType(AssetStreamingLocatorResponseArrayOutput{})
	pulumi.RegisterOutputType(AudioOutput{})
	pulumi.RegisterOutputType(AudioAnalyzerPresetOutput{})
	pulumi.RegisterOutputType(AudioAnalyzerPresetResponseOutput{})
	pulumi.RegisterOutputType(AudioOverlayOutput{})
	pulumi.RegisterOutputType(AudioOverlayResponseOutput{})
	pulumi.RegisterOutputType(AudioResponseOutput{})
	pulumi.RegisterOutputType(AudioTrackDescriptorOutput{})
	pulumi.RegisterOutputType(AudioTrackDescriptorResponseOutput{})
	pulumi.RegisterOutputType(BuiltInStandardEncoderPresetOutput{})
	pulumi.RegisterOutputType(BuiltInStandardEncoderPresetResponseOutput{})
	pulumi.RegisterOutputType(CbcsDrmConfigurationOutput{})
	pulumi.RegisterOutputType(CbcsDrmConfigurationPtrOutput{})
	pulumi.RegisterOutputType(CbcsDrmConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CbcsDrmConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(CencDrmConfigurationOutput{})
	pulumi.RegisterOutputType(CencDrmConfigurationPtrOutput{})
	pulumi.RegisterOutputType(CencDrmConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CencDrmConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCbcsOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCbcsPtrOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCbcsResponseOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCbcsResponsePtrOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCencOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCencPtrOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCencResponseOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCencResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyClearKeyConfigurationOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyClearKeyConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyFairPlayConfigurationOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyFairPlayConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyFairPlayOfflineRentalConfigurationOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyFairPlayOfflineRentalConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyFairPlayOfflineRentalConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyFairPlayOfflineRentalConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyOpenRestrictionOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyOpenRestrictionResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyOptionOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyOptionArrayOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyOptionResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyOptionResponseArrayOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyConfigurationOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionPtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyLicenseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyLicenseArrayOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyLicenseResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyLicenseResponseArrayOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyPlayRightOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyPlayRightPtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyPlayRightResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyPlayRightResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyRsaTokenKeyOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyRsaTokenKeyResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicySymmetricTokenKeyOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicySymmetricTokenKeyResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyTokenClaimOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyTokenClaimArrayOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyTokenClaimResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyTokenClaimResponseArrayOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyTokenRestrictionOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyTokenRestrictionResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyUnknownConfigurationOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyUnknownConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyUnknownRestrictionOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyUnknownRestrictionResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyWidevineConfigurationOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyWidevineConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyX509CertificateTokenKeyOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyX509CertificateTokenKeyResponseOutput{})
	pulumi.RegisterOutputType(CopyAudioOutput{})
	pulumi.RegisterOutputType(CopyAudioResponseOutput{})
	pulumi.RegisterOutputType(CopyVideoOutput{})
	pulumi.RegisterOutputType(CopyVideoResponseOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesPtrOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesResponseOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesResponsePtrOutput{})
	pulumi.RegisterOutputType(DefaultKeyOutput{})
	pulumi.RegisterOutputType(DefaultKeyPtrOutput{})
	pulumi.RegisterOutputType(DefaultKeyResponseOutput{})
	pulumi.RegisterOutputType(DefaultKeyResponsePtrOutput{})
	pulumi.RegisterOutputType(DeinterlaceOutput{})
	pulumi.RegisterOutputType(DeinterlacePtrOutput{})
	pulumi.RegisterOutputType(DeinterlaceResponseOutput{})
	pulumi.RegisterOutputType(DeinterlaceResponsePtrOutput{})
	pulumi.RegisterOutputType(EdgeUsageDataCollectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(EdgeUsageDataCollectionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(EdgeUsageDataEventHubResponseOutput{})
	pulumi.RegisterOutputType(EdgeUsageDataEventHubResponsePtrOutput{})
	pulumi.RegisterOutputType(EnabledProtocolsOutput{})
	pulumi.RegisterOutputType(EnabledProtocolsPtrOutput{})
	pulumi.RegisterOutputType(EnabledProtocolsResponseOutput{})
	pulumi.RegisterOutputType(EnabledProtocolsResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvelopeEncryptionOutput{})
	pulumi.RegisterOutputType(EnvelopeEncryptionPtrOutput{})
	pulumi.RegisterOutputType(EnvelopeEncryptionResponseOutput{})
	pulumi.RegisterOutputType(EnvelopeEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(FaceDetectorPresetOutput{})
	pulumi.RegisterOutputType(FaceDetectorPresetResponseOutput{})
	pulumi.RegisterOutputType(FilterTrackPropertyConditionOutput{})
	pulumi.RegisterOutputType(FilterTrackPropertyConditionArrayOutput{})
	pulumi.RegisterOutputType(FilterTrackPropertyConditionResponseOutput{})
	pulumi.RegisterOutputType(FilterTrackPropertyConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(FilterTrackSelectionOutput{})
	pulumi.RegisterOutputType(FilterTrackSelectionArrayOutput{})
	pulumi.RegisterOutputType(FilterTrackSelectionResponseOutput{})
	pulumi.RegisterOutputType(FilterTrackSelectionResponseArrayOutput{})
	pulumi.RegisterOutputType(FiltersOutput{})
	pulumi.RegisterOutputType(FiltersPtrOutput{})
	pulumi.RegisterOutputType(FiltersResponseOutput{})
	pulumi.RegisterOutputType(FiltersResponsePtrOutput{})
	pulumi.RegisterOutputType(FirstQualityOutput{})
	pulumi.RegisterOutputType(FirstQualityPtrOutput{})
	pulumi.RegisterOutputType(FirstQualityResponseOutput{})
	pulumi.RegisterOutputType(FirstQualityResponsePtrOutput{})
	pulumi.RegisterOutputType(FromAllInputFileOutput{})
	pulumi.RegisterOutputType(FromAllInputFileResponseOutput{})
	pulumi.RegisterOutputType(FromEachInputFileOutput{})
	pulumi.RegisterOutputType(FromEachInputFileResponseOutput{})
	pulumi.RegisterOutputType(H264LayerOutput{})
	pulumi.RegisterOutputType(H264LayerArrayOutput{})
	pulumi.RegisterOutputType(H264LayerResponseOutput{})
	pulumi.RegisterOutputType(H264LayerResponseArrayOutput{})
	pulumi.RegisterOutputType(H264VideoOutput{})
	pulumi.RegisterOutputType(H264VideoResponseOutput{})
	pulumi.RegisterOutputType(H265LayerOutput{})
	pulumi.RegisterOutputType(H265LayerArrayOutput{})
	pulumi.RegisterOutputType(H265LayerResponseOutput{})
	pulumi.RegisterOutputType(H265LayerResponseArrayOutput{})
	pulumi.RegisterOutputType(H265VideoOutput{})
	pulumi.RegisterOutputType(H265VideoResponseOutput{})
	pulumi.RegisterOutputType(HlsOutput{})
	pulumi.RegisterOutputType(HlsPtrOutput{})
	pulumi.RegisterOutputType(HlsResponseOutput{})
	pulumi.RegisterOutputType(HlsResponsePtrOutput{})
	pulumi.RegisterOutputType(IPAccessControlOutput{})
	pulumi.RegisterOutputType(IPAccessControlPtrOutput{})
	pulumi.RegisterOutputType(IPAccessControlResponseOutput{})
	pulumi.RegisterOutputType(IPAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(IPRangeOutput{})
	pulumi.RegisterOutputType(IPRangeArrayOutput{})
	pulumi.RegisterOutputType(IPRangeResponseOutput{})
	pulumi.RegisterOutputType(IPRangeResponseArrayOutput{})
	pulumi.RegisterOutputType(ImageOutput{})
	pulumi.RegisterOutputType(ImageFormatOutput{})
	pulumi.RegisterOutputType(ImageFormatResponseOutput{})
	pulumi.RegisterOutputType(ImageResponseOutput{})
	pulumi.RegisterOutputType(InputFileOutput{})
	pulumi.RegisterOutputType(InputFileResponseOutput{})
	pulumi.RegisterOutputType(JobErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(JobErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(JobErrorResponseOutput{})
	pulumi.RegisterOutputType(JobInputAssetOutput{})
	pulumi.RegisterOutputType(JobInputAssetResponseOutput{})
	pulumi.RegisterOutputType(JobInputClipOutput{})
	pulumi.RegisterOutputType(JobInputClipArrayOutput{})
	pulumi.RegisterOutputType(JobInputClipResponseOutput{})
	pulumi.RegisterOutputType(JobInputClipResponseArrayOutput{})
	pulumi.RegisterOutputType(JobInputHttpOutput{})
	pulumi.RegisterOutputType(JobInputHttpResponseOutput{})
	pulumi.RegisterOutputType(JobInputSequenceOutput{})
	pulumi.RegisterOutputType(JobInputSequenceResponseOutput{})
	pulumi.RegisterOutputType(JobInputsOutput{})
	pulumi.RegisterOutputType(JobInputsResponseOutput{})
	pulumi.RegisterOutputType(JobOutputAssetOutput{})
	pulumi.RegisterOutputType(JobOutputAssetArrayOutput{})
	pulumi.RegisterOutputType(JobOutputAssetResponseOutput{})
	pulumi.RegisterOutputType(JobOutputAssetResponseArrayOutput{})
	pulumi.RegisterOutputType(JpgFormatOutput{})
	pulumi.RegisterOutputType(JpgFormatResponseOutput{})
	pulumi.RegisterOutputType(JpgImageOutput{})
	pulumi.RegisterOutputType(JpgImageResponseOutput{})
	pulumi.RegisterOutputType(JpgLayerOutput{})
	pulumi.RegisterOutputType(JpgLayerArrayOutput{})
	pulumi.RegisterOutputType(JpgLayerResponseOutput{})
	pulumi.RegisterOutputType(JpgLayerResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingPtrOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingResponseOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointArrayOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointResponseOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(LiveEventInputTypeOutput{})
	pulumi.RegisterOutputType(LiveEventInputTypePtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputAccessControlOutput{})
	pulumi.RegisterOutputType(LiveEventInputAccessControlPtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputAccessControlResponseOutput{})
	pulumi.RegisterOutputType(LiveEventInputAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputResponseOutput{})
	pulumi.RegisterOutputType(LiveEventInputResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputTrackSelectionOutput{})
	pulumi.RegisterOutputType(LiveEventInputTrackSelectionArrayOutput{})
	pulumi.RegisterOutputType(LiveEventInputTrackSelectionResponseOutput{})
	pulumi.RegisterOutputType(LiveEventInputTrackSelectionResponseArrayOutput{})
	pulumi.RegisterOutputType(LiveEventOutputTranscriptionTrackOutput{})
	pulumi.RegisterOutputType(LiveEventOutputTranscriptionTrackPtrOutput{})
	pulumi.RegisterOutputType(LiveEventOutputTranscriptionTrackResponseOutput{})
	pulumi.RegisterOutputType(LiveEventOutputTranscriptionTrackResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewPtrOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlPtrOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlResponseOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewResponseOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventTranscriptionOutput{})
	pulumi.RegisterOutputType(LiveEventTranscriptionArrayOutput{})
	pulumi.RegisterOutputType(LiveEventTranscriptionResponseOutput{})
	pulumi.RegisterOutputType(LiveEventTranscriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(MediaServiceIdentityOutput{})
	pulumi.RegisterOutputType(MediaServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(MediaServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(MediaServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(Mp4FormatOutput{})
	pulumi.RegisterOutputType(Mp4FormatResponseOutput{})
	pulumi.RegisterOutputType(MultiBitrateFormatOutput{})
	pulumi.RegisterOutputType(MultiBitrateFormatResponseOutput{})
	pulumi.RegisterOutputType(NoEncryptionOutput{})
	pulumi.RegisterOutputType(NoEncryptionPtrOutput{})
	pulumi.RegisterOutputType(NoEncryptionResponseOutput{})
	pulumi.RegisterOutputType(NoEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(OutputFileOutput{})
	pulumi.RegisterOutputType(OutputFileArrayOutput{})
	pulumi.RegisterOutputType(OutputFileResponseOutput{})
	pulumi.RegisterOutputType(OutputFileResponseArrayOutput{})
	pulumi.RegisterOutputType(PngFormatOutput{})
	pulumi.RegisterOutputType(PngFormatResponseOutput{})
	pulumi.RegisterOutputType(PngImageOutput{})
	pulumi.RegisterOutputType(PngImageResponseOutput{})
	pulumi.RegisterOutputType(PngLayerOutput{})
	pulumi.RegisterOutputType(PngLayerArrayOutput{})
	pulumi.RegisterOutputType(PngLayerResponseOutput{})
	pulumi.RegisterOutputType(PngLayerResponseArrayOutput{})
	pulumi.RegisterOutputType(PresentationTimeRangeOutput{})
	pulumi.RegisterOutputType(PresentationTimeRangePtrOutput{})
	pulumi.RegisterOutputType(PresentationTimeRangeResponseOutput{})
	pulumi.RegisterOutputType(PresentationTimeRangeResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(RectangleOutput{})
	pulumi.RegisterOutputType(RectanglePtrOutput{})
	pulumi.RegisterOutputType(RectangleResponseOutput{})
	pulumi.RegisterOutputType(RectangleResponsePtrOutput{})
	pulumi.RegisterOutputType(SelectAudioTrackByAttributeOutput{})
	pulumi.RegisterOutputType(SelectAudioTrackByAttributeResponseOutput{})
	pulumi.RegisterOutputType(SelectAudioTrackByIdOutput{})
	pulumi.RegisterOutputType(SelectAudioTrackByIdResponseOutput{})
	pulumi.RegisterOutputType(SelectVideoTrackByAttributeOutput{})
	pulumi.RegisterOutputType(SelectVideoTrackByAttributeResponseOutput{})
	pulumi.RegisterOutputType(SelectVideoTrackByIdOutput{})
	pulumi.RegisterOutputType(SelectVideoTrackByIdResponseOutput{})
	pulumi.RegisterOutputType(StandardEncoderPresetOutput{})
	pulumi.RegisterOutputType(StandardEncoderPresetResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseArrayOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlPtrOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlResponseOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(StreamingLocatorContentKeyOutput{})
	pulumi.RegisterOutputType(StreamingLocatorContentKeyArrayOutput{})
	pulumi.RegisterOutputType(StreamingLocatorContentKeyResponseOutput{})
	pulumi.RegisterOutputType(StreamingLocatorContentKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(StreamingPathResponseOutput{})
	pulumi.RegisterOutputType(StreamingPathResponseArrayOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeyOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeyArrayOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeyResponseOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeysOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeysPtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeysResponseOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeysResponsePtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyFairPlayConfigurationOutput{})
	pulumi.RegisterOutputType(StreamingPolicyFairPlayConfigurationPtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyFairPlayConfigurationResponseOutput{})
	pulumi.RegisterOutputType(StreamingPolicyFairPlayConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyPlayReadyConfigurationOutput{})
	pulumi.RegisterOutputType(StreamingPolicyPlayReadyConfigurationPtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyPlayReadyConfigurationResponseOutput{})
	pulumi.RegisterOutputType(StreamingPolicyPlayReadyConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyWidevineConfigurationOutput{})
	pulumi.RegisterOutputType(StreamingPolicyWidevineConfigurationPtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyWidevineConfigurationResponseOutput{})
	pulumi.RegisterOutputType(StreamingPolicyWidevineConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TrackPropertyConditionOutput{})
	pulumi.RegisterOutputType(TrackPropertyConditionArrayOutput{})
	pulumi.RegisterOutputType(TrackPropertyConditionResponseOutput{})
	pulumi.RegisterOutputType(TrackPropertyConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(TrackSelectionOutput{})
	pulumi.RegisterOutputType(TrackSelectionArrayOutput{})
	pulumi.RegisterOutputType(TrackSelectionResponseOutput{})
	pulumi.RegisterOutputType(TrackSelectionResponseArrayOutput{})
	pulumi.RegisterOutputType(TransformOutputTypeOutput{})
	pulumi.RegisterOutputType(TransformOutputTypeArrayOutput{})
	pulumi.RegisterOutputType(TransformOutputResponseOutput{})
	pulumi.RegisterOutputType(TransformOutputResponseArrayOutput{})
	pulumi.RegisterOutputType(TransportStreamFormatOutput{})
	pulumi.RegisterOutputType(TransportStreamFormatResponseOutput{})
	pulumi.RegisterOutputType(UtcClipTimeOutput{})
	pulumi.RegisterOutputType(UtcClipTimeResponseOutput{})
	pulumi.RegisterOutputType(VideoOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerPresetOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerPresetResponseOutput{})
	pulumi.RegisterOutputType(VideoOverlayOutput{})
	pulumi.RegisterOutputType(VideoOverlayResponseOutput{})
	pulumi.RegisterOutputType(VideoResponseOutput{})
	pulumi.RegisterOutputType(VideoTrackDescriptorOutput{})
	pulumi.RegisterOutputType(VideoTrackDescriptorResponseOutput{})
}
