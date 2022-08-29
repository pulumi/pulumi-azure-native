


package blueprint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssignmentLockSettings struct {
	ExcludedActions    []string `pulumi:"excludedActions"`
	ExcludedPrincipals []string `pulumi:"excludedPrincipals"`
	Mode               *string  `pulumi:"mode"`
}





type AssignmentLockSettingsInput interface {
	pulumi.Input

	ToAssignmentLockSettingsOutput() AssignmentLockSettingsOutput
	ToAssignmentLockSettingsOutputWithContext(context.Context) AssignmentLockSettingsOutput
}

type AssignmentLockSettingsArgs struct {
	ExcludedActions    pulumi.StringArrayInput `pulumi:"excludedActions"`
	ExcludedPrincipals pulumi.StringArrayInput `pulumi:"excludedPrincipals"`
	Mode               pulumi.StringPtrInput   `pulumi:"mode"`
}

func (AssignmentLockSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentLockSettings)(nil)).Elem()
}

func (i AssignmentLockSettingsArgs) ToAssignmentLockSettingsOutput() AssignmentLockSettingsOutput {
	return i.ToAssignmentLockSettingsOutputWithContext(context.Background())
}

func (i AssignmentLockSettingsArgs) ToAssignmentLockSettingsOutputWithContext(ctx context.Context) AssignmentLockSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentLockSettingsOutput)
}

func (i AssignmentLockSettingsArgs) ToAssignmentLockSettingsPtrOutput() AssignmentLockSettingsPtrOutput {
	return i.ToAssignmentLockSettingsPtrOutputWithContext(context.Background())
}

func (i AssignmentLockSettingsArgs) ToAssignmentLockSettingsPtrOutputWithContext(ctx context.Context) AssignmentLockSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentLockSettingsOutput).ToAssignmentLockSettingsPtrOutputWithContext(ctx)
}









type AssignmentLockSettingsPtrInput interface {
	pulumi.Input

	ToAssignmentLockSettingsPtrOutput() AssignmentLockSettingsPtrOutput
	ToAssignmentLockSettingsPtrOutputWithContext(context.Context) AssignmentLockSettingsPtrOutput
}

type assignmentLockSettingsPtrType AssignmentLockSettingsArgs

func AssignmentLockSettingsPtr(v *AssignmentLockSettingsArgs) AssignmentLockSettingsPtrInput {
	return (*assignmentLockSettingsPtrType)(v)
}

func (*assignmentLockSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentLockSettings)(nil)).Elem()
}

func (i *assignmentLockSettingsPtrType) ToAssignmentLockSettingsPtrOutput() AssignmentLockSettingsPtrOutput {
	return i.ToAssignmentLockSettingsPtrOutputWithContext(context.Background())
}

func (i *assignmentLockSettingsPtrType) ToAssignmentLockSettingsPtrOutputWithContext(ctx context.Context) AssignmentLockSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentLockSettingsPtrOutput)
}

type AssignmentLockSettingsOutput struct{ *pulumi.OutputState }

func (AssignmentLockSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentLockSettings)(nil)).Elem()
}

func (o AssignmentLockSettingsOutput) ToAssignmentLockSettingsOutput() AssignmentLockSettingsOutput {
	return o
}

func (o AssignmentLockSettingsOutput) ToAssignmentLockSettingsOutputWithContext(ctx context.Context) AssignmentLockSettingsOutput {
	return o
}

func (o AssignmentLockSettingsOutput) ToAssignmentLockSettingsPtrOutput() AssignmentLockSettingsPtrOutput {
	return o.ToAssignmentLockSettingsPtrOutputWithContext(context.Background())
}

func (o AssignmentLockSettingsOutput) ToAssignmentLockSettingsPtrOutputWithContext(ctx context.Context) AssignmentLockSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignmentLockSettings) *AssignmentLockSettings {
		return &v
	}).(AssignmentLockSettingsPtrOutput)
}

func (o AssignmentLockSettingsOutput) ExcludedActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AssignmentLockSettings) []string { return v.ExcludedActions }).(pulumi.StringArrayOutput)
}

func (o AssignmentLockSettingsOutput) ExcludedPrincipals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AssignmentLockSettings) []string { return v.ExcludedPrincipals }).(pulumi.StringArrayOutput)
}

func (o AssignmentLockSettingsOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignmentLockSettings) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type AssignmentLockSettingsPtrOutput struct{ *pulumi.OutputState }

func (AssignmentLockSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentLockSettings)(nil)).Elem()
}

func (o AssignmentLockSettingsPtrOutput) ToAssignmentLockSettingsPtrOutput() AssignmentLockSettingsPtrOutput {
	return o
}

func (o AssignmentLockSettingsPtrOutput) ToAssignmentLockSettingsPtrOutputWithContext(ctx context.Context) AssignmentLockSettingsPtrOutput {
	return o
}

func (o AssignmentLockSettingsPtrOutput) Elem() AssignmentLockSettingsOutput {
	return o.ApplyT(func(v *AssignmentLockSettings) AssignmentLockSettings {
		if v != nil {
			return *v
		}
		var ret AssignmentLockSettings
		return ret
	}).(AssignmentLockSettingsOutput)
}

func (o AssignmentLockSettingsPtrOutput) ExcludedActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssignmentLockSettings) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedActions
	}).(pulumi.StringArrayOutput)
}

func (o AssignmentLockSettingsPtrOutput) ExcludedPrincipals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssignmentLockSettings) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedPrincipals
	}).(pulumi.StringArrayOutput)
}

