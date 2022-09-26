


package v20210501preview

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

func (o AccountEncryptionOutput) Identity() ResourceIdentityPtrOutput {
	return o.ApplyT(func(v AccountEncryption) *ResourceIdentity { return v.Identity }).(ResourceIdentityPtrOutput)
}

func (o AccountEncryptionOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v AccountEncryption) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

func (o AccountEncryptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AccountEncryption) string { return v.Type }).(pulumi.StringOutput)
}

type AccountEncryptionResponse struct {
	Identity           *ResourceIdentityResponse   `pulumi:"identity"`
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Status             string                      `pulumi:"status"`
	Type               string                      `pulumi:"type"`
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

type EccTokenKey struct {
	Alg  string `pulumi:"alg"`
	Kid  string `pulumi:"kid"`
	Type string `pulumi:"type"`
	X    string `pulumi:"x"`
	Y    string `pulumi:"y"`
}

type EccTokenKeyResponse struct {
	Alg  string `pulumi:"alg"`
	Kid  string `pulumi:"kid"`
	Type string `pulumi:"type"`
	X    string `pulumi:"x"`
	Y    string `pulumi:"y"`
}

type EndpointResponse struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
	Type        string  `pulumi:"type"`
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

type RsaTokenKeyResponse struct {
	Alg  string `pulumi:"alg"`
	E    string `pulumi:"e"`
	Kid  string `pulumi:"kid"`
	N    string `pulumi:"n"`
	Type string `pulumi:"type"`
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

type UserAssignedManagedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
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

type VideoFlagsResponse struct {
	CanStream   bool `pulumi:"canStream"`
	HasData     bool `pulumi:"hasData"`
	IsRecording bool `pulumi:"isRecording"`
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

func (o VideoFlagsResponseOutput) CanStream() pulumi.BoolOutput {
	return o.ApplyT(func(v VideoFlagsResponse) bool { return v.CanStream }).(pulumi.BoolOutput)
}

func (o VideoFlagsResponseOutput) HasData() pulumi.BoolOutput {
	return o.ApplyT(func(v VideoFlagsResponse) bool { return v.HasData }).(pulumi.BoolOutput)
}

func (o VideoFlagsResponseOutput) IsRecording() pulumi.BoolOutput {
	return o.ApplyT(func(v VideoFlagsResponse) bool { return v.IsRecording }).(pulumi.BoolOutput)
}

type VideoMediaInfoResponse struct {
	SegmentLength string `pulumi:"segmentLength"`
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

func (o VideoMediaInfoResponseOutput) SegmentLength() pulumi.StringOutput {
	return o.ApplyT(func(v VideoMediaInfoResponse) string { return v.SegmentLength }).(pulumi.StringOutput)
}

type VideoStreamingResponse struct {
	ArchiveBaseUrl *string `pulumi:"archiveBaseUrl"`
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

func (o VideoStreamingResponseOutput) ArchiveBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VideoStreamingResponse) *string { return v.ArchiveBaseUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountEncryptionOutput{})
	pulumi.RegisterOutputType(AccountEncryptionResponseOutput{})
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
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TokenClaimOutput{})
	pulumi.RegisterOutputType(TokenClaimArrayOutput{})
	pulumi.RegisterOutputType(TokenClaimResponseOutput{})
	pulumi.RegisterOutputType(TokenClaimResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAssignedManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedManagedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityPtrOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityResponseOutput{})
	pulumi.RegisterOutputType(VideoAnalyzerIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(VideoFlagsResponseOutput{})
	pulumi.RegisterOutputType(VideoMediaInfoResponseOutput{})
	pulumi.RegisterOutputType(VideoStreamingResponseOutput{})
}
