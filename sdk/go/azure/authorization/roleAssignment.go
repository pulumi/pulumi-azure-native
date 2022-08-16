


package authorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleAssignment struct {
	pulumi.CustomResourceState

	Condition                          pulumi.StringPtrOutput `pulumi:"condition"`
	ConditionVersion                   pulumi.StringPtrOutput `pulumi:"conditionVersion"`
	CreatedBy                          pulumi.StringOutput    `pulumi:"createdBy"`
	CreatedOn                          pulumi.StringOutput    `pulumi:"createdOn"`
	DelegatedManagedIdentityResourceId pulumi.StringPtrOutput `pulumi:"delegatedManagedIdentityResourceId"`
	Description                        pulumi.StringPtrOutput `pulumi:"description"`
	Name                               pulumi.StringOutput    `pulumi:"name"`
	PrincipalId                        pulumi.StringOutput    `pulumi:"principalId"`
	PrincipalType                      pulumi.StringPtrOutput `pulumi:"principalType"`
	RoleDefinitionId                   pulumi.StringOutput    `pulumi:"roleDefinitionId"`
	Scope                              pulumi.StringOutput    `pulumi:"scope"`
	Type                               pulumi.StringOutput    `pulumi:"type"`
	UpdatedBy                          pulumi.StringOutput    `pulumi:"updatedBy"`
	UpdatedOn                          pulumi.StringOutput    `pulumi:"updatedOn"`
}


func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.RoleDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'RoleDefinitionId'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if isZero(args.PrincipalType) {
		args.PrincipalType = pulumi.StringPtr("User")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20150701:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20171001preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180101preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180901preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200301preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200401preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200801preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20201001preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20220401:RoleAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource RoleAssignment
	err := ctx.RegisterResource("azure-native:authorization:RoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentState, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	var resource RoleAssignment
	err := ctx.ReadResource("azure-native:authorization:RoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type roleAssignmentState struct {
}

type RoleAssignmentState struct {
}

func (RoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentState)(nil)).Elem()
}

type roleAssignmentArgs struct {
	Condition                          *string `pulumi:"condition"`
	ConditionVersion                   *string `pulumi:"conditionVersion"`
	DelegatedManagedIdentityResourceId *string `pulumi:"delegatedManagedIdentityResourceId"`
	Description                        *string `pulumi:"description"`
	PrincipalId                        string  `pulumi:"principalId"`
	PrincipalType                      *string `pulumi:"principalType"`
	RoleAssignmentName                 *string `pulumi:"roleAssignmentName"`
	RoleDefinitionId                   string  `pulumi:"roleDefinitionId"`
	Scope                              string  `pulumi:"scope"`
}


type RoleAssignmentArgs struct {
	Condition                          pulumi.StringPtrInput
	ConditionVersion                   pulumi.StringPtrInput
	DelegatedManagedIdentityResourceId pulumi.StringPtrInput
	Description                        pulumi.StringPtrInput
	PrincipalId                        pulumi.StringInput
	PrincipalType                      pulumi.StringPtrInput
	RoleAssignmentName                 pulumi.StringPtrInput
	RoleDefinitionId                   pulumi.StringInput
	Scope                              pulumi.StringInput
}

func (RoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArgs)(nil)).Elem()
}

type RoleAssignmentInput interface {
	pulumi.Input

	ToRoleAssignmentOutput() RoleAssignmentOutput
	ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput
}

func (*RoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil)).Elem()
}

func (i *RoleAssignment) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return i.ToRoleAssignmentOutputWithContext(context.Background())
}

func (i *RoleAssignment) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentOutput)
}

type RoleAssignmentOutput struct{ *pulumi.OutputState }

func (RoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil)).Elem()
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return o
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return o
}

func (o RoleAssignmentOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.Condition }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentOutput) ConditionVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.ConditionVersion }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) DelegatedManagedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.DelegatedManagedIdentityResourceId }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) UpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.UpdatedBy }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) UpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.UpdatedOn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RoleAssignmentOutput{})
}
