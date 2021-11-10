


package v20201001preview

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
	if args.PrincipalType == nil {
		args.PrincipalType = pulumi.StringPtr("User")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:RoleAssignment"),
		},
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
	})
	opts = append(opts, aliases)
	var resource RoleAssignment
	err := ctx.RegisterResource("azure-native:authorization/v20201001preview:RoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentState, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	var resource RoleAssignment
	err := ctx.ReadResource("azure-native:authorization/v20201001preview:RoleAssignment", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*RoleAssignment)(nil))
}

func (i *RoleAssignment) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return i.ToRoleAssignmentOutputWithContext(context.Background())
}

func (i *RoleAssignment) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentOutput)
}

type RoleAssignmentOutput struct{ *pulumi.OutputState }

func (RoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssignment)(nil))
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return o
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RoleAssignmentOutput{})
}
