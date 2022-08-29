


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleManagementPolicyAssignment struct {
	pulumi.CustomResourceState

	Name                       pulumi.StringOutput                      `pulumi:"name"`
	PolicyAssignmentProperties PolicyAssignmentPropertiesResponseOutput `pulumi:"policyAssignmentProperties"`
	PolicyId                   pulumi.StringPtrOutput                   `pulumi:"policyId"`
	RoleDefinitionId           pulumi.StringPtrOutput                   `pulumi:"roleDefinitionId"`
	Scope                      pulumi.StringPtrOutput                   `pulumi:"scope"`
	Type                       pulumi.StringOutput                      `pulumi:"type"`
}


func NewRoleManagementPolicyAssignment(ctx *pulumi.Context,
	name string, args *RoleManagementPolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleManagementPolicyAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:RoleManagementPolicyAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20201001:RoleManagementPolicyAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource RoleManagementPolicyAssignment
	err := ctx.RegisterResource("azure-native:authorization/v20201001preview:RoleManagementPolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRoleManagementPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleManagementPolicyAssignmentState, opts ...pulumi.ResourceOption) (*RoleManagementPolicyAssignment, error) {
	var resource RoleManagementPolicyAssignment
	err := ctx.ReadResource("azure-native:authorization/v20201001preview:RoleManagementPolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type roleManagementPolicyAssignmentState struct {
}

type RoleManagementPolicyAssignmentState struct {
}

func (RoleManagementPolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleManagementPolicyAssignmentState)(nil)).Elem()
}

type roleManagementPolicyAssignmentArgs struct {
	PolicyId                           *string `pulumi:"policyId"`
	RoleDefinitionId                   *string `pulumi:"roleDefinitionId"`
	RoleManagementPolicyAssignmentName *string `pulumi:"roleManagementPolicyAssignmentName"`
	Scope                              string  `pulumi:"scope"`
}


type RoleManagementPolicyAssignmentArgs struct {
	PolicyId                           pulumi.StringPtrInput
	RoleDefinitionId                   pulumi.StringPtrInput
	RoleManagementPolicyAssignmentName pulumi.StringPtrInput
	Scope                              pulumi.StringInput
}

func (RoleManagementPolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleManagementPolicyAssignmentArgs)(nil)).Elem()
}

type RoleManagementPolicyAssignmentInput interface {
	pulumi.Input

	ToRoleManagementPolicyAssignmentOutput() RoleManagementPolicyAssignmentOutput
	ToRoleManagementPolicyAssignmentOutputWithContext(ctx context.Context) RoleManagementPolicyAssignmentOutput
}

func (*RoleManagementPolicyAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleManagementPolicyAssignment)(nil)).Elem()
}

func (i *RoleManagementPolicyAssignment) ToRoleManagementPolicyAssignmentOutput() RoleManagementPolicyAssignmentOutput {
	return i.ToRoleManagementPolicyAssignmentOutputWithContext(context.Background())
}

func (i *RoleManagementPolicyAssignment) ToRoleManagementPolicyAssignmentOutputWithContext(ctx context.Context) RoleManagementPolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleManagementPolicyAssignmentOutput)
}

type RoleManagementPolicyAssignmentOutput struct{ *pulumi.OutputState }

func (RoleManagementPolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleManagementPolicyAssignment)(nil)).Elem()
}

func (o RoleManagementPolicyAssignmentOutput) ToRoleManagementPolicyAssignmentOutput() RoleManagementPolicyAssignmentOutput {
	return o
}

func (o RoleManagementPolicyAssignmentOutput) ToRoleManagementPolicyAssignmentOutputWithContext(ctx context.Context) RoleManagementPolicyAssignmentOutput {
	return o
}

func (o RoleManagementPolicyAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RoleManagementPolicyAssignmentOutput) PolicyAssignmentProperties() PolicyAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) PolicyAssignmentPropertiesResponseOutput {
		return v.PolicyAssignmentProperties
	}).(PolicyAssignmentPropertiesResponseOutput)
}

func (o RoleManagementPolicyAssignmentOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) pulumi.StringPtrOutput { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o RoleManagementPolicyAssignmentOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) pulumi.StringPtrOutput { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o RoleManagementPolicyAssignmentOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o RoleManagementPolicyAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleManagementPolicyAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RoleManagementPolicyAssignmentOutput{})
}
