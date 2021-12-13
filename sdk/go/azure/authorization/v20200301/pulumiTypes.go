


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Identity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
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

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
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

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
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

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
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
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ParameterDefinitionsValue struct {
	AllowedValues []interface{}                      `pulumi:"allowedValues"`
	DefaultValue  interface{}                        `pulumi:"defaultValue"`
	Metadata      *ParameterDefinitionsValueMetadata `pulumi:"metadata"`
	Type          *string                            `pulumi:"type"`
}





type ParameterDefinitionsValueInput interface {
	pulumi.Input

	ToParameterDefinitionsValueOutput() ParameterDefinitionsValueOutput
	ToParameterDefinitionsValueOutputWithContext(context.Context) ParameterDefinitionsValueOutput
}

type ParameterDefinitionsValueArgs struct {
	AllowedValues pulumi.ArrayInput                         `pulumi:"allowedValues"`
	DefaultValue  pulumi.Input                              `pulumi:"defaultValue"`
	Metadata      ParameterDefinitionsValueMetadataPtrInput `pulumi:"metadata"`
	Type          pulumi.StringPtrInput                     `pulumi:"type"`
}

func (ParameterDefinitionsValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValue)(nil)).Elem()
}

func (i ParameterDefinitionsValueArgs) ToParameterDefinitionsValueOutput() ParameterDefinitionsValueOutput {
	return i.ToParameterDefinitionsValueOutputWithContext(context.Background())
}

func (i ParameterDefinitionsValueArgs) ToParameterDefinitionsValueOutputWithContext(ctx context.Context) ParameterDefinitionsValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionsValueOutput)
}





type ParameterDefinitionsValueMapInput interface {
	pulumi.Input

	ToParameterDefinitionsValueMapOutput() ParameterDefinitionsValueMapOutput
	ToParameterDefinitionsValueMapOutputWithContext(context.Context) ParameterDefinitionsValueMapOutput
}

type ParameterDefinitionsValueMap map[string]ParameterDefinitionsValueInput

func (ParameterDefinitionsValueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinitionsValue)(nil)).Elem()
}

func (i ParameterDefinitionsValueMap) ToParameterDefinitionsValueMapOutput() ParameterDefinitionsValueMapOutput {
	return i.ToParameterDefinitionsValueMapOutputWithContext(context.Background())
}

func (i ParameterDefinitionsValueMap) ToParameterDefinitionsValueMapOutputWithContext(ctx context.Context) ParameterDefinitionsValueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionsValueMapOutput)
}

type ParameterDefinitionsValueOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValue)(nil)).Elem()
}

func (o ParameterDefinitionsValueOutput) ToParameterDefinitionsValueOutput() ParameterDefinitionsValueOutput {
	return o
}

func (o ParameterDefinitionsValueOutput) ToParameterDefinitionsValueOutputWithContext(ctx context.Context) ParameterDefinitionsValueOutput {
	return o
}

func (o ParameterDefinitionsValueOutput) AllowedValues() pulumi.ArrayOutput {
	return o.ApplyT(func(v ParameterDefinitionsValue) []interface{} { return v.AllowedValues }).(pulumi.ArrayOutput)
}

func (o ParameterDefinitionsValueOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterDefinitionsValue) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterDefinitionsValueOutput) Metadata() ParameterDefinitionsValueMetadataPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValue) *ParameterDefinitionsValueMetadata { return v.Metadata }).(ParameterDefinitionsValueMetadataPtrOutput)
}

func (o ParameterDefinitionsValueOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValue) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionsValueMapOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinitionsValue)(nil)).Elem()
}

func (o ParameterDefinitionsValueMapOutput) ToParameterDefinitionsValueMapOutput() ParameterDefinitionsValueMapOutput {
	return o
}

func (o ParameterDefinitionsValueMapOutput) ToParameterDefinitionsValueMapOutputWithContext(ctx context.Context) ParameterDefinitionsValueMapOutput {
	return o
}

func (o ParameterDefinitionsValueMapOutput) MapIndex(k pulumi.StringInput) ParameterDefinitionsValueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterDefinitionsValue {
		return vs[0].(map[string]ParameterDefinitionsValue)[vs[1].(string)]
	}).(ParameterDefinitionsValueOutput)
}

