


package v20220601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Identity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
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

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
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

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                            `pulumi:"principalId"`
	TenantId               string                                            `pulumi:"tenantId"`
	Type                   *string                                           `pulumi:"type"`
	UserAssignedIdentities map[string]IdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
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

func (o IdentityResponseOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
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

func (o IdentityResponsePtrOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
}

type IdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type IdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutput() IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o IdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type IdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutput() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) IdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]IdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(IdentityResponseUserAssignedIdentitiesOutput)
}

type NonComplianceMessage struct {
	Message                     string  `pulumi:"message"`
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
}





type NonComplianceMessageInput interface {
	pulumi.Input

	ToNonComplianceMessageOutput() NonComplianceMessageOutput
	ToNonComplianceMessageOutputWithContext(context.Context) NonComplianceMessageOutput
}

type NonComplianceMessageArgs struct {
	Message                     pulumi.StringInput    `pulumi:"message"`
	PolicyDefinitionReferenceId pulumi.StringPtrInput `pulumi:"policyDefinitionReferenceId"`
}

func (NonComplianceMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NonComplianceMessage)(nil)).Elem()
}

func (i NonComplianceMessageArgs) ToNonComplianceMessageOutput() NonComplianceMessageOutput {
	return i.ToNonComplianceMessageOutputWithContext(context.Background())
}

func (i NonComplianceMessageArgs) ToNonComplianceMessageOutputWithContext(ctx context.Context) NonComplianceMessageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonComplianceMessageOutput)
}





type NonComplianceMessageArrayInput interface {
	pulumi.Input

	ToNonComplianceMessageArrayOutput() NonComplianceMessageArrayOutput
	ToNonComplianceMessageArrayOutputWithContext(context.Context) NonComplianceMessageArrayOutput
}

type NonComplianceMessageArray []NonComplianceMessageInput

func (NonComplianceMessageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NonComplianceMessage)(nil)).Elem()
}

func (i NonComplianceMessageArray) ToNonComplianceMessageArrayOutput() NonComplianceMessageArrayOutput {
	return i.ToNonComplianceMessageArrayOutputWithContext(context.Background())
}

func (i NonComplianceMessageArray) ToNonComplianceMessageArrayOutputWithContext(ctx context.Context) NonComplianceMessageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonComplianceMessageArrayOutput)
}

type NonComplianceMessageOutput struct{ *pulumi.OutputState }

func (NonComplianceMessageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NonComplianceMessage)(nil)).Elem()
}

func (o NonComplianceMessageOutput) ToNonComplianceMessageOutput() NonComplianceMessageOutput {
	return o
}

func (o NonComplianceMessageOutput) ToNonComplianceMessageOutputWithContext(ctx context.Context) NonComplianceMessageOutput {
	return o
}

func (o NonComplianceMessageOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v NonComplianceMessage) string { return v.Message }).(pulumi.StringOutput)
}

func (o NonComplianceMessageOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NonComplianceMessage) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

type NonComplianceMessageArrayOutput struct{ *pulumi.OutputState }

func (NonComplianceMessageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NonComplianceMessage)(nil)).Elem()
}

func (o NonComplianceMessageArrayOutput) ToNonComplianceMessageArrayOutput() NonComplianceMessageArrayOutput {
	return o
}

func (o NonComplianceMessageArrayOutput) ToNonComplianceMessageArrayOutputWithContext(ctx context.Context) NonComplianceMessageArrayOutput {
	return o
}

func (o NonComplianceMessageArrayOutput) Index(i pulumi.IntInput) NonComplianceMessageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NonComplianceMessage {
		return vs[0].([]NonComplianceMessage)[vs[1].(int)]
	}).(NonComplianceMessageOutput)
}

type NonComplianceMessageResponse struct {
	Message                     string  `pulumi:"message"`
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
}

type NonComplianceMessageResponseOutput struct{ *pulumi.OutputState }

