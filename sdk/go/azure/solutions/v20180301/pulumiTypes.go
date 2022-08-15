


package v20180301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationArtifactResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
	Uri  string `pulumi:"uri"`
}

type ApplicationArtifactResponseOutput struct{ *pulumi.OutputState }

func (ApplicationArtifactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationArtifactResponse)(nil)).Elem()
}

func (o ApplicationArtifactResponseOutput) ToApplicationArtifactResponseOutput() ApplicationArtifactResponseOutput {
	return o
}

func (o ApplicationArtifactResponseOutput) ToApplicationArtifactResponseOutputWithContext(ctx context.Context) ApplicationArtifactResponseOutput {
	return o
}

func (o ApplicationArtifactResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationArtifactResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationArtifactResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationArtifactResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ApplicationArtifactResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationArtifactResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type ApplicationArtifactResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArtifactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationArtifactResponse)(nil)).Elem()
}

func (o ApplicationArtifactResponseArrayOutput) ToApplicationArtifactResponseArrayOutput() ApplicationArtifactResponseArrayOutput {
	return o
}

func (o ApplicationArtifactResponseArrayOutput) ToApplicationArtifactResponseArrayOutputWithContext(ctx context.Context) ApplicationArtifactResponseArrayOutput {
	return o
}

func (o ApplicationArtifactResponseArrayOutput) Index(i pulumi.IntInput) ApplicationArtifactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationArtifactResponse {
		return vs[0].([]ApplicationArtifactResponse)[vs[1].(int)]
	}).(ApplicationArtifactResponseOutput)
}

type ApplicationAuthorization struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}





type ApplicationAuthorizationInput interface {
	pulumi.Input

	ToApplicationAuthorizationOutput() ApplicationAuthorizationOutput
	ToApplicationAuthorizationOutputWithContext(context.Context) ApplicationAuthorizationOutput
}

type ApplicationAuthorizationArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (ApplicationAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationAuthorization)(nil)).Elem()
}

func (i ApplicationAuthorizationArgs) ToApplicationAuthorizationOutput() ApplicationAuthorizationOutput {
	return i.ToApplicationAuthorizationOutputWithContext(context.Background())
}

func (i ApplicationAuthorizationArgs) ToApplicationAuthorizationOutputWithContext(ctx context.Context) ApplicationAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAuthorizationOutput)
}





type ApplicationAuthorizationArrayInput interface {
	pulumi.Input

	ToApplicationAuthorizationArrayOutput() ApplicationAuthorizationArrayOutput
	ToApplicationAuthorizationArrayOutputWithContext(context.Context) ApplicationAuthorizationArrayOutput
}

type ApplicationAuthorizationArray []ApplicationAuthorizationInput

func (ApplicationAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationAuthorization)(nil)).Elem()
}

func (i ApplicationAuthorizationArray) ToApplicationAuthorizationArrayOutput() ApplicationAuthorizationArrayOutput {
	return i.ToApplicationAuthorizationArrayOutputWithContext(context.Background())
}

func (i ApplicationAuthorizationArray) ToApplicationAuthorizationArrayOutputWithContext(ctx context.Context) ApplicationAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAuthorizationArrayOutput)
}

type ApplicationAuthorizationOutput struct{ *pulumi.OutputState }

func (ApplicationAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationAuthorization)(nil)).Elem()
}

func (o ApplicationAuthorizationOutput) ToApplicationAuthorizationOutput() ApplicationAuthorizationOutput {
	return o
}

func (o ApplicationAuthorizationOutput) ToApplicationAuthorizationOutputWithContext(ctx context.Context) ApplicationAuthorizationOutput {
	return o
}

func (o ApplicationAuthorizationOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationAuthorization) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ApplicationAuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationAuthorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type ApplicationAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationAuthorization)(nil)).Elem()
}

func (o ApplicationAuthorizationArrayOutput) ToApplicationAuthorizationArrayOutput() ApplicationAuthorizationArrayOutput {
	return o
}

func (o ApplicationAuthorizationArrayOutput) ToApplicationAuthorizationArrayOutputWithContext(ctx context.Context) ApplicationAuthorizationArrayOutput {
	return o
}

func (o ApplicationAuthorizationArrayOutput) Index(i pulumi.IntInput) ApplicationAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationAuthorization {
		return vs[0].([]ApplicationAuthorization)[vs[1].(int)]
	}).(ApplicationAuthorizationOutput)
}

type ApplicationAuthorizationResponse struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}

type ApplicationAuthorizationResponseOutput struct{ *pulumi.OutputState }

func (ApplicationAuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationAuthorizationResponse)(nil)).Elem()
}

func (o ApplicationAuthorizationResponseOutput) ToApplicationAuthorizationResponseOutput() ApplicationAuthorizationResponseOutput {
	return o
}

func (o ApplicationAuthorizationResponseOutput) ToApplicationAuthorizationResponseOutputWithContext(ctx context.Context) ApplicationAuthorizationResponseOutput {
	return o
}

func (o ApplicationAuthorizationResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationAuthorizationResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ApplicationAuthorizationResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationAuthorizationResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type ApplicationAuthorizationResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationAuthorizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationAuthorizationResponse)(nil)).Elem()
}

func (o ApplicationAuthorizationResponseArrayOutput) ToApplicationAuthorizationResponseArrayOutput() ApplicationAuthorizationResponseArrayOutput {
	return o
}

func (o ApplicationAuthorizationResponseArrayOutput) ToApplicationAuthorizationResponseArrayOutputWithContext(ctx context.Context) ApplicationAuthorizationResponseArrayOutput {
	return o
}

func (o ApplicationAuthorizationResponseArrayOutput) Index(i pulumi.IntInput) ApplicationAuthorizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationAuthorizationResponse {
		return vs[0].([]ApplicationAuthorizationResponse)[vs[1].(int)]
	}).(ApplicationAuthorizationResponseOutput)
}

type ApplicationBillingDetailsDefinitionResponse struct {
	ResourceUsageId *string `pulumi:"resourceUsageId"`
}

type ApplicationBillingDetailsDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ApplicationBillingDetailsDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationBillingDetailsDefinitionResponse)(nil)).Elem()
}