type ParameterDefinitionsValueMetadata struct {
	Description *string `pulumi:"description"`
	DisplayName *string `pulumi:"displayName"`
}





type ParameterDefinitionsValueMetadataInput interface {
	pulumi.Input

	ToParameterDefinitionsValueMetadataOutput() ParameterDefinitionsValueMetadataOutput
	ToParameterDefinitionsValueMetadataOutputWithContext(context.Context) ParameterDefinitionsValueMetadataOutput
}

type ParameterDefinitionsValueMetadataArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
}

func (ParameterDefinitionsValueMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValueMetadata)(nil)).Elem()
}

func (i ParameterDefinitionsValueMetadataArgs) ToParameterDefinitionsValueMetadataOutput() ParameterDefinitionsValueMetadataOutput {
	return i.ToParameterDefinitionsValueMetadataOutputWithContext(context.Background())
}

func (i ParameterDefinitionsValueMetadataArgs) ToParameterDefinitionsValueMetadataOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionsValueMetadataOutput)
}

func (i ParameterDefinitionsValueMetadataArgs) ToParameterDefinitionsValueMetadataPtrOutput() ParameterDefinitionsValueMetadataPtrOutput {
	return i.ToParameterDefinitionsValueMetadataPtrOutputWithContext(context.Background())
}

func (i ParameterDefinitionsValueMetadataArgs) ToParameterDefinitionsValueMetadataPtrOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionsValueMetadataOutput).ToParameterDefinitionsValueMetadataPtrOutputWithContext(ctx)
}









type ParameterDefinitionsValueMetadataPtrInput interface {
	pulumi.Input

	ToParameterDefinitionsValueMetadataPtrOutput() ParameterDefinitionsValueMetadataPtrOutput
	ToParameterDefinitionsValueMetadataPtrOutputWithContext(context.Context) ParameterDefinitionsValueMetadataPtrOutput
}

type parameterDefinitionsValueMetadataPtrType ParameterDefinitionsValueMetadataArgs

func ParameterDefinitionsValueMetadataPtr(v *ParameterDefinitionsValueMetadataArgs) ParameterDefinitionsValueMetadataPtrInput {
	return (*parameterDefinitionsValueMetadataPtrType)(v)
}

func (*parameterDefinitionsValueMetadataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterDefinitionsValueMetadata)(nil)).Elem()
}

func (i *parameterDefinitionsValueMetadataPtrType) ToParameterDefinitionsValueMetadataPtrOutput() ParameterDefinitionsValueMetadataPtrOutput {
	return i.ToParameterDefinitionsValueMetadataPtrOutputWithContext(context.Background())
}

func (i *parameterDefinitionsValueMetadataPtrType) ToParameterDefinitionsValueMetadataPtrOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionsValueMetadataPtrOutput)
}

type ParameterDefinitionsValueMetadataOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValueMetadata)(nil)).Elem()
}

func (o ParameterDefinitionsValueMetadataOutput) ToParameterDefinitionsValueMetadataOutput() ParameterDefinitionsValueMetadataOutput {
	return o
}

func (o ParameterDefinitionsValueMetadataOutput) ToParameterDefinitionsValueMetadataOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataOutput {
	return o
}

func (o ParameterDefinitionsValueMetadataOutput) ToParameterDefinitionsValueMetadataPtrOutput() ParameterDefinitionsValueMetadataPtrOutput {
	return o.ToParameterDefinitionsValueMetadataPtrOutputWithContext(context.Background())
}

func (o ParameterDefinitionsValueMetadataOutput) ToParameterDefinitionsValueMetadataPtrOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParameterDefinitionsValueMetadata) *ParameterDefinitionsValueMetadata {
		return &v
	}).(ParameterDefinitionsValueMetadataPtrOutput)
}

func (o ParameterDefinitionsValueMetadataOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueMetadata) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueMetadataOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueMetadata) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionsValueMetadataPtrOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueMetadataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterDefinitionsValueMetadata)(nil)).Elem()
}

func (o ParameterDefinitionsValueMetadataPtrOutput) ToParameterDefinitionsValueMetadataPtrOutput() ParameterDefinitionsValueMetadataPtrOutput {
	return o
}

func (o ParameterDefinitionsValueMetadataPtrOutput) ToParameterDefinitionsValueMetadataPtrOutputWithContext(ctx context.Context) ParameterDefinitionsValueMetadataPtrOutput {
	return o
}