func (NonComplianceMessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NonComplianceMessageResponse)(nil)).Elem()
}

func (o NonComplianceMessageResponseOutput) ToNonComplianceMessageResponseOutput() NonComplianceMessageResponseOutput {
	return o
}

func (o NonComplianceMessageResponseOutput) ToNonComplianceMessageResponseOutputWithContext(ctx context.Context) NonComplianceMessageResponseOutput {
	return o
}

func (o NonComplianceMessageResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v NonComplianceMessageResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o NonComplianceMessageResponseOutput) PolicyDefinitionReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NonComplianceMessageResponse) *string { return v.PolicyDefinitionReferenceId }).(pulumi.StringPtrOutput)
}

type NonComplianceMessageResponseArrayOutput struct{ *pulumi.OutputState }

func (NonComplianceMessageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NonComplianceMessageResponse)(nil)).Elem()
}

func (o NonComplianceMessageResponseArrayOutput) ToNonComplianceMessageResponseArrayOutput() NonComplianceMessageResponseArrayOutput {
	return o
}

func (o NonComplianceMessageResponseArrayOutput) ToNonComplianceMessageResponseArrayOutputWithContext(ctx context.Context) NonComplianceMessageResponseArrayOutput {
	return o
}

func (o NonComplianceMessageResponseArrayOutput) Index(i pulumi.IntInput) NonComplianceMessageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NonComplianceMessageResponse {
		return vs[0].([]NonComplianceMessageResponse)[vs[1].(int)]
	}).(NonComplianceMessageResponseOutput)
}

type Override struct {
	Kind      *string    `pulumi:"kind"`
	Selectors []Selector `pulumi:"selectors"`
	Value     *string    `pulumi:"value"`
}





type OverrideInput interface {
	pulumi.Input

	ToOverrideOutput() OverrideOutput
	ToOverrideOutputWithContext(context.Context) OverrideOutput
}

