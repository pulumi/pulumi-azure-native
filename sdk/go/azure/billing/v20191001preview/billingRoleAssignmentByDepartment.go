


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BillingRoleAssignmentByDepartment struct {
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


func NewBillingRoleAssignmentByDepartment(ctx *pulumi.Context,
	name string, args *BillingRoleAssignmentByDepartmentArgs, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByDepartment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountName == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountName'")
	}
	if args.DepartmentName == nil {
		return nil, errors.New("invalid value for required argument 'DepartmentName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:billing:BillingRoleAssignmentByDepartment"),
		},
	})
	opts = append(opts, aliases)
	var resource BillingRoleAssignmentByDepartment
	err := ctx.RegisterResource("azure-native:billing/v20191001preview:BillingRoleAssignmentByDepartment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBillingRoleAssignmentByDepartment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingRoleAssignmentByDepartmentState, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByDepartment, error) {
	var resource BillingRoleAssignmentByDepartment
	err := ctx.ReadResource("azure-native:billing/v20191001preview:BillingRoleAssignmentByDepartment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type billingRoleAssignmentByDepartmentState struct {
}

type BillingRoleAssignmentByDepartmentState struct {
}

func (BillingRoleAssignmentByDepartmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByDepartmentState)(nil)).Elem()
}

type billingRoleAssignmentByDepartmentArgs struct {
	BillingAccountName        string  `pulumi:"billingAccountName"`
	BillingRoleAssignmentName *string `pulumi:"billingRoleAssignmentName"`
	DepartmentName            string  `pulumi:"departmentName"`
	PrincipalId               *string `pulumi:"principalId"`
	PrincipalTenantId         *string `pulumi:"principalTenantId"`
	RoleDefinitionId          *string `pulumi:"roleDefinitionId"`
	UserAuthenticationType    *string `pulumi:"userAuthenticationType"`
	UserEmailAddress          *string `pulumi:"userEmailAddress"`
}


type BillingRoleAssignmentByDepartmentArgs struct {
	BillingAccountName        pulumi.StringInput
	BillingRoleAssignmentName pulumi.StringPtrInput
	DepartmentName            pulumi.StringInput
	PrincipalId               pulumi.StringPtrInput
	PrincipalTenantId         pulumi.StringPtrInput
	RoleDefinitionId          pulumi.StringPtrInput
	UserAuthenticationType    pulumi.StringPtrInput
	UserEmailAddress          pulumi.StringPtrInput
}

func (BillingRoleAssignmentByDepartmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByDepartmentArgs)(nil)).Elem()
}

type BillingRoleAssignmentByDepartmentInput interface {
	pulumi.Input

	ToBillingRoleAssignmentByDepartmentOutput() BillingRoleAssignmentByDepartmentOutput
	ToBillingRoleAssignmentByDepartmentOutputWithContext(ctx context.Context) BillingRoleAssignmentByDepartmentOutput
}

func (*BillingRoleAssignmentByDepartment) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingRoleAssignmentByDepartment)(nil)).Elem()
}

func (i *BillingRoleAssignmentByDepartment) ToBillingRoleAssignmentByDepartmentOutput() BillingRoleAssignmentByDepartmentOutput {
	return i.ToBillingRoleAssignmentByDepartmentOutputWithContext(context.Background())
}

func (i *BillingRoleAssignmentByDepartment) ToBillingRoleAssignmentByDepartmentOutputWithContext(ctx context.Context) BillingRoleAssignmentByDepartmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingRoleAssignmentByDepartmentOutput)
}

type BillingRoleAssignmentByDepartmentOutput struct{ *pulumi.OutputState }

func (BillingRoleAssignmentByDepartmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingRoleAssignmentByDepartment)(nil)).Elem()
}

func (o BillingRoleAssignmentByDepartmentOutput) ToBillingRoleAssignmentByDepartmentOutput() BillingRoleAssignmentByDepartmentOutput {
	return o
}

func (o BillingRoleAssignmentByDepartmentOutput) ToBillingRoleAssignmentByDepartmentOutputWithContext(ctx context.Context) BillingRoleAssignmentByDepartmentOutput {
	return o
}

func (o BillingRoleAssignmentByDepartmentOutput) CreatedByPrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringOutput { return v.CreatedByPrincipalId }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByDepartmentOutput) CreatedByPrincipalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringOutput { return v.CreatedByPrincipalTenantId }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByDepartmentOutput) CreatedByUserEmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringOutput { return v.CreatedByUserEmailAddress }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByDepartmentOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByDepartmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByDepartmentOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringPtrOutput { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByDepartmentOutput) PrincipalTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringPtrOutput { return v.PrincipalTenantId }).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByDepartmentOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringPtrOutput { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByDepartmentOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByDepartmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o BillingRoleAssignmentByDepartmentOutput) UserAuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringPtrOutput { return v.UserAuthenticationType }).(pulumi.StringPtrOutput)
}

func (o BillingRoleAssignmentByDepartmentOutput) UserEmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByDepartment) pulumi.StringPtrOutput { return v.UserEmailAddress }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BillingRoleAssignmentByDepartmentOutput{})
}