func (o AssignmentLockSettingsPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentLockSettings) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type AssignmentLockSettingsResponse struct {
	ExcludedActions    []string `pulumi:"excludedActions"`
	ExcludedPrincipals []string `pulumi:"excludedPrincipals"`
	Mode               *string  `pulumi:"mode"`
}

type AssignmentLockSettingsResponseOutput struct{ *pulumi.OutputState }

func (AssignmentLockSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentLockSettingsResponse)(nil)).Elem()
}

func (o AssignmentLockSettingsResponseOutput) ToAssignmentLockSettingsResponseOutput() AssignmentLockSettingsResponseOutput {
	return o
}

func (o AssignmentLockSettingsResponseOutput) ToAssignmentLockSettingsResponseOutputWithContext(ctx context.Context) AssignmentLockSettingsResponseOutput {
	return o
}

func (o AssignmentLockSettingsResponseOutput) ExcludedActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AssignmentLockSettingsResponse) []string { return v.ExcludedActions }).(pulumi.StringArrayOutput)
}

func (o AssignmentLockSettingsResponseOutput) ExcludedPrincipals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AssignmentLockSettingsResponse) []string { return v.ExcludedPrincipals }).(pulumi.StringArrayOutput)
}

func (o AssignmentLockSettingsResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignmentLockSettingsResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type AssignmentLockSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AssignmentLockSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentLockSettingsResponse)(nil)).Elem()
}

func (o AssignmentLockSettingsResponsePtrOutput) ToAssignmentLockSettingsResponsePtrOutput() AssignmentLockSettingsResponsePtrOutput {
	return o
}

func (o AssignmentLockSettingsResponsePtrOutput) ToAssignmentLockSettingsResponsePtrOutputWithContext(ctx context.Context) AssignmentLockSettingsResponsePtrOutput {
	return o
}

func (o AssignmentLockSettingsResponsePtrOutput) Elem() AssignmentLockSettingsResponseOutput {
	return o.ApplyT(func(v *AssignmentLockSettingsResponse) AssignmentLockSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AssignmentLockSettingsResponse
		return ret
	}).(AssignmentLockSettingsResponseOutput)
}

func (o AssignmentLockSettingsResponsePtrOutput) ExcludedActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssignmentLockSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedActions
	}).(pulumi.StringArrayOutput)
}

func (o AssignmentLockSettingsResponsePtrOutput) ExcludedPrincipals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssignmentLockSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedPrincipals
	}).(pulumi.StringArrayOutput)
}

func (o AssignmentLockSettingsResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentLockSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type AssignmentStatusResponse struct {
	LastModified     string   `pulumi:"lastModified"`
	ManagedResources []string `pulumi:"managedResources"`
	TimeCreated      string   `pulumi:"timeCreated"`
}

type AssignmentStatusResponseOutput struct{ *pulumi.OutputState }

func (AssignmentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentStatusResponse)(nil)).Elem()
}

func (o AssignmentStatusResponseOutput) ToAssignmentStatusResponseOutput() AssignmentStatusResponseOutput {
	return o
}

func (o AssignmentStatusResponseOutput) ToAssignmentStatusResponseOutputWithContext(ctx context.Context) AssignmentStatusResponseOutput {
	return o
}

func (o AssignmentStatusResponseOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentStatusResponse) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o AssignmentStatusResponseOutput) ManagedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AssignmentStatusResponse) []string { return v.ManagedResources }).(pulumi.StringArrayOutput)
}

func (o AssignmentStatusResponseOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentStatusResponse) string { return v.TimeCreated }).(pulumi.StringOutput)
}

type BlueprintStatusResponse struct {
	LastModified string `pulumi:"lastModified"`
	TimeCreated  string `pulumi:"timeCreated"`
}

type BlueprintStatusResponseOutput struct{ *pulumi.OutputState }

func (BlueprintStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlueprintStatusResponse)(nil)).Elem()
}

func (o BlueprintStatusResponseOutput) ToBlueprintStatusResponseOutput() BlueprintStatusResponseOutput {
	return o
}

func (o BlueprintStatusResponseOutput) ToBlueprintStatusResponseOutputWithContext(ctx context.Context) BlueprintStatusResponseOutput {
	return o
}

func (o BlueprintStatusResponseOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v BlueprintStatusResponse) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o BlueprintStatusResponseOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v BlueprintStatusResponse) string { return v.TimeCreated }).(pulumi.StringOutput)
}

type KeyVaultReference struct {
	Id string `pulumi:"id"`
}





type KeyVaultReferenceInput interface {
	pulumi.Input

	ToKeyVaultReferenceOutput() KeyVaultReferenceOutput
	ToKeyVaultReferenceOutputWithContext(context.Context) KeyVaultReferenceOutput
}

type KeyVaultReferenceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (KeyVaultReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReference)(nil)).Elem()
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferenceOutput() KeyVaultReferenceOutput {
	return i.ToKeyVaultReferenceOutputWithContext(context.Background())
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferenceOutputWithContext(ctx context.Context) KeyVaultReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceOutput)
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return i.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceOutput).ToKeyVaultReferencePtrOutputWithContext(ctx)
}









type KeyVaultReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput
	ToKeyVaultReferencePtrOutputWithContext(context.Context) KeyVaultReferencePtrOutput
}

type keyVaultReferencePtrType KeyVaultReferenceArgs

func KeyVaultReferencePtr(v *KeyVaultReferenceArgs) KeyVaultReferencePtrInput {
	return (*keyVaultReferencePtrType)(v)
}