func (o ParameterDefinitionsValueMetadataPtrOutput) Elem() ParameterDefinitionsValueMetadataOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueMetadata) ParameterDefinitionsValueMetadata {
		if v != nil {
			return *v
		}
		var ret ParameterDefinitionsValueMetadata
		return ret
	}).(ParameterDefinitionsValueMetadataOutput)
}

func (o ParameterDefinitionsValueMetadataPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueMetadata) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueMetadataPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueMetadata) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

type ParameterDefinitionsValueResponse struct {
	AllowedValues []interface{}                              `pulumi:"allowedValues"`
	DefaultValue  interface{}                                `pulumi:"defaultValue"`
	Metadata      *ParameterDefinitionsValueResponseMetadata `pulumi:"metadata"`
	Type          *string                                    `pulumi:"type"`
}

type ParameterDefinitionsValueResponseOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValueResponse)(nil)).Elem()
}

func (o ParameterDefinitionsValueResponseOutput) ToParameterDefinitionsValueResponseOutput() ParameterDefinitionsValueResponseOutput {
	return o
}

func (o ParameterDefinitionsValueResponseOutput) ToParameterDefinitionsValueResponseOutputWithContext(ctx context.Context) ParameterDefinitionsValueResponseOutput {
	return o
}

func (o ParameterDefinitionsValueResponseOutput) AllowedValues() pulumi.ArrayOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponse) []interface{} { return v.AllowedValues }).(pulumi.ArrayOutput)
}

func (o ParameterDefinitionsValueResponseOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponse) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterDefinitionsValueResponseOutput) Metadata() ParameterDefinitionsValueResponseMetadataPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponse) *ParameterDefinitionsValueResponseMetadata {
		return v.Metadata
	}).(ParameterDefinitionsValueResponseMetadataPtrOutput)
}

func (o ParameterDefinitionsValueResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionsValueResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinitionsValueResponse)(nil)).Elem()
}

func (o ParameterDefinitionsValueResponseMapOutput) ToParameterDefinitionsValueResponseMapOutput() ParameterDefinitionsValueResponseMapOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMapOutput) ToParameterDefinitionsValueResponseMapOutputWithContext(ctx context.Context) ParameterDefinitionsValueResponseMapOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterDefinitionsValueResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterDefinitionsValueResponse {
		return vs[0].(map[string]ParameterDefinitionsValueResponse)[vs[1].(string)]
	}).(ParameterDefinitionsValueResponseOutput)
}

type ParameterDefinitionsValueResponseMetadata struct {
	Description *string `pulumi:"description"`
	DisplayName *string `pulumi:"displayName"`
}

type ParameterDefinitionsValueResponseMetadataOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueResponseMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionsValueResponseMetadata)(nil)).Elem()
}

func (o ParameterDefinitionsValueResponseMetadataOutput) ToParameterDefinitionsValueResponseMetadataOutput() ParameterDefinitionsValueResponseMetadataOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMetadataOutput) ToParameterDefinitionsValueResponseMetadataOutputWithContext(ctx context.Context) ParameterDefinitionsValueResponseMetadataOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMetadataOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponseMetadata) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueResponseMetadataOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionsValueResponseMetadata) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

type ParameterDefinitionsValueResponseMetadataPtrOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionsValueResponseMetadataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterDefinitionsValueResponseMetadata)(nil)).Elem()
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) ToParameterDefinitionsValueResponseMetadataPtrOutput() ParameterDefinitionsValueResponseMetadataPtrOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) ToParameterDefinitionsValueResponseMetadataPtrOutputWithContext(ctx context.Context) ParameterDefinitionsValueResponseMetadataPtrOutput {
	return o
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) Elem() ParameterDefinitionsValueResponseMetadataOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueResponseMetadata) ParameterDefinitionsValueResponseMetadata {
		if v != nil {
			return *v
		}
		var ret ParameterDefinitionsValueResponseMetadata
		return ret
	}).(ParameterDefinitionsValueResponseMetadataOutput)
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueResponseMetadata) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionsValueResponseMetadataPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterDefinitionsValueResponseMetadata) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

type ParameterValuesValue struct {
	Value interface{} `pulumi:"value"`
}





