


package v20200201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Authorization struct {
	DelegatedRoleDefinitionIds []string `pulumi:"delegatedRoleDefinitionIds"`
	PrincipalId                string   `pulumi:"principalId"`
	PrincipalIdDisplayName     *string  `pulumi:"principalIdDisplayName"`
	RoleDefinitionId           string   `pulumi:"roleDefinitionId"`
}





type AuthorizationInput interface {
	pulumi.Input

	ToAuthorizationOutput() AuthorizationOutput
	ToAuthorizationOutputWithContext(context.Context) AuthorizationOutput
}

type AuthorizationArgs struct {
	DelegatedRoleDefinitionIds pulumi.StringArrayInput `pulumi:"delegatedRoleDefinitionIds"`
	PrincipalId                pulumi.StringInput      `pulumi:"principalId"`
	PrincipalIdDisplayName     pulumi.StringPtrInput   `pulumi:"principalIdDisplayName"`
	RoleDefinitionId           pulumi.StringInput      `pulumi:"roleDefinitionId"`
}

func (AuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Authorization)(nil)).Elem()
}

func (i AuthorizationArgs) ToAuthorizationOutput() AuthorizationOutput {
	return i.ToAuthorizationOutputWithContext(context.Background())
}

func (i AuthorizationArgs) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationOutput)
}





type AuthorizationArrayInput interface {
	pulumi.Input

	ToAuthorizationArrayOutput() AuthorizationArrayOutput
	ToAuthorizationArrayOutputWithContext(context.Context) AuthorizationArrayOutput
}

type AuthorizationArray []AuthorizationInput

func (AuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Authorization)(nil)).Elem()
}

func (i AuthorizationArray) ToAuthorizationArrayOutput() AuthorizationArrayOutput {
	return i.ToAuthorizationArrayOutputWithContext(context.Background())
}

func (i AuthorizationArray) ToAuthorizationArrayOutputWithContext(ctx context.Context) AuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationArrayOutput)
}

type AuthorizationOutput struct{ *pulumi.OutputState }

func (AuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Authorization)(nil)).Elem()
}

func (o AuthorizationOutput) ToAuthorizationOutput() AuthorizationOutput {
	return o
}

func (o AuthorizationOutput) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return o
}

func (o AuthorizationOutput) DelegatedRoleDefinitionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Authorization) []string { return v.DelegatedRoleDefinitionIds }).(pulumi.StringArrayOutput)
}

func (o AuthorizationOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v Authorization) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o AuthorizationOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Authorization) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

func (o AuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v Authorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type AuthorizationArrayOutput struct{ *pulumi.OutputState }

func (AuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Authorization)(nil)).Elem()
}

func (o AuthorizationArrayOutput) ToAuthorizationArrayOutput() AuthorizationArrayOutput {
	return o
}

func (o AuthorizationArrayOutput) ToAuthorizationArrayOutputWithContext(ctx context.Context) AuthorizationArrayOutput {
	return o
}

func (o AuthorizationArrayOutput) Index(i pulumi.IntInput) AuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Authorization {
		return vs[0].([]Authorization)[vs[1].(int)]
	}).(AuthorizationOutput)
}

type AuthorizationResponse struct {
	DelegatedRoleDefinitionIds []string `pulumi:"delegatedRoleDefinitionIds"`
	PrincipalId                string   `pulumi:"principalId"`
	PrincipalIdDisplayName     *string  `pulumi:"principalIdDisplayName"`
	RoleDefinitionId           string   `pulumi:"roleDefinitionId"`
}





type AuthorizationResponseInput interface {
	pulumi.Input

	ToAuthorizationResponseOutput() AuthorizationResponseOutput
	ToAuthorizationResponseOutputWithContext(context.Context) AuthorizationResponseOutput
}

type AuthorizationResponseArgs struct {
	DelegatedRoleDefinitionIds pulumi.StringArrayInput `pulumi:"delegatedRoleDefinitionIds"`
	PrincipalId                pulumi.StringInput      `pulumi:"principalId"`
	PrincipalIdDisplayName     pulumi.StringPtrInput   `pulumi:"principalIdDisplayName"`
	RoleDefinitionId           pulumi.StringInput      `pulumi:"roleDefinitionId"`
}

func (AuthorizationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationResponse)(nil)).Elem()
}

func (i AuthorizationResponseArgs) ToAuthorizationResponseOutput() AuthorizationResponseOutput {
	return i.ToAuthorizationResponseOutputWithContext(context.Background())
}

func (i AuthorizationResponseArgs) ToAuthorizationResponseOutputWithContext(ctx context.Context) AuthorizationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationResponseOutput)
}





type AuthorizationResponseArrayInput interface {
	pulumi.Input

	ToAuthorizationResponseArrayOutput() AuthorizationResponseArrayOutput
	ToAuthorizationResponseArrayOutputWithContext(context.Context) AuthorizationResponseArrayOutput
}

type AuthorizationResponseArray []AuthorizationResponseInput

func (AuthorizationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthorizationResponse)(nil)).Elem()
}

func (i AuthorizationResponseArray) ToAuthorizationResponseArrayOutput() AuthorizationResponseArrayOutput {
	return i.ToAuthorizationResponseArrayOutputWithContext(context.Background())
}

func (i AuthorizationResponseArray) ToAuthorizationResponseArrayOutputWithContext(ctx context.Context) AuthorizationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationResponseArrayOutput)
}

type AuthorizationResponseOutput struct{ *pulumi.OutputState }

func (AuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationResponse)(nil)).Elem()
}

func (o AuthorizationResponseOutput) ToAuthorizationResponseOutput() AuthorizationResponseOutput {
	return o
}

func (o AuthorizationResponseOutput) ToAuthorizationResponseOutputWithContext(ctx context.Context) AuthorizationResponseOutput {
	return o
}

func (o AuthorizationResponseOutput) DelegatedRoleDefinitionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthorizationResponse) []string { return v.DelegatedRoleDefinitionIds }).(pulumi.StringArrayOutput)
}

func (o AuthorizationResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o AuthorizationResponseOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthorizationResponse) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

func (o AuthorizationResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type AuthorizationResponseArrayOutput struct{ *pulumi.OutputState }

func (AuthorizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthorizationResponse)(nil)).Elem()
}

func (o AuthorizationResponseArrayOutput) ToAuthorizationResponseArrayOutput() AuthorizationResponseArrayOutput {
	return o
}

func (o AuthorizationResponseArrayOutput) ToAuthorizationResponseArrayOutputWithContext(ctx context.Context) AuthorizationResponseArrayOutput {
	return o
}

func (o AuthorizationResponseArrayOutput) Index(i pulumi.IntInput) AuthorizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthorizationResponse {
		return vs[0].([]AuthorizationResponse)[vs[1].(int)]
	}).(AuthorizationResponseOutput)
}

type EligibleApprover struct {
	PrincipalId            string  `pulumi:"principalId"`
	PrincipalIdDisplayName *string `pulumi:"principalIdDisplayName"`
}





type EligibleApproverInput interface {
	pulumi.Input

	ToEligibleApproverOutput() EligibleApproverOutput
	ToEligibleApproverOutputWithContext(context.Context) EligibleApproverOutput
}

type EligibleApproverArgs struct {
	PrincipalId            pulumi.StringInput    `pulumi:"principalId"`
	PrincipalIdDisplayName pulumi.StringPtrInput `pulumi:"principalIdDisplayName"`
}

func (EligibleApproverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleApprover)(nil)).Elem()
}

func (i EligibleApproverArgs) ToEligibleApproverOutput() EligibleApproverOutput {
	return i.ToEligibleApproverOutputWithContext(context.Background())
}

func (i EligibleApproverArgs) ToEligibleApproverOutputWithContext(ctx context.Context) EligibleApproverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleApproverOutput)
}





type EligibleApproverArrayInput interface {
	pulumi.Input

	ToEligibleApproverArrayOutput() EligibleApproverArrayOutput
	ToEligibleApproverArrayOutputWithContext(context.Context) EligibleApproverArrayOutput
}

type EligibleApproverArray []EligibleApproverInput

func (EligibleApproverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleApprover)(nil)).Elem()
}

func (i EligibleApproverArray) ToEligibleApproverArrayOutput() EligibleApproverArrayOutput {
	return i.ToEligibleApproverArrayOutputWithContext(context.Background())
}

func (i EligibleApproverArray) ToEligibleApproverArrayOutputWithContext(ctx context.Context) EligibleApproverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleApproverArrayOutput)
}

type EligibleApproverOutput struct{ *pulumi.OutputState }

func (EligibleApproverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleApprover)(nil)).Elem()
}

func (o EligibleApproverOutput) ToEligibleApproverOutput() EligibleApproverOutput {
	return o
}

func (o EligibleApproverOutput) ToEligibleApproverOutputWithContext(ctx context.Context) EligibleApproverOutput {
	return o
}

func (o EligibleApproverOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleApprover) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o EligibleApproverOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EligibleApprover) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

