


package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AttestationServiceCreationSpecificParams struct {
	PolicySigningCertificates *JSONWebKeySet `pulumi:"policySigningCertificates"`
}





type AttestationServiceCreationSpecificParamsInput interface {
	pulumi.Input

	ToAttestationServiceCreationSpecificParamsOutput() AttestationServiceCreationSpecificParamsOutput
	ToAttestationServiceCreationSpecificParamsOutputWithContext(context.Context) AttestationServiceCreationSpecificParamsOutput
}

type AttestationServiceCreationSpecificParamsArgs struct {
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

func (i AttestationServiceCreationSpecificParamsArgs) ToAttestationServiceCreationSpecificParamsPtrOutput() AttestationServiceCreationSpecificParamsPtrOutput {
	return i.ToAttestationServiceCreationSpecificParamsPtrOutputWithContext(context.Background())
}

func (i AttestationServiceCreationSpecificParamsArgs) ToAttestationServiceCreationSpecificParamsPtrOutputWithContext(ctx context.Context) AttestationServiceCreationSpecificParamsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationServiceCreationSpecificParamsOutput).ToAttestationServiceCreationSpecificParamsPtrOutputWithContext(ctx)
}









type AttestationServiceCreationSpecificParamsPtrInput interface {
	pulumi.Input

	ToAttestationServiceCreationSpecificParamsPtrOutput() AttestationServiceCreationSpecificParamsPtrOutput
	ToAttestationServiceCreationSpecificParamsPtrOutputWithContext(context.Context) AttestationServiceCreationSpecificParamsPtrOutput
}

type attestationServiceCreationSpecificParamsPtrType AttestationServiceCreationSpecificParamsArgs

func AttestationServiceCreationSpecificParamsPtr(v *AttestationServiceCreationSpecificParamsArgs) AttestationServiceCreationSpecificParamsPtrInput {
	return (*attestationServiceCreationSpecificParamsPtrType)(v)
}

func (*attestationServiceCreationSpecificParamsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationServiceCreationSpecificParams)(nil)).Elem()
}

func (i *attestationServiceCreationSpecificParamsPtrType) ToAttestationServiceCreationSpecificParamsPtrOutput() AttestationServiceCreationSpecificParamsPtrOutput {
	return i.ToAttestationServiceCreationSpecificParamsPtrOutputWithContext(context.Background())
}

func (i *attestationServiceCreationSpecificParamsPtrType) ToAttestationServiceCreationSpecificParamsPtrOutputWithContext(ctx context.Context) AttestationServiceCreationSpecificParamsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationServiceCreationSpecificParamsPtrOutput)
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

func (o AttestationServiceCreationSpecificParamsOutput) ToAttestationServiceCreationSpecificParamsPtrOutput() AttestationServiceCreationSpecificParamsPtrOutput {
	return o.ToAttestationServiceCreationSpecificParamsPtrOutputWithContext(context.Background())
}

func (o AttestationServiceCreationSpecificParamsOutput) ToAttestationServiceCreationSpecificParamsPtrOutputWithContext(ctx context.Context) AttestationServiceCreationSpecificParamsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AttestationServiceCreationSpecificParams) *AttestationServiceCreationSpecificParams {
		return &v
	}).(AttestationServiceCreationSpecificParamsPtrOutput)
}

func (o AttestationServiceCreationSpecificParamsOutput) PolicySigningCertificates() JSONWebKeySetPtrOutput {
	return o.ApplyT(func(v AttestationServiceCreationSpecificParams) *JSONWebKeySet { return v.PolicySigningCertificates }).(JSONWebKeySetPtrOutput)
}

type AttestationServiceCreationSpecificParamsPtrOutput struct{ *pulumi.OutputState }

func (AttestationServiceCreationSpecificParamsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationServiceCreationSpecificParams)(nil)).Elem()
}

func (o AttestationServiceCreationSpecificParamsPtrOutput) ToAttestationServiceCreationSpecificParamsPtrOutput() AttestationServiceCreationSpecificParamsPtrOutput {
	return o
}

func (o AttestationServiceCreationSpecificParamsPtrOutput) ToAttestationServiceCreationSpecificParamsPtrOutputWithContext(ctx context.Context) AttestationServiceCreationSpecificParamsPtrOutput {
	return o
}

func (o AttestationServiceCreationSpecificParamsPtrOutput) Elem() AttestationServiceCreationSpecificParamsOutput {
	return o.ApplyT(func(v *AttestationServiceCreationSpecificParams) AttestationServiceCreationSpecificParams {
		if v != nil {
			return *v
		}
		var ret AttestationServiceCreationSpecificParams
		return ret
	}).(AttestationServiceCreationSpecificParamsOutput)
}

func (o AttestationServiceCreationSpecificParamsPtrOutput) PolicySigningCertificates() JSONWebKeySetPtrOutput {
	return o.ApplyT(func(v *AttestationServiceCreationSpecificParams) *JSONWebKeySet {
		if v == nil {
			return nil
		}
		return v.PolicySigningCertificates
	}).(JSONWebKeySetPtrOutput)
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

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                             `pulumi:"id"`
	Name                              pulumi.StringInput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                             `pulumi:"provisioningState"`
	Type                              pulumi.StringInput                             `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
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

func init() {
	pulumi.RegisterOutputType(AttestationServiceCreationSpecificParamsOutput{})
	pulumi.RegisterOutputType(AttestationServiceCreationSpecificParamsPtrOutput{})
	pulumi.RegisterOutputType(JSONWebKeyOutput{})
	pulumi.RegisterOutputType(JSONWebKeyArrayOutput{})
	pulumi.RegisterOutputType(JSONWebKeySetOutput{})
	pulumi.RegisterOutputType(JSONWebKeySetPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