func (*keyVaultReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultReference)(nil)).Elem()
}

func (i *keyVaultReferencePtrType) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return i.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultReferencePtrType) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferencePtrOutput)
}

type KeyVaultReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReference)(nil)).Elem()
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferenceOutput() KeyVaultReferenceOutput {
	return o
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferenceOutputWithContext(ctx context.Context) KeyVaultReferenceOutput {
	return o
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return o.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultReference) *KeyVaultReference {
		return &v
	}).(KeyVaultReferencePtrOutput)
}

func (o KeyVaultReferenceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReference) string { return v.Id }).(pulumi.StringOutput)
}

type KeyVaultReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultReference)(nil)).Elem()
}

func (o KeyVaultReferencePtrOutput) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return o
}

func (o KeyVaultReferencePtrOutput) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return o
}

func (o KeyVaultReferencePtrOutput) Elem() KeyVaultReferenceOutput {
	return o.ApplyT(func(v *KeyVaultReference) KeyVaultReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultReference
		return ret
	}).(KeyVaultReferenceOutput)
}

func (o KeyVaultReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultReference) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type KeyVaultReferenceResponse struct {
	Id string `pulumi:"id"`
}

type KeyVaultReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReferenceResponse)(nil)).Elem()
}

func (o KeyVaultReferenceResponseOutput) ToKeyVaultReferenceResponseOutput() KeyVaultReferenceResponseOutput {
	return o
}

func (o KeyVaultReferenceResponseOutput) ToKeyVaultReferenceResponseOutputWithContext(ctx context.Context) KeyVaultReferenceResponseOutput {
	return o
}

func (o KeyVaultReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

type KeyVaultReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultReferenceResponse)(nil)).Elem()
}

func (o KeyVaultReferenceResponsePtrOutput) ToKeyVaultReferenceResponsePtrOutput() KeyVaultReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultReferenceResponsePtrOutput) ToKeyVaultReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultReferenceResponsePtrOutput) Elem() KeyVaultReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultReferenceResponse) KeyVaultReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultReferenceResponse
		return ret
	}).(KeyVaultReferenceResponseOutput)
}

func (o KeyVaultReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ManagedServiceIdentity struct {
	PrincipalId            *string                         `pulumi:"principalId"`
	TenantId               *string                         `pulumi:"tenantId"`
	Type                   string                          `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentity `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	PrincipalId            pulumi.StringPtrInput        `pulumi:"principalId"`
	TenantId               pulumi.StringPtrInput        `pulumi:"tenantId"`
	Type                   pulumi.StringInput           `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityMapInput `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() UserAssignedIdentityMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]UserAssignedIdentity { return v.UserAssignedIdentities }).(UserAssignedIdentityMapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            *string                                 `pulumi:"principalId"`
	TenantId               *string                                 `pulumi:"tenantId"`
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ParameterDefinition struct {
	AllowedValues []interface{} `pulumi:"allowedValues"`
	DefaultValue  interface{}   `pulumi:"defaultValue"`
	Description   *string       `pulumi:"description"`
	DisplayName   *string       `pulumi:"displayName"`
	StrongType    *string       `pulumi:"strongType"`
	Type          string        `pulumi:"type"`
}





type ParameterDefinitionInput interface {
	pulumi.Input

	ToParameterDefinitionOutput() ParameterDefinitionOutput
	ToParameterDefinitionOutputWithContext(context.Context) ParameterDefinitionOutput
}

type ParameterDefinitionArgs struct {
	AllowedValues pulumi.ArrayInput     `pulumi:"allowedValues"`
	DefaultValue  pulumi.Input          `pulumi:"defaultValue"`
	Description   pulumi.StringPtrInput `pulumi:"description"`
	DisplayName   pulumi.StringPtrInput `pulumi:"displayName"`
	StrongType    pulumi.StringPtrInput `pulumi:"strongType"`
	Type          pulumi.StringInput    `pulumi:"type"`
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





type ParameterDefinitionMapInput interface {
	pulumi.Input

	ToParameterDefinitionMapOutput() ParameterDefinitionMapOutput
	ToParameterDefinitionMapOutputWithContext(context.Context) ParameterDefinitionMapOutput
}

type ParameterDefinitionMap map[string]ParameterDefinitionInput

func (ParameterDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinition)(nil)).Elem()
}

func (i ParameterDefinitionMap) ToParameterDefinitionMapOutput() ParameterDefinitionMapOutput {
	return i.ToParameterDefinitionMapOutputWithContext(context.Background())
}

func (i ParameterDefinitionMap) ToParameterDefinitionMapOutputWithContext(ctx context.Context) ParameterDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionMapOutput)
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

func (o ParameterDefinitionOutput) AllowedValues() pulumi.ArrayOutput {
	return o.ApplyT(func(v ParameterDefinition) []interface{} { return v.AllowedValues }).(pulumi.ArrayOutput)
}

func (o ParameterDefinitionOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterDefinition) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinition) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinition) *string { return v.StrongType }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDefinition) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterDefinitionMapOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinition)(nil)).Elem()
}

func (o ParameterDefinitionMapOutput) ToParameterDefinitionMapOutput() ParameterDefinitionMapOutput {
	return o
}

func (o ParameterDefinitionMapOutput) ToParameterDefinitionMapOutputWithContext(ctx context.Context) ParameterDefinitionMapOutput {
	return o
}

