


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BillingRoleAssignmentByBillingAccount struct {
	pulumi.CustomResourceState

	CreatedByPrincipalId       pulumi.StringOutput    `pulumi:"createdByPrincipalId"`
	CreatedByPrincipalTenantId pulumi.StringOutput    `pulumi:"createdByPrincipalTenantId"`
	CreatedByUserEmailAddress  pulumi.StringOutput    `pulumi:"createdByUserEmailAddress"`
	CreatedOn                  pulumi.StringOutput    `pulumi:"createdOn"`
	Name                       pulumi.StringOutput    `pulumi:"name"`
	PrincipalId                pulumi.StringPtrOutput `pulumi:"principalId"`
	PrincipalTenantId          pulumi.StringPtrOutput `pulumi:"principalTenantId"`
	RoleDefinitionId           pulumi.StringPtrOutput `pulumi:"roleDefinitionId"`
	Scope                      pulumi.StringOutput    `pulumi:"scope"`
	Type                       pulumi.StringOutput    `pulumi:"type"`
	UserAuthenticationType     pulumi.StringPtrOutput `pulumi:"userAuthenticationType"`
	UserEmailAddress           pulumi.StringPtrOutput `pulumi:"userEmailAddress"`
}


func NewBillingRoleAssignmentByBillingAccount(ctx *pulumi.Context,
	name string, args *BillingRoleAssignmentByBillingAccountArgs, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByBillingAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountName == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:billing:BillingRoleAssignmentByBillingAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource BillingRoleAssignmentByBillingAccount
	err := ctx.RegisterResource("azure-native:billing/v20191001preview:BillingRoleAssignmentByBillingAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBillingRoleAssignmentByBillingAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingRoleAssignmentByBillingAccountState, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByBillingAccount, error) {
	var resource BillingRoleAssignmentByBillingAccount
	err := ctx.ReadResource("azure-native:billing/v20191001preview:BillingRoleAssignmentByBillingAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type billingRoleAssignmentByBillingAccountState struct {
}

type BillingRoleAssignmentByBillingAccountState struct {
}

func (BillingRoleAssignmentByBillingAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByBillingAccountState)(nil)).Elem()
}

type billingRoleAssignmentByBillingAccountArgs struct {
	BillingAccountName        string  `pulumi:"billingAccountName"`
	BillingRoleAssignmentName *string `pulumi:"billingRoleAssignmentName"`
	PrincipalId               *string `pulumi:"principalId"`
	PrincipalTenantId         *string `pulumi:"principalTenantId"`
	RoleDefinitionId          *string `pulumi:"roleDefinitionId"`
	UserAuthenticationType    *string `pulumi:"userAuthenticationType"`
	UserEmailAddress          *string `pulumi:"userEmailAddress"`
}


type BillingRoleAssignmentByBillingAccountArgs struct {
	BillingAccountName        pulumi.StringInput
	BillingRoleAssignmentName pulumi.StringPtrInput
	PrincipalId               pulumi.StringPtrInput
	PrincipalTenantId         pulumi.StringPtrInput
	RoleDefinitionId          pulumi.StringPtrInput
	UserAuthenticationType    pulumi.StringPtrInput
	UserEmailAddress          pulumi.StringPtrInput
}

func (BillingRoleAssignmentByBillingAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByBillingAccountArgs)(nil)).Elem()
}

type BillingRoleAssignmentByBillingAccountInput interface {
	pulumi.Input

	ToBillingRoleAssignmentByBillingAccountOutput() BillingRoleAssignmentByBillingAccountOutput
	ToBillingRoleAssignmentByBillingAccountOutputWithContext(ctx context.Context) BillingRoleAssignmentByBillingAccountOutput
}

func (*BillingRoleAssignmentByBillingAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingRoleAssignmentByBillingAccount)(nil)).Elem()
}

func (i *BillingRoleAssignmentByBillingAccount) ToBillingRoleAssignmentByBillingAccountOutput() BillingRoleAssignmentByBillingAccountOutput {
	return i.ToBillingRoleAssignmentByBillingAccountOutputWithContext(context.Background())
}

func (i *BillingRoleAssignmentByBillingAccount) ToBillingRoleAssignmentByBillingAccountOutputWithContext(ctx context.Context) BillingRoleAssignmentByBillingAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingRoleAssignmentByBillingAccountOutput)
}

type BillingRoleAssignmentByBillingAccountOutput struct{ *pulumi.OutputState }

func (BillingRoleAssignmentByBillingAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingRoleAssignmentByBillingAccount)(nil)).Elem()
}

func (o BillingRoleAssignmentByBillingAccountOutput) ToBillingRoleAssignmentByBillingAccountOutput() BillingRoleAssignmentByBillingAccountOutput {
	return o
}

func (o BillingRoleAssignmentByBillingAccountOutput) ToBillingRoleAssignmentByBillingAccountOutputWithContext(ctx context.Context) BillingRoleAssignmentByBillingAccountOutput {
	return o
}

func (o BillingRoleAssignmentByBillingAccountOutput) CreatedByPrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.CreatedByPrincipalId }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByBillingAccountOutput) CreatedByPrincipalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput {
		return v.CreatedByPrincipalTenantId
	}).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByBillingAccountOutput) CreatedByUserEmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.CreatedByUserEmailAddress }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByBillingAccountOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByBillingAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByBillingAccountOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringPtrOutput { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByBillingAccountOutput) PrincipalTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringPtrOutput { return v.PrincipalTenantId }).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByBillingAccountOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringPtrOutput { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByBillingAccountOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByBillingAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByBillingAccountOutput) UserAuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringPtrOutput { return v.UserAuthenticationType }).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByBillingAccountOutput) UserEmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringPtrOutput { return v.UserEmailAddress }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BillingRoleAssignmentByBillingAccountOutput{})
}