type EligibleApproverArrayOutput struct{ *pulumi.OutputState }

func (EligibleApproverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleApprover)(nil)).Elem()
}

func (o EligibleApproverArrayOutput) ToEligibleApproverArrayOutput() EligibleApproverArrayOutput {
	return o
}

func (o EligibleApproverArrayOutput) ToEligibleApproverArrayOutputWithContext(ctx context.Context) EligibleApproverArrayOutput {
	return o
}

func (o EligibleApproverArrayOutput) Index(i pulumi.IntInput) EligibleApproverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EligibleApprover {
		return vs[0].([]EligibleApprover)[vs[1].(int)]
	}).(EligibleApproverOutput)
}

type EligibleApproverResponse struct {
	PrincipalId            string  `pulumi:"principalId"`
	PrincipalIdDisplayName *string `pulumi:"principalIdDisplayName"`
}





type EligibleApproverResponseInput interface {
	pulumi.Input

	ToEligibleApproverResponseOutput() EligibleApproverResponseOutput
	ToEligibleApproverResponseOutputWithContext(context.Context) EligibleApproverResponseOutput
}

type EligibleApproverResponseArgs struct {
	PrincipalId            pulumi.StringInput    `pulumi:"principalId"`
	PrincipalIdDisplayName pulumi.StringPtrInput `pulumi:"principalIdDisplayName"`
}

func (EligibleApproverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleApproverResponse)(nil)).Elem()
}

func (i EligibleApproverResponseArgs) ToEligibleApproverResponseOutput() EligibleApproverResponseOutput {
	return i.ToEligibleApproverResponseOutputWithContext(context.Background())
}

func (i EligibleApproverResponseArgs) ToEligibleApproverResponseOutputWithContext(ctx context.Context) EligibleApproverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleApproverResponseOutput)
}





type EligibleApproverResponseArrayInput interface {
	pulumi.Input

	ToEligibleApproverResponseArrayOutput() EligibleApproverResponseArrayOutput
	ToEligibleApproverResponseArrayOutputWithContext(context.Context) EligibleApproverResponseArrayOutput
}

type EligibleApproverResponseArray []EligibleApproverResponseInput

func (EligibleApproverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleApproverResponse)(nil)).Elem()
}

func (i EligibleApproverResponseArray) ToEligibleApproverResponseArrayOutput() EligibleApproverResponseArrayOutput {
	return i.ToEligibleApproverResponseArrayOutputWithContext(context.Background())
}

func (i EligibleApproverResponseArray) ToEligibleApproverResponseArrayOutputWithContext(ctx context.Context) EligibleApproverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleApproverResponseArrayOutput)
}

type EligibleApproverResponseOutput struct{ *pulumi.OutputState }

func (EligibleApproverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleApproverResponse)(nil)).Elem()
}

func (o EligibleApproverResponseOutput) ToEligibleApproverResponseOutput() EligibleApproverResponseOutput {
	return o
}

func (o EligibleApproverResponseOutput) ToEligibleApproverResponseOutputWithContext(ctx context.Context) EligibleApproverResponseOutput {
	return o
}

func (o EligibleApproverResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleApproverResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o EligibleApproverResponseOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EligibleApproverResponse) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

type EligibleApproverResponseArrayOutput struct{ *pulumi.OutputState }

func (EligibleApproverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleApproverResponse)(nil)).Elem()
}

func (o EligibleApproverResponseArrayOutput) ToEligibleApproverResponseArrayOutput() EligibleApproverResponseArrayOutput {
	return o
}

func (o EligibleApproverResponseArrayOutput) ToEligibleApproverResponseArrayOutputWithContext(ctx context.Context) EligibleApproverResponseArrayOutput {
	return o
}

func (o EligibleApproverResponseArrayOutput) Index(i pulumi.IntInput) EligibleApproverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EligibleApproverResponse {
		return vs[0].([]EligibleApproverResponse)[vs[1].(int)]
	}).(EligibleApproverResponseOutput)
}

type EligibleAuthorization struct {
	JustInTimeAccessPolicy *JustInTimeAccessPolicy `pulumi:"justInTimeAccessPolicy"`
	PrincipalId            string                  `pulumi:"principalId"`
	PrincipalIdDisplayName *string                 `pulumi:"principalIdDisplayName"`
	RoleDefinitionId       string                  `pulumi:"roleDefinitionId"`
}





type EligibleAuthorizationInput interface {
	pulumi.Input

	ToEligibleAuthorizationOutput() EligibleAuthorizationOutput
	ToEligibleAuthorizationOutputWithContext(context.Context) EligibleAuthorizationOutput
}

type EligibleAuthorizationArgs struct {
	JustInTimeAccessPolicy JustInTimeAccessPolicyPtrInput `pulumi:"justInTimeAccessPolicy"`
	PrincipalId            pulumi.StringInput             `pulumi:"principalId"`
	PrincipalIdDisplayName pulumi.StringPtrInput          `pulumi:"principalIdDisplayName"`
	RoleDefinitionId       pulumi.StringInput             `pulumi:"roleDefinitionId"`
}

func (EligibleAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleAuthorization)(nil)).Elem()
}

func (i EligibleAuthorizationArgs) ToEligibleAuthorizationOutput() EligibleAuthorizationOutput {
	return i.ToEligibleAuthorizationOutputWithContext(context.Background())
}

func (i EligibleAuthorizationArgs) ToEligibleAuthorizationOutputWithContext(ctx context.Context) EligibleAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleAuthorizationOutput)
}





type EligibleAuthorizationArrayInput interface {
	pulumi.Input

	ToEligibleAuthorizationArrayOutput() EligibleAuthorizationArrayOutput
	ToEligibleAuthorizationArrayOutputWithContext(context.Context) EligibleAuthorizationArrayOutput
}

type EligibleAuthorizationArray []EligibleAuthorizationInput

func (EligibleAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleAuthorization)(nil)).Elem()
}

func (i EligibleAuthorizationArray) ToEligibleAuthorizationArrayOutput() EligibleAuthorizationArrayOutput {
	return i.ToEligibleAuthorizationArrayOutputWithContext(context.Background())
}

func (i EligibleAuthorizationArray) ToEligibleAuthorizationArrayOutputWithContext(ctx context.Context) EligibleAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleAuthorizationArrayOutput)
}

type EligibleAuthorizationOutput struct{ *pulumi.OutputState }

func (EligibleAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleAuthorization)(nil)).Elem()
}

func (o EligibleAuthorizationOutput) ToEligibleAuthorizationOutput() EligibleAuthorizationOutput {
	return o
}

func (o EligibleAuthorizationOutput) ToEligibleAuthorizationOutputWithContext(ctx context.Context) EligibleAuthorizationOutput {
	return o
}

func (o EligibleAuthorizationOutput) JustInTimeAccessPolicy() JustInTimeAccessPolicyPtrOutput {
	return o.ApplyT(func(v EligibleAuthorization) *JustInTimeAccessPolicy { return v.JustInTimeAccessPolicy }).(JustInTimeAccessPolicyPtrOutput)
}

func (o EligibleAuthorizationOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleAuthorization) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o EligibleAuthorizationOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EligibleAuthorization) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

func (o EligibleAuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleAuthorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type EligibleAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (EligibleAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleAuthorization)(nil)).Elem()
}

func (o EligibleAuthorizationArrayOutput) ToEligibleAuthorizationArrayOutput() EligibleAuthorizationArrayOutput {
	return o
}

func (o EligibleAuthorizationArrayOutput) ToEligibleAuthorizationArrayOutputWithContext(ctx context.Context) EligibleAuthorizationArrayOutput {
	return o
}

func (o EligibleAuthorizationArrayOutput) Index(i pulumi.IntInput) EligibleAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EligibleAuthorization {
		return vs[0].([]EligibleAuthorization)[vs[1].(int)]
	}).(EligibleAuthorizationOutput)
}

type EligibleAuthorizationResponse struct {
	JustInTimeAccessPolicy *JustInTimeAccessPolicyResponse `pulumi:"justInTimeAccessPolicy"`
	PrincipalId            string                          `pulumi:"principalId"`
	PrincipalIdDisplayName *string                         `pulumi:"principalIdDisplayName"`
	RoleDefinitionId       string                          `pulumi:"roleDefinitionId"`
}





type EligibleAuthorizationResponseInput interface {
	pulumi.Input

	ToEligibleAuthorizationResponseOutput() EligibleAuthorizationResponseOutput
	ToEligibleAuthorizationResponseOutputWithContext(context.Context) EligibleAuthorizationResponseOutput
}

type EligibleAuthorizationResponseArgs struct {
	JustInTimeAccessPolicy JustInTimeAccessPolicyResponsePtrInput `pulumi:"justInTimeAccessPolicy"`
	PrincipalId            pulumi.StringInput                     `pulumi:"principalId"`
	PrincipalIdDisplayName pulumi.StringPtrInput                  `pulumi:"principalIdDisplayName"`
	RoleDefinitionId       pulumi.StringInput                     `pulumi:"roleDefinitionId"`
}