func (o ApplicationBillingDetailsDefinitionResponseOutput) ToApplicationBillingDetailsDefinitionResponseOutput() ApplicationBillingDetailsDefinitionResponseOutput {
	return o
}

func (o ApplicationBillingDetailsDefinitionResponseOutput) ToApplicationBillingDetailsDefinitionResponseOutputWithContext(ctx context.Context) ApplicationBillingDetailsDefinitionResponseOutput {
	return o
}

func (o ApplicationBillingDetailsDefinitionResponseOutput) ResourceUsageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationBillingDetailsDefinitionResponse) *string { return v.ResourceUsageId }).(pulumi.StringPtrOutput)
}

type ApplicationClientDetailsResponse struct {
	ApplicationId *string `pulumi:"applicationId"`
	Oid           *string `pulumi:"oid"`
	Puid          *string `pulumi:"puid"`
}

type ApplicationClientDetailsResponseOutput struct{ *pulumi.OutputState }

func (ApplicationClientDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationClientDetailsResponse)(nil)).Elem()
}

func (o ApplicationClientDetailsResponseOutput) ToApplicationClientDetailsResponseOutput() ApplicationClientDetailsResponseOutput {
	return o
}

func (o ApplicationClientDetailsResponseOutput) ToApplicationClientDetailsResponseOutputWithContext(ctx context.Context) ApplicationClientDetailsResponseOutput {
	return o
}

func (o ApplicationClientDetailsResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationClientDetailsResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o ApplicationClientDetailsResponseOutput) Oid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationClientDetailsResponse) *string { return v.Oid }).(pulumi.StringPtrOutput)
}

func (o ApplicationClientDetailsResponseOutput) Puid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationClientDetailsResponse) *string { return v.Puid }).(pulumi.StringPtrOutput)
}

type ApplicationDefinitionArtifact struct {
	Name string                  `pulumi:"name"`
	Type ApplicationArtifactType `pulumi:"type"`
	Uri  string                  `pulumi:"uri"`
}





type ApplicationDefinitionArtifactInput interface {
	pulumi.Input

	ToApplicationDefinitionArtifactOutput() ApplicationDefinitionArtifactOutput
	ToApplicationDefinitionArtifactOutputWithContext(context.Context) ApplicationDefinitionArtifactOutput
}

type ApplicationDefinitionArtifactArgs struct {
	Name pulumi.StringInput           `pulumi:"name"`
	Type ApplicationArtifactTypeInput `pulumi:"type"`
	Uri  pulumi.StringInput           `pulumi:"uri"`
}

func (ApplicationDefinitionArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDefinitionArtifact)(nil)).Elem()
}

func (i ApplicationDefinitionArtifactArgs) ToApplicationDefinitionArtifactOutput() ApplicationDefinitionArtifactOutput {
	return i.ToApplicationDefinitionArtifactOutputWithContext(context.Background())
}

func (i ApplicationDefinitionArtifactArgs) ToApplicationDefinitionArtifactOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDefinitionArtifactOutput)
}





type ApplicationDefinitionArtifactArrayInput interface {
	pulumi.Input

	ToApplicationDefinitionArtifactArrayOutput() ApplicationDefinitionArtifactArrayOutput
	ToApplicationDefinitionArtifactArrayOutputWithContext(context.Context) ApplicationDefinitionArtifactArrayOutput
}

type ApplicationDefinitionArtifactArray []ApplicationDefinitionArtifactInput

func (ApplicationDefinitionArtifactArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationDefinitionArtifact)(nil)).Elem()
}

func (i ApplicationDefinitionArtifactArray) ToApplicationDefinitionArtifactArrayOutput() ApplicationDefinitionArtifactArrayOutput {
	return i.ToApplicationDefinitionArtifactArrayOutputWithContext(context.Background())
}

func (i ApplicationDefinitionArtifactArray) ToApplicationDefinitionArtifactArrayOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationDefinitionArtifactArrayOutput)
}

type ApplicationDefinitionArtifactOutput struct{ *pulumi.OutputState }

func (ApplicationDefinitionArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDefinitionArtifact)(nil)).Elem()
}

func (o ApplicationDefinitionArtifactOutput) ToApplicationDefinitionArtifactOutput() ApplicationDefinitionArtifactOutput {
	return o
}

func (o ApplicationDefinitionArtifactOutput) ToApplicationDefinitionArtifactOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactOutput {
	return o
}

func (o ApplicationDefinitionArtifactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifact) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationDefinitionArtifactOutput) Type() ApplicationArtifactTypeOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifact) ApplicationArtifactType { return v.Type }).(ApplicationArtifactTypeOutput)
}

func (o ApplicationDefinitionArtifactOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifact) string { return v.Uri }).(pulumi.StringOutput)
}

type ApplicationDefinitionArtifactArrayOutput struct{ *pulumi.OutputState }

func (ApplicationDefinitionArtifactArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationDefinitionArtifact)(nil)).Elem()
}

func (o ApplicationDefinitionArtifactArrayOutput) ToApplicationDefinitionArtifactArrayOutput() ApplicationDefinitionArtifactArrayOutput {
	return o
}

func (o ApplicationDefinitionArtifactArrayOutput) ToApplicationDefinitionArtifactArrayOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactArrayOutput {
	return o
}

func (o ApplicationDefinitionArtifactArrayOutput) Index(i pulumi.IntInput) ApplicationDefinitionArtifactOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationDefinitionArtifact {
		return vs[0].([]ApplicationDefinitionArtifact)[vs[1].(int)]
	}).(ApplicationDefinitionArtifactOutput)
}

type ApplicationDefinitionArtifactResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
	Uri  string `pulumi:"uri"`
}

type ApplicationDefinitionArtifactResponseOutput struct{ *pulumi.OutputState }

func (ApplicationDefinitionArtifactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationDefinitionArtifactResponse)(nil)).Elem()
}

func (o ApplicationDefinitionArtifactResponseOutput) ToApplicationDefinitionArtifactResponseOutput() ApplicationDefinitionArtifactResponseOutput {
	return o
}

func (o ApplicationDefinitionArtifactResponseOutput) ToApplicationDefinitionArtifactResponseOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactResponseOutput {
	return o
}

func (o ApplicationDefinitionArtifactResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifactResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationDefinitionArtifactResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifactResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ApplicationDefinitionArtifactResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationDefinitionArtifactResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type ApplicationDefinitionArtifactResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationDefinitionArtifactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationDefinitionArtifactResponse)(nil)).Elem()
}

func (o ApplicationDefinitionArtifactResponseArrayOutput) ToApplicationDefinitionArtifactResponseArrayOutput() ApplicationDefinitionArtifactResponseArrayOutput {
	return o
}

func (o ApplicationDefinitionArtifactResponseArrayOutput) ToApplicationDefinitionArtifactResponseArrayOutputWithContext(ctx context.Context) ApplicationDefinitionArtifactResponseArrayOutput {
	return o
}

func (o ApplicationDefinitionArtifactResponseArrayOutput) Index(i pulumi.IntInput) ApplicationDefinitionArtifactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationDefinitionArtifactResponse {
		return vs[0].([]ApplicationDefinitionArtifactResponse)[vs[1].(int)]
	}).(ApplicationDefinitionArtifactResponseOutput)
}

type ApplicationJitAccessPolicy struct {
	JitAccessEnabled         bool                    `pulumi:"jitAccessEnabled"`
	JitApprovalMode          *string                 `pulumi:"jitApprovalMode"`
	JitApprovers             []JitApproverDefinition `pulumi:"jitApprovers"`
	MaximumJitAccessDuration *string                 `pulumi:"maximumJitAccessDuration"`
}





type ApplicationJitAccessPolicyInput interface {
	pulumi.Input

	ToApplicationJitAccessPolicyOutput() ApplicationJitAccessPolicyOutput
	ToApplicationJitAccessPolicyOutputWithContext(context.Context) ApplicationJitAccessPolicyOutput
}

type ApplicationJitAccessPolicyArgs struct {
	JitAccessEnabled         pulumi.BoolInput                `pulumi:"jitAccessEnabled"`
	JitApprovalMode          pulumi.StringPtrInput           `pulumi:"jitApprovalMode"`
	JitApprovers             JitApproverDefinitionArrayInput `pulumi:"jitApprovers"`
	MaximumJitAccessDuration pulumi.StringPtrInput           `pulumi:"maximumJitAccessDuration"`
}

func (ApplicationJitAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationJitAccessPolicy)(nil)).Elem()
}

func (i ApplicationJitAccessPolicyArgs) ToApplicationJitAccessPolicyOutput() ApplicationJitAccessPolicyOutput {
	return i.ToApplicationJitAccessPolicyOutputWithContext(context.Background())
}

func (i ApplicationJitAccessPolicyArgs) ToApplicationJitAccessPolicyOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationJitAccessPolicyOutput)
}

func (i ApplicationJitAccessPolicyArgs) ToApplicationJitAccessPolicyPtrOutput() ApplicationJitAccessPolicyPtrOutput {
	return i.ToApplicationJitAccessPolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationJitAccessPolicyArgs) ToApplicationJitAccessPolicyPtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationJitAccessPolicyOutput).ToApplicationJitAccessPolicyPtrOutputWithContext(ctx)
}









type ApplicationJitAccessPolicyPtrInput interface {
	pulumi.Input

	ToApplicationJitAccessPolicyPtrOutput() ApplicationJitAccessPolicyPtrOutput
	ToApplicationJitAccessPolicyPtrOutputWithContext(context.Context) ApplicationJitAccessPolicyPtrOutput
}

type applicationJitAccessPolicyPtrType ApplicationJitAccessPolicyArgs

func ApplicationJitAccessPolicyPtr(v *ApplicationJitAccessPolicyArgs) ApplicationJitAccessPolicyPtrInput {
	return (*applicationJitAccessPolicyPtrType)(v)
}

func (*applicationJitAccessPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationJitAccessPolicy)(nil)).Elem()
}

func (i *applicationJitAccessPolicyPtrType) ToApplicationJitAccessPolicyPtrOutput() ApplicationJitAccessPolicyPtrOutput {
	return i.ToApplicationJitAccessPolicyPtrOutputWithContext(context.Background())
}

func (i *applicationJitAccessPolicyPtrType) ToApplicationJitAccessPolicyPtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationJitAccessPolicyPtrOutput)
}

type ApplicationJitAccessPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationJitAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationJitAccessPolicy)(nil)).Elem()
}

func (o ApplicationJitAccessPolicyOutput) ToApplicationJitAccessPolicyOutput() ApplicationJitAccessPolicyOutput {
	return o
}

func (o ApplicationJitAccessPolicyOutput) ToApplicationJitAccessPolicyOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyOutput {
	return o
}

func (o ApplicationJitAccessPolicyOutput) ToApplicationJitAccessPolicyPtrOutput() ApplicationJitAccessPolicyPtrOutput {
	return o.ToApplicationJitAccessPolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationJitAccessPolicyOutput) ToApplicationJitAccessPolicyPtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationJitAccessPolicy) *ApplicationJitAccessPolicy {
		return &v
	}).(ApplicationJitAccessPolicyPtrOutput)
}

func (o ApplicationJitAccessPolicyOutput) JitAccessEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicy) bool { return v.JitAccessEnabled }).(pulumi.BoolOutput)
}

func (o ApplicationJitAccessPolicyOutput) JitApprovalMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicy) *string { return v.JitApprovalMode }).(pulumi.StringPtrOutput)
}

func (o ApplicationJitAccessPolicyOutput) JitApprovers() JitApproverDefinitionArrayOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicy) []JitApproverDefinition { return v.JitApprovers }).(JitApproverDefinitionArrayOutput)
}

func (o ApplicationJitAccessPolicyOutput) MaximumJitAccessDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicy) *string { return v.MaximumJitAccessDuration }).(pulumi.StringPtrOutput)
}

type ApplicationJitAccessPolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationJitAccessPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationJitAccessPolicy)(nil)).Elem()
}

func (o ApplicationJitAccessPolicyPtrOutput) ToApplicationJitAccessPolicyPtrOutput() ApplicationJitAccessPolicyPtrOutput {
	return o
}

func (o ApplicationJitAccessPolicyPtrOutput) ToApplicationJitAccessPolicyPtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyPtrOutput {
	return o
}