type ParameterValuesValueInput interface {
	pulumi.Input

	ToParameterValuesValueOutput() ParameterValuesValueOutput
	ToParameterValuesValueOutputWithContext(context.Context) ParameterValuesValueOutput
}

type ParameterValuesValueArgs struct {
	Value pulumi.Input `pulumi:"value"`
}

func (ParameterValuesValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValuesValue)(nil)).Elem()
}

func (i ParameterValuesValueArgs) ToParameterValuesValueOutput() ParameterValuesValueOutput {
	return i.ToParameterValuesValueOutputWithContext(context.Background())
}

func (i ParameterValuesValueArgs) ToParameterValuesValueOutputWithContext(ctx context.Context) ParameterValuesValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterValuesValueOutput)
}





type ParameterValuesValueMapInput interface {
	pulumi.Input

	ToParameterValuesValueMapOutput() ParameterValuesValueMapOutput
	ToParameterValuesValueMapOutputWithContext(context.Context) ParameterValuesValueMapOutput
}

type ParameterValuesValueMap map[string]ParameterValuesValueInput

func (ParameterValuesValueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValuesValue)(nil)).Elem()
}

func (i ParameterValuesValueMap) ToParameterValuesValueMapOutput() ParameterValuesValueMapOutput {
	return i.ToParameterValuesValueMapOutputWithContext(context.Background())
}

func (i ParameterValuesValueMap) ToParameterValuesValueMapOutputWithContext(ctx context.Context) ParameterValuesValueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterValuesValueMapOutput)
}

type ParameterValuesValueOutput struct{ *pulumi.OutputState }

func (ParameterValuesValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValuesValue)(nil)).Elem()
}

func (o ParameterValuesValueOutput) ToParameterValuesValueOutput() ParameterValuesValueOutput {
	return o
}

func (o ParameterValuesValueOutput) ToParameterValuesValueOutputWithContext(ctx context.Context) ParameterValuesValueOutput {
	return o
}

func (o ParameterValuesValueOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterValuesValue) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ParameterValuesValueMapOutput struct{ *pulumi.OutputState }

func (ParameterValuesValueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValuesValue)(nil)).Elem()
}

func (o ParameterValuesValueMapOutput) ToParameterValuesValueMapOutput() ParameterValuesValueMapOutput {
	return o
}

func (o ParameterValuesValueMapOutput) ToParameterValuesValueMapOutputWithContext(ctx context.Context) ParameterValuesValueMapOutput {
	return o
}

func (o ParameterValuesValueMapOutput) MapIndex(k pulumi.StringInput) ParameterValuesValueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterValuesValue {
		return vs[0].(map[string]ParameterValuesValue)[vs[1].(string)]
	}).(ParameterValuesValueOutput)
}

type ParameterValuesValueResponse struct {
	Value interface{} `pulumi:"value"`
}

type ParameterValuesValueResponseOutput struct{ *pulumi.OutputState }

func (ParameterValuesValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValuesValueResponse)(nil)).Elem()
}

func (o ParameterValuesValueResponseOutput) ToParameterValuesValueResponseOutput() ParameterValuesValueResponseOutput {
	return o
}

func (o ParameterValuesValueResponseOutput) ToParameterValuesValueResponseOutputWithContext(ctx context.Context) ParameterValuesValueResponseOutput {
	return o
}

func (o ParameterValuesValueResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterValuesValueResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ParameterValuesValueResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterValuesValueResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValuesValueResponse)(nil)).Elem()
}

func (o ParameterValuesValueResponseMapOutput) ToParameterValuesValueResponseMapOutput() ParameterValuesValueResponseMapOutput {
	return o
}

func (o ParameterValuesValueResponseMapOutput) ToParameterValuesValueResponseMapOutputWithContext(ctx context.Context) ParameterValuesValueResponseMapOutput {
	return o
}

func (o ParameterValuesValueResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterValuesValueResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterValuesValueResponse {
		return vs[0].(map[string]ParameterValuesValueResponse)[vs[1].(string)]
	}).(ParameterValuesValueResponseOutput)
}

type PolicyDefinitionGroup struct {
	AdditionalMetadataId *string `pulumi:"additionalMetadataId"`
	Category             *string `pulumi:"category"`
	Description          *string `pulumi:"description"`
	DisplayName          *string `pulumi:"displayName"`
	Name                 string  `pulumi:"name"`
}