func (EligibleAuthorizationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleAuthorizationResponse)(nil)).Elem()
}

func (i EligibleAuthorizationResponseArgs) ToEligibleAuthorizationResponseOutput() EligibleAuthorizationResponseOutput {
	return i.ToEligibleAuthorizationResponseOutputWithContext(context.Background())
}

func (i EligibleAuthorizationResponseArgs) ToEligibleAuthorizationResponseOutputWithContext(ctx context.Context) EligibleAuthorizationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleAuthorizationResponseOutput)
}





type EligibleAuthorizationResponseArrayInput interface {
	pulumi.Input

	ToEligibleAuthorizationResponseArrayOutput() EligibleAuthorizationResponseArrayOutput
	ToEligibleAuthorizationResponseArrayOutputWithContext(context.Context) EligibleAuthorizationResponseArrayOutput
}

type EligibleAuthorizationResponseArray []EligibleAuthorizationResponseInput

func (EligibleAuthorizationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleAuthorizationResponse)(nil)).Elem()
}

func (i EligibleAuthorizationResponseArray) ToEligibleAuthorizationResponseArrayOutput() EligibleAuthorizationResponseArrayOutput {
	return i.ToEligibleAuthorizationResponseArrayOutputWithContext(context.Background())
}

func (i EligibleAuthorizationResponseArray) ToEligibleAuthorizationResponseArrayOutputWithContext(ctx context.Context) EligibleAuthorizationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleAuthorizationResponseArrayOutput)
}

type EligibleAuthorizationResponseOutput struct{ *pulumi.OutputState }

func (EligibleAuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleAuthorizationResponse)(nil)).Elem()
}

func (o EligibleAuthorizationResponseOutput) ToEligibleAuthorizationResponseOutput() EligibleAuthorizationResponseOutput {
	return o
}

func (o EligibleAuthorizationResponseOutput) ToEligibleAuthorizationResponseOutputWithContext(ctx context.Context) EligibleAuthorizationResponseOutput {
	return o
}

func (o EligibleAuthorizationResponseOutput) JustInTimeAccessPolicy() JustInTimeAccessPolicyResponsePtrOutput {
	return o.ApplyT(func(v EligibleAuthorizationResponse) *JustInTimeAccessPolicyResponse { return v.JustInTimeAccessPolicy }).(JustInTimeAccessPolicyResponsePtrOutput)
}

func (o EligibleAuthorizationResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleAuthorizationResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o EligibleAuthorizationResponseOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EligibleAuthorizationResponse) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

func (o EligibleAuthorizationResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleAuthorizationResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type EligibleAuthorizationResponseArrayOutput struct{ *pulumi.OutputState }

func (EligibleAuthorizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleAuthorizationResponse)(nil)).Elem()
}

func (o EligibleAuthorizationResponseArrayOutput) ToEligibleAuthorizationResponseArrayOutput() EligibleAuthorizationResponseArrayOutput {
	return o
}

func (o EligibleAuthorizationResponseArrayOutput) ToEligibleAuthorizationResponseArrayOutputWithContext(ctx context.Context) EligibleAuthorizationResponseArrayOutput {
	return o
}

func (o EligibleAuthorizationResponseArrayOutput) Index(i pulumi.IntInput) EligibleAuthorizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EligibleAuthorizationResponse {
		return vs[0].([]EligibleAuthorizationResponse)[vs[1].(int)]
	}).(EligibleAuthorizationResponseOutput)
}

type JustInTimeAccessPolicy struct {
	ManagedByTenantApprovers  []EligibleApprover `pulumi:"managedByTenantApprovers"`
	MaximumActivationDuration *string            `pulumi:"maximumActivationDuration"`
	MultiFactorAuthProvider   string             `pulumi:"multiFactorAuthProvider"`
}





type JustInTimeAccessPolicyInput interface {
	pulumi.Input

	ToJustInTimeAccessPolicyOutput() JustInTimeAccessPolicyOutput
	ToJustInTimeAccessPolicyOutputWithContext(context.Context) JustInTimeAccessPolicyOutput
}

type JustInTimeAccessPolicyArgs struct {
	ManagedByTenantApprovers  EligibleApproverArrayInput `pulumi:"managedByTenantApprovers"`
	MaximumActivationDuration pulumi.StringPtrInput      `pulumi:"maximumActivationDuration"`
	MultiFactorAuthProvider   pulumi.StringInput         `pulumi:"multiFactorAuthProvider"`
}

func (JustInTimeAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JustInTimeAccessPolicy)(nil)).Elem()
}

func (i JustInTimeAccessPolicyArgs) ToJustInTimeAccessPolicyOutput() JustInTimeAccessPolicyOutput {
	return i.ToJustInTimeAccessPolicyOutputWithContext(context.Background())
}

func (i JustInTimeAccessPolicyArgs) ToJustInTimeAccessPolicyOutputWithContext(ctx context.Context) JustInTimeAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JustInTimeAccessPolicyOutput)
}

func (i JustInTimeAccessPolicyArgs) ToJustInTimeAccessPolicyPtrOutput() JustInTimeAccessPolicyPtrOutput {
	return i.ToJustInTimeAccessPolicyPtrOutputWithContext(context.Background())
}

func (i JustInTimeAccessPolicyArgs) ToJustInTimeAccessPolicyPtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JustInTimeAccessPolicyOutput).ToJustInTimeAccessPolicyPtrOutputWithContext(ctx)
}









type JustInTimeAccessPolicyPtrInput interface {
	pulumi.Input

	ToJustInTimeAccessPolicyPtrOutput() JustInTimeAccessPolicyPtrOutput
	ToJustInTimeAccessPolicyPtrOutputWithContext(context.Context) JustInTimeAccessPolicyPtrOutput
}

type justInTimeAccessPolicyPtrType JustInTimeAccessPolicyArgs

func JustInTimeAccessPolicyPtr(v *JustInTimeAccessPolicyArgs) JustInTimeAccessPolicyPtrInput {
	return (*justInTimeAccessPolicyPtrType)(v)
}

func (*justInTimeAccessPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JustInTimeAccessPolicy)(nil)).Elem()
}

func (i *justInTimeAccessPolicyPtrType) ToJustInTimeAccessPolicyPtrOutput() JustInTimeAccessPolicyPtrOutput {
	return i.ToJustInTimeAccessPolicyPtrOutputWithContext(context.Background())
}

func (i *justInTimeAccessPolicyPtrType) ToJustInTimeAccessPolicyPtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JustInTimeAccessPolicyPtrOutput)
}

type JustInTimeAccessPolicyOutput struct{ *pulumi.OutputState }

func (JustInTimeAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JustInTimeAccessPolicy)(nil)).Elem()
}

func (o JustInTimeAccessPolicyOutput) ToJustInTimeAccessPolicyOutput() JustInTimeAccessPolicyOutput {
	return o
}

func (o JustInTimeAccessPolicyOutput) ToJustInTimeAccessPolicyOutputWithContext(ctx context.Context) JustInTimeAccessPolicyOutput {
	return o
}

func (o JustInTimeAccessPolicyOutput) ToJustInTimeAccessPolicyPtrOutput() JustInTimeAccessPolicyPtrOutput {
	return o.ToJustInTimeAccessPolicyPtrOutputWithContext(context.Background())
}

func (o JustInTimeAccessPolicyOutput) ToJustInTimeAccessPolicyPtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JustInTimeAccessPolicy) *JustInTimeAccessPolicy {
		return &v
	}).(JustInTimeAccessPolicyPtrOutput)
}

func (o JustInTimeAccessPolicyOutput) ManagedByTenantApprovers() EligibleApproverArrayOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicy) []EligibleApprover { return v.ManagedByTenantApprovers }).(EligibleApproverArrayOutput)
}

func (o JustInTimeAccessPolicyOutput) MaximumActivationDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicy) *string { return v.MaximumActivationDuration }).(pulumi.StringPtrOutput)
}

func (o JustInTimeAccessPolicyOutput) MultiFactorAuthProvider() pulumi.StringOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicy) string { return v.MultiFactorAuthProvider }).(pulumi.StringOutput)
}

type JustInTimeAccessPolicyPtrOutput struct{ *pulumi.OutputState }

func (JustInTimeAccessPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JustInTimeAccessPolicy)(nil)).Elem()
}

func (o JustInTimeAccessPolicyPtrOutput) ToJustInTimeAccessPolicyPtrOutput() JustInTimeAccessPolicyPtrOutput {
	return o
}

func (o JustInTimeAccessPolicyPtrOutput) ToJustInTimeAccessPolicyPtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyPtrOutput {
	return o
}

func (o JustInTimeAccessPolicyPtrOutput) Elem() JustInTimeAccessPolicyOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicy) JustInTimeAccessPolicy {
		if v != nil {
			return *v
		}
		var ret JustInTimeAccessPolicy
		return ret
	}).(JustInTimeAccessPolicyOutput)
}