type OverrideArgs struct {
	Kind      pulumi.StringPtrInput `pulumi:"kind"`
	Selectors SelectorArrayInput    `pulumi:"selectors"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (OverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Override)(nil)).Elem()
}

func (i OverrideArgs) ToOverrideOutput() OverrideOutput {
	return i.ToOverrideOutputWithContext(context.Background())
}

func (i OverrideArgs) ToOverrideOutputWithContext(ctx context.Context) OverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverrideOutput)
}





type OverrideArrayInput interface {
	pulumi.Input

	ToOverrideArrayOutput() OverrideArrayOutput
	ToOverrideArrayOutputWithContext(context.Context) OverrideArrayOutput
}

type OverrideArray []OverrideInput

func (OverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Override)(nil)).Elem()
}

func (i OverrideArray) ToOverrideArrayOutput() OverrideArrayOutput {
	return i.ToOverrideArrayOutputWithContext(context.Background())
}

func (i OverrideArray) ToOverrideArrayOutputWithContext(ctx context.Context) OverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverrideArrayOutput)
}

type OverrideOutput struct{ *pulumi.OutputState }

func (OverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Override)(nil)).Elem()
}

func (o OverrideOutput) ToOverrideOutput() OverrideOutput {
	return o
}

func (o OverrideOutput) ToOverrideOutputWithContext(ctx context.Context) OverrideOutput {
	return o
}

func (o OverrideOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Override) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o OverrideOutput) Selectors() SelectorArrayOutput {
	return o.ApplyT(func(v Override) []Selector { return v.Selectors }).(SelectorArrayOutput)
}

func (o OverrideOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Override) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type OverrideArrayOutput struct{ *pulumi.OutputState }

func (OverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Override)(nil)).Elem()
}

func (o OverrideArrayOutput) ToOverrideArrayOutput() OverrideArrayOutput {
	return o
}

func (o OverrideArrayOutput) ToOverrideArrayOutputWithContext(ctx context.Context) OverrideArrayOutput {
	return o
}

func (o OverrideArrayOutput) Index(i pulumi.IntInput) OverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Override {
		return vs[0].([]Override)[vs[1].(int)]
	}).(OverrideOutput)
}

type OverrideResponse struct {
	Kind      *string            `pulumi:"kind"`
	Selectors []SelectorResponse `pulumi:"selectors"`
	Value     *string            `pulumi:"value"`
}

type OverrideResponseOutput struct{ *pulumi.OutputState }

func (OverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OverrideResponse)(nil)).Elem()
}

func (o OverrideResponseOutput) ToOverrideResponseOutput() OverrideResponseOutput {
	return o
}

func (o OverrideResponseOutput) ToOverrideResponseOutputWithContext(ctx context.Context) OverrideResponseOutput {
	return o
}

func (o OverrideResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverrideResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o OverrideResponseOutput) Selectors() SelectorResponseArrayOutput {
	return o.ApplyT(func(v OverrideResponse) []SelectorResponse { return v.Selectors }).(SelectorResponseArrayOutput)
}

func (o OverrideResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OverrideResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type OverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (OverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OverrideResponse)(nil)).Elem()
}

func (o OverrideResponseArrayOutput) ToOverrideResponseArrayOutput() OverrideResponseArrayOutput {
	return o
}

func (o OverrideResponseArrayOutput) ToOverrideResponseArrayOutputWithContext(ctx context.Context) OverrideResponseArrayOutput {
	return o
}

func (o OverrideResponseArrayOutput) Index(i pulumi.IntInput) OverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OverrideResponse {
		return vs[0].([]OverrideResponse)[vs[1].(int)]
	}).(OverrideResponseOutput)
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

type ResourceSelector struct {
	Name      *string    `pulumi:"name"`
	Selectors []Selector `pulumi:"selectors"`
}





type ResourceSelectorInput interface {
	pulumi.Input

	ToResourceSelectorOutput() ResourceSelectorOutput
	ToResourceSelectorOutputWithContext(context.Context) ResourceSelectorOutput
}

type ResourceSelectorArgs struct {
	Name      pulumi.StringPtrInput `pulumi:"name"`
	Selectors SelectorArrayInput    `pulumi:"selectors"`
}

func (ResourceSelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSelector)(nil)).Elem()
}

func (i ResourceSelectorArgs) ToResourceSelectorOutput() ResourceSelectorOutput {
	return i.ToResourceSelectorOutputWithContext(context.Background())
}

func (i ResourceSelectorArgs) ToResourceSelectorOutputWithContext(ctx context.Context) ResourceSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSelectorOutput)
}





type ResourceSelectorArrayInput interface {
	pulumi.Input

	ToResourceSelectorArrayOutput() ResourceSelectorArrayOutput
	ToResourceSelectorArrayOutputWithContext(context.Context) ResourceSelectorArrayOutput
}

type ResourceSelectorArray []ResourceSelectorInput

func (ResourceSelectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceSelector)(nil)).Elem()
}

func (i ResourceSelectorArray) ToResourceSelectorArrayOutput() ResourceSelectorArrayOutput {
	return i.ToResourceSelectorArrayOutputWithContext(context.Background())
}

func (i ResourceSelectorArray) ToResourceSelectorArrayOutputWithContext(ctx context.Context) ResourceSelectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSelectorArrayOutput)
}

type ResourceSelectorOutput struct{ *pulumi.OutputState }

func (ResourceSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSelector)(nil)).Elem()
}

func (o ResourceSelectorOutput) ToResourceSelectorOutput() ResourceSelectorOutput {
	return o
}

func (o ResourceSelectorOutput) ToResourceSelectorOutputWithContext(ctx context.Context) ResourceSelectorOutput {
	return o
}

func (o ResourceSelectorOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSelector) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceSelectorOutput) Selectors() SelectorArrayOutput {
	return o.ApplyT(func(v ResourceSelector) []Selector { return v.Selectors }).(SelectorArrayOutput)
}

type ResourceSelectorArrayOutput struct{ *pulumi.OutputState }

func (ResourceSelectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceSelector)(nil)).Elem()
}

func (o ResourceSelectorArrayOutput) ToResourceSelectorArrayOutput() ResourceSelectorArrayOutput {
	return o
}

func (o ResourceSelectorArrayOutput) ToResourceSelectorArrayOutputWithContext(ctx context.Context) ResourceSelectorArrayOutput {
	return o
}

func (o ResourceSelectorArrayOutput) Index(i pulumi.IntInput) ResourceSelectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceSelector {
		return vs[0].([]ResourceSelector)[vs[1].(int)]
	}).(ResourceSelectorOutput)
}

type ResourceSelectorResponse struct {
	Name      *string            `pulumi:"name"`
	Selectors []SelectorResponse `pulumi:"selectors"`
}

type ResourceSelectorResponseOutput struct{ *pulumi.OutputState }

func (ResourceSelectorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSelectorResponse)(nil)).Elem()
}

func (o ResourceSelectorResponseOutput) ToResourceSelectorResponseOutput() ResourceSelectorResponseOutput {
	return o
}

func (o ResourceSelectorResponseOutput) ToResourceSelectorResponseOutputWithContext(ctx context.Context) ResourceSelectorResponseOutput {
	return o
}

func (o ResourceSelectorResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSelectorResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceSelectorResponseOutput) Selectors() SelectorResponseArrayOutput {
	return o.ApplyT(func(v ResourceSelectorResponse) []SelectorResponse { return v.Selectors }).(SelectorResponseArrayOutput)
}

type ResourceSelectorResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceSelectorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceSelectorResponse)(nil)).Elem()
}

func (o ResourceSelectorResponseArrayOutput) ToResourceSelectorResponseArrayOutput() ResourceSelectorResponseArrayOutput {
	return o
}

func (o ResourceSelectorResponseArrayOutput) ToResourceSelectorResponseArrayOutputWithContext(ctx context.Context) ResourceSelectorResponseArrayOutput {
	return o
}

func (o ResourceSelectorResponseArrayOutput) Index(i pulumi.IntInput) ResourceSelectorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceSelectorResponse {
		return vs[0].([]ResourceSelectorResponse)[vs[1].(int)]
	}).(ResourceSelectorResponseOutput)
}

type Selector struct {
	In    []string `pulumi:"in"`
	Kind  *string  `pulumi:"kind"`
	NotIn []string `pulumi:"notIn"`
}





type SelectorInput interface {
	pulumi.Input

	ToSelectorOutput() SelectorOutput
	ToSelectorOutputWithContext(context.Context) SelectorOutput
}

type SelectorArgs struct {
	In    pulumi.StringArrayInput `pulumi:"in"`
	Kind  pulumi.StringPtrInput   `pulumi:"kind"`
	NotIn pulumi.StringArrayInput `pulumi:"notIn"`
}

func (SelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Selector)(nil)).Elem()
}

func (i SelectorArgs) ToSelectorOutput() SelectorOutput {
	return i.ToSelectorOutputWithContext(context.Background())
}

func (i SelectorArgs) ToSelectorOutputWithContext(ctx context.Context) SelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectorOutput)
}





type SelectorArrayInput interface {
	pulumi.Input

	ToSelectorArrayOutput() SelectorArrayOutput
	ToSelectorArrayOutputWithContext(context.Context) SelectorArrayOutput
}

type SelectorArray []SelectorInput

func (SelectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Selector)(nil)).Elem()
}

func (i SelectorArray) ToSelectorArrayOutput() SelectorArrayOutput {
	return i.ToSelectorArrayOutputWithContext(context.Background())
}

func (i SelectorArray) ToSelectorArrayOutputWithContext(ctx context.Context) SelectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectorArrayOutput)
}

type SelectorOutput struct{ *pulumi.OutputState }

func (SelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Selector)(nil)).Elem()
}

func (o SelectorOutput) ToSelectorOutput() SelectorOutput {
	return o
}

func (o SelectorOutput) ToSelectorOutputWithContext(ctx context.Context) SelectorOutput {
	return o
}

func (o SelectorOutput) In() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Selector) []string { return v.In }).(pulumi.StringArrayOutput)
}

func (o SelectorOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Selector) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SelectorOutput) NotIn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Selector) []string { return v.NotIn }).(pulumi.StringArrayOutput)
}

type SelectorArrayOutput struct{ *pulumi.OutputState }

func (SelectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Selector)(nil)).Elem()
}

func (o SelectorArrayOutput) ToSelectorArrayOutput() SelectorArrayOutput {
	return o
}

func (o SelectorArrayOutput) ToSelectorArrayOutputWithContext(ctx context.Context) SelectorArrayOutput {
	return o
}

func (o SelectorArrayOutput) Index(i pulumi.IntInput) SelectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Selector {
		return vs[0].([]Selector)[vs[1].(int)]
	}).(SelectorOutput)
}

type SelectorResponse struct {
	In    []string `pulumi:"in"`
	Kind  *string  `pulumi:"kind"`
	NotIn []string `pulumi:"notIn"`
}

type SelectorResponseOutput struct{ *pulumi.OutputState }

func (SelectorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectorResponse)(nil)).Elem()
}

func (o SelectorResponseOutput) ToSelectorResponseOutput() SelectorResponseOutput {
	return o
}

func (o SelectorResponseOutput) ToSelectorResponseOutputWithContext(ctx context.Context) SelectorResponseOutput {
	return o
}

func (o SelectorResponseOutput) In() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SelectorResponse) []string { return v.In }).(pulumi.StringArrayOutput)
}

func (o SelectorResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelectorResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SelectorResponseOutput) NotIn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SelectorResponse) []string { return v.NotIn }).(pulumi.StringArrayOutput)
}

type SelectorResponseArrayOutput struct{ *pulumi.OutputState }

func (SelectorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SelectorResponse)(nil)).Elem()
}

func (o SelectorResponseArrayOutput) ToSelectorResponseArrayOutput() SelectorResponseArrayOutput {
	return o
}

func (o SelectorResponseArrayOutput) ToSelectorResponseArrayOutputWithContext(ctx context.Context) SelectorResponseArrayOutput {
	return o
}

func (o SelectorResponseArrayOutput) Index(i pulumi.IntInput) SelectorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SelectorResponse {
		return vs[0].([]SelectorResponse)[vs[1].(int)]
	}).(SelectorResponseOutput)
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

func init() {
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(NonComplianceMessageOutput{})
	pulumi.RegisterOutputType(NonComplianceMessageArrayOutput{})
	pulumi.RegisterOutputType(NonComplianceMessageResponseOutput{})
	pulumi.RegisterOutputType(NonComplianceMessageResponseArrayOutput{})
	pulumi.RegisterOutputType(OverrideOutput{})
	pulumi.RegisterOutputType(OverrideArrayOutput{})
	pulumi.RegisterOutputType(OverrideResponseOutput{})
	pulumi.RegisterOutputType(OverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueMapOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueResponseOutput{})
	pulumi.RegisterOutputType(ParameterValuesValueResponseMapOutput{})
	pulumi.RegisterOutputType(ResourceSelectorOutput{})
	pulumi.RegisterOutputType(ResourceSelectorArrayOutput{})
	pulumi.RegisterOutputType(ResourceSelectorResponseOutput{})
	pulumi.RegisterOutputType(ResourceSelectorResponseArrayOutput{})
	pulumi.RegisterOutputType(SelectorOutput{})
	pulumi.RegisterOutputType(SelectorArrayOutput{})
	pulumi.RegisterOutputType(SelectorResponseOutput{})
	pulumi.RegisterOutputType(SelectorResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
