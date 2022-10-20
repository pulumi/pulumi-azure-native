


package v20220215preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MongoDBResourceMongoRoleDefinition struct {
	pulumi.CustomResourceState

	DatabaseName pulumi.StringPtrOutput       `pulumi:"databaseName"`
	Name         pulumi.StringOutput          `pulumi:"name"`
	Privileges   PrivilegeResponseArrayOutput `pulumi:"privileges"`
	RoleName     pulumi.StringPtrOutput       `pulumi:"roleName"`
	Roles        RoleResponseArrayOutput      `pulumi:"roles"`
	Type         pulumi.StringOutput          `pulumi:"type"`
}


func NewMongoDBResourceMongoRoleDefinition(ctx *pulumi.Context,
	name string, args *MongoDBResourceMongoRoleDefinitionArgs, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoRoleDefinition, error) {
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
			Type: pulumi.String("azure-native:documentdb:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:MongoDBResourceMongoRoleDefinition"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:MongoDBResourceMongoRoleDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource MongoDBResourceMongoRoleDefinition
	err := ctx.RegisterResource("azure-native:documentdb/v20220215preview:MongoDBResourceMongoRoleDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMongoDBResourceMongoRoleDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoDBResourceMongoRoleDefinitionState, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoRoleDefinition, error) {
	var resource MongoDBResourceMongoRoleDefinition
	err := ctx.ReadResource("azure-native:documentdb/v20220215preview:MongoDBResourceMongoRoleDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mongoDBResourceMongoRoleDefinitionState struct {
}

type MongoDBResourceMongoRoleDefinitionState struct {
}

func (MongoDBResourceMongoRoleDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoRoleDefinitionState)(nil)).Elem()
}

type mongoDBResourceMongoRoleDefinitionArgs struct {
	AccountName           string                   `pulumi:"accountName"`
	DatabaseName          *string                  `pulumi:"databaseName"`
	MongoRoleDefinitionId *string                  `pulumi:"mongoRoleDefinitionId"`
	Privileges            []Privilege              `pulumi:"privileges"`
	ResourceGroupName     string                   `pulumi:"resourceGroupName"`
	RoleName              *string                  `pulumi:"roleName"`
	Roles                 []Role                   `pulumi:"roles"`
	Type                  *MongoRoleDefinitionType `pulumi:"type"`
}


type MongoDBResourceMongoRoleDefinitionArgs struct {
	AccountName           pulumi.StringInput
	DatabaseName          pulumi.StringPtrInput
	MongoRoleDefinitionId pulumi.StringPtrInput
	Privileges            PrivilegeArrayInput
	ResourceGroupName     pulumi.StringInput
	RoleName              pulumi.StringPtrInput
	Roles                 RoleArrayInput
	Type                  MongoRoleDefinitionTypePtrInput
}

func (MongoDBResourceMongoRoleDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoRoleDefinitionArgs)(nil)).Elem()
}

type MongoDBResourceMongoRoleDefinitionInput interface {
	pulumi.Input

	ToMongoDBResourceMongoRoleDefinitionOutput() MongoDBResourceMongoRoleDefinitionOutput
	ToMongoDBResourceMongoRoleDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoRoleDefinitionOutput
}

func (*MongoDBResourceMongoRoleDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBResourceMongoRoleDefinition)(nil)).Elem()
}

func (i *MongoDBResourceMongoRoleDefinition) ToMongoDBResourceMongoRoleDefinitionOutput() MongoDBResourceMongoRoleDefinitionOutput {
	return i.ToMongoDBResourceMongoRoleDefinitionOutputWithContext(context.Background())
}

func (i *MongoDBResourceMongoRoleDefinition) ToMongoDBResourceMongoRoleDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoRoleDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBResourceMongoRoleDefinitionOutput)
}

type MongoDBResourceMongoRoleDefinitionOutput struct{ *pulumi.OutputState }

func (MongoDBResourceMongoRoleDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBResourceMongoRoleDefinition)(nil)).Elem()
}

func (o MongoDBResourceMongoRoleDefinitionOutput) ToMongoDBResourceMongoRoleDefinitionOutput() MongoDBResourceMongoRoleDefinitionOutput {
	return o
}

func (o MongoDBResourceMongoRoleDefinitionOutput) ToMongoDBResourceMongoRoleDefinitionOutputWithContext(ctx context.Context) MongoDBResourceMongoRoleDefinitionOutput {
	return o
}

func (o MongoDBResourceMongoRoleDefinitionOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) pulumi.StringPtrOutput { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o MongoDBResourceMongoRoleDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MongoDBResourceMongoRoleDefinitionOutput) Privileges() PrivilegeResponseArrayOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) PrivilegeResponseArrayOutput { return v.Privileges }).(PrivilegeResponseArrayOutput)
}

func (o MongoDBResourceMongoRoleDefinitionOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) pulumi.StringPtrOutput { return v.RoleName }).(pulumi.StringPtrOutput)
}

func (o MongoDBResourceMongoRoleDefinitionOutput) Roles() RoleResponseArrayOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) RoleResponseArrayOutput { return v.Roles }).(RoleResponseArrayOutput)
}

func (o MongoDBResourceMongoRoleDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDBResourceMongoRoleDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MongoDBResourceMongoRoleDefinitionOutput{})
}