func (o JustInTimeAccessPolicyPtrOutput) ManagedByTenantApprovers() EligibleApproverArrayOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicy) []EligibleApprover {
		if v == nil {
			return nil
		}
		return v.ManagedByTenantApprovers
	}).(EligibleApproverArrayOutput)
}

func (o JustInTimeAccessPolicyPtrOutput) MaximumActivationDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicy) *string {
		if v == nil {
			return nil
		}
		return v.MaximumActivationDuration
	}).(pulumi.StringPtrOutput)
}

func (o JustInTimeAccessPolicyPtrOutput) MultiFactorAuthProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.MultiFactorAuthProvider
	}).(pulumi.StringPtrOutput)
}

type JustInTimeAccessPolicyResponse struct {
	ManagedByTenantApprovers  []EligibleApproverResponse `pulumi:"managedByTenantApprovers"`
	MaximumActivationDuration *string                    `pulumi:"maximumActivationDuration"`
	MultiFactorAuthProvider   string                     `pulumi:"multiFactorAuthProvider"`
}





type JustInTimeAccessPolicyResponseInput interface {
	pulumi.Input

	ToJustInTimeAccessPolicyResponseOutput() JustInTimeAccessPolicyResponseOutput
	ToJustInTimeAccessPolicyResponseOutputWithContext(context.Context) JustInTimeAccessPolicyResponseOutput
}

type JustInTimeAccessPolicyResponseArgs struct {
	ManagedByTenantApprovers  EligibleApproverResponseArrayInput `pulumi:"managedByTenantApprovers"`
	MaximumActivationDuration pulumi.StringPtrInput              `pulumi:"maximumActivationDuration"`
	MultiFactorAuthProvider   pulumi.StringInput                 `pulumi:"multiFactorAuthProvider"`
}

func (JustInTimeAccessPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JustInTimeAccessPolicyResponse)(nil)).Elem()
}

func (i JustInTimeAccessPolicyResponseArgs) ToJustInTimeAccessPolicyResponseOutput() JustInTimeAccessPolicyResponseOutput {
	return i.ToJustInTimeAccessPolicyResponseOutputWithContext(context.Background())
}

func (i JustInTimeAccessPolicyResponseArgs) ToJustInTimeAccessPolicyResponseOutputWithContext(ctx context.Context) JustInTimeAccessPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JustInTimeAccessPolicyResponseOutput)
}

func (i JustInTimeAccessPolicyResponseArgs) ToJustInTimeAccessPolicyResponsePtrOutput() JustInTimeAccessPolicyResponsePtrOutput {
	return i.ToJustInTimeAccessPolicyResponsePtrOutputWithContext(context.Background())
}

func (i JustInTimeAccessPolicyResponseArgs) ToJustInTimeAccessPolicyResponsePtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JustInTimeAccessPolicyResponseOutput).ToJustInTimeAccessPolicyResponsePtrOutputWithContext(ctx)
}









type JustInTimeAccessPolicyResponsePtrInput interface {
	pulumi.Input

	ToJustInTimeAccessPolicyResponsePtrOutput() JustInTimeAccessPolicyResponsePtrOutput
	ToJustInTimeAccessPolicyResponsePtrOutputWithContext(context.Context) JustInTimeAccessPolicyResponsePtrOutput
}

type justInTimeAccessPolicyResponsePtrType JustInTimeAccessPolicyResponseArgs

func JustInTimeAccessPolicyResponsePtr(v *JustInTimeAccessPolicyResponseArgs) JustInTimeAccessPolicyResponsePtrInput {
	return (*justInTimeAccessPolicyResponsePtrType)(v)
}

func (*justInTimeAccessPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JustInTimeAccessPolicyResponse)(nil)).Elem()
}

func (i *justInTimeAccessPolicyResponsePtrType) ToJustInTimeAccessPolicyResponsePtrOutput() JustInTimeAccessPolicyResponsePtrOutput {
	return i.ToJustInTimeAccessPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *justInTimeAccessPolicyResponsePtrType) ToJustInTimeAccessPolicyResponsePtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JustInTimeAccessPolicyResponsePtrOutput)
}

type JustInTimeAccessPolicyResponseOutput struct{ *pulumi.OutputState }

func (JustInTimeAccessPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JustInTimeAccessPolicyResponse)(nil)).Elem()
}

func (o JustInTimeAccessPolicyResponseOutput) ToJustInTimeAccessPolicyResponseOutput() JustInTimeAccessPolicyResponseOutput {
	return o
}

func (o JustInTimeAccessPolicyResponseOutput) ToJustInTimeAccessPolicyResponseOutputWithContext(ctx context.Context) JustInTimeAccessPolicyResponseOutput {
	return o
}

func (o JustInTimeAccessPolicyResponseOutput) ToJustInTimeAccessPolicyResponsePtrOutput() JustInTimeAccessPolicyResponsePtrOutput {
	return o.ToJustInTimeAccessPolicyResponsePtrOutputWithContext(context.Background())
}

func (o JustInTimeAccessPolicyResponseOutput) ToJustInTimeAccessPolicyResponsePtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JustInTimeAccessPolicyResponse) *JustInTimeAccessPolicyResponse {
		return &v
	}).(JustInTimeAccessPolicyResponsePtrOutput)
}

func (o JustInTimeAccessPolicyResponseOutput) ManagedByTenantApprovers() EligibleApproverResponseArrayOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicyResponse) []EligibleApproverResponse { return v.ManagedByTenantApprovers }).(EligibleApproverResponseArrayOutput)
}

func (o JustInTimeAccessPolicyResponseOutput) MaximumActivationDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicyResponse) *string { return v.MaximumActivationDuration }).(pulumi.StringPtrOutput)
}

func (o JustInTimeAccessPolicyResponseOutput) MultiFactorAuthProvider() pulumi.StringOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicyResponse) string { return v.MultiFactorAuthProvider }).(pulumi.StringOutput)
}

type JustInTimeAccessPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (JustInTimeAccessPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JustInTimeAccessPolicyResponse)(nil)).Elem()
}

func (o JustInTimeAccessPolicyResponsePtrOutput) ToJustInTimeAccessPolicyResponsePtrOutput() JustInTimeAccessPolicyResponsePtrOutput {
	return o
}

func (o JustInTimeAccessPolicyResponsePtrOutput) ToJustInTimeAccessPolicyResponsePtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyResponsePtrOutput {
	return o
}

func (o JustInTimeAccessPolicyResponsePtrOutput) Elem() JustInTimeAccessPolicyResponseOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicyResponse) JustInTimeAccessPolicyResponse {
		if v != nil {
			return *v
		}
		var ret JustInTimeAccessPolicyResponse
		return ret
	}).(JustInTimeAccessPolicyResponseOutput)
}

func (o JustInTimeAccessPolicyResponsePtrOutput) ManagedByTenantApprovers() EligibleApproverResponseArrayOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicyResponse) []EligibleApproverResponse {
		if v == nil {
			return nil
		}
		return v.ManagedByTenantApprovers
	}).(EligibleApproverResponseArrayOutput)
}

func (o JustInTimeAccessPolicyResponsePtrOutput) MaximumActivationDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaximumActivationDuration
	}).(pulumi.StringPtrOutput)
}

func (o JustInTimeAccessPolicyResponsePtrOutput) MultiFactorAuthProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MultiFactorAuthProvider
	}).(pulumi.StringPtrOutput)
}

