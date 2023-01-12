


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttestationServiceCreationSpecificParams struct {
	AttestationPolicy         *string        `pulumi:"attestationPolicy"`
	PolicySigningCertificates *JSONWebKeySet `pulumi:"policySigningCertificates"`
}





type AttestationServiceCreationSpecificParamsInput interface {
	pulumi.Input

	ToAttestationServiceCreationSpecificParamsOutput() AttestationServiceCreationSpecificParamsOutput
	ToAttestationServiceCreationSpecificParamsOutputWithContext(context.Context) AttestationServiceCreationSpecificParamsOutput
}

type AttestationServiceCreationSpecificParamsArgs struct {
	AttestationPolicy         pulumi.StringPtrInput `pulumi:"attestationPolicy"`
	PolicySigningCertificates JSONWebKeySetPtrInput `pulumi:"policySigningCertificates"`
}

func (AttestationServiceCreationSpecificParamsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestationServiceCreationSpecificParams)(nil)).Elem()
}

func (i AttestationServiceCreationSpecificParamsArgs) ToAttestationServiceCreationSpecificParamsOutput() AttestationServiceCreationSpecificParamsOutput {
	return i.ToAttestationServiceCreationSpecificParamsOutputWithContext(context.Background())
}

func (i AttestationServiceCreationSpecificParamsArgs) ToAttestationServiceCreationSpecificParamsOutputWithContext(ctx context.Context) AttestationServiceCreationSpecificParamsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationServiceCreationSpecificParamsOutput)
}

type AttestationServiceCreationSpecificParamsOutput struct{ *pulumi.OutputState }

func (AttestationServiceCreationSpecificParamsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestationServiceCreationSpecificParams)(nil)).Elem()
}

func (o AttestationServiceCreationSpecificParamsOutput) ToAttestationServiceCreationSpecificParamsOutput() AttestationServiceCreationSpecificParamsOutput {
	return o
}

func (o AttestationServiceCreationSpecificParamsOutput) ToAttestationServiceCreationSpecificParamsOutputWithContext(ctx context.Context) AttestationServiceCreationSpecificParamsOutput {
	return o
}

func (o AttestationServiceCreationSpecificParamsOutput) AttestationPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttestationServiceCreationSpecificParams) *string { return v.AttestationPolicy }).(pulumi.StringPtrOutput)
}

func (o AttestationServiceCreationSpecificParamsOutput) PolicySigningCertificates() JSONWebKeySetPtrOutput {
	return o.ApplyT(func(v AttestationServiceCreationSpecificParams) *JSONWebKeySet { return v.PolicySigningCertificates }).(JSONWebKeySetPtrOutput)
}

type JSONWebKey struct {
	Alg *string  `pulumi:"alg"`
	Crv *string  `pulumi:"crv"`
	D   *string  `pulumi:"d"`
	Dp  *string  `pulumi:"dp"`
	Dq  *string  `pulumi:"dq"`
	E   *string  `pulumi:"e"`
	K   *string  `pulumi:"k"`
	Kid *string  `pulumi:"kid"`
	Kty string   `pulumi:"kty"`
	N   *string  `pulumi:"n"`
	P   *string  `pulumi:"p"`
	Q   *string  `pulumi:"q"`
	Qi  *string  `pulumi:"qi"`
	Use *string  `pulumi:"use"`
	X   *string  `pulumi:"x"`
	X5c []string `pulumi:"x5c"`
	Y   *string  `pulumi:"y"`
}





type JSONWebKeyInput interface {
	pulumi.Input

	ToJSONWebKeyOutput() JSONWebKeyOutput
	ToJSONWebKeyOutputWithContext(context.Context) JSONWebKeyOutput
}

type JSONWebKeyArgs struct {
	Alg pulumi.StringPtrInput   `pulumi:"alg"`
	Crv pulumi.StringPtrInput   `pulumi:"crv"`
	D   pulumi.StringPtrInput   `pulumi:"d"`
	Dp  pulumi.StringPtrInput   `pulumi:"dp"`
	Dq  pulumi.StringPtrInput   `pulumi:"dq"`
	E   pulumi.StringPtrInput   `pulumi:"e"`
	K   pulumi.StringPtrInput   `pulumi:"k"`
	Kid pulumi.StringPtrInput   `pulumi:"kid"`
	Kty pulumi.StringInput      `pulumi:"kty"`
	N   pulumi.StringPtrInput   `pulumi:"n"`
	P   pulumi.StringPtrInput   `pulumi:"p"`
	Q   pulumi.StringPtrInput   `pulumi:"q"`
	Qi  pulumi.StringPtrInput   `pulumi:"qi"`
	Use pulumi.StringPtrInput   `pulumi:"use"`
	X   pulumi.StringPtrInput   `pulumi:"x"`
	X5c pulumi.StringArrayInput `pulumi:"x5c"`
	Y   pulumi.StringPtrInput   `pulumi:"y"`
}