func (o ApplicationJitAccessPolicyPtrOutput) Elem() ApplicationJitAccessPolicyOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicy) ApplicationJitAccessPolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationJitAccessPolicy
		return ret
	}).(ApplicationJitAccessPolicyOutput)
}

func (o ApplicationJitAccessPolicyPtrOutput) JitAccessEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicy) *bool {
		if v == nil {
			return nil
		}
		return &v.JitAccessEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationJitAccessPolicyPtrOutput) JitApprovalMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicy) *string {
		if v == nil {
			return nil
		}
		return v.JitApprovalMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationJitAccessPolicyPtrOutput) JitApprovers() JitApproverDefinitionArrayOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicy) []JitApproverDefinition {
		if v == nil {
			return nil
		}
		return v.JitApprovers
	}).(JitApproverDefinitionArrayOutput)
}

func (o ApplicationJitAccessPolicyPtrOutput) MaximumJitAccessDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicy) *string {
		if v == nil {
			return nil
		}
		return v.MaximumJitAccessDuration
	}).(pulumi.StringPtrOutput)
}

type ApplicationJitAccessPolicyResponse struct {
	JitAccessEnabled         bool                            `pulumi:"jitAccessEnabled"`
	JitApprovalMode          *string                         `pulumi:"jitApprovalMode"`
	JitApprovers             []JitApproverDefinitionResponse `pulumi:"jitApprovers"`
	MaximumJitAccessDuration *string                         `pulumi:"maximumJitAccessDuration"`
}

type ApplicationJitAccessPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationJitAccessPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationJitAccessPolicyResponse)(nil)).Elem()
}

func (o ApplicationJitAccessPolicyResponseOutput) ToApplicationJitAccessPolicyResponseOutput() ApplicationJitAccessPolicyResponseOutput {
	return o
}

func (o ApplicationJitAccessPolicyResponseOutput) ToApplicationJitAccessPolicyResponseOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyResponseOutput {
	return o
}

func (o ApplicationJitAccessPolicyResponseOutput) JitAccessEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicyResponse) bool { return v.JitAccessEnabled }).(pulumi.BoolOutput)
}

func (o ApplicationJitAccessPolicyResponseOutput) JitApprovalMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicyResponse) *string { return v.JitApprovalMode }).(pulumi.StringPtrOutput)
}

func (o ApplicationJitAccessPolicyResponseOutput) JitApprovers() JitApproverDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicyResponse) []JitApproverDefinitionResponse { return v.JitApprovers }).(JitApproverDefinitionResponseArrayOutput)
}

func (o ApplicationJitAccessPolicyResponseOutput) MaximumJitAccessDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationJitAccessPolicyResponse) *string { return v.MaximumJitAccessDuration }).(pulumi.StringPtrOutput)
}

type ApplicationJitAccessPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationJitAccessPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationJitAccessPolicyResponse)(nil)).Elem()
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) ToApplicationJitAccessPolicyResponsePtrOutput() ApplicationJitAccessPolicyResponsePtrOutput {
	return o
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) ToApplicationJitAccessPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationJitAccessPolicyResponsePtrOutput {
	return o
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) Elem() ApplicationJitAccessPolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicyResponse) ApplicationJitAccessPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationJitAccessPolicyResponse
		return ret
	}).(ApplicationJitAccessPolicyResponseOutput)
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) JitAccessEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.JitAccessEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) JitApprovalMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.JitApprovalMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) JitApprovers() JitApproverDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicyResponse) []JitApproverDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.JitApprovers
	}).(JitApproverDefinitionResponseArrayOutput)
}

func (o ApplicationJitAccessPolicyResponsePtrOutput) MaximumJitAccessDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationJitAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaximumJitAccessDuration
	}).(pulumi.StringPtrOutput)
}

type ApplicationPackageContactResponse struct {
	ContactName *string `pulumi:"contactName"`
	Email       string  `pulumi:"email"`
	Phone       string  `pulumi:"phone"`
}

type ApplicationPackageContactResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPackageContactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageContactResponse)(nil)).Elem()
}

func (o ApplicationPackageContactResponseOutput) ToApplicationPackageContactResponseOutput() ApplicationPackageContactResponseOutput {
	return o
}

func (o ApplicationPackageContactResponseOutput) ToApplicationPackageContactResponseOutputWithContext(ctx context.Context) ApplicationPackageContactResponseOutput {
	return o
}

func (o ApplicationPackageContactResponseOutput) ContactName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageContactResponse) *string { return v.ContactName }).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageContactResponseOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageContactResponse) string { return v.Email }).(pulumi.StringOutput)
}

func (o ApplicationPackageContactResponseOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageContactResponse) string { return v.Phone }).(pulumi.StringOutput)
}

type ApplicationPackageSupportUrlsResponse struct {
	GovernmentCloud *string `pulumi:"governmentCloud"`
	PublicAzure     *string `pulumi:"publicAzure"`
}

type ApplicationPackageSupportUrlsResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPackageSupportUrlsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageSupportUrlsResponse)(nil)).Elem()
}

func (o ApplicationPackageSupportUrlsResponseOutput) ToApplicationPackageSupportUrlsResponseOutput() ApplicationPackageSupportUrlsResponseOutput {
	return o
}

func (o ApplicationPackageSupportUrlsResponseOutput) ToApplicationPackageSupportUrlsResponseOutputWithContext(ctx context.Context) ApplicationPackageSupportUrlsResponseOutput {
	return o
}

func (o ApplicationPackageSupportUrlsResponseOutput) GovernmentCloud() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageSupportUrlsResponse) *string { return v.GovernmentCloud }).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageSupportUrlsResponseOutput) PublicAzure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageSupportUrlsResponse) *string { return v.PublicAzure }).(pulumi.StringPtrOutput)
}

type ApplicationPolicy struct {
	Name               *string `pulumi:"name"`
	Parameters         *string `pulumi:"parameters"`
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
}





type ApplicationPolicyInput interface {
	pulumi.Input

	ToApplicationPolicyOutput() ApplicationPolicyOutput
	ToApplicationPolicyOutputWithContext(context.Context) ApplicationPolicyOutput
}

