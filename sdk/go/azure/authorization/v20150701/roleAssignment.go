


package v20150701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleAssignment struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                             `pulumi:"name"`
	Properties RoleAssignmentPropertiesWithScopeResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                             `pulumi:"type"`
}


func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:RoleAssignment"),
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
	})
	opts = append(opts, aliases)
	var resource RoleAssignment
	err := ctx.RegisterResource("azure-native:authorization/v20150701:RoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentState, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	var resource RoleAssignment
	err := ctx.ReadResource("azure-native:authorization/v20150701:RoleAssignment", name, id, state, &resource, opts...)
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
	Properties         RoleAssignmentProperties `pulumi:"properties"`
	RoleAssignmentName *string                  `pulumi:"roleAssignmentName"`
	Scope              string                   `pulumi:"scope"`
}


type RoleAssignmentArgs struct {
	Properties         RoleAssignmentPropertiesInput
	RoleAssignmentName pulumi.StringPtrInput
	Scope              pulumi.StringInput
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

func init() {
	pulumi.RegisterOutputType(RoleAssignmentOutput{})
}
