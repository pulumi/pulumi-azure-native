


package v20221012preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentRoleResponse struct {
	Description string `pulumi:"description"`
	RoleName    string `pulumi:"roleName"`
}

type EnvironmentRoleResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentRoleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentRoleResponse)(nil)).Elem()
}

func (o EnvironmentRoleResponseOutput) ToEnvironmentRoleResponseOutput() EnvironmentRoleResponseOutput {
	return o
}

func (o EnvironmentRoleResponseOutput) ToEnvironmentRoleResponseOutputWithContext(ctx context.Context) EnvironmentRoleResponseOutput {
	return o
}

func (o EnvironmentRoleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentRoleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o EnvironmentRoleResponseOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentRoleResponse) string { return v.RoleName }).(pulumi.StringOutput)
}

type EnvironmentRoleResponseMapOutput struct{ *pulumi.OutputState }

func (EnvironmentRoleResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EnvironmentRoleResponse)(nil)).Elem()
}

func (o EnvironmentRoleResponseMapOutput) ToEnvironmentRoleResponseMapOutput() EnvironmentRoleResponseMapOutput {
	return o
}

func (o EnvironmentRoleResponseMapOutput) ToEnvironmentRoleResponseMapOutputWithContext(ctx context.Context) EnvironmentRoleResponseMapOutput {
	return o
}

func (o EnvironmentRoleResponseMapOutput) MapIndex(k pulumi.StringInput) EnvironmentRoleResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EnvironmentRoleResponse {
		return vs[0].(map[string]EnvironmentRoleResponse)[vs[1].(string)]
	}).(EnvironmentRoleResponseOutput)
}

type GitCatalog struct {
	Branch           *string `pulumi:"branch"`
	Path             *string `pulumi:"path"`
	SecretIdentifier *string `pulumi:"secretIdentifier"`
	Uri              *string `pulumi:"uri"`
}





type GitCatalogInput interface {
	pulumi.Input

	ToGitCatalogOutput() GitCatalogOutput
	ToGitCatalogOutputWithContext(context.Context) GitCatalogOutput
}

type GitCatalogArgs struct {
	Branch           pulumi.StringPtrInput `pulumi:"branch"`
	Path             pulumi.StringPtrInput `pulumi:"path"`
	SecretIdentifier pulumi.StringPtrInput `pulumi:"secretIdentifier"`
	Uri              pulumi.StringPtrInput `pulumi:"uri"`
}

func (GitCatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitCatalog)(nil)).Elem()
}

func (i GitCatalogArgs) ToGitCatalogOutput() GitCatalogOutput {
	return i.ToGitCatalogOutputWithContext(context.Background())
}

func (i GitCatalogArgs) ToGitCatalogOutputWithContext(ctx context.Context) GitCatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitCatalogOutput)
}

func (i GitCatalogArgs) ToGitCatalogPtrOutput() GitCatalogPtrOutput {
	return i.ToGitCatalogPtrOutputWithContext(context.Background())
}

func (i GitCatalogArgs) ToGitCatalogPtrOutputWithContext(ctx context.Context) GitCatalogPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitCatalogOutput).ToGitCatalogPtrOutputWithContext(ctx)
}









type GitCatalogPtrInput interface {
	pulumi.Input

	ToGitCatalogPtrOutput() GitCatalogPtrOutput
	ToGitCatalogPtrOutputWithContext(context.Context) GitCatalogPtrOutput
}

type gitCatalogPtrType GitCatalogArgs

func GitCatalogPtr(v *GitCatalogArgs) GitCatalogPtrInput {
	return (*gitCatalogPtrType)(v)
}

func (*gitCatalogPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitCatalog)(nil)).Elem()
}

func (i *gitCatalogPtrType) ToGitCatalogPtrOutput() GitCatalogPtrOutput {
	return i.ToGitCatalogPtrOutputWithContext(context.Background())
}

func (i *gitCatalogPtrType) ToGitCatalogPtrOutputWithContext(ctx context.Context) GitCatalogPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitCatalogPtrOutput)
}

type GitCatalogOutput struct{ *pulumi.OutputState }

func (GitCatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitCatalog)(nil)).Elem()
}