type ApplicationPolicyArgs struct {
	Name               pulumi.StringPtrInput `pulumi:"name"`
	Parameters         pulumi.StringPtrInput `pulumi:"parameters"`
	PolicyDefinitionId pulumi.StringPtrInput `pulumi:"policyDefinitionId"`
}

func (ApplicationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPolicy)(nil)).Elem()
}

func (i ApplicationPolicyArgs) ToApplicationPolicyOutput() ApplicationPolicyOutput {
	return i.ToApplicationPolicyOutputWithContext(context.Background())
}

func (i ApplicationPolicyArgs) ToApplicationPolicyOutputWithContext(ctx context.Context) ApplicationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPolicyOutput)
}





type ApplicationPolicyArrayInput interface {
	pulumi.Input

	ToApplicationPolicyArrayOutput() ApplicationPolicyArrayOutput
	ToApplicationPolicyArrayOutputWithContext(context.Context) ApplicationPolicyArrayOutput
}

type ApplicationPolicyArray []ApplicationPolicyInput

func (ApplicationPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPolicy)(nil)).Elem()
}

func (i ApplicationPolicyArray) ToApplicationPolicyArrayOutput() ApplicationPolicyArrayOutput {
	return i.ToApplicationPolicyArrayOutputWithContext(context.Background())
}

func (i ApplicationPolicyArray) ToApplicationPolicyArrayOutputWithContext(ctx context.Context) ApplicationPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPolicyArrayOutput)
}

type ApplicationPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPolicy)(nil)).Elem()
}

func (o ApplicationPolicyOutput) ToApplicationPolicyOutput() ApplicationPolicyOutput {
	return o
}

func (o ApplicationPolicyOutput) ToApplicationPolicyOutputWithContext(ctx context.Context) ApplicationPolicyOutput {
	return o
}

func (o ApplicationPolicyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicy) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationPolicyOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicy) *string { return v.Parameters }).(pulumi.StringPtrOutput)
}

func (o ApplicationPolicyOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicy) *string { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

type ApplicationPolicyArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPolicy)(nil)).Elem()
}

func (o ApplicationPolicyArrayOutput) ToApplicationPolicyArrayOutput() ApplicationPolicyArrayOutput {
	return o
}

func (o ApplicationPolicyArrayOutput) ToApplicationPolicyArrayOutputWithContext(ctx context.Context) ApplicationPolicyArrayOutput {
	return o
}

func (o ApplicationPolicyArrayOutput) Index(i pulumi.IntInput) ApplicationPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationPolicy {
		return vs[0].([]ApplicationPolicy)[vs[1].(int)]
	}).(ApplicationPolicyOutput)
}

type ApplicationPolicyResponse struct {
	Name               *string `pulumi:"name"`
	Parameters         *string `pulumi:"parameters"`
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
}

type ApplicationPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPolicyResponse)(nil)).Elem()
}

func (o ApplicationPolicyResponseOutput) ToApplicationPolicyResponseOutput() ApplicationPolicyResponseOutput {
	return o
}

func (o ApplicationPolicyResponseOutput) ToApplicationPolicyResponseOutputWithContext(ctx context.Context) ApplicationPolicyResponseOutput {
	return o
}

func (o ApplicationPolicyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationPolicyResponseOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicyResponse) *string { return v.Parameters }).(pulumi.StringPtrOutput)
}

func (o ApplicationPolicyResponseOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPolicyResponse) *string { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

type ApplicationPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPolicyResponse)(nil)).Elem()
}

func (o ApplicationPolicyResponseArrayOutput) ToApplicationPolicyResponseArrayOutput() ApplicationPolicyResponseArrayOutput {
	return o
}

func (o ApplicationPolicyResponseArrayOutput) ToApplicationPolicyResponseArrayOutputWithContext(ctx context.Context) ApplicationPolicyResponseArrayOutput {
	return o
}

func (o ApplicationPolicyResponseArrayOutput) Index(i pulumi.IntInput) ApplicationPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationPolicyResponse {
		return vs[0].([]ApplicationPolicyResponse)[vs[1].(int)]
	}).(ApplicationPolicyResponseOutput)
}

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
	PrincipalId            string                                          `pulumi:"principalId"`
	TenantId               string                                          `pulumi:"tenantId"`
	Type                   *string                                         `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedResourceIdentityResponse `pulumi:"userAssignedIdentities"`
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

func (o IdentityResponseOutput) UserAssignedIdentities() UserAssignedResourceIdentityResponseMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]UserAssignedResourceIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedResourceIdentityResponseMapOutput)
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

func (o IdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedResourceIdentityResponseMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]UserAssignedResourceIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedResourceIdentityResponseMapOutput)
}

type JitApproverDefinition struct {
	DisplayName *string `pulumi:"displayName"`
	Id          string  `pulumi:"id"`
	Type        *string `pulumi:"type"`
}





type JitApproverDefinitionInput interface {
	pulumi.Input

	ToJitApproverDefinitionOutput() JitApproverDefinitionOutput
	ToJitApproverDefinitionOutputWithContext(context.Context) JitApproverDefinitionOutput
}

type JitApproverDefinitionArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Id          pulumi.StringInput    `pulumi:"id"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (JitApproverDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitApproverDefinition)(nil)).Elem()
}

func (i JitApproverDefinitionArgs) ToJitApproverDefinitionOutput() JitApproverDefinitionOutput {
	return i.ToJitApproverDefinitionOutputWithContext(context.Background())
}

func (i JitApproverDefinitionArgs) ToJitApproverDefinitionOutputWithContext(ctx context.Context) JitApproverDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitApproverDefinitionOutput)
}





type JitApproverDefinitionArrayInput interface {
	pulumi.Input

	ToJitApproverDefinitionArrayOutput() JitApproverDefinitionArrayOutput
	ToJitApproverDefinitionArrayOutputWithContext(context.Context) JitApproverDefinitionArrayOutput
}

type JitApproverDefinitionArray []JitApproverDefinitionInput

func (JitApproverDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitApproverDefinition)(nil)).Elem()
}

func (i JitApproverDefinitionArray) ToJitApproverDefinitionArrayOutput() JitApproverDefinitionArrayOutput {
	return i.ToJitApproverDefinitionArrayOutputWithContext(context.Background())
}

