


package documentdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlResourceSqlRoleDefinition struct {
	pulumi.CustomResourceState

	AssignableScopes pulumi.StringArrayOutput      `pulumi:"assignableScopes"`
	Name             pulumi.StringOutput           `pulumi:"name"`
	Permissions      PermissionResponseArrayOutput `pulumi:"permissions"`
	RoleName         pulumi.StringPtrOutput        `pulumi:"roleName"`
	Type             pulumi.StringOutput           `pulumi:"type"`
}


func NewSqlResourceSqlRoleDefinition(ctx *pulumi.Context,
	name string, args *SqlResourceSqlRoleDefinitionArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlRoleDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:documentdb:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200601preview:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210301preview:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210401preview:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210415:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210515:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210615:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210701preview:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:SqlResourceSqlRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20211015:SqlResourceSqlRoleDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlResourceSqlRoleDefinition
	err := ctx.RegisterResource("azure-native:documentdb:SqlResourceSqlRoleDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlResourceSqlRoleDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlRoleDefinitionState, opts ...pulumi.ResourceOption) (*SqlResourceSqlRoleDefinition, error) {
	var resource SqlResourceSqlRoleDefinition
	err := ctx.ReadResource("azure-native:documentdb:SqlResourceSqlRoleDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlResourceSqlRoleDefinitionState struct {
}

type SqlResourceSqlRoleDefinitionState struct {
}

func (SqlResourceSqlRoleDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlRoleDefinitionState)(nil)).Elem()
}

type sqlResourceSqlRoleDefinitionArgs struct {
	AccountName       string              `pulumi:"accountName"`
	AssignableScopes  []string            `pulumi:"assignableScopes"`
	Permissions       []Permission        `pulumi:"permissions"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	RoleDefinitionId  *string             `pulumi:"roleDefinitionId"`
	RoleName          *string             `pulumi:"roleName"`
	Type              *RoleDefinitionType `pulumi:"type"`
}


type SqlResourceSqlRoleDefinitionArgs struct {
	AccountName       pulumi.StringInput
	AssignableScopes  pulumi.StringArrayInput
	Permissions       PermissionArrayInput
	ResourceGroupName pulumi.StringInput
	RoleDefinitionId  pulumi.StringPtrInput
	RoleName          pulumi.StringPtrInput
	Type              RoleDefinitionTypePtrInput
}

func (SqlResourceSqlRoleDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlRoleDefinitionArgs)(nil)).Elem()
}

type SqlResourceSqlRoleDefinitionInput interface {
	pulumi.Input

	ToSqlResourceSqlRoleDefinitionOutput() SqlResourceSqlRoleDefinitionOutput
	ToSqlResourceSqlRoleDefinitionOutputWithContext(ctx context.Context) SqlResourceSqlRoleDefinitionOutput
}

func (*SqlResourceSqlRoleDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlResourceSqlRoleDefinition)(nil))
}

func (i *SqlResourceSqlRoleDefinition) ToSqlResourceSqlRoleDefinitionOutput() SqlResourceSqlRoleDefinitionOutput {
	return i.ToSqlResourceSqlRoleDefinitionOutputWithContext(context.Background())
}

func (i *SqlResourceSqlRoleDefinition) ToSqlResourceSqlRoleDefinitionOutputWithContext(ctx context.Context) SqlResourceSqlRoleDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlRoleDefinitionOutput)
}

type SqlResourceSqlRoleDefinitionOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlRoleDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlResourceSqlRoleDefinition)(nil))
}

func (o SqlResourceSqlRoleDefinitionOutput) ToSqlResourceSqlRoleDefinitionOutput() SqlResourceSqlRoleDefinitionOutput {
	return o
}

func (o SqlResourceSqlRoleDefinitionOutput) ToSqlResourceSqlRoleDefinitionOutputWithContext(ctx context.Context) SqlResourceSqlRoleDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlResourceSqlRoleDefinitionInput)(nil)).Elem(), &SqlResourceSqlRoleDefinition{})
	pulumi.RegisterOutputType(SqlResourceSqlRoleDefinitionOutput{})
}