func (o GitCatalogOutput) ToGitCatalogOutput() GitCatalogOutput {
	return o
}

func (o GitCatalogOutput) ToGitCatalogOutputWithContext(ctx context.Context) GitCatalogOutput {
	return o
}

func (o GitCatalogOutput) ToGitCatalogPtrOutput() GitCatalogPtrOutput {
	return o.ToGitCatalogPtrOutputWithContext(context.Background())
}

func (o GitCatalogOutput) ToGitCatalogPtrOutputWithContext(ctx context.Context) GitCatalogPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitCatalog) *GitCatalog {
		return &v
	}).(GitCatalogPtrOutput)
}

func (o GitCatalogOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitCatalog) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o GitCatalogOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitCatalog) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o GitCatalogOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitCatalog) *string { return v.SecretIdentifier }).(pulumi.StringPtrOutput)
}

func (o GitCatalogOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitCatalog) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type GitCatalogPtrOutput struct{ *pulumi.OutputState }

func (GitCatalogPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitCatalog)(nil)).Elem()
}

func (o GitCatalogPtrOutput) ToGitCatalogPtrOutput() GitCatalogPtrOutput {
	return o
}

func (o GitCatalogPtrOutput) ToGitCatalogPtrOutputWithContext(ctx context.Context) GitCatalogPtrOutput {
	return o
}

func (o GitCatalogPtrOutput) Elem() GitCatalogOutput {
	return o.ApplyT(func(v *GitCatalog) GitCatalog {
		if v != nil {
			return *v
		}
		var ret GitCatalog
		return ret
	}).(GitCatalogOutput)
}

func (o GitCatalogPtrOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitCatalog) *string {
		if v == nil {
			return nil
		}
		return v.Branch
	}).(pulumi.StringPtrOutput)
}

func (o GitCatalogPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitCatalog) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o GitCatalogPtrOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitCatalog) *string {
		if v == nil {
			return nil
		}
		return v.SecretIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o GitCatalogPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitCatalog) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type GitCatalogResponse struct {
	Branch           *string `pulumi:"branch"`
	Path             *string `pulumi:"path"`
	SecretIdentifier *string `pulumi:"secretIdentifier"`
	Uri              *string `pulumi:"uri"`
}

type GitCatalogResponseOutput struct{ *pulumi.OutputState }

func (GitCatalogResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitCatalogResponse)(nil)).Elem()
}

func (o GitCatalogResponseOutput) ToGitCatalogResponseOutput() GitCatalogResponseOutput {
	return o
}

func (o GitCatalogResponseOutput) ToGitCatalogResponseOutputWithContext(ctx context.Context) GitCatalogResponseOutput {
	return o
}

func (o GitCatalogResponseOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitCatalogResponse) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o GitCatalogResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitCatalogResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o GitCatalogResponseOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitCatalogResponse) *string { return v.SecretIdentifier }).(pulumi.StringPtrOutput)
}

func (o GitCatalogResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitCatalogResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type GitCatalogResponsePtrOutput struct{ *pulumi.OutputState }

func (GitCatalogResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitCatalogResponse)(nil)).Elem()
}

func (o GitCatalogResponsePtrOutput) ToGitCatalogResponsePtrOutput() GitCatalogResponsePtrOutput {
	return o
}

func (o GitCatalogResponsePtrOutput) ToGitCatalogResponsePtrOutputWithContext(ctx context.Context) GitCatalogResponsePtrOutput {
	return o
}

func (o GitCatalogResponsePtrOutput) Elem() GitCatalogResponseOutput {
	return o.ApplyT(func(v *GitCatalogResponse) GitCatalogResponse {
		if v != nil {
			return *v
		}
		var ret GitCatalogResponse
		return ret
	}).(GitCatalogResponseOutput)
}

func (o GitCatalogResponsePtrOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitCatalogResponse) *string {
		if v == nil {
			return nil
		}
		return v.Branch
	}).(pulumi.StringPtrOutput)
}

func (o GitCatalogResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitCatalogResponse) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o GitCatalogResponsePtrOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitCatalogResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecretIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o GitCatalogResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitCatalogResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type ImageReference struct {
	Id        *string `pulumi:"id"`
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
}