type PolicyDefinitionGroupInput interface {
	pulumi.Input

	ToPolicyDefinitionGroupOutput() PolicyDefinitionGroupOutput
	ToPolicyDefinitionGroupOutputWithContext(context.Context) PolicyDefinitionGroupOutput
}

type PolicyDefinitionGroupArgs struct {
	AdditionalMetadataId pulumi.StringPtrInput `pulumi:"additionalMetadataId"`
	Category             pulumi.StringPtrInput `pulumi:"category"`
	Description          pulumi.StringPtrInput `pulumi:"description"`
	DisplayName          pulumi.StringPtrInput `pulumi:"displayName"`
	Name                 pulumi.StringInput    `pulumi:"name"`
}

func (PolicyDefinitionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionGroup)(nil)).Elem()
}

func (i PolicyDefinitionGroupArgs) ToPolicyDefinitionGroupOutput() PolicyDefinitionGroupOutput {
	return i.ToPolicyDefinitionGroupOutputWithContext(context.Background())
}

func (i PolicyDefinitionGroupArgs) ToPolicyDefinitionGroupOutputWithContext(ctx context.Context) PolicyDefinitionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionGroupOutput)
}





type PolicyDefinitionGroupArrayInput interface {
	pulumi.Input

	ToPolicyDefinitionGroupArrayOutput() PolicyDefinitionGroupArrayOutput
	ToPolicyDefinitionGroupArrayOutputWithContext(context.Context) PolicyDefinitionGroupArrayOutput
}

type PolicyDefinitionGroupArray []PolicyDefinitionGroupInput

func (PolicyDefinitionGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionGroup)(nil)).Elem()
}

func (i PolicyDefinitionGroupArray) ToPolicyDefinitionGroupArrayOutput() PolicyDefinitionGroupArrayOutput {
	return i.ToPolicyDefinitionGroupArrayOutputWithContext(context.Background())
}

func (i PolicyDefinitionGroupArray) ToPolicyDefinitionGroupArrayOutputWithContext(ctx context.Context) PolicyDefinitionGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionGroupArrayOutput)
}

type PolicyDefinitionGroupOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionGroup)(nil)).Elem()
}

func (o PolicyDefinitionGroupOutput) ToPolicyDefinitionGroupOutput() PolicyDefinitionGroupOutput {
	return o
}

func (o PolicyDefinitionGroupOutput) ToPolicyDefinitionGroupOutputWithContext(ctx context.Context) PolicyDefinitionGroupOutput {
	return o
}

func (o PolicyDefinitionGroupOutput) AdditionalMetadataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroup) *string { return v.AdditionalMetadataId }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroup) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroup) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroup) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyDefinitionGroup) string { return v.Name }).(pulumi.StringOutput)
}

type PolicyDefinitionGroupArrayOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionGroup)(nil)).Elem()
}

func (o PolicyDefinitionGroupArrayOutput) ToPolicyDefinitionGroupArrayOutput() PolicyDefinitionGroupArrayOutput {
	return o
}

func (o PolicyDefinitionGroupArrayOutput) ToPolicyDefinitionGroupArrayOutputWithContext(ctx context.Context) PolicyDefinitionGroupArrayOutput {
	return o
}

func (o PolicyDefinitionGroupArrayOutput) Index(i pulumi.IntInput) PolicyDefinitionGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyDefinitionGroup {
		return vs[0].([]PolicyDefinitionGroup)[vs[1].(int)]
	}).(PolicyDefinitionGroupOutput)
}

type PolicyDefinitionGroupResponse struct {
	AdditionalMetadataId *string `pulumi:"additionalMetadataId"`
	Category             *string `pulumi:"category"`
	Description          *string `pulumi:"description"`
	DisplayName          *string `pulumi:"displayName"`
	Name                 string  `pulumi:"name"`
}

type PolicyDefinitionGroupResponseOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionGroupResponse)(nil)).Elem()
}

func (o PolicyDefinitionGroupResponseOutput) ToPolicyDefinitionGroupResponseOutput() PolicyDefinitionGroupResponseOutput {
	return o
}

func (o PolicyDefinitionGroupResponseOutput) ToPolicyDefinitionGroupResponseOutputWithContext(ctx context.Context) PolicyDefinitionGroupResponseOutput {
	return o
}

