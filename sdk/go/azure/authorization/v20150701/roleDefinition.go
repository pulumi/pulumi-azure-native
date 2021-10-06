


package v20150701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleDefinition struct {
	pulumi.CustomResourceState

	AssignableScopes pulumi.StringArrayOutput      `pulumi:"assignableScopes"`
	Description      pulumi.StringPtrOutput        `pulumi:"description"`
	Name             pulumi.StringOutput           `pulumi:"name"`
	Permissions      PermissionResponseArrayOutput `pulumi:"permissions"`
	RoleName         pulumi.StringPtrOutput        `pulumi:"roleName"`
	RoleType         pulumi.StringPtrOutput        `pulumi:"roleType"`
	Type             pulumi.StringOutput           `pulumi:"type"`
}


func NewRoleDefinition(ctx *pulumi.Context,
	name string, args *RoleDefinitionArgs, opts ...pulumi.ResourceOption) (*RoleDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:authorization/v20150701:RoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization:RoleDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization:RoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180101preview:RoleDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20180101preview:RoleDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource RoleDefinition
	err := ctx.RegisterResource("azure-native:authorization/v20150701:RoleDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRoleDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleDefinitionState, opts ...pulumi.ResourceOption) (*RoleDefinition, error) {
	var resource RoleDefinition
	err := ctx.ReadResource("azure-native:authorization/v20150701:RoleDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type roleDefinitionState struct {
}

type RoleDefinitionState struct {
}

func (RoleDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleDefinitionState)(nil)).Elem()
}

type roleDefinitionArgs struct {
	AssignableScopes []string     `pulumi:"assignableScopes"`
	Description      *string      `pulumi:"description"`
	Permissions      []Permission `pulumi:"permissions"`
	RoleDefinitionId *string      `pulumi:"roleDefinitionId"`
	RoleName         *string      `pulumi:"roleName"`
	RoleType         *string      `pulumi:"roleType"`
	Scope            string       `pulumi:"scope"`
}


type RoleDefinitionArgs struct {
	AssignableScopes pulumi.StringArrayInput
	Description      pulumi.StringPtrInput
	Permissions      PermissionArrayInput
	RoleDefinitionId pulumi.StringPtrInput
	RoleName         pulumi.StringPtrInput
	RoleType         pulumi.StringPtrInput
	Scope            pulumi.StringInput
}

func (RoleDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleDefinitionArgs)(nil)).Elem()
}

type RoleDefinitionInput interface {
	pulumi.Input

	ToRoleDefinitionOutput() RoleDefinitionOutput
	ToRoleDefinitionOutputWithContext(ctx context.Context) RoleDefinitionOutput
}

func (*RoleDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleDefinition)(nil))
}

func (i *RoleDefinition) ToRoleDefinitionOutput() RoleDefinitionOutput {
	return i.ToRoleDefinitionOutputWithContext(context.Background())
}

func (i *RoleDefinition) ToRoleDefinitionOutputWithContext(ctx context.Context) RoleDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleDefinitionOutput)
}

type RoleDefinitionOutput struct{ *pulumi.OutputState }

func (RoleDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleDefinition)(nil))
}

func (o RoleDefinitionOutput) ToRoleDefinitionOutput() RoleDefinitionOutput {
	return o
}

func (o RoleDefinitionOutput) ToRoleDefinitionOutputWithContext(ctx context.Context) RoleDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RoleDefinitionOutput{})
}