func (JSONWebKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JSONWebKey)(nil)).Elem()
}

func (i JSONWebKeyArgs) ToJSONWebKeyOutput() JSONWebKeyOutput {
	return i.ToJSONWebKeyOutputWithContext(context.Background())
}

func (i JSONWebKeyArgs) ToJSONWebKeyOutputWithContext(ctx context.Context) JSONWebKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JSONWebKeyOutput)
}





type JSONWebKeyArrayInput interface {
	pulumi.Input

	ToJSONWebKeyArrayOutput() JSONWebKeyArrayOutput
	ToJSONWebKeyArrayOutputWithContext(context.Context) JSONWebKeyArrayOutput
}

type JSONWebKeyArray []JSONWebKeyInput

func (JSONWebKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JSONWebKey)(nil)).Elem()
}

func (i JSONWebKeyArray) ToJSONWebKeyArrayOutput() JSONWebKeyArrayOutput {
	return i.ToJSONWebKeyArrayOutputWithContext(context.Background())
}

func (i JSONWebKeyArray) ToJSONWebKeyArrayOutputWithContext(ctx context.Context) JSONWebKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JSONWebKeyArrayOutput)
}

type JSONWebKeyOutput struct{ *pulumi.OutputState }

func (JSONWebKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JSONWebKey)(nil)).Elem()
}

func (o JSONWebKeyOutput) ToJSONWebKeyOutput() JSONWebKeyOutput {
	return o
}

func (o JSONWebKeyOutput) ToJSONWebKeyOutputWithContext(ctx context.Context) JSONWebKeyOutput {
	return o
}

func (o JSONWebKeyOutput) Alg() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.Alg }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) Crv() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.Crv }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) D() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.D }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) Dp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.Dp }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) Dq() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.Dq }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) E() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.E }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) K() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.K }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) Kid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.Kid }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) Kty() pulumi.StringOutput {
	return o.ApplyT(func(v JSONWebKey) string { return v.Kty }).(pulumi.StringOutput)
}

func (o JSONWebKeyOutput) N() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.N }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) P() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.P }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) Q() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.Q }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) Qi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.Qi }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) Use() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.Use }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) X() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.X }).(pulumi.StringPtrOutput)
}

func (o JSONWebKeyOutput) X5c() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JSONWebKey) []string { return v.X5c }).(pulumi.StringArrayOutput)
}

func (o JSONWebKeyOutput) Y() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JSONWebKey) *string { return v.Y }).(pulumi.StringPtrOutput)
}

type JSONWebKeyArrayOutput struct{ *pulumi.OutputState }

func (JSONWebKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JSONWebKey)(nil)).Elem()
}

func (o JSONWebKeyArrayOutput) ToJSONWebKeyArrayOutput() JSONWebKeyArrayOutput {
	return o
}

func (o JSONWebKeyArrayOutput) ToJSONWebKeyArrayOutputWithContext(ctx context.Context) JSONWebKeyArrayOutput {
	return o
}

func (o JSONWebKeyArrayOutput) Index(i pulumi.IntInput) JSONWebKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JSONWebKey {
		return vs[0].([]JSONWebKey)[vs[1].(int)]
	}).(JSONWebKeyOutput)
}

type JSONWebKeySet struct {
	Keys []JSONWebKey `pulumi:"keys"`
}





type JSONWebKeySetInput interface {
	pulumi.Input

	ToJSONWebKeySetOutput() JSONWebKeySetOutput
	ToJSONWebKeySetOutputWithContext(context.Context) JSONWebKeySetOutput
}

type JSONWebKeySetArgs struct {
	Keys JSONWebKeyArrayInput `pulumi:"keys"`
}

func (JSONWebKeySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JSONWebKeySet)(nil)).Elem()
}