func (o ParameterDefinitionMapOutput) MapIndex(k pulumi.StringInput) ParameterDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterDefinition {
		return vs[0].(map[string]ParameterDefinition)[vs[1].(string)]
	}).(ParameterDefinitionOutput)
}

type ParameterDefinitionResponse struct {
	AllowedValues []interface{} `pulumi:"allowedValues"`
	DefaultValue  interface{}   `pulumi:"defaultValue"`
	Description   *string       `pulumi:"description"`
	DisplayName   *string       `pulumi:"displayName"`
	StrongType    *string       `pulumi:"strongType"`
	Type          string        `pulumi:"type"`
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

func (o ParameterDefinitionResponseOutput) AllowedValues() pulumi.ArrayOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) []interface{} { return v.AllowedValues }).(pulumi.ArrayOutput)
}

func (o ParameterDefinitionResponseOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterDefinitionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionResponseOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) *string { return v.StrongType }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterDefinitionResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinitionResponse)(nil)).Elem()
}

func (o ParameterDefinitionResponseMapOutput) ToParameterDefinitionResponseMapOutput() ParameterDefinitionResponseMapOutput {
	return o
}

func (o ParameterDefinitionResponseMapOutput) ToParameterDefinitionResponseMapOutputWithContext(ctx context.Context) ParameterDefinitionResponseMapOutput {
	return o
}

func (o ParameterDefinitionResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterDefinitionResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterDefinitionResponse {
		return vs[0].(map[string]ParameterDefinitionResponse)[vs[1].(string)]
	}).(ParameterDefinitionResponseOutput)
}

type ParameterValue struct {
	Reference *SecretValueReference `pulumi:"reference"`
	Value     interface{}           `pulumi:"value"`
}





type ParameterValueInput interface {
	pulumi.Input

	ToParameterValueOutput() ParameterValueOutput
	ToParameterValueOutputWithContext(context.Context) ParameterValueOutput
}

type ParameterValueArgs struct {
	Reference SecretValueReferencePtrInput `pulumi:"reference"`
	Value     pulumi.Input                 `pulumi:"value"`
}

func (ParameterValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValue)(nil)).Elem()
}

func (i ParameterValueArgs) ToParameterValueOutput() ParameterValueOutput {
	return i.ToParameterValueOutputWithContext(context.Background())
}

func (i ParameterValueArgs) ToParameterValueOutputWithContext(ctx context.Context) ParameterValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterValueOutput)
}





type ParameterValueMapInput interface {
	pulumi.Input

	ToParameterValueMapOutput() ParameterValueMapOutput
	ToParameterValueMapOutputWithContext(context.Context) ParameterValueMapOutput
}

type ParameterValueMap map[string]ParameterValueInput

func (ParameterValueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValue)(nil)).Elem()
}

func (i ParameterValueMap) ToParameterValueMapOutput() ParameterValueMapOutput {
	return i.ToParameterValueMapOutputWithContext(context.Background())
}

func (i ParameterValueMap) ToParameterValueMapOutputWithContext(ctx context.Context) ParameterValueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterValueMapOutput)
}

type ParameterValueOutput struct{ *pulumi.OutputState }

func (ParameterValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValue)(nil)).Elem()
}

func (o ParameterValueOutput) ToParameterValueOutput() ParameterValueOutput {
	return o
}

func (o ParameterValueOutput) ToParameterValueOutputWithContext(ctx context.Context) ParameterValueOutput {
	return o
}

func (o ParameterValueOutput) Reference() SecretValueReferencePtrOutput {
	return o.ApplyT(func(v ParameterValue) *SecretValueReference { return v.Reference }).(SecretValueReferencePtrOutput)
}

func (o ParameterValueOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterValue) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ParameterValueMapOutput struct{ *pulumi.OutputState }

func (ParameterValueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValue)(nil)).Elem()
}

func (o ParameterValueMapOutput) ToParameterValueMapOutput() ParameterValueMapOutput {
	return o
}

func (o ParameterValueMapOutput) ToParameterValueMapOutputWithContext(ctx context.Context) ParameterValueMapOutput {
	return o
}

func (o ParameterValueMapOutput) MapIndex(k pulumi.StringInput) ParameterValueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterValue {
		return vs[0].(map[string]ParameterValue)[vs[1].(string)]
	}).(ParameterValueOutput)
}

type ParameterValueResponse struct {
	Reference *SecretValueReferenceResponse `pulumi:"reference"`
	Value     interface{}                   `pulumi:"value"`
}

type ParameterValueResponseOutput struct{ *pulumi.OutputState }

func (ParameterValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValueResponse)(nil)).Elem()
}

func (o ParameterValueResponseOutput) ToParameterValueResponseOutput() ParameterValueResponseOutput {
	return o
}

func (o ParameterValueResponseOutput) ToParameterValueResponseOutputWithContext(ctx context.Context) ParameterValueResponseOutput {
	return o
}

func (o ParameterValueResponseOutput) Reference() SecretValueReferenceResponsePtrOutput {
	return o.ApplyT(func(v ParameterValueResponse) *SecretValueReferenceResponse { return v.Reference }).(SecretValueReferenceResponsePtrOutput)
}

func (o ParameterValueResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterValueResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ParameterValueResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterValueResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValueResponse)(nil)).Elem()
}

func (o ParameterValueResponseMapOutput) ToParameterValueResponseMapOutput() ParameterValueResponseMapOutput {
	return o
}