func (o PolicyDefinitionGroupResponseOutput) AdditionalMetadataId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroupResponse) *string { return v.AdditionalMetadataId }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroupResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroupResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionGroupResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyDefinitionGroupResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyDefinitionGroupResponse) string { return v.Name }).(pulumi.StringOutput)
}

type PolicyDefinitionGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionGroupResponse)(nil)).Elem()
}

func (o PolicyDefinitionGroupResponseArrayOutput) ToPolicyDefinitionGroupResponseArrayOutput() PolicyDefinitionGroupResponseArrayOutput {
	return o
}

func (o PolicyDefinitionGroupResponseArrayOutput) ToPolicyDefinitionGroupResponseArrayOutputWithContext(ctx context.Context) PolicyDefinitionGroupResponseArrayOutput {
	return o
}

func (o PolicyDefinitionGroupResponseArrayOutput) Index(i pulumi.IntInput) PolicyDefinitionGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyDefinitionGroupResponse {
		return vs[0].([]PolicyDefinitionGroupResponse)[vs[1].(int)]
	}).(PolicyDefinitionGroupResponseOutput)
}

type PolicyDefinitionReference struct {
	GroupNames                  []string                        `pulumi:"groupNames"`
	Parameters                  map[string]ParameterValuesValue `pulumi:"parameters"`
	PolicyDefinitionId          string                          `pulumi:"policyDefinitionId"`
	PolicyDefinitionReferenceId *string                         `pulumi:"policyDefinitionReferenceId"`
}





type PolicyDefinitionReferenceInput interface {
	pulumi.Input

	ToPolicyDefinitionReferenceOutput() PolicyDefinitionReferenceOutput
	ToPolicyDefinitionReferenceOutputWithContext(context.Context) PolicyDefinitionReferenceOutput
}

type PolicyDefinitionReferenceArgs struct {
	GroupNames                  pulumi.StringArrayInput      `pulumi:"groupNames"`
	Parameters                  ParameterValuesValueMapInput `pulumi:"parameters"`
	PolicyDefinitionId          pulumi.StringInput           `pulumi:"policyDefinitionId"`
	PolicyDefinitionReferenceId pulumi.StringPtrInput        `pulumi:"policyDefinitionReferenceId"`
}

func (PolicyDefinitionReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionReference)(nil)).Elem()
}

func (i PolicyDefinitionReferenceArgs) ToPolicyDefinitionReferenceOutput() PolicyDefinitionReferenceOutput {
	return i.ToPolicyDefinitionReferenceOutputWithContext(context.Background())
}

func (i PolicyDefinitionReferenceArgs) ToPolicyDefinitionReferenceOutputWithContext(ctx context.Context) PolicyDefinitionReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionReferenceOutput)
}





type PolicyDefinitionReferenceArrayInput interface {
	pulumi.Input

	ToPolicyDefinitionReferenceArrayOutput() PolicyDefinitionReferenceArrayOutput
	ToPolicyDefinitionReferenceArrayOutputWithContext(context.Context) PolicyDefinitionReferenceArrayOutput
}

type PolicyDefinitionReferenceArray []PolicyDefinitionReferenceInput

func (PolicyDefinitionReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionReference)(nil)).Elem()
}

func (i PolicyDefinitionReferenceArray) ToPolicyDefinitionReferenceArrayOutput() PolicyDefinitionReferenceArrayOutput {
	return i.ToPolicyDefinitionReferenceArrayOutputWithContext(context.Background())
}

func (i PolicyDefinitionReferenceArray) ToPolicyDefinitionReferenceArrayOutputWithContext(ctx context.Context) PolicyDefinitionReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionReferenceArrayOutput)
}

type PolicyDefinitionReferenceOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionReference)(nil)).Elem()
}

func (o PolicyDefinitionReferenceOutput) ToPolicyDefinitionReferenceOutput() PolicyDefinitionReferenceOutput {
	return o
}

func (o PolicyDefinitionReferenceOutput) ToPolicyDefinitionReferenceOutputWithContext(ctx context.Context) PolicyDefinitionReferenceOutput {
	return o
}

func (o PolicyDefinitionReferenceOutput) GroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PolicyDefinitionReference) []string { return v.GroupNames }).(pulumi.StringArrayOutput)
}