func (i JitApproverDefinitionArray) ToJitApproverDefinitionArrayOutputWithContext(ctx context.Context) JitApproverDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitApproverDefinitionArrayOutput)
}

type JitApproverDefinitionOutput struct{ *pulumi.OutputState }

func (JitApproverDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitApproverDefinition)(nil)).Elem()
}

func (o JitApproverDefinitionOutput) ToJitApproverDefinitionOutput() JitApproverDefinitionOutput {
	return o
}

func (o JitApproverDefinitionOutput) ToJitApproverDefinitionOutputWithContext(ctx context.Context) JitApproverDefinitionOutput {
	return o
}

func (o JitApproverDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitApproverDefinition) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o JitApproverDefinitionOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitApproverDefinition) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitApproverDefinitionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitApproverDefinition) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JitApproverDefinitionArrayOutput struct{ *pulumi.OutputState }

func (JitApproverDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitApproverDefinition)(nil)).Elem()
}

func (o JitApproverDefinitionArrayOutput) ToJitApproverDefinitionArrayOutput() JitApproverDefinitionArrayOutput {
	return o
}

func (o JitApproverDefinitionArrayOutput) ToJitApproverDefinitionArrayOutputWithContext(ctx context.Context) JitApproverDefinitionArrayOutput {
	return o
}

func (o JitApproverDefinitionArrayOutput) Index(i pulumi.IntInput) JitApproverDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitApproverDefinition {
		return vs[0].([]JitApproverDefinition)[vs[1].(int)]
	}).(JitApproverDefinitionOutput)
}

type JitApproverDefinitionResponse struct {
	DisplayName *string `pulumi:"displayName"`
	Id          string  `pulumi:"id"`
	Type        *string `pulumi:"type"`
}

type JitApproverDefinitionResponseOutput struct{ *pulumi.OutputState }

func (JitApproverDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitApproverDefinitionResponse)(nil)).Elem()
}

func (o JitApproverDefinitionResponseOutput) ToJitApproverDefinitionResponseOutput() JitApproverDefinitionResponseOutput {
	return o
}

func (o JitApproverDefinitionResponseOutput) ToJitApproverDefinitionResponseOutputWithContext(ctx context.Context) JitApproverDefinitionResponseOutput {
	return o
}

func (o JitApproverDefinitionResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitApproverDefinitionResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o JitApproverDefinitionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitApproverDefinitionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitApproverDefinitionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitApproverDefinitionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JitApproverDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (JitApproverDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitApproverDefinitionResponse)(nil)).Elem()
}

func (o JitApproverDefinitionResponseArrayOutput) ToJitApproverDefinitionResponseArrayOutput() JitApproverDefinitionResponseArrayOutput {
	return o
}

func (o JitApproverDefinitionResponseArrayOutput) ToJitApproverDefinitionResponseArrayOutputWithContext(ctx context.Context) JitApproverDefinitionResponseArrayOutput {
	return o
}

func (o JitApproverDefinitionResponseArrayOutput) Index(i pulumi.IntInput) JitApproverDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitApproverDefinitionResponse {
		return vs[0].([]JitApproverDefinitionResponse)[vs[1].(int)]
	}).(JitApproverDefinitionResponseOutput)
}

type JitAuthorizationPolicies struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}





type JitAuthorizationPoliciesInput interface {
	pulumi.Input

	ToJitAuthorizationPoliciesOutput() JitAuthorizationPoliciesOutput
	ToJitAuthorizationPoliciesOutputWithContext(context.Context) JitAuthorizationPoliciesOutput
}

type JitAuthorizationPoliciesArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (JitAuthorizationPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitAuthorizationPolicies)(nil)).Elem()
}

func (i JitAuthorizationPoliciesArgs) ToJitAuthorizationPoliciesOutput() JitAuthorizationPoliciesOutput {
	return i.ToJitAuthorizationPoliciesOutputWithContext(context.Background())
}

func (i JitAuthorizationPoliciesArgs) ToJitAuthorizationPoliciesOutputWithContext(ctx context.Context) JitAuthorizationPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitAuthorizationPoliciesOutput)
}





type JitAuthorizationPoliciesArrayInput interface {
	pulumi.Input

	ToJitAuthorizationPoliciesArrayOutput() JitAuthorizationPoliciesArrayOutput
	ToJitAuthorizationPoliciesArrayOutputWithContext(context.Context) JitAuthorizationPoliciesArrayOutput
}

type JitAuthorizationPoliciesArray []JitAuthorizationPoliciesInput

func (JitAuthorizationPoliciesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitAuthorizationPolicies)(nil)).Elem()
}

func (i JitAuthorizationPoliciesArray) ToJitAuthorizationPoliciesArrayOutput() JitAuthorizationPoliciesArrayOutput {
	return i.ToJitAuthorizationPoliciesArrayOutputWithContext(context.Background())
}

func (i JitAuthorizationPoliciesArray) ToJitAuthorizationPoliciesArrayOutputWithContext(ctx context.Context) JitAuthorizationPoliciesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitAuthorizationPoliciesArrayOutput)
}

type JitAuthorizationPoliciesOutput struct{ *pulumi.OutputState }

func (JitAuthorizationPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitAuthorizationPolicies)(nil)).Elem()
}

func (o JitAuthorizationPoliciesOutput) ToJitAuthorizationPoliciesOutput() JitAuthorizationPoliciesOutput {
	return o
}

func (o JitAuthorizationPoliciesOutput) ToJitAuthorizationPoliciesOutputWithContext(ctx context.Context) JitAuthorizationPoliciesOutput {
	return o
}

func (o JitAuthorizationPoliciesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v JitAuthorizationPolicies) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o JitAuthorizationPoliciesOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v JitAuthorizationPolicies) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type JitAuthorizationPoliciesArrayOutput struct{ *pulumi.OutputState }

func (JitAuthorizationPoliciesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitAuthorizationPolicies)(nil)).Elem()
}

func (o JitAuthorizationPoliciesArrayOutput) ToJitAuthorizationPoliciesArrayOutput() JitAuthorizationPoliciesArrayOutput {
	return o
}

func (o JitAuthorizationPoliciesArrayOutput) ToJitAuthorizationPoliciesArrayOutputWithContext(ctx context.Context) JitAuthorizationPoliciesArrayOutput {
	return o
}

