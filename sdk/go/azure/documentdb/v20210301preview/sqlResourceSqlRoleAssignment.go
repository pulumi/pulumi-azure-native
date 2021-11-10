


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlResourceSqlRoleAssignment struct {
	pulumi.CustomResourceState

	Name             pulumi.StringOutput    `pulumi:"name"`
	PrincipalId      pulumi.StringPtrOutput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringPtrOutput `pulumi:"roleDefinitionId"`
	Scope            pulumi.StringPtrOutput `pulumi:"scope"`
	Type             pulumi.StringOutput    `pulumi:"type"`
}


func NewSqlResourceSqlRoleAssignment(ctx *pulumi.Context,
	name string, args *SqlResourceSqlRoleAssignmentArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlRoleAssignment, error) {
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
			Type: pulumi.String("azure-native:documentdb:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:SqlResourceSqlRoleAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlResourceSqlRoleAssignment
	err := ctx.RegisterResource("azure-native:documentdb/v20210301preview:SqlResourceSqlRoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlResourceSqlRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlRoleAssignmentState, opts ...pulumi.ResourceOption) (*SqlResourceSqlRoleAssignment, error) {
	var resource SqlResourceSqlRoleAssignment
	err := ctx.ReadResource("azure-native:documentdb/v20210301preview:SqlResourceSqlRoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlResourceSqlRoleAssignmentState struct {
}

type SqlResourceSqlRoleAssignmentState struct {
}

func (SqlResourceSqlRoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlRoleAssignmentState)(nil)).Elem()
}

type sqlResourceSqlRoleAssignmentArgs struct {
	AccountName       string  `pulumi:"accountName"`
	PrincipalId       *string `pulumi:"principalId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RoleAssignmentId  *string `pulumi:"roleAssignmentId"`
	RoleDefinitionId  *string `pulumi:"roleDefinitionId"`
	Scope             *string `pulumi:"scope"`
}


type SqlResourceSqlRoleAssignmentArgs struct {
	AccountName       pulumi.StringInput
	PrincipalId       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RoleAssignmentId  pulumi.StringPtrInput
	RoleDefinitionId  pulumi.StringPtrInput
	Scope             pulumi.StringPtrInput
}

func (SqlResourceSqlRoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlRoleAssignmentArgs)(nil)).Elem()
}

type SqlResourceSqlRoleAssignmentInput interface {
	pulumi.Input

	ToSqlResourceSqlRoleAssignmentOutput() SqlResourceSqlRoleAssignmentOutput
	ToSqlResourceSqlRoleAssignmentOutputWithContext(ctx context.Context) SqlResourceSqlRoleAssignmentOutput
}

func (*SqlResourceSqlRoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlResourceSqlRoleAssignment)(nil))
}

func (i *SqlResourceSqlRoleAssignment) ToSqlResourceSqlRoleAssignmentOutput() SqlResourceSqlRoleAssignmentOutput {
	return i.ToSqlResourceSqlRoleAssignmentOutputWithContext(context.Background())
}

func (i *SqlResourceSqlRoleAssignment) ToSqlResourceSqlRoleAssignmentOutputWithContext(ctx context.Context) SqlResourceSqlRoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlRoleAssignmentOutput)
}

type SqlResourceSqlRoleAssignmentOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlRoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlResourceSqlRoleAssignment)(nil))
}

func (o SqlResourceSqlRoleAssignmentOutput) ToSqlResourceSqlRoleAssignmentOutput() SqlResourceSqlRoleAssignmentOutput {
	return o
}

func (o SqlResourceSqlRoleAssignmentOutput) ToSqlResourceSqlRoleAssignmentOutputWithContext(ctx context.Context) SqlResourceSqlRoleAssignmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlRoleAssignmentOutput{})
}