type ImageReferenceInput interface {
	pulumi.Input

	ToImageReferenceOutput() ImageReferenceOutput
	ToImageReferenceOutputWithContext(context.Context) ImageReferenceOutput
}

type ImageReferenceArgs struct {
	Id        pulumi.StringPtrInput `pulumi:"id"`
	Offer     pulumi.StringPtrInput `pulumi:"offer"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput `pulumi:"sku"`
}

func (ImageReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (i ImageReferenceArgs) ToImageReferenceOutput() ImageReferenceOutput {
	return i.ToImageReferenceOutputWithContext(context.Background())
}

func (i ImageReferenceArgs) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceOutput)
}

type ImageReferenceOutput struct{ *pulumi.OutputState }

func (ImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (o ImageReferenceOutput) ToImageReferenceOutput() ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

type ImageReferenceResponse struct {
	ExactVersion string  `pulumi:"exactVersion"`
	Id           *string `pulumi:"id"`
	Offer        *string `pulumi:"offer"`
	Publisher    *string `pulumi:"publisher"`
	Sku          *string `pulumi:"sku"`
}

type ImageReferenceResponseOutput struct{ *pulumi.OutputState }

func (ImageReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReferenceResponse)(nil)).Elem()
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutput() ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutputWithContext(ctx context.Context) ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) ExactVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ImageReferenceResponse) string { return v.ExactVersion }).(pulumi.StringOutput)
}

func (o ImageReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

type ImageValidationErrorDetailsResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}

type ImageValidationErrorDetailsResponseOutput struct{ *pulumi.OutputState }

func (ImageValidationErrorDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageValidationErrorDetailsResponse)(nil)).Elem()
}

func (o ImageValidationErrorDetailsResponseOutput) ToImageValidationErrorDetailsResponseOutput() ImageValidationErrorDetailsResponseOutput {
	return o
}

func (o ImageValidationErrorDetailsResponseOutput) ToImageValidationErrorDetailsResponseOutputWithContext(ctx context.Context) ImageValidationErrorDetailsResponseOutput {
	return o
}

func (o ImageValidationErrorDetailsResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageValidationErrorDetailsResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ImageValidationErrorDetailsResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageValidationErrorDetailsResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ManagedServiceIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
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

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
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

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
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

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment struct {
	Roles map[string]interface{} `pulumi:"roles"`
}





type ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentInput interface {
	pulumi.Input

	ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput() ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput
	ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutputWithContext(context.Context) ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput
}

type ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentArgs struct {
	Roles pulumi.MapInput `pulumi:"roles"`
}

func (ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment)(nil)).Elem()
}

func (i ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentArgs) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput() ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput {
	return i.ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutputWithContext(context.Background())
}

func (i ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentArgs) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutputWithContext(ctx context.Context) ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput)
}

func (i ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentArgs) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput() ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput {
	return i.ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutputWithContext(context.Background())
}

func (i ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentArgs) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutputWithContext(ctx context.Context) ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput).ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutputWithContext(ctx)
}









type ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrInput interface {
	pulumi.Input

	ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput() ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput
	ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutputWithContext(context.Context) ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput
}

type projectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrType ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentArgs

func ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtr(v *ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentArgs) ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrInput {
	return (*projectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrType)(v)
}

func (*projectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment)(nil)).Elem()
}

func (i *projectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrType) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput() ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput {
	return i.ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutputWithContext(context.Background())
}

func (i *projectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrType) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutputWithContext(ctx context.Context) ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput)
}

type ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput struct{ *pulumi.OutputState }

func (ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment)(nil)).Elem()
}

func (o ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput() ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput {
	return o
}

func (o ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutputWithContext(ctx context.Context) ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput {
	return o
}

func (o ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput() ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput {
	return o.ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutputWithContext(context.Background())
}

func (o ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutputWithContext(ctx context.Context) ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment) *ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment {
		return &v
	}).(ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput)
}

func (o ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput) Roles() pulumi.MapOutput {
	return o.ApplyT(func(v ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment) map[string]interface{} {
		return v.Roles
	}).(pulumi.MapOutput)
}

type ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput struct{ *pulumi.OutputState }

func (ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment)(nil)).Elem()
}

func (o ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput() ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput {
	return o
}

func (o ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput) ToProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutputWithContext(ctx context.Context) ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput {
	return o
}

func (o ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput) Elem() ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput {
	return o.ApplyT(func(v *ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment) ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment {
		if v != nil {
			return *v
		}
		var ret ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment
		return ret
	}).(ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput)
}

func (o ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput) Roles() pulumi.MapOutput {
	return o.ApplyT(func(v *ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Roles
	}).(pulumi.MapOutput)
}

type ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment struct {
	Roles map[string]EnvironmentRoleResponse `pulumi:"roles"`
}

type ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutput struct{ *pulumi.OutputState }

func (ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment)(nil)).Elem()
}

func (o ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutput) ToProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutput() ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutput {
	return o
}

func (o ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutput) ToProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutputWithContext(ctx context.Context) ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutput {
	return o
}

func (o ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutput) Roles() EnvironmentRoleResponseMapOutput {
	return o.ApplyT(func(v ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment) map[string]EnvironmentRoleResponse {
		return v.Roles
	}).(EnvironmentRoleResponseMapOutput)
}

type ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput struct{ *pulumi.OutputState }

func (ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment)(nil)).Elem()
}

func (o ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput) ToProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput() ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput {
	return o
}

func (o ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput) ToProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutputWithContext(ctx context.Context) ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput {
	return o
}

func (o ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput) Elem() ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutput {
	return o.ApplyT(func(v *ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment) ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment {
		if v != nil {
			return *v
		}
		var ret ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment
		return ret
	}).(ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutput)
}

func (o ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput) Roles() EnvironmentRoleResponseMapOutput {
	return o.ApplyT(func(v *ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment) map[string]EnvironmentRoleResponse {
		if v == nil {
			return nil
		}
		return v.Roles
	}).(EnvironmentRoleResponseMapOutput)
}

type Sku struct {
	Capacity *int     `pulumi:"capacity"`
	Family   *string  `pulumi:"family"`
	Name     string   `pulumi:"name"`
	Size     *string  `pulumi:"size"`
	Tier     *SkuTier `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     SkuTierPtrInput       `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v Sku) *SkuTier { return v.Tier }).(SkuTierPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
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

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
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

type UserRoleAssignment struct {
	Roles map[string]interface{} `pulumi:"roles"`
}





type UserRoleAssignmentInput interface {
	pulumi.Input

	ToUserRoleAssignmentOutput() UserRoleAssignmentOutput
	ToUserRoleAssignmentOutputWithContext(context.Context) UserRoleAssignmentOutput
}

type UserRoleAssignmentArgs struct {
	Roles pulumi.MapInput `pulumi:"roles"`
}

func (UserRoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRoleAssignment)(nil)).Elem()
}

func (i UserRoleAssignmentArgs) ToUserRoleAssignmentOutput() UserRoleAssignmentOutput {
	return i.ToUserRoleAssignmentOutputWithContext(context.Background())
}

func (i UserRoleAssignmentArgs) ToUserRoleAssignmentOutputWithContext(ctx context.Context) UserRoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRoleAssignmentOutput)
}





type UserRoleAssignmentMapInput interface {
	pulumi.Input

	ToUserRoleAssignmentMapOutput() UserRoleAssignmentMapOutput
	ToUserRoleAssignmentMapOutputWithContext(context.Context) UserRoleAssignmentMapOutput
}

type UserRoleAssignmentMap map[string]UserRoleAssignmentInput

func (UserRoleAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserRoleAssignment)(nil)).Elem()
}

func (i UserRoleAssignmentMap) ToUserRoleAssignmentMapOutput() UserRoleAssignmentMapOutput {
	return i.ToUserRoleAssignmentMapOutputWithContext(context.Background())
}

func (i UserRoleAssignmentMap) ToUserRoleAssignmentMapOutputWithContext(ctx context.Context) UserRoleAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRoleAssignmentMapOutput)
}

type UserRoleAssignmentOutput struct{ *pulumi.OutputState }

func (UserRoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRoleAssignment)(nil)).Elem()
}

func (o UserRoleAssignmentOutput) ToUserRoleAssignmentOutput() UserRoleAssignmentOutput {
	return o
}

func (o UserRoleAssignmentOutput) ToUserRoleAssignmentOutputWithContext(ctx context.Context) UserRoleAssignmentOutput {
	return o
}

func (o UserRoleAssignmentOutput) Roles() pulumi.MapOutput {
	return o.ApplyT(func(v UserRoleAssignment) map[string]interface{} { return v.Roles }).(pulumi.MapOutput)
}

type UserRoleAssignmentMapOutput struct{ *pulumi.OutputState }

func (UserRoleAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserRoleAssignment)(nil)).Elem()
}

func (o UserRoleAssignmentMapOutput) ToUserRoleAssignmentMapOutput() UserRoleAssignmentMapOutput {
	return o
}

func (o UserRoleAssignmentMapOutput) ToUserRoleAssignmentMapOutputWithContext(ctx context.Context) UserRoleAssignmentMapOutput {
	return o
}

func (o UserRoleAssignmentMapOutput) MapIndex(k pulumi.StringInput) UserRoleAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserRoleAssignment {
		return vs[0].(map[string]UserRoleAssignment)[vs[1].(string)]
	}).(UserRoleAssignmentOutput)
}

type UserRoleAssignmentResponse struct {
	Roles map[string]EnvironmentRoleResponse `pulumi:"roles"`
}

type UserRoleAssignmentResponseOutput struct{ *pulumi.OutputState }

func (UserRoleAssignmentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRoleAssignmentResponse)(nil)).Elem()
}

func (o UserRoleAssignmentResponseOutput) ToUserRoleAssignmentResponseOutput() UserRoleAssignmentResponseOutput {
	return o
}

func (o UserRoleAssignmentResponseOutput) ToUserRoleAssignmentResponseOutputWithContext(ctx context.Context) UserRoleAssignmentResponseOutput {
	return o
}

func (o UserRoleAssignmentResponseOutput) Roles() EnvironmentRoleResponseMapOutput {
	return o.ApplyT(func(v UserRoleAssignmentResponse) map[string]EnvironmentRoleResponse { return v.Roles }).(EnvironmentRoleResponseMapOutput)
}

type UserRoleAssignmentResponseMapOutput struct{ *pulumi.OutputState }

func (UserRoleAssignmentResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserRoleAssignmentResponse)(nil)).Elem()
}

func (o UserRoleAssignmentResponseMapOutput) ToUserRoleAssignmentResponseMapOutput() UserRoleAssignmentResponseMapOutput {
	return o
}

func (o UserRoleAssignmentResponseMapOutput) ToUserRoleAssignmentResponseMapOutputWithContext(ctx context.Context) UserRoleAssignmentResponseMapOutput {
	return o
}

func (o UserRoleAssignmentResponseMapOutput) MapIndex(k pulumi.StringInput) UserRoleAssignmentResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserRoleAssignmentResponse {
		return vs[0].(map[string]UserRoleAssignmentResponse)[vs[1].(string)]
	}).(UserRoleAssignmentResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentRoleResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentRoleResponseMapOutput{})
	pulumi.RegisterOutputType(GitCatalogOutput{})
	pulumi.RegisterOutputType(GitCatalogPtrOutput{})
	pulumi.RegisterOutputType(GitCatalogResponseOutput{})
	pulumi.RegisterOutputType(GitCatalogResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponseOutput{})
	pulumi.RegisterOutputType(ImageValidationErrorDetailsResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentOutput{})
	pulumi.RegisterOutputType(ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrOutput{})
	pulumi.RegisterOutputType(ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentOutput{})
	pulumi.RegisterOutputType(ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(UserRoleAssignmentOutput{})
	pulumi.RegisterOutputType(UserRoleAssignmentMapOutput{})
	pulumi.RegisterOutputType(UserRoleAssignmentResponseOutput{})
	pulumi.RegisterOutputType(UserRoleAssignmentResponseMapOutput{})
}