func (i JSONWebKeySetArgs) ToJSONWebKeySetOutput() JSONWebKeySetOutput {
	return i.ToJSONWebKeySetOutputWithContext(context.Background())
}

func (i JSONWebKeySetArgs) ToJSONWebKeySetOutputWithContext(ctx context.Context) JSONWebKeySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JSONWebKeySetOutput)
}

func (i JSONWebKeySetArgs) ToJSONWebKeySetPtrOutput() JSONWebKeySetPtrOutput {
	return i.ToJSONWebKeySetPtrOutputWithContext(context.Background())
}

func (i JSONWebKeySetArgs) ToJSONWebKeySetPtrOutputWithContext(ctx context.Context) JSONWebKeySetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JSONWebKeySetOutput).ToJSONWebKeySetPtrOutputWithContext(ctx)
}









type JSONWebKeySetPtrInput interface {
	pulumi.Input

	ToJSONWebKeySetPtrOutput() JSONWebKeySetPtrOutput
	ToJSONWebKeySetPtrOutputWithContext(context.Context) JSONWebKeySetPtrOutput
}

type jsonwebKeySetPtrType JSONWebKeySetArgs

func JSONWebKeySetPtr(v *JSONWebKeySetArgs) JSONWebKeySetPtrInput {
	return (*jsonwebKeySetPtrType)(v)
}

func (*jsonwebKeySetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JSONWebKeySet)(nil)).Elem()
}

func (i *jsonwebKeySetPtrType) ToJSONWebKeySetPtrOutput() JSONWebKeySetPtrOutput {
	return i.ToJSONWebKeySetPtrOutputWithContext(context.Background())
}

func (i *jsonwebKeySetPtrType) ToJSONWebKeySetPtrOutputWithContext(ctx context.Context) JSONWebKeySetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JSONWebKeySetPtrOutput)
}

type JSONWebKeySetOutput struct{ *pulumi.OutputState }

func (JSONWebKeySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JSONWebKeySet)(nil)).Elem()
}

func (o JSONWebKeySetOutput) ToJSONWebKeySetOutput() JSONWebKeySetOutput {
	return o
}

func (o JSONWebKeySetOutput) ToJSONWebKeySetOutputWithContext(ctx context.Context) JSONWebKeySetOutput {
	return o
}

func (o JSONWebKeySetOutput) ToJSONWebKeySetPtrOutput() JSONWebKeySetPtrOutput {
	return o.ToJSONWebKeySetPtrOutputWithContext(context.Background())
}

func (o JSONWebKeySetOutput) ToJSONWebKeySetPtrOutputWithContext(ctx context.Context) JSONWebKeySetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JSONWebKeySet) *JSONWebKeySet {
		return &v
	}).(JSONWebKeySetPtrOutput)
}

func (o JSONWebKeySetOutput) Keys() JSONWebKeyArrayOutput {
	return o.ApplyT(func(v JSONWebKeySet) []JSONWebKey { return v.Keys }).(JSONWebKeyArrayOutput)
}

type JSONWebKeySetPtrOutput struct{ *pulumi.OutputState }

func (JSONWebKeySetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JSONWebKeySet)(nil)).Elem()
}

func (o JSONWebKeySetPtrOutput) ToJSONWebKeySetPtrOutput() JSONWebKeySetPtrOutput {
	return o
}

func (o JSONWebKeySetPtrOutput) ToJSONWebKeySetPtrOutputWithContext(ctx context.Context) JSONWebKeySetPtrOutput {
	return o
}

func (o JSONWebKeySetPtrOutput) Elem() JSONWebKeySetOutput {
	return o.ApplyT(func(v *JSONWebKeySet) JSONWebKeySet {
		if v != nil {
			return *v
		}
		var ret JSONWebKeySet
		return ret
	}).(JSONWebKeySetOutput)
}

func (o JSONWebKeySetPtrOutput) Keys() JSONWebKeyArrayOutput {
	return o.ApplyT(func(v *JSONWebKeySet) []JSONWebKey {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(JSONWebKeyArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AttestationServiceCreationSpecificParamsOutput{})
	pulumi.RegisterOutputType(JSONWebKeyOutput{})
	pulumi.RegisterOutputType(JSONWebKeyArrayOutput{})
	pulumi.RegisterOutputType(JSONWebKeySetOutput{})
	pulumi.RegisterOutputType(JSONWebKeySetPtrOutput{})
}