type Plan struct {
	Name      string `pulumi:"name"`
	Product   string `pulumi:"product"`
	Publisher string `pulumi:"publisher"`
	Version   string `pulumi:"version"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Name      pulumi.StringInput `pulumi:"name"`
	Product   pulumi.StringInput `pulumi:"product"`
	Publisher pulumi.StringInput `pulumi:"publisher"`
	Version   pulumi.StringInput `pulumi:"version"`
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
	Name      string `pulumi:"name"`
	Product   string `pulumi:"product"`
	Publisher string `pulumi:"publisher"`
	Version   string `pulumi:"version"`
}





type PlanResponseInput interface {
	pulumi.Input

	ToPlanResponseOutput() PlanResponseOutput
	ToPlanResponseOutputWithContext(context.Context) PlanResponseOutput
}

type PlanResponseArgs struct {
	Name      pulumi.StringInput `pulumi:"name"`
	Product   pulumi.StringInput `pulumi:"product"`
	Publisher pulumi.StringInput `pulumi:"publisher"`
	Version   pulumi.StringInput `pulumi:"version"`
}

func (PlanResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (i PlanResponseArgs) ToPlanResponseOutput() PlanResponseOutput {
	return i.ToPlanResponseOutputWithContext(context.Background())
}

func (i PlanResponseArgs) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponseOutput)
}

func (i PlanResponseArgs) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return i.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (i PlanResponseArgs) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponseOutput).ToPlanResponsePtrOutputWithContext(ctx)
}









type PlanResponsePtrInput interface {
	pulumi.Input

	ToPlanResponsePtrOutput() PlanResponsePtrOutput
	ToPlanResponsePtrOutputWithContext(context.Context) PlanResponsePtrOutput
}

type planResponsePtrType PlanResponseArgs

func PlanResponsePtr(v *PlanResponseArgs) PlanResponsePtrInput {
	return (*planResponsePtrType)(v)
}

func (*planResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (i *planResponsePtrType) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return i.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (i *planResponsePtrType) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponsePtrOutput)
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

func (o PlanResponseOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (o PlanResponseOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlanResponse) *PlanResponse {
		return &v
	}).(PlanResponsePtrOutput)
}

func (o PlanResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Product }).(pulumi.StringOutput)
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

type RegistrationAssignmentProperties struct {
	RegistrationDefinitionId string `pulumi:"registrationDefinitionId"`
}





type RegistrationAssignmentPropertiesInput interface {
	pulumi.Input

	ToRegistrationAssignmentPropertiesOutput() RegistrationAssignmentPropertiesOutput
	ToRegistrationAssignmentPropertiesOutputWithContext(context.Context) RegistrationAssignmentPropertiesOutput
}

type RegistrationAssignmentPropertiesArgs struct {
	RegistrationDefinitionId pulumi.StringInput `pulumi:"registrationDefinitionId"`
}

func (RegistrationAssignmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentProperties)(nil)).Elem()
}

func (i RegistrationAssignmentPropertiesArgs) ToRegistrationAssignmentPropertiesOutput() RegistrationAssignmentPropertiesOutput {
	return i.ToRegistrationAssignmentPropertiesOutputWithContext(context.Background())
}

func (i RegistrationAssignmentPropertiesArgs) ToRegistrationAssignmentPropertiesOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesOutput)
}

func (i RegistrationAssignmentPropertiesArgs) ToRegistrationAssignmentPropertiesPtrOutput() RegistrationAssignmentPropertiesPtrOutput {
	return i.ToRegistrationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i RegistrationAssignmentPropertiesArgs) ToRegistrationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesOutput).ToRegistrationAssignmentPropertiesPtrOutputWithContext(ctx)
}









type RegistrationAssignmentPropertiesPtrInput interface {
	pulumi.Input

	ToRegistrationAssignmentPropertiesPtrOutput() RegistrationAssignmentPropertiesPtrOutput
	ToRegistrationAssignmentPropertiesPtrOutputWithContext(context.Context) RegistrationAssignmentPropertiesPtrOutput
}

type registrationAssignmentPropertiesPtrType RegistrationAssignmentPropertiesArgs

func RegistrationAssignmentPropertiesPtr(v *RegistrationAssignmentPropertiesArgs) RegistrationAssignmentPropertiesPtrInput {
	return (*registrationAssignmentPropertiesPtrType)(v)
}

func (*registrationAssignmentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationAssignmentProperties)(nil)).Elem()
}

func (i *registrationAssignmentPropertiesPtrType) ToRegistrationAssignmentPropertiesPtrOutput() RegistrationAssignmentPropertiesPtrOutput {
	return i.ToRegistrationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i *registrationAssignmentPropertiesPtrType) ToRegistrationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesPtrOutput)
}

type RegistrationAssignmentPropertiesOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentProperties)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesOutput) ToRegistrationAssignmentPropertiesOutput() RegistrationAssignmentPropertiesOutput {
	return o
}

func (o RegistrationAssignmentPropertiesOutput) ToRegistrationAssignmentPropertiesOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesOutput {
	return o
}

func (o RegistrationAssignmentPropertiesOutput) ToRegistrationAssignmentPropertiesPtrOutput() RegistrationAssignmentPropertiesPtrOutput {
	return o.ToRegistrationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (o RegistrationAssignmentPropertiesOutput) ToRegistrationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationAssignmentProperties) *RegistrationAssignmentProperties {
		return &v
	}).(RegistrationAssignmentPropertiesPtrOutput)
}

func (o RegistrationAssignmentPropertiesOutput) RegistrationDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentProperties) string { return v.RegistrationDefinitionId }).(pulumi.StringOutput)
}

type RegistrationAssignmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationAssignmentProperties)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesPtrOutput) ToRegistrationAssignmentPropertiesPtrOutput() RegistrationAssignmentPropertiesPtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesPtrOutput) ToRegistrationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesPtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesPtrOutput) Elem() RegistrationAssignmentPropertiesOutput {
	return o.ApplyT(func(v *RegistrationAssignmentProperties) RegistrationAssignmentProperties {
		if v != nil {
			return *v
		}
		var ret RegistrationAssignmentProperties
		return ret
	}).(RegistrationAssignmentPropertiesOutput)
}

func (o RegistrationAssignmentPropertiesPtrOutput) RegistrationDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.RegistrationDefinitionId
	}).(pulumi.StringPtrOutput)
}

type RegistrationAssignmentPropertiesResponse struct {
	ProvisioningState        string                                                         `pulumi:"provisioningState"`
	RegistrationDefinition   RegistrationAssignmentPropertiesResponseRegistrationDefinition `pulumi:"registrationDefinition"`
	RegistrationDefinitionId string                                                         `pulumi:"registrationDefinitionId"`
}





type RegistrationAssignmentPropertiesResponseInput interface {
	pulumi.Input

	ToRegistrationAssignmentPropertiesResponseOutput() RegistrationAssignmentPropertiesResponseOutput
	ToRegistrationAssignmentPropertiesResponseOutputWithContext(context.Context) RegistrationAssignmentPropertiesResponseOutput
}

type RegistrationAssignmentPropertiesResponseArgs struct {
	ProvisioningState        pulumi.StringInput                                                  `pulumi:"provisioningState"`
	RegistrationDefinition   RegistrationAssignmentPropertiesResponseRegistrationDefinitionInput `pulumi:"registrationDefinition"`
	RegistrationDefinitionId pulumi.StringInput                                                  `pulumi:"registrationDefinitionId"`
}

func (RegistrationAssignmentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentPropertiesResponse)(nil)).Elem()
}

func (i RegistrationAssignmentPropertiesResponseArgs) ToRegistrationAssignmentPropertiesResponseOutput() RegistrationAssignmentPropertiesResponseOutput {
	return i.ToRegistrationAssignmentPropertiesResponseOutputWithContext(context.Background())
}

func (i RegistrationAssignmentPropertiesResponseArgs) ToRegistrationAssignmentPropertiesResponseOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesResponseOutput)
}

func (i RegistrationAssignmentPropertiesResponseArgs) ToRegistrationAssignmentPropertiesResponsePtrOutput() RegistrationAssignmentPropertiesResponsePtrOutput {
	return i.ToRegistrationAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i RegistrationAssignmentPropertiesResponseArgs) ToRegistrationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesResponseOutput).ToRegistrationAssignmentPropertiesResponsePtrOutputWithContext(ctx)
}









type RegistrationAssignmentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToRegistrationAssignmentPropertiesResponsePtrOutput() RegistrationAssignmentPropertiesResponsePtrOutput
	ToRegistrationAssignmentPropertiesResponsePtrOutputWithContext(context.Context) RegistrationAssignmentPropertiesResponsePtrOutput
}

type registrationAssignmentPropertiesResponsePtrType RegistrationAssignmentPropertiesResponseArgs

func RegistrationAssignmentPropertiesResponsePtr(v *RegistrationAssignmentPropertiesResponseArgs) RegistrationAssignmentPropertiesResponsePtrInput {
	return (*registrationAssignmentPropertiesResponsePtrType)(v)
}

func (*registrationAssignmentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationAssignmentPropertiesResponse)(nil)).Elem()
}

func (i *registrationAssignmentPropertiesResponsePtrType) ToRegistrationAssignmentPropertiesResponsePtrOutput() RegistrationAssignmentPropertiesResponsePtrOutput {
	return i.ToRegistrationAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *registrationAssignmentPropertiesResponsePtrType) ToRegistrationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesResponsePtrOutput)
}

type RegistrationAssignmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentPropertiesResponse)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesResponseOutput) ToRegistrationAssignmentPropertiesResponseOutput() RegistrationAssignmentPropertiesResponseOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponseOutput) ToRegistrationAssignmentPropertiesResponseOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponseOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponseOutput) ToRegistrationAssignmentPropertiesResponsePtrOutput() RegistrationAssignmentPropertiesResponsePtrOutput {
	return o.ToRegistrationAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o RegistrationAssignmentPropertiesResponseOutput) ToRegistrationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationAssignmentPropertiesResponse) *RegistrationAssignmentPropertiesResponse {
		return &v
	}).(RegistrationAssignmentPropertiesResponsePtrOutput)
}

func (o RegistrationAssignmentPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RegistrationAssignmentPropertiesResponseOutput) RegistrationDefinition() RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponse) RegistrationAssignmentPropertiesResponseRegistrationDefinition {
		return v.RegistrationDefinition
	}).(RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput)
}

func (o RegistrationAssignmentPropertiesResponseOutput) RegistrationDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponse) string { return v.RegistrationDefinitionId }).(pulumi.StringOutput)
}

type RegistrationAssignmentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationAssignmentPropertiesResponse)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesResponsePtrOutput) ToRegistrationAssignmentPropertiesResponsePtrOutput() RegistrationAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponsePtrOutput) ToRegistrationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponsePtrOutput) Elem() RegistrationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponse) RegistrationAssignmentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret RegistrationAssignmentPropertiesResponse
		return ret
	}).(RegistrationAssignmentPropertiesResponseOutput)
}

func (o RegistrationAssignmentPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePtrOutput) RegistrationDefinition() RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponse) *RegistrationAssignmentPropertiesResponseRegistrationDefinition {
		if v == nil {
			return nil
		}
		return &v.RegistrationDefinition
	}).(RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePtrOutput) RegistrationDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RegistrationDefinitionId
	}).(pulumi.StringPtrOutput)
}

type RegistrationAssignmentPropertiesResponseProperties struct {
	Authorizations             []AuthorizationResponse         `pulumi:"authorizations"`
	Description                *string                         `pulumi:"description"`
	EligibleAuthorizations     []EligibleAuthorizationResponse `pulumi:"eligibleAuthorizations"`
	ManagedByTenantId          *string                         `pulumi:"managedByTenantId"`
	ManagedByTenantName        *string                         `pulumi:"managedByTenantName"`
	ManageeTenantId            *string                         `pulumi:"manageeTenantId"`
	ManageeTenantName          *string                         `pulumi:"manageeTenantName"`
	ProvisioningState          *string                         `pulumi:"provisioningState"`
	RegistrationDefinitionName *string                         `pulumi:"registrationDefinitionName"`
}





type RegistrationAssignmentPropertiesResponsePropertiesInput interface {
	pulumi.Input

	ToRegistrationAssignmentPropertiesResponsePropertiesOutput() RegistrationAssignmentPropertiesResponsePropertiesOutput
	ToRegistrationAssignmentPropertiesResponsePropertiesOutputWithContext(context.Context) RegistrationAssignmentPropertiesResponsePropertiesOutput
}

type RegistrationAssignmentPropertiesResponsePropertiesArgs struct {
	Authorizations             AuthorizationResponseArrayInput         `pulumi:"authorizations"`
	Description                pulumi.StringPtrInput                   `pulumi:"description"`
	EligibleAuthorizations     EligibleAuthorizationResponseArrayInput `pulumi:"eligibleAuthorizations"`
	ManagedByTenantId          pulumi.StringPtrInput                   `pulumi:"managedByTenantId"`
	ManagedByTenantName        pulumi.StringPtrInput                   `pulumi:"managedByTenantName"`
	ManageeTenantId            pulumi.StringPtrInput                   `pulumi:"manageeTenantId"`
	ManageeTenantName          pulumi.StringPtrInput                   `pulumi:"manageeTenantName"`
	ProvisioningState          pulumi.StringPtrInput                   `pulumi:"provisioningState"`
	RegistrationDefinitionName pulumi.StringPtrInput                   `pulumi:"registrationDefinitionName"`
}

func (RegistrationAssignmentPropertiesResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentPropertiesResponseProperties)(nil)).Elem()
}

func (i RegistrationAssignmentPropertiesResponsePropertiesArgs) ToRegistrationAssignmentPropertiesResponsePropertiesOutput() RegistrationAssignmentPropertiesResponsePropertiesOutput {
	return i.ToRegistrationAssignmentPropertiesResponsePropertiesOutputWithContext(context.Background())
}

func (i RegistrationAssignmentPropertiesResponsePropertiesArgs) ToRegistrationAssignmentPropertiesResponsePropertiesOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesResponsePropertiesOutput)
}

func (i RegistrationAssignmentPropertiesResponsePropertiesArgs) ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutput() RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return i.ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i RegistrationAssignmentPropertiesResponsePropertiesArgs) ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesResponsePropertiesOutput).ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutputWithContext(ctx)
}









type RegistrationAssignmentPropertiesResponsePropertiesPtrInput interface {
	pulumi.Input

	ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutput() RegistrationAssignmentPropertiesResponsePropertiesPtrOutput
	ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutputWithContext(context.Context) RegistrationAssignmentPropertiesResponsePropertiesPtrOutput
}

type registrationAssignmentPropertiesResponsePropertiesPtrType RegistrationAssignmentPropertiesResponsePropertiesArgs

func RegistrationAssignmentPropertiesResponsePropertiesPtr(v *RegistrationAssignmentPropertiesResponsePropertiesArgs) RegistrationAssignmentPropertiesResponsePropertiesPtrInput {
	return (*registrationAssignmentPropertiesResponsePropertiesPtrType)(v)
}

func (*registrationAssignmentPropertiesResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationAssignmentPropertiesResponseProperties)(nil)).Elem()
}

func (i *registrationAssignmentPropertiesResponsePropertiesPtrType) ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutput() RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return i.ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *registrationAssignmentPropertiesResponsePropertiesPtrType) ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesResponsePropertiesPtrOutput)
}

type RegistrationAssignmentPropertiesResponsePropertiesOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentPropertiesResponseProperties)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ToRegistrationAssignmentPropertiesResponsePropertiesOutput() RegistrationAssignmentPropertiesResponsePropertiesOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ToRegistrationAssignmentPropertiesResponsePropertiesOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePropertiesOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutput() RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return o.ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationAssignmentPropertiesResponseProperties) *RegistrationAssignmentPropertiesResponseProperties {
		return &v
	}).(RegistrationAssignmentPropertiesResponsePropertiesPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) Authorizations() AuthorizationResponseArrayOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) []AuthorizationResponse {
		return v.Authorizations
	}).(AuthorizationResponseArrayOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) EligibleAuthorizations() EligibleAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) []EligibleAuthorizationResponse {
		return v.EligibleAuthorizations
	}).(EligibleAuthorizationResponseArrayOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ManagedByTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.ManagedByTenantId }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ManagedByTenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.ManagedByTenantName }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ManageeTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.ManageeTenantId }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ManageeTenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.ManageeTenantName }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) RegistrationDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string {
		return v.RegistrationDefinitionName
	}).(pulumi.StringPtrOutput)
}

type RegistrationAssignmentPropertiesResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationAssignmentPropertiesResponseProperties)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutput() RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) Elem() RegistrationAssignmentPropertiesResponsePropertiesOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) RegistrationAssignmentPropertiesResponseProperties {
		if v != nil {
			return *v
		}
		var ret RegistrationAssignmentPropertiesResponseProperties
		return ret
	}).(RegistrationAssignmentPropertiesResponsePropertiesOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) Authorizations() AuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) []AuthorizationResponse {
		if v == nil {
			return nil
		}
		return v.Authorizations
	}).(AuthorizationResponseArrayOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) EligibleAuthorizations() EligibleAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) []EligibleAuthorizationResponse {
		if v == nil {
			return nil
		}
		return v.EligibleAuthorizations
	}).(EligibleAuthorizationResponseArrayOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ManagedByTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManagedByTenantId
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ManagedByTenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManagedByTenantName
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ManageeTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManageeTenantId
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ManageeTenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManageeTenantName
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) RegistrationDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationDefinitionName
	}).(pulumi.StringPtrOutput)
}

type RegistrationAssignmentPropertiesResponseRegistrationDefinition struct {
	Id         string                                              `pulumi:"id"`
	Name       string                                              `pulumi:"name"`
	Plan       *PlanResponse                                       `pulumi:"plan"`
	Properties *RegistrationAssignmentPropertiesResponseProperties `pulumi:"properties"`
	Type       string                                              `pulumi:"type"`
}





type RegistrationAssignmentPropertiesResponseRegistrationDefinitionInput interface {
	pulumi.Input

	ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput() RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput
	ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionOutputWithContext(context.Context) RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput
}

type RegistrationAssignmentPropertiesResponseRegistrationDefinitionArgs struct {
	Id         pulumi.StringInput                                         `pulumi:"id"`
	Name       pulumi.StringInput                                         `pulumi:"name"`
	Plan       PlanResponsePtrInput                                       `pulumi:"plan"`
	Properties RegistrationAssignmentPropertiesResponsePropertiesPtrInput `pulumi:"properties"`
	Type       pulumi.StringInput                                         `pulumi:"type"`
}

func (RegistrationAssignmentPropertiesResponseRegistrationDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentPropertiesResponseRegistrationDefinition)(nil)).Elem()
}

func (i RegistrationAssignmentPropertiesResponseRegistrationDefinitionArgs) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput() RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput {
	return i.ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionOutputWithContext(context.Background())
}

func (i RegistrationAssignmentPropertiesResponseRegistrationDefinitionArgs) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput)
}

func (i RegistrationAssignmentPropertiesResponseRegistrationDefinitionArgs) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput() RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput {
	return i.ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutputWithContext(context.Background())
}

func (i RegistrationAssignmentPropertiesResponseRegistrationDefinitionArgs) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput).ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutputWithContext(ctx)
}









type RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrInput interface {
	pulumi.Input

	ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput() RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput
	ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutputWithContext(context.Context) RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput
}

type registrationAssignmentPropertiesResponseRegistrationDefinitionPtrType RegistrationAssignmentPropertiesResponseRegistrationDefinitionArgs

func RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtr(v *RegistrationAssignmentPropertiesResponseRegistrationDefinitionArgs) RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrInput {
	return (*registrationAssignmentPropertiesResponseRegistrationDefinitionPtrType)(v)
}

func (*registrationAssignmentPropertiesResponseRegistrationDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationAssignmentPropertiesResponseRegistrationDefinition)(nil)).Elem()
}

func (i *registrationAssignmentPropertiesResponseRegistrationDefinitionPtrType) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput() RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput {
	return i.ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutputWithContext(context.Background())
}

func (i *registrationAssignmentPropertiesResponseRegistrationDefinitionPtrType) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput)
}

type RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentPropertiesResponseRegistrationDefinition)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput() RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput() RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput {
	return o.ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutputWithContext(context.Background())
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationAssignmentPropertiesResponseRegistrationDefinition) *RegistrationAssignmentPropertiesResponseRegistrationDefinition {
		return &v
	}).(RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseRegistrationDefinition) string { return v.Id }).(pulumi.StringOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseRegistrationDefinition) string { return v.Name }).(pulumi.StringOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseRegistrationDefinition) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) Properties() RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseRegistrationDefinition) *RegistrationAssignmentPropertiesResponseProperties {
		return v.Properties
	}).(RegistrationAssignmentPropertiesResponsePropertiesPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseRegistrationDefinition) string { return v.Type }).(pulumi.StringOutput)
}

type RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationAssignmentPropertiesResponseRegistrationDefinition)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput() RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput) Elem() RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseRegistrationDefinition) RegistrationAssignmentPropertiesResponseRegistrationDefinition {
		if v != nil {
			return *v
		}
		var ret RegistrationAssignmentPropertiesResponseRegistrationDefinition
		return ret
	}).(RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseRegistrationDefinition) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseRegistrationDefinition) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseRegistrationDefinition) *PlanResponse {
		if v == nil {
			return nil
		}
		return v.Plan
	}).(PlanResponsePtrOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput) Properties() RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseRegistrationDefinition) *RegistrationAssignmentPropertiesResponseProperties {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(RegistrationAssignmentPropertiesResponsePropertiesPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseRegistrationDefinition) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type RegistrationDefinitionProperties struct {
	Authorizations             []Authorization         `pulumi:"authorizations"`
	Description                *string                 `pulumi:"description"`
	EligibleAuthorizations     []EligibleAuthorization `pulumi:"eligibleAuthorizations"`
	ManagedByTenantId          string                  `pulumi:"managedByTenantId"`
	RegistrationDefinitionName *string                 `pulumi:"registrationDefinitionName"`
}





type RegistrationDefinitionPropertiesInput interface {
	pulumi.Input

	ToRegistrationDefinitionPropertiesOutput() RegistrationDefinitionPropertiesOutput
	ToRegistrationDefinitionPropertiesOutputWithContext(context.Context) RegistrationDefinitionPropertiesOutput
}

type RegistrationDefinitionPropertiesArgs struct {
	Authorizations             AuthorizationArrayInput         `pulumi:"authorizations"`
	Description                pulumi.StringPtrInput           `pulumi:"description"`
	EligibleAuthorizations     EligibleAuthorizationArrayInput `pulumi:"eligibleAuthorizations"`
	ManagedByTenantId          pulumi.StringInput              `pulumi:"managedByTenantId"`
	RegistrationDefinitionName pulumi.StringPtrInput           `pulumi:"registrationDefinitionName"`
}

func (RegistrationDefinitionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationDefinitionProperties)(nil)).Elem()
}

func (i RegistrationDefinitionPropertiesArgs) ToRegistrationDefinitionPropertiesOutput() RegistrationDefinitionPropertiesOutput {
	return i.ToRegistrationDefinitionPropertiesOutputWithContext(context.Background())
}

func (i RegistrationDefinitionPropertiesArgs) ToRegistrationDefinitionPropertiesOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationDefinitionPropertiesOutput)
}

func (i RegistrationDefinitionPropertiesArgs) ToRegistrationDefinitionPropertiesPtrOutput() RegistrationDefinitionPropertiesPtrOutput {
	return i.ToRegistrationDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (i RegistrationDefinitionPropertiesArgs) ToRegistrationDefinitionPropertiesPtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationDefinitionPropertiesOutput).ToRegistrationDefinitionPropertiesPtrOutputWithContext(ctx)
}









type RegistrationDefinitionPropertiesPtrInput interface {
	pulumi.Input

	ToRegistrationDefinitionPropertiesPtrOutput() RegistrationDefinitionPropertiesPtrOutput
	ToRegistrationDefinitionPropertiesPtrOutputWithContext(context.Context) RegistrationDefinitionPropertiesPtrOutput
}

type registrationDefinitionPropertiesPtrType RegistrationDefinitionPropertiesArgs

func RegistrationDefinitionPropertiesPtr(v *RegistrationDefinitionPropertiesArgs) RegistrationDefinitionPropertiesPtrInput {
	return (*registrationDefinitionPropertiesPtrType)(v)
}

func (*registrationDefinitionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationDefinitionProperties)(nil)).Elem()
}

func (i *registrationDefinitionPropertiesPtrType) ToRegistrationDefinitionPropertiesPtrOutput() RegistrationDefinitionPropertiesPtrOutput {
	return i.ToRegistrationDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (i *registrationDefinitionPropertiesPtrType) ToRegistrationDefinitionPropertiesPtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationDefinitionPropertiesPtrOutput)
}

type RegistrationDefinitionPropertiesOutput struct{ *pulumi.OutputState }

func (RegistrationDefinitionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationDefinitionProperties)(nil)).Elem()
}

func (o RegistrationDefinitionPropertiesOutput) ToRegistrationDefinitionPropertiesOutput() RegistrationDefinitionPropertiesOutput {
	return o
}

func (o RegistrationDefinitionPropertiesOutput) ToRegistrationDefinitionPropertiesOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesOutput {
	return o
}

func (o RegistrationDefinitionPropertiesOutput) ToRegistrationDefinitionPropertiesPtrOutput() RegistrationDefinitionPropertiesPtrOutput {
	return o.ToRegistrationDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (o RegistrationDefinitionPropertiesOutput) ToRegistrationDefinitionPropertiesPtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationDefinitionProperties) *RegistrationDefinitionProperties {
		return &v
	}).(RegistrationDefinitionPropertiesPtrOutput)
}

func (o RegistrationDefinitionPropertiesOutput) Authorizations() AuthorizationArrayOutput {
	return o.ApplyT(func(v RegistrationDefinitionProperties) []Authorization { return v.Authorizations }).(AuthorizationArrayOutput)
}

func (o RegistrationDefinitionPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationDefinitionProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesOutput) EligibleAuthorizations() EligibleAuthorizationArrayOutput {
	return o.ApplyT(func(v RegistrationDefinitionProperties) []EligibleAuthorization { return v.EligibleAuthorizations }).(EligibleAuthorizationArrayOutput)
}

func (o RegistrationDefinitionPropertiesOutput) ManagedByTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationDefinitionProperties) string { return v.ManagedByTenantId }).(pulumi.StringOutput)
}

func (o RegistrationDefinitionPropertiesOutput) RegistrationDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationDefinitionProperties) *string { return v.RegistrationDefinitionName }).(pulumi.StringPtrOutput)
}

type RegistrationDefinitionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (RegistrationDefinitionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationDefinitionProperties)(nil)).Elem()
}

func (o RegistrationDefinitionPropertiesPtrOutput) ToRegistrationDefinitionPropertiesPtrOutput() RegistrationDefinitionPropertiesPtrOutput {
	return o
}

func (o RegistrationDefinitionPropertiesPtrOutput) ToRegistrationDefinitionPropertiesPtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesPtrOutput {
	return o
}

func (o RegistrationDefinitionPropertiesPtrOutput) Elem() RegistrationDefinitionPropertiesOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) RegistrationDefinitionProperties {
		if v != nil {
			return *v
		}
		var ret RegistrationDefinitionProperties
		return ret
	}).(RegistrationDefinitionPropertiesOutput)
}

func (o RegistrationDefinitionPropertiesPtrOutput) Authorizations() AuthorizationArrayOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) []Authorization {
		if v == nil {
			return nil
		}
		return v.Authorizations
	}).(AuthorizationArrayOutput)
}

func (o RegistrationDefinitionPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesPtrOutput) EligibleAuthorizations() EligibleAuthorizationArrayOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) []EligibleAuthorization {
		if v == nil {
			return nil
		}
		return v.EligibleAuthorizations
	}).(EligibleAuthorizationArrayOutput)
}

func (o RegistrationDefinitionPropertiesPtrOutput) ManagedByTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ManagedByTenantId
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesPtrOutput) RegistrationDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationDefinitionName
	}).(pulumi.StringPtrOutput)
}

type RegistrationDefinitionPropertiesResponse struct {
	Authorizations             []AuthorizationResponse         `pulumi:"authorizations"`
	Description                *string                         `pulumi:"description"`
	EligibleAuthorizations     []EligibleAuthorizationResponse `pulumi:"eligibleAuthorizations"`
	ManagedByTenantId          string                          `pulumi:"managedByTenantId"`
	ManagedByTenantName        string                          `pulumi:"managedByTenantName"`
	ProvisioningState          string                          `pulumi:"provisioningState"`
	RegistrationDefinitionName *string                         `pulumi:"registrationDefinitionName"`
}





type RegistrationDefinitionPropertiesResponseInput interface {
	pulumi.Input

	ToRegistrationDefinitionPropertiesResponseOutput() RegistrationDefinitionPropertiesResponseOutput
	ToRegistrationDefinitionPropertiesResponseOutputWithContext(context.Context) RegistrationDefinitionPropertiesResponseOutput
}

type RegistrationDefinitionPropertiesResponseArgs struct {
	Authorizations             AuthorizationResponseArrayInput         `pulumi:"authorizations"`
	Description                pulumi.StringPtrInput                   `pulumi:"description"`
	EligibleAuthorizations     EligibleAuthorizationResponseArrayInput `pulumi:"eligibleAuthorizations"`
	ManagedByTenantId          pulumi.StringInput                      `pulumi:"managedByTenantId"`
	ManagedByTenantName        pulumi.StringInput                      `pulumi:"managedByTenantName"`
	ProvisioningState          pulumi.StringInput                      `pulumi:"provisioningState"`
	RegistrationDefinitionName pulumi.StringPtrInput                   `pulumi:"registrationDefinitionName"`
}

func (RegistrationDefinitionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationDefinitionPropertiesResponse)(nil)).Elem()
}

func (i RegistrationDefinitionPropertiesResponseArgs) ToRegistrationDefinitionPropertiesResponseOutput() RegistrationDefinitionPropertiesResponseOutput {
	return i.ToRegistrationDefinitionPropertiesResponseOutputWithContext(context.Background())
}

func (i RegistrationDefinitionPropertiesResponseArgs) ToRegistrationDefinitionPropertiesResponseOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationDefinitionPropertiesResponseOutput)
}

func (i RegistrationDefinitionPropertiesResponseArgs) ToRegistrationDefinitionPropertiesResponsePtrOutput() RegistrationDefinitionPropertiesResponsePtrOutput {
	return i.ToRegistrationDefinitionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i RegistrationDefinitionPropertiesResponseArgs) ToRegistrationDefinitionPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationDefinitionPropertiesResponseOutput).ToRegistrationDefinitionPropertiesResponsePtrOutputWithContext(ctx)
}









type RegistrationDefinitionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToRegistrationDefinitionPropertiesResponsePtrOutput() RegistrationDefinitionPropertiesResponsePtrOutput
	ToRegistrationDefinitionPropertiesResponsePtrOutputWithContext(context.Context) RegistrationDefinitionPropertiesResponsePtrOutput
}

type registrationDefinitionPropertiesResponsePtrType RegistrationDefinitionPropertiesResponseArgs

func RegistrationDefinitionPropertiesResponsePtr(v *RegistrationDefinitionPropertiesResponseArgs) RegistrationDefinitionPropertiesResponsePtrInput {
	return (*registrationDefinitionPropertiesResponsePtrType)(v)
}

func (*registrationDefinitionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationDefinitionPropertiesResponse)(nil)).Elem()
}

func (i *registrationDefinitionPropertiesResponsePtrType) ToRegistrationDefinitionPropertiesResponsePtrOutput() RegistrationDefinitionPropertiesResponsePtrOutput {
	return i.ToRegistrationDefinitionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *registrationDefinitionPropertiesResponsePtrType) ToRegistrationDefinitionPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationDefinitionPropertiesResponsePtrOutput)
}

type RegistrationDefinitionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RegistrationDefinitionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationDefinitionPropertiesResponse)(nil)).Elem()
}

func (o RegistrationDefinitionPropertiesResponseOutput) ToRegistrationDefinitionPropertiesResponseOutput() RegistrationDefinitionPropertiesResponseOutput {
	return o
}

func (o RegistrationDefinitionPropertiesResponseOutput) ToRegistrationDefinitionPropertiesResponseOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesResponseOutput {
	return o
}

func (o RegistrationDefinitionPropertiesResponseOutput) ToRegistrationDefinitionPropertiesResponsePtrOutput() RegistrationDefinitionPropertiesResponsePtrOutput {
	return o.ToRegistrationDefinitionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o RegistrationDefinitionPropertiesResponseOutput) ToRegistrationDefinitionPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationDefinitionPropertiesResponse) *RegistrationDefinitionPropertiesResponse {
		return &v
	}).(RegistrationDefinitionPropertiesResponsePtrOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) Authorizations() AuthorizationResponseArrayOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) []AuthorizationResponse { return v.Authorizations }).(AuthorizationResponseArrayOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) EligibleAuthorizations() EligibleAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) []EligibleAuthorizationResponse {
		return v.EligibleAuthorizations
	}).(EligibleAuthorizationResponseArrayOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) ManagedByTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) string { return v.ManagedByTenantId }).(pulumi.StringOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) ManagedByTenantName() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) string { return v.ManagedByTenantName }).(pulumi.StringOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) RegistrationDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) *string { return v.RegistrationDefinitionName }).(pulumi.StringPtrOutput)
}

type RegistrationDefinitionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (RegistrationDefinitionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationDefinitionPropertiesResponse)(nil)).Elem()
}

func (o RegistrationDefinitionPropertiesResponsePtrOutput) ToRegistrationDefinitionPropertiesResponsePtrOutput() RegistrationDefinitionPropertiesResponsePtrOutput {
	return o
}

func (o RegistrationDefinitionPropertiesResponsePtrOutput) ToRegistrationDefinitionPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesResponsePtrOutput {
	return o
}

func (o RegistrationDefinitionPropertiesResponsePtrOutput) Elem() RegistrationDefinitionPropertiesResponseOutput {
	return o.ApplyT(func(v *RegistrationDefinitionPropertiesResponse) RegistrationDefinitionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret RegistrationDefinitionPropertiesResponse
		return ret
	}).(RegistrationDefinitionPropertiesResponseOutput)
}

func (o RegistrationDefinitionPropertiesResponsePtrOutput) Authorizations() AuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *RegistrationDefinitionPropertiesResponse) []AuthorizationResponse {
		if v == nil {
			return nil
		}
		return v.Authorizations
	}).(AuthorizationResponseArrayOutput)
}

func (o RegistrationDefinitionPropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationDefinitionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesResponsePtrOutput) EligibleAuthorizations() EligibleAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *RegistrationDefinitionPropertiesResponse) []EligibleAuthorizationResponse {
		if v == nil {
			return nil
		}
		return v.EligibleAuthorizations
	}).(EligibleAuthorizationResponseArrayOutput)
}

func (o RegistrationDefinitionPropertiesResponsePtrOutput) ManagedByTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationDefinitionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ManagedByTenantId
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesResponsePtrOutput) ManagedByTenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationDefinitionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ManagedByTenantName
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationDefinitionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesResponsePtrOutput) RegistrationDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationDefinitionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationDefinitionName
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationOutput{})
	pulumi.RegisterOutputType(AuthorizationArrayOutput{})
	pulumi.RegisterOutputType(AuthorizationResponseOutput{})
	pulumi.RegisterOutputType(AuthorizationResponseArrayOutput{})
	pulumi.RegisterOutputType(EligibleApproverOutput{})
	pulumi.RegisterOutputType(EligibleApproverArrayOutput{})
	pulumi.RegisterOutputType(EligibleApproverResponseOutput{})
	pulumi.RegisterOutputType(EligibleApproverResponseArrayOutput{})
	pulumi.RegisterOutputType(EligibleAuthorizationOutput{})
	pulumi.RegisterOutputType(EligibleAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(EligibleAuthorizationResponseOutput{})
	pulumi.RegisterOutputType(EligibleAuthorizationResponseArrayOutput{})
	pulumi.RegisterOutputType(JustInTimeAccessPolicyOutput{})
	pulumi.RegisterOutputType(JustInTimeAccessPolicyPtrOutput{})
	pulumi.RegisterOutputType(JustInTimeAccessPolicyResponseOutput{})
	pulumi.RegisterOutputType(JustInTimeAccessPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponsePtrOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesResponsePropertiesOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesResponseRegistrationDefinitionPtrOutput{})
	pulumi.RegisterOutputType(RegistrationDefinitionPropertiesOutput{})
	pulumi.RegisterOutputType(RegistrationDefinitionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(RegistrationDefinitionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RegistrationDefinitionPropertiesResponsePtrOutput{})
}