func (o PolicyDefinitionReferenceOutput) Parameters() ParameterValuesValueMapOutput {
	return o.ApplyT(func(v PolicyDefinitionReference) map[string]ParameterValuesValue { return v.Parameters }).(ParameterValuesValueMapOutput)
}

func (o PolicyDefinitionReferenceOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyDefinitionReference) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o PolicyDefinitionReferenceOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionReference) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

type PolicyDefinitionReferenceArrayOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionReference)(nil)).Elem()
}

func (o PolicyDefinitionReferenceArrayOutput) ToPolicyDefinitionReferenceArrayOutput() PolicyDefinitionReferenceArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceArrayOutput) ToPolicyDefinitionReferenceArrayOutputWithContext(ctx context.Context) PolicyDefinitionReferenceArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceArrayOutput) Index(i pulumi.IntInput) PolicyDefinitionReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyDefinitionReference {
		return vs[0].([]PolicyDefinitionReference)[vs[1].(int)]
	}).(PolicyDefinitionReferenceOutput)
}

type PolicyDefinitionReferenceResponse struct {
	GroupNames                  []string                                `pulumi:"groupNames"`
	Parameters                  map[string]ParameterValuesValueResponse `pulumi:"parameters"`
	PolicyDefinitionId          string                                  `pulumi:"policyDefinitionId"`
	PolicyDefinitionReferenceId *string                                 `pulumi:"policyDefinitionReferenceId"`
}

type PolicyDefinitionReferenceResponseOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionReferenceResponse)(nil)).Elem()
}

func (o PolicyDefinitionReferenceResponseOutput) ToPolicyDefinitionReferenceResponseOutput() PolicyDefinitionReferenceResponseOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseOutput) ToPolicyDefinitionReferenceResponseOutputWithContext(ctx context.Context) PolicyDefinitionReferenceResponseOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseOutput) GroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PolicyDefinitionReferenceResponse) []string { return v.GroupNames }).(pulumi.StringArrayOutput)
}

func (o PolicyDefinitionReferenceResponseOutput) Parameters() ParameterValuesValueResponseMapOutput {
	return o.ApplyT(func(v PolicyDefinitionReferenceResponse) map[string]ParameterValuesValueResponse { return v.Parameters }).(ParameterValuesValueResponseMapOutput)
}

func (o PolicyDefinitionReferenceResponseOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyDefinitionReferenceResponse) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o PolicyDefinitionReferenceResponseOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionReferenceResponse) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

type PolicyDefinitionReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionReferenceResponse)(nil)).Elem()
}

func (o PolicyDefinitionReferenceResponseArrayOutput) ToPolicyDefinitionReferenceResponseArrayOutput() PolicyDefinitionReferenceResponseArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseArrayOutput) ToPolicyDefinitionReferenceResponseArrayOutputWithContext(ctx context.Context) PolicyDefinitionReferenceResponseArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseArrayOutput) Index(i pulumi.IntInput) PolicyDefinitionReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyDefinitionReferenceResponse {
		return vs[0].([]PolicyDefinitionReferenceResponse)[vs[1].(int)]
	}).(PolicyDefinitionReferenceResponseOutput)
}

type PolicySku struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type PolicySkuInput interface {
	pulumi.Input

	ToPolicySkuOutput() PolicySkuOutput
	ToPolicySkuOutputWithContext(context.Context) PolicySkuOutput
}

type PolicySkuArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (PolicySkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySku)(nil)).Elem()
}

func (i PolicySkuArgs) ToPolicySkuOutput() PolicySkuOutput {
	return i.ToPolicySkuOutputWithContext(context.Background())
}

func (i PolicySkuArgs) ToPolicySkuOutputWithContext(ctx context.Context) PolicySkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySkuOutput)
}

func (i PolicySkuArgs) ToPolicySkuPtrOutput() PolicySkuPtrOutput {
	return i.ToPolicySkuPtrOutputWithContext(context.Background())
}

func (i PolicySkuArgs) ToPolicySkuPtrOutputWithContext(ctx context.Context) PolicySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySkuOutput).ToPolicySkuPtrOutputWithContext(ctx)
}









