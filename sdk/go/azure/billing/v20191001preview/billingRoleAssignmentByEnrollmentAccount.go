


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BillingRoleAssignmentByEnrollmentAccount struct {
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


func NewBillingRoleAssignmentByEnrollmentAccount(ctx *pulumi.Context,
	name string, args *BillingRoleAssignmentByEnrollmentAccountArgs, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByEnrollmentAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountName == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountName'")
	}
	if args.EnrollmentAccountName == nil {
		return nil, errors.New("invalid value for required argument 'EnrollmentAccountName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:billing:BillingRoleAssignmentByEnrollmentAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource BillingRoleAssignmentByEnrollmentAccount
	err := ctx.RegisterResource("azure-native:billing/v20191001preview:BillingRoleAssignmentByEnrollmentAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBillingRoleAssignmentByEnrollmentAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingRoleAssignmentByEnrollmentAccountState, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByEnrollmentAccount, error) {
	var resource BillingRoleAssignmentByEnrollmentAccount
	err := ctx.ReadResource("azure-native:billing/v20191001preview:BillingRoleAssignmentByEnrollmentAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type billingRoleAssignmentByEnrollmentAccountState struct {
}

type BillingRoleAssignmentByEnrollmentAccountState struct {
}

func (BillingRoleAssignmentByEnrollmentAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByEnrollmentAccountState)(nil)).Elem()
}

type billingRoleAssignmentByEnrollmentAccountArgs struct {
	BillingAccountName        string  `pulumi:"billingAccountName"`
	BillingRoleAssignmentName *string `pulumi:"billingRoleAssignmentName"`
	EnrollmentAccountName     string  `pulumi:"enrollmentAccountName"`
	PrincipalId               *string `pulumi:"principalId"`
	PrincipalTenantId         *string `pulumi:"principalTenantId"`
	RoleDefinitionId          *string `pulumi:"roleDefinitionId"`
	UserAuthenticationType    *string `pulumi:"userAuthenticationType"`
	UserEmailAddress          *string `pulumi:"userEmailAddress"`
}


type BillingRoleAssignmentByEnrollmentAccountArgs struct {
	BillingAccountName        pulumi.StringInput
	BillingRoleAssignmentName pulumi.StringPtrInput
	EnrollmentAccountName     pulumi.StringInput
	PrincipalId               pulumi.StringPtrInput
	PrincipalTenantId         pulumi.StringPtrInput
	RoleDefinitionId          pulumi.StringPtrInput
	UserAuthenticationType    pulumi.StringPtrInput
	UserEmailAddress          pulumi.StringPtrInput
}

func (BillingRoleAssignmentByEnrollmentAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByEnrollmentAccountArgs)(nil)).Elem()
}

type BillingRoleAssignmentByEnrollmentAccountInput interface {
	pulumi.Input

	ToBillingRoleAssignmentByEnrollmentAccountOutput() BillingRoleAssignmentByEnrollmentAccountOutput
	ToBillingRoleAssignmentByEnrollmentAccountOutputWithContext(ctx context.Context) BillingRoleAssignmentByEnrollmentAccountOutput
}

func (*BillingRoleAssignmentByEnrollmentAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingRoleAssignmentByEnrollmentAccount)(nil)).Elem()
}

func (i *BillingRoleAssignmentByEnrollmentAccount) ToBillingRoleAssignmentByEnrollmentAccountOutput() BillingRoleAssignmentByEnrollmentAccountOutput {
	return i.ToBillingRoleAssignmentByEnrollmentAccountOutputWithContext(context.Background())
}

func (i *BillingRoleAssignmentByEnrollmentAccount) ToBillingRoleAssignmentByEnrollmentAccountOutputWithContext(ctx context.Context) BillingRoleAssignmentByEnrollmentAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingRoleAssignmentByEnrollmentAccountOutput)
}

type BillingRoleAssignmentByEnrollmentAccountOutput struct{ *pulumi.OutputState }

func (BillingRoleAssignmentByEnrollmentAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingRoleAssignmentByEnrollmentAccount)(nil)).Elem()
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) ToBillingRoleAssignmentByEnrollmentAccountOutput() BillingRoleAssignmentByEnrollmentAccountOutput {
	return o
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) ToBillingRoleAssignmentByEnrollmentAccountOutputWithContext(ctx context.Context) BillingRoleAssignmentByEnrollmentAccountOutput {
	return o
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) CreatedByPrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringOutput { return v.CreatedByPrincipalId }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) CreatedByPrincipalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringOutput {
		return v.CreatedByPrincipalTenantId
	}).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) CreatedByUserEmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringOutput {
		return v.CreatedByUserEmailAddress
	}).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringPtrOutput { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) PrincipalTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringPtrOutput { return v.PrincipalTenantId }).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringPtrOutput { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) UserAuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringPtrOutput {
		return v.UserAuthenticationType
	}).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByEnrollmentAccountOutput) UserEmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByEnrollmentAccount) pulumi.StringPtrOutput { return v.UserEmailAddress }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BillingRoleAssignmentByEnrollmentAccountOutput{})
}