func (o ParameterValueResponseMapOutput) ToParameterValueResponseMapOutputWithContext(ctx context.Context) ParameterValueResponseMapOutput {
	return o
}

func (o ParameterValueResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterValueResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterValueResponse {
		return vs[0].(map[string]ParameterValueResponse)[vs[1].(string)]
	}).(ParameterValueResponseOutput)
}

type ResourceGroupDefinition struct {
	DependsOn   []string          `pulumi:"dependsOn"`
	Description *string           `pulumi:"description"`
	DisplayName *string           `pulumi:"displayName"`
	Location    *string           `pulumi:"location"`
	Name        *string           `pulumi:"name"`
	StrongType  *string           `pulumi:"strongType"`
	Tags        map[string]string `pulumi:"tags"`
}





type ResourceGroupDefinitionInput interface {
	pulumi.Input

	ToResourceGroupDefinitionOutput() ResourceGroupDefinitionOutput
	ToResourceGroupDefinitionOutputWithContext(context.Context) ResourceGroupDefinitionOutput
}

type ResourceGroupDefinitionArgs struct {
	DependsOn   pulumi.StringArrayInput `pulumi:"dependsOn"`
	Description pulumi.StringPtrInput   `pulumi:"description"`
	DisplayName pulumi.StringPtrInput   `pulumi:"displayName"`
	Location    pulumi.StringPtrInput   `pulumi:"location"`
	Name        pulumi.StringPtrInput   `pulumi:"name"`
	StrongType  pulumi.StringPtrInput   `pulumi:"strongType"`
	Tags        pulumi.StringMapInput   `pulumi:"tags"`
}

func (ResourceGroupDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupDefinition)(nil)).Elem()
}

func (i ResourceGroupDefinitionArgs) ToResourceGroupDefinitionOutput() ResourceGroupDefinitionOutput {
	return i.ToResourceGroupDefinitionOutputWithContext(context.Background())
}

func (i ResourceGroupDefinitionArgs) ToResourceGroupDefinitionOutputWithContext(ctx context.Context) ResourceGroupDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupDefinitionOutput)
}





type ResourceGroupDefinitionMapInput interface {
	pulumi.Input

	ToResourceGroupDefinitionMapOutput() ResourceGroupDefinitionMapOutput
	ToResourceGroupDefinitionMapOutputWithContext(context.Context) ResourceGroupDefinitionMapOutput
}

type ResourceGroupDefinitionMap map[string]ResourceGroupDefinitionInput

func (ResourceGroupDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupDefinition)(nil)).Elem()
}

func (i ResourceGroupDefinitionMap) ToResourceGroupDefinitionMapOutput() ResourceGroupDefinitionMapOutput {
	return i.ToResourceGroupDefinitionMapOutputWithContext(context.Background())
}

func (i ResourceGroupDefinitionMap) ToResourceGroupDefinitionMapOutputWithContext(ctx context.Context) ResourceGroupDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupDefinitionMapOutput)
}

type ResourceGroupDefinitionOutput struct{ *pulumi.OutputState }

func (ResourceGroupDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupDefinition)(nil)).Elem()
}

func (o ResourceGroupDefinitionOutput) ToResourceGroupDefinitionOutput() ResourceGroupDefinitionOutput {
	return o
}

func (o ResourceGroupDefinitionOutput) ToResourceGroupDefinitionOutputWithContext(ctx context.Context) ResourceGroupDefinitionOutput {
	return o
}

func (o ResourceGroupDefinitionOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) []string { return v.DependsOn }).(pulumi.StringArrayOutput)
}

func (o ResourceGroupDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) *string { return v.StrongType }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ResourceGroupDefinitionMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupDefinition)(nil)).Elem()
}

func (o ResourceGroupDefinitionMapOutput) ToResourceGroupDefinitionMapOutput() ResourceGroupDefinitionMapOutput {
	return o
}

func (o ResourceGroupDefinitionMapOutput) ToResourceGroupDefinitionMapOutputWithContext(ctx context.Context) ResourceGroupDefinitionMapOutput {
	return o
}

func (o ResourceGroupDefinitionMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceGroupDefinition {
		return vs[0].(map[string]ResourceGroupDefinition)[vs[1].(string)]
	}).(ResourceGroupDefinitionOutput)
}

type ResourceGroupDefinitionResponse struct {
	DependsOn   []string          `pulumi:"dependsOn"`
	Description *string           `pulumi:"description"`
	DisplayName *string           `pulumi:"displayName"`
	Location    *string           `pulumi:"location"`
	Name        *string           `pulumi:"name"`
	StrongType  *string           `pulumi:"strongType"`
	Tags        map[string]string `pulumi:"tags"`
}

type ResourceGroupDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ResourceGroupDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupDefinitionResponse)(nil)).Elem()
}

func (o ResourceGroupDefinitionResponseOutput) ToResourceGroupDefinitionResponseOutput() ResourceGroupDefinitionResponseOutput {
	return o
}

func (o ResourceGroupDefinitionResponseOutput) ToResourceGroupDefinitionResponseOutputWithContext(ctx context.Context) ResourceGroupDefinitionResponseOutput {
	return o
}

func (o ResourceGroupDefinitionResponseOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) []string { return v.DependsOn }).(pulumi.StringArrayOutput)
}

func (o ResourceGroupDefinitionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionResponseOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) *string { return v.StrongType }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type ResourceGroupDefinitionResponseMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupDefinitionResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupDefinitionResponse)(nil)).Elem()
}