type PolicySkuPtrInput interface {
	pulumi.Input

	ToPolicySkuPtrOutput() PolicySkuPtrOutput
	ToPolicySkuPtrOutputWithContext(context.Context) PolicySkuPtrOutput
}

type policySkuPtrType PolicySkuArgs

func PolicySkuPtr(v *PolicySkuArgs) PolicySkuPtrInput {
	return (*policySkuPtrType)(v)
}

func (*policySkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySku)(nil)).Elem()
}

func (i *policySkuPtrType) ToPolicySkuPtrOutput() PolicySkuPtrOutput {
	return i.ToPolicySkuPtrOutputWithContext(context.Background())
}

func (i *policySkuPtrType) ToPolicySkuPtrOutputWithContext(ctx context.Context) PolicySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySkuPtrOutput)
}

type PolicySkuOutput struct{ *pulumi.OutputState }

func (PolicySkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySku)(nil)).Elem()
}

func (o PolicySkuOutput) ToPolicySkuOutput() PolicySkuOutput {
	return o
}

func (o PolicySkuOutput) ToPolicySkuOutputWithContext(ctx context.Context) PolicySkuOutput {
	return o
}

func (o PolicySkuOutput) ToPolicySkuPtrOutput() PolicySkuPtrOutput {
	return o.ToPolicySkuPtrOutputWithContext(context.Background())
}

func (o PolicySkuOutput) ToPolicySkuPtrOutputWithContext(ctx context.Context) PolicySkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicySku) *PolicySku {
		return &v
	}).(PolicySkuPtrOutput)
}

func (o PolicySkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PolicySku) string { return v.Name }).(pulumi.StringOutput)
}

func (o PolicySkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type PolicySkuPtrOutput struct{ *pulumi.OutputState }

func (PolicySkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySku)(nil)).Elem()
}

func (o PolicySkuPtrOutput) ToPolicySkuPtrOutput() PolicySkuPtrOutput {
	return o
}

func (o PolicySkuPtrOutput) ToPolicySkuPtrOutputWithContext(ctx context.Context) PolicySkuPtrOutput {
	return o
}

func (o PolicySkuPtrOutput) Elem() PolicySkuOutput {
	return o.ApplyT(func(v *PolicySku) PolicySku {
		if v != nil {
			return *v
		}
		var ret PolicySku
		return ret
	}).(PolicySkuOutput)
}

func (o PolicySkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PolicySkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type PolicySkuResponse struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}

type PolicySkuResponseOutput struct{ *pulumi.OutputState }

func (PolicySkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySkuResponse)(nil)).Elem()
}

func (o PolicySkuResponseOutput) ToPolicySkuResponseOutput() PolicySkuResponseOutput {
	return o
}

func (o PolicySkuResponseOutput) ToPolicySkuResponseOutputWithContext(ctx context.Context) PolicySkuResponseOutput {
	return o
}

func (o PolicySkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PolicySkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PolicySkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type PolicySkuResponsePtrOutput struct{ *pulumi.OutputState }

func (PolicySkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySkuResponse)(nil)).Elem()
}

func (o PolicySkuResponsePtrOutput) ToPolicySkuResponsePtrOutput() PolicySkuResponsePtrOutput {
	return o
}

func (o PolicySkuResponsePtrOutput) ToPolicySkuResponsePtrOutputWithContext(ctx context.Context) PolicySkuResponsePtrOutput {
	return o
}

func (o PolicySkuResponsePtrOutput) Elem() PolicySkuResponseOutput {
	return o.ApplyT(func(v *PolicySkuResponse) PolicySkuResponse {
		if v != nil {
			return *v
		}
		var ret PolicySkuResponse
		return ret
	}).(PolicySkuResponseOutput)
}

func (o PolicySkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PolicySkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueMapOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueMetadataOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueMetadataPtrOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueResponseOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueResponseMapOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueResponseMetadataOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionsValueResponseMetadataPtrOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueMapOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueResponseOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueResponseMapOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionGroupOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionGroupArrayOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionGroupResponseOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceArrayOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceResponseOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(PolicySkuOutput{})
	pulumi.RegisterOutputType(PolicySkuPtrOutput{})
	pulumi.RegisterOutputType(PolicySkuResponseOutput{})
	pulumi.RegisterOutputType(PolicySkuResponsePtrOutput{})
}