func (o JitAuthorizationPoliciesArrayOutput) Index(i pulumi.IntInput) JitAuthorizationPoliciesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitAuthorizationPolicies {
		return vs[0].([]JitAuthorizationPolicies)[vs[1].(int)]
	}).(JitAuthorizationPoliciesOutput)
}

type JitAuthorizationPoliciesResponse struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}

type JitAuthorizationPoliciesResponseOutput struct{ *pulumi.OutputState }

func (JitAuthorizationPoliciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitAuthorizationPoliciesResponse)(nil)).Elem()
}

func (o JitAuthorizationPoliciesResponseOutput) ToJitAuthorizationPoliciesResponseOutput() JitAuthorizationPoliciesResponseOutput {
	return o
}

func (o JitAuthorizationPoliciesResponseOutput) ToJitAuthorizationPoliciesResponseOutputWithContext(ctx context.Context) JitAuthorizationPoliciesResponseOutput {
	return o
}

func (o JitAuthorizationPoliciesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v JitAuthorizationPoliciesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o JitAuthorizationPoliciesResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v JitAuthorizationPoliciesResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type JitAuthorizationPoliciesResponseArrayOutput struct{ *pulumi.OutputState }

func (JitAuthorizationPoliciesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitAuthorizationPoliciesResponse)(nil)).Elem()
}

func (o JitAuthorizationPoliciesResponseArrayOutput) ToJitAuthorizationPoliciesResponseArrayOutput() JitAuthorizationPoliciesResponseArrayOutput {
	return o
}

func (o JitAuthorizationPoliciesResponseArrayOutput) ToJitAuthorizationPoliciesResponseArrayOutputWithContext(ctx context.Context) JitAuthorizationPoliciesResponseArrayOutput {
	return o
}

func (o JitAuthorizationPoliciesResponseArrayOutput) Index(i pulumi.IntInput) JitAuthorizationPoliciesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitAuthorizationPoliciesResponse {
		return vs[0].([]JitAuthorizationPoliciesResponse)[vs[1].(int)]
	}).(JitAuthorizationPoliciesResponseOutput)
}

type JitSchedulingPolicy struct {
	Duration  string `pulumi:"duration"`
	StartTime string `pulumi:"startTime"`
	Type      string `pulumi:"type"`
}





type JitSchedulingPolicyInput interface {
	pulumi.Input

	ToJitSchedulingPolicyOutput() JitSchedulingPolicyOutput
	ToJitSchedulingPolicyOutputWithContext(context.Context) JitSchedulingPolicyOutput
}

type JitSchedulingPolicyArgs struct {
	Duration  pulumi.StringInput `pulumi:"duration"`
	StartTime pulumi.StringInput `pulumi:"startTime"`
	Type      pulumi.StringInput `pulumi:"type"`
}

func (JitSchedulingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitSchedulingPolicy)(nil)).Elem()
}

func (i JitSchedulingPolicyArgs) ToJitSchedulingPolicyOutput() JitSchedulingPolicyOutput {
	return i.ToJitSchedulingPolicyOutputWithContext(context.Background())
}

func (i JitSchedulingPolicyArgs) ToJitSchedulingPolicyOutputWithContext(ctx context.Context) JitSchedulingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitSchedulingPolicyOutput)
}

type JitSchedulingPolicyOutput struct{ *pulumi.OutputState }

func (JitSchedulingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitSchedulingPolicy)(nil)).Elem()
}

func (o JitSchedulingPolicyOutput) ToJitSchedulingPolicyOutput() JitSchedulingPolicyOutput {
	return o
}

func (o JitSchedulingPolicyOutput) ToJitSchedulingPolicyOutputWithContext(ctx context.Context) JitSchedulingPolicyOutput {
	return o
}

func (o JitSchedulingPolicyOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v JitSchedulingPolicy) string { return v.Duration }).(pulumi.StringOutput)
}

func (o JitSchedulingPolicyOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v JitSchedulingPolicy) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o JitSchedulingPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JitSchedulingPolicy) string { return v.Type }).(pulumi.StringOutput)
}

type JitSchedulingPolicyResponse struct {
	Duration  string `pulumi:"duration"`
	StartTime string `pulumi:"startTime"`
	Type      string `pulumi:"type"`
}

type JitSchedulingPolicyResponseOutput struct{ *pulumi.OutputState }

func (JitSchedulingPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitSchedulingPolicyResponse)(nil)).Elem()
}

func (o JitSchedulingPolicyResponseOutput) ToJitSchedulingPolicyResponseOutput() JitSchedulingPolicyResponseOutput {
	return o
}

func (o JitSchedulingPolicyResponseOutput) ToJitSchedulingPolicyResponseOutputWithContext(ctx context.Context) JitSchedulingPolicyResponseOutput {
	return o
}

func (o JitSchedulingPolicyResponseOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v JitSchedulingPolicyResponse) string { return v.Duration }).(pulumi.StringOutput)
}

func (o JitSchedulingPolicyResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v JitSchedulingPolicyResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o JitSchedulingPolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JitSchedulingPolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ManagedIdentityTokenResponse struct {
	AccessToken           *string `pulumi:"accessToken"`
	AuthorizationAudience *string `pulumi:"authorizationAudience"`
	ExpiresIn             *string `pulumi:"expiresIn"`
	ExpiresOn             *string `pulumi:"expiresOn"`
	NotBefore             *string `pulumi:"notBefore"`
	ResourceId            *string `pulumi:"resourceId"`
	TokenType             *string `pulumi:"tokenType"`
}

type ManagedIdentityTokenResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTokenResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityTokenResponse)(nil)).Elem()
}

func (o ManagedIdentityTokenResponseOutput) ToManagedIdentityTokenResponseOutput() ManagedIdentityTokenResponseOutput {
	return o
}

func (o ManagedIdentityTokenResponseOutput) ToManagedIdentityTokenResponseOutputWithContext(ctx context.Context) ManagedIdentityTokenResponseOutput {
	return o
}

func (o ManagedIdentityTokenResponseOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityTokenResponse) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityTokenResponseOutput) AuthorizationAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityTokenResponse) *string { return v.AuthorizationAudience }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityTokenResponseOutput) ExpiresIn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityTokenResponse) *string { return v.ExpiresIn }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityTokenResponseOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityTokenResponse) *string { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityTokenResponseOutput) NotBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityTokenResponse) *string { return v.NotBefore }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityTokenResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityTokenResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityTokenResponseOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityTokenResponse) *string { return v.TokenType }).(pulumi.StringPtrOutput)
}

type ManagedIdentityTokenResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTokenResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedIdentityTokenResponse)(nil)).Elem()
}

func (o ManagedIdentityTokenResponseArrayOutput) ToManagedIdentityTokenResponseArrayOutput() ManagedIdentityTokenResponseArrayOutput {
	return o
}

func (o ManagedIdentityTokenResponseArrayOutput) ToManagedIdentityTokenResponseArrayOutputWithContext(ctx context.Context) ManagedIdentityTokenResponseArrayOutput {
	return o
}

func (o ManagedIdentityTokenResponseArrayOutput) Index(i pulumi.IntInput) ManagedIdentityTokenResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedIdentityTokenResponse {
		return vs[0].([]ManagedIdentityTokenResponse)[vs[1].(int)]
	}).(ManagedIdentityTokenResponseOutput)
}

type Plan struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
	Version       string  `pulumi:"version"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Name          pulumi.StringInput    `pulumi:"name"`
	Product       pulumi.StringInput    `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringInput    `pulumi:"publisher"`
	Version       pulumi.StringInput    `pulumi:"version"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

func (i PlanArgs) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput).ToPlanPtrOutputWithContext(ctx)
}









type PlanPtrInput interface {
	pulumi.Input

	ToPlanPtrOutput() PlanPtrOutput
	ToPlanPtrOutputWithContext(context.Context) PlanPtrOutput
}

type planPtrType PlanArgs

func PlanPtr(v *PlanArgs) PlanPtrInput {
	return (*planPtrType)(v)
}

func (*planPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (i *planPtrType) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i *planPtrType) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanPtrOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o.ToPlanPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Plan) *Plan {
		return &v
	}).(PlanPtrOutput)
}

func (o PlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Name }).(pulumi.StringOutput)
}

func (o PlanOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Product }).(pulumi.StringOutput)
}

func (o PlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o PlanOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Version }).(pulumi.StringOutput)
}

type PlanPtrOutput struct{ *pulumi.OutputState }

func (PlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (o PlanPtrOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) Elem() PlanOutput {
	return o.ApplyT(func(v *Plan) Plan {
		if v != nil {
			return *v
		}
		var ret Plan
		return ret
	}).(PlanOutput)
}

func (o PlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type PlanResponse struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
	Version       string  `pulumi:"version"`
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Product }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Version }).(pulumi.StringOutput)
}

type PlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) Elem() PlanResponseOutput {
	return o.ApplyT(func(v *PlanResponse) PlanResponse {
		if v != nil {
			return *v
		}
		var ret PlanResponse
		return ret
	}).(PlanResponseOutput)
}

func (o PlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Model    *string `pulumi:"model"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Model    pulumi.StringPtrInput `pulumi:"model"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Model }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Model
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Model    *string `pulumi:"model"`
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

func (o SkuResponseOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Model }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Model
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type UserAssignedResourceIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
}

type UserAssignedResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedResourceIdentityResponse)(nil)).Elem()
}

func (o UserAssignedResourceIdentityResponseOutput) ToUserAssignedResourceIdentityResponseOutput() UserAssignedResourceIdentityResponseOutput {
	return o
}

func (o UserAssignedResourceIdentityResponseOutput) ToUserAssignedResourceIdentityResponseOutputWithContext(ctx context.Context) UserAssignedResourceIdentityResponseOutput {
	return o
}

func (o UserAssignedResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o UserAssignedResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type UserAssignedResourceIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedResourceIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedResourceIdentityResponse)(nil)).Elem()
}

func (o UserAssignedResourceIdentityResponseMapOutput) ToUserAssignedResourceIdentityResponseMapOutput() UserAssignedResourceIdentityResponseMapOutput {
	return o
}

func (o UserAssignedResourceIdentityResponseMapOutput) ToUserAssignedResourceIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedResourceIdentityResponseMapOutput {
	return o
}

func (o UserAssignedResourceIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedResourceIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedResourceIdentityResponse {
		return vs[0].(map[string]UserAssignedResourceIdentityResponse)[vs[1].(string)]
	}).(UserAssignedResourceIdentityResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationArtifactResponseOutput{})
	pulumi.RegisterOutputType(ApplicationArtifactResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationAuthorizationOutput{})
	pulumi.RegisterOutputType(ApplicationAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationAuthorizationResponseOutput{})
	pulumi.RegisterOutputType(ApplicationAuthorizationResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationBillingDetailsDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ApplicationClientDetailsResponseOutput{})
	pulumi.RegisterOutputType(ApplicationDefinitionArtifactOutput{})
	pulumi.RegisterOutputType(ApplicationDefinitionArtifactArrayOutput{})
	pulumi.RegisterOutputType(ApplicationDefinitionArtifactResponseOutput{})
	pulumi.RegisterOutputType(ApplicationDefinitionArtifactResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationJitAccessPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationJitAccessPolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationJitAccessPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationJitAccessPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationPackageContactResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPackageSupportUrlsResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationPolicyArrayOutput{})
	pulumi.RegisterOutputType(ApplicationPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(JitApproverDefinitionOutput{})
	pulumi.RegisterOutputType(JitApproverDefinitionArrayOutput{})
	pulumi.RegisterOutputType(JitApproverDefinitionResponseOutput{})
	pulumi.RegisterOutputType(JitApproverDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(JitAuthorizationPoliciesOutput{})
	pulumi.RegisterOutputType(JitAuthorizationPoliciesArrayOutput{})
	pulumi.RegisterOutputType(JitAuthorizationPoliciesResponseOutput{})
	pulumi.RegisterOutputType(JitAuthorizationPoliciesResponseArrayOutput{})
	pulumi.RegisterOutputType(JitSchedulingPolicyOutput{})
	pulumi.RegisterOutputType(JitSchedulingPolicyResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTokenResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTokenResponseArrayOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedResourceIdentityResponseMapOutput{})
}