func (o ResourceGroupDefinitionResponseMapOutput) ToResourceGroupDefinitionResponseMapOutput() ResourceGroupDefinitionResponseMapOutput {
	return o
}

func (o ResourceGroupDefinitionResponseMapOutput) ToResourceGroupDefinitionResponseMapOutputWithContext(ctx context.Context) ResourceGroupDefinitionResponseMapOutput {
	return o
}

func (o ResourceGroupDefinitionResponseMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupDefinitionResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceGroupDefinitionResponse {
		return vs[0].(map[string]ResourceGroupDefinitionResponse)[vs[1].(string)]
	}).(ResourceGroupDefinitionResponseOutput)
}

type ResourceGroupValue struct {
	Location *string `pulumi:"location"`
	Name     *string `pulumi:"name"`
}





type ResourceGroupValueInput interface {
	pulumi.Input

	ToResourceGroupValueOutput() ResourceGroupValueOutput
	ToResourceGroupValueOutputWithContext(context.Context) ResourceGroupValueOutput
}

type ResourceGroupValueArgs struct {
	Location pulumi.StringPtrInput `pulumi:"location"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (ResourceGroupValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupValue)(nil)).Elem()
}

func (i ResourceGroupValueArgs) ToResourceGroupValueOutput() ResourceGroupValueOutput {
	return i.ToResourceGroupValueOutputWithContext(context.Background())
}

func (i ResourceGroupValueArgs) ToResourceGroupValueOutputWithContext(ctx context.Context) ResourceGroupValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupValueOutput)
}





type ResourceGroupValueMapInput interface {
	pulumi.Input

	ToResourceGroupValueMapOutput() ResourceGroupValueMapOutput
	ToResourceGroupValueMapOutputWithContext(context.Context) ResourceGroupValueMapOutput
}

type ResourceGroupValueMap map[string]ResourceGroupValueInput

func (ResourceGroupValueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupValue)(nil)).Elem()
}

func (i ResourceGroupValueMap) ToResourceGroupValueMapOutput() ResourceGroupValueMapOutput {
	return i.ToResourceGroupValueMapOutputWithContext(context.Background())
}

func (i ResourceGroupValueMap) ToResourceGroupValueMapOutputWithContext(ctx context.Context) ResourceGroupValueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupValueMapOutput)
}

type ResourceGroupValueOutput struct{ *pulumi.OutputState }

func (ResourceGroupValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupValue)(nil)).Elem()
}

func (o ResourceGroupValueOutput) ToResourceGroupValueOutput() ResourceGroupValueOutput {
	return o
}

func (o ResourceGroupValueOutput) ToResourceGroupValueOutputWithContext(ctx context.Context) ResourceGroupValueOutput {
	return o
}

func (o ResourceGroupValueOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupValue) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupValueOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupValue) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ResourceGroupValueMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupValueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupValue)(nil)).Elem()
}

func (o ResourceGroupValueMapOutput) ToResourceGroupValueMapOutput() ResourceGroupValueMapOutput {
	return o
}

func (o ResourceGroupValueMapOutput) ToResourceGroupValueMapOutputWithContext(ctx context.Context) ResourceGroupValueMapOutput {
	return o
}

func (o ResourceGroupValueMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupValueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceGroupValue {
		return vs[0].(map[string]ResourceGroupValue)[vs[1].(string)]
	}).(ResourceGroupValueOutput)
}

type ResourceGroupValueResponse struct {
	Location *string `pulumi:"location"`
	Name     *string `pulumi:"name"`
}

type ResourceGroupValueResponseOutput struct{ *pulumi.OutputState }

func (ResourceGroupValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupValueResponse)(nil)).Elem()
}

func (o ResourceGroupValueResponseOutput) ToResourceGroupValueResponseOutput() ResourceGroupValueResponseOutput {
	return o
}

func (o ResourceGroupValueResponseOutput) ToResourceGroupValueResponseOutputWithContext(ctx context.Context) ResourceGroupValueResponseOutput {
	return o
}

func (o ResourceGroupValueResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupValueResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupValueResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupValueResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ResourceGroupValueResponseMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupValueResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupValueResponse)(nil)).Elem()
}

func (o ResourceGroupValueResponseMapOutput) ToResourceGroupValueResponseMapOutput() ResourceGroupValueResponseMapOutput {
	return o
}

func (o ResourceGroupValueResponseMapOutput) ToResourceGroupValueResponseMapOutputWithContext(ctx context.Context) ResourceGroupValueResponseMapOutput {
	return o
}

func (o ResourceGroupValueResponseMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupValueResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceGroupValueResponse {
		return vs[0].(map[string]ResourceGroupValueResponse)[vs[1].(string)]
	}).(ResourceGroupValueResponseOutput)
}

type SecretValueReference struct {
	KeyVault      KeyVaultReference `pulumi:"keyVault"`
	SecretName    string            `pulumi:"secretName"`
	SecretVersion *string           `pulumi:"secretVersion"`
}





type SecretValueReferenceInput interface {
	pulumi.Input

	ToSecretValueReferenceOutput() SecretValueReferenceOutput
	ToSecretValueReferenceOutputWithContext(context.Context) SecretValueReferenceOutput
}

type SecretValueReferenceArgs struct {
	KeyVault      KeyVaultReferenceInput `pulumi:"keyVault"`
	SecretName    pulumi.StringInput     `pulumi:"secretName"`
	SecretVersion pulumi.StringPtrInput  `pulumi:"secretVersion"`
}

func (SecretValueReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretValueReference)(nil)).Elem()
}

func (i SecretValueReferenceArgs) ToSecretValueReferenceOutput() SecretValueReferenceOutput {
	return i.ToSecretValueReferenceOutputWithContext(context.Background())
}

func (i SecretValueReferenceArgs) ToSecretValueReferenceOutputWithContext(ctx context.Context) SecretValueReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretValueReferenceOutput)
}

func (i SecretValueReferenceArgs) ToSecretValueReferencePtrOutput() SecretValueReferencePtrOutput {
	return i.ToSecretValueReferencePtrOutputWithContext(context.Background())
}

func (i SecretValueReferenceArgs) ToSecretValueReferencePtrOutputWithContext(ctx context.Context) SecretValueReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretValueReferenceOutput).ToSecretValueReferencePtrOutputWithContext(ctx)
}









type SecretValueReferencePtrInput interface {
	pulumi.Input

	ToSecretValueReferencePtrOutput() SecretValueReferencePtrOutput
	ToSecretValueReferencePtrOutputWithContext(context.Context) SecretValueReferencePtrOutput
}

type secretValueReferencePtrType SecretValueReferenceArgs

func SecretValueReferencePtr(v *SecretValueReferenceArgs) SecretValueReferencePtrInput {
	return (*secretValueReferencePtrType)(v)
}

func (*secretValueReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretValueReference)(nil)).Elem()
}

func (i *secretValueReferencePtrType) ToSecretValueReferencePtrOutput() SecretValueReferencePtrOutput {
	return i.ToSecretValueReferencePtrOutputWithContext(context.Background())
}

func (i *secretValueReferencePtrType) ToSecretValueReferencePtrOutputWithContext(ctx context.Context) SecretValueReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretValueReferencePtrOutput)
}

type SecretValueReferenceOutput struct{ *pulumi.OutputState }

func (SecretValueReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretValueReference)(nil)).Elem()
}

func (o SecretValueReferenceOutput) ToSecretValueReferenceOutput() SecretValueReferenceOutput {
	return o
}

func (o SecretValueReferenceOutput) ToSecretValueReferenceOutputWithContext(ctx context.Context) SecretValueReferenceOutput {
	return o
}

func (o SecretValueReferenceOutput) ToSecretValueReferencePtrOutput() SecretValueReferencePtrOutput {
	return o.ToSecretValueReferencePtrOutputWithContext(context.Background())
}

func (o SecretValueReferenceOutput) ToSecretValueReferencePtrOutputWithContext(ctx context.Context) SecretValueReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretValueReference) *SecretValueReference {
		return &v
	}).(SecretValueReferencePtrOutput)
}

func (o SecretValueReferenceOutput) KeyVault() KeyVaultReferenceOutput {
	return o.ApplyT(func(v SecretValueReference) KeyVaultReference { return v.KeyVault }).(KeyVaultReferenceOutput)
}

func (o SecretValueReferenceOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v SecretValueReference) string { return v.SecretName }).(pulumi.StringOutput)
}

func (o SecretValueReferenceOutput) SecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretValueReference) *string { return v.SecretVersion }).(pulumi.StringPtrOutput)
}

type SecretValueReferencePtrOutput struct{ *pulumi.OutputState }

func (SecretValueReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretValueReference)(nil)).Elem()
}

func (o SecretValueReferencePtrOutput) ToSecretValueReferencePtrOutput() SecretValueReferencePtrOutput {
	return o
}

func (o SecretValueReferencePtrOutput) ToSecretValueReferencePtrOutputWithContext(ctx context.Context) SecretValueReferencePtrOutput {
	return o
}

func (o SecretValueReferencePtrOutput) Elem() SecretValueReferenceOutput {
	return o.ApplyT(func(v *SecretValueReference) SecretValueReference {
		if v != nil {
			return *v
		}
		var ret SecretValueReference
		return ret
	}).(SecretValueReferenceOutput)
}

func (o SecretValueReferencePtrOutput) KeyVault() KeyVaultReferencePtrOutput {
	return o.ApplyT(func(v *SecretValueReference) *KeyVaultReference {
		if v == nil {
			return nil
		}
		return &v.KeyVault
	}).(KeyVaultReferencePtrOutput)
}

func (o SecretValueReferencePtrOutput) SecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretValueReference) *string {
		if v == nil {
			return nil
		}
		return &v.SecretName
	}).(pulumi.StringPtrOutput)
}

func (o SecretValueReferencePtrOutput) SecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretValueReference) *string {
		if v == nil {
			return nil
		}
		return v.SecretVersion
	}).(pulumi.StringPtrOutput)
}

type SecretValueReferenceResponse struct {
	KeyVault      KeyVaultReferenceResponse `pulumi:"keyVault"`
	SecretName    string                    `pulumi:"secretName"`
	SecretVersion *string                   `pulumi:"secretVersion"`
}

type SecretValueReferenceResponseOutput struct{ *pulumi.OutputState }

func (SecretValueReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretValueReferenceResponse)(nil)).Elem()
}

func (o SecretValueReferenceResponseOutput) ToSecretValueReferenceResponseOutput() SecretValueReferenceResponseOutput {
	return o
}

func (o SecretValueReferenceResponseOutput) ToSecretValueReferenceResponseOutputWithContext(ctx context.Context) SecretValueReferenceResponseOutput {
	return o
}

func (o SecretValueReferenceResponseOutput) KeyVault() KeyVaultReferenceResponseOutput {
	return o.ApplyT(func(v SecretValueReferenceResponse) KeyVaultReferenceResponse { return v.KeyVault }).(KeyVaultReferenceResponseOutput)
}

func (o SecretValueReferenceResponseOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v SecretValueReferenceResponse) string { return v.SecretName }).(pulumi.StringOutput)
}

func (o SecretValueReferenceResponseOutput) SecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretValueReferenceResponse) *string { return v.SecretVersion }).(pulumi.StringPtrOutput)
}

type SecretValueReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (SecretValueReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretValueReferenceResponse)(nil)).Elem()
}

func (o SecretValueReferenceResponsePtrOutput) ToSecretValueReferenceResponsePtrOutput() SecretValueReferenceResponsePtrOutput {
	return o
}

func (o SecretValueReferenceResponsePtrOutput) ToSecretValueReferenceResponsePtrOutputWithContext(ctx context.Context) SecretValueReferenceResponsePtrOutput {
	return o
}

func (o SecretValueReferenceResponsePtrOutput) Elem() SecretValueReferenceResponseOutput {
	return o.ApplyT(func(v *SecretValueReferenceResponse) SecretValueReferenceResponse {
		if v != nil {
			return *v
		}
		var ret SecretValueReferenceResponse
		return ret
	}).(SecretValueReferenceResponseOutput)
}

func (o SecretValueReferenceResponsePtrOutput) KeyVault() KeyVaultReferenceResponsePtrOutput {
	return o.ApplyT(func(v *SecretValueReferenceResponse) *KeyVaultReferenceResponse {
		if v == nil {
			return nil
		}
		return &v.KeyVault
	}).(KeyVaultReferenceResponsePtrOutput)
}

func (o SecretValueReferenceResponsePtrOutput) SecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretValueReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecretName
	}).(pulumi.StringPtrOutput)
}

func (o SecretValueReferenceResponsePtrOutput) SecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretValueReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecretVersion
	}).(pulumi.StringPtrOutput)
}

type UserAssignedIdentity struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}





type UserAssignedIdentityInput interface {
	pulumi.Input

	ToUserAssignedIdentityOutput() UserAssignedIdentityOutput
	ToUserAssignedIdentityOutputWithContext(context.Context) UserAssignedIdentityOutput
}

type UserAssignedIdentityArgs struct {
	ClientId    pulumi.StringPtrInput `pulumi:"clientId"`
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (UserAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentity)(nil)).Elem()
}

func (i UserAssignedIdentityArgs) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return i.ToUserAssignedIdentityOutputWithContext(context.Background())
}

func (i UserAssignedIdentityArgs) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityOutput)
}





type UserAssignedIdentityMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityMapOutput() UserAssignedIdentityMapOutput
	ToUserAssignedIdentityMapOutputWithContext(context.Context) UserAssignedIdentityMapOutput
}

type UserAssignedIdentityMap map[string]UserAssignedIdentityInput

func (UserAssignedIdentityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentity)(nil)).Elem()
}

func (i UserAssignedIdentityMap) ToUserAssignedIdentityMapOutput() UserAssignedIdentityMapOutput {
	return i.ToUserAssignedIdentityMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityMap) ToUserAssignedIdentityMapOutputWithContext(ctx context.Context) UserAssignedIdentityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityMapOutput)
}

type UserAssignedIdentityOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentity)(nil)).Elem()
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentity) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserAssignedIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentity)(nil)).Elem()
}

func (o UserAssignedIdentityMapOutput) ToUserAssignedIdentityMapOutput() UserAssignedIdentityMapOutput {
	return o
}

func (o UserAssignedIdentityMapOutput) ToUserAssignedIdentityMapOutputWithContext(ctx context.Context) UserAssignedIdentityMapOutput {
	return o
}

func (o UserAssignedIdentityMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentity {
		return vs[0].(map[string]UserAssignedIdentity)[vs[1].(string)]
	}).(UserAssignedIdentityOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AssignmentLockSettingsOutput{})
	pulumi.RegisterOutputType(AssignmentLockSettingsPtrOutput{})
	pulumi.RegisterOutputType(AssignmentLockSettingsResponseOutput{})
	pulumi.RegisterOutputType(AssignmentLockSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AssignmentStatusResponseOutput{})
	pulumi.RegisterOutputType(BlueprintStatusResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionMapOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionResponseMapOutput{})
	pulumi.RegisterOutputType(ParameterValueOutput{})
	pulumi.RegisterOutputType(ParameterValueMapOutput{})
	pulumi.RegisterOutputType(ParameterValueResponseOutput{})
	pulumi.RegisterOutputType(ParameterValueResponseMapOutput{})
	pulumi.RegisterOutputType(ResourceGroupDefinitionOutput{})
	pulumi.RegisterOutputType(ResourceGroupDefinitionMapOutput{})
	pulumi.RegisterOutputType(ResourceGroupDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ResourceGroupDefinitionResponseMapOutput{})
	pulumi.RegisterOutputType(ResourceGroupValueOutput{})
	pulumi.RegisterOutputType(ResourceGroupValueMapOutput{})
	pulumi.RegisterOutputType(ResourceGroupValueResponseOutput{})
	pulumi.RegisterOutputType(ResourceGroupValueResponseMapOutput{})
	pulumi.RegisterOutputType(SecretValueReferenceOutput{})
	pulumi.RegisterOutputType(SecretValueReferencePtrOutput{})
	pulumi.RegisterOutputType(SecretValueReferenceResponseOutput{})
	pulumi.RegisterOutputType(SecretValueReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityMapOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}
